package fjson

import (
	_b "encoding/json"
	_a "io"
	_g "os"

	_de "bitbucket.org/shenghui0779/gopdf/common"
	_bf "bitbucket.org/shenghui0779/gopdf/core"
	_c "bitbucket.org/shenghui0779/gopdf/model"
)

// FieldData represents form field data loaded from JSON file.
type FieldData struct{ _da []fieldValue }

// FieldImageValues implements model.FieldImageProvider interface.
func (_ge *FieldData) FieldImageValues() (map[string]*_c.Image, error) {
	_abc := make(map[string]*_c.Image)
	for _, _afa := range _ge._da {
		if _afa.ImageValue != nil {
			_abc[_afa.Name] = _afa.ImageValue
		}
	}
	return _abc, nil
}

// JSON returns the field data as a string in JSON format.
func (_bec FieldData) JSON() (string, error) {
	_gd, _bee := _b.MarshalIndent(_bec._da, "", "\u0020\u0020\u0020\u0020")
	return string(_gd), _bee
}

// FieldValues implements model.FieldValueProvider interface.
func (_bfg *FieldData) FieldValues() (map[string]_bf.PdfObject, error) {
	_ccg := make(map[string]_bf.PdfObject)
	for _, _deb := range _bfg._da {
		if len(_deb.Value) > 0 {
			_ccg[_deb.Name] = _bf.MakeString(_deb.Value)
		}
	}
	return _ccg, nil
}

// LoadFromJSONFile loads form field data from a JSON file.
func LoadFromJSONFile(filePath string) (*FieldData, error) {
	_gf, _cf := _g.Open(filePath)
	if _cf != nil {
		return nil, _cf
	}
	defer _gf.Close()
	return LoadFromJSON(_gf)
}

// SetImageFromFile assign image file to a specific field identified by fieldName.
func (_eb *FieldData) SetImageFromFile(fieldName string, imagePath string, opt []string) error {
	_becd, _ba := _g.Open(imagePath)
	if _ba != nil {
		return _ba
	}
	defer _becd.Close()
	_fd, _ba := _c.ImageHandling.Read(_becd)
	if _ba != nil {
		_de.Log.Error("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _ba)
		return _ba
	}
	return _eb.SetImage(fieldName, _fd, opt)
}

// SetImage assign model.Image to a specific field identified by fieldName.
func (_defc *FieldData) SetImage(fieldName string, img *_c.Image, opt []string) error {
	_ecg := fieldValue{Name: fieldName, ImageValue: img, Options: opt}
	_defc._da = append(_defc._da, _ecg)
	return nil
}

type fieldValue struct {
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	ImageValue *_c.Image `json:"-"`

	// Options lists allowed values if present.
	Options []string `json:"options,omitempty"`
}

// LoadFromPDFFile loads form field data from a PDF file.
func LoadFromPDFFile(filePath string) (*FieldData, error) {
	_fbc, _ab := _g.Open(filePath)
	if _ab != nil {
		return nil, _ab
	}
	defer _fbc.Close()
	return LoadFromPDF(_fbc)
}

// LoadFromPDF loads form field data from a PDF.
func LoadFromPDF(rs _a.ReadSeeker) (*FieldData, error) {
	_ec, _bd := _c.NewPdfReader(rs)
	if _bd != nil {
		return nil, _bd
	}
	if _ec.AcroForm == nil {
		return nil, nil
	}
	var _ca []fieldValue
	_dc := _ec.AcroForm.AllFields()
	for _, _ag := range _dc {
		var _af []string
		_afc := make(map[string]struct{})
		_bg, _be := _ag.FullName()
		if _be != nil {
			return nil, _be
		}
		if _dd, _dg := _ag.V.(*_bf.PdfObjectString); _dg {
			_ca = append(_ca, fieldValue{Name: _bg, Value: _dd.Decoded()})
			continue
		}
		var _gc string
		for _, _fc := range _ag.Annotations {
			_cfa, _ed := _bf.GetName(_fc.AS)
			if _ed {
				_gc = _cfa.String()
			}
			_ce, _aff := _bf.GetDict(_fc.AP)
			if !_aff {
				continue
			}
			_fa, _ := _bf.GetDict(_ce.Get("\u004e"))
			for _, _fag := range _fa.Keys() {
				_bdg := _fag.String()
				if _, _gcb := _afc[_bdg]; !_gcb {
					_af = append(_af, _bdg)
					_afc[_bdg] = struct{}{}
				}
			}
			_eee, _ := _bf.GetDict(_ce.Get("\u0044"))
			for _, _ef := range _eee.Keys() {
				_cb := _ef.String()
				if _, _cc := _afc[_cb]; !_cc {
					_af = append(_af, _cb)
					_afc[_cb] = struct{}{}
				}
			}
		}
		_edd := fieldValue{Name: _bg, Value: _gc, Options: _af}
		_ca = append(_ca, _edd)
	}
	_cd := FieldData{_da: _ca}
	return &_cd, nil
}

// LoadFromJSON loads JSON form data from `r`.
func LoadFromJSON(r _a.Reader) (*FieldData, error) {
	var _e FieldData
	_ee := _b.NewDecoder(r).Decode(&_e._da)
	if _ee != nil {
		return nil, _ee
	}
	return &_e, nil
}
