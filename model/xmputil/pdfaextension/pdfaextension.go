package pdfaextension

import (
	_d "fmt"
	_bf "reflect"
	_a "sort"
	_gd "strings"
	_b "sync"

	_e "github.com/trimmer-io/go-xmp/models/dc"
	_c "github.com/trimmer-io/go-xmp/models/pdf"
	_da "github.com/trimmer-io/go-xmp/models/xmp_base"
	_ae "github.com/trimmer-io/go-xmp/models/xmp_mm"
	_ga "github.com/trimmer-io/go-xmp/xmp"
)

// RegisterSchema registers schema extension definition.
func RegisterSchema(ns *_ga.Namespace, schema *Schema) {
	_ef.Lock()
	defer _ef.Unlock()
	_ab[ns.URI] = schema
}
func init() {
	_ga.Register(Namespace, _ga.XmpMetadata)
	_ga.Register(SchemaNS)
	_ga.Register(PropertyNS)
	_ga.Register(ValueTypeNS)
	_ga.Register(FieldNS)
}

const (
	ValueTypeNameBoolean        ValueTypeName = "\u0042o\u006f\u006c\u0065\u0061\u006e"
	ValueTypeNameDate           ValueTypeName = "\u0044\u0061\u0074\u0065"
	ValueTypeNameInteger        ValueTypeName = "\u0049n\u0074\u0065\u0067\u0065\u0072"
	ValueTypeNameReal           ValueTypeName = "\u0052\u0065\u0061\u006c"
	ValueTypeNameText           ValueTypeName = "\u0054\u0065\u0078\u0074"
	ValueTypeNameAgentName      ValueTypeName = "\u0041g\u0065\u006e\u0074\u004e\u0061\u006de"
	ValueTypeNameProperName     ValueTypeName = "\u0050\u0072\u006f\u0070\u0065\u0072\u004e\u0061\u006d\u0065"
	ValueTypeNameXPath          ValueTypeName = "\u0058\u0050\u0061t\u0068"
	ValueTypeNameGUID           ValueTypeName = "\u0047\u0055\u0049\u0044"
	ValueTypeNameLocale         ValueTypeName = "\u004c\u006f\u0063\u0061\u006c\u0065"
	ValueTypeNameMIMEType       ValueTypeName = "\u004d\u0049\u004d\u0045\u0054\u0079\u0070\u0065"
	ValueTypeNameRenditionClass ValueTypeName = "\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006eC\u006c\u0061\u0073\u0073"
	ValueTypeNameResourceRef    ValueTypeName = "R\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0052\u0065\u0066"
	ValueTypeNameURL            ValueTypeName = "\u0055\u0052\u004c"
	ValueTypeNameURI            ValueTypeName = "\u0055\u0052\u0049"
	ValueTypeNameVersion        ValueTypeName = "\u0056e\u0072\u0073\u0069\u006f\u006e"
)

// MarshalXMP implements xmp.Marshaler interface.
func (_gb Properties) MarshalXMP(e *_ga.Encoder, node *_ga.Node, m _ga.Model) error {
	return _ga.MarshalArray(e, node, _gb.Typ(), _gb)
}

// SetSchema sets the schema into given model.
func (_fa *Model) SetSchema(namespaceURI string, s Schema) {
	for _de := 0; _de < len(_fa.Schemas); _de++ {
		if _fa.Schemas[_de].NamespaceURI == namespaceURI {
			_fa.Schemas[_de] = s
			return
		}
	}
	_fa.Schemas = append(_fa.Schemas, s)
}

// AltOfValueTypeName gets the ValueTypeName of the alt of given value type names.
func AltOfValueTypeName(vt ValueTypeName) ValueTypeName { return "\u0061\u006c\u0074\u0020" + vt }

// GetTag implements xmp.Model interface.
func (_bfb *Model) GetTag(tag string) (string, error) {
	_aea, _acg := _ga.GetNativeField(_bfb, tag)
	if _acg != nil {
		return "", _d.Errorf("\u0025\u0073\u003a\u0020\u0025\u0076", Namespace.GetName(), _acg)
	}
	return _aea, nil
}

// Can implements xmp.Model interface.
func (_fg *Model) Can(nsName string) bool { return Namespace.GetName() == nsName }

const (
	PropertyCategoryUndefined PropertyCategory = iota
	PropertyCategoryInternal
	PropertyCategoryExternal
)

