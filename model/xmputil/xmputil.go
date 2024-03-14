package xmputil

import (
	_bd "errors"
	_fe "fmt"
	_d "strconv"
	_a "time"

	_aab "bitbucket.org/shenghui0779/gopdf/core"
	_ac "bitbucket.org/shenghui0779/gopdf/internal/timeutils"
	_aa "bitbucket.org/shenghui0779/gopdf/internal/uuid"
	_ba "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaextension"
	_g "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaid"
	_be "github.com/trimmer-io/go-xmp/models/pdf"
	_db "github.com/trimmer-io/go-xmp/models/xmp_mm"
	_b "github.com/trimmer-io/go-xmp/xmp"
)

// MediaManagementVersion is the version of the media management xmp metadata.
type MediaManagementVersion struct {
	VersionID  string
	ModifyDate _a.Time
	Comments   string
	Modifier   string
}

// Document is an implementation of the xmp document.
// It is a wrapper over go-xmp/xmp.Document that provides some Pdf predefined functionality.
type Document struct{ _fg *_b.Document }

// PdfInfoOptions are the options used for setting pdf info.
type PdfInfoOptions struct {
	InfoDict   _aab.PdfObject
	PdfVersion string
	Copyright  string
	Marked     bool

	// Overwrite if set to true, overwrites all values found in the current pdf info xmp model to the ones provided.
	Overwrite bool
}

