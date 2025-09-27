package ccittfax

import (
	_c "errors"
	_b "io"
	_g "math"

	_fa "unitechio/gopdf/gopdf/internal/bitwise"
)

func (_baec *Decoder) fetch() error {
	if _baec._ecf == -1 {
		return nil
	}
	if _baec._bb < _baec._ecf {
		return nil
	}
	_baec._ecf = 0
	_ce := _baec.decodeRow()
	if _ce != nil {
		if !_c.Is(_ce, _b.EOF) {
			return _ce
		}
		if _baec._ecf != 0 {
			return _ce
		}
		_baec._ecf = -1
	}
	_baec._bb = 0
	return nil
}

func init() {
	_d = &treeNode{_cdgg: true, _efe: _ge}
	_a = &treeNode{_efe: _df, _ccdfb: _d}
	_a._dcecg = _a
	_e = &tree{_adcga: &treeNode{}}
	if _dcg := _e.fillWithNode(12, 0, _a); _dcg != nil {
		panic(_dcg.Error())
	}
	if _ba := _e.fillWithNode(12, 1, _d); _ba != nil {
		panic(_ba.Error())
	}
	_fd = &tree{_adcga: &treeNode{}}
	for _ae := 0; _ae < len(_ff); _ae++ {
		for _fb := 0; _fb < len(_ff[_ae]); _fb++ {
			if _bcc := _fd.fill(_ae+2, int(_ff[_ae][_fb]), int(_egd[_ae][_fb])); _bcc != nil {
				panic(_bcc.Error())
			}
		}
	}
	if _ef := _fd.fillWithNode(12, 0, _a); _ef != nil {
		panic(_ef.Error())
	}
	if _fc := _fd.fillWithNode(12, 1, _d); _fc != nil {
		panic(_fc.Error())
	}
	_bc = &tree{_adcga: &treeNode{}}
	for _ab := 0; _ab < len(_ggb); _ab++ {
		for _ag := 0; _ag < len(_ggb[_ab]); _ag++ {
			if _cc := _bc.fill(_ab+4, int(_ggb[_ab][_ag]), int(_aa[_ab][_ag])); _cc != nil {
				panic(_cc.Error())
			}
		}
	}
	if _fda := _bc.fillWithNode(12, 0, _a); _fda != nil {
		panic(_fda.Error())
	}
	if _fab := _bc.fillWithNode(12, 1, _d); _fab != nil {
		panic(_fab.Error())
	}
	_dc = &tree{_adcga: &treeNode{}}
	if _gg := _dc.fill(4, 1, _dce); _gg != nil {
		panic(_gg.Error())
	}
	if _eg := _dc.fill(3, 1, _fdd); _eg != nil {
		panic(_eg.Error())
	}
	if _agf := _dc.fill(1, 1, 0); _agf != nil {
		panic(_agf.Error())
	}
	if _cg := _dc.fill(3, 3, 1); _cg != nil {
		panic(_cg.Error())
	}
	if _gd := _dc.fill(6, 3, 2); _gd != nil {
		panic(_gd.Error())
	}
	if _bd := _dc.fill(7, 3, 3); _bd != nil {
		panic(_bd.Error())
	}
	if _geg := _dc.fill(3, 2, -1); _geg != nil {
		panic(_geg.Error())
	}
	if _af := _dc.fill(6, 2, -2); _af != nil {
		panic(_af.Error())
	}
	if _cb := _dc.fill(7, 2, -3); _cb != nil {
		panic(_cb.Error())
	}
}

func _fbc(_ccga []byte, _eab int) int {
	if _eab >= len(_ccga) {
		return _eab
	}
	if _eab < -1 {
		_eab = -1
	}
	var _eeba byte
	if _eab > -1 {
		_eeba = _ccga[_eab]
	} else {
		_eeba = _gbaf
	}
	_dgd := _eab + 1
	for _dgd < len(_ccga) {
		if _ccga[_dgd] != _eeba {
			break
		}
		_dgd++
	}
	return _dgd
}

func _fdc(_gcaa, _cbcg []byte, _ace int, _adcg bool) int {
	_dfg := _fbc(_cbcg, _ace)
	if _dfg < len(_cbcg) && (_ace == -1 && _cbcg[_dfg] == _gbaf || _ace >= 0 && _ace < len(_gcaa) && _gcaa[_ace] == _cbcg[_dfg] || _ace >= len(_gcaa) && _adcg && _cbcg[_dfg] == _gbaf || _ace >= len(_gcaa) && !_adcg && _cbcg[_dfg] == _geadg) {
		_dfg = _fbc(_cbcg, _dfg)
	}
	return _dfg
}

func (_ed *Decoder) decodeRow() (_dab error) {
	if !_ed._fe && _ed._gba > 0 && _ed._gba == _ed._gde {
		return _b.EOF
	}
	switch _ed._fcdd {
	case _ccf:
		_dab = _ed.decodeRowType2()
	case _ec:
		_dab = _ed.decodeRowType4()
	case _gcf:
		_dab = _ed.decodeRowType6()
	}
	if _dab != nil {
		return _dab
	}
	_dffa := 0
	_eba := true
	_ed._db = 0
	for _dcb := 0; _dcb < _ed._bae; _dcb++ {
		_cgc := _ed._dcd
		if _dcb != _ed._bae {
			_cgc = _ed._gbd[_dcb]
		}
		if _cgc > _ed._dcd {
			_cgc = _ed._dcd
		}
		_fgg := _dffa / 8
		for _dffa%8 != 0 && _cgc-_dffa > 0 {
			var _bbf byte
			if !_eba {
				_bbf = 1 << uint(7-(_dffa%8))
			}
			_ed._gbb[_fgg] |= _bbf
			_dffa++
		}
		if _dffa%8 == 0 {
			_fgg = _dffa / 8
			var _fagf byte
			if !_eba {
				_fagf = 0xff
			}
			for _cgc-_dffa > 7 {
				_ed._gbb[_fgg] = _fagf
				_dffa += 8
				_fgg++
			}
		}
		for _cgc-_dffa > 0 {
			if _dffa%8 == 0 {
				_ed._gbb[_fgg] = 0
			}
			var _beg byte
			if !_eba {
				_beg = 1 << uint(7-(_dffa%8))
			}
			_ed._gbb[_fgg] |= _beg
			_dffa++
		}
		_eba = !_eba
	}
	if _dffa != _ed._dcd {
		return _c.New("\u0073\u0075\u006d\u0020\u006f\u0066 \u0072\u0075\u006e\u002d\u006c\u0065\u006e\u0067\u0074\u0068\u0073\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074 \u0065\u0071\u0075\u0061\u006c\u0020\u0073\u0063\u0061\u006e\u0020\u006c\u0069\u006ee\u0020w\u0069\u0064\u0074\u0068")
	}
	_ed._ecf = (_dffa + 7) / 8
	_ed._gde++
	return nil
}

func _abfb(_cdd []byte, _bgad, _afc, _bdaa int) ([]byte, int) {
	_ccgg := _bdfg(_afc, _bdaa)
	_cdd, _bgad = _bdd(_cdd, _bgad, _ccgg)
	return _cdd, _bgad
}

func (_bgaa *Encoder) appendEncodedRow(_gaa, _ggab []byte, _cdca int) []byte {
	if len(_gaa) > 0 && _cdca != 0 && !_bgaa.EncodedByteAlign {
		_gaa[len(_gaa)-1] = _gaa[len(_gaa)-1] | _ggab[0]
		_gaa = append(_gaa, _ggab[1:]...)
	} else {
		_gaa = append(_gaa, _ggab...)
	}
	return _gaa
}

