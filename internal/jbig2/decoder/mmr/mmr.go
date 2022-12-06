package mmr

import (
	_d "errors"
	_g "fmt"
	_c "io"

	_af "bitbucket.org/shenghui0779/gopdf/common"
	_f "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_da "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
)

func (_ba *Decoder) createLittleEndianTable(_fde [][3]int) ([]*code, error) {
	_bad := make([]*code, _ga+1)
	for _gff := 0; _gff < len(_fde); _gff++ {
		_gbc := _gf(_fde[_gff])
		if _gbc._b <= _ed {
			_ece := _ed - _gbc._b
			_bg := _gbc._fg << uint(_ece)
			for _dc := (1 << uint(_ece)) - 1; _dc >= 0; _dc-- {
				_gba := _bg | _dc
				_bad[_gba] = _gbc
			}
		} else {
			_cdg := _gbc._fg >> uint(_gbc._b-_ed)
			if _bad[_cdg] == nil {
				var _gbb = _gf([3]int{})
				_gbb._e = make([]*code, _fbb+1)
				_bad[_cdg] = _gbb
			}
			if _gbc._b <= _ed+_dfg {
				_cbg := _ed + _dfg - _gbc._b
				_ggd := (_gbc._fg << uint(_cbg)) & _fbb
				_bad[_cdg]._fb = true
				for _gdb := (1 << uint(_cbg)) - 1; _gdb >= 0; _gdb-- {
					_bad[_cdg]._e[_ggd|_gdb] = _gbc
				}
			} else {
				return nil, _d.New("\u0043\u006f\u0064\u0065\u0020\u0074a\u0062\u006c\u0065\u0020\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077\u0020i\u006e\u0020\u004d\u004d\u0052\u0044\u0065c\u006f\u0064\u0065\u0072")
			}
		}
	}
	return _bad, nil
}
func (_bgg *runData) uncompressGetNextCodeLittleEndian() (int, error) {
	_fae := _bgg._aga - _bgg._ce
	if _fae < 0 || _fae > 24 {
		_gcf := (_bgg._aga >> 3) - _bgg._feb
		if _gcf >= _bgg._bd {
			_gcf += _bgg._feb
			if _be := _bgg.fillBuffer(_gcf); _be != nil {
				return 0, _be
			}
			_gcf -= _bgg._feb
		}
		_aegb := (uint32(_bgg._fag[_gcf]&0xFF) << 16) | (uint32(_bgg._fag[_gcf+1]&0xFF) << 8) | (uint32(_bgg._fag[_gcf+2] & 0xFF))
		_cgd := uint32(_bgg._aga & 7)
		_aegb <<= _cgd
		_bgg._ecd = int(_aegb)
	} else {
		_bga := _bgg._ce & 7
		_gdd := 7 - _bga
		if _fae <= _gdd {
			_bgg._ecd <<= uint(_fae)
		} else {
			_geae := (_bgg._ce >> 3) + 3 - _bgg._feb
			if _geae >= _bgg._bd {
				_geae += _bgg._feb
				if _eed := _bgg.fillBuffer(_geae); _eed != nil {
					return 0, _eed
				}
				_geae -= _bgg._feb
			}
			_bga = 8 - _bga
			for {
				_bgg._ecd <<= uint(_bga)
				_bgg._ecd |= int(uint(_bgg._fag[_geae]) & 0xFF)
				_fae -= _bga
				_geae++
				_bga = 8
				if !(_fae >= 8) {
					break
				}
			}
			_bgg._ecd <<= uint(_fae)
		}
	}
	_bgg._ce = _bgg._aga
	return _bgg._ecd, nil
}
func (_afc *Decoder) UncompressMMR() (_abc *_da.Bitmap, _aba error) {
	_abc = _da.New(_afc._fbbb, _afc._gb)
	_abd := make([]int, _abc.Width+5)
	_ebe := make([]int, _abc.Width+5)
	_ebe[0] = _abc.Width
	_gdg := 1
	var _ge int
	for _cg := 0; _cg < _abc.Height; _cg++ {
		_ge, _aba = _afc.uncompress2d(_afc._gd, _ebe, _gdg, _abd, _abc.Width)
		if _aba != nil {
			return nil, _aba
		}
		if _ge == EOF {
			break
		}
		if _ge > 0 {
			_aba = _afc.fillBitmap(_abc, _cg, _abd, _ge)
			if _aba != nil {
				return nil, _aba
			}
		}
		_ebe, _abd = _abd, _ebe
		_gdg = _ge
	}
	if _aba = _afc.detectAndSkipEOL(); _aba != nil {
		return nil, _aba
	}
	_afc._gd.align()
	return _abc, nil
}
func (_ee *runData) uncompressGetCode(_dge []*code) (*code, error) {
	return _ee.uncompressGetCodeLittleEndian(_dge)
}
func _cb(_ea, _ccb int) int {
	if _ea > _ccb {
		return _ccb
	}
	return _ea
}
func (_ad *code) String() string {
	return _g.Sprintf("\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064", _ad._b, _ad._fg, _ad._ca)
}
func New(r _f.StreamReader, width, height int, dataOffset, dataLength int64) (*Decoder, error) {
	_fd := &Decoder{_fbbb: width, _gb: height}
	_cbf, _ab := _f.NewSubstreamReader(r, uint64(dataOffset), uint64(dataLength))
	if _ab != nil {
		return nil, _ab
	}
	_eaf, _ab := _ebea(_cbf)
	if _ab != nil {
		return nil, _ab
	}
	_fd._gd = _eaf
	if _fdc := _fd.initTables(); _fdc != nil {
		return nil, _fdc
	}
	return _fd, nil
}

