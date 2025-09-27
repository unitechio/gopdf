package sighandler

import (
	_dc "bytes"
	_ff "crypto"
	_ba "crypto/rand"
	_ca "crypto/rsa"
	_df "crypto/x509"
	_d "crypto/x509/pkix"
	_ad "encoding/asn1"
	_f "encoding/hex"
	_b "errors"
	_ce "fmt"
	_cc "hash"
	_g "math/big"
	_a "strings"
	_cd "time"

	_dd "unitechio/gopdf/gopdf/common"
	_da "unitechio/gopdf/gopdf/core"
	_fe "unitechio/gopdf/gopdf/model"
	_dg "unitechio/gopdf/gopdf/model/mdp"
	_e "unitechio/gopdf/gopdf/model/sigutil"
	_ga "github.com/unidoc/pkcs7"
	_dfa "github.com/unidoc/timestamp"
)

// Sign adds a new reference to signature's references array.
func (_bae *DocMDPHandler) Sign(sig *_fe.PdfSignature, digest _fe.Hasher) error {
	return _bae._ed.Sign(sig, digest)
}

func (_cafa *docTimeStamp) getCertificate(_beee *_fe.PdfSignature) (*_df.Certificate, error) {
	_fdcd, _bcddf := _beee.GetCerts()
	if _bcddf != nil {
		return nil, _bcddf
	}
	return _fdcd[0], nil
}

// NewEtsiPAdESLevelLT creates a new Adobe.PPKLite ETSI.CAdES.detached Level LT signature handler.
func NewEtsiPAdESLevelLT(privateKey *_ca.PrivateKey, certificate *_df.Certificate, caCert *_df.Certificate, certificateTimestampServerURL string, appender *_fe.PdfAppender) (_fe.SignatureHandler, error) {
	_cda := appender.Reader.DSS
	if _cda == nil {
		_cda = _fe.NewDSS()
	}
	if _fad := _cda.GenerateHashMaps(); _fad != nil {
		return nil, _fad
	}
	return &etsiPAdES{_ec: certificate, _db: privateKey, _aab: caCert, _eb: certificateTimestampServerURL, CertClient: _e.NewCertClient(), OCSPClient: _e.NewOCSPClient(), CRLClient: _e.NewCRLClient(), _be: appender, _gde: _cda}, nil
}

// Validate implementation of the SignatureHandler interface
// This check is impossible without checking the document's content.
// Please, use ValidateWithOpts with the PdfParser.
func (_cac *DocMDPHandler) Validate(sig *_fe.PdfSignature, digest _fe.Hasher) (_fe.SignatureValidationResult, error) {
	return _fe.SignatureValidationResult{}, _b.New("i\u006d\u0070\u006f\u0073\u0073\u0069b\u006c\u0065\u0020\u0076\u0061\u006ci\u0064\u0061\u0074\u0069\u006f\u006e\u0020w\u0069\u0074\u0068\u006f\u0075\u0074\u0020\u0070\u0061\u0072s\u0065")
}

// InitSignature initialises the PdfSignature.
func (_fgbb *docTimeStamp) InitSignature(sig *_fe.PdfSignature) error {
	_fegb := *_fgbb
	sig.Type = _da.MakeName("\u0044\u006f\u0063T\u0069\u006d\u0065\u0053\u0074\u0061\u006d\u0070")
	sig.Handler = &_fegb
	sig.Filter = _da.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _da.MakeName("\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031")
	sig.Reference = nil
	if _fgbb._dade > 0 {
		sig.Contents = _da.MakeHexString(string(make([]byte, _fgbb._dade)))
	} else {
		_cfcf, _efcf := _fgbb.NewDigest(sig)
		if _efcf != nil {
			return _efcf
		}
		_cfcf.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
		if _efcf = _fegb.Sign(sig, _cfcf); _efcf != nil {
			return _efcf
		}
		_fgbb._dade = _fegb._dade
	}
	return nil
}

// RevocationInfoArchival is OIDAttributeAdobeRevocation attribute.
type RevocationInfoArchival struct {
	Crl          []_ad.RawValue `asn1:"explicit,tag:0,optional"`
	Ocsp         []_ad.RawValue `asn1:"explicit,tag:1,optional"`
	OtherRevInfo []_ad.RawValue `asn1:"explicit,tag:2,optional"`
}

const _acfb = _ff.SHA1

