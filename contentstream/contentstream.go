package contentstream

import (
	_g "bufio"
	_e "bytes"
	_bb "encoding/hex"
	_b "errors"
	_fe "fmt"
	_gf "image/color"
	_bbg "image/jpeg"
	_ff "io"
	_bg "math"
	_f "regexp"
	_a "strconv"

	_fb "unitechio/gopdf/gopdf/common"
	_ed "unitechio/gopdf/gopdf/core"
	_ac "unitechio/gopdf/gopdf/internal/imageutil"
	_de "unitechio/gopdf/gopdf/internal/transform"
	_gb "unitechio/gopdf/gopdf/model"
)

// Add_B_starred appends 'B*' operand to the content stream:
// Fill and then stroke the path (even-odd rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_cfg *ContentCreator) Add_B_starred() *ContentCreator {
	_cdac := ContentStreamOperation{}
	_cdac.Operand = "\u0042\u002a"
	_cfg._cg = append(_cfg._cg, &_cdac)
	return _cfg
}

func (_eccd *ContentStreamParser) parseNumber() (_ed.PdfObject, error) {
	return _ed.ParseNumber(_eccd._cac)
}

func (_dba *ContentStreamInlineImage) String() string {
	_bgb := _fe.Sprintf("I\u006el\u0069\u006e\u0065\u0049\u006d\u0061\u0067\u0065(\u006c\u0065\u006e\u003d%d\u0029\u000a", len(_dba._bce))
	if _dba.BitsPerComponent != nil {
		_bgb += "\u002d\u0020\u0042\u0050\u0043\u0020" + _dba.BitsPerComponent.WriteString() + "\u000a"
	}
	if _dba.ColorSpace != nil {
		_bgb += "\u002d\u0020\u0043S\u0020" + _dba.ColorSpace.WriteString() + "\u000a"
	}
	if _dba.Decode != nil {
		_bgb += "\u002d\u0020\u0044\u0020" + _dba.Decode.WriteString() + "\u000a"
	}
	if _dba.DecodeParms != nil {
		_bgb += "\u002d\u0020\u0044P\u0020" + _dba.DecodeParms.WriteString() + "\u000a"
	}
	if _dba.Filter != nil {
		_bgb += "\u002d\u0020\u0046\u0020" + _dba.Filter.WriteString() + "\u000a"
	}
	if _dba.Height != nil {
		_bgb += "\u002d\u0020\u0048\u0020" + _dba.Height.WriteString() + "\u000a"
	}
	if _dba.ImageMask != nil {
		_bgb += "\u002d\u0020\u0049M\u0020" + _dba.ImageMask.WriteString() + "\u000a"
	}
	if _dba.Intent != nil {
		_bgb += "\u002d \u0049\u006e\u0074\u0065\u006e\u0074 " + _dba.Intent.WriteString() + "\u000a"
	}
	if _dba.Interpolate != nil {
		_bgb += "\u002d\u0020\u0049\u0020" + _dba.Interpolate.WriteString() + "\u000a"
	}
	if _dba.Width != nil {
		_bgb += "\u002d\u0020\u0057\u0020" + _dba.Width.WriteString() + "\u000a"
	}
	return _bgb
}

func _eff(_dced _gb.PdfColorspace) bool {
	_, _gfb := _dced.(*_gb.PdfColorspaceSpecialPattern)
	return _gfb
}

func (_cedgb *ContentStreamProcessor) handleCommand_cs(_bcdf *ContentStreamOperation, _fdfee *_gb.PdfPageResources) error {
	if len(_bcdf.Params) < 1 {
		_fb.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0043\u0053\u0020\u0063\u006f\u006d\u006d\u0061\u006e\u0064\u002c\u0020s\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u006f\u0076\u0065\u0072")
		return _b.New("\u0074o\u006f \u0066\u0065\u0077\u0020\u0070a\u0072\u0061m\u0065\u0074\u0065\u0072\u0073")
	}
	if len(_bcdf.Params) > 1 {
		_fb.Log.Debug("\u0043\u0053\u0020\u0063\u006f\u006d\u006d\u0061n\u0064\u0020\u0077it\u0068\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020-\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075i\u006e\u0067")
		return _b.New("\u0074\u006f\u006f\u0020ma\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	}
	_cagc, _efbf := _bcdf.Params[0].(*_ed.PdfObjectName)
	if !_efbf {
		_fb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020C\u0053\u0020\u0063o\u006d\u006d\u0061n\u0064\u0020w\u0069\u0074\u0068\u0020\u0069\u006ev\u0061li\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u002c\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072")
		return _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_edeg, _dfga := _cedgb.getColorspace(string(*_cagc), _fdfee)
	if _dfga != nil {
		return _dfga
	}
	_cedgb._eced.ColorspaceNonStroking = _edeg
	_fcfa, _dfga := _cedgb.getInitialColor(_edeg)
	if _dfga != nil {
		return _dfga
	}
	_cedgb._eced.ColorNonStroking = _fcfa
	return nil
}

// Add_gs adds 'gs' operand to the content stream: Set the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_dge *ContentCreator) Add_gs(dictName _ed.PdfObjectName) *ContentCreator {
	_cf := ContentStreamOperation{}
	_cf.Operand = "\u0067\u0073"
	_cf.Params = _gfae([]_ed.PdfObjectName{dictName})
	_dge._cg = append(_dge._cg, &_cf)
	return _dge
}

func (_badc *ContentStreamProcessor) handleCommand_cm(_gbcc *ContentStreamOperation, _cdda *_gb.PdfPageResources) error {
	if len(_gbcc.Params) != 6 {
		_fb.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0063\u006d\u003a\u0020\u0025\u0064", len(_gbcc.Params))
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_gegb, _cfec := _ed.GetNumbersAsFloat(_gbcc.Params)
	if _cfec != nil {
		return _cfec
	}
	_begg := _de.NewMatrix(_gegb[0], _gegb[1], _gegb[2], _gegb[3], _gegb[4], _gegb[5])
	_badc._eced.CTM.Concat(_begg)
	return nil
}

// String is same as Bytes() except returns as a string for convenience.
func (_ef *ContentCreator) String() string { return string(_ef._cg.Bytes()) }

func (_dcff *ContentStreamProcessor) handleCommand_g(_bbab *ContentStreamOperation, _fbg *_gb.PdfPageResources) error {
	_bcaca := _gb.NewPdfColorspaceDeviceGray()
	if len(_bbab.Params) != _bcaca.GetNumComponents() {
		_fb.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020p\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0067")
		_fb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_bbab.Params), _bcaca)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_ggaa, _gfeb := _bcaca.ColorFromPdfObjects(_bbab.Params)
	if _gfeb != nil {
		_fb.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0068\u0061\u006e\u0064\u006c\u0065\u0043o\u006d\u006d\u0061\u006e\u0064\u005f\u0067\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061r\u0061\u006d\u0073\u002e\u0020c\u0073\u003d\u0025\u0054\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bcaca, _bbab, _gfeb)
		return _gfeb
	}
	_dcff._eced.ColorspaceNonStroking = _bcaca
	_dcff._eced.ColorNonStroking = _ggaa
	return nil
}

// Add_m adds 'm' operand to the content stream: Move the current point to (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_gae *ContentCreator) Add_m(x, y float64) *ContentCreator {
	_dc := ContentStreamOperation{}
	_dc.Operand = "\u006d"
	_dc.Params = _gfbfb([]float64{x, y})
	_gae._cg = append(_gae._cg, &_dc)
	return _gae
}

// Add_RG appends 'RG' operand to the content stream:
// Set the stroking colorspace to DeviceRGB and sets the r,g,b colors (0-1 each).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_fab *ContentCreator) Add_RG(r, g, b float64) *ContentCreator {
	_cga := ContentStreamOperation{}
	_cga.Operand = "\u0052\u0047"
	_cga.Params = _gfbfb([]float64{r, g, b})
	_fab._cg = append(_fab._cg, &_cga)
	return _fab
}

// Add_TL appends 'TL' operand to the content stream:
// Set leading.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_efc *ContentCreator) Add_TL(leading float64) *ContentCreator {
	_acf := ContentStreamOperation{}
	_acf.Operand = "\u0054\u004c"
	_acf.Params = _gfbfb([]float64{leading})
	_efc._cg = append(_efc._cg, &_acf)
	return _efc
}

// Add_W_starred appends 'W*' operand to the content stream:
// Modify the current clipping path by intersecting with the current path (even odd rule).
//
// See section 8.5.4 "Clipping Path Operators" and Table 61 (p. 146 PDF32000_2008).
func (_fdc *ContentCreator) Add_W_starred() *ContentCreator {
	_gac := ContentStreamOperation{}
	_gac.Operand = "\u0057\u002a"
	_fdc._cg = append(_fdc._cg, &_gac)
	return _fdc
}

// Add_j adds 'j' operand to the content stream: Set the line join style (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_daa *ContentCreator) Add_j(lineJoinStyle string) *ContentCreator {
	_ba := ContentStreamOperation{}
	_ba.Operand = "\u006a"
	_ba.Params = _gfae([]_ed.PdfObjectName{_ed.PdfObjectName(lineJoinStyle)})
	_daa._cg = append(_daa._cg, &_ba)
	return _daa
}

// SetNonStrokingColor sets the non-stroking `color` where color can be one of
// PdfColorDeviceGray, PdfColorDeviceRGB, or PdfColorDeviceCMYK.
func (_bbfg *ContentCreator) SetNonStrokingColor(color _gb.PdfColor) *ContentCreator {
	switch _bdg := color.(type) {
	case *_gb.PdfColorDeviceGray:
		_bbfg.Add_g(_bdg.Val())
	case *_gb.PdfColorDeviceRGB:
		_bbfg.Add_rg(_bdg.R(), _bdg.G(), _bdg.B())
	case *_gb.PdfColorDeviceCMYK:
		_bbfg.Add_k(_bdg.C(), _bdg.M(), _bdg.Y(), _bdg.K())
	case *_gb.PdfColorPatternType2:
		_bbfg.Add_cs(*_ed.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
		_bbfg.Add_scn_pattern(_bdg.PatternName)
	case *_gb.PdfColorPatternType3:
		_bbfg.Add_cs(*_ed.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
		_bbfg.Add_scn_pattern(_bdg.PatternName)
	default:
		_fb.Log.Debug("\u0053\u0065\u0074N\u006f\u006e\u0053\u0074\u0072\u006f\u006b\u0069\u006e\u0067\u0043\u006f\u006c\u006f\u0072\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020c\u006f\u006c\u006f\u0072\u003a\u0020\u0025\u0054", _bdg)
	}
	return _bbfg
}

// String returns `ops.Bytes()` as a string.
func (_ccg *ContentStreamOperations) String() string { return string(_ccg.Bytes()) }

func (_ecbe *ContentStreamParser) parseString() (*_ed.PdfObjectString, error) {
	_ecbe._cac.ReadByte()
	var _egf []byte
	_gdcd := 1
	for {
		_babf, _cfc := _ecbe._cac.Peek(1)
		if _cfc != nil {
			return _ed.MakeString(string(_egf)), _cfc
		}
		if _babf[0] == '\\' {
			_ecbe._cac.ReadByte()
			_dbe, _fagbd := _ecbe._cac.ReadByte()
			if _fagbd != nil {
				return _ed.MakeString(string(_egf)), _fagbd
			}
			if _ed.IsOctalDigit(_dbe) {
				_ffd, _agba := _ecbe._cac.Peek(2)
				if _agba != nil {
					return _ed.MakeString(string(_egf)), _agba
				}
				var _bgbc []byte
				_bgbc = append(_bgbc, _dbe)
				for _, _gefe := range _ffd {
					if _ed.IsOctalDigit(_gefe) {
						_bgbc = append(_bgbc, _gefe)
					} else {
						break
					}
				}
				_ecbe._cac.Discard(len(_bgbc) - 1)
				_fb.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _bgbc)
				_gabd, _agba := _a.ParseUint(string(_bgbc), 8, 32)
				if _agba != nil {
					return _ed.MakeString(string(_egf)), _agba
				}
				_egf = append(_egf, byte(_gabd))
				continue
			}
			switch _dbe {
			case 'n':
				_egf = append(_egf, '\n')
			case 'r':
				_egf = append(_egf, '\r')
			case 't':
				_egf = append(_egf, '\t')
			case 'b':
				_egf = append(_egf, '\b')
			case 'f':
				_egf = append(_egf, '\f')
			case '(':
				_egf = append(_egf, '(')
			case ')':
				_egf = append(_egf, ')')
			case '\\':
				_egf = append(_egf, '\\')
			}
			continue
		} else if _babf[0] == '(' {
			_gdcd++
		} else if _babf[0] == ')' {
			_gdcd--
			if _gdcd == 0 {
				_ecbe._cac.ReadByte()
				break
			}
		}
		_bedac, _ := _ecbe._cac.ReadByte()
		_egf = append(_egf, _bedac)
	}
	return _ed.MakeString(string(_egf)), nil
}

// Add_SC appends 'SC' operand to the content stream:
// Set color for stroking operations.  Input: c1, ..., cn.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_ebc *ContentCreator) Add_SC(c ...float64) *ContentCreator {
	_fbd := ContentStreamOperation{}
	_fbd.Operand = "\u0053\u0043"
	_fbd.Params = _gfbfb(c)
	_ebc._cg = append(_ebc._cg, &_fbd)
	return _ebc
}

// Add_B appends 'B' operand to the content stream:
// Fill and then stroke the path (nonzero winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_bgc *ContentCreator) Add_B() *ContentCreator {
	_gbc := ContentStreamOperation{}
	_gbc.Operand = "\u0042"
	_bgc._cg = append(_bgc._cg, &_gbc)
	return _bgc
}

