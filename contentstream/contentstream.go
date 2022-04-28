package contentstream

import (
	_cc "bufio"
	_gc "bytes"
	_f "encoding/hex"
	_a "errors"
	_fc "fmt"
	_g "image/color"
	_ac "image/jpeg"
	_ccg "io"
	_ec "math"
	_c "strconv"

	_eb "bitbucket.org/shenghui0779/gopdf/common"
	_d "bitbucket.org/shenghui0779/gopdf/core"
	_af "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_cf "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_cg "bitbucket.org/shenghui0779/gopdf/model"
)

// Add_Tstar appends 'T*' operand to the content stream:
// Move to the start of next line.
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_cgd *ContentCreator) Add_Tstar() *ContentCreator {
	_dfd := ContentStreamOperation{}
	_dfd.Operand = "\u0054\u002a"
	_cgd._acd = append(_cgd._acd, &_dfd)
	return _cgd
}

// Add_q adds 'q' operand to the content stream: Pushes the current graphics state on the stack.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_cbbc *ContentCreator) Add_q() *ContentCreator {
	_dbg := ContentStreamOperation{}
	_dbg.Operand = "\u0071"
	_cbbc._acd = append(_cbbc._acd, &_dbg)
	return _cbbc
}

// GetEncoder returns the encoder of the inline image.
func (_aebb *ContentStreamInlineImage) GetEncoder() (_d.StreamEncoder, error) { return _gcf(_aebb) }

// Add_i adds 'i' operand to the content stream: Set the flatness tolerance in the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_efd *ContentCreator) Add_i(flatness float64) *ContentCreator {
	_ede := ContentStreamOperation{}
	_ede.Operand = "\u0069"
	_ede.Params = _gbfad([]float64{flatness})
	_efd._acd = append(_efd._acd, &_ede)
	return _efd
}

// IsMask checks if an image is a mask.
// The image mask entry in the image dictionary specifies that the image data shall be used as a stencil
// mask for painting in the current color. The mask data is 1bpc, grayscale.
func (_gcde *ContentStreamInlineImage) IsMask() (bool, error) {
	if _gcde.ImageMask != nil {
		_cega, _ddbg := _gcde.ImageMask.(*_d.PdfObjectBool)
		if !_ddbg {
			_eb.Log.Debug("\u0049m\u0061\u0067\u0065\u0020\u006d\u0061\u0073\u006b\u0020\u006e\u006ft\u0020\u0061\u0020\u0062\u006f\u006f\u006c\u0065\u0061\u006e")
			return false, _a.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065")
		}
		return bool(*_cega), nil
	}
	return false, nil
}

// Add_b appends 'b' operand to the content stream:
// Close, fill and then stroke the path (nonzero winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_ddb *ContentCreator) Add_b() *ContentCreator {
	_aeb := ContentStreamOperation{}
	_aeb.Operand = "\u0062"
	_ddb._acd = append(_ddb._acd, &_aeb)
	return _ddb
}

// Add_Tc appends 'Tc' operand to the content stream:
// Set character spacing.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_ccf *ContentCreator) Add_Tc(charSpace float64) *ContentCreator {
	_fff := ContentStreamOperation{}
	_fff.Operand = "\u0054\u0063"
	_fff.Params = _gbfad([]float64{charSpace})
	_ccf._acd = append(_ccf._acd, &_fff)
	return _ccf
}

// Add_SCN appends 'SCN' operand to the content stream:
// Same as SC but supports more colorspaces.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_ffe *ContentCreator) Add_SCN(c ...float64) *ContentCreator {
	_aebf := ContentStreamOperation{}
	_aebf.Operand = "\u0053\u0043\u004e"
	_aebf.Params = _gbfad(c)
	_ffe._acd = append(_ffe._acd, &_aebf)
	return _ffe
}

// Add_Tf appends 'Tf' operand to the content stream:
// Set font and font size specified by font resource `fontName` and `fontSize`.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_adg *ContentCreator) Add_Tf(fontName _d.PdfObjectName, fontSize float64) *ContentCreator {
	_afc := ContentStreamOperation{}
	_afc.Operand = "\u0054\u0066"
	_afc.Params = _befb([]_d.PdfObjectName{fontName})
	_afc.Params = append(_afc.Params, _gbfad([]float64{fontSize})...)
	_adg._acd = append(_adg._acd, &_afc)
	return _adg
}

// Add_SCN_pattern appends 'SCN' operand to the content stream for pattern `name`:
// SCN with name attribute (for pattern). Syntax: c1 ... cn name SCN.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_fbf *ContentCreator) Add_SCN_pattern(name _d.PdfObjectName, c ...float64) *ContentCreator {
	_abc := ContentStreamOperation{}
	_abc.Operand = "\u0053\u0043\u004e"
	_abc.Params = _gbfad(c)
	_abc.Params = append(_abc.Params, _d.MakeName(string(name)))
	_fbf._acd = append(_fbf._acd, &_abc)
	return _fbf
}

// RotateDeg applies a rotation to the transformation matrix.
func (_fbe *ContentCreator) RotateDeg(angle float64) *ContentCreator {
	_bdf := _ec.Cos(angle * _ec.Pi / 180.0)
	_dcd := _ec.Sin(angle * _ec.Pi / 180.0)
	_ccc := -_ec.Sin(angle * _ec.Pi / 180.0)
	_efc := _ec.Cos(angle * _ec.Pi / 180.0)
	return _fbe.Add_cm(_bdf, _dcd, _ccc, _efc, 0, 0)
}

// Add_W_starred appends 'W*' operand to the content stream:
// Modify the current clipping path by intersecting with the current path (even odd rule).
//
// See section 8.5.4 "Clipping Path Operators" and Table 61 (p. 146 PDF32000_2008).
func (_deg *ContentCreator) Add_W_starred() *ContentCreator {
	_bab := ContentStreamOperation{}
	_bab.Operand = "\u0057\u002a"
	_deg._acd = append(_deg._acd, &_bab)
	return _deg
}

// Add_b_starred appends 'b*' operand to the content stream:
// Close, fill and then stroke the path (even-odd winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_gaa *ContentCreator) Add_b_starred() *ContentCreator {
	_aebd := ContentStreamOperation{}
	_aebd.Operand = "\u0062\u002a"
	_gaa._acd = append(_gaa._acd, &_aebd)
	return _gaa
}

// GetColorSpace returns the colorspace of the inline image.
func (_gbg *ContentStreamInlineImage) GetColorSpace(resources *_cg.PdfPageResources) (_cg.PdfColorspace, error) {
	if _gbg.ColorSpace == nil {
		_eb.Log.Debug("\u0049\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076i\u006e\u0067\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u002c\u0020\u0061\u0073\u0073\u0075\u006di\u006e\u0067\u0020\u0047\u0072a\u0079")
		return _cg.NewPdfColorspaceDeviceGray(), nil
	}
	if _cdcca, _accf := _gbg.ColorSpace.(*_d.PdfObjectArray); _accf {
		return _agee(_cdcca)
	}
	_edfa, _ceed := _gbg.ColorSpace.(*_d.PdfObjectName)
	if !_ceed {
		_eb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020I\u006e\u0076\u0061\u006c\u0069\u0064 \u006f\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u003b\u0025\u002bv\u0029", _gbg.ColorSpace, _gbg.ColorSpace)
		return nil, _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_edfa == "\u0047" || *_edfa == "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079" {
		return _cg.NewPdfColorspaceDeviceGray(), nil
	} else if *_edfa == "\u0052\u0047\u0042" || *_edfa == "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B" {
		return _cg.NewPdfColorspaceDeviceRGB(), nil
	} else if *_edfa == "\u0043\u004d\u0059\u004b" || *_edfa == "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b" {
		return _cg.NewPdfColorspaceDeviceCMYK(), nil
	} else if *_edfa == "\u0049" || *_edfa == "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		return nil, _a.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0049\u006e\u0064e\u0078 \u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065")
	} else {
		if resources.ColorSpace == nil {
			_eb.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0073", *_edfa)
			return nil, _a.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0063o\u006c\u006fr\u0073\u0070\u0061\u0063\u0065")
		}
		_dagg, _fbc := resources.GetColorspaceByName(*_edfa)
		if !_fbc {
			_eb.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0073", *_edfa)
			return nil, _a.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0063o\u006c\u006fr\u0073\u0070\u0061\u0063\u0065")
		}
		return _dagg, nil
	}
}

// Operations returns the list of operations.
func (_de *ContentCreator) Operations() *ContentStreamOperations { return &_de._acd }

// Add_W appends 'W' operand to the content stream:
// Modify the current clipping path by intersecting with the current path (nonzero winding rule).
//
// See section 8.5.4 "Clipping Path Operators" and Table 61 (p. 146 PDF32000_2008).
func (_dcg *ContentCreator) Add_W() *ContentCreator {
	_bca := ContentStreamOperation{}
	_bca.Operand = "\u0057"
	_dcg._acd = append(_dcg._acd, &_bca)
	return _dcg
}

// Add_c adds 'c' operand to the content stream: Append a Bezier curve to the current path from
// the current point to (x3,y3) with (x1,x1) and (x2,y2) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_cce *ContentCreator) Add_c(x1, y1, x2, y2, x3, y3 float64) *ContentCreator {
	_afe := ContentStreamOperation{}
	_afe.Operand = "\u0063"
	_afe.Params = _gbfad([]float64{x1, y1, x2, y2, x3, y3})
	_cce._acd = append(_cce._acd, &_afe)
	return _cce
}

// Add_Do adds 'Do' operation to the content stream:
// Displays an XObject (image or form) specified by `name`.
//
// See section 8.8 "External Objects" and Table 87 (pp. 209-220 PDF32000_2008).
func (_ffd *ContentCreator) Add_Do(name _d.PdfObjectName) *ContentCreator {
	_ddd := ContentStreamOperation{}
	_ddd.Operand = "\u0044\u006f"
	_ddd.Params = _befb([]_d.PdfObjectName{name})
	_ffd._acd = append(_ffd._acd, &_ddd)
	return _ffd
}

// Add_G appends 'G' operand to the content stream:
// Set the stroking colorspace to DeviceGray and sets the gray level (0-1).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_agb *ContentCreator) Add_G(gray float64) *ContentCreator {
	_aef := ContentStreamOperation{}
	_aef.Operand = "\u0047"
	_aef.Params = _gbfad([]float64{gray})
	_agb._acd = append(_agb._acd, &_aef)
	return _agb
}

// Add_RG appends 'RG' operand to the content stream:
// Set the stroking colorspace to DeviceRGB and sets the r,g,b colors (0-1 each).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_geb *ContentCreator) Add_RG(r, g, b float64) *ContentCreator {
	_bbad := ContentStreamOperation{}
	_bbad.Operand = "\u0052\u0047"
	_bbad.Params = _gbfad([]float64{r, g, b})
	_geb._acd = append(_geb._acd, &_bbad)
	return _geb
}

var (
	ErrInvalidOperand = _a.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
)

// Add_scn appends 'scn' operand to the content stream:
// Same as SC but for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_efb *ContentCreator) Add_scn(c ...float64) *ContentCreator {
	_cgf := ContentStreamOperation{}
	_cgf.Operand = "\u0073\u0063\u006e"
	_cgf.Params = _gbfad(c)
	_efb._acd = append(_efb._acd, &_cgf)
	return _efb
}
func (_cgae *ContentStreamParser) parseBool() (_d.PdfObjectBool, error) {
	_caec, _agff := _cgae._cecb.Peek(4)
	if _agff != nil {
		return _d.PdfObjectBool(false), _agff
	}
	if (len(_caec) >= 4) && (string(_caec[:4]) == "\u0074\u0072\u0075\u0065") {
		_cgae._cecb.Discard(4)
		return _d.PdfObjectBool(true), nil
	}
	_caec, _agff = _cgae._cecb.Peek(5)
	if _agff != nil {
		return _d.PdfObjectBool(false), _agff
	}
	if (len(_caec) >= 5) && (string(_caec[:5]) == "\u0066\u0061\u006cs\u0065") {
		_cgae._cecb.Discard(5)
		return _d.PdfObjectBool(false), nil
	}
	return _d.PdfObjectBool(false), _a.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

// Transform returns coordinates x, y transformed by the CTM.
func (_dbgd *GraphicsState) Transform(x, y float64) (float64, float64) {
	return _dbgd.CTM.Transform(x, y)
}

// Add_Tw appends 'Tw' operand to the content stream:
// Set word spacing.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_df *ContentCreator) Add_Tw(wordSpace float64) *ContentCreator {
	_fdga := ContentStreamOperation{}
	_fdga.Operand = "\u0054\u0077"
	_fdga.Params = _gbfad([]float64{wordSpace})
	_df._acd = append(_df._acd, &_fdga)
	return _df
}

// Add_Q adds 'Q' operand to the content stream: Pops the most recently stored state from the stack.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_ea *ContentCreator) Add_Q() *ContentCreator {
	_ged := ContentStreamOperation{}
	_ged.Operand = "\u0051"
	_ea._acd = append(_ea._acd, &_ged)
	return _ea
}
func (_bbgb *ContentStreamParser) parseNull() (_d.PdfObjectNull, error) {
	_, _gad := _bbgb._cecb.Discard(4)
	return _d.PdfObjectNull{}, _gad
}

// Add_g appends 'g' operand to the content stream:
// Same as G but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_afa *ContentCreator) Add_g(gray float64) *ContentCreator {
	_ceac := ContentStreamOperation{}
	_ceac.Operand = "\u0067"
	_ceac.Params = _gbfad([]float64{gray})
	_afa._acd = append(_afa._acd, &_ceac)
	return _afa
}

// Add_Tj appends 'Tj' operand to the content stream:
// Show a text string.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_egf *ContentCreator) Add_Tj(textstr _d.PdfObjectString) *ContentCreator {
	_cgfe := ContentStreamOperation{}
	_cgfe.Operand = "\u0054\u006a"
	_cgfe.Params = _fgag([]_d.PdfObjectString{textstr})
	_egf._acd = append(_egf._acd, &_cgfe)
	return _egf
}
func (_ace *ContentStreamProcessor) handleCommand_rg(_cdfa *ContentStreamOperation, _addf *_cg.PdfPageResources) error {
	_bgf := _cg.NewPdfColorspaceDeviceRGB()
	if len(_cdfa.Params) != _bgf.GetNumComponents() {
		_eb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_eb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_cdfa.Params), _bgf)
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_gcdd, _efac := _bgf.ColorFromPdfObjects(_cdfa.Params)
	if _efac != nil {
		return _efac
	}
	_ace._cff.ColorspaceNonStroking = _bgf
	_ace._cff.ColorNonStroking = _gcdd
	return nil
}

