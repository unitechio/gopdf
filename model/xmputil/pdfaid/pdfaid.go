package pdfaid

import (
	_d "fmt"

	_b "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaextension"
	_f "github.com/trimmer-io/go-xmp/xmp"
)

var _ _f.Model = (*Model)(nil)

func init() { _f.Register(Namespace, _f.XmpMetadata); _b.RegisterSchema(Namespace, &Schema) }

// SyncModel implements xmp.Model interface.
func (_aa *Model) SyncModel(d *_f.Document) error { return nil }

var Namespace = _f.NewNamespace("\u0070\u0064\u0066\u0061\u0069\u0064", "\u0068\u0074\u0074p\u003a\u002f\u002f\u0077w\u0077\u002e\u0061\u0069\u0069\u006d\u002eo\u0072\u0067\u002f\u0070\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0069\u0064\u002f", NewModel)

// GetTag implements xmp.Model interface.
func (_fb *Model) GetTag(tag string) (string, error) {
	_ga, _fbg := _f.GetNativeField(_fb, tag)
	if _fbg != nil {
		return "", _d.Errorf("\u0025\u0073\u003a\u0020\u0025\u0076", Namespace.GetName(), _fbg)
	}
	return _ga, nil
}

// Can implements xmp.Model interface.
func (_ea *Model) Can(nsName string) bool { return Namespace.GetName() == nsName }

// SyncFromXMP implements xmp.Model interface.
func (_cb *Model) SyncFromXMP(d *_f.Document) error { return nil }

// NewModel creates a new model.
func NewModel(name string) _f.Model { return &Model{} }

// MakeModel gets or create sa new model for PDF/A ID namespace.
func MakeModel(d *_f.Document) (*Model, error) {
	_c, _e := d.MakeModel(Namespace)
	if _e != nil {
		return nil, _e
	}
	return _c.(*Model), nil
}

// Model is the XMP model for the PdfA metadata.
type Model struct {
	Part        int    `xmp:"pdfaid:part"`
	Conformance string `xmp:"pdfaid:conformance"`
}

// SyncToXMP implements xmp.Model interface.
func (_fc *Model) SyncToXMP(d *_f.Document) error { return nil }

// SetTag implements xmp.Model interface.
func (_bc *Model) SetTag(tag, value string) error {
	if _bf := _f.SetNativeField(_bc, tag, value); _bf != nil {
		return _d.Errorf("\u0025\u0073\u003a\u0020\u0025\u0076", Namespace.GetName(), _bf)
	}
	return nil
}

// CanTag implements xmp.Model interface.
func (_bg *Model) CanTag(tag string) bool { _, _ba := _f.GetNativeField(_bg, tag); return _ba == nil }

// Namespaces implements xmp.Model interface.
func (_dab *Model) Namespaces() _f.NamespaceList { return _f.NamespaceList{Namespace} }

var Schema = _b.Schema{NamespaceURI: Namespace.URI, Prefix: Namespace.Name, Schema: "\u0050D\u0046/\u0041\u0020\u0049\u0044\u0020\u0053\u0063\u0068\u0065\u006d\u0061", Property: []_b.Property{{Category: _b.PropertyCategoryInternal, Description: "\u0050\u0061\u0072\u0074 o\u0066\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064", Name: "\u0070\u0061\u0072\u0074", ValueType: _b.ValueTypeNameInteger}, {Category: _b.PropertyCategoryInternal, Description: "A\u006d\u0065\u006e\u0064\u006d\u0065n\u0074\u0020\u006f\u0066\u0020\u0050\u0044\u0046\u002fA\u0020\u0073\u0074a\u006ed\u0061\u0072\u0064", Name: "\u0061\u006d\u0064", ValueType: _b.ValueTypeNameText}, {Category: _b.PropertyCategoryInternal, Description: "C\u006f\u006e\u0066\u006f\u0072\u006da\u006e\u0063\u0065\u0020\u006c\u0065v\u0065\u006c\u0020\u006f\u0066\u0020\u0050D\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064", Name: "c\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065", ValueType: _b.ValueTypeNameText}}, ValueType: nil}
