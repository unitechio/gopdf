package security

import (
	_cg "bytes"
	_fc "crypto/aes"
	_ga "crypto/cipher"
	_ea "crypto/md5"
	_e "crypto/rand"
	_ge "crypto/rc4"
	_bg "crypto/sha256"
	_gc "crypto/sha512"
	_c "encoding/binary"
	_d "errors"
	_a "fmt"
	_b "hash"
	_f "io"
	_gcd "math"

	_ed "bitbucket.org/shenghui0779/gopdf/common"
)

var _ StdHandler = stdHandlerR6{}

func (_aff errInvalidField) Error() string {
	return _a.Sprintf("\u0025s\u003a\u0020e\u0078\u0070\u0065\u0063t\u0065\u0064\u0020%\u0073\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074o \u0062\u0065\u0020%\u0064\u0020b\u0079\u0074\u0065\u0073\u002c\u0020g\u006f\u0074 \u0025\u0064", _aff.Func, _aff.Field, _aff.Exp, _aff.Got)
}
func (_egg stdHandlerR6) alg2b(R int, _eab, _cda, _bed []byte) ([]byte, error) {
	if R == 5 {
		return _ffc(_eab)
	}
	return _cac(_eab, _cda, _bed)
}

type stdHandlerR6 struct{}

// GenerateParams generates and sets O and U parameters for the encryption dictionary.
// It expects R, P and EncryptMetadata fields to be set.
func (_cab stdHandlerR4) GenerateParams(d *StdEncryptDict, opass, upass []byte) ([]byte, error) {
	O, _aec := _cab.alg3(d.R, upass, opass)
	if _aec != nil {
		_ed.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", _aec)
		return nil, _aec
	}
	d.O = O
	_ed.Log.Trace("\u0067\u0065\u006e\u0020\u004f\u003a\u0020\u0025\u0020\u0078", O)
	_dab := _cab.alg2(d, upass)
	U, _aec := _cab.alg5(_dab, upass)
	if _aec != nil {
		_ed.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", _aec)
		return nil, _aec
	}
	d.U = U
	_ed.Log.Trace("\u0067\u0065\u006e\u0020\u0055\u003a\u0020\u0025\u0020\u0078", U)
	return _dab, nil
}

// Authenticate implements StdHandler interface.
func (_fgb stdHandlerR4) Authenticate(d *StdEncryptDict, pass []byte) ([]byte, Permissions, error) {
	_ed.Log.Trace("\u0044\u0065b\u0075\u0067\u0067\u0069n\u0067\u0020a\u0075\u0074\u0068\u0065\u006e\u0074\u0069\u0063a\u0074\u0069\u006f\u006e\u0020\u002d\u0020\u006f\u0077\u006e\u0065\u0072 \u0070\u0061\u0073\u0073")
	_afe, _gda := _fgb.alg7(d, pass)
	if _gda != nil {
		return nil, 0, _gda
	}
	if _afe != nil {
		_ed.Log.Trace("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e")
		return _afe, PermOwner, nil
	}
	_ed.Log.Trace("\u0044\u0065bu\u0067\u0067\u0069n\u0067\u0020\u0061\u0075the\u006eti\u0063\u0061\u0074\u0069\u006f\u006e\u0020- \u0075\u0073\u0065\u0072\u0020\u0070\u0061s\u0073")
	_afe, _gda = _fgb.alg6(d, pass)
	if _gda != nil {
		return nil, 0, _gda
	}
	if _afe != nil {
		_ed.Log.Trace("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e")
		return _afe, d.P, nil
	}
	return nil, 0, nil
}
func (_fe stdHandlerR4) alg7(_cbf *StdEncryptDict, _aa []byte) ([]byte, error) {
	_daa := _fe.alg3Key(_cbf.R, _aa)
	_ecce := make([]byte, len(_cbf.O))
	if _cbf.R == 2 {
		_gfec, _acf := _ge.NewCipher(_daa)
		if _acf != nil {
			return nil, _d.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072")
		}
		_gfec.XORKeyStream(_ecce, _cbf.O)
	} else if _cbf.R >= 3 {
		_bff := append([]byte{}, _cbf.O...)
		for _ceg := 0; _ceg < 20; _ceg++ {
			_dfb := append([]byte{}, _daa...)
			for _cee := 0; _cee < len(_daa); _cee++ {
				_dfb[_cee] ^= byte(19 - _ceg)
			}
			_dc, _aea := _ge.NewCipher(_dfb)
			if _aea != nil {
				return nil, _d.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072")
			}
			_dc.XORKeyStream(_ecce, _bff)
			_bff = append([]byte{}, _ecce...)
		}
	} else {
		return nil, _d.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R")
	}
	_fcff, _fef := _fe.alg6(_cbf, _ecce)
	if _fef != nil {
		return nil, nil
	}
	return _fcff, nil
}

