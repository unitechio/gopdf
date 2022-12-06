package bitwise

import (
	_g "encoding/binary"
	_e "errors"
	_f "io"

	_a "bitbucket.org/shenghui0779/gopdf/common"
	_ee "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

type Reader struct {
	_adb []byte
	_bab byte
	_faa byte
	_ce  int64
	_ffd int
	_gag int
	_db  int64
	_cdf byte
	_ded byte
	_bae int
}

func (_gbc *Reader) ReadBool() (bool, error) { return _gbc.readBool() }
func BufferedMSB() *BufferedWriter           { return &BufferedWriter{_bb: true} }
func (_ecfc *SubstreamReader) Mark()         { _ecfc._gbf = _ecfc._be; _ecfc._dec = _ecfc._aaf }

var _ BinaryWriter = &BufferedWriter{}

func (_ggd *BufferedWriter) byteCapacity() int {
	_ea := len(_ggd._eed) - _ggd._bc
	if _ggd._b != 0 {
		_ea--
	}
	return _ea
}
func (_dca *SubstreamReader) ReadByte() (byte, error) {
	if _dca._aaf == 0 {
		return _dca.readBufferByte()
	}
	return _dca.readUnalignedByte()
}
func (_aca *BufferedWriter) WriteBits(bits uint64, number int) (_gga int, _ad error) {
	const _faf = "\u0042u\u0066\u0066\u0065\u0072e\u0064\u0057\u0072\u0069\u0074e\u0072.\u0057r\u0069\u0074\u0065\u0072\u0042\u0069\u0074s"
	if number < 0 || number > 64 {
		return 0, _ee.Errorf(_faf, "\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", number)
	}
	_fe := number / 8
	if _fe > 0 {
		_bf := number - _fe*8
		for _fg := _fe - 1; _fg >= 0; _fg-- {
			_eec := byte((bits >> uint(_fg*8+_bf)) & 0xff)
			if _ad = _aca.WriteByte(_eec); _ad != nil {
				return _gga, _ee.Wrapf(_ad, _faf, "\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _fe-_fg+1)
			}
		}
		number -= _fe * 8
		if number == 0 {
			return _fe, nil
		}
	}
	var _bd int
	for _bg := 0; _bg < number; _bg++ {
		if _aca._bb {
			_bd = int((bits >> uint(number-1-_bg)) & 0x1)
		} else {
			_bd = int(bits & 0x1)
			bits >>= 1
		}
		if _ad = _aca.WriteBit(_bd); _ad != nil {
			return _gga, _ee.Wrapf(_ad, _faf, "\u0062i\u0074\u003a\u0020\u0025\u0064", _bg)
		}
	}
	return _fe, nil
}
func (_eac *Reader) ReadBit() (_bbd int, _dfb error) {
	_fef, _dfb := _eac.readBool()
	if _dfb != nil {
		return 0, _dfb
	}
	if _fef {
		_bbd = 1
	}
	return _bbd, nil
}

var _ _f.Writer = &BufferedWriter{}

func (_ccc *Reader) BitPosition() int  { return int(_ccc._faa) }
func (_cf *Reader) Align() (_ced byte) { _ced = _cf._faa; _cf._faa = 0; return _ced }

type BitWriter interface {
	WriteBit(_ecg int) error
	WriteBits(_ebf uint64, _ecf int) (_edg int, _aag error)
	FinishByte()
	SkipBits(_cc int) error
}

func (_fc *SubstreamReader) readUnalignedByte() (_fdge byte, _efd error) {
	_dcda := _fc._aaf
	_fdge = _fc._eae << (8 - _dcda)
	_fc._eae, _efd = _fc.readBufferByte()
	if _efd != nil {
		return 0, _efd
	}
	_fdge |= _fc._eae >> _dcda
	_fc._eae &= 1<<_dcda - 1
	return _fdge, nil
}
func (_fee *SubstreamReader) Align() (_bga byte) {
	_bga = _fee._aaf
	_fee._aaf = 0
	return _bga
}
func (_dce *Reader) ConsumeRemainingBits() (uint64, error) {
	if _dce._faa != 0 {
		return _dce.ReadBits(_dce._faa)
	}
	return 0, nil
}
func (_gfg *Reader) readBool() (_acb bool, _bda error) {
	if _gfg._faa == 0 {
		_gfg._bab, _bda = _gfg.readBufferByte()
		if _bda != nil {
			return false, _bda
		}
		_acb = (_gfg._bab & 0x80) != 0
		_gfg._bab, _gfg._faa = _gfg._bab&0x7f, 7
		return _acb, nil
	}
	_gfg._faa--
	_acb = (_gfg._bab & (1 << _gfg._faa)) != 0
	_gfg._bab &= 1<<_gfg._faa - 1
	return _acb, nil
}
func (_ccf *SubstreamReader) readBufferByte() (byte, error) {
	if _ccf._be >= _ccf._ace {
		return 0, _f.EOF
	}
	if _ccf._be >= _ccf._dcf || _ccf._be < _ccf._cdff {
		if _fed := _ccf.fillBuffer(); _fed != nil {
			return 0, _fed
		}
	}
	_abb := _ccf._da[_ccf._be-_ccf._cdff]
	_ccf._be++
	return _abb, nil
}

var _ BinaryWriter = &Writer{}
var (
	_ _f.Reader     = &Reader{}
	_ _f.ByteReader = &Reader{}
	_ _f.Seeker     = &Reader{}
	_ StreamReader  = &Reader{}
)

func (_bdbc *SubstreamReader) ReadBool() (bool, error) { return _bdbc.readBool() }
func (_aa *BufferedWriter) ResetBitIndex()             { _aa._b = 0 }
func (_cge *Writer) WriteBit(bit int) error {
	switch bit {
	case 0, 1:
		return _cge.writeBit(uint8(bit))
	}
	return _ee.Error("\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074", "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0062\u0069\u0074\u0020v\u0061\u006c\u0075\u0065")
}
func (_gde *Writer) byteCapacity() int {
	_bdbfd := len(_gde._cde) - _gde._bed
	if _gde._dbe != 0 {
		_bdbfd--
	}
	return _bdbfd
}

type StreamReader interface {
	_f.Reader
	_f.ByteReader
	_f.Seeker
	Align() byte
	BitPosition() int
	Mark()
	Length() uint64
	ReadBit() (int, error)
	ReadBits(_cad byte) (uint64, error)
	ReadBool() (bool, error)
	ReadUint32() (uint32, error)
	Reset()
	StreamPosition() int64
}
type SubstreamReader struct {
	_be   uint64
	_ebgf StreamReader
	_aff  uint64
	_ace  uint64
	_da   []byte
	_cdff uint64
	_dcf  uint64
	_eae  byte
	_aaf  byte
	_gbf  uint64
	_dec  byte
}

const (
	_ff = 64
	_ae = int(^uint(0) >> 1)
)

func (_edb *Reader) ReadBits(n byte) (_cag uint64, _bdb error) {
	if n < _edb._faa {
		_ade := _edb._faa - n
		_cag = uint64(_edb._bab >> _ade)
		_edb._bab &= 1<<_ade - 1
		_edb._faa = _ade
		return _cag, nil
	}
	if n > _edb._faa {
		if _edb._faa > 0 {
			_cag = uint64(_edb._bab)
			n -= _edb._faa
		}
		for n >= 8 {
			_ecb, _ef := _edb.readBufferByte()
			if _ef != nil {
				return 0, _ef
			}
			_cag = _cag<<8 + uint64(_ecb)
			n -= 8
		}
		if n > 0 {
			if _edb._bab, _bdb = _edb.readBufferByte(); _bdb != nil {
				return 0, _bdb
			}
			_agf := 8 - n
			_cag = _cag<<n + uint64(_edb._bab>>_agf)
			_edb._bab &= 1<<_agf - 1
			_edb._faa = _agf
		} else {
			_edb._faa = 0
		}
		return _cag, nil
	}
	_edb._faa = 0
	return uint64(_edb._bab), nil
}
func (_eea *Writer) writeBit(_cgag uint8) error {
	if len(_eea._cde)-1 < _eea._bed {
		return _f.EOF
	}
	_cccf := _eea._dbe
	if _eea._affb {
		_cccf = 7 - _eea._dbe
	}
	_eea._cde[_eea._bed] |= byte(uint16(_cgag<<_cccf) & 0xff)
	_eea._dbe++
	if _eea._dbe == 8 {
		_eea._bed++
		_eea._dbe = 0
	}
	return nil
}
func (_cfa *Reader) StreamPosition() int64 { return _cfa._ce }
func (_cga *SubstreamReader) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case _f.SeekStart:
		_cga._be = uint64(offset)
	case _f.SeekCurrent:
		_cga._be += uint64(offset)
	case _f.SeekEnd:
		_cga._be = _cga._ace + uint64(offset)
	default:
		return 0, _e.New("\u0072\u0065\u0061d\u0065\u0072\u002e\u0053\u0075\u0062\u0073\u0074\u0072\u0065\u0061\u006d\u0052\u0065\u0061\u0064\u0065\u0072\u002e\u0053\u0065\u0065\u006b\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0077\u0068\u0065\u006e\u0063\u0065")
	}
	_cga._aaf = 0
	return int64(_cga._be), nil
}
func (_cca *Writer) WriteByte(c byte) error { return _cca.writeByte(c) }
func (_cg *Reader) ReadByte() (byte, error) {
	if _cg._faa == 0 {
		return _cg.readBufferByte()
	}
	return _cg.readUnalignedByte()
}
func NewWriterMSB(data []byte) *Writer { return &Writer{_cde: data, _affb: true} }
func (_eca *Reader) read(_ebe []byte) (int, error) {
	if _eca._ce >= int64(len(_eca._adb)) {
		return 0, _f.EOF
	}
	_eca._gag = -1
	_fd := copy(_ebe, _eca._adb[_eca._ce:])
	_eca._ce += int64(_fd)
	return _fd, nil
}
func (_gc *Reader) readUnalignedByte() (_aagg byte, _gafb error) {
	_dcd := _gc._faa
	_aagg = _gc._bab << (8 - _dcd)
	_gc._bab, _gafb = _gc.readBufferByte()
	if _gafb != nil {
		return 0, _gafb
	}
	_aagg |= _gc._bab >> _dcd
	_gc._bab &= 1<<_dcd - 1
	return _aagg, nil
}
func (_ca *BufferedWriter) writeShiftedBytes(_bdg []byte) int {
	for _, _ab := range _bdg {
		_ca.writeByte(_ab)
	}
	return len(_bdg)
}
func (_adg *Reader) ReadUint32() (uint32, error) {
	_baec := make([]byte, 4)
	_, _dfff := _adg.Read(_baec)
	if _dfff != nil {
		return 0, _dfff
	}
	return _g.BigEndian.Uint32(_baec), nil
}
func (_de *BufferedWriter) WriteBit(bit int) error {
	if bit != 1 && bit != 0 {
		return _ee.Errorf("\u0042\u0075\u0066fe\u0072\u0065\u0064\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074", "\u0062\u0069\u0074\u0020\u0076\u0061\u006cu\u0065\u0020\u006du\u0073\u0074\u0020\u0062e\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0030\u002c\u0031\u007d\u0020\u0062\u0075\u0074\u0020\u0069\u0073\u003a\u0020\u0025\u0064", bit)
	}
	if len(_de._eed)-1 < _de._bc {
		_de.expandIfNeeded(1)
	}
	_eg := _de._b
	if _de._bb {
		_eg = 7 - _de._b
	}
	_de._eed[_de._bc] |= byte(uint16(bit<<_eg) & 0xff)
	_de._b++
	if _de._b == 8 {
		_de._bc++
		_de._b = 0
	}
	return nil
}

