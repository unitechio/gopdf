package syncmap

import _a "sync"

func (_ge *RuneStringMap) Read(r rune) (string, bool) {
	_ge._dg.RLock()
	defer _ge._dg.RUnlock()
	_edg, _ee := _ge._cda[r]
	return _edg, _ee
}
func NewByteRuneMap(m map[byte]rune) *ByteRuneMap { return &ByteRuneMap{_ae: m} }
func (_da *ByteRuneMap) Length() int              { _da._g.RLock(); defer _da._g.RUnlock(); return len(_da._ae) }
func MakeByteRuneMap(length int) *ByteRuneMap     { return &ByteRuneMap{_ae: make(map[byte]rune, length)} }

type RuneByteMap struct {
	_eac map[rune]byte
	_abd _a.RWMutex
}

func MakeRuneSet(length int) *RuneSet { return &RuneSet{_cg: make(map[rune]struct{}, length)} }
func (_acg *RuneStringMap) Write(r rune, s string) {
	_acg._dg.Lock()
	defer _acg._dg.Unlock()
	_acg._cda[r] = s
}
func (_f *ByteRuneMap) Read(b byte) (rune, bool) {
	_f._g.RLock()
	defer _f._g.RUnlock()
	_b, _c := _f._ae[b]
	return _b, _c
}
func (_gf *RuneSet) Length() int                        { _gf._bgc.RLock(); defer _gf._bgc.RUnlock(); return len(_gf._cg) }
func NewStringRuneMap(m map[string]rune) *StringRuneMap { return &StringRuneMap{_fgc: m} }

type RuneStringMap struct {
	_cda map[rune]string
	_dg  _a.RWMutex
}

func (_gg *RuneUint16Map) Length() int {
	_gg._abgd.RLock()
	defer _gg._abgd.RUnlock()
	return len(_gg._gd)
}
func (_efb *RuneStringMap) Range(f func(_cff rune, _bc string) (_cde bool)) {
	_efb._dg.RLock()
	defer _efb._dg.RUnlock()
	for _de, _gfd := range _efb._cda {
		if f(_de, _gfd) {
			break
		}
	}
}
func MakeRuneByteMap(length int) *RuneByteMap {
	_ea := make(map[rune]byte, length)
	return &RuneByteMap{_eac: _ea}
}
func (_ga *RuneSet) Write(r rune) {
	_ga._bgc.Lock()
	defer _ga._bgc.Unlock()
	_ga._cg[r] = struct{}{}
}
func NewStringsMap(tuples []StringsTuple) *StringsMap {
	_gadf := map[string]string{}
	for _, _baa := range tuples {
		_gadf[_baa.Key] = _baa.Value
	}
	return &StringsMap{_gfe: _gadf}
}
func (_cf *RuneSet) Range(f func(_gc rune) (_fga bool)) {
	_cf._bgc.RLock()
	defer _cf._bgc.RUnlock()
	for _bgf := range _cf._cg {
		if f(_bgf) {
			break
		}
	}
}
func NewRuneStringMap(m map[rune]string) *RuneStringMap { return &RuneStringMap{_cda: m} }
func (_gbdd *StringsMap) Copy() *StringsMap {
	_gbdd._edge.RLock()
	defer _gbdd._edge.RUnlock()
	_fgd := map[string]string{}
	for _cb, _bdf := range _gbdd._gfe {
		_fgd[_cb] = _bdf
	}
	return &StringsMap{_gfe: _fgd}
}
func MakeRuneUint16Map(length int) *RuneUint16Map {
	return &RuneUint16Map{_gd: make(map[rune]uint16, length)}
}

type StringsMap struct {
	_gfe  map[string]string
	_edge _a.RWMutex
}
type RuneUint16Map struct {
	_gd   map[rune]uint16
	_abgd _a.RWMutex
}
type StringRuneMap struct {
	_fgc map[string]rune
	_dd  _a.RWMutex
}

func (_cd *RuneSet) Exists(r rune) bool {
	_cd._bgc.RLock()
	defer _cd._bgc.RUnlock()
	_, _bb := _cd._cg[r]
	return _bb
}
func (_ff *RuneByteMap) Read(r rune) (byte, bool) {
	_ff._abd.RLock()
	defer _ff._abd.RUnlock()
	_edf, _ba := _ff._eac[r]
	return _edf, _ba
}

type RuneSet struct {
	_cg  map[rune]struct{}
	_bgc _a.RWMutex
}

