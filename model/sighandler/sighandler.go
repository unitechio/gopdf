package sighandler

import (
	_e "bytes"
	_dc "crypto"
	_fg "crypto/rand"
	_c "crypto/rsa"
	_g "crypto/x509"
	_fcb "crypto/x509/pkix"
	_fc "encoding/asn1"
	_bb "errors"
	_ba "fmt"
	_d "hash"
	_a "math/big"
	_f "time"

	_cg "bitbucket.org/shenghui0779/gopdf/core"
	_fa "bitbucket.org/shenghui0779/gopdf/model"
	_bbg "bitbucket.org/shenghui0779/gopdf/model/mdp"
	_ag "bitbucket.org/shenghui0779/gopdf/model/sigutil"
	_ce "github.com/unidoc/pkcs7"
	_be "github.com/unidoc/timestamp"
)

// InitSignature initialises the PdfSignature.
func (_eed *adobePKCS7Detached) InitSignature(sig *_fa.PdfSignature) error {
	if !_eed._beg {
		if _eed._ggg == nil {
			return _bb.New("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c")
		}
		if _eed._ed == nil {
			return _bb.New("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c")
		}
	}
	_abe := *_eed
	sig.Handler = &_abe
	sig.Filter = _cg.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _cg.MakeName("\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064")
	sig.Reference = nil
	_gbg, _bae := _abe.NewDigest(sig)
	if _bae != nil {
		return _bae
	}
	_gbg.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
	return _abe.Sign(sig, _gbg)
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature
func (_dba *adobePKCS7Detached) IsApplicable(sig *_fa.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064"
}

// SignFunc represents a custom signing function. The function should return
// the computed signature.
type SignFunc func(_edb *_fa.PdfSignature, _ad _fa.Hasher) ([]byte, error)

// Sign adds a new reference to signature's references array.
func (_cf *DocMDPHandler) Sign(sig *_fa.PdfSignature, digest _fa.Hasher) error {
	return _cf._aa.Sign(sig, digest)
}

// Validate validates PdfSignature.
func (_dfc *adobePKCS7Detached) Validate(sig *_fa.PdfSignature, digest _fa.Hasher) (_fa.SignatureValidationResult, error) {
	_edf := sig.Contents.Bytes()
	_dbg, _bgb := _ce.Parse(_edf)
	if _bgb != nil {
		return _fa.SignatureValidationResult{}, _bgb
	}
	_abc := digest.(*_e.Buffer)
	_dbg.Content = _abc.Bytes()
	if _bgb = _dbg.Verify(); _bgb != nil {
		return _fa.SignatureValidationResult{}, _bgb
	}
	return _fa.SignatureValidationResult{IsSigned: true, IsVerified: true}, nil
}

// DocTimeStampOpts defines options for configuring the timestamp handler.
type DocTimeStampOpts struct {

	// SignatureSize is the estimated size of the signature contents in bytes.
	// If not provided, a default signature size of 4192 is used.
	// The signing process will report the model.ErrSignNotEnoughSpace error
	// if the estimated signature size is smaller than the actual size of the
	// signature.
	SignatureSize int

	// Client is the timestamp client used to make the signature request.
	// If no client is provided, a default one is used.
	Client *_ag.TimestampClient
}

// NewDigest creates a new digest.
func (_gbc *adobePKCS7Detached) NewDigest(sig *_fa.PdfSignature) (_fa.Hasher, error) {
	return _e.NewBuffer(nil), nil
}

// NewDocTimeStamp creates a new DocTimeStamp signature handler.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
// NOTE: the handler will do a mock Sign when initializing the signature
// in order to estimate the signature size. Use NewDocTimeStampWithOpts
// for providing the signature size.
func NewDocTimeStamp(timestampServerURL string, hashAlgorithm _dc.Hash) (_fa.SignatureHandler, error) {
	return &docTimeStamp{_debg: timestampServerURL, _eca: hashAlgorithm}, nil
}

type docTimeStamp struct {
	_debg string
	_eca  _dc.Hash
	_dcc  int
	_acc  *_ag.TimestampClient
}

func _eag(_cda []byte, _cdag int) (_ebf []byte) {
	_caa := len(_cda)
	if _caa > _cdag {
		_caa = _cdag
	}
	_ebf = make([]byte, _cdag)
	copy(_ebf[len(_ebf)-_caa:], _cda)
	return
}

type timestampInfo struct {
	Version        int
	Policy         _fc.RawValue
	MessageImprint struct {
		HashAlgorithm _fcb.AlgorithmIdentifier
		HashedMessage []byte
	}
	SerialNumber    _fc.RawValue
	GeneralizedTime _f.Time
}

// InitSignature initialises the PdfSignature.
func (_fee *adobeX509RSASHA1) InitSignature(sig *_fa.PdfSignature) error {
	if _fee._dfg == nil {
		return _bb.New("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c")
	}
	if _fee._gaf == nil && _fee._ead == nil {
		return _bb.New("\u006d\u0075\u0073\u0074\u0020\u0070\u0072o\u0076\u0069\u0064e\u0020\u0065\u0069t\u0068\u0065r\u0020\u0061\u0020\u0070\u0072\u0069v\u0061te\u0020\u006b\u0065\u0079\u0020\u006f\u0072\u0020\u0061\u0020\u0073\u0069\u0067\u006e\u0069\u006e\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	}
	_ggd := *_fee
	sig.Handler = &_ggd
	sig.Filter = _cg.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _cg.MakeName("\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031")
	sig.Cert = _cg.MakeString(string(_ggd._dfg.Raw))
	sig.Reference = nil
	_ac, _bbc := _ggd.NewDigest(sig)
	if _bbc != nil {
		return _bbc
	}
	_ac.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
	return _ggd.sign(sig, _ac, _fee._gc)
}

// Sign sets the Contents fields.
func (_dgc *adobePKCS7Detached) Sign(sig *_fa.PdfSignature, digest _fa.Hasher) error {
	if _dgc._beg {
		_ff := _dgc._af
		if _ff <= 0 {
			_ff = 8192
		}
		sig.Contents = _cg.MakeHexString(string(make([]byte, _ff)))
		return nil
	}
	_bd := digest.(*_e.Buffer)
	_ggf, _ffb := _ce.NewSignedData(_bd.Bytes())
	if _ffb != nil {
		return _ffb
	}
	if _cfb := _ggf.AddSigner(_dgc._ggg, _dgc._ed, _ce.SignerInfoConfig{}); _cfb != nil {
		return _cfb
	}
	_ggf.Detach()
	_ge, _ffb := _ggf.Finish()
	if _ffb != nil {
		return _ffb
	}
	_gag := make([]byte, 8192)
	copy(_gag, _ge)
	sig.Contents = _cg.MakeHexString(string(_gag))
	return nil
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_ggc *docTimeStamp) IsApplicable(sig *_fa.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031"
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
	Algorithm _dc.Hash
}

func (_efc *adobeX509RSASHA1) sign(_ffbe *_fa.PdfSignature, _eeb _fa.Hasher, _edc bool) error {
	if !_edc {
		return _efc.Sign(_ffbe, _eeb)
	}
	_gbge, _egd := _efc._dfg.PublicKey.(*_c.PublicKey)
	if !_egd {
		return _ba.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0075\u0062\u006c\u0069\u0063\u0020\u006b\u0065y\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", _gbge)
	}
	_egg, _cfc := _fc.Marshal(make([]byte, _gbge.Size()))
	if _cfc != nil {
		return _cfc
	}
	_ffbe.Contents = _cg.MakeHexString(string(_egg))
	return nil
}

// InitSignature initialization of the DocMDP signature.
func (_fb *DocMDPHandler) InitSignature(sig *_fa.PdfSignature) error {
	_cad := _fb._aa.InitSignature(sig)
	if _cad != nil {
		return _cad
	}
	sig.Handler = _fb
	if sig.Reference == nil {
		sig.Reference = _cg.MakeArray()
	}
	sig.Reference.Append(_fa.NewPdfSignatureReferenceDocMDP(_fa.NewPdfTransformParamsDocMDP(_fb.Permission)).ToPdfObject())
	return nil
}

// Sign sets the Contents fields for the PdfSignature.
func (_fgfa *docTimeStamp) Sign(sig *_fa.PdfSignature, digest _fa.Hasher) error {
	_ae, _fbfd := _ag.NewTimestampRequest(digest.(*_e.Buffer), &_be.RequestOptions{Hash: _fgfa._eca, Certificates: true})
	if _fbfd != nil {
		return _fbfd
	}
	_cbc := _fgfa._acc
	if _cbc == nil {
		_cbc = _ag.NewTimestampClient()
	}
	_ceb, _fbfd := _cbc.GetEncodedToken(_fgfa._debg, _ae)
	if _fbfd != nil {
		return _fbfd
	}
	_faf := len(_ceb)
	if _fgfa._dcc > 0 && _faf > _fgfa._dcc {
		return _fa.ErrSignNotEnoughSpace
	}
	if _faf > 0 {
		_fgfa._dcc = _faf + 128
	}
	sig.Contents = _cg.MakeHexString(string(_ceb))
	return nil
}
func (_bdc *adobeX509RSASHA1) getHashAlgorithm(_beb *_fa.PdfSignature) (_dc.Hash, error) {
	_feef, _afc := _bdc.getCertificate(_beb)
	if _afc != nil {
		if _bdc._gff != 0 {
			return _bdc._gff, nil
		}
		return _de, _afc
	}
	if _beb.Contents != nil {
		_abcf := _beb.Contents.Bytes()
		var _fec []byte
		if _, _cde := _fc.Unmarshal(_abcf, &_fec); _cde == nil {
			_fbf := _fga(_feef.PublicKey.(*_c.PublicKey), _fec)
			if _fbf > 0 {
				return _fbf, nil
			}
		}
	}
	if _bdc._gff != 0 {
		return _bdc._gff, nil
	}
	return _de, nil
}

// NewEmptyAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached
// signature handler. The generated signature is empty and of size signatureLen.
// The signatureLen parameter can be 0 for the signature validation.
func NewEmptyAdobePKCS7Detached(signatureLen int) (_fa.SignatureHandler, error) {
	return &adobePKCS7Detached{_beg: true, _af: signatureLen}, nil
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_ef *DocMDPHandler) IsApplicable(sig *_fa.PdfSignature) bool {
	_df := false
	for _, _gf := range sig.Reference.Elements() {
		if _da, _bg := _cg.GetDict(_gf); _bg {
			if _ea, _dd := _cg.GetNameVal(_da.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064")); _dd {
				if _ea != "\u0044\u006f\u0063\u004d\u0044\u0050" {
					return false
				}
				if _dad, _gb := _cg.GetDict(_da.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073")); _gb {
					_, _fe := _cg.GetNumberAsInt64(_dad.Get("\u0050"))
					if _fe != nil {
						return false
					}
					_df = true
					break
				}
			}
		}
	}
	return _df && _ef._aa.IsApplicable(sig)
}

const _de = _dc.SHA1

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_dfb *adobeX509RSASHA1) IsApplicable(sig *_fa.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031"
}

// NewDigest creates a new digest.
func (_bga *adobeX509RSASHA1) NewDigest(sig *_fa.PdfSignature) (_fa.Hasher, error) {
	if _bbda, _cdb := _bga.getHashAlgorithm(sig); _bbda != 0 && _cdb == nil {
		return _bbda.New(), nil
	}
	return _de.New(), nil
}

// NewDocTimeStampWithOpts returns a new DocTimeStamp configured using the
// specified options. If no options are provided, default options will be used.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
func NewDocTimeStampWithOpts(timestampServerURL string, hashAlgorithm _dc.Hash, opts *DocTimeStampOpts) (_fa.SignatureHandler, error) {
	if opts == nil {
		opts = &DocTimeStampOpts{}
	}
	if opts.SignatureSize <= 0 {
		opts.SignatureSize = 4192
	}
	return &docTimeStamp{_debg: timestampServerURL, _eca: hashAlgorithm, _dcc: opts.SignatureSize, _acc: opts.Client}, nil
}
func (_ga *adobePKCS7Detached) getCertificate(_fcd *_fa.PdfSignature) (*_g.Certificate, error) {
	if _ga._ggg != nil {
		return _ga._ggg, nil
	}
	_cfg, _gd := _fcd.GetCerts()
	if _gd != nil {
		return nil, _gd
	}
	return _cfg[0], nil
}

type adobeX509RSASHA1 struct {
	_gaf *_c.PrivateKey
	_dfg *_g.Certificate
	_ead SignFunc
	_gc  bool
	_gff _dc.Hash
}

func _bde(_bad _fc.ObjectIdentifier) (_dc.Hash, error) {
	switch {
	case _bad.Equal(_ce.OIDDigestAlgorithmSHA1), _bad.Equal(_ce.OIDDigestAlgorithmECDSASHA1), _bad.Equal(_ce.OIDDigestAlgorithmDSA), _bad.Equal(_ce.OIDDigestAlgorithmDSASHA1), _bad.Equal(_ce.OIDEncryptionAlgorithmRSA):
		return _dc.SHA1, nil
	case _bad.Equal(_ce.OIDDigestAlgorithmSHA256), _bad.Equal(_ce.OIDDigestAlgorithmECDSASHA256):
		return _dc.SHA256, nil
	case _bad.Equal(_ce.OIDDigestAlgorithmSHA384), _bad.Equal(_ce.OIDDigestAlgorithmECDSASHA384):
		return _dc.SHA384, nil
	case _bad.Equal(_ce.OIDDigestAlgorithmSHA512), _bad.Equal(_ce.OIDDigestAlgorithmECDSASHA512):
		return _dc.SHA512, nil
	}
	return _dc.Hash(0), _ce.ErrUnsupportedAlgorithm
}

// NewAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobePKCS7Detached(privateKey *_c.PrivateKey, certificate *_g.Certificate) (_fa.SignatureHandler, error) {
	return &adobePKCS7Detached{_ggg: certificate, _ed: privateKey}, nil
}

// NewDigest creates a new digest.
func (_dbe *DocMDPHandler) NewDigest(sig *_fa.PdfSignature) (_fa.Hasher, error) {
	return _dbe._aa.NewDigest(sig)
}

// DocMDPHandler describes handler for the DocMDP realization.
type DocMDPHandler struct {
	_aa        _fa.SignatureHandler
	Permission _bbg.DocMDPPermission
}

func (_agc *docTimeStamp) getCertificate(_bef *_fa.PdfSignature) (*_g.Certificate, error) {
	_gfa, _dbaa := _bef.GetCerts()
	if _dbaa != nil {
		return nil, _dbaa
	}
	return _gfa[0], nil
}

// NewAdobeX509RSASHA1CustomWithOpts creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. The
// handler is configured based on the provided options. If no options are
// provided, default options will be used. Both the certificate and the sign
// function can be nil for the signature validation.
func NewAdobeX509RSASHA1CustomWithOpts(certificate *_g.Certificate, signFunc SignFunc, opts *AdobeX509RSASHA1Opts) (_fa.SignatureHandler, error) {
	if opts == nil {
		opts = &AdobeX509RSASHA1Opts{}
	}
	return &adobeX509RSASHA1{_dfg: certificate, _ead: signFunc, _gc: opts.EstimateSize, _gff: opts.Algorithm}, nil
}
func (_cgb *adobeX509RSASHA1) getCertificate(_fgd *_fa.PdfSignature) (*_g.Certificate, error) {
	if _cgb._dfg != nil {
		return _cgb._dfg, nil
	}
	_eb, _bbe := _fgd.GetCerts()
	if _bbe != nil {
		return nil, _bbe
	}
	return _eb[0], nil
}

// NewDigest creates a new digest.
func (_cb *docTimeStamp) NewDigest(sig *_fa.PdfSignature) (_fa.Hasher, error) {
	return _e.NewBuffer(nil), nil
}

// Validate implementation of the SignatureHandler interface
// This check is impossible without checking the document's content.
// Please, use ValidateWithOpts with the PdfParser.
func (_cfd *DocMDPHandler) Validate(sig *_fa.PdfSignature, digest _fa.Hasher) (_fa.SignatureValidationResult, error) {
	return _fa.SignatureValidationResult{}, _bb.New("i\u006d\u0070\u006f\u0073\u0073\u0069b\u006c\u0065\u0020\u0076\u0061\u006ci\u0064\u0061\u0074\u0069\u006f\u006e\u0020w\u0069\u0074\u0068\u006f\u0075\u0074\u0020\u0070\u0061\u0072s\u0065")
}
func _fga(_gdg *_c.PublicKey, _adb []byte) _dc.Hash {
	_dde := _gdg.Size()
	if _dde != len(_adb) {
		return 0
	}
	_efcc := func(_cgd *_a.Int, _fgb *_c.PublicKey, _gfb *_a.Int) *_a.Int {
		_cc := _a.NewInt(int64(_fgb.E))
		_cgd.Exp(_gfb, _cc, _fgb.N)
		return _cgd
	}
	_bbga := new(_a.Int).SetBytes(_adb)
	_cdbe := _efcc(new(_a.Int), _gdg, _bbga)
	_ffg := _eag(_cdbe.Bytes(), _dde)
	if _ffg[0] != 0 || _ffg[1] != 1 {
		return 0
	}
	_cga := []struct {
		Hash   _dc.Hash
		Prefix []byte
	}{{Hash: _dc.SHA1, Prefix: []byte{0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14}}, {Hash: _dc.SHA256, Prefix: []byte{0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20}}, {Hash: _dc.SHA384, Prefix: []byte{0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30}}, {Hash: _dc.SHA512, Prefix: []byte{0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40}}, {Hash: _dc.RIPEMD160, Prefix: []byte{0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14}}}
	for _, _fcg := range _cga {
		_ffa := _fcg.Hash.Size()
		_deb := len(_fcg.Prefix) + _ffa
		if _e.Equal(_ffg[_dde-_deb:_dde-_ffa], _fcg.Prefix) {
			return _fcg.Hash
		}
	}
	return 0
}

type adobePKCS7Detached struct {
	_ed  *_c.PrivateKey
	_ggg *_g.Certificate
	_beg bool
	_af  int
}

// NewDocMDPHandler returns the new DocMDP handler with the specific DocMDP restriction level.
func NewDocMDPHandler(handler _fa.SignatureHandler, permission _bbg.DocMDPPermission) (_fa.SignatureHandler, error) {
	return &DocMDPHandler{_aa: handler, Permission: permission}, nil
}

// Validate validates PdfSignature.
func (_dcd *docTimeStamp) Validate(sig *_fa.PdfSignature, digest _fa.Hasher) (_fa.SignatureValidationResult, error) {
	_bcb := sig.Contents.Bytes()
	_bfb, _ccd := _ce.Parse(_bcb)
	if _ccd != nil {
		return _fa.SignatureValidationResult{}, _ccd
	}
	if _ccd = _bfb.Verify(); _ccd != nil {
		return _fa.SignatureValidationResult{}, _ccd
	}
	var _bbb timestampInfo
	_, _ccd = _fc.Unmarshal(_bfb.Content, &_bbb)
	if _ccd != nil {
		return _fa.SignatureValidationResult{}, _ccd
	}
	_feg, _ccd := _bde(_bbb.MessageImprint.HashAlgorithm.Algorithm)
	if _ccd != nil {
		return _fa.SignatureValidationResult{}, _ccd
	}
	_gfe := _feg.New()
	_eda := digest.(*_e.Buffer)
	_gfe.Write(_eda.Bytes())
	_eff := _gfe.Sum(nil)
	_bcd := _fa.SignatureValidationResult{IsSigned: true, IsVerified: _e.Equal(_eff, _bbb.MessageImprint.HashedMessage), GeneralizedTime: _bbb.GeneralizedTime}
	return _bcd, nil
}

// InitSignature initialises the PdfSignature.
func (_ebe *docTimeStamp) InitSignature(sig *_fa.PdfSignature) error {
	_ebb := *_ebe
	sig.Handler = &_ebb
	sig.Filter = _cg.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _cg.MakeName("\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031")
	sig.Reference = nil
	if _ebe._dcc > 0 {
		sig.Contents = _cg.MakeHexString(string(make([]byte, _ebe._dcc)))
	} else {
		_fac, _ebea := _ebe.NewDigest(sig)
		if _ebea != nil {
			return _ebea
		}
		_fac.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
		if _ebea = _ebb.Sign(sig, _fac); _ebea != nil {
			return _ebea
		}
		_ebe._dcc = _ebb._dcc
	}
	return nil
}

// Sign sets the Contents fields for the PdfSignature.
func (_dca *adobeX509RSASHA1) Sign(sig *_fa.PdfSignature, digest _fa.Hasher) error {
	var _fgf []byte
	var _gba error
	if _dca._ead != nil {
		_fgf, _gba = _dca._ead(sig, digest)
		if _gba != nil {
			return _gba
		}
	} else {
		_ada, _bee := digest.(_d.Hash)
		if !_bee {
			return _bb.New("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_abf := _de
		if _dca._gff != 0 {
			_abf = _dca._gff
		}
		_fgf, _gba = _c.SignPKCS1v15(_fg.Reader, _dca._gaf, _abf, _ada.Sum(nil))
		if _gba != nil {
			return _gba
		}
	}
	_fgf, _gba = _fc.Marshal(_fgf)
	if _gba != nil {
		return _gba
	}
	sig.Contents = _cg.MakeHexString(string(_fgf))
	return nil
}

// NewAdobeX509RSASHA1 creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler. Both the private key and the
// certificate can be nil for the signature validation.
func NewAdobeX509RSASHA1(privateKey *_c.PrivateKey, certificate *_g.Certificate) (_fa.SignatureHandler, error) {
	return &adobeX509RSASHA1{_dfg: certificate, _gaf: privateKey}, nil
}

// Validate validates PdfSignature.
func (_cgbg *adobeX509RSASHA1) Validate(sig *_fa.PdfSignature, digest _fa.Hasher) (_fa.SignatureValidationResult, error) {
	_aca, _eae := _cgbg.getCertificate(sig)
	if _eae != nil {
		return _fa.SignatureValidationResult{}, _eae
	}
	_dbf := sig.Contents.Bytes()
	var _ace []byte
	if _, _cec := _fc.Unmarshal(_dbf, &_ace); _cec != nil {
		return _fa.SignatureValidationResult{}, _cec
	}
	_baee, _ged := digest.(_d.Hash)
	if !_ged {
		return _fa.SignatureValidationResult{}, _bb.New("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_bfe, _ := _cgbg.getHashAlgorithm(sig)
	if _bfe == 0 {
		_bfe = _de
	}
	if _fed := _c.VerifyPKCS1v15(_aca.PublicKey.(*_c.PublicKey), _bfe, _baee.Sum(nil), _ace); _fed != nil {
		return _fa.SignatureValidationResult{}, _fed
	}
	return _fa.SignatureValidationResult{IsSigned: true, IsVerified: true}, nil
}

// ValidateWithOpts validates a PDF signature by checking PdfReader or PdfParser by the DiffPolicy
// params describes parameters for the DocMDP checks.
func (_dg *DocMDPHandler) ValidateWithOpts(sig *_fa.PdfSignature, digest _fa.Hasher, params _fa.SignatureHandlerDocMDPParams) (_fa.SignatureValidationResult, error) {
	_ee, _ec := _dg._aa.Validate(sig, digest)
	if _ec != nil {
		return _ee, _ec
	}
	_dga := params.Parser
	if _dga == nil {
		return _fa.SignatureValidationResult{}, _bb.New("p\u0061r\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027t\u0020\u0062\u0065\u0020nu\u006c\u006c")
	}
	if !_ee.IsVerified {
		return _ee, nil
	}
	_aae := params.DiffPolicy
	if _aae == nil {
		_aae = _bbg.NewDefaultDiffPolicy()
	}
	for _bf := 0; _bf <= _dga.GetRevisionNumber(); _bf++ {
		_eg, _ca := _dga.GetRevision(_bf)
		if _ca != nil {
			return _fa.SignatureValidationResult{}, _ca
		}
		_ab := _eg.GetTrailer()
		if _ab == nil {
			return _fa.SignatureValidationResult{}, _bb.New("\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064\u0020\u0074\u0068\u0065\u0020\u0074r\u0061i\u006c\u0065\u0072\u0020\u006f\u0062\u006ae\u0063\u0074")
		}
		_fae, _dac := _cg.GetDict(_ab.Get("\u0052\u006f\u006f\u0074"))
		if !_dac {
			return _fa.SignatureValidationResult{}, _bb.New("\u0075n\u0064\u0065\u0066\u0069n\u0065\u0064\u0020\u0074\u0068e\u0020r\u006fo\u0074\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
		_bff, _dac := _cg.GetDict(_fae.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
		if !_dac {
			continue
		}
		_cd, _dac := _cg.GetArray(_bff.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if !_dac {
			continue
		}
		for _, _dab := range _cd.Elements() {
			_dbd, _dadd := _cg.GetDict(_dab)
			if !_dadd {
				continue
			}
			_aga, _dadd := _cg.GetDict(_dbd.Get("\u0056"))
			if !_dadd {
				continue
			}
			if _cg.EqualObjects(_aga.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"), sig.Contents) {
				_ee.DiffResults, _ca = _aae.ReviewFile(_eg, _dga, &_bbg.MDPParameters{DocMDPLevel: _dg.Permission})
				if _ca != nil {
					return _fa.SignatureValidationResult{}, _ca
				}
				_ee.IsVerified = _ee.DiffResults.IsPermitted()
				return _ee, nil
			}
		}
	}
	return _fa.SignatureValidationResult{}, _bb.New("\u0064\u006f\u006e\u0027\u0074\u0020\u0066o\u0075\u006e\u0064 \u0074\u0068\u0069\u0073 \u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073")
}

// NewAdobeX509RSASHA1Custom creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. Both the
// certificate and the sign function can be nil for the signature validation.
// NOTE: the handler will do a mock Sign when initializing the signature in
// order to estimate the signature size. Use NewAdobeX509RSASHA1CustomWithOpts
// for configuring the handler to estimate the signature size.
func NewAdobeX509RSASHA1Custom(certificate *_g.Certificate, signFunc SignFunc) (_fa.SignatureHandler, error) {
	return &adobeX509RSASHA1{_dfg: certificate, _ead: signFunc}, nil
}