// Namespaces implements xmp.Model interface.
func (_ag *Model) Namespaces() _ga.NamespaceList { return _ga.NamespaceList{Namespace} }
func (_ba Schemas) Typ() _ga.ArrayType           { return _ga.ArrayTypeUnordered }

// Typ gets array type of the field value types.
func (_cg ValueTypes) Typ() _ga.ArrayType { return _ga.ArrayTypeOrdered }

// PropertyCategory is the property category enumerator.
type PropertyCategory int

// ClosedChoiceValueTypeName gets the closed choice of provided value type name.
func ClosedChoiceValueTypeName(vt ValueTypeName) ValueTypeName {
	return "\u0043\u006c\u006f\u0073\u0065\u0064\u0020\u0043\u0068\u006f\u0069\u0063e\u0020\u006f\u0066\u0020" + vt
}

// UnmarshalXMP implements xmp.Unmarshaler interface.
func (_efd *Schemas) UnmarshalXMP(d *_ga.Decoder, node *_ga.Node, m _ga.Model) error {
	return _ga.UnmarshalArray(d, node, _efd.Typ(), _efd)
}

const (
	ValueTypeResourceEvent ValueTypeName = "\u0052\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0045\u0076\u0065\u006e\u0074"
)

var (
	Namespace   = _ga.NewNamespace("\u0070\u0064\u0066\u0061\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e", "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0077\u0077\u0077\u002e\u0061\u0069\u0069\u006d\u002e\u006f\u0072\u0067\u002f\u0070\u0064\u0066\u0061/\u006e\u0073\u002f\u0065\u0078t\u0065\u006es\u0069\u006f\u006e\u002f", NewModel)
	SchemaNS    = _ga.NewNamespace("\u0070\u0064\u0066\u0061\u0053\u0063\u0068\u0065\u006d\u0061", "h\u0074\u0074\u0070\u003a\u002f\u002fw\u0077\u0077\u002e\u0061\u0069\u0069m\u002e\u006f\u0072\u0067\u002f\u0070\u0064f\u0061\u002f\u006e\u0073\u002f\u0073\u0063\u0068\u0065\u006da\u0023", nil)
	PropertyNS  = _ga.NewNamespace("\u0070\u0064\u0066a\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0079", "\u0068\u0074t\u0070\u003a\u002f\u002fw\u0077\u0077.\u0061\u0069\u0069\u006d\u002e\u006f\u0072\u0067/\u0070\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0070\u0072\u006f\u0070e\u0072\u0074\u0079\u0023", nil)
	ValueTypeNS = _ga.NewNamespace("\u0070\u0064\u0066\u0061\u0054\u0079\u0070\u0065", "\u0068\u0074\u0074\u0070\u003a\u002f/\u0077\u0077\u0077\u002e\u0061\u0069\u0069\u006d\u002e\u006f\u0072\u0067\u002fp\u0064\u0066\u0061\u002f\u006e\u0073\u002ft\u0079\u0070\u0065\u0023", nil)
	FieldNS     = _ga.NewNamespace("\u0070d\u0066\u0061\u0046\u0069\u0065\u006cd", "\u0068\u0074\u0074p:\u002f\u002f\u0077\u0077\u0077\u002e\u0061\u0069\u0069m\u002eo\u0072g\u002fp\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0066\u0069\u0065\u006c\u0064\u0023", nil)
)

func init() {
	_ab = map[string]*Schema{_ae.NsXmpMM.URI: &XmpMediaManagementSchema, "\u0078\u006d\u0070\u0069\u0064\u0071": &XmpIDQualSchema, _c.NsPDF.URI: &PdfSchema, "\u0073\u0074\u0045v\u0074": &FieldResourceEventSchema, "\u0073\u0074\u0056e\u0072": &FieldVersionSchema}
}