func (_gfee *ContentStreamInlineImage) toImageBase(_acc *_gb.PdfPageResources) (*_ac.ImageBase, error) {
	if _gfee._gfgb != nil {
		return _gfee._gfgb, nil
	}
	_gdb := _ac.ImageBase{}
	if _gfee.Height == nil {
		return nil, _b.New("\u0068e\u0069\u0067\u0068\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_bbc, _afff := _gfee.Height.(*_ed.PdfObjectInteger)
	if !_afff {
		return nil, _b.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0068e\u0069\u0067\u0068\u0074")
	}
	_gdb.Height = int(*_bbc)
	if _gfee.Width == nil {
		return nil, _b.New("\u0077\u0069\u0064th\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_gdg, _afff := _gfee.Width.(*_ed.PdfObjectInteger)
	if !_afff {
		return nil, _b.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0069\u0064\u0074\u0068")
	}
	_gdb.Width = int(*_gdg)
	_ebd, _adddd := _gfee.IsMask()
	if _adddd != nil {
		return nil, _adddd
	}
	if _ebd {
		_gdb.BitsPerComponent = 1
		_gdb.ColorComponents = 1
	} else {
		if _gfee.BitsPerComponent == nil {
			_fb.Log.Debug("\u0049\u006el\u0069\u006e\u0065\u0020\u0042\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u002d\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u0038")
			_gdb.BitsPerComponent = 8
		} else {
			_bdga, _gabg := _gfee.BitsPerComponent.(*_ed.PdfObjectInteger)
			if !_gabg {
				_fb.Log.Debug("E\u0072\u0072\u006f\u0072\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0062\u0069\u0074\u0073 p\u0065\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u0076al\u0075\u0065,\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _gfee.BitsPerComponent)
				return nil, _b.New("\u0042\u0050\u0043\u0020\u0054\u0079\u0070\u0065\u0020e\u0072\u0072\u006f\u0072")
			}
			_gdb.BitsPerComponent = int(*_bdga)
		}
		if _gfee.ColorSpace != nil {
			_gdad, _gffg := _gfee.GetColorSpace(_acc)
			if _gffg != nil {
				return nil, _gffg
			}
			_gdb.ColorComponents = _gdad.GetNumComponents()
		} else {
			_fb.Log.Debug("\u0049\u006el\u0069\u006e\u0065\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006e\u006f\u0074\u0020\u0073p\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u002d\u0020\u0061\u0073\u0073\u0075m\u0069\u006eg\u0020\u0031\u0020\u0063o\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			_gdb.ColorComponents = 1
		}
	}
	if _ecb, _dda := _ed.GetArray(_gfee.Decode); _dda {
		_gdb.Decode, _adddd = _ecb.ToFloat64Array()
		if _adddd != nil {
			return nil, _adddd
		}
	}
	_gfee._gfgb = &_gdb
	return _gfee._gfgb, nil
}

// Add_Do adds 'Do' operation to the content stream:
// Displays an XObject (image or form) specified by `name`.
//
// See section 8.8 "External Objects" and Table 87 (pp. 209-220 PDF32000_2008).
func (_ag *ContentCreator) Add_Do(name _ed.PdfObjectName) *ContentCreator {
	_ccgb := ContentStreamOperation{}
	_ccgb.Operand = "\u0044\u006f"
	_ccgb.Params = _gfae([]_ed.PdfObjectName{name})
	_ag._cg = append(_ag._cg, &_ccgb)
	return _ag
}

func (_cgbd *ContentStreamProcessor) handleCommand_k(_dgce *ContentStreamOperation, _acef *_gb.PdfPageResources) error {
	_faae := _gb.NewPdfColorspaceDeviceCMYK()
	if len(_dgce.Params) != _faae.GetNumComponents() {
		_fb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_fb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_dgce.Params), _faae)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_beg, _aegb := _faae.ColorFromPdfObjects(_dgce.Params)
	if _aegb != nil {
		return _aegb
	}
	_cgbd._eced.ColorspaceNonStroking = _faae
	_cgbd._eced.ColorNonStroking = _beg
	return nil
}

// Add_w adds 'w' operand to the content stream, which sets the line width.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_caa *ContentCreator) Add_w(lineWidth float64) *ContentCreator {
	_cde := ContentStreamOperation{}
	_cde.Operand = "\u0077"
	_cde.Params = _gfbfb([]float64{lineWidth})
	_caa._cg = append(_caa._cg, &_cde)
	return _caa
}

func (_gabga *ContentStreamParser) parseArray() (*_ed.PdfObjectArray, error) {
	_eaab := _ed.MakeArray()
	_gabga._cac.ReadByte()
	for {
		_gabga.skipSpaces()
		_caf, _aad := _gabga._cac.Peek(1)
		if _aad != nil {
			return _eaab, _aad
		}
		if _caf[0] == ']' {
			_gabga._cac.ReadByte()
			break
		}
		_cebg, _, _aad := _gabga.parseObject()
		if _aad != nil {
			return _eaab, _aad
		}
		_eaab.Append(_cebg)
	}
	return _eaab, nil
}

// ContentStreamOperations is a slice of ContentStreamOperations.
type ContentStreamOperations []*ContentStreamOperation

// GetEncoder returns the encoder of the inline image.
func (_acge *ContentStreamInlineImage) GetEncoder() (_ed.StreamEncoder, error) { return _egb(_acge) }

// Add_cs appends 'cs' operand to the content stream:
// Same as CS but for non-stroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_adcb *ContentCreator) Add_cs(name _ed.PdfObjectName) *ContentCreator {
	_fgg := ContentStreamOperation{}
	_fgg.Operand = "\u0063\u0073"
	_fgg.Params = _gfae([]_ed.PdfObjectName{name})
	_adcb._cg = append(_adcb._cg, &_fgg)
	return _adcb
}

// Translate applies a simple x-y translation to the transformation matrix.
func (_ab *ContentCreator) Translate(tx, ty float64) *ContentCreator {
	return _ab.Add_cm(1, 0, 0, 1, tx, ty)
}

func (_fcb *ContentStreamProcessor) handleCommand_RG(_ffb *ContentStreamOperation, _cbab *_gb.PdfPageResources) error {
	_abe := _gb.NewPdfColorspaceDeviceRGB()
	if len(_ffb.Params) != _abe.GetNumComponents() {
		_fb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020R\u0047")
		_fb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_ffb.Params), _abe)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_ccfb, _fgf := _abe.ColorFromPdfObjects(_ffb.Params)
	if _fgf != nil {
		return _fgf
	}
	_fcb._eced.ColorspaceStroking = _abe
	_fcb._eced.ColorStroking = _ccfb
	return nil
}

// Add_y appends 'y' operand to the content stream: Append a Bezier curve to the current path from the
// current point to (x3,y3) with (x1, y1) and (x3,y3) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_cdd *ContentCreator) Add_y(x1, y1, x3, y3 float64) *ContentCreator {
	_cbb := ContentStreamOperation{}
	_cbb.Operand = "\u0079"
	_cbb.Params = _gfbfb([]float64{x1, y1, x3, y3})
	_cdd._cg = append(_cdd._cg, &_cbb)
	return _cdd
}

func (_cdad *ContentStreamParser) parseOperand() (*_ed.PdfObjectString, error) {
	var _bgec []byte
	for {
		_dgfd, _bgbce := _cdad._cac.Peek(1)
		if _bgbce != nil {
			return _ed.MakeString(string(_bgec)), _bgbce
		}
		if _ed.IsDelimiter(_dgfd[0]) {
			break
		}
		if _ed.IsWhiteSpace(_dgfd[0]) {
			break
		}
		_bcee, _ := _cdad._cac.ReadByte()
		_bgec = append(_bgec, _bcee)
	}
	return _ed.MakeString(string(_bgec)), nil
}

// Add_h appends 'h' operand to the content stream:
// Close the current subpath by adding a line between the current position and the starting position.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_bc *ContentCreator) Add_h() *ContentCreator {
	_bfgb := ContentStreamOperation{}
	_bfgb.Operand = "\u0068"
	_bc._cg = append(_bc._cg, &_bfgb)
	return _bc
}

// Add_Tstar appends 'T*' operand to the content stream:
// Move to the start of next line.
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_dacc *ContentCreator) Add_Tstar() *ContentCreator {
	_bcd := ContentStreamOperation{}
	_bcd.Operand = "\u0054\u002a"
	_dacc._cg = append(_dacc._cg, &_bcd)
	return _dacc
}

// Add_i adds 'i' operand to the content stream: Set the flatness tolerance in the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_cab *ContentCreator) Add_i(flatness float64) *ContentCreator {
	_bfg := ContentStreamOperation{}
	_bfg.Operand = "\u0069"
	_bfg.Params = _gfbfb([]float64{flatness})
	_cab._cg = append(_cab._cg, &_bfg)
	return _cab
}

// Add_Tc appends 'Tc' operand to the content stream:
// Set character spacing.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_acgd *ContentCreator) Add_Tc(charSpace float64) *ContentCreator {
	_ccb := ContentStreamOperation{}
	_ccb.Operand = "\u0054\u0063"
	_ccb.Params = _gfbfb([]float64{charSpace})
	_acgd._cg = append(_acgd._cg, &_ccb)
	return _acgd
}

// Add_quotes appends `"` operand to the content stream:
// Move to next line and show a string, using `aw` and `ac` as word
// and character spacing respectively.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_cba *ContentCreator) Add_quotes(textstr _ed.PdfObjectString, aw, ac float64) *ContentCreator {
	_gaa := ContentStreamOperation{}
	_gaa.Operand = "\u0022"
	_gaa.Params = _gfbfb([]float64{aw, ac})
	_gaa.Params = append(_gaa.Params, _bfe([]_ed.PdfObjectString{textstr})...)
	_cba._cg = append(_cba._cg, &_gaa)
	return _cba
}

// Add_cm adds 'cm' operation to the content stream: Modifies the current transformation matrix (ctm)
// of the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_dg *ContentCreator) Add_cm(a, b, c, d, e, f float64) *ContentCreator {
	_dgg := ContentStreamOperation{}
	_dgg.Operand = "\u0063\u006d"
	_dgg.Params = _gfbfb([]float64{a, b, c, d, e, f})
	_dg._cg = append(_dg._cg, &_dgg)
	return _dg
}

func (_dgca *ContentStreamParser) skipComments() error {
	if _, _gada := _dgca.skipSpaces(); _gada != nil {
		return _gada
	}
	_ece := true
	for {
		_efbc, _gef := _dgca._cac.Peek(1)
		if _gef != nil {
			_fb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _gef.Error())
			return _gef
		}
		if _ece && _efbc[0] != '%' {
			return nil
		}
		_ece = false
		if (_efbc[0] != '\r') && (_efbc[0] != '\n') {
			_dgca._cac.ReadByte()
		} else {
			break
		}
	}
	return _dgca.skipComments()
}

// Push pushes `gs` on the `gsStack`.
func (_bdcee *GraphicStateStack) Push(gs GraphicsState) { *_bdcee = append(*_bdcee, gs) }

// GraphicStateStack represents a stack of GraphicsState.
type GraphicStateStack []GraphicsState

// Add_BMC appends 'BMC' operand to the content stream:
// Begins a marked-content sequence terminated by a balancing EMC operator.
// `tag` shall be a name object indicating the role or significance of
// the sequence.
//
// See section 14.6 "Marked Content" and Table 320 (p. 561 PDF32000_2008).
func (_defd *ContentCreator) Add_BMC(tag _ed.PdfObjectName) *ContentCreator {
	_cea := ContentStreamOperation{}
	_cea.Operand = "\u0042\u004d\u0043"
	_cea.Params = _gfae([]_ed.PdfObjectName{tag})
	_defd._cg = append(_defd._cg, &_cea)
	return _defd
}

// Add_J adds 'J' operand to the content stream: Set the line cap style (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_ecf *ContentCreator) Add_J(lineCapStyle string) *ContentCreator {
	_cca := ContentStreamOperation{}
	_cca.Operand = "\u004a"
	_cca.Params = _gfae([]_ed.PdfObjectName{_ed.PdfObjectName(lineCapStyle)})
	_ecf._cg = append(_ecf._cg, &_cca)
	return _ecf
}

func (_cbbe *ContentStreamProcessor) handleCommand_rg(_daac *ContentStreamOperation, _fgc *_gb.PdfPageResources) error {
	_gfbf := _gb.NewPdfColorspaceDeviceRGB()
	if len(_daac.Params) != _gfbf.GetNumComponents() {
		_fb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_fb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_daac.Params), _gfbf)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_fcgf, _fdcea := _gfbf.ColorFromPdfObjects(_daac.Params)
	if _fdcea != nil {
		return _fdcea
	}
	_cbbe._eced.ColorspaceNonStroking = _gfbf
	_cbbe._eced.ColorNonStroking = _fcgf
	return nil
}

// Add_M adds 'M' operand to the content stream: Set the miter limit (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_fc *ContentCreator) Add_M(miterlimit float64) *ContentCreator {
	_dbd := ContentStreamOperation{}
	_dbd.Operand = "\u004d"
	_dbd.Params = _gfbfb([]float64{miterlimit})
	_fc._cg = append(_fc._cg, &_dbd)
	return _fc
}

func (_fbgf *ContentStreamProcessor) handleCommand_K(_aee *ContentStreamOperation, _acag *_gb.PdfPageResources) error {
	_bedg := _gb.NewPdfColorspaceDeviceCMYK()
	if len(_aee.Params) != _bedg.GetNumComponents() {
		_fb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_fb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_aee.Params), _bedg)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_gdaf, _dffd := _bedg.ColorFromPdfObjects(_aee.Params)
	if _dffd != nil {
		return _dffd
	}
	_fbgf._eced.ColorspaceStroking = _bedg
	_fbgf._eced.ColorStroking = _gdaf
	return nil
}

// GraphicsState is a basic graphics state implementation for PDF processing.
// Initially only implementing and tracking a portion of the information specified. Easy to add more.
type GraphicsState struct {
	ColorspaceStroking    _gb.PdfColorspace
	ColorspaceNonStroking _gb.PdfColorspace
	ColorStroking         _gb.PdfColor
	ColorNonStroking      _gb.PdfColor
	CTM                   _de.Matrix
}

// WriteString outputs the object as it is to be written to file.
func (_ceg *ContentStreamInlineImage) WriteString() string {
	var _aea _e.Buffer
	_ebf := ""
	if _ceg.BitsPerComponent != nil {
		_ebf += "\u002f\u0042\u0050C\u0020" + _ceg.BitsPerComponent.WriteString() + "\u000a"
	}
	if _ceg.ColorSpace != nil {
		_ebf += "\u002f\u0043\u0053\u0020" + _ceg.ColorSpace.WriteString() + "\u000a"
	}
	if _ceg.Decode != nil {
		_ebf += "\u002f\u0044\u0020" + _ceg.Decode.WriteString() + "\u000a"
	}
	if _ceg.DecodeParms != nil {
		_ebf += "\u002f\u0044\u0050\u0020" + _ceg.DecodeParms.WriteString() + "\u000a"
	}
	if _ceg.Filter != nil {
		_ebf += "\u002f\u0046\u0020" + _ceg.Filter.WriteString() + "\u000a"
	}
	if _ceg.Height != nil {
		_ebf += "\u002f\u0048\u0020" + _ceg.Height.WriteString() + "\u000a"
	}
	if _ceg.ImageMask != nil {
		_ebf += "\u002f\u0049\u004d\u0020" + _ceg.ImageMask.WriteString() + "\u000a"
	}
	if _ceg.Intent != nil {
		_ebf += "\u002f\u0049\u006e\u0074\u0065\u006e\u0074\u0020" + _ceg.Intent.WriteString() + "\u000a"
	}
	if _ceg.Interpolate != nil {
		_ebf += "\u002f\u0049\u0020" + _ceg.Interpolate.WriteString() + "\u000a"
	}
	if _ceg.Width != nil {
		_ebf += "\u002f\u0057\u0020" + _ceg.Width.WriteString() + "\u000a"
	}
	_aea.WriteString(_ebf)
	_aea.WriteString("\u0049\u0044\u0020")
	_aea.Write(_ceg._bce)
	_aea.WriteString("\u000a\u0045\u0049\u000a")
	return _aea.String()
}

