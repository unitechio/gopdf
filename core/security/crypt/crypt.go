package crypt

import (
	_e "crypto/aes"
	_d "crypto/cipher"
	_f "crypto/md5"
	_gg "crypto/rand"
	_a "crypto/rc4"
	_b "fmt"
	_g "io"

	_dg "bitbucket.org/shenghui0779/gopdf/common"
	_ce "bitbucket.org/shenghui0779/gopdf/core/security"
)

func init()                         { _ede("\u0041\u0045\u0053V\u0032", _cf) }
func (filterIdentity) Name() string { return "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079" }

// Name implements Filter interface.
func (filterAESV3) Name() string { return "\u0041\u0045\u0053V\u0033" }
func _aga(_ebc string) (filterFunc, error) {
	_cef := _dac[_ebc]
	if _cef == nil {
		return nil, _b.Errorf("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0072\u0079p\u0074 \u0066\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0071", _ebc)
	}
	return _cef, nil
}

var _ Filter = filterAESV3{}

// Name implements Filter interface.
func (filterAESV2) Name() string      { return "\u0041\u0045\u0053V\u0032" }
func (filterIdentity) KeyLength() int { return 0 }

type filterAES struct{}

// KeyLength implements Filter interface.
func (_aed filterV2) KeyLength() int { return _aed._ae }

type filterAESV3 struct{ filterAES }

func init()                                       { _ede("\u0041\u0045\u0053V\u0033", _adg) }
func (filterIdentity) HandlerVersion() (V, R int) { return }

// PDFVersion implements Filter interface.
func (_dfe filterV2) PDFVersion() [2]int { return [2]int{} }

type filterFunc func(_afg FilterDict) (Filter, error)
type filterAESV2 struct{ filterAES }

func (filterIdentity) PDFVersion() [2]int { return [2]int{} }

// Name implements Filter interface.
func (filterV2) Name() string { return "\u0056\u0032" }
func _bc(_cge, _ec uint32, _ea []byte, _dae bool) ([]byte, error) {
	_cc := make([]byte, len(_ea)+5)
	copy(_cc, _ea)
	for _bcd := 0; _bcd < 3; _bcd++ {
		_edd := byte((_cge >> uint32(8*_bcd)) & 0xff)
		_cc[_bcd+len(_ea)] = _edd
	}
	for _cfb := 0; _cfb < 2; _cfb++ {
		_fbd := byte((_ec >> uint32(8*_cfb)) & 0xff)
		_cc[_cfb+len(_ea)+3] = _fbd
	}
	if _dae {
		_cc = append(_cc, 0x73)
		_cc = append(_cc, 0x41)
		_cc = append(_cc, 0x6C)
		_cc = append(_cc, 0x54)
	}
	_ef := _f.New()
	_ef.Write(_cc)
	_gb := _ef.Sum(nil)
	if len(_ea)+5 < 16 {
		return _gb[0 : len(_ea)+5], nil
	}
	return _gb, nil
}

var _ Filter = filterAESV2{}

// NewFilter creates CryptFilter from a corresponding dictionary.
func NewFilter(d FilterDict) (Filter, error) {
	_ge, _cbb := _aga(d.CFM)
	if _cbb != nil {
		return nil, _cbb
	}
	_bfb, _cbb := _ge(d)
	if _cbb != nil {
		return nil, _cbb
	}
	return _bfb, nil
}

var _ Filter = filterV2{}

// NewFilterV2 creates a RC4-based filter with a specified key length (in bytes).
func NewFilterV2(length int) Filter {
	_bab, _ab := _gd(FilterDict{Length: length})
	if _ab != nil {
		_dg.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020R\u0043\u0034\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _ab)
		return filterV2{_ae: length}
	}
	return _bab
}

// KeyLength implements Filter interface.
func (filterAESV3) KeyLength() int { return 256 / 8 }
func init()                        { _ede("\u0056\u0032", _gd) }

// HandlerVersion implements Filter interface.
func (filterAESV3) HandlerVersion() (V, R int)                            { V, R = 5, 6; return }
func (filterIdentity) DecryptBytes(p []byte, okey []byte) ([]byte, error) { return p, nil }

// DecryptBytes implements Filter interface.
func (filterV2) DecryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_cgda, _ac := _a.NewCipher(okey)
	if _ac != nil {
		return nil, _ac
	}
	_dg.Log.Trace("\u0052\u00434\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078", buf)
	_cgda.XORKeyStream(buf, buf)
	_dg.Log.Trace("\u0074o\u003a\u0020\u0025\u0020\u0078", buf)
	return buf, nil
}

// NewFilterAESV2 creates an AES-based filter with a 128 bit key (AESV2).
func NewFilterAESV2() Filter {
	_bb, _ad := _cf(FilterDict{})
	if _ad != nil {
		_dg.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _ad)
		return filterAESV2{}
	}
	return _bb
}

