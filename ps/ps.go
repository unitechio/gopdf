package ps

import (
	_ggd "bufio"
	_gg "bytes"
	_d "errors"
	_eg "fmt"
	_g "io"
	_gf "math"

	_a "bitbucket.org/shenghui0779/gopdf/common"
	_af "bitbucket.org/shenghui0779/gopdf/core"
)

func (_agdf *PSOperand) gt(_bga *PSStack) error {
	_bac, _dad := _bga.PopNumberAsFloat64()
	if _dad != nil {
		return _dad
	}
	_fcg, _dad := _bga.PopNumberAsFloat64()
	if _dad != nil {
		return _dad
	}
	if _gf.Abs(_fcg-_bac) < _dc {
		_aed := _bga.Push(MakeBool(false))
		return _aed
	} else if _fcg > _bac {
		_abb := _bga.Push(MakeBool(true))
		return _abb
	} else {
		_eddd := _bga.Push(MakeBool(false))
		return _eddd
	}
}

// NewPSProgram returns an empty, initialized PSProgram.
func NewPSProgram() *PSProgram { return &PSProgram{} }

func (_bgee *PSOperand) xor(_agbc *PSStack) error {
	_dadc, _cafe := _agbc.Pop()
	if _cafe != nil {
		return _cafe
	}
	_fad, _cafe := _agbc.Pop()
	if _cafe != nil {
		return _cafe
	}
	if _ggca, _fbebf := _dadc.(*PSBoolean); _fbebf {
		_bgf, _ffc := _fad.(*PSBoolean)
		if !_ffc {
			return ErrTypeCheck
		}
		_cafe = _agbc.Push(MakeBool(_ggca.Val != _bgf.Val))
		return _cafe
	}
	if _acaf, _ageg := _dadc.(*PSInteger); _ageg {
		_bcdd, _egfd := _fad.(*PSInteger)
		if !_egfd {
			return ErrTypeCheck
		}
		_cafe = _agbc.Push(MakeInteger(_acaf.Val ^ _bcdd.Val))
		return _cafe
	}
	return ErrTypeCheck
}

// String returns a string representation of the stack.
func (_bfb *PSStack) String() string {
	_fdff := "\u005b\u0020"
	for _, _gcag := range *_bfb {
		_fdff += _gcag.String()
		_fdff += "\u0020"
	}
	_fdff += "\u005d"
	return _fdff
}

func (_bcd *PSOperand) ifelse(_eeb *PSStack) error {
	_dgc, _fee := _eeb.Pop()
	if _fee != nil {
		return _fee
	}
	_cdag, _fee := _eeb.Pop()
	if _fee != nil {
		return _fee
	}
	_edf, _fee := _eeb.Pop()
	if _fee != nil {
		return _fee
	}
	_bgef, _bcb := _dgc.(*PSProgram)
	if !_bcb {
		return ErrTypeCheck
	}
	_ecb, _bcb := _cdag.(*PSProgram)
	if !_bcb {
		return ErrTypeCheck
	}
	_fga, _bcb := _edf.(*PSBoolean)
	if !_bcb {
		return ErrTypeCheck
	}
	if _fga.Val {
		_ggcg := _ecb.Exec(_eeb)
		return _ggcg
	}
	_fee = _bgef.Exec(_eeb)
	return _fee
}
func (_fgf *PSReal) String() string { return _eg.Sprintf("\u0025\u002e\u0035\u0066", _fgf.Val) }
func (_fca *PSOperand) dup(_edae *PSStack) error {
	_acg, _gbfb := _edae.Pop()
	if _gbfb != nil {
		return _gbfb
	}
	_gbfb = _edae.Push(_acg)
	if _gbfb != nil {
		return _gbfb
	}
	_gbfb = _edae.Push(_acg.Duplicate())
	return _gbfb
}

var ErrStackUnderflow = _d.New("\u0073t\u0061c\u006b\u0020\u0075\u006e\u0064\u0065\u0072\u0066\u006c\u006f\u0077")

func (_dg *PSReal) Duplicate() PSObject { _aff := PSReal{}; _aff.Val = _dg.Val; return &_aff }
func (_ca *PSOperand) DebugString() string {
	return _eg.Sprintf("\u006fp\u003a\u0027\u0025\u0073\u0027", *_ca)
}

// PopInteger specificially pops an integer from the top of the stack, returning the value as an int.
func (_adf *PSStack) PopInteger() (int, error) {
	_dcbd, _bcbfg := _adf.Pop()
	if _bcbfg != nil {
		return 0, _bcbfg
	}
	if _bggd, _bee := _dcbd.(*PSInteger); _bee {
		return _bggd.Val, nil
	}
	return 0, ErrTypeCheck
}

func (_cgb *PSOperand) Duplicate() PSObject {
	_dfb := *_cgb
	return &_dfb
}

func (_ecff *PSOperand) log(_fba *PSStack) error {
	_eag, _abe := _fba.PopNumberAsFloat64()
	if _abe != nil {
		return _abe
	}
	_ggf := _gf.Log10(_eag)
	_abe = _fba.Push(MakeReal(_ggf))
	return _abe
}

func (_bgbe *PSOperand) le(_bfd *PSStack) error {
	_gae, _aeee := _bfd.PopNumberAsFloat64()
	if _aeee != nil {
		return _aeee
	}
	_eba, _aeee := _bfd.PopNumberAsFloat64()
	if _aeee != nil {
		return _aeee
	}
	if _gf.Abs(_eba-_gae) < _dc {
		_bacb := _bfd.Push(MakeBool(true))
		return _bacb
	} else if _eba < _gae {
		_ege := _bfd.Push(MakeBool(true))
		return _ege
	} else {
		_dbdf := _bfd.Push(MakeBool(false))
		return _dbdf
	}
}

const _dc = 0.000001

// PSOperand represents a Postscript operand (text string).
type PSOperand string