func (_ag *ByteRuneMap) Write(b byte, r rune) {
	_ag._g.Lock()
	defer _ag._g.Unlock()
	_ag._ae[b] = r
}
func (_cdb *StringRuneMap) Read(g string) (rune, bool) {
	_cdb._dd.RLock()
	defer _cdb._dd.RUnlock()
	_cfgd, _bd := _cdb._fgc[g]
	return _cfgd, _bd
}
func (_cfg *RuneUint16Map) RangeDelete(f func(_efg rune, _bbg uint16) (_af bool, _gcfg bool)) {
	_cfg._abgd.Lock()
	defer _cfg._abgd.Unlock()
	for _dc, _gef := range _cfg._gd {
		_gad, _cag := f(_dc, _gef)
		if _gad {
			delete(_cfg._gd, _dc)
		}
		if _cag {
			break
		}
	}
}
func (_afa *RuneUint16Map) Delete(r rune) {
	_afa._abgd.Lock()
	defer _afa._abgd.Unlock()
	delete(_afa._gd, r)
}
func (_bcd *StringsMap) Range(f func(_gba, _fe string) (_bf bool)) {
	_bcd._edge.RLock()
	defer _bcd._edge.RUnlock()
	for _ad, _faa := range _bcd._gfe {
		if f(_ad, _faa) {
			break
		}
	}
}
func (_fd *StringsMap) Write(g1, g2 string) {
	_fd._edge.Lock()
	defer _fd._edge.Unlock()
	_fd._gfe[g1] = g2
}
func (_fca *RuneStringMap) Length() int {
	_fca._dg.RLock()
	defer _fca._dg.RUnlock()
	return len(_fca._cda)
}

type ByteRuneMap struct {
	_ae map[byte]rune
	_g  _a.RWMutex
}

func (_ddd *StringRuneMap) Length() int {
	_ddd._dd.RLock()
	defer _ddd._dd.RUnlock()
	return len(_ddd._fgc)
}
func (_aa *RuneUint16Map) Range(f func(_be rune, _ffd uint16) (_df bool)) {
	_aa._abgd.RLock()
	defer _aa._abgd.RUnlock()
	for _gb, _gccd := range _aa._gd {
		if f(_gb, _gccd) {
			break
		}
	}
}
func (_bg *ByteRuneMap) Range(f func(_fa byte, _fg rune) (_d bool)) {
	_bg._g.RLock()
	defer _bg._g.RUnlock()
	for _ab, _ed := range _bg._ae {
		if f(_ab, _ed) {
			break
		}
	}
}
func (_caa *RuneUint16Map) Read(r rune) (uint16, bool) {
	_caa._abgd.RLock()
	defer _caa._abgd.RUnlock()
	_gcc, _fgac := _caa._gd[r]
	return _gcc, _fgac
}

type StringsTuple struct {
	Key, Value string
}

func (_abg *RuneByteMap) Range(f func(_dad rune, _ca byte) (_fge bool)) {
	_abg._abd.RLock()
	defer _abg._abd.RUnlock()
	for _ef, _ac := range _abg._eac {
		if f(_ef, _ac) {
			break
		}
	}
}
func (_bcb *StringsMap) Read(g string) (string, bool) {
	_bcb._edge.RLock()
	defer _bcb._edge.RUnlock()
	_fae, _dee := _bcb._gfe[g]
	return _fae, _dee
}
func (_gcf *RuneUint16Map) Write(r rune, g uint16) {
	_gcf._abgd.Lock()
	defer _gcf._abgd.Unlock()
	_gcf._gd[r] = g
}
func (_eb *RuneByteMap) Write(r rune, b byte) {
	_eb._abd.Lock()
	defer _eb._abd.Unlock()
	_eb._eac[r] = b
}
func (_gadc *StringRuneMap) Range(f func(_gdg string, _gcg rune) (_ce bool)) {
	_gadc._dd.RLock()
	defer _gadc._dd.RUnlock()
	for _dec, _gbd := range _gadc._fgc {
		if f(_dec, _gbd) {
			break
		}
	}
}
func (_fff *RuneByteMap) Length() int {
	_fff._abd.RLock()
	defer _fff._abd.RUnlock()
	return len(_fff._eac)
}
func (_bbc *StringRuneMap) Write(g string, r rune) {
	_bbc._dd.Lock()
	defer _bbc._dd.Unlock()
	_bbc._fgc[g] = r
}
