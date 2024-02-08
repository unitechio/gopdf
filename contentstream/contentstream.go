package contentstream

import (
	_af "bufio"
	_ge "bytes"
	_a "encoding/hex"
	_b "errors"
	_ag "fmt"
	_f "image/color"
	_fg "image/jpeg"
	_ec "io"
	_ff "math"
	_g "regexp"
	_e "strconv"

	_db "bitbucket.org/shenghui0779/gopdf/common"
	_dd "bitbucket.org/shenghui0779/gopdf/core"
	_ad "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_bc "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_ba "bitbucket.org/shenghui0779/gopdf/model"
)

// Add_j adds 'j' operand to the content stream: Set the line join style (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_dc *ContentCreator) Add_j(lineJoinStyle string) *ContentCreator {
	_cd := ContentStreamOperation{}
	_cd.Operand = "\u006a"
	_cd.Params = _adbc([]_dd.PdfObjectName{_dd.PdfObjectName(lineJoinStyle)})
	_dc._bg = append(_dc._bg, &_cd)
	return _dc
}

// Add_EMC appends 'EMC' operand to the content stream:
// Ends a marked-content sequence.
//
// See section 14.6 "Marked Content" and Table 320 (p. 561 PDF32000_2008).
func (_dbce *ContentCreator) Add_EMC() *ContentCreator {
	_fdfb := ContentStreamOperation{}
	_fdfb.Operand = "\u0045\u004d\u0043"
	_dbce._bg = append(_dbce._bg, &_fdfb)
	return _dbce
}
func (_cgae *ContentStreamParser) parseNumber() (_dd.PdfObject, error) {
	return _dd.ParseNumber(_cgae._efe)
}

// ContentCreator is a builder for PDF content streams.
type ContentCreator struct{ _bg ContentStreamOperations }

// SetStrokingColor sets the stroking `color` where color can be one of
// PdfColorDeviceGray, PdfColorDeviceRGB, or PdfColorDeviceCMYK.
func (_gg *ContentCreator) SetStrokingColor(color _ba.PdfColor) *ContentCreator {
	switch _ecff := color.(type) {
	case *_ba.PdfColorDeviceGray:
		_gg.Add_G(_ecff.Val())
	case *_ba.PdfColorDeviceRGB:
		_gg.Add_RG(_ecff.R(), _ecff.G(), _ecff.B())
	case *_ba.PdfColorDeviceCMYK:
		_gg.Add_K(_ecff.C(), _ecff.M(), _ecff.Y(), _ecff.K())
	case *_ba.PdfColorPatternType2:
		_gg.Add_CS(*_dd.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
		_gg.Add_SCN_pattern(_ecff.PatternName)
	case *_ba.PdfColorPatternType3:
		_gg.Add_CS(*_dd.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
		_gg.Add_SCN_pattern(_ecff.PatternName)
	default:
		_db.Log.Debug("\u0053\u0065\u0074\u0053\u0074\u0072\u006f\u006b\u0069\u006e\u0067\u0043\u006fl\u006f\u0072\u003a\u0020\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006fr\u003a\u0020\u0025\u0054", _ecff)
	}
	return _gg
}

// Add_SC appends 'SC' operand to the content stream:
// Set color for stroking operations.  Input: c1, ..., cn.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_ccg *ContentCreator) Add_SC(c ...float64) *ContentCreator {
	_febf := ContentStreamOperation{}
	_febf.Operand = "\u0053\u0043"
	_febf.Params = _bdce(c)
	_ccg._bg = append(_ccg._bg, &_febf)
	return _ccg
}

// Add_v appends 'v' operand to the content stream: Append a Bezier curve to the current path from the
// current point to (x3,y3) with the current point and (x2,y2) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_cdg *ContentCreator) Add_v(x2, y2, x3, y3 float64) *ContentCreator {
	_gf := ContentStreamOperation{}
	_gf.Operand = "\u0076"
	_gf.Params = _bdce([]float64{x2, y2, x3, y3})
	_cdg._bg = append(_cdg._bg, &_gf)
	return _cdg
}

// Process processes the entire list of operations. Maintains the graphics state that is passed to any
// handlers that are triggered during processing (either on specific operators or all).
func (_fff *ContentStreamProcessor) Process(resources *_ba.PdfPageResources) error {
	_fff._gea.ColorspaceStroking = _ba.NewPdfColorspaceDeviceGray()
	_fff._gea.ColorspaceNonStroking = _ba.NewPdfColorspaceDeviceGray()
	_fff._gea.ColorStroking = _ba.NewPdfColorDeviceGray(0)
	_fff._gea.ColorNonStroking = _ba.NewPdfColorDeviceGray(0)
	_fff._gea.CTM = _bc.IdentityMatrix()
	for _, _dge := range _fff._aac {
		var _eca error
		switch _dge.Operand {
		case "\u0071":
			_fff._cee.Push(_fff._gea)
		case "\u0051":
			if len(_fff._cee) == 0 {
				_db.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0060\u0051\u0060\u0020\u006f\u0070e\u0072\u0061\u0074\u006f\u0072\u002e\u0020\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0073\u0020\u0073\u0074\u0061\u0074\u0065 \u0073\u0074\u0061\u0063\u006b\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079.\u0020\u0053\u006bi\u0070\u0070\u0069\u006e\u0067\u002e")
				continue
			}
			_fff._gea = _fff._cee.Pop()
		case "\u0043\u0053":
			_eca = _fff.handleCommand_CS(_dge, resources)
		case "\u0063\u0073":
			_eca = _fff.handleCommand_cs(_dge, resources)
		case "\u0053\u0043":
			_eca = _fff.handleCommand_SC(_dge, resources)
		case "\u0053\u0043\u004e":
			_eca = _fff.handleCommand_SCN(_dge, resources)
		case "\u0073\u0063":
			_eca = _fff.handleCommand_sc(_dge, resources)
		case "\u0073\u0063\u006e":
			_eca = _fff.handleCommand_scn(_dge, resources)
		case "\u0047":
			_eca = _fff.handleCommand_G(_dge, resources)
		case "\u0067":
			_eca = _fff.handleCommand_g(_dge, resources)
		case "\u0052\u0047":
			_eca = _fff.handleCommand_RG(_dge, resources)
		case "\u0072\u0067":
			_eca = _fff.handleCommand_rg(_dge, resources)
		case "\u004b":
			_eca = _fff.handleCommand_K(_dge, resources)
		case "\u006b":
			_eca = _fff.handleCommand_k(_dge, resources)
		case "\u0063\u006d":
			_eca = _fff.handleCommand_cm(_dge, resources)
		}
		if _eca != nil {
			_db.Log.Debug("\u0050\u0072\u006f\u0063\u0065\u0073s\u006f\u0072\u0020\u0068\u0061\u006e\u0064\u006c\u0069\u006e\u0067\u0020\u0065r\u0072\u006f\u0072\u0020\u0028\u0025\u0073)\u003a\u0020\u0025\u0076", _dge.Operand, _eca)
			_db.Log.Debug("\u004f\u0070\u0065r\u0061\u006e\u0064\u003a\u0020\u0025\u0023\u0076", _dge.Operand)
			return _eca
		}
		for _, _becc := range _fff._efbe {
			var _dbcb error
			if _becc.Condition.All() {
				_dbcb = _becc.Handler(_dge, _fff._gea, resources)
			} else if _becc.Condition.Operand() && _dge.Operand == _becc.Operand {
				_dbcb = _becc.Handler(_dge, _fff._gea, resources)
			}
			if _dbcb != nil {
				_db.Log.Debug("P\u0072\u006f\u0063\u0065\u0073\u0073o\u0072\u0020\u0068\u0061\u006e\u0064\u006c\u0065\u0072 \u0065\u0072\u0072o\u0072:\u0020\u0025\u0076", _dbcb)
				return _dbcb
			}
		}
	}
	return nil
}

// Scale applies x-y scaling to the transformation matrix.
func (_fgac *ContentCreator) Scale(sx, sy float64) *ContentCreator {
	return _fgac.Add_cm(sx, 0, 0, sy, 0, 0)
}
func (_cgeb *ContentStreamParser) skipSpaces() (int, error) {
	_ebgc := 0
	for {
		_cdd, _ggc := _cgeb._efe.Peek(1)
		if _ggc != nil {
			return 0, _ggc
		}
		if _dd.IsWhiteSpace(_cdd[0]) {
			_cgeb._efe.ReadByte()
			_ebgc++
		} else {
			break
		}
	}
	return _ebgc, nil
}

// Add_S appends 'S' operand to the content stream: Stroke the path.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_bfe *ContentCreator) Add_S() *ContentCreator {
	_fbbd := ContentStreamOperation{}
	_fbbd.Operand = "\u0053"
	_bfe._bg = append(_bfe._bg, &_fbbd)
	return _bfe
}
func (_bcga *ContentStreamProcessor) handleCommand_rg(_daefb *ContentStreamOperation, _dddce *_ba.PdfPageResources) error {
	_daec := _ba.NewPdfColorspaceDeviceRGB()
	if len(_daefb.Params) != _daec.GetNumComponents() {
		_db.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_db.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_daefb.Params), _daec)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_dfdc, _cdbg := _daec.ColorFromPdfObjects(_daefb.Params)
	if _cdbg != nil {
		return _cdbg
	}
	_bcga._gea.ColorspaceNonStroking = _daec
	_bcga._gea.ColorNonStroking = _dfdc
	return nil
}

// SetNonStrokingColor sets the non-stroking `color` where color can be one of
// PdfColorDeviceGray, PdfColorDeviceRGB, or PdfColorDeviceCMYK.
func (_degg *ContentCreator) SetNonStrokingColor(color _ba.PdfColor) *ContentCreator {
	switch _fae := color.(type) {
	case *_ba.PdfColorDeviceGray:
		_degg.Add_g(_fae.Val())
	case *_ba.PdfColorDeviceRGB:
		_degg.Add_rg(_fae.R(), _fae.G(), _fae.B())
	case *_ba.PdfColorDeviceCMYK:
		_degg.Add_k(_fae.C(), _fae.M(), _fae.Y(), _fae.K())
	case *_ba.PdfColorPatternType2:
		_degg.Add_cs(*_dd.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
		_degg.Add_scn_pattern(_fae.PatternName)
	case *_ba.PdfColorPatternType3:
		_degg.Add_cs(*_dd.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
		_degg.Add_scn_pattern(_fae.PatternName)
	default:
		_db.Log.Debug("\u0053\u0065\u0074N\u006f\u006e\u0053\u0074\u0072\u006f\u006b\u0069\u006e\u0067\u0043\u006f\u006c\u006f\u0072\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020c\u006f\u006c\u006f\u0072\u003a\u0020\u0025\u0054", _fae)
	}
	return _degg
}

// Add_Ts appends 'Ts' operand to the content stream:
// Set text rise.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_ccf *ContentCreator) Add_Ts(rise float64) *ContentCreator {
	_def := ContentStreamOperation{}
	_def.Operand = "\u0054\u0073"
	_def.Params = _bdce([]float64{rise})
	_ccf._bg = append(_ccf._bg, &_def)
	return _ccf
}

// GraphicsState is a basic graphics state implementation for PDF processing.
// Initially only implementing and tracking a portion of the information specified. Easy to add more.
type GraphicsState struct {
	ColorspaceStroking    _ba.PdfColorspace
	ColorspaceNonStroking _ba.PdfColorspace
	ColorStroking         _ba.PdfColor
	ColorNonStroking      _ba.PdfColor
	CTM                   _bc.Matrix
}

// Add_TD appends 'TD' operand to the content stream:
// Move to start of next line with offset (`tx`, `ty`).
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_aff *ContentCreator) Add_TD(tx, ty float64) *ContentCreator {
	_ac := ContentStreamOperation{}
	_ac.Operand = "\u0054\u0044"
	_ac.Params = _bdce([]float64{tx, ty})
	_aff._bg = append(_aff._bg, &_ac)
	return _aff
}

// Add_B_starred appends 'B*' operand to the content stream:
// Fill and then stroke the path (even-odd rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_fab *ContentCreator) Add_B_starred() *ContentCreator {
	_bec := ContentStreamOperation{}
	_bec.Operand = "\u0042\u002a"
	_fab._bg = append(_fab._bg, &_bec)
	return _fab
}

// Add_sh appends 'sh' operand to the content stream:
// Paints the shape and colour shading described by a shading dictionary specified by `name`,
// subject to the current clipping path
//
// See section 8.7.4 "Shading Patterns" and Table 77 (p. 190 PDF32000_2008).
func (_cbca *ContentCreator) Add_sh(name _dd.PdfObjectName) *ContentCreator {
	_bgg := ContentStreamOperation{}
	_bgg.Operand = "\u0073\u0068"
	_bgg.Params = _adbc([]_dd.PdfObjectName{name})
	_cbca._bg = append(_cbca._bg, &_bgg)
	return _cbca
}
func _ccda(_gba *ContentStreamInlineImage, _dcc *_dd.PdfObjectDictionary) (*_dd.FlateEncoder, error) {
	_cde := _dd.NewFlateEncoder()
	if _gba._feaga != nil {
		_cde.SetImage(_gba._feaga)
	}
	if _dcc == nil {
		_cbbe := _gba.DecodeParms
		if _cbbe != nil {
			_ege, _cfaf := _dd.GetDict(_cbbe)
			if !_cfaf {
				_db.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079 \u0028%\u0054\u0029", _cbbe)
				return nil, _ag.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
			_dcc = _ege
		}
	}
	if _dcc == nil {
		return _cde, nil
	}
	_db.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _dcc.String())
	_bbbd := _dcc.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _bbbd == nil {
		_db.Log.Debug("E\u0072\u0072o\u0072\u003a\u0020\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0066\u0072\u006f\u006d\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073 \u002d\u0020\u0043\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020\u0077\u0069t\u0068\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u00281\u0029")
	} else {
		_bef, _fgf := _bbbd.(*_dd.PdfObjectInteger)
		if !_fgf {
			_db.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _bbbd)
			return nil, _ag.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_cde.Predictor = int(*_bef)
	}
	_bbbd = _dcc.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _bbbd != nil {
		_dga, _gdg := _bbbd.(*_dd.PdfObjectInteger)
		if !_gdg {
			_db.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _ag.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_cde.BitsPerComponent = int(*_dga)
	}
	if _cde.Predictor > 1 {
		_cde.Columns = 1
		_bbbd = _dcc.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _bbbd != nil {
			_eec, _cba := _bbbd.(*_dd.PdfObjectInteger)
			if !_cba {
				return nil, _ag.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_cde.Columns = int(*_eec)
		}
		_cde.Colors = 1
		_geg := _dcc.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _geg != nil {
			_ccea, _cdge := _geg.(*_dd.PdfObjectInteger)
			if !_cdge {
				return nil, _ag.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_cde.Colors = int(*_ccea)
		}
	}
	return _cde, nil
}

// Add_B appends 'B' operand to the content stream:
// Fill and then stroke the path (nonzero winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_gcb *ContentCreator) Add_B() *ContentCreator {
	_bdf := ContentStreamOperation{}
	_bdf.Operand = "\u0042"
	_gcb._bg = append(_gcb._bg, &_bdf)
	return _gcb
}

// ContentStreamProcessor defines a data structure and methods for processing a content stream, keeping track of the
// current graphics state, and allowing external handlers to define their own functions as a part of the processing,
// for example rendering or extracting certain information.
type ContentStreamProcessor struct {
	_cee  GraphicStateStack
	_aac  []*ContentStreamOperation
	_gea  GraphicsState
	_efbe []handlerEntry
	_dag  int
}

// Add_Tj appends 'Tj' operand to the content stream:
// Show a text string.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_egb *ContentCreator) Add_Tj(textstr _dd.PdfObjectString) *ContentCreator {
	_bbgga := ContentStreamOperation{}
	_bbgga.Operand = "\u0054\u006a"
	_bbgga.Params = _gab([]_dd.PdfObjectString{textstr})
	_egb._bg = append(_egb._bg, &_bbgga)
	return _egb
}

// Add_l adds 'l' operand to the content stream:
// Append a straight line segment from the current point to (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_ebc *ContentCreator) Add_l(x, y float64) *ContentCreator {
	_ebg := ContentStreamOperation{}
	_ebg.Operand = "\u006c"
	_ebg.Params = _bdce([]float64{x, y})
	_ebc._bg = append(_ebc._bg, &_ebg)
	return _ebc
}
func _acce(_gadf string) bool {
	_, _fgbc := _ecdd[_gadf]
	return _fgbc
}

