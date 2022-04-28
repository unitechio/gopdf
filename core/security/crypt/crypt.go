package crypt

import (
	_db "crypto/aes"
	_e "crypto/cipher"
	_b "crypto/md5"
	_ga "crypto/rand"
	_g "crypto/rc4"
	_d "fmt"
	_f "io"

	_c "bitbucket.org/shenghui0779/gopdf/common"
	_fb "bitbucket.org/shenghui0779/gopdf/core/security"
)

func init() { _dbe("\u0041\u0045\u0053V\u0032", _gd) }

// KeyLength implements Filter interface.
func (filterAESV3) KeyLength() int { return 256 / 8 }

type filterV2 struct{ _egc int }

func (filterIdentity) Name() string { return "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" }

// PDFVersion implements Filter interface.
func (filterAESV2) PDFVersion() [2]int { return [2]int{1, 5} }
func _dbe(_ba string, _fgg filterFunc) {
	if _, _bfeb := _bfd[_ba]; _bfeb {
		panic("\u0061l\u0072e\u0061\u0064\u0079\u0020\u0072e\u0067\u0069s\u0074\u0065\u0072\u0065\u0064")
	}
	_bfd[_ba] = _fgg
}

// PDFVersion implements Filter interface.
func (filterAESV3) PDFVersion() [2]int { return [2]int{2, 0} }
func _acf(_cea, _gab uint32, _eg []byte, _bfg bool) ([]byte, error) {
	_cdd := make([]byte, len(_eg)+5)
	for _dd := 0; _dd < len(_eg); _dd++ {
		_cdd[_dd] = _eg[_dd]
	}
	for _ef := 0; _ef < 3; _ef++ {
		_ag := byte((_cea >> uint32(8*_ef)) & 0xff)
		_cdd[_ef+len(_eg)] = _ag
	}
	for _dbg := 0; _dbg < 2; _dbg++ {
		_cb := byte((_gab >> uint32(8*_dbg)) & 0xff)
		_cdd[_dbg+len(_eg)+3] = _cb
	}
	if _bfg {
		_cdd = append(_cdd, 0x73)
		_cdd = append(_cdd, 0x41)
		_cdd = append(_cdd, 0x6C)
		_cdd = append(_cdd, 0x54)
	}
	_ccb := _b.New()
	_ccb.Write(_cdd)
	_aa := _ccb.Sum(nil)
	if len(_eg)+5 < 16 {
		return _aa[0 : len(_eg)+5], nil
	}
	return _aa, nil
}

// NewFilterAESV2 creates an AES-based filter with a 128 bit key (AESV2).
func NewFilterAESV2() Filter {
	_fg, _bf := _gd(FilterDict{})
	if _bf != nil {
		_c.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _bf)
		return filterAESV2{}
	}
	return _fg
}

var _ Filter = filterV2{}

func init() { _dbe("\u0041\u0045\u0053V\u0033", _cd) }

// DecryptBytes implements Filter interface.
func (filterV2) DecryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_ecb, _be := _g.NewCipher(okey)
	if _be != nil {
		return nil, _be
	}
	_c.Log.Trace("\u0052\u00434\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078", buf)
	_ecb.XORKeyStream(buf, buf)
	_c.Log.Trace("\u0074o\u003a\u0020\u0025\u0020\u0078", buf)
	return buf, nil
}

// HandlerVersion implements Filter interface.
func (filterAESV3) HandlerVersion() (V, R int) { V, R = 5, 6; return }

type filterFunc func(_fgf FilterDict) (Filter, error)

func (filterIdentity) HandlerVersion() (V, R int) { return }

// KeyLength implements Filter interface.
func (_cg filterV2) KeyLength() int { return _cg._egc }

// NewFilterAESV3 creates an AES-based filter with a 256 bit key (AESV3).
func NewFilterAESV3() Filter {
	_ee, _gf := _cd(FilterDict{})
	if _gf != nil {
		_c.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0033\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _gf)
		return filterAESV3{}
	}
	return _ee
}

// NewFilterV2 creates a RC4-based filter with a specified key length (in bytes).
func NewFilterV2(length int) Filter {
	_edc, _ecg := _gbe(FilterDict{Length: length})
	if _ecg != nil {
		_c.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020R\u0043\u0034\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _ecg)
		return filterV2{_egc: length}
	}
	return _edc
}

// NewIdentity creates an identity filter that bypasses all data without changes.
func NewIdentity() Filter { return filterIdentity{} }

// MakeKey implements Filter interface.
func (filterAESV3) MakeKey(_, _ uint32, ekey []byte) ([]byte, error) { return ekey, nil }

// KeyLength implements Filter interface.
func (filterAESV2) KeyLength() int { return 128 / 8 }