type filterV2 struct{ _ae int }

func (filterAES) DecryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_ced, _ba := _e.NewCipher(okey)
	if _ba != nil {
		return nil, _ba
	}
	if len(buf) < 16 {
		_dg.Log.Debug("\u0045R\u0052\u004f\u0052\u0020\u0041\u0045\u0053\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0062\u0075\u0066\u0020\u0025\u0073", buf)
		return buf, _b.Errorf("\u0041\u0045\u0053\u003a B\u0075\u0066\u0020\u006c\u0065\u006e\u0020\u003c\u0020\u0031\u0036\u0020\u0028\u0025d\u0029", len(buf))
	}
	_ddg := buf[:16]
	buf = buf[16:]
	if len(buf)%16 != 0 {
		_dg.Log.Debug("\u0020\u0069\u0076\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(_ddg), _ddg)
		_dg.Log.Debug("\u0062\u0075\u0066\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
		return buf, _b.Errorf("\u0041\u0045\u0053\u0020\u0062\u0075\u0066\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069p\u006c\u0065\u0020\u006f\u0066 \u0031\u0036 \u0028\u0025\u0064\u0029", len(buf))
	}
	_ca := _d.NewCBCDecrypter(_ced, _ddg)
	_dg.Log.Trace("A\u0045\u0053\u0020\u0044ec\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	_dg.Log.Trace("\u0063\u0068\u006f\u0070\u0020\u0041\u0045\u0053\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u0020\u0028\u0025\u0064\u0029\u003a \u0025\u0020\u0078", len(buf), buf)
	_ca.CryptBlocks(buf, buf)
	_dg.Log.Trace("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	if len(buf) == 0 {
		_dg.Log.Trace("\u0045\u006d\u0070\u0074\u0079\u0020b\u0075\u0066\u002c\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0065\u006d\u0070\u0074\u0079\u0020\u0073t\u0072\u0069\u006e\u0067")
		return buf, nil
	}
	_fba := int(buf[len(buf)-1])
	if _fba > len(buf) {
		_dg.Log.Debug("\u0049\u006c\u006c\u0065g\u0061\u006c\u0020\u0070\u0061\u0064\u0020\u006c\u0065\u006eg\u0074h\u0020\u0028\u0025\u0064\u0020\u003e\u0020%\u0064\u0029", _fba, len(buf))
		return buf, _b.Errorf("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0070a\u0064\u0020l\u0065\u006e\u0067\u0074\u0068")
	}
	buf = buf[:len(buf)-_fba]
	return buf, nil
}
func (filterIdentity) EncryptBytes(p []byte, okey []byte) ([]byte, error) { return p, nil }

// MakeKey implements Filter interface.
func (filterAESV3) MakeKey(_, _ uint32, ekey []byte) ([]byte, error) { return ekey, nil }

// EncryptBytes implements Filter interface.
func (filterV2) EncryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_gf, _fa := _a.NewCipher(okey)
	if _fa != nil {
		return nil, _fa
	}
	_dg.Log.Trace("\u0052\u00434\u0020\u0045\u006ec\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078", buf)
	_gf.XORKeyStream(buf, buf)
	_dg.Log.Trace("\u0074o\u003a\u0020\u0025\u0020\u0078", buf)
	return buf, nil
}

type filterIdentity struct{}

func _adg(_cgd FilterDict) (Filter, error) {
	if _cgd.Length == 256 {
		_dg.Log.Debug("\u0041\u0045S\u0056\u0033\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029", _cgd.Length)
		_cgd.Length /= 8
	}
	if _cgd.Length != 0 && _cgd.Length != 32 {
		return nil, _b.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0033\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029", _cgd.Length)
	}
	return filterAESV3{}, nil
}

// NewIdentity creates an identity filter that bypasses all data without changes.
func NewIdentity() Filter { return filterIdentity{} }

// PDFVersion implements Filter interface.
func (filterAESV3) PDFVersion() [2]int { return [2]int{2, 0} }

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
	MakeKey(_be, _cag uint32, _bbc []byte) ([]byte, error)

	// EncryptBytes encrypts a buffer using object encryption key, as returned by MakeKey.
	// Implementation may reuse a buffer and encrypt data in-place.
	EncryptBytes(_bdc []byte, _def []byte) ([]byte, error)

	// DecryptBytes decrypts a buffer using object encryption key, as returned by MakeKey.
	// Implementation may reuse a buffer and decrypt data in-place.
	DecryptBytes(_agb []byte, _cb []byte) ([]byte, error)
}

// FilterDict represents information from a CryptFilter dictionary.
type FilterDict struct {
	CFM       string
	AuthEvent _ce.AuthEvent
	Length    int
}