var XmpIDQualSchema = Schema{NamespaceURI: "\u0068t\u0074\u0070:\u002f\u002f\u006e\u0073.\u0061\u0064\u006fb\u0065\u002e\u0063\u006f\u006d\u002f\u0078\u006d\u0070/I\u0064\u0065\u006et\u0069\u0066i\u0065\u0072\u002f\u0071\u0075\u0061l\u002f\u0031.\u0030\u002f", Prefix: "\u0078\u006d\u0070\u0069\u0064\u0071", Schema: "X\u004dP\u0020\u0078\u006d\u0070\u0069\u0064\u0071\u0020q\u0075\u0061\u006c\u0069fi\u0065\u0072", Property: []Property{{Category: PropertyCategoryInternal, Name: "\u0053\u0063\u0068\u0065\u006d\u0065", Description: "\u0041\u0020\u0071\u0075\u0061\u006c\u0069\u0066\u0069\u0065\u0072\u0020\u0070\u0072o\u0076\u0069\u0064i\u006e\u0067\u0020\u0074h\u0065\u0020\u006e\u0020\u0061\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u006c \u0069\u0064\u0065n\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u0068\u0065\u006d\u0065\u0020\u0075\u0073ed\u0020\u0066\u006fr\u0020\u0061\u006e\u0020\u0069\u0074\u0065\u006d \u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0078\u006d\u0070\u003a\u0049\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u002e", ValueType: ValueTypeNameText}}, ValueType: nil}
var PdfSchema = Schema{NamespaceURI: "\u0068\u0074\u0074\u0070:\u002f\u002f\u006e\u0073\u002e\u0061\u0064\u006f\u0062\u0065.\u0063o\u006d\u002f\u0070\u0064\u0066\u002f\u0031.\u0033\u002f", Prefix: "\u0070\u0064\u0066", Schema: "\u0041\u0064o\u0062\u0065\u0020P\u0044\u0046\u0020\u0053\u0063\u0068\u0065\u006d\u0061", Property: Properties{{Category: PropertyCategoryInternal, Description: "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", Name: "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", ValueType: ValueTypeNameText}, {Category: PropertyCategoryInternal, Description: "\u0054\u0068\u0065\u0020\u0050\u0044F\u0020\u0066\u0069l\u0065\u0020\u0076e\u0072\u0073\u0069\u006f\u006e\u0020\u0028\u0066\u006f\u0072 \u0065\u0078\u0061\u006d\u0070le\u003a\u0020\u0031\u002e\u0030\u002c\u0020\u0031\u002e\u0033\u002c\u0020\u0061\u006e\u0064\u0020\u0073\u006f\u0020\u006f\u006e\u0029\u002e", Name: "\u0050\u0044\u0046\u0056\u0065\u0072\u0073\u0069\u006f\u006e", ValueType: ValueTypeNameText}, {Category: PropertyCategoryInternal, Description: "\u0054\u0068\u0065\u0020\u006ea\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0074\u006f\u006fl\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0064\u0020\u0074\u0068\u0065\u0020\u0050\u0044\u0046\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074.", Name: "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", ValueType: ValueTypeNameAgentName}, {Category: PropertyCategoryInternal, Description: "\u0049\u006e\u0064\u0069\u0063\u0061\u0074\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0074\u0068\u0069s\u0020\u0069\u0073\u0020\u0061\u0020\u0072i\u0067\u0068\u0074\u0073\u002d\u006d\u0061\u006e\u0061\u0067\u0065d\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u002e\u002e", Name: "\u004d\u0061\u0072\u006b\u0065\u0064", ValueType: ValueTypeNameBoolean}}}

// Schemas is the array of xmp metadata extension resources.
type Schemas []Schema

// MarshalXMP implements xmp.Marshaler interface.
func (_fad Schemas) MarshalXMP(e *_ga.Encoder, node *_ga.Node, m _ga.Model) error {
	return _ga.MarshalArray(e, node, _fad.Typ(), _fad)
}

// FieldValueType is a schema that describes a field in a structured type.
type FieldValueType struct {
	Name        string        `xmp:"pdfaField:description"`
	Description string        `xmp:"pdfaField:name"`
	ValueType   ValueTypeName `xmp:"pdfaField:valueType"`
}

// SyncModel implements xmp.Model interface.
func (_eef *Model) SyncModel(d *_ga.Document) error { return nil }

// Properties is a list of properties.
type Properties []Property

// SyncFromXMP implements xmp.Model interface.
func (_gg *Model) SyncFromXMP(d *_ga.Document) error { return nil }

// GetSchema for provided namespace.
func GetSchema(namespaceURI string) (*Schema, bool) {
	_ef.RLock()
	defer _ef.RUnlock()
	_eadcc, _abe := _ab[namespaceURI]
	return _eadcc, _abe
}