// Add_ri adds 'ri' operand to the content stream, which sets the color rendering intent.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_afb *ContentCreator) Add_ri(intent _dd.PdfObjectName) *ContentCreator {
	_faf := ContentStreamOperation{}
	_faf.Operand = "\u0072\u0069"
	_faf.Params = _adbc([]_dd.PdfObjectName{intent})
	_afb._bg = append(_afb._bg, &_faf)
	return _afb
}
func (_gcbd *ContentStreamProcessor) handleCommand_k(_dccff *ContentStreamOperation, _gda *_ba.PdfPageResources) error {
	_fgag := _ba.NewPdfColorspaceDeviceCMYK()
	if len(_dccff.Params) != _fgag.GetNumComponents() {
		_db.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_db.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_dccff.Params), _fgag)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_accc, _fced := _fgag.ColorFromPdfObjects(_dccff.Params)
	if _fced != nil {
		return _fced
	}
	_gcbd._gea.ColorspaceNonStroking = _fgag
	_gcbd._gea.ColorNonStroking = _accc
	return nil
}

// Add_f_starred appends 'f*' operand to the content stream.
// f*: Fill the path using the even-odd rule to determine fill region.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_eg *ContentCreator) Add_f_starred() *ContentCreator {
	_gfc := ContentStreamOperation{}
	_gfc.Operand = "\u0066\u002a"
	_eg._bg = append(_eg._bg, &_gfc)
	return _eg
}

// ContentStreamParser represents a content stream parser for parsing content streams in PDFs.
type ContentStreamParser struct{ _efe *_af.Reader }

func _bgga(_gede []int64) []_dd.PdfObject {
	var _cbdc []_dd.PdfObject
	for _, _addc := range _gede {
		_cbdc = append(_cbdc, _dd.MakeInteger(_addc))
	}
	return _cbdc
}

// Add_BT appends 'BT' operand to the content stream:
// Begin text.
//
// See section 9.4 "Text Objects" and Table 107 (p. 256 PDF32000_2008).
func (_dfd *ContentCreator) Add_BT() *ContentCreator {
	_cbcf := ContentStreamOperation{}
	_cbcf.Operand = "\u0042\u0054"
	_dfd._bg = append(_dfd._bg, &_cbcf)
	return _dfd
}

// Operand returns true if `hce` is equivalent to HandlerConditionEnumOperand.
func (_dfbf HandlerConditionEnum) Operand() bool { return _dfbf == HandlerConditionEnumOperand }

// Add_k appends 'k' operand to the content stream:
// Same as K but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_deg *ContentCreator) Add_k(c, m, y, k float64) *ContentCreator {
	_gaa := ContentStreamOperation{}
	_gaa.Operand = "\u006b"
	_gaa.Params = _bdce([]float64{c, m, y, k})
	_deg._bg = append(_deg._bg, &_gaa)
	return _deg
}

var _ecdd = map[string]struct{}{"\u0062": struct{}{}, "\u0042": struct{}{}, "\u0062\u002a": struct{}{}, "\u0042\u002a": struct{}{}, "\u0042\u0044\u0043": struct{}{}, "\u0042\u0049": struct{}{}, "\u0042\u004d\u0043": struct{}{}, "\u0042\u0054": struct{}{}, "\u0042\u0058": struct{}{}, "\u0063": struct{}{}, "\u0063\u006d": struct{}{}, "\u0043\u0053": struct{}{}, "\u0063\u0073": struct{}{}, "\u0064": struct{}{}, "\u0064\u0030": struct{}{}, "\u0064\u0031": struct{}{}, "\u0044\u006f": struct{}{}, "\u0044\u0050": struct{}{}, "\u0045\u0049": struct{}{}, "\u0045\u004d\u0043": struct{}{}, "\u0045\u0054": struct{}{}, "\u0045\u0058": struct{}{}, "\u0066": struct{}{}, "\u0046": struct{}{}, "\u0066\u002a": struct{}{}, "\u0047": struct{}{}, "\u0067": struct{}{}, "\u0067\u0073": struct{}{}, "\u0068": struct{}{}, "\u0069": struct{}{}, "\u0049\u0044": struct{}{}, "\u006a": struct{}{}, "\u004a": struct{}{}, "\u004b": struct{}{}, "\u006b": struct{}{}, "\u006c": struct{}{}, "\u006d": struct{}{}, "\u004d": struct{}{}, "\u004d\u0050": struct{}{}, "\u006e": struct{}{}, "\u0071": struct{}{}, "\u0051": struct{}{}, "\u0072\u0065": struct{}{}, "\u0052\u0047": struct{}{}, "\u0072\u0067": struct{}{}, "\u0072\u0069": struct{}{}, "\u0073": struct{}{}, "\u0053": struct{}{}, "\u0053\u0043": struct{}{}, "\u0073\u0063": struct{}{}, "\u0053\u0043\u004e": struct{}{}, "\u0073\u0063\u006e": struct{}{}, "\u0073\u0068": struct{}{}, "\u0054\u002a": struct{}{}, "\u0054\u0063": struct{}{}, "\u0054\u0064": struct{}{}, "\u0054\u0044": struct{}{}, "\u0054\u0066": struct{}{}, "\u0054\u006a": struct{}{}, "\u0054\u004a": struct{}{}, "\u0054\u004c": struct{}{}, "\u0054\u006d": struct{}{}, "\u0054\u0072": struct{}{}, "\u0054\u0073": struct{}{}, "\u0054\u0077": struct{}{}, "\u0054\u007a": struct{}{}, "\u0076": struct{}{}, "\u0077": struct{}{}, "\u0057": struct{}{}, "\u0057\u002a": struct{}{}, "\u0079": struct{}{}, "\u0027": struct{}{}, "\u0022": struct{}{}}

// Translate applies a simple x-y translation to the transformation matrix.
func (_be *ContentCreator) Translate(tx, ty float64) *ContentCreator {
	return _be.Add_cm(1, 0, 0, 1, tx, ty)
}

var _agbd = _g.MustCompile("\u005e\u002f\u007b\u0032\u002c\u007d")

func _adbc(_aceee []_dd.PdfObjectName) []_dd.PdfObject {
	var _gceb []_dd.PdfObject
	for _, _gceg := range _aceee {
		_gceb = append(_gceb, _dd.MakeName(string(_gceg)))
	}
	return _gceb
}
func (_dbcde *ContentStreamProcessor) getColorspace(_cdeb string, _cdfa *_ba.PdfPageResources) (_ba.PdfColorspace, error) {
	switch _cdeb {
	case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
		return _ba.NewPdfColorspaceDeviceGray(), nil
	case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
		return _ba.NewPdfColorspaceDeviceRGB(), nil
	case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
		return _ba.NewPdfColorspaceDeviceCMYK(), nil
	case "\u0050a\u0074\u0074\u0065\u0072\u006e":
		return _ba.NewPdfColorspaceSpecialPattern(), nil
	}
	_gcg, _bcgc := _cdfa.GetColorspaceByName(_dd.PdfObjectName(_cdeb))
	if _bcgc {
		return _gcg, nil
	}
	switch _cdeb {
	case "\u0043a\u006c\u0047\u0072\u0061\u0079":
		return _ba.NewPdfColorspaceCalGray(), nil
	case "\u0043\u0061\u006c\u0052\u0047\u0042":
		return _ba.NewPdfColorspaceCalRGB(), nil
	case "\u004c\u0061\u0062":
		return _ba.NewPdfColorspaceLab(), nil
	}
	_db.Log.Debug("\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063e\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074\u0065\u0064\u003a\u0020\u0025\u0073", _cdeb)
	return nil, _ag.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065:\u0020\u0025\u0073", _cdeb)
}

// All returns true if `hce` is equivalent to HandlerConditionEnumAllOperands.
func (_ffag HandlerConditionEnum) All() bool { return _ffag == HandlerConditionEnumAllOperands }
func (_egd *ContentStreamParser) parseArray() (*_dd.PdfObjectArray, error) {
	_bdfa := _dd.MakeArray()
	_egd._efe.ReadByte()
	for {
		_egd.skipSpaces()
		_badb, _ggg := _egd._efe.Peek(1)
		if _ggg != nil {
			return _bdfa, _ggg
		}
		if _badb[0] == ']' {
			_egd._efe.ReadByte()
			break
		}
		_gcf, _, _ggg := _egd.parseObject()
		if _ggg != nil {
			return _bdfa, _ggg
		}
		_bdfa.Append(_gcf)
	}
	return _bdfa, nil
}

// Add_b_starred appends 'b*' operand to the content stream:
// Close, fill and then stroke the path (even-odd winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_fdb *ContentCreator) Add_b_starred() *ContentCreator {
	_de := ContentStreamOperation{}
	_de.Operand = "\u0062\u002a"
	_fdb._bg = append(_fdb._bg, &_de)
	return _fdb
}

// String returns `ops.Bytes()` as a string.
func (_bf *ContentStreamOperations) String() string { return string(_bf.Bytes()) }

// ToImage exports the inline image to Image which can be transformed or exported easily.
// Page resources are needed to look up colorspace information.
func (_bgb *ContentStreamInlineImage) ToImage(resources *_ba.PdfPageResources) (*_ba.Image, error) {
	_bdcb, _gfa := _bgb.toImageBase(resources)
	if _gfa != nil {
		return nil, _gfa
	}
	_cgca, _gfa := _cfc(_bgb)
	if _gfa != nil {
		return nil, _gfa
	}
	_dfb, _bdd := _dd.GetDict(_bgb.DecodeParms)
	if _bdd {
		_cgca.UpdateParams(_dfb)
	}
	_db.Log.Trace("\u0065n\u0063o\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076\u0020\u0025\u0054", _cgca, _cgca)
	_db.Log.Trace("\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065:\u0020\u0025\u002b\u0076", _bgb)
	_gafaa, _gfa := _cgca.DecodeBytes(_bgb._ega)
	if _gfa != nil {
		return nil, _gfa
	}
	_gfgb := &_ba.Image{Width: int64(_bdcb.Width), Height: int64(_bdcb.Height), BitsPerComponent: int64(_bdcb.BitsPerComponent), ColorComponents: _bdcb.ColorComponents, Data: _gafaa}
	if len(_bdcb.Decode) > 0 {
		for _cda := 0; _cda < len(_bdcb.Decode); _cda++ {
			_bdcb.Decode[_cda] *= float64((int(1) << uint(_bdcb.BitsPerComponent)) - 1)
		}
		_gfgb.SetDecode(_bdcb.Decode)
	}
	return _gfgb, nil
}
func (_cbf *ContentStreamProcessor) handleCommand_G(_geea *ContentStreamOperation, _gce *_ba.PdfPageResources) error {
	_fedd := _ba.NewPdfColorspaceDeviceGray()
	if len(_geea.Params) != _fedd.GetNumComponents() {
		_db.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_db.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_geea.Params), _fedd)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_fgcg, _aecc := _fedd.ColorFromPdfObjects(_geea.Params)
	if _aecc != nil {
		return _aecc
	}
	_cbf._gea.ColorspaceStroking = _fedd
	_cbf._gea.ColorStroking = _fgcg
	return nil
}

// NewInlineImageFromImage makes a new content stream inline image object from an image.
func NewInlineImageFromImage(img _ba.Image, encoder _dd.StreamEncoder) (*ContentStreamInlineImage, error) {
	if encoder == nil {
		encoder = _dd.NewRawEncoder()
	}
	encoder.UpdateParams(img.GetParamsDict())
	_dba := ContentStreamInlineImage{}
	if img.ColorComponents == 1 {
		_dba.ColorSpace = _dd.MakeName("\u0047")
	} else if img.ColorComponents == 3 {
		_dba.ColorSpace = _dd.MakeName("\u0052\u0047\u0042")
	} else if img.ColorComponents == 4 {
		_dba.ColorSpace = _dd.MakeName("\u0043\u004d\u0059\u004b")
	} else {
		_db.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006db\u0065\u0072\u0020o\u0066\u0020c\u006f\u006c\u006f\u0072\u0020\u0063o\u006dpo\u006e\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0072\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0064", img.ColorComponents)
		return nil, _b.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020c\u006fl\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073")
	}
	_dba.BitsPerComponent = _dd.MakeInteger(img.BitsPerComponent)
	_dba.Width = _dd.MakeInteger(img.Width)
	_dba.Height = _dd.MakeInteger(img.Height)
	_aee, _gafa := encoder.EncodeBytes(img.Data)
	if _gafa != nil {
		return nil, _gafa
	}
	_dba._ega = _aee
	_cef := encoder.GetFilterName()
	if _cef != _dd.StreamEncodingFilterNameRaw {
		_dba.Filter = _dd.MakeName(_cef)
	}
	return &_dba, nil
}

