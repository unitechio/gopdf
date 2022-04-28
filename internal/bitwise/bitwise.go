package bitwise

import (
	_f "encoding/binary"
	_gg "errors"
	_b "io"

	_c "bitbucket.org/shenghui0779/gopdf/common"
	_d "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

type SubstreamReader struct {
	_ebe  uint64
	_gdca StreamReader
	_efe  uint64
	_daa  uint64
	_eaf  []byte
	_dba  uint64
	_fcg  uint64
	_gbdf byte
	_abdd byte
	_ebb  uint64
	_dca  byte
}

func (_ggbe *Reader) ReadBit() (_gfe int, _ddd error) {
	_df, _ddd := _ggbe.readBool()
	if _ddd != nil {
		return 0, _ddd
	}
	if _df {
		_gfe = 1
	}
	return _gfe, nil
}
func (_cge *Writer) byteCapacity() int {
	_efc := len(_cge._acbb) - _cge._aeab
	if _cge._dce != 0 {
		_efc--
	}
	return _efc
}
func (_dfe *SubstreamReader) Mark() { _dfe._ebb = _dfe._ebe; _dfe._dca = _dfe._abdd }
func (_cgd *BufferedWriter) writeShiftedBytes(_ede []byte) int {
	for _, _fdf := range _ede {
		_cgd.writeByte(_fdf)
	}
	return len(_ede)
}

var _ _b.Writer = &BufferedWriter{}

func (_ggb *Reader) Length() uint64        { return uint64(len(_ggb._abf)) }
func (_aa *BufferedWriter) ResetBitIndex() { _aa._fc = 0 }
func (_ecg *Writer) Data() []byte          { return _ecg._acbb }
func (_aaa *Reader) Read(p []byte) (_cbf int, _bgd error) {
	if _aaa._ec == 0 {
		return _aaa.read(p)
	}
	for ; _cbf < len(p); _cbf++ {
		if p[_cbf], _bgd = _aaa.readUnalignedByte(); _bgd != nil {
			return 0, _bgd
		}
	}
	return _cbf, nil
}
func (_gf *BufferedWriter) WriteBit(bit int) error {
	if bit != 1 && bit != 0 {
		return _d.Errorf("\u0042\u0075\u0066fe\u0072\u0065\u0064\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074", "\u0062\u0069\u0074\u0020\u0076\u0061\u006cu\u0065\u0020\u006du\u0073\u0074\u0020\u0062e\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0030\u002c\u0031\u007d\u0020\u0062\u0075\u0074\u0020\u0069\u0073\u003a\u0020\u0025\u0064", bit)
	}
	if len(_gf._a)-1 < _gf._bb {
		_gf.expandIfNeeded(1)
	}
	_ea := _gf._fc
	if _gf._db {
		_ea = 7 - _gf._fc
	}
	_gf._a[_gf._bb] |= byte(uint16(bit<<_ea) & 0xff)
	_gf._fc++
	if _gf._fc == 8 {
		_gf._bb++
		_gf._fc = 0
	}
	return nil
}
func (_cgdd *SubstreamReader) ReadBit() (_bfe int, _fgg error) {
	_gfb, _fgg := _cgdd.readBool()
	if _fgg != nil {
		return 0, _fgg
	}
	if _gfb {
		_bfe = 1
	}
	return _bfe, nil
}
func (_aea *Reader) Mark() {
	_aea._bd = _aea._cda
	_aea._cdd = _aea._ec
	_aea._eegc = _aea._gbd
	_aea._bcb = _aea._acd
}
func (_dc *Reader) Align() (_bcg byte) {
	_bcg = _dc._ec
	_dc._ec = 0
	return _bcg
}
func (_bfg *SubstreamReader) ReadBool() (bool, error) { return _bfg.readBool() }
func (_bfb *Writer) UseMSB() bool                     { return _bfb._dbaae }

var (
	_ _b.Reader     = &Reader{}
	_ _b.ByteReader = &Reader{}
	_ _b.Seeker     = &Reader{}
	_ StreamReader  = &Reader{}
)

type BinaryWriter interface {
	BitWriter
	_b.Writer
	_b.ByteWriter
	Data() []byte
}

func (_gce *SubstreamReader) ReadBits(n byte) (_fgf uint64, _cba error) {
	if n < _gce._abdd {
		_cdb := _gce._abdd - n
		_fgf = uint64(_gce._gbdf >> _cdb)
		_gce._gbdf &= 1<<_cdb - 1
		_gce._abdd = _cdb
		return _fgf, nil
	}
	if n > _gce._abdd {
		if _gce._abdd > 0 {
			_fgf = uint64(_gce._gbdf)
			n -= _gce._abdd
		}
		var _aafe byte
		for n >= 8 {
			_aafe, _cba = _gce.readBufferByte()
			if _cba != nil {
				return 0, _cba
			}
			_fgf = _fgf<<8 + uint64(_aafe)
			n -= 8
		}
		if n > 0 {
			if _gce._gbdf, _cba = _gce.readBufferByte(); _cba != nil {
				return 0, _cba
			}
			_ebc := 8 - n
			_fgf = _fgf<<n + uint64(_gce._gbdf>>_ebc)
			_gce._gbdf &= 1<<_ebc - 1
			_gce._abdd = _ebc
		} else {
			_gce._abdd = 0
		}
		return _fgf, nil
	}
	_gce._abdd = 0
	return uint64(_gce._gbdf), nil
}
func (_dcb *SubstreamReader) readUnalignedByte() (_dad byte, _aad error) {
	_edc := _dcb._abdd
	_dad = _dcb._gbdf << (8 - _edc)
	_dcb._gbdf, _aad = _dcb.readBufferByte()
	if _aad != nil {
		return 0, _aad
	}
	_dad |= _dcb._gbdf >> _edc
	_dcb._gbdf &= 1<<_edc - 1
	return _dad, nil
}
func (_aaab *SubstreamReader) Reset() { _aaab._ebe = _aaab._ebb; _aaab._abdd = _aaab._dca }
func (_fe *Reader) Seek(offset int64, whence int) (int64, error) {
	_fe._fab = -1
	var _bce int64
	switch whence {
	case _b.SeekStart:
		_bce = offset
	case _b.SeekCurrent:
		_bce = _fe._cda + offset
	case _b.SeekEnd:
		_bce = int64(len(_fe._abf)) + offset
	default:
		return 0, _gg.New("\u0072\u0065\u0061de\u0072\u002e\u0052\u0065\u0061\u0064\u0065\u0072\u002eS\u0065e\u006b:\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0068\u0065\u006e\u0063\u0065")
	}
	if _bce < 0 {
		return 0, _gg.New("\u0072\u0065a\u0064\u0065\u0072\u002eR\u0065\u0061d\u0065\u0072\u002e\u0053\u0065\u0065\u006b\u003a \u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e")
	}
	_fe._cda = _bce
	_fe._ec = 0
	return _bce, nil
}
func NewWriterMSB(data []byte) *Writer           { return &Writer{_acbb: data, _dbaae: true} }
func (_ceg *SubstreamReader) Align() (_agc byte) { _agc = _ceg._abdd; _ceg._abdd = 0; return _agc }
func (_gec *BufferedWriter) byteCapacity() int {
	_eg := len(_gec._a) - _gec._bb
	if _gec._fc != 0 {
		_eg--
	}
	return _eg
}
func (_gfae *SubstreamReader) ReadUint32() (uint32, error) {
	_gac := make([]byte, 4)
	_, _eee := _gfae.Read(_gac)
	if _eee != nil {
		return 0, _eee
	}
	return _f.BigEndian.Uint32(_gac), nil
}
func (_dd *Reader) BitPosition() int { return int(_dd._ec) }
func (_bf *BufferedWriter) SkipBits(skip int) error {
	if skip == 0 {
		return nil
	}
	_ged := int(_bf._fc) + skip
	if _ged >= 0 && _ged < 8 {
		_bf._fc = uint8(_ged)
		return nil
	}
	_ged = int(_bf._fc) + _bf._bb*8 + skip
	if _ged < 0 {
		return _d.Errorf("\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073", "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_gdg := _ged / 8
	_fcd := _ged % 8
	_bf._fc = uint8(_fcd)
	if _ac := _gdg - _bf._bb; _ac > 0 && len(_bf._a)-1 < _gdg {
		if _bf._fc != 0 {
			_ac++
		}
		_bf.expandIfNeeded(_ac)
	}
	_bf._bb = _gdg
	return nil
}
func (_bceg *SubstreamReader) fillBuffer() error {
	if uint64(_bceg._gdca.StreamPosition()) != _bceg._ebe+_bceg._efe {
		_, _ded := _bceg._gdca.Seek(int64(_bceg._ebe+_bceg._efe), _b.SeekStart)
		if _ded != nil {
			return _ded
		}
	}
	_bceg._dba = _bceg._ebe
	_cgc := _ece(uint64(len(_bceg._eaf)), _bceg._daa-_bceg._ebe)
	_fbac := make([]byte, _cgc)
	_fbc, _aebb := _bceg._gdca.Read(_fbac)
	if _aebb != nil {
		return _aebb
	}
	for _dfeb := uint64(0); _dfeb < _cgc; _dfeb++ {
		_bceg._eaf[_dfeb] = _fbac[_dfeb]
	}
	_bceg._fcg = _bceg._dba + uint64(_fbc)
	return nil
}
func (_geca *SubstreamReader) Length() uint64 { return _geca._daa }
func (_ggg *Writer) writeBit(_bgb uint8) error {
	if len(_ggg._acbb)-1 < _ggg._aeab {
		return _b.EOF
	}
	_gcg := _ggg._dce
	if _ggg._dbaae {
		_gcg = 7 - _ggg._dce
	}
	_ggg._acbb[_ggg._aeab] |= byte(uint16(_bgb<<_gcg) & 0xff)
	_ggg._dce++
	if _ggg._dce == 8 {
		_ggg._aeab++
		_ggg._dce = 0
	}
	return nil
}
func (_ceea *Writer) ResetBit() { _ceea._dce = 0 }
func (_ef *BufferedWriter) writeByte(_ga byte) {
	switch {
	case _ef._fc == 0:
		_ef._a[_ef._bb] = _ga
		_ef._bb++
	case _ef._db:
		_ef._a[_ef._bb] |= _ga >> _ef._fc
		_ef._bb++
		_ef._a[_ef._bb] = byte(uint16(_ga) << (8 - _ef._fc) & 0xff)
	default:
		_ef._a[_ef._bb] |= byte(uint16(_ga) << _ef._fc & 0xff)
		_ef._bb++
		_ef._a[_ef._bb] = _ga >> (8 - _ef._fc)
	}
}
func (_af *BufferedWriter) WriteBits(bits uint64, number int) (_gdf int, _ag error) {
	const _eeg = "\u0042u\u0066\u0066\u0065\u0072e\u0064\u0057\u0072\u0069\u0074e\u0072.\u0057r\u0069\u0074\u0065\u0072\u0042\u0069\u0074s"
	if number < 0 || number > 64 {
		return 0, _d.Errorf(_eeg, "\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", number)
	}
	_ed := number / 8
	if _ed > 0 {
		_fd := number - _ed*8
		for _aae := _ed - 1; _aae >= 0; _aae-- {
			_cee := byte((bits >> uint(_aae*8+_fd)) & 0xff)
			if _ag = _af.WriteByte(_cee); _ag != nil {
				return _gdf, _d.Wrapf(_ag, _eeg, "\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _ed-_aae+1)
			}
		}
		number -= _ed * 8
		if number == 0 {
			return _ed, nil
		}
	}
	var _ba int
	for _fa := 0; _fa < number; _fa++ {
		if _af._db {
			_ba = int((bits >> uint(number-1-_fa)) & 0x1)
		} else {
			_ba = int(bits & 0x1)
			bits >>= 1
		}
		if _ag = _af.WriteBit(_ba); _ag != nil {
			return _gdf, _d.Wrapf(_ag, _eeg, "\u0062i\u0074\u003a\u0020\u0025\u0064", _fa)
		}
	}
	return _ed, nil
}
func (_eb *Reader) ReadUint32() (uint32, error) {
	_gc := make([]byte, 4)
	_, _abd := _eb.Read(_gc)
	if _abd != nil {
		return 0, _abd
	}
	return _f.BigEndian.Uint32(_gc), nil
}
func (_bbg *SubstreamReader) ReadByte() (byte, error) {
	if _bbg._abdd == 0 {
		return _bbg.readBufferByte()
	}
	return _bbg.readUnalignedByte()
}
func (_adc *SubstreamReader) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case _b.SeekStart:
		_adc._ebe = uint64(offset)
	case _b.SeekCurrent:
		_adc._ebe += uint64(offset)
	case _b.SeekEnd:
		_adc._ebe = _adc._daa + uint64(offset)
	default:
		return 0, _gg.New("\u0072\u0065\u0061d\u0065\u0072\u002e\u0053\u0075\u0062\u0073\u0074\u0072\u0065\u0061\u006d\u0052\u0065\u0061\u0064\u0065\u0072\u002e\u0053\u0065\u0065\u006b\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0077\u0068\u0065\u006e\u0063\u0065")
	}
	_adc._abdd = 0
	return int64(_adc._ebe), nil
}

var _ BinaryWriter = &Writer{}

func (_gba *Reader) readUnalignedByte() (_dag byte, _gfa error) {
	_cdf := _gba._ec
	_dag = _gba._gbd << (8 - _cdf)
	_gba._gbd, _gfa = _gba.readBufferByte()
	if _gfa != nil {
		return 0, _gfa
	}
	_dag |= _gba._gbd >> _cdf
	_gba._gbd &= 1<<_cdf - 1
	return _dag, nil
}
func BufferedMSB() *BufferedWriter { return &BufferedWriter{_db: true} }
func (_bab *Reader) read(_dgaa []byte) (int, error) {
	if _bab._cda >= int64(len(_bab._abf)) {
		return 0, _b.EOF
	}
	_bab._fab = -1
	_bcgb := copy(_dgaa, _bab._abf[_bab._cda:])
	_bab._cda += int64(_bcgb)
	return _bcgb, nil
}
func NewWriter(data []byte) *Writer { return &Writer{_acbb: data} }
func (_cfa *BufferedWriter) WriteByte(bt byte) error {
	if _cfa._bb > len(_cfa._a)-1 || (_cfa._bb == len(_cfa._a)-1 && _cfa._fc != 0) {
		_cfa.expandIfNeeded(1)
	}
	_cfa.writeByte(bt)
	return nil
}
func _ece(_afc, _gceb uint64) uint64 {
	if _afc < _gceb {
		return _afc
	}
	return _gceb
}
func (_fba *Reader) readBool() (_afe bool, _aaf error) {
	if _fba._ec == 0 {
		_fba._gbd, _aaf = _fba.readBufferByte()
		if _aaf != nil {
			return false, _aaf
		}
		_afe = (_fba._gbd & 0x80) != 0
		_fba._gbd, _fba._ec = _fba._gbd&0x7f, 7
		return _afe, nil
	}
	_fba._ec--
	_afe = (_fba._gbd & (1 << _fba._ec)) != 0
	_fba._gbd &= 1<<_fba._ec - 1
	return _afe, nil
}

type BufferedWriter struct {
	_a  []byte
	_fc uint8
	_bb int
	_db bool
}

func (_fdg *Reader) ReadBool() (bool, error) { return _fdg.readBool() }
func (_gcd *Writer) writeByte(_gcf byte) error {
	if _gcd._aeab > len(_gcd._acbb)-1 {
		return _b.EOF
	}
	if _gcd._aeab == len(_gcd._acbb)-1 && _gcd._dce != 0 {
		return _b.EOF
	}
	if _gcd._dce == 0 {
		_gcd._acbb[_gcd._aeab] = _gcf
		_gcd._aeab++
		return nil
	}
	if _gcd._dbaae {
		_gcd._acbb[_gcd._aeab] |= _gcf >> _gcd._dce
		_gcd._aeab++
		_gcd._acbb[_gcd._aeab] = byte(uint16(_gcf) << (8 - _gcd._dce) & 0xff)
	} else {
		_gcd._acbb[_gcd._aeab] |= byte(uint16(_gcf) << _gcd._dce & 0xff)
		_gcd._aeab++
		_gcd._acbb[_gcd._aeab] = _gcf >> (8 - _gcd._dce)
	}
	return nil
}
func (_bdb *SubstreamReader) BitPosition() int { return int(_bdb._abdd) }

type Writer struct {
	_acbb  []byte
	_dce   uint8
	_aeab  int
	_dbaae bool
}

func (_ee *BufferedWriter) Write(d []byte) (int, error) {
	_ee.expandIfNeeded(len(d))
	if _ee._fc == 0 {
		return _ee.writeFullBytes(d), nil
	}
	return _ee.writeShiftedBytes(d), nil
}
func (_cb *BufferedWriter) expandIfNeeded(_eed int) {
	if !_cb.tryGrowByReslice(_eed) {
		_cb.grow(_eed)
	}
}
func (_da *BufferedWriter) Data() []byte { return _da._a }

type Reader struct {
	_abf  []byte
	_gbd  byte
	_ec   byte
	_cda  int64
	_acd  int
	_fab  int
	_bd   int64
	_cdd  byte
	_eegc byte
	_bcb  int
}

func (_edg *Writer) WriteBits(bits uint64, number int) (_dgaag int, _agg error) {
	const _gcaa = "\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065r\u0042\u0069\u0074\u0073"
	if number < 0 || number > 64 {
		return 0, _d.Errorf(_gcaa, "\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", number)
	}
	if number == 0 {
		return 0, nil
	}
	_gef := number / 8
	if _gef > 0 {
		_bff := number - _gef*8
		for _ddc := _gef - 1; _ddc >= 0; _ddc-- {
			_bdc := byte((bits >> uint(_ddc*8+_bff)) & 0xff)
			if _agg = _edg.WriteByte(_bdc); _agg != nil {
				return _dgaag, _d.Wrapf(_agg, _gcaa, "\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _gef-_ddc+1)
			}
		}
		number -= _gef * 8
		if number == 0 {
			return _gef, nil
		}
	}
	var _eeb int
	for _bdge := 0; _bdge < number; _bdge++ {
		if _edg._dbaae {
			_eeb = int((bits >> uint(number-1-_bdge)) & 0x1)
		} else {
			_eeb = int(bits & 0x1)
			bits >>= 1
		}
		if _agg = _edg.WriteBit(_eeb); _agg != nil {
			return _dgaag, _d.Wrapf(_agg, _gcaa, "\u0062i\u0074\u003a\u0020\u0025\u0064", _bdge)
		}
	}
	return _gef, nil
}
func (_baab *Writer) Write(p []byte) (int, error) {
	if len(p) > _baab.byteCapacity() {
		return 0, _b.EOF
	}
	for _, _gae := range p {
		if _ebd := _baab.writeByte(_gae); _ebd != nil {
			return 0, _ebd
		}
	}
	return len(p), nil
}
func (_dge *BufferedWriter) tryGrowByReslice(_cef int) bool {
	if _bc := len(_dge._a); _cef <= cap(_dge._a)-_bc {
		_dge._a = _dge._a[:_bc+_cef]
		return true
	}
	return false
}
func (_bdg *Reader) StreamPosition() int64 { return _bdg._cda }
func NewSubstreamReader(r StreamReader, offset, length uint64) (*SubstreamReader, error) {
	if r == nil {
		return nil, _gg.New("\u0072o\u006ft\u0020\u0072\u0065\u0061\u0064e\u0072\u0020i\u0073\u0020\u006e\u0069\u006c")
	}
	_c.Log.Trace("\u004e\u0065\u0077\u0053\u0075\u0062\u0073\u0074r\u0065\u0061\u006dRe\u0061\u0064\u0065\u0072\u0020\u0061t\u0020\u006f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0064\u0020\u0077\u0069\u0074h\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a \u0025\u0064", offset, length)
	return &SubstreamReader{_gdca: r, _efe: offset, _daa: length, _eaf: make([]byte, length)}, nil
}

var _ BinaryWriter = &BufferedWriter{}

type BitWriter interface {
	WriteBit(_ae int) error
	WriteBits(_dga uint64, _baa int) (_cec int, _ab error)
	FinishByte()
	SkipBits(_baag int) error
}

func (_bcbg *SubstreamReader) readBool() (_daf bool, _be error) {
	if _bcbg._abdd == 0 {
		_bcbg._gbdf, _be = _bcbg.readBufferByte()
		if _be != nil {
			return false, _be
		}
		_daf = (_bcbg._gbdf & 0x80) != 0
		_bcbg._gbdf, _bcbg._abdd = _bcbg._gbdf&0x7f, 7
		return _daf, nil
	}
	_bcbg._abdd--
	_daf = (_bcbg._gbdf & (1 << _bcbg._abdd)) != 0
	_bcbg._gbdf &= 1<<_bcbg._abdd - 1
	return _daf, nil
}
func (_fdfd *Reader) ReadByte() (byte, error) {
	if _fdfd._ec == 0 {
		return _fdfd.readBufferByte()
	}
	return _fdfd.readUnalignedByte()
}
func (_caf *SubstreamReader) readBufferByte() (byte, error) {
	if _caf._ebe >= _caf._daa {
		return 0, _b.EOF
	}
	if _caf._ebe >= _caf._fcg || _caf._ebe < _caf._dba {
		if _bgde := _caf.fillBuffer(); _bgde != nil {
			return 0, _bgde
		}
	}
	_cbd := _caf._eaf[_caf._ebe-_caf._dba]
	_caf._ebe++
	return _cbd, nil
}
func (_dg *BufferedWriter) grow(_de int) {
	if _dg._a == nil && _de < _cf {
		_dg._a = make([]byte, _de, _cf)
		return
	}
	_gedd := len(_dg._a)
	if _dg._fc != 0 {
		_gedd++
	}
	_aab := cap(_dg._a)
	switch {
	case _de <= _aab/2-_gedd:
		_c.Log.Trace("\u005b\u0042\u0075\u0066\u0066\u0065r\u0065\u0064\u0057\u0072\u0069t\u0065\u0072\u005d\u0020\u0067\u0072o\u0077\u0020\u002d\u0020\u0072e\u0073\u006c\u0069\u0063\u0065\u0020\u006f\u006e\u006c\u0079\u002e\u0020L\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0043\u0061\u0070\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u006e\u003a\u0020'\u0025\u0064\u0027", len(_dg._a), cap(_dg._a), _de)
		_c.Log.Trace("\u0020\u006e\u0020\u003c\u003d\u0020\u0063\u0020\u002f\u0020\u0032\u0020\u002d\u006d\u002e \u0043:\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u006d\u003a\u0020\u0027\u0025\u0064\u0027", _aab, _gedd)
		copy(_dg._a, _dg._a[_dg.fullOffset():])
	case _aab > _gb-_aab-_de:
		_c.Log.Error("\u0042\u0055F\u0046\u0045\u0052 \u0074\u006f\u006f\u0020\u006c\u0061\u0072\u0067\u0065")
		return
	default:
		_ad := make([]byte, 2*_aab+_de)
		copy(_ad, _dg._a)
		_dg._a = _ad
	}
	_dg._a = _dg._a[:_gedd+_de]
}
func (_cdg *SubstreamReader) Read(b []byte) (_acb int, _fag error) {
	if _cdg._ebe >= _cdg._daa {
		_c.Log.Trace("\u0053\u0074\u0072e\u0061\u006d\u0050\u006fs\u003a\u0020\u0027\u0025\u0064\u0027\u0020>\u003d\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027", _cdg._ebe, _cdg._daa)
		return 0, _b.EOF
	}
	for ; _acb < len(b); _acb++ {
		if b[_acb], _fag = _cdg.readUnalignedByte(); _fag != nil {
			if _fag == _b.EOF {
				return _acb, nil
			}
			return 0, _fag
		}
	}
	return _acb, nil
}
func (_dbaa *SubstreamReader) StreamPosition() int64 { return int64(_dbaa._ebe) }
func (_cdad *Writer) FinishByte() {
	if _cdad._dce == 0 {
		return
	}
	_cdad._dce = 0
	_cdad._aeab++
}
func (_dafb *Writer) SkipBits(skip int) error {
	const _ffd = "\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073"
	if skip == 0 {
		return nil
	}
	_fggf := int(_dafb._dce) + skip
	if _fggf >= 0 && _fggf < 8 {
		_dafb._dce = uint8(_fggf)
		return nil
	}
	_fggf = int(_dafb._dce) + _dafb._aeab*8 + skip
	if _fggf < 0 {
		return _d.Errorf(_ffd, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_dcc := _fggf / 8
	_ecf := _fggf % 8
	_c.Log.Trace("\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073")
	_c.Log.Trace("\u0042\u0069\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0042\u0079\u0074\u0065\u0049n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0046\u0075\u006c\u006c\u0042\u0069\u0074\u0073\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u004c\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027,\u0020\u0043\u0061p\u003a\u0020\u0027\u0025\u0064\u0027", _dafb._dce, _dafb._aeab, int(_dafb._dce)+(_dafb._aeab)*8, len(_dafb._acbb), cap(_dafb._acbb))
	_c.Log.Trace("S\u006b\u0069\u0070\u003a\u0020\u0027%\u0064\u0027\u002c\u0020\u0064\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062i\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025d\u0027", skip, _fggf, _ecf)
	_dafb._dce = uint8(_ecf)
	if _aceb := _dcc - _dafb._aeab; _aceb > 0 && len(_dafb._acbb)-1 < _dcc {
		_c.Log.Trace("\u0042\u0079\u0074e\u0044\u0069\u0066\u0066\u003a\u0020\u0025\u0064", _aceb)
		return _d.Errorf(_ffd, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_dafb._aeab = _dcc
	_c.Log.Trace("\u0042\u0069\u0074I\u006e\u0064\u0065\u0078:\u0020\u0027\u0025\u0064\u0027\u002c\u0020B\u0079\u0074\u0065\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027", _dafb._dce, _dafb._aeab)
	return nil
}
func (_eef *BufferedWriter) fullOffset() int {
	_cg := _eef._bb
	if _eef._fc != 0 {
		_cg++
	}
	return _cg
}
func NewReader(data []byte) *Reader { return &Reader{_abf: data} }

var _ _b.ByteWriter = &BufferedWriter{}

func (_fcec *Writer) WriteByte(c byte) error { return _fcec.writeByte(c) }
func (_cd *BufferedWriter) FinishByte() {
	if _cd._fc == 0 {
		return
	}
	_cd._fc = 0
	_cd._bb++
}
func (_e *BufferedWriter) Len() int { return _e.byteCapacity() }
func (_dgd *Reader) ConsumeRemainingBits() (uint64, error) {
	if _dgd._ec != 0 {
		return _dgd.ReadBits(_dgd._ec)
	}
	return 0, nil
}
func (_aff *SubstreamReader) Offset() uint64 { return _aff._efe }
func (_egf *Reader) Reset() {
	_egf._cda = _egf._bd
	_egf._ec = _egf._cdd
	_egf._gbd = _egf._eegc
	_egf._acd = _egf._bcb
}

const (
	_cf = 64
	_gb = int(^uint(0) >> 1)
)

func (_dee *Writer) WriteBit(bit int) error {
	switch bit {
	case 0, 1:
		return _dee.writeBit(uint8(bit))
	}
	return _d.Error("\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074", "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0062\u0069\u0074\u0020v\u0061\u006c\u0075\u0065")
}
func (_aefb *Reader) ReadBits(n byte) (_bcc uint64, _dgdd error) {
	if n < _aefb._ec {
		_dfa := _aefb._ec - n
		_bcc = uint64(_aefb._gbd >> _dfa)
		_aefb._gbd &= 1<<_dfa - 1
		_aefb._ec = _dfa
		return _bcc, nil
	}
	if n > _aefb._ec {
		if _aefb._ec > 0 {
			_bcc = uint64(_aefb._gbd)
			n -= _aefb._ec
		}
		for n >= 8 {
			_ff, _fg := _aefb.readBufferByte()
			if _fg != nil {
				return 0, _fg
			}
			_bcc = _bcc<<8 + uint64(_ff)
			n -= 8
		}
		if n > 0 {
			if _aefb._gbd, _dgdd = _aefb.readBufferByte(); _dgdd != nil {
				return 0, _dgdd
			}
			_aeb := 8 - n
			_bcc = _bcc<<n + uint64(_aefb._gbd>>_aeb)
			_aefb._gbd &= 1<<_aeb - 1
			_aefb._ec = _aeb
		} else {
			_aefb._ec = 0
		}
		return _bcc, nil
	}
	_aefb._ec = 0
	return uint64(_aefb._gbd), nil
}
func (_cdc *BufferedWriter) writeFullBytes(_efd []byte) int {
	_fb := copy(_cdc._a[_cdc.fullOffset():], _efd)
	_cdc._bb += _fb
	return _fb
}
func (_ge *BufferedWriter) Reset() {
	_ge._a = _ge._a[:0]
	_ge._bb = 0
	_ge._fc = 0
}
func (_gdc *Reader) readBufferByte() (byte, error) {
	if _gdc._cda >= int64(len(_gdc._abf)) {
		return 0, _b.EOF
	}
	_gdc._fab = -1
	_fce := _gdc._abf[_gdc._cda]
	_gdc._cda++
	_gdc._acd = int(_fce)
	return _fce, nil
}

type StreamReader interface {
	_b.Reader
	_b.ByteReader
	_b.Seeker
	Align() byte
	BitPosition() int
	Mark()
	Length() uint64
	ReadBit() (int, error)
	ReadBits(_cc byte) (uint64, error)
	ReadBool() (bool, error)
	ReadUint32() (uint32, error)
	Reset()
	StreamPosition() int64
}