var (
	_ab = map[string]*Schema{}
	_ef _b.RWMutex
)

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (_bc *PropertyCategory) UnmarshalText(in []byte) error {
	switch string(in) {
	case "\u0069\u006e\u0074\u0065\u0072\u006e\u0061\u006c":
		*_bc = PropertyCategoryInternal
	case "\u0065\u0078\u0074\u0065\u0072\u006e\u0061\u006c":
		*_bc = PropertyCategoryExternal
	default:
		*_bc = PropertyCategoryUndefined
	}
	return nil
}

var XmpMediaManagementSchema = Schema{NamespaceURI: "\u0068\u0074\u0074p\u003a\u002f\u002f\u006es\u002e\u0061\u0064\u006f\u0062\u0065\u002ec\u006f\u006d\u002f\u0078\u0061\u0070\u002f\u0031\u002e\u0030\u002f\u006d\u006d\u002f", Prefix: "\u0078\u006d\u0070M\u004d", Schema: "X\u004dP\u0020\u004d\u0065\u0064\u0069\u0061\u0020\u004da\u006e\u0061\u0067\u0065me\u006e\u0074", Property: []Property{{Category: PropertyCategoryInternal, Description: "\u0041\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u006fr\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0066\u0072\u006f\u006d\u0020w\u0068\u0069\u0063\u0068\u0020\u0074\u0068\u0069\u0073\u0020\u006f\u006e\u0065\u0020i\u0073\u0020\u0064e\u0072\u0069\u0076\u0065\u0064\u002e", Name: "D\u0065\u0072\u0069\u0076\u0065\u0064\u0046\u0072\u006f\u006d", ValueType: ValueTypeNameResourceRef}, {Category: PropertyCategoryInternal, Description: "\u0054\u0068\u0065\u0020\u0063\u006fm\u006d\u006f\u006e\u0020\u0069d\u0065\u006e\u0074\u0069\u0066\u0069e\u0072\u0020\u0066\u006f\u0072 \u0061\u006c\u006c\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0073 \u0061\u006e\u0064\u0020\u0072\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u0020o\u0066\u0020\u0061\u0020\u0072\u0065\u0073\u006fu\u0072\u0063\u0065", Name: "\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0049\u0044", ValueType: ValueTypeNameURI}, {Category: PropertyCategoryInternal, Description: "UU\u0049\u0044 \u0062\u0061\u0073\u0065\u0064\u0020\u0069\u0064\u0065n\u0074\u0069\u0066\u0069\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0020\u0069\u006e\u0063\u0061\u0072\u006e\u0061\u0074i\u006fn\u0020\u006f\u0066\u0020\u0061\u0020\u0064\u006fc\u0075m\u0065\u006et", Name: "\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0049\u0044", ValueType: ValueTypeNameURI}, {Category: PropertyCategoryInternal, Description: "\u0054\u0068\u0065\u0020\u0063\u006fm\u006d\u006f\u006e\u0020\u0069d\u0065\u006e\u0074\u0069\u0066\u0069e\u0072\u0020\u0066\u006f\u0072 \u0061\u006c\u006c\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0073 \u0061\u006e\u0064\u0020\u0072\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u0020o\u0066\u0020\u0061\u0020\u0064\u006f\u0063\u0075m\u0065\u006e\u0074", Name: "\u004fr\u0069g\u0069\u006e\u0061\u006c\u0044o\u0063\u0075m\u0065\u006e\u0074\u0049\u0044", ValueType: ValueTypeNameURI}, {Category: PropertyCategoryInternal, Description: "\u0054\u0068\u0065 \u0072\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0020\u006e\u0061\u006d\u0065\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0069\u0073 \u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065", Name: "\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006eC\u006c\u0061\u0073\u0073", ValueType: ValueTypeNameRenditionClass}, {Category: PropertyCategoryInternal, Description: "\u0043\u0061n\u0020\u0062\u0065\u0020\u0075\u0073\u0065d\u0020t\u006f\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0072\u0065\u006e\u0064\u0069t\u0069\u006f\u006e\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0074\u0068a\u0074 \u0061r\u0065\u0020\u0074o\u006f\u0020\u0063\u006f\u006d\u0070\u006c\u0065\u0078\u0020\u006f\u0072\u0020\u0076\u0065\u0072\u0062o\u0073\u0065\u0020\u0074\u006f\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0020\u0069\u006e", Name: "\u0052e\u006ed\u0069\u0074\u0069\u006f\u006e\u0050\u0061\u0072\u0061\u006d\u0073", ValueType: ValueTypeNameText}, {Category: PropertyCategoryInternal, Description: "\u0054\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u0069\u0064\u0065\u006e\u0074i\u0066\u0069\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0074\u0068i\u0073\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u002e", Name: "\u0056e\u0072\u0073\u0069\u006f\u006e\u0049D", ValueType: ValueTypeNameText}, {Category: PropertyCategoryExternal, Description: "\u0054\u0068\u0065\u0020\u0076\u0065r\u0073\u0069\u006f\u006e\u0020\u0068\u0069\u0073\u0074\u006f\u0072\u0079\u0020\u0061\u0073\u0073\u006f\u0063\u0069\u0061t\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0020\u0074\u0068\u0069\u0073\u0020\u0072e\u0073o\u0075\u0072\u0063\u0065", Name: "\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u0073", ValueType: SeqOfValueTypeName(ValueTypeNameVersion)}, {Category: PropertyCategoryInternal, Description: "\u0054\u0068\u0065\u0020\u0072\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0072\u0065\u0073\u006f\u0075r\u0063\u0065\u0027\u0073\u0020\u006d\u0061n\u0061\u0067\u0065\u0072", Name: "\u004da\u006e\u0061\u0067\u0065\u0072", ValueType: ValueTypeNameAgentName}, {Category: PropertyCategoryInternal, Description: "\u0054\u0068\u0065 \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0027\u0073\u0020\u006d\u0061\u006e\u0061\u0067\u0065\u0072 \u0076\u0061\u0072\u0069\u0061\u006e\u0074\u002e", Name: "\u004d\u0061\u006e\u0061\u0067\u0065\u0072\u0056\u0061r\u0069\u0061\u006e\u0074", ValueType: ValueTypeNameText}, {Category: PropertyCategoryInternal, Description: "\u0054\u0068\u0065 \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0027\u0073\u0020\u006d\u0061\u006e\u0061\u0067\u0065\u0072 \u0076\u0061\u0072\u0069\u0061\u006e\u0074\u002e", Name: "\u004d\u0061\u006e\u0061\u0067\u0065\u0072\u0056\u0061r\u0069\u0061\u006e\u0074", ValueType: ValueTypeNameText}}}