type mmrCode int

const (
	EOF  = -3
	_aed = -2
	EOL  = -1
	_ed  = 8
	_ga  = (1 << _ed) - 1
	_dfg = 5
	_fbb = (1 << _dfg) - 1
)
const (
	_aaf  int  = 1024 << 7
	_dff  int  = 3
	_bbgg uint = 24
)

var (
	_bb   = [][3]int{{4, 0x1, int(_ag)}, {3, 0x1, int(_cce)}, {1, 0x1, int(_aa)}, {3, 0x3, int(_gg)}, {6, 0x3, int(_ec)}, {7, 0x3, int(_afe)}, {3, 0x2, int(_ccc)}, {6, 0x2, int(_cag)}, {7, 0x2, int(_eb)}, {10, 0xf, int(_gc)}, {12, 0xf, int(_df)}, {12, 0x1, int(EOL)}}
	_bbg  = [][3]int{{4, 0x07, 2}, {4, 0x08, 3}, {4, 0x0B, 4}, {4, 0x0C, 5}, {4, 0x0E, 6}, {4, 0x0F, 7}, {5, 0x12, 128}, {5, 0x13, 8}, {5, 0x14, 9}, {5, 0x1B, 64}, {5, 0x07, 10}, {5, 0x08, 11}, {6, 0x17, 192}, {6, 0x18, 1664}, {6, 0x2A, 16}, {6, 0x2B, 17}, {6, 0x03, 13}, {6, 0x34, 14}, {6, 0x35, 15}, {6, 0x07, 1}, {6, 0x08, 12}, {7, 0x13, 26}, {7, 0x17, 21}, {7, 0x18, 28}, {7, 0x24, 27}, {7, 0x27, 18}, {7, 0x28, 24}, {7, 0x2B, 25}, {7, 0x03, 22}, {7, 0x37, 256}, {7, 0x04, 23}, {7, 0x08, 20}, {7, 0xC, 19}, {8, 0x12, 33}, {8, 0x13, 34}, {8, 0x14, 35}, {8, 0x15, 36}, {8, 0x16, 37}, {8, 0x17, 38}, {8, 0x1A, 31}, {8, 0x1B, 32}, {8, 0x02, 29}, {8, 0x24, 53}, {8, 0x25, 54}, {8, 0x28, 39}, {8, 0x29, 40}, {8, 0x2A, 41}, {8, 0x2B, 42}, {8, 0x2C, 43}, {8, 0x2D, 44}, {8, 0x03, 30}, {8, 0x32, 61}, {8, 0x33, 62}, {8, 0x34, 63}, {8, 0x35, 0}, {8, 0x36, 320}, {8, 0x37, 384}, {8, 0x04, 45}, {8, 0x4A, 59}, {8, 0x4B, 60}, {8, 0x5, 46}, {8, 0x52, 49}, {8, 0x53, 50}, {8, 0x54, 51}, {8, 0x55, 52}, {8, 0x58, 55}, {8, 0x59, 56}, {8, 0x5A, 57}, {8, 0x5B, 58}, {8, 0x64, 448}, {8, 0x65, 512}, {8, 0x67, 640}, {8, 0x68, 576}, {8, 0x0A, 47}, {8, 0x0B, 48}, {9, 0x01, _aed}, {9, 0x98, 1472}, {9, 0x99, 1536}, {9, 0x9A, 1600}, {9, 0x9B, 1728}, {9, 0xCC, 704}, {9, 0xCD, 768}, {9, 0xD2, 832}, {9, 0xD3, 896}, {9, 0xD4, 960}, {9, 0xD5, 1024}, {9, 0xD6, 1088}, {9, 0xD7, 1152}, {9, 0xD8, 1216}, {9, 0xD9, 1280}, {9, 0xDA, 1344}, {9, 0xDB, 1408}, {10, 0x01, _aed}, {11, 0x01, _aed}, {11, 0x08, 1792}, {11, 0x0C, 1856}, {11, 0x0D, 1920}, {12, 0x00, EOF}, {12, 0x01, EOL}, {12, 0x12, 1984}, {12, 0x13, 2048}, {12, 0x14, 2112}, {12, 0x15, 2176}, {12, 0x16, 2240}, {12, 0x17, 2304}, {12, 0x1C, 2368}, {12, 0x1D, 2432}, {12, 0x1E, 2496}, {12, 0x1F, 2560}}
	_aedg = [][3]int{{2, 0x02, 3}, {2, 0x03, 2}, {3, 0x02, 1}, {3, 0x03, 4}, {4, 0x02, 6}, {4, 0x03, 5}, {5, 0x03, 7}, {6, 0x04, 9}, {6, 0x05, 8}, {7, 0x04, 10}, {7, 0x05, 11}, {7, 0x07, 12}, {8, 0x04, 13}, {8, 0x07, 14}, {9, 0x01, _aed}, {9, 0x18, 15}, {10, 0x01, _aed}, {10, 0x17, 16}, {10, 0x18, 17}, {10, 0x37, 0}, {10, 0x08, 18}, {10, 0x0F, 64}, {11, 0x01, _aed}, {11, 0x17, 24}, {11, 0x18, 25}, {11, 0x28, 23}, {11, 0x37, 22}, {11, 0x67, 19}, {11, 0x68, 20}, {11, 0x6C, 21}, {11, 0x08, 1792}, {11, 0x0C, 1856}, {11, 0x0D, 1920}, {12, 0x00, EOF}, {12, 0x01, EOL}, {12, 0x12, 1984}, {12, 0x13, 2048}, {12, 0x14, 2112}, {12, 0x15, 2176}, {12, 0x16, 2240}, {12, 0x17, 2304}, {12, 0x1C, 2368}, {12, 0x1D, 2432}, {12, 0x1E, 2496}, {12, 0x1F, 2560}, {12, 0x24, 52}, {12, 0x27, 55}, {12, 0x28, 56}, {12, 0x2B, 59}, {12, 0x2C, 60}, {12, 0x33, 320}, {12, 0x34, 384}, {12, 0x35, 448}, {12, 0x37, 53}, {12, 0x38, 54}, {12, 0x52, 50}, {12, 0x53, 51}, {12, 0x54, 44}, {12, 0x55, 45}, {12, 0x56, 46}, {12, 0x57, 47}, {12, 0x58, 57}, {12, 0x59, 58}, {12, 0x5A, 61}, {12, 0x5B, 256}, {12, 0x64, 48}, {12, 0x65, 49}, {12, 0x66, 62}, {12, 0x67, 63}, {12, 0x68, 30}, {12, 0x69, 31}, {12, 0x6A, 32}, {12, 0x6B, 33}, {12, 0x6C, 40}, {12, 0x6D, 41}, {12, 0xC8, 128}, {12, 0xC9, 192}, {12, 0xCA, 26}, {12, 0xCB, 27}, {12, 0xCC, 28}, {12, 0xCD, 29}, {12, 0xD2, 34}, {12, 0xD3, 35}, {12, 0xD4, 36}, {12, 0xD5, 37}, {12, 0xD6, 38}, {12, 0xD7, 39}, {12, 0xDA, 42}, {12, 0xDB, 43}, {13, 0x4A, 640}, {13, 0x4B, 704}, {13, 0x4C, 768}, {13, 0x4D, 832}, {13, 0x52, 1280}, {13, 0x53, 1344}, {13, 0x54, 1408}, {13, 0x55, 1472}, {13, 0x5A, 1536}, {13, 0x5B, 1600}, {13, 0x64, 1664}, {13, 0x65, 1728}, {13, 0x6C, 512}, {13, 0x6D, 576}, {13, 0x72, 896}, {13, 0x73, 960}, {13, 0x74, 1024}, {13, 0x75, 1088}, {13, 0x76, 1152}, {13, 0x77, 1216}}
)

