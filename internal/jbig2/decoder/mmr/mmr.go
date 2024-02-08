package mmr

import (
	_fc "errors"
	_f "fmt"
	_fd "io"

	_ed "bitbucket.org/shenghui0779/gopdf/common"
	_g "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_d "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
)

func (_gf *code) String() string {
	return _f.Sprintf("\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064", _gf._dd, _gf._eg, _gf._ee)
}

type runData struct {
	_bfe   *_g.Reader
	_daa   int
	_bagdd int
	_ffa   int
	_dgd   []byte
	_fdd   int
	_baf   int
}

func (_agc *runData) uncompressGetNextCodeLittleEndian() (int, error) {
	_cfd := _agc._daa - _agc._bagdd
	if _cfd < 0 || _cfd > 24 {
		_dbag := (_agc._daa >> 3) - _agc._fdd
		if _dbag >= _agc._baf {
			_dbag += _agc._fdd
			if _fbd := _agc.fillBuffer(_dbag); _fbd != nil {
				return 0, _fbd
			}
			_dbag -= _agc._fdd
		}
		_efba := (uint32(_agc._dgd[_dbag]&0xFF) << 16) | (uint32(_agc._dgd[_dbag+1]&0xFF) << 8) | (uint32(_agc._dgd[_dbag+2] & 0xFF))
		_eefd := uint32(_agc._daa & 7)
		_efba <<= _eefd
		_agc._ffa = int(_efba)
	} else {
		_ddgb := _agc._bagdd & 7
		_gec := 7 - _ddgb
		if _cfd <= _gec {
			_agc._ffa <<= uint(_cfd)
		} else {
			_efe := (_agc._bagdd >> 3) + 3 - _agc._fdd
			if _efe >= _agc._baf {
				_efe += _agc._fdd
				if _aee := _agc.fillBuffer(_efe); _aee != nil {
					return 0, _aee
				}
				_efe -= _agc._fdd
			}
			_ddgb = 8 - _ddgb
			for {
				_agc._ffa <<= uint(_ddgb)
				_agc._ffa |= int(uint(_agc._dgd[_efe]) & 0xFF)
				_cfd -= _ddgb
				_efe++
				_ddgb = 8
				if !(_cfd >= 8) {
					break
				}
			}
			_agc._ffa <<= uint(_cfd)
		}
	}
	_agc._bagdd = _agc._daa
	return _agc._ffa, nil
}
func (_ebf *runData) align() { _ebf._daa = ((_ebf._daa + 7) >> 3) << 3 }
func _ebd(_acc *_g.Reader) (*runData, error) {
	_beec := &runData{_bfe: _acc, _daa: 0, _bagdd: 1}
	_fad := _cc(_ede(_cdf, int(_acc.Length())), _bga)
	_beec._dgd = make([]byte, _fad)
	if _gag := _beec.fillBuffer(0); _gag != nil {
		if _gag == _fd.EOF {
			_beec._dgd = make([]byte, 10)
			_ed.Log.Debug("F\u0069\u006c\u006c\u0042uf\u0066e\u0072\u0020\u0066\u0061\u0069l\u0065\u0064\u003a\u0020\u0025\u0076", _gag)
		} else {
			return nil, _gag
		}
	}
	return _beec, nil
}

const (
	_ea mmrCode = iota
	_de
	_ef
	_gb
	_cg
	_dea
	_fda
	_db
	_ce
	_aa
	_def
)
const (
	EOF  = -3
	_cee = -2
	EOL  = -1
	_be  = 8
	_df  = (1 << _be) - 1
	_gg  = 5
	_fe  = (1 << _gg) - 1
)

type mmrCode int