// FilterDict represents information from a CryptFilter dictionary.
type FilterDict struct {
	CFM       string
	AuthEvent _fb.AuthEvent
	Length    int
}

// PDFVersion implements Filter interface.
func (_bd filterV2) PDFVersion() [2]int { return [2]int{} }
func init()                             { _dbe("\u0056\u0032", _gbe) }
func _fbd(_cgf string) (filterFunc, error) {
	_gdg := _bfd[_cgf]
	if _gdg == nil {
		return nil, _d.Errorf("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0072\u0079p\u0074 \u0066\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0071", _cgf)
	}
	return _gdg, nil
}
func (filterIdentity) MakeKey(objNum, genNum uint32, fkey []byte) ([]byte, error) { return fkey, nil }

type filterAESV2 struct{ filterAES }
type filterAES struct{}

func (filterAES) EncryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_bg, _af := _db.NewCipher(okey)
	if _af != nil {
		return nil, _af
	}
	_c.Log.Trace("A\u0045\u0053\u0020\u0045nc\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	const _ce = _db.BlockSize
	_afb := _ce - len(buf)%_ce
	for _gca := 0; _gca < _afb; _gca++ {
		buf = append(buf, byte(_afb))
	}
	_c.Log.Trace("\u0050a\u0064d\u0065\u0064\u0020\u0074\u006f \u0025\u0064 \u0062\u0079\u0074\u0065\u0073", len(buf))
	_cc := make([]byte, _ce+len(buf))
	_gb := _cc[:_ce]
	if _, _fd := _f.ReadFull(_ga.Reader, _gb); _fd != nil {
		return nil, _fd
	}
	_bga := _e.NewCBCEncrypter(_bg, _gb)
	_bga.CryptBlocks(_cc[_ce:], buf)
	buf = _cc
	_c.Log.Trace("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	return buf, nil
}

// Name implements Filter interface.
func (filterAESV3) Name() string { return "\u0041\u0045\u0053V\u0033" }

// HandlerVersion implements Filter interface.
func (_da filterV2) HandlerVersion() (V, R int) { V, R = 2, 3; return }

// Name implements Filter interface.
func (filterAESV2) Name() string { return "\u0041\u0045\u0053V\u0032" }

var _ Filter = filterAESV3{}

// EncryptBytes implements Filter interface.
func (filterV2) EncryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_cgc, _ad := _g.NewCipher(okey)
	if _ad != nil {
		return nil, _ad
	}
	_c.Log.Trace("\u0052\u00434\u0020\u0045\u006ec\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078", buf)
	_cgc.XORKeyStream(buf, buf)
	_c.Log.Trace("\u0074o\u003a\u0020\u0025\u0020\u0078", buf)
	return buf, nil
}
func (filterIdentity) DecryptBytes(p []byte, okey []byte) ([]byte, error) { return p, nil }

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
	MakeKey(_bfe, _gcf uint32, _bdb []byte) ([]byte, error)

	// EncryptBytes encrypts a buffer using object encryption key, as returned by MakeKey.
	// Implementation may reuse a buffer and encrypt data in-place.
	EncryptBytes(_bc []byte, _cdc []byte) ([]byte, error)

	// DecryptBytes decrypts a buffer using object encryption key, as returned by MakeKey.
	// Implementation may reuse a buffer and decrypt data in-place.
	DecryptBytes(_bfda []byte, _cf []byte) ([]byte, error)
}

func (filterIdentity) KeyLength() int { return 0 }