// Add_Tw appends 'Tw' operand to the content stream:
// Set word spacing.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_bbge *ContentCreator) Add_Tw(wordSpace float64) *ContentCreator {
	_faee := ContentStreamOperation{}
	_faee.Operand = "\u0054\u0077"
	_faee.Params = _bdce([]float64{wordSpace})
	_bbge._bg = append(_bbge._bg, &_faee)
	return _bbge
}

// HandlerConditionEnum represents the type of operand content stream processor (handler).
// The handler may process a single specific named operand or all operands.
type HandlerConditionEnum int

// Add_TJ appends 'TJ' operand to the content stream:
// Show one or more text string. Array of numbers (displacement) and strings.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_ddcf *ContentCreator) Add_TJ(vals ..._dd.PdfObject) *ContentCreator {
	_fgeg := ContentStreamOperation{}
	_fgeg.Operand = "\u0054\u004a"
	_fgeg.Params = []_dd.PdfObject{_dd.MakeArray(vals...)}
	_ddcf._bg = append(_ddcf._bg, &_fgeg)
	return _ddcf
}

// NewContentCreator returns a new initialized ContentCreator.
func NewContentCreator() *ContentCreator {
	_fbb := &ContentCreator{}
	_fbb._bg = ContentStreamOperations{}
	return _fbb
}
func (_ca *ContentStreamOperations) isWrapped() bool {
	if len(*_ca) < 2 {
		return false
	}
	_bca := 0
	for _, _feb := range *_ca {
		if _feb.Operand == "\u0071" {
			_bca++
		} else if _feb.Operand == "\u0051" {
			_bca--
		} else {
			if _bca < 1 {
				return false
			}
		}
	}
	return _bca == 0
}
func (_cdca *ContentStreamParser) parseHexString() (*_dd.PdfObjectString, error) {
	_cdca._efe.ReadByte()
	_dccf := []byte("\u0030\u0031\u0032\u003345\u0036\u0037\u0038\u0039\u0061\u0062\u0063\u0064\u0065\u0066\u0041\u0042\u0043\u0044E\u0046")
	var _bcg []byte
	for {
		_cdca.skipSpaces()
		_fce, _dcfed := _cdca._efe.Peek(1)
		if _dcfed != nil {
			return _dd.MakeString(""), _dcfed
		}
		if _fce[0] == '>' {
			_cdca._efe.ReadByte()
			break
		}
		_dad, _ := _cdca._efe.ReadByte()
		if _ge.IndexByte(_dccf, _dad) >= 0 {
			_bcg = append(_bcg, _dad)
		}
	}
	if len(_bcg)%2 == 1 {
		_bcg = append(_bcg, '0')
	}
	_eaac, _ := _a.DecodeString(string(_bcg))
	return _dd.MakeHexString(string(_eaac)), nil
}

// Add_h appends 'h' operand to the content stream:
// Close the current subpath by adding a line between the current position and the starting position.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_ef *ContentCreator) Add_h() *ContentCreator {
	_dfg := ContentStreamOperation{}
	_dfg.Operand = "\u0068"
	_ef._bg = append(_ef._bg, &_dfg)
	return _ef
}

// NewContentStreamProcessor returns a new ContentStreamProcessor for operations `ops`.
func NewContentStreamProcessor(ops []*ContentStreamOperation) *ContentStreamProcessor {
	_fcfd := ContentStreamProcessor{}
	_fcfd._cee = GraphicStateStack{}
	_gaac := GraphicsState{}
	_fcfd._gea = _gaac
	_fcfd._efbe = []handlerEntry{}
	_fcfd._dag = 0
	_fcfd._aac = ops
	return &_fcfd
}

// Push pushes `gs` on the `gsStack`.
func (_dgab *GraphicStateStack) Push(gs GraphicsState) { *_dgab = append(*_dgab, gs) }

// Add_w adds 'w' operand to the content stream, which sets the line width.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_ecd *ContentCreator) Add_w(lineWidth float64) *ContentCreator {
	_faa := ContentStreamOperation{}
	_faa.Operand = "\u0077"
	_faa.Params = _bdce([]float64{lineWidth})
	_ecd._bg = append(_ecd._bg, &_faa)
	return _ecd
}

