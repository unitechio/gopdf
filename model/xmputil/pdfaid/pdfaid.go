package pdfaid

import (
	_g "fmt"

	_ga "unitechio/gopdf/gopdf/model/xmputil/pdfaextension"
	_e "github.com/trimmer-io/go-xmp/xmp"
)

// SyncFromXMP implements xmp.Model interface.
func (_bf *Model) SyncFromXMP(d *_e.Document) error { return nil }
func init()                                         { _e.Register(Namespace, _e.XmpMetadata); _ga.RegisterSchema(Namespace, &Schema) }

var _ _e.Model = (*Model)(nil)

// SyncModel implements xmp.Model interface.
func (_be *Model) SyncModel(d *_e.Document) error { return nil }

// SyncToXMP implements xmp.Model interface.
func (_ef *Model) SyncToXMP(d *_e.Document) error { return nil }

// Can implements xmp.Model interface.
func (_eg *Model) Can(nsName string) bool { return Namespace.GetName() == nsName }

// NewModel creates a new model.
func NewModel(name string) _e.Model { return &Model{} }

// Namespaces implements xmp.Model interface.
func (_gae *Model) Namespaces() _e.NamespaceList { return _e.NamespaceList{Namespace} }

// Model is the XMP model for the PdfA metadata.
type Model struct {
	Part        int    `xmp:"pdfaid:part"`
	Conformance string `xmp:"pdfaid:conformance"`
}

var (
	Schema    = _ga.Schema{NamespaceURI: Namespace.URI, Prefix: Namespace.Name, Schema: "\u0050D\u0046/\u0041\u0020\u0049\u0044\u0020\u0053\u0063\u0068\u0065\u006d\u0061", Property: []_ga.Property{{Category: _ga.PropertyCategoryInternal, Description: "\u0050\u0061\u0072\u0074 o\u0066\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064", Name: "\u0070\u0061\u0072\u0074", ValueType: _ga.ValueTypeNameInteger}, {Category: _ga.PropertyCategoryInternal, Description: "A\u006d\u0065\u006e\u0064\u006d\u0065n\u0074\u0020\u006f\u0066\u0020\u0050\u0044\u0046\u002fA\u0020\u0073\u0074a\u006ed\u0061\u0072\u0064", Name: "\u0061\u006d\u0064", ValueType: _ga.ValueTypeNameText}, {Category: _ga.PropertyCategoryInternal, Description: "C\u006f\u006e\u0066\u006f\u0072\u006da\u006e\u0063\u0065\u0020\u006c\u0065v\u0065\u006c\u0020\u006f\u0066\u0020\u0050D\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064", Name: "c\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065", ValueType: _ga.ValueTypeNameText}}, ValueType: nil}
	Namespace = _e.NewNamespace("\u0070\u0064\u0066\u0061\u0069\u0064", "\u0068\u0074\u0074p\u003a\u002f\u002f\u0077w\u0077\u002e\u0061\u0069\u0069\u006d\u002eo\u0072\u0067\u002f\u0070\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0069\u0064\u002f", NewModel)
)

// GetTag implements xmp.Model interface.
func (_fa *Model) GetTag(tag string) (string, error) {
	_c, _a := _e.GetNativeField(_fa, tag)
	if _a != nil {
		return "", _g.Errorf("\u0025\u0073\u003a\u0020\u0025\u0076", Namespace.GetName(), _a)
	}
	return _c, nil
}

// SetTag implements xmp.Model interface.
func (_gg *Model) SetTag(tag, value string) error {
	if _ee := _e.SetNativeField(_gg, tag, value); _ee != nil {
		return _g.Errorf("\u0025\u0073\u003a\u0020\u0025\u0076", Namespace.GetName(), _ee)
	}
	return nil
}

// MakeModel gets or create sa new model for PDF/A ID namespace.
func MakeModel(d *_e.Document) (*Model, error) {
	_fd, _fb := d.MakeModel(Namespace)
	if _fb != nil {
		return nil, _fb
	}
	return _fd.(*Model), nil
}

// CanTag implements xmp.Model interface.
func (_gc *Model) CanTag(tag string) bool { _, _db := _e.GetNativeField(_gc, tag); return _db == nil }
