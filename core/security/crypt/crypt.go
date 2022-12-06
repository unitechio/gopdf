package crypt

import (
	_ce "crypto/aes"
	_de "crypto/cipher"
	_bc "crypto/md5"
	_f "crypto/rand"
	_g "crypto/rc4"
	_b "fmt"
	_c "io"

	_fd "bitbucket.org/shenghui0779/gopdf/common"
	_cg "bitbucket.org/shenghui0779/gopdf/core/security"
)

func init()                                                               { _cgc("\u0041\u0045\u0053V\u0032", _ef) }
func (filterIdentity) EncryptBytes(p []byte, okey []byte) ([]byte, error) { return p, nil }

var (
	_egg = make(map[string]filterFunc)
)

func _bf(_cec FilterDict) (Filter, error) {
	if _cec.Length == 256 {
		_fd.Log.Debug("\u0041\u0045S\u0056\u0033\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029", _cec.Length)
		_cec.Length /= 8
	}
	if _cec.Length != 0 && _cec.Length != 32 {
		return nil, _b.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0033\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029", _cec.Length)
	}
	return filterAESV3{}, nil
}
func _da(_daa string) (filterFunc, error) {
	_bdf := _egg[_daa]
	if _bdf == nil {
		return nil, _b.Errorf("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0072\u0079p\u0074 \u0066\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0071", _daa)
	}
	return _bdf, nil
}

// NewIdentity creates an identity filter that bypasses all data without changes.
func NewIdentity() Filter { return filterIdentity{} }

