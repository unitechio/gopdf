package xmputil

import (
	_e "errors"
	_cg "fmt"
	_f "strconv"
	_gf "time"

	_fea "bitbucket.org/shenghui0779/gopdf/core"
	_eg "bitbucket.org/shenghui0779/gopdf/internal/timeutils"
	_ea "bitbucket.org/shenghui0779/gopdf/internal/uuid"
	_a "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaextension"
	_cd "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaid"
	_c "github.com/trimmer-io/go-xmp/models/pdf"
	_fe "github.com/trimmer-io/go-xmp/models/xmp_mm"
	_g "github.com/trimmer-io/go-xmp/xmp"
)

// MediaManagementDerivedFrom is a structure that contains references of identifiers and versions
// from which given document was derived.
type MediaManagementDerivedFrom struct {
	OriginalDocumentID GUID
	DocumentID         GUID
	InstanceID         GUID
	VersionID          string
}

// GetPdfInfo gets the document pdf info.
func (_fec *Document) GetPdfInfo() (*PdfInfo, bool) {
	_ee, _df := _fec._dg.FindModel(_c.NsPDF).(*_c.PDFInfo)
	if !_df {
		return nil, false
	}
	_bb := PdfInfo{}
	var _bcd *_fea.PdfObjectDictionary
	_bb.Copyright = _ee.Copyright
	_bb.PdfVersion = _ee.PDFVersion
	_bb.Marked = bool(_ee.Marked)
	_ebf := func(_ead string, _egd _fea.PdfObject) {
		if _bcd == nil {
			_bcd = _fea.MakeDict()
		}
		_bcd.Set(_fea.PdfObjectName(_ead), _egd)
	}
	if len(_ee.Title) > 0 {
		_ebf("\u0054\u0069\u0074l\u0065", _fea.MakeString(_ee.Title.Default()))
	}
	if len(_ee.Author) > 0 {
		_ebf("\u0041\u0075\u0074\u0068\u006f\u0072", _fea.MakeString(_ee.Author[0]))
	}
	if _ee.Keywords != "" {
		_ebf("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _fea.MakeString(_ee.Keywords))
	}
	if len(_ee.Subject) > 0 {
		_ebf("\u0053u\u0062\u006a\u0065\u0063\u0074", _fea.MakeString(_ee.Subject.Default()))
	}
	if _ee.Creator != "" {
		_ebf("\u0043r\u0065\u0061\u0074\u006f\u0072", _fea.MakeString(string(_ee.Creator)))
	}
	if _ee.Producer != "" {
		_ebf("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _fea.MakeString(string(_ee.Producer)))
	}
	if _ee.Trapped {
		_ebf("\u0054r\u0061\u0070\u0070\u0065\u0064", _fea.MakeName("\u0054\u0072\u0075\u0065"))
	}
	if !_ee.CreationDate.IsZero() {
		_ebf("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _fea.MakeString(_eg.FormatPdfTime(_ee.CreationDate.Value())))
	}
	if !_ee.ModifyDate.IsZero() {
		_ebf("\u004do\u0064\u0044\u0061\u0074\u0065", _fea.MakeString(_eg.FormatPdfTime(_ee.ModifyDate.Value())))
	}
	_bb.InfoDict = _bcd
	return &_bb, true
}

// SetPdfAExtension sets the pdfaExtension XMP metadata.
func (_ed *Document) SetPdfAExtension() error {
	_db, _ca := _a.MakeModel(_ed._dg)
	if _ca != nil {
		return _ca
	}
	if _ca = _a.FillModel(_ed._dg, _db); _ca != nil {
		return _ca
	}
	if _ca = _db.SyncToXMP(_ed._dg); _ca != nil {
		return _ca
	}
	return nil
}

// NewDocument creates a new document without any previous xmp information.
func NewDocument() *Document { _cge := _g.NewDocument(); return &Document{_dg: _cge} }

// GetGoXmpDocument gets direct access to the go-xmp.Document.
// All changes done to specified document would result in change of this document 'd'.
func (_de *Document) GetGoXmpDocument() *_g.Document { return _de._dg }

// SetMediaManagement sets up XMP media management metadata: namespace xmpMM.
func (_bf *Document) SetMediaManagement(options *MediaManagementOptions) error {
	_ege, _ede := _fe.MakeModel(_bf._dg)
	if _ede != nil {
		return _ede
	}
	if options == nil {
		options = new(MediaManagementOptions)
	}
	_dd := _fe.ResourceRef{}
	if _ege.OriginalDocumentID.IsZero() {
		if options.OriginalDocumentID != "" {
			_ege.OriginalDocumentID = _g.GUID(options.OriginalDocumentID)
		} else {
			_ffg, _ffa := _ea.NewUUID()
			if _ffa != nil {
				return _ffa
			}
			_ege.OriginalDocumentID = _g.GUID(_ffg.String())
		}
	} else {
		_dd.OriginalDocumentID = _ege.OriginalDocumentID
	}
	switch {
	case options.DocumentID != "":
		_ege.DocumentID = _g.GUID(options.DocumentID)
	case options.NewDocumentID || _ege.DocumentID.IsZero():
		if !_ege.DocumentID.IsZero() {
			_dd.DocumentID = _ege.DocumentID
		}
		_gb, _cb := _ea.NewUUID()
		if _cb != nil {
			return _cb
		}
		_ege.DocumentID = _g.GUID(_gb.String())
	}
	if !_ege.InstanceID.IsZero() {
		_dd.InstanceID = _ege.InstanceID
	}
	_ege.InstanceID = _g.GUID(options.InstanceID)
	if _ege.InstanceID == "" {
		_fc, _ddc := _ea.NewUUID()
		if _ddc != nil {
			return _ddc
		}
		_ege.InstanceID = _g.GUID(_fc.String())
	}
	if !_dd.IsZero() {
		_ege.DerivedFrom = &_dd
	}
	_fed := options.VersionID
	if _ege.VersionID != "" {
		_ebg, _aca := _f.Atoi(_ege.VersionID)
		if _aca != nil {
			_fed = _f.Itoa(len(_ege.Versions) + 1)
		} else {
			_fed = _f.Itoa(_ebg + 1)
		}
	}
	if _fed == "" {
		_fed = "\u0031"
	}
	_ege.VersionID = _fed
	_edc := options.ModifyDate
	if _edc.IsZero() {
		_edc = _gf.Now()
	}
	if _ede = _ege.SyncToXMP(_bf._dg); _ede != nil {
		return _ede
	}
	return nil
}

// LoadDocument loads up the xmp document from provided input stream.
func LoadDocument(stream []byte) (*Document, error) {
	_ga := _g.NewDocument()
	if _b := _g.Unmarshal(stream, _ga); _b != nil {
		return nil, _b
	}
	return &Document{_dg: _ga}, nil
}

// MediaManagementOptions are the options for the Media management xmp metadata.
type MediaManagementOptions struct {

	// OriginalDocumentID  as media is imported and projects is started, an original-document ID
	// must be created to identify a new document. This identifies a document as a conceptual entity.
	// By default, this value is generated.
	OriginalDocumentID string

	// NewDocumentID is a flag which generates a new Document identifier while setting media management.
	// This value should be set to true only if the document is stored and saved as new document.
	// Otherwise, if the document is modified and overwrites previous file, it should be set to false.
	NewDocumentID bool

	// DocumentID when a document is copied to a new file path or converted to a new format with
	// Save As, another new document ID should usually be assigned. This identifies a general version or
	// branch of a document. You can use it to track different versions or extracted portions of a document
	// with the same original-document ID.
	// By default, this value is generated if NewDocumentID is true or previous doesn't exist.
	DocumentID string

	// InstanceID to track a document’s editing history, you must assign a new instance ID
	// whenever a document is saved after any changes. This uniquely identifies an exact version of a
	// document. It is used in resource references (to identify both the document or part itself and the
	// referenced or referencing documents), and in document-history resource events (to identify the
	// document instance that resulted from the change).
	// By default, this value is generated.
	InstanceID string

	// DerivedFrom references the source document from which this one is derived,
	// typically through a Save As operation that changes the file name or format. It is a minimal reference;
	// missing components can be assumed to be unchanged. For example, a new version might only need
	// to specify the instance ID and version number of the previous version, or a rendition might only need
	// to specify the instance ID and rendition class of the original.
	// By default, the derived from structure is filled from previous XMP metadata (if exists).
	DerivedFrom string

	// VersionID are meant to associate the document with a product version that is part of a release process. They can be useful in tracking the
	// document history, but should not be used to identify a document uniquely in any context.
	// Usually it simply works by incrementing integers 1,2,3...
	// By default, this values is incremented or set to the next version number.
	VersionID string

	// ModifyComment is a comment to given modification
	ModifyComment string

	// ModifyDate is a custom modification date for the versions.
	// By default, this would be set to time.Now().
	ModifyDate _gf.Time

	// Modifier is a person who did the modification.
	Modifier string
}

// SetPdfInfo sets the pdf info into selected document.
func (_bg *Document) SetPdfInfo(options *PdfInfoOptions) error {
	if options == nil {
		return _e.New("\u006ei\u006c\u0020\u0070\u0064\u0066\u0020\u006f\u0070\u0074\u0069\u006fn\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_ab, _ce := _c.MakeModel(_bg._dg)
	if _ce != nil {
		return _ce
	}
	if options.Overwrite {
		*_ab = _c.PDFInfo{}
	}
	if options.InfoDict != nil {
		_gaf, _def := _fea.GetDict(options.InfoDict)
		if !_def {
			return _cg.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", options.InfoDict)
		}
		var _gc *_fea.PdfObjectString
		for _, _ac := range _gaf.Keys() {
			switch _ac {
			case "\u0054\u0069\u0074l\u0065":
				_gc, _def = _fea.GetString(_gaf.Get("\u0054\u0069\u0074l\u0065"))
				if _def {
					_ab.Title = _g.NewAltString(_gc)
				}
			case "\u0041\u0075\u0074\u0068\u006f\u0072":
				_gc, _def = _fea.GetString(_gaf.Get("\u0041\u0075\u0074\u0068\u006f\u0072"))
				if _def {
					_ab.Author = _g.NewStringList(_gc.String())
				}
			case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":
				_gc, _def = _fea.GetString(_gaf.Get("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"))
				if _def {
					_ab.Keywords = _gc.String()
				}
			case "\u0043r\u0065\u0061\u0074\u006f\u0072":
				_gc, _def = _fea.GetString(_gaf.Get("\u0043r\u0065\u0061\u0074\u006f\u0072"))
				if _def {
					_ab.Creator = _g.AgentName(_gc.String())
				}
			case "\u0053u\u0062\u006a\u0065\u0063\u0074":
				_gc, _def = _fea.GetString(_gaf.Get("\u0053u\u0062\u006a\u0065\u0063\u0074"))
				if _def {
					_ab.Subject = _g.NewAltString(_gc.String())
				}
			case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":
				_gc, _def = _fea.GetString(_gaf.Get("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"))
				if _def {
					_ab.Producer = _g.AgentName(_gc.String())
				}
			case "\u0054r\u0061\u0070\u0070\u0065\u0064":
				_ffb, _cgd := _fea.GetName(_gaf.Get("\u0054r\u0061\u0070\u0070\u0065\u0064"))
				if _cgd {
					switch _ffb.String() {
					case "\u0054\u0072\u0075\u0065":
						_ab.Trapped = true
					case "\u0046\u0061\u006cs\u0065":
						_ab.Trapped = false
					default:
						_ab.Trapped = true
					}
				}
			case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":
				if _gcd, _feb := _fea.GetString(_gaf.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065")); _feb && _gcd.String() != "" {
					_gfb, _cdd := _eg.ParsePdfTime(_gcd.String())
					if _cdd != nil {
						return _cg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077", _cdd)
					}
					_ab.CreationDate = _g.NewDate(_gfb)
				}
			case "\u004do\u0064\u0044\u0061\u0074\u0065":
				if _bd, _be := _fea.GetString(_gaf.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _be && _bd.String() != "" {
					_fdf, _eab := _eg.ParsePdfTime(_bd.String())
					if _eab != nil {
						return _cg.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _eab)
					}
					_ab.ModifyDate = _g.NewDate(_fdf)
				}
			}
		}
	}
	if options.PdfVersion != "" {
		_ab.PDFVersion = options.PdfVersion
	}
	if options.Marked {
		_ab.Marked = _g.Bool(options.Marked)
	}
	if options.Copyright != "" {
		_ab.Copyright = options.Copyright
	}
	if _ce = _ab.SyncToXMP(_bg._dg); _ce != nil {
		return _ce
	}
	return nil
}

// Document is an implementation of the xmp document.
// It is a wrapper over go-xmp/xmp.Document that provides some Pdf predefined functionality.
type Document struct{ _dg *_g.Document }

// GUID is a string representing a globally unique identifier.
type GUID string

// MediaManagementVersion is the version of the media management xmp metadata.
type MediaManagementVersion struct {
	VersionID  string
	ModifyDate _gf.Time
	Comments   string
	Modifier   string
}

// MediaManagement are the values from the document media management metadata.
type MediaManagement struct {

	// OriginalDocumentID  as media is imported and projects is started, an original-document ID
	// must be created to identify a new document. This identifies a document as a conceptual entity.
	OriginalDocumentID GUID

	// DocumentID when a document is copied to a new file path or converted to a new format with
	// Save As, another new document ID should usually be assigned. This identifies a general version or
	// branch of a document. You can use it to track different versions or extracted portions of a document
	// with the same original-document ID.
	DocumentID GUID

	// InstanceID to track a document’s editing history, you must assign a new instance ID
	// whenever a document is saved after any changes. This uniquely identifies an exact version of a
	// document. It is used in resource references (to identify both the document or part itself and the
	// referenced or referencing documents), and in document-history resource events (to identify the
	// document instance that resulted from the change).
	InstanceID GUID

	// DerivedFrom references the source document from which this one is derived,
	// typically through a Save As operation that changes the file name or format. It is a minimal reference;
	// missing components can be assumed to be unchanged. For example, a new version might only need
	// to specify the instance ID and version number of the previous version, or a rendition might only need
	// to specify the instance ID and rendition class of the original.
	DerivedFrom *MediaManagementDerivedFrom

	// VersionID are meant to associate the document with a product version that is part of a release process. They can be useful in tracking the
	// document history, but should not be used to identify a document uniquely in any context.
	// Usually it simply works by incrementing integers 1,2,3...
	VersionID string

	// Versions is the history of the document versions along with the comments, timestamps and issuers.
	Versions []MediaManagementVersion
}

// SetPdfAID sets up pdfaid xmp metadata.
// In example: Part: '1' Conformance: 'B' states for PDF/A 1B.
func (_fda *Document) SetPdfAID(part int, conformance string) error {
	_fef, _eac := _cd.MakeModel(_fda._dg)
	if _eac != nil {
		return _eac
	}
	_fef.Part = part
	_fef.Conformance = conformance
	if _caf := _fef.SyncToXMP(_fda._dg); _caf != nil {
		return _caf
	}
	return nil
}

// PdfAID is the result of the XMP pdfaid metadata.
type PdfAID struct {
	Part        int
	Conformance string
}

// Marshal the document into xml byte stream.
func (_ge *Document) Marshal() ([]byte, error) {
	if _ge._dg.IsDirty() {
		if _ff := _ge._dg.SyncModels(); _ff != nil {
			return nil, _ff
		}
	}
	return _g.Marshal(_ge._dg)
}

// GetPdfAID gets the pdfaid xmp metadata model.
func (_cafd *Document) GetPdfAID() (*PdfAID, bool) {
	_ebgd, _ba := _cafd._dg.FindModel(_cd.Namespace).(*_cd.Model)
	if !_ba {
		return nil, false
	}
	return &PdfAID{Part: _ebgd.Part, Conformance: _ebgd.Conformance}, true
}

// GetPdfaExtensionSchemas gets a pdfa extension schemas.
func (_ef *Document) GetPdfaExtensionSchemas() ([]_a.Schema, error) {
	_cc := _ef._dg.FindModel(_a.Namespace)
	if _cc == nil {
		return nil, nil
	}
	_bc, _cf := _cc.(*_a.Model)
	if !_cf {
		return nil, _cg.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u006d\u006f\u0064\u0065l f\u006fr \u0070\u0064\u0066\u0061\u0045\u0078\u0074en\u0073\u0069\u006f\u006e\u0073\u003a\u0020%\u0054", _cc)
	}
	return _bc.Schemas, nil
}

// GetMediaManagement gets the media management metadata from provided xmp document.
func (_fcd *Document) GetMediaManagement() (*MediaManagement, bool) {
	_egg := _fe.FindModel(_fcd._dg)
	if _egg == nil {
		return nil, false
	}
	_ebb := make([]MediaManagementVersion, len(_egg.Versions))
	for _ec, _gd := range _egg.Versions {
		_ebb[_ec] = MediaManagementVersion{VersionID: _gd.Version, ModifyDate: _gd.ModifyDate.Value(), Comments: _gd.Comments, Modifier: _gd.Modifier}
	}
	_gbg := &MediaManagement{OriginalDocumentID: GUID(_egg.OriginalDocumentID.Value()), DocumentID: GUID(_egg.DocumentID.Value()), InstanceID: GUID(_egg.InstanceID.Value()), VersionID: _egg.VersionID, Versions: _ebb}
	if _egg.DerivedFrom != nil {
		_gbg.DerivedFrom = &MediaManagementDerivedFrom{OriginalDocumentID: GUID(_egg.DerivedFrom.OriginalDocumentID), DocumentID: GUID(_egg.DerivedFrom.DocumentID), InstanceID: GUID(_egg.DerivedFrom.InstanceID), VersionID: _egg.DerivedFrom.VersionID}
	}
	return _gbg, true
}

// PdfInfo is the xmp document pdf info.
type PdfInfo struct {
	InfoDict   _fea.PdfObject
	PdfVersion string
	Copyright  string
	Marked     bool
}

// PdfInfoOptions are the options used for setting pdf info.
type PdfInfoOptions struct {
	InfoDict   _fea.PdfObject
	PdfVersion string
	Copyright  string
	Marked     bool

	// Overwrite if set to true, overwrites all values found in the current pdf info xmp model to the ones provided.
	Overwrite bool
}

// MarshalIndent the document into xml byte stream with predefined prefix and indent.
func (_eb *Document) MarshalIndent(prefix, indent string) ([]byte, error) {
	if _eb._dg.IsDirty() {
		if _fd := _eb._dg.SyncModels(); _fd != nil {
			return nil, _fd
		}
	}
	return _g.MarshalIndent(_eb._dg, prefix, indent)
}