func (_acg *Decoder) getNextChangingElement(_abf int, _ccdf bool) int {
	_geb := 0
	if !_ccdf {
		_geb = 1
	}
	_gca := int(uint32(_acg._db)&0xFFFFFFFE) + _geb
	if _gca > 2 {
		_gca -= 2
	}
	if _abf == 0 {
		return _gca
	}
	for _cgde := _gca; _cgde < _acg._gdc; _cgde += 2 {
		if _abf < _acg._feb[_cgde] {
			_acg._db = _cgde
			return _cgde
		}
	}
	return -1
}

func init() {
	_gc = make(map[int]code)
	_gc[0] = code{Code: 13<<8 | 3<<6, BitsWritten: 10}
	_gc[1] = code{Code: 2 << (5 + 8), BitsWritten: 3}
	_gc[2] = code{Code: 3 << (6 + 8), BitsWritten: 2}
	_gc[3] = code{Code: 2 << (6 + 8), BitsWritten: 2}
	_gc[4] = code{Code: 3 << (5 + 8), BitsWritten: 3}
	_gc[5] = code{Code: 3 << (4 + 8), BitsWritten: 4}
	_gc[6] = code{Code: 2 << (4 + 8), BitsWritten: 4}
	_gc[7] = code{Code: 3 << (3 + 8), BitsWritten: 5}
	_gc[8] = code{Code: 5 << (2 + 8), BitsWritten: 6}
	_gc[9] = code{Code: 4 << (2 + 8), BitsWritten: 6}
	_gc[10] = code{Code: 4 << (1 + 8), BitsWritten: 7}
	_gc[11] = code{Code: 5 << (1 + 8), BitsWritten: 7}
	_gc[12] = code{Code: 7 << (1 + 8), BitsWritten: 7}
	_gc[13] = code{Code: 4 << 8, BitsWritten: 8}
	_gc[14] = code{Code: 7 << 8, BitsWritten: 8}
	_gc[15] = code{Code: 12 << 8, BitsWritten: 9}
	_gc[16] = code{Code: 5<<8 | 3<<6, BitsWritten: 10}
	_gc[17] = code{Code: 6 << 8, BitsWritten: 10}
	_gc[18] = code{Code: 2 << 8, BitsWritten: 10}
	_gc[19] = code{Code: 12<<8 | 7<<5, BitsWritten: 11}
	_gc[20] = code{Code: 13 << 8, BitsWritten: 11}
	_gc[21] = code{Code: 13<<8 | 4<<5, BitsWritten: 11}
	_gc[22] = code{Code: 6<<8 | 7<<5, BitsWritten: 11}
	_gc[23] = code{Code: 5 << 8, BitsWritten: 11}
	_gc[24] = code{Code: 2<<8 | 7<<5, BitsWritten: 11}
	_gc[25] = code{Code: 3 << 8, BitsWritten: 11}
	_gc[26] = code{Code: 12<<8 | 10<<4, BitsWritten: 12}
	_gc[27] = code{Code: 12<<8 | 11<<4, BitsWritten: 12}
	_gc[28] = code{Code: 12<<8 | 12<<4, BitsWritten: 12}
	_gc[29] = code{Code: 12<<8 | 13<<4, BitsWritten: 12}
	_gc[30] = code{Code: 6<<8 | 8<<4, BitsWritten: 12}
	_gc[31] = code{Code: 6<<8 | 9<<4, BitsWritten: 12}
	_gc[32] = code{Code: 6<<8 | 10<<4, BitsWritten: 12}
	_gc[33] = code{Code: 6<<8 | 11<<4, BitsWritten: 12}
	_gc[34] = code{Code: 13<<8 | 2<<4, BitsWritten: 12}
	_gc[35] = code{Code: 13<<8 | 3<<4, BitsWritten: 12}
	_gc[36] = code{Code: 13<<8 | 4<<4, BitsWritten: 12}
	_gc[37] = code{Code: 13<<8 | 5<<4, BitsWritten: 12}
	_gc[38] = code{Code: 13<<8 | 6<<4, BitsWritten: 12}
	_gc[39] = code{Code: 13<<8 | 7<<4, BitsWritten: 12}
	_gc[40] = code{Code: 6<<8 | 12<<4, BitsWritten: 12}
	_gc[41] = code{Code: 6<<8 | 13<<4, BitsWritten: 12}
	_gc[42] = code{Code: 13<<8 | 10<<4, BitsWritten: 12}
	_gc[43] = code{Code: 13<<8 | 11<<4, BitsWritten: 12}
	_gc[44] = code{Code: 5<<8 | 4<<4, BitsWritten: 12}
	_gc[45] = code{Code: 5<<8 | 5<<4, BitsWritten: 12}
	_gc[46] = code{Code: 5<<8 | 6<<4, BitsWritten: 12}
	_gc[47] = code{Code: 5<<8 | 7<<4, BitsWritten: 12}
	_gc[48] = code{Code: 6<<8 | 4<<4, BitsWritten: 12}
	_gc[49] = code{Code: 6<<8 | 5<<4, BitsWritten: 12}
	_gc[50] = code{Code: 5<<8 | 2<<4, BitsWritten: 12}
	_gc[51] = code{Code: 5<<8 | 3<<4, BitsWritten: 12}
	_gc[52] = code{Code: 2<<8 | 4<<4, BitsWritten: 12}
	_gc[53] = code{Code: 3<<8 | 7<<4, BitsWritten: 12}
	_gc[54] = code{Code: 3<<8 | 8<<4, BitsWritten: 12}
	_gc[55] = code{Code: 2<<8 | 7<<4, BitsWritten: 12}
	_gc[56] = code{Code: 2<<8 | 8<<4, BitsWritten: 12}
	_gc[57] = code{Code: 5<<8 | 8<<4, BitsWritten: 12}
	_gc[58] = code{Code: 5<<8 | 9<<4, BitsWritten: 12}
	_gc[59] = code{Code: 2<<8 | 11<<4, BitsWritten: 12}
	_gc[60] = code{Code: 2<<8 | 12<<4, BitsWritten: 12}
	_gc[61] = code{Code: 5<<8 | 10<<4, BitsWritten: 12}
	_gc[62] = code{Code: 6<<8 | 6<<4, BitsWritten: 12}
	_gc[63] = code{Code: 6<<8 | 7<<4, BitsWritten: 12}
	_gegg = make(map[int]code)
	_gegg[0] = code{Code: 53 << 8, BitsWritten: 8}
	_gegg[1] = code{Code: 7 << (2 + 8), BitsWritten: 6}
	_gegg[2] = code{Code: 7 << (4 + 8), BitsWritten: 4}
	_gegg[3] = code{Code: 8 << (4 + 8), BitsWritten: 4}
	_gegg[4] = code{Code: 11 << (4 + 8), BitsWritten: 4}
	_gegg[5] = code{Code: 12 << (4 + 8), BitsWritten: 4}
	_gegg[6] = code{Code: 14 << (4 + 8), BitsWritten: 4}
	_gegg[7] = code{Code: 15 << (4 + 8), BitsWritten: 4}
	_gegg[8] = code{Code: 19 << (3 + 8), BitsWritten: 5}
	_gegg[9] = code{Code: 20 << (3 + 8), BitsWritten: 5}
	_gegg[10] = code{Code: 7 << (3 + 8), BitsWritten: 5}
	_gegg[11] = code{Code: 8 << (3 + 8), BitsWritten: 5}
	_gegg[12] = code{Code: 8 << (2 + 8), BitsWritten: 6}
	_gegg[13] = code{Code: 3 << (2 + 8), BitsWritten: 6}
	_gegg[14] = code{Code: 52 << (2 + 8), BitsWritten: 6}
	_gegg[15] = code{Code: 53 << (2 + 8), BitsWritten: 6}
	_gegg[16] = code{Code: 42 << (2 + 8), BitsWritten: 6}
	_gegg[17] = code{Code: 43 << (2 + 8), BitsWritten: 6}
	_gegg[18] = code{Code: 39 << (1 + 8), BitsWritten: 7}
	_gegg[19] = code{Code: 12 << (1 + 8), BitsWritten: 7}
	_gegg[20] = code{Code: 8 << (1 + 8), BitsWritten: 7}
	_gegg[21] = code{Code: 23 << (1 + 8), BitsWritten: 7}
	_gegg[22] = code{Code: 3 << (1 + 8), BitsWritten: 7}
	_gegg[23] = code{Code: 4 << (1 + 8), BitsWritten: 7}
	_gegg[24] = code{Code: 40 << (1 + 8), BitsWritten: 7}
	_gegg[25] = code{Code: 43 << (1 + 8), BitsWritten: 7}
	_gegg[26] = code{Code: 19 << (1 + 8), BitsWritten: 7}
	_gegg[27] = code{Code: 36 << (1 + 8), BitsWritten: 7}
	_gegg[28] = code{Code: 24 << (1 + 8), BitsWritten: 7}
	_gegg[29] = code{Code: 2 << 8, BitsWritten: 8}
	_gegg[30] = code{Code: 3 << 8, BitsWritten: 8}
	_gegg[31] = code{Code: 26 << 8, BitsWritten: 8}
	_gegg[32] = code{Code: 27 << 8, BitsWritten: 8}
	_gegg[33] = code{Code: 18 << 8, BitsWritten: 8}
	_gegg[34] = code{Code: 19 << 8, BitsWritten: 8}
	_gegg[35] = code{Code: 20 << 8, BitsWritten: 8}
	_gegg[36] = code{Code: 21 << 8, BitsWritten: 8}
	_gegg[37] = code{Code: 22 << 8, BitsWritten: 8}
	_gegg[38] = code{Code: 23 << 8, BitsWritten: 8}
	_gegg[39] = code{Code: 40 << 8, BitsWritten: 8}
	_gegg[40] = code{Code: 41 << 8, BitsWritten: 8}
	_gegg[41] = code{Code: 42 << 8, BitsWritten: 8}
	_gegg[42] = code{Code: 43 << 8, BitsWritten: 8}
	_gegg[43] = code{Code: 44 << 8, BitsWritten: 8}
	_gegg[44] = code{Code: 45 << 8, BitsWritten: 8}
	_gegg[45] = code{Code: 4 << 8, BitsWritten: 8}
	_gegg[46] = code{Code: 5 << 8, BitsWritten: 8}
	_gegg[47] = code{Code: 10 << 8, BitsWritten: 8}
	_gegg[48] = code{Code: 11 << 8, BitsWritten: 8}
	_gegg[49] = code{Code: 82 << 8, BitsWritten: 8}
	_gegg[50] = code{Code: 83 << 8, BitsWritten: 8}
	_gegg[51] = code{Code: 84 << 8, BitsWritten: 8}
	_gegg[52] = code{Code: 85 << 8, BitsWritten: 8}
	_gegg[53] = code{Code: 36 << 8, BitsWritten: 8}
	_gegg[54] = code{Code: 37 << 8, BitsWritten: 8}
	_gegg[55] = code{Code: 88 << 8, BitsWritten: 8}
	_gegg[56] = code{Code: 89 << 8, BitsWritten: 8}
	_gegg[57] = code{Code: 90 << 8, BitsWritten: 8}
	_gegg[58] = code{Code: 91 << 8, BitsWritten: 8}
	_gegg[59] = code{Code: 74 << 8, BitsWritten: 8}
	_gegg[60] = code{Code: 75 << 8, BitsWritten: 8}
	_gegg[61] = code{Code: 50 << 8, BitsWritten: 8}
	_gegg[62] = code{Code: 51 << 8, BitsWritten: 8}
	_gegg[63] = code{Code: 52 << 8, BitsWritten: 8}
	_gdf = make(map[int]code)
	_gdf[64] = code{Code: 3<<8 | 3<<6, BitsWritten: 10}
	_gdf[128] = code{Code: 12<<8 | 8<<4, BitsWritten: 12}
	_gdf[192] = code{Code: 12<<8 | 9<<4, BitsWritten: 12}
	_gdf[256] = code{Code: 5<<8 | 11<<4, BitsWritten: 12}
	_gdf[320] = code{Code: 3<<8 | 3<<4, BitsWritten: 12}
	_gdf[384] = code{Code: 3<<8 | 4<<4, BitsWritten: 12}
	_gdf[448] = code{Code: 3<<8 | 5<<4, BitsWritten: 12}
	_gdf[512] = code{Code: 3<<8 | 12<<3, BitsWritten: 13}
	_gdf[576] = code{Code: 3<<8 | 13<<3, BitsWritten: 13}
	_gdf[640] = code{Code: 2<<8 | 10<<3, BitsWritten: 13}
	_gdf[704] = code{Code: 2<<8 | 11<<3, BitsWritten: 13}
	_gdf[768] = code{Code: 2<<8 | 12<<3, BitsWritten: 13}
	_gdf[832] = code{Code: 2<<8 | 13<<3, BitsWritten: 13}
	_gdf[896] = code{Code: 3<<8 | 18<<3, BitsWritten: 13}
	_gdf[960] = code{Code: 3<<8 | 19<<3, BitsWritten: 13}
	_gdf[1024] = code{Code: 3<<8 | 20<<3, BitsWritten: 13}
	_gdf[1088] = code{Code: 3<<8 | 21<<3, BitsWritten: 13}
	_gdf[1152] = code{Code: 3<<8 | 22<<3, BitsWritten: 13}
	_gdf[1216] = code{Code: 119 << 3, BitsWritten: 13}
	_gdf[1280] = code{Code: 2<<8 | 18<<3, BitsWritten: 13}
	_gdf[1344] = code{Code: 2<<8 | 19<<3, BitsWritten: 13}
	_gdf[1408] = code{Code: 2<<8 | 20<<3, BitsWritten: 13}
	_gdf[1472] = code{Code: 2<<8 | 21<<3, BitsWritten: 13}
	_gdf[1536] = code{Code: 2<<8 | 26<<3, BitsWritten: 13}
	_gdf[1600] = code{Code: 2<<8 | 27<<3, BitsWritten: 13}
	_gdf[1664] = code{Code: 3<<8 | 4<<3, BitsWritten: 13}
	_gdf[1728] = code{Code: 3<<8 | 5<<3, BitsWritten: 13}
	_cf = make(map[int]code)
	_cf[64] = code{Code: 27 << (3 + 8), BitsWritten: 5}
	_cf[128] = code{Code: 18 << (3 + 8), BitsWritten: 5}
	_cf[192] = code{Code: 23 << (2 + 8), BitsWritten: 6}
	_cf[256] = code{Code: 55 << (1 + 8), BitsWritten: 7}
	_cf[320] = code{Code: 54 << 8, BitsWritten: 8}
	_cf[384] = code{Code: 55 << 8, BitsWritten: 8}
	_cf[448] = code{Code: 100 << 8, BitsWritten: 8}
	_cf[512] = code{Code: 101 << 8, BitsWritten: 8}
	_cf[576] = code{Code: 104 << 8, BitsWritten: 8}
	_cf[640] = code{Code: 103 << 8, BitsWritten: 8}
	_cf[704] = code{Code: 102 << 8, BitsWritten: 9}
	_cf[768] = code{Code: 102<<8 | 1<<7, BitsWritten: 9}
	_cf[832] = code{Code: 105 << 8, BitsWritten: 9}
	_cf[896] = code{Code: 105<<8 | 1<<7, BitsWritten: 9}
	_cf[960] = code{Code: 106 << 8, BitsWritten: 9}
	_cf[1024] = code{Code: 106<<8 | 1<<7, BitsWritten: 9}
	_cf[1088] = code{Code: 107 << 8, BitsWritten: 9}
	_cf[1152] = code{Code: 107<<8 | 1<<7, BitsWritten: 9}
	_cf[1216] = code{Code: 108 << 8, BitsWritten: 9}
	_cf[1280] = code{Code: 108<<8 | 1<<7, BitsWritten: 9}
	_cf[1344] = code{Code: 109 << 8, BitsWritten: 9}
	_cf[1408] = code{Code: 109<<8 | 1<<7, BitsWritten: 9}
	_cf[1472] = code{Code: 76 << 8, BitsWritten: 9}
	_cf[1536] = code{Code: 76<<8 | 1<<7, BitsWritten: 9}
	_cf[1600] = code{Code: 77 << 8, BitsWritten: 9}
	_cf[1664] = code{Code: 24 << (2 + 8), BitsWritten: 6}
	_cf[1728] = code{Code: 77<<8 | 1<<7, BitsWritten: 9}
	_bcb = make(map[int]code)
	_bcb[1792] = code{Code: 1 << 8, BitsWritten: 11}
	_bcb[1856] = code{Code: 1<<8 | 4<<5, BitsWritten: 11}
	_bcb[1920] = code{Code: 1<<8 | 5<<5, BitsWritten: 11}
	_bcb[1984] = code{Code: 1<<8 | 2<<4, BitsWritten: 12}
	_bcb[2048] = code{Code: 1<<8 | 3<<4, BitsWritten: 12}
	_bcb[2112] = code{Code: 1<<8 | 4<<4, BitsWritten: 12}
	_bcb[2176] = code{Code: 1<<8 | 5<<4, BitsWritten: 12}
	_bcb[2240] = code{Code: 1<<8 | 6<<4, BitsWritten: 12}
	_bcb[2304] = code{Code: 1<<8 | 7<<4, BitsWritten: 12}
	_bcb[2368] = code{Code: 1<<8 | 12<<4, BitsWritten: 12}
	_bcb[2432] = code{Code: 1<<8 | 13<<4, BitsWritten: 12}
	_bcb[2496] = code{Code: 1<<8 | 14<<4, BitsWritten: 12}
	_bcb[2560] = code{Code: 1<<8 | 15<<4, BitsWritten: 12}
	_gb = make(map[int]byte)
	_gb[0] = 0xFF
	_gb[1] = 0xFE
	_gb[2] = 0xFC
	_gb[3] = 0xF8
	_gb[4] = 0xF0
	_gb[5] = 0xE0
	_gb[6] = 0xC0
	_gb[7] = 0x80
	_gb[8] = 0x00
}