// GetPdfInfo gets the document pdf info.
func (_dag *Document) GetPdfInfo() (*PdfInfo, bool) {
	_gb, _fgf := _dag._fg.FindModel(_be.NsPDF).(*_be.PDFInfo)
	if !_fgf {
		return nil, false
	}
	_ebd := PdfInfo{}
	var _ecd *_aab.PdfObjectDictionary
	_ebd.Copyright = _gb.Copyright
	_ebd.PdfVersion = _gb.PDFVersion
	_ebd.Marked = bool(_gb.Marked)
	_cgd := func(_dac string, _fa _aab.PdfObject) {
		if _ecd == nil {
			_ecd = _aab.MakeDict()
		}
		_ecd.Set(_aab.PdfObjectName(_dac), _fa)
	}
	if len(_gb.Title) > 0 {
		_cgd("\u0054\u0069\u0074l\u0065", _aab.MakeString(_gb.Title.Default()))
	}
	if len(_gb.Author) > 0 {
		_cgd("\u0041\u0075\u0074\u0068\u006f\u0072", _aab.MakeString(_gb.Author[0]))
	}
	if _gb.Keywords != "" {
		_cgd("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _aab.MakeString(_gb.Keywords))
	}
	if len(_gb.Subject) > 0 {
		_cgd("\u0053u\u0062\u006a\u0065\u0063\u0074", _aab.MakeString(_gb.Subject.Default()))
	}
	if _gb.Creator != "" {
		_cgd("\u0043r\u0065\u0061\u0074\u006f\u0072", _aab.MakeString(string(_gb.Creator)))
	}
	if _gb.Producer != "" {
		_cgd("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _aab.MakeString(string(_gb.Producer)))
	}
	if _gb.Trapped {
		_cgd("\u0054r\u0061\u0070\u0070\u0065\u0064", _aab.MakeName("\u0054\u0072\u0075\u0065"))
	}
	if !_gb.CreationDate.IsZero() {
		_cgd("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _aab.MakeString(_ac.FormatPdfTime(_gb.CreationDate.Value())))
	}
	if !_gb.ModifyDate.IsZero() {
		_cgd("\u004do\u0064\u0044\u0061\u0074\u0065", _aab.MakeString(_ac.FormatPdfTime(_gb.ModifyDate.Value())))
	}
	_ebd.InfoDict = _ecd
	return &_ebd, true
}

// LoadDocument loads up the xmp document from provided input stream.
func LoadDocument(stream []byte) (*Document, error) {
	_bg := _b.NewDocument()
	if _gd := _b.Unmarshal(stream, _bg); _gd != nil {
		return nil, _gd
	}
	return &Document{_fg: _bg}, nil
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
	ModifyDate _a.Time

	// Modifier is a person who did the modification.
	Modifier string
}

// GetGoXmpDocument gets direct access to the go-xmp.Document.
// All changes done to specified document would result in change of this document 'd'.
func (_fge *Document) GetGoXmpDocument() *_b.Document { return _fge._fg }

// GetPdfAID gets the pdfaid xmp metadata model.
func (_bc *Document) GetPdfAID() (*PdfAID, bool) {
	_cd, _eea := _bc._fg.FindModel(_g.Namespace).(*_g.Model)
	if !_eea {
		return nil, false
	}
	return &PdfAID{Part: _cd.Part, Conformance: _cd.Conformance}, true
}

// PdfInfo is the xmp document pdf info.
type PdfInfo struct {
	InfoDict   _aab.PdfObject
	PdfVersion string
	Copyright  string
	Marked     bool
}

// NewDocument creates a new document without any previous xmp information.
func NewDocument() *Document { _ff := _b.NewDocument(); return &Document{_fg: _ff} }

// SetPdfAID sets up pdfaid xmp metadata.
// In example: Part: '1' Conformance: 'B' states for PDF/A 1B.
func (_eba *Document) SetPdfAID(part int, conformance string) error {
	_fgfb, _cfd := _g.MakeModel(_eba._fg)
	if _cfd != nil {
		return _cfd
	}
	_fgfb.Part = part
	_fgfb.Conformance = conformance
	if _cgg := _fgfb.SyncToXMP(_eba._fg); _cgg != nil {
		return _cgg
	}
	return nil
}

// MediaManagementDerivedFrom is a structure that contains references of identifiers and versions
// from which given document was derived.
type MediaManagementDerivedFrom struct {
	OriginalDocumentID GUID
	DocumentID         GUID
	InstanceID         GUID
	VersionID          string
}

// Marshal the document into xml byte stream.
func (_e *Document) Marshal() ([]byte, error) {
	if _e._fg.IsDirty() {
		if _fgd := _e._fg.SyncModels(); _fgd != nil {
			return nil, _fgd
		}
	}
	return _b.Marshal(_e._fg)
}

// GetMediaManagement gets the media management metadata from provided xmp document.
func (_bda *Document) GetMediaManagement() (*MediaManagement, bool) {
	_gc := _db.FindModel(_bda._fg)
	if _gc == nil {
		return nil, false
	}
	_aca := make([]MediaManagementVersion, len(_gc.Versions))
	for _eef, _bdd := range _gc.Versions {
		_aca[_eef] = MediaManagementVersion{VersionID: _bdd.Version, ModifyDate: _bdd.ModifyDate.Value(), Comments: _bdd.Comments, Modifier: _bdd.Modifier}
	}
	_fcfd := &MediaManagement{OriginalDocumentID: GUID(_gc.OriginalDocumentID.Value()), DocumentID: GUID(_gc.DocumentID.Value()), InstanceID: GUID(_gc.InstanceID.Value()), VersionID: _gc.VersionID, Versions: _aca}
	if _gc.DerivedFrom != nil {
		_fcfd.DerivedFrom = &MediaManagementDerivedFrom{OriginalDocumentID: GUID(_gc.DerivedFrom.OriginalDocumentID), DocumentID: GUID(_gc.DerivedFrom.DocumentID), InstanceID: GUID(_gc.DerivedFrom.InstanceID), VersionID: _gc.DerivedFrom.VersionID}
	}
	return _fcfd, true
}

// SetPdfInfo sets the pdf info into selected document.
func (_ga *Document) SetPdfInfo(options *PdfInfoOptions) error {
	if options == nil {
		return _bd.New("\u006ei\u006c\u0020\u0070\u0064\u0066\u0020\u006f\u0070\u0074\u0069\u006fn\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_efd, _cb := _be.MakeModel(_ga._fg)
	if _cb != nil {
		return _cb
	}
	if options.Overwrite {
		*_efd = _be.PDFInfo{}
	}
	if options.InfoDict != nil {
		_cf, _aac := _aab.GetDict(options.InfoDict)
		if !_aac {
			return _fe.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", options.InfoDict)
		}
		var _dab *_aab.PdfObjectString
		for _, _af := range _cf.Keys() {
			switch _af {
			case "\u0054\u0069\u0074l\u0065":
				_dab, _aac = _aab.GetString(_cf.Get("\u0054\u0069\u0074l\u0065"))
				if _aac {
					_efd.Title = _b.NewAltString(_dab)
				}
			case "\u0041\u0075\u0074\u0068\u006f\u0072":
				_dab, _aac = _aab.GetString(_cf.Get("\u0041\u0075\u0074\u0068\u006f\u0072"))
				if _aac {
					_efd.Author = _b.NewStringList(_dab.String())
				}
			case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":
				_dab, _aac = _aab.GetString(_cf.Get("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"))
				if _aac {
					_efd.Keywords = _dab.String()
				}
			case "\u0043r\u0065\u0061\u0074\u006f\u0072":
				_dab, _aac = _aab.GetString(_cf.Get("\u0043r\u0065\u0061\u0074\u006f\u0072"))
				if _aac {
					_efd.Creator = _b.AgentName(_dab.String())
				}
			case "\u0053u\u0062\u006a\u0065\u0063\u0074":
				_dab, _aac = _aab.GetString(_cf.Get("\u0053u\u0062\u006a\u0065\u0063\u0074"))
				if _aac {
					_efd.Subject = _b.NewAltString(_dab.String())
				}
			case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":
				_dab, _aac = _aab.GetString(_cf.Get("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"))
				if _aac {
					_efd.Producer = _b.AgentName(_dab.String())
				}
			case "\u0054r\u0061\u0070\u0070\u0065\u0064":
				_fc, _bb := _aab.GetName(_cf.Get("\u0054r\u0061\u0070\u0070\u0065\u0064"))
				if _bb {
					switch _fc.String() {
					case "\u0054\u0072\u0075\u0065":
						_efd.Trapped = true
					case "\u0046\u0061\u006cs\u0065":
						_efd.Trapped = false
					default:
						_efd.Trapped = true
					}
				}
			case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":
				if _fd, _ec := _aab.GetString(_cf.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065")); _ec && _fd.String() != "" {
					_gf, _ebe := _ac.ParsePdfTime(_fd.String())
					if _ebe != nil {
						return _fe.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077", _ebe)
					}
					_efd.CreationDate = _b.NewDate(_gf)
				}
			case "\u004do\u0064\u0044\u0061\u0074\u0065":
				if _gda, _fcd := _aab.GetString(_cf.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _fcd && _gda.String() != "" {
					_cg, _eg := _ac.ParsePdfTime(_gda.String())
					if _eg != nil {
						return _fe.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _eg)
					}
					_efd.ModifyDate = _b.NewDate(_cg)
				}
			}
		}
	}
	if options.PdfVersion != "" {
		_efd.PDFVersion = options.PdfVersion
	}
	if options.Marked {
		_efd.Marked = _b.Bool(options.Marked)
	}
	if options.Copyright != "" {
		_efd.Copyright = options.Copyright
	}
	if _cb = _efd.SyncToXMP(_ga._fg); _cb != nil {
		return _cb
	}
	return nil
}

// MarshalIndent the document into xml byte stream with predefined prefix and indent.
func (_aaf *Document) MarshalIndent(prefix, indent string) ([]byte, error) {
	if _aaf._fg.IsDirty() {
		if _c := _aaf._fg.SyncModels(); _c != nil {
			return nil, _c
		}
	}
	return _b.MarshalIndent(_aaf._fg, prefix, indent)
}

// SetPdfAExtension sets the pdfaExtension XMP metadata.
func (_ae *Document) SetPdfAExtension() error {
	_ef, _bf := _ba.MakeModel(_ae._fg)
	if _bf != nil {
		return _bf
	}
	if _bf = _ba.FillModel(_ae._fg, _ef); _bf != nil {
		return _bf
	}
	if _bf = _ef.SyncToXMP(_ae._fg); _bf != nil {
		return _bf
	}
	return nil
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

// PdfAID is the result of the XMP pdfaid metadata.
type PdfAID struct {
	Part        int
	Conformance string
}

// SetMediaManagement sets up XMP media management metadata: namespace xmpMM.
func (_fcf *Document) SetMediaManagement(options *MediaManagementOptions) error {
	_fde, _cbc := _db.MakeModel(_fcf._fg)
	if _cbc != nil {
		return _cbc
	}
	if options == nil {
		options = new(MediaManagementOptions)
	}
	_gba := _db.ResourceRef{}
	if _fde.OriginalDocumentID.IsZero() {
		if options.OriginalDocumentID != "" {
			_fde.OriginalDocumentID = _b.GUID(options.OriginalDocumentID)
		} else {
			_ebeg, _cc := _aa.NewUUID()
			if _cc != nil {
				return _cc
			}
			_fde.OriginalDocumentID = _b.GUID(_ebeg.String())
		}
	} else {
		_gba.OriginalDocumentID = _fde.OriginalDocumentID
	}
	switch {
	case options.DocumentID != "":
		_fde.DocumentID = _b.GUID(options.DocumentID)
	case options.NewDocumentID || _fde.DocumentID.IsZero():
		if !_fde.DocumentID.IsZero() {
			_gba.DocumentID = _fde.DocumentID
		}
		_bdc, _fcc := _aa.NewUUID()
		if _fcc != nil {
			return _fcc
		}
		_fde.DocumentID = _b.GUID(_bdc.String())
	}
	if !_fde.InstanceID.IsZero() {
		_gba.InstanceID = _fde.InstanceID
	}
	_fde.InstanceID = _b.GUID(options.InstanceID)
	if _fde.InstanceID == "" {
		_dg, _bgd := _aa.NewUUID()
		if _bgd != nil {
			return _bgd
		}
		_fde.InstanceID = _b.GUID(_dg.String())
	}
	if !_gba.IsZero() {
		_fde.DerivedFrom = &_gba
	}
	_ad := options.VersionID
	if _fde.VersionID != "" {
		_gfe, _ca := _d.Atoi(_fde.VersionID)
		if _ca != nil {
			_ad = _d.Itoa(len(_fde.Versions) + 1)
		} else {
			_ad = _d.Itoa(_gfe + 1)
		}
	}
	if _ad == "" {
		_ad = "\u0031"
	}
	_fde.VersionID = _ad
	if _cbc = _fde.SyncToXMP(_fcf._fg); _cbc != nil {
		return _cbc
	}
	return nil
}

// GetPdfaExtensionSchemas gets a pdfa extension schemas.
func (_fec *Document) GetPdfaExtensionSchemas() ([]_ba.Schema, error) {
	_da := _fec._fg.FindModel(_ba.Namespace)
	if _da == nil {
		return nil, nil
	}
	_ee, _ea := _da.(*_ba.Model)
	if !_ea {
		return nil, _fe.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u006d\u006f\u0064\u0065l f\u006fr \u0070\u0064\u0066\u0061\u0045\u0078\u0074en\u0073\u0069\u006f\u006e\u0073\u003a\u0020%\u0054", _da)
	}
	return _ee.Schemas, nil
}

// GUID is a string representing a globally unique identifier.
type GUID string
