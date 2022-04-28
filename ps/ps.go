package ps

import (
	_gc "bufio"
	_g "bytes"
	_ba "errors"
	_baf "fmt"
	_a "io"
	_ac "math"

	_f "bitbucket.org/shenghui0779/gopdf/common"
	_af "bitbucket.org/shenghui0779/gopdf/core"
)

// Pop pops an object from the top of the stack.
func (_dcb *PSStack) Pop() (PSObject, error) {
	if len(*_dcb) < 1 {
		return nil, ErrStackUnderflow
	}
	_edg := (*_dcb)[len(*_dcb)-1]
	*_dcb = (*_dcb)[0 : len(*_dcb)-1]
	return _edg, nil
}
func (_egcf *PSParser) parseOperand() (*PSOperand, error) {
	var _dfe []byte
	for {
		_daa, _eba := _egcf._geca.Peek(1)
		if _eba != nil {
			if _eba == _a.EOF {
				break
			}
			return nil, _eba
		}
		if _af.IsDelimiter(_daa[0]) {
			break
		}
		if _af.IsWhiteSpace(_daa[0]) {
			break
		}
		_bdg, _ := _egcf._geca.ReadByte()
		_dfe = append(_dfe, _bdg)
	}
	if len(_dfe) == 0 {
		return nil, _ba.New("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029")
	}
	return MakeOperand(string(_dfe)), nil
}
func (_dgag *PSOperand) sin(_afce *PSStack) error {
	_cgcc, _baee := _afce.PopNumberAsFloat64()
	if _baee != nil {
		return _baee
	}
	_gfd := _ac.Sin(_cgcc * _ac.Pi / 180.0)
	_baee = _afce.Push(MakeReal(_gfd))
	return _baee
}
func (_abea *PSOperand) abs(_bda *PSStack) error {
	_dbf, _cf := _bda.Pop()
	if _cf != nil {
		return _cf
	}
	if _eac, _ff := _dbf.(*PSReal); _ff {
		_bed := _eac.Val
		if _bed < 0 {
			_cf = _bda.Push(MakeReal(-_bed))
		} else {
			_cf = _bda.Push(MakeReal(_bed))
		}
	} else if _aeg, _fad := _dbf.(*PSInteger); _fad {
		_gg := _aeg.Val
		if _gg < 0 {
			_cf = _bda.Push(MakeInteger(-_gg))
		} else {
			_cf = _bda.Push(MakeInteger(_gg))
		}
	} else {
		return ErrTypeCheck
	}
	return _cf
}
func (_ea *PSInteger) DebugString() string {
	return _baf.Sprintf("\u0069\u006e\u0074\u003a\u0025\u0064", _ea.Val)
}
func (_fcb *PSOperand) ifelse(_fead *PSStack) error {
	_cbb, _gaeea := _fead.Pop()
	if _gaeea != nil {
		return _gaeea
	}
	_cdc, _gaeea := _fead.Pop()
	if _gaeea != nil {
		return _gaeea
	}
	_ebf, _gaeea := _fead.Pop()
	if _gaeea != nil {
		return _gaeea
	}
	_fdge, _efa := _cbb.(*PSProgram)
	if !_efa {
		return ErrTypeCheck
	}
	_dbgb, _efa := _cdc.(*PSProgram)
	if !_efa {
		return ErrTypeCheck
	}
	_geda, _efa := _ebf.(*PSBoolean)
	if !_efa {
		return ErrTypeCheck
	}
	if _geda.Val {
		_bgaf := _dbgb.Exec(_fead)
		return _bgaf
	}
	_gaeea = _fdge.Exec(_fead)
	return _gaeea
}
func (_gae *PSOperand) cos(_aca *PSStack) error {
	_bcc, _fbg := _aca.PopNumberAsFloat64()
	if _fbg != nil {
		return _fbg
	}
	_ec := _ac.Cos(_bcc * _ac.Pi / 180.0)
	_fbg = _aca.Push(MakeReal(_ec))
	return _fbg
}
func (_gbb *PSOperand) ne(_cab *PSStack) error {
	_agfb := _gbb.eq(_cab)
	if _agfb != nil {
		return _agfb
	}
	_agfb = _gbb.not(_cab)
	return _agfb
}
func (_caf *PSOperand) div(_afd *PSStack) error {
	_fbb, _aaa := _afd.Pop()
	if _aaa != nil {
		return _aaa
	}
	_dba, _aaa := _afd.Pop()
	if _aaa != nil {
		return _aaa
	}
	_ece, _dde := _fbb.(*PSReal)
	_cdb, _gca := _fbb.(*PSInteger)
	if !_dde && !_gca {
		return ErrTypeCheck
	}
	if _dde && _ece.Val == 0 {
		return ErrUndefinedResult
	}
	if _gca && _cdb.Val == 0 {
		return ErrUndefinedResult
	}
	_fcg, _aaff := _dba.(*PSReal)
	_gee, _ggc := _dba.(*PSInteger)
	if !_aaff && !_ggc {
		return ErrTypeCheck
	}
	var _cfb float64
	if _aaff {
		_cfb = _fcg.Val
	} else {
		_cfb = float64(_gee.Val)
	}
	if _dde {
		_cfb /= _ece.Val
	} else {
		_cfb /= float64(_cdb.Val)
	}
	_aaa = _afd.Push(MakeReal(_cfb))
	return _aaa
}

// PSReal represents a real number.
type PSReal struct{ Val float64 }

func (_bedd *PSOperand) pop(_gafa *PSStack) error {
	_, _fdf := _gafa.Pop()
	if _fdf != nil {
		return _fdf
	}
	return nil
}

// MakeBool returns a new PSBoolean object initialized with `val`.
func MakeBool(val bool) *PSBoolean { _fgde := PSBoolean{}; _fgde.Val = val; return &_fgde }

// PSBoolean represents a boolean value.
type PSBoolean struct{ Val bool }