func (_begd *Decoder) looseFetchEOL() (bool, error) {
	_faae, _cfea := _begd._be.ReadBits(12)
	if _cfea != nil {
		return false, _cfea
	}
	switch _faae {
	case 0x1:
		return true, nil
	case 0x0:
		for {
			_dae, _cab := _begd._be.ReadBool()
			if _cab != nil {
				return false, _cab
			}
			if _dae {
				return true, nil
			}
		}
	default:
		return false, nil
	}
}

const (
	_ tiffType = iota
	_ccf
	_ec
	_gcf
)

type code struct {
	Code        uint16
	BitsWritten int
}

var _egd = [...][]uint16{{3, 2}, {1, 4}, {6, 5}, {7}, {9, 8}, {10, 11, 12}, {13, 14}, {15}, {16, 17, 0, 18, 64}, {24, 25, 23, 22, 19, 20, 21, 1792, 1856, 1920}, {1984, 2048, 2112, 2176, 2240, 2304, 2368, 2432, 2496, 2560, 52, 55, 56, 59, 60, 320, 384, 448, 53, 54, 50, 51, 44, 45, 46, 47, 57, 58, 61, 256, 48, 49, 62, 63, 30, 31, 32, 33, 40, 41, 128, 192, 26, 27, 28, 29, 34, 35, 36, 37, 38, 39, 42, 43}, {640, 704, 768, 832, 1280, 1344, 1408, 1472, 1536, 1600, 1664, 1728, 512, 576, 896, 960, 1024, 1088, 1152, 1216}}