// HandlerFunc is the function syntax that the ContentStreamProcessor handler must implement.
type HandlerFunc func(_eeba *ContentStreamOperation, _efa GraphicsState, _gbd *_cg.PdfPageResources) error

// Add_rg appends 'rg' operand to the content stream:
// Same as RG but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_gb *ContentCreator) Add_rg(r, g, b float64) *ContentCreator {
	_gac := ContentStreamOperation{}
	_gac.Operand = "\u0072\u0067"
	_gac.Params = _gbfad([]float64{r, g, b})
	_gb._acd = append(_gb._acd, &_gac)
	return _gb
}

// Add_h appends 'h' operand to the content stream:
// Close the current subpath by adding a line between the current position and the starting position.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_dce *ContentCreator) Add_h() *ContentCreator {
	_fef := ContentStreamOperation{}
	_fef.Operand = "\u0068"
	_dce._acd = append(_dce._acd, &_fef)
	return _dce
}

// ContentStreamProcessor defines a data structure and methods for processing a content stream, keeping track of the
// current graphics state, and allowing external handlers to define their own functions as a part of the processing,
// for example rendering or extracting certain information.
type ContentStreamProcessor struct {
	_afb  GraphicStateStack
	_aaba []*ContentStreamOperation
	_cff  GraphicsState
	_gfce []handlerEntry
	_beb  int
}

// Add_y appends 'y' operand to the content stream: Append a Bezier curve to the current path from the
// current point to (x3,y3) with (x1, y1) and (x3,y3) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_ced *ContentCreator) Add_y(x1, y1, x3, y3 float64) *ContentCreator {
	_fgc := ContentStreamOperation{}
	_fgc.Operand = "\u0079"
	_fgc.Params = _gbfad([]float64{x1, y1, x3, y3})
	_ced._acd = append(_ced._acd, &_fgc)
	return _ced
}

// Bytes converts the content stream operations to a content stream byte presentation, i.e. the kind that can be
// stored as a PDF stream or string format.
func (_bea *ContentCreator) Bytes() []byte { return _bea._acd.Bytes() }

// Pop pops and returns the topmost GraphicsState off the `gsStack`.
func (_cfae *GraphicStateStack) Pop() GraphicsState {
	_cde := (*_cfae)[len(*_cfae)-1]
	*_cfae = (*_cfae)[:len(*_cfae)-1]
	return _cde
}

// Add_w adds 'w' operand to the content stream, which sets the line width.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_cea *ContentCreator) Add_w(lineWidth float64) *ContentCreator {
	_eaf := ContentStreamOperation{}
	_eaf.Operand = "\u0077"
	_eaf.Params = _gbfad([]float64{lineWidth})
	_cea._acd = append(_cea._acd, &_eaf)
	return _cea
}
func (_fede *ContentStreamProcessor) handleCommand_K(_gcb *ContentStreamOperation, _gfd *_cg.PdfPageResources) error {
	_dgbg := _cg.NewPdfColorspaceDeviceCMYK()
	if len(_gcb.Params) != _dgbg.GetNumComponents() {
		_eb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_eb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_gcb.Params), _dgbg)
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_gfda, _dgbb := _dgbg.ColorFromPdfObjects(_gcb.Params)
	if _dgbb != nil {
		return _dgbb
	}
	_fede._cff.ColorspaceStroking = _dgbg
	_fede._cff.ColorStroking = _gfda
	return nil
}

// Add_Tz appends 'Tz' operand to the content stream:
// Set horizontal scaling.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_ffa *ContentCreator) Add_Tz(scale float64) *ContentCreator {
	_gba := ContentStreamOperation{}
	_gba.Operand = "\u0054\u007a"
	_gba.Params = _gbfad([]float64{scale})
	_ffa._acd = append(_ffa._acd, &_gba)
	return _ffa
}

// Add_Td appends 'Td' operand to the content stream:
// Move to start of next line with offset (`tx`, `ty`).
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_ead *ContentCreator) Add_Td(tx, ty float64) *ContentCreator {
	_eg := ContentStreamOperation{}
	_eg.Operand = "\u0054\u0064"
	_eg.Params = _gbfad([]float64{tx, ty})
	_ead._acd = append(_ead._acd, &_eg)
	return _ead
}
func _agee(_aggb _d.PdfObject) (_cg.PdfColorspace, error) {
	_dbfg, _efe := _aggb.(*_d.PdfObjectArray)
	if !_efe {
		_eb.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0049\u006ev\u0061\u006c\u0069d\u0020\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020cs\u0020\u006e\u006ft\u0020\u0069n\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025#\u0076\u0029", _aggb)
		return nil, _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _dbfg.Len() != 4 {
		_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061r\u0072\u0061\u0079\u002c\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0034\u0020\u0028\u0025\u0064\u0029", _dbfg.Len())
		return nil, _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_dabg, _efe := _dbfg.Get(0).(*_d.PdfObjectName)
	if !_efe {
		_eb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0066\u0069\u0072s\u0074 \u0065\u006c\u0065\u006de\u006e\u0074 \u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0023\u0076\u0029", *_dbfg)
		return nil, _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_dabg != "\u0049" && *_dabg != "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		_eb.Log.Debug("\u0045\u0072r\u006f\u0072\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064 \u0063\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0021\u003d\u0020\u0049\u0020\u0028\u0067\u006f\u0074\u003a\u0020\u0025\u0076\u0029", *_dabg)
		return nil, _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_dabg, _efe = _dbfg.Get(1).(*_d.PdfObjectName)
	if !_efe {
		_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072r\u0061\u0079\u0020\u0032\u006e\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0061\u0072\u0072a\u0079\u003a\u0020\u0025\u0023v\u0029", *_dbfg)
		return nil, _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_dabg != "\u0047" && *_dabg != "\u0052\u0047\u0042" && *_dabg != "\u0043\u004d\u0059\u004b" && *_dabg != "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079" && *_dabg != "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B" && *_dabg != "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b" {
		_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072r\u0061\u0079\u0020\u0032\u006e\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0021\u003d\u0020\u0047\u002f\u0052\u0047\u0042\u002f\u0043\u004d\u0059\u004b\u0020\u0028g\u006f\u0074\u003a\u0020\u0025v\u0029", *_dabg)
		return nil, _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_fbcc := ""
	switch *_dabg {
	case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
		_fbcc = "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079"
	case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
		_fbcc = "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B"
	case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
		_fbcc = "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
	}
	_gefa := _d.MakeArray(_d.MakeName("\u0049n\u0064\u0065\u0078\u0065\u0064"), _d.MakeName(_fbcc), _dbfg.Get(2), _dbfg.Get(3))
	return _cg.NewPdfColorspaceFromPdfObject(_gefa)
}

// Wrap ensures that the contentstream is wrapped within a balanced q ... Q expression.
func (_bba *ContentCreator) Wrap() { _bba._acd.WrapIfNeeded() }
func _fgag(_gedf []_d.PdfObjectString) []_d.PdfObject {
	var _cge []_d.PdfObject
	for _, _gfga := range _gedf {
		_cge = append(_cge, _d.MakeString(_gfga.Str()))
	}
	return _cge
}
func (_efcdc *ContentStreamProcessor) handleCommand_g(_cab *ContentStreamOperation, _ecc *_cg.PdfPageResources) error {
	_ded := _cg.NewPdfColorspaceDeviceGray()
	if len(_cab.Params) != _ded.GetNumComponents() {
		_eb.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020p\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0067")
		_eb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_cab.Params), _ded)
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_bcgg, _eaaa := _ded.ColorFromPdfObjects(_cab.Params)
	if _eaaa != nil {
		_eb.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0068\u0061\u006e\u0064\u006c\u0065\u0043o\u006d\u006d\u0061\u006e\u0064\u005f\u0067\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061r\u0061\u006d\u0073\u002e\u0020c\u0073\u003d\u0025\u0054\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ded, _cab, _eaaa)
		return _eaaa
	}
	_efcdc._cff.ColorspaceNonStroking = _ded
	_efcdc._cff.ColorNonStroking = _bcgg
	return nil
}

// WriteString outputs the object as it is to be written to file.
func (_edc *ContentStreamInlineImage) WriteString() string {
	var _ega _gc.Buffer
	_gea := ""
	if _edc.BitsPerComponent != nil {
		_gea += "\u002f\u0042\u0050C\u0020" + _edc.BitsPerComponent.WriteString() + "\u000a"
	}
	if _edc.ColorSpace != nil {
		_gea += "\u002f\u0043\u0053\u0020" + _edc.ColorSpace.WriteString() + "\u000a"
	}
	if _edc.Decode != nil {
		_gea += "\u002f\u0044\u0020" + _edc.Decode.WriteString() + "\u000a"
	}
	if _edc.DecodeParms != nil {
		_gea += "\u002f\u0044\u0050\u0020" + _edc.DecodeParms.WriteString() + "\u000a"
	}
	if _edc.Filter != nil {
		_gea += "\u002f\u0046\u0020" + _edc.Filter.WriteString() + "\u000a"
	}
	if _edc.Height != nil {
		_gea += "\u002f\u0048\u0020" + _edc.Height.WriteString() + "\u000a"
	}
	if _edc.ImageMask != nil {
		_gea += "\u002f\u0049\u004d\u0020" + _edc.ImageMask.WriteString() + "\u000a"
	}
	if _edc.Intent != nil {
		_gea += "\u002f\u0049\u006e\u0074\u0065\u006e\u0074\u0020" + _edc.Intent.WriteString() + "\u000a"
	}
	if _edc.Interpolate != nil {
		_gea += "\u002f\u0049\u0020" + _edc.Interpolate.WriteString() + "\u000a"
	}
	if _edc.Width != nil {
		_gea += "\u002f\u0057\u0020" + _edc.Width.WriteString() + "\u000a"
	}
	_ega.WriteString(_gea)
	_ega.WriteString("\u0049\u0044\u0020")
	_ega.Write(_edc._fece)
	_ega.WriteString("\u000a\u0045\u0049\u000a")
	return _ega.String()
}

// Add_EMC appends 'EMC' operand to the content stream:
// Ends a marked-content sequence.
//
// See section 14.6 "Marked Content" and Table 320 (p. 561 PDF32000_2008).
func (_bbc *ContentCreator) Add_EMC() *ContentCreator {
	_dfg := ContentStreamOperation{}
	_dfg.Operand = "\u0045\u004d\u0043"
	_bbc._acd = append(_bbc._acd, &_dfg)
	return _bbc
}

// HandlerConditionEnum represents the type of operand content stream processor (handler).
// The handler may process a single specific named operand or all operands.
type HandlerConditionEnum int

// GraphicsState is a basic graphics state implementation for PDF processing.
// Initially only implementing and tracking a portion of the information specified. Easy to add more.
type GraphicsState struct {
	ColorspaceStroking    _cg.PdfColorspace
	ColorspaceNonStroking _cg.PdfColorspace
	ColorStroking         _cg.PdfColor
	ColorNonStroking      _cg.PdfColor
	CTM                   _cf.Matrix
}

// Add_l adds 'l' operand to the content stream:
// Append a straight line segment from the current point to (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_ccce *ContentCreator) Add_l(x, y float64) *ContentCreator {
	_dab := ContentStreamOperation{}
	_dab.Operand = "\u006c"
	_dab.Params = _gbfad([]float64{x, y})
	_ccce._acd = append(_ccce._acd, &_dab)
	return _ccce
}

// Add_K appends 'K' operand to the content stream:
// Set the stroking colorspace to DeviceCMYK and sets the c,m,y,k color (0-1 each component).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_fde *ContentCreator) Add_K(c, m, y, k float64) *ContentCreator {
	_dga := ContentStreamOperation{}
	_dga.Operand = "\u004b"
	_dga.Params = _gbfad([]float64{c, m, y, k})
	_fde._acd = append(_fde._acd, &_dga)
	return _fde
}

const (
	HandlerConditionEnumOperand HandlerConditionEnum = iota
	HandlerConditionEnumAllOperands
)

func (_addd *ContentStreamProcessor) handleCommand_RG(_bacda *ContentStreamOperation, _ece *_cg.PdfPageResources) error {
	_faae := _cg.NewPdfColorspaceDeviceRGB()
	if len(_bacda.Params) != _faae.GetNumComponents() {
		_eb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020R\u0047")
		_eb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_bacda.Params), _faae)
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_dcfb, _bec := _faae.ColorFromPdfObjects(_bacda.Params)
	if _bec != nil {
		return _bec
	}
	_addd._cff.ColorspaceStroking = _faae
	_addd._cff.ColorStroking = _dcfb
	return nil
}

// All returns true if `hce` is equivalent to HandlerConditionEnumAllOperands.
func (_feb HandlerConditionEnum) All() bool { return _feb == HandlerConditionEnumAllOperands }

// Add_J adds 'J' operand to the content stream: Set the line cap style (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_bda *ContentCreator) Add_J(lineCapStyle string) *ContentCreator {
	_gg := ContentStreamOperation{}
	_gg.Operand = "\u004a"
	_gg.Params = _befb([]_d.PdfObjectName{_d.PdfObjectName(lineCapStyle)})
	_bda._acd = append(_bda._acd, &_gg)
	return _bda
}
func (_ecdf *ContentStreamParser) parseOperand() (*_d.PdfObjectString, error) {
	var _fcgd []byte
	for {
		_egff, _dae := _ecdf._cecb.Peek(1)
		if _dae != nil {
			return _d.MakeString(string(_fcgd)), _dae
		}
		if _d.IsDelimiter(_egff[0]) {
			break
		}
		if _d.IsWhiteSpace(_egff[0]) {
			break
		}
		_bdfc, _ := _ecdf._cecb.ReadByte()
		_fcgd = append(_fcgd, _bdfc)
	}
	return _d.MakeString(string(_fcgd)), nil
}

// ContentStreamParser represents a content stream parser for parsing content streams in PDFs.
type ContentStreamParser struct{ _cecb *_cc.Reader }