func (_cc *PSReal) DebugString() string {
	return _baf.Sprintf("\u0072e\u0061\u006c\u003a\u0025\u002e\u0035f", _cc.Val)
}
func (_bab *PSBoolean) DebugString() string {
	return _baf.Sprintf("\u0062o\u006f\u006c\u003a\u0025\u0076", _bab.Val)
}
func (_fgg *PSOperand) sqrt(_gcac *PSStack) error {
	_dfcb, _ace := _gcac.PopNumberAsFloat64()
	if _ace != nil {
		return _ace
	}
	if _dfcb < 0 {
		return ErrRangeCheck
	}
	_gcfg := _ac.Sqrt(_dfcb)
	_ace = _gcac.Push(MakeReal(_gcfg))
	return _ace
}
func (_dc *PSBoolean) String() string { return _baf.Sprintf("\u0025\u0076", _dc.Val) }
func (_dbg *PSOperand) dup(_ecee *PSStack) error {
	_fdg, _efe := _ecee.Pop()
	if _efe != nil {
		return _efe
	}
	_efe = _ecee.Push(_fdg)
	if _efe != nil {
		return _efe
	}
	_efe = _ecee.Push(_fdg.Duplicate())
	return _efe
}
func (_ccbb *PSOperand) exp(_gfce *PSStack) error {
	_ffb, _dbc := _gfce.PopNumberAsFloat64()
	if _dbc != nil {
		return _dbc
	}
	_ccc, _dbc := _gfce.PopNumberAsFloat64()
	if _dbc != nil {
		return _dbc
	}
	if _ac.Abs(_ffb) < 1 && _ccc < 0 {
		return ErrUndefinedResult
	}
	_cdff := _ac.Pow(_ccc, _ffb)
	_dbc = _gfce.Push(MakeReal(_cdff))
	return _dbc
}
func (_dfc *PSOperand) copy(_afbb *PSStack) error {
	_aea, _fba := _afbb.PopInteger()
	if _fba != nil {
		return _fba
	}
	if _aea < 0 {
		return ErrRangeCheck
	}
	if _aea > len(*_afbb) {
		return ErrRangeCheck
	}
	*_afbb = append(*_afbb, (*_afbb)[len(*_afbb)-_aea:]...)
	return nil
}

// NewPSProgram returns an empty, initialized PSProgram.
func NewPSProgram() *PSProgram              { return &PSProgram{} }
func (_afb *PSBoolean) Duplicate() PSObject { _gfa := PSBoolean{}; _gfa.Val = _afb.Val; return &_gfa }

// NewPSExecutor returns an initialized PSExecutor for an input `program`.
func NewPSExecutor(program *PSProgram) *PSExecutor {
	_fa := &PSExecutor{}
	_fa.Stack = NewPSStack()
	_fa._acf = program
	return _fa
}
func (_cad *PSParser) parseBool() (*PSBoolean, error) {
	_bgab, _ffg := _cad._geca.Peek(4)
	if _ffg != nil {
		return MakeBool(false), _ffg
	}
	if (len(_bgab) >= 4) && (string(_bgab[:4]) == "\u0074\u0072\u0075\u0065") {
		_cad._geca.Discard(4)
		return MakeBool(true), nil
	}
	_bgab, _ffg = _cad._geca.Peek(5)
	if _ffg != nil {
		return MakeBool(false), _ffg
	}
	if (len(_bgab) >= 5) && (string(_bgab[:5]) == "\u0066\u0061\u006cs\u0065") {
		_cad._geca.Discard(5)
		return MakeBool(false), nil
	}
	return MakeBool(false), _ba.New("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg")
}

// Push pushes an object on top of the stack.
func (_baae *PSStack) Push(obj PSObject) error {
	if len(*_baae) > 100 {
		return ErrStackOverflow
	}
	*_baae = append(*_baae, obj)
	return nil
}
func (_gfed *PSParser) skipSpaces() (int, error) {
	_fdag := 0
	for {
		_gacg, _baga := _gfed._geca.Peek(1)
		if _baga != nil {
			return 0, _baga
		}
		if _af.IsWhiteSpace(_gacg[0]) {
			_gfed._geca.ReadByte()
			_fdag++
		} else {
			break
		}
	}
	return _fdag, nil
}

var ErrUnsupportedOperand = _ba.New("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064")

func (_ebef *PSOperand) ge(_eadd *PSStack) error {
	_fgc, _bdb := _eadd.PopNumberAsFloat64()
	if _bdb != nil {
		return _bdb
	}
	_faa, _bdb := _eadd.PopNumberAsFloat64()
	if _bdb != nil {
		return _bdb
	}
	if _ac.Abs(_faa-_fgc) < _fe {
		_fbfb := _eadd.Push(MakeBool(true))
		return _fbfb
	} else if _faa > _fgc {
		_bgb := _eadd.Push(MakeBool(true))
		return _bgb
	} else {
		_agd := _eadd.Push(MakeBool(false))
		return _agd
	}
}
func (_fd *PSProgram) DebugString() string {
	_ead := "\u007b\u0020"
	for _, _bd := range *_fd {
		_ead += _bd.DebugString()
		_ead += "\u0020"
	}
	_ead += "\u007d"
	return _ead
}

// PSParser is a basic Postscript parser.
type PSParser struct{ _geca *_gc.Reader }

