package ps

import (
	_ad "bufio"
	_ga "bytes"
	_a "errors"
	_d "fmt"
	_e "io"
	_b "math"

	_aa "bitbucket.org/shenghui0779/gopdf/common"
	_be "bitbucket.org/shenghui0779/gopdf/core"
)

func (_babc *PSParser) parseBool() (*PSBoolean, error) {
	_fcd, _gge := _babc._bebfa.Peek(4)
	if _gge != nil {
		return MakeBool(false), _gge
	}
	if (len(_fcd) >= 4) && (string(_fcd[:4]) == "\u0074\u0072\u0075\u0065") {
		_babc._bebfa.Discard(4)
		return MakeBool(true), nil
	}
	_fcd, _gge = _babc._bebfa.Peek(5)
	if _gge != nil {
		return MakeBool(false), _gge
	}
	if (len(_fcd) >= 5) && (string(_fcd[:5]) == "\u0066\u0061\u006cs\u0065") {
		_babc._bebfa.Discard(5)
		return MakeBool(false), nil
	}
	return MakeBool(false), _a.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

// PopInteger specificially pops an integer from the top of the stack, returning the value as an int.
func (_bcfd *PSStack) PopInteger() (int, error) {
	_fceb, _aaad := _bcfd.Pop()
	if _aaad != nil {
		return 0, _aaad
	}
	if _abc, _ceae := _fceb.(*PSInteger); _ceae {
		return _abc.Val, nil
	}
	return 0, ErrTypeCheck
}
func (_eg *PSReal) Duplicate() PSObject {
	_gfg := PSReal{}
	_gfg.Val = _eg.Val
	return &_gfg
}

// String returns a string representation of the stack.
func (_bgad *PSStack) String() string {
	_gfcb := "\u005b\u0020"
	for _, _gdc := range *_bgad {
		_gfcb += _gdc.String()
		_gfcb += "\u0020"
	}
	_gfcb += "\u005d"
	return _gfcb
}
func (_gce *PSOperand) Duplicate() PSObject { _fd := *_gce; return &_fd }
func (_fb *PSOperand) add(_cga *PSStack) error {
	_fca, _cd := _cga.Pop()
	if _cd != nil {
		return _cd
	}
	_ac, _cd := _cga.Pop()
	if _cd != nil {
		return _cd
	}
	_fg, _bba := _fca.(*PSReal)
	_cgcb, _df := _fca.(*PSInteger)
	if !_bba && !_df {
		return ErrTypeCheck
	}
	_cc, _aba := _ac.(*PSReal)
	_bec, _baf := _ac.(*PSInteger)
	if !_aba && !_baf {
		return ErrTypeCheck
	}
	if _df && _baf {
		_aef := _cgcb.Val + _bec.Val
		_fcc := _cga.Push(MakeInteger(_aef))
		return _fcc
	}
	var _cca float64
	if _bba {
		_cca = _fg.Val
	} else {
		_cca = float64(_cgcb.Val)
	}
	if _aba {
		_cca += _cc.Val
	} else {
		_cca += float64(_bec.Val)
	}
	_cd = _cga.Push(MakeReal(_cca))
	return _cd
}

var ErrRangeCheck = _a.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")

func (_gb *PSOperand) cos(_beg *PSStack) error {
	_cbb, _ffb := _beg.PopNumberAsFloat64()
	if _ffb != nil {
		return _ffb
	}
	_dgc := _b.Cos(_cbb * _b.Pi / 180.0)
	_ffb = _beg.Push(MakeReal(_dgc))
	return _ffb
}

// PSExecutor has its own execution stack and is used to executre a PS routine (program).
type PSExecutor struct {
	Stack *PSStack
	_f    *PSProgram
}

func (_gff *PSOperand) exp(_acf *PSStack) error {
	_dcf, _fge := _acf.PopNumberAsFloat64()
	if _fge != nil {
		return _fge
	}
	_gga, _fge := _acf.PopNumberAsFloat64()
	if _fge != nil {
		return _fge
	}
	if _b.Abs(_dcf) < 1 && _gga < 0 {
		return ErrUndefinedResult
	}
	_ded := _b.Pow(_gga, _dcf)
	_fge = _acf.Push(MakeReal(_ded))
	return _fge
}

// NewPSProgram returns an empty, initialized PSProgram.
func NewPSProgram() *PSProgram { return &PSProgram{} }
func (_fa *PSOperand) cvi(_egb *PSStack) error {
	_defc, _ffa := _egb.Pop()
	if _ffa != nil {
		return _ffa
	}
	if _gfd, _gecb := _defc.(*PSReal); _gecb {
		_bgg := int(_gfd.Val)
		_ffa = _egb.Push(MakeInteger(_bgg))
	} else if _ee, _ddd := _defc.(*PSInteger); _ddd {
		_dcc := _ee.Val
		_ffa = _egb.Push(MakeInteger(_dcc))
	} else {
		return ErrTypeCheck
	}
	return _ffa
}

// PSObject represents a postscript object.
type PSObject interface {

	// Duplicate makes a fresh copy of the PSObject.
	Duplicate() PSObject

	// DebugString returns a descriptive representation of the PSObject with more information than String()
	// for debugging purposes.
	DebugString() string

	// String returns a string representation of the PSObject.
	String() string
}

// MakeInteger returns a new PSInteger object initialized with `val`.
func MakeInteger(val int) *PSInteger { _afb := PSInteger{}; _afb.Val = val; return &_afb }

var ErrUnsupportedOperand = _a.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")

// Empty empties the stack.
func (_bgcc *PSStack) Empty() { *_bgcc = []PSObject{} }

// Push pushes an object on top of the stack.
func (_eca *PSStack) Push(obj PSObject) error {
	if len(*_eca) > 100 {
		return ErrStackOverflow
	}
	*_eca = append(*_eca, obj)
	return nil
}
func (_bcf *PSOperand) or(_gac *PSStack) error {
	_ece, _aafc := _gac.Pop()
	if _aafc != nil {
		return _aafc
	}
	_cfb, _aafc := _gac.Pop()
	if _aafc != nil {
		return _aafc
	}
	if _abf, _eae := _ece.(*PSBoolean); _eae {
		_fbc, _bfde := _cfb.(*PSBoolean)
		if !_bfde {
			return ErrTypeCheck
		}
		_aafc = _gac.Push(MakeBool(_abf.Val || _fbc.Val))
		return _aafc
	}
	if _bfec, _aegf := _ece.(*PSInteger); _aegf {
		_gaa, _abd := _cfb.(*PSInteger)
		if !_abd {
			return ErrTypeCheck
		}
		_aafc = _gac.Push(MakeInteger(_bfec.Val | _gaa.Val))
		return _aafc
	}
	return ErrTypeCheck
}
func (_dc *PSOperand) DebugString() string {
	return _d.Sprintf("\u006fp\u003a\u0027\u0025\u0073\u0027", *_dc)
}
func (_cac *PSOperand) abs(_efa *PSStack) error {
	_ba, _cacg := _efa.Pop()
	if _cacg != nil {
		return _cacg
	}
	if _af, _bfca := _ba.(*PSReal); _bfca {
		_dg := _af.Val
		if _dg < 0 {
			_cacg = _efa.Push(MakeReal(-_dg))
		} else {
			_cacg = _efa.Push(MakeReal(_dg))
		}
	} else if _dec, _gced := _ba.(*PSInteger); _gced {
		_edc := _dec.Val
		if _edc < 0 {
			_cacg = _efa.Push(MakeInteger(-_edc))
		} else {
			_cacg = _efa.Push(MakeInteger(_edc))
		}
	} else {
		return ErrTypeCheck
	}
	return _cacg
}
func (_bde *PSOperand) div(_age *PSStack) error {
	_bfe, _cgg := _age.Pop()
	if _cgg != nil {
		return _cgg
	}
	_dbd, _cgg := _age.Pop()
	if _cgg != nil {
		return _cgg
	}
	_eafg, _begf := _bfe.(*PSReal)
	_cfa, _baa := _bfe.(*PSInteger)
	if !_begf && !_baa {
		return ErrTypeCheck
	}
	if _begf && _eafg.Val == 0 {
		return ErrUndefinedResult
	}
	if _baa && _cfa.Val == 0 {
		return ErrUndefinedResult
	}
	_dcd, _ddfa := _dbd.(*PSReal)
	_dbf, _dad := _dbd.(*PSInteger)
	if !_ddfa && !_dad {
		return ErrTypeCheck
	}
	var _fed float64
	if _ddfa {
		_fed = _dcd.Val
	} else {
		_fed = float64(_dbf.Val)
	}
	if _begf {
		_fed /= _eafg.Val
	} else {
		_fed /= float64(_cfa.Val)
	}
	_cgg = _age.Push(MakeReal(_fed))
	return _cgg
}

var ErrTypeCheck = _a.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")

func (_fdd *PSOperand) atan(_efae *PSStack) error {
	_aacc, _cce := _efae.PopNumberAsFloat64()
	if _cce != nil {
		return _cce
	}
	_cdf, _cce := _efae.PopNumberAsFloat64()
	if _cce != nil {
		return _cce
	}
	if _aacc == 0 {
		var _bef error
		if _cdf < 0 {
			_bef = _efae.Push(MakeReal(270))
		} else {
			_bef = _efae.Push(MakeReal(90))
		}
		return _bef
	}
	_deg := _cdf / _aacc
	_dbb := _b.Atan(_deg) * 180 / _b.Pi
	_cce = _efae.Push(MakeReal(_dbb))
	return _cce
}

// Exec executes the operand `op` in the state specified by `stack`.
func (_cb *PSOperand) Exec(stack *PSStack) error {
	_cgcd := ErrUnsupportedOperand
	switch *_cb {
	case "\u0061\u0062\u0073":
		_cgcd = _cb.abs(stack)
	case "\u0061\u0064\u0064":
		_cgcd = _cb.add(stack)
	case "\u0061\u006e\u0064":
		_cgcd = _cb.and(stack)
	case "\u0061\u0074\u0061\u006e":
		_cgcd = _cb.atan(stack)
	case "\u0062\u0069\u0074\u0073\u0068\u0069\u0066\u0074":
		_cgcd = _cb.bitshift(stack)
	case "\u0063e\u0069\u006c\u0069\u006e\u0067":
		_cgcd = _cb.ceiling(stack)
	case "\u0063\u006f\u0070\u0079":
		_cgcd = _cb.copy(stack)
	case "\u0063\u006f\u0073":
		_cgcd = _cb.cos(stack)
	case "\u0063\u0076\u0069":
		_cgcd = _cb.cvi(stack)
	case "\u0063\u0076\u0072":
		_cgcd = _cb.cvr(stack)
	case "\u0064\u0069\u0076":
		_cgcd = _cb.div(stack)
	case "\u0064\u0075\u0070":
		_cgcd = _cb.dup(stack)
	case "\u0065\u0071":
		_cgcd = _cb.eq(stack)
	case "\u0065\u0078\u0063\u0068":
		_cgcd = _cb.exch(stack)
	case "\u0065\u0078\u0070":
		_cgcd = _cb.exp(stack)
	case "\u0066\u006c\u006fo\u0072":
		_cgcd = _cb.floor(stack)
	case "\u0067\u0065":
		_cgcd = _cb.ge(stack)
	case "\u0067\u0074":
		_cgcd = _cb.gt(stack)
	case "\u0069\u0064\u0069\u0076":
		_cgcd = _cb.idiv(stack)
	case "\u0069\u0066":
		_cgcd = _cb.ifCondition(stack)
	case "\u0069\u0066\u0065\u006c\u0073\u0065":
		_cgcd = _cb.ifelse(stack)
	case "\u0069\u006e\u0064e\u0078":
		_cgcd = _cb.index(stack)
	case "\u006c\u0065":
		_cgcd = _cb.le(stack)
	case "\u006c\u006f\u0067":
		_cgcd = _cb.log(stack)
	case "\u006c\u006e":
		_cgcd = _cb.ln(stack)
	case "\u006c\u0074":
		_cgcd = _cb.lt(stack)
	case "\u006d\u006f\u0064":
		_cgcd = _cb.mod(stack)
	case "\u006d\u0075\u006c":
		_cgcd = _cb.mul(stack)
	case "\u006e\u0065":
		_cgcd = _cb.ne(stack)
	case "\u006e\u0065\u0067":
		_cgcd = _cb.neg(stack)
	case "\u006e\u006f\u0074":
		_cgcd = _cb.not(stack)
	case "\u006f\u0072":
		_cgcd = _cb.or(stack)
	case "\u0070\u006f\u0070":
		_cgcd = _cb.pop(stack)
	case "\u0072\u006f\u0075n\u0064":
		_cgcd = _cb.round(stack)
	case "\u0072\u006f\u006c\u006c":
		_cgcd = _cb.roll(stack)
	case "\u0073\u0069\u006e":
		_cgcd = _cb.sin(stack)
	case "\u0073\u0071\u0072\u0074":
		_cgcd = _cb.sqrt(stack)
	case "\u0073\u0075\u0062":
		_cgcd = _cb.sub(stack)
	case "\u0074\u0072\u0075\u006e\u0063\u0061\u0074\u0065":
		_cgcd = _cb.truncate(stack)
	case "\u0078\u006f\u0072":
		_cgcd = _cb.xor(stack)
	}
	return _cgcd
}
func (_aaf *PSOperand) bitshift(_ddf *PSStack) error {
	_bge, _gceg := _ddf.PopInteger()
	if _gceg != nil {
		return _gceg
	}
	_abe, _gceg := _ddf.PopInteger()
	if _gceg != nil {
		return _gceg
	}
	var _ag int
	if _bge >= 0 {
		_ag = _abe << uint(_bge)
	} else {
		_ag = _abe >> uint(-_bge)
	}
	_gceg = _ddf.Push(MakeInteger(_ag))
	return _gceg
}

// MakeBool returns a new PSBoolean object initialized with `val`.
func MakeBool(val bool) *PSBoolean { _egbc := PSBoolean{}; _egbc.Val = val; return &_egbc }

// PSOperand represents a Postscript operand (text string).
type PSOperand string

func (_fbb *PSOperand) mod(_dde *PSStack) error {
	_cgcc, _dfg := _dde.Pop()
	if _dfg != nil {
		return _dfg
	}
	_aeacf, _dfg := _dde.Pop()
	if _dfg != nil {
		return _dfg
	}
	_bag, _bfdf := _cgcc.(*PSInteger)
	if !_bfdf {
		return ErrTypeCheck
	}
	if _bag.Val == 0 {
		return ErrUndefinedResult
	}
	_deca, _bfdf := _aeacf.(*PSInteger)
	if !_bfdf {
		return ErrTypeCheck
	}
	_fcbg := _deca.Val % _bag.Val
	_dfg = _dde.Push(MakeInteger(_fcbg))
	return _dfg
}
func (_bfcce *PSParser) parseNumber() (PSObject, error) {
	_cdfe, _dfa := _be.ParseNumber(_bfcce._bebfa)
	if _dfa != nil {
		return nil, _dfa
	}
	switch _fda := _cdfe.(type) {
	case *_be.PdfObjectFloat:
		return MakeReal(float64(*_fda)), nil
	case *_be.PdfObjectInteger:
		return MakeInteger(int(*_fda)), nil
	}
	return nil, _d.Errorf("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _cdfe)
}
func (_def *PSOperand) copy(_bed *PSStack) error {
	_bfcag, _dda := _bed.PopInteger()
	if _dda != nil {
		return _dda
	}
	if _bfcag < 0 {
		return ErrRangeCheck
	}
	if _bfcag > len(*_bed) {
		return ErrRangeCheck
	}
	*_bed = append(*_bed, (*_bed)[len(*_bed)-_bfcag:]...)
	return nil
}
func (_ffbd *PSOperand) ifCondition(_cgab *PSStack) error {
	_ced, _cgb := _cgab.Pop()
	if _cgb != nil {
		return _cgb
	}
	_fbe, _cgb := _cgab.Pop()
	if _cgb != nil {
		return _cgb
	}
	_dfc, _dege := _ced.(*PSProgram)
	if !_dege {
		return ErrTypeCheck
	}
	_eeg, _dege := _fbe.(*PSBoolean)
	if !_dege {
		return ErrTypeCheck
	}
	if _eeg.Val {
		_gfga := _dfc.Exec(_cgab)
		return _gfga
	}
	return nil
}

// PSObjectArrayToFloat64Array converts []PSObject into a []float64 array. Each PSObject must represent a number,
// otherwise a ErrTypeCheck error occurs.
func PSObjectArrayToFloat64Array(objects []PSObject) ([]float64, error) {
	var _gc []float64
	for _, _c := range objects {
		if _beb, _ae := _c.(*PSInteger); _ae {
			_gc = append(_gc, float64(_beb.Val))
		} else if _da, _eag := _c.(*PSReal); _eag {
			_gc = append(_gc, _da.Val)
		} else {
			return nil, ErrTypeCheck
		}
	}
	return _gc, nil
}
func (_gg *PSProgram) String() string {
	_dd := "\u007b\u0020"
	for _, _ca := range *_gg {
		_dd += _ca.String()
		_dd += "\u0020"
	}
	_dd += "\u007d"
	return _dd
}
func (_fdef *PSOperand) index(_gda *PSStack) error {
	_fcec, _cbe := _gda.Pop()
	if _cbe != nil {
		return _cbe
	}
	_bgfd, _cbeb := _fcec.(*PSInteger)
	if !_cbeb {
		return ErrTypeCheck
	}
	if _bgfd.Val < 0 {
		return ErrRangeCheck
	}
	if _bgfd.Val > len(*_gda)-1 {
		return ErrStackUnderflow
	}
	_afg := (*_gda)[len(*_gda)-1-_bgfd.Val]
	_cbe = _gda.Push(_afg.Duplicate())
	return _cbe
}

// Pop pops an object from the top of the stack.
func (_fgac *PSStack) Pop() (PSObject, error) {
	if len(*_fgac) < 1 {
		return nil, ErrStackUnderflow
	}
	_dab := (*_fgac)[len(*_fgac)-1]
	*_fgac = (*_fgac)[0 : len(*_fgac)-1]
	return _dab, nil
}
func (_ead *PSReal) DebugString() string {
	return _d.Sprintf("\u0072e\u0061\u006c\u003a\u0025\u002e\u0035f", _ead.Val)
}
func (_cdgc *PSOperand) exch(_ccd *PSStack) error {
	_faf, _aeec := _ccd.Pop()
	if _aeec != nil {
		return _aeec
	}
	_fce, _aeec := _ccd.Pop()
	if _aeec != nil {
		return _aeec
	}
	_aeec = _ccd.Push(_faf)
	if _aeec != nil {
		return _aeec
	}
	_aeec = _ccd.Push(_fce)
	return _aeec
}
func (_eea *PSOperand) sub(_dffg *PSStack) error {
	_dfgb, _fgc := _dffg.Pop()
	if _fgc != nil {
		return _fgc
	}
	_cbed, _fgc := _dffg.Pop()
	if _fgc != nil {
		return _fgc
	}
	_dbag, _gfdb := _dfgb.(*PSReal)
	_dgbe, _afd := _dfgb.(*PSInteger)
	if !_gfdb && !_afd {
		return ErrTypeCheck
	}
	_beef, _bgdd := _cbed.(*PSReal)
	_dggb, _dcfd := _cbed.(*PSInteger)
	if !_bgdd && !_dcfd {
		return ErrTypeCheck
	}
	if _afd && _dcfd {
		_acdf := _dggb.Val - _dgbe.Val
		_abg := _dffg.Push(MakeInteger(_acdf))
		return _abg
	}
	var _fgbc float64 = 0
	if _bgdd {
		_fgbc = _beef.Val
	} else {
		_fgbc = float64(_dggb.Val)
	}
	if _gfdb {
		_fgbc -= _dbag.Val
	} else {
		_fgbc -= float64(_dgbe.Val)
	}
	_fgc = _dffg.Push(MakeReal(_fgbc))
	return _fgc
}
func (_geeb *PSOperand) ge(_geebb *PSStack) error {
	_dcde, _abec := _geebb.PopNumberAsFloat64()
	if _abec != nil {
		return _abec
	}
	_gecg, _abec := _geebb.PopNumberAsFloat64()
	if _abec != nil {
		return _abec
	}
	if _b.Abs(_gecg-_dcde) < _aac {
		_fdg := _geebb.Push(MakeBool(true))
		return _fdg
	} else if _gecg > _dcde {
		_cbbf := _geebb.Push(MakeBool(true))
		return _cbbf
	} else {
		_bdb := _geebb.Push(MakeBool(false))
		return _bdb
	}
}
func (_ddb *PSOperand) floor(_ggb *PSStack) error {
	_ged, _caa := _ggb.Pop()
	if _caa != nil {
		return _caa
	}
	if _gbbf, _efaf := _ged.(*PSReal); _efaf {
		_caa = _ggb.Push(MakeReal(_b.Floor(_gbbf.Val)))
	} else if _fbg, _cgd := _ged.(*PSInteger); _cgd {
		_caa = _ggb.Push(MakeInteger(_fbg.Val))
	} else {
		return ErrTypeCheck
	}
	return _caa
}
func (_gad *PSOperand) dup(_fbf *PSStack) error {
	_dce, _dgge := _fbf.Pop()
	if _dgge != nil {
		return _dgge
	}
	_dgge = _fbf.Push(_dce)
	if _dgge != nil {
		return _dgge
	}
	_dgge = _fbf.Push(_dce.Duplicate())
	return _dgge
}
func _bggb(_gca int) int {
	if _gca < 0 {
		return -_gca
	}
	return _gca
}

// PSParser is a basic Postscript parser.
type PSParser struct{ _bebfa *_ad.Reader }

func (_egcc *PSOperand) ne(_deb *PSStack) error {
	_aeacfe := _egcc.eq(_deb)
	if _aeacfe != nil {
		return _aeacfe
	}
	_aeacfe = _egcc.not(_deb)
	return _aeacfe
}
func (_fbdb *PSOperand) xor(_ffe *PSStack) error {
	_gfca, _fbcg := _ffe.Pop()
	if _fbcg != nil {
		return _fbcg
	}
	_ffd, _fbcg := _ffe.Pop()
	if _fbcg != nil {
		return _fbcg
	}
	if _cfdb, _eddg := _gfca.(*PSBoolean); _eddg {
		_cacb, _ebde := _ffd.(*PSBoolean)
		if !_ebde {
			return ErrTypeCheck
		}
		_fbcg = _ffe.Push(MakeBool(_cfdb.Val != _cacb.Val))
		return _fbcg
	}
	if _afaa, _caee := _gfca.(*PSInteger); _caee {
		_gebg, _dafa := _ffd.(*PSInteger)
		if !_dafa {
			return ErrTypeCheck
		}
		_fbcg = _ffe.Push(MakeInteger(_afaa.Val ^ _gebg.Val))
		return _fbcg
	}
	return ErrTypeCheck
}

// MakeOperand returns a new PSOperand object based on string `val`.
func MakeOperand(val string) *PSOperand { _gfbg := PSOperand(val); return &_gfbg }

// PSProgram defines a Postscript program which is a series of PS objects (arguments, commands, programs etc).
type PSProgram []PSObject

func (_cgea *PSOperand) neg(_ddbb *PSStack) error {
	_cec, _gfgf := _ddbb.Pop()
	if _gfgf != nil {
		return _gfgf
	}
	if _efc, _gdg := _cec.(*PSReal); _gdg {
		_gfgf = _ddbb.Push(MakeReal(-_efc.Val))
		return _gfgf
	} else if _cbbdg, _caea := _cec.(*PSInteger); _caea {
		_gfgf = _ddbb.Push(MakeInteger(-_cbbdg.Val))
		return _gfgf
	} else {
		return ErrTypeCheck
	}
}
func (_fc *PSBoolean) String() string { return _d.Sprintf("\u0025\u0076", _fc.Val) }
func (_gceaf *PSOperand) pop(_dgb *PSStack) error {
	_, _gaab := _dgb.Pop()
	if _gaab != nil {
		return _gaab
	}
	return nil
}
func (_acd *PSOperand) mul(_eddf *PSStack) error {
	_eee, _fdec := _eddf.Pop()
	if _fdec != nil {
		return _fdec
	}
	_agg, _fdec := _eddf.Pop()
	if _fdec != nil {
		return _fdec
	}
	_ade, _cafe := _eee.(*PSReal)
	_abee, _afea := _eee.(*PSInteger)
	if !_cafe && !_afea {
		return ErrTypeCheck
	}
	_egc, _aff := _agg.(*PSReal)
	_gcea, _dag := _agg.(*PSInteger)
	if !_aff && !_dag {
		return ErrTypeCheck
	}
	if _afea && _dag {
		_bdf := _abee.Val * _gcea.Val
		_agc := _eddf.Push(MakeInteger(_bdf))
		return _agc
	}
	var _bfa float64
	if _cafe {
		_bfa = _ade.Val
	} else {
		_bfa = float64(_abee.Val)
	}
	if _aff {
		_bfa *= _egc.Val
	} else {
		_bfa *= float64(_gcea.Val)
	}
	_fdec = _eddf.Push(MakeReal(_bfa))
	return _fdec
}
func (_gbbfd *PSOperand) log(_efb *PSStack) error {
	_gfb, _cbbd := _efb.PopNumberAsFloat64()
	if _cbbd != nil {
		return _cbbd
	}
	_caag := _b.Log10(_gfb)
	_cbbd = _efb.Push(MakeReal(_caag))
	return _cbbd
}

// NewPSParser returns a new instance of the PDF Postscript parser from input data.
func NewPSParser(content []byte) *PSParser {
	_bff := PSParser{}
	_bcff := _ga.NewBuffer(content)
	_bff._bebfa = _ad.NewReader(_bcff)
	return &_bff
}

// Parse parses the postscript and store as a program that can be executed.
func (_ecgg *PSParser) Parse() (*PSProgram, error) {
	_ecgg.skipSpaces()
	_ccaa, _fga := _ecgg._bebfa.Peek(2)
	if _fga != nil {
		return nil, _fga
	}
	if _ccaa[0] != '{' {
		return nil, _a.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0053\u0020\u0050\u0072\u006f\u0067\u0072\u0061\u006d\u0020\u006e\u006f\u0074\u0020\u0073t\u0061\u0072\u0074\u0069\u006eg\u0020\u0077i\u0074\u0068\u0020\u007b")
	}
	_eagd, _fga := _ecgg.parseFunction()
	if _fga != nil && _fga != _e.EOF {
		return nil, _fga
	}
	return _eagd, _fga
}

var ErrStackOverflow = _a.New("\u0073\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065r\u0066\u006c\u006f\u0077")

func (_ebe *PSOperand) String() string { return string(*_ebe) }
func (_ccc *PSOperand) eq(_daa *PSStack) error {
	_edd, _fccd := _daa.Pop()
	if _fccd != nil {
		return _fccd
	}
	_fcca, _fccd := _daa.Pop()
	if _fccd != nil {
		return _fccd
	}
	_fgd, _afa := _edd.(*PSBoolean)
	_bga, _ace := _fcca.(*PSBoolean)
	if _afa || _ace {
		var _ddg error
		if _afa && _ace {
			_ddg = _daa.Push(MakeBool(_fgd.Val == _bga.Val))
		} else {
			_ddg = _daa.Push(MakeBool(false))
		}
		return _ddg
	}
	var _ecc float64
	var _bgae float64
	if _gbb, _caf := _edd.(*PSInteger); _caf {
		_ecc = float64(_gbb.Val)
	} else if _bbab, _bege := _edd.(*PSReal); _bege {
		_ecc = _bbab.Val
	} else {
		return ErrTypeCheck
	}
	if _cafc, _bgc := _fcca.(*PSInteger); _bgc {
		_bgae = float64(_cafc.Val)
	} else if _cfd, _eff := _fcca.(*PSReal); _eff {
		_bgae = _cfd.Val
	} else {
		return ErrTypeCheck
	}
	if _b.Abs(_bgae-_ecc) < _aac {
		_fccd = _daa.Push(MakeBool(true))
	} else {
		_fccd = _daa.Push(MakeBool(false))
	}
	return _fccd
}

// NewPSStack returns an initialized PSStack.
func NewPSStack() *PSStack { return &PSStack{} }
func (_aed *PSOperand) idiv(_deeg *PSStack) error {
	_fde, _gcg := _deeg.Pop()
	if _gcg != nil {
		return _gcg
	}
	_adc, _gcg := _deeg.Pop()
	if _gcg != nil {
		return _gcg
	}
	_bee, _cee := _fde.(*PSInteger)
	if !_cee {
		return ErrTypeCheck
	}
	if _bee.Val == 0 {
		return ErrUndefinedResult
	}
	_fffa, _cee := _adc.(*PSInteger)
	if !_cee {
		return ErrTypeCheck
	}
	_adf := _fffa.Val / _bee.Val
	_gcg = _deeg.Push(MakeInteger(_adf))
	return _gcg
}
func (_bf *PSInteger) String() string { return _d.Sprintf("\u0025\u0064", _bf.Val) }

// PSReal represents a real number.
type PSReal struct{ Val float64 }

func (_aea *PSBoolean) Duplicate() PSObject { _fcb := PSBoolean{}; _fcb.Val = _aea.Val; return &_fcb }

// PopNumberAsFloat64 pops and return the numeric value of the top of the stack as a float64.
// Real or integer only.
func (_bebfg *PSStack) PopNumberAsFloat64() (float64, error) {
	_bcd, _gcdb := _bebfg.Pop()
	if _gcdb != nil {
		return 0, _gcdb
	}
	if _cbd, _babf := _bcd.(*PSReal); _babf {
		return _cbd.Val, nil
	} else if _aedeb, _bcgf := _bcd.(*PSInteger); _bcgf {
		return float64(_aedeb.Val), nil
	} else {
		return 0, ErrTypeCheck
	}
}

// Append appends an object to the PSProgram.
func (_ab *PSProgram) Append(obj PSObject) { *_ab = append(*_ab, obj) }
func (_cf *PSProgram) DebugString() string {
	_eb := "\u007b\u0020"
	for _, _ef := range *_cf {
		_eb += _ef.DebugString()
		_eb += "\u0020"
	}
	_eb += "\u007d"
	return _eb
}
func (_abb *PSParser) parseOperand() (*PSOperand, error) {
	var _aaae []byte
	for {
		_ddbe, _aecc := _abb._bebfa.Peek(1)
		if _aecc != nil {
			if _aecc == _e.EOF {
				break
			}
			return nil, _aecc
		}
		if _be.IsDelimiter(_ddbe[0]) {
			break
		}
		if _be.IsWhiteSpace(_ddbe[0]) {
			break
		}
		_bafe, _ := _abb._bebfa.ReadByte()
		_aaae = append(_aaae, _bafe)
	}
	if len(_aaae) == 0 {
		return nil, _a.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029")
	}
	return MakeOperand(string(_aaae)), nil
}
func (_eagf *PSOperand) sqrt(_bad *PSStack) error {
	_bgbb, _fgdd := _bad.PopNumberAsFloat64()
	if _fgdd != nil {
		return _fgdd
	}
	if _bgbb < 0 {
		return ErrRangeCheck
	}
	_aagb := _b.Sqrt(_bgbb)
	_fgdd = _bad.Push(MakeReal(_aagb))
	return _fgdd
}

var ErrStackUnderflow = _a.New("\u0073t\u0061c\u006b\u0020\u0075\u006e\u0064\u0065\u0072\u0066\u006c\u006f\u0077")

// DebugString returns a descriptive string representation of the stack - intended for debugging.
func (_aae *PSStack) DebugString() string {
	_bgce := "\u005b\u0020"
	for _, _caed := range *_aae {
		_bgce += _caed.DebugString()
		_bgce += "\u0020"
	}
	_bgce += "\u005d"
	return _bgce
}
func (_dff *PSOperand) round(_eegg *PSStack) error {
	_bbcb, _edgcb := _eegg.Pop()
	if _edgcb != nil {
		return _edgcb
	}
	if _aefg, _cef := _bbcb.(*PSReal); _cef {
		_edgcb = _eegg.Push(MakeReal(_b.Floor(_aefg.Val + 0.5)))
	} else if _ffbe, _cbbdgd := _bbcb.(*PSInteger); _cbbdgd {
		_edgcb = _eegg.Push(MakeInteger(_ffbe.Val))
	} else {
		return ErrTypeCheck
	}
	return _edgcb
}

// Execute executes the program for an input parameters `objects` and returns a slice of output objects.
func (_cg *PSExecutor) Execute(objects []PSObject) ([]PSObject, error) {
	for _, _bc := range objects {
		_daf := _cg.Stack.Push(_bc)
		if _daf != nil {
			return nil, _daf
		}
	}
	_ge := _cg._f.Exec(_cg.Stack)
	if _ge != nil {
		_aa.Log.Debug("\u0045x\u0065c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ge)
		return nil, _ge
	}
	_gf := []PSObject(*_cg.Stack)
	_cg.Stack.Empty()
	return _gf, nil
}
func (_befc *PSOperand) ln(_bea *PSStack) error {
	_fffg, _accd := _bea.PopNumberAsFloat64()
	if _accd != nil {
		return _accd
	}
	_beec := _b.Log(_fffg)
	_accd = _bea.Push(MakeReal(_beec))
	return _accd
}
func (_effb *PSOperand) le(_eecf *PSStack) error {
	_acc, _ggd := _eecf.PopNumberAsFloat64()
	if _ggd != nil {
		return _ggd
	}
	_eab, _ggd := _eecf.PopNumberAsFloat64()
	if _ggd != nil {
		return _ggd
	}
	if _b.Abs(_eab-_acc) < _aac {
		_ecg := _eecf.Push(MakeBool(true))
		return _ecg
	} else if _eab < _acc {
		_aeed := _eecf.Push(MakeBool(true))
		return _aeed
	} else {
		_aaag := _eecf.Push(MakeBool(false))
		return _aaag
	}
}
func (_ggbg *PSOperand) not(_fab *PSStack) error {
	_fgg, _bcg := _fab.Pop()
	if _bcg != nil {
		return _bcg
	}
	if _gfdc, _cea := _fgg.(*PSBoolean); _cea {
		_bcg = _fab.Push(MakeBool(!_gfdc.Val))
		return _bcg
	} else if _aec, _ccb := _fgg.(*PSInteger); _ccb {
		_bcg = _fab.Push(MakeInteger(^_aec.Val))
		return _bcg
	} else {
		return ErrTypeCheck
	}
}
func (_fcbf *PSParser) parseFunction() (*PSProgram, error) {
	_ddae, _ := _fcbf._bebfa.ReadByte()
	if _ddae != '{' {
		return nil, _a.New("\u0069\u006ev\u0061\u006c\u0069d\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	}
	_bfg := NewPSProgram()
	for {
		_fcbf.skipSpaces()
		_cefc, _cabe := _fcbf._bebfa.Peek(2)
		if _cabe != nil {
			if _cabe == _e.EOF {
				break
			}
			return nil, _cabe
		}
		_aa.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_cefc))
		if _cefc[0] == '}' {
			_aa.Log.Trace("\u0045\u004f\u0046 \u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
			_fcbf._bebfa.ReadByte()
			break
		} else if _cefc[0] == '{' {
			_aa.Log.Trace("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e!")
			_gea, _acg := _fcbf.parseFunction()
			if _acg != nil {
				return nil, _acg
			}
			_bfg.Append(_gea)
		} else if _be.IsDecimalDigit(_cefc[0]) || (_cefc[0] == '-' && _be.IsDecimalDigit(_cefc[1])) {
			_aa.Log.Trace("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!")
			_gef, _geaf := _fcbf.parseNumber()
			if _geaf != nil {
				return nil, _geaf
			}
			_bfg.Append(_gef)
		} else {
			_aa.Log.Trace("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f")
			_cefc, _ = _fcbf._bebfa.Peek(5)
			_dbgd := string(_cefc)
			_aa.Log.Trace("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073", _dbgd)
			if (len(_dbgd) > 4) && (_dbgd[:5] == "\u0066\u0061\u006cs\u0065") {
				_ggac, _gdb := _fcbf.parseBool()
				if _gdb != nil {
					return nil, _gdb
				}
				_bfg.Append(_ggac)
			} else if (len(_dbgd) > 3) && (_dbgd[:4] == "\u0074\u0072\u0075\u0065") {
				_ddbd, _geg := _fcbf.parseBool()
				if _geg != nil {
					return nil, _geg
				}
				_bfg.Append(_ddbd)
			} else {
				_ggaa, _ebag := _fcbf.parseOperand()
				if _ebag != nil {
					return nil, _ebag
				}
				_bfg.Append(_ggaa)
			}
		}
	}
	return _bfg, nil
}

const _aac = 0.000001

func (_ege *PSParser) skipSpaces() (int, error) {
	_eceg := 0
	for {
		_fgad, _feb := _ege._bebfa.Peek(1)
		if _feb != nil {
			return 0, _feb
		}
		if _be.IsWhiteSpace(_fgad[0]) {
			_ege._bebfa.ReadByte()
			_eceg++
		} else {
			break
		}
	}
	return _eceg, nil
}

// PSStack defines a stack of PSObjects. PSObjects can be pushed on or pull from the stack.
type PSStack []PSObject

func (_bd *PSInteger) DebugString() string {
	return _d.Sprintf("\u0069\u006e\u0074\u003a\u0025\u0064", _bd.Val)
}
func (_gbe *PSOperand) lt(_begg *PSStack) error {
	_aag, _cgbe := _begg.PopNumberAsFloat64()
	if _cgbe != nil {
		return _cgbe
	}
	_defd, _cgbe := _begg.PopNumberAsFloat64()
	if _cgbe != nil {
		return _cgbe
	}
	if _b.Abs(_defd-_aag) < _aac {
		_caagg := _begg.Push(MakeBool(false))
		return _caagg
	} else if _defd < _aag {
		_beea := _begg.Push(MakeBool(true))
		return _beea
	} else {
		_dba := _begg.Push(MakeBool(false))
		return _dba
	}
}
func (_feg *PSOperand) and(_gfc *PSStack) error {
	_ggc, _cab := _gfc.Pop()
	if _cab != nil {
		return _cab
	}
	_ecd, _cab := _gfc.Pop()
	if _cab != nil {
		return _cab
	}
	if _eba, _bac := _ggc.(*PSBoolean); _bac {
		_gd, _cdg := _ecd.(*PSBoolean)
		if !_cdg {
			return ErrTypeCheck
		}
		_cab = _gfc.Push(MakeBool(_eba.Val && _gd.Val))
		return _cab
	}
	if _edg, _db := _ggc.(*PSInteger); _db {
		_dcg, _afe := _ecd.(*PSInteger)
		if !_afe {
			return ErrTypeCheck
		}
		_cab = _gfc.Push(MakeInteger(_edg.Val & _dcg.Val))
		return _cab
	}
	return ErrTypeCheck
}
func (_eec *PSOperand) ifelse(_ebg *PSStack) error {
	_fcf, _bgf := _ebg.Pop()
	if _bgf != nil {
		return _bgf
	}
	_bab, _bgf := _ebg.Pop()
	if _bgf != nil {
		return _bgf
	}
	_dccc, _bgf := _ebg.Pop()
	if _bgf != nil {
		return _bgf
	}
	_cdc, _cfg := _fcf.(*PSProgram)
	if !_cfg {
		return ErrTypeCheck
	}
	_geb, _cfg := _bab.(*PSProgram)
	if !_cfg {
		return ErrTypeCheck
	}
	_aefb, _cfg := _dccc.(*PSBoolean)
	if !_cfg {
		return ErrTypeCheck
	}
	if _aefb.Val {
		_fbd := _geb.Exec(_ebg)
		return _fbd
	}
	_bgf = _cdc.Exec(_ebg)
	return _bgf
}

var ErrUndefinedResult = _a.New("\u0075\u006e\u0064\u0065fi\u006e\u0065\u0064\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0065\u0072\u0072o\u0072")

func (_ec *PSBoolean) DebugString() string {
	return _d.Sprintf("\u0062o\u006f\u006c\u003a\u0025\u0076", _ec.Val)
}
func (_bdg *PSOperand) gt(_fec *PSStack) error {
	_dbg, _afef := _fec.PopNumberAsFloat64()
	if _afef != nil {
		return _afef
	}
	_bbc, _afef := _fec.PopNumberAsFloat64()
	if _afef != nil {
		return _afef
	}
	if _b.Abs(_bbc-_dbg) < _aac {
		_bgd := _fec.Push(MakeBool(false))
		return _bgd
	} else if _bbc > _dbg {
		_fgb := _fec.Push(MakeBool(true))
		return _fgb
	} else {
		_gbc := _fec.Push(MakeBool(false))
		return _gbc
	}
}

// NewPSExecutor returns an initialized PSExecutor for an input `program`.
func NewPSExecutor(program *PSProgram) *PSExecutor {
	_ea := &PSExecutor{}
	_ea.Stack = NewPSStack()
	_ea._f = program
	return _ea
}
func (_ada *PSReal) String() string { return _d.Sprintf("\u0025\u002e\u0035\u0066", _ada.Val) }
func (_gdf *PSOperand) ceiling(_cae *PSStack) error {
	_gcd, _aaa := _cae.Pop()
	if _aaa != nil {
		return _aaa
	}
	if _cdd, _afc := _gcd.(*PSReal); _afc {
		_aaa = _cae.Push(MakeReal(_b.Ceil(_cdd.Val)))
	} else if _bfd, _ebd := _gcd.(*PSInteger); _ebd {
		_aaa = _cae.Push(MakeInteger(_bfd.Val))
	} else {
		_aaa = ErrTypeCheck
	}
	return _aaa
}
func (_gab *PSOperand) roll(_ffbg *PSStack) error {
	_bgec, _cfc := _ffbg.Pop()
	if _cfc != nil {
		return _cfc
	}
	_ffg, _cfc := _ffbg.Pop()
	if _cfc != nil {
		return _cfc
	}
	_ggae, _eade := _bgec.(*PSInteger)
	if !_eade {
		return ErrTypeCheck
	}
	_decd, _eade := _ffg.(*PSInteger)
	if !_eade {
		return ErrTypeCheck
	}
	if _decd.Val < 0 {
		return ErrRangeCheck
	}
	if _decd.Val == 0 || _decd.Val == 1 {
		return nil
	}
	if _decd.Val > len(*_ffbg) {
		return ErrStackUnderflow
	}
	for _fef := 0; _fef < _bggb(_ggae.Val); _fef++ {
		var _bfcc []PSObject
		_bfcc = (*_ffbg)[len(*_ffbg)-(_decd.Val) : len(*_ffbg)]
		if _ggae.Val > 0 {
			_eed := _bfcc[len(_bfcc)-1]
			_bfcc = append([]PSObject{_eed}, _bfcc[0:len(_bfcc)-1]...)
		} else {
			_fddc := _bfcc[len(_bfcc)-_decd.Val]
			_bfcc = append(_bfcc[1:], _fddc)
		}
		_bdeb := append((*_ffbg)[0:len(*_ffbg)-_decd.Val], _bfcc...)
		_ffbg = &_bdeb
	}
	return nil
}
func (_fddd *PSOperand) truncate(_fddb *PSStack) error {
	_adad, _aede := _fddb.Pop()
	if _aede != nil {
		return _aede
	}
	if _gcee, _gfff := _adad.(*PSReal); _gfff {
		_gfbc := int(_gcee.Val)
		_aede = _fddb.Push(MakeReal(float64(_gfbc)))
	} else if _beeac, _edf := _adad.(*PSInteger); _edf {
		_aede = _fddb.Push(MakeInteger(_beeac.Val))
	} else {
		return ErrTypeCheck
	}
	return _aede
}
func (_ce *PSProgram) Duplicate() PSObject {
	_aeg := &PSProgram{}
	for _, _bebf := range *_ce {
		_aeg.Append(_bebf.Duplicate())
	}
	return _aeg
}
func (_ed *PSInteger) Duplicate() PSObject { _bb := PSInteger{}; _bb.Val = _ed.Val; return &_bb }

// Exec executes the program, typically leaving output values on the stack.
func (_fcg *PSProgram) Exec(stack *PSStack) error {
	for _, _aeac := range *_fcg {
		var _cgc error
		switch _eaf := _aeac.(type) {
		case *PSInteger:
			_ff := _eaf
			_cgc = stack.Push(_ff)
		case *PSReal:
			_gec := _eaf
			_cgc = stack.Push(_gec)
		case *PSBoolean:
			_aead := _eaf
			_cgc = stack.Push(_aead)
		case *PSProgram:
			_bfc := _eaf
			_cgc = stack.Push(_bfc)
		case *PSOperand:
			_de := _eaf
			_cgc = _de.Exec(stack)
		default:
			return ErrTypeCheck
		}
		if _cgc != nil {
			return _cgc
		}
	}
	return nil
}
func (_fbgf *PSOperand) sin(_caac *PSStack) error {
	_cfab, _gffg := _caac.PopNumberAsFloat64()
	if _gffg != nil {
		return _gffg
	}
	_fdee := _b.Sin(_cfab * _b.Pi / 180.0)
	_gffg = _caac.Push(MakeReal(_fdee))
	return _gffg
}
func (_cddb *PSOperand) cvr(_fff *PSStack) error {
	_eef, _dgg := _fff.Pop()
	if _dgg != nil {
		return _dgg
	}
	if _cge, _edgc := _eef.(*PSReal); _edgc {
		_dgg = _fff.Push(MakeReal(_cge.Val))
	} else if _gee, _dee := _eef.(*PSInteger); _dee {
		_dgg = _fff.Push(MakeReal(float64(_gee.Val)))
	} else {
		return ErrTypeCheck
	}
	return _dgg
}

// PSInteger represents an integer.
type PSInteger struct{ Val int }

// PSBoolean represents a boolean value.
type PSBoolean struct{ Val bool }

// MakeReal returns a new PSReal object initialized with `val`.
func MakeReal(val float64) *PSReal { _adac := PSReal{}; _adac.Val = val; return &_adac }