// Process processes the entire list of operations. Maintains the graphics state that is passed to any
// handlers that are triggered during processing (either on specific operators or all).
func (_ffec *ContentStreamProcessor) Process(resources *_cg.PdfPageResources) error {
	_ffec._cff.ColorspaceStroking = _cg.NewPdfColorspaceDeviceGray()
	_ffec._cff.ColorspaceNonStroking = _cg.NewPdfColorspaceDeviceGray()
	_ffec._cff.ColorStroking = _cg.NewPdfColorDeviceGray(0)
	_ffec._cff.ColorNonStroking = _cg.NewPdfColorDeviceGray(0)
	_ffec._cff.CTM = _cf.IdentityMatrix()
	for _, _aag := range _ffec._aaba {
		var _fag error
		switch _aag.Operand {
		case "\u0071":
			_ffec._afb.Push(_ffec._cff)
		case "\u0051":
			if len(_ffec._afb) == 0 {
				_eb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0060\u0051\u0060\u0020\u006f\u0070e\u0072\u0061\u0074\u006f\u0072\u002e\u0020\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0073\u0020\u0073\u0074\u0061\u0074\u0065 \u0073\u0074\u0061\u0063\u006b\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079.\u0020\u0053\u006bi\u0070\u0070\u0069\u006e\u0067\u002e")
				continue
			}
			_ffec._cff = _ffec._afb.Pop()
		case "\u0043\u0053":
			_fag = _ffec.handleCommand_CS(_aag, resources)
		case "\u0063\u0073":
			_fag = _ffec.handleCommand_cs(_aag, resources)
		case "\u0053\u0043":
			_fag = _ffec.handleCommand_SC(_aag, resources)
		case "\u0053\u0043\u004e":
			_fag = _ffec.handleCommand_SCN(_aag, resources)
		case "\u0073\u0063":
			_fag = _ffec.handleCommand_sc(_aag, resources)
		case "\u0073\u0063\u006e":
			_fag = _ffec.handleCommand_scn(_aag, resources)
		case "\u0047":
			_fag = _ffec.handleCommand_G(_aag, resources)
		case "\u0067":
			_fag = _ffec.handleCommand_g(_aag, resources)
		case "\u0052\u0047":
			_fag = _ffec.handleCommand_RG(_aag, resources)
		case "\u0072\u0067":
			_fag = _ffec.handleCommand_rg(_aag, resources)
		case "\u004b":
			_fag = _ffec.handleCommand_K(_aag, resources)
		case "\u006b":
			_fag = _ffec.handleCommand_k(_aag, resources)
		case "\u0063\u006d":
			_fag = _ffec.handleCommand_cm(_aag, resources)
		}
		if _fag != nil {
			_eb.Log.Debug("\u0050\u0072\u006f\u0063\u0065\u0073s\u006f\u0072\u0020\u0068\u0061\u006e\u0064\u006c\u0069\u006e\u0067\u0020\u0065r\u0072\u006f\u0072\u0020\u0028\u0025\u0073)\u003a\u0020\u0025\u0076", _aag.Operand, _fag)
			_eb.Log.Debug("\u004f\u0070\u0065r\u0061\u006e\u0064\u003a\u0020\u0025\u0023\u0076", _aag.Operand)
			return _fag
		}
		for _, _ecaa := range _ffec._gfce {
			var _fgee error
			if _ecaa.Condition.All() {
				_fgee = _ecaa.Handler(_aag, _ffec._cff, resources)
			} else if _ecaa.Condition.Operand() && _aag.Operand == _ecaa.Operand {
				_fgee = _ecaa.Handler(_aag, _ffec._cff, resources)
			}
			if _fgee != nil {
				_eb.Log.Debug("P\u0072\u006f\u0063\u0065\u0073\u0073o\u0072\u0020\u0068\u0061\u006e\u0064\u006c\u0065\u0072 \u0065\u0072\u0072o\u0072:\u0020\u0025\u0076", _fgee)
				return _fgee
			}
		}
	}
	return nil
}

// Add_v appends 'v' operand to the content stream: Append a Bezier curve to the current path from the
// current point to (x3,y3) with the current point and (x2,y2) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_ggg *ContentCreator) Add_v(x2, y2, x3, y3 float64) *ContentCreator {
	_gab := ContentStreamOperation{}
	_gab.Operand = "\u0076"
	_gab.Params = _gbfad([]float64{x2, y2, x3, y3})
	_ggg._acd = append(_ggg._acd, &_gab)
	return _ggg
}

