package contentstream

import (
	_b "bufio"
	_de "bytes"
	_g "encoding/hex"
	_a "errors"
	_d "fmt"
	_be "image/color"
	_ag "image/jpeg"
	_e "io"
	_ed "math"
	_cb "strconv"

	_gc "bitbucket.org/shenghui0779/gopdf/common"
	_gb "bitbucket.org/shenghui0779/gopdf/core"
	_bef "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_df "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_ef "bitbucket.org/shenghui0779/gopdf/model"
)

// SetNonStrokingColor sets the non-stroking `color` where color can be one of
// PdfColorDeviceGray, PdfColorDeviceRGB, or PdfColorDeviceCMYK.
func (_fee *ContentCreator) SetNonStrokingColor(color _ef.PdfColor) *ContentCreator {
	switch _dbb := color.(type) {
	case *_ef.PdfColorDeviceGray:
		_fee.Add_g(_dbb.Val())
	case *_ef.PdfColorDeviceRGB:
		_fee.Add_rg(_dbb.R(), _dbb.G(), _dbb.B())
	case *_ef.PdfColorDeviceCMYK:
		_fee.Add_k(_dbb.C(), _dbb.M(), _dbb.Y(), _dbb.K())
	default:
		_gc.Log.Debug("\u0053\u0065\u0074N\u006f\u006e\u0053\u0074\u0072\u006f\u006b\u0069\u006e\u0067\u0043\u006f\u006c\u006f\u0072\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020c\u006f\u006c\u006f\u0072\u003a\u0020\u0025\u0054", _dbb)
	}
	return _fee
}
func _fddd(_abf *ContentStreamInlineImage) (*_gb.DCTEncoder, error) {
	_ffcf := _gb.NewDCTEncoder()
	_gfa := _de.NewReader(_abf._cbac)
	_cac, _ega := _ag.DecodeConfig(_gfa)
	if _ega != nil {
		_gc.Log.Debug("\u0045\u0072\u0072or\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u003a\u0020\u0025\u0073", _ega)
		return nil, _ega
	}
	switch _cac.ColorModel {
	case _be.RGBAModel:
		_ffcf.BitsPerComponent = 8
		_ffcf.ColorComponents = 3
	case _be.RGBA64Model:
		_ffcf.BitsPerComponent = 16
		_ffcf.ColorComponents = 3
	case _be.GrayModel:
		_ffcf.BitsPerComponent = 8
		_ffcf.ColorComponents = 1
	case _be.Gray16Model:
		_ffcf.BitsPerComponent = 16
		_ffcf.ColorComponents = 1
	case _be.CMYKModel:
		_ffcf.BitsPerComponent = 8
		_ffcf.ColorComponents = 4
	case _be.YCbCrModel:
		_ffcf.BitsPerComponent = 8
		_ffcf.ColorComponents = 3
	default:
		return nil, _a.New("\u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u006d\u006f\u0064\u0065\u006c")
	}
	_ffcf.Width = _cac.Width
	_ffcf.Height = _cac.Height
	_gc.Log.Trace("\u0044\u0043T\u0020\u0045\u006ec\u006f\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076", _ffcf)
	return _ffcf, nil
}

// Parse parses all commands in content stream, returning a list of operation data.
func (_fdb *ContentStreamParser) Parse() (*ContentStreamOperations, error) {
	_fdc := ContentStreamOperations{}
	for {
		_eeg := ContentStreamOperation{}
		for {
			_dfd, _dff, _dfbb := _fdb.parseObject()
			if _dfbb != nil {
				if _dfbb == _e.EOF {
					return &_fdc, nil
				}
				return &_fdc, _dfbb
			}
			if _dff {
				_eeg.Operand, _ = _gb.GetStringVal(_dfd)
				_fdc = append(_fdc, &_eeg)
				break
			} else {
				_eeg.Params = append(_eeg.Params, _dfd)
			}
		}
		if _eeg.Operand == "\u0042\u0049" {
			_bbe, _ddd := _fdb.ParseInlineImage()
			if _ddd != nil {
				return &_fdc, _ddd
			}
			_eeg.Params = append(_eeg.Params, _bbe)
		}
	}
}
func _cfd(_bfed _ef.PdfColorspace) bool {
	_, _edfd := _bfed.(*_ef.PdfColorspaceSpecialPattern)
	return _edfd
}

// Push pushes `gs` on the `gsStack`.
func (_ced *GraphicStateStack) Push(gs GraphicsState) { *_ced = append(*_ced, gs) }

// Add_c adds 'c' operand to the content stream: Append a Bezier curve to the current path from
// the current point to (x3,y3) with (x1,x1) and (x2,y2) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_dd *ContentCreator) Add_c(x1, y1, x2, y2, x3, y3 float64) *ContentCreator {
	_deg := ContentStreamOperation{}
	_deg.Operand = "\u0063"
	_deg.Params = _bcd([]float64{x1, y1, x2, y2, x3, y3})
	_dd._ec = append(_dd._ec, &_deg)
	return _dd
}

// Add_l adds 'l' operand to the content stream:
// Append a straight line segment from the current point to (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_bdfc *ContentCreator) Add_l(x, y float64) *ContentCreator {
	_abb := ContentStreamOperation{}
	_abb.Operand = "\u006c"
	_abb.Params = _bcd([]float64{x, y})
	_bdfc._ec = append(_bdfc._ec, &_abb)
	return _bdfc
}

// SetStrokingColor sets the stroking `color` where color can be one of
// PdfColorDeviceGray, PdfColorDeviceRGB, or PdfColorDeviceCMYK.
func (_bde *ContentCreator) SetStrokingColor(color _ef.PdfColor) *ContentCreator {
	switch _cbfbc := color.(type) {
	case *_ef.PdfColorDeviceGray:
		_bde.Add_G(_cbfbc.Val())
	case *_ef.PdfColorDeviceRGB:
		_bde.Add_RG(_cbfbc.R(), _cbfbc.G(), _cbfbc.B())
	case *_ef.PdfColorDeviceCMYK:
		_bde.Add_K(_cbfbc.C(), _cbfbc.M(), _cbfbc.Y(), _cbfbc.K())
	default:
		_gc.Log.Debug("\u0053\u0065\u0074\u0053\u0074\u0072\u006f\u006b\u0069\u006e\u0067\u0043\u006fl\u006f\u0072\u003a\u0020\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006fr\u003a\u0020\u0025\u0054", _cbfbc)
	}
	return _bde
}

// Add_b appends 'b' operand to the content stream:
// Close, fill and then stroke the path (nonzero winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_bed *ContentCreator) Add_b() *ContentCreator {
	_gab := ContentStreamOperation{}
	_gab.Operand = "\u0062"
	_bed._ec = append(_bed._ec, &_gab)
	return _bed
}

// Add_b_starred appends 'b*' operand to the content stream:
// Close, fill and then stroke the path (even-odd winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_fgb *ContentCreator) Add_b_starred() *ContentCreator {
	_fb := ContentStreamOperation{}
	_fb.Operand = "\u0062\u002a"
	_fgb._ec = append(_fgb._ec, &_fb)
	return _fgb
}

// Add_i adds 'i' operand to the content stream: Set the flatness tolerance in the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_bdf *ContentCreator) Add_i(flatness float64) *ContentCreator {
	_affg := ContentStreamOperation{}
	_affg.Operand = "\u0069"
	_affg.Params = _bcd([]float64{flatness})
	_bdf._ec = append(_bdf._ec, &_affg)
	return _bdf
}

// HandlerFunc is the function syntax that the ContentStreamProcessor handler must implement.
type HandlerFunc func(_cgeb *ContentStreamOperation, _bdeb GraphicsState, _fegd *_ef.PdfPageResources) error