func New(r *_g.Reader, width, height int, dataOffset, dataLength int64) (*Decoder, error) {
	_dge := &Decoder{_cca: width, _aec: height}
	_eea, _ec := r.NewPartialReader(int(dataOffset), int(dataLength), false)
	if _ec != nil {
		return nil, _ec
	}
	_eef, _ec := _ebd(_eea)
	if _ec != nil {
		return nil, _ec
	}
	_, _ec = r.Seek(_eea.RelativePosition(), _fd.SeekCurrent)
	if _ec != nil {
		return nil, _ec
	}
	_dge._efg = _eef
	if _cea := _dge.initTables(); _cea != nil {
		return nil, _cea
	}
	return _dge, nil
}
func _cc(_ae, _bd int) int {
	if _ae > _bd {
		return _bd
	}
	return _ae
}
func (_ebg *Decoder) uncompress1d(_cbd *runData, _dcg []int, _dca int) (int, error) {
	var (
		_gba = true
		_gdf int
		_bb  *code
		_deg int
		_fdf error
	)
_cdad:
	for _gdf < _dca {
	_eab:
		for {
			if _gba {
				_bb, _fdf = _cbd.uncompressGetCode(_ebg._cd)
				if _fdf != nil {
					return 0, _fdf
				}
			} else {
				_bb, _fdf = _cbd.uncompressGetCode(_ebg._bg)
				if _fdf != nil {
					return 0, _fdf
				}
			}
			_cbd._daa += _bb._dd
			if _bb._ee < 0 {
				break _cdad
			}
			_gdf += _bb._ee
			if _bb._ee < 64 {
				_gba = !_gba
				_dcg[_deg] = _gdf
				_deg++
				break _eab
			}
		}
	}
	if _dcg[_deg] != _dca {
		_dcg[_deg] = _dca
	}
	_eeb := EOL
	if _bb != nil && _bb._ee != EOL {
		_eeb = _deg
	}
	return _eeb, nil
}
func _b(_cb [3]int) *code { return &code{_dd: _cb[0], _eg: _cb[1], _ee: _cb[2]} }
func (_eed *Decoder) UncompressMMR() (_eaa *_d.Bitmap, _ff error) {
	_eaa = _d.New(_eed._cca, _eed._aec)
	_bgb := make([]int, _eaa.Width+5)
	_da := make([]int, _eaa.Width+5)
	_da[0] = _eaa.Width
	_efb := 1
	var _eb int
	for _dfd := 0; _dfd < _eaa.Height; _dfd++ {
		_eb, _ff = _eed.uncompress2d(_eed._efg, _da, _efb, _bgb, _eaa.Width)
		if _ff != nil {
			return nil, _ff
		}
		if _eb == EOF {
			break
		}
		if _eb > 0 {
			_ff = _eed.fillBitmap(_eaa, _dfd, _bgb, _eb)
			if _ff != nil {
				return nil, _ff
			}
		}
		_da, _bgb = _bgb, _da
		_efb = _eb
	}
	if _ff = _eed.detectAndSkipEOL(); _ff != nil {
		return nil, _ff
	}
	_eed._efg.align()
	return _eaa, nil
}
func (_afb *Decoder) initTables() (_bag error) {
	if _afb._cd == nil {
		_afb._cd, _bag = _afb.createLittleEndianTable(_ggg)
		if _bag != nil {
			return
		}
		_afb._bg, _bag = _afb.createLittleEndianTable(_eac)
		if _bag != nil {
			return
		}
		_afb._dg, _bag = _afb.createLittleEndianTable(_dc)
		if _bag != nil {
			return
		}
	}
	return nil
}

type Decoder struct {
	_cca, _aec int
	_efg       *runData
	_cd        []*code
	_bg        []*code
	_dg        []*code
}

