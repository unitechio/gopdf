package sampling

import (
	_ff "io"

	_b "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_d "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
)

type Writer struct {
	_af       _d.ImageBase
	_dgca     *_b.Writer
	_gff, _aa int
	_cad      bool
}

func NewReader(img _d.ImageBase) *Reader {
	return &Reader{_g: _b.NewReader(img.Data), _fd: img, _be: img.ColorComponents, _dg: img.BytesPerLine*8 != img.ColorComponents*img.BitsPerComponent*img.Width}
}

type Reader struct {
	_fd          _d.ImageBase
	_g           *_b.Reader
	_gb, _a, _be int
	_dg          bool
}

func (_ef *Writer) WriteSample(sample uint32) error {
	if _, _cadf := _ef._dgca.WriteBits(uint64(sample), _ef._af.BitsPerComponent); _cadf != nil {
		return _cadf
	}
	_ef._aa--
	if _ef._aa == 0 {
		_ef._aa = _ef._af.ColorComponents
		_ef._gff++
	}
	if _ef._gff == _ef._af.Width {
		if _ef._cad {
			_ef._dgca.FinishByte()
		}
		_ef._gff = 0
	}
	return nil
}
func ResampleBytes(data []byte, bitsPerSample int) []uint32 {
	var _bc []uint32
	_dfe := bitsPerSample
	var _ee uint32
	var _fb byte
	_gd := 0
	_dgb := 0
	_dgg := 0
	for _dgg < len(data) {
		if _gd > 0 {
			_ge := _gd
			if _dfe < _ge {
				_ge = _dfe
			}
			_ee = (_ee << uint(_ge)) | uint32(_fb>>uint(8-_ge))
			_gd -= _ge
			if _gd > 0 {
				_fb = _fb << uint(_ge)
			} else {
				_fb = 0
			}
			_dfe -= _ge
			if _dfe == 0 {
				_bc = append(_bc, _ee)
				_dfe = bitsPerSample
				_ee = 0
				_dgb++
			}
		} else {
			_ca := data[_dgg]
			_dgg++
			_ag := 8
			if _dfe < _ag {
				_ag = _dfe
			}
			_gd = 8 - _ag
			_ee = (_ee << uint(_ag)) | uint32(_ca>>uint(_gd))
			if _ag < 8 {
				_fb = _ca << uint(_ag)
			}
			_dfe -= _ag
			if _dfe == 0 {
				_bc = append(_bc, _ee)
				_dfe = bitsPerSample
				_ee = 0
				_dgb++
			}
		}
	}
	for _gd >= bitsPerSample {
		_bd := _gd
		if _dfe < _bd {
			_bd = _dfe
		}
		_ee = (_ee << uint(_bd)) | uint32(_fb>>uint(8-_bd))
		_gd -= _bd
		if _gd > 0 {
			_fb = _fb << uint(_bd)
		} else {
			_fb = 0
		}
		_dfe -= _bd
		if _dfe == 0 {
			_bc = append(_bc, _ee)
			_dfe = bitsPerSample
			_ee = 0
			_dgb++
		}
	}
	return _bc
}

type SampleReader interface {
	ReadSample() (uint32, error)
	ReadSamples(_c []uint32) error
}

func ResampleUint32(data []uint32, bitsPerInputSample int, bitsPerOutputSample int) []uint32 {
	var _bg []uint32
	_gg := bitsPerOutputSample
	var _bf uint32
	var _cd uint32
	_bgg := 0
	_gcb := 0
	_cf := 0
	for _cf < len(data) {
		if _bgg > 0 {
			_cfg := _bgg
			if _gg < _cfg {
				_cfg = _gg
			}
			_bf = (_bf << uint(_cfg)) | (_cd >> uint(bitsPerInputSample-_cfg))
			_bgg -= _cfg
			if _bgg > 0 {
				_cd = _cd << uint(_cfg)
			} else {
				_cd = 0
			}
			_gg -= _cfg
			if _gg == 0 {
				_bg = append(_bg, _bf)
				_gg = bitsPerOutputSample
				_bf = 0
				_gcb++
			}
		} else {
			_dgc := data[_cf]
			_cf++
			_ce := bitsPerInputSample
			if _gg < _ce {
				_ce = _gg
			}
			_bgg = bitsPerInputSample - _ce
			_bf = (_bf << uint(_ce)) | (_dgc >> uint(_bgg))
			if _ce < bitsPerInputSample {
				_cd = _dgc << uint(_ce)
			}
			_gg -= _ce
			if _gg == 0 {
				_bg = append(_bg, _bf)
				_gg = bitsPerOutputSample
				_bf = 0
				_gcb++
			}
		}
	}
	for _bgg >= bitsPerOutputSample {
		_fc := _bgg
		if _gg < _fc {
			_fc = _gg
		}
		_bf = (_bf << uint(_fc)) | (_cd >> uint(bitsPerInputSample-_fc))
		_bgg -= _fc
		if _bgg > 0 {
			_cd = _cd << uint(_fc)
		} else {
			_cd = 0
		}
		_gg -= _fc
		if _gg == 0 {
			_bg = append(_bg, _bf)
			_gg = bitsPerOutputSample
			_bf = 0
			_gcb++
		}
	}
	if _gg > 0 && _gg < bitsPerOutputSample {
		_bf <<= uint(_gg)
		_bg = append(_bg, _bf)
	}
	return _bg
}
func (_gc *Reader) ReadSamples(samples []uint32) (_df error) {
	for _cc := 0; _cc < len(samples); _cc++ {
		samples[_cc], _df = _gc.ReadSample()
		if _df != nil {
			return _df
		}
	}
	return nil
}
func (_e *Reader) ReadSample() (uint32, error) {
	if _e._a == _e._fd.Height {
		return 0, _ff.EOF
	}
	_gf, _de := _e._g.ReadBits(byte(_e._fd.BitsPerComponent))
	if _de != nil {
		return 0, _de
	}
	_e._be--
	if _e._be == 0 {
		_e._be = _e._fd.ColorComponents
		_e._gb++
	}
	if _e._gb == _e._fd.Width {
		if _e._dg {
			_e._g.ConsumeRemainingBits()
		}
		_e._gb = 0
		_e._a++
	}
	return uint32(_gf), nil
}
func (_fdc *Writer) WriteSamples(samples []uint32) error {
	for _afg := 0; _afg < len(samples); _afg++ {
		if _dfd := _fdc.WriteSample(samples[_afg]); _dfd != nil {
			return _dfd
		}
	}
	return nil
}

type SampleWriter interface {
	WriteSample(_cg uint32) error
	WriteSamples(_cce []uint32) error
}

func NewWriter(img _d.ImageBase) *Writer {
	return &Writer{_dgca: _b.NewWriterMSB(img.Data), _af: img, _aa: img.ColorComponents, _cad: img.BytesPerLine*8 != img.ColorComponents*img.BitsPerComponent*img.Width}
}
