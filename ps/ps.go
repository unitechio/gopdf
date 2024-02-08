package ps

import (
	_gba "bufio"
	_gb "bytes"
	_c "errors"
	_f "fmt"
	_a "io"
	_gbg "math"

	_e "bitbucket.org/shenghui0779/gopdf/common"
	_ca "bitbucket.org/shenghui0779/gopdf/core"
)

// PSExecutor has its own execution stack and is used to executre a PS routine (program).
type PSExecutor struct {
	Stack *PSStack
	_b    *PSProgram
}

func (_fcga *PSOperand) dup(_bgc *PSStack) error {
	_aeaf, _bfc := _bgc.Pop()
	if _bfc != nil {
		return _bfc
	}
	_bfc = _bgc.Push(_aeaf)
	if _bfc != nil {
		return _bfc
	}
	_bfc = _bgc.Push(_aeaf.Duplicate())
	return _bfc
}
func (_gae *PSOperand) div(_cca *PSStack) error {
	_gga, _bbb := _cca.Pop()
	if _bbb != nil {
		return _bbb
	}
	_gcad, _bbb := _cca.Pop()
	if _bbb != nil {
		return _bbb
	}
	_gdd, _ebb := _gga.(*PSReal)
	_ffgf, _dcb := _gga.(*PSInteger)
	if !_ebb && !_dcb {
		return ErrTypeCheck
	}
	if _ebb && _gdd.Val == 0 {
		return ErrUndefinedResult
	}
	if _dcb && _ffgf.Val == 0 {
		return ErrUndefinedResult
	}
	_dfc, _eced := _gcad.(*PSReal)
	_aaa, _agc := _gcad.(*PSInteger)
	if !_eced && !_agc {
		return ErrTypeCheck
	}
	var _dca float64
	if _eced {
		_dca = _dfc.Val
	} else {
		_dca = float64(_aaa.Val)
	}
	if _ebb {
		_dca /= _gdd.Val
	} else {
		_dca /= float64(_ffgf.Val)
	}
	_bbb = _cca.Push(MakeReal(_dca))
	return _bbb
}
func (_gcb *PSOperand) roll(_dbca *PSStack) error {
	_fef, _eece := _dbca.Pop()
	if _eece != nil {
		return _eece
	}
	_cfdcf, _eece := _dbca.Pop()
	if _eece != nil {
		return _eece
	}
	_ecgg, _abac := _fef.(*PSInteger)
	if !_abac {
		return ErrTypeCheck
	}
	_dedb, _abac := _cfdcf.(*PSInteger)
	if !_abac {
		return ErrTypeCheck
	}
	if _dedb.Val < 0 {
		return ErrRangeCheck
	}
	if _dedb.Val == 0 || _dedb.Val == 1 {
		return nil
	}
	if _dedb.Val > len(*_dbca) {
		return ErrStackUnderflow
	}
	for _gcf := 0; _gcf < _bbbe(_ecgg.Val); _gcf++ {
		var _cddf []PSObject
		_cddf = (*_dbca)[len(*_dbca)-(_dedb.Val) : len(*_dbca)]
		if _ecgg.Val > 0 {
			_gaa := _cddf[len(_cddf)-1]
			_cddf = append([]PSObject{_gaa}, _cddf[0:len(_cddf)-1]...)
		} else {
			_dcbae := _cddf[len(_cddf)-_dedb.Val]
			_cddf = append(_cddf[1:], _dcbae)
		}
		_cddg := append((*_dbca)[0:len(*_dbca)-_dedb.Val], _cddf...)
		_dbca = &_cddg
	}
	return nil
}
func (_afg *PSOperand) or(_gbbcc *PSStack) error {
	_edd, _gfa := _gbbcc.Pop()
	if _gfa != nil {
		return _gfa
	}
	_aedd, _gfa := _gbbcc.Pop()
	if _gfa != nil {
		return _gfa
	}
	if _feb, _baaa := _edd.(*PSBoolean); _baaa {
		_dbga, _cgb := _aedd.(*PSBoolean)
		if !_cgb {
			return ErrTypeCheck
		}
		_gfa = _gbbcc.Push(MakeBool(_feb.Val || _dbga.Val))
		return _gfa
	}
	if _cfbc, _acfe := _edd.(*PSInteger); _acfe {
		_cfg, _gead := _aedd.(*PSInteger)
		if !_gead {
			return ErrTypeCheck
		}
		_gfa = _gbbcc.Push(MakeInteger(_cfbc.Val | _cfg.Val))
		return _gfa
	}
	return ErrTypeCheck
}
func (_ac *PSBoolean) String() string { return _f.Sprintf("\u0025\u0076", _ac.Val) }
func (_ddgb *PSOperand) truncate(_aacc *PSStack) error {
	_cafc, _gbgf := _aacc.Pop()
	if _gbgf != nil {
		return _gbgf
	}
	if _dgec, _beg := _cafc.(*PSReal); _beg {
		_cbfg := int(_dgec.Val)
		_gbgf = _aacc.Push(MakeReal(float64(_cbfg)))
	} else if _ddd, _daea := _cafc.(*PSInteger); _daea {
		_gbgf = _aacc.Push(MakeInteger(_ddd.Val))
	} else {
		return ErrTypeCheck
	}
	return _gbgf
}

// NewPSExecutor returns an initialized PSExecutor for an input `program`.
func NewPSExecutor(program *PSProgram) *PSExecutor {
	_bf := &PSExecutor{}
	_bf.Stack = NewPSStack()
	_bf._b = program
	return _bf
}

// MakeBool returns a new PSBoolean object initialized with `val`.
func MakeBool(val bool) *PSBoolean         { _bbfc := PSBoolean{}; _bbfc.Val = val; return &_bbfc }
func (_df *PSInteger) Duplicate() PSObject { _cdd := PSInteger{}; _cdd.Val = _df.Val; return &_cdd }