// NewHandlerR4 creates a new standard security handler for R<=4.
func NewHandlerR4(id0 string, length int) StdHandler { return stdHandlerR4{ID0: id0, Length: length} }
func (_abe *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%_abe._ab != 0 {
		_ed.Log.Error("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073")
		return
	}
	if len(dst) < len(src) {
		_ed.Log.Error("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074")
		return
	}
	for len(src) > 0 {
		_abe._gg.Encrypt(dst, src[:_abe._ab])
		src = src[_abe._ab:]
		dst = dst[_abe._ab:]
	}
}
func _cac(_fdd, _cga, _ecea []byte) ([]byte, error) {
	var (
		_efg, _eag, _acg _b.Hash
	)
	_efg = _bg.New()
	_eee := make([]byte, 64)
	_bgb := _efg
	_bgb.Write(_fdd)
	K := _bgb.Sum(_eee[:0])
	_bba := make([]byte, 64*(127+64+48))
	_fce := func(_dabe int) ([]byte, error) {
		_aba := len(_cga) + len(K) + len(_ecea)
		_bfe := _bba[:_aba]
		_bec := copy(_bfe, _cga)
		_bec += copy(_bfe[_bec:], K[:])
		_bec += copy(_bfe[_bec:], _ecea)
		if _bec != _aba {
			_ed.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0072\u006f\u0075\u006ed\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u007ae\u002e")
			return nil, _d.New("\u0077\u0072\u006f\u006e\u0067\u0020\u0073\u0069\u007a\u0065")
		}
		K1 := _bba[:_aba*64]
		_dcg(K1, _aba)
		_fcfa, _gfd := _ebe(K[0:16])
		if _gfd != nil {
			return nil, _gfd
		}
		_dega := _ga.NewCBCEncrypter(_fcfa, K[16:32])
		_dega.CryptBlocks(K1, K1)
		E := K1
		_abg := 0
		for _ffd := 0; _ffd < 16; _ffd++ {
			_abg += int(E[_ffd] % 3)
		}
		var _gaee _b.Hash
		switch _abg % 3 {
		case 0:
			_gaee = _efg
		case 1:
			if _eag == nil {
				_eag = _gc.New384()
			}
			_gaee = _eag
		case 2:
			if _acg == nil {
				_acg = _gc.New()
			}
			_gaee = _acg
		}
		_gaee.Reset()
		_gaee.Write(E)
		K = _gaee.Sum(_eee[:0])
		return E, nil
	}
	for _agf := 0; ; {
		E, _daf := _fce(_agf)
		if _daf != nil {
			return nil, _daf
		}
		_bccb := E[len(E)-1]
		_agf++
		if _agf >= 64 && _bccb <= uint8(_agf-32) {
			break
		}
	}
	return K[:32], nil
}

type ecbDecrypter ecb