// PSExecutor has its own execution stack and is used to executre a PS routine (program).
type PSExecutor struct {
	Stack *PSStack
	_ed   *PSProgram
}

func (_bfeb *PSOperand) bitshift(_fef *PSStack) error {
	_dag, _cc := _fef.PopInteger()
	if _cc != nil {
		return _cc
	}
	_ggg, _cc := _fef.PopInteger()
	if _cc != nil {
		return _cc
	}
	var _dgde int
	if _dag >= 0 {
		_dgde = _ggg << uint(_dag)
	} else {
		_dgde = _ggg >> uint(-_dag)
	}
	_cc = _fef.Push(MakeInteger(_dgde))
	return _cc
}

// Parse parses the postscript and store as a program that can be executed.
func (_daee *PSParser) Parse() (*PSProgram, error) {
	_daee.skipSpaces()
	_gbeg, _ffe := _daee._gbea.Peek(2)
	if _ffe != nil {
		return nil, _ffe
	}
	if _gbeg[0] != '{' {
		return nil, _d.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0053\u0020\u0050\u0072\u006f\u0067\u0072\u0061\u006d\u0020\u006e\u006f\u0074\u0020\u0073t\u0061\u0072\u0074\u0069\u006eg\u0020\u0077i\u0074\u0068\u0020\u007b")
	}
	_ccde, _ffe := _daee.parseFunction()
	if _ffe != nil && _ffe != _g.EOF {
		return nil, _ffe
	}
	return _ccde, _ffe
}

func (_afb *PSOperand) cvr(_bcf *PSStack) error {
	_ddg, _dca := _bcf.Pop()
	if _dca != nil {
		return _dca
	}
	if _ddc, _feae := _ddg.(*PSReal); _feae {
		_dca = _bcf.Push(MakeReal(_ddc.Val))
	} else if _cdbf, _eda := _ddg.(*PSInteger); _eda {
		_dca = _bcf.Push(MakeReal(float64(_cdbf.Val)))
	} else {
		return ErrTypeCheck
	}
	return _dca
}
func (_bgd *PSOperand) String() string { return string(*_bgd) }
func (_cgbd *PSOperand) or(_dae *PSStack) error {
	_gfga, _bea := _dae.Pop()
	if _bea != nil {
		return _bea
	}
	_daef, _bea := _dae.Pop()
	if _bea != nil {
		return _bea
	}
	if _ace, _fgae := _gfga.(*PSBoolean); _fgae {
		_afe, _dgdd := _daef.(*PSBoolean)
		if !_dgdd {
			return ErrTypeCheck
		}
		_bea = _dae.Push(MakeBool(_ace.Val || _afe.Val))
		return _bea
	}
	if _geda, _ggdb := _gfga.(*PSInteger); _ggdb {
		_bag, _fcef := _daef.(*PSInteger)
		if !_fcef {
			return ErrTypeCheck
		}
		_bea = _dae.Push(MakeInteger(_geda.Val | _bag.Val))
		return _bea
	}
	return ErrTypeCheck
}

func (_geea *PSOperand) ln(_cdbb *PSStack) error {
	_aefe, _gge := _cdbb.PopNumberAsFloat64()
	if _gge != nil {
		return _gge
	}
	_fbf := _gf.Log(_aefe)
	_gge = _cdbb.Push(MakeReal(_fbf))
	return _gge
}

func (_acfa *PSOperand) lt(_afdg *PSStack) error {
	_gbab, _fbeb := _afdg.PopNumberAsFloat64()
	if _fbeb != nil {
		return _fbeb
	}
	_gcc, _fbeb := _afdg.PopNumberAsFloat64()
	if _fbeb != nil {
		return _fbeb
	}
	if _gf.Abs(_gcc-_gbab) < _dc {
		_ggdea := _afdg.Push(MakeBool(false))
		return _ggdea
	} else if _gcc < _gbab {
		_gabc := _afdg.Push(MakeBool(true))
		return _gabc
	} else {
		_ddcd := _afdg.Push(MakeBool(false))
		return _ddcd
	}
}

func (_ddbg *PSOperand) neg(_ccb *PSStack) error {
	_aaca, _ggcf := _ccb.Pop()
	if _ggcf != nil {
		return _ggcf
	}
	if _bcbd, _ebf := _aaca.(*PSReal); _ebf {
		_ggcf = _ccb.Push(MakeReal(-_bcbd.Val))
		return _ggcf
	} else if _ebcd, _dgg := _aaca.(*PSInteger); _dgg {
		_ggcf = _ccb.Push(MakeInteger(-_ebcd.Val))
		return _ggcf
	} else {
		return ErrTypeCheck
	}
}