// NewFilter creates CryptFilter from a corresponding dictionary.
func NewFilter(d FilterDict) (Filter, error) {
	_beg, _bba := _fbd(d.CFM)
	if _bba != nil {
		return nil, _bba
	}
	_ca, _bba := _beg(d)
	if _bba != nil {
		return nil, _bba
	}
	return _ca, nil
}
func (filterAES) DecryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_ac, _df := _db.NewCipher(okey)
	if _df != nil {
		return nil, _df
	}
	if len(buf) < 16 {
		_c.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0041\u0045\u0053\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0062\u0075\u0066\u0020\u0025\u0073", buf)
		return buf, _d.Errorf("\u0041\u0045\u0053\u003a B\u0075\u0066\u0020\u006c\u0065\u006e\u0020\u003c\u0020\u0031\u0036\u0020\u0028\u0025d\u0029", len(buf))
	}
	_fdb := buf[:16]
	buf = buf[16:]
	if len(buf)%16 != 0 {
		_c.Log.Debug("\u0020\u0069\u0076\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(_fdb), _fdb)
		_c.Log.Debug("\u0062\u0075\u0066\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
		return buf, _d.Errorf("\u0041\u0045\u0053\u0020\u0062\u0075\u0066\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069p\u006c\u0065\u0020\u006f\u0066 \u0031\u0036 \u0028\u0025\u0064\u0029", len(buf))
	}
	_bb := _e.NewCBCDecrypter(_ac, _fdb)
	_c.Log.Trace("A\u0045\u0053\u0020\u0044ec\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	_c.Log.Trace("\u0063\u0068\u006f\u0070\u0020\u0041\u0045\u0053\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u0020\u0028\u0025\u0064\u0029\u003a \u0025\u0020\u0078", len(buf), buf)
	_bb.CryptBlocks(buf, buf)
	_c.Log.Trace("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	if len(buf) == 0 {
		_c.Log.Trace("\u0045\u006d\u0070\u0074\u0079\u0020b\u0075\u0066\u002c\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0065\u006d\u0070\u0074\u0079\u0020\u0073t\u0072\u0069\u006e\u0067")
		return buf, nil
	}
	_ec := int(buf[len(buf)-1])
	if _ec > len(buf) {
		_c.Log.Debug("\u0049\u006c\u006c\u0065g\u0061\u006c\u0020\u0070\u0061\u0064\u0020\u006c\u0065\u006eg\u0074h\u0020\u0028\u0025\u0064\u0020\u003e\u0020%\u0064\u0029", _ec, len(buf))
		return buf, _d.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0070a\u0064\u0020l\u0065\u006e\u0067\u0074\u0068")
	}
	buf = buf[:len(buf)-_ec]
	return buf, nil
}
func _gd(_ed FilterDict) (Filter, error) {
	if _ed.Length == 128 {
		_c.Log.Debug("\u0041\u0045S\u0056\u0032\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029", _ed.Length)
		_ed.Length /= 8
	}
	if _ed.Length != 0 && _ed.Length != 16 {
		return nil, _d.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0032\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029", _ed.Length)
	}
	return filterAESV2{}, nil
}
func (filterIdentity) PDFVersion() [2]int { return [2]int{} }

// HandlerVersion implements Filter interface.
func (filterAESV2) HandlerVersion() (V, R int) { V, R = 4, 4; return }

type filterAESV3 struct{ filterAES }

func _cd(_eb FilterDict) (Filter, error) {
	if _eb.Length == 256 {
		_c.Log.Debug("\u0041\u0045S\u0056\u0033\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029", _eb.Length)
		_eb.Length /= 8
	}
	if _eb.Length != 0 && _eb.Length != 32 {
		return nil, _d.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0033\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029", _eb.Length)
	}
	return filterAESV3{}, nil
}
func _gbe(_eef FilterDict) (Filter, error) {
	if _eef.Length%8 != 0 {
		return nil, _d.Errorf("\u0063\u0072\u0079p\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020o\u0066\u0020\u0038\u0020\u0028\u0025\u0064\u0029", _eef.Length)
	}
	if _eef.Length < 5 || _eef.Length > 16 {
		if _eef.Length == 40 || _eef.Length == 64 || _eef.Length == 128 {
			_c.Log.Debug("\u0053\u0054\u0041\u004e\u0044AR\u0044\u0020V\u0049\u004f\u004c\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0043\u0072\u0079\u0070\u0074\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072s\u0020\u0074\u006f \u0062\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u0074\u0073\u0020\u0072\u0061t\u0068\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0020-\u0020\u0061s\u0073u\u006d\u0069\u006e\u0067\u0020\u0062\u0069t\u0073\u0020\u0028\u0025\u0064\u0029", _eef.Length)
			_eef.Length /= 8
		} else {
			return nil, _d.Errorf("\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074h\u0020\u006e\u006f\u0074\u0020\u0069\u006e \u0072\u0061\u006e\u0067\u0065\u0020\u0034\u0030\u0020\u002d\u00201\u0032\u0038\u0020\u0062\u0069\u0074\u0020\u0028\u0025\u0064\u0029", _eef.Length)
		}
	}
	return filterV2{_egc: _eef.Length}, nil
}

type filterIdentity struct{}

var _ Filter = filterAESV2{}

// MakeKey implements Filter interface.
func (filterAESV2) MakeKey(objNum, genNum uint32, ekey []byte) ([]byte, error) {
	return _acf(objNum, genNum, ekey, true)
}

var (
	_bfd = make(map[string]filterFunc)
)

// MakeKey implements Filter interface.
func (_fga filterV2) MakeKey(objNum, genNum uint32, ekey []byte) ([]byte, error) {
	return _acf(objNum, genNum, ekey, false)
}
func (filterIdentity) EncryptBytes(p []byte, okey []byte) ([]byte, error) { return p, nil }

// Name implements Filter interface.
func (filterV2) Name() string { return "\u0056\u0032" }
