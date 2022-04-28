package sigutil

import (
	_d "bytes"
	_e "crypto"
	_ge "crypto/x509"
	_ca "encoding/asn1"
	_ef "encoding/pem"
	_g "errors"
	_ad "fmt"
	_f "io"
	_ab "io/ioutil"
	_df "net/http"
	_c "time"

	_gea "bitbucket.org/shenghui0779/gopdf/common"
	_b "github.com/unidoc/timestamp"
	_ag "golang.org/x/crypto/ocsp"
)

// TimestampClient represents a RFC 3161 timestamp client.
// It is used to obtain signed tokens from timestamp authority servers.
type TimestampClient struct {

	// HTTPClient is the HTTP client used to make timestamp requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_df.Client
}

// GetEncodedToken executes the timestamp request and returns the DER encoded
// timestamp token bytes.
func (_bg *TimestampClient) GetEncodedToken(serverURL string, req *_b.Request) ([]byte, error) {
	if serverURL == "" {
		return nil, _ad.Errorf("\u006d\u0075\u0073\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061m\u0070\u0020\u0073\u0065\u0072\u0076\u0065r\u0020\u0055\u0052\u004c")
	}
	if req == nil {
		return nil, _ad.Errorf("\u0074\u0069\u006de\u0073\u0074\u0061\u006dp\u0020\u0072\u0065\u0071\u0075\u0065\u0073t\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c")
	}
	_afa, _cbd := req.Marshal()
	if _cbd != nil {
		return nil, _cbd
	}
	_dff := _bg.HTTPClient
	if _dff == nil {
		_dff = _fae()
	}
	_da, _cbd := _dff.Post(serverURL, "a\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0074\u0069\u006d\u0065\u0073t\u0061\u006d\u0070-\u0071u\u0065\u0072\u0079", _d.NewBuffer(_afa))
	if _cbd != nil {
		return nil, _cbd
	}
	defer _da.Body.Close()
	_cac, _cbd := _ab.ReadAll(_da.Body)
	if _cbd != nil {
		return nil, _cbd
	}
	if _da.StatusCode != _df.StatusOK {
		return nil, _ad.Errorf("\u0075\u006e\u0065x\u0070\u0065\u0063\u0074e\u0064\u0020\u0048\u0054\u0054\u0050\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0064", _da.StatusCode)
	}
	var _eef struct {
		Version _ca.RawValue
		Content _ca.RawValue
	}
	if _, _cbd = _ca.Unmarshal(_cac, &_eef); _cbd != nil {
		return nil, _cbd
	}
	return _eef.Content.FullBytes, nil
}

// MakeRequest makes a OCSP request to the specified server and returns
// the parsed and raw responses. If a server URL is not provided, it is
// extracted from the certificate.
func (_ccg *OCSPClient) MakeRequest(serverURL string, cert, issuer *_ge.Certificate) (*_ag.Response, []byte, error) {
	if _ccg.HTTPClient == nil {
		_ccg.HTTPClient = _fae()
	}
	if serverURL == "" {
		if len(cert.OCSPServer) == 0 {
			return nil, nil, _g.New("\u0063e\u0072\u0074i\u0066\u0069\u0063a\u0074\u0065\u0020\u0064\u006f\u0065\u0073 \u006e\u006f\u0074\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0079\u0020\u004f\u0043S\u0050\u0020\u0073\u0065\u0072\u0076\u0065\u0072\u0073")
		}
		serverURL = cert.OCSPServer[0]
	}
	_cb, _ga := _ag.CreateRequest(cert, issuer, &_ag.RequestOptions{Hash: _ccg.Hash})
	if _ga != nil {
		return nil, nil, _ga
	}
	_eea, _ga := _ccg.HTTPClient.Post(serverURL, "\u0061p\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u006fc\u0073\u0070\u002d\u0072\u0065\u0071\u0075\u0065\u0073\u0074", _d.NewReader(_cb))
	if _ga != nil {
		return nil, nil, _ga
	}
	defer _eea.Body.Close()
	_afg, _ga := _ab.ReadAll(_eea.Body)
	if _ga != nil {
		return nil, nil, _ga
	}
	if _gaf, _ := _ef.Decode(_afg); _gaf != nil {
		_afg = _gaf.Bytes
	}
	_ba, _ga := _ag.ParseResponseForCert(_afg, cert, issuer)
	if _ga != nil {
		return nil, nil, _ga
	}
	return _ba, _afg, nil
}

// NewCRLClient returns a new CRL client.
func NewCRLClient() *CRLClient { return &CRLClient{HTTPClient: _fae()} }

// CRLClient represents a CRL (Certificate revocation list) client.
// It is used to request revocation data from CRL servers.
type CRLClient struct {

	// HTTPClient is the HTTP client used to make CRL requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_df.Client
}

// NewCertClient returns a new certificate client.
func NewCertClient() *CertClient { return &CertClient{HTTPClient: _fae()} }

