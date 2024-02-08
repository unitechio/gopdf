package sighandler

import (
	_ge "bytes"
	_ffb "crypto"
	_eg "crypto/rand"
	_gg "crypto/rsa"
	_bc "crypto/x509"
	_dfc "crypto/x509/pkix"
	_b "encoding/asn1"
	_eb "encoding/hex"
	_a "errors"
	_df "fmt"
	_d "hash"
	_f "math/big"
	_g "strings"
	_ff "time"

	_ba "bitbucket.org/shenghui0779/gopdf/common"
	_be "bitbucket.org/shenghui0779/gopdf/core"
	_ggg "bitbucket.org/shenghui0779/gopdf/model"
	_ec "bitbucket.org/shenghui0779/gopdf/model/mdp"
	_de "bitbucket.org/shenghui0779/gopdf/model/sigutil"
	_bg "github.com/unidoc/pkcs7"
	_da "github.com/unidoc/timestamp"
)

func (_cbb *etsiPAdES) makeTimestampRequest(_egf string, _egfg []byte) (_b.RawValue, error) {
	_edd := _ffb.SHA512.New()
	_edd.Write(_egfg)
	_ab := _edd.Sum(nil)
	_fg := _da.Request{HashAlgorithm: _ffb.SHA512, HashedMessage: _ab, Certificates: true, Extensions: nil, ExtraExtensions: nil}
	_gbg := _de.NewTimestampClient()
	_agf, _cfec := _gbg.GetEncodedToken(_egf, &_fg)
	if _cfec != nil {
		return _b.NullRawValue, _cfec
	}
	return _b.RawValue{FullBytes: _agf}, nil
}

// Sign adds a new reference to signature's references array.
func (_af *DocMDPHandler) Sign(sig *_ggg.PdfSignature, digest _ggg.Hasher) error {
	return _af._ef.Sign(sig, digest)
}

// NewDigest creates a new digest.
func (_bf *DocMDPHandler) NewDigest(sig *_ggg.PdfSignature) (_ggg.Hasher, error) {
	return _bf._ef.NewDigest(sig)
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_c *DocMDPHandler) IsApplicable(sig *_ggg.PdfSignature) bool {
	_cd := false
	for _, _ecb := range sig.Reference.Elements() {
		if _fc, _ad := _be.GetDict(_ecb); _ad {
			if _ebf, _ffc := _be.GetNameVal(_fc.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064")); _ffc {
				if _ebf != "\u0044\u006f\u0063\u004d\u0044\u0050" {
					return false
				}
				if _ea, _aa := _be.GetDict(_fc.Get("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073")); _aa {
					_, _dg := _be.GetNumberAsInt64(_ea.Get("\u0050"))
					if _dg != nil {
						return false
					}
					_cd = true
					break
				}
			}
		}
	}
	return _cd && _c._ef.IsApplicable(sig)
}

// NewAdobeX509RSASHA1 creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler. Both the private key and the
// certificate can be nil for the signature validation.
func NewAdobeX509RSASHA1(privateKey *_gg.PrivateKey, certificate *_bc.Certificate) (_ggg.SignatureHandler, error) {
	return &adobeX509RSASHA1{_dfg: certificate, _abc: privateKey}, nil
}

type etsiPAdES struct {
	_ca   *_gg.PrivateKey
	_gge  *_bc.Certificate
	_bcd  bool
	_eaeg bool
	_fd   *_bc.Certificate
	_ffd  string

	// CertClient is the client used to retrieve certificates.
	CertClient *_de.CertClient

	// OCSPClient is the client used to retrieve OCSP validation information.
	OCSPClient *_de.OCSPClient

	// CRLClient is the client used to retrieve CRL validation information.
	CRLClient *_de.CRLClient
	_fca      *_ggg.PdfAppender
	_bcdf     *_ggg.DSS
}

