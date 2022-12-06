package pdfaid

import (
	_g "fmt"

	_bg "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaextension"
	_b "github.com/trimmer-io/go-xmp/xmp"
)

var Schema = _bg.Schema{NamespaceURI: Namespace.URI, Prefix: Namespace.Name, Schema: "\u0050D\u0046/\u0041\u0020\u0049\u0044\u0020\u0053\u0063\u0068\u0065\u006d\u0061", Property: []_bg.Property{{Category: _bg.PropertyCategoryInternal, Description: "\u0050\u0061\u0072\u0074 o\u0066\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064", Name: "\u0070\u0061\u0072\u0074", ValueType: _bg.ValueTypeNameInteger}, {Category: _bg.PropertyCategoryInternal, Description: "A\u006d\u0065\u006e\u0064\u006d\u0065n\u0074\u0020\u006f\u0066\u0020\u0050\u0044\u0046\u002fA\u0020\u0073\u0074a\u006ed\u0061\u0072\u0064", Name: "\u0061\u006d\u0064", ValueType: _bg.ValueTypeNameText}, {Category: _bg.PropertyCategoryInternal, Description: "C\u006f\u006e\u0066\u006f\u0072\u006da\u006e\u0063\u0065\u0020\u006c\u0065v\u0065\u006c\u0020\u006f\u0066\u0020\u0050D\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064", Name: "c\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065", ValueType: _bg.ValueTypeNameText}}, ValueType: nil}

func init() { _b.Register(Namespace, _b.XmpMetadata); _bg.RegisterSchema(Namespace, &Schema) }

// SyncToXMP implements xmp.Model interface.
func (_gg *Model) SyncToXMP(d *_b.Document) error { return nil }

// SyncModel implements xmp.Model interface.
func (_gadb *Model) SyncModel(d *_b.Document) error { return nil }

// Model is the XMP model for the PdfA metadata.
type Model struct {
	Part        int    `xmp:"pdfaid:part"`
	Conformance string `xmp:"pdfaid:conformance"`
}

// MakeModel gets or create sa new model for PDF/A ID namespace.
func MakeModel(d *_b.Document) (*Model, error) {
	_gd, _cg := d.MakeModel(Namespace)
	if _cg != nil {
		return nil, _cg
	}
	return _gd.(*Model), nil
}

var Namespace = _b.NewNamespace("\u0070\u0064\u0066\u0061\u0069\u0064", "\u0068\u0074\u0074p\u003a\u002f\u002f\u0077w\u0077\u002e\u0061\u0069\u0069\u006d\u002eo\u0072\u0067\u002f\u0070\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0069\u0064\u002f", NewModel)

// CanTag implements xmp.Model interface.
func (_f *Model) CanTag(tag string) bool { _, _d := _b.GetNativeField(_f, tag); return _d == nil }

// SyncFromXMP implements xmp.Model interface.
func (_gad *Model) SyncFromXMP(d *_b.Document) error { return nil }

// Namespaces implements xmp.Model interface.
func (_ga *Model) Namespaces() _b.NamespaceList { return _b.NamespaceList{Namespace} }

// Can implements xmp.Model interface.
func (_a *Model) Can(nsName string) bool { return Namespace.GetName() == nsName }

// GetTag implements xmp.Model interface.
func (_gb *Model) GetTag(tag string) (string, error) {
	_fb, _da := _b.GetNativeField(_gb, tag)
	if _da != nil {
		return "", _g.Errorf("\u0025\u0073\u003a\u0020\u0025\u0076", Namespace.GetName(), _da)
	}
	return _fb, nil
}

// SetTag implements xmp.Model interface.
func (_ea *Model) SetTag(tag, value string) error {
	if _gc := _b.SetNativeField(_ea, tag, value); _gc != nil {
		return _g.Errorf("\u0025\u0073\u003a\u0020\u0025\u0076", Namespace.GetName(), _gc)
	}
	return nil
}

var _ _b.Model = (*Model)(nil)

// NewModel creates a new model.
func NewModel(name string) _b.Model { return &Model{} }
