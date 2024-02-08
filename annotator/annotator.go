package annotator

import (
	_c "bytes"
	_f "errors"
	_bga "image"
	_dg "math"
	_g "strings"
	_ad "unicode"

	_d "bitbucket.org/shenghui0779/gopdf/common"
	_bg "bitbucket.org/shenghui0779/gopdf/contentstream"
	_fb "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_cd "bitbucket.org/shenghui0779/gopdf/core"
	_a "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_bb "bitbucket.org/shenghui0779/gopdf/model"
)

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
	_ca                  *AppearanceStyle
}

// TextFieldOptions defines optional parameter for a text field in a form.
type TextFieldOptions struct {
	MaxLen int
	Value  string
}

// NewFormResetButtonField would create a reset button in specified page according to the parameter in `FormResetActionOptions`.
func NewFormResetButtonField(page *_bb.PdfPage, opt FormResetActionOptions) (*_bb.PdfFieldButton, error) {
	_eaaa := _bb.NewPdfActionResetForm()
	_eaaa.Fields = opt.Fields
	_eaaa.Flags = _cd.MakeInteger(0)
	if opt.IsExclusionList {
		_eaaa.Flags = _cd.MakeInteger(1)
	}
	_dgeg, _acb := _ecc(page, opt.Rectangle, "\u0062\u0074\u006e\u0052\u0065\u0073\u0065\u0074", opt.Label, opt.LabelColor, opt.Font, opt.FontSize, _eaaa.ToPdfObject())
	if _acb != nil {
		return nil, _acb
	}
	return _dgeg, nil
}

// SignatureImagePosition specifies the image signature location relative to the text signature.
// If text signature is not defined, this position will be ignored.
type SignatureImagePosition int