// UnmarshalXMP implements xmp.Unmarshaler interface.
func (_fb *Properties) UnmarshalXMP(d *_ga.Decoder, node *_ga.Node, m _ga.Model) error {
	return _ga.UnmarshalArray(d, node, _fb.Typ(), _fb)
}

var FieldVersionSchema = Schema{NamespaceURI: "\u0068\u0074\u0074p\u003a\u002f\u002f\u006e\u0073\u002e\u0061\u0064\u006f\u0062\u0065\u002e\u0063\u006f\u006d\u002f\u0078\u0061\u0070\u002f\u0031\u002e\u0030\u002f\u0073\u0054\u0079\u0070\u0065/\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u0023", Prefix: "\u0073\u0074\u0056e\u0072", Schema: "\u0042a\u0073\u0069\u0063\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0074y\u0070\u0065\u0020\u0056\u0065\u0072\u0073\u0069\u006f\u006e", Property: []Property{{Category: PropertyCategoryInternal, Description: "\u0043\u006fmm\u0065\u006e\u0074s\u0020\u0063\u006f\u006ecer\u006ein\u0067\u0020\u0077\u0068\u0061\u0074\u0020wa\u0073\u0020\u0063\u0068\u0061\u006e\u0067e\u0064", Name: "\u0063\u006f\u006d\u006d\u0065\u006e\u0074\u0073", ValueType: ValueTypeNameText}, {Category: PropertyCategoryInternal, Description: "\u0048\u0069\u0067\u0068\u0020\u006c\u0065\u0076\u0065\u006c\u002c\u0020\u0066\u006f\u0072\u006d\u0061\u006c\u0020\u0064e\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e\u0020\u006f\u0066\u0020\u0077\u0068\u0061\u0074\u0020\u006f\u0070\u0065r\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0068\u0065\u0020\u0075\u0073\u0065\u0072 \u0070\u0065\u0072f\u006f\u0072\u006d\u0065\u0064\u002e", Name: "\u0065\u0076\u0065n\u0074", ValueType: ValueTypeResourceEvent}, {Category: PropertyCategoryInternal, Description: "\u0054\u0068e\u0020\u0064\u0061\u0074\u0065\u0020\u006f\u006e\u0020\u0077\u0068\u0069\u0063\u0068\u0020\u0074\u0068\u0069\u0073\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0065\u0063\u006b\u0065\u0064\u0020\u0069\u006e\u002e", Name: "\u006d\u006f\u0064\u0069\u0066\u0079\u0044\u0061\u0074\u0065", ValueType: ValueTypeNameDate}, {Category: PropertyCategoryInternal, Description: "\u0054\u0068e\u0020\u0070\u0065\u0072s\u006f\u006e \u0077\u0068\u006f\u0020\u006d\u006f\u0064\u0069f\u0069\u0065\u0064\u0020\u0074\u0068\u0069\u0073\u0020\u0076\u0065\u0072s\u0069\u006f\u006e\u002e", Name: "\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0072", ValueType: ValueTypeNameProperName}, {Category: PropertyCategoryInternal, Description: "\u0054\u0068\u0065 n\u0065\u0077\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u002e", Name: "\u0076e\u0072\u0073\u0069\u006f\u006e", ValueType: ValueTypeNameText}}, ValueType: nil}