// DecryptBytes implements Filter interface.
func (filterV2) DecryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_bff, _efb := _g.NewCipher(okey)
	if _efb != nil {
		return nil, _efb
	}
	_fd.Log.Trace("\u0052\u00434\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078", buf)
	_bff.XORKeyStream(buf, buf)
	_fd.Log.Trace("\u0074o\u003a\u0020\u0025\u0020\u0078", buf)
	return buf, nil
}
func (filterAES) DecryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_cac, _ba := _ce.NewCipher(okey)
	if _ba != nil {
		return nil, _ba
	}
	if len(buf) < 16 {
		_fd.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0041\u0045\u0053\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0062\u0075\u0066\u0020\u0025\u0073", buf)
		return buf, _b.Errorf("\u0041\u0045\u0053\u003a B\u0075\u0066\u0020\u006c\u0065\u006e\u0020\u003c\u0020\u0031\u0036\u0020\u0028\u0025d\u0029", len(buf))
	}
	_ed := buf[:16]
	buf = buf[16:]
	if len(buf)%16 != 0 {
		_fd.Log.Debug("\u0020\u0069\u0076\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(_ed), _ed)
		_fd.Log.Debug("\u0062\u0075\u0066\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
		return buf, _b.Errorf("\u0041\u0045\u0053\u0020\u0062\u0075\u0066\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069p\u006c\u0065\u0020\u006f\u0066 \u0031\u0036 \u0028\u0025\u0064\u0029", len(buf))
	}
	_cea := _de.NewCBCDecrypter(_cac, _ed)
	_fd.Log.Trace("A\u0045\u0053\u0020\u0044ec\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	_fd.Log.Trace("\u0063\u0068\u006f\u0070\u0020\u0041\u0045\u0053\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u0020\u0028\u0025\u0064\u0029\u003a \u0025\u0020\u0078", len(buf), buf)
	_cea.CryptBlocks(buf, buf)
	_fd.Log.Trace("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	if len(buf) == 0 {
		_fd.Log.Trace("\u0045\u006d\u0070\u0074\u0079\u0020b\u0075\u0066\u002c\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0065\u006d\u0070\u0074\u0079\u0020\u0073t\u0072\u0069\u006e\u0067")
		return buf, nil
	}
	_dg := int(buf[len(buf)-1])
	if _dg > len(buf) {
		_fd.Log.Debug("\u0049\u006c\u006c\u0065g\u0061\u006c\u0020\u0070\u0061\u0064\u0020\u006c\u0065\u006eg\u0074h\u0020\u0028\u0025\u0064\u0020\u003e\u0020%\u0064\u0029", _dg, len(buf))
		return buf, _b.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0070a\u0064\u0020l\u0065\u006e\u0067\u0074\u0068")
	}
	buf = buf[:len(buf)-_dg]
	return buf, nil
}

type filterAES struct{}

// HandlerVersion implements Filter interface.
func (_gaf filterV2) HandlerVersion() (V, R int) { V, R = 2, 3; return }
func init()                                      { _cgc("\u0041\u0045\u0053V\u0033", _bf) }

// KeyLength implements Filter interface.
func (filterAESV2) KeyLength() int { return 128 / 8 }

// Name implements Filter interface.
func (filterAESV3) Name() string { return "\u0041\u0045\u0053V\u0033" }

type filterAESV3 struct{ filterAES }
type filterAESV2 struct{ filterAES }

// HandlerVersion implements Filter interface.
func (filterAESV2) HandlerVersion() (V, R int) { V, R = 4, 4; return }
func _cgc(_gg string, _gbg filterFunc) {
	if _, _ceec := _egg[_gg]; _ceec {
		panic("\u0061l\u0072e\u0061\u0064\u0079\u0020\u0072e\u0067\u0069s\u0074\u0065\u0072\u0065\u0064")
	}
	_egg[_gg] = _gbg
}

type filterV2 struct{ _ga int }

// NewFilterV2 creates a RC4-based filter with a specified key length (in bytes).
func NewFilterV2(length int) Filter {
	_cd, _cgf := _ccg(FilterDict{Length: length})
	if _cgf != nil {
		_fd.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020R\u0043\u0034\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _cgf)
		return filterV2{_ga: length}
	}
	return _cd
}

// PDFVersion implements Filter interface.
func (filterAESV3) PDFVersion() [2]int { return [2]int{2, 0} }
func (filterAES) EncryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_dbg, _bcg := _ce.NewCipher(okey)
	if _bcg != nil {
		return nil, _bcg
	}
	_fd.Log.Trace("A\u0045\u0053\u0020\u0045nc\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	const _ca = _ce.BlockSize
	_ac := _ca - len(buf)%_ca
	for _cca := 0; _cca < _ac; _cca++ {
		buf = append(buf, byte(_ac))
	}
	_fd.Log.Trace("\u0050a\u0064d\u0065\u0064\u0020\u0074\u006f \u0025\u0064 \u0062\u0079\u0074\u0065\u0073", len(buf))
	_af := make([]byte, _ca+len(buf))
	_ad := _af[:_ca]
	if _, _gb := _c.ReadFull(_f.Reader, _ad); _gb != nil {
		return nil, _gb
	}
	_ae := _de.NewCBCEncrypter(_dbg, _ad)
	_ae.CryptBlocks(_af[_ca:], buf)
	buf = _af
	_fd.Log.Trace("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	return buf, nil
}

// Name implements Filter interface.
func (filterV2) Name() string { return "\u0056\u0032" }

// Filter is a common interface for crypt filter methods.
type Filter interface {

	// Name returns a name of the filter that should be used in CFM field of Encrypt dictionary.
	Name() string

	// KeyLength returns a length of the encryption key in bytes.
	KeyLength() int

	// PDFVersion reports the minimal version of PDF document that introduced this filter.
	PDFVersion() [2]int

	// HandlerVersion reports V and R parameters that should be used for this filter.
	HandlerVersion() (V, R int)

	// MakeKey generates a object encryption key based on file encryption key and object numbers.
	// Used only for legacy filters - AESV3 doesn't change the key for each object.
	MakeKey(_eb, _cf uint32, _gc []byte) ([]byte, error)

	// EncryptBytes encrypts a buffer using object encryption key, as returned by MakeKey.
	// Implementation may reuse a buffer and encrypt data in-place.
	EncryptBytes(_gfd []byte, _gaa []byte) ([]byte, error)

	// DecryptBytes decrypts a buffer using object encryption key, as returned by MakeKey.
	// Implementation may reuse a buffer and decrypt data in-place.
	DecryptBytes(_eeb []byte, _aea []byte) ([]byte, error)
}

func init() { _cgc("\u0056\u0032", _ccg) }

// NewFilterAESV2 creates an AES-based filter with a 128 bit key (AESV2).
func NewFilterAESV2() Filter {
	_cc, _e := _ef(FilterDict{})
	if _e != nil {
		_fd.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _e)
		return filterAESV2{}
	}
	return _cc
}

// Name implements Filter interface.
func (filterAESV2) Name() string { return "\u0041\u0045\u0053V\u0032" }

// MakeKey implements Filter interface.
func (filterAESV3) MakeKey(_, _ uint32, ekey []byte) ([]byte, error) { return ekey, nil }

var _ Filter = filterV2{}

// NewFilter creates CryptFilter from a corresponding dictionary.
func NewFilter(d FilterDict) (Filter, error) {
	_egb, _fdg := _da(d.CFM)
	if _fdg != nil {
		return nil, _fdg
	}
	_bac, _fdg := _egb(d)
	if _fdg != nil {
		return nil, _fdg
	}
	return _bac, nil
}

// PDFVersion implements Filter interface.
func (filterAESV2) PDFVersion() [2]int { return [2]int{1, 5} }
func (filterIdentity) Name() string    { return "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" }

// FilterDict represents information from a CryptFilter dictionary.
type FilterDict struct {
	CFM       string
	AuthEvent _cg.AuthEvent
	Length    int
}

// HandlerVersion implements Filter interface.
func (filterAESV3) HandlerVersion() (V, R int) { V, R = 5, 6; return }
func _def(_bdd, _aed uint32, _dd []byte, _edg bool) ([]byte, error) {
	_ee := make([]byte, len(_dd)+5)
	for _afg := 0; _afg < len(_dd); _afg++ {
		_ee[_afg] = _dd[_afg]
	}
	for _ag := 0; _ag < 3; _ag++ {
		_eg := byte((_bdd >> uint32(8*_ag)) & 0xff)
		_ee[_ag+len(_dd)] = _eg
	}
	for _cee := 0; _cee < 2; _cee++ {
		_ceg := byte((_aed >> uint32(8*_cee)) & 0xff)
		_ee[_cee+len(_dd)+3] = _ceg
	}
	if _edg {
		_ee = append(_ee, 0x73)
		_ee = append(_ee, 0x41)
		_ee = append(_ee, 0x6C)
		_ee = append(_ee, 0x54)
	}
	_dee := _bc.New()
	_dee.Write(_ee)
	_df := _dee.Sum(nil)
	if len(_dd)+5 < 16 {
		return _df[0 : len(_dd)+5], nil
	}
	return _df, nil
}
func _ccg(_aee FilterDict) (Filter, error) {
	if _aee.Length%8 != 0 {
		return nil, _b.Errorf("\u0063\u0072\u0079p\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020o\u0066\u0020\u0038\u0020\u0028\u0025\u0064\u0029", _aee.Length)
	}
	if _aee.Length < 5 || _aee.Length > 16 {
		if _aee.Length == 40 || _aee.Length == 64 || _aee.Length == 128 {
			_fd.Log.Debug("\u0053\u0054\u0041\u004e\u0044AR\u0044\u0020V\u0049\u004f\u004c\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0043\u0072\u0079\u0070\u0074\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072s\u0020\u0074\u006f \u0062\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u0074\u0073\u0020\u0072\u0061t\u0068\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0020-\u0020\u0061s\u0073u\u006d\u0069\u006e\u0067\u0020\u0062\u0069t\u0073\u0020\u0028\u0025\u0064\u0029", _aee.Length)
			_aee.Length /= 8
		} else {
			return nil, _b.Errorf("\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074h\u0020\u006e\u006f\u0074\u0020\u0069\u006e \u0072\u0061\u006e\u0067\u0065\u0020\u0034\u0030\u0020\u002d\u00201\u0032\u0038\u0020\u0062\u0069\u0074\u0020\u0028\u0025\u0064\u0029", _aee.Length)
		}
	}
	return filterV2{_ga: _aee.Length}, nil
}
func (filterIdentity) MakeKey(objNum, genNum uint32, fkey []byte) ([]byte, error) { return fkey, nil }

// NewFilterAESV3 creates an AES-based filter with a 256 bit key (AESV3).
func NewFilterAESV3() Filter {
	_dc, _ec := _bf(FilterDict{})
	if _ec != nil {
		_fd.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0033\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _ec)
		return filterAESV3{}
	}
	return _dc
}