// Add_Tj appends 'Tj' operand to the content stream:
// Show a text string.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_cebf *ContentCreator) Add_Tj(textstr _ed.PdfObjectString) *ContentCreator {
	_fdaf := ContentStreamOperation{}
	_fdaf.Operand = "\u0054\u006a"
	_fdaf.Params = _bfe([]_ed.PdfObjectString{textstr})
	_cebf._cg = append(_cebf._cg, &_fdaf)
	return _cebf
}

// Add_TD appends 'TD' operand to the content stream:
// Move to start of next line with offset (`tx`, `ty`).
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_gbe *ContentCreator) Add_TD(tx, ty float64) *ContentCreator {
	_dacf := ContentStreamOperation{}
	_dacf.Operand = "\u0054\u0044"
	_dacf.Params = _gfbfb([]float64{tx, ty})
	_gbe._cg = append(_gbe._cg, &_dacf)
	return _gbe
}

// NewContentCreator returns a new initialized ContentCreator.
func NewContentCreator() *ContentCreator {
	_dac := &ContentCreator{}
	_dac._cg = ContentStreamOperations{}
	return _dac
}

// Add_scn appends 'scn' operand to the content stream:
// Same as SC but for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_fga *ContentCreator) Add_scn(c ...float64) *ContentCreator {
	_caba := ContentStreamOperation{}
	_caba.Operand = "\u0073\u0063\u006e"
	_caba.Params = _gfbfb(c)
	_fga._cg = append(_fga._cg, &_caba)
	return _fga
}

func (_ead *ContentStreamParser) parseNull() (_ed.PdfObjectNull, error) {
	_, _ebad := _ead._cac.Discard(4)
	return _ed.PdfObjectNull{}, _ebad
}

// Add_q adds 'q' operand to the content stream: Pushes the current graphics state on the stack.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_eg *ContentCreator) Add_q() *ContentCreator {
	_gee := ContentStreamOperation{}
	_gee.Operand = "\u0071"
	_eg._cg = append(_eg._cg, &_gee)
	return _eg
}

// Operand returns true if `hce` is equivalent to HandlerConditionEnumOperand.
func (_bbca HandlerConditionEnum) Operand() bool { return _bbca == HandlerConditionEnumOperand }

func _gfae(_fggd []_ed.PdfObjectName) []_ed.PdfObject {
	var _eea []_ed.PdfObject
	for _, _fdff := range _fggd {
		_eea = append(_eea, _ed.MakeName(string(_fdff)))
	}
	return _eea
}

// Add_K appends 'K' operand to the content stream:
// Set the stroking colorspace to DeviceCMYK and sets the c,m,y,k color (0-1 each component).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_fag *ContentCreator) Add_K(c, m, y, k float64) *ContentCreator {
	_bbf := ContentStreamOperation{}
	_bbf.Operand = "\u004b"
	_bbf.Params = _gfbfb([]float64{c, m, y, k})
	_fag._cg = append(_fag._cg, &_bbf)
	return _fag
}

// SetStrokingColor sets the stroking `color` where color can be one of
// PdfColorDeviceGray, PdfColorDeviceRGB, or PdfColorDeviceCMYK.
func (_baf *ContentCreator) SetStrokingColor(color _gb.PdfColor) *ContentCreator {
	switch _faf := color.(type) {
	case *_gb.PdfColorDeviceGray:
		_baf.Add_G(_faf.Val())
	case *_gb.PdfColorDeviceRGB:
		_baf.Add_RG(_faf.R(), _faf.G(), _faf.B())
	case *_gb.PdfColorDeviceCMYK:
		_baf.Add_K(_faf.C(), _faf.M(), _faf.Y(), _faf.K())
	case *_gb.PdfColorPatternType2:
		_baf.Add_CS(*_ed.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
		_baf.Add_SCN_pattern(_faf.PatternName)
	case *_gb.PdfColorPatternType3:
		_baf.Add_CS(*_ed.MakeName("\u0050a\u0074\u0074\u0065\u0072\u006e"))
		_baf.Add_SCN_pattern(_faf.PatternName)
	default:
		_fb.Log.Debug("\u0053\u0065\u0074\u0053\u0074\u0072\u006f\u006b\u0069\u006e\u0067\u0043\u006fl\u006f\u0072\u003a\u0020\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006fr\u003a\u0020\u0025\u0054", _faf)
	}
	return _baf
}

const (
	HandlerConditionEnumOperand HandlerConditionEnum = iota
	HandlerConditionEnumAllOperands
)

func (_gba *ContentStreamOperations) isWrapped() bool {
	if len(*_gba) < 2 {
		return false
	}
	_fd := 0
	for _, _ce := range *_gba {
		if _ce.Operand == "\u0071" {
			_fd++
		} else if _ce.Operand == "\u0051" {
			_fd--
		} else {
			if _fd < 1 {
				return false
			}
		}
	}
	return _fd == 0
}

// ContentStreamProcessor defines a data structure and methods for processing a content stream, keeping track of the
// current graphics state, and allowing external handlers to define their own functions as a part of the processing,
// for example rendering or extracting certain information.
type ContentStreamProcessor struct {
	_dcf  GraphicStateStack
	_gcc  []*ContentStreamOperation
	_eced GraphicsState
	_efd  []handlerEntry
	_bgff int
}

// Add_g appends 'g' operand to the content stream:
// Same as G but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_bffa *ContentCreator) Add_g(gray float64) *ContentCreator {
	_aba := ContentStreamOperation{}
	_aba.Operand = "\u0067"
	_aba.Params = _gfbfb([]float64{gray})
	_bffa._cg = append(_bffa._cg, &_aba)
	return _bffa
}

// Add_SCN appends 'SCN' operand to the content stream:
// Same as SC but supports more colorspaces.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_gab *ContentCreator) Add_SCN(c ...float64) *ContentCreator {
	_fgbc := ContentStreamOperation{}
	_fgbc.Operand = "\u0053\u0043\u004e"
	_fgbc.Params = _gfbfb(c)
	_gab._cg = append(_gab._cg, &_fgbc)
	return _gab
}

// HasUnclosedQ checks if all the `q` operator is properly closed by `Q` operator.
func (_eda *ContentStreamOperations) HasUnclosedQ() bool {
	_c := 0
	for _, _ea := range *_eda {
		if _ea.Operand == "\u0071" {
			_c++
		} else if _ea.Operand == "\u0051" {
			_c--
		}
	}
	return _c != 0
}

// Add_Tf appends 'Tf' operand to the content stream:
// Set font and font size specified by font resource `fontName` and `fontSize`.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_def *ContentCreator) Add_Tf(fontName _ed.PdfObjectName, fontSize float64) *ContentCreator {
	_bdf := ContentStreamOperation{}
	_bdf.Operand = "\u0054\u0066"
	_bdf.Params = _gfae([]_ed.PdfObjectName{fontName})
	_bdf.Params = append(_bdf.Params, _gfbfb([]float64{fontSize})...)
	_def._cg = append(_def._cg, &_bdf)
	return _def
}

// HandlerFunc is the function syntax that the ContentStreamProcessor handler must implement.
type HandlerFunc func(_eccb *ContentStreamOperation, _efg GraphicsState, _caff *_gb.PdfPageResources) error

func _gfbfb(_fbcb []float64) []_ed.PdfObject {
	var _adcc []_ed.PdfObject
	for _, _ebgcb := range _fbcb {
		_adcc = append(_adcc, _ed.MakeFloat(_ebgcb))
	}
	return _adcc
}

// Add_EMC appends 'EMC' operand to the content stream:
// Ends a marked-content sequence.
//
// See section 14.6 "Marked Content" and Table 320 (p. 561 PDF32000_2008).
func (_gbd *ContentCreator) Add_EMC() *ContentCreator {
	_be := ContentStreamOperation{}
	_be.Operand = "\u0045\u004d\u0043"
	_gbd._cg = append(_gbd._cg, &_be)
	return _gbd
}

// Operations returns the list of operations.
func (_ffg *ContentCreator) Operations() *ContentStreamOperations { return &_ffg._cg }

// Add_scn_pattern appends 'scn' operand to the content stream for pattern `name`:
// scn with name attribute (for pattern). Syntax: c1 ... cn name scn.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_agc *ContentCreator) Add_scn_pattern(name _ed.PdfObjectName, c ...float64) *ContentCreator {
	_afg := ContentStreamOperation{}
	_afg.Operand = "\u0073\u0063\u006e"
	_afg.Params = _gfbfb(c)
	_afg.Params = append(_afg.Params, _ed.MakeName(string(name)))
	_agc._cg = append(_agc._cg, &_afg)
	return _agc
}

// NewInlineImageFromImage makes a new content stream inline image object from an image.
func NewInlineImageFromImage(img _gb.Image, encoder _ed.StreamEncoder) (*ContentStreamInlineImage, error) {
	if encoder == nil {
		encoder = _ed.NewRawEncoder()
	}
	encoder.UpdateParams(img.GetParamsDict())
	_fbe := ContentStreamInlineImage{}
	if img.ColorComponents == 1 {
		_fbe.ColorSpace = _ed.MakeName("\u0047")
	} else if img.ColorComponents == 3 {
		_fbe.ColorSpace = _ed.MakeName("\u0052\u0047\u0042")
	} else if img.ColorComponents == 4 {
		_fbe.ColorSpace = _ed.MakeName("\u0043\u004d\u0059\u004b")
	} else {
		_fb.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006db\u0065\u0072\u0020o\u0066\u0020c\u006f\u006c\u006f\u0072\u0020\u0063o\u006dpo\u006e\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0072\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0064", img.ColorComponents)
		return nil, _b.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020c\u006fl\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073")
	}
	_fbe.BitsPerComponent = _ed.MakeInteger(img.BitsPerComponent)
	_fbe.Width = _ed.MakeInteger(img.Width)
	_fbe.Height = _ed.MakeInteger(img.Height)
	_fbc, _gdd := encoder.EncodeBytes(img.Data)
	if _gdd != nil {
		return nil, _gdd
	}
	_fbe._bce = _fbc
	_bbb := encoder.GetFilterName()
	if _bbb != _ed.StreamEncodingFilterNameRaw {
		_fbe.Filter = _ed.MakeName(_bbb)
	}
	return &_fbe, nil
}

// Add_d adds 'd' operand to the content stream: Set the line dash pattern.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_ada *ContentCreator) Add_d(dashArray []int64, dashPhase int64) *ContentCreator {
	_feb := ContentStreamOperation{}
	_feb.Operand = "\u0064"
	_feb.Params = []_ed.PdfObject{}
	_feb.Params = append(_feb.Params, _ed.MakeArrayFromIntegers64(dashArray))
	_feb.Params = append(_feb.Params, _ed.MakeInteger(dashPhase))
	_ada._cg = append(_ada._cg, &_feb)
	return _ada
}

// Add_BT appends 'BT' operand to the content stream:
// Begin text.
//
// See section 9.4 "Text Objects" and Table 107 (p. 256 PDF32000_2008).
func (_fcg *ContentCreator) Add_BT() *ContentCreator {
	_cgc := ContentStreamOperation{}
	_cgc.Operand = "\u0042\u0054"
	_fcg._cg = append(_fcg._cg, &_cgc)
	return _fcg
}

// Add_f_starred appends 'f*' operand to the content stream.
// f*: Fill the path using the even-odd rule to determine fill region.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_gaf *ContentCreator) Add_f_starred() *ContentCreator {
	_febc := ContentStreamOperation{}
	_febc.Operand = "\u0066\u002a"
	_gaf._cg = append(_gaf._cg, &_febc)
	return _gaf
}

var (
	ErrInvalidOperand = _b.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
	ErrEarlyExit      = _b.New("\u0074\u0065\u0072\u006di\u006e\u0061\u0074\u0065\u0020\u0070\u0072\u006f\u0063\u0065s\u0073 \u0065\u0061\u0072\u006c\u0079\u0020\u0065x\u0069\u0074")
)

// Add_v appends 'v' operand to the content stream: Append a Bezier curve to the current path from the
// current point to (x3,y3) with the current point and (x2,y2) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_egd *ContentCreator) Add_v(x2, y2, x3, y3 float64) *ContentCreator {
	_cbc := ContentStreamOperation{}
	_cbc.Operand = "\u0076"
	_cbc.Params = _gfbfb([]float64{x2, y2, x3, y3})
	_egd._cg = append(_egd._cg, &_cbc)
	return _egd
}

// Add_Tm appends 'Tm' operand to the content stream:
// Set the text line matrix.
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_gfe *ContentCreator) Add_Tm(a, b, c, d, e, f float64) *ContentCreator {
	_eebg := ContentStreamOperation{}
	_eebg.Operand = "\u0054\u006d"
	_eebg.Params = _gfbfb([]float64{a, b, c, d, e, f})
	_gfe._cg = append(_gfe._cg, &_eebg)
	return _gfe
}

// Parse parses all commands in content stream, returning a list of operation data.
func (_ebcc *ContentStreamParser) Parse() (*ContentStreamOperations, error) {
	_bedf := ContentStreamOperations{}
	for {
		_ccf := ContentStreamOperation{}
		for {
			_bcc, _ged, _beb := _ebcc.parseObject()
			if _beb != nil {
				if _beb == _ff.EOF {
					return &_bedf, nil
				}
				return &_bedf, _beb
			}
			if _ged {
				_ccf.Operand, _ = _ed.GetStringVal(_bcc)
				_bedf = append(_bedf, &_ccf)
				break
			} else {
				_ccf.Params = append(_ccf.Params, _bcc)
			}
		}
		if _ccf.Operand == "\u0042\u0049" {
			_afc, _adce := _ebcc.ParseInlineImage()
			if _adce != nil {
				return &_bedf, _adce
			}
			_ccf.Params = append(_ccf.Params, _afc)
		}
	}
}

