package sighandler

import (
	_dc "bytes"
	_ga "crypto"
	_ca "crypto/rand"
	_a "crypto/rsa"
	_de "crypto/x509"
	_b "crypto/x509/pkix"
	_f "encoding/asn1"
	_cf "errors"
	_gb "fmt"
	_d "hash"
	_g "math/big"
	_e "time"

	_ce "bitbucket.org/shenghui0779/gopdf/core"
	_cefa "bitbucket.org/shenghui0779/gopdf/model"
	_cef "bitbucket.org/shenghui0779/gopdf/model/mdp"
	_eb "bitbucket.org/shenghui0779/gopdf/model/sigutil"
	_db "github.com/unidoc/pkcs7"
	_gf "github.com/unidoc/timestamp"
)

// NewDocTimeStamp creates a new DocTimeStamp signature handler.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
// NOTE: the handler will do a mock Sign when initializing the signature
// in order to estimate the signature size. Use NewDocTimeStampWithOpts
// for providing the signature size.
func NewDocTimeStamp(timestampServerURL string, hashAlgorithm _ga.Hash) (_cefa.SignatureHandler, error) {
	return &docTimeStamp{_fc: timestampServerURL, _fec: hashAlgorithm}, nil
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_gebb *adobeX509RSASHA1) IsApplicable(sig *_cefa.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031"
}
func (_dedd *adobeX509RSASHA1) sign(_aaa *_cefa.PdfSignature, _eee _cefa.Hasher, _feg bool) error {
	if !_feg {
		return _dedd.Sign(_aaa, _eee)
	}
	_fgf, _gfc := _dedd._aef.PublicKey.(*_a.PublicKey)
	if !_gfc {
		return _gb.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0075\u0062\u006c\u0069\u0063\u0020\u006b\u0065y\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", _fgf)
	}
	_ea, _afe := _f.Marshal(make([]byte, _fgf.Size()))
	if _afe != nil {
		return _afe
	}
	_aaa.Contents = _ce.MakeHexString(string(_ea))
	return nil
}

// InitSignature initialization of the DocMDP signature.
func (_eef *DocMDPHandler) InitSignature(sig *_cefa.PdfSignature) error {
	_gc := _eef._ab.InitSignature(sig)
	if _gc != nil {
		return _gc
	}
	sig.Handler = _eef
	if sig.Reference == nil {
		sig.Reference = _ce.MakeArray()
	}
	sig.Reference.Append(_cefa.NewPdfSignatureReferenceDocMDP(_cefa.NewPdfTransformParamsDocMDP(_eef.Permission)).ToPdfObject())
	return nil
}
func _gde(_fb _f.ObjectIdentifier) (_ga.Hash, error) {
	switch {
	case _fb.Equal(_db.OIDDigestAlgorithmSHA1), _fb.Equal(_db.OIDDigestAlgorithmECDSASHA1), _fb.Equal(_db.OIDDigestAlgorithmDSA), _fb.Equal(_db.OIDDigestAlgorithmDSASHA1), _fb.Equal(_db.OIDEncryptionAlgorithmRSA):
		return _ga.SHA1, nil
	case _fb.Equal(_db.OIDDigestAlgorithmSHA256), _fb.Equal(_db.OIDDigestAlgorithmECDSASHA256):
		return _ga.SHA256, nil
	case _fb.Equal(_db.OIDDigestAlgorithmSHA384), _fb.Equal(_db.OIDDigestAlgorithmECDSASHA384):
		return _ga.SHA384, nil
	case _fb.Equal(_db.OIDDigestAlgorithmSHA512), _fb.Equal(_db.OIDDigestAlgorithmECDSASHA512):
		return _ga.SHA512, nil
	}
	return _ga.Hash(0), _db.ErrUnsupportedAlgorithm
}

type docTimeStamp struct {
	_fc   string
	_fec  _ga.Hash
	_cagf int
}

