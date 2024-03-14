package bitwise

import (
	_b "encoding/binary"
	_g "errors"
	_aa "fmt"
	_ag "io"

	_ae "bitbucket.org/shenghui0779/gopdf/common"
	_d "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

var _ _ag.ByteWriter = &BufferedWriter{}

func (_cca *Writer) Data() []byte { return _cca._ccgg }
func (_ca *BufferedWriter) WriteByte(bt byte) error {
	if _ca._dc > len(_ca._add)-1 || (_ca._dc == len(_ca._add)-1 && _ca._adb != 0) {
		_ca.expandIfNeeded(1)
	}
	_ca.writeByte(bt)
	return nil
}

func (_aea *BufferedWriter) tryGrowByReslice(_eac int) bool {
	if _efa := len(_aea._add); _eac <= cap(_aea._add)-_efa {
		_aea._add = _aea._add[:_efa+_eac]
		return true
	}
	return false
}

func (_eaf *BufferedWriter) expandIfNeeded(_fd int) {
	if !_eaf.tryGrowByReslice(_fd) {
		_eaf.grow(_fd)
	}
}

func (_agd *Reader) ReadUint32() (uint32, error) {
	_bdc := make([]byte, 4)
	_, _cg := _agd.Read(_bdc)
	if _cg != nil {
		return 0, _cg
	}
	return _b.BigEndian.Uint32(_bdc), nil
}

var _ BinaryWriter = &Writer{}

func (_fdde *Reader) readBufferByte() (byte, error) {
	if _fdde._gba >= int64(_fdde._ec._ecf) {
		return 0, _ag.EOF
	}
	_fdde._acc = -1
	_fe := _fdde._ec._edc[int64(_fdde._ec._ebe)+_fdde._gba]
	_fdde._gba++
	_fdde._ebg = int(_fe)
	return _fe, nil
}

func (_gce *Reader) Reset() {
	_gce._gba = _gce._ccbb
	_gce._de = _gce._ccc
	_gce._ece = _gce._afa
	_gce._ebg = _gce._bg
}

func (_agc *Reader) ReadBit() (_df int, _ggf error) {
	_cbf, _ggf := _agc.readBool()
	if _ggf != nil {
		return 0, _ggf
	}
	if _cbf {
		_df = 1
	}
	return _df, nil
}

const (
	_c  = 64
	_ad = int(^uint(0) >> 1)
)

func (_agb *BufferedWriter) FinishByte() {
	if _agb._adb == 0 {
		return
	}
	_agb._adb = 0
	_agb._dc++
}

type BinaryWriter interface {
	BitWriter
	_ag.Writer
	_ag.ByteWriter
	Data() []byte
}

func (_bdde *Reader) NewPartialReader(offset, length int, relative bool) (*Reader, error) {
	if offset < 0 {
		return nil, _g.New("p\u0061\u0072\u0074\u0069\u0061\u006c\u0020\u0072\u0065\u0061\u0064\u0065\u0072\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062e \u006e\u0065\u0067a\u0074i\u0076\u0065")
	}
	if relative {
		offset = _bdde._ec._ebe + offset
	}
	if length > 0 {
		_ccca := len(_bdde._ec._edc)
		if relative {
			_ccca = _bdde._ec._ecf
		}
		if offset+length > _ccca {
			return nil, _aa.Errorf("\u0070\u0061r\u0074\u0069\u0061l\u0020\u0072\u0065\u0061\u0064e\u0072\u0020\u006f\u0066\u0066se\u0074\u0028\u0025\u0064\u0029\u002b\u006c\u0065\u006e\u0067\u0074\u0068\u0028\u0025\u0064\u0029\u003d\u0025d\u0020i\u0073\u0020\u0067\u0072\u0065\u0061ter\u0020\u0074\u0068\u0061\u006e\u0020\u0074\u0068\u0065\u0020\u006f\u0072ig\u0069n\u0061\u006c\u0020\u0072e\u0061d\u0065r\u0020\u006ce\u006e\u0067th\u003a\u0020\u0025\u0064", offset, length, offset+length, _bdde._ec._ecf)
		}
	}
	if length < 0 {
		_bf := len(_bdde._ec._edc)
		if relative {
			_bf = _bdde._ec._ecf
		}
		length = _bf - offset
	}
	return &Reader{_ec: readerSource{_edc: _bdde._ec._edc, _ecf: length, _ebe: offset}}, nil
}
func (_gcd *Reader) AbsolutePosition() int64 { return _gcd._gba + int64(_gcd._ec._ebe) }
func NewWriterMSB(data []byte) *Writer       { return &Writer{_ccgg: data, _cae: true} }
func (_ebf *BufferedWriter) writeByte(_bd byte) {
	switch {
	case _ebf._adb == 0:
		_ebf._add[_ebf._dc] = _bd
		_ebf._dc++
	case _ebf._gd:
		_ebf._add[_ebf._dc] |= _bd >> _ebf._adb
		_ebf._dc++
		_ebf._add[_ebf._dc] = byte(uint16(_bd) << (8 - _ebf._adb) & 0xff)
	default:
		_ebf._add[_ebf._dc] |= byte(uint16(_bd) << _ebf._adb & 0xff)
		_ebf._dc++
		_ebf._add[_ebf._dc] = _bd >> (8 - _ebf._adb)
	}
}

func (_cb *BufferedWriter) WriteBits(bits uint64, number int) (_fg int, _bed error) {
	const _gb = "\u0042u\u0066\u0066\u0065\u0072e\u0064\u0057\u0072\u0069\u0074e\u0072.\u0057r\u0069\u0074\u0065\u0072\u0042\u0069\u0074s"
	if number < 0 || number > 64 {
		return 0, _d.Errorf(_gb, "\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", number)
	}
	_ef := number / 8
	if _ef > 0 {
		_cbg := number - _ef*8
		for _ce := _ef - 1; _ce >= 0; _ce-- {
			_fc := byte((bits >> uint(_ce*8+_cbg)) & 0xff)
			if _bed = _cb.WriteByte(_fc); _bed != nil {
				return _fg, _d.Wrapf(_bed, _gb, "\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _ef-_ce+1)
			}
		}
		number -= _ef * 8
		if number == 0 {
			return _ef, nil
		}
	}
	var _cbd int
	for _ccb := 0; _ccb < number; _ccb++ {
		if _cb._gd {
			_cbd = int((bits >> uint(number-1-_ccb)) & 0x1)
		} else {
			_cbd = int(bits & 0x1)
			bits >>= 1
		}
		if _bed = _cb.WriteBit(_cbd); _bed != nil {
			return _fg, _d.Wrapf(_bed, _gb, "\u0062i\u0074\u003a\u0020\u0025\u0064", _ccb)
		}
	}
	return _ef, nil
}
func NewWriter(data []byte) *Writer         { return &Writer{_ccgg: data} }
func (_ff *Reader) ReadBool() (bool, error) { return _ff.readBool() }
func (_egb *Writer) writeBit(_fca uint8) error {
	if len(_egb._ccgg)-1 < _egb._gee {
		return _ag.EOF
	}
	_ceg := _egb._eba
	if _egb._cae {
		_ceg = 7 - _egb._eba
	}
	_egb._ccgg[_egb._gee] |= byte(uint16(_fca<<_ceg) & 0xff)
	_egb._eba++
	if _egb._eba == 8 {
		_egb._gee++
		_egb._eba = 0
	}
	return nil
}

func (_fgc *BufferedWriter) writeFullBytes(_ged []byte) int {
	_gcb := copy(_fgc._add[_fgc.fullOffset():], _ged)
	_fgc._dc += _gcb
	return _gcb
}
func (_da *BufferedWriter) Reset()        { _da._add = _da._add[:0]; _da._dc = 0; _da._adb = 0 }
func BufferedMSB() *BufferedWriter        { return &BufferedWriter{_gd: true} }
func (_e *BufferedWriter) ResetBitIndex() { _e._adb = 0 }

var _ _ag.Writer = &BufferedWriter{}

func (_bfb *Reader) Length() uint64 { return uint64(_bfb._ec._ecf) }

type StreamReader interface {
	_ag.Reader
	_ag.ByteReader
	_ag.Seeker
	Align() byte
	BitPosition() int
	Mark()
	Length() uint64
	ReadBit() (int, error)
	ReadBits(_fcc byte) (uint64, error)
	ReadBool() (bool, error)
	ReadUint32() (uint32, error)
	Reset()
	AbsolutePosition() int64
}

func (_fdd *Reader) BitPosition() int { return int(_fdd._de) }
func (_dee *Writer) Write(p []byte) (int, error) {
	if len(p) > _dee.byteCapacity() {
		return 0, _ag.EOF
	}
	for _, _egg := range p {
		if _fdcc := _dee.writeByte(_egg); _fdcc != nil {
			return 0, _fdcc
		}
	}
	return len(p), nil
}

func (_bdf *Writer) WriteBits(bits uint64, number int) (_cccae int, _accd error) {
	const _edce = "\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065r\u0042\u0069\u0074\u0073"
	if number < 0 || number > 64 {
		return 0, _d.Errorf(_edce, "\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", number)
	}
	if number == 0 {
		return 0, nil
	}
	_agba := number / 8
	if _agba > 0 {
		_acce := number - _agba*8
		for _gbg := _agba - 1; _gbg >= 0; _gbg-- {
			_cdg := byte((bits >> uint(_gbg*8+_acce)) & 0xff)
			if _accd = _bdf.WriteByte(_cdg); _accd != nil {
				return _cccae, _d.Wrapf(_accd, _edce, "\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _agba-_gbg+1)
			}
		}
		number -= _agba * 8
		if number == 0 {
			return _agba, nil
		}
	}
	var _abf int
	for _fddb := 0; _fddb < number; _fddb++ {
		if _bdf._cae {
			_abf = int((bits >> uint(number-1-_fddb)) & 0x1)
		} else {
			_abf = int(bits & 0x1)
			bits >>= 1
		}
		if _accd = _bdf.WriteBit(_abf); _accd != nil {
			return _cccae, _d.Wrapf(_accd, _edce, "\u0062i\u0074\u003a\u0020\u0025\u0064", _fddb)
		}
	}
	return _agba, nil
}

func (_dbg *Writer) FinishByte() {
	if _dbg._eba == 0 {
		return
	}
	_dbg._eba = 0
	_dbg._gee++
}
func (_ebb *Reader) RelativePosition() int64 { return _ebb._gba }
func (_aad *Reader) ReadBits(n byte) (_bcc uint64, _eg error) {
	if n < _aad._de {
		_bec := _aad._de - n
		_bcc = uint64(_aad._ece >> _bec)
		_aad._ece &= 1<<_bec - 1
		_aad._de = _bec
		return _bcc, nil
	}
	if n > _aad._de {
		if _aad._de > 0 {
			_bcc = uint64(_aad._ece)
			n -= _aad._de
		}
		for n >= 8 {
			_afb, _eafe := _aad.readBufferByte()
			if _eafe != nil {
				return 0, _eafe
			}
			_bcc = _bcc<<8 + uint64(_afb)
			n -= 8
		}
		if n > 0 {
			if _aad._ece, _eg = _aad.readBufferByte(); _eg != nil {
				return 0, _eg
			}
			_cd := 8 - n
			_bcc = _bcc<<n + uint64(_aad._ece>>_cd)
			_aad._ece &= 1<<_cd - 1
			_aad._de = _cd
		} else {
			_aad._de = 0
		}
		return _bcc, nil
	}
	_aad._de = 0
	return uint64(_aad._ece), nil
}

func (_ba *BufferedWriter) fullOffset() int {
	_ge := _ba._dc
	if _ba._adb != 0 {
		_ge++
	}
	return _ge
}
func (_edb *Writer) UseMSB() bool { return _edb._cae }
func (_ced *BufferedWriter) byteCapacity() int {
	_gcc := len(_ced._add) - _ced._dc
	if _ced._adb != 0 {
		_gcc--
	}
	return _gcc
}

func NewReader(data []byte) *Reader {
	return &Reader{_ec: readerSource{_edc: data, _ecf: len(data), _ebe: 0}}
}

func (_gc *BufferedWriter) SkipBits(skip int) error {
	if skip == 0 {
		return nil
	}
	_adbe := int(_gc._adb) + skip
	if _adbe >= 0 && _adbe < 8 {
		_gc._adb = uint8(_adbe)
		return nil
	}
	_adbe = int(_gc._adb) + _gc._dc*8 + skip
	if _adbe < 0 {
		return _d.Errorf("\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073", "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_f := _adbe / 8
	_cc := _adbe % 8
	_gc._adb = uint8(_cc)
	if _ab := _f - _gc._dc; _ab > 0 && len(_gc._add)-1 < _f {
		if _gc._adb != 0 {
			_ab++
		}
		_gc.expandIfNeeded(_ab)
	}
	_gc._dc = _f
	return nil
}

func (_gbb *Reader) Seek(offset int64, whence int) (int64, error) {
	_gbb._acc = -1
	_gbb._de = 0
	_gbb._ece = 0
	_gbb._ebg = 0
	var _afdb int64
	switch whence {
	case _ag.SeekStart:
		_afdb = offset
	case _ag.SeekCurrent:
		_afdb = _gbb._gba + offset
	case _ag.SeekEnd:
		_afdb = int64(_gbb._ec._ecf) + offset
	default:
		return 0, _g.New("\u0072\u0065\u0061de\u0072\u002e\u0052\u0065\u0061\u0064\u0065\u0072\u002eS\u0065e\u006b:\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0068\u0065\u006e\u0063\u0065")
	}
	if _afdb < 0 {
		return 0, _g.New("\u0072\u0065a\u0064\u0065\u0072\u002eR\u0065\u0061d\u0065\u0072\u002e\u0053\u0065\u0065\u006b\u003a \u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e")
	}
	_gbb._gba = _afdb
	_gbb._de = 0
	return _afdb, nil
}

func (_ga *BufferedWriter) grow(_af int) {
	if _ga._add == nil && _af < _c {
		_ga._add = make([]byte, _af, _c)
		return
	}
	_bc := len(_ga._add)
	if _ga._adb != 0 {
		_bc++
	}
	_cf := cap(_ga._add)
	switch {
	case _af <= _cf/2-_bc:
		_ae.Log.Trace("\u005b\u0042\u0075\u0066\u0066\u0065r\u0065\u0064\u0057\u0072\u0069t\u0065\u0072\u005d\u0020\u0067\u0072o\u0077\u0020\u002d\u0020\u0072e\u0073\u006c\u0069\u0063\u0065\u0020\u006f\u006e\u006c\u0079\u002e\u0020L\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0043\u0061\u0070\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u006e\u003a\u0020'\u0025\u0064\u0027", len(_ga._add), cap(_ga._add), _af)
		_ae.Log.Trace("\u0020\u006e\u0020\u003c\u003d\u0020\u0063\u0020\u002f\u0020\u0032\u0020\u002d\u006d\u002e \u0043:\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u006d\u003a\u0020\u0027\u0025\u0064\u0027", _cf, _bc)
		copy(_ga._add, _ga._add[_ga.fullOffset():])
	case _cf > _ad-_cf-_af:
		_ae.Log.Error("\u0042\u0055F\u0046\u0045\u0052 \u0074\u006f\u006f\u0020\u006c\u0061\u0072\u0067\u0065")
		return
	default:
		_eb := make([]byte, 2*_cf+_af)
		copy(_eb, _ga._add)
		_ga._add = _eb
	}
	_ga._add = _ga._add[:_bc+_af]
}

func (_ac *BufferedWriter) Write(d []byte) (int, error) {
	_ac.expandIfNeeded(len(d))
	if _ac._adb == 0 {
		return _ac.writeFullBytes(d), nil
	}
	return _ac.writeShiftedBytes(d), nil
}
func (_aef *BufferedWriter) Data() []byte { return _aef._add }

type BitWriter interface {
	WriteBit(_daf int) error
	WriteBits(_cab uint64, _ed int) (_gea int, _aab error)
	FinishByte()
	SkipBits(_ccg int) error
}

func (_fdc *Reader) read(_afdd []byte) (int, error) {
	if _fdc._gba >= int64(_fdc._ec._ecf) {
		return 0, _ag.EOF
	}
	_fdc._acc = -1
	_dag := copy(_afdd, _fdc._ec._edc[(int64(_fdc._ec._ebe)+_fdc._gba):(_fdc._ec._ebe+_fdc._ec._ecf)])
	_fdc._gba += int64(_dag)
	return _dag, nil
}

func (_fbf *Reader) readUnalignedByte() (_bde byte, _ebga error) {
	_ebd := _fbf._de
	_bde = _fbf._ece << (8 - _ebd)
	_fbf._ece, _ebga = _fbf.readBufferByte()
	if _ebga != nil {
		return 0, _ebga
	}
	_bde |= _fbf._ece >> _ebd
	_fbf._ece &= 1<<_ebd - 1
	return _bde, nil
}

func (_gec *Reader) ReadByte() (byte, error) {
	if _gec._de == 0 {
		return _gec.readBufferByte()
	}
	return _gec.readUnalignedByte()
}

type BufferedWriter struct {
	_add []byte
	_adb uint8
	_dc  int
	_gd  bool
}

func (_beb *BufferedWriter) Len() int        { return _beb.byteCapacity() }
func (_cabf *Writer) WriteByte(c byte) error { return _cabf.writeByte(c) }

type Reader struct {
	_ec   readerSource
	_ece  byte
	_de   byte
	_gba  int64
	_ebg  int
	_acc  int
	_ccbb int64
	_ccc  byte
	_afa  byte
	_bg   int
}

func (_ecg *Writer) ResetBit() { _ecg._eba = 0 }
func (_dae *Reader) Align() (_fgca byte) {
	_fgca = _dae._de
	_dae._de = 0
	return _fgca
}

func (_fa *BufferedWriter) WriteBit(bit int) error {
	if bit != 1 && bit != 0 {
		return _d.Errorf("\u0042\u0075\u0066fe\u0072\u0065\u0064\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074", "\u0062\u0069\u0074\u0020\u0076\u0061\u006cu\u0065\u0020\u006du\u0073\u0074\u0020\u0062e\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0030\u002c\u0031\u007d\u0020\u0062\u0075\u0074\u0020\u0069\u0073\u003a\u0020\u0025\u0064", bit)
	}
	if len(_fa._add)-1 < _fa._dc {
		_fa.expandIfNeeded(1)
	}
	_ea := _fa._adb
	if _fa._gd {
		_ea = 7 - _fa._adb
	}
	_fa._add[_fa._dc] |= byte(uint16(bit<<_ea) & 0xff)
	_fa._adb++
	if _fa._adb == 8 {
		_fa._dc++
		_fa._adb = 0
	}
	return nil
}

func (_eca *Writer) WriteBit(bit int) error {
	switch bit {
	case 0, 1:
		return _eca.writeBit(uint8(bit))
	}
	return _d.Error("\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074", "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0062\u0069\u0074\u0020v\u0061\u006c\u0075\u0065")
}

func (_ccbc *Writer) byteCapacity() int {
	_bdg := len(_ccbc._ccgg) - _ccbc._gee
	if _ccbc._eba != 0 {
		_bdg--
	}
	return _bdg
}

func (_fag *Reader) ConsumeRemainingBits() (uint64, error) {
	if _fag._de != 0 {
		return _fag.ReadBits(_fag._de)
	}
	return 0, nil
}

func (_fda *Reader) Mark() {
	_fda._ccbb = _fda._gba
	_fda._ccc = _fda._de
	_fda._afa = _fda._ece
	_fda._bg = _fda._ebg
}

func (_bdd *BufferedWriter) writeShiftedBytes(_gf []byte) int {
	for _, _afd := range _gf {
		_bdd.writeByte(_afd)
	}
	return len(_gf)
}
func (_fcb *Reader) AbsoluteLength() uint64 { return uint64(len(_fcb._ec._edc)) }
func (_aaf *Writer) writeByte(_cce byte) error {
	if _aaf._gee > len(_aaf._ccgg)-1 {
		return _ag.EOF
	}
	if _aaf._gee == len(_aaf._ccgg)-1 && _aaf._eba != 0 {
		return _ag.EOF
	}
	if _aaf._eba == 0 {
		_aaf._ccgg[_aaf._gee] = _cce
		_aaf._gee++
		return nil
	}
	if _aaf._cae {
		_aaf._ccgg[_aaf._gee] |= _cce >> _aaf._eba
		_aaf._gee++
		_aaf._ccgg[_aaf._gee] = byte(uint16(_cce) << (8 - _aaf._eba) & 0xff)
	} else {
		_aaf._ccgg[_aaf._gee] |= byte(uint16(_cce) << _aaf._eba & 0xff)
		_aaf._gee++
		_aaf._ccgg[_aaf._gee] = _cce >> (8 - _aaf._eba)
	}
	return nil
}

var (
	_ _ag.Reader     = &Reader{}
	_ _ag.ByteReader = &Reader{}
	_ _ag.Seeker     = &Reader{}
	_ StreamReader   = &Reader{}
)

type readerSource struct {
	_edc []byte
	_ebe int
	_ecf int
}

var _ BinaryWriter = &BufferedWriter{}

func (_gaf *Reader) readBool() (_fddd bool, _afbb error) {
	if _gaf._de == 0 {
		_gaf._ece, _afbb = _gaf.readBufferByte()
		if _afbb != nil {
			return false, _afbb
		}
		_fddd = (_gaf._ece & 0x80) != 0
		_gaf._ece, _gaf._de = _gaf._ece&0x7f, 7
		return _fddd, nil
	}
	_gaf._de--
	_fddd = (_gaf._ece & (1 << _gaf._de)) != 0
	_gaf._ece &= 1<<_gaf._de - 1
	return _fddd, nil
}

type Writer struct {
	_ccgg []byte
	_eba  uint8
	_gee  int
	_cae  bool
}

func (_gdf *Writer) SkipBits(skip int) error {
	const _dca = "\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073"
	if skip == 0 {
		return nil
	}
	_dd := int(_gdf._eba) + skip
	if _dd >= 0 && _dd < 8 {
		_gdf._eba = uint8(_dd)
		return nil
	}
	_dd = int(_gdf._eba) + _gdf._gee*8 + skip
	if _dd < 0 {
		return _d.Errorf(_dca, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_eafc := _dd / 8
	_ace := _dd % 8
	_ae.Log.Trace("\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073")
	_ae.Log.Trace("\u0042\u0069\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0042\u0079\u0074\u0065\u0049n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0046\u0075\u006c\u006c\u0042\u0069\u0074\u0073\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u004c\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027,\u0020\u0043\u0061p\u003a\u0020\u0027\u0025\u0064\u0027", _gdf._eba, _gdf._gee, int(_gdf._eba)+(_gdf._gee)*8, len(_gdf._ccgg), cap(_gdf._ccgg))
	_ae.Log.Trace("S\u006b\u0069\u0070\u003a\u0020\u0027%\u0064\u0027\u002c\u0020\u0064\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062i\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025d\u0027", skip, _dd, _ace)
	_gdf._eba = uint8(_ace)
	if _dbb := _eafc - _gdf._gee; _dbb > 0 && len(_gdf._ccgg)-1 < _eafc {
		_ae.Log.Trace("\u0042\u0079\u0074e\u0044\u0069\u0066\u0066\u003a\u0020\u0025\u0064", _dbb)
		return _d.Errorf(_dca, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_gdf._gee = _eafc
	_ae.Log.Trace("\u0042\u0069\u0074I\u006e\u0064\u0065\u0078:\u0020\u0027\u0025\u0064\u0027\u002c\u0020B\u0079\u0074\u0065\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027", _gdf._eba, _gdf._gee)
	return nil
}

func (_fdg *Reader) Read(p []byte) (_db int, _gg error) {
	if _fdg._de == 0 {
		return _fdg.read(p)
	}
	for ; _db < len(p); _db++ {
		if p[_db], _gg = _fdg.readUnalignedByte(); _gg != nil {
			return 0, _gg
		}
	}
	return _db, nil
}
