package annotator

import (
	_c "bytes"
	_bd "errors"
	_e "image"
	_cf "math"
	_ge "strings"
	_g "unicode"

	_a "bitbucket.org/shenghui0779/gopdf/common"
	_b "bitbucket.org/shenghui0779/gopdf/contentstream"
	_af "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_dd "bitbucket.org/shenghui0779/gopdf/core"
	_d "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_fa "bitbucket.org/shenghui0779/gopdf/model"
)

func _dace(_egd []*SignatureLine, _bffd *SignatureFieldOpts) (*_dd.PdfObjectDictionary, error) {
	if _bffd == nil {
		_bffd = NewSignatureFieldOpts()
	}
	var _bbc error
	var _cdcb *_dd.PdfObjectName
	_aegd := _bffd.Font
	if _aegd != nil {
		_cae, _ := _aegd.GetFontDescriptor()
		if _cae != nil {
			if _eba, _dfagb := _cae.FontName.(*_dd.PdfObjectName); _dfagb {
				_cdcb = _eba
			}
		}
		if _cdcb == nil {
			_cdcb = _dd.MakeName("\u0046\u006f\u006et\u0031")
		}
	} else {
		if _aegd, _bbc = _fa.NewStandard14Font("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a"); _bbc != nil {
			return nil, _bbc
		}
		_cdcb = _dd.MakeName("\u0048\u0065\u006c\u0076")
	}
	_bdaba := _bffd.FontSize
	if _bdaba <= 0 {
		_bdaba = 10
	}
	if _bffd.LineHeight <= 0 {
		_bffd.LineHeight = 1
	}
	_efbb := _bffd.LineHeight * _bdaba
	_agda, _eebf := _aegd.GetRuneMetrics(' ')
	if !_eebf {
		return nil, _bd.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
	}
	_fedb := _agda.Wx
	var _bbca float64
	var _fde []string
	for _, _aebd := range _egd {
		if _aebd.Text == "" {
			continue
		}
		_eced := _aebd.Text
		if _aebd.Desc != "" {
			_eced = _aebd.Desc + "\u003a\u0020" + _eced
		}
		_fde = append(_fde, _eced)
		var _dfab float64
		for _, _ebgda := range _eced {
			_aedf, _fdf := _aegd.GetRuneMetrics(_ebgda)
			if !_fdf {
				continue
			}
			_dfab += _aedf.Wx
		}
		if _dfab > _bbca {
			_bbca = _dfab
		}
	}
	_bbca = _bbca * _bdaba / 1000.0
	_fgaa := float64(len(_fde)) * _efbb
	_ggbcf := _bffd.Image != nil
	_abdg := _bffd.Rect
	if _abdg == nil {
		_abdg = []float64{0, 0, _bbca, _fgaa}
		if _ggbcf {
			_abdg[2] = _bbca * 2
			_abdg[3] = _fgaa * 2
		}
		_bffd.Rect = _abdg
	}
	_aabf := _abdg[2] - _abdg[0]
	_bagf := _abdg[3] - _abdg[1]
	_bac, _aecf := _abdg, _abdg
	var _dba, _cgb float64
	if _ggbcf && len(_fde) > 0 {
		if _bffd.ImagePosition <= SignatureImageRight {
			_dddgc := []float64{_abdg[0], _abdg[1], _abdg[0] + (_aabf / 2), _abdg[3]}
			_gbfe := []float64{_abdg[0] + (_aabf / 2), _abdg[1], _abdg[2], _abdg[3]}
			if _bffd.ImagePosition == SignatureImageLeft {
				_bac, _aecf = _dddgc, _gbfe
			} else {
				_bac, _aecf = _gbfe, _dddgc
			}
		} else {
			_cdb := []float64{_abdg[0], _abdg[1], _abdg[2], _abdg[1] + (_bagf / 2)}
			_fabf := []float64{_abdg[0], _abdg[1] + (_bagf / 2), _abdg[2], _abdg[3]}
			if _bffd.ImagePosition == SignatureImageTop {
				_bac, _aecf = _fabf, _cdb
			} else {
				_bac, _aecf = _cdb, _fabf
			}
		}
	}
	_dba = _aecf[2] - _aecf[0]
	_cgb = _aecf[3] - _aecf[1]
	var _cfdg float64
	if _bffd.AutoSize {
		if _bbca > _dba || _fgaa > _cgb {
			_gdg := _cf.Min(_dba/_bbca, _cgb/_fgaa)
			_bdaba *= _gdg
		}
		_efbb = _bffd.LineHeight * _bdaba
		_cfdg += (_cgb - float64(len(_fde))*_efbb) / 2
	}
	_cfcd := _b.NewContentCreator()
	_gaf := _fa.NewPdfPageResources()
	_gaf.SetFontByName(*_cdcb, _aegd.ToPdfObject())
	if _bffd.BorderSize <= 0 {
		_bffd.BorderSize = 0
		_bffd.BorderColor = _fa.NewPdfColorDeviceGray(1)
	}
	_cfcd.Add_q()
	if _bffd.FillColor != nil {
		_cfcd.SetNonStrokingColor(_bffd.FillColor)
	}
	if _bffd.BorderColor != nil {
		_cfcd.SetStrokingColor(_bffd.BorderColor)
	}
	_cfcd.Add_w(_bffd.BorderSize).Add_re(_abdg[0], _abdg[1], _aabf, _bagf)
	if _bffd.FillColor != nil && _bffd.BorderColor != nil {
		_cfcd.Add_B()
	} else if _bffd.FillColor != nil {
		_cfcd.Add_f()
	} else if _bffd.BorderColor != nil {
		_cfcd.Add_S()
	}
	_cfcd.Add_Q()
	if _bffd.WatermarkImage != nil {
		_abf := []float64{_abdg[0], _abdg[1], _abdg[2], _abdg[3]}
		_fbcg, _cacc, _bcd := _gea(_bffd.WatermarkImage, "\u0049\u006d\u0061\u0067\u0065\u0057\u0061\u0074\u0065r\u006d\u0061\u0072\u006b", _bffd, _abf, _cfcd)
		if _bcd != nil {
			return nil, _bcd
		}
		_gaf.SetXObjectImageByName(*_fbcg, _cacc)
	}
	_cfcd.Add_q()
	_cfcd.Translate(_aecf[0], _aecf[3]-_efbb-_cfdg)
	_cfcd.Add_BT()
	_ade := _aegd.Encoder()
	for _, _gdeg := range _fde {
		var _ffge []byte
		for _, _aga := range _gdeg {
			if _g.IsSpace(_aga) {
				if len(_ffge) > 0 {
					_cfcd.SetNonStrokingColor(_bffd.TextColor).Add_Tf(*_cdcb, _bdaba).Add_TL(_efbb).Add_TJ([]_dd.PdfObject{_dd.MakeStringFromBytes(_ffge)}...)
					_ffge = nil
				}
				_cfcd.Add_Tf(*_cdcb, _bdaba).Add_TL(_efbb).Add_TJ([]_dd.PdfObject{_dd.MakeFloat(-_fedb)}...)
			} else {
				_ffge = append(_ffge, _ade.Encode(string(_aga))...)
			}
		}
		if len(_ffge) > 0 {
			_cfcd.SetNonStrokingColor(_bffd.TextColor).Add_Tf(*_cdcb, _bdaba).Add_TL(_efbb).Add_TJ([]_dd.PdfObject{_dd.MakeStringFromBytes(_ffge)}...)
		}
		_cfcd.Add_Td(0, -_efbb)
	}
	_cfcd.Add_ET()
	_cfcd.Add_Q()
	if _ggbcf {
		_agdg, _eabd, _fdfd := _gea(_bffd.Image, "\u0049\u006d\u0061\u0067\u0065\u0053\u0069\u0067\u006ea\u0074\u0075\u0072\u0065", _bffd, _bac, _cfcd)
		if _fdfd != nil {
			return nil, _fdfd
		}
		_gaf.SetXObjectImageByName(*_agdg, _eabd)
	}
	_ffe := _fa.NewXObjectForm()
	_ffe.Resources = _gaf
	_ffe.BBox = _dd.MakeArrayFromFloats(_abdg)
	_ffe.SetContentStream(_cfcd.Bytes(), _afge())
	_dgge := _dd.MakeDict()
	_dgge.Set("\u004e", _ffe.ToPdfObject())
	return _dgge, nil
}

// NewSignatureLine returns a new signature line displayed as a part of the
// signature field appearance.
func NewSignatureLine(desc, text string) *SignatureLine {
	return &SignatureLine{Desc: desc, Text: text}
}

func _bga(_eeg *InkAnnotationDef) (*_dd.PdfObjectDictionary, *_fa.PdfRectangle, error) {
	_gbcf := _fa.NewXObjectForm()
	_bbea, _fbed, _ccba := _aebe(_eeg)
	if _ccba != nil {
		return nil, nil, _ccba
	}
	_ccba = _gbcf.SetContentStream(_bbea, nil)
	if _ccba != nil {
		return nil, nil, _ccba
	}
	_gbcf.BBox = _fbed.ToPdfObject()
	_gbcf.Resources = _fa.NewPdfPageResources()
	_gbcf.Resources.ProcSet = _dd.MakeArray(_dd.MakeName("\u0050\u0044\u0046"))
	_cdf := _dd.MakeDict()
	_cdf.Set("\u004e", _gbcf.ToPdfObject())
	return _cdf, _fbed, nil
}

func _dgcg(_gacd RectangleAnnotationDef, _deff string) ([]byte, *_fa.PdfRectangle, *_fa.PdfRectangle, error) {
	_bbccg := _af.Rectangle{X: 0, Y: 0, Width: _gacd.Width, Height: _gacd.Height, FillEnabled: _gacd.FillEnabled, FillColor: _gacd.FillColor, BorderEnabled: _gacd.BorderEnabled, BorderWidth: 2 * _gacd.BorderWidth, BorderColor: _gacd.BorderColor, Opacity: _gacd.Opacity}
	_ddfef, _edfb, _fbbaa := _bbccg.Draw(_deff)
	if _fbbaa != nil {
		return nil, nil, nil, _fbbaa
	}
	_dbaed := &_fa.PdfRectangle{}
	_dbaed.Llx = _gacd.X + _edfb.Llx
	_dbaed.Lly = _gacd.Y + _edfb.Lly
	_dbaed.Urx = _gacd.X + _edfb.Urx
	_dbaed.Ury = _gacd.Y + _edfb.Ury
	return _ddfef, _edfb, _dbaed, nil
}

// InkAnnotationDef holds base information for constructing an ink annotation.
type InkAnnotationDef struct {
	// Paths is the array of stroked paths which compose the annotation.
	Paths []_af.Path

	// Color is the color of the line. Default to black.
	Color *_fa.PdfColorDeviceRGB

	// LineWidth is the width of the line.
	LineWidth float64
}

// FormResetActionOptions holds options for creating a form reset button.
type FormResetActionOptions struct {
	// Rectangle holds the button position, size, and color.
	Rectangle _af.Rectangle

	// Label specifies the text that would be displayed on the button.
	Label string

	// LabelColor specifies the button label color.
	LabelColor _fa.PdfColor

	// Font specifies a font used for rendering the button label.
	// When omitted it will fallback to use a Helvetica font.
	Font *_fa.PdfFont

	// FontSize specifies the font size used in rendering the button label.
	// The default font size is 12pt.
	FontSize *float64

	// Fields specifies list of fields that could be resetted.
	// This list may contain indirect object to fields or field names.
	Fields *_dd.PdfObjectArray

	// IsExclusionList specifies that the fields in the `Fields` array would be excluded form reset process.
	IsExclusionList bool
}

// FieldAppearance implements interface model.FieldAppearanceGenerator and generates appearance streams
// for fields taking into account what value is in the field. A common use case is for generating the
// appearance stream prior to flattening fields.
//
// If `OnlyIfMissing` is true, the field appearance is generated only for fields that do not have an
// appearance stream specified.
// If `RegenerateTextFields` is true, all text fields are regenerated (even if OnlyIfMissing is true).
type FieldAppearance struct {
	OnlyIfMissing        bool
	RegenerateTextFields bool
	_eb                  *AppearanceStyle
}

// AppearanceStyle defines style parameters for appearance stream generation.
type AppearanceStyle struct {
	// How much of Rect height to fill when autosizing text.
	AutoFontSizeFraction float64

	// CheckmarkRune is a rune used for check mark in checkboxes (for ZapfDingbats font).
	CheckmarkRune rune
	BorderSize    float64
	BorderColor   _fa.PdfColor
	FillColor     _fa.PdfColor

	// Multiplier for lineheight for multi line text.
	MultilineLineHeight   float64
	MultilineVAlignMiddle bool

	// Visual guide checking alignment of field contents (debugging).
	DrawAlignmentReticle bool

	// Allow field MK appearance characteristics to override style settings.
	AllowMK bool

	// Fonts holds appearance styles for fonts.
	Fonts *AppearanceFontStyle

	// MarginLeft represents the amount of space to leave on the left side of
	// the form field bounding box when generating appearances (default: 2.0).
	MarginLeft *float64
}

// NewFormSubmitButtonField would create a submit button in specified page according to the parameter in `FormSubmitActionOptions`.
func NewFormSubmitButtonField(page *_fa.PdfPage, opt FormSubmitActionOptions) (*_fa.PdfFieldButton, error) {
	_ccad := int64(_gbed)
	if opt.IsExclusionList {
		_ccad |= _bfee
	}
	if opt.IncludeEmptyFields {
		_ccad |= _fgb
	}
	if opt.SubmitAsPDF {
		_ccad |= _dbe
	}
	_fgge := _fa.NewPdfActionSubmitForm()
	_fgge.Flags = _dd.MakeInteger(_ccad)
	_fgge.F = _fa.NewPdfFilespec()
	if opt.Fields != nil {
		_fgge.Fields = opt.Fields
	}
	_fgge.F.F = _dd.MakeString(opt.Url)
	_fgge.F.FS = _dd.MakeName("\u0055\u0052\u004c")
	_dfce, _ggcg := _ega(page, opt.Rectangle, "\u0062t\u006e\u0053\u0075\u0062\u006d\u0069t", opt.Label, opt.LabelColor, opt.Font, opt.FontSize, _fgge.ToPdfObject())
	if _ggcg != nil {
		return nil, _ggcg
	}
	return _dfce, nil
}

