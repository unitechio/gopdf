package syncmap

import _ee "sync"

func NewRuneStringMap(m map[rune]string) *RuneStringMap { return &RuneStringMap{_ff: m} }
func (_fd *RuneByteMap) Length() int                    { _fd._fg.RLock(); defer _fd._fg.RUnlock(); return len(_fd._fcf) }
func (_ggd *StringsMap) Copy() *StringsMap {
	_ggd._gfa.RLock()
	defer _ggd._gfa.RUnlock()
	_gfg := map[string]string{}
	for _cddd, _cac := range _ggd._fdd {
		_gfg[_cddd] = _cac
	}
	return &StringsMap{_fdd: _gfg}
}

type RuneSet struct {
	_faf map[rune]struct{}
	_gbd _ee.RWMutex
}

func (_fbdc *RuneSet) Length() int {
	_fbdc._gbd.RLock()
	defer _fbdc._gbd.RUnlock()
	return len(_fbdc._faf)
}
func NewStringsMap(tuples []StringsTuple) *StringsMap {
	_ffc := map[string]string{}
	for _, _bbb := range tuples {
		_ffc[_bbb.Key] = _bbb.Value
	}
	return &StringsMap{_fdd: _ffc}
}
func (_b *ByteRuneMap) Range(f func(_fc byte, _g rune) (_c bool)) {
	_b._d.RLock()
	defer _b._d.RUnlock()
	for _de, _dd := range _b._a {
		if f(_de, _dd) {
			break
		}
	}
}
func (_ca *RuneByteMap) Read(r rune) (byte, bool) {
	_ca._fg.RLock()
	defer _ca._fg.RUnlock()
	_ed, _ag := _ca._fcf[r]
	return _ed, _ag
}
func (_decf *RuneUint16Map) Range(f func(_fag rune, _ecb uint16) (_gba bool)) {
	_decf._ba.RLock()
	defer _decf._ba.RUnlock()
	for _ffe, _eag := range _decf._gc {
		if f(_ffe, _eag) {
			break
		}
	}
}

type RuneStringMap struct {
	_ff map[rune]string
	_ea _ee.RWMutex
}

func NewByteRuneMap(m map[byte]rune) *ByteRuneMap { return &ByteRuneMap{_a: m} }
func (_fce *StringsMap) Write(g1, g2 string) {
	_fce._gfa.Lock()
	defer _fce._gfa.Unlock()
	_fce._fdd[g1] = g2
}
func (_eg *ByteRuneMap) Write(b byte, r rune) { _eg._d.Lock(); defer _eg._d.Unlock(); _eg._a[b] = r }
func (_eed *RuneStringMap) Length() int {
	_eed._ea.RLock()
	defer _eed._ea.RUnlock()
	return len(_eed._ff)
}
func (_dfa *StringsMap) Read(g string) (string, bool) {
	_dfa._gfa.RLock()
	defer _dfa._gfa.RUnlock()
	_efb, _ae := _dfa._fdd[g]
	return _efb, _ae
}
func (_abg *StringRuneMap) Range(f func(_gcbe string, _dgb rune) (_bbg bool)) {
	_abg._fgf.RLock()
	defer _abg._fgf.RUnlock()
	for _ggc, _ede := range _abg._egb {
		if f(_ggc, _ede) {
			break
		}
	}
}
func (_cg *RuneUint16Map) Read(r rune) (uint16, bool) {
	_cg._ba.RLock()
	defer _cg._ba.RUnlock()
	_adb, _cdd := _cg._gc[r]
	return _adb, _cdd
}

type ByteRuneMap struct {
	_a map[byte]rune
	_d _ee.RWMutex
}