// Add_cs appends 'cs' operand to the content stream:
// Same as CS but for non-stroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_fc *ContentCreator) Add_cs(name _gb.PdfObjectName) *ContentCreator {
	_cda := ContentStreamOperation{}
	_cda.Operand = "\u0063\u0073"
	_cda.Params = _afec([]_gb.PdfObjectName{name})
	_fc._ec = append(_fc._ec, &_cda)
	return _fc
}
func (_dddb *ContentStreamProcessor) handleCommand_G(_fgbf *ContentStreamOperation, _abfe *_ef.PdfPageResources) error {
	_fbbe := _ef.NewPdfColorspaceDeviceGray()
	if len(_fgbf.Params) != _fbbe.GetNumComponents() {
		_gc.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_gc.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_fgbf.Params), _fbbe)
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_dcgc, _aaag := _fbbe.ColorFromPdfObjects(_fgbf.Params)
	if _aaag != nil {
		return _aaag
	}
	_dddb._efaee.ColorspaceStroking = _fbbe
	_dddb._efaee.ColorStroking = _dcgc
	return nil
}
func (_bb *ContentStreamOperations) isWrapped() bool {
	if len(*_bb) < 2 {
		return false
	}
	_f := 0
	for _, _agf := range *_bb {
		if _agf.Operand == "\u0071" {
			_f++
		} else if _agf.Operand == "\u0051" {
			_f--
		} else {
			if _f < 1 {
				return false
			}
		}
	}
	return _f == 0
}
func _caff(_efeg *ContentStreamInlineImage, _eb *_gb.PdfObjectDictionary) (*_gb.LZWEncoder, error) {
	_baa := _gb.NewLZWEncoder()
	if _eb == nil {
		if _efeg.DecodeParms != nil {
			_aeea, _edg := _gb.GetDict(_efeg.DecodeParms)
			if !_edg {
				_gc.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079 \u0028%\u0054\u0029", _efeg.DecodeParms)
				return nil, _d.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
			_eb = _aeea
		}
	}
	if _eb == nil {
		return _baa, nil
	}
	_ged := _eb.Get("E\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
	if _ged != nil {
		_fgde, _ffc := _ged.(*_gb.PdfObjectInteger)
		if !_ffc {
			_gc.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a \u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u006e\u0075\u006d\u0065\u0072i\u0063 \u0028\u0025\u0054\u0029", _ged)
			return nil, _d.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065")
		}
		if *_fgde != 0 && *_fgde != 1 {
			return nil, _d.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0045\u0061\u0072\u006c\u0079\u0043\u0068\u0061\u006e\u0067\u0065\u0020\u0076\u0061\u006c\u0075e\u0020\u0028\u006e\u006f\u0074 \u0030\u0020o\u0072\u0020\u0031\u0029")
		}
		_baa.EarlyChange = int(*_fgde)
	} else {
		_baa.EarlyChange = 1
	}
	_ged = _eb.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _ged != nil {
		_eaf, _dbe := _ged.(*_gb.PdfObjectInteger)
		if !_dbe {
			_gc.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _ged)
			return nil, _d.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_baa.Predictor = int(*_eaf)
	}
	_ged = _eb.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _ged != nil {
		_def, _gdg := _ged.(*_gb.PdfObjectInteger)
		if !_gdg {
			_gc.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _d.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_baa.BitsPerComponent = int(*_def)
	}
	if _baa.Predictor > 1 {
		_baa.Columns = 1
		_ged = _eb.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _ged != nil {
			_fdgc, _cdf := _ged.(*_gb.PdfObjectInteger)
			if !_cdf {
				return nil, _d.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_baa.Columns = int(*_fdgc)
		}
		_baa.Colors = 1
		_ged = _eb.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _ged != nil {
			_bgfa, _aeba := _ged.(*_gb.PdfObjectInteger)
			if !_aeba {
				return nil, _d.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_baa.Colors = int(*_bgfa)
		}
	}
	_gc.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _eb.String())
	return _baa, nil
}

// ParseInlineImage parses an inline image from a content stream, both reading its properties and binary data.
// When called, "BI" has already been read from the stream.  This function
// finishes reading through "EI" and then returns the ContentStreamInlineImage.
func (_bbac *ContentStreamParser) ParseInlineImage() (*ContentStreamInlineImage, error) {
	_ecbe := ContentStreamInlineImage{}
	for {
		_bbac.skipSpaces()
		_dead, _eacb, _aeeg := _bbac.parseObject()
		if _aeeg != nil {
			return nil, _aeeg
		}
		if !_eacb {
			_cgb, _dbd := _gb.GetName(_dead)
			if !_dbd {
				_gc.Log.Debug("\u0049\u006e\u0076\u0061\u006ci\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067e\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0020\u0028\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u0029\u0020\u002d\u0020\u0025T", _dead)
				return nil, _d.Errorf("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067e\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0020\u0028\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067\u0020\u006e\u0061\u006d\u0065\u0029\u0020\u002d\u0020\u0025T", _dead)
			}
			_ddb, _ecf, _aafc := _bbac.parseObject()
			if _aafc != nil {
				return nil, _aafc
			}
			if _ecf {
				return nil, _d.Errorf("\u006eo\u0074\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0069\u006e\u0067 \u0061\u006e\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
			}
			switch *_cgb {
			case "\u0042\u0050\u0043", "\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074":
				_ecbe.BitsPerComponent = _ddb
			case "\u0043\u0053", "\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065":
				_ecbe.ColorSpace = _ddb
			case "\u0044", "\u0044\u0065\u0063\u006f\u0064\u0065":
				_ecbe.Decode = _ddb
			case "\u0044\u0050", "D\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073":
				_ecbe.DecodeParms = _ddb
			case "\u0046", "\u0046\u0069\u006c\u0074\u0065\u0072":
				_ecbe.Filter = _ddb
			case "\u0048", "\u0048\u0065\u0069\u0067\u0068\u0074":
				_ecbe.Height = _ddb
			case "\u0049\u004d", "\u0049m\u0061\u0067\u0065\u004d\u0061\u0073k":
				_ecbe.ImageMask = _ddb
			case "\u0049\u006e\u0074\u0065\u006e\u0074":
				_ecbe.Intent = _ddb
			case "\u0049", "I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065":
				_ecbe.Interpolate = _ddb
			case "\u0057", "\u0057\u0069\u0064t\u0068":
				_ecbe.Width = _ddb
			case "\u004c\u0065\u006e\u0067\u0074\u0068", "\u0053u\u0062\u0074\u0079\u0070\u0065", "\u0054\u0079\u0070\u0065":
				_gc.Log.Debug("\u0049\u0067\u006e\u006fr\u0069\u006e\u0067\u0020\u0069\u006e\u006c\u0069\u006e\u0065 \u0070a\u0072\u0061\u006d\u0065\u0074\u0065\u0072 \u0025\u0073", *_cgb)
			default:
				return nil, _d.Errorf("\u0075\u006e\u006b\u006e\u006f\u0077n\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0020\u0025\u0073", *_cgb)
			}
		}
		if _eacb {
			_dfb, _efd := _dead.(*_gb.PdfObjectString)
			if !_efd {
				return nil, _d.Errorf("\u0066a\u0069\u006ce\u0064\u0020\u0074o\u0020\u0072\u0065\u0061\u0064\u0020\u0069n\u006c\u0069\u006e\u0065\u0020\u0069m\u0061\u0067\u0065\u0020\u002d\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
			}
			if _dfb.Str() == "\u0045\u0049" {
				_gc.Log.Trace("\u0049n\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020f\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e\u002e\u002e")
				return &_ecbe, nil
			} else if _dfb.Str() == "\u0049\u0044" {
				_gc.Log.Trace("\u0049\u0044\u0020\u0073\u0074\u0061\u0072\u0074")
				_geef, _ccef := _bbac._dcf.Peek(1)
				if _ccef != nil {
					return nil, _ccef
				}
				if _gb.IsWhiteSpace(_geef[0]) {
					_bbac._dcf.Discard(1)
				}
				_ecbe._cbac = []byte{}
				_bfa := 0
				var _agbc []byte
				for {
					_cefd, _cfag := _bbac._dcf.ReadByte()
					if _cfag != nil {
						_gc.Log.Debug("\u0055\u006e\u0061\u0062\u006ce\u0020\u0074\u006f\u0020\u0066\u0069\u006e\u0064\u0020\u0065\u006e\u0064\u0020o\u0066\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0045\u0049\u0020\u0069\u006e\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u0061\u0074a")
						return nil, _cfag
					}
					if _bfa == 0 {
						if _gb.IsWhiteSpace(_cefd) {
							_agbc = []byte{}
							_agbc = append(_agbc, _cefd)
							_bfa = 1
						} else if _cefd == 'E' {
							_agbc = append(_agbc, _cefd)
							_bfa = 2
						} else {
							_ecbe._cbac = append(_ecbe._cbac, _cefd)
						}
					} else if _bfa == 1 {
						_agbc = append(_agbc, _cefd)
						if _cefd == 'E' {
							_bfa = 2
						} else {
							_ecbe._cbac = append(_ecbe._cbac, _agbc...)
							_agbc = []byte{}
							if _gb.IsWhiteSpace(_cefd) {
								_bfa = 1
							} else {
								_bfa = 0
							}
						}
					} else if _bfa == 2 {
						_agbc = append(_agbc, _cefd)
						if _cefd == 'I' {
							_bfa = 3
						} else {
							_ecbe._cbac = append(_ecbe._cbac, _agbc...)
							_agbc = []byte{}
							_bfa = 0
						}
					} else if _bfa == 3 {
						_agbc = append(_agbc, _cefd)
						if _gb.IsWhiteSpace(_cefd) {
							_bfgc, _cdb := _bbac._dcf.Peek(20)
							if _cdb != nil && _cdb != _e.EOF {
								return nil, _cdb
							}
							_cdda := NewContentStreamParser(string(_bfgc))
							_baaa := true
							for _aecg := 0; _aecg < 3; _aecg++ {
								_dcaa, _dfbf, _acae := _cdda.parseObject()
								if _acae != nil {
									if _acae == _e.EOF {
										break
									}
									_baaa = false
									continue
								}
								if _dfbf && !_ebg(_dcaa.String()) {
									_baaa = false
									break
								}
							}
							if _baaa {
								if len(_ecbe._cbac) > 100 {
									_gc.Log.Trace("\u0049\u006d\u0061\u0067\u0065\u0020\u0073\u0074\u0072\u0065\u0061m\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078 \u002e\u002e\u002e", len(_ecbe._cbac), _ecbe._cbac[:100])
								} else {
									_gc.Log.Trace("\u0049\u006d\u0061\u0067e \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025 \u0078", len(_ecbe._cbac), _ecbe._cbac)
								}
								return &_ecbe, nil
							}
						}
						_ecbe._cbac = append(_ecbe._cbac, _agbc...)
						_agbc = []byte{}
						_bfa = 0
					}
				}
			}
		}
	}
}

// Scale applies x-y scaling to the transformation matrix.
func (_db *ContentCreator) Scale(sx, sy float64) *ContentCreator {
	return _db.Add_cm(sx, 0, 0, sy, 0, 0)
}
func (_cccb *ContentStreamProcessor) handleCommand_SC(_afc *ContentStreamOperation, _fdge *_ef.PdfPageResources) error {
	_dfba := _cccb._efaee.ColorspaceStroking
	if len(_afc.Params) != _dfba.GetNumComponents() {
		_gc.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_gc.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_afc.Params), _dfba)
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_bcf, _fbbg := _dfba.ColorFromPdfObjects(_afc.Params)
	if _fbbg != nil {
		return _fbbg
	}
	_cccb._efaee.ColorStroking = _bcf
	return nil
}

// Add_W_starred appends 'W*' operand to the content stream:
// Modify the current clipping path by intersecting with the current path (even odd rule).
//
// See section 8.5.4 "Clipping Path Operators" and Table 61 (p. 146 PDF32000_2008).
func (_bfd *ContentCreator) Add_W_starred() *ContentCreator {
	_bea := ContentStreamOperation{}
	_bea.Operand = "\u0057\u002a"
	_bfd._ec = append(_bfd._ec, &_bea)
	return _bfd
}

// Add_f_starred appends 'f*' operand to the content stream.
// f*: Fill the path using the even-odd rule to determine fill region.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_ceg *ContentCreator) Add_f_starred() *ContentCreator {
	_fed := ContentStreamOperation{}
	_fed.Operand = "\u0066\u002a"
	_ceg._ec = append(_ceg._ec, &_fed)
	return _ceg
}

// Bytes converts the content stream operations to a content stream byte presentation, i.e. the kind that can be
// stored as a PDF stream or string format.
func (_af *ContentCreator) Bytes() []byte { return _af._ec.Bytes() }

// Add_Do adds 'Do' operation to the content stream:
// Displays an XObject (image or form) specified by `name`.
//
// See section 8.8 "External Objects" and Table 87 (pp. 209-220 PDF32000_2008).
func (_fe *ContentCreator) Add_Do(name _gb.PdfObjectName) *ContentCreator {
	_bcb := ContentStreamOperation{}
	_bcb.Operand = "\u0044\u006f"
	_bcb.Params = _afec([]_gb.PdfObjectName{name})
	_fe._ec = append(_fe._ec, &_bcb)
	return _fe
}

// ContentStreamInlineImage is a representation of an inline image in a Content stream. Everything between the BI and EI operands.
// ContentStreamInlineImage implements the core.PdfObject interface although strictly it is not a PDF object.
type ContentStreamInlineImage struct {
	BitsPerComponent _gb.PdfObject
	ColorSpace       _gb.PdfObject
	Decode           _gb.PdfObject
	DecodeParms      _gb.PdfObject
	Filter           _gb.PdfObject
	Height           _gb.PdfObject
	ImageMask        _gb.PdfObject
	Intent           _gb.PdfObject
	Interpolate      _gb.PdfObject
	Width            _gb.PdfObject
	_cbac            []byte
	_adfg            *_bef.ImageBase
}

// Add_j adds 'j' operand to the content stream: Set the line join style (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_cc *ContentCreator) Add_j(lineJoinStyle string) *ContentCreator {
	_ge := ContentStreamOperation{}
	_ge.Operand = "\u006a"
	_ge.Params = _afec([]_gb.PdfObjectName{_gb.PdfObjectName(lineJoinStyle)})
	_cc._ec = append(_cc._ec, &_ge)
	return _cc
}
func (_cfa *ContentStreamInlineImage) toImageBase(_cfe *_ef.PdfPageResources) (*_bef.ImageBase, error) {
	if _cfa._adfg != nil {
		return _cfa._adfg, nil
	}
	_aedc := _bef.ImageBase{}
	if _cfa.Height == nil {
		return nil, _a.New("\u0068e\u0069\u0067\u0068\u0074\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_fba, _fce := _cfa.Height.(*_gb.PdfObjectInteger)
	if !_fce {
		return nil, _a.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0068e\u0069\u0067\u0068\u0074")
	}
	_aedc.Height = int(*_fba)
	if _cfa.Width == nil {
		return nil, _a.New("\u0077\u0069\u0064th\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067")
	}
	_afg, _fce := _cfa.Width.(*_gb.PdfObjectInteger)
	if !_fce {
		return nil, _a.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0069\u0064\u0074\u0068")
	}
	_aedc.Width = int(*_afg)
	_gabf, _gabe := _cfa.IsMask()
	if _gabe != nil {
		return nil, _gabe
	}
	if _gabf {
		_aedc.BitsPerComponent = 1
		_aedc.ColorComponents = 1
	} else {
		if _cfa.BitsPerComponent == nil {
			_gc.Log.Debug("\u0049\u006el\u0069\u006e\u0065\u0020\u0042\u0069\u0074\u0073\u0020\u0070\u0065\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u002d\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u0038")
			_aedc.BitsPerComponent = 8
		} else {
			_gbd, _feec := _cfa.BitsPerComponent.(*_gb.PdfObjectInteger)
			if !_feec {
				_gc.Log.Debug("E\u0072\u0072\u006f\u0072\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0062\u0069\u0074\u0073 p\u0065\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0020\u0076al\u0075\u0065,\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _cfa.BitsPerComponent)
				return nil, _a.New("\u0042\u0050\u0043\u0020\u0054\u0079\u0070\u0065\u0020e\u0072\u0072\u006f\u0072")
			}
			_aedc.BitsPerComponent = int(*_gbd)
		}
		if _cfa.ColorSpace != nil {
			_geee, _gfef := _cfa.GetColorSpace(_cfe)
			if _gfef != nil {
				return nil, _gfef
			}
			_aedc.ColorComponents = _geee.GetNumComponents()
		} else {
			_gc.Log.Debug("\u0049\u006el\u0069\u006e\u0065\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063e\u0020\u006e\u006f\u0074\u0020\u0073p\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u002d\u0020\u0061\u0073\u0073\u0075m\u0069\u006eg\u0020\u0031\u0020\u0063o\u006c\u006f\u0072\u0020\u0063o\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			_aedc.ColorComponents = 1
		}
	}
	if _aebc, _bgfaa := _gb.GetArray(_cfa.Decode); _bgfaa {
		_aedc.Decode, _gabe = _aebc.ToFloat64Array()
		if _gabe != nil {
			return nil, _gabe
		}
	}
	_cfa._adfg = &_aedc
	return _cfa._adfg, nil
}

// Add_gs adds 'gs' operand to the content stream: Set the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_dab *ContentCreator) Add_gs(dictName _gb.PdfObjectName) *ContentCreator {
	_bcg := ContentStreamOperation{}
	_bcg.Operand = "\u0067\u0073"
	_bcg.Params = _afec([]_gb.PdfObjectName{dictName})
	_dab._ec = append(_dab._ec, &_bcg)
	return _dab
}

// Add_Tw appends 'Tw' operand to the content stream:
// Set word spacing.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_bce *ContentCreator) Add_Tw(wordSpace float64) *ContentCreator {
	_bgcd := ContentStreamOperation{}
	_bgcd.Operand = "\u0054\u0077"
	_bgcd.Params = _bcd([]float64{wordSpace})
	_bce._ec = append(_bce._ec, &_bgcd)
	return _bce
}

// Add_ri adds 'ri' operand to the content stream, which sets the color rendering intent.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_eac *ContentCreator) Add_ri(intent _gb.PdfObjectName) *ContentCreator {
	_fdd := ContentStreamOperation{}
	_fdd.Operand = "\u0072\u0069"
	_fdd.Params = _afec([]_gb.PdfObjectName{intent})
	_eac._ec = append(_eac._ec, &_fdd)
	return _eac
}

// Add_y appends 'y' operand to the content stream: Append a Bezier curve to the current path from the
// current point to (x3,y3) with (x1, y1) and (x3,y3) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_afb *ContentCreator) Add_y(x1, y1, x3, y3 float64) *ContentCreator {
	_fab := ContentStreamOperation{}
	_fab.Operand = "\u0079"
	_fab.Params = _bcd([]float64{x1, y1, x3, y3})
	_afb._ec = append(_afb._ec, &_fab)
	return _afb
}
func _afeb(_cbba []int64) []_gb.PdfObject {
	var _bfga []_gb.PdfObject
	for _, _fbce := range _cbba {
		_bfga = append(_bfga, _gb.MakeInteger(_fbce))
	}
	return _bfga
}

// Add_J adds 'J' operand to the content stream: Set the line cap style (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_gcc *ContentCreator) Add_J(lineCapStyle string) *ContentCreator {
	_fde := ContentStreamOperation{}
	_fde.Operand = "\u004a"
	_fde.Params = _afec([]_gb.PdfObjectName{_gb.PdfObjectName(lineCapStyle)})
	_gcc._ec = append(_gcc._ec, &_fde)
	return _gcc
}
func (_eceab *ContentStreamProcessor) handleCommand_K(_caaf *ContentStreamOperation, _afae *_ef.PdfPageResources) error {
	_bdg := _ef.NewPdfColorspaceDeviceCMYK()
	if len(_caaf.Params) != _bdg.GetNumComponents() {
		_gc.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_gc.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_caaf.Params), _bdg)
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_cegd, _fccg := _bdg.ColorFromPdfObjects(_caaf.Params)
	if _fccg != nil {
		return _fccg
	}
	_eceab._efaee.ColorspaceStroking = _bdg
	_eceab._efaee.ColorStroking = _cegd
	return nil
}

// Add_sh appends 'sh' operand to the content stream:
// Paints the shape and colour shading described by a shading dictionary specified by `name`,
// subject to the current clipping path
//
// See section 8.7.4 "Shading Patterns" and Table 77 (p. 190 PDF32000_2008).
func (_abg *ContentCreator) Add_sh(name _gb.PdfObjectName) *ContentCreator {
	_bgg := ContentStreamOperation{}
	_bgg.Operand = "\u0073\u0068"
	_bgg.Params = _afec([]_gb.PdfObjectName{name})
	_abg._ec = append(_abg._ec, &_bgg)
	return _abg
}
func (_dcdc *ContentStreamParser) parseOperand() (*_gb.PdfObjectString, error) {
	var _gba []byte
	for {
		_bfdg, _cgf := _dcdc._dcf.Peek(1)
		if _cgf != nil {
			return _gb.MakeString(string(_gba)), _cgf
		}
		if _gb.IsDelimiter(_bfdg[0]) {
			break
		}
		if _gb.IsWhiteSpace(_bfdg[0]) {
			break
		}
		_fceb, _ := _dcdc._dcf.ReadByte()
		_gba = append(_gba, _fceb)
	}
	return _gb.MakeString(string(_gba)), nil
}

// Add_Ts appends 'Ts' operand to the content stream:
// Set text rise.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_edee *ContentCreator) Add_Ts(rise float64) *ContentCreator {
	_gae := ContentStreamOperation{}
	_gae.Operand = "\u0054\u0073"
	_gae.Params = _bcd([]float64{rise})
	_edee._ec = append(_edee._ec, &_gae)
	return _edee
}

// Add_g appends 'g' operand to the content stream:
// Same as G but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_bff *ContentCreator) Add_g(gray float64) *ContentCreator {
	_dag := ContentStreamOperation{}
	_dag.Operand = "\u0067"
	_dag.Params = _bcd([]float64{gray})
	_bff._ec = append(_bff._ec, &_dag)
	return _bff
}