// Sign sets the Contents fields.
func (_bfb *adobePKCS7Detached) Sign(sig *_ggg.PdfSignature, digest _ggg.Hasher) error {
	if _bfb._ddbab {
		_fefc := _bfb._gddd
		if _fefc <= 0 {
			_fefc = 8192
		}
		sig.Contents = _be.MakeHexString(string(make([]byte, _fefc)))
		return nil
	}
	_ecc, _acdg := digest.(*_ge.Buffer)
	if !_acdg {
		return _df.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_efc, _caba := _bg.NewSignedData(_ecc.Bytes())
	if _caba != nil {
		return _caba
	}
	if _ebcf := _efc.AddSigner(_bfb._efe, _bfb._acd, _bg.SignerInfoConfig{}); _ebcf != nil {
		return _ebcf
	}
	_efc.Detach()
	_fcce, _caba := _efc.Finish()
	if _caba != nil {
		return _caba
	}
	_bgg := make([]byte, 8192)
	copy(_bgg, _fcce)
	sig.Contents = _be.MakeHexString(string(_bgg))
	return nil
}

// Validate validates PdfSignature.
func (_fgg *adobeX509RSASHA1) Validate(sig *_ggg.PdfSignature, digest _ggg.Hasher) (_ggg.SignatureValidationResult, error) {
	_ecbc, _cff := _fgg.getCertificate(sig)
	if _cff != nil {
		return _ggg.SignatureValidationResult{}, _cff
	}
	_dfccc := sig.Contents.Bytes()
	var _cca []byte
	if _, _ddeb := _b.Unmarshal(_dfccc, &_cca); _ddeb != nil {
		return _ggg.SignatureValidationResult{}, _ddeb
	}
	_dfea, _daeg := digest.(_d.Hash)
	if !_daeg {
		return _ggg.SignatureValidationResult{}, _a.New("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
	}
	_dbb, _ := _fgg.getHashAlgorithm(sig)
	if _dbb == 0 {
		_dbb = _afe
	}
	if _aeb := _gg.VerifyPKCS1v15(_ecbc.PublicKey.(*_gg.PublicKey), _dbb, _dfea.Sum(nil), _cca); _aeb != nil {
		return _ggg.SignatureValidationResult{}, _aeb
	}
	return _ggg.SignatureValidationResult{IsSigned: true, IsVerified: true}, nil
}
func (_eddd *adobeX509RSASHA1) getHashAlgorithm(_afb *_ggg.PdfSignature) (_ffb.Hash, error) {
	_gfg, _ede := _eddd.getCertificate(_afb)
	if _ede != nil {
		if _eddd._bcc != 0 {
			return _eddd._bcc, nil
		}
		return _afe, _ede
	}
	if _afb.Contents != nil {
		_cfg := _afb.Contents.Bytes()
		var _egdb []byte
		if _, _faeb := _b.Unmarshal(_cfg, &_egdb); _faeb == nil {
			_gfd := _gbd(_gfg.PublicKey.(*_gg.PublicKey), _egdb)
			if _gfd > 0 {
				return _gfd, nil
			}
		}
	}
	if _eddd._bcc != 0 {
		return _eddd._bcc, nil
	}
	return _afe, nil
}
func _cea(_bafc _b.ObjectIdentifier) (_ffb.Hash, error) {
	switch {
	case _bafc.Equal(_bg.OIDDigestAlgorithmSHA1), _bafc.Equal(_bg.OIDDigestAlgorithmECDSASHA1), _bafc.Equal(_bg.OIDDigestAlgorithmDSA), _bafc.Equal(_bg.OIDDigestAlgorithmDSASHA1), _bafc.Equal(_bg.OIDEncryptionAlgorithmRSA):
		return _ffb.SHA1, nil
	case _bafc.Equal(_bg.OIDDigestAlgorithmSHA256), _bafc.Equal(_bg.OIDDigestAlgorithmECDSASHA256):
		return _ffb.SHA256, nil
	case _bafc.Equal(_bg.OIDDigestAlgorithmSHA384), _bafc.Equal(_bg.OIDDigestAlgorithmECDSASHA384):
		return _ffb.SHA384, nil
	case _bafc.Equal(_bg.OIDDigestAlgorithmSHA512), _bafc.Equal(_bg.OIDDigestAlgorithmECDSASHA512):
		return _ffb.SHA512, nil
	}
	return _ffb.Hash(0), _bg.ErrUnsupportedAlgorithm
}

// NewDigest creates a new digest.
func (_cge *etsiPAdES) NewDigest(_ *_ggg.PdfSignature) (_ggg.Hasher, error) {
	return _ge.NewBuffer(nil), nil
}

type docTimeStamp struct {
	_acg  string
	_eded _ffb.Hash
	_cfd  int
	_cfga *_de.TimestampClient
}

// NewAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobePKCS7Detached(privateKey *_gg.PrivateKey, certificate *_bc.Certificate) (_ggg.SignatureHandler, error) {
	return &adobePKCS7Detached{_efe: certificate, _acd: privateKey}, nil
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_ecg *etsiPAdES) IsApplicable(sig *_ggg.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064"
}

// InitSignature initialises the PdfSignature.
func (_faec *adobeX509RSASHA1) InitSignature(sig *_ggg.PdfSignature) error {
	if _faec._dfg == nil {
		return _a.New("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c")
	}
	if _faec._abc == nil && _faec._egb == nil {
		return _a.New("\u006d\u0075\u0073\u0074\u0020\u0070\u0072o\u0076\u0069\u0064e\u0020\u0065\u0069t\u0068\u0065r\u0020\u0061\u0020\u0070\u0072\u0069v\u0061te\u0020\u006b\u0065\u0079\u0020\u006f\u0072\u0020\u0061\u0020\u0073\u0069\u0067\u006e\u0069\u006e\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	}
	_dee := *_faec
	sig.Handler = &_dee
	sig.Filter = _be.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _be.MakeName("\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031")
	sig.Cert = _be.MakeString(string(_dee._dfg.Raw))
	sig.Reference = nil
	_fefe, _bcdfb := _dee.NewDigest(sig)
	if _bcdfb != nil {
		return _bcdfb
	}
	_fefe.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
	return _dee.sign(sig, _fefe, _faec._fcd)
}
func (_fbf *adobeX509RSASHA1) getCertificate(_ddc *_ggg.PdfSignature) (*_bc.Certificate, error) {
	if _fbf._dfg != nil {
		return _fbf._dfg, nil
	}
	_fdd, _bea := _ddc.GetCerts()
	if _bea != nil {
		return nil, _bea
	}
	return _fdd[0], nil
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_edfe *docTimeStamp) IsApplicable(sig *_ggg.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031"
}

// NewDigest creates a new digest.
func (_ada *adobeX509RSASHA1) NewDigest(sig *_ggg.PdfSignature) (_ggg.Hasher, error) {
	if _bef, _bcdfbb := _ada.getHashAlgorithm(sig); _bef != 0 && _bcdfbb == nil {
		return _bef.New(), nil
	}
	return _afe.New(), nil
}

// NewAdobeX509RSASHA1CustomWithOpts creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. The
// handler is configured based on the provided options. If no options are
// provided, default options will be used. Both the certificate and the sign
// function can be nil for the signature validation.
func NewAdobeX509RSASHA1CustomWithOpts(certificate *_bc.Certificate, signFunc SignFunc, opts *AdobeX509RSASHA1Opts) (_ggg.SignatureHandler, error) {
	if opts == nil {
		opts = &AdobeX509RSASHA1Opts{}
	}
	return &adobeX509RSASHA1{_dfg: certificate, _egb: signFunc, _fcd: opts.EstimateSize, _bcc: opts.Algorithm}, nil
}
func (_acc *etsiPAdES) buildCertChain(_cda, _ffde []*_bc.Certificate) ([]*_bc.Certificate, map[string]*_bc.Certificate, error) {
	_dde := map[string]*_bc.Certificate{}
	for _, _aff := range _cda {
		_dde[_aff.Subject.CommonName] = _aff
	}
	_agd := _cda
	for _, _fbb := range _ffde {
		_affa := _fbb.Subject.CommonName
		if _, _dgc := _dde[_affa]; _dgc {
			continue
		}
		_dde[_affa] = _fbb
		_agd = append(_agd, _fbb)
	}
	if len(_agd) == 0 {
		return nil, nil, _ggg.ErrSignNoCertificates
	}
	var _cbe error
	for _eaeb := _agd[0]; _eaeb != nil && !_acc.CertClient.IsCA(_eaeb); {
		var _dgff *_bc.Certificate
		_, _bcdfa := _dde[_eaeb.Issuer.CommonName]
		if !_bcdfa {
			if _dgff, _cbe = _acc.CertClient.GetIssuer(_eaeb); _cbe != nil {
				_ba.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043\u006f\u0075\u006cd\u0020\u006e\u006f\u0074\u0020\u0072\u0065tr\u0069\u0065\u0076\u0065 \u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061te\u0020\u0069s\u0073\u0075\u0065\u0072\u003a\u0020\u0025\u0076", _cbe)
				break
			}
			_dde[_eaeb.Issuer.CommonName] = _dgff
			_agd = append(_agd, _dgff)
		} else {
			break
		}
		_eaeb = _dgff
	}
	return _agd, _dde, nil
}

// Validate validates PdfSignature.
func (_ebcg *adobePKCS7Detached) Validate(sig *_ggg.PdfSignature, digest _ggg.Hasher) (_ggg.SignatureValidationResult, error) {
	_accf := sig.Contents.Bytes()
	_eff, _badg := _bg.Parse(_accf)
	if _badg != nil {
		return _ggg.SignatureValidationResult{}, _badg
	}
	_cbec, _eba := digest.(*_ge.Buffer)
	if !_eba {
		return _ggg.SignatureValidationResult{}, _df.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_eff.Content = _cbec.Bytes()
	if _badg = _eff.Verify(); _badg != nil {
		return _ggg.SignatureValidationResult{}, _badg
	}
	return _ggg.SignatureValidationResult{IsSigned: true, IsVerified: true}, nil
}

// NewEtsiPAdESLevelT creates a new Adobe.PPKLite ETSI.CAdES.detached Level T signature handler.
func NewEtsiPAdESLevelT(privateKey *_gg.PrivateKey, certificate *_bc.Certificate, caCert *_bc.Certificate, certificateTimestampServerURL string) (_ggg.SignatureHandler, error) {
	return &etsiPAdES{_gge: certificate, _ca: privateKey, _fd: caCert, _ffd: certificateTimestampServerURL}, nil
}
func (_dbe *etsiPAdES) getCerts(_gggg []*_bc.Certificate) ([][]byte, error) {
	_geg := make([][]byte, 0, len(_gggg))
	for _, _cag := range _gggg {
		_geg = append(_geg, _cag.Raw)
	}
	return _geg, nil
}

const _afe = _ffb.SHA1

// NewDocMDPHandler returns the new DocMDP handler with the specific DocMDP restriction level.
func NewDocMDPHandler(handler _ggg.SignatureHandler, permission _ec.DocMDPPermission) (_ggg.SignatureHandler, error) {
	return &DocMDPHandler{_ef: handler, Permission: permission}, nil
}

// NewDigest creates a new digest.
func (_fea *docTimeStamp) NewDigest(sig *_ggg.PdfSignature) (_ggg.Hasher, error) {
	return _ge.NewBuffer(nil), nil
}

type adobeX509RSASHA1 struct {
	_abc *_gg.PrivateKey
	_dfg *_bc.Certificate
	_egb SignFunc
	_fcd bool
	_bcc _ffb.Hash
}

// Sign sets the Contents fields for the PdfSignature.
func (_aeg *etsiPAdES) Sign(sig *_ggg.PdfSignature, digest _ggg.Hasher) error {
	_aga, _ffa := digest.(*_ge.Buffer)
	if !_ffa {
		return _df.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_dbed, _abf := _bg.NewSignedData(_aga.Bytes())
	if _abf != nil {
		return _abf
	}
	_dbed.SetDigestAlgorithm(_bg.OIDDigestAlgorithmSHA256)
	_fef := _bg.SignerInfoConfig{}
	_aab := _ffb.SHA256.New()
	_aab.Write(_aeg._gge.Raw)
	var _efb struct {
		Seq struct{ Seq struct{ Value []byte } }
	}
	_efb.Seq.Seq.Value = _aab.Sum(nil)
	var _deb []*_bc.Certificate
	var _agc []*_bc.Certificate
	if _aeg._fd != nil {
		_agc = []*_bc.Certificate{_aeg._fd}
	}
	_cbd := RevocationInfoArchival{Crl: []_b.RawValue{}, Ocsp: []_b.RawValue{}, OtherRevInfo: []_b.RawValue{}}
	_aaa := 0
	if _aeg._fca != nil && len(_aeg._ffd) > 0 {
		_bec, _fgc := _aeg.makeTimestampRequest(_aeg._ffd, ([]byte)(""))
		if _fgc != nil {
			return _fgc
		}
		_ce, _fgc := _da.Parse(_bec.FullBytes)
		if _fgc != nil {
			return _fgc
		}
		_deb = append(_deb, _ce.Certificates...)
	}
	if _aeg._fca != nil {
		_gc, _bae := _aeg.addDss([]*_bc.Certificate{_aeg._gge}, _agc, &_cbd)
		if _bae != nil {
			return _bae
		}
		_aaa += _gc
		if len(_deb) > 0 {
			_gc, _bae = _aeg.addDss(_deb, nil, &_cbd)
			if _bae != nil {
				return _bae
			}
			_aaa += _gc
		}
		if !_aeg._eaeg {
			_aeg._fca.SetDSS(_aeg._bcdf)
		}
	}
	_fef.ExtraSignedAttributes = append(_fef.ExtraSignedAttributes, _bg.Attribute{Type: _bg.OIDAttributeSigningCertificateV2, Value: _efb}, _bg.Attribute{Type: _bg.OIDAttributeAdobeRevocation, Value: _cbd})
	if _ddb := _dbed.AddSignerChainPAdES(_aeg._gge, _aeg._ca, _agc, _fef); _ddb != nil {
		return _ddb
	}
	_dbed.Detach()
	if len(_aeg._ffd) > 0 {
		_cab := _dbed.GetSignedData().SignerInfos[0].EncryptedDigest
		_dcb, _bde := _aeg.makeTimestampRequest(_aeg._ffd, _cab)
		if _bde != nil {
			return _bde
		}
		_bde = _dbed.AddTimestampTokenToSigner(0, _dcb.FullBytes)
		if _bde != nil {
			return _bde
		}
	}
	_beea, _abf := _dbed.Finish()
	if _abf != nil {
		return _abf
	}
	_ee := make([]byte, len(_beea)+1024*2+_aaa)
	copy(_ee, _beea)
	sig.Contents = _be.MakeHexString(string(_ee))
	if !_aeg._eaeg && _aeg._bcdf != nil {
		_aab = _ffb.SHA1.New()
		_aab.Write(_ee)
		_gd := _g.ToUpper(_eb.EncodeToString(_aab.Sum(nil)))
		if _gd != "" {
			_aeg._bcdf.VRI[_gd] = &_ggg.VRI{Cert: _aeg._bcdf.Certs, OCSP: _aeg._bcdf.OCSPs, CRL: _aeg._bcdf.CRLs}
		}
		_aeg._fca.SetDSS(_aeg._bcdf)
	}
	return nil
}

// InitSignature initialises the PdfSignature.
func (_cg *etsiPAdES) InitSignature(sig *_ggg.PdfSignature) error {
	if !_cg._bcd {
		if _cg._gge == nil {
			return _a.New("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c")
		}
		if _cg._ca == nil {
			return _a.New("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c")
		}
	}
	_edf := *_cg
	sig.Handler = &_edf
	sig.Filter = _be.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _be.MakeName("\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064")
	sig.Reference = nil
	_gga, _ag := _edf.NewDigest(sig)
	if _ag != nil {
		return _ag
	}
	_, _ag = _gga.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
	if _ag != nil {
		return _ag
	}
	_edf._eaeg = true
	_ag = _edf.Sign(sig, _gga)
	_edf._eaeg = false
	return _ag
}

// Sign sets the Contents fields for the PdfSignature.
func (_cdf *adobeX509RSASHA1) Sign(sig *_ggg.PdfSignature, digest _ggg.Hasher) error {
	var _aedb []byte
	var _bebf error
	if _cdf._egb != nil {
		_aedb, _bebf = _cdf._egb(sig, digest)
		if _bebf != nil {
			return _bebf
		}
	} else {
		_gfa, _gaa := digest.(_d.Hash)
		if !_gaa {
			return _a.New("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072")
		}
		_gab := _afe
		if _cdf._bcc != 0 {
			_gab = _cdf._bcc
		}
		_aedb, _bebf = _gg.SignPKCS1v15(_eg.Reader, _cdf._abc, _gab, _gfa.Sum(nil))
		if _bebf != nil {
			return _bebf
		}
	}
	_aedb, _bebf = _b.Marshal(_aedb)
	if _bebf != nil {
		return _bebf
	}
	sig.Contents = _be.MakeHexString(string(_aedb))
	return nil
}

// NewAdobeX509RSASHA1Custom creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. Both the
// certificate and the sign function can be nil for the signature validation.
// NOTE: the handler will do a mock Sign when initializing the signature in
// order to estimate the signature size. Use NewAdobeX509RSASHA1CustomWithOpts
// for configuring the handler to estimate the signature size.
func NewAdobeX509RSASHA1Custom(certificate *_bc.Certificate, signFunc SignFunc) (_ggg.SignatureHandler, error) {
	return &adobeX509RSASHA1{_dfg: certificate, _egb: signFunc}, nil
}

// NewDocTimeStamp creates a new DocTimeStamp signature handler.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
// NOTE: the handler will do a mock Sign when initializing the signature
// in order to estimate the signature size. Use NewDocTimeStampWithOpts
// for providing the signature size.
func NewDocTimeStamp(timestampServerURL string, hashAlgorithm _ffb.Hash) (_ggg.SignatureHandler, error) {
	return &docTimeStamp{_acg: timestampServerURL, _eded: hashAlgorithm}, nil
}
func (_fbe *etsiPAdES) getCRLs(_gba []*_bc.Certificate) ([][]byte, error) {
	_dff := make([][]byte, 0, len(_gba))
	for _, _ga := range _gba {
		for _, _ddfa := range _ga.CRLDistributionPoints {
			if _fbe.CertClient.IsCA(_ga) {
				continue
			}
			_afd, _ae := _fbe.CRLClient.MakeRequest(_ddfa, _ga)
			if _ae != nil {
				_ba.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0043R\u004c\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074 \u0065\u0072\u0072o\u0072:\u0020\u0025\u0076", _ae)
				continue
			}
			_dff = append(_dff, _afd)
		}
	}
	return _dff, nil
}
func (_bbd *etsiPAdES) getOCSPs(_caa []*_bc.Certificate, _bdb map[string]*_bc.Certificate) ([][]byte, error) {
	_abg := make([][]byte, 0, len(_caa))
	for _, _ggae := range _caa {
		for _, _ac := range _ggae.OCSPServer {
			if _bbd.CertClient.IsCA(_ggae) {
				continue
			}
			_bbc, _ebe := _bdb[_ggae.Issuer.CommonName]
			if !_ebe {
				_ba.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075\u0065\u0073\u0074\u003a\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
				continue
			}
			_, _dc, _dbc := _bbd.OCSPClient.MakeRequest(_ac, _ggae, _bbc)
			if _dbc != nil {
				_ba.Log.Debug("\u0057\u0041\u0052\u004e:\u0020\u004f\u0043\u0053\u0050\u0020\u0072\u0065\u0071\u0075e\u0073t\u0020\u0065\u0072\u0072\u006f\u0072\u003a \u0025\u0076", _dbc)
				continue
			}
			_abg = append(_abg, _dc)
		}
	}
	return _abg, nil
}

type adobePKCS7Detached struct {
	_acd   *_gg.PrivateKey
	_efe   *_bc.Certificate
	_ddbab bool
	_gddd  int
}

// InitSignature initialises the PdfSignature.
func (_ecd *docTimeStamp) InitSignature(sig *_ggg.PdfSignature) error {
	_daa := *_ecd
	sig.Type = _be.MakeName("\u0044\u006f\u0063T\u0069\u006d\u0065\u0053\u0074\u0061\u006d\u0070")
	sig.Handler = &_daa
	sig.Filter = _be.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _be.MakeName("\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031")
	sig.Reference = nil
	if _ecd._cfd > 0 {
		sig.Contents = _be.MakeHexString(string(make([]byte, _ecd._cfd)))
	} else {
		_gagf, _deed := _ecd.NewDigest(sig)
		if _deed != nil {
			return _deed
		}
		_gagf.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
		if _deed = _daa.Sign(sig, _gagf); _deed != nil {
			return _deed
		}
		_ecd._cfd = _daa._cfd
	}
	return nil
}
func (_cad *etsiPAdES) addDss(_egdd, _gdd []*_bc.Certificate, _ddba *RevocationInfoArchival) (int, error) {
	_fgb, _fcc, _agfe := _cad.buildCertChain(_egdd, _gdd)
	if _agfe != nil {
		return 0, _agfe
	}
	_gf, _agfe := _cad.getCerts(_fgb)
	if _agfe != nil {
		return 0, _agfe
	}
	var _bcf, _ggf [][]byte
	if _cad.OCSPClient != nil {
		_bcf, _agfe = _cad.getOCSPs(_fgb, _fcc)
		if _agfe != nil {
			return 0, _agfe
		}
	}
	if _cad.CRLClient != nil {
		_ggf, _agfe = _cad.getCRLs(_fgb)
		if _agfe != nil {
			return 0, _agfe
		}
	}
	if !_cad._eaeg {
		_, _agfe = _cad._bcdf.AddCerts(_gf)
		if _agfe != nil {
			return 0, _agfe
		}
		_, _agfe = _cad._bcdf.AddOCSPs(_bcf)
		if _agfe != nil {
			return 0, _agfe
		}
		_, _agfe = _cad._bcdf.AddCRLs(_ggf)
		if _agfe != nil {
			return 0, _agfe
		}
	}
	_bdc := 0
	for _, _cgg := range _ggf {
		_bdc += len(_cgg)
		_ddba.Crl = append(_ddba.Crl, _b.RawValue{FullBytes: _cgg})
	}
	for _, _aed := range _bcf {
		_bdc += len(_aed)
		_ddba.Ocsp = append(_ddba.Ocsp, _b.RawValue{FullBytes: _aed})
	}
	return _bdc, nil
}

// NewDigest creates a new digest.
func (_eada *adobePKCS7Detached) NewDigest(sig *_ggg.PdfSignature) (_ggg.Hasher, error) {
	return _ge.NewBuffer(nil), nil
}

// Sign sets the Contents fields for the PdfSignature.
func (_cga *docTimeStamp) Sign(sig *_ggg.PdfSignature, digest _ggg.Hasher) error {
	_dbcb, _abgg := _de.NewTimestampRequest(digest.(*_ge.Buffer), &_da.RequestOptions{Hash: _cga._eded, Certificates: true})
	if _abgg != nil {
		return _abgg
	}
	_aac := _cga._cfga
	if _aac == nil {
		_aac = _de.NewTimestampClient()
	}
	_gef, _abgg := _aac.GetEncodedToken(_cga._acg, _dbcb)
	if _abgg != nil {
		return _abgg
	}
	_eab := len(_gef)
	if _cga._cfd > 0 && _eab > _cga._cfd {
		return _ggg.ErrSignNotEnoughSpace
	}
	if _eab > 0 {
		_cga._cfd = _eab + 128
	}
	if sig.Contents != nil {
		_dgab := sig.Contents.Bytes()
		copy(_dgab, _gef)
		_gef = _dgab
	}
	sig.Contents = _be.MakeHexString(string(_gef))
	return nil
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
	Client *_de.TimestampClient
}

// InitSignature initialization of the DocMDP signature.
func (_gbf *DocMDPHandler) InitSignature(sig *_ggg.PdfSignature) error {
	_gbc := _gbf._ef.InitSignature(sig)
	if _gbc != nil {
		return _gbc
	}
	sig.Handler = _gbf
	if sig.Reference == nil {
		sig.Reference = _be.MakeArray()
	}
	sig.Reference.Append(_ggg.NewPdfSignatureReferenceDocMDP(_ggg.NewPdfTransformParamsDocMDP(_gbf.Permission)).ToPdfObject())
	return nil
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
	Algorithm _ffb.Hash
}

// IsApplicable returns true if the signature handler is applicable for the PdfSignature
func (_gff *adobePKCS7Detached) IsApplicable(sig *_ggg.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064"
}
func (_efbe *adobeX509RSASHA1) sign(_baeb *_ggg.PdfSignature, _dbbd _ggg.Hasher, _bge bool) error {
	if !_bge {
		return _efbe.Sign(_baeb, _dbbd)
	}
	_feg, _dfccg := _efbe._dfg.PublicKey.(*_gg.PublicKey)
	if !_dfccg {
		return _df.Errorf("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0075\u0062\u006c\u0069\u0063\u0020\u006b\u0065y\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054", _feg)
	}
	_fab, _bcda := _b.Marshal(make([]byte, _feg.Size()))
	if _bcda != nil {
		return _bcda
	}
	_baeb.Contents = _be.MakeHexString(string(_fab))
	return nil
}

// Validate validates PdfSignature.
func (_bbda *docTimeStamp) Validate(sig *_ggg.PdfSignature, digest _ggg.Hasher) (_ggg.SignatureValidationResult, error) {
	_cde := sig.Contents.Bytes()
	_gdf, _bgdd := _bg.Parse(_cde)
	if _bgdd != nil {
		return _ggg.SignatureValidationResult{}, _bgdd
	}
	if _bgdd = _gdf.Verify(); _bgdd != nil {
		return _ggg.SignatureValidationResult{}, _bgdd
	}
	var _cce timestampInfo
	_, _bgdd = _b.Unmarshal(_gdf.Content, &_cce)
	if _bgdd != nil {
		return _ggg.SignatureValidationResult{}, _bgdd
	}
	_eec, _bgdd := _cea(_cce.MessageImprint.HashAlgorithm.Algorithm)
	if _bgdd != nil {
		return _ggg.SignatureValidationResult{}, _bgdd
	}
	_gegg := _eec.New()
	_ecfd, _gae := digest.(*_ge.Buffer)
	if !_gae {
		return _ggg.SignatureValidationResult{}, _df.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_gegg.Write(_ecfd.Bytes())
	_bgb := _gegg.Sum(nil)
	_ffbg := _ggg.SignatureValidationResult{IsSigned: true, IsVerified: _ge.Equal(_bgb, _cce.MessageImprint.HashedMessage), GeneralizedTime: _cce.GeneralizedTime}
	return _ffbg, nil
}

// NewEtsiPAdESLevelLT creates a new Adobe.PPKLite ETSI.CAdES.detached Level LT signature handler.
func NewEtsiPAdESLevelLT(privateKey *_gg.PrivateKey, certificate *_bc.Certificate, caCert *_bc.Certificate, certificateTimestampServerURL string, appender *_ggg.PdfAppender) (_ggg.SignatureHandler, error) {
	_bag := appender.Reader.DSS
	if _bag == nil {
		_bag = _ggg.NewDSS()
	}
	if _bad := _bag.GenerateHashMaps(); _bad != nil {
		return nil, _bad
	}
	return &etsiPAdES{_gge: certificate, _ca: privateKey, _fd: caCert, _ffd: certificateTimestampServerURL, CertClient: _de.NewCertClient(), OCSPClient: _de.NewOCSPClient(), CRLClient: _de.NewCRLClient(), _fca: appender, _bcdf: _bag}, nil
}

// NewEtsiPAdESLevelB creates a new Adobe.PPKLite ETSI.CAdES.detached Level B signature handler.
func NewEtsiPAdESLevelB(privateKey *_gg.PrivateKey, certificate *_bc.Certificate, caCert *_bc.Certificate) (_ggg.SignatureHandler, error) {
	return &etsiPAdES{_gge: certificate, _ca: privateKey, _fd: caCert}, nil
}

// NewDocTimeStampWithOpts returns a new DocTimeStamp configured using the
// specified options. If no options are provided, default options will be used.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
func NewDocTimeStampWithOpts(timestampServerURL string, hashAlgorithm _ffb.Hash, opts *DocTimeStampOpts) (_ggg.SignatureHandler, error) {
	if opts == nil {
		opts = &DocTimeStampOpts{}
	}
	if opts.SignatureSize <= 0 {
		opts.SignatureSize = 4192
	}
	return &docTimeStamp{_acg: timestampServerURL, _eded: hashAlgorithm, _cfd: opts.SignatureSize, _cfga: opts.Client}, nil
}
func (_dgb *docTimeStamp) getCertificate(_fabb *_ggg.PdfSignature) (*_bc.Certificate, error) {
	_cagc, _agcf := _fabb.GetCerts()
	if _agcf != nil {
		return nil, _agcf
	}
	return _cagc[0], nil
}

// ValidateWithOpts validates a PDF signature by checking PdfReader or PdfParser by the DiffPolicy
// params describes parameters for the DocMDP checks.
func (_cc *DocMDPHandler) ValidateWithOpts(sig *_ggg.PdfSignature, digest _ggg.Hasher, params _ggg.SignatureHandlerDocMDPParams) (_ggg.SignatureValidationResult, error) {
	_bb, _gea := _cc._ef.Validate(sig, digest)
	if _gea != nil {
		return _bb, _gea
	}
	_aad := params.Parser
	if _aad == nil {
		return _ggg.SignatureValidationResult{}, _a.New("p\u0061r\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027t\u0020\u0062\u0065\u0020nu\u006c\u006c")
	}
	if !_bb.IsVerified {
		return _bb, nil
	}
	_cb := params.DiffPolicy
	if _cb == nil {
		_cb = _ec.NewDefaultDiffPolicy()
	}
	for _cf := 0; _cf <= _aad.GetRevisionNumber(); _cf++ {
		_geb, _dae := _aad.GetRevision(_cf)
		if _dae != nil {
			return _ggg.SignatureValidationResult{}, _dae
		}
		_bee := _geb.GetTrailer()
		if _bee == nil {
			return _ggg.SignatureValidationResult{}, _a.New("\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064\u0020\u0074\u0068\u0065\u0020\u0074r\u0061i\u006c\u0065\u0072\u0020\u006f\u0062\u006ae\u0063\u0074")
		}
		_gb, _eae := _be.GetDict(_bee.Get("\u0052\u006f\u006f\u0074"))
		if !_eae {
			return _ggg.SignatureValidationResult{}, _a.New("\u0075n\u0064\u0065\u0066\u0069n\u0065\u0064\u0020\u0074\u0068e\u0020r\u006fo\u0074\u0020\u006f\u0062\u006a\u0065\u0063t")
		}
		_bd, _eae := _be.GetDict(_gb.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
		if !_eae {
			continue
		}
		_cfe, _eae := _be.GetArray(_bd.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if !_eae {
			continue
		}
		for _, _fa := range _cfe.Elements() {
			_db, _ead := _be.GetDict(_fa)
			if !_ead {
				continue
			}
			_efd, _ead := _be.GetDict(_db.Get("\u0056"))
			if !_ead {
				continue
			}
			if _be.EqualObjects(_efd.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"), sig.Contents) {
				_bb.DiffResults, _dae = _cb.ReviewFile(_geb, _aad, &_ec.MDPParameters{DocMDPLevel: _cc.Permission})
				if _dae != nil {
					return _ggg.SignatureValidationResult{}, _dae
				}
				_bb.IsVerified = _bb.DiffResults.IsPermitted()
				return _bb, nil
			}
		}
	}
	return _ggg.SignatureValidationResult{}, _a.New("\u0064\u006f\u006e\u0027\u0074\u0020\u0066o\u0075\u006e\u0064 \u0074\u0068\u0069\u0073 \u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073")
}

// RevocationInfoArchival is OIDAttributeAdobeRevocation attribute.
type RevocationInfoArchival struct {
	Crl          []_b.RawValue `asn1:"explicit,tag:0,optional"`
	Ocsp         []_b.RawValue `asn1:"explicit,tag:1,optional"`
	OtherRevInfo []_b.RawValue `asn1:"explicit,tag:2,optional"`
}

// NewEmptyAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached
// signature handler. The generated signature is empty and of size signatureLen.
// The signatureLen parameter can be 0 for the signature validation.
func NewEmptyAdobePKCS7Detached(signatureLen int) (_ggg.SignatureHandler, error) {
	return &adobePKCS7Detached{_ddbab: true, _gddd: signatureLen}, nil
}

// Validate implementation of the SignatureHandler interface
// This check is impossible without checking the document's content.
// Please, use ValidateWithOpts with the PdfParser.
func (_ed *DocMDPHandler) Validate(sig *_ggg.PdfSignature, digest _ggg.Hasher) (_ggg.SignatureValidationResult, error) {
	return _ggg.SignatureValidationResult{}, _a.New("i\u006d\u0070\u006f\u0073\u0073\u0069b\u006c\u0065\u0020\u0076\u0061\u006ci\u0064\u0061\u0074\u0069\u006f\u006e\u0020w\u0069\u0074\u0068\u006f\u0075\u0074\u0020\u0070\u0061\u0072s\u0065")
}

// SignFunc represents a custom signing function. The function should return
// the computed signature.
type SignFunc func(_dcc *_ggg.PdfSignature, _gdac _ggg.Hasher) ([]byte, error)

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_dcd *adobeX509RSASHA1) IsApplicable(sig *_ggg.PdfSignature) bool {
	if sig == nil || sig.Filter == nil || sig.SubFilter == nil {
		return false
	}
	return (*sig.Filter == "A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053" || *sig.Filter == "\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065") && *sig.SubFilter == "\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031"
}
func (_beb *adobePKCS7Detached) getCertificate(_dga *_ggg.PdfSignature) (*_bc.Certificate, error) {
	if _beb._efe != nil {
		return _beb._efe, nil
	}
	_gda, _fad := _dga.GetCerts()
	if _fad != nil {
		return nil, _fad
	}
	return _gda[0], nil
}

// DocMDPHandler describes handler for the DocMDP realization.
type DocMDPHandler struct {
	_ef        _ggg.SignatureHandler
	Permission _ec.DocMDPPermission
}

func _gbd(_ega *_gg.PublicKey, _ebd []byte) _ffb.Hash {
	_feb := _ega.Size()
	if _feb != len(_ebd) {
		return 0
	}
	_dda := func(_eaca *_f.Int, _fdgcc *_gg.PublicKey, _baf *_f.Int) *_f.Int {
		_fagd := _f.NewInt(int64(_fdgcc.E))
		_eaca.Exp(_baf, _fagd, _fdgcc.N)
		return _eaca
	}
	_gcb := new(_f.Int).SetBytes(_ebd)
	_aegb := _dda(new(_f.Int), _ega, _gcb)
	_egef := _edea(_aegb.Bytes(), _feb)
	if _egef[0] != 0 || _egef[1] != 1 {
		return 0
	}
	_gfe := []struct {
		Hash   _ffb.Hash
		Prefix []byte
	}{{Hash: _ffb.SHA1, Prefix: []byte{0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14}}, {Hash: _ffb.SHA256, Prefix: []byte{0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20}}, {Hash: _ffb.SHA384, Prefix: []byte{0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30}}, {Hash: _ffb.SHA512, Prefix: []byte{0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40}}, {Hash: _ffb.RIPEMD160, Prefix: []byte{0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14}}}
	for _, _bdce := range _gfe {
		_beg := _bdce.Hash.Size()
		_abb := len(_bdce.Prefix) + _beg
		if _ge.Equal(_egef[_feb-_abb:_feb-_beg], _bdce.Prefix) {
			return _bdce.Hash
		}
	}
	return 0
}

// Validate validates PdfSignature.
func (_cbg *etsiPAdES) Validate(sig *_ggg.PdfSignature, digest _ggg.Hasher) (_ggg.SignatureValidationResult, error) {
	_dec := sig.Contents.Bytes()
	_fae, _gddf := _bg.Parse(_dec)
	if _gddf != nil {
		return _ggg.SignatureValidationResult{}, _gddf
	}
	_fdg, _bgd := digest.(*_ge.Buffer)
	if !_bgd {
		return _ggg.SignatureValidationResult{}, _df.Errorf("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073")
	}
	_fae.Content = _fdg.Bytes()
	if _gddf = _fae.Verify(); _gddf != nil {
		return _ggg.SignatureValidationResult{}, _gddf
	}
	_fge := false
	_fdgc := false
	var _ecfa _ff.Time
	for _, _bcg := range _fae.Signers {
		_decc := _bcg.EncryptedDigest
		var _cabb RevocationInfoArchival
		_gddf = _fae.UnmarshalSignedAttribute(_bg.OIDAttributeAdobeRevocation, &_cabb)
		if _gddf == nil {
			if len(_cabb.Crl) > 0 {
				_fdgc = true
			}
			if len(_cabb.Ocsp) > 0 {
				_fge = true
			}
		}
		for _, _eef := range _bcg.UnauthenticatedAttributes {
			if _eef.Type.Equal(_bg.OIDAttributeTimeStampToken) {
				_gca, _gde := _da.Parse(_eef.Value.Bytes)
				if _gde != nil {
					return _ggg.SignatureValidationResult{}, _gde
				}
				_ecfa = _gca.Time
				_ege := _gca.HashAlgorithm.New()
				_ege.Write(_decc)
				if !_ge.Equal(_ege.Sum(nil), _gca.HashedMessage) {
					return _ggg.SignatureValidationResult{}, _df.Errorf("\u0048\u0061\u0073\u0068\u0020i\u006e\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061\u006d\u0070\u0020\u0069s\u0020\u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020\u0066\u0072\u006f\u006d\u0020\u0070\u006b\u0063\u0073\u0037")
				}
				break
			}
		}
	}
	_caf := _ggg.SignatureValidationResult{IsSigned: true, IsVerified: true, IsCrlFound: _fdgc, IsOcspFound: _fge, GeneralizedTime: _ecfa}
	return _caf, nil
}

type timestampInfo struct {
	Version        int
	Policy         _b.RawValue
	MessageImprint struct {
		HashAlgorithm _dfc.AlgorithmIdentifier
		HashedMessage []byte
	}
	SerialNumber    _b.RawValue
	GeneralizedTime _ff.Time
}

func _edea(_agce []byte, _baff int) (_fbbc []byte) {
	_gce := len(_agce)
	if _gce > _baff {
		_gce = _baff
	}
	_fbbc = make([]byte, _baff)
	copy(_fbbc[len(_fbbc)-_gce:], _agce)
	return
}

// InitSignature initialises the PdfSignature.
func (_fag *adobePKCS7Detached) InitSignature(sig *_ggg.PdfSignature) error {
	if !_fag._ddbab {
		if _fag._efe == nil {
			return _a.New("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c")
		}
		if _fag._acd == nil {
			return _a.New("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c")
		}
	}
	_acdd := *_fag
	sig.Handler = &_acdd
	sig.Filter = _be.MakeName("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")
	sig.SubFilter = _be.MakeName("\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064")
	sig.Reference = nil
	_deg, _ebc := _acdd.NewDigest(sig)
	if _ebc != nil {
		return _ebc
	}
	_deg.Write([]byte("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"))
	return _acdd.Sign(sig, _deg)
}