// Add_Tr appends 'Tr' operand to the content stream:
// Set text rendering mode.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_gge *ContentCreator) Add_Tr(render int64) *ContentCreator {
	_bga := ContentStreamOperation{}
	_bga.Operand = "\u0054\u0072"
	_bga.Params = _ccdbd([]int64{render})
	_gge._acd = append(_gge._acd, &_bga)
	return _gge
}
func (_ddda *ContentStreamParser) parseDict() (*_d.PdfObjectDictionary, error) {
	_eb.Log.Trace("\u0052\u0065\u0061\u0064i\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074 \u0073t\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0021")
	_cfb := _d.MakeDict()
	_dfda, _ := _ddda._cecb.ReadByte()
	if _dfda != '<' {
		return nil, _a.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	_dfda, _ = _ddda._cecb.ReadByte()
	if _dfda != '<' {
		return nil, _a.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	for {
		_ddda.skipSpaces()
		_agbd, _afd := _ddda._cecb.Peek(2)
		if _afd != nil {
			return nil, _afd
		}
		_eb.Log.Trace("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_agbd), string(_agbd))
		if (_agbd[0] == '>') && (_agbd[1] == '>') {
			_eb.Log.Trace("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
			_ddda._cecb.ReadByte()
			_ddda._cecb.ReadByte()
			break
		}
		_eb.Log.Trace("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021")
		_fgcd, _afd := _ddda.parseName()
		_eb.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _fgcd)
		if _afd != nil {
			_eb.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073", _afd)
			return nil, _afd
		}
		if len(_fgcd) > 4 && _fgcd[len(_fgcd)-4:] == "\u006e\u0075\u006c\u006c" {
			_bfb := _fgcd[0 : len(_fgcd)-4]
			_eb.Log.Trace("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029", _fgcd)
			_eb.Log.Trace("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c", _bfb)
			_ddda.skipSpaces()
			_ebgd, _ := _ddda._cecb.Peek(1)
			if _ebgd[0] == '/' {
				_cfb.Set(_bfb, _d.MakeNull())
				continue
			}
		}
		_ddda.skipSpaces()
		_afac, _, _afd := _ddda.parseObject()
		if _afd != nil {
			return nil, _afd
		}
		_cfb.Set(_fgcd, _afac)
		_eb.Log.Trace("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073", _fgcd, _afac.String())
	}
	return _cfb, nil
}
func (_ffaf *ContentStreamProcessor) handleCommand_CS(_bfgb *ContentStreamOperation, _gadb *_cg.PdfPageResources) error {
	if len(_bfgb.Params) < 1 {
		_eb.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0063\u0073\u0020\u0063\u006f\u006d\u006d\u0061\u006e\u0064\u002c\u0020s\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u006f\u0076\u0065\u0072")
		return _a.New("\u0074o\u006f \u0066\u0065\u0077\u0020\u0070a\u0072\u0061m\u0065\u0074\u0065\u0072\u0073")
	}
	if len(_bfgb.Params) > 1 {
		_eb.Log.Debug("\u0063\u0073\u0020\u0063\u006f\u006d\u006d\u0061n\u0064\u0020\u0077it\u0068\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020-\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075i\u006e\u0067")
		return _a.New("\u0074\u006f\u006f\u0020ma\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	}
	_ggbb, _ebe := _bfgb.Params[0].(*_d.PdfObjectName)
	if !_ebe {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020c\u0073\u0020\u0063o\u006d\u006d\u0061n\u0064\u0020w\u0069\u0074\u0068\u0020\u0069\u006ev\u0061li\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u002c\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072")
		return _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_gbfd, _cbfg := _ffaf.getColorspace(string(*_ggbb), _gadb)
	if _cbfg != nil {
		return _cbfg
	}
	_ffaf._cff.ColorspaceStroking = _gbfd
	_fba, _cbfg := _ffaf.getInitialColor(_gbfd)
	if _cbfg != nil {
		return _cbfg
	}
	_ffaf._cff.ColorStroking = _fba
	return nil
}

// Add_n appends 'n' operand to the content stream:
// End the path without filling or stroking.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_cdg *ContentCreator) Add_n() *ContentCreator {
	_bed := ContentStreamOperation{}
	_bed.Operand = "\u006e"
	_cdg._acd = append(_cdg._acd, &_bed)
	return _cdg
}

// Translate applies a simple x-y translation to the transformation matrix.
func (_bbb *ContentCreator) Translate(tx, ty float64) *ContentCreator {
	return _bbb.Add_cm(1, 0, 0, 1, tx, ty)
}

// ContentStreamOperation represents an operation in PDF contentstream which consists of
// an operand and parameters.
type ContentStreamOperation struct {
	Params  []_d.PdfObject
	Operand string
}

// Add_cs appends 'cs' operand to the content stream:
// Same as CS but for non-stroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_gfb *ContentCreator) Add_cs(name _d.PdfObjectName) *ContentCreator {
	_dca := ContentStreamOperation{}
	_dca.Operand = "\u0063\u0073"
	_dca.Params = _befb([]_d.PdfObjectName{name})
	_gfb._acd = append(_gfb._acd, &_dca)
	return _gfb
}

// Add_s appends 's' operand to the content stream: Close and stroke the path.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_daa *ContentCreator) Add_s() *ContentCreator {
	_dge := ContentStreamOperation{}
	_dge.Operand = "\u0073"
	_daa._acd = append(_daa._acd, &_dge)
	return _daa
}

// Add_sh appends 'sh' operand to the content stream:
// Paints the shape and colour shading described by a shading dictionary specified by `name`,
// subject to the current clipping path
//
// See section 8.7.4 "Shading Patterns" and Table 77 (p. 190 PDF32000_2008).
func (_eab *ContentCreator) Add_sh(name _d.PdfObjectName) *ContentCreator {
	_gbc := ContentStreamOperation{}
	_gbc.Operand = "\u0073\u0068"
	_gbc.Params = _befb([]_d.PdfObjectName{name})
	_eab._acd = append(_eab._acd, &_gbc)
	return _eab
}

// Add_TL appends 'TL' operand to the content stream:
// Set leading.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_ffea *ContentCreator) Add_TL(leading float64) *ContentCreator {
	_ca := ContentStreamOperation{}
	_ca.Operand = "\u0054\u004c"
	_ca.Params = _gbfad([]float64{leading})
	_ffea._acd = append(_ffea._acd, &_ca)
	return _ffea
}

// NewContentCreator returns a new initialized ContentCreator.
func NewContentCreator() *ContentCreator {
	_ffc := &ContentCreator{}
	_ffc._acd = ContentStreamOperations{}
	return _ffc
}

// SetStrokingColor sets the stroking `color` where color can be one of
// PdfColorDeviceGray, PdfColorDeviceRGB, or PdfColorDeviceCMYK.
func (_dgg *ContentCreator) SetStrokingColor(color _cg.PdfColor) *ContentCreator {
	switch _gfc := color.(type) {
	case *_cg.PdfColorDeviceGray:
		_dgg.Add_G(_gfc.Val())
	case *_cg.PdfColorDeviceRGB:
		_dgg.Add_RG(_gfc.R(), _gfc.G(), _gfc.B())
	case *_cg.PdfColorDeviceCMYK:
		_dgg.Add_K(_gfc.C(), _gfc.M(), _gfc.Y(), _gfc.K())
	default:
		_eb.Log.Debug("\u0053\u0065\u0074\u0053\u0074\u0072\u006f\u006b\u0069\u006e\u0067\u0043\u006fl\u006f\u0072\u003a\u0020\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006fr\u003a\u0020\u0025\u0054", _gfc)
	}
	return _dgg
}

// ContentStreamOperations is a slice of ContentStreamOperations.
type ContentStreamOperations []*ContentStreamOperation

func (_eac *ContentStreamParser) parseNumber() (_d.PdfObject, error) {
	return _d.ParseNumber(_eac._cecb)
}
func (_ebea *ContentStreamProcessor) handleCommand_cm(_gcc *ContentStreamOperation, _cgac *_cg.PdfPageResources) error {
	if len(_gcc.Params) != 6 {
		_eb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0063\u006d\u003a\u0020\u0025\u0064", len(_gcc.Params))
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_cebc, _eagb := _d.GetNumbersAsFloat(_gcc.Params)
	if _eagb != nil {
		return _eagb
	}
	_ccea := _cf.NewMatrix(_cebc[0], _cebc[1], _cebc[2], _cebc[3], _cebc[4], _cebc[5])
	_ebea._cff.CTM.Concat(_ccea)
	return nil
}
func _fa(_bdg *ContentStreamInlineImage) (*_d.MultiEncoder, error) {
	_aegg := _d.NewMultiEncoder()
	var _ebf *_d.PdfObjectDictionary
	var _fcc []_d.PdfObject
	if _bcg := _bdg.DecodeParms; _bcg != nil {
		_dccc, _adc := _bcg.(*_d.PdfObjectDictionary)
		if _adc {
			_ebf = _dccc
		}
		_bcee, _fefa := _bcg.(*_d.PdfObjectArray)
		if _fefa {
			for _, _egb := range _bcee.Elements() {
				if _cbee, _aaf := _egb.(*_d.PdfObjectDictionary); _aaf {
					_fcc = append(_fcc, _cbee)
				} else {
					_fcc = append(_fcc, nil)
				}
			}
		}
	}
	_bbbg := _bdg.Filter
	if _bbbg == nil {
		return nil, _fc.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	_gdd, _daaf := _bbbg.(*_d.PdfObjectArray)
	if !_daaf {
		return nil, _fc.Errorf("m\u0075\u006c\u0074\u0069\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0063\u0061\u006e\u0020\u006f\u006el\u0079\u0020\u0062\u0065\u0020\u006d\u0061\u0064\u0065\u0020fr\u006f\u006d\u0020a\u0072r\u0061\u0079")
	}
	for _fec, _cbed := range _gdd.Elements() {
		_gafc, _bdaa := _cbed.(*_d.PdfObjectName)
		if !_bdaa {
			return nil, _fc.Errorf("\u006d\u0075l\u0074\u0069\u0020\u0066i\u006c\u0074e\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061 \u006e\u0061\u006d\u0065")
		}
		var _cdf _d.PdfObject
		if _ebf != nil {
			_cdf = _ebf
		} else {
			if len(_fcc) > 0 {
				if _fec >= len(_fcc) {
					return nil, _fc.Errorf("\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0073\u0020\u0069\u006e\u0020d\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006d\u0073\u0020a\u0072\u0072\u0061\u0079")
				}
				_cdf = _fcc[_fec]
			}
		}
		var _gbb *_d.PdfObjectDictionary
		if _fcfa, _efg := _cdf.(*_d.PdfObjectDictionary); _efg {
			_gbb = _fcfa
		}
		if *_gafc == _d.StreamEncodingFilterNameFlate || *_gafc == "\u0046\u006c" {
			_eef, _fdac := _cec(_bdg, _gbb)
			if _fdac != nil {
				return nil, _fdac
			}
			_aegg.AddEncoder(_eef)
		} else if *_gafc == _d.StreamEncodingFilterNameLZW {
			_edea, _ffcd := _ceea(_bdg, _gbb)
			if _ffcd != nil {
				return nil, _ffcd
			}
			_aegg.AddEncoder(_edea)
		} else if *_gafc == _d.StreamEncodingFilterNameASCIIHex {
			_bbf := _d.NewASCIIHexEncoder()
			_aegg.AddEncoder(_bbf)
		} else if *_gafc == _d.StreamEncodingFilterNameASCII85 || *_gafc == "\u0041\u0038\u0035" {
			_cgfc := _d.NewASCII85Encoder()
			_aegg.AddEncoder(_cgfc)
		} else {
			_eb.Log.Error("U\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0066\u0069l\u0074\u0065\u0072\u0020\u0025\u0073", *_gafc)
			return nil, _fc.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u0066\u0069\u006c\u0074er \u0069n \u006d\u0075\u006c\u0074\u0069\u0020\u0066il\u0074\u0065\u0072\u0020\u0061\u0072\u0072a\u0079")
		}
	}
	return _aegg, nil
}
func _gcf(_dfb *ContentStreamInlineImage) (_d.StreamEncoder, error) {
	if _dfb.Filter == nil {
		return _d.NewRawEncoder(), nil
	}
	_cgb, _adga := _dfb.Filter.(*_d.PdfObjectName)
	if !_adga {
		_bbd, _bfe := _dfb.Filter.(*_d.PdfObjectArray)
		if !_bfe {
			return nil, _fc.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0072 \u0041\u0072\u0072\u0061\u0079\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
		if _bbd.Len() == 0 {
			return _d.NewRawEncoder(), nil
		}
		if _bbd.Len() != 1 {
			_ceg, _fge := _fa(_dfb)
			if _fge != nil {
				_eb.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0063\u0072\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u006d\u0075\u006c\u0074i\u0020\u0065\u006e\u0063\u006f\u0064\u0065r\u003a\u0020\u0025\u0076", _fge)
				return nil, _fge
			}
			_eb.Log.Trace("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063:\u0020\u0025\u0073\u000a", _ceg)
			return _ceg, nil
		}
		_cga := _bbd.Get(0)
		_cgb, _bfe = _cga.(*_d.PdfObjectName)
		if !_bfe {
			return nil, _fc.Errorf("\u0066\u0069l\u0074\u0065\u0072\u0020a\u0072\u0072a\u0079\u0020\u006d\u0065\u006d\u0062\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
	}
	switch *_cgb {
	case "\u0041\u0048\u0078", "\u0041\u0053\u0043\u0049\u0049\u0048\u0065\u0078\u0044e\u0063\u006f\u0064\u0065":
		return _d.NewASCIIHexEncoder(), nil
	case "\u0041\u0038\u0035", "\u0041\u0053\u0043\u0049\u0049\u0038\u0035\u0044\u0065\u0063\u006f\u0064\u0065":
		return _d.NewASCII85Encoder(), nil
	case "\u0044\u0043\u0054", "\u0044C\u0054\u0044\u0065\u0063\u006f\u0064e":
		return _bdae(_dfb)
	case "\u0046\u006c", "F\u006c\u0061\u0074\u0065\u0044\u0065\u0063\u006f\u0064\u0065":
		return _cec(_dfb, nil)
	case "\u004c\u005a\u0057", "\u004cZ\u0057\u0044\u0065\u0063\u006f\u0064e":
		return _ceea(_dfb, nil)
	case "\u0043\u0043\u0046", "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065":
		return _d.NewCCITTFaxEncoder(), nil
	case "\u0052\u004c", "\u0052u\u006eL\u0065\u006e\u0067\u0074\u0068\u0044\u0065\u0063\u006f\u0064\u0065":
		return _d.NewRunLengthEncoder(), nil
	default:
		_eb.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065 \u0069\u006d\u0061\u0067\u0065\u0020\u0065n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u003a\u0020\u0025\u0073", *_cgb)
		return nil, _a.New("\u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006el\u0069n\u0065 \u0065n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006f\u0064")
	}
}

// Add_CS appends 'CS' operand to the content stream:
// Set the current colorspace for stroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_edf *ContentCreator) Add_CS(name _d.PdfObjectName) *ContentCreator {
	_abd := ContentStreamOperation{}
	_abd.Operand = "\u0043\u0053"
	_abd.Params = _befb([]_d.PdfObjectName{name})
	_edf._acd = append(_edf._acd, &_abd)
	return _edf
}

// Add_Tm appends 'Tm' operand to the content stream:
// Set the text line matrix.
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_dgb *ContentCreator) Add_Tm(a, b, c, d, e, f float64) *ContentCreator {
	_aeg := ContentStreamOperation{}
	_aeg.Operand = "\u0054\u006d"
	_aeg.Params = _gbfad([]float64{a, b, c, d, e, f})
	_dgb._acd = append(_dgb._acd, &_aeg)
	return _dgb
}

// Add_Ts appends 'Ts' operand to the content stream:
// Set text rise.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_abe *ContentCreator) Add_Ts(rise float64) *ContentCreator {
	_cdc := ContentStreamOperation{}
	_cdc.Operand = "\u0054\u0073"
	_cdc.Params = _gbfad([]float64{rise})
	_abe._acd = append(_abe._acd, &_cdc)
	return _abe
}
func (_dgff *ContentStreamParser) parseArray() (*_d.PdfObjectArray, error) {
	_gbce := _d.MakeArray()
	_dgff._cecb.ReadByte()
	for {
		_dgff.skipSpaces()
		_gbcb, _fbg := _dgff._cecb.Peek(1)
		if _fbg != nil {
			return _gbce, _fbg
		}
		if _gbcb[0] == ']' {
			_dgff._cecb.ReadByte()
			break
		}
		_bbgd, _, _fbg := _dgff.parseObject()
		if _fbg != nil {
			return _gbce, _fbg
		}
		_gbce.Append(_bbgd)
	}
	return _gbce, nil
}
func (_fcfb *ContentStreamInlineImage) String() string {
	_edg := _fc.Sprintf("I\u006el\u0069\u006e\u0065\u0049\u006d\u0061\u0067\u0065(\u006c\u0065\u006e\u003d%d\u0029\u000a", len(_fcfb._fece))
	if _fcfb.BitsPerComponent != nil {
		_edg += "\u002d\u0020\u0042\u0050\u0043\u0020" + _fcfb.BitsPerComponent.WriteString() + "\u000a"
	}
	if _fcfb.ColorSpace != nil {
		_edg += "\u002d\u0020\u0043S\u0020" + _fcfb.ColorSpace.WriteString() + "\u000a"
	}
	if _fcfb.Decode != nil {
		_edg += "\u002d\u0020\u0044\u0020" + _fcfb.Decode.WriteString() + "\u000a"
	}
	if _fcfb.DecodeParms != nil {
		_edg += "\u002d\u0020\u0044P\u0020" + _fcfb.DecodeParms.WriteString() + "\u000a"
	}
	if _fcfb.Filter != nil {
		_edg += "\u002d\u0020\u0046\u0020" + _fcfb.Filter.WriteString() + "\u000a"
	}
	if _fcfb.Height != nil {
		_edg += "\u002d\u0020\u0048\u0020" + _fcfb.Height.WriteString() + "\u000a"
	}
	if _fcfb.ImageMask != nil {
		_edg += "\u002d\u0020\u0049M\u0020" + _fcfb.ImageMask.WriteString() + "\u000a"
	}
	if _fcfb.Intent != nil {
		_edg += "\u002d \u0049\u006e\u0074\u0065\u006e\u0074 " + _fcfb.Intent.WriteString() + "\u000a"
	}
	if _fcfb.Interpolate != nil {
		_edg += "\u002d\u0020\u0049\u0020" + _fcfb.Interpolate.WriteString() + "\u000a"
	}
	if _fcfb.Width != nil {
		_edg += "\u002d\u0020\u0057\u0020" + _fcfb.Width.WriteString() + "\u000a"
	}
	return _edg
}

// NewInlineImageFromImage makes a new content stream inline image object from an image.
func NewInlineImageFromImage(img _cg.Image, encoder _d.StreamEncoder) (*ContentStreamInlineImage, error) {
	if encoder == nil {
		encoder = _d.NewRawEncoder()
	}
	encoder.UpdateParams(img.GetParamsDict())
	_cecc := ContentStreamInlineImage{}
	if img.ColorComponents == 1 {
		_cecc.ColorSpace = _d.MakeName("\u0047")
	} else if img.ColorComponents == 3 {
		_cecc.ColorSpace = _d.MakeName("\u0052\u0047\u0042")
	} else if img.ColorComponents == 4 {
		_cecc.ColorSpace = _d.MakeName("\u0043\u004d\u0059\u004b")
	} else {
		_eb.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006db\u0065\u0072\u0020o\u0066\u0020c\u006f\u006c\u006f\u0072\u0020\u0063o\u006dpo\u006e\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0072\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0064", img.ColorComponents)
		return nil, _a.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020c\u006fl\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073")
	}
	_cecc.BitsPerComponent = _d.MakeInteger(img.BitsPerComponent)
	_cecc.Width = _d.MakeInteger(img.Width)
	_cecc.Height = _d.MakeInteger(img.Height)
	_abb, _aab := encoder.EncodeBytes(img.Data)
	if _aab != nil {
		return nil, _aab
	}
	_cecc._fece = _abb
	_cced := encoder.GetFilterName()
	if _cced != _d.StreamEncodingFilterNameRaw {
		_cecc.Filter = _d.MakeName(_cced)
	}
	return &_cecc, nil
}

// Operand returns true if `hce` is equivalent to HandlerConditionEnumOperand.
func (_fgefe HandlerConditionEnum) Operand() bool { return _fgefe == HandlerConditionEnumOperand }

// ContentStreamInlineImage is a representation of an inline image in a Content stream. Everything between the BI and EI operands.
// ContentStreamInlineImage implements the core.PdfObject interface although strictly it is not a PDF object.
type ContentStreamInlineImage struct {
	BitsPerComponent _d.PdfObject
	ColorSpace       _d.PdfObject
	Decode           _d.PdfObject
	DecodeParms      _d.PdfObject
	Filter           _d.PdfObject
	Height           _d.PdfObject
	ImageMask        _d.PdfObject
	Intent           _d.PdfObject
	Interpolate      _d.PdfObject
	Width            _d.PdfObject
	_fece            []byte
	_cgbc            *_af.ImageBase
}

// ParseInlineImage parses an inline image from a content stream, both reading its properties and binary data.
// When called, "BI" has already been read from the stream.  This function
// finishes reading through "EI" and then returns the ContentStreamInlineImage.
func (_aeea *ContentStreamParser) ParseInlineImage() (*ContentStreamInlineImage, error) {
	_gcfc := ContentStreamInlineImage{}
	for {
		_aeea.skipSpaces()
		_dfde, _cbf, _cbd := _aeea.parseObject()
		if _cbd != nil {
			return nil, _cbd
		}
		if !_cbf {
			_dfeb, _fgbb := _d.GetName(_dfde)
			if !_fgbb {
				_eb.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067e\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0020\u0028\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u0029\u0020\u002d\u0020\u0025T", _dfde)
				return nil, _fc.Errorf("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067e\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0020\u0028\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u0029\u0020\u002d\u0020\u0025T", _dfde)
			}
			_dea, _fgbe, _dced := _aeea.parseObject()
			if _dced != nil {
				return nil, _dced
			}
			if _fgbe {
				return nil, _fc.Errorf("\u006eo\u0074\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067 \u0061\u006e\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
			}
			switch *_dfeb {
			case "\u0042\u0050\u0043", "\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074":
				_gcfc.BitsPerComponent = _dea
			case "\u0043\u0053", "\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065":
				_gcfc.ColorSpace = _dea
			case "\u0044", "\u0044\u0065\u0063\u006f\u0064\u0065":
				_gcfc.Decode = _dea
			case "\u0044\u0050", "D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073":
				_gcfc.DecodeParms = _dea
			case "\u0046", "\u0046\u0069\u006c\u0074\u0065\u0072":
				_gcfc.Filter = _dea
			case "\u0048", "\u0048\u0065\u0069\u0067\u0068\u0074":
				_gcfc.Height = _dea
			case "\u0049\u004d", "\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k":
				_gcfc.ImageMask = _dea
			case "\u0049\u006e\u0074\u0065\u006e\u0074":
				_gcfc.Intent = _dea
			case "\u0049", "I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065":
				_gcfc.Interpolate = _dea
			case "\u0057", "\u0057\u0069\u0064t\u0068":
				_gcfc.Width = _dea
			case "\u004c\u0065\u006e\u0067\u0074\u0068", "\u0053u\u0062\u0074\u0079\u0070\u0065", "\u0054\u0079\u0070\u0065":
				_eb.Log.Debug("\u0049\u0067\u006e\u006fr\u0069\u006e\u0067\u0020\u0069\u006e\u006c\u0069\u006e\u0065 \u0070a\u0072\u0061\u006d\u0065\u0074\u0065\u0072 \u0025\u0073", *_dfeb)
			default:
				return nil, _fc.Errorf("\u0075\u006e\u006b\u006e\u006f\u0077n\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0020\u0025\u0073", *_dfeb)
			}
		}
		if _cbf {
			_bbaf, _gbf := _dfde.(*_d.PdfObjectString)
			if !_gbf {
				return nil, _fc.Errorf("\u0066a\u0069\u006ce\u0064\u0020\u0074o\u0020\u0072\u0065\u0061\u0064\u0020\u0069n\u006c\u0069\u006e\u0065\u0020\u0069m\u0061\u0067\u0065\u0020\u002d\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
			}
			if _bbaf.Str() == "\u0045\u0049" {
				_eb.Log.Trace("\u0049n\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020f\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e\u002e\u002e")
				return &_gcfc, nil
			} else if _bbaf.Str() == "\u0049\u0044" {
				_eb.Log.Trace("\u0049\u0044\u0020\u0073\u0074\u0061\u0072\u0074")
				_ebbg, _gca := _aeea._cecb.Peek(1)
				if _gca != nil {
					return nil, _gca
				}
				if _d.IsWhiteSpace(_ebbg[0]) {
					_aeea._cecb.Discard(1)
				}
				_gcfc._fece = []byte{}
				_dggc := 0
				var _ffg []byte
				for {
					_fce, _ege := _aeea._cecb.ReadByte()
					if _ege != nil {
						_eb.Log.Debug("\u0055\u006e\u0061\u0062\u006ce\u0020\u0074\u006f\u0020\u0066\u0069\u006e\u0064\u0020\u0065\u006e\u0064\u0020o\u0066\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0045\u0049\u0020\u0069\u006e\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u0061\u0074a")
						return nil, _ege
					}
					if _dggc == 0 {
						if _d.IsWhiteSpace(_fce) {
							_ffg = []byte{}
							_ffg = append(_ffg, _fce)
							_dggc = 1
						} else if _fce == 'E' {
							_ffg = append(_ffg, _fce)
							_dggc = 2
						} else {
							_gcfc._fece = append(_gcfc._fece, _fce)
						}
					} else if _dggc == 1 {
						_ffg = append(_ffg, _fce)
						if _fce == 'E' {
							_dggc = 2
						} else {
							_gcfc._fece = append(_gcfc._fece, _ffg...)
							_ffg = []byte{}
							if _d.IsWhiteSpace(_fce) {
								_dggc = 1
							} else {
								_dggc = 0
							}
						}
					} else if _dggc == 2 {
						_ffg = append(_ffg, _fce)
						if _fce == 'I' {
							_dggc = 3
						} else {
							_gcfc._fece = append(_gcfc._fece, _ffg...)
							_ffg = []byte{}
							_dggc = 0
						}
					} else if _dggc == 3 {
						_ffg = append(_ffg, _fce)
						if _d.IsWhiteSpace(_fce) {
							_abgb, _ebd := _aeea._cecb.Peek(20)
							if _ebd != nil && _ebd != _ccg.EOF {
								return nil, _ebd
							}
							_bggd := NewContentStreamParser(string(_abgb))
							_gcaf := true
							for _accc := 0; _accc < 3; _accc++ {
								_dfdc, _bedc, _bdca := _bggd.parseObject()
								if _bdca != nil {
									if _bdca == _ccg.EOF {
										break
									}
									_gcaf = false
									continue
								}
								if _bedc && !_cae(_dfdc.String()) {
									_gcaf = false
									break
								}
							}
							if _gcaf {
								if len(_gcfc._fece) > 100 {
									_eb.Log.Trace("\u0049\u006d\u0061\u0067\u0065\u0020\u0073\u0074\u0072\u0065\u0061m\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078 \u002e\u002e\u002e", len(_gcfc._fece), _gcfc._fece[:100])
								} else {
									_eb.Log.Trace("\u0049\u006d\u0061\u0067e \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025 \u0078", len(_gcfc._fece), _gcfc._fece)
								}
								return &_gcfc, nil
							}
						}
						_gcfc._fece = append(_gcfc._fece, _ffg...)
						_ffg = []byte{}
						_dggc = 0
					}
				}
			}
		}
	}
}

// Add_TJ appends 'TJ' operand to the content stream:
// Show one or more text string. Array of numbers (displacement) and strings.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_bfc *ContentCreator) Add_TJ(vals ..._d.PdfObject) *ContentCreator {
	_cef := ContentStreamOperation{}
	_cef.Operand = "\u0054\u004a"
	_cef.Params = []_d.PdfObject{_d.MakeArray(vals...)}
	_bfc._acd = append(_bfc._acd, &_cef)
	return _bfc
}

// Add_M adds 'M' operand to the content stream: Set the miter limit (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_bg *ContentCreator) Add_M(miterlimit float64) *ContentCreator {
	_bef := ContentStreamOperation{}
	_bef.Operand = "\u004d"
	_bef.Params = _gbfad([]float64{miterlimit})
	_bg._acd = append(_bg._acd, &_bef)
	return _bg
}

// Add_d adds 'd' operand to the content stream: Set the line dash pattern.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_eba *ContentCreator) Add_d(dashArray []int64, dashPhase int64) *ContentCreator {
	_gdc := ContentStreamOperation{}
	_gdc.Operand = "\u0064"
	_gdc.Params = []_d.PdfObject{}
	_gdc.Params = append(_gdc.Params, _d.MakeArrayFromIntegers64(dashArray))
	_gdc.Params = append(_gdc.Params, _d.MakeInteger(dashPhase))
	_eba._acd = append(_eba._acd, &_gdc)
	return _eba
}
func _ceea(_ffb *ContentStreamInlineImage, _acf *_d.PdfObjectDictionary) (*_d.LZWEncoder, error) {
	_gcea := _d.NewLZWEncoder()
	if _acf == nil {
		if _ffb.DecodeParms != nil {
			_acc, _cdcc := _d.GetDict(_ffb.DecodeParms)
			if !_cdcc {
				_eb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079 \u0028%\u0054\u0029", _ffb.DecodeParms)
				return nil, _fc.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
			_acf = _acc
		}
	}
	if _acf == nil {
		return _gcea, nil
	}
	_bbce := _acf.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
	if _bbce != nil {
		_bdc, _ccee := _bbce.(*_d.PdfObjectInteger)
		if !_ccee {
			_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a \u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u006e\u0075\u006d\u0065\u0072i\u0063 \u0028\u0025\u0054\u0029", _bbce)
			return nil, _fc.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
		}
		if *_bdc != 0 && *_bdc != 1 {
			return nil, _fc.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0076\u0061\u006c\u0075e\u0020\u0028\u006e\u006f\u0074 \u0030\u0020o\u0072\u0020\u0031\u0029")
		}
		_gcea.EarlyChange = int(*_bdc)
	} else {
		_gcea.EarlyChange = 1
	}
	_bbce = _acf.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _bbce != nil {
		_gga, _gafe := _bbce.(*_d.PdfObjectInteger)
		if !_gafe {
			_eb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _bbce)
			return nil, _fc.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_gcea.Predictor = int(*_gga)
	}
	_bbce = _acf.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _bbce != nil {
		_ccb, _eaa := _bbce.(*_d.PdfObjectInteger)
		if !_eaa {
			_eb.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _fc.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_gcea.BitsPerComponent = int(*_ccb)
	}
	if _gcea.Predictor > 1 {
		_gcea.Columns = 1
		_bbce = _acf.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _bbce != nil {
			_bdd, _gcg := _bbce.(*_d.PdfObjectInteger)
			if !_gcg {
				return nil, _fc.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_gcea.Columns = int(*_bdd)
		}
		_gcea.Colors = 1
		_bbce = _acf.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _bbce != nil {
			_cgc, _fgab := _bbce.(*_d.PdfObjectInteger)
			if !_fgab {
				return nil, _fc.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_gcea.Colors = int(*_cgc)
		}
	}
	_eb.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _acf.String())
	return _gcea, nil
}
func (_bcgf *ContentStreamParser) parseName() (_d.PdfObjectName, error) {
	_agac := ""
	_bfg := false
	for {
		_cgbcc, _cdd := _bcgf._cecb.Peek(1)
		if _cdd == _ccg.EOF {
			break
		}
		if _cdd != nil {
			return _d.PdfObjectName(_agac), _cdd
		}
		if !_bfg {
			if _cgbcc[0] == '/' {
				_bfg = true
				_bcgf._cecb.ReadByte()
			} else {
				_eb.Log.Error("N\u0061\u006d\u0065\u0020\u0073\u0074a\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069\u0074h\u0020\u0025\u0073 \u0028%\u0020\u0078\u0029", _cgbcc, _cgbcc)
				return _d.PdfObjectName(_agac), _fc.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _cgbcc[0])
			}
		} else {
			if _d.IsWhiteSpace(_cgbcc[0]) {
				break
			} else if (_cgbcc[0] == '/') || (_cgbcc[0] == '[') || (_cgbcc[0] == '(') || (_cgbcc[0] == ']') || (_cgbcc[0] == '<') || (_cgbcc[0] == '>') {
				break
			} else if _cgbcc[0] == '#' {
				_gff, _ccef := _bcgf._cecb.Peek(3)
				if _ccef != nil {
					return _d.PdfObjectName(_agac), _ccef
				}
				_bcgf._cecb.Discard(3)
				_aabf, _ccef := _f.DecodeString(string(_gff[1:3]))
				if _ccef != nil {
					return _d.PdfObjectName(_agac), _ccef
				}
				_agac += string(_aabf)
			} else {
				_dbc, _ := _bcgf._cecb.ReadByte()
				_agac += string(_dbc)
			}
		}
	}
	return _d.PdfObjectName(_agac), nil
}

// Add_gs adds 'gs' operand to the content stream: Set the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_eag *ContentCreator) Add_gs(dictName _d.PdfObjectName) *ContentCreator {
	_dg := ContentStreamOperation{}
	_dg.Operand = "\u0067\u0073"
	_dg.Params = _befb([]_d.PdfObjectName{dictName})
	_eag._acd = append(_eag._acd, &_dg)
	return _eag
}

// Add_B_starred appends 'B*' operand to the content stream:
// Fill and then stroke the path (even-odd rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_fee *ContentCreator) Add_B_starred() *ContentCreator {
	_cfa := ContentStreamOperation{}
	_cfa.Operand = "\u0042\u002a"
	_fee._acd = append(_fee._acd, &_cfa)
	return _fee
}
func _cded(_ccdb _cg.PdfColorspace) bool {
	_, _cegd := _ccdb.(*_cg.PdfColorspaceSpecialPattern)
	return _cegd
}

// Add_k appends 'k' operand to the content stream:
// Same as K but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_cbe *ContentCreator) Add_k(c, m, y, k float64) *ContentCreator {
	_ad := ContentStreamOperation{}
	_ad.Operand = "\u006b"
	_ad.Params = _gbfad([]float64{c, m, y, k})
	_cbe._acd = append(_cbe._acd, &_ad)
	return _cbe
}

// Add_TD appends 'TD' operand to the content stream:
// Move to start of next line with offset (`tx`, `ty`).
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_bce *ContentCreator) Add_TD(tx, ty float64) *ContentCreator {
	_dda := ContentStreamOperation{}
	_dda.Operand = "\u0054\u0044"
	_dda.Params = _gbfad([]float64{tx, ty})
	_bce._acd = append(_bce._acd, &_dda)
	return _bce
}

// Add_SC appends 'SC' operand to the content stream:
// Set color for stroking operations.  Input: c1, ..., cn.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_ccd *ContentCreator) Add_SC(c ...float64) *ContentCreator {
	_geg := ContentStreamOperation{}
	_geg.Operand = "\u0053\u0043"
	_geg.Params = _gbfad(c)
	_ccd._acd = append(_ccd._acd, &_geg)
	return _ccd
}

