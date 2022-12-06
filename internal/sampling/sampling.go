package sampling

import (
	_g "io"

	_e "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_cb "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
)

func NewReader(img _cb.ImageBase) *Reader {
	return &Reader{_cg: _e.NewReader(img.Data), _d: img, _da: img.ColorComponents, _a: img.BytesPerLine*8 != img.ColorComponents*img.BitsPerComponent*img.Width}
}
func (_cge *Writer) WriteSamples(samples []uint32) error {
	for _ae := 0; _ae < len(samples); _ae++ {
		if _gf := _cge.WriteSample(samples[_ae]); _gf != nil {
			return _gf
		}
	}
	return nil
}

type SampleReader interface {
	ReadSample() (uint32, error)
	ReadSamples(_cd []uint32) error
}
type Reader struct {
	_d             _cb.ImageBase
	_cg            *_e.Reader
	_cdg, _eg, _da int
	_a             bool
}

func (_ee *Writer) WriteSample(sample uint32) error {
	if _, _efa := _ee._ca.WriteBits(uint64(sample), _ee._dfa.BitsPerComponent); _efa != nil {
		return _efa
	}
	_ee._eda--
	if _ee._eda == 0 {
		_ee._eda = _ee._dfa.ColorComponents
		_ee._cac++
	}
	if _ee._cac == _ee._dfa.Width {
		if _ee._bde {
			_ee._ca.FinishByte()
		}
		_ee._cac = 0
	}
	return nil
}
func ResampleUint32(data []uint32, bitsPerInputSample int, bitsPerOutputSample int) []uint32 {
	var _efg []uint32
	_cdf := bitsPerOutputSample
	var _df uint32
	var _cbd uint32
	_afg := 0
	_f := 0
	_dc := 0
	for _dc < len(data) {
		if _afg > 0 {
			_fg := _afg
			if _cdf < _fg {
				_fg = _cdf
			}
			_df = (_df << uint(_fg)) | (_cbd >> uint(bitsPerInputSample-_fg))
			_afg -= _fg
			if _afg > 0 {
				_cbd = _cbd << uint(_fg)
			} else {
				_cbd = 0
			}
			_cdf -= _fg
			if _cdf == 0 {
				_efg = append(_efg, _df)
				_cdf = bitsPerOutputSample
				_df = 0
				_f++
			}
		} else {
			_bb := data[_dc]
			_dc++
			_ede := bitsPerInputSample
			if _cdf < _ede {
				_ede = _cdf
			}
			_afg = bitsPerInputSample - _ede
			_df = (_df << uint(_ede)) | (_bb >> uint(_afg))
			if _ede < bitsPerInputSample {
				_cbd = _bb << uint(_ede)
			}
			_cdf -= _ede
			if _cdf == 0 {
				_efg = append(_efg, _df)
				_cdf = bitsPerOutputSample
				_df = 0
				_f++
			}
		}
	}
	for _afg >= bitsPerOutputSample {
		_ege := _afg
		if _cdf < _ege {
			_ege = _cdf
		}
		_df = (_df << uint(_ege)) | (_cbd >> uint(bitsPerInputSample-_ege))
		_afg -= _ege
		if _afg > 0 {
			_cbd = _cbd << uint(_ege)
		} else {
			_cbd = 0
		}
		_cdf -= _ege
		if _cdf == 0 {
			_efg = append(_efg, _df)
			_cdf = bitsPerOutputSample
			_df = 0
			_f++
		}
	}
	if _cdf > 0 && _cdf < bitsPerOutputSample {
		_df <<= uint(_cdf)
		_efg = append(_efg, _df)
	}
	return _efg
}

type Writer struct {
	_dfa       _cb.ImageBase
	_ca        *_e.Writer
	_cac, _eda int
	_bde       bool
}

func (_gb *Reader) ReadSamples(samples []uint32) (_b error) {
	for _ce := 0; _ce < len(samples); _ce++ {
		samples[_ce], _b = _gb.ReadSample()
		if _b != nil {
			return _b
		}
	}
	return nil
}
func (_ef *Reader) ReadSample() (uint32, error) {
	if _ef._eg == _ef._d.Height {
		return 0, _g.EOF
	}
	_ed, _cbe := _ef._cg.ReadBits(byte(_ef._d.BitsPerComponent))
	if _cbe != nil {
		return 0, _cbe
	}
	_ef._da--
	if _ef._da == 0 {
		_ef._da = _ef._d.ColorComponents
		_ef._cdg++
	}
	if _ef._cdg == _ef._d.Width {
		if _ef._a {
			_ef._cg.ConsumeRemainingBits()
		}
		_ef._cdg = 0
		_ef._eg++
	}
	return uint32(_ed), nil
}

type SampleWriter interface {
	WriteSample(_be uint32) error
	WriteSamples(_fc []uint32) error
}

func NewWriter(img _cb.ImageBase) *Writer {
	return &Writer{_ca: _e.NewWriterMSB(img.Data), _dfa: img, _eda: img.ColorComponents, _bde: img.BytesPerLine*8 != img.ColorComponents*img.BitsPerComponent*img.Width}
}
func ResampleBytes(data []byte, bitsPerSample int) []uint32 {
	var _efb []uint32
	_bg := bitsPerSample
	var _bc uint32
	var _bgb byte
	_de := 0
	_bd := 0
	_bcf := 0
	for _bcf < len(data) {
		if _de > 0 {
			_af := _de
			if _bg < _af {
				_af = _bg
			}
			_bc = (_bc << uint(_af)) | uint32(_bgb>>uint(8-_af))
			_de -= _af
			if _de > 0 {
				_bgb = _bgb << uint(_af)
			} else {
				_bgb = 0
			}
			_bg -= _af
			if _bg == 0 {
				_efb = append(_efb, _bc)
				_bg = bitsPerSample
				_bc = 0
				_bd++
			}
		} else {
			_gd := data[_bcf]
			_bcf++
			_gc := 8
			if _bg < _gc {
				_gc = _bg
			}
			_de = 8 - _gc
			_bc = (_bc << uint(_gc)) | uint32(_gd>>uint(_de))
			if _gc < 8 {
				_bgb = _gd << uint(_gc)
			}
			_bg -= _gc
			if _bg == 0 {
				_efb = append(_efb, _bc)
				_bg = bitsPerSample
				_bc = 0
				_bd++
			}
		}
	}
	for _de >= bitsPerSample {
		_ceg := _de
		if _bg < _ceg {
			_ceg = _bg
		}
		_bc = (_bc << uint(_ceg)) | uint32(_bgb>>uint(8-_ceg))
		_de -= _ceg
		if _de > 0 {
			_bgb = _bgb << uint(_ceg)
		} else {
			_bgb = 0
		}
		_bg -= _ceg
		if _bg == 0 {
			_efb = append(_efb, _bc)
			_bg = bitsPerSample
			_bc = 0
			_bd++
		}
	}
	return _efb
}