// Add_RG appends 'RG' operand to the content stream:
// Set the stroking colorspace to DeviceRGB and sets the r,g,b colors (0-1 each).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_deca *ContentCreator) Add_RG(r, g, b float64) *ContentCreator {
	_cdac := ContentStreamOperation{}
	_cdac.Operand = "\u0052\u0047"
	_cdac.Params = _bcd([]float64{r, g, b})
	_deca._ec = append(_deca._ec, &_cdac)
	return _deca
}

// Add_SC appends 'SC' operand to the content stream:
// Set color for stroking operations.  Input: c1, ..., cn.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_bab *ContentCreator) Add_SC(c ...float64) *ContentCreator {
	_bbc := ContentStreamOperation{}
	_bbc.Operand = "\u0053\u0043"
	_bbc.Params = _bcd(c)
	_bab._ec = append(_bab._ec, &_bbc)
	return _bab
}

// Add_CS appends 'CS' operand to the content stream:
// Set the current colorspace for stroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_dgb *ContentCreator) Add_CS(name _gb.PdfObjectName) *ContentCreator {
	_ff := ContentStreamOperation{}
	_ff.Operand = "\u0043\u0053"
	_ff.Params = _afec([]_gb.PdfObjectName{name})
	_dgb._ec = append(_dgb._ec, &_ff)
	return _dgb
}

// Add_m adds 'm' operand to the content stream: Move the current point to (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_cec *ContentCreator) Add_m(x, y float64) *ContentCreator {
	_cbd := ContentStreamOperation{}
	_cbd.Operand = "\u006d"
	_cbd.Params = _bcd([]float64{x, y})
	_cec._ec = append(_cec._ec, &_cbd)
	return _cec
}

// Add_quote appends "'" operand to the content stream:
// Move to next line and show a string.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_ecda *ContentCreator) Add_quote(textstr _gb.PdfObjectString) *ContentCreator {
	_adf := ContentStreamOperation{}
	_adf.Operand = "\u0027"
	_adf.Params = _ccbe([]_gb.PdfObjectString{textstr})
	_ecda._ec = append(_ecda._ec, &_adf)
	return _ecda
}

// Translate applies a simple x-y translation to the transformation matrix.
func (_bga *ContentCreator) Translate(tx, ty float64) *ContentCreator {
	return _bga.Add_cm(1, 0, 0, 1, tx, ty)
}

// Add_scn_pattern appends 'scn' operand to the content stream for pattern `name`:
// scn with name attribute (for pattern). Syntax: c1 ... cn name scn.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_cba *ContentCreator) Add_scn_pattern(name _gb.PdfObjectName, c ...float64) *ContentCreator {
	_bgad := ContentStreamOperation{}
	_bgad.Operand = "\u0073\u0063\u006e"
	_bgad.Params = _bcd(c)
	_bgad.Params = append(_bgad.Params, _gb.MakeName(string(name)))
	_cba._ec = append(_cba._ec, &_bgad)
	return _cba
}

// Add_quotes appends `"` operand to the content stream:
// Move to next line and show a string, using `aw` and `ac` as word
// and character spacing respectively.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_gfd *ContentCreator) Add_quotes(textstr _gb.PdfObjectString, aw, ac float64) *ContentCreator {
	_ade := ContentStreamOperation{}
	_ade.Operand = "\u0022"
	_ade.Params = _bcd([]float64{aw, ac})
	_ade.Params = append(_ade.Params, _ccbe([]_gb.PdfObjectString{textstr})...)
	_gfd._ec = append(_gfd._ec, &_ade)
	return _gfd
}

// Add_cm adds 'cm' operation to the content stream: Modifies the current transformation matrix (ctm)
// of the graphics state.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_daa *ContentCreator) Add_cm(a, b, c, d, e, f float64) *ContentCreator {
	_fda := ContentStreamOperation{}
	_fda.Operand = "\u0063\u006d"
	_fda.Params = _bcd([]float64{a, b, c, d, e, f})
	_daa._ec = append(_daa._ec, &_fda)
	return _daa
}

// Add_f appends 'f' operand to the content stream:
// Fill the path using the nonzero winding number rule to determine fill region.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_cf *ContentCreator) Add_f() *ContentCreator {
	_gffe := ContentStreamOperation{}
	_gffe.Operand = "\u0066"
	_cf._ec = append(_cf._ec, &_gffe)
	return _cf
}

// Add_EMC appends 'EMC' operand to the content stream:
// Ends a marked-content sequence.
//
// See section 14.6 "Marked Content" and Table 320 (p. 561 PDF32000_2008).
func (_fddc *ContentCreator) Add_EMC() *ContentCreator {
	_abd := ContentStreamOperation{}
	_abd.Operand = "\u0045\u004d\u0043"
	_fddc._ec = append(_fddc._ec, &_abd)
	return _fddc
}
func _aca(_agea *ContentStreamInlineImage, _dfe *_gb.PdfObjectDictionary) (*_gb.FlateEncoder, error) {
	_edad := _gb.NewFlateEncoder()
	if _agea._adfg != nil {
		_edad.SetImage(_agea._adfg)
	}
	if _dfe == nil {
		_aec := _agea.DecodeParms
		if _aec != nil {
			_fagf, _fdf := _gb.GetDict(_aec)
			if !_fdf {
				_gc.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079 \u0028%\u0054\u0029", _aec)
				return nil, _d.Errorf("\u0069\u006e\u0076\u0061li\u0064\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073")
			}
			_dfe = _fagf
		}
	}
	if _dfe == nil {
		return _edad, nil
	}
	_gc.Log.Trace("\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006ds\u003a\u0020\u0025\u0073", _dfe.String())
	_ggd := _dfe.Get("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr")
	if _ggd == nil {
		_gc.Log.Debug("E\u0072\u0072o\u0072\u003a\u0020\u0050\u0072\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067 \u0066\u0072\u006f\u006d\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u006d\u0073 \u002d\u0020\u0043\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020\u0077\u0069t\u0068\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u00281\u0029")
	} else {
		_caa, _aabdd := _ggd.(*_gb.PdfObjectInteger)
		if !_aabdd {
			_gc.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0050\u0072\u0065d\u0069\u0063\u0074\u006f\u0072\u0020\u0073pe\u0063\u0069\u0066\u0069e\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u006f\u0074 n\u0075\u006de\u0072\u0069\u0063\u0020\u0028\u0025\u0054\u0029", _ggd)
			return nil, _d.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0072\u0065\u0064i\u0063\u0074\u006f\u0072")
		}
		_edad.Predictor = int(*_caa)
	}
	_ggd = _dfe.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
	if _ggd != nil {
		_cdde, _agd := _ggd.(*_gb.PdfObjectInteger)
		if !_agd {
			_gc.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064\u0020\u0042i\u0074\u0073\u0050\u0065\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
			return nil, _d.Errorf("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0042\u0069\u0074\u0073\u0050e\u0072\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")
		}
		_edad.BitsPerComponent = int(*_cdde)
	}
	if _edad.Predictor > 1 {
		_edad.Columns = 1
		_ggd = _dfe.Get("\u0043o\u006c\u0075\u006d\u006e\u0073")
		if _ggd != nil {
			_eae, _aaa := _ggd.(*_gb.PdfObjectInteger)
			if !_aaa {
				return nil, _d.Errorf("\u0070r\u0065\u0064\u0069\u0063\u0074\u006f\u0072\u0020\u0063\u006f\u006cu\u006d\u006e\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
			}
			_edad.Columns = int(*_eae)
		}
		_edad.Colors = 1
		_bae := _dfe.Get("\u0043\u006f\u006c\u006f\u0072\u0073")
		if _bae != nil {
			_gffg, _dbba := _bae.(*_gb.PdfObjectInteger)
			if !_dbba {
				return nil, _d.Errorf("\u0070\u0072\u0065d\u0069\u0063\u0074\u006fr\u0020\u0063\u006f\u006c\u006f\u0072\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
			}
			_edad.Colors = int(*_gffg)
		}
	}
	return _edad, nil
}
func (_eaec *ContentStreamParser) parseArray() (*_gb.PdfObjectArray, error) {
	_bdc := _gb.MakeArray()
	_eaec._dcf.ReadByte()
	for {
		_eaec.skipSpaces()
		_bgaa, _fbfd := _eaec._dcf.Peek(1)
		if _fbfd != nil {
			return _bdc, _fbfd
		}
		if _bgaa[0] == ']' {
			_eaec._dcf.ReadByte()
			break
		}
		_bbea, _, _fbfd := _eaec.parseObject()
		if _fbfd != nil {
			return _bdc, _fbfd
		}
		_bdc.Append(_bbea)
	}
	return _bdc, nil
}

// Add_d adds 'd' operand to the content stream: Set the line dash pattern.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_ae *ContentCreator) Add_d(dashArray []int64, dashPhase int64) *ContentCreator {
	_gee := ContentStreamOperation{}
	_gee.Operand = "\u0064"
	_gee.Params = []_gb.PdfObject{}
	_gee.Params = append(_gee.Params, _gb.MakeArrayFromIntegers64(dashArray))
	_gee.Params = append(_gee.Params, _gb.MakeInteger(dashPhase))
	_ae._ec = append(_ae._ec, &_gee)
	return _ae
}

// Add_SCN_pattern appends 'SCN' operand to the content stream for pattern `name`:
// SCN with name attribute (for pattern). Syntax: c1 ... cn name SCN.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_ceff *ContentCreator) Add_SCN_pattern(name _gb.PdfObjectName, c ...float64) *ContentCreator {
	_bbbf := ContentStreamOperation{}
	_bbbf.Operand = "\u0053\u0043\u004e"
	_bbbf.Params = _bcd(c)
	_bbbf.Params = append(_bbbf.Params, _gb.MakeName(string(name)))
	_ceff._ec = append(_ceff._ec, &_bbbf)
	return _ceff
}

// Add_K appends 'K' operand to the content stream:
// Set the stroking colorspace to DeviceCMYK and sets the c,m,y,k color (0-1 each component).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_acga *ContentCreator) Add_K(c, m, y, k float64) *ContentCreator {
	_bdfaf := ContentStreamOperation{}
	_bdfaf.Operand = "\u004b"
	_bdfaf.Params = _bcd([]float64{c, m, y, k})
	_acga._ec = append(_acga._ec, &_bdfaf)
	return _acga
}
func (_deeae *ContentStreamProcessor) handleCommand_RG(_ffa *ContentStreamOperation, _fdef *_ef.PdfPageResources) error {
	_gccef := _ef.NewPdfColorspaceDeviceRGB()
	if len(_ffa.Params) != _gccef.GetNumComponents() {
		_gc.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020R\u0047")
		_gc.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_ffa.Params), _gccef)
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_bcgg, _deee := _gccef.ColorFromPdfObjects(_ffa.Params)
	if _deee != nil {
		return _deee
	}
	_deeae._efaee.ColorspaceStroking = _gccef
	_deeae._efaee.ColorStroking = _bcgg
	return nil
}
func (_ace *ContentStreamParser) parseDict() (*_gb.PdfObjectDictionary, error) {
	_gc.Log.Trace("\u0052\u0065\u0061\u0064i\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074 \u0073t\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0021")
	_edb := _gb.MakeDict()
	_efaf, _ := _ace._dcf.ReadByte()
	if _efaf != '<' {
		return nil, _a.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	_efaf, _ = _ace._dcf.ReadByte()
	if _efaf != '<' {
		return nil, _a.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074")
	}
	for {
		_ace.skipSpaces()
		_bdca, _edeb := _ace._dcf.Peek(2)
		if _edeb != nil {
			return nil, _edeb
		}
		_gc.Log.Trace("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021", string(_bdca), string(_bdca))
		if (_bdca[0] == '>') && (_bdca[1] == '>') {
			_gc.Log.Trace("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
			_ace._dcf.ReadByte()
			_ace._dcf.ReadByte()
			break
		}
		_gc.Log.Trace("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021")
		_adbb, _edeb := _ace.parseName()
		_gc.Log.Trace("\u004be\u0079\u003a\u0020\u0025\u0073", _adbb)
		if _edeb != nil {
			_gc.Log.Debug("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073", _edeb)
			return nil, _edeb
		}
		if len(_adbb) > 4 && _adbb[len(_adbb)-4:] == "\u006e\u0075\u006c\u006c" {
			_eea := _adbb[0 : len(_adbb)-4]
			_gc.Log.Trace("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029", _adbb)
			_gc.Log.Trace("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c", _eea)
			_ace.skipSpaces()
			_bfe, _ := _ace._dcf.Peek(1)
			if _bfe[0] == '/' {
				_edb.Set(_eea, _gb.MakeNull())
				continue
			}
		}
		_ace.skipSpaces()
		_edge, _, _edeb := _ace.parseObject()
		if _edeb != nil {
			return nil, _edeb
		}
		_edb.Set(_adbb, _edge)
		_gc.Log.Trace("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073", _adbb, _edge.String())
	}
	return _edb, nil
}

// String is same as Bytes() except returns as a string for convenience.
func (_agb *ContentCreator) String() string { return string(_agb._ec.Bytes()) }

// Add_q adds 'q' operand to the content stream: Pushes the current graphics state on the stack.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_aa *ContentCreator) Add_q() *ContentCreator {
	_cae := ContentStreamOperation{}
	_cae.Operand = "\u0071"
	_aa._ec = append(_aa._ec, &_cae)
	return _aa
}

// Add_n appends 'n' operand to the content stream:
// End the path without filling or stroking.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_bfg *ContentCreator) Add_n() *ContentCreator {
	_cbb := ContentStreamOperation{}
	_cbb.Operand = "\u006e"
	_bfg._ec = append(_bfg._ec, &_cbb)
	return _bfg
}
func (_ggc *ContentStreamParser) parseNull() (_gb.PdfObjectNull, error) {
	_, _fff := _ggc._dcf.Discard(4)
	return _gb.PdfObjectNull{}, _fff
}

// Add_BT appends 'BT' operand to the content stream:
// Begin text.
//
// See section 9.4 "Text Objects" and Table 107 (p. 256 PDF32000_2008).
func (_fgc *ContentCreator) Add_BT() *ContentCreator {
	_fgba := ContentStreamOperation{}
	_fgba.Operand = "\u0042\u0054"
	_fgc._ec = append(_fgc._ec, &_fgba)
	return _fgc
}

// NewContentStreamProcessor returns a new ContentStreamProcessor for operations `ops`.
func NewContentStreamProcessor(ops []*ContentStreamOperation) *ContentStreamProcessor {
	_ffcfb := ContentStreamProcessor{}
	_ffcfb._egcc = GraphicStateStack{}
	_cfcg := GraphicsState{}
	_ffcfb._efaee = _cfcg
	_ffcfb._dcb = []handlerEntry{}
	_ffcfb._gbfc = 0
	_ffcfb._abdd = ops
	return &_ffcfb
}

// HandlerConditionEnum represents the type of operand content stream processor (handler).
// The handler may process a single specific named operand or all operands.
type HandlerConditionEnum int