func (_ecba *Decoder) decode2D() error {
	_ecba._gdc = _ecba._bae
	_ecba._gbd, _ecba._feb = _ecba._feb, _ecba._gbd
	_dffad := true
	var (
		_cfag bool
		_cgca int
		_bef  error
	)
	_ecba._bae = 0
_cca:
	for _cgca < _ecba._dcd {
		_cbb := _dc._adcga
		for {
			_cfag, _bef = _ecba._be.ReadBool()
			if _bef != nil {
				return _bef
			}
			_cbb = _cbb.walk(_cfag)
			if _cbb == nil {
				continue _cca
			}
			if !_cbb._cdgg {
				continue
			}
			switch _cbb._efe {
			case _fdd:
				var _eef int
				if _dffad {
					_eef, _bef = _ecba.decodeRun(_bc)
				} else {
					_eef, _bef = _ecba.decodeRun(_fd)
				}
				if _bef != nil {
					return _bef
				}
				_cgca += _eef
				_ecba._gbd[_ecba._bae] = _cgca
				_ecba._bae++
				if _dffad {
					_eef, _bef = _ecba.decodeRun(_fd)
				} else {
					_eef, _bef = _ecba.decodeRun(_bc)
				}
				if _bef != nil {
					return _bef
				}
				_cgca += _eef
				_ecba._gbd[_ecba._bae] = _cgca
				_ecba._bae++
			case _dce:
				_ebf := _ecba.getNextChangingElement(_cgca, _dffad) + 1
				if _ebf >= _ecba._gdc {
					_cgca = _ecba._dcd
				} else {
					_cgca = _ecba._feb[_ebf]
				}
			default:
				_bfgg := _ecba.getNextChangingElement(_cgca, _dffad)
				if _bfgg >= _ecba._gdc || _bfgg == -1 {
					_cgca = _ecba._dcd + _cbb._efe
				} else {
					_cgca = _ecba._feb[_bfgg] + _cbb._efe
				}
				_ecba._gbd[_ecba._bae] = _cgca
				_ecba._bae++
				_dffad = !_dffad
			}
			continue _cca
		}
	}
	return nil
}

