package fjson

import (
	_c "encoding/json"
	_d "io"
	_f "os"

	_fa "bitbucket.org/shenghui0779/gopdf/common"
	_b "bitbucket.org/shenghui0779/gopdf/core"
	_da "bitbucket.org/shenghui0779/gopdf/model"
)

// FieldValues implements model.FieldValueProvider interface.
func (_fca *FieldData) FieldValues() (map[string]_b.PdfObject, error) {
	_gbd := make(map[string]_b.PdfObject)
	for _, _bbc := range _fca._bf {
		if len(_bbc.Value) > 0 {
			_gbd[_bbc.Name] = _b.MakeString(_bbc.Value)
		}
	}
	return _gbd, nil
}

// LoadFromPDF loads form field data from a PDF.
func LoadFromPDF(rs _d.ReadSeeker) (*FieldData, error) {
	_ga, _ge := _da.NewPdfReader(rs)
	if _ge != nil {
		return nil, _ge
	}
	if _ga.AcroForm == nil {
		return nil, nil
	}
	var _db []fieldValue
	_ceb := _ga.AcroForm.AllFields()
	for _, _a := range _ceb {
		var _fb []string
		_eb := make(map[string]struct{})
		_gea, _gb := _a.FullName()
		if _gb != nil {
			return nil, _gb
		}
		if _de, _cec := _a.V.(*_b.PdfObjectString); _cec {
			_db = append(_db, fieldValue{Name: _gea, Value: _de.Decoded()})
			continue
		}
		var _ef string
		for _, _cg := range _a.Annotations {
			_cb, _ebg := _b.GetName(_cg.AS)
			if _ebg {
				_ef = _cb.String()
			}
			_gec, _fba := _b.GetDict(_cg.AP)
			if !_fba {
				continue
			}
			_fbc, _ := _b.GetDict(_gec.Get("\u004e"))
			for _, _ac := range _fbc.Keys() {
				_cbb := _ac.String()
				if _, _gd := _eb[_cbb]; !_gd {
					_fb = append(_fb, _cbb)
					_eb[_cbb] = struct{}{}
				}
			}
			_ee, _ := _b.GetDict(_gec.Get("\u0044"))
			for _, _fc := range _ee.Keys() {
				_cgc := _fc.String()
				if _, _ag := _eb[_cgc]; !_ag {
					_fb = append(_fb, _cgc)
					_eb[_cgc] = struct{}{}
				}
			}
		}
		_bd := fieldValue{Name: _gea, Value: _ef, Options: _fb}
		_db = append(_db, _bd)
	}
	_be := FieldData{_bf: _db}
	return &_be, nil
}

// FieldImageValues implements model.FieldImageProvider interface.
func (_ec *FieldData) FieldImageValues() (map[string]*_da.Image, error) {
	_bgg := make(map[string]*_da.Image)
	for _, _agc := range _ec._bf {
		if _agc.ImageValue != nil {
			_bgg[_agc.Name] = _agc.ImageValue
		}
	}
	return _bgg, nil
}

// LoadFromJSON loads JSON form data from `r`.
func LoadFromJSON(r _d.Reader) (*FieldData, error) {
	var _fe FieldData
	_bg := _c.NewDecoder(r).Decode(&_fe._bf)
	if _bg != nil {
		return nil, _bg
	}
	return &_fe, nil
}

// SetImage assign model.Image to a specific field identified by fieldName.
func (_bgff *FieldData) SetImage(fieldName string, img *_da.Image, opt []string) error {
	_caf := fieldValue{Name: fieldName, ImageValue: img, Options: opt}
	_bgff._bf = append(_bgff._bf, _caf)
	return nil
}

// LoadFromPDFFile loads form field data from a PDF file.
func LoadFromPDFFile(filePath string) (*FieldData, error) {
	_ea, _ebd := _f.Open(filePath)
	if _ebd != nil {
		return nil, _ebd
	}
	defer _ea.Close()
	return LoadFromPDF(_ea)
}

// SetImageFromFile assign image file to a specific field identified by fieldName.
func (_gbe *FieldData) SetImageFromFile(fieldName string, imagePath string, opt []string) error {
	_eaf, _cd := _f.Open(imagePath)
	if _cd != nil {
		return _cd
	}
	defer _eaf.Close()
	_eef, _cd := _da.ImageHandling.Read(_eaf)
	if _cd != nil {
		_fa.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _cd)
		return _cd
	}
	return _gbe.SetImage(fieldName, _eef, opt)
}

// LoadFromJSONFile loads form field data from a JSON file.
func LoadFromJSONFile(filePath string) (*FieldData, error) {
	_e, _ce := _f.Open(filePath)
	if _ce != nil {
		return nil, _ce
	}
	defer _e.Close()
	return LoadFromJSON(_e)
}

type fieldValue struct {
	Name       string     `json:"name"`
	Value      string     `json:"value"`
	ImageValue *_da.Image `json:"-"`

	// Options lists allowed values if present.
	Options []string `json:"options,omitempty"`
}

// FieldData represents form field data loaded from JSON file.
type FieldData struct{ _bf []fieldValue }

// JSON returns the field data as a string in JSON format.
func (_ad FieldData) JSON() (string, error) {
	_gad, _fg := _c.MarshalIndent(_ad._bf, "", "\u0020\u0020\u0020\u0020")
	return string(_gad), _fg
}
