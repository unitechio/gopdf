package crypt

import (
	_a "crypto/aes"
	_fg "crypto/cipher"
	_ae "crypto/md5"
	_aa "crypto/rand"
	_d "crypto/rc4"
	_g "fmt"
	_e "io"

	_dg "unitechio/gopdf/gopdf/common"
	_dc "unitechio/gopdf/gopdf/core/security"
)

func init() { _gf("\u0041\u0045\u0053V\u0032", _fc) }

// NewFilterV2 creates a RC4-based filter with a specified key length (in bytes).
func NewFilterV2(length int) Filter {
	_abg, _ea := _gc(FilterDict{Length: length})
	if _ea != nil {
		_dg.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020R\u0043\u0034\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _ea)
		return filterV2{_dbc: length}
	}
	return _abg
}

var _ Filter = filterAESV2{}

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
	MakeKey(_ecff, _fde uint32, _eab []byte) ([]byte, error)

	// EncryptBytes encrypts a buffer using object encryption key, as returned by MakeKey.
	// Implementation may reuse a buffer and encrypt data in-place.
	EncryptBytes(_bed []byte, _dd []byte) ([]byte, error)

	// DecryptBytes decrypts a buffer using object encryption key, as returned by MakeKey.
	// Implementation may reuse a buffer and decrypt data in-place.
	DecryptBytes(_de []byte, _bdg []byte) ([]byte, error)
}

var _ Filter = filterV2{}

// PDFVersion implements Filter interface.
func (filterAESV3) PDFVersion() [2]int { return [2]int{2, 0} }

// KeyLength implements Filter interface.
func (_ee filterV2) KeyLength() int { return _ee._dbc }

func _gc(_bg FilterDict) (Filter, error) {
	if _bg.Length%8 != 0 {
		return nil, _g.Errorf("\u0063\u0072\u0079p\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020o\u0066\u0020\u0038\u0020\u0028\u0025\u0064\u0029", _bg.Length)
	}
	if _bg.Length < 5 || _bg.Length > 16 {
		if _bg.Length == 40 || _bg.Length == 64 || _bg.Length == 128 {
			_dg.Log.Debug("\u0053\u0054\u0041\u004e\u0044AR\u0044\u0020V\u0049\u004f\u004c\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0043\u0072\u0079\u0070\u0074\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072s\u0020\u0074\u006f \u0062\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u0074\u0073\u0020\u0072\u0061t\u0068\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0020-\u0020\u0061s\u0073u\u006d\u0069\u006e\u0067\u0020\u0062\u0069t\u0073\u0020\u0028\u0025\u0064\u0029", _bg.Length)
			_bg.Length /= 8
		} else {
			return nil, _g.Errorf("\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074h\u0020\u006e\u006f\u0074\u0020\u0069\u006e \u0072\u0061\u006e\u0067\u0065\u0020\u0034\u0030\u0020\u002d\u00201\u0032\u0038\u0020\u0062\u0069\u0074\u0020\u0028\u0025\u0064\u0029", _bg.Length)
		}
	}
	return filterV2{_dbc: _bg.Length}, nil
}

type filterAES struct{}

func _bgg(_aee, _ga uint32, _bga []byte, _bd bool) ([]byte, error) {
	_ecd := make([]byte, len(_bga)+5)
	copy(_ecd, _bga)
	for _cc := 0; _cc < 3; _cc++ {
		_ed := byte((_aee >> uint32(8*_cc)) & 0xff)
		_ecd[_cc+len(_bga)] = _ed
	}
	for _ad := 0; _ad < 2; _ad++ {
		_adf := byte((_ga >> uint32(8*_ad)) & 0xff)
		_ecd[_ad+len(_bga)+3] = _adf
	}
	if _bd {
		_ecd = append(_ecd, 0x73)
		_ecd = append(_ecd, 0x41)
		_ecd = append(_ecd, 0x6C)
		_ecd = append(_ecd, 0x54)
	}
	_adc := _ae.New()
	_adc.Write(_ecd)
	_cf := _adc.Sum(nil)
	if len(_bga)+5 < 16 {
		return _cf[0 : len(_bga)+5], nil
	}
	return _cf, nil
}
func init() { _gf("\u0041\u0045\u0053V\u0033", _ffb) }

// Name implements Filter interface.
func (filterV2) Name() string { return "\u0056\u0032" }

type filterAESV2 struct{ filterAES }

var _ Filter = filterAESV3{}

// HandlerVersion implements Filter interface.
func (_gcc filterV2) HandlerVersion() (V, R int) { V, R = 2, 3; return }

type filterAESV3 struct{ filterAES }