func (_deac *runData) uncompressGetCodeLittleEndian(_ebc []*code) (*code, error) {
	_agad, _cdb := _deac.uncompressGetNextCodeLittleEndian()
	if _cdb != nil {
		_ed.Log.Debug("\u0055n\u0063\u006fm\u0070\u0072\u0065\u0073s\u0047\u0065\u0074N\u0065\u0078\u0074\u0043\u006f\u0064\u0065\u004c\u0069tt\u006c\u0065\u0045n\u0064\u0069a\u006e\u0020\u0066\u0061\u0069\u006ce\u0064\u003a \u0025\u0076", _cdb)
		return nil, _cdb
	}
	_agad &= 0xffffff
	_aca := _agad >> (_bdc - _be)
	_cfa := _ebc[_aca]
	if _cfa != nil && _cfa._c {
		_aca = (_agad >> (_bdc - _be - _gg)) & _fe
		_cfa = _cfa._a[_aca]
	}
	return _cfa, nil
}
func _ede(_ag, _cf int) int {
	if _ag < _cf {
		return _cf
	}
	return _ag
}
func (_cgd *Decoder) uncompress2d(_ac *runData, _bbg []int, _bee int, _bage []int, _agf int) (int, error) {
	var (
		_ecb  int
		_ffe  int
		_dec  int
		_ecd  = true
		_ffb  error
		_cbdb *code
	)
	_bbg[_bee] = _agf
	_bbg[_bee+1] = _agf
	_bbg[_bee+2] = _agf + 1
	_bbg[_bee+3] = _agf + 1
_eaf:
	for _dec < _agf {
		_cbdb, _ffb = _ac.uncompressGetCode(_cgd._dg)
		if _ffb != nil {
			return EOL, nil
		}
		if _cbdb == nil {
			_ac._daa++
			break _eaf
		}
		_ac._daa += _cbdb._dd
		switch mmrCode(_cbdb._ee) {
		case _ef:
			_dec = _bbg[_ecb]
		case _gb:
			_dec = _bbg[_ecb] + 1
		case _fda:
			_dec = _bbg[_ecb] - 1
		case _de:
			for {
				var _gbg []*code
				if _ecd {
					_gbg = _cgd._cd
				} else {
					_gbg = _cgd._bg
				}
				_cbdb, _ffb = _ac.uncompressGetCode(_gbg)
				if _ffb != nil {
					return 0, _ffb
				}
				if _cbdb == nil {
					break _eaf
				}
				_ac._daa += _cbdb._dd
				if _cbdb._ee < 64 {
					if _cbdb._ee < 0 {
						_bage[_ffe] = _dec
						_ffe++
						_cbdb = nil
						break _eaf
					}
					_dec += _cbdb._ee
					_bage[_ffe] = _dec
					_ffe++
					break
				}
				_dec += _cbdb._ee
			}
			_ead := _dec
		_fb:
			for {
				var _bcf []*code
				if !_ecd {
					_bcf = _cgd._cd
				} else {
					_bcf = _cgd._bg
				}
				_cbdb, _ffb = _ac.uncompressGetCode(_bcf)
				if _ffb != nil {
					return 0, _ffb
				}
				if _cbdb == nil {
					break _eaf
				}
				_ac._daa += _cbdb._dd
				if _cbdb._ee < 64 {
					if _cbdb._ee < 0 {
						_bage[_ffe] = _dec
						_ffe++
						break _eaf
					}
					_dec += _cbdb._ee
					if _dec < _agf || _dec != _ead {
						_bage[_ffe] = _dec
						_ffe++
					}
					break _fb
				}
				_dec += _cbdb._ee
			}
			for _dec < _agf && _bbg[_ecb] <= _dec {
				_ecb += 2
			}
			continue _eaf
		case _ea:
			_ecb++
			_dec = _bbg[_ecb]
			_ecb++
			continue _eaf
		case _cg:
			_dec = _bbg[_ecb] + 2
		case _db:
			_dec = _bbg[_ecb] - 2
		case _dea:
			_dec = _bbg[_ecb] + 3
		case _ce:
			_dec = _bbg[_ecb] - 3
		default:
			if _ac._daa == 12 && _cbdb._ee == EOL {
				_ac._daa = 0
				if _, _ffb = _cgd.uncompress1d(_ac, _bbg, _agf); _ffb != nil {
					return 0, _ffb
				}
				_ac._daa++
				if _, _ffb = _cgd.uncompress1d(_ac, _bage, _agf); _ffb != nil {
					return 0, _ffb
				}
				_cbe, _ecc := _cgd.uncompress1d(_ac, _bbg, _agf)
				if _ecc != nil {
					return EOF, _ecc
				}
				_ac._daa++
				return _cbe, nil
			}
			_dec = _agf
			continue _eaf
		}
		if _dec <= _agf {
			_ecd = !_ecd
			_bage[_ffe] = _dec
			_ffe++
			if _ecb > 0 {
				_ecb--
			} else {
				_ecb++
			}
			for _dec < _agf && _bbg[_ecb] <= _dec {
				_ecb += 2
			}
		}
	}
	if _bage[_ffe] != _agf {
		_bage[_ffe] = _agf
	}
	if _cbdb == nil {
		return EOL, nil
	}
	return _ffe, nil
}
func (_ddg *Decoder) detectAndSkipEOL() error {
	for {
		_ab, _ba := _ddg._efg.uncompressGetCode(_ddg._dg)
		if _ba != nil {
			return _ba
		}
		if _ab != nil && _ab._ee == EOL {
			_ddg._efg._daa += _ab._dd
		} else {
			return nil
		}
	}
}