// Exec executes the operand `op` in the state specified by `stack`.
func (_fb *PSOperand) Exec(stack *PSStack) error {
	_fe := ErrUnsupportedOperand
	switch *_fb {
	case "\u0061\u0062\u0073":
		_fe = _fb.abs(stack)
	case "\u0061\u0064\u0064":
		_fe = _fb.add(stack)
	case "\u0061\u006e\u0064":
		_fe = _fb.and(stack)
	case "\u0061\u0074\u0061\u006e":
		_fe = _fb.atan(stack)
	case "\u0062\u0069\u0074\u0073\u0068\u0069\u0066\u0074":
		_fe = _fb.bitshift(stack)
	case "\u0063e\u0069\u006c\u0069\u006e\u0067":
		_fe = _fb.ceiling(stack)
	case "\u0063\u006f\u0070\u0079":
		_fe = _fb.copy(stack)
	case "\u0063\u006f\u0073":
		_fe = _fb.cos(stack)
	case "\u0063\u0076\u0069":
		_fe = _fb.cvi(stack)
	case "\u0063\u0076\u0072":
		_fe = _fb.cvr(stack)
	case "\u0064\u0069\u0076":
		_fe = _fb.div(stack)
	case "\u0064\u0075\u0070":
		_fe = _fb.dup(stack)
	case "\u0065\u0071":
		_fe = _fb.eq(stack)
	case "\u0065\u0078\u0063\u0068":
		_fe = _fb.exch(stack)
	case "\u0065\u0078\u0070":
		_fe = _fb.exp(stack)
	case "\u0066\u006c\u006fo\u0072":
		_fe = _fb.floor(stack)
	case "\u0067\u0065":
		_fe = _fb.ge(stack)
	case "\u0067\u0074":
		_fe = _fb.gt(stack)
	case "\u0069\u0064\u0069\u0076":
		_fe = _fb.idiv(stack)
	case "\u0069\u0066":
		_fe = _fb.ifCondition(stack)
	case "\u0069\u0066\u0065\u006c\u0073\u0065":
		_fe = _fb.ifelse(stack)
	case "\u0069\u006e\u0064e\u0078":
		_fe = _fb.index(stack)
	case "\u006c\u0065":
		_fe = _fb.le(stack)
	case "\u006c\u006f\u0067":
		_fe = _fb.log(stack)
	case "\u006c\u006e":
		_fe = _fb.ln(stack)
	case "\u006c\u0074":
		_fe = _fb.lt(stack)
	case "\u006d\u006f\u0064":
		_fe = _fb.mod(stack)
	case "\u006d\u0075\u006c":
		_fe = _fb.mul(stack)
	case "\u006e\u0065":
		_fe = _fb.ne(stack)
	case "\u006e\u0065\u0067":
		_fe = _fb.neg(stack)
	case "\u006e\u006f\u0074":
		_fe = _fb.not(stack)
	case "\u006f\u0072":
		_fe = _fb.or(stack)
	case "\u0070\u006f\u0070":
		_fe = _fb.pop(stack)
	case "\u0072\u006f\u0075n\u0064":
		_fe = _fb.round(stack)
	case "\u0072\u006f\u006c\u006c":
		_fe = _fb.roll(stack)
	case "\u0073\u0069\u006e":
		_fe = _fb.sin(stack)
	case "\u0073\u0071\u0072\u0074":
		_fe = _fb.sqrt(stack)
	case "\u0073\u0075\u0062":
		_fe = _fb.sub(stack)
	case "\u0074\u0072\u0075\u006e\u0063\u0061\u0074\u0065":
		_fe = _fb.truncate(stack)
	case "\u0078\u006f\u0072":
		_fe = _fb.xor(stack)
	}
	return _fe
}

func (_cdaa *PSOperand) idiv(_edg *PSStack) error {
	_dbgd, _dff := _edg.Pop()
	if _dff != nil {
		return _dff
	}
	_cbf, _dff := _edg.Pop()
	if _dff != nil {
		return _dff
	}
	_eeec, _dgdea := _dbgd.(*PSInteger)
	if !_dgdea {
		return ErrTypeCheck
	}
	if _eeec.Val == 0 {
		return ErrUndefinedResult
	}
	_gda, _dgdea := _cbf.(*PSInteger)
	if !_dgdea {
		return ErrTypeCheck
	}
	_gbaa := _gda.Val / _eeec.Val
	_dff = _edg.Push(MakeInteger(_gbaa))
	return _dff
}

func (_ee *PSReal) DebugString() string {
	return _eg.Sprintf("\u0072e\u0061\u006c\u003a\u0025\u002e\u0035f", _ee.Val)
}

func (_bge *PSOperand) exp(_egbd *PSStack) error {
	_bcaea, _dcc := _egbd.PopNumberAsFloat64()
	if _dcc != nil {
		return _dcc
	}
	_ddga, _dcc := _egbd.PopNumberAsFloat64()
	if _dcc != nil {
		return _dcc
	}
	if _gf.Abs(_bcaea) < 1 && _ddga < 0 {
		return ErrUndefinedResult
	}
	_cda := _gf.Pow(_ddga, _bcaea)
	_dcc = _egbd.Push(MakeReal(_cda))
	return _dcc
}

// PSObjectArrayToFloat64Array converts []PSObject into a []float64 array. Each PSObject must represent a number,
// otherwise a ErrTypeCheck error occurs.
func PSObjectArrayToFloat64Array(objects []PSObject) ([]float64, error) {
	var _da []float64
	for _, _b := range objects {
		if _ef, _c := _b.(*PSInteger); _c {
			_da = append(_da, float64(_ef.Val))
		} else if _ec, _bg := _b.(*PSReal); _bg {
			_da = append(_da, _ec.Val)
		} else {
			return nil, ErrTypeCheck
		}
	}
	return _da, nil
}

func (_ag *PSOperand) and(_gab *PSStack) error {
	_bgb, _fd := _gab.Pop()
	if _fd != nil {
		return _fd
	}
	_ded, _fd := _gab.Pop()
	if _fd != nil {
		return _fd
	}
	if _agb, _gfa := _bgb.(*PSBoolean); _gfa {
		_aae, _gbe := _ded.(*PSBoolean)
		if !_gbe {
			return ErrTypeCheck
		}
		_fd = _gab.Push(MakeBool(_agb.Val && _aae.Val))
		return _fd
	}
	if _bfe, _dec := _bgb.(*PSInteger); _dec {
		_ecf, _ce := _ded.(*PSInteger)
		if !_ce {
			return ErrTypeCheck
		}
		_fd = _gab.Push(MakeInteger(_bfe.Val & _ecf.Val))
		return _fd
	}
	return ErrTypeCheck
}

func (_dcda *PSOperand) ge(_dac *PSStack) error {
	_gaf, _cca := _dac.PopNumberAsFloat64()
	if _cca != nil {
		return _cca
	}
	_ecce, _cca := _dac.PopNumberAsFloat64()
	if _cca != nil {
		return _cca
	}
	if _gf.Abs(_ecce-_gaf) < _dc {
		_caed := _dac.Push(MakeBool(true))
		return _caed
	} else if _ecce > _gaf {
		_ggde := _dac.Push(MakeBool(true))
		return _ggde
	} else {
		_ccg := _dac.Push(MakeBool(false))
		return _ccg
	}
}

