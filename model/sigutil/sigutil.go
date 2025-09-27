package sigutil

import (
	_g "bytes"
	_gc "crypto"
	_bda "crypto/x509"
	_c "encoding/asn1"
	_e "encoding/pem"
	_bg "errors"
	_de "fmt"
	_bd "io"
	_bge "net/http"
	_f "time"

	_fe "unitechio/gopdf/gopdf/common"
	_gg "github.com/unidoc/timestamp"
	_fc "golang.org/x/crypto/ocsp"
)

// Get retrieves the certificate at the specified URL.
func (_fa *CertClient) Get(url string) (*_bda.Certificate, error) {
	if _fa.HTTPClient == nil {
		_fa.HTTPClient = _bcb()
	}
	_ca, _eb := _fa.HTTPClient.Get(url)
	if _eb != nil {
		return nil, _eb
	}
	defer _ca.Body.Close()
	_bc, _eb := _bd.ReadAll(_ca.Body)
	if _eb != nil {
		return nil, _eb
	}
	if _ggd, _ := _e.Decode(_bc); _ggd != nil {
		_bc = _ggd.Bytes
	}
	_cc, _eb := _bda.ParseCertificate(_bc)
	if _eb != nil {
		return nil, _eb
	}
	return _cc, nil
}

// MakeRequest makes a CRL request to the specified server and returns the
// response. If a server URL is not provided, it is extracted from the certificate.
func (_ec *CRLClient) MakeRequest(serverURL string, cert *_bda.Certificate) ([]byte, error) {
	if _ec.HTTPClient == nil {
		_ec.HTTPClient = _bcb()
	}
	if serverURL == "" {
		if len(cert.CRLDistributionPoints) == 0 {
			return nil, _bg.New("\u0063e\u0072\u0074i\u0066\u0069\u0063\u0061t\u0065\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070ec\u0069\u0066\u0079 \u0061\u006ey\u0020\u0043\u0052\u004c\u0020\u0073e\u0072\u0076e\u0072\u0073")
		}
		serverURL = cert.CRLDistributionPoints[0]
	}
	_ae, _ebc := _ec.HTTPClient.Get(serverURL)
	if _ebc != nil {
		return nil, _ebc
	}
	defer _ae.Body.Close()
	_bce, _ebc := _bd.ReadAll(_ae.Body)
	if _ebc != nil {
		return nil, _ebc
	}
	if _gcg, _ := _e.Decode(_bce); _gcg != nil {
		_bce = _gcg.Bytes
	}
	return _bce, nil
}

// OCSPClient represents a OCSP (Online Certificate Status Protocol) client.
// It is used to request revocation data from OCSP servers.
type OCSPClient struct {
	// HTTPClient is the HTTP client used to make OCSP requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_bge.Client

	// Hash is the hash function  used when constructing the OCSP
	// requests. If zero, SHA-1 will be used.
	Hash _gc.Hash
}

// NewTimestampClient returns a new timestamp client.
func NewTimestampClient() *TimestampClient { return &TimestampClient{HTTPClient: _bcb()} }

// NewCRLClient returns a new CRL client.
func NewCRLClient() *CRLClient { return &CRLClient{HTTPClient: _bcb()} }

// MakeRequest makes a OCSP request to the specified server and returns
// the parsed and raw responses. If a server URL is not provided, it is
// extracted from the certificate.
func (_ff *OCSPClient) MakeRequest(serverURL string, cert, issuer *_bda.Certificate) (*_fc.Response, []byte, error) {
	if _ff.HTTPClient == nil {
		_ff.HTTPClient = _bcb()
	}
	if serverURL == "" {
		if len(cert.OCSPServer) == 0 {
			return nil, nil, _bg.New("\u0063e\u0072\u0074i\u0066\u0069\u0063a\u0074\u0065\u0020\u0064\u006f\u0065\u0073 \u006e\u006f\u0074\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0079\u0020\u004f\u0043S\u0050\u0020\u0073\u0065\u0072\u0076\u0065\u0072\u0073")
		}
		serverURL = cert.OCSPServer[0]
	}
	_ed, _fg := _fc.CreateRequest(cert, issuer, &_fc.RequestOptions{Hash: _ff.Hash})
	if _fg != nil {
		return nil, nil, _fg
	}
	_ef, _fg := _ff.HTTPClient.Post(serverURL, "\u0061p\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u006fc\u0073\u0070\u002d\u0072\u0065\u0071\u0075\u0065\u0073\u0074", _g.NewReader(_ed))
	if _fg != nil {
		return nil, nil, _fg
	}
	defer _ef.Body.Close()
	_dff, _fg := _bd.ReadAll(_ef.Body)
	if _fg != nil {
		return nil, nil, _fg
	}
	if _ab, _ := _e.Decode(_dff); _ab != nil {
		_dff = _ab.Bytes
	}
	_ag, _fg := _fc.ParseResponseForCert(_dff, cert, issuer)
	if _fg != nil {
		return nil, nil, _fg
	}
	return _ag, _dff, nil
}

