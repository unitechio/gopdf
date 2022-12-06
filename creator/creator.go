package creator

import (
	_b "bytes"
	_cc "encoding/xml"
	_d "errors"
	_eg "fmt"
	_ba "image"
	_ef "io"
	_a "math"
	_e "os"
	_p "path"
	_pf "path/filepath"
	_ca "regexp"
	_ce "sort"
	_ae "strconv"
	_ed "strings"
	_g "text/template"
	_ee "unicode"

	_df "bitbucket.org/shenghui0779/gopdf/common"
	_eee "bitbucket.org/shenghui0779/gopdf/contentstream"
	_cec "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_fg "bitbucket.org/shenghui0779/gopdf/core"
	_fb "bitbucket.org/shenghui0779/gopdf/internal/integrations/unichart"
	_edf "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_f "bitbucket.org/shenghui0779/gopdf/model"
	_aec "github.com/gorilla/i18n/linebreak"
	_ab "github.com/unidoc/unichart/render"
	_af "golang.org/x/text/unicode/bidi"
)

// SetHorizontalAlignment sets the cell's horizontal alignment of content.
// Can be one of:
// - CellHorizontalAlignmentLeft
// - CellHorizontalAlignmentCenter
// - CellHorizontalAlignmentRight
func (_ggfgc *TableCell) SetHorizontalAlignment(halign CellHorizontalAlignment) {
	_ggfgc._eaeac = halign
}
func (_ecabd *Invoice) drawInformation() *Table {
	_beeb := _ecef(2)
	_gegd := append([][2]*InvoiceCell{_ecabd._edeb, _ecabd._fgcge, _ecabd._debe}, _ecabd._agef...)
	for _, _fbgf := range _gegd {
		_cccfg, _egce := _fbgf[0], _fbgf[1]
		if _egce.Value == "" {
			continue
		}
		_ebff := _beeb.NewCell()
		_ebff.SetBackgroundColor(_cccfg.BackgroundColor)
		_ecabd.setCellBorder(_ebff, _cccfg)
		_bdgda := _ggaa(_cccfg.TextStyle)
		_bdgda.Append(_cccfg.Value)
		_bdgda.SetMargins(0, 0, 2, 1)
		_ebff.SetContent(_bdgda)
		_ebff = _beeb.NewCell()
		_ebff.SetBackgroundColor(_egce.BackgroundColor)
		_ecabd.setCellBorder(_ebff, _egce)
		_bdgda = _ggaa(_egce.TextStyle)
		_bdgda.Append(_egce.Value)
		_bdgda.SetMargins(0, 0, 2, 1)
		_ebff.SetContent(_bdgda)
	}
	return _beeb
}

// SetTextOverflow controls the behavior of paragraph text which
// does not fit in the available space.
func (_fccgg *StyledParagraph) SetTextOverflow(textOverflow TextOverflow) { _fccgg._gcb = textOverflow }

// GetMargins returns the margins of the TOC line: left, right, top, bottom.
func (_ccbad *TOCLine) GetMargins() (float64, float64, float64, float64) {
	_cedf := &_ccbad._face._afdba
	return _ccbad._ecdbc, _cedf.Right, _cedf.Top, _cedf.Bottom
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

// SetEncoder sets the encoding/compression mechanism for the image.
func (_aeed *Image) SetEncoder(encoder _fg.StreamEncoder) { _aeed._facba = encoder }

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
func (_bbbe *Creator) SetPdfWriterAccessFunc(pdfWriterAccessFunc func(_faeff *_f.PdfWriter) error) {
	_bbbe._eege = pdfWriterAccessFunc
}

// CreateFrontPage sets a function to generate a front Page.
func (_dga *Creator) CreateFrontPage(genFrontPageFunc func(_aebce FrontpageFunctionArgs)) {
	_dga._bcec = genFrontPageFunc
}

// FitMode returns the fit mode of the image.
func (_ege *Image) FitMode() FitMode { return _ege._aedff }

type marginDrawable interface {
	VectorDrawable
	GetMargins() (float64, float64, float64, float64)
}

// Marker returns the marker used for the list items.
// The marker instance can be used the change the text and the style
// of newly added list items.
func (_gdbe *List) Marker() *TextChunk { return &_gdbe._fbfe }

// LevelOffset returns the amount of space an indentation level occupies.
func (_ggge *TOCLine) LevelOffset() float64 { return _ggge._aaaa }

// NewRectangle creates a new Rectangle with default parameters
// with left corner at (x,y) and width, height as specified.
func (_acg *Creator) NewRectangle(x, y, width, height float64) *Rectangle {
	return _aeac(x, y, width, height)
}

// WriteToFile writes the Creator output to file specified by path.
func (_gggfa *Creator) WriteToFile(filename string) error {
	abspath, _cagd := _pf.Abs(filename)

	if _cagd != nil {
		return _cagd
	}

	if _cagd = _e.MkdirAll(_p.Dir(abspath), 0775); _cagd != nil {
		return _cagd
	}

	_bgc, _cagd := _e.OpenFile(abspath, _e.O_RDWR|_e.O_CREATE|_e.O_TRUNC, 0775)

	if _cagd != nil {
		return _cagd
	}

	defer _bgc.Close()

	return _gggfa.Write(_bgc)
}

// NewPageBreak create a new page break.
func (_fcafg *Creator) NewPageBreak() *PageBreak { return _acdfe() }

// SetBackgroundColor sets the cell's background color.
func (_cafdda *TableCell) SetBackgroundColor(col Color) { _cafdda._fgddf = col }

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
func (_ebfaf *TableCell) SetContent(vd VectorDrawable) error {
	switch _aaccg := vd.(type) {
	case *Paragraph:
		if _aaccg._bbbeg {
			_aaccg._bbfc = true
		}
		_ebfaf._cbgcb = vd
	case *StyledParagraph:
		if _aaccg._ffbb {
			_aaccg._ecbf = true
		}
		_ebfaf._cbgcb = vd
	case *Image, *Chart, *Table, *Division, *List, *Rectangle, *Ellipse, *Line:
		_ebfaf._cbgcb = vd
	default:
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0063e\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0079p\u0065\u0020\u0025\u0054", vd)
		return _fg.ErrTypeError
	}
	return nil
}
func (_gbbd *templateProcessor) parseTextOverflowAttr(_ffef, _agfg string) TextOverflow {
	_df.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020\u0074e\u0078\u0074\u0020o\u0076\u0065\u0072\u0066\u006c\u006f\u0077\u0020\u0061tt\u0072\u0069\u0062u\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060,\u0020\u0025s\u0029\u002e", _ffef, _agfg)
	_dfced := map[string]TextOverflow{"\u0076i\u0073\u0069\u0062\u006c\u0065": TextOverflowVisible, "\u0068\u0069\u0064\u0064\u0065\u006e": TextOverflowHidden}[_agfg]
	return _dfced
}

// CellBorderSide defines the table cell's border side.
type CellBorderSide int

func (_fbcg *templateProcessor) parseBoolAttr(_feae, _cbeda string) bool {
	_df.Log.Debug("P\u0061\u0072\u0073\u0069\u006e\u0067 \u0062\u006f\u006f\u006c\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _feae, _cbeda)
	_fbcag, _ := _ae.ParseBool(_cbeda)
	return _cbeda == "" || _fbcag
}

// FitMode returns the fit mode of the line.
func (_gffg *Line) FitMode() FitMode { return _gffg._gggcb }

// SetTitle sets the title of the invoice.
func (_fbbd *Invoice) SetTitle(title string) { _fbbd._bgdd = title }

// SetNoteHeadingStyle sets the style properties used to render the heading
// of the invoice note sections.
func (_cegd *Invoice) SetNoteHeadingStyle(style TextStyle) { _cegd._cacg = style }

const (
	TextOverflowVisible TextOverflow = iota
	TextOverflowHidden
)

// SetPos sets the Block's positioning to absolute mode with the specified coordinates.
func (_gd *Block) SetPos(x, y float64) { _gd._ec = PositionAbsolute; _gd._fa = x; _gd._ea = y }
func (_ffdc *templateProcessor) renderNode(_egabf *templateNode) error {
	_efgf := _egabf._cedba
	if _efgf == nil {
		return nil
	}
	_adcf := _egabf._ddgde.Name.Local
	_cebe, _edee := _accge[_adcf]
	if !_edee {
		_df.Log.Debug("\u0049\u006e\u0076\u0061l\u0069\u0064\u0020\u0074\u0061\u0067\u003a\u0020\u0060\u0025s\u0060.\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u002e", _adcf)
		return nil
	}
	var _gafgc interface{}
	if _egabf._dfdd != nil && _egabf._dfdd._cedba != nil {
		_aefe := _egabf._dfdd._ddgde.Name.Local
		if _, _edee = _cebe._cedge[_aefe]; !_edee {
			_df.Log.Debug("\u0060%\u0073\u0060 \u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072e\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020`\u0025\u0073\u0060\u0020\u006e\u006f\u0064\u0065\u002e", _aefe, _adcf)
			return _dbea
		}
		_gafgc = _egabf._dfdd._cedba
	} else {
		_bcecb := "\u0063r\u0065\u0061\u0074\u006f\u0072"
		switch _ffdc._cfeb.(type) {
		case *Block:
			_bcecb = "\u0062\u006c\u006fc\u006b"
		}
		if _, _edee = _cebe._cedge[_bcecb]; !_edee {
			_df.Log.Debug("\u0060%\u0073\u0060 \u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0076\u0061l\u0069\u0064\u0020\u0070\u0061\u0072e\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020`\u0025\u0073\u0060\u0020\u006e\u006f\u0064\u0065\u002e", _bcecb, _adcf)
			return _dbea
		}
		_gafgc = _ffdc._cfeb
	}
	switch _gceb := _gafgc.(type) {
	case componentRenderer:
		_fgfbb, _eagd := _efgf.(Drawable)
		if !_eagd {
			_df.Log.Error("\u0043\u006f\u006d\u0070\u006f\u006ee\u006e\u0074\u0020\u0028\u0025\u0054\u0029\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0061\u0020\u0064\u0072\u0061\u0077a\u0062\u006c\u0065\u002e", _efgf)
			return _gaeg
		}
		return _gceb.Draw(_fgfbb)
	case *Division:
		switch _bgcc := _efgf.(type) {
		case *Background:
			_gceb.SetBackground(_bgcc)
		case VectorDrawable:
			return _gceb.Add(_bgcc)
		}
	case *TableCell:
		_bgge, _fdeaa := _efgf.(VectorDrawable)
		if !_fdeaa {
			_df.Log.Error("C\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u0028\u0025\u0054\u0029\u0020\u0069\u0073\u0020\u006eo\u0074\u0020\u0061\u0020\u0076\u0065\u0063\u0074\u006f\u0072 d\u0072\u0061\u0077a\u0062l\u0065\u002e", _efgf)
			return _gaeg
		}
		return _gceb.SetContent(_bgge)
	case *StyledParagraph:
		_bfbeb, _dfac := _efgf.(*TextChunk)
		if !_dfac {
			_df.Log.Error("C\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0020\u0028\u0025\u0054\u0029 \u0069\u0073\u0020\u006e\u006f\u0074\u0020a\u0020\u0074\u0065\u0078\u0074\u0020\u0063\u0068\u0075\u006ek\u002e", _efgf)
			return _gaeg
		}
		_gceb.appendChunk(_bfbeb)
	case *Chapter:
		switch _dedfc := _efgf.(type) {
		case *Chapter:
			return nil
		case *Paragraph:
			if _egabf._ddgde.Name.Local == "\u0063h\u0061p\u0074\u0065\u0072\u002d\u0068\u0065\u0061\u0064\u0069\u006e\u0067" {
				return nil
			}
			return _gceb.Add(_dedfc)
		case Drawable:
			return _gceb.Add(_dedfc)
		}
	}
	return nil
}

// TitleStyle returns the style properties used to render the invoice title.
func (_ecab *Invoice) TitleStyle() TextStyle { return _ecab._acgbf }

// SetFontSize sets the font size in document units (points).
func (_gafg *Paragraph) SetFontSize(fontSize float64) { _gafg._gced = fontSize }
func (_gfed *Invoice) drawAddress(_fedd *InvoiceAddress) []*StyledParagraph {
	var _aadb []*StyledParagraph
	if _fedd.Heading != "" {
		_bbee := _ggaa(_gfed._eeea)
		_bbee.SetMargins(0, 0, 0, 7)
		_bbee.Append(_fedd.Heading)
		_aadb = append(_aadb, _bbee)
	}
	_adgd := _ggaa(_gfed._ddf)
	_adgd.SetLineHeight(1.2)
	_geff := _fedd.Separator
	if _geff == "" {
		_geff = _gfed._ceb
	}
	_abde := _fedd.City
	if _fedd.State != "" {
		if _abde != "" {
			_abde += _geff
		}
		_abde += _fedd.State
	}
	if _fedd.Zip != "" {
		if _abde != "" {
			_abde += _geff
		}
		_abde += _fedd.Zip
	}
	if _fedd.Name != "" {
		_adgd.Append(_fedd.Name + "\u000a")
	}
	if _fedd.Street != "" {
		_adgd.Append(_fedd.Street + "\u000a")
	}
	if _fedd.Street2 != "" {
		_adgd.Append(_fedd.Street2 + "\u000a")
	}
	if _abde != "" {
		_adgd.Append(_abde + "\u000a")
	}
	if _fedd.Country != "" {
		_adgd.Append(_fedd.Country + "\u000a")
	}
	_dcgd := _ggaa(_gfed._ddf)
	_dcgd.SetLineHeight(1.2)
	_dcgd.SetMargins(0, 0, 7, 0)
	if _fedd.Phone != "" {
		_dcgd.Append(_fedd.fmtLine(_fedd.Phone, "\u0050h\u006f\u006e\u0065\u003a\u0020", _fedd.HidePhoneLabel))
	}
	if _fedd.Email != "" {
		_dcgd.Append(_fedd.fmtLine(_fedd.Email, "\u0045m\u0061\u0069\u006c\u003a\u0020", _fedd.HideEmailLabel))
	}
	_aadb = append(_aadb, _adgd, _dcgd)
	return _aadb
}

// NewBlock creates a new Block with specified width and height.
func NewBlock(width float64, height float64) *Block {
	_fc := &Block{}
	_fc._ga = &_eee.ContentStreamOperations{}
	_fc._eb = _f.NewPdfPageResources()
	_fc._bf = width
	_fc._eda = height
	return _fc
}

// NewList creates a new list.
func (_ffgb *Creator) NewList() *List { return _gcffd(_ffgb.NewTextStyle()) }

// AddLine appends a new line to the invoice line items table.
func (_cfdf *Invoice) AddLine(values ...string) []*InvoiceCell {
	_dfeg := len(_cfdf._afeb)
	var _aeaf []*InvoiceCell
	for _ecfg, _gegg := range values {
		_cbcba := _cfdf.newCell(_gegg, _cfdf._defdd)
		if _ecfg < _dfeg {
			_cbcba.Alignment = _cfdf._afeb[_ecfg].Alignment
		}
		_aeaf = append(_aeaf, _cbcba)
	}
	_cfdf._eead = append(_cfdf._eead, _aeaf)
	return _aeaf
}

// Height returns the Block's height.
func (_aeb *Block) Height() float64 { return _aeb._eda }

// DrawWithContext draws the Block using the specified drawing context.
func (_gea *Block) DrawWithContext(d Drawable, ctx DrawContext) error {
	_dca, _, _fgd := d.GeneratePageBlocks(ctx)
	if _fgd != nil {
		return _fgd
	}
	if len(_dca) != 1 {
		return _d.New("\u0074\u006f\u006f\u0020ma\u006e\u0079\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0062\u006c\u006f\u0063k\u0073")
	}
	for _, _fgdc := range _dca {
		if _dea := _gea.mergeBlocks(_fgdc); _dea != nil {
			return _dea
		}
	}
	return nil
}
func _acdfe() *PageBreak { return &PageBreak{} }
func _gdaeb(_ffbga *templateProcessor, _eecb *templateNode) (interface{}, error) {
	return _ffbga.parseChapterHeading(_eecb)
}
func _ceeg(_eeeda *Block, _ggeb *Image, _ggcg DrawContext) (DrawContext, error) {
	_gffe := _ggcg
	_ebfe := 1
	_aceg := _fg.PdfObjectName(_eg.Sprintf("\u0049\u006d\u0067%\u0064", _ebfe))
	for _eeeda._eb.HasXObjectByName(_aceg) {
		_ebfe++
		_aceg = _fg.PdfObjectName(_eg.Sprintf("\u0049\u006d\u0067%\u0064", _ebfe))
	}
	_fagf := _eeeda._eb.SetXObjectImageByName(_aceg, _ggeb._aabc)
	if _fagf != nil {
		return _ggcg, _fagf
	}
	_efaef := 0
	_cfce := _fg.PdfObjectName(_eg.Sprintf("\u0047\u0053\u0025\u0064", _efaef))
	for _eeeda._eb.HasExtGState(_cfce) {
		_efaef++
		_cfce = _fg.PdfObjectName(_eg.Sprintf("\u0047\u0053\u0025\u0064", _efaef))
	}
	_adgc := _fg.MakeDict()
	_adgc.Set("\u0042\u004d", _fg.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
	if _ggeb._edgb < 1.0 {
		_adgc.Set("\u0043\u0041", _fg.MakeFloat(_ggeb._edgb))
		_adgc.Set("\u0063\u0061", _fg.MakeFloat(_ggeb._edgb))
	}
	_fagf = _eeeda._eb.AddExtGState(_cfce, _fg.MakeIndirectObject(_adgc))
	if _fagf != nil {
		return _ggcg, _fagf
	}
	_fgfeb := _ggeb.Width()
	_ddgda := _ggeb.Height()
	_, _dbafe := _ggeb.rotatedSize()
	_ccbe := _ggcg.X
	_cgba := _ggcg.PageHeight - _ggcg.Y - _ddgda
	if _ggeb._ccec.IsRelative() {
		_cgba -= (_dbafe - _ddgda) / 2
		switch _ggeb._cedbd {
		case HorizontalAlignmentCenter:
			_ccbe += (_ggcg.Width - _fgfeb) / 2
		case HorizontalAlignmentRight:
			_ccbe = _ggcg.PageWidth - _ggcg.Margins.Right - _ggeb._dbfe.Right - _fgfeb
		}
	}
	_ccde := _ggeb._afgea
	_egad := _eee.NewContentCreator()
	_egad.Add_gs(_cfce)
	_egad.Translate(_ccbe, _cgba)
	if _ccde != 0 {
		_egad.Translate(_fgfeb/2, _ddgda/2)
		_egad.RotateDeg(_ccde)
		_egad.Translate(-_fgfeb/2, -_ddgda/2)
	}
	_egad.Scale(_fgfeb, _ddgda).Add_Do(_aceg)
	_dagd := _egad.Operations()
	_dagd.WrapIfNeeded()
	_eeeda.addContents(_dagd)
	if _ggeb._ccec.IsRelative() {
		_ggcg.Y += _dbafe
		_ggcg.Height -= _dbafe
		return _ggcg, nil
	}
	return _gffe, nil
}

// Angle returns the block rotation angle in degrees.
func (_bc *Block) Angle() float64 { return _bc._ag }

// SetFillColor sets the fill color.
func (_fbgd *Ellipse) SetFillColor(col Color) { _fbgd._cedc = col }

// PageFinalize sets a function to be called for each page before finalization
// (i.e. the last stage of page processing before they get written out).
// The callback function allows final touch-ups for each page, and it
// provides information that might not be known at other stages of designing
// the document (e.g. the total number of pages). Unlike the header/footer
// functions, which are limited to the top/bottom margins of the page, the
// finalize function can be used draw components anywhere on the current page.
func (_dbaf *Creator) PageFinalize(pageFinalizeFunc func(_bdcc PageFinalizeFunctionArgs) error) {
	_dbaf._ebd = pageFinalizeFunc
}

// Chapter is used to arrange multiple drawables (paragraphs, images, etc) into a single section.
// The concept is the same as a book or a report chapter.
type Chapter struct {
	_aebc        int
	_bbe         string
	_dbf         *Paragraph
	_gecb        []Drawable
	_gcd         int
	_bdc         bool
	_fcde        bool
	_fabf        Positioning
	_ccdd, _adgg float64
	_ccf         Margins
	_cadbb       *Chapter
	_gdba        *TOC
	_cdfd        *_f.Outline
	_baf         *_f.OutlineItem
	_aced        uint
}

// SetBorderColor sets the border color.
func (_fegc *Polygon) SetBorderColor(color Color) { _fegc._ccbeg.BorderColor = _bcfe(color) }

const (
	DefaultHorizontalScaling = 100
)

// SetWidth sets line width.
func (_effb *Curve) SetWidth(width float64) { _effb._bbf = width }
func (_cgbc *Image) makeXObject() error {
	_ddgc := _cgbc._facba
	if _ddgc == nil {
		_ddgc = _fg.NewFlateEncoder()
	}
	_ggcae, _fgdda := _f.NewXObjectImageFromImage(_cgbc._aaca, nil, _ddgc)
	if _fgdda != nil {
		_df.Log.Error("\u0046\u0061\u0069le\u0064\u0020\u0074\u006f\u0020\u0063\u0072\u0065\u0061t\u0065 \u0078o\u0062j\u0065\u0063\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _fgdda)
		return _fgdda
	}
	_cgbc._aabc = _ggcae
	return nil
}

// NewPage adds a new Page to the Creator and sets as the active Page.
func (_dgge *Creator) NewPage() *_f.PdfPage {
	_dcfe := _dgge.newPage()
	_dgge._gadb = append(_dgge._gadb, _dcfe)
	_dgge._bfec.Page++
	return _dcfe
}
func (_cbbg *Division) ctxHeight(_fgb float64) float64 {
	_fgb -= _cbbg._fccf.Left + _cbbg._fccf.Right + _cbbg._gfeb.Left + _cbbg._gfeb.Right
	var _efdf float64
	for _, _feee := range _cbbg._aefd {
		_efdf += _dcge(_feee, _fgb)
	}
	return _efdf
}

// Width is not used. The list component is designed to fill into the available
// width depending on the context. Returns 0.
func (_fadf *List) Width() float64 { return 0 }
func (_dfee *templateProcessor) parseTableCell(_fbfcg *templateNode) (interface{}, error) {
	if _fbfcg._dfdd == nil {
		_df.Log.Error("\u0054\u0061\u0062\u006c\u0065\u0020\u0063\u0065\u006c\u006c\u0020\u0070\u0061\u0072\u0065n\u0074 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e")
		return nil, _dbea
	}
	_ecgdg, _facdb := _fbfcg._dfdd._cedba.(*Table)
	if !_facdb {
		_df.Log.Error("\u0054\u0061\u0062\u006c\u0065\u0020\u0063\u0065\u006c\u006c\u0020\u0070\u0061\u0072\u0065\u006e\u0074\u0020\u0028\u0025\u0054\u0029\u0020\u0069s\u0020\u006e\u006f\u0074\u0020a\u0020\u0074a\u0062\u006c\u0065\u002e", _fbfcg._dfdd._cedba)
		return nil, _dbea
	}
	var _cfgbd, _gbee int64
	for _, _gege := range _fbfcg._ddgde.Attr {
		_acbe := _gege.Value
		switch _adada := _gege.Name.Local; _adada {
		case "\u0063o\u006c\u0073\u0070\u0061\u006e":
			_cfgbd = _dfee.parseInt64Attr(_adada, _acbe)
		case "\u0072o\u0077\u0073\u0070\u0061\u006e":
			_gbee = _dfee.parseInt64Attr(_adada, _acbe)
		}
	}
	if _cfgbd <= 0 {
		_cfgbd = 1
	}
	if _gbee <= 0 {
		_gbee = 1
	}
	_dccaf := _ecgdg.MultiCell(int(_gbee), int(_cfgbd))
	for _, _fdbe := range _fbfcg._ddgde.Attr {
		_bddgfa := _fdbe.Value
		switch _efdae := _fdbe.Name.Local; _efdae {
		case "\u0069\u006e\u0064\u0065\u006e\u0074":
			_dccaf.SetIndent(_dfee.parseFloatAttr(_efdae, _bddgfa))
		case "\u0061\u006c\u0069g\u006e":
			_dccaf.SetHorizontalAlignment(_dfee.parseCellAlignmentAttr(_efdae, _bddgfa))
		case "\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c\u002da\u006c\u0069\u0067\u006e":
			_dccaf.SetVerticalAlignment(_dfee.parseCellVerticalAlignmentAttr(_efdae, _bddgfa))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0073\u0074\u0079\u006c\u0065":
			_dccaf.SetSideBorderStyle(CellBorderSideAll, _dfee.parseCellBorderStyleAttr(_efdae, _bddgfa))
		case "\u0062\u006fr\u0064\u0065\u0072-\u0073\u0074\u0079\u006c\u0065\u002d\u0074\u006f\u0070":
			_dccaf.SetSideBorderStyle(CellBorderSideTop, _dfee.parseCellBorderStyleAttr(_efdae, _bddgfa))
		case "\u0062\u006f\u0072\u0064er\u002d\u0073\u0074\u0079\u006c\u0065\u002d\u0062\u006f\u0074\u0074\u006f\u006d":
			_dccaf.SetSideBorderStyle(CellBorderSideBottom, _dfee.parseCellBorderStyleAttr(_efdae, _bddgfa))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0073\u0074\u0079\u006c\u0065-\u006c\u0065\u0066\u0074":
			_dccaf.SetSideBorderStyle(CellBorderSideLeft, _dfee.parseCellBorderStyleAttr(_efdae, _bddgfa))
		case "\u0062o\u0072d\u0065\u0072\u002d\u0073\u0074y\u006c\u0065-\u0072\u0069\u0067\u0068\u0074":
			_dccaf.SetSideBorderStyle(CellBorderSideRight, _dfee.parseCellBorderStyleAttr(_efdae, _bddgfa))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068":
			_dccaf.SetSideBorderWidth(CellBorderSideAll, _dfee.parseFloatAttr(_efdae, _bddgfa))
		case "\u0062\u006fr\u0064\u0065\u0072-\u0077\u0069\u0064\u0074\u0068\u002d\u0074\u006f\u0070":
			_dccaf.SetSideBorderWidth(CellBorderSideTop, _dfee.parseFloatAttr(_efdae, _bddgfa))
		case "\u0062\u006f\u0072\u0064er\u002d\u0077\u0069\u0064\u0074\u0068\u002d\u0062\u006f\u0074\u0074\u006f\u006d":
			_dccaf.SetSideBorderWidth(CellBorderSideBottom, _dfee.parseFloatAttr(_efdae, _bddgfa))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068-\u006c\u0065\u0066\u0074":
			_dccaf.SetSideBorderWidth(CellBorderSideLeft, _dfee.parseFloatAttr(_efdae, _bddgfa))
		case "\u0062o\u0072d\u0065\u0072\u002d\u0077\u0069d\u0074\u0068-\u0072\u0069\u0067\u0068\u0074":
			_dccaf.SetSideBorderWidth(CellBorderSideRight, _dfee.parseFloatAttr(_efdae, _bddgfa))
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072":
			_dccaf.SetSideBorderColor(CellBorderSideAll, _dfee.parseColorAttr(_efdae, _bddgfa))
		case "\u0062\u006fr\u0064\u0065\u0072-\u0063\u006f\u006c\u006f\u0072\u002d\u0074\u006f\u0070":
			_dccaf.SetSideBorderColor(CellBorderSideTop, _dfee.parseColorAttr(_efdae, _bddgfa))
		case "\u0062\u006f\u0072\u0064er\u002d\u0063\u006f\u006c\u006f\u0072\u002d\u0062\u006f\u0074\u0074\u006f\u006d":
			_dccaf.SetSideBorderColor(CellBorderSideBottom, _dfee.parseColorAttr(_efdae, _bddgfa))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072-\u006c\u0065\u0066\u0074":
			_dccaf.SetSideBorderColor(CellBorderSideLeft, _dfee.parseColorAttr(_efdae, _bddgfa))
		case "\u0062o\u0072d\u0065\u0072\u002d\u0063\u006fl\u006f\u0072-\u0072\u0069\u0067\u0068\u0074":
			_dccaf.SetSideBorderColor(CellBorderSideRight, _dfee.parseColorAttr(_efdae, _bddgfa))
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u006c\u0069\u006e\u0065\u002ds\u0074\u0079\u006c\u0065":
			_dccaf.SetBorderLineStyle(_dfee.parseLineStyleAttr(_efdae, _bddgfa))
		case "\u0062\u0061c\u006b\u0067\u0072o\u0075\u006e\u0064\u002d\u0063\u006f\u006c\u006f\u0072":
			_dccaf.SetBackgroundColor(_dfee.parseColorAttr(_efdae, _bddgfa))
		case "\u0063o\u006c\u0073\u0070\u0061\u006e", "\u0072o\u0077\u0073\u0070\u0061\u006e":
		default:
			_df.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0074\u0061\u0062\u006c\u0065\u0020\u0063\u0065\u006c\u006c\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e", _efdae)
		}
	}
	return _dccaf, nil
}

// IsAbsolute checks if the positioning is absolute.
func (_egaa Positioning) IsAbsolute() bool { return _egaa == PositionAbsolute }

// EnableWordWrap sets the paragraph word wrap flag.
func (_debgf *StyledParagraph) EnableWordWrap(val bool) { _debgf._bggd = val }
func (_eagag *templateProcessor) parseTable(_agceb *templateNode) (interface{}, error) {
	var _caacc int64
	for _, _eaaac := range _agceb._ddgde.Attr {
		_afabg := _eaaac.Value
		switch _ggabg := _eaaac.Name.Local; _ggabg {
		case "\u0063o\u006c\u0075\u006d\u006e\u0073":
			_caacc = _eagag.parseInt64Attr(_ggabg, _afabg)
		}
	}
	if _caacc <= 0 {
		_df.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006eu\u006d\u0062e\u0072\u0020\u006f\u0066\u0020\u0074\u0061\u0062\u006ce\u0020\u0063\u006f\u006cu\u006d\u006e\u0073\u003a\u0020\u0025\u0064\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0031\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020m\u0061\u0079\u0020b\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e", _caacc)
		_caacc = 1
	}
	_eafd := _eagag.creator.NewTable(int(_caacc))
	for _, _afdbd := range _agceb._ddgde.Attr {
		_fgfa := _afdbd.Value
		switch _edeg := _afdbd.Name.Local; _edeg {
		case "\u0063\u006f\u006c\u0075\u006d\u006e\u002d\u0077\u0069\u0064\u0074\u0068\u0073":
			_eafd.SetColumnWidths(_eagag.parseFloatArray(_edeg, _fgfa)...)
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_dacc := _eagag.parseMarginAttr(_edeg, _fgfa)
			_eafd.SetMargins(_dacc.Left, _dacc.Right, _dacc.Top, _dacc.Bottom)
		case "\u0078":
			_eafd.SetPos(_eagag.parseFloatAttr(_edeg, _fgfa), _eafd._eebda)
		case "\u0079":
			_eafd.SetPos(_eafd._fdcbfa, _eagag.parseFloatAttr(_edeg, _fgfa))
		case "\u0068\u0065a\u0064\u0065\u0072-\u0073\u0074\u0061\u0072\u0074\u002d\u0072\u006f\u0077":
			_eafd._fdfff = int(_eagag.parseInt64Attr(_edeg, _fgfa))
		case "\u0068\u0065\u0061\u0064\u0065\u0072\u002d\u0065\u006ed\u002d\u0072\u006f\u0077":
			_eafd._bdffd = int(_eagag.parseInt64Attr(_edeg, _fgfa))
		case "\u0065n\u0061b\u006c\u0065\u002d\u0072\u006f\u0077\u002d\u0077\u0072\u0061\u0070":
			_eafd.EnableRowWrap(_eagag.parseBoolAttr(_edeg, _fgfa))
		case "\u0065\u006ea\u0062\u006c\u0065-\u0070\u0061\u0067\u0065\u002d\u0077\u0072\u0061\u0070":
			_eafd.EnablePageWrap(_eagag.parseBoolAttr(_edeg, _fgfa))
		case "\u0063o\u006c\u0075\u006d\u006e\u0073":
		default:
			_df.Log.Debug("\u0055n\u0073\u0075p\u0070\u006f\u0072\u0074e\u0064\u0020\u0074a\u0062\u006c\u0065\u0020\u0061\u0074\u0074\u0072\u0069bu\u0074\u0065\u003a \u0060\u0025s\u0060\u002e\u0020\u0053\u006b\u0069p\u0070\u0069n\u0067\u002e", _edeg)
		}
	}
	if _eafd._fdfff != 0 && _eafd._bdffd != 0 {
		_cfdb := _eafd.SetHeaderRows(_eafd._fdfff, _eafd._bdffd)
		if _cfdb != nil {
			_df.Log.Debug("\u0043\u006ful\u0064\u0020\u006eo\u0074\u0020\u0073\u0065t t\u0061bl\u0065\u0020\u0068\u0065\u0061\u0064\u0065r \u0072\u006f\u0077\u0073\u003a\u0020\u0025v\u002e", _cfdb)
		}
	} else {
		_eafd._fdfff = 0
		_eafd._bdffd = 0
	}
	return _eafd, nil
}
func _fcfea(_cfba *_f.PdfAnnotationLink) *_f.PdfAnnotationLink {
	if _cfba == nil {
		return nil
	}
	_gacab := _f.NewPdfAnnotationLink()
	_gacab.BS = _cfba.BS
	_gacab.A = _cfba.A
	if _gfacc, _degf := _cfba.GetAction(); _degf == nil && _gfacc != nil {
		_gacab.SetAction(_gfacc)
	}
	if _fdga, _bgdda := _cfba.Dest.(*_fg.PdfObjectArray); _bgdda {
		_gacab.Dest = _fg.MakeArray(_fdga.Elements()...)
	}
	return _gacab
}

const (
	CellHorizontalAlignmentLeft CellHorizontalAlignment = iota
	CellHorizontalAlignmentCenter
	CellHorizontalAlignmentRight
)

func (_gdeag *Table) clone() *Table {
	_aaae := *_gdeag
	_aaae._ceeb = make([]float64, len(_gdeag._ceeb))
	copy(_aaae._ceeb, _gdeag._ceeb)
	_aaae._agdgd = make([]float64, len(_gdeag._agdgd))
	copy(_aaae._agdgd, _gdeag._agdgd)
	_aaae._ebfb = make([]*TableCell, 0, len(_gdeag._ebfb))
	for _, _fcecd := range _gdeag._ebfb {
		_ffeg := *_fcecd
		_ffeg._fbdc = &_aaae
		_aaae._ebfb = append(_aaae._ebfb, &_ffeg)
	}
	return &_aaae
}

// SetMaxLines sets the maximum number of lines before the paragraph
// text is truncated.
func (_egac *Paragraph) SetMaxLines(maxLines int) { _egac._eedca = maxLines; _egac.wrapText() }

// EnableRowWrap controls whether rows are wrapped across pages.
// NOTE: Currently, row wrapping is supported for rows using StyledParagraphs.
func (_fbeg *Table) EnableRowWrap(enable bool) { _fbeg._ggbfc = enable }
func (_acbca *templateProcessor) parseLineStyleAttr(_ddcg, _ecaf string) _cec.LineStyle {
	_df.Log.Debug("\u0050\u0061\u0072\u0073\u0069n\u0067\u0020\u006c\u0069\u006e\u0065\u0020\u0073\u0074\u0079\u006c\u0065\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _ddcg, _ecaf)
	_bfcca := map[string]_cec.LineStyle{"\u0073\u006f\u006ci\u0064": _cec.LineStyleSolid, "\u0064\u0061\u0073\u0068\u0065\u0064": _cec.LineStyleDashed}[_ecaf]
	return _bfcca
}

type pageTransformations struct {
	_daa  *_edf.Matrix
	_gdd  bool
	_cbcb bool
}

// Width returns the current page width.
func (_gcfb *Creator) Width() float64 { return _gcfb._gee }
func (_cbda *templateProcessor) parseColorAttr(_gfedd, _adbe string) Color {
	_df.Log.Debug("\u0050\u0061rs\u0069\u006e\u0067 \u0063\u006f\u006c\u006fr a\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020(`\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _gfedd, _adbe)
	_bcdfb := ColorBlack
	if _adbe == "" {
		return _bcdfb
	}
	_ebgfa, _dgcde := _cbda._daae.ColorMap[_adbe]
	if _dgcde {
		return _ebgfa
	}
	if _adbe[0] == '#' {
		return ColorRGBFromHex(_adbe)
	}
	return _bcdfb
}

// NoteHeadingStyle returns the style properties used to render the heading of
// the invoice note sections.
func (_fbda *Invoice) NoteHeadingStyle() TextStyle { return _fbda._cacg }

// TextRenderingMode determines whether showing text shall cause glyph
// outlines to be stroked, filled, used as a clipping boundary, or some
// combination of the three.
// See section 9.3 "Text State Parameters and Operators" and
// Table 106 (pp. 254-255 PDF32000_2008).
type TextRenderingMode int

// Paragraph represents text drawn with a specified font and can wrap across lines and pages.
// By default it occupies the available width in the drawing context.
type Paragraph struct {
	_ebea        string
	_ceda        *_f.PdfFont
	_gced        float64
	_abbg        float64
	_bfbc        Color
	_badg        TextAlignment
	_bbfc        bool
	_bdga        float64
	_eedca       int
	_bbbeg       bool
	_fccd        float64
	_eacf        Margins
	_beeba       Positioning
	_fbbde       float64
	_agdg        float64
	_accg, _agee float64
	_cead        []string
}

func (_fecd *templateProcessor) parseCellVerticalAlignmentAttr(_abgdd, _gddcfg string) CellVerticalAlignment {
	_df.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u0065\u006c\u006c\u0020\u0076\u0065r\u0074\u0069\u0063\u0061\u006c\u0020\u0061\u006c\u0069\u0067\u006e\u006d\u0065n\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a (\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _abgdd, _gddcfg)
	_aeeff := map[string]CellVerticalAlignment{"\u0074\u006f\u0070": CellVerticalAlignmentTop, "\u006d\u0069\u0064\u0064\u006c\u0065": CellVerticalAlignmentMiddle, "\u0062\u006f\u0074\u0074\u006f\u006d": CellVerticalAlignmentBottom}[_gddcfg]
	return _aeeff
}

// InvoiceCell represents any cell belonging to a table from the invoice
// template. The main tables are the invoice information table, the line
// items table and totals table. Contains the text value of the cell and
// the style properties of the cell.
type InvoiceCell struct {
	InvoiceCellProps
	Value string
}

// TextChunk represents a chunk of text along with a particular style.
type TextChunk struct {

	// The text that is being rendered in the PDF.
	Text string

	// The style of the text being rendered.
	Style  TextStyle
	_ceddd *_f.PdfAnnotation
	_aaagg bool
}

// SetPageMargins sets the page margins: left, right, top, bottom.
// The default page margins are 10% of document width.
func (_eace *Creator) SetPageMargins(left, right, top, bottom float64) {
	_eace._fade.Left = left
	_eace._fade.Right = right
	_eace._fade.Top = top
	_eace._fade.Bottom = bottom
}

// Scale scales Image by a constant factor, both width and height.
func (_agfc *Image) Scale(xFactor, yFactor float64) {
	_agfc._bcfd = xFactor * _agfc._bcfd
	_agfc._gddc = yFactor * _agfc._gddc
}
func (_faed *templateProcessor) parseInt64Array(_gdbb, _eeag string) []int64 {
	_df.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020\u0069\u006e\u0074\u0036\u0034\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060%\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _gdbb, _eeag)
	_dcafe := _ed.Fields(_eeag)
	_fceba := make([]int64, 0, len(_dcafe))
	for _, _ebbe := range _dcafe {
		_acac, _ := _ae.ParseInt(_ebbe, 10, 64)
		_fceba = append(_fceba, _acac)
	}
	return _fceba
}

// SetRowHeight sets the height for a specified row.
func (_gdae *Table) SetRowHeight(row int, h float64) error {
	if row < 1 || row > len(_gdae._ceeb) {
		return _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_gdae._ceeb[row-1] = h
	return nil
}

// SetBorderColor sets the border color.
func (_bdgg *PolyBezierCurve) SetBorderColor(color Color) { _bdgg._cegf.BorderColor = _bcfe(color) }
func (_gafa *StyledParagraph) getTextLineWidth(_bgce []*TextChunk) float64 {
	var _aaec float64
	_cecd := len(_bgce)
	for _fgfd, _bcdf := range _bgce {
		_eggd := &_bcdf.Style
		_gfedc := len(_bcdf.Text)
		for _efca, _dfaf := range _bcdf.Text {
			if _dfaf == '\u000A' {
				continue
			}
			_dfdc, _fgeb := _eggd.Font.GetRuneMetrics(_dfaf)
			if !_fgeb {
				_df.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _dfaf)
				return -1
			}
			_aaec += _eggd.FontSize * _dfdc.Wx * _eggd.horizontalScale()
			if _dfaf != ' ' && (_fgfd != _cecd-1 || _efca != _gfedc-1) {
				_aaec += _eggd.CharSpacing * 1000.0
			}
		}
	}
	return _aaec
}
func (_edaec *TextStyle) horizontalScale() float64 { return _edaec.HorizontalScaling / 100 }

// SetIndent sets the left offset of the list when nested into another list.
func (_agce *List) SetIndent(indent float64) { _agce._dgab = indent; _agce._ggbf = false }

// SellerAddress returns the seller address used in the invoice template.
func (_dbg *Invoice) SellerAddress() *InvoiceAddress { return _dbg._bgfe }

// HeaderFunctionArgs holds the input arguments to a header drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type HeaderFunctionArgs struct {
	PageNum    int
	TotalPages int
}

// NewImageFromFile creates an Image from a file.
func (_daag *Creator) NewImageFromFile(path string) (*Image, error) { return _fffeg(path) }
func (_efbc *pageTransformations) transformBlock(_bfcb *Block) {
	if _efbc._daa != nil {
		_bfcb.transform(*_efbc._daa)
	}
}
func _ebcb(_dbfg string, _fbgb TextStyle) *Paragraph {
	_dgaba := &Paragraph{_ebea: _dbfg, _ceda: _fbgb.Font, _gced: _fbgb.FontSize, _abbg: 1.0, _bbfc: true, _bbbeg: true, _badg: TextAlignmentLeft, _fccd: 0, _accg: 1, _agee: 1, _beeba: PositionRelative}
	_dgaba.SetColor(_fbgb.Color)
	return _dgaba
}
func (_fdeg *TOCLine) prepareParagraph(_ececb *StyledParagraph, _bbgff DrawContext) {
	_eacc := _fdeg.Title.Text
	if _fdeg.Number.Text != "" {
		_eacc = "\u0020" + _eacc
	}
	_eacc += "\u0020"
	_agedg := _fdeg.Page.Text
	if _agedg != "" {
		_agedg = "\u0020" + _agedg
	}
	_ececb._eedad = []*TextChunk{{Text: _fdeg.Number.Text, Style: _fdeg.Number.Style, _ceddd: _fdeg.getLineLink()}, {Text: _eacc, Style: _fdeg.Title.Style, _ceddd: _fdeg.getLineLink()}, {Text: _agedg, Style: _fdeg.Page.Style, _ceddd: _fdeg.getLineLink()}}
	_ececb.wrapText()
	_gdac := len(_ececb._bfaa)
	if _gdac == 0 {
		return
	}
	_decg := _bbgff.Width*1000 - _ececb.getTextLineWidth(_ececb._bfaa[_gdac-1])
	_edfbfg := _ececb.getTextLineWidth([]*TextChunk{&_fdeg.Separator})
	_cfff := int(_decg / _edfbfg)
	_abdfe := _ed.Repeat(_fdeg.Separator.Text, _cfff)
	_acab := _fdeg.Separator.Style
	_gdbfg := _ececb.Insert(2, _abdfe)
	_gdbfg.Style = _acab
	_gdbfg._ceddd = _fdeg.getLineLink()
	_decg = _decg - float64(_cfff)*_edfbfg
	if _decg > 500 {
		_gbfcf, _gabeb := _acab.Font.GetRuneMetrics(' ')
		if _gabeb && _decg > _gbfcf.Wx {
			_befdb := int(_decg / _gbfcf.Wx)
			if _befdb > 0 {
				_efgfb := _acab
				_efgfb.FontSize = 1
				_gdbfg = _ececb.Insert(2, _ed.Repeat("\u0020", _befdb))
				_gdbfg.Style = _efgfb
				_gdbfg._ceddd = _fdeg.getLineLink()
			}
		}
	}
}

// GetMargins returns the Paragraph's margins: left, right, top, bottom.
func (_gfafd *Paragraph) GetMargins() (float64, float64, float64, float64) {
	return _gfafd._eacf.Left, _gfafd._eacf.Right, _gfafd._eacf.Top, _gfafd._eacf.Bottom
}

// SetStyleTop sets border style for top side.
func (_fcaf *border) SetStyleTop(style CellBorderStyle) { _fcaf._faaa = style }

// ColorRGBFrom8bit creates a Color from 8-bit (0-255) r,g,b values.
// Example:
//   red := ColorRGBFrom8Bit(255, 0, 0)
func ColorRGBFrom8bit(r, g, b byte) Color {
	return rgbColor{_fabb: float64(r) / 255.0, _abab: float64(g) / 255.0, _gggf: float64(b) / 255.0}
}

// SkipOver skips over a specified number of rows and cols.
func (_cafb *Table) SkipOver(rows, cols int) {
	_faddc := rows*_cafb._cfee + cols - 1
	if _faddc < 0 {
		_df.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	_cafb._edag += _faddc
}

// NewStyledParagraph creates a new styled paragraph.
// Default attributes:
// Font: Helvetica,
// Font size: 10
// Encoding: WinAnsiEncoding
// Wrap: enabled
// Text color: black
func (_gaga *Creator) NewStyledParagraph() *StyledParagraph { return _ggaa(_gaga.NewTextStyle()) }

// SetLevel sets the indentation level of the TOC line.
func (_gccd *TOCLine) SetLevel(level uint) {
	_gccd._dgcda = level
	_gccd._face._afdba.Left = _gccd._ecdbc + float64(_gccd._dgcda-1)*_gccd._aaaa
}

type containerDrawable interface {
	Drawable

	// ContainerComponent checks if the component is allowed to be added into provided 'container' and returns
	// preprocessed copy of itself. If the component is not changed it is allowed to return itself in a callback way.
	// If the component is not compatible with provided container this method should return an error.
	ContainerComponent(_ceef Drawable) (Drawable, error)
}

func (_egca *Chapter) headingText() string {
	_cbdc := _egca._bbe
	if _dgdc := _egca.headingNumber(); _dgdc != "" {
		_cbdc = _eg.Sprintf("\u0025\u0073\u0020%\u0073", _dgdc, _cbdc)
	}
	return _cbdc
}

// String implements error interface.
func (_fdgd UnsupportedRuneError) Error() string { return _fdgd.Message }

// GetRowHeight returns the height of the specified row.
func (_agfd *Table) GetRowHeight(row int) (float64, error) {
	if row < 1 || row > len(_agfd._ceeb) {
		return 0, _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	return _agfd._ceeb[row-1], nil
}

// NewStyledTOCLine creates a new table of contents line with the provided style.
func (_becg *Creator) NewStyledTOCLine(number, title, page TextChunk, level uint, style TextStyle) *TOCLine {
	return _ffggc(number, title, page, level, style)
}

// SetAngle sets the rotation angle of the text.
func (_ddag *Paragraph) SetAngle(angle float64) { _ddag._fccd = angle }

// Width returns the width of the Paragraph.
func (_gfef *StyledParagraph) Width() float64 {
	if _gfef._ecbf && int(_gfef._bbef) > 0 {
		return _gfef._bbef
	}
	return _gfef.getTextWidth() / 1000.0
}

// Lines returns all the rows of the invoice line items table.
func (_gebe *Invoice) Lines() [][]*InvoiceCell { return _gebe._eead }

// Width returns the width of the chart. In relative positioning mode,
// all the available context width is used at render time.
func (_cffg *Chart) Width() float64 { return float64(_cffg._fbaa.Width()) }

// SetNumber sets the number of the invoice.
func (_acb *Invoice) SetNumber(number string) (*InvoiceCell, *InvoiceCell) {
	_acb._edeb[1].Value = number
	return _acb._edeb[0], _acb._edeb[1]
}
func (_fcba *templateProcessor) parseParagraph(_aaage *templateNode, _efbg *Paragraph) (interface{}, error) {
	if _efbg == nil {
		_efbg = _fcba.creator.NewParagraph("")
	}
	for _, _agbe := range _aaage._ddgde.Attr {
		_afadb := _agbe.Value
		switch _egab := _agbe.Name.Local; _egab {
		case "\u0066\u006f\u006e\u0074":
			_efbg.SetFont(_fcba.parseFontAttr(_egab, _afadb))
		case "\u0066o\u006e\u0074\u002d\u0073\u0069\u007ae":
			_efbg.SetFontSize(_fcba.parseFloatAttr(_egab, _afadb))
		case "\u0074\u0065\u0078\u0074\u002d\u0061\u006c\u0069\u0067\u006e":
			_efbg.SetTextAlignment(_fcba.parseTextAlignmentAttr(_egab, _afadb))
		case "l\u0069\u006e\u0065\u002d\u0068\u0065\u0069\u0067\u0068\u0074":
			_efbg.SetLineHeight(_fcba.parseFloatAttr(_egab, _afadb))
		case "e\u006e\u0061\u0062\u006c\u0065\u002d\u0077\u0072\u0061\u0070":
			_efbg.SetEnableWrap(_fcba.parseBoolAttr(_egab, _afadb))
		case "\u0063\u006f\u006co\u0072":
			_efbg.SetColor(_fcba.parseColorAttr(_egab, _afadb))
		case "\u0078":
			_efbg.SetPos(_fcba.parseFloatAttr(_egab, _afadb), _efbg._agdg)
		case "\u0079":
			_efbg.SetPos(_efbg._fbbde, _fcba.parseFloatAttr(_egab, _afadb))
		case "\u0061\u006e\u0067l\u0065":
			_efbg.SetAngle(_fcba.parseFloatAttr(_egab, _afadb))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_fdcef := _fcba.parseMarginAttr(_egab, _afadb)
			_efbg.SetMargins(_fdcef.Left, _fdcef.Right, _fdcef.Top, _fdcef.Bottom)
		case "\u006da\u0078\u002d\u006c\u0069\u006e\u0065s":
			_efbg.SetMaxLines(int(_fcba.parseInt64Attr(_egab, _afadb)))
		default:
			_df.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065\u0064\u0020pa\u0072\u0061\u0067\u0072\u0061\u0070h\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073`\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069n\u0067\u002e", _egab)
		}
	}
	return _efbg, nil
}

// SetBorderColor sets the border color.
func (_ffd *Ellipse) SetBorderColor(col Color) { _ffd._agcc = col }

// MoveY moves the drawing context to absolute position y.
func (_aaba *Creator) MoveY(y float64) { _aaba._bfec.Y = y }

// SetAngle sets the rotation angle of the text.
func (_gaee *StyledParagraph) SetAngle(angle float64) { _gaee._daab = angle }
func (_add *Invoice) generateLineBlocks(_eaed DrawContext) ([]*Block, DrawContext, error) {
	_eadd := _ecef(len(_add._afeb))
	_eadd.SetMargins(0, 0, 25, 0)
	for _, _febff := range _add._afeb {
		_agcfe := _ggaa(_febff.TextStyle)
		_agcfe.SetMargins(0, 0, 1, 0)
		_agcfe.Append(_febff.Value)
		_dcec := _eadd.NewCell()
		_dcec.SetHorizontalAlignment(_febff.Alignment)
		_dcec.SetBackgroundColor(_febff.BackgroundColor)
		_add.setCellBorder(_dcec, _febff)
		_dcec.SetContent(_agcfe)
	}
	for _, _aeafb := range _add._eead {
		for _, _ffcg := range _aeafb {
			_bbab := _ggaa(_ffcg.TextStyle)
			_bbab.SetMargins(0, 0, 3, 2)
			_bbab.Append(_ffcg.Value)
			_bddgf := _eadd.NewCell()
			_bddgf.SetHorizontalAlignment(_ffcg.Alignment)
			_bddgf.SetBackgroundColor(_ffcg.BackgroundColor)
			_add.setCellBorder(_bddgf, _ffcg)
			_bddgf.SetContent(_bbab)
		}
	}
	return _eadd.GeneratePageBlocks(_eaed)
}

var (
	ColorBlack  = ColorRGBFromArithmetic(0, 0, 0)
	ColorWhite  = ColorRGBFromArithmetic(1, 1, 1)
	ColorRed    = ColorRGBFromArithmetic(1, 0, 0)
	ColorGreen  = ColorRGBFromArithmetic(0, 1, 0)
	ColorBlue   = ColorRGBFromArithmetic(0, 0, 1)
	ColorYellow = ColorRGBFromArithmetic(1, 1, 0)
)

func (_gge *Block) mergeBlocks(_cdf *Block) error {
	_fga := _ggg(_gge._ga, _gge._eb, _cdf._ga, _cdf._eb)
	if _fga != nil {
		return _fga
	}
	for _, _aba := range _cdf._cad {
		_gge.AddAnnotation(_aba)
	}
	return nil
}

// SetStyle sets the style for all the line components: number, title,
// separator, page.
func (_bdgbg *TOCLine) SetStyle(style TextStyle) {
	_bdgbg.Number.Style = style
	_bdgbg.Title.Style = style
	_bdgbg.Separator.Style = style
	_bdgbg.Page.Style = style
}

// SetPositioning sets Rectangle's position attribute.
func (_gefff *Rectangle) SetPositioning(position Positioning) { _gefff._fcbgb = position }

// GeneratePageBlocks draws the rectangle on a new block representing the page. Implements the Drawable interface.
func (_baggc *Rectangle) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_bbgb := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_afabc := _cec.Rectangle{Opacity: 1.0, X: _baggc._bagb, Y: ctx.PageHeight - _baggc._cbe - _baggc._ecca, Height: _baggc._ecca, Width: _baggc._ecfc, BorderRadiusTopLeft: _baggc._ecdc, BorderRadiusTopRight: _baggc._aaegc, BorderRadiusBottomLeft: _baggc._bbde, BorderRadiusBottomRight: _baggc._dada}
	if _baggc._fcbgb == PositionRelative {
		_afabc.X = ctx.X
		_afabc.Y = ctx.PageHeight - ctx.Y - _baggc._ecca
	}
	if _baggc._aebg != nil {
		_afabc.FillEnabled = true
		_afabc.FillColor = _bcfe(_baggc._aebg)
	}
	if _baggc._afcb != nil && _baggc._bffaf > 0 {
		_afabc.BorderEnabled = true
		_afabc.BorderColor = _bcfe(_baggc._afcb)
		_afabc.BorderWidth = _baggc._bffaf
	}
	_bbdb, _accgf := _bbgb.setOpacity(_baggc._ffcbg, _baggc._beeac)
	if _accgf != nil {
		return nil, ctx, _accgf
	}
	_bgbf, _, _accgf := _afabc.Draw(_bbdb)
	if _accgf != nil {
		return nil, ctx, _accgf
	}
	if _accgf = _bbgb.addContentsByString(string(_bgbf)); _accgf != nil {
		return nil, ctx, _accgf
	}
	return []*Block{_bbgb}, ctx, nil
}

// SetStyleBottom sets border style for bottom side.
func (_cac *border) SetStyleBottom(style CellBorderStyle) { _cac._aeba = style }
func (_bbag *StyledParagraph) split(_acda DrawContext) (_eeedag, _cefc *StyledParagraph, _aaafdf error) {
	if _aaafdf = _bbag.wrapChunks(false); _aaafdf != nil {
		return nil, nil, _aaafdf
	}
	_egdbc := func(_dddb []*TextChunk, _fccdb []*TextChunk) []*TextChunk {
		if len(_fccdb) == 0 {
			return _dddb
		}
		_bdbc := len(_dddb)
		if _bdbc == 0 {
			return append(_dddb, _fccdb...)
		}
		if _dddb[_bdbc-1].Style == _fccdb[0].Style {
			_dddb[_bdbc-1].Text += _fccdb[0].Text
		} else {
			_dddb = append(_dddb, _fccdb[0])
		}
		return append(_dddb, _fccdb[1:]...)
	}
	_gbdf := func(_ccgc *StyledParagraph, _efeb []*TextChunk) *StyledParagraph {
		if len(_efeb) == 0 {
			return nil
		}
		_cgcg := *_ccgc
		_cgcg._eedad = _efeb
		return &_cgcg
	}
	var (
		_faeab float64
		_acae  []*TextChunk
		_degd  []*TextChunk
	)
	for _, _adgfc := range _bbag._bfaa {
		var _fggad float64
		_decf := make([]*TextChunk, 0, len(_adgfc))
		for _, _ecccb := range _adgfc {
			if _agdef := _ecccb.Style.FontSize; _agdef > _fggad {
				_fggad = _agdef
			}
			_decf = append(_decf, _ecccb.clone())
		}
		_fggad *= _bbag._bged
		if _bbag._geagd.IsRelative() {
			if _faeab+_fggad > _acda.Height {
				_degd = _egdbc(_degd, _decf)
			} else {
				_acae = _egdbc(_acae, _decf)
			}
		}
		_faeab += _fggad
	}
	_bbag._bfaa = nil
	if len(_degd) == 0 {
		return _bbag, nil, nil
	}
	return _gbdf(_bbag, _acae), _gbdf(_bbag, _degd), nil
}

// Inline returns whether the inline mode of the division is active.
func (_bffa *Division) Inline() bool { return _bffa._ddda }

// FooterFunctionArgs holds the input arguments to a footer drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type FooterFunctionArgs struct {
	PageNum    int
	TotalPages int
}

// GetCoords returns the coordinates of the Ellipse's center (xc,yc).
func (_gace *Ellipse) GetCoords() (float64, float64) { return _gace._agdd, _gace._ddgg }

// ColorRGBFromArithmetic creates a Color from arithmetic color values (0-1).
// Example:
//   green := ColorRGBFromArithmetic(0.0, 1.0, 0.0)
func ColorRGBFromArithmetic(r, g, b float64) Color {
	return rgbColor{_fabb: _a.Max(_a.Min(r, 1.0), 0.0), _abab: _a.Max(_a.Min(g, 1.0), 0.0), _gggf: _a.Max(_a.Min(b, 1.0), 0.0)}
}

// SetFont sets the Paragraph's font.
func (_bfag *Paragraph) SetFont(font *_f.PdfFont) { _bfag._ceda = font }
func (_aaag *Invoice) newCell(_gdbf string, _gceg InvoiceCellProps) *InvoiceCell {
	return &InvoiceCell{_gceg, _gdbf}
}

// Add appends a new item to the list.
// The supported components are: *Paragraph, *StyledParagraph and *List.
// Returns the marker used for the newly added item. The returned marker
// object can be used to change the text and style of the marker for the
// current item.
func (_bbge *List) Add(item VectorDrawable) (*TextChunk, error) {
	_gbbc := &listItem{_cdef: item, _gade: _bbge._fbfe}
	switch _gafbcg := item.(type) {
	case *Paragraph:
	case *StyledParagraph:
	case *List:
		if _gafbcg._ggbf {
			_gafbcg._dgab = 15
		}
	default:
		return nil, _d.New("\u0074\u0068i\u0073\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066\u0020\u0064\u0072\u0061\u0077\u0061\u0062\u006c\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u0020\u006c\u0069\u0073\u0074")
	}
	_bbge._eddg = append(_bbge._eddg, _gbbc)
	return &_gbbc._gade, nil
}

// AppendCurve appends a Bezier curve to the filled curve.
func (_geaeb *FilledCurve) AppendCurve(curve _cec.CubicBezierCurve) *FilledCurve {
	_geaeb._bgf = append(_geaeb._bgf, curve)
	return _geaeb
}

// SetMargins sets the margins of the line.
// NOTE: line margins are only applied if relative positioning is used.
func (_acfc *Line) SetMargins(left, right, top, bottom float64) {
	_acfc._dacfg.Left = left
	_acfc._dacfg.Right = right
	_acfc._dacfg.Top = top
	_acfc._dacfg.Bottom = bottom
}

// SetHeading sets the text and the style of the heading of the TOC component.
func (_aaebc *TOC) SetHeading(text string, style TextStyle) {
	_aecfe := _aaebc.Heading()
	_aecfe.Reset()
	_ceebf := _aecfe.Append(text)
	_ceebf.Style = style
}

// Sections returns the custom content sections of the invoice as
// title-content pairs.
func (_fggg *Invoice) Sections() [][2]string { return _fggg._bgfee }

// Line defines a line between point 1 (X1, Y1) and point 2 (X2, Y2).
// The line width, color, style (solid or dashed) and opacity can be
// configured. Implements the Drawable interface.
type Line struct {
	_ggfd  float64
	_fadc  float64
	_dccb  float64
	_gdec  float64
	_ggdcb Color
	_bfcbd _cec.LineStyle
	_cgcbf float64
	_gbde  []int64
	_afbf  int64
	_ffgbg float64
	_edcee Positioning
	_gggcb FitMode
	_dacfg Margins
}

func (_gdce *Invoice) drawSection(_aead, _abgb string) []*StyledParagraph {
	var _eage []*StyledParagraph
	if _aead != "" {
		_defb := _ggaa(_gdce._cacg)
		_defb.SetMargins(0, 0, 0, 5)
		_defb.Append(_aead)
		_eage = append(_eage, _defb)
	}
	if _abgb != "" {
		_cabf := _ggaa(_gdce._egf)
		_cabf.Append(_abgb)
		_eage = append(_eage, _cabf)
	}
	return _eage
}
func _ggg(_cce *_eee.ContentStreamOperations, _cece *_f.PdfPageResources, _aaa *_eee.ContentStreamOperations, _fbe *_f.PdfPageResources) error {
	_bae := map[_fg.PdfObjectName]_fg.PdfObjectName{}
	_daf := map[_fg.PdfObjectName]_fg.PdfObjectName{}
	_cbb := map[_fg.PdfObjectName]_fg.PdfObjectName{}
	_fcf := map[_fg.PdfObjectName]_fg.PdfObjectName{}
	_egb := map[_fg.PdfObjectName]_fg.PdfObjectName{}
	_ccd := map[_fg.PdfObjectName]_fg.PdfObjectName{}
	for _, _fad := range *_aaa {
		switch _fad.Operand {
		case "\u0044\u006f":
			if len(_fad.Params) == 1 {
				if _ebga, _gac := _fad.Params[0].(*_fg.PdfObjectName); _gac {
					if _, _gaca := _bae[*_ebga]; !_gaca {
						var _bcee _fg.PdfObjectName
						_gdb, _ := _fbe.GetXObjectByName(*_ebga)
						if _gdb != nil {
							_bcee = *_ebga
							for {
								_ebbf, _ := _cece.GetXObjectByName(_bcee)
								if _ebbf == nil || _ebbf == _gdb {
									break
								}
								_bcee = _bcee + "\u0030"
							}
						}
						_cece.SetXObjectByName(_bcee, _gdb)
						_bae[*_ebga] = _bcee
					}
					_efd := _bae[*_ebga]
					_fad.Params[0] = &_efd
				}
			}
		case "\u0054\u0066":
			if len(_fad.Params) == 2 {
				if _ffg, _efdd := _fad.Params[0].(*_fg.PdfObjectName); _efdd {
					if _, _cae := _daf[*_ffg]; !_cae {
						_dgc, _fcae := _fbe.GetFontByName(*_ffg)
						_gbg := *_ffg
						if _fcae && _dgc != nil {
							_gbg = _agf(_ffg.String(), _dgc, _cece)
						}
						_cece.SetFontByName(_gbg, _dgc)
						_daf[*_ffg] = _gbg
					}
					_afe := _daf[*_ffg]
					_fad.Params[0] = &_afe
				}
			}
		case "\u0043\u0053", "\u0063\u0073":
			if len(_fad.Params) == 1 {
				if _cge, _cedb := _fad.Params[0].(*_fg.PdfObjectName); _cedb {
					if _, _baa := _cbb[*_cge]; !_baa {
						var _geg _fg.PdfObjectName
						_bdb, _gc := _fbe.GetColorspaceByName(*_cge)
						if _gc {
							_geg = *_cge
							for {
								_gfd, _cbfd := _cece.GetColorspaceByName(_geg)
								if !_cbfd || _bdb == _gfd {
									break
								}
								_geg = _geg + "\u0030"
							}
							_cece.SetColorspaceByName(_geg, _bdb)
							_cbb[*_cge] = _geg
						} else {
							_df.Log.Debug("C\u006fl\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064")
						}
					}
					if _dbb, _cbfc := _cbb[*_cge]; _cbfc {
						_fad.Params[0] = &_dbb
					} else {
						_df.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0043\u006f\u006co\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", *_cge)
					}
				}
			}
		case "\u0053\u0043\u004e", "\u0073\u0063\u006e":
			if len(_fad.Params) == 1 {
				if _cef, _cadb := _fad.Params[0].(*_fg.PdfObjectName); _cadb {
					if _, _efa := _fcf[*_cef]; !_efa {
						var _cfg _fg.PdfObjectName
						_dfgc, _faa := _fbe.GetPatternByName(*_cef)
						if _faa {
							_cfg = *_cef
							for {
								_caef, _ffb := _cece.GetPatternByName(_cfg)
								if !_ffb || _caef == _dfgc {
									break
								}
								_cfg = _cfg + "\u0030"
							}
							_cfe := _cece.SetPatternByName(_cfg, _dfgc.ToPdfObject())
							if _cfe != nil {
								return _cfe
							}
							_fcf[*_cef] = _cfg
						}
					}
					if _gde, _gggce := _fcf[*_cef]; _gggce {
						_fad.Params[0] = &_gde
					}
				}
			}
		case "\u0073\u0068":
			if len(_fad.Params) == 1 {
				if _dce, _dec := _fad.Params[0].(*_fg.PdfObjectName); _dec {
					if _, _acf := _egb[*_dce]; !_acf {
						var _bdg _fg.PdfObjectName
						_abf, _fab := _fbe.GetShadingByName(*_dce)
						if _fab {
							_bdg = *_dce
							for {
								_cab, _bgg := _cece.GetShadingByName(_bdg)
								if !_bgg || _abf == _cab {
									break
								}
								_bdg = _bdg + "\u0030"
							}
							_abae := _cece.SetShadingByName(_bdg, _abf.ToPdfObject())
							if _abae != nil {
								_df.Log.Debug("E\u0052\u0052\u004f\u0052 S\u0065t\u0020\u0073\u0068\u0061\u0064i\u006e\u0067\u003a\u0020\u0025\u0076", _abae)
								return _abae
							}
							_egb[*_dce] = _bdg
						} else {
							_df.Log.Debug("\u0053\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
						}
					}
					if _fff, _gfdb := _egb[*_dce]; _gfdb {
						_fad.Params[0] = &_fff
					} else {
						_df.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020S\u0068\u0061\u0064\u0069\u006e\u0067\u0020\u0025\u0073 \u006e\u006f\u0074 \u0066o\u0075\u006e\u0064", *_dce)
					}
				}
			}
		case "\u0067\u0073":
			if len(_fad.Params) == 1 {
				if _cbc, _bfc := _fad.Params[0].(*_fg.PdfObjectName); _bfc {
					if _, _fba := _ccd[*_cbc]; !_fba {
						var _eeeb _fg.PdfObjectName
						_egc, _dad := _fbe.GetExtGState(*_cbc)
						if _dad {
							_eeeb = *_cbc
							_egd := 1
							for {
								_ccb, _ad := _cece.GetExtGState(_eeeb)
								if !_ad || _egc == _ccb {
									break
								}
								_eeeb = _fg.PdfObjectName(_eg.Sprintf("\u0047\u0053\u0025\u0064", _egd))
								_egd++
							}
						}
						_cece.AddExtGState(_eeeb, _egc)
						_ccd[*_cbc] = _eeeb
					}
					_dde := _ccd[*_cbc]
					_fad.Params[0] = &_dde
				}
			}
		}
		*_cce = append(*_cce, _fad)
	}
	return nil
}

// RotatedSize returns the width and height of the rotated block.
func (_gab *Block) RotatedSize() (float64, float64) {
	_, _, _cd, _bcd := _feda(_gab._bf, _gab._eda, _gab._ag)
	return _cd, _bcd
}

// SetStyleRight sets border style for right side.
func (_aeef *border) SetStyleRight(style CellBorderStyle) { _aeef._bgad = style }

type rgbColor struct{ _fabb, _abab, _gggf float64 }

func (_afce *StyledParagraph) getLineMetrics(_edebg int) (_bedb, _efc, _gacaff float64) {
	if _afce._bfaa == nil || len(_afce._bfaa) == 0 {
		_afce.wrapText()
	}
	if _edebg < 0 || _edebg > len(_afce._bfaa)-1 {
		_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020p\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u0020\u006c\u0069\u006e\u0065 \u0069\u006e\u0064\u0065\u0078\u0020\u0025\u0064\u002e\u0020\u0052\u0065tu\u0072\u006e\u0069\u006e\u0067\u0020\u0030\u002c\u0020\u0030", _edebg)
		return 0, 0, 0
	}
	_efbb := _afce._bfaa[_edebg]
	for _, _fdea := range _efbb {
		_eedgf, _fafa := _fdea.Style.Font.GetFontDescriptor()
		if _fafa != nil {
			_df.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020U\u006e\u0061\u0062\u006ce t\u006f g\u0065\u0074\u0020\u0066\u006f\u006e\u0074 d\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
		}
		var _cdeg, _ccdb float64
		if _eedgf != nil {
			if _cdeg, _fafa = _eedgf.GetCapHeight(); _fafa != nil {
				_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0043\u0061\u0070\u0048\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _fafa)
			}
			if _ccdb, _fafa = _eedgf.GetDescent(); _fafa != nil {
				_df.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020U\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0067\u0065t\u0020\u0066\u006f\u006e\u0074\u0020\u0044\u0065\u0073\u0063\u0065\u006et\u003a\u0020\u0025\u0076", _fafa)
			}
		}
		if int(_cdeg) <= 0 {
			_df.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0043\u0061p\u0020\u0048\u0065ig\u0068\u0074\u0020\u006e\u006f\u0074 \u0061\u0076\u0061\u0069\u006c\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065\u0074t\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u00310\u0030\u0030")
			_cdeg = 1000
		}
		if _dbba := _cdeg / 1000.0 * _fdea.Style.FontSize; _dbba > _bedb {
			_bedb = _dbba
		}
		if _cdagf := _ccdb / 1000.0 * _fdea.Style.FontSize; _cdagf < _gacaff {
			_gacaff = _cdagf
		}
		if _gfbf := _fdea.Style.FontSize; _gfbf > _efc {
			_efc = _gfbf
		}
	}
	return _bedb, _efc, _gacaff
}

// MoveRight moves the drawing context right by relative displacement dx (negative goes left).
func (_cagfg *Creator) MoveRight(dx float64) { _cagfg._bfec.X += dx }

// SetText replaces all the text of the paragraph with the specified one.
func (_fbde *StyledParagraph) SetText(text string) *TextChunk {
	_fbde.Reset()
	return _fbde.Append(text)
}

// Write output of creator to io.Writer interface.
func (_gbgdb *Creator) Write(ws _ef.Writer) error {
	if _fgfe := _gbgdb.Finalize(); _fgfe != nil {
		return _fgfe
	}
	_bag := _f.NewPdfWriter()
	_bag.SetOptimizer(_gbgdb._agg)
	if _gbgdb._eafg != nil {
		_eedc := _bag.SetForms(_gbgdb._eafg)
		if _eedc != nil {
			_df.Log.Debug("F\u0061\u0069\u006c\u0075\u0072\u0065\u003a\u0020\u0025\u0076", _eedc)
			return _eedc
		}
	}
	if _gbgdb._afed != nil {
		_bag.AddOutlineTree(_gbgdb._afed)
	} else if _gbgdb._dfd != nil && _gbgdb.AddOutlines {
		_bag.AddOutlineTree(&_gbgdb._dfd.ToPdfOutline().PdfOutlineTreeNode)
	}
	if _gbgdb._bcca != nil {
		if _caac := _bag.SetPageLabels(_gbgdb._bcca); _caac != nil {
			_df.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020C\u006f\u0075\u006c\u0064 no\u0074 s\u0065\u0074\u0020\u0070\u0061\u0067\u0065 l\u0061\u0062\u0065\u006c\u0073\u003a\u0020%\u0076", _caac)
			return _caac
		}
	}
	if _gbgdb._dedb != nil {
		for _, _bbd := range _gbgdb._dedb {
			_ggegg := _bbd.SubsetRegistered()
			if _ggegg != nil {
				_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043\u006f\u0075\u006c\u0064\u0020\u006e\u006ft\u0020s\u0075\u0062\u0073\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0025\u0076", _ggegg)
				return _ggegg
			}
		}
	}
	if _gbgdb._eege != nil {
		_gacf := _gbgdb._eege(&_bag)
		if _gacf != nil {
			_df.Log.Debug("F\u0061\u0069\u006c\u0075\u0072\u0065\u003a\u0020\u0025\u0076", _gacf)
			return _gacf
		}
	}
	for _, _abea := range _gbgdb._gadb {
		_abfd := _bag.AddPage(_abea)
		if _abfd != nil {
			_df.Log.Error("\u0046\u0061\u0069\u006ced\u0020\u0074\u006f\u0020\u0061\u0064\u0064\u0020\u0050\u0061\u0067\u0065\u003a\u0020%\u0076", _abfd)
			return _abfd
		}
	}
	_bceb := _bag.Write(ws)
	if _bceb != nil {
		return _bceb
	}
	return nil
}
func (_cbfca *templateProcessor) parseTextChunk(_gfgd *templateNode) (interface{}, error) {
	_ffacb := NewTextChunk("", _cbfca.creator.NewTextStyle())
	for _, _afgeba := range _gfgd._ddgde.Attr {
		_agcad := _afgeba.Value
		switch _cacgd := _afgeba.Name.Local; _cacgd {
		case "\u0063\u006f\u006co\u0072":
			_ffacb.Style.Color = _cbfca.parseColorAttr(_cacgd, _agcad)
		case "\u006f\u0075\u0074\u006c\u0069\u006e\u0065\u002d\u0063\u006f\u006c\u006f\u0072":
			_ffacb.Style.OutlineColor = _cbfca.parseColorAttr(_cacgd, _agcad)
		case "\u0066\u006f\u006e\u0074":
			_ffacb.Style.Font = _cbfca.parseFontAttr(_cacgd, _agcad)
		case "\u0066o\u006e\u0074\u002d\u0073\u0069\u007ae":
			_ffacb.Style.FontSize = _cbfca.parseFloatAttr(_cacgd, _agcad)
		case "\u006f\u0075\u0074l\u0069\u006e\u0065\u002d\u0073\u0069\u007a\u0065":
			_ffacb.Style.OutlineSize = _cbfca.parseFloatAttr(_cacgd, _agcad)
		case "\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u002d\u0073\u0070a\u0063\u0069\u006e\u0067":
			_ffacb.Style.CharSpacing = _cbfca.parseFloatAttr(_cacgd, _agcad)
		case "\u0068o\u0072i\u007a\u006f\u006e\u0074\u0061l\u002d\u0073c\u0061\u006c\u0069\u006e\u0067":
			_ffacb.Style.HorizontalScaling = _cbfca.parseFloatAttr(_cacgd, _agcad)
		case "\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067-\u006d\u006f\u0064\u0065":
			_ffacb.Style.RenderingMode = _cbfca.parseTextRenderingModeAttr(_cacgd, _agcad)
		case "\u0075n\u0064\u0065\u0072\u006c\u0069\u006ee":
			_ffacb.Style.Underline = _cbfca.parseBoolAttr(_cacgd, _agcad)
		case "\u0075n\u0064e\u0072\u006c\u0069\u006e\u0065\u002d\u0063\u006f\u006c\u006f\u0072":
			_ffacb.Style.UnderlineStyle.Color = _cbfca.parseColorAttr(_cacgd, _agcad)
		case "\u0075\u006ed\u0065\u0072\u006ci\u006e\u0065\u002d\u006f\u0066\u0066\u0073\u0065\u0074":
			_ffacb.Style.UnderlineStyle.Offset = _cbfca.parseFloatAttr(_cacgd, _agcad)
		case "\u0075\u006e\u0064\u0065rl\u0069\u006e\u0065\u002d\u0074\u0068\u0069\u0063\u006b\u006e\u0065\u0073\u0073":
			_ffacb.Style.UnderlineStyle.Thickness = _cbfca.parseFloatAttr(_cacgd, _agcad)
		case "\u0074e\u0078\u0074\u002d\u0072\u0069\u0073e":
			_ffacb.Style.TextRise = _cbfca.parseFloatAttr(_cacgd, _agcad)
		default:
			_df.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0063\u0068\u0075\u006e\u006b\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e", _cacgd)
		}
	}
	return _ffacb, nil
}

// GetMargins returns the Paragraph's margins: left, right, top, bottom.
func (_eebd *StyledParagraph) GetMargins() (float64, float64, float64, float64) {
	return _eebd._afdba.Left, _eebd._afdba.Right, _eebd._afdba.Top, _eebd._afdba.Bottom
}

// SetWidthBottom sets border width for bottom.
func (_ded *border) SetWidthBottom(bw float64) { _ded._gbe = bw }

// NewChapter creates a new chapter with the specified title as the heading.
func (_cfa *Creator) NewChapter(title string) *Chapter {
	_cfa._dedf++
	_eced := _cfa.NewTextStyle()
	_eced.FontSize = 16
	return _bac(nil, _cfa._cddg, _cfa._dfd, title, _cfa._dedf, _eced)
}

// FitMode defines resizing options of an object inside a container.
type FitMode int

// SetAnnotation sets a annotation on a TextChunk.
func (_eeddb *TextChunk) SetAnnotation(annotation *_f.PdfAnnotation) { _eeddb._ceddd = annotation }

// NewChart creates a new creator drawable based on the provided
// unichart chart component.
func NewChart(chart _ab.ChartRenderable) *Chart { return _adb(chart) }

// NewImageFromData creates an Image from image data.
func (_dgdd *Creator) NewImageFromData(data []byte) (*Image, error) { return _bcae(data) }
func (_dfg *Block) transform(_edc _edf.Matrix) {
	_bbg := _eee.NewContentCreator().Add_cm(_edc[0], _edc[1], _edc[3], _edc[4], _edc[6], _edc[7]).Operations()
	*_dfg._ga = append(*_bbg, *_dfg._ga...)
	_dfg._ga.WrapIfNeeded()
}

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_fbfc *List) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var _bdda float64
	var _aebd []*StyledParagraph
	for _, _acdf := range _fbfc._eddg {
		_cgbgg := _ggaa(_fbfc._edae)
		_cgbgg.SetEnableWrap(false)
		_cgbgg.SetTextAlignment(TextAlignmentRight)
		_cgbgg.Append(_acdf._gade.Text).Style = _acdf._gade.Style
		_effe := _cgbgg.getTextWidth() / 1000.0 / ctx.Width
		if _bdda < _effe {
			_bdda = _effe
		}
		_aebd = append(_aebd, _cgbgg)
	}
	_dadg := _ecef(2)
	_dadg.SetColumnWidths(_bdda, 1-_bdda)
	_dadg.SetMargins(_fbfc._dgab, 0, 0, 0)
	for _fedg, _acca := range _fbfc._eddg {
		_dffg := _dadg.NewCell()
		_dffg.SetIndent(0)
		_dffg.SetContent(_aebd[_fedg])
		_dffg = _dadg.NewCell()
		_dffg.SetIndent(0)
		_dffg.SetContent(_acca._cdef)
	}
	return _dadg.GeneratePageBlocks(ctx)
}

// Width returns the width of the line.
// NOTE: Depending on the fit mode the line is set to use, its width may be
// calculated at runtime (e.g. when using FitModeFillWidth).
func (_aeadce *Line) Width() float64 { return _a.Abs(_aeadce._dccb - _aeadce._ggfd) }

// Title returns the title of the invoice.
func (_gfea *Invoice) Title() string { return _gfea._bgdd }

// Logo returns the logo of the invoice.
func (_fadg *Invoice) Logo() *Image { return _fadg._fbgcc }

// SetMargins sets the Paragraph's margins.
func (_ccfe *Paragraph) SetMargins(left, right, top, bottom float64) {
	_ccfe._eacf.Left = left
	_ccfe._eacf.Right = right
	_ccfe._eacf.Top = top
	_ccfe._eacf.Bottom = bottom
}

// DrawHeader sets a function to draw a header on created output pages.
func (_bgd *Creator) DrawHeader(drawHeaderFunc func(_fgff *Block, _cgec HeaderFunctionArgs)) {
	_bgd._eacd = drawHeaderFunc
}

// AddExternalLink adds a new external link to the paragraph.
// The text parameter represents the text that is displayed and the url
// parameter sets the destionation of the link.
func (_edaa *StyledParagraph) AddExternalLink(text, url string) *TextChunk {
	_ebce := NewTextChunk(text, _edaa._bdbf)
	_ebce._ceddd = _bfbbe(url)
	return _edaa.appendChunk(_ebce)
}

// SetAngle sets Image rotation angle in degrees.
func (_gggda *Image) SetAngle(angle float64) { _gggda._afgea = angle }
func _bcafd(_fgbee *templateProcessor, _babdd *templateNode) (interface{}, error) {
	return _fgbee.parseStyledParagraph(_babdd)
}

// SetWidth set the Image's document width to specified w. This does not change the raw image data, i.e.
// no actual scaling of data is performed. That is handled by the PDF viewer.
func (_afa *Image) SetWidth(w float64) { _afa._bcfd = w }

// SetLineOpacity sets the line opacity.
func (_daec *Polyline) SetLineOpacity(opacity float64) { _daec._afbcd = opacity }
func (_ddfce *StyledParagraph) getMaxLineWidth() float64 {
	if _ddfce._bfaa == nil || len(_ddfce._bfaa) == 0 {
		_ddfce.wrapText()
	}
	var _gfdc float64
	for _, _eedcf := range _ddfce._bfaa {
		_fdfb := _ddfce.getTextLineWidth(_eedcf)
		if _fdfb > _gfdc {
			_gfdc = _fdfb
		}
	}
	return _gfdc
}

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_cdgfc *Invoice) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gdcc := ctx
	_afaa := []func(_aafd DrawContext) ([]*Block, DrawContext, error){_cdgfc.generateHeaderBlocks, _cdgfc.generateInformationBlocks, _cdgfc.generateLineBlocks, _cdgfc.generateTotalBlocks, _cdgfc.generateNoteBlocks}
	var _ddaab []*Block
	for _, _cdcd := range _afaa {
		_dfbd, _fccfg, _gadbb := _cdcd(ctx)
		if _gadbb != nil {
			return _ddaab, ctx, _gadbb
		}
		if len(_ddaab) == 0 {
			_ddaab = _dfbd
		} else if len(_dfbd) > 0 {
			_ddaab[len(_ddaab)-1].mergeBlocks(_dfbd[0])
			_ddaab = append(_ddaab, _dfbd[1:]...)
		}
		ctx = _fccfg
	}
	if _cdgfc._dfgf.IsRelative() {
		ctx.X = _gdcc.X
	}
	if _cdgfc._dfgf.IsAbsolute() {
		return _ddaab, _gdcc, nil
	}
	return _ddaab, ctx, nil
}

// NewTextChunk returns a new text chunk instance.
func NewTextChunk(text string, style TextStyle) *TextChunk {
	return &TextChunk{Text: text, Style: style}
}

type templateNode struct {
	_cedba interface{}
	_ddgde _cc.StartElement
	_dfdd  *templateNode
}

// Rows returns the total number of rows the table has.
func (_gfeae *Table) Rows() int { return _gfeae._fdff }

// EnableFontSubsetting enables font subsetting for `font` when the creator output is written to file.
// Embeds only the subset of the runes/glyphs that are actually used to display the file.
// Subsetting can reduce the size of fonts significantly.
func (_baef *Creator) EnableFontSubsetting(font *_f.PdfFont) { _baef._dedb = append(_baef._dedb, font) }
func _fcabb(_dgbfd *_f.PdfFont) TextStyle {
	return TextStyle{Color: ColorRGBFrom8bit(0, 0, 238), Font: _dgbfd, FontSize: 10, OutlineSize: 1, HorizontalScaling: DefaultHorizontalScaling, UnderlineStyle: TextDecorationLineStyle{Offset: 1, Thickness: 1}}
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
func (_gbeb *Chapter) Add(d Drawable) error {
	if Drawable(_gbeb) == d {
		_df.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u006e\u006f\u0074 \u0061\u0064\u0064\u0020\u0069\u0074\u0073\u0065\u006c\u0066")
		return _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	switch _ead := d.(type) {
	case *Paragraph, *StyledParagraph, *Image, *Chart, *Table, *Division, *List, *Rectangle, *Ellipse, *Line, *Block, *PageBreak, *Chapter:
		_gbeb._gecb = append(_gbeb._gecb, d)
	case containerDrawable:
		_fbcb, _aaf := _ead.ContainerComponent(_gbeb)
		if _aaf != nil {
			return _aaf
		}
		_gbeb._gecb = append(_gbeb._gecb, _fbcb)
	default:
		_df.Log.Debug("\u0055n\u0073u\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u003a\u0020\u0025\u0054", d)
		return _d.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	return nil
}

// NewLine creates a new line between (x1, y1) to (x2, y2),
// using default attributes.
// NOTE: In relative positioning mode, `x1` and `y1` are calculated using the
// current context and `x2`, `y2` are used only to calculate the position of
// the second point in relation to the first one (used just as a measurement
// of size). Furthermore, when the fit mode is set to fill the context width,
// `x2` is set to the right edge coordinate of the context.
func (_eaec *Creator) NewLine(x1, y1, x2, y2 float64) *Line { return _agddb(x1, y1, x2, y2) }
func (_aecf *Invoice) generateHeaderBlocks(_aacd DrawContext) ([]*Block, DrawContext, error) {
	_debgb := _ggaa(_aecf._acgbf)
	_debgb.SetEnableWrap(true)
	_debgb.Append(_aecf._bgdd)
	_gefb := _ecef(2)
	if _aecf._fbgcc != nil {
		_gbabc := _gefb.NewCell()
		_gbabc.SetHorizontalAlignment(CellHorizontalAlignmentLeft)
		_gbabc.SetVerticalAlignment(CellVerticalAlignmentMiddle)
		_gbabc.SetIndent(0)
		_gbabc.SetContent(_aecf._fbgcc)
		_aecf._fbgcc.ScaleToHeight(_debgb.Height() + 20)
	} else {
		_gefb.SkipCells(1)
	}
	_dfffa := _gefb.NewCell()
	_dfffa.SetHorizontalAlignment(CellHorizontalAlignmentRight)
	_dfffa.SetVerticalAlignment(CellVerticalAlignmentMiddle)
	_dfffa.SetContent(_debgb)
	return _gefb.GeneratePageBlocks(_aacd)
}

// The Image type is used to draw an image onto PDF.
type Image struct {
	_aabc        *_f.XObjectImage
	_aaca        *_f.Image
	_afgea       float64
	_bcfd, _gddc float64
	_cbbgg, _gbf float64
	_ccec        Positioning
	_cedbd       HorizontalAlignment
	_cgee        float64
	_ddab        float64
	_edgb        float64
	_dbfe        Margins
	_bcfc, _efaa float64
	_facba       _fg.StreamEncoder
	_aedff       FitMode
}

// VectorDrawable is a Drawable with a specified width and height.
type VectorDrawable interface {
	Drawable

	// Width returns the width of the Drawable.
	Width() float64

	// Height returns the height of the Drawable.
	Height() float64
}

// SetLinePageStyle sets the style for the page part of all new lines
// of the table of contents.
func (_cgagg *TOC) SetLinePageStyle(style TextStyle) { _cgagg._ffed = style }

// SetMargins sets the margins for the Image (in relative mode): left, right, top, bottom.
func (_befe *Image) SetMargins(left, right, top, bottom float64) {
	_befe._dbfe.Left = left
	_befe._dbfe.Right = right
	_befe._dbfe.Top = top
	_befe._dbfe.Bottom = bottom
}
func (_gfffd *Invoice) generateInformationBlocks(_eeadd DrawContext) ([]*Block, DrawContext, error) {
	_eebf := _ggaa(_gfffd._ffdf)
	_eebf.SetMargins(0, 0, 0, 20)
	_baab := _gfffd.drawAddress(_gfffd._bgfe)
	_baab = append(_baab, _eebf)
	_baab = append(_baab, _gfffd.drawAddress(_gfffd._ebca)...)
	_eccd := _cdde()
	for _, _acgf := range _baab {
		_eccd.Add(_acgf)
	}
	_fggae := _gfffd.drawInformation()
	_abdf := _ecef(2)
	_abdf.SetMargins(0, 0, 25, 0)
	_fbcf := _abdf.NewCell()
	_fbcf.SetIndent(0)
	_fbcf.SetContent(_eccd)
	_fbcf = _abdf.NewCell()
	_fbcf.SetContent(_fggae)
	return _abdf.GeneratePageBlocks(_eeadd)
}

type componentRenderer interface {
	Draw(_gcfc Drawable) error
}

func _ffab(_bdgdf float64, _aecc float64) float64 { return _a.Round(_bdgdf/_aecc) * _aecc }

// AddTextItem appends a new item with the specified text to the list.
// The method creates a styled paragraph with the specified text and returns
// it so that the item style can be customized.
// The method also returns the marker used for the newly added item.
// The marker object can be used to change the text and style of the marker
// for the current item.
func (_ageb *List) AddTextItem(text string) (*StyledParagraph, *TextChunk, error) {
	_fdaf := _ggaa(_ageb._edae)
	_fdaf.Append(text)
	_caff, _dgce := _ageb.Add(_fdaf)
	return _fdaf, _caff, _dgce
}
func _aeac(_aafa, _fabd, _fbgg, _dgdf float64) *Rectangle {
	return &Rectangle{_bagb: _aafa, _cbe: _fabd, _ecfc: _fbgg, _ecca: _dgdf, _afcb: ColorBlack, _bffaf: 1.0, _ffcbg: 1.0, _beeac: 1.0, _fcbgb: PositionAbsolute}
}

// Terms returns the terms and conditions section of the invoice as a
// title-content pair.
func (_ddad *Invoice) Terms() (string, string) { return _ddad._dcdf[0], _ddad._dcdf[1] }

// SetBorderColor sets the border color for the path.
func (_bafg *FilledCurve) SetBorderColor(color Color) { _bafg._aac = color }

// NewSubchapter creates a new child chapter with the specified title.
func (_eeed *Chapter) NewSubchapter(title string) *Chapter {
	_ade := _dadda(_eeed._dbf._ceda)
	_ade.FontSize = 14
	_eeed._gcd++
	_bed := _bac(_eeed, _eeed._gdba, _eeed._cdfd, title, _eeed._gcd, _ade)
	_eeed.Add(_bed)
	return _bed
}
func _fabgd(_aaga, _cdca TextStyle) *Invoice {
	_facd := &Invoice{_bgdd: "\u0049N\u0056\u004f\u0049\u0043\u0045", _ceb: "\u002c\u0020", _ffdf: _aaga, _gfdd: _cdca}
	_facd._bgfe = &InvoiceAddress{Separator: _facd._ceb}
	_facd._ebca = &InvoiceAddress{Heading: "\u0042i\u006c\u006c\u0020\u0074\u006f", Separator: _facd._ceb}
	_fegd := ColorRGBFrom8bit(245, 245, 245)
	_fdgg := ColorRGBFrom8bit(155, 155, 155)
	_facd._acgbf = _cdca
	_facd._acgbf.Color = _fdgg
	_facd._acgbf.FontSize = 20
	_facd._ddf = _aaga
	_facd._eeea = _cdca
	_facd._egf = _aaga
	_facd._cacg = _cdca
	_facd._fcabd = _facd.NewCellProps()
	_facd._fcabd.BackgroundColor = _fegd
	_facd._fcabd.TextStyle = _cdca
	_facd._cbgg = _facd.NewCellProps()
	_facd._cbgg.TextStyle = _cdca
	_facd._cbgg.BackgroundColor = _fegd
	_facd._cbgg.BorderColor = _fegd
	_facd._defdd = _facd.NewCellProps()
	_facd._defdd.BorderColor = _fegd
	_facd._defdd.BorderSides = []CellBorderSide{CellBorderSideBottom}
	_facd._defdd.Alignment = CellHorizontalAlignmentRight
	_facd._fabge = _facd.NewCellProps()
	_facd._fabge.Alignment = CellHorizontalAlignmentRight
	_facd._edeb = [2]*InvoiceCell{_facd.newCell("\u0049\u006e\u0076\u006f\u0069\u0063\u0065\u0020\u006eu\u006d\u0062\u0065\u0072", _facd._fcabd), _facd.newCell("", _facd._fcabd)}
	_facd._fgcge = [2]*InvoiceCell{_facd.newCell("\u0044\u0061\u0074\u0065", _facd._fcabd), _facd.newCell("", _facd._fcabd)}
	_facd._debe = [2]*InvoiceCell{_facd.newCell("\u0044\u0075\u0065\u0020\u0044\u0061\u0074\u0065", _facd._fcabd), _facd.newCell("", _facd._fcabd)}
	_facd._faeg = [2]*InvoiceCell{_facd.newCell("\u0053\u0075\u0062\u0074\u006f\u0074\u0061\u006c", _facd._fabge), _facd.newCell("", _facd._fabge)}
	_gdef := _facd._fabge
	_gdef.TextStyle = _cdca
	_gdef.BackgroundColor = _fegd
	_gdef.BorderColor = _fegd
	_facd._defg = [2]*InvoiceCell{_facd.newCell("\u0054\u006f\u0074a\u006c", _gdef), _facd.newCell("", _gdef)}
	_facd._dbe = [2]string{"\u004e\u006f\u0074e\u0073", ""}
	_facd._dcdf = [2]string{"T\u0065r\u006d\u0073\u0020\u0061\u006e\u0064\u0020\u0063o\u006e\u0064\u0069\u0074io\u006e\u0073", ""}
	_facd._afeb = []*InvoiceCell{_facd.newColumn("D\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e", CellHorizontalAlignmentLeft), _facd.newColumn("\u0051\u0075\u0061\u006e\u0074\u0069\u0074\u0079", CellHorizontalAlignmentRight), _facd.newColumn("\u0055\u006e\u0069\u0074\u0020\u0070\u0072\u0069\u0063\u0065", CellHorizontalAlignmentRight), _facd.newColumn("\u0041\u006d\u006f\u0075\u006e\u0074", CellHorizontalAlignmentRight)}
	return _facd
}

// SetBorderOpacity sets the border opacity.
func (_aefc *CurvePolygon) SetBorderOpacity(opacity float64) { _aefc._cedec = opacity }

const (
	CellVerticalAlignmentTop CellVerticalAlignment = iota
	CellVerticalAlignmentMiddle
	CellVerticalAlignmentBottom
)

// ScaleToHeight scale Image to a specified height h, maintaining the aspect ratio.
func (_feeg *Image) ScaleToHeight(h float64) {
	_fecf := _feeg._bcfd / _feeg._gddc
	_feeg._gddc = h
	_feeg._bcfd = h * _fecf
}

// InvoiceCellProps holds all style properties for an invoice cell.
type InvoiceCellProps struct {
	TextStyle       TextStyle
	Alignment       CellHorizontalAlignment
	BackgroundColor Color
	BorderColor     Color
	BorderWidth     float64
	BorderSides     []CellBorderSide
}

func (_abga *Image) rotatedSize() (float64, float64) {
	_abefec := _abga._bcfd
	_caeca := _abga._gddc
	_dccad := _abga._afgea
	if _dccad == 0 {
		return _abefec, _caeca
	}
	_cegg := _cec.Path{Points: []_cec.Point{_cec.NewPoint(0, 0).Rotate(_dccad), _cec.NewPoint(_abefec, 0).Rotate(_dccad), _cec.NewPoint(0, _caeca).Rotate(_dccad), _cec.NewPoint(_abefec, _caeca).Rotate(_dccad)}}.GetBoundingBox()
	return _cegg.Width, _cegg.Height
}

// SetBorderWidth sets the border width.
func (_abec *CurvePolygon) SetBorderWidth(borderWidth float64) { _abec._dbca.BorderWidth = borderWidth }

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
	_aefd []VectorDrawable
	_edfc Positioning
	_fccf Margins
	_gfeb Margins
	_ddda bool
	_dcd  bool
	_bgeg *Background
}

// SetColorLeft sets border color for left.
func (_bcc *border) SetColorLeft(col Color) { _bcc._bfe = col }
func _fegf(_dcgaf *_f.PdfRectangle, _dfbddb _edf.Matrix) *_f.PdfRectangle {
	var _bcbfa _f.PdfRectangle
	_bcbfa.Llx, _bcbfa.Lly = _dfbddb.Transform(_dcgaf.Llx, _dcgaf.Lly)
	_bcbfa.Urx, _bcbfa.Ury = _dfbddb.Transform(_dcgaf.Urx, _dcgaf.Ury)
	_bcbfa.Normalize()
	return &_bcbfa
}

// AddressStyle returns the style properties used to render the content of
// the invoice address sections.
func (_ecbe *Invoice) AddressStyle() TextStyle { return _ecbe._ddf }

// Height returns the current page height.
func (_afbcc *Creator) Height() float64 { return _afbcc._agca }

// SetIndent sets the cell's left indent.
func (_eggg *TableCell) SetIndent(indent float64) { _eggg._cdgg = indent }

// SetLineHeight sets the line height (1.0 default).
func (_adafa *StyledParagraph) SetLineHeight(lineheight float64) { _adafa._bged = lineheight }
func (_gadg *Creator) setActivePage(_cfge *_f.PdfPage)           { _gadg._efde = _cfge }

// AddLine adds a new line with the provided style to the table of contents.
func (_efaed *TOC) AddLine(line *TOCLine) *TOCLine {
	if line == nil {
		return nil
	}
	_efaed._fbagb = append(_efaed._fbagb, line)
	return line
}

// GeneratePageBlocks draws the polygon on a new block representing the page.
// Implements the Drawable interface.
func (_bgeb *Polygon) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_aebaa := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_gdgf, _gdfb := _aebaa.setOpacity(_bgeb._fead, _bgeb._afab)
	if _gdfb != nil {
		return nil, ctx, _gdfb
	}
	_cbcf := _bgeb._ccbeg
	_cbcf.FillEnabled = _cbcf.FillColor != nil
	_cbcf.BorderEnabled = _cbcf.BorderColor != nil && _cbcf.BorderWidth > 0
	_cgedd := _cbcf.Points
	for _dcgg := range _cgedd {
		for _bebge := range _cgedd[_dcgg] {
			_efddc := &_cgedd[_dcgg][_bebge]
			_efddc.Y = ctx.PageHeight - _efddc.Y
		}
	}
	_aadg, _, _gdfb := _cbcf.Draw(_gdgf)
	if _gdfb != nil {
		return nil, ctx, _gdfb
	}
	if _gdfb = _aebaa.addContentsByString(string(_aadg)); _gdfb != nil {
		return nil, ctx, _gdfb
	}
	return []*Block{_aebaa}, ctx, nil
}
func (_bcega *Creator) getActivePage() *_f.PdfPage {
	if _bcega._efde == nil {
		if len(_bcega._gadb) == 0 {
			return nil
		}
		return _bcega._gadb[len(_bcega._gadb)-1]
	}
	return _bcega._efde
}
func _gfff(_edfb, _fcab *_f.PdfPageResources) error {
	_fdb, _ := _edfb.GetColorspaces()
	if _fdb != nil && len(_fdb.Colorspaces) > 0 {
		for _ebf, _afbc := range _fdb.Colorspaces {
			_beg := *_fg.MakeName(_ebf)
			if _fcab.HasColorspaceByName(_beg) {
				continue
			}
			_aega := _fcab.SetColorspaceByName(_beg, _afbc)
			if _aega != nil {
				return _aega
			}
		}
	}
	return nil
}

// MoveDown moves the drawing context down by relative displacement dy (negative goes up).
func (_eece *Creator) MoveDown(dy float64) { _eece._bfec.Y += dy }

// PageFinalizeFunctionArgs holds the input arguments provided to the page
// finalize callback function which can be set using Creator.PageFinalize.
type PageFinalizeFunctionArgs struct {
	PageNum    int
	PageWidth  float64
	PageHeight float64
	TOCPages   int
	TotalPages int
}

func _gcffd(_gccf TextStyle) *List {
	return &List{_fbfe: TextChunk{Text: "\u2022\u0020", Style: _gccf}, _dgab: 0, _ggbf: true, _acdb: PositionRelative, _edae: _gccf}
}

type border struct {
	_cdfc     float64
	_aaab     float64
	_dgf      float64
	_gbgc     float64
	_edg      Color
	_bfe      Color
	_afeg     float64
	_bdd      Color
	_gbe      float64
	_bfae     Color
	_fe       float64
	_gad      Color
	_ebc      float64
	LineStyle _cec.LineStyle
	_afg      CellBorderStyle
	_bgad     CellBorderStyle
	_faaa     CellBorderStyle
	_aeba     CellBorderStyle
}

// Positioning returns the type of positioning the line is set to use.
func (_gafbc *Line) Positioning() Positioning { return _gafbc._edcee }

// GeneratePageBlocks generates the page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages. Implements the Drawable interface.
func (_agfcd *StyledParagraph) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_ecga := ctx
	var _gfaa []*Block
	_gdfe := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _agfcd._geagd.IsRelative() {
		ctx.X += _agfcd._afdba.Left
		ctx.Y += _agfcd._afdba.Top
		ctx.Width -= _agfcd._afdba.Left + _agfcd._afdba.Right
		ctx.Height -= _agfcd._afdba.Top
		_agfcd.SetWidth(ctx.Width)
	} else {
		if int(_agfcd._bbef) <= 0 {
			_agfcd.SetWidth(_agfcd.getTextWidth() / 1000.0)
		}
		ctx.X = _agfcd._ggbc
		ctx.Y = _agfcd._aeegd
	}
	if _agfcd._bccd != nil {
		_agfcd._bccd(_agfcd, ctx)
	}
	if _ecaa := _agfcd.wrapText(); _ecaa != nil {
		return nil, ctx, _ecaa
	}
	_adfe := _agfcd._bfaa
	for {
		_adfa, _faaf, _aegbc := _gebb(_gdfe, _agfcd, _adfe, ctx)
		if _aegbc != nil {
			_df.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _aegbc)
			return nil, ctx, _aegbc
		}
		ctx = _adfa
		_gfaa = append(_gfaa, _gdfe)
		if _adfe = _faaf; len(_faaf) == 0 {
			break
		}
		_gdfe = NewBlock(ctx.PageWidth, ctx.PageHeight)
		ctx.Page++
		_adfa = ctx
		_adfa.Y = ctx.Margins.Top
		_adfa.X = ctx.Margins.Left + _agfcd._afdba.Left
		_adfa.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
		_adfa.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _agfcd._afdba.Left - _agfcd._afdba.Right
		ctx = _adfa
	}
	if _agfcd._geagd.IsRelative() {
		ctx.Y += _agfcd._afdba.Bottom
		ctx.Height -= _agfcd._afdba.Bottom
		if !ctx.Inline {
			ctx.X = _ecga.X
			ctx.Width = _ecga.Width
		}
		return _gfaa, ctx, nil
	}
	return _gfaa, _ecga, nil
}

// SetWidthLeft sets border width for left.
func (_fbd *border) SetWidthLeft(bw float64) { _fbd._afeg = bw }
func (_afgg *StyledParagraph) wrapWordChunks() {
	if !_afgg._bggd {
		return
	}
	var _aacc []*TextChunk
	for _, _aegb := range _afgg._eedad {
		_bagbg := []rune(_aegb.Text)
		if len(_aacc) > 0 {
			if len(_bagbg) == 1 && _ee.IsPunct(_bagbg[0]) {
				_cegfd := []rune(_aacc[len(_aacc)-1].Text)
				_aacc[len(_aacc)-1].Text = string(append(_cegfd, _bagbg[0]))
				continue
			} else {
				_, _gedc := _ae.Atoi(_aegb.Text)
				if _gedc == nil {
					_gdfad := []rune(_aacc[len(_aacc)-1].Text)
					_faacc := len(_gdfad)
					_, _caaac := _ae.Atoi(string(_gdfad[_faacc-2]))
					if _caaac == nil && _ee.IsPunct(_gdfad[_faacc-1]) {
						_aacc[len(_aacc)-1].Text = string(append(_gdfad, _bagbg...))
						continue
					}
				}
			}
		}
		_bcgb, _eeaf := _adaeb(_aegb.Text)
		if _eeaf != nil {
			_df.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0062\u0072\u0065\u0061\u006b\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020w\u006f\u0072\u0064\u0073\u003a\u0020\u0025\u0076", _eeaf)
			_bcgb = []string{_aegb.Text}
		}
		for _, _cggdc := range _bcgb {
			_gbaa := NewTextChunk(_cggdc, _aegb.Style)
			_aacc = append(_aacc, _gbaa)
		}
	}
	if len(_aacc) > 0 {
		_afgg._eedad = _aacc
	}
}
func (_ebbb *Creator) newPage() *_f.PdfPage {
	_cfde := _f.NewPdfPage()
	_bdfd := _ebbb._dda[0]
	_fabg := _ebbb._dda[1]
	_gcc := _f.PdfRectangle{Llx: 0, Lly: 0, Urx: _bdfd, Ury: _fabg}
	_cfde.MediaBox = &_gcc
	_ebbb._gee = _bdfd
	_ebbb._agca = _fabg
	_ebbb.initContext()
	return _cfde
}

// SetWidth sets the the Paragraph width. This is essentially the wrapping width, i.e. the width the
// text can extend to prior to wrapping over to next line.
func (_gcab *Paragraph) SetWidth(width float64) { _gcab._bdga = width; _gcab.wrapText() }
func _edfbf(_ecge *Creator, _ecee _ef.Reader, _fefc interface{}, _bbcf *TemplateOptions, _adec componentRenderer) error {
	if _ecge == nil {
		_df.Log.Error("\u0043\u0072\u0065a\u0074\u006f\u0072\u0020i\u006e\u0073\u0074\u0061\u006e\u0063\u0065 \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c\u002e")
		return _aeag
	}
	_ccac := _b.NewBuffer(nil)
	if _, _gacad := _ef.Copy(_ccac, _ecee); _gacad != nil {
		return _gacad
	}
	_cdefe := _g.FuncMap{"\u0064\u0069\u0063\u0074": _agcaa}
	if _bbcf != nil && _bbcf.HelperFuncMap != nil {
		for _aadgg, _gabb := range _bbcf.HelperFuncMap {
			if _, _bcaf := _cdefe[_aadgg]; _bcaf {
				_df.Log.Debug("\u0043\u0061\u006e\u006e\u006f\u0074 \u006f\u0076\u0065r\u0072\u0069\u0064e\u0020\u0062\u0075\u0069\u006c\u0074\u002d\u0069\u006e\u0020`\u0025\u0073\u0060\u0020\u0068el\u0070\u0065\u0072\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _aadgg)
				continue
			}
			_cdefe[_aadgg] = _gabb
		}
	}
	_dfbdc, _ebef := _g.New("").Funcs(_cdefe).Parse(_ccac.String())
	if _ebef != nil {
		return _ebef
	}
	_ccac.Reset()
	if _bbfb := _dfbdc.Execute(_ccac, _fefc); _bbfb != nil {
		return _bbfb
	}
	return _gfde(_ecge, _ccac.Bytes(), _bbcf, _adec).run()
}

// SetSubtotal sets the subtotal of the invoice.
func (_cga *Invoice) SetSubtotal(value string) { _cga._faeg[1].Value = value }

// DashPattern returns the dash pattern of the line.
func (_cggec *Line) DashPattern() (_gbgcf []int64, _ddcb int64) { return _cggec._gbde, _cggec._afbf }

// Draw processes the specified Drawable widget and generates blocks that can
// be rendered to the output document. The generated blocks can span over one
// or more pages. Additional pages are added if the contents go over the current
// page. Each generated block is assigned to the creator page it will be
// rendered to. In order to render the generated blocks to the creator pages,
// call Finalize, Write or WriteToFile.
func (_cde *Creator) Draw(d Drawable) error {
	if _cde.getActivePage() == nil {
		_cde.NewPage()
	}
	_beee, _ada, _gabe := d.GeneratePageBlocks(_cde._bfec)
	if _gabe != nil {
		return _gabe
	}
	if len(_ada._fdcc) > 0 {
		_cde.Errors = append(_cde.Errors, _ada._fdcc...)
	}
	for _eae, _acee := range _beee {
		if _eae > 0 {
			_cde.NewPage()
		}
		_dcce := _cde.getActivePage()
		if _efg, _fbed := _cde._caec[_dcce]; _fbed {
			if _ecb := _efg.mergeBlocks(_acee); _ecb != nil {
				return _ecb
			}
			if _adaf := _gfff(_acee._eb, _efg._eb); _adaf != nil {
				return _adaf
			}
		} else {
			_cde._caec[_dcce] = _acee
		}
	}
	_cde._bfec.X = _ada.X
	_cde._bfec.Y = _ada.Y
	_cde._bfec.Height = _ada.PageHeight - _ada.Y - _ada.Margins.Bottom
	return nil
}

// SetBorderWidth sets the border width.
func (_ccfc *Ellipse) SetBorderWidth(bw float64) { _ccfc._fbec = bw }

// SetBorderRadius sets the radius of the background corners.
func (_aa *Background) SetBorderRadius(topLeft, topRight, bottomLeft, bottomRight float64) {
	_aa.BorderRadiusTopLeft = topLeft
	_aa.BorderRadiusTopRight = topRight
	_aa.BorderRadiusBottomLeft = bottomLeft
	_aa.BorderRadiusBottomRight = bottomRight
}
func _ebeca(_aggfg *templateProcessor, _ddfb *templateNode) (interface{}, error) {
	return _aggfg.parseChart(_ddfb)
}

// SetLink makes the line an internal link.
// The text parameter represents the text that is displayed.
// The user is taken to the specified page, at the specified x and y
// coordinates. Position 0, 0 is at the top left of the page.
func (_gdecf *TOCLine) SetLink(page int64, x, y float64) {
	_gdecf._cegge = x
	_gdecf._gaace = y
	_gdecf._cbfg = page
	_fbaeb := _gdecf._face._bdbf.Color
	_gdecf.Number.Style.Color = _fbaeb
	_gdecf.Title.Style.Color = _fbaeb
	_gdecf.Separator.Style.Color = _fbaeb
	_gdecf.Page.Style.Color = _fbaeb
}

// SetAddressHeadingStyle sets the style properties used to render the
// heading of the invoice address sections.
func (_bedg *Invoice) SetAddressHeadingStyle(style TextStyle) { _bedg._eeea = style }

// LineWidth returns the width of the line.
func (_gfcf *Line) LineWidth() float64 { return _gfcf._ffgbg }

// Opacity returns the opacity of the line.
func (_edfea *Line) Opacity() float64 { return _edfea._cgcbf }

// SetMargins sets the margins of the chart component.
func (_ega *Chart) SetMargins(left, right, top, bottom float64) {
	_ega._daea.Left = left
	_ega._daea.Right = right
	_ega._daea.Top = top
	_ega._daea.Bottom = bottom
}
func (_aacf *templateProcessor) parseInt64Attr(_cega, _ecgdgf string) int64 {
	_df.Log.Debug("\u0050\u0061rs\u0069\u006e\u0067 \u0069\u006e\u0074\u00364 a\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020(`\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _cega, _ecgdgf)
	_egbg, _ := _ae.ParseInt(_ecgdgf, 10, 64)
	return _egbg
}

// GetMargins returns the Chapter's margin: left, right, top, bottom.
func (_gcg *Chapter) GetMargins() (float64, float64, float64, float64) {
	return _gcg._ccf.Left, _gcg._ccf.Right, _gcg._ccf.Top, _gcg._ccf.Bottom
}

// Notes returns the notes section of the invoice as a title-content pair.
func (_fggf *Invoice) Notes() (string, string) { return _fggf._dbe[0], _fggf._dbe[1] }
func _bfbbe(_egceb string) *_f.PdfAnnotation {
	_aggff := _f.NewPdfAnnotationLink()
	_agec := _f.NewBorderStyle()
	_agec.SetBorderWidth(0)
	_aggff.BS = _agec.ToPdfObject()
	_ecbd := _f.NewPdfActionURI()
	_ecbd.URI = _fg.MakeString(_egceb)
	_aggff.SetAction(_ecbd.PdfAction)
	return _aggff.PdfAnnotation
}

// ScaleToHeight scales the Block to a specified height, maintaining the same aspect ratio.
func (_cdc *Block) ScaleToHeight(h float64) { _da := h / _cdc._eda; _cdc.Scale(_da, _da) }
func (_cbf *Block) translate(_eceb, _cgb float64) {
	_cbd := _eee.NewContentCreator().Translate(_eceb, -_cgb).Operations()
	*_cbf._ga = append(*_cbd, *_cbf._ga...)
	_cbf._ga.WrapIfNeeded()
}
func _fdcb(_dddf *Block, _ddabe *Paragraph, _acdg DrawContext) (DrawContext, error) {
	_ecdba := 1
	_acbcg := _fg.PdfObjectName("\u0046\u006f\u006e\u0074" + _ae.Itoa(_ecdba))
	for _dddf._eb.HasFontByName(_acbcg) {
		_ecdba++
		_acbcg = _fg.PdfObjectName("\u0046\u006f\u006e\u0074" + _ae.Itoa(_ecdba))
	}
	_dggee := _dddf._eb.SetFontByName(_acbcg, _ddabe._ceda.ToPdfObject())
	if _dggee != nil {
		return _acdg, _dggee
	}
	_ddabe.wrapText()
	_faad := _eee.NewContentCreator()
	_faad.Add_q()
	_bafc := _acdg.PageHeight - _acdg.Y - _ddabe._gced*_ddabe._abbg
	_faad.Translate(_acdg.X, _bafc)
	if _ddabe._fccd != 0 {
		_faad.RotateDeg(_ddabe._fccd)
	}
	_faad.Add_BT().SetNonStrokingColor(_bcfe(_ddabe._bfbc)).Add_Tf(_acbcg, _ddabe._gced).Add_TL(_ddabe._gced * _ddabe._abbg)
	for _acfb, _debb := range _ddabe._cead {
		if _acfb != 0 {
			_faad.Add_Tstar()
		}
		_aege := []rune(_debb)
		_adde := 0.0
		_gbgf := 0
		for _aaafd, _bced := range _aege {
			if _bced == ' ' {
				_gbgf++
				continue
			}
			if _bced == '\u000A' {
				continue
			}
			_ccbac, _bbaab := _ddabe._ceda.GetRuneMetrics(_bced)
			if !_bbaab {
				_df.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0075\u006e\u0065\u0020\u0069=\u0025\u0064\u0020\u0072\u0075\u006e\u0065=\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0020\u0069n\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0073\u0020\u0025\u0073", _aaafd, _bced, _bced, _ddabe._ceda.BaseFont(), _ddabe._ceda.Subtype())
				return _acdg, _d.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0067\u006c\u0079p\u0068")
			}
			_adde += _ddabe._gced * _ccbac.Wx
		}
		var _eaee []_fg.PdfObject
		_dfbc, _gfg := _ddabe._ceda.GetRuneMetrics(' ')
		if !_gfg {
			return _acdg, _d.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
		}
		_aeca := _dfbc.Wx
		switch _ddabe._badg {
		case TextAlignmentJustify:
			if _gbgf > 0 && _acfb < len(_ddabe._cead)-1 {
				_aeca = (_ddabe._bdga*1000.0 - _adde) / float64(_gbgf) / _ddabe._gced
			}
		case TextAlignmentCenter:
			_ggcbf := _adde + float64(_gbgf)*_aeca*_ddabe._gced
			_bcecd := (_ddabe._bdga*1000.0 - _ggcbf) / 2 / _ddabe._gced
			_eaee = append(_eaee, _fg.MakeFloat(-_bcecd))
		case TextAlignmentRight:
			_gbgcg := _adde + float64(_gbgf)*_aeca*_ddabe._gced
			_adfc := (_ddabe._bdga*1000.0 - _gbgcg) / _ddabe._gced
			_eaee = append(_eaee, _fg.MakeFloat(-_adfc))
		}
		_cfcg := _ddabe._ceda.Encoder()
		var _cacdg []byte
		for _, _adggg := range _aege {
			if _adggg == '\u000A' {
				continue
			}
			if _adggg == ' ' {
				if len(_cacdg) > 0 {
					_eaee = append(_eaee, _fg.MakeStringFromBytes(_cacdg))
					_cacdg = nil
				}
				_eaee = append(_eaee, _fg.MakeFloat(-_aeca))
			} else {
				if _, _edfd := _cfcg.RuneToCharcode(_adggg); !_edfd {
					_dggee = UnsupportedRuneError{Message: _eg.Sprintf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0075\u006e\u0065 \u0069\u006e\u0020\u0074\u0065\u0078\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0023\u0078\u0020\u0028\u0025\u0063\u0029", _adggg, _adggg), Rune: _adggg}
					_acdg._fdcc = append(_acdg._fdcc, _dggee)
					_df.Log.Debug(_dggee.Error())
					if _acdg._eag <= 0 {
						continue
					}
					_adggg = _acdg._eag
				}
				_cacdg = append(_cacdg, _cfcg.Encode(string(_adggg))...)
			}
		}
		if len(_cacdg) > 0 {
			_eaee = append(_eaee, _fg.MakeStringFromBytes(_cacdg))
		}
		_faad.Add_TJ(_eaee...)
	}
	_faad.Add_ET()
	_faad.Add_Q()
	_acfg := _faad.Operations()
	_acfg.WrapIfNeeded()
	_dddf.addContents(_acfg)
	if _ddabe._beeba.IsRelative() {
		_dbab := _ddabe.Height()
		_acdg.Y += _dbab
		_acdg.Height -= _dbab
		if _acdg.Inline {
			_acdg.X += _ddabe.Width() + _ddabe._eacf.Right
		}
	}
	return _acdg, nil
}
func (_dgcd *templateProcessor) parseChart(_abdc *templateNode) (interface{}, error) {
	var _agff string
	for _, _afecf := range _abdc._ddgde.Attr {
		_eefb := _afecf.Value
		switch _efaag := _afecf.Name.Local; _efaag {
		case "\u0073\u0072\u0063":
			_agff = _eefb
		}
	}
	if _agff == "" {
		_df.Log.Error("\u0043\u0068\u0061\u0072\u0074\u0020\u0060\u0073\u0072\u0063\u0060\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062e\u0020\u0065m\u0070\u0074\u0079\u002e")
		return nil, _adbab
	}
	_ecfa, _dabe := _dgcd._daae.ChartMap[_agff]
	if !_dabe {
		_df.Log.Error("\u0043\u006ful\u0064\u0020\u006eo\u0074\u0020\u0066\u0069nd \u0063ha\u0072\u0074\u0020\u0072\u0065\u0073\u006fur\u0063\u0065\u003a\u0020\u0060\u0025\u0073`\u002e", _agff)
		return nil, _adbab
	}
	_adfd := NewChart(_ecfa)
	for _, _aafdg := range _abdc._ddgde.Attr {
		_babd := _aafdg.Value
		switch _ddggf := _aafdg.Name.Local; _ddggf {
		case "\u0078":
			_adfd.SetPos(_dgcd.parseFloatAttr(_ddggf, _babd), _adfd._dbc)
		case "\u0079":
			_adfd.SetPos(_adfd._eca, _dgcd.parseFloatAttr(_ddggf, _babd))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_fcecb := _dgcd.parseMarginAttr(_ddggf, _babd)
			_adfd.SetMargins(_fcecb.Left, _fcecb.Right, _fcecb.Top, _fcecb.Bottom)
		case "\u0077\u0069\u0064t\u0068":
			_adfd._fbaa.SetWidth(int(_dgcd.parseFloatAttr(_ddggf, _babd)))
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_adfd._fbaa.SetHeight(int(_dgcd.parseFloatAttr(_ddggf, _babd)))
		case "\u0073\u0072\u0063":
		default:
			_df.Log.Debug("\u0055n\u0073\u0075p\u0070\u006f\u0072\u0074e\u0064\u0020\u0063h\u0061\u0072\u0074\u0020\u0061\u0074\u0074\u0072\u0069bu\u0074\u0065\u003a \u0060\u0025s\u0060\u002e\u0020\u0053\u006b\u0069p\u0070\u0069n\u0067\u002e", _ddggf)
		}
	}
	return _adfd, nil
}

// SetFillColor sets the fill color.
func (_eggc *Polygon) SetFillColor(color Color) { _eggc._ccbeg.FillColor = _bcfe(color) }

// SetMargins sets the Paragraph's margins.
func (_ggdf *StyledParagraph) SetMargins(left, right, top, bottom float64) {
	_ggdf._afdba.Left = left
	_ggdf._afdba.Right = right
	_ggdf._afdba.Top = top
	_ggdf._afdba.Bottom = bottom
}

// SetDueDate sets the due date of the invoice.
func (_bdff *Invoice) SetDueDate(dueDate string) (*InvoiceCell, *InvoiceCell) {
	_bdff._debe[1].Value = dueDate
	return _bdff._debe[0], _bdff._debe[1]
}
func (_bedeg *templateProcessor) parseFontAttr(_bggc, _cecgb string) *_f.PdfFont {
	_df.Log.Debug("P\u0061\u0072\u0073\u0069\u006e\u0067 \u0066\u006f\u006e\u0074\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _bggc, _cecgb)
	_ccgea := _bedeg.creator._ffgd
	if _cecgb == "" {
		return _ccgea
	}
	_becga := _ed.Split(_cecgb, "\u002c")
	for _, _edgdb := range _becga {
		_edgdb = _ed.TrimSpace(_edgdb)
		if _edgdb == "" {
			continue
		}
		_bccg, _aeeac := _bedeg._daae.FontMap[_cecgb]
		if _aeeac {
			return _bccg
		}
		_bdgf, _aeeac := map[string]_f.StdFontName{"\u0063o\u0075\u0072\u0069\u0065\u0072": _f.CourierName, "\u0063\u006f\u0075r\u0069\u0065\u0072\u002d\u0062\u006f\u006c\u0064": _f.CourierBoldName, "\u0063o\u0075r\u0069\u0065\u0072\u002d\u006f\u0062\u006c\u0069\u0071\u0075\u0065": _f.CourierObliqueName, "c\u006fu\u0072\u0069\u0065\u0072\u002d\u0062\u006f\u006cd\u002d\u006f\u0062\u006ciq\u0075\u0065": _f.CourierBoldObliqueName, "\u0068e\u006c\u0076\u0065\u0074\u0069\u0063a": _f.HelveticaName, "\u0068\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061-\u0062\u006f\u006c\u0064": _f.HelveticaBoldName, "\u0068\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061\u002d\u006f\u0062l\u0069\u0071\u0075\u0065": _f.HelveticaObliqueName, "\u0068\u0065\u006c\u0076et\u0069\u0063\u0061\u002d\u0062\u006f\u006c\u0064\u002d\u006f\u0062\u006c\u0069\u0071u\u0065": _f.HelveticaBoldObliqueName, "\u0073\u0079\u006d\u0062\u006f\u006c": _f.SymbolName, "\u007a\u0061\u0070\u0066\u002d\u0064\u0069\u006e\u0067\u0062\u0061\u0074\u0073": _f.ZapfDingbatsName, "\u0074\u0069\u006de\u0073": _f.TimesRomanName, "\u0074\u0069\u006d\u0065\u0073\u002d\u0062\u006f\u006c\u0064": _f.TimesBoldName, "\u0074\u0069\u006de\u0073\u002d\u0069\u0074\u0061\u006c\u0069\u0063": _f.TimesItalicName, "\u0074\u0069\u006d\u0065\u0073\u002d\u0062\u006f\u006c\u0064\u002d\u0069t\u0061\u006c\u0069\u0063": _f.TimesBoldItalicName}[_cecgb]
		if _aeeac {
			if _ebee, _fcfd := _f.NewStandard14Font(_bdgf); _fcfd == nil {
				return _ebee
			}
		}
		if _adcb := _bedeg.parseAttrPropList(_edgdb); len(_adcb) > 0 {
			if _abee, _afdbb := _adcb["\u0070\u0061\u0074\u0068"]; _afdbb {
				_ceecde := _f.NewPdfFontFromTTFFile
				if _bega, _babde := _adcb["\u0074\u0079\u0070\u0065"]; _babde && _bega == "\u0063o\u006d\u0070\u006f\u0073\u0069\u0074e" {
					_ceecde = _f.NewCompositePdfFontFromTTFFile
				}
				if _bdab, _cebbd := _ceecde(_abee); _cebbd != nil {
					_df.Log.Debug("\u0043\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0060\u0025\u0073\u0060\u003a %\u0076\u002e", _abee, _cebbd)
				} else {
					return _bdab
				}
			}
		}
	}
	return _ccgea
}

// SetPositioning sets Ellipse's position attribute.
func (_dfgg *Ellipse) SetPositioning(position Positioning) { _dfgg._cfagc = position }

// SetLineWidth sets the line width.
func (_gecbe *Line) SetLineWidth(width float64) { _gecbe._ffgbg = width }

// SetLineNumberStyle sets the style for the numbers part of all new lines
// of the table of contents.
func (_baag *TOC) SetLineNumberStyle(style TextStyle) { _baag._ccdbe = style }

// SetBorderWidth sets the border width.
func (_dgff *Polygon) SetBorderWidth(borderWidth float64) { _dgff._ccbeg.BorderWidth = borderWidth }
func _ggaa(_bgfgf TextStyle) *StyledParagraph {
	return &StyledParagraph{_eedad: []*TextChunk{}, _gdbeb: _bgfgf, _bdbf: _fcabb(_bgfgf.Font), _bged: 1.0, _bcege: TextAlignmentLeft, _ecbf: true, _ffbb: true, _bggd: false, _daab: 0, _fgaea: 1, _fbgba: 1, _geagd: PositionRelative}
}
func _agddb(_ffgbf, _beec, _bbeb, _ecdf float64) *Line {
	return &Line{_ggfd: _ffgbf, _fadc: _beec, _dccb: _bbeb, _gdec: _ecdf, _ggdcb: ColorBlack, _cgcbf: 1.0, _ffgbg: 1.0, _gbde: []int64{1, 1}, _edcee: PositionAbsolute}
}
func (_egg *FilledCurve) draw(_ebec string) ([]byte, *_f.PdfRectangle, error) {
	_cggd := _cec.NewCubicBezierPath()
	for _, _adaff := range _egg._bgf {
		_cggd = _cggd.AppendCurve(_adaff)
	}
	creator := _eee.NewContentCreator()
	creator.Add_q()
	if _egg.FillEnabled && _egg._bffg != nil {
		creator.SetNonStrokingColor(_bcfe(_egg._bffg))
	}
	if _egg.BorderEnabled {
		if _egg._aac != nil {
			creator.SetStrokingColor(_bcfe(_egg._aac))
		}
		creator.Add_w(_egg.BorderWidth)
	}
	if len(_ebec) > 1 {
		creator.Add_gs(_fg.PdfObjectName(_ebec))
	}
	_cec.DrawBezierPathWithCreator(_cggd, creator)
	creator.Add_h()
	if _egg.FillEnabled && _egg.BorderEnabled {
		creator.Add_B()
	} else if _egg.FillEnabled {
		creator.Add_f()
	} else if _egg.BorderEnabled {
		creator.Add_S()
	}
	creator.Add_Q()
	_eeaa := _cggd.GetBoundingBox()
	if _egg.BorderEnabled {
		_eeaa.Height += _egg.BorderWidth
		_eeaa.Width += _egg.BorderWidth
		_eeaa.X -= _egg.BorderWidth / 2
		_eeaa.Y -= _egg.BorderWidth / 2
	}
	_ffag := &_f.PdfRectangle{}
	_ffag.Llx = _eeaa.X
	_ffag.Lly = _eeaa.Y
	_ffag.Urx = _eeaa.X + _eeaa.Width
	_ffag.Ury = _eeaa.Y + _eeaa.Height
	return creator.Bytes(), _ffag, nil
}

// SetOptimizer sets the optimizer to optimize PDF before writing.
func (_cdae *Creator) SetOptimizer(optimizer _f.Optimizer) { _cdae._agg = optimizer }

// Date returns the invoice date description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_fdag *Invoice) Date() (*InvoiceCell, *InvoiceCell) { return _fdag._fgcge[0], _fdag._fgcge[1] }

// Height returns the height of the Paragraph. The height is calculated based on the input text and
// how it is wrapped within the container. Does not include Margins.
func (_cgcc *Paragraph) Height() float64 {
	_cgcc.wrapText()
	return float64(len(_cgcc._cead)) * _cgcc._abbg * _cgcc._gced
}

// NewTOCLine creates a new table of contents line with the default style.
func (_cdce *Creator) NewTOCLine(number, title, page string, level uint) *TOCLine {
	return _eafcg(number, title, page, level, _cdce.NewTextStyle())
}

// Width is not used. Not used as a Table element is designed to fill into
// available width depending on the context. Returns 0.
func (_dffb *Table) Width() float64 { return 0 }

const (
	PositionRelative Positioning = iota
	PositionAbsolute
)

// GeneratePageBlocks draws the curve onto page blocks.
func (_cdaed *Curve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gdga := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_deg := _eee.NewContentCreator()
	_deg.Add_q().Add_w(_cdaed._bbf).SetStrokingColor(_bcfe(_cdaed._cddgg)).Add_m(_cdaed._ggce, ctx.PageHeight-_cdaed._bdcce).Add_v(_cdaed._dafb, ctx.PageHeight-_cdaed._feec, _cdaed._fgaeb, ctx.PageHeight-_cdaed._aaeb).Add_S().Add_Q()
	_agb := _gdga.addContentsByString(_deg.String())
	if _agb != nil {
		return nil, ctx, _agb
	}
	return []*Block{_gdga}, ctx, nil
}

// SetPadding sets the padding of the component. The padding represents
// inner margins which are applied around the contents of the division.
// The background of the component is not affected by its padding.
func (_bffb *Division) SetPadding(left, right, top, bottom float64) {
	_bffb._gfeb.Left = left
	_bffb._gfeb.Right = right
	_bffb._gfeb.Top = top
	_bffb._gfeb.Bottom = bottom
}

// Margins returns the margins of the list: left, right, top, bottom.
func (_bfecf *List) Margins() (float64, float64, float64, float64) {
	return _bfecf._bdccd.Left, _bfecf._bdccd.Right, _bfecf._bdccd.Top, _bfecf._bdccd.Bottom
}
func _edcgg(_beag *templateProcessor, _efce *templateNode) (interface{}, error) {
	return _beag.parseTable(_efce)
}
func (_aabb *Paragraph) getTextLineWidth(_cggf string) float64 {
	var _agdae float64
	for _, _dcfgc := range _cggf {
		if _dcfgc == '\u000A' {
			continue
		}
		_cdbf, _fedgd := _aabb._ceda.GetRuneMetrics(_dcfgc)
		if !_fedgd {
			_df.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052u\u006e\u0065\u0020\u0063\u0068a\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0028\u0072\u0075\u006e\u0065\u0020\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0029", _dcfgc, _dcfgc)
			return -1
		}
		_agdae += _aabb._gced * _cdbf.Wx
	}
	return _agdae
}

// Fit fits the chunk into the specified bounding box, cropping off the
// remainder in a new chunk, if it exceeds the specified dimensions.
// NOTE: The method assumes a line height of 1.0. In order to account for other
// line height values, the passed in height must be divided by the line height:
// height = height / lineHeight
func (_ccgb *TextChunk) Fit(width, height float64) (*TextChunk, error) {
	_adac, _dbgf := _ccgb.Wrap(width)
	if _dbgf != nil {
		return nil, _dbgf
	}
	_dbedf := int(height / _ccgb.Style.FontSize)
	if _dbedf >= len(_adac) {
		return nil, nil
	}
	_fcce := "\u000a"
	_ccgb.Text = _ed.Replace(_ed.Join(_adac[:_dbedf], "\u0020"), _fcce+"\u0020", _fcce, -1)
	_debbd := _ed.Replace(_ed.Join(_adac[_dbedf:], "\u0020"), _fcce+"\u0020", _fcce, -1)
	return NewTextChunk(_debbd, _ccgb.Style), nil
}

// SkipCells skips over a specified number of cells in the table.
func (_gaaed *Table) SkipCells(num int) {
	if num < 0 {
		_df.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	_gaaed._edag += num
}

// SetAddressStyle sets the style properties used to render the content of
// the invoice address sections.
func (_daagg *Invoice) SetAddressStyle(style TextStyle) { _daagg._ddf = style }

// Length calculates and returns the length of the line.
func (_cdab *Line) Length() float64 {
	return _a.Sqrt(_a.Pow(_cdab._dccb-_cdab._ggfd, 2.0) + _a.Pow(_cdab._gdec-_cdab._fadc, 2.0))
}

// Subtotal returns the invoice subtotal description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_beddf *Invoice) Subtotal() (*InvoiceCell, *InvoiceCell) {
	return _beddf._faeg[0], _beddf._faeg[1]
}

// DrawTemplate renders the template provided through the specified reader,
// using the specified `data` and `options`.
// Creator templates are first executed as text/template *Template instances,
// so the specified `data` is inserted within the template.
// The second phase of processing is actually parsing the template, translating
// it into creator components and rendering them using the provided options.
// Both the `data` and `options` parameters can be nil.
func (_abb *Creator) DrawTemplate(r _ef.Reader, data interface{}, options *TemplateOptions) error {
	return _edfbf(_abb, r, data, options, _abb)
}

// SetMargins sets the Table's left, right, top, bottom margins.
func (_dbcb *Table) SetMargins(left, right, top, bottom float64) {
	_dbcb._ddcc.Left = left
	_dbcb._ddcc.Right = right
	_dbcb._ddcc.Top = top
	_dbcb._ddcc.Bottom = bottom
}

// ColorCMYKFromArithmetic creates a Color from arithmetic color values (0-1).
// Example:
//   green := ColorCMYKFromArithmetic(1.0, 0.0, 1.0, 0.0)
func ColorCMYKFromArithmetic(c, m, y, k float64) Color {
	return cmykColor{_fefa: _a.Max(_a.Min(c, 1.0), 0.0), _dcefe: _a.Max(_a.Min(m, 1.0), 0.0), _eefc: _a.Max(_a.Min(y, 1.0), 0.0), _bef: _a.Max(_a.Min(k, 1.0), 0.0)}
}
func (_dbbf *pageTransformations) transformPage(_gddg *_f.PdfPage) error {
	if _fbee := _dbbf.applyFlip(_gddg); _fbee != nil {
		return _fbee
	}
	return nil
}
func _gbef() *FilledCurve {
	_deaf := FilledCurve{}
	_deaf._bgf = []_cec.CubicBezierCurve{}
	return &_deaf
}

// SetLineWidth sets the line width.
func (_cbaa *Polyline) SetLineWidth(lineWidth float64) { _cbaa._bcg.LineWidth = lineWidth }

// GeneratePageBlocks draws the chart onto a block.
func (_abe *Chart) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gbc := ctx
	_bgag := _abe._ebe.IsRelative()
	var _agde []*Block
	if _bgag {
		_eab := 1.0
		_afd := _abe._daea.Top
		if float64(_abe._fbaa.Height()) > ctx.Height-_abe._daea.Top {
			_agde = []*Block{NewBlock(ctx.PageWidth, ctx.PageHeight-ctx.Y)}
			var _ccg error
			if _, ctx, _ccg = _acdfe().GeneratePageBlocks(ctx); _ccg != nil {
				return nil, ctx, _ccg
			}
			_afd = 0
		}
		ctx.X += _abe._daea.Left + _eab
		ctx.Y += _afd
		ctx.Width -= _abe._daea.Left + _abe._daea.Right + 2*_eab
		ctx.Height -= _afd
		_abe._fbaa.SetWidth(int(ctx.Width))
	} else {
		ctx.X = _abe._eca
		ctx.Y = _abe._dbc
	}
	_acde := _eee.NewContentCreator()
	_acde.Translate(0, ctx.PageHeight)
	_acde.Scale(1, -1)
	_acde.Translate(ctx.X, ctx.Y)
	_cged := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_abe._fbaa.Render(_fb.NewRenderer(_acde, _cged._eb), nil)
	if _caaa := _cged.addContentsByString(_acde.String()); _caaa != nil {
		return nil, ctx, _caaa
	}
	if _bgag {
		_egcf := _abe.Height() + _abe._daea.Bottom
		ctx.Y += _egcf
		ctx.Height -= _egcf
	} else {
		ctx = _gbc
	}
	_agde = append(_agde, _cged)
	return _agde, ctx, nil
}

// Columns returns all the columns in the invoice line items table.
func (_gddcc *Invoice) Columns() []*InvoiceCell { return _gddcc._afeb }

// Link returns link information for this line.
func (_bgecd *TOCLine) Link() (_fagg int64, _caagd, _ecabeb float64) {
	return _bgecd._cbfg, _bgecd._cegge, _bgecd._gaace
}

// Width returns Rectangle's document width.
func (_fcbfb *Rectangle) Width() float64 { return _fcbfb._ecfc }

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_edde *TOCLine) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_deadf := ctx
	_dfbb, ctx, _gdegd := _edde._face.GeneratePageBlocks(ctx)
	if _gdegd != nil {
		return _dfbb, ctx, _gdegd
	}
	if _edde._dcafec.IsRelative() {
		ctx.X = _deadf.X
	}
	if _edde._dcafec.IsAbsolute() {
		return _dfbb, _deadf, nil
	}
	return _dfbb, ctx, nil
}

// ColorRGBFromHex converts color hex code to rgb color for using with creator.
// NOTE: If there is a problem interpreting the string, then will use black color and log a debug message.
// Example hex code: #ffffff -> (1,1,1) white.
func ColorRGBFromHex(hexStr string) Color {
	_bacc := rgbColor{}
	if (len(hexStr) != 4 && len(hexStr) != 7) || hexStr[0] != '#' {
		_df.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
		return _bacc
	}
	var _ddge, _cgg, _dee int
	if len(hexStr) == 4 {
		var _adeb, _fed, _gagf int
		_fbgc, _cddd := _eg.Sscanf(hexStr, "\u0023\u0025\u0031\u0078\u0025\u0031\u0078\u0025\u0031\u0078", &_adeb, &_fed, &_gagf)
		if _cddd != nil {
			_df.Log.Debug("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020h\u0065\u0078\u0020\u0063\u006f\u0064\u0065:\u0020\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", hexStr, _cddd)
			return _bacc
		}
		if _fbgc != 3 {
			_df.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
			return _bacc
		}
		_ddge = _adeb*16 + _adeb
		_cgg = _fed*16 + _fed
		_dee = _gagf*16 + _gagf
	} else {
		_faef, _ccc := _eg.Sscanf(hexStr, "\u0023\u0025\u0032\u0078\u0025\u0032\u0078\u0025\u0032\u0078", &_ddge, &_cgg, &_dee)
		if _ccc != nil {
			_df.Log.Debug("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073", hexStr)
			return _bacc
		}
		if _faef != 3 {
			_df.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0073,\u0020\u006e\u0020\u0021\u003d\u0020\u0033 \u0028\u0025\u0064\u0029", hexStr, _faef)
			return _bacc
		}
	}
	_fgaag := float64(_ddge) / 255.0
	_fefb := float64(_cgg) / 255.0
	_fcb := float64(_dee) / 255.0
	_bacc._fabb = _fgaag
	_bacc._abab = _fefb
	_bacc._gggf = _fcb
	return _bacc
}

// GetMargins returns the Image's margins: left, right, top, bottom.
func (_effg *Image) GetMargins() (float64, float64, float64, float64) {
	return _effg._dbfe.Left, _effg._dbfe.Right, _effg._dbfe.Top, _effg._dbfe.Bottom
}
func (_dega *templateProcessor) addNodeText(_fbbfc *templateNode, _ccee string) error {
	_beefb := _fbbfc._cedba
	if _beefb == nil {
		return nil
	}
	switch _dceb := _beefb.(type) {
	case *TextChunk:
		_dceb.Text = _ccee
	case *Paragraph:
		switch _fbbfc._ddgde.Name.Local {
		case "\u0063h\u0061p\u0074\u0065\u0072\u002d\u0068\u0065\u0061\u0064\u0069\u006e\u0067":
			if _fbbfc._dfdd != nil {
				if _bbbg, _gfgc := _fbbfc._dfdd._cedba.(*Chapter); _gfgc {
					_bbbg._bbe = _ccee
					_dceb.SetText(_bbbg.headingText())
				}
			}
		default:
			_dceb.SetText(_ccee)
		}
	}
	return nil
}
func (_gcfaf *StyledParagraph) appendChunk(_dade *TextChunk) *TextChunk {
	_gcfaf._eedad = append(_gcfaf._eedad, _dade)
	_gcfaf.wrapText()
	return _dade
}

// NewBlockFromPage creates a Block from a PDF Page.  Useful for loading template pages as blocks
// from a PDF document and additional content with the creator.
func NewBlockFromPage(page *_f.PdfPage) (*Block, error) {
	_aag := &Block{}
	_afb, _ceg := page.GetAllContentStreams()
	if _ceg != nil {
		return nil, _ceg
	}
	_fcc := _eee.NewContentStreamParser(_afb)
	_dfb, _ceg := _fcc.Parse()
	if _ceg != nil {
		return nil, _ceg
	}
	_dfb.WrapIfNeeded()
	_aag._ga = _dfb
	if page.Resources != nil {
		_aag._eb = page.Resources
	} else {
		_aag._eb = _f.NewPdfPageResources()
	}
	_fgc, _ceg := page.GetMediaBox()
	if _ceg != nil {
		return nil, _ceg
	}
	if _fgc.Llx != 0 || _fgc.Lly != 0 {
		_aag.translate(-_fgc.Llx, _fgc.Lly)
	}
	_aag._bf = _fgc.Urx - _fgc.Llx
	_aag._eda = _fgc.Ury - _fgc.Lly
	if page.Rotate != nil {
		_aag._ag = -float64(*page.Rotate)
	}
	return _aag, nil
}

// Width returns the Block's width.
func (_ge *Block) Width() float64 { return _ge._bf }

// Wrap wraps the text of the chunk into lines based on its style and the
// specified width.
func (_fdcf *TextChunk) Wrap(width float64) ([]string, error) {
	if int(width) <= 0 {
		return []string{_fdcf.Text}, nil
	}
	var _badb []string
	var _cbcbf []rune
	var _aabd float64
	var _abfb []float64
	_bgagd := _fdcf.Style
	_bgacg := _gaafg(_fdcf.Text)
	for _, _egga := range _fdcf.Text {
		if _egga == '\u000A' {
			_becb := _daed(string(_cbcbf), _bgacg)
			_badb = append(_badb, _ed.TrimRightFunc(_becb, _ee.IsSpace)+string(_egga))
			_cbcbf = nil
			_aabd = 0
			_abfb = nil
			continue
		}
		_defag := _egga == ' '
		_dedbc, _beccd := _bgagd.Font.GetRuneMetrics(_egga)
		if !_beccd {
			_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006det\u0072i\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064!\u0020\u0072\u0075\u006e\u0065\u003d\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0020\u0066o\u006e\u0074\u003d\u0025\u0073\u0020\u0025\u0023\u0071", _egga, _egga, _bgagd.Font.BaseFont(), _bgagd.Font.Subtype())
			_df.Log.Trace("\u0046o\u006e\u0074\u003a\u0020\u0025\u0023v", _bgagd.Font)
			_df.Log.Trace("\u0045\u006e\u0063o\u0064\u0065\u0072\u003a\u0020\u0025\u0023\u0076", _bgagd.Font.Encoder())
			return nil, _d.New("\u0067\u006c\u0079\u0070\u0068\u0020\u0063\u0068\u0061\u0072\u0020m\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
		}
		_gefge := _bgagd.FontSize * _dedbc.Wx
		_bcaac := _gefge
		if !_defag {
			_bcaac = _gefge + _bgagd.CharSpacing*1000.0
		}
		if _aabd+_gefge > width*1000.0 {
			_egfe := -1
			if !_defag {
				for _ebfec := len(_cbcbf) - 1; _ebfec >= 0; _ebfec-- {
					if _cbcbf[_ebfec] == ' ' {
						_egfe = _ebfec
						break
					}
				}
			}
			_febg := string(_cbcbf)
			if _egfe > 0 {
				_febg = string(_cbcbf[0 : _egfe+1])
				_cbcbf = append(_cbcbf[_egfe+1:], _egga)
				_abfb = append(_abfb[_egfe+1:], _bcaac)
				_aabd = 0
				for _, _gggcg := range _abfb {
					_aabd += _gggcg
				}
			} else {
				if _defag {
					_cbcbf = []rune{}
					_abfb = []float64{}
					_aabd = 0
				} else {
					_cbcbf = []rune{_egga}
					_abfb = []float64{_bcaac}
					_aabd = _bcaac
				}
			}
			_febg = _daed(_febg, _bgacg)
			_badb = append(_badb, _ed.TrimRightFunc(_febg, _ee.IsSpace))
		} else {
			_cbcbf = append(_cbcbf, _egga)
			_aabd += _bcaac
			_abfb = append(_abfb, _bcaac)
		}
	}
	if len(_cbcbf) > 0 {
		_fcacg := string(_cbcbf)
		_fcacg = _daed(_fcacg, _bgacg)
		_badb = append(_badb, _fcacg)
	}
	return _badb, nil
}
func (_fcfbc *Invoice) generateTotalBlocks(_fcbf DrawContext) ([]*Block, DrawContext, error) {
	_efdfa := _ecef(4)
	_efdfa.SetMargins(0, 0, 10, 10)
	_affd := [][2]*InvoiceCell{_fcfbc._faeg}
	_affd = append(_affd, _fcfbc._gce...)
	_affd = append(_affd, _fcfbc._defg)
	for _, _daagf := range _affd {
		_fadgf, _cecg := _daagf[0], _daagf[1]
		if _cecg.Value == "" {
			continue
		}
		_efdfa.SkipCells(2)
		_bgcg := _efdfa.NewCell()
		_bgcg.SetBackgroundColor(_fadgf.BackgroundColor)
		_bgcg.SetHorizontalAlignment(_cecg.Alignment)
		_fcfbc.setCellBorder(_bgcg, _fadgf)
		_dcga := _ggaa(_fadgf.TextStyle)
		_dcga.SetMargins(0, 0, 2, 1)
		_dcga.Append(_fadgf.Value)
		_bgcg.SetContent(_dcga)
		_bgcg = _efdfa.NewCell()
		_bgcg.SetBackgroundColor(_cecg.BackgroundColor)
		_bgcg.SetHorizontalAlignment(_cecg.Alignment)
		_fcfbc.setCellBorder(_bgcg, _fadgf)
		_dcga = _ggaa(_cecg.TextStyle)
		_dcga.SetMargins(0, 0, 2, 1)
		_dcga.Append(_cecg.Value)
		_bgcg.SetContent(_dcga)
	}
	return _efdfa.GeneratePageBlocks(_fcbf)
}

// SetTextAlignment sets the horizontal alignment of the text within the space provided.
func (_cdag *StyledParagraph) SetTextAlignment(align TextAlignment) { _cdag._bcege = align }

// Height returns the height of the division, assuming all components are
// stacked on top of each other.
func (_fcbg *Division) Height() float64 {
	var _bab float64
	for _, _fgdf := range _fcbg._aefd {
		switch _cfag := _fgdf.(type) {
		case marginDrawable:
			_, _, _fdf, _eaca := _cfag.GetMargins()
			_bab += _cfag.Height() + _fdf + _eaca
		default:
			_bab += _cfag.Height()
		}
	}
	return _bab
}

// SetForms adds an Acroform to a PDF file.  Sets the specified form for writing.
func (_gaba *Creator) SetForms(form *_f.PdfAcroForm) error { _gaba._eafg = form; return nil }

// MoveTo moves the drawing context to absolute coordinates (x, y).
func (_befb *Creator) MoveTo(x, y float64) { _befb._bfec.X = x; _befb._bfec.Y = y }
func (_afba *Block) addContents(_ace *_eee.ContentStreamOperations) {
	_afba._ga.WrapIfNeeded()
	_ace.WrapIfNeeded()
	*_afba._ga = append(*_afba._ga, *_ace...)
}
func _baec(_faea *_f.Image) (*Image, error) {
	_cccf := float64(_faea.Width)
	_dabb := float64(_faea.Height)
	return &Image{_aaca: _faea, _cbbgg: _cccf, _gbf: _dabb, _bcfd: _cccf, _gddc: _dabb, _afgea: 0, _edgb: 1.0, _ccec: PositionRelative}, nil
}

// SetOpacity sets opacity for Image.
func (_efae *Image) SetOpacity(opacity float64) { _efae._edgb = opacity }

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

// SetMargins sets the margins TOC line.
func (_bggfc *TOCLine) SetMargins(left, right, top, bottom float64) {
	_bggfc._ecdbc = left
	_gccde := &_bggfc._face._afdba
	_gccde.Left = _bggfc._ecdbc + float64(_bggfc._dgcda-1)*_bggfc._aaaa
	_gccde.Right = right
	_gccde.Top = top
	_gccde.Bottom = bottom
}

// SetSideBorderWidth sets the cell's side border width.
func (_cbec *TableCell) SetSideBorderWidth(side CellBorderSide, width float64) {
	switch side {
	case CellBorderSideAll:
		_cbec._cfdc = width
		_cbec._efgc = width
		_cbec._bfbdc = width
		_cbec._bfgd = width
	case CellBorderSideTop:
		_cbec._cfdc = width
	case CellBorderSideBottom:
		_cbec._efgc = width
	case CellBorderSideLeft:
		_cbec._bfbdc = width
	case CellBorderSideRight:
		_cbec._bfgd = width
	}
}

// SetDate sets the date of the invoice.
func (_bdbb *Invoice) SetDate(date string) (*InvoiceCell, *InvoiceCell) {
	_bdbb._fgcge[1].Value = date
	return _bdbb._fgcge[0], _bdbb._fgcge[1]
}

// NewCurvePolygon creates a new curve polygon.
func (_egcab *Creator) NewCurvePolygon(rings [][]_cec.CubicBezierCurve) *CurvePolygon {
	return _dbff(rings)
}

var (
	PageSizeA3     = PageSize{297 * PPMM, 420 * PPMM}
	PageSizeA4     = PageSize{210 * PPMM, 297 * PPMM}
	PageSizeA5     = PageSize{148 * PPMM, 210 * PPMM}
	PageSizeLetter = PageSize{8.5 * PPI, 11 * PPI}
	PageSizeLegal  = PageSize{8.5 * PPI, 14 * PPI}
)

// SetSideBorderStyle sets the cell's side border style.
func (_gaaeda *TableCell) SetSideBorderStyle(side CellBorderSide, style CellBorderStyle) {
	switch side {
	case CellBorderSideAll:
		_gaaeda._dbfgg = style
		_gaaeda._dgbfe = style
		_gaaeda._dbd = style
		_gaaeda._cebbe = style
	case CellBorderSideTop:
		_gaaeda._dbfgg = style
	case CellBorderSideBottom:
		_gaaeda._dgbfe = style
	case CellBorderSideLeft:
		_gaaeda._dbd = style
	case CellBorderSideRight:
		_gaaeda._cebbe = style
	}
}

// GeneratePageBlocks draws the composite Bezier curve on a new block
// representing the page. Implements the Drawable interface.
func (_abdb *PolyBezierCurve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_ebgd := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_dffe, _cdac := _ebgd.setOpacity(_abdb._bbeag, _abdb._bcbe)
	if _cdac != nil {
		return nil, ctx, _cdac
	}
	_ccgd := _abdb._cegf
	_ccgd.FillEnabled = _ccgd.FillColor != nil
	var (
		_ggcc   = ctx.PageHeight
		_bgadcf = _ccgd.Curves
		_ffce   = make([]_cec.CubicBezierCurve, 0, len(_ccgd.Curves))
	)
	for _gcfa := range _ccgd.Curves {
		_agdcd := _bgadcf[_gcfa]
		_agdcd.P0.Y = _ggcc - _agdcd.P0.Y
		_agdcd.P1.Y = _ggcc - _agdcd.P1.Y
		_agdcd.P2.Y = _ggcc - _agdcd.P2.Y
		_agdcd.P3.Y = _ggcc - _agdcd.P3.Y
		_ffce = append(_ffce, _agdcd)
	}
	_ccgd.Curves = _ffce
	defer func() { _ccgd.Curves = _bgadcf }()
	_dbbd, _, _cdac := _ccgd.Draw(_dffe)
	if _cdac != nil {
		return nil, ctx, _cdac
	}
	if _cdac = _ebgd.addContentsByString(string(_dbbd)); _cdac != nil {
		return nil, ctx, _cdac
	}
	return []*Block{_ebgd}, ctx, nil
}

// SetBorderRadius sets the radius of the rectangle corners.
func (_dgdeb *Rectangle) SetBorderRadius(topLeft, topRight, bottomLeft, bottomRight float64) {
	_dgdeb._ecdc = topLeft
	_dgdeb._aaegc = topRight
	_dgdeb._bbde = bottomLeft
	_dgdeb._dada = bottomRight
}

type cmykColor struct{ _fefa, _dcefe, _eefc, _bef float64 }

// NewPolyBezierCurve creates a new composite Bezier (polybezier) curve.
func (_ddef *Creator) NewPolyBezierCurve(curves []_cec.CubicBezierCurve) *PolyBezierCurve {
	return _fgcgae(curves)
}
func (_de *Block) addContentsByString(_ebg string) error {
	_fccb := _eee.NewContentStreamParser(_ebg)
	_cg, _eaa := _fccb.Parse()
	if _eaa != nil {
		return _eaa
	}
	_de._ga.WrapIfNeeded()
	_cg.WrapIfNeeded()
	*_de._ga = append(*_de._ga, *_cg...)
	return nil
}

// SetColumnWidths sets the fractional column widths.
// Each width should be in the range 0-1 and is a fraction of the table width.
// The number of width inputs must match number of columns, otherwise an error is returned.
func (_gbfb *Table) SetColumnWidths(widths ...float64) error {
	if len(widths) != _gbfb._cfee {
		_df.Log.Debug("M\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0069\u006e\u0067\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020\u0077\u0069\u0064\u0074\u0068\u0073\u0020\u0061nd\u0020\u0063\u006fl\u0075m\u006e\u0073")
		return _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_gbfb._agdgd = widths
	return nil
}

// Height returns Ellipse's document height.
func (_bebg *Ellipse) Height() float64 { return _bebg._gdeg }

// SetMargins sets the margins of the paragraph.
func (_acbc *List) SetMargins(left, right, top, bottom float64) {
	_acbc._bdccd.Left = left
	_acbc._bdccd.Right = right
	_acbc._bdccd.Top = top
	_acbc._bdccd.Bottom = bottom
}
func _daed(_baga string, _agfdf bool) string {
	_ceddb := _baga
	if _ceddb == "" {
		return ""
	}
	_dgaf := _af.Paragraph{}
	_, _gfeeg := _dgaf.SetString(_baga)
	if _gfeeg != nil {
		return _ceddb
	}
	_aafae, _gfeeg := _dgaf.Order()
	if _gfeeg != nil {
		return _ceddb
	}
	_efcb := _aafae.NumRuns()
	_cegad := make([]string, _efcb)
	for _agaaa := 0; _agaaa < _aafae.NumRuns(); _agaaa++ {
		_gcbaf := _aafae.Run(_agaaa)
		_fdbd := _gcbaf.String()
		if _gcbaf.Direction() == _af.RightToLeft {
			_fdbd = _af.ReverseString(_fdbd)
		}
		if _agfdf {
			_cegad[_agaaa] = _fdbd
		} else {
			_cegad[_efcb-1] = _fdbd
		}
		_efcb--
	}
	if len(_cegad) != _aafae.NumRuns() {
		return _baga
	}
	_ceddb = _ed.Join(_cegad, "")
	return _ceddb
}

// SetBorder sets the cell's border style.
func (_gedef *TableCell) SetBorder(side CellBorderSide, style CellBorderStyle, width float64) {
	if style == CellBorderStyleSingle && side == CellBorderSideAll {
		_gedef._dbd = CellBorderStyleSingle
		_gedef._bfbdc = width
		_gedef._dgbfe = CellBorderStyleSingle
		_gedef._efgc = width
		_gedef._cebbe = CellBorderStyleSingle
		_gedef._bfgd = width
		_gedef._dbfgg = CellBorderStyleSingle
		_gedef._cfdc = width
	} else if style == CellBorderStyleDouble && side == CellBorderSideAll {
		_gedef._dbd = CellBorderStyleDouble
		_gedef._bfbdc = width
		_gedef._dgbfe = CellBorderStyleDouble
		_gedef._efgc = width
		_gedef._cebbe = CellBorderStyleDouble
		_gedef._bfgd = width
		_gedef._dbfgg = CellBorderStyleDouble
		_gedef._cfdc = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideLeft {
		_gedef._dbd = style
		_gedef._bfbdc = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideBottom {
		_gedef._dgbfe = style
		_gedef._efgc = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideRight {
		_gedef._cebbe = style
		_gedef._bfgd = width
	} else if (style == CellBorderStyleSingle || style == CellBorderStyleDouble) && side == CellBorderSideTop {
		_gedef._dbfgg = style
		_gedef._cfdc = width
	}
}
func (_defa *templateProcessor) parseTextRenderingModeAttr(_cceb, _fcdd string) TextRenderingMode {
	_df.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0074\u0065\u0078\u0074\u0020\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u006d\u006f\u0064e\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a \u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _cceb, _fcdd)
	_gcea := map[string]TextRenderingMode{"\u0066\u0069\u006c\u006c": TextRenderingModeFill, "\u0073\u0074\u0072\u006f\u006b\u0065": TextRenderingModeStroke, "f\u0069\u006c\u006c\u002d\u0073\u0074\u0072\u006f\u006b\u0065": TextRenderingModeFillStroke, "\u0069n\u0076\u0069\u0073\u0069\u0062\u006ce": TextRenderingModeInvisible, "\u0066i\u006c\u006c\u002d\u0063\u006c\u0069p": TextRenderingModeFillClip, "s\u0074\u0072\u006f\u006b\u0065\u002d\u0063\u006c\u0069\u0070": TextRenderingModeStrokeClip, "\u0066\u0069l\u006c\u002d\u0073t\u0072\u006f\u006b\u0065\u002d\u0063\u006c\u0069\u0070": TextRenderingModeFillStrokeClip, "\u0063\u006c\u0069\u0070": TextRenderingModeClip}[_fcdd]
	return _gcea
}

// Table allows organizing content in an rows X columns matrix, which can spawn across multiple pages.
type Table struct {
	_fdff           int
	_cfee           int
	_edag           int
	_agdgd          []float64
	_ceeb           []float64
	_dcgce          float64
	_ebfb           []*TableCell
	_fbca           []int
	_bfed           Positioning
	_fdcbfa, _eebda float64
	_ddcc           Margins
	_dabbc          bool
	_fdfff          int
	_bdffd          int
	_ggbfc          bool
	_fcacc          bool
}

// SetFillOpacity sets the fill opacity.
func (_gaceg *PolyBezierCurve) SetFillOpacity(opacity float64) { _gaceg._bbeag = opacity }

// Level returns the indentation level of the TOC line.
func (_aabce *TOCLine) Level() uint { return _aabce._dgcda }

// ColorCMYKFrom8bit creates a Color from c,m,y,k values (0-100).
// Example:
//   red := ColorCMYKFrom8Bit(0, 100, 100, 0)
func ColorCMYKFrom8bit(c, m, y, k byte) Color {
	return cmykColor{_fefa: _a.Min(float64(c), 100) / 100.0, _dcefe: _a.Min(float64(m), 100) / 100.0, _eefc: _a.Min(float64(y), 100) / 100.0, _bef: _a.Min(float64(k), 100) / 100.0}
}
func (_afdb cmykColor) ToRGB() (float64, float64, float64) {
	_cdfb := _afdb._bef
	return 1 - (_afdb._fefa*(1-_cdfb) + _cdfb), 1 - (_afdb._dcefe*(1-_cdfb) + _cdfb), 1 - (_afdb._eefc*(1-_cdfb) + _cdfb)
}

// GeneratePageBlocks draws the line on a new block representing the page.
// Implements the Drawable interface.
func (_baccc *Line) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var (
		_cegdf        []*Block
		_dgbb         = NewBlock(ctx.PageWidth, ctx.PageHeight)
		_cgcd         = ctx
		_gccg, _faega = _baccc._ggfd, ctx.PageHeight - _baccc._fadc
		_bgaa, _gdab  = _baccc._dccb, ctx.PageHeight - _baccc._gdec
	)
	_gffb := _baccc._edcee.IsRelative()
	if _gffb {
		ctx.X += _baccc._dacfg.Left
		ctx.Y += _baccc._dacfg.Top
		ctx.Width -= _baccc._dacfg.Left + _baccc._dacfg.Right
		ctx.Height -= _baccc._dacfg.Top + _baccc._dacfg.Bottom
		_gccg, _faega, _bgaa, _gdab = _baccc.computeCoords(ctx)
		if _baccc.Height() > ctx.Height {
			_cegdf = append(_cegdf, _dgbb)
			_dgbb = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_efed := ctx
			_efed.Y = ctx.Margins.Top + _baccc._dacfg.Top
			_efed.X = ctx.Margins.Left + _baccc._dacfg.Left
			_efed.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _baccc._dacfg.Top - _baccc._dacfg.Bottom
			_efed.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _baccc._dacfg.Left - _baccc._dacfg.Right
			ctx = _efed
			_gccg, _faega, _bgaa, _gdab = _baccc.computeCoords(ctx)
		}
	}
	_abfg := _cec.BasicLine{X1: _gccg, Y1: _faega, X2: _bgaa, Y2: _gdab, LineColor: _bcfe(_baccc._ggdcb), Opacity: _baccc._cgcbf, LineWidth: _baccc._ffgbg, LineStyle: _baccc._bfcbd, DashArray: _baccc._gbde, DashPhase: _baccc._afbf}
	_afgd, _gacc := _dgbb.setOpacity(1.0, _baccc._cgcbf)
	if _gacc != nil {
		return nil, ctx, _gacc
	}
	_bede, _, _gacc := _abfg.Draw(_afgd)
	if _gacc != nil {
		return nil, ctx, _gacc
	}
	if _gacc = _dgbb.addContentsByString(string(_bede)); _gacc != nil {
		return nil, ctx, _gacc
	}
	if _gffb {
		ctx.X = _cgcd.X
		ctx.Width = _cgcd.Width
		_fgba := _baccc.Height()
		ctx.Y += _fgba + _baccc._dacfg.Bottom
		ctx.Height -= _fgba
	} else {
		ctx = _cgcd
	}
	_cegdf = append(_cegdf, _dgbb)
	return _cegdf, ctx, nil
}

// AddAnnotation adds an annotation to the current block.
// The annotation will be added to the page the block will be rendered on.
func (_ece *Block) AddAnnotation(annotation *_f.PdfAnnotation) {
	for _, _aed := range _ece._cad {
		if _aed == annotation {
			return
		}
	}
	_ece._cad = append(_ece._cad, annotation)
}
func (_cgbg *Invoice) newColumn(_abege string, _cbga CellHorizontalAlignment) *InvoiceCell {
	_agdag := &InvoiceCell{_cgbg._cbgg, _abege}
	_agdag.Alignment = _cbga
	return _agdag
}

// Height returns the height of the chart.
func (_ggag *Chart) Height() float64 { return float64(_ggag._fbaa.Height()) }

// EnablePageWrap controls whether the table is wrapped across pages.
// If disabled, the table is moved in its entirety on a new page, if it
// does not fit in the available height. By default, page wrapping is enabled.
// If the height of the table is larger than an entire page, wrapping is
// enabled automatically in order to avoid unwanted behavior.
func (_bfge *Table) EnablePageWrap(enable bool) { _bfge._fcacc = enable }

// GetCoords returns the (x1, y1), (x2, y2) points defining the Line.
func (_cacd *Line) GetCoords() (float64, float64, float64, float64) {
	return _cacd._ggfd, _cacd._fadc, _cacd._dccb, _cacd._gdec
}

// NewColumn returns a new column for the line items invoice table.
func (_becd *Invoice) NewColumn(description string) *InvoiceCell {
	return _becd.newColumn(description, CellHorizontalAlignmentLeft)
}
func (_bdbeb *TableCell) height(_fbeb float64) float64 {
	var _fbeebd float64
	switch _cfb := _bdbeb._cbgcb.(type) {
	case *Paragraph:
		if _cfb._bbfc {
			_cfb.SetWidth(_fbeb - _bdbeb._cdgg - _cfb._eacf.Left - _cfb._eacf.Right)
		}
		_fbeebd = _cfb.Height() + _cfb._eacf.Top + _cfb._eacf.Bottom + 0.5*_cfb._gced*_cfb._abbg
	case *StyledParagraph:
		if _cfb._ecbf {
			_cfb.SetWidth(_fbeb - _bdbeb._cdgg - _cfb._afdba.Left - _cfb._afdba.Right)
		}
		_fbeebd = _cfb.Height() + _cfb._afdba.Top + _cfb._afdba.Bottom + 0.5*_cfb.getTextHeight()
	case *Image:
		_cfb.applyFitMode(_fbeb - _bdbeb._cdgg)
		_fbeebd = _cfb.Height() + _cfb._dbfe.Top + _cfb._dbfe.Bottom
	case *Table:
		_cfb.updateRowHeights(_fbeb - _bdbeb._cdgg - _cfb._ddcc.Left - _cfb._ddcc.Right)
		_fbeebd = _cfb.Height() + _cfb._ddcc.Top + _cfb._ddcc.Bottom
	case *List:
		_fbeebd = _cfb.tableHeight(_fbeb-_bdbeb._cdgg) + _cfb._bdccd.Top + _cfb._bdccd.Bottom
	case *Division:
		_fbeebd = _cfb.ctxHeight(_fbeb-_bdbeb._cdgg) + _cfb._fccf.Top + _cfb._fccf.Bottom
	case *Chart:
		_fbeebd = _cfb.Height() + _cfb._daea.Top + _cfb._daea.Bottom
	case *Rectangle:
		_fbeebd = _cfb.Height()
	case *Ellipse:
		_fbeebd = _cfb.Height()
	case *Line:
		_fbeebd = _cfb.Height() + _cfb._dacfg.Top + _cfb._dacfg.Bottom
	}
	return _fbeebd
}

// Text sets the text content of the Paragraph.
func (_daac *Paragraph) Text() string { return _daac._ebea }
func _bcfe(_bedd Color) _f.PdfColor {
	if _bedd == nil {
		_bedd = ColorBlack
	}
	switch _egde := _bedd.(type) {
	case cmykColor:
		return _f.NewPdfColorDeviceCMYK(_egde._fefa, _egde._dcefe, _egde._eefc, _egde._bef)
	}
	return _f.NewPdfColorDeviceRGB(_bedd.ToRGB())
}

// Width returns Image's document width.
func (_dgdef *Image) Width() float64 { return _dgdef._bcfd }

// GetOptimizer returns current PDF optimizer.
func (_bdf *Creator) GetOptimizer() _f.Optimizer { return _bdf._agg }

// TOC represents a table of contents component.
// It consists of a paragraph heading and a collection of
// table of contents lines.
// The representation of a table of contents line is as follows:
//       [number] [title]      [separator] [page]
// e.g.: Chapter1 Introduction ........... 1
type TOC struct {
	_gcdb   *StyledParagraph
	_fbagb  []*TOCLine
	_ccdbe  TextStyle
	_ecfgfa TextStyle
	_cacgb  TextStyle
	_ffed   TextStyle
	_ggdcba string
	_fgbea  float64
	_bcbf   Margins
	_ebfge  Positioning
	_ffbd   TextStyle
	_abbdf  bool
}

// TableCell defines a table cell which can contain a Drawable as content.
type TableCell struct {
	_fgddf       Color
	_degeb       _cec.LineStyle
	_dbd         CellBorderStyle
	_egcc        Color
	_bfbdc       float64
	_dgbfe       CellBorderStyle
	_dbee        Color
	_efgc        float64
	_cebbe       CellBorderStyle
	_gfag        Color
	_bfgd        float64
	_dbfgg       CellBorderStyle
	_ddec        Color
	_cfdc        float64
	_gbdd, _dbdf int
	_fefef       int
	_cegef       int
	_cbgcb       VectorDrawable
	_eaeac       CellHorizontalAlignment
	_geac        CellVerticalAlignment
	_cdgg        float64
	_fbdc        *Table
}

func (_gaf *Block) duplicate() *Block {
	_gf := &Block{}
	*_gf = *_gaf
	_caa := _eee.ContentStreamOperations{}
	_caa = append(_caa, *_gaf._ga...)
	_gf._ga = &_caa
	return _gf
}

// Width returns Ellipse's document width.
func (_cegb *Ellipse) Width() float64 { return _cegb._acea }

// Scale block by specified factors in the x and y directions.
func (_ced *Block) Scale(sx, sy float64) {
	_bfa := _eee.NewContentCreator().Scale(sx, sy).Operations()
	*_ced._ga = append(*_bfa, *_ced._ga...)
	_ced._ga.WrapIfNeeded()
	_ced._bf *= sx
	_ced._eda *= sy
}

// CreateTableOfContents sets a function to generate table of contents.
func (_efee *Creator) CreateTableOfContents(genTOCFunc func(_bfcc *TOC) error) {
	_efee._bdcf = genTOCFunc
}

// SetColor sets the line color. Use ColorRGBFromHex, ColorRGBFrom8bit or
// ColorRGBFromArithmetic to create the color object.
func (_ebecd *Line) SetColor(color Color) { _ebecd._ggdcb = color }
func _cdde() *Division                    { return &Division{_dcd: true} }

// SetFillColor sets the fill color.
func (_dgde *CurvePolygon) SetFillColor(color Color) { _dgde._dbca.FillColor = _bcfe(color) }

// ScaleToWidth scale Image to a specified width w, maintaining the aspect ratio.
func (_agdeb *Image) ScaleToWidth(w float64) {
	_fceb := _agdeb._gddc / _agdeb._bcfd
	_agdeb._bcfd = w
	_agdeb._gddc = w * _fceb
}

// Add adds a new line with the default style to the table of contents.
func (_abbfb *TOC) Add(number, title, page string, level uint) *TOCLine {
	_deecg := _abbfb.AddLine(_ffggc(TextChunk{Text: number, Style: _abbfb._ccdbe}, TextChunk{Text: title, Style: _abbfb._ecfgfa}, TextChunk{Text: page, Style: _abbfb._ffed}, level, _abbfb._ffbd))
	if _deecg == nil {
		return nil
	}
	_afcba := &_abbfb._bcbf
	_deecg.SetMargins(_afcba.Left, _afcba.Right, _afcba.Top, _afcba.Bottom)
	_deecg.SetLevelOffset(_abbfb._fgbea)
	_deecg.Separator.Text = _abbfb._ggdcba
	_deecg.Separator.Style = _abbfb._cacgb
	return _deecg
}

// InsertColumn inserts a column in the line items table at the specified index.
func (_ggbg *Invoice) InsertColumn(index uint, description string) *InvoiceCell {
	_cbca := uint(len(_ggbg._afeb))
	if index > _cbca {
		index = _cbca
	}
	_ebdb := _ggbg.NewColumn(description)
	_ggbg._afeb = append(_ggbg._afeb[:index], append([]*InvoiceCell{_ebdb}, _ggbg._afeb[index:]...)...)
	return _ebdb
}

// SetFillOpacity sets the fill opacity.
func (_badf *CurvePolygon) SetFillOpacity(opacity float64) { _badf._cdge = opacity }

// NewPolyline creates a new polyline.
func (_ggb *Creator) NewPolyline(points []_cec.Point) *Polyline { return _fcge(points) }

// Height returns the height of the list.
func (_abad *List) Height() float64 {
	var _cgfg float64
	for _, _cfae := range _abad._eddg {
		_cgfg += _cfae._cdef.Height()
	}
	return _cgfg
}

// GetMargins returns the margins of the chart (left, right, top, bottom).
func (_agcg *Chart) GetMargins() (float64, float64, float64, float64) {
	return _agcg._daea.Left, _agcg._daea.Right, _agcg._daea.Top, _agcg._daea.Bottom
}

// SetPos sets absolute positioning with specified coordinates.
func (_adeg *StyledParagraph) SetPos(x, y float64) {
	_adeg._geagd = PositionAbsolute
	_adeg._ggbc = x
	_adeg._aeegd = y
}

// GeneratePageBlocks generate the Page blocks.  Multiple blocks are generated if the contents wrap
// over multiple pages.
func (_fcg *Chapter) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_ebgb := ctx
	if _fcg._fabf.IsRelative() {
		ctx.X += _fcg._ccf.Left
		ctx.Y += _fcg._ccf.Top
		ctx.Width -= _fcg._ccf.Left + _fcg._ccf.Right
		ctx.Height -= _fcg._ccf.Top
	}
	_fecb, _fbbf, _fbg := _fcg._dbf.GeneratePageBlocks(ctx)
	if _fbg != nil {
		return _fecb, ctx, _fbg
	}
	ctx = _fbbf
	_egdg := ctx.X
	_cbba := ctx.Y - _fcg._dbf.Height()
	_bec := int64(ctx.Page)
	_ggd := _fcg.headingNumber()
	_bddg := _fcg.headingText()
	if _fcg._fcde {
		_gabg := _fcg._gdba.Add(_ggd, _fcg._bbe, _ae.FormatInt(_bec, 10), _fcg._aced)
		if _fcg._gdba._abbdf {
			_gabg.SetLink(_bec, _egdg, _cbba)
		}
	}
	if _fcg._baf == nil {
		_fcg._baf = _f.NewOutlineItem(_bddg, _f.NewOutlineDest(_bec-1, _egdg, _cbba))
		if _fcg._cadbb != nil {
			_fcg._cadbb._baf.Add(_fcg._baf)
		} else {
			_fcg._cdfd.Add(_fcg._baf)
		}
	} else {
		_feaf := &_fcg._baf.Dest
		_feaf.Page = _bec - 1
		_feaf.X = _egdg
		_feaf.Y = _cbba
	}
	for _, _edce := range _fcg._gecb {
		_agcb, _gfc, _gga := _edce.GeneratePageBlocks(ctx)
		if _gga != nil {
			return _fecb, ctx, _gga
		}
		if len(_agcb) < 1 {
			continue
		}
		_fecb[len(_fecb)-1].mergeBlocks(_agcb[0])
		_fecb = append(_fecb, _agcb[1:]...)
		ctx = _gfc
	}
	if _fcg._fabf.IsRelative() {
		ctx.X = _ebgb.X
	}
	if _fcg._fabf.IsAbsolute() {
		return _fecb, _ebgb, nil
	}
	return _fecb, ctx, nil
}
func _dbff(_eea [][]_cec.CubicBezierCurve) *CurvePolygon {
	return &CurvePolygon{_dbca: &_cec.CurvePolygon{Rings: _eea}, _cdge: 1.0, _cedec: 1.0}
}

// SetBorderWidth sets the border width.
func (_ebaa *PolyBezierCurve) SetBorderWidth(borderWidth float64) {
	_ebaa._cegf.BorderWidth = borderWidth
}

// GetMargins returns the left, right, top, bottom Margins.
func (_fefe *Table) GetMargins() (float64, float64, float64, float64) {
	return _fefe._ddcc.Left, _fefe._ddcc.Right, _fefe._ddcc.Top, _fefe._ddcc.Bottom
}

// SetTOC sets the table of content component of the creator.
// This method should be used when building a custom table of contents.
func (_adggf *Creator) SetTOC(toc *TOC) {
	if toc == nil {
		return
	}
	_adggf._cddg = toc
}

type templateProcessor struct {
	creator *Creator
	_ceac   []byte
	_daae   *TemplateOptions
	_cfeb   componentRenderer
}

// Append adds a new text chunk to the paragraph.
func (_ggea *StyledParagraph) Append(text string) *TextChunk {
	_bbec := NewTextChunk(text, _ggea._gdbeb)
	return _ggea.appendChunk(_bbec)
}
func _febe(_gbge *templateProcessor, _defe *templateNode) (interface{}, error) {
	return _gbge.parseChapter(_defe)
}

// SetNotes sets the notes section of the invoice.
func (_dgfa *Invoice) SetNotes(title, content string) { _dgfa._dbe = [2]string{title, content} }
func (_cbdgc *templateProcessor) parseDivision(_edcbc *templateNode) (interface{}, error) {
	_gfbgd := _cbdgc.creator.NewDivision()
	for _, _geaaf := range _edcbc._ddgde.Attr {
		_aged := _geaaf.Value
		switch _ecde := _geaaf.Name.Local; _ecde {
		case "\u0065\u006ea\u0062\u006c\u0065-\u0070\u0061\u0067\u0065\u002d\u0077\u0072\u0061\u0070":
			_gfbgd.EnablePageWrap(_cbdgc.parseBoolAttr(_ecde, _aged))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_bfab := _cbdgc.parseMarginAttr(_ecde, _aged)
			_gfbgd.SetMargins(_bfab.Left, _bfab.Right, _bfab.Top, _bfab.Bottom)
		case "\u0070a\u0064\u0064\u0069\u006e\u0067":
			_cefdd := _cbdgc.parseMarginAttr(_ecde, _aged)
			_gfbgd.SetPadding(_cefdd.Left, _cefdd.Right, _cefdd.Top, _cefdd.Bottom)
		default:
			_df.Log.Debug("U\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0064\u0069\u0076\u0069\u0073\u0069on\u0020\u0061\u0074\u0074r\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025s`\u002e\u0020S\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _ecde)
		}
	}
	return _gfbgd, nil
}

const (
	CellBorderStyleNone CellBorderStyle = iota
	CellBorderStyleSingle
	CellBorderStyleDouble
)

func (_edcb *Creator) initContext() {
	_edcb._bfec.X = _edcb._fade.Left
	_edcb._bfec.Y = _edcb._fade.Top
	_edcb._bfec.Width = _edcb._gee - _edcb._fade.Right - _edcb._fade.Left
	_edcb._bfec.Height = _edcb._agca - _edcb._fade.Bottom - _edcb._fade.Top
	_edcb._bfec.PageHeight = _edcb._agca
	_edcb._bfec.PageWidth = _edcb._gee
	_edcb._bfec.Margins = _edcb._fade
	_edcb._bfec._eag = _edcb.UnsupportedCharacterReplacement
}

// CurvePolygon represents a curve polygon shape.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type CurvePolygon struct {
	_dbca  *_cec.CurvePolygon
	_cdge  float64
	_cedec float64
}

func _egcac(_bdfgg string, _aadbbc, _afacf TextStyle) *TOC {
	_fbcd := _afacf
	_fbcd.FontSize = 14
	_bbfg := _ggaa(_fbcd)
	_bbfg.SetEnableWrap(true)
	_bbfg.SetTextAlignment(TextAlignmentLeft)
	_bbfg.SetMargins(0, 0, 0, 5)
	_befgg := _bbfg.Append(_bdfgg)
	_befgg.Style = _fbcd
	return &TOC{_gcdb: _bbfg, _fbagb: []*TOCLine{}, _ccdbe: _aadbbc, _ecfgfa: _aadbbc, _cacgb: _aadbbc, _ffed: _aadbbc, _ggdcba: "\u002e", _fgbea: 10, _bcbf: Margins{0, 0, 2, 2}, _ebfge: PositionRelative, _ffbd: _aadbbc, _abbdf: true}
}

// TextAlignment options for paragraph.
type TextAlignment int

// SetSideBorderColor sets the cell's side border color.
func (_acbbb *TableCell) SetSideBorderColor(side CellBorderSide, col Color) {
	switch side {
	case CellBorderSideAll:
		_acbbb._ddec = col
		_acbbb._dbee = col
		_acbbb._egcc = col
		_acbbb._gfag = col
	case CellBorderSideTop:
		_acbbb._ddec = col
	case CellBorderSideBottom:
		_acbbb._dbee = col
	case CellBorderSideLeft:
		_acbbb._egcc = col
	case CellBorderSideRight:
		_acbbb._gfag = col
	}
}

var (
	_cffa  = _ca.MustCompile("\u0028[\u005cw\u002d\u005d\u002b\u0029\u005c(\u0027\u0028.\u002b\u0029\u0027\u005c\u0029")
	_aeag  = _d.New("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0063\u0072\u0065a\u0074\u006f\u0072\u0020\u0069\u006e\u0073t\u0061\u006e\u0063\u0065")
	_dbea  = _d.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074e\u0020p\u0061\u0072\u0065\u006e\u0074\u0020\u006eo\u0064\u0065")
	_gaeg  = _d.New("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020c\u0068\u0069\u006cd\u0020n\u006f\u0064\u0065")
	_adbab = _d.New("\u0069n\u0076\u0061\u006c\u0069d\u0020\u0074\u0065\u006d\u0070l\u0061t\u0065 \u0072\u0065\u0073\u006f\u0075\u0072\u0063e")
)

func _gfde(_dgfd *Creator, _daecg []byte, _beafg *TemplateOptions, _adea componentRenderer) *templateProcessor {
	if _beafg == nil {
		_beafg = &TemplateOptions{}
	}
	_beafg.init()
	if _adea == nil {
		_adea = _dgfd
	}
	return &templateProcessor{creator: _dgfd, _ceac: _daecg, _daae: _beafg, _cfeb: _adea}
}

// TextStyle is a collection of properties that can be assigned to a text chunk.
type TextStyle struct {

	// Color represents the color of the text.
	Color Color

	// OutlineColor represents the color of the text outline.
	OutlineColor Color

	// Font represents the font the text will use.
	Font *_f.PdfFont

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
	_eag   rune
	_fdcc  []error
}

// TextVerticalAlignment controls the vertical position of the text
// in a styled paragraph.
type TextVerticalAlignment int

func _gebb(_ecag *Block, _dabdc *StyledParagraph, _abadc [][]*TextChunk, _fgaeg DrawContext) (DrawContext, [][]*TextChunk, error) {
	_cdfdb := 1
	_cebb := _fg.PdfObjectName(_eg.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _cdfdb))
	for _ecag._eb.HasFontByName(_cebb) {
		_cdfdb++
		_cebb = _fg.PdfObjectName(_eg.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _cdfdb))
	}
	_bfda := _ecag._eb.SetFontByName(_cebb, _dabdc._gdbeb.Font.ToPdfObject())
	if _bfda != nil {
		return _fgaeg, nil, _bfda
	}
	_cdfdb++
	_fbbdef := _cebb
	_gfbg := _dabdc._gdbeb.FontSize
	_dggc := _dabdc._geagd.IsRelative()
	var _aegf [][]_fg.PdfObjectName
	var _ebbff [][]*TextChunk
	var _gcag float64
	for _gacfd, _cccd := range _abadc {
		var _egdd []_fg.PdfObjectName
		var _bcgc float64
		if len(_cccd) > 0 {
			_bcgc = _cccd[0].Style.FontSize
		}
		for _, _agba := range _cccd {
			_cdfg := _agba.Style
			if _agba.Text != "" && _cdfg.FontSize > _bcgc {
				_bcgc = _cdfg.FontSize
			}
			_cebb = _fg.PdfObjectName(_eg.Sprintf("\u0046\u006f\u006e\u0074\u0025\u0064", _cdfdb))
			_adga := _ecag._eb.SetFontByName(_cebb, _cdfg.Font.ToPdfObject())
			if _adga != nil {
				return _fgaeg, nil, _adga
			}
			_egdd = append(_egdd, _cebb)
			_cdfdb++
		}
		_bcgc *= _dabdc._bged
		if _dggc && _gcag+_bcgc > _fgaeg.Height {
			_ebbff = _abadc[_gacfd:]
			_abadc = _abadc[:_gacfd]
			break
		}
		_gcag += _bcgc
		_aegf = append(_aegf, _egdd)
	}
	_daage, _ffbc, _ccaae := _dabdc.getLineMetrics(0)
	_accc, _aegea := _daage*_dabdc._bged, _ffbc*_dabdc._bged
	if len(_abadc) == 0 {
		return _fgaeg, _ebbff, nil
	}
	_efda := _eee.NewContentCreator()
	_efda.Add_q()
	_fcac := _aegea
	if _dabdc._ceab == TextVerticalAlignmentCenter {
		_fcac = _ffbc + (_daage+_ccaae-_ffbc)/2 + (_aegea-_ffbc)/2
	}
	_eggf := _fgaeg.PageHeight - _fgaeg.Y - _fcac
	_efda.Translate(_fgaeg.X, _eggf)
	_gebg := _eggf
	if _dabdc._daab != 0 {
		_efda.RotateDeg(_dabdc._daab)
	}
	if _dabdc._gcb == TextOverflowHidden {
		_efda.Add_re(0, -_gcag+_accc+1, _dabdc._bbef, _gcag).Add_W().Add_n()
	}
	_efda.Add_BT()
	var _adda []*_cec.BasicLine
	for _agac, _ccef := range _abadc {
		_dadfe := _fgaeg.X
		var _ebcee float64
		if len(_ccef) > 0 {
			_ebcee = _ccef[0].Style.FontSize
		}
		_daage, _, _ccaae = _dabdc.getLineMetrics(_agac)
		_aegea = (_daage + _ccaae)
		for _, _gfebe := range _ccef {
			_bbecd := &_gfebe.Style
			if _gfebe.Text != "" && _bbecd.FontSize > _ebcee {
				_ebcee = _bbecd.FontSize
			}
			if _aegea > _ebcee {
				_ebcee = _aegea
			}
		}
		if _agac != 0 {
			_efda.Add_TD(0, -_ebcee*_dabdc._bged)
			_gebg -= _ebcee * _dabdc._bged
		}
		_efddcc := _agac == len(_abadc)-1
		var (
			_eacea float64
			_cggc  float64
			_fcfg  float64
			_gegb  uint
		)
		var _afbb []float64
		for _, _gfba := range _ccef {
			_bebgec := &_gfba.Style
			if _bebgec.FontSize > _cggc {
				_cggc = _bebgec.FontSize
			}
			if _aegea > _cggc {
				_cggc = _aegea
			}
			_gfbe, _efeg := _bebgec.Font.GetRuneMetrics(' ')
			if !_efeg {
				return _fgaeg, nil, _d.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
			}
			var _fdgc uint
			var _edge float64
			_cgac := len(_gfba.Text)
			for _fffgg, _ceefc := range _gfba.Text {
				if _ceefc == ' ' {
					_fdgc++
					continue
				}
				if _ceefc == '\u000A' {
					continue
				}
				_babag, _gfdfg := _bebgec.Font.GetRuneMetrics(_ceefc)
				if !_gfdfg {
					_df.Log.Debug("\u0055\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0075\u006ee\u0020%\u0076\u0020\u0069\u006e\u0020\u0066\u006fn\u0074\u000a", _ceefc)
					return _fgaeg, nil, _d.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u0078\u0074\u0020\u0067\u006c\u0079p\u0068")
				}
				_edge += _bebgec.FontSize * _babag.Wx * _bebgec.horizontalScale()
				if _fffgg != _cgac-1 {
					_edge += _bebgec.CharSpacing * 1000.0
				}
			}
			_afbb = append(_afbb, _edge)
			_eacea += _edge
			_fcfg += float64(_fdgc) * _gfbe.Wx * _bebgec.FontSize * _bebgec.horizontalScale()
			_gegb += _fdgc
		}
		_cggc *= _dabdc._bged
		var _aceb []_fg.PdfObject
		_aaea := _dabdc._bbef * 1000.0
		if _dabdc._bcege == TextAlignmentJustify {
			if _gegb > 0 && !_efddcc {
				_fcfg = (_aaea - _eacea) / float64(_gegb) / _gfbg
			}
		} else if _dabdc._bcege == TextAlignmentCenter {
			_ebcbf := (_aaea - _eacea - _fcfg) / 2
			_fdacf := _ebcbf / _gfbg
			_aceb = append(_aceb, _fg.MakeFloat(-_fdacf))
			_dadfe += _ebcbf / 1000.0
		} else if _dabdc._bcege == TextAlignmentRight {
			_gaae := (_aaea - _eacea - _fcfg)
			_gabd := _gaae / _gfbg
			_aceb = append(_aceb, _fg.MakeFloat(-_gabd))
			_dadfe += _gaae / 1000.0
		}
		if len(_aceb) > 0 {
			_efda.Add_Tf(_fbbdef, _gfbg).Add_TL(_gfbg * _dabdc._bged).Add_TJ(_aceb...)
		}
		for _fdcbf, _babf := range _ccef {
			_cgag := &_babf.Style
			_eaebb := _fbbdef
			_cdeb := _gfbg
			_ggbfa := _cgag.OutlineColor != nil
			_fcaee := _cgag.HorizontalScaling != DefaultHorizontalScaling
			_fgfecb := _cgag.OutlineSize != 1
			if _fgfecb {
				_efda.Add_w(_cgag.OutlineSize)
			}
			_caagf := _cgag.RenderingMode != TextRenderingModeFill
			if _caagf {
				_efda.Add_Tr(int64(_cgag.RenderingMode))
			}
			_gebd := _cgag.CharSpacing != 0
			if _gebd {
				_efda.Add_Tc(_cgag.CharSpacing)
			}
			_cdaea := _cgag.TextRise != 0
			if _cdaea {
				_efda.Add_Ts(_cgag.TextRise)
			}
			if _dabdc._bcege != TextAlignmentJustify || _efddcc {
				_becgc, _gfeacg := _cgag.Font.GetRuneMetrics(' ')
				if !_gfeacg {
					return _fgaeg, nil, _d.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
				}
				_eaebb = _aegf[_agac][_fdcbf]
				_cdeb = _cgag.FontSize
				_fcfg = _becgc.Wx * _cgag.horizontalScale()
			}
			_bdfg := _cgag.Font.Encoder()
			var _gddcf []byte
			for _, _dfggc := range _babf.Text {
				if _dfggc == '\u000A' {
					continue
				}
				if _dfggc == ' ' {
					if len(_gddcf) > 0 {
						if _ggbfa {
							_efda.SetStrokingColor(_bcfe(_cgag.OutlineColor))
						}
						if _fcaee {
							_efda.Add_Tz(_cgag.HorizontalScaling)
						}
						_efda.SetNonStrokingColor(_bcfe(_cgag.Color)).Add_Tf(_aegf[_agac][_fdcbf], _cgag.FontSize).Add_TJ([]_fg.PdfObject{_fg.MakeStringFromBytes(_gddcf)}...)
						_gddcf = nil
					}
					if _fcaee {
						_efda.Add_Tz(DefaultHorizontalScaling)
					}
					_efda.Add_Tf(_eaebb, _cdeb).Add_TJ([]_fg.PdfObject{_fg.MakeFloat(-_fcfg)}...)
					_afbb[_fdcbf] += _fcfg * _cdeb
				} else {
					if _, _bccdb := _bdfg.RuneToCharcode(_dfggc); !_bccdb {
						_bfda = UnsupportedRuneError{Message: _eg.Sprintf("\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0072\u0075\u006e\u0065 \u0069\u006e\u0020\u0074\u0065\u0078\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u003a\u0020\u0025\u0023\u0078\u0020\u0028\u0025\u0063\u0029", _dfggc, _dfggc), Rune: _dfggc}
						_fgaeg._fdcc = append(_fgaeg._fdcc, _bfda)
						_df.Log.Debug(_bfda.Error())
						if _fgaeg._eag <= 0 {
							continue
						}
						_dfggc = _fgaeg._eag
					}
					_gddcf = append(_gddcf, _bdfg.Encode(string(_dfggc))...)
				}
			}
			if len(_gddcf) > 0 {
				if _ggbfa {
					_efda.SetStrokingColor(_bcfe(_cgag.OutlineColor))
				}
				if _fcaee {
					_efda.Add_Tz(_cgag.HorizontalScaling)
				}
				_efda.SetNonStrokingColor(_bcfe(_cgag.Color)).Add_Tf(_aegf[_agac][_fdcbf], _cgag.FontSize).Add_TJ([]_fg.PdfObject{_fg.MakeStringFromBytes(_gddcf)}...)
			}
			_cagd := _afbb[_fdcbf] / 1000.0
			if _cgag.Underline {
				_acdee := _cgag.UnderlineStyle.Color
				if _acdee == nil {
					_acdee = _babf.Style.Color
				}
				_bdcg, _dffea, _bbcb := _acdee.ToRGB()
				_fgfdg := _dadfe - _fgaeg.X
				_addaa := _gebg - _eggf + _cgag.TextRise - _cgag.UnderlineStyle.Offset
				_adda = append(_adda, &_cec.BasicLine{X1: _fgfdg, Y1: _addaa, X2: _fgfdg + _cagd, Y2: _addaa, LineWidth: _babf.Style.UnderlineStyle.Thickness, LineColor: _f.NewPdfColorDeviceRGB(_bdcg, _dffea, _bbcb)})
			}
			if _babf._ceddd != nil {
				var _gabeg *_fg.PdfObjectArray
				if !_babf._aaagg {
					switch _dcfc := _babf._ceddd.GetContext().(type) {
					case *_f.PdfAnnotationLink:
						_gabeg = _fg.MakeArray()
						_dcfc.Rect = _gabeg
						_ccefe, _gaad := _dcfc.Dest.(*_fg.PdfObjectArray)
						if _gaad && _ccefe.Len() == 5 {
							_dfdf, _fcbb := _ccefe.Get(1).(*_fg.PdfObjectName)
							if _fcbb && _dfdf.String() == "\u0058\u0059\u005a" {
								_eaaae, _beca := _fg.GetNumberAsFloat(_ccefe.Get(3))
								if _beca == nil {
									_ccefe.Set(3, _fg.MakeFloat(_fgaeg.PageHeight-_eaaae))
								}
							}
						}
					}
					_babf._aaagg = true
				}
				if _gabeg != nil {
					_dfeb := _cec.NewPoint(_dadfe-_fgaeg.X, _gebg+_cgag.TextRise-_eggf).Rotate(_dabdc._daab)
					_dfeb.X += _fgaeg.X
					_dfeb.Y += _eggf
					_feef, _fcefa, _cfcec, _eeaec := _feda(_cagd, _cggc, _dabdc._daab)
					_dfeb.X += _feef
					_dfeb.Y += _fcefa
					_gabeg.Clear()
					_gabeg.Append(_fg.MakeFloat(_dfeb.X))
					_gabeg.Append(_fg.MakeFloat(_dfeb.Y))
					_gabeg.Append(_fg.MakeFloat(_dfeb.X + _cfcec))
					_gabeg.Append(_fg.MakeFloat(_dfeb.Y + _eeaec))
				}
				_ecag.AddAnnotation(_babf._ceddd)
			}
			_dadfe += _cagd
			if _fgfecb {
				_efda.Add_w(1.0)
			}
			if _ggbfa {
				_efda.Add_RG(0.0, 0.0, 0.0)
			}
			if _caagf {
				_efda.Add_Tr(int64(TextRenderingModeFill))
			}
			if _gebd {
				_efda.Add_Tc(0)
			}
			if _cdaea {
				_efda.Add_Ts(0)
			}
			if _fcaee {
				_efda.Add_Tz(DefaultHorizontalScaling)
			}
		}
	}
	_efda.Add_ET()
	for _, _bbdf := range _adda {
		_efda.SetStrokingColor(_bbdf.LineColor).Add_w(_bbdf.LineWidth).Add_m(_bbdf.X1, _bbdf.Y1).Add_l(_bbdf.X2, _bbdf.Y2).Add_s()
	}
	_efda.Add_Q()
	_fbbe := _efda.Operations()
	_fbbe.WrapIfNeeded()
	_ecag.addContents(_fbbe)
	if _dggc {
		_gbgfc := _gcag
		_fgaeg.Y += _gbgfc
		_fgaeg.Height -= _gbgfc
		if _fgaeg.Inline {
			_fgaeg.X += _dabdc.Width() + _dabdc._afdba.Right
		}
	}
	return _fgaeg, _ebbff, nil
}

// SetPageLabels adds the specified page labels to the PDF file generated
// by the creator. See section 12.4.2 "Page Labels" (p. 382 PDF32000_2008).
// NOTE: for existing PDF files, the page label ranges object can be obtained
// using the model.PDFReader's GetPageLabels method.
func (_cea *Creator) SetPageLabels(pageLabels _fg.PdfObject) { _cea._bcca = pageLabels }
func (_dc *Block) setOpacity(_acc float64, _cca float64) (string, error) {
	if (_acc < 0 || _acc >= 1.0) && (_cca < 0 || _cca >= 1.0) {
		return "", nil
	}
	_gg := 0
	_aeg := _eg.Sprintf("\u0047\u0053\u0025\u0064", _gg)
	for _dc._eb.HasExtGState(_fg.PdfObjectName(_aeg)) {
		_gg++
		_aeg = _eg.Sprintf("\u0047\u0053\u0025\u0064", _gg)
	}
	_fcd := _fg.MakeDict()
	if _acc >= 0 && _acc < 1.0 {
		_fcd.Set("\u0063\u0061", _fg.MakeFloat(_acc))
	}
	if _cca >= 0 && _cca < 1.0 {
		_fcd.Set("\u0043\u0041", _fg.MakeFloat(_cca))
	}
	_bb := _dc._eb.AddExtGState(_fg.PdfObjectName(_aeg), _fcd)
	if _bb != nil {
		return "", _bb
	}
	return _aeg, nil
}

// PageBreak represents a page break for a chapter.
type PageBreak struct{}

// SetVerticalAlignment set the cell's vertical alignment of content.
// Can be one of:
// - CellHorizontalAlignmentTop
// - CellHorizontalAlignmentMiddle
// - CellHorizontalAlignmentBottom
func (_daeca *TableCell) SetVerticalAlignment(valign CellVerticalAlignment) { _daeca._geac = valign }

// SetBorderColor sets the cell's border color.
func (_eada *TableCell) SetBorderColor(col Color) {
	_eada._egcc = col
	_eada._dbee = col
	_eada._gfag = col
	_eada._ddec = col
}

// AddSubtable copies the cells of the subtable in the table, starting with the
// specified position. The table row and column indices are 1-based, which
// makes the position of the first cell of the first row of the table 1,1.
// The table is automatically extended if the subtable exceeds its columns.
// This can happen when the subtable has more columns than the table or when
// one or more columns of the subtable starting from the specified position
// exceed the last column of the table.
func (_cfed *Table) AddSubtable(row, col int, subtable *Table) {
	for _, _aggcd := range subtable._ebfb {
		_gaeb := &TableCell{}
		*_gaeb = *_aggcd
		_gaeb._fbdc = _cfed
		_gaeb._dbdf += col - 1
		if _gfeag := _cfed._cfee - (_gaeb._dbdf - 1); _gfeag < _gaeb._cegef {
			_cfed._cfee += _gaeb._cegef - _gfeag
			_cfed.resetColumnWidths()
			_df.Log.Debug("\u0054a\u0062l\u0065\u003a\u0020\u0073\u0075\u0062\u0074\u0061\u0062\u006c\u0065 \u0065\u0078\u0063\u0065e\u0064\u0073\u0020\u0064\u0065s\u0074\u0069\u006e\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0061\u0062\u006c\u0065\u002e\u0020\u0045\u0078\u0070\u0061\u006e\u0064\u0069\u006e\u0067\u0020\u0074\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0025\u0064\u0020\u0063\u006fl\u0075\u006d\u006e\u0073\u002e", _cfed._cfee)
		}
		_gaeb._gbdd += row - 1
		_fcff := subtable._ceeb[_aggcd._gbdd-1]
		if _gaeb._gbdd > _cfed._fdff {
			for _gaeb._gbdd > _cfed._fdff {
				_cfed._fdff++
				_cfed._ceeb = append(_cfed._ceeb, _cfed._dcgce)
			}
			_cfed._ceeb[_gaeb._gbdd-1] = _fcff
		} else {
			_cfed._ceeb[_gaeb._gbdd-1] = _a.Max(_cfed._ceeb[_gaeb._gbdd-1], _fcff)
		}
		_cfed._ebfb = append(_cfed._ebfb, _gaeb)
	}
	_ce.Slice(_cfed._ebfb, func(_gfda, _bcdff int) bool {
		_cdfga := _cfed._ebfb[_gfda]._gbdd
		_gggb := _cfed._ebfb[_bcdff]._gbdd
		if _cdfga < _gggb {
			return true
		}
		if _cdfga > _gggb {
			return false
		}
		return _cfed._ebfb[_gfda]._dbdf < _cfed._ebfb[_bcdff]._dbdf
	})
}

// Insert adds a new text chunk at the specified position in the paragraph.
func (_eefe *StyledParagraph) Insert(index uint, text string) *TextChunk {
	_fddc := uint(len(_eefe._eedad))
	if index > _fddc {
		index = _fddc
	}
	_ddff := NewTextChunk(text, _eefe._gdbeb)
	_eefe._eedad = append(_eefe._eedad[:index], append([]*TextChunk{_ddff}, _eefe._eedad[index:]...)...)
	_eefe.wrapText()
	return _ddff
}

const (
	CellBorderSideLeft CellBorderSide = iota
	CellBorderSideRight
	CellBorderSideTop
	CellBorderSideBottom
	CellBorderSideAll
)

func _agcaa(_dbffa ...interface{}) (map[string]interface{}, error) {
	_fcegf := len(_dbffa)
	if _fcegf%2 != 0 {
		_df.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065r\u0073\u0020\u0066\u006f\u0072\u0020\u0063\u0072\u0065\u0061\u0074i\u006e\u0067\u0020\u006d\u0061\u0070\u003a\u0020\u0025\u0064\u002e", _fcegf)
		return nil, _fg.ErrRangeError
	}
	_gdcgbd := map[string]interface{}{}
	for _bcfcf := 0; _bcfcf < _fcegf; _bcfcf += 2 {
		_cafa, _ccfdf := _dbffa[_bcfcf].(string)
		if !_ccfdf {
			_df.Log.Error("\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006d\u0061\u0070 \u006b\u0065\u0079\u0020t\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u0029\u002e\u0020\u0045\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u002e", _dbffa[_bcfcf])
			return nil, _fg.ErrTypeError
		}
		_gdcgbd[_cafa] = _dbffa[_bcfcf+1]
	}
	return _gdcgbd, nil
}

// MoveX moves the drawing context to absolute position x.
func (_cffgc *Creator) MoveX(x float64) { _cffgc._bfec.X = x }

// GetCoords returns coordinates of border.
func (_ggeg *border) GetCoords() (float64, float64) { return _ggeg._cdfc, _ggeg._aaab }
func (_cedcf *Table) wrapRow(_caegd int, _ddae DrawContext, _daabd float64) (bool, error) {
	if !_cedcf._ggbfc {
		return false, nil
	}
	var (
		_daad  = _cedcf._ebfb[_caegd]
		_gbba  = -1
		_cbgc  []*TableCell
		_ccbaf float64
		_agcbb bool
		_ffbbe = make([]float64, 0, len(_cedcf._agdgd))
	)
	_cacb := func(_bgadd *TableCell, _ceegd VectorDrawable, _dcggg bool) *TableCell {
		_ccdf := *_bgadd
		_ccdf._cbgcb = _ceegd
		if _dcggg {
			_ccdf._gbdd++
		}
		return &_ccdf
	}
	_cgeg := func(_dcggc int, _bfeb VectorDrawable) {
		var _cbea float64 = -1
		if _bfeb == nil {
			if _agcec := _ffbbe[_dcggc-_caegd]; _agcec > _ddae.Height {
				_bfeb = _cedcf._ebfb[_dcggc]._cbgcb
				_cedcf._ebfb[_dcggc]._cbgcb = nil
				_ffbbe[_dcggc-_caegd] = 0
				_cbea = _agcec
			}
		}
		_decd := _cacb(_cedcf._ebfb[_dcggc], _bfeb, true)
		_cbgc = append(_cbgc, _decd)
		if _cbea < 0 {
			_cbea = _decd.height(_ddae.Width)
		}
		if _cbea > _ccbaf {
			_ccbaf = _cbea
		}
	}
	for _faeb := _caegd; _faeb < len(_cedcf._ebfb); _faeb++ {
		_dbed := _cedcf._ebfb[_faeb]
		if _daad._gbdd != _dbed._gbdd {
			_gbba = _faeb
			break
		}
		_ddae.Width = _dbed.width(_cedcf._agdgd, _daabd)
		var _bafd VectorDrawable
		switch _febd := _dbed._cbgcb.(type) {
		case *StyledParagraph:
			if _gbdeb := _dbed.height(_ddae.Width); _gbdeb > _ddae.Height {
				_gcedc := _ddae
				_gcedc.Height = _a.Floor(_ddae.Height - _febd._afdba.Top - _febd._afdba.Bottom - 0.5*_febd.getTextHeight())
				_gadgb, _addd, _acbf := _febd.split(_gcedc)
				if _acbf != nil {
					return false, _acbf
				}
				if _gadgb != nil && _addd != nil {
					_febd = _gadgb
					_dbed = _cacb(_dbed, _gadgb, false)
					_cedcf._ebfb[_faeb] = _dbed
					_bafd = _addd
					_agcbb = true
				}
			}
		case *Division:
			if _eecdd := _dbed.height(_ddae.Width); _eecdd > _ddae.Height {
				_dgdg := _ddae
				_dgdg.Height = _a.Floor(_ddae.Height - _febd._fccf.Top - _febd._fccf.Bottom)
				_bdfb, _daabdc := _febd.split(_dgdg)
				if _bdfb != nil && _daabdc != nil {
					_febd = _bdfb
					_dbed = _cacb(_dbed, _bdfb, false)
					_cedcf._ebfb[_faeb] = _dbed
					_bafd = _daabdc
					_agcbb = true
					if _bdfb._bgeg != nil {
						_bdfb._bgeg.BorderRadiusBottomLeft = 0
						_bdfb._bgeg.BorderRadiusBottomRight = 0
					}
					if _daabdc._bgeg != nil {
						_daabdc._bgeg.BorderRadiusTopLeft = 0
						_daabdc._bgeg.BorderRadiusTopRight = 0
					}
				}
			}
		}
		_ffbbe = append(_ffbbe, _dbed.height(_ddae.Width))
		if _agcbb {
			if _cbgc == nil {
				_cbgc = make([]*TableCell, 0, len(_cedcf._agdgd))
				for _cddf := _caegd; _cddf < _faeb; _cddf++ {
					_cgeg(_cddf, nil)
				}
			}
			_cgeg(_faeb, _bafd)
		}
	}
	var _cacf float64
	for _, _begga := range _ffbbe {
		if _begga > _cacf {
			_cacf = _begga
		}
	}
	if _agcbb && _cacf < _ddae.Height {
		if _gbba < 0 {
			_gbba = len(_cedcf._ebfb)
		}
		_cdagg := _cedcf._ebfb[_gbba-1]._gbdd + _cedcf._ebfb[_gbba-1]._fefef - 1
		for _cgeb := _gbba; _cgeb < len(_cedcf._ebfb); _cgeb++ {
			_cedcf._ebfb[_cgeb]._gbdd++
		}
		_cedcf._ebfb = append(_cedcf._ebfb[:_gbba], append(_cbgc, _cedcf._ebfb[_gbba:]...)...)
		_cedcf._ceeb = append(_cedcf._ceeb[:_cdagg], append([]float64{_ccbaf}, _cedcf._ceeb[_cdagg:]...)...)
		_cedcf._ceeb[_daad._gbdd+_daad._fefef-2] = _cacf
	}
	return _agcbb, nil
}

// PageSize represents the page size as a 2 element array representing the width and height in PDF document units (points).
type PageSize [2]float64

// Color interface represents colors in the PDF creator.
type Color interface {
	ToRGB() (float64, float64, float64)
}

// SetTextVerticalAlignment sets the vertical alignment of the text within the
// bounds of the styled paragraph.
func (_bbdab *StyledParagraph) SetTextVerticalAlignment(align TextVerticalAlignment) {
	_bbdab._ceab = align
}

// SetBorderColor sets border color.
func (_baad *Rectangle) SetBorderColor(col Color) { _baad._afcb = col }

// Cols returns the total number of columns the table has.
func (_eeca *Table) Cols() int { return _eeca._cfee }

// SetText sets the text content of the Paragraph.
func (_adgdc *Paragraph) SetText(text string) { _adgdc._ebea = text }
func (_gdecb *Table) updateRowHeights(_bdbd float64) {
	for _, _bgfc := range _gdecb._ebfb {
		_gcddg := _bgfc.width(_gdecb._agdgd, _bdbd)
		_bafe := _gdecb._ceeb[_bgfc._gbdd+_bgfc._fefef-2]
		if _agded := _bgfc.height(_gcddg); _agded > _bafe {
			_cbgaf := _agded / float64(_bgfc._fefef)
			for _aade := 1; _aade <= _bgfc._fefef; _aade++ {
				if _cbgaf > _gdecb._ceeb[_bgfc._gbdd+_aade-2] {
					_gdecb._ceeb[_bgfc._gbdd+_aade-2] = _cbgaf
				}
			}
		}
	}
}
func _bfeg(_fadb [][]_cec.Point) *Polygon {
	return &Polygon{_ccbeg: &_cec.Polygon{Points: _fadb}, _fead: 1.0, _afab: 1.0}
}

// Context returns the current drawing context.
func (_aab *Creator) Context() DrawContext { return _aab._bfec }

// DueDate returns the invoice due date description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_becf *Invoice) DueDate() (*InvoiceCell, *InvoiceCell) { return _becf._debe[0], _becf._debe[1] }

// Color returns the color of the line.
func (_cedef *Line) Color() Color { return _cedef._ggdcb }
func _dfc(_cdd, _bceg, _bgc, _eba float64) *border {
	_fgae := &border{}
	_fgae._cdfc = _cdd
	_fgae._aaab = _bceg
	_fgae._dgf = _bgc
	_fgae._gbgc = _eba
	_fgae._gad = ColorBlack
	_fgae._bdd = ColorBlack
	_fgae._bfe = ColorBlack
	_fgae._bfae = ColorBlack
	_fgae._ebc = 0
	_fgae._gbe = 0
	_fgae._afeg = 0
	_fgae._fe = 0
	_fgae.LineStyle = _cec.LineStyleSolid
	return _fgae
}
func (_gecc *templateProcessor) parseAttrPropList(_ecgc string) map[string]string {
	_cfbe := _ed.Fields(_ecgc)
	if len(_cfbe) == 0 {
		return nil
	}
	_acegd := map[string]string{}
	for _, _caage := range _cfbe {
		_gcgfd := _cffa.FindStringSubmatch(_caage)
		if len(_gcgfd) < 3 {
			continue
		}
		_dedd, _ddac := _ed.TrimSpace(_gcgfd[1]), _gcgfd[2]
		if _dedd == "" {
			continue
		}
		_acegd[_dedd] = _ddac
	}
	return _acegd
}
func _fcef(_beb, _ggfg, _ddgd, _cdea, _bcbg, _ggab float64) *Curve {
	_bbad := &Curve{}
	_bbad._ggce = _beb
	_bbad._bdcce = _ggfg
	_bbad._dafb = _ddgd
	_bbad._feec = _cdea
	_bbad._fgaeb = _bcbg
	_bbad._aaeb = _ggab
	_bbad._cddgg = ColorBlack
	_bbad._bbf = 1.0
	return _bbad
}
func (_fedf *StyledParagraph) wrapText() error { return _fedf.wrapChunks(true) }

// Invoice represents a configurable invoice template.
type Invoice struct {
	_bgdd  string
	_fbgcc *Image
	_ebca  *InvoiceAddress
	_bgfe  *InvoiceAddress
	_ceb   string
	_edeb  [2]*InvoiceCell
	_fgcge [2]*InvoiceCell
	_debe  [2]*InvoiceCell
	_agef  [][2]*InvoiceCell
	_afeb  []*InvoiceCell
	_eead  [][]*InvoiceCell
	_faeg  [2]*InvoiceCell
	_defg  [2]*InvoiceCell
	_gce   [][2]*InvoiceCell
	_dbe   [2]string
	_dcdf  [2]string
	_bgfee [][2]string
	_ffdf  TextStyle
	_gfdd  TextStyle
	_acgbf TextStyle
	_ddf   TextStyle
	_eeea  TextStyle
	_egf   TextStyle
	_cacg  TextStyle
	_fcabd InvoiceCellProps
	_cbgg  InvoiceCellProps
	_defdd InvoiceCellProps
	_fabge InvoiceCellProps
	_dfgf  Positioning
}

func (_cdbg *Table) resetColumnWidths() {
	_cdbg._agdgd = []float64{}
	_bgedf := float64(1.0) / float64(_cdbg._cfee)
	for _cdfab := 0; _cdfab < _cdbg._cfee; _cdfab++ {
		_cdbg._agdgd = append(_cdbg._agdgd, _bgedf)
	}
}

// SetTitleStyle sets the style properties of the invoice title.
func (_eaga *Invoice) SetTitleStyle(style TextStyle) { _eaga._acgbf = style }

// HorizontalAlignment represents the horizontal alignment of components
// within a page.
type HorizontalAlignment int

// Margins returns the margins of the component.
func (_gaab *Division) Margins() (_bbda, _ecae, _fcfb, _eeedg float64) {
	return _gaab._fccf.Left, _gaab._fccf.Right, _gaab._fccf.Top, _gaab._fccf.Bottom
}

// RotateDeg rotates the current active page by angle degrees.  An error is returned on failure,
// which can be if there is no currently active page, or the angleDeg is not a multiple of 90 degrees.
func (_gca *Creator) RotateDeg(angleDeg int64) error {
	_dfdg := _gca.getActivePage()
	if _dfdg == nil {
		_df.Log.Debug("F\u0061\u0069\u006c\u0020\u0074\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0065\u003a\u0020\u006e\u006f\u0020p\u0061\u0067\u0065\u0020\u0063\u0075\u0072\u0072\u0065\u006etl\u0079\u0020\u0061c\u0074i\u0076\u0065")
		return _d.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	if angleDeg%90 != 0 {
		_df.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0067e\u0020\u0072\u006f\u0074\u0061\u0074\u0069on\u0020\u0061\u006e\u0067l\u0065\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006dul\u0074\u0069p\u006c\u0065\u0020\u006f\u0066\u0020\u0039\u0030")
		return _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	var _ceec int64
	if _dfdg.Rotate != nil {
		_ceec = *(_dfdg.Rotate)
	}
	_ceec += angleDeg
	_dfdg.Rotate = &_ceec
	return nil
}

// SetLineSeparatorStyle sets the style for the separator part of all new
// lines of the table of contents.
func (_dagdg *TOC) SetLineSeparatorStyle(style TextStyle) { _dagdg._cacgb = style }

// SetOpacity sets the opacity of the line (0-1).
func (_bbbd *Line) SetOpacity(opacity float64) { _bbbd._cgcbf = opacity }

// NoteStyle returns the style properties used to render the content of the
// invoice note sections.
func (_gef *Invoice) NoteStyle() TextStyle { return _gef._egf }
func _effa(_fdbg *Table, _afec DrawContext) ([]*Block, DrawContext, error) {
	var _gbdb []*Block
	_dbad := NewBlock(_afec.PageWidth, _afec.PageHeight)
	_fdbg.updateRowHeights(_afec.Width - _fdbg._ddcc.Left - _fdbg._ddcc.Right)
	_agbc := _fdbg._ddcc.Top
	if _fdbg._bfed.IsRelative() && !_fdbg._fcacc {
		_cbee := _fdbg.Height()
		if _cbee > _afec.Height-_fdbg._ddcc.Top && _cbee <= _afec.PageHeight-_afec.Margins.Top-_afec.Margins.Bottom {
			_gbdb = []*Block{NewBlock(_afec.PageWidth, _afec.PageHeight-_afec.Y)}
			var _gcba error
			if _, _afec, _gcba = _acdfe().GeneratePageBlocks(_afec); _gcba != nil {
				return nil, _afec, _gcba
			}
			_agbc = 0
		}
	}
	_aafab := _afec
	if _fdbg._bfed.IsAbsolute() {
		_afec.X = _fdbg._fdcbfa
		_afec.Y = _fdbg._eebda
	} else {
		_afec.X += _fdbg._ddcc.Left
		_afec.Y += _agbc
		_afec.Width -= _fdbg._ddcc.Left + _fdbg._ddcc.Right
		_afec.Height -= _agbc
	}
	_ecagg := _afec.Width
	_gaafd := _afec.X
	_cefa := _afec.Y
	_fbff := _afec.Height
	_ecfgf := 0
	_cafdd, _dbaa := -1, -1
	if _fdbg._dabbc {
		for _eaaf, _agbfe := range _fdbg._ebfb {
			if _agbfe._gbdd < _fdbg._fdfff {
				continue
			}
			if _agbfe._gbdd > _fdbg._bdffd {
				break
			}
			if _cafdd < 0 {
				_cafdd = _eaaf
			}
			_dbaa = _eaaf
		}
	}
	if _defge := _fdbg.wrapContent(_afec); _defge != nil {
		return nil, _afec, _defge
	}
	_fdbg.updateRowHeights(_afec.Width - _fdbg._ddcc.Left - _fdbg._ddcc.Right)
	var (
		_agga  bool
		_cfeeg int
		_cbed  int
		_daba  bool
		_aaaeg int
		_bada  error
	)
	for _bccac := 0; _bccac < len(_fdbg._ebfb); _bccac++ {
		_cbggd := _fdbg._ebfb[_bccac]
		_fcda := _cbggd.width(_fdbg._agdgd, _ecagg)
		_gefc := float64(0.0)
		for _caba := 0; _caba < _cbggd._dbdf-1; _caba++ {
			_gefc += _fdbg._agdgd[_caba] * _ecagg
		}
		_egcfa := float64(0.0)
		for _geaa := _ecfgf; _geaa < _cbggd._gbdd-1; _geaa++ {
			_egcfa += _fdbg._ceeb[_geaa]
		}
		_afec.Height = _fbff - _egcfa
		_bfccb := float64(0.0)
		for _bgec := 0; _bgec < _cbggd._fefef; _bgec++ {
			_bfccb += _fdbg._ceeb[_cbggd._gbdd+_bgec-1]
		}
		_fbbbd := _daba && _cbggd._gbdd != _aaaeg
		_aaaeg = _cbggd._gbdd
		if _fbbbd || _bfccb > _afec.Height {
			if _fdbg._ggbfc && !_daba {
				_daba, _bada = _fdbg.wrapRow(_bccac, _afec, _ecagg)
				if _bada != nil {
					return nil, _afec, _bada
				}
				if _daba {
					_bccac--
					continue
				}
			}
			_gbdb = append(_gbdb, _dbad)
			_dbad = NewBlock(_afec.PageWidth, _afec.PageHeight)
			_gaafd = _afec.Margins.Left + _fdbg._ddcc.Left
			_cefa = _afec.Margins.Top
			_afec.Height = _afec.PageHeight - _afec.Margins.Top - _afec.Margins.Bottom
			_afec.Page++
			_fbff = _afec.Height
			_ecfgf = _cbggd._gbdd - 1
			_egcfa = 0
			_daba = false
			if _fdbg._dabbc && _cafdd >= 0 {
				_cfeeg = _bccac
				_bccac = _cafdd - 1
				_cbed = _ecfgf
				_ecfgf = _fdbg._fdfff - 1
				_agga = true
				if _cbggd._fefef > (_fdbg._fdff-_aaaeg) || (_cbggd._fefef > 1 && _bccac < 0) {
					_df.Log.Debug("\u0054a\u0062\u006ce\u0020\u0068\u0065a\u0064\u0065\u0072\u0020\u0072\u006f\u0077s\u0070\u0061\u006e\u0020\u0065\u0078c\u0065\u0065\u0064\u0073\u0020\u0061\u0076\u0061\u0069\u006c\u0061b\u006c\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u002e")
					_agga = false
					_cafdd, _dbaa = -1, -1
				}
				continue
			}
			if _fbbbd {
				_bccac--
				continue
			}
		}
		_afec.Width = _fcda
		_afec.X = _gaafd + _gefc
		_afec.Y = _cefa + _egcfa
		_dcbb := _dfc(_afec.X, _afec.Y, _fcda, _bfccb)
		if _cbggd._fgddf != nil {
			_dcbb.SetFillColor(_cbggd._fgddf)
		}
		_dcbb.LineStyle = _cbggd._degeb
		_dcbb._afg = _cbggd._dbd
		_dcbb._bgad = _cbggd._cebbe
		_dcbb._faaa = _cbggd._dbfgg
		_dcbb._aeba = _cbggd._dgbfe
		if _cbggd._egcc != nil {
			_dcbb.SetColorLeft(_cbggd._egcc)
		}
		if _cbggd._dbee != nil {
			_dcbb.SetColorBottom(_cbggd._dbee)
		}
		if _cbggd._gfag != nil {
			_dcbb.SetColorRight(_cbggd._gfag)
		}
		if _cbggd._ddec != nil {
			_dcbb.SetColorTop(_cbggd._ddec)
		}
		_dcbb.SetWidthBottom(_cbggd._efgc)
		_dcbb.SetWidthLeft(_cbggd._bfbdc)
		_dcbb.SetWidthRight(_cbggd._bfgd)
		_dcbb.SetWidthTop(_cbggd._cfdc)
		_cagge := _dbad.Draw(_dcbb)
		if _cagge != nil {
			_df.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cagge)
		}
		if _cbggd._cbgcb != nil {
			_begg := _cbggd._cbgcb.Width()
			_eafa := _cbggd._cbgcb.Height()
			_bdfdf := 0.0
			switch _eagc := _cbggd._cbgcb.(type) {
			case *Paragraph:
				if _eagc._bbfc {
					_begg = _eagc.getMaxLineWidth() / 1000.0
				}
				_begg += _eagc._eacf.Left + _eagc._eacf.Right
				_eafa += _eagc._eacf.Top + _eagc._eacf.Bottom
			case *StyledParagraph:
				if _eagc._ecbf {
					_begg = _eagc.getMaxLineWidth() / 1000.0
				}
				_adba, _ffge, _fgda := _eagc.getLineMetrics(0)
				_dbcc, _febc := _adba*_eagc._bged, _ffge*_eagc._bged
				if _eagc._ceab == TextVerticalAlignmentCenter {
					_bdfdf = _febc - (_ffge + (_adba+_fgda-_ffge)/2 + (_febc-_ffge)/2)
				}
				if len(_eagc._bfaa) == 1 {
					_eafa = _dbcc
				} else {
					_eafa = _eafa - _febc + _dbcc
				}
				_bdfdf += _dbcc - _febc
				switch _cbggd._geac {
				case CellVerticalAlignmentTop:
					_bdfdf += _dbcc * 0.5
				case CellVerticalAlignmentBottom:
					_bdfdf -= _dbcc * 0.5
				}
				_begg += _eagc._afdba.Left + _eagc._afdba.Right
				_eafa += _eagc._afdba.Top + _eagc._afdba.Bottom
			case *Table:
				_begg = _fcda
			case *List:
				_begg = _fcda
			case *Division:
				_begg = _fcda
			case *Chart:
				_begg = _fcda
			case *Line:
				_eafa += _eagc._dacfg.Top + _eagc._dacfg.Bottom
				_bdfdf -= _eagc.Height() / 2
			}
			switch _cbggd._eaeac {
			case CellHorizontalAlignmentLeft:
				_afec.X += _cbggd._cdgg
				_afec.Width -= _cbggd._cdgg
			case CellHorizontalAlignmentCenter:
				if _dddd := _fcda - _begg; _dddd > 0 {
					_afec.X += _dddd / 2
					_afec.Width -= _dddd / 2
				}
			case CellHorizontalAlignmentRight:
				if _fcda > _begg {
					_afec.X = _afec.X + _fcda - _begg - _cbggd._cdgg
					_afec.Width -= _cbggd._cdgg
				}
			}
			_gbggcg := _afec.Y
			_cfdaa := _afec.Height
			_afec.Y += _bdfdf
			switch _cbggd._geac {
			case CellVerticalAlignmentTop:
			case CellVerticalAlignmentMiddle:
				if _gdff := _bfccb - _eafa; _gdff > 0 {
					_afec.Y += _gdff / 2
					_afec.Height -= _gdff / 2
				}
			case CellVerticalAlignmentBottom:
				if _bfccb > _eafa {
					_afec.Y = _afec.Y + _bfccb - _eafa
					_afec.Height = _bfccb
				}
			}
			_dgef := _dbad.DrawWithContext(_cbggd._cbgcb, _afec)
			if _dgef != nil {
				_df.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dgef)
			}
			_afec.Y = _gbggcg
			_afec.Height = _cfdaa
		}
		_afec.Y += _bfccb
		_afec.Height -= _bfccb
		if _agga && _bccac+1 > _dbaa {
			_cefa += _egcfa + _bfccb
			_fbff -= _bfccb + _egcfa
			_ecfgf = _cbed
			_bccac = _cfeeg - 1
			_agga = false
		}
	}
	_gbdb = append(_gbdb, _dbad)
	if _fdbg._bfed.IsAbsolute() {
		return _gbdb, _aafab, nil
	}
	_afec.X = _aafab.X
	_afec.Width = _aafab.Width
	_afec.Y += _fdbg._ddcc.Bottom
	_afec.Height -= _fdbg._ddcc.Bottom
	return _gbdb, _afec, nil
}

// SetStyle sets the style of the line (solid or dashed).
func (_bgfa *Line) SetStyle(style _cec.LineStyle) { _bgfa._bfcbd = style }

// Block contains a portion of PDF Page contents. It has a width and a position and can
// be placed anywhere on a Page.  It can even contain a whole Page, and is used in the creator
// where each Drawable object can output one or more blocks, each representing content for separate pages
// (typically needed when Page breaks occur).
type Block struct {
	_ga      *_eee.ContentStreamOperations
	_eb      *_f.PdfPageResources
	_ec      Positioning
	_fa, _ea float64
	_bf      float64
	_eda     float64
	_ag      float64
	_ac      Margins
	_cad     []*_f.PdfAnnotation
}

// Flip flips the active page on the specified axes.
// If `flipH` is true, the page is flipped horizontally. Similarly, if `flipV`
// is true, the page is flipped vertically. If both are true, the page is
// flipped both horizontally and vertically.
// NOTE: the flip transformations are applied when the creator is finalized,
// which is at write time in most cases.
func (_bgagb *Creator) Flip(flipH, flipV bool) error {
	_gdcb := _bgagb.getActivePage()
	if _gdcb == nil {
		return _d.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	_gdg, _bdfdd := _bgagb._gba[_gdcb]
	if !_bdfdd {
		_gdg = &pageTransformations{}
		_bgagb._gba[_gdcb] = _gdg
	}
	_gdg._gdd = flipH
	_gdg._cbcb = flipV
	return nil
}
func _gaafg(_ddbf string) bool {
	_dbeae := func(_feadb rune) bool { return _feadb == '\u000A' }
	_ebfeg := _ed.TrimFunc(_ddbf, _dbeae)
	_ebffg := _af.Paragraph{}
	_, _bcaec := _ebffg.SetString(_ebfeg)
	if _bcaec != nil {
		return true
	}
	_eceg, _bcaec := _ebffg.Order()
	if _bcaec != nil {
		return true
	}
	if _eceg.NumRuns() < 1 {
		return true
	}
	return _ebffg.IsLeftToRight()
}

// BuyerAddress returns the buyer address used in the invoice template.
func (_bcac *Invoice) BuyerAddress() *InvoiceAddress { return _bcac._ebca }

// SetLogo sets the logo of the invoice.
func (_fgdb *Invoice) SetLogo(logo *Image) { _fgdb._fbgcc = logo }

// Style returns the style of the line.
func (_cagg *Line) Style() _cec.LineStyle { return _cagg._bfcbd }

// FilledCurve represents a closed path of Bezier curves with a border and fill.
type FilledCurve struct {
	_bgf          []_cec.CubicBezierCurve
	FillEnabled   bool
	_bffg         Color
	BorderEnabled bool
	BorderWidth   float64
	_aac          Color
}

func _bac(_gcfd *Chapter, _ebgea *TOC, _adc *_f.Outline, _ecf string, _def int, _dab TextStyle) *Chapter {
	var _fgaa uint = 1
	if _gcfd != nil {
		_fgaa = _gcfd._aced + 1
	}
	_acd := &Chapter{_aebc: _def, _bbe: _ecf, _bdc: true, _fcde: true, _cadbb: _gcfd, _gdba: _ebgea, _cdfd: _adc, _gecb: []Drawable{}, _aced: _fgaa}
	_gfee := _ebcb(_acd.headingText(), _dab)
	_gfee.SetFont(_dab.Font)
	_gfee.SetFontSize(_dab.FontSize)
	_acd._dbf = _gfee
	return _acd
}

// Positioning represents the positioning type for drawing creator components (relative/absolute).
type Positioning int

// FrontpageFunctionArgs holds the input arguments to a front page drawing function.
// It is designed as a struct, so additional parameters can be added in the future with backwards
// compatibility.
type FrontpageFunctionArgs struct {
	PageNum    int
	TotalPages int
}

// MultiRowCell makes a new cell with the specified row span and inserts it
// into the table at the current position.
func (_ccadb *Table) MultiRowCell(rowspan int) *TableCell { return _ccadb.MultiCell(rowspan, 1) }

// CellVerticalAlignment defines the table cell's vertical alignment.
type CellVerticalAlignment int

// Width returns the cell's width based on the input draw context.
func (_cdcc *TableCell) Width(ctx DrawContext) float64 {
	_cdgd := float64(0.0)
	for _cadd := 0; _cadd < _cdcc._cegef; _cadd++ {
		_cdgd += _cdcc._fbdc._agdgd[_cdcc._dbdf+_cadd-1]
	}
	_ddbe := ctx.Width * _cdgd
	return _ddbe
}

// AddInternalLink adds a new internal link to the paragraph.
// The text parameter represents the text that is displayed.
// The user is taken to the specified page, at the specified x and y
// coordinates. Position 0, 0 is at the top left of the page.
// The zoom of the destination page is controlled with the zoom
// parameter. Pass in 0 to keep the current zoom value.
func (_egea *StyledParagraph) AddInternalLink(text string, page int64, x, y, zoom float64) *TextChunk {
	_addb := NewTextChunk(text, _egea._bdbf)
	_addb._ceddd = _eabf(page-1, x, y, zoom)
	return _egea.appendChunk(_addb)
}

// Creator is a wrapper around functionality for creating PDF reports and/or adding new
// content onto imported PDF pages, etc.
type Creator struct {

	// Errors keeps error messages that should not interrupt pdf processing and to be checked later.
	Errors []error

	// UnsupportedCharacterReplacement is character that will be used to replace unsupported glyph.
	// The value will be passed to drawing context.
	UnsupportedCharacterReplacement rune
	_gadb                           []*_f.PdfPage
	_caec                           map[*_f.PdfPage]*Block
	_gba                            map[*_f.PdfPage]*pageTransformations
	_efde                           *_f.PdfPage
	_dda                            PageSize
	_bfec                           DrawContext
	_fade                           Margins
	_gee, _agca                     float64
	_dedf                           int
	_bcec                           func(_cede FrontpageFunctionArgs)
	_bdcf                           func(_agcf *TOC) error
	_eacd                           func(_bfd *Block, _ggf HeaderFunctionArgs)
	_eeb                            func(_ccge *Block, _cabe FooterFunctionArgs)
	_ebd                            func(_fge PageFinalizeFunctionArgs) error
	_eege                           func(_eade *_f.PdfWriter) error
	_cgga                           bool

	// Controls whether a table of contents will be generated.
	AddTOC bool

	// CustomTOC specifies if the TOC is rendered by the user.
	// When the `CustomTOC` field is set to `true`, the default TOC component is not rendered.
	// Instead the TOC is drawn by the user, in the callback provided to
	// the `Creator.CreateTableOfContents` method.
	// If `CustomTOC` is set to `false`, the callback provided to
	// `Creator.CreateTableOfContents` customizes the style of the automatically generated TOC component.
	CustomTOC bool
	_cddg     *TOC

	// Controls whether outlines will be generated.
	AddOutlines bool
	_dfd        *_f.Outline
	_afed       *_f.PdfOutlineTreeNode
	_eafg       *_f.PdfAcroForm
	_bcca       _fg.PdfObject
	_agg        _f.Optimizer
	_dedb       []*_f.PdfFont
	_ffgd       *_f.PdfFont
	_beea       *_f.PdfFont
}

// CurRow returns the currently active cell's row number.
func (_affa *Table) CurRow() int { _cffed := (_affa._edag-1)/_affa._cfee + 1; return _cffed }
func _gdbg(_bdba *templateProcessor, _gcbd *templateNode) (interface{}, error) {
	return _bdba.parseImage(_gcbd)
}
func (_fgf *Block) drawToPage(_ff *_f.PdfPage) error {
	_dba := &_eee.ContentStreamOperations{}
	if _ff.Resources == nil {
		_ff.Resources = _f.NewPdfPageResources()
	}
	_fca := _ggg(_dba, _ff.Resources, _fgf._ga, _fgf._eb)
	if _fca != nil {
		return _fca
	}
	if _fca = _gfff(_fgf._eb, _ff.Resources); _fca != nil {
		return _fca
	}
	if _fca = _ff.AppendContentBytes(_dba.Bytes(), true); _fca != nil {
		return _fca
	}
	for _, _bbb := range _fgf._cad {
		_ff.AddAnnotation(_bbb)
	}
	return nil
}

// GeneratePageBlocks generates a page break block.
func (_ffbgd *PageBreak) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_aeefg := []*Block{NewBlock(ctx.PageWidth, ctx.PageHeight-ctx.Y), NewBlock(ctx.PageWidth, ctx.PageHeight)}
	ctx.Page++
	_ccbc := ctx
	_ccbc.Y = ctx.Margins.Top
	_ccbc.X = ctx.Margins.Left
	_ccbc.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
	_ccbc.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right
	ctx = _ccbc
	return _aeefg, ctx, nil
}

// NewImageFromGoImage creates an Image from a go image.Image data structure.
func (_gda *Creator) NewImageFromGoImage(goimg _ba.Image) (*Image, error) { return _ababa(goimg) }

// GeneratePageBlocks draws the rectangle on a new block representing the page.
func (_fegg *Ellipse) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_age := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_dege := _cec.Circle{X: _fegg._agdd - _fegg._acea/2, Y: ctx.PageHeight - _fegg._ddgg - _fegg._gdeg/2, Width: _fegg._acea, Height: _fegg._gdeg, Opacity: 1.0, BorderWidth: _fegg._fbec}
	if _fegg._cfagc == PositionRelative {
		_dege.X = ctx.X
		_dege.Y = ctx.PageHeight - ctx.Y - _fegg._gdeg
	}
	if _fegg._cedc != nil {
		_dege.FillEnabled = true
		_dege.FillColor = _bcfe(_fegg._cedc)
	}
	if _fegg._agcc != nil {
		_dege.BorderEnabled = false
		if _fegg._fbec > 0 {
			_dege.BorderEnabled = true
		}
		_dege.BorderColor = _bcfe(_fegg._agcc)
		_dege.BorderWidth = _fegg._fbec
	}
	_faf, _, _cgd := _dege.Draw("")
	if _cgd != nil {
		return nil, ctx, _cgd
	}
	_cgd = _age.addContentsByString(string(_faf))
	if _cgd != nil {
		return nil, ctx, _cgd
	}
	return []*Block{_age}, ctx, nil
}

// SkipRows skips over a specified number of rows in the table.
func (_ebbga *Table) SkipRows(num int) {
	_eeff := num*_ebbga._cfee - 1
	if _eeff < 0 {
		_df.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0073\u006b\u0069\u0070\u0020b\u0061\u0063\u006b\u0020\u0074\u006f\u0020\u0070\u0072\u0065\u0076\u0069\u006f\u0075\u0073\u0020\u0063\u0065\u006c\u006c\u0073")
		return
	}
	_ebbga._edag += _eeff
}
func (_gdfc *templateProcessor) parseChapterHeading(_afdgc *templateNode) (interface{}, error) {
	if _afdgc._dfdd == nil {
		_df.Log.Error("\u0043\u0068a\u0070\u0074\u0065\u0072 \u0068\u0065a\u0064\u0069\u006e\u0067\u0020\u0070\u0061\u0072e\u006e\u0074\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c\u002e")
		return nil, _dbea
	}
	_bbaca, _ceca := _afdgc._dfdd._cedba.(*Chapter)
	if !_ceca {
		_df.Log.Error("\u0043h\u0061\u0070t\u0065\u0072\u0020h\u0065\u0061\u0064\u0069\u006e\u0067\u0020p\u0061\u0072\u0065\u006e\u0074\u0020(\u0025\u0054\u0029\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020a\u0020\u0063\u0068\u0061\u0070\u0074\u0065\u0072\u002e", _afdgc._dfdd._cedba)
		return nil, _dbea
	}
	_dfbdb := _bbaca.GetHeading()
	if _, _fcgee := _gdfc.parseParagraph(_afdgc, _dfbdb); _fcgee != nil {
		return nil, _fcgee
	}
	return _dfbdb, nil
}

// GetMargins returns the margins of the line: left, right, top, bottom.
func (_caca *Line) GetMargins() (float64, float64, float64, float64) {
	return _caca._dacfg.Left, _caca._dacfg.Right, _caca._dacfg.Top, _caca._dacfg.Bottom
}

const (
	TextAlignmentLeft TextAlignment = iota
	TextAlignmentRight
	TextAlignmentCenter
	TextAlignmentJustify
)

// NewEllipse creates a new ellipse centered at (xc,yc) with a width and height specified.
func (_gcfg *Creator) NewEllipse(xc, yc, width, height float64) *Ellipse {
	return _bbaf(xc, yc, width, height)
}

// NewCellProps returns the default properties of an invoice cell.
func (_bcaa *Invoice) NewCellProps() InvoiceCellProps {
	_ceefg := ColorRGBFrom8bit(255, 255, 255)
	return InvoiceCellProps{TextStyle: _bcaa._ffdf, Alignment: CellHorizontalAlignmentLeft, BackgroundColor: _ceefg, BorderColor: _ceefg, BorderWidth: 1, BorderSides: []CellBorderSide{CellBorderSideAll}}
}

type templateTag struct {
	_cedge map[string]struct{}
	_bgff  func(*templateProcessor, *templateNode) (interface{}, error)
}

// NewCurve returns new instance of Curve between points (x1,y1) and (x2, y2) with control point (cx,cy).
func (_fdc *Creator) NewCurve(x1, y1, cx, cy, x2, y2 float64) *Curve {
	return _fcef(x1, y1, cx, cy, x2, y2)
}

var PPI float64 = 72

// SetTotal sets the total of the invoice.
func (_ccaa *Invoice) SetTotal(value string) { _ccaa._defg[1].Value = value }
func (_aada *Image) applyFitMode(_babb float64) {
	_babb -= _aada._dbfe.Left + _aada._dbfe.Right
	switch _aada._aedff {
	case FitModeFillWidth:
		_aada.ScaleToWidth(_babb)
	}
}

// GeneratePageBlocks generates the page blocks for the Division component.
// Multiple blocks are generated if the contents wrap over multiple pages.
func (_agcfa *Division) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	var (
		_eaeb []*Block
		_afca bool
		_fgdd error
		_gfcc = _agcfa._edfc.IsRelative()
		_geag = _agcfa._fccf.Top
	)
	if _gfcc && !_agcfa._dcd && !_agcfa._ddda {
		_dbffb := _agcfa.ctxHeight(ctx.Width)
		if _dbffb > ctx.Height-_agcfa._fccf.Top && _dbffb <= ctx.PageHeight-ctx.Margins.Top-ctx.Margins.Bottom {
			if _eaeb, ctx, _fgdd = _acdfe().GeneratePageBlocks(ctx); _fgdd != nil {
				return nil, ctx, _fgdd
			}
			_afca = true
			_geag = 0
		}
	}
	_gfeg := ctx
	_cgea := ctx
	if _gfcc {
		ctx.X += _agcfa._fccf.Left
		ctx.Y += _geag
		ctx.Width -= _agcfa._fccf.Left + _agcfa._fccf.Right
		ctx.Height -= _geag
		_cgea = ctx
		ctx.X += _agcfa._gfeb.Left
		ctx.Y += _agcfa._gfeb.Top
		ctx.Width -= _agcfa._gfeb.Left + _agcfa._gfeb.Right
		ctx.Height -= _agcfa._gfeb.Top
		ctx.Margins.Top += _agcfa._gfeb.Top
		ctx.Margins.Bottom += _agcfa._gfeb.Bottom
		ctx.Margins.Left += _agcfa._fccf.Left + _agcfa._gfeb.Left
		ctx.Margins.Right += _agcfa._fccf.Right + _agcfa._gfeb.Right
	}
	ctx.Inline = _agcfa._ddda
	_egcg := ctx
	_ffbg := ctx
	var _gae float64
	for _, _ggca := range _agcfa._aefd {
		if ctx.Inline {
			if (ctx.X-_egcg.X)+_ggca.Width() <= ctx.Width {
				ctx.Y = _ffbg.Y
				ctx.Height = _ffbg.Height
			} else {
				ctx.X = _egcg.X
				ctx.Width = _egcg.Width
				_ffbg.Y += _gae
				_ffbg.Height -= _gae
				_gae = 0
			}
		}
		_fcec, _acgb, _efbf := _ggca.GeneratePageBlocks(ctx)
		if _efbf != nil {
			_df.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006eg\u0020p\u0061\u0067\u0065\u0020\u0062\u006c\u006f\u0063\u006b\u0073\u003a\u0020\u0025\u0076", _efbf)
			return nil, ctx, _efbf
		}
		if len(_fcec) < 1 {
			continue
		}
		if len(_eaeb) > 0 {
			_eaeb[len(_eaeb)-1].mergeBlocks(_fcec[0])
			_eaeb = append(_eaeb, _fcec[1:]...)
		} else {
			if _dcb := _fcec[0]._ga; _dcb == nil || len(*_dcb) == 0 {
				_afca = true
			}
			_eaeb = append(_eaeb, _fcec[0:]...)
		}
		if ctx.Inline {
			if ctx.Page != _acgb.Page {
				_egcg.Y = ctx.Margins.Top
				_egcg.Height = ctx.PageHeight - ctx.Margins.Top
				_ffbg.Y = _egcg.Y
				_ffbg.Height = _egcg.Height
				_gae = _acgb.Height - _egcg.Height
			} else {
				if _gdf := ctx.Height - _acgb.Height; _gdf > _gae {
					_gae = _gdf
				}
			}
		} else {
			_acgb.X = ctx.X
		}
		ctx = _acgb
	}
	ctx.Inline = _gfeg.Inline
	ctx.Margins = _gfeg.Margins
	if _gfcc {
		ctx.X = _gfeg.X
		ctx.Width = _gfeg.Width
		ctx.Y += _agcfa._gfeb.Bottom
		ctx.Height -= _agcfa._gfeb.Bottom
	}
	if _agcfa._bgeg != nil {
		_eaeb, _fgdd = _agcfa.drawBackground(_eaeb, _cgea, ctx, _afca)
		if _fgdd != nil {
			return nil, ctx, _fgdd
		}
	}
	if _agcfa._edfc.IsAbsolute() {
		return _eaeb, _gfeg, nil
	}
	ctx.Y += _agcfa._fccf.Bottom
	ctx.Height -= _agcfa._fccf.Bottom
	return _eaeb, ctx, nil
}
func _adb(_bee _ab.ChartRenderable) *Chart {
	return &Chart{_fbaa: _bee, _ebe: PositionRelative, _daea: Margins{Top: 10, Bottom: 10}}
}

var _accge = map[string]*templateTag{"\u0070a\u0072\u0061\u0067\u0072\u0061\u0070h": {_cedge: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bgff: _bcafd}, "\u0074\u0065\u0078\u0074\u002d\u0063\u0068\u0075\u006e\u006b": {_cedge: map[string]struct{}{"\u0070a\u0072\u0061\u0067\u0072\u0061\u0070h": {}}, _bgff: _geaf}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {_cedge: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bgff: _dadae}, "\u0074\u0061\u0062l\u0065": {_cedge: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bgff: _edcgg}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {_cedge: map[string]struct{}{"\u0074\u0061\u0062l\u0065": {}}, _bgff: _baafd}, "\u006c\u0069\u006e\u0065": {_cedge: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bgff: _gaef}, "\u0069\u006d\u0061g\u0065": {_cedge: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bgff: _gdbg}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {_cedge: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bgff: _febe}, "\u0063h\u0061p\u0074\u0065\u0072\u002d\u0068\u0065\u0061\u0064\u0069\u006e\u0067": {_cedge: map[string]struct{}{"\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bgff: _gdaeb}, "\u0063\u0068\u0061r\u0074": {_cedge: map[string]struct{}{"\u0063r\u0065\u0061\u0074\u006f\u0072": {}, "\u0062\u006c\u006fc\u006b": {}, "\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}, "\u0074\u0061\u0062\u006c\u0065\u002d\u0063\u0065\u006c\u006c": {}, "\u0063h\u0061\u0070\u0074\u0065\u0072": {}}, _bgff: _ebeca}, "\u0062\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064": {_cedge: map[string]struct{}{"\u0064\u0069\u0076\u0069\u0073\u0069\u006f\u006e": {}}, _bgff: _abbd}}

func (_cdgdb *templateProcessor) parseMarginAttr(_bdce, _cbbb string) Margins {
	_df.Log.Debug("\u0050\u0061r\u0073\u0069\u006e\u0067 \u006d\u0061r\u0067\u0069\u006e\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c \u0025\u0073\u0029\u002e", _bdce, _cbbb)
	_acff := Margins{}
	switch _ecaaf := _ed.Fields(_cbbb); len(_ecaaf) {
	case 1:
		_acff.Top, _ = _ae.ParseFloat(_ecaaf[0], 64)
		_acff.Bottom = _acff.Top
		_acff.Left = _acff.Top
		_acff.Right = _acff.Top
	case 2:
		_acff.Top, _ = _ae.ParseFloat(_ecaaf[0], 64)
		_acff.Bottom = _acff.Top
		_acff.Left, _ = _ae.ParseFloat(_ecaaf[1], 64)
		_acff.Right = _acff.Left
	case 3:
		_acff.Top, _ = _ae.ParseFloat(_ecaaf[0], 64)
		_acff.Left, _ = _ae.ParseFloat(_ecaaf[1], 64)
		_acff.Right = _acff.Left
		_acff.Bottom, _ = _ae.ParseFloat(_ecaaf[2], 64)
	case 4:
		_acff.Top, _ = _ae.ParseFloat(_ecaaf[0], 64)
		_acff.Right, _ = _ae.ParseFloat(_ecaaf[1], 64)
		_acff.Bottom, _ = _ae.ParseFloat(_ecaaf[2], 64)
		_acff.Left, _ = _ae.ParseFloat(_ecaaf[3], 64)
	}
	return _acff
}

// Height returns Image's document height.
func (_fgdce *Image) Height() float64 { return _fgdce._gddc }

// GetMargins returns the Block's margins: left, right, top, bottom.
func (_cb *Block) GetMargins() (float64, float64, float64, float64) {
	return _cb._ac.Left, _cb._ac.Right, _cb._ac.Top, _cb._ac.Bottom
}
func (_fdcg *Invoice) generateNoteBlocks(_ddabf DrawContext) ([]*Block, DrawContext, error) {
	_edfe := _cdde()
	_aeeg := append([][2]string{_fdcg._dbe, _fdcg._dcdf}, _fdcg._bgfee...)
	for _, _abc := range _aeeg {
		if _abc[1] != "" {
			_caeg := _fdcg.drawSection(_abc[0], _abc[1])
			for _, _egag := range _caeg {
				_edfe.Add(_egag)
			}
			_gafb := _ggaa(_fdcg._ffdf)
			_gafb.SetMargins(0, 0, 10, 0)
			_edfe.Add(_gafb)
		}
	}
	return _edfe.GeneratePageBlocks(_ddabf)
}

// TextOverflow determines the behavior of paragraph text which does
// not fit in the available space.
type TextOverflow int

// SetLineSeparator sets the separator for all new lines of the table of contents.
func (_fcgg *TOC) SetLineSeparator(separator string) { _fcgg._ggdcba = separator }

// NewCell returns a new invoice table cell.
func (_eadc *Invoice) NewCell(value string) *InvoiceCell {
	return _eadc.newCell(value, _eadc.NewCellProps())
}
func (_eabac *Table) wrapContent(_bbff DrawContext) error {
	if _eabac._ggbfc {
		return nil
	}
	_facbf := func(_afdg *TableCell, _ccae int, _beab int, _beaf int) (_bfaeg int) {
		if _beaf < 1 {
			return -1
		}
		_bdcd := 0
		for _gfdff := _beab + 1; _gfdff < len(_eabac._ebfb)-1; _gfdff++ {
			_bdef := _eabac._ebfb[_gfdff]
			if _bdef._gbdd == _beaf {
				_bdcd = _gfdff
				if (_bdef._dbdf < _afdg._dbdf && _eabac._cfee > _bdef._dbdf) || _afdg._dbdf < _eabac._cfee {
					continue
				}
				break
			}
		}
		_ddeb := float64(0.0)
		for _gbad := 0; _gbad < _afdg._fefef; _gbad++ {
			_ddeb += _eabac._ceeb[_afdg._gbdd+_gbad-1]
		}
		_effga := float64(0.0)
		for _cfgef := 0; _cfgef < _afdg._cegef; _cfgef++ {
			_effga += _eabac._agdgd[_afdg._dbdf+_cfgef-1]
		}
		var (
			_afdf VectorDrawable
			_adge = false
		)
		switch _cbffc := _afdg._cbgcb.(type) {
		case *StyledParagraph:
			_fbeaf := _bbff
			_fbeaf.Height = _a.Floor(_ddeb - _cbffc._afdba.Top - _cbffc._afdba.Bottom - 0.5*_cbffc.getTextHeight())
			_fbeaf.Width = _effga
			_egeb, _aadbc, _eccbg := _cbffc.split(_fbeaf)
			if _eccbg != nil {
				_df.Log.Error("\u0045\u0072\u0072o\u0072\u0020\u0077\u0072a\u0070\u0020\u0073\u0074\u0079\u006c\u0065d\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u003a\u0020\u0025\u0076", _eccbg.Error())
			}
			if _egeb != nil && _aadbc != nil {
				_eabac._ebfb[_beab]._cbgcb = _egeb
				_afdf = _aadbc
				_adge = true
			}
		}
		_eabac._ebfb[_beab]._fefef = _afdg._fefef
		_bbff.Height = _bbff.PageHeight - _bbff.Margins.Top - _bbff.Margins.Bottom
		_ffbe := _afdg.cloneProps(nil)
		if _adge {
			_ffbe._cbgcb = _afdf
		}
		_ffbe._fefef = _ccae - 1
		_ffbe._gbdd = _beaf + 1
		_ffbe._dbdf = _afdg._dbdf
		_eabac._ebfb = append(_eabac._ebfb[:_bdcd+1], append([]*TableCell{_ffbe}, _eabac._ebfb[_bdcd+1:]...)...)
		return _bdcd + 1
	}
	_cfaga := float64(0.0)
	_afcbd := 0
	_bdgdg := -1
	for _bgfgc, _fdca := range _eabac._ebfb {
		if _bdgdg == _bgfgc {
			_afcbd = _fdca._gbdd
			_cfaga = 0.0
		}
		if _fdca._fefef < 2 {
			if _afcbd < _fdca._gbdd && _bgfgc > _bdgdg && _cfaga < _bbff.Height {
				_cfaga += _eabac._ceeb[_fdca._gbdd-1]
			}
			_afcbd = _fdca._gbdd
			continue
		}
		if _cfaga < 1 && _bgfgc == _bdgdg {
			_cfaga += _eabac._ceeb[_fdca._gbdd-1]
		}
		_gfbad := float64(0.0)
		_bfbg := -1
		_bcacf := -1
		_cggbb := 0
		for _ggccea := 0; _ggccea < _fdca._fefef; _ggccea++ {
			if (_gfbad + _eabac._ceeb[_fdca._gbdd+_ggccea-1]) > (_bbff.Height - _cfaga) {
				break
			}
			_gfbad += _eabac._ceeb[_fdca._gbdd+_ggccea-1]
			_bcacf = _fdca._gbdd + _ggccea
			_bfbg = _fdca._fefef - _ggccea
			_cggbb++
		}
		if _bfbg > 0 && _fdca._fefef > _cggbb {
			_fdca._fefef = _cggbb
			_bdgdg = _facbf(_fdca, _bfbg, _bgfgc, _bcacf)
			_afcbd = _bcacf
		}
	}
	return nil
}
func _bbaf(_eeeba, _fbdf, _fdce, _abgf float64) *Ellipse {
	_gggd := &Ellipse{}
	_gggd._agdd = _eeeba
	_gggd._ddgg = _fbdf
	_gggd._acea = _fdce
	_gggd._gdeg = _abgf
	_gggd._cfagc = PositionAbsolute
	_gggd._agcc = ColorBlack
	_gggd._fbec = 1.0
	return _gggd
}
func _fffeg(_cdad string) (*Image, error) {
	_cba, _adf := _e.Open(_cdad)
	if _adf != nil {
		return nil, _adf
	}
	defer _cba.Close()
	_adgf, _adf := _f.ImageHandling.Read(_cba)
	if _adf != nil {
		_df.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _adf)
		return nil, _adf
	}
	return _baec(_adgf)
}

// GeneratePageBlocks generates the page blocks.  Multiple blocks are generated if the contents wrap
// over multiple pages. Implements the Drawable interface.
func (_cccb *Paragraph) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_egcee := ctx
	var _afcc []*Block
	_abfdd := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _cccb._beeba.IsRelative() {
		ctx.X += _cccb._eacf.Left
		ctx.Y += _cccb._eacf.Top
		ctx.Width -= _cccb._eacf.Left + _cccb._eacf.Right
		ctx.Height -= _cccb._eacf.Top
		_cccb.SetWidth(ctx.Width)
		if _cccb.Height() > ctx.Height {
			_afcc = append(_afcc, _abfdd)
			_abfdd = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_cfec := ctx
			_cfec.Y = ctx.Margins.Top
			_cfec.X = ctx.Margins.Left + _cccb._eacf.Left
			_cfec.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
			_cfec.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _cccb._eacf.Left - _cccb._eacf.Right
			ctx = _cfec
		}
	} else {
		if int(_cccb._bdga) <= 0 {
			_cccb.SetWidth(_cccb.getTextWidth())
		}
		ctx.X = _cccb._fbbde
		ctx.Y = _cccb._agdg
	}
	ctx, _dfce := _fdcb(_abfdd, _cccb, ctx)
	if _dfce != nil {
		_df.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dfce)
		return nil, ctx, _dfce
	}
	_afcc = append(_afcc, _abfdd)
	if _cccb._beeba.IsRelative() {
		ctx.Y += _cccb._eacf.Bottom
		ctx.Height -= _cccb._eacf.Bottom
		if !ctx.Inline {
			ctx.X = _egcee.X
			ctx.Width = _egcee.Width
		}
		return _afcc, ctx, nil
	}
	return _afcc, _egcee, nil
}
func (_abbfg *templateProcessor) parseTextVerticalAlignmentAttr(_bffgc, _dabef string) TextVerticalAlignment {
	_df.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0074\u0065\u0078\u0074\u0020\u0076\u0065r\u0074\u0069\u0063\u0061\u006c\u0020\u0061\u006c\u0069\u0067\u006e\u006d\u0065n\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a (\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _bffgc, _dabef)
	_gffca := map[string]TextVerticalAlignment{"\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065": TextVerticalAlignmentBaseline, "\u0063\u0065\u006e\u0074\u0065\u0072": TextVerticalAlignmentCenter}[_dabef]
	return _gffca
}

// GetHeading returns the chapter heading paragraph. Used to give access to address style: font, sizing etc.
func (_begd *Chapter) GetHeading() *Paragraph { return _begd._dbf }

// UnsupportedRuneError is an error that occurs when there is unsupported glyph being used.
type UnsupportedRuneError struct {
	Message string
	Rune    rune
}

// GeneratePageBlocks draws the filled curve on page blocks.
func (_fece *FilledCurve) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_gfdbf := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_aaaf, _, _bfff := _fece.draw("")
	if _bfff != nil {
		return nil, ctx, _bfff
	}
	_bfff = _gfdbf.addContentsByString(string(_aaaf))
	if _bfff != nil {
		return nil, ctx, _bfff
	}
	return []*Block{_gfdbf}, ctx, nil
}

// Curve represents a cubic Bezier curve with a control point.
type Curve struct {
	_ggce  float64
	_bdcce float64
	_dafb  float64
	_feec  float64
	_fgaeb float64
	_aaeb  float64
	_cddgg Color
	_bbf   float64
}

func _ccgf(_befge *_f.PdfAnnotation) *_f.PdfAnnotation {
	if _befge == nil {
		return nil
	}
	var _dfae *_f.PdfAnnotation
	switch _bfefd := _befge.GetContext().(type) {
	case *_f.PdfAnnotationLink:
		if _dfea := _fcfea(_bfefd); _dfea != nil {
			_dfae = _dfea.PdfAnnotation
		}
	}
	return _dfae
}

// Polyline represents a slice of points that are connected as straight lines.
// Implements the Drawable interface and can be rendered using the Creator.
type Polyline struct {
	_bcg   *_cec.Polyline
	_afbcd float64
}

// SetInline sets the inline mode of the division.
func (_bdgd *Division) SetInline(inline bool) { _bdgd._ddda = inline }
func _dcge(_dcfg VectorDrawable, _bbadg float64) float64 {
	switch _eaba := _dcfg.(type) {
	case *Paragraph:
		if _eaba._bbfc {
			_eaba.SetWidth(_bbadg - _eaba._eacf.Left - _eaba._eacf.Right)
		}
		return _eaba.Height() + _eaba._eacf.Top + _eaba._eacf.Bottom
	case *StyledParagraph:
		if _eaba._ecbf {
			_eaba.SetWidth(_bbadg - _eaba._afdba.Left - _eaba._afdba.Right)
		}
		return _eaba.Height() + _eaba._afdba.Top + _eaba._afdba.Bottom
	case *Image:
		_eaba.applyFitMode(_bbadg)
		return _eaba.Height() + _eaba._dbfe.Top + _eaba._dbfe.Bottom
	case marginDrawable:
		_, _, _fbbg, _fagbg := _eaba.GetMargins()
		return _eaba.Height() + _fbbg + _fagbg
	default:
		return _eaba.Height()
	}
}
func _ababa(_acedg _ba.Image) (*Image, error) {
	_bdgb, _gffc := _f.ImageHandling.NewImageFromGoImage(_acedg)
	if _gffc != nil {
		return nil, _gffc
	}
	return _baec(_bdgb)
}

const (
	TextVerticalAlignmentBaseline TextVerticalAlignment = iota
	TextVerticalAlignmentCenter
)

// NewTOC creates a new table of contents.
func (_ccag *Creator) NewTOC(title string) *TOC {
	_feba := _ccag.NewTextStyle()
	_feba.Font = _ccag._beea
	return _egcac(title, _ccag.NewTextStyle(), _feba)
}

// SetColorRight sets border color for right.
func (_cdg *border) SetColorRight(col Color) { _cdg._bfae = col }
func _abbd(_cfedd *templateProcessor, _gfgf *templateNode) (interface{}, error) {
	return _cfedd.parseBackground(_gfgf)
}

// SetLineHeight sets the line height (1.0 default).
func (_fgag *Paragraph) SetLineHeight(lineheight float64) { _fgag._abbg = lineheight }

// SetPos sets the absolute position. Changes object positioning to absolute.
func (_adag *Image) SetPos(x, y float64) {
	_adag._ccec = PositionAbsolute
	_adag._cgee = x
	_adag._ddab = y
}

type listItem struct {
	_cdef VectorDrawable
	_gade TextChunk
}

// SetMargins sets the Block's left, right, top, bottom, margins.
func (_bce *Block) SetMargins(left, right, top, bottom float64) {
	_bce._ac.Left = left
	_bce._ac.Right = right
	_bce._ac.Top = top
	_bce._ac.Bottom = bottom
}
func (_gdfg *templateProcessor) parseTextAlignmentAttr(_gaeef, _cceg string) TextAlignment {
	_df.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020t\u0065\u0078\u0074\u0020\u0061\u006c\u0069g\u006e\u006d\u0065\u006e\u0074\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028`\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _gaeef, _cceg)
	_abefa := map[string]TextAlignment{"\u006c\u0065\u0066\u0074": TextAlignmentLeft, "\u0072\u0069\u0067h\u0074": TextAlignmentRight, "\u0063\u0065\u006e\u0074\u0065\u0072": TextAlignmentCenter, "\u006au\u0073\u0074\u0069\u0066\u0079": TextAlignmentJustify}[_cceg]
	return _abefa
}
func (_aaef *StyledParagraph) wrapChunks(_bde bool) error {
	if !_aaef._ecbf || int(_aaef._bbef) <= 0 {
		_aaef._bfaa = [][]*TextChunk{_aaef._eedad}
		return nil
	}
	if _aaef._bggd {
		_aaef.wrapWordChunks()
	}
	_aaef._bfaa = [][]*TextChunk{}
	var _bfad []*TextChunk
	var _dfceg float64
	_ceag := _ee.IsSpace
	if !_bde {
		_ceag = func(rune) bool { return false }
	}
	_febfe := _ffab(_aaef._bbef*1000.0, 0.000001)
	for _, _acfca := range _aaef._eedad {
		_cgef := _acfca.Style
		_dgbad := _acfca._ceddd
		var (
			_agbdg []rune
			_gdcgb []float64
		)
		_fbef := _gaafg(_acfca.Text)
		for _, _dffed := range _acfca.Text {
			if _dffed == '\u000A' {
				if !_bde {
					_agbdg = append(_agbdg, _dffed)
				}
				_bfad = append(_bfad, &TextChunk{Text: _ed.TrimRightFunc(string(_agbdg), _ceag), Style: _cgef, _ceddd: _ccgf(_dgbad)})
				_aaef._bfaa = append(_aaef._bfaa, _bfad)
				_bfad = nil
				_dfceg = 0
				_agbdg = nil
				_gdcgb = nil
				continue
			}
			_ddcde := _dffed == ' '
			_abge, _ccad := _cgef.Font.GetRuneMetrics(_dffed)
			if !_ccad {
				_df.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _dffed)
				return _d.New("\u0067\u006c\u0079\u0070\u0068\u0020\u0063\u0068\u0061\u0072\u0020m\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006d\u0069\u0073s\u0069\u006e\u0067")
			}
			_dabd := _cgef.FontSize * _abge.Wx * _cgef.horizontalScale()
			_gbggc := _dabd
			if !_ddcde {
				_gbggc = _dabd + _cgef.CharSpacing*1000.0
			}
			if _dfceg+_dabd > _febfe {
				if _aaef._bggd {
					if len(_bfad) > 0 {
						_aaef._bfaa = append(_aaef._bfaa, _bfad)
						_bfad = []*TextChunk{}
					}
					_agbdg = append(_agbdg, _dffed)
					_gdcgb = append(_gdcgb, _gbggc)
					_gafd := -1
					if !_ddcde {
						for _cgdf := len(_agbdg) - 1; _cgdf >= 0; _cgdf-- {
							if _agbdg[_cgdf] == ' ' {
								_gafd = _cgdf
								break
							}
						}
					}
					if _gafd >= 0 {
						_agbdg = _agbdg[_gafd+1:]
						_gdcgb = _gdcgb[_gafd+1:]
					}
					_dfceg = 0
					for _, _ecgd := range _gdcgb {
						_dfceg += _ecgd
					}
					continue
				}
				_bgfgg := -1
				if !_ddcde {
					for _ecdcd := len(_agbdg) - 1; _ecdcd >= 0; _ecdcd-- {
						if _agbdg[_ecdcd] == ' ' {
							_bgfgg = _ecdcd
							break
						}
					}
				}
				_eceag := string(_agbdg)
				if _bgfgg >= 0 {
					_eceag = string(_agbdg[0 : _bgfgg+1])
					_agbdg = _agbdg[_bgfgg+1:]
					_agbdg = append(_agbdg, _dffed)
					_gdcgb = _gdcgb[_bgfgg+1:]
					_gdcgb = append(_gdcgb, _gbggc)
					_dfceg = 0
					for _, _eeae := range _gdcgb {
						_dfceg += _eeae
					}
				} else {
					if _ddcde {
						_dfceg = 0
						_agbdg = []rune{}
						_gdcgb = []float64{}
					} else {
						_dfceg = _gbggc
						_agbdg = []rune{_dffed}
						_gdcgb = []float64{_gbggc}
					}
				}
				_eceag = _daed(_eceag, _fbef)
				if !_bde && _ddcde {
					_eceag += "\u0020"
				}
				_bfad = append(_bfad, &TextChunk{Text: _ed.TrimRightFunc(_eceag, _ceag), Style: _cgef, _ceddd: _ccgf(_dgbad)})
				_aaef._bfaa = append(_aaef._bfaa, _bfad)
				_bfad = []*TextChunk{}
			} else {
				_dfceg += _gbggc
				_agbdg = append(_agbdg, _dffed)
				_gdcgb = append(_gdcgb, _gbggc)
			}
		}
		if len(_agbdg) > 0 {
			_cafd := _daed(string(_agbdg), _fbef)
			_bfad = append(_bfad, &TextChunk{Text: _cafd, Style: _cgef, _ceddd: _ccgf(_dgbad)})
		}
	}
	if len(_bfad) > 0 {
		_aaef._bfaa = append(_aaef._bfaa, _bfad)
	}
	return nil
}

const (
	FitModeNone FitMode = iota
	FitModeFillWidth
)

// Height returns the total height of all rows.
func (_gefg *Table) Height() float64 {
	_caccg := float64(0.0)
	for _, _cagga := range _gefg._ceeb {
		_caccg += _cagga
	}
	return _caccg
}

// SetFitMode sets the fit mode of the line.
// NOTE: The fit mode is only applied if relative positioning is used.
func (_fgfec *Line) SetFitMode(fitMode FitMode) { _fgfec._gggcb = fitMode }

// Reset removes all the text chunks the paragraph contains.
func (_egacb *StyledParagraph) Reset() { _egacb._eedad = []*TextChunk{} }

// Polygon represents a polygon shape.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Polygon struct {
	_ccbeg *_cec.Polygon
	_fead  float64
	_afab  float64
}

// TemplateOptions contains options and resources to use when rendering
// a template with a Creator instance.
// All the resources in the map fields can be referenced by their
// name/key in the template which is rendered using the options instance.
type TemplateOptions struct {

	// Helper functions map.
	HelperFuncMap _g.FuncMap

	// Named resource maps.
	FontMap  map[string]*_f.PdfFont
	ImageMap map[string]*_f.Image
	ColorMap map[string]Color
	ChartMap map[string]_ab.ChartRenderable
}

// Height returns Rectangle's document height.
func (_befbc *Rectangle) Height() float64 { return _befbc._ecca }
func (_aabg *templateProcessor) parseFloatAttr(_cggeg, _fagdb string) float64 {
	_df.Log.Debug("\u0050\u0061rs\u0069\u006e\u0067 \u0066\u006c\u006f\u0061t a\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020(`\u0025\u0073\u0060\u002c\u0020\u0025\u0073)\u002e", _cggeg, _fagdb)
	_cadg, _ := _ae.ParseFloat(_fagdb, 64)
	return _cadg
}
func (_gabgg *templateProcessor) parseFloatArray(_gbefe, _dgaca string) []float64 {
	_df.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020\u0066\u006c\u006f\u0061\u0074\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060%\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _gbefe, _dgaca)
	_caccb := _ed.Fields(_dgaca)
	_gcec := make([]float64, 0, len(_caccb))
	for _, _gddccd := range _caccb {
		_gddb, _ := _ae.ParseFloat(_gddccd, 64)
		_gcec = append(_gcec, _gddb)
	}
	return _gcec
}
func (_edb rgbColor) ToRGB() (float64, float64, float64) { return _edb._fabb, _edb._abab, _edb._gggf }

// New creates a new instance of the PDF Creator.
func New() *Creator {
	_ebbfc := &Creator{}
	_ebbfc._gadb = []*_f.PdfPage{}
	_ebbfc._caec = map[*_f.PdfPage]*Block{}
	_ebbfc._gba = map[*_f.PdfPage]*pageTransformations{}
	_ebbfc.SetPageSize(PageSizeLetter)
	_bbea := 0.1 * _ebbfc._gee
	_ebbfc._fade.Left = _bbea
	_ebbfc._fade.Right = _bbea
	_ebbfc._fade.Top = _bbea
	_ebbfc._fade.Bottom = _bbea
	var _gbed error
	_ebbfc._ffgd, _gbed = _f.NewStandard14Font(_f.HelveticaName)
	if _gbed != nil {
		_ebbfc._ffgd = _f.DefaultFont()
	}
	_ebbfc._beea, _gbed = _f.NewStandard14Font(_f.HelveticaBoldName)
	if _gbed != nil {
		_ebbfc._ffgd = _f.DefaultFont()
	}
	_ebbfc._cddg = _ebbfc.NewTOC("\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0065\u006e\u0074\u0073")
	_ebbfc.AddOutlines = true
	_ebbfc._dfd = _f.NewOutline()
	return _ebbfc
}

// NewInvoice returns an instance of an empty invoice.
func (_aggc *Creator) NewInvoice() *Invoice {
	_eaea := _aggc.NewTextStyle()
	_eaea.Font = _aggc._beea
	return _fabgd(_aggc.NewTextStyle(), _eaea)
}

const (
	HorizontalAlignmentLeft HorizontalAlignment = iota
	HorizontalAlignmentCenter
	HorizontalAlignmentRight
)

// SetPos sets the Table's positioning to absolute mode and specifies the upper-left corner
// coordinates as (x,y).
// Note that this is only sensible to use when the table does not wrap over multiple pages.
// TODO: Should be able to set width too (not just based on context/relative positioning mode).
func (_ggcce *Table) SetPos(x, y float64) {
	_ggcce._bfed = PositionAbsolute
	_ggcce._fdcbfa = x
	_ggcce._eebda = y
}

// Total returns the invoice total description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_fgfg *Invoice) Total() (*InvoiceCell, *InvoiceCell) { return _fgfg._defg[0], _fgfg._defg[1] }

// SetLineStyle sets the style for all the line components: number, title,
// separator, page. The style is applied only for new lines added to the
// TOC component.
func (_fgfad *TOC) SetLineStyle(style TextStyle) {
	_fgfad.SetLineNumberStyle(style)
	_fgfad.SetLineTitleStyle(style)
	_fgfad.SetLineSeparatorStyle(style)
	_fgfad.SetLinePageStyle(style)
}

// SetFillOpacity sets the fill opacity.
func (_bgeba *Rectangle) SetFillOpacity(opacity float64) { _bgeba._ffcbg = opacity }
func (_gfdcd *TableCell) width(_ecbff []float64, _bbgf float64) float64 {
	_gbfc := float64(0.0)
	for _eecec := 0; _eecec < _gfdcd._cegef; _eecec++ {
		_gbfc += _ecbff[_gfdcd._dbdf+_eecec-1]
	}
	return _gbfc * _bbgf
}

var PPMM = float64(72 * 1.0 / 25.4)

// SetColumns overwrites any columns in the line items table. This should be
// called before AddLine.
func (_ddfc *Invoice) SetColumns(cols []*InvoiceCell) { _ddfc._afeb = cols }

// InfoLines returns all the rows in the invoice information table as
// description-value cell pairs.
func (_beeg *Invoice) InfoLines() [][2]*InvoiceCell {
	_ebbfa := [][2]*InvoiceCell{_beeg._edeb, _beeg._fgcge, _beeg._debe}
	return append(_ebbfa, _beeg._agef...)
}

// CellHorizontalAlignment defines the table cell's horizontal alignment.
type CellHorizontalAlignment int

func (_dfag *templateProcessor) parseChapter(_cgccg *templateNode) (interface{}, error) {
	_gcfe := _dfag.creator.NewChapter
	if _cgccg._dfdd != nil {
		if _ddedg, _egdda := _cgccg._dfdd._cedba.(*Chapter); _egdda {
			_gcfe = _ddedg.NewSubchapter
		}
	}
	_ffde := _gcfe("")
	for _, _deba := range _cgccg._ddgde.Attr {
		_egfgc := _deba.Value
		switch _aggf := _deba.Name.Local; _aggf {
		case "\u0073\u0068\u006f\u0077\u002d\u006e\u0075\u006d\u0062e\u0072\u0069\u006e\u0067":
			_ffde.SetShowNumbering(_dfag.parseBoolAttr(_aggf, _egfgc))
		case "\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u002d\u0069n\u002d\u0074\u006f\u0063":
			_ffde.SetIncludeInTOC(_dfag.parseBoolAttr(_aggf, _egfgc))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_fede := _dfag.parseMarginAttr(_aggf, _egfgc)
			_ffde.SetMargins(_fede.Left, _fede.Right, _fede.Top, _fede.Bottom)
		default:
			_df.Log.Debug("\u0055\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0068\u0061\u0070\u0074\u0065\u0072\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _aggf)
		}
	}
	return _ffde, nil
}
func _fcge(_fbea []_cec.Point) *Polyline {
	return &Polyline{_bcg: &_cec.Polyline{Points: _fbea, LineColor: _f.NewPdfColorDeviceRGB(0, 0, 0), LineWidth: 1.0}, _afbcd: 1.0}
}

// SetLineMargins sets the margins for all new lines of the table of contents.
func (_fabc *TOC) SetLineMargins(left, right, top, bottom float64) {
	_debca := &_fabc._bcbf
	_debca.Left = left
	_debca.Right = right
	_debca.Top = top
	_debca.Bottom = bottom
}
func _geaf(_aafg *templateProcessor, _fbcad *templateNode) (interface{}, error) {
	return _aafg.parseTextChunk(_fbcad)
}

// DrawTemplate renders the template provided through the specified reader,
// using the specified `data` and `options`.
// Creator templates are first executed as text/template *Template instances,
// so the specified `data` is inserted within the template.
// The second phase of processing is actually parsing the template, translating
// it into creator components and rendering them using the provided options.
// Both the `data` and `options` parameters can be nil.
func (_ebb *Block) DrawTemplate(c *Creator, r _ef.Reader, data interface{}, options *TemplateOptions) error {
	return _edfbf(c, r, data, options, _ebb)
}

// GetIndent get the cell's left indent.
func (_dggb *TableCell) GetIndent() float64 { return _dggb._cdgg }

// AddInfo is used to append a piece of invoice information in the template
// information table.
func (_gcffa *Invoice) AddInfo(description, value string) (*InvoiceCell, *InvoiceCell) {
	_dfaa := [2]*InvoiceCell{_gcffa.newCell(description, _gcffa._fcabd), _gcffa.newCell(value, _gcffa._fcabd)}
	_gcffa._agef = append(_gcffa._agef, _dfaa)
	return _dfaa[0], _dfaa[1]
}

// SetLineTitleStyle sets the style for the title part of all new lines
// of the table of contents.
func (_egggb *TOC) SetLineTitleStyle(style TextStyle) { _egggb._ecfgfa = style }
func (_ccba *pageTransformations) applyFlip(_bcbb *_f.PdfPage) error {
	_dded, _cgc := _ccba._gdd, _ccba._cbcb
	if !_dded && !_cgc {
		return nil
	}
	if _bcbb == nil {
		return _d.New("\u006e\u006f\u0020\u0070\u0061\u0067\u0065\u0020\u0061c\u0074\u0069\u0076\u0065")
	}
	_gbab, _fgcf := _bcbb.GetMediaBox()
	if _fgcf != nil {
		return _fgcf
	}
	_fffe, _eedg := _gbab.Width(), _gbab.Height()
	_gadd, _fgcf := _bcbb.GetRotate()
	if _fgcf != nil {
		_df.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _fgcf.Error())
	}
	if _cege := _gadd%360 != 0 && _gadd%90 == 0; _cege {
		if _fgfb := (360 + _gadd%360) % 360; _fgfb == 90 || _fgfb == 270 {
			_dded, _cgc = _cgc, _dded
		}
	}
	_feb, _dff := 1.0, 0.0
	if _dded {
		_feb, _dff = -1.0, -_fffe
	}
	_abeg, _aeeaa := 1.0, 0.0
	if _cgc {
		_abeg, _aeeaa = -1.0, -_eedg
	}
	_adcd := _eee.NewContentCreator().Scale(_feb, _abeg).Translate(_dff, _aeeaa)
	_dac, _fgcf := _fg.MakeStream(_adcd.Bytes(), _fg.NewFlateEncoder())
	if _fgcf != nil {
		return _fgcf
	}
	_gbd := _fg.MakeArray(_dac)
	_gbd.Append(_bcbb.GetContentStreamObjs()...)
	_bcbb.Contents = _gbd
	return nil
}

// SetLineLevelOffset sets the amount of space an indentation level occupies
// for all new lines of the table of contents.
func (_aefcf *TOC) SetLineLevelOffset(levelOffset float64) { _aefcf._fgbea = levelOffset }

// List represents a list of items.
// The representation of a list item is as follows:
//       [marker] [content]
// e.g.:        • This is the content of the item.
// The supported components to add content to list items are:
// - Paragraph
// - StyledParagraph
// - List
type List struct {
	_eddg  []*listItem
	_bdccd Margins
	_fbfe  TextChunk
	_dgab  float64
	_ggbf  bool
	_acdb  Positioning
	_edae  TextStyle
}

// SetBorderColor sets the border color.
func (_efdef *CurvePolygon) SetBorderColor(color Color) { _efdef._dbca.BorderColor = _bcfe(color) }
func (_fged *templateProcessor) run() error {
	_cagde := _cc.NewDecoder(_b.NewReader(_fged._ceac))
	var _cgfb *templateNode
	for {
		_gbcc, _cfedg := _cagde.Token()
		if _cfedg != nil {
			if _cfedg == _ef.EOF {
				return nil
			}
			return _cfedg
		}
		if _gbcc == nil {
			break
		}
		switch _abag := _gbcc.(type) {
		case _cc.StartElement:
			_df.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006eg\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0073\u0074\u0061r\u0074\u0020\u0074\u0061\u0067\u003a\u0020`\u0025\u0073\u0060\u002e", _abag.Name.Local)
			_bbac, _ffeb := _accge[_abag.Name.Local]
			if !_ffeb {
				_df.Log.Debug("\u0055\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0074a\u0067\u003a\u0020\u0060\u0025\u0073`\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e\u0020\u004f\u0075t\u0070\u0075t\u0020\u006d\u0061\u0079 \u0062\u0065\u0020\u0069\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u002e", _abag.Name.Local)
				continue
			}
			_cgfb = &templateNode{_ddgde: _abag, _dfdd: _cgfb}
			if _cded := _bbac._bgff; _cded != nil {
				_cgfb._cedba, _cfedg = _cded(_fged, _cgfb)
				if _cfedg != nil {
					return _cfedg
				}
			}
		case _cc.EndElement:
			_df.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020t\u0065\u006d\u0070\u006c\u0061\u0074\u0065 \u0065\u006e\u0064\u0020\u0074\u0061\u0067\u003a\u0020\u0060\u0025\u0073\u0060\u002e", _abag.Name.Local)
			if _cgfb != nil {
				if _cgfb._cedba != nil {
					if _fbfb := _fged.renderNode(_cgfb); _fbfb != nil {
						return _fbfb
					}
				}
				_cgfb = _cgfb._dfdd
			}
		case _cc.CharData:
			if _cgfb != nil && _cgfb._cedba != nil {
				if _egcgg := _fged.addNodeText(_cgfb, string(_abag)); _egcgg != nil {
					return _egcgg
				}
			}
		case _cc.Comment:
			_df.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020t\u0065\u006d\u0070\u006c\u0061\u0074\u0065 \u0063\u006f\u006d\u006d\u0065\u006e\u0074\u003a\u0020\u0060\u0025\u0073\u0060\u002e", string(_abag))
		}
	}
	return nil
}
func (_gddgc *templateProcessor) parseCellBorderStyleAttr(_bcfg, _egdgf string) CellBorderStyle {
	_df.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020c\u0065\u006c\u006c b\u006f\u0072\u0064\u0065\u0072\u0020s\u0074\u0079\u006c\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a \u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025s\u0029\u002e", _bcfg, _egdgf)
	_eded := map[string]CellBorderStyle{"\u006e\u006f\u006e\u0065": CellBorderStyleNone, "\u0073\u0069\u006e\u0067\u006c\u0065": CellBorderStyleSingle, "\u0064\u006f\u0075\u0062\u006c\u0065": CellBorderStyleDouble}[_egdgf]
	return _eded
}

// IsRelative checks if the positioning is relative.
func (_bfbb Positioning) IsRelative() bool { return _bfbb == PositionRelative }

// SetDashPattern sets the dash pattern of the line.
// NOTE: the dash pattern is taken into account only if the style of the
// line is set to dashed.
func (_bgfg *Line) SetDashPattern(dashArray []int64, dashPhase int64) {
	_bgfg._gbde = dashArray
	_bgfg._afbf = dashPhase
}

// SetBorderOpacity sets the border opacity.
func (_baea *Polygon) SetBorderOpacity(opacity float64) { _baea._afab = opacity }
func (_bead *Table) moveToNextAvailableCell() int {
	_deafe := (_bead._edag-1)%(_bead._cfee) + 1
	for {
		if _deafe-1 >= len(_bead._fbca) {
			return _deafe
		} else if _bead._fbca[_deafe-1] == 0 {
			return _deafe
		} else {
			_bead._edag++
			_bead._fbca[_deafe-1]--
		}
		_deafe++
	}
}
func (_dcacf *templateProcessor) parsePositioningAttr(_bfef, _agedd string) Positioning {
	_df.Log.Debug("\u0050\u0061\u0072s\u0069\u006e\u0067\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0069\u006e\u0067\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060%\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _bfef, _agedd)
	_efeda := map[string]Positioning{"\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065": PositionRelative, "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065": PositionAbsolute}[_agedd]
	return _efeda
}

// SetIncludeInTOC sets a flag to indicate whether or not to include in tOC.
func (_cda *Chapter) SetIncludeInTOC(includeInTOC bool) { _cda._fcde = includeInTOC }

// GeneratePageBlocks generate the Page blocks. Draws the Image on a block, implementing the Drawable interface.
func (_bacg *Image) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	if _bacg._aabc == nil {
		if _edea := _bacg.makeXObject(); _edea != nil {
			return nil, ctx, _edea
		}
	}
	var _dgbf []*Block
	_cdda := ctx
	_dfe := NewBlock(ctx.PageWidth, ctx.PageHeight)
	if _bacg._ccec.IsRelative() {
		_bacg.applyFitMode(ctx.Width)
		ctx.X += _bacg._dbfe.Left
		ctx.Y += _bacg._dbfe.Top
		ctx.Width -= _bacg._dbfe.Left + _bacg._dbfe.Right
		ctx.Height -= _bacg._dbfe.Top + _bacg._dbfe.Bottom
		if _bacg._gddc > ctx.Height {
			_dgbf = append(_dgbf, _dfe)
			_dfe = NewBlock(ctx.PageWidth, ctx.PageHeight)
			ctx.Page++
			_abbb := ctx
			_abbb.Y = ctx.Margins.Top + _bacg._dbfe.Top
			_abbb.X = ctx.Margins.Left + _bacg._dbfe.Left
			_abbb.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom - _bacg._dbfe.Top - _bacg._dbfe.Bottom
			_abbb.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right - _bacg._dbfe.Left - _bacg._dbfe.Right
			ctx = _abbb
		}
	} else {
		ctx.X = _bacg._cgee
		ctx.Y = _bacg._ddab
	}
	ctx, _baed := _ceeg(_dfe, _bacg, ctx)
	if _baed != nil {
		return nil, ctx, _baed
	}
	_dgbf = append(_dgbf, _dfe)
	if _bacg._ccec.IsAbsolute() {
		ctx = _cdda
	} else {
		ctx.X = _cdda.X
		ctx.Width = _cdda.Width
		ctx.Y += _bacg._dbfe.Bottom
	}
	return _dgbf, ctx, nil
}

// Chart represents a chart drawable.
// It is used to render unichart chart components using a creator instance.
type Chart struct {
	_fbaa _ab.ChartRenderable
	_ebe  Positioning
	_eca  float64
	_dbc  float64
	_daea Margins
}

func (_feggcb *templateProcessor) parseLine(_gffbg *templateNode) (interface{}, error) {
	_ccgg := _feggcb.creator.NewLine(0, 0, 0, 0)
	for _, _gefa := range _gffbg._ddgde.Attr {
		_fdaeb := _gefa.Value
		switch _fdccc := _gefa.Name.Local; _fdccc {
		case "\u0078\u0031":
			_ccgg._ggfd = _feggcb.parseFloatAttr(_fdccc, _fdaeb)
		case "\u0079\u0031":
			_ccgg._fadc = _feggcb.parseFloatAttr(_fdccc, _fdaeb)
		case "\u0078\u0032":
			_ccgg._dccb = _feggcb.parseFloatAttr(_fdccc, _fdaeb)
		case "\u0079\u0032":
			_ccgg._gdec = _feggcb.parseFloatAttr(_fdccc, _fdaeb)
		case "\u0074h\u0069\u0063\u006b\u006e\u0065\u0073s":
			_ccgg.SetLineWidth(_feggcb.parseFloatAttr(_fdccc, _fdaeb))
		case "\u0063\u006f\u006co\u0072":
			_ccgg.SetColor(_feggcb.parseColorAttr(_fdccc, _fdaeb))
		case "\u0073\u0074\u0079l\u0065":
			_ccgg.SetStyle(_feggcb.parseLineStyleAttr(_fdccc, _fdaeb))
		case "\u0064\u0061\u0073\u0068\u002d\u0061\u0072\u0072\u0061\u0079":
			_ccgg.SetDashPattern(_feggcb.parseInt64Array(_fdccc, _fdaeb), _ccgg._afbf)
		case "\u0064\u0061\u0073\u0068\u002d\u0070\u0068\u0061\u0073\u0065":
			_ccgg.SetDashPattern(_ccgg._gbde, _feggcb.parseInt64Attr(_fdccc, _fdaeb))
		case "\u006fp\u0061\u0063\u0069\u0074\u0079":
			_ccgg.SetOpacity(_feggcb.parseFloatAttr(_fdccc, _fdaeb))
		case "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e":
			_ccgg.SetPositioning(_feggcb.parsePositioningAttr(_fdccc, _fdaeb))
		case "\u0066\u0069\u0074\u002d\u006d\u006f\u0064\u0065":
			_ccgg.SetFitMode(_feggcb.parseFitModeAttr(_fdccc, _fdaeb))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_fgdbb := _feggcb.parseMarginAttr(_fdccc, _fdaeb)
			_ccgg.SetMargins(_fgdbb.Left, _fgdbb.Right, _fgdbb.Top, _fgdbb.Bottom)
		default:
			_df.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u006c\u0069\u006e\u0065 \u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _fdccc)
		}
	}
	return _ccgg, nil
}

// Width is not used as the division component is designed to fill all the
// available space, depending on the context. Returns 0.
func (_cdba *Division) Width() float64 { return 0 }

// SetTextAlignment sets the horizontal alignment of the text within the space provided.
func (_cfaa *Paragraph) SetTextAlignment(align TextAlignment) { _cfaa._badg = align }

// Height returns the height of the Paragraph. The height is calculated based on the input text and how it is wrapped
// within the container. Does not include Margins.
func (_dbef *StyledParagraph) Height() float64 {
	_dbef.wrapText()
	var _aebge float64
	for _, _ddffb := range _dbef._bfaa {
		var _gaac float64
		for _, _ecabe := range _ddffb {
			_gacaf := _dbef._bged * _ecabe.Style.FontSize
			if _gacaf > _gaac {
				_gaac = _gacaf
			}
		}
		_aebge += _gaac
	}
	return _aebge
}

// SetWidthTop sets border width for top.
func (_bgac *border) SetWidthTop(bw float64) { _bgac._ebc = bw }
func (_bdaba *templateProcessor) parseCellAlignmentAttr(_fcdaa, _efdfg string) CellHorizontalAlignment {
	_df.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020c\u0065\u006c\u006c\u0020\u0061\u006c\u0069g\u006e\u006d\u0065\u006e\u0074\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028`\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029\u002e", _fcdaa, _efdfg)
	_fdfg := map[string]CellHorizontalAlignment{"\u006c\u0065\u0066\u0074": CellHorizontalAlignmentLeft, "\u0063\u0065\u006e\u0074\u0065\u0072": CellHorizontalAlignmentCenter, "\u0072\u0069\u0067h\u0074": CellHorizontalAlignmentRight}[_efdfg]
	return _fdfg
}

// GeneratePageBlocks generate the Page blocks. Multiple blocks are generated
// if the contents wrap over multiple pages.
func (_cfecb *TOC) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_befeb := ctx
	_begb, ctx, _aafe := _cfecb._gcdb.GeneratePageBlocks(ctx)
	if _aafe != nil {
		return _begb, ctx, _aafe
	}
	for _, _cbdcb := range _cfecb._fbagb {
		_dgeb := _cbdcb._cbfg
		if !_cfecb._abbdf {
			_cbdcb._cbfg = 0
		}
		_gdabc, _afga, _bdfa := _cbdcb.GeneratePageBlocks(ctx)
		_cbdcb._cbfg = _dgeb
		if _bdfa != nil {
			return _begb, ctx, _bdfa
		}
		if len(_gdabc) < 1 {
			continue
		}
		_begb[len(_begb)-1].mergeBlocks(_gdabc[0])
		_begb = append(_begb, _gdabc[1:]...)
		ctx = _afga
	}
	if _cfecb._ebfge.IsRelative() {
		ctx.X = _befeb.X
	}
	if _cfecb._ebfge.IsAbsolute() {
		return _begb, _befeb, nil
	}
	return _begb, ctx, nil
}

// NewCell makes a new cell and inserts it into the table at the current position.
func (_bgeda *Table) NewCell() *TableCell { return _bgeda.MultiCell(1, 1) }

// EnablePageWrap controls whether the division is wrapped across pages.
// If disabled, the division is moved in its entirety on a new page, if it
// does not fit in the available height. By default, page wrapping is enabled.
// If the height of the division is larger than an entire page, wrapping is
// enabled automatically in order to avoid unwanted behavior.
// Currently, page wrapping can only be disabled for vertical divisions.
func (_ggfe *Division) EnablePageWrap(enable bool) { _ggfe._dcd = enable }
func _baafd(_beggg *templateProcessor, _addcd *templateNode) (interface{}, error) {
	return _beggg.parseTableCell(_addcd)
}
func _eafcg(_ebgfd, _ebaaf, _ceed string, _ecfd uint, _eedga TextStyle) *TOCLine {
	return _ffggc(TextChunk{Text: _ebgfd, Style: _eedga}, TextChunk{Text: _ebaaf, Style: _eedga}, TextChunk{Text: _ceed, Style: _eedga}, _ecfd, _eedga)
}

// SetLineColor sets the line color.
func (_edaea *Polyline) SetLineColor(color Color) { _edaea._bcg.LineColor = _bcfe(color) }

// Ellipse defines an ellipse with a center at (xc,yc) and a specified width and height.  The ellipse can have a colored
// fill and/or border with a specified width.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Ellipse struct {
	_agdd  float64
	_ddgg  float64
	_acea  float64
	_gdeg  float64
	_cfagc Positioning
	_cedc  Color
	_agcc  Color
	_fbec  float64
}

func (_cfbg *templateProcessor) parseBackground(_dgbg *templateNode) (interface{}, error) {
	_cfcc := &Background{}
	for _, _bffd := range _dgbg._ddgde.Attr {
		_dccbg := _bffd.Value
		switch _fagde := _bffd.Name.Local; _fagde {
		case "\u0066\u0069\u006c\u006c\u002d\u0063\u006f\u006c\u006f\u0072":
			_cfcc.FillColor = _cfbg.parseColorAttr(_fagde, _dccbg)
		case "\u0062\u006f\u0072d\u0065\u0072\u002d\u0063\u006f\u006c\u006f\u0072":
			_cfcc.BorderColor = _cfbg.parseColorAttr(_fagde, _dccbg)
		case "b\u006f\u0072\u0064\u0065\u0072\u002d\u0073\u0069\u007a\u0065":
			_cfcc.BorderSize = _cfbg.parseFloatAttr(_fagde, _dccbg)
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0072\u0061\u0064\u0069\u0075\u0073":
			_bcfde, _becgf, _eabc, _ddgb := _cfbg.parseBorderRadiusAttr(_fagde, _dccbg)
			_cfcc.SetBorderRadius(_bcfde, _becgf, _ddgb, _eabc)
		case "\u0062\u006f\u0072\u0064er\u002d\u0074\u006f\u0070\u002d\u006c\u0065\u0066\u0074\u002d\u0072\u0061\u0064\u0069u\u0073":
			_cfcc.BorderRadiusTopLeft = _cfbg.parseFloatAttr(_fagde, _dccbg)
		case "\u0062\u006f\u0072de\u0072\u002d\u0074\u006f\u0070\u002d\u0072\u0069\u0067\u0068\u0074\u002d\u0072\u0061\u0064\u0069\u0075\u0073":
			_cfcc.BorderRadiusTopRight = _cfbg.parseFloatAttr(_fagde, _dccbg)
		case "\u0062o\u0072\u0064\u0065\u0072-\u0062\u006f\u0074\u0074\u006fm\u002dl\u0065f\u0074\u002d\u0072\u0061\u0064\u0069\u0075s":
			_cfcc.BorderRadiusBottomLeft = _cfbg.parseFloatAttr(_fagde, _dccbg)
		case "\u0062\u006f\u0072\u0064\u0065\u0072\u002d\u0062\u006f\u0074\u0074o\u006d\u002d\u0072\u0069\u0067\u0068\u0074\u002d\u0072\u0061d\u0069\u0075\u0073":
			_cfcc.BorderRadiusBottomRight = _cfbg.parseFloatAttr(_fagde, _dccbg)
		default:
			_df.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0062\u0061\u0063\u006b\u0067\u0072\u006f\u0075\u006e\u0064\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006bi\u0070p\u0069\u006e\u0067\u002e", _fagde)
		}
	}
	return _cfcc, nil
}

// SetShowLinks sets visibility of links for the TOC lines.
func (_bbdfe *TOC) SetShowLinks(showLinks bool) { _bbdfe._abbdf = showLinks }

// SetWidth sets the the Paragraph width. This is essentially the wrapping width,
// i.e. the width the text can extend to prior to wrapping over to next line.
func (_eaae *StyledParagraph) SetWidth(width float64) { _eaae._bbef = width; _eaae.wrapText() }

// Margins represents page margins or margins around an element.
type Margins struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

// SetFitMode sets the fit mode of the image.
// NOTE: The fit mode is only applied if relative positioning is used.
func (_cfc *Image) SetFitMode(fitMode FitMode) { _cfc._aedff = fitMode }

// GeneratePageBlocks draws the block contents on a template Page block.
// Implements the Drawable interface.
func (_dfa *Block) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_bg := _eee.NewContentCreator()
	_dg, _efb := _dfa.Width(), _dfa.Height()
	if _dfa._ec.IsRelative() {
		_bg.Translate(ctx.X, ctx.PageHeight-ctx.Y-_efb)
	} else {
		_bg.Translate(_dfa._fa, ctx.PageHeight-_dfa._ea-_efb)
	}
	_eec := _efb
	if _dfa._ag != 0 {
		_bg.Translate(_dg/2, _efb/2)
		_bg.RotateDeg(_dfa._ag)
		_bg.Translate(-_dg/2, -_efb/2)
		_, _eec = _dfa.RotatedSize()
	}
	if _dfa._ec.IsRelative() {
		ctx.Y += _eec
	}
	_cee := _dfa.duplicate()
	_bgb := append(*_bg.Operations(), *_cee._ga...)
	_bgb.WrapIfNeeded()
	_cee._ga = &_bgb
	return []*Block{_cee}, ctx, nil
}
func (_fdbgc *TemplateOptions) init() {
	if _fdbgc.FontMap == nil {
		_fdbgc.FontMap = map[string]*_f.PdfFont{}
	}
	if _fdbgc.ImageMap == nil {
		_fdbgc.ImageMap = map[string]*_f.Image{}
	}
	if _fdbgc.ColorMap == nil {
		_fdbgc.ColorMap = map[string]Color{}
	}
	if _fdbgc.ChartMap == nil {
		_fdbgc.ChartMap = map[string]_ab.ChartRenderable{}
	}
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
func (_fadgg *Paragraph) SetColor(col Color) { _fadgg._bfbc = col }

// GeneratePageBlocks generates the table page blocks. Multiple blocks are
// generated if the contents wrap over multiple pages.
// Implements the Drawable interface.
func (_dbfed *Table) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_fadcd := _dbfed
	if _dbfed._ggbfc {
		_fadcd = _dbfed.clone()
	}
	return _effa(_fadcd, ctx)
}

// PolyBezierCurve represents a composite curve that is the result of joining
// multiple cubic Bezier curves.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type PolyBezierCurve struct {
	_cegf  *_cec.PolyBezierCurve
	_bbeag float64
	_bcbe  float64
}

func (_facbd *templateProcessor) loadImageFromSrc(_cbeee string) (*Image, error) {
	if _cbeee == "" {
		_df.Log.Error("\u0049\u006d\u0061\u0067\u0065\u0020\u0060\u0073\u0072\u0063\u0060\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062e\u0020\u0065m\u0070\u0074\u0079\u002e")
		return nil, _adbab
	}
	_aaff := _ed.Split(_cbeee, "\u002c")
	for _, _egbb := range _aaff {
		_egbb = _ed.TrimSpace(_egbb)
		if _egbb == "" {
			continue
		}
		_dbcgc, _gdca := _facbd._daae.ImageMap[_egbb]
		if _gdca {
			return _baec(_dbcgc)
		}
		if _ecefc := _facbd.parseAttrPropList(_egbb); len(_ecefc) > 0 {
			if _gaff, _ebbfe := _ecefc["\u0070\u0061\u0074\u0068"]; _ebbfe {
				if _aegc, _ebgbf := _fffeg(_gaff); _ebgbf != nil {
					_df.Log.Debug("\u0043\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020l\u006f\u0061\u0064\u0020\u0069\u006d\u0061g\u0065\u0020\u0060\u0025\u0073\u0060\u003a\u0020\u0025\u0076\u002e", _gaff, _ebgbf)
				} else {
					return _aegc, nil
				}
			}
		}
	}
	_df.Log.Error("\u0043\u006ful\u0064\u0020\u006eo\u0074\u0020\u0066\u0069nd \u0069ma\u0067\u0065\u0020\u0072\u0065\u0073\u006fur\u0063\u0065\u003a\u0020\u0060\u0025\u0073`\u002e", _cbeee)
	return nil, _adbab
}
func (_feggc *TableCell) cloneProps(_bdffg VectorDrawable) *TableCell {
	_cafe := *_feggc
	_cafe._cbgcb = _bdffg
	return &_cafe
}

// SetFillColor sets background color for border.
func (_gfdf *border) SetFillColor(col Color) { _gfdf._edg = col }
func (_fdac *Division) drawBackground(_eedd []*Block, _fde, _cbdb DrawContext, _effc bool) ([]*Block, error) {
	_febf := len(_eedd)
	if _febf == 0 || _fdac._bgeg == nil {
		return _eedd, nil
	}
	_fdda := make([]*Block, 0, len(_eedd))
	for _cbg, _gddf := range _eedd {
		var (
			_febaa = _fdac._bgeg.BorderRadiusTopLeft
			_eega  = _fdac._bgeg.BorderRadiusTopRight
			_dbcg  = _fdac._bgeg.BorderRadiusBottomLeft
			_cgff  = _fdac._bgeg.BorderRadiusBottomRight
		)
		_ebcd := _fde
		_ebcd.Page += _cbg
		if _cbg == 0 {
			if _effc {
				_fdda = append(_fdda, _gddf)
				continue
			}
			if _febf == 1 {
				_ebcd.Height = _cbdb.Y - _fde.Y
			}
		} else {
			_ebcd.X = _ebcd.Margins.Left + _fdac._fccf.Left
			_ebcd.Y = _ebcd.Margins.Top
			_ebcd.Width = _ebcd.PageWidth - _ebcd.Margins.Left - _ebcd.Margins.Right - _fdac._fccf.Left - _fdac._fccf.Right
			if _cbg == _febf-1 {
				_ebcd.Height = _cbdb.Y - _ebcd.Margins.Top - _fdac._fccf.Top
			} else {
				_ebcd.Height = _ebcd.PageHeight - _ebcd.Margins.Top - _ebcd.Margins.Bottom
			}
			if !_effc {
				_febaa = 0
				_eega = 0
			}
		}
		if _febf > 1 && _cbg != _febf-1 {
			_dbcg = 0
			_cgff = 0
		}
		_aefb := _aeac(_ebcd.X, _ebcd.Y, _ebcd.Width, _ebcd.Height)
		_aefb.SetFillColor(_fdac._bgeg.FillColor)
		_aefb.SetBorderColor(_fdac._bgeg.BorderColor)
		_aefb.SetBorderWidth(_fdac._bgeg.BorderSize)
		_aefb.SetBorderRadius(_febaa, _eega, _dbcg, _cgff)
		_aca, _, _gfb := _aefb.GeneratePageBlocks(_ebcd)
		if _gfb != nil {
			return nil, _gfb
		}
		if len(_aca) == 0 {
			continue
		}
		_bffbe := _aca[0]
		if _gfb = _bffbe.mergeBlocks(_gddf); _gfb != nil {
			return nil, _gfb
		}
		_fdda = append(_fdda, _bffbe)
	}
	return _fdda, nil
}

// GeneratePageBlocks draws the composite curve polygon on a new block
// representing the page. Implements the Drawable interface.
func (_cbbf *CurvePolygon) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_bcbc := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_abbf, _dbfc := _bcbc.setOpacity(_cbbf._cdge, _cbbf._cedec)
	if _dbfc != nil {
		return nil, ctx, _dbfc
	}
	_ffae := _cbbf._dbca
	_ffae.FillEnabled = _ffae.FillColor != nil
	_ffae.BorderEnabled = _ffae.BorderColor != nil && _ffae.BorderWidth > 0
	var (
		_fecg = ctx.PageHeight
		_aebf = _ffae.Rings
		_afgc = make([][]_cec.CubicBezierCurve, 0, len(_ffae.Rings))
	)
	for _, _fbf := range _aebf {
		_gcgc := make([]_cec.CubicBezierCurve, 0, len(_fbf))
		for _, _fbbb := range _fbf {
			_cbff := _fbbb
			_cbff.P0.Y = _fecg - _cbff.P0.Y
			_cbff.P1.Y = _fecg - _cbff.P1.Y
			_cbff.P2.Y = _fecg - _cbff.P2.Y
			_cbff.P3.Y = _fecg - _cbff.P3.Y
			_gcgc = append(_gcgc, _cbff)
		}
		_afgc = append(_afgc, _gcgc)
	}
	_ffae.Rings = _afgc
	defer func() { _ffae.Rings = _aebf }()
	_ffff, _, _dbfc := _ffae.Draw(_abbf)
	if _dbfc != nil {
		return nil, ctx, _dbfc
	}
	if _dbfc = _bcbc.addContentsByString(string(_ffff)); _dbfc != nil {
		return nil, ctx, _dbfc
	}
	return []*Block{_bcbc}, ctx, nil
}
func (_cfgb *templateProcessor) parseStyledParagraph(_befg *templateNode) (interface{}, error) {
	_edcg := _cfgb.creator.NewStyledParagraph()
	for _, _afgeb := range _befg._ddgde.Attr {
		_ebae := _afgeb.Value
		switch _aadbb := _afgeb.Name.Local; _aadbb {
		case "\u0074\u0065\u0078\u0074\u002d\u0061\u006c\u0069\u0067\u006e":
			_edcg.SetTextAlignment(_cfgb.parseTextAlignmentAttr(_aadbb, _ebae))
		case "\u0076\u0065\u0072\u0074ic\u0061\u006c\u002d\u0074\u0065\u0078\u0074\u002d\u0061\u006c\u0069\u0067\u006e":
			_edcg.SetTextVerticalAlignment(_cfgb.parseTextVerticalAlignmentAttr(_aadbb, _ebae))
		case "l\u0069\u006e\u0065\u002d\u0068\u0065\u0069\u0067\u0068\u0074":
			_edcg.SetLineHeight(_cfgb.parseFloatAttr(_aadbb, _ebae))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_befd := _cfgb.parseMarginAttr(_aadbb, _ebae)
			_edcg.SetMargins(_befd.Left, _befd.Right, _befd.Top, _befd.Bottom)
		case "e\u006e\u0061\u0062\u006c\u0065\u002d\u0077\u0072\u0061\u0070":
			_edcg.SetEnableWrap(_cfgb.parseBoolAttr(_aadbb, _ebae))
		case "\u0065\u006ea\u0062\u006c\u0065-\u0077\u006f\u0072\u0064\u002d\u0077\u0072\u0061\u0070":
			_edcg.EnableWordWrap(_cfgb.parseBoolAttr(_aadbb, _ebae))
		case "\u0074\u0065\u0078\u0074\u002d\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077":
			_edcg.SetTextOverflow(_cfgb.parseTextOverflowAttr(_aadbb, _ebae))
		case "\u0078":
			_edcg.SetPos(_cfgb.parseFloatAttr(_aadbb, _ebae), _edcg._aeegd)
		case "\u0079":
			_edcg.SetPos(_edcg._ggbc, _cfgb.parseFloatAttr(_aadbb, _ebae))
		case "\u0061\u006e\u0067l\u0065":
			_edcg.SetAngle(_cfgb.parseFloatAttr(_aadbb, _ebae))
		default:
			_df.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064 \u0073\u0074\u0079l\u0065\u0064 \u0070\u0061\u0072\u0061\u0067\u0072a\u0070h \u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0060\u0025\u0073\u0060\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e", _aadbb)
		}
	}
	return _edcg, nil
}

// AddSection adds a new content section at the end of the invoice.
func (_fccg *Invoice) AddSection(title, content string) {
	_fccg._bgfee = append(_fccg._bgfee, [2]string{title, content})
}

// ScaleToWidth scales the Block to a specified width, maintaining the same aspect ratio.
func (_gff *Block) ScaleToWidth(w float64) { _be := w / _gff._bf; _gff.Scale(_be, _be) }
func (_abcc *Paragraph) getTextWidth() float64 {
	_dabf := 0.0
	for _, _agbg := range _abcc._ebea {
		if _agbg == '\u000A' {
			continue
		}
		_ddb, _ccfa := _abcc._ceda.GetRuneMetrics(_agbg)
		if !_ccfa {
			_df.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0052u\u006e\u0065\u0020\u0063\u0068a\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0028\u0072\u0075\u006e\u0065\u0020\u0030\u0078\u0025\u0030\u0034\u0078\u003d\u0025\u0063\u0029", _agbg, _agbg)
			return -1
		}
		_dabf += _abcc._gced * _ddb.Wx
	}
	return _dabf
}

// StyledParagraph represents text drawn with a specified font and can wrap across lines and pages.
// By default occupies the available width in the drawing context.
type StyledParagraph struct {
	_eedad []*TextChunk
	_gdbeb TextStyle
	_bdbf  TextStyle
	_bcege TextAlignment
	_ceab  TextVerticalAlignment
	_bged  float64
	_ecbf  bool
	_bbef  float64
	_bggd  bool
	_ffbb  bool
	_gcb   TextOverflow
	_daab  float64
	_afdba Margins
	_geagd Positioning
	_ggbc  float64
	_aeegd float64
	_fgaea float64
	_fbgba float64
	_bfaa  [][]*TextChunk
	_bccd  func(_dgba *StyledParagraph, _dffa DrawContext)
}

// CellBorderStyle defines the table cell's border style.
type CellBorderStyle int

// SetBorderOpacity sets the border opacity.
func (_fded *Rectangle) SetBorderOpacity(opacity float64) { _fded._beeac = opacity }

// SetLevelOffset sets the amount of space an indentation level occupies.
func (_cbabc *TOCLine) SetLevelOffset(levelOffset float64) {
	_cbabc._aaaa = levelOffset
	_cbabc._face._afdba.Left = _cbabc._ecdbc + float64(_cbabc._dgcda-1)*_cbabc._aaaa
}

// AddTotalLine adds a new line in the invoice totals table.
func (_gfab *Invoice) AddTotalLine(desc, value string) (*InvoiceCell, *InvoiceCell) {
	_badfa := &InvoiceCell{_gfab._fabge, desc}
	_bcacc := &InvoiceCell{_gfab._fabge, value}
	_gfab._gce = append(_gfab._gce, [2]*InvoiceCell{_badfa, _bcacc})
	return _badfa, _bcacc
}

// SetFillColor sets the fill color.
func (_dfab *PolyBezierCurve) SetFillColor(color Color) { _dfab._cegf.FillColor = _bcfe(color) }
func _feda(_cfefd, _bfbdb, _ebab float64) (_eebe, _dgfb, _adadaa, _cbcg float64) {
	if _ebab == 0 {
		return 0, 0, _cfefd, _bfbdb
	}
	_edaf := _cec.Path{Points: []_cec.Point{_cec.NewPoint(0, 0).Rotate(_ebab), _cec.NewPoint(_cfefd, 0).Rotate(_ebab), _cec.NewPoint(0, _bfbdb).Rotate(_ebab), _cec.NewPoint(_cfefd, _bfbdb).Rotate(_ebab)}}.GetBoundingBox()
	return _edaf.X, _edaf.Y, _edaf.Width, _edaf.Height
}
func _ecef(_dffc int) *Table {
	_afggf := &Table{_cfee: _dffc, _dcgce: 10.0, _agdgd: []float64{}, _ceeb: []float64{}, _ebfb: []*TableCell{}, _fbca: make([]int, _dffc), _fcacc: true}
	_afggf.resetColumnWidths()
	return _afggf
}
func _bcae(_gcff []byte) (*Image, error) {
	_dcaf := _b.NewReader(_gcff)
	_agda, _bfcbe := _f.ImageHandling.Read(_dcaf)
	if _bfcbe != nil {
		_df.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _bfcbe)
		return nil, _bfcbe
	}
	return _baec(_agda)
}

// SetBorderWidth sets the border width.
func (_cecc *Rectangle) SetBorderWidth(bw float64) { _cecc._bffaf = bw }

// Drawable is a widget that can be used to draw with the Creator.
type Drawable interface {

	// GeneratePageBlocks draw onto blocks representing Page contents. As the content can wrap over many pages, multiple
	// templates are returned, one per Page.  The function also takes a draw context containing information
	// where to draw (if relative positioning) and the available height to draw on accounting for Margins etc.
	GeneratePageBlocks(_caf DrawContext) ([]*Block, DrawContext, error)
}

// SetMargins sets the Chapter margins: left, right, top, bottom.
// Typically not needed as the creator's page margins are used.
func (_cdb *Chapter) SetMargins(left, right, top, bottom float64) {
	_cdb._ccf.Left = left
	_cdb._ccf.Right = right
	_cdb._ccf.Top = top
	_cdb._ccf.Bottom = bottom
}

// SetBorderLineStyle sets border style (currently dashed or plain).
func (_cccdf *TableCell) SetBorderLineStyle(style _cec.LineStyle) { _cccdf._degeb = style }

// SetFillColor sets the fill color.
func (_dabg *Rectangle) SetFillColor(col Color) { _dabg._aebg = col }
func _agf(_bca string, _fdae _fg.PdfObject, _bcb *_f.PdfPageResources) _fg.PdfObjectName {
	_eed := _ed.TrimRightFunc(_ed.TrimSpace(_bca), func(_aee rune) bool { return _ee.IsNumber(_aee) })
	if _eed == "" {
		_eed = "\u0046\u006f\u006e\u0074"
	}
	_dgg := 0
	_faac := _fg.PdfObjectName(_bca)
	for {
		_edd, _fac := _bcb.GetFontByName(_faac)
		if !_fac || _edd == _fdae {
			break
		}
		_dgg++
		_faac = _fg.PdfObjectName(_eg.Sprintf("\u0025\u0073\u0025\u0064", _eed, _dgg))
	}
	return _faac
}
func (_cegbf *Invoice) setCellBorder(_bgdc *TableCell, _gffeb *InvoiceCell) {
	for _, _acbb := range _gffeb.BorderSides {
		_bgdc.SetBorder(_acbb, CellBorderStyleSingle, _gffeb.BorderWidth)
	}
	_bgdc.SetBorderColor(_gffeb.BorderColor)
}

// NewTable create a new Table with a specified number of columns.
func (_feed *Creator) NewTable(cols int) *Table { return _ecef(cols) }
func (_ffaa *templateProcessor) parseHorizontalAlignmentAttr(_edagb, _ebdbg string) HorizontalAlignment {
	_df.Log.Debug("\u0050\u0061\u0072\u0073\u0069n\u0067\u0020\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u0020a\u006c\u0069\u0067\u006e\u006d\u0065\u006e\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c\u0020\u0025\u0073\u0029.", _edagb, _ebdbg)
	_dbec := map[string]HorizontalAlignment{"\u006c\u0065\u0066\u0074": HorizontalAlignmentLeft, "\u0063\u0065\u006e\u0074\u0065\u0072": HorizontalAlignmentCenter, "\u0072\u0069\u0067h\u0074": HorizontalAlignmentRight}[_ebdbg]
	return _dbec
}
func (_fedc *templateProcessor) parseBorderRadiusAttr(_afacg, _addbg string) (_gaaca, _abgg, _gafab, _cbge float64) {
	_df.Log.Debug("\u0050a\u0072\u0073i\u006e\u0067\u0020\u0062o\u0072\u0064\u0065r\u0020\u0072\u0061\u0064\u0069\u0075\u0073\u0020\u0061tt\u0072\u0069\u0062u\u0074\u0065:\u0020\u0028\u0060\u0025\u0073\u0060,\u0020\u0025s\u0029\u002e", _afacg, _addbg)
	switch _baff := _ed.Fields(_addbg); len(_baff) {
	case 1:
		_gaaca, _ = _ae.ParseFloat(_baff[0], 64)
		_abgg = _gaaca
		_gafab = _gaaca
		_cbge = _gaaca
	case 2:
		_gaaca, _ = _ae.ParseFloat(_baff[0], 64)
		_gafab = _gaaca
		_abgg, _ = _ae.ParseFloat(_baff[1], 64)
		_cbge = _abgg
	case 3:
		_gaaca, _ = _ae.ParseFloat(_baff[0], 64)
		_abgg, _ = _ae.ParseFloat(_baff[1], 64)
		_cbge = _abgg
		_gafab, _ = _ae.ParseFloat(_baff[2], 64)
	case 4:
		_gaaca, _ = _ae.ParseFloat(_baff[0], 64)
		_abgg, _ = _ae.ParseFloat(_baff[1], 64)
		_gafab, _ = _ae.ParseFloat(_baff[2], 64)
		_cbge, _ = _ae.ParseFloat(_baff[3], 64)
	}
	return _gaaca, _abgg, _gafab, _cbge
}

// TotalLines returns all the rows in the invoice totals table as
// description-value cell pairs.
func (_aede *Invoice) TotalLines() [][2]*InvoiceCell {
	_dfbe := [][2]*InvoiceCell{_aede._faeg}
	_dfbe = append(_dfbe, _aede._gce...)
	return append(_dfbe, _aede._defg)
}
func (_agddc *Paragraph) wrapText() error {
	if !_agddc._bbfc || int(_agddc._bdga) <= 0 {
		_agddc._cead = []string{_agddc._ebea}
		return nil
	}
	_dadf := NewTextChunk(_agddc._ebea, TextStyle{Font: _agddc._ceda, FontSize: _agddc._gced})
	_aadd, _cdfeb := _dadf.Wrap(_agddc._bdga)
	if _cdfeb != nil {
		return _cdfeb
	}
	if _agddc._eedca > 0 && len(_aadd) > _agddc._eedca {
		_aadd = _aadd[:_agddc._eedca]
	}
	_agddc._cead = _aadd
	return nil
}
func (_aga *StyledParagraph) getTextWidth() float64 {
	var _ceegf float64
	_gffd := len(_aga._eedad)
	for _afac, _ddcd := range _aga._eedad {
		_ebgac := &_ddcd.Style
		_bddd := len(_ddcd.Text)
		for _bfde, _fffegd := range _ddcd.Text {
			if _fffegd == '\u000A' {
				continue
			}
			_agbf, _deeb := _ebgac.Font.GetRuneMetrics(_fffegd)
			if !_deeb {
				_df.Log.Debug("\u0052\u0075\u006e\u0065\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069c\u0073 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0021\u0020\u0025\u0076\u000a", _fffegd)
				return -1
			}
			_ceegf += _ebgac.FontSize * _agbf.Wx * _ebgac.horizontalScale()
			if _fffegd != ' ' && (_afac != _gffd-1 || _bfde != _bddd-1) {
				_ceegf += _ebgac.CharSpacing * 1000.0
			}
		}
	}
	return _ceegf
}

// NewFilledCurve returns a instance of filled curve.
func (_afge *Creator) NewFilledCurve() *FilledCurve { return _gbef() }
func _adaeb(_acdec string) ([]string, error) {
	var (
		_adce  []string
		_bafga []rune
	)
	for _, _bfbeba := range _acdec {
		if _bfbeba == '\u000A' {
			if len(_bafga) > 0 {
				_adce = append(_adce, string(_bafga))
			}
			_adce = append(_adce, string(_bfbeba))
			_bafga = nil
			continue
		}
		_bafga = append(_bafga, _bfbeba)
	}
	if len(_bafga) > 0 {
		_adce = append(_adce, string(_bafga))
	}
	var _cfgec []string
	for _, _abac := range _adce {
		_cfdcdf := []rune(_abac)
		_fcabbb := _aec.NewScanner(_cfdcdf)
		var _baagc []rune
		for _fabfa := 0; _fabfa < len(_cfdcdf); _fabfa++ {
			_, _ddagf, _eedadg := _fcabbb.Next()
			if _eedadg != nil {
				return nil, _eedadg
			}
			if _ddagf == _aec.BreakProhibited || _ee.IsSpace(_cfdcdf[_fabfa]) {
				_baagc = append(_baagc, _cfdcdf[_fabfa])
				if _ee.IsSpace(_cfdcdf[_fabfa]) {
					_cfgec = append(_cfgec, string(_baagc))
					_baagc = []rune{}
				}
				continue
			} else {
				if len(_baagc) > 0 {
					_cfgec = append(_cfgec, string(_baagc))
				}
				_baagc = []rune{_cfdcdf[_fabfa]}
			}
		}
		if len(_baagc) > 0 {
			_cfgec = append(_cfgec, string(_baagc))
		}
	}
	return _cfgec, nil
}

// SetStyleLeft sets border style for left side.
func (_facb *border) SetStyleLeft(style CellBorderStyle) { _facb._afg = style }
func (_agdb *Paragraph) getMaxLineWidth() float64 {
	if _agdb._cead == nil || len(_agdb._cead) == 0 {
		_agdb.wrapText()
	}
	var _abed float64
	for _, _fdfe := range _agdb._cead {
		_cdbb := _agdb.getTextLineWidth(_fdfe)
		if _cdbb > _abed {
			_abed = _cdbb
		}
	}
	return _abed
}

// Lines returns all the lines the table of contents has.
func (_cebf *TOC) Lines() []*TOCLine { return _cebf._fbagb }

// SetHorizontalAlignment sets the horizontal alignment of the image.
func (_gfeea *Image) SetHorizontalAlignment(alignment HorizontalAlignment) { _gfeea._cedbd = alignment }

// SetTerms sets the terms and conditions section of the invoice.
func (_eccb *Invoice) SetTerms(title, content string) { _eccb._dcdf = [2]string{title, content} }
func (_agbd *StyledParagraph) getTextHeight() float64 {
	var _bage float64
	for _, _cabc := range _agbd._eedad {
		_gdcg := _cabc.Style.FontSize * _agbd._bged
		if _gdcg > _bage {
			_bage = _gdcg
		}
	}
	return _bage
}

// SetSellerAddress sets the seller address of the invoice.
func (_gfaf *Invoice) SetSellerAddress(address *InvoiceAddress) { _gfaf._bgfe = address }

// NewTextStyle creates a new text style object which can be used to style
// chunks of text.
// Default attributes:
// Font: Helvetica
// Font size: 10
// Encoding: WinAnsiEncoding
// Text color: black
func (_cacc *Creator) NewTextStyle() TextStyle { return _dadda(_cacc._ffgd) }

// SetOutlineTree adds the specified outline tree to the PDF file generated
// by the creator. Adding an external outline tree disables the automatic
// generation of outlines done by the creator for the relevant components.
func (_abg *Creator) SetOutlineTree(outlineTree *_f.PdfOutlineTreeNode) { _abg._afed = outlineTree }
func (_daefb *TextChunk) clone() *TextChunk {
	_fgaad := *_daefb
	_fgaad._ceddd = _ccgf(_daefb._ceddd)
	return &_fgaad
}

// DrawFooter sets a function to draw a footer on created output pages.
func (_gfac *Creator) DrawFooter(drawFooterFunc func(_dgb *Block, _fagb FooterFunctionArgs)) {
	_gfac._eeb = drawFooterFunc
}

// AddressHeadingStyle returns the style properties used to render the
// heading of the invoice address sections.
func (_ccece *Invoice) AddressHeadingStyle() TextStyle { return _ccece._gfdd }

// NewDivision returns a new Division container component.
func (_edbb *Creator) NewDivision() *Division { return _cdde() }

// SetEnableWrap sets the line wrapping enabled flag.
func (_eagab *Paragraph) SetEnableWrap(enableWrap bool) {
	_eagab._bbfc = enableWrap
	_eagab._bbbeg = false
}

// AppendColumn appends a column to the line items table.
func (_beddd *Invoice) AppendColumn(description string) *InvoiceCell {
	_baaf := _beddd.NewColumn(description)
	_beddd._afeb = append(_beddd._afeb, _baaf)
	return _baaf
}
func (_gacb *List) tableHeight(_bcbd float64) float64 {
	var _egfg float64
	for _, _ffdff := range _gacb._eddg {
		switch _eaaa := _ffdff._cdef.(type) {
		case *Paragraph:
			_dddac := _eaaa
			if _dddac._bbfc {
				_dddac.SetWidth(_bcbd)
			}
			_egfg += _dddac.Height() + _dddac._eacf.Bottom + _dddac._eacf.Bottom
			_egfg += 0.5 * _dddac._gced * _dddac._abbg
		case *StyledParagraph:
			_bfffa := _eaaa
			if _bfffa._ecbf {
				_bfffa.SetWidth(_bcbd)
			}
			_egfg += _bfffa.Height() + _bfffa._afdba.Top + _bfffa._afdba.Bottom
			_egfg += 0.5 * _bfffa.getTextHeight()
		default:
			_egfg += _ffdff._cdef.Height()
		}
	}
	return _egfg
}
func (_dcgc *Chapter) headingNumber() string {
	var _bgcd string
	if _dcgc._bdc {
		if _dcgc._aebc != 0 {
			_bgcd = _ae.Itoa(_dcgc._aebc) + "\u002e"
		}
		if _dcgc._cadbb != nil {
			_bdae := _dcgc._cadbb.headingNumber()
			if _bdae != "" {
				_bgcd = _bdae + _bgcd
			}
		}
	}
	return _bgcd
}

// SetHeaderRows turns the selected table rows into headers that are repeated
// for every page the table spans. startRow and endRow are inclusive.
func (_gcge *Table) SetHeaderRows(startRow, endRow int) error {
	if startRow <= 0 {
		return _d.New("\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0073\u0074\u0061\u0072\u0074\u0020r\u006f\u0077\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	if endRow <= 0 {
		return _d.New("\u0068\u0065a\u0064\u0065\u0072\u0020e\u006e\u0064 \u0072\u006f\u0077\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	if startRow > endRow {
		return _d.New("\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0072\u006f\u0077\u0020\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0074\u0068\u0065 \u0065\u006e\u0064\u0020\u0072o\u0077")
	}
	_gcge._dabbc = true
	_gcge._fdfff = startRow
	_gcge._bdffd = endRow
	return nil
}

// GetHorizontalAlignment returns the horizontal alignment of the image.
func (_deed *Image) GetHorizontalAlignment() HorizontalAlignment { return _deed._cedbd }
func _eabf(_ggfeg int64, _eeee, _ddfcc, _dacg float64) *_f.PdfAnnotation {
	_agdf := _f.NewPdfAnnotationLink()
	_bbece := _f.NewBorderStyle()
	_bbece.SetBorderWidth(0)
	_agdf.BS = _bbece.ToPdfObject()
	if _ggfeg < 0 {
		_ggfeg = 0
	}
	_agdf.Dest = _fg.MakeArray(_fg.MakeInteger(_ggfeg), _fg.MakeName("\u0058\u0059\u005a"), _fg.MakeFloat(_eeee), _fg.MakeFloat(_ddfcc), _fg.MakeFloat(_dacg))
	return _agdf.PdfAnnotation
}

// TOC returns the table of contents component of the creator.
func (_ggdc *Creator) TOC() *TOC { return _ggdc._cddg }

// SetNoteStyle sets the style properties used to render the content of the
// invoice note sections.
func (_dbcf *Invoice) SetNoteStyle(style TextStyle) { _dbcf._egf = style }

// CurCol returns the currently active cell's column number.
func (_egcd *Table) CurCol() int { _ceaa := (_egcd._edag-1)%(_egcd._cfee) + 1; return _ceaa }

// GetCoords returns coordinates of the Rectangle's upper left corner (x,y).
func (_fecbf *Rectangle) GetCoords() (float64, float64) { return _fecbf._bagb, _fecbf._cbe }
func (_cgge *InvoiceAddress) fmtLine(_dcbg, _gded string, _bddb bool) string {
	if _bddb {
		_gded = ""
	}
	return _eg.Sprintf("\u0025\u0073\u0025s\u000a", _gded, _dcbg)
}

// GeneratePageBlocks implements drawable interface.
func (_aeeb *border) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_fbc := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_gfdg := _aeeb._cdfc
	_fce := ctx.PageHeight - _aeeb._aaab
	if _aeeb._edg != nil {
		_dgfc := _cec.Rectangle{Opacity: 1.0, X: _aeeb._cdfc, Y: ctx.PageHeight - _aeeb._aaab - _aeeb._gbgc, Height: _aeeb._gbgc, Width: _aeeb._dgf}
		_dgfc.FillEnabled = true
		_dgfc.FillColor = _bcfe(_aeeb._edg)
		_dgfc.BorderEnabled = false
		_eaad, _, _ddc := _dgfc.Draw("")
		if _ddc != nil {
			return nil, ctx, _ddc
		}
		_ddc = _fbc.addContentsByString(string(_eaad))
		if _ddc != nil {
			return nil, ctx, _ddc
		}
	}
	_fbae := _aeeb._ebc
	_fea := _aeeb._gbe
	_dcf := _aeeb._afeg
	_bgbg := _aeeb._fe
	_acfa := _aeeb._ebc
	if _aeeb._faaa == CellBorderStyleDouble {
		_acfa += 2 * _fbae
	}
	_gbgg := _aeeb._gbe
	if _aeeb._aeba == CellBorderStyleDouble {
		_gbgg += 2 * _fea
	}
	_gbb := _aeeb._afeg
	if _aeeb._afg == CellBorderStyleDouble {
		_gbb += 2 * _dcf
	}
	_gfa := _aeeb._fe
	if _aeeb._bgad == CellBorderStyleDouble {
		_gfa += 2 * _bgbg
	}
	_bfb := (_acfa - _gbb) / 2
	_dcef := (_acfa - _gfa) / 2
	_fef := (_gbgg - _gbb) / 2
	_fag := (_gbgg - _gfa) / 2
	if _aeeb._ebc != 0 {
		_fee := _gfdg
		_eef := _fce
		if _aeeb._faaa == CellBorderStyleDouble {
			_eef -= _fbae
			_cedg := _cec.BasicLine{LineColor: _bcfe(_aeeb._gad), Opacity: 1.0, LineWidth: _aeeb._ebc, LineStyle: _aeeb.LineStyle, X1: _fee - _acfa/2 + _bfb, Y1: _eef + 2*_fbae, X2: _fee + _acfa/2 - _dcef + _aeeb._dgf, Y2: _eef + 2*_fbae}
			_cag, _, _cace := _cedg.Draw("")
			if _cace != nil {
				return nil, ctx, _cace
			}
			_cace = _fbc.addContentsByString(string(_cag))
			if _cace != nil {
				return nil, ctx, _cace
			}
		}
		_eeda := _cec.BasicLine{LineWidth: _aeeb._ebc, Opacity: 1.0, LineColor: _bcfe(_aeeb._gad), LineStyle: _aeeb.LineStyle, X1: _fee - _acfa/2 + _bfb + (_gbb - _aeeb._afeg), Y1: _eef, X2: _fee + _acfa/2 - _dcef + _aeeb._dgf - (_gfa - _aeeb._fe), Y2: _eef}
		_adg, _, _fdaa := _eeda.Draw("")
		if _fdaa != nil {
			return nil, ctx, _fdaa
		}
		_fdaa = _fbc.addContentsByString(string(_adg))
		if _fdaa != nil {
			return nil, ctx, _fdaa
		}
	}
	if _aeeb._gbe != 0 {
		_ffc := _gfdg
		_fced := _fce - _aeeb._gbgc
		if _aeeb._aeba == CellBorderStyleDouble {
			_fced += _fea
			_eeg := _cec.BasicLine{LineWidth: _aeeb._gbe, Opacity: 1.0, LineColor: _bcfe(_aeeb._bdd), LineStyle: _aeeb.LineStyle, X1: _ffc - _gbgg/2 + _fef, Y1: _fced - 2*_fea, X2: _ffc + _gbgg/2 - _fag + _aeeb._dgf, Y2: _fced - 2*_fea}
			_deb, _, _geb := _eeg.Draw("")
			if _geb != nil {
				return nil, ctx, _geb
			}
			_geb = _fbc.addContentsByString(string(_deb))
			if _geb != nil {
				return nil, ctx, _geb
			}
		}
		_ddg := _cec.BasicLine{LineWidth: _aeeb._gbe, Opacity: 1.0, LineColor: _bcfe(_aeeb._bdd), LineStyle: _aeeb.LineStyle, X1: _ffc - _gbgg/2 + _fef + (_gbb - _aeeb._afeg), Y1: _fced, X2: _ffc + _gbgg/2 - _fag + _aeeb._dgf - (_gfa - _aeeb._fe), Y2: _fced}
		_gdea, _, _gcf := _ddg.Draw("")
		if _gcf != nil {
			return nil, ctx, _gcf
		}
		_gcf = _fbc.addContentsByString(string(_gdea))
		if _gcf != nil {
			return nil, ctx, _gcf
		}
	}
	if _aeeb._afeg != 0 {
		_bcdc := _gfdg
		_gabc := _fce
		if _aeeb._afg == CellBorderStyleDouble {
			_bcdc += _dcf
			_gfe := _cec.BasicLine{LineWidth: _aeeb._afeg, Opacity: 1.0, LineColor: _bcfe(_aeeb._bfe), LineStyle: _aeeb.LineStyle, X1: _bcdc - 2*_dcf, Y1: _gabc + _gbb/2 + _bfb, X2: _bcdc - 2*_dcf, Y2: _gabc - _gbb/2 - _fef - _aeeb._gbgc}
			_bda, _, _ddd := _gfe.Draw("")
			if _ddd != nil {
				return nil, ctx, _ddd
			}
			_ddd = _fbc.addContentsByString(string(_bda))
			if _ddd != nil {
				return nil, ctx, _ddd
			}
		}
		_bff := _cec.BasicLine{LineWidth: _aeeb._afeg, Opacity: 1.0, LineColor: _bcfe(_aeeb._bfe), LineStyle: _aeeb.LineStyle, X1: _bcdc, Y1: _gabc + _gbb/2 + _bfb - (_acfa - _aeeb._ebc), X2: _bcdc, Y2: _gabc - _gbb/2 - _fef - _aeeb._gbgc + (_gbgg - _aeeb._gbe)}
		_eff, _, _fec := _bff.Draw("")
		if _fec != nil {
			return nil, ctx, _fec
		}
		_fec = _fbc.addContentsByString(string(_eff))
		if _fec != nil {
			return nil, ctx, _fec
		}
	}
	if _aeeb._fe != 0 {
		_debg := _gfdg + _aeeb._dgf
		_fdg := _fce
		if _aeeb._bgad == CellBorderStyleDouble {
			_debg -= _bgbg
			_fgg := _cec.BasicLine{LineWidth: _aeeb._fe, Opacity: 1.0, LineColor: _bcfe(_aeeb._bfae), LineStyle: _aeeb.LineStyle, X1: _debg + 2*_bgbg, Y1: _fdg + _gfa/2 + _dcef, X2: _debg + 2*_bgbg, Y2: _fdg - _gfa/2 - _fag - _aeeb._gbgc}
			_aef, _, _eaf := _fgg.Draw("")
			if _eaf != nil {
				return nil, ctx, _eaf
			}
			_eaf = _fbc.addContentsByString(string(_aef))
			if _eaf != nil {
				return nil, ctx, _eaf
			}
		}
		_ddce := _cec.BasicLine{LineWidth: _aeeb._fe, Opacity: 1.0, LineColor: _bcfe(_aeeb._bfae), LineStyle: _aeeb.LineStyle, X1: _debg, Y1: _fdg + _gfa/2 + _dcef - (_acfa - _aeeb._ebc), X2: _debg, Y2: _fdg - _gfa/2 - _fag - _aeeb._gbgc + (_gbgg - _aeeb._gbe)}
		_dgd, _, _ffcb := _ddce.Draw("")
		if _ffcb != nil {
			return nil, ctx, _ffcb
		}
		_ffcb = _fbc.addContentsByString(string(_dgd))
		if _ffcb != nil {
			return nil, ctx, _ffcb
		}
	}
	return []*Block{_fbc}, ctx, nil
}

// MultiCell makes a new cell with the specified row span and col span
// and inserts it into the table at the current position.
func (_ggcgd *Table) MultiCell(rowspan, colspan int) *TableCell {
	_ggcgd._edag++
	_dbadf := (_ggcgd.moveToNextAvailableCell()-1)%(_ggcgd._cfee) + 1
	_dbbc := (_ggcgd._edag-1)/_ggcgd._cfee + 1
	for _dbbc > _ggcgd._fdff {
		_ggcgd._fdff++
		_ggcgd._ceeb = append(_ggcgd._ceeb, _ggcgd._dcgce)
	}
	_cfeg := &TableCell{}
	_cfeg._gbdd = _dbbc
	_cfeg._dbdf = _dbadf
	_cfeg._cdgg = 5
	_cfeg._dbd = CellBorderStyleNone
	_cfeg._degeb = _cec.LineStyleSolid
	_cfeg._eaeac = CellHorizontalAlignmentLeft
	_cfeg._geac = CellVerticalAlignmentTop
	_cfeg._bfbdc = 0
	_cfeg._efgc = 0
	_cfeg._bfgd = 0
	_cfeg._cfdc = 0
	_bgda := ColorBlack
	_cfeg._egcc = _bgda
	_cfeg._dbee = _bgda
	_cfeg._gfag = _bgda
	_cfeg._ddec = _bgda
	if rowspan < 1 {
		_df.Log.Debug("\u0054\u0061\u0062\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0072\u006f\u0077\u0073\u0070a\u006e\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061t\u0020\u0031\u0020\u0028\u0025\u0064\u0029\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063e\u006c\u006c\u0020\u0072\u006f\u0077s\u0070\u0061n\u0020\u0074o\u00201\u002e", rowspan)
		rowspan = 1
	}
	_aaad := _ggcgd._fdff - (_cfeg._gbdd - 1)
	if rowspan > _aaad {
		_df.Log.Debug("\u0054\u0061b\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0072\u006f\u0077\u0073\u0070\u0061\u006e\u0020\u0028\u0025d\u0029\u0020\u0065\u0078\u0063\u0065e\u0064\u0073\u0020\u0072\u0065\u006d\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0072o\u0077\u0073 \u0028\u0025\u0064\u0029.\u0020\u0041\u0064\u0064\u0069n\u0067\u0020\u0072\u006f\u0077\u0073\u002e", rowspan, _aaad)
		_ggcgd._fdff += rowspan - 1
		for _eaag := 0; _eaag <= rowspan-_aaad; _eaag++ {
			_ggcgd._ceeb = append(_ggcgd._ceeb, _ggcgd._dcgce)
		}
	}
	for _dadd := 0; _dadd < colspan && _dbadf+_dadd-1 < len(_ggcgd._fbca); _dadd++ {
		_ggcgd._fbca[_dbadf+_dadd-1] = rowspan - 1
	}
	_cfeg._fefef = rowspan
	if colspan < 1 {
		_df.Log.Debug("\u0054\u0061\u0062\u006c\u0065\u003a\u0020\u0063\u0065\u006c\u006c\u0020\u0063\u006f\u006c\u0073\u0070a\u006e\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061n\u0020\u0031\u0020\u0028\u0025\u0064\u0029\u002e\u0020\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0063e\u006c\u006c\u0020\u0063\u006f\u006cs\u0070\u0061n\u0020\u0074o\u00201\u002e", colspan)
		colspan = 1
	}
	_fddg := _ggcgd._cfee - (_cfeg._dbdf - 1)
	if colspan > _fddg {
		_df.Log.Debug("\u0054\u0061\u0062\u006c\u0065:\u0020\u0063\u0065\u006c\u006c\u0020\u0063o\u006c\u0073\u0070\u0061\u006e\u0020\u0028\u0025\u0064\u0029\u0020\u0065\u0078\u0063\u0065\u0065\u0064\u0073\u0020\u0072\u0065\u006d\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0072\u006f\u0077\u0020\u0063\u006f\u006c\u0073\u0020\u0028\u0025d\u0029\u002e\u0020\u0041\u0064\u006a\u0075\u0073\u0074\u0069\u006e\u0067 \u0063\u006f\u006c\u0073\u0070\u0061n\u002e", colspan, _fddg)
		colspan = _fddg
	}
	_cfeg._cegef = colspan
	_ggcgd._edag += colspan - 1
	_ggcgd._ebfb = append(_ggcgd._ebfb, _cfeg)
	_cfeg._fbdc = _ggcgd
	return _cfeg
}

// NewImage create a new image from a unidoc image (model.Image).
func (_dbbb *Creator) NewImage(img *_f.Image) (*Image, error) { return _baec(img) }

// SetBorderOpacity sets the border opacity.
func (_aaeg *PolyBezierCurve) SetBorderOpacity(opacity float64) { _aaeg._bcbe = opacity }

// SetFillOpacity sets the fill opacity.
func (_fagc *Polygon) SetFillOpacity(opacity float64) { _fagc._fead = opacity }
func _gaef(_fbecd *templateProcessor, _gcda *templateNode) (interface{}, error) {
	return _fbecd.parseLine(_gcda)
}

// SetWidthRight sets border width for right.
func (_agd *border) SetWidthRight(bw float64) { _agd._fe = bw }

// Indent returns the left offset of the list when nested into another list.
func (_dfbec *List) Indent() float64 { return _dfbec._dgab }

// Width returns the width of the Paragraph.
func (_eccc *Paragraph) Width() float64 {
	if _eccc._bbfc && int(_eccc._bdga) > 0 {
		return _eccc._bdga
	}
	return _eccc.getTextWidth() / 1000.0
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
	_face *StyledParagraph

	// Holds the text and style of the number part of the TOC line.
	Number TextChunk

	// Holds the text and style of the title part of the TOC line.
	Title TextChunk

	// Holds the text and style of the separator part of the TOC line.
	Separator TextChunk

	// Holds the text and style of the page part of the TOC line.
	Page    TextChunk
	_ecdbc  float64
	_dgcda  uint
	_aaaa   float64
	_dcafec Positioning
	_cegge  float64
	_gaace  float64
	_cbfg   int64
}

// Number returns the invoice number description and value cells.
// The returned values can be used to customize the styles of the cells.
func (_bdbe *Invoice) Number() (*InvoiceCell, *InvoiceCell) { return _bdbe._edeb[0], _bdbe._edeb[1] }

// AddPage adds the specified page to the creator.
// NOTE: If the page has a Rotate flag, the creator will take care of
// transforming the contents to maintain the correct orientation.
func (_aecd *Creator) AddPage(page *_f.PdfPage) error {
	_bdfc, _gbgd := page.GetMediaBox()
	if _gbgd != nil {
		_df.Log.Debug("\u0046\u0061\u0069l\u0065\u0064\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0070\u0061g\u0065\u0020\u006d\u0065\u0064\u0069\u0061\u0062\u006f\u0078\u003a\u0020\u0025\u0076", _gbgd)
		return _gbgd
	}
	_bdfc.Normalize()
	_geae, _fada := _bdfc.Llx, _bdfc.Lly
	_dcac := _bdfc
	if _cefe := page.CropBox; _cefe != nil && *_cefe != *_bdfc {
		_cefe.Normalize()
		_geae, _fada = _cefe.Llx, _cefe.Lly
		_dcac = _cefe
	}
	_abgd := _edf.IdentityMatrix()
	_fdd, _gbgd := page.GetRotate()
	if _gbgd != nil {
		_df.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a", _gbgd.Error())
	}
	_fagd := _fdd%360 != 0 && _fdd%90 == 0
	if _fagd {
		_fbeeb := float64((360 + _fdd%360) % 360)
		if _fbeeb == 90 {
			_abgd = _abgd.Translate(_dcac.Width(), 0)
		} else if _fbeeb == 180 {
			_abgd = _abgd.Translate(_dcac.Width(), _dcac.Height())
		} else if _fbeeb == 270 {
			_abgd = _abgd.Translate(0, _dcac.Height())
		}
		_abgd = _abgd.Mult(_edf.RotationMatrix(_fbeeb * _a.Pi / 180))
		_abgd = _abgd.Round(0.000001)
		_ecd := _fegf(_dcac, _abgd)
		_dcac = _ecd
		_dcac.Normalize()
	}
	if _geae != 0 || _fada != 0 {
		_abgd = _edf.TranslationMatrix(_geae, _fada).Mult(_abgd)
	}
	if !_abgd.Identity() {
		_abgd = _abgd.Round(0.000001)
		_aecd._gba[page] = &pageTransformations{_daa: &_abgd}
	}
	_aecd._gee = _dcac.Width()
	_aecd._agca = _dcac.Height()
	_aecd.initContext()
	_aecd._gadb = append(_aecd._gadb, page)
	_aecd._bfec.Page++
	return nil
}
func (_caffb *templateProcessor) parseImage(_decc *templateNode) (interface{}, error) {
	var _fdeb string
	for _, _defc := range _decc._ddgde.Attr {
		_dgac := _defc.Value
		switch _gage := _defc.Name.Local; _gage {
		case "\u0073\u0072\u0063":
			_fdeb = _dgac
		}
	}
	_adae, _ggdcd := _caffb.loadImageFromSrc(_fdeb)
	if _ggdcd != nil {
		return nil, _ggdcd
	}
	for _, _effea := range _decc._ddgde.Attr {
		_ffcc := _effea.Value
		switch _cabcb := _effea.Name.Local; _cabcb {
		case "\u0061\u006c\u0069g\u006e":
			_adae.SetHorizontalAlignment(_caffb.parseHorizontalAlignmentAttr(_cabcb, _ffcc))
		case "\u006fp\u0061\u0063\u0069\u0074\u0079":
			_adae.SetOpacity(_caffb.parseFloatAttr(_cabcb, _ffcc))
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			_aaefb := _caffb.parseMarginAttr(_cabcb, _ffcc)
			_adae.SetMargins(_aaefb.Left, _aaefb.Right, _aaefb.Top, _aaefb.Bottom)
		case "\u0066\u0069\u0074\u002d\u006d\u006f\u0064\u0065":
			_adae.SetFitMode(_caffb.parseFitModeAttr(_cabcb, _ffcc))
		case "\u0078":
			_adae.SetPos(_caffb.parseFloatAttr(_cabcb, _ffcc), _adae._ddab)
		case "\u0079":
			_adae.SetPos(_adae._cgee, _caffb.parseFloatAttr(_cabcb, _ffcc))
		case "\u0077\u0069\u0064t\u0068":
			_adae.SetWidth(_caffb.parseFloatAttr(_cabcb, _ffcc))
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_adae.SetHeight(_caffb.parseFloatAttr(_cabcb, _ffcc))
		case "\u0061\u006e\u0067l\u0065":
			_adae.SetAngle(_caffb.parseFloatAttr(_cabcb, _ffcc))
		case "\u0073\u0072\u0063":
		default:
			_df.Log.Debug("\u0055n\u0073\u0075p\u0070\u006f\u0072\u0074e\u0064\u0020\u0069m\u0061\u0067\u0065\u0020\u0061\u0074\u0074\u0072\u0069bu\u0074\u0065\u003a \u0060\u0025s\u0060\u002e\u0020\u0053\u006b\u0069p\u0070\u0069n\u0067\u002e", _cabcb)
		}
	}
	return _adae, nil
}

// Heading returns the heading component of the table of contents.
func (_gbbaa *TOC) Heading() *StyledParagraph { return _gbbaa._gcdb }

// Padding returns the padding of the component.
func (_ddeg *Division) Padding() (_bea, _deea, _cgf, _egdb float64) {
	return _ddeg._gfeb.Left, _ddeg._gfeb.Right, _ddeg._gfeb.Top, _ddeg._gfeb.Bottom
}

// NewParagraph creates a new text paragraph.
// Default attributes:
// Font: Helvetica,
// Font size: 10
// Encoding: WinAnsiEncoding
// Wrap: enabled
// Text color: black
func (_fgcgd *Creator) NewParagraph(text string) *Paragraph {
	return _ebcb(text, _fgcgd.NewTextStyle())
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

// ConvertToBinary converts current image data into binary (Bi-level image) format.
// If provided image is RGB or GrayScale the function converts it into binary image
// using histogram auto threshold method.
func (_gcgf *Image) ConvertToBinary() error { return _gcgf._aaca.ConvertToBinary() }
func _dadae(_ffaed *templateProcessor, _caed *templateNode) (interface{}, error) {
	return _ffaed.parseDivision(_caed)
}
func (_aefeg *templateProcessor) parseFitModeAttr(_dagb, _edaeag string) FitMode {
	_df.Log.Debug("\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u0069\u0074\u0020\u006do\u0064\u0065\u0020\u0061\u0074\u0074r\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0028\u0060\u0025\u0073\u0060\u002c \u0025\u0073\u0029\u002e", _dagb, _edaeag)
	_cada := map[string]FitMode{"\u006e\u006f\u006e\u0065": FitModeNone, "\u0066\u0069\u006c\u006c\u002d\u0077\u0069\u0064\u0074\u0068": FitModeFillWidth}[_edaeag]
	return _cada
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
func (_eedb *Creator) SetPageSize(size PageSize) {
	_eedb._dda = size
	_eedb._gee = size[0]
	_eedb._agca = size[1]
	_gbgca := 0.1 * _eedb._gee
	_eedb._fade.Left = _gbgca
	_eedb._fade.Right = _gbgca
	_eedb._fade.Top = _gbgca
	_eedb._fade.Bottom = _gbgca
}

// MultiColCell makes a new cell with the specified column span and inserts it
// into the table at the current position.
func (_bdaf *Table) MultiColCell(colspan int) *TableCell { return _bdaf.MultiCell(1, colspan) }

// SetHeight sets the Image's document height to specified h.
func (_ababae *Image) SetHeight(h float64) { _ababae._gddc = h }

// SetPos sets the position of the chart to the specified coordinates.
// This method sets the chart to use absolute positioning.
func (_bad *Chart) SetPos(x, y float64) { _bad._ebe = PositionAbsolute; _bad._eca = x; _bad._dbc = y }

// Height returns the height of the line.
func (_eaeaa *Line) Height() float64 {
	_becca := _eaeaa._ffgbg
	if _eaeaa._ggfd == _eaeaa._dccb {
		_becca /= 2
	}
	return _a.Abs(_eaeaa._gdec-_eaeaa._fadc) + _becca
}

// SetEnableWrap sets the line wrapping enabled flag.
func (_bffge *StyledParagraph) SetEnableWrap(enableWrap bool) {
	_bffge._ecbf = enableWrap
	_bffge._ffbb = false
}

// SetShowNumbering sets a flag to indicate whether or not to show chapter numbers as part of title.
func (_bfee *Chapter) SetShowNumbering(show bool) {
	_bfee._bdc = show
	_bfee._dbf.SetText(_bfee.headingText())
}

// Draw draws the drawable d on the block.
// Note that the drawable must not wrap, i.e. only return one block. Otherwise an error is returned.
func (_caag *Block) Draw(d Drawable) error {
	_bga := DrawContext{}
	_bga.Width = _caag._bf
	_bga.Height = _caag._eda
	_bga.PageWidth = _caag._bf
	_bga.PageHeight = _caag._eda
	_bga.X = 0
	_bga.Y = 0
	_gec, _, _ffe := d.GeneratePageBlocks(_bga)
	if _ffe != nil {
		return _ffe
	}
	if len(_gec) != 1 {
		return _d.New("\u0074\u006f\u006f\u0020ma\u006e\u0079\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0062\u006c\u006f\u0063k\u0073")
	}
	for _, _cf := range _gec {
		if _dcc := _caag.mergeBlocks(_cf); _dcc != nil {
			return _dcc
		}
	}
	return nil
}
func (_cdceb *Division) split(_bcbcd DrawContext) (_gdde, _bfeeg *Division) {
	var (
		_gbeba      float64
		_aff, _cedd []VectorDrawable
	)
	_ddgef := _bcbcd.Width - _cdceb._fccf.Left - _cdceb._fccf.Right - _cdceb._gfeb.Left - _cdceb._gfeb.Right
	for _fceg, _cgcb := range _cdceb._aefd {
		_gbeba += _dcge(_cgcb, _ddgef)
		if _gbeba < _bcbcd.Height {
			_aff = append(_aff, _cgcb)
		} else {
			_cedd = _cdceb._aefd[_fceg:]
			break
		}
	}
	if len(_aff) > 0 {
		_gdde = _cdde()
		*_gdde = *_cdceb
		_gdde._aefd = _aff
		if _cdceb._bgeg != nil {
			_gdde._bgeg = &Background{}
			*_gdde._bgeg = *_cdceb._bgeg
		}
	}
	if len(_cedd) > 0 {
		_bfeeg = _cdde()
		*_bfeeg = *_cdceb
		_bfeeg._aefd = _cedd
		if _cdceb._bgeg != nil {
			_bfeeg._bgeg = &Background{}
			*_bfeeg._bgeg = *_cdceb._bgeg
		}
	}
	return _gdde, _bfeeg
}
func _ffggc(_gbbfc, _fffgc, _bedbd TextChunk, _cdee uint, _eafae TextStyle) *TOCLine {
	_bcgcg := _ggaa(_eafae)
	_bcgcg.SetEnableWrap(true)
	_bcgcg.SetTextAlignment(TextAlignmentLeft)
	_bcgcg.SetMargins(0, 0, 2, 2)
	_egbf := &TOCLine{_face: _bcgcg, Number: _gbbfc, Title: _fffgc, Page: _bedbd, Separator: TextChunk{Text: "\u002e", Style: _eafae}, _ecdbc: 0, _dgcda: _cdee, _aaaa: 10, _dcafec: PositionRelative}
	_bcgcg._afdba.Left = _egbf._ecdbc + float64(_egbf._dgcda-1)*_egbf._aaaa
	_bcgcg._bccd = _egbf.prepareParagraph
	return _egbf
}
func (_egdgd *Line) computeCoords(_cdbc DrawContext) (_cggbd, _cefd, _ggdb, _afad float64) {
	_cggbd = _cdbc.X
	_ggdb = _cggbd + _egdgd._dccb - _egdgd._ggfd
	_edgd := _egdgd._ffgbg
	if _egdgd._ggfd == _egdgd._dccb {
		_edgd /= 2
	}
	if _egdgd._fadc < _egdgd._gdec {
		_cefd = _cdbc.PageHeight - _cdbc.Y - _edgd
		_afad = _cefd - _egdgd._gdec + _egdgd._fadc
	} else {
		_afad = _cdbc.PageHeight - _cdbc.Y - _edgd
		_cefd = _afad - _egdgd._fadc + _egdgd._gdec
	}
	switch _egdgd._gggcb {
	case FitModeFillWidth:
		_ggdb = _cggbd + _cdbc.Width
	}
	return _cggbd, _cefd, _ggdb, _afad
}

// SetBuyerAddress sets the buyer address of the invoice.
func (_aeebb *Invoice) SetBuyerAddress(address *InvoiceAddress) { _aeebb._ebca = address }

// SetColor sets the line color.
func (_ccfd *Curve) SetColor(col Color) { _ccfd._cddgg = col }

// SetFillColor sets the fill color for the path.
func (_cabg *FilledCurve) SetFillColor(color Color) { _cabg._bffg = color }

// NewPolygon creates a new polygon.
func (_dgcc *Creator) NewPolygon(points [][]_cec.Point) *Polygon { return _bfeg(points) }

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
func (_aea *Creator) Finalize() error {
	if _aea._cgga {
		return nil
	}
	_fffg := len(_aea._gadb)
	_aad := 0
	if _aea._bcec != nil {
		_eedbe := *_aea
		_aea._gadb = nil
		_aea._efde = nil
		_aea.initContext()
		_ecg := FrontpageFunctionArgs{PageNum: 1, TotalPages: _fffg}
		_aea._bcec(_ecg)
		_aad += len(_aea._gadb)
		_aea._gadb = _eedbe._gadb
		_aea._efde = _eedbe._efde
	}
	if _aea.AddTOC {
		_aea.initContext()
		_aea._bfec.Page = _aad + 1
		if _aea.CustomTOC && _aea._bdcf != nil {
			_gfcg := *_aea
			_aea._gadb = nil
			_aea._efde = nil
			if _feg := _aea._bdcf(_aea._cddg); _feg != nil {
				return _feg
			}
			_aad += len(_aea._gadb)
			_aea._gadb = _gfcg._gadb
			_aea._efde = _gfcg._efde
		} else {
			if _aea._bdcf != nil {
				if _bdcb := _aea._bdcf(_aea._cddg); _bdcb != nil {
					return _bdcb
				}
			}
			_bbc, _, _ggc := _aea._cddg.GeneratePageBlocks(_aea._bfec)
			if _ggc != nil {
				_df.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0065\u0020\u0062\u006c\u006f\u0063\u006b\u0073: \u0025\u0076", _ggc)
				return _ggc
			}
			_aad += len(_bbc)
		}
		_aagc := _aea._cddg.Lines()
		for _, _afdbg := range _aagc {
			_gfec, _afc := _ae.Atoi(_afdbg.Page.Text)
			if _afc != nil {
				continue
			}
			_afdbg.Page.Text = _ae.Itoa(_gfec + _aad)
			_afdbg._cbfg += int64(_aad)
		}
	}
	_agdc := false
	var _ddaa []*_f.PdfPage
	if _aea._bcec != nil {
		_abef := *_aea
		_aea._gadb = nil
		_aea._efde = nil
		_fbag := FrontpageFunctionArgs{PageNum: 1, TotalPages: _fffg}
		_aea._bcec(_fbag)
		_fffg += len(_aea._gadb)
		_ddaa = _aea._gadb
		_aea._gadb = append(_aea._gadb, _abef._gadb...)
		_aea._efde = _abef._efde
		_agdc = true
	}
	var _cfgg []*_f.PdfPage
	if _aea.AddTOC {
		_aea.initContext()
		if _aea.CustomTOC && _aea._bdcf != nil {
			_gcce := *_aea
			_aea._gadb = nil
			_aea._efde = nil
			if _abefe := _aea._bdcf(_aea._cddg); _abefe != nil {
				_df.Log.Debug("\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074i\u006e\u0067\u0020\u0054\u004f\u0043\u003a\u0020\u0025\u0076", _abefe)
				return _abefe
			}
			_cfgg = _aea._gadb
			_fffg += len(_cfgg)
			_aea._gadb = _gcce._gadb
			_aea._efde = _gcce._efde
		} else {
			if _aea._bdcf != nil {
				if _afdc := _aea._bdcf(_aea._cddg); _afdc != nil {
					_df.Log.Debug("\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074i\u006e\u0067\u0020\u0054\u004f\u0043\u003a\u0020\u0025\u0076", _afdc)
					return _afdc
				}
			}
			_dcad, _, _ := _aea._cddg.GeneratePageBlocks(_aea._bfec)
			for _, _defd := range _dcad {
				_defd.SetPos(0, 0)
				_fffg++
				_cffga := _aea.newPage()
				_cfgg = append(_cfgg, _cffga)
				_aea.setActivePage(_cffga)
				_aea.Draw(_defd)
			}
		}
		if _agdc {
			_gcdd := _ddaa
			_ecc := _aea._gadb[len(_ddaa):]
			_aea._gadb = append([]*_f.PdfPage{}, _gcdd...)
			_aea._gadb = append(_aea._gadb, _cfgg...)
			_aea._gadb = append(_aea._gadb, _ecc...)
		} else {
			_aea._gadb = append(_cfgg, _aea._gadb...)
		}
	}
	if _aea._dfd != nil && _aea.AddOutlines {
		var _dbac func(_gcddb *_f.OutlineItem)
		_dbac = func(_ccdg *_f.OutlineItem) {
			_ccdg.Dest.Page += int64(_aad)
			if _ecdb := int(_ccdg.Dest.Page); _ecdb >= 0 && _ecdb < len(_aea._gadb) {
				_ccdg.Dest.PageObj = _aea._gadb[_ecdb].GetPageAsIndirectObject()
			} else {
				_df.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _ecdb)
			}
			_ccdg.Dest.Y = _aea._agca - _ccdg.Dest.Y
			_dag := _ccdg.Items()
			for _, _ccff := range _dag {
				_dbac(_ccff)
			}
		}
		_gbea := _aea._dfd.Items()
		for _, _dge := range _gbea {
			_dbac(_dge)
		}
		if _aea.AddTOC {
			var _dfff int
			if _agdc {
				_dfff = len(_ddaa)
			}
			_ababd := _f.NewOutlineDest(int64(_dfff), 0, _aea._agca)
			if _dfff >= 0 && _dfff < len(_aea._gadb) {
				_ababd.PageObj = _aea._gadb[_dfff].GetPageAsIndirectObject()
			} else {
				_df.Log.Debug("\u0057\u0041R\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u0064", _dfff)
			}
			_aea._dfd.Insert(0, _f.NewOutlineItem("\u0054\u0061\u0062\u006c\u0065\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0065\u006e\u0074\u0073", _ababd))
		}
	}
	for _cggb, _cbdg := range _aea._gadb {
		_aea.setActivePage(_cbdg)
		if _aea._ebd != nil {
			_dacf, _gaa, _ddea := _cbdg.Size()
			if _ddea != nil {
				return _ddea
			}
			_eafc := PageFinalizeFunctionArgs{PageNum: _cggb + 1, PageWidth: _dacf, PageHeight: _gaa, TOCPages: len(_cfgg), TotalPages: _fffg}
			if _ecea := _aea._ebd(_eafc); _ecea != nil {
				_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0070\u0061\u0067\u0065\u0020\u0066\u0069\u006e\u0061\u006c\u0069\u007a\u0065 \u0063\u0061\u006c\u006c\u0062\u0061\u0063k\u003a\u0020\u0025\u0076", _ecea)
				return _ecea
			}
		}
		if _aea._eacd != nil {
			_caee := NewBlock(_aea._gee, _aea._fade.Top)
			_efdb := HeaderFunctionArgs{PageNum: _cggb + 1, TotalPages: _fffg}
			_aea._eacd(_caee, _efdb)
			_caee.SetPos(0, 0)
			if _fcga := _aea.Draw(_caee); _fcga != nil {
				_df.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069n\u0067 \u0068e\u0061\u0064\u0065\u0072\u003a\u0020\u0025v", _fcga)
				return _fcga
			}
		}
		if _aea._eeb != nil {
			_dbfb := NewBlock(_aea._gee, _aea._fade.Bottom)
			_bfcgf := FooterFunctionArgs{PageNum: _cggb + 1, TotalPages: _fffg}
			_aea._eeb(_dbfb, _bfcgf)
			_dbfb.SetPos(0, _aea._agca-_dbfb._eda)
			if _feea := _aea.Draw(_dbfb); _feea != nil {
				_df.Log.Debug("\u0045R\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069n\u0067 \u0066o\u006f\u0074\u0065\u0072\u003a\u0020\u0025v", _feea)
				return _feea
			}
		}
		_daaa, _fgga := _aea._gba[_cbdg]
		if _abfe, _cagf := _aea._caec[_cbdg]; _cagf {
			if _fgga {
				_daaa.transformBlock(_abfe)
			}
			if _ecec := _abfe.drawToPage(_cbdg); _ecec != nil {
				_df.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0064\u0072\u0061\u0077\u0069\u006e\u0067\u0020\u0070\u0061\u0067\u0065\u0020%\u0064\u0020\u0062\u006c\u006f\u0063\u006bs\u003a\u0020\u0025\u0076", _cggb+1, _ecec)
				return _ecec
			}
		}
		if _fgga {
			if _ffa := _daaa.transformPage(_cbdg); _ffa != nil {
				_df.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0074\u0072\u0061\u006e\u0073f\u006f\u0072\u006d\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0076", _ffa)
				return _ffa
			}
		}
	}
	_aea._cgga = true
	return nil
}

// SetColorTop sets border color for top.
func (_bcf *border) SetColorTop(col Color) { _bcf._gad = col }

// SetMargins sets the margins of the component. The margins are applied
// around the division.
func (_ebfa *Division) SetMargins(left, right, top, bottom float64) {
	_ebfa._fccf.Left = left
	_ebfa._fccf.Right = right
	_ebfa._fccf.Top = top
	_ebfa._fccf.Bottom = bottom
}

// SetBackground sets the background properties of the component.
func (_gcdf *Division) SetBackground(background *Background) { _gcdf._bgeg = background }

// SetColorBottom sets border color for bottom.
func (_begc *border) SetColorBottom(col Color) { _begc._bdd = col }
func _eabe(_fbfa *_e.File) ([]*_f.PdfPage, error) {
	_eagf, _dbbde := _f.NewPdfReader(_fbfa)
	if _dbbde != nil {
		return nil, _dbbde
	}
	_eceeg, _dbbde := _eagf.GetNumPages()
	if _dbbde != nil {
		return nil, _dbbde
	}
	var _bffdc []*_f.PdfPage
	for _ffdg := 0; _ffdg < _eceeg; _ffdg++ {
		_cbcbfa, _gaag := _eagf.GetPage(_ffdg + 1)
		if _gaag != nil {
			return nil, _gaag
		}
		_bffdc = append(_bffdc, _cbcbfa)
	}
	return _bffdc, nil
}
func _fgcgae(_bagg []_cec.CubicBezierCurve) *PolyBezierCurve {
	return &PolyBezierCurve{_cegf: &_cec.PolyBezierCurve{Curves: _bagg, BorderColor: _f.NewPdfColorDeviceRGB(0, 0, 0), BorderWidth: 1.0}, _bbeag: 1.0, _bcbe: 1.0}
}

// SetPos sets absolute positioning with specified coordinates.
func (_ceff *Paragraph) SetPos(x, y float64) {
	_ceff._beeba = PositionAbsolute
	_ceff._fbbde = x
	_ceff._agdg = y
}

// SetPositioning sets the positioning of the line (absolute or relative).
func (_gfeac *Line) SetPositioning(positioning Positioning) { _gfeac._edcee = positioning }
func _dadda(_cgcge *_f.PdfFont) TextStyle {
	return TextStyle{Color: ColorRGBFrom8bit(0, 0, 0), Font: _cgcge, FontSize: 10, OutlineSize: 1, HorizontalScaling: DefaultHorizontalScaling, UnderlineStyle: TextDecorationLineStyle{Offset: 1, Thickness: 1}}
}
func (_fdbee *TOCLine) getLineLink() *_f.PdfAnnotation {
	if _fdbee._cbfg <= 0 {
		return nil
	}
	return _eabf(_fdbee._cbfg-1, _fdbee._cegge, _fdbee._gaace, 0)
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
func (_fgde *Division) Add(d VectorDrawable) error {
	switch _bfdb := d.(type) {
	case *Paragraph, *StyledParagraph, *Image, *Chart, *Rectangle, *Ellipse, *Line, *Table, *Division:
	case containerDrawable:
		_fdbb, _eecd := _bfdb.ContainerComponent(_fgde)
		if _eecd != nil {
			return _eecd
		}
		_dcca, _fgcga := _fdbb.(VectorDrawable)
		if !_fgcga {
			return _eg.Errorf("\u0072\u0065\u0073\u0075\u006ct\u0020\u006f\u0066\u0020\u0043\u006f\u006et\u0061\u0069\u006e\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u002d\u0020\u0025\u0054\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0056\u0065c\u0074\u006f\u0072\u0044\u0072\u0061\u0077\u0061\u0062\u006c\u0065\u0020i\u006e\u0074\u0065\u0072\u0066\u0061c\u0065", _fdbb)
		}
		d = _dcca
	default:
		return _d.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070e\u0020i\u006e\u0020\u0044\u0069\u0076\u0069\u0073i\u006f\u006e")
	}
	_fgde._aefd = append(_fgde._aefd, d)
	return nil
}

// SetAngle sets the rotation angle in degrees.
func (_agc *Block) SetAngle(angleDeg float64) { _agc._ag = angleDeg }

// GeneratePageBlocks draws the polyline on a new block representing the page.
// Implements the Drawable interface.
func (_gbgfe *Polyline) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	_cedbc := NewBlock(ctx.PageWidth, ctx.PageHeight)
	_ffac, _egbd := _cedbc.setOpacity(_gbgfe._afbcd, _gbgfe._afbcd)
	if _egbd != nil {
		return nil, ctx, _egbd
	}
	_bbgd := _gbgfe._bcg.Points
	for _cfdg := range _bbgd {
		_fgca := &_bbgd[_cfdg]
		_fgca.Y = ctx.PageHeight - _fgca.Y
	}
	_gede, _, _egbd := _gbgfe._bcg.Draw(_ffac)
	if _egbd != nil {
		return nil, ctx, _egbd
	}
	if _egbd = _cedbc.addContentsByString(string(_gede)); _egbd != nil {
		return nil, ctx, _egbd
	}
	return []*Block{_cedbc}, ctx, nil
}

// Rectangle defines a rectangle with upper left corner at (x,y) and a specified width and height.  The rectangle
// can have a colored fill and/or border with a specified width.
// Implements the Drawable interface and can be drawn on PDF using the Creator.
type Rectangle struct {
	_bagb  float64
	_cbe   float64
	_ecfc  float64
	_ecca  float64
	_fcbgb Positioning
	_aebg  Color
	_ffcbg float64
	_afcb  Color
	_bffaf float64
	_beeac float64
	_ecdc  float64
	_aaegc float64
	_bbde  float64
	_dada  float64
}
