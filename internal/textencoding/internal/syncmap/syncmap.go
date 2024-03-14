package syncmap

import _a "sync"

func (_dc *RuneSet) Exists(r rune) bool {
	_dc._ee.RLock()
	defer _dc._ee.RUnlock()
	_, _ba := _dc._gdc[r]
	return _ba
}

func (_fbe *RuneStringMap) Write(r rune, s string) {
	_fbe._eg.Lock()
	defer _fbe._eg.Unlock()
	_fbe._gdcc[r] = s
}

func MakeRuneByteMap(length int) *RuneByteMap {
	_bfd := make(map[rune]byte, length)
	return &RuneByteMap{_ec: _bfd}
}

func (_ecc *StringsMap) Copy() *StringsMap {
	_ecc._gf.RLock()
	defer _ecc._gf.RUnlock()
	_baca := map[string]string{}
	for _cfa, _cfb := range _ecc._bda {
		_baca[_cfa] = _cfb
	}
	return &StringsMap{_bda: _baca}
}
func (_gg *RuneSet) Length() int                  { _gg._ee.RLock(); defer _gg._ee.RUnlock(); return len(_gg._gdc) }
func NewByteRuneMap(m map[byte]rune) *ByteRuneMap { return &ByteRuneMap{_af: m} }
func (_fg *ByteRuneMap) Range(f func(_ab byte, _fe rune) (_bf bool)) {
	_fg._g.RLock()
	defer _fg._g.RUnlock()
	for _cb, _ed := range _fg._af {
		if f(_cb, _ed) {
			break
		}
	}
}

type ByteRuneMap struct {
	_af map[byte]rune
	_g  _a.RWMutex
}

func (_cad *StringsMap) Read(g string) (string, bool) {
	_cad._gf.RLock()
	defer _cad._gf.RUnlock()
	_cebb, _bbd := _cad._bda[g]
	return _cebb, _bbd
}

func (_ea *RuneStringMap) Length() int {
	_ea._eg.RLock()
	defer _ea._eg.RUnlock()
	return len(_ea._gdcc)
}

func (_bcc *StringRuneMap) Write(g string, r rune) {
	_bcc._ad.Lock()
	defer _bcc._ad.Unlock()
	_bcc._abb[g] = r
}

type RuneSet struct {
	_gdc map[rune]struct{}
	_ee  _a.RWMutex
}

func (_dga *RuneUint16Map) Write(r rune, g uint16) {
	_dga._egc.Lock()
	defer _dga._egc.Unlock()
	_dga._def[r] = g
}

func (_bd *RuneByteMap) Range(f func(_ce rune, _d byte) (_abd bool)) {
	_bd._acd.RLock()
	defer _bd._acd.RUnlock()
	for _cf, _de := range _bd._ec {
		if f(_cf, _de) {
			break
		}
	}
}

func (_gda *RuneUint16Map) Length() int {
	_gda._egc.RLock()
	defer _gda._egc.RUnlock()
	return len(_gda._def)
}

func (_cbd *RuneSet) Write(r rune) {
	_cbd._ee.Lock()
	defer _cbd._ee.Unlock()
	_cbd._gdc[r] = struct{}{}
}

func (_ffb *RuneUint16Map) RangeDelete(f func(_bfb rune, _dd uint16) (_ece bool, _gdf bool)) {
	_ffb._egc.Lock()
	defer _ffb._egc.Unlock()
	for _agd, _bfe := range _ffb._def {
		_fgf, _gdd := f(_agd, _bfe)
		if _fgf {
			delete(_ffb._def, _agd)
		}
		if _gdd {
			break
		}
	}
}

func (_gcc *StringRuneMap) Read(g string) (rune, bool) {
	_gcc._ad.RLock()
	defer _gcc._ad.RUnlock()
	_ae, _gea := _gcc._abb[g]
	return _ae, _gea
}

func (_ca *RuneUint16Map) Delete(r rune) {
	_ca._egc.Lock()
	defer _ca._egc.Unlock()
	delete(_ca._def, r)
}

type StringRuneMap struct {
	_abb map[string]rune
	_ad  _a.RWMutex
}

