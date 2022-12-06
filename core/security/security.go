package security

import (
	_g "bytes"
	_b "crypto/aes"
	_f "crypto/cipher"
	_fe "crypto/md5"
	_ee "crypto/rand"
	_c "crypto/rc4"
	_dg "crypto/sha256"
	_a "crypto/sha512"
	_bc "encoding/binary"
	_fg "errors"
	_aa "fmt"
	_eg "hash"
	_e "io"
	_gf "math"

	_aab "bitbucket.org/shenghui0779/gopdf/common"
)

func _fa(_ea, _efd string, _ff int, _af []byte) error {
	if len(_af) < _ff {
		return errInvalidField{Func: _ea, Field: _efd, Exp: _ff, Got: len(_af)}
	}
	return nil
}

type ecbEncrypter ecb

func (_be *ecbEncrypter) BlockSize() int { return _be._bg }
func (_bfa stdHandlerR6) alg10(_gea *StdEncryptDict, _ebfea []byte) error {
	if _cdge := _fa("\u0061\u006c\u00671\u0030", "\u004b\u0065\u0079", 32, _ebfea); _cdge != nil {
		return _cdge
	}
	_edg := uint64(uint32(_gea.P)) | (_gf.MaxUint32 << 32)
	Perms := make([]byte, 16)
	_bc.LittleEndian.PutUint64(Perms[:8], _edg)
	if _gea.EncryptMetadata {
		Perms[8] = 'T'
	} else {
		Perms[8] = 'F'
	}
	copy(Perms[9:12], "\u0061\u0064\u0062")
	if _, _acaa := _e.ReadFull(_ee.Reader, Perms[12:16]); _acaa != nil {
		return _acaa
	}
	_cad, _eee := _egd(_ebfea[:32])
	if _eee != nil {
		return _eee
	}
	_ab := _ef(_cad)
	_ab.CryptBlocks(Perms, Perms)
	_gea.Perms = Perms[:16]
	return nil
}
func (stdHandlerR4) paddedPass(_eca []byte) []byte {
	_gc := make([]byte, 32)
	_ece := copy(_gc, _eca)
	for ; _ece < 32; _ece++ {
		_gc[_ece] = _fc[_ece-len(_eca)]
	}
	return _gc
}

type ecb struct {
	_gd _f.Block
	_bg int
}

// NewHandlerR6 creates a new standard security handler for R=5 and R=6.
func NewHandlerR6() StdHandler { return stdHandlerR6{} }

// Permissions is a bitmask of access permissions for a PDF file.
type Permissions uint32

