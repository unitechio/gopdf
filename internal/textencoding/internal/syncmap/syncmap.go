package syncmap

import _f "sync"

func MakeRuneByteMap(length int) *RuneByteMap {
	_fd := make(map[rune]byte, length)
	return &RuneByteMap{_bb: _fd}
}
func (_ee *ByteRuneMap) Write(b byte, r rune) { _ee._c.Lock(); defer _ee._c.Unlock(); _ee._e[b] = r }
func (_beb *RuneSet) Range(f func(_daee rune) (_egc bool)) {
	_beb._cef.RLock()
	defer _beb._cef.RUnlock()
	for _eb := range _beb._df {
		if f(_eb) {
			break
		}
	}
}
func NewByteRuneMap(m map[byte]rune) *ByteRuneMap { return &ByteRuneMap{_e: m} }
func (_fdc *StringsMap) Write(g1, g2 string) {
	_fdc._dfdd.Lock()
	defer _fdc._dfdd.Unlock()
	_fdc._fdf[g1] = g2
}
func (_fc *RuneByteMap) Range(f func(_cf rune, _da byte) (_bf bool)) {
	_fc._ce.RLock()
	defer _fc._ce.RUnlock()
	for _cg, _dae := range _fc._bb {
		if f(_cg, _dae) {
			break
		}
	}
}
func (_aea *StringRuneMap) Read(g string) (rune, bool) {
	_aea._dgc.RLock()
	defer _aea._dgc.RUnlock()
	_ab, _bebd := _aea._eae[g]
	return _ab, _bebd
}
func (_ga *RuneByteMap) Read(r rune) (byte, bool) {
	_ga._ce.RLock()
	defer _ga._ce.RUnlock()
	_gb, _eg := _ga._bb[r]
	return _gb, _eg
}

type RuneByteMap struct {
	_bb map[rune]byte
	_ce _f.RWMutex
}

func NewStringRuneMap(m map[string]rune) *StringRuneMap { return &StringRuneMap{_eae: m} }

type RuneStringMap struct {
	_agb map[rune]string
	_ad  _f.RWMutex
}

func (_bc *ByteRuneMap) Range(f func(_bcc byte, _a rune) (_d bool)) {
	_bc._c.RLock()
	defer _bc._c.RUnlock()
	for _be, _gd := range _bc._e {
		if f(_be, _gd) {
			break
		}
	}
}
func (_gg *RuneUint16Map) Write(r rune, g uint16) {
	_gg._bea.Lock()
	defer _gg._bea.Unlock()
	_gg._aed[r] = g
}
func (_cea *RuneUint16Map) Range(f func(_gc rune, _gab uint16) (_cda bool)) {
	_cea._bea.RLock()
	defer _cea._bea.RUnlock()
	for _bec, _aga := range _cea._aed {
		if f(_bec, _aga) {
			break
		}
	}
}
func (_gf *ByteRuneMap) Read(b byte) (rune, bool) {
	_gf._c.RLock()
	defer _gf._c.RUnlock()
	_ea, _b := _gf._e[b]
	return _ea, _b
}
func (_bg *ByteRuneMap) Length() int {
	_bg._c.RLock()
	defer _bg._c.RUnlock()
	return len(_bg._e)
}

type ByteRuneMap struct {
	_e map[byte]rune
	_c _f.RWMutex
}
type StringRuneMap struct {
	_eae map[string]rune
	_dgc _f.RWMutex
}

func NewStringsMap(tuples []StringsTuple) *StringsMap {
	_beg := map[string]string{}
	for _, _dac := range tuples {
		_beg[_dac.Key] = _dac.Value
	}
	return &StringsMap{_fdf: _beg}
}
func (_ae *RuneSet) Length() int { _ae._cef.RLock(); defer _ae._cef.RUnlock(); return len(_ae._df) }

type RuneUint16Map struct {
	_aed map[rune]uint16
	_bea _f.RWMutex
}