// Add_Tj appends 'Tj' operand to the content stream:
// Show a text string.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_age *ContentCreator) Add_Tj(textstr _gb.PdfObjectString) *ContentCreator {
	_aee := ContentStreamOperation{}
	_aee.Operand = "\u0054\u006a"
	_aee.Params = _ccbe([]_gb.PdfObjectString{textstr})
	_age._ec = append(_age._ec, &_aee)
	return _age
}
func (_ecgc *ContentStreamParser) skipSpaces() (int, error) {
	_dga := 0
	for {
		_fgbaa, _cad := _ecgc._dcf.Peek(1)
		if _cad != nil {
			return 0, _cad
		}
		if _gb.IsWhiteSpace(_fgbaa[0]) {
			_ecgc._dcf.ReadByte()
			_dga++
		} else {
			break
		}
	}
	return _dga, nil
}

// Add_W appends 'W' operand to the content stream:
// Modify the current clipping path by intersecting with the current path (nonzero winding rule).
//
// See section 8.5.4 "Clipping Path Operators" and Table 61 (p. 146 PDF32000_2008).
func (_fag *ContentCreator) Add_W() *ContentCreator {
	_fdac := ContentStreamOperation{}
	_fdac.Operand = "\u0057"
	_fag._ec = append(_fag._ec, &_fdac)
	return _fag
}

// Add_B appends 'B' operand to the content stream:
// Fill and then stroke the path (nonzero winding number rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_gec *ContentCreator) Add_B() *ContentCreator {
	_ccg := ContentStreamOperation{}
	_ccg.Operand = "\u0042"
	_gec._ec = append(_gec._ec, &_ccg)
	return _gec
}
func (_edab *ContentStreamProcessor) handleCommand_cm(_ggde *ContentStreamOperation, _gfbdf *_ef.PdfPageResources) error {
	if len(_ggde.Params) != 6 {
		_gc.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0063\u006d\u003a\u0020\u0025\u0064", len(_ggde.Params))
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_fbc, _ecfb := _gb.GetNumbersAsFloat(_ggde.Params)
	if _ecfb != nil {
		return _ecfb
	}
	_bccb := _df.NewMatrix(_fbc[0], _fbc[1], _fbc[2], _fbc[3], _fbc[4], _fbc[5])
	_edab._efaee.CTM.Concat(_bccb)
	return nil
}

// All returns true if `hce` is equivalent to HandlerConditionEnumAllOperands.
func (_fdfd HandlerConditionEnum) All() bool { return _fdfd == HandlerConditionEnumAllOperands }
func (_gbf *ContentStreamParser) parseObject() (_abcd _gb.PdfObject, _bcga bool, _gfee error) {
	_gbf.skipSpaces()
	for {
		_fbb, _cfaf := _gbf._dcf.Peek(2)
		if _cfaf != nil {
			return nil, false, _cfaf
		}
		_gc.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_fbb))
		if _fbb[0] == '%' {
			_gbf.skipComments()
			continue
		} else if _fbb[0] == '/' {
			_gfge, _dgbc := _gbf.parseName()
			_gc.Log.Trace("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027", _gfge)
			return &_gfge, false, _dgbc
		} else if _fbb[0] == '(' {
			_gc.Log.Trace("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!")
			_gdaa, _cge := _gbf.parseString()
			return _gdaa, false, _cge
		} else if _fbb[0] == '<' && _fbb[1] != '<' {
			_gc.Log.Trace("\u002d\u003e\u0048\u0065\u0078\u0020\u0053\u0074\u0072\u0069\u006e\u0067\u0021")
			_gfc, _efafb := _gbf.parseHexString()
			return _gfc, false, _efafb
		} else if _fbb[0] == '[' {
			_gc.Log.Trace("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021")
			_bbgf, _gcb := _gbf.parseArray()
			return _bbgf, false, _gcb
		} else if _gb.IsFloatDigit(_fbb[0]) || (_fbb[0] == '-' && _gb.IsFloatDigit(_fbb[1])) || (_fbb[0] == '+' && _gb.IsFloatDigit(_fbb[1])) {
			_gc.Log.Trace("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!")
			_acdf, _agda := _gbf.parseNumber()
			return _acdf, false, _agda
		} else if _fbb[0] == '<' && _fbb[1] == '<' {
			_efee, _ddbe := _gbf.parseDict()
			return _efee, false, _ddbe
		} else {
			_gc.Log.Trace("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f")
			_fbb, _ = _gbf._dcf.Peek(5)
			_efdd := string(_fbb)
			_gc.Log.Trace("\u0063\u006f\u006e\u0074\u0020\u0050\u0065\u0065\u006b\u0020\u0073\u0074r\u003a\u0020\u0025\u0073", _efdd)
			if (len(_efdd) > 3) && (_efdd[:4] == "\u006e\u0075\u006c\u006c") {
				_dcdg, _ecbbb := _gbf.parseNull()
				return &_dcdg, false, _ecbbb
			} else if (len(_efdd) > 4) && (_efdd[:5] == "\u0066\u0061\u006cs\u0065") {
				_dfdd, _acec := _gbf.parseBool()
				return &_dfdd, false, _acec
			} else if (len(_efdd) > 3) && (_efdd[:4] == "\u0074\u0072\u0075\u0065") {
				_gaagc, _ebb := _gbf.parseBool()
				return &_gaagc, false, _ebb
			}
			_dde, _dfg := _gbf.parseOperand()
			if _dfg != nil {
				return _dde, false, _dfg
			}
			if len(_dde.String()) < 1 {
				return _dde, false, ErrInvalidOperand
			}
			return _dde, true, nil
		}
	}
}

// NewInlineImageFromImage makes a new content stream inline image object from an image.
func NewInlineImageFromImage(img _ef.Image, encoder _gb.StreamEncoder) (*ContentStreamInlineImage, error) {
	if encoder == nil {
		encoder = _gb.NewRawEncoder()
	}
	encoder.UpdateParams(img.GetParamsDict())
	_ecc := ContentStreamInlineImage{}
	if img.ColorComponents == 1 {
		_ecc.ColorSpace = _gb.MakeName("\u0047")
	} else if img.ColorComponents == 3 {
		_ecc.ColorSpace = _gb.MakeName("\u0052\u0047\u0042")
	} else if img.ColorComponents == 4 {
		_ecc.ColorSpace = _gb.MakeName("\u0043\u004d\u0059\u004b")
	} else {
		_gc.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006db\u0065\u0072\u0020o\u0066\u0020c\u006f\u006c\u006f\u0072\u0020\u0063o\u006dpo\u006e\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0072\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0064", img.ColorComponents)
		return nil, _a.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020c\u006fl\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073")
	}
	_ecc.BitsPerComponent = _gb.MakeInteger(img.BitsPerComponent)
	_ecc.Width = _gb.MakeInteger(img.Width)
	_ecc.Height = _gb.MakeInteger(img.Height)
	_daga, _dabe := encoder.EncodeBytes(img.Data)
	if _dabe != nil {
		return nil, _dabe
	}
	_ecc._cbac = _daga
	_adb := encoder.GetFilterName()
	if _adb != _gb.StreamEncodingFilterNameRaw {
		_ecc.Filter = _gb.MakeName(_adb)
	}
	return &_ecc, nil
}

// NewContentCreator returns a new initialized ContentCreator.
func NewContentCreator() *ContentCreator {
	_eg := &ContentCreator{}
	_eg._ec = ContentStreamOperations{}
	return _eg
}

// AddOperand adds a specified operand.
func (_bgc *ContentCreator) AddOperand(op ContentStreamOperation) *ContentCreator {
	_bgc._ec = append(_bgc._ec, &op)
	return _bgc
}

