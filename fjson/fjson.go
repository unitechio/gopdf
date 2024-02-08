package fjson

import (
	_f "encoding/json"
	_gc "io"
	_e "os"

	_c "bitbucket.org/shenghui0779/gopdf/common"
	_b "bitbucket.org/shenghui0779/gopdf/core"
	_ge "bitbucket.org/shenghui0779/gopdf/model"
)

// JSON returns the field data as a string in JSON format.
func (_ab FieldData) JSON() (string, error) {
	_dcg, _gcf := _f.MarshalIndent(_ab._fg, "", "\u0020\u0020\u0020\u0020")
	return string(_dcg), _gcf
}

// FieldValues implements model.FieldValueProvider interface.
func (_gcca *FieldData) FieldValues() (map[string]_b.PdfObject, error) {
	_ac := make(map[string]_b.PdfObject)
	for _, _dbd := range _gcca._fg {
		if len(_dbd.Value) > 0 {
			_ac[_dbd.Name] = _b.MakeString(_dbd.Value)
		}
	}
	return _ac, nil
}

// LoadFromPDFFile loads form field data from a PDF file.
func LoadFromPDFFile(filePath string) (*FieldData, error) {
	_edd, _ffgd := _e.Open(filePath)
	if _ffgd != nil {
		return nil, _ffgd
	}
	defer _edd.Close()
	return LoadFromPDF(_edd)
}

// LoadFromJSON loads JSON form data from `r`.
func LoadFromJSON(r _gc.Reader) (*FieldData, error) {
	var _d FieldData
	_ba := _f.NewDecoder(r).Decode(&_d._fg)
	if _ba != nil {
		return nil, _ba
	}
	return &_d, nil
}

type fieldValue struct {
	Name       string     `json:"name"`
	Value      string     `json:"value"`
	ImageValue *_ge.Image `json:"-"`

	// Options lists allowed values if present.
	Options []string `json:"options,omitempty"`
}

// SetImageFromFile assign image file to a specific field identified by fieldName.
func (_edc *FieldData) SetImageFromFile(fieldName string, imagePath string, opt []string) error {
	_gba, _gbb := _e.Open(imagePath)
	if _gbb != nil {
		return _gbb
	}
	defer _gba.Close()
	_ebgb, _gbb := _ge.ImageHandling.Read(_gba)
	if _gbb != nil {
		_c.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _gbb)
		return _gbb
	}
	return _edc.SetImage(fieldName, _ebgb, opt)
}

// LoadFromPDF loads form field data from a PDF.
func LoadFromPDF(rs _gc.ReadSeeker) (*FieldData, error) {
	_ed, _a := _ge.NewPdfReader(rs)
	if _a != nil {
		return nil, _a
	}
	if _ed.AcroForm == nil {
		return nil, nil
	}
	var _bb []fieldValue
	_cb := _ed.AcroForm.AllFields()
	for _, _db := range _cb {
		var _gf []string
		_cdd := make(map[string]struct{})
		_ff, _bf := _db.FullName()
		if _bf != nil {
			return nil, _bf
		}
		if _da, _gee := _db.V.(*_b.PdfObjectString); _gee {
			_bb = append(_bb, fieldValue{Name: _ff, Value: _da.Decoded()})
			continue
		}
		var _fc string
		for _, _gce := range _db.Annotations {
			_bab, _ebd := _b.GetName(_gce.AS)
			if _ebd {
				_fc = _bab.String()
			}
			_df, _ffg := _b.GetDict(_gce.AP)
			if !_ffg {
				continue
			}
			_fe, _ := _b.GetDict(_df.Get("\u004e"))
			for _, _gg := range _fe.Keys() {
				_ggg := _gg.String()
				if _, _eac := _cdd[_ggg]; !_eac {
					_gf = append(_gf, _ggg)
					_cdd[_ggg] = struct{}{}
				}
			}
			_cg, _ := _b.GetDict(_df.Get("\u0044"))
			for _, _eab := range _cg.Keys() {
				_ebg := _eab.String()
				if _, _cc := _cdd[_ebg]; !_cc {
					_gf = append(_gf, _ebg)
					_cdd[_ebg] = struct{}{}
				}
			}
		}
		_gcc := fieldValue{Name: _ff, Value: _fc, Options: _gf}
		_bb = append(_bb, _gcc)
	}
	_be := FieldData{_fg: _bb}
	return &_be, nil
}

// FieldData represents form field data loaded from JSON file.
type FieldData struct{ _fg []fieldValue }

// SetImage assign model.Image to a specific field identified by fieldName.
func (_gga *FieldData) SetImage(fieldName string, img *_ge.Image, opt []string) error {
	_aba := fieldValue{Name: fieldName, ImageValue: img, Options: opt}
	_gga._fg = append(_gga._fg, _aba)
	return nil
}

// LoadFromJSONFile loads form field data from a JSON file.
func LoadFromJSONFile(filePath string) (*FieldData, error) {
	_eb, _ea := _e.Open(filePath)
	if _ea != nil {
		return nil, _ea
	}
	defer _eb.Close()
	return LoadFromJSON(_eb)
}

// FieldImageValues implements model.FieldImageProvider interface.
func (_gb *FieldData) FieldImageValues() (map[string]*_ge.Image, error) {
	_fab := make(map[string]*_ge.Image)
	for _, _ga := range _gb._fg {
		if _ga.ImageValue != nil {
			_fab[_ga.Name] = _ga.ImageValue
		}
	}
	return _fab, nil
}