func (filterAES) EncryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_db, _be := _a.NewCipher(okey)
	if _be != nil {
		return nil, _be
	}
	_dg.Log.Trace("A\u0045\u0053\u0020\u0045nc\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	const _df = _a.BlockSize
	_ebd := _df - len(buf)%_df
	for _af := 0; _af < _ebd; _af++ {
		buf = append(buf, byte(_ebd))
	}
	_dg.Log.Trace("\u0050a\u0064d\u0065\u0064\u0020\u0074\u006f \u0025\u0064 \u0062\u0079\u0074\u0065\u0073", len(buf))
	_aac := make([]byte, _df+len(buf))
	_dgfc := _aac[:_df]
	if _, _c := _e.ReadFull(_aa.Reader, _dgfc); _c != nil {
		return nil, _c
	}
	_gd := _fg.NewCBCEncrypter(_db, _dgfc)
	_gd.CryptBlocks(_aac[_df:], buf)
	buf = _aac
	_dg.Log.Trace("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	return buf, nil
}

// NewIdentity creates an identity filter that bypasses all data without changes.
func NewIdentity() Filter { return filterIdentity{} }

// HandlerVersion implements Filter interface.
func (filterAESV3) HandlerVersion() (V, R int) { V, R = 5, 6; return }

// HandlerVersion implements Filter interface.
func (filterAESV2) HandlerVersion() (V, R int) { V, R = 4, 4; return }

func _fgeg(_afc string) (filterFunc, error) {
	_aab := _cce[_afc]
	if _aab == nil {
		return nil, _g.Errorf("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0072\u0079p\u0074 \u0066\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0071", _afc)
	}
	return _aab, nil
}

func (filterAES) DecryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_dbd, _fa := _a.NewCipher(okey)
	if _fa != nil {
		return nil, _fa
	}
	if len(buf) < 16 {
		_dg.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0041\u0045\u0053\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0062\u0075\u0066\u0020\u0025\u0073", buf)
		return buf, _g.Errorf("\u0041\u0045\u0053\u003a B\u0075\u0066\u0020\u006c\u0065\u006e\u0020\u003c\u0020\u0031\u0036\u0020\u0028\u0025d\u0029", len(buf))
	}
	_ebdc := buf[:16]
	buf = buf[16:]
	if len(buf)%16 != 0 {
		_dg.Log.Debug("\u0020\u0069\u0076\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(_ebdc), _ebdc)
		_dg.Log.Debug("\u0062\u0075\u0066\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
		return buf, _g.Errorf("\u0041\u0045\u0053\u0020\u0062\u0075\u0066\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069p\u006c\u0065\u0020\u006f\u0066 \u0031\u0036 \u0028\u0025\u0064\u0029", len(buf))
	}
	_ec := _fg.NewCBCDecrypter(_dbd, _ebdc)
	_dg.Log.Trace("A\u0045\u0053\u0020\u0044ec\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	_dg.Log.Trace("\u0063\u0068\u006f\u0070\u0020\u0041\u0045\u0053\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u0020\u0028\u0025\u0064\u0029\u003a \u0025\u0020\u0078", len(buf), buf)
	_ec.CryptBlocks(buf, buf)
	_dg.Log.Trace("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	if len(buf) == 0 {
		_dg.Log.Trace("\u0045\u006d\u0070\u0074\u0079\u0020b\u0075\u0066\u002c\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0065\u006d\u0070\u0074\u0079\u0020\u0073t\u0072\u0069\u006e\u0067")
		return buf, nil
	}
	_fd := int(buf[len(buf)-1])
	if _fd > len(buf) {
		_dg.Log.Debug("\u0049\u006c\u006c\u0065g\u0061\u006c\u0020\u0070\u0061\u0064\u0020\u006c\u0065\u006eg\u0074h\u0020\u0028\u0025\u0064\u0020\u003e\u0020%\u0064\u0029", _fd, len(buf))
		return buf, _g.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0070a\u0064\u0020l\u0065\u006e\u0067\u0074\u0068")
	}
	buf = buf[:len(buf)-_fd]
	return buf, nil
}

// NewFilterAESV2 creates an AES-based filter with a 128 bit key (AESV2).
func NewFilterAESV2() Filter {
	_dcb, _eb := _fc(FilterDict{})
	if _eb != nil {
		_dg.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _eb)
		return filterAESV2{}
	}
	return _dcb
}
func init() { _gf("\u0056\u0032", _gc) }
func _gf(_gae string, _fda filterFunc) {
	if _, _gcb := _cce[_gae]; _gcb {
		panic("\u0061l\u0072e\u0061\u0064\u0079\u0020\u0072e\u0067\u0069s\u0074\u0065\u0072\u0065\u0064")
	}
	_cce[_gae] = _fda
}

type filterIdentity struct{}

// MakeKey implements Filter interface.
func (filterAESV3) MakeKey(_, _ uint32, ekey []byte) ([]byte, error) { return ekey, nil }

// KeyLength implements Filter interface.
func (filterAESV3) KeyLength() int { return 256 / 8 }

// KeyLength implements Filter interface.
func (filterAESV2) KeyLength() int { return 128 / 8 }

// NewFilter creates CryptFilter from a corresponding dictionary.
func NewFilter(d FilterDict) (Filter, error) {
	_dee, _gcca := _fgeg(d.CFM)
	if _gcca != nil {
		return nil, _gcca
	}
	_ag, _gcca := _dee(d)
	if _gcca != nil {
		return nil, _gcca
	}
	return _ag, nil
}

type filterV2 struct{ _dbc int }

var _cce = make(map[string]filterFunc)

func (filterIdentity) KeyLength() int { return 0 }

// Name implements Filter interface.
func (filterAESV2) Name() string { return "\u0041\u0045\u0053V\u0032" }

type filterFunc func(_aef FilterDict) (Filter, error)

// PDFVersion implements Filter interface.
func (_aae filterV2) PDFVersion() [2]int { return [2]int{} }

func _fc(_ab FilterDict) (Filter, error) {
	if _ab.Length == 128 {
		_dg.Log.Debug("\u0041\u0045S\u0056\u0032\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029", _ab.Length)
		_ab.Length /= 8
	}
	if _ab.Length != 0 && _ab.Length != 16 {
		return nil, _g.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0032\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029", _ab.Length)
	}
	return filterAESV2{}, nil
}

// EncryptBytes implements Filter interface.
func (filterV2) EncryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_gb, _dbcb := _d.NewCipher(okey)
	if _dbcb != nil {
		return nil, _dbcb
	}
	_dg.Log.Trace("\u0052\u00434\u0020\u0045\u006ec\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078", buf)
	_gb.XORKeyStream(buf, buf)
	_dg.Log.Trace("\u0074o\u003a\u0020\u0025\u0020\u0078", buf)
	return buf, nil
}
func (filterIdentity) MakeKey(objNum, genNum uint32, fkey []byte) ([]byte, error) { return fkey, nil }
func (filterIdentity) EncryptBytes(p []byte, okey []byte) ([]byte, error)         { return p, nil }

// FilterDict represents information from a CryptFilter dictionary.
type FilterDict struct {
	CFM       string
	AuthEvent _dc.AuthEvent
	Length    int
}

// NewFilterAESV3 creates an AES-based filter with a 256 bit key (AESV3).
func NewFilterAESV3() Filter {
	_fcc, _ff := _ffb(FilterDict{})
	if _ff != nil {
		_dg.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0033\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _ff)
		return filterAESV3{}
	}
	return _fcc
}
func (filterIdentity) PDFVersion() [2]int { return [2]int{} }

