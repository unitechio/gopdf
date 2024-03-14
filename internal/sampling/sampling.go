package sampling

import (
	_a "io"

	_e "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_c "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
)

func ResampleUint32(data []uint32, bitsPerInputSample int, bitsPerOutputSample int) []uint32 {
	var _ega []uint32
	_egab := bitsPerOutputSample
	var _cd uint32
	var _ba uint32
	_gb := 0
	_gdb := 0
	_abd := 0
	for _abd < len(data) {
		if _gb > 0 {
			_eec := _gb
			if _egab < _eec {
				_eec = _egab
			}
			_cd = (_cd << uint(_eec)) | (_ba >> uint(bitsPerInputSample-_eec))
			_gb -= _eec
			if _gb > 0 {
				_ba = _ba << uint(_eec)
			} else {
				_ba = 0
			}
			_egab -= _eec
			if _egab == 0 {
				_ega = append(_ega, _cd)
				_egab = bitsPerOutputSample
				_cd = 0
				_gdb++
			}
		} else {
			_fb := data[_abd]
			_abd++
			_fbb := bitsPerInputSample
			if _egab < _fbb {
				_fbb = _egab
			}
			_gb = bitsPerInputSample - _fbb
			_cd = (_cd << uint(_fbb)) | (_fb >> uint(_gb))
			if _fbb < bitsPerInputSample {
				_ba = _fb << uint(_fbb)
			}
			_egab -= _fbb
			if _egab == 0 {
				_ega = append(_ega, _cd)
				_egab = bitsPerOutputSample
				_cd = 0
				_gdb++
			}
		}
	}
	for _gb >= bitsPerOutputSample {
		_gdd := _gb
		if _egab < _gdd {
			_gdd = _egab
		}
		_cd = (_cd << uint(_gdd)) | (_ba >> uint(bitsPerInputSample-_gdd))
		_gb -= _gdd
		if _gb > 0 {
			_ba = _ba << uint(_gdd)
		} else {
			_ba = 0
		}
		_egab -= _gdd
		if _egab == 0 {
			_ega = append(_ega, _cd)
			_egab = bitsPerOutputSample
			_cd = 0
			_gdb++
		}
	}
	if _egab > 0 && _egab < bitsPerOutputSample {
		_cd <<= uint(_egab)
		_ega = append(_ega, _cd)
	}
	return _ega
}

func (_gd *Reader) ReadSamples(samples []uint32) (_ca error) {
	for _ce := 0; _ce < len(samples); _ce++ {
		samples[_ce], _ca = _gd.ReadSample()
		if _ca != nil {
			return _ca
		}
	}
	return nil
}

func ResampleBytes(data []byte, bitsPerSample int) []uint32 {
	var _cgf []uint32
	_fc := bitsPerSample
	var _dd uint32
	var _ee byte
	_da := 0
	_gf := 0
	_eeg := 0
	for _eeg < len(data) {
		if _da > 0 {
			_ede := _da
			if _fc < _ede {
				_ede = _fc
			}
			_dd = (_dd << uint(_ede)) | uint32(_ee>>uint(8-_ede))
			_da -= _ede
			if _da > 0 {
				_ee = _ee << uint(_ede)
			} else {
				_ee = 0
			}
			_fc -= _ede
			if _fc == 0 {
				_cgf = append(_cgf, _dd)
				_fc = bitsPerSample
				_dd = 0
				_gf++
			}
		} else {
			_ab := data[_eeg]
			_eeg++
			_gg := 8
			if _fc < _gg {
				_gg = _fc
			}
			_da = 8 - _gg
			_dd = (_dd << uint(_gg)) | uint32(_ab>>uint(_da))
			if _gg < 8 {
				_ee = _ab << uint(_gg)
			}
			_fc -= _gg
			if _fc == 0 {
				_cgf = append(_cgf, _dd)
				_fc = bitsPerSample
				_dd = 0
				_gf++
			}
		}
	}
	for _da >= bitsPerSample {
		_cgb := _da
		if _fc < _cgb {
			_cgb = _fc
		}
		_dd = (_dd << uint(_cgb)) | uint32(_ee>>uint(8-_cgb))
		_da -= _cgb
		if _da > 0 {
			_ee = _ee << uint(_cgb)
		} else {
			_ee = 0
		}
		_fc -= _cgb
		if _fc == 0 {
			_cgf = append(_cgf, _dd)
			_fc = bitsPerSample
			_dd = 0
			_gf++
		}
	}
	return _cgf
}

func (_ga *Writer) WriteSamples(samples []uint32) error {
	for _af := 0; _af < len(samples); _af++ {
		if _fca := _ga.WriteSample(samples[_af]); _fca != nil {
			return _fca
		}
	}
	return nil
}

func NewReader(img _c.ImageBase) *Reader {
	return &Reader{_cf: _e.NewReader(img.Data), _f: img, _d: img.ColorComponents, _de: img.BytesPerLine*8 != img.ColorComponents*img.BitsPerComponent*img.Width}
}

type Writer struct {
	_cc      _c.ImageBase
	_dbd     *_e.Writer
	_ge, _ec int
	_ecg     bool
}
type SampleWriter interface {
	WriteSample(_db uint32) error
	WriteSamples(_dbc []uint32) error
}

func (_g *Reader) ReadSample() (uint32, error) {
	if _g._ad == _g._f.Height {
		return 0, _a.EOF
	}
	_ed, _eb := _g._cf.ReadBits(byte(_g._f.BitsPerComponent))
	if _eb != nil {
		return 0, _eb
	}
	_g._d--
	if _g._d == 0 {
		_g._d = _g._f.ColorComponents
		_g._cg++
	}
	if _g._cg == _g._f.Width {
		if _g._de {
			_g._cf.ConsumeRemainingBits()
		}
		_g._cg = 0
		_g._ad++
	}
	return uint32(_ed), nil
}

type SampleReader interface {
	ReadSample() (uint32, error)
	ReadSamples(_eg []uint32) error
}

func NewWriter(img _c.ImageBase) *Writer {
	return &Writer{_dbd: _e.NewWriterMSB(img.Data), _cc: img, _ec: img.ColorComponents, _ecg: img.BytesPerLine*8 != img.ColorComponents*img.BitsPerComponent*img.Width}
}

type Reader struct {
	_f           _c.ImageBase
	_cf          *_e.Reader
	_cg, _ad, _d int
	_de          bool
}

func (_eeb *Writer) WriteSample(sample uint32) error {
	if _, _dg := _eeb._dbd.WriteBits(uint64(sample), _eeb._cc.BitsPerComponent); _dg != nil {
		return _dg
	}
	_eeb._ec--
	if _eeb._ec == 0 {
		_eeb._ec = _eeb._cc.ColorComponents
		_eeb._ge++
	}
	if _eeb._ge == _eeb._cc.Width {
		if _eeb._ecg {
			_eeb._dbd.FinishByte()
		}
		_eeb._ge = 0
	}
	return nil
}