// Add_J adds 'J' operand to the content stream: Set the line cap style (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_cgg *ContentCreator) Add_J(lineCapStyle string) *ContentCreator {
	_fec := ContentStreamOperation{}
	_fec.Operand = "\u004a"
	_fec.Params = _adbc([]_dd.PdfObjectName{_dd.PdfObjectName(lineCapStyle)})
	_cgg._bg = append(_cgg._bg, &_fec)
	return _cgg
}
func (_eeg *ContentStreamParser) parseString() (*_dd.PdfObjectString, error) {
	_eeg._efe.ReadByte()
	var _gadec []byte
	_fdbd := 1
	for {
		_bfc, _baff := _eeg._efe.Peek(1)
		if _baff != nil {
			return _dd.MakeString(string(_gadec)), _baff
		}
		if _bfc[0] == '\\' {
			_eeg._efe.ReadByte()
			_gbef, _cca := _eeg._efe.ReadByte()
			if _cca != nil {
				return _dd.MakeString(string(_gadec)), _cca
			}
			if _dd.IsOctalDigit(_gbef) {
				_dddc, _fafg := _eeg._efe.Peek(2)
				if _fafg != nil {
					return _dd.MakeString(string(_gadec)), _fafg
				}
				var _eaec []byte
				_eaec = append(_eaec, _gbef)
				for _, _ecfc := range _dddc {
					if _dd.IsOctalDigit(_ecfc) {
						_eaec = append(_eaec, _ecfc)
					} else {
						break
					}
				}
				_eeg._efe.Discard(len(_eaec) - 1)
				_db.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _eaec)
				_ffdc, _fafg := _e.ParseUint(string(_eaec), 8, 32)
				if _fafg != nil {
					return _dd.MakeString(string(_gadec)), _fafg
				}
				_gadec = append(_gadec, byte(_ffdc))
				continue
			}
			switch _gbef {
			case 'n':
				_gadec = append(_gadec, '\n')
			case 'r':
				_gadec = append(_gadec, '\r')
			case 't':
				_gadec = append(_gadec, '\t')
			case 'b':
				_gadec = append(_gadec, '\b')
			case 'f':
				_gadec = append(_gadec, '\f')
			case '(':
				_gadec = append(_gadec, '(')
			case ')':
				_gadec = append(_gadec, ')')
			case '\\':
				_gadec = append(_gadec, '\\')
			}
			continue
		} else if _bfc[0] == '(' {
			_fdbd++
		} else if _bfc[0] == ')' {
			_fdbd--
			if _fdbd == 0 {
				_eeg._efe.ReadByte()
				break
			}
		}
		_dgg, _ := _eeg._efe.ReadByte()
		_gadec = append(_gadec, _dgg)
	}
	return _dd.MakeString(string(_gadec)), nil
}
func (_cgebc *ContentStreamProcessor) handleCommand_CS(_cgce *ContentStreamOperation, _afc *_ba.PdfPageResources) error {
	if len(_cgce.Params) < 1 {
		_db.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0063\u0073\u0020\u0063\u006f\u006d\u006d\u0061\u006e\u0064\u002c\u0020s\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u006f\u0076\u0065\u0072")
		return _b.New("\u0074o\u006f \u0066\u0065\u0077\u0020\u0070a\u0072\u0061m\u0065\u0074\u0065\u0072\u0073")
	}
	if len(_cgce.Params) > 1 {
		_db.Log.Debug("\u0063\u0073\u0020\u0063\u006f\u006d\u006d\u0061n\u0064\u0020\u0077it\u0068\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020-\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075i\u006e\u0067")
		return _b.New("\u0074\u006f\u006f\u0020ma\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	}
	_adfcd, _cfdb := _cgce.Params[0].(*_dd.PdfObjectName)
	if !_cfdb {
		_db.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020c\u0073\u0020\u0063o\u006d\u006d\u0061n\u0064\u0020w\u0069\u0074\u0068\u0020\u0069\u006ev\u0061li\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u002c\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072")
		return _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_efde, _afa := _cgebc.getColorspace(string(*_adfcd), _afc)
	if _afa != nil {
		return _afa
	}
	_cgebc._gea.ColorspaceStroking = _efde
	_dade, _afa := _cgebc.getInitialColor(_efde)
	if _afa != nil {
		return _afa
	}
	_cgebc._gea.ColorStroking = _dade
	return nil
}
func (_ceg *ContentStreamParser) parseDict() (*_dd.PdfObjectDictionary, error) {
	_db.Log.Trace("\u0052\u0065\u0061\u0064i\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074 \u0073t\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0021")
	_acbd := _dd.MakeDict()
	_fafgb, _ := _ceg._efe.ReadByte()
	if _fafgb != '<' {
		return nil, _b.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	_fafgb, _ = _ceg._efe.ReadByte()
	if _fafgb != '<' {
		return nil, _b.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	for {
		_ceg.skipSpaces()
		_daef, _fecf := _ceg._efe.Peek(2)
		if _fecf != nil {
			return nil, _fecf
		}
		_db.Log.Trace("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_daef), string(_daef))
		if (_daef[0] == '>') && (_daef[1] == '>') {
			_db.Log.Trace("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
			_ceg._efe.ReadByte()
			_ceg._efe.ReadByte()
			break
		}
		_db.Log.Trace("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021")
		_bee, _fecf := _ceg.parseName()
		_db.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _bee)
		if _fecf != nil {
			_db.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073", _fecf)
			return nil, _fecf
		}
		if len(_bee) > 4 && _bee[len(_bee)-4:] == "\u006e\u0075\u006c\u006c" {
			_feabd := _bee[0 : len(_bee)-4]
			_db.Log.Trace("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029", _bee)
			_db.Log.Trace("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c", _feabd)
			_ceg.skipSpaces()
			_dbde, _ := _ceg._efe.Peek(1)
			if _dbde[0] == '/' {
				_acbd.Set(_feabd, _dd.MakeNull())
				continue
			}
		}
		_ceg.skipSpaces()
		_abbc, _, _fecf := _ceg.parseObject()
		if _fecf != nil {
			return nil, _fecf
		}
		_acbd.Set(_bee, _abbc)
		_db.Log.Trace("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073", _bee, _abbc.String())
	}
	return _acbd, nil
}
func (_fdbf *ContentStreamParser) parseName() (_dd.PdfObjectName, error) {
	_baf := ""
	_gbf := false
	for {
		_aggf, _ddb := _fdbf._efe.Peek(1)
		if _ddb == _ec.EOF {
			break
		}
		if _ddb != nil {
			return _dd.PdfObjectName(_baf), _ddb
		}
		if !_gbf {
			if _aggf[0] == '/' {
				_gbf = true
				_fdbf._efe.ReadByte()
			} else {
				_db.Log.Error("N\u0061\u006d\u0065\u0020\u0073\u0074a\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069\u0074h\u0020\u0025\u0073 \u0028%\u0020\u0078\u0029", _aggf, _aggf)
				return _dd.PdfObjectName(_baf), _ag.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _aggf[0])
			}
		} else {
			if _dd.IsWhiteSpace(_aggf[0]) {
				break
			} else if (_aggf[0] == '/') || (_aggf[0] == '[') || (_aggf[0] == '(') || (_aggf[0] == ']') || (_aggf[0] == '<') || (_aggf[0] == '>') {
				break
			} else if _aggf[0] == '#' {
				_aaag, _gdc := _fdbf._efe.Peek(3)
				if _gdc != nil {
					return _dd.PdfObjectName(_baf), _gdc
				}
				_fdbf._efe.Discard(3)
				_egcf, _gdc := _a.DecodeString(string(_aaag[1:3]))
				if _gdc != nil {
					return _dd.PdfObjectName(_baf), _gdc
				}
				_baf += string(_egcf)
			} else {
				_egaf, _ := _fdbf._efe.ReadByte()
				_baf += string(_egaf)
			}
		}
	}
	return _dd.PdfObjectName(_baf), nil
}

// GraphicStateStack represents a stack of GraphicsState.
type GraphicStateStack []GraphicsState
type handlerEntry struct {
	Condition HandlerConditionEnum
	Operand   string
	Handler   HandlerFunc
}

// Add_Tr appends 'Tr' operand to the content stream:
// Set text rendering mode.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_ged *ContentCreator) Add_Tr(render int64) *ContentCreator {
	_abfb := ContentStreamOperation{}
	_abfb.Operand = "\u0054\u0072"
	_abfb.Params = _bgga([]int64{render})
	_ged._bg = append(_ged._bg, &_abfb)
	return _ged
}
func (_ebge *ContentStreamParser) parseNull() (_dd.PdfObjectNull, error) {
	_, _effc := _ebge._efe.Discard(4)
	return _dd.PdfObjectNull{}, _effc
}

// AddOperand adds a specified operand.
func (_afg *ContentCreator) AddOperand(op ContentStreamOperation) *ContentCreator {
	_afg._bg = append(_afg._bg, &op)
	return _afg
}

// WrapIfNeeded wraps the entire contents within q ... Q.  If unbalanced, then adds extra Qs at the end.
// Only does if needed. Ensures that when adding new content, one start with all states
// in the default condition.
func (_dbc *ContentStreamOperations) WrapIfNeeded() *ContentStreamOperations {
	if len(*_dbc) == 0 {
		return _dbc
	}
	if _dbc.isWrapped() {
		return _dbc
	}
	*_dbc = append([]*ContentStreamOperation{{Operand: "\u0071"}}, *_dbc...)
	_fgb := 0
	for _, _cb := range *_dbc {
		if _cb.Operand == "\u0071" {
			_fgb++
		} else if _cb.Operand == "\u0051" {
			_fgb--
		}
	}
	for _fgb > 0 {
		*_dbc = append(*_dbc, &ContentStreamOperation{Operand: "\u0051"})
		_fgb--
	}
	return _dbc
}
func (_gfb *ContentStreamProcessor) handleCommand_sc(_dcfa *ContentStreamOperation, _bbeb *_ba.PdfPageResources) error {
	_gag := _gfb._gea.ColorspaceNonStroking
	if !_fgaaf(_gag) {
		if len(_dcfa.Params) != _gag.GetNumComponents() {
			_db.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_db.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_dcfa.Params), _gag)
			return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_befe, _gadc := _gag.ColorFromPdfObjects(_dcfa.Params)
	if _gadc != nil {
		return _gadc
	}
	_gfb._gea.ColorNonStroking = _befe
	return nil
}
func (_ddd *ContentStreamInlineImage) toImageBase(_gcce *_ba.PdfPageResources) (*_ad.ImageBase, error) {
	if _ddd._feaga != nil {
		return _ddd._feaga, nil
	}
	_dbe := _ad.ImageBase{}
	if _ddd.Height == nil {
		return nil, _b.New("\u0068e\u0069\u0067\u0068\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_ecg, _afbb := _ddd.Height.(*_dd.PdfObjectInteger)
	if !_afbb {
		return nil, _b.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0068e\u0069\u0067\u0068\u0074")
	}
	_dbe.Height = int(*_ecg)
	if _ddd.Width == nil {
		return nil, _b.New("\u0077\u0069\u0064th\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_deae, _afbb := _ddd.Width.(*_dd.PdfObjectInteger)
	if !_afbb {
		return nil, _b.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0069\u0064\u0074\u0068")
	}
	_dbe.Width = int(*_deae)
	_feagf, _abbd := _ddd.IsMask()
	if _abbd != nil {
		return nil, _abbd
	}
	if _feagf {
		_dbe.BitsPerComponent = 1
		_dbe.ColorComponents = 1
	} else {
		if _ddd.BitsPerComponent == nil {
			_db.Log.Debug("\u0049\u006el\u0069\u006e\u0065\u0020\u0042\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u002d\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u0038")
			_dbe.BitsPerComponent = 8
		} else {
			_bge, _edcb := _ddd.BitsPerComponent.(*_dd.PdfObjectInteger)
			if !_edcb {
				_db.Log.Debug("E\u0072\u0072\u006f\u0072\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0062\u0069\u0074\u0073 p\u0065\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u0076al\u0075\u0065,\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _ddd.BitsPerComponent)
				return nil, _b.New("\u0042\u0050\u0043\u0020\u0054\u0079\u0070\u0065\u0020e\u0072\u0072\u006f\u0072")
			}
			_dbe.BitsPerComponent = int(*_bge)
		}
		if _ddd.ColorSpace != nil {
			_feab, _egae := _ddd.GetColorSpace(_gcce)
			if _egae != nil {
				return nil, _egae
			}
			_dbe.ColorComponents = _feab.GetNumComponents()
		} else {
			_db.Log.Debug("\u0049\u006el\u0069\u006e\u0065\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006e\u006f\u0074\u0020\u0073p\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u002d\u0020\u0061\u0073\u0073\u0075m\u0069\u006eg\u0020\u0031\u0020\u0063o\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			_dbe.ColorComponents = 1
		}
	}
	if _bbcg, _ggd := _dd.GetArray(_ddd.Decode); _ggd {
		_dbe.Decode, _abbd = _bbcg.ToFloat64Array()
		if _abbd != nil {
			return nil, _abbd
		}
	}
	_ddd._feaga = &_dbe
	return _ddd._feaga, nil
}
func _ggbb(_bgd *ContentStreamInlineImage) (*_dd.MultiEncoder, error) {
	_fad := _dd.NewMultiEncoder()
	var _egee *_dd.PdfObjectDictionary
	var _dce []_dd.PdfObject
	if _fgee := _bgd.DecodeParms; _fgee != nil {
		_acb, _gbac := _fgee.(*_dd.PdfObjectDictionary)
		if _gbac {
			_egee = _acb
		}
		_fgc, _dec := _fgee.(*_dd.PdfObjectArray)
		if _dec {
			for _, _bgce := range _fgc.Elements() {
				if _eaa, _agg := _bgce.(*_dd.PdfObjectDictionary); _agg {
					_dce = append(_dce, _eaa)
				} else {
					_dce = append(_dce, nil)
				}
			}
		}
	}
	_fged := _bgd.Filter
	if _fged == nil {
		return nil, _ag.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	_gad, _gade := _fged.(*_dd.PdfObjectArray)
	if !_gade {
		return nil, _ag.Errorf("m\u0075\u006c\u0074\u0069\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0063\u0061\u006e\u0020\u006f\u006el\u0079\u0020\u0062\u0065\u0020\u006d\u0061\u0064\u0065\u0020fr\u006f\u006d\u0020a\u0072r\u0061\u0079")
	}
	for _bfg, _deb := range _gad.Elements() {
		_bcad, _dbda := _deb.(*_dd.PdfObjectName)
		if !_dbda {
			return nil, _ag.Errorf("\u006d\u0075l\u0074\u0069\u0020\u0066i\u006c\u0074e\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061 \u006e\u0061\u006d\u0065")
		}
		var _ce _dd.PdfObject
		if _egee != nil {
			_ce = _egee
		} else {
			if len(_dce) > 0 {
				if _bfg >= len(_dce) {
					return nil, _ag.Errorf("\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0073\u0020\u0069\u006e\u0020d\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006d\u0073\u0020a\u0072\u0072\u0061\u0079")
				}
				_ce = _dce[_bfg]
			}
		}
		var _cfb *_dd.PdfObjectDictionary
		if _fbbb, _ffg := _ce.(*_dd.PdfObjectDictionary); _ffg {
			_cfb = _fbbb
		}
		if *_bcad == _dd.StreamEncodingFilterNameFlate || *_bcad == "\u0046\u006c" {
			_bbbg, _gbcg := _ccda(_bgd, _cfb)
			if _gbcg != nil {
				return nil, _gbcg
			}
			_fad.AddEncoder(_bbbg)
		} else if *_bcad == _dd.StreamEncodingFilterNameLZW {
			_ccca, _cccab := _cge(_bgd, _cfb)
			if _cccab != nil {
				return nil, _cccab
			}
			_fad.AddEncoder(_ccca)
		} else if *_bcad == _dd.StreamEncodingFilterNameASCIIHex {
			_ced := _dd.NewASCIIHexEncoder()
			_fad.AddEncoder(_ced)
		} else if *_bcad == _dd.StreamEncodingFilterNameASCII85 || *_bcad == "\u0041\u0038\u0035" {
			_cfbf := _dd.NewASCII85Encoder()
			_fad.AddEncoder(_cfbf)
		} else {
			_db.Log.Error("U\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0066\u0069l\u0074\u0065\u0072\u0020\u0025\u0073", *_bcad)
			return nil, _ag.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u0066\u0069\u006c\u0074er \u0069n \u006d\u0075\u006c\u0074\u0069\u0020\u0066il\u0074\u0065\u0072\u0020\u0061\u0072\u0072a\u0079")
		}
	}
	return _fad, nil
}

// Add_d adds 'd' operand to the content stream: Set the line dash pattern.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_gc *ContentCreator) Add_d(dashArray []int64, dashPhase int64) *ContentCreator {
	_ageg := ContentStreamOperation{}
	_ageg.Operand = "\u0064"
	_ageg.Params = []_dd.PdfObject{}
	_ageg.Params = append(_ageg.Params, _dd.MakeArrayFromIntegers64(dashArray))
	_ageg.Params = append(_ageg.Params, _dd.MakeInteger(dashPhase))
	_gc._bg = append(_gc._bg, &_ageg)
	return _gc
}

// ContentStreamInlineImage is a representation of an inline image in a Content stream. Everything between the BI and EI operands.
// ContentStreamInlineImage implements the core.PdfObject interface although strictly it is not a PDF object.
type ContentStreamInlineImage struct {
	BitsPerComponent _dd.PdfObject
	ColorSpace       _dd.PdfObject
	Decode           _dd.PdfObject
	DecodeParms      _dd.PdfObject
	Filter           _dd.PdfObject
	Height           _dd.PdfObject
	ImageMask        _dd.PdfObject
	Intent           _dd.PdfObject
	Interpolate      _dd.PdfObject
	Width            _dd.PdfObject
	_ega             []byte
	_feaga           *_ad.ImageBase
}

// GetColorSpace returns the colorspace of the inline image.
func (_fbc *ContentStreamInlineImage) GetColorSpace(resources *_ba.PdfPageResources) (_ba.PdfColorspace, error) {
	if _fbc.ColorSpace == nil {
		_db.Log.Debug("\u0049\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076i\u006e\u0067\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u002c\u0020\u0061\u0073\u0073\u0075\u006di\u006e\u0067\u0020\u0047\u0072a\u0079")
		return _ba.NewPdfColorspaceDeviceGray(), nil
	}
	if _fba, _fccg := _fbc.ColorSpace.(*_dd.PdfObjectArray); _fccg {
		return _ccabf(_fba)
	}
	_defg, _eef := _fbc.ColorSpace.(*_dd.PdfObjectName)
	if !_eef {
		_db.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020I\u006e\u0076\u0061\u006c\u0069\u0064 \u006f\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u003b\u0025\u002bv\u0029", _fbc.ColorSpace, _fbc.ColorSpace)
		return nil, _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_defg == "\u0047" || *_defg == "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079" {
		return _ba.NewPdfColorspaceDeviceGray(), nil
	} else if *_defg == "\u0052\u0047\u0042" || *_defg == "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B" {
		return _ba.NewPdfColorspaceDeviceRGB(), nil
	} else if *_defg == "\u0043\u004d\u0059\u004b" || *_defg == "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b" {
		return _ba.NewPdfColorspaceDeviceCMYK(), nil
	} else if *_defg == "\u0049" || *_defg == "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		return nil, _b.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0049\u006e\u0064e\u0078 \u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065")
	} else {
		if resources.ColorSpace == nil {
			_db.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0073", *_defg)
			return nil, _b.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0063o\u006c\u006fr\u0073\u0070\u0061\u0063\u0065")
		}
		_edbb, _gca := resources.GetColorspaceByName(*_defg)
		if !_gca {
			_db.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0073", *_defg)
			return nil, _b.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0063o\u006c\u006fr\u0073\u0070\u0061\u0063\u0065")
		}
		return _edbb, nil
	}
}
func (_bcc *ContentStreamProcessor) handleCommand_SC(_fbgd *ContentStreamOperation, _ebb *_ba.PdfPageResources) error {
	_fbba := _bcc._gea.ColorspaceStroking
	if len(_fbgd.Params) != _fbba.GetNumComponents() {
		_db.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_db.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_fbgd.Params), _fbba)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_edf, _addd := _fbba.ColorFromPdfObjects(_fbgd.Params)
	if _addd != nil {
		return _addd
	}
	_bcc._gea.ColorStroking = _edf
	return nil
}
func (_baee *ContentStreamParser) parseObject() (_gde _dd.PdfObject, _badea bool, _cdgea error) {
	_baee.skipSpaces()
	for {
		_cadd, _adfc := _baee._efe.Peek(2)
		if _adfc != nil {
			return nil, false, _adfc
		}
		_db.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_cadd))
		if _cadd[0] == '%' {
			_baee.skipComments()
			continue
		} else if _cadd[0] == '/' {
			_bdb, _aec := _baee.parseName()
			_db.Log.Trace("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027", _bdb)
			return &_bdb, false, _aec
		} else if _cadd[0] == '(' {
			_db.Log.Trace("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!")
			_defc, _aged := _baee.parseString()
			return _defc, false, _aged
		} else if _cadd[0] == '<' && _cadd[1] != '<' {
			_db.Log.Trace("\u002d\u003e\u0048\u0065\u0078\u0020\u0053\u0074\u0072\u0069\u006e\u0067\u0021")
			_efec, _bdba := _baee.parseHexString()
			return _efec, false, _bdba
		} else if _cadd[0] == '[' {
			_db.Log.Trace("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021")
			_cegf, _adcc := _baee.parseArray()
			return _cegf, false, _adcc
		} else if _dd.IsFloatDigit(_cadd[0]) || (_cadd[0] == '-' && _dd.IsFloatDigit(_cadd[1])) || (_cadd[0] == '+' && _dd.IsFloatDigit(_cadd[1])) {
			_db.Log.Trace("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!")
			_cgf, _defa := _baee.parseNumber()
			return _cgf, false, _defa
		} else if _cadd[0] == '<' && _cadd[1] == '<' {
			_gafg, _ebga := _baee.parseDict()
			return _gafg, false, _ebga
		} else {
			_db.Log.Trace("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f")
			_cadd, _ = _baee._efe.Peek(5)
			_cbg := string(_cadd)
			_db.Log.Trace("\u0063\u006f\u006e\u0074\u0020\u0050\u0065\u0065\u006b\u0020\u0073\u0074r\u003a\u0020\u0025\u0073", _cbg)
			if (len(_cbg) > 3) && (_cbg[:4] == "\u006e\u0075\u006c\u006c") {
				_dgge, _egg := _baee.parseNull()
				return &_dgge, false, _egg
			} else if (len(_cbg) > 4) && (_cbg[:5] == "\u0066\u0061\u006cs\u0065") {
				_dega, _eefe := _baee.parseBool()
				return &_dega, false, _eefe
			} else if (len(_cbg) > 3) && (_cbg[:4] == "\u0074\u0072\u0075\u0065") {
				_bdbc, _efb := _baee.parseBool()
				return &_bdbc, false, _efb
			}
			_bddc, _ffcf := _baee.parseOperand()
			if _ffcf != nil {
				return _bddc, false, _ffcf
			}
			if len(_bddc.String()) < 1 {
				return _bddc, false, ErrInvalidOperand
			}
			return _bddc, true, nil
		}
	}
}

// ParseInlineImage parses an inline image from a content stream, both reading its properties and binary data.
// When called, "BI" has already been read from the stream.  This function
// finishes reading through "EI" and then returns the ContentStreamInlineImage.
func (_aabf *ContentStreamParser) ParseInlineImage() (*ContentStreamInlineImage, error) {
	_cdfg := ContentStreamInlineImage{}
	for {
		_aabf.skipSpaces()
		_fcfa, _cdc, _ceb := _aabf.parseObject()
		if _ceb != nil {
			return nil, _ceb
		}
		if !_cdc {
			_eedd, _efd := _dd.GetName(_fcfa)
			if !_efd {
				_db.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067e\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0020\u0028\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u0029\u0020\u002d\u0020\u0025T", _fcfa)
				return nil, _ag.Errorf("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067e\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0020\u0028\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u0029\u0020\u002d\u0020\u0025T", _fcfa)
			}
			_ffb, _fee, _afd := _aabf.parseObject()
			if _afd != nil {
				return nil, _afd
			}
			if _fee {
				return nil, _ag.Errorf("\u006eo\u0074\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067 \u0061\u006e\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
			}
			switch *_eedd {
			case "\u0042\u0050\u0043", "\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074":
				_cdfg.BitsPerComponent = _ffb
			case "\u0043\u0053", "\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065":
				_cdfg.ColorSpace = _ffb
			case "\u0044", "\u0044\u0065\u0063\u006f\u0064\u0065":
				_cdfg.Decode = _ffb
			case "\u0044\u0050", "D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073":
				_cdfg.DecodeParms = _ffb
			case "\u0046", "\u0046\u0069\u006c\u0074\u0065\u0072":
				_cdfg.Filter = _ffb
			case "\u0048", "\u0048\u0065\u0069\u0067\u0068\u0074":
				_cdfg.Height = _ffb
			case "\u0049\u004d", "\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k":
				_cdfg.ImageMask = _ffb
			case "\u0049\u006e\u0074\u0065\u006e\u0074":
				_cdfg.Intent = _ffb
			case "\u0049", "I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065":
				_cdfg.Interpolate = _ffb
			case "\u0057", "\u0057\u0069\u0064t\u0068":
				_cdfg.Width = _ffb
			case "\u004c\u0065\u006e\u0067\u0074\u0068", "\u0053u\u0062\u0074\u0079\u0070\u0065", "\u0054\u0079\u0070\u0065":
				_db.Log.Debug("\u0049\u0067\u006e\u006fr\u0069\u006e\u0067\u0020\u0069\u006e\u006c\u0069\u006e\u0065 \u0070a\u0072\u0061\u006d\u0065\u0074\u0065\u0072 \u0025\u0073", *_eedd)
			default:
				return nil, _ag.Errorf("\u0075\u006e\u006b\u006e\u006f\u0077n\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0020\u0025\u0073", *_eedd)
			}
		}
		if _cdc {
			_adc, _fafcd := _fcfa.(*_dd.PdfObjectString)
			if !_fafcd {
				return nil, _ag.Errorf("\u0066a\u0069\u006ce\u0064\u0020\u0074o\u0020\u0072\u0065\u0061\u0064\u0020\u0069n\u006c\u0069\u006e\u0065\u0020\u0069m\u0061\u0067\u0065\u0020\u002d\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
			}
			if _adc.Str() == "\u0045\u0049" {
				_db.Log.Trace("\u0049n\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020f\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e\u002e\u002e")
				return &_cdfg, nil
			} else if _adc.Str() == "\u0049\u0044" {
				_db.Log.Trace("\u0049\u0044\u0020\u0073\u0074\u0061\u0072\u0074")
				_ecgc, _cbcg := _aabf._efe.Peek(1)
				if _cbcg != nil {
					return nil, _cbcg
				}
				if _dd.IsWhiteSpace(_ecgc[0]) {
					_aabf._efe.Discard(1)
				}
				_cdfg._ega = []byte{}
				_edd := 0
				var _fadd []byte
				for {
					_afda, _bfa := _aabf._efe.ReadByte()
					if _bfa != nil {
						_db.Log.Debug("\u0055\u006e\u0061\u0062\u006ce\u0020\u0074\u006f\u0020\u0066\u0069\u006e\u0064\u0020\u0065\u006e\u0064\u0020o\u0066\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0045\u0049\u0020\u0069\u006e\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u0061\u0074a")
						return nil, _bfa
					}
					if _edd == 0 {
						if _dd.IsWhiteSpace(_afda) {
							_fadd = []byte{}
							_fadd = append(_fadd, _afda)
							_edd = 1
						} else if _afda == 'E' {
							_fadd = append(_fadd, _afda)
							_edd = 2
						} else {
							_cdfg._ega = append(_cdfg._ega, _afda)
						}
					} else if _edd == 1 {
						_fadd = append(_fadd, _afda)
						if _afda == 'E' {
							_edd = 2
						} else {
							_cdfg._ega = append(_cdfg._ega, _fadd...)
							_fadd = []byte{}
							if _dd.IsWhiteSpace(_afda) {
								_edd = 1
							} else {
								_edd = 0
							}
						}
					} else if _edd == 2 {
						_fadd = append(_fadd, _afda)
						if _afda == 'I' {
							_edd = 3
						} else {
							_cdfg._ega = append(_cdfg._ega, _fadd...)
							_fadd = []byte{}
							_edd = 0
						}
					} else if _edd == 3 {
						_fadd = append(_fadd, _afda)
						if _dd.IsWhiteSpace(_afda) {
							_gccea, _geef := _aabf._efe.Peek(20)
							if _geef != nil && _geef != _ec.EOF {
								return nil, _geef
							}
							_bcf := NewContentStreamParser(string(_gccea))
							_ebfg := true
							for _dgd := 0; _dgd < 3; _dgd++ {
								_eebf, _aaaa, _agb := _bcf.parseObject()
								if _agb != nil {
									if _agb == _ec.EOF {
										break
									}
									_ebfg = false
									continue
								}
								if _aaaa && !_acce(_eebf.String()) {
									_ebfg = false
									break
								}
							}
							if _ebfg {
								if len(_cdfg._ega) > 100 {
									_db.Log.Trace("\u0049\u006d\u0061\u0067\u0065\u0020\u0073\u0074\u0072\u0065\u0061m\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078 \u002e\u002e\u002e", len(_cdfg._ega), _cdfg._ega[:100])
								} else {
									_db.Log.Trace("\u0049\u006d\u0061\u0067e \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025 \u0078", len(_cdfg._ega), _cdfg._ega)
								}
								return &_cdfg, nil
							}
						}
						_cdfg._ega = append(_cdfg._ega, _fadd...)
						_fadd = []byte{}
						_edd = 0
					}
				}
			}
		}
	}
}
func _bdce(_efbf []float64) []_dd.PdfObject {
	var _eedb []_dd.PdfObject
	for _, _gga := range _efbf {
		_eedb = append(_eedb, _dd.MakeFloat(_gga))
	}
	return _eedb
}

// GetEncoder returns the encoder of the inline image.
func (_adbd *ContentStreamInlineImage) GetEncoder() (_dd.StreamEncoder, error) { return _cfc(_adbd) }

// HandlerFunc is the function syntax that the ContentStreamProcessor handler must implement.
type HandlerFunc func(_dgcg *ContentStreamOperation, _fgg GraphicsState, _fbff *_ba.PdfPageResources) error

// Add_q adds 'q' operand to the content stream: Pushes the current graphics state on the stack.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_abb *ContentCreator) Add_q() *ContentCreator {
	_cfa := ContentStreamOperation{}
	_cfa.Operand = "\u0071"
	_abb._bg = append(_abb._bg, &_cfa)
	return _abb
}
func (_faae *ContentStreamProcessor) handleCommand_RG(_aed *ContentStreamOperation, _fcb *_ba.PdfPageResources) error {
	_accf := _ba.NewPdfColorspaceDeviceRGB()
	if len(_aed.Params) != _accf.GetNumComponents() {
		_db.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020R\u0047")
		_db.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_aed.Params), _accf)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_fcbc, _cbcfb := _accf.ColorFromPdfObjects(_aed.Params)
	if _cbcfb != nil {
		return _cbcfb
	}
	_faae._gea.ColorspaceStroking = _accf
	_faae._gea.ColorStroking = _fcbc
	return nil
}
func (_fcf *ContentStreamInlineImage) String() string {
	_bfgb := _ag.Sprintf("I\u006el\u0069\u006e\u0065\u0049\u006d\u0061\u0067\u0065(\u006c\u0065\u006e\u003d%d\u0029\u000a", len(_fcf._ega))
	if _fcf.BitsPerComponent != nil {
		_bfgb += "\u002d\u0020\u0042\u0050\u0043\u0020" + _fcf.BitsPerComponent.WriteString() + "\u000a"
	}
	if _fcf.ColorSpace != nil {
		_bfgb += "\u002d\u0020\u0043S\u0020" + _fcf.ColorSpace.WriteString() + "\u000a"
	}
	if _fcf.Decode != nil {
		_bfgb += "\u002d\u0020\u0044\u0020" + _fcf.Decode.WriteString() + "\u000a"
	}
	if _fcf.DecodeParms != nil {
		_bfgb += "\u002d\u0020\u0044P\u0020" + _fcf.DecodeParms.WriteString() + "\u000a"
	}
	if _fcf.Filter != nil {
		_bfgb += "\u002d\u0020\u0046\u0020" + _fcf.Filter.WriteString() + "\u000a"
	}
	if _fcf.Height != nil {
		_bfgb += "\u002d\u0020\u0048\u0020" + _fcf.Height.WriteString() + "\u000a"
	}
	if _fcf.ImageMask != nil {
		_bfgb += "\u002d\u0020\u0049M\u0020" + _fcf.ImageMask.WriteString() + "\u000a"
	}
	if _fcf.Intent != nil {
		_bfgb += "\u002d \u0049\u006e\u0074\u0065\u006e\u0074 " + _fcf.Intent.WriteString() + "\u000a"
	}
	if _fcf.Interpolate != nil {
		_bfgb += "\u002d\u0020\u0049\u0020" + _fcf.Interpolate.WriteString() + "\u000a"
	}
	if _fcf.Width != nil {
		_bfgb += "\u002d\u0020\u0057\u0020" + _fcf.Width.WriteString() + "\u000a"
	}
	return _bfgb
}

// NewContentStreamParser creates a new instance of the content stream parser from an input content
// stream string.
func NewContentStreamParser(contentStr string) *ContentStreamParser {
	_eae := ContentStreamParser{}
	contentStr = string(_agbd.ReplaceAll([]byte(contentStr), []byte("\u002f")))
	_dff := _ge.NewBufferString(contentStr + "\u000a")
	_eae._efe = _af.NewReader(_dff)
	return &_eae
}

// Add_i adds 'i' operand to the content stream: Set the flatness tolerance in the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_bbe *ContentCreator) Add_i(flatness float64) *ContentCreator {
	_fbf := ContentStreamOperation{}
	_fbf.Operand = "\u0069"
	_fbf.Params = _bdce([]float64{flatness})
	_bbe._bg = append(_bbe._bg, &_fbf)
	return _bbe
}

// Pop pops and returns the topmost GraphicsState off the `gsStack`.
func (_fabf *GraphicStateStack) Pop() GraphicsState {
	_faef := (*_fabf)[len(*_fabf)-1]
	*_fabf = (*_fabf)[:len(*_fabf)-1]
	return _faef
}

// Add_scn_pattern appends 'scn' operand to the content stream for pattern `name`:
// scn with name attribute (for pattern). Syntax: c1 ... cn name scn.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_dfc *ContentCreator) Add_scn_pattern(name _dd.PdfObjectName, c ...float64) *ContentCreator {
	_ecb := ContentStreamOperation{}
	_ecb.Operand = "\u0073\u0063\u006e"
	_ecb.Params = _bdce(c)
	_ecb.Params = append(_ecb.Params, _dd.MakeName(string(name)))
	_dfc._bg = append(_dfc._bg, &_ecb)
	return _dfc
}

// Add_Tm appends 'Tm' operand to the content stream:
// Set the text line matrix.
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_acc *ContentCreator) Add_Tm(a, b, c, d, e, f float64) *ContentCreator {
	_ccce := ContentStreamOperation{}
	_ccce.Operand = "\u0054\u006d"
	_ccce.Params = _bdce([]float64{a, b, c, d, e, f})
	_acc._bg = append(_acc._bg, &_ccce)
	return _acc
}
func (_fdgg *ContentStreamProcessor) getInitialColor(_aef _ba.PdfColorspace) (_ba.PdfColor, error) {
	switch _dfdd := _aef.(type) {
	case *_ba.PdfColorspaceDeviceGray:
		return _ba.NewPdfColorDeviceGray(0.0), nil
	case *_ba.PdfColorspaceDeviceRGB:
		return _ba.NewPdfColorDeviceRGB(0.0, 0.0, 0.0), nil
	case *_ba.PdfColorspaceDeviceCMYK:
		return _ba.NewPdfColorDeviceCMYK(0.0, 0.0, 0.0, 1.0), nil
	case *_ba.PdfColorspaceCalGray:
		return _ba.NewPdfColorCalGray(0.0), nil
	case *_ba.PdfColorspaceCalRGB:
		return _ba.NewPdfColorCalRGB(0.0, 0.0, 0.0), nil
	case *_ba.PdfColorspaceLab:
		_fed := 0.0
		_dbb := 0.0
		_becb := 0.0
		if _dfdd.Range[0] > 0 {
			_fed = _dfdd.Range[0]
		}
		if _dfdd.Range[2] > 0 {
			_dbb = _dfdd.Range[2]
		}
		return _ba.NewPdfColorLab(_fed, _dbb, _becb), nil
	case *_ba.PdfColorspaceICCBased:
		if _dfdd.Alternate == nil {
			_db.Log.Trace("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065d\u0020-\u0020\u0061\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0066\u0061\u006c\u006c\u0020\u0062a\u0063\u006b\u0020\u0028\u004e\u0020\u003d\u0020\u0025\u0064\u0029", _dfdd.N)
			if _dfdd.N == 1 {
				_db.Log.Trace("\u0046\u0061\u006c\u006c\u0069\u006e\u0067\u0020\u0062\u0061\u0063k\u0020\u0074\u006f\u0020\u0044\u0065\u0076\u0069\u0063\u0065G\u0072\u0061\u0079")
				return _fdgg.getInitialColor(_ba.NewPdfColorspaceDeviceGray())
			} else if _dfdd.N == 3 {
				_db.Log.Trace("\u0046a\u006c\u006c\u0069\u006eg\u0020\u0062\u0061\u0063\u006b \u0074o\u0020D\u0065\u0076\u0069\u0063\u0065\u0052\u0047B")
				return _fdgg.getInitialColor(_ba.NewPdfColorspaceDeviceRGB())
			} else if _dfdd.N == 4 {
				_db.Log.Trace("\u0046\u0061\u006c\u006c\u0069\u006e\u0067\u0020\u0062\u0061\u0063k\u0020\u0074\u006f\u0020\u0044\u0065\u0076\u0069\u0063\u0065C\u004d\u0059\u004b")
				return _fdgg.getInitialColor(_ba.NewPdfColorspaceDeviceCMYK())
			} else {
				return nil, _b.New("a\u006c\u0074\u0065\u0072\u006e\u0061t\u0065\u0020\u0073\u0070\u0061\u0063e\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0049C\u0043")
			}
		}
		return _fdgg.getInitialColor(_dfdd.Alternate)
	case *_ba.PdfColorspaceSpecialIndexed:
		if _dfdd.Base == nil {
			return nil, _b.New("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0062\u0061\u0073e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069f\u0069\u0065\u0064")
		}
		return _fdgg.getInitialColor(_dfdd.Base)
	case *_ba.PdfColorspaceSpecialSeparation:
		if _dfdd.AlternateSpace == nil {
			return nil, _b.New("\u0061\u006ct\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069fi\u0065\u0064")
		}
		return _fdgg.getInitialColor(_dfdd.AlternateSpace)
	case *_ba.PdfColorspaceDeviceN:
		if _dfdd.AlternateSpace == nil {
			return nil, _b.New("\u0061\u006ct\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069fi\u0065\u0064")
		}
		return _fdgg.getInitialColor(_dfdd.AlternateSpace)
	case *_ba.PdfColorspaceSpecialPattern:
		return _ba.NewPdfColorPattern(), nil
	}
	_db.Log.Debug("Un\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0066\u006f\u0072\u0020\u0075\u006e\u006b\u006e\u006fw\u006e \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065:\u0020\u0025T", _aef)
	return nil, _b.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065")
}

// HasUnclosedQ checks if all the `q` operator is properly closed by `Q` operator.
func (_ed *ContentStreamOperations) HasUnclosedQ() bool {
	_bab := 0
	for _, _c := range *_ed {
		if _c.Operand == "\u0071" {
			_bab++
		} else if _c.Operand == "\u0051" {
			_bab--
		}
	}
	return _bab != 0
}

var (
	ErrInvalidOperand = _b.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
	ErrEarlyExit      = _b.New("\u0074\u0065\u0072\u006di\u006e\u0061\u0074\u0065\u0020\u0070\u0072\u006f\u0063\u0065s\u0073 \u0065\u0061\u0072\u006c\u0079\u0020\u0065x\u0069\u0074")
)

func (_dbbb *ContentStreamProcessor) handleCommand_scn(_aebd *ContentStreamOperation, _efea *_ba.PdfPageResources) error {
	_fafgc := _dbbb._gea.ColorspaceNonStroking
	if !_fgaaf(_fafgc) {
		if len(_aebd.Params) != _fafgc.GetNumComponents() {
			_db.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_db.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_aebd.Params), _fafgc)
			return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_gdf, _gdfe := _fafgc.ColorFromPdfObjects(_aebd.Params)
	if _gdfe != nil {
		_db.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c \u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0063o\u006co\u0072\u0020\u0066\u0072\u006f\u006d\u0020\u0070\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076 \u0028\u0043\u0053\u0020\u0069\u0073\u0020\u0025\u002b\u0076\u0029", _aebd.Params, _fafgc)
		return _gdfe
	}
	_dbbb._gea.ColorNonStroking = _gdf
	return nil
}

// Add_K appends 'K' operand to the content stream:
// Set the stroking colorspace to DeviceCMYK and sets the c,m,y,k color (0-1 each component).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_bece *ContentCreator) Add_K(c, m, y, k float64) *ContentCreator {
	_ebf := ContentStreamOperation{}
	_ebf.Operand = "\u004b"
	_ebf.Params = _bdce([]float64{c, m, y, k})
	_bece._bg = append(_bece._bg, &_ebf)
	return _bece
}

// ExtractText parses and extracts all text data in content streams and returns as a string.
// Does not take into account Encoding table, the output is simply the character codes.
//
// Deprecated: More advanced text extraction is offered in package extractor with character encoding support.
func (_gae *ContentStreamParser) ExtractText() (string, error) {
	_bb, _bae := _gae.Parse()
	if _bae != nil {
		return "", _bae
	}
	_bag := false
	_eb, _fb := float64(-1), float64(-1)
	_fa := ""
	for _, _ab := range *_bb {
		if _ab.Operand == "\u0042\u0054" {
			_bag = true
		} else if _ab.Operand == "\u0045\u0054" {
			_bag = false
		}
		if _ab.Operand == "\u0054\u0064" || _ab.Operand == "\u0054\u0044" || _ab.Operand == "\u0054\u002a" {
			_fa += "\u000a"
		}
		if _ab.Operand == "\u0054\u006d" {
			if len(_ab.Params) != 6 {
				continue
			}
			_gb, _fef := _ab.Params[4].(*_dd.PdfObjectFloat)
			if !_fef {
				_bd, _cad := _ab.Params[4].(*_dd.PdfObjectInteger)
				if !_cad {
					continue
				}
				_gb = _dd.MakeFloat(float64(*_bd))
			}
			_ee, _fef := _ab.Params[5].(*_dd.PdfObjectFloat)
			if !_fef {
				_bbf, _cce := _ab.Params[5].(*_dd.PdfObjectInteger)
				if !_cce {
					continue
				}
				_ee = _dd.MakeFloat(float64(*_bbf))
			}
			if _fb == -1 {
				_fb = float64(*_ee)
			} else if _fb > float64(*_ee) {
				_fa += "\u000a"
				_eb = float64(*_gb)
				_fb = float64(*_ee)
				continue
			}
			if _eb == -1 {
				_eb = float64(*_gb)
			} else if _eb < float64(*_gb) {
				_fa += "\u0009"
				_eb = float64(*_gb)
			}
		}
		if _bag && _ab.Operand == "\u0054\u004a" {
			if len(_ab.Params) < 1 {
				continue
			}
			_df, _ada := _ab.Params[0].(*_dd.PdfObjectArray)
			if !_ada {
				return "", _ag.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0070\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0020\u0074y\u0070\u0065\u002c\u0020\u006e\u006f\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _ab.Params[0])
			}
			for _, _fga := range _df.Elements() {
				switch _cag := _fga.(type) {
				case *_dd.PdfObjectString:
					_fa += _cag.Str()
				case *_dd.PdfObjectFloat:
					if *_cag < -100 {
						_fa += "\u0020"
					}
				case *_dd.PdfObjectInteger:
					if *_cag < -100 {
						_fa += "\u0020"
					}
				}
			}
		} else if _bag && _ab.Operand == "\u0054\u006a" {
			if len(_ab.Params) < 1 {
				continue
			}
			_baa, _fge := _ab.Params[0].(*_dd.PdfObjectString)
			if !_fge {
				return "", _ag.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072\u0020\u0074\u0079p\u0065\u002c\u0020\u006e\u006f\u0074\u0020\u0073\u0074\u0072\u0069\u006e\u0067 \u0028\u0025\u0054\u0029", _ab.Params[0])
			}
			_fa += _baa.Str()
		}
	}
	return _fa, nil
}
func _ccabf(_ebcf _dd.PdfObject) (_ba.PdfColorspace, error) {
	_gfdf, _fabb := _ebcf.(*_dd.PdfObjectArray)
	if !_fabb {
		_db.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0049\u006ev\u0061\u006c\u0069d\u0020\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020cs\u0020\u006e\u006ft\u0020\u0069n\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025#\u0076\u0029", _ebcf)
		return nil, _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _gfdf.Len() != 4 {
		_db.Log.Debug("\u0045\u0072\u0072\u006f\u0072:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061r\u0072\u0061\u0079\u002c\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0034\u0020\u0028\u0025\u0064\u0029", _gfdf.Len())
		return nil, _b.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_eccb, _fabb := _gfdf.Get(0).(*_dd.PdfObjectName)
	if !_fabb {
		_db.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0066\u0069\u0072s\u0074 \u0065\u006c\u0065\u006de\u006e\u0074 \u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0023\u0076\u0029", *_gfdf)
		return nil, _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_eccb != "\u0049" && *_eccb != "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		_db.Log.Debug("\u0045\u0072r\u006f\u0072\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064 \u0063\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0021\u003d\u0020\u0049\u0020\u0028\u0067\u006f\u0074\u003a\u0020\u0025\u0076\u0029", *_eccb)
		return nil, _b.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_eccb, _fabb = _gfdf.Get(1).(*_dd.PdfObjectName)
	if !_fabb {
		_db.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072r\u0061\u0079\u0020\u0032\u006e\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0061\u0072\u0072a\u0079\u003a\u0020\u0025\u0023v\u0029", *_gfdf)
		return nil, _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_eccb != "\u0047" && *_eccb != "\u0052\u0047\u0042" && *_eccb != "\u0043\u004d\u0059\u004b" && *_eccb != "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079" && *_eccb != "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B" && *_eccb != "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b" {
		_db.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072r\u0061\u0079\u0020\u0032\u006e\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0021\u003d\u0020\u0047\u002f\u0052\u0047\u0042\u002f\u0043\u004d\u0059\u004b\u0020\u0028g\u006f\u0074\u003a\u0020\u0025v\u0029", *_eccb)
		return nil, _b.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_aba := ""
	switch *_eccb {
	case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
		_aba = "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079"
	case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
		_aba = "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B"
	case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
		_aba = "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
	}
	_bcb := _dd.MakeArray(_dd.MakeName("\u0049n\u0064\u0065\u0078\u0065\u0064"), _dd.MakeName(_aba), _gfdf.Get(2), _gfdf.Get(3))
	return _ba.NewPdfColorspaceFromPdfObject(_bcb)
}

// Add_quote appends "'" operand to the content stream:
// Move to next line and show a string.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_beb *ContentCreator) Add_quote(textstr _dd.PdfObjectString) *ContentCreator {
	_bcd := ContentStreamOperation{}
	_bcd.Operand = "\u0027"
	_bcd.Params = _gab([]_dd.PdfObjectString{textstr})
	_beb._bg = append(_beb._bg, &_bcd)
	return _beb
}
func (_gef *ContentStreamProcessor) handleCommand_SCN(_dagd *ContentStreamOperation, _egcb *_ba.PdfPageResources) error {
	_aaea := _gef._gea.ColorspaceStroking
	if !_fgaaf(_aaea) {
		if len(_dagd.Params) != _aaea.GetNumComponents() {
			_db.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_db.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_dagd.Params), _aaea)
			return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_bacb, _eaee := _aaea.ColorFromPdfObjects(_dagd.Params)
	if _eaee != nil {
		return _eaee
	}
	_gef._gea.ColorStroking = _bacb
	return nil
}

// Add_Td appends 'Td' operand to the content stream:
// Move to start of next line with offset (`tx`, `ty`).
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_aeg *ContentCreator) Add_Td(tx, ty float64) *ContentCreator {
	_edg := ContentStreamOperation{}
	_edg.Operand = "\u0054\u0064"
	_edg.Params = _bdce([]float64{tx, ty})
	_aeg._bg = append(_aeg._bg, &_edg)
	return _aeg
}

// Add_s appends 's' operand to the content stream: Close and stroke the path.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_gbe *ContentCreator) Add_s() *ContentCreator {
	_aab := ContentStreamOperation{}
	_aab.Operand = "\u0073"
	_gbe._bg = append(_gbe._bg, &_aab)
	return _gbe
}

// Transform returns coordinates x, y transformed by the CTM.
func (_dfce *GraphicsState) Transform(x, y float64) (float64, float64) {
	return _dfce.CTM.Transform(x, y)
}

// Parse parses all commands in content stream, returning a list of operation data.
func (_cdgc *ContentStreamParser) Parse() (*ContentStreamOperations, error) {
	_gfad := ContentStreamOperations{}
	for {
		_ffc := ContentStreamOperation{}
		for {
			_fbg, _fdbg, _cbac := _cdgc.parseObject()
			if _cbac != nil {
				if _cbac == _ec.EOF {
					return &_gfad, nil
				}
				return &_gfad, _cbac
			}
			if _fdbg {
				_ffc.Operand, _ = _dd.GetStringVal(_fbg)
				_gfad = append(_gfad, &_ffc)
				break
			} else {
				_ffc.Params = append(_ffc.Params, _fbg)
			}
		}
		if _ffc.Operand == "\u0042\u0049" {
			_aeeb, _bage := _cdgc.ParseInlineImage()
			if _bage != nil {
				return &_gfad, _bage
			}
			_ffc.Params = append(_ffc.Params, _aeeb)
		}
	}
}

// Add_y appends 'y' operand to the content stream: Append a Bezier curve to the current path from the
// current point to (x3,y3) with (x1, y1) and (x3,y3) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_da *ContentCreator) Add_y(x1, y1, x3, y3 float64) *ContentCreator {
	_afga := ContentStreamOperation{}
	_afga.Operand = "\u0079"
	_afga.Params = _bdce([]float64{x1, y1, x3, y3})
	_da._bg = append(_da._bg, &_afga)
	return _da
}

// Bytes converts the content stream operations to a content stream byte presentation, i.e. the kind that can be
// stored as a PDF stream or string format.
func (_abd *ContentCreator) Bytes() []byte { return _abd._bg.Bytes() }

// Add_TL appends 'TL' operand to the content stream:
// Set leading.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_dcdd *ContentCreator) Add_TL(leading float64) *ContentCreator {
	_dcfe := ContentStreamOperation{}
	_dcfe.Operand = "\u0054\u004c"
	_dcfe.Params = _bdce([]float64{leading})
	_dcdd._bg = append(_dcdd._bg, &_dcfe)
	return _dcdd
}
func (_dee *ContentStreamProcessor) handleCommand_cs(_cefd *ContentStreamOperation, _cbd *_ba.PdfPageResources) error {
	if len(_cefd.Params) < 1 {
		_db.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0043\u0053\u0020\u0063\u006f\u006d\u006d\u0061\u006e\u0064\u002c\u0020s\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u006f\u0076\u0065\u0072")
		return _b.New("\u0074o\u006f \u0066\u0065\u0077\u0020\u0070a\u0072\u0061m\u0065\u0074\u0065\u0072\u0073")
	}
	if len(_cefd.Params) > 1 {
		_db.Log.Debug("\u0043\u0053\u0020\u0063\u006f\u006d\u006d\u0061n\u0064\u0020\u0077it\u0068\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020-\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075i\u006e\u0067")
		return _b.New("\u0074\u006f\u006f\u0020ma\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	}
	_cdce, _aecg := _cefd.Params[0].(*_dd.PdfObjectName)
	if !_aecg {
		_db.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020C\u0053\u0020\u0063o\u006d\u006d\u0061n\u0064\u0020w\u0069\u0074\u0068\u0020\u0069\u006ev\u0061li\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u002c\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072")
		return _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_aeed, _gcdb := _dee.getColorspace(string(*_cdce), _cbd)
	if _gcdb != nil {
		return _gcdb
	}
	_dee._gea.ColorspaceNonStroking = _aeed
	_ffdf, _gcdb := _dee.getInitialColor(_aeed)
	if _gcdb != nil {
		return _gcdb
	}
	_dee._gea.ColorNonStroking = _ffdf
	return nil
}
func (_bfef *ContentStreamParser) parseOperand() (*_dd.PdfObjectString, error) {
	var _eaed []byte
	for {
		_gddc, _gbd := _bfef._efe.Peek(1)
		if _gbd != nil {
			return _dd.MakeString(string(_eaed)), _gbd
		}
		if _dd.IsDelimiter(_gddc[0]) {
			break
		}
		if _dd.IsWhiteSpace(_gddc[0]) {
			break
		}
		_ebe, _ := _bfef._efe.ReadByte()
		_eaed = append(_eaed, _ebe)
	}
	return _dd.MakeString(string(_eaed)), nil
}
func _dbd(_dae *ContentStreamInlineImage) (*_dd.DCTEncoder, error) {
	_aeb := _dd.NewDCTEncoder()
	_cdf := _ge.NewReader(_dae._ega)
	_edcg, _feag := _fg.DecodeConfig(_cdf)
	if _feag != nil {
		_db.Log.Debug("\u0045\u0072\u0072or\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _feag)
		return nil, _feag
	}
	switch _edcg.ColorModel {
	case _f.RGBAModel:
		_aeb.BitsPerComponent = 8
		_aeb.ColorComponents = 3
	case _f.RGBA64Model:
		_aeb.BitsPerComponent = 16
		_aeb.ColorComponents = 3
	case _f.GrayModel:
		_aeb.BitsPerComponent = 8
		_aeb.ColorComponents = 1
	case _f.Gray16Model:
		_aeb.BitsPerComponent = 16
		_aeb.ColorComponents = 1
	case _f.CMYKModel:
		_aeb.BitsPerComponent = 8
		_aeb.ColorComponents = 4
	case _f.YCbCrModel:
		_aeb.BitsPerComponent = 8
		_aeb.ColorComponents = 3
	default:
		return nil, _b.New("\u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006d\u006f\u0064\u0065\u006c")
	}
	_aeb.Width = _edcg.Width
	_aeb.Height = _edcg.Height
	_db.Log.Trace("\u0044\u0043T\u0020\u0045\u006ec\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076", _aeb)
	return _aeb, nil
}

// Add_g appends 'g' operand to the content stream:
// Same as G but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_fbed *ContentCreator) Add_g(gray float64) *ContentCreator {
	_afgd := ContentStreamOperation{}
	_afgd.Operand = "\u0067"
	_afgd.Params = _bdce([]float64{gray})
	_fbed._bg = append(_fbed._bg, &_afgd)
	return _fbed
}

// Add_ET appends 'ET' operand to the content stream:
// End text.
//
// See section 9.4 "Text Objects" and Table 107 (p. 256 PDF32000_2008).
func (_fdc *ContentCreator) Add_ET() *ContentCreator {
	_ddc := ContentStreamOperation{}
	_ddc.Operand = "\u0045\u0054"
	_fdc._bg = append(_fdc._bg, &_ddc)
	return _fdc
}

// Add_SCN_pattern appends 'SCN' operand to the content stream for pattern `name`:
// SCN with name attribute (for pattern). Syntax: c1 ... cn name SCN.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_fafc *ContentCreator) Add_SCN_pattern(name _dd.PdfObjectName, c ...float64) *ContentCreator {
	_cga := ContentStreamOperation{}
	_cga.Operand = "\u0053\u0043\u004e"
	_cga.Params = _bdce(c)
	_cga.Params = append(_cga.Params, _dd.MakeName(string(name)))
	_fafc._bg = append(_fafc._bg, &_cga)
	return _fafc
}

// Add_Tz appends 'Tz' operand to the content stream:
// Set horizontal scaling.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_bbgg *ContentCreator) Add_Tz(scale float64) *ContentCreator {
	_abf := ContentStreamOperation{}
	_abf.Operand = "\u0054\u007a"
	_abf.Params = _bdce([]float64{scale})
	_bbgg._bg = append(_bbgg._bg, &_abf)
	return _bbgg
}

// Add_BMC appends 'BMC' operand to the content stream:
// Begins a marked-content sequence terminated by a balancing EMC operator.
// `tag` shall be a name object indicating the role or significance of
// the sequence.
//
// See section 14.6 "Marked Content" and Table 320 (p. 561 PDF32000_2008).
func (_aabg *ContentCreator) Add_BMC(tag _dd.PdfObjectName) *ContentCreator {
	_dgf := ContentStreamOperation{}
	_dgf.Operand = "\u0042\u004d\u0043"
	_dgf.Params = _adbc([]_dd.PdfObjectName{tag})
	_aabg._bg = append(_aabg._bg, &_dgf)
	return _aabg
}
func _gab(_cedf []_dd.PdfObjectString) []_dd.PdfObject {
	var _fdfbd []_dd.PdfObject
	for _, _adce := range _cedf {
		_fdfbd = append(_fdfbd, _dd.MakeString(_adce.Str()))
	}
	return _fdfbd
}

// Add_W appends 'W' operand to the content stream:
// Modify the current clipping path by intersecting with the current path (nonzero winding rule).
//
// See section 8.5.4 "Clipping Path Operators" and Table 61 (p. 146 PDF32000_2008).
func (_edcd *ContentCreator) Add_W() *ContentCreator {
	_eeb := ContentStreamOperation{}
	_eeb.Operand = "\u0057"
	_edcd._bg = append(_edcd._bg, &_eeb)
	return _edcd
}

// Add_Tc appends 'Tc' operand to the content stream:
// Set character spacing.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_ffd *ContentCreator) Add_Tc(charSpace float64) *ContentCreator {
	_gfd := ContentStreamOperation{}
	_gfd.Operand = "\u0054\u0063"
	_gfd.Params = _bdce([]float64{charSpace})
	_ffd._bg = append(_ffd._bg, &_gfd)
	return _ffd
}

// ContentStreamOperation represents an operation in PDF contentstream which consists of
// an operand and parameters.
type ContentStreamOperation struct {
	Params  []_dd.PdfObject
	Operand string
}

// Add_n appends 'n' operand to the content stream:
// End the path without filling or stroking.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_adf *ContentCreator) Add_n() *ContentCreator {
	_fbe := ContentStreamOperation{}
	_fbe.Operand = "\u006e"
	_adf._bg = append(_adf._bg, &_fbe)
	return _adf
}