// Add_W appends 'W' operand to the content stream:
// Modify the current clipping path by intersecting with the current path (nonzero winding rule).
//
// See section 8.5.4 "Clipping Path Operators" and Table 61 (p. 146 PDF32000_2008).
func (_gda *ContentCreator) Add_W() *ContentCreator {
	_ggb := ContentStreamOperation{}
	_ggb.Operand = "\u0057"
	_gda._cg = append(_gda._cg, &_ggb)
	return _gda
}

func (_edbc *ContentStreamProcessor) getColorspace(_acgde string, _ebgc *_gb.PdfPageResources) (_gb.PdfColorspace, error) {
	switch _acgde {
	case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
		return _gb.NewPdfColorspaceDeviceGray(), nil
	case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
		return _gb.NewPdfColorspaceDeviceRGB(), nil
	case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
		return _gb.NewPdfColorspaceDeviceCMYK(), nil
	case "\u0050a\u0074\u0074\u0065\u0072\u006e":
		return _gb.NewPdfColorspaceSpecialPattern(), nil
	}
	_cegf, _facg := _ebgc.GetColorspaceByName(_ed.PdfObjectName(_acgde))
	if _facg {
		return _cegf, nil
	}
	switch _acgde {
	case "\u0043a\u006c\u0047\u0072\u0061\u0079":
		return _gb.NewPdfColorspaceCalGray(), nil
	case "\u0043\u0061\u006c\u0052\u0047\u0042":
		return _gb.NewPdfColorspaceCalRGB(), nil
	case "\u004c\u0061\u0062":
		return _gb.NewPdfColorspaceLab(), nil
	}
	_fb.Log.Debug("\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063e\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074\u0065\u0064\u003a\u0020\u0025\u0073", _acgde)
	return nil, _fe.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065:\u0020\u0025\u0073", _acgde)
}

var _cdf = map[string]struct{}{"\u0062": {}, "\u0042": {}, "\u0062\u002a": {}, "\u0042\u002a": {}, "\u0042\u0044\u0043": {}, "\u0042\u0049": {}, "\u0042\u004d\u0043": {}, "\u0042\u0054": {}, "\u0042\u0058": {}, "\u0063": {}, "\u0063\u006d": {}, "\u0043\u0053": {}, "\u0063\u0073": {}, "\u0064": {}, "\u0064\u0030": {}, "\u0064\u0031": {}, "\u0044\u006f": {}, "\u0044\u0050": {}, "\u0045\u0049": {}, "\u0045\u004d\u0043": {}, "\u0045\u0054": {}, "\u0045\u0058": {}, "\u0066": {}, "\u0046": {}, "\u0066\u002a": {}, "\u0047": {}, "\u0067": {}, "\u0067\u0073": {}, "\u0068": {}, "\u0069": {}, "\u0049\u0044": {}, "\u006a": {}, "\u004a": {}, "\u004b": {}, "\u006b": {}, "\u006c": {}, "\u006d": {}, "\u004d": {}, "\u004d\u0050": {}, "\u006e": {}, "\u0071": {}, "\u0051": {}, "\u0072\u0065": {}, "\u0052\u0047": {}, "\u0072\u0067": {}, "\u0072\u0069": {}, "\u0073": {}, "\u0053": {}, "\u0053\u0043": {}, "\u0073\u0063": {}, "\u0053\u0043\u004e": {}, "\u0073\u0063\u006e": {}, "\u0073\u0068": {}, "\u0054\u002a": {}, "\u0054\u0063": {}, "\u0054\u0064": {}, "\u0054\u0044": {}, "\u0054\u0066": {}, "\u0054\u006a": {}, "\u0054\u004a": {}, "\u0054\u004c": {}, "\u0054\u006d": {}, "\u0054\u0072": {}, "\u0054\u0073": {}, "\u0054\u0077": {}, "\u0054\u007a": {}, "\u0076": {}, "\u0077": {}, "\u0057": {}, "\u0057\u002a": {}, "\u0079": {}, "\u0027": {}, "\u0022": {}}

// Add_rg appends 'rg' operand to the content stream:
// Same as RG but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_bfd *ContentCreator) Add_rg(r, g, b float64) *ContentCreator {
	_dee := ContentStreamOperation{}
	_dee.Operand = "\u0072\u0067"
	_dee.Params = _gfbfb([]float64{r, g, b})
	_bfd._cg = append(_bfd._cg, &_dee)
	return _bfd
}

// Add_l adds 'l' operand to the content stream:
// Append a straight line segment from the current point to (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_gdf *ContentCreator) Add_l(x, y float64) *ContentCreator {
	_cec := ContentStreamOperation{}
	_cec.Operand = "\u006c"
	_cec.Params = _gfbfb([]float64{x, y})
	_gdf._cg = append(_gdf._cg, &_cec)
	return _gdf
}

