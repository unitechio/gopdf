package xmputil

import (
	_fg "errors"
	_gf "fmt"
	_g "strconv"
	_d "time"

	_ea "bitbucket.org/shenghui0779/gopdf/core"
	_bb "bitbucket.org/shenghui0779/gopdf/internal/timeutils"
	_a "bitbucket.org/shenghui0779/gopdf/internal/uuid"
	_dd "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaextension"
	_eg "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaid"
	_b "github.com/trimmer-io/go-xmp/models/pdf"
	_fe "github.com/trimmer-io/go-xmp/models/xmp_mm"
	_f "github.com/trimmer-io/go-xmp/xmp"
)

// GetPdfInfo gets the document pdf info.
func (_feg *Document) GetPdfInfo() (*PdfInfo, bool) {
	_eba, _ec := _feg._c.FindModel(_b.NsPDF).(*_b.PDFInfo)
	if !_ec {
		return nil, false
	}
	_eac := PdfInfo{}
	var _ead *_ea.PdfObjectDictionary
	_eac.Copyright = _eba.Copyright
	_eac.PdfVersion = _eba.PDFVersion
	_eac.Marked = bool(_eba.Marked)
	_ad := func(_acb string, _abg _ea.PdfObject) {
		if _ead == nil {
			_ead = _ea.MakeDict()
		}
		_ead.Set(_ea.PdfObjectName(_acb), _abg)
	}
	if len(_eba.Title) > 0 {
		_ad("\u0054\u0069\u0074l\u0065", _ea.MakeString(_eba.Title.Default()))
	}
	if len(_eba.Author) > 0 {
		_ad("\u0041\u0075\u0074\u0068\u006f\u0072", _ea.MakeString(_eba.Author[0]))
	}
	if _eba.Keywords != "" {
		_ad("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _ea.MakeString(_eba.Keywords))
	}
	if len(_eba.Subject) > 0 {
		_ad("\u0053u\u0062\u006a\u0065\u0063\u0074", _ea.MakeString(_eba.Subject.Default()))
	}
	if _eba.Creator != "" {
		_ad("\u0043r\u0065\u0061\u0074\u006f\u0072", _ea.MakeString(string(_eba.Creator)))
	}
	if _eba.Producer != "" {
		_ad("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _ea.MakeString(string(_eba.Producer)))
	}
	if _eba.Trapped {
		_ad("\u0054r\u0061\u0070\u0070\u0065\u0064", _ea.MakeName("\u0054\u0072\u0075\u0065"))
	}
	if !_eba.CreationDate.IsZero() {
		_ad("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _ea.MakeString(_bb.FormatPdfTime(_eba.CreationDate.Value())))
	}
	if !_eba.ModifyDate.IsZero() {
		_ad("\u004do\u0064\u0044\u0061\u0074\u0065", _ea.MakeString(_bb.FormatPdfTime(_eba.ModifyDate.Value())))
	}
	_eac.InfoDict = _ead
	return &_eac, true
}

// MarshalIndent the document into xml byte stream with predefined prefix and indent.
func (_gfe *Document) MarshalIndent(prefix, indent string) ([]byte, error) {
	if _gfe._c.IsDirty() {
		if _cf := _gfe._c.SyncModels(); _cf != nil {
			return nil, _cf
		}
	}
	return _f.MarshalIndent(_gfe._c, prefix, indent)
}

// MediaManagementVersion is the version of the media management xmp metadata.
type MediaManagementVersion struct {
	VersionID  string
	ModifyDate _d.Time
	Comments   string
	Modifier   string
}

// SetPdfAExtension sets the pdfaExtension XMP metadata.
func (_fa *Document) SetPdfAExtension() error {
	_ab, _gc := _dd.MakeModel(_fa._c)
	if _gc != nil {
		return _gc
	}
	if _gc = _dd.FillModel(_fa._c, _ab); _gc != nil {
		return _gc
	}
	if _gc = _ab.SyncToXMP(_fa._c); _gc != nil {
		return _gc
	}
	return nil
}

// Marshal the document into xml byte stream.
func (_bc *Document) Marshal() ([]byte, error) {
	if _bc._c.IsDirty() {
		if _be := _bc._c.SyncModels(); _be != nil {
			return nil, _be
		}
	}
	return _f.Marshal(_bc._c)
}

// PdfAID is the result of the XMP pdfaid metadata.
type PdfAID struct {
	Part        int
	Conformance string
}

// SetMediaManagement sets up XMP media management metadata: namespace xmpMM.
func (_aaf *Document) SetMediaManagement(options *MediaManagementOptions) error {
	_bbe, _cbf := _fe.MakeModel(_aaf._c)
	if _cbf != nil {
		return _cbf
	}
	if options == nil {
		options = new(MediaManagementOptions)
	}
	_abd := _fe.ResourceRef{}
	if _bbe.OriginalDocumentID.IsZero() {
		if options.OriginalDocumentID != "" {
			_bbe.OriginalDocumentID = _f.GUID(options.OriginalDocumentID)
		} else {
			_aba, _cdf := _a.NewUUID()
			if _cdf != nil {
				return _cdf
			}
			_bbe.OriginalDocumentID = _f.GUID(_aba.String())
		}
	} else {
		_abd.OriginalDocumentID = _bbe.OriginalDocumentID
	}
	switch {
	case options.DocumentID != "":
		_bbe.DocumentID = _f.GUID(options.DocumentID)
	case options.NewDocumentID || _bbe.DocumentID.IsZero():
		if !_bbe.DocumentID.IsZero() {
			_abd.DocumentID = _bbe.DocumentID
		}
		_ca, _gb := _a.NewUUID()
		if _gb != nil {
			return _gb
		}
		_bbe.DocumentID = _f.GUID(_ca.String())
	}
	if !_bbe.InstanceID.IsZero() {
		_abd.InstanceID = _bbe.InstanceID
	}
	_bbe.InstanceID = _f.GUID(options.InstanceID)
	if _bbe.InstanceID == "" {
		_ebc, _dgg := _a.NewUUID()
		if _dgg != nil {
			return _dgg
		}
		_bbe.InstanceID = _f.GUID(_ebc.String())
	}
	if !_abd.IsZero() {
		_bbe.DerivedFrom = &_abd
	}
	_fad := options.VersionID
	if _bbe.VersionID != "" {
		_ade, _bad := _g.Atoi(_bbe.VersionID)
		if _bad != nil {
			_fad = _g.Itoa(len(_bbe.Versions) + 1)
		} else {
			_fad = _g.Itoa(_ade + 1)
		}
	}
	if _fad == "" {
		_fad = "\u0031"
	}
	_bbe.VersionID = _fad
	if _cbf = _bbe.SyncToXMP(_aaf._c); _cbf != nil {
		return _cbf
	}
	return nil
}

// PdfInfo is the xmp document pdf info.
type PdfInfo struct {
	InfoDict   _ea.PdfObject
	PdfVersion string
	Copyright  string
	Marked     bool
}

// SetPdfAID sets up pdfaid xmp metadata.
// In example: Part: '1' Conformance: 'B' states for PDF/A 1B.
func (_fba *Document) SetPdfAID(part int, conformance string) error {
	_acd, _fc := _eg.MakeModel(_fba._c)
	if _fc != nil {
		return _fc
	}
	_acd.Part = part
	_acd.Conformance = conformance
	if _fea := _acd.SyncToXMP(_fba._c); _fea != nil {
		return _fea
	}
	return nil
}

// SetPdfInfo sets the pdf info into selected document.
func (_gfg *Document) SetPdfInfo(options *PdfInfoOptions) error {
	if options == nil {
		return _fg.New("\u006ei\u006c\u0020\u0070\u0064\u0066\u0020\u006f\u0070\u0074\u0069\u006fn\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_ege, _fgf := _b.MakeModel(_gfg._c)
	if _fgf != nil {
		return _fgf
	}
	if options.Overwrite {
		*_ege = _b.PDFInfo{}
	}
	if options.InfoDict != nil {
		_gad, _cb := _ea.GetDict(options.InfoDict)
		if !_cb {
			return _gf.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", options.InfoDict)
		}
		var _cc *_ea.PdfObjectString
		for _, _ccd := range _gad.Keys() {
			switch _ccd {
			case "\u0054\u0069\u0074l\u0065":
				_cc, _cb = _ea.GetString(_gad.Get("\u0054\u0069\u0074l\u0065"))
				if _cb {
					_ege.Title = _f.NewAltString(_cc)
				}
			case "\u0041\u0075\u0074\u0068\u006f\u0072":
				_cc, _cb = _ea.GetString(_gad.Get("\u0041\u0075\u0074\u0068\u006f\u0072"))
				if _cb {
					_ege.Author = _f.NewStringList(_cc.String())
				}
			case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":
				_cc, _cb = _ea.GetString(_gad.Get("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"))
				if _cb {
					_ege.Keywords = _cc.String()
				}
			case "\u0043r\u0065\u0061\u0074\u006f\u0072":
				_cc, _cb = _ea.GetString(_gad.Get("\u0043r\u0065\u0061\u0074\u006f\u0072"))
				if _cb {
					_ege.Creator = _f.AgentName(_cc.String())
				}
			case "\u0053u\u0062\u006a\u0065\u0063\u0074":
				_cc, _cb = _ea.GetString(_gad.Get("\u0053u\u0062\u006a\u0065\u0063\u0074"))
				if _cb {
					_ege.Subject = _f.NewAltString(_cc.String())
				}
			case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":
				_cc, _cb = _ea.GetString(_gad.Get("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"))
				if _cb {
					_ege.Producer = _f.AgentName(_cc.String())
				}
			case "\u0054r\u0061\u0070\u0070\u0065\u0064":
				_ac, _ae := _ea.GetName(_gad.Get("\u0054r\u0061\u0070\u0070\u0065\u0064"))
				if _ae {
					switch _ac.String() {
					case "\u0054\u0072\u0075\u0065":
						_ege.Trapped = true
					case "\u0046\u0061\u006cs\u0065":
						_ege.Trapped = false
					default:
						_ege.Trapped = true
					}
				}
			case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":
				if _gce, _eb := _ea.GetString(_gad.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065")); _eb && _gce.String() != "" {
					_gadf, _bd := _bb.ParsePdfTime(_gce.String())
					if _bd != nil {
						return _gf.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077", _bd)
					}
					_ege.CreationDate = _f.NewDate(_gadf)
				}
			case "\u004do\u0064\u0044\u0061\u0074\u0065":
				if _gcg, _fb := _ea.GetString(_gad.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _fb && _gcg.String() != "" {
					_gg, _aa := _bb.ParsePdfTime(_gcg.String())
					if _aa != nil {
						return _gf.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _aa)
					}
					_ege.ModifyDate = _f.NewDate(_gg)
				}
			}
		}
	}
	if options.PdfVersion != "" {
		_ege.PDFVersion = options.PdfVersion
	}
	if options.Marked {
		_ege.Marked = _f.Bool(options.Marked)
	}
	if options.Copyright != "" {
		_ege.Copyright = options.Copyright
	}
	if _fgf = _ege.SyncToXMP(_gfg._c); _fgf != nil {
		return _fgf
	}
	return nil
}

// GetGoXmpDocument gets direct access to the go-xmp.Document.
// All changes done to specified document would result in change of this document 'd'.
func (_bca *Document) GetGoXmpDocument() *_f.Document { return _bca._c }

// GetPdfaExtensionSchemas gets a pdfa extension schemas.
func (_gcf *Document) GetPdfaExtensionSchemas() ([]_dd.Schema, error) {
	_ba := _gcf._c.FindModel(_dd.Namespace)
	if _ba == nil {
		return nil, nil
	}
	_cd, _dg := _ba.(*_dd.Model)
	if !_dg {
		return nil, _gf.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u006d\u006f\u0064\u0065l f\u006fr \u0070\u0064\u0066\u0061\u0045\u0078\u0074en\u0073\u0069\u006f\u006e\u0073\u003a\u0020%\u0054", _ba)
	}
	return _cd.Schemas, nil
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
	ModifyDate _d.Time

	// Modifier is a person who did the modification.
	Modifier string
}

// MediaManagementDerivedFrom is a structure that contains references of identifiers and versions
// from which given document was derived.
type MediaManagementDerivedFrom struct {
	OriginalDocumentID GUID
	DocumentID         GUID
	InstanceID         GUID
	VersionID          string
}

// NewDocument creates a new document without any previous xmp information.
func NewDocument() *Document { _ga := _f.NewDocument(); return &Document{_c: _ga} }

// PdfInfoOptions are the options used for setting pdf info.
type PdfInfoOptions struct {
	InfoDict   _ea.PdfObject
	PdfVersion string
	Copyright  string
	Marked     bool

	// Overwrite if set to true, overwrites all values found in the current pdf info xmp model to the ones provided.
	Overwrite bool
}

// GUID is a string representing a globally unique identifier.
type GUID string

// GetPdfAID gets the pdfaid xmp metadata model.
func (_bbb *Document) GetPdfAID() (*PdfAID, bool) {
	_ag, _df := _bbb._c.FindModel(_eg.Namespace).(*_eg.Model)
	if !_df {
		return nil, false
	}
	return &PdfAID{Part: _ag.Part, Conformance: _ag.Conformance}, true
}

// Document is an implementation of the xmp document.
// It is a wrapper over go-xmp/xmp.Document that provides some Pdf predefined functionality.
type Document struct{ _c *_f.Document }

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

// GetMediaManagement gets the media management metadata from provided xmp document.
func (_eda *Document) GetMediaManagement() (*MediaManagement, bool) {
	_gac := _fe.FindModel(_eda._c)
	if _gac == nil {
		return nil, false
	}
	_eaf := make([]MediaManagementVersion, len(_gac.Versions))
	for _ccb, _fda := range _gac.Versions {
		_eaf[_ccb] = MediaManagementVersion{VersionID: _fda.Version, ModifyDate: _fda.ModifyDate.Value(), Comments: _fda.Comments, Modifier: _fda.Modifier}
	}
	_af := &MediaManagement{OriginalDocumentID: GUID(_gac.OriginalDocumentID.Value()), DocumentID: GUID(_gac.DocumentID.Value()), InstanceID: GUID(_gac.InstanceID.Value()), VersionID: _gac.VersionID, Versions: _eaf}
	if _gac.DerivedFrom != nil {
		_af.DerivedFrom = &MediaManagementDerivedFrom{OriginalDocumentID: GUID(_gac.DerivedFrom.OriginalDocumentID), DocumentID: GUID(_gac.DerivedFrom.DocumentID), InstanceID: GUID(_gac.DerivedFrom.InstanceID), VersionID: _gac.DerivedFrom.VersionID}
	}
	return _af, true
}

// LoadDocument loads up the xmp document from provided input stream.
func LoadDocument(stream []byte) (*Document, error) {
	_fd := _f.NewDocument()
	if _gff := _f.Unmarshal(stream, _fd); _gff != nil {
		return nil, _gff
	}
	return &Document{_c: _fd}, nil
}