var (
	_dc  = [][3]int{{4, 0x1, int(_ea)}, {3, 0x1, int(_de)}, {1, 0x1, int(_ef)}, {3, 0x3, int(_gb)}, {6, 0x3, int(_cg)}, {7, 0x3, int(_dea)}, {3, 0x2, int(_fda)}, {6, 0x2, int(_db)}, {7, 0x2, int(_ce)}, {10, 0xf, int(_aa)}, {12, 0xf, int(_def)}, {12, 0x1, int(EOL)}}
	_ggg = [][3]int{{4, 0x07, 2}, {4, 0x08, 3}, {4, 0x0B, 4}, {4, 0x0C, 5}, {4, 0x0E, 6}, {4, 0x0F, 7}, {5, 0x12, 128}, {5, 0x13, 8}, {5, 0x14, 9}, {5, 0x1B, 64}, {5, 0x07, 10}, {5, 0x08, 11}, {6, 0x17, 192}, {6, 0x18, 1664}, {6, 0x2A, 16}, {6, 0x2B, 17}, {6, 0x03, 13}, {6, 0x34, 14}, {6, 0x35, 15}, {6, 0x07, 1}, {6, 0x08, 12}, {7, 0x13, 26}, {7, 0x17, 21}, {7, 0x18, 28}, {7, 0x24, 27}, {7, 0x27, 18}, {7, 0x28, 24}, {7, 0x2B, 25}, {7, 0x03, 22}, {7, 0x37, 256}, {7, 0x04, 23}, {7, 0x08, 20}, {7, 0xC, 19}, {8, 0x12, 33}, {8, 0x13, 34}, {8, 0x14, 35}, {8, 0x15, 36}, {8, 0x16, 37}, {8, 0x17, 38}, {8, 0x1A, 31}, {8, 0x1B, 32}, {8, 0x02, 29}, {8, 0x24, 53}, {8, 0x25, 54}, {8, 0x28, 39}, {8, 0x29, 40}, {8, 0x2A, 41}, {8, 0x2B, 42}, {8, 0x2C, 43}, {8, 0x2D, 44}, {8, 0x03, 30}, {8, 0x32, 61}, {8, 0x33, 62}, {8, 0x34, 63}, {8, 0x35, 0}, {8, 0x36, 320}, {8, 0x37, 384}, {8, 0x04, 45}, {8, 0x4A, 59}, {8, 0x4B, 60}, {8, 0x5, 46}, {8, 0x52, 49}, {8, 0x53, 50}, {8, 0x54, 51}, {8, 0x55, 52}, {8, 0x58, 55}, {8, 0x59, 56}, {8, 0x5A, 57}, {8, 0x5B, 58}, {8, 0x64, 448}, {8, 0x65, 512}, {8, 0x67, 640}, {8, 0x68, 576}, {8, 0x0A, 47}, {8, 0x0B, 48}, {9, 0x01, _cee}, {9, 0x98, 1472}, {9, 0x99, 1536}, {9, 0x9A, 1600}, {9, 0x9B, 1728}, {9, 0xCC, 704}, {9, 0xCD, 768}, {9, 0xD2, 832}, {9, 0xD3, 896}, {9, 0xD4, 960}, {9, 0xD5, 1024}, {9, 0xD6, 1088}, {9, 0xD7, 1152}, {9, 0xD8, 1216}, {9, 0xD9, 1280}, {9, 0xDA, 1344}, {9, 0xDB, 1408}, {10, 0x01, _cee}, {11, 0x01, _cee}, {11, 0x08, 1792}, {11, 0x0C, 1856}, {11, 0x0D, 1920}, {12, 0x00, EOF}, {12, 0x01, EOL}, {12, 0x12, 1984}, {12, 0x13, 2048}, {12, 0x14, 2112}, {12, 0x15, 2176}, {12, 0x16, 2240}, {12, 0x17, 2304}, {12, 0x1C, 2368}, {12, 0x1D, 2432}, {12, 0x1E, 2496}, {12, 0x1F, 2560}}
	_eac = [][3]int{{2, 0x02, 3}, {2, 0x03, 2}, {3, 0x02, 1}, {3, 0x03, 4}, {4, 0x02, 6}, {4, 0x03, 5}, {5, 0x03, 7}, {6, 0x04, 9}, {6, 0x05, 8}, {7, 0x04, 10}, {7, 0x05, 11}, {7, 0x07, 12}, {8, 0x04, 13}, {8, 0x07, 14}, {9, 0x01, _cee}, {9, 0x18, 15}, {10, 0x01, _cee}, {10, 0x17, 16}, {10, 0x18, 17}, {10, 0x37, 0}, {10, 0x08, 18}, {10, 0x0F, 64}, {11, 0x01, _cee}, {11, 0x17, 24}, {11, 0x18, 25}, {11, 0x28, 23}, {11, 0x37, 22}, {11, 0x67, 19}, {11, 0x68, 20}, {11, 0x6C, 21}, {11, 0x08, 1792}, {11, 0x0C, 1856}, {11, 0x0D, 1920}, {12, 0x00, EOF}, {12, 0x01, EOL}, {12, 0x12, 1984}, {12, 0x13, 2048}, {12, 0x14, 2112}, {12, 0x15, 2176}, {12, 0x16, 2240}, {12, 0x17, 2304}, {12, 0x1C, 2368}, {12, 0x1D, 2432}, {12, 0x1E, 2496}, {12, 0x1F, 2560}, {12, 0x24, 52}, {12, 0x27, 55}, {12, 0x28, 56}, {12, 0x2B, 59}, {12, 0x2C, 60}, {12, 0x33, 320}, {12, 0x34, 384}, {12, 0x35, 448}, {12, 0x37, 53}, {12, 0x38, 54}, {12, 0x52, 50}, {12, 0x53, 51}, {12, 0x54, 44}, {12, 0x55, 45}, {12, 0x56, 46}, {12, 0x57, 47}, {12, 0x58, 57}, {12, 0x59, 58}, {12, 0x5A, 61}, {12, 0x5B, 256}, {12, 0x64, 48}, {12, 0x65, 49}, {12, 0x66, 62}, {12, 0x67, 63}, {12, 0x68, 30}, {12, 0x69, 31}, {12, 0x6A, 32}, {12, 0x6B, 33}, {12, 0x6C, 40}, {12, 0x6D, 41}, {12, 0xC8, 128}, {12, 0xC9, 192}, {12, 0xCA, 26}, {12, 0xCB, 27}, {12, 0xCC, 28}, {12, 0xCD, 29}, {12, 0xD2, 34}, {12, 0xD3, 35}, {12, 0xD4, 36}, {12, 0xD5, 37}, {12, 0xD6, 38}, {12, 0xD7, 39}, {12, 0xDA, 42}, {12, 0xDB, 43}, {13, 0x4A, 640}, {13, 0x4B, 704}, {13, 0x4C, 768}, {13, 0x4D, 832}, {13, 0x52, 1280}, {13, 0x53, 1344}, {13, 0x54, 1408}, {13, 0x55, 1472}, {13, 0x5A, 1536}, {13, 0x5B, 1600}, {13, 0x64, 1664}, {13, 0x65, 1728}, {13, 0x6C, 512}, {13, 0x6D, 576}, {13, 0x72, 896}, {13, 0x73, 960}, {13, 0x74, 1024}, {13, 0x75, 1088}, {13, 0x76, 1152}, {13, 0x77, 1216}}
)

