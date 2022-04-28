package security

import (
	_fe "bytes"
	_a "crypto/aes"
	_f "crypto/cipher"
	_ed "crypto/md5"
	_c "crypto/rand"
	_g "crypto/rc4"
	_ec "crypto/sha256"
	_d "crypto/sha512"
	_da "encoding/binary"
	_ad "errors"
	_gb "fmt"
	_bc "hash"
	_e "io"
	_fb "math"

	_cf "bitbucket.org/shenghui0779/gopdf/common"
)

// GenerateParams is the algorithm opposite to alg2a (R>=5).
// It generates U,O,UE,OE,Perms fields using AESv3 encryption.
// There is no algorithm number assigned to this function in the spec.
// It expects R, P and EncryptMetadata fields to be set.
func (_bca stdHandlerR6) GenerateParams(d *StdEncryptDict, opass, upass []byte) ([]byte, error) {
	_bcf := make([]byte, 32)
	if _, _eee := _e.ReadFull(_c.Reader, _bcf); _eee != nil {
		return nil, _eee
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
	if _daeb := _bca.alg8(d, _bcf, upass); _daeb != nil {
		return nil, _daeb
	}
	if _abdb := _bca.alg9(d, _bcf, opass); _abdb != nil {
		return nil, _abdb
	}
	if d.R == 5 {
		return _bcf, nil
	}
	if _fcd := _bca.alg10(d, _bcf); _fcd != nil {
		return nil, _fcd
	}
	return _bcf, nil
}
func (_gfe stdHandlerR6) alg2a(_eafc *StdEncryptDict, _fec []byte) ([]byte, Permissions, error) {
	if _afc := _gba("\u0061\u006c\u00672\u0061", "\u004f", 48, _eafc.O); _afc != nil {
		return nil, 0, _afc
	}
	if _gdc := _gba("\u0061\u006c\u00672\u0061", "\u0055", 48, _eafc.U); _gdc != nil {
		return nil, 0, _gdc
	}
	if len(_fec) > 127 {
		_fec = _fec[:127]
	}
	_eafgg, _eb := _gfe.alg12(_eafc, _fec)
	if _eb != nil {
		return nil, 0, _eb
	}
	var (
		_ddd  []byte
		_gcf  []byte
		_geba []byte
	)
	var _bgf Permissions
	if len(_eafgg) != 0 {
		_bgf = PermOwner
		_ecf := make([]byte, len(_fec)+8+48)
		_gge := copy(_ecf, _fec)
		_gge += copy(_ecf[_gge:], _eafc.O[40:48])
		copy(_ecf[_gge:], _eafc.U[0:48])
		_ddd = _ecf
		_gcf = _eafc.OE
		_geba = _eafc.U[0:48]
	} else {
		_eafgg, _eb = _gfe.alg11(_eafc, _fec)
		if _eb == nil && len(_eafgg) == 0 {
			_eafgg, _eb = _gfe.alg11(_eafc, []byte(""))
		}
		if _eb != nil {
			return nil, 0, _eb
		} else if len(_eafgg) == 0 {
			return nil, 0, nil
		}
		_bgf = _eafc.P
		_cdg := make([]byte, len(_fec)+8)
		_cba := copy(_cdg, _fec)
		copy(_cdg[_cba:], _eafc.U[40:48])
		_ddd = _cdg
		_gcf = _eafc.UE
		_geba = nil
	}
	if _caae := _gba("\u0061\u006c\u00672\u0061", "\u004b\u0065\u0079", 32, _gcf); _caae != nil {
		return nil, 0, _caae
	}
	_gcf = _gcf[:32]
	_fc, _eb := _gfe.alg2b(_eafc.R, _ddd, _fec, _geba)
	if _eb != nil {
		return nil, 0, _eb
	}
	_dabg, _eb := _a.NewCipher(_fc[:32])
	if _eb != nil {
		return nil, 0, _eb
	}
	_ee := make([]byte, _a.BlockSize)
	_cgb := _f.NewCBCDecrypter(_dabg, _ee)
	_gbf := make([]byte, 32)
	_cgb.CryptBlocks(_gbf, _gcf)
	if _eafc.R == 5 {
		return _gbf, _bgf, nil
	}
	_eb = _gfe.alg13(_eafc, _gbf)
	if _eb != nil {
		return nil, 0, _eb
	}
	return _gbf, _bgf, nil
}
func _ccb(_gaf []byte, _caaa int) {
	_bgaf := _caaa
	for _bgaf < len(_gaf) {
		copy(_gaf[_bgaf:], _gaf[:_bgaf])
		_bgaf *= 2
	}
}

type errInvalidField struct {
	Func  string
	Field string
	Exp   int
	Got   int
}
type stdHandlerR6 struct{}

// StdEncryptDict is a set of additional fields used in standard encryption dictionary.
type StdEncryptDict struct {
	R               int
	P               Permissions
	EncryptMetadata bool
	O, U            []byte
	OE, UE          []byte
	Perms           []byte
}

func (_ggb stdHandlerR6) alg10(_egc *StdEncryptDict, _dae []byte) error {
	if _ede := _gba("\u0061\u006c\u00671\u0030", "\u004b\u0065\u0079", 32, _dae); _ede != nil {
		return _ede
	}
	_bgbc := uint64(uint32(_egc.P)) | (_fb.MaxUint32 << 32)
	Perms := make([]byte, 16)
	_da.LittleEndian.PutUint64(Perms[:8], _bgbc)
	if _egc.EncryptMetadata {
		Perms[8] = 'T'
	} else {
		Perms[8] = 'F'
	}
	copy(Perms[9:12], "\u0061\u0064\u0062")
	if _, _cff := _e.ReadFull(_c.Reader, Perms[12:16]); _cff != nil {
		return _cff
	}
	_gga, _fdd := _abc(_dae[:32])
	if _fdd != nil {
		return _fdd
	}
	_eaba := _caa(_gga)
	_eaba.CryptBlocks(Perms, Perms)
	_egc.Perms = Perms[:16]
	return nil
}
func (_gcc stdHandlerR6) alg8(_edf *StdEncryptDict, _efe []byte, _edg []byte) error {
	if _egef := _gba("\u0061\u006c\u0067\u0038", "\u004b\u0065\u0079", 32, _efe); _egef != nil {
		return _egef
	}
	var _dcg [16]byte
	if _, _acc := _e.ReadFull(_c.Reader, _dcg[:]); _acc != nil {
		return _acc
	}
	_bgc := _dcg[0:8]
	_ebd := _dcg[8:16]
	_eabc := make([]byte, len(_edg)+len(_bgc))
	_fce := copy(_eabc, _edg)
	copy(_eabc[_fce:], _bgc)
	_ffc, _bfab := _gcc.alg2b(_edf.R, _eabc, _edg, nil)
	if _bfab != nil {
		return _bfab
	}
	U := make([]byte, len(_ffc)+len(_bgc)+len(_ebd))
	_fce = copy(U, _ffc[:32])
	_fce += copy(U[_fce:], _bgc)
	copy(U[_fce:], _ebd)
	_edf.U = U
	_fce = len(_edg)
	copy(_eabc[_fce:], _ebd)
	_ffc, _bfab = _gcc.alg2b(_edf.R, _eabc, _edg, nil)
	if _bfab != nil {
		return _bfab
	}
	_deb, _bfab := _abc(_ffc[:32])
	if _bfab != nil {
		return _bfab
	}
	_badg := make([]byte, _a.BlockSize)
	_dgf := _f.NewCBCEncrypter(_deb, _badg)
	UE := make([]byte, 32)
	_dgf.CryptBlocks(UE, _efe[:32])
	_edf.UE = UE
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

type ecb struct {
	_ff _f.Block
	_ca int
}

func (stdHandlerR4) paddedPass(_ae []byte) []byte {
	_aeb := make([]byte, 32)
	_gc := copy(_aeb, _ae)
	for ; _gc < 32; _gc++ {
		_aeb[_gc] = _cbd[_gc-len(_ae)]
	}
	return _aeb
}
func _daf(_gf _f.Block) *ecb { return &ecb{_ff: _gf, _ca: _gf.BlockSize()} }

const _cbd = "\x28\277\116\136\x4e\x75\x8a\x41\x64\000\x4e\x56\377" + "\xfa\001\010\056\x2e\x00\xb6\xd0\x68\076\x80\x2f\014" + "\251\xfe\x64\x53\x69\172"

func (_feg *ecbEncrypter) BlockSize() int { return _feg._ca }

// GenerateParams generates and sets O and U parameters for the encryption dictionary.
// It expects R, P and EncryptMetadata fields to be set.
func (_bdd stdHandlerR4) GenerateParams(d *StdEncryptDict, opass, upass []byte) ([]byte, error) {
	O, _ceb := _bdd.alg3(d.R, upass, opass)
	if _ceb != nil {
		_cf.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", _ceb)
		return nil, _ceb
	}
	d.O = O
	_cf.Log.Trace("\u0067\u0065\u006e\u0020\u004f\u003a\u0020\u0025\u0020\u0078", O)
	_cad := _bdd.alg2(d, upass)
	U, _ceb := _bdd.alg5(_cad, upass)
	if _ceb != nil {
		_cf.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", _ceb)
		return nil, _ceb
	}
	d.U = U
	_cf.Log.Trace("\u0067\u0065\u006e\u0020\u0055\u003a\u0020\u0025\u0020\u0078", U)
	return _cad, nil
}

// StdHandler is an interface for standard security handlers.
type StdHandler interface {

	// GenerateParams uses owner and user passwords to set encryption parameters and generate an encryption key.
	// It assumes that R, P and EncryptMetadata are already set.
	GenerateParams(_bd *StdEncryptDict, _ffb, _ac []byte) ([]byte, error)

	// Authenticate uses encryption dictionary parameters and the password to calculate
	// the document encryption key. It also returns permissions that should be granted to a user.
	// In case of failed authentication, it returns empty key and zero permissions with no error.
	Authenticate(_abe *StdEncryptDict, _ea []byte) ([]byte, Permissions, error)
}

func (_ba *ecbDecrypter) BlockSize() int { return _ba._ca }
func (_gg *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%_gg._ca != 0 {
		_cf.Log.Error("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073")
		return
	}
	if len(dst) < len(src) {
		_cf.Log.Error("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074")
		return
	}
	for len(src) > 0 {
		_gg._ff.Encrypt(dst, src[:_gg._ca])
		src = src[_gg._ca:]
		dst = dst[_gg._ca:]
	}
}
func (_cfa stdHandlerR4) alg4(_dafd []byte, _ceg []byte) ([]byte, error) {
	_fa, _caf := _g.NewCipher(_dafd)
	if _caf != nil {
		return nil, _ad.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_dba := []byte(_cbd)
	_add := make([]byte, len(_dba))
	_fa.XORKeyStream(_add, _dba)
	return _add, nil
}
func _gfg(_dace []byte) ([]byte, error) {
	_aac := _ec.New()
	_aac.Write(_dace)
	return _aac.Sum(nil), nil
}
func (_bf stdHandlerR4) alg6(_bfc *StdEncryptDict, _gcd []byte) ([]byte, error) {
	var (
		_ace []byte
		_egb error
	)
	_ggg := _bf.alg2(_bfc, _gcd)
	if _bfc.R == 2 {
		_ace, _egb = _bf.alg4(_ggg, _gcd)
	} else if _bfc.R >= 3 {
		_ace, _egb = _bf.alg5(_ggg, _gcd)
	} else {
		return nil, _ad.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R")
	}
	if _egb != nil {
		return nil, _egb
	}
	_cf.Log.Trace("\u0063\u0068\u0065\u0063k:\u0020\u0025\u0020\u0078\u0020\u003d\u003d\u0020\u0025\u0020\u0078\u0020\u003f", string(_ace), string(_bfc.U))
	_dcb := _ace
	_bfe := _bfc.U
	if _bfc.R >= 3 {
		if len(_dcb) > 16 {
			_dcb = _dcb[0:16]
		}
		if len(_bfe) > 16 {
			_bfe = _bfe[0:16]
		}
	}
	if !_fe.Equal(_dcb, _bfe) {
		return nil, nil
	}
	return _ggg, nil
}

// Permissions is a bitmask of access permissions for a PDF file.
type Permissions uint32

func (_gfb stdHandlerR6) alg9(_agg *StdEncryptDict, _gccb []byte, _cfg []byte) error {
	if _ffd := _gba("\u0061\u006c\u0067\u0039", "\u004b\u0065\u0079", 32, _gccb); _ffd != nil {
		return _ffd
	}
	if _gccbe := _gba("\u0061\u006c\u0067\u0039", "\u0055", 48, _agg.U); _gccbe != nil {
		return _gccbe
	}
	var _bbe [16]byte
	if _, _aeg := _e.ReadFull(_c.Reader, _bbe[:]); _aeg != nil {
		return _aeg
	}
	_dage := _bbe[0:8]
	_fed := _bbe[8:16]
	_abgd := _agg.U[:48]
	_aebg := make([]byte, len(_cfg)+len(_dage)+len(_abgd))
	_gee := copy(_aebg, _cfg)
	_gee += copy(_aebg[_gee:], _dage)
	_gee += copy(_aebg[_gee:], _abgd)
	_acg, _cbbd := _gfb.alg2b(_agg.R, _aebg, _cfg, _abgd)
	if _cbbd != nil {
		return _cbbd
	}
	O := make([]byte, len(_acg)+len(_dage)+len(_fed))
	_gee = copy(O, _acg[:32])
	_gee += copy(O[_gee:], _dage)
	_gee += copy(O[_gee:], _fed)
	_agg.O = O
	_gee = len(_cfg)
	_gee += copy(_aebg[_gee:], _fed)
	_acg, _cbbd = _gfb.alg2b(_agg.R, _aebg, _cfg, _abgd)
	if _cbbd != nil {
		return _cbbd
	}
	_fgc, _cbbd := _abc(_acg[:32])
	if _cbbd != nil {
		return _cbbd
	}
	_fbe := make([]byte, _a.BlockSize)
	_gaeb := _f.NewCBCEncrypter(_fgc, _fbe)
	OE := make([]byte, 32)
	_gaeb.CryptBlocks(OE, _gccb[:32])
	_agg.OE = OE
	return nil
}

// NewHandlerR6 creates a new standard security handler for R=5 and R=6.
func NewHandlerR6() StdHandler { return stdHandlerR6{} }

var _ StdHandler = stdHandlerR6{}

func (_de stdHandlerR4) alg3Key(R int, _dbb []byte) []byte {
	_eaf := _ed.New()
	_ce := _de.paddedPass(_dbb)
	_eaf.Write(_ce)
	if R >= 3 {
		for _gd := 0; _gd < 50; _gd++ {
			_ege := _eaf.Sum(nil)
			_eaf = _ed.New()
			_eaf.Write(_ege)
		}
	}
	_gdf := _eaf.Sum(nil)
	if R == 2 {
		_gdf = _gdf[0:5]
	} else {
		_gdf = _gdf[0 : _de.Length/8]
	}
	return _gdf
}
func (_geb stdHandlerR4) alg7(_bbd *StdEncryptDict, _cec []byte) ([]byte, error) {
	_bgbd := _geb.alg3Key(_bbd.R, _cec)
	_agf := make([]byte, len(_bbd.O))
	if _bbd.R == 2 {
		_bfcf, _af := _g.NewCipher(_bgbd)
		if _af != nil {
			return nil, _ad.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072")
		}
		_bfcf.XORKeyStream(_agf, _bbd.O)
	} else if _bbd.R >= 3 {
		_ced := append([]byte{}, _bbd.O...)
		for _bge := 0; _bge < 20; _bge++ {
			_bba := append([]byte{}, _bgbd...)
			for _efg := 0; _efg < len(_bgbd); _efg++ {
				_bba[_efg] ^= byte(19 - _bge)
			}
			_fdac, _bga := _g.NewCipher(_bba)
			if _bga != nil {
				return nil, _ad.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072")
			}
			_fdac.XORKeyStream(_agf, _ced)
			_ced = append([]byte{}, _agf...)
		}
	} else {
		return nil, _ad.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R")
	}
	_gaa, _bfa := _geb.alg6(_bbd, _agf)
	if _bfa != nil {
		return nil, nil
	}
	return _gaa, nil
}
func _gba(_cd, _dd string, _ge int, _bg []byte) error {
	if len(_bg) < _ge {
		return errInvalidField{Func: _cd, Field: _dd, Exp: _ge, Got: len(_bg)}
	}
	return nil
}
func (_cc stdHandlerR4) alg5(_dg []byte, _gbb []byte) ([]byte, error) {
	_abg := _ed.New()
	_abg.Write([]byte(_cbd))
	_abg.Write([]byte(_cc.ID0))
	_eaa := _abg.Sum(nil)
	_cf.Log.Trace("\u0061\u006c\u0067\u0035")
	_cf.Log.Trace("\u0065k\u0065\u0079\u003a\u0020\u0025\u0020x", _dg)
	_cf.Log.Trace("\u0049D\u003a\u0020\u0025\u0020\u0078", _cc.ID0)
	if len(_eaa) != 16 {
		return nil, _ad.New("\u0068a\u0073\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074\u0020\u0031\u0036\u0020\u0062\u0079\u0074\u0065\u0073")
	}
	_gae, _gaeg := _g.NewCipher(_dg)
	if _gaeg != nil {
		return nil, _ad.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_aa := make([]byte, 16)
	_gae.XORKeyStream(_aa, _eaa)
	_ef := make([]byte, len(_dg))
	for _fd := 0; _fd < 19; _fd++ {
		for _agc := 0; _agc < len(_dg); _agc++ {
			_ef[_agc] = _dg[_agc] ^ byte(_fd+1)
		}
		_gae, _gaeg = _g.NewCipher(_ef)
		if _gaeg != nil {
			return nil, _ad.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
		}
		_gae.XORKeyStream(_aa, _aa)
		_cf.Log.Trace("\u0069\u0020\u003d\u0020\u0025\u0064\u002c\u0020\u0065\u006b\u0065\u0079:\u0020\u0025\u0020\u0078", _fd, _ef)
		_cf.Log.Trace("\u0069\u0020\u003d\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0020\u0078", _fd, _aa)
	}
	_fda := make([]byte, 32)
	for _bb := 0; _bb < 16; _bb++ {
		_fda[_bb] = _aa[_bb]
	}
	_, _gaeg = _c.Read(_fda[16:32])
	if _gaeg != nil {
		return nil, _ad.New("\u0066a\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065n\u0020r\u0061n\u0064\u0020\u006e\u0075\u006d\u0062\u0065r")
	}
	return _fda, nil
}

// Authenticate implements StdHandler interface.
func (_bed stdHandlerR4) Authenticate(d *StdEncryptDict, pass []byte) ([]byte, Permissions, error) {
	_cf.Log.Trace("\u0044\u0065b\u0075\u0067\u0067\u0069n\u0067\u0020a\u0075\u0074\u0068\u0065\u006e\u0074\u0069\u0063a\u0074\u0069\u006f\u006e\u0020\u002d\u0020\u006f\u0077\u006e\u0065\u0072 \u0070\u0061\u0073\u0073")
	_afg, _bda := _bed.alg7(d, pass)
	if _bda != nil {
		return nil, 0, _bda
	}
	if _afg != nil {
		_cf.Log.Trace("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e")
		return _afg, PermOwner, nil
	}
	_cf.Log.Trace("\u0044\u0065bu\u0067\u0067\u0069n\u0067\u0020\u0061\u0075the\u006eti\u0063\u0061\u0074\u0069\u006f\u006e\u0020- \u0075\u0073\u0065\u0072\u0020\u0070\u0061s\u0073")
	_afg, _bda = _bed.alg6(d, pass)
	if _bda != nil {
		return nil, 0, _bda
	}
	if _afg != nil {
		_cf.Log.Trace("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e")
		return _afg, d.P, nil
	}
	return nil, 0, nil
}
func _cb(_ab _f.Block) _f.BlockMode { return (*ecbDecrypter)(_daf(_ab)) }
func (_dab stdHandlerR4) alg3(R int, _dbd, _cbb []byte) ([]byte, error) {
	var _ag []byte
	if len(_cbb) > 0 {
		_ag = _dab.alg3Key(R, _cbb)
	} else {
		_ag = _dab.alg3Key(R, _dbd)
	}
	_cea, _eafg := _g.NewCipher(_ag)
	if _eafg != nil {
		return nil, _ad.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_be := _dab.paddedPass(_dbd)
	_dac := make([]byte, len(_be))
	_cea.XORKeyStream(_dac, _be)
	if R >= 3 {
		_ga := make([]byte, len(_ag))
		for _cfd := 0; _cfd < 19; _cfd++ {
			for _eab := 0; _eab < len(_ag); _eab++ {
				_ga[_eab] = _ag[_eab] ^ byte(_cfd+1)
			}
			_dacd, _abb := _g.NewCipher(_ga)
			if _abb != nil {
				return nil, _ad.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
			}
			_dacd.XORKeyStream(_dac, _dac)
		}
	}
	return _dac, nil
}
func (_dagd stdHandlerR6) alg13(_acf *StdEncryptDict, _cgf []byte) error {
	if _fbc := _gba("\u0061\u006c\u00671\u0033", "\u004b\u0065\u0079", 32, _cgf); _fbc != nil {
		return _fbc
	}
	if _cddg := _gba("\u0061\u006c\u00671\u0033", "\u0050\u0065\u0072m\u0073", 16, _acf.Perms); _cddg != nil {
		return _cddg
	}
	_bbbb := make([]byte, 16)
	copy(_bbbb, _acf.Perms[:16])
	_abdg, _cgfg := _a.NewCipher(_cgf[:32])
	if _cgfg != nil {
		return _cgfg
	}
	_edb := _cb(_abdg)
	_edb.CryptBlocks(_bbbb, _bbbb)
	if !_fe.Equal(_bbbb[9:12], []byte("\u0061\u0064\u0062")) {
		return _ad.New("\u0064\u0065\u0063o\u0064\u0065\u0064\u0020p\u0065\u0072\u006d\u0069\u0073\u0073\u0069o\u006e\u0073\u0020\u0061\u0072\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_dbf := Permissions(_da.LittleEndian.Uint32(_bbbb[0:4]))
	if _dbf != _acf.P {
		return _ad.New("\u0070\u0065r\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0061il\u0065\u0064")
	}
	var _aeaf bool
	if _bbbb[8] == 'T' {
		_aeaf = true
	} else if _bbbb[8] == 'F' {
		_aeaf = false
	} else {
		return _ad.New("\u0064\u0065\u0063\u006f\u0064\u0065\u0064 \u006d\u0065\u0074a\u0064\u0061\u0074\u0061 \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0067\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	if _aeaf != _acf.EncryptMetadata {
		return _ad.New("\u006d\u0065t\u0061\u0064\u0061\u0074a\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e \u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066a\u0069\u006c\u0065\u0064")
	}
	return nil
}

// Allowed checks if a set of permissions can be granted.
func (_dc Permissions) Allowed(p2 Permissions) bool { return _dc&p2 == p2 }

type stdHandlerR4 struct {
	Length int
	ID0    string
}

func (_ddc stdHandlerR4) alg2(_dagg *StdEncryptDict, _adg []byte) []byte {
	_cf.Log.Trace("\u0061\u006c\u0067\u0032")
	_dcc := _ddc.paddedPass(_adg)
	_db := _ed.New()
	_db.Write(_dcc)
	_db.Write(_dagg.O)
	var _bgb [4]byte
	_da.LittleEndian.PutUint32(_bgb[:], uint32(_dagg.P))
	_db.Write(_bgb[:])
	_cf.Log.Trace("\u0067o\u0020\u0050\u003a\u0020\u0025\u0020x", _bgb)
	_db.Write([]byte(_ddc.ID0))
	_cf.Log.Trace("\u0074\u0068\u0069\u0073\u002e\u0052\u0020\u003d\u0020\u0025d\u0020\u0065\u006e\u0063\u0072\u0079\u0070t\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u0025\u0076", _dagg.R, _dagg.EncryptMetadata)
	if (_dagg.R >= 4) && !_dagg.EncryptMetadata {
		_db.Write([]byte{0xff, 0xff, 0xff, 0xff})
	}
	_eg := _db.Sum(nil)
	if _dagg.R >= 3 {
		_db = _ed.New()
		for _bce := 0; _bce < 50; _bce++ {
			_db.Reset()
			_db.Write(_eg[0 : _ddc.Length/8])
			_eg = _db.Sum(nil)
		}
	}
	if _dagg.R >= 3 {
		return _eg[0 : _ddc.Length/8]
	}
	return _eg[0:5]
}
func (_aea stdHandlerR6) alg11(_eef *StdEncryptDict, _dcgf []byte) ([]byte, error) {
	if _ceac := _gba("\u0061\u006c\u00671\u0031", "\u0055", 48, _eef.U); _ceac != nil {
		return nil, _ceac
	}
	_acgc := make([]byte, len(_dcgf)+8)
	_fff := copy(_acgc, _dcgf)
	_fff += copy(_acgc[_fff:], _eef.U[32:40])
	_bcea, _dddf := _aea.alg2b(_eef.R, _acgc, _dcgf, nil)
	if _dddf != nil {
		return nil, _dddf
	}
	_bcea = _bcea[:32]
	if !_fe.Equal(_bcea, _eef.U[:32]) {
		return nil, nil
	}
	return _bcea, nil
}

type ecbDecrypter ecb

func _ada(_ggga, _cga, _ddf []byte) ([]byte, error) {
	var (
		_eba, _cab, _df _bc.Hash
	)
	_eba = _ec.New()
	_ccbb := make([]byte, 64)
	_eddg := _eba
	_eddg.Write(_ggga)
	K := _eddg.Sum(_ccbb[:0])
	_gad := make([]byte, 64*(127+64+48))
	_gbad := func(_efb int) ([]byte, error) {
		_gbd := len(_cga) + len(K) + len(_ddf)
		_aad := _gad[:_gbd]
		_acee := copy(_aad, _cga)
		_acee += copy(_aad[_acee:], K[:])
		_acee += copy(_aad[_acee:], _ddf)
		if _acee != _gbd {
			_cf.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0072\u006f\u0075\u006ed\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u007ae\u002e")
			return nil, _ad.New("\u0077\u0072\u006f\u006e\u0067\u0020\u0073\u0069\u007a\u0065")
		}
		K1 := _gad[:_gbd*64]
		_ccb(K1, _gbd)
		_aaa, _aaaf := _abc(K[0:16])
		if _aaaf != nil {
			return nil, _aaaf
		}
		_dgb := _f.NewCBCEncrypter(_aaa, K[16:32])
		_dgb.CryptBlocks(K1, K1)
		E := K1
		_bbf := 0
		for _eaae := 0; _eaae < 16; _eaae++ {
			_bbf += int(E[_eaae] % 3)
		}
		var _aaaa _bc.Hash
		switch _bbf % 3 {
		case 0:
			_aaaa = _eba
		case 1:
			if _cab == nil {
				_cab = _d.New384()
			}
			_aaaa = _cab
		case 2:
			if _df == nil {
				_df = _d.New()
			}
			_aaaa = _df
		}
		_aaaa.Reset()
		_aaaa.Write(E)
		K = _aaaa.Sum(_ccbb[:0])
		return E, nil
	}
	for _fdg := 0; ; {
		E, _eca := _gbad(_fdg)
		if _eca != nil {
			return nil, _eca
		}
		_cgc := E[len(E)-1]
		_fdg++
		if _fdg >= 64 && _cgc <= uint8(_fdg-32) {
			break
		}
	}
	return K[:32], nil
}

var _ StdHandler = stdHandlerR4{}

func _abc(_gcg []byte) (_f.Block, error) {
	_eff, _aga := _a.NewCipher(_gcg)
	if _aga != nil {
		_cf.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0063\u0069p\u0068\u0065r\u003a\u0020\u0025\u0076", _aga)
		return nil, _aga
	}
	return _eff, nil
}

// AuthEvent is an event type that triggers authentication.
type AuthEvent string

func (_dddc stdHandlerR6) alg2b(R int, _abd, _agcg, _bad []byte) ([]byte, error) {
	if R == 5 {
		return _gfg(_abd)
	}
	return _ada(_abd, _agcg, _bad)
}
func (_cg errInvalidField) Error() string {
	return _gb.Sprintf("\u0025s\u003a\u0020e\u0078\u0070\u0065\u0063t\u0065\u0064\u0020%\u0073\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074o \u0062\u0065\u0020%\u0064\u0020b\u0079\u0074\u0065\u0073\u002c\u0020g\u006f\u0074 \u0025\u0064", _cg.Func, _cg.Field, _cg.Exp, _cg.Got)
}
func (_aab stdHandlerR6) alg12(_dbe *StdEncryptDict, _fgb []byte) ([]byte, error) {
	if _bbb := _gba("\u0061\u006c\u00671\u0032", "\u0055", 48, _dbe.U); _bbb != nil {
		return nil, _bbb
	}
	if _edde := _gba("\u0061\u006c\u00671\u0032", "\u004f", 48, _dbe.O); _edde != nil {
		return nil, _edde
	}
	_cdd := make([]byte, len(_fgb)+8+48)
	_efbd := copy(_cdd, _fgb)
	_efbd += copy(_cdd[_efbd:], _dbe.O[32:40])
	_efbd += copy(_cdd[_efbd:], _dbe.U[0:48])
	_ffg, _dbbc := _aab.alg2b(_dbe.R, _cdd, _fgb, _dbe.U[0:48])
	if _dbbc != nil {
		return nil, _dbbc
	}
	_ffg = _ffg[:32]
	if !_fe.Equal(_ffg, _dbe.O[:32]) {
		return nil, nil
	}
	return _ffg, nil
}

// NewHandlerR4 creates a new standard security handler for R<=4.
func NewHandlerR4(id0 string, length int) StdHandler { return stdHandlerR4{ID0: id0, Length: length} }
func _caa(_dag _f.Block) _f.BlockMode                { return (*ecbEncrypter)(_daf(_dag)) }

// Authenticate implements StdHandler interface.
func (_cedc stdHandlerR6) Authenticate(d *StdEncryptDict, pass []byte) ([]byte, Permissions, error) {
	return _cedc.alg2a(d, pass)
}

const (
	EventDocOpen = AuthEvent("\u0044o\u0063\u004f\u0070\u0065\u006e")
	EventEFOpen  = AuthEvent("\u0045\u0046\u004f\u0070\u0065\u006e")
)

func (_ggc *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%_ggc._ca != 0 {
		_cf.Log.Error("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073")
		return
	}
	if len(dst) < len(src) {
		_cf.Log.Error("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0064\u0065\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074")
		return
	}
	for len(src) > 0 {
		_ggc._ff.Decrypt(dst, src[:_ggc._ca])
		src = src[_ggc._ca:]
		dst = dst[_ggc._ca:]
	}
}

type ecbEncrypter ecb
