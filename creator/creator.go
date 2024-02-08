package creator

import (
	_g "bytes"
	_e "encoding/xml"
	_fa "errors"
	_df "fmt"
	_c "image"
	_ae "io"
	_b "math"
	_ed "os"
	_p "path"
	_pf "path/filepath"
	_fag "regexp"
	_f "sort"
	_a "strconv"
	_dc "strings"
	_dg "text/template"
	_cd "unicode"

	_ca "bitbucket.org/shenghui0779/gopdf/common"
	_bdb "bitbucket.org/shenghui0779/gopdf/contentstream"
	_fc "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_fe "bitbucket.org/shenghui0779/gopdf/core"
	_cc "bitbucket.org/shenghui0779/gopdf/internal/graphic2d/svg"
	_ec "bitbucket.org/shenghui0779/gopdf/internal/integrations/unichart"
	_bd "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_ggc "bitbucket.org/shenghui0779/gopdf/model"
	_ee "github.com/gorilla/i18n/linebreak"
	_gg "github.com/unidoc/unichart/render"
	_cag "golang.org/x/text/unicode/bidi"
)

// Paragraph represents text drawn with a specified font and can wrap across lines and pages.
// By default it occupies the available width in the drawing context.
type Paragraph struct {
	_age          string
	_fggb         *_ggc.PdfFont
	_fcbfa        float64
	_dacae        float64
	_bged         Color
	_abd          TextAlignment
	_bebg         bool
	_dcdb         float64
	_bebd         int
	_ggad         bool
	_dgeg         float64
	_fgcbf        Margins
	_abea         Positioning
	_bcec         float64
	_cbcca        float64
	_geca, _bfeff float64
	_cbcf         []string
}

// SetPos sets absolute positioning with specified coordinates.
func (_bffd *Paragraph) SetPos(x, y float64) {
	_bffd._abea = PositionAbsolute
	_bffd._bcec = x
	_bffd._cbcca = y
}
func _dbfbb(_bbabb *_e.Decoder) (int, int) { return 0, 0 }

// SetFillColor sets the fill color.
func (_fdbfe *PolyBezierCurve) SetFillColor(color Color) {
	_fdbfe._bbaec = color
	_fdbfe._ffbb.FillColor = _dbac(color)
}

// SetBorderWidth sets the border width of the rectangle.
func (_abcge *Rectangle) SetBorderWidth(bw float64) { _abcge._decec = bw }

var PPMM = float64(72 * 1.0 / 25.4)

// ColorCMYKFrom8bit creates a Color from c,m,y,k values (0-100).
// Example:
//
//	red := ColorCMYKFrom8Bit(0, 100, 100, 0)
func ColorCMYKFrom8bit(c, m, y, k byte) Color {
	return cmykColor{_badc: _b.Min(float64(c), 100) / 100.0, _dea: _b.Min(float64(m), 100) / 100.0, _gaa: _b.Min(float64(y), 100) / 100.0, _ccb: _b.Min(float64(k), 100) / 100.0}
}

// Notes returns the notes section of the invoice as a title-content pair.
func (_fbbg *Invoice) Notes() (string, string) { return _fbbg._ecf[0], _fbbg._ecf[1] }
func (_dbeg *Image) makeXObject() error {
	_bfbb, _faa := _ggc.NewXObjectImageFromImage(_dbeg._dffd, nil, _dbeg._ggfc)
	if _faa != nil {
		_ca.Log.Error("\u0046\u0061\u0069le\u0064\u0020\u0074\u006f\u0020\u0063\u0072\u0065\u0061t\u0065 \u0078o\u0062j\u0065\u0063\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _faa)
		return _faa
	}
	_dbeg._edee = _bfbb
	return nil
}

// SetFillColor sets the fill color of the ellipse.
func (_aae *Ellipse) SetFillColor(col Color) { _aae._dbadf = col }

type listItem struct {
	_gdceb VectorDrawable
	_eeed  TextChunk
}

// SetAntiAlias enables anti alias config.
//
// Anti alias is disabled by default.
func (_bbcac *LinearShading) SetAntiAlias(enable bool) { _bbcac._aecc.SetAntiAlias(enable) }

// SetStyleTop sets border style for top side.
func (_fdc *border) SetStyleTop(style CellBorderStyle) { _fdc._deccc = style }

// NewPage adds a new Page to the Creator and sets as the active Page.
func (_fcac *Creator) NewPage() *_ggc.PdfPage {
	_efbe := _fcac.newPage()
	_fcac._cec = append(_fcac._cec, _efbe)
	_fcac._eacd.Page++
	return _efbe
}

// SetAngle sets the rotation angle of the text.
func (_cfbf *StyledParagraph) SetAngle(angle float64) { _cfbf._adag = angle }

// SetAnchor set gradient position anchor.
// Default to center.
func (_gdcfg *RadialShading) SetAnchor(anchor AnchorPoint) { _gdcfg._dedge = anchor }

// SetVerticalAlignment set the cell's vertical alignment of content.
// Can be one of:
// - CellHorizontalAlignmentTop
// - CellHorizontalAlignmentMiddle
// - CellHorizontalAlignmentBottom
func (_ccgge *TableCell) SetVerticalAlignment(valign CellVerticalAlignment) { _ccgge._fddca = valign }

// Line defines a line between point 1 (X1, Y1) and point 2 (X2, Y2).
// The line width, color, style (solid or dashed) and opacity can be
// configured. Implements the Drawable interface.
type Line struct {
	_daed  float64
	_gfaa  float64
	_eddg  float64
	_eabb  float64
	_cdgfa Color
	_ceac  _fc.LineStyle
	_ggbab float64
	_fgbfg []int64
	_bgfed int64
	_adf   float64
	_fgef  Positioning
	_afad  FitMode
	_bccc  Margins
}

// The Image type is used to draw an image onto PDF.
type Image struct {
	_edee         *_ggc.XObjectImage
	_dffd         *_ggc.Image
	_eadgc        float64
	_cagba, _efef float64
	_cegf, _bdff  float64
	_ebfa         Positioning
	_afdag        HorizontalAlignment
	_fce          float64
	_gdd          float64
	_feef         float64
	_cbea         Margins
	_dfdc, _afgg  float64
	_ggfc         _fe.StreamEncoder
	_geaf         FitMode
}

// Vertical returns total vertical (top + bottom) margin.
func (_bfaa *Margins) Vertical() float64 { return _bfaa.Bottom + _bfaa.Top }
func (_caba *Invoice) drawSection(_fccbg, _dedd string) []*StyledParagraph {
	var _ggef []*StyledParagraph
	if _fccbg != "" {
		_cgag := _egdc(_caba._daag)
		_cgag.SetMargins(0, 0, 0, 5)
		_cgag.Append(_fccbg)
		_ggef = append(_ggef, _cgag)
	}
	if _dedd != "" {
		_fcbb := _egdc(_caba._dcfab)
		_fcbb.Append(_dedd)
		_ggef = append(_ggef, _fcbb)
	}
	return _ggef
}

// SetBorderOpacity sets the border opacity.
func (_fcfd *Polygon) SetBorderOpacity(opacity float64) { _fcfd._dffbb = opacity }

// PageFinalize sets a function to be called for each page before finalization
// (i.e. the last stage of page processing before they get written out).
// The callback function allows final touch-ups for each page, and it
// provides information that might not be known at other stages of designing
// the document (e.g. the total number of pages). Unlike the header/footer
// functions, which are limited to the top/bottom margins of the page, the
// finalize function can be used draw components anywhere on the current page.
func (_ebad *Creator) PageFinalize(pageFinalizeFunc func(_aad PageFinalizeFunctionArgs) error) {
	_ebad._dcc = pageFinalizeFunc
}

// Height returns the current page height.
func (_eced *Creator) Height() float64 { return _eced._ffc }

// SetAngle would set the angle at which the gradient is rendered.
//
// The default angle would be 0 where the gradient would be rendered from left to right side.
func (_gada *LinearShading) SetAngle(angle float64) { _gada._fcfe = angle }

// TOCLine represents a line in a table of contents.
// The component can be used both in the context of a
// table of contents component and as a standalone component.
// The representation of a table of contents line is as follows:
/*
         [number] [title]      [separator] [page]
   e.g.: Chapter1 Introduction ........... 1
*/
type TOCLine struct {
	_ecde *StyledParagraph

	// Holds the text and style of the number part of the TOC line.
	Number TextChunk

	// Holds the text and style of the title part of the TOC line.
	Title TextChunk

	// Holds the text and style of the separator part of the TOC line.
	Separator TextChunk

	// Holds the text and style of the page part of the TOC line.
	Page   TextChunk
	_adbfe float64
	_daebd uint
	_aeeeg float64
	_dggfc Positioning
	_efdgc float64
	_cafda float64
	_eefcc int64
}

// NewTextStyle creates a new text style object which can be used to style
// chunks of text.
// Default attributes:
// Font: Helvetica
// Font size: 10
// Encoding: WinAnsiEncoding
// Text color: black
func (_cecc *Creator) NewTextStyle() TextStyle { return _eabad(_cecc._gade) }

// DrawHeader sets a function to draw a header on created output pages.
func (_facc *Creator) DrawHeader(drawHeaderFunc func(_cdfg *Block, _dbee HeaderFunctionArgs)) {
	_facc._fdcb = drawHeaderFunc
}

// Height returns the height of the graphic svg.
func (_dcfa *GraphicSVG) Height() float64 { return _dcfa._dgaf.Height }

// AddColorStop add color stop information for rendering gradient.
func (_adaa *shading) AddColorStop(color Color, point float64) {
	_adaa._ggcd = append(_adaa._ggcd, _cafcfa(color, point))
}

type containerDrawable interface {
	Drawable

	// ContainerComponent checks if the component is allowed to be added into provided 'container' and returns
	// preprocessed copy of itself. If the component is not changed it is allowed to return itself in a callback way.
	// If the component is not compatible with provided container this method should return an error.
	ContainerComponent(_aaaf Drawable) (Drawable, error)
}

// Width returns the width of the ellipse.
func (_ddeac *Ellipse) Width() float64 { return _ddeac._eded }

// SetBorderWidth sets the border width.
func (_gdbg *CurvePolygon) SetBorderWidth(borderWidth float64) { _gdbg._ecae.BorderWidth = borderWidth }

type cmykColor struct{ _badc, _dea, _gaa, _ccb float64 }

func (_bcg *Block) duplicate() *Block {
	_eb := &Block{}
	*_eb = *_bcg
	_fgc := _bdb.ContentStreamOperations{}
	_fgc = append(_fgc, *_bcg._cad...)
	_eb._cad = &_fgc
	return _eb
}

// SetInline sets the inline mode of the division.
func (_dggge *Division) SetInline(inline bool) { _dggge._aeadf = inline }
func _gfeg(_gfabd *templateProcessor, _gdeac *templateNode) (interface{}, error) {
	return _gfabd.parseListItem(_gdeac)
}

// Scale scales the ellipse dimensions by the specified factors.
func (_ffgfb *Ellipse) Scale(xFactor, yFactor float64) {
	_ffgfb._eded = xFactor * _ffgfb._eded
	_ffgfb._dabc = yFactor * _ffgfb._dabc
}

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_afaad *TOC) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gcbf := ctx
	_eaag, ctx, _fbadf := _afaad._ecagg.GeneratePageBlocks(ctx)
	if _fbadf != nil {
		return _eaag, ctx, _fbadf
	}
	for _, _eacca := range _afaad._fbeeg {
		_fbcc := _eacca._eefcc
		if !_afaad._dbaf {
			_eacca._eefcc = 0
		}
		_ddcce, _ddfae, _fbcgff := _eacca.GeneratePageBlocks(ctx)
		_eacca._eefcc = _fbcc
		if _fbcgff != nil {
			return _eaag, ctx, _fbcgff
		}
		if len(_ddcce) < 1 {
			continue
		}
		_eaag[len(_eaag)-1].mergeBlocks(_ddcce[0])
		_eaag = append(_eaag, _ddcce[1:]...)
		ctx = _ddfae
	}
	if _afaad._fcffe.IsRelative() {
		ctx.X = _gcbf.X
	}
	if _afaad._fcffe.IsAbsolute() {
		return _eaag, _gcbf, nil
	}
	return _eaag, ctx, nil
}

// SetTitleStyle sets the style properties of the invoice title.
func (_acbce *Invoice) SetTitleStyle(style TextStyle) { _acbce._eefc = style }

// InvoiceCellProps holds all style properties for an invoice cell.
type InvoiceCellProps struct {
	TextStyle       TextStyle
	Alignment       CellHorizontalAlignment
	BackgroundColor Color
	BorderColor     Color
	BorderWidth     float64
	BorderSides     []CellBorderSide
}

// SetLineOpacity sets the line opacity.
func (_gdff *Polyline) SetLineOpacity(opacity float64) { _gdff._afcb = opacity }

// Logo returns the logo of the invoice.
func (_agfb *Invoice) Logo() *Image { return _agfb._bfge }

// Date returns the invoice date description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_gfba *Invoice) Date() (*InvoiceCell, *InvoiceCell) { return _gfba._ddbe[0], _gfba._ddbe[1] }

// ToRGB implements interface Color.
// Note: It's not directly used since shading color works differently than regular color.
func (_abga *LinearShading) ToRGB() (float64, float64, float64) { return 0, 0, 0 }
func _bdcfa(_ecgff string) (*GraphicSVG, error) {
	_bcggb, _dedg := _cc.ParseFromString(_ecgff)
	if _dedg != nil {
		return nil, _dedg
	}
	return _cgee(_bcggb)
}

// EnablePageWrap controls whether the division is wrapped across pages.
// If disabled, the division is moved in its entirety on a new page, if it
// does not fit in the available height. By default, page wrapping is enabled.
// If the height of the division is larger than an entire page, wrapping is
// enabled automatically in order to avoid unwanted behavior.
// Currently, page wrapping can only be disabled for vertical divisions.
func (_acaa *Division) EnablePageWrap(enable bool) { _acaa._ggdbd = enable }

// NewImageFromFile creates an Image from a file.
func (_eggb *Creator) NewImageFromFile(path string) (*Image, error) { return _ggde(path) }

// FitMode defines resizing options of an object inside a container.
type FitMode int

// SetAddressStyle sets the style properties used to render the content of
// the invoice address sections.
func (_bgbfc *Invoice) SetAddressStyle(style TextStyle) { _bgbfc._acecg = style }
func _bebddd(_egee, _cacf, _cbade TextChunk, _aeefc uint, _acgeg TextStyle) *TOCLine {
	_gdee := _egdc(_acgeg)
	_gdee.SetEnableWrap(true)
	_gdee.SetTextAlignment(TextAlignmentLeft)
	_gdee.SetMargins(0, 0, 2, 2)
	_fedf := &TOCLine{_ecde: _gdee, Number: _egee, Title: _cacf, Page: _cbade, Separator: TextChunk{Text: "\u002e", Style: _acgeg}, _adbfe: 0, _daebd: _aeefc, _aeeeg: 10, _dggfc: PositionRelative}
	_gdee._fbgbc.Left = _fedf._adbfe + float64(_fedf._daebd-1)*_fedf._aeeeg
	_gdee._dbdfe = _fedf.prepareParagraph
	return _fedf
}

// GeneratePageBlocks generate the Page blocks.  Multiple blocks are generated if the contents wrap
// over multiple pages.
func (_eegd *Chapter) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_bad := ctx
	if _eegd._dddd.IsRelative() {
		ctx.X += _eegd._ecga.Left
		ctx.Y += _eegd._ecga.Top
		ctx.Width -= _eegd._ecga.Left + _eegd._ecga.Right
		ctx.Height -= _eegd._ecga.Top
	}
	_bge, _cdec, _cbed := _eegd._ddd.GeneratePageBlocks(ctx)
	if _cbed != nil {
		return _bge, ctx, _cbed
	}
	ctx = _cdec
	_eafa := ctx.X
	_faec := ctx.Y - _eegd._ddd.Height()
	_bcc := int64(ctx.Page)
	_gcdd := _eegd.headingNumber()
	_gfac := _eegd.headingText()
	if _eegd._afg {
		_ffe := _eegd._bdde.Add(_gcdd, _eegd._gbcb, _a.FormatInt(_bcc, 10), _eegd._ede)
		if _eegd._bdde._dbaf {
			_ffe.SetLink(_bcc, _eafa, _faec)
		}
	}
	if _eegd._agd == nil {
		_eegd._agd = _ggc.NewOutlineItem(_gfac, _ggc.NewOutlineDest(_bcc-1, _eafa, _faec))
		if _eegd._beae != nil {
			_eegd._beae._agd.Add(_eegd._agd)
		} else {
			_eegd._ggf.Add(_eegd._agd)
		}
	} else {
		_gcf := &_eegd._agd.Dest
		_gcf.Page = _bcc - 1
		_gcf.X = _eafa
		_gcf.Y = _faec
	}
	for _, _ffac := range _eegd._gbac {
		_gag, _cfdf, _badd := _ffac.GeneratePageBlocks(ctx)
		if _badd != nil {
			return _bge, ctx, _badd
		}
		if len(_gag) < 1 {
			continue
		}
		_bge[len(_bge)-1].mergeBlocks(_gag[0])
		_bge = append(_bge, _gag[1:]...)
		ctx = _cfdf
	}
	if _eegd._dddd.IsRelative() {
		ctx.X = _bad.X
	}
	if _eegd._dddd.IsAbsolute() {
		return _bge, _bad, nil
	}
	return _bge, ctx, nil
}

var (
	PageSizeA3     = PageSize{297 * PPMM, 420 * PPMM}
	PageSizeA4     = PageSize{210 * PPMM, 297 * PPMM}
	PageSizeA5     = PageSize{148 * PPMM, 210 * PPMM}
	PageSizeLetter = PageSize{8.5 * PPI, 11 * PPI}
	PageSizeLegal  = PageSize{8.5 * PPI, 14 * PPI}
)

// NewImage create a new image from a unidoc image (model.Image).
func (_bgba *Creator) NewImage(img *_ggc.Image) (*Image, error) { return _ggba(img) }

// ConvertToBinary converts current image data into binary (Bi-level image) format.
// If provided image is RGB or GrayScale the function converts it into binary image
// using histogram auto threshold method.
func (_gdg *Image) ConvertToBinary() error { return _gdg._dffd.ConvertToBinary() }

// SetHeaderRows turns the selected table rows into headers that are repeated
// for every page the table spans. startRow and endRow are inclusive.
func (_fdfdg *Table) SetHeaderRows(startRow, endRow int) error {
	if startRow <= 0 {
		return _fa.New("\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0073\u0074\u0061\u0072\u0074\u0020r\u006f\u0077\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	if endRow <= 0 {
		return _fa.New("\u0068\u0065a\u0064\u0065\u0072\u0020e\u006e\u0064 \u0072\u006f\u0077\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	if startRow > endRow {
		return _fa.New("\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0072\u006f\u0077\u0020\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0074\u0068\u0065 \u0065\u006e\u0064\u0020\u0072o\u0077")
	}
	_fdfdg._dgfg = true
	_fdfdg._aggfe = startRow
	_fdfdg._fecf = endRow
	return nil
}

// Height returns the height of the line.
func (_cfad *Line) Height() float64 {
	_bbfc := _cfad._adf
	if _cfad._daed == _cfad._eddg {
		_bbfc /= 2
	}
	return _b.Abs(_cfad._eabb-_cfad._gfaa) + _bbfc
}

// ScaleToHeight sets the graphic svg scaling factor with the given height.
func (_gadb *GraphicSVG) ScaleToHeight(h float64) {
	_eggg := _gadb._dgaf.Width / _gadb._dgaf.Height
	_gadb._dgaf.Height = h
	_gadb._dgaf.Width = h * _eggg
	_gadb._dgaf.SetScaling(_eggg, _eggg)
}

// SetPos sets the Table's positioning to absolute mode and specifies the upper-left corner
// coordinates as (x,y).
// Note that this is only sensible to use when the table does not wrap over multiple pages.
// TODO: Should be able to set width too (not just based on context/relative positioning mode).
func (_fgff *Table) SetPos(x, y float64) {
	_fgff._bbdf = PositionAbsolute
	_fgff._fedd = x
	_fgff._degda = y
}

// CreateTableOfContents sets a function to generate table of contents.
func (_bag *Creator) CreateTableOfContents(genTOCFunc func(_gfb *TOC) error) { _bag._gfgf = genTOCFunc }
func _facd(_dfg *Chapter, _gbd *TOC, _ecbf *_ggc.Outline, _fdcd string, _gcb int, _eafca TextStyle) *Chapter {
	var _bceg uint = 1
	if _dfg != nil {
		_bceg = _dfg._ede + 1
	}
	_eeeed := &Chapter{_cfde: _gcb, _gbcb: _fdcd, _fgg: true, _afg: true, _beae: _dfg, _bdde: _gbd, _ggf: _ecbf, _gbac: []Drawable{}, _ede: _bceg}
	_defg := _agbf(_eeeed.headingText(), _eafca)
	_defg.SetFont(_eafca.Font)
	_defg.SetFontSize(_eafca.FontSize)
	_eeeed._ddd = _defg
	return _eeeed
}

// GeneratePageBlocks draws the filled curve on page blocks.
func (_ccgd *FilledCurve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gcec := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_bbga, _, _bgbe := _ccgd.draw(_gcec, "")
	if _bgbe != nil {
		return nil, ctx, _bgbe
	}
	_bgbe = _gcec.addContentsByString(string(_bbga))
	if _bgbe != nil {
		return nil, ctx, _bgbe
	}
	return []*Block{_gcec}, ctx, nil
}

// GetMargins returns the margins of the graphic svg (left, right, top, bottom).
func (_baddc *GraphicSVG) GetMargins() (float64, float64, float64, float64) {
	return _baddc._eaff.Left, _baddc._eaff.Right, _baddc._eaff.Top, _baddc._eaff.Bottom
}

// SetWidth sets the the Paragraph width. This is essentially the wrapping width, i.e. the width the
// text can extend to prior to wrapping over to next line.
func (_cege *Paragraph) SetWidth(width float64) { _cege._dcdb = width; _cege.wrapText() }

// IsRelative checks if the positioning is relative.
func (_cgea Positioning) IsRelative() bool { return _cgea == PositionRelative }
func (_ddda *templateProcessor) parseTextOverflowAttr(_ceeb, _cbfc string) TextOverflow {
	_ca.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020\u0074e\u0078\u0074\u0020o\u0076\u0065\u0072\u0066\u006c\u006f\u0077\u0020\u0061tt\u0072\u0069\u0062u\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060,\u0020\u0025s\u0029\u002e", _ceeb, _cbfc)
	_effgdb := map[string]TextOverflow{"\u0076i\u0073\u0069\u0062\u006c\u0065": TextOverflowVisible, "\u0068\u0069\u0064\u0064\u0065\u006e": TextOverflowHidden}[_cbfc]
	return _effgdb
}

// GeneratePageBlocks draws the line on a new block representing the page.
// Implements the Drawable interface.
func (_adfb *Line) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var (
		_cfddd         []*Block
		_fgbcg         = NewBlock(ctx.PageWidth, ctx.PageHeight)
		_ddef          = ctx
		_faecg, _decce = _adfb._daed, ctx.PageHeight - _adfb._gfaa
		_fdceb, _aegb  = _adfb._eddg, ctx.PageHeight - _adfb._eabb
	)
	_bccd := _adfb._fgef.IsRelative()
	if _bccd {
		ctx.X += _adfb._bccc.Left
		ctx.Y += _adfb._bccc.Top
		ctx.Width -= _adfb._bccc.Left + _adfb._bccc.Right
		ctx.Height -= _adfb._bccc.Top + _adfb._bccc.Bottom
		_faecg, _decce, _fdceb, _aegb = _adfb.computeCoords(ctx)
		if _adfb.Height() > ctx.Height {
			_cfddd = append(_cfddd, _fgbcg)
			_fgbcg = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_cadba := ctx
			_cadba.Y = ctx.Margins.Top + _adfb._bccc.Top
			_cadba.X = ctx.Margins.Left + _adfb._bccc.Left
			_cadba.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _adfb._bccc.Top - _adfb._bccc.Bottom
			_cadba.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _adfb._bccc.Left - _adfb._bccc.Right
			ctx = _cadba
			_faecg, _decce, _fdceb, _aegb = _adfb.computeCoords(ctx)
		}
	}
	_fgdc := _fc.BasicLine{X1: _faecg, Y1: _decce, X2: _fdceb, Y2: _aegb, LineColor: _dbac(_adfb._cdgfa), Opacity: _adfb._ggbab, LineWidth: _adfb._adf, LineStyle: _adfb._ceac, DashArray: _adfb._fgbfg, DashPhase: _adfb._bgfed}
	_ffda, _gfdf := _fgbcg.setOpacity(1.0, _adfb._ggbab)
	if _gfdf != nil {
		return nil, ctx, _gfdf
	}
	_aceee, _, _gfdf := _fgdc.Draw(_ffda)
	if _gfdf != nil {
		return nil, ctx, _gfdf
	}
	if _gfdf = _fgbcg.addContentsByString(string(_aceee)); _gfdf != nil {
		return nil, ctx, _gfdf
	}
	if _bccd {
		ctx.X = _ddef.X
		ctx.Width = _ddef.Width
		_gdbca := _adfb.Height()
		ctx.Y += _gdbca + _adfb._bccc.Bottom
		ctx.Height -= _gdbca
	} else {
		ctx = _ddef
	}
	_cfddd = append(_cfddd, _fgbcg)
	return _cfddd, ctx, nil
}

// IsAbsolute checks if the positioning is absolute.
func (_ccfa Positioning) IsAbsolute() bool { return _ccfa == PositionAbsolute }

// NewTable create a new Table with a specified number of columns.
func (_bdc *Creator) NewTable(cols int) *Table { return _gdec(cols) }
func (_dedgc *templateProcessor) nodeError(_eecde *templateNode, _abbd string, _gfff ...interface{}) error {
	return _df.Errorf(_dedgc.getNodeErrorLocation(_eecde, _abbd, _gfff...))
}

// InvoiceCell represents any cell belonging to a table from the invoice
// template. The main tables are the invoice information table, the line
// items table and totals table. Contains the text value of the cell and
// the style properties of the cell.
type InvoiceCell struct {
	InvoiceCellProps
	Value string
}

func (_beca *Creator) wrapPageIfNeeded(_cffaa *_ggc.PdfPage) (*_ggc.PdfPage, error) {
	_agc, _fgba := _cffaa.GetAllContentStreams()
	if _fgba != nil {
		return nil, _fgba
	}
	_bcab := _bdb.NewContentStreamParser(_agc)
	_gccb, _fgba := _bcab.Parse()
	if _fgba != nil {
		return nil, _fgba
	}
	if !_gccb.HasUnclosedQ() {
		return nil, nil
	}
	_gccb.WrapIfNeeded()
	_fccg, _fgba := _fe.MakeStream(_gccb.Bytes(), _fe.NewFlateEncoder())
	if _fgba != nil {
		return nil, _fgba
	}
	_cffaa.Contents = _fe.MakeArray(_fccg)
	return _cffaa, nil
}
func _gcba(_efce *templateProcessor, _fcbae *templateNode) (interface{}, error) {
	return _efce.parseTableCell(_fcbae)
}

// SetBorderRadius sets the radius of the background corners.
func (_bg *Background) SetBorderRadius(topLeft, topRight, bottomLeft, bottomRight float64) {
	_bg.BorderRadiusTopLeft = topLeft
	_bg.BorderRadiusTopRight = topRight
	_bg.BorderRadiusBottomLeft = bottomLeft
	_bg.BorderRadiusBottomRight = bottomRight
}

// SetHeading sets the text and the style of the heading of the TOC component.
func (_adegea *TOC) SetHeading(text string, style TextStyle) {
	_ddddg := _adegea.Heading()
	_ddddg.Reset()
	_eeccb := _ddddg.Append(text)
	_eeccb.Style = style
}

// SetAngle sets Image rotation angle in degrees.
func (_dgbeb *Image) SetAngle(angle float64) { _dgbeb._eadgc = angle }
func _eabad(_faca *_ggc.PdfFont) TextStyle {
	return TextStyle{Color: ColorRGBFrom8bit(0, 0, 0), Font: _faca, FontSize: 10, OutlineSize: 1, HorizontalScaling: DefaultHorizontalScaling, UnderlineStyle: TextDecorationLineStyle{Offset: 1, Thickness: 1}}
}

// SetAddressHeadingStyle sets the style properties used to render the
// heading of the invoice address sections.
func (_bgfe *Invoice) SetAddressHeadingStyle(style TextStyle) { _bgfe._debc = style }
func (_abgd *Creator) newPage() *_ggc.PdfPage {
	_bde := _ggc.NewPdfPage()
	_cfc := _abgd._gccce[0]
	_abe := _abgd._gccce[1]
	_cda := _ggc.PdfRectangle{Llx: 0, Lly: 0, Urx: _cfc, Ury: _abe}
	_bde.MediaBox = &_cda
	_abgd._abf = _cfc
	_abgd._ffc = _abe
	_abgd.initContext()
	return _bde
}

// ScaleToWidth scale Image to a specified width w, maintaining the aspect ratio.
func (_bgeg *Image) ScaleToWidth(w float64) {
	_eecg := _bgeg._efef / _bgeg._cagba
	_bgeg._cagba = w
	_bgeg._efef = w * _eecg
}

// SetPageLabels adds the specified page labels to the PDF file generated
// by the creator. See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
// NOTE: for existing PDF files, the page label ranges object can be obtained
// using the model.PDFReader's GetPageLabels method.
func (_egf *Creator) SetPageLabels(pageLabels _fe.PdfObject) { _egf._dbgc = pageLabels }

// SetFillOpacity sets the fill opacity of the ellipse.
func (_dedc *Ellipse) SetFillOpacity(opacity float64) { _dedc._ddbf = opacity }

// NewStyledParagraph creates a new styled paragraph.
// Default attributes:
// Font: Helvetica,
// Font size: 10
// Encoding: WinAnsiEncoding
// Wrap: enabled
// Text color: black
func (_bcbdd *Creator) NewStyledParagraph() *StyledParagraph { return _egdc(_bcbdd.NewTextStyle()) }

// DashPattern returns the dash pattern of the line.
func (_ceda *Line) DashPattern() (_fdbe []int64, _ccaf int64) { return _ceda._fgbfg, _ceda._bgfed }

// DrawTemplate renders the template provided through the specified reader,
// using the specified `data` and `options`.
// Creator templates are first executed as text/template *Template instances,
// so the specified `data` is inserted within the template.
// The second phase of processing is actually parsing the template, translating
// it into creator components and rendering them using the provided options.
// Both the `data` and `options` parameters can be nil.
func (_fee *Block) DrawTemplate(c *Creator, r _ae.Reader, data interface{}, options *TemplateOptions) error {
	return _dced(c, r, data, options, _fee)
}

// SetMargins sets the margins of the rectangle.
// NOTE: rectangle margins are only applied if relative positioning is used.
func (_fddbd *Rectangle) SetMargins(left, right, top, bottom float64) {
	_fddbd._ebfce.Left = left
	_fddbd._ebfce.Right = right
	_fddbd._ebfce.Top = top
	_fddbd._ebfce.Bottom = bottom
}
func (_adcf *Invoice) generateInformationBlocks(_adb DrawContext) ([]*Block, DrawContext, error) {
	_agad := _egdc(_adcf._bdcd)
	_agad.SetMargins(0, 0, 0, 20)
	_acff := _adcf.drawAddress(_adcf._cfbbb)
	_acff = append(_acff, _agad)
	_acff = append(_acff, _adcf.drawAddress(_adcf._dcfc)...)
	_gcfb := _ffcc()
	for _, _bebf := range _acff {
		_gcfb.Add(_bebf)
	}
	_bgce := _adcf.drawInformation()
	_edeef := _gdec(2)
	_edeef.SetMargins(0, 0, 25, 0)
	_dfbf := _edeef.NewCell()
	_dfbf.SetIndent(0)
	_dfbf.SetContent(_gcfb)
	_dfbf = _edeef.NewCell()
	_dfbf.SetContent(_bgce)
	return _edeef.GeneratePageBlocks(_adb)
}

// WriteToFile writes the Creator output to file specified by path.
func (_gdccf *Creator) WriteToFile(outputPath string) error {
	abspath, _cagd := _pf.Abs(outputPath)
	if _cagd != nil {
		return _cagd
	}
	if _cagd = _ed.MkdirAll(_p.Dir(abspath), 0775); _cagd != nil {
		return _cagd
	}
	_bgc, _cagd := _ed.OpenFile(abspath, _ed.O_RDWR|_ed.O_CREATE|_ed.O_TRUNC, 0775)
	if _cagd != nil {
		return _cagd
	}
	defer _bgc.Close()
	return _gdccf.Write(_bgc)
}

// SetFillOpacity sets the fill opacity.
func (_abac *PolyBezierCurve) SetFillOpacity(opacity float64) { _abac._bcdgc = opacity }

// Positioning returns the type of positioning the rectangle is set to use.
func (_fcaae *Rectangle) Positioning() Positioning { return _fcaae._bbgac }

// SetColor sets the line color. Use ColorRGBFromHex, ColorRGBFrom8bit or
// ColorRGBFromArithmetic to create the color object.
func (_ceca *Line) SetColor(color Color) { _ceca._cdgfa = color }

// SetLineWidth sets the line width.
func (_beef *Line) SetLineWidth(width float64) { _beef._adf = width }
func (_aebfgc *templateProcessor) parseList(_efgdb *templateNode) (interface{}, error) {
	_bgfcf := _aebfgc.creator.NewList()
	for _, _cbgbc := range _efgdb._gbdee.Attr {
		_cbbe := _cbgbc.Value
		switch _bbba := _cbgbc.Name.Local; _bbba {
		case "\u0069\u006e\u0064\u0065\u006e\u0074":
			_bgfcf.SetIndent(_aebfgc.parseFloatAttr(_bbba, _cbbe))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_dcfcf := _aebfgc.parseMarginAttr(_bbba, _cbbe)
			_bgfcf.SetMargins(_dcfcf.Left, _dcfcf.Right, _dcfcf.Top, _dcfcf.Bottom)
		default:
			_aebfgc.nodeLogDebug(_efgdb, "\u0055\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u006c\u0069\u0073\u0074 \u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _bbba)
		}
	}
	return _bgfcf, nil
}
func _gdad() *FilledCurve {
	_cfab := FilledCurve{}
	_cfab._babe = []_fc.CubicBezierCurve{}
	return &_cfab
}

// NewRectangle creates a new rectangle with the left corner at (`x`, `y`),
// having the specified width and height.
// NOTE: In relative positioning mode, `x` and `y` are calculated using the
// current context. Furthermore, when the fit mode is set to fill the available
// space, the rectangle is scaled so that it occupies the entire context width
// while maintaining the original aspect ratio.
func (_baad *Creator) NewRectangle(x, y, width, height float64) *Rectangle {
	return _fbbgc(x, y, width, height)
}
func (_ffdg *Division) drawBackground(_gece []*Block, _dfd, _gbaa DrawContext, _eae bool) ([]*Block, error) {
	_gdccfe := len(_gece)
	if _gdccfe == 0 || _ffdg._fbgb == nil {
		return _gece, nil
	}
	_fcgb := make([]*Block, 0, len(_gece))
	for _baccg, _ddeg := range _gece {
		var (
			_aaad  = _ffdg._fbgb.BorderRadiusTopLeft
			_dfaec = _ffdg._fbgb.BorderRadiusTopRight
			_egb   = _ffdg._fbgb.BorderRadiusBottomLeft
			_acfg  = _ffdg._fbgb.BorderRadiusBottomRight
		)
		_ddfa := _dfd
		_ddfa.Page += _baccg
		if _baccg == 0 {
			if _eae {
				_fcgb = append(_fcgb, _ddeg)
				continue
			}
			if _gdccfe == 1 {
				_ddfa.Height = _gbaa.Y - _dfd.Y
			}
		} else {
			_ddfa.X = _ddfa.Margins.Left + _ffdg._debb.Left
			_ddfa.Y = _ddfa.Margins.Top
			_ddfa.Width = _ddfa.PageWidth - _ddfa.Margins.Left - _ddfa.Margins.Right - _ffdg._debb.Left - _ffdg._debb.Right
			if _baccg == _gdccfe-1 {
				_ddfa.Height = _gbaa.Y - _ddfa.Margins.Top - _ffdg._debb.Top
			} else {
				_ddfa.Height = _ddfa.PageHeight - _ddfa.Margins.Top - _ddfa.Margins.Bottom
			}
			if !_eae {
				_aaad = 0
				_dfaec = 0
			}
		}
		if _gdccfe > 1 && _baccg != _gdccfe-1 {
			_egb = 0
			_acfg = 0
		}
		_baeff := _fbbgc(_ddfa.X, _ddfa.Y, _ddfa.Width, _ddfa.Height)
		_baeff.SetFillColor(_ffdg._fbgb.FillColor)
		_baeff.SetBorderColor(_ffdg._fbgb.BorderColor)
		_baeff.SetBorderWidth(_ffdg._fbgb.BorderSize)
		_baeff.SetBorderRadius(_aaad, _dfaec, _egb, _acfg)
		_bfbfa, _, _eccab := _baeff.GeneratePageBlocks(_ddfa)
		if _eccab != nil {
			return nil, _eccab
		}
		if len(_bfbfa) == 0 {
			continue
		}
		_deba := _bfbfa[0]
		if _eccab = _deba.mergeBlocks(_ddeg); _eccab != nil {
			return nil, _eccab
		}
		_fcgb = append(_fcgb, _deba)
	}
	return _fcgb, nil
}

// SetMargins sets the margins for the Image (in relative mode): left, right, top, bottom.
func (_cgddd *Image) SetMargins(left, right, top, bottom float64) {
	_cgddd._cbea.Left = left
	_cgddd._cbea.Right = right
	_cgddd._cbea.Top = top
	_cgddd._cbea.Bottom = bottom
}

// ToRGB implements interface Color.
// Note: It's not directly used since shading color works differently than regular color.
func (_acge *RadialShading) ToRGB() (float64, float64, float64) { return 0, 0, 0 }

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

// Lines returns all the rows of the invoice line items table.
func (_adac *Invoice) Lines() [][]*InvoiceCell { return _adac._fbac }

// SetMargins sets the margins TOC line.
func (_gcfgb *TOCLine) SetMargins(left, right, top, bottom float64) {
	_gcfgb._adbfe = left
	_aeeea := &_gcfgb._ecde._fbgbc
	_aeeea.Left = _gcfgb._adbfe + float64(_gcfgb._daebd-1)*_gcfgb._aeeeg
	_aeeea.Right = right
	_aeeea.Top = top
	_aeeea.Bottom = bottom
}

// NewLinearGradientColor creates a linear gradient color that could act as a color in other components.
func (_cgdf *Creator) NewLinearGradientColor(colorPoints []*ColorPoint) *LinearShading {
	return _dgbfg(colorPoints)
}
func (_dabcg *StyledParagraph) getTextLineWidth(_gggeg []*TextChunk) float64 {
	var _aacf float64
	_gdfcb := len(_gggeg)
	for _bdfeg, _eaef := range _gggeg {
		_gdfed := &_eaef.Style
		_gcdf := len(_eaef.Text)
		for _ccdeb, _cgfed := range _eaef.Text {
			if _cgfed == '\u000A' {
				continue
			}
			_eecd, _cdbb := _gdfed.Font.GetRuneMetrics(_cgfed)
			if !_cdbb {
				_ca.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _cgfed)
				return -1
			}
			_aacf += _gdfed.FontSize * _eecd.Wx * _gdfed.horizontalScale()
			if _cgfed != ' ' && (_bdfeg != _gdfcb-1 || _ccdeb != _gcdf-1) {
				_aacf += _gdfed.CharSpacing * 1000.0
			}
		}
	}
	return _aacf
}

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_eddgg *List) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var _becb float64
	var _gaaag []*StyledParagraph
	for _, _aeed := range _eddgg._bede {
		_ccge := _egdc(_eddgg._dagb)
		_ccge.SetEnableWrap(false)
		_ccge.SetTextAlignment(TextAlignmentRight)
		_ccge.Append(_aeed._eeed.Text).Style = _aeed._eeed.Style
		_aadb := _ccge.getTextWidth() / 1000.0 / ctx.Width
		if _becb < _aadb {
			_becb = _aadb
		}
		_gaaag = append(_gaaag, _ccge)
	}
	_gffeg := _gdec(2)
	_gffeg.SetColumnWidths(_becb, 1-_becb)
	_gffeg.SetMargins(_eddgg._fed.Left+_eddgg._egab, _eddgg._fed.Right, _eddgg._fed.Top, _eddgg._fed.Bottom)
	_gffeg.EnableRowWrap(true)
	for _efgd, _afag := range _eddgg._bede {
		_cgaf := _gffeg.NewCell()
		_cgaf.SetIndent(0)
		_cgaf.SetContent(_gaaag[_efgd])
		_cgaf = _gffeg.NewCell()
		_cgaf.SetIndent(0)
		_cgaf.SetContent(_afag._gdceb)
	}
	return _gffeg.GeneratePageBlocks(ctx)
}

// FitMode returns the fit mode of the image.
func (_gfge *Image) FitMode() FitMode { return _gfge._geaf }
func (_abb *Ellipse) applyFitMode(_aadff float64) {
	_aadff -= _abb._cadg.Left + _abb._cadg.Right
	switch _abb._gbfg {
	case FitModeFillWidth:
		_abb.ScaleToWidth(_aadff)
	}
}
func _afac(_adgde float64, _ffca float64, _ddgdd float64, _aadeb float64, _gebfd []*ColorPoint) *RadialShading {
	return &RadialShading{_fcacc: &shading{_dedad: ColorWhite, _daff: false, _fgdd: []bool{false, false}, _ggcd: _gebfd}, _cafd: _adgde, _eagf: _ffca, _efdd: _ddgdd, _cagbac: _aadeb, _dedge: AnchorCenter}
}
func _fbbgc(_efcc, _fcfb, _ggadc, _ceaa float64) *Rectangle {
	return &Rectangle{_cgedd: _efcc, _eeff: _fcfb, _gfad: _ggadc, _fefg: _ceaa, _bbgac: PositionAbsolute, _geee: 1.0, _ecce: ColorBlack, _decec: 1.0, _ffgd: 1.0}
}

// GetMargins returns the margins of the rectangle: left, right, top, bottom.
func (_dbaga *Rectangle) GetMargins() (float64, float64, float64, float64) {
	return _dbaga._ebfce.Left, _dbaga._ebfce.Right, _dbaga._ebfce.Top, _dbaga._ebfce.Bottom
}

// Heading returns the heading component of the table of contents.
func (_fafd *TOC) Heading() *StyledParagraph { return _fafd._ecagg }

// SetBoundingBox set gradient color bounding box where the gradient would be rendered.
func (_efea *LinearShading) SetBoundingBox(x, y, width, height float64) {
	_efea._cbgf = &_ggc.PdfRectangle{Llx: x, Lly: y, Urx: x + width, Ury: y + height}
}

// MultiRowCell makes a new cell with the specified row span and inserts it
// into the table at the current position.
func (_ecdc *Table) MultiRowCell(rowspan int) *TableCell { return _ecdc.MultiCell(rowspan, 1) }

// AddColorStop add color stop info for rendering gradient color.
func (_agge *LinearShading) AddColorStop(color Color, point float64) {
	_agge._aecc.AddColorStop(color, point)
}

// GeneratePageBlocks draws the block contents on a template Page block.
// Implements the Drawable interface.
func (_cf *Block) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_cb := _bd.IdentityMatrix()
	_db, _gcg := _cf.Width(), _cf.Height()
	if _cf._ba.IsRelative() {
		_cb = _cb.Translate(ctx.X, ctx.PageHeight-ctx.Y-_gcg)
	} else {
		_cb = _cb.Translate(_cf._gb, ctx.PageHeight-_cf._gf-_gcg)
	}
	_fb := _gcg
	if _cf._de != 0 {
		_cb = _cb.Translate(_db/2, _gcg/2).Rotate(_cf._de*_b.Pi/180.0).Translate(-_db/2, -_gcg/2)
		_, _fb = _cf.RotatedSize()
	}
	if _cf._ba.IsRelative() {
		ctx.Y += _fb
	}
	_ef := _bdb.NewContentCreator()
	_ef.Add_cm(_cb[0], _cb[1], _cb[3], _cb[4], _cb[6], _cb[7])
	_gdf := _cf.duplicate()
	_ggd := append(*_ef.Operations(), *_gdf._cad...)
	_ggd.WrapIfNeeded()
	_gdf._cad = &_ggd
	for _, _cdb := range _cf._ga {
		_cfg, _aeb := _fe.GetArray(_cdb.Rect)
		if !_aeb || _cfg.Len() != 4 {
			_ca.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u0069\u006e\u0076\u0061\u006ci\u0064 \u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0052\u0065\u0063\u0074\u0020\u0066\u0069\u0065l\u0064\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _cdb.Rect)
			continue
		}
		_deb, _bgb := _ggc.NewPdfRectangle(*_cfg)
		if _bgb != nil {
			_ca.Log.Debug("\u0057A\u0052N\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074 \u0070\u0061\u0072\u0073e\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020\u0052\u0065\u0063\u0074\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061y\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u002e", _bgb)
			continue
		}
		_deb.Transform(_cb)
		_cdb.Rect = _deb.ToPdfObject()
	}
	return []*Block{_gdf}, ctx, nil
}

// Width returns the width of the chart. In relative positioning mode,
// all the available context width is used at render time.
func (_caff *Chart) Width() float64 { return float64(_caff._ebbf.Width()) }

// NewParagraph creates a new text paragraph.
// Default attributes:
// Font: Helvetica,
// Font size: 10
// Encoding: WinAnsiEncoding
// Wrap: enabled
// Text color: black
func (_aeff *Creator) NewParagraph(text string) *Paragraph { return _agbf(text, _aeff.NewTextStyle()) }

// TotalLines returns all the rows in the invoice totals table as
// description-value cell pairs.
func (_bcegb *Invoice) TotalLines() [][2]*InvoiceCell {
	_bfce := [][2]*InvoiceCell{_bcegb._befc}
	_bfce = append(_bfce, _bcegb._edb...)
	return append(_bfce, _bcegb._ada)
}

// ColorRGBFromArithmetic creates a Color from arithmetic color values (0-1).
// Example:
//
//	green := ColorRGBFromArithmetic(0.0, 1.0, 0.0)
func ColorRGBFromArithmetic(r, g, b float64) Color {
	return rgbColor{_efdc: _b.Max(_b.Min(r, 1.0), 0.0), _fgbf: _b.Max(_b.Min(g, 1.0), 0.0), _acbd: _b.Max(_b.Min(b, 1.0), 0.0)}
}

// SetPageMargins sets the page margins: left, right, top, bottom.
// The default page margins are 10% of document width.
func (_cbab *Creator) SetPageMargins(left, right, top, bottom float64) {
	_cbab._gcgd.Left = left
	_cbab._gcgd.Right = right
	_cbab._gcgd.Top = top
	_cbab._gcgd.Bottom = bottom
}

// AddressHeadingStyle returns the style properties used to render the
// heading of the invoice address sections.
func (_afcg *Invoice) AddressHeadingStyle() TextStyle { return _afcg._febc }

// Width returns the width of the line.
// NOTE: Depending on the fit mode the line is set to use, its width may be
// calculated at runtime (e.g. when using FitModeFillWidth).
func (_cbb *Line) Width() float64 { return _b.Abs(_cbb._eddg - _cbb._daed) }

// Number returns the invoice number description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_febb *Invoice) Number() (*InvoiceCell, *InvoiceCell) { return _febb._geeb[0], _febb._geeb[1] }

// AddExternalLink adds a new external link to the paragraph.
// The text parameter represents the text that is displayed and the url
// parameter sets the destionation of the link.
func (_adaf *StyledParagraph) AddExternalLink(text, url string) *TextChunk {
	_cdaad := NewTextChunk(text, _adaf._aecf)
	_cdaad._dfcb = _gfcae(url)
	return _adaf.appendChunk(_cdaad)
}

// ColorRGBFromHex converts color hex code to rgb color for using with creator.
// NOTE: If there is a problem interpreting the string, then will use black color and log a debug message.
// Example hex code: #ffffff -> (1,1,1) white.
func ColorRGBFromHex(hexStr string) Color {
	_fcca := rgbColor{}
	if (len(hexStr) != 4 && len(hexStr) != 7) || hexStr[0] != '#' {
		_ca.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
		return _fcca
	}
	var _debe, _cgcf, _cggg int
	if len(hexStr) == 4 {
		var _cbgb, _dfag, _bacb int
		_eaca, _egc := _df.Sscanf(hexStr, "\u0023\u0025\u0031\u0078\u0025\u0031\u0078\u0025\u0031\u0078", &_cbgb, &_dfag, &_bacb)
		if _egc != nil {
			_ca.Log.Debug("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020h\u0065\u0078\u0020\u0063\u006f\u0064\u0065:\u0020\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", hexStr, _egc)
			return _fcca
		}
		if _eaca != 3 {
			_ca.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
			return _fcca
		}
		_debe = _cbgb*16 + _cbgb
		_cgcf = _dfag*16 + _dfag
		_cggg = _bacb*16 + _bacb
	} else {
		_ecgg, _cace := _df.Sscanf(hexStr, "\u0023\u0025\u0032\u0078\u0025\u0032\u0078\u0025\u0032\u0078", &_debe, &_cgcf, &_cggg)
		if _cace != nil {
			_ca.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
			return _fcca
		}
		if _ecgg != 3 {
			_ca.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0073,\u0020\u006e\u0020\u0021\u003d\u0020\u0033 \u0028\u0025\u0064\u0029", hexStr, _ecgg)
			return _fcca
		}
	}
	_fgab := float64(_debe) / 255.0
	_bfbf := float64(_cgcf) / 255.0
	_ggdg := float64(_cggg) / 255.0
	_fcca._efdc = _fgab
	_fcca._fgbf = _bfbf
	_fcca._acbd = _ggdg
	return _fcca
}

// ScaleToHeight scales the Block to a specified height, maintaining the same aspect ratio.
func (_bf *Block) ScaleToHeight(h float64) { _cada := h / _bf._gfe; _bf.Scale(_cada, _cada) }

// GetMargins returns the left, right, top, bottom Margins.
func (_cfcc *Table) GetMargins() (float64, float64, float64, float64) {
	return _cfcc._gfbcc.Left, _cfcc._gfbcc.Right, _cfcc._gfbcc.Top, _cfcc._gfbcc.Bottom
}

// TextOverflow determines the behavior of paragraph text which does
// not fit in the available space.
type TextOverflow int

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
func (_effg *Creator) Finalize() error {
	if _effg._acead {
		return nil
	}
	_eabe := len(_effg._cec)
	_beag := 0
	if _effg._fbd != nil {
		_acbc := *_effg
		_effg._cec = nil
		_effg._bcba = nil
		_effg.initContext()
		_gcaf := FrontpageFunctionArgs{PageNum: 1, TotalPages: _eabe}
		_effg._fbd(_gcaf)
		_beag += len(_effg._cec)
		_effg._cec = _acbc._cec
		_effg._bcba = _acbc._bcba
	}
	if _effg.AddTOC {
		_effg.initContext()
		_effg._eacd.Page = _beag + 1
		if _effg.CustomTOC && _effg._gfgf != nil {
			_bcf := *_effg
			_effg._cec = nil
			_effg._bcba = nil
			if _facg := _effg._gfgf(_effg._effa); _facg != nil {
				return _facg
			}
			_beag += len(_effg._cec)
			_effg._cec = _bcf._cec
			_effg._bcba = _bcf._bcba
		} else {
			if _effg._gfgf != nil {
				if _fdd := _effg._gfgf(_effg._effa); _fdd != nil {
					return _fdd
				}
			}
			_gdb, _, _fffd := _effg._effa.GeneratePageBlocks(_effg._eacd)
			if _fffd != nil {
				_ca.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0065\u0020\u0062\u006c\u006f\u0063\u006b\u0073: \u0025\u0076", _fffd)
				return _fffd
			}
			_beag += len(_gdb)
		}
		_cdc := _effg._effa.Lines()
		for _, _fge := range _cdc {
			_cce, _caa := _a.Atoi(_fge.Page.Text)
			if _caa != nil {
				continue
			}
			_fge.Page.Text = _a.Itoa(_cce + _beag)
			_fge._eefcc += int64(_beag)
		}
	}
	_dfgd := false
	var _eabc []*_ggc.PdfPage
	if _effg._fbd != nil {
		_geae := *_effg
		_effg._cec = nil
		_effg._bcba = nil
		_cfdd := FrontpageFunctionArgs{PageNum: 1, TotalPages: _eabe}
		_effg._fbd(_cfdd)
		_eabe += len(_effg._cec)
		_eabc = _effg._cec
		_effg._cec = append(_effg._cec, _geae._cec...)
		_effg._bcba = _geae._bcba
		_dfgd = true
	}
	var _agce []*_ggc.PdfPage
	if _effg.AddTOC {
		_effg.initContext()
		if _effg.CustomTOC && _effg._gfgf != nil {
			_cafca := *_effg
			_effg._cec = nil
			_effg._bcba = nil
			if _bcfg := _effg._gfgf(_effg._effa); _bcfg != nil {
				_ca.Log.Debug("\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074i\u006e\u0067\u0020\u0054\u004f\u0043\u003a\u0020\u0025\u0076", _bcfg)
				return _bcfg
			}
			_agce = _effg._cec
			_eabe += len(_agce)
			_effg._cec = _cafca._cec
			_effg._bcba = _cafca._bcba
		} else {
			if _effg._gfgf != nil {
				if _edef := _effg._gfgf(_effg._effa); _edef != nil {
					_ca.Log.Debug("\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074i\u006e\u0067\u0020\u0054\u004f\u0043\u003a\u0020\u0025\u0076", _edef)
					return _edef
				}
			}
			_deff, _, _ := _effg._effa.GeneratePageBlocks(_effg._eacd)
			for _, _fbce := range _deff {
				_fbce.SetPos(0, 0)
				_eabe++
				_accgd := _effg.newPage()
				_agce = append(_agce, _accgd)
				_effg.setActivePage(_accgd)
				_effg.Draw(_fbce)
			}
		}
		if _dfgd {
			_deffg := _eabc
			_bfae := _effg._cec[len(_eabc):]
			_effg._cec = append([]*_ggc.PdfPage{}, _deffg...)
			_effg._cec = append(_effg._cec, _agce...)
			_effg._cec = append(_effg._cec, _bfae...)
		} else {
			_effg._cec = append(_agce, _effg._cec...)
		}
	}
	if _effg._gaec != nil && _effg.AddOutlines {
		var _dcac func(_fda *_ggc.OutlineItem)
		_dcac = func(_bba *_ggc.OutlineItem) {
			_bba.Dest.Page += int64(_beag)
			if _fba := int(_bba.Dest.Page); _fba >= 0 && _fba < len(_effg._cec) {
				_bba.Dest.PageObj = _effg._cec[_fba].GetPageAsIndirectObject()
			} else {
				_ca.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _fba)
			}
			_bba.Dest.Y = _effg._ffc - _bba.Dest.Y
			_gdae := _bba.Items()
			for _, _fbfd := range _gdae {
				_dcac(_fbfd)
			}
		}
		_gaaa := _effg._gaec.Items()
		for _, _acac := range _gaaa {
			_dcac(_acac)
		}
		if _effg.AddTOC {
			var _bbe int
			if _dfgd {
				_bbe = len(_eabc)
			}
			_adcb := _ggc.NewOutlineDest(int64(_bbe), 0, _effg._ffc)
			if _bbe >= 0 && _bbe < len(_effg._cec) {
				_adcb.PageObj = _effg._cec[_bbe].GetPageAsIndirectObject()
			} else {
				_ca.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _bbe)
			}
			_effg._gaec.Insert(0, _ggc.NewOutlineItem("\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0065\u006e\u0074\u0073", _adcb))
		}
	}
	for _cefe, _bcd := range _effg._cec {
		_effg.setActivePage(_bcd)
		if _effg._dcc != nil {
			_efcf, _fbgg, _dda := _bcd.Size()
			if _dda != nil {
				return _dda
			}
			_gdfa := PageFinalizeFunctionArgs{PageNum: _cefe + 1, PageWidth: _efcf, PageHeight: _fbgg, TOCPages: len(_agce), TotalPages: _eabe}
			if _bcea := _effg._dcc(_gdfa); _bcea != nil {
				_ca.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0070\u0061\u0067\u0065\u0020\u0066\u0069\u006e\u0061\u006c\u0069\u007a\u0065 \u0063\u0061\u006c\u006c\u0062\u0061\u0063k\u003a\u0020\u0025\u0076", _bcea)
				return _bcea
			}
		}
		if _effg._fdcb != nil {
			_bab := NewBlock(_effg._abf, _effg._gcgd.Top)
			_acf := HeaderFunctionArgs{PageNum: _cefe + 1, TotalPages: _eabe}
			_effg._fdcb(_bab, _acf)
			_bab.SetPos(0, 0)
			if _fgbc := _effg.Draw(_bab); _fgbc != nil {
				_ca.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069n\u0067 \u0068e\u0061\u0064\u0065\u0072\u003a\u0020\u0025v", _fgbc)
				return _fgbc
			}
		}
		if _effg._cgeab != nil {
			_edf := NewBlock(_effg._abf, _effg._gcgd.Bottom)
			_fdcec := FooterFunctionArgs{PageNum: _cefe + 1, TotalPages: _eabe}
			_effg._cgeab(_edf, _fdcec)
			_edf.SetPos(0, _effg._ffc-_edf._gfe)
			if _ebef := _effg.Draw(_edf); _ebef != nil {
				_ca.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069n\u0067 \u0066o\u006f\u0074\u0065\u0072\u003a\u0020\u0025v", _ebef)
				return _ebef
			}
		}
		_gcge, _dgbd := _effg._afbc[_bcd]
		if _aead, _bcfe := _effg._gfab[_bcd]; _bcfe {
			if _dgbd {
				_gcge.transformBlock(_aead)
			}
			if _cbd := _aead.drawToPage(_bcd); _cbd != nil {
				_ca.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069\u006e\u0067\u0020\u0070\u0061\u0067\u0065\u0020%\u0064\u0020\u0062\u006c\u006f\u0063\u006bs\u003a\u0020\u0025\u0076", _cefe+1, _cbd)
				return _cbd
			}
		}
		if _dgbd {
			if _cgfg := _gcge.transformPage(_bcd); _cgfg != nil {
				_ca.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0074\u0072\u0061\u006e\u0073f\u006f\u0072\u006d\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0076", _cgfg)
				return _cgfg
			}
		}
	}
	_effg._acead = true
	return nil
}

// NewFilledCurve returns a instance of filled curve.
func (_dfgc *Creator) NewFilledCurve() *FilledCurve { return _gdad() }
func _ffcc() *Division                              { return &Division{_ggdbd: true} }

// Chapter is used to arrange multiple drawables (paragraphs, images, etc) into a single section.
// The concept is the same as a book or a report chapter.
type Chapter struct {
	_cfde         int
	_gbcb         string
	_ddd          *Paragraph
	_gbac         []Drawable
	_bcbd         int
	_fgg          bool
	_afg          bool
	_dddd         Positioning
	_edcb, _bggdb float64
	_ecga         Margins
	_beae         *Chapter
	_bdde         *TOC
	_ggf          *_ggc.Outline
	_agd          *_ggc.OutlineItem
	_ede          uint
}

func (_fegdba *TOCLine) prepareParagraph(_ggab *StyledParagraph, _cdeec DrawContext) {
	_bbec := _fegdba.Title.Text
	if _fegdba.Number.Text != "" {
		_bbec = "\u0020" + _bbec
	}
	_bbec += "\u0020"
	_fegf := _fegdba.Page.Text
	if _fegf != "" {
		_fegf = "\u0020" + _fegf
	}
	_ggab._cdffa = []*TextChunk{{Text: _fegdba.Number.Text, Style: _fegdba.Number.Style, _dfcb: _fegdba.getLineLink()}, {Text: _bbec, Style: _fegdba.Title.Style, _dfcb: _fegdba.getLineLink()}, {Text: _fegf, Style: _fegdba.Page.Style, _dfcb: _fegdba.getLineLink()}}
	_ggab.wrapText()
	_degec := len(_ggab._aabba)
	if _degec == 0 {
		return
	}
	_ffgfe := _cdeec.Width*1000 - _ggab.getTextLineWidth(_ggab._aabba[_degec-1])
	_bdbd := _ggab.getTextLineWidth([]*TextChunk{&_fegdba.Separator})
	_adeggd := int(_ffgfe / _bdbd)
	_ecfgd := _dc.Repeat(_fegdba.Separator.Text, _adeggd)
	_gagfb := _fegdba.Separator.Style
	_abbc := _ggab.Insert(2, _ecfgd)
	_abbc.Style = _gagfb
	_abbc._dfcb = _fegdba.getLineLink()
	_ffgfe = _ffgfe - float64(_adeggd)*_bdbd
	if _ffgfe > 500 {
		_bgfdd, _edge := _gagfb.Font.GetRuneMetrics(' ')
		if _edge && _ffgfe > _bgfdd.Wx {
			_eafaf := int(_ffgfe / _bgfdd.Wx)
			if _eafaf > 0 {
				_dfeeb := _gagfb
				_dfeeb.FontSize = 1
				_abbc = _ggab.Insert(2, _dc.Repeat("\u0020", _eafaf))
				_abbc.Style = _dfeeb
				_abbc._dfcb = _fegdba.getLineLink()
			}
		}
	}
}
func _bgbee(_gcce *_ggc.PdfAnnotation) *_ggc.PdfAnnotation {
	if _gcce == nil {
		return nil
	}
	var _dbfg *_ggc.PdfAnnotation
	switch _fdaab := _gcce.GetContext().(type) {
	case *_ggc.PdfAnnotationLink:
		if _dbcg := _egedc(_fdaab); _dbcg != nil {
			_dbfg = _dbcg.PdfAnnotation
		}
	}
	return _dbfg
}
func (_gfd *Block) addContentsByString(_bae string) error {
	_bb := _bdb.NewContentStreamParser(_bae)
	_ccc, _fbf := _bb.Parse()
	if _fbf != nil {
		return _fbf
	}
	_gfd._cad.WrapIfNeeded()
	_ccc.WrapIfNeeded()
	*_gfd._cad = append(*_gfd._cad, *_ccc...)
	return nil
}

// GetMargins returns the margins of the chart (left, right, top, bottom).
func (_ffgf *Chart) GetMargins() (float64, float64, float64, float64) {
	return _ffgf._gcff.Left, _ffgf._gcff.Right, _ffgf._gcff.Top, _ffgf._gcff.Bottom
}

// Text sets the text content of the Paragraph.
func (_edbb *Paragraph) Text() string { return _edbb._age }

// LinearShading holds data for rendering a linear shading gradient.
type LinearShading struct {
	_aecc *shading
	_cbgf *_ggc.PdfRectangle
	_fcfe float64
}

func (_dedae *templateProcessor) parsePositioningAttr(_deccag, _ffcb string) Positioning {
	_ca.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0069\u006e\u0067\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060%\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _deccag, _ffcb)
	_cbbb := map[string]Positioning{"\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065": PositionRelative, "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065": PositionAbsolute}[_ffcb]
	return _cbbb
}

// CurvePolygon represents a curve polygon shape.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type CurvePolygon struct {
	_ecae  *_fc.CurvePolygon
	_gfage float64
	_fgcc  float64
	_fdge  Color
}

// SkipOver skips over a specified number of rows and cols.
func (_dbae *Table) SkipOver(rows, cols int) {
	_bgdaf := rows*_dbae._aeaa + cols - 1
	if _bgdaf < 0 {
		_ca.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	for _dgege := 0; _dgege < _bgdaf; _dgege++ {
		_dbae.NewCell()
	}
}

// SetStyleRight sets border style for right side.
func (_acbab *border) SetStyleRight(style CellBorderStyle) { _acbab._dabg = style }
func _daf(_bgbc string) string {
	_dca := _dac.FindAllString(_bgbc, -1)
	if len(_dca) == 0 {
		_bgbc = _bgbc + "\u0030"
	} else {
		_bac, _aebb := _a.Atoi(_dca[len(_dca)-1])
		if _aebb != nil {
			_ca.Log.Debug("\u0045r\u0072\u006f\u0072 \u0063\u006f\u006ev\u0065rt\u0069\u006e\u0067\u0020\u0064\u0069\u0067i\u0074\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020\u006e\u0061\u006de,\u0020f\u0061\u006c\u006c\u0062\u0061\u0063k\u0020\u0074\u006f\u0020\u0062a\u0073\u0069\u0063\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u003a \u0025\u0076", _aebb)
			_bgbc = _bgbc + "\u0030"
		} else {
			_bac++
			_gdce := _dc.LastIndex(_bgbc, _dca[len(_dca)-1])
			if _gdce == -1 {
				_bgbc = _df.Sprintf("\u0025\u0073\u0025\u0064", _bgbc[:len(_bgbc)-1], _bac)
			} else {
				_bgbc = _bgbc[:_gdce] + _a.Itoa(_bac)
			}
		}
	}
	return _bgbc
}
func _cbcc(_bfcc string) (*GraphicSVG, error) {
	_bcdae, _cdbef := _cc.ParseFromFile(_bfcc)
	if _cdbef != nil {
		return nil, _cdbef
	}
	return _cgee(_bcdae)
}

// AddLine adds a new line with the provided style to the table of contents.
func (_gbba *TOC) AddLine(line *TOCLine) *TOCLine {
	if line == nil {
		return nil
	}
	_gbba._fbeeg = append(_gbba._fbeeg, line)
	return line
}
func (_ded *Division) ctxHeight(_ecgd float64) float64 {
	_ecgd -= _ded._debb.Left + _ded._debb.Right + _ded._agda.Left + _ded._agda.Right
	var _gabe float64
	for _, _agga := range _ded._bfdf {
		_gabe += _cefg(_agga, _ecgd)
	}
	return _gabe
}
func _dbac(_ddb Color) _ggc.PdfColor {
	if _ddb == nil {
		_ddb = ColorBlack
	}
	switch _eebg := _ddb.(type) {
	case cmykColor:
		return _ggc.NewPdfColorDeviceCMYK(_eebg._badc, _eebg._dea, _eebg._gaa, _eebg._ccb)
	case *LinearShading:
		return _ggc.NewPdfColorPatternType2()
	case *RadialShading:
		return _ggc.NewPdfColorPatternType3()
	}
	return _ggc.NewPdfColorDeviceRGB(_ddb.ToRGB())
}

// GeneratePageBlocks draws the ellipse on a new block representing the page.
func (_bdcf *Ellipse) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var (
		_eaeb []*Block
		_fgge = NewBlock(ctx.PageWidth, ctx.PageHeight)
		_cfff = ctx
	)
	_ffb := _bdcf._acgd.IsRelative()
	if _ffb {
		_bdcf.applyFitMode(ctx.Width)
		ctx.X += _bdcf._cadg.Left
		ctx.Y += _bdcf._cadg.Top
		ctx.Width -= _bdcf._cadg.Left + _bdcf._cadg.Right
		ctx.Height -= _bdcf._cadg.Top + _bdcf._cadg.Bottom
		if _bdcf._dabc > ctx.Height {
			_eaeb = append(_eaeb, _fgge)
			_fgge = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_gbce := ctx
			_gbce.Y = ctx.Margins.Top + _bdcf._cadg.Top
			_gbce.X = ctx.Margins.Left + _bdcf._cadg.Left
			_gbce.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _bdcf._cadg.Top - _bdcf._cadg.Bottom
			_gbce.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _bdcf._cadg.Left - _bdcf._cadg.Right
			ctx = _gbce
		}
	} else {
		ctx.X = _bdcf._eebe - _bdcf._eded/2
		ctx.Y = _bdcf._beb - _bdcf._dabc/2
	}
	_beec := _fc.Circle{X: ctx.X, Y: ctx.PageHeight - ctx.Y - _bdcf._dabc, Width: _bdcf._eded, Height: _bdcf._dabc, BorderWidth: _bdcf._fgcca, Opacity: 1.0}
	if _bdcf._dbadf != nil {
		_beec.FillEnabled = true
		_eddd := _dbac(_bdcf._dbadf)
		_gage := _aede(_fgge, _eddd, _bdcf._dbadf, func() Rectangle {
			return Rectangle{_cgedd: _beec.X, _eeff: _beec.Y, _gfad: _beec.Width, _fefg: _beec.Height}
		})
		if _gage != nil {
			return nil, ctx, _gage
		}
		_beec.FillColor = _eddd
	}
	if _bdcf._fegd != nil {
		_beec.BorderEnabled = false
		if _bdcf._fgcca > 0 {
			_beec.BorderEnabled = true
		}
		_beec.BorderColor = _dbac(_bdcf._fegd)
		_beec.BorderWidth = _bdcf._fgcca
	}
	_gggf, _defad := _fgge.setOpacity(_bdcf._ddbf, _bdcf._dadc)
	if _defad != nil {
		return nil, ctx, _defad
	}
	_eef, _, _defad := _beec.Draw(_gggf)
	if _defad != nil {
		return nil, ctx, _defad
	}
	_defad = _fgge.addContentsByString(string(_eef))
	if _defad != nil {
		return nil, ctx, _defad
	}
	if _ffb {
		ctx.X = _cfff.X
		ctx.Width = _cfff.Width
		ctx.Y += _bdcf._dabc + _bdcf._cadg.Bottom
		ctx.Height -= _bdcf._dabc
	} else {
		ctx = _cfff
	}
	_eaeb = append(_eaeb, _fgge)
	return _eaeb, ctx, nil
}

// NewCurve returns new instance of Curve between points (x1,y1) and (x2, y2) with control point (cx,cy).
func (_gbdc *Creator) NewCurve(x1, y1, cx, cy, x2, y2 float64) *Curve {
	return _cdae(x1, y1, cx, cy, x2, y2)
}

const (
	HorizontalAlignmentLeft HorizontalAlignment = iota
	HorizontalAlignmentCenter
	HorizontalAlignmentRight
)
const (
	TextOverflowVisible TextOverflow = iota
	TextOverflowHidden
)

// NewImageFromData creates an Image from image data.
func (_cddab *Creator) NewImageFromData(data []byte) (*Image, error) { return _cdce(data) }

// SetDate sets the date of the invoice.
func (_cgba *Invoice) SetDate(date string) (*InvoiceCell, *InvoiceCell) {
	_cgba._ddbe[1].Value = date
	return _cgba._ddbe[0], _cgba._ddbe[1]
}

const (
	PositionRelative Positioning = iota
	PositionAbsolute
)

// SetMargins sets the margins of the line.
// NOTE: line margins are only applied if relative positioning is used.
func (_ccee *Line) SetMargins(left, right, top, bottom float64) {
	_ccee._bccc.Left = left
	_ccee._bccc.Right = right
	_ccee._bccc.Top = top
	_ccee._bccc.Bottom = bottom
}

// TextChunk represents a chunk of text along with a particular style.
type TextChunk struct {

	// The text that is being rendered in the PDF.
	Text string

	// The style of the text being rendered.
	Style  TextStyle
	_dfcb  *_ggc.PdfAnnotation
	_abacb bool

	// The vertical alignment of the text chunk.
	VerticalAlignment TextVerticalAlignment
}

// SetLineHeight sets the line height (1.0 default).
func (_dgdfd *StyledParagraph) SetLineHeight(lineheight float64) { _dgdfd._babcf = lineheight }
func (_eebgc *templateProcessor) parseChapterHeading(_fgggab *templateNode) (interface{}, error) {
	if _fgggab._fbcg == nil {
		_eebgc.nodeLogError(_fgggab, "\u0043\u0068a\u0070\u0074\u0065\u0072 \u0068\u0065a\u0064\u0069\u006e\u0067\u0020\u0070\u0061\u0072e\u006e\u0074\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c\u002e")
		return nil, _gfaba
	}
	_ggdd, _gdcada := _fgggab._fbcg._caacd.(*Chapter)
	if !_gdcada {
		_eebgc.nodeLogError(_fgggab, "\u0043h\u0061\u0070t\u0065\u0072\u0020h\u0065\u0061\u0064\u0069\u006e\u0067\u0020p\u0061\u0072\u0065\u006e\u0074\u0020(\u0025\u0054\u0029\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020a\u0020\u0063\u0068\u0061\u0070\u0074\u0065\u0072\u002e", _fgggab._fbcg._caacd)
		return nil, _gfaba
	}
	_bbge := _ggdd.GetHeading()
	if _, _bggg := _eebgc.parseParagraph(_fgggab, _bbge); _bggg != nil {
		return nil, _bggg
	}
	return _bbge, nil
}

// Terms returns the terms and conditions section of the invoice as a
// title-content pair.
func (_cgfe *Invoice) Terms() (string, string) { return _cgfe._aacd[0], _cgfe._aacd[1] }

type templateProcessor struct {
	creator *Creator
	_gcgb   []byte
	_affcb  *TemplateOptions
	_gcfaeb componentRenderer
	_eccabe string
}

func (_efgca *Table) sortCells() {
	_f.Slice(_efgca._cacca, func(_egba, _gfagd int) bool {
		_afde := _efgca._cacca[_egba]._deded
		_gdbec := _efgca._cacca[_gfagd]._deded
		if _afde < _gdbec {
			return true
		}
		if _afde > _gdbec {
			return false
		}
		return _efgca._cacca[_egba]._eafbd < _efgca._cacca[_gfagd]._eafbd
	})
}

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
func (_afb *Chapter) Add(d Drawable) error {
	if Drawable(_afb) == d {
		_ca.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u006e\u006f\u0074 \u0061\u0064\u0064\u0020\u0069\u0074\u0073\u0065\u006c\u0066")
		return _fa.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	switch _dgf := d.(type) {
	case *Paragraph, *StyledParagraph, *Image, *Chart, *Table, *Division, *List, *Rectangle, *Ellipse, *Line, *Block, *PageBreak, *Chapter:
		_afb._gbac = append(_afb._gbac, d)
	case containerDrawable:
		_cade, _dbca := _dgf.ContainerComponent(_afb)
		if _dbca != nil {
			return _dbca
		}
		_afb._gbac = append(_afb._gbac, _cade)
	default:
		_ca.Log.Debug("\u0055n\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u003a\u0020\u0025\u0054", d)
		return _fa.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return nil
}

// Positioning represents the positioning type for drawing creator components (relative/absolute).
type Positioning int

// Chart represents a chart drawable.
// It is used to render unichart chart components using a creator instance.
type Chart struct {
	_ebbf _gg.ChartRenderable
	_fbc  Positioning
	_bfe  float64
	_gbbc float64
	_gcff Margins
}

// DrawTemplate renders the template provided through the specified reader,
// using the specified `data` and `options`.
// Creator templates are first executed as text/template *Template instances,
// so the specified `data` is inserted within the template.
// The second phase of processing is actually parsing the template, translating
// it into creator components and rendering them using the provided options.
// Both the `data` and `options` parameters can be nil.
func (_cgd *Creator) DrawTemplate(r _ae.Reader, data interface{}, options *TemplateOptions) error {
	return _dced(_cgd, r, data, options, _cgd)
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

// Total returns the invoice total description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_eeeb *Invoice) Total() (*InvoiceCell, *InvoiceCell) { return _eeeb._ada[0], _eeeb._ada[1] }

// SetCoords sets the upper left corner coordinates of the rectangle.
func (_aadd *Rectangle) SetCoords(x, y float64) { _aadd._cgedd = x; _aadd._eeff = y }

// GeneratePageBlocks implements drawable interface.
func (_gecd *border) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_ebf := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_deg := _gecd._agf
	_fcad := ctx.PageHeight - _gecd._gea
	if _gecd._cde != nil {
		_geg := _fc.Rectangle{Opacity: 1.0, X: _gecd._agf, Y: ctx.PageHeight - _gecd._gea - _gecd._aebc, Height: _gecd._aebc, Width: _gecd._dcf}
		_geg.FillEnabled = true
		_acea := _dbac(_gecd._cde)
		_gff := _aede(_ebf, _acea, _gecd._cde, func() Rectangle {
			return Rectangle{_cgedd: _geg.X, _eeff: _geg.Y, _gfad: _geg.Width, _fefg: _geg.Height}
		})
		if _gff != nil {
			return nil, ctx, _gff
		}
		_geg.FillColor = _acea
		_geg.BorderEnabled = false
		_fcgd, _, _gff := _geg.Draw("")
		if _gff != nil {
			return nil, ctx, _gff
		}
		_gff = _ebf.addContentsByString(string(_fcgd))
		if _gff != nil {
			return nil, ctx, _gff
		}
	}
	_ddfb := _gecd._cdgaa
	_cgc := _gecd._gaf
	_dffg := _gecd._gcd
	_febe := _gecd._ebb
	_gffd := _gecd._cdgaa
	if _gecd._deccc == CellBorderStyleDouble {
		_gffd += 2 * _ddfb
	}
	_eea := _gecd._gaf
	if _gecd._cffe == CellBorderStyleDouble {
		_eea += 2 * _cgc
	}
	_edc := _gecd._gcd
	if _gecd._acba == CellBorderStyleDouble {
		_edc += 2 * _dffg
	}
	_cbg := _gecd._ebb
	if _gecd._dabg == CellBorderStyleDouble {
		_cbg += 2 * _febe
	}
	_fdbb := (_gffd - _edc) / 2
	_fdbf := (_gffd - _cbg) / 2
	_cfd := (_eea - _edc) / 2
	_abcg := (_eea - _cbg) / 2
	if _gecd._cdgaa != 0 {
		_efb := _deg
		_ebc := _fcad
		if _gecd._deccc == CellBorderStyleDouble {
			_ebc -= _ddfb
			_cacc := _fc.BasicLine{LineColor: _dbac(_gecd._feg), Opacity: 1.0, LineWidth: _gecd._cdgaa, LineStyle: _gecd.LineStyle, X1: _efb - _gffd/2 + _fdbb, Y1: _ebc + 2*_ddfb, X2: _efb + _gffd/2 - _fdbf + _gecd._dcf, Y2: _ebc + 2*_ddfb}
			_gfa, _, _fcc := _cacc.Draw("")
			if _fcc != nil {
				return nil, ctx, _fcc
			}
			_fcc = _ebf.addContentsByString(string(_gfa))
			if _fcc != nil {
				return nil, ctx, _fcc
			}
		}
		_cffad := _fc.BasicLine{LineWidth: _gecd._cdgaa, Opacity: 1.0, LineColor: _dbac(_gecd._feg), LineStyle: _gecd.LineStyle, X1: _efb - _gffd/2 + _fdbb + (_edc - _gecd._gcd), Y1: _ebc, X2: _efb + _gffd/2 - _fdbf + _gecd._dcf - (_cbg - _gecd._ebb), Y2: _ebc}
		_bggd, _, _gce := _cffad.Draw("")
		if _gce != nil {
			return nil, ctx, _gce
		}
		_gce = _ebf.addContentsByString(string(_bggd))
		if _gce != nil {
			return nil, ctx, _gce
		}
	}
	if _gecd._gaf != 0 {
		_fga := _deg
		_gfae := _fcad - _gecd._aebc
		if _gecd._cffe == CellBorderStyleDouble {
			_gfae += _cgc
			_faef := _fc.BasicLine{LineWidth: _gecd._gaf, Opacity: 1.0, LineColor: _dbac(_gecd._agg), LineStyle: _gecd.LineStyle, X1: _fga - _eea/2 + _cfd, Y1: _gfae - 2*_cgc, X2: _fga + _eea/2 - _abcg + _gecd._dcf, Y2: _gfae - 2*_cgc}
			_fcbg, _, _bbb := _faef.Draw("")
			if _bbb != nil {
				return nil, ctx, _bbb
			}
			_bbb = _ebf.addContentsByString(string(_fcbg))
			if _bbb != nil {
				return nil, ctx, _bbb
			}
		}
		_ebff := _fc.BasicLine{LineWidth: _gecd._gaf, Opacity: 1.0, LineColor: _dbac(_gecd._agg), LineStyle: _gecd.LineStyle, X1: _fga - _eea/2 + _cfd + (_edc - _gecd._gcd), Y1: _gfae, X2: _fga + _eea/2 - _abcg + _gecd._dcf - (_cbg - _gecd._ebb), Y2: _gfae}
		_cee, _, _aafa := _ebff.Draw("")
		if _aafa != nil {
			return nil, ctx, _aafa
		}
		_aafa = _ebf.addContentsByString(string(_cee))
		if _aafa != nil {
			return nil, ctx, _aafa
		}
	}
	if _gecd._gcd != 0 {
		_acc := _deg
		_ebab := _fcad
		if _gecd._acba == CellBorderStyleDouble {
			_acc += _dffg
			_daa := _fc.BasicLine{LineWidth: _gecd._gcd, Opacity: 1.0, LineColor: _dbac(_gecd._fcg), LineStyle: _gecd.LineStyle, X1: _acc - 2*_dffg, Y1: _ebab + _edc/2 + _fdbb, X2: _acc - 2*_dffg, Y2: _ebab - _edc/2 - _cfd - _gecd._aebc}
			_ecb, _, _afc := _daa.Draw("")
			if _afc != nil {
				return nil, ctx, _afc
			}
			_afc = _ebf.addContentsByString(string(_ecb))
			if _afc != nil {
				return nil, ctx, _afc
			}
		}
		_gafb := _fc.BasicLine{LineWidth: _gecd._gcd, Opacity: 1.0, LineColor: _dbac(_gecd._fcg), LineStyle: _gecd.LineStyle, X1: _acc, Y1: _ebab + _edc/2 + _fdbb - (_gffd - _gecd._cdgaa), X2: _acc, Y2: _ebab - _edc/2 - _cfd - _gecd._aebc + (_eea - _gecd._gaf)}
		_gead, _, _eeae := _gafb.Draw("")
		if _eeae != nil {
			return nil, ctx, _eeae
		}
		_eeae = _ebf.addContentsByString(string(_gead))
		if _eeae != nil {
			return nil, ctx, _eeae
		}
	}
	if _gecd._ebb != 0 {
		_acee := _deg + _gecd._dcf
		_ced := _fcad
		if _gecd._dabg == CellBorderStyleDouble {
			_acee -= _febe
			_afe := _fc.BasicLine{LineWidth: _gecd._ebb, Opacity: 1.0, LineColor: _dbac(_gecd._add), LineStyle: _gecd.LineStyle, X1: _acee + 2*_febe, Y1: _ced + _cbg/2 + _fdbf, X2: _acee + 2*_febe, Y2: _ced - _cbg/2 - _abcg - _gecd._aebc}
			_eabg, _, _fagd := _afe.Draw("")
			if _fagd != nil {
				return nil, ctx, _fagd
			}
			_fagd = _ebf.addContentsByString(string(_eabg))
			if _fagd != nil {
				return nil, ctx, _fagd
			}
		}
		_ffaf := _fc.BasicLine{LineWidth: _gecd._ebb, Opacity: 1.0, LineColor: _dbac(_gecd._add), LineStyle: _gecd.LineStyle, X1: _acee, Y1: _ced + _cbg/2 + _fdbf - (_gffd - _gecd._cdgaa), X2: _acee, Y2: _ced - _cbg/2 - _abcg - _gecd._aebc + (_eea - _gecd._gaf)}
		_cged, _, _daac := _ffaf.Draw("")
		if _daac != nil {
			return nil, ctx, _daac
		}
		_daac = _ebf.addContentsByString(string(_cged))
		if _daac != nil {
			return nil, ctx, _daac
		}
	}
	return []*Block{_ebf}, ctx, nil
}

// GeneratePageBlocks draws the polygon on a new block representing the page.
// Implements the Drawable interface.
func (_gcfd *Polygon) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_bbda := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_bcbed, _aeea := _bbda.setOpacity(_gcfd._befea, _gcfd._dffbb)
	if _aeea != nil {
		return nil, ctx, _aeea
	}
	_faefd := _gcfd._gbda
	_faefd.FillEnabled = _faefd.FillColor != nil
	_faefd.BorderEnabled = _faefd.BorderColor != nil && _faefd.BorderWidth > 0
	_bdaf := _faefd.Points
	_acdbe := _ggc.PdfRectangle{}
	_bcedc := false
	for _febcc := range _bdaf {
		for _gdfg := range _bdaf[_febcc] {
			_dbec := &_bdaf[_febcc][_gdfg]
			_dbec.Y = ctx.PageHeight - _dbec.Y
			if !_bcedc {
				_acdbe.Llx = _dbec.X
				_acdbe.Lly = _dbec.Y
				_acdbe.Urx = _dbec.X
				_acdbe.Ury = _dbec.Y
				_bcedc = true
			} else {
				_acdbe.Llx = _b.Min(_acdbe.Llx, _dbec.X)
				_acdbe.Lly = _b.Min(_acdbe.Lly, _dbec.Y)
				_acdbe.Urx = _b.Max(_acdbe.Urx, _dbec.X)
				_acdbe.Ury = _b.Max(_acdbe.Ury, _dbec.Y)
			}
		}
	}
	if _faefd.FillEnabled {
		_dbdd := _aede(_bbda, _gcfd._gbda.FillColor, _gcfd._ddcb, func() Rectangle {
			return Rectangle{_cgedd: _acdbe.Llx, _eeff: _acdbe.Lly, _gfad: _acdbe.Width(), _fefg: _acdbe.Height()}
		})
		if _dbdd != nil {
			return nil, ctx, _dbdd
		}
	}
	_gacd, _, _aeea := _faefd.Draw(_bcbed)
	if _aeea != nil {
		return nil, ctx, _aeea
	}
	if _aeea = _bbda.addContentsByString(string(_gacd)); _aeea != nil {
		return nil, ctx, _aeea
	}
	return []*Block{_bbda}, ctx, nil
}

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
func (_fbdc *Creator) SetPdfWriterAccessFunc(pdfWriterAccessFunc func(_effgd *_ggc.PdfWriter) error) {
	_fbdc._bggc = pdfWriterAccessFunc
}

// FitMode returns the fit mode of the rectangle.
func (_geeg *Rectangle) FitMode() FitMode { return _geeg._gaef }

// CellHorizontalAlignment defines the table cell's horizontal alignment.
type CellHorizontalAlignment int

func (_fdgb *pageTransformations) applyFlip(_cgce *_ggc.PdfPage) error {
	_ebg, _dbad := _fdgb._gfag, _fdgb._fcaa
	if !_ebg && !_dbad {
		return nil
	}
	if _cgce == nil {
		return _fa.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	_dfae, _gbad := _cgce.GetMediaBox()
	if _gbad != nil {
		return _gbad
	}
	_ecad, _cdfc := _dfae.Width(), _dfae.Height()
	_cfgf, _gbad := _cgce.GetRotate()
	if _gbad != nil {
		_ca.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _gbad.Error())
	}
	if _dfb := _cfgf%360 != 0 && _cfgf%90 == 0; _dfb {
		if _decf := (360 + _cfgf%360) % 360; _decf == 90 || _decf == 270 {
			_ebg, _dbad = _dbad, _ebg
		}
	}
	_bdf, _bda := 1.0, 0.0
	if _ebg {
		_bdf, _bda = -1.0, -_ecad
	}
	_cbgd, _caga := 1.0, 0.0
	if _dbad {
		_cbgd, _caga = -1.0, -_cdfc
	}
	_fdgbd := _bdb.NewContentCreator().Scale(_bdf, _cbgd).Translate(_bda, _caga)
	_aabb, _gbad := _fe.MakeStream(_fdgbd.Bytes(), _fe.NewFlateEncoder())
	if _gbad != nil {
		return _gbad
	}
	_dad := _fe.MakeArray(_aabb)
	_dad.Append(_cgce.GetContentStreamObjs()...)
	_cgce.Contents = _dad
	return nil
}

// SetOpacity sets the opacity of the line (0-1).
func (_fccf *Line) SetOpacity(opacity float64) { _fccf._ggbab = opacity }

// FitMode returns the fit mode of the ellipse.
func (_fecc *Ellipse) FitMode() FitMode { return _fecc._gbfg }

// SetMargins sets the margins of the component. The margins are applied
// around the division.
func (_afa *Division) SetMargins(left, right, top, bottom float64) {
	_afa._debb.Left = left
	_afa._debb.Right = right
	_afa._debb.Top = top
	_afa._debb.Bottom = bottom
}

// Padding returns the padding of the component.
func (_bbdc *Division) Padding() (_eedb, _gfgb, _fabb, _gceda float64) {
	return _bbdc._agda.Left, _bbdc._agda.Right, _bbdc._agda.Top, _bbdc._agda.Bottom
}
func (_aa *Block) addContents(_bgbg *_bdb.ContentStreamOperations) {
	_aa._cad.WrapIfNeeded()
	_bgbg.WrapIfNeeded()
	*_aa._cad = append(*_aa._cad, *_bgbg...)
}

// SetFontSize sets the font size in document units (points).
func (_dabf *Paragraph) SetFontSize(fontSize float64) { _dabf._fcbfa = fontSize }

// SetIndent sets the left offset of the list when nested into another list.
func (_egbb *List) SetIndent(indent float64) { _egbb._egab = indent; _egbb._ceeg = false }

// AddSection adds a new content section at the end of the invoice.
func (_efag *Invoice) AddSection(title, content string) {
	_efag._agdc = append(_efag._agdc, [2]string{title, content})
}

// GeneratePageBlocks generates the page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages. Implements the Drawable interface.
func (_bfedd *StyledParagraph) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_ccgg := ctx
	var _dcffb []*Block
	_fgac := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _bfedd._gbga.IsRelative() {
		ctx.X += _bfedd._fbgbc.Left
		ctx.Y += _bfedd._fbgbc.Top
		ctx.Width -= _bfedd._fbgbc.Left + _bfedd._fbgbc.Right
		ctx.Height -= _bfedd._fbgbc.Top
		_bfedd.SetWidth(ctx.Width)
	} else {
		if int(_bfedd._gcadd) <= 0 {
			_bfedd.SetWidth(_bfedd.getTextWidth() / 1000.0)
		}
		ctx.X = _bfedd._gcfda
		ctx.Y = _bfedd._cfge
	}
	if _bfedd._dbdfe != nil {
		_bfedd._dbdfe(_bfedd, ctx)
	}
	if _edeb := _bfedd.wrapText(); _edeb != nil {
		return nil, ctx, _edeb
	}
	_fcgbg := _bfedd._aabba
	for {
		_bdgdb, _agac, _efdfc := _cfdeg(_fgac, _bfedd, _fcgbg, ctx)
		if _efdfc != nil {
			_ca.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _efdfc)
			return nil, ctx, _efdfc
		}
		ctx = _bdgdb
		_dcffb = append(_dcffb, _fgac)
		if _fcgbg = _agac; len(_agac) == 0 {
			break
		}
		_fgac = NewBlock(ctx.PageWidth, ctx.PageHeight)
		ctx.Page++
		_bdgdb = ctx
		_bdgdb.Y = ctx.Margins.Top
		_bdgdb.X = ctx.Margins.Left + _bfedd._fbgbc.Left
		_bdgdb.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
		_bdgdb.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _bfedd._fbgbc.Left - _bfedd._fbgbc.Right
		ctx = _bdgdb
	}
	if _bfedd._gbga.IsRelative() {
		ctx.Y += _bfedd._fbgbc.Bottom
		ctx.Height -= _bfedd._fbgbc.Bottom
		if !ctx.Inline {
			ctx.X = _ccgg.X
			ctx.Width = _ccgg.Width
		}
		return _dcffb, ctx, nil
	}
	return _dcffb, _ccgg, nil
}

// SetTextAlignment sets the horizontal alignment of the text within the space provided.
func (_cdcc *Paragraph) SetTextAlignment(align TextAlignment) { _cdcc._abd = align }
func (_dbeed *Paragraph) getTextLineWidth(_eccg string) float64 {
	var _begf float64
	for _, _bgbfb := range _eccg {
		if _bgbfb == '\u000A' {
			continue
		}
		_fabec, _egdf := _dbeed._fggb.GetRuneMetrics(_bgbfb)
		if !_egdf {
			_ca.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052u\u006e\u0065\u0020\u0063\u0068a\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0028\u0072\u0075\u006e\u0065\u0020\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0029", _bgbfb, _bgbfb)
			return -1
		}
		_begf += _dbeed._fcbfa * _fabec.Wx
	}
	return _begf
}

// GeneratePageBlocks draws the rectangle on a new block representing the page. Implements the Drawable interface.
func (_agbfd *Rectangle) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var (
		_gcgf  []*Block
		_adfbd = NewBlock(ctx.PageWidth, ctx.PageHeight)
		_daae  = ctx
		_fedgc = _agbfd._decec / 2
	)
	_cbee := _agbfd._bbgac.IsRelative()
	if _cbee {
		_agbfd.applyFitMode(ctx.Width)
		ctx.X += _agbfd._ebfce.Left + _fedgc
		ctx.Y += _agbfd._ebfce.Top + _fedgc
		ctx.Width -= _agbfd._ebfce.Left + _agbfd._ebfce.Right
		ctx.Height -= _agbfd._ebfce.Top + _agbfd._ebfce.Bottom
		if _agbfd._fefg > ctx.Height {
			_gcgf = append(_gcgf, _adfbd)
			_adfbd = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_dgddb := ctx
			_dgddb.Y = ctx.Margins.Top + _agbfd._ebfce.Top + _fedgc
			_dgddb.X = ctx.Margins.Left + _agbfd._ebfce.Left + _fedgc
			_dgddb.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _agbfd._ebfce.Top - _agbfd._ebfce.Bottom
			_dgddb.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _agbfd._ebfce.Left - _agbfd._ebfce.Right
			ctx = _dgddb
		}
	} else {
		ctx.X = _agbfd._cgedd
		ctx.Y = _agbfd._eeff
	}
	_gaca := _fc.Rectangle{X: ctx.X, Y: ctx.PageHeight - ctx.Y - _agbfd._fefg, Width: _agbfd._gfad, Height: _agbfd._fefg, BorderRadiusTopLeft: _agbfd._gagef, BorderRadiusTopRight: _agbfd._efaga, BorderRadiusBottomLeft: _agbfd._fbcea, BorderRadiusBottomRight: _agbfd._fdcc, Opacity: 1.0}
	if _agbfd._bgca != nil {
		_gaca.FillEnabled = true
		_ecbc := _dbac(_agbfd._bgca)
		_gfeb := _aede(_adfbd, _ecbc, _agbfd._bgca, func() Rectangle {
			return Rectangle{_cgedd: _gaca.X, _eeff: _gaca.Y, _gfad: _gaca.Width, _fefg: _gaca.Height}
		})
		if _gfeb != nil {
			return nil, ctx, _gfeb
		}
		_gaca.FillColor = _ecbc
	}
	if _agbfd._ecce != nil && _agbfd._decec > 0 {
		_gaca.BorderEnabled = true
		_gaca.BorderColor = _dbac(_agbfd._ecce)
		_gaca.BorderWidth = _agbfd._decec
	}
	_ggcb, _bdcfag := _adfbd.setOpacity(_agbfd._geee, _agbfd._ffgd)
	if _bdcfag != nil {
		return nil, ctx, _bdcfag
	}
	_ceeac, _, _bdcfag := _gaca.Draw(_ggcb)
	if _bdcfag != nil {
		return nil, ctx, _bdcfag
	}
	if _bdcfag = _adfbd.addContentsByString(string(_ceeac)); _bdcfag != nil {
		return nil, ctx, _bdcfag
	}
	if _cbee {
		ctx.X = _daae.X
		ctx.Width = _daae.Width
		_aedf := _agbfd._fefg + _fedgc
		ctx.Y += _aedf + _agbfd._ebfce.Bottom
		ctx.Height -= _aedf
	} else {
		ctx = _daae
	}
	_gcgf = append(_gcgf, _adfbd)
	return _gcgf, ctx, nil
}
func _ddbg(_gcee _c.Image) (*Image, error) {
	_dddb, _acec := _ggc.ImageHandling.NewImageFromGoImage(_gcee)
	if _acec != nil {
		return nil, _acec
	}
	return _ggba(_dddb)
}

// SetDueDate sets the due date of the invoice.
func (_bcebg *Invoice) SetDueDate(dueDate string) (*InvoiceCell, *InvoiceCell) {
	_bcebg._fdfc[1].Value = dueDate
	return _bcebg._fdfc[0], _bcebg._fdfc[1]
}

// Margins returns the margins of the component.
func (_bcdf *Division) Margins() (_fgag, _geef, _bbgc, _caggb float64) {
	return _bcdf._debb.Left, _bcdf._debb.Right, _bcdf._debb.Top, _bcdf._debb.Bottom
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
func (_bcadf *TableCell) SetContent(vd VectorDrawable) error {
	switch _fcfa := vd.(type) {
	case *Paragraph:
		if _fcfa._ggad {
			_fcfa._bebg = true
		}
		_bcadf._efbbe = vd
	case *StyledParagraph:
		if _fcfa._facb {
			_fcfa._gfbb = true
		}
		_bcadf._efbbe = vd
	case *Image, *Chart, *Table, *Division, *List, *Rectangle, *Ellipse, *Line:
		_bcadf._efbbe = vd
	default:
		_ca.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0063e\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0079p\u0065\u0020\u0025\u0054", vd)
		return _fe.ErrTypeError
	}
	return nil
}

// Table allows organizing content in an rows X columns matrix, which can spawn across multiple pages.
type Table struct {
	_fgbfga       int
	_aeaa         int
	_ggfb         int
	_abbbf        []float64
	_begb         []float64
	_ddfaf        float64
	_cacca        []*TableCell
	_degfe        []int
	_bbdf         Positioning
	_fedd, _degda float64
	_gfbcc        Margins
	_dgfg         bool
	_aggfe        int
	_fecf         int
	_eebgf        bool
	_gbgc         bool
}

// Scale block by specified factors in the x and y directions.
func (_cff *Block) Scale(sx, sy float64) {
	_cddd := _bdb.NewContentCreator().Scale(sx, sy).Operations()
	*_cff._cad = append(*_cddd, *_cff._cad...)
	_cff._cad.WrapIfNeeded()
	_cff._ecd *= sx
	_cff._gfe *= sy
}

// SetBorderOpacity sets the border opacity of the rectangle.
func (_gaea *Rectangle) SetBorderOpacity(opacity float64) { _gaea._ffgd = opacity }

// MoveX moves the drawing context to absolute position x.
func (_dgdd *Creator) MoveX(x float64) { _dgdd._eacd.X = x }

// SetBackgroundColor set background color of the shading area.
//
// By default the background color is set to white.
func (_edec *shading) SetBackgroundColor(backgroundColor Color) { _edec._dedad = backgroundColor }

// SetBorderColor sets the border color of the ellipse.
func (_bgaee *Ellipse) SetBorderColor(col Color) { _bgaee._fegd = col }

// AddTextItem appends a new item with the specified text to the list.
// The method creates a styled paragraph with the specified text and returns
// it so that the item style can be customized.
// The method also returns the marker used for the newly added item.
// The marker object can be used to change the text and style of the marker
// for the current item.
func (_cgga *List) AddTextItem(text string) (*StyledParagraph, *TextChunk, error) {
	_fedg := _egdc(_cgga._dagb)
	_fedg.Append(text)
	_abbbc, _acaf := _cgga.Add(_fedg)
	return _fedg, _abbbc, _acaf
}

// SetTextVerticalAlignment sets the vertical alignment of the text within the
// bounds of the styled paragraph.
//
// Note: Currently Styled Paragraph doesn't support TextVerticalAlignmentBottom
// as that option only used for aligning text chunks.
//
// In order to change the vertical alignment of individual text chunks, use TextChunk.VerticalAlignment.
func (_eged *StyledParagraph) SetTextVerticalAlignment(align TextVerticalAlignment) {
	_eged._edbcb = align
}
func (_dbadc *Invoice) drawAddress(_dgee *InvoiceAddress) []*StyledParagraph {
	var _bfaf []*StyledParagraph
	if _dgee.Heading != "" {
		_abaf := _egdc(_dbadc._debc)
		_abaf.SetMargins(0, 0, 0, 7)
		_abaf.Append(_dgee.Heading)
		_bfaf = append(_bfaf, _abaf)
	}
	_ecac := _egdc(_dbadc._acecg)
	_ecac.SetLineHeight(1.2)
	_eafb := _dgee.Separator
	if _eafb == "" {
		_eafb = _dbadc._cedd
	}
	_fgbab := _dgee.City
	if _dgee.State != "" {
		if _fgbab != "" {
			_fgbab += _eafb
		}
		_fgbab += _dgee.State
	}
	if _dgee.Zip != "" {
		if _fgbab != "" {
			_fgbab += _eafb
		}
		_fgbab += _dgee.Zip
	}
	if _dgee.Name != "" {
		_ecac.Append(_dgee.Name + "\u000a")
	}
	if _dgee.Street != "" {
		_ecac.Append(_dgee.Street + "\u000a")
	}
	if _dgee.Street2 != "" {
		_ecac.Append(_dgee.Street2 + "\u000a")
	}
	if _fgbab != "" {
		_ecac.Append(_fgbab + "\u000a")
	}
	if _dgee.Country != "" {
		_ecac.Append(_dgee.Country + "\u000a")
	}
	_bdfe := _egdc(_dbadc._acecg)
	_bdfe.SetLineHeight(1.2)
	_bdfe.SetMargins(0, 0, 7, 0)
	if _dgee.Phone != "" {
		_bdfe.Append(_dgee.fmtLine(_dgee.Phone, "\u0050h\u006f\u006e\u0065\u003a\u0020", _dgee.HidePhoneLabel))
	}
	if _dgee.Email != "" {
		_bdfe.Append(_dgee.fmtLine(_dgee.Email, "\u0045m\u0061\u0069\u006c\u003a\u0020", _dgee.HideEmailLabel))
	}
	_bfaf = append(_bfaf, _ecac, _bdfe)
	return _bfaf
}
func (_gede *Table) wrapRow(_cagaa int, _bacbe DrawContext, _fabfbe float64) (bool, error) {
	if !_gede._eebgf {
		return false, nil
	}
	var (
		_egecc  = _gede._cacca[_cagaa]
		_bfgeb  = -1
		_adeb   []*TableCell
		_bbbc   float64
		_cdaadg bool
		_cdcge  = make([]float64, 0, len(_gede._abbbf))
	)
	_agbc := func(_dbcc *TableCell, _babcac VectorDrawable, _agfe bool) *TableCell {
		_cadd := *_dbcc
		_cadd._efbbe = _babcac
		if _agfe {
			_cadd._deded++
		}
		return &_cadd
	}
	_faggc := func(_agdg int, _bdcg VectorDrawable) {
		var _fegeb float64 = -1
		if _bdcg == nil {
			if _bbbb := _cdcge[_agdg-_cagaa]; _bbbb > _bacbe.Height {
				_bdcg = _gede._cacca[_agdg]._efbbe
				_gede._cacca[_agdg]._efbbe = nil
				_cdcge[_agdg-_cagaa] = 0
				_fegeb = _bbbb
			}
		}
		_deead := _agbc(_gede._cacca[_agdg], _bdcg, true)
		_adeb = append(_adeb, _deead)
		if _fegeb < 0 {
			_fegeb = _deead.height(_bacbe.Width)
		}
		if _fegeb > _bbbc {
			_bbbc = _fegeb
		}
	}
	for _efda := _cagaa; _efda < len(_gede._cacca); _efda++ {
		_bgcbc := _gede._cacca[_efda]
		if _egecc._deded != _bgcbc._deded {
			_bfgeb = _efda
			break
		}
		_bacbe.Width = _bgcbc.width(_gede._abbbf, _fabfbe)
		_cfda := _bgcbc.height(_bacbe.Width)
		var _aeeee VectorDrawable
		switch _dcffc := _bgcbc._efbbe.(type) {
		case *StyledParagraph:
			if _cfda > _bacbe.Height {
				_fbgd := _bacbe
				_fbgd.Height = _b.Floor(_bacbe.Height - _dcffc._fbgbc.Top - _dcffc._fbgbc.Bottom - 0.5*_dcffc.getTextHeight())
				_fbaeg, _fgfd, _cebce := _dcffc.split(_fbgd)
				if _cebce != nil {
					return false, _cebce
				}
				if _fbaeg != nil && _fgfd != nil {
					_dcffc = _fbaeg
					_bgcbc = _agbc(_bgcbc, _fbaeg, false)
					_gede._cacca[_efda] = _bgcbc
					_aeeee = _fgfd
					_cdaadg = true
				}
				_cfda = _bgcbc.height(_bacbe.Width)
			}
		case *Division:
			if _cfda > _bacbe.Height {
				_deegf := _bacbe
				_deegf.Height = _b.Floor(_bacbe.Height - _dcffc._debb.Top - _dcffc._debb.Bottom)
				_fcgge, _aeedc := _dcffc.split(_deegf)
				if _fcgge != nil && _aeedc != nil {
					_dcffc = _fcgge
					_bgcbc = _agbc(_bgcbc, _fcgge, false)
					_gede._cacca[_efda] = _bgcbc
					_aeeee = _aeedc
					_cdaadg = true
					if _fcgge._fbgb != nil {
						_fcgge._fbgb.BorderRadiusBottomLeft = 0
						_fcgge._fbgb.BorderRadiusBottomRight = 0
					}
					if _aeedc._fbgb != nil {
						_aeedc._fbgb.BorderRadiusTopLeft = 0
						_aeedc._fbgb.BorderRadiusTopRight = 0
					}
					_cfda = _bgcbc.height(_bacbe.Width)
				}
			}
		case *List:
			if _cfda > _bacbe.Height {
				_bcae := _bacbe
				_bcae.Height = _b.Floor(_bacbe.Height - _dcffc._fed.Vertical())
				_bebfg, _egggb := _dcffc.split(_bcae)
				if _bebfg != nil {
					_dcffc = _bebfg
					_bgcbc = _agbc(_bgcbc, _bebfg, false)
					_gede._cacca[_efda] = _bgcbc
				}
				if _egggb != nil {
					_aeeee = _egggb
					_cdaadg = true
				}
				_cfda = _bgcbc.height(_bacbe.Width)
			}
		}
		_cdcge = append(_cdcge, _cfda)
		if _cdaadg {
			if _adeb == nil {
				_adeb = make([]*TableCell, 0, len(_gede._abbbf))
				for _acgde := _cagaa; _acgde < _efda; _acgde++ {
					_faggc(_acgde, nil)
				}
			}
			_faggc(_efda, _aeeee)
		}
	}
	var _eacdg float64
	for _, _fead := range _cdcge {
		if _fead > _eacdg {
			_eacdg = _fead
		}
	}
	if _cdaadg && _eacdg < _bacbe.Height {
		if _bfgeb < 0 {
			_bfgeb = len(_gede._cacca)
		}
		_ebade := _gede._cacca[_bfgeb-1]._deded + _gede._cacca[_bfgeb-1]._ffgdb - 1
		for _febbd := _bfgeb; _febbd < len(_gede._cacca); _febbd++ {
			_gede._cacca[_febbd]._deded++
		}
		_gede._cacca = append(_gede._cacca[:_bfgeb], append(_adeb, _gede._cacca[_bfgeb:]...)...)
		_gede._begb = append(_gede._begb[:_ebade], append([]float64{_bbbc}, _gede._begb[_ebade:]...)...)
		_gede._begb[_egecc._deded+_egecc._ffgdb-2] = _eacdg
	}
	return _cdaadg, nil
}

// Write output of creator to io.Writer interface.
func (_fccb *Creator) Write(ws _ae.Writer) error {
	if _ebgc := _fccb.Finalize(); _ebgc != nil {
		return _ebgc
	}
	_gcde := _ggc.NewPdfWriter()
	_gcde.SetOptimizer(_fccb._eecf)
	if _fccb._ddc != nil {
		_degf := _gcde.SetForms(_fccb._ddc)
		if _degf != nil {
			_ca.Log.Debug("F\u0061\u0069\u006c\u0075\u0072\u0065\u003a\u0020\u0025\u0076", _degf)
			return _degf
		}
	}
	if _fccb._beeb != nil {
		_gcde.AddOutlineTree(_fccb._beeb)
	} else if _fccb._gaec != nil && _fccb.AddOutlines {
		_gcde.AddOutlineTree(&_fccb._gaec.ToPdfOutline().PdfOutlineTreeNode)
	}
	if _fccb._dbgc != nil {
		if _effb := _gcde.SetPageLabels(_fccb._dbgc); _effb != nil {
			_ca.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020C\u006f\u0075\u006c\u0064 no\u0074 s\u0065\u0074\u0020\u0070\u0061\u0067\u0065 l\u0061\u0062\u0065\u006c\u0073\u003a\u0020%\u0076", _effb)
			return _effb
		}
	}
	if _fccb._feff != nil {
		for _, _bdg := range _fccb._feff {
			_bgge := _bdg.SubsetRegistered()
			if _bgge != nil {
				_ca.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u006f\u0075\u006c\u0064\u0020\u006e\u006ft\u0020s\u0075\u0062\u0073\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0025\u0076", _bgge)
				return _bgge
			}
		}
	}
	if _fccb._bggc != nil {
		_gdaf := _fccb._bggc(&_gcde)
		if _gdaf != nil {
			_ca.Log.Debug("F\u0061\u0069\u006c\u0075\u0072\u0065\u003a\u0020\u0025\u0076", _gdaf)
			return _gdaf
		}
	}
	for _, _efbb := range _fccb._cec {
		_cded := _gcde.AddPage(_efbb)
		if _cded != nil {
			_ca.Log.Error("\u0046\u0061\u0069\u006ced\u0020\u0074\u006f\u0020\u0061\u0064\u0064\u0020\u0050\u0061\u0067\u0065\u003a\u0020%\u0076", _cded)
			return _cded
		}
	}
	_abef := _gcde.Write(ws)
	if _abef != nil {
		return _abef
	}
	return nil
}
func (_ffdec *templateProcessor) parseFontAttr(_fdcgfg, _gfcca string) *_ggc.PdfFont {
	_ca.Log.Debug("P\u0061\u0072\u0073\u0069\u006e\u0067 \u0066\u006f\u006e\u0074\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _fdcgfg, _gfcca)
	_bcac := _ffdec.creator._gade
	if _gfcca == "" {
		return _bcac
	}
	_gdfda := _dc.Split(_gfcca, "\u002c")
	for _, _aeae := range _gdfda {
		_aeae = _dc.TrimSpace(_aeae)
		if _aeae == "" {
			continue
		}
		_fgfed, _fdffb := _ffdec._affcb.FontMap[_gfcca]
		if _fdffb {
			return _fgfed
		}
		_agebd, _fdffb := map[string]_ggc.StdFontName{"\u0063o\u0075\u0072\u0069\u0065\u0072": _ggc.CourierName, "\u0063\u006f\u0075r\u0069\u0065\u0072\u002d\u0062\u006f\u006c\u0064": _ggc.CourierBoldName, "\u0063o\u0075r\u0069\u0065\u0072\u002d\u006f\u0062\u006c\u0069\u0071\u0075\u0065": _ggc.CourierObliqueName, "c\u006fu\u0072\u0069\u0065\u0072\u002d\u0062\u006f\u006cd\u002d\u006f\u0062\u006ciq\u0075\u0065": _ggc.CourierBoldObliqueName, "\u0068e\u006c\u0076\u0065\u0074\u0069\u0063a": _ggc.HelveticaName, "\u0068\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061-\u0062\u006f\u006c\u0064": _ggc.HelveticaBoldName, "\u0068\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061\u002d\u006f\u0062l\u0069\u0071\u0075\u0065": _ggc.HelveticaObliqueName, "\u0068\u0065\u006c\u0076et\u0069\u0063\u0061\u002d\u0062\u006f\u006c\u0064\u002d\u006f\u0062\u006c\u0069\u0071u\u0065": _ggc.HelveticaBoldObliqueName, "\u0073\u0079\u006d\u0062\u006f\u006c": _ggc.SymbolName, "\u007a\u0061\u0070\u0066\u002d\u0064\u0069\u006e\u0067\u0062\u0061\u0074\u0073": _ggc.ZapfDingbatsName, "\u0074\u0069\u006de\u0073": _ggc.TimesRomanName, "\u0074\u0069\u006d\u0065\u0073\u002d\u0062\u006f\u006c\u0064": _ggc.TimesBoldName, "\u0074\u0069\u006de\u0073\u002d\u0069\u0074\u0061\u006c\u0069\u0063": _ggc.TimesItalicName, "\u0074\u0069\u006d\u0065\u0073\u002d\u0062\u006f\u006c\u0064\u002d\u0069t\u0061\u006c\u0069\u0063": _ggc.TimesBoldItalicName}[_gfcca]
		if _fdffb {
			if _acaea, _bedeb := _ggc.NewStandard14Font(_agebd); _bedeb == nil {
				return _acaea
			}
		}
		if _bdffa := _ffdec.parseAttrPropList(_aeae); len(_bdffa) > 0 {
			if _gbbfb, _agba := _bdffa["\u0070\u0061\u0074\u0068"]; _agba {
				_aefb := _ggc.NewPdfFontFromTTFFile
				if _dcccf, _dfdg := _bdffa["\u0074\u0079\u0070\u0065"]; _dfdg && _dcccf == "\u0063o\u006d\u0070\u006f\u0073\u0069\u0074e" {
					_aefb = _ggc.NewCompositePdfFontFromTTFFile
				}
				if _cagdg, _ededf := _aefb(_gbbfb); _ededf != nil {
					_ca.Log.Debug("\u0043\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0060\u0025\u0073\u0060\u003a %\u0076\u002e", _gbbfb, _ededf)
				} else {
					return _cagdg
				}
			}
		}
	}
	return _bcac
}

// SetOutlineTree adds the specified outline tree to the PDF file generated
// by the creator. Adding an external outline tree disables the automatic
// generation of outlines done by the creator for the relevant components.
func (_dae *Creator) SetOutlineTree(outlineTree *_ggc.PdfOutlineTreeNode) { _dae._beeb = outlineTree }
func (_aaebf *Table) clone() *Table {
	_affb := *_aaebf
	_affb._begb = make([]float64, len(_aaebf._begb))
	copy(_affb._begb, _aaebf._begb)
	_affb._abbbf = make([]float64, len(_aaebf._abbbf))
	copy(_affb._abbbf, _aaebf._abbbf)
	_affb._cacca = make([]*TableCell, 0, len(_aaebf._cacca))
	for _, _dfbb := range _aaebf._cacca {
		_gegff := *_dfbb
		_gegff._ddggg = &_affb
		_affb._cacca = append(_affb._cacca, &_gegff)
	}
	return &_affb
}

// SetBuyerAddress sets the buyer address of the invoice.
func (_cgbf *Invoice) SetBuyerAddress(address *InvoiceAddress) { _cgbf._dcfc = address }

// AppendCurve appends a Bezier curve to the filled curve.
func (_dacc *FilledCurve) AppendCurve(curve _fc.CubicBezierCurve) *FilledCurve {
	_dacc._babe = append(_dacc._babe, curve)
	return _dacc
}
func (_adad *templateProcessor) parseListMarker(_deafc *templateNode) (interface{}, error) {
	if _deafc._fbcg == nil {
		_adad.nodeLogError(_deafc, "\u004c\u0069\u0073\u0074\u0020\u006da\u0072\u006b\u0065\u0072\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0063a\u006e\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c\u002e")
		return nil, _gfaba
	}
	var _eddbg *TextChunk
	switch _cgeb := _deafc._fbcg._caacd.(type) {
	case *List:
		_eddbg = &_cgeb._ebgba
	case *listItem:
		_eddbg = &_cgeb._eeed
	default:
		_adad.nodeLogError(_deafc, "\u0025\u0076 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u006e\u006f\u0064\u0065\u0020\u0066\u006f\u0072\u0020\u006c\u0069\u0073\u0074\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e", _cgeb)
		return nil, _gfaba
	}
	if _, _bcdb := _adad.parseTextChunk(_deafc, _eddbg); _bcdb != nil {
		_adad.nodeLogError(_deafc, "\u0043\u006f\u0075ld\u0020\u006e\u006f\u0074\u0020\u0070\u0061\u0072\u0073e\u0020l\u0069s\u0074 \u006d\u0061\u0072\u006b\u0065\u0072\u003a\u0020\u0060\u0025\u0076\u0060\u002e", _bcdb)
		return nil, nil
	}
	return _eddbg, nil
}

// SetNoteStyle sets the style properties used to render the content of the
// invoice note sections.
func (_gbbd *Invoice) SetNoteStyle(style TextStyle) { _gbbd._dcfab = style }

// AddressStyle returns the style properties used to render the content of
// the invoice address sections.
func (_badbc *Invoice) AddressStyle() TextStyle { return _badbc._acecg }

// String implements error interface.
func (_eada UnsupportedRuneError) Error() string { return _eada.Message }
func (_edcf *listItem) ctxHeight(_bgcee float64) float64 {
	var _gbdbf float64
	switch _eaaf := _edcf._gdceb.(type) {
	case *Paragraph:
		if _eaaf._bebg {
			_eaaf.SetWidth(_bgcee - _eaaf._fgcbf.Horizontal())
		}
		_gbdbf = _eaaf.Height() + _eaaf._fgcbf.Vertical()
		_gbdbf += 0.5 * _eaaf._fcbfa * _eaaf._dacae
	case *StyledParagraph:
		if _eaaf._gfbb {
			_eaaf.SetWidth(_bgcee - _eaaf._fbgbc.Horizontal())
		}
		_gbdbf = _eaaf.Height() + _eaaf._fbgbc.Vertical()
		_gbdbf += 0.5 * _eaaf.getTextHeight()
	case *List:
		_fgad := _bgcee - _edcf._eeed.Width() - _eaaf._fed.Horizontal() - _eaaf._egab
		_gbdbf = _eaaf.ctxHeight(_fgad) + _eaaf._fed.Vertical()
	case *Image:
		_gbdbf = _eaaf.Height() + _eaaf._cbea.Vertical()
	case *Division:
		_caced := _bgcee - _edcf._eeed.Width() - _eaaf._debb.Horizontal()
		_gbdbf = _eaaf.ctxHeight(_caced) + _eaaf._debb.Vertical()
	case *Table:
		_bbcba := _bgcee - _edcf._eeed.Width() - _eaaf._gfbcc.Horizontal()
		_eaaf.updateRowHeights(_bbcba)
		_gbdbf = _eaaf.Height() + _eaaf._gfbcc.Vertical()
	default:
		_gbdbf = _edcf._gdceb.Height()
	}
	return _gbdbf
}

// FooterFunctionArgs holds the input arguments to a footer drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type FooterFunctionArgs struct {
	PageNum    int
	TotalPages int
}

// Ellipse defines an ellipse with a center at (xc,yc) and a specified width and height.  The ellipse can have a colored
// fill and/or border with a specified width.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Ellipse struct {
	_eebe  float64
	_beb   float64
	_eded  float64
	_dabc  float64
	_acgd  Positioning
	_dbadf Color
	_ddbf  float64
	_fegd  Color
	_fgcca float64
	_dadc  float64
	_cadg  Margins
	_gbfg  FitMode
}

// NewTOC creates a new table of contents.
func (_edeg *Creator) NewTOC(title string) *TOC {
	_babc := _edeg.NewTextStyle()
	_babc.Font = _edeg._eacc
	return _dfede(title, _edeg.NewTextStyle(), _babc)
}

// TemplateOptions contains options and resources to use when rendering
// a template with a Creator instance.
// All the resources in the map fields can be referenced by their
// name/key in the template which is rendered using the options instance.
type TemplateOptions struct {

	// HelperFuncMap is used to define functions which can be accessed
	// inside the rendered templates by their assigned names.
	HelperFuncMap _dg.FuncMap

	// SubtemplateMap contains templates which can be rendered alongside
	// the main template. They can be accessed using their assigned names
	// in the main template or in the other subtemplates.
	// Subtemplates defined inside the subtemplates specified in the map
	// can be accessed directly.
	// All resources available to the main template are also available
	// to the subtemplates.
	SubtemplateMap map[string]_ae.Reader

	// FontMap contains pre-loaded fonts which can be accessed
	// inside the rendered templates by their assigned names.
	FontMap map[string]*_ggc.PdfFont

	// ImageMap contains pre-loaded images which can be accessed
	// inside the rendered templates by their assigned names.
	ImageMap map[string]*_ggc.Image

	// ColorMap contains colors which can be accessed
	// inside the rendered templates by their assigned names.
	ColorMap map[string]Color

	// ChartMap contains charts which can be accessed
	// inside the rendered templates by their assigned names.
	ChartMap map[string]_gg.ChartRenderable
}

func (_addec *templateProcessor) parseBorderRadiusAttr(_egfc, _gdffe string) (_dagdg, _cfbbd, _fafeb, _bgfa float64) {
	_ca.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020\u0062o\u0072\u0064\u0065r\u0020\u0072\u0061\u0064\u0069\u0075\u0073\u0020\u0061tt\u0072\u0069\u0062u\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060,\u0020\u0025s\u0029\u002e", _egfc, _gdffe)
	switch _gbaf := _dc.Fields(_gdffe); len(_gbaf) {
	case 1:
		_dagdg, _ = _a.ParseFloat(_gbaf[0], 64)
		_cfbbd = _dagdg
		_fafeb = _dagdg
		_bgfa = _dagdg
	case 2:
		_dagdg, _ = _a.ParseFloat(_gbaf[0], 64)
		_fafeb = _dagdg
		_cfbbd, _ = _a.ParseFloat(_gbaf[1], 64)
		_bgfa = _cfbbd
	case 3:
		_dagdg, _ = _a.ParseFloat(_gbaf[0], 64)
		_cfbbd, _ = _a.ParseFloat(_gbaf[1], 64)
		_bgfa = _cfbbd
		_fafeb, _ = _a.ParseFloat(_gbaf[2], 64)
	case 4:
		_dagdg, _ = _a.ParseFloat(_gbaf[0], 64)
		_cfbbd, _ = _a.ParseFloat(_gbaf[1], 64)
		_fafeb, _ = _a.ParseFloat(_gbaf[2], 64)
		_bgfa, _ = _a.ParseFloat(_gbaf[3], 64)
	}
	return _dagdg, _cfbbd, _fafeb, _bgfa
}

// Level returns the indentation level of the TOC line.
func (_bcbca *TOCLine) Level() uint { return _bcbca._daebd }

// BuyerAddress returns the buyer address used in the invoice template.
func (_fbe *Invoice) BuyerAddress() *InvoiceAddress { return _fbe._dcfc }

// DrawWithContext draws the Block using the specified drawing context.
func (_eda *Block) DrawWithContext(d Drawable, ctx DrawContext) error {
	_ccg, _, _cddg := d.GeneratePageBlocks(ctx)
	if _cddg != nil {
		return _cddg
	}
	if len(_ccg) != 1 {
		return ErrContentNotFit
	}
	for _, _dff := range _ccg {
		if _bef := _eda.mergeBlocks(_dff); _bef != nil {
			return _bef
		}
	}
	return nil
}
func _cddc(_ebdec string) ([]string, error) {
	var (
		_facbe []string
		_fdef  []rune
	)
	for _, _bdabc := range _ebdec {
		if _bdabc == '\u000A' {
			if len(_fdef) > 0 {
				_facbe = append(_facbe, string(_fdef))
			}
			_facbe = append(_facbe, string(_bdabc))
			_fdef = nil
			continue
		}
		_fdef = append(_fdef, _bdabc)
	}
	if len(_fdef) > 0 {
		_facbe = append(_facbe, string(_fdef))
	}
	var _eddbcd []string
	for _, _eebdg := range _facbe {
		_bacg := []rune(_eebdg)
		_dggb := _ee.NewScanner(_bacg)
		var _gefb []rune
		for _fecdgb := 0; _fecdgb < len(_bacg); _fecdgb++ {
			_, _gffbg, _agdaa := _dggb.Next()
			if _agdaa != nil {
				return nil, _agdaa
			}
			if _gffbg == _ee.BreakProhibited || _cd.IsSpace(_bacg[_fecdgb]) {
				_gefb = append(_gefb, _bacg[_fecdgb])
				if _cd.IsSpace(_bacg[_fecdgb]) {
					_eddbcd = append(_eddbcd, string(_gefb))
					_gefb = []rune{}
				}
				continue
			} else {
				if len(_gefb) > 0 {
					_eddbcd = append(_eddbcd, string(_gefb))
				}
				_gefb = []rune{_bacg[_fecdgb]}
			}
		}
		if len(_gefb) > 0 {
			_eddbcd = append(_eddbcd, string(_gefb))
		}
	}
	return _eddbcd, nil
}

// ScaleToWidth scales the ellipse to the specified width. The height of
// the ellipse is scaled so that the aspect ratio is maintained.
func (_eceb *Ellipse) ScaleToWidth(w float64) {
	_bbc := _eceb._dabc / _eceb._eded
	_eceb._eded = w
	_eceb._dabc = w * _bbc
}

// NewPolygon creates a new polygon.
func (_aabcb *Creator) NewPolygon(points [][]_fc.Point) *Polygon { return _bgcb(points) }
func _fdbcg(_decg *templateProcessor, _cecd *templateNode) (interface{}, error) {
	return _decg.parseBackground(_cecd)
}

// SetAngle sets the rotation angle in degrees.
func (_bgf *Block) SetAngle(angleDeg float64) { _bgf._de = angleDeg }

// GeneratePageBlocks draw graphic svg into block.
func (_edce *GraphicSVG) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_defb := ctx
	_aecg := _edce._dgdb.IsRelative()
	var _cdgf []*Block
	if _aecg {
		_bbae := 1.0
		_dbfa := _edce._eaff.Top
		if _edce._dgaf.Height > ctx.Height-_edce._eaff.Top {
			_cdgf = []*Block{NewBlock(ctx.PageWidth, ctx.PageHeight-ctx.Y)}
			var _ggaa error
			if _, ctx, _ggaa = _cbag().GeneratePageBlocks(ctx); _ggaa != nil {
				return nil, ctx, _ggaa
			}
			_dbfa = 0
		}
		ctx.X += _edce._eaff.Left + _bbae
		ctx.Y += _dbfa
		ctx.Width -= _edce._eaff.Left + _edce._eaff.Right + 2*_bbae
		ctx.Height -= _dbfa
	} else {
		ctx.X = _edce._bgag
		ctx.Y = _edce._gcfge
	}
	_egbg := _bdb.NewContentCreator()
	_egbg.Translate(0, ctx.PageHeight)
	_egbg.Scale(1, -1)
	_egbg.Translate(ctx.X, ctx.Y)
	_abbb := _edce._dgaf.Width / _edce._dgaf.ViewBox.W
	_bdcb := _edce._dgaf.Height / _edce._dgaf.ViewBox.H
	_bgbcd := 0.0
	_bagg := 0.0
	if _aecg {
		_bgbcd = _edce._bgag - (_edce._dgaf.ViewBox.X * _b.Max(_abbb, _bdcb))
		_bagg = _edce._gcfge - (_edce._dgaf.ViewBox.Y * _b.Max(_abbb, _bdcb))
	}
	_edce._dgaf.ToContentCreator(_egbg, _abbb, _bdcb, _bgbcd, _bagg)
	_dbag := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _fgfe := _dbag.addContentsByString(_egbg.String()); _fgfe != nil {
		return nil, ctx, _fgfe
	}
	if _aecg {
		_gbab := _edce.Height() + _edce._eaff.Bottom
		ctx.Y += _gbab
		ctx.Height -= _gbab
	} else {
		ctx = _defb
	}
	_cdgf = append(_cdgf, _dbag)
	return _cdgf, ctx, nil
}

// ScaleToWidth sets the graphic svg scaling factor with the given width.
func (_ddgd *GraphicSVG) ScaleToWidth(w float64) {
	_ececd := _ddgd._dgaf.Height / _ddgd._dgaf.Width
	_ddgd._dgaf.Width = w
	_ddgd._dgaf.Height = w * _ececd
	_ddgd._dgaf.SetScaling(_ececd, _ececd)
}

// Width returns the current page width.
func (_aba *Creator) Width() float64 { return _aba._abf }
func (_cbff *List) markerWidth() float64 {
	var _dccga float64
	for _, _fdbg := range _cbff._bede {
		_aaadb := _egdc(_cbff._dagb)
		_aaadb.SetEnableWrap(false)
		_aaadb.SetTextAlignment(TextAlignmentRight)
		_aaadb.Append(_fdbg._eeed.Text).Style = _fdbg._eeed.Style
		_dagf := _aaadb.getTextWidth() / 1000.0
		if _dccga < _dagf {
			_dccga = _dagf
		}
	}
	return _dccga
}
func (_fdf rgbColor) ToRGB() (float64, float64, float64) { return _fdf._efdc, _fdf._fgbf, _fdf._acbd }

// Columns returns all the columns in the invoice line items table.
func (_bcgge *Invoice) Columns() []*InvoiceCell { return _bcgge._eag }

// SetFillColor sets background color for border.
func (_aeba *border) SetFillColor(col Color) { _aeba._cde = col }

// GetMargins returns the margins of the ellipse: left, right, top, bottom.
func (_eebec *Ellipse) GetMargins() (float64, float64, float64, float64) {
	return _eebec._cadg.Left, _eebec._cadg.Right, _eebec._cadg.Top, _eebec._cadg.Bottom
}

// GeneratePageBlocks draws the chart onto a block.
func (_gedc *Chart) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_fcbf := ctx
	_bgga := _gedc._fbc.IsRelative()
	var _dgd []*Block
	if _bgga {
		_agaa := 1.0
		_gced := _gedc._gcff.Top
		if float64(_gedc._ebbf.Height()) > ctx.Height-_gedc._gcff.Top {
			_dgd = []*Block{NewBlock(ctx.PageWidth, ctx.PageHeight-ctx.Y)}
			var _aebf error
			if _, ctx, _aebf = _cbag().GeneratePageBlocks(ctx); _aebf != nil {
				return nil, ctx, _aebf
			}
			_gced = 0
		}
		ctx.X += _gedc._gcff.Left + _agaa
		ctx.Y += _gced
		ctx.Width -= _gedc._gcff.Left + _gedc._gcff.Right + 2*_agaa
		ctx.Height -= _gced
		_gedc._ebbf.SetWidth(int(ctx.Width))
	} else {
		ctx.X = _gedc._bfe
		ctx.Y = _gedc._gbbc
	}
	_ece := _bdb.NewContentCreator()
	_ece.Translate(0, ctx.PageHeight)
	_ece.Scale(1, -1)
	_ece.Translate(ctx.X, ctx.Y)
	_bfdd := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_gedc._ebbf.Render(_ec.NewRenderer(_ece, _bfdd._ge), nil)
	if _ccf := _bfdd.addContentsByString(_ece.String()); _ccf != nil {
		return nil, ctx, _ccf
	}
	if _bgga {
		_gcfg := _gedc.Height() + _gedc._gcff.Bottom
		ctx.Y += _gcfg
		ctx.Height -= _gcfg
	} else {
		ctx = _fcbf
	}
	_dgd = append(_dgd, _bfdd)
	return _dgd, ctx, nil
}

// EnableWordWrap sets the paragraph word wrap flag.
func (_edcfa *StyledParagraph) EnableWordWrap(val bool) { _edcfa._cbeg = val }
func (_dffdb *List) split(_beaga DrawContext) (_adga, _fdfa *List) {
	var (
		_gdcad        float64
		_ffbd, _dfeec []*listItem
	)
	_adace := _beaga.Width - _dffdb._fed.Horizontal() - _dffdb._egab - _dffdb.markerWidth()
	_eadfe := _dffdb.markerWidth()
	for _gfbg, _fcge := range _dffdb._bede {
		_cdeb := _fcge.ctxHeight(_adace)
		_gdcad += _cdeb
		if _gdcad <= _beaga.Height {
			_ffbd = append(_ffbd, _fcge)
		} else {
			switch _bead := _fcge._gdceb.(type) {
			case *List:
				_baggg := _beaga
				_baggg.Height = _b.Floor(_cdeb - (_gdcad - _beaga.Height))
				_cede, _fbdce := _bead.split(_baggg)
				if _cede != nil {
					_fcbc := _fcf()
					_fcbc._eeed = _fcge._eeed
					_fcbc._gdceb = _cede
					_ffbd = append(_ffbd, _fcbc)
				}
				if _fbdce != nil {
					_gfgea := _bead._ebgba.Style.FontSize
					_dgggb, _dbga := _bead._ebgba.Style.Font.GetRuneMetrics(' ')
					if _dbga {
						_gfgea = _bead._ebgba.Style.FontSize * _dgggb.Wx * _bead._ebgba.Style.horizontalScale() / 1000.0
					}
					_fdda := _dc.Repeat("\u0020", int(_eadfe/_gfgea))
					_ecbd := _fcf()
					_ecbd._eeed = *NewTextChunk(_fdda, _bead._ebgba.Style)
					_ecbd._gdceb = _fbdce
					_dfeec = append(_dfeec, _ecbd)
					_dfeec = append(_dfeec, _dffdb._bede[_gfbg+1:]...)
				}
			default:
				_dfeec = _dffdb._bede[_gfbg:]
			}
			if len(_dfeec) > 0 {
				break
			}
		}
	}
	if len(_ffbd) > 0 {
		_adga = _edbe(_dffdb._dagb)
		*_adga = *_dffdb
		_adga._bede = _ffbd
	}
	if len(_dfeec) > 0 {
		_fdfa = _edbe(_dffdb._dagb)
		*_fdfa = *_dffdb
		_fdfa._bede = _dfeec
	}
	return _adga, _fdfa
}

// Subtotal returns the invoice subtotal description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_edgc *Invoice) Subtotal() (*InvoiceCell, *InvoiceCell) { return _edgc._befc[0], _edgc._befc[1] }
func (_ebee *templateProcessor) parseColor(_febf string) Color {
	if _febf == "" {
		return nil
	}
	_ccbb, _ddee := _ebee._affcb.ColorMap[_febf]
	if _ddee {
		return _ccbb
	}
	if _febf[0] == '#' {
		return ColorRGBFromHex(_febf)
	}
	return nil
}
func (_gcdc *templateProcessor) parseCellAlignmentAttr(_dgfad, _dceb string) CellHorizontalAlignment {
	_ca.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020c\u0065\u006c\u006c\u0020\u0061\u006c\u0069g\u006e\u006d\u0065\u006e\u0074\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028`\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _dgfad, _dceb)
	_aegf := map[string]CellHorizontalAlignment{"\u006c\u0065\u0066\u0074": CellHorizontalAlignmentLeft, "\u0063\u0065\u006e\u0074\u0065\u0072": CellHorizontalAlignmentCenter, "\u0072\u0069\u0067h\u0074": CellHorizontalAlignmentRight}[_dceb]
	return _aegf
}
func (_dbcf *Paragraph) wrapText() error {
	if !_dbcf._bebg || int(_dbcf._dcdb) <= 0 {
		_dbcf._cbcf = []string{_dbcf._age}
		return nil
	}
	_gfbd := NewTextChunk(_dbcf._age, TextStyle{Font: _dbcf._fggb, FontSize: _dbcf._fcbfa})
	_eaaa, _acdf := _gfbd.Wrap(_dbcf._dcdb)
	if _acdf != nil {
		return _acdf
	}
	if _dbcf._bebd > 0 && len(_eaaa) > _dbcf._bebd {
		_eaaa = _eaaa[:_dbcf._bebd]
	}
	_dbcf._cbcf = _eaaa
	return nil
}

// NewChart creates a new creator drawable based on the provided
// unichart chart component.
func NewChart(chart _gg.ChartRenderable) *Chart { return _bga(chart) }
func _cgee(_efdg *_cc.GraphicSVG) (*GraphicSVG, error) {
	return &GraphicSVG{_dgaf: _efdg, _dgdb: PositionRelative, _eaff: Margins{Top: 10, Bottom: 10}}, nil
}

// Angle returns the block rotation angle in degrees.
func (_eeg *Block) Angle() float64 { return _eeg._de }

// Inline returns whether the inline mode of the division is active.
func (_cgde *Division) Inline() bool { return _cgde._aeadf }
func (_ggce *Invoice) generateTotalBlocks(_gffec DrawContext) ([]*Block, DrawContext, error) {
	_bafa := _gdec(4)
	_bafa.SetMargins(0, 0, 10, 10)
	_cfacd := [][2]*InvoiceCell{_ggce._befc}
	_cfacd = append(_cfacd, _ggce._edb...)
	_cfacd = append(_cfacd, _ggce._ada)
	for _, _bbcb := range _cfacd {
		_geed, _gbdd := _bbcb[0], _bbcb[1]
		if _gbdd.Value == "" {
			continue
		}
		_bafa.SkipCells(2)
		_eade := _bafa.NewCell()
		_eade.SetBackgroundColor(_geed.BackgroundColor)
		_eade.SetHorizontalAlignment(_gbdd.Alignment)
		_ggce.setCellBorder(_eade, _geed)
		_cgfec := _egdc(_geed.TextStyle)
		_cgfec.SetMargins(0, 0, 2, 1)
		_cgfec.Append(_geed.Value)
		_eade.SetContent(_cgfec)
		_eade = _bafa.NewCell()
		_eade.SetBackgroundColor(_gbdd.BackgroundColor)
		_eade.SetHorizontalAlignment(_gbdd.Alignment)
		_ggce.setCellBorder(_eade, _geed)
		_cgfec = _egdc(_gbdd.TextStyle)
		_cgfec.SetMargins(0, 0, 2, 1)
		_cgfec.Append(_gbdd.Value)
		_eade.SetContent(_cgfec)
	}
	return _bafa.GeneratePageBlocks(_gffec)
}

// NewPolyline creates a new polyline.
func (_gggg *Creator) NewPolyline(points []_fc.Point) *Polyline { return _ecaea(points) }

// ScaleToHeight scale Image to a specified height h, maintaining the aspect ratio.
func (_acfb *Image) ScaleToHeight(h float64) {
	_faeg := _acfb._cagba / _acfb._efef
	_acfb._efef = h
	_acfb._cagba = h * _faeg
}
func (_dagbf *templateProcessor) parseBackground(_ebeb *templateNode) (interface{}, error) {
	_ggdfc := &Background{}
	for _, _cebge := range _ebeb._gbdee.Attr {
		_ecdg := _cebge.Value
		switch _gegb := _cebge.Name.Local; _gegb {
		case "\u0066\u0069\u006c\u006c\u002d\u0063\u006f\u006c\u006f\u0072":
			_ggdfc.FillColor = _dagbf.parseColorAttr(_gegb, _ecdg)
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072":
			_ggdfc.BorderColor = _dagbf.parseColorAttr(_gegb, _ecdg)
		case "b\u006f\u0072\u0064\u0065\u0072\u002d\u0073\u0069\u007a\u0065":
			_ggdfc.BorderSize = _dagbf.parseFloatAttr(_gegb, _ecdg)
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0072\u0061\u0064\u0069\u0075\u0073":
			_ccbede, _affcf, _edga, _gcagd := _dagbf.parseBorderRadiusAttr(_gegb, _ecdg)
			_ggdfc.SetBorderRadius(_ccbede, _affcf, _gcagd, _edga)
		case "\u0062\u006f\u0072\u0064er\u002d\u0074\u006f\u0070\u002d\u006c\u0065\u0066\u0074\u002d\u0072\u0061\u0064\u0069u\u0073":
			_ggdfc.BorderRadiusTopLeft = _dagbf.parseFloatAttr(_gegb, _ecdg)
		case "\u0062\u006f\u0072de\u0072\u002d\u0074\u006f\u0070\u002d\u0072\u0069\u0067\u0068\u0074\u002d\u0072\u0061\u0064\u0069\u0075\u0073":
			_ggdfc.BorderRadiusTopRight = _dagbf.parseFloatAttr(_gegb, _ecdg)
		case "\u0062o\u0072\u0064\u0065\u0072-\u0062\u006f\u0074\u0074\u006fm\u002dl\u0065f\u0074\u002d\u0072\u0061\u0064\u0069\u0075s":
			_ggdfc.BorderRadiusBottomLeft = _dagbf.parseFloatAttr(_gegb, _ecdg)
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0062\u006f\u0074\u0074o\u006d\u002d\u0072\u0069\u0067\u0068\u0074\u002d\u0072\u0061d\u0069\u0075\u0073":
			_ggdfc.BorderRadiusBottomRight = _dagbf.parseFloatAttr(_gegb, _ecdg)
		default:
			_dagbf.nodeLogDebug(_ebeb, "\u0055\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0062\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e", _gegb)
		}
	}
	return _ggdfc, nil
}

// Width is not used. The list component is designed to fill into the available
// width depending on the context. Returns 0.
func (_agcc *List) Width() float64 { return 0 }

// SetFitMode sets the fit mode of the line.
// NOTE: The fit mode is only applied if relative positioning is used.
func (_bbaf *Line) SetFitMode(fitMode FitMode) { _bbaf._afad = fitMode }

// SetCoords sets the center coordinates of the ellipse.
func (_ebec *Ellipse) SetCoords(xc, yc float64) { _ebec._eebe = xc; _ebec._beb = yc }
func (_cfcb *StyledParagraph) appendChunk(_eggbf *TextChunk) *TextChunk {
	_cfcb._cdffa = append(_cfcb._cdffa, _eggbf)
	_cfcb.wrapText()
	return _eggbf
}

// SetSideBorderStyle sets the cell's side border style.
func (_afdf *TableCell) SetSideBorderStyle(side CellBorderSide, style CellBorderStyle) {
	switch side {
	case CellBorderSideAll:
		_afdf._aaddc = style
		_afdf._cbafe = style
		_afdf._ffdeb = style
		_afdf._dfbfd = style
	case CellBorderSideTop:
		_afdf._aaddc = style
	case CellBorderSideBottom:
		_afdf._cbafe = style
	case CellBorderSideLeft:
		_afdf._ffdeb = style
	case CellBorderSideRight:
		_afdf._dfbfd = style
	}
}

// NewCellProps returns the default properties of an invoice cell.
func (_geafb *Invoice) NewCellProps() InvoiceCellProps {
	_bcebge := ColorRGBFrom8bit(255, 255, 255)
	return InvoiceCellProps{TextStyle: _geafb._bdcd, Alignment: CellHorizontalAlignmentLeft, BackgroundColor: _bcebge, BorderColor: _bcebge, BorderWidth: 1, BorderSides: []CellBorderSide{CellBorderSideAll}}
}

// Opacity returns the opacity of the line.
func (_fcacb *Line) Opacity() float64 { return _fcacb._ggbab }

// SetFitMode sets the fit mode of the rectangle.
// NOTE: The fit mode is only applied if relative positioning is used.
func (_abee *Rectangle) SetFitMode(fitMode FitMode) { _abee._gaef = fitMode }

// SetBorderWidth sets the border width.
func (_dggf *Polygon) SetBorderWidth(borderWidth float64) { _dggf._gbda.BorderWidth = borderWidth }

// FillOpacity returns the fill opacity of the ellipse (0-1).
func (_ddbd *Ellipse) FillOpacity() float64 { return _ddbd._ddbf }

// GeneratePageBlocks generates the page blocks.  Multiple blocks are generated if the contents wrap
// over multiple pages. Implements the Drawable interface.
func (_cefge *Paragraph) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_efaf := ctx
	var _cdffb []*Block
	_bdea := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _cefge._abea.IsRelative() {
		ctx.X += _cefge._fgcbf.Left
		ctx.Y += _cefge._fgcbf.Top
		ctx.Width -= _cefge._fgcbf.Left + _cefge._fgcbf.Right
		ctx.Height -= _cefge._fgcbf.Top
		_cefge.SetWidth(ctx.Width)
		if _cefge.Height() > ctx.Height {
			_cdffb = append(_cdffb, _bdea)
			_bdea = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_cedaf := ctx
			_cedaf.Y = ctx.Margins.Top
			_cedaf.X = ctx.Margins.Left + _cefge._fgcbf.Left
			_cedaf.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
			_cedaf.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _cefge._fgcbf.Left - _cefge._fgcbf.Right
			ctx = _cedaf
		}
	} else {
		if int(_cefge._dcdb) <= 0 {
			_cefge.SetWidth(_cefge.getTextWidth())
		}
		ctx.X = _cefge._bcec
		ctx.Y = _cefge._cbcca
	}
	ctx, _dgbf := _cdgc(_bdea, _cefge, ctx)
	if _dgbf != nil {
		_ca.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dgbf)
		return nil, ctx, _dgbf
	}
	_cdffb = append(_cdffb, _bdea)
	if _cefge._abea.IsRelative() {
		ctx.Y += _cefge._fgcbf.Bottom
		ctx.Height -= _cefge._fgcbf.Bottom
		if !ctx.Inline {
			ctx.X = _efaf.X
			ctx.Width = _efaf.Width
		}
		return _cdffb, ctx, nil
	}
	return _cdffb, _efaf, nil
}

// SetText sets the text content of the Paragraph.
func (_eeca *Paragraph) SetText(text string) { _eeca._age = text }

// SetOptimizer sets the optimizer to optimize PDF before writing.
func (_cab *Creator) SetOptimizer(optimizer _ggc.Optimizer) { _cab._eecf = optimizer }

// AddInternalLink adds a new internal link to the paragraph.
// The text parameter represents the text that is displayed.
// The user is taken to the specified page, at the specified x and y
// coordinates. Position 0, 0 is at the top left of the page.
// The zoom of the destination page is controlled with the zoom
// parameter. Pass in 0 to keep the current zoom value.
func (_cebb *StyledParagraph) AddInternalLink(text string, page int64, x, y, zoom float64) *TextChunk {
	_eedc := NewTextChunk(text, _cebb._aecf)
	_eedc._dfcb = _bfefg(page-1, x, y, zoom)
	return _cebb.appendChunk(_eedc)
}

// Positioning returns the type of positioning the line is set to use.
func (_agdcg *Line) Positioning() Positioning { return _agdcg._fgef }

// SetLineTitleStyle sets the style for the title part of all new lines
// of the table of contents.
func (_gdab *TOC) SetLineTitleStyle(style TextStyle) { _gdab._cbgbd = style }

// CurCol returns the currently active cell's column number.
func (_bbdfe *Table) CurCol() int { _egae := (_bbdfe._ggfb-1)%(_bbdfe._aeaa) + 1; return _egae }

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
func (_geb *Creator) SetPageSize(size PageSize) {
	_geb._gccce = size
	_geb._abf = size[0]
	_geb._ffc = size[1]
	_eafd := 0.1 * _geb._abf
	_geb._gcgd.Left = _eafd
	_geb._gcgd.Right = _eafd
	_geb._gcgd.Top = _eafd
	_geb._gcgd.Bottom = _eafd
}
func _gdec(_aeee int) *Table {
	_gaaf := &Table{_aeaa: _aeee, _ddfaf: 10.0, _abbbf: []float64{}, _begb: []float64{}, _cacca: []*TableCell{}, _degfe: make([]int, _aeee), _gbgc: true}
	_gaaf.resetColumnWidths()
	return _gaaf
}

// ScaleToWidth scales the rectangle to the specified width. The height of
// the rectangle is scaled so that the aspect ratio is maintained.
func (_febg *Rectangle) ScaleToWidth(w float64) {
	_faece := _febg._fefg / _febg._gfad
	_febg._gfad = w
	_febg._fefg = w * _faece
}
func _fcf() *listItem { return &listItem{} }

// SetFillColor sets the fill color of the rectangle.
func (_dbbf *Rectangle) SetFillColor(col Color) { _dbbf._bgca = col }

// SetColorRight sets border color for right.
func (_dgge *border) SetColorRight(col Color) { _dgge._add = col }

// SetLineMargins sets the margins for all new lines of the table of contents.
func (_bfgb *TOC) SetLineMargins(left, right, top, bottom float64) {
	_cbffc := &_bfgb._dfbae
	_cbffc.Left = left
	_cbffc.Right = right
	_cbffc.Top = top
	_cbffc.Bottom = bottom
}

// TOC represents a table of contents component.
// It consists of a paragraph heading and a collection of
// table of contents lines.
// The representation of a table of contents line is as follows:
//
//	[number] [title]      [separator] [page]
//
// e.g.: Chapter1 Introduction ........... 1
type TOC struct {
	_ecagg *StyledParagraph
	_fbeeg []*TOCLine
	_febed TextStyle
	_cbgbd TextStyle
	_effdg TextStyle
	_abgad TextStyle
	_feaba string
	_ebgfa float64
	_dfbae Margins
	_fcffe Positioning
	_cceeb TextStyle
	_dbaf  bool
}

// SetFitMode sets the fit mode of the image.
// NOTE: The fit mode is only applied if relative positioning is used.
func (_bdgc *Image) SetFitMode(fitMode FitMode) { _bdgc._geaf = fitMode }

// ToPdfShadingPattern generates a new model.PdfShadingPatternType2 object.
func (_fadb *LinearShading) ToPdfShadingPattern() *_ggc.PdfShadingPatternType2 {
	_gadca, _cgbe, _bdeeb := _fadb._aecc._dedad.ToRGB()
	_fdcecb := _fadb.shadingModel()
	_fdcecb.PdfShading.Background = _fe.MakeArrayFromFloats([]float64{_gadca, _cgbe, _bdeeb})
	_dgef := _ggc.NewPdfShadingPatternType2()
	_dgef.Shading = _fdcecb
	return _dgef
}

// SetLineNumberStyle sets the style for the numbers part of all new lines
// of the table of contents.
func (_bbgg *TOC) SetLineNumberStyle(style TextStyle) { _bbgg._febed = style }

// SetFillOpacity sets the fill opacity of the rectangle.
func (_eddba *Rectangle) SetFillOpacity(opacity float64) { _eddba._geee = opacity }

// NewColorPoint creates a new color and point object for use in the gradient rendering process.
func NewColorPoint(color Color, point float64) *ColorPoint { return _cafcfa(color, point) }
func _effe(_ecab *Block, _efege *Image, _beg DrawContext) (DrawContext, error) {
	_aeffg := _beg
	_dccd := 1
	_dfga := _fe.PdfObjectName(_df.Sprintf("\u0049\u006d\u0067%\u0064", _dccd))
	for _ecab._ge.HasXObjectByName(_dfga) {
		_dccd++
		_dfga = _fe.PdfObjectName(_df.Sprintf("\u0049\u006d\u0067%\u0064", _dccd))
	}
	_gfga := _ecab._ge.SetXObjectImageByName(_dfga, _efege._edee)
	if _gfga != nil {
		return _beg, _gfga
	}
	_badb := 0
	_dfda := _fe.PdfObjectName(_df.Sprintf("\u0047\u0053\u0025\u0064", _badb))
	for _ecab._ge.HasExtGState(_dfda) {
		_badb++
		_dfda = _fe.PdfObjectName(_df.Sprintf("\u0047\u0053\u0025\u0064", _badb))
	}
	_fcag := _fe.MakeDict()
	_fcag.Set("\u0042\u004d", _fe.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
	if _efege._feef < 1.0 {
		_fcag.Set("\u0043\u0041", _fe.MakeFloat(_efege._feef))
		_fcag.Set("\u0063\u0061", _fe.MakeFloat(_efege._feef))
	}
	_gfga = _ecab._ge.AddExtGState(_dfda, _fe.MakeIndirectObject(_fcag))
	if _gfga != nil {
		return _beg, _gfga
	}
	_agdfd := _efege.Width()
	_acbad := _efege.Height()
	_, _feba := _efege.rotatedSize()
	_cfeb := _beg.X
	_fffe := _beg.PageHeight - _beg.Y - _acbad
	if _efege._ebfa.IsRelative() {
		_fffe -= (_feba - _acbad) / 2
		switch _efege._afdag {
		case HorizontalAlignmentCenter:
			_cfeb += (_beg.Width - _agdfd) / 2
		case HorizontalAlignmentRight:
			_cfeb = _beg.PageWidth - _beg.Margins.Right - _efege._cbea.Right - _agdfd
		}
	}
	_aege := _efege._eadgc
	_babca := _bdb.NewContentCreator()
	_babca.Add_gs(_dfda)
	_babca.Translate(_cfeb, _fffe)
	if _aege != 0 {
		_babca.Translate(_agdfd/2, _acbad/2)
		_babca.RotateDeg(_aege)
		_babca.Translate(-_agdfd/2, -_acbad/2)
	}
	_babca.Scale(_agdfd, _acbad).Add_Do(_dfga)
	_aebe := _babca.Operations()
	_aebe.WrapIfNeeded()
	_ecab.addContents(_aebe)
	if _efege._ebfa.IsRelative() {
		_beg.Y += _feba
		_beg.Height -= _feba
		return _beg, nil
	}
	return _aeffg, nil
}

// SetBorderWidth sets the border width.
func (_bggee *PolyBezierCurve) SetBorderWidth(borderWidth float64) {
	_bggee._ffbb.BorderWidth = borderWidth
}

// TOC returns the table of contents component of the creator.
func (_ecec *Creator) TOC() *TOC { return _ecec._effa }

// Reset removes all the text chunks the paragraph contains.
func (_efcfb *StyledParagraph) Reset() { _efcfb._cdffa = []*TextChunk{} }

// SetHorizontalAlignment sets the horizontal alignment of the image.
func (_dge *Image) SetHorizontalAlignment(alignment HorizontalAlignment) { _dge._afdag = alignment }

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
func (_cbcg *Paragraph) SetColor(col Color) { _cbcg._bged = col }
func (_ffgfad *templateProcessor) parseColorAttr(_bdbcb, _bafc string) Color {
	_ca.Log.Debug("\u0050\u0061rs\u0069\u006e\u0067 \u0063\u006f\u006c\u006fr a\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020(`\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _bdbcb, _bafc)
	_bafc = _dc.TrimSpace(_bafc)
	if _dc.HasPrefix(_bafc, "\u006c\u0069n\u0065\u0061\u0072-\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0028") && _dc.HasSuffix(_bafc, "\u0029") && len(_bafc) > 17 {
		return _ffgfad.parseLinearGradientAttr(_ffgfad.creator, _bafc)
	}
	if _dc.HasPrefix(_bafc, "\u0072\u0061d\u0069\u0061\u006c-\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0028") && _dc.HasSuffix(_bafc, "\u0029") && len(_bafc) > 17 {
		return _ffgfad.parseRadialGradientAttr(_ffgfad.creator, _bafc)
	}
	if _eadeb := _ffgfad.parseColor(_bafc); _eadeb != nil {
		return _eadeb
	}
	return ColorBlack
}
func _cfdeg(_gbcc *Block, _geba *StyledParagraph, _dbfbg [][]*TextChunk, _cced DrawContext) (DrawContext, [][]*TextChunk, error) {
	_fdcgd := 1
	_fcdcf := _fe.PdfObjectName(_df.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _fdcgd))
	for _gbcc._ge.HasFontByName(_fcdcf) {
		_fdcgd++
		_fcdcf = _fe.PdfObjectName(_df.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _fdcgd))
	}
	_adab := _gbcc._ge.SetFontByName(_fcdcf, _geba._egcg.Font.ToPdfObject())
	if _adab != nil {
		return _cced, nil, _adab
	}
	_fdcgd++
	_cbae := _fcdcf
	_dbcac := _geba._egcg.FontSize
	_cadac := _geba._gbga.IsRelative()
	var _gbcf [][]_fe.PdfObjectName
	var _cbgbg [][]*TextChunk
	var _gddcd float64
	for _gdeb, _gaac := range _dbfbg {
		var _ffffb []_fe.PdfObjectName
		var _degfb float64
		if len(_gaac) > 0 {
			_degfb = _gaac[0].Style.FontSize
		}
		for _, _aeef := range _gaac {
			_cfbaeg := _aeef.Style
			if _aeef.Text != "" && _cfbaeg.FontSize > _degfb {
				_degfb = _cfbaeg.FontSize
			}
			if _degfb > _cced.PageHeight {
				return _cced, nil, _fa.New("\u0050\u0061\u0072\u0061\u0067\u0072a\u0070\u0068\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0063\u0061\u006e\u0027\u0074\u0020\u0062\u0065\u0020\u006ca\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0070\u0061\u0067\u0065 \u0068e\u0069\u0067\u0068\u0074")
			}
			_fcdcf = _fe.PdfObjectName(_df.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _fdcgd))
			_adcg := _gbcc._ge.SetFontByName(_fcdcf, _cfbaeg.Font.ToPdfObject())
			if _adcg != nil {
				return _cced, nil, _adcg
			}
			_ffffb = append(_ffffb, _fcdcf)
			_fdcgd++
		}
		_degfb *= _geba._babcf
		if _cadac && _gddcd+_degfb > _cced.Height {
			_cbgbg = _dbfbg[_gdeb:]
			_dbfbg = _dbfbg[:_gdeb]
			break
		}
		_gddcd += _degfb
		_gbcf = append(_gbcf, _ffffb)
	}
	_fbdef, _efdff, _cafb := _geba.getLineMetrics(0)
	_fgagb, _gfbe := _fbdef*_geba._babcf, _efdff*_geba._babcf
	if len(_dbfbg) == 0 {
		return _cced, _cbgbg, nil
	}
	_fcbac := _bdb.NewContentCreator()
	_fcbac.Add_q()
	_geebb := _gfbe
	if _geba._edbcb == TextVerticalAlignmentCenter {
		_geebb = _efdff + (_fbdef+_cafb-_efdff)/2 + (_gfbe-_efdff)/2
	}
	_aabce := _cced.PageHeight - _cced.Y - _geebb
	_fcbac.Translate(_cced.X, _aabce)
	_befba := _aabce
	if _geba._adag != 0 {
		_fcbac.RotateDeg(_geba._adag)
	}
	if _geba._cfec == TextOverflowHidden {
		_fcbac.Add_re(0, -_gddcd+_fgagb+1, _geba._gcadd, _gddcd).Add_W().Add_n()
	}
	_fcbac.Add_BT()
	_ffaa := 0.0
	var _dafce []*_fc.BasicLine
	for _fcaac, _adaac := range _dbfbg {
		_bdge := _cced.X
		var _dbecg float64
		if len(_adaac) > 0 {
			_dbecg = _adaac[0].Style.FontSize
		}
		_fbdef, _, _cafb = _geba.getLineMetrics(_fcaac)
		_gfbe = (_fbdef + _cafb)
		for _, _dacd := range _adaac {
			_fbbe := &_dacd.Style
			if _dacd.Text != "" && _fbbe.FontSize > _dbecg {
				_dbecg = _fbbe.FontSize
			}
			if _gfbe > _dbecg {
				_dbecg = _gfbe
			}
		}
		if _fcaac != 0 {
			_fcbac.Add_TD(0, -_dbecg*_geba._babcf+_ffaa)
			_befba -= _dbecg*_geba._babcf + _ffaa
			_ffaa = 0.0
		}
		_afeg := _fcaac == len(_dbfbg)-1
		var (
			_ccac float64
			_bbgd float64
			_adcd *fontMetrics
			_afcf float64
			_adgg uint
		)
		var _dbgf []float64
		for _, _gcecf := range _adaac {
			_bgdc := &_gcecf.Style
			if _bgdc.FontSize > _bbgd {
				_bbgd = _bgdc.FontSize
				_adcd = _ecfaf(_gcecf.Style.Font, _bgdc.FontSize)
			}
			if _gfbe > _bbgd {
				_bbgd = _gfbe
			}
			_gafa, _fgae := _bgdc.Font.GetRuneMetrics(' ')
			if !_fgae {
				return _cced, nil, _fa.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
			}
			var _bcbaf uint
			var _gcgda float64
			_ddcd := len(_gcecf.Text)
			for _ebgd, _bddfa := range _gcecf.Text {
				if _bddfa == ' ' {
					_bcbaf++
					continue
				}
				if _bddfa == '\u000A' {
					continue
				}
				_deag, _acab := _bgdc.Font.GetRuneMetrics(_bddfa)
				if !_acab {
					_ca.Log.Debug("\u0055\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0075\u006ee\u0020%\u0076\u0020\u0069\u006e\u0020\u0066\u006fn\u0074\u000a", _bddfa)
					return _cced, nil, _fa.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0067\u006c\u0079p\u0068")
				}
				_gcgda += _bgdc.FontSize * _deag.Wx * _bgdc.horizontalScale()
				if _ebgd != _ddcd-1 {
					_gcgda += _bgdc.CharSpacing * 1000.0
				}
			}
			_dbgf = append(_dbgf, _gcgda)
			_ccac += _gcgda
			_afcf += float64(_bcbaf) * _gafa.Wx * _bgdc.FontSize * _bgdc.horizontalScale()
			_adgg += _bcbaf
		}
		_bbgd *= _geba._babcf
		var _fffda []_fe.PdfObject
		_fbgce := _geba._gcadd * 1000.0
		if _geba._eaeg == TextAlignmentJustify {
			if _adgg > 0 && !_afeg {
				_afcf = (_fbgce - _ccac) / float64(_adgg) / _dbcac
			}
		} else if _geba._eaeg == TextAlignmentCenter {
			_cecgb := (_fbgce - _ccac - _afcf) / 2
			_abfgf := _cecgb / _dbcac
			_fffda = append(_fffda, _fe.MakeFloat(-_abfgf))
			_bdge += _cecgb / 1000.0
		} else if _geba._eaeg == TextAlignmentRight {
			_dgdbc := (_fbgce - _ccac - _afcf)
			_fbef := _dgdbc / _dbcac
			_fffda = append(_fffda, _fe.MakeFloat(-_fbef))
			_bdge += _dgdbc / 1000.0
		}
		if len(_fffda) > 0 {
			_fcbac.Add_Tf(_cbae, _dbcac).Add_TL(_dbcac * _geba._babcf).Add_TJ(_fffda...)
		}
		_edcba := 0.0
		for _egacea, _cabee := range _adaac {
			_ebgdg := &_cabee.Style
			_bfcg := _cbae
			_cbgac := _dbcac
			_dgced := _ebgdg.OutlineColor != nil
			_agee := _ebgdg.HorizontalScaling != DefaultHorizontalScaling
			_ffeda := _ebgdg.OutlineSize != 1
			if _ffeda {
				_fcbac.Add_w(_ebgdg.OutlineSize)
			}
			_bbadee := _ebgdg.RenderingMode != TextRenderingModeFill
			if _bbadee {
				_fcbac.Add_Tr(int64(_ebgdg.RenderingMode))
			}
			_gecbb := _ebgdg.CharSpacing != 0
			if _gecbb {
				_fcbac.Add_Tc(_ebgdg.CharSpacing)
			}
			_dgegf := _ebgdg.TextRise != 0
			if _dgegf {
				_fcbac.Add_Ts(_ebgdg.TextRise)
			}
			if _cabee.VerticalAlignment != TextVerticalAlignmentBaseline {
				_ffbbd := _ecfaf(_cabee.Style.Font, _ebgdg.FontSize)
				switch _cabee.VerticalAlignment {
				case TextVerticalAlignmentCenter:
					_edcba = _adcd._fbcdg/2 - _ffbbd._fbcdg/2
				case TextVerticalAlignmentBottom:
					_edcba = _adcd._afebc - _ffbbd._afebc
				case TextVerticalAlignmentTop:
					_edcba = _efdff - _ebgdg.FontSize
				}
				if _edcba != 0.0 {
					_fcbac.Translate(0, _edcba)
				}
			}
			if _geba._eaeg != TextAlignmentJustify || _afeg {
				_begc, _ddaec := _ebgdg.Font.GetRuneMetrics(' ')
				if !_ddaec {
					return _cced, nil, _fa.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
				}
				_bfcg = _gbcf[_fcaac][_egacea]
				_cbgac = _ebgdg.FontSize
				_afcf = _begc.Wx * _ebgdg.horizontalScale()
			}
			_afcgd := _ebgdg.Font.Encoder()
			var _gffbd []byte
			for _, _fabfb := range _cabee.Text {
				if _fabfb == '\u000A' {
					continue
				}
				if _fabfb == ' ' {
					if len(_gffbd) > 0 {
						if _dgced {
							_fcbac.SetStrokingColor(_dbac(_ebgdg.OutlineColor))
						}
						if _agee {
							_fcbac.Add_Tz(_ebgdg.HorizontalScaling)
						}
						_fcbac.SetNonStrokingColor(_dbac(_ebgdg.Color)).Add_Tf(_gbcf[_fcaac][_egacea], _ebgdg.FontSize).Add_TJ([]_fe.PdfObject{_fe.MakeStringFromBytes(_gffbd)}...)
						_gffbd = nil
					}
					if _agee {
						_fcbac.Add_Tz(DefaultHorizontalScaling)
					}
					_fcbac.Add_Tf(_bfcg, _cbgac).Add_TJ([]_fe.PdfObject{_fe.MakeFloat(-_afcf)}...)
					_dbgf[_egacea] += _afcf * _cbgac
				} else {
					if _, _ecag := _afcgd.RuneToCharcode(_fabfb); !_ecag {
						_adab = UnsupportedRuneError{Message: _df.Sprintf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0075\u006e\u0065 \u0069\u006e\u0020\u0074\u0065\u0078\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0023\u0078\u0020\u0028\u0025\u0063\u0029", _fabfb, _fabfb), Rune: _fabfb}
						_cced._bcbc = append(_cced._bcbc, _adab)
						_ca.Log.Debug(_adab.Error())
						if _cced._dgce <= 0 {
							continue
						}
						_fabfb = _cced._dgce
					}
					_gffbd = append(_gffbd, _afcgd.Encode(string(_fabfb))...)
				}
			}
			if len(_gffbd) > 0 {
				if _dgced {
					_fcbac.SetStrokingColor(_dbac(_ebgdg.OutlineColor))
				}
				if _agee {
					_fcbac.Add_Tz(_ebgdg.HorizontalScaling)
				}
				_fcbac.SetNonStrokingColor(_dbac(_ebgdg.Color)).Add_Tf(_gbcf[_fcaac][_egacea], _ebgdg.FontSize).Add_TJ([]_fe.PdfObject{_fe.MakeStringFromBytes(_gffbd)}...)
			}
			_fgdg := _dbgf[_egacea] / 1000.0
			if _ebgdg.Underline {
				_gffef := _ebgdg.UnderlineStyle.Color
				if _gffef == nil {
					_gffef = _cabee.Style.Color
				}
				_fgbcc, _edaf, _acfdg := _gffef.ToRGB()
				_fgbcgf := _bdge - _cced.X
				_ccgdd := _befba - _aabce + _ebgdg.TextRise - _ebgdg.UnderlineStyle.Offset
				_dafce = append(_dafce, &_fc.BasicLine{X1: _fgbcgf, Y1: _ccgdd, X2: _fgbcgf + _fgdg, Y2: _ccgdd, LineWidth: _cabee.Style.UnderlineStyle.Thickness, LineColor: _ggc.NewPdfColorDeviceRGB(_fgbcc, _edaf, _acfdg)})
			}
			if _cabee._dfcb != nil {
				var _afbba *_fe.PdfObjectArray
				if !_cabee._abacb {
					switch _bdade := _cabee._dfcb.GetContext().(type) {
					case *_ggc.PdfAnnotationLink:
						_afbba = _fe.MakeArray()
						_bdade.Rect = _afbba
						_abfdb, _cgdg := _bdade.Dest.(*_fe.PdfObjectArray)
						if _cgdg && _abfdb.Len() == 5 {
							_dccf, _agcee := _abfdb.Get(1).(*_fe.PdfObjectName)
							if _agcee && _dccf.String() == "\u0058\u0059\u005a" {
								_bcad, _afec := _fe.GetNumberAsFloat(_abfdb.Get(3))
								if _afec == nil {
									_abfdb.Set(3, _fe.MakeFloat(_cced.PageHeight-_bcad))
								}
							}
						}
					}
					_cabee._abacb = true
				}
				if _afbba != nil {
					_dbbe := _fc.NewPoint(_bdge-_cced.X, _befba+_ebgdg.TextRise-_aabce).Rotate(_geba._adag)
					_dbbe.X += _cced.X
					_dbbe.Y += _aabce
					_ffcf, _egda, _dcdfb, _dfed := _gcdfe(_fgdg, _bbgd, _geba._adag)
					_dbbe.X += _ffcf
					_dbbe.Y += _egda
					_afbba.Clear()
					_afbba.Append(_fe.MakeFloat(_dbbe.X))
					_afbba.Append(_fe.MakeFloat(_dbbe.Y))
					_afbba.Append(_fe.MakeFloat(_dbbe.X + _dcdfb))
					_afbba.Append(_fe.MakeFloat(_dbbe.Y + _dfed))
				}
				_gbcc.AddAnnotation(_cabee._dfcb)
			}
			_bdge += _fgdg
			if _ffeda {
				_fcbac.Add_w(1.0)
			}
			if _dgced {
				_fcbac.Add_RG(0.0, 0.0, 0.0)
			}
			if _bbadee {
				_fcbac.Add_Tr(int64(TextRenderingModeFill))
			}
			if _gecbb {
				_fcbac.Add_Tc(0)
			}
			if _dgegf {
				_fcbac.Add_Ts(0)
			}
			if _agee {
				_fcbac.Add_Tz(DefaultHorizontalScaling)
			}
			if _edcba != 0.0 {
				_fcbac.Translate(0, -_edcba)
				_edcba = 0.0
			}
		}
	}
	_fcbac.Add_ET()
	for _, _decfb := range _dafce {
		_fcbac.SetStrokingColor(_decfb.LineColor).Add_w(_decfb.LineWidth).Add_m(_decfb.X1, _decfb.Y1).Add_l(_decfb.X2, _decfb.Y2).Add_s()
	}
	_fcbac.Add_Q()
	_abfdd := _fcbac.Operations()
	_abfdd.WrapIfNeeded()
	_gbcc.addContents(_abfdd)
	if _cadac {
		_bdecc := _gddcd
		_cced.Y += _bdecc
		_cced.Height -= _bdecc
		if _cced.Inline {
			_cced.X += _geba.Width() + _geba._fbgbc.Right
		}
	}
	return _cced, _cbgbg, nil
}

// StyledParagraph represents text drawn with a specified font and can wrap across lines and pages.
// By default occupies the available width in the drawing context.
type StyledParagraph struct {
	_cdffa []*TextChunk
	_egcg  TextStyle
	_aecf  TextStyle
	_eaeg  TextAlignment
	_edbcb TextVerticalAlignment
	_babcf float64
	_gfbb  bool
	_gcadd float64
	_cbeg  bool
	_facb  bool
	_cfec  TextOverflow
	_adag  float64
	_fbgbc Margins
	_gbga  Positioning
	_gcfda float64
	_cfge  float64
	_bfagd float64
	_bdbc  float64
	_aabba [][]*TextChunk
	_dbdfe func(_afaf *StyledParagraph, _bbade DrawContext)
}

func _cdgc(_gfbag *Block, _gfacb *Paragraph, _gebfa DrawContext) (DrawContext, error) {
	_dadfc := 1
	_dfdca := _fe.PdfObjectName("\u0046\u006f\u006e\u0074" + _a.Itoa(_dadfc))
	for _gfbag._ge.HasFontByName(_dfdca) {
		_dadfc++
		_dfdca = _fe.PdfObjectName("\u0046\u006f\u006e\u0074" + _a.Itoa(_dadfc))
	}
	_bdec := _gfbag._ge.SetFontByName(_dfdca, _gfacb._fggb.ToPdfObject())
	if _bdec != nil {
		return _gebfa, _bdec
	}
	_gfacb.wrapText()
	_egggd := _bdb.NewContentCreator()
	_egggd.Add_q()
	_cbabe := _gebfa.PageHeight - _gebfa.Y - _gfacb._fcbfa*_gfacb._dacae
	_egggd.Translate(_gebfa.X, _cbabe)
	if _gfacb._dgeg != 0 {
		_egggd.RotateDeg(_gfacb._dgeg)
	}
	_decb := _dbac(_gfacb._bged)
	_bdec = _aede(_gfbag, _decb, _gfacb._bged, func() Rectangle {
		return Rectangle{_cgedd: _gebfa.X, _eeff: _cbabe, _gfad: _gfacb.getMaxLineWidth() / 1000.0, _fefg: _gfacb.Height()}
	})
	if _bdec != nil {
		return _gebfa, _bdec
	}
	_egggd.Add_BT().SetNonStrokingColor(_decb).Add_Tf(_dfdca, _gfacb._fcbfa).Add_TL(_gfacb._fcbfa * _gfacb._dacae)
	for _dede, _beed := range _gfacb._cbcf {
		if _dede != 0 {
			_egggd.Add_Tstar()
		}
		_bbccd := []rune(_beed)
		_gfdfg := 0.0
		_defae := 0
		for _abbbd, _abbbde := range _bbccd {
			if _abbbde == ' ' {
				_defae++
				continue
			}
			if _abbbde == '\u000A' {
				continue
			}
			_egbgb, _gbaae := _gfacb._fggb.GetRuneMetrics(_abbbde)
			if !_gbaae {
				_ca.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0075\u006e\u0065\u0020\u0069=\u0025\u0064\u0020\u0072\u0075\u006e\u0065=\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0020\u0069n\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0073\u0020\u0025\u0073", _abbbd, _abbbde, _abbbde, _gfacb._fggb.BaseFont(), _gfacb._fggb.Subtype())
				return _gebfa, _fa.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0067\u006c\u0079p\u0068")
			}
			_gfdfg += _gfacb._fcbfa * _egbgb.Wx
		}
		var _afbf []_fe.PdfObject
		_efae, _dbd := _gfacb._fggb.GetRuneMetrics(' ')
		if !_dbd {
			return _gebfa, _fa.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
		}
		_eabf := _efae.Wx
		switch _gfacb._abd {
		case TextAlignmentJustify:
			if _defae > 0 && _dede < len(_gfacb._cbcf)-1 {
				_eabf = (_gfacb._dcdb*1000.0 - _gfdfg) / float64(_defae) / _gfacb._fcbfa
			}
		case TextAlignmentCenter:
			_egace := _gfdfg + float64(_defae)*_eabf*_gfacb._fcbfa
			_efdf := (_gfacb._dcdb*1000.0 - _egace) / 2 / _gfacb._fcbfa
			_afbf = append(_afbf, _fe.MakeFloat(-_efdf))
		case TextAlignmentRight:
			_dade := _gfdfg + float64(_defae)*_eabf*_gfacb._fcbfa
			_dbff := (_gfacb._dcdb*1000.0 - _dade) / _gfacb._fcbfa
			_afbf = append(_afbf, _fe.MakeFloat(-_dbff))
		}
		_egag := _gfacb._fggb.Encoder()
		var _bdgd []byte
		for _, _gadg := range _bbccd {
			if _gadg == '\u000A' {
				continue
			}
			if _gadg == ' ' {
				if len(_bdgd) > 0 {
					_afbf = append(_afbf, _fe.MakeStringFromBytes(_bdgd))
					_bdgd = nil
				}
				_afbf = append(_afbf, _fe.MakeFloat(-_eabf))
			} else {
				if _, _baca := _egag.RuneToCharcode(_gadg); !_baca {
					_bdec = UnsupportedRuneError{Message: _df.Sprintf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0075\u006e\u0065 \u0069\u006e\u0020\u0074\u0065\u0078\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0023\u0078\u0020\u0028\u0025\u0063\u0029", _gadg, _gadg), Rune: _gadg}
					_gebfa._bcbc = append(_gebfa._bcbc, _bdec)
					_ca.Log.Debug(_bdec.Error())
					if _gebfa._dgce <= 0 {
						continue
					}
					_gadg = _gebfa._dgce
				}
				_bdgd = append(_bdgd, _egag.Encode(string(_gadg))...)
			}
		}
		if len(_bdgd) > 0 {
			_afbf = append(_afbf, _fe.MakeStringFromBytes(_bdgd))
		}
		_egggd.Add_TJ(_afbf...)
	}
	_egggd.Add_ET()
	_egggd.Add_Q()
	_fdcgf := _egggd.Operations()
	_fdcgf.WrapIfNeeded()
	_gfbag.addContents(_fdcgf)
	if _gfacb._abea.IsRelative() {
		_aabg := _gfacb.Height()
		_gebfa.Y += _aabg
		_gebfa.Height -= _aabg
		if _gebfa.Inline {
			_gebfa.X += _gfacb.Width() + _gfacb._fgcbf.Right
		}
	}
	return _gebfa, nil
}

// Scale scales Image by a constant factor, both width and height.
func (_dbb *Image) Scale(xFactor, yFactor float64) {
	_dbb._cagba = xFactor * _dbb._cagba
	_dbb._efef = yFactor * _dbb._efef
}

// SetLineWidth sets the line width.
func (_aaeb *Polyline) SetLineWidth(lineWidth float64) { _aaeb._edbcg.LineWidth = lineWidth }

// BorderOpacity returns the border opacity of the ellipse (0-1).
func (_bgaec *Ellipse) BorderOpacity() float64 { return _bgaec._dadc }

const (
	FitModeNone FitMode = iota
	FitModeFillWidth
)

// SetPos sets the position of the graphic svg to the specified coordinates.
// This method sets the graphic svg to use absolute positioning.
func (_baabe *GraphicSVG) SetPos(x, y float64) {
	_baabe._dgdb = PositionAbsolute
	_baabe._bgag = x
	_baabe._gcfge = y
}

// Color interface represents colors in the PDF creator.
type Color interface {
	ToRGB() (float64, float64, float64)
}

// SetRowHeight sets the height for a specified row.
func (_fcde *Table) SetRowHeight(row int, h float64) error {
	if row < 1 || row > len(_fcde._begb) {
		return _fa.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_fcde._begb[row-1] = h
	return nil
}
func (_fagbe *templateProcessor) parseTable(_bdcff *templateNode) (interface{}, error) {
	var _ffceb int64
	for _, _daffa := range _bdcff._gbdee.Attr {
		_gbed := _daffa.Value
		switch _fbbf := _daffa.Name.Local; _fbbf {
		case "\u0063o\u006c\u0075\u006d\u006e\u0073":
			_ffceb = _fagbe.parseInt64Attr(_fbbf, _gbed)
		}
	}
	if _ffceb <= 0 {
		_fagbe.nodeLogDebug(_bdcff, "\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006eu\u006d\u0062e\u0072\u0020\u006f\u0066\u0020\u0074\u0061\u0062\u006ce\u0020\u0063\u006f\u006cu\u006d\u006e\u0073\u003a\u0020\u0025\u0064\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0031\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020m\u0061\u0079\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e", _ffceb)
		_ffceb = 1
	}
	_dcbgg := _fagbe.creator.NewTable(int(_ffceb))
	for _, _bagf := range _bdcff._gbdee.Attr {
		_ddgede := _bagf.Value
		switch _cgfbc := _bagf.Name.Local; _cgfbc {
		case "\u0063\u006f\u006c\u0075\u006d\u006e\u002d\u0077\u0069\u0064\u0074\u0068\u0073":
			_dcbgg.SetColumnWidths(_fagbe.parseFloatArray(_cgfbc, _ddgede)...)
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_fbfda := _fagbe.parseMarginAttr(_cgfbc, _ddgede)
			_dcbgg.SetMargins(_fbfda.Left, _fbfda.Right, _fbfda.Top, _fbfda.Bottom)
		case "\u0078":
			_dcbgg.SetPos(_fagbe.parseFloatAttr(_cgfbc, _ddgede), _dcbgg._degda)
		case "\u0079":
			_dcbgg.SetPos(_dcbgg._fedd, _fagbe.parseFloatAttr(_cgfbc, _ddgede))
		case "\u0068\u0065a\u0064\u0065\u0072-\u0073\u0074\u0061\u0072\u0074\u002d\u0072\u006f\u0077":
			_dcbgg._aggfe = int(_fagbe.parseInt64Attr(_cgfbc, _ddgede))
		case "\u0068\u0065\u0061\u0064\u0065\u0072\u002d\u0065\u006ed\u002d\u0072\u006f\u0077":
			_dcbgg._fecf = int(_fagbe.parseInt64Attr(_cgfbc, _ddgede))
		case "\u0065n\u0061b\u006c\u0065\u002d\u0072\u006f\u0077\u002d\u0077\u0072\u0061\u0070":
			_dcbgg.EnableRowWrap(_fagbe.parseBoolAttr(_cgfbc, _ddgede))
		case "\u0065\u006ea\u0062\u006c\u0065-\u0070\u0061\u0067\u0065\u002d\u0077\u0072\u0061\u0070":
			_dcbgg.EnablePageWrap(_fagbe.parseBoolAttr(_cgfbc, _ddgede))
		case "\u0063o\u006c\u0075\u006d\u006e\u0073":
		default:
			_fagbe.nodeLogDebug(_bdcff, "\u0055n\u0073\u0075p\u0070\u006f\u0072\u0074e\u0064\u0020\u0074a\u0062\u006c\u0065\u0020\u0061\u0074\u0074\u0072\u0069bu\u0074\u0065\u003a \u0060\u0025s\u0060\u002e\u0020\u0053\u006b\u0069p\u0070\u0069n\u0067\u002e", _cgfbc)
		}
	}
	if _dcbgg._aggfe != 0 && _dcbgg._fecf != 0 {
		_abdg := _dcbgg.SetHeaderRows(_dcbgg._aggfe, _dcbgg._fecf)
		if _abdg != nil {
			_fagbe.nodeLogDebug(_bdcff, "\u0043\u006ful\u0064\u0020\u006eo\u0074\u0020\u0073\u0065t t\u0061bl\u0065\u0020\u0068\u0065\u0061\u0064\u0065r \u0072\u006f\u0077\u0073\u003a\u0020\u0025v\u002e", _abdg)
		}
	} else {
		_dcbgg._aggfe = 0
		_dcbgg._fecf = 0
	}
	return _dcbgg, nil
}
func (_bedebd *templateProcessor) parseTextAlignmentAttr(_bfdbc, _agbce string) TextAlignment {
	_ca.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020t\u0065\u0078\u0074\u0020\u0061\u006c\u0069g\u006e\u006d\u0065\u006e\u0074\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028`\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _bfdbc, _agbce)
	_beba := map[string]TextAlignment{"\u006c\u0065\u0066\u0074": TextAlignmentLeft, "\u0072\u0069\u0067h\u0074": TextAlignmentRight, "\u0063\u0065\u006e\u0074\u0065\u0072": TextAlignmentCenter, "\u006au\u0073\u0074\u0069\u0066\u0079": TextAlignmentJustify}[_agbce]
	return _beba
}
func (_fbcb *Image) rotatedSize() (float64, float64) {
	_dfgb := _fbcb._cagba
	_bdac := _fbcb._efef
	_bcag := _fbcb._eadgc
	if _bcag == 0 {
		return _dfgb, _bdac
	}
	_aadfe := _fc.Path{Points: []_fc.Point{_fc.NewPoint(0, 0).Rotate(_bcag), _fc.NewPoint(_dfgb, 0).Rotate(_bcag), _fc.NewPoint(0, _bdac).Rotate(_bcag), _fc.NewPoint(_dfgb, _bdac).Rotate(_bcag)}}.GetBoundingBox()
	return _aadfe.Width, _aadfe.Height
}

// SetExtends specifies whether ot extend the shading beyond the starting and ending points.
//
// Text extends is set to `[]bool{false, false}` by default.
func (_ecbde *RadialShading) SetExtends(start bool, end bool) { _ecbde._fcacc.SetExtends(start, end) }

// SetWidth sets the width of the rectangle.
func (_feae *Rectangle) SetWidth(width float64) { _feae._gfad = width }

var _dac = _fag.MustCompile("\u005c\u0064\u002b")

// SellerAddress returns the seller address used in the invoice template.
func (_eggge *Invoice) SellerAddress() *InvoiceAddress { return _eggge._cfbbb }

// NewGraphicSVGFromString creates a graphic SVG from a SVG string.
func NewGraphicSVGFromString(svgStr string) (*GraphicSVG, error) { return _bdcfa(svgStr) }

// NewInvoice returns an instance of an empty invoice.
func (_bfdg *Creator) NewInvoice() *Invoice {
	_gdfcf := _bfdg.NewTextStyle()
	_gdfcf.Font = _bfdg._eacc
	return _gcag(_bfdg.NewTextStyle(), _gdfcf)
}
func _ffa(_bfd string, _cdga _fe.PdfObject, _dgg *_ggc.PdfPageResources) _fe.PdfObjectName {
	_befe := _dc.TrimRightFunc(_dc.TrimSpace(_bfd), func(_dcbb rune) bool { return _cd.IsNumber(_dcbb) })
	if _befe == "" {
		_befe = "\u0046\u006f\u006e\u0074"
	}
	_eafc := 0
	_gee := _fe.PdfObjectName(_bfd)
	for {
		_efd, _cbf := _dgg.GetFontByName(_gee)
		if !_cbf || _efd == _cdga {
			break
		}
		_eafc++
		_gee = _fe.PdfObjectName(_df.Sprintf("\u0025\u0073\u0025\u0064", _befe, _eafc))
	}
	return _gee
}

// Polyline represents a slice of points that are connected as straight lines.
// Implements the Drawable interface and can be rendered using the Creator.
type Polyline struct {
	_edbcg *_fc.Polyline
	_afcb  float64
}

func (_ggg cmykColor) ToRGB() (float64, float64, float64) {
	_dfee := _ggg._ccb
	return 1 - (_ggg._badc*(1-_dfee) + _dfee), 1 - (_ggg._dea*(1-_dfee) + _dfee), 1 - (_ggg._gaa*(1-_dfee) + _dfee)
}

// SetMargins sets the Paragraph's margins.
func (_bdbbb *StyledParagraph) SetMargins(left, right, top, bottom float64) {
	_bdbbb._fbgbc.Left = left
	_bdbbb._fbgbc.Right = right
	_bdbbb._fbgbc.Top = top
	_bdbbb._fbgbc.Bottom = bottom
}
func (_fdcdd *templateProcessor) parseInt64Attr(_dabe, _cecab string) int64 {
	_ca.Log.Debug("\u0050\u0061rs\u0069\u006e\u0067 \u0069\u006e\u0074\u00364 a\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020(`\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _dabe, _cecab)
	_bebdd, _ := _a.ParseInt(_cecab, 10, 64)
	return _bebdd
}

// SetEnableWrap sets the line wrapping enabled flag.
func (_gedb *Paragraph) SetEnableWrap(enableWrap bool) { _gedb._bebg = enableWrap; _gedb._ggad = false }

// SetPositioning sets the positioning of the ellipse (absolute or relative).
func (_daee *Ellipse) SetPositioning(position Positioning) { _daee._acgd = position }

// SetShowNumbering sets a flag to indicate whether or not to show chapter numbers as part of title.
func (_eedf *Chapter) SetShowNumbering(show bool) {
	_eedf._fgg = show
	_eedf._ddd.SetText(_eedf.headingText())
}

// BorderOpacity returns the border opacity of the rectangle (0-1).
func (_bbee *Rectangle) BorderOpacity() float64 { return _bbee._ffgd }

// GeneratePageBlocks generate the Page blocks. Draws the Image on a block, implementing the Drawable interface.
func (_cgbb *Image) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	if _cgbb._edee == nil {
		if _afdd := _cgbb.makeXObject(); _afdd != nil {
			return nil, ctx, _afdd
		}
	}
	var _gdfdd []*Block
	_abce := ctx
	_fgcb := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _cgbb._ebfa.IsRelative() {
		_cgbb.applyFitMode(ctx.Width)
		ctx.X += _cgbb._cbea.Left
		ctx.Y += _cgbb._cbea.Top
		ctx.Width -= _cgbb._cbea.Left + _cgbb._cbea.Right
		ctx.Height -= _cgbb._cbea.Top + _cgbb._cbea.Bottom
		if _cgbb._efef > ctx.Height {
			_gdfdd = append(_gdfdd, _fgcb)
			_fgcb = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_bbaef := ctx
			_bbaef.Y = ctx.Margins.Top + _cgbb._cbea.Top
			_bbaef.X = ctx.Margins.Left + _cgbb._cbea.Left
			_bbaef.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _cgbb._cbea.Top - _cgbb._cbea.Bottom
			_bbaef.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _cgbb._cbea.Left - _cgbb._cbea.Right
			ctx = _bbaef
		}
	} else {
		ctx.X = _cgbb._fce
		ctx.Y = _cgbb._gdd
	}
	ctx, _ebffg := _effe(_fgcb, _cgbb, ctx)
	if _ebffg != nil {
		return nil, ctx, _ebffg
	}
	_gdfdd = append(_gdfdd, _fgcb)
	if _cgbb._ebfa.IsAbsolute() {
		ctx = _abce
	} else {
		ctx.X = _abce.X
		ctx.Width = _abce.Width
		ctx.Y += _cgbb._cbea.Bottom
	}
	return _gdfdd, ctx, nil
}

// TextStyle is a collection of properties that can be assigned to a text chunk.
type TextStyle struct {

	// Color represents the color of the text.
	Color Color

	// OutlineColor represents the color of the text outline.
	OutlineColor Color

	// Font represents the font the text will use.
	Font *_ggc.PdfFont

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

// SetColor sets the line color.
func (_gde *Curve) SetColor(col Color) { _gde._ccad = col }

// HorizontalAlignment represents the horizontal alignment of components
// within a page.
type HorizontalAlignment int

// Insert adds a new text chunk at the specified position in the paragraph.
func (_daage *StyledParagraph) Insert(index uint, text string) *TextChunk {
	_aafc := uint(len(_daage._cdffa))
	if index > _aafc {
		index = _aafc
	}
	_efccb := NewTextChunk(text, _daage._egcg)
	_daage._cdffa = append(_daage._cdffa[:index], append([]*TextChunk{_efccb}, _daage._cdffa[index:]...)...)
	_daage.wrapText()
	return _efccb
}

const (
	TextAlignmentLeft TextAlignment = iota
	TextAlignmentRight
	TextAlignmentCenter
	TextAlignmentJustify
)
const (
	CellBorderStyleNone CellBorderStyle = iota
	CellBorderStyleSingle
	CellBorderStyleDouble
)

// SetBackgroundColor sets the cell's background color.
func (_fdced *TableCell) SetBackgroundColor(col Color) { _fdced._afbg = col }

// Width is not used as the division component is designed to fill all the
// available space, depending on the context. Returns 0.
func (_aagd *Division) Width() float64 { return 0 }

// GetCoords returns the upper left corner coordinates of the rectangle (`x`, `y`).
func (_ddca *Rectangle) GetCoords() (float64, float64) { return _ddca._cgedd, _ddca._eeff }

type rgbColor struct{ _efdc, _fgbf, _acbd float64 }

// NewLine creates a new line between (x1, y1) to (x2, y2),
// using default attributes.
// NOTE: In relative positioning mode, `x1` and `y1` are calculated using the
// current context and `x2`, `y2` are used only to calculate the position of
// the second point in relation to the first one (used just as a measurement
// of size). Furthermore, when the fit mode is set to fill the context width,
// `x2` is set to the right edge coordinate of the context.
func (_ccag *Creator) NewLine(x1, y1, x2, y2 float64) *Line { return _dfgcg(x1, y1, x2, y2) }
func (_ecfg *Invoice) generateNoteBlocks(_cbaf DrawContext) ([]*Block, DrawContext, error) {
	_aaef := _ffcc()
	_fddb := append([][2]string{_ecfg._ecf, _ecfg._aacd}, _ecfg._agdc...)
	for _, _bff := range _fddb {
		if _bff[1] != "" {
			_ebcg := _ecfg.drawSection(_bff[0], _bff[1])
			for _, _dcd := range _ebcg {
				_aaef.Add(_dcd)
			}
			_dfde := _egdc(_ecfg._bdcd)
			_dfde.SetMargins(0, 0, 10, 0)
			_aaef.Add(_dfde)
		}
	}
	return _aaef.GeneratePageBlocks(_cbaf)
}

// SetMargins sets the Chapter margins: left, right, top, bottom.
// Typically not needed as the creator's page margins are used.
func (_baef *Chapter) SetMargins(left, right, top, bottom float64) {
	_baef._ecga.Left = left
	_baef._ecga.Right = right
	_baef._ecga.Top = top
	_baef._ecga.Bottom = bottom
}

// SetColPosition sets cell column position.
func (_bdbgc *TableCell) SetColPosition(col int) { _bdbgc._eafbd = col }
func (_fdg *Chapter) headingNumber() string {
	var _aebaa string
	if _fdg._fgg {
		if _fdg._cfde != 0 {
			_aebaa = _a.Itoa(_fdg._cfde) + "\u002e"
		}
		if _fdg._beae != nil {
			_aabc := _fdg._beae.headingNumber()
			if _aabc != "" {
				_aebaa = _aabc + _aebaa
			}
		}
	}
	return _aebaa
}

// Block contains a portion of PDF Page contents. It has a width and a position and can
// be placed anywhere on a Page.  It can even contain a whole Page, and is used in the creator
// where each Drawable object can output one or more blocks, each representing content for separate pages
// (typically needed when Page breaks occur).
type Block struct {
	_cad     *_bdb.ContentStreamOperations
	_ge      *_ggc.PdfPageResources
	_ba      Positioning
	_gb, _gf float64
	_ecd     float64
	_gfe     float64
	_de      float64
	_bc      Margins
	_ga      []*_ggc.PdfAnnotation
}

func (_gebe *FilledCurve) draw(_cgeg *Block, _dafe string) ([]byte, *_ggc.PdfRectangle, error) {
	_debaf := _fc.NewCubicBezierPath()
	for _, _deef := range _gebe._babe {
		_debaf = _debaf.AppendCurve(_deef)
	}
	creator := _bdb.NewContentCreator()
	creator.Add_q()
	if _gebe.FillEnabled && _gebe._ddddf != nil {
		_cdbe := _dbac(_gebe._ddddf)
		_fad := _aede(_cgeg, _cdbe, _gebe._ddddf, func() Rectangle {
			_cfgg := _fc.NewCubicBezierPath()
			for _, _caffec := range _gebe._babe {
				_cfgg = _cfgg.AppendCurve(_caffec)
			}
			_bcfa := _cfgg.GetBoundingBox()
			if _gebe.BorderEnabled {
				_bcfa.Height += _gebe.BorderWidth
				_bcfa.Width += _gebe.BorderWidth
				_bcfa.X -= _gebe.BorderWidth / 2
				_bcfa.Y -= _gebe.BorderWidth / 2
			}
			return Rectangle{_cgedd: _bcfa.X, _eeff: _bcfa.Y, _gfad: _bcfa.Width, _fefg: _bcfa.Height}
		})
		if _fad != nil {
			return nil, nil, _fad
		}
		creator.SetNonStrokingColor(_cdbe)
	}
	if _gebe.BorderEnabled {
		if _gebe._eadf != nil {
			creator.SetStrokingColor(_dbac(_gebe._eadf))
		}
		creator.Add_w(_gebe.BorderWidth)
	}
	if len(_dafe) > 1 {
		creator.Add_gs(_fe.PdfObjectName(_dafe))
	}
	_fc.DrawBezierPathWithCreator(_debaf, creator)
	creator.Add_h()
	if _gebe.FillEnabled && _gebe.BorderEnabled {
		creator.Add_B()
	} else if _gebe.FillEnabled {
		creator.Add_f()
	} else if _gebe.BorderEnabled {
		creator.Add_S()
	}
	creator.Add_Q()
	_ceb := _debaf.GetBoundingBox()
	if _gebe.BorderEnabled {
		_ceb.Height += _gebe.BorderWidth
		_ceb.Width += _gebe.BorderWidth
		_ceb.X -= _gebe.BorderWidth / 2
		_ceb.Y -= _gebe.BorderWidth / 2
	}
	_aade := &_ggc.PdfRectangle{}
	_aade.Llx = _ceb.X
	_aade.Lly = _ceb.Y
	_aade.Urx = _ceb.X + _ceb.Width
	_aade.Ury = _ceb.Y + _ceb.Height
	return creator.Bytes(), _aade, nil
}

// GetMargins returns the Chapter's margin: left, right, top, bottom.
func (_efeg *Chapter) GetMargins() (float64, float64, float64, float64) {
	return _efeg._ecga.Left, _efeg._ecga.Right, _efeg._ecga.Top, _efeg._ecga.Bottom
}
func _bgcb(_baeb [][]_fc.Point) *Polygon {
	return &Polygon{_gbda: &_fc.Polygon{Points: _baeb}, _befea: 1.0, _dffbb: 1.0}
}

var (
	ErrContentNotFit = _fa.New("\u0043\u0061\u006e\u006e\u006ft\u0020\u0066\u0069\u0074\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020i\u006e\u0074\u006f\u0020\u0061\u006e\u0020\u0065\u0078\u0069\u0073\u0074\u0069\u006e\u0067\u0020\u0073\u0070\u0061\u0063\u0065")
)

// SetSideBorderColor sets the cell's side border color.
func (_dgbebd *TableCell) SetSideBorderColor(side CellBorderSide, col Color) {
	switch side {
	case CellBorderSideAll:
		_dgbebd._afga = col
		_dgbebd._dagd = col
		_dgbebd._egde = col
		_dgbebd._aeec = col
	case CellBorderSideTop:
		_dgbebd._afga = col
	case CellBorderSideBottom:
		_dgbebd._dagd = col
	case CellBorderSideLeft:
		_dgbebd._egde = col
	case CellBorderSideRight:
		_dgbebd._aeec = col
	}
}

// Width is not used. Not used as a Table element is designed to fill into
// available width depending on the context. Returns 0.
func (_ecda *Table) Width() float64 { return 0 }
func _dbba(_fcffd *templateProcessor, _afadb *templateNode) (interface{}, error) {
	return _fcffd.parseImage(_afadb)
}

// SetPadding sets the padding of the component. The padding represents
// inner margins which are applied around the contents of the division.
// The background of the component is not affected by its padding.
func (_cafcg *Division) SetPadding(left, right, top, bottom float64) {
	_cafcg._agda.Left = left
	_cafcg._agda.Right = right
	_cafcg._agda.Top = top
	_cafcg._agda.Bottom = bottom
}

// RadialShading holds information that will be used to render a radial shading.
type RadialShading struct {
	_fcacc  *shading
	_bbccc  *_ggc.PdfRectangle
	_dedge  AnchorPoint
	_cafd   float64
	_eagf   float64
	_efdd   float64
	_cagbac float64
}

func (_fcfba *templateProcessor) parseRadialGradientAttr(creator *Creator, _ededa string) Color {
	_fcgf := ColorBlack
	if _ededa == "" {
		return _fcgf
	}
	var (
		_adef  error
		_cdbga = 0.0
		_ccfe  = 0.0
		_efge  = -1.0
		_cadaa = _dc.Split(_ededa[16:len(_ededa)-1], "\u002c")
	)
	_ecbg := _dc.Fields(_cadaa[0])
	if len(_ecbg) == 2 && _dc.TrimSpace(_ecbg[0])[0] != '#' {
		_cdbga, _adef = _a.ParseFloat(_ecbg[0], 64)
		if _adef != nil {
			_ca.Log.Debug("\u0046a\u0069\u006ce\u0064\u0020\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0072a\u0064\u0069\u0061\u006c\u0020\u0067r\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0058\u0020\u0070\u006fs\u0069\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076", _adef)
		}
		_ccfe, _adef = _a.ParseFloat(_ecbg[1], 64)
		if _adef != nil {
			_ca.Log.Debug("\u0046a\u0069\u006ce\u0064\u0020\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0072a\u0064\u0069\u0061\u006c\u0020\u0067r\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0059\u0020\u0070\u006fs\u0069\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076", _adef)
		}
		_cadaa = _cadaa[1:]
	}
	_adefc := _dc.TrimSpace(_cadaa[0])
	if _adefc[0] != '#' {
		_efge, _adef = _a.ParseFloat(_adefc, 64)
		if _adef != nil {
			_ca.Log.Debug("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0070\u0061\u0072\u0073\u0069\u006eg\u0020\u0072\u0061\u0064\u0069\u0061l\u0020\u0067\u0072\u0061\u0064\u0069\u0065\u006e\u0074\u0020\u0073\u0069\u007ae\u003a\u0020\u0025\u0076", _adef)
		}
		_cadaa = _cadaa[1:]
	}
	_gbcdd, _bedd := _fcfba.processGradientColorPair(_cadaa)
	if _gbcdd == nil || _bedd == nil {
		return _fcgf
	}
	_dfgdd := creator.NewRadialGradientColor(_cdbga, _ccfe, 0, _efge, []*ColorPoint{})
	for _adeg := 0; _adeg < len(_gbcdd); _adeg++ {
		_dfgdd.AddColorStop(_gbcdd[_adeg], _bedd[_adeg])
	}
	return _dfgdd
}

// SetPositioning sets the positioning of the rectangle (absolute or relative).
func (_gdadc *Rectangle) SetPositioning(position Positioning) { _gdadc._bbgac = position }

// SetIndent sets the cell's left indent.
func (_agaf *TableCell) SetIndent(indent float64) { _agaf._bbgcg = indent }

// SetColumnWidths sets the fractional column widths.
// Each width should be in the range 0-1 and is a fraction of the table width.
// The number of width inputs must match number of columns, otherwise an error is returned.
func (_bbaff *Table) SetColumnWidths(widths ...float64) error {
	if len(widths) != _bbaff._aeaa {
		_ca.Log.Debug("M\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0069\u006e\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020\u0077\u0069\u0064\u0074\u0068\u0073\u0020\u0061nd\u0020\u0063\u006fl\u0075m\u006e\u0073")
		return _fa.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_bbaff._abbbf = widths
	return nil
}
func (_gdbfa *templateProcessor) getNodeErrorLocation(_gcdgd *templateNode, _ddfg string, _dgbg ...interface{}) string {
	_adec := _df.Sprintf(_ddfg, _dgbg...)
	_bbccf := _df.Sprintf("\u0025\u0064", _gcdgd._gcfag)
	if _gcdgd._febd != 0 {
		_bbccf = _df.Sprintf("\u0025\u0064\u003a%\u0064", _gcdgd._febd, _gcdgd._fgfag)
	}
	if _gdbfa._eccabe != "" {
		return _df.Sprintf("\u0025\u0073\u0020\u005b\u0025\u0073\u003a\u0025\u0073\u005d", _adec, _gdbfa._eccabe, _bbccf)
	}
	return _df.Sprintf("\u0025s\u0020\u005b\u0025\u0073\u005d", _adec, _bbccf)
}

// NewGraphicSVGFromFile creates a graphic SVG from a file.
func NewGraphicSVGFromFile(path string) (*GraphicSVG, error) { return _cbcc(path) }

// SetColorBottom sets border color for bottom.
func (_dgcf *border) SetColorBottom(col Color) { _dgcf._agg = col }

// GeneratePageBlocks generates the page blocks for the Division component.
// Multiple blocks are generated if the contents wrap over multiple pages.
func (_dgbe *Division) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var (
		_edgg  []*Block
		_dgbee bool
		_dfgcc error
		_dafa  = _dgbe._badg.IsRelative()
		_gccd  = _dgbe._debb.Top
	)
	if _dafa && !_dgbe._ggdbd && !_dgbe._aeadf {
		_cbfg := _dgbe.ctxHeight(ctx.Width)
		if _cbfg > ctx.Height-_dgbe._debb.Top && _cbfg <= ctx.PageHeight-ctx.Margins.Top-ctx.Margins.Bottom {
			if _edgg, ctx, _dfgcc = _cbag().GeneratePageBlocks(ctx); _dfgcc != nil {
				return nil, ctx, _dfgcc
			}
			_dgbee = true
			_gccd = 0
		}
	}
	_daca := ctx
	_ccea := ctx
	if _dafa {
		ctx.X += _dgbe._debb.Left
		ctx.Y += _gccd
		ctx.Width -= _dgbe._debb.Left + _dgbe._debb.Right
		ctx.Height -= _gccd
		_ccea = ctx
		ctx.X += _dgbe._agda.Left
		ctx.Y += _dgbe._agda.Top
		ctx.Width -= _dgbe._agda.Left + _dgbe._agda.Right
		ctx.Height -= _dgbe._agda.Top
		ctx.Margins.Top += _dgbe._agda.Top
		ctx.Margins.Bottom += _dgbe._agda.Bottom
		ctx.Margins.Left += _dgbe._debb.Left + _dgbe._agda.Left
		ctx.Margins.Right += _dgbe._debb.Right + _dgbe._agda.Right
	}
	ctx.Inline = _dgbe._aeadf
	_cdbg := ctx
	_dggd := ctx
	var _eddb float64
	for _, _gcab := range _dgbe._bfdf {
		if ctx.Inline {
			if (ctx.X-_cdbg.X)+_gcab.Width() <= ctx.Width {
				ctx.Y = _dggd.Y
				ctx.Height = _dggd.Height
			} else {
				ctx.X = _cdbg.X
				ctx.Width = _cdbg.Width
				_dggd.Y += _eddb
				_dggd.Height -= _eddb
				_eddb = 0
			}
		}
		_egec, _gga, _bgab := _gcab.GeneratePageBlocks(ctx)
		if _bgab != nil {
			_ca.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006eg\u0020p\u0061\u0067\u0065\u0020\u0062\u006c\u006f\u0063\u006b\u0073\u003a\u0020\u0025\u0076", _bgab)
			return nil, ctx, _bgab
		}
		if len(_egec) < 1 {
			continue
		}
		if len(_edgg) > 0 {
			_edgg[len(_edgg)-1].mergeBlocks(_egec[0])
			_edgg = append(_edgg, _egec[1:]...)
		} else {
			if _ceg := _egec[0]._cad; _ceg == nil || len(*_ceg) == 0 {
				_dgbee = true
			}
			_edgg = append(_edgg, _egec[0:]...)
		}
		if ctx.Inline {
			if ctx.Page != _gga.Page {
				_cdbg.Y = ctx.Margins.Top
				_cdbg.Height = ctx.PageHeight - ctx.Margins.Top
				_dggd.Y = _cdbg.Y
				_dggd.Height = _cdbg.Height
				_eddb = _gga.Height - _cdbg.Height
			} else {
				if _fabea := ctx.Height - _gga.Height; _fabea > _eddb {
					_eddb = _fabea
				}
			}
		} else {
			_gga.X = ctx.X
		}
		ctx = _gga
	}
	ctx.Inline = _daca.Inline
	ctx.Margins = _daca.Margins
	if _dafa {
		ctx.X = _daca.X
		ctx.Width = _daca.Width
		ctx.Y += _dgbe._agda.Bottom
		ctx.Height -= _dgbe._agda.Bottom
	}
	if _dgbe._fbgb != nil {
		_edgg, _dfgcc = _dgbe.drawBackground(_edgg, _ccea, ctx, _dgbee)
		if _dfgcc != nil {
			return nil, ctx, _dfgcc
		}
	}
	if _dgbe._badg.IsAbsolute() {
		return _edgg, _daca, nil
	}
	ctx.Y += _dgbe._debb.Bottom
	ctx.Height -= _dgbe._debb.Bottom
	return _edgg, ctx, nil
}

// MoveY moves the drawing context to absolute position y.
func (_cceg *Creator) MoveY(y float64) { _cceg._eacd.Y = y }
func _gcag(_acbcf, _edcg TextStyle) *Invoice {
	_befce := &Invoice{_bfbbd: "\u0049N\u0056\u004f\u0049\u0043\u0045", _cedd: "\u002c\u0020", _bdcd: _acbcf, _febc: _edcg}
	_befce._cfbbb = &InvoiceAddress{Separator: _befce._cedd}
	_befce._dcfc = &InvoiceAddress{Heading: "\u0042i\u006c\u006c\u0020\u0074\u006f", Separator: _befce._cedd}
	_ffeg := ColorRGBFrom8bit(245, 245, 245)
	_aaec := ColorRGBFrom8bit(155, 155, 155)
	_befce._eefc = _edcg
	_befce._eefc.Color = _aaec
	_befce._eefc.FontSize = 20
	_befce._acecg = _acbcf
	_befce._debc = _edcg
	_befce._dcfab = _acbcf
	_befce._daag = _edcg
	_befce._fcba = _befce.NewCellProps()
	_befce._fcba.BackgroundColor = _ffeg
	_befce._fcba.TextStyle = _edcg
	_befce._dece = _befce.NewCellProps()
	_befce._dece.TextStyle = _edcg
	_befce._dece.BackgroundColor = _ffeg
	_befce._dece.BorderColor = _ffeg
	_befce._degfg = _befce.NewCellProps()
	_befce._degfg.BorderColor = _ffeg
	_befce._degfg.BorderSides = []CellBorderSide{CellBorderSideBottom}
	_befce._degfg.Alignment = CellHorizontalAlignmentRight
	_befce._ddde = _befce.NewCellProps()
	_befce._ddde.Alignment = CellHorizontalAlignmentRight
	_befce._geeb = [2]*InvoiceCell{_befce.newCell("\u0049\u006e\u0076\u006f\u0069\u0063\u0065\u0020\u006eu\u006d\u0062\u0065\u0072", _befce._fcba), _befce.newCell("", _befce._fcba)}
	_befce._ddbe = [2]*InvoiceCell{_befce.newCell("\u0044\u0061\u0074\u0065", _befce._fcba), _befce.newCell("", _befce._fcba)}
	_befce._fdfc = [2]*InvoiceCell{_befce.newCell("\u0044\u0075\u0065\u0020\u0044\u0061\u0074\u0065", _befce._fcba), _befce.newCell("", _befce._fcba)}
	_befce._befc = [2]*InvoiceCell{_befce.newCell("\u0053\u0075\u0062\u0074\u006f\u0074\u0061\u006c", _befce._ddde), _befce.newCell("", _befce._ddde)}
	_bcaa := _befce._ddde
	_bcaa.TextStyle = _edcg
	_bcaa.BackgroundColor = _ffeg
	_bcaa.BorderColor = _ffeg
	_befce._ada = [2]*InvoiceCell{_befce.newCell("\u0054\u006f\u0074a\u006c", _bcaa), _befce.newCell("", _bcaa)}
	_befce._ecf = [2]string{"\u004e\u006f\u0074e\u0073", ""}
	_befce._aacd = [2]string{"T\u0065r\u006d\u0073\u0020\u0061\u006e\u0064\u0020\u0063o\u006e\u0064\u0069\u0074io\u006e\u0073", ""}
	_befce._eag = []*InvoiceCell{_befce.newColumn("D\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e", CellHorizontalAlignmentLeft), _befce.newColumn("\u0051\u0075\u0061\u006e\u0074\u0069\u0074\u0079", CellHorizontalAlignmentRight), _befce.newColumn("\u0055\u006e\u0069\u0074\u0020\u0070\u0072\u0069\u0063\u0065", CellHorizontalAlignmentRight), _befce.newColumn("\u0041\u006d\u006f\u0075\u006e\u0074", CellHorizontalAlignmentRight)}
	return _befce
}
func (_gffb *StyledParagraph) wrapText() error { return _gffb.wrapChunks(true) }

// SetLinePageStyle sets the style for the page part of all new lines
// of the table of contents.
func (_bagbb *TOC) SetLinePageStyle(style TextStyle) { _bagbb._abgad = style }

// Width returns the width of the Paragraph.
func (_adgd *Paragraph) Width() float64 {
	if _adgd._bebg && int(_adgd._dcdb) > 0 {
		return _adgd._dcdb
	}
	return _adgd.getTextWidth() / 1000.0
}

// SetAntiAlias enables anti alias config.
//
// Anti alias is disabled by default.
func (_cgdef *RadialShading) SetAntiAlias(enable bool) { _cgdef._fcacc.SetAntiAlias(enable) }

// ScaleToWidth scales the Block to a specified width, maintaining the same aspect ratio.
func (_gbf *Block) ScaleToWidth(w float64) { _gcc := w / _gbf._ecd; _gbf.Scale(_gcc, _gcc) }

// Width returns the Block's width.
func (_cdd *Block) Width() float64 { return _cdd._ecd }

// SetLineLevelOffset sets the amount of space an indentation level occupies
// for all new lines of the table of contents.
func (_gfda *TOC) SetLineLevelOffset(levelOffset float64) { _gfda._ebgfa = levelOffset }
func (_dcfd *Creator) initContext() {
	_dcfd._eacd.X = _dcfd._gcgd.Left
	_dcfd._eacd.Y = _dcfd._gcgd.Top
	_dcfd._eacd.Width = _dcfd._abf - _dcfd._gcgd.Right - _dcfd._gcgd.Left
	_dcfd._eacd.Height = _dcfd._ffc - _dcfd._gcgd.Bottom - _dcfd._gcgd.Top
	_dcfd._eacd.PageHeight = _dcfd._ffc
	_dcfd._eacd.PageWidth = _dcfd._abf
	_dcfd._eacd.Margins = _dcfd._gcgd
	_dcfd._eacd._dgce = _dcfd.UnsupportedCharacterReplacement
}

// PolyBezierCurve represents a composite curve that is the result of joining
// multiple cubic Bezier curves.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type PolyBezierCurve struct {
	_ffbb  *_fc.PolyBezierCurve
	_bcdgc float64
	_gfcgc float64
	_bbaec Color
}

// GetIndent get the cell's left indent.
func (_ecfa *TableCell) GetIndent() float64 { return _ecfa._bbgcg }

// Margins returns the margins of the list: left, right, top, bottom.
func (_ecebd *List) Margins() (float64, float64, float64, float64) {
	return _ecebd._fed.Left, _ecebd._fed.Right, _ecebd._fed.Top, _ecebd._fed.Bottom
}
func _cefg(_fdac VectorDrawable, _fddf float64) float64 {
	switch _cbdfc := _fdac.(type) {
	case *Paragraph:
		if _cbdfc._bebg {
			_cbdfc.SetWidth(_fddf - _cbdfc._fgcbf.Left - _cbdfc._fgcbf.Right)
		}
		return _cbdfc.Height() + _cbdfc._fgcbf.Top + _cbdfc._fgcbf.Bottom
	case *StyledParagraph:
		if _cbdfc._gfbb {
			_cbdfc.SetWidth(_fddf - _cbdfc._fbgbc.Left - _cbdfc._fbgbc.Right)
		}
		return _cbdfc.Height() + _cbdfc._fbgbc.Top + _cbdfc._fbgbc.Bottom
	case *Image:
		_cbdfc.applyFitMode(_fddf)
		return _cbdfc.Height() + _cbdfc._cbea.Top + _cbdfc._cbea.Bottom
	case *Rectangle:
		_cbdfc.applyFitMode(_fddf)
		return _cbdfc.Height() + _cbdfc._ebfce.Top + _cbdfc._ebfce.Bottom + _cbdfc._decec
	case *Ellipse:
		_cbdfc.applyFitMode(_fddf)
		return _cbdfc.Height() + _cbdfc._cadg.Top + _cbdfc._cadg.Bottom
	case *Division:
		return _cbdfc.ctxHeight(_fddf) + _cbdfc._debb.Top + _cbdfc._debb.Bottom + _cbdfc._agda.Top + _cbdfc._agda.Bottom
	case *Table:
		_cbdfc.updateRowHeights(_fddf - _cbdfc._gfbcc.Left - _cbdfc._gfbcc.Right)
		return _cbdfc.Height() + _cbdfc._gfbcc.Top + _cbdfc._gfbcc.Bottom
	case *List:
		return _cbdfc.ctxHeight(_fddf) + _cbdfc._fed.Top + _cbdfc._fed.Bottom
	case marginDrawable:
		_, _, _abcf, _adde := _cbdfc.GetMargins()
		return _cbdfc.Height() + _abcf + _adde
	default:
		return _cbdfc.Height()
	}
}

// SetFitMode sets the fit mode of the ellipse.
// NOTE: The fit mode is only applied if relative positioning is used.
func (_eaad *Ellipse) SetFitMode(fitMode FitMode) { _eaad._gbfg = fitMode }

// GetMargins returns the margins of the TOC line: left, right, top, bottom.
func (_gcgdf *TOCLine) GetMargins() (float64, float64, float64, float64) {
	_dgcee := &_gcgdf._ecde._fbgbc
	return _gcgdf._adbfe, _dgcee.Right, _dgcee.Top, _dgcee.Bottom
}
func (_ffag *StyledParagraph) getMaxLineWidth() float64 {
	if _ffag._aabba == nil || len(_ffag._aabba) == 0 {
		_ffag.wrapText()
	}
	var _bgfdg float64
	for _, _cedf := range _ffag._aabba {
		_abeed := _ffag.getTextLineWidth(_cedf)
		if _abeed > _bgfdg {
			_bgfdg = _abeed
		}
	}
	return _bgfdg
}
func (_bagce *templateProcessor) parseRectangle(_dgac *templateNode) (interface{}, error) {
	_ebdc := _bagce.creator.NewRectangle(0, 0, 0, 0)
	for _, _ceaf := range _dgac._gbdee.Attr {
		_gaad := _ceaf.Value
		switch _daeb := _ceaf.Name.Local; _daeb {
		case "\u0078":
			_ebdc._cgedd = _bagce.parseFloatAttr(_daeb, _gaad)
		case "\u0079":
			_ebdc._eeff = _bagce.parseFloatAttr(_daeb, _gaad)
		case "\u0077\u0069\u0064t\u0068":
			_ebdc.SetWidth(_bagce.parseFloatAttr(_daeb, _gaad))
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_ebdc.SetHeight(_bagce.parseFloatAttr(_daeb, _gaad))
		case "\u0066\u0069\u006c\u006c\u002d\u0063\u006f\u006c\u006f\u0072":
			_ebdc.SetFillColor(_bagce.parseColorAttr(_daeb, _gaad))
		case "\u0066\u0069\u006cl\u002d\u006f\u0070\u0061\u0063\u0069\u0074\u0079":
			_ebdc.SetFillOpacity(_bagce.parseFloatAttr(_daeb, _gaad))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072":
			_ebdc.SetBorderColor(_bagce.parseColorAttr(_daeb, _gaad))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u006f\u0070a\u0063\u0069\u0074\u0079":
			_ebdc.SetBorderOpacity(_bagce.parseFloatAttr(_daeb, _gaad))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068":
			_ebdc.SetBorderWidth(_bagce.parseFloatAttr(_daeb, _gaad))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0072\u0061\u0064\u0069\u0075\u0073":
			_dafb, _ggbf, _aedfg, _fcea := _bagce.parseBorderRadiusAttr(_daeb, _gaad)
			_ebdc.SetBorderRadius(_dafb, _ggbf, _fcea, _aedfg)
		case "\u0062\u006f\u0072\u0064er\u002d\u0074\u006f\u0070\u002d\u006c\u0065\u0066\u0074\u002d\u0072\u0061\u0064\u0069u\u0073":
			_ebdc._gagef = _bagce.parseFloatAttr(_daeb, _gaad)
		case "\u0062\u006f\u0072de\u0072\u002d\u0074\u006f\u0070\u002d\u0072\u0069\u0067\u0068\u0074\u002d\u0072\u0061\u0064\u0069\u0075\u0073":
			_ebdc._efaga = _bagce.parseFloatAttr(_daeb, _gaad)
		case "\u0062o\u0072\u0064\u0065\u0072-\u0062\u006f\u0074\u0074\u006fm\u002dl\u0065f\u0074\u002d\u0072\u0061\u0064\u0069\u0075s":
			_ebdc._fbcea = _bagce.parseFloatAttr(_daeb, _gaad)
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0062\u006f\u0074\u0074o\u006d\u002d\u0072\u0069\u0067\u0068\u0074\u002d\u0072\u0061d\u0069\u0075\u0073":
			_ebdc._fdcc = _bagce.parseFloatAttr(_daeb, _gaad)
		case "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e":
			_ebdc.SetPositioning(_bagce.parsePositioningAttr(_daeb, _gaad))
		case "\u0066\u0069\u0074\u002d\u006d\u006f\u0064\u0065":
			_ebdc.SetFitMode(_bagce.parseFitModeAttr(_daeb, _gaad))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_ggfca := _bagce.parseMarginAttr(_daeb, _gaad)
			_ebdc.SetMargins(_ggfca.Left, _ggfca.Right, _ggfca.Top, _ggfca.Bottom)
		default:
			_bagce.nodeLogDebug(_dgac, "\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065\u0064\u0020re\u0063\u0074\u0061\u006e\u0067\u006ce\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073`\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u002e", _daeb)
		}
	}
	return _ebdc, nil
}

// SetBorderOpacity sets the border opacity.
func (_cbec *CurvePolygon) SetBorderOpacity(opacity float64) { _cbec._fgcc = opacity }

// SetIncludeInTOC sets a flag to indicate whether or not to include in tOC.
func (_ggb *Chapter) SetIncludeInTOC(includeInTOC bool) { _ggb._afg = includeInTOC }

// MultiCell makes a new cell with the specified row span and col span
// and inserts it into the table at the current position.
func (_babed *Table) MultiCell(rowspan, colspan int) *TableCell {
	_babed._ggfb++
	_dcfcc := (_babed.moveToNextAvailableCell()-1)%(_babed._aeaa) + 1
	_dbdee := (_babed._ggfb-1)/_babed._aeaa + 1
	for _dbdee > _babed._fgbfga {
		_babed._fgbfga++
		_babed._begb = append(_babed._begb, _babed._ddfaf)
	}
	_bfefa := &TableCell{}
	_bfefa._deded = _dbdee
	_bfefa._eafbd = _dcfcc
	_bfefa._bbgcg = 5
	_bfefa._ffdeb = CellBorderStyleNone
	_bfefa._fcagf = _fc.LineStyleSolid
	_bfefa._abad = CellHorizontalAlignmentLeft
	_bfefa._fddca = CellVerticalAlignmentTop
	_bfefa._egbd = 0
	_bfefa._egaeb = 0
	_bfefa._ddbdf = 0
	_bfefa._caffa = 0
	_fged := ColorBlack
	_bfefa._egde = _fged
	_bfefa._dagd = _fged
	_bfefa._aeec = _fged
	_bfefa._afga = _fged
	if rowspan < 1 {
		_ca.Log.Debug("\u0054\u0061\u0062\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0072\u006f\u0077\u0073\u0070a\u006e\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061t\u0020\u0031\u0020\u0028\u0025\u0064\u0029\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063e\u006c\u006c\u0020\u0072\u006f\u0077s\u0070\u0061n\u0020\u0074o\u00201\u002e", rowspan)
		rowspan = 1
	}
	_becce := _babed._fgbfga - (_bfefa._deded - 1)
	if rowspan > _becce {
		_ca.Log.Debug("\u0054\u0061b\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0072\u006f\u0077\u0073\u0070\u0061\u006e\u0020\u0028\u0025d\u0029\u0020\u0065\u0078\u0063\u0065e\u0064\u0073\u0020\u0072\u0065\u006d\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0072o\u0077\u0073 \u0028\u0025\u0064\u0029.\u0020\u0041\u0064\u0064\u0069n\u0067\u0020\u0072\u006f\u0077\u0073\u002e", rowspan, _becce)
		_babed._fgbfga += rowspan - 1
		for _ggee := 0; _ggee <= rowspan-_becce; _ggee++ {
			_babed._begb = append(_babed._begb, _babed._ddfaf)
		}
	}
	for _eeedg := 0; _eeedg < colspan && _dcfcc+_eeedg-1 < len(_babed._degfe); _eeedg++ {
		_babed._degfe[_dcfcc+_eeedg-1] = rowspan - 1
	}
	_bfefa._ffgdb = rowspan
	if colspan < 1 {
		_ca.Log.Debug("\u0054\u0061\u0062\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0063\u006f\u006c\u0073\u0070a\u006e\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061n\u0020\u0031\u0020\u0028\u0025\u0064\u0029\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063e\u006c\u006c\u0020\u0063\u006f\u006cs\u0070\u0061n\u0020\u0074o\u00201\u002e", colspan)
		colspan = 1
	}
	_bcdd := _babed._aeaa - (_bfefa._eafbd - 1)
	if colspan > _bcdd {
		_ca.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0065\u006c\u006c\u0020\u0063o\u006c\u0073\u0070\u0061\u006e\u0020\u0028\u0025\u0064\u0029\u0020\u0065\u0078\u0063\u0065\u0065\u0064\u0073\u0020\u0072\u0065\u006d\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0072\u006f\u0077\u0020\u0063\u006f\u006c\u0073\u0020\u0028\u0025d\u0029\u002e\u0020\u0041\u0064\u006a\u0075\u0073\u0074\u0069\u006e\u0067 \u0063\u006f\u006c\u0073\u0070\u0061n\u002e", colspan, _bcdd)
		colspan = _bcdd
	}
	_bfefa._abfda = colspan
	_babed._ggfb += colspan - 1
	_babed._cacca = append(_babed._cacca, _bfefa)
	_bfefa._ddggg = _babed
	return _bfefa
}

// GeneratePageBlocks generates the table page blocks. Multiple blocks are
// generated if the contents wrap over multiple pages.
// Implements the Drawable interface.
func (_fbab *Table) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_bbddg := _fbab
	if _fbab._eebgf {
		_bbddg = _fbab.clone()
	}
	return _aebfg(_bbddg, ctx)
}
func _agbf(_fdcgc string, _dgba TextStyle) *Paragraph {
	_dgbcg := &Paragraph{_age: _fdcgc, _fggb: _dgba.Font, _fcbfa: _dgba.FontSize, _dacae: 1.0, _bebg: true, _ggad: true, _abd: TextAlignmentLeft, _dgeg: 0, _geca: 1, _bfeff: 1, _abea: PositionRelative}
	_dgbcg.SetColor(_dgba.Color)
	return _dgbcg
}

// ColorPoint is a pair of Color and a relative point where the color
// would be rendered.
type ColorPoint struct {
	_afea  Color
	_dbcfd float64
}

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

// SetLineColor sets the line color.
func (_dafd *Polyline) SetLineColor(color Color) { _dafd._edbcg.LineColor = _dbac(color) }

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_afab *TOCLine) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_aefae := ctx
	_ebada, ctx, _fbcgb := _afab._ecde.GeneratePageBlocks(ctx)
	if _fbcgb != nil {
		return _ebada, ctx, _fbcgb
	}
	if _afab._dggfc.IsRelative() {
		ctx.X = _aefae.X
	}
	if _afab._dggfc.IsAbsolute() {
		return _ebada, _aefae, nil
	}
	return _ebada, ctx, nil
}

// BorderWidth returns the border width of the ellipse.
func (_ggbc *Ellipse) BorderWidth() float64 { return _ggbc._fgcca }

// NewBlockFromPage creates a Block from a PDF Page.  Useful for loading template pages as blocks
// from a PDF document and additional content with the creator.
func NewBlockFromPage(page *_ggc.PdfPage) (*Block, error) {
	_dgc := &Block{}
	_gd, _gc := page.GetAllContentStreams()
	if _gc != nil {
		return nil, _gc
	}
	_dga := _bdb.NewContentStreamParser(_gd)
	_bgg, _gc := _dga.Parse()
	if _gc != nil {
		return nil, _gc
	}
	_bgg.WrapIfNeeded()
	_dgc._cad = _bgg
	if page.Resources != nil {
		_dgc._ge = page.Resources
	} else {
		_dgc._ge = _ggc.NewPdfPageResources()
	}
	_fg, _gc := page.GetMediaBox()
	if _gc != nil {
		return nil, _gc
	}
	if _fg.Llx != 0 || _fg.Lly != 0 {
		_dgc.translate(-_fg.Llx, _fg.Lly)
	}
	_dgc._ecd = _fg.Urx - _fg.Llx
	_dgc._gfe = _fg.Ury - _fg.Lly
	if page.Rotate != nil {
		_dgc._de = -float64(*page.Rotate)
	}
	return _dgc, nil
}
func (_ddgae *templateProcessor) parseTextChunk(_efdb *templateNode, _bedf *TextChunk) (interface{}, error) {
	if _efdb._fbcg == nil {
		_ddgae.nodeLogError(_efdb, "\u0054\u0065\u0078\u0074\u0020\u0063\u0068\u0075\u006e\u006b\u0020\u0070\u0061\u0072\u0065n\u0074 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e")
		return nil, _gfaba
	}
	var (
		_ebfg  = _ddgae.creator.NewTextStyle()
		_dgdcc bool
	)
	for _, _fbbc := range _efdb._gbdee.Attr {
		if _fbbc.Name.Local == "\u006c\u0069\u006e\u006b" {
			_dggc, _gfef := _efdb._fbcg._caacd.(*StyledParagraph)
			if !_gfef {
				_ddgae.nodeLogError(_efdb, "\u004c\u0069\u006e\u006b \u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065 \u006f\u006e\u006c\u0079\u0020\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0062\u006c\u0065\u0020\u0074\u006f \u0070\u0061\u0072\u0061\u0067r\u0061\u0070\u0068\u0027\u0073\u0020\u0074\u0065\u0078\u0074\u0020\u0063\u0068\u0075\u006e\u006b\u002e")
				_dgdcc = true
			} else {
				_ebfg = _dggc._aecf
			}
			break
		}
	}
	if _bedf == nil {
		_bedf = NewTextChunk("", _ebfg)
	}
	for _, _fcaag := range _efdb._gbdee.Attr {
		_defe := _fcaag.Value
		switch _ceefe := _fcaag.Name.Local; _ceefe {
		case "\u0063\u006f\u006co\u0072":
			_bedf.Style.Color = _ddgae.parseColorAttr(_ceefe, _defe)
		case "\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u002d\u0063\u006f\u006c\u006f\u0072":
			_bedf.Style.OutlineColor = _ddgae.parseColorAttr(_ceefe, _defe)
		case "\u0066\u006f\u006e\u0074":
			_bedf.Style.Font = _ddgae.parseFontAttr(_ceefe, _defe)
		case "\u0066o\u006e\u0074\u002d\u0073\u0069\u007ae":
			_bedf.Style.FontSize = _ddgae.parseFloatAttr(_ceefe, _defe)
		case "\u006f\u0075\u0074l\u0069\u006e\u0065\u002d\u0073\u0069\u007a\u0065":
			_bedf.Style.OutlineSize = _ddgae.parseFloatAttr(_ceefe, _defe)
		case "\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u002d\u0073\u0070a\u0063\u0069\u006e\u0067":
			_bedf.Style.CharSpacing = _ddgae.parseFloatAttr(_ceefe, _defe)
		case "\u0068o\u0072i\u007a\u006f\u006e\u0074\u0061l\u002d\u0073c\u0061\u006c\u0069\u006e\u0067":
			_bedf.Style.HorizontalScaling = _ddgae.parseFloatAttr(_ceefe, _defe)
		case "\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067-\u006d\u006f\u0064\u0065":
			_bedf.Style.RenderingMode = _ddgae.parseTextRenderingModeAttr(_ceefe, _defe)
		case "\u0075n\u0064\u0065\u0072\u006c\u0069\u006ee":
			_bedf.Style.Underline = _ddgae.parseBoolAttr(_ceefe, _defe)
		case "\u0075n\u0064e\u0072\u006c\u0069\u006e\u0065\u002d\u0063\u006f\u006c\u006f\u0072":
			_bedf.Style.UnderlineStyle.Color = _ddgae.parseColorAttr(_ceefe, _defe)
		case "\u0075\u006ed\u0065\u0072\u006ci\u006e\u0065\u002d\u006f\u0066\u0066\u0073\u0065\u0074":
			_bedf.Style.UnderlineStyle.Offset = _ddgae.parseFloatAttr(_ceefe, _defe)
		case "\u0075\u006e\u0064\u0065rl\u0069\u006e\u0065\u002d\u0074\u0068\u0069\u0063\u006b\u006e\u0065\u0073\u0073":
			_bedf.Style.UnderlineStyle.Thickness = _ddgae.parseFloatAttr(_ceefe, _defe)
		case "\u006c\u0069\u006e\u006b":
			if !_dgdcc {
				_bedf._dfcb = _ddgae.parseLinkAttr(_ceefe, _defe)
			}
		case "\u0074e\u0078\u0074\u002d\u0072\u0069\u0073e":
			_bedf.Style.TextRise = _ddgae.parseFloatAttr(_ceefe, _defe)
		default:
			_ddgae.nodeLogDebug(_efdb, "\u0055\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0063\u0068\u0075\u006e\u006b\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e", _ceefe)
		}
	}
	return _bedf, nil
}

// SetHeight sets the height of the rectangle.
func (_bccg *Rectangle) SetHeight(height float64) { _bccg._fefg = height }

// CurRow returns the currently active cell's row number.
func (_bfced *Table) CurRow() int { _dadfd := (_bfced._ggfb-1)/_bfced._aeaa + 1; return _dadfd }

// Drawable is a widget that can be used to draw with the Creator.
type Drawable interface {

	// GeneratePageBlocks draw onto blocks representing Page contents. As the content can wrap over many pages, multiple
	// templates are returned, one per Page.  The function also takes a draw context containing information
	// where to draw (if relative positioning) and the available height to draw on accounting for Margins etc.
	GeneratePageBlocks(_cgdd DrawContext) ([]*Block, DrawContext, error)
}

func (_ageee *templateProcessor) parseCellVerticalAlignmentAttr(_eagc, _fbdb string) CellVerticalAlignment {
	_ca.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u0065\u006c\u006c\u0020\u0076\u0065r\u0074\u0069\u0063\u0061\u006c\u0020\u0061\u006c\u0069\u0067\u006e\u006d\u0065n\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a (\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _eagc, _fbdb)
	_aaaed := map[string]CellVerticalAlignment{"\u0074\u006f\u0070": CellVerticalAlignmentTop, "\u006d\u0069\u0064\u0064\u006c\u0065": CellVerticalAlignmentMiddle, "\u0062\u006f\u0074\u0074\u006f\u006d": CellVerticalAlignmentBottom}[_fbdb]
	return _aaaed
}

// SetFillColor sets the fill color.
func (_bcfgb *Polygon) SetFillColor(color Color) {
	_bcfgb._ddcb = color
	_bcfgb._gbda.FillColor = _dbac(color)
}

// SetBorderColor sets the border color.
func (_babg *CurvePolygon) SetBorderColor(color Color) { _babg._ecae.BorderColor = _dbac(color) }
func (_abfd *Paragraph) getTextWidth() float64 {
	_gecdb := 0.0
	for _, _ebba := range _abfd._age {
		if _ebba == '\u000A' {
			continue
		}
		_cfee, _eddfb := _abfd._fggb.GetRuneMetrics(_ebba)
		if !_eddfb {
			_ca.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052u\u006e\u0065\u0020\u0063\u0068a\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0028\u0072\u0075\u006e\u0065\u0020\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0029", _ebba, _ebba)
			return -1
		}
		_gecdb += _abfd._fcbfa * _cfee.Wx
	}
	return _gecdb
}

// SetLink makes the line an internal link.
// The text parameter represents the text that is displayed.
// The user is taken to the specified page, at the specified x and y
// coordinates. Position 0, 0 is at the top left of the page.
func (_cbbdd *TOCLine) SetLink(page int64, x, y float64) {
	_cbbdd._efdgc = x
	_cbbdd._cafda = y
	_cbbdd._eefcc = page
	_fggc := _cbbdd._ecde._aecf.Color
	_cbbdd.Number.Style.Color = _fggc
	_cbbdd.Title.Style.Color = _fggc
	_cbbdd.Separator.Style.Color = _fggc
	_cbbdd.Page.Style.Color = _fggc
}

// SetBorderOpacity sets the border opacity of the ellipse.
func (_cgdc *Ellipse) SetBorderOpacity(opacity float64) { _cgdc._dadc = opacity }

// SetBorderColor sets the border color.
func (_fgcg *PolyBezierCurve) SetBorderColor(color Color) { _fgcg._ffbb.BorderColor = _dbac(color) }
func (_agcfa *StyledParagraph) wrapChunks(_fgfg bool) error {
	if !_agcfa._gfbb || int(_agcfa._gcadd) <= 0 {
		_agcfa._aabba = [][]*TextChunk{_agcfa._cdffa}
		return nil
	}
	if _agcfa._cbeg {
		_agcfa.wrapWordChunks()
	}
	_agcfa._aabba = [][]*TextChunk{}
	var _agdb []*TextChunk
	var _fcdc float64
	_fbaa := _cd.IsSpace
	if !_fgfg {
		_fbaa = func(rune) bool { return false }
	}
	_caafe := _cece(_agcfa._gcadd*1000.0, 0.000001)
	for _, _bbfaf := range _agcfa._cdffa {
		_added := _bbfaf.Style
		_fdbef := _bbfaf._dfcb
		_egce := _bbfaf.VerticalAlignment
		var (
			_effc []rune
			_feac []float64
		)
		_dagbe := _ddfgf(_bbfaf.Text)
		for _, _bbdcg := range _bbfaf.Text {
			if _bbdcg == '\u000A' {
				if !_fgfg {
					_effc = append(_effc, _bbdcg)
				}
				_agdb = append(_agdb, &TextChunk{Text: _dc.TrimRightFunc(string(_effc), _fbaa), Style: _added, _dfcb: _bgbee(_fdbef), VerticalAlignment: _egce})
				_agcfa._aabba = append(_agcfa._aabba, _agdb)
				_agdb = nil
				_fcdc = 0
				_effc = nil
				_feac = nil
				continue
			}
			_cgdac := _bbdcg == ' '
			_fabbf, _aaeff := _added.Font.GetRuneMetrics(_bbdcg)
			if !_aaeff {
				_ca.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _bbdcg)
				return _fa.New("\u0067\u006c\u0079\u0070\u0068\u0020\u0063\u0068\u0061\u0072\u0020m\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
			}
			_cefd := _added.FontSize * _fabbf.Wx * _added.horizontalScale()
			_bgabg := _cefd
			if !_cgdac {
				_bgabg = _cefd + _added.CharSpacing*1000.0
			}
			if _fcdc+_cefd > _caafe {
				_cfggf := -1
				if !_cgdac {
					for _acad := len(_effc) - 1; _acad >= 0; _acad-- {
						if _effc[_acad] == ' ' {
							_cfggf = _acad
							break
						}
					}
				}
				if _agcfa._cbeg {
					_afcba := len(_agdb)
					if _afcba > 0 {
						_agdb[_afcba-1].Text = _dc.TrimRightFunc(_agdb[_afcba-1].Text, _fbaa)
						_agcfa._aabba = append(_agcfa._aabba, _agdb)
						_agdb = []*TextChunk{}
					}
					_effc = append(_effc, _bbdcg)
					_feac = append(_feac, _bgabg)
					if _cfggf >= 0 {
						_effc = _effc[_cfggf+1:]
						_feac = _feac[_cfggf+1:]
					}
					_fcdc = 0
					for _, _eefcf := range _feac {
						_fcdc += _eefcf
					}
					if _fcdc > _caafe {
						_agae := string(_effc[:len(_effc)-1])
						_agae = _baecf(_agae, _dagbe)
						if !_fgfg && _cgdac {
							_agae += "\u0020"
						}
						_agdb = append(_agdb, &TextChunk{Text: _dc.TrimRightFunc(_agae, _fbaa), Style: _added, _dfcb: _bgbee(_fdbef), VerticalAlignment: _egce})
						_agcfa._aabba = append(_agcfa._aabba, _agdb)
						_agdb = []*TextChunk{}
						_effc = []rune{_bbdcg}
						_feac = []float64{_bgabg}
						_fcdc = _bgabg
					}
					continue
				}
				_ccae := string(_effc)
				if _cfggf >= 0 {
					_ccae = string(_effc[0 : _cfggf+1])
					_effc = _effc[_cfggf+1:]
					_effc = append(_effc, _bbdcg)
					_feac = _feac[_cfggf+1:]
					_feac = append(_feac, _bgabg)
					_fcdc = 0
					for _, _ecbca := range _feac {
						_fcdc += _ecbca
					}
				} else {
					if _cgdac {
						_fcdc = 0
						_effc = []rune{}
						_feac = []float64{}
					} else {
						_fcdc = _bgabg
						_effc = []rune{_bbdcg}
						_feac = []float64{_bgabg}
					}
				}
				_ccae = _baecf(_ccae, _dagbe)
				if !_fgfg && _cgdac {
					_ccae += "\u0020"
				}
				_agdb = append(_agdb, &TextChunk{Text: _dc.TrimRightFunc(_ccae, _fbaa), Style: _added, _dfcb: _bgbee(_fdbef), VerticalAlignment: _egce})
				_agcfa._aabba = append(_agcfa._aabba, _agdb)
				_agdb = []*TextChunk{}
			} else {
				_fcdc += _bgabg
				_effc = append(_effc, _bbdcg)
				_feac = append(_feac, _bgabg)
			}
		}
		if len(_effc) > 0 {
			_cafcgd := _baecf(string(_effc), _dagbe)
			_agdb = append(_agdb, &TextChunk{Text: _cafcgd, Style: _added, _dfcb: _bgbee(_fdbef), VerticalAlignment: _egce})
		}
	}
	if len(_agdb) > 0 {
		_agcfa._aabba = append(_agcfa._aabba, _agdb)
	}
	return nil
}

// Margins represents page margins or margins around an element.
type Margins struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

// GetRowHeight returns the height of the specified row.
func (_gecg *Table) GetRowHeight(row int) (float64, error) {
	if row < 1 || row > len(_gecg._begb) {
		return 0, _fa.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	return _gecg._begb[row-1], nil
}

// Width returns the width of the specified text chunk.
func (_gdead *TextChunk) Width() float64 {
	var (
		_cbde  float64
		_feefe = _gdead.Style
	)
	for _, _cfae := range _gdead.Text {
		_fggea, _bacab := _feefe.Font.GetRuneMetrics(_cfae)
		if !_bacab {
			_ca.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006det\u0072i\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064!\u0020\u0072\u0075\u006e\u0065\u003d\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0020\u0066o\u006e\u0074\u003d\u0025\u0073\u0020\u0025\u0023\u0071", _cfae, _cfae, _feefe.Font.BaseFont(), _feefe.Font.Subtype())
			_ca.Log.Trace("\u0046o\u006e\u0074\u003a\u0020\u0025\u0023v", _feefe.Font)
			_ca.Log.Trace("\u0045\u006e\u0063o\u0064\u0065\u0072\u003a\u0020\u0025\u0023\u0076", _feefe.Font.Encoder())
		}
		_aebfb := _feefe.FontSize * _fggea.Wx
		_ggcf := _aebfb
		if _cfae != ' ' {
			_ggcf = _aebfb + _feefe.CharSpacing*1000.0
		}
		_cbde += _ggcf
	}
	return _cbde / 1000.0
}

// SetPos sets the position of the chart to the specified coordinates.
// This method sets the chart to use absolute positioning.
func (_aga *Chart) SetPos(x, y float64) { _aga._fbc = PositionAbsolute; _aga._bfe = x; _aga._gbbc = y }
func (_cage *templateProcessor) run() error {
	_fbbee := _e.NewDecoder(_g.NewReader(_cage._gcgb))
	var _cdfe *templateNode
	for {
		_faba, _afed := _fbbee.Token()
		if _afed != nil {
			if _afed == _ae.EOF {
				return nil
			}
			return _afed
		}
		if _faba == nil {
			break
		}
		_bddeb, _fggga := _dbfbb(_fbbee)
		_cgdca := _fbbee.InputOffset()
		switch _egdde := _faba.(type) {
		case _e.StartElement:
			_ca.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006eg\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0073\u0074\u0061r\u0074\u0020\u0074\u0061\u0067\u003a\u0020`\u0025\u0073\u0060\u002e", _egdde.Name.Local)
			_beeg, _bfdgc := _acaac[_egdde.Name.Local]
			if !_bfdgc {
				if _cage._eccabe == "" {
					if _bddeb != 0 {
						_ca.Log.Debug("\u0055n\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u006dp\u006c\u0061\u0074\u0065 \u0074\u0061\u0067\u0020\u003c%\u0073\u003e\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e\u0020\u005b%\u0064\u003a\u0025\u0064\u005d", _egdde.Name.Local, _bddeb, _fggga)
					} else {
						_ca.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u0074\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u002e\u0020\u0053\u006b\u0069\u0070\u0070i\u006e\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072e\u0063\u0074\u002e\u0020\u005b%\u0064\u005d", _egdde.Name.Local, _cgdca)
					}
				} else {
					if _bddeb != 0 {
						_ca.Log.Debug("\u0055\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0074e\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0074\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065 \u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e\u0020\u005b%\u0073\u003a\u0025\u0064\u003a\u0025d\u005d", _egdde.Name.Local, _cage._eccabe, _bddeb, _fggga)
					} else {
						_ca.Log.Debug("\u0055n\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u006dp\u006c\u0061\u0074\u0065 \u0074\u0061\u0067\u0020\u003c%\u0073\u003e\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e\u0020\u005b%\u0073\u003a\u0025\u0064\u005d", _egdde.Name.Local, _cage._eccabe, _cgdca)
					}
				}
				continue
			}
			_cdfe = &templateNode{_gbdee: _egdde, _fbcg: _cdfe, _febd: _bddeb, _fgfag: _fggga, _gcfag: _cgdca}
			if _gcbcg := _beeg._bcddf; _gcbcg != nil {
				_cdfe._caacd, _afed = _gcbcg(_cage, _cdfe)
				if _afed != nil {
					return _afed
				}
			}
		case _e.EndElement:
			_ca.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020t\u0065\u006d\u0070\u006c\u0061\u0074\u0065 \u0065\u006e\u0064\u0020\u0074\u0061\u0067\u003a\u0020\u0060\u0025\u0073\u0060\u002e", _egdde.Name.Local)
			if _cdfe != nil {
				if _cdfe._caacd != nil {
					if _cdedf := _cage.renderNode(_cdfe); _cdedf != nil {
						return _cdedf
					}
				}
				_cdfe = _cdfe._fbcg
			}
		case _e.CharData:
			if _cdfe != nil && _cdfe._caacd != nil {
				if _dcbe := _cage.addNodeText(_cdfe, string(_egdde)); _dcbe != nil {
					return _dcbe
				}
			}
		case _e.Comment:
			_ca.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020t\u0065\u006d\u0070\u006c\u0061\u0074\u0065 \u0063\u006f\u006d\u006d\u0065\u006e\u0074\u003a\u0020\u0060\u0025\u0073\u0060\u002e", string(_egdde))
		}
	}
	return nil
}

// SetBorderWidth sets the border width of the ellipse.
func (_gffe *Ellipse) SetBorderWidth(bw float64) { _gffe._fgcca = bw }

const (
	TextVerticalAlignmentBaseline TextVerticalAlignment = iota
	TextVerticalAlignmentCenter
	TextVerticalAlignmentBottom
	TextVerticalAlignmentTop
)

// SetTitle sets the title of the invoice.
func (_gebf *Invoice) SetTitle(title string) { _gebf._bfbbd = title }
func (_fbdd *StyledParagraph) split(_deadf DrawContext) (_cabc, _gaee *StyledParagraph, _bggcc error) {
	if _bggcc = _fbdd.wrapChunks(false); _bggcc != nil {
		return nil, nil, _bggcc
	}
	if len(_fbdd._aabba) == 1 && _fbdd._babcf > _deadf.Height {
		return _fbdd, nil, nil
	}
	_dgdc := func(_feab []*TextChunk, _aaee []*TextChunk) []*TextChunk {
		if len(_aaee) == 0 {
			return _feab
		}
		_bbcg := len(_feab)
		if _bbcg == 0 {
			return append(_feab, _aaee...)
		}
		if _feab[_bbcg-1].Style == _aaee[0].Style {
			_feab[_bbcg-1].Text += _aaee[0].Text
		} else {
			_feab = append(_feab, _aaee[0])
		}
		return append(_feab, _aaee[1:]...)
	}
	_aega := func(_bbab *StyledParagraph, _cacg []*TextChunk) *StyledParagraph {
		if len(_cacg) == 0 {
			return nil
		}
		_eeea := *_bbab
		_eeea._cdffa = _cacg
		return &_eeea
	}
	var (
		_acbcg float64
		_dabd  []*TextChunk
		_deec  []*TextChunk
	)
	for _, _feca := range _fbdd._aabba {
		var _ggfe float64
		_efedb := make([]*TextChunk, 0, len(_feca))
		for _, _dabgf := range _feca {
			if _eddda := _dabgf.Style.FontSize; _eddda > _ggfe {
				_ggfe = _eddda
			}
			_efedb = append(_efedb, _dabgf.clone())
		}
		_ggfe *= _fbdd._babcf
		if _fbdd._gbga.IsRelative() {
			if _acbcg+_ggfe > _deadf.Height {
				_deec = _dgdc(_deec, _efedb)
			} else {
				_dabd = _dgdc(_dabd, _efedb)
			}
		}
		_acbcg += _ggfe
	}
	_fbdd._aabba = nil
	if len(_deec) == 0 {
		return _fbdd, nil, nil
	}
	return _aega(_fbdd, _dabd), _aega(_fbdd, _deec), nil
}

// HeaderFunctionArgs holds the input arguments to a header drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type HeaderFunctionArgs struct {
	PageNum    int
	TotalPages int
}

// SetStyle sets the style for all the line components: number, title,
// separator, page.
func (_eecfe *TOCLine) SetStyle(style TextStyle) {
	_eecfe.Number.Style = style
	_eecfe.Title.Style = style
	_eecfe.Separator.Style = style
	_eecfe.Page.Style = style
}

// SetLevel sets the indentation level of the TOC line.
func (_cfdac *TOCLine) SetLevel(level uint) {
	_cfdac._daebd = level
	_cfdac._ecde._fbgbc.Left = _cfdac._adbfe + float64(_cfdac._daebd-1)*_cfdac._aeeeg
}

// AddSubtable copies the cells of the subtable in the table, starting with the
// specified position. The table row and column indices are 1-based, which
// makes the position of the first cell of the first row of the table 1,1.
// The table is automatically extended if the subtable exceeds its columns.
// This can happen when the subtable has more columns than the table or when
// one or more columns of the subtable starting from the specified position
// exceed the last column of the table.
func (_dfece *Table) AddSubtable(row, col int, subtable *Table) {
	for _, _ebce := range subtable._cacca {
		_gegd := &TableCell{}
		*_gegd = *_ebce
		_gegd._ddggg = _dfece
		_gegd._eafbd += col - 1
		if _cffag := _dfece._aeaa - (_gegd._eafbd - 1); _cffag < _gegd._abfda {
			_dfece._aeaa += _gegd._abfda - _cffag
			_dfece.resetColumnWidths()
			_ca.Log.Debug("\u0054a\u0062l\u0065\u003a\u0020\u0073\u0075\u0062\u0074\u0061\u0062\u006c\u0065 \u0065\u0078\u0063\u0065e\u0064\u0073\u0020\u0064\u0065s\u0074\u0069\u006e\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0061\u0062\u006c\u0065\u002e\u0020\u0045\u0078\u0070\u0061\u006e\u0064\u0069\u006e\u0067\u0020\u0074\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0025\u0064\u0020\u0063\u006fl\u0075\u006d\u006e\u0073\u002e", _dfece._aeaa)
		}
		_gegd._deded += row - 1
		_adcef := subtable._begb[_ebce._deded-1]
		if _gegd._deded > _dfece._fgbfga {
			for _gegd._deded > _dfece._fgbfga {
				_dfece._fgbfga++
				_dfece._begb = append(_dfece._begb, _dfece._ddfaf)
			}
			_dfece._begb[_gegd._deded-1] = _adcef
		} else {
			_dfece._begb[_gegd._deded-1] = _b.Max(_dfece._begb[_gegd._deded-1], _adcef)
		}
		_dfece._cacca = append(_dfece._cacca, _gegd)
	}
	_dfece.sortCells()
}
func _ggba(_eecb *_ggc.Image) (*Image, error) {
	_dfec := float64(_eecb.Width)
	_gcdg := float64(_eecb.Height)
	return &Image{_dffd: _eecb, _cegf: _dfec, _bdff: _gcdg, _cagba: _dfec, _efef: _gcdg, _eadgc: 0, _feef: 1.0, _ebfa: PositionRelative}, nil
}

// GeneratePageBlocks generates a page break block.
func (_fcedb *PageBreak) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gfbgf := []*Block{NewBlock(ctx.PageWidth, ctx.PageHeight-ctx.Y), NewBlock(ctx.PageWidth, ctx.PageHeight)}
	ctx.Page++
	_baaa := ctx
	_baaa.Y = ctx.Margins.Top
	_baaa.X = ctx.Margins.Left
	_baaa.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
	_baaa.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right
	ctx = _baaa
	return _gfbgf, ctx, nil
}

// SetNotes sets the notes section of the invoice.
func (_edbc *Invoice) SetNotes(title, content string) { _edbc._ecf = [2]string{title, content} }
func (_dgade *templateProcessor) addNodeText(_dfba *templateNode, _fdbea string) error {
	_eedd := _dfba._caacd
	if _eedd == nil {
		return nil
	}
	switch _bfdfe := _eedd.(type) {
	case *TextChunk:
		_bfdfe.Text = _fdbea
	case *Paragraph:
		switch _dfba._gbdee.Name.Local {
		case "\u0063h\u0061p\u0074\u0065\u0072\u002d\u0068\u0065\u0061\u0064\u0069\u006e\u0067":
			if _dfba._fbcg != nil {
				if _cega, _abge := _dfba._fbcg._caacd.(*Chapter); _abge {
					_cega._gbcb = _fdbea
					_bfdfe.SetText(_cega.headingText())
				}
			}
		default:
			_bfdfe.SetText(_fdbea)
		}
	}
	return nil
}

// SetColorTop sets border color for top.
func (_cfb *border) SetColorTop(col Color) { _cfb._feg = col }

const (
	DefaultHorizontalScaling = 100
)

func (_ggeg *templateProcessor) parseInt64Array(_dcgc, _edcc string) []int64 {
	_ca.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020\u0069\u006e\u0074\u0036\u0034\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060%\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _dcgc, _edcc)
	_geefa := _dc.Fields(_edcc)
	_dbfbd := make([]int64, 0, len(_geefa))
	for _, _afbd := range _geefa {
		_egbf, _ := _a.ParseInt(_afbd, 10, 64)
		_dbfbd = append(_dbfbd, _egbf)
	}
	return _dbfbd
}
func (_aacfd *TableCell) width(_egaea []float64, _cfdfaa float64) float64 {
	_bbdb := float64(0.0)
	for _abed := 0; _abed < _aacfd._abfda; _abed++ {
		_bbdb += _egaea[_aacfd._eafbd+_abed-1]
	}
	return _bbdb * _cfdfaa
}
func _ggde(_bdbg string) (*Image, error) {
	_ffde, _ceea := _ed.Open(_bdbg)
	if _ceea != nil {
		return nil, _ceea
	}
	defer _ffde.Close()
	_fbad, _ceea := _ggc.ImageHandling.Read(_ffde)
	if _ceea != nil {
		_ca.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _ceea)
		return nil, _ceea
	}
	return _ggba(_fbad)
}
func (_ecea *Invoice) generateHeaderBlocks(_edda DrawContext) ([]*Block, DrawContext, error) {
	_dgcef := _egdc(_ecea._eefc)
	_dgcef.SetEnableWrap(true)
	_dgcef.Append(_ecea._bfbbd)
	_aaed := _gdec(2)
	if _ecea._bfge != nil {
		_egece := _aaed.NewCell()
		_egece.SetHorizontalAlignment(CellHorizontalAlignmentLeft)
		_egece.SetVerticalAlignment(CellVerticalAlignmentMiddle)
		_egece.SetIndent(0)
		_egece.SetContent(_ecea._bfge)
		_ecea._bfge.ScaleToHeight(_dgcef.Height() + 20)
	} else {
		_aaed.SkipCells(1)
	}
	_bgbaf := _aaed.NewCell()
	_bgbaf.SetHorizontalAlignment(CellHorizontalAlignmentRight)
	_bgbaf.SetVerticalAlignment(CellVerticalAlignmentMiddle)
	_bgbaf.SetContent(_dgcef)
	return _aaed.GeneratePageBlocks(_edda)
}
func _bfcfd(_dbdcf *templateProcessor, _gecf *templateNode) (interface{}, error) {
	return _dbdcf.parsePageBreak(_gecf)
}

// SetMargins sets the margins of the chart component.
func (_cdbc *Chart) SetMargins(left, right, top, bottom float64) {
	_cdbc._gcff.Left = left
	_cdbc._gcff.Right = right
	_cdbc._gcff.Top = top
	_cdbc._gcff.Bottom = bottom
}

// GetCoords returns the center coordinates of ellipse (`xc`, `yc`).
func (_cadb *Ellipse) GetCoords() (float64, float64) { return _cadb._eebe, _cadb._beb }

// GetMargins returns the margins of the line: left, right, top, bottom.
func (_eaea *Line) GetMargins() (float64, float64, float64, float64) {
	return _eaea._bccc.Left, _eaea._bccc.Right, _eaea._bccc.Top, _eaea._bccc.Bottom
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
	_bede  []*listItem
	_fed   Margins
	_ebgba TextChunk
	_egab  float64
	_ceeg  bool
	_bgdd  Positioning
	_dagb  TextStyle
}

func (_bbfab *templateProcessor) parseBoolAttr(_fbdff, _dadg string) bool {
	_ca.Log.Debug("P\u0061\u0072\u0073\u0069\u006e\u0067 \u0062\u006f\u006f\u006c\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _fbdff, _dadg)
	_cbad, _ := _a.ParseBool(_dadg)
	return _dadg == "" || _cbad
}
func (_cfcbb *templateProcessor) parseListItem(_dbcad *templateNode) (interface{}, error) {
	if _dbcad._fbcg == nil {
		_cfcbb.nodeLogError(_dbcad, "\u004c\u0069\u0073t\u0020\u0069\u0074\u0065m\u0020\u0070\u0061\u0072\u0065\u006e\u0074 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e")
		return nil, _gfaba
	}
	_abaa, _aebed := _dbcad._fbcg._caacd.(*List)
	if !_aebed {
		_cfcbb.nodeLogError(_dbcad, "\u004c\u0069s\u0074\u0020\u0069\u0074\u0065\u006d\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u004cis\u0074\u002e")
		return nil, _gfaba
	}
	_dbadfb := _fcf()
	_dbadfb._eeed = _abaa._ebgba
	return _dbadfb, nil
}

// GetMargins returns the Paragraph's margins: left, right, top, bottom.
func (_agfbg *Paragraph) GetMargins() (float64, float64, float64, float64) {
	return _agfbg._fgcbf.Left, _agfbg._fgcbf.Right, _agfbg._fgcbf.Top, _agfbg._fgcbf.Bottom
}

// Height returns the height of the division, assuming all components are
// stacked on top of each other.
func (_ecca *Division) Height() float64 {
	var _gegfa float64
	for _, _aced := range _ecca._bfdf {
		switch _efgg := _aced.(type) {
		case marginDrawable:
			_, _, _cggf, _caffc := _efgg.GetMargins()
			_gegfa += _efgg.Height() + _cggf + _caffc
		default:
			_gegfa += _efgg.Height()
		}
	}
	return _gegfa
}

// EnableFontSubsetting enables font subsetting for `font` when the creator output is written to file.
// Embeds only the subset of the runes/glyphs that are actually used to display the file.
// Subsetting can reduce the size of fonts significantly.
func (_eace *Creator) EnableFontSubsetting(font *_ggc.PdfFont) {
	_eace._feff = append(_eace._feff, font)
}

// NewTOCLine creates a new table of contents line with the default style.
func (_bedb *Creator) NewTOCLine(number, title, page string, level uint) *TOCLine {
	return _bbcaf(number, title, page, level, _bedb.NewTextStyle())
}

// NewRadialGradientColor creates a radial gradient color that could act as a color in other componenents.
// Note: The innerRadius must be smaller than outerRadius for the circle to render properly.
func (_dabgb *Creator) NewRadialGradientColor(x float64, y float64, innerRadius float64, outerRadius float64, colorPoints []*ColorPoint) *RadialShading {
	return _afac(x, y, innerRadius, outerRadius, colorPoints)
}

// SetBorderLineStyle sets border style (currently dashed or plain).
func (_ecfc *TableCell) SetBorderLineStyle(style _fc.LineStyle) { _ecfc._fcagf = style }

// Draw processes the specified Drawable widget and generates blocks that can
// be rendered to the output document. The generated blocks can span over one
// or more pages. Additional pages are added if the contents go over the current
// page. Each generated block is assigned to the creator page it will be
// rendered to. In order to render the generated blocks to the creator pages,
// call Finalize, Write or WriteToFile.
func (_bcdg *Creator) Draw(d Drawable) error {
	if _bcdg.getActivePage() == nil {
		_bcdg.NewPage()
	}
	_dcbdb, _dcgb, _dcgf := d.GeneratePageBlocks(_bcdg._eacd)
	if _dcgf != nil {
		return _dcgf
	}
	if len(_dcgb._bcbc) > 0 {
		_bcdg.Errors = append(_bcdg.Errors, _dcgb._bcbc...)
	}
	for _ffge, _fgf := range _dcbdb {
		if _ffge > 0 {
			_bcdg.NewPage()
		}
		_cbfb := _bcdg.getActivePage()
		if _aceb, _baf := _bcdg._gfab[_cbfb]; _baf {
			if _egfe := _aceb.mergeBlocks(_fgf); _egfe != nil {
				return _egfe
			}
			if _fbb := _ega(_fgf._ge, _aceb._ge); _fbb != nil {
				return _fbb
			}
		} else {
			_bcdg._gfab[_cbfb] = _fgf
		}
	}
	_bcdg._eacd.X = _dcgb.X
	_bcdg._eacd.Y = _dcgb.Y
	_bcdg._eacd.Height = _dcgb.PageHeight - _dcgb.Y - _dcgb.Margins.Bottom
	return nil
}
func (_gbge *templateProcessor) parseCellBorderStyleAttr(_eafg, _adceg string) CellBorderStyle {
	_ca.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020c\u0065\u006c\u006c b\u006f\u0072\u0064\u0065\u0072\u0020s\u0074\u0079\u006c\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a \u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025s\u0029\u002e", _eafg, _adceg)
	_dbabg := map[string]CellBorderStyle{"\u006e\u006f\u006e\u0065": CellBorderStyleNone, "\u0073\u0069\u006e\u0067\u006c\u0065": CellBorderStyleSingle, "\u0064\u006f\u0075\u0062\u006c\u0065": CellBorderStyleDouble}[_adceg]
	return _dbabg
}
func (_babcg *Invoice) newColumn(_ebdb string, _fgbg CellHorizontalAlignment) *InvoiceCell {
	_edae := &InvoiceCell{_babcg._dece, _ebdb}
	_edae.Alignment = _fgbg
	return _edae
}
func _dfgcg(_cafcff, _baefd, _fafa, _bgcdc float64) *Line {
	return &Line{_daed: _cafcff, _gfaa: _baefd, _eddg: _fafa, _eabb: _bgcdc, _cdgfa: ColorBlack, _ggbab: 1.0, _adf: 1.0, _fgbfg: []int64{1, 1}, _fgef: PositionAbsolute}
}
func (_gfcb *templateProcessor) parseEllipse(_efegeb *templateNode) (interface{}, error) {
	_daeddd := _gfcb.creator.NewEllipse(0, 0, 0, 0)
	for _, _afff := range _efegeb._gbdee.Attr {
		_fbbce := _afff.Value
		switch _eacg := _afff.Name.Local; _eacg {
		case "\u0063\u0078":
			_daeddd._eebe = _gfcb.parseFloatAttr(_eacg, _fbbce)
		case "\u0063\u0079":
			_daeddd._beb = _gfcb.parseFloatAttr(_eacg, _fbbce)
		case "\u0077\u0069\u0064t\u0068":
			_daeddd.SetWidth(_gfcb.parseFloatAttr(_eacg, _fbbce))
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_daeddd.SetHeight(_gfcb.parseFloatAttr(_eacg, _fbbce))
		case "\u0066\u0069\u006c\u006c\u002d\u0063\u006f\u006c\u006f\u0072":
			_daeddd.SetFillColor(_gfcb.parseColorAttr(_eacg, _fbbce))
		case "\u0066\u0069\u006cl\u002d\u006f\u0070\u0061\u0063\u0069\u0074\u0079":
			_daeddd.SetFillOpacity(_gfcb.parseFloatAttr(_eacg, _fbbce))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072":
			_daeddd.SetBorderColor(_gfcb.parseColorAttr(_eacg, _fbbce))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u006f\u0070a\u0063\u0069\u0074\u0079":
			_daeddd.SetBorderOpacity(_gfcb.parseFloatAttr(_eacg, _fbbce))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068":
			_daeddd.SetBorderWidth(_gfcb.parseFloatAttr(_eacg, _fbbce))
		case "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e":
			_daeddd.SetPositioning(_gfcb.parsePositioningAttr(_eacg, _fbbce))
		case "\u0066\u0069\u0074\u002d\u006d\u006f\u0064\u0065":
			_daeddd.SetFitMode(_gfcb.parseFitModeAttr(_eacg, _fbbce))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_bbeeb := _gfcb.parseMarginAttr(_eacg, _fbbce)
			_daeddd.SetMargins(_bbeeb.Left, _bbeeb.Right, _bbeeb.Top, _bbeeb.Bottom)
		default:
			_gfcb.nodeLogDebug(_efegeb, "\u0055\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u006c\u0069\u0070\u0073\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _eacg)
		}
	}
	return _daeddd, nil
}

// Invoice represents a configurable invoice template.
type Invoice struct {
	_bfbbd string
	_bfge  *Image
	_dcfc  *InvoiceAddress
	_cfbbb *InvoiceAddress
	_cedd  string
	_geeb  [2]*InvoiceCell
	_ddbe  [2]*InvoiceCell
	_fdfc  [2]*InvoiceCell
	_cbga  [][2]*InvoiceCell
	_eag   []*InvoiceCell
	_fbac  [][]*InvoiceCell
	_befc  [2]*InvoiceCell
	_ada   [2]*InvoiceCell
	_edb   [][2]*InvoiceCell
	_ecf   [2]string
	_aacd  [2]string
	_agdc  [][2]string
	_bdcd  TextStyle
	_febc  TextStyle
	_eefc  TextStyle
	_acecg TextStyle
	_debc  TextStyle
	_dcfab TextStyle
	_daag  TextStyle
	_fcba  InvoiceCellProps
	_dece  InvoiceCellProps
	_degfg InvoiceCellProps
	_ddde  InvoiceCellProps
	_edeee Positioning
}

// SetTextAlignment sets the horizontal alignment of the text within the space provided.
func (_dbge *StyledParagraph) SetTextAlignment(align TextAlignment) { _dbge._eaeg = align }
func _cdae(_eadb, _gdbc, _cdaa, _ebgb, _cgfb, _ccbed float64) *Curve {
	_fdea := &Curve{}
	_fdea._ebbc = _eadb
	_fdea._gbae = _gdbc
	_fdea._cfac = _cdaa
	_fdea._cfdb = _ebgb
	_fdea._bceb = _cgfb
	_fdea._bfea = _ccbed
	_fdea._ccad = ColorBlack
	_fdea._bceaa = 1.0
	return _fdea
}

// NewStyledTOCLine creates a new table of contents line with the provided style.
func (_cbc *Creator) NewStyledTOCLine(number, title, page TextChunk, level uint, style TextStyle) *TOCLine {
	return _bebddd(number, title, page, level, style)
}
func (_bace *Invoice) drawInformation() *Table {
	_dgag := _gdec(2)
	_fgga := append([][2]*InvoiceCell{_bace._geeb, _bace._ddbe, _bace._fdfc}, _bace._cbga...)
	for _, _ccde := range _fgga {
		_gagg, _gaaaf := _ccde[0], _ccde[1]
		if _gaaaf.Value == "" {
			continue
		}
		_cfgc := _dgag.NewCell()
		_cfgc.SetBackgroundColor(_gagg.BackgroundColor)
		_bace.setCellBorder(_cfgc, _gagg)
		_fggf := _egdc(_gagg.TextStyle)
		_fggf.Append(_gagg.Value)
		_fggf.SetMargins(0, 0, 2, 1)
		_cfgc.SetContent(_fggf)
		_cfgc = _dgag.NewCell()
		_cfgc.SetBackgroundColor(_gaaaf.BackgroundColor)
		_bace.setCellBorder(_cfgc, _gaaaf)
		_fggf = _egdc(_gaaaf.TextStyle)
		_fggf.Append(_gaaaf.Value)
		_fggf.SetMargins(0, 0, 2, 1)
		_cfgc.SetContent(_fggf)
	}
	return _dgag
}

// ColorCMYKFromArithmetic creates a Color from arithmetic color values (0-1).
// Example:
//
//	green := ColorCMYKFromArithmetic(1.0, 0.0, 1.0, 0.0)
func ColorCMYKFromArithmetic(c, m, y, k float64) Color {
	return cmykColor{_badc: _b.Max(_b.Min(c, 1.0), 0.0), _dea: _b.Max(_b.Min(m, 1.0), 0.0), _gaa: _b.Max(_b.Min(y, 1.0), 0.0), _ccb: _b.Max(_b.Min(k, 1.0), 0.0)}
}

// GeneratePageBlocks draws the polyline on a new block representing the page.
// Implements the Drawable interface.
func (_dcbda *Polyline) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_bfag := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_faga, _fabf := _bfag.setOpacity(_dcbda._afcb, _dcbda._afcb)
	if _fabf != nil {
		return nil, ctx, _fabf
	}
	_dgcgb := _dcbda._edbcg.Points
	for _gdcb := range _dgcgb {
		_ebbe := &_dgcgb[_gdcb]
		_ebbe.Y = ctx.PageHeight - _ebbe.Y
	}
	_dbab, _, _fabf := _dcbda._edbcg.Draw(_faga)
	if _fabf != nil {
		return nil, ctx, _fabf
	}
	if _fabf = _bfag.addContentsByString(string(_dbab)); _fabf != nil {
		return nil, ctx, _fabf
	}
	return []*Block{_bfag}, ctx, nil
}

// SetBackgroundColor set background color of the shading area.
//
// By default the background color is set to white.
func (_gfbc *LinearShading) SetBackgroundColor(backgroundColor Color) {
	_gfbc._aecc.SetBackgroundColor(backgroundColor)
}

// SetTextOverflow controls the behavior of paragraph text which
// does not fit in the available space.
func (_fdag *StyledParagraph) SetTextOverflow(textOverflow TextOverflow) { _fdag._cfec = textOverflow }

// SetMargins sets the Paragraph's margins.
func (_efdca *Paragraph) SetMargins(left, right, top, bottom float64) {
	_efdca._fgcbf.Left = left
	_efdca._fgcbf.Right = right
	_efdca._fgcbf.Top = top
	_efdca._fgcbf.Bottom = bottom
}
func _ega(_gad, _feb *_ggc.PdfPageResources) error {
	_efgc, _ := _gad.GetColorspaces()
	if _efgc != nil && len(_efgc.Colorspaces) > 0 {
		for _acg, _edab := range _efgc.Colorspaces {
			_eca := *_fe.MakeName(_acg)
			if _feb.HasColorspaceByName(_eca) {
				continue
			}
			_bdbf := _feb.SetColorspaceByName(_eca, _edab)
			if _bdbf != nil {
				return _bdbf
			}
		}
	}
	return nil
}
func _dbaec(_fdgc *templateProcessor, _bbgacg *templateNode) (interface{}, error) {
	return _fdgc.parseEllipse(_bbgacg)
}
func (_cgabd *templateProcessor) parseFloatArray(_fcedbd, _gdefg string) []float64 {
	_ca.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020\u0066\u006c\u006f\u0061\u0074\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060%\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _fcedbd, _gdefg)
	_fdebe := _dc.Fields(_gdefg)
	_ddad := make([]float64, 0, len(_fdebe))
	for _, _fbbfa := range _fdebe {
		_afgba, _ := _a.ParseFloat(_fbbfa, 64)
		_ddad = append(_ddad, _afgba)
	}
	return _ddad
}

// GraphicSVG represents a drawable graphic SVG.
// It is used to render the graphic SVG components using a creator instance.
type GraphicSVG struct {
	_dgaf  *_cc.GraphicSVG
	_dgdb  Positioning
	_bgag  float64
	_gcfge float64
	_eaff  Margins
}

// FilledCurve represents a closed path of Bezier curves with a border and fill.
type FilledCurve struct {
	_babe         []_fc.CubicBezierCurve
	FillEnabled   bool
	_ddddf        Color
	BorderEnabled bool
	BorderWidth   float64
	_eadf         Color
}

func _ddf(_bcb *_bdb.ContentStreamOperations, _cba *_ggc.PdfPageResources, _ggdb *_bdb.ContentStreamOperations, _gdcf *_ggc.PdfPageResources) error {
	_eebb := map[_fe.PdfObjectName]_fe.PdfObjectName{}
	_cca := map[_fe.PdfObjectName]_fe.PdfObjectName{}
	_bced := map[_fe.PdfObjectName]_fe.PdfObjectName{}
	_cbe := map[_fe.PdfObjectName]_fe.PdfObjectName{}
	_efa := map[_fe.PdfObjectName]_fe.PdfObjectName{}
	_bcgc := map[_fe.PdfObjectName]_fe.PdfObjectName{}
	for _, _dcg := range *_ggdb {
		switch _dcg.Operand {
		case "\u0044\u006f":
			if len(_dcg.Params) == 1 {
				if _edd, _dee := _dcg.Params[0].(*_fe.PdfObjectName); _dee {
					if _, _ged := _eebb[*_edd]; !_ged {
						var _dfa _fe.PdfObjectName
						_dcbd, _ := _gdcf.GetXObjectByName(*_edd)
						if _dcbd != nil {
							_dfa = *_edd
							for {
								_ad, _ := _cba.GetXObjectByName(_dfa)
								if _ad == nil || _ad == _dcbd {
									break
								}
								_dfa = *_fe.MakeName(_daf(_dfa.String()))
							}
						}
						_cba.SetXObjectByName(_dfa, _dcbd)
						_eebb[*_edd] = _dfa
					}
					_ecg := _eebb[*_edd]
					_dcg.Params[0] = &_ecg
				}
			}
		case "\u0054\u0066":
			if len(_dcg.Params) == 2 {
				if _dde, _eab := _dcg.Params[0].(*_fe.PdfObjectName); _eab {
					if _, _gfg := _cca[*_dde]; !_gfg {
						_dbc, _ffg := _gdcf.GetFontByName(*_dde)
						_debg := *_dde
						if _ffg && _dbc != nil {
							_debg = _ffa(_dde.String(), _dbc, _cba)
						}
						_cba.SetFontByName(_debg, _dbc)
						_cca[*_dde] = _debg
					}
					_bddf := _cca[*_dde]
					_dcg.Params[0] = &_bddf
				}
			}
		case "\u0043\u0053", "\u0063\u0073":
			if len(_dcg.Params) == 1 {
				if _cfa, _fcdf := _dcg.Params[0].(*_fe.PdfObjectName); _fcdf {
					if _, _cffa := _bced[*_cfa]; !_cffa {
						var _abg _fe.PdfObjectName
						_af, _gdfc := _gdcf.GetColorspaceByName(*_cfa)
						if _gdfc {
							_abg = *_cfa
							for {
								_befa, _ebe := _cba.GetColorspaceByName(_abg)
								if !_ebe || _af == _befa {
									break
								}
								_abg = *_fe.MakeName(_daf(_abg.String()))
							}
							_cba.SetColorspaceByName(_abg, _af)
							_bced[*_cfa] = _abg
						} else {
							_ca.Log.Debug("C\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064")
						}
					}
					if _egd, _fac := _bced[*_cfa]; _fac {
						_dcg.Params[0] = &_egd
					} else {
						_ca.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0043\u006f\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", *_cfa)
					}
				}
			}
		case "\u0053\u0043\u004e", "\u0073\u0063\u006e":
			if len(_dcg.Params) == 1 {
				if _eee, _gfc := _dcg.Params[0].(*_fe.PdfObjectName); _gfc {
					if _, _bcaf := _cbe[*_eee]; !_bcaf {
						var _fff _fe.PdfObjectName
						_gef, _eeee := _gdcf.GetPatternByName(*_eee)
						if _eeee {
							_fff = *_eee
							for {
								_eed, _bee := _cba.GetPatternByName(_fff)
								if !_bee || _eed == _gef {
									break
								}
								_fff = *_fe.MakeName(_daf(_fff.String()))
							}
							_ebd := _cba.SetPatternByName(_fff, _gef.ToPdfObject())
							if _ebd != nil {
								return _ebd
							}
							_cbe[*_eee] = _fff
						}
					}
					if _abc, _dab := _cbe[*_eee]; _dab {
						_dcg.Params[0] = &_abc
					}
				}
			}
		case "\u0073\u0068":
			if len(_dcg.Params) == 1 {
				if _fae, _ecc := _dcg.Params[0].(*_fe.PdfObjectName); _ecc {
					if _, _dgbc := _efa[*_fae]; !_dgbc {
						var _cfaa _fe.PdfObjectName
						_fgd, _cadaf := _gdcf.GetShadingByName(*_fae)
						if _cadaf {
							_cfaa = *_fae
							for {
								_ac, _acd := _cba.GetShadingByName(_cfaa)
								if !_acd || _fgd == _ac {
									break
								}
								_cfaa = *_fe.MakeName(_daf(_cfaa.String()))
							}
							_fbg := _cba.SetShadingByName(_cfaa, _fgd.ToPdfObject())
							if _fbg != nil {
								_ca.Log.Debug("E\u0052\u0052\u004f\u0052 S\u0065t\u0020\u0073\u0068\u0061\u0064i\u006e\u0067\u003a\u0020\u0025\u0076", _fbg)
								return _fbg
							}
							_efa[*_fae] = _cfaa
						} else {
							_ca.Log.Debug("\u0053\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
						}
					}
					if _fdb, _ead := _efa[*_fae]; _ead {
						_dcg.Params[0] = &_fdb
					} else {
						_ca.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020S\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0025\u0073 \u006e\u006f\u0074 \u0066o\u0075\u006e\u0064", *_fae)
					}
				}
			}
		case "\u0067\u0073":
			if len(_dcg.Params) == 1 {
				if _edg, _acb := _dcg.Params[0].(*_fe.PdfObjectName); _acb {
					if _, _baa := _bcgc[*_edg]; !_baa {
						var _bed _fe.PdfObjectName
						_efec, _ace := _gdcf.GetExtGState(*_edg)
						if _ace {
							_bed = *_edg
							for {
								_efg, _fec := _cba.GetExtGState(_bed)
								if !_fec || _efec == _efg {
									break
								}
								_bed = *_fe.MakeName(_daf(_bed.String()))
							}
						}
						_cba.AddExtGState(_bed, _efec)
						_bcgc[*_edg] = _bed
					}
					_fca := _bcgc[*_edg]
					_dcg.Params[0] = &_fca
				}
			}
		}
		*_bcb = append(*_bcb, _dcg)
	}
	return nil
}
func _dedafb(_edfd *templateProcessor, _adbb *templateNode) (interface{}, error) {
	return _edfd.parseTable(_adbb)
}

// SetMargins sets the margins of the ellipse.
// NOTE: ellipse margins are only applied if relative positioning is used.
func (_abcgc *Ellipse) SetMargins(left, right, top, bottom float64) {
	_abcgc._cadg.Left = left
	_abcgc._cadg.Right = right
	_abcgc._cadg.Top = top
	_abcgc._cadg.Bottom = bottom
}
func _dfede(_cbbec string, _eeeaa, _ccda TextStyle) *TOC {
	_gbee := _ccda
	_gbee.FontSize = 14
	_dacce := _egdc(_gbee)
	_dacce.SetEnableWrap(true)
	_dacce.SetTextAlignment(TextAlignmentLeft)
	_dacce.SetMargins(0, 0, 0, 5)
	_dabff := _dacce.Append(_cbbec)
	_dabff.Style = _gbee
	return &TOC{_ecagg: _dacce, _fbeeg: []*TOCLine{}, _febed: _eeeaa, _cbgbd: _eeeaa, _effdg: _eeeaa, _abgad: _eeeaa, _feaba: "\u002e", _ebgfa: 10, _dfbae: Margins{0, 0, 2, 2}, _fcffe: PositionRelative, _cceeb: _eeeaa, _dbaf: true}
}

// Curve represents a cubic Bezier curve with a control point.
type Curve struct {
	_ebbc  float64
	_gbae  float64
	_cfac  float64
	_cfdb  float64
	_bceb  float64
	_bfea  float64
	_ccad  Color
	_bceaa float64
}

// Add appends a new item to the list.
// The supported components are: *Paragraph, *StyledParagraph, *Division, *Image, *Table, and *List.
// Returns the marker used for the newly added item. The returned marker
// object can be used to change the text and style of the marker for the
// current item.
func (_afebf *List) Add(item VectorDrawable) (*TextChunk, error) {
	_fecd := &listItem{_gdceb: item, _eeed: _afebf._ebgba}
	switch _dgfa := item.(type) {
	case *Paragraph:
	case *StyledParagraph:
	case *List:
		if _dgfa._ceeg {
			_dgfa._egab = 15
		}
	case *Division:
	case *Image:
	case *Table:
	default:
		return nil, _fa.New("\u0074\u0068i\u0073\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0064\u0072\u0061\u0077\u0061\u0062\u006c\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u0020\u006c\u0069\u0073\u0074")
	}
	_afebf._bede = append(_afebf._bede, _fecd)
	return &_fecd._eeed, nil
}
func _agccg(_fgdda *_ed.File) ([]*_ggc.PdfPage, error) {
	_ecbe, _fcaca := _ggc.NewPdfReader(_fgdda)
	if _fcaca != nil {
		return nil, _fcaca
	}
	_dabgba, _fcaca := _ecbe.GetNumPages()
	if _fcaca != nil {
		return nil, _fcaca
	}
	var _cdcab []*_ggc.PdfPage
	for _dbgg := 0; _dbgg < _dabgba; _dbgg++ {
		_dgcc, _ceebb := _ecbe.GetPage(_dbgg + 1)
		if _ceebb != nil {
			return nil, _ceebb
		}
		_cdcab = append(_cdcab, _dgcc)
	}
	return _cdcab, nil
}

// GetHorizontalAlignment returns the horizontal alignment of the image.
func (_cbdfa *Image) GetHorizontalAlignment() HorizontalAlignment { return _cbdfa._afdag }

// UnsupportedRuneError is an error that occurs when there is unsupported glyph being used.
type UnsupportedRuneError struct {
	Message string
	Rune    rune
}

func (_efcg *StyledParagraph) wrapWordChunks() {
	if !_efcg._cbeg {
		return
	}
	var (
		_bfca  []*TextChunk
		_decca *_ggc.PdfFont
	)
	for _, _dbbc := range _efcg._cdffa {
		_bafe := []rune(_dbbc.Text)
		if _decca == nil {
			_decca = _dbbc.Style.Font
		}
		_dcbf := _dbbc._dfcb
		_afcac := _dbbc.VerticalAlignment
		if len(_bfca) > 0 {
			if len(_bafe) == 1 && _cd.IsPunct(_bafe[0]) && _dbbc.Style.Font == _decca {
				_cadc := []rune(_bfca[len(_bfca)-1].Text)
				_bfca[len(_bfca)-1].Text = string(append(_cadc, _bafe[0]))
				continue
			} else {
				_, _ebedd := _a.Atoi(_dbbc.Text)
				if _ebedd == nil {
					_gbgd := []rune(_bfca[len(_bfca)-1].Text)
					_bdgg := len(_gbgd)
					if _bdgg >= 2 {
						_, _cbafc := _a.Atoi(string(_gbgd[_bdgg-2]))
						if _cbafc == nil && _cd.IsPunct(_gbgd[_bdgg-1]) {
							_bfca[len(_bfca)-1].Text = string(append(_gbgd, _bafe...))
							continue
						}
					}
				}
			}
		}
		_aafad, _bdaa := _cddc(_dbbc.Text)
		if _bdaa != nil {
			_ca.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0062\u0072\u0065\u0061\u006b\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020w\u006f\u0072\u0064\u0073\u003a\u0020\u0025\u0076", _bdaa)
			_aafad = []string{_dbbc.Text}
		}
		for _, _ebfcb := range _aafad {
			_aaga := NewTextChunk(_ebfcb, _dbbc.Style)
			_aaga._dfcb = _bgbee(_dcbf)
			_aaga.VerticalAlignment = _afcac
			_bfca = append(_bfca, _aaga)
		}
		_decca = _dbbc.Style.Font
	}
	if len(_bfca) > 0 {
		_efcg._cdffa = _bfca
	}
}
func _gfcae(_caaa string) *_ggc.PdfAnnotation {
	_gbfc := _ggc.NewPdfAnnotationLink()
	_dgcgc := _ggc.NewBorderStyle()
	_dgcgc.SetBorderWidth(0)
	_gbfc.BS = _dgcgc.ToPdfObject()
	_geege := _ggc.NewPdfActionURI()
	_geege.URI = _fe.MakeString(_caaa)
	_gbfc.SetAction(_geege.PdfAction)
	return _gbfc.PdfAnnotation
}
func _egdc(_fcaf TextStyle) *StyledParagraph {
	return &StyledParagraph{_cdffa: []*TextChunk{}, _egcg: _fcaf, _aecf: _dbeag(_fcaf.Font), _babcf: 1.0, _eaeg: TextAlignmentLeft, _gfbb: true, _facb: true, _cbeg: false, _adag: 0, _bfagd: 1, _bdbc: 1, _gbga: PositionRelative}
}

// NewTextChunk returns a new text chunk instance.
func NewTextChunk(text string, style TextStyle) *TextChunk {
	return &TextChunk{Text: text, Style: style, VerticalAlignment: TextVerticalAlignmentBaseline}
}

// SetSideBorderWidth sets the cell's side border width.
func (_fddg *TableCell) SetSideBorderWidth(side CellBorderSide, width float64) {
	switch side {
	case CellBorderSideAll:
		_fddg._caffa = width
		_fddg._egaeb = width
		_fddg._egbd = width
		_fddg._ddbdf = width
	case CellBorderSideTop:
		_fddg._caffa = width
	case CellBorderSideBottom:
		_fddg._egaeb = width
	case CellBorderSideLeft:
		_fddg._egbd = width
	case CellBorderSideRight:
		_fddg._ddbdf = width
	}
}

// AddPage adds the specified page to the creator.
// NOTE: If the page has a Rotate flag, the creator will take care of
// transforming the contents to maintain the correct orientation.
func (_efc *Creator) AddPage(page *_ggc.PdfPage) error {
	_aabf, _abag := _efc.wrapPageIfNeeded(page)
	if _abag != nil {
		return _abag
	}
	if _aabf != nil {
		page = _aabf
	}
	_cef, _abag := page.GetMediaBox()
	if _abag != nil {
		_ca.Log.Debug("\u0046\u0061\u0069l\u0065\u0064\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0070\u0061g\u0065\u0020\u006d\u0065\u0064\u0069\u0061\u0062\u006f\u0078\u003a\u0020\u0025\u0076", _abag)
		return _abag
	}
	_cef.Normalize()
	_aef, _cfe := _cef.Llx, _cef.Lly
	_aca := _cef
	if _dbfb := page.CropBox; _dbfb != nil && *_dbfb != *_cef {
		_dbfb.Normalize()
		_aef, _cfe = _dbfb.Llx, _dbfb.Lly
		_aca = _dbfb
	}
	_cfba := _bd.IdentityMatrix()
	_cfbb, _abag := page.GetRotate()
	if _abag != nil {
		_ca.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _abag.Error())
	}
	_deea := _cfbb%360 != 0 && _cfbb%90 == 0
	if _deea {
		_gda := float64((360 + _cfbb%360) % 360)
		if _gda == 90 {
			_cfba = _cfba.Translate(_aca.Width(), 0)
		} else if _gda == 180 {
			_cfba = _cfba.Translate(_aca.Width(), _aca.Height())
		} else if _gda == 270 {
			_cfba = _cfba.Translate(0, _aca.Height())
		}
		_cfba = _cfba.Mult(_bd.RotationMatrix(_gda * _b.Pi / 180))
		_cfba = _cfba.Round(0.000001)
		_geff := _gddf(_aca, _cfba)
		_aca = _geff
		_aca.Normalize()
	}
	if _aef != 0 || _cfe != 0 {
		_cfba = _bd.TranslationMatrix(_aef, _cfe).Mult(_cfba)
	}
	if !_cfba.Identity() {
		_cfba = _cfba.Round(0.000001)
		_efc._afbc[page] = &pageTransformations{_dafc: &_cfba}
	}
	_efc._abf = _aca.Width()
	_efc._ffc = _aca.Height()
	_efc.initContext()
	_efc._cec = append(_efc._cec, page)
	_efc._eacd.Page++
	return nil
}

// NewColumn returns a new column for the line items invoice table.
func (_fbgc *Invoice) NewColumn(description string) *InvoiceCell {
	return _fbgc.newColumn(description, CellHorizontalAlignmentLeft)
}
func (_bfba *Table) wrapContent(_bgfba DrawContext) error {
	if _bfba._eebgf {
		return nil
	}
	_bfba.sortCells()
	_edad := func(_fccad *TableCell, _dabfe int, _gafg int, _efcd int) (_bgagg int) {
		if _efcd < 1 {
			return -1
		}
		_gdfcfe := 0
		for _addbf := _gafg + 1; _addbf < len(_bfba._cacca)-1; _addbf++ {
			_egbgd := _bfba._cacca[_addbf]
			if _egbgd._deded == _efcd && _gdfcfe != _gafg {
				_gdfcfe = _addbf
				if (_egbgd._eafbd < _fccad._eafbd && _bfba._aeaa > _egbgd._eafbd) || _fccad._eafbd < _bfba._aeaa {
					continue
				}
				break
			}
		}
		_cbgce := float64(0.0)
		for _affc := 0; _affc < _fccad._ffgdb; _affc++ {
			_cbgce += _bfba._begb[_fccad._deded+_affc-1]
		}
		_deca := _fccad.width(_bfba._abbbf, _bgfba.Width)
		var (
			_eccc  VectorDrawable
			_dafde = false
		)
		switch _begg := _fccad._efbbe.(type) {
		case *StyledParagraph:
			_dbbffe := _bgfba
			_dbbffe.Height = _b.Floor(_cbgce - _begg._fbgbc.Top - _begg._fbgbc.Bottom - 0.5*_begg.getTextHeight())
			_dbbffe.Width = _deca
			_ddcdb, _ceddf, _dbdc := _begg.split(_dbbffe)
			if _dbdc != nil {
				_ca.Log.Error("\u0045\u0072\u0072o\u0072\u0020\u0077\u0072a\u0070\u0020\u0073\u0074\u0079\u006c\u0065d\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u003a\u0020\u0025\u0076", _dbdc.Error())
			}
			if _ddcdb != nil && _ceddf != nil {
				_bfba._cacca[_gafg]._efbbe = _ddcdb
				_eccc = _ceddf
				_dafde = true
			}
		}
		_bfba._cacca[_gafg]._ffgdb = _fccad._ffgdb
		_bgfba.Height = _bgfba.PageHeight - _bgfba.Margins.Top - _bgfba.Margins.Bottom
		_bfff := _fccad.cloneProps(nil)
		if _dafde {
			_bfff._efbbe = _eccc
		}
		_bfff._ffgdb = _dabfe
		_bfff._deded = _efcd + 1
		_bfff._eafbd = _fccad._eafbd
		if _bfff._deded+_bfff._ffgdb-1 > _bfba._fgbfga {
			for _fdaa := _bfba._fgbfga; _fdaa < _bfff._deded+_bfff._ffgdb-1; _fdaa++ {
				_bfba._fgbfga++
				_bfba._begb = append(_bfba._begb, _bfba._ddfaf)
			}
		}
		_bfba._cacca = append(_bfba._cacca[:_gdfcfe+1], append([]*TableCell{_bfff}, _bfba._cacca[_gdfcfe+1:]...)...)
		return _gdfcfe + 1
	}
	_bdcgb := func(_dcaa *TableCell, _gdef int, _efbee int, _adbg float64) (_gddd int) {
		_fcbd := _dcaa.width(_bfba._abbbf, _bgfba.Width)
		_feadc := _adbg
		_fffc := 1
		_agafg := _bgfba.Height
		if _agafg > 0 {
			for _feadc > _agafg {
				_feadc -= _bgfba.Height
				_agafg = _bgfba.PageHeight - _bgfba.Margins.Top - _bgfba.Margins.Bottom
				_fffc++
			}
		}
		var (
			_aefab VectorDrawable
			_fcdcg = false
		)
		switch _faac := _dcaa._efbbe.(type) {
		case *StyledParagraph:
			_dbdg := _bgfba
			_dbdg.Height = _b.Floor(_bgfba.Height - _faac._fbgbc.Top - _faac._fbgbc.Bottom - 0.5*_faac.getTextHeight())
			_dbdg.Width = _fcbd
			_abba, _bfefc, _fbcda := _faac.split(_dbdg)
			if _fbcda != nil {
				_ca.Log.Error("\u0045\u0072\u0072o\u0072\u0020\u0077\u0072a\u0070\u0020\u0073\u0074\u0079\u006c\u0065d\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u003a\u0020\u0025\u0076", _fbcda.Error())
			}
			if _abba != nil && _bfefc != nil {
				_bfba._cacca[_gdef]._efbbe = _abba
				_aefab = _bfefc
				_fcdcg = true
			}
		}
		if _fffc < 2 {
			return -1
		}
		if _bfba._cacca[_gdef]._deded+_fffc-1 > _bfba._fgbfga {
			for _bfafg := 0; _bfafg < _fffc; _bfafg++ {
				_bfba._fgbfga++
				_bfba._begb = append(_bfba._begb, _bfba._ddfaf)
			}
		}
		_aeeg := _adbg / float64(_fffc)
		for _deaa := 0; _deaa < _fffc; _deaa++ {
			_bfba._begb[_efbee+_deaa-1] = _aeeg
		}
		_bgfba.Height = _bgfba.PageHeight - _bgfba.Margins.Top - _bgfba.Margins.Bottom
		_fdagd := _dcaa.cloneProps(nil)
		if _fcdcg {
			_fdagd._efbbe = _aefab
		}
		_fdagd._ffgdb = 1
		_fdagd._deded = _efbee + _fffc - 1
		_fdagd._eafbd = _dcaa._eafbd
		_bfba._cacca = append(_bfba._cacca, _fdagd)
		return len(_bfba._cacca)
	}
	_bbef := 1
	_eccca := -1
	for _debd := 0; _debd < len(_bfba._cacca); _debd++ {
		_cbfe := _bfba._cacca[_debd]
		if _eccca == _debd {
			_bbef = _cbfe._deded
		}
		if _cbfe._ffgdb < 2 {
			if _fabc := _bfba._begb[_cbfe._deded-1]; _fabc > _bgfba.Height {
				_eccca = _bdcgb(_cbfe, _debd, _cbfe._deded, _fabc)
				continue
			}
			continue
		}
		_eegf := float64(0)
		for _dgaff := 0; _dgaff < _cbfe._ffgdb; _dgaff++ {
			_eegf += _bfba._begb[_cbfe._deded+_dgaff-1]
		}
		_aebbb := float64(0)
		for _dffgd := _bbef - 1; _dffgd < _cbfe._deded-1; _dffgd++ {
			_aebbb += _bfba._begb[_dffgd]
		}
		if _eegf <= (_bgfba.Height - _aebbb) {
			continue
		}
		_aebbc := float64(0.0)
		_cfef := _cbfe._ffgdb
		_ededg := -1
		_fedb := 1
		for _aebbe := 1; _aebbe <= _cbfe._ffgdb; _aebbe++ {
			if (_aebbc + _bfba._begb[_cbfe._deded+_aebbe-2]) > (_bgfba.Height - _aebbb) {
				_fedb--
				break
			}
			_ededg = _cbfe._deded + _aebbe - 1
			_cfef = _cbfe._ffgdb - _aebbe
			_aebbc += _bfba._begb[_cbfe._deded+_aebbe-2]
			_fedb++
		}
		if _cbfe._ffgdb == _cfef {
			_bgfba.Height = _bgfba.PageHeight - _bgfba.Margins.Top - _bgfba.Margins.Bottom
			_bbef = _cbfe._deded
			_debd--
			continue
		}
		if _cfef > 0 && _cbfe._ffgdb > _fedb {
			_cbfe._ffgdb = _fedb
			_eccca = _edad(_cbfe, _cfef, _debd, _ededg)
			if _debd+1 == _eccca {
				_debd--
			}
		}
		_bbef = _cbfe._deded
	}
	_bfba.sortCells()
	return nil
}

// FitMode returns the fit mode of the line.
func (_bfgd *Line) FitMode() FitMode { return _bfgd._afad }

// Positioning returns the type of positioning the ellipse is set to use.
func (_cdgaf *Ellipse) Positioning() Positioning { return _cdgaf._acgd }

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

// GetOptimizer returns current PDF optimizer.
func (_aebfe *Creator) GetOptimizer() _ggc.Optimizer { return _aebfe._eecf }

// Cols returns the total number of columns the table has.
func (_gaff *Table) Cols() int { return _gaff._aeaa }
func (_bbg *pageTransformations) transformPage(_ggge *_ggc.PdfPage) error {
	if _bbd := _bbg.applyFlip(_ggge); _bbd != nil {
		return _bbd
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
	_dgce  rune
	_bcbc  []error
}

// SetMargins sets the Block's left, right, top, bottom, margins.
func (_cg *Block) SetMargins(left, right, top, bottom float64) {
	_cg._bc.Left = left
	_cg._bc.Right = right
	_cg._bc.Top = top
	_cg._bc.Bottom = bottom
}

// SetBackgroundColor set background color of the shading area.
//
// By default the background color is set to white.
func (_acebb *RadialShading) SetBackgroundColor(backgroundColor Color) {
	_acebb._fcacc.SetBackgroundColor(backgroundColor)
}

// MoveTo moves the drawing context to absolute coordinates (x, y).
func (_gbde *Creator) MoveTo(x, y float64) { _gbde._eacd.X = x; _gbde._eacd.Y = y }

type shading struct {
	_dedad Color
	_daff  bool
	_fgdd  []bool
	_ggcd  []*ColorPoint
}

// AddPatternResource adds pattern dictionary inside the resources dictionary.
func (_debcf *RadialShading) AddPatternResource(block *Block) (_bfed _fe.PdfObjectName, _gdfff error) {
	_cbce := 1
	_eaee := _fe.PdfObjectName("\u0050" + _a.Itoa(_cbce))
	for block._ge.HasPatternByName(_eaee) {
		_cbce++
		_eaee = _fe.PdfObjectName("\u0050" + _a.Itoa(_cbce))
	}
	if _ebbcd := block._ge.SetPatternByName(_eaee, _debcf.ToPdfShadingPattern().ToPdfObject()); _ebbcd != nil {
		return "", _ebbcd
	}
	return _eaee, nil
}

// SetBorderOpacity sets the border opacity.
func (_ggfce *PolyBezierCurve) SetBorderOpacity(opacity float64) { _ggfce._gfcgc = opacity }
func (_bgeda *Rectangle) applyFitMode(_cegee float64) {
	_cegee -= _bgeda._ebfce.Left + _bgeda._ebfce.Right + _bgeda._decec
	switch _bgeda._gaef {
	case FitModeFillWidth:
		_bgeda.ScaleToWidth(_cegee)
	}
}

// AnchorPoint defines anchor point where the center position of the radial gradient would be calculated.
type AnchorPoint int

func _egedc(_dcbfb *_ggc.PdfAnnotationLink) *_ggc.PdfAnnotationLink {
	if _dcbfb == nil {
		return nil
	}
	_eacdeb := _ggc.NewPdfAnnotationLink()
	_eacdeb.BS = _dcbfb.BS
	_eacdeb.A = _dcbfb.A
	if _fccbc, _bgee := _dcbfb.GetAction(); _bgee == nil && _fccbc != nil {
		_eacdeb.SetAction(_fccbc)
	}
	if _eadbg, _bbafe := _dcbfb.Dest.(*_fe.PdfObjectArray); _bbafe {
		_eacdeb.Dest = _fe.MakeArray(_eadbg.Elements()...)
	}
	return _eacdeb
}

// EnableRowWrap controls whether rows are wrapped across pages.
// NOTE: Currently, row wrapping is supported for rows using StyledParagraphs.
func (_agec *Table) EnableRowWrap(enable bool) { _agec._eebgf = enable }

// SetHeight sets the height of the ellipse.
func (_fbdf *Ellipse) SetHeight(height float64) { _fbdf._dabc = height }

// SetBorder sets the cell's border style.
func (_cgfc *TableCell) SetBorder(side CellBorderSide, style CellBorderStyle, width float64) {
	if style == CellBorderStyleSingle && side == CellBorderSideAll {
		_cgfc._ffdeb = CellBorderStyleSingle
		_cgfc._egbd = width
		_cgfc._cbafe = CellBorderStyleSingle
		_cgfc._egaeb = width
		_cgfc._dfbfd = CellBorderStyleSingle
		_cgfc._ddbdf = width
		_cgfc._aaddc = CellBorderStyleSingle
		_cgfc._caffa = width
	} else if style == CellBorderStyleDouble && side == CellBorderSideAll {
		_cgfc._ffdeb = CellBorderStyleDouble
		_cgfc._egbd = width
		_cgfc._cbafe = CellBorderStyleDouble
		_cgfc._egaeb = width
		_cgfc._dfbfd = CellBorderStyleDouble
		_cgfc._ddbdf = width
		_cgfc._aaddc = CellBorderStyleDouble
		_cgfc._caffa = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideLeft {
		_cgfc._ffdeb = style
		_cgfc._egbd = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideBottom {
		_cgfc._cbafe = style
		_cgfc._egaeb = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideRight {
		_cgfc._dfbfd = style
		_cgfc._ddbdf = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideTop {
		_cgfc._aaddc = style
		_cgfc._caffa = width
	}
}
func (_cfga *templateProcessor) parseTableCell(_fdbd *templateNode) (interface{}, error) {
	if _fdbd._fbcg == nil {
		_cfga.nodeLogError(_fdbd, "\u0054\u0061\u0062\u006c\u0065\u0020\u0063\u0065\u006c\u006c\u0020\u0070\u0061\u0072\u0065n\u0074 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e")
		return nil, _gfaba
	}
	_agdbb, _baaab := _fdbd._fbcg._caacd.(*Table)
	if !_baaab {
		_cfga.nodeLogError(_fdbd, "\u0054\u0061\u0062\u006c\u0065\u0020\u0063\u0065\u006c\u006c\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0028\u0025\u0054\u0029\u0020\u0069s\u0020\u006e\u006f\u0074\u0020a\u0020\u0074a\u0062\u006c\u0065\u002e", _fdbd._fbcg._caacd)
		return nil, _gfaba
	}
	var _bcce, _gfcag int64
	for _, _ceace := range _fdbd._gbdee.Attr {
		_bfeab := _ceace.Value
		switch _cbccf := _ceace.Name.Local; _cbccf {
		case "\u0063o\u006c\u0073\u0070\u0061\u006e":
			_bcce = _cfga.parseInt64Attr(_cbccf, _bfeab)
		case "\u0072o\u0077\u0073\u0070\u0061\u006e":
			_gfcag = _cfga.parseInt64Attr(_cbccf, _bfeab)
		}
	}
	if _bcce <= 0 {
		_bcce = 1
	}
	if _gfcag <= 0 {
		_gfcag = 1
	}
	_cdfb := _agdbb.MultiCell(int(_gfcag), int(_bcce))
	for _, _beaac := range _fdbd._gbdee.Attr {
		_adae := _beaac.Value
		switch _edabc := _beaac.Name.Local; _edabc {
		case "\u0069\u006e\u0064\u0065\u006e\u0074":
			_cdfb.SetIndent(_cfga.parseFloatAttr(_edabc, _adae))
		case "\u0061\u006c\u0069g\u006e":
			_cdfb.SetHorizontalAlignment(_cfga.parseCellAlignmentAttr(_edabc, _adae))
		case "\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c\u002da\u006c\u0069\u0067\u006e":
			_cdfb.SetVerticalAlignment(_cfga.parseCellVerticalAlignmentAttr(_edabc, _adae))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0073\u0074\u0079\u006c\u0065":
			_cdfb.SetSideBorderStyle(CellBorderSideAll, _cfga.parseCellBorderStyleAttr(_edabc, _adae))
		case "\u0062\u006fr\u0064\u0065\u0072-\u0073\u0074\u0079\u006c\u0065\u002d\u0074\u006f\u0070":
			_cdfb.SetSideBorderStyle(CellBorderSideTop, _cfga.parseCellBorderStyleAttr(_edabc, _adae))
		case "\u0062\u006f\u0072\u0064er\u002d\u0073\u0074\u0079\u006c\u0065\u002d\u0062\u006f\u0074\u0074\u006f\u006d":
			_cdfb.SetSideBorderStyle(CellBorderSideBottom, _cfga.parseCellBorderStyleAttr(_edabc, _adae))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0073\u0074\u0079\u006c\u0065-\u006c\u0065\u0066\u0074":
			_cdfb.SetSideBorderStyle(CellBorderSideLeft, _cfga.parseCellBorderStyleAttr(_edabc, _adae))
		case "\u0062o\u0072d\u0065\u0072\u002d\u0073\u0074y\u006c\u0065-\u0072\u0069\u0067\u0068\u0074":
			_cdfb.SetSideBorderStyle(CellBorderSideRight, _cfga.parseCellBorderStyleAttr(_edabc, _adae))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068":
			_cdfb.SetSideBorderWidth(CellBorderSideAll, _cfga.parseFloatAttr(_edabc, _adae))
		case "\u0062\u006fr\u0064\u0065\u0072-\u0077\u0069\u0064\u0074\u0068\u002d\u0074\u006f\u0070":
			_cdfb.SetSideBorderWidth(CellBorderSideTop, _cfga.parseFloatAttr(_edabc, _adae))
		case "\u0062\u006f\u0072\u0064er\u002d\u0077\u0069\u0064\u0074\u0068\u002d\u0062\u006f\u0074\u0074\u006f\u006d":
			_cdfb.SetSideBorderWidth(CellBorderSideBottom, _cfga.parseFloatAttr(_edabc, _adae))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068-\u006c\u0065\u0066\u0074":
			_cdfb.SetSideBorderWidth(CellBorderSideLeft, _cfga.parseFloatAttr(_edabc, _adae))
		case "\u0062o\u0072d\u0065\u0072\u002d\u0077\u0069d\u0074\u0068-\u0072\u0069\u0067\u0068\u0074":
			_cdfb.SetSideBorderWidth(CellBorderSideRight, _cfga.parseFloatAttr(_edabc, _adae))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072":
			_cdfb.SetSideBorderColor(CellBorderSideAll, _cfga.parseColorAttr(_edabc, _adae))
		case "\u0062\u006fr\u0064\u0065\u0072-\u0063\u006f\u006c\u006f\u0072\u002d\u0074\u006f\u0070":
			_cdfb.SetSideBorderColor(CellBorderSideTop, _cfga.parseColorAttr(_edabc, _adae))
		case "\u0062\u006f\u0072\u0064er\u002d\u0063\u006f\u006c\u006f\u0072\u002d\u0062\u006f\u0074\u0074\u006f\u006d":
			_cdfb.SetSideBorderColor(CellBorderSideBottom, _cfga.parseColorAttr(_edabc, _adae))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072-\u006c\u0065\u0066\u0074":
			_cdfb.SetSideBorderColor(CellBorderSideLeft, _cfga.parseColorAttr(_edabc, _adae))
		case "\u0062o\u0072d\u0065\u0072\u002d\u0063\u006fl\u006f\u0072-\u0072\u0069\u0067\u0068\u0074":
			_cdfb.SetSideBorderColor(CellBorderSideRight, _cfga.parseColorAttr(_edabc, _adae))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u006c\u0069\u006e\u0065\u002ds\u0074\u0079\u006c\u0065":
			_cdfb.SetBorderLineStyle(_cfga.parseLineStyleAttr(_edabc, _adae))
		case "\u0062\u0061c\u006b\u0067\u0072o\u0075\u006e\u0064\u002d\u0063\u006f\u006c\u006f\u0072":
			_cdfb.SetBackgroundColor(_cfga.parseColorAttr(_edabc, _adae))
		case "\u0063o\u006c\u0073\u0070\u0061\u006e", "\u0072o\u0077\u0073\u0070\u0061\u006e":
		default:
			_cfga.nodeLogDebug(_fdbd, "\u0055\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0063\u0065\u006c\u006c\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e", _edabc)
		}
	}
	return _cdfb, nil
}
func _dbeag(_adege *_ggc.PdfFont) TextStyle {
	return TextStyle{Color: ColorRGBFrom8bit(0, 0, 238), Font: _adege, FontSize: 10, OutlineSize: 1, HorizontalScaling: DefaultHorizontalScaling, UnderlineStyle: TextDecorationLineStyle{Offset: 1, Thickness: 1}}
}

type templateTag struct {
	_bdcdf map[string]struct{}
	_bcddf func(*templateProcessor, *templateNode) (interface{}, error)
}

func (_gac *Chapter) headingText() string {
	_fdce := _gac._gbcb
	if _ddea := _gac.headingNumber(); _ddea != "" {
		_fdce = _df.Sprintf("\u0025\u0073\u0020%\u0073", _ddea, _fdce)
	}
	return _fdce
}
func _cfgb(_bdef *templateProcessor, _eaeac *templateNode) (interface{}, error) {
	return _bdef.parseStyledParagraph(_eaeac)
}

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_caac *Invoice) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_effd := ctx
	_bcbea := []func(_cea DrawContext) ([]*Block, DrawContext, error){_caac.generateHeaderBlocks, _caac.generateInformationBlocks, _caac.generateLineBlocks, _caac.generateTotalBlocks, _caac.generateNoteBlocks}
	var _edfc []*Block
	for _, _bfdb := range _bcbea {
		_gcbe, _eefg, _gdag := _bfdb(ctx)
		if _gdag != nil {
			return _edfc, ctx, _gdag
		}
		if len(_edfc) == 0 {
			_edfc = _gcbe
		} else if len(_gcbe) > 0 {
			_edfc[len(_edfc)-1].mergeBlocks(_gcbe[0])
			_edfc = append(_edfc, _gcbe[1:]...)
		}
		ctx = _eefg
	}
	if _caac._edeee.IsRelative() {
		ctx.X = _effd.X
	}
	if _caac._edeee.IsAbsolute() {
		return _edfc, _effd, nil
	}
	return _edfc, ctx, nil
}
func (_adc *Creator) getActivePage() *_ggc.PdfPage {
	if _adc._bcba == nil {
		if len(_adc._cec) == 0 {
			return nil
		}
		return _adc._cec[len(_adc._cec)-1]
	}
	return _adc._bcba
}
func (_ddagg *templateProcessor) parseAttrPropList(_eebd string) map[string]string {
	_fdgd := _dc.Fields(_eebd)
	if len(_fdgd) == 0 {
		return nil
	}
	_cbac := map[string]string{}
	for _, _bbaed := range _fdgd {
		_geade := _cccd.FindStringSubmatch(_bbaed)
		if len(_geade) < 3 {
			continue
		}
		_dfegg, _baba := _dc.TrimSpace(_geade[1]), _geade[2]
		if _dfegg == "" {
			continue
		}
		_cbac[_dfegg] = _baba
	}
	return _cbac
}

// SetTotal sets the total of the invoice.
func (_bgcd *Invoice) SetTotal(value string) { _bgcd._ada[1].Value = value }
func (_egef *templateProcessor) parseStyledParagraph(_ecaef *templateNode) (interface{}, error) {
	_aadebf := _egef.creator.NewStyledParagraph()
	for _, _ggfd := range _ecaef._gbdee.Attr {
		_dgeec := _ggfd.Value
		switch _cgab := _ggfd.Name.Local; _cgab {
		case "\u0074\u0065\u0078\u0074\u002d\u0061\u006c\u0069\u0067\u006e":
			_aadebf.SetTextAlignment(_egef.parseTextAlignmentAttr(_cgab, _dgeec))
		case "\u0076\u0065\u0072\u0074ic\u0061\u006c\u002d\u0074\u0065\u0078\u0074\u002d\u0061\u006c\u0069\u0067\u006e":
			_aadebf.SetTextVerticalAlignment(_egef.parseTextVerticalAlignmentAttr(_cgab, _dgeec))
		case "l\u0069\u006e\u0065\u002d\u0068\u0065\u0069\u0067\u0068\u0074":
			_aadebf.SetLineHeight(_egef.parseFloatAttr(_cgab, _dgeec))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_ggag := _egef.parseMarginAttr(_cgab, _dgeec)
			_aadebf.SetMargins(_ggag.Left, _ggag.Right, _ggag.Top, _ggag.Bottom)
		case "e\u006e\u0061\u0062\u006c\u0065\u002d\u0077\u0072\u0061\u0070":
			_aadebf.SetEnableWrap(_egef.parseBoolAttr(_cgab, _dgeec))
		case "\u0065\u006ea\u0062\u006c\u0065-\u0077\u006f\u0072\u0064\u002d\u0077\u0072\u0061\u0070":
			_aadebf.EnableWordWrap(_egef.parseBoolAttr(_cgab, _dgeec))
		case "\u0074\u0065\u0078\u0074\u002d\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077":
			_aadebf.SetTextOverflow(_egef.parseTextOverflowAttr(_cgab, _dgeec))
		case "\u0078":
			_aadebf.SetPos(_egef.parseFloatAttr(_cgab, _dgeec), _aadebf._cfge)
		case "\u0079":
			_aadebf.SetPos(_aadebf._gcfda, _egef.parseFloatAttr(_cgab, _dgeec))
		case "\u0061\u006e\u0067l\u0065":
			_aadebf.SetAngle(_egef.parseFloatAttr(_cgab, _dgeec))
		default:
			_egef.nodeLogDebug(_ecaef, "\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0073\u0074\u0079l\u0065\u0064\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0060\u0025\u0073`.\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _cgab)
		}
	}
	return _aadebf, nil
}

// GeneratePageBlocks draws the composite curve polygon on a new block
// representing the page. Implements the Drawable interface.
func (_aadf *CurvePolygon) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_egca := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_agdf, _fefd := _egca.setOpacity(_aadf._gfage, _aadf._fgcc)
	if _fefd != nil {
		return nil, ctx, _fefd
	}
	_dcae := _aadf._ecae
	_dcae.FillEnabled = _dcae.FillColor != nil
	_dcae.BorderEnabled = _dcae.BorderColor != nil && _dcae.BorderWidth > 0
	var (
		_beaa = ctx.PageHeight
		_aee  = _dcae.Rings
		_cdca = make([][]_fc.CubicBezierCurve, 0, len(_dcae.Rings))
	)
	_bbfa := _ggc.PdfRectangle{}
	if len(_aee) > 0 && len(_aee[0]) > 0 {
		_beaaf := _aee[0][0]
		_beaaf.P0.Y = _beaa - _beaaf.P0.Y
		_beaaf.P1.Y = _beaa - _beaaf.P1.Y
		_beaaf.P2.Y = _beaa - _beaaf.P2.Y
		_beaaf.P3.Y = _beaa - _beaaf.P3.Y
		_bbfa = _beaaf.GetBounds()
	}
	for _, _eadg := range _aee {
		_cgcc := make([]_fc.CubicBezierCurve, 0, len(_eadg))
		for _, _eaa := range _eadg {
			_adg := _eaa
			_adg.P0.Y = _beaa - _adg.P0.Y
			_adg.P1.Y = _beaa - _adg.P1.Y
			_adg.P2.Y = _beaa - _adg.P2.Y
			_adg.P3.Y = _beaa - _adg.P3.Y
			_cgcc = append(_cgcc, _adg)
			_egdd := _adg.GetBounds()
			_bbfa.Llx = _b.Min(_bbfa.Llx, _egdd.Llx)
			_bbfa.Lly = _b.Min(_bbfa.Lly, _egdd.Lly)
			_bbfa.Urx = _b.Max(_bbfa.Urx, _egdd.Urx)
			_bbfa.Ury = _b.Max(_bbfa.Ury, _egdd.Ury)
		}
		_cdca = append(_cdca, _cgcc)
	}
	_dcae.Rings = _cdca
	defer func() { _dcae.Rings = _aee }()
	if _dcae.FillEnabled {
		_cbcd := _aede(_egca, _aadf._ecae.FillColor, _aadf._fdge, func() Rectangle {
			return Rectangle{_cgedd: _bbfa.Llx, _eeff: _bbfa.Lly, _gfad: _bbfa.Width(), _fefg: _bbfa.Height()}
		})
		if _cbcd != nil {
			return nil, ctx, _cbcd
		}
	}
	_bcda, _, _fefd := _dcae.Draw(_agdf)
	if _fefd != nil {
		return nil, ctx, _fefd
	}
	if _fefd = _egca.addContentsByString(string(_bcda)); _fefd != nil {
		return nil, ctx, _fefd
	}
	return []*Block{_egca}, ctx, nil
}
func _cafcfa(_aggf Color, _fddaa float64) *ColorPoint {
	return &ColorPoint{_afea: _aggf, _dbcfd: _fddaa}
}

// DrawFooter sets a function to draw a footer on created output pages.
func (_bfddc *Creator) DrawFooter(drawFooterFunc func(_ccfc *Block, _ecgf FooterFunctionArgs)) {
	_bfddc._cgeab = drawFooterFunc
}

// SetWidth set the Image's document width to specified w. This does not change the raw image data, i.e.
// no actual scaling of data is performed. That is handled by the PDF viewer.
func (_efbbc *Image) SetWidth(w float64) { _efbbc._cagba = w }
func _gecbe(_aeeb *templateProcessor, _bfgc *templateNode) (interface{}, error) {
	return _aeeb.parseDivision(_bfgc)
}
func _aebfg(_efede *Table, _ccedf DrawContext) ([]*Block, DrawContext, error) {
	var _agage []*Block
	_cbcdg := NewBlock(_ccedf.PageWidth, _ccedf.PageHeight)
	_efede.updateRowHeights(_ccedf.Width - _efede._gfbcc.Left - _efede._gfbcc.Right)
	_fdfb := _efede._gfbcc.Top
	if _efede._bbdf.IsRelative() && !_efede._gbgc {
		_dggee := _efede.Height()
		if _dggee > _ccedf.Height-_efede._gfbcc.Top && _dggee <= _ccedf.PageHeight-_ccedf.Margins.Top-_ccedf.Margins.Bottom {
			_agage = []*Block{NewBlock(_ccedf.PageWidth, _ccedf.PageHeight-_ccedf.Y)}
			var _acae error
			if _, _ccedf, _acae = _cbag().GeneratePageBlocks(_ccedf); _acae != nil {
				return nil, _ccedf, _acae
			}
			_fdfb = 0
		}
	}
	_dbbd := _ccedf
	if _efede._bbdf.IsAbsolute() {
		_ccedf.X = _efede._fedd
		_ccedf.Y = _efede._degda
	} else {
		_ccedf.X += _efede._gfbcc.Left
		_ccedf.Y += _fdfb
		_ccedf.Width -= _efede._gfbcc.Left + _efede._gfbcc.Right
		_ccedf.Height -= _fdfb
	}
	_egdcb := _ccedf.Width
	_dada := _ccedf.X
	_aeca := _ccedf.Y
	_ceef := _ccedf.Height
	_gdeg := 0
	_agdff, _gdbcaa := -1, -1
	if _efede._dgfg {
		for _dcdbg, _adgb := range _efede._cacca {
			if _adgb._deded < _efede._aggfe {
				continue
			}
			if _adgb._deded > _efede._fecf {
				break
			}
			if _agdff < 0 {
				_agdff = _dcdbg
			}
			_gdbcaa = _dcdbg
		}
	}
	if _ddgg := _efede.wrapContent(_ccedf); _ddgg != nil {
		return nil, _ccedf, _ddgg
	}
	_efede.updateRowHeights(_ccedf.Width - _efede._gfbcc.Left - _efede._gfbcc.Right)
	var (
		_egcag bool
		_geaec int
		_bgegd int
		_gbcde bool
		_cfcg  int
		_edbd  error
	)
	for _aacg := 0; _aacg < len(_efede._cacca); _aacg++ {
		_dbgcc := _efede._cacca[_aacg]
		if _agea, _facce := _efede.getLastCellFromCol(_dbgcc._eafbd); _agea == _aacg {
			if (_facce._deded + _facce._ffgdb - 1) < _efede._fgbfga {
				for _aceeg := _dbgcc._deded; _aceeg < _efede._fgbfga; _aceeg++ {
					_bgcg := &TableCell{}
					_bgcg._deded = _aceeg + 1
					_bgcg._ffgdb = 1
					_bgcg._eafbd = _dbgcc._eafbd
					_efede._cacca = append(_efede._cacca, _bgcg)
				}
			}
		}
		_gdcg := _dbgcc.width(_efede._abbbf, _egdcb)
		_gceb := float64(0.0)
		for _efbef := 0; _efbef < _dbgcc._eafbd-1; _efbef++ {
			_gceb += _efede._abbbf[_efbef] * _egdcb
		}
		_eege := float64(0.0)
		for _bgcdb := _gdeg; _bgcdb < _dbgcc._deded-1; _bgcdb++ {
			_eege += _efede._begb[_bgcdb]
		}
		_ccedf.Height = _ceef - _eege
		_cedfd := float64(0.0)
		for _bcggeg := 0; _bcggeg < _dbgcc._ffgdb; _bcggeg++ {
			_cedfd += _efede._begb[_dbgcc._deded+_bcggeg-1]
		}
		_cbeef := _gbcde && _dbgcc._deded != _cfcg
		_cfcg = _dbgcc._deded
		if _cbeef || _cedfd > _ccedf.Height {
			if _efede._eebgf && !_gbcde {
				_gbcde, _edbd = _efede.wrapRow(_aacg, _ccedf, _egdcb)
				if _edbd != nil {
					return nil, _ccedf, _edbd
				}
				if _gbcde {
					_aacg--
					continue
				}
			}
			_agage = append(_agage, _cbcdg)
			_cbcdg = NewBlock(_ccedf.PageWidth, _ccedf.PageHeight)
			_dada = _ccedf.Margins.Left + _efede._gfbcc.Left
			_aeca = _ccedf.Margins.Top
			_ccedf.Height = _ccedf.PageHeight - _ccedf.Margins.Top - _ccedf.Margins.Bottom
			_ccedf.Page++
			_ceef = _ccedf.Height
			_gdeg = _dbgcc._deded - 1
			_eege = 0
			_gbcde = false
			if _efede._dgfg && _agdff >= 0 {
				_geaec = _aacg
				_aacg = _agdff - 1
				_bgegd = _gdeg
				_gdeg = _efede._aggfe - 1
				_egcag = true
				if _dbgcc._ffgdb > (_efede._fgbfga-_cfcg) || (_dbgcc._ffgdb > 1 && _aacg < 0) {
					_ca.Log.Debug("\u0054a\u0062\u006ce\u0020\u0068\u0065a\u0064\u0065\u0072\u0020\u0072\u006f\u0077s\u0070\u0061\u006e\u0020\u0065\u0078c\u0065\u0065\u0064\u0073\u0020\u0061\u0076\u0061\u0069\u006c\u0061b\u006c\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u002e")
					_egcag = false
					_agdff, _gdbcaa = -1, -1
				}
				continue
			}
			if _cbeef {
				_aacg--
				continue
			}
		}
		_ccedf.Width = _gdcg
		_ccedf.X = _dada + _gceb
		_ccedf.Y = _aeca + _eege
		if _cedfd > _ccedf.PageHeight-_ccedf.Margins.Top-_ccedf.Margins.Bottom {
			_cedfd = _ccedf.PageHeight - _ccedf.Margins.Top - _ccedf.Margins.Bottom
		}
		_ccgfb := _fbfe(_ccedf.X, _ccedf.Y, _gdcg, _cedfd)
		if _dbgcc._afbg != nil {
			_ccgfb.SetFillColor(_dbgcc._afbg)
		}
		_ccgfb.LineStyle = _dbgcc._fcagf
		_ccgfb._acba = _dbgcc._ffdeb
		_ccgfb._dabg = _dbgcc._dfbfd
		_ccgfb._deccc = _dbgcc._aaddc
		_ccgfb._cffe = _dbgcc._cbafe
		if _dbgcc._egde != nil {
			_ccgfb.SetColorLeft(_dbgcc._egde)
		}
		if _dbgcc._dagd != nil {
			_ccgfb.SetColorBottom(_dbgcc._dagd)
		}
		if _dbgcc._aeec != nil {
			_ccgfb.SetColorRight(_dbgcc._aeec)
		}
		if _dbgcc._afga != nil {
			_ccgfb.SetColorTop(_dbgcc._afga)
		}
		_ccgfb.SetWidthBottom(_dbgcc._egaeb)
		_ccgfb.SetWidthLeft(_dbgcc._egbd)
		_ccgfb.SetWidthRight(_dbgcc._ddbdf)
		_ccgfb.SetWidthTop(_dbgcc._caffa)
		_eecdd := NewBlock(_cbcdg._ecd, _cbcdg._gfe)
		_acecd := _cbcdg.Draw(_ccgfb)
		if _acecd != nil {
			_ca.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _acecd)
		}
		if _dbgcc._efbbe != nil {
			_eeaa := _dbgcc._efbbe.Width()
			_egbbd := _dbgcc._efbbe.Height()
			_fdae := 0.0
			switch _daedd := _dbgcc._efbbe.(type) {
			case *Paragraph:
				if _daedd._bebg {
					_eeaa = _daedd.getMaxLineWidth() / 1000.0
				}
				_eeaa += _daedd._fgcbf.Left + _daedd._fgcbf.Right
				_egbbd += _daedd._fgcbf.Top + _daedd._fgcbf.Bottom
			case *StyledParagraph:
				if _daedd._gfbb {
					_eeaa = _daedd.getMaxLineWidth() / 1000.0
				}
				_feee, _gcea, _gdbgb := _daedd.getLineMetrics(0)
				_dggeee, _bedg := _feee*_daedd._babcf, _gcea*_daedd._babcf
				if _daedd._edbcb == TextVerticalAlignmentCenter {
					_fdae = _bedg - (_gcea + (_feee+_gdbgb-_gcea)/2 + (_bedg-_gcea)/2)
				}
				if len(_daedd._aabba) == 1 {
					_egbbd = _dggeee
				} else {
					_egbbd = _egbbd - _bedg + _dggeee
				}
				_fdae += _dggeee - _bedg
				switch _dbgcc._fddca {
				case CellVerticalAlignmentTop:
					_fdae += _dggeee * 0.5
				case CellVerticalAlignmentBottom:
					_fdae -= _dggeee * 0.5
				}
				_eeaa += _daedd._fbgbc.Left + _daedd._fbgbc.Right
				_egbbd += _daedd._fbgbc.Top + _daedd._fbgbc.Bottom
			case *Table:
				_eeaa = _gdcg
			case *List:
				_eeaa = _gdcg
			case *Division:
				_eeaa = _gdcg
			case *Chart:
				_eeaa = _gdcg
			case *Line:
				_egbbd += _daedd._bccc.Top + _daedd._bccc.Bottom
				_fdae -= _daedd.Height() / 2
			case *Image:
				_eeaa += _daedd._cbea.Left + _daedd._cbea.Right
				_egbbd += _daedd._cbea.Top + _daedd._cbea.Bottom
			}
			switch _dbgcc._abad {
			case CellHorizontalAlignmentLeft:
				_ccedf.X += _dbgcc._bbgcg
				_ccedf.Width -= _dbgcc._bbgcg
			case CellHorizontalAlignmentCenter:
				if _cebc := _gdcg - _eeaa; _cebc > 0 {
					_ccedf.X += _cebc / 2
					_ccedf.Width -= _cebc / 2
				}
			case CellHorizontalAlignmentRight:
				if _gdcg > _eeaa {
					_ccedf.X = _ccedf.X + _gdcg - _eeaa - _dbgcc._bbgcg
					_ccedf.Width -= _dbgcc._bbgcg
				}
			}
			_cgced := _ccedf.Y
			_bcgb := _ccedf.Height
			_ccedf.Y += _fdae
			switch _dbgcc._fddca {
			case CellVerticalAlignmentTop:
			case CellVerticalAlignmentMiddle:
				if _cfgfb := _cedfd - _egbbd; _cfgfb > 0 {
					_ccedf.Y += _cfgfb / 2
					_ccedf.Height -= _cfgfb / 2
				}
			case CellVerticalAlignmentBottom:
				if _cedfd > _egbbd {
					_ccedf.Y = _ccedf.Y + _cedfd - _egbbd
					_ccedf.Height = _cedfd
				}
			}
			_egbbf := _cbcdg.DrawWithContext(_dbgcc._efbbe, _ccedf)
			if _egbbf != nil {
				if _fa.Is(_egbbf, ErrContentNotFit) && !_cbeef {
					_cbcdg = _eecdd
					_cbeef = true
					_aacg--
					continue
				}
				_ca.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _egbbf)
			}
			_ccedf.Y = _cgced
			_ccedf.Height = _bcgb
		}
		_ccedf.Y += _cedfd
		_ccedf.Height -= _cedfd
		if _egcag && _aacg+1 > _gdbcaa {
			_aeca += _eege + _cedfd
			_ceef -= _cedfd + _eege
			_gdeg = _bgegd
			_aacg = _geaec - 1
			_egcag = false
		}
	}
	_agage = append(_agage, _cbcdg)
	if _efede._bbdf.IsAbsolute() {
		return _agage, _dbbd, nil
	}
	_ccedf.X = _dbbd.X
	_ccedf.Width = _dbbd.Width
	_ccedf.Y += _efede._gfbcc.Bottom
	_ccedf.Height -= _efede._gfbcc.Bottom
	return _agage, _ccedf, nil
}
func (_aafce *Table) moveToNextAvailableCell() int {
	_fggd := (_aafce._ggfb-1)%(_aafce._aeaa) + 1
	for {
		if _fggd-1 >= len(_aafce._degfe) {
			if _aafce._degfe[0] == 0 {
				return _fggd
			}
			_fggd = 1
		} else if _aafce._degfe[_fggd-1] == 0 {
			return _fggd
		}
		_aafce._ggfb++
		_aafce._degfe[_fggd-1]--
		_fggd++
	}
}

// GeneratePageBlocks draws the composite Bezier curve on a new block
// representing the page. Implements the Drawable interface.
func (_bgbcc *PolyBezierCurve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_cbca := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_gcbef, _ddge := _cbca.setOpacity(_bgbcc._bcdgc, _bgbcc._gfcgc)
	if _ddge != nil {
		return nil, ctx, _ddge
	}
	_ecgb := _bgbcc._ffbb
	_ecgb.FillEnabled = _ecgb.FillColor != nil
	var (
		_gdde = ctx.PageHeight
		_dege = _ecgb.Curves
		_dgff = make([]_fc.CubicBezierCurve, 0, len(_ecgb.Curves))
	)
	_fgbfc := _ggc.PdfRectangle{}
	for _cecg := range _ecgb.Curves {
		_facdg := _dege[_cecg]
		_facdg.P0.Y = _gdde - _facdg.P0.Y
		_facdg.P1.Y = _gdde - _facdg.P1.Y
		_facdg.P2.Y = _gdde - _facdg.P2.Y
		_facdg.P3.Y = _gdde - _facdg.P3.Y
		_dgff = append(_dgff, _facdg)
		_eegdf := _facdg.GetBounds()
		if _cecg == 0 {
			_fgbfc = _eegdf
		} else {
			_fgbfc.Llx = _b.Min(_fgbfc.Llx, _eegdf.Llx)
			_fgbfc.Lly = _b.Min(_fgbfc.Lly, _eegdf.Lly)
			_fgbfc.Urx = _b.Max(_fgbfc.Urx, _eegdf.Urx)
			_fgbfc.Ury = _b.Max(_fgbfc.Ury, _eegdf.Ury)
		}
	}
	_ecgb.Curves = _dgff
	defer func() { _ecgb.Curves = _dege }()
	if _ecgb.FillEnabled {
		_addb := _aede(_cbca, _bgbcc._ffbb.FillColor, _bgbcc._bbaec, func() Rectangle {
			return Rectangle{_cgedd: _fgbfc.Llx, _eeff: _fgbfc.Lly, _gfad: _fgbfc.Width(), _fefg: _fgbfc.Height()}
		})
		if _addb != nil {
			return nil, ctx, _addb
		}
	}
	_bcfad, _, _ddge := _ecgb.Draw(_gcbef)
	if _ddge != nil {
		return nil, ctx, _ddge
	}
	if _ddge = _cbca.addContentsByString(string(_bcfad)); _ddge != nil {
		return nil, ctx, _ddge
	}
	return []*Block{_cbca}, ctx, nil
}

// MoveRight moves the drawing context right by relative displacement dx (negative goes left).
func (_bgd *Creator) MoveRight(dx float64) { _bgd._eacd.X += dx }

// FillColor returns the fill color of the rectangle.
func (_bfde *Rectangle) FillColor() Color { return _bfde._bgca }

// AddTotalLine adds a new line in the invoice totals table.
func (_bgaf *Invoice) AddTotalLine(desc, value string) (*InvoiceCell, *InvoiceCell) {
	_efbbf := &InvoiceCell{_bgaf._ddde, desc}
	_caea := &InvoiceCell{_bgaf._ddde, value}
	_bgaf._edb = append(_bgaf._edb, [2]*InvoiceCell{_efbbf, _caea})
	return _efbbf, _caea
}
func (_dadf *Division) split(_adcbb DrawContext) (_eceg, _bbad *Division) {
	var (
		_bcfbe      float64
		_cggd, _bfc []VectorDrawable
	)
	_deaf := _adcbb.Width - _dadf._debb.Left - _dadf._debb.Right - _dadf._agda.Left - _dadf._agda.Right
	for _edca, _ddcc := range _dadf._bfdf {
		_bcfbe += _cefg(_ddcc, _deaf)
		if _bcfbe < _adcbb.Height {
			_cggd = append(_cggd, _ddcc)
		} else {
			_bfc = _dadf._bfdf[_edca:]
			break
		}
	}
	if len(_cggd) > 0 {
		_eceg = _ffcc()
		*_eceg = *_dadf
		_eceg._bfdf = _cggd
		if _dadf._fbgb != nil {
			_eceg._fbgb = &Background{}
			*_eceg._fbgb = *_dadf._fbgb
		}
	}
	if len(_bfc) > 0 {
		_bbad = _ffcc()
		*_bbad = *_dadf
		_bbad._bfdf = _bfc
		if _dadf._fbgb != nil {
			_bbad._fbgb = &Background{}
			*_bbad._fbgb = *_dadf._fbgb
		}
	}
	return _eceg, _bbad
}

// MultiColCell makes a new cell with the specified column span and inserts it
// into the table at the current position.
func (_gcabg *Table) MultiColCell(colspan int) *TableCell { return _gcabg.MultiCell(1, colspan) }

// SetFillOpacity sets the fill opacity.
func (_cbdf *CurvePolygon) SetFillOpacity(opacity float64) { _cbdf._gfage = opacity }

// InsertColumn inserts a column in the line items table at the specified index.
func (_ddga *Invoice) InsertColumn(index uint, description string) *InvoiceCell {
	_gcfc := uint(len(_ddga._eag))
	if index > _gcfc {
		index = _gcfc
	}
	_gbbb := _ddga.NewColumn(description)
	_ddga._eag = append(_ddga._eag[:index], append([]*InvoiceCell{_gbbb}, _ddga._eag[index:]...)...)
	return _gbbb
}
func (_dcbgd *TextChunk) clone() *TextChunk {
	_fcfc := *_dcbgd
	_fcfc._dfcb = _bgbee(_dcbgd._dfcb)
	return &_fcfc
}
func (_gdga *StyledParagraph) getTextHeight() float64 {
	var _adfc float64
	for _, _gaeg := range _gdga._cdffa {
		_bebga := _gaeg.Style.FontSize * _gdga._babcf
		if _bebga > _adfc {
			_adfc = _bebga
		}
	}
	return _adfc
}
func (_egfbc *templateProcessor) parseFloatAttr(_cdffd, _afcbd string) float64 {
	_ca.Log.Debug("\u0050\u0061rs\u0069\u006e\u0067 \u0066\u006c\u006f\u0061t a\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020(`\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _cdffd, _afcbd)
	_adgc, _ := _a.ParseFloat(_afcbd, 64)
	return _adgc
}
func (_bcega *Table) resetColumnWidths() {
	_bcega._abbbf = []float64{}
	_ddged := float64(1.0) / float64(_bcega._aeaa)
	for _fdab := 0; _fdab < _bcega._aeaa; _fdab++ {
		_bcega._abbbf = append(_bcega._abbbf, _ddged)
	}
}
func _fbfe(_gae, _eff, _eddf, _bdbb float64) *border {
	_acdb := &border{}
	_acdb._agf = _gae
	_acdb._gea = _eff
	_acdb._dcf = _eddf
	_acdb._aebc = _bdbb
	_acdb._feg = ColorBlack
	_acdb._agg = ColorBlack
	_acdb._fcg = ColorBlack
	_acdb._add = ColorBlack
	_acdb._cdgaa = 0
	_acdb._gaf = 0
	_acdb._gcd = 0
	_acdb._ebb = 0
	_acdb.LineStyle = _fc.LineStyleSolid
	return _acdb
}

// Width returns the cell's width based on the input draw context.
func (_efcb *TableCell) Width(ctx DrawContext) float64 {
	_accc := float64(0.0)
	for _gfde := 0; _gfde < _efcb._abfda; _gfde++ {
		_accc += _efcb._ddggg._abbbf[_efcb._eafbd+_gfde-1]
	}
	_afdg := ctx.Width * _accc
	return _afdg
}

// SetForms adds an Acroform to a PDF file.  Sets the specified form for writing.
func (_bddd *Creator) SetForms(form *_ggc.PdfAcroForm) error { _bddd._ddc = form; return nil }

// GeneratePageBlocks draws the curve onto page blocks.
func (_egfb *Curve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_dcbg := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_baab := _bdb.NewContentCreator()
	_baab.Add_q().Add_w(_egfb._bceaa).SetStrokingColor(_dbac(_egfb._ccad)).Add_m(_egfb._ebbc, ctx.PageHeight-_egfb._gbae).Add_v(_egfb._cfac, ctx.PageHeight-_egfb._cfdb, _egfb._bceb, ctx.PageHeight-_egfb._bfea).Add_S().Add_Q()
	_becc := _dcbg.addContentsByString(_baab.String())
	if _becc != nil {
		return nil, ctx, _becc
	}
	return []*Block{_dcbg}, ctx, nil
}

// FrontpageFunctionArgs holds the input arguments to a front page drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type FrontpageFunctionArgs struct {
	PageNum    int
	TotalPages int
}

// NewPageBreak create a new page break.
func (_gebg *Creator) NewPageBreak() *PageBreak { return _cbag() }

// Height returns the height of the chart.
func (_dggg *Chart) Height() float64 { return float64(_dggg._ebbf.Height()) }
func _bga(_gab _gg.ChartRenderable) *Chart {
	return &Chart{_ebbf: _gab, _fbc: PositionRelative, _gcff: Margins{Top: 10, Bottom: 10}}
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

func (_acga *templateProcessor) parseDivision(_abae *templateNode) (interface{}, error) {
	_cfdfac := _acga.creator.NewDivision()
	for _, _agada := range _abae._gbdee.Attr {
		_bgaa := _agada.Value
		switch _bdab := _agada.Name.Local; _bdab {
		case "\u0065\u006ea\u0062\u006c\u0065-\u0070\u0061\u0067\u0065\u002d\u0077\u0072\u0061\u0070":
			_cfdfac.EnablePageWrap(_acga.parseBoolAttr(_bdab, _bgaa))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_ebde := _acga.parseMarginAttr(_bdab, _bgaa)
			_cfdfac.SetMargins(_ebde.Left, _ebde.Right, _ebde.Top, _ebde.Bottom)
		case "\u0070a\u0064\u0064\u0069\u006e\u0067":
			_dgage := _acga.parseMarginAttr(_bdab, _bgaa)
			_cfdfac.SetPadding(_dgage.Left, _dgage.Right, _dgage.Top, _dgage.Bottom)
		default:
			_acga.nodeLogDebug(_abae, "U\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0064\u0069\u0076\u0069\u0073\u0069on\u0020\u0061\u0074\u0074r\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025s`\u002e\u0020S\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _bdab)
		}
	}
	return _cfdfac, nil
}

// SetFont sets the Paragraph's font.
func (_dce *Paragraph) SetFont(font *_ggc.PdfFont) { _dce._fggb = font }
func (_ddbgf *templateProcessor) parseTextVerticalAlignmentAttr(_bcagf, _gacdd string) TextVerticalAlignment {
	_ca.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0074\u0065\u0078\u0074\u0020\u0076\u0065r\u0074\u0069\u0063\u0061\u006c\u0020\u0061\u006c\u0069\u0067\u006e\u006d\u0065n\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a (\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _bcagf, _gacdd)
	_bdbcbc := map[string]TextVerticalAlignment{"\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065": TextVerticalAlignmentBaseline, "\u0063\u0065\u006e\u0074\u0065\u0072": TextVerticalAlignmentCenter}[_gacdd]
	return _bdbcbc
}

// SkipRows skips over a specified number of rows in the table.
func (_baec *Table) SkipRows(num int) {
	_cagd := num*_baec._aeaa - 1
	if _cagd < 0 {
		_ca.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	for _bcbee := 0; _bcbee < _cagd; _bcbee++ {
		_baec.NewCell()
	}
}
func _gcdfe(_cdeag, _dfbe, _faggb float64) (_beacg, _begfe, _aeefe, _baaca float64) {
	if _faggb == 0 {
		return 0, 0, _cdeag, _dfbe
	}
	_bgfbc := _fc.Path{Points: []_fc.Point{_fc.NewPoint(0, 0).Rotate(_faggb), _fc.NewPoint(_cdeag, 0).Rotate(_faggb), _fc.NewPoint(0, _dfbe).Rotate(_faggb), _fc.NewPoint(_cdeag, _dfbe).Rotate(_faggb)}}.GetBoundingBox()
	return _bgfbc.X, _bgfbc.Y, _bgfbc.Width, _bgfbc.Height
}

// SetBackground sets the background properties of the component.
func (_fbae *Division) SetBackground(background *Background) { _fbae._fbgb = background }

// SkipCells skips over a specified number of cells in the table.
func (_ebgbb *Table) SkipCells(num int) {
	if num < 0 {
		_ca.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	for _ceec := 0; _ceec < num; _ceec++ {
		_ebgbb.NewCell()
	}
}

// NewCurvePolygon creates a new curve polygon.
func (_gegf *Creator) NewCurvePolygon(rings [][]_fc.CubicBezierCurve) *CurvePolygon {
	return _cagg(rings)
}

// SetMaxLines sets the maximum number of lines before the paragraph
// text is truncated.
func (_dfge *Paragraph) SetMaxLines(maxLines int) { _dfge._bebd = maxLines; _dfge.wrapText() }

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
	_bfdf  []VectorDrawable
	_badg  Positioning
	_debb  Margins
	_agda  Margins
	_aeadf bool
	_ggdbd bool
	_fbgb  *Background
}

func (_ab *Block) translate(_cge, _decc float64) {
	_fcb := _bdb.NewContentCreator().Translate(_cge, -_decc).Operations()
	*_ab._cad = append(*_fcb, *_ab._cad...)
	_ab._cad.WrapIfNeeded()
}
func _bgbgb(_faeda *Creator, _afdeg string, _egbef []byte, _fagag *TemplateOptions, _fabg componentRenderer) *templateProcessor {
	if _fagag == nil {
		_fagag = &TemplateOptions{}
	}
	_fagag.init()
	if _fabg == nil {
		_fabg = _faeda
	}
	return &templateProcessor{creator: _faeda, _gcgb: _egbef, _affcb: _fagag, _gcfaeb: _fabg, _eccabe: _afdeg}
}

// CellVerticalAlignment defines the table cell's vertical alignment.
type CellVerticalAlignment int

func _dced(_caffg *Creator, _ffga _ae.Reader, _fgbff interface{}, _aegga *TemplateOptions, _cecgd componentRenderer) error {
	if _caffg == nil {
		_ca.Log.Error("\u0043\u0072\u0065a\u0074\u006f\u0072\u0020i\u006e\u0073\u0074\u0061\u006e\u0063\u0065 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e")
		return _ddead
	}
	_gadcad := ""
	if _ddbb, _efccf := _ffga.(*_ed.File); _efccf {
		_gadcad = _ddbb.Name()
	}
	_eegc := _g.NewBuffer(nil)
	if _, _gceed := _ae.Copy(_eegc, _ffga); _gceed != nil {
		return _gceed
	}
	_fggee := _dg.FuncMap{"\u0064\u0069\u0063\u0074": _agbgc}
	if _aegga != nil && _aegga.HelperFuncMap != nil {
		for _eaae, _fefc := range _aegga.HelperFuncMap {
			if _, _dfea := _fggee[_eaae]; _dfea {
				_ca.Log.Debug("\u0043\u0061\u006e\u006e\u006f\u0074 \u006f\u0076\u0065r\u0072\u0069\u0064e\u0020\u0062\u0075\u0069\u006c\u0074\u002d\u0069\u006e\u0020`\u0025\u0073\u0060\u0020\u0068el\u0070\u0065\u0072\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _eaae)
				continue
			}
			_fggee[_eaae] = _fefc
		}
	}
	_gdccb, _eeec := _dg.New("").Funcs(_fggee).Parse(_eegc.String())
	if _eeec != nil {
		return _eeec
	}
	if _aegga != nil && _aegga.SubtemplateMap != nil {
		for _acbb, _aabge := range _aegga.SubtemplateMap {
			if _acbb == "" {
				_ca.Log.Debug("\u0053\u0075\u0062\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u0063\u0061\u006en\u006f\u0074\u0020\u0062\u0065\u0020\u0065\u006d\u0070\u0074\u0079\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067.\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065 \u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e")
				continue
			}
			if _aabge == nil {
				_ca.Log.Debug("S\u0075\u0062t\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0063\u0061\u006e\u006eo\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079 \u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e")
				continue
			}
			_egaeg := _g.NewBuffer(nil)
			if _, _fadf := _ae.Copy(_egaeg, _aabge); _fadf != nil {
				return _fadf
			}
			if _, _dfdf := _gdccb.New(_acbb).Parse(_egaeg.String()); _dfdf != nil {
				return _dfdf
			}
		}
	}
	_eegc.Reset()
	if _agca := _gdccb.Execute(_eegc, _fgbff); _agca != nil {
		return _agca
	}
	return _bgbgb(_caffg, _gadcad, _eegc.Bytes(), _aegga, _cecgd).run()
}

// Creator is a wrapper around functionality for creating PDF reports and/or adding new
// content onto imported PDF pages, etc.
type Creator struct {

	// Errors keeps error messages that should not interrupt pdf processing and to be checked later.
	Errors []error

	// UnsupportedCharacterReplacement is character that will be used to replace unsupported glyph.
	// The value will be passed to drawing context.
	UnsupportedCharacterReplacement rune
	_cec                            []*_ggc.PdfPage
	_gfab                           map[*_ggc.PdfPage]*Block
	_afbc                           map[*_ggc.PdfPage]*pageTransformations
	_bcba                           *_ggc.PdfPage
	_gccce                          PageSize
	_eacd                           DrawContext
	_gcgd                           Margins
	_abf, _ffc                      float64
	_caffe                          int
	_fbd                            func(_fab FrontpageFunctionArgs)
	_gfgf                           func(_facf *TOC) error
	_fdcb                           func(_dgbb *Block, _faf HeaderFunctionArgs)
	_cgeab                          func(_bec *Block, _cecb FooterFunctionArgs)
	_dcc                            func(_gcfa PageFinalizeFunctionArgs) error
	_bggc                           func(_cae *_ggc.PdfWriter) error
	_acead                          bool

	// Controls whether a table of contents will be generated.
	AddTOC bool

	// CustomTOC specifies if the TOC is rendered by the user.
	// When the `CustomTOC` field is set to `true`, the default TOC component is not rendered.
	// Instead the TOC is drawn by the user, in the callback provided to
	// the `Creator.CreateTableOfContents` method.
	// If `CustomTOC` is set to `false`, the callback provided to
	// `Creator.CreateTableOfContents` customizes the style of the automatically generated TOC component.
	CustomTOC bool
	_effa     *TOC

	// Controls whether outlines will be generated.
	AddOutlines bool
	_gaec       *_ggc.Outline
	_beeb       *_ggc.PdfOutlineTreeNode
	_ddc        *_ggc.PdfAcroForm
	_dbgc       _fe.PdfObject
	_eecf       _ggc.Optimizer
	_feff       []*_ggc.PdfFont
	_gade       *_ggc.PdfFont
	_eacc       *_ggc.PdfFont
}

// SetAngle sets the rotation angle of the text.
func (_egdb *Paragraph) SetAngle(angle float64) { _egdb._dgeg = angle }

// SetDashPattern sets the dash pattern of the line.
// NOTE: the dash pattern is taken into account only if the style of the
// line is set to dashed.
func (_babb *Line) SetDashPattern(dashArray []int64, dashPhase int64) {
	_babb._fgbfg = dashArray
	_babb._bgfed = dashPhase
}

// SetPos sets the absolute position. Changes object positioning to absolute.
func (_bbbd *Image) SetPos(x, y float64) {
	_bbbd._ebfa = PositionAbsolute
	_bbbd._fce = x
	_bbbd._gdd = y
}

// SetColorLeft sets border color for left.
func (_gbbe *border) SetColorLeft(col Color) { _gbbe._fcg = col }

// FillColor returns the fill color of the ellipse.
func (_badce *Ellipse) FillColor() Color { return _badce._dbadf }

// SetFillOpacity sets the fill opacity.
func (_gfec *Polygon) SetFillOpacity(opacity float64) { _gfec._befea = opacity }

// CellBorderSide defines the table cell's border side.
type CellBorderSide int

// Height returns the total height of all rows.
func (_aedg *Table) Height() float64 {
	_gbef := float64(0.0)
	for _, _fceb := range _aedg._begb {
		_gbef += _fceb
	}
	return _gbef
}

// SetExtends specifies whether to extend the shading beyond the starting and ending points.
//
// Text extends is set to `[]bool{false, false}` by default.
func (_bgfb *shading) SetExtends(start bool, end bool) { _bgfb._fgdd = []bool{start, end} }

// PageSize represents the page size as a 2 element array representing the width and height in PDF document units (points).
type PageSize [2]float64

func _aede(_aceca *Block, _gecb _ggc.PdfColor, _gcad Color, _egacb func() Rectangle) error {
	switch _efbf := _gecb.(type) {
	case *_ggc.PdfColorPatternType2:
		_dead, _dffa := _gcad.(*LinearShading)
		if !_dffa {
			return _df.Errorf("\u0043\u006f\u006c\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u004c\u0069\u006e\u0065\u0061\u0072\u0053\u0068\u0061d\u0069\u006e\u0067")
		}
		_gafe := _egacb()
		_dead.SetBoundingBox(_gafe._cgedd, _gafe._eeff, _gafe._gfad, _gafe._fefg)
		_egfg, _bbca := _dead.AddPatternResource(_aceca)
		if _bbca != nil {
			return _df.Errorf("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0074\u006f \u0072\u0065\u0073\u006f\u0075r\u0063\u0065s\u003a\u0020\u0025\u0076", _bbca)
		}
		_efbf.PatternName = _egfg
	case *_ggc.PdfColorPatternType3:
		_gedcf, _eaddc := _gcad.(*RadialShading)
		if !_eaddc {
			return _df.Errorf("\u0043\u006f\u006c\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0052\u0061\u0064\u0069\u0061\u006c\u0053\u0068\u0061d\u0069\u006e\u0067")
		}
		_adbca := _egacb()
		_gedcf.SetBoundingBox(_adbca._cgedd, _adbca._eeff, _adbca._gfad, _adbca._fefg)
		_fbbd, _egff := _gedcf.AddPatternResource(_aceca)
		if _egff != nil {
			return _df.Errorf("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0064\u0064\u0069\u006e\u0067\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0074\u006f \u0072\u0065\u0073\u006f\u0075r\u0063\u0065s\u003a\u0020\u0025\u0076", _egff)
		}
		_efbf.PatternName = _fbbd
	}
	return nil
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
func (_bdga *Division) Add(d VectorDrawable) error {
	switch _ffed := d.(type) {
	case *Paragraph, *StyledParagraph, *Image, *Chart, *Rectangle, *Ellipse, *Line, *Table, *Division, *List:
	case containerDrawable:
		_gdfd, _fbcd := _ffed.ContainerComponent(_bdga)
		if _fbcd != nil {
			return _fbcd
		}
		_ccagb, _bdee := _gdfd.(VectorDrawable)
		if !_bdee {
			return _df.Errorf("\u0072\u0065\u0073\u0075\u006ct\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0061\u0069\u006e\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u002d\u0020\u0025\u0054\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0056\u0065c\u0074\u006f\u0072\u0044\u0072\u0061\u0077\u0061\u0062\u006c\u0065\u0020i\u006e\u0074\u0065\u0072\u0066\u0061c\u0065", _gdfd)
		}
		d = _ccagb
	default:
		return _fa.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070e\u0020i\u006e\u0020\u0044\u0069\u0076\u0069\u0073i\u006f\u006e")
	}
	_bdga._bfdf = append(_bdga._bfdf, d)
	return nil
}

// SetWidthRight sets border width for right.
func (_dbe *border) SetWidthRight(bw float64) { _dbe._ebb = bw }
func (_aecgg *Line) computeCoords(_gefc DrawContext) (_dgceb, _gecde, _aff, _dgbdf float64) {
	_dgceb = _gefc.X
	_aff = _dgceb + _aecgg._eddg - _aecgg._daed
	_bdcdb := _aecgg._adf
	if _aecgg._daed == _aecgg._eddg {
		_bdcdb /= 2
	}
	if _aecgg._gfaa < _aecgg._eabb {
		_gecde = _gefc.PageHeight - _gefc.Y - _bdcdb
		_dgbdf = _gecde - _aecgg._eabb + _aecgg._gfaa
	} else {
		_dgbdf = _gefc.PageHeight - _gefc.Y - _bdcdb
		_gecde = _dgbdf - _aecgg._gfaa + _aecgg._eabb
	}
	switch _aecgg._afad {
	case FitModeFillWidth:
		_aff = _dgceb + _gefc.Width
	}
	return _dgceb, _gecde, _aff, _dgbdf
}
func (_ffcd *templateProcessor) processGradientColorPair(_dcffe []string) (_ffgef []Color, _bccb []float64) {
	for _, _afba := range _dcffe {
		var (
			_ebca  = _dc.Fields(_afba)
			_gbfed = len(_ebca)
		)
		if _gbfed == 0 {
			continue
		}
		_deadfa := ""
		if _gbfed > 1 {
			_deadfa = _dc.TrimSpace(_ebca[1])
		}
		_faea := -1.0
		if _dc.HasSuffix(_deadfa, "\u0025") {
			_gdfge, _cebgf := _a.ParseFloat(_deadfa[:len(_deadfa)-1], 64)
			if _cebgf != nil {
				_ca.Log.Debug("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0070\u0061\u0072s\u0069\u006e\u0067\u0020\u0070\u006f\u0069n\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _cebgf)
			}
			_faea = _gdfge / 100.0
		}
		_geceg := _ffcd.parseColor(_dc.TrimSpace(_ebca[0]))
		if _geceg != nil {
			_ffgef = append(_ffgef, _geceg)
			_bccb = append(_bccb, _faea)
		}
	}
	if len(_ffgef) != len(_bccb) {
		_ca.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u006c\u0069\u006e\u0065\u0061\u0072\u0020\u0067\u0072\u0061\u0064i\u0065\u006e\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0064\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0021")
		return nil, nil
	}
	_eebac := -1
	_babge := 0.0
	for _aecgb, _dgbac := range _bccb {
		if _dgbac == -1.0 {
			if _aecgb == 0 {
				_dgbac = 0.0
				_bccb[_aecgb] = 0.0
				continue
			}
			_eebac++
			if _aecgb < len(_bccb)-1 {
				continue
			} else {
				_dgbac = 1.0
				_bccb[_aecgb] = 1.0
			}
		}
		_adgbc := _eebac + 1
		for _abgg := _aecgb - _eebac; _abgg < _aecgb; _abgg++ {
			_bccb[_abgg] = _babge + (float64(_abgg) * (_dgbac - _babge) / float64(_adgbc))
		}
		_babge = _dgbac
		_eebac = -1
	}
	return _ffgef, _bccb
}

// Horizontal returns total horizontal (left + right) margin.
func (_gege *Margins) Horizontal() float64 { return _gege.Left + _gege.Right }

// SetAnnotation sets a annotation on a TextChunk.
func (_aafee *TextChunk) SetAnnotation(annotation *_ggc.PdfAnnotation) { _aafee._dfcb = annotation }

// SetAntiAlias enables anti alias config.
//
// Anti alias is disabled by default.
func (_egbe *shading) SetAntiAlias(enable bool) { _egbe._daff = enable }

// Height returns the height of the Paragraph. The height is calculated based on the input text and
// how it is wrapped within the container. Does not include Margins.
func (_cgfd *Paragraph) Height() float64 {
	_cgfd.wrapText()
	return float64(len(_cgfd._cbcf)) * _cgfd._dacae * _cgfd._fcbfa
}

const (
	CellHorizontalAlignmentLeft CellHorizontalAlignment = iota
	CellHorizontalAlignmentCenter
	CellHorizontalAlignmentRight
)

// SetShowLinks sets visibility of links for the TOC lines.
func (_gbdad *TOC) SetShowLinks(showLinks bool) { _gbdad._dbaf = showLinks }
func (_cgfcd *templateProcessor) parseLinearGradientAttr(creator *Creator, _efdcag string) Color {
	_fgee := ColorBlack
	if _efdcag == "" {
		return _fgee
	}
	_fbfaa := creator.NewLinearGradientColor([]*ColorPoint{})
	_fbfaa.SetExtends(true, true)
	var (
		_addg  = _dc.Split(_efdcag[16:len(_efdcag)-1], "\u002c")
		_cbeea = _dc.TrimSpace(_addg[0])
	)
	if _dc.HasSuffix(_cbeea, "\u0064\u0065\u0067") {
		_dbcfdg, _ffef := _a.ParseFloat(_cbeea[:len(_cbeea)-3], 64)
		if _ffef != nil {
			_ca.Log.Debug("\u0046\u0061\u0069\u006c\u0065\u0064 \u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0067\u0072\u0061\u0064\u0069e\u006e\u0074\u0020\u0061\u006e\u0067\u006ce\u003a\u0020\u0025\u0076", _ffef)
		} else {
			_fbfaa.SetAngle(_dbcfdg)
		}
		_addg = _addg[1:]
	}
	_agbg, _fcff := _cgfcd.processGradientColorPair(_addg)
	if _agbg == nil || _fcff == nil {
		return _fgee
	}
	for _dabb := 0; _dabb < len(_agbg); _dabb++ {
		_fbfaa.AddColorStop(_agbg[_dabb], _fcff[_dabb])
	}
	return _fbfaa
}

// ScaleToHeight scales the rectangle to the specified height. The width of
// the rectangle is scaled so that the aspect ratio is maintained.
func (_eeab *Rectangle) ScaleToHeight(h float64) {
	_cebg := _eeab._gfad / _eeab._fefg
	_eeab._fefg = h
	_eeab._gfad = h * _cebg
}
func (_bfg *Block) drawToPage(_aaf *_ggc.PdfPage) error {
	_def := &_bdb.ContentStreamOperations{}
	if _aaf.Resources == nil {
		_aaf.Resources = _ggc.NewPdfPageResources()
	}
	_ff := _ddf(_def, _aaf.Resources, _bfg._cad, _bfg._ge)
	if _ff != nil {
		return _ff
	}
	if _ff = _ega(_bfg._ge, _aaf.Resources); _ff != nil {
		return _ff
	}
	if _ff = _aaf.AppendContentBytes(_def.Bytes(), true); _ff != nil {
		return _ff
	}
	for _, _cafc := range _bfg._ga {
		_aaf.AddAnnotation(_cafc)
	}
	return nil
}

// EnablePageWrap controls whether the table is wrapped across pages.
// If disabled, the table is moved in its entirety on a new page, if it
// does not fit in the available height. By default, page wrapping is enabled.
// If the height of the table is larger than an entire page, wrapping is
// enabled automatically in order to avoid unwanted behavior.
func (_aagg *Table) EnablePageWrap(enable bool) { _aagg._gbgc = enable }

// SetStyleLeft sets border style for left side.
func (_gec *border) SetStyleLeft(style CellBorderStyle) { _gec._acba = style }

// Height returns the Block's height.
func (_fef *Block) Height() float64 { return _fef._gfe }

// AddAnnotation adds an annotation to the current block.
// The annotation will be added to the page the block will be rendered on.
func (_eeb *Block) AddAnnotation(annotation *_ggc.PdfAnnotation) {
	for _, _dd := range _eeb._ga {
		if _dd == annotation {
			return
		}
	}
	_eeb._ga = append(_eeb._ga, annotation)
}

// TextRenderingMode determines whether showing text shall cause glyph
// outlines to be stroked, filled, used as a clipping boundary, or some
// combination of the three.
// See section 9.3 "Text State Parameters and Operators" and
// Table 106 (pp. 254-255 PDF32000_2008).
type TextRenderingMode int

// SetEnableWrap sets the line wrapping enabled flag.
func (_fccgb *StyledParagraph) SetEnableWrap(enableWrap bool) {
	_fccgb._gfbb = enableWrap
	_fccgb._facb = false
}
func _ecaea(_accf []_fc.Point) *Polyline {
	return &Polyline{_edbcg: &_fc.Polyline{Points: _accf, LineColor: _ggc.NewPdfColorDeviceRGB(0, 0, 0), LineWidth: 1.0}, _afcb: 1.0}
}

// SetPos sets absolute positioning with specified coordinates.
func (_bbeea *StyledParagraph) SetPos(x, y float64) {
	_bbeea._gbga = PositionAbsolute
	_bbeea._gcfda = x
	_bbeea._cfge = y
}

// SetMargins sets the Table's left, right, top, bottom margins.
func (_feedg *Table) SetMargins(left, right, top, bottom float64) {
	_feedg._gfbcc.Left = left
	_feedg._gfbcc.Right = right
	_feedg._gfbcc.Top = top
	_feedg._gfbcc.Bottom = bottom
}
func (_ebfgb *templateProcessor) loadImageFromSrc(_gfcaf string) (*Image, error) {
	if _gfcaf == "" {
		_ca.Log.Error("\u0049\u006d\u0061\u0067\u0065\u0020\u0060\u0073\u0072\u0063\u0060\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062e\u0020\u0065m\u0070\u0074\u0079\u002e")
		return nil, _gggfb
	}
	_fbbeg := _dc.Split(_gfcaf, "\u002c")
	for _, _feeea := range _fbbeg {
		_feeea = _dc.TrimSpace(_feeea)
		if _feeea == "" {
			continue
		}
		_adbce, _agaca := _ebfgb._affcb.ImageMap[_feeea]
		if _agaca {
			return _ggba(_adbce)
		}
		if _fdcgfgg := _ebfgb.parseAttrPropList(_feeea); len(_fdcgfgg) > 0 {
			if _abgeg, _gcage := _fdcgfgg["\u0070\u0061\u0074\u0068"]; _gcage {
				if _baeg, _afaa := _ggde(_abgeg); _afaa != nil {
					_ca.Log.Debug("\u0043\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0069\u006d\u0061g\u0065\u0020\u0060\u0025\u0073\u0060\u003a\u0020\u0025\u0076\u002e", _abgeg, _afaa)
				} else {
					return _baeg, nil
				}
			}
		}
	}
	_ca.Log.Error("\u0043\u006ful\u0064\u0020\u006eo\u0074\u0020\u0066\u0069nd \u0069ma\u0067\u0065\u0020\u0072\u0065\u0073\u006fur\u0063\u0065\u003a\u0020\u0060\u0025\u0073`\u002e", _gfcaf)
	return nil, _gggfb
}
func (_bbcd *Invoice) newCell(_gbfe string, _fced InvoiceCellProps) *InvoiceCell {
	return &InvoiceCell{_fced, _gbfe}
}
func (_ccadf *templateProcessor) parseChart(_ecbb *templateNode) (interface{}, error) {
	var _acda string
	for _, _cfdea := range _ecbb._gbdee.Attr {
		_eeedf := _cfdea.Value
		switch _effcc := _cfdea.Name.Local; _effcc {
		case "\u0073\u0072\u0063":
			_acda = _eeedf
		}
	}
	if _acda == "" {
		_ccadf.nodeLogError(_ecbb, "\u0043\u0068\u0061\u0072\u0074\u0020\u0060\u0073\u0072\u0063\u0060\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062e\u0020\u0065m\u0070\u0074\u0079\u002e")
		return nil, _gggfb
	}
	_accfb, _gcgba := _ccadf._affcb.ChartMap[_acda]
	if !_gcgba {
		_ccadf.nodeLogError(_ecbb, "\u0043\u006ful\u0064\u0020\u006eo\u0074\u0020\u0066\u0069nd \u0063ha\u0072\u0074\u0020\u0072\u0065\u0073\u006fur\u0063\u0065\u003a\u0020\u0060\u0025\u0073`\u002e", _acda)
		return nil, _gggfb
	}
	_fdbc := NewChart(_accfb)
	for _, _eadef := range _ecbb._gbdee.Attr {
		_efbbd := _eadef.Value
		switch _gfcbb := _eadef.Name.Local; _gfcbb {
		case "\u0078":
			_fdbc.SetPos(_ccadf.parseFloatAttr(_gfcbb, _efbbd), _fdbc._gbbc)
		case "\u0079":
			_fdbc.SetPos(_fdbc._bfe, _ccadf.parseFloatAttr(_gfcbb, _efbbd))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_cfgff := _ccadf.parseMarginAttr(_gfcbb, _efbbd)
			_fdbc.SetMargins(_cfgff.Left, _cfgff.Right, _cfgff.Top, _cfgff.Bottom)
		case "\u0077\u0069\u0064t\u0068":
			_fdbc._ebbf.SetWidth(int(_ccadf.parseFloatAttr(_gfcbb, _efbbd)))
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_fdbc._ebbf.SetHeight(int(_ccadf.parseFloatAttr(_gfcbb, _efbbd)))
		case "\u0073\u0072\u0063":
		default:
			_ccadf.nodeLogDebug(_ecbb, "\u0055n\u0073\u0075p\u0070\u006f\u0072\u0074e\u0064\u0020\u0063h\u0061\u0072\u0074\u0020\u0061\u0074\u0074\u0072\u0069bu\u0074\u0065\u003a \u0060\u0025s\u0060\u002e\u0020\u0053\u006b\u0069p\u0070\u0069n\u0067\u002e", _gfcbb)
		}
	}
	return _fdbc, nil
}
func _caaf(_fgefg []_fc.CubicBezierCurve) *PolyBezierCurve {
	return &PolyBezierCurve{_ffbb: &_fc.PolyBezierCurve{Curves: _fgefg, BorderColor: _ggc.NewPdfColorDeviceRGB(0, 0, 0), BorderWidth: 1.0}, _bcdgc: 1.0, _gfcgc: 1.0}
}
func _cdce(_dccg []byte) (*Image, error) {
	_cabe := _g.NewReader(_dccg)
	_febec, _gdfe := _ggc.ImageHandling.Read(_cabe)
	if _gdfe != nil {
		_ca.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _gdfe)
		return nil, _gdfe
	}
	return _ggba(_febec)
}
func _bfefg(_edfdd int64, _bgdfc, _ddcf, _badca float64) *_ggc.PdfAnnotation {
	_eafbe := _ggc.NewPdfAnnotationLink()
	_cadf := _ggc.NewBorderStyle()
	_cadf.SetBorderWidth(0)
	_eafbe.BS = _cadf.ToPdfObject()
	if _edfdd < 0 {
		_edfdd = 0
	}
	_eafbe.Dest = _fe.MakeArray(_fe.MakeInteger(_edfdd), _fe.MakeName("\u0058\u0059\u005a"), _fe.MakeFloat(_bgdfc), _fe.MakeFloat(_ddcf), _fe.MakeFloat(_badca))
	return _eafbe.PdfAnnotation
}

// SetWidth sets the width of the ellipse.
func (_afda *Ellipse) SetWidth(width float64) { _afda._eded = width }
func (_decbe *templateProcessor) parseParagraph(_caffcf *templateNode, _ccgea *Paragraph) (interface{}, error) {
	if _ccgea == nil {
		_ccgea = _decbe.creator.NewParagraph("")
	}
	for _, _cdcga := range _caffcf._gbdee.Attr {
		_acedc := _cdcga.Value
		switch _cfgaf := _cdcga.Name.Local; _cfgaf {
		case "\u0066\u006f\u006e\u0074":
			_ccgea.SetFont(_decbe.parseFontAttr(_cfgaf, _acedc))
		case "\u0066o\u006e\u0074\u002d\u0073\u0069\u007ae":
			_ccgea.SetFontSize(_decbe.parseFloatAttr(_cfgaf, _acedc))
		case "\u0074\u0065\u0078\u0074\u002d\u0061\u006c\u0069\u0067\u006e":
			_ccgea.SetTextAlignment(_decbe.parseTextAlignmentAttr(_cfgaf, _acedc))
		case "l\u0069\u006e\u0065\u002d\u0068\u0065\u0069\u0067\u0068\u0074":
			_ccgea.SetLineHeight(_decbe.parseFloatAttr(_cfgaf, _acedc))
		case "e\u006e\u0061\u0062\u006c\u0065\u002d\u0077\u0072\u0061\u0070":
			_ccgea.SetEnableWrap(_decbe.parseBoolAttr(_cfgaf, _acedc))
		case "\u0063\u006f\u006co\u0072":
			_ccgea.SetColor(_decbe.parseColorAttr(_cfgaf, _acedc))
		case "\u0078":
			_ccgea.SetPos(_decbe.parseFloatAttr(_cfgaf, _acedc), _ccgea._cbcca)
		case "\u0079":
			_ccgea.SetPos(_ccgea._bcec, _decbe.parseFloatAttr(_cfgaf, _acedc))
		case "\u0061\u006e\u0067l\u0065":
			_ccgea.SetAngle(_decbe.parseFloatAttr(_cfgaf, _acedc))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_aeaad := _decbe.parseMarginAttr(_cfgaf, _acedc)
			_ccgea.SetMargins(_aeaad.Left, _aeaad.Right, _aeaad.Top, _aeaad.Bottom)
		case "\u006da\u0078\u002d\u006c\u0069\u006e\u0065s":
			_ccgea.SetMaxLines(int(_decbe.parseInt64Attr(_cfgaf, _acedc)))
		default:
			_decbe.nodeLogDebug(_caffcf, "\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065\u0064\u0020pa\u0072\u0061\u0067\u0072\u0061\u0070h\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073`\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u002e", _cfgaf)
		}
	}
	return _ccgea, nil
}

type fontMetrics struct {
	_bacbec float64
	_fbcdg  float64
	_eeeae  float64
	_afebc  float64
}
type componentRenderer interface{ Draw(_cddf Drawable) error }

// SetHorizontalAlignment sets the cell's horizontal alignment of content.
// Can be one of:
// - CellHorizontalAlignmentLeft
// - CellHorizontalAlignmentCenter
// - CellHorizontalAlignmentRight
func (_fafe *TableCell) SetHorizontalAlignment(halign CellHorizontalAlignment) { _fafe._abad = halign }

// BorderColor returns the border color of the rectangle.
func (_bbeg *Rectangle) BorderColor() Color { return _bbeg._ecce }

// NewSubchapter creates a new child chapter with the specified title.
func (_aac *Chapter) NewSubchapter(title string) *Chapter {
	_ddg := _eabad(_aac._ddd._fggb)
	_ddg.FontSize = 14
	_aac._bcbd++
	_dbg := _facd(_aac, _aac._bdde, _aac._ggf, title, _aac._bcbd, _ddg)
	_aac.Add(_dbg)
	return _dbg
}
func (_cbcde *templateProcessor) parseHorizontalAlignmentAttr(_aebag, _efbeeb string) HorizontalAlignment {
	_ca.Log.Debug("\u0050\u0061\u0072\u0073\u0069n\u0067\u0020\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u0020a\u006c\u0069\u0067\u006e\u006d\u0065\u006e\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029.", _aebag, _efbeeb)
	_ecdag := map[string]HorizontalAlignment{"\u006c\u0065\u0066\u0074": HorizontalAlignmentLeft, "\u0063\u0065\u006e\u0074\u0065\u0072": HorizontalAlignmentCenter, "\u0072\u0069\u0067h\u0074": HorizontalAlignmentRight}[_efbeeb]
	return _ecdag
}

// Add adds a new line with the default style to the table of contents.
func (_dcgd *TOC) Add(number, title, page string, level uint) *TOCLine {
	_bdba := _dcgd.AddLine(_bebddd(TextChunk{Text: number, Style: _dcgd._febed}, TextChunk{Text: title, Style: _dcgd._cbgbd}, TextChunk{Text: page, Style: _dcgd._abgad}, level, _dcgd._cceeb))
	if _bdba == nil {
		return nil
	}
	_ceedg := &_dcgd._dfbae
	_bdba.SetMargins(_ceedg.Left, _ceedg.Right, _ceedg.Top, _ceedg.Bottom)
	_bdba.SetLevelOffset(_dcgd._ebgfa)
	_bdba.Separator.Text = _dcgd._feaba
	_bdba.Separator.Style = _dcgd._effdg
	return _bdba
}
func (_gged *templateProcessor) nodeLogDebug(_bgdg *templateNode, _ggaf string, _cffb ...interface{}) {
	_ca.Log.Debug(_gged.getNodeErrorLocation(_bgdg, _ggaf, _cffb...))
}
func _gbcbd(_bgcab *templateProcessor, _dcef *templateNode) (interface{}, error) {
	return _bgcab.parseRectangle(_dcef)
}

// CreateFrontPage sets a function to generate a front Page.
func (_fabe *Creator) CreateFrontPage(genFrontPageFunc func(_gadd FrontpageFunctionArgs)) {
	_fabe._fbd = genFrontPageFunc
}
func _cbag() *PageBreak { return &PageBreak{} }

type marginDrawable interface {
	VectorDrawable
	GetMargins() (float64, float64, float64, float64)
}

// NewList creates a new list.
func (_egac *Creator) NewList() *List { return _edbe(_egac.NewTextStyle()) }
func _cgcb(_ecega *templateProcessor, _fbaga *templateNode) (interface{}, error) {
	return _ecega.parseChapter(_fbaga)
}

// AddShadingResource adds shading dictionary inside the resources dictionary.
func (_dbdf *RadialShading) AddShadingResource(block *Block) (_eefb _fe.PdfObjectName, _daaed error) {
	_aggg := 1
	_eefb = _fe.PdfObjectName("\u0053\u0068" + _a.Itoa(_aggg))
	for block._ge.HasShadingByName(_eefb) {
		_aggg++
		_eefb = _fe.PdfObjectName("\u0053\u0068" + _a.Itoa(_aggg))
	}
	if _ffab := block._ge.SetShadingByName(_eefb, _dbdf.shadingModel().ToPdfObject()); _ffab != nil {
		return "", _ffab
	}
	return _eefb, nil
}

// InfoLines returns all the rows in the invoice information table as
// description-value cell pairs.
func (_cegb *Invoice) InfoLines() [][2]*InvoiceCell {
	_gdbf := [][2]*InvoiceCell{_cegb._geeb, _cegb._ddbe, _cegb._fdfc}
	return append(_gdbf, _cegb._cbga...)
}
func _ecfaf(_edggb *_ggc.PdfFont, _egedcf float64) *fontMetrics {
	_dggdd := &fontMetrics{}
	if _edggb == nil {
		_ca.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0069s\u0020\u006e\u0069\u006c")
		return _dggdd
	}
	_faag, _egfgd := _edggb.GetFontDescriptor()
	if _egfgd != nil {
		_ca.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0067\u0065t\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063ri\u0070\u0074\u006fr\u003a \u0025\u0076", _egfgd)
		return _dggdd
	}
	if _dggdd._bacbec, _egfgd = _faag.GetCapHeight(); _egfgd != nil {
		_ca.Log.Trace("\u0057\u0041\u0052\u004e\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0067\u0065\u0074\u0020f\u006f\u006e\u0074\u0020\u0063\u0061\u0070\u0020\u0068\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _egfgd)
	}
	if int(_dggdd._bacbec) <= 0 {
		_ca.Log.Trace("\u0057\u0041\u0052\u004e\u003a\u0020\u0043\u0061p\u0020\u0048\u0065ig\u0068\u0074\u0020\u006e\u006f\u0074 \u0061\u0076\u0061\u0069\u006c\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065\u0074t\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u00310\u0030\u0030")
		_dggdd._bacbec = 1000
	}
	_dggdd._bacbec *= _egedcf / 1000.0
	if _dggdd._fbcdg, _egfgd = _faag.GetXHeight(); _egfgd != nil {
		_ca.Log.Trace("\u0057\u0041R\u004e\u003a\u0020\u0055n\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0078\u002d\u0068\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _egfgd)
	}
	_dggdd._fbcdg *= _egedcf / 1000.0
	if _dggdd._eeeae, _egfgd = _faag.GetAscent(); _egfgd != nil {
		_ca.Log.Trace("W\u0041\u0052\u004e\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0066\u006fn\u0074\u0020\u0061\u0073\u0063\u0065\u006e\u0074\u003a\u0020%\u0076", _egfgd)
	}
	_dggdd._eeeae *= _egedcf / 1000.0
	if _dggdd._afebc, _egfgd = _faag.GetDescent(); _egfgd != nil {
		_ca.Log.Trace("\u0057\u0041RN\u003a\u0020\u0055n\u0061\u0062\u006c\u0065 to\u0020ge\u0074\u0020\u0066\u006f\u006e\u0074\u0020de\u0073\u0063\u0065\u006e\u0074\u003a\u0020%\u0076", _egfgd)
	}
	_dggdd._afebc *= _egedcf / 1000.0
	return _dggdd
}

// AddInfo is used to append a piece of invoice information in the template
// information table.
func (_bdfc *Invoice) AddInfo(description, value string) (*InvoiceCell, *InvoiceCell) {
	_bgdf := [2]*InvoiceCell{_bdfc.newCell(description, _bdfc._fcba), _bdfc.newCell(value, _bdfc._fcba)}
	_bdfc._cbga = append(_bdfc._cbga, _bgdf)
	return _bgdf[0], _bgdf[1]
}

// Scale scales the rectangle dimensions by the specified factors.
func (_agcf *Rectangle) Scale(xFactor, yFactor float64) {
	_agcf._gfad = xFactor * _agcf._gfad
	_agcf._fefg = yFactor * _agcf._fefg
}
func (_aaae *RadialShading) shadingModel() *_ggc.PdfShadingType3 {
	_bgcf, _ebfb, _gbe := _aaae._fcacc._dedad.ToRGB()
	var _fagb _fc.Point
	switch _aaae._dedge {
	case AnchorBottomLeft:
		_fagb = _fc.Point{X: _aaae._bbccc.Llx, Y: _aaae._bbccc.Lly}
	case AnchorBottomRight:
		_fagb = _fc.Point{X: _aaae._bbccc.Urx, Y: _aaae._bbccc.Ury - _aaae._bbccc.Height()}
	case AnchorTopLeft:
		_fagb = _fc.Point{X: _aaae._bbccc.Llx, Y: _aaae._bbccc.Lly + _aaae._bbccc.Height()}
	case AnchorTopRight:
		_fagb = _fc.Point{X: _aaae._bbccc.Urx, Y: _aaae._bbccc.Ury}
	case AnchorLeft:
		_fagb = _fc.Point{X: _aaae._bbccc.Llx, Y: _aaae._bbccc.Lly + _aaae._bbccc.Height()/2}
	case AnchorTop:
		_fagb = _fc.Point{X: _aaae._bbccc.Llx + _aaae._bbccc.Width()/2, Y: _aaae._bbccc.Ury}
	case AnchorRight:
		_fagb = _fc.Point{X: _aaae._bbccc.Urx, Y: _aaae._bbccc.Lly + _aaae._bbccc.Height()/2}
	case AnchorBottom:
		_fagb = _fc.Point{X: _aaae._bbccc.Urx + _aaae._bbccc.Width()/2, Y: _aaae._bbccc.Lly}
	default:
		_fagb = _fc.NewPoint(_aaae._bbccc.Llx+_aaae._bbccc.Width()/2, _aaae._bbccc.Lly+_aaae._bbccc.Height()/2)
	}
	_adce := _aaae._efdd
	_aafe := _aaae._cagbac
	_bfeb := _fagb.X + _aaae._cafd
	_fefe := _fagb.Y + _aaae._eagf
	if _adce == -1.0 {
		_adce = 0.0
	}
	if _aafe == -1.0 {
		var _gaaea []float64
		_accfd := _b.Pow(_bfeb-_aaae._bbccc.Llx, 2) + _b.Pow(_fefe-_aaae._bbccc.Lly, 2)
		_gaaea = append(_gaaea, _b.Abs(_accfd))
		_degd := _b.Pow(_bfeb-_aaae._bbccc.Llx, 2) + _b.Pow(_aaae._bbccc.Lly+_aaae._bbccc.Height()-_fefe, 2)
		_gaaea = append(_gaaea, _b.Abs(_degd))
		_ggggf := _b.Pow(_aaae._bbccc.Urx-_bfeb, 2) + _b.Pow(_fefe-_aaae._bbccc.Ury-_aaae._bbccc.Height(), 2)
		_gaaea = append(_gaaea, _b.Abs(_ggggf))
		_aacb := _b.Pow(_aaae._bbccc.Urx-_bfeb, 2) + _b.Pow(_aaae._bbccc.Ury-_fefe, 2)
		_gaaea = append(_gaaea, _b.Abs(_aacb))
		_f.Slice(_gaaea, func(_eggdb, _edcd int) bool { return _eggdb > _edcd })
		_aafe = _b.Sqrt(_gaaea[0])
	}
	_egaae := &_ggc.PdfRectangle{Llx: _bfeb - _aafe, Lly: _fefe - _aafe, Urx: _bfeb + _aafe, Ury: _fefe + _aafe}
	_gabc := _ggc.NewPdfShadingType3()
	_gabc.PdfShading.ShadingType = _fe.MakeInteger(3)
	_gabc.PdfShading.ColorSpace = _ggc.NewPdfColorspaceDeviceRGB()
	_gabc.PdfShading.Background = _fe.MakeArrayFromFloats([]float64{_bgcf, _ebfb, _gbe})
	_gabc.PdfShading.BBox = _egaae
	_gabc.PdfShading.AntiAlias = _fe.MakeBool(_aaae._fcacc._daff)
	_gabc.Coords = _fe.MakeArrayFromFloats([]float64{_bfeb, _fefe, _adce, _bfeb, _fefe, _aafe})
	_gabc.Domain = _fe.MakeArrayFromFloats([]float64{0.0, 1.0})
	_gabc.Extend = _fe.MakeArray(_fe.MakeBool(_aaae._fcacc._fgdd[0]), _fe.MakeBool(_aaae._fcacc._fgdd[1]))
	_gabc.Function = _aaae._fcacc.generatePdfFunctions()
	return _gabc
}

// NewBlock creates a new Block with specified width and height.
func NewBlock(width float64, height float64) *Block {
	_cagb := &Block{}
	_cagb._cad = &_bdb.ContentStreamOperations{}
	_cagb._ge = _ggc.NewPdfPageResources()
	_cagb._ecd = width
	_cagb._gfe = height
	return _cagb
}

// Style returns the style of the line.
func (_gagf *Line) Style() _fc.LineStyle { return _gagf._ceac }

// Draw draws the drawable d on the block.
// Note that the drawable must not wrap, i.e. only return one block. Otherwise an error is returned.
func (_bfb *Block) Draw(d Drawable) error {
	_bdd := DrawContext{}
	_bdd.Width = _bfb._ecd
	_bdd.Height = _bfb._gfe
	_bdd.PageWidth = _bfb._ecd
	_bdd.PageHeight = _bfb._gfe
	_bdd.X = 0
	_bdd.Y = 0
	_dcb, _, _efe := d.GeneratePageBlocks(_bdd)
	if _efe != nil {
		return _efe
	}
	if len(_dcb) != 1 {
		return ErrContentNotFit
	}
	for _, _aag := range _dcb {
		if _cdfa := _bfb.mergeBlocks(_aag); _cdfa != nil {
			return _cdfa
		}
	}
	return nil
}

var _acaac = map[string]*templateTag{"\u0070a\u0072\u0061\u0067\u0072\u0061\u0070h": {_bdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": {}}, _bcddf: _cfgb}, "\u0074\u0065\u0078\u0074\u002d\u0063\u0068\u0075\u006e\u006b": {_bdcdf: map[string]struct{}{"\u0070a\u0072\u0061\u0067\u0072\u0061\u0070h": {}}, _bcddf: _gfbgg}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {_bdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": {}}, _bcddf: _gecbe}, "\u0074\u0061\u0062l\u0065": {_bdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": {}}, _bcddf: _dedafb}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {_bdcdf: map[string]struct{}{"\u0074\u0061\u0062l\u0065": {}}, _bcddf: _gcba}, "\u006c\u0069\u006e\u0065": {_bdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bcddf: _gedea}, "\u0072e\u0063\u0074\u0061\u006e\u0067\u006ce": {_bdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bcddf: _gbcbd}, "\u0065l\u006c\u0069\u0070\u0073\u0065": {_bdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bcddf: _dbaec}, "\u0069\u006d\u0061g\u0065": {_bdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": {}}, _bcddf: _dbba}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {_bdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bcddf: _cgcb}, "\u0063h\u0061p\u0074\u0065\u0072\u002d\u0068\u0065\u0061\u0064\u0069\u006e\u0067": {_bdcdf: map[string]struct{}{"\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bcddf: _ggdeb}, "\u0063\u0068\u0061r\u0074": {_bdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bcddf: _fgfc}, "\u0070\u0061\u0067\u0065\u002d\u0062\u0072\u0065\u0061\u006b": {_bdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bcddf: _bfcfd}, "\u0062\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064": {_bdcdf: map[string]struct{}{"\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}}, _bcddf: _fdbcg}, "\u006c\u0069\u0073\u0074": {_bdcdf: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": {}}, _bcddf: _cebcg}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": {_bdcdf: map[string]struct{}{"\u006c\u0069\u0073\u0074": {}}, _bcddf: _gfeg}, "l\u0069\u0073\u0074\u002d\u006d\u0061\u0072\u006b\u0065\u0072": {_bdcdf: map[string]struct{}{"\u006c\u0069\u0073\u0074": {}, "\u006ci\u0073\u0074\u002d\u0069\u0074\u0065m": {}}, _bcddf: _eafef}}

// RotateDeg rotates the current active page by angle degrees.  An error is returned on failure,
// which can be if there is no currently active page, or the angleDeg is not a multiple of 90 degrees.
func (_bgae *Creator) RotateDeg(angleDeg int64) error {
	_aea := _bgae.getActivePage()
	if _aea == nil {
		_ca.Log.Debug("F\u0061\u0069\u006c\u0020\u0074\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0065\u003a\u0020\u006e\u006f\u0020p\u0061\u0067\u0065\u0020\u0063\u0075\u0072\u0072\u0065\u006etl\u0079\u0020\u0061c\u0074i\u0076\u0065")
		return _fa.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	if angleDeg%90 != 0 {
		_ca.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067e\u0020\u0072\u006f\u0074\u0061\u0074\u0069on\u0020\u0061\u006e\u0067l\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006dul\u0074\u0069p\u006c\u0065\u0020\u006f\u0066\u0020\u0039\u0030")
		return _fa.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	var _fdga int64
	if _aea.Rotate != nil {
		_fdga = *(_aea.Rotate)
	}
	_fdga += angleDeg
	_aea.Rotate = &_fdga
	return nil
}

// Width returns the width of the rectangle.
// NOTE: the returned value does not include the border width of the rectangle.
func (_gcbc *Rectangle) Width() float64 { return _gcbc._gfad }
func (_aefdd *Paragraph) getMaxLineWidth() float64 {
	if _aefdd._cbcf == nil || len(_aefdd._cbcf) == 0 {
		_aefdd.wrapText()
	}
	var _cgafa float64
	for _, _eccd := range _aefdd._cbcf {
		_bbbe := _aefdd.getTextLineWidth(_eccd)
		if _bbbe > _cgafa {
			_cgafa = _bbbe
		}
	}
	return _cgafa
}

// VectorDrawable is a Drawable with a specified width and height.
type VectorDrawable interface {
	Drawable

	// Width returns the width of the Drawable.
	Width() float64

	// Height returns the height of the Drawable.
	Height() float64
}

// Width returns the width of the Paragraph.
func (_caedf *StyledParagraph) Width() float64 {
	if _caedf._gfbb && int(_caedf._gcadd) > 0 {
		return _caedf._gcadd
	}
	return _caedf.getTextWidth() / 1000.0
}

var PPI float64 = 72

func (_eaf *Block) mergeBlocks(_bca *Block) error {
	_cafcf := _ddf(_eaf._cad, _eaf._ge, _bca._cad, _bca._ge)
	if _cafcf != nil {
		return _cafcf
	}
	for _, _dag := range _bca._ga {
		_eaf.AddAnnotation(_dag)
	}
	return nil
}

// Title returns the title of the invoice.
func (_dbbg *Invoice) Title() string { return _dbbg._bfbbd }
func (_bbbg *TableCell) height(_eaffa float64) float64 {
	var _adbf float64
	switch _becg := _bbbg._efbbe.(type) {
	case *Paragraph:
		if _becg._bebg {
			_becg.SetWidth(_eaffa - _bbbg._bbgcg - _becg._fgcbf.Left - _becg._fgcbf.Right)
		}
		_adbf = _becg.Height() + _becg._fgcbf.Top + _becg._fgcbf.Bottom + 0.5*_becg._fcbfa*_becg._dacae
	case *StyledParagraph:
		if _becg._gfbb {
			_becg.SetWidth(_eaffa - _bbbg._bbgcg - _becg._fbgbc.Left - _becg._fbgbc.Right)
		}
		_adbf = _becg.Height() + _becg._fbgbc.Top + _becg._fbgbc.Bottom + 0.5*_becg.getTextHeight()
	case *Image:
		_becg.applyFitMode(_eaffa - _bbbg._bbgcg)
		_adbf = _becg.Height() + _becg._cbea.Top + _becg._cbea.Bottom
	case *Table:
		_becg.updateRowHeights(_eaffa - _bbbg._bbgcg - _becg._gfbcc.Left - _becg._gfbcc.Right)
		_adbf = _becg.Height() + _becg._gfbcc.Top + _becg._gfbcc.Bottom
	case *List:
		_adbf = _becg.ctxHeight(_eaffa-_bbbg._bbgcg) + _becg._fed.Top + _becg._fed.Bottom
	case *Division:
		_adbf = _becg.ctxHeight(_eaffa-_bbbg._bbgcg) + _becg._debb.Top + _becg._debb.Bottom + _becg._agda.Top + _becg._agda.Bottom
	case *Chart:
		_adbf = _becg.Height() + _becg._gcff.Top + _becg._gcff.Bottom
	case *Rectangle:
		_becg.applyFitMode(_eaffa - _bbbg._bbgcg)
		_adbf = _becg.Height() + _becg._ebfce.Top + _becg._ebfce.Bottom + _becg._decec
	case *Ellipse:
		_becg.applyFitMode(_eaffa - _bbbg._bbgcg)
		_adbf = _becg.Height() + _becg._cadg.Top + _becg._cadg.Bottom
	case *Line:
		_adbf = _becg.Height() + _becg._bccc.Top + _becg._bccc.Bottom
	}
	return _adbf
}

// Rows returns the total number of rows the table has.
func (_aegg *Table) Rows() int { return _aegg._fgbfga }

// Indent returns the left offset of the list when nested into another list.
func (_eccac *List) Indent() float64 { return _eccac._egab }

// GetMargins returns the Paragraph's margins: left, right, top, bottom.
func (_fagg *StyledParagraph) GetMargins() (float64, float64, float64, float64) {
	return _fagg._fbgbc.Left, _fagg._fbgbc.Right, _fagg._fbgbc.Top, _fagg._fbgbc.Bottom
}
func (_dbaa *pageTransformations) transformBlock(_dfef *Block) {
	if _dbaa._dafc != nil {
		_dfef.transform(*_dbaa._dafc)
	}
}

// SetLevelOffset sets the amount of space an indentation level occupies.
func (_gaggb *TOCLine) SetLevelOffset(levelOffset float64) {
	_gaggb._aeeeg = levelOffset
	_gaggb._ecde._fbgbc.Left = _gaggb._adbfe + float64(_gaggb._daebd-1)*_gaggb._aeeeg
}
func (_bfa *Block) transform(_ege _bd.Matrix) {
	_cdg := _bdb.NewContentCreator().Add_cm(_ege[0], _ege[1], _ege[3], _ege[4], _ege[6], _ege[7]).Operations()
	*_bfa._cad = append(*_cdg, *_bfa._cad...)
	_bfa._cad.WrapIfNeeded()
}

// BorderWidth returns the border width of the rectangle.
func (_cdfac *Rectangle) BorderWidth() float64 { return _cdfac._decec }

// PageBreak represents a page break for a chapter.
type PageBreak struct{}

func _cece(_abbg float64, _aebff float64) float64 { return _b.Round(_abbg/_aebff) * _aebff }

// SetTOC sets the table of content component of the creator.
// This method should be used when building a custom table of contents.
func (_caed *Creator) SetTOC(toc *TOC) {
	if toc == nil {
		return
	}
	_caed._effa = toc
}
func (_fcadd *LinearShading) shadingModel() *_ggc.PdfShadingType2 {
	_aefa := _fc.NewPoint(_fcadd._cbgf.Llx+_fcadd._cbgf.Width()/2, _fcadd._cbgf.Lly+_fcadd._cbgf.Height()/2)
	_bded := _fc.NewPoint(_fcadd._cbgf.Llx, _fcadd._cbgf.Lly+_fcadd._cbgf.Height()/2).Add(-_aefa.X, -_aefa.Y).Rotate(_fcadd._fcfe).Add(_aefa.X, _aefa.Y)
	_bded = _fc.NewPoint(_b.Max(_b.Min(_bded.X, _fcadd._cbgf.Urx), _fcadd._cbgf.Llx), _b.Max(_b.Min(_bded.Y, _fcadd._cbgf.Ury), _fcadd._cbgf.Lly))
	_eabgf := _fc.NewPoint(_fcadd._cbgf.Urx, _fcadd._cbgf.Lly+_fcadd._cbgf.Height()/2).Add(-_aefa.X, -_aefa.Y).Rotate(_fcadd._fcfe).Add(_aefa.X, _aefa.Y)
	_eabgf = _fc.NewPoint(_b.Min(_b.Max(_eabgf.X, _fcadd._cbgf.Llx), _fcadd._cbgf.Urx), _b.Min(_b.Max(_eabgf.Y, _fcadd._cbgf.Lly), _fcadd._cbgf.Ury))
	_edcag := _ggc.NewPdfShadingType2()
	_edcag.PdfShading.ShadingType = _fe.MakeInteger(2)
	_edcag.PdfShading.ColorSpace = _ggc.NewPdfColorspaceDeviceRGB()
	_edcag.PdfShading.AntiAlias = _fe.MakeBool(_fcadd._aecc._daff)
	_edcag.Coords = _fe.MakeArrayFromFloats([]float64{_bded.X, _bded.Y, _eabgf.X, _eabgf.Y})
	_edcag.Extend = _fe.MakeArray(_fe.MakeBool(_fcadd._aecc._fgdd[0]), _fe.MakeBool(_fcadd._aecc._fgdd[1]))
	_edcag.Function = _fcadd._aecc.generatePdfFunctions()
	return _edcag
}

// AddShadingResource adds shading dictionary inside the resources dictionary.
func (_afca *LinearShading) AddShadingResource(block *Block) (_ebffgg _fe.PdfObjectName, _geaa error) {
	_fede := 1
	_ebffgg = _fe.PdfObjectName("\u0053\u0068" + _a.Itoa(_fede))
	for block._ge.HasShadingByName(_ebffgg) {
		_fede++
		_ebffgg = _fe.PdfObjectName("\u0053\u0068" + _a.Itoa(_fede))
	}
	if _dedaf := block._ge.SetShadingByName(_ebffgg, _afca.shadingModel().ToPdfObject()); _dedaf != nil {
		return "", _dedaf
	}
	return _ebffgg, nil
}

// Width returns Image's document width.
func (_feaa *Image) Width() float64 { return _feaa._cagba }

// AppendColumn appends a column to the line items table.
func (_eaade *Invoice) AppendColumn(description string) *InvoiceCell {
	_dgad := _eaade.NewColumn(description)
	_eaade._eag = append(_eaade._eag, _dgad)
	return _dgad
}

// Sections returns the custom content sections of the invoice as
// title-content pairs.
func (_cccb *Invoice) Sections() [][2]string { return _cccb._agdc }

// Color returns the color of the line.
func (_gadc *Line) Color() Color { return _gadc._cdgfa }

// SetBorderRadius sets the radius of the rectangle corners.
func (_fbfa *Rectangle) SetBorderRadius(topLeft, topRight, bottomLeft, bottomRight float64) {
	_fbfa._gagef = topLeft
	_fbfa._efaga = topRight
	_fbfa._fbcea = bottomLeft
	_fbfa._fdcc = bottomRight
}
func (_fddd *TemplateOptions) init() {
	if _fddd.SubtemplateMap == nil {
		_fddd.SubtemplateMap = map[string]_ae.Reader{}
	}
	if _fddd.FontMap == nil {
		_fddd.FontMap = map[string]*_ggc.PdfFont{}
	}
	if _fddd.ImageMap == nil {
		_fddd.ImageMap = map[string]*_ggc.Image{}
	}
	if _fddd.ColorMap == nil {
		_fddd.ColorMap = map[string]Color{}
	}
	if _fddd.ChartMap == nil {
		_fddd.ChartMap = map[string]_gg.ChartRenderable{}
	}
}

// Height returns the height of the ellipse.
func (_gacg *Ellipse) Height() float64 { return _gacg._dabc }
func _bbcaf(_cbfa, _bage, _fdeaa string, _gcdde uint, _ebggd TextStyle) *TOCLine {
	return _bebddd(TextChunk{Text: _cbfa, Style: _ebggd}, TextChunk{Text: _bage, Style: _ebggd}, TextChunk{Text: _fdeaa, Style: _ebggd}, _gcdde, _ebggd)
}

// NewCell makes a new cell and inserts it into the table at the current position.
func (_eedba *Table) NewCell() *TableCell { return _eedba.MultiCell(1, 1) }

// AddPatternResource adds pattern dictionary inside the resources dictionary.
func (_cbecb *LinearShading) AddPatternResource(block *Block) (_cfebfg _fe.PdfObjectName, _bbgb error) {
	_gbdbe := 1
	_ggfcf := _fe.PdfObjectName("\u0050" + _a.Itoa(_gbdbe))
	for block._ge.HasPatternByName(_ggfcf) {
		_gbdbe++
		_ggfcf = _fe.PdfObjectName("\u0050" + _a.Itoa(_gbdbe))
	}
	if _caae := block._ge.SetPatternByName(_ggfcf, _cbecb.ToPdfShadingPattern().ToPdfObject()); _caae != nil {
		return "", _caae
	}
	return _ggfcf, nil
}

// DueDate returns the invoice due date description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_cgedf *Invoice) DueDate() (*InvoiceCell, *InvoiceCell) {
	return _cgedf._fdfc[0], _cgedf._fdfc[1]
}
func _fgfc(_afdfg *templateProcessor, _aebba *templateNode) (interface{}, error) {
	return _afdfg.parseChart(_aebba)
}

// SetWidthTop sets border width for top.
func (_gfcg *border) SetWidthTop(bw float64) { _gfcg._cdgaa = bw }

// SetRowPosition sets cell row position.
func (_gafc *TableCell) SetRowPosition(row int) { _gafc._deded = row }

// Context returns the current drawing context.
func (_bfef *Creator) Context() DrawContext { return _bfef._eacd }

// TableCell defines a table cell which can contain a Drawable as content.
type TableCell struct {
	_afbg          Color
	_fcagf         _fc.LineStyle
	_ffdeb         CellBorderStyle
	_egde          Color
	_egbd          float64
	_cbafe         CellBorderStyle
	_dagd          Color
	_egaeb         float64
	_dfbfd         CellBorderStyle
	_aeec          Color
	_ddbdf         float64
	_aaddc         CellBorderStyle
	_afga          Color
	_caffa         float64
	_deded, _eafbd int
	_ffgdb         int
	_abfda         int
	_efbbe         VectorDrawable
	_abad          CellHorizontalAlignment
	_fddca         CellVerticalAlignment
	_bbgcg         float64
	_ddggg         *Table
}

// SetNumber sets the number of the invoice.
func (_edfa *Invoice) SetNumber(number string) (*InvoiceCell, *InvoiceCell) {
	_edfa._geeb[1].Value = number
	return _edfa._geeb[0], _edfa._geeb[1]
}

// Flip flips the active page on the specified axes.
// If `flipH` is true, the page is flipped horizontally. Similarly, if `flipV`
// is true, the page is flipped vertically. If both are true, the page is
// flipped both horizontally and vertically.
// NOTE: the flip transformations are applied when the creator is finalized,
// which is at write time in most cases.
func (_cgf *Creator) Flip(flipH, flipV bool) error {
	_cdea := _cgf.getActivePage()
	if _cdea == nil {
		return _fa.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	_fcgg, _ebae := _cgf._afbc[_cdea]
	if !_ebae {
		_fcgg = &pageTransformations{}
		_cgf._afbc[_cdea] = _fcgg
	}
	_fcgg._gfag = flipH
	_fcgg._fcaa = flipV
	return nil
}

var (
	_cccd  = _fag.MustCompile("\u0028[\u005cw\u002d\u005d\u002b\u0029\u005c(\u0027\u0028.\u002b\u0029\u0027\u005c\u0029")
	_ddead = _fa.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0063\u0072\u0065a\u0074\u006f\u0072\u0020\u0069\u006e\u0073t\u0061\u006e\u0063\u0065")
	_gfaba = _fa.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074e\u0020p\u0061\u0072\u0065\u006e\u0074\u0020\u006eo\u0064\u0065")
	_dfgbg = _fa.New("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020c\u0068\u0069\u006cd\u0020n\u006f\u0064\u0065")
	_gggfb = _fa.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0074\u0065\u006d\u0070l\u0061t\u0065 \u0072\u0065\u0073\u006f\u0075\u0072\u0063e")
)

// GetMargins returns the Block's margins: left, right, top, bottom.
func (_caf *Block) GetMargins() (float64, float64, float64, float64) {
	return _caf._bc.Left, _caf._bc.Right, _caf._bc.Top, _caf._bc.Bottom
}
func (_eddbc *TextStyle) horizontalScale() float64 { return _eddbc.HorizontalScaling / 100 }
func (_bcadd *templateProcessor) nodeLogError(_feaca *templateNode, _faedf string, _gafd ...interface{}) {
	_ca.Log.Error(_bcadd.getNodeErrorLocation(_feaca, _faedf, _gafd...))
}
func (_adegg *templateProcessor) parseFitModeAttr(_abcec, _bcacb string) FitMode {
	_ca.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u0069\u0074\u0020\u006do\u0064\u0065\u0020\u0061\u0074\u0074r\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c \u0025\u0073\u0029\u002e", _abcec, _bcacb)
	_dfagb := map[string]FitMode{"\u006e\u006f\u006e\u0065": FitModeNone, "\u0066\u0069\u006c\u006c\u002d\u0077\u0069\u0064\u0074\u0068": FitModeFillWidth}[_bcacb]
	return _dfagb
}

var (
	ColorBlack  = ColorRGBFromArithmetic(0, 0, 0)
	ColorWhite  = ColorRGBFromArithmetic(1, 1, 1)
	ColorRed    = ColorRGBFromArithmetic(1, 0, 0)
	ColorGreen  = ColorRGBFromArithmetic(0, 1, 0)
	ColorBlue   = ColorRGBFromArithmetic(0, 0, 1)
	ColorYellow = ColorRGBFromArithmetic(1, 1, 0)
)

// NewEllipse creates a new ellipse with the center at (`xc`, `yc`),
// having the specified width and height.
// NOTE: In relative positioning mode, `xc` and `yc` are calculated using the
// current context. Furthermore, when the fit mode is set to fill the available
// space, the ellipse is scaled so that it occupies the entire context width
// while maintaining the original aspect ratio.
func (_fcadf *Creator) NewEllipse(xc, yc, width, height float64) *Ellipse {
	return _dccb(xc, yc, width, height)
}

// SetNoteHeadingStyle sets the style properties used to render the heading
// of the invoice note sections.
func (_gbcd *Invoice) SetNoteHeadingStyle(style TextStyle) { _gbcd._daag = style }

// SetBorderColor sets border color of the rectangle.
func (_fddc *Rectangle) SetBorderColor(col Color) { _fddc._ecce = col }

// SetLineHeight sets the line height (1.0 default).
func (_gcgc *Paragraph) SetLineHeight(lineheight float64) { _gcgc._dacae = lineheight }

// SetWidth sets line width.
func (_fcade *Curve) SetWidth(width float64) { _fcade._bceaa = width }

// GetCoords returns coordinates of border.
func (_afd *border) GetCoords() (float64, float64) { return _afd._agf, _afd._gea }

// SetFillColor sets the fill color.
func (_baadb *CurvePolygon) SetFillColor(color Color) {
	_baadb._fdge = color
	_baadb._ecae.FillColor = _dbac(color)
}

// CellBorderStyle defines the table cell's border style.
type CellBorderStyle int

func _eafef(_dbea *templateProcessor, _abcfd *templateNode) (interface{}, error) {
	return _dbea.parseListMarker(_abcfd)
}

// Marker returns the marker used for the list items.
// The marker instance can be used the change the text and the style
// of newly added list items.
func (_cgda *List) Marker() *TextChunk { return &_cgda._ebgba }

// NewDivision returns a new Division container component.
func (_bcfb *Creator) NewDivision() *Division { return _ffcc() }

// GetHeading returns the chapter heading paragraph. Used to give access to address style: font, sizing etc.
func (_baee *Chapter) GetHeading() *Paragraph { return _baee._ddd }

// SetHeight sets the Image's document height to specified h.
func (_afeb *Image) SetHeight(h float64) { _afeb._efef = h }

const (
	CellVerticalAlignmentTop CellVerticalAlignment = iota
	CellVerticalAlignmentMiddle
	CellVerticalAlignmentBottom
)

func (_eagg *Invoice) generateLineBlocks(_bgff DrawContext) ([]*Block, DrawContext, error) {
	_cffab := _gdec(len(_eagg._eag))
	_cffab.SetMargins(0, 0, 25, 0)
	for _, _gddc := range _eagg._eag {
		_eece := _egdc(_gddc.TextStyle)
		_eece.SetMargins(0, 0, 1, 0)
		_eece.Append(_gddc.Value)
		_abff := _cffab.NewCell()
		_abff.SetHorizontalAlignment(_gddc.Alignment)
		_abff.SetBackgroundColor(_gddc.BackgroundColor)
		_eagg.setCellBorder(_abff, _gddc)
		_abff.SetContent(_eece)
	}
	for _, _ccgf := range _eagg._fbac {
		for _, _acbaa := range _ccgf {
			_gbbf := _egdc(_acbaa.TextStyle)
			_gbbf.SetMargins(0, 0, 3, 2)
			_gbbf.Append(_acbaa.Value)
			_ffee := _cffab.NewCell()
			_ffee.SetHorizontalAlignment(_acbaa.Alignment)
			_ffee.SetBackgroundColor(_acbaa.BackgroundColor)
			_eagg.setCellBorder(_ffee, _acbaa)
			_ffee.SetContent(_gbbf)
		}
	}
	return _cffab.GeneratePageBlocks(_bgff)
}

// Height returns the height of the rectangle.
// NOTE: the returned value does not include the border width of the rectangle.
func (_fbee *Rectangle) Height() float64 { return _fbee._fefg }

// Wrap wraps the text of the chunk into lines based on its style and the
// specified width.
func (_aeaf *TextChunk) Wrap(width float64) ([]string, error) {
	if int(width) <= 0 {
		return []string{_aeaf.Text}, nil
	}
	var _cagaf []string
	var _cbfee []rune
	var _eefaf float64
	var _dggcc []float64
	_ffbdd := _aeaf.Style
	_edbba := _ddfgf(_aeaf.Text)
	for _, _dbfc := range _aeaf.Text {
		if _dbfc == '\u000A' {
			_bdbcc := _baecf(string(_cbfee), _edbba)
			_cagaf = append(_cagaf, _dc.TrimRightFunc(_bdbcc, _cd.IsSpace)+string(_dbfc))
			_cbfee = nil
			_eefaf = 0
			_dggcc = nil
			continue
		}
		_aada := _dbfc == ' '
		_fcfdf, _ggbg := _ffbdd.Font.GetRuneMetrics(_dbfc)
		if !_ggbg {
			_ca.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006det\u0072i\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064!\u0020\u0072\u0075\u006e\u0065\u003d\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0020\u0066o\u006e\u0074\u003d\u0025\u0073\u0020\u0025\u0023\u0071", _dbfc, _dbfc, _ffbdd.Font.BaseFont(), _ffbdd.Font.Subtype())
			_ca.Log.Trace("\u0046o\u006e\u0074\u003a\u0020\u0025\u0023v", _ffbdd.Font)
			_ca.Log.Trace("\u0045\u006e\u0063o\u0064\u0065\u0072\u003a\u0020\u0025\u0023\u0076", _ffbdd.Font.Encoder())
			return nil, _fa.New("\u0067\u006c\u0079\u0070\u0068\u0020\u0063\u0068\u0061\u0072\u0020m\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
		}
		_dbfbgf := _ffbdd.FontSize * _fcfdf.Wx
		_ffbba := _dbfbgf
		if !_aada {
			_ffbba = _dbfbgf + _ffbdd.CharSpacing*1000.0
		}
		if _eefaf+_dbfbgf > width*1000.0 {
			_fbcgf := -1
			if !_aada {
				for _aaddb := len(_cbfee) - 1; _aaddb >= 0; _aaddb-- {
					if _cbfee[_aaddb] == ' ' {
						_fbcgf = _aaddb
						break
					}
				}
			}
			_aacc := string(_cbfee)
			if _fbcgf > 0 {
				_aacc = string(_cbfee[0 : _fbcgf+1])
				_cbfee = append(_cbfee[_fbcgf+1:], _dbfc)
				_dggcc = append(_dggcc[_fbcgf+1:], _ffbba)
				_eefaf = 0
				for _, _acaef := range _dggcc {
					_eefaf += _acaef
				}
			} else {
				if _aada {
					_cbfee = []rune{}
					_dggcc = []float64{}
					_eefaf = 0
				} else {
					_cbfee = []rune{_dbfc}
					_dggcc = []float64{_ffbba}
					_eefaf = _ffbba
				}
			}
			_aacc = _baecf(_aacc, _edbba)
			_cagaf = append(_cagaf, _dc.TrimRightFunc(_aacc, _cd.IsSpace))
		} else {
			_cbfee = append(_cbfee, _dbfc)
			_eefaf += _ffbba
			_dggcc = append(_dggcc, _ffbba)
		}
	}
	if len(_cbfee) > 0 {
		_bffc := string(_cbfee)
		_bffc = _baecf(_bffc, _edbba)
		_cagaf = append(_cagaf, _bffc)
	}
	return _cagaf, nil
}
func _cagg(_bacc [][]_fc.CubicBezierCurve) *CurvePolygon {
	return &CurvePolygon{_ecae: &_fc.CurvePolygon{Rings: _bacc}, _gfage: 1.0, _fgcc: 1.0}
}

// Fit fits the chunk into the specified bounding box, cropping off the
// remainder in a new chunk, if it exceeds the specified dimensions.
// NOTE: The method assumes a line height of 1.0. In order to account for other
// line height values, the passed in height must be divided by the line height:
// height = height / lineHeight
func (_gcga *TextChunk) Fit(width, height float64) (*TextChunk, error) {
	_bcdc, _ggfg := _gcga.Wrap(width)
	if _ggfg != nil {
		return nil, _ggfg
	}
	_dbgcf := int(height / _gcga.Style.FontSize)
	if _dbgcf >= len(_bcdc) {
		return nil, nil
	}
	_bdcc := "\u000a"
	_gcga.Text = _dc.Replace(_dc.Join(_bcdc[:_dbgcf], "\u0020"), _bdcc+"\u0020", _bdcc, -1)
	_gcfgd := _dc.Replace(_dc.Join(_bcdc[_dbgcf:], "\u0020"), _bdcc+"\u0020", _bdcc, -1)
	return NewTextChunk(_gcfgd, _gcga.Style), nil
}

// SetOpacity sets opacity for Image.
func (_eaebb *Image) SetOpacity(opacity float64) { _eaebb._feef = opacity }

// SetText replaces all the text of the paragraph with the specified one.
func (_deeg *StyledParagraph) SetText(text string) *TextChunk {
	_deeg.Reset()
	return _deeg.Append(text)
}

// SetBorderColor sets the cell's border color.
func (_bgdag *TableCell) SetBorderColor(col Color) {
	_bgdag._egde = col
	_bgdag._dagd = col
	_bgdag._aeec = col
	_bgdag._afga = col
}
func _gfbgg(_aggfa *templateProcessor, _ccfbc *templateNode) (interface{}, error) {
	return _aggfa.parseTextChunk(_ccfbc, nil)
}
func (_fafb *templateProcessor) parseMarginAttr(_bcdfb, _eadea string) Margins {
	_ca.Log.Debug("\u0050\u0061r\u0073\u0069\u006e\u0067 \u006d\u0061r\u0067\u0069\u006e\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c \u0025\u0073\u0029\u002e", _bcdfb, _eadea)
	_adaace := Margins{}
	switch _dgffb := _dc.Fields(_eadea); len(_dgffb) {
	case 1:
		_adaace.Top, _ = _a.ParseFloat(_dgffb[0], 64)
		_adaace.Bottom = _adaace.Top
		_adaace.Left = _adaace.Top
		_adaace.Right = _adaace.Top
	case 2:
		_adaace.Top, _ = _a.ParseFloat(_dgffb[0], 64)
		_adaace.Bottom = _adaace.Top
		_adaace.Left, _ = _a.ParseFloat(_dgffb[1], 64)
		_adaace.Right = _adaace.Left
	case 3:
		_adaace.Top, _ = _a.ParseFloat(_dgffb[0], 64)
		_adaace.Left, _ = _a.ParseFloat(_dgffb[1], 64)
		_adaace.Right = _adaace.Left
		_adaace.Bottom, _ = _a.ParseFloat(_dgffb[2], 64)
	case 4:
		_adaace.Top, _ = _a.ParseFloat(_dgffb[0], 64)
		_adaace.Right, _ = _a.ParseFloat(_dgffb[1], 64)
		_adaace.Bottom, _ = _a.ParseFloat(_dgffb[2], 64)
		_adaace.Left, _ = _a.ParseFloat(_dgffb[3], 64)
	}
	return _adaace
}

// Height returns the height of the list.
func (_gggb *List) Height() float64 {
	var _ededb float64
	for _, _daccg := range _gggb._bede {
		_ededb += _daccg.ctxHeight(_gggb.Width())
	}
	return _ededb
}

// ColorRGBFrom8bit creates a Color from 8-bit (0-255) r,g,b values.
// Example:
//
//	red := ColorRGBFrom8Bit(255, 0, 0)
func ColorRGBFrom8bit(r, g, b byte) Color {
	return rgbColor{_efdc: float64(r) / 255.0, _fgbf: float64(g) / 255.0, _acbd: float64(b) / 255.0}
}

// SetSubtotal sets the subtotal of the invoice.
func (_ebbg *Invoice) SetSubtotal(value string) { _ebbg._befc[1].Value = value }
func _dgbfg(_feefb []*ColorPoint) *LinearShading {
	return &LinearShading{_aecc: &shading{_dedad: ColorWhite, _daff: false, _fgdd: []bool{false, false}, _ggcd: _feefb}, _cbgf: &_ggc.PdfRectangle{}}
}
func (_eefea *templateProcessor) parsePageBreak(_eacce *templateNode) (interface{}, error) {
	return _cbag(), nil
}

// AddLine appends a new line to the invoice line items table.
func (_ggdc *Invoice) AddLine(values ...string) []*InvoiceCell {
	_eadd := len(_ggdc._eag)
	var _gaae []*InvoiceCell
	for _fgbaa, _bcgca := range values {
		_gfcge := _ggdc.newCell(_bcgca, _ggdc._degfg)
		if _fgbaa < _eadd {
			_gfcge.Alignment = _ggdc._eag[_fgbaa].Alignment
		}
		_gaae = append(_gaae, _gfcge)
	}
	_ggdc._fbac = append(_ggdc._fbac, _gaae)
	return _gaae
}

// SetWidthBottom sets border width for bottom.
func (_ffd *border) SetWidthBottom(bw float64) { _ffd._gaf = bw }

// SetWidth sets the the Paragraph width. This is essentially the wrapping width,
// i.e. the width the text can extend to prior to wrapping over to next line.
func (_dgfe *StyledParagraph) SetWidth(width float64) { _dgfe._gcadd = width; _dgfe.wrapText() }

// SetLogo sets the logo of the invoice.
func (_ebfc *Invoice) SetLogo(logo *Image) { _ebfc._bfge = logo }
func (_eadge *templateProcessor) parseLineStyleAttr(_fafg, _fbed string) _fc.LineStyle {
	_ca.Log.Debug("\u0050\u0061\u0072\u0073\u0069n\u0067\u0020\u006c\u0069\u006e\u0065\u0020\u0073\u0074\u0079\u006c\u0065\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _fafg, _fbed)
	_fdace := map[string]_fc.LineStyle{"\u0073\u006f\u006ci\u0064": _fc.LineStyleSolid, "\u0064\u0061\u0073\u0068\u0065\u0064": _fc.LineStyleDashed}[_fbed]
	return _fdace
}
func _ddfgf(_ccfea string) bool {
	_cbbg := func(_cgdff rune) bool { return _cgdff == '\u000A' }
	_ebda := _dc.TrimFunc(_ccfea, _cbbg)
	_fefcc := _cag.Paragraph{}
	_, _bcdcc := _fefcc.SetString(_ebda)
	if _bcdcc != nil {
		return true
	}
	_bdeb, _bcdcc := _fefcc.Order()
	if _bcdcc != nil {
		return true
	}
	if _bdeb.NumRuns() < 1 {
		return true
	}
	return _fefcc.IsLeftToRight()
}

// RotatedSize returns the width and height of the rotated block.
func (_dec *Block) RotatedSize() (float64, float64) {
	_, _, _gca, _eba := _gcdfe(_dec._ecd, _dec._gfe, _dec._de)
	return _gca, _eba
}
func (_adebg *templateProcessor) renderNode(_gbca *templateNode) error {
	_dgea := _gbca._caacd
	if _dgea == nil {
		return nil
	}
	_efecd := _gbca._gbdee.Name.Local
	_dacb, _eeaec := _acaac[_efecd]
	if !_eeaec {
		_adebg.nodeLogDebug(_gbca, "I\u006e\u0076\u0061\u006c\u0069\u0064 \u0074\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u002e \u0053\u006b\u0069p\u0070i\u006e\u0067\u002e", _efecd)
		return nil
	}
	var _efefa interface{}
	if _gbca._fbcg != nil && _gbca._fbcg._caacd != nil {
		_adff := _gbca._fbcg._gbdee.Name.Local
		if _, _eeaec = _dacb._bdcdf[_adff]; !_eeaec {
			_adebg.nodeLogDebug(_gbca, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073\u003e \u0069\u0073\u0020no\u0074\u0020\u0061\u0020\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u003c\u0025\u0073\u003e\u0020\u0074a\u0067\u002e", _adff, _efecd)
			return _gfaba
		}
		_efefa = _gbca._fbcg._caacd
	} else {
		_ebede := "\u0063r\u0065\u0061\u0074\u006f\u0072"
		switch _adebg._gcfaeb.(type) {
		case *Block:
			_ebede = "\u0062\u006c\u006fc\u006b"
		}
		if _, _eeaec = _dacb._bdcdf[_ebede]; !_eeaec {
			_adebg.nodeLogDebug(_gbca, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073\u003e \u0069\u0073\u0020no\u0074\u0020\u0061\u0020\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u003c\u0025\u0073\u003e\u0020\u0074a\u0067\u002e", _ebede, _efecd)
			return _gfaba
		}
		_efefa = _adebg._gcfaeb
	}
	switch _efga := _efefa.(type) {
	case componentRenderer:
		_fccbgd, _ebbcb := _dgea.(Drawable)
		if !_ebbcb {
			_adebg.nodeLogError(_gbca, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054\u0029\u0020\u0069s\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0072\u0061\u0077\u0061\u0062\u006c\u0065\u002e", _efecd, _dgea)
			return _dfgbg
		}
		_bgddg := _efga.Draw(_fccbgd)
		if _bgddg != nil {
			return _adebg.nodeError(_gbca, "\u0043\u0061\u006en\u006f\u0074\u0020\u0064r\u0061\u0077\u0073\u0020\u0074\u0061\u0067 \u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054\u0029\u003a\u0020\u0025\u0073\u002e", _efecd, _dgea, _bgddg)
		}
	case *Division:
		switch _agde := _dgea.(type) {
		case *Background:
			_efga.SetBackground(_agde)
		case VectorDrawable:
			_gfgff := _efga.Add(_agde)
			if _gfgff != nil {
				return _adebg.nodeError(_gbca, "\u0043a\u006e\u006eo\u0074\u0020\u0061d\u0064\u0020\u0074\u0061\u0067\u0020\u003c%\u0073\u003e\u0020\u0028\u0025\u0054)\u0020\u0069\u006e\u0074\u006f\u0020\u0061\u0020\u0044\u0069\u0076i\u0073\u0069\u006f\u006e\u003a\u0020\u0025\u0073\u002e", _efecd, _dgea, _gfgff)
			}
		}
	case *TableCell:
		_bdddd, _dabdg := _dgea.(VectorDrawable)
		if !_dabdg {
			_adebg.nodeLogError(_gbca, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054\u0029 \u0069\u0073\u0020\u006e\u006f\u0074 \u0061\u0020\u0076\u0065\u0063\u0074\u006f\u0072\u0020\u0064\u0072\u0061\u0077a\u0062\u006c\u0065\u002e", _efecd, _dgea)
			return _dfgbg
		}
		_fccd := _efga.SetContent(_bdddd)
		if _fccd != nil {
			return _adebg.nodeError(_gbca, "C\u0061\u006e\u006e\u006f\u0074\u0020\u0061\u0064\u0064 \u0074\u0061\u0067\u0020\u003c\u0025\u0073> \u0028\u0025\u0054\u0029 \u0069\u006e\u0074\u006f\u0020\u0061\u0020\u0074\u0061bl\u0065\u0020c\u0065\u006c\u006c\u003a\u0020\u0025\u0073\u002e", _efecd, _dgea, _fccd)
		}
	case *StyledParagraph:
		_aecef, _gbefc := _dgea.(*TextChunk)
		if !_gbefc {
			_adebg.nodeLogError(_gbca, "\u0054\u0061\u0067 <\u0025\u0073\u003e\u0020\u0028\u0025\u0054\u0029\u0020i\u0073 \u006eo\u0074 \u0061\u0020\u0074\u0065\u0078\u0074\u0020\u0063\u0068\u0075\u006e\u006b\u002e", _efecd, _dgea)
			return _dfgbg
		}
		_efga.appendChunk(_aecef)
	case *Chapter:
		switch _eecc := _dgea.(type) {
		case *Chapter:
			return nil
		case *Paragraph:
			if _gbca._gbdee.Name.Local == "\u0063h\u0061p\u0074\u0065\u0072\u002d\u0068\u0065\u0061\u0064\u0069\u006e\u0067" {
				return nil
			}
			_bbac := _efga.Add(_eecc)
			if _bbac != nil {
				return _adebg.nodeError(_gbca, "\u0043a\u006e\u006eo\u0074\u0020\u0061\u0064d\u0020\u0074\u0061g\u0020\u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054) \u0069\u006e\u0074o\u0020\u0061 \u0043\u0068\u0061\u0070\u0074\u0065r\u003a\u0020%\u0073\u002e", _efecd, _dgea, _bbac)
			}
		case Drawable:
			_debca := _efga.Add(_eecc)
			if _debca != nil {
				return _adebg.nodeError(_gbca, "\u0043a\u006e\u006eo\u0074\u0020\u0061\u0064d\u0020\u0074\u0061g\u0020\u003c\u0025\u0073\u003e\u0020\u0028\u0025\u0054) \u0069\u006e\u0074o\u0020\u0061 \u0043\u0068\u0061\u0070\u0074\u0065r\u003a\u0020%\u0073\u002e", _efecd, _dgea, _debca)
			}
		}
	case *List:
		switch _fcfgc := _dgea.(type) {
		case *TextChunk:
		case *listItem:
			_efga._bede = append(_efga._bede, _fcfgc)
		default:
			_adebg.nodeLogError(_gbca, "\u0054\u0061\u0067\u0020\u003c\u0025\u0073>\u0020\u0028\u0025T\u0029\u0020\u0069\u0073 \u006e\u006f\u0074\u0020\u0061\u0020\u006c\u0069\u0073\u0074\u0020\u0069\u0074\u0065\u006d\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _efecd, _dgea)
		}
	case *listItem:
		switch _dgacc := _dgea.(type) {
		case *TextChunk:
		case *StyledParagraph:
			_efga._gdceb = _dgacc
		case *List:
			if _dgacc._ceeg {
				_dgacc._egab = 15
			}
			_efga._gdceb = _dgacc
		case *Image:
			_efga._gdceb = _dgacc
		case *Division:
			_efga._gdceb = _dgacc
		case *Table:
			_efga._gdceb = _dgacc
		default:
			_adebg.nodeLogError(_gbca, "\u0054\u0061\u0067\u0020\u003c%\u0073\u003e\u0020\u0028\u0025\u0054\u0029\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u006c\u0069\u0073\u0074\u002e", _efecd, _dgea)
			return _dfgbg
		}
	}
	return nil
}

// GetMargins returns the Image's margins: left, right, top, bottom.
func (_cfdfa *Image) GetMargins() (float64, float64, float64, float64) {
	return _cfdfa._cbea.Left, _cfdfa._cbea.Right, _cfdfa._cbea.Top, _cfdfa._cbea.Bottom
}
func _agbgc(_ffbf ...interface{}) (map[string]interface{}, error) {
	_fefdd := len(_ffbf)
	if _fefdd%2 != 0 {
		_ca.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u0073\u0020\u0066\u006f\u0072\u0020\u0063\u0072\u0065\u0061\u0074i\u006e\u0067\u0020\u006d\u0061\u0070\u003a\u0020\u0025\u0064\u002e", _fefdd)
		return nil, _fe.ErrRangeError
	}
	_fbdg := map[string]interface{}{}
	for _cdaeb := 0; _cdaeb < _fefdd; _cdaeb += 2 {
		_adda, _edbce := _ffbf[_cdaeb].(string)
		if !_edbce {
			_ca.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006d\u0061\u0070 \u006b\u0065\u0079\u0020t\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u002e\u0020\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u002e", _ffbf[_cdaeb])
			return nil, _fe.ErrTypeError
		}
		_fbdg[_adda] = _ffbf[_cdaeb+1]
	}
	return _fbdg, nil
}

// NoteStyle returns the style properties used to render the content of the
// invoice note sections.
func (_ggcg *Invoice) NoteStyle() TextStyle { return _ggcg._dcfab }

// SetBorderColor sets the border color.
func (_effeg *Polygon) SetBorderColor(color Color) { _effeg._gbda.BorderColor = _dbac(color) }
func (_ffgg *Table) getLastCellFromCol(_ebgg int) (int, *TableCell) {
	for _fdeg := len(_ffgg._cacca) - 1; _fdeg >= 0; _fdeg-- {
		if _ffgg._cacca[_fdeg]._eafbd == _ebgg {
			return _fdeg, _ffgg._cacca[_fdeg]
		}
	}
	return 0, nil
}
func _baecf(_gbbbf string, _bgdcb bool) string {
	_faggbb := _gbbbf
	if _faggbb == "" {
		return ""
	}
	_ffede := _cag.Paragraph{}
	_, _eebag := _ffede.SetString(_gbbbf)
	if _eebag != nil {
		return _faggbb
	}
	_eabcc, _eebag := _ffede.Order()
	if _eebag != nil {
		return _faggbb
	}
	_cgeaf := _eabcc.NumRuns()
	_bbbgb := make([]string, _cgeaf)
	for _ccafe := 0; _ccafe < _eabcc.NumRuns(); _ccafe++ {
		_dgdg := _eabcc.Run(_ccafe)
		_cdcf := _dgdg.String()
		if _dgdg.Direction() == _cag.RightToLeft {
			_cdcf = _cag.ReverseString(_cdcf)
		}
		if _bgdcb {
			_bbbgb[_ccafe] = _cdcf
		} else {
			_bbbgb[_cgeaf-1] = _cdcf
		}
		_cgeaf--
	}
	if len(_bbbgb) != _eabcc.NumRuns() {
		return _gbbbf
	}
	_faggbb = _dc.Join(_bbbgb, "")
	return _faggbb
}

// SetWidthLeft sets border width for left.
func (_gbb *border) SetWidthLeft(bw float64) { _gbb._gcd = bw }

// Link returns link information for this line.
func (_gcada *TOCLine) Link() (_fbca int64, _dfecg, _bdaaf float64) {
	return _gcada._eefcc, _gcada._efdgc, _gcada._cafda
}
func (_ggfbf *TOCLine) getLineLink() *_ggc.PdfAnnotation {
	if _ggfbf._eefcc <= 0 {
		return nil
	}
	return _bfefg(_ggfbf._eefcc-1, _ggfbf._efdgc, _ggfbf._cafda, 0)
}

// SetExtends specifies whether ot extend the shading beyond the starting and ending points.
//
// Text extends is set to `[]bool{false, false}` by default.
func (_adgaa *LinearShading) SetExtends(start bool, end bool) { _adgaa._aecc.SetExtends(start, end) }

// Polygon represents a polygon shape.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Polygon struct {
	_gbda  *_fc.Polygon
	_befea float64
	_dffbb float64
	_ddcb  Color
}

const (
	CellBorderSideLeft CellBorderSide = iota
	CellBorderSideRight
	CellBorderSideTop
	CellBorderSideBottom
	CellBorderSideAll
)

// GetCoords returns the (x1, y1), (x2, y2) points defining the Line.
func (_eecec *Line) GetCoords() (float64, float64, float64, float64) {
	return _eecec._daed, _eecec._gfaa, _eecec._eddg, _eecec._eabb
}
func (_gdgg *InvoiceAddress) fmtLine(_bbcc, _bacbg string, _bgda bool) string {
	if _bgda {
		_bacbg = ""
	}
	return _df.Sprintf("\u0025\u0073\u0025s\u000a", _bacbg, _bbcc)
}

// SetPositioning sets the positioning of the line (absolute or relative).
func (_fdca *Line) SetPositioning(positioning Positioning) { _fdca._fgef = positioning }

// SetLineSeparator sets the separator for all new lines of the table of contents.
func (_bgcc *TOC) SetLineSeparator(separator string) { _bgcc._feaba = separator }

// Height returns the height of the Paragraph. The height is calculated based on the input text and how it is wrapped
// within the container. Does not include Margins.
func (_eafe *StyledParagraph) Height() float64 {
	_eafe.wrapText()
	var _baeec float64
	for _, _ccdb := range _eafe._aabba {
		var _cbgag float64
		for _, _ffgec := range _ccdb {
			_gdbe := _eafe._babcf * _ffgec.Style.FontSize
			if _gdbe > _cbgag {
				_cbgag = _gdbe
			}
		}
		_baeec += _cbgag
	}
	return _baeec
}
func _dccb(_ddae, _gfca, _bbea, _ebfe float64) *Ellipse {
	return &Ellipse{_eebe: _ddae, _beb: _gfca, _eded: _bbea, _dabc: _ebfe, _acgd: PositionAbsolute, _ddbf: 1.0, _fegd: ColorBlack, _fgcca: 1.0, _dadc: 1.0}
}
func (_dcga *templateProcessor) parseTextRenderingModeAttr(_afcag, _ggea string) TextRenderingMode {
	_ca.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0074\u0065\u0078\u0074\u0020\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u006d\u006f\u0064e\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a \u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _afcag, _ggea)
	_ccdbg := map[string]TextRenderingMode{"\u0066\u0069\u006c\u006c": TextRenderingModeFill, "\u0073\u0074\u0072\u006f\u006b\u0065": TextRenderingModeStroke, "f\u0069\u006c\u006c\u002d\u0073\u0074\u0072\u006f\u006b\u0065": TextRenderingModeFillStroke, "\u0069n\u0076\u0069\u0073\u0069\u0062\u006ce": TextRenderingModeInvisible, "\u0066i\u006c\u006c\u002d\u0063\u006c\u0069p": TextRenderingModeFillClip, "s\u0074\u0072\u006f\u006b\u0065\u002d\u0063\u006c\u0069\u0070": TextRenderingModeStrokeClip, "\u0066\u0069l\u006c\u002d\u0073t\u0072\u006f\u006b\u0065\u002d\u0063\u006c\u0069\u0070": TextRenderingModeFillStrokeClip, "\u0063\u006c\u0069\u0070": TextRenderingModeClip}[_ggea]
	return _ccdbg
}

// SetStyle sets the style of the line (solid or dashed).
func (_deda *Line) SetStyle(style _fc.LineStyle) { _deda._ceac = style }
func (_ccfcg *shading) generatePdfFunctions() []_ggc.PdfFunction {
	if len(_ccfcg._ggcd) == 0 {
		return nil
	} else if len(_ccfcg._ggcd) <= 2 {
		_bdbff, _bdgac, _fcaaf := _ccfcg._ggcd[0]._afea.ToRGB()
		_faed, _ddag, _bgfcg := _ccfcg._ggcd[len(_ccfcg._ggcd)-1]._afea.ToRGB()
		return []_ggc.PdfFunction{&_ggc.PdfFunctionType2{Domain: []float64{0.0, 1.0}, Range: []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}, N: 1, C0: []float64{_bdbff, _bdgac, _fcaaf}, C1: []float64{_faed, _ddag, _bgfcg}}}
	} else {
		_faefdd := []_ggc.PdfFunction{}
		_beefd := []float64{}
		for _bagd := 0; _bagd < len(_ccfcg._ggcd)-1; _bagd++ {
			_fege, _dcfb, _gdcbf := _ccfcg._ggcd[_bagd]._afea.ToRGB()
			_eefe, _gedf, _gefe := _ccfcg._ggcd[_bagd+1]._afea.ToRGB()
			_bfdbe := &_ggc.PdfFunctionType2{Domain: []float64{0.0, 1.0}, Range: []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}, N: 1, C0: []float64{_fege, _dcfb, _gdcbf}, C1: []float64{_eefe, _gedf, _gefe}}
			_faefdd = append(_faefdd, _bfdbe)
			if _bagd > 0 {
				_beefd = append(_beefd, _ccfcg._ggcd[_bagd]._dbcfd)
			}
		}
		_daffe := []float64{}
		for range _faefdd {
			_daffe = append(_daffe, []float64{0.0, 1.0}...)
		}
		return []_ggc.PdfFunction{&_ggc.PdfFunctionType3{Domain: []float64{0.0, 1.0}, Range: []float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0}, Functions: _faefdd, Bounds: _beefd, Encode: _daffe}}
	}
}

// Length calculates and returns the length of the line.
func (_cfbae *Line) Length() float64 {
	return _b.Sqrt(_b.Pow(_cfbae._eddg-_cfbae._daed, 2.0) + _b.Pow(_cfbae._eabb-_cfbae._gfaa, 2.0))
}
func (_eg *Block) setOpacity(_aeg float64, _ag float64) (string, error) {
	if (_aeg < 0 || _aeg >= 1.0) && (_ag < 0 || _ag >= 1.0) {
		return "", nil
	}
	_gge := 0
	_bgc := _df.Sprintf("\u0047\u0053\u0025\u0064", _gge)
	for _eg._ge.HasExtGState(_fe.PdfObjectName(_bgc)) {
		_gge++
		_bgc = _df.Sprintf("\u0047\u0053\u0025\u0064", _gge)
	}
	_fd := _fe.MakeDict()
	if _aeg >= 0 && _aeg < 1.0 {
		_fd.Set("\u0063\u0061", _fe.MakeFloat(_aeg))
	}
	if _ag >= 0 && _ag < 1.0 {
		_fd.Set("\u0043\u0041", _fe.MakeFloat(_ag))
	}
	_cdf := _eg._ge.AddExtGState(_fe.PdfObjectName(_bgc), _fd)
	if _cdf != nil {
		return "", _cdf
	}
	return _bgc, nil
}

// BorderColor returns the border color of the ellipse.
func (_geefb *Ellipse) BorderColor() Color { return _geefb._fegd }

// Append adds a new text chunk to the paragraph.
func (_aaba *StyledParagraph) Append(text string) *TextChunk {
	_dddec := NewTextChunk(text, _aaba._egcg)
	return _aaba.appendChunk(_dddec)
}

// SetMargins sets the margins of the paragraph.
func (_bbbf *List) SetMargins(left, right, top, bottom float64) {
	_bbbf._fed.Left = left
	_bbbf._fed.Right = right
	_bbbf._fed.Top = top
	_bbbf._fed.Bottom = bottom
}

// MoveDown moves the drawing context down by relative displacement dy (negative goes up).
func (_bgfd *Creator) MoveDown(dy float64) { _bgfd._eacd.Y += dy }

// SetSellerAddress sets the seller address of the invoice.
func (_cecbg *Invoice) SetSellerAddress(address *InvoiceAddress) { _cecbg._cfbbb = address }

// SetBorderColor sets the border color for the path.
func (_bgac *FilledCurve) SetBorderColor(color Color) { _bgac._eadf = color }

// ToPdfShadingPattern generates a new model.PdfShadingPatternType3 object.
func (_dffe *RadialShading) ToPdfShadingPattern() *_ggc.PdfShadingPatternType3 {
	_cffc, _ebeg, _bfbe := _dffe._fcacc._dedad.ToRGB()
	_gace := _dffe.shadingModel()
	_gace.PdfShading.Background = _fe.MakeArrayFromFloats([]float64{_cffc, _ebeg, _bfbe})
	_fdff := _ggc.NewPdfShadingPatternType3()
	_fdff.Shading = _gace
	return _fdff
}
func (_fdfd *Table) updateRowHeights(_cddag float64) {
	for _, _ageb := range _fdfd._cacca {
		_aeeac := _ageb.width(_fdfd._abbbf, _cddag)
		_dbde := _ageb.height(_aeeac)
		_aadc := _fdfd._begb[_ageb._deded+_ageb._ffgdb-2]
		if _ageb._ffgdb > 1 {
			_fgccd := 0.0
			_fccbe := _fdfd._begb[_ageb._deded-1 : (_ageb._deded + _ageb._ffgdb - 1)]
			for _, _cebd := range _fccbe {
				_fgccd += _cebd
			}
			if _dbde <= _fgccd {
				continue
			}
		}
		if _dbde > _aadc {
			_bbaa := _dbde / float64(_ageb._ffgdb)
			if _bbaa > _aadc {
				for _ggac := 1; _ggac <= _ageb._ffgdb; _ggac++ {
					if _bbaa > _fdfd._begb[_ageb._deded+_ggac-2] {
						_fdfd._begb[_ageb._deded+_ggac-2] = _bbaa
					}
				}
			}
		}
	}
}
func (_bafb *templateProcessor) parseChapter(_edbdc *templateNode) (interface{}, error) {
	_cfbe := _bafb.creator.NewChapter
	if _edbdc._fbcg != nil {
		if _fffed, _feda := _edbdc._fbcg._caacd.(*Chapter); _feda {
			_cfbe = _fffed.NewSubchapter
		}
	}
	_gfdeg := _cfbe("")
	for _, _fgedd := range _edbdc._gbdee.Attr {
		_fcfbb := _fgedd.Value
		switch _edbf := _fgedd.Name.Local; _edbf {
		case "\u0073\u0068\u006f\u0077\u002d\u006e\u0075\u006d\u0062e\u0072\u0069\u006e\u0067":
			_gfdeg.SetShowNumbering(_bafb.parseBoolAttr(_edbf, _fcfbb))
		case "\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u002d\u0069n\u002d\u0074\u006f\u0063":
			_gfdeg.SetIncludeInTOC(_bafb.parseBoolAttr(_edbf, _fcfbb))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_baaaa := _bafb.parseMarginAttr(_edbf, _fcfbb)
			_gfdeg.SetMargins(_baaaa.Left, _baaaa.Right, _baaaa.Top, _baaaa.Bottom)
		default:
			_bafb.nodeLogDebug(_edbdc, "\u0055\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0068\u0061\u0070\u0074\u0065\u0072\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _edbf)
		}
	}
	return _gfdeg, nil
}

type border struct {
	_agf      float64
	_gea      float64
	_dcf      float64
	_aebc     float64
	_cde      Color
	_fcg      Color
	_gcd      float64
	_agg      Color
	_gaf      float64
	_add      Color
	_ebb      float64
	_feg      Color
	_cdgaa    float64
	LineStyle _fc.LineStyle
	_acba     CellBorderStyle
	_dabg     CellBorderStyle
	_deccc    CellBorderStyle
	_cffe     CellBorderStyle
}

// Scale sets the scale ratio with `X` factor and `Y` factor for the graphic svg.
func (_bfcf *GraphicSVG) Scale(xFactor, yFactor float64) {
	_bfcf._dgaf.Width = xFactor * _bfcf._dgaf.Width
	_bfcf._dgaf.Height = yFactor * _bfcf._dgaf.Height
	_bfcf._dgaf.SetScaling(xFactor, yFactor)
}

// NewCell returns a new invoice table cell.
func (_bcbe *Invoice) NewCell(value string) *InvoiceCell {
	return _bcbe.newCell(value, _bcbe.NewCellProps())
}

// TitleStyle returns the style properties used to render the invoice title.
func (_bagc *Invoice) TitleStyle() TextStyle { return _bagc._eefc }

type templateNode struct {
	_caacd interface{}
	_gbdee _e.StartElement
	_fbcg  *templateNode
	_febd  int
	_fgfag int
	_gcfag int64
}

// LineWidth returns the width of the line.
func (_ecead *Line) LineWidth() float64 { return _ecead._adf }

// SetColumns overwrites any columns in the line items table. This should be
// called before AddLine.
func (_eeba *Invoice) SetColumns(cols []*InvoiceCell) { _eeba._eag = cols }

// SetBoundingBox set gradient color bounding box where the gradient would be rendered.
func (_gfcc *RadialShading) SetBoundingBox(x, y, width, height float64) {
	_gfcc._bbccc = &_ggc.PdfRectangle{Llx: x, Lly: y, Urx: x + width, Ury: y + height}
}
func (_dbagg *Image) applyFitMode(_ffff float64) {
	_ffff -= _dbagg._cbea.Left + _dbagg._cbea.Right
	switch _dbagg._geaf {
	case FitModeFillWidth:
		_dbagg.ScaleToWidth(_ffff)
	}
}

// NoteHeadingStyle returns the style properties used to render the heading of
// the invoice note sections.
func (_agb *Invoice) NoteHeadingStyle() TextStyle { return _agb._daag }

// SetStyleBottom sets border style for bottom side.
func (_gccc *border) SetStyleBottom(style CellBorderStyle) { _gccc._cffe = style }

// SetFillColor sets the fill color for the path.
func (_cbgc *FilledCurve) SetFillColor(color Color) { _cbgc._ddddf = color }
func (_caddc *templateProcessor) parseLine(_fecfg *templateNode) (interface{}, error) {
	_eabfb := _caddc.creator.NewLine(0, 0, 0, 0)
	for _, _fdfde := range _fecfg._gbdee.Attr {
		_aaaec := _fdfde.Value
		switch _aeede := _fdfde.Name.Local; _aeede {
		case "\u0078\u0031":
			_eabfb._daed = _caddc.parseFloatAttr(_aeede, _aaaec)
		case "\u0079\u0031":
			_eabfb._gfaa = _caddc.parseFloatAttr(_aeede, _aaaec)
		case "\u0078\u0032":
			_eabfb._eddg = _caddc.parseFloatAttr(_aeede, _aaaec)
		case "\u0079\u0032":
			_eabfb._eabb = _caddc.parseFloatAttr(_aeede, _aaaec)
		case "\u0074h\u0069\u0063\u006b\u006e\u0065\u0073s":
			_eabfb.SetLineWidth(_caddc.parseFloatAttr(_aeede, _aaaec))
		case "\u0063\u006f\u006co\u0072":
			_eabfb.SetColor(_caddc.parseColorAttr(_aeede, _aaaec))
		case "\u0073\u0074\u0079l\u0065":
			_eabfb.SetStyle(_caddc.parseLineStyleAttr(_aeede, _aaaec))
		case "\u0064\u0061\u0073\u0068\u002d\u0061\u0072\u0072\u0061\u0079":
			_eabfb.SetDashPattern(_caddc.parseInt64Array(_aeede, _aaaec), _eabfb._bgfed)
		case "\u0064\u0061\u0073\u0068\u002d\u0070\u0068\u0061\u0073\u0065":
			_eabfb.SetDashPattern(_eabfb._fgbfg, _caddc.parseInt64Attr(_aeede, _aaaec))
		case "\u006fp\u0061\u0063\u0069\u0074\u0079":
			_eabfb.SetOpacity(_caddc.parseFloatAttr(_aeede, _aaaec))
		case "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e":
			_eabfb.SetPositioning(_caddc.parsePositioningAttr(_aeede, _aaaec))
		case "\u0066\u0069\u0074\u002d\u006d\u006f\u0064\u0065":
			_eabfb.SetFitMode(_caddc.parseFitModeAttr(_aeede, _aaaec))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_cbgdb := _caddc.parseMarginAttr(_aeede, _aaaec)
			_eabfb.SetMargins(_cbgdb.Left, _cbgdb.Right, _cbgdb.Top, _cbgdb.Bottom)
		default:
			_caddc.nodeLogDebug(_fecfg, "\u0055\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u006c\u0069\u006e\u0065 \u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _aeede)
		}
	}
	return _eabfb, nil
}
func (_gfaf *templateProcessor) parseLinkAttr(_geec, _aefad string) *_ggc.PdfAnnotation {
	_aefad = _dc.TrimSpace(_aefad)
	if _dc.HasPrefix(_aefad, "\u0075\u0072\u006c(\u0027") && _dc.HasSuffix(_aefad, "\u0027\u0029") && len(_aefad) > 7 {
		return _gfcae(_aefad[5 : len(_aefad)-2])
	}
	if _dc.HasPrefix(_aefad, "\u0070\u0061\u0067e\u0028") && _dc.HasSuffix(_aefad, "\u0029") && len(_aefad) > 6 {
		var (
			_ggada  error
			_gegc   int64
			_agfa   float64
			_babedc float64
			_acgad  = 1.0
			_adadc  = _dc.Split(_aefad[5:len(_aefad)-1], "\u002c")
		)
		_gegc, _ggada = _a.ParseInt(_dc.TrimSpace(_adadc[0]), 10, 64)
		if _ggada != nil {
			_ca.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0070\u0061\u0067\u0065\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u003a\u0020\u0025\u0076", _ggada)
			return nil
		}
		if len(_adadc) >= 2 {
			_agfa, _ggada = _a.ParseFloat(_dc.TrimSpace(_adadc[1]), 64)
			if _ggada != nil {
				_ca.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0070\u0061\u0072\u0073\u0069\u006eg\u0020\u0058\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u003a\u0020\u0025\u0076", _ggada)
				return nil
			}
		}
		if len(_adadc) >= 3 {
			_babedc, _ggada = _a.ParseFloat(_dc.TrimSpace(_adadc[2]), 64)
			if _ggada != nil {
				_ca.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0070\u0061\u0072\u0073\u0069\u006eg\u0020\u0059\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u003a\u0020\u0025\u0076", _ggada)
				return nil
			}
		}
		if len(_adadc) >= 4 {
			_acgad, _ggada = _a.ParseFloat(_dc.TrimSpace(_adadc[3]), 64)
			if _ggada != nil {
				_ca.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u007a\u006f\u006f\u006d\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u003a\u0020\u0025\u0076", _ggada)
				return nil
			}
		}
		return _bfefg(_gegc-1, _agfa, _babedc, _acgad)
	}
	return nil
}

// AddColorStop add color stop info for rendering gradient color.
func (_eaba *RadialShading) AddColorStop(color Color, point float64) {
	_eaba._fcacc.AddColorStop(color, point)
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

// TextVerticalAlignment controls the vertical position of the text
// in a styled paragraph.
type TextVerticalAlignment int

// SetPos sets the Block's positioning to absolute mode with the specified coordinates.
func (_gba *Block) SetPos(x, y float64) { _gba._ba = PositionAbsolute; _gba._gb = x; _gba._gf = y }

// Height returns Image's document height.
func (_bcggd *Image) Height() float64 { return _bcggd._efef }

// LevelOffset returns the amount of space an indentation level occupies.
func (_aaff *TOCLine) LevelOffset() float64 { return _aaff._aeeeg }

// NewPolyBezierCurve creates a new composite Bezier (polybezier) curve.
func (_ggdf *Creator) NewPolyBezierCurve(curves []_fc.CubicBezierCurve) *PolyBezierCurve {
	return _caaf(curves)
}
func (_efbc *List) ctxHeight(_fgfa float64) float64 {
	_fgfa -= _efbc._egab
	var _bbdcb float64
	for _, _baac := range _efbc._bede {
		_bbdcb += _baac.ctxHeight(_fgfa)
	}
	return _bbdcb
}
func (_fbea *StyledParagraph) getLineMetrics(_gffg int) (_cgaa, _aace, _fbfb float64) {
	if _fbea._aabba == nil || len(_fbea._aabba) == 0 {
		_fbea.wrapText()
	}
	if _gffg < 0 || _gffg > len(_fbea._aabba)-1 {
		_ca.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020p\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u0020\u006c\u0069\u006e\u0065 \u0069\u006e\u0064\u0065\u0078\u0020\u0025\u0064\u002e\u0020\u0052\u0065tu\u0072\u006e\u0069\u006e\u0067\u0020\u0030\u002c\u0020\u0030", _gffg)
		return 0, 0, 0
	}
	_bcgd := _fbea._aabba[_gffg]
	for _, _bcca := range _bcgd {
		_bfede := _ecfaf(_bcca.Style.Font, _bcca.Style.FontSize)
		if _bfede._bacbec > _cgaa {
			_cgaa = _bfede._bacbec
		}
		if _bfede._afebc < _fbfb {
			_fbfb = _bfede._afebc
		}
		if _fbde := _bcca.Style.FontSize; _fbde > _aace {
			_aace = _fbde
		}
	}
	return _cgaa, _aace, _fbfb
}

// SetMargins sets the margins of the graphic svg component.
func (_cfdbf *GraphicSVG) SetMargins(left, right, top, bottom float64) {
	_cfdbf._eaff.Left = left
	_cfdbf._eaff.Right = right
	_cfdbf._eaff.Top = top
	_cfdbf._eaff.Bottom = bottom
}

// SetLineStyle sets the style for all the line components: number, title,
// separator, page. The style is applied only for new lines added to the
// TOC component.
func (_ggbce *TOC) SetLineStyle(style TextStyle) {
	_ggbce.SetLineNumberStyle(style)
	_ggbce.SetLineTitleStyle(style)
	_ggbce.SetLineSeparatorStyle(style)
	_ggbce.SetLinePageStyle(style)
}
func _gedea(_cgccg *templateProcessor, _fegdb *templateNode) (interface{}, error) {
	return _cgccg.parseLine(_fegdb)
}
func _gddf(_fdfaf *_ggc.PdfRectangle, _aecca _bd.Matrix) *_ggc.PdfRectangle {
	var _dccgc _ggc.PdfRectangle
	_dccgc.Llx, _dccgc.Lly = _aecca.Transform(_fdfaf.Llx, _fdfaf.Lly)
	_dccgc.Urx, _dccgc.Ury = _aecca.Transform(_fdfaf.Urx, _fdfaf.Ury)
	_dccgc.Normalize()
	return &_dccgc
}

// ScaleToHeight scales the ellipse to the specified height. The width of
// the ellipse is scaled so that the aspect ratio is maintained.
func (_becd *Ellipse) ScaleToHeight(h float64) {
	_eebbe := _becd._eded / _becd._dabc
	_becd._dabc = h
	_becd._eded = h * _eebbe
}

// TextAlignment options for paragraph.
type TextAlignment int

func (_acgb *Creator) setActivePage(_ccdd *_ggc.PdfPage) { _acgb._bcba = _ccdd }
func (_eefa *Invoice) setCellBorder(_bfga *TableCell, _bdgcc *InvoiceCell) {
	for _, _deefa := range _bdgcc.BorderSides {
		_bfga.SetBorder(_deefa, CellBorderStyleSingle, _bdgcc.BorderWidth)
	}
	_bfga.SetBorderColor(_bdgcc.BorderColor)
}
func (_bdbbe *TableCell) cloneProps(_gfdfd VectorDrawable) *TableCell {
	_cdag := *_bdbbe
	_cdag._efbbe = _gfdfd
	return &_cdag
}

// SetTerms sets the terms and conditions section of the invoice.
func (_gbg *Invoice) SetTerms(title, content string) { _gbg._aacd = [2]string{title, content} }

// Lines returns all the lines the table of contents has.
func (_dbed *TOC) Lines() []*TOCLine { return _dbed._fbeeg }
func (_ggcc *templateProcessor) parseImage(_cfgfd *templateNode) (interface{}, error) {
	var _acfe string
	for _, _fffcb := range _cfgfd._gbdee.Attr {
		_gacge := _fffcb.Value
		switch _ebgbc := _fffcb.Name.Local; _ebgbc {
		case "\u0073\u0072\u0063":
			_acfe = _gacge
		}
	}
	_dcee, _ccec := _ggcc.loadImageFromSrc(_acfe)
	if _ccec != nil {
		return nil, _ccec
	}
	for _, _dedb := range _cfgfd._gbdee.Attr {
		_aafg := _dedb.Value
		switch _gcca := _dedb.Name.Local; _gcca {
		case "\u0061\u006c\u0069g\u006e":
			_dcee.SetHorizontalAlignment(_ggcc.parseHorizontalAlignmentAttr(_gcca, _aafg))
		case "\u006fp\u0061\u0063\u0069\u0074\u0079":
			_dcee.SetOpacity(_ggcc.parseFloatAttr(_gcca, _aafg))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_ceeda := _ggcc.parseMarginAttr(_gcca, _aafg)
			_dcee.SetMargins(_ceeda.Left, _ceeda.Right, _ceeda.Top, _ceeda.Bottom)
		case "\u0066\u0069\u0074\u002d\u006d\u006f\u0064\u0065":
			_dcee.SetFitMode(_ggcc.parseFitModeAttr(_gcca, _aafg))
		case "\u0078":
			_dcee.SetPos(_ggcc.parseFloatAttr(_gcca, _aafg), _dcee._gdd)
		case "\u0079":
			_dcee.SetPos(_dcee._fce, _ggcc.parseFloatAttr(_gcca, _aafg))
		case "\u0077\u0069\u0064t\u0068":
			_dcee.SetWidth(_ggcc.parseFloatAttr(_gcca, _aafg))
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_dcee.SetHeight(_ggcc.parseFloatAttr(_gcca, _aafg))
		case "\u0061\u006e\u0067l\u0065":
			_dcee.SetAngle(_ggcc.parseFloatAttr(_gcca, _aafg))
		case "\u0073\u0072\u0063":
		default:
			_ggcc.nodeLogDebug(_cfgfd, "\u0055n\u0073\u0075p\u0070\u006f\u0072\u0074e\u0064\u0020\u0069m\u0061\u0067\u0065\u0020\u0061\u0074\u0074\u0072\u0069bu\u0074\u0065\u003a \u0060\u0025s\u0060\u002e\u0020\u0053\u006b\u0069p\u0070\u0069n\u0067\u002e", _gcca)
		}
	}
	return _dcee, nil
}
func _ggdeb(_gdda *templateProcessor, _becbf *templateNode) (interface{}, error) {
	return _gdda.parseChapterHeading(_becbf)
}

// New creates a new instance of the PDF Creator.
func New() *Creator {
	_gbbg := &Creator{}
	_gbbg._cec = []*_ggc.PdfPage{}
	_gbbg._gfab = map[*_ggc.PdfPage]*Block{}
	_gbbg._afbc = map[*_ggc.PdfPage]*pageTransformations{}
	_gbbg.SetPageSize(PageSizeLetter)
	_aec := 0.1 * _gbbg._abf
	_gbbg._gcgd.Left = _aec
	_gbbg._gcgd.Right = _aec
	_gbbg._gcgd.Top = _aec
	_gbbg._gcgd.Bottom = _aec
	var _bcgg error
	_gbbg._gade, _bcgg = _ggc.NewStandard14Font(_ggc.HelveticaName)
	if _bcgg != nil {
		_gbbg._gade = _ggc.DefaultFont()
	}
	_gbbg._eacc, _bcgg = _ggc.NewStandard14Font(_ggc.HelveticaBoldName)
	if _bcgg != nil {
		_gbbg._gade = _ggc.DefaultFont()
	}
	_gbbg._effa = _gbbg.NewTOC("\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0065\u006e\u0074\u0073")
	_gbbg.AddOutlines = true
	_gbbg._gaec = _ggc.NewOutline()
	return _gbbg
}
func (_beac *StyledParagraph) getTextWidth() float64 {
	var _bdad float64
	_fdad := len(_beac._cdffa)
	for _dccc, _dcdf := range _beac._cdffa {
		_fgggd := &_dcdf.Style
		_cdgfd := len(_dcdf.Text)
		for _dbbff, _fcfg := range _dcdf.Text {
			if _fcfg == '\u000A' {
				continue
			}
			_beece, _aagc := _fgggd.Font.GetRuneMetrics(_fcfg)
			if !_aagc {
				_ca.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _fcfg)
				return -1
			}
			_bdad += _fgggd.FontSize * _beece.Wx * _fgggd.horizontalScale()
			if _fcfg != ' ' && (_dccc != _fdad-1 || _dbbff != _cdgfd-1) {
				_bdad += _fgggd.CharSpacing * 1000.0
			}
		}
	}
	return _bdad
}

// SetLineSeparatorStyle sets the style for the separator part of all new
// lines of the table of contents.
func (_eaaff *TOC) SetLineSeparatorStyle(style TextStyle) { _eaaff._effdg = style }

// NewChapter creates a new chapter with the specified title as the heading.
func (_ffgfa *Creator) NewChapter(title string) *Chapter {
	_ffgfa._caffe++
	_cecf := _ffgfa.NewTextStyle()
	_cecf.FontSize = 16
	return _facd(nil, _ffgfa._effa, _ffgfa._gaec, title, _ffgfa._caffe, _cecf)
}

// Rectangle defines a rectangle with upper left corner at (x,y) and a specified width and height.  The rectangle
// can have a colored fill and/or border with a specified width.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Rectangle struct {
	_cgedd float64
	_eeff  float64
	_gfad  float64
	_fefg  float64
	_bbgac Positioning
	_bgca  Color
	_geee  float64
	_ecce  Color
	_decec float64
	_ffgd  float64
	_gagef float64
	_efaga float64
	_fbcea float64
	_fdcc  float64
	_ebfce Margins
	_gaef  FitMode
}
type pageTransformations struct {
	_dafc *_bd.Matrix
	_gfag bool
	_fcaa bool
}

// SetEncoder sets the encoding/compression mechanism for the image.
func (_gdca *Image) SetEncoder(encoder _fe.StreamEncoder) { _gdca._ggfc = encoder }

// NewImageFromGoImage creates an Image from a go image.Image data structure.
func (_aece *Creator) NewImageFromGoImage(goimg _c.Image) (*Image, error) { return _ddbg(goimg) }

// Width returns the width of the graphic svg.
func (_deeb *GraphicSVG) Width() float64 { return _deeb._dgaf.Width }
func _cebcg(_bggea *templateProcessor, _geedf *templateNode) (interface{}, error) {
	return _bggea.parseList(_geedf)
}
func _edbe(_gfcaa TextStyle) *List {
	return &List{_ebgba: TextChunk{Text: "\u2022\u0020", Style: _gfcaa}, _egab: 0, _ceeg: true, _bgdd: PositionRelative, _dagb: _gfcaa}
}

// FillOpacity returns the fill opacity of the rectangle (0-1).
func (_acfd *Rectangle) FillOpacity() float64 { return _acfd._geee }