// Validate validates PdfSignature.
func (_eecb *adobePKCS7Detached) Validate(sig *_fe.PdfSignature, digest _fe.Hasher) (_fe.SignatureValidationResult, error) {
	_cfc := sig.Contents.Bytes()
	_gdeb, _bfe := _ga.Parse(_cfc)
	if _bfe != nil {
		return _fe.SignatureValidationResult{}, _bfe
	}
	_eccc, _gcgg := digest.(*_dc.Buffer)
	if !_gcgg {
		return _fe.SignatureValidationResult{}, _ce.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_gdeb.Content = _eccc.Bytes()
	if _bfe = _gdeb.Verify(); _bfe != nil {
		return _fe.SignatureValidationResult{}, _bfe
	}
	return _fe.SignatureValidationResult{IsSigned: true, IsVerified: true}, nil
}

// NewAdobeX509RSASHA1 creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler. Both the private key and the
// certificate can be nil for the signature validation.
func NewAdobeX509RSASHA1(privateKey *_ca.PrivateKey, certificate *_df.Certificate) (_fe.SignatureHandler, error) {
	return &adobeX509RSASHA1{_dga: certificate, _cccd: privateKey}, nil
}

// NewDocMDPHandler returns the new DocMDP handler with the specific DocMDP restriction level.
func NewDocMDPHandler(handler _fe.SignatureHandler, permission _dg.DocMDPPermission) (_fe.SignatureHandler, error) {
	return &DocMDPHandler{_ed: handler, Permission: permission}, nil
}

// NewDocTimeStamp creates a new DocTimeStamp signature handler.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
// NOTE: the handler will do a mock Sign when initializing the signature
// in order to estimate the signature size. Use NewDocTimeStampWithOpts
// for providing the signature size.
func NewDocTimeStamp(timestampServerURL string, hashAlgorithm _ff.Hash) (_fe.SignatureHandler, error) {
	return &docTimeStamp{_cccf: timestampServerURL, _bea: hashAlgorithm}, nil
}

// Sign sets the Contents fields.
func (_defd *adobePKCS7Detached) Sign(sig *_fe.PdfSignature, digest _fe.Hasher) error {
	if _defd._bebe {
		_dad := _defd._baba
		if _dad <= 0 {
			_dad = 8192
		}
		sig.Contents = _da.MakeHexString(string(make([]byte, _dad)))
		return nil
	}
	_dab, _ceecc := digest.(*_dc.Buffer)
	if !_ceecc {
		return _ce.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_ecce, _ffbg := _ga.NewSignedData(_dab.Bytes())
	if _ffbg != nil {
		return _ffbg
	}
	if _aaf := _ecce.AddSigner(_defd._dce, _defd._eca, _ga.SignerInfoConfig{}); _aaf != nil {
		return _aaf
	}
	_ecce.Detach()
	_fea, _ffbg := _ecce.Finish()
	if _ffbg != nil {
		return _ffbg
	}
	_aeb := make([]byte, 8192)
	copy(_aeb, _fea)
	sig.Contents = _da.MakeHexString(string(_aeb))
	return nil
}

// Sign sets the Contents fields for the PdfSignature.
func (_ecbf *docTimeStamp) Sign(sig *_fe.PdfSignature, digest _fe.Hasher) error {
	_edg, _egfe := _e.NewTimestampRequest(digest.(*_dc.Buffer), &_dfa.RequestOptions{Hash: _ecbf._bea, Certificates: true})
	if _egfe != nil {
		return _egfe
	}
	_ffbb := _ecbf._fced
	if _ffbb == nil {
		_ffbb = _e.NewTimestampClient()
	}
	_ceff, _egfe := _ffbb.GetEncodedToken(_ecbf._cccf, _edg)
	if _egfe != nil {
		return _egfe
	}
	_ebae := len(_ceff)
	if _ecbf._dade > 0 && _ebae > _ecbf._dade {
		return _fe.ErrSignNotEnoughSpace
	}
	if _ebae > 0 {
		_ecbf._dade = _ebae + 128
	}
	if sig.Contents != nil {
		_aeg := sig.Contents.Bytes()
		copy(_aeg, _ceff)
		_ceff = _aeg
	}
	sig.Contents = _da.MakeHexString(string(_ceff))
	return nil
}

// NewDigest creates a new digest.
func (_gbeb *docTimeStamp) NewDigest(sig *_fe.PdfSignature) (_fe.Hasher, error) {
	return _dc.NewBuffer(nil), nil
}

// Validate validates PdfSignature.
func (_eac *adobeX509RSASHA1) Validate(sig *_fe.PdfSignature, digest _fe.Hasher) (_fe.SignatureValidationResult, error) {
	_cgcd, _eee := _eac.getCertificate(sig)
	if _eee != nil {
		return _fe.SignatureValidationResult{}, _eee
	}
	_bcd := sig.Contents.Bytes()
	var _ddac []byte
	if _, _cbga := _ad.Unmarshal(_bcd, &_ddac); _cbga != nil {
		return _fe.SignatureValidationResult{}, _cbga
	}
	_gca, _fef := digest.(_cc.Hash)
	if !_fef {
		return _fe.SignatureValidationResult{}, _b.New("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_daa, _ := _eac.getHashAlgorithm(sig)
	if _daa == 0 {
		_daa = _acfb
	}
	if _bfd := _ca.VerifyPKCS1v15(_cgcd.PublicKey.(*_ca.PublicKey), _daa, _gca.Sum(nil), _ddac); _bfd != nil {
		return _fe.SignatureValidationResult{}, _bfd
	}
	return _fe.SignatureValidationResult{IsSigned: true, IsVerified: true}, nil
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_ecd *etsiPAdES) IsApplicable(sig *_fe.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064"
}

func _cbgc(_dfd *_ca.PublicKey, _bcea []byte) _ff.Hash {
	_eab := _dfd.Size()
	if _eab != len(_bcea) {
		return 0
	}
	_dfce := func(_afae *_g.Int, _fge *_ca.PublicKey, _eba *_g.Int) *_g.Int {
		_abe := _g.NewInt(int64(_fge.E))
		_afae.Exp(_eba, _abe, _fge.N)
		return _afae
	}
	_bfdf := new(_g.Int).SetBytes(_bcea)
	_abb := _dfce(new(_g.Int), _dfd, _bfdf)
	_gcf := _eacc(_abb.Bytes(), _eab)
	if _gcf[0] != 0 || _gcf[1] != 1 {
		return 0
	}
	_cebg := []struct {
		Hash   _ff.Hash
		Prefix []byte
	}{{Hash: _ff.SHA1, Prefix: []byte{0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14}}, {Hash: _ff.SHA256, Prefix: []byte{0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20}}, {Hash: _ff.SHA384, Prefix: []byte{0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30}}, {Hash: _ff.SHA512, Prefix: []byte{0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40}}, {Hash: _ff.RIPEMD160, Prefix: []byte{0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14}}}
	for _, _fcd := range _cebg {
		_cga := _fcd.Hash.Size()
		_deb := len(_fcd.Prefix) + _cga
		if _dc.Equal(_gcf[_eab-_deb:_eab-_cga], _fcd.Prefix) {
			return _fcd.Hash
		}
	}
	return 0
}

// InitSignature initialises the PdfSignature.
func (_de *etsiPAdES) InitSignature(sig *_fe.PdfSignature) error {
	if !_de._acd {
		if _de._ec == nil {
			return _b.New("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c")
		}
		if _de._db == nil {
			return _b.New("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c")
		}
	}
	_ffc := *_de
	sig.Handler = &_ffc
	sig.Filter = _da.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _da.MakeName("\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064")
	sig.Reference = nil
	_ee, _cefb := _ffc.NewDigest(sig)
	if _cefb != nil {
		return _cefb
	}
	_, _cefb = _ee.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
	if _cefb != nil {
		return _cefb
	}
	_ffc._bab = true
	_cefb = _ffc.Sign(sig, _ee)
	_ffc._bab = false
	return _cefb
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_bfb *docTimeStamp) IsApplicable(sig *_fe.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031"
}

// SignFunc represents a custom signing function. The function should return
// the computed signature.
type SignFunc func(_cbc *_fe.PdfSignature, _bbf _fe.Hasher) ([]byte, error)

// NewEtsiPAdESLevelT creates a new Adobe.PPKLite ETSI.CAdES.detached Level T signature handler.
func NewEtsiPAdESLevelT(privateKey *_ca.PrivateKey, certificate *_df.Certificate, caCert *_df.Certificate, certificateTimestampServerURL string) (_fe.SignatureHandler, error) {
	return &etsiPAdES{_ec: certificate, _db: privateKey, _aab: caCert, _eb: certificateTimestampServerURL}, nil
}

// InitSignature initialises the PdfSignature.
func (_fffb *adobePKCS7Detached) InitSignature(sig *_fe.PdfSignature) error {
	if !_fffb._bebe {
		if _fffb._dce == nil {
			return _b.New("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c")
		}
		if _fffb._eca == nil {
			return _b.New("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c")
		}
	}
	_acg := *_fffb
	sig.Handler = &_acg
	sig.Filter = _da.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _da.MakeName("\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064")
	sig.Reference = nil
	_bbe, _ggdf := _acg.NewDigest(sig)
	if _ggdf != nil {
		return _ggdf
	}
	_bbe.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
	return _acg.Sign(sig, _bbe)
}

type adobeX509RSASHA1 struct {
	_cccd *_ca.PrivateKey
	_dga  *_df.Certificate
	_cag  SignFunc
	_cafe bool
	_fca  _ff.Hash
}

func (_ecc *etsiPAdES) addDss(_fdc, _geab []*_df.Certificate, _bga *RevocationInfoArchival) (int, error) {
	_bce, _gfd, _bcc := _ecc.buildCertChain(_fdc, _geab)
	if _bcc != nil {
		return 0, _bcc
	}
	_gfe, _bcc := _ecc.getCerts(_bce)
	if _bcc != nil {
		return 0, _bcc
	}
	var _ceeg, _gag [][]byte
	if _ecc.OCSPClient != nil {
		_ceeg, _bcc = _ecc.getOCSPs(_bce, _gfd)
		if _bcc != nil {
			return 0, _bcc
		}
	}
	if _ecc.CRLClient != nil {
		_gag, _bcc = _ecc.getCRLs(_bce)
		if _bcc != nil {
			return 0, _bcc
		}
	}
	if !_ecc._bab {
		_, _bcc = _ecc._gde.AddCerts(_gfe)
		if _bcc != nil {
			return 0, _bcc
		}
		_, _bcc = _ecc._gde.AddOCSPs(_ceeg)
		if _bcc != nil {
			return 0, _bcc
		}
		_, _bcc = _ecc._gde.AddCRLs(_gag)
		if _bcc != nil {
			return 0, _bcc
		}
	}
	_bgb := 0
	for _, _cde := range _gag {
		_bgb += len(_cde)
		_bga.Crl = append(_bga.Crl, _ad.RawValue{FullBytes: _cde})
	}
	for _, _ddf := range _ceeg {
		_bgb += len(_ddf)
		_bga.Ocsp = append(_bga.Ocsp, _ad.RawValue{FullBytes: _ddf})
	}
	return _bgb, nil
}

func (_ecfc *adobeX509RSASHA1) getCertificate(_dfc *_fe.PdfSignature) (*_df.Certificate, error) {
	if _ecfc._dga != nil {
		return _ecfc._dga, nil
	}
	_gfef, _gefc := _dfc.GetCerts()
	if _gefc != nil {
		return nil, _gefc
	}
	return _gfef[0], nil
}

// NewDigest creates a new digest.
func (_beg *etsiPAdES) NewDigest(_ *_fe.PdfSignature) (_fe.Hasher, error) {
	return _dc.NewBuffer(nil), nil
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_gaf *adobeX509RSASHA1) IsApplicable(sig *_fe.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031"
}

// Sign sets the Contents fields for the PdfSignature.
func (_ded *etsiPAdES) Sign(sig *_fe.PdfSignature, digest _fe.Hasher) error {
	_ceb, _ccfd := digest.(*_dc.Buffer)
	if !_ccfd {
		return _ce.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_caaf, _dea := _ga.NewSignedData(_ceb.Bytes())
	if _dea != nil {
		return _dea
	}
	_caaf.SetDigestAlgorithm(_ga.OIDDigestAlgorithmSHA256)
	_ddb := _ga.SignerInfoConfig{}
	_dbc := _ff.SHA256.New()
	_dbc.Write(_ded._ec.Raw)
	var _bee struct {
		Seq struct{ Seq struct{ Value []byte } }
	}
	_bee.Seq.Seq.Value = _dbc.Sum(nil)
	var _aag []*_df.Certificate
	var _fag []*_df.Certificate
	if _ded._aab != nil {
		_fag = []*_df.Certificate{_ded._aab}
	}
	_beb := RevocationInfoArchival{Crl: []_ad.RawValue{}, Ocsp: []_ad.RawValue{}, OtherRevInfo: []_ad.RawValue{}}
	_ceec := 0
	if _ded._be != nil && len(_ded._eb) > 0 {
		_af, _fgc := _ded.makeTimestampRequest(_ded._eb, ([]byte)(""))
		if _fgc != nil {
			return _fgc
		}
		_bg, _fgc := _dfa.Parse(_af.FullBytes)
		if _fgc != nil {
			return _fgc
		}
		_aag = append(_aag, _bg.Certificates...)
	}
	if _ded._be != nil {
		_cbg, _ddbc := _ded.addDss([]*_df.Certificate{_ded._ec}, _fag, &_beb)
		if _ddbc != nil {
			return _ddbc
		}
		_ceec += _cbg
		if len(_aag) > 0 {
			_cbg, _ddbc = _ded.addDss(_aag, nil, &_beb)
			if _ddbc != nil {
				return _ddbc
			}
			_ceec += _cbg
		}
		if !_ded._bab {
			_ded._be.SetDSS(_ded._gde)
		}
	}
	_ddb.ExtraSignedAttributes = append(_ddb.ExtraSignedAttributes, _ga.Attribute{Type: _ga.OIDAttributeSigningCertificateV2, Value: _bee}, _ga.Attribute{Type: _ga.OIDAttributeAdobeRevocation, Value: _beb})
	if _bbbd := _caaf.AddSignerChainPAdES(_ded._ec, _ded._db, _fag, _ddb); _bbbd != nil {
		return _bbbd
	}
	_caaf.Detach()
	if len(_ded._eb) > 0 {
		_adb := _caaf.GetSignedData().SignerInfos[0].EncryptedDigest
		_ege, _bbbc := _ded.makeTimestampRequest(_ded._eb, _adb)
		if _bbbc != nil {
			return _bbbc
		}
		_bbbc = _caaf.AddTimestampTokenToSigner(0, _ege.FullBytes)
		if _bbbc != nil {
			return _bbbc
		}
	}
	_ebc, _dea := _caaf.Finish()
	if _dea != nil {
		return _dea
	}
	_bfg := make([]byte, len(_ebc)+1024*2+_ceec)
	copy(_bfg, _ebc)
	sig.Contents = _da.MakeHexString(string(_bfg))
	if !_ded._bab && _ded._gde != nil {
		_dbc = _ff.SHA1.New()
		_dbc.Write(_bfg)
		_ebb := _a.ToUpper(_f.EncodeToString(_dbc.Sum(nil)))
		if _ebb != "" {
			_ded._gde.VRI[_ebb] = &_fe.VRI{Cert: _ded._gde.Certs, OCSP: _ded._gde.OCSPs, CRL: _ded._gde.CRLs}
		}
		_ded._be.SetDSS(_ded._gde)
	}
	return nil
}

type docTimeStamp struct {
	_cccf string
	_bea  _ff.Hash
	_dade int
	_fced *_e.TimestampClient
}

func (_ddfe *adobeX509RSASHA1) getHashAlgorithm(_feba *_fe.PdfSignature) (_ff.Hash, error) {
	_dbb, _gec := _ddfe.getCertificate(_feba)
	if _gec != nil {
		if _ddfe._fca != 0 {
			return _ddfe._fca, nil
		}
		return _acfb, _gec
	}
	if _feba.Contents != nil {
		_eef := _feba.Contents.Bytes()
		var _afc []byte
		if _, _gfg := _ad.Unmarshal(_eef, &_afc); _gfg == nil {
			_faf := _cbgc(_dbb.PublicKey.(*_ca.PublicKey), _afc)
			if _faf > 0 {
				return _faf, nil
			}
		}
	}
	if _ddfe._fca != 0 {
		return _ddfe._fca, nil
	}
	return _acfb, nil
}

func (_dgbg *etsiPAdES) buildCertChain(_fg, _gdb []*_df.Certificate) ([]*_df.Certificate, map[string]*_df.Certificate, error) {
	_eda := map[string]*_df.Certificate{}
	for _, _gbe := range _fg {
		_eda[_gbe.Subject.CommonName] = _gbe
	}
	_adda := _fg
	for _, _aac := range _gdb {
		_cdc := _aac.Subject.CommonName
		if _, _ffe := _eda[_cdc]; _ffe {
			continue
		}
		_eda[_cdc] = _aac
		_adda = append(_adda, _aac)
	}
	if len(_adda) == 0 {
		return nil, nil, _fe.ErrSignNoCertificates
	}
	var _dgbgb error
	for _dge := _adda[0]; _dge != nil && !_dgbg.CertClient.IsCA(_dge); {
		var _abg *_df.Certificate
		_, _caa := _eda[_dge.Issuer.CommonName]
		if !_caa {
			if _abg, _dgbgb = _dgbg.CertClient.GetIssuer(_dge); _dgbgb != nil {
				_dd.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u006f\u0075\u006cd\u0020\u006e\u006f\u0074\u0020\u0072\u0065tr\u0069\u0065\u0076\u0065 \u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061te\u0020\u0069s\u0073\u0075\u0065\u0072\u003a\u0020\u0025\u0076", _dgbgb)
				break
			}
			_eda[_dge.Issuer.CommonName] = _abg
			_adda = append(_adda, _abg)
		} else {
			break
		}
		_dge = _abg
	}
	return _adda, _eda, nil
}

// NewAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobePKCS7Detached(privateKey *_ca.PrivateKey, certificate *_df.Certificate) (_fe.SignatureHandler, error) {
	return &adobePKCS7Detached{_dce: certificate, _eca: privateKey}, nil
}

// Validate validates PdfSignature.
func (_ffb *etsiPAdES) Validate(sig *_fe.PdfSignature, digest _fe.Hasher) (_fe.SignatureValidationResult, error) {
	_ggd := sig.Contents.Bytes()
	_fc, _ea := _ga.Parse(_ggd)
	if _ea != nil {
		return _fe.SignatureValidationResult{}, _ea
	}
	_ceda, _egf := digest.(*_dc.Buffer)
	if !_egf {
		return _fe.SignatureValidationResult{}, _ce.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_fc.Content = _ceda.Bytes()
	if _ea = _fc.Verify(); _ea != nil {
		return _fe.SignatureValidationResult{}, _ea
	}
	_efc := false
	_fce := false
	var _fda _cd.Time
	for _, _gga := range _fc.Signers {
		_gebg := _gga.EncryptedDigest
		var _gef RevocationInfoArchival
		_ea = _fc.UnmarshalSignedAttribute(_ga.OIDAttributeAdobeRevocation, &_gef)
		if _ea == nil {
			if len(_gef.Crl) > 0 {
				_fce = true
			}
			if len(_gef.Ocsp) > 0 {
				_efc = true
			}
		}
		for _, _dgc := range _gga.UnauthenticatedAttributes {
			if _dgc.Type.Equal(_ga.OIDAttributeTimeStampToken) {
				_bdd, _afa := _dfa.Parse(_dgc.Value.Bytes)
				if _afa != nil {
					return _fe.SignatureValidationResult{}, _afa
				}
				_fda = _bdd.Time
				_dde := _bdd.HashAlgorithm.New()
				_dde.Write(_gebg)
				if !_dc.Equal(_dde.Sum(nil), _bdd.HashedMessage) {
					return _fe.SignatureValidationResult{}, _ce.Errorf("\u0048\u0061\u0073\u0068\u0020i\u006e\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061\u006d\u0070\u0020\u0069s\u0020\u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020\u0066\u0072\u006f\u006d\u0020\u0070\u006b\u0063\u0073\u0037")
				}
				break
			}
		}
	}
	_adbb := _fe.SignatureValidationResult{IsSigned: true, IsVerified: true, IsCrlFound: _fce, IsOcspFound: _efc, GeneralizedTime: _fda}
	return _adbb, nil
}

// NewEtsiPAdESLevelB creates a new Adobe.PPKLite ETSI.CAdES.detached Level B signature handler.
func NewEtsiPAdESLevelB(privateKey *_ca.PrivateKey, certificate *_df.Certificate, caCert *_df.Certificate) (_fe.SignatureHandler, error) {
	return &etsiPAdES{_ec: certificate, _db: privateKey, _aab: caCert}, nil
}

type adobePKCS7Detached struct {
	_eca  *_ca.PrivateKey
	_dce  *_df.Certificate
	_bebe bool
	_baba int
}

// NewEmptyAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached
// signature handler. The generated signature is empty and of size signatureLen.
// The signatureLen parameter can be 0 for the signature validation.
func NewEmptyAdobePKCS7Detached(signatureLen int) (_fe.SignatureHandler, error) {
	return &adobePKCS7Detached{_bebe: true, _baba: signatureLen}, nil
}

// ValidateWithOpts validates a PDF signature by checking PdfReader or PdfParser by the DiffPolicy
// params describes parameters for the DocMDP checks.
func (_fec *DocMDPHandler) ValidateWithOpts(sig *_fe.PdfSignature, digest _fe.Hasher, params _fe.SignatureHandlerDocMDPParams) (_fe.SignatureValidationResult, error) {
	_bf, _ccc := _fec._ed.Validate(sig, digest)
	if _ccc != nil {
		return _bf, _ccc
	}
	_aa := params.Parser
	if _aa == nil {
		return _fe.SignatureValidationResult{}, _b.New("p\u0061r\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027t\u0020\u0062\u0065\u0020nu\u006c\u006c")
	}
	if !_bf.IsVerified {
		return _bf, nil
	}
	_edf := params.DiffPolicy
	if _edf == nil {
		_edf = _dg.NewDefaultDiffPolicy()
	}
	for _cef := 0; _cef <= _aa.GetRevisionNumber(); _cef++ {
		_gea, _gfc := _aa.GetRevision(_cef)
		if _gfc != nil {
			return _fe.SignatureValidationResult{}, _gfc
		}
		_ced := _gea.GetTrailer()
		if _ced == nil {
			return _fe.SignatureValidationResult{}, _b.New("\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064\u0020\u0074\u0068\u0065\u0020\u0074r\u0061i\u006c\u0065\u0072\u0020\u006f\u0062\u006ae\u0063\u0074")
		}
		_gd, _eg := _da.GetDict(_ced.Get("\u0052\u006f\u006f\u0074"))
		if !_eg {
			return _fe.SignatureValidationResult{}, _b.New("\u0075n\u0064\u0065\u0066\u0069n\u0065\u0064\u0020\u0074\u0068e\u0020r\u006fo\u0074\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
		_cdf, _eg := _da.GetDict(_gd.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
		if !_eg {
			continue
		}
		_ef, _eg := _da.GetArray(_cdf.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if !_eg {
			continue
		}
		for _, _fa := range _ef.Elements() {
			_ccb, _fb := _da.GetDict(_fa)
			if !_fb {
				continue
			}
			_cfe, _fb := _da.GetDict(_ccb.Get("\u0056"))
			if !_fb {
				continue
			}
			if _da.EqualObjects(_cfe.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"), sig.Contents) {
				_bf.DiffResults, _gfc = _edf.ReviewFile(_gea, _aa, &_dg.MDPParameters{DocMDPLevel: _fec.Permission})
				if _gfc != nil {
					return _fe.SignatureValidationResult{}, _gfc
				}
				_bf.IsVerified = _bf.DiffResults.IsPermitted()
				return _bf, nil
			}
		}
	}
	return _fe.SignatureValidationResult{}, _b.New("\u0064\u006f\u006e\u0027\u0074\u0020\u0066o\u0075\u006e\u0064 \u0074\u0068\u0069\u0073 \u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073")
}

// NewDigest creates a new digest.
func (_dbf *adobePKCS7Detached) NewDigest(sig *_fe.PdfSignature) (_fe.Hasher, error) {
	return _dc.NewBuffer(nil), nil
}

// NewAdobeX509RSASHA1Custom creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. Both the
// certificate and the sign function can be nil for the signature validation.
// NOTE: the handler will do a mock Sign when initializing the signature in
// order to estimate the signature size. Use NewAdobeX509RSASHA1CustomWithOpts
// for configuring the handler to estimate the signature size.
func NewAdobeX509RSASHA1Custom(certificate *_df.Certificate, signFunc SignFunc) (_fe.SignatureHandler, error) {
	return &adobeX509RSASHA1{_dga: certificate, _cag: signFunc}, nil
}

func (_gbf *adobePKCS7Detached) getCertificate(_feb *_fe.PdfSignature) (*_df.Certificate, error) {
	if _gbf._dce != nil {
		return _gbf._dce, nil
	}
	_cgc, _deg := _feb.GetCerts()
	if _deg != nil {
		return nil, _deg
	}
	return _cgc[0], nil
}

func (_cfg *etsiPAdES) makeTimestampRequest(_cfa string, _dda []byte) (_ad.RawValue, error) {
	_cce := _ff.SHA512.New()
	_cce.Write(_dda)
	_gg := _cce.Sum(nil)
	_ecf := _dfa.Request{HashAlgorithm: _ff.SHA512, HashedMessage: _gg, Certificates: true, Extensions: nil, ExtraExtensions: nil}
	_eec := _e.NewTimestampClient()
	_bd, _cg := _eec.GetEncodedToken(_cfa, &_ecf)
	if _cg != nil {
		return _ad.NullRawValue, _cg
	}
	return _ad.RawValue{FullBytes: _bd}, nil
}

type etsiPAdES struct {
	_db  *_ca.PrivateKey
	_ec  *_df.Certificate
	_acd bool
	_bab bool
	_aab *_df.Certificate
	_eb  string

	// CertClient is the client used to retrieve certificates.
	CertClient *_e.CertClient

	// OCSPClient is the client used to retrieve OCSP validation information.
	OCSPClient *_e.OCSPClient

	// CRLClient is the client used to retrieve CRL validation information.
	CRLClient *_e.CRLClient
	_be       *_fe.PdfAppender
	_gde      *_fe.DSS
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
	Algorithm _ff.Hash
}

// InitSignature initialization of the DocMDP signature.
func (_add *DocMDPHandler) InitSignature(sig *_fe.PdfSignature) error {
	_dcd := _add._ed.InitSignature(sig)
	if _dcd != nil {
		return _dcd
	}
	sig.Handler = _add
	if sig.Reference == nil {
		sig.Reference = _da.MakeArray()
	}
	sig.Reference.Append(_fe.NewPdfSignatureReferenceDocMDP(_fe.NewPdfTransformParamsDocMDP(_add.Permission)).ToPdfObject())
	return nil
}

// Validate validates PdfSignature.
func (_dgae *docTimeStamp) Validate(sig *_fe.PdfSignature, digest _fe.Hasher) (_fe.SignatureValidationResult, error) {
	_cbe := sig.Contents.Bytes()
	_ccbe, _cggb := _ga.Parse(_cbe)
	if _cggb != nil {
		return _fe.SignatureValidationResult{}, _cggb
	}
	if _cggb = _ccbe.Verify(); _cggb != nil {
		return _fe.SignatureValidationResult{}, _cggb
	}
	var _ace timestampInfo
	_, _cggb = _ad.Unmarshal(_ccbe.Content, &_ace)
	if _cggb != nil {
		return _fe.SignatureValidationResult{}, _cggb
	}
	_aafb, _cggb := _cfd(_ace.MessageImprint.HashAlgorithm.Algorithm)
	if _cggb != nil {
		return _fe.SignatureValidationResult{}, _cggb
	}
	_fgeb := _aafb.New()
	_ggb, _eeg := digest.(*_dc.Buffer)
	if !_eeg {
		return _fe.SignatureValidationResult{}, _ce.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_fgeb.Write(_ggb.Bytes())
	_cbce := _fgeb.Sum(nil)
	_gafd := _fe.SignatureValidationResult{IsSigned: true, IsVerified: _dc.Equal(_cbce, _ace.MessageImprint.HashedMessage), GeneralizedTime: _ace.GeneralizedTime}
	return _gafd, nil
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature
func (_ebe *adobePKCS7Detached) IsApplicable(sig *_fe.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064"
}

func _cfd(_ffbd _ad.ObjectIdentifier) (_ff.Hash, error) {
	switch {
	case _ffbd.Equal(_ga.OIDDigestAlgorithmSHA1), _ffbd.Equal(_ga.OIDDigestAlgorithmECDSASHA1), _ffbd.Equal(_ga.OIDDigestAlgorithmDSA), _ffbd.Equal(_ga.OIDDigestAlgorithmDSASHA1), _ffbd.Equal(_ga.OIDEncryptionAlgorithmRSA):
		return _ff.SHA1, nil
	case _ffbd.Equal(_ga.OIDDigestAlgorithmSHA256), _ffbd.Equal(_ga.OIDDigestAlgorithmECDSASHA256):
		return _ff.SHA256, nil
	case _ffbd.Equal(_ga.OIDDigestAlgorithmSHA384), _ffbd.Equal(_ga.OIDDigestAlgorithmECDSASHA384):
		return _ff.SHA384, nil
	case _ffbd.Equal(_ga.OIDDigestAlgorithmSHA512), _ffbd.Equal(_ga.OIDDigestAlgorithmECDSASHA512):
		return _ff.SHA512, nil
	}
	return _ff.Hash(0), _ga.ErrUnsupportedAlgorithm
}

func (_bc *etsiPAdES) getCerts(_fbf []*_df.Certificate) ([][]byte, error) {
	_baee := make([][]byte, 0, len(_fbf))
	for _, _aea := range _fbf {
		_baee = append(_baee, _aea.Raw)
	}
	return _baee, nil
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_caf *DocMDPHandler) IsApplicable(sig *_fe.PdfSignature) bool {
	_ge := false
	for _, _gb := range sig.Reference.Elements() {
		if _gf, _dcb := _da.GetDict(_gb); _dcb {
			if _fd, _ae := _da.GetNameVal(_gf.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064")); _ae {
				if _fd != "\u0044\u006f\u0063\u004d\u0044\u0050" {
					return false
				}
				if _ag, _baf := _da.GetDict(_gf.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073")); _baf {
					_, _dgg := _da.GetNumberAsInt64(_ag.Get("\u0050"))
					if _dgg != nil {
						return false
					}
					_ge = true
					break
				}
			}
		}
	}
	return _ge && _caf._ed.IsApplicable(sig)
}

// NewDocTimeStampWithOpts returns a new DocTimeStamp configured using the
// specified options. If no options are provided, default options will be used.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
func NewDocTimeStampWithOpts(timestampServerURL string, hashAlgorithm _ff.Hash, opts *DocTimeStampOpts) (_fe.SignatureHandler, error) {
	if opts == nil {
		opts = &DocTimeStampOpts{}
	}
	if opts.SignatureSize <= 0 {
		opts.SignatureSize = 4192
	}
	return &docTimeStamp{_cccf: timestampServerURL, _bea: hashAlgorithm, _dade: opts.SignatureSize, _fced: opts.Client}, nil
}

func (_fecg *etsiPAdES) getCRLs(_defb []*_df.Certificate) ([][]byte, error) {
	_deff := make([][]byte, 0, len(_defb))
	for _, _ab := range _defb {
		for _, _efe := range _ab.CRLDistributionPoints {
			if _fecg.CertClient.IsCA(_ab) {
				continue
			}
			_ceef, _fee := _fecg.CRLClient.MakeRequest(_efe, _ab)
			if _fee != nil {
				_dd.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043R\u004c\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074 \u0065\u0072\u0072o\u0072:\u0020\u0025\u0076", _fee)
				continue
			}
			_deff = append(_deff, _ceef)
		}
	}
	return _deff, nil
}

// NewAdobeX509RSASHA1CustomWithOpts creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. The
// handler is configured based on the provided options. If no options are
// provided, default options will be used. Both the certificate and the sign
// function can be nil for the signature validation.
func NewAdobeX509RSASHA1CustomWithOpts(certificate *_df.Certificate, signFunc SignFunc, opts *AdobeX509RSASHA1Opts) (_fe.SignatureHandler, error) {
	if opts == nil {
		opts = &AdobeX509RSASHA1Opts{}
	}
	return &adobeX509RSASHA1{_dga: certificate, _cag: signFunc, _cafe: opts.EstimateSize, _fca: opts.Algorithm}, nil
}

// DocMDPHandler describes handler for the DocMDP realization.
type DocMDPHandler struct {
	_ed        _fe.SignatureHandler
	Permission _dg.DocMDPPermission
}

// NewDigest creates a new digest.
func (_agc *adobeX509RSASHA1) NewDigest(sig *_fe.PdfSignature) (_fe.Hasher, error) {
	if _feg, _fcea := _agc.getHashAlgorithm(sig); _feg != 0 && _fcea == nil {
		return _feg.New(), nil
	}
	return _acfb.New(), nil
}

func _eacc(_ebcd []byte, _gdc int) (_cab []byte) {
	_cge := len(_ebcd)
	if _cge > _gdc {
		_cge = _gdc
	}
	_cab = make([]byte, _gdc)
	copy(_cab[len(_cab)-_cge:], _ebcd)
	return
}

// NewDigest creates a new digest.
func (_ccf *DocMDPHandler) NewDigest(sig *_fe.PdfSignature) (_fe.Hasher, error) {
	return _ccf._ed.NewDigest(sig)
}

// Sign sets the Contents fields for the PdfSignature.
func (_acfd *adobeX509RSASHA1) Sign(sig *_fe.PdfSignature, digest _fe.Hasher) error {
	var _dadb []byte
	var _bcdc error
	if _acfd._cag != nil {
		_dadb, _bcdc = _acfd._cag(sig, digest)
		if _bcdc != nil {
			return _bcdc
		}
	} else {
		_aabe, _fcg := digest.(_cc.Hash)
		if !_fcg {
			return _b.New("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_dfb := _acfb
		if _acfd._fca != 0 {
			_dfb = _acfd._fca
		}
		_dadb, _bcdc = _ca.SignPKCS1v15(_ba.Reader, _acfd._cccd, _dfb, _aabe.Sum(nil))
		if _bcdc != nil {
			return _bcdc
		}
	}
	_dadb, _bcdc = _ad.Marshal(_dadb)
	if _bcdc != nil {
		return _bcdc
	}
	sig.Contents = _da.MakeHexString(string(_dadb))
	return nil
}

func (_gc *etsiPAdES) getOCSPs(_dgb []*_df.Certificate, _ffa map[string]*_df.Certificate) ([][]byte, error) {
	_bafa := make([][]byte, 0, len(_dgb))
	for _, _cefd := range _dgb {
		for _, _gcb := range _cefd.OCSPServer {
			if _gc.CertClient.IsCA(_cefd) {
				continue
			}
			_bb, _gda := _ffa[_cefd.Issuer.CommonName]
			if !_gda {
				_dd.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074\u003a\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
				continue
			}
			_, _def, _gcg := _gc.OCSPClient.MakeRequest(_gcb, _cefd, _bb)
			if _gcg != nil {
				_dd.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075e\u0073t\u0020\u0065\u0072\u0072\u006f\u0072\u003a \u0025\u0076", _gcg)
				continue
			}
			_bafa = append(_bafa, _def)
		}
	}
	return _bafa, nil
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
	Client *_e.TimestampClient
}

func (_aebg *adobeX509RSASHA1) sign(_bfgc *_fe.PdfSignature, _gdbe _fe.Hasher, _cgcb bool) error {
	if !_cgcb {
		return _aebg.Sign(_bfgc, _gdbe)
	}
	_gbb, _ega := _aebg._dga.PublicKey.(*_ca.PublicKey)
	if !_ega {
		return _ce.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0075\u0062\u006c\u0069\u0063\u0020\u006b\u0065y\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", _gbb)
	}
	_bdg, _dbfc := _ad.Marshal(make([]byte, _gbb.Size()))
	if _dbfc != nil {
		return _dbfc
	}
	_bfgc.Contents = _da.MakeHexString(string(_bdg))
	return nil
}

type timestampInfo struct {
	Version        int
	Policy         _ad.RawValue
	MessageImprint struct {
		HashAlgorithm _d.AlgorithmIdentifier
		HashedMessage []byte
	}
	SerialNumber    _ad.RawValue
	GeneralizedTime _cd.Time
}

// InitSignature initialises the PdfSignature.
func (_dcdg *adobeX509RSASHA1) InitSignature(sig *_fe.PdfSignature) error {
	if _dcdg._dga == nil {
		return _b.New("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c")
	}
	if _dcdg._cccd == nil && _dcdg._cag == nil {
		return _b.New("\u006d\u0075\u0073\u0074\u0020\u0070\u0072o\u0076\u0069\u0064e\u0020\u0065\u0069t\u0068\u0065r\u0020\u0061\u0020\u0070\u0072\u0069v\u0061te\u0020\u006b\u0065\u0079\u0020\u006f\u0072\u0020\u0061\u0020\u0073\u0069\u0067\u006e\u0069\u006e\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	}
	_ddec := *_dcdg
	sig.Handler = &_ddec
	sig.Filter = _da.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _da.MakeName("\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031")
	sig.Cert = _da.MakeString(string(_ddec._dga.Raw))
	sig.Reference = nil
	_fgb, _acga := _ddec.NewDigest(sig)
	if _acga != nil {
		return _acga
	}
	_fgb.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
	return _ddec.sign(sig, _fgb, _dcdg._cafe)
}
