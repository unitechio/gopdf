package annotator

import (
	_ee "bytes"
	_bb "errors"
	_ga "image"
	_d "math"
	_ef "strings"
	_f "unicode"

	_g "bitbucket.org/shenghui0779/gopdf/common"
	_c "bitbucket.org/shenghui0779/gopdf/contentstream"
	_efb "bitbucket.org/shenghui0779/gopdf/contentstream/draw"
	_gac "bitbucket.org/shenghui0779/gopdf/core"
	_b "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_eg "bitbucket.org/shenghui0779/gopdf/model"
)

func _ff(_ecc CircleAnnotationDef, _gacf string) ([]byte, *_eg.PdfRectangle, *_eg.PdfRectangle, error) {
	_ggd := _efb.Circle{X: _ecc.X, Y: _ecc.Y, Width: _ecc.Width, Height: _ecc.Height, FillEnabled: _ecc.FillEnabled, FillColor: _ecc.FillColor, BorderEnabled: _ecc.BorderEnabled, BorderWidth: _ecc.BorderWidth, BorderColor: _ecc.BorderColor, Opacity: _ecc.Opacity}
	_ea, _dg, _cc := _ggd.Draw(_gacf)
	if _cc != nil {
		return nil, nil, nil, _cc
	}
	_bcf := &_eg.PdfRectangle{}
	_bcf.Llx = _ecc.X + _dg.Llx
	_bcf.Lly = _ecc.Y + _dg.Lly
	_bcf.Urx = _ecc.X + _dg.Urx
	_bcf.Ury = _ecc.Y + _dg.Ury
	return _ea, _dg, _bcf, nil
}
func _egf(_gad *_eg.PdfAnnotationWidget, _ddb *_eg.PdfFieldText, _bab *_eg.PdfPageResources, _gec AppearanceStyle) (*_gac.PdfObjectDictionary, error) {
	_eccb := _eg.NewPdfPageResources()
	_fcfg, _ecd := _gac.GetArray(_gad.Rect)
	if !_ecd {
		return nil, _bb.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_agb, _bed := _eg.NewPdfRectangle(*_fcfg)
	if _bed != nil {
		return nil, _bed
	}
	_dad, _eae := _agb.Width(), _agb.Height()
	_bg, _gade := _dad, _eae
	_fab, _bef := _gac.GetDict(_gad.MK)
	if _bef {
		_fee, _ := _gac.GetDict(_gad.BS)
		_ddbc := _gec.applyAppearanceCharacteristics(_fab, _fee, nil)
		if _ddbc != nil {
			return nil, _ddbc
		}
	}
	_geg, _bef := _gac.GetIntVal(_ddb.MaxLen)
	if !_bef {
		return nil, _bb.New("\u006d\u0061\u0078\u006c\u0065\u006e\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
	}
	if _geg <= 0 {
		return nil, _bb.New("\u006d\u0061\u0078\u004c\u0065\u006e\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	_baeb := _dad / float64(_geg)
	_dgd, _bed := _c.NewContentStreamParser(_bgg(_ddb.PdfField)).Parse()
	if _bed != nil {
		return nil, _bed
	}
	_feeb := _c.NewContentCreator()
	if _gec.BorderSize > 0 {
		_efc(_feeb, _gec, _dad, _eae)
	}
	if _gec.DrawAlignmentReticle {
		_gbg := _gec
		_gbg.BorderSize = 0.2
		_gdgd(_feeb, _gbg, _dad, _eae)
	}
	_feeb.Add_BMC("\u0054\u0078")
	_feeb.Add_q()
	_, _eae = _gec.applyRotation(_fab, _dad, _eae, _feeb)
	_feeb.Add_BT()
	_bcbg, _efbe, _bed := _gec.processDA(_ddb.PdfField, _dgd, _bab, _eccb, _feeb)
	if _bed != nil {
		return nil, _bed
	}
	_ae := _bcbg.Font
	_fag := _gac.MakeName(_bcbg.Name)
	_cfe := _bcbg.Size
	_fgd := _cfe == 0
	if _fgd && _efbe {
		_cfe = _eae * _gec.AutoFontSizeFraction
	}
	_fagf := _ae.Encoder()
	if _fagf == nil {
		_g.Log.Debug("\u0057\u0041RN\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0069\u0073\u0020\u006e\u0069l\u002e\u0020\u0041\u0073s\u0075\u006d\u0069\u006eg \u0069\u0064e\u006et\u0069\u0074\u0079\u0020\u0065\u006ec\u006f\u0064\u0065r\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
		_fagf = _b.NewIdentityTextEncoder("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048")
	}
	var _fge string
	if _edg, _bce := _gac.GetString(_ddb.V); _bce {
		_fge = _edg.Decoded()
	}
	_feeb.Add_Tf(*_fag, _cfe)
	var _aabe float64
	for _, _eef := range _fge {
		_dabd, _cad := _ae.GetRuneMetrics(_eef)
		if !_cad {
			_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0052\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0066\u006fn\u0074\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067 \u006f\u0076\u0065\u0072", _eef)
			continue
		}
		_fbc := _dabd.Wy
		if int(_fbc) <= 0 {
			_fbc = _dabd.Wx
		}
		if _fbc > _aabe {
			_aabe = _fbc
		}
	}
	if int(_aabe) == 0 {
		_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0074o\u0020\u0064\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u006d\u0061x\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u0073\u0069\u007a\u0065\u0020- \u0075\u0073\u0069\u006e\u0067\u0020\u0031\u0030\u0030\u0030")
		_aabe = 1000
	}
	_gcd, _bed := _ae.GetFontDescriptor()
	if _bed != nil {
		_g.Log.Debug("\u0045\u0072ro\u0072\u003a\u0020U\u006e\u0061\u0062\u006ce t\u006f g\u0065\u0074\u0020\u0066\u006f\u006e\u0074 d\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
	}
	var _bbcd float64
	if _gcd != nil {
		_bbcd, _bed = _gcd.GetCapHeight()
		if _bed != nil {
			_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0043\u0061\u0070\u0048\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _bed)
		}
	}
	if int(_bbcd) <= 0 {
		_g.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u0061\u0070\u0048e\u0069\u0067\u0068\u0074\u0020\u006e\u006ft \u0061\u0076\u0061\u0069l\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065tt\u0069\u006eg\u0020\u0074\u006f\u0020\u0031\u0030\u0030\u0030")
		_bbcd = 1000.0
	}
	_bbag := _bbcd / 1000.0 * _cfe
	_ecb := 0.0
	_gbf := 1.0 * _cfe * (_aabe / 1000.0)
	{
		_aae := _gbf
		if _fgd && _ecb+_aae > _eae {
			_cfe = 0.95 * (_eae - _ecb)
			_bbag = _bbcd / 1000.0 * _cfe
		}
		if _eae > _bbag {
			_ecb = (_eae - _bbag) / 2.0
		}
	}
	_feeb.Add_Td(0, _ecb)
	if _eeb, _db := _gac.GetIntVal(_ddb.Q); _db {
		switch _eeb {
		case 2:
			if len(_fge) < _geg {
				_eefc := float64(_geg-len(_fge)) * _baeb
				_feeb.Add_Td(_eefc, 0)
			}
		}
	}
	for _cb, _eadc := range _fge {
		_dde := _cf
		if _gec.MarginLeft != nil {
			_dde = *_gec.MarginLeft
		}
		_ce := string(_eadc)
		if _fagf != nil {
			_adgg, _cafd := _ae.GetRuneMetrics(_eadc)
			if !_cafd {
				_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0052\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0066\u006fn\u0074\u003a\u0020\u0025\u0076\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067 \u006f\u0076\u0065\u0072", _eadc)
				continue
			}
			_ce = string(_fagf.Encode(_ce))
			_aed := _cfe * _adgg.Wx / 1000.0
			_efg := (_baeb - _aed) / 2
			_dde = _efg
		}
		_feeb.Add_Td(_dde, 0)
		_feeb.Add_Tj(*_gac.MakeString(_ce))
		if _cb != len(_fge)-1 {
			_feeb.Add_Td(_baeb-_dde, 0)
		}
	}
	_feeb.Add_ET()
	_feeb.Add_Q()
	_feeb.Add_EMC()
	_egg := _eg.NewXObjectForm()
	_egg.Resources = _eccb
	_egg.BBox = _gac.MakeArrayFromFloats([]float64{0, 0, _bg, _gade})
	_egg.SetContentStream(_feeb.Bytes(), _ddba())
	_dcc := _gac.MakeDict()
	_dcc.Set("\u004e", _egg.ToPdfObject())
	return _dcc, nil
}
func _aeec(_ceb RectangleAnnotationDef, _dcca string) ([]byte, *_eg.PdfRectangle, *_eg.PdfRectangle, error) {
	_bcc := _efb.Rectangle{X: 0, Y: 0, Width: _ceb.Width, Height: _ceb.Height, FillEnabled: _ceb.FillEnabled, FillColor: _ceb.FillColor, BorderEnabled: _ceb.BorderEnabled, BorderWidth: 2 * _ceb.BorderWidth, BorderColor: _ceb.BorderColor, Opacity: _ceb.Opacity}
	_cfgc, _cadg, _abg := _bcc.Draw(_dcca)
	if _abg != nil {
		return nil, nil, nil, _abg
	}
	_cbf := &_eg.PdfRectangle{}
	_cbf.Llx = _ceb.X + _cadg.Llx
	_cbf.Lly = _ceb.Y + _cadg.Lly
	_cbf.Urx = _ceb.X + _cadg.Urx
	_cbf.Ury = _ceb.Y + _cadg.Ury
	return _cfgc, _cadg, _cbf, nil
}

// SignatureLine represents a line of information in the signature field appearance.
type SignatureLine struct {
	Desc string
	Text string
}

// CreateLineAnnotation creates a line annotation object that can be added to page PDF annotations.
func CreateLineAnnotation(lineDef LineAnnotationDef) (*_eg.PdfAnnotation, error) {
	_cdccb := _eg.NewPdfAnnotationLine()
	_cdccb.L = _gac.MakeArrayFromFloats([]float64{lineDef.X1, lineDef.Y1, lineDef.X2, lineDef.Y2})
	_bdec := _gac.MakeName("\u004e\u006f\u006e\u0065")
	if lineDef.LineEndingStyle1 == _efb.LineEndingStyleArrow {
		_bdec = _gac.MakeName("C\u006c\u006f\u0073\u0065\u0064\u0041\u0072\u0072\u006f\u0077")
	}
	_bfcf := _gac.MakeName("\u004e\u006f\u006e\u0065")
	if lineDef.LineEndingStyle2 == _efb.LineEndingStyleArrow {
		_bfcf = _gac.MakeName("C\u006c\u006f\u0073\u0065\u0064\u0041\u0072\u0072\u006f\u0077")
	}
	_cdccb.LE = _gac.MakeArray(_bdec, _bfcf)
	if lineDef.Opacity < 1.0 {
		_cdccb.CA = _gac.MakeFloat(lineDef.Opacity)
	}
	_fdfc, _fff, _ecgf := lineDef.LineColor.R(), lineDef.LineColor.G(), lineDef.LineColor.B()
	_cdccb.IC = _gac.MakeArrayFromFloats([]float64{_fdfc, _fff, _ecgf})
	_cdccb.C = _gac.MakeArrayFromFloats([]float64{_fdfc, _fff, _ecgf})
	_bcdf := _eg.NewBorderStyle()
	_bcdf.SetBorderWidth(lineDef.LineWidth)
	_cdccb.BS = _bcdf.ToPdfObject()
	_bga, _gfdf, _bgc := _afge(lineDef)
	if _bgc != nil {
		return nil, _bgc
	}
	_cdccb.AP = _bga
	_cdccb.Rect = _gac.MakeArrayFromFloats([]float64{_gfdf.Llx, _gfdf.Lly, _gfdf.Urx, _gfdf.Ury})
	return _cdccb.PdfAnnotation, nil
}

// NewSignatureFieldOpts returns a new initialized instance of options
// used to generate a signature appearance.
func NewSignatureFieldOpts() *SignatureFieldOpts {
	return &SignatureFieldOpts{Font: _eg.DefaultFont(), FontSize: 10, LineHeight: 1, AutoSize: true, TextColor: _eg.NewPdfColorDeviceGray(0), BorderColor: _eg.NewPdfColorDeviceGray(0), FillColor: _eg.NewPdfColorDeviceGray(1), Encoder: _gac.NewFlateEncoder(), ImagePosition: SignatureImageLeft}
}

// CreateCircleAnnotation creates a circle/ellipse annotation object with appearance stream that can be added to
// page PDF annotations.
func CreateCircleAnnotation(circDef CircleAnnotationDef) (*_eg.PdfAnnotation, error) {
	_a := _eg.NewPdfAnnotationCircle()
	if circDef.BorderEnabled {
		_bc, _bf, _dd := circDef.BorderColor.R(), circDef.BorderColor.G(), circDef.BorderColor.B()
		_a.C = _gac.MakeArrayFromFloats([]float64{_bc, _bf, _dd})
		_fc := _eg.NewBorderStyle()
		_fc.SetBorderWidth(circDef.BorderWidth)
		_a.BS = _fc.ToPdfObject()
	}
	if circDef.FillEnabled {
		_aa, _ec, _cd := circDef.FillColor.R(), circDef.FillColor.G(), circDef.FillColor.B()
		_a.IC = _gac.MakeArrayFromFloats([]float64{_aa, _ec, _cd})
	} else {
		_a.IC = _gac.MakeArrayFromIntegers([]int{})
	}
	if circDef.Opacity < 1.0 {
		_a.CA = _gac.MakeFloat(circDef.Opacity)
	}
	_be, _cdc, _fe := _da(circDef)
	if _fe != nil {
		return nil, _fe
	}
	_a.AP = _be
	_a.Rect = _gac.MakeArrayFromFloats([]float64{_cdc.Llx, _cdc.Lly, _cdc.Urx, _cdc.Ury})
	return _a.PdfAnnotation, nil
}

// SetStyle applies appearance `style` to `fa`.
func (_cbbd *ImageFieldAppearance) SetStyle(style AppearanceStyle) { _cbbd._gfea = &style }

// RectangleAnnotationDef is a rectangle defined with a specified Width and Height and a lower left corner at (X,Y).
// The rectangle can optionally have a border and a filling color.
// The Width/Height includes the border (if any specified).
type RectangleAnnotationDef struct {
	X             float64
	Y             float64
	Width         float64
	Height        float64
	FillEnabled   bool
	FillColor     *_eg.PdfColorDeviceRGB
	BorderEnabled bool
	BorderWidth   float64
	BorderColor   *_eg.PdfColorDeviceRGB
	Opacity       float64
}

const (
	_ba  quadding = 0
	_ag  quadding = 1
	_ege quadding = 2
	_cf  float64  = 2.0
)

func _gdgd(_agfd *_c.ContentCreator, _gaca AppearanceStyle, _cabc, _cdee float64) {
	_agfd.Add_q().Add_re(0, 0, _cabc, _cdee).Add_re(0, _cdee/2, _cabc, _cdee/2).Add_re(0, 0, _cabc, _cdee).Add_re(_cabc/2, 0, _cabc/2, _cdee).Add_w(_gaca.BorderSize).SetStrokingColor(_gaca.BorderColor).SetNonStrokingColor(_gaca.FillColor).Add_B().Add_Q()
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
	Font *_eg.PdfFont

	// FontSize specifies the size of the text content.
	FontSize float64

	// LineHeight specifies the height of a line of text in the appearance annotation.
	LineHeight float64

	// TextColor represents the color of the text content displayed.
	TextColor _eg.PdfColor

	// FillColor represents the background color of the appearance annotation area.
	FillColor _eg.PdfColor

	// BorderSize represents border size of the appearance annotation area.
	BorderSize float64

	// BorderColor represents the border color of the appearance annotation area.
	BorderColor _eg.PdfColor

	// WatermarkImage specifies the image used as a watermark that will be rendered
	// behind the signature.
	WatermarkImage _ga.Image

	// Image represents the image used for the signature appearance.
	Image _ga.Image

	// Encoder specifies the image encoder used for image signature. Defaults to flate encoder.
	Encoder _gac.StreamEncoder

	// ImagePosition specifies the image location relative to the text signature.
	ImagePosition SignatureImagePosition
}

func _da(_gg CircleAnnotationDef) (*_gac.PdfObjectDictionary, *_eg.PdfRectangle, error) {
	_fd := _eg.NewXObjectForm()
	_fd.Resources = _eg.NewPdfPageResources()
	_aad := ""
	if _gg.Opacity < 1.0 {
		_dc := _gac.MakeDict()
		_dc.Set("\u0063\u0061", _gac.MakeFloat(_gg.Opacity))
		_dc.Set("\u0043\u0041", _gac.MakeFloat(_gg.Opacity))
		_fg := _fd.Resources.AddExtGState("\u0067\u0073\u0031", _dc)
		if _fg != nil {
			_g.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020\u0065\u0078\u0074g\u0073\u0074\u0061t\u0065 \u0067\u0073\u0031")
			return nil, nil, _fg
		}
		_aad = "\u0067\u0073\u0031"
	}
	_df, _bd, _bcg, _cg := _ff(_gg, _aad)
	if _cg != nil {
		return nil, nil, _cg
	}
	_cg = _fd.SetContentStream(_df, nil)
	if _cg != nil {
		return nil, nil, _cg
	}
	_fd.BBox = _bd.ToPdfObject()
	_ed := _gac.MakeDict()
	_ed.Set("\u004e", _fd.ToPdfObject())
	return _ed, _bcg, nil
}

// Style returns the appearance style of `fa`. If not specified, returns default style.
func (_eaff ImageFieldAppearance) Style() AppearanceStyle {
	if _eaff._gfea != nil {
		return *_eaff._gfea
	}
	return AppearanceStyle{BorderSize: 0.0, BorderColor: _eg.NewPdfColorDeviceGray(0), FillColor: _eg.NewPdfColorDeviceGray(1), DrawAlignmentReticle: false}
}

// ImageFieldOptions defines optional parameters for a push button with image attach capability form field.
type ImageFieldOptions struct {
	Image *_eg.Image
	_bfgc AppearanceStyle
}

func (_ddc *AppearanceStyle) applyRotation(_dfdd *_gac.PdfObjectDictionary, _abe, _ced float64, _acd *_c.ContentCreator) (float64, float64) {
	if !_ddc.AllowMK {
		return _abe, _ced
	}
	if _dfdd == nil {
		return _abe, _ced
	}
	_gfga, _ := _gac.GetNumberAsFloat(_dfdd.Get("\u0052"))
	if _gfga == 0 {
		return _abe, _ced
	}
	_bee := -_gfga
	_fdb := _efb.Path{Points: []_efb.Point{_efb.NewPoint(0, 0).Rotate(_bee), _efb.NewPoint(_abe, 0).Rotate(_bee), _efb.NewPoint(0, _ced).Rotate(_bee), _efb.NewPoint(_abe, _ced).Rotate(_bee)}}.GetBoundingBox()
	_acd.RotateDeg(_gfga)
	_acd.Translate(_fdb.X, _fdb.Y)
	return _fdb.Width, _fdb.Height
}

type quadding int

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
	_dfe                 *AppearanceStyle
}

func _gfd(_ddbac *_eg.PdfFieldButton, _fdee *_eg.PdfAnnotationWidget, _eceg AppearanceStyle) (*_gac.PdfObjectDictionary, error) {
	_deeae, _dgde := _gac.GetArray(_fdee.Rect)
	if !_dgde {
		return nil, _bb.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_dggc, _cca := _eg.NewPdfRectangle(*_deeae)
	if _cca != nil {
		return nil, _cca
	}
	_eaeb, _fbed := _dggc.Width(), _dggc.Height()
	_babgc := _c.NewContentCreator()
	if _eceg.BorderSize > 0 {
		_efc(_babgc, _eceg, _eaeb, _fbed)
	}
	if _eceg.DrawAlignmentReticle {
		_eca := _eceg
		_eca.BorderSize = 0.2
		_gdgd(_babgc, _eca, _eaeb, _fbed)
	}
	_gccf := _ddbac.GetFillImage()
	_ddab, _cca := _cbd(_eaeb, _fbed, _gccf, _eceg)
	if _cca != nil {
		return nil, _cca
	}
	_dfdgb, _fecg := _gac.GetDict(_fdee.MK)
	if _fecg {
		_dfdgb.Set("\u006c", _ddab.ToPdfObject())
	}
	_cefc := _gac.MakeDict()
	_cefc.Set("\u0046\u0052\u004d", _ddab.ToPdfObject())
	_bbe := _eg.NewPdfPageResources()
	_bbe.ProcSet = _gac.MakeArray(_gac.MakeName("\u0050\u0044\u0046"))
	_bbe.XObject = _cefc
	_eggf := _eaeb - 2
	_bbec := _fbed - 2
	_babgc.Add_q()
	_babgc.Add_re(1, 1, _eggf, _bbec)
	_babgc.Add_W()
	_babgc.Add_n()
	_eggf -= 2
	_bbec -= 2
	_babgc.Add_q()
	_babgc.Add_re(2, 2, _eggf, _bbec)
	_babgc.Add_W()
	_babgc.Add_n()
	_cgea := _d.Min(_eggf/float64(_gccf.Width), _bbec/float64(_gccf.Height))
	_babgc.Add_cm(_cgea, 0, 0, _cgea, (_eaeb/2)-(float64(_gccf.Width)*_cgea/2)+2, 2)
	_babgc.Add_Do("\u0046\u0052\u004d")
	_babgc.Add_Q()
	_babgc.Add_Q()
	_fdfb := _eg.NewXObjectForm()
	_fdfb.FormType = _gac.MakeInteger(1)
	_fdfb.Resources = _bbe
	_fdfb.BBox = _gac.MakeArrayFromFloats([]float64{0, 0, _eaeb, _fbed})
	_fdfb.Matrix = _gac.MakeArrayFromFloats([]float64{1.0, 0.0, 0.0, 1.0, 0.0, 0.0})
	_fdfb.SetContentStream(_babgc.Bytes(), _ddba())
	_efbb := _gac.MakeDict()
	_efbb.Set("\u004e", _fdfb.ToPdfObject())
	return _efbb, nil
}

// GenerateAppearanceDict generates an appearance dictionary for widget annotation `wa` for the `field` in `form`.
// Implements interface model.FieldAppearanceGenerator.
func (_cdd FieldAppearance) GenerateAppearanceDict(form *_eg.PdfAcroForm, field *_eg.PdfField, wa *_eg.PdfAnnotationWidget) (*_gac.PdfObjectDictionary, error) {
	_g.Log.Trace("\u0047\u0065n\u0065\u0072\u0061\u0074e\u0041\u0070p\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0044i\u0063\u0074\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u0020\u0056:\u0020\u0025\u002b\u0076", field.PartialName(), field.V)
	_, _ca := field.GetContext().(*_eg.PdfFieldText)
	_dga, _dac := _gac.GetDict(wa.AP)
	if _dac && _cdd.OnlyIfMissing && (!_ca || !_cdd.RegenerateTextFields) {
		_g.Log.Trace("\u0041\u006c\u0072\u0065a\u0064\u0079\u0020\u0070\u006f\u0070\u0075\u006c\u0061\u0074e\u0064 \u002d\u0020\u0069\u0067\u006e\u006f\u0072i\u006e\u0067")
		return _dga, nil
	}
	if form.DR == nil {
		form.DR = _eg.NewPdfPageResources()
	}
	switch _fb := field.GetContext().(type) {
	case *_eg.PdfFieldText:
		_eb := _fb
		switch {
		case _eb.Flags().Has(_eg.FieldFlagPassword):
			return nil, nil
		case _eb.Flags().Has(_eg.FieldFlagFileSelect):
			return nil, nil
		case _eb.Flags().Has(_eg.FieldFlagComb):
			if _eb.MaxLen != nil {
				_fga, _eac := _egf(wa, _eb, form.DR, _cdd.Style())
				if _eac != nil {
					return nil, _eac
				}
				return _fga, nil
			}
		}
		_gaa, _ad := _ece(wa, _eb, form.DR, _cdd.Style())
		if _ad != nil {
			return nil, _ad
		}
		return _gaa, nil
	case *_eg.PdfFieldButton:
		_eged := _fb
		if _eged.IsCheckbox() {
			_gd, _ade := _ffe(wa, _eged, form.DR, _cdd.Style())
			if _ade != nil {
				return nil, _ade
			}
			return _gd, nil
		}
		_g.Log.Debug("\u0054\u004f\u0044\u004f\u003a\u0020\u0055\u004e\u0048\u0041\u004e\u0044\u004c\u0045\u0044 \u0062u\u0074\u0074\u006f\u006e\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u002b\u0076", _eged.GetType())
	case *_eg.PdfFieldChoice:
		_cdcc := _fb
		switch {
		case _cdcc.Flags().Has(_eg.FieldFlagCombo):
			_bag, _bba := _cgffg(form, wa, _cdcc, _cdd.Style())
			if _bba != nil {
				return nil, _bba
			}
			return _bag, nil
		default:
			_g.Log.Debug("\u0054\u004f\u0044\u004f\u003a\u0020\u0055N\u0048\u0041\u004eD\u004c\u0045\u0044\u0020c\u0068\u006f\u0069\u0063\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0077\u0069\u0074\u0068\u0020\u0066\u006c\u0061\u0067\u0073\u003a\u0020\u0025\u0073", _cdcc.Flags().String())
		}
	default:
		_g.Log.Debug("\u0054\u004f\u0044\u004f\u003a\u0020\u0055\u004e\u0048\u0041N\u0044\u004c\u0045\u0044\u0020\u0066\u0069e\u006c\u0064\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054", _fb)
	}
	return nil, nil
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

// WrapContentStream ensures that the entire content stream for a `page` is wrapped within q ... Q operands.
// Ensures that following operands that are added are not affected by additional operands that are added.
// Implements interface model.ContentStreamWrapper.
func (_edbg FieldAppearance) WrapContentStream(page *_eg.PdfPage) error {
	_gbd, _aef := page.GetAllContentStreams()
	if _aef != nil {
		return _aef
	}
	_dcg := _c.NewContentStreamParser(_gbd)
	_gfb, _aef := _dcg.Parse()
	if _aef != nil {
		return _aef
	}
	_gfb.WrapIfNeeded()
	_bffe := []string{_gfb.String()}
	return page.SetContentStreams(_bffe, _ddba())
}

// LineAnnotationDef defines a line between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none
// (regular line), or arrows at either end.  The line also has a specified width, color and opacity.
type LineAnnotationDef struct {
	X1               float64
	Y1               float64
	X2               float64
	Y2               float64
	LineColor        *_eg.PdfColorDeviceRGB
	Opacity          float64
	LineWidth        float64
	LineEndingStyle1 _efb.LineEndingStyle
	LineEndingStyle2 _efb.LineEndingStyle
}

func (_begd *AppearanceStyle) applyAppearanceCharacteristics(_ffca *_gac.PdfObjectDictionary, _fad *_gac.PdfObjectDictionary, _ecbc *_eg.PdfFont) error {
	if !_begd.AllowMK {
		return nil
	}
	if CA, _faf := _gac.GetString(_ffca.Get("\u0043\u0041")); _faf && _ecbc != nil {
		_cdef := CA.Bytes()
		if len(_cdef) != 0 {
			_gegef := []rune(_ecbc.Encoder().Decode(_cdef))
			if len(_gegef) == 1 {
				_begd.CheckmarkRune = _gegef[0]
			}
		}
	}
	if BC, _aec := _gac.GetArray(_ffca.Get("\u0042\u0043")); _aec {
		_fddb, _dcec := BC.ToFloat64Array()
		if _dcec != nil {
			return _dcec
		}
		switch len(_fddb) {
		case 1:
			_begd.BorderColor = _eg.NewPdfColorDeviceGray(_fddb[0])
		case 3:
			_begd.BorderColor = _eg.NewPdfColorDeviceRGB(_fddb[0], _fddb[1], _fddb[2])
		case 4:
			_begd.BorderColor = _eg.NewPdfColorDeviceCMYK(_fddb[0], _fddb[1], _fddb[2], _fddb[3])
		default:
			_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0042\u0043\u0020\u002d\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0028\u0025\u0064)", len(_fddb))
		}
		if _fad != nil {
			if _cgbf, _acb := _gac.GetNumberAsFloat(_fad.Get("\u0057")); _acb == nil {
				_begd.BorderSize = _cgbf
			}
		}
	}
	if BG, _gacab := _gac.GetArray(_ffca.Get("\u0042\u0047")); _gacab {
		_babg, _edbf := BG.ToFloat64Array()
		if _edbf != nil {
			return _edbf
		}
		switch len(_babg) {
		case 1:
			_begd.FillColor = _eg.NewPdfColorDeviceGray(_babg[0])
		case 3:
			_begd.FillColor = _eg.NewPdfColorDeviceRGB(_babg[0], _babg[1], _babg[2])
		case 4:
			_begd.FillColor = _eg.NewPdfColorDeviceCMYK(_babg[0], _babg[1], _babg[2], _babg[3])
		default:
			_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0042\u0047\u0020\u002d\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0028\u0025\u0064)", len(_babg))
		}
	}
	return nil
}
func _cbd(_bfbe, _ccca float64, _egbb *_eg.Image, _eee AppearanceStyle) (*_eg.XObjectForm, error) {
	_edce, _fbgc := _eg.NewXObjectImageFromImage(_egbb, nil, _gac.NewFlateEncoder())
	if _fbgc != nil {
		return nil, _fbgc
	}
	_edce.Decode = _gac.MakeArrayFromFloats([]float64{0.0, 1.0, 0.0, 1.0, 0.0, 1.0})
	_efge := _eg.NewPdfPageResources()
	_efge.ProcSet = _gac.MakeArray(_gac.MakeName("\u0050\u0044\u0046"), _gac.MakeName("\u0049\u006d\u0061\u0067\u0065\u0043"))
	_efge.SetXObjectImageByName(_gac.PdfObjectName("\u0049\u006d\u0030"), _edce)
	_fgg := _c.NewContentCreator()
	_fgg.Add_q()
	_fgg.Add_cm(float64(_egbb.Width), 0, 0, float64(_egbb.Height), 0, 0)
	_fgg.Add_Do("\u0049\u006d\u0030")
	_fgg.Add_Q()
	_acc := _eg.NewXObjectForm()
	_acc.FormType = _gac.MakeInteger(1)
	_acc.BBox = _gac.MakeArrayFromFloats([]float64{0, 0, float64(_egbb.Width), float64(_egbb.Height)})
	_acc.Resources = _efge
	_acc.SetContentStream(_fgg.Bytes(), _ddba())
	return _acc, nil
}
func (_ceae *AppearanceStyle) processDA(_aba *_eg.PdfField, _fcc *_c.ContentStreamOperations, _bcgd, _agfg *_eg.PdfPageResources, _fefb *_c.ContentCreator) (*AppearanceFont, bool, error) {
	var _gcg *AppearanceFont
	var _fdde bool
	if _ceae.Fonts != nil {
		if _ceae.Fonts.Fallback != nil {
			_gcg = _ceae.Fonts.Fallback
		}
		if _eab := _ceae.Fonts.FieldFallbacks; _eab != nil {
			if _eadcb, _daca := _eab[_aba.PartialName()]; _daca {
				_gcg = _eadcb
			} else if _fbee, _bfe := _aba.FullName(); _bfe == nil {
				if _cef, _gbeg := _eab[_fbee]; _gbeg {
					_gcg = _cef
				}
			}
		}
		_fdde = _ceae.Fonts.ForceReplace
	}
	var _gebc string
	var _bbb float64
	var _dcf bool
	if _fcc != nil {
		for _, _cfc := range *_fcc {
			if _cfc.Operand == "\u0054\u0066" && len(_cfc.Params) == 2 {
				if _dcfe, _cgdb := _gac.GetNameVal(_cfc.Params[0]); _cgdb {
					_gebc = _dcfe
				}
				if _gbb, _ffec := _gac.GetNumberAsFloat(_cfc.Params[1]); _ffec == nil {
					_bbb = _gbb
				}
				_dcf = true
				continue
			}
			_fefb.AddOperand(*_cfc)
		}
	}
	var _ggb *AppearanceFont
	var _dca _gac.PdfObject
	if _fdde && _gcg != nil {
		_ggb = _gcg
	} else {
		if _bcgd != nil && _gebc != "" {
			if _bfbg, _fdgef := _bcgd.GetFontByName(*_gac.MakeName(_gebc)); _fdgef {
				if _fda, _aee := _eg.NewPdfFontFromPdfObject(_bfbg); _aee == nil {
					_dca = _bfbg
					_ggb = &AppearanceFont{Name: _gebc, Font: _fda, Size: _bbb}
				} else {
					_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052:\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006fa\u0064\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0066\u006f\u006e\u0074\u003a\u0020\u0025\u0076", _aee)
				}
			}
		}
		if _ggb == nil && _gcg != nil {
			_ggb = _gcg
		}
		if _ggb == nil {
			_gccd, _eacb := _eg.NewStandard14Font("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a")
			if _eacb != nil {
				return nil, false, _eacb
			}
			_ggb = &AppearanceFont{Name: "\u0048\u0065\u006c\u0076", Font: _gccd, Size: _bbb}
		}
	}
	if _ggb.Size <= 0 && _ceae.Fonts != nil && _ceae.Fonts.FallbackSize > 0 {
		_ggb.Size = _ceae.Fonts.FallbackSize
	}
	_acba := *_gac.MakeName(_ggb.Name)
	if _dca == nil {
		_dca = _ggb.Font.ToPdfObject()
	}
	if _bcgd != nil && !_bcgd.HasFontByName(_acba) {
		_bcgd.SetFontByName(_acba, _dca)
	}
	if _agfg != nil && !_agfg.HasFontByName(_acba) {
		_agfg.SetFontByName(_acba, _dca)
	}
	return _ggb, _dcf, nil
}

// CreateRectangleAnnotation creates a rectangle annotation object that can be added to page PDF annotations.
func CreateRectangleAnnotation(rectDef RectangleAnnotationDef) (*_eg.PdfAnnotation, error) {
	_abeb := _eg.NewPdfAnnotationSquare()
	if rectDef.BorderEnabled {
		_dfc, _affg, _cbge := rectDef.BorderColor.R(), rectDef.BorderColor.G(), rectDef.BorderColor.B()
		_abeb.C = _gac.MakeArrayFromFloats([]float64{_dfc, _affg, _cbge})
		_fbcg := _eg.NewBorderStyle()
		_fbcg.SetBorderWidth(rectDef.BorderWidth)
		_abeb.BS = _fbcg.ToPdfObject()
	}
	if rectDef.FillEnabled {
		_cgfc, _ccf, _bfed := rectDef.FillColor.R(), rectDef.FillColor.G(), rectDef.FillColor.B()
		_abeb.IC = _gac.MakeArrayFromFloats([]float64{_cgfc, _ccf, _bfed})
	} else {
		_abeb.IC = _gac.MakeArrayFromIntegers([]int{})
	}
	if rectDef.Opacity < 1.0 {
		_abeb.CA = _gac.MakeFloat(rectDef.Opacity)
	}
	_bage, _eedg, _abec := _dbef(rectDef)
	if _abec != nil {
		return nil, _abec
	}
	_abeb.AP = _bage
	_abeb.Rect = _gac.MakeArrayFromFloats([]float64{_eedg.Llx, _eedg.Lly, _eedg.Urx, _eedg.Ury})
	return _abeb.PdfAnnotation, nil
}

// AppearanceStyle defines style parameters for appearance stream generation.
type AppearanceStyle struct {

	// How much of Rect height to fill when autosizing text.
	AutoFontSizeFraction float64

	// CheckmarkRune is a rune used for check mark in checkboxes (for ZapfDingbats font).
	CheckmarkRune rune
	BorderSize    float64
	BorderColor   _eg.PdfColor
	FillColor     _eg.PdfColor

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

// NewTextField generates a new text field with partial name `name` at location
// specified by `rect` on given `page` and with field specific options `opt`.
func NewTextField(page *_eg.PdfPage, name string, rect []float64, opt TextFieldOptions) (*_eg.PdfFieldText, error) {
	if page == nil {
		return nil, _bb.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _bb.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _bb.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_caca := _eg.NewPdfField()
	_gcga := &_eg.PdfFieldText{}
	_caca.SetContext(_gcga)
	_gcga.PdfField = _caca
	_gcga.T = _gac.MakeString(name)
	if opt.MaxLen > 0 {
		_gcga.MaxLen = _gac.MakeInteger(int64(opt.MaxLen))
	}
	if len(opt.Value) > 0 {
		_gcga.V = _gac.MakeString(opt.Value)
	}
	_agbf := _eg.NewPdfAnnotationWidget()
	_agbf.Rect = _gac.MakeArrayFromFloats(rect)
	_agbf.P = page.ToPdfObject()
	_agbf.F = _gac.MakeInteger(4)
	_agbf.Parent = _gcga.ToPdfObject()
	_gcga.Annotations = append(_gcga.Annotations, _agbf)
	return _gcga, nil
}
func _dbef(_fdc RectangleAnnotationDef) (*_gac.PdfObjectDictionary, *_eg.PdfRectangle, error) {
	_faea := _eg.NewXObjectForm()
	_faea.Resources = _eg.NewPdfPageResources()
	_fed := ""
	if _fdc.Opacity < 1.0 {
		_ffd := _gac.MakeDict()
		_ffd.Set("\u0063\u0061", _gac.MakeFloat(_fdc.Opacity))
		_ffd.Set("\u0043\u0041", _gac.MakeFloat(_fdc.Opacity))
		_gcb := _faea.Resources.AddExtGState("\u0067\u0073\u0031", _ffd)
		if _gcb != nil {
			_g.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020\u0065\u0078\u0074g\u0073\u0074\u0061t\u0065 \u0067\u0073\u0031")
			return nil, nil, _gcb
		}
		_fed = "\u0067\u0073\u0031"
	}
	_ffad, _fagd, _cccf, _bgag := _aeec(_fdc, _fed)
	if _bgag != nil {
		return nil, nil, _bgag
	}
	_bgag = _faea.SetContentStream(_ffad, nil)
	if _bgag != nil {
		return nil, nil, _bgag
	}
	_faea.BBox = _fagd.ToPdfObject()
	_daf := _gac.MakeDict()
	_daf.Set("\u004e", _faea.ToPdfObject())
	return _daf, _cccf, nil
}

// GenerateAppearanceDict generates an appearance dictionary for widget annotation `wa` for the `field` in `form`.
// Implements interface model.FieldAppearanceGenerator.
func (_deed ImageFieldAppearance) GenerateAppearanceDict(form *_eg.PdfAcroForm, field *_eg.PdfField, wa *_eg.PdfAnnotationWidget) (*_gac.PdfObjectDictionary, error) {
	_, _edc := field.GetContext().(*_eg.PdfFieldButton)
	if !_edc {
		_g.Log.Trace("C\u006f\u0075\u006c\u0064\u0020\u006fn\u006c\u0079\u0020\u0068\u0061\u006ed\u006c\u0065\u0020\u0062\u0075\u0074\u0074o\u006e\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067")
		return nil, nil
	}
	_bdgb, _dgg := _gac.GetDict(wa.AP)
	if _dgg && _deed.OnlyIfMissing {
		_g.Log.Trace("\u0041\u006c\u0072\u0065a\u0064\u0079\u0020\u0070\u006f\u0070\u0075\u006c\u0061\u0074e\u0064 \u002d\u0020\u0069\u0067\u006e\u006f\u0072i\u006e\u0067")
		return _bdgb, nil
	}
	if form.DR == nil {
		form.DR = _eg.NewPdfPageResources()
	}
	switch _gegeb := field.GetContext().(type) {
	case *_eg.PdfFieldButton:
		if _gegeb.IsPush() {
			_cdab, _fcec := _gfd(_gegeb, wa, _deed.Style())
			if _fcec != nil {
				return nil, _fcec
			}
			return _cdab, nil
		}
	}
	return nil, nil
}
func _efc(_fdd *_c.ContentCreator, _cdbf AppearanceStyle, _fbg, _dfg float64) {
	_fdd.Add_q().Add_re(0, 0, _fbg, _dfg).Add_w(_cdbf.BorderSize).SetStrokingColor(_cdbf.BorderColor).SetNonStrokingColor(_cdbf.FillColor).Add_B().Add_Q()
}

// NewImageField generates a new image field with partial name `name` at location `rect`
// on specified `page` and with field specific options `opt`.
func NewImageField(page *_eg.PdfPage, name string, rect []float64, opt ImageFieldOptions) (*_eg.PdfFieldButton, error) {
	if page == nil {
		return nil, _bb.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _bb.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _bb.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_egbg := _eg.NewPdfField()
	_abcf := &_eg.PdfFieldButton{}
	_abcf.PdfField = _egbg
	_egbg.SetContext(_abcf)
	_abcf.SetType(_eg.ButtonTypePush)
	_abcf.T = _gac.MakeString(name)
	_fdf := _eg.NewPdfAnnotationWidget()
	_fdf.Rect = _gac.MakeArrayFromFloats(rect)
	_fdf.P = page.ToPdfObject()
	_fdf.F = _gac.MakeInteger(4)
	_fdf.Parent = _abcf.ToPdfObject()
	_dcbc := rect[2] - rect[0]
	_bcd := rect[3] - rect[1]
	_bcbf := opt._bfgc
	_dacd := _c.NewContentCreator()
	if _bcbf.BorderSize > 0 {
		_efc(_dacd, _bcbf, _dcbc, _bcd)
	}
	if _bcbf.DrawAlignmentReticle {
		_aeef := _bcbf
		_aeef.BorderSize = 0.2
		_gdgd(_dacd, _aeef, _dcbc, _bcd)
	}
	_bafb, _gadc := _cbd(_dcbc, _bcd, opt.Image, _bcbf)
	if _gadc != nil {
		return nil, _gadc
	}
	_dade, _gabc := _gac.GetDict(_fdf.MK)
	if _gabc {
		_dade.Set("\u006c", _bafb.ToPdfObject())
	}
	_cbg := _gac.MakeDict()
	_cbg.Set("\u0046\u0052\u004d", _bafb.ToPdfObject())
	_abaa := _eg.NewPdfPageResources()
	_abaa.ProcSet = _gac.MakeArray(_gac.MakeName("\u0050\u0044\u0046"))
	_abaa.XObject = _cbg
	_bcdg := _dcbc - 2
	_cgfg := _bcd - 2
	_dacd.Add_q()
	_dacd.Add_re(1, 1, _bcdg, _cgfg)
	_dacd.Add_W()
	_dacd.Add_n()
	_bcdg -= 2
	_cgfg -= 2
	_dacd.Add_q()
	_dacd.Add_re(2, 2, _bcdg, _cgfg)
	_dacd.Add_W()
	_dacd.Add_n()
	_cggc := _d.Min(_bcdg/float64(opt.Image.Width), _cgfg/float64(opt.Image.Height))
	_dacd.Add_cm(_cggc, 0, 0, _cggc, (_dcbc/2)-(float64(opt.Image.Width)*_cggc/2)+2, 2)
	_dacd.Add_Do("\u0046\u0052\u004d")
	_dacd.Add_Q()
	_dacd.Add_Q()
	_gfbe := _eg.NewXObjectForm()
	_gfbe.FormType = _gac.MakeInteger(1)
	_gfbe.Resources = _abaa
	_gfbe.BBox = _gac.MakeArrayFromFloats([]float64{0, 0, _dcbc, _bcd})
	_gfbe.Matrix = _gac.MakeArrayFromFloats([]float64{1.0, 0.0, 0.0, 1.0, 0.0, 0.0})
	_gfbe.SetContentStream(_dacd.Bytes(), _ddba())
	_dcgf := _gac.MakeDict()
	_dcgf.Set("\u004e", _gfbe.ToPdfObject())
	_fdf.AP = _dcgf
	_abcf.Annotations = append(_abcf.Annotations, _fdf)
	return _abcf, nil
}

// Style returns the appearance style of `fa`. If not specified, returns default style.
func (_aada FieldAppearance) Style() AppearanceStyle {
	if _aada._dfe != nil {
		return *_aada._dfe
	}
	_ab := _cf
	return AppearanceStyle{AutoFontSizeFraction: 0.65, CheckmarkRune: '✔', BorderSize: 0.0, BorderColor: _eg.NewPdfColorDeviceGray(0), FillColor: _eg.NewPdfColorDeviceGray(1), MultilineLineHeight: 1.2, MultilineVAlignMiddle: false, DrawAlignmentReticle: false, AllowMK: true, MarginLeft: &_ab}
}

// ComboboxFieldOptions defines optional parameters for a combobox form field.
type ComboboxFieldOptions struct {

	// Choices is the list of string values that can be selected.
	Choices []string
}

// ImageFieldAppearance implements interface model.FieldAppearanceGenerator and generates appearance streams
// for attaching an image to a button field.
type ImageFieldAppearance struct {
	OnlyIfMissing bool
	_gfea         *AppearanceStyle
}

func _aeg(_afb *_eg.PdfField, _edgc, _edfb float64, _geb string, _dcd AppearanceStyle, _eggg *_c.ContentStreamOperations, _gdc *_eg.PdfPageResources, _cacd *_gac.PdfObjectDictionary) (*_eg.XObjectForm, error) {
	_gfg := _eg.NewPdfPageResources()
	_ddae, _bfg := _edgc, _edfb
	_adec := _c.NewContentCreator()
	if _dcd.BorderSize > 0 {
		_efc(_adec, _dcd, _edgc, _edfb)
	}
	if _dcd.DrawAlignmentReticle {
		_gfe := _dcd
		_gfe.BorderSize = 0.2
		_gdgd(_adec, _gfe, _edgc, _edfb)
	}
	_adec.Add_BMC("\u0054\u0078")
	_adec.Add_q()
	_adec.Add_BT()
	_edgc, _edfb = _dcd.applyRotation(_cacd, _edgc, _edfb, _adec)
	_cee, _cae, _eed := _dcd.processDA(_afb, _eggg, _gdc, _gfg, _adec)
	if _eed != nil {
		return nil, _eed
	}
	_addb := _cee.Font
	_dbed := _cee.Size
	_gdb := _gac.MakeName(_cee.Name)
	_ebg := _dbed == 0
	if _ebg && _cae {
		_dbed = _edfb * _dcd.AutoFontSizeFraction
	}
	_deea := _addb.Encoder()
	if _deea == nil {
		_g.Log.Debug("\u0057\u0041RN\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0069\u0073\u0020\u006e\u0069l\u002e\u0020\u0041\u0073s\u0075\u006d\u0069\u006eg \u0069\u0064e\u006et\u0069\u0074\u0079\u0020\u0065\u006ec\u006f\u0064\u0065r\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
		_deea = _b.NewIdentityTextEncoder("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048")
	}
	if len(_geb) == 0 {
		return nil, nil
	}
	_fea := _cf
	if _dcd.MarginLeft != nil {
		_fea = *_dcd.MarginLeft
	}
	_aadb := 0.0
	if _deea != nil {
		for _, _cga := range _geb {
			_efa, _gag := _addb.GetRuneMetrics(_cga)
			if !_gag {
				_g.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u006d\u0065\u0074r\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u002d\u0020\u0073k\u0069p\u0070\u0069\u006e\u0067", _cga)
				continue
			}
			_aadb += _efa.Wx
		}
		_geb = string(_deea.Encode(_geb))
	}
	if _dbed == 0 || _ebg && _aadb > 0 && _fea+_aadb*_dbed/1000.0 > _edgc {
		_dbed = 0.95 * 1000.0 * (_edgc - _fea) / _aadb
	}
	_aggg := 1.0 * _dbed
	_bdf := 2.0
	{
		_bdbc := _aggg
		if _ebg && _bdf+_bdbc > _edfb {
			_dbed = 0.95 * (_edfb - _bdf)
			_aggg = 1.0 * _dbed
			_bdbc = _aggg
		}
		if _edfb > _bdbc {
			_bdf = (_edfb - _bdbc) / 2.0
			_bdf += 1.50
		}
	}
	_adec.Add_Tf(*_gdb, _dbed)
	_adec.Add_Td(_fea, _bdf)
	_adec.Add_Tj(*_gac.MakeString(_geb))
	_adec.Add_ET()
	_adec.Add_Q()
	_adec.Add_EMC()
	_caa := _eg.NewXObjectForm()
	_caa.Resources = _gfg
	_caa.BBox = _gac.MakeArrayFromFloats([]float64{0, 0, _ddae, _bfg})
	_caa.SetContentStream(_adec.Bytes(), _ddba())
	return _caa, nil
}
func _ddba() _gac.StreamEncoder { return _gac.NewFlateEncoder() }

const (
	SignatureImageLeft SignatureImagePosition = iota
	SignatureImageRight
	SignatureImageTop
	SignatureImageBottom
)

// CircleAnnotationDef defines a circle annotation or ellipse at position (X, Y) and Width and Height.
// The annotation has various style parameters including Fill and Border options and Opacity.
type CircleAnnotationDef struct {
	X             float64
	Y             float64
	Width         float64
	Height        float64
	FillEnabled   bool
	FillColor     *_eg.PdfColorDeviceRGB
	BorderEnabled bool
	BorderWidth   float64
	BorderColor   *_eg.PdfColorDeviceRGB
	Opacity       float64
}

// TextFieldOptions defines optional parameter for a text field in a form.
type TextFieldOptions struct {
	MaxLen int
	Value  string
}

func _afge(_aga LineAnnotationDef) (*_gac.PdfObjectDictionary, *_eg.PdfRectangle, error) {
	_gagg := _eg.NewXObjectForm()
	_gagg.Resources = _eg.NewPdfPageResources()
	_dbae := ""
	if _aga.Opacity < 1.0 {
		_dgce := _gac.MakeDict()
		_dgce.Set("\u0063\u0061", _gac.MakeFloat(_aga.Opacity))
		_gdbe := _gagg.Resources.AddExtGState("\u0067\u0073\u0031", _dgce)
		if _gdbe != nil {
			_g.Log.Debug("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020\u0065\u0078\u0074g\u0073\u0074\u0061t\u0065 \u0067\u0073\u0031")
			return nil, nil, _gdbe
		}
		_dbae = "\u0067\u0073\u0031"
	}
	_bfbgb, _ffg, _bfec, _bedd := _cdfe(_aga, _dbae)
	if _bedd != nil {
		return nil, nil, _bedd
	}
	_bedd = _gagg.SetContentStream(_bfbgb, nil)
	if _bedd != nil {
		return nil, nil, _bedd
	}
	_gagg.BBox = _ffg.ToPdfObject()
	_gcdg := _gac.MakeDict()
	_gcdg.Set("\u004e", _gagg.ToPdfObject())
	return _gcdg, _bfec, nil
}

// SetStyle applies appearance `style` to `fa`.
func (_edb *FieldAppearance) SetStyle(style AppearanceStyle) { _edb._dfe = &style }
func _bgg(_fbe *_eg.PdfField) string {
	if _fbe == nil {
		return ""
	}
	_cgb, _efgg := _fbe.GetContext().(*_eg.PdfFieldText)
	if !_efgg {
		return _bgg(_fbe.Parent)
	}
	if _cgb.DA != nil {
		return _cgb.DA.Str()
	}
	return _bgg(_cgb.Parent)
}
func _cgffg(_bcab *_eg.PdfAcroForm, _cdb *_eg.PdfAnnotationWidget, _ecdb *_eg.PdfFieldChoice, _bbfb AppearanceStyle) (*_gac.PdfObjectDictionary, error) {
	_dbc, _edae := _gac.GetArray(_cdb.Rect)
	if !_edae {
		return nil, _bb.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_ecg, _edgd := _eg.NewPdfRectangle(*_dbc)
	if _edgd != nil {
		return nil, _edgd
	}
	_cbe, _adb := _ecg.Width(), _ecg.Height()
	_g.Log.Debug("\u0043\u0068\u006f\u0069\u0063\u0065\u002c\u0020\u0077\u0061\u0020\u0042S\u003a\u0020\u0025\u0076", _cdb.BS)
	_gcc, _edgd := _c.NewContentStreamParser(_bgg(_ecdb.PdfField)).Parse()
	if _edgd != nil {
		return nil, _edgd
	}
	_gbe, _dcb := _gac.GetDict(_cdb.MK)
	if _dcb {
		_dbf, _ := _gac.GetDict(_cdb.BS)
		_dfef := _bbfb.applyAppearanceCharacteristics(_gbe, _dbf, nil)
		if _dfef != nil {
			return nil, _dfef
		}
	}
	_gcde := _gac.MakeDict()
	for _, _ddaf := range _ecdb.Opt.Elements() {
		if _eaf, _cda := _gac.GetArray(_ddaf); _cda && _eaf.Len() == 2 {
			_ddaf = _eaf.Get(1)
		}
		var _cea string
		if _bagd, _afd := _gac.GetString(_ddaf); _afd {
			_cea = _bagd.Decoded()
		} else if _aabeb, _bdg := _gac.GetName(_ddaf); _bdg {
			_cea = _aabeb.String()
		} else {
			_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u004f\u0070\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006de\u002f\u0073\u0074\u0072\u0069\u006e\u0067 \u002d\u0020\u0025\u0054", _ddaf)
			return nil, _bb.New("\u006e\u006f\u0074\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u002f\u0073t\u0072\u0069\u006e\u0067")
		}
		if len(_cea) > 0 {
			_bcfe, _cbea := _aeg(_ecdb.PdfField, _cbe, _adb, _cea, _bbfb, _gcc, _bcab.DR, _gbe)
			if _cbea != nil {
				return nil, _cbea
			}
			_gcde.Set(*_gac.MakeName(_cea), _bcfe.ToPdfObject())
		}
	}
	_dbeb := _gac.MakeDict()
	_dbeb.Set("\u004e", _gcde)
	return _dbeb, nil
}

// NewCheckboxField generates a new checkbox field with partial name `name` at location `rect`
// on specified `page` and with field specific options `opt`.
func NewCheckboxField(page *_eg.PdfPage, name string, rect []float64, opt CheckboxFieldOptions) (*_eg.PdfFieldButton, error) {
	if page == nil {
		return nil, _bb.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _bb.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _bb.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_dded, _ggg := _eg.NewStandard14Font(_eg.ZapfDingbatsName)
	if _ggg != nil {
		return nil, _ggg
	}
	_fade := _eg.NewPdfField()
	_dgff := &_eg.PdfFieldButton{}
	_fade.SetContext(_dgff)
	_dgff.PdfField = _fade
	_dgff.T = _gac.MakeString(name)
	_dgff.SetType(_eg.ButtonTypeCheckbox)
	_bac := "\u004f\u0066\u0066"
	if opt.Checked {
		_bac = "\u0059\u0065\u0073"
	}
	_dgff.V = _gac.MakeName(_bac)
	_gace := _eg.NewPdfAnnotationWidget()
	_gace.Rect = _gac.MakeArrayFromFloats(rect)
	_gace.P = page.ToPdfObject()
	_gace.F = _gac.MakeInteger(4)
	_gace.Parent = _dgff.ToPdfObject()
	_fafa := rect[2] - rect[0]
	_gbef := rect[3] - rect[1]
	var _bbcde _ee.Buffer
	_bbcde.WriteString("\u0071\u000a")
	_bbcde.WriteString("\u0030 \u0030\u0020\u0031\u0020\u0072\u0067\n")
	_bbcde.WriteString("\u0042\u0054\u000a")
	_bbcde.WriteString("\u002f\u005a\u0061D\u0062\u0020\u0031\u0032\u0020\u0054\u0066\u000a")
	_bbcde.WriteString("\u0045\u0054\u000a")
	_bbcde.WriteString("\u0051\u000a")
	_gbbf := _c.NewContentCreator()
	_gbbf.Add_q()
	_gbbf.Add_rg(0, 0, 1)
	_gbbf.Add_BT()
	_gbbf.Add_Tf(*_gac.MakeName("\u005a\u0061\u0044\u0062"), 12)
	_gbbf.Add_Td(0, 0)
	_gbbf.Add_ET()
	_gbbf.Add_Q()
	_dfdda := _eg.NewXObjectForm()
	_dfdda.SetContentStream(_gbbf.Bytes(), _gac.NewRawEncoder())
	_dfdda.BBox = _gac.MakeArrayFromFloats([]float64{0, 0, _fafa, _gbef})
	_dfdda.Resources = _eg.NewPdfPageResources()
	_dfdda.Resources.SetFontByName("\u005a\u0061\u0044\u0062", _dded.ToPdfObject())
	_gbbf = _c.NewContentCreator()
	_gbbf.Add_q()
	_gbbf.Add_re(0, 0, _fafa, _gbef)
	_gbbf.Add_W().Add_n()
	_gbbf.Add_rg(0, 0, 1)
	_gbbf.Translate(0, 3.0)
	_gbbf.Add_BT()
	_gbbf.Add_Tf(*_gac.MakeName("\u005a\u0061\u0044\u0062"), 12)
	_gbbf.Add_Td(0, 0)
	_gbbf.Add_Tj(*_gac.MakeString("\u0034"))
	_gbbf.Add_ET()
	_gbbf.Add_Q()
	_dba := _eg.NewXObjectForm()
	_dba.SetContentStream(_gbbf.Bytes(), _gac.NewRawEncoder())
	_dba.BBox = _gac.MakeArrayFromFloats([]float64{0, 0, _fafa, _gbef})
	_dba.Resources = _eg.NewPdfPageResources()
	_dba.Resources.SetFontByName("\u005a\u0061\u0044\u0062", _dded.ToPdfObject())
	_cdcgf := _gac.MakeDict()
	_cdcgf.Set("\u004f\u0066\u0066", _dfdda.ToPdfObject())
	_cdcgf.Set("\u0059\u0065\u0073", _dba.ToPdfObject())
	_gba := _gac.MakeDict()
	_gba.Set("\u004e", _cdcgf)
	_gace.AP = _gba
	_gace.AS = _gac.MakeName(_bac)
	_dgff.Annotations = append(_dgff.Annotations, _gace)
	return _dgff, nil
}

// SignatureImagePosition specifies the image signature location relative to the text signature.
// If text signature is not defined, this position will be ignored.
type SignatureImagePosition int

func _eacg(_beeg _ga.Image, _bfa string, _feg *SignatureFieldOpts, _dccd []float64, _bbce *_c.ContentCreator) (*_gac.PdfObjectName, *_eg.XObjectImage, error) {
	_cdaf, _gde := _eg.DefaultImageHandler{}.NewImageFromGoImage(_beeg)
	if _gde != nil {
		return nil, nil, _gde
	}
	_cacfg, _gde := _eg.NewXObjectImageFromImage(_cdaf, nil, _feg.Encoder)
	if _gde != nil {
		return nil, nil, _gde
	}
	_bbbe, _edgb := float64(*_cacfg.Width), float64(*_cacfg.Height)
	_adab := _dccd[2] - _dccd[0]
	_ccec := _dccd[3] - _dccd[1]
	if _feg.AutoSize {
		_fead := _d.Min(_adab/_bbbe, _ccec/_edgb)
		_bbbe *= _fead
		_edgb *= _fead
		_dccd[0] = _dccd[0] + (_adab / 2) - (_bbbe / 2)
		_dccd[1] = _dccd[1] + (_ccec / 2) - (_edgb / 2)
	}
	var _egba *_gac.PdfObjectName
	if _bfcd, _cede := _gac.GetName(_cacfg.Name); _cede {
		_egba = _bfcd
	} else {
		_egba = _gac.MakeName(_bfa)
	}
	if _bbce != nil {
		_bbce.Add_q().Translate(_dccd[0], _dccd[1]).Scale(_bbbe, _edgb).Add_Do(*_egba).Add_Q()
	} else {
		return nil, nil, _bb.New("\u0043\u006f\u006e\u0074en\u0074\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u0075l\u006c")
	}
	return _egba, _cacfg, nil
}

// NewComboboxField generates a new combobox form field with partial name `name` at location `rect`
// on specified `page` and with field specific options `opt`.
func NewComboboxField(page *_eg.PdfPage, name string, rect []float64, opt ComboboxFieldOptions) (*_eg.PdfFieldChoice, error) {
	if page == nil {
		return nil, _bb.New("\u0070a\u0067e\u0020\u006e\u006f\u0074\u0020s\u0070\u0065c\u0069\u0066\u0069\u0065\u0064")
	}
	if len(name) <= 0 {
		return nil, _bb.New("\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0061\u0074\u0074\u0072\u0069\u0062u\u0074e\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065\u0064")
	}
	if len(rect) != 4 {
		return nil, _bb.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_ddbe := _eg.NewPdfField()
	_cbb := &_eg.PdfFieldChoice{}
	_ddbe.SetContext(_cbb)
	_cbb.PdfField = _ddbe
	_cbb.T = _gac.MakeString(name)
	_cbb.Opt = _gac.MakeArray()
	for _, _gee := range opt.Choices {
		_cbb.Opt.Append(_gac.MakeString(_gee))
	}
	_cbb.SetFlag(_eg.FieldFlagCombo)
	_cced := _eg.NewPdfAnnotationWidget()
	_cced.Rect = _gac.MakeArrayFromFloats(rect)
	_cced.P = page.ToPdfObject()
	_cced.F = _gac.MakeInteger(4)
	_cced.Parent = _cbb.ToPdfObject()
	_cbb.Annotations = append(_cbb.Annotations, _cced)
	return _cbb, nil
}
func _cdfe(_ggbc LineAnnotationDef, _cdce string) ([]byte, *_eg.PdfRectangle, *_eg.PdfRectangle, error) {
	_bcabe := _efb.Line{X1: 0, Y1: 0, X2: _ggbc.X2 - _ggbc.X1, Y2: _ggbc.Y2 - _ggbc.Y1, LineColor: _ggbc.LineColor, Opacity: _ggbc.Opacity, LineWidth: _ggbc.LineWidth, LineEndingStyle1: _ggbc.LineEndingStyle1, LineEndingStyle2: _ggbc.LineEndingStyle2}
	_cabf, _afdb, _aca := _bcabe.Draw(_cdce)
	if _aca != nil {
		return nil, nil, nil, _aca
	}
	_fceb := &_eg.PdfRectangle{}
	_fceb.Llx = _ggbc.X1 + _afdb.Llx
	_fceb.Lly = _ggbc.Y1 + _afdb.Lly
	_fceb.Urx = _ggbc.X1 + _afdb.Urx
	_fceb.Ury = _ggbc.Y1 + _afdb.Ury
	return _cabf, _afdb, _fceb, nil
}

// WrapContentStream ensures that the entire content stream for a `page` is wrapped within q ... Q operands.
// Ensures that following operands that are added are not affected by additional operands that are added.
// Implements interface model.ContentStreamWrapper.
func (_ecaa ImageFieldAppearance) WrapContentStream(page *_eg.PdfPage) error {
	_fgdf, _caef := page.GetAllContentStreams()
	if _caef != nil {
		return _caef
	}
	_gbdf := _c.NewContentStreamParser(_fgdf)
	_fgbe, _caef := _gbdf.Parse()
	if _caef != nil {
		return _caef
	}
	_fgbe.WrapIfNeeded()
	_ebf := []string{_fgbe.String()}
	return page.SetContentStreams(_ebf, _ddba())
}

// AppearanceFont represents a font used for generating the appearance of a
// field in the filling/flattening process.
type AppearanceFont struct {

	// Name represents the name of the font which will be added to the
	// AcroForm resources (DR).
	Name string

	// Font represents the actual font used for the field appearance.
	Font *_eg.PdfFont

	// Size represents the size of the font used for the field appearance.
	// If the font size is 0, the value of the FallbackSize field of the
	// AppearanceFontStyle is used, if set. Otherwise, the font size is
	// calculated based on the available annotation height and on the
	// AutoFontSizeFraction field of the AppearanceStyle.
	Size float64
}

func _bggb(_gdd []*SignatureLine, _age *SignatureFieldOpts) (*_gac.PdfObjectDictionary, error) {
	if _age == nil {
		_age = NewSignatureFieldOpts()
	}
	var _bfee error
	var _dec *_gac.PdfObjectName
	_cfa := _age.Font
	if _cfa != nil {
		_ace, _ := _cfa.GetFontDescriptor()
		if _ace != nil {
			if _bdd, _cgdg := _ace.FontName.(*_gac.PdfObjectName); _cgdg {
				_dec = _bdd
			}
		}
		if _dec == nil {
			_dec = _gac.MakeName("\u0046\u006f\u006et\u0031")
		}
	} else {
		if _cfa, _bfee = _eg.NewStandard14Font("\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a"); _bfee != nil {
			return nil, _bfee
		}
		_dec = _gac.MakeName("\u0048\u0065\u006c\u0076")
	}
	_dfdg := _age.FontSize
	if _dfdg <= 0 {
		_dfdg = 10
	}
	if _age.LineHeight <= 0 {
		_age.LineHeight = 1
	}
	_agga := _age.LineHeight * _dfdg
	_cfef, _edac := _cfa.GetRuneMetrics(' ')
	if !_edac {
		return nil, _bb.New("\u0074\u0068e \u0066\u006f\u006et\u0020\u0064\u006f\u0065s n\u006ft \u0068\u0061\u0076\u0065\u0020\u0061\u0020sp\u0061\u0063\u0065\u0020\u0067\u006c\u0079p\u0068")
	}
	_egb := _cfef.Wx
	var _ecbf float64
	var _ebb []string
	for _, _ggbb := range _gdd {
		if _ggbb.Text == "" {
			continue
		}
		_dgfa := _ggbb.Text
		if _ggbb.Desc != "" {
			_dgfa = _ggbb.Desc + "\u003a\u0020" + _dgfa
		}
		_ebb = append(_ebb, _dgfa)
		var _ecbb float64
		for _, _dgc := range _dgfa {
			_gadg, _ddfd := _cfa.GetRuneMetrics(_dgc)
			if !_ddfd {
				continue
			}
			_ecbb += _gadg.Wx
		}
		if _ecbb > _ecbf {
			_ecbf = _ecbb
		}
	}
	_ecbf = _ecbf * _dfdg / 1000.0
	_cddg := float64(len(_ebb)) * _agga
	_bffb := _age.Image != nil
	_bdab := _age.Rect
	if _bdab == nil {
		_bdab = []float64{0, 0, _ecbf, _cddg}
		if _bffb {
			_bdab[2] = _ecbf * 2
			_bdab[3] = _cddg * 2
		}
		_age.Rect = _bdab
	}
	_fbef := _bdab[2] - _bdab[0]
	_dece := _bdab[3] - _bdab[1]
	_fde, _bdfg := _bdab, _bdab
	var _eafe, _ddef float64
	if _bffb && len(_ebb) > 0 {
		if _age.ImagePosition <= SignatureImageRight {
			_egef := []float64{_bdab[0], _bdab[1], _bdab[0] + (_fbef / 2), _bdab[3]}
			_cgag := []float64{_bdab[0] + (_fbef / 2), _bdab[1], _bdab[2], _bdab[3]}
			if _age.ImagePosition == SignatureImageLeft {
				_fde, _bdfg = _egef, _cgag
			} else {
				_fde, _bdfg = _cgag, _egef
			}
		} else {
			_cdba := []float64{_bdab[0], _bdab[1], _bdab[2], _bdab[1] + (_dece / 2)}
			_afa := []float64{_bdab[0], _bdab[1] + (_dece / 2), _bdab[2], _bdab[3]}
			if _age.ImagePosition == SignatureImageTop {
				_fde, _bdfg = _afa, _cdba
			} else {
				_fde, _bdfg = _cdba, _afa
			}
		}
	}
	_eafe = _bdfg[2] - _bdfg[0]
	_ddef = _bdfg[3] - _bdfg[1]
	var _fca float64
	if _age.AutoSize {
		if _ecbf > _eafe || _cddg > _ddef {
			_eag := _d.Min(_eafe/_ecbf, _ddef/_cddg)
			_dfdg *= _eag
		}
		_agga = _age.LineHeight * _dfdg
		_fca += (_ddef - float64(len(_ebb))*_agga) / 2
	}
	_cff := _c.NewContentCreator()
	_afg := _eg.NewPdfPageResources()
	_afg.SetFontByName(*_dec, _cfa.ToPdfObject())
	if _age.BorderSize <= 0 {
		_age.BorderSize = 0
		_age.BorderColor = _eg.NewPdfColorDeviceGray(1)
	}
	if _age.BorderColor == nil {
		_age.BorderColor = _eg.NewPdfColorDeviceGray(1)
	}
	if _age.FillColor == nil {
		_age.FillColor = _eg.NewPdfColorDeviceGray(1)
	}
	_cff.Add_q().SetNonStrokingColor(_age.FillColor).SetStrokingColor(_age.BorderColor).Add_w(_age.BorderSize).Add_re(_bdab[0], _bdab[1], _fbef, _dece).Add_B().Add_Q()
	if _age.WatermarkImage != nil {
		_edga := []float64{_bdab[0], _bdab[1], _bdab[2], _bdab[3]}
		_gef, _cdcg, _fcfb := _eacg(_age.WatermarkImage, "\u0049\u006d\u0061\u0067\u0065\u0057\u0061\u0074\u0065r\u006d\u0061\u0072\u006b", _age, _edga, _cff)
		if _fcfb != nil {
			return nil, _fcfb
		}
		_afg.SetXObjectImageByName(*_gef, _cdcg)
	}
	_cff.Add_q()
	_cff.Translate(_bdfg[0], _bdfg[3]-_agga-_fca)
	_cff.Add_BT()
	_ffea := _cfa.Encoder()
	for _, _fdgf := range _ebb {
		var _aff []byte
		for _, _ggf := range _fdgf {
			if _f.IsSpace(_ggf) {
				if len(_aff) > 0 {
					_cff.SetNonStrokingColor(_age.TextColor).Add_Tf(*_dec, _dfdg).Add_TL(_agga).Add_TJ([]_gac.PdfObject{_gac.MakeStringFromBytes(_aff)}...)
					_aff = nil
				}
				_cff.Add_Tf(*_dec, _dfdg).Add_TL(_agga).Add_TJ([]_gac.PdfObject{_gac.MakeFloat(-_egb)}...)
			} else {
				_aff = append(_aff, _ffea.Encode(string(_ggf))...)
			}
		}
		if len(_aff) > 0 {
			_cff.SetNonStrokingColor(_age.TextColor).Add_Tf(*_dec, _dfdg).Add_TL(_agga).Add_TJ([]_gac.PdfObject{_gac.MakeStringFromBytes(_aff)}...)
		}
		_cff.Add_Td(0, -_agga)
	}
	_cff.Add_ET()
	_cff.Add_Q()
	if _bffb {
		_fgcd, _cgbff, _aeea := _eacg(_age.Image, "\u0049\u006d\u0061\u0067\u0065\u0053\u0069\u0067\u006ea\u0074\u0075\u0072\u0065", _age, _fde, _cff)
		if _aeea != nil {
			return nil, _aeea
		}
		_afg.SetXObjectImageByName(*_fgcd, _cgbff)
	}
	_begc := _eg.NewXObjectForm()
	_begc.Resources = _afg
	_begc.BBox = _gac.MakeArrayFromFloats(_bdab)
	_begc.SetContentStream(_cff.Bytes(), _ddba())
	_aegd := _gac.MakeDict()
	_aegd.Set("\u004e", _begc.ToPdfObject())
	return _aegd, nil
}

// CheckboxFieldOptions defines optional parameters for a checkbox field a form.
type CheckboxFieldOptions struct{ Checked bool }

// NewSignatureLine returns a new signature line displayed as a part of the
// signature field appearance.
func NewSignatureLine(desc, text string) *SignatureLine {
	return &SignatureLine{Desc: desc, Text: text}
}

// NewSignatureField returns a new signature field with a visible appearance
// containing the specified signature lines and styled according to the
// specified options.
func NewSignatureField(signature *_eg.PdfSignature, lines []*SignatureLine, opts *SignatureFieldOpts) (*_eg.PdfFieldSignature, error) {
	if signature == nil {
		return nil, _bb.New("\u0073\u0069\u0067na\u0074\u0075\u0072\u0065\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	_ecge, _abb := _bggb(lines, opts)
	if _abb != nil {
		return nil, _abb
	}
	_eagg := _eg.NewPdfFieldSignature(signature)
	_eagg.Rect = _gac.MakeArrayFromFloats(opts.Rect)
	_eagg.AP = _ecge
	return _eagg, nil
}
func _ffe(_dacg *_eg.PdfAnnotationWidget, _bge *_eg.PdfFieldButton, _cgg *_eg.PdfPageResources, _bfb AppearanceStyle) (*_gac.PdfObjectDictionary, error) {
	_cgff, _beg := _gac.GetArray(_dacg.Rect)
	if !_beg {
		return nil, _bb.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_eccc, _dfd := _eg.NewPdfRectangle(*_cgff)
	if _dfd != nil {
		return nil, _dfd
	}
	_agg, _cfg := _eccc.Width(), _eccc.Height()
	_fdge, _dfbd := _agg, _cfg
	_g.Log.Debug("\u0043\u0068\u0065\u0063kb\u006f\u0078\u002c\u0020\u0077\u0061\u0020\u0042\u0053\u003a\u0020\u0025\u0076", _dacg.BS)
	_bagc, _dfd := _eg.NewStandard14Font("\u005a\u0061\u0070f\u0044\u0069\u006e\u0067\u0062\u0061\u0074\u0073")
	if _dfd != nil {
		return nil, _dfd
	}
	_eefg, _eefcc := _gac.GetDict(_dacg.MK)
	if _eefcc {
		_gege, _ := _gac.GetDict(_dacg.BS)
		_af := _bfb.applyAppearanceCharacteristics(_eefg, _gege, _bagc)
		if _af != nil {
			return nil, _af
		}
	}
	_cccd := _eg.NewXObjectForm()
	{
		_bbcf := _c.NewContentCreator()
		if _bfb.BorderSize > 0 {
			_efc(_bbcf, _bfb, _agg, _cfg)
		}
		if _bfb.DrawAlignmentReticle {
			_fgcg := _bfb
			_fgcg.BorderSize = 0.2
			_gdgd(_bbcf, _fgcg, _agg, _cfg)
		}
		_agg, _cfg = _bfb.applyRotation(_eefg, _agg, _cfg, _bbcf)
		_ecdg := _bfb.AutoFontSizeFraction * _cfg
		_gae, _aaec := _bagc.GetRuneMetrics(_bfb.CheckmarkRune)
		if !_aaec {
			return nil, _bb.New("\u0067l\u0079p\u0068\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
		}
		_bfc := _bagc.Encoder()
		_fcb := _bfc.Encode(string(_bfb.CheckmarkRune))
		_bdc := _gae.Wx * _ecdg / 1000.0
		_bgf := 705.0
		_edd := _bgf / 1000.0 * _ecdg
		_bcee := _cf
		if _bfb.MarginLeft != nil {
			_bcee = *_bfb.MarginLeft
		}
		_bde := 1.0
		if _bdc < _agg {
			_bcee = (_agg - _bdc) / 2.0
		}
		if _edd < _cfg {
			_bde = (_cfg - _edd) / 2.0
		}
		_bbcf.Add_q().Add_g(0).Add_BT().Add_Tf("\u005a\u0061\u0044\u0062", _ecdg).Add_Td(_bcee, _bde).Add_Tj(*_gac.MakeStringFromBytes(_fcb)).Add_ET().Add_Q()
		_cccd.Resources = _eg.NewPdfPageResources()
		_cccd.Resources.SetFontByName("\u005a\u0061\u0044\u0062", _bagc.ToPdfObject())
		_cccd.BBox = _gac.MakeArrayFromFloats([]float64{0, 0, _fdge, _dfbd})
		_cccd.SetContentStream(_bbcf.Bytes(), _ddba())
	}
	_bca := _eg.NewXObjectForm()
	{
		_eade := _c.NewContentCreator()
		if _bfb.BorderSize > 0 {
			_efc(_eade, _bfb, _agg, _cfg)
		}
		_bca.BBox = _gac.MakeArrayFromFloats([]float64{0, 0, _fdge, _dfbd})
		_bca.SetContentStream(_eade.Bytes(), _ddba())
	}
	_dbe := _gac.MakeDict()
	_dbe.Set("\u004f\u0066\u0066", _bca.ToPdfObject())
	_dbe.Set("\u0059\u0065\u0073", _cccd.ToPdfObject())
	_dge := _gac.MakeDict()
	_dge.Set("\u004e", _dbe)
	return _dge, nil
}
func _ece(_cce *_eg.PdfAnnotationWidget, _gab *_eg.PdfFieldText, _ccb *_eg.PdfPageResources, _gdf AppearanceStyle) (*_gac.PdfObjectDictionary, error) {
	_aaa := _eg.NewPdfPageResources()
	_fcf, _gf := _gac.GetArray(_cce.Rect)
	if !_gf {
		return nil, _bb.New("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0052\u0065\u0063\u0074")
	}
	_bdb, _dab := _eg.NewPdfRectangle(*_fcf)
	if _dab != nil {
		return nil, _dab
	}
	_eda, _eaa := _bdb.Width(), _bdb.Height()
	_dff, _add := _eda, _eaa
	_aab, _ggdg := _gac.GetDict(_cce.MK)
	if _ggdg {
		_efd, _ := _gac.GetDict(_cce.BS)
		_cac := _gdf.applyAppearanceCharacteristics(_aab, _efd, nil)
		if _cac != nil {
			return nil, _cac
		}
	}
	_bbc, _dab := _c.NewContentStreamParser(_bgg(_gab.PdfField)).Parse()
	if _dab != nil {
		return nil, _dab
	}
	_cde := _c.NewContentCreator()
	if _gdf.BorderSize > 0 {
		_efc(_cde, _gdf, _eda, _eaa)
	}
	if _gdf.DrawAlignmentReticle {
		_cacf := _gdf
		_cacf.BorderSize = 0.2
		_gdgd(_cde, _cacf, _eda, _eaa)
	}
	_cde.Add_BMC("\u0054\u0078")
	_cde.Add_q()
	_eda, _eaa = _gdf.applyRotation(_aab, _eda, _eaa, _cde)
	_cde.Add_BT()
	_bae, _ead, _dab := _gdf.processDA(_gab.PdfField, _bbc, _ccb, _aaa, _cde)
	if _dab != nil {
		return nil, _dab
	}
	_gga := _bae.Font
	_abc := _bae.Size
	_ffa := _gac.MakeName(_bae.Name)
	_fae := _abc == 0
	if _fae && _ead {
		_abc = _eaa * _gdf.AutoFontSizeFraction
	}
	_bbf := _gga.Encoder()
	if _bbf == nil {
		_g.Log.Debug("\u0057\u0041RN\u003a\u0020\u0066\u006f\u006e\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0069\u0073\u0020\u006e\u0069l\u002e\u0020\u0041\u0073s\u0075\u006d\u0069\u006eg \u0069\u0064e\u006et\u0069\u0074\u0079\u0020\u0065\u006ec\u006f\u0064\u0065r\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
		_bbf = _b.NewIdentityTextEncoder("\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048")
	}
	_ccc, _dab := _gga.GetFontDescriptor()
	if _dab != nil {
		_g.Log.Debug("\u0045\u0072ro\u0072\u003a\u0020U\u006e\u0061\u0062\u006ce t\u006f g\u0065\u0074\u0020\u0066\u006f\u006e\u0074 d\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072")
	}
	var _edf string
	if _ac, _bff := _gac.GetString(_gab.V); _bff {
		_edf = _ac.Decoded()
	}
	if len(_edf) == 0 {
		return nil, nil
	}
	_gge := []string{_edf}
	_cfb := false
	if _gab.Flags().Has(_eg.FieldFlagMultiline) {
		_cfb = true
		_edf = _ef.Replace(_edf, "\u000d\u000a", "\u000a", -1)
		_edf = _ef.Replace(_edf, "\u000d", "\u000a", -1)
		_gge = _ef.Split(_edf, "\u000a")
	}
	_ccg := make([]string, len(_gge))
	copy(_ccg, _gge)
	_efe := _gdf.MultilineLineHeight
	_ada := 0.0
	_aade := 0
	if _bbf != nil {
		for _abc >= 0 {
			_eaad := make([]string, len(_gge))
			copy(_eaad, _gge)
			_ddg := make([]string, len(_ccg))
			copy(_ddg, _ccg)
			_ada = 0.0
			_aade = 0
			_eea := len(_eaad)
			_cdf := 0
			for _cdf < _eea {
				var _bcb float64
				_bbfg := -1
				_dda := _cf
				if _gdf.MarginLeft != nil {
					_dda = *_gdf.MarginLeft
				}
				for _de, _dfb := range _eaad[_cdf] {
					if _dfb == ' ' {
						_bbfg = _de
						_bcb = _dda
					}
					_baf, _cab := _gga.GetRuneMetrics(_dfb)
					if !_cab {
						_g.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0072\u0075\u006e\u0065\u0020\u006d\u0065\u0074r\u0069\u0063\u0073\u0020\u0066\u006f\u0072\u0020\u0025\u0076\u0020\u002d\u0020\u0073k\u0069p\u0070\u0069\u006e\u0067", _dfb)
						continue
					}
					_dda += _baf.Wx
					if _cfb && !_fae && _abc*_dda/1000.0 > _eda && _bbfg > 0 {
						_fce := _eaad[_cdf][_bbfg+1:]
						_ddf := _ddg[_cdf][_bbfg+1:]
						if _cdf < len(_eaad)-1 {
							_eaad = append(_eaad[:_cdf+1], _eaad[_cdf:]...)
							_eaad[_cdf+1] = _fce
							_ddg = append(_ddg[:_cdf+1], _ddg[_cdf:]...)
							_ddg[_cdf+1] = _ddf
						} else {
							_eaad = append(_eaad, _fce)
							_ddg = append(_ddg, _ddf)
						}
						_eea++
						_eaad[_cdf] = _eaad[_cdf][0:_bbfg]
						_ddg[_cdf] = _ddg[_cdf][0:_bbfg]
						_dda = _bcb
						break
					}
				}
				if _dda > _ada {
					_ada = _dda
				}
				_eaad[_cdf] = string(_bbf.Encode(_eaad[_cdf]))
				if len(_eaad[_cdf]) > 0 {
					_aade++
				}
				_cdf++
			}
			_fdg := _abc
			if _aade > 1 {
				_fdg *= _efe
			}
			_gb := float64(_aade) * _fdg
			if _fae || _gb <= _eaa {
				_gge = _eaad
				_ccg = _ddg
				break
			}
			_abc--
		}
	}
	_def := _cf
	if _gdf.MarginLeft != nil {
		_def = *_gdf.MarginLeft
	}
	if _abc == 0 || _fae && _ada > 0 && _def+_ada*_abc/1000.0 > _eda {
		_abc = 0.95 * 1000.0 * (_eda - _def) / _ada
	}
	_addg := _ba
	{
		if _abd, _dee := _gac.GetIntVal(_gab.Q); _dee {
			switch _abd {
			case 0:
				_addg = _ba
			case 1:
				_addg = _ag
			case 2:
				_addg = _ege
			default:
				_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0071\u0075\u0061\u0064\u0064\u0069\u006e\u0067\u003a\u0020%\u0064\u0020\u002d\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u006c\u0065ft\u0020\u0061\u006c\u0069\u0067\u006e\u006d\u0065\u006e\u0074", _abd)
			}
		}
	}
	_ge := _abc
	if _cfb && _aade > 1 {
		_ge = _efe * _abc
	}
	var _gdg float64
	if _ccc != nil {
		_gdg, _dab = _ccc.GetCapHeight()
		if _dab != nil {
			_g.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0067\u0065\u0074 \u0066\u006f\u006e\u0074\u0020\u0043\u0061\u0070\u0048\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076", _dab)
		}
	}
	if int(_gdg) <= 0 {
		_g.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u0061\u0070\u0048e\u0069\u0067\u0068\u0074\u0020\u006e\u006ft \u0061\u0076\u0061\u0069l\u0061\u0062\u006c\u0065\u0020\u002d\u0020\u0073\u0065tt\u0069\u006eg\u0020\u0074\u006f\u0020\u0031\u0030\u0030\u0030")
		_gdg = 1000
	}
	_fgb := _gdg / 1000.0 * _abc
	_eeg := 0.0
	{
		_fgc := float64(_aade) * _ge
		if _fae && _eeg+_fgc > _eaa {
			_abc = 0.95 * (_eaa - _eeg) / float64(_aade)
			_ge = _abc
			if _cfb && _aade > 1 {
				_ge = _efe * _abc
			}
			_fgb = _gdg / 1000.0 * _abc
			_fgc = float64(_aade) * _ge
		}
		if _eaa > _fgc {
			if _cfb {
				if _gdf.MultilineVAlignMiddle {
					_bda := (_eaa - (_fgc + _fgb)) / 2.0
					_fef := _bda + _fgc + _fgb - _ge
					_eeg = _fef
					if _aade > 1 {
						_eeg = _eeg + (_fgc / _abc * float64(_aade)) - _ge - _fgb
					}
					if _eeg < _fgc {
						_eeg = (_eaa - _fgb) / 2.0
					}
				} else {
					_eeg = _eaa - _ge
					if _eeg > _abc {
						_eeg -= _abc * 0.5
					}
				}
			} else {
				_eeg = (_eaa - _fgb) / 2.0
			}
		}
	}
	_cde.Add_Tf(*_ffa, _abc)
	_cde.Add_Td(_def, _eeg)
	_caf := _def
	_dce := _def
	for _agf, _gc := range _gge {
		_cgf := 0.0
		for _, _adg := range _ccg[_agf] {
			_cfbd, _cfd := _gga.GetRuneMetrics(_adg)
			if !_cfd {
				continue
			}
			_cgf += _cfbd.Wx
		}
		_cag := _cgf / 1000.0 * _abc
		_fec := _eda - _cag
		var _bea float64
		switch _addg {
		case _ba:
			_bea = _caf
		case _ag:
			_bea = _fec / 2
		case _ege:
			_bea = _fec
		}
		_def = _bea - _dce
		if _def > 0.0 {
			_cde.Add_Td(_def, 0)
		}
		_dce = _bea
		_cde.Add_Tj(*_gac.MakeString(_gc))
		if _agf < len(_gge)-1 {
			_cde.Add_Td(0, -_abc*_efe)
		}
	}
	_cde.Add_ET()
	_cde.Add_Q()
	_cde.Add_EMC()
	_ffc := _eg.NewXObjectForm()
	_ffc.Resources = _aaa
	_ffc.BBox = _gac.MakeArrayFromFloats([]float64{0, 0, _dff, _add})
	_ffc.SetContentStream(_cde.Bytes(), _ddba())
	_acf := _gac.MakeDict()
	_acf.Set("\u004e", _ffc.ToPdfObject())
	return _acf, nil
}
