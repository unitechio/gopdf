package fjson

import (
	_d "encoding/json"
	_a "io"
	_da "os"

	_ad "bitbucket.org/shenghui0779/gopdf/common"
	_fb "bitbucket.org/shenghui0779/gopdf/core"
	_f "bitbucket.org/shenghui0779/gopdf/model"
)

// FieldValues implements model.FieldValueProvider interface.
func (_bg *FieldData) FieldValues() (map[string]_fb.PdfObject, error) {
	_bfg := make(map[string]_fb.PdfObject)
	for _, _gd := range _bg._g {
		if len(_gd.Value) > 0 {
			_bfg[_gd.Name] = _fb.MakeString(_gd.Value)
		}
	}
	return _bfg, nil
}

// SetImageFromFile assign image file to a specific field identified by fieldName.
func (_bfba *FieldData) SetImageFromFile(fieldName string, imagePath string, opt []string) error {
	_dcc, _bfc := _da.Open(imagePath)
	if _bfc != nil {
		return _bfc
	}
	defer _dcc.Close()
	_cfe, _bfc := _f.ImageHandling.Read(_dcc)
	if _bfc != nil {
		_ad.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _bfc)
		return _bfc
	}
	return _bfba.SetImage(fieldName, _cfe, opt)
}

// LoadFromJSONFile loads form field data from a JSON file.
func LoadFromJSONFile(filePath string) (*FieldData, error) {
	_ga, _ea := _da.Open(filePath)
	if _ea != nil {
		return nil, _ea
	}
	defer _ga.Close()
	return LoadFromJSON(_ga)
}

// JSON returns the field data as a string in JSON format.
func (_ae FieldData) JSON() (string, error) {
	_bd, _bfd := _d.MarshalIndent(_ae._g, "", "\u0020\u0020\u0020\u0020")
	return string(_bd), _bfd
}

// LoadFromPDFFile loads form field data from a PDF file.
func LoadFromPDFFile(filePath string) (*FieldData, error) {
	_dg, _ee := _da.Open(filePath)
	if _ee != nil {
		return nil, _ee
	}
	defer _dg.Close()
	return LoadFromPDF(_dg)
}

// FieldImageValues implements model.FieldImageProvider interface.
func (_fc *FieldData) FieldImageValues() (map[string]*_f.Image, error) {
	_gda := make(map[string]*_f.Image)
	for _, _aee := range _fc._g {
		if _aee.ImageValue != nil {
			_gda[_aee.Name] = _aee.ImageValue
		}
	}
	return _gda, nil
}

// FieldData represents form field data loaded from JSON file.
type FieldData struct{ _g []fieldValue }
type fieldValue struct {
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	ImageValue *_f.Image `json:"-"`

	// Options lists allowed values if present.
	Options []string `json:"options,omitempty"`
}

// LoadFromPDF loads form field data from a PDF.
func LoadFromPDF(rs _a.ReadSeeker) (*FieldData, error) {
	_b, _fbb := _f.NewPdfReader(rs)
	if _fbb != nil {
		return nil, _fbb
	}
	if _b.AcroForm == nil {
		return nil, nil
	}
	var _ed []fieldValue
	_acg := _b.AcroForm.AllFields()
	for _, _be := range _acg {
		var _fe []string
		_cf := make(map[string]struct{})
		_aa, _add := _be.FullName()
		if _add != nil {
			return nil, _add
		}
		if _ec, _fea := _be.V.(*_fb.PdfObjectString); _fea {
			_ed = append(_ed, fieldValue{Name: _aa, Value: _ec.Decoded()})
			continue
		}
		var _dc string
		for _, _beb := range _be.Annotations {
			_fead, _eae := _fb.GetName(_beb.AS)
			if _eae {
				_dc = _fead.String()
			}
			_adc, _gg := _fb.GetDict(_beb.AP)
			if !_gg {
				continue
			}
			_ff, _ := _fb.GetDict(_adc.Get("\u004e"))
			for _, _ggd := range _ff.Keys() {
				_ce := _ggd.String()
				if _, _eb := _cf[_ce]; !_eb {
					_fe = append(_fe, _ce)
					_cf[_ce] = struct{}{}
				}
			}
			_ede, _ := _fb.GetDict(_adc.Get("\u0044"))
			for _, _cc := range _ede.Keys() {
				_gb := _cc.String()
				if _, _eg := _cf[_gb]; !_eg {
					_fe = append(_fe, _gb)
					_cf[_gb] = struct{}{}
				}
			}
		}
		_ca := fieldValue{Name: _aa, Value: _dc, Options: _fe}
		_ed = append(_ed, _ca)
	}
	_db := FieldData{_g: _ed}
	return &_db, nil
}

// SetImage assign model.Image to a specific field identified by fieldName.
func (_bc *FieldData) SetImage(fieldName string, img *_f.Image, opt []string) error {
	_cee := fieldValue{Name: fieldName, ImageValue: img, Options: opt}
	_bc._g = append(_bc._g, _cee)
	return nil
}

// LoadFromJSON loads JSON form data from `r`.
func LoadFromJSON(r _a.Reader) (*FieldData, error) {
	var _ac FieldData
	_c := _d.NewDecoder(r).Decode(&_ac._g)
	if _c != nil {
		return nil, _c
	}
	return &_ac, nil
}
