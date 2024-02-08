package bitwise

import (
	_dd "encoding/binary"
	_g "errors"
	_ff "fmt"
	_f "io"

	_b "bitbucket.org/shenghui0779/gopdf/common"
	_c "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

func (_bfe *BufferedWriter) byteCapacity() int {
	_dad := len(_bfe._fg) - _bfe._a
	if _bfe._fc != 0 {
		_dad--
	}
	return _dad
}
func (_ec *BufferedWriter) Data() []byte { return _ec._fg }
func (_fd *Writer) SkipBits(skip int) error {
	const _bfb = "\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073"
	if skip == 0 {
		return nil
	}
	_fcf := int(_fd._gefd) + skip
	if _fcf >= 0 && _fcf < 8 {
		_fd._gefd = uint8(_fcf)
		return nil
	}
	_fcf = int(_fd._gefd) + _fd._aba*8 + skip
	if _fcf < 0 {
		return _c.Errorf(_bfb, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_bae := _fcf / 8
	_cefd := _fcf % 8
	_b.Log.Trace("\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073")
	_b.Log.Trace("\u0042\u0069\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0042\u0079\u0074\u0065\u0049n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0046\u0075\u006c\u006c\u0042\u0069\u0074\u0073\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u004c\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027,\u0020\u0043\u0061p\u003a\u0020\u0027\u0025\u0064\u0027", _fd._gefd, _fd._aba, int(_fd._gefd)+(_fd._aba)*8, len(_fd._bda), cap(_fd._bda))
	_b.Log.Trace("S\u006b\u0069\u0070\u003a\u0020\u0027%\u0064\u0027\u002c\u0020\u0064\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062i\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025d\u0027", skip, _fcf, _cefd)
	_fd._gefd = uint8(_cefd)
	if _bdf := _bae - _fd._aba; _bdf > 0 && len(_fd._bda)-1 < _bae {
		_b.Log.Trace("\u0042\u0079\u0074e\u0044\u0069\u0066\u0066\u003a\u0020\u0025\u0064", _bdf)
		return _c.Errorf(_bfb, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_fd._aba = _bae
	_b.Log.Trace("\u0042\u0069\u0074I\u006e\u0064\u0065\u0078:\u0020\u0027\u0025\u0064\u0027\u002c\u0020B\u0079\u0074\u0065\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027", _fd._gefd, _fd._aba)
	return nil
}
func (_ca *Reader) ReadUint32() (uint32, error) {
	_eea := make([]byte, 4)
	_, _abd := _ca.Read(_eea)
	if _abd != nil {
		return 0, _abd
	}
	return _dd.BigEndian.Uint32(_eea), nil
}
func (_bdg *Reader) read(_dbf []byte) (int, error) {
	if _bdg._add >= int64(_bdg._cef._bef) {
		return 0, _f.EOF
	}
	_bdg._cgb = -1
	_caa := copy(_dbf, _bdg._cef._ee[(int64(_bdg._cef._af)+_bdg._add):(_bdg._cef._af+_bdg._cef._bef)])
	_bdg._add += int64(_caa)
	return _caa, nil
}
func (_aef *Reader) readUnalignedByte() (_aed byte, _cca error) {
	_eec := _aef._ead
	_aed = _aef._ffd << (8 - _eec)
	_aef._ffd, _cca = _aef.readBufferByte()
	if _cca != nil {
		return 0, _cca
	}
	_aed |= _aef._ffd >> _eec
	_aef._ffd &= 1<<_eec - 1
	return _aed, nil
}
func (_bac *Writer) byteCapacity() int {
	_bcg := len(_bac._bda) - _bac._aba
	if _bac._gefd != 0 {
		_bcg--
	}
	return _bcg
}

type Writer struct {
	_bda  []byte
	_gefd uint8
	_aba  int
	_dda  bool
}

func (_fab *Reader) readBool() (_cdc bool, _cfg error) {
	if _fab._ead == 0 {
		_fab._ffd, _cfg = _fab.readBufferByte()
		if _cfg != nil {
			return false, _cfg
		}
		_cdc = (_fab._ffd & 0x80) != 0
		_fab._ffd, _fab._ead = _fab._ffd&0x7f, 7
		return _cdc, nil
	}
	_fab._ead--
	_cdc = (_fab._ffd & (1 << _fab._ead)) != 0
	_fab._ffd &= 1<<_fab._ead - 1
	return _cdc, nil
}
func (_bed *Writer) Write(p []byte) (int, error) {
	if len(p) > _bed.byteCapacity() {
		return 0, _f.EOF
	}
	for _, _abc := range p {
		if _aedf := _bed.writeByte(_abc); _aedf != nil {
			return 0, _aedf
		}
	}
	return len(p), nil
}
func (_faa *Writer) WriteBit(bit int) error {
	switch bit {
	case 0, 1:
		return _faa.writeBit(uint8(bit))
	}
	return _c.Error("\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074", "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0062\u0069\u0074\u0020v\u0061\u006c\u0075\u0065")
}
func (_dcd *Reader) Read(p []byte) (_gcf int, _bd error) {
	if _dcd._ead == 0 {
		return _dcd.read(p)
	}
	for ; _gcf < len(p); _gcf++ {
		if p[_gcf], _bd = _dcd.readUnalignedByte(); _bd != nil {
			return 0, _bd
		}
	}
	return _gcf, nil
}
func (_eg *Reader) readBufferByte() (byte, error) {
	if _eg._add >= int64(_eg._cef._bef) {
		return 0, _f.EOF
	}
	_eg._cgb = -1
	_bde := _eg._cef._ee[int64(_eg._cef._af)+_eg._add]
	_eg._add++
	_eg._ecg = int(_bde)
	return _bde, nil
}
func (_eb *BufferedWriter) writeByte(_fa byte) {
	switch {
	case _eb._fc == 0:
		_eb._fg[_eb._a] = _fa
		_eb._a++
	case _eb._fb:
		_eb._fg[_eb._a] |= _fa >> _eb._fc
		_eb._a++
		_eb._fg[_eb._a] = byte(uint16(_fa) << (8 - _eb._fc) & 0xff)
	default:
		_eb._fg[_eb._a] |= byte(uint16(_fa) << _eb._fc & 0xff)
		_eb._a++
		_eb._fg[_eb._a] = _fa >> (8 - _eb._fc)
	}
}
func (_edb *BufferedWriter) tryGrowByReslice(_bfa int) bool {
	if _gf := len(_edb._fg); _bfa <= cap(_edb._fg)-_gf {
		_edb._fg = _edb._fg[:_gf+_bfa]
		return true
	}
	return false
}
func (_ea *BufferedWriter) writeFullBytes(_fbf []byte) int {
	_fgg := copy(_ea._fg[_ea.fullOffset():], _fbf)
	_ea._a += _fgg
	return _fgg
}

type BufferedWriter struct {
	_fg []byte
	_fc uint8
	_a  int
	_fb bool
}

func (_fe *BufferedWriter) SkipBits(skip int) error {
	if skip == 0 {
		return nil
	}
	_ce := int(_fe._fc) + skip
	if _ce >= 0 && _ce < 8 {
		_fe._fc = uint8(_ce)
		return nil
	}
	_ce = int(_fe._fc) + _fe._a*8 + skip
	if _ce < 0 {
		return _c.Errorf("\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073", "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_da := _ce / 8
	_ac := _ce % 8
	_fe._fc = uint8(_ac)
	if _gg := _da - _fe._a; _gg > 0 && len(_fe._fg)-1 < _da {
		if _fe._fc != 0 {
			_gg++
		}
		_fe.expandIfNeeded(_gg)
	}
	_fe._a = _da
	return nil
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
	ReadBits(_ffb byte) (uint64, error)
	ReadBool() (bool, error)
	ReadUint32() (uint32, error)
	Reset()
	AbsolutePosition() int64
}

func (_ba *BufferedWriter) writeShiftedBytes(_edf []byte) int {
	for _, _ggc := range _edf {
		_ba.writeByte(_ggc)
	}
	return len(_edf)
}

var (
	_ _f.Reader     = &Reader{}
	_ _f.ByteReader = &Reader{}
	_ _f.Seeker     = &Reader{}
	_ StreamReader  = &Reader{}
)

func (_gdd *Writer) UseMSB() bool { return _gdd._dda }
func (_dc *Reader) NewPartialReader(offset, length int, relative bool) (*Reader, error) {
	if offset < 0 {
		return nil, _g.New("p\u0061\u0072\u0074\u0069\u0061\u006c\u0020\u0072\u0065\u0061\u0064\u0065\u0072\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062e \u006e\u0065\u0067a\u0074i\u0076\u0065")
	}
	if relative {
		offset = _dc._cef._af + offset
	}
	if length > 0 {
		_gga := len(_dc._cef._ee)
		if relative {
			_gga = _dc._cef._bef
		}
		if offset+length > _gga {
			return nil, _ff.Errorf("\u0070\u0061r\u0074\u0069\u0061l\u0020\u0072\u0065\u0061\u0064e\u0072\u0020\u006f\u0066\u0066se\u0074\u0028\u0025\u0064\u0029\u002b\u006c\u0065\u006e\u0067\u0074\u0068\u0028\u0025\u0064\u0029\u003d\u0025d\u0020i\u0073\u0020\u0067\u0072\u0065\u0061ter\u0020\u0074\u0068\u0061\u006e\u0020\u0074\u0068\u0065\u0020\u006f\u0072ig\u0069n\u0061\u006c\u0020\u0072e\u0061d\u0065r\u0020\u006ce\u006e\u0067th\u003a\u0020\u0025\u0064", offset, length, offset+length, _dc._cef._bef)
		}
	}
	if length < 0 {
		_bfag := len(_dc._cef._ee)
		if relative {
			_bfag = _dc._cef._bef
		}
		length = _bfag - offset
	}
	return &Reader{_cef: readerSource{_ee: _dc._cef._ee, _bef: length, _af: offset}}, nil
}

type BinaryWriter interface {
	BitWriter
	_f.Writer
	_f.ByteWriter
	Data() []byte
}

func (_be *BufferedWriter) ResetBitIndex() { _be._fc = 0 }
func (_fbfa *Reader) ConsumeRemainingBits() (uint64, error) {
	if _fbfa._ead != 0 {
		return _fbfa.ReadBits(_fbfa._ead)
	}
	return 0, nil
}

var _ BinaryWriter = &BufferedWriter{}

func (_fce *BufferedWriter) Reset() { _fce._fg = _fce._fg[:0]; _fce._a = 0; _fce._fc = 0 }

type readerSource struct {
	_ee  []byte
	_af  int
	_bef int
}

func (_gba *Writer) WriteByte(c byte) error { return _gba.writeByte(c) }
func (_ceg *Reader) BitPosition() int       { return int(_ceg._ead) }
func (_geb *Writer) ResetBit()              { _geb._gefd = 0 }
func (_fgb *BufferedWriter) Write(d []byte) (int, error) {
	_fgb.expandIfNeeded(len(d))
	if _fgb._fc == 0 {
		return _fgb.writeFullBytes(d), nil
	}
	return _fgb.writeShiftedBytes(d), nil
}
func (_ed *BufferedWriter) Len() int { return _ed.byteCapacity() }
func NewReader(data []byte) *Reader {
	return &Reader{_cef: readerSource{_ee: data, _bef: len(data), _af: 0}}
}

type Reader struct {
	_cef readerSource
	_ffd byte
	_ead byte
	_add int64
	_ecg int
	_cgb int
	_aa  int64
	_abf byte
	_ddd byte
	_ef  int
}

func (_eab *Reader) Seek(offset int64, whence int) (int64, error) {
	_eab._cgb = -1
	_eab._ead = 0
	_eab._ffd = 0
	_eab._ecg = 0
	var _aee int64
	switch whence {
	case _f.SeekStart:
		_aee = offset
	case _f.SeekCurrent:
		_aee = _eab._add + offset
	case _f.SeekEnd:
		_aee = int64(_eab._cef._bef) + offset
	default:
		return 0, _g.New("\u0072\u0065\u0061de\u0072\u002e\u0052\u0065\u0061\u0064\u0065\u0072\u002eS\u0065e\u006b:\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0068\u0065\u006e\u0063\u0065")
	}
	if _aee < 0 {
		return 0, _g.New("\u0072\u0065a\u0064\u0065\u0072\u002eR\u0065\u0061d\u0065\u0072\u002e\u0053\u0065\u0065\u006b\u003a \u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e")
	}
	_eab._add = _aee
	_eab._ead = 0
	return _aee, nil
}
func (_fcb *Reader) Align() (_gbd byte) { _gbd = _fcb._ead; _fcb._ead = 0; return _gbd }
func (_fgbg *BufferedWriter) fullOffset() int {
	_cd := _fgbg._a
	if _fgbg._fc != 0 {
		_cd++
	}
	return _cd
}
func (_gb *BufferedWriter) WriteBit(bit int) error {
	if bit != 1 && bit != 0 {
		return _c.Errorf("\u0042\u0075\u0066fe\u0072\u0065\u0064\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074", "\u0062\u0069\u0074\u0020\u0076\u0061\u006cu\u0065\u0020\u006du\u0073\u0074\u0020\u0062e\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0030\u002c\u0031\u007d\u0020\u0062\u0075\u0074\u0020\u0069\u0073\u003a\u0020\u0025\u0064", bit)
	}
	if len(_gb._fg)-1 < _gb._a {
		_gb.expandIfNeeded(1)
	}
	_bf := _gb._fc
	if _gb._fb {
		_bf = 7 - _gb._fc
	}
	_gb._fg[_gb._a] |= byte(uint16(bit<<_bf) & 0xff)
	_gb._fc++
	if _gb._fc == 8 {
		_gb._a++
		_gb._fc = 0
	}
	return nil
}
func (_ecf *Writer) WriteBits(bits uint64, number int) (_fggc int, _beg error) {
	const _ecgd = "\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065r\u0042\u0069\u0074\u0073"
	if number < 0 || number > 64 {
		return 0, _c.Errorf(_ecgd, "\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", number)
	}
	if number == 0 {
		return 0, nil
	}
	_edd := number / 8
	if _edd > 0 {
		_febf := number - _edd*8
		for _gfg := _edd - 1; _gfg >= 0; _gfg-- {
			_eba := byte((bits >> uint(_gfg*8+_febf)) & 0xff)
			if _beg = _ecf.WriteByte(_eba); _beg != nil {
				return _fggc, _c.Wrapf(_beg, _ecgd, "\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _edd-_gfg+1)
			}
		}
		number -= _edd * 8
		if number == 0 {
			return _edd, nil
		}
	}
	var _bgc int
	for _cb := 0; _cb < number; _cb++ {
		if _ecf._dda {
			_bgc = int((bits >> uint(number-1-_cb)) & 0x1)
		} else {
			_bgc = int(bits & 0x1)
			bits >>= 1
		}
		if _beg = _ecf.WriteBit(_bgc); _beg != nil {
			return _fggc, _c.Wrapf(_beg, _ecgd, "\u0062i\u0074\u003a\u0020\u0025\u0064", _cb)
		}
	}
	return _edd, nil
}

var _ _f.ByteWriter = &BufferedWriter{}
var _ _f.Writer = &BufferedWriter{}

func (_cf *BufferedWriter) WriteBits(bits uint64, number int) (_dgb int, _bg error) {
	const _bba = "\u0042u\u0066\u0066\u0065\u0072e\u0064\u0057\u0072\u0069\u0074e\u0072.\u0057r\u0069\u0074\u0065\u0072\u0042\u0069\u0074s"
	if number < 0 || number > 64 {
		return 0, _c.Errorf(_bba, "\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", number)
	}
	_ga := number / 8
	if _ga > 0 {
		_fbg := number - _ga*8
		for _cc := _ga - 1; _cc >= 0; _cc-- {
			_ab := byte((bits >> uint(_cc*8+_fbg)) & 0xff)
			if _bg = _cf.WriteByte(_ab); _bg != nil {
				return _dgb, _c.Wrapf(_bg, _bba, "\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _ga-_cc+1)
			}
		}
		number -= _ga * 8
		if number == 0 {
			return _ga, nil
		}
	}
	var _ad int
	for _ge := 0; _ge < number; _ge++ {
		if _cf._fb {
			_ad = int((bits >> uint(number-1-_ge)) & 0x1)
		} else {
			_ad = int(bits & 0x1)
			bits >>= 1
		}
		if _bg = _cf.WriteBit(_ad); _bg != nil {
			return _dgb, _c.Wrapf(_bg, _bba, "\u0062i\u0074\u003a\u0020\u0025\u0064", _ge)
		}
	}
	return _ga, nil
}
func (_cegc *Writer) writeBit(_gac uint8) error {
	if len(_cegc._bda)-1 < _cegc._aba {
		return _f.EOF
	}
	_gde := _cegc._gefd
	if _cegc._dda {
		_gde = 7 - _cegc._gefd
	}
	_cegc._bda[_cegc._aba] |= byte(uint16(_gac<<_gde) & 0xff)
	_cegc._gefd++
	if _cegc._gefd == 8 {
		_cegc._aba++
		_cegc._gefd = 0
	}
	return nil
}
func (_gd *Reader) Mark() {
	_gd._aa = _gd._add
	_gd._abf = _gd._ead
	_gd._ddd = _gd._ffd
	_gd._ef = _gd._ecg
}
func (_bbe *Reader) Reset() {
	_bbe._add = _bbe._aa
	_bbe._ead = _bbe._abf
	_bbe._ffd = _bbe._ddd
	_bbe._ecg = _bbe._ef
}
func (_fgc *Reader) ReadBool() (bool, error)  { return _fgc.readBool() }
func (_bafa *Reader) AbsolutePosition() int64 { return _bafa._add + int64(_bafa._cef._af) }
func (_feg *Writer) FinishByte() {
	if _feg._gefd == 0 {
		return
	}
	_feg._gefd = 0
	_feg._aba++
}
func NewWriter(data []byte) *Writer          { return &Writer{_bda: data} }
func (_efd *Reader) RelativePosition() int64 { return _efd._add }
func (_ag *BufferedWriter) FinishByte() {
	if _ag._fc == 0 {
		return
	}
	_ag._fc = 0
	_ag._a++
}

const (
	_gc = 64
	_dg = int(^uint(0) >> 1)
)

func (_fba *Writer) Data() []byte      { return _fba._bda }
func NewWriterMSB(data []byte) *Writer { return &Writer{_bda: data, _dda: true} }
func (_afg *Reader) ReadBits(n byte) (_edgb uint64, _edc error) {
	if n < _afg._ead {
		_cdb := _afg._ead - n
		_edgb = uint64(_afg._ffd >> _cdb)
		_afg._ffd &= 1<<_cdb - 1
		_afg._ead = _cdb
		return _edgb, nil
	}
	if n > _afg._ead {
		if _afg._ead > 0 {
			_edgb = uint64(_afg._ffd)
			n -= _afg._ead
		}
		for n >= 8 {
			_baf, _ccd := _afg.readBufferByte()
			if _ccd != nil {
				return 0, _ccd
			}
			_edgb = _edgb<<8 + uint64(_baf)
			n -= 8
		}
		if n > 0 {
			if _afg._ffd, _edc = _afg.readBufferByte(); _edc != nil {
				return 0, _edc
			}
			_gec := 8 - n
			_edgb = _edgb<<n + uint64(_afg._ffd>>_gec)
			_afg._ffd &= 1<<_gec - 1
			_afg._ead = _gec
		} else {
			_afg._ead = 0
		}
		return _edgb, nil
	}
	_afg._ead = 0
	return uint64(_afg._ffd), nil
}

type BitWriter interface {
	WriteBit(_de int) error
	WriteBits(_fac uint64, _ggb int) (_bgb int, _bfd error)
	FinishByte()
	SkipBits(_fgbcc int) error
}

func (_gef *Reader) AbsoluteLength() uint64 { return uint64(len(_gef._cef._ee)) }
func (_feb *BufferedWriter) grow(_acg int) {
	if _feb._fg == nil && _acg < _gc {
		_feb._fg = make([]byte, _acg, _gc)
		return
	}
	_bea := len(_feb._fg)
	if _feb._fc != 0 {
		_bea++
	}
	_fgbc := cap(_feb._fg)
	switch {
	case _acg <= _fgbc/2-_bea:
		_b.Log.Trace("\u005b\u0042\u0075\u0066\u0066\u0065r\u0065\u0064\u0057\u0072\u0069t\u0065\u0072\u005d\u0020\u0067\u0072o\u0077\u0020\u002d\u0020\u0072e\u0073\u006c\u0069\u0063\u0065\u0020\u006f\u006e\u006c\u0079\u002e\u0020L\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0043\u0061\u0070\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u006e\u003a\u0020'\u0025\u0064\u0027", len(_feb._fg), cap(_feb._fg), _acg)
		_b.Log.Trace("\u0020\u006e\u0020\u003c\u003d\u0020\u0063\u0020\u002f\u0020\u0032\u0020\u002d\u006d\u002e \u0043:\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u006d\u003a\u0020\u0027\u0025\u0064\u0027", _fgbc, _bea)
		copy(_feb._fg, _feb._fg[_feb.fullOffset():])
	case _fgbc > _dg-_fgbc-_acg:
		_b.Log.Error("\u0042\u0055F\u0046\u0045\u0052 \u0074\u006f\u006f\u0020\u006c\u0061\u0072\u0067\u0065")
		return
	default:
		_db := make([]byte, 2*_fgbc+_acg)
		copy(_db, _feb._fg)
		_feb._fg = _db
	}
	_feb._fg = _feb._fg[:_bea+_acg]
}
func BufferedMSB() *BufferedWriter { return &BufferedWriter{_fb: true} }
func (_bgd *Reader) ReadByte() (byte, error) {
	if _bgd._ead == 0 {
		return _bgd.readBufferByte()
	}
	return _bgd.readUnalignedByte()
}
func (_fec *BufferedWriter) expandIfNeeded(_cg int) {
	if !_fec.tryGrowByReslice(_cg) {
		_fec.grow(_cg)
	}
}
func (_cge *Reader) Length() uint64 { return uint64(_cge._cef._bef) }
func (_bbg *BufferedWriter) WriteByte(bt byte) error {
	if _bbg._a > len(_bbg._fg)-1 || (_bbg._a == len(_bbg._fg)-1 && _bbg._fc != 0) {
		_bbg.expandIfNeeded(1)
	}
	_bbg.writeByte(bt)
	return nil
}

var _ BinaryWriter = &Writer{}

func (_cfb *Reader) ReadBit() (_bbb int, _bfc error) {
	_aag, _bfc := _cfb.readBool()
	if _bfc != nil {
		return 0, _bfc
	}
	if _aag {
		_bbb = 1
	}
	return _bbb, nil
}
func (_aaa *Writer) writeByte(_fcbb byte) error {
	if _aaa._aba > len(_aaa._bda)-1 {
		return _f.EOF
	}
	if _aaa._aba == len(_aaa._bda)-1 && _aaa._gefd != 0 {
		return _f.EOF
	}
	if _aaa._gefd == 0 {
		_aaa._bda[_aaa._aba] = _fcbb
		_aaa._aba++
		return nil
	}
	if _aaa._dda {
		_aaa._bda[_aaa._aba] |= _fcbb >> _aaa._gefd
		_aaa._aba++
		_aaa._bda[_aaa._aba] = byte(uint16(_fcbb) << (8 - _aaa._gefd) & 0xff)
	} else {
		_aaa._bda[_aaa._aba] |= byte(uint16(_fcbb) << _aaa._gefd & 0xff)
		_aaa._aba++
		_aaa._bda[_aaa._aba] = _fcbb >> (8 - _aaa._gefd)
	}
	return nil
}