// ValueTypeName is the name of the value type.
type ValueTypeName string

// MarshalXMP implements xmp.Marshaler interface.
func (_fdg ValueTypes) MarshalXMP(e *_ga.Encoder, node *_ga.Node, m _ga.Model) error {
	return _ga.MarshalArray(e, node, _fdg.Typ(), _fdg)
}

// Schema is the pdfa extension schema.
type Schema struct {

	// NamespaceURI is schema namespace URI.
	NamespaceURI string `xmp:"pdfaSchema:namespaceURI"`

	// Prefix is preferred schema namespace prefix.
	Prefix string `xmp:"pdfaSchema:prefix"`

	// Schema is optional description.
	Schema string `xmp:"pdfaSchema:schema"`

	// Property is description of schema properties.
	Property Properties `xmp:"pdfaSchema:property"`

	// ValueType is description of schema-specific value types.
	ValueType ValueTypes `xmp:"pdfaSchema:valueType"`
}

// CanTag implements xmp.Model interface.
func (_bb *Model) CanTag(tag string) bool {
	_, _bbd := _ga.GetNativeField(_bb, tag)
	return _bbd == nil
}

// SyncToXMP implements xmp.Model interface.
func (_gf *Model) SyncToXMP(d *_ga.Document) error { return nil }

// MarshalText implements encoding.TextMarshaler interface.
func (_ec PropertyCategory) MarshalText() ([]byte, error) {
	switch _ec {
	case PropertyCategoryInternal:
		return []byte("\u0069\u006e\u0074\u0065\u0072\u006e\u0061\u006c"), nil
	case PropertyCategoryExternal:
		return []byte("\u0065\u0078\u0074\u0065\u0072\u006e\u0061\u006c"), nil
	case PropertyCategoryUndefined:
		return []byte(""), nil
	default:
		return nil, _d.Errorf("\u0075\u006ed\u0065\u0066\u0069\u006ee\u0064\u0020p\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0020c\u0061\u0074\u0065\u0067\u006f\u0072\u0079\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076", _ec)
	}
}
func _dad(_ac _ga.Model, _ff *_ga.Namespace, _gag Property) bool {
	_bfa := _bf.ValueOf(_ac).Elem()
	if _bfa.Kind() == _bf.Ptr {
		_bfa = _bfa.Elem()
	}
	_cdg := _bfa.Type()
	for _eaba := 0; _eaba < _cdg.NumField(); _eaba++ {
		_eb := _cdg.Field(_eaba)
		_dbd := _eb.Tag.Get("\u0078\u006d\u0070")
		if _dbd == "" {
			continue
		}
		if !_gd.HasPrefix(_dbd, _ff.Name) {
			continue
		}
		_bg := _gd.IndexRune(_dbd, ':')
		if _bg == -1 {
			continue
		}
		_bd := _dbd[_bg+1:]
		if _bd == _gag.Name {
			_eaf := _bfa.Field(_eaba)
			return !_eaf.IsZero()
		}
	}
	return false
}