func (_gfd *Decoder) decodeG32D() error {
	_gfd._gdc = _gfd._bae
	_gfd._gbd, _gfd._feb = _gfd._feb, _gfd._gbd
	_ecb := true
	var (
		_ceb bool
		_bec int
		_cfe error
	)
	_gfd._bae = 0
_ggbg:
	for _bec < _gfd._dcd {
		_eea := _dc._adcga
		for {
			_ceb, _cfe = _gfd._be.ReadBool()
			if _cfe != nil {
				return _cfe
			}
			_eea = _eea.walk(_ceb)
			if _eea == nil {
				continue _ggbg
			}
			if !_eea._cdgg {
				continue
			}
			switch _eea._efe {
			case _fdd:
				var _fbd int
				if _ecb {
					_fbd, _cfe = _gfd.decodeRun(_bc)
				} else {
					_fbd, _cfe = _gfd.decodeRun(_fd)
				}
				if _cfe != nil {
					return _cfe
				}
				_bec += _fbd
				_gfd._gbd[_gfd._bae] = _bec
				_gfd._bae++
				if _ecb {
					_fbd, _cfe = _gfd.decodeRun(_fd)
				} else {
					_fbd, _cfe = _gfd.decodeRun(_bc)
				}
				if _cfe != nil {
					return _cfe
				}
				_bec += _fbd
				_gfd._gbd[_gfd._bae] = _bec
				_gfd._bae++
			case _dce:
				_bafg := _gfd.getNextChangingElement(_bec, _ecb) + 1
				if _bafg >= _gfd._gdc {
					_bec = _gfd._dcd
				} else {
					_bec = _gfd._feb[_bafg]
				}
			default:
				_gec := _gfd.getNextChangingElement(_bec, _ecb)
				if _gec >= _gfd._gdc || _gec == -1 {
					_bec = _gfd._dcd + _eea._efe
				} else {
					_bec = _gfd._feb[_gec] + _eea._efe
				}
				_gfd._gbd[_gfd._bae] = _bec
				_gfd._bae++
				_ecb = !_ecb
			}
			continue _ggbg
		}
	}
	return nil
}

func _cgf(_bddc [][]byte) [][]byte {
	_aged := make([]byte, len(_bddc[0]))
	for _dge := range _aged {
		_aged[_dge] = _gbaf
	}
	_bddc = append(_bddc, []byte{})
	for _eae := len(_bddc) - 1; _eae > 0; _eae-- {
		_bddc[_eae] = _bddc[_eae-1]
	}
	_bddc[0] = _aged
	return _bddc
}

var _aa = [...][]uint16{{2, 3, 4, 5, 6, 7}, {128, 8, 9, 64, 10, 11}, {192, 1664, 16, 17, 13, 14, 15, 1, 12}, {26, 21, 28, 27, 18, 24, 25, 22, 256, 23, 20, 19}, {33, 34, 35, 36, 37, 38, 31, 32, 29, 53, 54, 39, 40, 41, 42, 43, 44, 30, 61, 62, 63, 0, 320, 384, 45, 59, 60, 46, 49, 50, 51, 52, 55, 56, 57, 58, 448, 512, 640, 576, 47, 48}, {1472, 1536, 1600, 1728, 704, 768, 832, 896, 960, 1024, 1088, 1152, 1216, 1280, 1344, 1408}, {}, {1792, 1856, 1920}, {1984, 2048, 2112, 2176, 2240, 2304, 2368, 2432, 2496, 2560}}

type Decoder struct {
	_dcd  int
	_gba  int
	_gde  int
	_gbb  []byte
	_fbb  int
	_gbf  bool
	_da   bool
	_ga   bool
	_fcd  bool
	_gdg  bool
	_fe   bool
	_efgc bool
	_ecf  int
	_bb   int
	_feb  []int
	_gbd  []int
	_gdc  int
	_bae  int
	_dd   int
	_db   int
	_be   *_fa.Reader
	_fcdd tiffType
	_gfg  error
}