func (_agb stdHandlerR4) alg5(_ceg []byte, _gff []byte) ([]byte, error) {
	_cdg := _fe.New()
	_cdg.Write([]byte(_fc))
	_cdg.Write([]byte(_agb.ID0))
	_bcb := _cdg.Sum(nil)
	_aab.Log.Trace("\u0061\u006c\u0067\u0035")
	_aab.Log.Trace("\u0065k\u0065\u0079\u003a\u0020\u0025\u0020x", _ceg)
	_aab.Log.Trace("\u0049D\u003a\u0020\u0025\u0020\u0078", _agb.ID0)
	if len(_bcb) != 16 {
		return nil, _fg.New("\u0068a\u0073\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074\u0020\u0031\u0036\u0020\u0062\u0079\u0074\u0065\u0073")
	}
	_cc, _cef := _c.NewCipher(_ceg)
	if _cef != nil {
		return nil, _fg.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_ffa := make([]byte, 16)
	_cc.XORKeyStream(_ffa, _bcb)
	_df := make([]byte, len(_ceg))
	for _eeg := 0; _eeg < 19; _eeg++ {
		for _fdf := 0; _fdf < len(_ceg); _fdf++ {
			_df[_fdf] = _ceg[_fdf] ^ byte(_eeg+1)
		}
		_cc, _cef = _c.NewCipher(_df)
		if _cef != nil {
			return nil, _fg.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
		}
		_cc.XORKeyStream(_ffa, _ffa)
		_aab.Log.Trace("\u0069\u0020\u003d\u0020\u0025\u0064\u002c\u0020\u0065\u006b\u0065\u0079:\u0020\u0025\u0020\u0078", _eeg, _df)
		_aab.Log.Trace("\u0069\u0020\u003d\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0020\u0078", _eeg, _ffa)
	}
	_dgf := make([]byte, 32)
	for _bee := 0; _bee < 16; _bee++ {
		_dgf[_bee] = _ffa[_bee]
	}
	_, _cef = _ee.Read(_dgf[16:32])
	if _cef != nil {
		return nil, _fg.New("\u0066a\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065n\u0020r\u0061n\u0064\u0020\u006e\u0075\u006d\u0062\u0065r")
	}
	return _dgf, nil
}
func _egd(_fcb []byte) (_f.Block, error) {
	_fce, _bcbe := _b.NewCipher(_fcb)
	if _bcbe != nil {
		_aab.Log.Error("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0063\u0069p\u0068\u0065r\u003a\u0020\u0025\u0076", _bcbe)
		return nil, _bcbe
	}
	return _fce, nil
}

// GenerateParams is the algorithm opposite to alg2a (R>=5).
// It generates U,O,UE,OE,Perms fields using AESv3 encryption.
// There is no algorithm number assigned to this function in the spec.
// It expects R, P and EncryptMetadata fields to be set.
func (_ede stdHandlerR6) GenerateParams(d *StdEncryptDict, opass, upass []byte) ([]byte, error) {
	_dfa := make([]byte, 32)
	if _, _dgfb := _e.ReadFull(_ee.Reader, _dfa); _dgfb != nil {
		return nil, _dgfb
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
	if _caa := _ede.alg8(d, _dfa, upass); _caa != nil {
		return nil, _caa
	}
	if _bb := _ede.alg9(d, _dfa, opass); _bb != nil {
		return nil, _bb
	}
	if d.R == 5 {
		return _dfa, nil
	}
	if _gage := _ede.alg10(d, _dfa); _gage != nil {
		return nil, _gage
	}
	return _dfa, nil
}
func (_bceb stdHandlerR4) alg6(_fgf *StdEncryptDict, _edb []byte) ([]byte, error) {
	var (
		_bd []byte
		_db error
	)
	_bac := _bceb.alg2(_fgf, _edb)
	if _fgf.R == 2 {
		_bd, _db = _bceb.alg4(_bac, _edb)
	} else if _fgf.R >= 3 {
		_bd, _db = _bceb.alg5(_bac, _edb)
	} else {
		return nil, _fg.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R")
	}
	if _db != nil {
		return nil, _db
	}
	_aab.Log.Trace("\u0063\u0068\u0065\u0063k:\u0020\u0025\u0020\u0078\u0020\u003d\u003d\u0020\u0025\u0020\u0078\u0020\u003f", string(_bd), string(_fgf.U))
	_gge := _bd
	_aed := _fgf.U
	if _fgf.R >= 3 {
		if len(_gge) > 16 {
			_gge = _gge[0:16]
		}
		if len(_aed) > 16 {
			_aed = _aed[0:16]
		}
	}
	if !_g.Equal(_gge, _aed) {
		return nil, nil
	}
	return _bac, nil
}

var _ StdHandler = stdHandlerR6{}

func _dde(_bgd, _ddb, _afe []byte) ([]byte, error) {
	var (
		_ggd, _gcf, _ffd _eg.Hash
	)
	_ggd = _dg.New()
	_cg := make([]byte, 64)
	_agg := _ggd
	_agg.Write(_bgd)
	K := _agg.Sum(_cg[:0])
	_cca := make([]byte, 64*(127+64+48))
	_age := func(_gbg int) ([]byte, error) {
		_ebfe := len(_ddb) + len(K) + len(_afe)
		_beb := _cca[:_ebfe]
		_cag := copy(_beb, _ddb)
		_cag += copy(_beb[_cag:], K[:])
		_cag += copy(_beb[_cag:], _afe)
		if _cag != _ebfe {
			_aab.Log.Error("E\u0052\u0052\u004f\u0052\u003a\u0020u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0072\u006f\u0075\u006ed\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u007ae\u002e")
			return nil, _fg.New("\u0077\u0072\u006f\u006e\u0067\u0020\u0073\u0069\u007a\u0065")
		}
		K1 := _cca[:_ebfe*64]
		_bcca(K1, _ebfe)
		_fbf, _ddd := _egd(K[0:16])
		if _ddd != nil {
			return nil, _ddd
		}
		_ggcf := _f.NewCBCEncrypter(_fbf, K[16:32])
		_ggcf.CryptBlocks(K1, K1)
		E := K1
		_dea := 0
		for _bad := 0; _bad < 16; _bad++ {
			_dea += int(E[_bad] % 3)
		}
		var _eaa _eg.Hash
		switch _dea % 3 {
		case 0:
			_eaa = _ggd
		case 1:
			if _gcf == nil {
				_gcf = _a.New384()
			}
			_eaa = _gcf
		case 2:
			if _ffd == nil {
				_ffd = _a.New()
			}
			_eaa = _ffd
		}
		_eaa.Reset()
		_eaa.Write(E)
		K = _eaa.Sum(_cg[:0])
		return E, nil
	}
	for _gaa := 0; ; {
		E, _dda := _age(_gaa)
		if _dda != nil {
			return nil, _dda
		}
		_bdbg := E[len(E)-1]
		_gaa++
		if _gaa >= 64 && _bdbg <= uint8(_gaa-32) {
			break
		}
	}
	return K[:32], nil
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

// Allowed checks if a set of permissions can be granted.
func (_ba Permissions) Allowed(p2 Permissions) bool { return _ba&p2 == p2 }
func (_cb *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%_cb._bg != 0 {
		_aab.Log.Error("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073")
		return
	}
	if len(dst) < len(src) {
		_aab.Log.Error("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0064\u0065\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074")
		return
	}
	for len(src) > 0 {
		_cb._gd.Decrypt(dst, src[:_cb._bg])
		src = src[_cb._bg:]
		dst = dst[_cb._bg:]
	}
}
func (_ae errInvalidField) Error() string {
	return _aa.Sprintf("\u0025s\u003a\u0020e\u0078\u0070\u0065\u0063t\u0065\u0064\u0020%\u0073\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074o \u0062\u0065\u0020%\u0064\u0020b\u0079\u0074\u0065\u0073\u002c\u0020g\u006f\u0074 \u0025\u0064", _ae.Func, _ae.Field, _ae.Exp, _ae.Got)
}
func (_fbb stdHandlerR6) alg2b(R int, _bae, _fae, _fbbg []byte) ([]byte, error) {
	if R == 5 {
		return _fdg(_bae)
	}
	return _dde(_bae, _fae, _fbbg)
}
func (_baf stdHandlerR6) alg11(_deae *StdEncryptDict, _bcde []byte) ([]byte, error) {
	if _fgfb := _fa("\u0061\u006c\u00671\u0031", "\u0055", 48, _deae.U); _fgfb != nil {
		return nil, _fgfb
	}
	_abb := make([]byte, len(_bcde)+8)
	_bfg := copy(_abb, _bcde)
	_bfg += copy(_abb[_bfg:], _deae.U[32:40])
	_gbfb, _afee := _baf.alg2b(_deae.R, _abb, _bcde, nil)
	if _afee != nil {
		return nil, _afee
	}
	_gbfb = _gbfb[:32]
	if !_g.Equal(_gbfb, _deae.U[:32]) {
		return nil, nil
	}
	return _gbfb, nil
}
func (_bf *ecbDecrypter) BlockSize() int { return _bf._bg }
func (_fea stdHandlerR6) alg2a(_ebf *StdEncryptDict, _dcb []byte) ([]byte, Permissions, error) {
	if _ffg := _fa("\u0061\u006c\u00672\u0061", "\u004f", 48, _ebf.O); _ffg != nil {
		return nil, 0, _ffg
	}
	if _bdf := _fa("\u0061\u006c\u00672\u0061", "\u0055", 48, _ebf.U); _bdf != nil {
		return nil, 0, _bdf
	}
	if len(_dcb) > 127 {
		_dcb = _dcb[:127]
	}
	_aaa, _dgc := _fea.alg12(_ebf, _dcb)
	if _dgc != nil {
		return nil, 0, _dgc
	}
	var (
		_afgf []byte
		_bdb  []byte
		_dacc []byte
	)
	var _ccb Permissions
	if len(_aaa) != 0 {
		_ccb = PermOwner
		_cfa := make([]byte, len(_dcb)+8+48)
		_fed := copy(_cfa, _dcb)
		_fed += copy(_cfa[_fed:], _ebf.O[40:48])
		copy(_cfa[_fed:], _ebf.U[0:48])
		_afgf = _cfa
		_bdb = _ebf.OE
		_dacc = _ebf.U[0:48]
	} else {
		_aaa, _dgc = _fea.alg11(_ebf, _dcb)
		if _dgc == nil && len(_aaa) == 0 {
			_aaa, _dgc = _fea.alg11(_ebf, []byte(""))
		}
		if _dgc != nil {
			return nil, 0, _dgc
		} else if len(_aaa) == 0 {
			return nil, 0, nil
		}
		_ccb = _ebf.P
		_aaba := make([]byte, len(_dcb)+8)
		_ggc := copy(_aaba, _dcb)
		copy(_aaba[_ggc:], _ebf.U[40:48])
		_afgf = _aaba
		_bdb = _ebf.UE
		_dacc = nil
	}
	if _fb := _fa("\u0061\u006c\u00672\u0061", "\u004b\u0065\u0079", 32, _bdb); _fb != nil {
		return nil, 0, _fb
	}
	_bdb = _bdb[:32]
	_baa, _dgc := _fea.alg2b(_ebf.R, _afgf, _dcb, _dacc)
	if _dgc != nil {
		return nil, 0, _dgc
	}
	_bga, _dgc := _b.NewCipher(_baa[:32])
	if _dgc != nil {
		return nil, 0, _dgc
	}
	_cfb := make([]byte, _b.BlockSize)
	_aeaf := _f.NewCBCDecrypter(_bga, _cfb)
	_dgd := make([]byte, 32)
	_aeaf.CryptBlocks(_dgd, _bdb)
	if _ebf.R == 5 {
		return _dgd, _ccb, nil
	}
	_dgc = _fea.alg13(_ebf, _dgd)
	if _dgc != nil {
		return nil, 0, _dgc
	}
	return _dgd, _ccb, nil
}
func (_bed stdHandlerR4) alg7(_ccf *StdEncryptDict, _ffb []byte) ([]byte, error) {
	_acae := _bed.alg3Key(_ccf.R, _ffb)
	_efg := make([]byte, len(_ccf.O))
	if _ccf.R == 2 {
		_ffc, _fee := _c.NewCipher(_acae)
		if _fee != nil {
			return nil, _fg.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072")
		}
		_ffc.XORKeyStream(_efg, _ccf.O)
	} else if _ccf.R >= 3 {
		_gbdc := append([]byte{}, _ccf.O...)
		for _aeb := 0; _aeb < 20; _aeb++ {
			_eba := append([]byte{}, _acae...)
			for _ged := 0; _ged < len(_acae); _ged++ {
				_eba[_ged] ^= byte(19 - _aeb)
			}
			_ga, _afg := _c.NewCipher(_eba)
			if _afg != nil {
				return nil, _fg.New("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072")
			}
			_ga.XORKeyStream(_efg, _gbdc)
			_gbdc = append([]byte{}, _efg...)
		}
	} else {
		return nil, _fg.New("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R")
	}
	_ace, _dggd := _bed.alg6(_ccf, _efg)
	if _dggd != nil {
		return nil, nil
	}
	return _ace, nil
}
func (_ega stdHandlerR4) alg2(_gda *StdEncryptDict, _cd []byte) []byte {
	_aab.Log.Trace("\u0061\u006c\u0067\u0032")
	_aag := _ega.paddedPass(_cd)
	_bcc := _fe.New()
	_bcc.Write(_aag)
	_bcc.Write(_gda.O)
	var _aea [4]byte
	_bc.LittleEndian.PutUint32(_aea[:], uint32(_gda.P))
	_bcc.Write(_aea[:])
	_aab.Log.Trace("\u0067o\u0020\u0050\u003a\u0020\u0025\u0020x", _aea)
	_bcc.Write([]byte(_ega.ID0))
	_aab.Log.Trace("\u0074\u0068\u0069\u0073\u002e\u0052\u0020\u003d\u0020\u0025d\u0020\u0065\u006e\u0063\u0072\u0079\u0070t\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u0025\u0076", _gda.R, _gda.EncryptMetadata)
	if (_gda.R >= 4) && !_gda.EncryptMetadata {
		_bcc.Write([]byte{0xff, 0xff, 0xff, 0xff})
	}
	_fgg := _bcc.Sum(nil)
	if _gda.R >= 3 {
		_bcc = _fe.New()
		for _ad := 0; _ad < 50; _ad++ {
			_bcc.Reset()
			_bcc.Write(_fgg[0 : _ega.Length/8])
			_fgg = _bcc.Sum(nil)
		}
	}
	if _gda.R >= 3 {
		return _fgg[0 : _ega.Length/8]
	}
	return _fgg[0:5]
}

// AuthEvent is an event type that triggers authentication.
type AuthEvent string
type ecbDecrypter ecb

var _ StdHandler = stdHandlerR4{}

// StdHandler is an interface for standard security handlers.
type StdHandler interface {

	// GenerateParams uses owner and user passwords to set encryption parameters and generate an encryption key.
	// It assumes that R, P and EncryptMetadata are already set.
	GenerateParams(_gb *StdEncryptDict, _dc, _eb []byte) ([]byte, error)

	// Authenticate uses encryption dictionary parameters and the password to calculate
	// the document encryption key. It also returns permissions that should be granted to a user.
	// In case of failed authentication, it returns empty key and zero permissions with no error.
	Authenticate(_dgg *StdEncryptDict, _bce []byte) ([]byte, Permissions, error)
}

func (_gdd *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%_gdd._bg != 0 {
		_aab.Log.Error("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073")
		return
	}
	if len(dst) < len(src) {
		_aab.Log.Error("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074")
		return
	}
	for len(src) > 0 {
		_gdd._gd.Encrypt(dst, src[:_gdd._bg])
		src = src[_gdd._bg:]
		dst = dst[_gdd._bg:]
	}
}

const _fc = "\x28\277\116\136\x4e\x75\x8a\x41\x64\000\x4e\x56\377" + "\xfa\001\010\056\x2e\x00\xb6\xd0\x68\076\x80\x2f\014" + "\251\xfe\x64\x53\x69\172"

// GenerateParams generates and sets O and U parameters for the encryption dictionary.
// It expects R, P and EncryptMetadata fields to be set.
func (_gffe stdHandlerR4) GenerateParams(d *StdEncryptDict, opass, upass []byte) ([]byte, error) {
	O, _daf := _gffe.alg3(d.R, upass, opass)
	if _daf != nil {
		_aab.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", _daf)
		return nil, _daf
	}
	d.O = O
	_aab.Log.Trace("\u0067\u0065\u006e\u0020\u004f\u003a\u0020\u0025\u0020\u0078", O)
	_gde := _gffe.alg2(d, upass)
	U, _daf := _gffe.alg5(_gde, upass)
	if _daf != nil {
		_aab.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029", _daf)
		return nil, _daf
	}
	d.U = U
	_aab.Log.Trace("\u0067\u0065\u006e\u0020\u0055\u003a\u0020\u0025\u0020\u0078", U)
	return _gde, nil
}
func (_fgfe stdHandlerR6) alg12(_cabb *StdEncryptDict, _bfc []byte) ([]byte, error) {
	if _fgfg := _fa("\u0061\u006c\u00671\u0032", "\u0055", 48, _cabb.U); _fgfg != nil {
		return nil, _fgfg
	}
	if _caec := _fa("\u0061\u006c\u00671\u0032", "\u004f", 48, _cabb.O); _caec != nil {
		return nil, _caec
	}
	_dce := make([]byte, len(_bfc)+8+48)
	_edcc := copy(_dce, _bfc)
	_edcc += copy(_dce[_edcc:], _cabb.O[32:40])
	_edcc += copy(_dce[_edcc:], _cabb.U[0:48])
	_cfbg, _dcc := _fgfe.alg2b(_cabb.R, _dce, _bfc, _cabb.U[0:48])
	if _dcc != nil {
		return nil, _dcc
	}
	_cfbg = _cfbg[:32]
	if !_g.Equal(_cfbg, _cabb.O[:32]) {
		return nil, nil
	}
	return _cfbg, nil
}

type stdHandlerR4 struct {
	Length int
	ID0    string
}

func _bcca(_edc []byte, _ffcb int) {
	_bdg := _ffcb
	for _bdg < len(_edc) {
		copy(_edc[_bdg:], _edc[:_bdg])
		_bdg *= 2
	}
}

type stdHandlerR6 struct{}

func (_gdaa stdHandlerR4) alg3Key(R int, _eda []byte) []byte {
	_aca := _fe.New()
	_gbd := _gdaa.paddedPass(_eda)
	_aca.Write(_gbd)
	if R >= 3 {
		for _ge := 0; _ge < 50; _ge++ {
			_ggg := _aca.Sum(nil)
			_aca = _fe.New()
			_aca.Write(_ggg)
		}
	}
	_ag := _aca.Sum(nil)
	if R == 2 {
		_ag = _ag[0:5]
	} else {
		_ag = _ag[0 : _gdaa.Length/8]
	}
	return _ag
}
func (_dac stdHandlerR4) alg4(_egb []byte, _cae []byte) ([]byte, error) {
	_cbf, _fd := _c.NewCipher(_egb)
	if _fd != nil {
		return nil, _fg.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_de := []byte(_fc)
	_dae := make([]byte, len(_de))
	_cbf.XORKeyStream(_dae, _de)
	return _dae, nil
}

type errInvalidField struct {
	Func  string
	Field string
	Exp   int
	Got   int
}

const (
	PermOwner             = Permissions(_gf.MaxUint32)
	PermPrinting          = Permissions(1 << 2)
	PermModify            = Permissions(1 << 3)
	PermExtractGraphics   = Permissions(1 << 4)
	PermAnnotate          = Permissions(1 << 5)
	PermFillForms         = Permissions(1 << 8)
	PermDisabilityExtract = Permissions(1 << 9)
	PermRotateInsert      = Permissions(1 << 10)
	PermFullPrintQuality  = Permissions(1 << 11)
)

func _ac(_ed _f.Block) _f.BlockMode { return (*ecbDecrypter)(_ca(_ed)) }

// NewHandlerR4 creates a new standard security handler for R<=4.
func NewHandlerR4(id0 string, length int) StdHandler { return stdHandlerR4{ID0: id0, Length: length} }
func _fdg(_ecb []byte) ([]byte, error) {
	_cfc := _dg.New()
	_cfc.Write(_ecb)
	return _cfc.Sum(nil), nil
}

const (
	EventDocOpen = AuthEvent("\u0044o\u0063\u004f\u0070\u0065\u006e")
	EventEFOpen  = AuthEvent("\u0045\u0046\u004f\u0070\u0065\u006e")
)

func (_agc stdHandlerR6) alg8(_bec *StdEncryptDict, _gfc []byte, _fdb []byte) error {
	if _fgef := _fa("\u0061\u006c\u0067\u0038", "\u004b\u0065\u0079", 32, _gfc); _fgef != nil {
		return _fgef
	}
	var _ecg [16]byte
	if _, _egc := _e.ReadFull(_ee.Reader, _ecg[:]); _egc != nil {
		return _egc
	}
	_ccc := _ecg[0:8]
	_fcef := _ecg[8:16]
	_ggf := make([]byte, len(_fdb)+len(_ccc))
	_eceg := copy(_ggf, _fdb)
	copy(_ggf[_eceg:], _ccc)
	_ffcg, _aee := _agc.alg2b(_bec.R, _ggf, _fdb, nil)
	if _aee != nil {
		return _aee
	}
	U := make([]byte, len(_ffcg)+len(_ccc)+len(_fcef))
	_eceg = copy(U, _ffcg[:32])
	_eceg += copy(U[_eceg:], _ccc)
	copy(U[_eceg:], _fcef)
	_bec.U = U
	_eceg = len(_fdb)
	copy(_ggf[_eceg:], _fcef)
	_ffcg, _aee = _agc.alg2b(_bec.R, _ggf, _fdb, nil)
	if _aee != nil {
		return _aee
	}
	_gag, _aee := _egd(_ffcg[:32])
	if _aee != nil {
		return _aee
	}
	_gffeb := make([]byte, _b.BlockSize)
	_ggfd := _f.NewCBCEncrypter(_gag, _gffeb)
	UE := make([]byte, 32)
	_ggfd.CryptBlocks(UE, _gfc[:32])
	_bec.UE = UE
	return nil
}
func (_egad stdHandlerR6) alg9(_fdd *StdEncryptDict, _fbbga []byte, _cegg []byte) error {
	if _gded := _fa("\u0061\u006c\u0067\u0039", "\u004b\u0065\u0079", 32, _fbbga); _gded != nil {
		return _gded
	}
	if _cdf := _fa("\u0061\u006c\u0067\u0039", "\u0055", 48, _fdd.U); _cdf != nil {
		return _cdf
	}
	var _gdeg [16]byte
	if _, _dgb := _e.ReadFull(_ee.Reader, _gdeg[:]); _dgb != nil {
		return _dgb
	}
	_dgbc := _gdeg[0:8]
	_caea := _gdeg[8:16]
	_dbf := _fdd.U[:48]
	_fcd := make([]byte, len(_cegg)+len(_dgbc)+len(_dbf))
	_aedf := copy(_fcd, _cegg)
	_aedf += copy(_fcd[_aedf:], _dgbc)
	_aedf += copy(_fcd[_aedf:], _dbf)
	_cgf, _dff := _egad.alg2b(_fdd.R, _fcd, _cegg, _dbf)
	if _dff != nil {
		return _dff
	}
	O := make([]byte, len(_cgf)+len(_dgbc)+len(_caea))
	_aedf = copy(O, _cgf[:32])
	_aedf += copy(O[_aedf:], _dgbc)
	_aedf += copy(O[_aedf:], _caea)
	_fdd.O = O
	_aedf = len(_cegg)
	_aedf += copy(_fcd[_aedf:], _caea)
	_cgf, _dff = _egad.alg2b(_fdd.R, _fcd, _cegg, _dbf)
	if _dff != nil {
		return _dff
	}
	_aeg, _dff := _egd(_cgf[:32])
	if _dff != nil {
		return _dff
	}
	_gfa := make([]byte, _b.BlockSize)
	_gbf := _f.NewCBCEncrypter(_aeg, _gfa)
	OE := make([]byte, 32)
	_gbf.CryptBlocks(OE, _fbbga[:32])
	_fdd.OE = OE
	return nil
}

// Authenticate implements StdHandler interface.
func (_aaf stdHandlerR4) Authenticate(d *StdEncryptDict, pass []byte) ([]byte, Permissions, error) {
	_aab.Log.Trace("\u0044\u0065b\u0075\u0067\u0067\u0069n\u0067\u0020a\u0075\u0074\u0068\u0065\u006e\u0074\u0069\u0063a\u0074\u0069\u006f\u006e\u0020\u002d\u0020\u006f\u0077\u006e\u0065\u0072 \u0070\u0061\u0073\u0073")
	_cbb, _edba := _aaf.alg7(d, pass)
	if _edba != nil {
		return nil, 0, _edba
	}
	if _cbb != nil {
		_aab.Log.Trace("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e")
		return _cbb, PermOwner, nil
	}
	_aab.Log.Trace("\u0044\u0065bu\u0067\u0067\u0069n\u0067\u0020\u0061\u0075the\u006eti\u0063\u0061\u0074\u0069\u006f\u006e\u0020- \u0075\u0073\u0065\u0072\u0020\u0070\u0061s\u0073")
	_cbb, _edba = _aaf.alg6(d, pass)
	if _edba != nil {
		return nil, 0, _edba
	}
	if _cbb != nil {
		_aab.Log.Trace("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e")
		return _cbb, d.P, nil
	}
	return nil, 0, nil
}
func (_ffcd stdHandlerR6) alg13(_gbda *StdEncryptDict, _bebc []byte) error {
	if _ead := _fa("\u0061\u006c\u00671\u0033", "\u004b\u0065\u0079", 32, _bebc); _ead != nil {
		return _ead
	}
	if _cdgg := _fa("\u0061\u006c\u00671\u0033", "\u0050\u0065\u0072m\u0073", 16, _gbda.Perms); _cdgg != nil {
		return _cdgg
	}
	_ecae := make([]byte, 16)
	copy(_ecae, _gbda.Perms[:16])
	_cegb, _faa := _b.NewCipher(_bebc[:32])
	if _faa != nil {
		return _faa
	}
	_bdgb := _ac(_cegb)
	_bdgb.CryptBlocks(_ecae, _ecae)
	if !_g.Equal(_ecae[9:12], []byte("\u0061\u0064\u0062")) {
		return _fg.New("\u0064\u0065\u0063o\u0064\u0065\u0064\u0020p\u0065\u0072\u006d\u0069\u0073\u0073\u0069o\u006e\u0073\u0020\u0061\u0072\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	_agf := Permissions(_bc.LittleEndian.Uint32(_ecae[0:4]))
	if _agf != _gbda.P {
		return _fg.New("\u0070\u0065r\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0061il\u0065\u0064")
	}
	var _ecac bool
	if _ecae[8] == 'T' {
		_ecac = true
	} else if _ecae[8] == 'F' {
		_ecac = false
	} else {
		return _fg.New("\u0064\u0065\u0063\u006f\u0064\u0065\u0064 \u006d\u0065\u0074a\u0064\u0061\u0074\u0061 \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0067\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064")
	}
	if _ecac != _gbda.EncryptMetadata {
		return _fg.New("\u006d\u0065t\u0061\u0064\u0061\u0074a\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e \u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066a\u0069\u006c\u0065\u0064")
	}
	return nil
}
func (_cdd stdHandlerR4) alg3(R int, _cab, _aac []byte) ([]byte, error) {
	var _dd []byte
	if len(_aac) > 0 {
		_dd = _cdd.alg3Key(R, _aac)
	} else {
		_dd = _cdd.alg3Key(R, _cab)
	}
	_fge, _efe := _c.NewCipher(_dd)
	if _efe != nil {
		return nil, _fg.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
	}
	_dca := _cdd.paddedPass(_cab)
	_cf := make([]byte, len(_dca))
	_fge.XORKeyStream(_cf, _dca)
	if R >= 3 {
		_fad := make([]byte, len(_dd))
		for _fgd := 0; _fgd < 19; _fgd++ {
			for _bcd := 0; _bcd < len(_dd); _bcd++ {
				_fad[_bcd] = _dd[_bcd] ^ byte(_fgd+1)
			}
			_ce, _fcc := _c.NewCipher(_fad)
			if _fcc != nil {
				return nil, _fg.New("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068")
			}
			_ce.XORKeyStream(_cf, _cf)
		}
	}
	return _cf, nil
}
func _ca(_da _f.Block) *ecb         { return &ecb{_gd: _da, _bg: _da.BlockSize()} }
func _ef(_gg _f.Block) _f.BlockMode { return (*ecbEncrypter)(_ca(_gg)) }

// Authenticate implements StdHandler interface.
func (_dfg stdHandlerR6) Authenticate(d *StdEncryptDict, pass []byte) ([]byte, Permissions, error) {
	return _dfg.alg2a(d, pass)
}