func (_f *ByteRuneMap) Read(b byte) (rune, bool) {
	_f._d.RLock()
	defer _f._d.RUnlock()
	_ef, _aa := _f._a[b]
	return _ef, _aa
}
func (_gfd *StringRuneMap) Length() int {
	_gfd._fgf.RLock()
	defer _gfd._fgf.RUnlock()
	return len(_gfd._egb)
}
func (_gcb *StringRuneMap) Read(g string) (rune, bool) {
	_gcb._fgf.RLock()
	defer _gcb._fgf.RUnlock()
	_fbg, _bcf := _gcb._egb[g]
	return _fbg, _bcf
}
func (_bbba *StringsMap) Range(f func(_cee, _fea string) (_ecd bool)) {
	_bbba._gfa.RLock()
	defer _bbba._gfa.RUnlock()
	for _aab, _efda := range _bbba._fdd {
		if f(_aab, _efda) {
			break
		}
	}
}
func (_ab *RuneUint16Map) RangeDelete(f func(_dg rune, _bc uint16) (_aga bool, _fbc bool)) {
	_ab._ba.Lock()
	defer _ab._ba.Unlock()
	for _cfe, _fge := range _ab._gc {
		_dgd, _cfbe := f(_cfe, _fge)
		if _dgd {
			delete(_ab._gc, _cfe)
		}
		if _cfbe {
			break
		}
	}
}
func NewStringRuneMap(m map[string]rune) *StringRuneMap { return &StringRuneMap{_egb: m} }
func (_fe *RuneUint16Map) Delete(r rune)                { _fe._ba.Lock(); defer _fe._ba.Unlock(); delete(_fe._gc, r) }
func (_ge *RuneUint16Map) Write(r rune, g uint16) {
	_ge._ba.Lock()
	defer _ge._ba.Unlock()
	_ge._gc[r] = g
}
func (_ac *RuneSet) Write(r rune) { _ac._gbd.Lock(); defer _ac._gbd.Unlock(); _ac._faf[r] = struct{}{} }

type StringsTuple struct {
	Key, Value string
}

func MakeRuneByteMap(length int) *RuneByteMap {
	_ddf := make(map[rune]byte, length)
	return &RuneByteMap{_fcf: _ddf}
}
func MakeByteRuneMap(length int) *ByteRuneMap { return &ByteRuneMap{_a: make(map[byte]rune, length)} }
func (_fa *ByteRuneMap) Length() int {
	_fa._d.RLock()
	defer _fa._d.RUnlock()
	return len(_fa._a)
}
func (_da *RuneUint16Map) Length() int { _da._ba.RLock(); defer _da._ba.RUnlock(); return len(_da._gc) }
func (_dec *RuneStringMap) Range(f func(_dc rune, _eeb string) (_gbe bool)) {
	_dec._ea.RLock()
	defer _dec._ea.RUnlock()
	for _ec, _cfb := range _dec._ff {
		if f(_ec, _cfb) {
			break
		}
	}
}

type RuneUint16Map struct {
	_gc map[rune]uint16
	_ba _ee.RWMutex
}

func (_af *RuneStringMap) Read(r rune) (string, bool) {
	_af._ea.RLock()
	defer _af._ea.RUnlock()
	_cf, _bb := _af._ff[r]
	return _cf, _bb
}
func (_fcfe *RuneStringMap) Write(r rune, s string) {
	_fcfe._ea.Lock()
	defer _fcfe._ea.Unlock()
	_fcfe._ff[r] = s
}
func MakeRuneSet(length int) *RuneSet { return &RuneSet{_faf: make(map[rune]struct{}, length)} }
func (_fb *RuneSet) Exists(r rune) bool {
	_fb._gbd.RLock()
	defer _fb._gbd.RUnlock()
	_, _fbd := _fb._faf[r]
	return _fbd
}
func (_df *RuneByteMap) Write(r rune, b byte) {
	_df._fg.Lock()
	defer _df._fg.Unlock()
	_df._fcf[r] = b
}

type StringRuneMap struct {
	_egb map[string]rune
	_fgf _ee.RWMutex
}

func (_ddfa *RuneByteMap) Range(f func(_eea rune, _gb byte) (_gf bool)) {
	_ddfa._fg.RLock()
	defer _ddfa._fg.RUnlock()
	for _ad, _efd := range _ddfa._fcf {
		if f(_ad, _efd) {
			break
		}
	}
}
func (_fbe *StringRuneMap) Write(g string, r rune) {
	_fbe._fgf.Lock()
	defer _fbe._fgf.Unlock()
	_fbe._egb[g] = r
}
func (_ce *RuneSet) Range(f func(_cd rune) (_gg bool)) {
	_ce._gbd.RLock()
	defer _ce._gbd.RUnlock()
	for _agf := range _ce._faf {
		if f(_agf) {
			break
		}
	}
}

type StringsMap struct {
	_fdd map[string]string
	_gfa _ee.RWMutex
}

func MakeRuneUint16Map(length int) *RuneUint16Map {
	return &RuneUint16Map{_gc: make(map[rune]uint16, length)}
}

type RuneByteMap struct {
	_fcf map[rune]byte
	_fg  _ee.RWMutex
}