func _gf(_ae [3]int) *code { return &code{_b: _ae[0], _fg: _ae[1], _ca: _ae[2]} }

type code struct {
	_b  int
	_fg int
	_ca int
	_e  []*code
	_fb bool
}

func (_acb *Decoder) detectAndSkipEOL() error {
	for {
		_eaa, _cae := _acb._gd.uncompressGetCode(_acb._cda)
		if _cae != nil {
			return _cae
		}
		if _eaa != nil && _eaa._ca == EOL {
			_acb._gd._aga += _eaa._b
		} else {
			return nil
		}
	}
}
func (_gab *runData) align() { _gab._aga = ((_gab._aga + 7) >> 3) << 3 }
func (_bfg *Decoder) uncompress2d(_ccd *runData, _gad []int, _agg int, _cge []int, _bbfc int) (int, error) {
	var (
		_adg int
		_dbc int
		_bbd int
		_ffd = true
		_dd  error
		_edd *code
	)
	_gad[_agg] = _bbfc
	_gad[_agg+1] = _bbfc
	_gad[_agg+2] = _bbfc + 1
	_gad[_agg+3] = _bbfc + 1
_dag:
	for _bbd < _bbfc {
		_edd, _dd = _ccd.uncompressGetCode(_bfg._cda)
		if _dd != nil {
			return EOL, nil
		}
		if _edd == nil {
			_ccd._aga++
			break _dag
		}
		_ccd._aga += _edd._b
		switch mmrCode(_edd._ca) {
		case _aa:
			_bbd = _gad[_adg]
		case _gg:
			_bbd = _gad[_adg] + 1
		case _ccc:
			_bbd = _gad[_adg] - 1
		case _cce:
			for {
				var _aeg []*code
				if _ffd {
					_aeg = _bfg._gaf
				} else {
					_aeg = _bfg._ff
				}
				_edd, _dd = _ccd.uncompressGetCode(_aeg)
				if _dd != nil {
					return 0, _dd
				}
				if _edd == nil {
					break _dag
				}
				_ccd._aga += _edd._b
				if _edd._ca < 64 {
					if _edd._ca < 0 {
						_cge[_dbc] = _bbd
						_dbc++
						_edd = nil
						break _dag
					}
					_bbd += _edd._ca
					_cge[_dbc] = _bbd
					_dbc++
					break
				}
				_bbd += _edd._ca
			}
			_edg := _bbd
		_cdb:
			for {
				var _ddd []*code
				if !_ffd {
					_ddd = _bfg._gaf
				} else {
					_ddd = _bfg._ff
				}
				_edd, _dd = _ccd.uncompressGetCode(_ddd)
				if _dd != nil {
					return 0, _dd
				}
				if _edd == nil {
					break _dag
				}
				_ccd._aga += _edd._b
				if _edd._ca < 64 {
					if _edd._ca < 0 {
						_cge[_dbc] = _bbd
						_dbc++
						break _dag
					}
					_bbd += _edd._ca
					if _bbd < _bbfc || _bbd != _edg {
						_cge[_dbc] = _bbd
						_dbc++
					}
					break _cdb
				}
				_bbd += _edd._ca
			}
			for _bbd < _bbfc && _gad[_adg] <= _bbd {
				_adg += 2
			}
			continue _dag
		case _ag:
			_adg++
			_bbd = _gad[_adg]
			_adg++
			continue _dag
		case _ec:
			_bbd = _gad[_adg] + 2
		case _cag:
			_bbd = _gad[_adg] - 2
		case _afe:
			_bbd = _gad[_adg] + 3
		case _eb:
			_bbd = _gad[_adg] - 3
		default:
			if _ccd._aga == 12 && _edd._ca == EOL {
				_ccd._aga = 0
				if _, _dd = _bfg.uncompress1d(_ccd, _gad, _bbfc); _dd != nil {
					return 0, _dd
				}
				_ccd._aga++
				if _, _dd = _bfg.uncompress1d(_ccd, _cge, _bbfc); _dd != nil {
					return 0, _dd
				}
				_gbcb, _afb := _bfg.uncompress1d(_ccd, _gad, _bbfc)
				if _afb != nil {
					return EOF, _afb
				}
				_ccd._aga++
				return _gbcb, nil
			}
			_bbd = _bbfc
			continue _dag
		}
		if _bbd <= _bbfc {
			_ffd = !_ffd
			_cge[_dbc] = _bbd
			_dbc++
			if _adg > 0 {
				_adg--
			} else {
				_adg++
			}
			for _bbd < _bbfc && _gad[_adg] <= _bbd {
				_adg += 2
			}
		}
	}
	if _cge[_dbc] != _bbfc {
		_cge[_dbc] = _bbfc
	}
	if _edd == nil {
		return EOL, nil
	}
	return _dbc, nil
}
func (_bbdg *runData) fillBuffer(_bdee int) error {
	_bbdg._feb = _bdee
	_, _ccab := _bbdg._bc.Seek(int64(_bdee), _c.SeekStart)
	if _ccab != nil {
		if _ccab == _c.EOF {
			_af.Log.Debug("\u0053\u0065\u0061\u006b\u0020\u0045\u004f\u0046")
			_bbdg._bd = -1
		} else {
			return _ccab
		}
	}
	if _ccab == nil {
		_bbdg._bd, _ccab = _bbdg._bc.Read(_bbdg._fag)
		if _ccab != nil {
			if _ccab == _c.EOF {
				_af.Log.Trace("\u0052\u0065\u0061\u0064\u0020\u0045\u004f\u0046")
				_bbdg._bd = -1
			} else {
				return _ccab
			}
		}
	}
	if _bbdg._bd > -1 && _bbdg._bd < 3 {
		for _bbdg._bd < 3 {
			_dfd, _dee := _bbdg._bc.ReadByte()
			if _dee != nil {
				if _dee == _c.EOF {
					_bbdg._fag[_bbdg._bd] = 0
				} else {
					return _dee
				}
			} else {
				_bbdg._fag[_bbdg._bd] = _dfd & 0xFF
			}
			_bbdg._bd++
		}
	}
	_bbdg._bd -= 3
	if _bbdg._bd < 0 {
		_bbdg._fag = make([]byte, len(_bbdg._fag))
		_bbdg._bd = len(_bbdg._fag) - 3
	}
	return nil
}
func _ebea(_bgc *_f.SubstreamReader) (*runData, error) {
	_add := &runData{_bc: _bgc, _aga: 0, _ce: 1}
	_fec := _cb(_cc(_dff, int(_bgc.Length())), _aaf)
	_add._fag = make([]byte, _fec)
	if _bde := _add.fillBuffer(0); _bde != nil {
		if _bde == _c.EOF {
			_add._fag = make([]byte, 10)
			_af.Log.Debug("F\u0069\u006c\u006c\u0042uf\u0066e\u0072\u0020\u0066\u0061\u0069l\u0065\u0064\u003a\u0020\u0025\u0076", _bde)
		} else {
			return nil, _bde
		}
	}
	return _add, nil
}
func (_fbe *Decoder) uncompress1d(_gcca *runData, _dg []int, _bf int) (int, error) {
	var (
		_ebb = true
		_fc  int
		_eae *code
		_dab int
		_bgb error
	)
_cca:
	for _fc < _bf {
	_de:
		for {
			if _ebb {
				_eae, _bgb = _gcca.uncompressGetCode(_fbe._gaf)
				if _bgb != nil {
					return 0, _bgb
				}
			} else {
				_eae, _bgb = _gcca.uncompressGetCode(_fbe._ff)
				if _bgb != nil {
					return 0, _bgb
				}
			}
			_gcca._aga += _eae._b
			if _eae._ca < 0 {
				break _cca
			}
			_fc += _eae._ca
			if _eae._ca < 64 {
				_ebb = !_ebb
				_dg[_dab] = _fc
				_dab++
				break _de
			}
		}
	}
	if _dg[_dab] != _bf {
		_dg[_dab] = _bf
	}
	_ebee := EOL
	if _eae != nil && _eae._ca != EOL {
		_ebee = _dab
	}
	return _ebee, nil
}
func (_cba *runData) uncompressGetCodeLittleEndian(_gea []*code) (*code, error) {
	_aad, _bbb := _cba.uncompressGetNextCodeLittleEndian()
	if _bbb != nil {
		_af.Log.Debug("\u0055n\u0063\u006fm\u0070\u0072\u0065\u0073s\u0047\u0065\u0074N\u0065\u0078\u0074\u0043\u006f\u0064\u0065\u004c\u0069tt\u006c\u0065\u0045n\u0064\u0069a\u006e\u0020\u0066\u0061\u0069\u006ce\u0064\u003a \u0025\u0076", _bbb)
		return nil, _bbb
	}
	_aad &= 0xffffff
	_cf := _aad >> (_bbgg - _ed)
	_bfb := _gea[_cf]
	if _bfb != nil && _bfb._fb {
		_cf = (_aad >> (_bbgg - _ed - _dfg)) & _fbb
		_bfb = _bfb._e[_cf]
	}
	return _bfb, nil
}
func _cc(_ac, _cd int) int {
	if _ac < _cd {
		return _cd
	}
	return _ac
}