// Add_M adds 'M' operand to the content stream: Set the miter limit (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_fdg *ContentCreator) Add_M(miterlimit float64) *ContentCreator {
	_dde := ContentStreamOperation{}
	_dde.Operand = "\u004d"
	_dde.Params = _bdce([]float64{miterlimit})
	_fdg._bg = append(_fdg._bg, &_dde)
	return _fdg
}

// Add_W_starred appends 'W*' operand to the content stream:
// Modify the current clipping path by intersecting with the current path (even odd rule).
//
// See section 8.5.4 "Clipping Path Operators" and Table 61 (p. 146 PDF32000_2008).
func (_bea *ContentCreator) Add_W_starred() *ContentCreator {
	_eed := ContentStreamOperation{}
	_eed.Operand = "\u0057\u002a"
	_bea._bg = append(_bea._bg, &_eed)
	return _bea
}
func _cge(_cfd *ContentStreamInlineImage, _eecd *_dd.PdfObjectDictionary) (*_dd.LZWEncoder, error) {
	_ffa := _dd.NewLZWEncoder()
	if _eecd == nil {
		if _cfd.DecodeParms != nil {
			_cfe, _abe := _dd.GetDict(_cfd.DecodeParms)
			if !_abe {
				_db.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079 \u0028%\u0054\u0029", _cfd.DecodeParms)
				return nil, _ag.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
			_eecd = _cfe
		}
	}
	if _eecd == nil {
		return _ffa, nil
	}
	_dfge := _eecd.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
	if _dfge != nil {
		_fea, _dgc := _dfge.(*_dd.PdfObjectInteger)
		if !_dgc {
			_db.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a \u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u006e\u0075\u006d\u0065\u0072i\u0063 \u0028\u0025\u0054\u0029", _dfge)
			return nil, _ag.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
		}
		if *_fea != 0 && *_fea != 1 {
			return nil, _ag.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0076\u0061\u006c\u0075e\u0020\u0028\u006e\u006f\u0074 \u0030\u0020o\u0072\u0020\u0031\u0029")
		}
		_ffa.EarlyChange = int(*_fea)
	} else {
		_ffa.EarlyChange = 1
	}
	_dfge = _eecd.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _dfge != nil {
		_acee, _bade := _dfge.(*_dd.PdfObjectInteger)
		if !_bade {
			_db.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _dfge)
			return nil, _ag.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_ffa.Predictor = int(*_acee)
	}
	_dfge = _eecd.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _dfge != nil {
		_eab, _aag := _dfge.(*_dd.PdfObjectInteger)
		if !_aag {
			_db.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _ag.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_ffa.BitsPerComponent = int(*_eab)
	}
	if _ffa.Predictor > 1 {
		_ffa.Columns = 1
		_dfge = _eecd.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _dfge != nil {
			_bfeb, _ebcb := _dfge.(*_dd.PdfObjectInteger)
			if !_ebcb {
				return nil, _ag.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_ffa.Columns = int(*_bfeb)
		}
		_ffa.Colors = 1
		_dfge = _eecd.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _dfge != nil {
			_ggbc, _ccdd := _dfge.(*_dd.PdfObjectInteger)
			if !_ccdd {
				return nil, _ag.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_ffa.Colors = int(*_ggbc)
		}
	}
	_db.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _eecd.String())
	return _ffa, nil
}

// Bytes converts a set of content stream operations to a content stream byte presentation,
// i.e. the kind that can be stored as a PDF stream or string format.
func (_bac *ContentStreamOperations) Bytes() []byte {
	var _cg _ge.Buffer
	for _, _ga := range *_bac {
		if _ga == nil {
			continue
		}
		if _ga.Operand == "\u0042\u0049" {
			_cg.WriteString(_ga.Operand + "\u000a")
			_cg.WriteString(_ga.Params[0].WriteString())
		} else {
			for _, _ae := range _ga.Params {
				_cg.WriteString(_ae.WriteString())
				_cg.WriteString("\u0020")
			}
			_cg.WriteString(_ga.Operand + "\u000a")
		}
	}
	return _cg.Bytes()
}
func _cfc(_ace *ContentStreamInlineImage) (_dd.StreamEncoder, error) {
	if _ace.Filter == nil {
		return _dd.NewRawEncoder(), nil
	}
	_gd, _cbb := _ace.Filter.(*_dd.PdfObjectName)
	if !_cbb {
		_egc, _acf := _ace.Filter.(*_dd.PdfObjectArray)
		if !_acf {
			return nil, _ag.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0072 \u0041\u0072\u0072\u0061\u0079\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
		if _egc.Len() == 0 {
			return _dd.NewRawEncoder(), nil
		}
		if _egc.Len() != 1 {
			_beae, _dgb := _ggbb(_ace)
			if _dgb != nil {
				_db.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0063\u0072\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u006d\u0075\u006c\u0074i\u0020\u0065\u006e\u0063\u006f\u0064\u0065r\u003a\u0020\u0025\u0076", _dgb)
				return nil, _dgb
			}
			_db.Log.Trace("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063:\u0020\u0025\u0073\u000a", _beae)
			return _beae, nil
		}
		_ggb := _egc.Get(0)
		_gd, _acf = _ggb.(*_dd.PdfObjectName)
		if !_acf {
			return nil, _ag.Errorf("\u0066\u0069l\u0074\u0065\u0072\u0020a\u0072\u0072a\u0079\u0020\u006d\u0065\u006d\u0062\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
	}
	switch *_gd {
	case "\u0041\u0048\u0078", "\u0041\u0053\u0043\u0049\u0049\u0048\u0065\u0078\u0044e\u0063\u006f\u0064\u0065":
		return _dd.NewASCIIHexEncoder(), nil
	case "\u0041\u0038\u0035", "\u0041\u0053\u0043\u0049\u0049\u0038\u0035\u0044\u0065\u0063\u006f\u0064\u0065":
		return _dd.NewASCII85Encoder(), nil
	case "\u0044\u0043\u0054", "\u0044C\u0054\u0044\u0065\u0063\u006f\u0064e":
		return _dbd(_ace)
	case "\u0046\u006c", "F\u006c\u0061\u0074\u0065\u0044\u0065\u0063\u006f\u0064\u0065":
		return _ccda(_ace, nil)
	case "\u004c\u005a\u0057", "\u004cZ\u0057\u0044\u0065\u0063\u006f\u0064e":
		return _cge(_ace, nil)
	case "\u0043\u0043\u0046", "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065":
		return _dd.NewCCITTFaxEncoder(), nil
	case "\u0052\u004c", "\u0052u\u006eL\u0065\u006e\u0067\u0074\u0068\u0044\u0065\u0063\u006f\u0064\u0065":
		return _dd.NewRunLengthEncoder(), nil
	default:
		_db.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065 \u0069\u006d\u0061\u0067\u0065\u0020\u0065n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u003a\u0020\u0025\u0073", *_gd)
		return nil, _b.New("\u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006el\u0069n\u0065 \u0065n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006f\u0064")
	}
}
func (_gbae *ContentStreamProcessor) handleCommand_g(_dgbe *ContentStreamOperation, _ccab *_ba.PdfPageResources) error {
	_dgbc := _ba.NewPdfColorspaceDeviceGray()
	if len(_dgbe.Params) != _dgbc.GetNumComponents() {
		_db.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020p\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0067")
		_db.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_dgbe.Params), _dgbc)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_ebfa, _adg := _dgbc.ColorFromPdfObjects(_dgbe.Params)
	if _adg != nil {
		_db.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0068\u0061\u006e\u0064\u006c\u0065\u0043o\u006d\u006d\u0061\u006e\u0064\u005f\u0067\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061r\u0061\u006d\u0073\u002e\u0020c\u0073\u003d\u0025\u0054\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dgbc, _dgbe, _adg)
		return _adg
	}
	_gbae._gea.ColorspaceNonStroking = _dgbc
	_gbae._gea.ColorNonStroking = _ebfa
	return nil
}

// Operations returns the list of operations.
func (_bbg *ContentCreator) Operations() *ContentStreamOperations { return &_bbg._bg }

// Add_Tstar appends 'T*' operand to the content stream:
// Move to the start of next line.
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_bba *ContentCreator) Add_Tstar() *ContentCreator {
	_bgfb := ContentStreamOperation{}
	_bgfb.Operand = "\u0054\u002a"
	_bba._bg = append(_bba._bg, &_bgfb)
	return _bba
}

// WriteString outputs the object as it is to be written to file.
func (_bde *ContentStreamInlineImage) WriteString() string {
	var _gfg _ge.Buffer
	_fgaa := ""
	if _bde.BitsPerComponent != nil {
		_fgaa += "\u002f\u0042\u0050C\u0020" + _bde.BitsPerComponent.WriteString() + "\u000a"
	}
	if _bde.ColorSpace != nil {
		_fgaa += "\u002f\u0043\u0053\u0020" + _bde.ColorSpace.WriteString() + "\u000a"
	}
	if _bde.Decode != nil {
		_fgaa += "\u002f\u0044\u0020" + _bde.Decode.WriteString() + "\u000a"
	}
	if _bde.DecodeParms != nil {
		_fgaa += "\u002f\u0044\u0050\u0020" + _bde.DecodeParms.WriteString() + "\u000a"
	}
	if _bde.Filter != nil {
		_fgaa += "\u002f\u0046\u0020" + _bde.Filter.WriteString() + "\u000a"
	}
	if _bde.Height != nil {
		_fgaa += "\u002f\u0048\u0020" + _bde.Height.WriteString() + "\u000a"
	}
	if _bde.ImageMask != nil {
		_fgaa += "\u002f\u0049\u004d\u0020" + _bde.ImageMask.WriteString() + "\u000a"
	}
	if _bde.Intent != nil {
		_fgaa += "\u002f\u0049\u006e\u0074\u0065\u006e\u0074\u0020" + _bde.Intent.WriteString() + "\u000a"
	}
	if _bde.Interpolate != nil {
		_fgaa += "\u002f\u0049\u0020" + _bde.Interpolate.WriteString() + "\u000a"
	}
	if _bde.Width != nil {
		_fgaa += "\u002f\u0057\u0020" + _bde.Width.WriteString() + "\u000a"
	}
	_gfg.WriteString(_fgaa)
	_gfg.WriteString("\u0049\u0044\u0020")
	_gfg.Write(_bde._ega)
	_gfg.WriteString("\u000a\u0045\u0049\u000a")
	return _gfg.String()
}
func (_fecfc *ContentStreamProcessor) handleCommand_cm(_gac *ContentStreamOperation, _fdgd *_ba.PdfPageResources) error {
	if len(_gac.Params) != 6 {
		_db.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0063\u006d\u003a\u0020\u0025\u0064", len(_gac.Params))
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_aga, _dgcgd := _dd.GetNumbersAsFloat(_gac.Params)
	if _dgcgd != nil {
		return _dgcgd
	}
	_fbbc := _bc.NewMatrix(_aga[0], _aga[1], _aga[2], _aga[3], _aga[4], _aga[5])
	_fecfc._gea.CTM.Concat(_fbbc)
	return nil
}

// Add_cm adds 'cm' operation to the content stream: Modifies the current transformation matrix (ctm)
// of the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_cbc *ContentCreator) Add_cm(a, b, c, d, e, f float64) *ContentCreator {
	_bcab := ContentStreamOperation{}
	_bcab.Operand = "\u0063\u006d"
	_bcab.Params = _bdce([]float64{a, b, c, d, e, f})
	_cbc._bg = append(_cbc._bg, &_bcab)
	return _cbc
}

// Add_Do adds 'Do' operation to the content stream:
// Displays an XObject (image or form) specified by `name`.
//
// See section 8.8 "External Objects" and Table 87 (pp. 209-220 PDF32000_2008).
func (_dcd *ContentCreator) Add_Do(name _dd.PdfObjectName) *ContentCreator {
	_ead := ContentStreamOperation{}
	_ead.Operand = "\u0044\u006f"
	_ead.Params = _adbc([]_dd.PdfObjectName{name})
	_dcd._bg = append(_dcd._bg, &_ead)
	return _dcd
}

// Add_c adds 'c' operand to the content stream: Append a Bezier curve to the current path from
// the current point to (x3,y3) with (x1,x1) and (x2,y2) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_eag *ContentCreator) Add_c(x1, y1, x2, y2, x3, y3 float64) *ContentCreator {
	_bad := ContentStreamOperation{}
	_bad.Operand = "\u0063"
	_bad.Params = _bdce([]float64{x1, y1, x2, y2, x3, y3})
	_eag._bg = append(_eag._bg, &_bad)
	return _eag
}

// Add_Tf appends 'Tf' operand to the content stream:
// Set font and font size specified by font resource `fontName` and `fontSize`.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_gaf *ContentCreator) Add_Tf(fontName _dd.PdfObjectName, fontSize float64) *ContentCreator {
	_faba := ContentStreamOperation{}
	_faba.Operand = "\u0054\u0066"
	_faba.Params = _adbc([]_dd.PdfObjectName{fontName})
	_faba.Params = append(_faba.Params, _bdce([]float64{fontSize})...)
	_gaf._bg = append(_gaf._bg, &_faba)
	return _gaf
}
func _fgaaf(_eee _ba.PdfColorspace) bool {
	_, _gfde := _eee.(*_ba.PdfColorspaceSpecialPattern)
	return _gfde
}

// Add_SCN appends 'SCN' operand to the content stream:
// Same as SC but supports more colorspaces.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_aae *ContentCreator) Add_SCN(c ...float64) *ContentCreator {
	_ddf := ContentStreamOperation{}
	_ddf.Operand = "\u0053\u0043\u004e"
	_ddf.Params = _bdce(c)
	_aae._bg = append(_aae._bg, &_ddf)
	return _aae
}