// MakeRequest makes a CRL request to the specified server and returns the
// response. If a server URL is not provided, it is extracted from the certificate.
func (_ec *CRLClient) MakeRequest(serverURL string, cert *_ge.Certificate) ([]byte, error) {
	if _ec.HTTPClient == nil {
		_ec.HTTPClient = _fae()
	}
	if serverURL == "" {
		if len(cert.CRLDistributionPoints) == 0 {
			return nil, _g.New("\u0063e\u0072\u0074i\u0066\u0069\u0063\u0061t\u0065\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070ec\u0069\u0066\u0079 \u0061\u006ey\u0020\u0043\u0052\u004c\u0020\u0073e\u0072\u0076e\u0072\u0073")
		}
		serverURL = cert.CRLDistributionPoints[0]
	}
	_gcd, _fg := _ec.HTTPClient.Get(serverURL)
	if _fg != nil {
		return nil, _fg
	}
	defer _gcd.Body.Close()
	_ff, _fg := _ab.ReadAll(_gcd.Body)
	if _fg != nil {
		return nil, _fg
	}
	if _dc, _ := _ef.Decode(_ff); _dc != nil {
		_ff = _dc.Bytes
	}
	return _ff, nil
}

// NewTimestampRequest returns a new timestamp request based
// on the specified options.
func NewTimestampRequest(body _f.Reader, opts *_b.RequestOptions) (*_b.Request, error) {
	if opts == nil {
		opts = &_b.RequestOptions{}
	}
	if opts.Hash == 0 {
		opts.Hash = _e.SHA256
	}
	if !opts.Hash.Available() {
		return nil, _ge.ErrUnsupportedAlgorithm
	}
	_ddc := opts.Hash.New()
	if _, _dfe := _f.Copy(_ddc, body); _dfe != nil {
		return nil, _dfe
	}
	return &_b.Request{HashAlgorithm: opts.Hash, HashedMessage: _ddc.Sum(nil), Certificates: opts.Certificates, TSAPolicyOID: opts.TSAPolicyOID, Nonce: opts.Nonce}, nil
}

// NewOCSPClient returns a new OCSP client.
func NewOCSPClient() *OCSPClient { return &OCSPClient{HTTPClient: _fae(), Hash: _e.SHA1} }

// Get retrieves the certificate at the specified URL.
func (_gc *CertClient) Get(url string) (*_ge.Certificate, error) {
	if _gc.HTTPClient == nil {
		_gc.HTTPClient = _fae()
	}
	_ce, _dd := _gc.HTTPClient.Get(url)
	if _dd != nil {
		return nil, _dd
	}
	defer _ce.Body.Close()
	_cg, _dd := _ab.ReadAll(_ce.Body)
	if _dd != nil {
		return nil, _dd
	}
	if _de, _ := _ef.Decode(_cg); _de != nil {
		_cg = _de.Bytes
	}
	_bd, _dd := _ge.ParseCertificate(_cg)
	if _dd != nil {
		return nil, _dd
	}
	return _bd, nil
}

// IsCA returns true if the provided certificate appears to be a CA certificate.
func (_af *CertClient) IsCA(cert *_ge.Certificate) bool {
	return cert.IsCA && _d.Equal(cert.RawIssuer, cert.RawSubject)
}

// GetIssuer retrieves the issuer of the provided certificate.
func (_fe *CertClient) GetIssuer(cert *_ge.Certificate) (*_ge.Certificate, error) {
	for _, _ee := range cert.IssuingCertificateURL {
		_aa, _fa := _fe.Get(_ee)
		if _fa != nil {
			_gea.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074 \u0064\u006f\u0077\u006e\u006c\u006f\u0061\u0064\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0066\u006f\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0025\u0076\u003a\u0020\u0025\u0076", cert.Subject.CommonName, _fa)
			continue
		}
		return _aa, nil
	}
	return nil, _ad.Errorf("\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063a\u0074e\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064")
}

// OCSPClient represents a OCSP (Online Certificate Status Protocol) client.
// It is used to request revocation data from OCSP servers.
type OCSPClient struct {

	// HTTPClient is the HTTP client used to make OCSP requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_df.Client

	// Hash is the hash function  used when constructing the OCSP
	// requests. If zero, SHA-1 will be used.
	Hash _e.Hash
}

func _fae() *_df.Client { return &_df.Client{Timeout: 5 * _c.Second} }

// NewTimestampClient returns a new timestamp client.
func NewTimestampClient() *TimestampClient { return &TimestampClient{HTTPClient: _fae()} }

// CertClient represents a X.509 certificate client. Its primary purpose
// is to download certificates.
type CertClient struct {

	// HTTPClient is the HTTP client used to make certificate requests.
	// By default, an HTTP client with a 5 second timeout per request is used.
	HTTPClient *_df.Client
}