var (
	_eb  = _c.New("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0063\u006f\u0072\u0072\u0075\u0070\u0074\u0065\u0064\u0020\u0052T\u0043")
	_bde = _c.New("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0045\u004f\u004c\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
)

func (_bfg *Decoder) decoderRowType41D() error {
	if _bfg._efgc {
		_bfg._be.Align()
	}
	_bfg._be.Mark()
	var (
		_ega bool
		_dbb error
	)
	if _bfg._gdg {
		_ega, _dbb = _bfg.tryFetchEOL()
		if _dbb != nil {
			return _dbb
		}
		if !_ega {
			return _bde
		}
	} else {
		_ega, _dbb = _bfg.looseFetchEOL()
		if _dbb != nil {
			return _dbb
		}
	}
	if !_ega {
		_bfg._be.Reset()
	}
	if _ega && _bfg._fe {
		_bfg._be.Mark()
		for _aee := 0; _aee < 5; _aee++ {
			_ega, _dbb = _bfg.tryFetchEOL()
			if _dbb != nil {
				if _c.Is(_dbb, _b.EOF) {
					if _aee == 0 {
						break
					}
					return _eb
				}
			}
			if _ega {
				continue
			}
			if _aee > 0 {
				return _eb
			}
			break
		}
		if _ega {
			return _b.EOF
		}
		_bfg._be.Reset()
	}
	if _dbb = _bfg.decode1D(); _dbb != nil {
		return _dbb
	}
	return nil
}

func (_fcb *treeNode) set(_cdbc bool, _edg *treeNode) {
	if !_cdbc {
		_fcb._dcecg = _edg
	} else {
		_fcb._ccdfb = _edg
	}
}

var (
	_gc   map[int]code
	_gegg map[int]code
	_gdf  map[int]code
	_cf   map[int]code
	_bcb  map[int]code
	_gb   map[int]byte
	_ffg  = code{Code: 1 << 4, BitsWritten: 12}
	_ccd  = code{Code: 3 << 3, BitsWritten: 13}
	_ac   = code{Code: 2 << 3, BitsWritten: 13}
	_fg   = code{Code: 1 << 12, BitsWritten: 4}
	_fag  = code{Code: 1 << 13, BitsWritten: 3}
	_efg  = code{Code: 1 << 15, BitsWritten: 1}
	_ca   = code{Code: 3 << 13, BitsWritten: 3}
	_bcd  = code{Code: 3 << 10, BitsWritten: 6}
	_bf   = code{Code: 3 << 9, BitsWritten: 7}
	_afb  = code{Code: 2 << 13, BitsWritten: 3}
	_ccg  = code{Code: 2 << 10, BitsWritten: 6}
	_egf  = code{Code: 2 << 9, BitsWritten: 7}
)

func _bdfg(_ffad, _def int) code {
	var _bab code
	switch _def - _ffad {
	case -1:
		_bab = _ca
	case -2:
		_bab = _bcd
	case -3:
		_bab = _bf
	case 0:
		_bab = _efg
	case 1:
		_bab = _afb
	case 2:
		_bab = _ccg
	case 3:
		_bab = _egf
	}
	return _bab
}

func (_agd *Encoder) encodeG4(_dcgd [][]byte) []byte {
	_aaa := make([][]byte, len(_dcgd))
	copy(_aaa, _dcgd)
	_aaa = _cgf(_aaa)
	var _befe []byte
	var _gce int
	for _bccc := 1; _bccc < len(_aaa); _bccc++ {
		if _agd.Rows > 0 && !_agd.EndOfBlock && _bccc == (_agd.Rows+1) {
			break
		}
		var _cabf []byte
		var _egdb, _fedf, _dfb int
		_adge := _gce
		_dabc := -1
		for _dabc < len(_aaa[_bccc]) {
			_egdb = _fbc(_aaa[_bccc], _dabc)
			_fedf = _beb(_aaa[_bccc], _aaa[_bccc-1], _dabc)
			_dfb = _fbc(_aaa[_bccc-1], _fedf)
			if _dfb < _egdb {
				_cabf, _adge = _bdd(_cabf, _adge, _fg)
				_dabc = _dfb
			} else {
				if _g.Abs(float64(_fedf-_egdb)) > 3 {
					_cabf, _adge, _dabc = _ebd(_aaa[_bccc], _cabf, _adge, _dabc, _egdb)
				} else {
					_cabf, _adge = _abfb(_cabf, _adge, _egdb, _fedf)
					_dabc = _egdb
				}
			}
		}
		_befe = _agd.appendEncodedRow(_befe, _cabf, _gce)
		if _agd.EncodedByteAlign {
			_adge = 0
		}
		_gce = _adge % 8
	}
	if _agd.EndOfBlock {
		_deb, _ := _fff(_gce)
		_befe = _agd.appendEncodedRow(_befe, _deb, _gce)
	}
	return _befe
}

type treeNode struct {
	_dcecg *treeNode
	_ccdfb *treeNode
	_efe   int
	_ccfg  bool
	_cdgg  bool
}

var (
	_gbaf  byte = 1
	_geadg byte = 0
)

func _aeg(_ffa int, _abg bool) (code, int, bool) {
	if _ffa < 64 {
		if _abg {
			return _gegg[_ffa], 0, true
		}
		return _gc[_ffa], 0, true
	}
	_dfd := _ffa / 64
	if _dfd > 40 {
		return _bcb[2560], _ffa - 2560, false
	}
	if _dfd > 27 {
		return _bcb[_dfd*64], _ffa - _dfd*64, false
	}
	if _abg {
		return _cf[_dfd*64], _ffa - _dfd*64, false
	}
	return _gdf[_dfd*64], _ffa - _dfd*64, false
}

func _bdd(_cgdf []byte, _gbe int, _cfac code) ([]byte, int) {
	_bad := 0
	for _bad < _cfac.BitsWritten {
		_dda := _gbe / 8
		_fgae := _gbe % 8
		if _dda >= len(_cgdf) {
			_cgdf = append(_cgdf, 0)
		}
		_gag := 8 - _fgae
		_cff := _cfac.BitsWritten - _bad
		if _gag > _cff {
			_gag = _cff
		}
		if _bad < 8 {
			_cgdf[_dda] = _cgdf[_dda] | byte(_cfac.Code>>uint(8+_fgae-_bad))&_gb[8-_gag-_fgae]
		} else {
			_cgdf[_dda] = _cgdf[_dda] | (byte(_cfac.Code<<uint(_bad-8))&_gb[8-_gag])>>uint(_fgae)
		}
		_gbe += _gag
		_bad += _gag
	}
	return _cgdf, _gbe
}

func (_gaf *Encoder) Encode(pixels [][]byte) []byte {
	if _gaf.BlackIs1 {
		_gbaf = 0
		_geadg = 1
	} else {
		_gbaf = 1
		_geadg = 0
	}
	if _gaf.K == 0 {
		return _gaf.encodeG31D(pixels)
	}
	if _gaf.K > 0 {
		return _gaf.encodeG32D(pixels)
	}
	if _gaf.K < 0 {
		return _gaf.encodeG4(pixels)
	}
	return nil
}

func _cebf(_ccdb int) ([]byte, int) {
	var _bfe []byte
	for _bcde := 0; _bcde < 6; _bcde++ {
		_bfe, _ccdb = _bdd(_bfe, _ccdb, _ffg)
	}
	return _bfe, _ccdb % 8
}

func (_cga *Encoder) encodeG32D(_fcddd [][]byte) []byte {
	var _ada []byte
	var _fga int
	for _agbg := 0; _agbg < len(_fcddd); _agbg += _cga.K {
		if _cga.Rows > 0 && !_cga.EndOfBlock && _agbg == _cga.Rows {
			break
		}
		_fcf, _dbbf := _ddg(_fcddd[_agbg], _fga, _ccd)
		_ada = _cga.appendEncodedRow(_ada, _fcf, _fga)
		if _cga.EncodedByteAlign {
			_dbbf = 0
		}
		_fga = _dbbf
		for _ccdd := _agbg + 1; _ccdd < (_agbg+_cga.K) && _ccdd < len(_fcddd); _ccdd++ {
			if _cga.Rows > 0 && !_cga.EndOfBlock && _ccdd == _cga.Rows {
				break
			}
			_ecfd, _ggc := _bdd(nil, _fga, _ac)
			var _adg, _dbg, _ea int
			_daee := -1
			for _daee < len(_fcddd[_ccdd]) {
				_adg = _fbc(_fcddd[_ccdd], _daee)
				_dbg = _beb(_fcddd[_ccdd], _fcddd[_ccdd-1], _daee)
				_ea = _fbc(_fcddd[_ccdd-1], _dbg)
				if _ea < _adg {
					_ecfd, _ggc = _bafc(_ecfd, _ggc)
					_daee = _ea
				} else {
					if _g.Abs(float64(_dbg-_adg)) > 3 {
						_ecfd, _ggc, _daee = _ebd(_fcddd[_ccdd], _ecfd, _ggc, _daee, _adg)
					} else {
						_ecfd, _ggc = _abfb(_ecfd, _ggc, _adg, _dbg)
						_daee = _adg
					}
				}
			}
			_ada = _cga.appendEncodedRow(_ada, _ecfd, _fga)
			if _cga.EncodedByteAlign {
				_ggc = 0
			}
			_fga = _ggc % 8
		}
	}
	if _cga.EndOfBlock {
		_ecg, _ := _gfe(_fga)
		_ada = _cga.appendEncodedRow(_ada, _ecg, _fga)
	}
	return _ada
}

func (_eeb *Encoder) encodeG31D(_caa [][]byte) []byte {
	var _fdb []byte
	_dbd := 0
	for _gef := range _caa {
		if _eeb.Rows > 0 && !_eeb.EndOfBlock && _gef == _eeb.Rows {
			break
		}
		_gee, _gdd := _ddg(_caa[_gef], _dbd, _ffg)
		_fdb = _eeb.appendEncodedRow(_fdb, _gee, _dbd)
		if _eeb.EncodedByteAlign {
			_gdd = 0
		}
		_dbd = _gdd
	}
	if _eeb.EndOfBlock {
		_dfae, _ := _cebf(_dbd)
		_fdb = _eeb.appendEncodedRow(_fdb, _dfae, _dbd)
	}
	return _fdb
}

func (_edb *Decoder) decodeRun(_bdf *tree) (int, error) {
	var _edbc int
	_fgb := _bdf._adcga
	for {
		_gead, _fdf := _edb._be.ReadBool()
		if _fdf != nil {
			return 0, _fdf
		}
		_fgb = _fgb.walk(_gead)
		if _fgb == nil {
			return 0, _c.New("\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006f\u0064\u0065\u0020\u0069n\u0020H\u0075f\u0066m\u0061\u006e\u0020\u0052\u004c\u0045\u0020\u0073\u0074\u0072\u0065\u0061\u006d")
		}
		if _fgb._cdgg {
			_edbc += _fgb._efe
			switch {
			case _fgb._efe >= 64:
				_fgb = _bdf._adcga
			case _fgb._efe >= 0:
				return _edbc, nil
			default:
				return _edb._dcd, nil
			}
		}
	}
}

func _fff(_ceg int) ([]byte, int) {
	var _dca []byte
	for _eff := 0; _eff < 2; _eff++ {
		_dca, _ceg = _bdd(_dca, _ceg, _ffg)
	}
	return _dca, _ceg % 8
}

func (_fgbc *tree) fill(_cgdb, _febd, _acf int) error {
	_abb := _fgbc._adcga
	for _eaa := 0; _eaa < _cgdb; _eaa++ {
		_aac := _cgdb - 1 - _eaa
		_bbc := ((_febd >> uint(_aac)) & 1) != 0
		_cgb := _abb.walk(_bbc)
		if _cgb != nil {
			if _cgb._cdgg {
				return _c.New("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067")
			}
			_abb = _cgb
			continue
		}
		_cgb = &treeNode{}
		if _eaa == _cgdb-1 {
			_cgb._efe = _acf
			_cgb._cdgg = true
		}
		if _febd == 0 {
			_cgb._ccfg = true
		}
		_abb.set(_bbc, _cgb)
		_abb = _cgb
	}
	return nil
}

func (_bdc tiffType) String() string {
	switch _bdc {
	case _ccf:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u004d\u006f\u0064i\u0066\u0069\u0065\u0064\u0048\u0075\u0066\u0066\u006d\u0061n\u0052\u006c\u0065"
	case _ec:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0034"
	case _gcf:
		return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0036"
	default:
		return "\u0075n\u0064\u0065\u0066\u0069\u006e\u0065d"
	}
}

func (_cfa *Decoder) decodeRowType2() error {
	if _cfa._efgc {
		_cfa._be.Align()
	}
	if _gea := _cfa.decode1D(); _gea != nil {
		return _gea
	}
	return nil
}

func _beb(_fabf, _gcca []byte, _gbea int) int {
	_acaa := _fbc(_gcca, _gbea)
	if _acaa < len(_gcca) && (_gbea == -1 && _gcca[_acaa] == _gbaf || _gbea >= 0 && _gbea < len(_fabf) && _fabf[_gbea] == _gcca[_acaa] || _gbea >= len(_fabf) && _fabf[_gbea-1] != _gcca[_acaa]) {
		_acaa = _fbc(_gcca, _acaa)
	}
	return _acaa
}

func (_fbg *Decoder) decodeRowType6() error {
	if _fbg._efgc {
		_fbg._be.Align()
	}
	if _fbg._fe {
		_fbg._be.Mark()
		_ffe, _dfa := _fbg.tryFetchEOL()
		if _dfa != nil {
			return _dfa
		}
		if _ffe {
			_ffe, _dfa = _fbg.tryFetchEOL()
			if _dfa != nil {
				return _dfa
			}
			if _ffe {
				return _b.EOF
			}
		}
		_fbg._be.Reset()
	}
	return _fbg.decode2D()
}

type tiffType int

func (_bca *Decoder) tryFetchEOL1() (bool, error) {
	_agfg, _dcff := _bca._be.ReadBits(13)
	if _dcff != nil {
		return false, _dcff
	}
	return _agfg == 0x3, nil
}

func _aec(_debc []byte, _gga int, _gdfb int, _dg bool) ([]byte, int) {
	var (
		_ccff code
		_aaad bool
	)
	for !_aaad {
		_ccff, _gdfb, _aaad = _aeg(_gdfb, _dg)
		_debc, _gga = _bdd(_debc, _gga, _ccff)
	}
	return _debc, _gga
}

var _ggb = [...][]uint16{{0x7, 0x8, 0xb, 0xc, 0xe, 0xf}, {0x12, 0x13, 0x14, 0x1b, 0x7, 0x8}, {0x17, 0x18, 0x2a, 0x2b, 0x3, 0x34, 0x35, 0x7, 0x8}, {0x13, 0x17, 0x18, 0x24, 0x27, 0x28, 0x2b, 0x3, 0x37, 0x4, 0x8, 0xc}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1a, 0x1b, 0x2, 0x24, 0x25, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x3, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x4, 0x4a, 0x4b, 0x5, 0x52, 0x53, 0x54, 0x55, 0x58, 0x59, 0x5a, 0x5b, 0x64, 0x65, 0x67, 0x68, 0xa, 0xb}, {0x98, 0x99, 0x9a, 0x9b, 0xcc, 0xcd, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8, 0xd9, 0xda, 0xdb}, {}, {0x8, 0xc, 0xd}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1c, 0x1d, 0x1e, 0x1f}}