// MakeModel creates or gets a model from document.
func MakeModel(d *_ga.Document) (*Model, error) {
	_dc, _f := d.MakeModel(Namespace)
	if _f != nil {
		return nil, _f
	}
	return _dc.(*Model), nil
}

// UnmarshalXMP implements xmp.Unmarshaler interface.
func (_dg *ValueTypes) UnmarshalXMP(d *_ga.Decoder, node *_ga.Node, m _ga.Model) error {
	return _ga.UnmarshalArray(d, node, _dg.Typ(), _dg)
}

// FillModel fills in the document XMP model.
func FillModel(d *_ga.Document, extModel *Model) error {
	var _dca []*_ga.Namespace
	for _, _cb := range d.Namespaces() {
		_dca = append(_dca, _cb)
	}
	_a.Slice(_dca, func(_gab, _ea int) bool { return _dca[_gab].Name < _dca[_ea].Name })
	for _, _aa := range _dca {
		switch _aa {
		case Namespace, SchemaNS, PropertyNS, ValueTypeNS, FieldNS, _da.NsXmp, _e.NsDc:
			continue
		default:
			_ca, _ee := GetSchema(_aa.URI)
			if !_ee {
				continue
			}
			_db := d.FindModel(_aa)
			_bfe := *_ca
			_bfe.Property = Properties{}
			for _, _eab := range _ca.Property {
				_be := _dad(_db, _aa, _eab)
				if _be {
					_bfe.Property = append(_bfe.Property, _eab)
				}
			}
			if len(_bfe.Property) == 0 {
				continue
			}
			var _aef bool
			for _cdf, _ead := range extModel.Schemas {
				if _ead.Schema == _bfe.Schema {
					_aef = true
					extModel.Schemas[_cdf] = _bfe
					break
				}
			}
			if !_aef {
				extModel.Schemas = append(extModel.Schemas, _bfe)
			}
		}
	}
	return nil
}