func (_cgcf *PSOperand) mod(_eaaa *PSStack) error {
	_dbb, _fegb := _eaaa.Pop()
	if _fegb != nil {
		return _fegb
	}
	_bafe, _fegb := _eaaa.Pop()
	if _fegb != nil {
		return _fegb
	}
	_adbg, _afec := _dbb.(*PSInteger)
	if !_afec {
		return ErrTypeCheck
	}
	if _adbg.Val == 0 {
		return ErrUndefinedResult
	}
	_bdege, _afec := _bafe.(*PSInteger)
	if !_afec {
		return ErrTypeCheck
	}
	_ada := _bdege.Val % _adbg.Val
	_fegb = _eaaa.Push(MakeInteger(_ada))
	return _fegb
}
func (_aec *PSOperand) exch(_abb *PSStack) error {
	_gge, _gdb := _abb.Pop()
	if _gdb != nil {
		return _gdb
	}
	_ggf, _gdb := _abb.Pop()
	if _gdb != nil {
		return _gdb
	}
	_gdb = _abb.Push(_gge)
	if _gdb != nil {
		return _gdb
	}
	_gdb = _abb.Push(_ggf)
	return _gdb
}

// PSStack defines a stack of PSObjects. PSObjects can be pushed on or pull from the stack.
type PSStack []PSObject

var ErrStackOverflow = _ba.New("\u0073\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065r\u0066\u006c\u006f\u0077")

// PopInteger specificially pops an integer from the top of the stack, returning the value as an int.
func (_ddbg *PSStack) PopInteger() (int, error) {
	_ddfc, _ggca := _ddbg.Pop()
	if _ggca != nil {
		return 0, _ggca
	}
	if _dfb, _dfa := _ddfc.(*PSInteger); _dfa {
		return _dfb.Val, nil
	}
	return 0, ErrTypeCheck
}

