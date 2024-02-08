package sigutil

import (
	_a "bytes"
	_ba "crypto"
	_ae "crypto/x509"
	_aa "encoding/asn1"
	_ag "encoding/pem"
	_fb "errors"
	_bg "fmt"
	_g "io"
	_d "io/ioutil"
	_c "net/http"
	_b "time"

	_ce "bitbucket.org/shenghui0779/gopdf/common"
	_af "github.com/unidoc/timestamp"
	_ab "golang.org/x/crypto/ocsp"
)

// NewTimestampClient returns a new timestamp client.
func NewTimestampClient() *TimestampClient { return &TimestampClient{HTTPClient: _cde()} }

// GetIssuer retrieves the issuer of the provided certificate.
func (_gg *CertClient) GetIssuer(cert *_ae.Certificate) (*_ae.Certificate, error) {
	for _, _abd := range cert.IssuingCertificateURL {
		_e, _dg := _gg.Get(_abd)
		if _dg != nil {
			_ce.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074 \u0064\u006f\u0077\u006e\u006c\u006f\u0061\u0064\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0066\u006f\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0025\u0076\u003a\u0020\u0025\u0076", cert.Subject.CommonName, _dg)
			continue
		}
		return _e, nil
	}
	return nil, _bg.Errorf("\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063a\u0074e\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064")
}

// MakeRequest makes a CRL request to the specified server and returns the
// response. If a server URL is not provided, it is extracted from the certificate.
func (_cg *CRLClient) MakeRequest(serverURL string, cert *_ae.Certificate) ([]byte, error) {
	if _cg.HTTPClient == nil {
		_cg.HTTPClient = _cde()
	}
	if serverURL == "" {
		if len(cert.CRLDistributionPoints) == 0 {
			return nil, _fb.New("\u0063e\u0072\u0074i\u0066\u0069\u0063\u0061t\u0065\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070ec\u0069\u0066\u0079 \u0061\u006ey\u0020\u0043\u0052\u004c\u0020\u0073e\u0072\u0076e\u0072\u0073")
		}
		serverURL = cert.CRLDistributionPoints[0]
	}
	_ad, _fbb := _cg.HTTPClient.Get(serverURL)
	if _fbb != nil {
		return nil, _fbb
	}
	defer _ad.Body.Close()
	_ac, _fbb := _d.ReadAll(_ad.Body)
	if _fbb != nil {
		return nil, _fbb
	}
	if _fbf, _ := _ag.Decode(_ac); _fbf != nil {
		_ac = _fbf.Bytes
	}
	return _ac, nil
}

// TimestampClient represents a RFC 3161 timestamp client.
// It is used to obtain signed tokens from timestamp authority servers.
type TimestampClient struct {

	// HTTPClient is the HTTP client used to make timestamp requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_c.Client

	// Callbacks.
	BeforeHTTPRequest func(_aea *_c.Request) error
}

// Get retrieves the certificate at the specified URL.
func (_ga *CertClient) Get(url string) (*_ae.Certificate, error) {
	if _ga.HTTPClient == nil {
		_ga.HTTPClient = _cde()
	}
	_gf, _afe := _ga.HTTPClient.Get(url)
	if _afe != nil {
		return nil, _afe
	}
	defer _gf.Body.Close()
	_ca, _afe := _d.ReadAll(_gf.Body)
	if _afe != nil {
		return nil, _afe
	}
	if _gag, _ := _ag.Decode(_ca); _gag != nil {
		_ca = _gag.Bytes
	}
	_cd, _afe := _ae.ParseCertificate(_ca)
	if _afe != nil {
		return nil, _afe
	}
	return _cd, nil
}

// NewCRLClient returns a new CRL client.
func NewCRLClient() *CRLClient { return &CRLClient{HTTPClient: _cde()} }

// CRLClient represents a CRL (Certificate revocation list) client.
// It is used to request revocation data from CRL servers.
type CRLClient struct {

	// HTTPClient is the HTTP client used to make CRL requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_c.Client
}

// IsCA returns true if the provided certificate appears to be a CA certificate.
func (_gdc *CertClient) IsCA(cert *_ae.Certificate) bool {
	return cert.IsCA && _a.Equal(cert.RawIssuer, cert.RawSubject)
}