func _gacg(_gdg []byte, _bfe int) (_gca []byte) {
	_cacf := len(_gdg)
	if _cacf > _bfe {
		_cacf = _bfe
	}
	_gca = make([]byte, _bfe)
	copy(_gca[len(_gca)-_cacf:], _gdg)
	return
}

// NewEmptyAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached
// signature handler. The generated signature is empty and of size signatureLen.
// The signatureLen parameter can be 0 for the signature validation.
func NewEmptyAdobePKCS7Detached(signatureLen int) (_cefa.SignatureHandler, error) {
	return &adobePKCS7Detached{_bac: true, _baf: signatureLen}, nil
}

// InitSignature initialises the PdfSignature.
func (_fecg *docTimeStamp) InitSignature(sig *_cefa.PdfSignature) error {
	_daab := *_fecg
	sig.Handler = &_daab
	sig.Filter = _ce.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _ce.MakeName("\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031")
	sig.Reference = nil
	if _fecg._cagf > 0 {
		sig.Contents = _ce.MakeHexString(string(make([]byte, _fecg._cagf)))
	} else {
		_ffa, _gfg := _fecg.NewDigest(sig)
		if _gfg != nil {
			return _gfg
		}
		_ffa.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
		if _gfg = _daab.Sign(sig, _ffa); _gfg != nil {
			return _gfg
		}
		_fecg._cagf = _daab._cagf
	}
	return nil
}