// Wrap ensures that the contentstream is wrapped within a balanced q ... Q expression.
func (_aff *ContentCreator) Wrap() { _aff._ec.WrapIfNeeded() }
func _cdfa(_cab _gb.PdfObject) (_ef.PdfColorspace, error) {
	_dgba, _ceed := _cab.(*_gb.PdfObjectArray)
	if !_ceed {
		_gc.Log.Debug("\u0045r\u0072\u006fr\u003a\u0020\u0049\u006ev\u0061\u006c\u0069d\u0020\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020cs\u0020\u006e\u006ft\u0020\u0069n\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025#\u0076\u0029", _cab)
		return nil, _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if _dgba.Len() != 4 {
		_gc.Log.Debug("\u0045\u0072\u0072\u006f\u0072:\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061r\u0072\u0061\u0079\u002c\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0021\u003d\u0020\u0034\u0020\u0028\u0025\u0064\u0029", _dgba.Len())
		return nil, _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_eed, _ceed := _dgba.Get(0).(*_gb.PdfObjectName)
	if !_ceed {
		_gc.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0066\u0069\u0072s\u0074 \u0065\u006c\u0065\u006de\u006e\u0074 \u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0061\u0072\u0072\u0061\u0079\u003a\u0020\u0025\u0023\u0076\u0029", *_dgba)
		return nil, _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_eed != "\u0049" && *_eed != "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		_gc.Log.Debug("\u0045\u0072r\u006f\u0072\u003a\u0020\u0049n\u0076\u0061\u006c\u0069\u0064 \u0063\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0021\u003d\u0020\u0049\u0020\u0028\u0067\u006f\u0074\u003a\u0020\u0025\u0076\u0029", *_eed)
		return nil, _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_eed, _ceed = _dgba.Get(1).(*_gb.PdfObjectName)
	if !_ceed {
		_gc.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072r\u0061\u0079\u0020\u0032\u006e\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0061\u0072\u0072a\u0079\u003a\u0020\u0025\u0023v\u0029", *_dgba)
		return nil, _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_eed != "\u0047" && *_eed != "\u0052\u0047\u0042" && *_eed != "\u0043\u004d\u0059\u004b" && *_eed != "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079" && *_eed != "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B" && *_eed != "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b" {
		_gc.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u0073\u0020\u0061\u0072r\u0061\u0079\u0020\u0032\u006e\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0021\u003d\u0020\u0047\u002f\u0052\u0047\u0042\u002f\u0043\u004d\u0059\u004b\u0020\u0028g\u006f\u0074\u003a\u0020\u0025v\u0029", *_eed)
		return nil, _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
	}
	_aad := ""
	switch *_eed {
	case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
		_aad = "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079"
	case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
		_aad = "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B"
	case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
		_aad = "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
	}
	_gbag := _gb.MakeArray(_gb.MakeName("\u0049n\u0064\u0065\u0078\u0065\u0064"), _gb.MakeName(_aad), _dgba.Get(2), _dgba.Get(3))
	return _ef.NewPdfColorspaceFromPdfObject(_gbag)
}
func (_bage *ContentStreamInlineImage) String() string {
	_faf := _d.Sprintf("I\u006el\u0069\u006e\u0065\u0049\u006d\u0061\u0067\u0065(\u006c\u0065\u006e\u003d%d\u0029\u000a", len(_bage._cbac))
	if _bage.BitsPerComponent != nil {
		_faf += "\u002d\u0020\u0042\u0050\u0043\u0020" + _bage.BitsPerComponent.WriteString() + "\u000a"
	}
	if _bage.ColorSpace != nil {
		_faf += "\u002d\u0020\u0043S\u0020" + _bage.ColorSpace.WriteString() + "\u000a"
	}
	if _bage.Decode != nil {
		_faf += "\u002d\u0020\u0044\u0020" + _bage.Decode.WriteString() + "\u000a"
	}
	if _bage.DecodeParms != nil {
		_faf += "\u002d\u0020\u0044P\u0020" + _bage.DecodeParms.WriteString() + "\u000a"
	}
	if _bage.Filter != nil {
		_faf += "\u002d\u0020\u0046\u0020" + _bage.Filter.WriteString() + "\u000a"
	}
	if _bage.Height != nil {
		_faf += "\u002d\u0020\u0048\u0020" + _bage.Height.WriteString() + "\u000a"
	}
	if _bage.ImageMask != nil {
		_faf += "\u002d\u0020\u0049M\u0020" + _bage.ImageMask.WriteString() + "\u000a"
	}
	if _bage.Intent != nil {
		_faf += "\u002d \u0049\u006e\u0074\u0065\u006e\u0074 " + _bage.Intent.WriteString() + "\u000a"
	}
	if _bage.Interpolate != nil {
		_faf += "\u002d\u0020\u0049\u0020" + _bage.Interpolate.WriteString() + "\u000a"
	}
	if _bage.Width != nil {
		_faf += "\u002d\u0020\u0057\u0020" + _bage.Width.WriteString() + "\u000a"
	}
	return _faf
}

// Add_Tstar appends 'T*' operand to the content stream:
// Move to the start of next line.
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_gge *ContentCreator) Add_Tstar() *ContentCreator {
	_bdee := ContentStreamOperation{}
	_bdee.Operand = "\u0054\u002a"
	_gge._ec = append(_gge._ec, &_bdee)
	return _gge
}
func (_ccf *ContentStreamParser) parseName() (_gb.PdfObjectName, error) {
	_eba := ""
	_feb := false
	for {
		_gdf, _bdef := _ccf._dcf.Peek(1)
		if _bdef == _e.EOF {
			break
		}
		if _bdef != nil {
			return _gb.PdfObjectName(_eba), _bdef
		}
		if !_feb {
			if _gdf[0] == '/' {
				_feb = true
				_ccf._dcf.ReadByte()
			} else {
				_gc.Log.Error("N\u0061\u006d\u0065\u0020\u0073\u0074a\u0072\u0074\u0069\u006e\u0067\u0020\u0077\u0069\u0074h\u0020\u0025\u0073 \u0028%\u0020\u0078\u0029", _gdf, _gdf)
				return _gb.PdfObjectName(_eba), _d.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029", _gdf[0])
			}
		} else {
			if _gb.IsWhiteSpace(_gdf[0]) {
				break
			} else if (_gdf[0] == '/') || (_gdf[0] == '[') || (_gdf[0] == '(') || (_gdf[0] == ']') || (_gdf[0] == '<') || (_gdf[0] == '>') {
				break
			} else if _gdf[0] == '#' {
				_ggfe, _afd := _ccf._dcf.Peek(3)
				if _afd != nil {
					return _gb.PdfObjectName(_eba), _afd
				}
				_ccf._dcf.Discard(3)
				_agec, _afd := _g.DecodeString(string(_ggfe[1:3]))
				if _afd != nil {
					return _gb.PdfObjectName(_eba), _afd
				}
				_eba += string(_agec)
			} else {
				_daf, _ := _ccf._dcf.ReadByte()
				_eba += string(_daf)
			}
		}
	}
	return _gb.PdfObjectName(_eba), nil
}
func (_edgf *ContentStreamParser) parseHexString() (*_gb.PdfObjectString, error) {
	_edgf._dcf.ReadByte()
	_fga := []byte("\u0030\u0031\u0032\u003345\u0036\u0037\u0038\u0039\u0061\u0062\u0063\u0064\u0065\u0066\u0041\u0042\u0043\u0044E\u0046")
	var _cga []byte
	for {
		_edgf.skipSpaces()
		_ecga, _cfc := _edgf._dcf.Peek(1)
		if _cfc != nil {
			return _gb.MakeString(""), _cfc
		}
		if _ecga[0] == '>' {
			_edgf._dcf.ReadByte()
			break
		}
		_aafb, _ := _edgf._dcf.ReadByte()
		if _de.IndexByte(_fga, _aafb) >= 0 {
			_cga = append(_cga, _aafb)
		}
	}
	if len(_cga)%2 == 1 {
		_cga = append(_cga, '0')
	}
	_ege, _ := _g.DecodeString(string(_cga))
	return _gb.MakeHexString(string(_ege)), nil
}

// Add_Tm appends 'Tm' operand to the content stream:
// Set the text line matrix.
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_eaag *ContentCreator) Add_Tm(a, b, c, d, e, f float64) *ContentCreator {
	_bbca := ContentStreamOperation{}
	_bbca.Operand = "\u0054\u006d"
	_bbca.Params = _bcd([]float64{a, b, c, d, e, f})
	_eaag._ec = append(_eaag._ec, &_bbca)
	return _eaag
}

// String returns `ops.Bytes()` as a string.
func (_befb *ContentStreamOperations) String() string { return string(_befb.Bytes()) }

// Process processes the entire list of operations. Maintains the graphics state that is passed to any
// handlers that are triggered during processing (either on specific operators or all).
func (_fef *ContentStreamProcessor) Process(resources *_ef.PdfPageResources) error {
	_fef._efaee.ColorspaceStroking = _ef.NewPdfColorspaceDeviceGray()
	_fef._efaee.ColorspaceNonStroking = _ef.NewPdfColorspaceDeviceGray()
	_fef._efaee.ColorStroking = _ef.NewPdfColorDeviceGray(0)
	_fef._efaee.ColorNonStroking = _ef.NewPdfColorDeviceGray(0)
	_fef._efaee.CTM = _df.IdentityMatrix()
	for _, _bedg := range _fef._abdd {
		var _cde error
		switch _bedg.Operand {
		case "\u0071":
			_fef._egcc.Push(_fef._efaee)
		case "\u0051":
			if len(_fef._egcc) == 0 {
				_gc.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0060\u0051\u0060\u0020\u006f\u0070e\u0072\u0061\u0074\u006f\u0072\u002e\u0020\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0073\u0020\u0073\u0074\u0061\u0074\u0065 \u0073\u0074\u0061\u0063\u006b\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079.\u0020\u0053\u006bi\u0070\u0070\u0069\u006e\u0067\u002e")
				continue
			}
			_fef._efaee = _fef._egcc.Pop()
		case "\u0043\u0053":
			_cde = _fef.handleCommand_CS(_bedg, resources)
		case "\u0063\u0073":
			_cde = _fef.handleCommand_cs(_bedg, resources)
		case "\u0053\u0043":
			_cde = _fef.handleCommand_SC(_bedg, resources)
		case "\u0053\u0043\u004e":
			_cde = _fef.handleCommand_SCN(_bedg, resources)
		case "\u0073\u0063":
			_cde = _fef.handleCommand_sc(_bedg, resources)
		case "\u0073\u0063\u006e":
			_cde = _fef.handleCommand_scn(_bedg, resources)
		case "\u0047":
			_cde = _fef.handleCommand_G(_bedg, resources)
		case "\u0067":
			_cde = _fef.handleCommand_g(_bedg, resources)
		case "\u0052\u0047":
			_cde = _fef.handleCommand_RG(_bedg, resources)
		case "\u0072\u0067":
			_cde = _fef.handleCommand_rg(_bedg, resources)
		case "\u004b":
			_cde = _fef.handleCommand_K(_bedg, resources)
		case "\u006b":
			_cde = _fef.handleCommand_k(_bedg, resources)
		case "\u0063\u006d":
			_cde = _fef.handleCommand_cm(_bedg, resources)
		}
		if _cde != nil {
			_gc.Log.Debug("\u0050\u0072\u006f\u0063\u0065\u0073s\u006f\u0072\u0020\u0068\u0061\u006e\u0064\u006c\u0069\u006e\u0067\u0020\u0065r\u0072\u006f\u0072\u0020\u0028\u0025\u0073)\u003a\u0020\u0025\u0076", _bedg.Operand, _cde)
			_gc.Log.Debug("\u004f\u0070\u0065r\u0061\u006e\u0064\u003a\u0020\u0025\u0023\u0076", _bedg.Operand)
			return _cde
		}
		for _, _egbg := range _fef._dcb {
			var _ggee error
			if _egbg.Condition.All() {
				_ggee = _egbg.Handler(_bedg, _fef._efaee, resources)
			} else if _egbg.Condition.Operand() && _bedg.Operand == _egbg.Operand {
				_ggee = _egbg.Handler(_bedg, _fef._efaee, resources)
			}
			if _ggee != nil {
				_gc.Log.Debug("P\u0072\u006f\u0063\u0065\u0073\u0073o\u0072\u0020\u0068\u0061\u006e\u0064\u006c\u0065\u0072 \u0065\u0072\u0072o\u0072:\u0020\u0025\u0076", _ggee)
				return _ggee
			}
		}
	}
	return nil
}

// Add_k appends 'k' operand to the content stream:
// Same as K but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_aab *ContentCreator) Add_k(c, m, y, k float64) *ContentCreator {
	_geg := ContentStreamOperation{}
	_geg.Operand = "\u006b"
	_geg.Params = _bcd([]float64{c, m, y, k})
	_aab._ec = append(_aab._ec, &_geg)
	return _aab
}
func _bgf(_aabd *ContentStreamInlineImage) (_gb.StreamEncoder, error) {
	if _aabd.Filter == nil {
		return _gb.NewRawEncoder(), nil
	}
	_gbcf, _deea := _aabd.Filter.(*_gb.PdfObjectName)
	if !_deea {
		_fgd, _efa := _aabd.Filter.(*_gb.PdfObjectArray)
		if !_efa {
			return nil, _d.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006f\u0072 \u0041\u0072\u0072\u0061\u0079\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
		if _fgd.Len() == 0 {
			return _gb.NewRawEncoder(), nil
		}
		if _fgd.Len() != 1 {
			_cag, _aeb := _gfg(_aabd)
			if _aeb != nil {
				_gc.Log.Error("\u0046\u0061\u0069\u006c\u0065\u0064 \u0063\u0072\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u006d\u0075\u006c\u0074i\u0020\u0065\u006e\u0063\u006f\u0064\u0065r\u003a\u0020\u0025\u0076", _aeb)
				return nil, _aeb
			}
			_gc.Log.Trace("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063:\u0020\u0025\u0073\u000a", _cag)
			return _cag, nil
		}
		_fcb := _fgd.Get(0)
		_gbcf, _efa = _fcb.(*_gb.PdfObjectName)
		if !_efa {
			return nil, _d.Errorf("\u0066\u0069l\u0074\u0065\u0072\u0020a\u0072\u0072a\u0079\u0020\u006d\u0065\u006d\u0062\u0065\u0072 \u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061\u006d\u0065\u0020\u006fb\u006a\u0065\u0063\u0074")
		}
	}
	switch *_gbcf {
	case "\u0041\u0048\u0078", "\u0041\u0053\u0043\u0049\u0049\u0048\u0065\u0078\u0044e\u0063\u006f\u0064\u0065":
		return _gb.NewASCIIHexEncoder(), nil
	case "\u0041\u0038\u0035", "\u0041\u0053\u0043\u0049\u0049\u0038\u0035\u0044\u0065\u0063\u006f\u0064\u0065":
		return _gb.NewASCII85Encoder(), nil
	case "\u0044\u0043\u0054", "\u0044C\u0054\u0044\u0065\u0063\u006f\u0064e":
		return _fddd(_aabd)
	case "\u0046\u006c", "F\u006c\u0061\u0074\u0065\u0044\u0065\u0063\u006f\u0064\u0065":
		return _aca(_aabd, nil)
	case "\u004c\u005a\u0057", "\u004cZ\u0057\u0044\u0065\u0063\u006f\u0064e":
		return _caff(_aabd, nil)
	case "\u0043\u0043\u0046", "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065":
		return _gb.NewCCITTFaxEncoder(), nil
	case "\u0052\u004c", "\u0052u\u006eL\u0065\u006e\u0067\u0074\u0068\u0044\u0065\u0063\u006f\u0064\u0065":
		return _gb.NewRunLengthEncoder(), nil
	default:
		_gc.Log.Debug("\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065 \u0069\u006d\u0061\u0067\u0065\u0020\u0065n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u003a\u0020\u0025\u0073", *_gbcf)
		return nil, _a.New("\u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006el\u0069n\u0065 \u0065n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006f\u0064")
	}
}

// Add_TL appends 'TL' operand to the content stream:
// Set leading.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_bdb *ContentCreator) Add_TL(leading float64) *ContentCreator {
	_efe := ContentStreamOperation{}
	_efe.Operand = "\u0054\u004c"
	_efe.Params = _bcd([]float64{leading})
	_bdb._ec = append(_bdb._ec, &_efe)
	return _bdb
}

// GraphicStateStack represents a stack of GraphicsState.
type GraphicStateStack []GraphicsState

// Add_B_starred appends 'B*' operand to the content stream:
// Fill and then stroke the path (even-odd rule).
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_cef *ContentCreator) Add_B_starred() *ContentCreator {
	_gbc := ContentStreamOperation{}
	_gbc.Operand = "\u0042\u002a"
	_cef._ec = append(_cef._ec, &_gbc)
	return _cef
}

// Add_Q adds 'Q' operand to the content stream: Pops the most recently stored state from the stack.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_egc *ContentCreator) Add_Q() *ContentCreator {
	_deeb := ContentStreamOperation{}
	_deeb.Operand = "\u0051"
	_egc._ec = append(_egc._ec, &_deeb)
	return _egc
}

// Add_TD appends 'TD' operand to the content stream:
// Move to start of next line with offset (`tx`, `ty`).
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_faec *ContentCreator) Add_TD(tx, ty float64) *ContentCreator {
	_cccd := ContentStreamOperation{}
	_cccd.Operand = "\u0054\u0044"
	_cccd.Params = _bcd([]float64{tx, ty})
	_faec._ec = append(_faec._ec, &_cccd)
	return _faec
}

type handlerEntry struct {
	Condition HandlerConditionEnum
	Operand   string
	Handler   HandlerFunc
}

func (_agef *ContentStreamProcessor) handleCommand_g(_bbeb *ContentStreamOperation, _gcbc *_ef.PdfPageResources) error {
	_fdce := _ef.NewPdfColorspaceDeviceGray()
	if len(_bbeb.Params) != _fdce.GetNumComponents() {
		_gc.Log.Debug("\u0049\u006e\u0076al\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020p\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0067")
		_gc.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_bbeb.Params), _fdce)
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_fcbb, _bcfb := _fdce.ColorFromPdfObjects(_bbeb.Params)
	if _bcfb != nil {
		_gc.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0068\u0061\u006e\u0064\u006c\u0065\u0043o\u006d\u006d\u0061\u006e\u0064\u005f\u0067\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061r\u0061\u006d\u0073\u002e\u0020c\u0073\u003d\u0025\u0054\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fdce, _bbeb, _bcfb)
		return _bcfb
	}
	_agef._efaee.ColorspaceNonStroking = _fdce
	_agef._efaee.ColorNonStroking = _fcbb
	return nil
}

// Add_v appends 'v' operand to the content stream: Append a Bezier curve to the current path from the
// current point to (x3,y3) with the current point and (x2,y2) as control points.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_aeg *ContentCreator) Add_v(x2, y2, x3, y3 float64) *ContentCreator {
	_cce := ContentStreamOperation{}
	_cce.Operand = "\u0076"
	_cce.Params = _bcd([]float64{x2, y2, x3, y3})
	_aeg._ec = append(_aeg._ec, &_cce)
	return _aeg
}

// IsMask checks if an image is a mask.
// The image mask entry in the image dictionary specifies that the image data shall be used as a stencil
// mask for painting in the current color. The mask data is 1bpc, grayscale.
func (_fac *ContentStreamInlineImage) IsMask() (bool, error) {
	if _fac.ImageMask != nil {
		_agdf, _bbad := _fac.ImageMask.(*_gb.PdfObjectBool)
		if !_bbad {
			_gc.Log.Debug("\u0049m\u0061\u0067\u0065\u0020\u006d\u0061\u0073\u006b\u0020\u006e\u006ft\u0020\u0061\u0020\u0062\u006f\u006f\u006c\u0065\u0061\u006e")
			return false, _a.New("\u0069\u006e\u0076\u0061li\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0079\u0070\u0065")
		}
		return bool(*_agdf), nil
	}
	return false, nil
}
func _ebg(_bgca string) bool { _, _afe := _dbeb[_bgca]; return _afe }

// Add_Tf appends 'Tf' operand to the content stream:
// Set font and font size specified by font resource `fontName` and `fontSize`.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_feg *ContentCreator) Add_Tf(fontName _gb.PdfObjectName, fontSize float64) *ContentCreator {
	_ccc := ContentStreamOperation{}
	_ccc.Operand = "\u0054\u0066"
	_ccc.Params = _afec([]_gb.PdfObjectName{fontName})
	_ccc.Params = append(_ccc.Params, _bcd([]float64{fontSize})...)
	_feg._ec = append(_feg._ec, &_ccc)
	return _feg
}
func (_fdgcf *ContentStreamProcessor) handleCommand_rg(_cdaf *ContentStreamOperation, _cfb *_ef.PdfPageResources) error {
	_gde := _ef.NewPdfColorspaceDeviceRGB()
	if len(_cdaf.Params) != _gde.GetNumComponents() {
		_gc.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_gc.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_cdaf.Params), _gde)
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_bffd, _fdfdb := _gde.ColorFromPdfObjects(_cdaf.Params)
	if _fdfdb != nil {
		return _fdfdb
	}
	_fdgcf._efaee.ColorspaceNonStroking = _gde
	_fdgcf._efaee.ColorNonStroking = _bffd
	return nil
}