// MakeRequest makes a OCSP request to the specified server and returns
// the parsed and raw responses. If a server URL is not provided, it is
// extracted from the certificate.
func (_ceg *OCSPClient) MakeRequest(serverURL string, cert, issuer *_ae.Certificate) (*_ab.Response, []byte, error) {
	if _ceg.HTTPClient == nil {
		_ceg.HTTPClient = _cde()
	}
	if serverURL == "" {
		if len(cert.OCSPServer) == 0 {
			return nil, nil, _fb.New("\u0063e\u0072\u0074i\u0066\u0069\u0063a\u0074\u0065\u0020\u0064\u006f\u0065\u0073 \u006e\u006f\u0074\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0079\u0020\u004f\u0043S\u0050\u0020\u0073\u0065\u0072\u0076\u0065\u0072\u0073")
		}
		serverURL = cert.OCSPServer[0]
	}
	_db, _bgg := _ab.CreateRequest(cert, issuer, &_ab.RequestOptions{Hash: _ceg.Hash})
	if _bgg != nil {
		return nil, nil, _bgg
	}
	_cec, _bgg := _ceg.HTTPClient.Post(serverURL, "\u0061p\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u006fc\u0073\u0070\u002d\u0072\u0065\u0071\u0075\u0065\u0073\u0074", _a.NewReader(_db))
	if _bgg != nil {
		return nil, nil, _bgg
	}
	defer _cec.Body.Close()
	_bc, _bgg := _d.ReadAll(_cec.Body)
	if _bgg != nil {
		return nil, nil, _bgg
	}
	if _fg, _ := _ag.Decode(_bc); _fg != nil {
		_bc = _fg.Bytes
	}
	_gdcf, _bgg := _ab.ParseResponseForCert(_bc, cert, issuer)
	if _bgg != nil {
		return nil, nil, _bgg
	}
	return _gdcf, _bc, nil
}

// NewCertClient returns a new certificate client.
func NewCertClient() *CertClient { return &CertClient{HTTPClient: _cde()} }

// NewTimestampRequest returns a new timestamp request based
// on the specified options.
func NewTimestampRequest(body _g.Reader, opts *_af.RequestOptions) (*_af.Request, error) {
	if opts == nil {
		opts = &_af.RequestOptions{}
	}
	if opts.Hash == 0 {
		opts.Hash = _ba.SHA256
	}
	if !opts.Hash.Available() {
		return nil, _ae.ErrUnsupportedAlgorithm
	}
	_bca := opts.Hash.New()
	if _, _fe := _g.Copy(_bca, body); _fe != nil {
		return nil, _fe
	}
	return &_af.Request{HashAlgorithm: opts.Hash, HashedMessage: _bca.Sum(nil), Certificates: opts.Certificates, TSAPolicyOID: opts.TSAPolicyOID, Nonce: opts.Nonce}, nil
}

// CertClient represents a X.509 certificate client. Its primary purpose
// is to download certificates.
type CertClient struct {

	// HTTPClient is the HTTP client used to make certificate requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_c.Client
}

// GetEncodedToken executes the timestamp request and returns the DER encoded
// timestamp token bytes.
func (_fa *TimestampClient) GetEncodedToken(serverURL string, req *_af.Request) ([]byte, error) {
	if serverURL == "" {
		return nil, _bg.Errorf("\u006d\u0075\u0073\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061m\u0070\u0020\u0073\u0065\u0072\u0076\u0065r\u0020\u0055\u0052\u004c")
	}
	if req == nil {
		return nil, _bg.Errorf("\u0074\u0069\u006de\u0073\u0074\u0061\u006dp\u0020\u0072\u0065\u0071\u0075\u0065\u0073t\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	_ggf, _caa := req.Marshal()
	if _caa != nil {
		return nil, _caa
	}
	_caf, _caa := _c.NewRequest("\u0050\u004f\u0053\u0054", serverURL, _a.NewBuffer(_ggf))
	if _caa != nil {
		return nil, _caa
	}
	_caf.Header.Set("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065", "a\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0074\u0069\u006d\u0065\u0073t\u0061\u006d\u0070-\u0071u\u0065\u0072\u0079")
	if _fa.BeforeHTTPRequest != nil {
		if _gb := _fa.BeforeHTTPRequest(_caf); _gb != nil {
			return nil, _gb
		}
	}
	_ge := _fa.HTTPClient
	if _ge == nil {
		_ge = _cde()
	}
	_fc, _caa := _ge.Do(_caf)
	if _caa != nil {
		return nil, _caa
	}
	defer _fc.Body.Close()
	_bag, _caa := _d.ReadAll(_fc.Body)
	if _caa != nil {
		return nil, _caa
	}
	if _fc.StatusCode != _c.StatusOK {
		return nil, _bg.Errorf("\u0075\u006e\u0065x\u0070\u0065\u0063\u0074e\u0064\u0020\u0048\u0054\u0054\u0050\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0064", _fc.StatusCode)
	}
	var _ed struct {
		Version _aa.RawValue
		Content _aa.RawValue
	}
	if _, _caa = _aa.Unmarshal(_bag, &_ed); _caa != nil {
		return nil, _caa
	}
	return _ed.Content.FullBytes, nil
}

// NewOCSPClient returns a new OCSP client.
func NewOCSPClient() *OCSPClient { return &OCSPClient{HTTPClient: _cde(), Hash: _ba.SHA1} }

// OCSPClient represents a OCSP (Online Certificate Status Protocol) client.
// It is used to request revocation data from OCSP servers.
type OCSPClient struct {

	// HTTPClient is the HTTP client used to make OCSP requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_c.Client

	// Hash is the hash function  used when constructing the OCSP
	// requests. If zero, SHA-1 will be used.
	Hash _ba.Hash
}

func _cde() *_c.Client { return &_c.Client{Timeout: 5 * _b.Second} }