// CRLClient represents a CRL (Certificate revocation list) client.
// It is used to request revocation data from CRL servers.
type CRLClient struct {
	// HTTPClient is the HTTP client used to make CRL requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_bge.Client
}

// NewCertClient returns a new certificate client.
func NewCertClient() *CertClient { return &CertClient{HTTPClient: _bcb()} }

// CertClient represents a X.509 certificate client. Its primary purpose
// is to download certificates.
type CertClient struct {
	// HTTPClient is the HTTP client used to make certificate requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_bge.Client
}

// GetEncodedToken executes the timestamp request and returns the DER encoded
// timestamp token bytes.
func (_cbc *TimestampClient) GetEncodedToken(serverURL string, req *_gg.Request) ([]byte, error) {
	if serverURL == "" {
		return nil, _de.Errorf("\u006d\u0075\u0073\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061m\u0070\u0020\u0073\u0065\u0072\u0076\u0065r\u0020\u0055\u0052\u004c")
	}
	if req == nil {
		return nil, _de.Errorf("\u0074\u0069\u006de\u0073\u0074\u0061\u006dp\u0020\u0072\u0065\u0071\u0075\u0065\u0073t\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	_dc, _aff := req.Marshal()
	if _aff != nil {
		return nil, _aff
	}
	_gd, _aff := _bge.NewRequest("\u0050\u004f\u0053\u0054", serverURL, _g.NewBuffer(_dc))
	if _aff != nil {
		return nil, _aff
	}
	_gd.Header.Set("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065", "a\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0074\u0069\u006d\u0065\u0073t\u0061\u006d\u0070-\u0071u\u0065\u0072\u0079")
	if _cbc.BeforeHTTPRequest != nil {
		if _aeg := _cbc.BeforeHTTPRequest(_gd); _aeg != nil {
			return nil, _aeg
		}
	}
	_ccd := _cbc.HTTPClient
	if _ccd == nil {
		_ccd = _bcb()
	}
	_be, _aff := _ccd.Do(_gd)
	if _aff != nil {
		return nil, _aff
	}
	defer _be.Body.Close()
	_gde, _aff := _bd.ReadAll(_be.Body)
	if _aff != nil {
		return nil, _aff
	}
	if _be.StatusCode != _bge.StatusOK {
		return nil, _de.Errorf("\u0075\u006e\u0065x\u0070\u0065\u0063\u0074e\u0064\u0020\u0048\u0054\u0054\u0050\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0064", _be.StatusCode)
	}
	var _cbb struct {
		Version _c.RawValue
		Content _c.RawValue
	}
	if _, _aff = _c.Unmarshal(_gde, &_cbb); _aff != nil {
		return nil, _aff
	}
	return _cbb.Content.FullBytes, nil
}

// TimestampClient represents a RFC 3161 timestamp client.
// It is used to obtain signed tokens from timestamp authority servers.
type TimestampClient struct {
	// HTTPClient is the HTTP client used to make timestamp requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_bge.Client

	// Callbacks.
	BeforeHTTPRequest func(_bde *_bge.Request) error
}

func _bcb() *_bge.Client { return &_bge.Client{Timeout: 5 * _f.Second} }

// NewOCSPClient returns a new OCSP client.
func NewOCSPClient() *OCSPClient { return &OCSPClient{HTTPClient: _bcb(), Hash: _gc.SHA1} }

// IsCA returns true if the provided certificate appears to be a CA certificate.
func (_a *CertClient) IsCA(cert *_bda.Certificate) bool {
	return cert.IsCA && _g.Equal(cert.RawIssuer, cert.RawSubject)
}

// GetIssuer retrieves the issuer of the provided certificate.
func (_cb *CertClient) GetIssuer(cert *_bda.Certificate) (*_bda.Certificate, error) {
	for _, _cd := range cert.IssuingCertificateURL {
		_ba, _dd := _cb.Get(_cd)
		if _dd != nil {
			_fe.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074 \u0064\u006f\u0077\u006e\u006c\u006f\u0061\u0064\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0066\u006f\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0025\u0076\u003a\u0020\u0025\u0076", cert.Subject.CommonName, _dd)
			continue
		}
		return _ba, nil
	}
	return nil, _de.Errorf("\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063a\u0074e\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064")
}

// NewTimestampRequest returns a new timestamp request based
// on the specified options.
func NewTimestampRequest(body _bd.Reader, opts *_gg.RequestOptions) (*_gg.Request, error) {
	if opts == nil {
		opts = &_gg.RequestOptions{}
	}
	if opts.Hash == 0 {
		opts.Hash = _gc.SHA256
	}
	if !opts.Hash.Available() {
		return nil, _bda.ErrUnsupportedAlgorithm
	}
	_af := opts.Hash.New()
	if _, _ac := _bd.Copy(_af, body); _ac != nil {
		return nil, _ac
	}
	return &_gg.Request{HashAlgorithm: opts.Hash, HashedMessage: _af.Sum(nil), Certificates: opts.Certificates, TSAPolicyOID: opts.TSAPolicyOID, Nonce: opts.Nonce}, nil
}