// ToImage exports the inline image to Image which can be transformed or exported easily.
// Page resources are needed to look up colorspace information.
func (_eaagf *ContentStreamInlineImage) ToImage(resources *_ef.PdfPageResources) (*_ef.Image, error) {
	_gaag, _cbbc := _eaagf.toImageBase(resources)
	if _cbbc != nil {
		return nil, _cbbc
	}
	_gea, _cbbc := _bgf(_eaagf)
	if _cbbc != nil {
		return nil, _cbbc
	}
	_fbg, _dcd := _gb.GetDict(_eaagf.DecodeParms)
	if _dcd {
		_gea.UpdateParams(_fbg)
	}
	_gc.Log.Trace("\u0065n\u0063o\u0064\u0065\u0072\u003a\u0020\u0025\u002b\u0076\u0020\u0025\u0054", _gea, _gea)
	_gc.Log.Trace("\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065:\u0020\u0025\u002b\u0076", _eaagf)
	_faa, _cbbc := _gea.DecodeBytes(_eaagf._cbac)
	if _cbbc != nil {
		return nil, _cbbc
	}
	_cbc := &_ef.Image{Width: int64(_gaag.Width), Height: int64(_gaag.Height), BitsPerComponent: int64(_gaag.BitsPerComponent), ColorComponents: _gaag.ColorComponents, Data: _faa}
	if len(_gaag.Decode) > 0 {
		for _cacd := 0; _cacd < len(_gaag.Decode); _cacd++ {
			_gaag.Decode[_cacd] *= float64((int(1) << uint(_gaag.BitsPerComponent)) - 1)
		}
		_cbc.SetDecode(_gaag.Decode)
	}
	return _cbc, nil
}
func (_efag *ContentStreamProcessor) handleCommand_CS(_bad *ContentStreamOperation, _eab *_ef.PdfPageResources) error {
	if len(_bad.Params) < 1 {
		_gc.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0063\u0073\u0020\u0063\u006f\u006d\u006d\u0061\u006e\u0064\u002c\u0020s\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u006f\u0076\u0065\u0072")
		return _a.New("\u0074o\u006f \u0066\u0065\u0077\u0020\u0070a\u0072\u0061m\u0065\u0074\u0065\u0072\u0073")
	}
	if len(_bad.Params) > 1 {
		_gc.Log.Debug("\u0063\u0073\u0020\u0063\u006f\u006d\u006d\u0061n\u0064\u0020\u0077it\u0068\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020-\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075i\u006e\u0067")
		return _a.New("\u0074\u006f\u006f\u0020ma\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	}
	_adbe, _faea := _bad.Params[0].(*_gb.PdfObjectName)
	if !_faea {
		_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020c\u0073\u0020\u0063o\u006d\u006d\u0061n\u0064\u0020w\u0069\u0074\u0068\u0020\u0069\u006ev\u0061li\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u002c\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072")
		return _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_dffb, _bfc := _efag.getColorspace(string(*_adbe), _eab)
	if _bfc != nil {
		return _bfc
	}
	_efag._efaee.ColorspaceStroking = _dffb
	_gad, _bfc := _efag.getInitialColor(_dffb)
	if _bfc != nil {
		return _bfc
	}
	_efag._efaee.ColorStroking = _gad
	return nil
}
func (_fcg *ContentStreamParser) skipComments() error {
	if _, _ecbg := _fcg.skipSpaces(); _ecbg != nil {
		return _ecbg
	}
	_efc := true
	for {
		_dddg, _adca := _fcg._dcf.Peek(1)
		if _adca != nil {
			_gc.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _adca.Error())
			return _adca
		}
		if _efc && _dddg[0] != '%' {
			return nil
		}
		_efc = false
		if (_dddg[0] != '\r') && (_dddg[0] != '\n') {
			_fcg._dcf.ReadByte()
		} else {
			break
		}
	}
	return _fcg.skipComments()
}

// Add_rg appends 'rg' operand to the content stream:
// Same as RG but used for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_eeb *ContentCreator) Add_rg(r, g, b float64) *ContentCreator {
	_aed := ContentStreamOperation{}
	_aed.Operand = "\u0072\u0067"
	_aed.Params = _bcd([]float64{r, g, b})
	_eeb._ec = append(_eeb._ec, &_aed)
	return _eeb
}

// ContentCreator is a builder for PDF content streams.
type ContentCreator struct{ _ec ContentStreamOperations }

// Add_SCN appends 'SCN' operand to the content stream:
// Same as SC but supports more colorspaces.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_cdd *ContentCreator) Add_SCN(c ...float64) *ContentCreator {
	_fbf := ContentStreamOperation{}
	_fbf.Operand = "\u0053\u0043\u004e"
	_fbf.Params = _bcd(c)
	_cdd._ec = append(_cdd._ec, &_fbf)
	return _cdd
}

// Add_Tc appends 'Tc' operand to the content stream:
// Set character spacing.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_bfbe *ContentCreator) Add_Tc(charSpace float64) *ContentCreator {
	_dca := ContentStreamOperation{}
	_dca.Operand = "\u0054\u0063"
	_dca.Params = _bcd([]float64{charSpace})
	_bfbe._ec = append(_bfbe._ec, &_dca)
	return _bfbe
}

// Add_TJ appends 'TJ' operand to the content stream:
// Show one or more text string. Array of numbers (displacement) and strings.
//
// See section 9.4.3 "Text Showing Operators" and
// Table 209 (pp. 258-259 PDF32000_2008).
func (_fcd *ContentCreator) Add_TJ(vals ..._gb.PdfObject) *ContentCreator {
	_dabf := ContentStreamOperation{}
	_dabf.Operand = "\u0054\u004a"
	_dabf.Params = []_gb.PdfObject{_gb.MakeArray(vals...)}
	_fcd._ec = append(_fcd._ec, &_dabf)
	return _fcd
}
func (_dcg *ContentStreamProcessor) handleCommand_SCN(_afad *ContentStreamOperation, _cbg *_ef.PdfPageResources) error {
	_aedd := _dcg._efaee.ColorspaceStroking
	if !_cfd(_aedd) {
		if len(_afad.Params) != _aedd.GetNumComponents() {
			_gc.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_gc.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_afad.Params), _aedd)
			return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_bdff, _debe := _aedd.ColorFromPdfObjects(_afad.Params)
	if _debe != nil {
		return _debe
	}
	_dcg._efaee.ColorStroking = _bdff
	return nil
}

// Transform returns coordinates x, y transformed by the CTM.
func (_cgba *GraphicsState) Transform(x, y float64) (float64, float64) {
	return _cgba.CTM.Transform(x, y)
}

// Add_scn appends 'scn' operand to the content stream:
// Same as SC but for nonstroking operations.
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_bdfa *ContentCreator) Add_scn(c ...float64) *ContentCreator {
	_eaa := ContentStreamOperation{}
	_eaa.Operand = "\u0073\u0063\u006e"
	_eaa.Params = _bcd(c)
	_bdfa._ec = append(_bdfa._ec, &_eaa)
	return _bdfa
}
func (_ded *ContentStreamProcessor) handleCommand_k(_fcce *ContentStreamOperation, _bbfe *_ef.PdfPageResources) error {
	_dcfb := _ef.NewPdfColorspaceDeviceCMYK()
	if len(_fcce.Params) != _dcfb.GetNumComponents() {
		_gc.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
		_gc.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_fcce.Params), _dcfb)
		return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
	}
	_ccb, _fgac := _dcfb.ColorFromPdfObjects(_fcce.Params)
	if _fgac != nil {
		return _fgac
	}
	_ded._efaee.ColorspaceNonStroking = _dcfb
	_ded._efaee.ColorNonStroking = _ccb
	return nil
}

// WriteString outputs the object as it is to be written to file.
func (_gfb *ContentStreamInlineImage) WriteString() string {
	var _ecb _de.Buffer
	_babg := ""
	if _gfb.BitsPerComponent != nil {
		_babg += "\u002f\u0042\u0050C\u0020" + _gfb.BitsPerComponent.WriteString() + "\u000a"
	}
	if _gfb.ColorSpace != nil {
		_babg += "\u002f\u0043\u0053\u0020" + _gfb.ColorSpace.WriteString() + "\u000a"
	}
	if _gfb.Decode != nil {
		_babg += "\u002f\u0044\u0020" + _gfb.Decode.WriteString() + "\u000a"
	}
	if _gfb.DecodeParms != nil {
		_babg += "\u002f\u0044\u0050\u0020" + _gfb.DecodeParms.WriteString() + "\u000a"
	}
	if _gfb.Filter != nil {
		_babg += "\u002f\u0046\u0020" + _gfb.Filter.WriteString() + "\u000a"
	}
	if _gfb.Height != nil {
		_babg += "\u002f\u0048\u0020" + _gfb.Height.WriteString() + "\u000a"
	}
	if _gfb.ImageMask != nil {
		_babg += "\u002f\u0049\u004d\u0020" + _gfb.ImageMask.WriteString() + "\u000a"
	}
	if _gfb.Intent != nil {
		_babg += "\u002f\u0049\u006e\u0074\u0065\u006e\u0074\u0020" + _gfb.Intent.WriteString() + "\u000a"
	}
	if _gfb.Interpolate != nil {
		_babg += "\u002f\u0049\u0020" + _gfb.Interpolate.WriteString() + "\u000a"
	}
	if _gfb.Width != nil {
		_babg += "\u002f\u0057\u0020" + _gfb.Width.WriteString() + "\u000a"
	}
	_ecb.WriteString(_babg)
	_ecb.WriteString("\u0049\u0044\u0020")
	_ecb.Write(_gfb._cbac)
	_ecb.WriteString("\u000a\u0045\u0049\u000a")
	return _ecb.String()
}

// ContentStreamParser represents a content stream parser for parsing content streams in PDFs.
type ContentStreamParser struct{ _dcf *_b.Reader }

// Operand returns true if `hce` is equivalent to HandlerConditionEnumOperand.
func (_cbcgc HandlerConditionEnum) Operand() bool { return _cbcgc == HandlerConditionEnumOperand }
func _afec(_faegb []_gb.PdfObjectName) []_gb.PdfObject {
	var _efega []_gb.PdfObject
	for _, _ecec := range _faegb {
		_efega = append(_efega, _gb.MakeName(string(_ecec)))
	}
	return _efega
}

// Add_Td appends 'Td' operand to the content stream:
// Move to start of next line with offset (`tx`, `ty`).
//
// See section 9.4.2 "Text Positioning Operators" and
// Table 108 (pp. 257-258 PDF32000_2008).
func (_dea *ContentCreator) Add_Td(tx, ty float64) *ContentCreator {
	_cca := ContentStreamOperation{}
	_cca.Operand = "\u0054\u0064"
	_cca.Params = _bcd([]float64{tx, ty})
	_dea._ec = append(_dea._ec, &_cca)
	return _dea
}

// GetColorSpace returns the colorspace of the inline image.
func (_ggf *ContentStreamInlineImage) GetColorSpace(resources *_ef.PdfPageResources) (_ef.PdfColorspace, error) {
	if _ggf.ColorSpace == nil {
		_gc.Log.Debug("\u0049\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076i\u006e\u0067\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u002c\u0020\u0061\u0073\u0073\u0075\u006di\u006e\u0067\u0020\u0047\u0072a\u0079")
		return _ef.NewPdfColorspaceDeviceGray(), nil
	}
	if _gcgb, _feee := _ggf.ColorSpace.(*_gb.PdfObjectArray); _feee {
		return _cdfa(_gcgb)
	}
	_eace, _adc := _ggf.ColorSpace.(*_gb.PdfObjectName)
	if !_adc {
		_gc.Log.Debug("E\u0072\u0072\u006f\u0072\u003a\u0020I\u006e\u0076\u0061\u006c\u0069\u0064 \u006f\u0062\u006a\u0065\u0063\u0074\u0020t\u0079\u0070\u0065\u0020\u0028\u0025\u0054\u003b\u0025\u002bv\u0029", _ggf.ColorSpace, _ggf.ColorSpace)
		return nil, _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	if *_eace == "\u0047" || *_eace == "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079" {
		return _ef.NewPdfColorspaceDeviceGray(), nil
	} else if *_eace == "\u0052\u0047\u0042" || *_eace == "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B" {
		return _ef.NewPdfColorspaceDeviceRGB(), nil
	} else if *_eace == "\u0043\u004d\u0059\u004b" || *_eace == "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b" {
		return _ef.NewPdfColorspaceDeviceCMYK(), nil
	} else if *_eace == "\u0049" || *_eace == "\u0049n\u0064\u0065\u0078\u0065\u0064" {
		return nil, _a.New("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0049\u006e\u0064e\u0078 \u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063\u0065")
	} else {
		if resources.ColorSpace == nil {
			_gc.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0073", *_eace)
			return nil, _a.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0063o\u006c\u006fr\u0073\u0070\u0061\u0063\u0065")
		}
		_efae, _fcc := resources.GetColorspaceByName(*_eace)
		if !_fcc {
			_gc.Log.Debug("\u0045\u0072r\u006f\u0072\u002c\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006e\u006c\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0073", *_eace)
			return nil, _a.New("\u0075n\u006bn\u006f\u0077\u006e\u0020\u0063o\u006c\u006fr\u0073\u0070\u0061\u0063\u0065")
		}
		return _efae, nil
	}
}
func _bcd(_edeg []float64) []_gb.PdfObject {
	var _cgaf []_gb.PdfObject
	for _, _aegg := range _edeg {
		_cgaf = append(_cgaf, _gb.MakeFloat(_aegg))
	}
	return _cgaf
}

// ContentStreamOperations is a slice of ContentStreamOperations.
type ContentStreamOperations []*ContentStreamOperation

// ContentStreamProcessor defines a data structure and methods for processing a content stream, keeping track of the
// current graphics state, and allowing external handlers to define their own functions as a part of the processing,
// for example rendering or extracting certain information.
type ContentStreamProcessor struct {
	_egcc  GraphicStateStack
	_abdd  []*ContentStreamOperation
	_efaee GraphicsState
	_dcb   []handlerEntry
	_gbfc  int
}

// WrapIfNeeded wraps the entire contents within q ... Q.  If unbalanced, then adds extra Qs at the end.
// Only does if needed. Ensures that when adding new content, one start with all states
// in the default condition.
func (_gcg *ContentStreamOperations) WrapIfNeeded() *ContentStreamOperations {
	if len(*_gcg) == 0 {
		return _gcg
	}
	if _gcg.isWrapped() {
		return _gcg
	}
	*_gcg = append([]*ContentStreamOperation{{Operand: "\u0071"}}, *_gcg...)
	_ga := 0
	for _, _bc := range *_gcg {
		if _bc.Operand == "\u0071" {
			_ga++
		} else if _bc.Operand == "\u0051" {
			_ga--
		}
	}
	for _ga > 0 {
		*_gcg = append(*_gcg, &ContentStreamOperation{Operand: "\u0051"})
		_ga--
	}
	return _gcg
}

// Add_ET appends 'ET' operand to the content stream:
// End text.
//
// See section 9.4 "Text Objects" and Table 107 (p. 256 PDF32000_2008).
func (_ecd *ContentCreator) Add_ET() *ContentCreator {
	_gaf := ContentStreamOperation{}
	_gaf.Operand = "\u0045\u0054"
	_ecd._ec = append(_ecd._ec, &_gaf)
	return _ecd
}
func (_bgb *ContentStreamProcessor) getInitialColor(_eff _ef.PdfColorspace) (_ef.PdfColor, error) {
	switch _bcc := _eff.(type) {
	case *_ef.PdfColorspaceDeviceGray:
		return _ef.NewPdfColorDeviceGray(0.0), nil
	case *_ef.PdfColorspaceDeviceRGB:
		return _ef.NewPdfColorDeviceRGB(0.0, 0.0, 0.0), nil
	case *_ef.PdfColorspaceDeviceCMYK:
		return _ef.NewPdfColorDeviceCMYK(0.0, 0.0, 0.0, 1.0), nil
	case *_ef.PdfColorspaceCalGray:
		return _ef.NewPdfColorCalGray(0.0), nil
	case *_ef.PdfColorspaceCalRGB:
		return _ef.NewPdfColorCalRGB(0.0, 0.0, 0.0), nil
	case *_ef.PdfColorspaceLab:
		_geag := 0.0
		_edf := 0.0
		_abe := 0.0
		if _bcc.Range[0] > 0 {
			_geag = _bcc.Range[0]
		}
		if _bcc.Range[2] > 0 {
			_edf = _bcc.Range[2]
		}
		return _ef.NewPdfColorLab(_geag, _edf, _abe), nil
	case *_ef.PdfColorspaceICCBased:
		if _bcc.Alternate == nil {
			_gc.Log.Trace("\u0049\u0043\u0043\u0020\u0042\u0061\u0073\u0065\u0064\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065d\u0020-\u0020\u0061\u0074\u0074\u0065\u006d\u0070\u0074\u0069\u006e\u0067\u0020\u0066\u0061\u006c\u006c\u0020\u0062a\u0063\u006b\u0020\u0028\u004e\u0020\u003d\u0020\u0025\u0064\u0029", _bcc.N)
			if _bcc.N == 1 {
				_gc.Log.Trace("\u0046\u0061\u006c\u006c\u0069\u006e\u0067\u0020\u0062\u0061\u0063k\u0020\u0074\u006f\u0020\u0044\u0065\u0076\u0069\u0063\u0065G\u0072\u0061\u0079")
				return _bgb.getInitialColor(_ef.NewPdfColorspaceDeviceGray())
			} else if _bcc.N == 3 {
				_gc.Log.Trace("\u0046a\u006c\u006c\u0069\u006eg\u0020\u0062\u0061\u0063\u006b \u0074o\u0020D\u0065\u0076\u0069\u0063\u0065\u0052\u0047B")
				return _bgb.getInitialColor(_ef.NewPdfColorspaceDeviceRGB())
			} else if _bcc.N == 4 {
				_gc.Log.Trace("\u0046\u0061\u006c\u006c\u0069\u006e\u0067\u0020\u0062\u0061\u0063k\u0020\u0074\u006f\u0020\u0044\u0065\u0076\u0069\u0063\u0065C\u004d\u0059\u004b")
				return _bgb.getInitialColor(_ef.NewPdfColorspaceDeviceCMYK())
			} else {
				return nil, _a.New("a\u006c\u0074\u0065\u0072\u006e\u0061t\u0065\u0020\u0073\u0070\u0061\u0063e\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0049C\u0043")
			}
		}
		return _bgb.getInitialColor(_bcc.Alternate)
	case *_ef.PdfColorspaceSpecialIndexed:
		if _bcc.Base == nil {
			return nil, _a.New("\u0069\u006e\u0064\u0065\u0078\u0065\u0064\u0020\u0062\u0061\u0073e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069f\u0069\u0065\u0064")
		}
		return _bgb.getInitialColor(_bcc.Base)
	case *_ef.PdfColorspaceSpecialSeparation:
		if _bcc.AlternateSpace == nil {
			return nil, _a.New("\u0061\u006ct\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069fi\u0065\u0064")
		}
		return _bgb.getInitialColor(_bcc.AlternateSpace)
	case *_ef.PdfColorspaceDeviceN:
		if _bcc.AlternateSpace == nil {
			return nil, _a.New("\u0061\u006ct\u0065\u0072\u006e\u0061\u0074\u0065\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069fi\u0065\u0064")
		}
		return _bgb.getInitialColor(_bcc.AlternateSpace)
	case *_ef.PdfColorspaceSpecialPattern:
		return nil, nil
	}
	_gc.Log.Debug("Un\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0066\u006f\u0072\u0020\u0075\u006e\u006b\u006e\u006fw\u006e \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065:\u0020\u0025T", _eff)
	return nil, _a.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061c\u0065")
}

// Add_re appends 're' operand to the content stream:
// Append a rectangle to the current path as a complete subpath, with lower left corner (x,y).
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_acd *ContentCreator) Add_re(x, y, width, height float64) *ContentCreator {
	_ecg := ContentStreamOperation{}
	_ecg.Operand = "\u0072\u0065"
	_ecg.Params = _bcd([]float64{x, y, width, height})
	_acd._ec = append(_acd._ec, &_ecg)
	return _acd
}