func _fga(_fbf *_fa.PdfAnnotationWidget, _bde *_fa.PdfFieldButton, _fed *_fa.PdfPageResources, _cccc AppearanceStyle) (*_dd.PdfObjectDictionary, error) {
	_bcc, _dabb := _dd.GetArray(_fbf.Rect)
	if !_dabb {
		return nil, _bd.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_daa, _aae := _fa.NewPdfRectangle(*_bcc)
	if _aae != nil {
		return nil, _aae
	}
	_dea, _fbce := _daa.Width(), _daa.Height()
	_cga, _gff := _dea, _fbce
	_a.Log.Debug("\u0043\u0068\u0065\u0063kb\u006f\u0078\u002c\u0020\u0077\u0061\u0020\u0042\u0053\u003a\u0020\u0025\u0076", _fbf.BS)
	_ccg, _aae := _fa.NewStandard14Font("\u005a\u0061\u0070f\u0044\u0069\u006e\u0067\u0062\u0061\u0074\u0073")
	if _aae != nil {
		return nil, _aae
	}
	_dag, _afa := _dd.GetDict(_fbf.MK)
	if _afa {
		_cee, _ := _dd.GetDict(_fbf.BS)
		_gbf := _cccc.applyAppearanceCharacteristics(_dag, _cee, _ccg)
		if _gbf != nil {
			return nil, _gbf
		}
	}
	_dfg := _fa.NewXObjectForm()
	{
		_fbd := _b.NewContentCreator()
		if _cccc.BorderSize > 0 {
			_afdc(_fbd, _cccc, _dea, _fbce)
		}
		if _cccc.DrawAlignmentReticle {
			_ced := _cccc
			_ced.BorderSize = 0.2
			_cece(_fbd, _ced, _dea, _fbce)
		}
		_dea, _fbce = _cccc.applyRotation(_dag, _dea, _fbce, _fbd)
		_affa := _cccc.AutoFontSizeFraction * _fbce
		_decf, _dfc := _ccg.GetRuneMetrics(_cccc.CheckmarkRune)
		if !_dfc {
			return nil, _bd.New("\u0067l\u0079p\u0068\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
		}
		_efa := _ccg.Encoder()
		_agdd := _efa.Encode(string(_cccc.CheckmarkRune))
		_dcga := _decf.Wx * _affa / 1000.0
		_cag := 705.0
		_cbd := _cag / 1000.0 * _affa
		_ceg := _aeb
		if _cccc.MarginLeft != nil {
			_ceg = *_cccc.MarginLeft
		}
		_feca := 1.0
		if _dcga < _dea {
			_ceg = (_dea - _dcga) / 2.0
		}
		if _cbd < _fbce {
			_feca = (_fbce - _cbd) / 2.0
		}
		_fbd.Add_q().Add_g(0).Add_BT().Add_Tf("\u005a\u0061\u0044\u0062", _affa).Add_Td(_ceg, _feca).Add_Tj(*_dd.MakeStringFromBytes(_agdd)).Add_ET().Add_Q()
		_dfg.Resources = _fa.NewPdfPageResources()
		_dfg.Resources.SetFontByName("\u005a\u0061\u0044\u0062", _ccg.ToPdfObject())
		_dfg.BBox = _dd.MakeArrayFromFloats([]float64{0, 0, _cga, _gff})
		_dfg.SetContentStream(_fbd.Bytes(), _afge())
	}
	_bgd := _fa.NewXObjectForm()
	{
		_fdg := _b.NewContentCreator()
		if _cccc.BorderSize > 0 {
			_afdc(_fdg, _cccc, _dea, _fbce)
		}
		_bgd.BBox = _dd.MakeArrayFromFloats([]float64{0, 0, _cga, _gff})
		_bgd.SetContentStream(_fdg.Bytes(), _afge())
	}
	_ebgd := _dd.PdfObjectName("\u0059\u0065\u0073")
	_agf, _afa := _dd.GetDict(_fbf.PdfAnnotation.AP)
	if _afa && _agf != nil {
		_def := _dd.TraceToDirectObject(_agf.Get("\u004e"))
		switch _ggd := _def.(type) {
		case *_dd.PdfObjectDictionary:
			_fbdg := _ggd.Keys()
			for _, _ccd := range _fbdg {
				if _ccd != "\u004f\u0066\u0066" {
					_ebgd = _ccd
				}
			}
		}
	}
	_deaf := _dd.MakeDict()
	_deaf.Set("\u004f\u0066\u0066", _bgd.ToPdfObject())
	_deaf.Set(_ebgd, _dfg.ToPdfObject())
	_baf := _dd.MakeDict()
	_baf.Set("\u004e", _deaf)
	return _baf, nil
}

// RectangleAnnotationDef is a rectangle defined with a specified Width and Height and a lower left corner at (X,Y).
// The rectangle can optionally have a border and a filling color.
// The Width/Height includes the border (if any specified).
type RectangleAnnotationDef struct {
	X             float64
	Y             float64
	Width         float64
	Height        float64
	FillEnabled   bool
	FillColor     *_fa.PdfColorDeviceRGB
	BorderEnabled bool
	BorderWidth   float64
	BorderColor   *_fa.PdfColorDeviceRGB
	Opacity       float64
}

// NewSignatureFieldOpts returns a new initialized instance of options
// used to generate a signature appearance.
func NewSignatureFieldOpts() *SignatureFieldOpts {
	return &SignatureFieldOpts{Font: _fa.DefaultFont(), FontSize: 10, LineHeight: 1, AutoSize: true, TextColor: _fa.NewPdfColorDeviceGray(0), BorderColor: _fa.NewPdfColorDeviceGray(0), FillColor: _fa.NewPdfColorDeviceGray(1), Encoder: _dd.NewFlateEncoder(), ImagePosition: SignatureImageLeft}
}

func _beed(_bdb *_fa.PdfField) string {
	if _bdb == nil {
		return ""
	}
	_cfa, _fac := _bdb.GetContext().(*_fa.PdfFieldText)
	if !_fac {
		return _beed(_bdb.Parent)
	}
	if _cfa.DA != nil {
		return _cfa.DA.Str()
	}
	return _beed(_cfa.Parent)
}

// AppearanceFontStyle defines font style characteristics for form fields,
// used in the filling/flattening process.
type AppearanceFontStyle struct {
	// Fallback represents a global font fallback, used for fields which do
	// not specify a font in their default appearance (DA). The fallback is
	// also used if there is a font specified in the DA, but it is not
	// found in the AcroForm resources (DR).
	Fallback *AppearanceFont

	// FallbackSize represents a global font size fallback used for fields
	// which do not specify a font size in their default appearance (DA).
	// The fallback size is applied only if its value is larger than zero.
	FallbackSize float64

	// FieldFallbacks defines font fallbacks for specific fields. The map keys
	// represent the names of the fields (which can be specified by their
	// partial or full names). Specific field fallback fonts take precedence
	// over the global font fallback.
	FieldFallbacks map[string]*AppearanceFont

	// ForceReplace forces the replacement of fonts in the filling/flattening
	// process, even if the default appearance (DA) specifies a valid font.
	// If no fallback font is provided, setting this field has no effect.
	ForceReplace bool
}

// NewFormResetButtonField would create a reset button in specified page according to the parameter in `FormResetActionOptions`.
func NewFormResetButtonField(page *_fa.PdfPage, opt FormResetActionOptions) (*_fa.PdfFieldButton, error) {
	_egcg := _fa.NewPdfActionResetForm()
	_egcg.Fields = opt.Fields
	_egcg.Flags = _dd.MakeInteger(0)
	if opt.IsExclusionList {
		_egcg.Flags = _dd.MakeInteger(1)
	}
	_dceg, _afbd := _ega(page, opt.Rectangle, "\u0062\u0074\u006e\u0052\u0065\u0073\u0065\u0074", opt.Label, opt.LabelColor, opt.Font, opt.FontSize, _egcg.ToPdfObject())
	if _afbd != nil {
		return nil, _afbd
	}
	return _dceg, nil
}

func (_feeg *AppearanceStyle) applyAppearanceCharacteristics(_cef *_dd.PdfObjectDictionary, _fbda *_dd.PdfObjectDictionary, _ad *_fa.PdfFont) error {
	if !_feeg.AllowMK {
		return nil
	}
	if CA, _geeg := _dd.GetString(_cef.Get("\u0043\u0041")); _geeg && _ad != nil {
		_fgcc := CA.Bytes()
		if len(_fgcc) != 0 {
			_bfcd := []rune(_ad.Encoder().Decode(_fgcc))
			if len(_bfcd) == 1 {
				_feeg.CheckmarkRune = _bfcd[0]
			}
		}
	}
	if BC, _ggg := _dd.GetArray(_cef.Get("\u0042\u0043")); _ggg {
		_eca, _fdba := BC.ToFloat64Array()
		if _fdba != nil {
			return _fdba
		}
		switch len(_eca) {
		case 1:
			_feeg.BorderColor = _fa.NewPdfColorDeviceGray(_eca[0])
		case 3:
			_feeg.BorderColor = _fa.NewPdfColorDeviceRGB(_eca[0], _eca[1], _eca[2])
		case 4:
			_feeg.BorderColor = _fa.NewPdfColorDeviceCMYK(_eca[0], _eca[1], _eca[2], _eca[3])
		default:
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0042\u0043\u0020\u002d\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0028\u0025\u0064)", len(_eca))
		}
		if _fbda != nil {
			if _agfc, _ggbc := _dd.GetNumberAsFloat(_fbda.Get("\u0057")); _ggbc == nil {
				_feeg.BorderSize = _agfc
			}
		}
	}
	if BG, _bfbf := _dd.GetArray(_cef.Get("\u0042\u0047")); _bfbf {
		_ggge, _dcfe := BG.ToFloat64Array()
		if _dcfe != nil {
			return _dcfe
		}
		switch len(_ggge) {
		case 1:
			_feeg.FillColor = _fa.NewPdfColorDeviceGray(_ggge[0])
		case 3:
			_feeg.FillColor = _fa.NewPdfColorDeviceRGB(_ggge[0], _ggge[1], _ggge[2])
		case 4:
			_feeg.FillColor = _fa.NewPdfColorDeviceCMYK(_ggge[0], _ggge[1], _ggge[2], _ggge[3])
		default:
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0042\u0047\u0020\u002d\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0028\u0025\u0064)", len(_ggge))
		}
	}
	return nil
}

func _gfbg(_gbac _dd.PdfObject, _dcc *_fa.PdfPageResources) (*_dd.PdfObjectName, float64, bool) {
	var (
		_ddfg *_dd.PdfObjectName
		_cfgf float64
		_dfaa bool
	)
	if _fbec, _feee := _dd.GetDict(_gbac); _feee && _fbec != nil {
		_bagg := _dd.TraceToDirectObject(_fbec.Get("\u004e"))
		switch _bca := _bagg.(type) {
		case *_dd.PdfObjectStream:
			_bbgb, _eec := _dd.DecodeStream(_bca)
			if _eec != nil {
				_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0063\u006f\u006e\u0074e\u006e\u0074\u0020\u0073\u0074r\u0065\u0061m\u003a\u0020\u0025\u0076", _eec.Error())
				return nil, 0, false
			}
			_gcf, _eec := _b.NewContentStreamParser(string(_bbgb)).Parse()
			if _eec != nil {
				_a.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0075n\u0061\u0062l\u0065\u0020\u0070\u0061\u0072\u0073\u0065\u0020c\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061m\u003a\u0020\u0025\u0076", _eec.Error())
				return nil, 0, false
			}
			_cea := _b.NewContentStreamProcessor(*_gcf)
			_cea.AddHandler(_b.HandlerConditionEnumOperand, "\u0054\u0066", func(_deg *_b.ContentStreamOperation, _cfab _b.GraphicsState, _bcbee *_fa.PdfPageResources) error {
				if len(_deg.Params) == 2 {
					if _ebge, _dfgaf := _dd.GetName(_deg.Params[0]); _dfgaf {
						_ddfg = _ebge
					}
					if _ccdb, _fagf := _dd.GetNumberAsFloat(_deg.Params[1]); _fagf == nil {
						_cfgf = _ccdb
					}
					_dfaa = true
					return _b.ErrEarlyExit
				}
				return nil
			})
			_cea.Process(_dcc)
			return _ddfg, _cfgf, _dfaa
		}
	}
	return nil, 0, false
}

// ImageFieldAppearance implements interface model.FieldAppearanceGenerator and generates appearance streams
// for attaching an image to a button field.
type ImageFieldAppearance struct {
	OnlyIfMissing bool
	_aaf          *AppearanceStyle
}

// AppearanceFont represents a font used for generating the appearance of a
// field in the filling/flattening process.
type AppearanceFont struct {
	// Name represents the name of the font which will be added to the
	// AcroForm resources (DR).
	Name string

	// Font represents the actual font used for the field appearance.
	Font *_fa.PdfFont

	// Size represents the size of the font used for the field appearance.
	// If the font size is 0, the value of the FallbackSize field of the
	// AppearanceFontStyle is used, if set. Otherwise, the font size is
	// calculated based on the available annotation height and on the
	// AutoFontSizeFraction field of the AppearanceStyle.
	Size float64
}