func _bffg(_fca *ContentStreamInlineImage, _acb *_ed.PdfObjectDictionary) (*_ed.FlateEncoder, error) {
	_dbg := _ed.NewFlateEncoder()
	if _fca._gfgb != nil {
		_dbg.SetImage(_fca._gfgb)
	}
	if _acb == nil {
		_fbb := _fca.DecodeParms
		if _fbb != nil {
			_dgd, _ded := _ed.GetDict(_fbb)
			if !_ded {
				_fb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079 \u0028%\u0054\u0029", _fbb)
				return nil, _fe.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
			_acb = _dgd
		}
	}
	if _acb == nil {
		return _dbg, nil
	}
	_fb.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _acb.String())
	_bffb := _acb.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _bffb == nil {
		_fb.Log.Debug("E\u0072\u0072o\u0072\u003a\u0020\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0066\u0072\u006f\u006d\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073 \u002d\u0020\u0043\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020\u0077\u0069t\u0068\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u00281\u0029")
	} else {
		_aga, _bgda := _bffb.(*_ed.PdfObjectInteger)
		if !_bgda {
			_fb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _bffb)
			return nil, _fe.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_dbg.Predictor = int(*_aga)
	}
	_bffb = _acb.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _bffb != nil {
		_acfg, _eeba := _bffb.(*_ed.PdfObjectInteger)
		if !_eeba {
			_fb.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _fe.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_dbg.BitsPerComponent = int(*_acfg)
	}
	if _dbg.Predictor > 1 {
		_dbg.Columns = 1
		_bffb = _acb.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _bffb != nil {
			_cbca, _ccc := _bffb.(*_ed.PdfObjectInteger)
			if !_ccc {
				return nil, _fe.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_dbg.Columns = int(*_cbca)
		}
		_dbg.Colors = 1
		_gga := _acb.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _gga != nil {
			_bef, _decc := _gga.(*_ed.PdfObjectInteger)
			if !_decc {
				return nil, _fe.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_dbg.Colors = int(*_bef)
		}
	}
	return _dbg, nil
}

// Transform returns coordinates x, y transformed by the CTM.
func (_deg *GraphicsState) Transform(x, y float64) (float64, float64) {
	return _deg.CTM.Transform(x, y)
}

// IsMask checks if an image is a mask.
// The image mask entry in the image dictionary specifies that the image data shall be used as a stencil
// mask for painting in the current color. The mask data is 1bpc, grayscale.
func (_bfc *ContentStreamInlineImage) IsMask() (bool, error) {
	if _bfc.ImageMask != nil {
		_edaf, _cge := _bfc.ImageMask.(*_ed.PdfObjectBool)
		if !_cge {
			_fb.Log.Debug("\u0049m\u0061\u0067\u0065\u0020\u006d\u0061\u0073\u006b\u0020\u006e\u006ft\u0020\u0061\u0020\u0062\u006f\u006f\u006c\u0065\u0061\u006e")
			return false, _b.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065")
		}
		return bool(*_edaf), nil
	}
	return false, nil
}

// Add_s appends 's' operand to the content stream: Close and stroke the path.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_eb *ContentCreator) Add_s() *ContentCreator {
	_agg := ContentStreamOperation{}
	_agg.Operand = "\u0073"
	_eb._cg = append(_eb._cg, &_agg)
	return _eb
}

func _gaea(_fff *ContentStreamInlineImage) (*_ed.MultiEncoder, error) {
	_adga := _ed.NewMultiEncoder()
	var _agb *_ed.PdfObjectDictionary
	var _baa []_ed.PdfObject
	if _ccd := _fff.DecodeParms; _ccd != nil {
		_gbaec, _bed := _ccd.(*_ed.PdfObjectDictionary)
		if _bed {
			_agb = _gbaec
		}
		_aaa, _addb := _ccd.(*_ed.PdfObjectArray)
		if _addb {
			for _, _cdea := range _aaa.Elements() {
				if _fbbf, _fagb := _cdea.(*_ed.PdfObjectDictionary); _fagb {
					_baa = append(_baa, _fbbf)
				} else {
					_baa = append(_baa, nil)
				}
			}
		}
	}
	_baef := _fff.Filter
	if _baef == nil {
		return nil, _fe.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	_dca, _faa := _baef.(*_ed.PdfObjectArray)
	if !_faa {
		return nil, _fe.Errorf("m\u0075\u006c\u0074\u0069\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0063\u0061\u006e\u0020\u006f\u006el\u0079\u0020\u0062\u0065\u0020\u006d\u0061\u0064\u0065\u0020fr\u006f\u006d\u0020a\u0072r\u0061\u0079")
	}
	for _edb, _ega := range _dca.Elements() {
		_fgbd, _ced := _ega.(*_ed.PdfObjectName)
		if !_ced {
			return nil, _fe.Errorf("\u006d\u0075l\u0074\u0069\u0020\u0066i\u006c\u0074e\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061 \u006e\u0061\u006d\u0065")
		}
		var _fde _ed.PdfObject
		if _agb != nil {
			_fde = _agb
		} else {
			if len(_baa) > 0 {
				if _edb >= len(_baa) {
					return nil, _fe.Errorf("\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0073\u0020\u0069\u006e\u0020d\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006d\u0073\u0020a\u0072\u0072\u0061\u0079")
				}
				_fde = _baa[_edb]
			}
		}
		var _egbf *_ed.PdfObjectDictionary
		if _bbfc, _adf := _fde.(*_ed.PdfObjectDictionary); _adf {
			_egbf = _bbfc
		}
		if *_fgbd == _ed.StreamEncodingFilterNameFlate || *_fgbd == "\u0046\u006c" {
			_ecdd, _geg := _bffg(_fff, _egbf)
			if _geg != nil {
				return nil, _geg
			}
			_adga.AddEncoder(_ecdd)
		} else if *_fgbd == _ed.StreamEncodingFilterNameLZW {
			_gag, _efa := _ecfd(_fff, _egbf)
			if _efa != nil {
				return nil, _efa
			}
			_adga.AddEncoder(_gag)
		} else if *_fgbd == _ed.StreamEncodingFilterNameASCIIHex {
			_cbf := _ed.NewASCIIHexEncoder()
			_adga.AddEncoder(_cbf)
		} else if *_fgbd == _ed.StreamEncodingFilterNameASCII85 || *_fgbd == "\u0041\u0038\u0035" {
			_fbag := _ed.NewASCII85Encoder()
			_adga.AddEncoder(_fbag)
		} else {
			_fb.Log.Error("U\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0066\u0069l\u0074\u0065\u0072\u0020\u0025\u0073", *_fgbd)
			return nil, _fe.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u0066\u0069\u006c\u0074er \u0069n \u006d\u0075\u006c\u0074\u0069\u0020\u0066il\u0074\u0065\u0072\u0020\u0061\u0072\u0072a\u0079")
		}
	}
	return _adga, nil
}

// Add_quote appends "'" operand to the content stream:
// Move to next line and show a string.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_cgb *ContentCreator) Add_quote(textstr _ed.PdfObjectString) *ContentCreator {
	_cgbg := ContentStreamOperation{}
	_cgbg.Operand = "\u0027"
	_cgbg.Params = _bfe([]_ed.PdfObjectString{textstr})
	_cgb._cg = append(_cgb._cg, &_cgbg)
	return _cgb
}

// ContentStreamInlineImage is a representation of an inline image in a Content stream. Everything between the BI and EI operands.
// ContentStreamInlineImage implements the core.PdfObject interface although strictly it is not a PDF object.
type ContentStreamInlineImage struct {
	BitsPerComponent _ed.PdfObject
	ColorSpace       _ed.PdfObject
	Decode           _ed.PdfObject
	DecodeParms      _ed.PdfObject
	Filter           _ed.PdfObject
	Height           _ed.PdfObject
	ImageMask        _ed.PdfObject
	Intent           _ed.PdfObject
	Interpolate      _ed.PdfObject
	Width            _ed.PdfObject
	_bce             []byte
	_gfgb            *_ac.ImageBase
}

func _bfe(_ebfd []_ed.PdfObjectString) []_ed.PdfObject {
	var _bgbe []_ed.PdfObject
	for _, _ffa := range _ebfd {
		_bgbe = append(_bgbe, _ed.MakeString(_ffa.Str()))
	}
	return _bgbe
}

func (_gdgc *ContentStreamParser) parseName() (_ed.PdfObjectName, error) {
	_eafe := ""
	_bad := false
	for {
		_eed, _deae := _gdgc._cac.Peek(1)
		if _deae == _ff.EOF {
			break
		}
		if _deae != nil {
			return _ed.PdfObjectName(_eafe), _deae
		}
		if !_bad {
			if _eed[0] == '/' {
				_bad = true
				_gdgc._cac.ReadByte()
			} else {
				_fb.Log.Error("N\u0061\u006d\u0065\u0020\u0073\u0074a\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069\u0074h\u0020\u0025\u0073 \u0028%\u0020\u0078\u0029", _eed, _eed)
				return _ed.PdfObjectName(_eafe), _fe.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _eed[0])
			}
		} else {
			if _ed.IsWhiteSpace(_eed[0]) {
				break
			} else if (_eed[0] == '/') || (_eed[0] == '[') || (_eed[0] == '(') || (_eed[0] == ']') || (_eed[0] == '<') || (_eed[0] == '>') {
				break
			} else if _eed[0] == '#' {
				_adfa, _dggd := _gdgc._cac.Peek(3)
				if _dggd != nil {
					return _ed.PdfObjectName(_eafe), _dggd
				}
				_gdgc._cac.Discard(3)
				_cdae, _dggd := _bb.DecodeString(string(_adfa[1:3]))
				if _dggd != nil {
					return _ed.PdfObjectName(_eafe), _dggd
				}
				_eafe += string(_cdae)
			} else {
				_bac, _ := _gdgc._cac.ReadByte()
				_eafe += string(_bac)
			}
		}
	}
	return _ed.PdfObjectName(_eafe), nil
}

// Bytes converts a set of content stream operations to a content stream byte presentation,
// i.e. the kind that can be stored as a PDF stream or string format.
func (_fda *ContentStreamOperations) Bytes() []byte {
	var _cb _e.Buffer
	for _, _ge := range *_fda {
		if _ge == nil {
			continue
		}
		if _ge.Operand == "\u0042\u0049" {
			_cb.WriteString(_ge.Operand + "\u000a")
			_cb.WriteString(_ge.Params[0].WriteString())
		} else {
			for _, _cbg := range _ge.Params {
				_cb.WriteString(_cbg.WriteString())
				_cb.WriteString("\u0020")
			}
			_cb.WriteString(_ge.Operand + "\u000a")
		}
	}
	return _cb.Bytes()
}

// NewContentStreamParser creates a new instance of the content stream parser from an input content
// stream string.
func NewContentStreamParser(contentStr string) *ContentStreamParser {
	_bdff := ContentStreamParser{}
	contentStr = string(_bdd.ReplaceAll([]byte(contentStr), []byte("\u002f")))
	_acfc := _e.NewBufferString(contentStr + "\u000a")
	_bdff._cac = _g.NewReader(_acfc)
	return &_bdff
}

func (_cgf *ContentStreamParser) parseDict() (*_ed.PdfObjectDictionary, error) {
	_fb.Log.Trace("\u0052\u0065\u0061\u0064i\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074 \u0073t\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0021")
	_efcd := _ed.MakeDict()
	_cegd, _ := _cgf._cac.ReadByte()
	if _cegd != '<' {
		return nil, _b.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	_cegd, _ = _cgf._cac.ReadByte()
	if _cegd != '<' {
		return nil, _b.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	for {
		_cgf.skipSpaces()
		_eab, _bgdf := _cgf._cac.Peek(2)
		if _bgdf != nil {
			return nil, _bgdf
		}
		_fb.Log.Trace("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_eab), string(_eab))
		if (_eab[0] == '>') && (_eab[1] == '>') {
			_fb.Log.Trace("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
			_cgf._cac.ReadByte()
			_cgf._cac.ReadByte()
			break
		}
		_fb.Log.Trace("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021")
		_fdafb, _bgdf := _cgf.parseName()
		_fb.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _fdafb)
		if _bgdf != nil {
			_fb.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073", _bgdf)
			return nil, _bgdf
		}
		if len(_fdafb) > 4 && _fdafb[len(_fdafb)-4:] == "\u006e\u0075\u006c\u006c" {
			_bced := _fdafb[0 : len(_fdafb)-4]
			_fb.Log.Trace("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029", _fdafb)
			_fb.Log.Trace("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c", _bced)
			_cgf.skipSpaces()
			_cbbg, _ := _cgf._cac.Peek(1)
			if _cbbg[0] == '/' {
				_efcd.Set(_bced, _ed.MakeNull())
				continue
			}
		}
		_cgf.skipSpaces()
		_acfe, _, _bgdf := _cgf.parseObject()
		if _bgdf != nil {
			return nil, _bgdf
		}
		_efcd.Set(_fdafb, _acfe)
		_fb.Log.Trace("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073", _fdafb, _acfe.String())
	}
	return _efcd, nil
}

// Add_Ts appends 'Ts' operand to the content stream:
// Set text rise.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_afgd *ContentCreator) Add_Ts(rise float64) *ContentCreator {
	_add := ContentStreamOperation{}
	_add.Operand = "\u0054\u0073"
	_add.Params = _gfbfb([]float64{rise})
	_afgd._cg = append(_afgd._cg, &_add)
	return _afgd
}

// Add_ET appends 'ET' operand to the content stream:
// End text.
//
// See section 9.4 "Text Objects" and Table 107 (p. 256 PDF32000_2008).
func (_cag *ContentCreator) Add_ET() *ContentCreator {
	_acg := ContentStreamOperation{}
	_acg.Operand = "\u0045\u0054"
	_cag._cg = append(_cag._cg, &_acg)
	return _cag
}

// RotateDeg applies a rotation to the transformation matrix.
func (_ec *ContentCreator) RotateDeg(angle float64) *ContentCreator {
	_cgda := _bg.Cos(angle * _bg.Pi / 180.0)
	_gfge := _bg.Sin(angle * _bg.Pi / 180.0)
	_aa := -_bg.Sin(angle * _bg.Pi / 180.0)
	_bgd := _bg.Cos(angle * _bg.Pi / 180.0)
	return _ec.Add_cm(_cgda, _gfge, _aa, _bgd, 0, 0)
}

// Wrap ensures that the contentstream is wrapped within a balanced q ... Q expression.
func (_adc *ContentCreator) Wrap() { _adc._cg.WrapIfNeeded() }

// ContentCreator is a builder for PDF content streams.
type ContentCreator struct{ _cg ContentStreamOperations }

// HandlerConditionEnum represents the type of operand content stream processor (handler).
// The handler may process a single specific named operand or all operands.
type HandlerConditionEnum int

// Add_SCN_pattern appends 'SCN' operand to the content stream for pattern `name`:
// SCN with name attribute (for pattern). Syntax: c1 ... cn name SCN.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_efe *ContentCreator) Add_SCN_pattern(name _ed.PdfObjectName, c ...float64) *ContentCreator {
	_ebb := ContentStreamOperation{}
	_ebb.Operand = "\u0053\u0043\u004e"
	_ebb.Params = _gfbfb(c)
	_ebb.Params = append(_ebb.Params, _ed.MakeName(string(name)))
	_efe._cg = append(_efe._cg, &_ebb)
	return _efe
}

// Add_G appends 'G' operand to the content stream:
// Set the stroking colorspace to DeviceGray and sets the gray level (0-1).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_bgg *ContentCreator) Add_G(gray float64) *ContentCreator {
	_bba := ContentStreamOperation{}
	_bba.Operand = "\u0047"
	_bba.Params = _gfbfb([]float64{gray})
	_bgg._cg = append(_bgg._cg, &_bba)
	return _bgg
}

// Process processes the entire list of operations. Maintains the graphics state that is passed to any
// handlers that are triggered during processing (either on specific operators or all).
func (_bcca *ContentStreamProcessor) Process(resources *_gb.PdfPageResources) error {
	_bcca._eced.ColorspaceStroking = _gb.NewPdfColorspaceDeviceGray()
	_bcca._eced.ColorspaceNonStroking = _gb.NewPdfColorspaceDeviceGray()
	_bcca._eced.ColorStroking = _gb.NewPdfColorDeviceGray(0)
	_bcca._eced.ColorNonStroking = _gb.NewPdfColorDeviceGray(0)
	_bcca._eced.CTM = _de.IdentityMatrix()
	for _, _aef := range _bcca._gcc {
		var _abcf error
		switch _aef.Operand {
		case "\u0071":
			_bcca._dcf.Push(_bcca._eced)
		case "\u0051":
			if len(_bcca._dcf) == 0 {
				_fb.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0060\u0051\u0060\u0020\u006f\u0070e\u0072\u0061\u0074\u006f\u0072\u002e\u0020\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0073\u0020\u0073\u0074\u0061\u0074\u0065 \u0073\u0074\u0061\u0063\u006b\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079.\u0020\u0053\u006bi\u0070\u0070\u0069\u006e\u0067\u002e")
				continue
			}
			_bcca._eced = _bcca._dcf.Pop()
		case "\u0043\u0053":
			_abcf = _bcca.handleCommand_CS(_aef, resources)
		case "\u0063\u0073":
			_abcf = _bcca.handleCommand_cs(_aef, resources)
		case "\u0053\u0043":
			_abcf = _bcca.handleCommand_SC(_aef, resources)
		case "\u0053\u0043\u004e":
			_abcf = _bcca.handleCommand_SCN(_aef, resources)
		case "\u0073\u0063":
			_abcf = _bcca.handleCommand_sc(_aef, resources)
		case "\u0073\u0063\u006e":
			_abcf = _bcca.handleCommand_scn(_aef, resources)
		case "\u0047":
			_abcf = _bcca.handleCommand_G(_aef, resources)
		case "\u0067":
			_abcf = _bcca.handleCommand_g(_aef, resources)
		case "\u0052\u0047":
			_abcf = _bcca.handleCommand_RG(_aef, resources)
		case "\u0072\u0067":
			_abcf = _bcca.handleCommand_rg(_aef, resources)
		case "\u004b":
			_abcf = _bcca.handleCommand_K(_aef, resources)
		case "\u006b":
			_abcf = _bcca.handleCommand_k(_aef, resources)
		case "\u0063\u006d":
			_abcf = _bcca.handleCommand_cm(_aef, resources)
		}
		if _abcf != nil {
			_fb.Log.Debug("\u0050\u0072\u006f\u0063\u0065\u0073s\u006f\u0072\u0020\u0068\u0061\u006e\u0064\u006c\u0069\u006e\u0067\u0020\u0065r\u0072\u006f\u0072\u0020\u0028\u0025\u0073)\u003a\u0020\u0025\u0076", _aef.Operand, _abcf)
			_fb.Log.Debug("\u004f\u0070\u0065r\u0061\u006e\u0064\u003a\u0020\u0025\u0023\u0076", _aef.Operand)
			return _abcf
		}
		for _, _ggd := range _bcca._efd {
			var _daff error
			if _ggd.Condition.All() {
				_daff = _ggd.Handler(_aef, _bcca._eced, resources)
			} else if _ggd.Condition.Operand() && _aef.Operand == _ggd.Operand {
				_daff = _ggd.Handler(_aef, _bcca._eced, resources)
			}
			if _daff != nil {
				_fb.Log.Debug("P\u0072\u006f\u0063\u0065\u0073\u0073o\u0072\u0020\u0068\u0061\u006e\u0064\u006c\u0065\u0072 \u0065\u0072\u0072o\u0072:\u0020\u0025\u0076", _daff)
				return _daff
			}
		}
	}
	return nil
}

func (_bda *ContentStreamParser) parseObject() (_gdgg _ed.PdfObject, _eabg bool, _adda error) {
	_bda.skipSpaces()
	for {
		_eebf, _gdcg := _bda._cac.Peek(2)
		if _gdcg != nil {
			return nil, false, _gdcg
		}
		_fb.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_eebf))
		if _eebf[0] == '%' {
			_bda.skipComments()
			continue
		} else if _eebf[0] == '/' {
			_ccbf, _ebfbd := _bda.parseName()
			_fb.Log.Trace("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027", _ccbf)
			return &_ccbf, false, _ebfbd
		} else if _eebf[0] == '(' {
			_fb.Log.Trace("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!")
			_cddg, _ebef := _bda.parseString()
			return _cddg, false, _ebef
		} else if _eebf[0] == '<' && _eebf[1] != '<' {
			_fb.Log.Trace("\u002d\u003e\u0048\u0065\u0078\u0020\u0053\u0074\u0072\u0069\u006e\u0067\u0021")
			_dag, _dbaf := _bda.parseHexString()
			return _dag, false, _dbaf
		} else if _eebf[0] == '[' {
			_fb.Log.Trace("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021")
			_dcgf, _ddf := _bda.parseArray()
			return _dcgf, false, _ddf
		} else if _ed.IsFloatDigit(_eebf[0]) || (_eebf[0] == '-' && _ed.IsFloatDigit(_eebf[1])) || (_eebf[0] == '+' && _ed.IsFloatDigit(_eebf[1])) {
			_fb.Log.Trace("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!")
			_daf, _bfcc := _bda.parseNumber()
			return _daf, false, _bfcc
		} else if _eebf[0] == '<' && _eebf[1] == '<' {
			_bfb, _edbf := _bda.parseDict()
			return _bfb, false, _edbf
		} else {
			_fb.Log.Trace("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f")
			_eebf, _ = _bda._cac.Peek(5)
			_fcgb := string(_eebf)
			_fb.Log.Trace("\u0063\u006f\u006e\u0074\u0020\u0050\u0065\u0065\u006b\u0020\u0073\u0074r\u003a\u0020\u0025\u0073", _fcgb)
			if (len(_fcgb) > 3) && (_fcgb[:4] == "\u006e\u0075\u006c\u006c") {
				_fbdd, _bdce := _bda.parseNull()
				return &_fbdd, false, _bdce
			} else if (len(_fcgb) > 4) && (_fcgb[:5] == "\u0066\u0061\u006cs\u0065") {
				_acfgg, _bdee := _bda.parseBool()
				return &_acfgg, false, _bdee
			} else if (len(_fcgb) > 3) && (_fcgb[:4] == "\u0074\u0072\u0075\u0065") {
				_affc, _bgbd := _bda.parseBool()
				return &_affc, false, _bgbd
			}
			_aac, _gaca := _bda.parseOperand()
			if _gaca != nil {
				return _aac, false, _gaca
			}
			if len(_aac.String()) < 1 {
				return _aac, false, ErrInvalidOperand
			}
			return _aac, true, nil
		}
	}
}

// Add_CS appends 'CS' operand to the content stream:
// Set the current colorspace for stroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_dfb *ContentCreator) Add_CS(name _ed.PdfObjectName) *ContentCreator {
	_fgbe := ContentStreamOperation{}
	_fgbe.Operand = "\u0043\u0053"
	_fgbe.Params = _gfae([]_ed.PdfObjectName{name})
	_dfb._cg = append(_dfb._cg, &_fgbe)
	return _dfb
}