func (_ag *RuneSet) Exists(r rune) bool {
	_ag._cef.RLock()
	defer _ag._cef.RUnlock()
	_, _ff := _ag._df[r]
	return _ff
}
func (_cec *RuneStringMap) Read(r rune) (string, bool) {
	_cec._ad.RLock()
	defer _cec._ad.RUnlock()
	_gba, _eab := _cec._agb[r]
	return _gba, _eab
}
func (_db *StringRuneMap) Range(f func(_dfdb string, _feb rune) (_de bool)) {
	_db._dgc.RLock()
	defer _db._dgc.RUnlock()
	for _bcd, _adg := range _db._eae {
		if f(_bcd, _adg) {
			break
		}
	}
}
func (_ceac *RuneUint16Map) Length() int {
	_ceac._bea.RLock()
	defer _ceac._bea.RUnlock()
	return len(_ceac._aed)
}
func (_ca *RuneByteMap) Write(r rune, b byte) { _ca._ce.Lock(); defer _ca._ce.Unlock(); _ca._bb[r] = b }
func (_afb *StringsMap) Range(f func(_fed, _fb string) (_caa bool)) {
	_afb._dfdd.RLock()
	defer _afb._dfdd.RUnlock()
	for _cad, _eea := range _afb._fdf {
		if f(_cad, _eea) {
			break
		}
	}
}
func (_ac *RuneUint16Map) Read(r rune) (uint16, bool) {
	_ac._bea.RLock()
	defer _ac._bea.RUnlock()
	_fe, _cfg := _ac._aed[r]
	return _fe, _cfg
}
func (_cefc *StringRuneMap) Length() int {
	_cefc._dgc.RLock()
	defer _cefc._dgc.RUnlock()
	return len(_cefc._eae)
}

type RuneSet struct {
	_df  map[rune]struct{}
	_cef _f.RWMutex
}
type StringsMap struct {
	_fdf  map[string]string
	_dfdd _f.RWMutex
}

func (_ebe *RuneUint16Map) Delete(r rune) {
	_ebe._bea.Lock()
	defer _ebe._bea.Unlock()
	delete(_ebe._aed, r)
}
func NewRuneStringMap(m map[rune]string) *RuneStringMap { return &RuneStringMap{_agb: m} }
func (_bag *RuneUint16Map) RangeDelete(f func(_fff rune, _bgb uint16) (_dd bool, _ef bool)) {
	_bag._bea.Lock()
	defer _bag._bea.Unlock()
	for _ffg, _af := range _bag._aed {
		_egb, _dg := f(_ffg, _af)
		if _egb {
			delete(_bag._aed, _ffg)
		}
		if _dg {
			break
		}
	}
}
func MakeRuneSet(length int) *RuneSet { return &RuneSet{_df: make(map[rune]struct{}, length)} }
func (_gdg *RuneStringMap) Length() int {
	_gdg._ad.RLock()
	defer _gdg._ad.RUnlock()
	return len(_gdg._agb)
}
func (_ba *RuneSet) Write(r rune) { _ba._cef.Lock(); defer _ba._cef.Unlock(); _ba._df[r] = struct{}{} }
func (_gfb *StringsMap) Read(g string) (string, bool) {
	_gfb._dfdd.RLock()
	defer _gfb._dfdd.RUnlock()
	_acb, _dff := _gfb._fdf[g]
	return _acb, _dff
}
func (_egd *RuneByteMap) Length() int {
	_egd._ce.RLock()
	defer _egd._ce.RUnlock()
	return len(_egd._bb)
}
func (_gabg *StringsMap) Copy() *StringsMap {
	_gabg._dfdd.RLock()
	defer _gabg._dfdd.RUnlock()
	_baa := map[string]string{}
	for _bba, _bbb := range _gabg._fdf {
		_baa[_bba] = _bbb
	}
	return &StringsMap{_fdf: _baa}
}
func MakeByteRuneMap(length int) *ByteRuneMap { return &ByteRuneMap{_e: make(map[byte]rune, length)} }
func (_dfd *RuneStringMap) Write(r rune, s string) {
	_dfd._ad.Lock()
	defer _dfd._ad.Unlock()
	_dfd._agb[r] = s
}
func MakeRuneUint16Map(length int) *RuneUint16Map {
	return &RuneUint16Map{_aed: make(map[rune]uint16, length)}
}
func (_aa *StringRuneMap) Write(g string, r rune) {
	_aa._dgc.Lock()
	defer _aa._dgc.Unlock()
	_aa._eae[g] = r
}

type StringsTuple struct{ Key, Value string }

func (_gae *RuneStringMap) Range(f func(_cd rune, _bccc string) (_age bool)) {
	_gae._ad.RLock()
	defer _gae._ad.RUnlock()
	for _gfe, _bbe := range _gae._agb {
		if f(_gfe, _bbe) {
			break
		}
	}
}