// Execute executes the program for an input parameters `objects` and returns a slice of output objects.
func (_d *PSExecutor) Execute(objects []PSObject) ([]PSObject, error) {
	for _, _bb := range objects {
		_cg := _d.Stack.Push(_bb)
		if _cg != nil {
			return nil, _cg
		}
	}
	_eg := _d._acf.Exec(_d.Stack)
	if _eg != nil {
		_f.Log.Debug("\u0045x\u0065c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _eg)
		return nil, _eg
	}
	_ceb := []PSObject(*_d.Stack)
	_d.Stack.Empty()
	return _ceb, nil
}

// PSExecutor has its own execution stack and is used to executre a PS routine (program).
type PSExecutor struct {
	Stack *PSStack
	_acf  *PSProgram
}

func (_fadg *PSOperand) bitshift(_cff *PSStack) error {
	_gd, _cbc := _cff.PopInteger()
	if _cbc != nil {
		return _cbc
	}
	_aee, _cbc := _cff.PopInteger()
	if _cbc != nil {
		return _cbc
	}
	var _cea int
	if _gd >= 0 {
		_cea = _aee << uint(_gd)
	} else {
		_cea = _aee >> uint(-_gd)
	}
	_cbc = _cff.Push(MakeInteger(_cea))
	return _cbc
}

// PSOperand represents a Postscript operand (text string).
type PSOperand string

func (_gff *PSOperand) cvr(_edc *PSStack) error {
	_bga, _bafc := _edc.Pop()
	if _bafc != nil {
		return _bafc
	}
	if _aafa, _abg := _bga.(*PSReal); _abg {
		_bafc = _edc.Push(MakeReal(_aafa.Val))
	} else if _add, _fge := _bga.(*PSInteger); _fge {
		_bafc = _edc.Push(MakeReal(float64(_add.Val)))
	} else {
		return ErrTypeCheck
	}
	return _bafc
}
func (_ggg *PSOperand) lt(_egea *PSStack) error {
	_bdbc, _abbf := _egea.PopNumberAsFloat64()
	if _abbf != nil {
		return _abbf
	}
	_ceg, _abbf := _egea.PopNumberAsFloat64()
	if _abbf != nil {
		return _abbf
	}
	if _ac.Abs(_ceg-_bdbc) < _fe {
		_cde := _egea.Push(MakeBool(false))
		return _cde
	} else if _ceg < _bdbc {
		_gbd := _egea.Push(MakeBool(true))
		return _gbd
	} else {
		_bbe := _egea.Push(MakeBool(false))
		return _bbe
	}
}
func (_bg *PSOperand) ceiling(_dab *PSStack) error {
	_ddd, _cgg := _dab.Pop()
	if _cgg != nil {
		return _cgg
	}
	if _cd, _baa := _ddd.(*PSReal); _baa {
		_cgg = _dab.Push(MakeReal(_ac.Ceil(_cd.Val)))
	} else if _eade, _dgc := _ddd.(*PSInteger); _dgc {
		_cgg = _dab.Push(MakeInteger(_eade.Val))
	} else {
		_cgg = ErrTypeCheck
	}
	return _cgg
}

// PSInteger represents an integer.
type PSInteger struct{ Val int }

func (_eadb *PSOperand) le(_cafb *PSStack) error {
	_fac, _eeb := _cafb.PopNumberAsFloat64()
	if _eeb != nil {
		return _eeb
	}
	_fef, _eeb := _cafb.PopNumberAsFloat64()
	if _eeb != nil {
		return _eeb
	}
	if _ac.Abs(_fef-_fac) < _fe {
		_gfag := _cafb.Push(MakeBool(true))
		return _gfag
	} else if _fef < _fac {
		_bdeg := _cafb.Push(MakeBool(true))
		return _bdeg
	} else {
		_eab := _cafb.Push(MakeBool(false))
		return _eab
	}
}

// MakeInteger returns a new PSInteger object initialized with `val`.
func MakeInteger(val int) *PSInteger { _daga := PSInteger{}; _daga.Val = val; return &_daga }
func (_bgdg *PSParser) parseNumber() (PSObject, error) {
	_deae, _ccca := _af.ParseNumber(_bgdg._geca)
	if _ccca != nil {
		return nil, _ccca
	}
	switch _gbec := _deae.(type) {
	case *_af.PdfObjectFloat:
		return MakeReal(float64(*_gbec)), nil
	case *_af.PdfObjectInteger:
		return MakeInteger(int(*_gbec)), nil
	}
	return nil, _baf.Errorf("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054", _deae)
}

// PSProgram defines a Postscript program which is a series of PS objects (arguments, commands, programs etc).
type PSProgram []PSObject

func (_ceef *PSOperand) ln(_fcf *PSStack) error {
	_bae, _eaaf := _fcf.PopNumberAsFloat64()
	if _eaaf != nil {
		return _eaaf
	}
	_aba := _ac.Log(_bae)
	_eaaf = _fcf.Push(MakeReal(_aba))
	return _eaaf
}

// NewPSStack returns an initialized PSStack.
func NewPSStack() *PSStack { return &PSStack{} }

// PSObjectArrayToFloat64Array converts []PSObject into a []float64 array. Each PSObject must represent a number,
// otherwise a ErrTypeCheck error occurs.
func PSObjectArrayToFloat64Array(objects []PSObject) ([]float64, error) {
	var _aa []float64
	for _, _fb := range objects {
		if _e, _c := _fb.(*PSInteger); _c {
			_aa = append(_aa, float64(_e.Val))
		} else if _gb, _ce := _fb.(*PSReal); _ce {
			_aa = append(_aa, _gb.Val)
		} else {
			return nil, ErrTypeCheck
		}
	}
	return _aa, nil
}
func (_gec *PSOperand) round(_ggcf *PSStack) error {
	_acfgd, _bgdc := _ggcf.Pop()
	if _bgdc != nil {
		return _bgdc
	}
	if _ebg, _fdac := _acfgd.(*PSReal); _fdac {
		_bgdc = _ggcf.Push(MakeReal(_ac.Floor(_ebg.Val + 0.5)))
	} else if _ddg, _eec := _acfgd.(*PSInteger); _eec {
		_bgdc = _ggcf.Push(MakeInteger(_ddg.Val))
	} else {
		return ErrTypeCheck
	}
	return _bgdc
}

var ErrUndefinedResult = _ba.New("\u0075\u006e\u0064\u0065fi\u006e\u0065\u0064\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0065\u0072\u0072o\u0072")

// Parse parses the postscript and store as a program that can be executed.
func (_dbbd *PSParser) Parse() (*PSProgram, error) {
	_dbbd.skipSpaces()
	_gad, _deba := _dbbd._geca.Peek(2)
	if _deba != nil {
		return nil, _deba
	}
	if _gad[0] != '{' {
		return nil, _ba.New("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0053\u0020\u0050\u0072\u006f\u0067\u0072\u0061\u006d\u0020\u006e\u006f\u0074\u0020\u0073t\u0061\u0072\u0074\u0069\u006eg\u0020\u0077i\u0074\u0068\u0020\u007b")
	}
	_eadba, _deba := _dbbd.parseFunction()
	if _deba != nil && _deba != _a.EOF {
		return nil, _deba
	}
	return _eadba, _deba
}
func (_cbcc *PSOperand) xor(_cfe *PSStack) error {
	_abc, _cefa := _cfe.Pop()
	if _cefa != nil {
		return _cefa
	}
	_fadc, _cefa := _cfe.Pop()
	if _cefa != nil {
		return _cefa
	}
	if _fcfc, _begb := _abc.(*PSBoolean); _begb {
		_aegd, _edcc := _fadc.(*PSBoolean)
		if !_edcc {
			return ErrTypeCheck
		}
		_cefa = _cfe.Push(MakeBool(_fcfc.Val != _aegd.Val))
		return _cefa
	}
	if _cgga, _dfga := _abc.(*PSInteger); _dfga {
		_bafd, _ebd := _fadc.(*PSInteger)
		if !_ebd {
			return ErrTypeCheck
		}
		_cefa = _cfe.Push(MakeInteger(_cgga.Val ^ _bafd.Val))
		return _cefa
	}
	return ErrTypeCheck
}

var ErrStackUnderflow = _ba.New("\u0073t\u0061c\u006b\u0020\u0075\u006e\u0064\u0065\u0072\u0066\u006c\u006f\u0077")

// MakeOperand returns a new PSOperand object based on string `val`.
func MakeOperand(val string) *PSOperand { _gcbd := PSOperand(val); return &_gcbd }
func (_cdg *PSOperand) eq(_dge *PSStack) error {
	_fbf, _egg := _dge.Pop()
	if _egg != nil {
		return _egg
	}
	_ecc, _egg := _dge.Pop()
	if _egg != nil {
		return _egg
	}
	_ecf, _bacd := _fbf.(*PSBoolean)
	_dae, _gbg := _ecc.(*PSBoolean)
	if _bacd || _gbg {
		var _bgd error
		if _bacd && _gbg {
			_bgd = _dge.Push(MakeBool(_ecf.Val == _dae.Val))
		} else {
			_bgd = _dge.Push(MakeBool(false))
		}
		return _bgd
	}
	var _abee float64
	var _ged float64
	if _ede, _aeab := _fbf.(*PSInteger); _aeab {
		_abee = float64(_ede.Val)
	} else if _dbfe, _deb := _fbf.(*PSReal); _deb {
		_abee = _dbfe.Val
	} else {
		return ErrTypeCheck
	}
	if _gfeb, _gbgd := _ecc.(*PSInteger); _gbgd {
		_ged = float64(_gfeb.Val)
	} else if _cdf, _ege := _ecc.(*PSReal); _ege {
		_ged = _cdf.Val
	} else {
		return ErrTypeCheck
	}
	if _ac.Abs(_ged-_abee) < _fe {
		_egg = _dge.Push(MakeBool(true))
	} else {
		_egg = _dge.Push(MakeBool(false))
	}
	return _egg
}
func (_bdf *PSOperand) String() string { return string(*_bdf) }
func (_dbge *PSOperand) floor(_ade *PSStack) error {
	_fca, _egc := _ade.Pop()
	if _egc != nil {
		return _egc
	}
	if _bcea, _bgg := _fca.(*PSReal); _bgg {
		_egc = _ade.Push(MakeReal(_ac.Floor(_bcea.Val)))
	} else if _gfcb, _cfg := _fca.(*PSInteger); _cfg {
		_egc = _ade.Push(MakeInteger(_gfcb.Val))
	} else {
		return ErrTypeCheck
	}
	return _egc
}
func _fbbf(_dgaa int) int {
	if _dgaa < 0 {
		return -_dgaa
	}
	return _dgaa
}

// Empty empties the stack.
func (_dad *PSStack) Empty() { *_dad = []PSObject{} }

// Append appends an object to the PSProgram.
func (_ag *PSProgram) Append(obj PSObject) { *_ag = append(*_ag, obj) }
func (_aae *PSOperand) DebugString() string {
	return _baf.Sprintf("\u006fp\u003a\u0027\u0025\u0073\u0027", *_aae)
}
func (_ggcg *PSOperand) or(_aecf *PSStack) error {
	_acc, _ddbd := _aecf.Pop()
	if _ddbd != nil {
		return _ddbd
	}
	_gbdc, _ddbd := _aecf.Pop()
	if _ddbd != nil {
		return _ddbd
	}
	if _dccf, _acfg := _acc.(*PSBoolean); _acfg {
		_fbff, _cecg := _gbdc.(*PSBoolean)
		if !_cecg {
			return ErrTypeCheck
		}
		_ddbd = _aecf.Push(MakeBool(_dccf.Val || _fbff.Val))
		return _ddbd
	}
	if _cac, _baac := _acc.(*PSInteger); _baac {
		_dea, _gda := _gbdc.(*PSInteger)
		if !_gda {
			return ErrTypeCheck
		}
		_ddbd = _aecf.Push(MakeInteger(_cac.Val | _dea.Val))
		return _ddbd
	}
	return ErrTypeCheck
}
func (_bedg *PSOperand) truncate(_baad *PSStack) error {
	_dfg, _febd := _baad.Pop()
	if _febd != nil {
		return _febd
	}
	if _gac, _eceb := _dfg.(*PSReal); _eceb {
		_efd := int(_gac.Val)
		_febd = _baad.Push(MakeReal(float64(_efd)))
	} else if _ebc, _dda := _dfg.(*PSInteger); _dda {
		_febd = _baad.Push(MakeInteger(_ebc.Val))
	} else {
		return ErrTypeCheck
	}
	return _febd
}

var ErrRangeCheck = _ba.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")

func (_aac *PSOperand) gt(_fgd *PSStack) error {
	_cga, _beb := _fgd.PopNumberAsFloat64()
	if _beb != nil {
		return _beb
	}
	_aeeb, _beb := _fgd.PopNumberAsFloat64()
	if _beb != nil {
		return _beb
	}
	if _ac.Abs(_aeeb-_cga) < _fe {
		_aab := _fgd.Push(MakeBool(false))
		return _aab
	} else if _aeeb > _cga {
		_geg := _fgd.Push(MakeBool(true))
		return _geg
	} else {
		_fegd := _fgd.Push(MakeBool(false))
		return _fegd
	}
}
func (_cce *PSReal) String() string { return _baf.Sprintf("\u0025\u002e\u0035\u0066", _cce.Val) }
func (_aeb *PSOperand) add(_fea *PSStack) error {
	_ega, _ef := _fea.Pop()
	if _ef != nil {
		return _ef
	}
	_de, _ef := _fea.Pop()
	if _ef != nil {
		return _ef
	}
	_edd, _efg := _ega.(*PSReal)
	_ee, _dgf := _ega.(*PSInteger)
	if !_efg && !_dgf {
		return ErrTypeCheck
	}
	_ebe, _gaf := _de.(*PSReal)
	_gfg, _egf := _de.(*PSInteger)
	if !_gaf && !_egf {
		return ErrTypeCheck
	}
	if _dgf && _egf {
		_bde := _ee.Val + _gfg.Val
		_feg := _fea.Push(MakeInteger(_bde))
		return _feg
	}
	var _fc float64
	if _efg {
		_fc = _edd.Val
	} else {
		_fc = float64(_ee.Val)
	}
	if _gaf {
		_fc += _ebe.Val
	} else {
		_fc += float64(_gfg.Val)
	}
	_ef = _fea.Push(MakeReal(_fc))
	return _ef
}
func (_gfcee *PSOperand) roll(_dccd *PSStack) error {
	_cffe, _efge := _dccd.Pop()
	if _efge != nil {
		return _efge
	}
	_ceaa, _efge := _dccd.Pop()
	if _efge != nil {
		return _efge
	}
	_eff, _gdf := _cffe.(*PSInteger)
	if !_gdf {
		return ErrTypeCheck
	}
	_beg, _gdf := _ceaa.(*PSInteger)
	if !_gdf {
		return ErrTypeCheck
	}
	if _beg.Val < 0 {
		return ErrRangeCheck
	}
	if _beg.Val == 0 || _beg.Val == 1 {
		return nil
	}
	if _beg.Val > len(*_dccd) {
		return ErrStackUnderflow
	}
	for _fded := 0; _fded < _fbbf(_eff.Val); _fded++ {
		var _cffc []PSObject
		_cffc = (*_dccd)[len(*_dccd)-(_beg.Val) : len(*_dccd)]
		if _eff.Val > 0 {
			_gde := _cffc[len(_cffc)-1]
			_cffc = append([]PSObject{_gde}, _cffc[0:len(_cffc)-1]...)
		} else {
			_cfgb := _cffc[len(_cffc)-_beg.Val]
			_cffc = append(_cffc[1:], _cfgb)
		}
		_bag := append((*_dccd)[0:len(*_dccd)-_beg.Val], _cffc...)
		_dccd = &_bag
	}
	return nil
}

// String returns a string representation of the stack.
func (_fdc *PSStack) String() string {
	_ddae := "\u005b\u0020"
	for _, _dcf := range *_fdc {
		_ddae += _dcf.String()
		_ddae += "\u0020"
	}
	_ddae += "\u005d"
	return _ddae
}
func (_eb *PSProgram) String() string {
	_ad := "\u007b\u0020"
	for _, _bbb := range *_eb {
		_ad += _bbb.String()
		_ad += "\u0020"
	}
	_ad += "\u007d"
	return _ad
}

var ErrTypeCheck = _ba.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")

func (_eadbc *PSParser) parseFunction() (*PSProgram, error) {
	_gdag, _ := _eadbc._geca.ReadByte()
	if _gdag != '{' {
		return nil, _ba.New("\u0069\u006ev\u0061\u006c\u0069d\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
	}
	_ddf := NewPSProgram()
	for {
		_eadbc.skipSpaces()
		_gdcf, _dbe := _eadbc._geca.Peek(2)
		if _dbe != nil {
			if _dbe == _a.EOF {
				break
			}
			return nil, _dbe
		}
		_f.Log.Trace("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073", string(_gdcf))
		if _gdcf[0] == '}' {
			_f.Log.Trace("\u0045\u004f\u0046 \u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e")
			_eadbc._geca.ReadByte()
			break
		} else if _gdcf[0] == '{' {
			_f.Log.Trace("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e!")
			_bded, _fbgd := _eadbc.parseFunction()
			if _fbgd != nil {
				return nil, _fbgd
			}
			_ddf.Append(_bded)
		} else if _af.IsDecimalDigit(_gdcf[0]) || (_gdcf[0] == '-' && _af.IsDecimalDigit(_gdcf[1])) {
			_f.Log.Trace("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!")
			_bfab, _bdbf := _eadbc.parseNumber()
			if _bdbf != nil {
				return nil, _bdbf
			}
			_ddf.Append(_bfab)
		} else {
			_f.Log.Trace("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f")
			_gdcf, _ = _eadbc._geca.Peek(5)
			_cbce := string(_gdcf)
			_f.Log.Trace("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073", _cbce)
			if (len(_cbce) > 4) && (_cbce[:5] == "\u0066\u0061\u006cs\u0065") {
				_fbab, _bacde := _eadbc.parseBool()
				if _bacde != nil {
					return nil, _bacde
				}
				_ddf.Append(_fbab)
			} else if (len(_cbce) > 3) && (_cbce[:4] == "\u0074\u0072\u0075\u0065") {
				_gcd, _ced := _eadbc.parseBool()
				if _ced != nil {
					return nil, _ced
				}
				_ddf.Append(_gcd)
			} else {
				_cefb, _abcb := _eadbc.parseOperand()
				if _abcb != nil {
					return nil, _abcb
				}
				_ddf.Append(_cefb)
			}
		}
	}
	return _ddf, nil
}

// DebugString returns a descriptive string representation of the stack - intended for debugging.
func (_eag *PSStack) DebugString() string {
	_gbgg := "\u005b\u0020"
	for _, _fbgdc := range *_eag {
		_gbgg += _fbgdc.DebugString()
		_gbgg += "\u0020"
	}
	_gbgg += "\u005d"
	return _gbgg
}
func (_ab *PSInteger) Duplicate() PSObject {
	_gf := PSInteger{}
	_gf.Val = _ab.Val
	return &_gf
}
func (_bfd *PSOperand) ifCondition(_cgeg *PSStack) error {
	_bbd, _fade := _cgeg.Pop()
	if _fade != nil {
		return _fade
	}
	_bgac, _fade := _cgeg.Pop()
	if _fade != nil {
		return _fade
	}
	_eed, _dce := _bbd.(*PSProgram)
	if !_dce {
		return ErrTypeCheck
	}
	_fgb, _dce := _bgac.(*PSBoolean)
	if !_dce {
		return ErrTypeCheck
	}
	if _fgb.Val {
		_gaee := _eed.Exec(_cgeg)
		return _gaee
	}
	return nil
}

// PopNumberAsFloat64 pops and return the numeric value of the top of the stack as a float64.
// Real or integer only.
func (_dcad *PSStack) PopNumberAsFloat64() (float64, error) {
	_dfaa, _gcc := _dcad.Pop()
	if _gcc != nil {
		return 0, _gcc
	}
	if _fegg, _dbce := _dfaa.(*PSReal); _dbce {
		return _fegg.Val, nil
	} else if _fdef, _effdc := _dfaa.(*PSInteger); _effdc {
		return float64(_fdef.Val), nil
	} else {
		return 0, ErrTypeCheck
	}
}

// Exec executes the program, typically leaving output values on the stack.
func (_fde *PSProgram) Exec(stack *PSStack) error {
	for _, _dg := range *_fde {
		var _ca error
		switch _bc := _dg.(type) {
		case *PSInteger:
			_ccb := _bc
			_ca = stack.Push(_ccb)
		case *PSReal:
			_be := _bc
			_ca = stack.Push(_be)
		case *PSBoolean:
			_fed := _bc
			_ca = stack.Push(_fed)
		case *PSProgram:
			_baca := _bc
			_ca = stack.Push(_baca)
		case *PSOperand:
			_db := _bc
			_ca = _db.Exec(stack)
		default:
			return ErrTypeCheck
		}
		if _ca != nil {
			return _ca
		}
	}
	return nil
}
func (_cee *PSReal) Duplicate() PSObject {
	_ga := PSReal{}
	_ga.Val = _cee.Val
	return &_ga
}
func (_da *PSInteger) String() string { return _baf.Sprintf("\u0025\u0064", _da.Val) }
func (_fg *PSOperand) atan(_cge *PSStack) error {
	_aed, _def := _cge.PopNumberAsFloat64()
	if _def != nil {
		return _def
	}
	_bce, _def := _cge.PopNumberAsFloat64()
	if _def != nil {
		return _def
	}
	if _aed == 0 {
		var _eea error
		if _bce < 0 {
			_eea = _cge.Push(MakeReal(270))
		} else {
			_eea = _cge.Push(MakeReal(90))
		}
		return _eea
	}
	_cgc := _bce / _aed
	_fdeb := _ac.Atan(_cgc) * 180 / _ac.Pi
	_def = _cge.Push(MakeReal(_fdeb))
	return _def
}
func (_afe *PSOperand) index(_geb *PSStack) error {
	_dded, _faf := _geb.Pop()
	if _faf != nil {
		return _faf
	}
	_agf, _gedc := _dded.(*PSInteger)
	if !_gedc {
		return ErrTypeCheck
	}
	if _agf.Val < 0 {
		return ErrRangeCheck
	}
	if _agf.Val > len(*_geb)-1 {
		return ErrStackUnderflow
	}
	_dgad := (*_geb)[len(*_geb)-1-_agf.Val]
	_faf = _geb.Push(_dgad.Duplicate())
	return _faf
}
func (_aedf *PSOperand) cvi(_cdd *PSStack) error {
	_dga, _fae := _cdd.Pop()
	if _fae != nil {
		return _fae
	}
	if _gfgd, _ge := _dga.(*PSReal); _ge {
		_gbe := int(_gfgd.Val)
		_fae = _cdd.Push(MakeInteger(_gbe))
	} else if _ffc, _ddb := _dga.(*PSInteger); _ddb {
		_efb := _ffc.Val
		_fae = _cdd.Push(MakeInteger(_efb))
	} else {
		return ErrTypeCheck
	}
	return _fae
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

func (_dgfb *PSOperand) sub(_gfcg *PSStack) error {
	_gdg, _bfc := _gfcg.Pop()
	if _bfc != nil {
		return _bfc
	}
	_dag, _bfc := _gfcg.Pop()
	if _bfc != nil {
		return _bfc
	}
	_ffa, _gab := _gdg.(*PSReal)
	_gcfd, _cecb := _gdg.(*PSInteger)
	if !_gab && !_cecb {
		return ErrTypeCheck
	}
	_ggcc, _eda := _dag.(*PSReal)
	_abad, _gdga := _dag.(*PSInteger)
	if !_eda && !_gdga {
		return ErrTypeCheck
	}
	if _cecb && _gdga {
		_adeb := _abad.Val - _gcfd.Val
		_dbfd := _gfcg.Push(MakeInteger(_adeb))
		return _dbfd
	}
	var _bec float64 = 0
	if _eda {
		_bec = _ggcc.Val
	} else {
		_bec = float64(_abad.Val)
	}
	if _gab {
		_bec -= _ffa.Val
	} else {
		_bec -= float64(_gcfd.Val)
	}
	_bfc = _gfcg.Push(MakeReal(_bec))
	return _bfc
}
func (_eaaac *PSOperand) neg(_cfa *PSStack) error {
	_bdbe, _feb := _cfa.Pop()
	if _feb != nil {
		return _feb
	}
	if _gcb, _cebd := _bdbe.(*PSReal); _cebd {
		_feb = _cfa.Push(MakeReal(-_gcb.Val))
		return _feb
	} else if _fab, _fgf := _bdbe.(*PSInteger); _fgf {
		_feb = _cfa.Push(MakeInteger(-_fab.Val))
		return _feb
	} else {
		return ErrTypeCheck
	}
}

// MakeReal returns a new PSReal object initialized with `val`.
func MakeReal(val float64) *PSReal { _cfbd := PSReal{}; _cfbd.Val = val; return &_cfbd }
func (_ecb *PSOperand) log(_dca *PSStack) error {
	_caff, _adg := _dca.PopNumberAsFloat64()
	if _adg != nil {
		return _adg
	}
	_bbbe := _ac.Log10(_caff)
	_adg = _dca.Push(MakeReal(_bbbe))
	return _adg
}

// Exec executes the operand `op` in the state specified by `stack`.
func (_fbd *PSOperand) Exec(stack *PSStack) error {
	_ceba := ErrUnsupportedOperand
	switch *_fbd {
	case "\u0061\u0062\u0073":
		_ceba = _fbd.abs(stack)
	case "\u0061\u0064\u0064":
		_ceba = _fbd.add(stack)
	case "\u0061\u006e\u0064":
		_ceba = _fbd.and(stack)
	case "\u0061\u0074\u0061\u006e":
		_ceba = _fbd.atan(stack)
	case "\u0062\u0069\u0074\u0073\u0068\u0069\u0066\u0074":
		_ceba = _fbd.bitshift(stack)
	case "\u0063e\u0069\u006c\u0069\u006e\u0067":
		_ceba = _fbd.ceiling(stack)
	case "\u0063\u006f\u0070\u0079":
		_ceba = _fbd.copy(stack)
	case "\u0063\u006f\u0073":
		_ceba = _fbd.cos(stack)
	case "\u0063\u0076\u0069":
		_ceba = _fbd.cvi(stack)
	case "\u0063\u0076\u0072":
		_ceba = _fbd.cvr(stack)
	case "\u0064\u0069\u0076":
		_ceba = _fbd.div(stack)
	case "\u0064\u0075\u0070":
		_ceba = _fbd.dup(stack)
	case "\u0065\u0071":
		_ceba = _fbd.eq(stack)
	case "\u0065\u0078\u0063\u0068":
		_ceba = _fbd.exch(stack)
	case "\u0065\u0078\u0070":
		_ceba = _fbd.exp(stack)
	case "\u0066\u006c\u006fo\u0072":
		_ceba = _fbd.floor(stack)
	case "\u0067\u0065":
		_ceba = _fbd.ge(stack)
	case "\u0067\u0074":
		_ceba = _fbd.gt(stack)
	case "\u0069\u0064\u0069\u0076":
		_ceba = _fbd.idiv(stack)
	case "\u0069\u0066":
		_ceba = _fbd.ifCondition(stack)
	case "\u0069\u0066\u0065\u006c\u0073\u0065":
		_ceba = _fbd.ifelse(stack)
	case "\u0069\u006e\u0064e\u0078":
		_ceba = _fbd.index(stack)
	case "\u006c\u0065":
		_ceba = _fbd.le(stack)
	case "\u006c\u006f\u0067":
		_ceba = _fbd.log(stack)
	case "\u006c\u006e":
		_ceba = _fbd.ln(stack)
	case "\u006c\u0074":
		_ceba = _fbd.lt(stack)
	case "\u006d\u006f\u0064":
		_ceba = _fbd.mod(stack)
	case "\u006d\u0075\u006c":
		_ceba = _fbd.mul(stack)
	case "\u006e\u0065":
		_ceba = _fbd.ne(stack)
	case "\u006e\u0065\u0067":
		_ceba = _fbd.neg(stack)
	case "\u006e\u006f\u0074":
		_ceba = _fbd.not(stack)
	case "\u006f\u0072":
		_ceba = _fbd.or(stack)
	case "\u0070\u006f\u0070":
		_ceba = _fbd.pop(stack)
	case "\u0072\u006f\u0075n\u0064":
		_ceba = _fbd.round(stack)
	case "\u0072\u006f\u006c\u006c":
		_ceba = _fbd.roll(stack)
	case "\u0073\u0069\u006e":
		_ceba = _fbd.sin(stack)
	case "\u0073\u0071\u0072\u0074":
		_ceba = _fbd.sqrt(stack)
	case "\u0073\u0075\u0062":
		_ceba = _fbd.sub(stack)
	case "\u0074\u0072\u0075\u006e\u0063\u0061\u0074\u0065":
		_ceba = _fbd.truncate(stack)
	case "\u0078\u006f\u0072":
		_ceba = _fbd.xor(stack)
	}
	return _ceba
}

// NewPSParser returns a new instance of the PDF Postscript parser from input data.
func NewPSParser(content []byte) *PSParser {
	_ffd := PSParser{}
	_dcec := _g.NewBuffer(content)
	_ffd._geca = _gc.NewReader(_dcec)
	return &_ffd
}
func (_bfa *PSOperand) mul(_egac *PSStack) error {
	_gdc, _cec := _egac.Pop()
	if _cec != nil {
		return _cec
	}
	_abgg, _cec := _egac.Pop()
	if _cec != nil {
		return _cec
	}
	_babe, _dbae := _gdc.(*PSReal)
	_bede, _dfd := _gdc.(*PSInteger)
	if !_dbae && !_dfd {
		return ErrTypeCheck
	}
	_edb, _bcf := _abgg.(*PSReal)
	_eacf, _fbde := _abgg.(*PSInteger)
	if !_bcf && !_fbde {
		return ErrTypeCheck
	}
	if _dfd && _fbde {
		_bebf := _bede.Val * _eacf.Val
		_dbaa := _egac.Push(MakeInteger(_bebf))
		return _dbaa
	}
	var _cgeb float64
	if _dbae {
		_cgeb = _babe.Val
	} else {
		_cgeb = float64(_bede.Val)
	}
	if _bcf {
		_cgeb *= _edb.Val
	} else {
		_cgeb *= float64(_eacf.Val)
	}
	_cec = _egac.Push(MakeReal(_cgeb))
	return _cec
}
func (_adb *PSOperand) idiv(_dcg *PSStack) error {
	_cdgd, _eaa := _dcg.Pop()
	if _eaa != nil {
		return _eaa
	}
	_fcgb, _eaa := _dcg.Pop()
	if _eaa != nil {
		return _eaa
	}
	_cef, _eef := _cdgd.(*PSInteger)
	if !_eef {
		return ErrTypeCheck
	}
	if _cef.Val == 0 {
		return ErrUndefinedResult
	}
	_bgba, _eef := _fcgb.(*PSInteger)
	if !_eef {
		return ErrTypeCheck
	}
	_faea := _bgba.Val / _cef.Val
	_eaa = _dcg.Push(MakeInteger(_faea))
	return _eaa
}

const _fe = 0.000001

func (_gba *PSOperand) and(_fda *PSStack) error {
	_gfe, _gcf := _fda.Pop()
	if _gcf != nil {
		return _gcf
	}
	_bcg, _gcf := _fda.Pop()
	if _gcf != nil {
		return _gcf
	}
	if _df, _afc := _gfe.(*PSBoolean); _afc {
		_cbg, _efgb := _bcg.(*PSBoolean)
		if !_efgb {
			return ErrTypeCheck
		}
		_gcf = _fda.Push(MakeBool(_df.Val && _cbg.Val))
		return _gcf
	}
	if _dbfa, _dcd := _gfe.(*PSInteger); _dcd {
		_aaf, _cbe := _bcg.(*PSInteger)
		if !_cbe {
			return ErrTypeCheck
		}
		_gcf = _fda.Push(MakeInteger(_dbfa.Val & _aaf.Val))
		return _gcf
	}
	return ErrTypeCheck
}
func (_fdb *PSOperand) Duplicate() PSObject { _bf := *_fdb; return &_bf }
func (_faeac *PSOperand) not(_dcc *PSStack) error {
	_fgdc, _cdffg := _dcc.Pop()
	if _cdffg != nil {
		return _cdffg
	}
	if _eccg, _gaea := _fgdc.(*PSBoolean); _gaea {
		_cdffg = _dcc.Push(MakeBool(!_eccg.Val))
		return _cdffg
	} else if _eaf, _bea := _fgdc.(*PSInteger); _bea {
		_cdffg = _dcc.Push(MakeInteger(^_eaf.Val))
		return _cdffg
	} else {
		return ErrTypeCheck
	}
}
func (_abe *PSProgram) Duplicate() PSObject {
	_gfc := &PSProgram{}
	for _, _bac := range *_abe {
		_gfc.Append(_bac.Duplicate())
	}
	return _gfc
}