// DecryptBytes implements Filter interface.
func (filterV2) DecryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_ba, _dbda := _d.NewCipher(okey)
	if _dbda != nil {
		return nil, _dbda
	}
	_dg.Log.Trace("\u0052\u00434\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078", buf)
	_ba.XORKeyStream(buf, buf)
	_dg.Log.Trace("\u0074o\u003a\u0020\u0025\u0020\u0078", buf)
	return buf, nil
}

// MakeKey implements Filter interface.
func (filterAESV2) MakeKey(objNum, genNum uint32, ekey []byte) ([]byte, error) {
	return _bgg(objNum, genNum, ekey, true)
}
func (filterIdentity) HandlerVersion() (V, R int) { return }
func _ffb(_b FilterDict) (Filter, error) {
	if _b.Length == 256 {
		_dg.Log.Debug("\u0041\u0045S\u0056\u0033\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029", _b.Length)
		_b.Length /= 8
	}
	if _b.Length != 0 && _b.Length != 32 {
		return nil, _g.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0033\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029", _b.Length)
	}
	return filterAESV3{}, nil
}
func (filterIdentity) Name() string { return "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" }

// PDFVersion implements Filter interface.
func (filterAESV2) PDFVersion() [2]int                                    { return [2]int{1, 5} }
func (filterIdentity) DecryptBytes(p []byte, okey []byte) ([]byte, error) { return p, nil }

// Name implements Filter interface.
func (filterAESV3) Name() string { return "\u0041\u0045\u0053V\u0033" }

// MakeKey implements Filter interface.
func (_bef filterV2) MakeKey(objNum, genNum uint32, ekey []byte) ([]byte, error) {
	return _bgg(objNum, genNum, ekey, false)
}
