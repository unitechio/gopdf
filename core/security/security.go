package security

import (
	_dgd "bytes"
	_dg "crypto/aes"
	_b "crypto/cipher"
	_gc "crypto/md5"
	_gd "crypto/rand"
	_fg "crypto/rc4"
	_c "crypto/sha256"
	_dc "crypto/sha512"
	_bd "encoding/binary"
	_a "errors"
	_e "fmt"
	_f "hash"
	_d "io"
	_fb "math"

	_ab "unitechio/gopdf/gopdf/common"
)

type ecbDecrypter ecb

func _eccag(_baefb []byte) ([]byte, error) {
	_degf := _c.New()
	_degf.Write(_baefb)
	return _degf.Sum(nil), nil
}

func (_egad stdHandlerR6) alg13(_gaf *StdEncryptDict, _bec []byte) error {
	if _edd := _faa("\u0061\u006c\u00671\u0033", "\u004b\u0065\u0079", 32, _bec); _edd != nil {
		return _edd
	}
	if _adbb := _faa("\u0061\u006c\u00671\u0033", "\u0050\u0065\u0072m\u0073", 16, _gaf.Perms); _adbb != nil {
		return _adbb
	}
	_cfaa := make([]byte, 16)
	copy(_cfaa, _gaf.Perms[:16])
	_adfc, _fgd := _dg.NewCipher(_bec[:32])
	if _fgd != nil {
		return _fgd
	}
	_fgda := _gb(_adfc)
	_fgda.CryptBlocks(_cfaa, _cfaa)
	if !_dgd.Equal(_cfaa[9:12], []byte("\u0061\u0064\u0062")) {
		return _a.New("\u0064\u0065\u0063o\u0064\u0065\u0064\u0020p\u0065\u0072\u006d\u0069\u0073\u0073\u0069o\u006e\u0073\u0020\u0061\u0072\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_gdbe := Permissions(_bd.LittleEndian.Uint32(_cfaa[0:4]))
	if _gdbe != _gaf.P {
		return _a.New("\u0070\u0065r\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0061il\u0065\u0064")
	}
	var _feb bool
	if _cfaa[8] == 'T' {
		_feb = true
	} else if _cfaa[8] == 'F' {
		_feb = false
	} else {
		return _a.New("\u0064\u0065\u0063\u006f\u0064\u0065\u0064 \u006d\u0065\u0074a\u0064\u0061\u0074\u0061 \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0067\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	if _feb != _gaf.EncryptMetadata {
		return _a.New("\u006d\u0065t\u0061\u0064\u0061\u0074a\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e \u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066a\u0069\u006c\u0065\u0064")
	}
	return nil
}

const (
	PermOwner             = Permissions(_fb.MaxUint32)
	PermPrinting          = Permissions(1 << 2)
	PermModify            = Permissions(1 << 3)
	PermExtractGraphics   = Permissions(1 << 4)
	PermAnnotate          = Permissions(1 << 5)
	PermFillForms         = Permissions(1 << 8)
	PermDisabilityExtract = Permissions(1 << 9)
	PermRotateInsert      = Permissions(1 << 10)
	PermFullPrintQuality  = Permissions(1 << 11)
)

func (_adb stdHandlerR6) alg2b(R int, _efg, _gdbb, _fgge []byte) ([]byte, error) {
	if R == 5 {
		return _eccag(_efg)
	}
	return _ebf(_efg, _gdbb, _fgge)
}

func (_gba stdHandlerR6) alg11(_baed *StdEncryptDict, _ccf []byte) ([]byte, error) {
	if _bed := _faa("\u0061\u006c\u00671\u0031", "\u0055", 48, _baed.U); _bed != nil {
		return nil, _bed
	}
	_edc := make([]byte, len(_ccf)+8)
	_ebgf := copy(_edc, _ccf)
	_ebgf += copy(_edc[_ebgf:], _baed.U[32:40])
	_dfg, _cbaf := _gba.alg2b(_baed.R, _edc, _ccf, nil)
	if _cbaf != nil {
		return nil, _cbaf
	}
	_dfg = _dfg[:32]
	if !_dgd.Equal(_dfg, _baed.U[:32]) {
		return nil, nil
	}
	return _dfg, nil
}

func (_cc stdHandlerR4) alg2(_fgg *StdEncryptDict, _gbg []byte) []byte {
	_ab.Log.Trace("\u0061\u006c\u0067\u0032")
	_be := _cc.paddedPass(_gbg)
	_bae := _gc.New()
	_bae.Write(_be)
	_bae.Write(_fgg.O)
	var _baef [4]byte
	_bd.LittleEndian.PutUint32(_baef[:], uint32(_fgg.P))
	_bae.Write(_baef[:])
	_ab.Log.Trace("\u0067o\u0020\u0050\u003a\u0020\u0025\u0020x", _baef)
	_bae.Write([]byte(_cc.ID0))
	_ab.Log.Trace("\u0074\u0068\u0069\u0073\u002e\u0052\u0020\u003d\u0020\u0025d\u0020\u0065\u006e\u0063\u0072\u0079\u0070t\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u0025\u0076", _fgg.R, _fgg.EncryptMetadata)
	if (_fgg.R >= 4) && !_fgg.EncryptMetadata {
		_bae.Write([]byte{0xff, 0xff, 0xff, 0xff})
	}
	_gcc := _bae.Sum(nil)
	if _fgg.R >= 3 {
		_bae = _gc.New()
		for _cdc := 0; _cdc < 50; _cdc++ {
			_bae.Reset()
			_bae.Write(_gcc[0 : _cc.Length/8])
			_gcc = _bae.Sum(nil)
		}
	}
	if _fgg.R >= 3 {
		return _gcc[0 : _cc.Length/8]
	}
	return _gcc[0:5]
}
func _ae(_ea _b.Block) _b.BlockMode { return (*ecbEncrypter)(_dfc(_ea)) }
func _gfc(_agd []byte, _bag int) {
	_eee := _bag
	for _eee < len(_agd) {
		copy(_agd[_eee:], _agd[:_eee])
		_eee *= 2
	}
}

// GenerateParams generates and sets O and U parameters for the encryption dictionary.
// It expects R, P and EncryptMetadata fields to be set.
func (_cggb stdHandlerR4) GenerateParams(d *StdEncryptDict, opass, upass []byte) ([]byte, error) {
	O, _cbf := _cggb.alg3(d.R, upass, opass)
	if _cbf != nil {
		_ab.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", _cbf)
		return nil, _cbf
	}
	d.O = O
	_ab.Log.Trace("\u0067\u0065\u006e\u0020\u004f\u003a\u0020\u0025\u0020\u0078", O)
	_gae := _cggb.alg2(d, upass)
	U, _cbf := _cggb.alg5(_gae, upass)
	if _cbf != nil {
		_ab.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", _cbf)
		return nil, _cbf
	}
	d.U = U
	_ab.Log.Trace("\u0067\u0065\u006e\u0020\u0055\u003a\u0020\u0025\u0020\u0078", U)
	return _gae, nil
}

type ecb struct {
	_df _b.Block
	_bf int
}

// StdHandler is an interface for standard security handlers.
type StdHandler interface {
	// GenerateParams uses owner and user passwords to set encryption parameters and generate an encryption key.
	// It assumes that R, P and EncryptMetadata are already set.
	GenerateParams(_fgc *StdEncryptDict, _aef, _dd []byte) ([]byte, error)

	// Authenticate uses encryption dictionary parameters and the password to calculate
	// the document encryption key. It also returns permissions that should be granted to a user.
	// In case of failed authentication, it returns empty key and zero permissions with no error.
	Authenticate(_ee *StdEncryptDict, _fa []byte) ([]byte, Permissions, error)
}

// Allowed checks if a set of permissions can be granted.
func (_aed Permissions) Allowed(p2 Permissions) bool { return _aed&p2 == p2 }

// Permissions is a bitmask of access permissions for a PDF file.
type Permissions uint32

func _ebf(_cac, _fed, _aead []byte) ([]byte, error) {
	var _defd, _ffd, _abde _f.Hash
	_defd = _c.New()
	_acf := make([]byte, 64)
	_cfb := _defd
	_cfb.Write(_cac)
	K := _cfb.Sum(_acf[:0])
	_accg := make([]byte, 64*(127+64+48))
	_bg := func(_dac int) ([]byte, error) {
		_dde := len(_fed) + len(K) + len(_aead)
		_afag := _accg[:_dde]
		_fac := copy(_afag, _fed)
		_fac += copy(_afag[_fac:], K[:])
		_fac += copy(_afag[_fac:], _aead)
		if _fac != _dde {
			_ab.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0072\u006f\u0075\u006ed\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u007ae\u002e")
			return nil, _a.New("\u0077\u0072\u006f\u006e\u0067\u0020\u0073\u0069\u007a\u0065")
		}
		K1 := _accg[:_dde*64]
		_gfc(K1, _dde)
		_dedc, _cfa := _bdd(K[0:16])
		if _cfa != nil {
			return nil, _cfa
		}
		_dab := _b.NewCBCEncrypter(_dedc, K[16:32])
		_dab.CryptBlocks(K1, K1)
		E := K1
		_fdb := 0
		for _fdf := 0; _fdf < 16; _fdf++ {
			_fdb += int(E[_fdf] % 3)
		}
		var _fbbc _f.Hash
		switch _fdb % 3 {
		case 0:
			_fbbc = _defd
		case 1:
			if _ffd == nil {
				_ffd = _dc.New384()
			}
			_fbbc = _ffd
		case 2:
			if _abde == nil {
				_abde = _dc.New()
			}
			_fbbc = _abde
		}
		_fbbc.Reset()
		_fbbc.Write(E)
		K = _fbbc.Sum(_acf[:0])
		return E, nil
	}
	for _dfcd := 0; ; {
		E, _ddb := _bg(_dfcd)
		if _ddb != nil {
			return nil, _ddb
		}
		_ccda := E[len(E)-1]
		_dfcd++
		if _dfcd >= 64 && _ccda <= uint8(_dfcd-32) {
			break
		}
	}
	return K[:32], nil
}
func (_ff *ecbEncrypter) BlockSize() int { return _ff._bf }
func (_ed stdHandlerR4) alg6(_edf *StdEncryptDict, _fbb []byte) ([]byte, error) {
	var (
		_gg  []byte
		_aeb error
	)
	_befe := _ed.alg2(_edf, _fbb)
	if _edf.R == 2 {
		_gg, _aeb = _ed.alg4(_befe, _fbb)
	} else if _edf.R >= 3 {
		_gg, _aeb = _ed.alg5(_befe, _fbb)
	} else {
		return nil, _a.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R")
	}
	if _aeb != nil {
		return nil, _aeb
	}
	_ab.Log.Trace("\u0063\u0068\u0065\u0063k:\u0020\u0025\u0020\u0078\u0020\u003d\u003d\u0020\u0025\u0020\u0078\u0020\u003f", string(_gg), string(_edf.U))
	_acc := _gg
	_acg := _edf.U
	if _edf.R >= 3 {
		if len(_acc) > 16 {
			_acc = _acc[0:16]
		}
		if len(_acg) > 16 {
			_acg = _acg[0:16]
		}
	}
	if !_dgd.Equal(_acc, _acg) {
		return nil, nil
	}
	return _befe, nil
}

const _age = "\x28\277\116\136\x4e\x75\x8a\x41\x64\000\x4e\x56\377" + "\xfa\001\010\056\x2e\x00\xb6\xd0\x68\076\x80\x2f\014" + "\251\xfe\x64\x53\x69\172"

func (_aag stdHandlerR4) alg4(_ec []byte, _eb []byte) ([]byte, error) {
	_gac, _fc := _fg.NewCipher(_ec)
	if _fc != nil {
		return nil, _a.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_gaa := []byte(_age)
	_ad := make([]byte, len(_gaa))
	_gac.XORKeyStream(_ad, _gaa)
	return _ad, nil
}

type errInvalidField struct {
	Func  string
	Field string
	Exp   int
	Got   int
}

// NewHandlerR6 creates a new standard security handler for R=5 and R=6.
func NewHandlerR6() StdHandler { return stdHandlerR6{} }
func _dfc(_ga _b.Block) *ecb   { return &ecb{_df: _ga, _bf: _ga.BlockSize()} }

type stdHandlerR4 struct {
	Length int
	ID0    string
}

func (_bda stdHandlerR4) alg3(R int, _ffg, _cdf []byte) ([]byte, error) {
	var _acb []byte
	if len(_cdf) > 0 {
		_acb = _bda.alg3Key(R, _cdf)
	} else {
		_acb = _bda.alg3Key(R, _ffg)
	}
	_cf, _def := _fg.NewCipher(_acb)
	if _def != nil {
		return nil, _a.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_cgg := _bda.paddedPass(_ffg)
	_bde := make([]byte, len(_cgg))
	_cf.XORKeyStream(_bde, _cgg)
	if R >= 3 {
		_afa := make([]byte, len(_acb))
		for _ffc := 0; _ffc < 19; _ffc++ {
			for _deg := 0; _deg < len(_acb); _deg++ {
				_afa[_deg] = _acb[_deg] ^ byte(_ffc+1)
			}
			_aa, _dgf := _fg.NewCipher(_afa)
			if _dgf != nil {
				return nil, _a.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
			}
			_aa.XORKeyStream(_bde, _bde)
		}
	}
	return _bde, nil
}

func _bdd(_cec []byte) (_b.Block, error) {
	_dcb, _eccf := _dg.NewCipher(_cec)
	if _eccf != nil {
		_ab.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0063\u0069p\u0068\u0065r\u003a\u0020\u0025\u0076", _eccf)
		return nil, _eccf
	}
	return _dcb, nil
}

// GenerateParams is the algorithm opposite to alg2a (R>=5).
// It generates U,O,UE,OE,Perms fields using AESv3 encryption.
// There is no algorithm number assigned to this function in the spec.
// It expects R, P and EncryptMetadata fields to be set.
func (_abb stdHandlerR6) GenerateParams(d *StdEncryptDict, opass, upass []byte) ([]byte, error) {
	_edda := make([]byte, 32)
	if _, _bdf := _d.ReadFull(_gd.Reader, _edda); _bdf != nil {
		return nil, _bdf
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
	if _acdb := _abb.alg8(d, _edda, upass); _acdb != nil {
		return nil, _acdb
	}
	if _cgfb := _abb.alg9(d, _edda, opass); _cgfb != nil {
		return nil, _cgfb
	}
	if d.R == 5 {
		return _edda, nil
	}
	if _fda := _abb.alg10(d, _edda); _fda != nil {
		return nil, _fda
	}
	return _edda, nil
}

// AuthEvent is an event type that triggers authentication.
type AuthEvent string

var _ StdHandler = stdHandlerR4{}

func (_fcc stdHandlerR4) alg5(_da []byte, _eaed []byte) ([]byte, error) {
	_gdc := _gc.New()
	_gdc.Write([]byte(_age))
	_gdc.Write([]byte(_fcc.ID0))
	_abc := _gdc.Sum(nil)
	_ab.Log.Trace("\u0061\u006c\u0067\u0035")
	_ab.Log.Trace("\u0065k\u0065\u0079\u003a\u0020\u0025\u0020x", _da)
	_ab.Log.Trace("\u0049D\u003a\u0020\u0025\u0020\u0078", _fcc.ID0)
	if len(_abc) != 16 {
		return nil, _a.New("\u0068a\u0073\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074\u0020\u0031\u0036\u0020\u0062\u0079\u0074\u0065\u0073")
	}
	_bef, _gdb := _fg.NewCipher(_da)
	if _gdb != nil {
		return nil, _a.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_fccd := make([]byte, 16)
	_bef.XORKeyStream(_fccd, _abc)
	_db := make([]byte, len(_da))
	for _ega := 0; _ega < 19; _ega++ {
		for _gf := 0; _gf < len(_da); _gf++ {
			_db[_gf] = _da[_gf] ^ byte(_ega+1)
		}
		_bef, _gdb = _fg.NewCipher(_db)
		if _gdb != nil {
			return nil, _a.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
		}
		_bef.XORKeyStream(_fccd, _fccd)
		_ab.Log.Trace("\u0069\u0020\u003d\u0020\u0025\u0064\u002c\u0020\u0065\u006b\u0065\u0079:\u0020\u0025\u0020\u0078", _ega, _db)
		_ab.Log.Trace("\u0069\u0020\u003d\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0020\u0078", _ega, _fccd)
	}
	_dge := make([]byte, 32)
	for _ffa := 0; _ffa < 16; _ffa++ {
		_dge[_ffa] = _fccd[_ffa]
	}
	_, _gdb = _gd.Read(_dge[16:32])
	if _gdb != nil {
		return nil, _a.New("\u0066a\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065n\u0020r\u0061n\u0064\u0020\u006e\u0075\u006d\u0062\u0065r")
	}
	return _dge, nil
}

const (
	EventDocOpen = AuthEvent("\u0044o\u0063\u004f\u0070\u0065\u006e")
	EventEFOpen  = AuthEvent("\u0045\u0046\u004f\u0070\u0065\u006e")
)

var _ StdHandler = stdHandlerR6{}

func (_eaef stdHandlerR6) alg2a(_ecca *StdEncryptDict, _fgcc []byte) ([]byte, Permissions, error) {
	if _aeda := _faa("\u0061\u006c\u00672\u0061", "\u004f", 48, _ecca.O); _aeda != nil {
		return nil, 0, _aeda
	}
	if _gdca := _faa("\u0061\u006c\u00672\u0061", "\u0055", 48, _ecca.U); _gdca != nil {
		return nil, 0, _gdca
	}
	if len(_fgcc) > 127 {
		_fgcc = _fgcc[:127]
	}
	_gga, _bad := _eaef.alg12(_ecca, _fgcc)
	if _bad != nil {
		return nil, 0, _bad
	}
	var (
		_fgga []byte
		_fgf  []byte
		_fbc  []byte
	)
	var _ccd Permissions
	if len(_gga) != 0 {
		_ccd = PermOwner
		_caf := make([]byte, len(_fgcc)+8+48)
		_bc := copy(_caf, _fgcc)
		_bc += copy(_caf[_bc:], _ecca.O[40:48])
		copy(_caf[_bc:], _ecca.U[0:48])
		_fgga = _caf
		_fgf = _ecca.OE
		_fbc = _ecca.U[0:48]
	} else {
		_gga, _bad = _eaef.alg11(_ecca, _fgcc)
		if _bad == nil && len(_gga) == 0 {
			_gga, _bad = _eaef.alg11(_ecca, []byte(""))
		}
		if _bad != nil {
			return nil, 0, _bad
		} else if len(_gga) == 0 {
			return nil, 0, nil
		}
		_ccd = _ecca.P
		_gaab := make([]byte, len(_fgcc)+8)
		_bcg := copy(_gaab, _fgcc)
		copy(_gaab[_bcg:], _ecca.U[40:48])
		_fgga = _gaab
		_fgf = _ecca.UE
		_fbc = nil
	}
	if _ced := _faa("\u0061\u006c\u00672\u0061", "\u004b\u0065\u0079", 32, _fgf); _ced != nil {
		return nil, 0, _ced
	}
	_fgf = _fgf[:32]
	_fd, _bad := _eaef.alg2b(_ecca.R, _fgga, _fgcc, _fbc)
	if _bad != nil {
		return nil, 0, _bad
	}
	_gfe, _bad := _dg.NewCipher(_fd[:32])
	if _bad != nil {
		return nil, 0, _bad
	}
	_fe := make([]byte, _dg.BlockSize)
	_dcad := _b.NewCBCDecrypter(_gfe, _fe)
	_bcb := make([]byte, 32)
	_dcad.CryptBlocks(_bcb, _fgf)
	if _ecca.R == 5 {
		return _bcb, _ccd, nil
	}
	_bad = _eaef.alg13(_ecca, _bcb)
	if _bad != nil {
		return nil, 0, _bad
	}
	return _bcb, _ccd, nil
}
func (_ba *ecbDecrypter) BlockSize() int { return _ba._bf }
func (_bfb stdHandlerR6) alg9(_dedce *StdEncryptDict, _gcf []byte, _cffb []byte) error {
	if _cdg := _faa("\u0061\u006c\u0067\u0039", "\u004b\u0065\u0079", 32, _gcf); _cdg != nil {
		return _cdg
	}
	if _dgfc := _faa("\u0061\u006c\u0067\u0039", "\u0055", 48, _dedce.U); _dgfc != nil {
		return _dgfc
	}
	var _bddf [16]byte
	if _, _fcca := _d.ReadFull(_gd.Reader, _bddf[:]); _fcca != nil {
		return _fcca
	}
	_acfe := _bddf[0:8]
	_faf := _bddf[8:16]
	_gab := _dedce.U[:48]
	_eebe := make([]byte, len(_cffb)+len(_acfe)+len(_gab))
	_cbfa := copy(_eebe, _cffb)
	_cbfa += copy(_eebe[_cbfa:], _acfe)
	_cbfa += copy(_eebe[_cbfa:], _gab)
	_fag, _fdfc := _bfb.alg2b(_dedce.R, _eebe, _cffb, _gab)
	if _fdfc != nil {
		return _fdfc
	}
	O := make([]byte, len(_fag)+len(_acfe)+len(_faf))
	_cbfa = copy(O, _fag[:32])
	_cbfa += copy(O[_cbfa:], _acfe)
	_cbfa += copy(O[_cbfa:], _faf)
	_dedce.O = O
	_cbfa = len(_cffb)
	_cbfa += copy(_eebe[_cbfa:], _faf)
	_fag, _fdfc = _bfb.alg2b(_dedce.R, _eebe, _cffb, _gab)
	if _fdfc != nil {
		return _fdfc
	}
	_bbg, _fdfc := _bdd(_fag[:32])
	if _fdfc != nil {
		return _fdfc
	}
	_dega := make([]byte, _dg.BlockSize)
	_adg := _b.NewCBCEncrypter(_bbg, _dega)
	OE := make([]byte, 32)
	_adg.CryptBlocks(OE, _gcf[:32])
	_dedce.OE = OE
	return nil
}

// NewHandlerR4 creates a new standard security handler for R<=4.
func NewHandlerR4(id0 string, length int) StdHandler { return stdHandlerR4{ID0: id0, Length: length} }

func _faa(_eg, _ef string, _aea int, _cg []byte) error {
	if len(_cg) < _aea {
		return errInvalidField{Func: _eg, Field: _ef, Exp: _aea, Got: len(_cg)}
	}
	return nil
}

func (_fbd stdHandlerR4) alg7(_abcc *StdEncryptDict, _cbd []byte) ([]byte, error) {
	_egb := _fbd.alg3Key(_abcc.R, _cbd)
	_bb := make([]byte, len(_abcc.O))
	if _abcc.R == 2 {
		_ca, _aba := _fg.NewCipher(_egb)
		if _aba != nil {
			return nil, _a.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072")
		}
		_ca.XORKeyStream(_bb, _abcc.O)
	} else if _abcc.R >= 3 {
		_ce := append([]byte{}, _abcc.O...)
		for _dfd := 0; _dfd < 20; _dfd++ {
			_ddd := append([]byte{}, _egb...)
			for _ecc := 0; _ecc < len(_egb); _ecc++ {
				_ddd[_ecc] ^= byte(19 - _dfd)
			}
			_abd, _ebg := _fg.NewCipher(_ddd)
			if _ebg != nil {
				return nil, _a.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072")
			}
			_abd.XORKeyStream(_bb, _ce)
			_ce = append([]byte{}, _bb...)
		}
	} else {
		return nil, _a.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R")
	}
	_daa, _dba := _fbd.alg6(_abcc, _bb)
	if _dba != nil {
		return nil, nil
	}
	return _daa, nil
}

type stdHandlerR6 struct{}

func (_ag *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%_ag._bf != 0 {
		_ab.Log.Error("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073")
		return
	}
	if len(dst) < len(src) {
		_ab.Log.Error("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0064\u0065\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074")
		return
	}
	for len(src) > 0 {
		_ag._df.Decrypt(dst, src[:_ag._bf])
		src = src[_ag._bf:]
		dst = dst[_ag._bf:]
	}
}

func (_gdf errInvalidField) Error() string {
	return _e.Sprintf("\u0025s\u003a\u0020e\u0078\u0070\u0065\u0063t\u0065\u0064\u0020%\u0073\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074o \u0062\u0065\u0020%\u0064\u0020b\u0079\u0074\u0065\u0073\u002c\u0020g\u006f\u0074 \u0025\u0064", _gdf.Func, _gdf.Field, _gdf.Exp, _gdf.Got)
}

func (stdHandlerR4) paddedPass(_ac []byte) []byte {
	_dgda := make([]byte, 32)
	_cb := copy(_dgda, _ac)
	for ; _cb < 32; _cb++ {
		_dgda[_cb] = _age[_cb-len(_ac)]
	}
	return _dgda
}

func (_bdb *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%_bdb._bf != 0 {
		_ab.Log.Error("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073")
		return
	}
	if len(dst) < len(src) {
		_ab.Log.Error("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074")
		return
	}
	for len(src) > 0 {
		_bdb._df.Encrypt(dst, src[:_bdb._bf])
		src = src[_bdb._bf:]
		dst = dst[_bdb._bf:]
	}
}

func (_bbc stdHandlerR6) alg10(_dgb *StdEncryptDict, _cba []byte) error {
	if _baea := _faa("\u0061\u006c\u00671\u0030", "\u004b\u0065\u0079", 32, _cba); _baea != nil {
		return _baea
	}
	_cgf := uint64(uint32(_dgb.P)) | (_fb.MaxUint32 << 32)
	Perms := make([]byte, 16)
	_bd.LittleEndian.PutUint64(Perms[:8], _cgf)
	if _dgb.EncryptMetadata {
		Perms[8] = 'T'
	} else {
		Perms[8] = 'F'
	}
	copy(Perms[9:12], "\u0061\u0064\u0062")
	if _, _fdc := _d.ReadFull(_gd.Reader, Perms[12:16]); _fdc != nil {
		return _fdc
	}
	_cdfe, _aagb := _bdd(_cba[:32])
	if _aagb != nil {
		return _aagb
	}
	_acd := _ae(_cdfe)
	_acd.CryptBlocks(Perms, Perms)
	_dgb.Perms = Perms[:16]
	return nil
}

func (_cga stdHandlerR6) alg8(_cedb *StdEncryptDict, _gbf []byte, _ecd []byte) error {
	if _cecf := _faa("\u0061\u006c\u0067\u0038", "\u004b\u0065\u0079", 32, _gbf); _cecf != nil {
		return _cecf
	}
	var _ecdf [16]byte
	if _, _gaabf := _d.ReadFull(_gd.Reader, _ecdf[:]); _gaabf != nil {
		return _gaabf
	}
	_dad := _ecdf[0:8]
	_afd := _ecdf[8:16]
	_eeb := make([]byte, len(_ecd)+len(_dad))
	_aae := copy(_eeb, _ecd)
	copy(_eeb[_aae:], _dad)
	_gcb, _abg := _cga.alg2b(_cedb.R, _eeb, _ecd, nil)
	if _abg != nil {
		return _abg
	}
	U := make([]byte, len(_gcb)+len(_dad)+len(_afd))
	_aae = copy(U, _gcb[:32])
	_aae += copy(U[_aae:], _dad)
	copy(U[_aae:], _afd)
	_cedb.U = U
	_aae = len(_ecd)
	copy(_eeb[_aae:], _afd)
	_gcb, _abg = _cga.alg2b(_cedb.R, _eeb, _ecd, nil)
	if _abg != nil {
		return _abg
	}
	_aaf, _abg := _bdd(_gcb[:32])
	if _abg != nil {
		return _abg
	}
	_cff := make([]byte, _dg.BlockSize)
	_adf := _b.NewCBCEncrypter(_aaf, _cff)
	UE := make([]byte, 32)
	_adf.CryptBlocks(UE, _gbf[:32])
	_cedb.UE = UE
	return nil
}

func (_de stdHandlerR4) alg3Key(R int, _ge []byte) []byte {
	_bdba := _gc.New()
	_eae := _de.paddedPass(_ge)
	_bdba.Write(_eae)
	if R >= 3 {
		for _aeg := 0; _aeg < 50; _aeg++ {
			_af := _bdba.Sum(nil)
			_bdba = _gc.New()
			_bdba.Write(_af)
		}
	}
	_dca := _bdba.Sum(nil)
	if R == 2 {
		_dca = _dca[0:5]
	} else {
		_dca = _dca[0 : _de.Length/8]
	}
	return _dca
}

func (_bdec stdHandlerR6) alg12(_acac *StdEncryptDict, _bgd []byte) ([]byte, error) {
	if _fde := _faa("\u0061\u006c\u00671\u0032", "\u0055", 48, _acac.U); _fde != nil {
		return nil, _fde
	}
	if _dce := _faa("\u0061\u006c\u00671\u0032", "\u004f", 48, _acac.O); _dce != nil {
		return nil, _dce
	}
	_fab := make([]byte, len(_bgd)+8+48)
	_ece := copy(_fab, _bgd)
	_ece += copy(_fab[_ece:], _acac.O[32:40])
	_ece += copy(_fab[_ece:], _acac.U[0:48])
	_gcbc, _gfg := _bdec.alg2b(_acac.R, _fab, _bgd, _acac.U[0:48])
	if _gfg != nil {
		return nil, _gfg
	}
	_gcbc = _gcbc[:32]
	if !_dgd.Equal(_gcbc, _acac.O[:32]) {
		return nil, nil
	}
	return _gcbc, nil
}
func _gb(_cd _b.Block) _b.BlockMode { return (*ecbDecrypter)(_dfc(_cd)) }

// Authenticate implements StdHandler interface.
func (_efe stdHandlerR6) Authenticate(d *StdEncryptDict, pass []byte) ([]byte, Permissions, error) {
	return _efe.alg2a(d, pass)
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
type ecbEncrypter ecb

// Authenticate implements StdHandler interface.
func (_cfg stdHandlerR4) Authenticate(d *StdEncryptDict, pass []byte) ([]byte, Permissions, error) {
	_ab.Log.Trace("\u0044\u0065b\u0075\u0067\u0067\u0069n\u0067\u0020a\u0075\u0074\u0068\u0065\u006e\u0074\u0069\u0063a\u0074\u0069\u006f\u006e\u0020\u002d\u0020\u006f\u0077\u006e\u0065\u0072 \u0070\u0061\u0073\u0073")
	_dea, _dgeg := _cfg.alg7(d, pass)
	if _dgeg != nil {
		return nil, 0, _dgeg
	}
	if _dea != nil {
		_ab.Log.Trace("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e")
		return _dea, PermOwner, nil
	}
	_ab.Log.Trace("\u0044\u0065bu\u0067\u0067\u0069n\u0067\u0020\u0061\u0075the\u006eti\u0063\u0061\u0074\u0069\u006f\u006e\u0020- \u0075\u0073\u0065\u0072\u0020\u0070\u0061s\u0073")
	_dea, _dgeg = _cfg.alg6(d, pass)
	if _dgeg != nil {
		return nil, 0, _dgeg
	}
	if _dea != nil {
		_ab.Log.Trace("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e")
		return _dea, d.P, nil
	}
	return nil, 0, nil
}