// SignatureFieldOpts represents a set of options used to configure
// an appearance widget dictionary.
type SignatureFieldOpts struct {
	// Rect represents the area the signature annotation is displayed on.
	Rect []float64

	// AutoSize specifies if the content of the appearance should be
	// scaled to fit in the annotation rectangle.
	AutoSize bool

	// Font specifies the font of the text content.
	Font *_fa.PdfFont

	// FontSize specifies the size of the text content.
	FontSize float64

	// LineHeight specifies the height of a line of text in the appearance annotation.
	LineHeight float64

	// TextColor represents the color of the text content displayed.
	TextColor _fa.PdfColor

	// FillColor represents the background color of the appearance annotation area.
	FillColor _fa.PdfColor

	// BorderSize represents border size of the appearance annotation area.
	BorderSize float64

	// BorderColor represents the border color of the appearance annotation area.
	BorderColor _fa.PdfColor

	// WatermarkImage specifies the image used as a watermark that will be rendered
	// behind the signature.
	WatermarkImage _e.Image

	// Image represents the image used for the signature appearance.
	Image _e.Image

	// Encoder specifies the image encoder used for image signature. Defaults to flate encoder.
	Encoder _dd.StreamEncoder

	// ImagePosition specifies the image location relative to the text signature.
	ImagePosition SignatureImagePosition
}

// SetStyle applies appearance `style` to `fa`.
func (_cefe *ImageFieldAppearance) SetStyle(style AppearanceStyle) { _cefe._aaf = &style }