// Add_cs appends 'cs' operand to the content stream:
// Same as CS but for non-stroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_ecf *ContentCreator) Add_cs(name _dd.PdfObjectName) *ContentCreator {
	_adb := ContentStreamOperation{}
	_adb.Operand = "\u0063\u0073"
	_adb.Params = _adbc([]_dd.PdfObjectName{name})
	_ecf._bg = append(_ecf._bg, &_adb)
	return _ecf
}

// ContentStreamOperations is a slice of ContentStreamOperations.
type ContentStreamOperations []*ContentStreamOperation

// Add_m adds 'm' operand to the content stream: Move the current point to (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_bfd *ContentCreator) Add_m(x, y float64) *ContentCreator {
	_bga := ContentStreamOperation{}
	_bga.Operand = "\u006d"
	_bga.Params = _bdce([]float64{x, y})
	_bfd._bg = append(_bfd._bg, &_bga)
	return _bfd
}

// Add_gs adds 'gs' operand to the content stream: Set the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_dcf *ContentCreator) Add_gs(dictName _dd.PdfObjectName) *ContentCreator {
	_cac := ContentStreamOperation{}
	_cac.Operand = "\u0067\u0073"
	_cac.Params = _adbc([]_dd.PdfObjectName{dictName})
	_dcf._bg = append(_dcf._bg, &_cac)
	return _dcf
}