var _ _f.ByteWriter = &BufferedWriter{}

func (_gaf *Reader) readBufferByte() (byte, error) {
	if _gaf._ce >= int64(len(_gaf._adb)) {
		return 0, _f.EOF
	}
	_gaf._gag = -1
	_bbc := _gaf._adb[_gaf._ce]
	_gaf._ce++
	_gaf._ffd = int(_bbc)
	return _bbc, nil
}
func (_ag *BufferedWriter) Len() int { return _ag.byteCapacity() }
func (_eeb *BufferedWriter) expandIfNeeded(_gf int) {
	if !_eeb.tryGrowByReslice(_gf) {
		_eeb.grow(_gf)
	}
}
func (_def *SubstreamReader) readBool() (_eff bool, _efb error) {
	if _def._aaf == 0 {
		_def._eae, _efb = _def.readBufferByte()
		if _efb != nil {
			return false, _efb
		}
		_eff = (_def._eae & 0x80) != 0
		_def._eae, _def._aaf = _def._eae&0x7f, 7
		return _eff, nil
	}
	_def._aaf--
	_eff = (_def._eae & (1 << _def._aaf)) != 0
	_def._eae &= 1<<_def._aaf - 1
	return _eff, nil
}
func (_bcg *BufferedWriter) Data() []byte { return _bcg._eed }