func _dgcf(_fafd *ContentStreamInlineImage) (*_ed.DCTEncoder, error) {
	_dea := _ed.NewDCTEncoder()
	_feeb := _e.NewReader(_fafd._bce)
	_fcaf, _aab := _bbg.DecodeConfig(_feeb)
	if _aab != nil {
		_fb.Log.Debug("\u0045\u0072\u0072or\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _aab)
		return nil, _aab
	}
	switch _fcaf.ColorModel {
	case _gf.RGBAModel:
		_dea.BitsPerComponent = 8
		_dea.ColorComponents = 3
	case _gf.RGBA64Model:
		_dea.BitsPerComponent = 16
		_dea.ColorComponents = 3
	case _gf.GrayModel:
		_dea.BitsPerComponent = 8
		_dea.ColorComponents = 1
	case _gf.Gray16Model:
		_dea.BitsPerComponent = 16
		_dea.ColorComponents = 1
	case _gf.CMYKModel:
		_dea.BitsPerComponent = 8
		_dea.ColorComponents = 4
	case _gf.YCbCrModel:
		_dea.BitsPerComponent = 8
		_dea.ColorComponents = 3
	default:
		return nil, _b.New("\u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006d\u006f\u0064\u0065\u006c")
	}
	_dea.Width = _fcaf.Width
	_dea.Height = _fcaf.Height
	_fb.Log.Trace("\u0044\u0043T\u0020\u0045\u006ec\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076", _dea)
	return _dea, nil
}

// ExtractText parses and extracts all text data in content streams and returns as a string.
// Does not take into account Encoding table, the output is simply the character codes.
//
// Deprecated: More advanced text extraction is offered in package extractor with character encoding support.
func (_cce *ContentStreamParser) ExtractText() (string, error) {
	_bf, _dd := _cce.Parse()
	if _dd != nil {
		return "", _dd
	}
	_gfc := false
	_da, _fdf := float64(-1), float64(-1)
	_eaa := ""
	for _, _cd := range *_bf {
		if _cd.Operand == "\u0042\u0054" {
			_gfc = true
		} else if _cd.Operand == "\u0045\u0054" {
			_gfc = false
		}
		if _cd.Operand == "\u0054\u0064" || _cd.Operand == "\u0054\u0044" || _cd.Operand == "\u0054\u002a" {
			_eaa += "\u000a"
		}
		if _cd.Operand == "\u0054\u006d" {
			if len(_cd.Params) != 6 {
				continue
			}
			_eeb, _fgb := _cd.Params[4].(*_ed.PdfObjectFloat)
			if !_fgb {
				_aff, _cdc := _cd.Params[4].(*_ed.PdfObjectInteger)
				if !_cdc {
					continue
				}
				_eeb = _ed.MakeFloat(float64(*_aff))
			}
			_fba, _fgb := _cd.Params[5].(*_ed.PdfObjectFloat)
			if !_fgb {
				_df, _ad := _cd.Params[5].(*_ed.PdfObjectInteger)
				if !_ad {
					continue
				}
				_fba = _ed.MakeFloat(float64(*_df))
			}
			if _fdf == -1 {
				_fdf = float64(*_fba)
			} else if _fdf > float64(*_fba) {
				_eaa += "\u000a"
				_da = float64(*_eeb)
				_fdf = float64(*_fba)
				continue
			}
			if _da == -1 {
				_da = float64(*_eeb)
			} else if _da < float64(*_eeb) {
				_eaa += "\u0009"
				_da = float64(*_eeb)
			}
		}
		if _gfc && _cd.Operand == "\u0054\u004a" {
			if len(_cd.Params) < 1 {
				continue
			}
			_eeg, _cdcd := _cd.Params[0].(*_ed.PdfObjectArray)
			if !_cdcd {
				return "", _fe.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0070\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0020\u0074y\u0070\u0065\u002c\u0020\u006e\u006f\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _cd.Params[0])
			}
			for _, _gc := range _eeg.Elements() {
				switch _fa := _gc.(type) {
				case *_ed.PdfObjectString:
					_eaa += _fa.Str()
				case *_ed.PdfObjectFloat:
					if *_fa < -100 {
						_eaa += "\u0020"
					}
				case *_ed.PdfObjectInteger:
					if *_fa < -100 {
						_eaa += "\u0020"
					}
				}
			}
		} else if _gfc && _cd.Operand == "\u0054\u006a" {
			if len(_cd.Params) < 1 {
				continue
			}
			_gfa, _bd := _cd.Params[0].(*_ed.PdfObjectString)
			if !_bd {
				return "", _fe.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072\u0020\u0074\u0079p\u0065\u002c\u0020\u006e\u006f\u0074\u0020\u0073\u0074\u0072\u0069\u006e\u0067 \u0028\u0025\u0054\u0029", _cd.Params[0])
			}
			_eaa += _gfa.Str()
		}
	}
	return _eaa, nil
}

func (_gacg *ContentStreamProcessor) handleCommand_SC(_cdaa *ContentStreamOperation, _dce *_gb.PdfPageResources) error {
	_acca := _gacg._eced.ColorspaceStroking
	if len(_cdaa.Params) != _acca.GetNumComponents() {
		_fb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_fb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_cdaa.Params), _acca)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_fdfa, _bcac := _acca.ColorFromPdfObjects(_cdaa.Params)
	if _bcac != nil {
		return _bcac
	}
	_gacg._eced.ColorStroking = _fdfa
	return nil
}

// Add_ri adds 'ri' operand to the content stream, which sets the color rendering intent.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_gd *ContentCreator) Add_ri(intent _ed.PdfObjectName) *ContentCreator {
	_fea := ContentStreamOperation{}
	_fea.Operand = "\u0072\u0069"
	_fea.Params = _gfae([]_ed.PdfObjectName{intent})
	_gd._cg = append(_gd._cg, &_fea)
	return _gd
}

func (_gdgb *ContentStreamParser) parseBool() (_ed.PdfObjectBool, error) {
	_dae, _dcge := _gdgb._cac.Peek(4)
	if _dcge != nil {
		return _ed.PdfObjectBool(false), _dcge
	}
	if (len(_dae) >= 4) && (string(_dae[:4]) == "\u0074\u0072\u0075\u0065") {
		_gdgb._cac.Discard(4)
		return _ed.PdfObjectBool(true), nil
	}
	_dae, _dcge = _gdgb._cac.Peek(5)
	if _dcge != nil {
		return _ed.PdfObjectBool(false), _dcge
	}
	if (len(_dae) >= 5) && (string(_dae[:5]) == "\u0066\u0061\u006cs\u0065") {
		_gdgb._cac.Discard(5)
		return _ed.PdfObjectBool(false), nil
	}
	return _ed.PdfObjectBool(false), _b.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

func (_aabb *ContentStreamProcessor) handleCommand_SCN(_bgcc *ContentStreamOperation, _bbac *_gb.PdfPageResources) error {
	_egeg := _aabb._eced.ColorspaceStroking
	if !_eff(_egeg) {
		if len(_bgcc.Params) != _egeg.GetNumComponents() {
			_fb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_fb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_bgcc.Params), _egeg)
			return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_cgba, _gbaa := _egeg.ColorFromPdfObjects(_bgcc.Params)
	if _gbaa != nil {
		return _gbaa
	}
	_aabb._eced.ColorStroking = _cgba
	return nil
}

// Add_Tr appends 'Tr' operand to the content stream:
// Set text rendering mode.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_gdfg *ContentCreator) Add_Tr(render int64) *ContentCreator {
	_bdfd := ContentStreamOperation{}
	_bdfd.Operand = "\u0054\u0072"
	_bdfd.Params = _bfgc([]int64{render})
	_gdfg._cg = append(_gdfg._cg, &_bdfd)
	return _gdfg
}

var _bdd = _f.MustCompile("\u005e\u002f\u007b\u0032\u002c\u007d")

func _bfgc(_bbce []int64) []_ed.PdfObject {
	var _gec []_ed.PdfObject
	for _, _fggdc := range _bbce {
		_gec = append(_gec, _ed.MakeInteger(_fggdc))
	}
	return _gec
}

func (_egab *ContentStreamProcessor) handleCommand_G(_eabb *ContentStreamOperation, _adb *_gb.PdfPageResources) error {
	_daab := _gb.NewPdfColorspaceDeviceGray()
	if len(_eabb.Params) != _daab.GetNumComponents() {
		_fb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_fb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_eabb.Params), _daab)
		return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_dagd, _dbcb := _daab.ColorFromPdfObjects(_eabb.Params)
	if _dbcb != nil {
		return _dbcb
	}
	_egab._eced.ColorspaceStroking = _daab
	_egab._eced.ColorStroking = _dagd
	return nil
}

// Add_k appends 'k' operand to the content stream:
// Same as K but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_ebg *ContentCreator) Add_k(c, m, y, k float64) *ContentCreator {
	_fac := ContentStreamOperation{}
	_fac.Operand = "\u006b"
	_fac.Params = _gfbfb([]float64{c, m, y, k})
	_ebg._cg = append(_ebg._cg, &_fac)
	return _ebg
}

// ContentStreamOperation represents an operation in PDF contentstream which consists of
// an operand and parameters.
type ContentStreamOperation struct {
	Params  []_ed.PdfObject
	Operand string
}

// Add_c adds 'c' operand to the content stream: Append a Bezier curve to the current path from
// the current point to (x3,y3) with (x1,x1) and (x2,y2) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_ae *ContentCreator) Add_c(x1, y1, x2, y2, x3, y3 float64) *ContentCreator {
	_ecg := ContentStreamOperation{}
	_ecg.Operand = "\u0063"
	_ecg.Params = _gfbfb([]float64{x1, y1, x2, y2, x3, y3})
	_ae._cg = append(_ae._cg, &_ecg)
	return _ae
}

func (_fec *ContentStreamProcessor) handleCommand_sc(_ecff *ContentStreamOperation, _dcfg *_gb.PdfPageResources) error {
	_fbfc := _fec._eced.ColorspaceNonStroking
	if !_eff(_fbfc) {
		if len(_ecff.Params) != _fbfc.GetNumComponents() {
			_fb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_fb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_ecff.Params), _fbfc)
			return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_bddb, _bgfe := _fbfc.ColorFromPdfObjects(_ecff.Params)
	if _bgfe != nil {
		return _bgfe
	}
	_fec._eced.ColorNonStroking = _bddb
	return nil
}

// AddHandler adds a new ContentStreamProcessor `handler` of type `condition` for `operand`.
func (_abc *ContentStreamProcessor) AddHandler(condition HandlerConditionEnum, operand string, handler HandlerFunc) {
	_aaae := handlerEntry{}
	_aaae.Condition = condition
	_aaae.Operand = operand
	_aaae.Handler = handler
	_abc._efd = append(_abc._efd, _aaae)
}
func _gfd(_gdc string) bool { _, _daca := _cdf[_gdc]; return _daca }

type handlerEntry struct {
	Condition HandlerConditionEnum
	Operand   string
	Handler   HandlerFunc
}

// AddOperand adds a specified operand.
func (_dab *ContentCreator) AddOperand(op ContentStreamOperation) *ContentCreator {
	_dab._cg = append(_dab._cg, &op)
	return _dab
}

func (_dfgg *ContentStreamProcessor) handleCommand_scn(_bafb *ContentStreamOperation, _fegb *_gb.PdfPageResources) error {
	_bceg := _dfgg._eced.ColorspaceNonStroking
	if !_eff(_bceg) {
		if len(_bafb.Params) != _bceg.GetNumComponents() {
			_fb.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_fb.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_bafb.Params), _bceg)
			return _b.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_cfa, _fdfb := _bceg.ColorFromPdfObjects(_bafb.Params)
	if _fdfb != nil {
		_fb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c \u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0063o\u006co\u0072\u0020\u0066\u0072\u006f\u006d\u0020\u0070\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076 \u0028\u0043\u0053\u0020\u0069\u0073\u0020\u0025\u002b\u0076\u0029", _bafb.Params, _bceg)
		return _fdfb
	}
	_dfgg._eced.ColorNonStroking = _cfa
	return nil
}

// NewContentStreamProcessor returns a new ContentStreamProcessor for operations `ops`.
func NewContentStreamProcessor(ops []*ContentStreamOperation) *ContentStreamProcessor {
	_bcag := ContentStreamProcessor{}
	_bcag._dcf = GraphicStateStack{}
	_agga := GraphicsState{}
	_bcag._eced = _agga
	_bcag._efd = []handlerEntry{}
	_bcag._bgff = 0
	_bcag._gcc = ops
	return &_bcag
}

// Pop pops and returns the topmost GraphicsState off the `gsStack`.
func (_aabd *GraphicStateStack) Pop() GraphicsState {
	_fcgd := (*_aabd)[len(*_aabd)-1]
	*_aabd = (*_aabd)[:len(*_aabd)-1]
	return _fcgd
}

// All returns true if `hce` is equivalent to HandlerConditionEnumAllOperands.
func (_eac HandlerConditionEnum) All() bool { return _eac == HandlerConditionEnumAllOperands }