// Execute executes the program for an input parameters `objects` and returns a slice of output objects.
func (_fb *PSExecutor) Execute(objects []PSObject) ([]PSObject, error) {
	for _, _de := range objects {
		_fc := _fb.Stack.Push(_de)
		if _fc != nil {
			return nil, _fc
		}
	}
	_ff := _fb._b.Exec(_fb.Stack)
	if _ff != nil {
		_e.Log.Debug("\u0045x\u0065c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ff)
		return nil, _ff
	}
	_gc := []PSObject(*_fb.Stack)
	_fb.Stack.Empty()
	return _gc, nil
}

const _cf = 0.000001

// NewPSProgram returns an empty, initialized PSProgram.
func NewPSProgram() *PSProgram { return &PSProgram{} }
func _bbbe(_eed int) int {
	if _eed < 0 {
		return -_eed
	}
	return _eed
}

// PSStack defines a stack of PSObjects. PSObjects can be pushed on or pull from the stack.
type PSStack []PSObject

func (_aebd *PSOperand) ifelse(_ege *PSStack) error {
	_dba, _gbgg := _ege.Pop()
	if _gbgg != nil {
		return _gbgg
	}
	_bca, _gbgg := _ege.Pop()
	if _gbgg != nil {
		return _gbgg
	}
	_eee, _gbgg := _ege.Pop()
	if _gbgg != nil {
		return _gbgg
	}
	_be, _eag := _dba.(*PSProgram)
	if !_eag {
		return ErrTypeCheck
	}
	_fgd, _eag := _bca.(*PSProgram)
	if !_eag {
		return ErrTypeCheck
	}
	_dfeb, _eag := _eee.(*PSBoolean)
	if !_eag {
		return ErrTypeCheck
	}
	if _dfeb.Val {
		_fce := _fgd.Exec(_ege)
		return _fce
	}
	_gbgg = _be.Exec(_ege)
	return _gbgg
}
func (_dedf *PSOperand) floor(_ccb *PSStack) error {
	_cdf, _fe := _ccb.Pop()
	if _fe != nil {
		return _fe
	}
	if _bcg, _ce := _cdf.(*PSReal); _ce {
		_fe = _ccb.Push(MakeReal(_gbg.Floor(_bcg.Val)))
	} else if _eff, _aag := _cdf.(*PSInteger); _aag {
		_fe = _ccb.Push(MakeInteger(_eff.Val))
	} else {
		return ErrTypeCheck
	}
	return _fe
}

var ErrUnsupportedOperand = _c.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")

func (_ece *PSBoolean) Duplicate() PSObject {
	_cddd := PSBoolean{}
	_cddd.Val = _ece.Val
	return &_cddd
}
func (_adb *PSOperand) DebugString() string {
	return _f.Sprintf("\u006fp\u003a\u0027\u0025\u0073\u0027", *_adb)
}
func (_ecf *PSOperand) add(_fd *PSStack) error {
	_bfb, _geg := _fd.Pop()
	if _geg != nil {
		return _geg
	}
	_eae, _geg := _fd.Pop()
	if _geg != nil {
		return _geg
	}
	_ee, _gbgb := _bfb.(*PSReal)
	_bb, _da := _bfb.(*PSInteger)
	if !_gbgb && !_da {
		return ErrTypeCheck
	}
	_fcb, _cff := _eae.(*PSReal)
	_bg, _dbb := _eae.(*PSInteger)
	if !_cff && !_dbb {
		return ErrTypeCheck
	}
	if _da && _dbb {
		_ffggc := _bb.Val + _bg.Val
		_fdg := _fd.Push(MakeInteger(_ffggc))
		return _fdg
	}
	var _dgd float64
	if _gbgb {
		_dgd = _ee.Val
	} else {
		_dgd = float64(_bb.Val)
	}
	if _cff {
		_dgd += _fcb.Val
	} else {
		_dgd += float64(_bg.Val)
	}
	_geg = _fd.Push(MakeReal(_dgd))
	return _geg
}

// PSReal represents a real number.
type PSReal struct{ Val float64 }

// PSParser is a basic Postscript parser.
type PSParser struct{ _dcff *_gba.Reader }

// Empty empties the stack.
func (_adga *PSStack) Empty()          { *_adga = []PSObject{} }
func (_deg *PSOperand) String() string { return string(*_deg) }
func (_gefe *PSOperand) log(_eec *PSStack) error {
	_fdeb, _aacf := _eec.PopNumberAsFloat64()
	if _aacf != nil {
		return _aacf
	}
	_acbf := _gbg.Log10(_fdeb)
	_aacf = _eec.Push(MakeReal(_acbf))
	return _aacf
}
func (_bga *PSOperand) neg(_fgg *PSStack) error {
	_gde, _dga := _fgg.Pop()
	if _dga != nil {
		return _dga
	}
	if _bada, _efd := _gde.(*PSReal); _efd {
		_dga = _fgg.Push(MakeReal(-_bada.Val))
		return _dga
	} else if _cdcg, _bde := _gde.(*PSInteger); _bde {
		_dga = _fgg.Push(MakeInteger(-_cdcg.Val))
		return _dga
	} else {
		return ErrTypeCheck
	}
}

// Exec executes the operand `op` in the state specified by `stack`.
func (_aea *PSOperand) Exec(stack *PSStack) error {
	_db := ErrUnsupportedOperand
	switch *_aea {
	case "\u0061\u0062\u0073":
		_db = _aea.abs(stack)
	case "\u0061\u0064\u0064":
		_db = _aea.add(stack)
	case "\u0061\u006e\u0064":
		_db = _aea.and(stack)
	case "\u0061\u0074\u0061\u006e":
		_db = _aea.atan(stack)
	case "\u0062\u0069\u0074\u0073\u0068\u0069\u0066\u0074":
		_db = _aea.bitshift(stack)
	case "\u0063e\u0069\u006c\u0069\u006e\u0067":
		_db = _aea.ceiling(stack)
	case "\u0063\u006f\u0070\u0079":
		_db = _aea.copy(stack)
	case "\u0063\u006f\u0073":
		_db = _aea.cos(stack)
	case "\u0063\u0076\u0069":
		_db = _aea.cvi(stack)
	case "\u0063\u0076\u0072":
		_db = _aea.cvr(stack)
	case "\u0064\u0069\u0076":
		_db = _aea.div(stack)
	case "\u0064\u0075\u0070":
		_db = _aea.dup(stack)
	case "\u0065\u0071":
		_db = _aea.eq(stack)
	case "\u0065\u0078\u0063\u0068":
		_db = _aea.exch(stack)
	case "\u0065\u0078\u0070":
		_db = _aea.exp(stack)
	case "\u0066\u006c\u006fo\u0072":
		_db = _aea.floor(stack)
	case "\u0067\u0065":
		_db = _aea.ge(stack)
	case "\u0067\u0074":
		_db = _aea.gt(stack)
	case "\u0069\u0064\u0069\u0076":
		_db = _aea.idiv(stack)
	case "\u0069\u0066":
		_db = _aea.ifCondition(stack)
	case "\u0069\u0066\u0065\u006c\u0073\u0065":
		_db = _aea.ifelse(stack)
	case "\u0069\u006e\u0064e\u0078":
		_db = _aea.index(stack)
	case "\u006c\u0065":
		_db = _aea.le(stack)
	case "\u006c\u006f\u0067":
		_db = _aea.log(stack)
	case "\u006c\u006e":
		_db = _aea.ln(stack)
	case "\u006c\u0074":
		_db = _aea.lt(stack)
	case "\u006d\u006f\u0064":
		_db = _aea.mod(stack)
	case "\u006d\u0075\u006c":
		_db = _aea.mul(stack)
	case "\u006e\u0065":
		_db = _aea.ne(stack)
	case "\u006e\u0065\u0067":
		_db = _aea.neg(stack)
	case "\u006e\u006f\u0074":
		_db = _aea.not(stack)
	case "\u006f\u0072":
		_db = _aea.or(stack)
	case "\u0070\u006f\u0070":
		_db = _aea.pop(stack)
	case "\u0072\u006f\u0075n\u0064":
		_db = _aea.round(stack)
	case "\u0072\u006f\u006c\u006c":
		_db = _aea.roll(stack)
	case "\u0073\u0069\u006e":
		_db = _aea.sin(stack)
	case "\u0073\u0071\u0072\u0074":
		_db = _aea.sqrt(stack)
	case "\u0073\u0075\u0062":
		_db = _aea.sub(stack)
	case "\u0074\u0072\u0075\u006e\u0063\u0061\u0074\u0065":
		_db = _aea.truncate(stack)
	case "\u0078\u006f\u0072":
		_db = _aea.xor(stack)
	}
	return _db
}
func (_debd *PSOperand) ceiling(_dfe *PSStack) error {
	_adf, _dcc := _dfe.Pop()
	if _dcc != nil {
		return _dcc
	}
	if _cadf, _cafb := _adf.(*PSReal); _cafb {
		_dcc = _dfe.Push(MakeReal(_gbg.Ceil(_cadf.Val)))
	} else if _bdc, _dae := _adf.(*PSInteger); _dae {
		_dcc = _dfe.Push(MakeInteger(_bdc.Val))
	} else {
		_dcc = ErrTypeCheck
	}
	return _dcc
}

var ErrTypeCheck = _c.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")

func (_cabb *PSOperand) eq(_cga *PSStack) error {
	_bfcb, _bcc := _cga.Pop()
	if _bcc != nil {
		return _bcc
	}
	_acb, _bcc := _cga.Pop()
	if _bcc != nil {
		return _bcc
	}
	_abf, _ffad := _bfcb.(*PSBoolean)
	_dag, _aeba := _acb.(*PSBoolean)
	if _ffad || _aeba {
		var _bfa error
		if _ffad && _aeba {
			_bfa = _cga.Push(MakeBool(_abf.Val == _dag.Val))
		} else {
			_bfa = _cga.Push(MakeBool(false))
		}
		return _bfa
	}
	var _cdc float64
	var _cbg float64
	if _aaf, _bce := _bfcb.(*PSInteger); _bce {
		_cdc = float64(_aaf.Val)
	} else if _bfbe, _cda := _bfcb.(*PSReal); _cda {
		_cdc = _bfbe.Val
	} else {
		return ErrTypeCheck
	}
	if _dbg, _bab := _acb.(*PSInteger); _bab {
		_cbg = float64(_dbg.Val)
	} else if _ccg, _cgec := _acb.(*PSReal); _cgec {
		_cbg = _ccg.Val
	} else {
		return ErrTypeCheck
	}
	if _gbg.Abs(_cbg-_cdc) < _cf {
		_bcc = _cga.Push(MakeBool(true))
	} else {
		_bcc = _cga.Push(MakeBool(false))
	}
	return _bcc
}

// MakeOperand returns a new PSOperand object based on string `val`.
func MakeOperand(val string) *PSOperand { _ccc := PSOperand(val); return &_ccc }

var ErrStackUnderflow = _c.New("\u0073t\u0061c\u006b\u0020\u0075\u006e\u0064\u0065\u0072\u0066\u006c\u006f\u0077")

func (_bag *PSOperand) abs(_ea *PSStack) error {
	_ga, _cb := _ea.Pop()
	if _cb != nil {
		return _cb
	}
	if _gec, _cc := _ga.(*PSReal); _cc {
		_ed := _gec.Val
		if _ed < 0 {
			_cb = _ea.Push(MakeReal(-_ed))
		} else {
			_cb = _ea.Push(MakeReal(_ed))
		}
	} else if _bd, _deb := _ga.(*PSInteger); _deb {
		_gea := _bd.Val
		if _gea < 0 {
			_cb = _ea.Push(MakeInteger(-_gea))
		} else {
			_cb = _ea.Push(MakeInteger(_gea))
		}
	} else {
		return ErrTypeCheck
	}
	return _cb
}
func (_cbe *PSOperand) idiv(_bagb *PSStack) error {
	_adfb, _bdf := _bagb.Pop()
	if _bdf != nil {
		return _bdf
	}
	_fec, _bdf := _bagb.Pop()
	if _bdf != nil {
		return _bdf
	}
	_ggd, _bccc := _adfb.(*PSInteger)
	if !_bccc {
		return ErrTypeCheck
	}
	if _ggd.Val == 0 {
		return ErrUndefinedResult
	}
	_cgab, _bccc := _fec.(*PSInteger)
	if !_bccc {
		return ErrTypeCheck
	}
	_gfd := _cgab.Val / _ggd.Val
	_bdf = _bagb.Push(MakeInteger(_gfd))
	return _bdf
}
func (_cfae *PSOperand) mod(_dffe *PSStack) error {
	_agcd, _bcee := _dffe.Pop()
	if _bcee != nil {
		return _bcee
	}
	_fgcf, _bcee := _dffe.Pop()
	if _bcee != nil {
		return _bcee
	}
	_dde, _cdbe := _agcd.(*PSInteger)
	if !_cdbe {
		return ErrTypeCheck
	}
	if _dde.Val == 0 {
		return ErrUndefinedResult
	}
	_dbbg, _cdbe := _fgcf.(*PSInteger)
	if !_cdbe {
		return ErrTypeCheck
	}
	_fgf := _dbbg.Val % _dde.Val
	_bcee = _dffe.Push(MakeInteger(_fgf))
	return _bcee
}
func (_dfd *PSOperand) atan(_abdc *PSStack) error {
	_gda, _dab := _abdc.PopNumberAsFloat64()
	if _dab != nil {
		return _dab
	}
	_agf, _dab := _abdc.PopNumberAsFloat64()
	if _dab != nil {
		return _dab
	}
	if _gda == 0 {
		var _dd error
		if _agf < 0 {
			_dd = _abdc.Push(MakeReal(270))
		} else {
			_dd = _abdc.Push(MakeReal(90))
		}
		return _dd
	}
	_fbg := _agf / _gda
	_aba := _gbg.Atan(_fbg) * 180 / _gbg.Pi
	_dab = _abdc.Push(MakeReal(_aba))
	return _dab
}

// NewPSParser returns a new instance of the PDF Postscript parser from input data.
func NewPSParser(content []byte) *PSParser {
	_cbcb := PSParser{}
	_dddc := _gb.NewBuffer(content)
	_cbcb._dcff = _gba.NewReader(_dddc)
	return &_cbcb
}
func (_ag *PSReal) Duplicate() PSObject { _ec := PSReal{}; _ec.Val = _ag.Val; return &_ec }

// NewPSStack returns an initialized PSStack.
func NewPSStack() *PSStack { return &PSStack{} }
func (_gbcb *PSParser) parseNumber() (PSObject, error) {
	_ead, _aae := _ca.ParseNumber(_gbcb._dcff)
	if _aae != nil {
		return nil, _aae
	}
	switch _ecac := _ead.(type) {
	case *_ca.PdfObjectFloat:
		return MakeReal(float64(*_ecac)), nil
	case *_ca.PdfObjectInteger:
		return MakeInteger(int(*_ecac)), nil
	}
	return nil, _f.Errorf("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _ead)
}
func (_dcf *PSOperand) cvr(_cdb *PSStack) error {
	_edad, _gbb := _cdb.Pop()
	if _gbb != nil {
		return _gbb
	}
	if _abab, _baa := _edad.(*PSReal); _baa {
		_gbb = _cdb.Push(MakeReal(_abab.Val))
	} else if _ef, _gfc := _edad.(*PSInteger); _gfc {
		_gbb = _cdb.Push(MakeReal(float64(_ef.Val)))
	} else {
		return ErrTypeCheck
	}
	return _gbb
}
func (_ecbf *PSOperand) sin(_gffa *PSStack) error {
	_ffae, _dacc := _gffa.PopNumberAsFloat64()
	if _dacc != nil {
		return _dacc
	}
	_gaed := _gbg.Sin(_ffae * _gbg.Pi / 180.0)
	_dacc = _gffa.Push(MakeReal(_gaed))
	return _dacc
}
func (_aa *PSOperand) cvi(_eca *PSStack) error {
	_abg, _gca := _eca.Pop()
	if _gca != nil {
		return _gca
	}
	if _cfa, _fde := _abg.(*PSReal); _fde {
		_cge := int(_cfa.Val)
		_gca = _eca.Push(MakeInteger(_cge))
	} else if _gf, _fgca := _abg.(*PSInteger); _fgca {
		_fac := _gf.Val
		_gca = _eca.Push(MakeInteger(_fac))
	} else {
		return ErrTypeCheck
	}
	return _gca
}

// PSOperand represents a Postscript operand (text string).
type PSOperand string

var ErrUndefinedResult = _c.New("\u0075\u006e\u0064\u0065fi\u006e\u0065\u0064\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0065\u0072\u0072o\u0072")

func (_ad *PSInteger) String() string { return _f.Sprintf("\u0025\u0064", _ad.Val) }
func (_af *PSProgram) Duplicate() PSObject {
	_afd := &PSProgram{}
	for _, _ada := range *_af {
		_afd.Append(_ada.Duplicate())
	}
	return _afd
}

// PSProgram defines a Postscript program which is a series of PS objects (arguments, commands, programs etc).
type PSProgram []PSObject

func (_aed *PSOperand) copy(_fgc *PSStack) error {
	_dgeg, _eaec := _fgc.PopInteger()
	if _eaec != nil {
		return _eaec
	}
	if _dgeg < 0 {
		return ErrRangeCheck
	}
	if _dgeg > len(*_fgc) {
		return ErrRangeCheck
	}
	*_fgc = append(*_fgc, (*_fgc)[len(*_fgc)-_dgeg:]...)
	return nil
}

// MakeReal returns a new PSReal object initialized with `val`.
func MakeReal(val float64) *PSReal { _gfe := PSReal{}; _gfe.Val = val; return &_gfe }

// Parse parses the postscript and store as a program that can be executed.
func (_dbd *PSParser) Parse() (*PSProgram, error) {
	_dbd.skipSpaces()
	_cfdd, _cgc := _dbd._dcff.Peek(2)
	if _cgc != nil {
		return nil, _cgc
	}
	if _cfdd[0] != '{' {
		return nil, _c.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0053\u0020\u0050\u0072\u006f\u0067\u0072\u0061\u006d\u0020\u006e\u006f\u0074\u0020\u0073t\u0061\u0072\u0074\u0069\u006eg\u0020\u0077i\u0074\u0068\u0020\u007b")
	}
	_agd, _cgc := _dbd.parseFunction()
	if _cgc != nil && _cgc != _a.EOF {
		return nil, _cgc
	}
	return _agd, _cgc
}
func (_fdc *PSOperand) and(_faa *PSStack) error {
	_ggc, _debc := _faa.Pop()
	if _debc != nil {
		return _debc
	}
	_dgbg, _debc := _faa.Pop()
	if _debc != nil {
		return _debc
	}
	if _eda, _bfg := _ggc.(*PSBoolean); _bfg {
		_fca, _eaf := _dgbg.(*PSBoolean)
		if !_eaf {
			return ErrTypeCheck
		}
		_debc = _faa.Push(MakeBool(_eda.Val && _fca.Val))
		return _debc
	}
	if _cad, _dbf := _ggc.(*PSInteger); _dbf {
		_bba, _bad := _dgbg.(*PSInteger)
		if !_bad {
			return ErrTypeCheck
		}
		_debc = _faa.Push(MakeInteger(_cad.Val & _bba.Val))
		return _debc
	}
	return ErrTypeCheck
}
func (_afa *PSOperand) bitshift(_degf *PSStack) error {
	_cfd, _caf := _degf.PopInteger()
	if _caf != nil {
		return _caf
	}
	_ggcg, _caf := _degf.PopInteger()
	if _caf != nil {
		return _caf
	}
	var _ffa int
	if _cfd >= 0 {
		_ffa = _ggcg << uint(_cfd)
	} else {
		_ffa = _ggcg >> uint(-_cfd)
	}
	_caf = _degf.Push(MakeInteger(_ffa))
	return _caf
}

var ErrStackOverflow = _c.New("\u0073\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065r\u0066\u006c\u006f\u0077")

func (_bcgg *PSOperand) not(_ecg *PSStack) error {
	_bac, _aebg := _ecg.Pop()
	if _aebg != nil {
		return _aebg
	}
	if _ebe, _cfcb := _bac.(*PSBoolean); _cfcb {
		_aebg = _ecg.Push(MakeBool(!_ebe.Val))
		return _aebg
	} else if _cbc, _dgegc := _bac.(*PSInteger); _dgegc {
		_aebg = _ecg.Push(MakeInteger(^_cbc.Val))
		return _aebg
	} else {
		return ErrTypeCheck
	}
}
func (_dbc *PSOperand) exch(_bbd *PSStack) error {
	_ded, _dgdf := _bbd.Pop()
	if _dgdf != nil {
		return _dgdf
	}
	_cfb, _dgdf := _bbd.Pop()
	if _dgdf != nil {
		return _dgdf
	}
	_dgdf = _bbd.Push(_ded)
	if _dgdf != nil {
		return _dgdf
	}
	_dgdf = _bbd.Push(_cfb)
	return _dgdf
}
func (_efe *PSParser) parseBool() (*PSBoolean, error) {
	_gaec, _ecfe := _efe._dcff.Peek(4)
	if _ecfe != nil {
		return MakeBool(false), _ecfe
	}
	if (len(_gaec) >= 4) && (string(_gaec[:4]) == "\u0074\u0072\u0075\u0065") {
		_efe._dcff.Discard(4)
		return MakeBool(true), nil
	}
	_gaec, _ecfe = _efe._dcff.Peek(5)
	if _ecfe != nil {
		return MakeBool(false), _ecfe
	}
	if (len(_gaec) >= 5) && (string(_gaec[:5]) == "\u0066\u0061\u006cs\u0065") {
		_efe._dcff.Discard(5)
		return MakeBool(false), nil
	}
	return MakeBool(false), _c.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}
func (_cba *PSOperand) ifCondition(_dac *PSStack) error {
	_fbf, _geae := _dac.Pop()
	if _geae != nil {
		return _geae
	}
	_adg, _geae := _dac.Pop()
	if _geae != nil {
		return _geae
	}
	_ggg, _aac := _fbf.(*PSProgram)
	if !_aac {
		return ErrTypeCheck
	}
	_gcdb, _aac := _adg.(*PSBoolean)
	if !_aac {
		return ErrTypeCheck
	}
	if _gcdb.Val {
		_dea := _ggg.Exec(_dac)
		return _dea
	}
	return nil
}
func (_cgf *PSProgram) DebugString() string {
	_cfe := "\u007b\u0020"
	for _, _ffg := range *_cgf {
		_cfe += _ffg.DebugString()
		_cfe += "\u0020"
	}
	_cfe += "\u007d"
	return _cfe
}

// Append appends an object to the PSProgram.
func (_eg *PSProgram) Append(obj PSObject) { *_eg = append(*_eg, obj) }

// DebugString returns a descriptive string representation of the stack - intended for debugging.
func (_eccbe *PSStack) DebugString() string {
	_cfea := "\u005b\u0020"
	for _, _ebd := range *_eccbe {
		_cfea += _ebd.DebugString()
		_cfea += "\u0020"
	}
	_cfea += "\u005d"
	return _cfea
}
func (_efc *PSOperand) pop(_dded *PSStack) error {
	_, _fdbf := _dded.Pop()
	if _fdbf != nil {
		return _fdbf
	}
	return nil
}

// String returns a string representation of the stack.
func (_edcc *PSStack) String() string {
	_bfdc := "\u005b\u0020"
	for _, _ffbc := range *_edcc {
		_bfdc += _ffbc.String()
		_bfdc += "\u0020"
	}
	_bfdc += "\u005d"
	return _bfdc
}

// PSInteger represents an integer.
type PSInteger struct{ Val int }

func (_aad *PSParser) parseOperand() (*PSOperand, error) {
	var _edb []byte
	for {
		_cdcga, _bced := _aad._dcff.Peek(1)
		if _bced != nil {
			if _bced == _a.EOF {
				break
			}
			return nil, _bced
		}
		if _ca.IsDelimiter(_cdcga[0]) {
			break
		}
		if _ca.IsWhiteSpace(_cdcga[0]) {
			break
		}
		_aeafe, _ := _aad._dcff.ReadByte()
		_edb = append(_edb, _aeafe)
	}
	if len(_edb) == 0 {
		return nil, _c.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029")
	}
	return MakeOperand(string(_edb)), nil
}
func (_abc *PSParser) parseFunction() (*PSProgram, error) {
	_edfb, _ := _abc._dcff.ReadByte()
	if _edfb != '{' {
		return nil, _c.New("\u0069\u006ev\u0061\u006c\u0069d\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	}
	_cgdf := NewPSProgram()
	for {
		_abc.skipSpaces()
		_abc.skipComments()
		_egg, _dfed := _abc._dcff.Peek(2)
		if _dfed != nil {
			if _dfed == _a.EOF {
				break
			}
			return nil, _dfed
		}
		_e.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_egg))
		if _egg[0] == '}' {
			_e.Log.Trace("\u0045\u004f\u0046 \u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
			_abc._dcff.ReadByte()
			break
		} else if _egg[0] == '{' {
			_e.Log.Trace("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e!")
			_gegb, _gfb := _abc.parseFunction()
			if _gfb != nil {
				return nil, _gfb
			}
			_cgdf.Append(_gegb)
		} else if _ca.IsDecimalDigit(_egg[0]) || (_egg[0] == '-' && _ca.IsDecimalDigit(_egg[1])) {
			_e.Log.Trace("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!")
			_bgag, _bbdg := _abc.parseNumber()
			if _bbdg != nil {
				return nil, _bbdg
			}
			_cgdf.Append(_bgag)
		} else {
			_e.Log.Trace("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f")
			_egg, _ = _abc._dcff.Peek(5)
			_fegg := string(_egg)
			_e.Log.Trace("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073", _fegg)
			if (len(_fegg) > 4) && (_fegg[:5] == "\u0066\u0061\u006cs\u0065") {
				_fgb, _aagc := _abc.parseBool()
				if _aagc != nil {
					return nil, _aagc
				}
				_cgdf.Append(_fgb)
			} else if (len(_fegg) > 3) && (_fegg[:4] == "\u0074\u0072\u0075\u0065") {
				_gcage, _edc := _abc.parseBool()
				if _edc != nil {
					return nil, _edc
				}
				_cgdf.Append(_gcage)
			} else {
				_dbec, _fefb := _abc.parseOperand()
				if _fefb != nil {
					return nil, _fefb
				}
				_cgdf.Append(_dbec)
			}
		}
	}
	return _cgdf, nil
}

// Exec executes the program, typically leaving output values on the stack.
func (_cab *PSProgram) Exec(stack *PSStack) error {
	for _, _ffb := range *_cab {
		var _dge error
		switch _def := _ffb.(type) {
		case *PSInteger:
			_gcd := _def
			_dge = stack.Push(_gcd)
		case *PSReal:
			_dc := _def
			_dge = stack.Push(_dc)
		case *PSBoolean:
			_abd := _def
			_dge = stack.Push(_abd)
		case *PSProgram:
			_acd := _def
			_dge = stack.Push(_acd)
		case *PSOperand:
			_ggb := _def
			_dge = _ggb.Exec(stack)
		default:
			return ErrTypeCheck
		}
		if _dge != nil {
			return _dge
		}
	}
	return nil
}
func (_eba *PSReal) DebugString() string {
	return _f.Sprintf("\u0072e\u0061\u006c\u003a\u0025\u002e\u0035f", _eba.Val)
}
func (_agcf *PSOperand) exp(_geca *PSStack) error {
	_fcbc, _ged := _geca.PopNumberAsFloat64()
	if _ged != nil {
		return _ged
	}
	_bbc, _ged := _geca.PopNumberAsFloat64()
	if _ged != nil {
		return _ged
	}
	if _gbg.Abs(_fcbc) < 1 && _bbc < 0 {
		return ErrUndefinedResult
	}
	_efa := _gbg.Pow(_bbc, _fcbc)
	_ged = _geca.Push(MakeReal(_efa))
	return _ged
}
func (_cgdg *PSOperand) ln(_dec *PSStack) error {
	_ecc, _bef := _dec.PopNumberAsFloat64()
	if _bef != nil {
		return _bef
	}
	_dceb := _gbg.Log(_ecc)
	_bef = _dec.Push(MakeReal(_dceb))
	return _bef
}
func (_geaa *PSOperand) index(_aacb *PSStack) error {
	_aeaa, _bfgc := _aacb.Pop()
	if _bfgc != nil {
		return _bfgc
	}
	_bbce, _cgd := _aeaa.(*PSInteger)
	if !_cgd {
		return ErrTypeCheck
	}
	if _bbce.Val < 0 {
		return ErrRangeCheck
	}
	if _bbce.Val > len(*_aacb)-1 {
		return ErrStackUnderflow
	}
	_facb := (*_aacb)[len(*_aacb)-1-_bbce.Val]
	_bfgc = _aacb.Push(_facb.Duplicate())
	return _bfgc
}
func (_daf *PSOperand) sub(_cage *PSStack) error {
	_fed, _cgfd := _cage.Pop()
	if _cgfd != nil {
		return _cgfd
	}
	_edae, _cgfd := _cage.Pop()
	if _cgfd != nil {
		return _cgfd
	}
	_ecga, _dacf := _fed.(*PSReal)
	_bgad, _fgcb := _fed.(*PSInteger)
	if !_dacf && !_fgcb {
		return ErrTypeCheck
	}
	_gaea, _cbf := _edae.(*PSReal)
	_aga, _fegc := _edae.(*PSInteger)
	if !_cbf && !_fegc {
		return ErrTypeCheck
	}
	if _fgcb && _fegc {
		_bdfb := _aga.Val - _bgad.Val
		_geb := _cage.Push(MakeInteger(_bdfb))
		return _geb
	}
	var _fefg float64 = 0
	if _cbf {
		_fefg = _gaea.Val
	} else {
		_fefg = float64(_aga.Val)
	}
	if _dacf {
		_fefg -= _ecga.Val
	} else {
		_fefg -= float64(_bgad.Val)
	}
	_cgfd = _cage.Push(MakeReal(_fefg))
	return _cgfd
}
func (_dfb *PSParser) skipSpaces() (int, error) {
	_baea := 0
	for {
		_ccf, _dgfg := _dfb._dcff.Peek(1)
		if _dgfg != nil {
			return 0, _dgfg
		}
		if _ca.IsWhiteSpace(_ccf[0]) {
			_dfb._dcff.ReadByte()
			_baea++
		} else {
			break
		}
	}
	return _baea, nil
}

// Pop pops an object from the top of the stack.
func (_gaaf *PSStack) Pop() (PSObject, error) {
	if len(*_gaaf) < 1 {
		return nil, ErrStackUnderflow
	}
	_fff := (*_gaaf)[len(*_gaaf)-1]
	*_gaaf = (*_gaaf)[0 : len(*_gaaf)-1]
	return _fff, nil
}
func (_fba *PSOperand) gt(_bdb *PSStack) error {
	_dbe, _fab := _bdb.PopNumberAsFloat64()
	if _fab != nil {
		return _fab
	}
	_dff, _fab := _bdb.PopNumberAsFloat64()
	if _fab != nil {
		return _fab
	}
	if _gbg.Abs(_dff-_dbe) < _cf {
		_bgd := _bdb.Push(MakeBool(false))
		return _bgd
	} else if _dff > _dbe {
		_bfd := _bdb.Push(MakeBool(true))
		return _bfd
	} else {
		_effe := _bdb.Push(MakeBool(false))
		return _effe
	}
}
func (_cg *PSInteger) DebugString() string {
	return _f.Sprintf("\u0069\u006e\u0074\u003a\u0025\u0064", _cg.Val)
}

// Push pushes an object on top of the stack.
func (_fcea *PSStack) Push(obj PSObject) error {
	if len(*_fcea) > 100 {
		return ErrStackOverflow
	}
	*_fcea = append(*_fcea, obj)
	return nil
}

// PSObjectArrayToFloat64Array converts []PSObject into a []float64 array. Each PSObject must represent a number,
// otherwise a ErrTypeCheck error occurs.
func PSObjectArrayToFloat64Array(objects []PSObject) ([]float64, error) {
	var _cd []float64
	for _, _gg := range objects {
		if _fa, _ge := _gg.(*PSInteger); _ge {
			_cd = append(_cd, float64(_fa.Val))
		} else if _ba, _eb := _gg.(*PSReal); _eb {
			_cd = append(_cd, _ba.Val)
		} else {
			return nil, ErrTypeCheck
		}
	}
	return _cd, nil
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
func MakeInteger(val int) *PSInteger { _ddde := PSInteger{}; _ddde.Val = val; return &_ddde }
func (_fdec *PSOperand) xor(_eeec *PSStack) error {
	_fae, _bae := _eeec.Pop()
	if _bae != nil {
		return _bae
	}
	_dbff, _bae := _eeec.Pop()
	if _bae != nil {
		return _bae
	}
	if _febb, _agfb := _fae.(*PSBoolean); _agfb {
		_eab, _ced := _dbff.(*PSBoolean)
		if !_ced {
			return ErrTypeCheck
		}
		_bae = _eeec.Push(MakeBool(_febb.Val != _eab.Val))
		return _bae
	}
	if _eabd, _gdcbd := _fae.(*PSInteger); _gdcbd {
		_bgade, _ffadg := _dbff.(*PSInteger)
		if !_ffadg {
			return ErrTypeCheck
		}
		_bae = _eeec.Push(MakeInteger(_eabd.Val ^ _bgade.Val))
		return _bae
	}
	return ErrTypeCheck
}
func (_fcg *PSOperand) Duplicate() PSObject { _ffgg := *_fcg; return &_ffgg }
func (_gd *PSProgram) String() string {
	_dgb := "\u007b\u0020"
	for _, _fg := range *_gd {
		_dgb += _fg.String()
		_dgb += "\u0020"
	}
	_dgb += "\u007d"
	return _dgb
}
func (_cgfc *PSOperand) le(_gcag *PSStack) error {
	_dgf, _fadc := _gcag.PopNumberAsFloat64()
	if _fadc != nil {
		return _fadc
	}
	_efag, _fadc := _gcag.PopNumberAsFloat64()
	if _fadc != nil {
		return _fadc
	}
	if _gbg.Abs(_efag-_dgf) < _cf {
		_acf := _gcag.Push(MakeBool(true))
		return _acf
	} else if _efag < _dgf {
		_gce := _gcag.Push(MakeBool(true))
		return _gce
	} else {
		_gdc := _gcag.Push(MakeBool(false))
		return _gdc
	}
}

var ErrRangeCheck = _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")

func (_dcda *PSOperand) ne(_gff *PSStack) error {
	_dcg := _dcda.eq(_gff)
	if _dcg != nil {
		return _dcg
	}
	_dcg = _dcda.not(_gff)
	return _dcg
}

// PopNumberAsFloat64 pops and return the numeric value of the top of the stack as a float64.
// Real or integer only.
func (_fcd *PSStack) PopNumberAsFloat64() (float64, error) {
	_efad, _afb := _fcd.Pop()
	if _afb != nil {
		return 0, _afb
	}
	if _gee, _gag := _efad.(*PSReal); _gag {
		return _gee.Val, nil
	} else if _ggaa, _ega := _efad.(*PSInteger); _ega {
		return float64(_ggaa.Val), nil
	} else {
		return 0, ErrTypeCheck
	}
}
func (_dcba *PSOperand) mul(_cfdc *PSStack) error {
	_ccga, _fbgd := _cfdc.Pop()
	if _fbgd != nil {
		return _fbgd
	}
	_dceg, _fbgd := _cfdc.Pop()
	if _fbgd != nil {
		return _fbgd
	}
	_eccb, _gcab := _ccga.(*PSReal)
	_edf, _gaeg := _ccga.(*PSInteger)
	if !_gcab && !_gaeg {
		return ErrTypeCheck
	}
	_ecb, _bfge := _dceg.(*PSReal)
	_cgfe, _ggbe := _dceg.(*PSInteger)
	if !_bfge && !_ggbe {
		return ErrTypeCheck
	}
	if _gaeg && _ggbe {
		_fdd := _edf.Val * _cgfe.Val
		_fdb := _cfdc.Push(MakeInteger(_fdd))
		return _fdb
	}
	var _decc float64
	if _gcab {
		_decc = _eccb.Val
	} else {
		_decc = float64(_edf.Val)
	}
	if _bfge {
		_decc *= _ecb.Val
	} else {
		_decc *= float64(_cgfe.Val)
	}
	_fbgd = _cfdc.Push(MakeReal(_decc))
	return _fbgd
}

// PopInteger specificially pops an integer from the top of the stack, returning the value as an int.
func (_ecee *PSStack) PopInteger() (int, error) {
	_afgf, _daca := _ecee.Pop()
	if _daca != nil {
		return 0, _daca
	}
	if _bccb, _aeg := _afgf.(*PSInteger); _aeg {
		return _bccb.Val, nil
	}
	return 0, ErrTypeCheck
}
func (_bfbg *PSParser) skipComments() error {
	if _, _cddb := _bfbg.skipSpaces(); _cddb != nil {
		return _cddb
	}
	_gceg := true
	for {
		_efac, _gbcf := _bfbg._dcff.Peek(1)
		if _gbcf != nil {
			_e.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _gbcf.Error())
			return _gbcf
		}
		if _gceg && _efac[0] != '%' {
			return nil
		}
		_gceg = false
		if (_efac[0] != '\r') && (_efac[0] != '\n') {
			_bfbg._dcff.ReadByte()
		} else {
			break
		}
	}
	return _bfbg.skipComments()
}
func (_gdcb *PSOperand) round(_cdbd *PSStack) error {
	_fga, _bfba := _cdbd.Pop()
	if _bfba != nil {
		return _bfba
	}
	if _gbf, _bed := _fga.(*PSReal); _bed {
		_bfba = _cdbd.Push(MakeReal(_gbg.Floor(_gbf.Val + 0.5)))
	} else if _agce, _dda := _fga.(*PSInteger); _dda {
		_bfba = _cdbd.Push(MakeInteger(_agce.Val))
	} else {
		return ErrTypeCheck
	}
	return _bfba
}
func (_aeb *PSBoolean) DebugString() string {
	return _f.Sprintf("\u0062o\u006f\u006c\u003a\u0025\u0076", _aeb.Val)
}
func (_fag *PSOperand) cos(_bc *PSStack) error {
	_bfe, _aeac := _bc.PopNumberAsFloat64()
	if _aeac != nil {
		return _aeac
	}
	_gaf := _gbg.Cos(_bfe * _gbg.Pi / 180.0)
	_aeac = _bc.Push(MakeReal(_gaf))
	return _aeac
}
func (_ceb *PSOperand) ge(_feg *PSStack) error {
	_dce, _gbbc := _feg.PopNumberAsFloat64()
	if _gbbc != nil {
		return _gbbc
	}
	_gef, _gbbc := _feg.PopNumberAsFloat64()
	if _gbbc != nil {
		return _gbbc
	}
	if _gbg.Abs(_gef-_dce) < _cf {
		_debe := _feg.Push(MakeBool(true))
		return _debe
	} else if _gef > _dce {
		_ade := _feg.Push(MakeBool(true))
		return _ade
	} else {
		_badc := _feg.Push(MakeBool(false))
		return _badc
	}
}
func (_ae *PSReal) String() string { return _f.Sprintf("\u0025\u002e\u0035\u0066", _ae.Val) }
func (_ace *PSOperand) sqrt(_eeb *PSStack) error {
	_abag, _gbc := _eeb.PopNumberAsFloat64()
	if _gbc != nil {
		return _gbc
	}
	if _abag < 0 {
		return ErrRangeCheck
	}
	_gcdbg := _gbg.Sqrt(_abag)
	_gbc = _eeb.Push(MakeReal(_gcdbg))
	return _gbc
}
func (_ddg *PSOperand) lt(_geag *PSStack) error {
	_ccd, _cabe := _geag.PopNumberAsFloat64()
	if _cabe != nil {
		return _cabe
	}
	_dcd, _cabe := _geag.PopNumberAsFloat64()
	if _cabe != nil {
		return _cabe
	}
	if _gbg.Abs(_dcd-_ccd) < _cf {
		_gdca := _geag.Push(MakeBool(false))
		return _gdca
	} else if _dcd < _ccd {
		_gfdb := _geag.Push(MakeBool(true))
		return _gfdb
	} else {
		_cfc := _geag.Push(MakeBool(false))
		return _cfc
	}
}

// PSBoolean represents a boolean value.
type PSBoolean struct{ Val bool }