var FieldResourceEventSchema = Schema{NamespaceURI: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u006es\u002e\u0061\u0064ob\u0065\u002e\u0063\u006f\u006d\u002fx\u0061\u0070\u002f\u0031\u002e\u0030\u002f\u0073\u0054\u0079\u0070\u0065\u002f\u0052\u0065s\u006f\u0075\u0072\u0063\u0065\u0045\u0076\u0065n\u0074\u0023", Prefix: "\u0073\u0074\u0045v\u0074", Schema: "\u0044e\u0066\u0069n\u0069\u0074\u0069\u006fn\u0020\u006f\u0066 \u0062\u0061\u0073\u0069\u0063\u0020\u0076\u0061\u006cue\u0020\u0074\u0079p\u0065\u0020R\u0065\u0073\u006f\u0075\u0072\u0063e\u0045\u0076e\u006e\u0074", Property: []Property{{Category: PropertyCategoryInternal, Description: "\u0054he\u0020a\u0063t\u0069\u006f\u006e\u0020\u0074\u0068a\u0074\u0020\u006f\u0063c\u0075\u0072\u0072\u0065\u0064\u002e\u0020\u0044\u0065\u0066\u0069\u006e\u0065d \u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0061\u0072\u0065\u003a\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0065d\u002c \u0063\u006f\u0070\u0069\u0065\u0064\u002c\u0020\u0063\u0072\u0065\u0061\u0074e\u0064\u002c \u0063\u0072\u006fp\u0070\u0065\u0064\u002c\u0020\u0065\u0064\u0069\u0074ed\u002c\u0020\u0066i\u006c\u002d\u0074\u0065r\u0065\u0064\u002c\u0020\u0066\u006fr\u006d\u0061t\u0074\u0065\u0064\u002c\u0020\u0076\u0065\u0072s\u0069\u006f\u006e\u005f\u0075\u0070\u0064\u0061\u0074\u0065\u0064\u002c\u0020\u0070\u0072\u0069\u006e\u0074\u0065\u0064\u002c\u0020\u0070ubli\u0073\u0068\u0065\u0064\u002c\u0020\u006d\u0061\u006e\u0061\u0067\u0065\u0064\u002c\u0020\u0070\u0072\u006f\u0064\u0075\u0063\u0065\u0064\u002c\u0020\u0072\u0065\u0073i\u007ae\u0064.\u004e\u0065\u0077\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065 \u0076\u0065r\u0062\u0073 \u0069n\u0020\u0074\u0068\u0065\u0020\u0070a\u0073\u0074\u0020\u0074\u0065\u006e\u0073\u0065\u002e", Name: "\u0061\u0063\u0074\u0069\u006f\u006e", ValueType: "O\u0070\u0065\u006e\u0020\u0043\u0068\u006f\u0069\u0063\u0065"}, {Category: PropertyCategoryInternal, Description: "T\u0068\u0065\u0020\u0069\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0020\u0049\u0044\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u006d\u006f\u0064\u0069\u0066\u0069\u0065d \u0072\u0065\u0073o\u0075r\u0063\u0065", Name: "\u0069\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0049\u0044", ValueType: ValueTypeNameURI}, {Category: PropertyCategoryInternal, Description: "\u0041\u0064di\u0074\u0069\u006fn\u0061\u006c\u0020\u0064esc\u0072ip\u0074\u0069\u006f\u006e\u0020\u006f\u0066 t\u0068\u0065\u0020\u0061\u0063\u0074\u0069o\u006e", Name: "\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073", ValueType: ValueTypeNameText}, {Category: PropertyCategoryInternal, Description: "\u0054h\u0065\u0020s\u006f\u0066\u0074\u0077a\u0072\u0065\u0020a\u0067\u0065\u006e\u0074\u0020\u0074\u0068\u0061\u0074 p\u0065\u0072\u0066o\u0072\u006de\u0064\u0020\u0074\u0068\u0065\u0020a\u0063\u0074i\u006f\u006e", Name: "\u0073\u006f\u0066\u0074\u0077\u0061\u0072\u0065\u0041\u0067\u0065\u006e\u0074", ValueType: ValueTypeNameAgentName}, {Category: PropertyCategoryInternal, Description: "\u004f\u0070t\u0069\u006f\u006e\u0061\u006c\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061\u006d\u0070\u0020\u006f\u0066\u0020\u0077\u0068\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0020\u006f\u0063\u0063\u0075\u0072\u0072\u0065\u0064", Name: "\u0077\u0068\u0065\u006e", ValueType: ValueTypeNameDate}}, ValueType: nil}

// SeqOfValueTypeName gets a value type name of a sequence of input value type names.
func SeqOfValueTypeName(vt ValueTypeName) ValueTypeName { return "\u0073\u0065\u0071\u0020" + vt }

// SetTag implements xmp.Model interface.
func (_eadc *Model) SetTag(tag, value string) error {
	if _gabd := _ga.SetNativeField(_eadc, tag, value); _gabd != nil {
		return _d.Errorf("\u0025\u0073\u003a\u0020\u0025\u0076", Namespace.GetName(), _gabd)
	}
	return nil
}

// NewModel creates a new pdfAExtension model.
func NewModel(name string) _ga.Model { return &Model{} }

// Model is the pdfa extension metadata model.
type Model struct {
	Schemas Schemas `xmp:"pdfaExtension:schemas"`
}

// Typ gets the array type of properties.
func (_ebe Properties) Typ() _ga.ArrayType { return _ga.ArrayTypeOrdered }

// Property is a schema that describes single property.
type Property struct {

	// Category is the property category.
	Category PropertyCategory `xmp:"pdfaProperty:category"`

	// Description of the property.
	Description string `xmp:"pdfaProperty:description"`

	// Name is a property name.
	Name string `xmp:"pdfaProperty:name"`

	// ValueType is the property value type.
	ValueType ValueTypeName `xmp:"pdfaProperty:valueType"`
}

// ValueType is the pdfa extension value type schema.
type ValueType struct {
	Description  string           `xmp:"pdfaType:description"`
	Field        []FieldValueType `xmp:"pdfaType:field"`
	NamespaceURI string           `xmp:"pdfaType:namespaceURI"`
	Prefix       string           `xmp:"pdfaType:prefix"`
	Type         string           `xmp:"pdfaType:type"`
}

// IsZero checks if the resources list has no entries.
func (_fd Schemas) IsZero() bool { return len(_fd) == 0 }

// BagOfValueTypeName gets the ValueTypeName of the bag of provided value type names.
func BagOfValueTypeName(vt ValueTypeName) ValueTypeName { return "\u0062\u0061\u0067\u0020" + vt }

// ValueTypes is the slice of field value types.
type ValueTypes []FieldValueType