// Add_f appends 'f' operand to the content stream:
// Fill the path using the nonzero winding number rule to determine fill region.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_fcc *ContentCreator) Add_f() *ContentCreator {
	_gaee := ContentStreamOperation{}
	_gaee.Operand = "\u0066"
	_fcc._bg = append(_fcc._bg, &_gaee)
	return _fcc
}

// Add_CS appends 'CS' operand to the content stream:
// Set the current colorspace for stroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_eff *ContentCreator) Add_CS(name _dd.PdfObjectName) *ContentCreator {
	_dea := ContentStreamOperation{}
	_dea.Operand = "\u0043\u0053"
	_dea.Params = _adbc([]_dd.PdfObjectName{name})
	_eff._bg = append(_eff._bg, &_dea)
	return _eff
}

// Add_quotes appends `"` operand to the content stream:
// Move to next line and show a string, using `aw` and `ac` as word
// and character spacing respectively.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_ddcb *ContentCreator) Add_quotes(textstr _dd.PdfObjectString, aw, ac float64) *ContentCreator {
	_fca := ContentStreamOperation{}
	_fca.Operand = "\u0022"
	_fca.Params = _bdce([]float64{aw, ac})
	_fca.Params = append(_fca.Params, _gab([]_dd.PdfObjectString{textstr})...)
	_ddcb._bg = append(_ddcb._bg, &_fca)
	return _ddcb
}