// Add_BMC appends 'BMC' operand to the content stream:
// Begins a marked-content sequence terminated by a balancing EMC operator.
// `tag` shall be a name object indicating the role or significance of
// the sequence.
//
// See section 14.6 "Marked Content" and Table 320 (p. 561 PDF32000_2008).
func (_ccdc *ContentCreator) Add_BMC(tag _d.PdfObjectName) *ContentCreator {
	_dcc := ContentStreamOperation{}
	_dcc.Operand = "\u0042\u004d\u0043"
	_dcc.Params = _befb([]_d.PdfObjectName{tag})
	_ccdc._acd = append(_ccdc._acd, &_dcc)
	return _ccdc
}
func (_eed *ContentStreamParser) skipComments() error {
	if _, _fgef := _eed.skipSpaces(); _fgef != nil {
		return _fgef
	}
	_edda := true
	for {
		_gedg, _fdf := _eed._cecb.Peek(1)
		if _fdf != nil {
			_eb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _fdf.Error())
			return _fdf
		}
		if _edda && _gedg[0] != '%' {
			return nil
		}
		_edda = false
		if (_gedg[0] != '\r') && (_gedg[0] != '\n') {
			_eed._cecb.ReadByte()
		} else {
			break
		}
	}
	return _eed.skipComments()
}