type BufferedWriter struct {
	_eed []byte
	_b   uint8
	_bc  int
	_bb  bool
}
type Writer struct {
	_cde  []byte
	_dbe  uint8
	_bed  int
	_affb bool
}

func (_ffcg *SubstreamReader) Reset() { _ffcg._be = _ffcg._gbf; _ffcg._aaf = _ffcg._dec }
func (_ge *SubstreamReader) Read(b []byte) (_dbf int, _gagc error) {
	if _ge._be >= _ge._ace {
		_a.Log.Trace("\u0053\u0074\u0072e\u0061\u006d\u0050\u006fs\u003a\u0020\u0027\u0025\u0064\u0027\u0020>\u003d\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027", _ge._be, _ge._ace)
		return 0, _f.EOF
	}
	for ; _dbf < len(b); _dbf++ {
		if b[_dbf], _gagc = _ge.readUnalignedByte(); _gagc != nil {
			if _gagc == _f.EOF {
				return _dbf, nil
			}
			return 0, _gagc
		}
	}
	return _dbf, nil
}
func (_bggb *Writer) Data() []byte  { return _bggb._cde }
func (_abd *Writer) UseMSB() bool   { return _abd._affb }
func (_ddb *Reader) Length() uint64 { return uint64(len(_ddb._adb)) }
func (_gdg *BufferedWriter) writeByte(_dc byte) {
	switch {
	case _gdg._b == 0:
		_gdg._eed[_gdg._bc] = _dc
		_gdg._bc++
	case _gdg._bb:
		_gdg._eed[_gdg._bc] |= _dc >> _gdg._b
		_gdg._bc++
		_gdg._eed[_gdg._bc] = byte(uint16(_dc) << (8 - _gdg._b) & 0xff)
	default:
		_gdg._eed[_gdg._bc] |= byte(uint16(_dc) << _gdg._b & 0xff)
		_gdg._bc++
		_gdg._eed[_gdg._bc] = _dc >> (8 - _gdg._b)
	}
}
func (_gb *BufferedWriter) WriteByte(bt byte) error {
	if _gb._bc > len(_gb._eed)-1 || (_gb._bc == len(_gb._eed)-1 && _gb._b != 0) {
		_gb.expandIfNeeded(1)
	}
	_gb.writeByte(bt)
	return nil
}
func (_fa *BufferedWriter) FinishByte() {
	if _fa._b == 0 {
		return
	}
	_fa._b = 0
	_fa._bc++
}
func (_dfc *SubstreamReader) ReadBit() (_gac int, _edd error) {
	_bfd, _edd := _dfc.readBool()
	if _edd != nil {
		return 0, _edd
	}
	if _bfd {
		_gac = 1
	}
	return _gac, nil
}
func (_cb *SubstreamReader) Length() uint64 { return _cb._ace }
func (_ecaa *Writer) WriteBits(bits uint64, number int) (_cfg int, _bdbf error) {
	const _gfc = "\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065r\u0042\u0069\u0074\u0073"
	if number < 0 || number > 64 {
		return 0, _ee.Errorf(_gfc, "\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", number)
	}
	if number == 0 {
		return 0, nil
	}
	_gad := number / 8
	if _gad > 0 {
		_caa := number - _gad*8
		for _age := _gad - 1; _age >= 0; _age-- {
			_fcf := byte((bits >> uint(_age*8+_caa)) & 0xff)
			if _bdbf = _ecaa.WriteByte(_fcf); _bdbf != nil {
				return _cfg, _ee.Wrapf(_bdbf, _gfc, "\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _gad-_age+1)
			}
		}
		number -= _gad * 8
		if number == 0 {
			return _gad, nil
		}
	}
	var _cbf int
	for _bec := 0; _bec < number; _bec++ {
		if _ecaa._affb {
			_cbf = int((bits >> uint(number-1-_bec)) & 0x1)
		} else {
			_cbf = int(bits & 0x1)
			bits >>= 1
		}
		if _bdbf = _ecaa.WriteBit(_cbf); _bdbf != nil {
			return _cfg, _ee.Wrapf(_bdbf, _gfc, "\u0062i\u0074\u003a\u0020\u0025\u0064", _bec)
		}
	}
	return _gad, nil
}
func NewSubstreamReader(r StreamReader, offset, length uint64) (*SubstreamReader, error) {
	if r == nil {
		return nil, _e.New("\u0072o\u006ft\u0020\u0072\u0065\u0061\u0064e\u0072\u0020i\u0073\u0020\u006e\u0069\u006c")
	}
	const _gbg = 1000 * 1000
	_dge := length
	if _dge > _gbg {
		_dge = _gbg
	}
	_a.Log.Trace("\u004e\u0065\u0077\u0053\u0075\u0062\u0073\u0074r\u0065\u0061\u006dRe\u0061\u0064\u0065\u0072\u0020\u0061t\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0064\u0020\u0077\u0069\u0074h\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a \u0025\u0064", offset, length)
	return &SubstreamReader{_ebgf: r, _aff: offset, _ace: length, _da: make([]byte, _dge)}, nil
}
func (_eb *BufferedWriter) Reset() { _eb._eed = _eb._eed[:0]; _eb._bc = 0; _eb._b = 0 }
func (_fgb *BufferedWriter) fullOffset() int {
	_ba := _fgb._bc
	if _fgb._b != 0 {
		_ba++
	}
	return _ba
}
func (_fgg *Reader) Mark() {
	_fgg._db = _fgg._ce
	_fgg._cdf = _fgg._faa
	_fgg._ded = _fgg._bab
	_fgg._bae = _fgg._ffd
}
func (_cgd *Writer) SkipBits(skip int) error {
	const _gage = "\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073"
	if skip == 0 {
		return nil
	}
	_eba := int(_cgd._dbe) + skip
	if _eba >= 0 && _eba < 8 {
		_cgd._dbe = uint8(_eba)
		return nil
	}
	_eba = int(_cgd._dbe) + _cgd._bed*8 + skip
	if _eba < 0 {
		return _ee.Errorf(_gage, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_bgc := _eba / 8
	_dbg := _eba % 8
	_a.Log.Trace("\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073")
	_a.Log.Trace("\u0042\u0069\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0042\u0079\u0074\u0065\u0049n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0046\u0075\u006c\u006c\u0042\u0069\u0074\u0073\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u004c\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027,\u0020\u0043\u0061p\u003a\u0020\u0027\u0025\u0064\u0027", _cgd._dbe, _cgd._bed, int(_cgd._dbe)+(_cgd._bed)*8, len(_cgd._cde), cap(_cgd._cde))
	_a.Log.Trace("S\u006b\u0069\u0070\u003a\u0020\u0027%\u0064\u0027\u002c\u0020\u0064\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062i\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025d\u0027", skip, _eba, _dbg)
	_cgd._dbe = uint8(_dbg)
	if _fede := _bgc - _cgd._bed; _fede > 0 && len(_cgd._cde)-1 < _bgc {
		_a.Log.Trace("\u0042\u0079\u0074e\u0044\u0069\u0066\u0066\u003a\u0020\u0025\u0064", _fede)
		return _ee.Errorf(_gage, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_cgd._bed = _bgc
	_a.Log.Trace("\u0042\u0069\u0074I\u006e\u0064\u0065\u0078:\u0020\u0027\u0025\u0064\u0027\u002c\u0020B\u0079\u0074\u0065\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027", _cgd._dbe, _cgd._bed)
	return nil
}
func NewReader(data []byte) *Reader            { return &Reader{_adb: data} }
func (_bgg *SubstreamReader) BitPosition() int { return int(_bgg._aaf) }
func (_ffg *Reader) Seek(offset int64, whence int) (int64, error) {
	_ffg._gag = -1
	var _fba int64
	switch whence {
	case _f.SeekStart:
		_fba = offset
	case _f.SeekCurrent:
		_fba = _ffg._ce + offset
	case _f.SeekEnd:
		_fba = int64(len(_ffg._adb)) + offset
	default:
		return 0, _e.New("\u0072\u0065\u0061de\u0072\u002e\u0052\u0065\u0061\u0064\u0065\u0072\u002eS\u0065e\u006b:\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0068\u0065\u006e\u0063\u0065")
	}
	if _fba < 0 {
		return 0, _e.New("\u0072\u0065a\u0064\u0065\u0072\u002eR\u0065\u0061d\u0065\u0072\u002e\u0053\u0065\u0065\u006b\u003a \u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e")
	}
	_ffg._ce = _fba
	_ffg._faa = 0
	return _fba, nil
}
func (_baa *Reader) Read(p []byte) (_aae int, _dba error) {
	if _baa._faa == 0 {
		return _baa.read(p)
	}
	for ; _aae < len(p); _aae++ {
		if p[_aae], _dba = _baa.readUnalignedByte(); _dba != nil {
			return 0, _dba
		}
	}
	return _aae, nil
}
func (_bfg *Writer) writeByte(_gcg byte) error {
	if _bfg._bed > len(_bfg._cde)-1 {
		return _f.EOF
	}
	if _bfg._bed == len(_bfg._cde)-1 && _bfg._dbe != 0 {
		return _f.EOF
	}
	if _bfg._dbe == 0 {
		_bfg._cde[_bfg._bed] = _gcg
		_bfg._bed++
		return nil
	}
	if _bfg._affb {
		_bfg._cde[_bfg._bed] |= _gcg >> _bfg._dbe
		_bfg._bed++
		_bfg._cde[_bfg._bed] = byte(uint16(_gcg) << (8 - _bfg._dbe) & 0xff)
	} else {
		_bfg._cde[_bfg._bed] |= byte(uint16(_gcg) << _bfg._dbe & 0xff)
		_bfg._bed++
		_bfg._cde[_bfg._bed] = _gcg >> (8 - _bfg._dbe)
	}
	return nil
}
func _gbgc(_bgad, _cef uint64) uint64 {
	if _bgad < _cef {
		return _bgad
	}
	return _cef
}
func (_bdf *Writer) ResetBit() { _bdf._dbe = 0 }
func (_ffc *BufferedWriter) SkipBits(skip int) error {
	if skip == 0 {
		return nil
	}
	_c := int(_ffc._b) + skip
	if _c >= 0 && _c < 8 {
		_ffc._b = uint8(_c)
		return nil
	}
	_c = int(_ffc._b) + _ffc._bc*8 + skip
	if _c < 0 {
		return _ee.Errorf("\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073", "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_fag := _c / 8
	_gg := _c % 8
	_ffc._b = uint8(_gg)
	if _ec := _fag - _ffc._bc; _ec > 0 && len(_ffc._eed)-1 < _fag {
		if _ffc._b != 0 {
			_ec++
		}
		_ffc.expandIfNeeded(_ec)
	}
	_ffc._bc = _fag
	return nil
}
func (_ed *BufferedWriter) grow(_eef int) {
	if _ed._eed == nil && _eef < _ff {
		_ed._eed = make([]byte, _eef, _ff)
		return
	}
	_gge := len(_ed._eed)
	if _ed._b != 0 {
		_gge++
	}
	_gd := cap(_ed._eed)
	switch {
	case _eef <= _gd/2-_gge:
		_a.Log.Trace("\u005b\u0042\u0075\u0066\u0066\u0065r\u0065\u0064\u0057\u0072\u0069t\u0065\u0072\u005d\u0020\u0067\u0072o\u0077\u0020\u002d\u0020\u0072e\u0073\u006c\u0069\u0063\u0065\u0020\u006f\u006e\u006c\u0079\u002e\u0020L\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0043\u0061\u0070\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u006e\u003a\u0020'\u0025\u0064\u0027", len(_ed._eed), cap(_ed._eed), _eef)
		_a.Log.Trace("\u0020\u006e\u0020\u003c\u003d\u0020\u0063\u0020\u002f\u0020\u0032\u0020\u002d\u006d\u002e \u0043:\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u006d\u003a\u0020\u0027\u0025\u0064\u0027", _gd, _gge)
		copy(_ed._eed, _ed._eed[_ed.fullOffset():])
	case _gd > _ae-_gd-_eef:
		_a.Log.Error("\u0042\u0055F\u0046\u0045\u0052 \u0074\u006f\u006f\u0020\u006c\u0061\u0072\u0067\u0065")
		return
	default:
		_ggc := make([]byte, 2*_gd+_eef)
		copy(_ggc, _ed._eed)
		_ed._eed = _ggc
	}
	_ed._eed = _ed._eed[:_gge+_eef]
}
func (_agd *BufferedWriter) tryGrowByReslice(_dff int) bool {
	if _ebg := len(_agd._eed); _dff <= cap(_agd._eed)-_ebg {
		_agd._eed = _agd._eed[:_ebg+_dff]
		return true
	}
	return false
}
func (_egf *Reader) Reset() {
	_egf._ce = _egf._db
	_egf._faa = _egf._cdf
	_egf._bab = _egf._ded
	_egf._ffd = _egf._bae
}
func (_bff *Writer) Write(p []byte) (int, error) {
	if len(p) > _bff.byteCapacity() {
		return 0, _f.EOF
	}
	for _, _aga := range p {
		if _geb := _bff.writeByte(_aga); _geb != nil {
			return 0, _geb
		}
	}
	return len(p), nil
}
func (_cd *BufferedWriter) writeFullBytes(_ga []byte) int {
	_bfb := copy(_cd._eed[_cd.fullOffset():], _ga)
	_cd._bc += _bfb
	return _bfb
}
func (_baed *SubstreamReader) fillBuffer() error {
	if uint64(_baed._ebgf.StreamPosition()) != _baed._be+_baed._aff {
		_, _fea := _baed._ebgf.Seek(int64(_baed._be+_baed._aff), _f.SeekStart)
		if _fea != nil {
			return _fea
		}
	}
	_baed._cdff = _baed._be
	_ecbg := _gbgc(uint64(len(_baed._da)), _baed._ace-_baed._be)
	_egg := make([]byte, _ecbg)
	_bfa, _edge := _baed._ebgf.Read(_egg)
	if _edge != nil {
		return _edge
	}
	for _bfdc := uint64(0); _bfdc < _ecbg; _bfdc++ {
		_baed._da[_bfdc] = _egg[_bfdc]
	}
	_baed._dcf = _baed._cdff + uint64(_bfa)
	return nil
}
func (_fbg *SubstreamReader) Offset() uint64        { return _fbg._aff }
func (_ceb *SubstreamReader) StreamPosition() int64 { return int64(_ceb._be) }
func (_dgc *Writer) FinishByte() {
	if _dgc._dbe == 0 {
		return
	}
	_dgc._dbe = 0
	_dgc._bed++
}
func NewWriter(data []byte) *Writer { return &Writer{_cde: data} }
func (_fdg *SubstreamReader) ReadUint32() (uint32, error) {
	_abc := make([]byte, 4)
	_, _gcd := _fdg.Read(_abc)
	if _gcd != nil {
		return 0, _gcd
	}
	return _g.BigEndian.Uint32(_abc), nil
}
func (_bad *SubstreamReader) ReadBits(n byte) (_acd uint64, _gdgg error) {
	if n < _bad._aaf {
		_bea := _bad._aaf - n
		_acd = uint64(_bad._eae >> _bea)
		_bad._eae &= 1<<_bea - 1
		_bad._aaf = _bea
		return _acd, nil
	}
	if n > _bad._aaf {
		if _bad._aaf > 0 {
			_acd = uint64(_bad._eae)
			n -= _bad._aaf
		}
		var _dfe byte
		for n >= 8 {
			_dfe, _gdgg = _bad.readBufferByte()
			if _gdgg != nil {
				return 0, _gdgg
			}
			_acd = _acd<<8 + uint64(_dfe)
			n -= 8
		}
		if n > 0 {
			if _bad._eae, _gdgg = _bad.readBufferByte(); _gdgg != nil {
				return 0, _gdgg
			}
			_gdd := 8 - n
			_acd = _acd<<n + uint64(_bad._eae>>_gdd)
			_bad._eae &= 1<<_gdd - 1
			_bad._aaf = _gdd
		} else {
			_bad._aaf = 0
		}
		return _acd, nil
	}
	_bad._aaf = 0
	return uint64(_bad._eae), nil
}
func (_ac *BufferedWriter) Write(d []byte) (int, error) {
	_ac.expandIfNeeded(len(d))
	if _ac._b == 0 {
		return _ac.writeFullBytes(d), nil
	}
	return _ac.writeShiftedBytes(d), nil
}

type BinaryWriter interface {
	BitWriter
	_f.Writer
	_f.ByteWriter
	Data() []byte
}
