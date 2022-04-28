package xmputil

import (
	_gg "errors"
	_bce "fmt"
	_ga "strconv"
	_e "time"

	_dd "bitbucket.org/shenghui0779/gopdf/core"
	_gac "bitbucket.org/shenghui0779/gopdf/internal/timeutils"
	_gc "bitbucket.org/shenghui0779/gopdf/internal/uuid"
	_bc "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaextension"
	_ef "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaid"
	_b "github.com/trimmer-io/go-xmp/models/pdf"
	_db "github.com/trimmer-io/go-xmp/models/xmp_mm"
	_g "github.com/trimmer-io/go-xmp/xmp"
)

// MarshalIndent the document into xml byte stream with predefined prefix and indent.
func (_cb *Document) MarshalIndent(prefix, indent string) ([]byte, error) {
	if _cb._c.IsDirty() {
		if _cc := _cb._c.SyncModels(); _cc != nil {
			return nil, _cc
		}
	}
	return _g.MarshalIndent(_cb._c, prefix, indent)
}

// MediaManagementVersion is the version of the media management xmp metadata.
type MediaManagementVersion struct {
	VersionID  string
	ModifyDate _e.Time
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

// Marshal the document into xml byte stream.
func (_bcf *Document) Marshal() ([]byte, error) {
	if _bcf._c.IsDirty() {
		if _dc := _bcf._c.SyncModels(); _dc != nil {
			return nil, _dc
		}
	}
	return _g.Marshal(_bcf._c)
}

// GetPdfaExtensionSchemas gets a pdfa extension schemas.
func (_ca *Document) GetPdfaExtensionSchemas() ([]_bc.Schema, error) {
	_bca := _ca._c.FindModel(_bc.Namespace)
	if _bca == nil {
		return nil, nil
	}
	_ddf, _fa := _bca.(*_bc.Model)
	if !_fa {
		return nil, _bce.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u006d\u006f\u0064\u0065l f\u006fr \u0070\u0064\u0066\u0061\u0045\u0078\u0074en\u0073\u0069\u006f\u006e\u0073\u003a\u0020%\u0054", _bca)
	}
	return _ddf.Schemas, nil
}

// SetMediaManagement sets up XMP media management metadata: namespace xmpMM.
func (_fef *Document) SetMediaManagement(options *MediaManagementOptions) error {
	_ddfb, _ee := _db.MakeModel(_fef._c)
	if _ee != nil {
		return _ee
	}
	if options == nil {
		options = new(MediaManagementOptions)
	}
	_def := _db.ResourceRef{}
	if _ddfb.OriginalDocumentID.IsZero() {
		if options.OriginalDocumentID != "" {
			_ddfb.OriginalDocumentID = _g.GUID(options.OriginalDocumentID)
		} else {
			_gd, _geb := _gc.NewUUID()
			if _geb != nil {
				return _geb
			}
			_ddfb.OriginalDocumentID = _g.GUID(_gd.String())
		}
	} else {
		_def.OriginalDocumentID = _ddfb.OriginalDocumentID
	}
	switch {
	case options.DocumentID != "":
		_ddfb.DocumentID = _g.GUID(options.DocumentID)
	case options.NewDocumentID || _ddfb.DocumentID.IsZero():
		if !_ddfb.DocumentID.IsZero() {
			_def.DocumentID = _ddfb.DocumentID
		}
		_dda, _aa := _gc.NewUUID()
		if _aa != nil {
			return _aa
		}
		_ddfb.DocumentID = _g.GUID(_dda.String())
	}
	if !_ddfb.InstanceID.IsZero() {
		_def.InstanceID = _ddfb.InstanceID
	}
	_ddfb.InstanceID = _g.GUID(options.InstanceID)
	if _ddfb.InstanceID == "" {
		_bfb, _eea := _gc.NewUUID()
		if _eea != nil {
			return _eea
		}
		_ddfb.InstanceID = _g.GUID(_bfb.String())
	}
	if !_def.IsZero() {
		_ddfb.DerivedFrom = &_def
	}
	_ba := options.VersionID
	if _ddfb.VersionID != "" {
		_da, _bcc := _ga.Atoi(_ddfb.VersionID)
		if _bcc != nil {
			_ba = _ga.Itoa(len(_ddfb.Versions) + 1)
		} else {
			_ba = _ga.Itoa(_da + 1)
		}
	}
	if _ba == "" {
		_ba = "\u0031"
	}
	_ddfb.VersionID = _ba
	_bb := options.ModifyDate
	if _bb.IsZero() {
		_bb = _e.Now()
	}
	if _ee = _ddfb.SyncToXMP(_fef._c); _ee != nil {
		return _ee
	}
	return nil
}

// PdfAID is the result of the XMP pdfaid metadata.
type PdfAID struct {
	Part        int
	Conformance string
}

// LoadDocument loads up the xmp document from provided input stream.
func LoadDocument(stream []byte) (*Document, error) {
	_f := _g.NewDocument()
	if _ge := _g.Unmarshal(stream, _f); _ge != nil {
		return nil, _ge
	}
	return &Document{_c: _f}, nil
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
	ModifyDate _e.Time

	// Modifier is a person who did the modification.
	Modifier string
}

// SetPdfAExtension sets the pdfaExtension XMP metadata.
func (_dcb *Document) SetPdfAExtension() error {
	_fc, _ce := _bc.MakeModel(_dcb._c)
	if _ce != nil {
		return _ce
	}
	if _ce = _bc.FillModel(_dcb._c, _fc); _ce != nil {
		return _ce
	}
	if _ce = _fc.SyncToXMP(_dcb._c); _ce != nil {
		return _ce
	}
	return nil
}

// SetPdfAID sets up pdfaid xmp metadata.
// In example: Part: '1' Conformance: 'B' states for PDF/A 1B.
func (_ac *Document) SetPdfAID(part int, conformance string) error {
	_gag, _bg := _ef.MakeModel(_ac._c)
	if _bg != nil {
		return _bg
	}
	_gag.Part = part
	_gag.Conformance = conformance
	if _eag := _gag.SyncToXMP(_ac._c); _eag != nil {
		return _eag
	}
	return nil
}

// NewDocument creates a new document without any previous xmp information.
func NewDocument() *Document { _a := _g.NewDocument(); return &Document{_c: _a} }

// GetMediaManagement gets the media management metadata from provided xmp document.
func (_ddae *Document) GetMediaManagement() (*MediaManagement, bool) {
	_eeg := _db.FindModel(_ddae._c)
	if _eeg == nil {
		return nil, false
	}
	_bab := make([]MediaManagementVersion, len(_eeg.Versions))
	for _aba, _aab := range _eeg.Versions {
		_bab[_aba] = MediaManagementVersion{VersionID: _aab.Version, ModifyDate: _aab.ModifyDate.Value(), Comments: _aab.Comments, Modifier: _aab.Modifier}
	}
	_ddcb := &MediaManagement{OriginalDocumentID: GUID(_eeg.OriginalDocumentID.Value()), DocumentID: GUID(_eeg.DocumentID.Value()), InstanceID: GUID(_eeg.InstanceID.Value()), VersionID: _eeg.VersionID, Versions: _bab}
	if _eeg.DerivedFrom != nil {
		_ddcb.DerivedFrom = &MediaManagementDerivedFrom{OriginalDocumentID: GUID(_eeg.DerivedFrom.OriginalDocumentID), DocumentID: GUID(_eeg.DerivedFrom.DocumentID), InstanceID: GUID(_eeg.DerivedFrom.InstanceID), VersionID: _eeg.DerivedFrom.VersionID}
	}
	return _ddcb, true
}

// GUID is a string representing a globally unique identifier.
type GUID string

// Document is an implementation of the xmp document.
// It is a wrapper over go-xmp/xmp.Document that provides some Pdf predefined functionality.
type Document struct{ _c *_g.Document }

// PdfInfoOptions are the options used for setting pdf info.
type PdfInfoOptions struct {
	InfoDict   _dd.PdfObject
	PdfVersion string
	Copyright  string
	Marked     bool

	// Overwrite if set to true, overwrites all values found in the current pdf info xmp model to the ones provided.
	Overwrite bool
}

// PdfInfo is the xmp document pdf info.
type PdfInfo struct {
	InfoDict   _dd.PdfObject
	PdfVersion string
	Copyright  string
	Marked     bool
}

// GetPdfAID gets the pdfaid xmp metadata model.
func (_gaa *Document) GetPdfAID() (*PdfAID, bool) {
	_adc, _cce := _gaa._c.FindModel(_ef.Namespace).(*_ef.Model)
	if !_cce {
		return nil, false
	}
	return &PdfAID{Part: _adc.Part, Conformance: _adc.Conformance}, true
}

// SetPdfInfo sets the pdf info into selected document.
func (_ea *Document) SetPdfInfo(options *PdfInfoOptions) error {
	if options == nil {
		return _gg.New("\u006ei\u006c\u0020\u0070\u0064\u0066\u0020\u006f\u0070\u0074\u0069\u006fn\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_dcd, _cg := _b.MakeModel(_ea._c)
	if _cg != nil {
		return _cg
	}
	if options.Overwrite {
		*_dcd = _b.PDFInfo{}
	}
	if options.InfoDict != nil {
		_bf, _cd := _dd.GetDict(options.InfoDict)
		if !_cd {
			return _bce.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", options.InfoDict)
		}
		var _gef *_dd.PdfObjectString
		for _, _fb := range _bf.Keys() {
			switch _fb {
			case "\u0054\u0069\u0074l\u0065":
				_gef, _cd = _dd.GetString(_bf.Get("\u0054\u0069\u0074l\u0065"))
				if _cd {
					_dcd.Title = _g.NewAltString(_gef)
				}
			case "\u0041\u0075\u0074\u0068\u006f\u0072":
				_gef, _cd = _dd.GetString(_bf.Get("\u0041\u0075\u0074\u0068\u006f\u0072"))
				if _cd {
					_dcd.Author = _g.NewStringList(_gef.String())
				}
			case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":
				_gef, _cd = _dd.GetString(_bf.Get("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"))
				if _cd {
					_dcd.Keywords = _gef.String()
				}
			case "\u0043r\u0065\u0061\u0074\u006f\u0072":
				_gef, _cd = _dd.GetString(_bf.Get("\u0043r\u0065\u0061\u0074\u006f\u0072"))
				if _cd {
					_dcd.Creator = _g.AgentName(_gef.String())
				}
			case "\u0053u\u0062\u006a\u0065\u0063\u0074":
				_gef, _cd = _dd.GetString(_bf.Get("\u0053u\u0062\u006a\u0065\u0063\u0074"))
				if _cd {
					_dcd.Subject = _g.NewAltString(_gef.String())
				}
			case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":
				_gef, _cd = _dd.GetString(_bf.Get("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"))
				if _cd {
					_dcd.Producer = _g.AgentName(_gef.String())
				}
			case "\u0054r\u0061\u0070\u0070\u0065\u0064":
				_cdf, _be := _dd.GetName(_bf.Get("\u0054r\u0061\u0070\u0070\u0065\u0064"))
				if _be {
					switch _cdf.String() {
					case "\u0054\u0072\u0075\u0065":
						_dcd.Trapped = true
					case "\u0046\u0061\u006cs\u0065":
						_dcd.Trapped = false
					default:
						_dcd.Trapped = true
					}
				}
			case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":
				if _fac, _dbg := _dd.GetString(_bf.Get("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065")); _dbg && _fac.String() != "" {
					_ab, _ced := _gac.ParsePdfTime(_fac.String())
					if _ced != nil {
						return _bce.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077", _ced)
					}
					_dcd.CreationDate = _g.NewDate(_ab)
				}
			case "\u004do\u0064\u0044\u0061\u0074\u0065":
				if _fbc, _ed := _dd.GetString(_bf.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _ed && _fbc.String() != "" {
					_fg, _eg := _gac.ParsePdfTime(_fbc.String())
					if _eg != nil {
						return _bce.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _eg)
					}
					_dcd.ModifyDate = _g.NewDate(_fg)
				}
			}
		}
	}
	if options.PdfVersion != "" {
		_dcd.PDFVersion = options.PdfVersion
	}
	if options.Marked {
		_dcd.Marked = _g.Bool(options.Marked)
	}
	if options.Copyright != "" {
		_dcd.Copyright = options.Copyright
	}
	if _cg = _dcd.SyncToXMP(_ea._c); _cg != nil {
		return _cg
	}
	return nil
}

// GetPdfInfo gets the document pdf info.
func (_abf *Document) GetPdfInfo() (*PdfInfo, bool) {
	_beg, _ad := _abf._c.FindModel(_b.NsPDF).(*_b.PDFInfo)
	if !_ad {
		return nil, false
	}
	_cf := PdfInfo{}
	var _feb *_dd.PdfObjectDictionary
	_cf.Copyright = _beg.Copyright
	_cf.PdfVersion = _beg.PDFVersion
	_cf.Marked = bool(_beg.Marked)
	_gf := func(_de string, _bd _dd.PdfObject) {
		if _feb == nil {
			_feb = _dd.MakeDict()
		}
		_feb.Set(_dd.PdfObjectName(_de), _bd)
	}
	if len(_beg.Title) > 0 {
		_gf("\u0054\u0069\u0074l\u0065", _dd.MakeString(_beg.Title.Default()))
	}
	if len(_beg.Author) > 0 {
		_gf("\u0041\u0075\u0074\u0068\u006f\u0072", _dd.MakeString(_beg.Author[0]))
	}
	if _beg.Keywords != "" {
		_gf("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", _dd.MakeString(_beg.Keywords))
	}
	if len(_beg.Subject) > 0 {
		_gf("\u0053u\u0062\u006a\u0065\u0063\u0074", _dd.MakeString(_beg.Subject.Default()))
	}
	if _beg.Creator != "" {
		_gf("\u0043r\u0065\u0061\u0074\u006f\u0072", _dd.MakeString(string(_beg.Creator)))
	}
	if _beg.Producer != "" {
		_gf("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072", _dd.MakeString(string(_beg.Producer)))
	}
	if _beg.Trapped {
		_gf("\u0054r\u0061\u0070\u0070\u0065\u0064", _dd.MakeName("\u0054\u0072\u0075\u0065"))
	}
	if !_beg.CreationDate.IsZero() {
		_gf("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065", _dd.MakeString(_gac.FormatPdfTime(_beg.CreationDate.Value())))
	}
	if !_beg.ModifyDate.IsZero() {
		_gf("\u004do\u0064\u0044\u0061\u0074\u0065", _dd.MakeString(_gac.FormatPdfTime(_beg.ModifyDate.Value())))
	}
	_cf.InfoDict = _feb
	return &_cf, true
}

// GetGoXmpDocument gets direct access to the go-xmp.Document.
// All changes done to specified document would result in change of this document 'd'.
func (_fe *Document) GetGoXmpDocument() *_g.Document { return _fe._c }

// MediaManagementDerivedFrom is a structure that contains references of identifiers and versions
// from which given document was derived.
type MediaManagementDerivedFrom struct {
	OriginalDocumentID GUID
	DocumentID         GUID
	InstanceID         GUID
	VersionID          string
}