// KeyLength implements Filter interface.
func (filterAESV3) KeyLength() int { return 256 / 8 }
func _ef(_a FilterDict) (Filter, error) {
	if _a.Length == 128 {
		_fd.Log.Debug("\u0041\u0045S\u0056\u0032\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029", _a.Length)
		_a.Length /= 8
	}
	if _a.Length != 0 && _a.Length != 16 {
		return nil, _b.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0032\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029", _a.Length)
	}
	return filterAESV2{}, nil
}

// KeyLength implements Filter interface.
func (_ge filterV2) KeyLength() int                                       { return _ge._ga }
func (filterIdentity) DecryptBytes(p []byte, okey []byte) ([]byte, error) { return p, nil }
func (filterIdentity) HandlerVersion() (V, R int)                         { return }

// PDFVersion implements Filter interface.
func (_fg filterV2) PDFVersion() [2]int { return [2]int{} }

// EncryptBytes implements Filter interface.
func (filterV2) EncryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_cegd, _aff := _g.NewCipher(okey)
	if _aff != nil {
		return nil, _aff
	}
	_fd.Log.Trace("\u0052\u00434\u0020\u0045\u006ec\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078", buf)
	_cegd.XORKeyStream(buf, buf)
	_fd.Log.Trace("\u0074o\u003a\u0020\u0025\u0020\u0078", buf)
	return buf, nil
}
func (filterIdentity) PDFVersion() [2]int { return [2]int{} }

type filterFunc func(_ceaa FilterDict) (Filter, error)

var _ Filter = filterAESV3{}

func (filterIdentity) KeyLength() int { return 0 }

// MakeKey implements Filter interface.
func (filterAESV2) MakeKey(objNum, genNum uint32, ekey []byte) ([]byte, error) {
	return _def(objNum, genNum, ekey, true)
}

type filterIdentity struct{}

var _ Filter = filterAESV2{}

// MakeKey implements Filter interface.
func (_gf filterV2) MakeKey(objNum, genNum uint32, ekey []byte) ([]byte, error) {
	return _def(objNum, genNum, ekey, false)
}