func (_gc *Decoder) fillBitmap(_fa *_d.Bitmap, _bea int, _agg []int, _bc int) error {
	var _afe byte
	_ccb := 0
	_gcf := _fa.GetByteIndex(_ccb, _bea)
	for _deb := 0; _deb < _bc; _deb++ {
		_aga := byte(1)
		_cda := _agg[_deb]
		if (_deb & 1) == 0 {
			_aga = 0
		}
		for _ccb < _cda {
			_afe = (_afe << 1) | _aga
			_ccb++
			if (_ccb & 7) == 0 {
				if _dgf := _fa.SetByte(_gcf, _afe); _dgf != nil {
					return _dgf
				}
				_gcf++
				_afe = 0
			}
		}
	}
	if (_ccb & 7) != 0 {
		_afe <<= uint(8 - (_ccb & 7))
		if _bac := _fa.SetByte(_gcf, _afe); _bac != nil {
			return _bac
		}
	}
	return nil
}
func (_fga *runData) fillBuffer(_eeg int) error {
	_fga._fdd = _eeg
	_, _ggf := _fga._bfe.Seek(int64(_eeg), _fd.SeekStart)
	if _ggf != nil {
		if _ggf == _fd.EOF {
			_ed.Log.Debug("\u0053\u0065\u0061\u006b\u0020\u0045\u004f\u0046")
			_fga._baf = -1
		} else {
			return _ggf
		}
	}
	if _ggf == nil {
		_fga._baf, _ggf = _fga._bfe.Read(_fga._dgd)
		if _ggf != nil {
			if _ggf == _fd.EOF {
				_ed.Log.Trace("\u0052\u0065\u0061\u0064\u0020\u0045\u004f\u0046")
				_fga._baf = -1
			} else {
				return _ggf
			}
		}
	}
	if _fga._baf > -1 && _fga._baf < 3 {
		for _fga._baf < 3 {
			_gfe, _ggb := _fga._bfe.ReadByte()
			if _ggb != nil {
				if _ggb == _fd.EOF {
					_fga._dgd[_fga._baf] = 0
				} else {
					return _ggb
				}
			} else {
				_fga._dgd[_fga._baf] = _gfe & 0xFF
			}
			_fga._baf++
		}
	}
	_fga._baf -= 3
	if _fga._baf < 0 {
		_fga._dgd = make([]byte, len(_fga._dgd))
		_fga._baf = len(_fga._dgd) - 3
	}
	return nil
}

