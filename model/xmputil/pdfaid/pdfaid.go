package pdfaid

import (
	_ce "fmt"

	_eb "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaextension"
	_e "github.com/trimmer-io/go-xmp/xmp"
)

// SyncModel implements xmp.Model interface.
func (_ca *Model) SyncModel(d *_e.Document) error { return nil }
func init()                                       { _e.Register(Namespace, _e.XmpMetadata); _eb.RegisterSchema(Namespace, &Schema) }

// SetTag implements xmp.Model interface.
func (_bf *Model) SetTag(tag, value string) error {
	if _df := _e.SetNativeField(_bf, tag, value); _df != nil {
		return _ce.Errorf("\u0025\u0073\u003a\u0020\u0025\u0076", Namespace.GetName(), _df)
	}
	return nil
}

var Schema = _eb.Schema{NamespaceURI: Namespace.URI, Prefix: Namespace.Name, Schema: "\u0050D\u0046/\u0041\u0020\u0049\u0044\u0020\u0053\u0063\u0068\u0065\u006d\u0061", Property: []_eb.Property{{Category: _eb.PropertyCategoryInternal, Description: "\u0050\u0061\u0072\u0074 o\u0066\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064", Name: "\u0070\u0061\u0072\u0074", ValueType: _eb.ValueTypeNameInteger}, {Category: _eb.PropertyCategoryInternal, Description: "A\u006d\u0065\u006e\u0064\u006d\u0065n\u0074\u0020\u006f\u0066\u0020\u0050\u0044\u0046\u002fA\u0020\u0073\u0074a\u006ed\u0061\u0072\u0064", Name: "\u0061\u006d\u0064", ValueType: _eb.ValueTypeNameText}, {Category: _eb.PropertyCategoryInternal, Description: "C\u006f\u006e\u0066\u006f\u0072\u006da\u006e\u0063\u0065\u0020\u006c\u0065v\u0065\u006c\u0020\u006f\u0066\u0020\u0050D\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064", Name: "c\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065", ValueType: _eb.ValueTypeNameText}}, ValueType: nil}
var _ _e.Model = (*Model)(nil)

// SyncToXMP implements xmp.Model interface.
func (_b *Model) SyncToXMP(d *_e.Document) error { return nil }

// GetTag implements xmp.Model interface.
func (_f *Model) GetTag(tag string) (string, error) {
	_cf, _be := _e.GetNativeField(_f, tag)
	if _be != nil {
		return "", _ce.Errorf("\u0025\u0073\u003a\u0020\u0025\u0076", Namespace.GetName(), _be)
	}
	return _cf, nil
}

// NewModel creates a new model.
func NewModel(name string) _e.Model { return &Model{} }

// Model is the XMP model for the PdfA metadata.
type Model struct {
	Part        int    `xmp:"pdfaid:part"`
	Conformance string `xmp:"pdfaid:conformance"`
}

// CanTag implements xmp.Model interface.
func (_ag *Model) CanTag(tag string) bool { _, _db := _e.GetNativeField(_ag, tag); return _db == nil }

var Namespace = _e.NewNamespace("\u0070\u0064\u0066\u0061\u0069\u0064", "\u0068\u0074\u0074p\u003a\u002f\u002f\u0077w\u0077\u002e\u0061\u0069\u0069\u006d\u002eo\u0072\u0067\u002f\u0070\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0069\u0064\u002f", NewModel)

// MakeModel gets or create sa new model for PDF/A ID namespace.
func MakeModel(d *_e.Document) (*Model, error) {
	_g, _ac := d.MakeModel(Namespace)
	if _ac != nil {
		return nil, _ac
	}
	return _g.(*Model), nil
}

// Can implements xmp.Model interface.
func (_ga *Model) Can(nsName string) bool { return Namespace.GetName() == nsName }

// SyncFromXMP implements xmp.Model interface.
func (_ee *Model) SyncFromXMP(d *_e.Document) error { return nil }

// Namespaces implements xmp.Model interface.
func (_cb *Model) Namespaces() _e.NamespaceList { return _e.NamespaceList{Namespace} }
