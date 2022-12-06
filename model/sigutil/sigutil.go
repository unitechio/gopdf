package sigutil

import (
	_b "bytes"
	_a "crypto"
	_d "crypto/x509"
	_ca "encoding/asn1"
	_ga "encoding/pem"
	_c "errors"
	_cf "fmt"
	_ea "io"
	_g "io/ioutil"
	_f "net/http"
	_be "time"

	_gc "bitbucket.org/shenghui0779/gopdf/common"
	_eg "github.com/unidoc/timestamp"
	_gf "golang.org/x/crypto/ocsp"
)

// IsCA returns true if the provided certificate appears to be a CA certificate.
func (_df *CertClient) IsCA(cert *_d.Certificate) bool {
	return cert.IsCA && _b.Equal(cert.RawIssuer, cert.RawSubject)
}

// GetEncodedToken executes the timestamp request and returns the DER encoded
// timestamp token bytes.
func (_cfg *TimestampClient) GetEncodedToken(serverURL string, req *_eg.Request) ([]byte, error) {
	if serverURL == "" {
		return nil, _cf.Errorf("\u006d\u0075\u0073\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061m\u0070\u0020\u0073\u0065\u0072\u0076\u0065r\u0020\u0055\u0052\u004c")
	}
	if req == nil {
		return nil, _cf.Errorf("\u0074\u0069\u006de\u0073\u0074\u0061\u006dp\u0020\u0072\u0065\u0071\u0075\u0065\u0073t\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	_fe, _gfg := req.Marshal()
	if _gfg != nil {
		return nil, _gfg
	}
	_bd, _gfg := _f.NewRequest("\u0050\u004f\u0053\u0054", serverURL, _b.NewBuffer(_fe))
	if _gfg != nil {
		return nil, _gfg
	}
	_bd.Header.Set("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065", "a\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0074\u0069\u006d\u0065\u0073t\u0061\u006d\u0070-\u0071u\u0065\u0072\u0079")
	if _cfg.BeforeHTTPRequest != nil {
		if _beba := _cfg.BeforeHTTPRequest(_bd); _beba != nil {
			return nil, _beba
		}
	}
	_cac := _cfg.HTTPClient
	if _cac == nil {
		_cac = _bff()
	}
	_ae, _gfg := _cac.Do(_bd)
	if _gfg != nil {
		return nil, _gfg
	}
	defer _ae.Body.Close()
	_db, _gfg := _g.ReadAll(_ae.Body)
	if _gfg != nil {
		return nil, _gfg
	}
	if _ae.StatusCode != _f.StatusOK {
		return nil, _cf.Errorf("\u0075\u006e\u0065x\u0070\u0065\u0063\u0074e\u0064\u0020\u0048\u0054\u0054\u0050\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0064", _ae.StatusCode)
	}
	var _da struct {
		Version _ca.RawValue
		Content _ca.RawValue
	}
	if _, _gfg = _ca.Unmarshal(_db, &_da); _gfg != nil {
		return nil, _gfg
	}
	return _da.Content.FullBytes, nil
}

// CRLClient represents a CRL (Certificate revocation list) client.
// It is used to request revocation data from CRL servers.
type CRLClient struct {

	// HTTPClient is the HTTP client used to make CRL requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_f.Client
}

// MakeRequest makes a CRL request to the specified server and returns the
// response. If a server URL is not provided, it is extracted from the certificate.
func (_ffa *CRLClient) MakeRequest(serverURL string, cert *_d.Certificate) ([]byte, error) {
	if _ffa.HTTPClient == nil {
		_ffa.HTTPClient = _bff()
	}
	if serverURL == "" {
		if len(cert.CRLDistributionPoints) == 0 {
			return nil, _c.New("\u0063e\u0072\u0074i\u0066\u0069\u0063\u0061t\u0065\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070ec\u0069\u0066\u0079 \u0061\u006ey\u0020\u0043\u0052\u004c\u0020\u0073e\u0072\u0076e\u0072\u0073")
		}
		serverURL = cert.CRLDistributionPoints[0]
	}
	_ge, _cd := _ffa.HTTPClient.Get(serverURL)
	if _cd != nil {
		return nil, _cd
	}
	defer _ge.Body.Close()
	_de, _cd := _g.ReadAll(_ge.Body)
	if _cd != nil {
		return nil, _cd
	}
	if _cff, _ := _ga.Decode(_de); _cff != nil {
		_de = _cff.Bytes
	}
	return _de, nil
}

// TimestampClient represents a RFC 3161 timestamp client.
// It is used to obtain signed tokens from timestamp authority servers.
type TimestampClient struct {

	// HTTPClient is the HTTP client used to make timestamp requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_f.Client

	// Callbacks.
	BeforeHTTPRequest func(_bcb *_f.Request) error
}

// MakeRequest makes a OCSP request to the specified server and returns
// the parsed and raw responses. If a server URL is not provided, it is
// extracted from the certificate.
func (_ad *OCSPClient) MakeRequest(serverURL string, cert, issuer *_d.Certificate) (*_gf.Response, []byte, error) {
	if _ad.HTTPClient == nil {
		_ad.HTTPClient = _bff()
	}
	if serverURL == "" {
		if len(cert.OCSPServer) == 0 {
			return nil, nil, _c.New("\u0063e\u0072\u0074i\u0066\u0069\u0063a\u0074\u0065\u0020\u0064\u006f\u0065\u0073 \u006e\u006f\u0074\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0079\u0020\u004f\u0043S\u0050\u0020\u0073\u0065\u0072\u0076\u0065\u0072\u0073")
		}
		serverURL = cert.OCSPServer[0]
	}
	_agg, _cc := _gf.CreateRequest(cert, issuer, &_gf.RequestOptions{Hash: _ad.Hash})
	if _cc != nil {
		return nil, nil, _cc
	}
	_fg, _cc := _ad.HTTPClient.Post(serverURL, "\u0061p\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u006fc\u0073\u0070\u002d\u0072\u0065\u0071\u0075\u0065\u0073\u0074", _b.NewReader(_agg))
	if _cc != nil {
		return nil, nil, _cc
	}
	defer _fg.Body.Close()
	_ab, _cc := _g.ReadAll(_fg.Body)
	if _cc != nil {
		return nil, nil, _cc
	}
	if _eb, _ := _ga.Decode(_ab); _eb != nil {
		_ab = _eb.Bytes
	}
	_bc, _cc := _gf.ParseResponseForCert(_ab, cert, issuer)
	if _cc != nil {
		return nil, nil, _cc
	}
	return _bc, _ab, nil
}

// OCSPClient represents a OCSP (Online Certificate Status Protocol) client.
// It is used to request revocation data from OCSP servers.
type OCSPClient struct {

	// HTTPClient is the HTTP client used to make OCSP requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_f.Client

	// Hash is the hash function  used when constructing the OCSP
	// requests. If zero, SHA-1 will be used.
	Hash _a.Hash
}

// CertClient represents a X.509 certificate client. Its primary purpose
// is to download certificates.
type CertClient struct {

	// HTTPClient is the HTTP client used to make certificate requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_f.Client
}

// NewTimestampClient returns a new timestamp client.
func NewTimestampClient() *TimestampClient { return &TimestampClient{HTTPClient: _bff()} }

// NewTimestampRequest returns a new timestamp request based
// on the specified options.
func NewTimestampRequest(body _ea.Reader, opts *_eg.RequestOptions) (*_eg.Request, error) {
	if opts == nil {
		opts = &_eg.RequestOptions{}
	}
	if opts.Hash == 0 {
		opts.Hash = _a.SHA256
	}
	if !opts.Hash.Available() {
		return nil, _d.ErrUnsupportedAlgorithm
	}
	_dgaa := opts.Hash.New()
	if _, _ce := _ea.Copy(_dgaa, body); _ce != nil {
		return nil, _ce
	}
	return &_eg.Request{HashAlgorithm: opts.Hash, HashedMessage: _dgaa.Sum(nil), Certificates: opts.Certificates, TSAPolicyOID: opts.TSAPolicyOID, Nonce: opts.Nonce}, nil
}

// NewCRLClient returns a new CRL client.
func NewCRLClient() *CRLClient { return &CRLClient{HTTPClient: _bff()} }

// NewCertClient returns a new certificate client.
func NewCertClient() *CertClient { return &CertClient{HTTPClient: _bff()} }

// Get retrieves the certificate at the specified URL.
func (_ag *CertClient) Get(url string) (*_d.Certificate, error) {
	if _ag.HTTPClient == nil {
		_ag.HTTPClient = _bff()
	}
	_gd, _ef := _ag.HTTPClient.Get(url)
	if _ef != nil {
		return nil, _ef
	}
	defer _gd.Body.Close()
	_dc, _ef := _g.ReadAll(_gd.Body)
	if _ef != nil {
		return nil, _ef
	}
	if _fd, _ := _ga.Decode(_dc); _fd != nil {
		_dc = _fd.Bytes
	}
	_dg, _ef := _d.ParseCertificate(_dc)
	if _ef != nil {
		return nil, _ef
	}
	return _dg, nil
}
func _bff() *_f.Client { return &_f.Client{Timeout: 5 * _be.Second} }

// NewOCSPClient returns a new OCSP client.
func NewOCSPClient() *OCSPClient { return &OCSPClient{HTTPClient: _bff(), Hash: _a.SHA1} }

// GetIssuer retrieves the issuer of the provided certificate.
func (_beb *CertClient) GetIssuer(cert *_d.Certificate) (*_d.Certificate, error) {
	for _, _ff := range cert.IssuingCertificateURL {
		_fb, _dd := _beb.Get(_ff)
		if _dd != nil {
			_gc.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074 \u0064\u006f\u0077\u006e\u006c\u006f\u0061\u0064\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0066\u006f\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0025\u0076\u003a\u0020\u0025\u0076", cert.Subject.CommonName, _dd)
			continue
		}
		return _fb, nil
	}
	return nil, _cf.Errorf("\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063a\u0074e\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064")
}