func _ede(_ga string, _fac filterFunc) {
	if _, _acb := _dac[_ga]; _acb {
		panic("\u0061l\u0072e\u0061\u0064\u0079\u0020\u0072e\u0067\u0069s\u0074\u0065\u0072\u0065\u0064")
	}
	_dac[_ga] = _fac
}
func _cf(_af FilterDict) (Filter, error) {
	if _af.Length == 128 {
		_dg.Log.Debug("\u0041\u0045S\u0056\u0032\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029", _af.Length)
		_af.Length /= 8
	}
	if _af.Length != 0 && _af.Length != 16 {
		return nil, _b.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0032\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029", _af.Length)
	}
	return filterAESV2{}, nil
}

// NewFilterAESV3 creates an AES-based filter with a 256 bit key (AESV3).
func NewFilterAESV3() Filter {
	_cg, _bg := _adg(FilterDict{})
	if _bg != nil {
		_dg.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0033\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076", _bg)
		return filterAESV3{}
	}
	return _cg
}

// MakeKey implements Filter interface.
func (_cd filterV2) MakeKey(objNum, genNum uint32, ekey []byte) ([]byte, error) {
	return _bc(objNum, genNum, ekey, false)
}

// PDFVersion implements Filter interface.
func (filterAESV2) PDFVersion() [2]int { return [2]int{1, 5} }

var (
	_dac = make(map[string]filterFunc)
)

// MakeKey implements Filter interface.
func (filterAESV2) MakeKey(objNum, genNum uint32, ekey []byte) ([]byte, error) {
	return _bc(objNum, genNum, ekey, true)
}

// HandlerVersion implements Filter interface.
func (filterAESV2) HandlerVersion() (V, R int) { V, R = 4, 4; return }

// HandlerVersion implements Filter interface.
func (_bf filterV2) HandlerVersion() (V, R int)                                   { V, R = 2, 3; return }
func (filterIdentity) MakeKey(objNum, genNum uint32, fkey []byte) ([]byte, error) { return fkey, nil }

// KeyLength implements Filter interface.
func (filterAESV2) KeyLength() int { return 128 / 8 }
func (filterAES) EncryptBytes(buf []byte, okey []byte) ([]byte, error) {
	_de, _da := _e.NewCipher(okey)
	if _da != nil {
		return nil, _da
	}
	_dg.Log.Trace("A\u0045\u0053\u0020\u0045nc\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	const _bd = _e.BlockSize
	_fb := _bd - len(buf)%_bd
	for _dd := 0; _dd < _fb; _dd++ {
		buf = append(buf, byte(_fb))
	}
	_dg.Log.Trace("\u0050a\u0064d\u0065\u0064\u0020\u0074\u006f \u0025\u0064 \u0062\u0079\u0074\u0065\u0073", len(buf))
	_ed := make([]byte, _bd+len(buf))
	_edb := _ed[:_bd]
	if _, _cgc := _g.ReadFull(_gg.Reader, _edb); _cgc != nil {
		return nil, _cgc
	}
	_ag := _d.NewCBCEncrypter(_de, _edb)
	_ag.CryptBlocks(_ed[_bd:], buf)
	buf = _ed
	_dg.Log.Trace("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078", len(buf), buf)
	return buf, nil
}
func _gd(_dda FilterDict) (Filter, error) {
	if _dda.Length%8 != 0 {
		return nil, _b.Errorf("\u0063\u0072\u0079p\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020o\u0066\u0020\u0038\u0020\u0028\u0025\u0064\u0029", _dda.Length)
	}
	if _dda.Length < 5 || _dda.Length > 16 {
		if _dda.Length == 40 || _dda.Length == 64 || _dda.Length == 128 {
			_dg.Log.Debug("\u0053\u0054\u0041\u004e\u0044AR\u0044\u0020V\u0049\u004f\u004c\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0043\u0072\u0079\u0070\u0074\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072s\u0020\u0074\u006f \u0062\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u0074\u0073\u0020\u0072\u0061t\u0068\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0020-\u0020\u0061s\u0073u\u006d\u0069\u006e\u0067\u0020\u0062\u0069t\u0073\u0020\u0028\u0025\u0064\u0029", _dda.Length)
			_dda.Length /= 8
		} else {
			return nil, _b.Errorf("\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074h\u0020\u006e\u006f\u0074\u0020\u0069\u006e \u0072\u0061\u006e\u0067\u0065\u0020\u0034\u0030\u0020\u002d\u00201\u0032\u0038\u0020\u0062\u0069\u0074\u0020\u0028\u0025\u0064\u0029", _dda.Length)
		}
	}
	return filterV2{_ae: _dda.Length}, nil
}