func NewDecoder(data []byte, options DecodeOptions) (*Decoder, error) {
	_cd := &Decoder{_be: _fa.NewReader(data), _dcd: options.Columns, _gba: options.Rows, _fbb: options.DamagedRowsBeforeError, _gbb: make([]byte, (options.Columns+7)/8), _feb: make([]int, options.Columns+2), _gbd: make([]int, options.Columns+2), _efgc: options.EncodedByteAligned, _fcd: options.BlackIsOne, _gdg: options.EndOfLine, _fe: options.EndOfBlock}
	switch {
	case options.K == 0:
		_cd._fcdd = _ec
		if len(data) < 20 {
			return nil, _c.New("\u0074o\u006f\u0020\u0073\u0068o\u0072\u0074\u0020\u0063\u0063i\u0074t\u0066a\u0078\u0020\u0073\u0074\u0072\u0065\u0061m")
		}
		_dcec := data[:20]
		if _dcec[0] != 0 || (_dcec[1]>>4 != 1 && _dcec[1] != 1) {
			_cd._fcdd = _ccf
			_cgd := (uint16(_dcec[0])<<8 + uint16(_dcec[1]&0xff)) >> 4
			for _afe := 12; _afe < 160; _afe++ {
				_cgd = (_cgd << 1) + uint16((_dcec[_afe/8]>>uint16(7-(_afe%8)))&0x01)
				if _cgd&0xfff == 1 {
					_cd._fcdd = _ec
					break
				}
			}
		}
	case options.K < 0:
		_cd._fcdd = _gcf
	case options.K > 0:
		_cd._fcdd = _ec
		_cd._gbf = true
	}
	switch _cd._fcdd {
	case _ccf, _ec, _gcf:
	default:
		return nil, _c.New("\u0075\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u0063\u0069\u0074\u0074\u0066\u0061\u0078\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0072\u0020ty\u0070\u0065")
	}
	return _cd, nil
}

type tree struct{ _adcga *treeNode }

func (_bga *Decoder) decodeRowType4() error {
	if !_bga._gbf {
		return _bga.decoderRowType41D()
	}
	if _bga._efgc {
		_bga._be.Align()
	}
	_bga._be.Mark()
	_ad, _ced := _bga.tryFetchEOL()
	if _ced != nil {
		return _ced
	}
	if !_ad && _bga._gdg {
		_bga._dd++
		if _bga._dd > _bga._fbb {
			return _bde
		}
		_bga._be.Reset()
	}
	if !_ad {
		_bga._be.Reset()
	}
	_ee, _ced := _bga._be.ReadBool()
	if _ced != nil {
		return _ced
	}
	if _ee {
		if _ad && _bga._fe {
			if _ced = _bga.tryFetchRTC2D(); _ced != nil {
				return _ced
			}
		}
		_ced = _bga.decode1D()
	} else {
		_ced = _bga.decode2D()
	}
	if _ced != nil {
		return _ced
	}
	return nil
}