// Add_scn appends 'scn' operand to the content stream:
// Same as SC but for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_aaa *ContentCreator) Add_scn(c ...float64) *ContentCreator {
	_gbc := ContentStreamOperation{}
	_gbc.Operand = "\u0073\u0063\u006e"
	_gbc.Params = _bdce(c)
	_aaa._bg = append(_aaa._bg, &_gbc)
	return _aaa
}

// AddHandler adds a new ContentStreamProcessor `handler` of type `condition` for `operand`.
func (_ecc *ContentStreamProcessor) AddHandler(condition HandlerConditionEnum, operand string, handler HandlerFunc) {
	_ecce := handlerEntry{}
	_ecce.Condition = condition
	_ecce.Operand = operand
	_ecce.Handler = handler
	_ecc._efbe = append(_ecc._efbe, _ecce)
}

// Add_Q adds 'Q' operand to the content stream: Pops the most recently stored state from the stack.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_edb *ContentCreator) Add_Q() *ContentCreator {
	_bbb := ContentStreamOperation{}
	_bbb.Operand = "\u0051"
	_edb._bg = append(_edb._bg, &_bbb)
	return _edb
}

const (
	HandlerConditionEnumOperand HandlerConditionEnum = iota
	HandlerConditionEnumAllOperands
)

// RotateDeg applies a rotation to the transformation matrix.
func (_bdc *ContentCreator) RotateDeg(angle float64) *ContentCreator {
	_agd := _ff.Cos(angle * _ff.Pi / 180.0)
	_dbcd := _ff.Sin(angle * _ff.Pi / 180.0)
	_bgc := -_ff.Sin(angle * _ff.Pi / 180.0)
	_fd := _ff.Cos(angle * _ff.Pi / 180.0)
	return _bdc.Add_cm(_agd, _dbcd, _bgc, _fd, 0, 0)
}

// Add_rg appends 'rg' operand to the content stream:
// Same as RG but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_fdf *ContentCreator) Add_rg(r, g, b float64) *ContentCreator {
	_gcd := ContentStreamOperation{}
	_gcd.Operand = "\u0072\u0067"
	_gcd.Params = _bdce([]float64{r, g, b})
	_fdf._bg = append(_fdf._bg, &_gcd)
	return _fdf
}

// Add_re appends 're' operand to the content stream:
// Append a rectangle to the current path as a complete subpath, with lower left corner (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_ccd *ContentCreator) Add_re(x, y, width, height float64) *ContentCreator {
	_ccc := ContentStreamOperation{}
	_ccc.Operand = "\u0072\u0065"
	_ccc.Params = _bdce([]float64{x, y, width, height})
	_ccd._bg = append(_ccd._bg, &_ccc)
	return _ccd
}
func (_bfb *ContentStreamParser) parseBool() (_dd.PdfObjectBool, error) {
	_cdb, _gccf := _bfb._efe.Peek(4)
	if _gccf != nil {
		return _dd.PdfObjectBool(false), _gccf
	}
	if (len(_cdb) >= 4) && (string(_cdb[:4]) == "\u0074\u0072\u0075\u0065") {
		_bfb._efe.Discard(4)
		return _dd.PdfObjectBool(true), nil
	}
	_cdb, _gccf = _bfb._efe.Peek(5)
	if _gccf != nil {
		return _dd.PdfObjectBool(false), _gccf
	}
	if (len(_cdb) >= 5) && (string(_cdb[:5]) == "\u0066\u0061\u006cs\u0065") {
		_bfb._efe.Discard(5)
		return _dd.PdfObjectBool(false), nil
	}
	return _dd.PdfObjectBool(false), _b.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

// Add_RG appends 'RG' operand to the content stream:
// Set the stroking colorspace to DeviceRGB and sets the r,g,b colors (0-1 each).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_bbc *ContentCreator) Add_RG(r, g, b float64) *ContentCreator {
	_bgf := ContentStreamOperation{}
	_bgf.Operand = "\u0052\u0047"
	_bgf.Params = _bdce([]float64{r, g, b})
	_bbc._bg = append(_bbc._bg, &_bgf)
	return _bbc
}

// Add_G appends 'G' operand to the content stream:
// Set the stroking colorspace to DeviceGray and sets the gray level (0-1).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_gcc *ContentCreator) Add_G(gray float64) *ContentCreator {
	_bgca := ContentStreamOperation{}
	_bgca.Operand = "\u0047"
	_bgca.Params = _bdce([]float64{gray})
	_gcc._bg = append(_gcc._bg, &_bgca)
	return _gcc
}

// String is same as Bytes() except returns as a string for convenience.
func (_add *ContentCreator) String() string { return string(_add._bg.Bytes()) }

// Add_b appends 'b' operand to the content stream:
// Close, fill and then stroke the path (nonzero winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_cgc *ContentCreator) Add_b() *ContentCreator {
	_dg := ContentStreamOperation{}
	_dg.Operand = "\u0062"
	_cgc._bg = append(_cgc._bg, &_dg)
	return _cgc
}
func (_egef *ContentStreamParser) skipComments() error {
	if _, _cgcag := _egef.skipSpaces(); _cgcag != nil {
		return _cgcag
	}
	_gec := true
	for {
		_dgff, _bbga := _egef._efe.Peek(1)
		if _bbga != nil {
			_db.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _bbga.Error())
			return _bbga
		}
		if _gec && _dgff[0] != '%' {
			return nil
		}
		_gec = false
		if (_dgff[0] != '\r') && (_dgff[0] != '\n') {
			_egef._efe.ReadByte()
		} else {
			break
		}
	}
	return _egef.skipComments()
}
func (_afgf *ContentStreamProcessor) handleCommand_K(_egf *ContentStreamOperation, _cdba *_ba.PdfPageResources) error {
	_egba := _ba.NewPdfColorspaceDeviceCMYK()
	if len(_egf.Params) != _egba.GetNumComponents() {
		_db.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_db.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_egf.Params), _egba)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_cfbc, _eac := _egba.ColorFromPdfObjects(_egf.Params)
	if _eac != nil {
		return _eac
	}
	_afgf._gea.ColorspaceStroking = _egba
	_afgf._gea.ColorStroking = _cfbc
	return nil
}

// IsMask checks if an image is a mask.
// The image mask entry in the image dictionary specifies that the image data shall be used as a stencil
// mask for painting in the current color. The mask data is 1bpc, grayscale.
func (_febfa *ContentStreamInlineImage) IsMask() (bool, error) {
	if _febfa.ImageMask != nil {
		_dca, _ccb := _febfa.ImageMask.(*_dd.PdfObjectBool)
		if !_ccb {
			_db.Log.Debug("\u0049m\u0061\u0067\u0065\u0020\u006d\u0061\u0073\u006b\u0020\u006e\u006ft\u0020\u0061\u0020\u0062\u006f\u006f\u006c\u0065\u0061\u006e")
			return false, _b.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065")
		}
		return bool(*_dca), nil
	}
	return false, nil
}

// Wrap ensures that the contentstream is wrapped within a balanced q ... Q expression.
func (_cf *ContentCreator) Wrap() { _cf._bg.WrapIfNeeded() }