// MakeOperand returns a new PSOperand object based on string `val`.
func MakeOperand(val string) *PSOperand { _dcbc := PSOperand(val); return &_dcbc }

func (_cae *PSOperand) eq(_cce *PSStack) error {
	_aab, _edd := _cce.Pop()
	if _edd != nil {
		return _edd
	}
	_afd, _edd := _cce.Pop()
	if _edd != nil {
		return _edd
	}
	_fbee, _egf := _aab.(*PSBoolean)
	_agf, _bcae := _afd.(*PSBoolean)
	if _egf || _bcae {
		var _acf error
		if _egf && _bcae {
			_acf = _cce.Push(MakeBool(_fbee.Val == _agf.Val))
		} else {
			_acf = _cce.Push(MakeBool(false))
		}
		return _acf
	}
	var _aca float64
	var _age float64
	if _cfc, _egb := _aab.(*PSInteger); _egb {
		_aca = float64(_cfc.Val)
	} else if _fbb, _fgef := _aab.(*PSReal); _fgef {
		_aca = _fbb.Val
	} else {
		return ErrTypeCheck
	}
	if _dcae, _aged := _afd.(*PSInteger); _aged {
		_age = float64(_dcae.Val)
	} else if _abf, _bfa := _afd.(*PSReal); _bfa {
		_age = _abf.Val
	} else {
		return ErrTypeCheck
	}
	if _gf.Abs(_age-_aca) < _dc {
		_edd = _cce.Push(MakeBool(true))
	} else {
		_edd = _cce.Push(MakeBool(false))
	}
	return _edd
}
func (_eed *PSBoolean) Duplicate() PSObject { _db := PSBoolean{}; _db.Val = _eed.Val; return &_db }

var ErrStackOverflow = _d.New("\u0073\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065r\u0066\u006c\u006f\u0077")

func (_dbef *PSOperand) sub(_acbb *PSStack) error {
	_dgb, _dcb := _acbb.Pop()
	if _dcb != nil {
		return _dcb
	}
	_fae, _dcb := _acbb.Pop()
	if _dcb != nil {
		return _dcb
	}
	_gbbd, _cfcf := _dgb.(*PSReal)
	_dfg, _ddd := _dgb.(*PSInteger)
	if !_cfcf && !_ddd {
		return ErrTypeCheck
	}
	_bfdc, _acae := _fae.(*PSReal)
	_dcad, _dda := _fae.(*PSInteger)
	if !_acae && !_dda {
		return ErrTypeCheck
	}
	if _ddd && _dda {
		_dfge := _dcad.Val - _dfg.Val
		_ebg := _acbb.Push(MakeInteger(_dfge))
		return _ebg
	}
	var _beda float64 = 0
	if _acae {
		_beda = _bfdc.Val
	} else {
		_beda = float64(_dcad.Val)
	}
	if _cfcf {
		_beda -= _gbbd.Val
	} else {
		_beda -= float64(_dfg.Val)
	}
	_dcb = _acbb.Push(MakeReal(_beda))
	return _dcb
}

// Exec executes the program, typically leaving output values on the stack.
func (_aee *PSProgram) Exec(stack *PSStack) error {
	for _, _dbg := range *_aee {
		var _ga error
		switch _cd := _dbg.(type) {
		case *PSInteger:
			_ba := _cd
			_ga = stack.Push(_ba)
		case *PSReal:
			_fge := _cd
			_ga = stack.Push(_fge)
		case *PSBoolean:
			_bad := _cd
			_ga = stack.Push(_bad)
		case *PSProgram:
			_gbf := _cd
			_ga = stack.Push(_gbf)
		case *PSOperand:
			_eea := _cd
			_ga = _eea.Exec(stack)
		default:
			return ErrTypeCheck
		}
		if _ga != nil {
			return _ga
		}
	}
	return nil
}