// Sign sets the Contents fields.
func (_af *adobePKCS7Detached) Sign(sig *_cefa.PdfSignature, digest _cefa.Hasher) error {
	if _af._bac {
		_dfd := _af._baf
		if _dfd <= 0 {
			_dfd = 8192
		}
		sig.Contents = _ce.MakeHexString(string(make([]byte, _dfd)))
		return nil
	}
	_bd := digest.(*_dc.Buffer)
	_aad, _ggf := _db.NewSignedData(_bd.Bytes())
	if _ggf != nil {
		return _ggf
	}
	if _gcg := _aad.AddSigner(_af._dd, _af._eefb, _db.SignerInfoConfig{}); _gcg != nil {
		return _gcg
	}
	_aad.Detach()
	_agf, _ggf := _aad.Finish()
	if _ggf != nil {
		return _ggf
	}
	_bbb := make([]byte, 8192)
	copy(_bbb, _agf)
	sig.Contents = _ce.MakeHexString(string(_bbb))
	return nil
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature
func (_cc *adobePKCS7Detached) IsApplicable(sig *_cefa.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064"
}

type adobePKCS7Detached struct {
	_eefb *_a.PrivateKey
	_dd   *_de.Certificate
	_bac  bool
	_baf  int
}

// ValidateWithOpts validates a PDF signature by checking PdfReader or PdfParser by the DiffPolicy
// params describes parameters for the DocMDP checks.
func (_ac *DocMDPHandler) ValidateWithOpts(sig *_cefa.PdfSignature, digest _cefa.Hasher, params _cefa.SignatureHandlerDocMDPParams) (_cefa.SignatureValidationResult, error) {
	_ada, _ec := _ac._ab.Validate(sig, digest)
	if _ec != nil {
		return _ada, _ec
	}
	_ceg := params.Parser
	if _ceg == nil {
		return _cefa.SignatureValidationResult{}, _cf.New("p\u0061r\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027t\u0020\u0062\u0065\u0020nu\u006c\u006c")
	}
	if !_ada.IsVerified {
		return _ada, nil
	}
	_gaf := params.DiffPolicy
	if _gaf == nil {
		_gaf = _cef.NewDefaultDiffPolicy()
	}
	for _fab := 0; _fab <= _ceg.GetRevisionNumber(); _fab++ {
		_df, _ag := _ceg.GetRevision(_fab)
		if _ag != nil {
			return _cefa.SignatureValidationResult{}, _ag
		}
		_cdf := _df.GetTrailer()
		if _cdf == nil {
			return _cefa.SignatureValidationResult{}, _cf.New("\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064\u0020\u0074\u0068\u0065\u0020\u0074r\u0061i\u006c\u0065\u0072\u0020\u006f\u0062\u006ae\u0063\u0074")
		}
		_fef, _bcg := _ce.GetDict(_cdf.Get("\u0052\u006f\u006f\u0074"))
		if !_bcg {
			return _cefa.SignatureValidationResult{}, _cf.New("\u0075n\u0064\u0065\u0066\u0069n\u0065\u0064\u0020\u0074\u0068e\u0020r\u006fo\u0074\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
		_gab, _bcg := _ce.GetDict(_fef.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
		if !_bcg {
			continue
		}
		_ee, _bcg := _ce.GetArray(_gab.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if !_bcg {
			continue
		}
		for _, _gg := range _ee.Elements() {
			_fefe, _ffg := _ce.GetDict(_gg)
			if !_ffg {
				continue
			}
			_ebg, _ffg := _ce.GetDict(_fefe.Get("\u0056"))
			if !_ffg {
				continue
			}
			if _ce.EqualObjects(_ebg.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"), sig.Contents) {
				_ada.DiffResults, _ag = _gaf.ReviewFile(_df, _ceg, &_cef.MDPParameters{DocMDPLevel: _ac.Permission})
				if _ag != nil {
					return _cefa.SignatureValidationResult{}, _ag
				}
				_ada.IsVerified = _ada.DiffResults.IsPermitted()
				return _ada, nil
			}
		}
	}
	return _cefa.SignatureValidationResult{}, _cf.New("\u0064\u006f\u006e\u0027\u0074\u0020\u0066o\u0075\u006e\u0064 \u0074\u0068\u0069\u0073 \u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073")
}

// NewAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobePKCS7Detached(privateKey *_a.PrivateKey, certificate *_de.Certificate) (_cefa.SignatureHandler, error) {
	return &adobePKCS7Detached{_dd: certificate, _eefb: privateKey}, nil
}

// DocMDPHandler describes handler for the DocMDP realization.
type DocMDPHandler struct {
	_ab        _cefa.SignatureHandler
	Permission _cef.DocMDPPermission
}

// Sign adds a new reference to signature's references array.
func (_ae *DocMDPHandler) Sign(sig *_cefa.PdfSignature, digest _cefa.Hasher) error {
	return _ae._ab.Sign(sig, digest)
}

// NewAdobeX509RSASHA1 creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler. Both the private key and the
// certificate can be nil for the signature validation.
func NewAdobeX509RSASHA1(privateKey *_a.PrivateKey, certificate *_de.Certificate) (_cefa.SignatureHandler, error) {
	return &adobeX509RSASHA1{_aef: certificate, _ed: privateKey}, nil
}

// SignFunc represents a custom signing function. The function should return
// the computed signature.
type SignFunc func(_ddf *_cefa.PdfSignature, _egdf _cefa.Hasher) ([]byte, error)

// NewDocTimeStampWithOpts returns a new DocTimeStamp configured using the
// specified options. If no options are provided, default options will be used.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
func NewDocTimeStampWithOpts(timestampServerURL string, hashAlgorithm _ga.Hash, opts *DocTimeStampOpts) (_cefa.SignatureHandler, error) {
	if opts == nil {
		opts = &DocTimeStampOpts{}
	}
	if opts.SignatureSize <= 0 {
		opts.SignatureSize = 4192
	}
	return &docTimeStamp{_fc: timestampServerURL, _fec: hashAlgorithm, _cagf: opts.SignatureSize}, nil
}

// InitSignature initialises the PdfSignature.
func (_cbg *adobeX509RSASHA1) InitSignature(sig *_cefa.PdfSignature) error {
	if _cbg._aef == nil {
		return _cf.New("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c")
	}
	if _cbg._ed == nil && _cbg._ece == nil {
		return _cf.New("\u006d\u0075\u0073\u0074\u0020\u0070\u0072o\u0076\u0069\u0064e\u0020\u0065\u0069t\u0068\u0065r\u0020\u0061\u0020\u0070\u0072\u0069v\u0061te\u0020\u006b\u0065\u0079\u0020\u006f\u0072\u0020\u0061\u0020\u0073\u0069\u0067\u006e\u0069\u006e\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	}
	_adg := *_cbg
	sig.Handler = &_adg
	sig.Filter = _ce.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _ce.MakeName("\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031")
	sig.Cert = _ce.MakeString(string(_adg._aef.Raw))
	sig.Reference = nil
	_deb, _da := _adg.NewDigest(sig)
	if _da != nil {
		return _da
	}
	_deb.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
	return _adg.sign(sig, _deb, _cbg._aaf)
}

// NewAdobeX509RSASHA1CustomWithOpts creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. The
// handler is configured based on the provided options. If no options are
// provided, default options will be used. Both the certificate and the sign
// function can be nil for the signature validation.
func NewAdobeX509RSASHA1CustomWithOpts(certificate *_de.Certificate, signFunc SignFunc, opts *AdobeX509RSASHA1Opts) (_cefa.SignatureHandler, error) {
	if opts == nil {
		opts = &AdobeX509RSASHA1Opts{}
	}
	return &adobeX509RSASHA1{_aef: certificate, _ece: signFunc, _aaf: opts.EstimateSize, _edg: opts.Algorithm}, nil
}

const _ggg = _ga.SHA1

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_bfc *docTimeStamp) IsApplicable(sig *_cefa.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031"
}

// Validate validates PdfSignature.
func (_aa *adobePKCS7Detached) Validate(sig *_cefa.PdfSignature, digest _cefa.Hasher) (_cefa.SignatureValidationResult, error) {
	_bcgb := sig.Contents.Bytes()
	_cb, _ebc := _db.Parse(_bcgb)
	if _ebc != nil {
		return _cefa.SignatureValidationResult{}, _ebc
	}
	_cag := digest.(*_dc.Buffer)
	_cb.Content = _cag.Bytes()
	if _ebc = _cb.Verify(); _ebc != nil {
		return _cefa.SignatureValidationResult{}, _ebc
	}
	return _cefa.SignatureValidationResult{IsSigned: true, IsVerified: true}, nil
}

// NewDigest creates a new digest.
func (_gff *adobePKCS7Detached) NewDigest(sig *_cefa.PdfSignature) (_cefa.Hasher, error) {
	return _dc.NewBuffer(nil), nil
}

type timestampInfo struct {
	Version        int
	Policy         _f.RawValue
	MessageImprint struct {
		HashAlgorithm _b.AlgorithmIdentifier
		HashedMessage []byte
	}
	SerialNumber    _f.RawValue
	GeneralizedTime _e.Time
}

// NewDigest creates a new digest.
func (_fg *DocMDPHandler) NewDigest(sig *_cefa.PdfSignature) (_cefa.Hasher, error) {
	return _fg._ab.NewDigest(sig)
}

// NewDocMDPHandler returns the new DocMDP handler with the specific DocMDP restriction level.
func NewDocMDPHandler(handler _cefa.SignatureHandler, permission _cef.DocMDPPermission) (_cefa.SignatureHandler, error) {
	return &DocMDPHandler{_ab: handler, Permission: permission}, nil
}

// Sign sets the Contents fields for the PdfSignature.
func (_dbg *adobeX509RSASHA1) Sign(sig *_cefa.PdfSignature, digest _cefa.Hasher) error {
	var _aag []byte
	var _cga error
	if _dbg._ece != nil {
		_aag, _cga = _dbg._ece(sig, digest)
		if _cga != nil {
			return _cga
		}
	} else {
		_aae, _bbd := digest.(_d.Hash)
		if !_bbd {
			return _cf.New("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_geb := _ggg
		if _dbg._edg != 0 {
			_geb = _dbg._edg
		}
		_aag, _cga = _a.SignPKCS1v15(_ca.Reader, _dbg._ed, _geb, _aae.Sum(nil))
		if _cga != nil {
			return _cga
		}
	}
	_aag, _cga = _f.Marshal(_aag)
	if _cga != nil {
		return _cga
	}
	sig.Contents = _ce.MakeHexString(string(_aag))
	return nil
}

// NewAdobeX509RSASHA1Custom creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. Both the
// certificate and the sign function can be nil for the signature validation.
// NOTE: the handler will do a mock Sign when initializing the signature in
// order to estimate the signature size. Use NewAdobeX509RSASHA1CustomWithOpts
// for configuring the handler to estimate the signature size.
func NewAdobeX509RSASHA1Custom(certificate *_de.Certificate, signFunc SignFunc) (_cefa.SignatureHandler, error) {
	return &adobeX509RSASHA1{_aef: certificate, _ece: signFunc}, nil
}
func (_egb *adobePKCS7Detached) getCertificate(_egd *_cefa.PdfSignature) (*_de.Certificate, error) {
	if _egb._dd != nil {
		return _egb._dd, nil
	}
	_ceb, _ef := _egd.GetCerts()
	if _ef != nil {
		return nil, _ef
	}
	return _ceb[0], nil
}

// DocTimeStampOpts defines options for configuring the timestamp handler.
type DocTimeStampOpts struct {

	// SignatureSize is the estimated size of the signature contents in bytes.
	// If not provided, a default signature size of 4192 is used.
	// The signing process will report the model.ErrSignNotEnoughSpace error
	// if the estimated signature size is smaller than the actual size of the
	// signature.
	SignatureSize int
}

// Validate implementation of the SignatureHandler interface
// This check is impossible without checking the document's content.
// Please, use ValidateWithOpts with the PdfParser.
func (_adb *DocMDPHandler) Validate(sig *_cefa.PdfSignature, digest _cefa.Hasher) (_cefa.SignatureValidationResult, error) {
	return _cefa.SignatureValidationResult{}, _cf.New("i\u006d\u0070\u006f\u0073\u0073\u0069b\u006c\u0065\u0020\u0076\u0061\u006ci\u0064\u0061\u0074\u0069\u006f\u006e\u0020w\u0069\u0074\u0068\u006f\u0075\u0074\u0020\u0070\u0061\u0072s\u0065")
}

// Sign sets the Contents fields for the PdfSignature.
func (_dbed *docTimeStamp) Sign(sig *_cefa.PdfSignature, digest _cefa.Hasher) error {
	_bgf, _bea := _eb.NewTimestampRequest(digest.(*_dc.Buffer), &_gf.RequestOptions{Hash: _dbed._fec, Certificates: true})
	if _bea != nil {
		return _bea
	}
	_abc := _eb.NewTimestampClient()
	_gdb, _bea := _abc.GetEncodedToken(_dbed._fc, _bgf)
	if _bea != nil {
		return _bea
	}
	_eea := len(_gdb)
	if _dbed._cagf > 0 && _eea > _dbed._cagf {
		return _cefa.ErrSignNotEnoughSpace
	}
	if _eea > 0 {
		_dbed._cagf = _eea + 128
	}
	sig.Contents = _ce.MakeHexString(string(_gdb))
	return nil
}

// InitSignature initialises the PdfSignature.
func (_bg *adobePKCS7Detached) InitSignature(sig *_cefa.PdfSignature) error {
	if !_bg._bac {
		if _bg._dd == nil {
			return _cf.New("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c")
		}
		if _bg._eefb == nil {
			return _cf.New("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c")
		}
	}
	_cg := *_bg
	sig.Handler = &_cg
	sig.Filter = _ce.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _ce.MakeName("\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064")
	sig.Reference = nil
	_dbe, _ebb := _cg.NewDigest(sig)
	if _ebb != nil {
		return _ebb
	}
	_dbe.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
	return _cg.Sign(sig, _dbe)
}

// Validate validates PdfSignature.
func (_bed *docTimeStamp) Validate(sig *_cefa.PdfSignature, digest _cefa.Hasher) (_cefa.SignatureValidationResult, error) {
	_bfgb := sig.Contents.Bytes()
	_dfdg, _eac := _db.Parse(_bfgb)
	if _eac != nil {
		return _cefa.SignatureValidationResult{}, _eac
	}
	if _eac = _dfdg.Verify(); _eac != nil {
		return _cefa.SignatureValidationResult{}, _eac
	}
	var _fefa timestampInfo
	_, _eac = _f.Unmarshal(_dfdg.Content, &_fefa)
	if _eac != nil {
		return _cefa.SignatureValidationResult{}, _eac
	}
	_fefd, _eac := _gde(_fefa.MessageImprint.HashAlgorithm.Algorithm)
	if _eac != nil {
		return _cefa.SignatureValidationResult{}, _eac
	}
	_ace := _fefd.New()
	_ccbd := digest.(*_dc.Buffer)
	_ace.Write(_ccbd.Bytes())
	_ecf := _ace.Sum(nil)
	_aec := _cefa.SignatureValidationResult{IsSigned: true, IsVerified: _dc.Equal(_ecf, _fefa.MessageImprint.HashedMessage), GeneralizedTime: _fefa.GeneralizedTime}
	return _aec, nil
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_be *DocMDPHandler) IsApplicable(sig *_cefa.PdfSignature) bool {
	_abd := false
	for _, _cd := range sig.Reference.Elements() {
		if _bb, _fa := _ce.GetDict(_cd); _fa {
			if _ad, _dca := _ce.GetNameVal(_bb.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064")); _dca {
				if _ad != "\u0044\u006f\u0063\u004d\u0044\u0050" {
					return false
				}
				if _fe, _ba := _ce.GetDict(_bb.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073")); _ba {
					_, _bc := _ce.GetNumberAsInt64(_fe.Get("\u0050"))
					if _bc != nil {
						return false
					}
					_abd = true
					break
				}
			}
		}
	}
	return _abd && _be._ab.IsApplicable(sig)
}

// NewDigest creates a new digest.
func (_bdf *adobeX509RSASHA1) NewDigest(sig *_cefa.PdfSignature) (_cefa.Hasher, error) {
	if _gd, _gcga := _bdf.getHashAlgorithm(sig); _gd != 0 && _gcga == nil {
		return _gd.New(), nil
	}
	return _ggg.New(), nil
}
func (_bdg *docTimeStamp) getCertificate(_gcf *_cefa.PdfSignature) (*_de.Certificate, error) {
	_dbdb, _bfg := _gcf.GetCerts()
	if _bfg != nil {
		return nil, _bfg
	}
	return _dbdb[0], nil
}

// AdobeX509RSASHA1Opts defines options for configuring the adbe.x509.rsa_sha1
// signature handler.
type AdobeX509RSASHA1Opts struct {

	// EstimateSize specifies whether the size of the signature contents
	// should be estimated based on the modulus size of the public key
	// extracted from the signing certificate. If set to false, a mock Sign
	// call is made in order to estimate the size of the signature contents.
	EstimateSize bool

	// Algorithm specifies the algorithm used for performing signing.
	// If not specified, defaults to SHA1.
	Algorithm _ga.Hash
}

func (_dbf *adobeX509RSASHA1) getHashAlgorithm(_dff *_cefa.PdfSignature) (_ga.Hash, error) {
	_ccba, _bcb := _dbf.getCertificate(_dff)
	if _bcb != nil {
		if _dbf._edg != 0 {
			return _dbf._edg, nil
		}
		return _ggg, _bcb
	}
	if _dff.Contents != nil {
		_dbc := _dff.Contents.Bytes()
		var _dbcd []byte
		if _, _cea := _f.Unmarshal(_dbc, &_dbcd); _cea == nil {
			_gfd := _dbd(_ccba.PublicKey.(*_a.PublicKey), _dbcd)
			if _gfd > 0 {
				return _gfd, nil
			}
		}
	}
	if _dbf._edg != 0 {
		return _dbf._edg, nil
	}
	return _ggg, nil
}
func (_daa *adobeX509RSASHA1) getCertificate(_ege *_cefa.PdfSignature) (*_de.Certificate, error) {
	if _daa._aef != nil {
		return _daa._aef, nil
	}
	_cebd, _ccb := _ege.GetCerts()
	if _ccb != nil {
		return nil, _ccb
	}
	return _cebd[0], nil
}

// Validate validates PdfSignature.
func (_bf *adobeX509RSASHA1) Validate(sig *_cefa.PdfSignature, digest _cefa.Hasher) (_cefa.SignatureValidationResult, error) {
	_dbfb, _ge := _bf.getCertificate(sig)
	if _ge != nil {
		return _cefa.SignatureValidationResult{}, _ge
	}
	_fabe := sig.Contents.Bytes()
	var _cca []byte
	if _, _gbe := _f.Unmarshal(_fabe, &_cca); _gbe != nil {
		return _cefa.SignatureValidationResult{}, _gbe
	}
	_bdfb, _ded := digest.(_d.Hash)
	if !_ded {
		return _cefa.SignatureValidationResult{}, _cf.New("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_efa, _ := _bf.getHashAlgorithm(sig)
	if _efa == 0 {
		_efa = _ggg
	}
	if _cfb := _a.VerifyPKCS1v15(_dbfb.PublicKey.(*_a.PublicKey), _efa, _bdfb.Sum(nil), _cca); _cfb != nil {
		return _cefa.SignatureValidationResult{}, _cfb
	}
	return _cefa.SignatureValidationResult{IsSigned: true, IsVerified: true}, nil
}

type adobeX509RSASHA1 struct {
	_ed  *_a.PrivateKey
	_aef *_de.Certificate
	_ece SignFunc
	_aaf bool
	_edg _ga.Hash
}

// NewDigest creates a new digest.
func (_gge *docTimeStamp) NewDigest(sig *_cefa.PdfSignature) (_cefa.Hasher, error) {
	return _dc.NewBuffer(nil), nil
}
func _dbd(_dda *_a.PublicKey, _cdfb []byte) _ga.Hash {
	_gac := _dda.Size()
	if _gac != len(_cdfb) {
		return 0
	}
	_gdd := func(_fga *_g.Int, _afb *_a.PublicKey, _dfda *_g.Int) *_g.Int {
		_cbgb := _g.NewInt(int64(_afb.E))
		_fga.Exp(_dfda, _cbgb, _afb.N)
		return _fga
	}
	_cce := new(_g.Int).SetBytes(_cdfb)
	_cae := _gdd(new(_g.Int), _dda, _cce)
	_eba := _gacg(_cae.Bytes(), _gac)
	if _eba[0] != 0 || _eba[1] != 1 {
		return 0
	}
	_ebcb := []struct {
		Hash   _ga.Hash
		Prefix []byte
	}{{Hash: _ga.SHA1, Prefix: []byte{0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14}}, {Hash: _ga.SHA256, Prefix: []byte{0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20}}, {Hash: _ga.SHA384, Prefix: []byte{0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30}}, {Hash: _ga.SHA512, Prefix: []byte{0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40}}, {Hash: _ga.RIPEMD160, Prefix: []byte{0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14}}}
	for _, _eeb := range _ebcb {
		_dcc := _eeb.Hash.Size()
		_efd := len(_eeb.Prefix) + _dcc
		if _dc.Equal(_eba[_gac-_efd:_gac-_dcc], _eeb.Prefix) {
			return _eeb.Hash
		}
	}
	return 0
}