// Add_BMC appends 'BMC' operand to the content stream:
// Begins a marked-content sequence terminated by a balancing EMC operator.
// `tag` shall be a name object indicating the role or significance of
// the sequence.
//
// See section 14.6 "Marked Content" and Table 320 (p. 561 PDF32000_2008).
func (_cff *ContentCreator) Add_BMC(tag _gb.PdfObjectName) *ContentCreator {
	_bgd := ContentStreamOperation{}
	_bgd.Operand = "\u0042\u004d\u0043"
	_bgd.Params = _afec([]_gb.PdfObjectName{tag})
	_cff._ec = append(_cff._ec, &_bgd)
	return _cff
}

// Add_M adds 'M' operand to the content stream: Set the miter limit (graphics state).
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_bd *ContentCreator) Add_M(miterlimit float64) *ContentCreator {
	_gag := ContentStreamOperation{}
	_gag.Operand = "\u004d"
	_gag.Params = _bcd([]float64{miterlimit})
	_bd._ec = append(_bd._ec, &_gag)
	return _bd
}

// Add_S appends 'S' operand to the content stream: Stroke the path.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_bbg *ContentCreator) Add_S() *ContentCreator {
	_fae := ContentStreamOperation{}
	_fae.Operand = "\u0053"
	_bbg._ec = append(_bbg._ec, &_fae)
	return _bbg
}

// GetEncoder returns the encoder of the inline image.
func (_deb *ContentStreamInlineImage) GetEncoder() (_gb.StreamEncoder, error) { return _bgf(_deb) }
func _ccbe(_bgadd []_gb.PdfObjectString) []_gb.PdfObject {
	var _cfcc []_gb.PdfObject
	for _, _cbbbd := range _bgadd {
		_cfcc = append(_cfcc, _gb.MakeString(_cbbbd.Str()))
	}
	return _cfcc
}
func _gfg(_fca *ContentStreamInlineImage) (*_gb.MultiEncoder, error) {
	_bca := _gb.NewMultiEncoder()
	var _daac *_gb.PdfObjectDictionary
	var _cbfc []_gb.PdfObject
	if _ffd := _fca.DecodeParms; _ffd != nil {
		_gagd, _bcgb := _ffd.(*_gb.PdfObjectDictionary)
		if _bcgb {
			_daac = _gagd
		}
		_ece, _bee := _ffd.(*_gb.PdfObjectArray)
		if _bee {
			for _, _gcd := range _ece.Elements() {
				if _affa, _dbg := _gcd.(*_gb.PdfObjectDictionary); _dbg {
					_cbfc = append(_cbfc, _affa)
				} else {
					_cbfc = append(_cbfc, nil)
				}
			}
		}
	}
	_egb := _fca.Filter
	if _egb == nil {
		return nil, _d.Errorf("\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006d\u0069s\u0073\u0069\u006e\u0067")
	}
	_aaf, _caac := _egb.(*_gb.PdfObjectArray)
	if !_caac {
		return nil, _d.Errorf("m\u0075\u006c\u0074\u0069\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0063\u0061\u006e\u0020\u006f\u006el\u0079\u0020\u0062\u0065\u0020\u006d\u0061\u0064\u0065\u0020fr\u006f\u006d\u0020a\u0072r\u0061\u0079")
	}
	for _fea, _dgf := range _aaf.Elements() {
		_gda, _befa := _dgf.(*_gb.PdfObjectName)
		if !_befa {
			return nil, _d.Errorf("\u006d\u0075l\u0074\u0069\u0020\u0066i\u006c\u0074e\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0061 \u006e\u0061\u006d\u0065")
		}
		var _gga _gb.PdfObject
		if _daac != nil {
			_gga = _daac
		} else {
			if len(_cbfc) > 0 {
				if _fea >= len(_cbfc) {
					return nil, _d.Errorf("\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0073\u0020\u0069\u006e\u0020d\u0065\u0063\u006f\u0064\u0065\u0020\u0070\u0061\u0072\u0061\u006d\u0073\u0020a\u0072\u0072\u0061\u0079")
				}
				_gga = _cbfc[_fea]
			}
		}
		var _gcce *_gb.PdfObjectDictionary
		if _eef, _edc := _gga.(*_gb.PdfObjectDictionary); _edc {
			_gcce = _eef
		}
		if *_gda == _gb.StreamEncodingFilterNameFlate || *_gda == "\u0046\u006c" {
			_eebd, _cbe := _aca(_fca, _gcce)
			if _cbe != nil {
				return nil, _cbe
			}
			_bca.AddEncoder(_eebd)
		} else if *_gda == _gb.StreamEncodingFilterNameLZW {
			_dcag, _gaa := _caff(_fca, _gcce)
			if _gaa != nil {
				return nil, _gaa
			}
			_bca.AddEncoder(_dcag)
		} else if *_gda == _gb.StreamEncodingFilterNameASCIIHex {
			_bbf := _gb.NewASCIIHexEncoder()
			_bca.AddEncoder(_bbf)
		} else if *_gda == _gb.StreamEncodingFilterNameASCII85 || *_gda == "\u0041\u0038\u0035" {
			_bgfb := _gb.NewASCII85Encoder()
			_bca.AddEncoder(_bgfb)
		} else {
			_gc.Log.Error("U\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0066\u0069l\u0074\u0065\u0072\u0020\u0025\u0073", *_gda)
			return nil, _d.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u0066\u0069\u006c\u0074er \u0069n \u006d\u0075\u006c\u0074\u0069\u0020\u0066il\u0074\u0065\u0072\u0020\u0061\u0072\u0072a\u0079")
		}
	}
	return _bca, nil
}
func (_dbgf *ContentStreamParser) parseBool() (_gb.PdfObjectBool, error) {
	_beda, _cbbb := _dbgf._dcf.Peek(4)
	if _cbbb != nil {
		return _gb.PdfObjectBool(false), _cbbb
	}
	if (len(_beda) >= 4) && (string(_beda[:4]) == "\u0074\u0072\u0075\u0065") {
		_dbgf._dcf.Discard(4)
		return _gb.PdfObjectBool(true), nil
	}
	_beda, _cbbb = _dbgf._dcf.Peek(5)
	if _cbbb != nil {
		return _gb.PdfObjectBool(false), _cbbb
	}
	if (len(_beda) >= 5) && (string(_beda[:5]) == "\u0066\u0061\u006cs\u0065") {
		_dbgf._dcf.Discard(5)
		return _gb.PdfObjectBool(false), nil
	}
	return _gb.PdfObjectBool(false), _a.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}