// Add_j adds 'j' operand to the content stream: Set the line join style (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_dbd *ContentCreator) Add_j(lineJoinStyle string) *ContentCreator {
	_ggb := ContentStreamOperation{}
	_ggb.Operand = "\u006a"
	_ggb.Params = _befb([]_d.PdfObjectName{_d.PdfObjectName(lineJoinStyle)})
	_dbd._acd = append(_dbd._acd, &_ggb)
	return _dbd
}
func (_abdf *ContentStreamParser) parseHexString() (*_d.PdfObjectString, error) {
	_abdf._cecb.ReadByte()
	_add := []byte("\u0030\u0031\u0032\u003345\u0036\u0037\u0038\u0039\u0061\u0062\u0063\u0064\u0065\u0066\u0041\u0042\u0043\u0044E\u0046")
	var _ddff []byte
	for {
		_abdf.skipSpaces()
		_cefd, _begg := _abdf._cecb.Peek(1)
		if _begg != nil {
			return _d.MakeString(""), _begg
		}
		if _cefd[0] == '>' {
			_abdf._cecb.ReadByte()
			break
		}
		_ccbb, _ := _abdf._cecb.ReadByte()
		if _gc.IndexByte(_add, _ccbb) >= 0 {
			_ddff = append(_ddff, _ccbb)
		}
	}
	if len(_ddff)%2 == 1 {
		_ddff = append(_ddff, '0')
	}
	_gee, _ := _f.DecodeString(string(_ddff))
	return _d.MakeHexString(string(_gee)), nil
}
func (_cgcg *ContentStreamProcessor) getInitialColor(_bge _cg.PdfColorspace) (_cg.PdfColor, error) {
	switch _acdd := _bge.(type) {
	case *_cg.PdfColorspaceDeviceGray:
		return _cg.NewPdfColorDeviceGray(0.0), nil
	case *_cg.PdfColorspaceDeviceRGB:
		return _cg.NewPdfColorDeviceRGB(0.0, 0.0, 0.0), nil
	case *_cg.PdfColorspaceDeviceCMYK:
		return _cg.NewPdfColorDeviceCMYK(0.0, 0.0, 0.0, 1.0), nil
	case *_cg.PdfColorspaceCalGray:
		return _cg.NewPdfColorCalGray(0.0), nil
	case *_cg.PdfColorspaceCalRGB:
		return _cg.NewPdfColorCalRGB(0.0, 0.0, 0.0), nil
	case *_cg.PdfColorspaceLab:
		_bdffa := 0.0
		_gade := 0.0
		_acfb := 0.0
		if _acdd.Range[0] > 0 {
			_bdffa = _acdd.Range[0]
		}
		if _acdd.Range[2] > 0 {
			_gade = _acdd.Range[2]
		}
		return _cg.NewPdfColorLab(_bdffa, _gade, _acfb), nil
	case *_cg.PdfColorspaceICCBased:
		if _acdd.Alternate == nil {
			_eb.Log.Trace("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065d\u0020-\u0020\u0061\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0066\u0061\u006c\u006c\u0020\u0062a\u0063\u006b\u0020\u0028\u004e\u0020\u003d\u0020\u0025\u0064\u0029", _acdd.N)
			if _acdd.N == 1 {
				_eb.Log.Trace("\u0046\u0061\u006c\u006c\u0069\u006e\u0067\u0020\u0062\u0061\u0063k\u0020\u0074\u006f\u0020\u0044\u0065\u0076\u0069\u0063\u0065G\u0072\u0061\u0079")
				return _cgcg.getInitialColor(_cg.NewPdfColorspaceDeviceGray())
			} else if _acdd.N == 3 {
				_eb.Log.Trace("\u0046a\u006c\u006c\u0069\u006eg\u0020\u0062\u0061\u0063\u006b \u0074o\u0020D\u0065\u0076\u0069\u0063\u0065\u0052\u0047B")
				return _cgcg.getInitialColor(_cg.NewPdfColorspaceDeviceRGB())
			} else if _acdd.N == 4 {
				_eb.Log.Trace("\u0046\u0061\u006c\u006c\u0069\u006e\u0067\u0020\u0062\u0061\u0063k\u0020\u0074\u006f\u0020\u0044\u0065\u0076\u0069\u0063\u0065C\u004d\u0059\u004b")
				return _cgcg.getInitialColor(_cg.NewPdfColorspaceDeviceCMYK())
			} else {
				return nil, _a.New("a\u006c\u0074\u0065\u0072\u006e\u0061t\u0065\u0020\u0073\u0070\u0061\u0063e\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0049C\u0043")
			}
		}
		return _cgcg.getInitialColor(_acdd.Alternate)
	case *_cg.PdfColorspaceSpecialIndexed:
		if _acdd.Base == nil {
			return nil, _a.New("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0062\u0061\u0073e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069f\u0069\u0065\u0064")
		}
		return _cgcg.getInitialColor(_acdd.Base)
	case *_cg.PdfColorspaceSpecialSeparation:
		if _acdd.AlternateSpace == nil {
			return nil, _a.New("\u0061\u006ct\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069fi\u0065\u0064")
		}
		return _cgcg.getInitialColor(_acdd.AlternateSpace)
	case *_cg.PdfColorspaceDeviceN:
		if _acdd.AlternateSpace == nil {
			return nil, _a.New("\u0061\u006ct\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069fi\u0065\u0064")
		}
		return _cgcg.getInitialColor(_acdd.AlternateSpace)
	case *_cg.PdfColorspaceSpecialPattern:
		return nil, nil
	}
	_eb.Log.Debug("Un\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0066\u006f\u0072\u0020\u0075\u006e\u006b\u006e\u006fw\u006e \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065:\u0020\u0025T", _bge)
	return nil, _a.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065")
}
func _bdae(_fgd *ContentStreamInlineImage) (*_d.DCTEncoder, error) {
	_egg := _d.NewDCTEncoder()
	_gef := _gc.NewReader(_fgd._fece)
	_ecb, _gdg := _ac.DecodeConfig(_gef)
	if _gdg != nil {
		_eb.Log.Debug("\u0045\u0072\u0072or\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _gdg)
		return nil, _gdg
	}
	switch _ecb.ColorModel {
	case _g.RGBAModel:
		_egg.BitsPerComponent = 8
		_egg.ColorComponents = 3
	case _g.RGBA64Model:
		_egg.BitsPerComponent = 16
		_egg.ColorComponents = 3
	case _g.GrayModel:
		_egg.BitsPerComponent = 8
		_egg.ColorComponents = 1
	case _g.Gray16Model:
		_egg.BitsPerComponent = 16
		_egg.ColorComponents = 1
	case _g.CMYKModel:
		_egg.BitsPerComponent = 8
		_egg.ColorComponents = 4
	case _g.YCbCrModel:
		_egg.BitsPerComponent = 8
		_egg.ColorComponents = 3
	default:
		return nil, _a.New("\u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006d\u006f\u0064\u0065\u006c")
	}
	_egg.Width = _ecb.Width
	_egg.Height = _ecb.Height
	_eb.Log.Trace("\u0044\u0043T\u0020\u0045\u006ec\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076", _egg)
	return _egg, nil
}

// WrapIfNeeded wraps the entire contents within q ... Q.  If unbalanced, then adds extra Qs at the end.
// Only does if needed. Ensures that when adding new content, one start with all states
// in the default condition.
func (_fga *ContentStreamOperations) WrapIfNeeded() *ContentStreamOperations {
	if len(*_fga) == 0 {
		return _fga
	}
	if _fga.isWrapped() {
		return _fga
	}
	*_fga = append([]*ContentStreamOperation{{Operand: "\u0071"}}, *_fga...)
	_b := 0
	for _, _gce := range *_fga {
		if _gce.Operand == "\u0071" {
			_b++
		} else if _gce.Operand == "\u0051" {
			_b--
		}
	}
	for _b > 0 {
		*_fga = append(*_fga, &ContentStreamOperation{Operand: "\u0051"})
		_b--
	}
	return _fga
}
func (_gace *ContentStreamProcessor) getColorspace(_gbe string, _ceb *_cg.PdfPageResources) (_cg.PdfColorspace, error) {
	switch _gbe {
	case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
		return _cg.NewPdfColorspaceDeviceGray(), nil
	case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
		return _cg.NewPdfColorspaceDeviceRGB(), nil
	case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
		return _cg.NewPdfColorspaceDeviceCMYK(), nil
	case "\u0050a\u0074\u0074\u0065\u0072\u006e":
		return _cg.NewPdfColorspaceSpecialPattern(), nil
	}
	_bbda, _feg := _ceb.GetColorspaceByName(_d.PdfObjectName(_gbe))
	if _feg {
		return _bbda, nil
	}
	switch _gbe {
	case "\u0043a\u006c\u0047\u0072\u0061\u0079":
		return _cg.NewPdfColorspaceCalGray(), nil
	case "\u0043\u0061\u006c\u0052\u0047\u0042":
		return _cg.NewPdfColorspaceCalRGB(), nil
	case "\u004c\u0061\u0062":
		return _cg.NewPdfColorspaceLab(), nil
	}
	_eb.Log.Debug("\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063e\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074\u0065\u0064\u003a\u0020\u0025\u0073", _gbe)
	return nil, _fc.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065:\u0020\u0025\u0073", _gbe)
}

// GraphicStateStack represents a stack of GraphicsState.
type GraphicStateStack []GraphicsState

func (_aabc *ContentStreamProcessor) handleCommand_SC(_ecdb *ContentStreamOperation, _bfgf *_cg.PdfPageResources) error {
	_adb := _aabc._cff.ColorspaceStroking
	if len(_ecdb.Params) != _adb.GetNumComponents() {
		_eb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_eb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_ecdb.Params), _adb)
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_bbdd, _gfbc := _adb.ColorFromPdfObjects(_ecdb.Params)
	if _gfbc != nil {
		return _gfbc
	}
	_aabc._cff.ColorStroking = _bbdd
	return nil
}

// SetNonStrokingColor sets the non-stroking `color` where color can be one of
// PdfColorDeviceGray, PdfColorDeviceRGB, or PdfColorDeviceCMYK.
func (_dbf *ContentCreator) SetNonStrokingColor(color _cg.PdfColor) *ContentCreator {
	switch _abg := color.(type) {
	case *_cg.PdfColorDeviceGray:
		_dbf.Add_g(_abg.Val())
	case *_cg.PdfColorDeviceRGB:
		_dbf.Add_rg(_abg.R(), _abg.G(), _abg.B())
	case *_cg.PdfColorDeviceCMYK:
		_dbf.Add_k(_abg.C(), _abg.M(), _abg.Y(), _abg.K())
	default:
		_eb.Log.Debug("\u0053\u0065\u0074N\u006f\u006e\u0053\u0074\u0072\u006f\u006b\u0069\u006e\u0067\u0043\u006f\u006c\u006f\u0072\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020c\u006f\u006c\u006f\u0072\u003a\u0020\u0025\u0054", _abg)
	}
	return _dbf
}
func (_dddg *ContentStreamProcessor) handleCommand_SCN(_cebd *ContentStreamOperation, _fdgc *_cg.PdfPageResources) error {
	_eae := _dddg._cff.ColorspaceStroking
	if !_cded(_eae) {
		if len(_cebd.Params) != _eae.GetNumComponents() {
			_eb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_eb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_cebd.Params), _eae)
			return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_cdeg, _gaef := _eae.ColorFromPdfObjects(_cebd.Params)
	if _gaef != nil {
		return _gaef
	}
	_dddg._cff.ColorStroking = _cdeg
	return nil
}
func (_eee *ContentStreamProcessor) handleCommand_sc(_cefb *ContentStreamOperation, _edae *_cg.PdfPageResources) error {
	_efbg := _eee._cff.ColorspaceNonStroking
	if !_cded(_efbg) {
		if len(_cefb.Params) != _efbg.GetNumComponents() {
			_eb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_eb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_cefb.Params), _efbg)
			return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_adca, _ecg := _efbg.ColorFromPdfObjects(_cefb.Params)
	if _ecg != nil {
		return _ecg
	}
	_eee._cff.ColorNonStroking = _adca
	return nil
}

// Add_m adds 'm' operand to the content stream: Move the current point to (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_ggf *ContentCreator) Add_m(x, y float64) *ContentCreator {
	_bbg := ContentStreamOperation{}
	_bbg.Operand = "\u006d"
	_bbg.Params = _gbfad([]float64{x, y})
	_ggf._acd = append(_ggf._acd, &_bbg)
	return _ggf
}

// Add_BT appends 'BT' operand to the content stream:
// Begin text.
//
// See section 9.4 "Text Objects" and Table 107 (p. 256 PDF32000_2008).
func (_aeff *ContentCreator) Add_BT() *ContentCreator {
	_cee := ContentStreamOperation{}
	_cee.Operand = "\u0042\u0054"
	_aeff._acd = append(_aeff._acd, &_cee)
	return _aeff
}

// Scale applies x-y scaling to the transformation matrix.
func (_ef *ContentCreator) Scale(sx, sy float64) *ContentCreator {
	return _ef.Add_cm(sx, 0, 0, sy, 0, 0)
}

// NewContentStreamParser creates a new instance of the content stream parser from an input content
// stream string.
func NewContentStreamParser(contentStr string) *ContentStreamParser {
	_agg := ContentStreamParser{}
	_dac := _gc.NewBufferString(contentStr + "\u000a")
	_agg._cecb = _cc.NewReader(_dac)
	return &_agg
}
func (_fgcb *ContentStreamProcessor) handleCommand_k(_acec *ContentStreamOperation, _dcfa *_cg.PdfPageResources) error {
	_gggf := _cg.NewPdfColorspaceDeviceCMYK()
	if len(_acec.Params) != _gggf.GetNumComponents() {
		_eb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_eb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_acec.Params), _gggf)
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_cfaa, _bacc := _gggf.ColorFromPdfObjects(_acec.Params)
	if _bacc != nil {
		return _bacc
	}
	_fgcb._cff.ColorspaceNonStroking = _gggf
	_fgcb._cff.ColorNonStroking = _cfaa
	return nil
}

// AddOperand adds a specified operand.
func (_ce *ContentCreator) AddOperand(op ContentStreamOperation) *ContentCreator {
	_ce._acd = append(_ce._acd, &op)
	return _ce
}

// Add_ET appends 'ET' operand to the content stream:
// End text.
//
// See section 9.4 "Text Objects" and Table 107 (p. 256 PDF32000_2008).
func (_ggfc *ContentCreator) Add_ET() *ContentCreator {
	_eeb := ContentStreamOperation{}
	_eeb.Operand = "\u0045\u0054"
	_ggfc._acd = append(_ggfc._acd, &_eeb)
	return _ggfc
}