func (_gbfc *PSParser) parseFunction() (*PSProgram, error) {
	_bgfg, _ := _gbfc._gbea.ReadByte()
	if _bgfg != '{' {
		return nil, _d.New("\u0069\u006ev\u0061\u006c\u0069d\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	}
	_dffb := NewPSProgram()
	for {
		_gbfc.skipSpaces()
		_gbfc.skipComments()
		_faed, _bgfd := _gbfc._gbea.Peek(2)
		if _bgfd != nil {
			if _bgfd == _g.EOF {
				break
			}
			return nil, _bgfd
		}
		_a.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_faed))
		if _faed[0] == '}' {
			_a.Log.Trace("\u0045\u004f\u0046 \u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
			_gbfc._gbea.ReadByte()
			break
		} else if _faed[0] == '{' {
			_a.Log.Trace("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e!")
			_dge, _bbef := _gbfc.parseFunction()
			if _bbef != nil {
				return nil, _bbef
			}
			_dffb.Append(_dge)
		} else if _af.IsDecimalDigit(_faed[0]) || (_faed[0] == '-' && _af.IsDecimalDigit(_faed[1])) {
			_a.Log.Trace("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!")
			_fdf, _gcd := _gbfc.parseNumber()
			if _gcd != nil {
				return nil, _gcd
			}
			_dffb.Append(_fdf)
		} else {
			_a.Log.Trace("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f")
			_faed, _ = _gbfc._gbea.Peek(5)
			_cecfd := string(_faed)
			_a.Log.Trace("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073", _cecfd)
			if (len(_cecfd) > 4) && (_cecfd[:5] == "\u0066\u0061\u006cs\u0065") {
				_fcbf, _bdeb := _gbfc.parseBool()
				if _bdeb != nil {
					return nil, _bdeb
				}
				_dffb.Append(_fcbf)
			} else if (len(_cecfd) > 3) && (_cecfd[:4] == "\u0074\u0072\u0075\u0065") {
				_feb, _gdg := _gbfc.parseBool()
				if _gdg != nil {
					return nil, _gdg
				}
				_dffb.Append(_feb)
			} else {
				_fadc, _gcdd := _gbfc.parseOperand()
				if _gcdd != nil {
					return nil, _gcdd
				}
				_dffb.Append(_fadc)
			}
		}
	}
	return _dffb, nil
}

func (_aada *PSParser) skipSpaces() (int, error) {
	_eeg := 0
	for {
		_cbge, _dce := _aada._gbea.Peek(1)
		if _dce != nil {
			return 0, _dce
		}
		if _af.IsWhiteSpace(_cbge[0]) {
			_aada._gbea.ReadByte()
			_eeg++
		} else {
			break
		}
	}
	return _eeg, nil
}

func (_cbgec *PSParser) parseNumber() (PSObject, error) {
	_faa, _aabc := _af.ParseNumber(_cbgec._gbea)
	if _aabc != nil {
		return nil, _aabc
	}
	switch _fffe := _faa.(type) {
	case *_af.PdfObjectFloat:
		return MakeReal(float64(*_fffe)), nil
	case *_af.PdfObjectInteger:
		return MakeInteger(int(*_fffe)), nil
	}
	return nil, _eg.Errorf("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _faa)
}

func (_dbc *PSOperand) not(_egfc *PSStack) error {
	_cdba, _badc := _egfc.Pop()
	if _badc != nil {
		return _badc
	}
	if _cea, _egg := _cdba.(*PSBoolean); _egg {
		_badc = _egfc.Push(MakeBool(!_cea.Val))
		return _badc
	} else if _ffa, _gff := _cdba.(*PSInteger); _gff {
		_badc = _egfc.Push(MakeInteger(^_ffa.Val))
		return _badc
	} else {
		return ErrTypeCheck
	}
}

// NewPSStack returns an initialized PSStack.
func NewPSStack() *PSStack { return &PSStack{} }

var (
	ErrTypeCheck  = _d.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	ErrRangeCheck = _d.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
)

func (_ebc *PSBoolean) String() string { return _eg.Sprintf("\u0025\u0076", _ebc.Val) }

var ErrUndefinedResult = _d.New("\u0075\u006e\u0064\u0065fi\u006e\u0065\u0064\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0065\u0072\u0072o\u0072")

// PSReal represents a real number.
type PSReal struct{ Val float64 }

func (_edb *PSOperand) round(_add *PSStack) error {
	_deg, _cbgd := _add.Pop()
	if _cbgd != nil {
		return _cbgd
	}
	if _cbb, _gbde := _deg.(*PSReal); _gbde {
		_cbgd = _add.Push(MakeReal(_gf.Floor(_cbb.Val + 0.5)))
	} else if _accb, _afaf := _deg.(*PSInteger); _afaf {
		_cbgd = _add.Push(MakeInteger(_accb.Val))
	} else {
		return ErrTypeCheck
	}
	return _cbgd
}

// Append appends an object to the PSProgram.
func (_gd *PSProgram) Append(obj PSObject) { *_gd = append(*_gd, obj) }

func (_eage *PSOperand) sin(_cgdf *PSStack) error {
	_fff, _cdf := _cgdf.PopNumberAsFloat64()
	if _cdf != nil {
		return _cdf
	}
	_ebcdf := _gf.Sin(_fff * _gf.Pi / 180.0)
	_cdf = _cgdf.Push(MakeReal(_ebcdf))
	return _cdf
}
func (_cg *PSInteger) Duplicate() PSObject { _dd := PSInteger{}; _dd.Val = _cg.Val; return &_dd }

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

func (_ae *PSInteger) String() string { return _eg.Sprintf("\u0025\u0064", _ae.Val) }
func (_bdb *PSOperand) atan(_agg *PSStack) error {
	_dbdc, _ac := _agg.PopNumberAsFloat64()
	if _ac != nil {
		return _ac
	}
	_acc, _ac := _agg.PopNumberAsFloat64()
	if _ac != nil {
		return _ac
	}
	if _dbdc == 0 {
		var _fac error
		if _acc < 0 {
			_fac = _agg.Push(MakeReal(270))
		} else {
			_fac = _agg.Push(MakeReal(90))
		}
		return _fac
	}
	_aga := _acc / _dbdc
	_ceb := _gf.Atan(_aga) * 180 / _gf.Pi
	_ac = _agg.Push(MakeReal(_ceb))
	return _ac
}

func (_dbf *PSOperand) floor(_cbc *PSStack) error {
	_bdbg, _ddf := _cbc.Pop()
	if _ddf != nil {
		return _ddf
	}
	if _gaga, _eace := _bdbg.(*PSReal); _eace {
		_ddf = _cbc.Push(MakeReal(_gf.Floor(_gaga.Val)))
	} else if _ggdd, _geg := _bdbg.(*PSInteger); _geg {
		_ddf = _cbc.Push(MakeInteger(_ggdd.Val))
	} else {
		return ErrTypeCheck
	}
	return _ddf
}

// PSStack defines a stack of PSObjects. PSObjects can be pushed on or pull from the stack.
type PSStack []PSObject

func (_dcaf *PSOperand) ifCondition(_dbde *PSStack) error {
	_gbd, _fbed := _dbde.Pop()
	if _fbed != nil {
		return _fbed
	}
	_ddb, _fbed := _dbde.Pop()
	if _fbed != nil {
		return _fbed
	}
	_fcac, _eaa := _gbd.(*PSProgram)
	if !_eaa {
		return ErrTypeCheck
	}
	_ggbb, _eaa := _ddb.(*PSBoolean)
	if !_eaa {
		return ErrTypeCheck
	}
	if _ggbb.Val {
		_gbg := _fcac.Exec(_dbde)
		return _gbg
	}
	return nil
}

func (_bbb *PSOperand) ne(_eeea *PSStack) error {
	_bfc := _bbb.eq(_eeea)
	if _bfc != nil {
		return _bfc
	}
	_bfc = _bbb.not(_eeea)
	return _bfc
}

// DebugString returns a descriptive string representation of the stack - intended for debugging.
func (_adacg *PSStack) DebugString() string {
	_dfcf := "\u005b\u0020"
	for _, _agfgc := range *_adacg {
		_dfcf += _agfgc.DebugString()
		_dfcf += "\u0020"
	}
	_dfcf += "\u005d"
	return _dfcf
}

func (_cff *PSOperand) roll(_aedc *PSStack) error {
	_dace, _cgbb := _aedc.Pop()
	if _cgbb != nil {
		return _cgbb
	}
	_bde, _cgbb := _aedc.Pop()
	if _cgbb != nil {
		return _cgbb
	}
	_fda, _gdb := _dace.(*PSInteger)
	if !_gdb {
		return ErrTypeCheck
	}
	_geed, _gdb := _bde.(*PSInteger)
	if !_gdb {
		return ErrTypeCheck
	}
	if _geed.Val < 0 {
		return ErrRangeCheck
	}
	if _geed.Val == 0 || _geed.Val == 1 {
		return nil
	}
	if _geed.Val > len(*_aedc) {
		return ErrStackUnderflow
	}
	for _dba := 0; _dba < _dffe(_fda.Val); _dba++ {
		var _eeab []PSObject
		_eeab = (*_aedc)[len(*_aedc)-(_geed.Val) : len(*_aedc)]
		if _fda.Val > 0 {
			_cgg := _eeab[len(_eeab)-1]
			_eeab = append([]PSObject{_cgg}, _eeab[0:len(_eeab)-1]...)
		} else {
			_eabc := _eeab[len(_eeab)-_geed.Val]
			_eeab = append(_eeab[1:], _eabc)
		}
		_bed := append((*_aedc)[0:len(*_aedc)-_geed.Val], _eeab...)
		_aedc = &_bed
	}
	return nil
}

func (_ada *PSOperand) mod(_fbfg *PSStack) error {
	_cga, _aac := _fbfg.Pop()
	if _aac != nil {
		return _aac
	}
	_cec, _aac := _fbfg.Pop()
	if _aac != nil {
		return _aac
	}
	_fcb, _gfad := _cga.(*PSInteger)
	if !_gfad {
		return ErrTypeCheck
	}
	if _fcb.Val == 0 {
		return ErrUndefinedResult
	}
	_bdcd, _gfad := _cec.(*PSInteger)
	if !_gfad {
		return ErrTypeCheck
	}
	_aeff := _bdcd.Val % _fcb.Val
	_aac = _fbfg.Push(MakeInteger(_aeff))
	return _aac
}

func (_bbc *PSOperand) sqrt(_gdd *PSStack) error {
	_gde, _cebb := _gdd.PopNumberAsFloat64()
	if _cebb != nil {
		return _cebb
	}
	if _gde < 0 {
		return ErrRangeCheck
	}
	_gbaea := _gf.Sqrt(_gde)
	_cebb = _gdd.Push(MakeReal(_gbaea))
	return _cebb
}

// PSBoolean represents a boolean value.
type PSBoolean struct{ Val bool }

// Push pushes an object on top of the stack.
func (_cfa *PSStack) Push(obj PSObject) error {
	if len(*_cfa) > 100 {
		return ErrStackOverflow
	}
	*_cfa = append(*_cfa, obj)
	return nil
}

func (_ea *PSOperand) abs(_cgc *PSStack) error {
	_cf, _bf := _cgc.Pop()
	if _bf != nil {
		return _bf
	}
	if _dfe, _ecd := _cf.(*PSReal); _ecd {
		_fbe := _dfe.Val
		if _fbe < 0 {
			_bf = _cgc.Push(MakeReal(-_fbe))
		} else {
			_bf = _cgc.Push(MakeReal(_fbe))
		}
	} else if _cac, _fa := _cf.(*PSInteger); _fa {
		_efe := _cac.Val
		if _efe < 0 {
			_bf = _cgc.Push(MakeInteger(-_efe))
		} else {
			_bf = _cgc.Push(MakeInteger(_efe))
		}
	} else {
		return ErrTypeCheck
	}
	return _bf
}

// PopNumberAsFloat64 pops and return the numeric value of the top of the stack as a float64.
// Real or integer only.
func (_ebga *PSStack) PopNumberAsFloat64() (float64, error) {
	_bddb, _dbba := _ebga.Pop()
	if _dbba != nil {
		return 0, _dbba
	}
	if _cdac, _gcdda := _bddb.(*PSReal); _gcdda {
		return _cdac.Val, nil
	} else if _gbcd, _bcbff := _bddb.(*PSInteger); _bcbff {
		return float64(_gbcd.Val), nil
	} else {
		return 0, ErrTypeCheck
	}
}

func (_faec *PSParser) parseOperand() (*PSOperand, error) {
	var _ebb []byte
	for {
		_ffcf, _cccg := _faec._gbea.Peek(1)
		if _cccg != nil {
			if _cccg == _g.EOF {
				break
			}
			return nil, _cccg
		}
		if _af.IsDelimiter(_ffcf[0]) {
			break
		}
		if _af.IsWhiteSpace(_ffcf[0]) {
			break
		}
		_dbb, _ := _faec._gbea.ReadByte()
		_ebb = append(_ebb, _dbb)
	}
	if len(_ebb) == 0 {
		return nil, _d.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029")
	}
	return MakeOperand(string(_ebb)), nil
}

func (_gce *PSParser) parseBool() (*PSBoolean, error) {
	_eaad, _eged := _gce._gbea.Peek(4)
	if _eged != nil {
		return MakeBool(false), _eged
	}
	if (len(_eaad) >= 4) && (string(_eaad[:4]) == "\u0074\u0072\u0075\u0065") {
		_gce._gbea.Discard(4)
		return MakeBool(true), nil
	}
	_eaad, _eged = _gce._gbea.Peek(5)
	if _eged != nil {
		return MakeBool(false), _eged
	}
	if (len(_eaad) >= 5) && (string(_eaad[:5]) == "\u0066\u0061\u006cs\u0065") {
		_gce._gbea.Discard(5)
		return MakeBool(false), nil
	}
	return MakeBool(false), _d.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

func (_dcaa *PSOperand) truncate(_dfa *PSStack) error {
	_dege, _eec := _dfa.Pop()
	if _eec != nil {
		return _eec
	}
	if _cab, _ggfa := _dege.(*PSReal); _ggfa {
		_gageb := int(_cab.Val)
		_eec = _dfa.Push(MakeReal(float64(_gageb)))
	} else if _aecg, _aecc := _dege.(*PSInteger); _aecc {
		_eec = _dfa.Push(MakeInteger(_aecg.Val))
	} else {
		return ErrTypeCheck
	}
	return _eec
}

func (_bd *PSInteger) DebugString() string {
	return _eg.Sprintf("\u0069\u006e\u0074\u003a\u0025\u0064", _bd.Val)
}

func (_dfc *PSProgram) String() string {
	_dgf := "\u007b\u0020"
	for _, _aef := range *_dfc {
		_dgf += _aef.String()
		_dgf += "\u0020"
	}
	_dgf += "\u007d"
	return _dgf
}

func (_cecf *PSOperand) mul(_gbc *PSStack) error {
	_bfag, _gbcg := _gbc.Pop()
	if _gbcg != nil {
		return _gbcg
	}
	_ged, _gbcg := _gbc.Pop()
	if _gbcg != nil {
		return _gbcg
	}
	_edge, _dagc := _bfag.(*PSReal)
	_dbfe, _aabg := _bfag.(*PSInteger)
	if !_dagc && !_aabg {
		return ErrTypeCheck
	}
	_aad, _ccea := _ged.(*PSReal)
	_gcf, _bgc := _ged.(*PSInteger)
	if !_ccea && !_bgc {
		return ErrTypeCheck
	}
	if _aabg && _bgc {
		_egbdb := _dbfe.Val * _gcf.Val
		_dadb := _gbc.Push(MakeInteger(_egbdb))
		return _dadb
	}
	var _bffa float64
	if _dagc {
		_bffa = _edge.Val
	} else {
		_bffa = float64(_dbfe.Val)
	}
	if _ccea {
		_bffa *= _aad.Val
	} else {
		_bffa *= float64(_gcf.Val)
	}
	_gbcg = _gbc.Push(MakeReal(_bffa))
	return _gbcg
}

var ErrUnsupportedOperand = _d.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")

// Pop pops an object from the top of the stack.
func (_agfg *PSStack) Pop() (PSObject, error) {
	if len(*_agfg) < 1 {
		return nil, ErrStackUnderflow
	}
	_dagcg := (*_agfg)[len(*_agfg)-1]
	*_agfg = (*_agfg)[0 : len(*_agfg)-1]
	return _dagcg, nil
}

func (_cdb *PSOperand) ceiling(_bbd *PSStack) error {
	_bdd, _cba := _bbd.Pop()
	if _cba != nil {
		return _cba
	}
	if _fgee, _aaf := _bdd.(*PSReal); _aaf {
		_cba = _bbd.Push(MakeReal(_gf.Ceil(_fgee.Val)))
	} else if _ccd, _ab := _bdd.(*PSInteger); _ab {
		_cba = _bbd.Push(MakeInteger(_ccd.Val))
	} else {
		_cba = ErrTypeCheck
	}
	return _cba
}

func (_fdd *PSOperand) index(_ccad *PSStack) error {
	_dcag, _bab := _ccad.Pop()
	if _bab != nil {
		return _bab
	}
	_ddge, _dedb := _dcag.(*PSInteger)
	if !_dedb {
		return ErrTypeCheck
	}
	if _ddge.Val < 0 {
		return ErrRangeCheck
	}
	if _ddge.Val > len(*_ccad)-1 {
		return ErrStackUnderflow
	}
	_bcc := (*_ccad)[len(*_ccad)-1-_ddge.Val]
	_bab = _ccad.Push(_bcc.Duplicate())
	return _bab
}

// Execute executes the program for an input parameters `objects` and returns a slice of output objects.
func (_f *PSExecutor) Execute(objects []PSObject) ([]PSObject, error) {
	for _, _fc := range objects {
		_efc := _f.Stack.Push(_fc)
		if _efc != nil {
			return nil, _efc
		}
	}
	_gb := _f._ed.Exec(_f.Stack)
	if _gb != nil {
		_a.Log.Debug("\u0045x\u0065c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _gb)
		return nil, _gb
	}
	_fg := []PSObject(*_f.Stack)
	_f.Stack.Empty()
	return _fg, nil
}

// MakeReal returns a new PSReal object initialized with `val`.
func MakeReal(val float64) *PSReal { _gbge := PSReal{}; _gbge.Val = val; return &_gbge }

func _dffe(_gbee int) int {
	if _gbee < 0 {
		return -_gbee
	}
	return _gbee
}

// MakeBool returns a new PSBoolean object initialized with `val`.
func MakeBool(val bool) *PSBoolean { _acgb := PSBoolean{}; _acgb.Val = val; return &_acgb }

// Empty empties the stack.
func (_dged *PSStack) Empty() { *_dged = []PSObject{} }

func (_gage *PSOperand) pop(_agge *PSStack) error {
	_, _bcbf := _agge.Pop()
	if _bcbf != nil {
		return _bcbf
	}
	return nil
}

// PSInteger represents an integer.
type PSInteger struct{ Val int }

// PSProgram defines a Postscript program which is a series of PS objects (arguments, commands, programs etc).
type PSProgram []PSObject

func (_ff *PSOperand) copy(_bca *PSStack) error {
	_aec, _bdcc := _bca.PopInteger()
	if _bdcc != nil {
		return _bdcc
	}
	if _aec < 0 {
		return ErrRangeCheck
	}
	if _aec > len(*_bca) {
		return ErrRangeCheck
	}
	*_bca = append(*_bca, (*_bca)[len(*_bca)-_aec:]...)
	return nil
}

func (_eab *PSOperand) cvi(_gag *PSStack) error {
	_cgd, _eca := _gag.Pop()
	if _eca != nil {
		return _eca
	}
	if _ecaf, _ggb := _cgd.(*PSReal); _ggb {
		_bbe := int(_ecaf.Val)
		_eca = _gag.Push(MakeInteger(_bbe))
	} else if _acb, _cdc := _cgd.(*PSInteger); _cdc {
		_gba := _acb.Val
		_eca = _gag.Push(MakeInteger(_gba))
	} else {
		return ErrTypeCheck
	}
	return _eca
}

func (_gbb *PSProgram) Duplicate() PSObject {
	_bdc := &PSProgram{}
	for _, _bc := range *_gbb {
		_bdc.Append(_bc.Duplicate())
	}
	return _bdc
}

func (_fcd *PSOperand) add(_dbd *PSStack) error {
	_cb, _de := _dbd.Pop()
	if _de != nil {
		return _de
	}
	_aa, _de := _dbd.Pop()
	if _de != nil {
		return _de
	}
	_bcg, _afa := _cb.(*PSReal)
	_fea, _gfgc := _cb.(*PSInteger)
	if !_afa && !_gfgc {
		return ErrTypeCheck
	}
	_eef, _egc := _aa.(*PSReal)
	_eeac, _dfbc := _aa.(*PSInteger)
	if !_egc && !_dfbc {
		return ErrTypeCheck
	}
	if _gfgc && _dfbc {
		_gee := _fea.Val + _eeac.Val
		_eac := _dbd.Push(MakeInteger(_gee))
		return _eac
	}
	var _gc float64
	if _afa {
		_gc = _bcg.Val
	} else {
		_gc = float64(_fea.Val)
	}
	if _egc {
		_gc += _eef.Val
	} else {
		_gc += float64(_eeac.Val)
	}
	_de = _dbd.Push(MakeReal(_gc))
	return _de
}

// MakeInteger returns a new PSInteger object initialized with `val`.
func MakeInteger(val int) *PSInteger { _beb := PSInteger{}; _beb.Val = val; return &_beb }

func (_ecc *PSProgram) DebugString() string {
	_gfg := "\u007b\u0020"
	for _, _df := range *_ecc {
		_gfg += _df.DebugString()
		_gfg += "\u0020"
	}
	_gfg += "\u007d"
	return _gfg
}

func (_fag *PSOperand) exch(_dbe *PSStack) error {
	_cee, _abg := _dbe.Pop()
	if _abg != nil {
		return _abg
	}
	_edc, _abg := _dbe.Pop()
	if _abg != nil {
		return _abg
	}
	_abg = _dbe.Push(_cee)
	if _abg != nil {
		return _abg
	}
	_abg = _dbe.Push(_edc)
	return _abg
}

func (_ecfg *PSParser) skipComments() error {
	if _, _edac := _ecfg.skipSpaces(); _edac != nil {
		return _edac
	}
	_cacd := true
	for {
		_ccc, _egbb := _ecfg._gbea.Peek(1)
		if _egbb != nil {
			_a.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _egbb.Error())
			return _egbb
		}
		if _cacd && _ccc[0] != '%' {
			return nil
		}
		_cacd = false
		if (_ccc[0] != '\r') && (_ccc[0] != '\n') {
			_ecfg._gbea.ReadByte()
		} else {
			break
		}
	}
	return _ecfg.skipComments()
}

func (_ggc *PSOperand) cos(_dcd *PSStack) error {
	_dfef, _caf := _dcd.PopNumberAsFloat64()
	if _caf != nil {
		return _caf
	}
	_bda := _gf.Cos(_dfef * _gf.Pi / 180.0)
	_caf = _dcd.Push(MakeReal(_bda))
	return _caf
}

func (_ge *PSBoolean) DebugString() string {
	return _eg.Sprintf("\u0062o\u006f\u006c\u003a\u0025\u0076", _ge.Val)
}

// NewPSExecutor returns an initialized PSExecutor for an input `program`.
func NewPSExecutor(program *PSProgram) *PSExecutor {
	_eb := &PSExecutor{}
	_eb.Stack = NewPSStack()
	_eb._ed = program
	return _eb
}

// NewPSParser returns a new instance of the PDF Postscript parser from input data.
func NewPSParser(content []byte) *PSParser {
	_bffb := PSParser{}
	_dbeg := _gg.NewBuffer(content)
	_bffb._gbea = _ggd.NewReader(_dbeg)
	return &_bffb
}

func (_fbec *PSOperand) div(_cbg *PSStack) error {
	_gbae, _ad := _cbg.Pop()
	if _ad != nil {
		return _ad
	}
	_be, _ad := _cbg.Pop()
	if _ad != nil {
		return _ad
	}
	_bdcb, _agd := _gbae.(*PSReal)
	_bff, _feac := _gbae.(*PSInteger)
	if !_agd && !_feac {
		return ErrTypeCheck
	}
	if _agd && _bdcb.Val == 0 {
		return ErrUndefinedResult
	}
	if _feac && _bff.Val == 0 {
		return ErrUndefinedResult
	}
	_feaef, _cbd := _be.(*PSReal)
	_ead, _feaa := _be.(*PSInteger)
	if !_cbd && !_feaa {
		return ErrTypeCheck
	}
	var _gca float64
	if _cbd {
		_gca = _feaef.Val
	} else {
		_gca = float64(_ead.Val)
	}
	if _agd {
		_gca /= _bdcb.Val
	} else {
		_gca /= float64(_bff.Val)
	}
	_ad = _cbg.Push(MakeReal(_gca))
	return _ad
}

// PSParser is a basic Postscript parser.
type PSParser struct{ _gbea *_ggd.Reader }
