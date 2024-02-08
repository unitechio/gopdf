package sampling

import (
	_g "io"

	_f "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_a "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
)

func (_ga *Reader) ReadSamples(samples []uint32) (_d error) {
	for _ac := 0; _ac < len(samples); _ac++ {
		samples[_ac], _d = _ga.ReadSample()
		if _d != nil {
			return _d
		}
	}
	return nil
}
func NewReader(img _a.ImageBase) *Reader {
	return &Reader{_c: _f.NewReader(img.Data), _b: img, _eeb: img.ColorComponents, _eb: img.BytesPerLine*8 != img.ColorComponents*img.BitsPerComponent*img.Width}
}
func (_gb *Reader) ReadSample() (uint32, error) {
	if _gb._ae == _gb._b.Height {
		return 0, _g.EOF
	}
	_ed, _be := _gb._c.ReadBits(byte(_gb._b.BitsPerComponent))
	if _be != nil {
		return 0, _be
	}
	_gb._eeb--
	if _gb._eeb == 0 {
		_gb._eeb = _gb._b.ColorComponents
		_gb._fc++
	}
	if _gb._fc == _gb._b.Width {
		if _gb._eb {
			_gb._c.ConsumeRemainingBits()
		}
		_gb._fc = 0
		_gb._ae++
	}
	return uint32(_ed), nil
}

type Writer struct {
	_def      _a.ImageBase
	_ge       *_f.Writer
	_ca, _fbc int
	_fa       bool
}

func (_aae *Writer) WriteSample(sample uint32) error {
	if _, _deb := _aae._ge.WriteBits(uint64(sample), _aae._def.BitsPerComponent); _deb != nil {
		return _deb
	}
	_aae._fbc--
	if _aae._fbc == 0 {
		_aae._fbc = _aae._def.ColorComponents
		_aae._ca++
	}
	if _aae._ca == _aae._def.Width {
		if _aae._fa {
			_aae._ge.FinishByte()
		}
		_aae._ca = 0
	}
	return nil
}

type Reader struct {
	_b             _a.ImageBase
	_c             *_f.Reader
	_fc, _ae, _eeb int
	_eb            bool
}
type SampleWriter interface {
	WriteSample(_fe uint32) error
	WriteSamples(_bg []uint32) error
}

func ResampleUint32(data []uint32, bitsPerInputSample int, bitsPerOutputSample int) []uint32 {
	var _eee []uint32
	_af := bitsPerOutputSample
	var _bed uint32
	var _gf uint32
	_ag := 0
	_dae := 0
	_ec := 0
	for _ec < len(data) {
		if _ag > 0 {
			_cgd := _ag
			if _af < _cgd {
				_cgd = _af
			}
			_bed = (_bed << uint(_cgd)) | (_gf >> uint(bitsPerInputSample-_cgd))
			_ag -= _cgd
			if _ag > 0 {
				_gf = _gf << uint(_cgd)
			} else {
				_gf = 0
			}
			_af -= _cgd
			if _af == 0 {
				_eee = append(_eee, _bed)
				_af = bitsPerOutputSample
				_bed = 0
				_dae++
			}
		} else {
			_gd := data[_ec]
			_ec++
			_gbe := bitsPerInputSample
			if _af < _gbe {
				_gbe = _af
			}
			_ag = bitsPerInputSample - _gbe
			_bed = (_bed << uint(_gbe)) | (_gd >> uint(_ag))
			if _gbe < bitsPerInputSample {
				_gf = _gd << uint(_gbe)
			}
			_af -= _gbe
			if _af == 0 {
				_eee = append(_eee, _bed)
				_af = bitsPerOutputSample
				_bed = 0
				_dae++
			}
		}
	}
	for _ag >= bitsPerOutputSample {
		_ef := _ag
		if _af < _ef {
			_ef = _af
		}
		_bed = (_bed << uint(_ef)) | (_gf >> uint(bitsPerInputSample-_ef))
		_ag -= _ef
		if _ag > 0 {
			_gf = _gf << uint(_ef)
		} else {
			_gf = 0
		}
		_af -= _ef
		if _af == 0 {
			_eee = append(_eee, _bed)
			_af = bitsPerOutputSample
			_bed = 0
			_dae++
		}
	}
	if _af > 0 && _af < bitsPerOutputSample {
		_bed <<= uint(_af)
		_eee = append(_eee, _bed)
	}
	return _eee
}
func (_debb *Writer) WriteSamples(samples []uint32) error {
	for _dab := 0; _dab < len(samples); _dab++ {
		if _gea := _debb.WriteSample(samples[_dab]); _gea != nil {
			return _gea
		}
	}
	return nil
}
func NewWriter(img _a.ImageBase) *Writer {
	return &Writer{_ge: _f.NewWriterMSB(img.Data), _def: img, _fbc: img.ColorComponents, _fa: img.BytesPerLine*8 != img.ColorComponents*img.BitsPerComponent*img.Width}
}

type SampleReader interface {
	ReadSample() (uint32, error)
	ReadSamples(_ee []uint32) error
}

func ResampleBytes(data []byte, bitsPerSample int) []uint32 {
	var _da []uint32
	_dad := bitsPerSample
	var _dc uint32
	var _aa byte
	_de := 0
	_cg := 0
	_gc := 0
	for _gc < len(data) {
		if _de > 0 {
			_fg := _de
			if _dad < _fg {
				_fg = _dad
			}
			_dc = (_dc << uint(_fg)) | uint32(_aa>>uint(8-_fg))
			_de -= _fg
			if _de > 0 {
				_aa = _aa << uint(_fg)
			} else {
				_aa = 0
			}
			_dad -= _fg
			if _dad == 0 {
				_da = append(_da, _dc)
				_dad = bitsPerSample
				_dc = 0
				_cg++
			}
		} else {
			_fb := data[_gc]
			_gc++
			_eed := 8
			if _dad < _eed {
				_eed = _dad
			}
			_de = 8 - _eed
			_dc = (_dc << uint(_eed)) | uint32(_fb>>uint(_de))
			if _eed < 8 {
				_aa = _fb << uint(_eed)
			}
			_dad -= _eed
			if _dad == 0 {
				_da = append(_da, _dc)
				_dad = bitsPerSample
				_dc = 0
				_cg++
			}
		}
	}
	for _de >= bitsPerSample {
		_dg := _de
		if _dad < _dg {
			_dg = _dad
		}
		_dc = (_dc << uint(_dg)) | uint32(_aa>>uint(8-_dg))
		_de -= _dg
		if _de > 0 {
			_aa = _aa << uint(_dg)
		} else {
			_aa = 0
		}
		_dad -= _dg
		if _dad == 0 {
			_da = append(_da, _dc)
			_dad = bitsPerSample
			_dc = 0
			_cg++
		}
	}
	return _da
}