// Add_re appends 're' operand to the content stream:
// Append a rectangle to the current path as a complete subpath, with lower left corner (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_fcg *ContentCreator) Add_re(x, y, width, height float64) *ContentCreator {
	_eda := ContentStreamOperation{}
	_eda.Operand = "\u0072\u0065"
	_eda.Params = _gbfad([]float64{x, y, width, height})
	_fcg._acd = append(_fcg._acd, &_eda)
	return _fcg
}
func (_adgb *ContentStreamParser) parseObject() (_edeff _d.PdfObject, _dff bool, _ceab error) {
	_adgb.skipSpaces()
	for {
		_gfg, _cbc := _adgb._cecb.Peek(2)
		if _cbc != nil {
			return nil, false, _cbc
		}
		_eb.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_gfg))
		if _gfg[0] == '%' {
			_adgb.skipComments()
			continue
		} else if _gfg[0] == '/' {
			_dbgg, _cgba := _adgb.parseName()
			_eb.Log.Trace("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027", _dbgg)
			return &_dbgg, false, _cgba
		} else if _gfg[0] == '(' {
			_eb.Log.Trace("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!")
			_dfdd, _ebfa := _adgb.parseString()
			return _dfdd, false, _ebfa
		} else if _gfg[0] == '<' && _gfg[1] != '<' {
			_eb.Log.Trace("\u002d\u003e\u0048\u0065\u0078\u0020\u0053\u0074\u0072\u0069\u006e\u0067\u0021")
			_cfed, _gag := _adgb.parseHexString()
			return _cfed, false, _gag
		} else if _gfg[0] == '[' {
			_eb.Log.Trace("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021")
			_edff, _aead := _adgb.parseArray()
			return _edff, false, _aead
		} else if _d.IsFloatDigit(_gfg[0]) || (_gfg[0] == '-' && _d.IsFloatDigit(_gfg[1])) || (_gfg[0] == '+' && _d.IsFloatDigit(_gfg[1])) {
			_eb.Log.Trace("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!")
			_acce, _bbaa := _adgb.parseNumber()
			return _acce, false, _bbaa
		} else if _gfg[0] == '<' && _gfg[1] == '<' {
			_egdc, _dfgc := _adgb.parseDict()
			return _egdc, false, _dfgc
		} else {
			_eb.Log.Trace("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f")
			_gfg, _ = _adgb._cecb.Peek(5)
			_cbg := string(_gfg)
			_eb.Log.Trace("\u0063\u006f\u006e\u0074\u0020\u0050\u0065\u0065\u006b\u0020\u0073\u0074r\u003a\u0020\u0025\u0073", _cbg)
			if (len(_cbg) > 3) && (_cbg[:4] == "\u006e\u0075\u006c\u006c") {
				_eea, _degb := _adgb.parseNull()
				return &_eea, false, _degb
			} else if (len(_cbg) > 4) && (_cbg[:5] == "\u0066\u0061\u006cs\u0065") {
				_efga, _fae := _adgb.parseBool()
				return &_efga, false, _fae
			} else if (len(_cbg) > 3) && (_cbg[:4] == "\u0074\u0072\u0075\u0065") {
				_cedd, _gda := _adgb.parseBool()
				return &_cedd, false, _gda
			}
			_bdff, _eefb := _adgb.parseOperand()
			if _eefb != nil {
				return _bdff, false, _eefb
			}
			if len(_bdff.String()) < 1 {
				return _bdff, false, ErrInvalidOperand
			}
			return _bdff, true, nil
		}
	}
}

// Bytes converts a set of content stream operations to a content stream byte presentation,
// i.e. the kind that can be stored as a PDF stream or string format.
func (_dcf *ContentStreamOperations) Bytes() []byte {
	var _cb _gc.Buffer
	for _, _ab := range *_dcf {
		if _ab == nil {
			continue
		}
		if _ab.Operand == "\u0042\u0049" {
			_cb.WriteString(_ab.Operand + "\u000a")
			_cb.WriteString(_ab.Params[0].WriteString())
		} else {
			for _, _ba := range _ab.Params {
				_cb.WriteString(_ba.WriteString())
				_cb.WriteString("\u0020")
			}
			_cb.WriteString(_ab.Operand + "\u000a")
		}
	}
	return _cb.Bytes()
}
func (_dee *ContentStreamParser) parseString() (*_d.PdfObjectString, error) {
	_dee._cecb.ReadByte()
	var _cfe []byte
	_bee := 1
	for {
		_dfeg, _faa := _dee._cecb.Peek(1)
		if _faa != nil {
			return _d.MakeString(string(_cfe)), _faa
		}
		if _dfeg[0] == '\\' {
			_dee._cecb.ReadByte()
			_deae, _gcgg := _dee._cecb.ReadByte()
			if _gcgg != nil {
				return _d.MakeString(string(_cfe)), _gcgg
			}
			if _d.IsOctalDigit(_deae) {
				_cfd, _adgag := _dee._cecb.Peek(2)
				if _adgag != nil {
					return _d.MakeString(string(_cfe)), _adgag
				}
				var _adf []byte
				_adf = append(_adf, _deae)
				for _, _gcaff := range _cfd {
					if _d.IsOctalDigit(_gcaff) {
						_adf = append(_adf, _gcaff)
					} else {
						break
					}
				}
				_dee._cecb.Discard(len(_adf) - 1)
				_eb.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _adf)
				_eca, _adgag := _c.ParseUint(string(_adf), 8, 32)
				if _adgag != nil {
					return _d.MakeString(string(_cfe)), _adgag
				}
				_cfe = append(_cfe, byte(_eca))
				continue
			}
			switch _deae {
			case 'n':
				_cfe = append(_cfe, '\n')
			case 'r':
				_cfe = append(_cfe, '\r')
			case 't':
				_cfe = append(_cfe, '\t')
			case 'b':
				_cfe = append(_cfe, '\b')
			case 'f':
				_cfe = append(_cfe, '\f')
			case '(':
				_cfe = append(_cfe, '(')
			case ')':
				_cfe = append(_cfe, ')')
			case '\\':
				_cfe = append(_cfe, '\\')
			}
			continue
		} else if _dfeg[0] == '(' {
			_bee++
		} else if _dfeg[0] == ')' {
			_bee--
			if _bee == 0 {
				_dee._cecb.ReadByte()
				break
			}
		}
		_fed, _ := _dee._cecb.ReadByte()
		_cfe = append(_cfe, _fed)
	}
	return _d.MakeString(string(_cfe)), nil
}

// Add_S appends 'S' operand to the content stream: Stroke the path.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_bdb *ContentCreator) Add_S() *ContentCreator {
	_ebb := ContentStreamOperation{}
	_ebb.Operand = "\u0053"
	_bdb._acd = append(_bdb._acd, &_ebb)
	return _bdb
}

type handlerEntry struct {
	Condition HandlerConditionEnum
	Operand   string
	Handler   HandlerFunc
}

// ContentCreator is a builder for PDF content streams.
type ContentCreator struct{ _acd ContentStreamOperations }