func (_bd stdHandlerR4) alg4(_deg []byte, _fae []byte) ([]byte, error) {
	_cbg, _aed := _ge.NewCipher(_deg)
	if _aed != nil {
		return nil, _d.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_gdc := []byte(_cb)
	_gde := make([]byte, len(_gdc))
	_cbg.XORKeyStream(_gde, _gdc)
	return _gde, nil
}

// AuthEvent is an event type that triggers authentication.
type AuthEvent string

func (_fada stdHandlerR6) alg13(_gedc *StdEncryptDict, _dgbc []byte) error {
	if _abgb := _ac("\u0061\u006c\u00671\u0033", "\u004b\u0065\u0079", 32, _dgbc); _abgb != nil {
		return _abgb
	}
	if _bbg := _ac("\u0061\u006c\u00671\u0033", "\u0050\u0065\u0072m\u0073", 16, _gedc.Perms); _bbg != nil {
		return _bbg
	}
	_geg := make([]byte, 16)
	copy(_geg, _gedc.Perms[:16])
	_cbe, _egge := _fc.NewCipher(_dgbc[:32])
	if _egge != nil {
		return _egge
	}
	_eec := _fb(_cbe)
	_eec.CryptBlocks(_geg, _geg)
	if !_cg.Equal(_geg[9:12], []byte("\u0061\u0064\u0062")) {
		return _d.New("\u0064\u0065\u0063o\u0064\u0065\u0064\u0020p\u0065\u0072\u006d\u0069\u0073\u0073\u0069o\u006e\u0073\u0020\u0061\u0072\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_abee := Permissions(_c.LittleEndian.Uint32(_geg[0:4]))
	if _abee != _gedc.P {
		return _d.New("\u0070\u0065r\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0061il\u0065\u0064")
	}
	var _gfg bool
	if _geg[8] == 'T' {
		_gfg = true
	} else if _geg[8] == 'F' {
		_gfg = false
	} else {
		return _d.New("\u0064\u0065\u0063\u006f\u0064\u0065\u0064 \u006d\u0065\u0074a\u0064\u0061\u0074\u0061 \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0067\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	if _gfg != _gedc.EncryptMetadata {
		return _d.New("\u006d\u0065t\u0061\u0064\u0061\u0074a\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e \u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066a\u0069\u006c\u0065\u0064")
	}
	return nil
}
func (_afeg stdHandlerR6) alg12(_dgg *StdEncryptDict, _agg []byte) ([]byte, error) {
	if _gdec := _ac("\u0061\u006c\u00671\u0032", "\u0055", 48, _dgg.U); _gdec != nil {
		return nil, _gdec
	}
	if _cdeg := _ac("\u0061\u006c\u00671\u0032", "\u004f", 48, _dgg.O); _cdeg != nil {
		return nil, _cdeg
	}
	_fad := make([]byte, len(_agg)+8+48)
	_egc := copy(_fad, _agg)
	_egc += copy(_fad[_egc:], _dgg.O[32:40])
	_egc += copy(_fad[_egc:], _dgg.U[0:48])
	_egcf, _bfeb := _afeg.alg2b(_dgg.R, _fad, _agg, _dgg.U[0:48])
	if _bfeb != nil {
		return nil, _bfeb
	}
	_egcf = _egcf[:32]
	if !_cg.Equal(_egcf, _dgg.O[:32]) {
		return nil, nil
	}
	return _egcf, nil
}

var _ StdHandler = stdHandlerR4{}

func (_bc stdHandlerR4) alg3(R int, _bef, _cf []byte) ([]byte, error) {
	var _ba []byte
	if len(_cf) > 0 {
		_ba = _bc.alg3Key(R, _cf)
	} else {
		_ba = _bc.alg3Key(R, _bef)
	}
	_bbd, _ecc := _ge.NewCipher(_ba)
	if _ecc != nil {
		return nil, _d.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_cgf := _bc.paddedPass(_bef)
	_gfb := make([]byte, len(_cgf))
	_bbd.XORKeyStream(_gfb, _cgf)
	if R >= 3 {
		_fbd := make([]byte, len(_ba))
		for _gd := 0; _gd < 19; _gd++ {
			for _cd := 0; _cd < len(_ba); _cd++ {
				_fbd[_cd] = _ba[_cd] ^ byte(_gd+1)
			}
			_fcf, _gdf := _ge.NewCipher(_fbd)
			if _gdf != nil {
				return nil, _d.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
			}
			_fcf.XORKeyStream(_gfb, _gfb)
		}
	}
	return _gfb, nil
}
func (_ff stdHandlerR4) alg3Key(R int, _afff []byte) []byte {
	_afffb := _ea.New()
	_gfe := _ff.paddedPass(_afff)
	_afffb.Write(_gfe)
	if R >= 3 {
		for _edg := 0; _edg < 50; _edg++ {
			_ag := _afffb.Sum(nil)
			_afffb = _ea.New()
			_afffb.Write(_ag)
		}
	}
	_bb := _afffb.Sum(nil)
	if R == 2 {
		_bb = _bb[0:5]
	} else {
		_bb = _bb[0 : _ff.Length/8]
	}
	return _bb
}
func (_aeg stdHandlerR4) alg6(_eg *StdEncryptDict, _bdd []byte) ([]byte, error) {
	var (
		_abea []byte
		_gaeb error
	)
	_fde := _aeg.alg2(_eg, _bdd)
	if _eg.R == 2 {
		_abea, _gaeb = _aeg.alg4(_fde, _bdd)
	} else if _eg.R >= 3 {
		_abea, _gaeb = _aeg.alg5(_fde, _bdd)
	} else {
		return nil, _d.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R")
	}
	if _gaeb != nil {
		return nil, _gaeb
	}
	_ed.Log.Trace("\u0063\u0068\u0065\u0063k:\u0020\u0025\u0020\u0078\u0020\u003d\u003d\u0020\u0025\u0020\u0078\u0020\u003f", string(_abea), string(_eg.U))
	_gag := _abea
	_dfe := _eg.U
	if _eg.R >= 3 {
		if len(_gag) > 16 {
			_gag = _gag[0:16]
		}
		if len(_dfe) > 16 {
			_dfe = _dfe[0:16]
		}
	}
	if !_cg.Equal(_gag, _dfe) {
		return nil, nil
	}
	return _fde, nil
}

type stdHandlerR4 struct {
	Length int
	ID0    string
}

const _cb = "\x28\277\116\136\x4e\x75\x8a\x41\x64\000\x4e\x56\377" + "\xfa\001\010\056\x2e\x00\xb6\xd0\x68\076\x80\x2f\014" + "\251\xfe\x64\x53\x69\172"

type ecb struct {
	_gg _ga.Block
	_ab int
}

func (_eage stdHandlerR6) alg11(_feb *StdEncryptDict, _dga []byte) ([]byte, error) {
	if _gbd := _ac("\u0061\u006c\u00671\u0031", "\u0055", 48, _feb.U); _gbd != nil {
		return nil, _gbd
	}
	_edfg := make([]byte, len(_dga)+8)
	_gadf := copy(_edfg, _dga)
	_gadf += copy(_edfg[_gadf:], _feb.U[32:40])
	_accd, _cgg := _eage.alg2b(_feb.R, _edfg, _dga, nil)
	if _cgg != nil {
		return nil, _cgg
	}
	_accd = _accd[:32]
	if !_cg.Equal(_accd, _feb.U[:32]) {
		return nil, nil
	}
	return _accd, nil
}
func (stdHandlerR4) paddedPass(_caf []byte) []byte {
	_dd := make([]byte, 32)
	_ae := copy(_dd, _caf)
	for ; _ae < 32; _ae++ {
		_dd[_ae] = _cb[_ae-len(_caf)]
	}
	return _dd
}

// NewHandlerR6 creates a new standard security handler for R=5 and R=6.
func NewHandlerR6() StdHandler { return stdHandlerR6{} }
func _dcg(_gfa []byte, _ccce int) {
	_acc := _ccce
	for _acc < len(_gfa) {
		copy(_gfa[_acc:], _gfa[:_acc])
		_acc *= 2
	}
}
func _eb(_fa _ga.Block) *ecb { return &ecb{_gg: _fa, _ab: _fa.BlockSize()} }

const (
	PermOwner             = Permissions(_gcd.MaxUint32)
	PermPrinting          = Permissions(1 << 2)
	PermModify            = Permissions(1 << 3)
	PermExtractGraphics   = Permissions(1 << 4)
	PermAnnotate          = Permissions(1 << 5)
	PermFillForms         = Permissions(1 << 8)
	PermDisabilityExtract = Permissions(1 << 9)
	PermRotateInsert      = Permissions(1 << 10)
	PermFullPrintQuality  = Permissions(1 << 11)
)

// Permissions is a bitmask of access permissions for a PDF file.
type Permissions uint32

func (_da stdHandlerR4) alg2(_ced *StdEncryptDict, _be []byte) []byte {
	_ed.Log.Trace("\u0061\u006c\u0067\u0032")
	_gga := _da.paddedPass(_be)
	_de := _ea.New()
	_de.Write(_gga)
	_de.Write(_ced.O)
	var _eff [4]byte
	_c.LittleEndian.PutUint32(_eff[:], uint32(_ced.P))
	_de.Write(_eff[:])
	_ed.Log.Trace("\u0067o\u0020\u0050\u003a\u0020\u0025\u0020x", _eff)
	_de.Write([]byte(_da.ID0))
	_ed.Log.Trace("\u0074\u0068\u0069\u0073\u002e\u0052\u0020\u003d\u0020\u0025d\u0020\u0065\u006e\u0063\u0072\u0079\u0070t\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u0025\u0076", _ced.R, _ced.EncryptMetadata)
	if (_ced.R >= 4) && !_ced.EncryptMetadata {
		_de.Write([]byte{0xff, 0xff, 0xff, 0xff})
	}
	_db := _de.Sum(nil)
	if _ced.R >= 3 {
		_de = _ea.New()
		for _ggc := 0; _ggc < 50; _ggc++ {
			_de.Reset()
			_de.Write(_db[0 : _da.Length/8])
			_db = _de.Sum(nil)
		}
	}
	if _ced.R >= 3 {
		return _db[0 : _da.Length/8]
	}
	return _db[0:5]
}

// StdHandler is an interface for standard security handlers.
type StdHandler interface {

	// GenerateParams uses owner and user passwords to set encryption parameters and generate an encryption key.
	// It assumes that R, P and EncryptMetadata are already set.
	GenerateParams(_ef *StdEncryptDict, _gf, _gb []byte) ([]byte, error)

	// Authenticate uses encryption dictionary parameters and the password to calculate
	// the document encryption key. It also returns permissions that should be granted to a user.
	// In case of failed authentication, it returns empty key and zero permissions with no error.
	Authenticate(_gaa *StdEncryptDict, _ca []byte) ([]byte, Permissions, error)
}

func (_ebd stdHandlerR6) alg8(_fed *StdEncryptDict, _dda []byte, _gec []byte) error {
	if _bde := _ac("\u0061\u006c\u0067\u0038", "\u004b\u0065\u0079", 32, _dda); _bde != nil {
		return _bde
	}
	var _fee [16]byte
	if _, _dbc := _f.ReadFull(_e.Reader, _fee[:]); _dbc != nil {
		return _dbc
	}
	_eabf := _fee[0:8]
	_ebde := _fee[8:16]
	_dgf := make([]byte, len(_gec)+len(_eabf))
	_afb := copy(_dgf, _gec)
	copy(_dgf[_afb:], _eabf)
	_fbc, _cgfb := _ebd.alg2b(_fed.R, _dgf, _gec, nil)
	if _cgfb != nil {
		return _cgfb
	}
	U := make([]byte, len(_fbc)+len(_eabf)+len(_ebde))
	_afb = copy(U, _fbc[:32])
	_afb += copy(U[_afb:], _eabf)
	copy(U[_afb:], _ebde)
	_fed.U = U
	_afb = len(_gec)
	copy(_dgf[_afb:], _ebde)
	_fbc, _cgfb = _ebd.alg2b(_fed.R, _dgf, _gec, nil)
	if _cgfb != nil {
		return _cgfb
	}
	_cgag, _cgfb := _ebe(_fbc[:32])
	if _cgfb != nil {
		return _cgfb
	}
	_ffa := make([]byte, _fc.BlockSize)
	_cgfd := _ga.NewCBCEncrypter(_cgag, _ffa)
	UE := make([]byte, 32)
	_cgfd.CryptBlocks(UE, _dda[:32])
	_fed.UE = UE
	return nil
}
func _ffc(_ege []byte) ([]byte, error) {
	_dbb := _bg.New()
	_dbb.Write(_ege)
	return _dbb.Sum(nil), nil
}
func _ac(_fd, _df string, _ggg int, _dfg []byte) error {
	if len(_dfg) < _ggg {
		return errInvalidField{Func: _fd, Field: _df, Exp: _ggg, Got: len(_dfg)}
	}
	return nil
}
func (_gae *ecbEncrypter) BlockSize() int { return _gae._ab }

type ecbEncrypter ecb

func (_eabc stdHandlerR6) alg9(_ccd *StdEncryptDict, _ebdc []byte, _cca []byte) error {
	if _cffb := _ac("\u0061\u006c\u0067\u0039", "\u004b\u0065\u0079", 32, _ebdc); _cffb != nil {
		return _cffb
	}
	if _dgb := _ac("\u0061\u006c\u0067\u0039", "\u0055", 48, _ccd.U); _dgb != nil {
		return _dgb
	}
	var _eeef [16]byte
	if _, _afd := _f.ReadFull(_e.Reader, _eeef[:]); _afd != nil {
		return _afd
	}
	_age := _eeef[0:8]
	_fea := _eeef[8:16]
	_dfgg := _ccd.U[:48]
	_aac := make([]byte, len(_cca)+len(_age)+len(_dfgg))
	_bece := copy(_aac, _cca)
	_bece += copy(_aac[_bece:], _age)
	_bece += copy(_aac[_bece:], _dfgg)
	_bab, _bbc := _eabc.alg2b(_ccd.R, _aac, _cca, _dfgg)
	if _bbc != nil {
		return _bbc
	}
	O := make([]byte, len(_bab)+len(_age)+len(_fea))
	_bece = copy(O, _bab[:32])
	_bece += copy(O[_bece:], _age)
	_bece += copy(O[_bece:], _fea)
	_ccd.O = O
	_bece = len(_cca)
	_bece += copy(_aac[_bece:], _fea)
	_bab, _bbc = _eabc.alg2b(_ccd.R, _aac, _cca, _dfgg)
	if _bbc != nil {
		return _bbc
	}
	_becd, _bbc := _ebe(_bab[:32])
	if _bbc != nil {
		return _bbc
	}
	_edf := make([]byte, _fc.BlockSize)
	_aaee := _ga.NewCBCEncrypter(_becd, _edf)
	OE := make([]byte, 32)
	_aaee.CryptBlocks(OE, _ebdc[:32])
	_ccd.OE = OE
	return nil
}

// Allowed checks if a set of permissions can be granted.
func (_ad Permissions) Allowed(p2 Permissions) bool { return _ad&p2 == p2 }
func (_gef stdHandlerR6) alg2a(_bac *StdEncryptDict, _dg []byte) ([]byte, Permissions, error) {
	if _bgg := _ac("\u0061\u006c\u00672\u0061", "\u004f", 48, _bac.O); _bgg != nil {
		return nil, 0, _bgg
	}
	if _bcc := _ac("\u0061\u006c\u00672\u0061", "\u0055", 48, _bac.U); _bcc != nil {
		return nil, 0, _bcc
	}
	if len(_dg) > 127 {
		_dg = _dg[:127]
	}
	_abf, _egb := _gef.alg12(_bac, _dg)
	if _egb != nil {
		return nil, 0, _egb
	}
	var (
		_gff  []byte
		_bbdb []byte
		_ged  []byte
	)
	var _aae Permissions
	if len(_abf) != 0 {
		_aae = PermOwner
		_gcf := make([]byte, len(_dg)+8+48)
		_ccg := copy(_gcf, _dg)
		_ccg += copy(_gcf[_ccg:], _bac.O[40:48])
		copy(_gcf[_ccg:], _bac.U[0:48])
		_gff = _gcf
		_bbdb = _bac.OE
		_ged = _bac.U[0:48]
	} else {
		_abf, _egb = _gef.alg11(_bac, _dg)
		if _egb == nil && len(_abf) == 0 {
			_abf, _egb = _gef.alg11(_bac, []byte(""))
		}
		if _egb != nil {
			return nil, 0, _egb
		} else if len(_abf) == 0 {
			return nil, 0, nil
		}
		_aae = _bac.P
		_dfeb := make([]byte, len(_dg)+8)
		_dad := copy(_dfeb, _dg)
		copy(_dfeb[_dad:], _bac.U[40:48])
		_gff = _dfeb
		_bbdb = _bac.UE
		_ged = nil
	}
	if _aee := _ac("\u0061\u006c\u00672\u0061", "\u004b\u0065\u0079", 32, _bbdb); _aee != nil {
		return nil, 0, _aee
	}
	_bbdb = _bbdb[:32]
	_afa, _egb := _gef.alg2b(_bac.R, _gff, _dg, _ged)
	if _egb != nil {
		return nil, 0, _egb
	}
	_cbgg, _egb := _fc.NewCipher(_afa[:32])
	if _egb != nil {
		return nil, 0, _egb
	}
	_fgd := make([]byte, _fc.BlockSize)
	_gaga := _ga.NewCBCDecrypter(_cbgg, _fgd)
	_ece := make([]byte, 32)
	_gaga.CryptBlocks(_ece, _bbdb)
	if _bac.R == 5 {
		return _ece, _aae, nil
	}
	_egb = _gef.alg13(_bac, _ece)
	if _egb != nil {
		return nil, 0, _egb
	}
	return _ece, _aae, nil
}
func (_cc *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%_cc._ab != 0 {
		_ed.Log.Error("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073")
		return
	}
	if len(dst) < len(src) {
		_ed.Log.Error("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0064\u0065\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074")
		return
	}
	for len(src) > 0 {
		_cc._gg.Decrypt(dst, src[:_cc._ab])
		src = src[_cc._ab:]
		dst = dst[_cc._ab:]
	}
}

// Authenticate implements StdHandler interface.
func (_gce stdHandlerR6) Authenticate(d *StdEncryptDict, pass []byte) ([]byte, Permissions, error) {
	return _gce.alg2a(d, pass)
}
func _ebe(_cbc []byte) (_ga.Block, error) {
	_gbg, _ee := _fc.NewCipher(_cbc)
	if _ee != nil {
		_ed.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0063\u0069p\u0068\u0065r\u003a\u0020\u0025\u0076", _ee)
		return nil, _ee
	}
	return _gbg, nil
}

// StdEncryptDict is a set of additional fields used in standard encryption dictionary.
type StdEncryptDict struct {
	R               int
	P               Permissions
	EncryptMetadata bool
	O, U            []byte
	OE, UE          []byte
	Perms           []byte
}

const (
	EventDocOpen = AuthEvent("\u0044o\u0063\u004f\u0070\u0065\u006e")
	EventEFOpen  = AuthEvent("\u0045\u0046\u004f\u0070\u0065\u006e")
)

type errInvalidField struct {
	Func  string
	Field string
	Exp   int
	Got   int
}

func (_cae stdHandlerR4) alg5(_ccc []byte, _cff []byte) ([]byte, error) {
	_eba := _ea.New()
	_eba.Write([]byte(_cb))
	_eba.Write([]byte(_cae.ID0))
	_cde := _eba.Sum(nil)
	_ed.Log.Trace("\u0061\u006c\u0067\u0035")
	_ed.Log.Trace("\u0065k\u0065\u0079\u003a\u0020\u0025\u0020x", _ccc)
	_ed.Log.Trace("\u0049D\u003a\u0020\u0025\u0020\u0078", _cae.ID0)
	if len(_cde) != 16 {
		return nil, _d.New("\u0068a\u0073\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074\u0020\u0031\u0036\u0020\u0062\u0079\u0074\u0065\u0073")
	}
	_aca, _fcg := _ge.NewCipher(_ccc)
	if _fcg != nil {
		return nil, _d.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_fba := make([]byte, 16)
	_aca.XORKeyStream(_fba, _cde)
	_bda := make([]byte, len(_ccc))
	for _geb := 0; _geb < 19; _geb++ {
		for _dbg := 0; _dbg < len(_ccc); _dbg++ {
			_bda[_dbg] = _ccc[_dbg] ^ byte(_geb+1)
		}
		_aca, _fcg = _ge.NewCipher(_bda)
		if _fcg != nil {
			return nil, _d.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
		}
		_aca.XORKeyStream(_fba, _fba)
		_ed.Log.Trace("\u0069\u0020\u003d\u0020\u0025\u0064\u002c\u0020\u0065\u006b\u0065\u0079:\u0020\u0025\u0020\u0078", _geb, _bda)
		_ed.Log.Trace("\u0069\u0020\u003d\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0020\u0078", _geb, _fba)
	}
	_dae := make([]byte, 32)
	for _adb := 0; _adb < 16; _adb++ {
		_dae[_adb] = _fba[_adb]
	}
	_, _fcg = _e.Read(_dae[16:32])
	if _fcg != nil {
		return nil, _d.New("\u0066a\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065n\u0020r\u0061n\u0064\u0020\u006e\u0075\u006d\u0062\u0065r")
	}
	return _dae, nil
}
func _fg(_bf _ga.Block) _ga.BlockMode { return (*ecbEncrypter)(_eb(_bf)) }

// GenerateParams is the algorithm opposite to alg2a (R>=5).
// It generates U,O,UE,OE,Perms fields using AESv3 encryption.
// There is no algorithm number assigned to this function in the spec.
// It expects R, P and EncryptMetadata fields to be set.
func (_cdab stdHandlerR6) GenerateParams(d *StdEncryptDict, opass, upass []byte) ([]byte, error) {
	_cedb := make([]byte, 32)
	if _, _aecg := _f.ReadFull(_e.Reader, _cedb); _aecg != nil {
		return nil, _aecg
	}
	d.U = nil
	d.O = nil
	d.UE = nil
	d.OE = nil
	d.Perms = nil
	if len(upass) > 127 {
		upass = upass[:127]
	}
	if len(opass) > 127 {
		opass = opass[:127]
	}
	if _ebdd := _cdab.alg8(d, _cedb, upass); _ebdd != nil {
		return nil, _ebdd
	}
	if _afc := _cdab.alg9(d, _cedb, opass); _afc != nil {
		return nil, _afc
	}
	if d.R == 5 {
		return _cedb, nil
	}
	if _edgb := _cdab.alg10(d, _cedb); _edgb != nil {
		return nil, _edgb
	}
	return _cedb, nil
}
func (_aga stdHandlerR6) alg10(_fge *StdEncryptDict, _bffa []byte) error {
	if _cede := _ac("\u0061\u006c\u00671\u0030", "\u004b\u0065\u0079", 32, _bffa); _cede != nil {
		return _cede
	}
	_cabd := uint64(uint32(_fge.P)) | (_gcd.MaxUint32 << 32)
	Perms := make([]byte, 16)
	_c.LittleEndian.PutUint64(Perms[:8], _cabd)
	if _fge.EncryptMetadata {
		Perms[8] = 'T'
	} else {
		Perms[8] = 'F'
	}
	copy(Perms[9:12], "\u0061\u0064\u0062")
	if _, _affa := _f.ReadFull(_e.Reader, Perms[12:16]); _affa != nil {
		return _affa
	}
	_ebc, _ccb := _ebe(_bffa[:32])
	if _ccb != nil {
		return _ccb
	}
	_efge := _fg(_ebc)
	_efge.CryptBlocks(Perms, Perms)
	_fge.Perms = Perms[:16]
	return nil
}
func (_af *ecbDecrypter) BlockSize() int { return _af._ab }
func _fb(_ce _ga.Block) _ga.BlockMode    { return (*ecbDecrypter)(_eb(_ce)) }