func (_bagd *ContentStreamProcessor) handleCommand_cs(_gdd *ContentStreamOperation, _fdca *_ef.PdfPageResources) error {
	if len(_gdd.Params) < 1 {
		_gc.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0043\u0053\u0020\u0063\u006f\u006d\u006d\u0061\u006e\u0064\u002c\u0020s\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u006f\u0076\u0065\u0072")
		return _a.New("\u0074o\u006f \u0066\u0065\u0077\u0020\u0070a\u0072\u0061m\u0065\u0074\u0065\u0072\u0073")
	}
	if len(_gdd.Params) > 1 {
		_gc.Log.Debug("\u0043\u0053\u0020\u0063\u006f\u006d\u006d\u0061n\u0064\u0020\u0077it\u0068\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020-\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075i\u006e\u0067")
		return _a.New("\u0074\u006f\u006f\u0020ma\u006e\u0079\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	}
	_beb, _cea := _gdd.Params[0].(*_gb.PdfObjectName)
	if !_cea {
		_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020C\u0053\u0020\u0063o\u006d\u006d\u0061n\u0064\u0020w\u0069\u0074\u0068\u0020\u0069\u006ev\u0061li\u0064\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u002c\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u006f\u0076\u0065\u0072")
		return _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_feeb, _fedb := _bagd.getColorspace(string(*_beb), _fdca)
	if _fedb != nil {
		return _fedb
	}
	_bagd._efaee.ColorspaceNonStroking = _feeb
	_adfd, _fedb := _bagd.getInitialColor(_feeb)
	if _fedb != nil {
		return _fedb
	}
	_bagd._efaee.ColorNonStroking = _adfd
	return nil
}
func (_abc *ContentStreamParser) parseString() (*_gb.PdfObjectString, error) {
	_abc._dcf.ReadByte()
	var _edgg []byte
	_acaa := 1
	for {
		_cdg, _cbcg := _abc._dcf.Peek(1)
		if _cbcg != nil {
			return _gb.MakeString(string(_edgg)), _cbcg
		}
		if _cdg[0] == '\\' {
			_abc._dcf.ReadByte()
			_gecf, _ecbb := _abc._dcf.ReadByte()
			if _ecbb != nil {
				return _gb.MakeString(string(_edgg)), _ecbb
			}
			if _gb.IsOctalDigit(_gecf) {
				_fddf, _adbf := _abc._dcf.Peek(2)
				if _adbf != nil {
					return _gb.MakeString(string(_edgg)), _adbf
				}
				var _dfbd []byte
				_dfbd = append(_dfbd, _gecf)
				for _, _ffb := range _fddf {
					if _gb.IsOctalDigit(_ffb) {
						_dfbd = append(_dfbd, _ffb)
					} else {
						break
					}
				}
				_abc._dcf.Discard(len(_dfbd) - 1)
				_gc.Log.Trace("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022", _dfbd)
				_fdaa, _adbf := _cb.ParseUint(string(_dfbd), 8, 32)
				if _adbf != nil {
					return _gb.MakeString(string(_edgg)), _adbf
				}
				_edgg = append(_edgg, byte(_fdaa))
				continue
			}
			switch _gecf {
			case 'n':
				_edgg = append(_edgg, '\n')
			case 'r':
				_edgg = append(_edgg, '\r')
			case 't':
				_edgg = append(_edgg, '\t')
			case 'b':
				_edgg = append(_edgg, '\b')
			case 'f':
				_edgg = append(_edgg, '\f')
			case '(':
				_edgg = append(_edgg, '(')
			case ')':
				_edgg = append(_edgg, ')')
			case '\\':
				_edgg = append(_edgg, '\\')
			}
			continue
		} else if _cdg[0] == '(' {
			_acaa++
		} else if _cdg[0] == ')' {
			_acaa--
			if _acaa == 0 {
				_abc._dcf.ReadByte()
				break
			}
		}
		_ecea, _ := _abc._dcf.ReadByte()
		_edgg = append(_edgg, _ecea)
	}
	return _gb.MakeString(string(_edgg)), nil
}

// Add_Tr appends 'Tr' operand to the content stream:
// Set text rendering mode.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_fedf *ContentCreator) Add_Tr(render int64) *ContentCreator {
	_dfa := ContentStreamOperation{}
	_dfa.Operand = "\u0054\u0072"
	_dfa.Params = _afeb([]int64{render})
	_fedf._ec = append(_fedf._ec, &_dfa)
	return _fedf
}

// Add_w adds 'w' operand to the content stream, which sets the line width.
//
// See section 8.4.4 "Graphic State Operators" and Table 57 (pp. 135-136 PDF32000_2008).
func (_gff *ContentCreator) Add_w(lineWidth float64) *ContentCreator {
	_eda := ContentStreamOperation{}
	_eda.Operand = "\u0077"
	_eda.Params = _bcd([]float64{lineWidth})
	_gff._ec = append(_gff._ec, &_eda)
	return _gff
}
func (_fegf *ContentStreamProcessor) handleCommand_sc(_gfdc *ContentStreamOperation, _gdfg *_ef.PdfPageResources) error {
	_bfbd := _fegf._efaee.ColorspaceNonStroking
	if !_cfd(_bfbd) {
		if len(_gfdc.Params) != _bfbd.GetNumComponents() {
			_gc.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_gc.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_gfdc.Params), _bfbd)
			return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_gfbd, _bcbe := _bfbd.ColorFromPdfObjects(_gfdc.Params)
	if _bcbe != nil {
		return _bcbe
	}
	_fegf._efaee.ColorNonStroking = _gfbd
	return nil
}

// Add_G appends 'G' operand to the content stream:
// Set the stroking colorspace to DeviceGray and sets the gray level (0-1).
//
// See section 8.6.8 "Colour Operators" and Table 74 (p. 179-180 PDF32000_2008).
func (_fdg *ContentCreator) Add_G(gray float64) *ContentCreator {
	_gcgc := ContentStreamOperation{}
	_gcgc.Operand = "\u0047"
	_gcgc.Params = _bcd([]float64{gray})
	_fdg._ec = append(_fdg._ec, &_gcgc)
	return _fdg
}

// Bytes converts a set of content stream operations to a content stream byte presentation,
// i.e. the kind that can be stored as a PDF stream or string format.
func (_ad *ContentStreamOperations) Bytes() []byte {
	var _gbb _de.Buffer
	for _, _ba := range *_ad {
		if _ba == nil {
			continue
		}
		if _ba.Operand == "\u0042\u0049" {
			_gbb.WriteString(_ba.Operand + "\u000a")
			_gbb.WriteString(_ba.Params[0].WriteString())
		} else {
			for _, _gcga := range _ba.Params {
				_gbb.WriteString(_gcga.WriteString())
				_gbb.WriteString("\u0020")
			}
			_gbb.WriteString(_ba.Operand + "\u000a")
		}
	}
	return _gbb.Bytes()
}
func (_ada *ContentStreamProcessor) handleCommand_scn(_gedd *ContentStreamOperation, _gabg *_ef.PdfPageResources) error {
	_gagb := _ada._efaee.ColorspaceNonStroking
	if !_cfd(_gagb) {
		if len(_gedd.Params) != _gagb.GetNumComponents() {
			_gc.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0061\u0072\u0061m\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020S\u0043")
			_gc.Log.Debug("\u004e\u0075mb\u0065\u0072\u0020%\u0064\u0020\u006e\u006ft m\u0061tc\u0068\u0069\u006e\u0067\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u0020%\u0054", len(_gedd.Params), _gagb)
			return _a.New("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0065\u0074e\u0072\u0073")
		}
	}
	_aeef, _ffg := _gagb.ColorFromPdfObjects(_gedd.Params)
	if _ffg != nil {
		_gc.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c \u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0063o\u006co\u0072\u0020\u0066\u0072\u006f\u006d\u0020\u0070\u0061\u0072\u0061\u006d\u0073\u003a\u0020\u0025\u002b\u0076 \u0028\u0043\u0053\u0020\u0069\u0073\u0020\u0025\u002b\u0076\u0029", _gedd.Params, _gagb)
		return _ffg
	}
	_ada._efaee.ColorNonStroking = _aeef
	return nil
}

var (
	ErrInvalidOperand = _a.New("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")
)

// RotateDeg applies a rotation to the transformation matrix.
func (_bbb *ContentCreator) RotateDeg(angle float64) *ContentCreator {
	_afa := _ed.Cos(angle * _ed.Pi / 180.0)
	_gcf := _ed.Sin(angle * _ed.Pi / 180.0)
	_fa := -_ed.Sin(angle * _ed.Pi / 180.0)
	_fdab := _ed.Cos(angle * _ed.Pi / 180.0)
	return _bbb.Add_cm(_afa, _gcf, _fa, _fdab, 0, 0)
}

// Operations returns the list of operations.
func (_ea *ContentCreator) Operations() *ContentStreamOperations { return &_ea._ec }

// Pop pops and returns the topmost GraphicsState off the `gsStack`.
func (_fcba *GraphicStateStack) Pop() GraphicsState {
	_acc := (*_fcba)[len(*_fcba)-1]
	*_fcba = (*_fcba)[:len(*_fcba)-1]
	return _acc
}

// AddHandler adds a new ContentStreamProcessor `handler` of type `condition` for `operand`.
func (_edag *ContentStreamProcessor) AddHandler(condition HandlerConditionEnum, operand string, handler HandlerFunc) {
	_baf := handlerEntry{}
	_baf.Condition = condition
	_baf.Operand = operand
	_baf.Handler = handler
	_edag._dcb = append(_edag._dcb, _baf)
}

// GraphicsState is a basic graphics state implementation for PDF processing.
// Initially only implementing and tracking a portion of the information specified. Easy to add more.
type GraphicsState struct {
	ColorspaceStroking    _ef.PdfColorspace
	ColorspaceNonStroking _ef.PdfColorspace
	ColorStroking         _ef.PdfColor
	ColorNonStroking      _ef.PdfColor
	CTM                   _df.Matrix
}

// ExtractText parses and extracts all text data in content streams and returns as a string.
// Does not take into account Encoding table, the output is simply the character codes.
//
// Deprecated: More advanced text extraction is offered in package extractor with character encoding support.
func (_ca *ContentStreamParser) ExtractText() (string, error) {
	_fd, _cd := _ca.Parse()
	if _cd != nil {
		return "", _cd
	}
	_dg := false
	_cg, _bba := float64(-1), float64(-1)
	_gf := ""
	for _, _fg := range *_fd {
		if _fg.Operand == "\u0042\u0054" {
			_dg = true
		} else if _fg.Operand == "\u0045\u0054" {
			_dg = false
		}
		if _fg.Operand == "\u0054\u0064" || _fg.Operand == "\u0054\u0044" || _fg.Operand == "\u0054\u002a" {
			_gf += "\u000a"
		}
		if _fg.Operand == "\u0054\u006d" {
			if len(_fg.Params) != 6 {
				continue
			}
			_dee, _acg := _fg.Params[4].(*_gb.PdfObjectFloat)
			if !_acg {
				_caf, _dc := _fg.Params[4].(*_gb.PdfObjectInteger)
				if !_dc {
					continue
				}
				_dee = _gb.MakeFloat(float64(*_caf))
			}
			_gg, _acg := _fg.Params[5].(*_gb.PdfObjectFloat)
			if !_acg {
				_ce, _gfe := _fg.Params[5].(*_gb.PdfObjectInteger)
				if !_gfe {
					continue
				}
				_gg = _gb.MakeFloat(float64(*_ce))
			}
			if _bba == -1 {
				_bba = float64(*_gg)
			} else if _bba > float64(*_gg) {
				_gf += "\u000a"
				_cg = float64(*_dee)
				_bba = float64(*_gg)
				continue
			}
			if _cg == -1 {
				_cg = float64(*_dee)
			} else if _cg < float64(*_dee) {
				_gf += "\u0009"
				_cg = float64(*_dee)
			}
		}
		if _dg && _fg.Operand == "\u0054\u004a" {
			if len(_fg.Params) < 1 {
				continue
			}
			_da, _ab := _fg.Params[0].(*_gb.PdfObjectArray)
			if !_ab {
				return "", _d.Errorf("\u0069\u006ev\u0061\u006c\u0069\u0064 \u0070\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0020\u0074y\u0070\u0065\u002c\u0020\u006e\u006f\u0020\u0061\u0072\u0072\u0061\u0079 \u0028\u0025\u0054\u0029", _fg.Params[0])
			}
			for _, _adg := range _da.Elements() {
				switch _cbf := _adg.(type) {
				case *_gb.PdfObjectString:
					_gf += _cbf.Str()
				case *_gb.PdfObjectFloat:
					if *_cbf < -100 {
						_gf += "\u0020"
					}
				case *_gb.PdfObjectInteger:
					if *_cbf < -100 {
						_gf += "\u0020"
					}
				}
			}
		} else if _dg && _fg.Operand == "\u0054\u006a" {
			if len(_fg.Params) < 1 {
				continue
			}
			_ede, _bg := _fg.Params[0].(*_gb.PdfObjectString)
			if !_bg {
				return "", _d.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072\u0020\u0074\u0079p\u0065\u002c\u0020\u006e\u006f\u0074\u0020\u0073\u0074\u0072\u0069\u006e\u0067 \u0028\u0025\u0054\u0029", _fg.Params[0])
			}
			_gf += _ede.Str()
		}
	}
	return _gf, nil
}

// NewContentStreamParser creates a new instance of the content stream parser from an input content
// stream string.
func NewContentStreamParser(contentStr string) *ContentStreamParser {
	_dac := ContentStreamParser{}
	_ggec := _de.NewBufferString(contentStr + "\u000a")
	_dac._dcf = _b.NewReader(_ggec)
	return &_dac
}

const (
	HandlerConditionEnumOperand HandlerConditionEnum = iota
	HandlerConditionEnumAllOperands
)

// Add_s appends 's' operand to the content stream: Close and stroke the path.
//
// See section 8.5.3 "Path Painting Operators" and Table 60 (p. 143 PDF32000_2008).
func (_ee *ContentCreator) Add_s() *ContentCreator {
	_cbfb := ContentStreamOperation{}
	_cbfb.Operand = "\u0073"
	_ee._ec = append(_ee._ec, &_cbfb)
	return _ee
}

// ContentStreamOperation represents an operation in PDF contentstream which consists of
// an operand and parameters.
type ContentStreamOperation struct {
	Params  []_gb.PdfObject
	Operand string
}

// Add_Tz appends 'Tz' operand to the content stream:
// Set horizontal scaling.
//
// See section 9.3 "Text State Parameters and Operators" and
// Table 105 (pp. 251-252 PDF32000_2008).
func (_cdc *ContentCreator) Add_Tz(scale float64) *ContentCreator {
	_acgg := ContentStreamOperation{}
	_acgg.Operand = "\u0054\u007a"
	_acgg.Params = _bcd([]float64{scale})
	_cdc._ec = append(_cdc._ec, &_acgg)
	return _cdc
}
func (_gaac *ContentStreamProcessor) getColorspace(_edcb string, _bbeae *_ef.PdfPageResources) (_ef.PdfColorspace, error) {
	switch _edcb {
	case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
		return _ef.NewPdfColorspaceDeviceGray(), nil
	case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
		return _ef.NewPdfColorspaceDeviceRGB(), nil
	case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
		return _ef.NewPdfColorspaceDeviceCMYK(), nil
	case "\u0050a\u0074\u0074\u0065\u0072\u006e":
		return _ef.NewPdfColorspaceSpecialPattern(), nil
	}
	_dda, _cfef := _bbeae.GetColorspaceByName(_gb.PdfObjectName(_edcb))
	if _cfef {
		return _dda, nil
	}
	switch _edcb {
	case "\u0043a\u006c\u0047\u0072\u0061\u0079":
		return _ef.NewPdfColorspaceCalGray(), nil
	case "\u0043\u0061\u006c\u0052\u0047\u0042":
		return _ef.NewPdfColorspaceCalRGB(), nil
	case "\u004c\u0061\u0062":
		return _ef.NewPdfColorspaceLab(), nil
	}
	_gc.Log.Debug("\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0063\u006f\u006c\u006f\u0072\u0073\u0070a\u0063e\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074\u0065\u0064\u003a\u0020\u0025\u0073", _edcb)
	return nil, _d.Errorf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0063\u006f\u006c\u006f\u0072\u0073\u0070\u0061\u0063\u0065:\u0020\u0025\u0073", _edcb)
}

// Add_h appends 'h' operand to the content stream:
// Close the current subpath by adding a line between the current position and the starting position.
//
// See section 8.5.2 "Path Construction Operators" and Table 59 (pp. 140-141 PDF32000_2008).
func (_agfc *ContentCreator) Add_h() *ContentCreator {
	_cgg := ContentStreamOperation{}
	_cgg.Operand = "\u0068"
	_agfc._ec = append(_agfc._ec, &_cgg)
	return _agfc
}

var _dbeb = map[string]struct{}{"\u0062": struct{}{}, "\u0042": struct{}{}, "\u0062\u002a": struct{}{}, "\u0042\u002a": struct{}{}, "\u0042\u0044\u0043": struct{}{}, "\u0042\u0049": struct{}{}, "\u0042\u004d\u0043": struct{}{}, "\u0042\u0054": struct{}{}, "\u0042\u0058": struct{}{}, "\u0063": struct{}{}, "\u0063\u006d": struct{}{}, "\u0043\u0053": struct{}{}, "\u0063\u0073": struct{}{}, "\u0064": struct{}{}, "\u0064\u0030": struct{}{}, "\u0064\u0031": struct{}{}, "\u0044\u006f": struct{}{}, "\u0044\u0050": struct{}{}, "\u0045\u0049": struct{}{}, "\u0045\u004d\u0043": struct{}{}, "\u0045\u0054": struct{}{}, "\u0045\u0058": struct{}{}, "\u0066": struct{}{}, "\u0046": struct{}{}, "\u0066\u002a": struct{}{}, "\u0047": struct{}{}, "\u0067": struct{}{}, "\u0067\u0073": struct{}{}, "\u0068": struct{}{}, "\u0069": struct{}{}, "\u0049\u0044": struct{}{}, "\u006a": struct{}{}, "\u004a": struct{}{}, "\u004b": struct{}{}, "\u006b": struct{}{}, "\u006c": struct{}{}, "\u006d": struct{}{}, "\u004d": struct{}{}, "\u004d\u0050": struct{}{}, "\u006e": struct{}{}, "\u0071": struct{}{}, "\u0051": struct{}{}, "\u0072\u0065": struct{}{}, "\u0052\u0047": struct{}{}, "\u0072\u0067": struct{}{}, "\u0072\u0069": struct{}{}, "\u0073": struct{}{}, "\u0053": struct{}{}, "\u0053\u0043": struct{}{}, "\u0073\u0063": struct{}{}, "\u0053\u0043\u004e": struct{}{}, "\u0073\u0063\u006e": struct{}{}, "\u0073\u0068": struct{}{}, "\u0054\u002a": struct{}{}, "\u0054\u0063": struct{}{}, "\u0054\u0064": struct{}{}, "\u0054\u0044": struct{}{}, "\u0054\u0066": struct{}{}, "\u0054\u006a": struct{}{}, "\u0054\u004a": struct{}{}, "\u0054\u004c": struct{}{}, "\u0054\u006d": struct{}{}, "\u0054\u0072": struct{}{}, "\u0054\u0073": struct{}{}, "\u0054\u0077": struct{}{}, "\u0054\u007a": struct{}{}, "\u0076": struct{}{}, "\u0077": struct{}{}, "\u0057": struct{}{}, "\u0057\u002a": struct{}{}, "\u0079": struct{}{}, "\u0027": struct{}{}, "\u0022": struct{}{}}

func (_dba *ContentStreamParser) parseNumber() (_gb.PdfObject, error) {
	return _gb.ParseNumber(_dba._dcf)
}