func _cec(_dfgf *ContentStreamInlineImage, _dgf *_d.PdfObjectDictionary) (*_d.FlateEncoder, error) {
	_gcd := _d.NewFlateEncoder()
	if _dfgf._cgbc != nil {
		_gcd.SetImage(_dfgf._cgbc)
	}
	if _dgf == nil {
		_dfe := _dfgf.DecodeParms
		if _dfe != nil {
			_baa, _cgfea := _d.GetDict(_dfe)
			if !_cgfea {
				_eb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079 \u0028%\u0054\u0029", _dfe)
				return nil, _fc.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
			_dgf = _baa
		}
	}
	if _dgf == nil {
		return _gcd, nil
	}
	_eb.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _dgf.String())
	_daf := _dgf.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _daf == nil {
		_eb.Log.Debug("E\u0072\u0072o\u0072\u003a\u0020\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0066\u0072\u006f\u006d\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073 \u002d\u0020\u0043\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020\u0077\u0069t\u0068\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u00281\u0029")
	} else {
		_fcf, _caa := _daf.(*_d.PdfObjectInteger)
		if !_caa {
			_eb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _daf)
			return nil, _fc.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_gcd.Predictor = int(*_fcf)
	}
	_daf = _dgf.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _daf != nil {
		_bff, _egd := _daf.(*_d.PdfObjectInteger)
		if !_egd {
			_eb.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _fc.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_gcd.BitsPerComponent = int(*_bff)
	}
	if _gcd.Predictor > 1 {
		_gcd.Columns = 1
		_daf = _dgf.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _daf != nil {
			_abf, _dbb := _daf.(*_d.PdfObjectInteger)
			if !_dbb {
				return nil, _fc.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_gcd.Columns = int(*_abf)
		}
		_gcd.Colors = 1
		_bac := _dgf.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _bac != nil {
			_fda, _edef := _bac.(*_d.PdfObjectInteger)
			if !_edef {
				return nil, _fc.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_gcd.Colors = int(*_fda)
		}
	}
	return _gcd, nil
}

// Add_B appends 'B' operand to the content stream:
// Fill and then stroke the path (nonzero winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_ee *ContentCreator) Add_B() *ContentCreator {
	_ccgd := ContentStreamOperation{}
	_ccgd.Operand = "\u0042"
	_ee._acd = append(_ee._acd, &_ccgd)
	return _ee
}
func _cae(_bcac string) bool { _, _egaf := _aeba[_bcac]; return _egaf }
func (_gbag *ContentStreamParser) skipSpaces() (int, error) {
	_bdcae := 0
	for {
		_agaf, _afca := _gbag._cecb.Peek(1)
		if _afca != nil {
			return 0, _afca
		}
		if _d.IsWhiteSpace(_agaf[0]) {
			_gbag._cecb.ReadByte()
			_bdcae++
		} else {
			break
		}
	}
	return _bdcae, nil
}
func (_aea *ContentStreamInlineImage) toImageBase(_bacd *_cg.PdfPageResources) (*_af.ImageBase, error) {
	if _aea._cgbc != nil {
		return _aea._cgbc, nil
	}
	_ggd := _af.ImageBase{}
	if _aea.Height == nil {
		return nil, _a.New("\u0068e\u0069\u0067\u0068\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_aga, _age := _aea.Height.(*_d.PdfObjectInteger)
	if !_age {
		return nil, _a.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0068e\u0069\u0067\u0068\u0074")
	}
	_ggd.Height = int(*_aga)
	if _aea.Width == nil {
		return nil, _a.New("\u0077\u0069\u0064th\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_efbc, _age := _aea.Width.(*_d.PdfObjectInteger)
	if !_age {
		return nil, _a.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0069\u0064\u0074\u0068")
	}
	_ggd.Width = int(*_efbc)
	_aee, _fcfbb := _aea.IsMask()
	if _fcfbb != nil {
		return nil, _fcfbb
	}
	if _aee {
		_ggd.BitsPerComponent = 1
		_ggd.ColorComponents = 1
	} else {
		if _aea.BitsPerComponent == nil {
			_eb.Log.Debug("\u0049\u006el\u0069\u006e\u0065\u0020\u0042\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u002d\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u0038")
			_ggd.BitsPerComponent = 8
		} else {
			_fcd, _efdc := _aea.BitsPerComponent.(*_d.PdfObjectInteger)
			if !_efdc {
				_eb.Log.Debug("E\u0072\u0072\u006f\u0072\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0062\u0069\u0074\u0073 p\u0065\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u0076al\u0075\u0065,\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _aea.BitsPerComponent)
				return nil, _a.New("\u0042\u0050\u0043\u0020\u0054\u0079\u0070\u0065\u0020e\u0072\u0072\u006f\u0072")
			}
			_ggd.BitsPerComponent = int(*_fcd)
		}
		if _aea.ColorSpace != nil {
			_beg, _bgg := _aea.GetColorSpace(_bacd)
			if _bgg != nil {
				return nil, _bgg
			}
			_ggd.ColorComponents = _beg.GetNumComponents()
		} else {
			_eb.Log.Debug("\u0049\u006el\u0069\u006e\u0065\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006e\u006f\u0074\u0020\u0073p\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u002d\u0020\u0061\u0073\u0073\u0075m\u0069\u006eg\u0020\u0031\u0020\u0063o\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			_ggd.ColorComponents = 1
		}
	}
	if _ceeb, _cdga := _d.GetArray(_aea.Decode); _cdga {
		_ggd.Decode, _fcfbb = _ceeb.ToFloat64Array()
		if _fcfbb != nil {
			return nil, _fcfbb
		}
	}
	_aea._cgbc = &_ggd
	return _aea._cgbc, nil
}
func (_ga *ContentStreamOperations) isWrapped() bool {
	if len(*_ga) < 2 {
		return false
	}
	_fg := 0
	for _, _dc := range *_ga {
		if _dc.Operand == "\u0071" {
			_fg++
		} else if _dc.Operand == "\u0051" {
			_fg--
		} else {
			if _fg < 1 {
				return false
			}
		}
	}
	return _fg == 0
}
func _befb(_ecfa []_d.PdfObjectName) []_d.PdfObject {
	var _gged []_d.PdfObject
	for _, _baac := range _ecfa {
		_gged = append(_gged, _d.MakeName(string(_baac)))
	}
	return _gged
}
func _gbfad(_dgbe []float64) []_d.PdfObject {
	var _cefa []_d.PdfObject
	for _, _ebbb := range _dgbe {
		_cefa = append(_cefa, _d.MakeFloat(_ebbb))
	}
	return _cefa
}

// Add_cm adds 'cm' operation to the content stream: Modifies the current transformation matrix (ctm)
// of the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_bbab *ContentCreator) Add_cm(a, b, c, d, e, f float64) *ContentCreator {
	_agc := ContentStreamOperation{}
	_agc.Operand = "\u0063\u006d"
	_agc.Params = _gbfad([]float64{a, b, c, d, e, f})
	_bbab._acd = append(_bbab._acd, &_agc)
	return _bbab
}

// Add_quote appends "'" operand to the content stream:
// Move to next line and show a string.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_eafc *ContentCreator) Add_quote(textstr _d.PdfObjectString) *ContentCreator {
	_dag := ContentStreamOperation{}
	_dag.Operand = "\u0027"
	_dag.Params = _fgag([]_d.PdfObjectString{textstr})
	_eafc._acd = append(_eafc._acd, &_dag)
	return _eafc
}

// Add_scn_pattern appends 'scn' operand to the content stream for pattern `name`:
// scn with name attribute (for pattern). Syntax: c1 ... cn name scn.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_bdbe *ContentCreator) Add_scn_pattern(name _d.PdfObjectName, c ...float64) *ContentCreator {
	_aa := ContentStreamOperation{}
	_aa.Operand = "\u0073\u0063\u006e"
	_aa.Params = _gbfad(c)
	_aa.Params = append(_aa.Params, _d.MakeName(string(name)))
	_bdbe._acd = append(_bdbe._acd, &_aa)
	return _bdbe
}

// Add_ri adds 'ri' operand to the content stream, which sets the color rendering intent.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_fd *ContentCreator) Add_ri(intent _d.PdfObjectName) *ContentCreator {
	_gf := ContentStreamOperation{}
	_gf.Operand = "\u0072\u0069"
	_gf.Params = _befb([]_d.PdfObjectName{intent})
	_fd._acd = append(_fd._acd, &_gf)
	return _fd
}
func (_ecgg *ContentStreamProcessor) handleCommand_scn(_efcd *ContentStreamOperation, _fab *_cg.PdfPageResources) error {
	_afee := _ecgg._cff.ColorspaceNonStroking
	if !_cded(_afee) {
		if len(_efcd.Params) != _afee.GetNumComponents() {
			_eb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_eb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_efcd.Params), _afee)
			return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_bdde, _eec := _afee.ColorFromPdfObjects(_efcd.Params)
	if _eec != nil {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c \u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0063o\u006co\u0072\u0020\u0066\u0072\u006f\u006d\u0020\u0070\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076 \u0028\u0043\u0053\u0020\u0069\u0073\u0020\u0025\u002b\u0076\u0029", _efcd.Params, _afee)
		return _eec
	}
	_ecgg._cff.ColorNonStroking = _bdde
	return nil
}
func (_dgc *ContentStreamProcessor) handleCommand_G(_afae *ContentStreamOperation, _cdfg *_cg.PdfPageResources) error {
	_edab := _cg.NewPdfColorspaceDeviceGray()
	if len(_afae.Params) != _edab.GetNumComponents() {
		_eb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_eb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_afae.Params), _edab)
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_eeac, _cbcc := _edab.ColorFromPdfObjects(_afae.Params)
	if _cbcc != nil {
		return _cbcc
	}
	_dgc._cff.ColorspaceStroking = _edab
	_dgc._cff.ColorStroking = _eeac
	return nil
}
func _ccdbd(_ecbf []int64) []_d.PdfObject {
	var _agbb []_d.PdfObject
	for _, _gabg := range _ecbf {
		_agbb = append(_agbb, _d.MakeInteger(_gabg))
	}
	return _agbb
}
func (_eefg *ContentStreamProcessor) handleCommand_cs(_gbfa *ContentStreamOperation, _gdad *_cg.PdfPageResources) error {
	if len(_gbfa.Params) < 1 {
		_eb.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0043\u0053\u0020\u0063\u006f\u006d\u006d\u0061\u006e\u0064\u002c\u0020s\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u006f\u0076\u0065\u0072")
		return _a.New("\u0074o\u006f \u0066\u0065\u0077\u0020\u0070a\u0072\u0061m\u0065\u0074\u0065\u0072\u0073")
	}
	if len(_gbfa.Params) > 1 {
		_eb.Log.Debug("\u0043\u0053\u0020\u0063\u006f\u006d\u006d\u0061n\u0064\u0020\u0077it\u0068\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020-\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075i\u006e\u0067")
		return _a.New("\u0074\u006f\u006f\u0020ma\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	}
	_cag, _fecg := _gbfa.Params[0].(*_d.PdfObjectName)
	if !_fecg {
		_eb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020C\u0053\u0020\u0063o\u006d\u006d\u0061n\u0064\u0020w\u0069\u0074\u0068\u0020\u0069\u006ev\u0061li\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u002c\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072")
		return _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_ecf, _fbga := _eefg.getColorspace(string(*_cag), _gdad)
	if _fbga != nil {
		return _fbga
	}
	_eefg._cff.ColorspaceNonStroking = _ecf
	_dcab, _fbga := _eefg.getInitialColor(_ecf)
	if _fbga != nil {
		return _fbga
	}
	_eefg._cff.ColorNonStroking = _dcab
	return nil
}

// Push pushes `gs` on the `gsStack`.
func (_gbcf *GraphicStateStack) Push(gs GraphicsState) { *_gbcf = append(*_gbcf, gs) }

// Add_f appends 'f' operand to the content stream:
// Fill the path using the nonzero winding number rule to determine fill region.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_ffdf *ContentCreator) Add_f() *ContentCreator {
	_bgc := ContentStreamOperation{}
	_bgc.Operand = "\u0066"
	_ffdf._acd = append(_ffdf._acd, &_bgc)
	return _ffdf
}

// Add_quotes appends `"` operand to the content stream:
// Move to next line and show a string, using `aw` and `ac` as word
// and character spacing respectively.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_eeg *ContentCreator) Add_quotes(textstr _d.PdfObjectString, aw, ac float64) *ContentCreator {
	_cedg := ContentStreamOperation{}
	_cedg.Operand = "\u0022"
	_cedg.Params = _gbfad([]float64{aw, ac})
	_cedg.Params = append(_cedg.Params, _fgag([]_d.PdfObjectString{textstr})...)
	_eeg._acd = append(_eeg._acd, &_cedg)
	return _eeg
}

// String is same as Bytes() except returns as a string for convenience.
func (_bf *ContentCreator) String() string { return string(_bf._acd.Bytes()) }

// ExtractText parses and extracts all text data in content streams and returns as a string.
// Does not take into account Encoding table, the output is simply the character codes.
//
// Deprecated: More advanced text extraction is offered in package extractor with character encoding support.
func (_fb *ContentStreamParser) ExtractText() (string, error) {
	_be, _gced := _fb.Parse()
	if _gced != nil {
		return "", _gced
	}
	_afg := false
	_fe, _gcec := float64(-1), float64(-1)
	_dd := ""
	for _, _bb := range *_be {
		if _bb.Operand == "\u0042\u0054" {
			_afg = true
		} else if _bb.Operand == "\u0045\u0054" {
			_afg = false
		}
		if _bb.Operand == "\u0054\u0064" || _bb.Operand == "\u0054\u0044" || _bb.Operand == "\u0054\u002a" {
			_dd += "\u000a"
		}
		if _bb.Operand == "\u0054\u006d" {
			if len(_bb.Params) != 6 {
				continue
			}
			_bcb, _baf := _bb.Params[4].(*_d.PdfObjectFloat)
			if !_baf {
				_ge, _gaf := _bb.Params[4].(*_d.PdfObjectInteger)
				if !_gaf {
					continue
				}
				_bcb = _d.MakeFloat(float64(*_ge))
			}
			_gae, _baf := _bb.Params[5].(*_d.PdfObjectFloat)
			if !_baf {
				_ddf, _ed := _bb.Params[5].(*_d.PdfObjectInteger)
				if !_ed {
					continue
				}
				_gae = _d.MakeFloat(float64(*_ddf))
			}
			if _gcec == -1 {
				_gcec = float64(*_gae)
			} else if _gcec > float64(*_gae) {
				_dd += "\u000a"
				_fe = float64(*_bcb)
				_gcec = float64(*_gae)
				continue
			}
			if _fe == -1 {
				_fe = float64(*_bcb)
			} else if _fe < float64(*_bcb) {
				_dd += "\u0009"
				_fe = float64(*_bcb)
			}
		}
		if _afg && _bb.Operand == "\u0054\u004a" {
			if len(_bb.Params) < 1 {
				continue
			}
			_cbb, _cd := _bb.Params[0].(*_d.PdfObjectArray)
			if !_cd {
				return "", _fc.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0070\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0020\u0074y\u0070\u0065\u002c\u0020\u006e\u006f\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _bb.Params[0])
			}
			for _, _db := range _cbb.Elements() {
				switch _fgb := _db.(type) {
				case *_d.PdfObjectString:
					_dd += _fgb.Str()
				case *_d.PdfObjectFloat:
					if *_fgb < -100 {
						_dd += "\u0020"
					}
				case *_d.PdfObjectInteger:
					if *_fgb < -100 {
						_dd += "\u0020"
					}
				}
			}
		} else if _afg && _bb.Operand == "\u0054\u006a" {
			if len(_bb.Params) < 1 {
				continue
			}
			_da, _bd := _bb.Params[0].(*_d.PdfObjectString)
			if !_bd {
				return "", _fc.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072\u0020\u0074\u0079p\u0065\u002c\u0020\u006e\u006f\u0074\u0020\u0073\u0074\u0072\u0069\u006e\u0067 \u0028\u0025\u0054\u0029", _bb.Params[0])
			}
			_dd += _da.Str()
		}
	}
	return _dd, nil
}

// NewContentStreamProcessor returns a new ContentStreamProcessor for operations `ops`.
func NewContentStreamProcessor(ops []*ContentStreamOperation) *ContentStreamProcessor {
	_ggea := ContentStreamProcessor{}
	_ggea._afb = GraphicStateStack{}
	_bfbd := GraphicsState{}
	_ggea._cff = _bfbd
	_ggea._gfce = []handlerEntry{}
	_ggea._beb = 0
	_ggea._aaba = ops
	return &_ggea
}

// String returns `ops.Bytes()` as a string.
func (_ebg *ContentStreamOperations) String() string { return string(_ebg.Bytes()) }

// ToImage exports the inline image to Image which can be transformed or exported easily.
// Page resources are needed to look up colorspace information.
func (_aac *ContentStreamInlineImage) ToImage(resources *_cg.PdfPageResources) (*_cg.Image, error) {
	_def, _aeffg := _aac.toImageBase(resources)
	if _aeffg != nil {
		return nil, _aeffg
	}
	_abcf, _aeffg := _gcf(_aac)
	if _aeffg != nil {
		return nil, _aeffg
	}
	_dabe, _dege := _d.GetDict(_aac.DecodeParms)
	if _dege {
		_abcf.UpdateParams(_dabe)
	}
	_eb.Log.Trace("\u0065n\u0063o\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076\u0020\u0025\u0054", _abcf, _abcf)
	_eb.Log.Trace("\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065:\u0020\u0025\u002b\u0076", _aac)
	_gde, _aeffg := _abcf.DecodeBytes(_aac._fece)
	if _aeffg != nil {
		return nil, _aeffg
	}
	_fad := &_cg.Image{Width: int64(_def.Width), Height: int64(_def.Height), BitsPerComponent: int64(_def.BitsPerComponent), ColorComponents: _def.ColorComponents, Data: _gde}
	if len(_def.Decode) > 0 {
		for _dcdb := 0; _dcdb < len(_def.Decode); _dcdb++ {
			_def.Decode[_dcdb] *= float64((int(1) << uint(_def.BitsPerComponent)) - 1)
		}
		_fad.SetDecode(_def.Decode)
	}
	return _fad, nil
}

// Add_f_starred appends 'f*' operand to the content stream.
// f*: Fill the path using the even-odd rule to determine fill region.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_fdg *ContentCreator) Add_f_starred() *ContentCreator {
	_edd := ContentStreamOperation{}
	_edd.Operand = "\u0066\u002a"
	_fdg._acd = append(_fdg._acd, &_edd)
	return _fdg
}

// AddHandler adds a new ContentStreamProcessor `handler` of type `condition` for `operand`.
func (_fac *ContentStreamProcessor) AddHandler(condition HandlerConditionEnum, operand string, handler HandlerFunc) {
	_eabf := handlerEntry{}
	_eabf.Condition = condition
	_eabf.Operand = operand
	_eabf.Handler = handler
	_fac._gfce = append(_fac._gfce, _eabf)
}

// Parse parses all commands in content stream, returning a list of operation data.
func (_abga *ContentStreamParser) Parse() (*ContentStreamOperations, error) {
	_bad := ContentStreamOperations{}
	for {
		_gedd := ContentStreamOperation{}
		for {
			_gcef, _dba, _bgca := _abga.parseObject()
			if _bgca != nil {
				if _bgca == _ccg.EOF {
					return &_bad, nil
				}
				return &_bad, _bgca
			}
			if _dba {
				_gedd.Operand, _ = _d.GetStringVal(_gcef)
				_bad = append(_bad, &_gedd)
				break
			} else {
				_gedd.Params = append(_gedd.Params, _gcef)
			}
		}
		if _gedd.Operand == "\u0042\u0049" {
			_caae, _fefe := _abga.ParseInlineImage()
			if _fefe != nil {
				return &_bad, _fefe
			}
			_gedd.Params = append(_gedd.Params, _caae)
		}
	}
}

var _aeba = map[string]struct{}{"\u0062": struct{}{}, "\u0042": struct{}{}, "\u0062\u002a": struct{}{}, "\u0042\u002a": struct{}{}, "\u0042\u0044\u0043": struct{}{}, "\u0042\u0049": struct{}{}, "\u0042\u004d\u0043": struct{}{}, "\u0042\u0054": struct{}{}, "\u0042\u0058": struct{}{}, "\u0063": struct{}{}, "\u0063\u006d": struct{}{}, "\u0043\u0053": struct{}{}, "\u0063\u0073": struct{}{}, "\u0064": struct{}{}, "\u0064\u0030": struct{}{}, "\u0064\u0031": struct{}{}, "\u0044\u006f": struct{}{}, "\u0044\u0050": struct{}{}, "\u0045\u0049": struct{}{}, "\u0045\u004d\u0043": struct{}{}, "\u0045\u0054": struct{}{}, "\u0045\u0058": struct{}{}, "\u0066": struct{}{}, "\u0046": struct{}{}, "\u0066\u002a": struct{}{}, "\u0047": struct{}{}, "\u0067": struct{}{}, "\u0067\u0073": struct{}{}, "\u0068": struct{}{}, "\u0069": struct{}{}, "\u0049\u0044": struct{}{}, "\u006a": struct{}{}, "\u004a": struct{}{}, "\u004b": struct{}{}, "\u006b": struct{}{}, "\u006c": struct{}{}, "\u006d": struct{}{}, "\u004d": struct{}{}, "\u004d\u0050": struct{}{}, "\u006e": struct{}{}, "\u0071": struct{}{}, "\u0051": struct{}{}, "\u0072\u0065": struct{}{}, "\u0052\u0047": struct{}{}, "\u0072\u0067": struct{}{}, "\u0072\u0069": struct{}{}, "\u0073": struct{}{}, "\u0053": struct{}{}, "\u0053\u0043": struct{}{}, "\u0073\u0063": struct{}{}, "\u0053\u0043\u004e": struct{}{}, "\u0073\u0063\u006e": struct{}{}, "\u0073\u0068": struct{}{}, "\u0054\u002a": struct{}{}, "\u0054\u0063": struct{}{}, "\u0054\u0064": struct{}{}, "\u0054\u0044": struct{}{}, "\u0054\u0066": struct{}{}, "\u0054\u006a": struct{}{}, "\u0054\u004a": struct{}{}, "\u0054\u004c": struct{}{}, "\u0054\u006d": struct{}{}, "\u0054\u0072": struct{}{}, "\u0054\u0073": struct{}{}, "\u0054\u0077": struct{}{}, "\u0054\u007a": struct{}{}, "\u0076": struct{}{}, "\u0077": struct{}{}, "\u0057": struct{}{}, "\u0057\u002a": struct{}{}, "\u0079": struct{}{}, "\u0027": struct{}{}, "\u0022": struct{}{}}