func (_fbac *ContentStreamProcessor) handleCommand_CS(_eacb *ContentStreamOperation, _acbf *_gb.PdfPageResources) error {
	if len(_eacb.Params) < 1 {
		_fb.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0063\u0073\u0020\u0063\u006f\u006d\u006d\u0061\u006e\u0064\u002c\u0020s\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u006f\u0076\u0065\u0072")
		return _b.New("\u0074o\u006f \u0066\u0065\u0077\u0020\u0070a\u0072\u0061m\u0065\u0074\u0065\u0072\u0073")
	}
	if len(_eacb.Params) > 1 {
		_fb.Log.Debug("\u0063\u0073\u0020\u0063\u006f\u006d\u006d\u0061n\u0064\u0020\u0077it\u0068\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020-\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075i\u006e\u0067")
		return _b.New("\u0074\u006f\u006f\u0020ma\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	}
	_daad, _ebag := _eacb.Params[0].(*_ed.PdfObjectName)
	if !_ebag {
		_fb.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020c\u0073\u0020\u0063o\u006d\u006d\u0061n\u0064\u0020w\u0069\u0074\u0068\u0020\u0069\u006ev\u0061li\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u002c\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072")
		return _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_egc, _bgfg := _fbac.getColorspace(string(*_daad), _acbf)
	if _bgfg != nil {
		return _bgfg
	}
	_fbac._eced.ColorspaceStroking = _egc
	_eacc, _bgfg := _fbac.getInitialColor(_egc)
	if _bgfg != nil {
		return _bgfg
	}
	_fbac._eced.ColorStroking = _eacc
	return nil
}

// Bytes converts the content stream operations to a content stream byte presentation, i.e. the kind that can be
// stored as a PDF stream or string format.
func (_cgd *ContentCreator) Bytes() []byte { return _cgd._cg.Bytes() }

func (_dgdb *ContentStreamProcessor) getInitialColor(_faab _gb.PdfColorspace) (_gb.PdfColor, error) {
	switch _decf := _faab.(type) {
	case *_gb.PdfColorspaceDeviceGray:
		return _gb.NewPdfColorDeviceGray(0.0), nil
	case *_gb.PdfColorspaceDeviceRGB:
		return _gb.NewPdfColorDeviceRGB(0.0, 0.0, 0.0), nil
	case *_gb.PdfColorspaceDeviceCMYK:
		return _gb.NewPdfColorDeviceCMYK(0.0, 0.0, 0.0, 1.0), nil
	case *_gb.PdfColorspaceCalGray:
		return _gb.NewPdfColorCalGray(0.0), nil
	case *_gb.PdfColorspaceCalRGB:
		return _gb.NewPdfColorCalRGB(0.0, 0.0, 0.0), nil
	case *_gb.PdfColorspaceLab:
		_fef := 0.0
		_caag := 0.0
		_ebdf := 0.0
		if _decf.Range[0] > 0 {
			_fef = _decf.Range[0]
		}
		if _decf.Range[2] > 0 {
			_caag = _decf.Range[2]
		}
		return _gb.NewPdfColorLab(_fef, _caag, _ebdf), nil
	case *_gb.PdfColorspaceICCBased:
		if _decf.Alternate == nil {
			_fb.Log.Trace("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065d\u0020-\u0020\u0061\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0066\u0061\u006c\u006c\u0020\u0062a\u0063\u006b\u0020\u0028\u004e\u0020\u003d\u0020\u0025\u0064\u0029", _decf.N)
			if _decf.N == 1 {
				_fb.Log.Trace("\u0046\u0061\u006c\u006c\u0069\u006e\u0067\u0020\u0062\u0061\u0063k\u0020\u0074\u006f\u0020\u0044\u0065\u0076\u0069\u0063\u0065G\u0072\u0061\u0079")
				return _dgdb.getInitialColor(_gb.NewPdfColorspaceDeviceGray())
			} else if _decf.N == 3 {
				_fb.Log.Trace("\u0046a\u006c\u006c\u0069\u006eg\u0020\u0062\u0061\u0063\u006b \u0074o\u0020D\u0065\u0076\u0069\u0063\u0065\u0052\u0047B")
				return _dgdb.getInitialColor(_gb.NewPdfColorspaceDeviceRGB())
			} else if _decf.N == 4 {
				_fb.Log.Trace("\u0046\u0061\u006c\u006c\u0069\u006e\u0067\u0020\u0062\u0061\u0063k\u0020\u0074\u006f\u0020\u0044\u0065\u0076\u0069\u0063\u0065C\u004d\u0059\u004b")
				return _dgdb.getInitialColor(_gb.NewPdfColorspaceDeviceCMYK())
			} else {
				return nil, _b.New("a\u006c\u0074\u0065\u0072\u006e\u0061t\u0065\u0020\u0073\u0070\u0061\u0063e\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0049C\u0043")
			}
		}
		return _dgdb.getInitialColor(_decf.Alternate)
	case *_gb.PdfColorspaceSpecialIndexed:
		if _decf.Base == nil {
			return nil, _b.New("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0062\u0061\u0073e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069f\u0069\u0065\u0064")
		}
		return _dgdb.getInitialColor(_decf.Base)
	case *_gb.PdfColorspaceSpecialSeparation:
		if _decf.AlternateSpace == nil {
			return nil, _b.New("\u0061\u006ct\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069fi\u0065\u0064")
		}
		return _dgdb.getInitialColor(_decf.AlternateSpace)
	case *_gb.PdfColorspaceDeviceN:
		if _decf.AlternateSpace == nil {
			return nil, _b.New("\u0061\u006ct\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069fi\u0065\u0064")
		}
		return _dgdb.getInitialColor(_decf.AlternateSpace)
	case *_gb.PdfColorspaceSpecialPattern:
		return _gb.NewPdfColorPattern(), nil
	}
	_fb.Log.Debug("Un\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0066\u006f\u0072\u0020\u0075\u006e\u006b\u006e\u006fw\u006e \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065:\u0020\u0025T", _faab)
	return nil, _b.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065")
}

// Add_Tz appends 'Tz' operand to the content stream:
// Set horizontal scaling.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_gbcg *ContentCreator) Add_Tz(scale float64) *ContentCreator {
	_cdg := ContentStreamOperation{}
	_cdg.Operand = "\u0054\u007a"
	_cdg.Params = _gfbfb([]float64{scale})
	_gbcg._cg = append(_gbcg._cg, &_cdg)
	return _gbcg
}

func (_ebfb *ContentStreamParser) parseHexString() (*_ed.PdfObjectString, error) {
	_ebfb._cac.ReadByte()
	_cedg := []byte("\u0030\u0031\u0032\u003345\u0036\u0037\u0038\u0039\u0061\u0062\u0063\u0064\u0065\u0066\u0041\u0042\u0043\u0044E\u0046")
	var _feea []byte
	for {
		_ebfb.skipSpaces()
		_cdace, _bdc := _ebfb._cac.Peek(1)
		if _bdc != nil {
			return _ed.MakeString(""), _bdc
		}
		if _cdace[0] == '>' {
			_ebfb._cac.ReadByte()
			break
		}
		_gcab, _ := _ebfb._cac.ReadByte()
		if _e.IndexByte(_cedg, _gcab) >= 0 {
			_feea = append(_feea, _gcab)
		}
	}
	if len(_feea)%2 == 1 {
		_feea = append(_feea, '0')
	}
	_eba, _ := _bb.DecodeString(string(_feea))
	return _ed.MakeHexString(string(_eba)), nil
}

// Scale applies x-y scaling to the transformation matrix.
func (_db *ContentCreator) Scale(sx, sy float64) *ContentCreator {
	return _db.Add_cm(sx, 0, 0, sy, 0, 0)
}

// Add_TJ appends 'TJ' operand to the content stream:
// Show one or more text string. Array of numbers (displacement) and strings.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_adg *ContentCreator) Add_TJ(vals ..._ed.PdfObject) *ContentCreator {
	_ecd := ContentStreamOperation{}
	_ecd.Operand = "\u0054\u004a"
	_ecd.Params = []_ed.PdfObject{_ed.MakeArray(vals...)}
	_adg._cg = append(_adg._cg, &_ecd)
	return _adg
}

// ParseInlineImage parses an inline image from a content stream, both reading its properties and binary data.
// When called, "BI" has already been read from the stream.  This function
// finishes reading through "EI" and then returns the ContentStreamInlineImage.
func (_ebff *ContentStreamParser) ParseInlineImage() (*ContentStreamInlineImage, error) {
	_aaec := ContentStreamInlineImage{}
	for {
		_ebff.skipSpaces()
		_aeg, _dfc, _fge := _ebff.parseObject()
		if _fge != nil {
			return nil, _fge
		}
		if !_dfc {
			_ede, _cbbb := _ed.GetName(_aeg)
			if !_cbbb {
				_fb.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067e\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0020\u0028\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u0029\u0020\u002d\u0020\u0025T", _aeg)
				return nil, _fe.Errorf("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067e\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0020\u0028\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u0029\u0020\u002d\u0020\u0025T", _aeg)
			}
			_efce, _abg, _gcd := _ebff.parseObject()
			if _gcd != nil {
				return nil, _gcd
			}
			if _abg {
				return nil, _fe.Errorf("\u006eo\u0074\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067 \u0061\u006e\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
			}
			switch *_ede {
			case "\u0042\u0050\u0043", "\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074":
				_aaec.BitsPerComponent = _efce
			case "\u0043\u0053", "\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065":
				_aaec.ColorSpace = _efce
			case "\u0044", "\u0044\u0065\u0063\u006f\u0064\u0065":
				_aaec.Decode = _efce
			case "\u0044\u0050", "D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073":
				_aaec.DecodeParms = _efce
			case "\u0046", "\u0046\u0069\u006c\u0074\u0065\u0072":
				_aaec.Filter = _efce
			case "\u0048", "\u0048\u0065\u0069\u0067\u0068\u0074":
				_aaec.Height = _efce
			case "\u0049\u004d", "\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k":
				_aaec.ImageMask = _efce
			case "\u0049\u006e\u0074\u0065\u006e\u0074":
				_aaec.Intent = _efce
			case "\u0049", "I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065":
				_aaec.Interpolate = _efce
			case "\u0057", "\u0057\u0069\u0064t\u0068":
				_aaec.Width = _efce
			case "\u004c\u0065\u006e\u0067\u0074\u0068", "\u0053u\u0062\u0074\u0079\u0070\u0065", "\u0054\u0079\u0070\u0065":
				_fb.Log.Debug("\u0049\u0067\u006e\u006fr\u0069\u006e\u0067\u0020\u0069\u006e\u006c\u0069\u006e\u0065 \u0070a\u0072\u0061\u006d\u0065\u0074\u0065\u0072 \u0025\u0073", *_ede)
			default:
				return nil, _fe.Errorf("\u0075\u006e\u006b\u006e\u006f\u0077n\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0020\u0025\u0073", *_ede)
			}
		}
		if _dfc {
			_dbc, _bab := _aeg.(*_ed.PdfObjectString)
			if !_bab {
				return nil, _fe.Errorf("\u0066a\u0069\u006ce\u0064\u0020\u0074o\u0020\u0072\u0065\u0061\u0064\u0020\u0069n\u006c\u0069\u006e\u0065\u0020\u0069m\u0061\u0067\u0065\u0020\u002d\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
			}
			if _dbc.Str() == "\u0045\u0049" {
				_fb.Log.Trace("\u0049n\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020f\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e\u002e\u002e")
				return &_aaec, nil
			} else if _dbc.Str() == "\u0049\u0044" {
				_fb.Log.Trace("\u0049\u0044\u0020\u0073\u0074\u0061\u0072\u0074")
				_cgdf, _agab := _ebff._cac.Peek(1)
				if _agab != nil {
					return nil, _agab
				}
				if _ed.IsWhiteSpace(_cgdf[0]) {
					_ebff._cac.Discard(1)
				}
				_aaec._bce = []byte{}
				_abb := 0
				var _dcc []byte
				for {
					_ace, _acgf := _ebff._cac.ReadByte()
					if _acgf != nil {
						_fb.Log.Debug("\u0055\u006e\u0061\u0062\u006ce\u0020\u0074\u006f\u0020\u0066\u0069\u006e\u0064\u0020\u0065\u006e\u0064\u0020o\u0066\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0045\u0049\u0020\u0069\u006e\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u0061\u0074a")
						return nil, _acgf
					}
					if _abb == 0 {
						if _ed.IsWhiteSpace(_ace) {
							_dcc = []byte{}
							_dcc = append(_dcc, _ace)
							_abb = 1
						} else if _ace == 'E' {
							_dcc = append(_dcc, _ace)
							_abb = 2
						} else {
							_aaec._bce = append(_aaec._bce, _ace)
						}
					} else if _abb == 1 {
						_dcc = append(_dcc, _ace)
						if _ace == 'E' {
							_abb = 2
						} else {
							_aaec._bce = append(_aaec._bce, _dcc...)
							_dcc = []byte{}
							if _ed.IsWhiteSpace(_ace) {
								_abb = 1
							} else {
								_abb = 0
							}
						}
					} else if _abb == 2 {
						_dcc = append(_dcc, _ace)
						if _ace == 'I' {
							_abb = 3
						} else {
							_aaec._bce = append(_aaec._bce, _dcc...)
							_dcc = []byte{}
							_abb = 0
						}
					} else if _abb == 3 {
						_dcc = append(_dcc, _ace)
						if _ed.IsWhiteSpace(_ace) {
							_fbfd, _cffa := _ebff._cac.Peek(20)
							if _cffa != nil && _cffa != _ff.EOF {
								return nil, _cffa
							}
							_acfa := NewContentStreamParser(string(_fbfd))
							_adaa := true
							for _eec := 0; _eec < 3; _eec++ {
								_bdeg, _fege, _gad := _acfa.parseObject()
								if _gad != nil {
									if _gad == _ff.EOF {
										break
									}
									_adaa = false
									continue
								}
								if _fege && !_gfd(_bdeg.String()) {
									_adaa = false
									break
								}
							}
							if _adaa {
								if len(_aaec._bce) > 100 {
									_fb.Log.Trace("\u0049\u006d\u0061\u0067\u0065\u0020\u0073\u0074\u0072\u0065\u0061m\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078 \u002e\u002e\u002e", len(_aaec._bce), _aaec._bce[:100])
								} else {
									_fb.Log.Trace("\u0049\u006d\u0061\u0067e \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025 \u0078", len(_aaec._bce), _aaec._bce)
								}
								return &_aaec, nil
							}
						}
						_aaec._bce = append(_aaec._bce, _dcc...)
						_dcc = []byte{}
						_abb = 0
					}
				}
			}
		}
	}
}

func _ecfd(_faff *ContentStreamInlineImage, _gacd *_ed.PdfObjectDictionary) (*_ed.LZWEncoder, error) {
	_cbba := _ed.NewLZWEncoder()
	if _gacd == nil {
		if _faff.DecodeParms != nil {
			_bec, _bdbf := _ed.GetDict(_faff.DecodeParms)
			if !_bdbf {
				_fb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079 \u0028%\u0054\u0029", _faff.DecodeParms)
				return nil, _fe.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
			_gacd = _bec
		}
	}
	if _gacd == nil {
		return _cbba, nil
	}
	_fee := _gacd.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
	if _fee != nil {
		_bgeg, _ecc := _fee.(*_ed.PdfObjectInteger)
		if !_ecc {
			_fb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a \u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u006e\u0075\u006d\u0065\u0072i\u0063 \u0028\u0025\u0054\u0029", _fee)
			return nil, _fe.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
		}
		if *_bgeg != 0 && *_bgeg != 1 {
			return nil, _fe.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0076\u0061\u006c\u0075e\u0020\u0028\u006e\u006f\u0074 \u0030\u0020o\u0072\u0020\u0031\u0029")
		}
		_cbba.EarlyChange = int(*_bgeg)
	} else {
		_cbba.EarlyChange = 1
	}
	_fee = _gacd.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _fee != nil {
		_fdfe, _ccea := _fee.(*_ed.PdfObjectInteger)
		if !_ccea {
			_fb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _fee)
			return nil, _fe.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_cbba.Predictor = int(*_fdfe)
	}
	_fee = _gacd.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _fee != nil {
		_dfg, _aae := _fee.(*_ed.PdfObjectInteger)
		if !_aae {
			_fb.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _fe.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_cbba.BitsPerComponent = int(*_dfg)
	}
	if _cbba.Predictor > 1 {
		_cbba.Columns = 1
		_fee = _gacd.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _fee != nil {
			_bae, _feg := _fee.(*_ed.PdfObjectInteger)
			if !_feg {
				return nil, _fe.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_cbba.Columns = int(*_bae)
		}
		_cbba.Colors = 1
		_fee = _gacd.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _fee != nil {
			_ggf, _bggg := _fee.(*_ed.PdfObjectInteger)
			if !_bggg {
				return nil, _fe.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_cbba.Colors = int(*_ggf)
		}
	}
	_fb.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _gacd.String())
	return _cbba, nil
}

// Add_b appends 'b' operand to the content stream:
// Close, fill and then stroke the path (nonzero winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_cff *ContentCreator) Add_b() *ContentCreator {
	_cfe := ContentStreamOperation{}
	_cfe.Operand = "\u0062"
	_cff._cg = append(_cff._cg, &_cfe)
	return _cff
}

// Add_Td appends 'Td' operand to the content stream:
// Move to start of next line with offset (`tx`, `ty`).
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_bdb *ContentCreator) Add_Td(tx, ty float64) *ContentCreator {
	_aca := ContentStreamOperation{}
	_aca.Operand = "\u0054\u0064"
	_aca.Params = _gfbfb([]float64{tx, ty})
	_bdb._cg = append(_bdb._cg, &_aca)
	return _bdb
}

// Add_sh appends 'sh' operand to the content stream:
// Paints the shape and colour shading described by a shading dictionary specified by `name`,
// subject to the current clipping path
//
// See section 8.7.4 "Shading Patterns" and Table 77 (p. 190 PDF32000_2008).
func (_fdab *ContentCreator) Add_sh(name _ed.PdfObjectName) *ContentCreator {
	_fdce := ContentStreamOperation{}
	_fdce.Operand = "\u0073\u0068"
	_fdce.Params = _gfae([]_ed.PdfObjectName{name})
	_fdab._cg = append(_fdab._cg, &_fdce)
	return _fdab
}

func _bbff(_afd _ed.PdfObject) (_gb.PdfColorspace, error) {
	_fdd, _abbb := _afd.(*_ed.PdfObjectArray)
	if !_abbb {
		_fb.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0049\u006ev\u0061\u006c\u0069d\u0020\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020cs\u0020\u006e\u006ft\u0020\u0069n\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025#\u0076\u0029", _afd)
		return nil, _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _fdd.Len() != 4 {
		_fb.Log.Debug("\u0045\u0072\u0072\u006f\u0072:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061r\u0072\u0061\u0079\u002c\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0034\u0020\u0028\u0025\u0064\u0029", _fdd.Len())
		return nil, _b.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_fgce, _abbb := _fdd.Get(0).(*_ed.PdfObjectName)
	if !_abbb {
		_fb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0066\u0069\u0072s\u0074 \u0065\u006c\u0065\u006de\u006e\u0074 \u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0023\u0076\u0029", *_fdd)
		return nil, _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_fgce != "\u0049" && *_fgce != "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		_fb.Log.Debug("\u0045\u0072r\u006f\u0072\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064 \u0063\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0021\u003d\u0020\u0049\u0020\u0028\u0067\u006f\u0074\u003a\u0020\u0025\u0076\u0029", *_fgce)
		return nil, _b.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_fgce, _abbb = _fdd.Get(1).(*_ed.PdfObjectName)
	if !_abbb {
		_fb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072r\u0061\u0079\u0020\u0032\u006e\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0061\u0072\u0072a\u0079\u003a\u0020\u0025\u0023v\u0029", *_fdd)
		return nil, _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_fgce != "\u0047" && *_fgce != "\u0052\u0047\u0042" && *_fgce != "\u0043\u004d\u0059\u004b" && *_fgce != "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079" && *_fgce != "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B" && *_fgce != "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b" {
		_fb.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072r\u0061\u0079\u0020\u0032\u006e\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0021\u003d\u0020\u0047\u002f\u0052\u0047\u0042\u002f\u0043\u004d\u0059\u004b\u0020\u0028g\u006f\u0074\u003a\u0020\u0025v\u0029", *_fgce)
		return nil, _b.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_gbdg := ""
	switch *_fgce {
	case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
		_gbdg = "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079"
	case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
		_gbdg = "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B"
	case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
		_gbdg = "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
	}
	_abd := _ed.MakeArray(_ed.MakeName("\u0049n\u0064\u0065\u0078\u0065\u0064"), _ed.MakeName(_gbdg), _fdd.Get(2), _fdd.Get(3))
	return _gb.NewPdfColorspaceFromPdfObject(_abd)
}

// ToImage exports the inline image to Image which can be transformed or exported easily.
// Page resources are needed to look up colorspace information.
func (_beda *ContentStreamInlineImage) ToImage(resources *_gb.PdfPageResources) (*_gb.Image, error) {
	_edd, _bde := _beda.toImageBase(resources)
	if _bde != nil {
		return nil, _bde
	}
	_afb, _bde := _egb(_beda)
	if _bde != nil {
		return nil, _bde
	}
	_eaf, _ecca := _ed.GetDict(_beda.DecodeParms)
	if _ecca {
		_afb.UpdateParams(_eaf)
	}
	_fb.Log.Trace("\u0065n\u0063o\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076\u0020\u0025\u0054", _afb, _afb)
	_fb.Log.Trace("\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065:\u0020\u0025\u002b\u0076", _beda)
	_dged, _bde := _afb.DecodeBytes(_beda._bce)
	if _bde != nil {
		return nil, _bde
	}
	_aag := &_gb.Image{Width: int64(_edd.Width), Height: int64(_edd.Height), BitsPerComponent: int64(_edd.BitsPerComponent), ColorComponents: _edd.ColorComponents, Data: _dged}
	if len(_edd.Decode) > 0 {
		for _aaaa := 0; _aaaa < len(_edd.Decode); _aaaa++ {
			_edd.Decode[_aaaa] *= float64((int(1) << uint(_edd.BitsPerComponent)) - 1)
		}
		_aag.SetDecode(_edd.Decode)
	}
	return _aag, nil
}

// GetColorSpace returns the colorspace of the inline image.
func (_efb *ContentStreamInlineImage) GetColorSpace(resources *_gb.PdfPageResources) (_gb.PdfColorspace, error) {
	if _efb.ColorSpace == nil {
		_fb.Log.Debug("\u0049\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076i\u006e\u0067\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u002c\u0020\u0061\u0073\u0073\u0075\u006di\u006e\u0067\u0020\u0047\u0072a\u0079")
		return _gb.NewPdfColorspaceDeviceGray(), nil
	}
	if _dfff, _gbg := _efb.ColorSpace.(*_ed.PdfObjectArray); _gbg {
		return _bbff(_dfff)
	}
	_gbdf, _bca := _efb.ColorSpace.(*_ed.PdfObjectName)
	if !_bca {
		_fb.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020I\u006e\u0076\u0061\u006c\u0069\u0064 \u006f\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u003b\u0025\u002bv\u0029", _efb.ColorSpace, _efb.ColorSpace)
		return nil, _b.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_gbdf == "\u0047" || *_gbdf == "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079" {
		return _gb.NewPdfColorspaceDeviceGray(), nil
	} else if *_gbdf == "\u0052\u0047\u0042" || *_gbdf == "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B" {
		return _gb.NewPdfColorspaceDeviceRGB(), nil
	} else if *_gbdf == "\u0043\u004d\u0059\u004b" || *_gbdf == "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b" {
		return _gb.NewPdfColorspaceDeviceCMYK(), nil
	} else if *_gbdf == "\u0049" || *_gbdf == "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		return nil, _b.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0049\u006e\u0064e\u0078 \u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065")
	} else {
		if resources.ColorSpace == nil {
			_fb.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0073", *_gbdf)
			return nil, _b.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0063o\u006c\u006fr\u0073\u0070\u0061\u0063\u0065")
		}
		_bbe, _edc := resources.GetColorspaceByName(*_gbdf)
		if !_edc {
			_fb.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0073", *_gbdf)
			return nil, _b.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0063o\u006c\u006fr\u0073\u0070\u0061\u0063\u0065")
		}
		return _bbe, nil
	}
}

// WrapIfNeeded wraps the entire contents within q ... Q.  If unbalanced, then adds extra Qs at the end.
// Only does if needed. Ensures that when adding new content, one start with all states
// in the default condition.
func (_cc *ContentStreamOperations) WrapIfNeeded() *ContentStreamOperations {
	if len(*_cc) == 0 {
		return _cc
	}
	if _cc.isWrapped() {
		return _cc
	}
	*_cc = append([]*ContentStreamOperation{{Operand: "\u0071"}}, *_cc...)
	_gg := 0
	for _, _ca := range *_cc {
		if _ca.Operand == "\u0071" {
			_gg++
		} else if _ca.Operand == "\u0051" {
			_gg--
		}
	}
	for _gg > 0 {
		*_cc = append(*_cc, &ContentStreamOperation{Operand: "\u0051"})
		_gg--
	}
	return _cc
}

// Add_f appends 'f' operand to the content stream:
// Fill the path using the nonzero winding number rule to determine fill region.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_dec *ContentCreator) Add_f() *ContentCreator {
	_cfd := ContentStreamOperation{}
	_cfd.Operand = "\u0066"
	_dec._cg = append(_dec._cg, &_cfd)
	return _dec
}

// Add_Tw appends 'Tw' operand to the content stream:
// Set word spacing.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_dgc *ContentCreator) Add_Tw(wordSpace float64) *ContentCreator {
	_cagb := ContentStreamOperation{}
	_cagb.Operand = "\u0054\u0077"
	_cagb.Params = _gfbfb([]float64{wordSpace})
	_dgc._cg = append(_dgc._cg, &_cagb)
	return _dgc
}

// Add_b_starred appends 'b*' operand to the content stream:
// Close, fill and then stroke the path (even-odd winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_gfgga *ContentCreator) Add_b_starred() *ContentCreator {
	_bga := ContentStreamOperation{}
	_bga.Operand = "\u0062\u002a"
	_gfgga._cg = append(_gfgga._cg, &_bga)
	return _gfgga
}

// Add_n appends 'n' operand to the content stream:
// End the path without filling or stroking.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_bff *ContentCreator) Add_n() *ContentCreator {
	_gff := ContentStreamOperation{}
	_gff.Operand = "\u006e"
	_bff._cg = append(_bff._cg, &_gff)
	return _bff
}

func (_afgdc *ContentStreamParser) skipSpaces() (int, error) {
	_dga := 0
	for {
		_dcg, _bdec := _afgdc._cac.Peek(1)
		if _bdec != nil {
			return 0, _bdec
		}
		if _ed.IsWhiteSpace(_dcg[0]) {
			_afgdc._cac.ReadByte()
			_dga++
		} else {
			break
		}
	}
	return _dga, nil
}

// ContentStreamParser represents a content stream parser for parsing content streams in PDFs.
type ContentStreamParser struct{ _cac *_g.Reader }

func _egb(_cdca *ContentStreamInlineImage) (_ed.StreamEncoder, error) {
	if _cdca.Filter == nil {
		return _ed.NewRawEncoder(), nil
	}
	_fcf, _edf := _cdca.Filter.(*_ed.PdfObjectName)
	if !_edf {
		_adgd, _egdb := _cdca.Filter.(*_ed.PdfObjectArray)
		if !_egdb {
			return nil, _fe.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0072 \u0041\u0072\u0072\u0061\u0079\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
		if _adgd.Len() == 0 {
			return _ed.NewRawEncoder(), nil
		}
		if _adgd.Len() != 1 {
			_gbae, _addd := _gaea(_cdca)
			if _addd != nil {
				_fb.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0063\u0072\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u006d\u0075\u006c\u0074i\u0020\u0065\u006e\u0063\u006f\u0064\u0065r\u003a\u0020\u0025\u0076", _addd)
				return nil, _addd
			}
			_fb.Log.Trace("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063:\u0020\u0025\u0073\u000a", _gbae)
			return _gbae, nil
		}
		_bcf := _adgd.Get(0)
		_fcf, _egdb = _bcf.(*_ed.PdfObjectName)
		if !_egdb {
			return nil, _fe.Errorf("\u0066\u0069l\u0074\u0065\u0072\u0020a\u0072\u0072a\u0079\u0020\u006d\u0065\u006d\u0062\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
	}
	switch *_fcf {
	case "\u0041\u0048\u0078", "\u0041\u0053\u0043\u0049\u0049\u0048\u0065\u0078\u0044e\u0063\u006f\u0064\u0065":
		return _ed.NewASCIIHexEncoder(), nil
	case "\u0041\u0038\u0035", "\u0041\u0053\u0043\u0049\u0049\u0038\u0035\u0044\u0065\u0063\u006f\u0064\u0065":
		return _ed.NewASCII85Encoder(), nil
	case "\u0044\u0043\u0054", "\u0044C\u0054\u0044\u0065\u0063\u006f\u0064e":
		return _dgcf(_cdca)
	case "\u0046\u006c", "F\u006c\u0061\u0074\u0065\u0044\u0065\u0063\u006f\u0064\u0065":
		return _bffg(_cdca, nil)
	case "\u004c\u005a\u0057", "\u004cZ\u0057\u0044\u0065\u0063\u006f\u0064e":
		return _ecfd(_cdca, nil)
	case "\u0043\u0043\u0046", "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065":
		return _ed.NewCCITTFaxEncoder(), nil
	case "\u0052\u004c", "\u0052u\u006eL\u0065\u006e\u0067\u0074\u0068\u0044\u0065\u0063\u006f\u0064\u0065":
		return _ed.NewRunLengthEncoder(), nil
	default:
		_fb.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065 \u0069\u006d\u0061\u0067\u0065\u0020\u0065n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u003a\u0020\u0025\u0073", *_fcf)
		return nil, _b.New("\u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006el\u0069n\u0065 \u0065n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006f\u0064")
	}
}

// Add_re appends 're' operand to the content stream:
// Append a rectangle to the current path as a complete subpath, with lower left corner (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_cda *ContentCreator) Add_re(x, y, width, height float64) *ContentCreator {
	_bge := ContentStreamOperation{}
	_bge.Operand = "\u0072\u0065"
	_bge.Params = _gfbfb([]float64{x, y, width, height})
	_cda._cg = append(_cda._cg, &_bge)
	return _cda
}

// Add_S appends 'S' operand to the content stream: Stroke the path.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_dff *ContentCreator) Add_S() *ContentCreator {
	_fad := ContentStreamOperation{}
	_fad.Operand = "\u0053"
	_dff._cg = append(_dff._cg, &_fad)
	return _dff
}

// Add_Q adds 'Q' operand to the content stream: Pops the most recently stored state from the stack.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_ceb *ContentCreator) Add_Q() *ContentCreator {
	_gfgg := ContentStreamOperation{}
	_gfgg.Operand = "\u0051"
	_ceb._cg = append(_ceb._cg, &_gfgg)
	return _ceb
}