type code struct {
	_dd int
	_eg int
	_ee int
	_a  []*code
	_c  bool
}

const (
	_bga int  = 1024 << 7
	_cdf int  = 3
	_bdc uint = 24
)

func (_fdfe *runData) uncompressGetCode(_gggb []*code) (*code, error) {
	return _fdfe.uncompressGetCodeLittleEndian(_gggb)
}
func (_ga *Decoder) createLittleEndianTable(_eba [][3]int) ([]*code, error) {
	_dba := make([]*code, _df+1)
	for _dbc := 0; _dbc < len(_eba); _dbc++ {
		_dfb := _b(_eba[_dbc])
		if _dfb._dd <= _be {
			_cef := _be - _dfb._dd
			_agb := _dfb._eg << uint(_cef)
			for _af := (1 << uint(_cef)) - 1; _af >= 0; _af-- {
				_gd := _agb | _af
				_dba[_gd] = _dfb
			}
		} else {
			_bf := _dfb._eg >> uint(_dfb._dd-_be)
			if _dba[_bf] == nil {
				var _fg = _b([3]int{})
				_fg._a = make([]*code, _fe+1)
				_dba[_bf] = _fg
			}
			if _dfb._dd <= _be+_gg {
				_gaf := _be + _gg - _dfb._dd
				_ega := (_dfb._eg << uint(_gaf)) & _fe
				_dba[_bf]._c = true
				for _ge := (1 << uint(_gaf)) - 1; _ge >= 0; _ge-- {
					_dba[_bf]._a[_ega|_ge] = _dfb
				}
			} else {
				return nil, _fc.New("\u0043\u006f\u0064\u0065\u0020\u0074a\u0062\u006c\u0065\u0020\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077\u0020i\u006e\u0020\u004d\u004d\u0052\u0044\u0065c\u006f\u0064\u0065\u0072")
			}
		}
	}
	return _dba, nil
}