func (_ede *RuneUint16Map) Range(f func(_gb rune, _ceb uint16) (_fd bool)) {
	_ede._egc.RLock()
	defer _ede._egc.RUnlock()
	for _bcd, _baf := range _ede._def {
		if f(_bcd, _baf) {
			break
		}
	}
}
func (_b *ByteRuneMap) Write(b byte, r rune) { _b._g.Lock(); defer _b._g.Unlock(); _b._af[b] = r }
func (_ag *RuneStringMap) Range(f func(_dee rune, _eea string) (_gga bool)) {
	_ag._eg.RLock()
	defer _ag._eg.RUnlock()
	for _bb, _bc := range _ag._gdcc {
		if f(_bb, _bc) {
			break
		}
	}
}

func NewStringsMap(tuples []StringsTuple) *StringsMap {
	_dff := map[string]string{}
	for _, _aca := range tuples {
		_dff[_aca.Key] = _aca.Value
	}
	return &StringsMap{_bda: _dff}
}

func (_fb *RuneByteMap) Length() int {
	_fb._acd.RLock()
	defer _fb._acd.RUnlock()
	return len(_fb._ec)
}

type RuneStringMap struct {
	_gdcc map[rune]string
	_eg   _a.RWMutex
}

func (_ge *RuneSet) Range(f func(_aff rune) (_ff bool)) {
	_ge._ee.RLock()
	defer _ge._ee.RUnlock()
	for _ga := range _ge._gdc {
		if f(_ga) {
			break
		}
	}
}

func (_gbd *StringRuneMap) Length() int {
	_gbd._ad.RLock()
	defer _gbd._ad.RUnlock()
	return len(_gbd._abb)
}

type RuneUint16Map struct {
	_def map[rune]uint16
	_egc _a.RWMutex
}

func NewRuneStringMap(m map[rune]string) *RuneStringMap { return &RuneStringMap{_gdcc: m} }
func (_ef *RuneStringMap) Read(r rune) (string, bool) {
	_ef._eg.RLock()
	defer _ef._eg.RUnlock()
	_gc, _ggg := _ef._gdcc[r]
	return _gc, _ggg
}

type StringsMap struct {
	_bda map[string]string
	_gf  _a.RWMutex
}

func (_aa *RuneByteMap) Read(r rune) (byte, bool) {
	_aa._acd.RLock()
	defer _aa._acd.RUnlock()
	_fea, _fa := _aa._ec[r]
	return _fea, _fa
}
func MakeByteRuneMap(length int) *ByteRuneMap { return &ByteRuneMap{_af: make(map[byte]rune, length)} }

type RuneByteMap struct {
	_ec  map[rune]byte
	_acd _a.RWMutex
}

func NewStringRuneMap(m map[string]rune) *StringRuneMap { return &StringRuneMap{_abb: m} }
func (_bac *StringsMap) Range(f func(_edf, _cde string) (_ged bool)) {
	_bac._gf.RLock()
	defer _bac._gf.RUnlock()
	for _dffc, _eab := range _bac._bda {
		if f(_dffc, _eab) {
			break
		}
	}
}

type StringsTuple struct{ Key, Value string }

func (_dg *RuneUint16Map) Read(r rune) (uint16, bool) {
	_dg._egc.RLock()
	defer _dg._egc.RUnlock()
	_cd, _ead := _dg._def[r]
	return _cd, _ead
}

func (_gd *RuneByteMap) Write(r rune, b byte) {
	_gd._acd.Lock()
	defer _gd._acd.Unlock()
	_gd._ec[r] = b
}
func MakeRuneSet(length int) *RuneSet { return &RuneSet{_gdc: make(map[rune]struct{}, length)} }
func (_bfg *ByteRuneMap) Length() int { _bfg._g.RLock(); defer _bfg._g.RUnlock(); return len(_bfg._af) }
func (_gggb *StringsMap) Write(g1, g2 string) {
	_gggb._gf.Lock()
	defer _gggb._gf.Unlock()
	_gggb._bda[g1] = g2
}

func (_f *ByteRuneMap) Read(b byte) (rune, bool) {
	_f._g.RLock()
	defer _f._g.RUnlock()
	_c, _ac := _f._af[b]
	return _c, _ac
}

func MakeRuneUint16Map(length int) *RuneUint16Map {
	return &RuneUint16Map{_def: make(map[rune]uint16, length)}
}

func (_add *StringRuneMap) Range(f func(_ggf string, _egd rune) (_efe bool)) {
	_add._ad.RLock()
	defer _add._ad.RUnlock()
	for _eaf, _df := range _add._abb {
		if f(_eaf, _df) {
			break
		}
	}
}