func (_agb *Decoder) decode1D() error {
	var (
		_cdc int
		_edf error
	)
	_dcge := true
	_agb._bae = 0
	for {
		var _fed int
		if _dcge {
			_fed, _edf = _agb.decodeRun(_bc)
		} else {
			_fed, _edf = _agb.decodeRun(_fd)
		}
		if _edf != nil {
			return _edf
		}
		_cdc += _fed
		_agb._gbd[_agb._bae] = _cdc
		_agb._bae++
		_dcge = !_dcge
		if _cdc >= _agb._dcd {
			break
		}
	}
	return nil
}

func (_bda *Decoder) Read(in []byte) (int, error) {
	if _bda._gfg != nil {
		return 0, _bda._gfg
	}
	_fac := len(in)
	var (
		_baf int
		_ggg int
	)
	for _fac != 0 {
		if _bda._bb >= _bda._ecf {
			if _faa := _bda.fetch(); _faa != nil {
				_bda._gfg = _faa
				return 0, _faa
			}
		}
		if _bda._ecf == -1 {
			return _baf, _b.EOF
		}
		switch {
		case _fac <= _bda._ecf-_bda._bb:
			_bg := _bda._gbb[_bda._bb : _bda._bb+_fac]
			for _, _de := range _bg {
				if !_bda._fcd {
					_de = ^_de
				}
				in[_ggg] = _de
				_ggg++
			}
			_baf += len(_bg)
			_bda._bb += len(_bg)
			return _baf, nil
		default:
			_fce := _bda._gbb[_bda._bb:]
			for _, _cdf := range _fce {
				if !_bda._fcd {
					_cdf = ^_cdf
				}
				in[_ggg] = _cdf
				_ggg++
			}
			_baf += len(_fce)
			_bda._bb += len(_fce)
			_fac -= len(_fce)
		}
	}
	return _baf, nil
}

func (_ege *Decoder) tryFetchEOL() (bool, error) {
	_gfgc, _dcgg := _ege._be.ReadBits(12)
	if _dcgg != nil {
		return false, _dcgg
	}
	return _gfgc == 0x1, nil
}

func (_dee *treeNode) walk(_bea bool) *treeNode {
	if _bea {
		return _dee._ccdfb
	}
	return _dee._dcecg
}

func _bega(_edc []byte, _acag bool, _adab int) (int, int) {
	_cbba := 0
	for _adab < len(_edc) {
		if _acag {
			if _edc[_adab] != _gbaf {
				break
			}
		} else {
			if _edc[_adab] != _geadg {
				break
			}
		}
		_cbba++
		_adab++
	}
	return _cbba, _adab
}

var (
	_d   *treeNode
	_a   *treeNode
	_fd  *tree
	_bc  *tree
	_e   *tree
	_dc  *tree
	_ge  = -2000
	_df  = -1000
	_dce = -3000
	_fdd = -4000
)

func _bafc(_bcae []byte, _ead int) ([]byte, int) { return _bdd(_bcae, _ead, _fg) }
func (_cebb *Decoder) tryFetchRTC2D() (_cdg error) {
	_cebb._be.Mark()
	var _gbc bool
	for _gcc := 0; _gcc < 5; _gcc++ {
		_gbc, _cdg = _cebb.tryFetchEOL1()
		if _cdg != nil {
			if _c.Is(_cdg, _b.EOF) {
				if _gcc == 0 {
					break
				}
				return _eb
			}
		}
		if _gbc {
			continue
		}
		if _gcc > 0 {
			return _eb
		}
		break
	}
	if _gbc {
		return _b.EOF
	}
	_cebb._be.Reset()
	return _cdg
}

func _gfe(_cbbg int) ([]byte, int) {
	var _eaf []byte
	for _cbc := 0; _cbc < 6; _cbc++ {
		_eaf, _cbbg = _bdd(_eaf, _cbbg, _ccd)
	}
	return _eaf, _cbbg % 8
}

func (_bcdg *tree) fillWithNode(_dcbc, _bfc int, _fbe *treeNode) error {
	_aga := _bcdg._adcga
	for _ffd := 0; _ffd < _dcbc; _ffd++ {
		_bgd := uint(_dcbc - 1 - _ffd)
		_gcfg := ((_bfc >> _bgd) & 1) != 0
		_debcb := _aga.walk(_gcfg)
		if _debcb != nil {
			if _debcb._cdgg {
				return _c.New("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067")
			}
			_aga = _debcb
			continue
		}
		if _ffd == _dcbc-1 {
			_debcb = _fbe
		} else {
			_debcb = &treeNode{}
		}
		if _bfc == 0 {
			_debcb._ccfg = true
		}
		_aga.set(_gcfg, _debcb)
		_aga = _debcb
	}
	return nil
}

func _ddg(_ecgf []byte, _becd int, _gfea code) ([]byte, int) {
	_adc := true
	var _gad []byte
	_gad, _becd = _bdd(nil, _becd, _gfea)
	_cce := 0
	var _adga int
	for _cce < len(_ecgf) {
		_adga, _cce = _bega(_ecgf, _adc, _cce)
		_gad, _becd = _aec(_gad, _becd, _adga, _adc)
		_adc = !_adc
	}
	return _gad, _becd % 8
}

func _ebd(_bff, _ccdc []byte, _bba, _gadc, _debe int) ([]byte, int, int) {
	_gdgc := _fbc(_bff, _debe)
	_age := _gadc >= 0 && _bff[_gadc] == _gbaf || _gadc == -1
	_ccdc, _bba = _bdd(_ccdc, _bba, _fag)
	var _bfb int
	if _gadc > -1 {
		_bfb = _debe - _gadc
	} else {
		_bfb = _debe - _gadc - 1
	}
	_ccdc, _bba = _aec(_ccdc, _bba, _bfb, _age)
	_age = !_age
	_gfb := _gdgc - _debe
	_ccdc, _bba = _aec(_ccdc, _bba, _gfb, _age)
	_gadc = _gdgc
	return _ccdc, _bba, _gadc
}

var _ff = [...][]uint16{{0x2, 0x3}, {0x2, 0x3}, {0x2, 0x3}, {0x3}, {0x4, 0x5}, {0x4, 0x5, 0x7}, {0x4, 0x7}, {0x18}, {0x17, 0x18, 0x37, 0x8, 0xf}, {0x17, 0x18, 0x28, 0x37, 0x67, 0x68, 0x6c, 0x8, 0xc, 0xd}, {0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x1c, 0x1d, 0x1e, 0x1f, 0x24, 0x27, 0x28, 0x2b, 0x2c, 0x33, 0x34, 0x35, 0x37, 0x38, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xda, 0xdb}, {0x4a, 0x4b, 0x4c, 0x4d, 0x52, 0x53, 0x54, 0x55, 0x5a, 0x5b, 0x64, 0x65, 0x6c, 0x6d, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77}}

type DecodeOptions struct {
	Columns                int
	Rows                   int
	K                      int
	EncodedByteAligned     bool
	BlackIsOne             bool
	EndOfBlock             bool
	EndOfLine              bool
	DamagedRowsBeforeError int
}
type Encoder struct {
	K                      int
	EndOfLine              bool
	EncodedByteAlign       bool
	Columns                int
	Rows                   int
	EndOfBlock             bool
	BlackIs1               bool
	DamagedRowsBeforeError int
}