// GenerateAppearanceDict generates an appearance dictionary for widget annotation `wa` for the `field` in `form`.
// Implements interface model.FieldAppearanceGenerator.
func (_bbg FieldAppearance) GenerateAppearanceDict(form *_bb.PdfAcroForm, field *_bb.PdfField, wa *_bb.PdfAnnotationWidget) (*_cd.PdfObjectDictionary, error) {
	_d.Log.Trace("\u0047\u0065n\u0065\u0072\u0061\u0074e\u0041\u0070p\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0044i\u0063\u0074\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u0020\u0056:\u0020\u0025\u002b\u0076", field.PartialName(), field.V)
	_, _cdc := field.GetContext().(*_bb.PdfFieldText)
	_aggb, _eg := _cd.GetDict(wa.AP)
	if _eg && _bbg.OnlyIfMissing && (!_cdc || !_bbg.RegenerateTextFields) {
		_d.Log.Trace("\u0041\u006c\u0072\u0065a\u0064\u0079\u0020\u0070\u006f\u0070\u0075\u006c\u0061\u0074e\u0064 \u002d\u0020\u0069\u0067\u006e\u006f\u0072i\u006e\u0067")
		return _aggb, nil
	}
	if form.DR == nil {
		form.DR = _bb.NewPdfPageResources()
	}
	switch _bae := field.GetContext().(type) {
	case *_bb.PdfFieldText:
		_eb := _bae
		if _ef := _dbf(_eb.PdfField); _ef == "" {
			_eb.DA = form.DA
		}
		switch {
		case _eb.Flags().Has(_bb.FieldFlagPassword):
			return nil, nil
		case _eb.Flags().Has(_bb.FieldFlagFileSelect):
			return nil, nil
		case _eb.Flags().Has(_bb.FieldFlagComb):
			if _eb.MaxLen != nil {
				_bge, _cac := _fec(wa, _eb, form.DR, _bbg.Style())
				if _cac != nil {
					return nil, _cac
				}
				return _bge, nil
			}
		}
		_egd, _egdb := _df(wa, _eb, form.DR, _bbg.Style())
		if _egdb != nil {
			return nil, _egdb
		}
		return _egd, nil
	case *_bb.PdfFieldButton:
		_ga := _bae
		if _ga.IsCheckbox() {
			_dgb, _cae := _gbaf(wa, _ga, form.DR, _bbg.Style())
			if _cae != nil {
				return nil, _cae
			}
			return _dgb, nil
		}
		_d.Log.Debug("\u0054\u004f\u0044\u004f\u003a\u0020\u0055\u004e\u0048\u0041\u004e\u0044\u004c\u0045\u0044 \u0062u\u0074\u0074\u006f\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u002b\u0076", _ga.GetType())
	case *_bb.PdfFieldChoice:
		_adg := _bae
		switch {
		case _adg.Flags().Has(_bb.FieldFlagCombo):
			_cgd, _efb := _fggc(form, wa, _adg, _bbg.Style())
			if _efb != nil {
				return nil, _efb
			}
			return _cgd, nil
		default:
			_d.Log.Debug("\u0054\u004f\u0044\u004f\u003a\u0020\u0055N\u0048\u0041\u004eD\u004c\u0045\u0044\u0020c\u0068\u006f\u0069\u0063\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0077\u0069\u0074\u0068\u0020\u0066\u006c\u0061\u0067\u0073\u003a\u0020\u0025\u0073", _adg.Flags().String())
		}
	default:
		_d.Log.Debug("\u0054\u004f\u0044\u004f\u003a\u0020\u0055\u004e\u0048\u0041N\u0044\u004c\u0045\u0044\u0020\u0066\u0069e\u006c\u0064\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _bae)
	}
	return nil, nil
}
func _e(_gd CircleAnnotationDef) (*_cd.PdfObjectDictionary, *_bb.PdfRectangle, error) {
	_adc := _bb.NewXObjectForm()
	_adc.Resources = _bb.NewPdfPageResources()
	_agg := ""
	if _gd.Opacity < 1.0 {
		_fg := _cd.MakeDict()
		_fg.Set("\u0063\u0061", _cd.MakeFloat(_gd.Opacity))
		_fg.Set("\u0043\u0041", _cd.MakeFloat(_gd.Opacity))
		_cg := _adc.Resources.AddExtGState("\u0067\u0073\u0031", _fg)
		if _cg != nil {
			_d.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020\u0065\u0078\u0074g\u0073\u0074\u0061t\u0065 \u0067\u0073\u0031")
			return nil, nil, _cg
		}
		_agg = "\u0067\u0073\u0031"
	}
	_gcb, _fed, _fda, _gf := _ffb(_gd, _agg)
	if _gf != nil {
		return nil, nil, _gf
	}
	_gf = _adc.SetContentStream(_gcb, nil)
	if _gf != nil {
		return nil, nil, _gf
	}
	_adc.BBox = _fed.ToPdfObject()
	_de := _cd.MakeDict()
	_de.Set("\u004e", _adc.ToPdfObject())
	return _de, _fda, nil
}
func _dfbg(_fdfe RectangleAnnotationDef, _baee string) ([]byte, *_bb.PdfRectangle, *_bb.PdfRectangle, error) {
	_egeb := _fb.Rectangle{X: 0, Y: 0, Width: _fdfe.Width, Height: _fdfe.Height, FillEnabled: _fdfe.FillEnabled, FillColor: _fdfe.FillColor, BorderEnabled: _fdfe.BorderEnabled, BorderWidth: 2 * _fdfe.BorderWidth, BorderColor: _fdfe.BorderColor, Opacity: _fdfe.Opacity}
	_gdac, _agcg, _cace := _egeb.Draw(_baee)
	if _cace != nil {
		return nil, nil, nil, _cace
	}
	_fde := &_bb.PdfRectangle{}
	_fde.Llx = _fdfe.X + _agcg.Llx
	_fde.Lly = _fdfe.Y + _agcg.Lly
	_fde.Urx = _fdfe.X + _agcg.Urx
	_fde.Ury = _fdfe.Y + _agcg.Ury
	return _gdac, _agcg, _fde, nil
}

const (
	SignatureImageLeft SignatureImagePosition = iota
	SignatureImageRight
	SignatureImageTop
	SignatureImageBottom
)

func _eaed(_bea LineAnnotationDef) (*_cd.PdfObjectDictionary, *_bb.PdfRectangle, error) {
	_fcgag := _bb.NewXObjectForm()
	_fcgag.Resources = _bb.NewPdfPageResources()
	_fbca := ""
	if _bea.Opacity < 1.0 {
		_gced := _cd.MakeDict()
		_gced.Set("\u0063\u0061", _cd.MakeFloat(_bea.Opacity))
		_fgde := _fcgag.Resources.AddExtGState("\u0067\u0073\u0031", _gced)
		if _fgde != nil {
			_d.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020\u0065\u0078\u0074g\u0073\u0074\u0061t\u0065 \u0067\u0073\u0031")
			return nil, nil, _fgde
		}
		_fbca = "\u0067\u0073\u0031"
	}
	_feaf, _ggga, _feag, _bca := _cefbf(_bea, _fbca)
	if _bca != nil {
		return nil, nil, _bca
	}
	_bca = _fcgag.SetContentStream(_feaf, nil)
	if _bca != nil {
		return nil, nil, _bca
	}
	_fcgag.BBox = _ggga.ToPdfObject()
	_cfdd := _cd.MakeDict()
	_cfdd.Set("\u004e", _fcgag.ToPdfObject())
	return _cfdd, _feag, nil
}

// ImageFieldAppearance implements interface model.FieldAppearanceGenerator and generates appearance streams
// for attaching an image to a button field.
type ImageFieldAppearance struct {
	OnlyIfMissing bool
	_aade         *AppearanceStyle
}

// NewCheckboxField generates a new checkbox field with partial name `name` at location `rect`
// on specified `page` and with field specific options `opt`.
func NewCheckboxField(page *_bb.PdfPage, name string, rect []float64, opt CheckboxFieldOptions) (*_bb.PdfFieldButton, error) {
	if page == nil {
		return nil, _f.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _f.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _f.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_ffee, _bfc := _bb.NewStandard14Font(_bb.ZapfDingbatsName)
	if _bfc != nil {
		return nil, _bfc
	}
	_eabe := _bb.NewPdfField()
	_gcca := &_bb.PdfFieldButton{}
	_eabe.SetContext(_gcca)
	_gcca.PdfField = _eabe
	_gcca.T = _cd.MakeString(name)
	_gcca.SetType(_bb.ButtonTypeCheckbox)
	_bcf := "\u004f\u0066\u0066"
	if opt.Checked {
		_bcf = "\u0059\u0065\u0073"
	}
	_gcca.V = _cd.MakeName(_bcf)
	_cca := _bb.NewPdfAnnotationWidget()
	_cca.Rect = _cd.MakeArrayFromFloats(rect)
	_cca.P = page.ToPdfObject()
	_cca.F = _cd.MakeInteger(4)
	_cca.Parent = _gcca.ToPdfObject()
	_fcdc := rect[2] - rect[0]
	_bdcd := rect[3] - rect[1]
	var _gcgc _c.Buffer
	_gcgc.WriteString("\u0071\u000a")
	_gcgc.WriteString("\u0030 \u0030\u0020\u0031\u0020\u0072\u0067\n")
	_gcgc.WriteString("\u0042\u0054\u000a")
	_gcgc.WriteString("\u002f\u005a\u0061D\u0062\u0020\u0031\u0032\u0020\u0054\u0066\u000a")
	_gcgc.WriteString("\u0045\u0054\u000a")
	_gcgc.WriteString("\u0051\u000a")
	_acgd := _bg.NewContentCreator()
	_acgd.Add_q()
	_acgd.Add_rg(0, 0, 1)
	_acgd.Add_BT()
	_acgd.Add_Tf(*_cd.MakeName("\u005a\u0061\u0044\u0062"), 12)
	_acgd.Add_Td(0, 0)
	_acgd.Add_ET()
	_acgd.Add_Q()
	_dceg := _bb.NewXObjectForm()
	_dceg.SetContentStream(_acgd.Bytes(), _cd.NewRawEncoder())
	_dceg.BBox = _cd.MakeArrayFromFloats([]float64{0, 0, _fcdc, _bdcd})
	_dceg.Resources = _bb.NewPdfPageResources()
	_dceg.Resources.SetFontByName("\u005a\u0061\u0044\u0062", _ffee.ToPdfObject())
	_acgd = _bg.NewContentCreator()
	_acgd.Add_q()
	_acgd.Add_re(0, 0, _fcdc, _bdcd)
	_acgd.Add_W().Add_n()
	_acgd.Add_rg(0, 0, 1)
	_acgd.Translate(0, 3.0)
	_acgd.Add_BT()
	_acgd.Add_Tf(*_cd.MakeName("\u005a\u0061\u0044\u0062"), 12)
	_acgd.Add_Td(0, 0)
	_acgd.Add_Tj(*_cd.MakeString("\u0034"))
	_acgd.Add_ET()
	_acgd.Add_Q()
	_adff := _bb.NewXObjectForm()
	_adff.SetContentStream(_acgd.Bytes(), _cd.NewRawEncoder())
	_adff.BBox = _cd.MakeArrayFromFloats([]float64{0, 0, _fcdc, _bdcd})
	_adff.Resources = _bb.NewPdfPageResources()
	_adff.Resources.SetFontByName("\u005a\u0061\u0044\u0062", _ffee.ToPdfObject())
	_aff := _cd.MakeDict()
	_aff.Set("\u004f\u0066\u0066", _dceg.ToPdfObject())
	_aff.Set("\u0059\u0065\u0073", _adff.ToPdfObject())
	_dgag := _cd.MakeDict()
	_dgag.Set("\u004e", _aff)
	_cca.AP = _dgag
	_cca.AS = _cd.MakeName(_bcf)
	_gcca.Annotations = append(_gcca.Annotations, _cca)
	return _gcca, nil
}
func _ecc(_dacg *_bb.PdfPage, _effe _fb.Rectangle, _gaba string, _afce string, _eaea _bb.PdfColor, _gfgg *_bb.PdfFont, _gcee *float64, _eeeg _cd.PdfObject) (*_bb.PdfFieldButton, error) {
	_abd, _dgad := _effe.X, _effe.Y
	_dabd := _effe.Width
	_eeee := _effe.Height
	if _effe.FillColor == nil {
		_effe.FillColor = _bb.NewPdfColorDeviceGray(0.7)
	}
	if _eaea == nil {
		_eaea = _bb.NewPdfColorDeviceGray(0)
	}
	if _gfgg == nil {
		_cded, _bdbe := _bb.NewStandard14Font("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
		if _bdbe != nil {
			return nil, _bdbe
		}
		_gfgg = _cded
	}
	_bage := _bb.NewPdfField()
	_fgcc := &_bb.PdfFieldButton{}
	_bage.SetContext(_fgcc)
	_fgcc.PdfField = _bage
	_fgcc.T = _cd.MakeString(_gaba)
	_fgcc.SetType(_bb.ButtonTypePush)
	_fgcc.V = _cd.MakeName("\u004f\u0066\u0066")
	_fgcc.Ff = _cd.MakeInteger(4)
	_abga := _cd.MakeDict()
	_abga.Set(*_cd.MakeName("\u0043\u0041"), _cd.MakeString(_afce))
	_fecf, _geaf := _gfgg.GetFontDescriptor()
	if _geaf != nil {
		return nil, _geaf
	}
	_aed := _cd.MakeName("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
	_faaf := 12.0
	if _fecf != nil && _fecf.FontName != nil {
		_aed, _ = _cd.GetName(_fecf.FontName)
	}
	if _gcee != nil {
		_faaf = *_gcee
	}
	_gdagf := _bg.NewContentCreator()
	_gdagf.Add_q()
	_gdagf.SetNonStrokingColor(_effe.FillColor)
	_gdagf.Add_re(0, 0, _dabd, _eeee)
	_gdagf.Add_f()
	_gdagf.Add_Q()
	_gdagf.Add_q()
	_gdagf.Add_BT()
	_debf := 0.0
	for _, _bbgfe := range _afce {
		_affc, _bbdf := _gfgg.GetRuneMetrics(_bbgfe)
		if !_bbdf {
			_d.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u006d\u0065\u0074r\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u002d\u0020\u0073k\u0069p\u0070\u0069\u006e\u0067", _bbgfe)
			continue
		}
		_debf += _affc.Wx
	}
	_debf = _debf / 1000.0 * _faaf
	var _bab float64
	if _fecf != nil {
		_bab, _geaf = _fecf.GetCapHeight()
		if _geaf != nil {
			_d.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0043\u0061\u0070\u0048\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _geaf)
		}
	}
	if int(_bab) <= 0 {
		_d.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u0061\u0070\u0048e\u0069\u0067\u0068\u0074\u0020\u006e\u006ft \u0061\u0076\u0061\u0069l\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065tt\u0069\u006eg\u0020\u0074\u006f\u0020\u0031\u0030\u0030\u0030")
		_bab = 1000
	}
	_efba := _bab / 1000.0 * _faaf
	_dggb := (_eeee - _efba) / 2.0
	_fac := (_dabd - _debf) / 2.0
	_gdagf.Add_Tf(*_aed, _faaf)
	_gdagf.SetNonStrokingColor(_eaea)
	_gdagf.Add_Td(_fac, _dggb)
	_gdagf.Add_Tj(*_cd.MakeString(_afce))
	_gdagf.Add_ET()
	_gdagf.Add_Q()
	_deed := _bb.NewXObjectForm()
	_deed.SetContentStream(_gdagf.Bytes(), _cd.NewRawEncoder())
	_deed.BBox = _cd.MakeArrayFromFloats([]float64{0, 0, _dabd, _eeee})
	_deed.Resources = _bb.NewPdfPageResources()
	_deed.Resources.SetFontByName(*_aed, _gfgg.ToPdfObject())
	_gcec := _cd.MakeDict()
	_gcec.Set("\u004e", _deed.ToPdfObject())
	_ceg := _bb.NewPdfAnnotationWidget()
	_ceg.Rect = _cd.MakeArrayFromFloats([]float64{_abd, _dgad, _abd + _dabd, _dgad + _eeee})
	_ceg.P = _dacg.ToPdfObject()
	_ceg.F = _cd.MakeInteger(4)
	_ceg.Parent = _fgcc.ToPdfObject()
	_ceg.A = _eeeg
	_ceg.MK = _abga
	_ceg.AP = _gcec
	_fgcc.Annotations = append(_fgcc.Annotations, _ceg)
	return _fgcc, nil
}

// CircleAnnotationDef defines a circle annotation or ellipse at position (X, Y) and Width and Height.
// The annotation has various style parameters including Fill and Border options and Opacity.
type CircleAnnotationDef struct {
	X             float64
	Y             float64
	Width         float64
	Height        float64
	FillEnabled   bool
	FillColor     *_bb.PdfColorDeviceRGB
	BorderEnabled bool
	BorderWidth   float64
	BorderColor   *_bb.PdfColorDeviceRGB
	Opacity       float64
}

func _df(_fc *_bb.PdfAnnotationWidget, _da *_bb.PdfFieldText, _cda *_bb.PdfPageResources, _dae AppearanceStyle) (*_cd.PdfObjectDictionary, error) {
	_gcbc := _bb.NewPdfPageResources()
	_dac, _age := _cd.GetArray(_fc.Rect)
	if !_age {
		return nil, _f.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_aec, _dace := _bb.NewPdfRectangle(*_dac)
	if _dace != nil {
		return nil, _dace
	}
	_gfb, _fgf := _aec.Width(), _aec.Height()
	_dfa, _ec := _gfb, _fgf
	_egb, _egde := _cd.GetDict(_fc.MK)
	if _egde {
		_gg, _ := _cd.GetDict(_fc.BS)
		_fgd := _dae.applyAppearanceCharacteristics(_egb, _gg, nil)
		if _fgd != nil {
			return nil, _fgd
		}
	}
	_bgg, _dace := _bg.NewContentStreamParser(_dbf(_da.PdfField)).Parse()
	if _dace != nil {
		return nil, _dace
	}
	_bf := _bg.NewContentCreator()
	if _dae.BorderSize > 0 {
		_ddc(_bf, _dae, _gfb, _fgf)
	}
	if _dae.DrawAlignmentReticle {
		_db := _dae
		_db.BorderSize = 0.2
		_deggd(_bf, _db, _gfb, _fgf)
	}
	_bf.Add_BMC("\u0054\u0078")
	_bf.Add_q()
	_gfb, _fgf = _dae.applyRotation(_egb, _gfb, _fgf, _bf)
	_bf.Add_BT()
	_ggc, _fce, _dace := _dae.processDA(_da.PdfField, _bgg, _cda, _gcbc, _bf)
	if _dace != nil {
		return nil, _dace
	}
	_dag := _ggc.Font
	_ea := _ggc.Size
	_af := _cd.MakeName(_ggc.Name)
	if _da.Flags().Has(_bb.FieldFlagMultiline) && _da.MaxLen != nil {
		_d.Log.Debug("\u004c\u006f\u006f\u006b\u0020\u0066\u006f\u0072\u0020\u0041\u0050\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u0020\u0066\u006f\u0072 \u004e\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006fn\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
		if _fbd, _gec, _bag := _fecd(_fc.PdfAnnotation.AP, _cda); _bag {
			_af = _fbd
			_ea = _gec
			_fce = true
		}
	}
	_gda := _ea == 0
	if _gda && _fce {
		_ea = _fgf * _dae.AutoFontSizeFraction
	}
	_cbg := _dag.Encoder()
	if _cbg == nil {
		_d.Log.Debug("\u0057\u0041RN\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0069\u0073\u0020\u006e\u0069l\u002e\u0020\u0041\u0073s\u0075\u006d\u0069\u006eg \u0069\u0064e\u006et\u0069\u0074\u0079\u0020\u0065\u006ec\u006f\u0064\u0065r\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
		_cbg = _a.NewIdentityTextEncoder("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048")
	}
	_ee, _dace := _dag.GetFontDescriptor()
	if _dace != nil {
		_d.Log.Debug("\u0045\u0072ro\u0072\u003a\u0020U\u006e\u0061\u0062\u006ce t\u006f g\u0065\u0074\u0020\u0066\u006f\u006e\u0074 d\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
	}
	var _bbgg string
	if _dc, _ac := _cd.GetString(_da.V); _ac {
		_bbgg = _dc.Decoded()
	}
	if len(_bbgg) == 0 {
		return nil, nil
	}
	_cgde := []string{_bbgg}
	_cf := false
	if _da.Flags().Has(_bb.FieldFlagMultiline) {
		_cf = true
		_bbgg = _g.Replace(_bbgg, "\u000d\u000a", "\u000a", -1)
		_bbgg = _g.Replace(_bbgg, "\u000d", "\u000a", -1)
		_cgde = _g.Split(_bbgg, "\u000a")
	}
	_efc := make([]string, len(_cgde))
	copy(_efc, _cgde)
	_aeg := _dae.MultilineLineHeight
	_eae := 0.0
	_baa := 0
	if _cbg != nil {
		for _ea >= 0 {
			_bgce := make([]string, len(_cgde))
			copy(_bgce, _cgde)
			_gbd := make([]string, len(_efc))
			copy(_gbd, _efc)
			_eae = 0.0
			_baa = 0
			_dfd := len(_bgce)
			_acg := 0
			for _acg < _dfd {
				var _fceg float64
				_bd := -1
				_bfg := _dgg
				if _dae.MarginLeft != nil {
					_bfg = *_dae.MarginLeft
				}
				for _eaf, _eag := range _bgce[_acg] {
					if _eag == ' ' {
						_bd = _eaf
					}
					_fea, _ecg := _dag.GetRuneMetrics(_eag)
					if !_ecg {
						_d.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u006d\u0065\u0074r\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u002d\u0020\u0073k\u0069p\u0070\u0069\u006e\u0067", _eag)
						continue
					}
					_fceg = _bfg
					_bfg += _fea.Wx
					if _cf && !_gda && _ea*_bfg/1000.0 > _gfb {
						_daa := _eaf
						_ecd := _eaf
						if _bd > 0 {
							_daa = _bd + 1
							_ecd = _bd
						}
						_geb := _bgce[_acg][_daa:]
						_ffe := _gbd[_acg][_daa:]
						if _acg < len(_bgce)-1 {
							_bgce = append(_bgce[:_acg+1], _bgce[_acg:]...)
							_bgce[_acg+1] = _geb
							_gbd = append(_gbd[:_acg+1], _gbd[_acg:]...)
							_gbd[_acg+1] = _ffe
						} else {
							_bgce = append(_bgce, _geb)
							_gbd = append(_gbd, _ffe)
						}
						_bgce[_acg] = _bgce[_acg][0:_ecd]
						_gbd[_acg] = _gbd[_acg][0:_ecd]
						_dfd++
						_bfg = _fceg
						break
					}
				}
				if _bfg > _eae {
					_eae = _bfg
				}
				_bgce[_acg] = string(_cbg.Encode(_bgce[_acg]))
				if len(_bgce[_acg]) > 0 {
					_baa++
				}
				_acg++
			}
			_gba := _ea
			if _baa > 1 {
				_gba *= _aeg
			}
			_gdb := float64(_baa) * _gba
			if _gda || _gdb <= _fgf {
				_cgde = _bgce
				_efc = _gbd
				break
			}
			_ea--
		}
	}
	_fbdd := _dgg
	if _dae.MarginLeft != nil {
		_fbdd = *_dae.MarginLeft
	}
	if _ea == 0 || _gda && _eae > 0 && _fbdd+_eae*_ea/1000.0 > _gfb {
		_ea = 0.95 * 1000.0 * (_gfb - _fbdd) / _eae
	}
	_bbgc := _gb
	{
		if _bef, _fbdb := _cd.GetIntVal(_da.Q); _fbdb {
			switch _bef {
			case 0:
				_bbgc = _gb
			case 1:
				_bbgc = _cc
			case 2:
				_bbgc = _cb
			default:
				_d.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0071\u0075\u0061\u0064\u0064\u0069\u006e\u0067\u003a\u0020%\u0064\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u006c\u0065ft\u0020\u0061\u006c\u0069\u0067\u006e\u006d\u0065\u006e\u0074", _bef)
			}
		}
	}
	_ffbc := _ea
	if _cf && _baa > 1 {
		_ffbc = _aeg * _ea
	}
	var _ggf float64
	if _ee != nil {
		_ggf, _dace = _ee.GetCapHeight()
		if _dace != nil {
			_d.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0043\u0061\u0070\u0048\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _dace)
		}
	}
	if int(_ggf) <= 0 {
		_d.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u0061\u0070\u0048e\u0069\u0067\u0068\u0074\u0020\u006e\u006ft \u0061\u0076\u0061\u0069l\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065tt\u0069\u006eg\u0020\u0074\u006f\u0020\u0031\u0030\u0030\u0030")
		_ggf = 1000
	}
	_deg := _ggf / 1000.0 * _ea
	_ebg := 0.0
	{
		_ece := float64(_baa) * _ffbc
		if _gda && _ebg+_ece > _fgf {
			_ea = 0.95 * (_fgf - _ebg) / float64(_baa)
			_ffbc = _ea
			if _cf && _baa > 1 {
				_ffbc = _aeg * _ea
			}
			_deg = _ggf / 1000.0 * _ea
			_ece = float64(_baa) * _ffbc
		}
		if _fgf > _ece {
			if _cf {
				if _dae.MultilineVAlignMiddle {
					_ddb := (_fgf - (_ece + _deg)) / 2.0
					_bbc := _ddb + _ece + _deg - _ffbc
					_ebg = _bbc
					if _baa > 1 {
						_ebg = _ebg + (_ece / _ea * float64(_baa)) - _ffbc - _deg
					}
					if _ebg < _ece {
						_ebg = (_fgf - _deg) / 2.0
					}
				} else {
					_ebg = _fgf - _ffbc
					if _ebg > _ea {
						_aac := 0.0
						if _cf && _dae.MultilineLineHeight > 1 && _baa > 1 {
							_aac = _dae.MultilineLineHeight - 1
						}
						_ebg -= _ea * (0.5 - _aac)
					}
				}
			} else {
				_ebg = (_fgf - _deg) / 2.0
			}
		}
	}
	_bf.Add_Tf(*_af, _ea)
	_bf.Add_Td(_fbdd, _ebg)
	_degf := _fbdd
	_gdgd := _fbdd
	for _fceb, _fcf := range _cgde {
		_gcda := 0.0
		for _, _gef := range _efc[_fceb] {
			_dge, _fa := _dag.GetRuneMetrics(_gef)
			if !_fa {
				continue
			}
			_gcda += _dge.Wx
		}
		_egg := _gcda / 1000.0 * _ea
		_gdd := _gfb - _egg
		var _agc float64
		switch _bbgc {
		case _gb:
			_agc = _degf
		case _cc:
			_agc = _gdd / 2
		case _cb:
			_agc = _gdd
		}
		_fbdd = _agc - _gdgd
		if _fbdd > 0.0 {
			_bf.Add_Td(_fbdd, 0)
		}
		_gdgd = _agc
		_bf.Add_Tj(*_cd.MakeString(_fcf))
		if _fceb < len(_cgde)-1 {
			_bf.Add_Td(0, -_ea*_aeg)
		}
	}
	_bf.Add_ET()
	_bf.Add_Q()
	_bf.Add_EMC()
	_eeb := _bb.NewXObjectForm()
	_eeb.Resources = _gcbc
	_eeb.BBox = _cd.MakeArrayFromFloats([]float64{0, 0, _dfa, _ec})
	_eeb.SetContentStream(_bf.Bytes(), _bed())
	_bfd := _cd.MakeDict()
	_bfd.Set("\u004e", _eeb.ToPdfObject())
	return _bfd, nil
}
func _abgf(_ega []*SignatureLine, _eee *SignatureFieldOpts) (*_cd.PdfObjectDictionary, error) {
	if _eee == nil {
		_eee = NewSignatureFieldOpts()
	}
	var _ggfb error
	var _fegb *_cd.PdfObjectName
	_eac := _eee.Font
	if _eac != nil {
		_baag, _ := _eac.GetFontDescriptor()
		if _baag != nil {
			if _afg, _gdda := _baag.FontName.(*_cd.PdfObjectName); _gdda {
				_fegb = _afg
			}
		}
		if _fegb == nil {
			_fegb = _cd.MakeName("\u0046\u006f\u006et\u0031")
		}
	} else {
		if _eac, _ggfb = _bb.NewStandard14Font("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a"); _ggfb != nil {
			return nil, _ggfb
		}
		_fegb = _cd.MakeName("\u0048\u0065\u006c\u0076")
	}
	_fggcg := _eee.FontSize
	if _fggcg <= 0 {
		_fggcg = 10
	}
	if _eee.LineHeight <= 0 {
		_eee.LineHeight = 1
	}
	_bdadb := _eee.LineHeight * _fggcg
	_bbf, _aadf := _eac.GetRuneMetrics(' ')
	if !_aadf {
		return nil, _f.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
	}
	_dgc := _bbf.Wx
	var _ged float64
	var _dfda []string
	for _, _eggd := range _ega {
		if _eggd.Text == "" {
			continue
		}
		_gbc := _eggd.Text
		if _eggd.Desc != "" {
			_gbc = _eggd.Desc + "\u003a\u0020" + _gbc
		}
		_dfda = append(_dfda, _gbc)
		var _ege float64
		for _, _cebe := range _gbc {
			_agad, _fgbd := _eac.GetRuneMetrics(_cebe)
			if !_fgbd {
				continue
			}
			_ege += _agad.Wx
		}
		if _ege > _ged {
			_ged = _ege
		}
	}
	_ged = _ged * _fggcg / 1000.0
	_bdgb := float64(len(_dfda)) * _bdadb
	_gcc := _eee.Image != nil
	_ddba := _eee.Rect
	if _ddba == nil {
		_ddba = []float64{0, 0, _ged, _bdgb}
		if _gcc {
			_ddba[2] = _ged * 2
			_ddba[3] = _bdgb * 2
		}
		_eee.Rect = _ddba
	}
	_ddbg := _ddba[2] - _ddba[0]
	_caag := _ddba[3] - _ddba[1]
	_cdb, _bbe := _ddba, _ddba
	var _dgd, _cgaad float64
	if _gcc && len(_dfda) > 0 {
		if _eee.ImagePosition <= SignatureImageRight {
			_dagc := []float64{_ddba[0], _ddba[1], _ddba[0] + (_ddbg / 2), _ddba[3]}
			_dea := []float64{_ddba[0] + (_ddbg / 2), _ddba[1], _ddba[2], _ddba[3]}
			if _eee.ImagePosition == SignatureImageLeft {
				_cdb, _bbe = _dagc, _dea
			} else {
				_cdb, _bbe = _dea, _dagc
			}
		} else {
			_fbe := []float64{_ddba[0], _ddba[1], _ddba[2], _ddba[1] + (_caag / 2)}
			_bbgf := []float64{_ddba[0], _ddba[1] + (_caag / 2), _ddba[2], _ddba[3]}
			if _eee.ImagePosition == SignatureImageTop {
				_cdb, _bbe = _bbgf, _fbe
			} else {
				_cdb, _bbe = _fbe, _bbgf
			}
		}
	}
	_dgd = _bbe[2] - _bbe[0]
	_cgaad = _bbe[3] - _bbe[1]
	var _baea float64
	if _eee.AutoSize {
		if _ged > _dgd || _bdgb > _cgaad {
			_fgcf := _dg.Min(_dgd/_ged, _cgaad/_bdgb)
			_fggcg *= _fgcf
		}
		_bdadb = _eee.LineHeight * _fggcg
		_baea += (_cgaad - float64(len(_dfda))*_bdadb) / 2
	}
	_ggcc := _bg.NewContentCreator()
	_fdda := _bb.NewPdfPageResources()
	_fdda.SetFontByName(*_fegb, _eac.ToPdfObject())
	if _eee.BorderSize <= 0 {
		_eee.BorderSize = 0
		_eee.BorderColor = _bb.NewPdfColorDeviceGray(1)
	}
	_ggcc.Add_q()
	if _eee.FillColor != nil {
		_ggcc.SetNonStrokingColor(_eee.FillColor)
	}
	if _eee.BorderColor != nil {
		_ggcc.SetStrokingColor(_eee.BorderColor)
	}
	_ggcc.Add_w(_eee.BorderSize).Add_re(_ddba[0], _ddba[1], _ddbg, _caag)
	if _eee.FillColor != nil && _eee.BorderColor != nil {
		_ggcc.Add_B()
	} else if _eee.FillColor != nil {
		_ggcc.Add_f()
	} else if _eee.BorderColor != nil {
		_ggcc.Add_S()
	}
	_ggcc.Add_Q()
	if _eee.WatermarkImage != nil {
		_bfdd := []float64{_ddba[0], _ddba[1], _ddba[2], _ddba[3]}
		_bbeg, _gffg, _cec := _eed(_eee.WatermarkImage, "\u0049\u006d\u0061\u0067\u0065\u0057\u0061\u0074\u0065r\u006d\u0061\u0072\u006b", _eee, _bfdd, _ggcc)
		if _cec != nil {
			return nil, _cec
		}
		_fdda.SetXObjectImageByName(*_bbeg, _gffg)
	}
	_ggcc.Add_q()
	_ggcc.Translate(_bbe[0], _bbe[3]-_bdadb-_baea)
	_ggcc.Add_BT()
	_adaa := _eac.Encoder()
	for _, _egac := range _dfda {
		var _dcee []byte
		for _, _efg := range _egac {
			if _ad.IsSpace(_efg) {
				if len(_dcee) > 0 {
					_ggcc.SetNonStrokingColor(_eee.TextColor).Add_Tf(*_fegb, _fggcg).Add_TL(_bdadb).Add_TJ([]_cd.PdfObject{_cd.MakeStringFromBytes(_dcee)}...)
					_dcee = nil
				}
				_ggcc.Add_Tf(*_fegb, _fggcg).Add_TL(_bdadb).Add_TJ([]_cd.PdfObject{_cd.MakeFloat(-_dgc)}...)
			} else {
				_dcee = append(_dcee, _adaa.Encode(string(_efg))...)
			}
		}
		if len(_dcee) > 0 {
			_ggcc.SetNonStrokingColor(_eee.TextColor).Add_Tf(*_fegb, _fggcg).Add_TL(_bdadb).Add_TJ([]_cd.PdfObject{_cd.MakeStringFromBytes(_dcee)}...)
		}
		_ggcc.Add_Td(0, -_bdadb)
	}
	_ggcc.Add_ET()
	_ggcc.Add_Q()
	if _gcc {
		_cagdc, _ebbf, _gefe := _eed(_eee.Image, "\u0049\u006d\u0061\u0067\u0065\u0053\u0069\u0067\u006ea\u0074\u0075\u0072\u0065", _eee, _cdb, _ggcc)
		if _gefe != nil {
			return nil, _gefe
		}
		_fdda.SetXObjectImageByName(*_cagdc, _ebbf)
	}
	_deb := _bb.NewXObjectForm()
	_deb.Resources = _fdda
	_deb.BBox = _cd.MakeArrayFromFloats(_ddba)
	_deb.SetContentStream(_ggcc.Bytes(), _bed())
	_geeb := _cd.MakeDict()
	_geeb.Set("\u004e", _deb.ToPdfObject())
	return _geeb, nil
}

// WrapContentStream ensures that the entire content stream for a `page` is wrapped within q ... Q operands.
// Ensures that following operands that are added are not affected by additional operands that are added.
// Implements interface model.ContentStreamWrapper.
func (_gccb ImageFieldAppearance) WrapContentStream(page *_bb.PdfPage) error {
	_cab, _bgeb := page.GetAllContentStreams()
	if _bgeb != nil {
		return _bgeb
	}
	_abb := _bg.NewContentStreamParser(_cab)
	_cgc, _bgeb := _abb.Parse()
	if _bgeb != nil {
		return _bgeb
	}
	_cgc.WrapIfNeeded()
	_gccbf := []string{_cgc.String()}
	return page.SetContentStreams(_gccbf, _bed())
}

// WrapContentStream ensures that the entire content stream for a `page` is wrapped within q ... Q operands.
// Ensures that following operands that are added are not affected by additional operands that are added.
// Implements interface model.ContentStreamWrapper.
func (_fggd FieldAppearance) WrapContentStream(page *_bb.PdfPage) error {
	_caeb, _gdcf := page.GetAllContentStreams()
	if _gdcf != nil {
		return _gdcf
	}
	_eeg := _bg.NewContentStreamParser(_caeb)
	_ebec, _gdcf := _eeg.Parse()
	if _gdcf != nil {
		return _gdcf
	}
	_ebec.WrapIfNeeded()
	_dbg := []string{_ebec.String()}
	return page.SetContentStreams(_dbg, _bed())
}

// NewTextField generates a new text field with partial name `name` at location
// specified by `rect` on given `page` and with field specific options `opt`.
func NewTextField(page *_bb.PdfPage, name string, rect []float64, opt TextFieldOptions) (*_bb.PdfFieldText, error) {
	if page == nil {
		return nil, _f.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _f.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _f.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_edca := _bb.NewPdfField()
	_acgb := &_bb.PdfFieldText{}
	_edca.SetContext(_acgb)
	_acgb.PdfField = _edca
	_acgb.T = _cd.MakeString(name)
	if opt.MaxLen > 0 {
		_acgb.MaxLen = _cd.MakeInteger(int64(opt.MaxLen))
	}
	if len(opt.Value) > 0 {
		_acgb.V = _cd.MakeString(opt.Value)
	}
	_fgeg := _bb.NewPdfAnnotationWidget()
	_fgeg.Rect = _cd.MakeArrayFromFloats(rect)
	_fgeg.P = page.ToPdfObject()
	_fgeg.F = _cd.MakeInteger(4)
	_fgeg.Parent = _acgb.ToPdfObject()
	_acgb.Annotations = append(_acgb.Annotations, _fgeg)
	return _acgb, nil
}

// AppearanceStyle defines style parameters for appearance stream generation.
type AppearanceStyle struct {

	// How much of Rect height to fill when autosizing text.
	AutoFontSizeFraction float64

	// CheckmarkRune is a rune used for check mark in checkboxes (for ZapfDingbats font).
	CheckmarkRune rune
	BorderSize    float64
	BorderColor   _bb.PdfColor
	FillColor     _bb.PdfColor

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

// NewComboboxField generates a new combobox form field with partial name `name` at location `rect`
// on specified `page` and with field specific options `opt`.
func NewComboboxField(page *_bb.PdfPage, name string, rect []float64, opt ComboboxFieldOptions) (*_bb.PdfFieldChoice, error) {
	if page == nil {
		return nil, _f.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _f.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _f.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_dab := _bb.NewPdfField()
	_dfb := &_bb.PdfFieldChoice{}
	_dab.SetContext(_dfb)
	_dfb.PdfField = _dab
	_dfb.T = _cd.MakeString(name)
	_dfb.Opt = _cd.MakeArray()
	for _, _agab := range opt.Choices {
		_dfb.Opt.Append(_cd.MakeString(_agab))
	}
	_dfb.SetFlag(_bb.FieldFlagCombo)
	_aafe := _bb.NewPdfAnnotationWidget()
	_aafe.Rect = _cd.MakeArrayFromFloats(rect)
	_aafe.P = page.ToPdfObject()
	_aafe.F = _cd.MakeInteger(4)
	_aafe.Parent = _dfb.ToPdfObject()
	_dfb.Annotations = append(_dfb.Annotations, _aafe)
	return _dfb, nil
}

// FormSubmitActionOptions holds options for creating a form submit button.
type FormSubmitActionOptions struct {

	// Rectangle holds the button position, size, and color.
	Rectangle _fb.Rectangle

	// Url specifies the URL where the fieds will be submitted.
	Url string

	// Label specifies the text that would be displayed on the button.
	Label string

	// LabelColor specifies the button label color.
	LabelColor _bb.PdfColor

	// Font specifies a font used for rendering the button label.
	// When omitted it will fallback to use a Helvetica font.
	Font *_bb.PdfFont

	// FontSize specifies the font size used in rendering the button label.
	// The default font size is 12pt.
	FontSize *float64

	// Fields specifies list of fields that could be submitted.
	// This list may contain indirect object to fields or field names.
	Fields *_cd.PdfObjectArray

	// IsExclusionList specifies that the fields contain in `Fields` array would not be submitted.
	IsExclusionList bool

	// IncludeEmptyFields specifies if all fields would be submitted even though it's value is empty.
	IncludeEmptyFields bool

	// SubmitAsPDF specifies that the document shall be submitted as PDF.
	// If set then all the other flags shall be ignored.
	SubmitAsPDF bool
}

func _dbf(_ebcb *_bb.PdfField) string {
	if _ebcb == nil {
		return ""
	}
	_dce, _bdb := _ebcb.GetContext().(*_bb.PdfFieldText)
	if !_bdb {
		return _dbf(_ebcb.Parent)
	}
	if _dce.DA != nil {
		return _dce.DA.Str()
	}
	return _dbf(_dce.Parent)
}
func (_gfc *AppearanceFont) fillName() {
	if _gfc.Font == nil || _gfc.Name != "" {
		return
	}
	_ce := _gfc.Font.FontDescriptor()
	if _ce == nil || _ce.FontName == nil {
		return
	}
	_gfc.Name = _ce.FontName.String()
}

// NewFormSubmitButtonField would create a submit button in specified page according to the parameter in `FormSubmitActionOptions`.
func NewFormSubmitButtonField(page *_bb.PdfPage, opt FormSubmitActionOptions) (*_bb.PdfFieldButton, error) {
	_baeg := int64(_gbfc)
	if opt.IsExclusionList {
		_baeg |= _fag
	}
	if opt.IncludeEmptyFields {
		_baeg |= _gad
	}
	if opt.SubmitAsPDF {
		_baeg |= _gdbd
	}
	_gedgd := _bb.NewPdfActionSubmitForm()
	_gedgd.Flags = _cd.MakeInteger(_baeg)
	_gedgd.F = _bb.NewPdfFilespec()
	if opt.Fields != nil {
		_gedgd.Fields = opt.Fields
	}
	_gedgd.F.F = _cd.MakeString(opt.Url)
	_gedgd.F.FS = _cd.MakeName("\u0055\u0052\u004c")
	_fgda, _dffd := _ecc(page, opt.Rectangle, "\u0062t\u006e\u0053\u0075\u0062\u006d\u0069t", opt.Label, opt.LabelColor, opt.Font, opt.FontSize, _gedgd.ToPdfObject())
	if _dffd != nil {
		return nil, _dffd
	}
	return _fgda, nil
}
func _cefbf(_gceb LineAnnotationDef, _ecea string) ([]byte, *_bb.PdfRectangle, *_bb.PdfRectangle, error) {
	_gaca := _fb.Line{X1: 0, Y1: 0, X2: _gceb.X2 - _gceb.X1, Y2: _gceb.Y2 - _gceb.Y1, LineColor: _gceb.LineColor, Opacity: _gceb.Opacity, LineWidth: _gceb.LineWidth, LineEndingStyle1: _gceb.LineEndingStyle1, LineEndingStyle2: _gceb.LineEndingStyle2}
	_ebed, _dcf, _gccbg := _gaca.Draw(_ecea)
	if _gccbg != nil {
		return nil, nil, nil, _gccbg
	}
	_fgdg := &_bb.PdfRectangle{}
	_fgdg.Llx = _gceb.X1 + _dcf.Llx
	_fgdg.Lly = _gceb.Y1 + _dcf.Lly
	_fgdg.Urx = _gceb.X1 + _dcf.Urx
	_fgdg.Ury = _gceb.Y1 + _dcf.Ury
	return _ebed, _dcf, _fgdg, nil
}

// InkAnnotationDef holds base information for constructing an ink annotation.
type InkAnnotationDef struct {

	// Paths is the array of stroked paths which compose the annotation.
	Paths []_fb.Path

	// Color is the color of the line. Default to black.
	Color *_bb.PdfColorDeviceRGB

	// LineWidth is the width of the line.
	LineWidth float64
}

func (_bdfb *AppearanceStyle) processDA(_cbb *_bb.PdfField, _feg *_bg.ContentStreamOperations, _cefb, _aead *_bb.PdfPageResources, _eea *_bg.ContentCreator) (*AppearanceFont, bool, error) {
	var _ebea *AppearanceFont
	var _dbad bool
	if _bdfb.Fonts != nil {
		if _bdfb.Fonts.Fallback != nil {
			_ebea = _bdfb.Fonts.Fallback
		}
		if _cdab := _bdfb.Fonts.FieldFallbacks; _cdab != nil {
			if _gcga, _cad := _cdab[_cbb.PartialName()]; _cad {
				_ebea = _gcga
			} else if _fgb, _gfce := _cbb.FullName(); _gfce == nil {
				if _cgag, _ceb := _cdab[_fgb]; _ceb {
					_ebea = _cgag
				}
			}
		}
		if _ebea != nil {
			_ebea.fillName()
		}
		_dbad = _bdfb.Fonts.ForceReplace
	}
	var _dfab string
	var _ada float64
	var _fffd bool
	if _feg != nil {
		for _, _afc := range *_feg {
			if _afc.Operand == "\u0054\u0066" && len(_afc.Params) == 2 {
				if _adda, _edbb := _cd.GetNameVal(_afc.Params[0]); _edbb {
					_dfab = _adda
				}
				if _aab, _ffbcc := _cd.GetNumberAsFloat(_afc.Params[1]); _ffbcc == nil {
					_ada = _aab
				}
				_fffd = true
				continue
			}
			_eea.AddOperand(*_afc)
		}
	}
	var _dfdf *AppearanceFont
	var _fggb _cd.PdfObject
	if _dbad && _ebea != nil {
		_dfdf = _ebea
	} else {
		if _cefb != nil && _dfab != "" {
			if _febd, _cefbb := _cefb.GetFontByName(*_cd.MakeName(_dfab)); _cefbb {
				if _fgfg, _gfe := _bb.NewPdfFontFromPdfObject(_febd); _gfe == nil {
					_fggb = _febd
					_dfdf = &AppearanceFont{Name: _dfab, Font: _fgfg, Size: _ada}
				} else {
					_d.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006fa\u0064\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0025\u0076", _gfe)
				}
			}
		}
		if _dfdf == nil && _ebea != nil {
			_dfdf = _ebea
		}
		if _dfdf == nil {
			_agda, _bcec := _bb.NewStandard14Font("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
			if _bcec != nil {
				return nil, false, _bcec
			}
			_dfdf = &AppearanceFont{Name: "\u0048\u0065\u006c\u0076", Font: _agda, Size: _ada}
		}
	}
	if _dfdf.Size <= 0 && _bdfb.Fonts != nil && _bdfb.Fonts.FallbackSize > 0 {
		_dfdf.Size = _bdfb.Fonts.FallbackSize
	}
	_fcad := *_cd.MakeName(_dfdf.Name)
	if _fggb == nil {
		_fggb = _dfdf.Font.ToPdfObject()
	}
	if _cefb != nil && !_cefb.HasFontByName(_fcad) {
		_cefb.SetFontByName(_fcad, _fggb)
	}
	if _aead != nil && !_aead.HasFontByName(_fcad) {
		_aead.SetFontByName(_fcad, _fggb)
	}
	return _dfdf, _fffd, nil
}

// AppearanceFont represents a font used for generating the appearance of a
// field in the filling/flattening process.
type AppearanceFont struct {

	// Name represents the name of the font which will be added to the
	// AcroForm resources (DR).
	Name string

	// Font represents the actual font used for the field appearance.
	Font *_bb.PdfFont

	// Size represents the size of the font used for the field appearance.
	// If the font size is 0, the value of the FallbackSize field of the
	// AppearanceFontStyle is used, if set. Otherwise, the font size is
	// calculated based on the available annotation height and on the
	// AutoFontSizeFraction field of the AppearanceStyle.
	Size float64
}

func _gbaf(_fcc *_bb.PdfAnnotationWidget, _egbb *_bb.PdfFieldButton, _aecc *_bb.PdfPageResources, _bdgd AppearanceStyle) (*_cd.PdfObjectDictionary, error) {
	_cgb, _egge := _cd.GetArray(_fcc.Rect)
	if !_egge {
		return nil, _f.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_feb, _eda := _bb.NewPdfRectangle(*_cgb)
	if _eda != nil {
		return nil, _eda
	}
	_ccf, _dfg := _feb.Width(), _feb.Height()
	_aea, _ddf := _ccf, _dfg
	_d.Log.Debug("\u0043\u0068\u0065\u0063kb\u006f\u0078\u002c\u0020\u0077\u0061\u0020\u0042\u0053\u003a\u0020\u0025\u0076", _fcc.BS)
	_cbgf, _eda := _bb.NewStandard14Font("\u005a\u0061\u0070f\u0044\u0069\u006e\u0067\u0062\u0061\u0074\u0073")
	if _eda != nil {
		return nil, _eda
	}
	_adf, _caca := _cd.GetDict(_fcc.MK)
	if _caca {
		_acd, _ := _cd.GetDict(_fcc.BS)
		_edb := _bdgd.applyAppearanceCharacteristics(_adf, _acd, _cbgf)
		if _edb != nil {
			return nil, _edb
		}
	}
	_ebc := _bb.NewXObjectForm()
	{
		_degg := _bg.NewContentCreator()
		if _bdgd.BorderSize > 0 {
			_ddc(_degg, _bdgd, _ccf, _dfg)
		}
		if _bdgd.DrawAlignmentReticle {
			_ecfb := _bdgd
			_ecfb.BorderSize = 0.2
			_deggd(_degg, _ecfb, _ccf, _dfg)
		}
		_ccf, _dfg = _bdgd.applyRotation(_adf, _ccf, _dfg, _degg)
		_fgg := _bdgd.AutoFontSizeFraction * _dfg
		_ggd, _fdb := _cbgf.GetRuneMetrics(_bdgd.CheckmarkRune)
		if !_fdb {
			return nil, _f.New("\u0067l\u0079p\u0068\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
		}
		_bff := _cbgf.Encoder()
		_aega := _bff.Encode(string(_bdgd.CheckmarkRune))
		_caac := _ggd.Wx * _fgg / 1000.0
		_gbde := 705.0
		_ecbg := _gbde / 1000.0 * _fgg
		_add := _dgg
		if _bdgd.MarginLeft != nil {
			_add = *_bdgd.MarginLeft
		}
		_gefa := 1.0
		if _caac < _ccf {
			_add = (_ccf - _caac) / 2.0
		}
		if _ecbg < _dfg {
			_gefa = (_dfg - _ecbg) / 2.0
		}
		_degg.Add_q().Add_g(0).Add_BT().Add_Tf("\u005a\u0061\u0044\u0062", _fgg).Add_Td(_add, _gefa).Add_Tj(*_cd.MakeStringFromBytes(_aega)).Add_ET().Add_Q()
		_ebc.Resources = _bb.NewPdfPageResources()
		_ebc.Resources.SetFontByName("\u005a\u0061\u0044\u0062", _cbgf.ToPdfObject())
		_ebc.BBox = _cd.MakeArrayFromFloats([]float64{0, 0, _aea, _ddf})
		_ebc.SetContentStream(_degg.Bytes(), _bed())
	}
	_bce := _bb.NewXObjectForm()
	{
		_ebb := _bg.NewContentCreator()
		if _bdgd.BorderSize > 0 {
			_ddc(_ebb, _bdgd, _ccf, _dfg)
		}
		_bce.BBox = _cd.MakeArrayFromFloats([]float64{0, 0, _aea, _ddf})
		_bce.SetContentStream(_ebb.Bytes(), _bed())
	}
	_cde := _cd.PdfObjectName("\u0059\u0065\u0073")
	_bbb, _caca := _cd.GetDict(_fcc.PdfAnnotation.AP)
	if _caca && _bbb != nil {
		_fca := _cd.TraceToDirectObject(_bbb.Get("\u004e"))
		switch _cbag := _fca.(type) {
		case *_cd.PdfObjectDictionary:
			_caab := _cbag.Keys()
			for _, _dfae := range _caab {
				if _dfae != "\u004f\u0066\u0066" {
					_cde = _dfae
				}
			}
		}
	}
	_ddfa := _cd.MakeDict()
	_ddfa.Set("\u004f\u0066\u0066", _bce.ToPdfObject())
	_ddfa.Set(_cde, _ebc.ToPdfObject())
	_aegf := _cd.MakeDict()
	_aegf.Set("\u004e", _ddfa)
	return _aegf, nil
}

// NewSignatureLine returns a new signature line displayed as a part of the
// signature field appearance.
func NewSignatureLine(desc, text string) *SignatureLine {
	return &SignatureLine{Desc: desc, Text: text}
}
func _cede(_abbc *InkAnnotationDef) (*_cd.PdfObjectDictionary, *_bb.PdfRectangle, error) {
	_fee := _bb.NewXObjectForm()
	_gggc, _dcdg, _gdbdf := _cdbd(_abbc)
	if _gdbdf != nil {
		return nil, nil, _gdbdf
	}
	_gdbdf = _fee.SetContentStream(_gggc, nil)
	if _gdbdf != nil {
		return nil, nil, _gdbdf
	}
	_fee.BBox = _dcdg.ToPdfObject()
	_fee.Resources = _bb.NewPdfPageResources()
	_fee.Resources.ProcSet = _cd.MakeArray(_cd.MakeName("\u0050\u0044\u0046"))
	_fgbe := _cd.MakeDict()
	_fgbe.Set("\u004e", _fee.ToPdfObject())
	return _fgbe, _dcdg, nil
}
func _gde(_abf RectangleAnnotationDef) (*_cd.PdfObjectDictionary, *_bb.PdfRectangle, error) {
	_ecfga := _bb.NewXObjectForm()
	_ecfga.Resources = _bb.NewPdfPageResources()
	_fdde := ""
	if _abf.Opacity < 1.0 {
		_daag := _cd.MakeDict()
		_daag.Set("\u0063\u0061", _cd.MakeFloat(_abf.Opacity))
		_daag.Set("\u0043\u0041", _cd.MakeFloat(_abf.Opacity))
		_eagd := _ecfga.Resources.AddExtGState("\u0067\u0073\u0031", _daag)
		if _eagd != nil {
			_d.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020\u0065\u0078\u0074g\u0073\u0074\u0061t\u0065 \u0067\u0073\u0031")
			return nil, nil, _eagd
		}
		_fdde = "\u0067\u0073\u0031"
	}
	_beab, _efgg, _ceeb, _edcd := _dfbg(_abf, _fdde)
	if _edcd != nil {
		return nil, nil, _edcd
	}
	_edcd = _ecfga.SetContentStream(_beab, nil)
	if _edcd != nil {
		return nil, nil, _edcd
	}
	_ecfga.BBox = _efgg.ToPdfObject()
	_dbgb := _cd.MakeDict()
	_dbgb.Set("\u004e", _ecfga.ToPdfObject())
	return _dbgb, _ceeb, nil
}
func _eaff(_gdaf *_bb.PdfField, _ace, _fbde float64, _gdba string, _ccc AppearanceStyle, _bdad *_bg.ContentStreamOperations, _fgc *_bb.PdfPageResources, _dgf *_cd.PdfObjectDictionary) (*_bb.XObjectForm, error) {
	_edg := _bb.NewPdfPageResources()
	_fedg, _aad := _ace, _fbde
	_cdcg := _bg.NewContentCreator()
	if _ccc.BorderSize > 0 {
		_ddc(_cdcg, _ccc, _ace, _fbde)
	}
	if _ccc.DrawAlignmentReticle {
		_bdc := _ccc
		_bdc.BorderSize = 0.2
		_deggd(_cdcg, _bdc, _ace, _fbde)
	}
	_cdcg.Add_BMC("\u0054\u0078")
	_cdcg.Add_q()
	_cdcg.Add_BT()
	_ace, _fbde = _ccc.applyRotation(_dgf, _ace, _fbde, _cdcg)
	_cacc, _gcdg, _fgcb := _ccc.processDA(_gdaf, _bdad, _fgc, _edg, _cdcg)
	if _fgcb != nil {
		return nil, _fgcb
	}
	_fffa := _cacc.Font
	_bcc := _cacc.Size
	_edd := _cd.MakeName(_cacc.Name)
	_bcee := _bcc == 0
	if _bcee && _gcdg {
		_bcc = _fbde * _ccc.AutoFontSizeFraction
	}
	_ggcb := _fffa.Encoder()
	if _ggcb == nil {
		_d.Log.Debug("\u0057\u0041RN\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0069\u0073\u0020\u006e\u0069l\u002e\u0020\u0041\u0073s\u0075\u006d\u0069\u006eg \u0069\u0064e\u006et\u0069\u0074\u0079\u0020\u0065\u006ec\u006f\u0064\u0065r\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
		_ggcb = _a.NewIdentityTextEncoder("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048")
	}
	if len(_gdba) == 0 {
		return nil, nil
	}
	_cefe := _dgg
	if _ccc.MarginLeft != nil {
		_cefe = *_ccc.MarginLeft
	}
	_dbca := 0.0
	if _ggcb != nil {
		for _, _bdd := range _gdba {
			_gcg, _dff := _fffa.GetRuneMetrics(_bdd)
			if !_dff {
				_d.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u006d\u0065\u0074r\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u002d\u0020\u0073k\u0069p\u0070\u0069\u006e\u0067", _bdd)
				continue
			}
			_dbca += _gcg.Wx
		}
		_gdba = string(_ggcb.Encode(_gdba))
	}
	if _bcc == 0 || _bcee && _dbca > 0 && _cefe+_dbca*_bcc/1000.0 > _ace {
		_bcc = 0.95 * 1000.0 * (_ace - _cefe) / _dbca
	}
	_bdde := 1.0 * _bcc
	_eggg := 2.0
	{
		_gcbd := _bdde
		if _bcee && _eggg+_gcbd > _fbde {
			_bcc = 0.95 * (_fbde - _eggg)
			_bdde = 1.0 * _bcc
			_gcbd = _bdde
		}
		if _fbde > _gcbd {
			_eggg = (_fbde - _gcbd) / 2.0
			_eggg += 1.50
		}
	}
	_cdcg.Add_Tf(*_edd, _bcc)
	_cdcg.Add_Td(_cefe, _eggg)
	_cdcg.Add_Tj(*_cd.MakeString(_gdba))
	_cdcg.Add_ET()
	_cdcg.Add_Q()
	_cdcg.Add_EMC()
	_adfa := _bb.NewXObjectForm()
	_adfa.Resources = _edg
	_adfa.BBox = _cd.MakeArrayFromFloats([]float64{0, 0, _fedg, _aad})
	_adfa.SetContentStream(_cdcg.Bytes(), _bed())
	return _adfa, nil
}

// SetStyle applies appearance `style` to `fa`.
func (_abge *ImageFieldAppearance) SetStyle(style AppearanceStyle) { _abge._aade = &style }

// ComboboxFieldOptions defines optional parameters for a combobox form field.
type ComboboxFieldOptions struct {

	// Choices is the list of string values that can be selected.
	Choices []string
}

func _fggc(_febb *_bb.PdfAcroForm, _gee *_bb.PdfAnnotationWidget, _daab *_bb.PdfFieldChoice, _cfe AppearanceStyle) (*_cd.PdfObjectDictionary, error) {
	_acc, _ggg := _cd.GetArray(_gee.Rect)
	if !_ggg {
		return nil, _f.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_cgg, _fcaf := _bb.NewPdfRectangle(*_acc)
	if _fcaf != nil {
		return nil, _fcaf
	}
	_ebe, _fcd := _cgg.Width(), _cgg.Height()
	_d.Log.Debug("\u0043\u0068\u006f\u0069\u0063\u0065\u002c\u0020\u0077\u0061\u0020\u0042S\u003a\u0020\u0025\u0076", _gee.BS)
	_cdcd, _fcaf := _bg.NewContentStreamParser(_dbf(_daab.PdfField)).Parse()
	if _fcaf != nil {
		return nil, _fcaf
	}
	_cfee, _edaa := _cd.GetDict(_gee.MK)
	if _edaa {
		_faa, _ := _cd.GetDict(_gee.BS)
		_baaa := _cfe.applyAppearanceCharacteristics(_cfee, _faa, nil)
		if _baaa != nil {
			return nil, _baaa
		}
	}
	_cagd := _cd.MakeDict()
	for _, _bbge := range _daab.Opt.Elements() {
		if _cage, _fafd := _cd.GetArray(_bbge); _fafd && _cage.Len() == 2 {
			_bbge = _cage.Get(1)
		}
		var _dcg string
		if _gbg, _daca := _cd.GetString(_bbge); _daca {
			_dcg = _gbg.Decoded()
		} else if _fgeb, _efcc := _cd.GetName(_bbge); _efcc {
			_dcg = _fgeb.String()
		} else {
			_d.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u004f\u0070\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006de\u002f\u0073\u0074\u0072\u0069\u006e\u0067 \u002d\u0020\u0025\u0054", _bbge)
			return nil, _f.New("\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u002f\u0073t\u0072\u0069\u006e\u0067")
		}
		if len(_dcg) > 0 {
			_bbae, _dfe := _eaff(_daab.PdfField, _ebe, _fcd, _dcg, _cfe, _cdcd, _febb.DR, _cfee)
			if _dfe != nil {
				return nil, _dfe
			}
			_cagd.Set(*_cd.MakeName(_dcg), _bbae.ToPdfObject())
		}
	}
	_gcaf := _cd.MakeDict()
	_gcaf.Set("\u004e", _cagd)
	return _gcaf, nil
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

func _gacdc(_ced, _cdg float64, _cdcf *_bb.Image, _eaeae AppearanceStyle) (*_bb.XObjectForm, error) {
	_aeab, _fgcbd := _bb.NewXObjectImageFromImage(_cdcf, nil, _cd.NewFlateEncoder())
	if _fgcbd != nil {
		return nil, _fgcbd
	}
	_aeab.Decode = _cd.MakeArrayFromFloats([]float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0})
	_bfa := _bb.NewPdfPageResources()
	_bfa.ProcSet = _cd.MakeArray(_cd.MakeName("\u0050\u0044\u0046"), _cd.MakeName("\u0049\u006d\u0061\u0067\u0065\u0043"))
	_bfa.SetXObjectImageByName(_cd.PdfObjectName("\u0049\u006d\u0030"), _aeab)
	_gefc := _bg.NewContentCreator()
	_gefc.Add_q()
	_gefc.Add_cm(float64(_cdcf.Width), 0, 0, float64(_cdcf.Height), 0, 0)
	_gefc.Add_Do("\u0049\u006d\u0030")
	_gefc.Add_Q()
	_gcbg := _bb.NewXObjectForm()
	_gcbg.FormType = _cd.MakeInteger(1)
	_gcbg.BBox = _cd.MakeArrayFromFloats([]float64{0, 0, float64(_cdcf.Width), float64(_cdcf.Height)})
	_gcbg.Resources = _bfa
	_gcbg.SetContentStream(_gefc.Bytes(), _bed())
	return _gcbg, nil
}
func _bed() _cd.StreamEncoder { return _cd.NewFlateEncoder() }
func _ddc(_dcgb *_bg.ContentCreator, _cffd AppearanceStyle, _dad, _cfg float64) {
	_dcgb.Add_q().Add_re(0, 0, _dad, _cfg).Add_w(_cffd.BorderSize).SetStrokingColor(_cffd.BorderColor).SetNonStrokingColor(_cffd.FillColor).Add_B().Add_Q()
}

const (
	_gb  quadding = 0
	_cc  quadding = 1
	_cb  quadding = 2
	_dgg float64  = 2.0
)

func _ffb(_ffff CircleAnnotationDef, _gdc string) ([]byte, *_bb.PdfRectangle, *_bb.PdfRectangle, error) {
	_bc := _fb.Circle{X: _ffff.X, Y: _ffff.Y, Width: _ffff.Width, Height: _ffff.Height, FillEnabled: _ffff.FillEnabled, FillColor: _ffff.FillColor, BorderEnabled: _ffff.BorderEnabled, BorderWidth: _ffff.BorderWidth, BorderColor: _ffff.BorderColor, Opacity: _ffff.Opacity}
	_aa, _be, _aggd := _bc.Draw(_gdc)
	if _aggd != nil {
		return nil, nil, nil, _aggd
	}
	_dd := &_bb.PdfRectangle{}
	_dd.Llx = _ffff.X + _be.Llx
	_dd.Lly = _ffff.Y + _be.Lly
	_dd.Urx = _ffff.X + _be.Urx
	_dd.Ury = _ffff.Y + _be.Ury
	return _aa, _be, _dd, nil
}

// GenerateAppearanceDict generates an appearance dictionary for widget annotation `wa` for the `field` in `form`.
// Implements interface model.FieldAppearanceGenerator.
func (_edf ImageFieldAppearance) GenerateAppearanceDict(form *_bb.PdfAcroForm, field *_bb.PdfField, wa *_bb.PdfAnnotationWidget) (*_cd.PdfObjectDictionary, error) {
	_, _cee := field.GetContext().(*_bb.PdfFieldButton)
	if !_cee {
		_d.Log.Trace("C\u006f\u0075\u006c\u0064\u0020\u006fn\u006c\u0079\u0020\u0068\u0061\u006ed\u006c\u0065\u0020\u0062\u0075\u0074\u0074o\u006e\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067")
		return nil, nil
	}
	_fcfc, _befe := _cd.GetDict(wa.AP)
	if _befe && _edf.OnlyIfMissing {
		_d.Log.Trace("\u0041\u006c\u0072\u0065a\u0064\u0079\u0020\u0070\u006f\u0070\u0075\u006c\u0061\u0074e\u0064 \u002d\u0020\u0069\u0067\u006e\u006f\u0072i\u006e\u0067")
		return _fcfc, nil
	}
	if form.DR == nil {
		form.DR = _bb.NewPdfPageResources()
	}
	switch _cgf := field.GetContext().(type) {
	case *_bb.PdfFieldButton:
		if _cgf.IsPush() {
			_cce, _cfed := _fgge(_cgf, wa, _edf.Style())
			if _cfed != nil {
				return nil, _cfed
			}
			return _cce, nil
		}
	}
	return nil, nil
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
	FillColor     *_bb.PdfColorDeviceRGB
	BorderEnabled bool
	BorderWidth   float64
	BorderColor   *_bb.PdfColorDeviceRGB
	Opacity       float64
}

// Style returns the appearance style of `fa`. If not specified, returns default style.
func (_bbd FieldAppearance) Style() AppearanceStyle {
	if _bbd._ca != nil {
		return *_bbd._ca
	}
	_dga := _dgg
	return AppearanceStyle{AutoFontSizeFraction: 0.65, CheckmarkRune: '✔', BorderSize: 0.0, BorderColor: _bb.NewPdfColorDeviceGray(0), FillColor: _bb.NewPdfColorDeviceGray(1), MultilineLineHeight: 1.2, MultilineVAlignMiddle: false, DrawAlignmentReticle: false, AllowMK: true, MarginLeft: &_dga}
}
func (_ead *AppearanceStyle) applyAppearanceCharacteristics(_egbd *_cd.PdfObjectDictionary, _fdd *_cd.PdfObjectDictionary, _egdbg *_bb.PdfFont) error {
	if !_ead.AllowMK {
		return nil
	}
	if CA, _bbgcb := _cd.GetString(_egbd.Get("\u0043\u0041")); _bbgcb && _egdbg != nil {
		_aaca := CA.Bytes()
		if len(_aaca) != 0 {
			_bbbe := []rune(_egdbg.Encoder().Decode(_aaca))
			if len(_bbbe) == 1 {
				_ead.CheckmarkRune = _bbbe[0]
			}
		}
	}
	if BC, _egdec := _cd.GetArray(_egbd.Get("\u0042\u0043")); _egdec {
		_fdbf, _edc := BC.ToFloat64Array()
		if _edc != nil {
			return _edc
		}
		switch len(_fdbf) {
		case 1:
			_ead.BorderColor = _bb.NewPdfColorDeviceGray(_fdbf[0])
		case 3:
			_ead.BorderColor = _bb.NewPdfColorDeviceRGB(_fdbf[0], _fdbf[1], _fdbf[2])
		case 4:
			_ead.BorderColor = _bb.NewPdfColorDeviceCMYK(_fdbf[0], _fdbf[1], _fdbf[2], _fdbf[3])
		default:
			_d.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0042\u0043\u0020\u002d\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0028\u0025\u0064)", len(_fdbf))
		}
		if _fdd != nil {
			if _eefe, _edba := _cd.GetNumberAsFloat(_fdd.Get("\u0057")); _edba == nil {
				_ead.BorderSize = _eefe
			}
		}
	}
	if BG, _gcdgf := _cd.GetArray(_egbd.Get("\u0042\u0047")); _gcdgf {
		_cefg, _ade := BG.ToFloat64Array()
		if _ade != nil {
			return _ade
		}
		switch len(_cefg) {
		case 1:
			_ead.FillColor = _bb.NewPdfColorDeviceGray(_cefg[0])
		case 3:
			_ead.FillColor = _bb.NewPdfColorDeviceRGB(_cefg[0], _cefg[1], _cefg[2])
		case 4:
			_ead.FillColor = _bb.NewPdfColorDeviceCMYK(_cefg[0], _cefg[1], _cefg[2], _cefg[3])
		default:
			_d.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0042\u0047\u0020\u002d\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0028\u0025\u0064)", len(_cefg))
		}
	}
	return nil
}

// Style returns the appearance style of `fa`. If not specified, returns default style.
func (_abc ImageFieldAppearance) Style() AppearanceStyle {
	if _abc._aade != nil {
		return *_abc._aade
	}
	return AppearanceStyle{BorderSize: 0.0, BorderColor: _bb.NewPdfColorDeviceGray(0), FillColor: _bb.NewPdfColorDeviceGray(1), DrawAlignmentReticle: false}
}

// SetStyle applies appearance `style` to `fa`.
func (_cbd *FieldAppearance) SetStyle(style AppearanceStyle) { _cbd._ca = &style }
func _cdbd(_fga *InkAnnotationDef) ([]byte, *_bb.PdfRectangle, error) {
	_egdeb := [][]_fb.CubicBezierCurve{}
	for _, _aaea := range _fga.Paths {
		if _aaea.Length() == 0 {
			continue
		}
		_febc := _aaea.Points
		_gcce, _fcga, _fbg := _bgca(_febc)
		if _fbg != nil {
			return nil, nil, _fbg
		}
		if len(_gcce) != len(_fcga) {
			return nil, nil, _f.New("\u0049\u006e\u0065\u0071\u0075\u0061\u006c\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0063\u0061l\u0063\u0075\u006c\u0061\u0074\u0065\u0064\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0061\u006e\u0064\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0063\u006f\u006e\u0074\u0072o\u006c\u0020\u0070\u006f\u0069n\u0074")
		}
		_afff := []_fb.CubicBezierCurve{}
		for _eaccf := 0; _eaccf < len(_gcce); _eaccf++ {
			_afff = append(_afff, _fb.CubicBezierCurve{P0: _febc[_eaccf], P1: _gcce[_eaccf], P2: _fcga[_eaccf], P3: _febc[_eaccf+1]})
		}
		if len(_afff) > 0 {
			_egdeb = append(_egdeb, _afff)
		}
	}
	_bace, _cffc, _cbe := _dccg(_egdeb, _fga.Color, _fga.LineWidth)
	if _cbe != nil {
		return nil, nil, _cbe
	}
	return _bace, _cffc, nil
}

// CreateInkAnnotation creates an ink annotation object that can be added to the annotation list of a PDF page.
func CreateInkAnnotation(inkDef InkAnnotationDef) (*_bb.PdfAnnotation, error) {
	_cdgg := _bb.NewPdfAnnotationInk()
	_fffb := _cd.MakeArray()
	for _, _gfee := range inkDef.Paths {
		if _gfee.Length() == 0 {
			continue
		}
		_dceef := []float64{}
		for _, _aaab := range _gfee.Points {
			_dceef = append(_dceef, _aaab.X, _aaab.Y)
		}
		_fffb.Append(_cd.MakeArrayFromFloats(_dceef))
	}
	_cdgg.InkList = _fffb
	if inkDef.Color == nil {
		inkDef.Color = _bb.NewPdfColorDeviceRGB(0.0, 0.0, 0.0)
	}
	_cdgg.C = _cd.MakeArrayFromFloats([]float64{inkDef.Color.R(), inkDef.Color.G(), inkDef.Color.B()})
	_eeed, _ggfc, _cdcda := _cede(&inkDef)
	if _cdcda != nil {
		return nil, _cdcda
	}
	_cdgg.AP = _eeed
	_cdgg.Rect = _cd.MakeArrayFromFloats([]float64{_ggfc.Llx, _ggfc.Lly, _ggfc.Urx, _ggfc.Ury})
	return _cdgg.PdfAnnotation, nil
}
func _cbab(_geaa []float64) []float64 {
	var (
		_bdbea = len(_geaa)
		_gcbgb = make([]float64, _bdbea)
		_cbbb  = make([]float64, _bdbea)
	)
	_deeaf := 2.0
	_gcbgb[0] = _geaa[0] / _deeaf
	for _ddd := 1; _ddd < _bdbea; _ddd++ {
		_cbbb[_ddd] = 1 / _deeaf
		if _ddd < _bdbea-1 {
			_deeaf = 4.0
		} else {
			_deeaf = 3.5
		}
		_deeaf -= _cbbb[_ddd]
		_gcbgb[_ddd] = (_geaa[_ddd] - _gcbgb[_ddd-1]) / _deeaf
	}
	for _egf := 1; _egf < _bdbea; _egf++ {
		_gcbgb[_bdbea-_egf-1] -= _cbbb[_bdbea-_egf] * _gcbgb[_bdbea-_egf]
	}
	return _gcbgb
}
func _dccg(_dgeb [][]_fb.CubicBezierCurve, _eaeab *_bb.PdfColorDeviceRGB, _baga float64) ([]byte, *_bb.PdfRectangle, error) {
	_deea := _bg.NewContentCreator()
	_deea.Add_q().SetStrokingColor(_eaeab).Add_w(_baga)
	_deec := _fb.NewCubicBezierPath()
	for _, _afdc := range _dgeb {
		_deec.Curves = append(_deec.Curves, _afdc...)
		for _beda, _bggb := range _afdc {
			if _beda == 0 {
				_deea.Add_m(_bggb.P0.X, _bggb.P0.Y)
			} else {
				_deea.Add_l(_bggb.P0.X, _bggb.P0.Y)
			}
			_deea.Add_c(_bggb.P1.X, _bggb.P1.Y, _bggb.P2.X, _bggb.P2.Y, _bggb.P3.X, _bggb.P3.Y)
		}
	}
	_deea.Add_S().Add_Q()
	return _deea.Bytes(), _deec.GetBoundingBox().ToPdfRectangle(), nil
}

// SignatureLine represents a line of information in the signature field appearance.
type SignatureLine struct {
	Desc string
	Text string
}

// NewSignatureFieldOpts returns a new initialized instance of options
// used to generate a signature appearance.
func NewSignatureFieldOpts() *SignatureFieldOpts {
	return &SignatureFieldOpts{Font: _bb.DefaultFont(), FontSize: 10, LineHeight: 1, AutoSize: true, TextColor: _bb.NewPdfColorDeviceGray(0), BorderColor: _bb.NewPdfColorDeviceGray(0), FillColor: _bb.NewPdfColorDeviceGray(1), Encoder: _cd.NewFlateEncoder(), ImagePosition: SignatureImageLeft}
}
func _eed(_eafb _bga.Image, _fgdf string, _fdc *SignatureFieldOpts, _eaa []float64, _dec *_bg.ContentCreator) (*_cd.PdfObjectName, *_bb.XObjectImage, error) {
	_dgde, _cgab := _bb.DefaultImageHandler{}.NewImageFromGoImage(_eafb)
	if _cgab != nil {
		return nil, nil, _cgab
	}
	_afe, _cgab := _bb.NewXObjectImageFromImage(_dgde, nil, _fdc.Encoder)
	if _cgab != nil {
		return nil, nil, _cgab
	}
	_ccfd, _ccb := float64(*_afe.Width), float64(*_afe.Height)
	_bec := _eaa[2] - _eaa[0]
	_ebda := _eaa[3] - _eaa[1]
	if _fdc.AutoSize {
		_gfa := _dg.Min(_bec/_ccfd, _ebda/_ccb)
		_ccfd *= _gfa
		_ccb *= _gfa
		_eaa[0] = _eaa[0] + (_bec / 2) - (_ccfd / 2)
		_eaa[1] = _eaa[1] + (_ebda / 2) - (_ccb / 2)
	}
	var _gcaa *_cd.PdfObjectName
	if _dbe, _cbc := _cd.GetName(_afe.Name); _cbc {
		_gcaa = _dbe
	} else {
		_gcaa = _cd.MakeName(_fgdf)
	}
	if _dec != nil {
		_dec.Add_q().Translate(_eaa[0], _eaa[1]).Scale(_ccfd, _ccb).Add_Do(*_gcaa).Add_Q()
	} else {
		return nil, nil, _f.New("\u0043\u006f\u006e\u0074en\u0074\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u0075l\u006c")
	}
	return _gcaa, _afe, nil
}
func _bgca(_fbf []_fb.Point) (_adfe []_fb.Point, _babf []_fb.Point, _faad error) {
	_bdbg := len(_fbf) - 1
	if len(_fbf) < 1 {
		return nil, nil, _f.New("\u0041\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0074\u0077\u006f\u0020\u0070\u006f\u0069\u006e\u0074s \u0072e\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0074\u006f\u0020\u0063\u0061l\u0063\u0075\u006c\u0061\u0074\u0065\u0020\u0063\u0075\u0072\u0076\u0065\u0020\u0063\u006f\u006e\u0074r\u006f\u006c\u0020\u0070\u006f\u0069\u006e\u0074\u0073")
	}
	if _bdbg == 1 {
		_deae := _fb.Point{X: (2*_fbf[0].X + _fbf[1].X) / 3, Y: (2*_fbf[0].Y + _fbf[1].Y) / 3}
		_adfe = append(_adfe, _deae)
		_babf = append(_babf, _fb.Point{X: 2*_deae.X - _fbf[0].X, Y: 2*_deae.Y - _fbf[0].Y})
		return _adfe, _babf, nil
	}
	_agdc := make([]float64, _bdbg)
	for _aee := 1; _aee < _bdbg-1; _aee++ {
		_agdc[_aee] = 4*_fbf[_aee].X + 2*_fbf[_aee+1].X
	}
	_agdc[0] = _fbf[0].X + 2*_fbf[1].X
	_agdc[_bdbg-1] = (8*_fbf[_bdbg-1].X + _fbf[_bdbg].X) / 2.0
	_edga := _cbab(_agdc)
	for _cccf := 1; _cccf < _bdbg-1; _cccf++ {
		_agdc[_cccf] = 4*_fbf[_cccf].Y + 2*_fbf[_cccf+1].Y
	}
	_agdc[0] = _fbf[0].Y + 2*_fbf[1].Y
	_agdc[_bdbg-1] = (8*_fbf[_bdbg-1].Y + _fbf[_bdbg].Y) / 2.0
	_fafc := _cbab(_agdc)
	_adfe = make([]_fb.Point, _bdbg)
	_babf = make([]_fb.Point, _bdbg)
	for _fcda := 0; _fcda < _bdbg; _fcda++ {
		_adfe[_fcda] = _fb.Point{X: _edga[_fcda], Y: _fafc[_fcda]}
		if _fcda < _bdbg-1 {
			_babf[_fcda] = _fb.Point{X: 2*_fbf[_fcda+1].X - _edga[_fcda+1], Y: 2*_fbf[_fcda+1].Y - _fafc[_fcda+1]}
		} else {
			_babf[_fcda] = _fb.Point{X: (_fbf[_bdbg].X + _edga[_bdbg-1]) / 2, Y: (_fbf[_bdbg].Y + _fafc[_bdbg-1]) / 2}
		}
	}
	return _adfe, _babf, nil
}

// LineAnnotationDef defines a line between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none
// (regular line), or arrows at either end.  The line also has a specified width, color and opacity.
type LineAnnotationDef struct {
	X1               float64
	Y1               float64
	X2               float64
	Y2               float64
	LineColor        *_bb.PdfColorDeviceRGB
	Opacity          float64
	LineWidth        float64
	LineEndingStyle1 _fb.LineEndingStyle
	LineEndingStyle2 _fb.LineEndingStyle
}

func (_fgcd *AppearanceStyle) applyRotation(_geec *_cd.PdfObjectDictionary, _cbf, _gac float64, _gcab *_bg.ContentCreator) (float64, float64) {
	if !_fgcd.AllowMK {
		return _cbf, _gac
	}
	if _geec == nil {
		return _cbf, _gac
	}
	_dda, _ := _cd.GetNumberAsFloat(_geec.Get("\u0052"))
	if _dda == 0 {
		return _cbf, _gac
	}
	_aae := -_dda
	_fgce := _fb.Path{Points: []_fb.Point{_fb.NewPoint(0, 0).Rotate(_aae), _fb.NewPoint(_cbf, 0).Rotate(_aae), _fb.NewPoint(0, _gac).Rotate(_aae), _fb.NewPoint(_cbf, _gac).Rotate(_aae)}}.GetBoundingBox()
	_gcab.RotateDeg(_dda)
	_gcab.Translate(_fgce.X, _fgce.Y)
	return _fgce.Width, _fgce.Height
}

// FormResetActionOptions holds options for creating a form reset button.
type FormResetActionOptions struct {

	// Rectangle holds the button position, size, and color.
	Rectangle _fb.Rectangle

	// Label specifies the text that would be displayed on the button.
	Label string

	// LabelColor specifies the button label color.
	LabelColor _bb.PdfColor

	// Font specifies a font used for rendering the button label.
	// When omitted it will fallback to use a Helvetica font.
	Font *_bb.PdfFont

	// FontSize specifies the font size used in rendering the button label.
	// The default font size is 12pt.
	FontSize *float64

	// Fields specifies list of fields that could be resetted.
	// This list may contain indirect object to fields or field names.
	Fields *_cd.PdfObjectArray

	// IsExclusionList specifies that the fields in the `Fields` array would be excluded form reset process.
	IsExclusionList bool
}

// NewImageField generates a new image field with partial name `name` at location `rect`
// on specified `page` and with field specific options `opt`.
func NewImageField(page *_bb.PdfPage, name string, rect []float64, opt ImageFieldOptions) (*_bb.PdfFieldButton, error) {
	if page == nil {
		return nil, _f.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _f.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _f.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_efgb := _bb.NewPdfField()
	_fbc := &_bb.PdfFieldButton{}
	_fbc.PdfField = _efgb
	_efgb.SetContext(_fbc)
	_fbc.SetType(_bb.ButtonTypePush)
	_fbc.T = _cd.MakeString(name)
	_aeb := _bb.NewPdfAnnotationWidget()
	_aeb.Rect = _cd.MakeArrayFromFloats(rect)
	_aeb.P = page.ToPdfObject()
	_aeb.F = _cd.MakeInteger(4)
	_aeb.Parent = _fbc.ToPdfObject()
	_fdf := rect[2] - rect[0]
	_gfg := rect[3] - rect[1]
	_efce := opt._degc
	_gggf := _bg.NewContentCreator()
	if _efce.BorderSize > 0 {
		_ddc(_gggf, _efce, _fdf, _gfg)
	}
	if _efce.DrawAlignmentReticle {
		_eefa := _efce
		_eefa.BorderSize = 0.2
		_deggd(_gggf, _eefa, _fdf, _gfg)
	}
	_efe, _bbbf := _gacdc(_fdf, _gfg, opt.Image, _efce)
	if _bbbf != nil {
		return nil, _bbbf
	}
	_gfgd, _bbff := _cd.GetDict(_aeb.MK)
	if _bbff {
		_gfgd.Set("\u006c", _efe.ToPdfObject())
	}
	_fggg := _cd.MakeDict()
	_fggg.Set("\u0046\u0052\u004d", _efe.ToPdfObject())
	_ecfg := _bb.NewPdfPageResources()
	_ecfg.ProcSet = _cd.MakeArray(_cd.MakeName("\u0050\u0044\u0046"))
	_ecfg.XObject = _fggg
	_edgd := _fdf - 2
	_cfge := _gfg - 2
	_gggf.Add_q()
	_gggf.Add_re(1, 1, _edgd, _cfge)
	_gggf.Add_W()
	_gggf.Add_n()
	_edgd -= 2
	_cfge -= 2
	_gggf.Add_q()
	_gggf.Add_re(2, 2, _edgd, _cfge)
	_gggf.Add_W()
	_gggf.Add_n()
	_gabg := _dg.Min(_edgd/float64(opt.Image.Width), _cfge/float64(opt.Image.Height))
	_gggf.Add_cm(_gabg, 0, 0, _gabg, (_fdf/2)-(float64(opt.Image.Width)*_gabg/2)+2, 2)
	_gggf.Add_Do("\u0046\u0052\u004d")
	_gggf.Add_Q()
	_gggf.Add_Q()
	_efd := _bb.NewXObjectForm()
	_efd.FormType = _cd.MakeInteger(1)
	_efd.Resources = _ecfg
	_efd.BBox = _cd.MakeArrayFromFloats([]float64{0, 0, _fdf, _gfg})
	_efd.Matrix = _cd.MakeArrayFromFloats([]float64{1.0, 0.0, 0.0, 1.0, 0.0, 0.0})
	_efd.SetContentStream(_gggf.Bytes(), _bed())
	_daea := _cd.MakeDict()
	_daea.Set("\u004e", _efd.ToPdfObject())
	_aeb.AP = _daea
	_fbc.Annotations = append(_fbc.Annotations, _aeb)
	return _fbc, nil
}

// CreateRectangleAnnotation creates a rectangle annotation object that can be added to page PDF annotations.
func CreateRectangleAnnotation(rectDef RectangleAnnotationDef) (*_bb.PdfAnnotation, error) {
	_feeb := _bb.NewPdfAnnotationSquare()
	if rectDef.BorderEnabled {
		_agb, _bdgc, _cdf := rectDef.BorderColor.R(), rectDef.BorderColor.G(), rectDef.BorderColor.B()
		_feeb.C = _cd.MakeArrayFromFloats([]float64{_agb, _bdgc, _cdf})
		_fafcc := _bb.NewBorderStyle()
		_fafcc.SetBorderWidth(rectDef.BorderWidth)
		_feeb.BS = _fafcc.ToPdfObject()
	}
	if rectDef.FillEnabled {
		_bgee, _acgbf, _fbeg := rectDef.FillColor.R(), rectDef.FillColor.G(), rectDef.FillColor.B()
		_feeb.IC = _cd.MakeArrayFromFloats([]float64{_bgee, _acgbf, _fbeg})
	} else {
		_feeb.IC = _cd.MakeArrayFromIntegers([]int{})
	}
	if rectDef.Opacity < 1.0 {
		_feeb.CA = _cd.MakeFloat(rectDef.Opacity)
	}
	_fcdcg, _cfeb, _fdbd := _gde(rectDef)
	if _fdbd != nil {
		return nil, _fdbd
	}
	_feeb.AP = _fcdcg
	_feeb.Rect = _cd.MakeArrayFromFloats([]float64{_cfeb.Llx, _cfeb.Lly, _cfeb.Urx, _cfeb.Ury})
	return _feeb.PdfAnnotation, nil
}
func _fecd(_ceca _cd.PdfObject, _aecd *_bb.PdfPageResources) (*_cd.PdfObjectName, float64, bool) {
	var (
		_ebge *_cd.PdfObjectName
		_gea  float64
		_gedg bool
	)
	if _gdgb, _fcg := _cd.GetDict(_ceca); _fcg && _gdgb != nil {
		_dbee := _cd.TraceToDirectObject(_gdgb.Get("\u004e"))
		switch _dbac := _dbee.(type) {
		case *_cd.PdfObjectStream:
			_aaf, _gebf := _cd.DecodeStream(_dbac)
			if _gebf != nil {
				_d.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0063\u006f\u006e\u0074e\u006e\u0074\u0020\u0073\u0074r\u0065\u0061m\u003a\u0020\u0025\u0076", _gebf.Error())
				return nil, 0, false
			}
			_eaad, _gebf := _bg.NewContentStreamParser(string(_aaf)).Parse()
			if _gebf != nil {
				_d.Log.Debug("\u0045\u0052R\u004f\u0052\u0020\u0075n\u0061\u0062l\u0065\u0020\u0070\u0061\u0072\u0073\u0065\u0020c\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061m\u003a\u0020\u0025\u0076", _gebf.Error())
				return nil, 0, false
			}
			_acca := _bg.NewContentStreamProcessor(*_eaad)
			_acca.AddHandler(_bg.HandlerConditionEnumOperand, "\u0054\u0066", func(_bffe *_bg.ContentStreamOperation, _gffd _bg.GraphicsState, _aggg *_bb.PdfPageResources) error {
				if len(_bffe.Params) == 2 {
					if _cccb, _aag := _cd.GetName(_bffe.Params[0]); _aag {
						_ebge = _cccb
					}
					if _gga, _gede := _cd.GetNumberAsFloat(_bffe.Params[1]); _gede == nil {
						_gea = _gga
					}
					_gedg = true
					return _bg.ErrEarlyExit
				}
				return nil
			})
			_acca.Process(_aecd)
			return _ebge, _gea, _gedg
		}
	}
	return nil, 0, false
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
	Font *_bb.PdfFont

	// FontSize specifies the size of the text content.
	FontSize float64

	// LineHeight specifies the height of a line of text in the appearance annotation.
	LineHeight float64

	// TextColor represents the color of the text content displayed.
	TextColor _bb.PdfColor

	// FillColor represents the background color of the appearance annotation area.
	FillColor _bb.PdfColor

	// BorderSize represents border size of the appearance annotation area.
	BorderSize float64

	// BorderColor represents the border color of the appearance annotation area.
	BorderColor _bb.PdfColor

	// WatermarkImage specifies the image used as a watermark that will be rendered
	// behind the signature.
	WatermarkImage _bga.Image

	// Image represents the image used for the signature appearance.
	Image _bga.Image

	// Encoder specifies the image encoder used for image signature. Defaults to flate encoder.
	Encoder _cd.StreamEncoder

	// ImagePosition specifies the image location relative to the text signature.
	ImagePosition SignatureImagePosition
}

// NewSignatureField returns a new signature field with a visible appearance
// containing the specified signature lines and styled according to the
// specified options.
func NewSignatureField(signature *_bb.PdfSignature, lines []*SignatureLine, opts *SignatureFieldOpts) (*_bb.PdfFieldSignature, error) {
	if signature == nil {
		return nil, _f.New("\u0073\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	_eaec, _aaef := _abgf(lines, opts)
	if _aaef != nil {
		return nil, _aaef
	}
	_baf := _bb.NewPdfFieldSignature(signature)
	_baf.Rect = _cd.MakeArrayFromFloats(opts.Rect)
	_baf.AP = _eaec
	return _baf, nil
}

type quadding int

// CreateLineAnnotation creates a line annotation object that can be added to page PDF annotations.
func CreateLineAnnotation(lineDef LineAnnotationDef) (*_bb.PdfAnnotation, error) {
	_ecce := _bb.NewPdfAnnotationLine()
	_ecce.L = _cd.MakeArrayFromFloats([]float64{lineDef.X1, lineDef.Y1, lineDef.X2, lineDef.Y2})
	_bfdc := _cd.MakeName("\u004e\u006f\u006e\u0065")
	if lineDef.LineEndingStyle1 == _fb.LineEndingStyleArrow {
		_bfdc = _cd.MakeName("C\u006c\u006f\u0073\u0065\u0064\u0041\u0072\u0072\u006f\u0077")
	}
	_cega := _cd.MakeName("\u004e\u006f\u006e\u0065")
	if lineDef.LineEndingStyle2 == _fb.LineEndingStyleArrow {
		_cega = _cd.MakeName("C\u006c\u006f\u0073\u0065\u0064\u0041\u0072\u0072\u006f\u0077")
	}
	_ecce.LE = _cd.MakeArray(_bfdc, _cega)
	if lineDef.Opacity < 1.0 {
		_ecce.CA = _cd.MakeFloat(lineDef.Opacity)
	}
	_fffg, _cfgb, _aegg := lineDef.LineColor.R(), lineDef.LineColor.G(), lineDef.LineColor.B()
	_ecce.IC = _cd.MakeArrayFromFloats([]float64{_fffg, _cfgb, _aegg})
	_ecce.C = _cd.MakeArrayFromFloats([]float64{_fffg, _cfgb, _aegg})
	_eagf := _bb.NewBorderStyle()
	_eagf.SetBorderWidth(lineDef.LineWidth)
	_ecce.BS = _eagf.ToPdfObject()
	_dffa, _gbdbg, _ecbf := _eaed(lineDef)
	if _ecbf != nil {
		return nil, _ecbf
	}
	_ecce.AP = _dffa
	_ecce.Rect = _cd.MakeArrayFromFloats([]float64{_gbdbg.Llx, _gbdbg.Lly, _gbdbg.Urx, _gbdbg.Ury})
	return _ecce.PdfAnnotation, nil
}

const (
	_fag   = 1
	_gad   = 2
	_gbfc  = 4
	_fccb  = 8
	_gdag  = 16
	_cbcf  = 32
	_ebca  = 64
	_fbda  = 128
	_gdbd  = 256
	_agdaf = 512
	_dadd  = 1024
	_aaff  = 2048
	_ffdf  = 4096
)

func _fec(_aaa *_bb.PdfAnnotationWidget, _bdg *_bb.PdfFieldText, _bdf *_bb.PdfPageResources, _ebd AppearanceStyle) (*_cd.PdfObjectDictionary, error) {
	_bde := _bb.NewPdfPageResources()
	_gbe, _fba := _cd.GetArray(_aaa.Rect)
	if !_fba {
		return nil, _f.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_gce, _bba := _bb.NewPdfRectangle(*_gbe)
	if _bba != nil {
		return nil, _bba
	}
	_dca, _bbdb := _gce.Width(), _gce.Height()
	_dba, _gca := _dca, _bbdb
	_gbb, _egga := _cd.GetDict(_aaa.MK)
	if _egga {
		_eab, _ := _cd.GetDict(_aaa.BS)
		_cag := _ebd.applyAppearanceCharacteristics(_gbb, _eab, nil)
		if _cag != nil {
			return nil, _cag
		}
	}
	_cff, _egga := _cd.GetIntVal(_bdg.MaxLen)
	if !_egga {
		return nil, _f.New("\u006d\u0061\u0078\u006c\u0065\u006e\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
	}
	if _cff <= 0 {
		return nil, _f.New("\u006d\u0061\u0078\u004c\u0065\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_ecf := _dca / float64(_cff)
	_ggfd, _bba := _bg.NewContentStreamParser(_dbf(_bdg.PdfField)).Parse()
	if _bba != nil {
		return nil, _bba
	}
	_ecb := _bg.NewContentCreator()
	if _ebd.BorderSize > 0 {
		_ddc(_ecb, _ebd, _dca, _bbdb)
	}
	if _ebd.DrawAlignmentReticle {
		_ed := _ebd
		_ed.BorderSize = 0.2
		_deggd(_ecb, _ed, _dca, _bbdb)
	}
	_ecb.Add_BMC("\u0054\u0078")
	_ecb.Add_q()
	_, _bbdb = _ebd.applyRotation(_gbb, _dca, _bbdb, _ecb)
	_ecb.Add_BT()
	_cea, _gbdb, _bba := _ebd.processDA(_bdg.PdfField, _ggfd, _bdf, _bde, _ecb)
	if _bba != nil {
		return nil, _bba
	}
	_faf := _cea.Font
	_caf := _cd.MakeName(_cea.Name)
	_degfg := _cea.Size
	_befc := _degfg == 0
	if _befc && _gbdb {
		_degfg = _bbdb * _ebd.AutoFontSizeFraction
	}
	_cga := _faf.Encoder()
	if _cga == nil {
		_d.Log.Debug("\u0057\u0041RN\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0069\u0073\u0020\u006e\u0069l\u002e\u0020\u0041\u0073s\u0075\u006d\u0069\u006eg \u0069\u0064e\u006et\u0069\u0074\u0079\u0020\u0065\u006ec\u006f\u0064\u0065r\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
		_cga = _a.NewIdentityTextEncoder("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048")
	}
	var _gaf string
	if _bbcf, _cef := _cd.GetString(_bdg.V); _cef {
		_gaf = _bbcf.Decoded()
	}
	_ecb.Add_Tf(*_caf, _degfg)
	var _ab float64
	for _, _eff := range _gaf {
		_bda, _dbc := _faf.GetRuneMetrics(_eff)
		if !_dbc {
			_d.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0052\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0066\u006fn\u0074\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067 \u006f\u0076\u0065\u0072", _eff)
			continue
		}
		_adcd := _bda.Wy
		if int(_adcd) <= 0 {
			_adcd = _bda.Wx
		}
		if _adcd > _ab {
			_ab = _adcd
		}
	}
	if int(_ab) == 0 {
		_d.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074o\u0020\u0064\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u006d\u0061x\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u0073\u0069\u007a\u0065\u0020- \u0075\u0073\u0069\u006e\u0067\u0020\u0031\u0030\u0030\u0030")
		_ab = 1000
	}
	_afb, _bba := _faf.GetFontDescriptor()
	if _bba != nil {
		_d.Log.Debug("\u0045\u0072ro\u0072\u003a\u0020U\u006e\u0061\u0062\u006ce t\u006f g\u0065\u0074\u0020\u0066\u006f\u006e\u0074 d\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
	}
	var _gab float64
	if _afb != nil {
		_gab, _bba = _afb.GetCapHeight()
		if _bba != nil {
			_d.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0043\u0061\u0070\u0048\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _bba)
		}
	}
	if int(_gab) <= 0 {
		_d.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u0061\u0070\u0048e\u0069\u0067\u0068\u0074\u0020\u006e\u006ft \u0061\u0076\u0061\u0069l\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065tt\u0069\u006eg\u0020\u0074\u006f\u0020\u0031\u0030\u0030\u0030")
		_gab = 1000.0
	}
	_fecb := _gab / 1000.0 * _degfg
	_ggcd := 0.0
	_geg := 1.0 * _degfg * (_ab / 1000.0)
	{
		_fge := _geg
		if _befc && _ggcd+_fge > _bbdb {
			_degfg = 0.95 * (_bbdb - _ggcd)
			_fecb = _gab / 1000.0 * _degfg
		}
		if _bbdb > _fecb {
			_ggcd = (_bbdb - _fecb) / 2.0
		}
	}
	_ecb.Add_Td(0, _ggcd)
	if _dbd, _bfdf := _cd.GetIntVal(_bdg.Q); _bfdf {
		switch _dbd {
		case 2:
			if len(_gaf) < _cff {
				_caa := float64(_cff-len(_gaf)) * _ecf
				_ecb.Add_Td(_caa, 0)
			}
		}
	}
	for _bggg, _cafd := range _gaf {
		_dfc := _dgg
		if _ebd.MarginLeft != nil {
			_dfc = *_ebd.MarginLeft
		}
		_egbg := string(_cafd)
		if _cga != nil {
			_dee, _gdcc := _faf.GetRuneMetrics(_cafd)
			if !_gdcc {
				_d.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0052\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0066\u006fn\u0074\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067 \u006f\u0076\u0065\u0072", _cafd)
				continue
			}
			_egbg = string(_cga.Encode(_egbg))
			_afd := _degfg * _dee.Wx / 1000.0
			_cgaa := (_ecf - _afd) / 2
			_dfc = _cgaa
		}
		_ecb.Add_Td(_dfc, 0)
		_ecb.Add_Tj(*_cd.MakeString(_egbg))
		if _bggg != len(_gaf)-1 {
			_ecb.Add_Td(_ecf-_dfc, 0)
		}
	}
	_ecb.Add_ET()
	_ecb.Add_Q()
	_ecb.Add_EMC()
	_cba := _bb.NewXObjectForm()
	_cba.Resources = _bde
	_cba.BBox = _cd.MakeArrayFromFloats([]float64{0, 0, _dba, _gca})
	_cba.SetContentStream(_ecb.Bytes(), _bed())
	_ffd := _cd.MakeDict()
	_ffd.Set("\u004e", _cba.ToPdfObject())
	return _ffd, nil
}

// ImageFieldOptions defines optional parameters for a push button with image attach capability form field.
type ImageFieldOptions struct {
	Image *_bb.Image
	_degc AppearanceStyle
}

// CheckboxFieldOptions defines optional parameters for a checkbox field a form.
type CheckboxFieldOptions struct{ Checked bool }

func _fgge(_gacd *_bb.PdfFieldButton, _aggbb *_bb.PdfAnnotationWidget, _bfe AppearanceStyle) (*_cd.PdfObjectDictionary, error) {
	_gecb, _dgagf := _cd.GetArray(_aggbb.Rect)
	if !_dgagf {
		return nil, _f.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_fae, _bac := _bb.NewPdfRectangle(*_gecb)
	if _bac != nil {
		return nil, _bac
	}
	_baae, _deef := _fae.Width(), _fae.Height()
	_eacc := _bg.NewContentCreator()
	if _bfe.BorderSize > 0 {
		_ddc(_eacc, _bfe, _baae, _deef)
	}
	if _bfe.DrawAlignmentReticle {
		_aeae := _bfe
		_aeae.BorderSize = 0.2
		_deggd(_eacc, _aeae, _baae, _deef)
	}
	_acbc := _gacd.GetFillImage()
	_daee, _bac := _gacdc(_baae, _deef, _acbc, _bfe)
	if _bac != nil {
		return nil, _bac
	}
	_bedf, _efed := _cd.GetDict(_aggbb.MK)
	if _efed {
		_bedf.Set("\u006c", _daee.ToPdfObject())
	}
	_aef := _cd.MakeDict()
	_aef.Set("\u0046\u0052\u004d", _daee.ToPdfObject())
	_daead := _bb.NewPdfPageResources()
	_daead.ProcSet = _cd.MakeArray(_cd.MakeName("\u0050\u0044\u0046"))
	_daead.XObject = _aef
	_fcgg := _baae - 2
	_gfff := _deef - 2
	_eacc.Add_q()
	_eacc.Add_re(1, 1, _fcgg, _gfff)
	_eacc.Add_W()
	_eacc.Add_n()
	_fcgg -= 2
	_gfff -= 2
	_eacc.Add_q()
	_eacc.Add_re(2, 2, _fcgg, _gfff)
	_eacc.Add_W()
	_eacc.Add_n()
	_ffed := _dg.Min(_fcgg/float64(_acbc.Width), _gfff/float64(_acbc.Height))
	_eacc.Add_cm(_ffed, 0, 0, _ffed, (_baae/2)-(float64(_acbc.Width)*_ffed/2)+2, 2)
	_eacc.Add_Do("\u0046\u0052\u004d")
	_eacc.Add_Q()
	_eacc.Add_Q()
	_dcc := _bb.NewXObjectForm()
	_dcc.FormType = _cd.MakeInteger(1)
	_dcc.Resources = _daead
	_dcc.BBox = _cd.MakeArrayFromFloats([]float64{0, 0, _baae, _deef})
	_dcc.Matrix = _cd.MakeArrayFromFloats([]float64{1.0, 0.0, 0.0, 1.0, 0.0, 0.0})
	_dcc.SetContentStream(_eacc.Bytes(), _bed())
	_eca := _cd.MakeDict()
	_eca.Set("\u004e", _dcc.ToPdfObject())
	return _eca, nil
}

// CreateCircleAnnotation creates a circle/ellipse annotation object with appearance stream that can be added to
// page PDF annotations.
func CreateCircleAnnotation(circDef CircleAnnotationDef) (*_bb.PdfAnnotation, error) {
	_ae := _bb.NewPdfAnnotationCircle()
	if circDef.BorderEnabled {
		_ff, _fe, _ba := circDef.BorderColor.R(), circDef.BorderColor.G(), circDef.BorderColor.B()
		_ae.C = _cd.MakeArrayFromFloats([]float64{_ff, _fe, _ba})
		_bgc := _bb.NewBorderStyle()
		_bgc.SetBorderWidth(circDef.BorderWidth)
		_ae.BS = _bgc.ToPdfObject()
	}
	if circDef.FillEnabled {
		_ffg, _ag, _gc := circDef.FillColor.R(), circDef.FillColor.G(), circDef.FillColor.B()
		_ae.IC = _cd.MakeArrayFromFloats([]float64{_ffg, _ag, _gc})
	} else {
		_ae.IC = _cd.MakeArrayFromIntegers([]int{})
	}
	if circDef.Opacity < 1.0 {
		_ae.CA = _cd.MakeFloat(circDef.Opacity)
	}
	_fff, _gcd, _fd := _e(circDef)
	if _fd != nil {
		return nil, _fd
	}
	_ae.AP = _fff
	_ae.Rect = _cd.MakeArrayFromFloats([]float64{_gcd.Llx, _gcd.Lly, _gcd.Urx, _gcd.Ury})
	return _ae.PdfAnnotation, nil
}
func _deggd(_gbf *_bg.ContentCreator, _dcda AppearanceStyle, _bfdb, _gecd float64) {
	_gbf.Add_q().Add_re(0, 0, _bfdb, _gecd).Add_re(0, _gecd/2, _bfdb, _gecd/2).Add_re(0, 0, _bfdb, _gecd).Add_re(_bfdb/2, 0, _bfdb/2, _gecd).Add_w(_dcda.BorderSize).SetStrokingColor(_dcda.BorderColor).SetNonStrokingColor(_dcda.FillColor).Add_B().Add_Q()
}