type Decoder struct {
	_fbbb, _gb int
	_gd        *runData
	_gaf       []*code
	_ff        []*code
	_cda       []*code
}

func (_db *Decoder) fillBitmap(_ggg *_da.Bitmap, _eg int, _caf []int, _abb int) error {
	var _gbd byte
	_dad := 0
	_gfe := _ggg.GetByteIndex(_dad, _eg)
	for _abda := 0; _abda < _abb; _abda++ {
		_fa := byte(1)
		_fe := _caf[_abda]
		if (_abda & 1) == 0 {
			_fa = 0
		}
		for _dad < _fe {
			_gbd = (_gbd << 1) | _fa
			_dad++
			if (_dad & 7) == 0 {
				if _gfc := _ggg.SetByte(_gfe, _gbd); _gfc != nil {
					return _gfc
				}
				_gfe++
				_gbd = 0
			}
		}
	}
	if (_dad & 7) != 0 {
		_gbd <<= uint(8 - (_dad & 7))
		if _bbf := _ggg.SetByte(_gfe, _gbd); _bbf != nil {
			return _bbf
		}
	}
	return nil
}

const (
	_ag mmrCode = iota
	_cce
	_aa
	_gg
	_ec
	_afe
	_ccc
	_cag
	_eb
	_gc
	_df
)

func (_gcc *Decoder) initTables() (_ggb error) {
	if _gcc._gaf == nil {
		_gcc._gaf, _ggb = _gcc.createLittleEndianTable(_bbg)
		if _ggb != nil {
			return
		}
		_gcc._ff, _ggb = _gcc.createLittleEndianTable(_aedg)
		if _ggb != nil {
			return
		}
		_gcc._cda, _ggb = _gcc.createLittleEndianTable(_bb)
		if _ggb != nil {
			return
		}
	}
	return nil
}

type runData struct {
	_bc  *_f.SubstreamReader
	_aga int
	_ce  int
	_ecd int
	_fag []byte
	_feb int
	_bd  int
}