// GenerateAppearanceDict generates an appearance dictionary for widget annotation `wa` for the `field` in `form`.
// Implements interface model.FieldAppearanceGenerator.
func (_aa FieldAppearance) GenerateAppearanceDict(form *_fa.PdfAcroForm, field *_fa.PdfField, wa *_fa.PdfAnnotationWidget) (*_dd.PdfObjectDictionary, error) {
	_a.Log.Trace("\u0047\u0065n\u0065\u0072\u0061\u0074e\u0041\u0070p\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0044i\u0063\u0074\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u0020\u0056:\u0020\u0025\u002b\u0076", field.PartialName(), field.V)
	_, _gc := field.GetContext().(*_fa.PdfFieldText)
	_gcc, _bcf := _dd.GetDict(wa.AP)
	if _bcf && _aa.OnlyIfMissing && (!_gc || !_aa.RegenerateTextFields) {
		_a.Log.Trace("\u0041\u006c\u0072\u0065a\u0064\u0079\u0020\u0070\u006f\u0070\u0075\u006c\u0061\u0074e\u0064 \u002d\u0020\u0069\u0067\u006e\u006f\u0072i\u006e\u0067")
		return _gcc, nil
	}
	if form.DR == nil {
		form.DR = _fa.NewPdfPageResources()
	}
	switch _gcg := field.GetContext().(type) {
	case *_fa.PdfFieldText:
		_ec := _gcg
		if _cc := _beed(_ec.PdfField); _cc == "" {
			_ec.DA = form.DA
		}
		switch {
		case _ec.Flags().Has(_fa.FieldFlagPassword):
			return nil, nil
		case _ec.Flags().Has(_fa.FieldFlagFileSelect):
			return nil, nil
		case _ec.Flags().Has(_fa.FieldFlagComb):
			if _ec.MaxLen != nil {
				_bgf, _ef := _gdad(wa, _ec, form.DR, _aa.Style())
				if _ef != nil {
					return nil, _ef
				}
				return _bgf, nil
			}
		}
		_ddg, _bab := _beab(wa, _ec, form.DR, _aa.Style())
		if _bab != nil {
			return nil, _bab
		}
		return _ddg, nil
	case *_fa.PdfFieldButton:
		_gbee := _gcg
		if _gbee.IsCheckbox() {
			_bfd, _de := _fga(wa, _gbee, form.DR, _aa.Style())
			if _de != nil {
				return nil, _de
			}
			return _bfd, nil
		}
		_a.Log.Debug("\u0054\u004f\u0044\u004f\u003a\u0020\u0055\u004e\u0048\u0041\u004e\u0044\u004c\u0045\u0044 \u0062u\u0074\u0074\u006f\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u002b\u0076", _gbee.GetType())
	case *_fa.PdfFieldChoice:
		_afe := _gcg
		switch {
		case _afe.Flags().Has(_fa.FieldFlagCombo):
			_ea, _fca := _cfc(form, wa, _afe, _aa.Style())
			if _fca != nil {
				return nil, _fca
			}
			return _ea, nil
		default:
			_a.Log.Debug("\u0054\u004f\u0044\u004f\u003a\u0020\u0055N\u0048\u0041\u004eD\u004c\u0045\u0044\u0020c\u0068\u006f\u0069\u0063\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0077\u0069\u0074\u0068\u0020\u0066\u006c\u0061\u0067\u0073\u003a\u0020\u0025\u0073", _afe.Flags().String())
		}
	default:
		_a.Log.Debug("\u0054\u004f\u0044\u004f\u003a\u0020\u0055\u004e\u0048\u0041N\u0044\u004c\u0045\u0044\u0020\u0066\u0069e\u006c\u0064\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _gcg)
	}
	return nil, nil
}

// ComboboxFieldOptions defines optional parameters for a combobox form field.
type ComboboxFieldOptions struct {
	// Choices is the list of string values that can be selected.
	Choices []string
}

// NewImageField generates a new image field with partial name `name` at location `rect`
// on specified `page` and with field specific options `opt`.
func NewImageField(page *_fa.PdfPage, name string, rect []float64, opt ImageFieldOptions) (*_fa.PdfFieldButton, error) {
	if page == nil {
		return nil, _bd.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _bd.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _bd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_afdg := _fa.NewPdfField()
	_bffdg := &_fa.PdfFieldButton{}
	_bffdg.PdfField = _afdg
	_afdg.SetContext(_bffdg)
	_bffdg.SetType(_fa.ButtonTypePush)
	_bffdg.T = _dd.MakeString(name)
	_begg := _fa.NewPdfAnnotationWidget()
	_begg.Rect = _dd.MakeArrayFromFloats(rect)
	_begg.P = page.ToPdfObject()
	_begg.F = _dd.MakeInteger(4)
	_begg.Parent = _bffdg.ToPdfObject()
	_affad := rect[2] - rect[0]
	_gdfg := rect[3] - rect[1]
	_eff := opt._ecf
	_dad := _b.NewContentCreator()
	if _eff.BorderSize > 0 {
		_afdc(_dad, _eff, _affad, _gdfg)
	}
	if _eff.DrawAlignmentReticle {
		_efd := _eff
		_efd.BorderSize = 0.2
		_cece(_dad, _efd, _affad, _gdfg)
	}
	_cbgcb, _ceb := _dgce(_affad, _gdfg, opt.Image, _eff)
	if _ceb != nil {
		return nil, _ceb
	}
	_dbc, _efbcb := _dd.GetDict(_begg.MK)
	if _efbcb {
		_dbc.Set("\u006c", _cbgcb.ToPdfObject())
	}
	_egg := _dd.MakeDict()
	_egg.Set("\u0046\u0052\u004d", _cbgcb.ToPdfObject())
	_ccef := _fa.NewPdfPageResources()
	_ccef.ProcSet = _dd.MakeArray(_dd.MakeName("\u0050\u0044\u0046"))
	_ccef.XObject = _egg
	_ebba := _affad - 2
	_cca := _gdfg - 2
	_dad.Add_q()
	_dad.Add_re(1, 1, _ebba, _cca)
	_dad.Add_W()
	_dad.Add_n()
	_ebba -= 2
	_cca -= 2
	_dad.Add_q()
	_dad.Add_re(2, 2, _ebba, _cca)
	_dad.Add_W()
	_dad.Add_n()
	_baa := _cf.Min(_ebba/float64(opt.Image.Width), _cca/float64(opt.Image.Height))
	_dad.Add_cm(_baa, 0, 0, _baa, (_affad/2)-(float64(opt.Image.Width)*_baa/2)+2, 2)
	_dad.Add_Do("\u0046\u0052\u004d")
	_dad.Add_Q()
	_dad.Add_Q()
	_gefb := _fa.NewXObjectForm()
	_gefb.FormType = _dd.MakeInteger(1)
	_gefb.Resources = _ccef
	_gefb.BBox = _dd.MakeArrayFromFloats([]float64{0, 0, _affad, _gdfg})
	_gefb.Matrix = _dd.MakeArrayFromFloats([]float64{1.0, 0.0, 0.0, 1.0, 0.0, 0.0})
	_gefb.SetContentStream(_dad.Bytes(), _afge())
	_dfff := _dd.MakeDict()
	_dfff.Set("\u004e", _gefb.ToPdfObject())
	_begg.AP = _dfff
	_bffdg.Annotations = append(_bffdg.Annotations, _begg)
	return _bffdg, nil
}

func _abfa(_bcbd []_af.Point) (_gffc []_af.Point, _affae []_af.Point, _cged error) {
	_abgb := len(_bcbd) - 1
	if len(_bcbd) < 1 {
		return nil, nil, _bd.New("\u0041\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0074\u0077\u006f\u0020\u0070\u006f\u0069\u006e\u0074s \u0072e\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0074\u006f\u0020\u0063\u0061l\u0063\u0075\u006c\u0061\u0074\u0065\u0020\u0063\u0075\u0072\u0076\u0065\u0020\u0063\u006f\u006e\u0074r\u006f\u006c\u0020\u0070\u006f\u0069\u006e\u0074\u0073")
	}
	if _abgb == 1 {
		_fbcgd := _af.Point{X: (2*_bcbd[0].X + _bcbd[1].X) / 3, Y: (2*_bcbd[0].Y + _bcbd[1].Y) / 3}
		_gffc = append(_gffc, _fbcgd)
		_affae = append(_affae, _af.Point{X: 2*_fbcgd.X - _bcbd[0].X, Y: 2*_fbcgd.Y - _bcbd[0].Y})
		return _gffc, _affae, nil
	}
	_daaf := make([]float64, _abgb)
	for _adba := 1; _adba < _abgb-1; _adba++ {
		_daaf[_adba] = 4*_bcbd[_adba].X + 2*_bcbd[_adba+1].X
	}
	_daaf[0] = _bcbd[0].X + 2*_bcbd[1].X
	_daaf[_abgb-1] = (8*_bcbd[_abgb-1].X + _bcbd[_abgb].X) / 2.0
	_dbeb := _cdgd(_daaf)
	for _efdc := 1; _efdc < _abgb-1; _efdc++ {
		_daaf[_efdc] = 4*_bcbd[_efdc].Y + 2*_bcbd[_efdc+1].Y
	}
	_daaf[0] = _bcbd[0].Y + 2*_bcbd[1].Y
	_daaf[_abgb-1] = (8*_bcbd[_abgb-1].Y + _bcbd[_abgb].Y) / 2.0
	_dcfa := _cdgd(_daaf)
	_gffc = make([]_af.Point, _abgb)
	_affae = make([]_af.Point, _abgb)
	for _dafe := 0; _dafe < _abgb; _dafe++ {
		_gffc[_dafe] = _af.Point{X: _dbeb[_dafe], Y: _dcfa[_dafe]}
		if _dafe < _abgb-1 {
			_affae[_dafe] = _af.Point{X: 2*_bcbd[_dafe+1].X - _dbeb[_dafe+1], Y: 2*_bcbd[_dafe+1].Y - _dcfa[_dafe+1]}
		} else {
			_affae[_dafe] = _af.Point{X: (_bcbd[_abgb].X + _dbeb[_abgb-1]) / 2, Y: (_bcbd[_abgb].Y + _dcfa[_abgb-1]) / 2}
		}
	}
	return _gffc, _affae, nil
}

func _dbad(_cdcbb LineAnnotationDef, _bagca string) ([]byte, *_fa.PdfRectangle, *_fa.PdfRectangle, error) {
	_fdbc := _af.Line{X1: 0, Y1: 0, X2: _cdcbb.X2 - _cdcbb.X1, Y2: _cdcbb.Y2 - _cdcbb.Y1, LineColor: _cdcbb.LineColor, Opacity: _cdcbb.Opacity, LineWidth: _cdcbb.LineWidth, LineEndingStyle1: _cdcbb.LineEndingStyle1, LineEndingStyle2: _cdcbb.LineEndingStyle2}
	_ecab, _febd, _dfaf := _fdbc.Draw(_bagca)
	if _dfaf != nil {
		return nil, nil, nil, _dfaf
	}
	_fcde := &_fa.PdfRectangle{}
	_fcde.Llx = _cdcbb.X1 + _febd.Llx
	_fcde.Lly = _cdcbb.Y1 + _febd.Lly
	_fcde.Urx = _cdcbb.X1 + _febd.Urx
	_fcde.Ury = _cdcbb.Y1 + _febd.Ury
	return _ecab, _febd, _fcde, nil
}

// CheckboxFieldOptions defines optional parameters for a checkbox field a form.
type CheckboxFieldOptions struct{ Checked bool }

func _cfe(_afd *_fa.PdfField, _bbf, _ddda float64, _dddd string, _dbg AppearanceStyle, _gab *_b.ContentStreamOperations, _bcg *_fa.PdfPageResources, _edde *_dd.PdfObjectDictionary) (*_fa.XObjectForm, error) {
	_cbfe := _fa.NewPdfPageResources()
	_abec, _bcbc := _bbf, _ddda
	_fda := _b.NewContentCreator()
	if _dbg.BorderSize > 0 {
		_afdc(_fda, _dbg, _bbf, _ddda)
	}
	if _dbg.DrawAlignmentReticle {
		_fabg := _dbg
		_fabg.BorderSize = 0.2
		_cece(_fda, _fabg, _bbf, _ddda)
	}
	_fda.Add_BMC("\u0054\u0078")
	_fda.Add_q()
	_fda.Add_BT()
	_bbf, _ddda = _dbg.applyRotation(_edde, _bbf, _ddda, _fda)
	_ebgg, _bff, _eaf := _dbg.processDA(_afd, _gab, _bcg, _cbfe, _fda)
	if _eaf != nil {
		return nil, _eaf
	}
	_efe := _ebgg.Font
	_dgbe := _ebgg.Size
	_fad := _dd.MakeName(_ebgg.Name)
	_bbb := _dgbe == 0
	if _bbb && _bff {
		_dgbe = _ddda * _dbg.AutoFontSizeFraction
	}
	_ffad := _efe.Encoder()
	if _ffad == nil {
		_a.Log.Debug("\u0057\u0041RN\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0069\u0073\u0020\u006e\u0069l\u002e\u0020\u0041\u0073s\u0075\u006d\u0069\u006eg \u0069\u0064e\u006et\u0069\u0074\u0079\u0020\u0065\u006ec\u006f\u0064\u0065r\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
		_ffad = _d.NewIdentityTextEncoder("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048")
	}
	if len(_dddd) == 0 {
		return nil, nil
	}
	_bbbc := _aeb
	if _dbg.MarginLeft != nil {
		_bbbc = *_dbg.MarginLeft
	}
	_fede := 0.0
	if _ffad != nil {
		for _, _efed := range _dddd {
			_fcea, _egc := _efe.GetRuneMetrics(_efed)
			if !_egc {
				_a.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u006d\u0065\u0074r\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u002d\u0020\u0073k\u0069p\u0070\u0069\u006e\u0067", _efed)
				continue
			}
			_fede += _fcea.Wx
		}
		_dddd = string(_ffad.Encode(_dddd))
	}
	if _dgbe == 0 || _bbb && _fede > 0 && _bbbc+_fede*_dgbe/1000.0 > _bbf {
		_dgbe = 0.95 * 1000.0 * (_bbf - _bbbc) / _fede
	}
	_afg := 1.0 * _dgbe
	_cedf := 2.0
	{
		_cfb := _afg
		if _bbb && _cedf+_cfb > _ddda {
			_dgbe = 0.95 * (_ddda - _cedf)
			_afg = 1.0 * _dgbe
			_cfb = _afg
		}
		if _ddda > _cfb {
			_cedf = (_ddda - _cfb) / 2.0
			_cedf += 1.50
		}
	}
	_fda.Add_Tf(*_fad, _dgbe)
	_fda.Add_Td(_bbbc, _cedf)
	_fda.Add_Tj(*_dd.MakeString(_dddd))
	_fda.Add_ET()
	_fda.Add_Q()
	_fda.Add_EMC()
	_gae := _fa.NewXObjectForm()
	_gae.Resources = _cbfe
	_gae.BBox = _dd.MakeArrayFromFloats([]float64{0, 0, _abec, _bcbc})
	_gae.SetContentStream(_fda.Bytes(), _afge())
	return _gae, nil
}

func _cfc(_ecg *_fa.PdfAcroForm, _cedg *_fa.PdfAnnotationWidget, _gga *_fa.PdfFieldChoice, _gcea AppearanceStyle) (*_dd.PdfObjectDictionary, error) {
	_gca, _gbfc := _dd.GetArray(_cedg.Rect)
	if !_gbfc {
		return nil, _bd.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_dfa, _dafa := _fa.NewPdfRectangle(*_gca)
	if _dafa != nil {
		return nil, _dafa
	}
	_ggb, _dffb := _dfa.Width(), _dfa.Height()
	_a.Log.Debug("\u0043\u0068\u006f\u0069\u0063\u0065\u002c\u0020\u0077\u0061\u0020\u0042S\u003a\u0020\u0025\u0076", _cedg.BS)
	_ege, _dafa := _b.NewContentStreamParser(_beed(_gga.PdfField)).Parse()
	if _dafa != nil {
		return nil, _dafa
	}
	_aad, _dge := _dd.GetDict(_cedg.MK)
	if _dge {
		_gcgb, _ := _dd.GetDict(_cedg.BS)
		_afcf := _gcea.applyAppearanceCharacteristics(_aad, _gcgb, nil)
		if _afcf != nil {
			return nil, _afcf
		}
	}
	_fab := _dd.MakeDict()
	for _, _agfa := range _gga.Opt.Elements() {
		if _cedga, _dfag := _dd.GetArray(_agfa); _dfag && _cedga.Len() == 2 {
			_agfa = _cedga.Get(1)
		}
		var _afb string
		if _daad, _cgd := _dd.GetString(_agfa); _cgd {
			_afb = _daad.Decoded()
		} else if _ecd, _bbe := _dd.GetName(_agfa); _bbe {
			_afb = _ecd.String()
		} else {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u004f\u0070\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006de\u002f\u0073\u0074\u0072\u0069\u006e\u0067 \u002d\u0020\u0025\u0054", _agfa)
			return nil, _bd.New("\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u002f\u0073t\u0072\u0069\u006e\u0067")
		}
		if len(_afb) > 0 {
			_cbce, _age := _cfe(_gga.PdfField, _ggb, _dffb, _afb, _gcea, _ege, _ecg.DR, _aad)
			if _age != nil {
				return nil, _age
			}
			_fab.Set(*_dd.MakeName(_afb), _cbce.ToPdfObject())
		}
	}
	_dgd := _dd.MakeDict()
	_dgd.Set("\u004e", _fab)
	return _dgd, nil
}

func _ega(_cdg *_fa.PdfPage, _dbae _af.Rectangle, _geefg string, _acaf string, _gbba _fa.PdfColor, _ebdf *_fa.PdfFont, _ebc *float64, _fdeg _dd.PdfObject) (*_fa.PdfFieldButton, error) {
	_fbbe, _adea := _dbae.X, _dbae.Y
	_dcd := _dbae.Width
	_fdbbe := _dbae.Height
	if _dbae.FillColor == nil {
		_dbae.FillColor = _fa.NewPdfColorDeviceGray(0.7)
	}
	if _gbba == nil {
		_gbba = _fa.NewPdfColorDeviceGray(0)
	}
	if _ebdf == nil {
		_eebcf, _ggf := _fa.NewStandard14Font("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
		if _ggf != nil {
			return nil, _ggf
		}
		_ebdf = _eebcf
	}
	_gabg := _fa.NewPdfField()
	_abg := &_fa.PdfFieldButton{}
	_gabg.SetContext(_abg)
	_abg.PdfField = _gabg
	_abg.T = _dd.MakeString(_geefg)
	_abg.SetType(_fa.ButtonTypePush)
	_abg.V = _dd.MakeName("\u004f\u0066\u0066")
	_abg.Ff = _dd.MakeInteger(4)
	_edf := _dd.MakeDict()
	_edf.Set(*_dd.MakeName("\u0043\u0041"), _dd.MakeString(_acaf))
	_cgg, _fadd := _ebdf.GetFontDescriptor()
	if _fadd != nil {
		return nil, _fadd
	}
	_dcgd := _dd.MakeName("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
	_fdbf := 12.0
	if _cgg != nil && _cgg.FontName != nil {
		_dcgd, _ = _dd.GetName(_cgg.FontName)
	}
	if _ebc != nil {
		_fdbf = *_ebc
	}
	_edea := _b.NewContentCreator()
	_edea.Add_q()
	_edea.SetNonStrokingColor(_dbae.FillColor)
	_edea.Add_re(0, 0, _dcd, _fdbbe)
	_edea.Add_f()
	_edea.Add_Q()
	_edea.Add_q()
	_edea.Add_BT()
	_dcdd := 0.0
	for _, _fef := range _acaf {
		_egde, _add := _ebdf.GetRuneMetrics(_fef)
		if !_add {
			_a.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u006d\u0065\u0074r\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u002d\u0020\u0073k\u0069p\u0070\u0069\u006e\u0067", _fef)
			continue
		}
		_dcdd += _egde.Wx
	}
	_dcdd = _dcdd / 1000.0 * _fdbf
	var _dece float64
	if _cgg != nil {
		_dece, _fadd = _cgg.GetCapHeight()
		if _fadd != nil {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0043\u0061\u0070\u0048\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _fadd)
		}
	}
	if int(_dece) <= 0 {
		_a.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u0061\u0070\u0048e\u0069\u0067\u0068\u0074\u0020\u006e\u006ft \u0061\u0076\u0061\u0069l\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065tt\u0069\u006eg\u0020\u0074\u006f\u0020\u0031\u0030\u0030\u0030")
		_dece = 1000
	}
	_fbeb := _dece / 1000.0 * _fdbf
	_bcdb := (_fdbbe - _fbeb) / 2.0
	_dga := (_dcd - _dcdd) / 2.0
	_edea.Add_Tf(*_dcgd, _fdbf)
	_edea.SetNonStrokingColor(_gbba)
	_edea.Add_Td(_dga, _bcdb)
	_edea.Add_Tj(*_dd.MakeString(_acaf))
	_edea.Add_ET()
	_edea.Add_Q()
	_effc := _fa.NewXObjectForm()
	_effc.SetContentStream(_edea.Bytes(), _dd.NewRawEncoder())
	_effc.BBox = _dd.MakeArrayFromFloats([]float64{0, 0, _dcd, _fdbbe})
	_effc.Resources = _fa.NewPdfPageResources()
	_effc.Resources.SetFontByName(*_dcgd, _ebdf.ToPdfObject())
	_fagd := _dd.MakeDict()
	_fagd.Set("\u004e", _effc.ToPdfObject())
	_egf := _fa.NewPdfAnnotationWidget()
	_egf.Rect = _dd.MakeArrayFromFloats([]float64{_fbbe, _adea, _fbbe + _dcd, _adea + _fdbbe})
	_egf.P = _cdg.ToPdfObject()
	_egf.F = _dd.MakeInteger(4)
	_egf.Parent = _abg.ToPdfObject()
	_egf.A = _fdeg
	_egf.MK = _edf
	_egf.AP = _fagd
	_abg.Annotations = append(_abg.Annotations, _egf)
	return _abg, nil
}

// SignatureImagePosition specifies the image signature location relative to the text signature.
// If text signature is not defined, this position will be ignored.
type SignatureImagePosition int

func _dabc(_dgdb *_fa.PdfFieldButton, _egce *_fa.PdfAnnotationWidget, _ceff AppearanceStyle) (*_dd.PdfObjectDictionary, error) {
	_fdbbc, _egcea := _dd.GetArray(_egce.Rect)
	if !_egcea {
		return nil, _bd.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_fada, _adeac := _fa.NewPdfRectangle(*_fdbbc)
	if _adeac != nil {
		return nil, _adeac
	}
	_dcad, _eef := _fada.Width(), _fada.Height()
	_fbfc := _b.NewContentCreator()
	if _ceff.BorderSize > 0 {
		_afdc(_fbfc, _ceff, _dcad, _eef)
	}
	if _ceff.DrawAlignmentReticle {
		_ebged := _ceff
		_ebged.BorderSize = 0.2
		_cece(_fbfc, _ebged, _dcad, _eef)
	}
	_deec := _dgdb.GetFillImage()
	_babbb, _adeac := _dgce(_dcad, _eef, _deec, _ceff)
	if _adeac != nil {
		return nil, _adeac
	}
	_fdag, _ccdc := _dd.GetDict(_egce.MK)
	if _ccdc {
		_fdag.Set("\u006c", _babbb.ToPdfObject())
	}
	_gcca := _dd.MakeDict()
	_gcca.Set("\u0046\u0052\u004d", _babbb.ToPdfObject())
	_gbc := _fa.NewPdfPageResources()
	_gbc.ProcSet = _dd.MakeArray(_dd.MakeName("\u0050\u0044\u0046"))
	_gbc.XObject = _gcca
	_cfdb := _dcad - 2
	_ebcg := _eef - 2
	_fbfc.Add_q()
	_fbfc.Add_re(1, 1, _cfdb, _ebcg)
	_fbfc.Add_W()
	_fbfc.Add_n()
	_cfdb -= 2
	_ebcg -= 2
	_fbfc.Add_q()
	_fbfc.Add_re(2, 2, _cfdb, _ebcg)
	_fbfc.Add_W()
	_fbfc.Add_n()
	_bdcea := _cf.Min(_cfdb/float64(_deec.Width), _ebcg/float64(_deec.Height))
	_fbfc.Add_cm(_bdcea, 0, 0, _bdcea, (_dcad/2)-(float64(_deec.Width)*_bdcea/2)+2, 2)
	_fbfc.Add_Do("\u0046\u0052\u004d")
	_fbfc.Add_Q()
	_fbfc.Add_Q()
	_befa := _fa.NewXObjectForm()
	_befa.FormType = _dd.MakeInteger(1)
	_befa.Resources = _gbc
	_befa.BBox = _dd.MakeArrayFromFloats([]float64{0, 0, _dcad, _eef})
	_befa.Matrix = _dd.MakeArrayFromFloats([]float64{1.0, 0.0, 0.0, 1.0, 0.0, 0.0})
	_befa.SetContentStream(_fbfc.Bytes(), _afge())
	_aabff := _dd.MakeDict()
	_aabff.Set("\u004e", _befa.ToPdfObject())
	return _aabff, nil
}

func (_gcgc *AppearanceStyle) processDA(_bcfg *_fa.PdfField, _faea *_b.ContentStreamOperations, _afcg, _gdab *_fa.PdfPageResources, _eceg *_b.ContentCreator) (*AppearanceFont, bool, error) {
	var _ecb *AppearanceFont
	var _efcd bool
	if _gcgc.Fonts != nil {
		if _gcgc.Fonts.Fallback != nil {
			_ecb = _gcgc.Fonts.Fallback
		}
		if _ged := _gcgc.Fonts.FieldFallbacks; _ged != nil {
			if _cefc, _afde := _ged[_bcfg.PartialName()]; _afde {
				_ecb = _cefc
			} else if _abc, _cce := _bcfg.FullName(); _cce == nil {
				if _ecbb, _aag := _ged[_abc]; _aag {
					_ecb = _ecbb
				}
			}
		}
		if _ecb != nil {
			_ecb.fillName()
		}
		_efcd = _gcgc.Fonts.ForceReplace
	}
	var _dfe string
	var _aeg float64
	var _ace bool
	if _faea != nil {
		for _, _dcgae := range *_faea {
			if _dcgae.Operand == "\u0054\u0066" && len(_dcgae.Params) == 2 {
				if _dcgb, _dagd := _dd.GetNameVal(_dcgae.Params[0]); _dagd {
					_dfe = _dcgb
				}
				if _cefb, _ddaf := _dd.GetNumberAsFloat(_dcgae.Params[1]); _ddaf == nil {
					_aeg = _cefb
				}
				_ace = true
				continue
			}
			_eceg.AddOperand(*_dcgae)
		}
	}
	var _bed *AppearanceFont
	var _geb _dd.PdfObject
	if _efcd && _ecb != nil {
		_bed = _ecb
	} else {
		if _afcg != nil && _dfe != "" {
			if _aec, _dgc := _afcg.GetFontByName(*_dd.MakeName(_dfe)); _dgc {
				if _ceef, _caf := _fa.NewPdfFontFromPdfObject(_aec); _caf == nil {
					_geb = _aec
					_bed = &AppearanceFont{Name: _dfe, Font: _ceef, Size: _aeg}
				} else {
					_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006fa\u0064\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0025\u0076", _caf)
				}
			}
		}
		if _bed == nil && _ecb != nil {
			_bed = _ecb
		}
		if _bed == nil {
			_ffg, _eede := _fa.NewStandard14Font("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
			if _eede != nil {
				return nil, false, _eede
			}
			_bed = &AppearanceFont{Name: "\u0048\u0065\u006c\u0076", Font: _ffg, Size: _aeg}
		}
	}
	if _bed.Size <= 0 && _gcgc.Fonts != nil && _gcgc.Fonts.FallbackSize > 0 {
		_bed.Size = _gcgc.Fonts.FallbackSize
	}
	_eceba := *_dd.MakeName(_bed.Name)
	if _geb == nil {
		_geb = _bed.Font.ToPdfObject()
	}
	if _afcg != nil && !_afcg.HasFontByName(_eceba) {
		_afcg.SetFontByName(_eceba, _geb)
	}
	if _gdab != nil && !_gdab.HasFontByName(_eceba) {
		_gdab.SetFontByName(_eceba, _geb)
	}
	return _bed, _ace, nil
}

// FormSubmitActionOptions holds options for creating a form submit button.
type FormSubmitActionOptions struct {
	// Rectangle holds the button position, size, and color.
	Rectangle _af.Rectangle

	// Url specifies the URL where the fieds will be submitted.
	Url string

	// Label specifies the text that would be displayed on the button.
	Label string

	// LabelColor specifies the button label color.
	LabelColor _fa.PdfColor

	// Font specifies a font used for rendering the button label.
	// When omitted it will fallback to use a Helvetica font.
	Font *_fa.PdfFont

	// FontSize specifies the font size used in rendering the button label.
	// The default font size is 12pt.
	FontSize *float64

	// Fields specifies list of fields that could be submitted.
	// This list may contain indirect object to fields or field names.
	Fields *_dd.PdfObjectArray

	// IsExclusionList specifies that the fields contain in `Fields` array would not be submitted.
	IsExclusionList bool

	// IncludeEmptyFields specifies if all fields would be submitted even though it's value is empty.
	IncludeEmptyFields bool

	// SubmitAsPDF specifies that the document shall be submitted as PDF.
	// If set then all the other flags shall be ignored.
	SubmitAsPDF bool
}

func _afdc(_ecdc *_b.ContentCreator, _faa AppearanceStyle, _eceb, _fdgf float64) {
	_ecdc.Add_q().Add_re(0, 0, _eceb, _fdgf).Add_w(_faa.BorderSize).SetStrokingColor(_faa.BorderColor).SetNonStrokingColor(_faa.FillColor).Add_B().Add_Q()
}

// WrapContentStream ensures that the entire content stream for a `page` is wrapped within q ... Q operands.
// Ensures that following operands that are added are not affected by additional operands that are added.
// Implements interface model.ContentStreamWrapper.
func (_bbad FieldAppearance) WrapContentStream(page *_fa.PdfPage) error {
	_gacf, _fecg := page.GetAllContentStreams()
	if _fecg != nil {
		return _fecg
	}
	_eac := _b.NewContentStreamParser(_gacf)
	_ddde, _fecg := _eac.Parse()
	if _fecg != nil {
		return _fecg
	}
	_ddde.WrapIfNeeded()
	_cgaf := []string{_ddde.String()}
	return page.SetContentStreams(_cgaf, _afge())
}

func _dgec(_cbcdc [][]_af.CubicBezierCurve, _ggcf *_fa.PdfColorDeviceRGB, _fbfce float64) ([]byte, *_fa.PdfRectangle, error) {
	_daef := _b.NewContentCreator()
	_daef.Add_q().SetStrokingColor(_ggcf).Add_w(_fbfce)
	_bcec := _af.NewCubicBezierPath()
	for _, _gdfc := range _cbcdc {
		_bcec.Curves = append(_bcec.Curves, _gdfc...)
		for _bgef, _bagce := range _gdfc {
			if _bgef == 0 {
				_daef.Add_m(_bagce.P0.X, _bagce.P0.Y)
			} else {
				_daef.Add_l(_bagce.P0.X, _bagce.P0.Y)
			}
			_daef.Add_c(_bagce.P1.X, _bagce.P1.Y, _bagce.P2.X, _bagce.P2.Y, _bagce.P3.X, _bagce.P3.Y)
		}
	}
	_daef.Add_S().Add_Q()
	return _daef.Bytes(), _bcec.GetBoundingBox().ToPdfRectangle(), nil
}

// CreateInkAnnotation creates an ink annotation object that can be added to the annotation list of a PDF page.
func CreateInkAnnotation(inkDef InkAnnotationDef) (*_fa.PdfAnnotation, error) {
	_effb := _fa.NewPdfAnnotationInk()
	_efda := _dd.MakeArray()
	for _, _ffbe := range inkDef.Paths {
		if _ffbe.Length() == 0 {
			continue
		}
		_ebgc := []float64{}
		for _, _fdeb := range _ffbe.Points {
			_ebgc = append(_ebgc, _fdeb.X, _fdeb.Y)
		}
		_efda.Append(_dd.MakeArrayFromFloats(_ebgc))
	}
	_effb.InkList = _efda
	if inkDef.Color == nil {
		inkDef.Color = _fa.NewPdfColorDeviceRGB(0.0, 0.0, 0.0)
	}
	_effb.C = _dd.MakeArrayFromFloats([]float64{inkDef.Color.R(), inkDef.Color.G(), inkDef.Color.B()})
	_efbgc, _fbde, _gefc := _bga(&inkDef)
	if _gefc != nil {
		return nil, _gefc
	}
	_effb.AP = _efbgc
	_effb.Rect = _dd.MakeArrayFromFloats([]float64{_fbde.Llx, _fbde.Lly, _fbde.Urx, _fbde.Ury})
	return _effb.PdfAnnotation, nil
}

const (
	_bag quadding = 0
	_bb  quadding = 1
	_ebg quadding = 2
	_aeb float64  = 2.0
)

// TextFieldOptions defines optional parameter for a text field in a form.
type TextFieldOptions struct {
	MaxLen int
	Value  string
}

func _cece(_cac *_b.ContentCreator, _bbbd AppearanceStyle, _dddg, _ddc float64) {
	_cac.Add_q().Add_re(0, 0, _dddg, _ddc).Add_re(0, _ddc/2, _dddg, _ddc/2).Add_re(0, 0, _dddg, _ddc).Add_re(_dddg/2, 0, _dddg/2, _ddc).Add_w(_bbbd.BorderSize).SetStrokingColor(_bbbd.BorderColor).SetNonStrokingColor(_bbbd.FillColor).Add_B().Add_Q()
}

func _aebe(_abb *InkAnnotationDef) ([]byte, *_fa.PdfRectangle, error) {
	_gdgg := [][]_af.CubicBezierCurve{}
	for _, _bbcc := range _abb.Paths {
		if _bbcc.Length() == 0 {
			continue
		}
		_gec := _bbcc.Points
		_edeac, _cbdb, _cgbb := _abfa(_gec)
		if _cgbb != nil {
			return nil, nil, _cgbb
		}
		if len(_edeac) != len(_cbdb) {
			return nil, nil, _bd.New("\u0049\u006e\u0065\u0071\u0075\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0063\u0061l\u0063\u0075\u006c\u0061\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0061\u006e\u0064\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0063\u006f\u006e\u0074\u0072o\u006c\u0020\u0070\u006f\u0069n\u0074")
		}
		_eead := []_af.CubicBezierCurve{}
		for _feea := 0; _feea < len(_edeac); _feea++ {
			_eead = append(_eead, _af.CubicBezierCurve{P0: _gec[_feea], P1: _edeac[_feea], P2: _cbdb[_feea], P3: _gec[_feea+1]})
		}
		if len(_eead) > 0 {
			_gdgg = append(_gdgg, _eead)
		}
	}
	_dgea, _gebbe, _fcdaf := _dgec(_gdgg, _abb.Color, _abb.LineWidth)
	if _fcdaf != nil {
		return nil, nil, _fcdaf
	}
	return _dgea, _gebbe, nil
}

func _adag(_gbcfa LineAnnotationDef) (*_dd.PdfObjectDictionary, *_fa.PdfRectangle, error) {
	_ceeg := _fa.NewXObjectForm()
	_ceeg.Resources = _fa.NewPdfPageResources()
	_fbge := ""
	if _gbcfa.Opacity < 1.0 {
		_fefd := _dd.MakeDict()
		_fefd.Set("\u0063\u0061", _dd.MakeFloat(_gbcfa.Opacity))
		_eafb := _ceeg.Resources.AddExtGState("\u0067\u0073\u0031", _fefd)
		if _eafb != nil {
			_a.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020\u0065\u0078\u0074g\u0073\u0074\u0061t\u0065 \u0067\u0073\u0031")
			return nil, nil, _eafb
		}
		_fbge = "\u0067\u0073\u0031"
	}
	_cba, _bcgb, _acd, _bfdc := _dbad(_gbcfa, _fbge)
	if _bfdc != nil {
		return nil, nil, _bfdc
	}
	_bfdc = _ceeg.SetContentStream(_cba, nil)
	if _bfdc != nil {
		return nil, nil, _bfdc
	}
	_ceeg.BBox = _bcgb.ToPdfObject()
	_baaa := _dd.MakeDict()
	_baaa.Set("\u004e", _ceeg.ToPdfObject())
	return _baaa, _acd, nil
}

// NewComboboxField generates a new combobox form field with partial name `name` at location `rect`
// on specified `page` and with field specific options `opt`.
func NewComboboxField(page *_fa.PdfPage, name string, rect []float64, opt ComboboxFieldOptions) (*_fa.PdfFieldChoice, error) {
	if page == nil {
		return nil, _bd.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _bd.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _bd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_gbfcc := _fa.NewPdfField()
	_cfgc := &_fa.PdfFieldChoice{}
	_gbfcc.SetContext(_cfgc)
	_cfgc.PdfField = _gbfcc
	_cfgc.T = _dd.MakeString(name)
	_cfgc.Opt = _dd.MakeArray()
	for _, _babb := range opt.Choices {
		_cfgc.Opt.Append(_dd.MakeString(_babb))
	}
	_cfgc.SetFlag(_fa.FieldFlagCombo)
	_daec := _fa.NewPdfAnnotationWidget()
	_daec.Rect = _dd.MakeArrayFromFloats(rect)
	_daec.P = page.ToPdfObject()
	_daec.F = _dd.MakeInteger(4)
	_daec.Parent = _cfgc.ToPdfObject()
	_cfgc.Annotations = append(_cfgc.Annotations, _daec)
	return _cfgc, nil
}

func (_bg *AppearanceFont) fillName() {
	if _bg.Font == nil || _bg.Name != "" {
		return
	}
	_bdg := _bg.Font.FontDescriptor()
	if _bdg == nil || _bdg.FontName == nil {
		return
	}
	_bg.Name = _bdg.FontName.String()
}

// GenerateAppearanceDict generates an appearance dictionary for widget annotation `wa` for the `field` in `form`.
// Implements interface model.FieldAppearanceGenerator.
func (_cgf ImageFieldAppearance) GenerateAppearanceDict(form *_fa.PdfAcroForm, field *_fa.PdfField, wa *_fa.PdfAnnotationWidget) (*_dd.PdfObjectDictionary, error) {
	_, _ccgb := field.GetContext().(*_fa.PdfFieldButton)
	if !_ccgb {
		_a.Log.Trace("C\u006f\u0075\u006c\u0064\u0020\u006fn\u006c\u0079\u0020\u0068\u0061\u006ed\u006c\u0065\u0020\u0062\u0075\u0074\u0074o\u006e\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067")
		return nil, nil
	}
	_gaee, _cge := _dd.GetDict(wa.AP)
	if _cge && _cgf.OnlyIfMissing {
		_a.Log.Trace("\u0041\u006c\u0072\u0065a\u0064\u0079\u0020\u0070\u006f\u0070\u0075\u006c\u0061\u0074e\u0064 \u002d\u0020\u0069\u0067\u006e\u006f\u0072i\u006e\u0067")
		return _gaee, nil
	}
	if form.DR == nil {
		form.DR = _fa.NewPdfPageResources()
	}
	switch _ddfe := field.GetContext().(type) {
	case *_fa.PdfFieldButton:
		if _ddfe.IsPush() {
			_badf, _aeca := _dabc(_ddfe, wa, _cgf.Style())
			if _aeca != nil {
				return nil, _aeca
			}
			return _badf, nil
		}
	}
	return nil, nil
}

// CreateLineAnnotation creates a line annotation object that can be added to page PDF annotations.
func CreateLineAnnotation(lineDef LineAnnotationDef) (*_fa.PdfAnnotation, error) {
	_egdf := _fa.NewPdfAnnotationLine()
	_egdf.L = _dd.MakeArrayFromFloats([]float64{lineDef.X1, lineDef.Y1, lineDef.X2, lineDef.Y2})
	_bcae := _dd.MakeName("\u004e\u006f\u006e\u0065")
	if lineDef.LineEndingStyle1 == _af.LineEndingStyleArrow {
		_bcae = _dd.MakeName("C\u006c\u006f\u0073\u0065\u0064\u0041\u0072\u0072\u006f\u0077")
	}
	_badg := _dd.MakeName("\u004e\u006f\u006e\u0065")
	if lineDef.LineEndingStyle2 == _af.LineEndingStyleArrow {
		_badg = _dd.MakeName("C\u006c\u006f\u0073\u0065\u0064\u0041\u0072\u0072\u006f\u0077")
	}
	_egdf.LE = _dd.MakeArray(_bcae, _badg)
	if lineDef.Opacity < 1.0 {
		_egdf.CA = _dd.MakeFloat(lineDef.Opacity)
	}
	_ddcg, _addf, _fge := lineDef.LineColor.R(), lineDef.LineColor.G(), lineDef.LineColor.B()
	_egdf.IC = _dd.MakeArrayFromFloats([]float64{_ddcg, _addf, _fge})
	_egdf.C = _dd.MakeArrayFromFloats([]float64{_ddcg, _addf, _fge})
	_fea := _fa.NewBorderStyle()
	_fea.SetBorderWidth(lineDef.LineWidth)
	_egdf.BS = _fea.ToPdfObject()
	_febf, _efede, _aea := _adag(lineDef)
	if _aea != nil {
		return nil, _aea
	}
	_egdf.AP = _febf
	_egdf.Rect = _dd.MakeArrayFromFloats([]float64{_efede.Llx, _efede.Lly, _efede.Urx, _efede.Ury})
	return _egdf.PdfAnnotation, nil
}

// SignatureLine represents a line of information in the signature field appearance.
type SignatureLine struct {
	Desc string
	Text string
}

// LineAnnotationDef defines a line between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none
// (regular line), or arrows at either end.  The line also has a specified width, color and opacity.
type LineAnnotationDef struct {
	X1               float64
	Y1               float64
	X2               float64
	Y2               float64
	LineColor        *_fa.PdfColorDeviceRGB
	Opacity          float64
	LineWidth        float64
	LineEndingStyle1 _af.LineEndingStyle
	LineEndingStyle2 _af.LineEndingStyle
}

func _ab(_bce CircleAnnotationDef) (*_dd.PdfObjectDictionary, *_fa.PdfRectangle, error) {
	_cb := _fa.NewXObjectForm()
	_cb.Resources = _fa.NewPdfPageResources()
	_gd := ""
	if _bce.Opacity < 1.0 {
		_gef := _dd.MakeDict()
		_gef.Set("\u0063\u0061", _dd.MakeFloat(_bce.Opacity))
		_gef.Set("\u0043\u0041", _dd.MakeFloat(_bce.Opacity))
		_dc := _cb.Resources.AddExtGState("\u0067\u0073\u0031", _gef)
		if _dc != nil {
			_a.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020\u0065\u0078\u0074g\u0073\u0074\u0061t\u0065 \u0067\u0073\u0031")
			return nil, nil, _dc
		}
		_gd = "\u0067\u0073\u0031"
	}
	_ff, _ag, _bf, _fd := _fe(_bce, _gd)
	if _fd != nil {
		return nil, nil, _fd
	}
	_fd = _cb.SetContentStream(_ff, nil)
	if _fd != nil {
		return nil, nil, _fd
	}
	_cb.BBox = _ag.ToPdfObject()
	_cfg := _dd.MakeDict()
	_cfg.Set("\u004e", _cb.ToPdfObject())
	return _cfg, _bf, nil
}

func (_gfg *AppearanceStyle) applyRotation(_geef *_dd.PdfObjectDictionary, _cbgg, _afcd float64, _acb *_b.ContentCreator) (float64, float64) {
	if !_gfg.AllowMK {
		return _cbgg, _afcd
	}
	if _geef == nil {
		return _cbgg, _afcd
	}
	_bbef, _ := _dd.GetNumberAsFloat(_geef.Get("\u0052"))
	if _bbef == 0 {
		return _cbgg, _afcd
	}
	_ada := -_bbef
	_adb := _af.Path{Points: []_af.Point{_af.NewPoint(0, 0).Rotate(_ada), _af.NewPoint(_cbgg, 0).Rotate(_ada), _af.NewPoint(0, _afcd).Rotate(_ada), _af.NewPoint(_cbgg, _afcd).Rotate(_ada)}}.GetBoundingBox()
	_acb.RotateDeg(_bbef)
	_acb.Translate(_adb.X, _adb.Y)
	return _adb.Width, _adb.Height
}

type quadding int

func _cbac(_beee RectangleAnnotationDef) (*_dd.PdfObjectDictionary, *_fa.PdfRectangle, error) {
	_ccaa := _fa.NewXObjectForm()
	_ccaa.Resources = _fa.NewPdfPageResources()
	_gddg := ""
	if _beee.Opacity < 1.0 {
		_adef := _dd.MakeDict()
		_adef.Set("\u0063\u0061", _dd.MakeFloat(_beee.Opacity))
		_adef.Set("\u0043\u0041", _dd.MakeFloat(_beee.Opacity))
		_bdcg := _ccaa.Resources.AddExtGState("\u0067\u0073\u0031", _adef)
		if _bdcg != nil {
			_a.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020\u0065\u0078\u0074g\u0073\u0074\u0061t\u0065 \u0067\u0073\u0031")
			return nil, nil, _bdcg
		}
		_gddg = "\u0067\u0073\u0031"
	}
	_fdcd, _fggc, _agaa, _bbd := _dgcg(_beee, _gddg)
	if _bbd != nil {
		return nil, nil, _bbd
	}
	_bbd = _ccaa.SetContentStream(_fdcd, nil)
	if _bbd != nil {
		return nil, nil, _bbd
	}
	_ccaa.BBox = _fggc.ToPdfObject()
	_dged := _dd.MakeDict()
	_dged.Set("\u004e", _ccaa.ToPdfObject())
	return _dged, _agaa, nil
}

// WrapContentStream ensures that the entire content stream for a `page` is wrapped within q ... Q operands.
// Ensures that following operands that are added are not affected by additional operands that are added.
// Implements interface model.ContentStreamWrapper.
func (_ccb ImageFieldAppearance) WrapContentStream(page *_fa.PdfPage) error {
	_efedd, _cbcd := page.GetAllContentStreams()
	if _cbcd != nil {
		return _cbcd
	}
	_feg := _b.NewContentStreamParser(_efedd)
	_babc, _cbcd := _feg.Parse()
	if _cbcd != nil {
		return _cbcd
	}
	_babc.WrapIfNeeded()
	_dafc := []string{_babc.String()}
	return page.SetContentStreams(_dafc, _afge())
}

// CreateRectangleAnnotation creates a rectangle annotation object that can be added to page PDF annotations.
func CreateRectangleAnnotation(rectDef RectangleAnnotationDef) (*_fa.PdfAnnotation, error) {
	_ccde := _fa.NewPdfAnnotationSquare()
	if rectDef.BorderEnabled {
		_cdaa, _dbag, _aaeb := rectDef.BorderColor.R(), rectDef.BorderColor.G(), rectDef.BorderColor.B()
		_ccde.C = _dd.MakeArrayFromFloats([]float64{_cdaa, _dbag, _aaeb})
		_acfb := _fa.NewBorderStyle()
		_acfb.SetBorderWidth(rectDef.BorderWidth)
		_ccde.BS = _acfb.ToPdfObject()
	}
	if rectDef.FillEnabled {
		_adc, _agdf, _edfd := rectDef.FillColor.R(), rectDef.FillColor.G(), rectDef.FillColor.B()
		_ccde.IC = _dd.MakeArrayFromFloats([]float64{_adc, _agdf, _edfd})
	} else {
		_ccde.IC = _dd.MakeArrayFromIntegers([]int{})
	}
	if rectDef.Opacity < 1.0 {
		_ccde.CA = _dd.MakeFloat(rectDef.Opacity)
	}
	_acgc, _agfg, _eadd := _cbac(rectDef)
	if _eadd != nil {
		return nil, _eadd
	}
	_ccde.AP = _acgc
	_ccde.Rect = _dd.MakeArrayFromFloats([]float64{_agfg.Llx, _agfg.Lly, _agfg.Urx, _agfg.Ury})
	return _ccde.PdfAnnotation, nil
}

func _gdad(_bgca *_fa.PdfAnnotationWidget, _ddd *_fa.PdfFieldText, _ecee *_fa.PdfPageResources, _aef AppearanceStyle) (*_dd.PdfObjectDictionary, error) {
	_fdbb := _fa.NewPdfPageResources()
	_gfc, _eea := _dd.GetArray(_bgca.Rect)
	if !_eea {
		return nil, _bd.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_ddeg, _gbefe := _fa.NewPdfRectangle(*_gfc)
	if _gbefe != nil {
		return nil, _gbefe
	}
	_bead, _bbaf := _ddeg.Width(), _ddeg.Height()
	_bfg, _dee := _bead, _bbaf
	_eaea, _bcbe := _dd.GetDict(_bgca.MK)
	if _bcbe {
		_bgga, _ := _dd.GetDict(_bgca.BS)
		_gbbc := _aef.applyAppearanceCharacteristics(_eaea, _bgga, nil)
		if _gbbc != nil {
			return nil, _gbbc
		}
	}
	_fce, _bcbe := _dd.GetIntVal(_ddd.MaxLen)
	if !_bcbe {
		return nil, _bd.New("\u006d\u0061\u0078\u006c\u0065\u006e\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
	}
	if _fce <= 0 {
		return nil, _bd.New("\u006d\u0061\u0078\u004c\u0065\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_abaa := _bead / float64(_fce)
	_dabe, _gbefe := _b.NewContentStreamParser(_beed(_ddd.PdfField)).Parse()
	if _gbefe != nil {
		return nil, _gbefe
	}
	_fcda := _b.NewContentCreator()
	if _aef.BorderSize > 0 {
		_afdc(_fcda, _aef, _bead, _bbaf)
	}
	if _aef.DrawAlignmentReticle {
		_dcg := _aef
		_dcg.BorderSize = 0.2
		_cece(_fcda, _dcg, _bead, _bbaf)
	}
	_fcda.Add_BMC("\u0054\u0078")
	_fcda.Add_q()
	_, _bbaf = _aef.applyRotation(_eaea, _bead, _bbaf, _fcda)
	_fcda.Add_BT()
	_eaa, _bbg, _gbefe := _aef.processDA(_ddd.PdfField, _dabe, _ecee, _fdbb, _fcda)
	if _gbefe != nil {
		return nil, _gbefe
	}
	_gg := _eaa.Font
	_ebf := _dd.MakeName(_eaa.Name)
	_bfbb := _eaa.Size
	_beg := _bfbb == 0
	if _beg && _bbg {
		_bfbb = _bbaf * _aef.AutoFontSizeFraction
	}
	_fbc := _gg.Encoder()
	if _fbc == nil {
		_a.Log.Debug("\u0057\u0041RN\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0069\u0073\u0020\u006e\u0069l\u002e\u0020\u0041\u0073s\u0075\u006d\u0069\u006eg \u0069\u0064e\u006et\u0069\u0074\u0079\u0020\u0065\u006ec\u006f\u0064\u0065r\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
		_fbc = _d.NewIdentityTextEncoder("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048")
	}
	var _bgff string
	if _bge, _dcf := _dd.GetString(_ddd.V); _dcf {
		_bgff = _bge.Decoded()
	}
	_fcda.Add_Tf(*_ebf, _bfbb)
	var _efbg float64
	for _, _feb := range _bgff {
		_cbgc, _cec := _gg.GetRuneMetrics(_feb)
		if !_cec {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0052\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0066\u006fn\u0074\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067 \u006f\u0076\u0065\u0072", _feb)
			continue
		}
		_fbe := _cbgc.Wy
		if int(_fbe) <= 0 {
			_fbe = _cbgc.Wx
		}
		if _fbe > _efbg {
			_efbg = _fbe
		}
	}
	if int(_efbg) == 0 {
		_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074o\u0020\u0064\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u006d\u0061x\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u0073\u0069\u007a\u0065\u0020- \u0075\u0073\u0069\u006e\u0067\u0020\u0031\u0030\u0030\u0030")
		_efbg = 1000
	}
	_fgg, _gbefe := _gg.GetFontDescriptor()
	if _gbefe != nil {
		_a.Log.Debug("\u0045\u0072ro\u0072\u003a\u0020U\u006e\u0061\u0062\u006ce t\u006f g\u0065\u0074\u0020\u0066\u006f\u006e\u0074 d\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
	}
	var _fgc float64
	if _fgg != nil {
		_fgc, _gbefe = _fgg.GetCapHeight()
		if _gbefe != nil {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0043\u0061\u0070\u0048\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _gbefe)
		}
	}
	if int(_fgc) <= 0 {
		_a.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u0061\u0070\u0048e\u0069\u0067\u0068\u0074\u0020\u006e\u006ft \u0061\u0076\u0061\u0069l\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065tt\u0069\u006eg\u0020\u0074\u006f\u0020\u0031\u0030\u0030\u0030")
		_fgc = 1000.0
	}
	_cgc := _fgc / 1000.0 * _bfbb
	_fbg := 0.0
	_abd := 1.0 * _bfbb * (_efbg / 1000.0)
	{
		_bggf := _abd
		if _beg && _fbg+_bggf > _bbaf {
			_bfbb = 0.95 * (_bbaf - _fbg)
			_cgc = _fgc / 1000.0 * _bfbb
		}
		if _bbaf > _cgc {
			_fbg = (_bbaf - _cgc) / 2.0
		}
	}
	_fcda.Add_Td(0, _fbg)
	if _gfa, _aac := _dd.GetIntVal(_ddd.Q); _aac {
		switch _gfa {
		case 2:
			if len(_bgff) < _fce {
				_cbc := float64(_fce-len(_bgff)) * _abaa
				_fcda.Add_Td(_cbc, 0)
			}
		}
	}
	for _cbf, _dac := range _bgff {
		_dgb := _aeb
		if _aef.MarginLeft != nil {
			_dgb = *_aef.MarginLeft
		}
		_bgfa := string(_dac)
		if _fbc != nil {
			_dec, _affe := _gg.GetRuneMetrics(_dac)
			if !_affe {
				_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0052\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0066\u006fn\u0074\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067 \u006f\u0076\u0065\u0072", _dac)
				continue
			}
			_bgfa = string(_fbc.Encode(_bgfa))
			_efbc := _bfbb * _dec.Wx / 1000.0
			_dcge := (_abaa - _efbc) / 2
			_dgb = _dcge
		}
		_fcda.Add_Td(_dgb, 0)
		_fcda.Add_Tj(*_dd.MakeString(_bgfa))
		if _cbf != len(_bgff)-1 {
			_fcda.Add_Td(_abaa-_dgb, 0)
		}
	}
	_fcda.Add_ET()
	_fcda.Add_Q()
	_fcda.Add_EMC()
	_ead := _fa.NewXObjectForm()
	_ead.Resources = _fdbb
	_ead.BBox = _dd.MakeArrayFromFloats([]float64{0, 0, _bfg, _dee})
	_ead.SetContentStream(_fcda.Bytes(), _afge())
	_bcfb := _dd.MakeDict()
	_bcfb.Set("\u004e", _ead.ToPdfObject())
	return _bcfb, nil
}

// ImageFieldOptions defines optional parameters for a push button with image attach capability form field.
type ImageFieldOptions struct {
	Image *_fa.Image
	_ecf  AppearanceStyle
}

// NewCheckboxField generates a new checkbox field with partial name `name` at location `rect`
// on specified `page` and with field specific options `opt`.
func NewCheckboxField(page *_fa.PdfPage, name string, rect []float64, opt CheckboxFieldOptions) (*_fa.PdfFieldButton, error) {
	if page == nil {
		return nil, _bd.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _bd.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _bd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_cfdd, _bfcdb := _fa.NewStandard14Font(_fa.ZapfDingbatsName)
	if _bfcdb != nil {
		return nil, _bfcdb
	}
	_aca := _fa.NewPdfField()
	_bfe := &_fa.PdfFieldButton{}
	_aca.SetContext(_bfe)
	_bfe.PdfField = _aca
	_bfe.T = _dd.MakeString(name)
	_bfe.SetType(_fa.ButtonTypeCheckbox)
	_fbbf := "\u004f\u0066\u0066"
	if opt.Checked {
		_fbbf = "\u0059\u0065\u0073"
	}
	_bfe.V = _dd.MakeName(_fbbf)
	_ede := _fa.NewPdfAnnotationWidget()
	_ede.Rect = _dd.MakeArrayFromFloats(rect)
	_ede.P = page.ToPdfObject()
	_ede.F = _dd.MakeInteger(4)
	_ede.Parent = _bfe.ToPdfObject()
	_fadc := rect[2] - rect[0]
	_dbaf := rect[3] - rect[1]
	var _bfgd _c.Buffer
	_bfgd.WriteString("\u0071\u000a")
	_bfgd.WriteString("\u0030 \u0030\u0020\u0031\u0020\u0072\u0067\n")
	_bfgd.WriteString("\u0042\u0054\u000a")
	_bfgd.WriteString("\u002f\u005a\u0061D\u0062\u0020\u0031\u0032\u0020\u0054\u0066\u000a")
	_bfgd.WriteString("\u0045\u0054\u000a")
	_bfgd.WriteString("\u0051\u000a")
	_bdce := _b.NewContentCreator()
	_bdce.Add_q()
	_bdce.Add_rg(0, 0, 1)
	_bdce.Add_BT()
	_bdce.Add_Tf(*_dd.MakeName("\u005a\u0061\u0044\u0062"), 12)
	_bdce.Add_Td(0, 0)
	_bdce.Add_ET()
	_bdce.Add_Q()
	_bbgbb := _fa.NewXObjectForm()
	_bbgbb.SetContentStream(_bdce.Bytes(), _dd.NewRawEncoder())
	_bbgbb.BBox = _dd.MakeArrayFromFloats([]float64{0, 0, _fadc, _dbaf})
	_bbgbb.Resources = _fa.NewPdfPageResources()
	_bbgbb.Resources.SetFontByName("\u005a\u0061\u0044\u0062", _cfdd.ToPdfObject())
	_bdce = _b.NewContentCreator()
	_bdce.Add_q()
	_bdce.Add_re(0, 0, _fadc, _dbaf)
	_bdce.Add_W().Add_n()
	_bdce.Add_rg(0, 0, 1)
	_bdce.Translate(0, 3.0)
	_bdce.Add_BT()
	_bdce.Add_Tf(*_dd.MakeName("\u005a\u0061\u0044\u0062"), 12)
	_bdce.Add_Td(0, 0)
	_bdce.Add_Tj(*_dd.MakeString("\u0034"))
	_bdce.Add_ET()
	_bdce.Add_Q()
	_ggac := _fa.NewXObjectForm()
	_ggac.SetContentStream(_bdce.Bytes(), _dd.NewRawEncoder())
	_ggac.BBox = _dd.MakeArrayFromFloats([]float64{0, 0, _fadc, _dbaf})
	_ggac.Resources = _fa.NewPdfPageResources()
	_ggac.Resources.SetFontByName("\u005a\u0061\u0044\u0062", _cfdd.ToPdfObject())
	_faafa := _dd.MakeDict()
	_faafa.Set("\u004f\u0066\u0066", _bbgbb.ToPdfObject())
	_faafa.Set("\u0059\u0065\u0073", _ggac.ToPdfObject())
	_eaed := _dd.MakeDict()
	_eaed.Set("\u004e", _faafa)
	_ede.AP = _eaed
	_ede.AS = _dd.MakeName(_fbbf)
	_bfe.Annotations = append(_bfe.Annotations, _ede)
	return _bfe, nil
}
func _afge() _dd.StreamEncoder { return _dd.NewFlateEncoder() }
func _fe(_fb CircleAnnotationDef, _dab string) ([]byte, *_fa.PdfRectangle, *_fa.PdfRectangle, error) {
	_bcb := _af.Circle{X: _fb.X, Y: _fb.Y, Width: _fb.Width, Height: _fb.Height, FillEnabled: _fb.FillEnabled, FillColor: _fb.FillColor, BorderEnabled: _fb.BorderEnabled, BorderWidth: _fb.BorderWidth, BorderColor: _fb.BorderColor, Opacity: _fb.Opacity}
	_ed, _gda, _bea := _bcb.Draw(_dab)
	if _bea != nil {
		return nil, nil, nil, _bea
	}
	_cd := &_fa.PdfRectangle{}
	_cd.Llx = _fb.X + _gda.Llx
	_cd.Lly = _fb.Y + _gda.Lly
	_cd.Urx = _fb.X + _gda.Urx
	_cd.Ury = _fb.Y + _gda.Ury
	return _ed, _gda, _cd, nil
}

// Style returns the appearance style of `fa`. If not specified, returns default style.
func (_eeaa ImageFieldAppearance) Style() AppearanceStyle {
	if _eeaa._aaf != nil {
		return *_eeaa._aaf
	}
	return AppearanceStyle{BorderSize: 0.0, BorderColor: _fa.NewPdfColorDeviceGray(0), FillColor: _fa.NewPdfColorDeviceGray(1), DrawAlignmentReticle: false}
}

func _beab(_cda *_fa.PdfAnnotationWidget, _df *_fa.PdfFieldText, _bee *_fa.PdfPageResources, _gce AppearanceStyle) (*_dd.PdfObjectDictionary, error) {
	_bfc := _fa.NewPdfPageResources()
	_ee, _eg := _dd.GetArray(_cda.Rect)
	if !_eg {
		return nil, _bd.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_fec, _gba := _fa.NewPdfRectangle(*_ee)
	if _gba != nil {
		return nil, _gba
	}
	_ebb, _aab := _fec.Width(), _fec.Height()
	_daf, _cdc := _ebb, _aab
	_fcc, _ddef := _dd.GetDict(_cda.MK)
	if _ddef {
		_ce, _ := _dd.GetDict(_cda.BS)
		_afee := _gce.applyAppearanceCharacteristics(_fcc, _ce, nil)
		if _afee != nil {
			return nil, _afee
		}
	}
	_fbb, _gba := _b.NewContentStreamParser(_beed(_df.PdfField)).Parse()
	if _gba != nil {
		return nil, _gba
	}
	_gdac := _b.NewContentCreator()
	if _gce.BorderSize > 0 {
		_afdc(_gdac, _gce, _ebb, _aab)
	}
	if _gce.DrawAlignmentReticle {
		_ddf := _gce
		_ddf.BorderSize = 0.2
		_cece(_gdac, _ddf, _ebb, _aab)
	}
	_gdac.Add_BMC("\u0054\u0078")
	_gdac.Add_q()
	_ebb, _aab = _gce.applyRotation(_fcc, _ebb, _aab, _gdac)
	_gdac.Add_BT()
	_gfd, _bae, _gba := _gce.processDA(_df.PdfField, _fbb, _bee, _bfc, _gdac)
	if _gba != nil {
		return nil, _gba
	}
	_agd := _gfd.Font
	_gdf := _gfd.Size
	_fg := _dd.MakeName(_gfd.Name)
	if _df.Flags().Has(_fa.FieldFlagMultiline) && _df.MaxLen != nil {
		_a.Log.Debug("\u004c\u006f\u006f\u006b\u0020\u0066\u006f\u0072\u0020\u0041\u0050\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0020\u0066\u006f\u0072 \u004e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006fn\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
		if _bdc, _fag, _edc := _gfbg(_cda.PdfAnnotation.AP, _bee); _edc {
			_fg = _bdc
			_gdf = _fag
			_bae = true
		}
	}
	_db := _gdf == 0
	if _db && _bae {
		_gdf = _aab * _gce.AutoFontSizeFraction
	}
	_ca := _agd.Encoder()
	if _ca == nil {
		_a.Log.Debug("\u0057\u0041RN\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0069\u0073\u0020\u006e\u0069l\u002e\u0020\u0041\u0073s\u0075\u006d\u0069\u006eg \u0069\u0064e\u006et\u0069\u0074\u0079\u0020\u0065\u006ec\u006f\u0064\u0065r\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
		_ca = _d.NewIdentityTextEncoder("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048")
	}
	_efc, _gba := _agd.GetFontDescriptor()
	if _gba != nil {
		_a.Log.Debug("\u0045\u0072ro\u0072\u003a\u0020U\u006e\u0061\u0062\u006ce t\u006f g\u0065\u0074\u0020\u0066\u006f\u006e\u0074 d\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
	}
	var _bef string
	if _gde, _ebbf := _dd.GetString(_df.V); _ebbf {
		_bef = _gde.Decoded()
	}
	if len(_bef) == 0 {
		return nil, nil
	}
	_afc := []string{_bef}
	_fae := false
	if _df.Flags().Has(_fa.FieldFlagMultiline) {
		_fae = true
		_bef = _ge.Replace(_bef, "\u000d\u000a", "\u000a", -1)
		_bef = _ge.Replace(_bef, "\u000d", "\u000a", -1)
		_afc = _ge.Split(_bef, "\u000a")
	}
	_dae := make([]string, len(_afc))
	copy(_dae, _afc)
	_ccc := _gce.MultilineLineHeight
	_gbb := 0.0
	_cbg := 0
	if _ca != nil {
		for _gdf >= 0 {
			_dg := make([]string, len(_afc))
			copy(_dg, _afc)
			_bda := make([]string, len(_dae))
			copy(_bda, _dae)
			_gbb = 0.0
			_cbg = 0
			_gbef := len(_dg)
			_afec := 0
			for _afec < _gbef {
				var _efb float64
				_abe := -1
				_dgf := _aeb
				if _gce.MarginLeft != nil {
					_dgf = *_gce.MarginLeft
				}
				for _aff, _eae := range _dg[_afec] {
					if _eae == ' ' {
						_abe = _aff
					}
					_acg, _fcd := _agd.GetRuneMetrics(_eae)
					if !_fcd {
						_a.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u006d\u0065\u0074r\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u002d\u0020\u0073k\u0069p\u0070\u0069\u006e\u0067", _eae)
						continue
					}
					_efb = _dgf
					_dgf += _acg.Wx
					if _fae && !_db && _gdf*_dgf/1000.0 > _ebb {
						_agb := _aff
						_ebd := _aff
						if _abe > 0 {
							_agb = _abe + 1
							_ebd = _abe
						}
						_ffa := _dg[_afec][_agb:]
						_cfd := _bda[_afec][_agb:]
						if _afec < len(_dg)-1 {
							_dg = append(_dg[:_afec+1], _dg[_afec:]...)
							_dg[_afec+1] = _ffa
							_bda = append(_bda[:_afec+1], _bda[_afec:]...)
							_bda[_afec+1] = _cfd
						} else {
							_dg = append(_dg, _ffa)
							_bda = append(_bda, _cfd)
						}
						_dg[_afec] = _dg[_afec][0:_ebd]
						_bda[_afec] = _bda[_afec][0:_ebd]
						_gbef++
						_dgf = _efb
						break
					}
				}
				if _dgf > _gbb {
					_gbb = _dgf
				}
				_dg[_afec] = string(_ca.Encode(_dg[_afec]))
				if len(_dg[_afec]) > 0 {
					_cbg++
				}
				_afec++
			}
			_aaa := _gdf
			if _cbg > 1 {
				_aaa *= _ccc
			}
			_bad := float64(_cbg) * _aaa
			if _db || _bad <= _aab {
				_afc = _dg
				_dae = _bda
				break
			}
			_gdf--
		}
	}
	_ddab := _aeb
	if _gce.MarginLeft != nil {
		_ddab = *_gce.MarginLeft
	}
	if _gdf == 0 || _db && _gbb > 0 && _ddab+_gbb*_gdf/1000.0 > _ebb {
		_gdf = 0.95 * 1000.0 * (_ebb - _ddab) / _gbb
	}
	_bgc := _bag
	{
		if _gfb, _fdb := _dd.GetIntVal(_df.Q); _fdb {
			switch _gfb {
			case 0:
				_bgc = _bag
			case 1:
				_bgc = _bb
			case 2:
				_bgc = _ebg
			default:
				_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0071\u0075\u0061\u0064\u0064\u0069\u006e\u0067\u003a\u0020%\u0064\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u006c\u0065ft\u0020\u0061\u006c\u0069\u0067\u006e\u006d\u0065\u006e\u0074", _gfb)
			}
		}
	}
	_bgg := _gdf
	if _fae && _cbg > 1 {
		_bgg = _ccc * _gdf
	}
	var _eab float64
	if _efc != nil {
		_eab, _gba = _efc.GetCapHeight()
		if _gba != nil {
			_a.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0043\u0061\u0070\u0048\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _gba)
		}
	}
	if int(_eab) <= 0 {
		_a.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u0061\u0070\u0048e\u0069\u0067\u0068\u0074\u0020\u006e\u006ft \u0061\u0076\u0061\u0069l\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065tt\u0069\u006eg\u0020\u0074\u006f\u0020\u0031\u0030\u0030\u0030")
		_eab = 1000
	}
	_fcf := _eab / 1000.0 * _gdf
	_fee := 0.0
	{
		_eeb := float64(_cbg) * _bgg
		if _db && _fee+_eeb > _aab {
			_gdf = 0.95 * (_aab - _fee) / float64(_cbg)
			_bgg = _gdf
			if _fae && _cbg > 1 {
				_bgg = _ccc * _gdf
			}
			_fcf = _eab / 1000.0 * _gdf
			_eeb = float64(_cbg) * _bgg
		}
		if _aab > _eeb {
			if _fae {
				if _gce.MultilineVAlignMiddle {
					_gefe := (_aab - (_eeb + _fcf)) / 2.0
					_ga := _gefe + _eeb + _fcf - _bgg
					_fee = _ga
					if _cbg > 1 {
						_fee = _fee + (_eeb / _gdf * float64(_cbg)) - _bgg - _fcf
					}
					if _fee < _eeb {
						_fee = (_aab - _fcf) / 2.0
					}
				} else {
					_fee = _aab - _bgg
					if _fee > _gdf {
						_cdd := 0.0
						if _fae && _gce.MultilineLineHeight > 1 && _cbg > 1 {
							_cdd = _gce.MultilineLineHeight - 1
						}
						_fee -= _gdf * (0.5 - _cdd)
					}
				}
			} else {
				_fee = (_aab - _fcf) / 2.0
			}
		}
	}
	_gdac.Add_Tf(*_fg, _gdf)
	_gdac.Add_Td(_ddab, _fee)
	_dca := _ddab
	_ffb := _ddab
	for _cg, _dcb := range _afc {
		_gac := 0.0
		for _, _dff := range _dae[_cg] {
			_dce, _acf := _agd.GetRuneMetrics(_dff)
			if !_acf {
				continue
			}
			_gac += _dce.Wx
		}
		_bba := _gac / 1000.0 * _gdf
		_ece := _ebb - _bba
		var _gdb float64
		switch _bgc {
		case _bag:
			_gdb = _dca
		case _bb:
			_gdb = _ece / 2
		case _ebg:
			_gdb = _ece
		}
		_ddab = _gdb - _ffb
		if _ddab > 0.0 {
			_gdac.Add_Td(_ddab, 0)
		}
		_ffb = _gdb
		_gdac.Add_Tj(*_dd.MakeString(_dcb))
		if _cg < len(_afc)-1 {
			_gdac.Add_Td(0, -_gdf*_ccc)
		}
	}
	_gdac.Add_ET()
	_gdac.Add_Q()
	_gdac.Add_EMC()
	_bfb := _fa.NewXObjectForm()
	_bfb.Resources = _bfc
	_bfb.BBox = _dd.MakeArrayFromFloats([]float64{0, 0, _daf, _cdc})
	_bfb.SetContentStream(_gdac.Bytes(), _afge())
	_edd := _dd.MakeDict()
	_edd.Set("\u004e", _bfb.ToPdfObject())
	return _edd, nil
}

func _cdgd(_adgc []float64) []float64 {
	var (
		_gdd  = len(_adgc)
		_eggd = make([]float64, _gdd)
		_cffc = make([]float64, _gdd)
	)
	_cgea := 2.0
	_eggd[0] = _adgc[0] / _cgea
	for _fade := 1; _fade < _gdd; _fade++ {
		_cffc[_fade] = 1 / _cgea
		if _fade < _gdd-1 {
			_cgea = 4.0
		} else {
			_cgea = 3.5
		}
		_cgea -= _cffc[_fade]
		_eggd[_fade] = (_adgc[_fade] - _eggd[_fade-1]) / _cgea
	}
	for _caac := 1; _caac < _gdd; _caac++ {
		_eggd[_gdd-_caac-1] -= _cffc[_gdd-_caac] * _eggd[_gdd-_caac]
	}
	return _eggd
}

// NewTextField generates a new text field with partial name `name` at location
// specified by `rect` on given `page` and with field specific options `opt`.
func NewTextField(page *_fa.PdfPage, name string, rect []float64, opt TextFieldOptions) (*_fa.PdfFieldText, error) {
	if page == nil {
		return nil, _bd.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _bd.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _bd.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_gad := _fa.NewPdfField()
	_dbd := &_fa.PdfFieldText{}
	_gad.SetContext(_dbd)
	_dbd.PdfField = _gad
	_dbd.T = _dd.MakeString(name)
	if opt.MaxLen > 0 {
		_dbd.MaxLen = _dd.MakeInteger(int64(opt.MaxLen))
	}
	if len(opt.Value) > 0 {
		_dbd.V = _dd.MakeString(opt.Value)
	}
	_ccgf := _fa.NewPdfAnnotationWidget()
	_ccgf.Rect = _dd.MakeArrayFromFloats(rect)
	_ccgf.P = page.ToPdfObject()
	_ccgf.F = _dd.MakeInteger(4)
	_ccgf.Parent = _dbd.ToPdfObject()
	_dbd.Annotations = append(_dbd.Annotations, _ccgf)
	return _dbd, nil
}

// Style returns the appearance style of `fa`. If not specified, returns default style.
func (_dde FieldAppearance) Style() AppearanceStyle {
	if _dde._eb != nil {
		return *_dde._eb
	}
	_gbe := _aeb
	return AppearanceStyle{AutoFontSizeFraction: 0.65, CheckmarkRune: '✔', BorderSize: 0.0, BorderColor: _fa.NewPdfColorDeviceGray(0), FillColor: _fa.NewPdfColorDeviceGray(1), MultilineLineHeight: 1.2, MultilineVAlignMiddle: false, DrawAlignmentReticle: false, AllowMK: true, MarginLeft: &_gbe}
}

// NewSignatureField returns a new signature field with a visible appearance
// containing the specified signature lines and styled according to the
// specified options.
func NewSignatureField(signature *_fa.PdfSignature, lines []*SignatureLine, opts *SignatureFieldOpts) (*_fa.PdfFieldSignature, error) {
	if signature == nil {
		return nil, _bd.New("\u0073\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	_egge, _egfe := _dace(lines, opts)
	if _egfe != nil {
		return nil, _egfe
	}
	_bcfbg := _fa.NewPdfFieldSignature(signature)
	_bcfbg.Rect = _dd.MakeArrayFromFloats(opts.Rect)
	_bcfbg.AP = _egge
	return _bcfbg, nil
}

// CircleAnnotationDef defines a circle annotation or ellipse at position (X, Y) and Width and Height.
// The annotation has various style parameters including Fill and Border options and Opacity.
type CircleAnnotationDef struct {
	X             float64
	Y             float64
	Width         float64
	Height        float64
	FillEnabled   bool
	FillColor     *_fa.PdfColorDeviceRGB
	BorderEnabled bool
	BorderWidth   float64
	BorderColor   *_fa.PdfColorDeviceRGB
	Opacity       float64
}

const (
	SignatureImageLeft SignatureImagePosition = iota
	SignatureImageRight
	SignatureImageTop
	SignatureImageBottom
)

const (
	_bfee  = 1
	_fgb   = 2
	_gbed  = 4
	_cbga  = 8
	_bagc  = 16
	_agfcc = 32
	_dbga  = 64
	_gabe  = 128
	_dbe   = 256
	_ebdc  = 512
	_efgg  = 1024
	_ecde  = 2048
	_febc  = 4096
)

func _dgce(_gabd, _gfe float64, _efbe *_fa.Image, _aaac AppearanceStyle) (*_fa.XObjectForm, error) {
	_dgdd, _dgcc := _fa.NewXObjectImageFromImage(_efbe, nil, _dd.NewFlateEncoder())
	if _dgcc != nil {
		return nil, _dgcc
	}
	_dgdd.Decode = _dd.MakeArrayFromFloats([]float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0})
	_ffaa := _fa.NewPdfPageResources()
	_ffaa.ProcSet = _dd.MakeArray(_dd.MakeName("\u0050\u0044\u0046"), _dd.MakeName("\u0049\u006d\u0061\u0067\u0065\u0043"))
	_ffaa.SetXObjectImageByName(_dd.PdfObjectName("\u0049\u006d\u0030"), _dgdd)
	_cfce := _b.NewContentCreator()
	_cfce.Add_q()
	_cfce.Add_cm(float64(_efbe.Width), 0, 0, float64(_efbe.Height), 0, 0)
	_cfce.Add_Do("\u0049\u006d\u0030")
	_cfce.Add_Q()
	_dgab := _fa.NewXObjectForm()
	_dgab.FormType = _dd.MakeInteger(1)
	_dgab.BBox = _dd.MakeArrayFromFloats([]float64{0, 0, float64(_efbe.Width), float64(_efbe.Height)})
	_dgab.Resources = _ffaa
	_dgab.SetContentStream(_cfce.Bytes(), _afge())
	return _dgab, nil
}

func _gea(_faaf _e.Image, _bfbg string, _agfb *SignatureFieldOpts, _efg []float64, _gfga *_b.ContentCreator) (*_dd.PdfObjectName, *_fa.XObjectImage, error) {
	_eag, _aged := _fa.DefaultImageHandler{}.NewImageFromGoImage(_faaf)
	if _aged != nil {
		return nil, nil, _aged
	}
	_cff, _aged := _fa.NewXObjectImageFromImage(_eag, nil, _agfb.Encoder)
	if _aged != nil {
		return nil, nil, _aged
	}
	_agbf, _cbdd := float64(*_cff.Width), float64(*_cff.Height)
	_adg := _efg[2] - _efg[0]
	_adee := _efg[3] - _efg[1]
	if _agfb.AutoSize {
		_dacd := _cf.Min(_adg/_agbf, _adee/_cbdd)
		_agbf *= _dacd
		_cbdd *= _dacd
		_efg[0] = _efg[0] + (_adg / 2) - (_agbf / 2)
		_efg[1] = _efg[1] + (_adee / 2) - (_cbdd / 2)
	}
	var _aefe *_dd.PdfObjectName
	if _ddcc, _fabgd := _dd.GetName(_cff.Name); _fabgd {
		_aefe = _ddcc
	} else {
		_aefe = _dd.MakeName(_bfbg)
	}
	if _gfga != nil {
		_gfga.Add_q().Translate(_efg[0], _efg[1]).Scale(_agbf, _cbdd).Add_Do(*_aefe).Add_Q()
	} else {
		return nil, nil, _bd.New("\u0043\u006f\u006e\u0074en\u0074\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u0075l\u006c")
	}
	return _aefe, _cff, nil
}

// CreateCircleAnnotation creates a circle/ellipse annotation object with appearance stream that can be added to
// page PDF annotations.
func CreateCircleAnnotation(circDef CircleAnnotationDef) (*_fa.PdfAnnotation, error) {
	_ac := _fa.NewPdfAnnotationCircle()
	if circDef.BorderEnabled {
		_bc, _gf, _da := circDef.BorderColor.R(), circDef.BorderColor.G(), circDef.BorderColor.B()
		_ac.C = _dd.MakeArrayFromFloats([]float64{_bc, _gf, _da})
		_ba := _fa.NewBorderStyle()
		_ba.SetBorderWidth(circDef.BorderWidth)
		_ac.BS = _ba.ToPdfObject()
	}
	if circDef.FillEnabled {
		_geg, _ae, _dda := circDef.FillColor.R(), circDef.FillColor.G(), circDef.FillColor.B()
		_ac.IC = _dd.MakeArrayFromFloats([]float64{_geg, _ae, _dda})
	} else {
		_ac.IC = _dd.MakeArrayFromIntegers([]int{})
	}
	if circDef.Opacity < 1.0 {
		_ac.CA = _dd.MakeFloat(circDef.Opacity)
	}
	_gb, _be, _gee := _ab(circDef)
	if _gee != nil {
		return nil, _gee
	}
	_ac.AP = _gb
	_ac.Rect = _dd.MakeArrayFromFloats([]float64{_be.Llx, _be.Lly, _be.Urx, _be.Ury})
	return _ac.PdfAnnotation, nil
}

// SetStyle applies appearance `style` to `fa`.
func (_aba *FieldAppearance) SetStyle(style AppearanceStyle) { _aba._eb = &style }
