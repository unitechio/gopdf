package extractor

import (
	_gc "bytes"
	_e "errors"
	_a "fmt"
	_cfb "image"
	_ca "image/color"
	_f "io"
	_g "math"
	_cd "reflect"
	_c "regexp"
	_cf "sort"
	_cb "strings"
	_fg "unicode"
	_fe "unicode/utf8"

	_da "unitechio/gopdf/gopdf/common"
	_bb "unitechio/gopdf/gopdf/contentstream"
	_eb "unitechio/gopdf/gopdf/core"
	_ba "unitechio/gopdf/gopdf/internal/textencoding"
	_fa "unitechio/gopdf/gopdf/internal/transform"
	_ab "unitechio/gopdf/gopdf/model"
	_b "golang.org/x/image/draw"
	_cdb "golang.org/x/text/unicode/norm"
	_fc "golang.org/x/xerrors"
)

func (_ggb *Extractor) extractPageText(_dae string, _bde *_ab.PdfPageResources, _bbcf _fa.Matrix, _fbf int) (*PageText, int, int, error) {
	_da.Log.Trace("\u0065x\u0074\u0072\u0061\u0063t\u0050\u0061\u0067\u0065\u0054e\u0078t\u003a \u006c\u0065\u0076\u0065\u006c\u003d\u0025d", _fbf)
	_aed := &PageText{_cfce: _ggb._efe, _gcaf: _ggb._egg, _fcff: _ggb._gb}
	_addd := _agde(_ggb._efe)
	var _gee stateStack
	_adb := _cegc(_ggb, _bde, _bb.GraphicsState{}, &_addd, &_gee)
	_ebfg := shapesState{_bfb: _bbcf, _cdda: _fa.IdentityMatrix(), _dfff: _adb}
	var _fbb bool
	_fec := -1
	if _fbf > _cgcb {
		_bfd := _e.New("\u0066\u006f\u0072\u006d s\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077")
		_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0065\u0078\u0074\u0072\u0061\u0063\u0074\u0050\u0061\u0067\u0065\u0054\u0065\u0078\u0074\u002e\u0020\u0072\u0065\u0063u\u0072\u0073\u0069\u006f\u006e\u0020\u006c\u0065\u0076\u0065\u006c\u003d\u0025\u0064 \u0065r\u0072\u003d\u0025\u0076", _fbf, _bfd)
		return _aed, _addd._cbee, _addd._agec, _bfd
	}
	_dafb := _bb.NewContentStreamParser(_dae)
	_bbd, _eecd := _dafb.Parse()
	if _eecd != nil {
		_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020e\u0078\u0074\u0072a\u0063\u0074\u0050\u0061g\u0065\u0054\u0065\u0078\u0074\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _eecd)
		return _aed, _addd._cbee, _addd._agec, _eecd
	}
	_aed._ffef = _bbd
	_agd := _bb.NewContentStreamProcessor(*_bbd)
	_agd.AddHandler(_bb.HandlerConditionEnumAllOperands, "", func(_befe *_bb.ContentStreamOperation, _eba _bb.GraphicsState, _age *_ab.PdfPageResources) error {
		_dga := _befe.Operand
		if _eagc {
			_da.Log.Info("\u0026&\u0026\u0020\u006f\u0070\u003d\u0025s", _befe)
		}
		switch _dga {
		case "\u0071":
			if _caad {
				_da.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _ebfg._cdda)
			}
			_gee.push(&_addd)
		case "\u0051":
			if !_gee.empty() {
				_addd = *_gee.pop()
			}
			_ebfg._cdda = _eba.CTM
			if _caad {
				_da.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _ebfg._cdda)
			}
		case "\u0042\u0044\u0043":
			_cgb, _cfec := _eb.GetDict(_befe.Params[1])
			if !_cfec {
				_da.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0042D\u0043\u0020\u006f\u0070\u003d\u0025\u0073 \u0047\u0065\u0074\u0044\u0069\u0063\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _befe)
				return _eecd
			}
			_fea := _cgb.Get("\u004d\u0043\u0049\u0044")
			if _fea != nil {
				_cabd, _dce := _eb.GetIntVal(_fea)
				if !_dce {
					_da.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0042\u0044C\u0020\u006f\u0070=\u0025\u0073\u002e\u0020\u0042\u0061\u0064\u0020\u006eum\u0065\u0072\u0069c\u0061\u006c \u006f\u0062\u006a\u0065\u0063\u0074.\u0020\u006f=\u0025\u0073", _befe, _fea)
				}
				_fec = _cabd
			} else {
				_fec = -1
			}
		case "\u0045\u004d\u0043":
			_fec = -1
		case "\u0042\u0054":
			if _fbb {
				_da.Log.Debug("\u0042\u0054\u0020\u0063\u0061\u006c\u006c\u0065\u0064\u0020\u0077\u0068\u0069\u006c\u0065 \u0069n\u0020\u0061\u0020\u0074\u0065\u0078\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
				_aed._cbf = append(_aed._cbf, _adb._cecd...)
			}
			_fbb = true
			_daef := _eba
			_daef.CTM = _bbcf.Mult(_daef.CTM)
			_adb = _cegc(_ggb, _age, _daef, &_addd, &_gee)
			_ebfg._dfff = _adb
		case "\u0045\u0054":
			if !_fbb {
				_da.Log.Debug("\u0045\u0054\u0020ca\u006c\u006c\u0065\u0064\u0020\u006f\u0075\u0074\u0073i\u0064e\u0020o\u0066 \u0061\u0020\u0074\u0065\u0078\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
			}
			_fbb = false
			_aed._cbf = append(_aed._cbf, _adb._cecd...)
			_adb.reset()
		case "\u0054\u002a":
			_adb.nextLine()
		case "\u0054\u0064":
			if _fffe, _gab := _adb.checkOp(_befe, 2, true); !_fffe {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gab)
				return _gab
			}
			_adbg, _ccg, _eed := _dbadf(_befe.Params)
			if _eed != nil {
				return _eed
			}
			_adb.moveText(_adbg, _ccg)
		case "\u0054\u0044":
			if _dbf, _fcfgc := _adb.checkOp(_befe, 2, true); !_dbf {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fcfgc)
				return _fcfgc
			}
			_acc, _fda, _gda := _dbadf(_befe.Params)
			if _gda != nil {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gda)
				return _gda
			}
			_adb.moveTextSetLeading(_acc, _fda)
		case "\u0054\u006a":
			if _beba, _fecc := _adb.checkOp(_befe, 1, true); !_beba {
				_da.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0054\u006a\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d%\u0076", _befe, _fecc)
				return _fecc
			}
			_gbg := _eb.TraceToDirectObject(_befe.Params[0])
			_ffgb, _dddd := _eb.GetStringBytes(_gbg)
			if !_dddd {
				_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020T\u006a\u0020o\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074S\u0074\u0072\u0069\u006e\u0067\u0042\u0079\u0074\u0065\u0073\u0020\u0066a\u0069\u006c\u0065\u0064", _befe)
				return _eb.ErrTypeError
			}
			return _adb.showText(_gbg, _ffgb, _fec)
		case "\u0054\u004a":
			if _eff, _abb := _adb.checkOp(_befe, 1, true); !_eff {
				_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u004a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _abb)
				return _abb
			}
			_bfg, _baa := _eb.GetArray(_befe.Params[0])
			if !_baa {
				_da.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0054\u004a\u0020\u006f\u0070\u003d\u0025s\u0020G\u0065t\u0041r\u0072\u0061\u0079\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _befe)
				return _eecd
			}
			return _adb.showTextAdjusted(_bfg, _fec)
		case "\u0027":
			if _egb, _ecad := _adb.checkOp(_befe, 1, true); !_egb {
				_da.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u0027\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ecad)
				return _ecad
			}
			_fgd := _eb.TraceToDirectObject(_befe.Params[0])
			_fde, _cbbd := _eb.GetStringBytes(_fgd)
			if !_cbbd {
				_da.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020'\u0020\u006f\u0070\u003d%s \u0047et\u0053\u0074\u0072\u0069\u006e\u0067\u0042yt\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064", _befe)
				return _eb.ErrTypeError
			}
			_adb.nextLine()
			return _adb.showText(_fgd, _fde, _fec)
		case "\u0022":
			if _ceg, _ccb := _adb.checkOp(_befe, 3, true); !_ceg {
				_da.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u0022\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ccb)
				return _ccb
			}
			_aaf, _fage, _dgfe := _dbadf(_befe.Params[:2])
			if _dgfe != nil {
				return _dgfe
			}
			_dgb := _eb.TraceToDirectObject(_befe.Params[2])
			_dfe, _eaf := _eb.GetStringBytes(_dgb)
			if !_eaf {
				_da.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020\"\u0020\u006f\u0070\u003d%s \u0047et\u0053\u0074\u0072\u0069\u006e\u0067\u0042yt\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064", _befe)
				return _eb.ErrTypeError
			}
			_adb.setCharSpacing(_aaf)
			_adb.setWordSpacing(_fage)
			_adb.nextLine()
			return _adb.showText(_dgb, _dfe, _fec)
		case "\u0054\u004c":
			_fgbb, _eef := _fcce(_befe)
			if _eef != nil {
				_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u004c\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _eef)
				return _eef
			}
			_adb.setTextLeading(_fgbb)
		case "\u0054\u0063":
			_gga, _ade := _fcce(_befe)
			if _ade != nil {
				_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0063\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ade)
				return _ade
			}
			_adb.setCharSpacing(_gga)
		case "\u0054\u0066":
			if _dcf, _ecb := _adb.checkOp(_befe, 2, true); !_dcf {
				_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0066\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ecb)
				return _ecb
			}
			_gdd, _cgg := _eb.GetNameVal(_befe.Params[0])
			if !_cgg {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _befe)
				return _eb.ErrTypeError
			}
			_aaae, _ggaf := _eb.GetNumberAsFloat(_befe.Params[1])
			if !_cgg {
				_da.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u0046\u006c\u006f\u0061\u0074\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065d\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _befe, _ggaf)
				return _ggaf
			}
			_ggaf = _adb.setFont(_gdd, _aaae)
			_adb._bfcc = _fc.Is(_ggaf, _eb.ErrNotSupported)
			if _ggaf != nil && !_adb._bfcc {
				return _ggaf
			}
		case "\u0054\u006d":
			if _ece, _gdc := _adb.checkOp(_befe, 6, true); !_ece {
				_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u006d\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gdc)
				return _gdc
			}
			_beff, _ecef := _eb.GetNumbersAsFloat(_befe.Params)
			if _ecef != nil {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ecef)
				return _ecef
			}
			_adb.setTextMatrix(_beff)
		case "\u0054\u0072":
			if _gfa, _beed := _adb.checkOp(_befe, 1, true); !_gfa {
				_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0072\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _beed)
				return _beed
			}
			_edf, _dccb := _eb.GetIntVal(_befe.Params[0])
			if !_dccb {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0072\u0020\u006f\u0070\u003d\u0025\u0073 \u0047e\u0074\u0049\u006e\u0074\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _befe)
				return _eb.ErrTypeError
			}
			_adb.setTextRenderMode(_edf)
		case "\u0054\u0073":
			if _fee, _ddg := _adb.checkOp(_befe, 1, true); !_fee {
				_da.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ddg)
				return _ddg
			}
			_gea, _gbcd := _eb.GetNumberAsFloat(_befe.Params[0])
			if _gbcd != nil {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gbcd)
				return _gbcd
			}
			_adb.setTextRise(_gea)
		case "\u0054\u0077":
			if _cdg, _fbfa := _adb.checkOp(_befe, 1, true); !_cdg {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fbfa)
				return _fbfa
			}
			_ace, _dbg := _eb.GetNumberAsFloat(_befe.Params[0])
			if _dbg != nil {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dbg)
				return _dbg
			}
			_adb.setWordSpacing(_ace)
		case "\u0054\u007a":
			if _gdda, _dbgb := _adb.checkOp(_befe, 1, true); !_gdda {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dbgb)
				return _dbgb
			}
			_dca, _fbc := _eb.GetNumberAsFloat(_befe.Params[0])
			if _fbc != nil {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fbc)
				return _fbc
			}
			_adb.setHorizScaling(_dca)
		case "\u0063\u006d":
			_ebfg._cdda = _eba.CTM
			if _ebfg._cdda.Singular() {
				_aff := _fa.IdentityMatrix().Translate(_ebfg._cdda.Translation())
				_da.Log.Debug("S\u0069n\u0067\u0075\u006c\u0061\u0072\u0020\u0063\u0074m\u003d\u0025\u0073\u2192%s", _ebfg._cdda, _aff)
				_ebfg._cdda = _aff
			}
			if _caad {
				_da.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _ebfg._cdda)
			}
		case "\u006d":
			if len(_befe.Params) != 2 {
				_da.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006d\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _ga)
				return nil
			}
			_dfg, _ffgd := _eb.GetNumbersAsFloat(_befe.Params)
			if _ffgd != nil {
				return _ffgd
			}
			_ebfg.moveTo(_dfg[0], _dfg[1])
		case "\u006c":
			if len(_befe.Params) != 2 {
				_da.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006c\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _ga)
				return nil
			}
			_fggd, _geef := _eb.GetNumbersAsFloat(_befe.Params)
			if _geef != nil {
				return _geef
			}
			_ebfg.lineTo(_fggd[0], _fggd[1])
		case "\u0063":
			if len(_befe.Params) != 6 {
				return _ga
			}
			_cgde, _ebg := _eb.GetNumbersAsFloat(_befe.Params)
			if _ebg != nil {
				return _ebg
			}
			_da.Log.Debug("\u0043u\u0062\u0069\u0063\u0020b\u0065\u007a\u0069\u0065\u0072 \u0070a\u0072a\u006d\u0073\u003a\u0020\u0025\u002e\u0032f", _cgde)
			_ebfg.cubicTo(_cgde[0], _cgde[1], _cgde[2], _cgde[3], _cgde[4], _cgde[5])
		case "\u0076", "\u0079":
			if len(_befe.Params) != 4 {
				return _ga
			}
			_ffe, _fead := _eb.GetNumbersAsFloat(_befe.Params)
			if _fead != nil {
				return _fead
			}
			_da.Log.Debug("\u0043u\u0062\u0069\u0063\u0020b\u0065\u007a\u0069\u0065\u0072 \u0070a\u0072a\u006d\u0073\u003a\u0020\u0025\u002e\u0032f", _ffe)
			_ebfg.quadraticTo(_ffe[0], _ffe[1], _ffe[2], _ffe[3])
		case "\u0068":
			_ebfg.closePath()
		case "\u0072\u0065":
			if len(_befe.Params) != 4 {
				return _ga
			}
			_ddda, _ccc := _eb.GetNumbersAsFloat(_befe.Params)
			if _ccc != nil {
				return _ccc
			}
			_ebfg.drawRectangle(_ddda[0], _ddda[1], _ddda[2], _ddda[3])
			_ebfg.closePath()
		case "\u0053":
			_ebfg.stroke(&_aed._dab)
			_ebfg.clearPath()
		case "\u0073":
			_ebfg.closePath()
			_ebfg.stroke(&_aed._dab)
			_ebfg.clearPath()
		case "\u0046":
			_ebfg.fill(&_aed._cdbf)
			_ebfg.clearPath()
		case "\u0066", "\u0066\u002a":
			_ebfg.closePath()
			_ebfg.fill(&_aed._cdbf)
			_ebfg.clearPath()
		case "\u0042", "\u0042\u002a":
			_ebfg.fill(&_aed._cdbf)
			_ebfg.stroke(&_aed._dab)
			_ebfg.clearPath()
		case "\u0062", "\u0062\u002a":
			_ebfg.closePath()
			_ebfg.fill(&_aed._cdbf)
			_ebfg.stroke(&_aed._dab)
			_ebfg.clearPath()
		case "\u006e":
			_ebfg.clearPath()
		case "\u0044\u006f":
			if len(_befe.Params) == 0 {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0058\u004fbj\u0065c\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0070\u0065\u0072\u0061n\u0064\u0020\u0066\u006f\u0072\u0020\u0044\u006f\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072.\u0020\u0047\u006f\u0074\u0020\u0025\u002b\u0076\u002e", _befe.Params)
				return _eb.ErrRangeError
			}
			_deg, _cae := _eb.GetName(_befe.Params[0])
			if !_cae {
				_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0044\u006f\u0020\u006f\u0070e\u0072a\u0074\u006f\u0072\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006fp\u0065\u0072\u0061\u006e\u0064\u003a\u0020\u0025\u002b\u0076\u002e", _befe.Params[0])
				return _eb.ErrTypeError
			}
			_, _fabc := _age.GetXObjectByName(*_deg)
			if _fabc != _ab.XObjectTypeForm {
				break
			}
			_egdd, _cae := _ggb._fag[_deg.String()]
			if !_cae {
				_eefd, _dfec := _age.GetXObjectFormByName(*_deg)
				if _dfec != nil {
					_da.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dfec)
					return _dfec
				}
				_bfcf, _dfec := _eefd.GetContentStream()
				if _dfec != nil {
					_da.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dfec)
					return _dfec
				}
				_eab := _eefd.Resources
				if _eab == nil {
					_eab = _age
				}
				_efca := _eba.CTM
				if _aab, _fcgd := _eb.GetArray(_eefd.Matrix); _fcgd {
					_ffa, _fdf := _aab.GetAsFloat64Slice()
					if _fdf != nil {
						return _fdf
					}
					if len(_ffa) != 6 {
						return _ga
					}
					_babd := _fa.NewMatrix(_ffa[0], _ffa[1], _ffa[2], _ffa[3], _ffa[4], _ffa[5])
					_efca = _eba.CTM.Mult(_babd)
				}
				_ebd, _dbff, _abf, _dfec := _ggb.extractPageText(string(_bfcf), _eab, _bbcf.Mult(_efca), _fbf+1)
				if _dfec != nil {
					_da.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _dfec)
					return _dfec
				}
				_egdd = textResult{*_ebd, _dbff, _abf}
				_ggb._fag[_deg.String()] = _egdd
			}
			_ebfg._cdda = _eba.CTM
			if _caad {
				_da.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _ebfg._cdda)
			}
			_aed._cbf = append(_aed._cbf, _egdd._ccf._cbf...)
			_aed._dab = append(_aed._dab, _egdd._ccf._dab...)
			_aed._cdbf = append(_aed._cdbf, _egdd._ccf._cdbf...)
			_addd._cbee += _egdd._aec
			_addd._agec += _egdd._efg
		case "\u0072\u0067", "\u0067", "\u006b", "\u0063\u0073", "\u0073\u0063", "\u0073\u0063\u006e":
			_adb._fcfe.ColorspaceNonStroking = _eba.ColorspaceNonStroking
			_adb._fcfe.ColorNonStroking = _eba.ColorNonStroking
		case "\u0052\u0047", "\u0047", "\u004b", "\u0043\u0053", "\u0053\u0043", "\u0053\u0043\u004e":
			_adb._fcfe.ColorspaceStroking = _eba.ColorspaceStroking
			_adb._fcfe.ColorStroking = _eba.ColorStroking
		}
		return nil
	})
	_eecd = _agd.Process(_bde)
	return _aed, _addd._cbee, _addd._agec, _eecd
}

func (_ffcc rulingList) blocks(_daec, _cgdec *ruling) bool {
	if _daec._beadd > _cgdec._dgbc || _cgdec._beadd > _daec._dgbc {
		return false
	}
	_cgcag := _g.Max(_daec._beadd, _cgdec._beadd)
	_aagbf := _g.Min(_daec._dgbc, _cgdec._dgbc)
	if _daec._dbbce > _cgdec._dbbce {
		_daec, _cgdec = _cgdec, _daec
	}
	for _, _gccag := range _ffcc {
		if _daec._dbbce <= _gccag._dbbce+_efbg && _gccag._dbbce <= _cgdec._dbbce+_efbg && _gccag._beadd <= _aagbf && _cgcag <= _gccag._dgbc {
			return true
		}
	}
	return false
}

// ImageMark represents an image drawn on a page and its position in device coordinates.
// All coordinates are in device coordinates.
type ImageMark struct {
	Image *_ab.Image

	// Dimensions of the image as displayed in the PDF.
	Width  float64
	Height float64

	// Position of the image in PDF coordinates (lower left corner).
	X float64
	Y float64

	// Angle in degrees, if rotated.
	Angle float64
}

// TableCell is a cell in a TextTable.
type TableCell struct {
	_ab.PdfRectangle

	// Text is the extracted text.
	Text string

	// Marks returns the TextMarks corresponding to the text in Text.
	Marks TextMarkArray
}

func (_ecdf *textObject) newTextMark(_ddbga string, _fcab _fa.Matrix, _aega _fa.Point, _gbcc float64, _ffed *_ab.PdfFont, _fdfb float64, _dedgg, _afccg _ca.Color, _dgff _eb.PdfObject, _beegc []string, _gfc int, _dacf int) (textMark, bool) {
	_dgca := _fcab.Angle()
	_cgbc := _ggcg(_dgca, _efgd)
	var _fdad float64
	if _cgbc%180 != 90 {
		_fdad = _fcab.ScalingFactorY()
	} else {
		_fdad = _fcab.ScalingFactorX()
	}
	_ecdb := _ecefg(_fcab)
	_cabdf := _ab.PdfRectangle{Llx: _ecdb.X, Lly: _ecdb.Y, Urx: _aega.X, Ury: _aega.Y}
	switch _cgbc % 360 {
	case 90:
		_cabdf.Urx -= _fdad
	case 180:
		_cabdf.Ury -= _fdad
	case 270:
		_cabdf.Urx += _fdad
	case 0:
		_cabdf.Ury += _fdad
	default:
		_cgbc = 0
		_cabdf.Ury += _fdad
	}
	if _cabdf.Llx > _cabdf.Urx {
		_cabdf.Llx, _cabdf.Urx = _cabdf.Urx, _cabdf.Llx
	}
	if _cabdf.Lly > _cabdf.Ury {
		_cabdf.Lly, _cabdf.Ury = _cabdf.Ury, _cabdf.Lly
	}
	_agcg := true
	if _ecdf._bdg._efe.Width() > 0 {
		_dgea, _ecbgg := _geee(_cabdf, _ecdf._bdg._efe)
		if !_ecbgg {
			_agcg = false
			_da.Log.Debug("\u0054\u0065\u0078\u0074\u0020m\u0061\u0072\u006b\u0020\u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0070a\u0067\u0065\u002e\u0020\u0062\u0062\u006f\u0078\u003d\u0025\u0067\u0020\u006d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u0067\u0020\u0074\u0065\u0078\u0074\u003d\u0025q", _cabdf, _ecdf._bdg._efe, _ddbga)
		}
		_cabdf = _dgea
	}
	_eeab := _cabdf
	_gcab := _ecdf._bdg._efe
	switch _cgbc % 360 {
	case 90:
		_gcab.Urx, _gcab.Ury = _gcab.Ury, _gcab.Urx
		_eeab = _ab.PdfRectangle{Llx: _gcab.Urx - _cabdf.Ury, Urx: _gcab.Urx - _cabdf.Lly, Lly: _cabdf.Llx, Ury: _cabdf.Urx}
	case 180:
		_eeab = _ab.PdfRectangle{Llx: _gcab.Urx - _cabdf.Llx, Urx: _gcab.Urx - _cabdf.Urx, Lly: _gcab.Ury - _cabdf.Lly, Ury: _gcab.Ury - _cabdf.Ury}
	case 270:
		_gcab.Urx, _gcab.Ury = _gcab.Ury, _gcab.Urx
		_eeab = _ab.PdfRectangle{Llx: _cabdf.Ury, Urx: _cabdf.Lly, Lly: _gcab.Ury - _cabdf.Llx, Ury: _gcab.Ury - _cabdf.Urx}
	}
	if _eeab.Llx > _eeab.Urx {
		_eeab.Llx, _eeab.Urx = _eeab.Urx, _eeab.Llx
	}
	if _eeab.Lly > _eeab.Ury {
		_eeab.Lly, _eeab.Ury = _eeab.Ury, _eeab.Lly
	}
	_dagbb := textMark{_bbeb: _ddbga, PdfRectangle: _eeab, _fdea: _cabdf, _fdbdc: _ffed, _acaa: _fdad, _febab: _fdfb, _addb: _fcab, _ggef: _aega, _fbfge: _cgbc, _aegfa: _dedgg, _acbg: _afccg, _ggdb: _dgff, _cbfa: _beegc, Th: _ecdf._dbfd._fgc, Tw: _ecdf._dbfd._ebbfc, _gbdg: _dacf, _cbgg: _gfc}
	if _fbbg {
		_da.Log.Info("n\u0065\u0077\u0054\u0065\u0078\u0074M\u0061\u0072\u006b\u003a\u0020\u0073t\u0061\u0072\u0074\u003d\u0025\u002e\u0032f\u0020\u0065\u006e\u0064\u003d\u0025\u002e\u0032\u0066\u0020%\u0073", _ecdb, _aega, _dagbb.String())
	}
	return _dagbb, _agcg
}

func (_debe *compositeCell) updateBBox() {
	for _, _dbfc := range _debe.paraList {
		_debe.PdfRectangle = _abbb(_debe.PdfRectangle, _dbfc.PdfRectangle)
	}
}

type textLine struct {
	_ab.PdfRectangle
	_defc float64
	_gfed []*textWord
	_ead  float64
}

func _cega(_fcac, _efbc bounded) float64 {
	_bgag := _bbaa(_fcac, _efbc)
	if !_ccae(_bgag) {
		return _bgag
	}
	return _faeff(_fcac, _efbc)
}

func (_dgbf *textObject) checkOp(_ccd *_bb.ContentStreamOperation, _ccfe int, _egf bool) (_gcbf bool, _dbfa error) {
	if _dgbf == nil {
		var _efeb []_eb.PdfObject
		if _ccfe > 0 {
			_efeb = _ccd.Params
			if len(_efeb) > _ccfe {
				_efeb = _efeb[:_ccfe]
			}
		}
		_da.Log.Debug("\u0025\u0023q \u006f\u0070\u0065r\u0061\u006e\u0064\u0020out\u0073id\u0065\u0020\u0074\u0065\u0078\u0074\u002e p\u0061\u0072\u0061\u006d\u0073\u003d\u0025+\u0076", _ccd.Operand, _efeb)
	}
	if _ccfe >= 0 {
		if len(_ccd.Params) != _ccfe {
			if _egf {
				_dbfa = _e.New("\u0069n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0070\u0061r\u0061m\u0065t\u0065\u0072\u0020\u0063\u006f\u0075\u006et")
			}
			_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u0023\u0071\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020h\u0061\u0076\u0065\u0020\u0025\u0064\u0020i\u006e\u0070\u0075\u0074\u0020\u0070\u0061\u0072\u0061\u006d\u0073,\u0020\u0067\u006f\u0074\u0020\u0025\u0064\u0020\u0025\u002b\u0076", _ccd.Operand, _ccfe, len(_ccd.Params), _ccd.Params)
			return false, _dbfa
		}
	}
	return true, nil
}
func _dbdg(_gbdc, _gcbae float64) bool { return _gbdc/_g.Max(_addcd, _gcbae) < _cfggb }
func (_edbf *textObject) setFont(_bbae string, _ddde float64) error {
	if _edbf == nil {
		return nil
	}
	_edbf._dbfd._egbc = _ddde
	_gag, _deb := _edbf.getFont(_bbae)
	if _deb != nil {
		return _deb
	}
	_edbf._dbfd._ggda = _gag
	return nil
}

// Text gets the extracted text contained in `l`.
func (_gbaa *list) Text() string {
	_fddgd := &_cb.Builder{}
	_adecg := ""
	_affc(_gbaa, _fddgd, &_adecg)
	return _fddgd.String()
}

func _defbb(_fade []*textLine, _cdae string, _dcfe []*list) *list {
	return &list{_dbfeb: _fade, _dbga: _cdae, _ffaf: _dcfe}
}

// ToTextMark returns the public view of `tm`.
func (_bbebc *textMark) ToTextMark() TextMark {
	return TextMark{Text: _bbebc._bbeb, Original: _bbebc._cage, BBox: _bbebc._fdea, Font: _bbebc._fdbdc, FontSize: _bbebc._acaa, FillColor: _bbebc._aegfa, StrokeColor: _bbebc._acbg, Orientation: _bbebc._fbfge, DirectObject: _bbebc._ggdb, ObjString: _bbebc._cbfa, Tw: _bbebc.Tw, Th: _bbebc.Th, Tc: _bbebc._febab, Index: _bbebc._cbgg}
}

func (_bedb *textTable) reduceTiling(_ccfdc gridTiling, _fadab float64) *textTable {
	_facc := make([]int, 0, _bedb._faed)
	_dgdgb := make([]int, 0, _bedb._cbcb)
	_efcd := _ccfdc._cgeba
	_ebadb := _ccfdc._gfdac
	for _beage := 0; _beage < _bedb._faed; _beage++ {
		_fefd := _beage > 0 && _g.Abs(_ebadb[_beage-1]-_ebadb[_beage]) < _fadab && _bedb.emptyCompositeRow(_beage)
		if !_fefd {
			_facc = append(_facc, _beage)
		}
	}
	for _ebceg := 0; _ebceg < _bedb._cbcb; _ebceg++ {
		_aecge := _ebceg < _bedb._cbcb-1 && _g.Abs(_efcd[_ebceg+1]-_efcd[_ebceg]) < _fadab && _bedb.emptyCompositeColumn(_ebceg)
		if !_aecge {
			_dgdgb = append(_dgdgb, _ebceg)
		}
	}
	if len(_facc) == _bedb._faed && len(_dgdgb) == _bedb._cbcb {
		return _bedb
	}
	_dgcac := textTable{_cafe: _bedb._cafe, _cbcb: len(_dgdgb), _faed: len(_facc), _fcda: make(map[uint64]compositeCell, len(_dgdgb)*len(_facc))}
	if _cabbg {
		_da.Log.Info("\u0072\u0065\u0064\u0075c\u0065\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0025d\u0078%\u0064\u0020\u002d\u003e\u0020\u0025\u0064x\u0025\u0064", _bedb._cbcb, _bedb._faed, len(_dgdgb), len(_facc))
		_da.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0043\u006f\u006c\u0073\u003a\u0020\u0025\u002b\u0076", _dgdgb)
		_da.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0052\u006f\u0077\u0073\u003a\u0020\u0025\u002b\u0076", _facc)
	}
	for _afef, _edff := range _facc {
		for _bccgf, _gccde := range _dgdgb {
			_fgdg, _fdge := _bedb.getComposite(_gccde, _edff)
			if len(_fgdg) == 0 {
				continue
			}
			if _cabbg {
				_a.Printf("\u0020 \u0025\u0032\u0064\u002c \u0025\u0032\u0064\u0020\u0028%\u0032d\u002c \u0025\u0032\u0064\u0029\u0020\u0025\u0071\n", _bccgf, _afef, _gccde, _edff, _gggg(_fgdg.merge().text(), 50))
			}
			_dgcac.putComposite(_bccgf, _afef, _fgdg, _fdge)
		}
	}
	return &_dgcac
}

const _cgcb = 20

func _agecb(_dfdg []*textLine, _dddb map[float64][]*textLine) []*list {
	_acd := _ebede(_dddb)
	_adgbg := []*list{}
	if len(_acd) == 0 {
		return _adgbg
	}
	_ecgcac := _acd[0]
	_bdgd := 1
	_bfbb := _dddb[_ecgcac]
	for _gdfg, _cccf := range _bfbb {
		var _gfbg float64
		_aebg := []*list{}
		_dbad := _cccf._defc
		_cfbb := -1.0
		if _gdfg < len(_bfbb)-1 {
			_cfbb = _bfbb[_gdfg+1]._defc
		}
		if _bdgd < len(_acd) {
			_aebg = _eaea(_dfdg, _dddb, _acd, _bdgd, _dbad, _cfbb)
		}
		_gfbg = _cfbb
		if len(_aebg) > 0 {
			_bddd := _aebg[0]
			if len(_bddd._dbfeb) > 0 {
				_gfbg = _bddd._dbfeb[0]._defc
			}
		}
		_egfd := []*textLine{_cccf}
		_adba := _dcfd(_cccf, _dfdg, _acd, _dbad, _gfbg)
		_egfd = append(_egfd, _adba...)
		_eccf := _defbb(_egfd, "\u0062\u0075\u006c\u006c\u0065\u0074", _aebg)
		_eccf._gfec = _aeag(_egfd, "")
		_adgbg = append(_adgbg, _eccf)
	}
	return _adgbg
}

func _fbbae(_adee, _faeg _fa.Point) rulingKind {
	_bebc := _g.Abs(_adee.X - _faeg.X)
	_ffca := _g.Abs(_adee.Y - _faeg.Y)
	return _cddg(_bebc, _ffca, _addc)
}

func _bfcd(_bfbe *wordBag, _deab *textWord, _fgeg float64) bool {
	return _bfbe.Urx <= _deab.Llx && _deab.Llx < _bfbe.Urx+_fgeg
}

// String returns a description of `v`.
func (_cced *ruling) String() string {
	if _cced._ggcaf == _fbag {
		return "\u004e\u004f\u0054\u0020\u0052\u0055\u004c\u0049\u004e\u0047"
	}
	_bbfd, _aae := "\u0078", "\u0079"
	if _cced._ggcaf == _ceedg {
		_bbfd, _aae = "\u0079", "\u0078"
	}
	_gaba := ""
	if _cced._fbfc != 0.0 {
		_gaba = _a.Sprintf(" \u0077\u0069\u0064\u0074\u0068\u003d\u0025\u002e\u0032\u0066", _cced._fbfc)
	}
	return _a.Sprintf("\u0025\u00310\u0073\u0020\u0025\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u0020\u002d\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0028\u0025\u0036\u002e\u0032\u0066\u0029\u0020\u0025\u0073\u0020\u0025\u0076\u0025\u0073", _cced._ggcaf, _bbfd, _cced._dbbce, _aae, _cced._beadd, _cced._dgbc, _cced._dgbc-_cced._beadd, _cced._bfccg, _cced.Color, _gaba)
}

func _fcacf(_bdaa string) bool {
	for _, _cgfa := range _bdaa {
		if !_fg.IsSpace(_cgfa) {
			return false
		}
	}
	return true
}

func _ccbcd(_ffgc structElement) []structElement {
	_afcc := []structElement{}
	for _, _dedg := range _ffgc._baac {
		for _, _eebd := range _dedg._baac {
			for _, _cffdc := range _eebd._baac {
				if _cffdc._acfd == "\u004c" {
					_afcc = append(_afcc, _cffdc)
				}
			}
		}
	}
	return _afcc
}

func (_dfebg *structTreeRoot) buildList(_ccbcf map[int][]*textLine, _eegda _eb.PdfObject) []*list {
	if _dfebg == nil {
		_da.Log.Debug("\u0062\u0075\u0069\u006c\u0064\u004c\u0069\u0073\u0074\u003a\u0020t\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0069\u0073 \u006e\u0069\u006c")
		return nil
	}
	var _fcd *structElement
	_adbgf := []structElement{}
	if len(_dfebg._eede) == 1 {
		_caeg := _dfebg._eede[0]._acfd
		if _caeg == "\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074" || _caeg == "\u0053\u0065\u0063\u0074" || _caeg == "\u0050\u0061\u0072\u0074" || _caeg == "\u0044\u0069\u0076" || _caeg == "\u0041\u0072\u0074" {
			_fcd = &_dfebg._eede[0]
		}
	} else {
		_fcd = &structElement{_baac: _dfebg._eede, _acfd: _dfebg._aabe}
	}
	if _fcd == nil {
		_da.Log.Debug("\u0062\u0075\u0069\u006cd\u004c\u0069\u0073\u0074\u003a\u0020\u0074\u006f\u0070\u0045l\u0065m\u0065\u006e\u0074\u0020\u0069\u0073\u0020n\u0069\u006c")
		return nil
	}
	for _, _cegf := range _fcd._baac {
		if _cegf._acfd == "\u004c" {
			_adbgf = append(_adbgf, _cegf)
		} else if _cegf._acfd == "\u0054\u0061\u0062l\u0065" {
			_fedbf := _ccbcd(_cegf)
			_adbgf = append(_adbgf, _fedbf...)
		}
	}
	_cccd := _acad(_adbgf, _ccbcf, _eegda)
	var _eefb []*list
	for _, _aced := range _cccd {
		_dbbb := _dfcg(_aced)
		_eefb = append(_eefb, _dbbb...)
	}
	return _eefb
}

func (_aacb *textTable) toTextTable() TextTable {
	if _cabbg {
		_da.Log.Info("t\u006fT\u0065\u0078\u0074\u0054\u0061\u0062\u006c\u0065:\u0020\u0025\u0064\u0020x \u0025\u0064", _aacb._cbcb, _aacb._faed)
	}
	_fdde := make([][]TableCell, _aacb._faed)
	for _dfac := 0; _dfac < _aacb._faed; _dfac++ {
		_fdde[_dfac] = make([]TableCell, _aacb._cbcb)
		for _eefg := 0; _eefg < _aacb._cbcb; _eefg++ {
			_gacc := _aacb.get(_eefg, _dfac)
			if _gacc == nil {
				continue
			}
			if _cabbg {
				_a.Printf("\u0025\u0034\u0064 \u0025\u0032\u0064\u003a\u0020\u0025\u0073\u000a", _eefg, _dfac, _gacc)
			}
			_fdde[_dfac][_eefg].Text = _gacc.text()
			_gcfac := 0
			_fdde[_dfac][_eefg].Marks._beaf = _gacc.toTextMarks(&_gcfac)
		}
	}
	_fcgde := TextTable{W: _aacb._cbcb, H: _aacb._faed, Cells: _fdde}
	_fcgde.PdfRectangle = _aacb.bbox()
	return _fcgde
}

// PageTextOptions holds various options available in extraction process.
type PageTextOptions struct {
	_bgd   bool
	_dccbb bool
}

func (_gbf *textLine) pullWord(_efag *wordBag, _gdbg *textWord, _gdfb int) {
	_gbf.appendWord(_gdbg)
	_efag.removeWord(_gdbg, _gdfb)
}

func _fcce(_dfef *_bb.ContentStreamOperation) (float64, error) {
	if len(_dfef.Params) != 1 {
		_fdce := _e.New("\u0069n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0070\u0061r\u0061m\u0065t\u0065\u0072\u0020\u0063\u006f\u0075\u006et")
		_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u0023\u0071\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020h\u0061\u0076\u0065\u0020\u0025\u0064\u0020i\u006e\u0070\u0075\u0074\u0020\u0070\u0061\u0072\u0061\u006d\u0073,\u0020\u0067\u006f\u0074\u0020\u0025\u0064\u0020\u0025\u002b\u0076", _dfef.Operand, 1, len(_dfef.Params), _dfef.Params)
		return 0.0, _fdce
	}
	return _eb.GetNumberAsFloat(_dfef.Params[0])
}

func (_ecee *wordBag) absorb(_cabe *wordBag) {
	_geg := _cabe.makeRemovals()
	for _gcca, _cfgg := range _cabe._edgf {
		for _, _eafc := range _cfgg {
			_ecee.pullWord(_eafc, _gcca, _geg)
		}
	}
	_cabe.applyRemovals(_geg)
}

func (_ebaae *PageText) computeViews() {
	_fad := _ebaae.getParagraphs()
	_fbaa := new(_gc.Buffer)
	_fad.writeText(_fbaa)
	_ebaae._aeb = _fbaa.String()
	_ebaae._gddg = _fad.toTextMarks()
	_ebaae._aee = _fad.tables()
	if _cabbg {
		_da.Log.Info("\u0063\u006f\u006dpu\u0074\u0065\u0056\u0069\u0065\u0077\u0073\u003a\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u003d\u0025\u0064", len(_ebaae._aee))
	}
}

type textPara struct {
	_ab.PdfRectangle
	_gedc  _ab.PdfRectangle
	_fcbfe []*textLine
	_bcfcf *textTable
	_aagd  bool
	_ccfeb bool
	_fdccb *textPara
	_eefcd *textPara
	_aebbe *textPara
	_afeeg *textPara
	_bbda  []list
}

func (_bgdg *textLine) markWordBoundaries() {
	_gbgc := _gccd * _bgdg._ead
	for _ddff, _fafea := range _bgdg._gfed[1:] {
		if _cfba(_fafea, _bgdg._gfed[_ddff]) >= _gbgc {
			_fafea._bace = true
		}
	}
}

func _bbfc(_dbca []rulingList) (rulingList, rulingList) {
	var _gfca rulingList
	for _, _agcge := range _dbca {
		_gfca = append(_gfca, _agcge...)
	}
	return _gfca.vertsHorzs()
}

func (_fgfeb *TextMarkArray) exists(_fga TextMark) bool {
	for _, _cfeb := range _fgfeb.Elements() {
		if _cd.DeepEqual(_fga.DirectObject, _cfeb.DirectObject) && _cd.DeepEqual(_fga.BBox, _cfeb.BBox) && _cfeb.Text == _fga.Text {
			return true
		}
	}
	return false
}

func _cebb(_dcdd _ab.PdfRectangle) *ruling {
	return &ruling{_ggcaf: _eccce, _dbbce: _dcdd.Llx, _beadd: _dcdd.Lly, _dgbc: _dcdd.Ury}
}

// String returns a string describing `tm`.
func (_edbd TextMark) String() string {
	_efb := _edbd.BBox
	var _gdbd string
	if _edbd.Font != nil {
		_gdbd = _edbd.Font.String()
		if len(_gdbd) > 50 {
			_gdbd = _gdbd[:50] + "\u002e\u002e\u002e"
		}
	}
	var _fcgb string
	if _edbd.Meta {
		_fcgb = "\u0020\u002a\u004d\u002a"
	}
	return _a.Sprintf("\u007b\u0054\u0065\u0078t\u004d\u0061\u0072\u006b\u003a\u0020\u0025\u0064\u0020%\u0071\u003d\u0025\u0030\u0032\u0078\u0020\u0028\u0025\u0036\u002e\u0032\u0066\u002c\u0020\u0025\u0036\u002e2\u0066\u0029\u0020\u0028\u00256\u002e\u0032\u0066\u002c\u0020\u0025\u0036\u002e\u0032\u0066\u0029\u0020\u0025\u0073\u0025\u0073\u007d", _edbd.Offset, _edbd.Text, []rune(_edbd.Text), _efb.Llx, _efb.Lly, _efb.Urx, _efb.Ury, _gdbd, _fcgb)
}

func (_fcag paraList) llyRange(_bcda []int, _dgda, _bbagc float64) []int {
	_bbcff := len(_fcag)
	if _bbagc < _fcag[_bcda[0]].Lly || _dgda > _fcag[_bcda[_bbcff-1]].Lly {
		return nil
	}
	_edcb := _cf.Search(_bbcff, func(_bcbb int) bool { return _fcag[_bcda[_bcbb]].Lly >= _dgda })
	_adgc := _cf.Search(_bbcff, func(_gfdaf int) bool { return _fcag[_bcda[_gfdaf]].Lly > _bbagc })
	return _bcda[_edcb:_adgc]
}

func (_ccgb *textObject) getFont(_dbc string) (*_ab.PdfFont, error) {
	if _ccgb._bdg._eg != nil {
		_ecce, _fbca := _ccgb.getFontDict(_dbc)
		if _fbca != nil {
			_da.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0067\u0065\u0074\u0046\u006f\u006e\u0074:\u0020n\u0061m\u0065=\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073", _dbc, _fbca.Error())
			return nil, _fbca
		}
		_ccgb._bdg._bc++
		_dfgf, _dcb := _ccgb._bdg._eg[_ecce.String()]
		if _dcb {
			_dfgf._dba = _ccgb._bdg._bc
			return _dfgf._daba, nil
		}
	}
	_cbbg, _abdc := _ccgb.getFontDict(_dbc)
	if _abdc != nil {
		return nil, _abdc
	}
	_afga, _abdc := _ccgb.getFontDirect(_dbc)
	if _abdc != nil {
		return nil, _abdc
	}
	if _ccgb._bdg._eg != nil {
		_bdfd := fontEntry{_afga, _ccgb._bdg._bc}
		if len(_ccgb._bdg._eg) >= _fgbf {
			var _dedb []string
			for _fadf := range _ccgb._bdg._eg {
				_dedb = append(_dedb, _fadf)
			}
			_cf.Slice(_dedb, func(_dfde, _fbbc int) bool {
				return _ccgb._bdg._eg[_dedb[_dfde]]._dba < _ccgb._bdg._eg[_dedb[_fbbc]]._dba
			})
			delete(_ccgb._bdg._eg, _dedb[0])
		}
		_ccgb._bdg._eg[_cbbg.String()] = _bdfd
	}
	return _afga, nil
}
func _cbc(_abc _fa.Point) *subpath { return &subpath{_bfaa: []_fa.Point{_abc}} }
func _agde(_bead _ab.PdfRectangle) textState {
	return textState{_fgc: 100, _cdf: RenderModeFill, _cee: _bead}
}

func (_edacge *textTable) put(_fagd, _adbc int, _gcge *textPara) {
	_edacge._ebgbd[_fadg(_fagd, _adbc)] = _gcge
}

func (_fcccb *textPara) toTextMarks(_gbgag *int) []TextMark {
	if _fcccb._bcfcf == nil {
		return _fcccb.toCellTextMarks(_gbgag)
	}
	var _aeba []TextMark
	for _fddd := 0; _fddd < _fcccb._bcfcf._faed; _fddd++ {
		for _aeddd := 0; _aeddd < _fcccb._bcfcf._cbcb; _aeddd++ {
			_adcc := _fcccb._bcfcf.get(_aeddd, _fddd)
			if _adcc == nil {
				_aeba = _beec(_aeba, _gbgag, "\u0009")
			} else {
				_gfga := _adcc.toCellTextMarks(_gbgag)
				_aeba = append(_aeba, _gfga...)
			}
			_aeba = _beec(_aeba, _gbgag, "\u0020")
		}
		if _fddd < _fcccb._bcfcf._faed-1 {
			_aeba = _beec(_aeba, _gbgag, "\u000a")
		}
	}
	_agdeg := _fcccb._bcfcf
	if _agdeg.isExportable() {
		_aecg := _agdeg.toTextTable()
		_aeba = _cfdab(_aeba, &_aecg)
	}
	return _aeba
}

// String returns a description of `t`.
func (_dbac *textTable) String() string {
	return _a.Sprintf("\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u0020\u0025\u0074", _dbac._cbcb, _dbac._faed, _dbac._cafe)
}
func (_aaaee *textTable) bbox() _ab.PdfRectangle { return _aaaee.PdfRectangle }

type lists []*list

func _ebdd(_fgfa _ab.PdfRectangle) rulingKind {
	_gfcf := _fgfa.Width()
	_aebge := _fgfa.Height()
	if _gfcf > _aebge {
		if _gfcf >= _addc {
			return _ceedg
		}
	} else {
		if _aebge >= _addc {
			return _eccce
		}
	}
	return _fbag
}

func (_gccb *wordBag) getDepthIdx(_cfea float64) int {
	_bagd := _gccb.depthIndexes()
	_edgg := _aeef(_cfea)
	if _edgg < _bagd[0] {
		return _bagd[0]
	}
	if _edgg > _bagd[len(_bagd)-1] {
		return _bagd[len(_bagd)-1]
	}
	return _edgg
}

func (_ebce *wordBag) makeRemovals() map[int]map[*textWord]struct{} {
	_fggc := make(map[int]map[*textWord]struct{}, len(_ebce._edgf))
	for _dfffa := range _ebce._edgf {
		_fggc[_dfffa] = make(map[*textWord]struct{})
	}
	return _fggc
}

func (_ceeg *wordBag) removeWord(_abcg *textWord, _egdf int) {
	_effa := _ceeg._edgf[_egdf]
	_effa = _facad(_effa, _abcg)
	if len(_effa) == 0 {
		delete(_ceeg._edgf, _egdf)
	} else {
		_ceeg._edgf[_egdf] = _effa
	}
}

func (_acef compositeCell) parasBBox() (paraList, _ab.PdfRectangle) {
	return _acef.paraList, _acef.PdfRectangle
}

func (_dd *PageFonts) extractPageResourcesToFont(_faa *_ab.PdfPageResources) error {
	_ge, _fgfb := _eb.GetDict(_faa.Font)
	if !_fgfb {
		return _e.New(_ff)
	}
	for _, _fgb := range _ge.Keys() {
		var (
			_gac = true
			_aa  []byte
			_fgg string
		)
		_gef, _ffc := _faa.GetFontByName(_fgb)
		if !_ffc {
			return _e.New(_eggg)
		}
		_egd, _bec := _ab.NewPdfFontFromPdfObject(_gef)
		if _bec != nil {
			return _bec
		}
		_efd := _egd.FontDescriptor()
		_cbb := _egd.FontDescriptor().FontName.String()
		_bea := _egd.Subtype()
		if _efdd(_dd.Fonts, _cbb) {
			continue
		}
		if len(_egd.ToUnicode()) == 0 {
			_gac = false
		}
		if _efd.FontFile != nil {
			if _efc, _cffd := _eb.GetStream(_efd.FontFile); _cffd {
				_aa, _bec = _eb.DecodeStream(_efc)
				if _bec != nil {
					return _bec
				}
				_fgg = _cbb + "\u002e\u0070\u0066\u0062"
			}
		} else if _efd.FontFile2 != nil {
			if _bab, _fcg := _eb.GetStream(_efd.FontFile2); _fcg {
				_aa, _bec = _eb.DecodeStream(_bab)
				if _bec != nil {
					return _bec
				}
				_fgg = _cbb + "\u002e\u0074\u0074\u0066"
			}
		} else if _efd.FontFile3 != nil {
			if _cde, _db := _eb.GetStream(_efd.FontFile3); _db {
				_aa, _bec = _eb.DecodeStream(_cde)
				if _bec != nil {
					return _bec
				}
				_fgg = _cbb + "\u002e\u0063\u0066\u0066"
			}
		}
		if len(_fgg) < 1 {
			_da.Log.Debug(_ec)
		}
		_ag := Font{FontName: _cbb, PdfFont: _egd, IsCID: _egd.IsCID(), IsSimple: _egd.IsSimple(), ToUnicode: _gac, FontType: _bea, FontData: _aa, FontFileName: _fgg, FontDescriptor: _efd}
		_dd.Fonts = append(_dd.Fonts, _ag)
	}
	return nil
}

// ImageExtractOptions contains options for controlling image extraction from
// PDF pages.
type ImageExtractOptions struct{ IncludeInlineStencilMasks bool }

func (_egbd *textObject) setTextRise(_eced float64) {
	if _egbd == nil {
		return
	}
	_egbd._dbfd._ded = _eced
}

func _eefc(_ffaa []*textMark, _cebe _ab.PdfRectangle, _dafg rulingList, _edac []gridTiling, _gcea bool) paraList {
	_da.Log.Trace("\u006d\u0061\u006b\u0065\u0054\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u003a \u0025\u0064\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0073\u0020\u0070\u0061\u0067\u0065\u0053\u0069\u007a\u0065=\u0025\u002e\u0032\u0066", len(_ffaa), _cebe)
	if len(_ffaa) == 0 {
		return nil
	}
	_egcb := _eabcb(_ffaa, _cebe)
	if len(_egcb) == 0 {
		return nil
	}
	_dafg.log("\u006d\u0061\u006be\u0054\u0065\u0078\u0074\u0050\u0061\u0067\u0065")
	_ddgd, _bfde := _dafg.vertsHorzs()
	_aeea := _dbaf(_egcb, _cebe.Ury, _ddgd, _bfde)
	_abe := _affa(_aeea, _cebe.Ury, _ddgd, _bfde)
	_abe = _cgdc(_abe)
	_cdfae := make(paraList, 0, len(_abe))
	for _, _dcff := range _abe {
		_edfe := _dcff.arrangeText()
		if _edfe != nil {
			_cdfae = append(_cdfae, _edfe)
		}
	}
	if !_gcea && len(_cdfae) >= _afdc {
		_cdfae = _cdfae.extractTables(_edac)
	}
	_cdfae.sortReadingOrder()
	if !_gcea {
		_cdfae.sortTopoOrder()
	}
	_cdfae.log("\u0073\u006f\u0072te\u0064\u0020\u0069\u006e\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0064\u0065\u0072")
	return _cdfae
}

func (_agge *wordBag) highestWord(_cggg int, _ege, _ggca float64) *textWord {
	for _, _feadb := range _agge._edgf[_cggg] {
		if _ege <= _feadb._faace && _feadb._faace <= _ggca {
			return _feadb
		}
	}
	return nil
}

// Options extractor options.
type Options struct {
	// DisableDocumentTags specifies whether to use the document tags during list extraction.
	DisableDocumentTags bool

	// ApplyCropBox will extract page text based on page cropbox if set to `true`.
	ApplyCropBox bool

	// UseSimplerExtractionProcess will skip topological text ordering and table processing.
	//
	// NOTE: While normally the extra processing is beneficial, it can also lead to problems when it does not work.
	// Thus it is a flag to allow the user to control this process.
	//
	// Skipping some extraction processes would also lead to the reduced processing time.
	UseSimplerExtractionProcess bool
}

func (_dccbe *textLine) toTextMarks(_ccef *int) []TextMark {
	var _agcdf []TextMark
	for _, _fedf := range _dccbe._gfed {
		if _fedf._bace {
			_agcdf = _beec(_agcdf, _ccef, "\u0020")
		}
		_cgdcc := _fedf.toTextMarks(_ccef)
		_agcdf = append(_agcdf, _cgdcc...)
	}
	return _agcdf
}
func (_fcb *subpath) add(_agfd ..._fa.Point) { _fcb._bfaa = append(_fcb._bfaa, _agfd...) }
func _dgag(_dbcd float64) float64            { return _dddg * _g.Round(_dbcd/_dddg) }
func (_fcgee *textTable) get(_ecaea, _bgced int) *textPara {
	return _fcgee._ebgbd[_fadg(_ecaea, _bgced)]
}

func (_eaeg *wordBag) firstReadingIndex(_edbb int) int {
	_ged := _eaeg.firstWord(_edbb)._fcde
	_bfaf := float64(_edbb+1) * _aegf
	_efdgg := _bfaf + _dbab*_ged
	_bbdg := _edbb
	for _, _egge := range _eaeg.depthBand(_bfaf, _efdgg) {
		if _faeff(_eaeg.firstWord(_egge), _eaeg.firstWord(_bbdg)) < 0 {
			_bbdg = _egge
		}
	}
	return _bbdg
}

func (_bfcea *ruling) encloses(_bfcdb, _eece float64) bool {
	return _bfcea._beadd-_gded <= _bfcdb && _eece <= _bfcea._dgbc+_gded
}

func _aabg(_efee byte) bool {
	for _, _aaba := range _ebda {
		if []byte(_aaba)[0] == _efee {
			return true
		}
	}
	return false
}

func (_bagc rulingList) toGrids() []rulingList {
	if _egc {
		_da.Log.Info("t\u006f\u0047\u0072\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _bagc)
	}
	_baae := _bagc.intersections()
	if _egc {
		_da.Log.Info("\u0074\u006f\u0047r\u0069\u0064\u0073\u003a \u0076\u0065\u0063\u0073\u003d\u0025\u0064 \u0069\u006e\u0074\u0065\u0072\u0073\u0065\u0063\u0074\u0073\u003d\u0025\u0064\u0020", len(_bagc), len(_baae))
		for _, _ceaf := range _dcbe(_baae) {
			_a.Printf("\u00254\u0064\u003a\u0020\u0025\u002b\u0076\n", _ceaf, _baae[_ceaf])
		}
	}
	_dbcge := make(map[int]intSet, len(_bagc))
	for _babc := range _bagc {
		_adgba := _bagc.connections(_baae, _babc)
		if len(_adgba) > 0 {
			_dbcge[_babc] = _adgba
		}
	}
	if _egc {
		_da.Log.Info("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u0063\u006fn\u006e\u0065\u0063\u0074s=\u0025\u0064", len(_dbcge))
		for _, _gbccb := range _dcbe(_dbcge) {
			_a.Printf("\u00254\u0064\u003a\u0020\u0025\u002b\u0076\n", _gbccb, _dbcge[_gbccb])
		}
	}
	_ebcc := _cdcc(len(_bagc), func(_bgee, _gdbe int) bool {
		_efef, _bgbg := len(_dbcge[_bgee]), len(_dbcge[_gdbe])
		if _efef != _bgbg {
			return _efef > _bgbg
		}
		return _bagc.comp(_bgee, _gdbe)
	})
	if _egc {
		_da.Log.Info("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u006f\u0072d\u0065\u0072\u0069\u006eg=\u0025\u0076", _ebcc)
	}
	_fgad := [][]int{{_ebcc[0]}}
_afac:
	for _, _agfe := range _ebcc[1:] {
		for _bbdee, _efefb := range _fgad {
			for _, _ggcab := range _efefb {
				if _dbcge[_ggcab].has(_agfe) {
					_fgad[_bbdee] = append(_efefb, _agfe)
					continue _afac
				}
			}
		}
		_fgad = append(_fgad, []int{_agfe})
	}
	if _egc {
		_da.Log.Info("\u0074o\u0047r\u0069\u0064\u0073\u003a\u0020i\u0067\u0072i\u0064\u0073\u003d\u0025\u0076", _fgad)
	}
	_cf.SliceStable(_fgad, func(_gdadf, _cdcg int) bool { return len(_fgad[_gdadf]) > len(_fgad[_cdcg]) })
	for _, _dbea := range _fgad {
		_cf.Slice(_dbea, func(_ddgdc, _bgagg int) bool { return _bagc.comp(_dbea[_ddgdc], _dbea[_bgagg]) })
	}
	_cgea := make([]rulingList, len(_fgad))
	for _gfge, _agcgb := range _fgad {
		_aaed := make(rulingList, len(_agcgb))
		for _dcee, _ffac := range _agcgb {
			_aaed[_dcee] = _bagc[_ffac]
		}
		_cgea[_gfge] = _aaed
	}
	if _egc {
		_da.Log.Info("\u0074o\u0047r\u0069\u0064\u0073\u003a\u0020g\u0072\u0069d\u0073\u003d\u0025\u002b\u0076", _cgea)
	}
	var _ecab []rulingList
	for _, _cceff := range _cgea {
		if _afaf, _ggec := _cceff.isActualGrid(); _ggec {
			_cceff = _afaf
			_cceff = _cceff.snapToGroups()
			_ecab = append(_ecab, _cceff)
		}
	}
	if _egc {
		_fggf("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u0061\u0063t\u0075\u0061\u006c\u0047ri\u0064\u0073", _ecab)
		_da.Log.Info("\u0074\u006f\u0047\u0072\u0069\u0064\u0073\u003a\u0020\u0067\u0072\u0069\u0064\u0073\u003d%\u0064 \u0061\u0063\u0074\u0075\u0061\u006c\u0047\u0072\u0069\u0064\u0073\u003d\u0025\u0064", len(_cgea), len(_ecab))
	}
	return _ecab
}

func _gcbdd(_fded, _bbeaa int) int {
	if _fded < _bbeaa {
		return _fded
	}
	return _bbeaa
}

func (_gddf rulingList) sortStrict() {
	_cf.Slice(_gddf, func(_fdeb, _bceee int) bool {
		_deee, _degf := _gddf[_fdeb], _gddf[_bceee]
		_ebba, _gfcce := _deee._ggcaf, _degf._ggcaf
		if _ebba != _gfcce {
			return _ebba > _gfcce
		}
		_dbaff, _ccfg := _deee._dbbce, _degf._dbbce
		if !_ccae(_dbaff - _ccfg) {
			return _dbaff < _ccfg
		}
		_dbaff, _ccfg = _deee._beadd, _degf._beadd
		if _dbaff != _ccfg {
			return _dbaff < _ccfg
		}
		return _deee._dgbc < _degf._dgbc
	})
}

func _ccdb(_adcga _ab.PdfRectangle) *ruling {
	return &ruling{_ggcaf: _ceedg, _dbbce: _adcga.Lly, _beadd: _adcga.Llx, _dgbc: _adcga.Urx}
}

func _abbb(_bffb, _efdf _ab.PdfRectangle) _ab.PdfRectangle {
	return _ab.PdfRectangle{Llx: _g.Min(_bffb.Llx, _efdf.Llx), Lly: _g.Min(_bffb.Lly, _efdf.Lly), Urx: _g.Max(_bffb.Urx, _efdf.Urx), Ury: _g.Max(_bffb.Ury, _efdf.Ury)}
}
func (_ddcde intSet) add(_gbafa int) { _ddcde[_gbafa] = struct{}{} }

// Tables returns the tables extracted from the page.
func (_cgeb PageText) Tables() []TextTable {
	if _cabbg {
		_da.Log.Info("\u0054\u0061\u0062\u006c\u0065\u0073\u003a\u0020\u0025\u0064", len(_cgeb._aee))
	}
	return _cgeb._aee
}

func (_bgdf *textWord) computeText() string {
	_fadaba := make([]string, len(_bgdf._bacfd))
	for _gabfc, _cffce := range _bgdf._bacfd {
		_fadaba[_gabfc] = _cffce._bbeb
	}
	return _cb.Join(_fadaba, "")
}

type rulingList []*ruling

// New returns an Extractor instance for extracting content from the input PDF page.
func New(page *_ab.PdfPage) (*Extractor, error) { return NewWithOptions(page, nil) }
func (_adf gridTile) complete() bool            { return _adf.numBorders() == 4 }
func _bfee(_ccfgg _ab.PdfRectangle, _dgfg, _cdedg, _afcg, _edgcc *ruling) gridTile {
	_gcef := _ccfgg.Llx
	_fecg := _ccfgg.Urx
	_fddab := _ccfgg.Lly
	_cagd := _ccfgg.Ury
	return gridTile{PdfRectangle: _ccfgg, _feaa: _dgfg != nil && _dgfg.encloses(_fddab, _cagd), _gcde: _cdedg != nil && _cdedg.encloses(_fddab, _cagd), _bdec: _afcg != nil && _afcg.encloses(_gcef, _fecg), _dbgf: _edgcc != nil && _edgcc.encloses(_gcef, _fecg)}
}

type gridTiling struct {
	_ab.PdfRectangle
	_cgeba []float64
	_gfdac []float64
	_afce  map[float64]map[float64]gridTile
}

func (_agaa *textTable) computeBbox() _ab.PdfRectangle {
	var _ffege _ab.PdfRectangle
	_addbc := false
	for _gdgd := 0; _gdgd < _agaa._faed; _gdgd++ {
		for _cbfdg := 0; _cbfdg < _agaa._cbcb; _cbfdg++ {
			_eaab := _agaa.get(_cbfdg, _gdgd)
			if _eaab == nil {
				continue
			}
			if !_addbc {
				_ffege = _eaab.PdfRectangle
				_addbc = true
			} else {
				_ffege = _abbb(_ffege, _eaab.PdfRectangle)
			}
		}
	}
	return _ffege
}

func (_efgdg *textTable) putComposite(_ddca, _egbge int, _geag paraList, _bbgd _ab.PdfRectangle) {
	if len(_geag) == 0 {
		_da.Log.Error("\u0074\u0065xt\u0054\u0061\u0062l\u0065\u0029\u0020\u0070utC\u006fmp\u006f\u0073\u0069\u0074\u0065\u003a\u0020em\u0070\u0074\u0079\u0020\u0070\u0061\u0072a\u0073")
		return
	}
	_gabaa := compositeCell{PdfRectangle: _bbgd, paraList: _geag}
	if _cabbg {
		_a.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0070\u0075\u0074\u0043\u006f\u006d\u0070o\u0073i\u0074\u0065\u0028\u0025\u0064\u002c\u0025\u0064\u0029\u003c\u002d\u0025\u0073\u000a", _ddca, _egbge, _gabaa.String())
	}
	_gabaa.updateBBox()
	_efgdg._fcda[_fadg(_ddca, _egbge)] = _gabaa
}

func (_gcdd *shapesState) newSubPath() {
	_gcdd.clearPath()
	if _caad {
		_da.Log.Info("\u006e\u0065\u0077\u0053\u0075\u0062\u0050\u0061\u0074h\u003a\u0020\u0025\u0073", _gcdd)
	}
}

func _dcaf(_acge []pathSection) rulingList {
	_bbede(_acge)
	if _egc {
		_da.Log.Info("\u006d\u0061k\u0065\u0053\u0074\u0072\u006f\u006b\u0065\u0052\u0075\u006c\u0069\u006e\u0067\u0073\u003a\u0020\u0025\u0064\u0020\u0073\u0074\u0072ok\u0065\u0073", len(_acge))
	}
	var _babbd rulingList
	for _, _cbfg := range _acge {
		for _, _debec := range _cbfg._egfg {
			if len(_debec._bfaa) < 2 {
				continue
			}
			_bgeb := _debec._bfaa[0]
			for _, _dggd := range _debec._bfaa[1:] {
				if _baea, _feca := _gbccg(_bgeb, _dggd, _cbfg.Color); _feca {
					_babbd = append(_babbd, _baea)
				}
				_bgeb = _dggd
			}
		}
	}
	if _egc {
		_da.Log.Info("m\u0061\u006b\u0065\u0053tr\u006fk\u0065\u0052\u0075\u006c\u0069n\u0067\u0073\u003a\u0020\u0025\u0073", _babbd)
	}
	return _babbd
}

func (_dec pathSection) bbox() _ab.PdfRectangle {
	_faeb := _dec._egfg[0]._bfaa[0]
	_fdaa := _ab.PdfRectangle{Llx: _faeb.X, Urx: _faeb.X, Lly: _faeb.Y, Ury: _faeb.Y}
	_dggf := func(_dgad _fa.Point) {
		if _dgad.X < _fdaa.Llx {
			_fdaa.Llx = _dgad.X
		} else if _dgad.X > _fdaa.Urx {
			_fdaa.Urx = _dgad.X
		}
		if _dgad.Y < _fdaa.Lly {
			_fdaa.Lly = _dgad.Y
		} else if _dgad.Y > _fdaa.Ury {
			_fdaa.Ury = _dgad.Y
		}
	}
	for _, _eedf := range _dec._egfg[0]._bfaa[1:] {
		_dggf(_eedf)
	}
	for _, _cefd := range _dec._egfg[1:] {
		for _, _adc := range _cefd._bfaa {
			_dggf(_adc)
		}
	}
	return _fdaa
}

func (_gcg paraList) findTables(_gdbbd []gridTiling) []*textTable {
	_gcg.addNeighbours()
	_cf.Slice(_gcg, func(_dgfa, _feab int) bool { return _beadb(_gcg[_dgfa], _gcg[_feab]) < 0 })
	var _fgae []*textTable
	if _dbbc {
		_aaee := _gcg.findGridTables(_gdbbd)
		_fgae = append(_fgae, _aaee...)
	}
	if _gcbdb {
		_bfccc := _gcg.findTextTables()
		_fgae = append(_fgae, _bfccc...)
	}
	return _fgae
}

func (_agefa TextTable) getCellInfo(_fcffb TextMark) [][]int {
	for _fbbd, _geeg := range _agefa.Cells {
		for _bcde, _bbgf := range _geeg {
			_cabg := &_bbgf.Marks
			if _cabg.exists(_fcffb) {
				return [][]int{{_fbbd}, {_bcde}}
			}
		}
	}
	return nil
}

// ToText returns the page text as a single string.
// Deprecated: This function is deprecated and will be removed in a future major version. Please use
// Text() instead.
func (_bbeg PageText) ToText() string { return _bbeg.Text() }

func _aebba(_faab *textLine, _aaga []*textLine, _gabe []float64) float64 {
	var _bfdd float64 = -1
	for _, _dddgd := range _aaga {
		if _dddgd._defc > _faab._defc {
			if _g.Round(_dddgd.Llx) >= _g.Round(_faab.Llx) {
				_bfdd = _dddgd._defc
			} else {
				break
			}
		}
	}
	return _bfdd
}

func (_gbge *shapesState) fill(_egbg *[]pathSection) {
	_cad := pathSection{_egfg: _gbge._dagf, Color: _gbge._dfff.getFillColor()}
	*_egbg = append(*_egbg, _cad)
	if _egc {
		_egbda := _cad.bbox()
		_a.Printf("\u0020 \u0020\u0020\u0046\u0049\u004c\u004c\u003a %\u0032\u0064\u0020\u0066\u0069\u006c\u006c\u0073\u0020\u0028\u0025\u0064\u0020\u006ee\u0077\u0029 \u0073\u0073\u003d%\u0073\u0020\u0063\u006f\u006c\u006f\u0072\u003d\u0025\u0033\u0076\u0020\u0025\u0036\u002e\u0032f\u003d\u00256.\u0032\u0066\u0078%\u0036\u002e\u0032\u0066\u000a", len(*_egbg), len(_cad._egfg), _gbge, _cad.Color, _egbda, _egbda.Width(), _egbda.Height())
		if _gcbad {
			for _cafg, _bcag := range _cad._egfg {
				_a.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _cafg, _bcag)
				if _cafg == 10 {
					break
				}
			}
		}
	}
}

func (_ggcf paraList) tables() []TextTable {
	var _cdfe []TextTable
	if _cabbg {
		_da.Log.Info("\u0070\u0061\u0072\u0061\u0073\u002e\u0074\u0061\u0062\u006c\u0065\u0073\u003a")
	}
	for _, _ffcg := range _ggcf {
		_dgafg := _ffcg._bcfcf
		if _dgafg != nil && _dgafg.isExportable() {
			_cdfe = append(_cdfe, _dgafg.toTextTable())
		}
	}
	return _cdfe
}

func _geee(_deacb, _debda _ab.PdfRectangle) (_ab.PdfRectangle, bool) {
	if !_gega(_deacb, _debda) {
		return _ab.PdfRectangle{}, false
	}
	return _ab.PdfRectangle{Llx: _g.Max(_deacb.Llx, _debda.Llx), Urx: _g.Min(_deacb.Urx, _debda.Urx), Lly: _g.Max(_deacb.Lly, _debda.Lly), Ury: _g.Min(_deacb.Ury, _debda.Ury)}, true
}

func _egfcc(_agda []TextMark, _ceda *int, _gefde TextMark) []TextMark {
	_gefde.Offset = *_ceda
	_agda = append(_agda, _gefde)
	*_ceda += len(_gefde.Text)
	return _agda
}

func (_acafd rulingList) augmentGrid() (rulingList, rulingList) {
	_afgc, _edbe := _acafd.vertsHorzs()
	if len(_afgc) == 0 || len(_edbe) == 0 {
		return _afgc, _edbe
	}
	_gceae, _gfaa := _afgc, _edbe
	_bcbef := _afgc.bbox()
	_fadaa := _edbe.bbox()
	if _egc {
		_da.Log.Info("\u0061u\u0067\u006d\u0065\u006e\u0074\u0047\u0072\u0069\u0064\u003a\u0020b\u0062\u006f\u0078\u0056\u003d\u0025\u0036\u002e\u0032\u0066", _bcbef)
		_da.Log.Info("\u0061u\u0067\u006d\u0065\u006e\u0074\u0047\u0072\u0069\u0064\u003a\u0020b\u0062\u006f\u0078\u0048\u003d\u0025\u0036\u002e\u0032\u0066", _fadaa)
	}
	var _efcgd, _cdfec, _fbeag, _bgfad *ruling
	if _fadaa.Llx < _bcbef.Llx-_gded {
		_efcgd = &ruling{_bfccg: _agbf, _ggcaf: _eccce, _dbbce: _fadaa.Llx, _beadd: _bcbef.Lly, _dgbc: _bcbef.Ury}
		_afgc = append(rulingList{_efcgd}, _afgc...)
	}
	if _fadaa.Urx > _bcbef.Urx+_gded {
		_cdfec = &ruling{_bfccg: _agbf, _ggcaf: _eccce, _dbbce: _fadaa.Urx, _beadd: _bcbef.Lly, _dgbc: _bcbef.Ury}
		_afgc = append(_afgc, _cdfec)
	}
	if _bcbef.Lly < _fadaa.Lly-_gded {
		_fbeag = &ruling{_bfccg: _agbf, _ggcaf: _ceedg, _dbbce: _bcbef.Lly, _beadd: _fadaa.Llx, _dgbc: _fadaa.Urx}
		_edbe = append(rulingList{_fbeag}, _edbe...)
	}
	if _bcbef.Ury > _fadaa.Ury+_gded {
		_bgfad = &ruling{_bfccg: _agbf, _ggcaf: _ceedg, _dbbce: _bcbef.Ury, _beadd: _fadaa.Llx, _dgbc: _fadaa.Urx}
		_edbe = append(_edbe, _bgfad)
	}
	if len(_afgc)+len(_edbe) == len(_acafd) {
		return _gceae, _gfaa
	}
	_gbag := append(_afgc, _edbe...)
	_acafd.log("u\u006e\u0061\u0075\u0067\u006d\u0065\u006e\u0074\u0065\u0064")
	_gbag.log("\u0061u\u0067\u006d\u0065\u006e\u0074\u0065d")
	return _afgc, _edbe
}

// String returns a description of `k`.
func (_cggff rulingKind) String() string {
	_geaff, _cafb := _ddbc[_cggff]
	if !_cafb {
		return _a.Sprintf("\u004e\u006ft\u0020\u0061\u0020r\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0025\u0064", _cggff)
	}
	return _geaff
}

func (_bfgc *wordBag) empty(_gbgaf int) bool {
	_, _cbfb := _bfgc._edgf[_gbgaf]
	return !_cbfb
}

func _bbede(_bedbd []pathSection) {
	if _dddg < 0.0 {
		return
	}
	if _egc {
		_da.Log.Info("\u0067\u0072\u0061\u006e\u0075\u006c\u0061\u0072\u0069\u007a\u0065\u003a\u0020\u0025\u0064 \u0073u\u0062\u0070\u0061\u0074\u0068\u0020\u0073\u0065\u0063\u0074\u0069\u006f\u006e\u0073", len(_bedbd))
	}
	for _fcgcec, _dbdba := range _bedbd {
		for _ccgcg, _ebaed := range _dbdba._egfg {
			for _cffcg, _eagfa := range _ebaed._bfaa {
				_ebaed._bfaa[_cffcg] = _fa.Point{X: _dgag(_eagfa.X), Y: _dgag(_eagfa.Y)}
				if _egc {
					_fgga := _ebaed._bfaa[_cffcg]
					if !_gagg(_eagfa, _fgga) {
						_dabg := _fa.Point{X: _fgga.X - _eagfa.X, Y: _fgga.Y - _eagfa.Y}
						_a.Printf("\u0025\u0034d \u002d\u0020\u00254\u0064\u0020\u002d\u0020%4d\u003a %\u002e\u0032\u0066\u0020\u2192\u0020\u0025.2\u0066\u0020\u0028\u0025\u0067\u0029\u000a", _fcgcec, _ccgcg, _cffcg, _eagfa, _fgga, _dabg)
					}
				}
			}
		}
	}
}

func (_gcf *textObject) setTextLeading(_dgfb float64) {
	if _gcf == nil {
		return
	}
	_gcf._dbfd._abda = _dgfb
}

func (_ddgc rulingList) snapToGroups() rulingList {
	_dceg, _egbb := _ddgc.vertsHorzs()
	if len(_dceg) > 0 {
		_dceg = _dceg.snapToGroupsDirection()
	}
	if len(_egbb) > 0 {
		_egbb = _egbb.snapToGroupsDirection()
	}
	_bbgfa := append(_dceg, _egbb...)
	_bbgfa.log("\u0073\u006e\u0061p\u0054\u006f\u0047\u0072\u006f\u0075\u0070\u0073")
	return _bbgfa
}
func (_begf *subpath) last() _fa.Point { return _begf._bfaa[len(_begf._bfaa)-1] }

// ExtractFonts returns all font information from the page extractor, including
// font name, font type, the raw data of the embedded font file (if embedded), font descriptor and more.
//
// The argument `previousPageFonts` is used when trying to build a complete font catalog for multiple pages or the entire document.
// The entries from `previousPageFonts` are added to the returned result unless already included in the page, i.e. no duplicate entries.
//
// NOTE: If previousPageFonts is nil, all fonts from the page will be returned. Use it when building up a full list of fonts for a document or page range.
func (_cbe *Extractor) ExtractFonts(previousPageFonts *PageFonts) (*PageFonts, error) {
	_bbef := PageFonts{}
	_be := _bbef.extractPageResourcesToFont(_cbe._fgf)
	if _be != nil {
		return nil, _be
	}
	if previousPageFonts != nil {
		for _, _eca := range previousPageFonts.Fonts {
			if !_efdd(_bbef.Fonts, _eca.FontName) {
				_bbef.Fonts = append(_bbef.Fonts, _eca)
			}
		}
	}
	return &PageFonts{Fonts: _bbef.Fonts}, nil
}

func (_dddga *textTable) getDown() paraList {
	_eged := make(paraList, _dddga._cbcb)
	for _cecg := 0; _cecg < _dddga._cbcb; _cecg++ {
		_cceea := _dddga.get(_cecg, _dddga._faed-1)._afeeg
		if _cceea.taken() {
			return nil
		}
		_eged[_cecg] = _cceea
	}
	for _cgddg := 0; _cgddg < _dddga._cbcb-1; _cgddg++ {
		if _eged[_cgddg]._eefcd != _eged[_cgddg+1] {
			return nil
		}
	}
	return _eged
}
func (_dggb *wordBag) firstWord(_ceb int) *textWord { return _dggb._edgf[_ceb][0] }
func _abfe(_egca *list, _cbeg *string) string {
	_cded := _cb.Split(_egca._gfec, "\u000a")
	_fbe := &_cb.Builder{}
	for _, _eagcd := range _cded {
		if _eagcd != "" {
			_fbe.WriteString(*_cbeg)
			_fbe.WriteString(_eagcd)
			_fbe.WriteString("\u000a")
		}
	}
	return _fbe.String()
}
func (_bdbc paraList) sortTopoOrder()     { _cfcg := _bdbc.topoOrder(); _bdbc.reorder(_cfcg) }
func (_dffb *textPara) fontsize() float64 { return _dffb._fcbfe[0]._ead }

// ExtractPageText returns the text contents of `e` (an Extractor for a page) as a PageText.
// TODO(peterwilliams97): The stats complicate this function signature and aren't very useful.
//
//	Replace with a function like Extract() (*PageText, error)
func (_faca *Extractor) ExtractPageText() (*PageText, int, int, error) {
	_adda, _bba, _bbc, _cfe := _faca.extractPageText(_faca._ef, _faca._fgf, _fa.IdentityMatrix(), 0)
	if _cfe != nil && _cfe != _ab.ErrColorOutOfRange {
		return nil, 0, 0, _cfe
	}
	if _faca._bbe != nil {
		_adda._cbef._dccbb = _faca._bbe.UseSimplerExtractionProcess
	}
	_adda.computeViews()
	if _faca._bbe != nil {
		if _faca._bbe.ApplyCropBox && _faca._fcf != nil {
			_adda.ApplyArea(*_faca._fcf)
		}
		_adda._cbef._bgd = _faca._bbe.DisableDocumentTags
	}
	return _adda, _bba, _bbc, nil
}

func (_aeee *wordBag) allWords() []*textWord {
	var _gfae []*textWord
	for _, _cbdf := range _aeee._edgf {
		_gfae = append(_gfae, _cbdf...)
	}
	return _gfae
}

var _cfda string = "\u005e\u005b\u0061\u002d\u007a\u0041\u002dZ\u005d\u0028\u005c)\u007c\u005c\u002e)\u007c\u005e[\u005c\u0064\u005d\u002b\u0028\u005c)\u007c\\.\u0029\u007c\u005e\u005c\u0028\u005b\u0061\u002d\u007a\u0041\u002d\u005a\u005d\u005c\u0029\u007c\u005e\u005c\u0028\u005b\u005c\u0064\u005d\u002b\u005c\u0029"

func (_bdbb paraList) list() []*list {
	var _eabb []*textLine
	var _gddc []*textLine
	for _, _bdeg := range _bdbb {
		_bcfcb := _bdeg.getListLines()
		_eabb = append(_eabb, _bcfcb...)
		_gddc = append(_gddc, _bdeg._fcbfe...)
	}
	_aceg := _ecfd(_eabb)
	_ggcb := _agecb(_gddc, _aceg)
	return _ggcb
}

const _fgbf = 10

func (_gecf *shapesState) establishSubpath() *subpath {
	_dddf, _eaba := _gecf.lastpointEstablished()
	if !_eaba {
		_gecf._dagf = append(_gecf._dagf, _cbc(_dddf))
	}
	if len(_gecf._dagf) == 0 {
		return nil
	}
	_gecf._caef = false
	return _gecf._dagf[len(_gecf._dagf)-1]
}

func _cfdab(_eeaf []TextMark, _eccc *TextTable) []TextMark {
	var _ddbgf []TextMark
	for _, _acee := range _eeaf {
		_acee._fddc = true
		_acee._fabfe = _eccc
		_ddbgf = append(_ddbgf, _acee)
	}
	return _ddbgf
}

// TextMarkArray is a collection of TextMarks.
type (
	TextMarkArray struct{ _beaf []TextMark }
	textMark      struct {
		_ab.PdfRectangle
		_fbfge int
		_bbeb  string
		_cage  string
		_fdbdc *_ab.PdfFont
		_acaa  float64
		_febab float64
		_addb  _fa.Matrix
		_ggef  _fa.Point
		_fdea  _ab.PdfRectangle
		_aegfa _ca.Color
		_acbg  _ca.Color
		_ggdb  _eb.PdfObject
		_cbfa  []string
		Tw     float64
		Th     float64
		_gbdg  int
		_cbgg  int
	}
)

func (_fced *shapesState) devicePoint(_aebb, _aefgg float64) _fa.Point {
	_gfbfa := _fced._bfb.Mult(_fced._cdda)
	_aebb, _aefgg = _gfbfa.Transform(_aebb, _aefgg)
	return _fa.NewPoint(_aebb, _aefgg)
}

func (_fega *textTable) compositeRowCorridors() map[int][]float64 {
	_egdc := make(map[int][]float64, _fega._faed)
	if _cabbg {
		_da.Log.Info("c\u006f\u006d\u0070\u006f\u0073\u0069t\u0065\u0052\u006f\u0077\u0043\u006f\u0072\u0072\u0069d\u006f\u0072\u0073:\u0020h\u003d\u0025\u0064", _fega._faed)
	}
	for _dfdc := 1; _dfdc < _fega._faed; _dfdc++ {
		var _afedc []compositeCell
		for _eaece := 0; _eaece < _fega._cbcb; _eaece++ {
			if _fafde, _dbeb := _fega._fcda[_fadg(_eaece, _dfdc)]; _dbeb {
				_afedc = append(_afedc, _fafde)
			}
		}
		if len(_afedc) == 0 {
			continue
		}
		_cfabg := _ede(_afedc)
		_egdc[_dfdc] = _cfabg
		if _cabbg {
			_a.Printf("\u0020\u0020\u0020\u0025\u0032\u0064\u003a\u0020\u00256\u002e\u0032\u0066\u000a", _dfdc, _cfabg)
		}
	}
	return _egdc
}

func _deff(_dgfbe *paraList) map[int][]*textLine {
	_fbcc := map[int][]*textLine{}
	for _, _gada := range *_dgfbe {
		for _, _eecc := range _gada._fcbfe {
			if !_bbcdg(_eecc) {
				_da.Log.Debug("g\u0072\u006f\u0075p\u004c\u0069\u006e\u0065\u0073\u003a\u0020\u0054\u0068\u0065\u0020\u0074\u0065\u0078\u0074\u0020\u006c\u0069\u006e\u0065\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0073 \u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u006e\u0065 \u006d\u0063\u0069\u0064 \u006e\u0075\u006d\u0062e\u0072\u002e\u0020\u0049\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0073p\u006c\u0069\u0074\u002e")
				continue
			}
			_cdge := _eecc._gfed[0]._bacfd[0]._gbdg
			_fbcc[_cdge] = append(_fbcc[_cdge], _eecc)
		}
		if _gada._bcfcf != nil {
			_aabaa := _gada._bcfcf._ebgbd
			for _, _daaec := range _aabaa {
				for _, _cgf := range _daaec._fcbfe {
					if !_bbcdg(_cgf) {
						_da.Log.Debug("g\u0072\u006f\u0075p\u004c\u0069\u006e\u0065\u0073\u003a\u0020\u0054\u0068\u0065\u0020\u0074\u0065\u0078\u0074\u0020\u006c\u0069\u006e\u0065\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0073 \u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u006e\u0065 \u006d\u0063\u0069\u0064 \u006e\u0075\u006d\u0062e\u0072\u002e\u0020\u0049\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0073p\u006c\u0069\u0074\u002e")
						continue
					}
					_aeec := _cgf._gfed[0]._bacfd[0]._gbdg
					_fbcc[_aeec] = append(_fbcc[_aeec], _cgf)
				}
			}
		}
	}
	return _fbcc
}

func (_bbca *ruling) alignsSec(_gggdg *ruling) bool {
	const _dbafb = _efbg + 1.0
	return _bbca._beadd-_dbafb <= _gggdg._dgbc && _gggdg._beadd-_dbafb <= _bbca._dgbc
}

type shapesState struct {
	_cdda _fa.Matrix
	_bfb  _fa.Matrix
	_dagf []*subpath
	_caef bool
	_beaa _fa.Point
	_dfff *textObject
}

func (_afcbfg paraList) lines() []*textLine {
	var _cfbc []*textLine
	for _, _dbfbf := range _afcbfg {
		_cfbc = append(_cfbc, _dbfbf._fcbfe...)
	}
	return _cfbc
}

func _dcfd(_gcfe *textLine, _beedd []*textLine, _dfce []float64, _feee, _afca float64) []*textLine {
	_ccdc := []*textLine{}
	for _, _geabb := range _beedd {
		if _geabb._defc >= _feee {
			if _afca != -1 && _geabb._defc < _afca {
				if _geabb.text() != _gcfe.text() {
					if _g.Round(_geabb.Llx) < _g.Round(_gcfe.Llx) {
						break
					}
					_ccdc = append(_ccdc, _geabb)
				}
			} else if _afca == -1 {
				if _geabb._defc == _gcfe._defc {
					if _geabb.text() != _gcfe.text() {
						_ccdc = append(_ccdc, _geabb)
					}
					continue
				}
				_eadc := _aebba(_gcfe, _beedd, _dfce)
				if _eadc != -1 && _geabb._defc <= _eadc {
					_ccdc = append(_ccdc, _geabb)
				}
			}
		}
	}
	return _ccdc
}

func (_gaad *textObject) getStrokeColor() _ca.Color {
	return _eafe(_gaad._fcfe.ColorspaceStroking, _gaad._fcfe.ColorStroking)
}

// List returns all the list objects detected on the page.
// It detects all the bullet point Lists from a given pdf page and builds a slice of bullet list objects.
// A given bullet list object has a tree structure.
// Each bullet point list is extracted with the text content it contains and all the sub lists found under it as children in the tree.
// The rest content of the pdf is ignored and only text in the bullet point lists are extracted.
// The list extraction is done in two ways.
// 1. If the document is tagged then the lists are extracted using the tags provided in the document.
// 2. Otherwise the bullet lists are extracted from the raw text using regex matching.
// By default the document tag is used if available.
// However this can be disabled using `DisableDocumentTags` in the `Options` object.
// Sometimes disabling document tags option might give a better bullet list extraction if the document was tagged incorrectly.
//
//	    options := &Options{
//		     DisableDocumentTags: false, // this means use document tag if available
//	    }
//	    ex, err := NewWithOptions(page, options)
//	    // handle error
//	    pageText, _, _, err := ex.ExtractPageText()
//	    // handle error
//	    lists := pageText.List()
//	    txt := lists.Text()
func (_bdbe PageText) List() lists {
	_gdae := !_bdbe._cbef._bgd
	_deabe := _bdbe.getParagraphs()
	_gfac := true
	if _bdbe._gcaf == nil || *_bdbe._gcaf == nil {
		_gfac = false
	}
	_bbec := _deabe.list()
	if _gfac && _gdae {
		_aead := _deff(&_deabe)
		_abbd := &structTreeRoot{}
		_abbd.parseStructTreeRoot(*_bdbe._gcaf)
		if _abbd._eede == nil {
			_da.Log.Debug("\u004c\u0069\u0073\u0074\u003a\u0020\u0073t\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0064\u006f\u0065\u0073\u006e'\u0074\u0020\u0068\u0061\u0076e\u0020\u0061\u006e\u0079\u0020\u0063\u006f\u006e\u0074e\u006e\u0074\u002c\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0074\u0065\u0078\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0069\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u0020\u0069\u006e\u0073\u0074\u0065\u0061\u0064\u002e")
			return _bbec
		}
		_bbec = _abbd.buildList(_aead, _bdbe._fcff)
	}
	return _bbec
}

func (_gcce *wordBag) pullWord(_fge *textWord, _fdb int, _cdca map[int]map[*textWord]struct{}) {
	_gcce.PdfRectangle = _abbb(_gcce.PdfRectangle, _fge.PdfRectangle)
	if _fge._fcde > _gcce._caag {
		_gcce._caag = _fge._fcde
	}
	_gcce._edgf[_fdb] = append(_gcce._edgf[_fdb], _fge)
	_cdca[_fdb][_fge] = struct{}{}
}

func (_acdb rulingList) removeDuplicates() rulingList {
	if len(_acdb) == 0 {
		return nil
	}
	_acdb.sort()
	_cegg := rulingList{_acdb[0]}
	for _, _afeg := range _acdb[1:] {
		if _afeg.equals(_cegg[len(_cegg)-1]) {
			continue
		}
		_cegg = append(_cegg, _afeg)
	}
	return _cegg
}

func (_bdgc *ruling) gridIntersecting(_cac *ruling) bool {
	return _ccfac(_bdgc._beadd, _cac._beadd) && _ccfac(_bdgc._dgbc, _cac._dgbc)
}

// String returns a description of `l`.
func (_bfeb *textLine) String() string {
	return _a.Sprintf("\u0025\u002e2\u0066\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\"%\u0073\u0022", _bfeb._defc, _bfeb.PdfRectangle, _bfeb._ead, _bfeb.text())
}

func (_efcadf *textTable) depth() float64 {
	_cbbae := 1e10
	for _fbbf := 0; _fbbf < _efcadf._cbcb; _fbbf++ {
		_begge := _efcadf.get(_fbbf, 0)
		if _begge == nil || _begge._ccfeb {
			continue
		}
		_cbbae = _g.Min(_cbbae, _begge.depth())
	}
	return _cbbae
}

type rectRuling struct {
	_fdba rulingKind
	_feff markKind
	_ca.Color
	_ab.PdfRectangle
}

func (_aef *stateStack) empty() bool { return len(*_aef) == 0 }
func (_eceed *wordBag) sort() {
	for _, _ebdg := range _eceed._edgf {
		_cf.Slice(_ebdg, func(_bgad, _babb int) bool { return _faeff(_ebdg[_bgad], _ebdg[_babb]) < 0 })
	}
}

// String returns a description of `p`.
func (_gdecd *textPara) String() string {
	if _gdecd._ccfeb {
		return _a.Sprintf("\u0025\u0036\u002e\u0032\u0066\u0020\u005b\u0045\u004d\u0050\u0054\u0059\u005d", _gdecd.PdfRectangle)
	}
	_bbaeb := ""
	if _gdecd._bcfcf != nil {
		_bbaeb = _a.Sprintf("\u005b\u0025\u0064\u0078\u0025\u0064\u005d\u0020", _gdecd._bcfcf._cbcb, _gdecd._bcfcf._faed)
	}
	return _a.Sprintf("\u0025\u0036\u002e\u0032f \u0025\u0073\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u0020\u0025\u0071", _gdecd.PdfRectangle, _bbaeb, len(_gdecd._fcbfe), _gggg(_gdecd.text(), 50))
}

func _eedd(_fccc, _agdf _ab.PdfRectangle) bool {
	return _agdf.Llx <= _fccc.Urx && _fccc.Llx <= _agdf.Urx
}

type stateStack []*textState

func (_bdfg *PageText) getParagraphs() paraList {
	var _cfg rulingList
	if _bfcfg {
		_facb := _dcaf(_bdfg._dab)
		_cfg = append(_cfg, _facb...)
	}
	if _fbbb {
		_ccgd := _dggc(_bdfg._cdbf)
		_cfg = append(_cfg, _ccgd...)
	}
	_cfg, _cag := _cfg.toTilings()
	var _gec paraList
	_bga := len(_bdfg._cbf)
	for _cgbf := 0; _cgbf < 360 && _bga > 0; _cgbf += 90 {
		_bfa := make([]*textMark, 0, len(_bdfg._cbf)-_bga)
		for _, _aggc := range _bdfg._cbf {
			if _aggc._fbfge == _cgbf {
				_bfa = append(_bfa, _aggc)
			}
		}
		if len(_bfa) > 0 {
			_acf := _eefc(_bfa, _bdfg._cfce, _cfg, _cag, _bdfg._cbef._dccbb)
			_gec = append(_gec, _acf...)
			_bga -= len(_bfa)
		}
	}
	return _gec
}

func _adccb(_bdbaf []*textMark, _bbdae _ab.PdfRectangle) *textWord {
	_dbcdg := _bdbaf[0].PdfRectangle
	_efaf := _bdbaf[0]._acaa
	for _, _bgcae := range _bdbaf[1:] {
		_dbcdg = _abbb(_dbcdg, _bgcae.PdfRectangle)
		if _bgcae._acaa > _efaf {
			_efaf = _bgcae._acaa
		}
	}
	return &textWord{PdfRectangle: _dbcdg, _bacfd: _bdbaf, _faace: _bbdae.Ury - _dbcdg.Lly, _fcde: _efaf}
}

// ExtractPageImages returns the image contents of the page extractor, including data
// and position, size information for each image.
// A set of options to control page image extraction can be passed in. The options
// parameter can be nil for the default options. By default, inline stencil masks
// are not extracted.
func (_bge *Extractor) ExtractPageImages(options *ImageExtractOptions) (*PageImages, error) {
	_dc := &imageExtractContext{_bd: options}
	_gca := _dc.extractContentStreamImages(_bge._ef, _bge._fgf)
	if _gca != nil {
		return nil, _gca
	}
	return &PageImages{Images: _dc._daae}, nil
}

type event struct {
	_dcfab  float64
	_ccbfba bool
	_dbcaee int
}

func _gdac(_gebb, _bafg _ab.PdfRectangle) bool {
	return _gebb.Lly <= _bafg.Ury && _bafg.Lly <= _gebb.Ury
}

func _bgceg(_gdbfc *_ab.Image, _cacf _ca.Color) _cfb.Image {
	_geeb, _cdbg := int(_gdbfc.Width), int(_gdbfc.Height)
	_edgca := _cfb.NewRGBA(_cfb.Rect(0, 0, _geeb, _cdbg))
	for _gcfacd := 0; _gcfacd < _cdbg; _gcfacd++ {
		for _fgfbg := 0; _fgfbg < _geeb; _fgfbg++ {
			_befbb, _edgcd := _gdbfc.ColorAt(_fgfbg, _gcfacd)
			if _edgcd != nil {
				_da.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _fgfbg, _gcfacd)
				continue
			}
			_fbeb, _fadfg, _fgbbc, _ := _befbb.RGBA()
			var _dgegd _ca.Color
			if _fbeb+_fadfg+_fgbbc == 0 {
				_dgegd = _ca.Transparent
			} else {
				_dgegd = _cacf
			}
			_edgca.Set(_fgfbg, _gcfacd, _dgegd)
		}
	}
	return _edgca
}

func (_egfde *textTable) reduce() *textTable {
	_ggade := make([]int, 0, _egfde._faed)
	_bebb := make([]int, 0, _egfde._cbcb)
	for _deca := 0; _deca < _egfde._faed; _deca++ {
		if !_egfde.emptyCompositeRow(_deca) {
			_ggade = append(_ggade, _deca)
		}
	}
	for _cdcfe := 0; _cdcfe < _egfde._cbcb; _cdcfe++ {
		if !_egfde.emptyCompositeColumn(_cdcfe) {
			_bebb = append(_bebb, _cdcfe)
		}
	}
	if len(_ggade) == _egfde._faed && len(_bebb) == _egfde._cbcb {
		return _egfde
	}
	_gcgb := textTable{_cafe: _egfde._cafe, _cbcb: len(_bebb), _faed: len(_ggade), _ebgbd: make(map[uint64]*textPara, len(_bebb)*len(_ggade))}
	if _cabbg {
		_da.Log.Info("\u0072\u0065\u0064\u0075ce\u003a\u0020\u0025\u0064\u0078\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0064\u0078%\u0064", _egfde._cbcb, _egfde._faed, len(_bebb), len(_ggade))
		_da.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0043\u006f\u006c\u0073\u003a\u0020\u0025\u002b\u0076", _bebb)
		_da.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0052\u006f\u0077\u0073\u003a\u0020\u0025\u002b\u0076", _ggade)
	}
	for _gbe, _accb := range _ggade {
		for _ggeae, _edgcb := range _bebb {
			_cggd, _bggd := _egfde.getComposite(_edgcb, _accb)
			if _cggd == nil {
				continue
			}
			if _cabbg {
				_a.Printf("\u0020 \u0025\u0032\u0064\u002c \u0025\u0032\u0064\u0020\u0028%\u0032d\u002c \u0025\u0032\u0064\u0029\u0020\u0025\u0071\n", _ggeae, _gbe, _edgcb, _accb, _gggg(_cggd.merge().text(), 50))
			}
			_gcgb.putComposite(_ggeae, _gbe, _cggd, _bggd)
		}
	}
	return &_gcgb
}

func (_gafge paraList) xNeighbours(_bbegb float64) map[*textPara][]int {
	_cdab := make([]event, 2*len(_gafge))
	if _bbegb == 0 {
		for _fbdg, _ebbgf := range _gafge {
			_cdab[2*_fbdg] = event{_ebbgf.Llx, true, _fbdg}
			_cdab[2*_fbdg+1] = event{_ebbgf.Urx, false, _fbdg}
		}
	} else {
		for _egcc, _fgcf := range _gafge {
			_cdab[2*_egcc] = event{_fgcf.Llx - _bbegb*_fgcf.fontsize(), true, _egcc}
			_cdab[2*_egcc+1] = event{_fgcf.Urx + _bbegb*_fgcf.fontsize(), false, _egcc}
		}
	}
	return _gafge.eventNeighbours(_cdab)
}

func (_faf *stateStack) top() *textState {
	if _faf.empty() {
		return nil
	}
	return (*_faf)[_faf.size()-1]
}

func (_caf *textObject) setHorizScaling(_abff float64) {
	if _caf == nil {
		return
	}
	_caf._dbfd._fgc = _abff
}

// TextTable represents a table.
// Cells are ordered top-to-bottom, left-to-right.
// Cells[y] is the (0-offset) y'th row in the table.
// Cells[y][x] is the (0-offset) x'th column in the table.
type TextTable struct {
	_ab.PdfRectangle
	W, H  int
	Cells [][]TableCell
}
type textTable struct {
	_ab.PdfRectangle
	_cbcb, _faed int
	_cafe        bool
	_ebgbd       map[uint64]*textPara
	_fcda        map[uint64]compositeCell
}

func _eaegb(_dbgaab map[int][]float64) []int {
	_ggccb := make([]int, len(_dbgaab))
	_febdd := 0
	for _caddd := range _dbgaab {
		_ggccb[_febdd] = _caddd
		_febdd++
	}
	_cf.Ints(_ggccb)
	return _ggccb
}

func (_dbde *wordBag) scanBand(_aacd string, _ccac *wordBag, _aacda func(_gcbca *wordBag, _bgda *textWord) bool, _cba, _fbfe, _eegd float64, _deac, _eaa bool) int {
	_efcg := _ccac._caag
	var _aaaec map[int]map[*textWord]struct{}
	if !_deac {
		_aaaec = _dbde.makeRemovals()
	}
	_dfeb := _cbgd * _efcg
	_bdga := 0
	for _, _efab := range _dbde.depthBand(_cba-_dfeb, _fbfe+_dfeb) {
		if len(_dbde._edgf[_efab]) == 0 {
			continue
		}
		for _, _bdef := range _dbde._edgf[_efab] {
			if !(_cba-_dfeb <= _bdef._faace && _bdef._faace <= _fbfe+_dfeb) {
				continue
			}
			if !_aacda(_ccac, _bdef) {
				continue
			}
			_fddga := 2.0 * _g.Abs(_bdef._fcde-_ccac._caag) / (_bdef._fcde + _ccac._caag)
			_cgbd := _g.Max(_bdef._fcde/_ccac._caag, _ccac._caag/_bdef._fcde)
			_cgee := _g.Min(_fddga, _cgbd)
			if _eegd > 0 && _cgee > _eegd {
				continue
			}
			if _ccac.blocked(_bdef) {
				continue
			}
			if !_deac {
				_ccac.pullWord(_bdef, _efab, _aaaec)
			}
			_bdga++
			if !_eaa {
				if _bdef._faace < _cba {
					_cba = _bdef._faace
				}
				if _bdef._faace > _fbfe {
					_fbfe = _bdef._faace
				}
			}
			if _deac {
				break
			}
		}
	}
	if !_deac {
		_dbde.applyRemovals(_aaaec)
	}
	return _bdga
}
func (_cec *textObject) nextLine() { _cec.moveLP(0, -_cec._dbfd._abda) }
func (_bcdf rulingList) merge() *ruling {
	_dbdc := _bcdf[0]._dbbce
	_ebef := _bcdf[0]._beadd
	_gbdgf := _bcdf[0]._dgbc
	for _, _dbfad := range _bcdf[1:] {
		_dbdc += _dbfad._dbbce
		if _dbfad._beadd < _ebef {
			_ebef = _dbfad._beadd
		}
		if _dbfad._dgbc > _gbdgf {
			_gbdgf = _dbfad._dgbc
		}
	}
	_ddffa := &ruling{_ggcaf: _bcdf[0]._ggcaf, _bfccg: _bcdf[0]._bfccg, Color: _bcdf[0].Color, _dbbce: _dbdc / float64(len(_bcdf)), _beadd: _ebef, _dgbc: _gbdgf}
	if _gecc {
		_da.Log.Info("\u006de\u0072g\u0065\u003a\u0020\u0025\u0032d\u0020\u0076e\u0063\u0073\u0020\u0025\u0073", len(_bcdf), _ddffa)
		for _cccb, _gcafa := range _bcdf {
			_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _cccb, _gcafa)
		}
	}
	return _ddffa
}

func (_gdbcb gridTiling) log(_bdbg string) {
	if !_bcf {
		return
	}
	_da.Log.Info("\u0074i\u006ci\u006e\u0067\u003a\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0071", len(_gdbcb._cgeba), len(_gdbcb._gfdac), _bdbg)
	_a.Printf("\u0020\u0020\u0020l\u006c\u0078\u003d\u0025\u002e\u0032\u0066\u000a", _gdbcb._cgeba)
	_a.Printf("\u0020\u0020\u0020l\u006c\u0079\u003d\u0025\u002e\u0032\u0066\u000a", _gdbcb._gfdac)
	for _gadad, _fdgb := range _gdbcb._gfdac {
		_geege, _fdfbf := _gdbcb._afce[_fdgb]
		if !_fdfbf {
			continue
		}
		_a.Printf("%\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u000a", _gadad, _fdgb)
		for _fgfbe, _efbfe := range _gdbcb._cgeba {
			_ggfc, _dgcf := _geege[_efbfe]
			if !_dgcf {
				continue
			}
			_a.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _fgfbe, _ggfc.String())
		}
	}
}

const (
	_gacb  = true
	_bdcg  = true
	_becd  = true
	_fdfc  = false
	_agfg  = false
	_gfag  = 6
	_aad   = 3.0
	_cbg   = 200
	_dbbc  = true
	_gcbdb = true
	_bfcfg = true
	_fbbb  = true
	_ccfd  = false
)

func _gdeb(_fcgcf bounded) float64 { return -_fcgcf.bbox().Lly }

type structTreeRoot struct {
	_eede []structElement
	_aabe string
}

func (_gdfc *textPara) toCellTextMarks(_caagd *int) []TextMark {
	var _efdae []TextMark
	for _cfab, _efgg := range _gdfc._fcbfe {
		_bdba := _efgg.toTextMarks(_caagd)
		_gegd := _gacb && _efgg.endsInHyphen() && _cfab != len(_gdfc._fcbfe)-1
		if _gegd {
			_bdba = _acadc(_bdba, _caagd)
		}
		_efdae = append(_efdae, _bdba...)
		if !(_gegd || _cfab == len(_gdfc._fcbfe)-1) {
			_efdae = _beec(_efdae, _caagd, _agadc(_efgg._defc, _gdfc._fcbfe[_cfab+1]._defc))
		}
	}
	return _efdae
}

type ruling struct {
	_ggcaf rulingKind
	_bfccg markKind
	_ca.Color
	_dbbce float64
	_beadd float64
	_dgbc  float64
	_fbfc  float64
}

func (_dbcb *subpath) makeRectRuling(_cbgb _ca.Color) (*ruling, bool) {
	if _gddab {
		_da.Log.Info("\u006d\u0061\u006beR\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0070\u0061\u0074\u0068\u003d\u0025\u0076", _dbcb)
	}
	_gccbed := _dbcb._bfaa[:4]
	_fbdd := make(map[int]rulingKind, len(_gccbed))
	for _ddea, _fgbbf := range _gccbed {
		_bbad := _dbcb._bfaa[(_ddea+1)%4]
		_fbdd[_ddea] = _fcebe(_fgbbf, _bbad)
		if _gddab {
			_a.Printf("\u0025\u0034\u0064: \u0025\u0073\u0020\u003d\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u002d\u0020\u0025\u0036\u002e\u0032\u0066", _ddea, _fbdd[_ddea], _fgbbf, _bbad)
		}
	}
	if _gddab {
		_a.Printf("\u0020\u0020\u0020\u006b\u0069\u006e\u0064\u0073\u003d\u0025\u002b\u0076\u000a", _fbdd)
	}
	var _ecdd, _caadf []int
	for _fbda, _ccfee := range _fbdd {
		switch _ccfee {
		case _ceedg:
			_caadf = append(_caadf, _fbda)
		case _eccce:
			_ecdd = append(_ecdd, _fbda)
		}
	}
	if _gddab {
		_a.Printf("\u0020\u0020 \u0068\u006f\u0072z\u0073\u003d\u0025\u0064\u0020\u0025\u002b\u0076\u000a", len(_caadf), _caadf)
		_a.Printf("\u0020\u0020 \u0076\u0065\u0072t\u0073\u003d\u0025\u0064\u0020\u0025\u002b\u0076\u000a", len(_ecdd), _ecdd)
	}
	_cecb := (len(_caadf) == 2 && len(_ecdd) == 2) || (len(_caadf) == 2 && len(_ecdd) == 0 && _agbcf(_gccbed[_caadf[0]], _gccbed[_caadf[1]])) || (len(_ecdd) == 2 && len(_caadf) == 0 && _agdb(_gccbed[_ecdd[0]], _gccbed[_ecdd[1]]))
	if _gddab {
		_a.Printf(" \u0020\u0020\u0068\u006f\u0072\u007as\u003d\u0025\u0064\u0020\u0076\u0065\u0072\u0074\u0073=\u0025\u0064\u0020o\u006b=\u0025\u0074\u000a", len(_caadf), len(_ecdd), _cecb)
	}
	if !_cecb {
		if _gddab {
			_da.Log.Error("\u0021!\u006d\u0061\u006b\u0065R\u0065\u0063\u0074\u0052\u0075l\u0069n\u0067:\u0020\u0070\u0061\u0074\u0068\u003d\u0025v", _dbcb)
			_a.Printf(" \u0020\u0020\u0068\u006f\u0072\u007as\u003d\u0025\u0064\u0020\u0076\u0065\u0072\u0074\u0073=\u0025\u0064\u0020o\u006b=\u0025\u0074\u000a", len(_caadf), len(_ecdd), _cecb)
		}
		return &ruling{}, false
	}
	if len(_ecdd) == 0 {
		for _affd, _gfcd := range _fbdd {
			if _gfcd != _ceedg {
				_ecdd = append(_ecdd, _affd)
			}
		}
	}
	if len(_caadf) == 0 {
		for _abfa, _abbe := range _fbdd {
			if _abbe != _eccce {
				_caadf = append(_caadf, _abfa)
			}
		}
	}
	if _gddab {
		_da.Log.Info("\u006da\u006b\u0065R\u0065\u0063\u0074\u0052u\u006c\u0069\u006eg\u003a\u0020\u0068\u006f\u0072\u007a\u0073\u003d\u0025d \u0076\u0065\u0072t\u0073\u003d%\u0064\u0020\u0070\u006f\u0069\u006et\u0073\u003d%\u0064\u000a"+"\u0009\u0020\u0068o\u0072\u007a\u0073\u003d\u0025\u002b\u0076\u000a"+"\u0009\u0020\u0076e\u0072\u0074\u0073\u003d\u0025\u002b\u0076\u000a"+"\t\u0070\u006f\u0069\u006e\u0074\u0073\u003d\u0025\u002b\u0076", len(_caadf), len(_ecdd), len(_gccbed), _caadf, _ecdd, _gccbed)
	}
	var _cade, _gacbb, _dgfdg, _acbb _fa.Point
	if _gccbed[_caadf[0]].Y > _gccbed[_caadf[1]].Y {
		_dgfdg, _acbb = _gccbed[_caadf[0]], _gccbed[_caadf[1]]
	} else {
		_dgfdg, _acbb = _gccbed[_caadf[1]], _gccbed[_caadf[0]]
	}
	if _gccbed[_ecdd[0]].X > _gccbed[_ecdd[1]].X {
		_cade, _gacbb = _gccbed[_ecdd[0]], _gccbed[_ecdd[1]]
	} else {
		_cade, _gacbb = _gccbed[_ecdd[1]], _gccbed[_ecdd[0]]
	}
	_baecb := _ab.PdfRectangle{Llx: _cade.X, Urx: _gacbb.X, Lly: _acbb.Y, Ury: _dgfdg.Y}
	if _baecb.Llx > _baecb.Urx {
		_baecb.Llx, _baecb.Urx = _baecb.Urx, _baecb.Llx
	}
	if _baecb.Lly > _baecb.Ury {
		_baecb.Lly, _baecb.Ury = _baecb.Ury, _baecb.Lly
	}
	_eaec := rectRuling{PdfRectangle: _baecb, _fdba: _ebdd(_baecb), Color: _cbgb}
	if _eaec._fdba == _fbag {
		if _gddab {
			_da.Log.Error("\u006da\u006b\u0065\u0052\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006eg\u003a\u0020\u006b\u0069\u006e\u0064\u003d\u006e\u0069\u006c")
		}
		return nil, false
	}
	_dcade, _bgbc := _eaec.asRuling()
	if !_bgbc {
		if _gddab {
			_da.Log.Error("\u006da\u006b\u0065\u0052\u0065c\u0074\u0052\u0075\u006c\u0069n\u0067:\u0020!\u0069\u0073\u0052\u0075\u006c\u0069\u006eg")
		}
		return nil, false
	}
	if _egc {
		_a.Printf("\u0020\u0020\u0020\u0072\u003d\u0025\u0073\u000a", _dcade.String())
	}
	return _dcade, true
}

func (_cbbdc *textTable) getRight() paraList {
	_cfcgg := make(paraList, _cbbdc._faed)
	for _adag := 0; _adag < _cbbdc._faed; _adag++ {
		_dbgfd := _cbbdc.get(_cbbdc._cbcb-1, _adag)._eefcd
		if _dbgfd.taken() {
			return nil
		}
		_cfcgg[_adag] = _dbgfd
	}
	for _agdc := 0; _agdc < _cbbdc._faed-1; _agdc++ {
		if _cfcgg[_agdc]._afeeg != _cfcgg[_agdc+1] {
			return nil
		}
	}
	return _cfcgg
}
func _cfba(_bfdc, _ggf bounded) float64 { return _bfdc.bbox().Llx - _ggf.bbox().Urx }
func (_eaedf *subpath) isQuadrilateral() bool {
	if len(_eaedf._bfaa) < 4 || len(_eaedf._bfaa) > 5 {
		return false
	}
	if len(_eaedf._bfaa) == 5 {
		_cgcae := _eaedf._bfaa[0]
		_fffc := _eaedf._bfaa[4]
		if _cgcae.X != _fffc.X || _cgcae.Y != _fffc.Y {
			return false
		}
	}
	return true
}

// String returns a description of `b`.
func (_eabc *wordBag) String() string {
	var _ecea []string
	for _, _debb := range _eabc.depthIndexes() {
		_gefd := _eabc._edgf[_debb]
		for _, _cafd := range _gefd {
			_ecea = append(_ecea, _cafd._faaf)
		}
	}
	return _a.Sprintf("\u0025.\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065=\u0025\u002e\u0032\u0066\u0020\u0025\u0064\u0020\u0025\u0071", _eabc.PdfRectangle, _eabc._caag, len(_ecea), _ecea)
}

func (_ggd *textObject) setCharSpacing(_agf float64) {
	if _ggd == nil {
		return
	}
	_ggd._dbfd._edg = _agf
	if _bbf {
		_da.Log.Info("\u0073\u0065t\u0043\u0068\u0061\u0072\u0053\u0070\u0061\u0063\u0069\u006e\u0067\u003a\u0020\u0025\u002e\u0032\u0066\u0020\u0073\u0074\u0061\u0074e=\u0025\u0073", _agf, _ggd._dbfd.String())
	}
}

func (_cbed paraList) llyOrdering() []int {
	_gbfa := make([]int, len(_cbed))
	for _dgfbf := range _cbed {
		_gbfa[_dgfbf] = _dgfbf
	}
	_cf.SliceStable(_gbfa, func(_ffae, _egacc int) bool {
		_eega, _cfcd := _gbfa[_ffae], _gbfa[_egacc]
		return _cbed[_eega].Lly < _cbed[_cfcd].Lly
	})
	return _gbfa
}

func (_afccgf *wordBag) arrangeText() *textPara {
	_afccgf.sort()
	if _bdcg {
		_afccgf.removeDuplicates()
	}
	var _dgcb []*textLine
	for _, _cddb := range _afccgf.depthIndexes() {
		for !_afccgf.empty(_cddb) {
			_cfcec := _afccgf.firstReadingIndex(_cddb)
			_ccba := _afccgf.firstWord(_cfcec)
			_gfadd := _abag(_afccgf, _cfcec)
			_agcde := _ccba._fcde
			_dbcga := _ccba._faace - _cbgd*_agcde
			_feda := _ccba._faace + _cbgd*_agcde
			_begbc := _fffad * _agcde
			_gdbce := _aedd * _agcde
		_bcab:
			for {
				var _bbdc *textWord
				_aadf := 0
				for _, _cdcb := range _afccgf.depthBand(_dbcga, _feda) {
					_becb := _afccgf.highestWord(_cdcb, _dbcga, _feda)
					if _becb == nil {
						continue
					}
					_dbfb := _cfba(_becb, _gfadd._gfed[len(_gfadd._gfed)-1])
					if _dbfb < -_gdbce {
						break _bcab
					}
					if _dbfb > _begbc {
						continue
					}
					if _bbdc != nil && _faeff(_becb, _bbdc) >= 0 {
						continue
					}
					_bbdc = _becb
					_aadf = _cdcb
				}
				if _bbdc == nil {
					break
				}
				_gfadd.pullWord(_afccgf, _bbdc, _aadf)
			}
			_gfadd.markWordBoundaries()
			_dgcb = append(_dgcb, _gfadd)
		}
	}
	if len(_dgcb) == 0 {
		return nil
	}
	_cf.Slice(_dgcb, func(_dcfg, _dccf int) bool { return _cega(_dgcb[_dcfg], _dgcb[_dccf]) < 0 })
	_dgdg := _eabgf(_afccgf.PdfRectangle, _dgcb)
	if _dabe {
		_da.Log.Info("\u0061\u0072\u0072an\u0067\u0065\u0054\u0065\u0078\u0074\u0020\u0021\u0021\u0021\u0020\u0070\u0061\u0072\u0061\u003d\u0025\u0073", _dgdg.String())
		if _bdcf {
			for _ggdae, _fedg := range _dgdg._fcbfe {
				_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _ggdae, _fedg.String())
				if _dfae {
					for _fcec, _eccg := range _fedg._gfed {
						_a.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _fcec, _eccg.String())
						for _cfdbf, _acgd := range _eccg._bacfd {
							_a.Printf("\u00251\u0032\u0064\u003a\u0020\u0025\u0073\n", _cfdbf, _acgd.String())
						}
					}
				}
			}
		}
	}
	return _dgdg
}

func (_adff rulingList) mergePrimary() float64 {
	_dfcf := _adff[0]._dbbce
	for _, _gfgcf := range _adff[1:] {
		_dfcf += _gfgcf._dbbce
	}
	return _dfcf / float64(len(_adff))
}

// GetContentStreamOps returns the contentStreamOps field of `pt`.
func (_bdf *PageText) GetContentStreamOps() *_bb.ContentStreamOperations { return _bdf._ffef }

func (_bgff *shapesState) cubicTo(_gcae, _dega, _ada, _gfe, _cfd, _fgfebe float64) {
	if _caad {
		_da.Log.Info("\u0063\u0075\u0062\u0069\u0063\u0054\u006f\u003a")
	}
	_bgff.addPoint(_cfd, _fgfebe)
}

func (_cfcgb paraList) merge() *textPara {
	_da.Log.Trace("\u006d\u0065\u0072\u0067\u0065:\u0020\u0070\u0061\u0072\u0061\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_cfcgb))
	if len(_cfcgb) == 0 {
		return nil
	}
	_cfcgb.sortReadingOrder()
	_aade := _cfcgb[0].PdfRectangle
	_adgbf := _cfcgb[0]._fcbfe
	for _, _adga := range _cfcgb[1:] {
		_aade = _abbb(_aade, _adga.PdfRectangle)
		_adgbf = append(_adgbf, _adga._fcbfe...)
	}
	return _eabgf(_aade, _adgbf)
}

func (_bfag *subpath) close() {
	if !_gagg(_bfag._bfaa[0], _bfag.last()) {
		_bfag.add(_bfag._bfaa[0])
	}
	_bfag._gfde = true
	_bfag.removeDuplicates()
}

func _ceed(_cfge, _eaca *textPara) bool {
	if _cfge._ccfeb || _eaca._ccfeb {
		return true
	}
	return _ccae(_cfge.depth() - _eaca.depth())
}
func (_fdcb *stateStack) size() int { return len(*_fdcb) }

// String returns a string describing `ma`.
func (_bdee TextMarkArray) String() string {
	_afcb := len(_bdee._beaf)
	if _afcb == 0 {
		return "\u0045\u004d\u0050T\u0059"
	}
	_gad := _bdee._beaf[0]
	_gcd := _bdee._beaf[_afcb-1]
	return _a.Sprintf("\u007b\u0054\u0045\u0058\u0054\u004d\u0041\u0052K\u0041\u0052\u0052AY\u003a\u0020\u0025\u0064\u0020\u0065l\u0065\u006d\u0065\u006e\u0074\u0073\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025s\u000a\u0009\u0020\u006c\u0061\u0073\u0074\u003d%\u0073\u007d", _afcb, _gad, _gcd)
}
func _edcfb(_fcgff float64) bool { return _g.Abs(_fcgff) < _efbg }
func _ccae(_bbfed float64) bool  { return _g.Abs(_bbfed) < _gccbd }
func _ede(_bgaeb []compositeCell) []float64 {
	var _abcgc []*textLine
	_fgce := 0
	for _, _eddb := range _bgaeb {
		_fgce += len(_eddb.paraList)
		_abcgc = append(_abcgc, _eddb.lines()...)
	}
	_cf.Slice(_abcgc, func(_caga, _gaadf int) bool {
		_gffd, _eeae := _abcgc[_caga], _abcgc[_gaadf]
		_bbbgd, _ffcb := _gffd._defc, _eeae._defc
		if !_ccae(_bbbgd - _ffcb) {
			return _bbbgd < _ffcb
		}
		return _gffd.Llx < _eeae.Llx
	})
	if _cabbg {
		_a.Printf("\u0020\u0020\u0020 r\u006f\u0077\u0042\u006f\u0072\u0064\u0065\u0072\u0073:\u0020%\u0064 \u0070a\u0072\u0061\u0073\u0020\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u000a", _fgce, len(_abcgc))
		for _agae, _beffb := range _abcgc {
			_a.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _agae, _beffb)
		}
	}
	var _gfgae []float64
	_feebe := _abcgc[0]
	var _efac [][]*textLine
	_afbb := []*textLine{_feebe}
	for _eedag, _baff := range _abcgc[1:] {
		if _baff.Ury < _feebe.Lly {
			_efeea := 0.5 * (_baff.Ury + _feebe.Lly)
			if _cabbg {
				_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u003c\u0020\u0025\u0036.\u0032f\u0020\u0062\u006f\u0072\u0064\u0065\u0072\u003d\u0025\u0036\u002e\u0032\u0066\u000a"+"\u0009\u0020\u0071\u003d\u0025\u0073\u000a\u0009\u0020p\u003d\u0025\u0073\u000a", _eedag, _baff.Ury, _feebe.Lly, _efeea, _feebe, _baff)
			}
			_gfgae = append(_gfgae, _efeea)
			_efac = append(_efac, _afbb)
			_afbb = nil
		}
		_afbb = append(_afbb, _baff)
		if _baff.Lly < _feebe.Lly {
			_feebe = _baff
		}
	}
	if len(_afbb) > 0 {
		_efac = append(_efac, _afbb)
	}
	if _cabbg {
		_a.Printf(" \u0020\u0020\u0020\u0020\u0020\u0020 \u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064o\u0072\u0073\u003d%\u0036.\u0032\u0066\u000a", _gfgae)
	}
	if _cabbg {
		_da.Log.Info("\u0072\u006f\u0077\u003d\u0025\u0064", len(_bgaeb))
		for _gcfg, _babega := range _bgaeb {
			_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _gcfg, _babega)
		}
		_da.Log.Info("\u0067r\u006f\u0075\u0070\u0073\u003d\u0025d", len(_efac))
		for _bgbb, _agade := range _efac {
			_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0064\u000a", _bgbb, len(_agade))
			for _fedaf, _agcda := range _agade {
				_a.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _fedaf, _agcda)
			}
		}
	}
	_bbce := true
	for _bcfag, _fefff := range _efac {
		_deed := true
		for _fdebc, _dcdaf := range _bgaeb {
			if _cabbg {
				_a.Printf("\u0020\u0020\u0020\u007e\u007e\u007e\u0067\u0072\u006f\u0075\u0070\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u0020\u0063\u0065\u006cl\u0020\u0025\u0064\u0020\u006ff\u0020\u0025d\u0020\u0025\u0073\u000a", _bcfag, len(_efac), _fdebc, len(_bgaeb), _dcdaf)
			}
			if !_dcdaf.hasLines(_fefff) {
				if _cabbg {
					_a.Printf("\u0020\u0020\u0020\u0021\u0021\u0021\u0067\u0072\u006f\u0075\u0070\u0020\u0025d\u0020\u006f\u0066\u0020\u0025\u0064 \u0063\u0065\u006c\u006c\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064 \u004f\u0055\u0054\u000a", _bcfag, len(_efac), _fdebc, len(_bgaeb))
				}
				_deed = false
				break
			}
		}
		if !_deed {
			_bbce = false
			break
		}
	}
	if !_bbce {
		if _cabbg {
			_da.Log.Info("\u0072\u006f\u0077\u0020\u0063o\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u0020\u0064\u006f\u006e\u0027\u0074 \u0073\u0070\u0061\u006e\u0020\u0061\u006c\u006c\u0020\u0063\u0065\u006c\u006c\u0073\u0020\u0069\u006e\u0020\u0072\u006f\u0077\u002e\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006eg")
		}
		_gfgae = nil
	}
	if _cabbg && _gfgae != nil {
		_a.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u002a\u002a*\u0072\u006f\u0077\u0043\u006f\u0072\u0072i\u0064\u006f\u0072\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u000a", _gfgae)
	}
	return _gfgae
}

func _dfee(_agbce *_ab.Image, _ecde _ca.Color) _cfb.Image {
	_abee, _gcgg := int(_agbce.Width), int(_agbce.Height)
	_eddc := _cfb.NewRGBA(_cfb.Rect(0, 0, _abee, _gcgg))
	for _dbda := 0; _dbda < _gcgg; _dbda++ {
		for _cefe := 0; _cefe < _abee; _cefe++ {
			_aafeg, _egfdeb := _agbce.ColorAt(_cefe, _dbda)
			if _egfdeb != nil {
				_da.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _cefe, _dbda)
				continue
			}
			_aecb, _cddgc, _ddbbg, _ := _aafeg.RGBA()
			var _ecebc _ca.Color
			if _aecb+_cddgc+_ddbbg == 0 {
				_ecebc = _ecde
			} else {
				_ecebc = _ca.Transparent
			}
			_eddc.Set(_cefe, _dbda, _ecebc)
		}
	}
	return _eddc
}

func _faae(_dbcc string) bool {
	if _fe.RuneCountInString(_dbcc) < _fafa {
		return false
	}
	_agfa, _egcg := _fe.DecodeLastRuneInString(_dbcc)
	if _egcg <= 0 || !_fg.Is(_fg.Hyphen, _agfa) {
		return false
	}
	_agfa, _egcg = _fe.DecodeLastRuneInString(_dbcc[:len(_dbcc)-_egcg])
	return _egcg > 0 && !_fg.IsSpace(_agfa)
}
func _gagg(_bdfe, _dege _fa.Point) bool { return _bdfe.X == _dege.X && _bdfe.Y == _dege.Y }

// String returns a description of `state`.
func (_dada *textState) String() string {
	_bcg := "\u005bN\u004f\u0054\u0020\u0053\u0045\u0054]"
	if _dada._ggda != nil {
		_bcg = _dada._ggda.BaseFont()
	}
	return _a.Sprintf("\u0074\u0063\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0077\u003d\u0025\u002e\u0032\u0066 \u0074f\u0073\u003d\u0025\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071", _dada._edg, _dada._ebbfc, _dada._egbc, _bcg)
}

var (
	_dag = _e.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	_ga  = _e.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
)

func _becg(_adagd []*textWord, _ecaf int) []*textWord {
	_egbgeb := len(_adagd)
	copy(_adagd[_ecaf:], _adagd[_ecaf+1:])
	return _adagd[:_egbgeb-1]
}

var _ebda = []string{"\u2756", "\u27a2", "\u2713", "\u2022", "\uf0a7", "\u25a1", "\u2212", "\u25a0", "\u25aa", "\u006f"}

func (_fcedb *subpath) clear() { *_fcedb = subpath{} }
func _daaga(_ebed *textWord, _ccbc float64, _acbe, _ddga rulingList) *wordBag {
	_dadd := _aeef(_ebed._faace)
	_fca := []*textWord{_ebed}
	_aefe := wordBag{_edgf: map[int][]*textWord{_dadd: _fca}, PdfRectangle: _ebed.PdfRectangle, _caag: _ebed._fcde, _bfgd: _ccbc, _bbcd: _acbe, _bae: _ddga}
	return &_aefe
}

func (_acdg *textPara) depth() float64 {
	if _acdg._ccfeb {
		return -1.0
	}
	if len(_acdg._fcbfe) > 0 {
		return _acdg._fcbfe[0]._defc
	}
	return _acdg._bcfcf.depth()
}

func (_cgdb compositeCell) hasLines(_cbbf []*textLine) bool {
	for _dcgc, _gbgabc := range _cbbf {
		_bcaf := _gega(_cgdb.PdfRectangle, _gbgabc.PdfRectangle)
		if _cabbg {
			_a.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u005e\u005e\u005e\u0069\u006e\u0074\u0065\u0072\u0073e\u0063t\u0073\u003d\u0025\u0074\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u000a", _bcaf, _dcgc, len(_cbbf))
			_a.Printf("\u0020\u0020\u0020\u0020  \u005e\u005e\u005e\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u003d\u0025s\u000a", _cgdb)
			_a.Printf("\u0020 \u0020 \u0020\u0020\u0020\u006c\u0069\u006e\u0065\u003d\u0025\u0073\u000a", _gbgabc)
		}
		if _bcaf {
			return true
		}
	}
	return false
}

func (_dfcd *textPara) isAtom() *textTable {
	_efagfb := _dfcd
	_edcde := _dfcd._eefcd
	_bgce := _dfcd._afeeg
	if _edcde.taken() || _bgce.taken() {
		return nil
	}
	_gefdc := _edcde._afeeg
	if _gefdc.taken() || _gefdc != _bgce._eefcd {
		return nil
	}
	return _eedc(_efagfb, _edcde, _bgce, _gefdc)
}

// String returns a description of `k`.
func (_fdcdc markKind) String() string {
	_ffga, _eggae := _agfc[_fdcdc]
	if !_eggae {
		return _a.Sprintf("\u004e\u006f\u0074\u0020\u0061\u0020\u006d\u0061\u0072k\u003a\u0020\u0025\u0064", _fdcdc)
	}
	return _ffga
}

func (_dcaag *wordBag) applyRemovals(_fffec map[int]map[*textWord]struct{}) {
	for _efgc, _geafa := range _fffec {
		if len(_geafa) == 0 {
			continue
		}
		_gfda := _dcaag._edgf[_efgc]
		_cgcf := len(_gfda) - len(_geafa)
		if _cgcf == 0 {
			delete(_dcaag._edgf, _efgc)
			continue
		}
		_agaf := make([]*textWord, _cgcf)
		_gbb := 0
		for _, _ddcf := range _gfda {
			if _, _ccdg := _geafa[_ddcf]; !_ccdg {
				_agaf[_gbb] = _ddcf
				_gbb++
			}
		}
		_dcaag._edgf[_efgc] = _agaf
	}
}

func _eedc(_dgcfa, _bbcc, _afbce, _fdac *textPara) *textTable {
	_ebfd := &textTable{_cbcb: 2, _faed: 2, _ebgbd: make(map[uint64]*textPara, 4)}
	_ebfd.put(0, 0, _dgcfa)
	_ebfd.put(1, 0, _bbcc)
	_ebfd.put(0, 1, _afbce)
	_ebfd.put(1, 1, _fdac)
	return _ebfd
}

func _dggc(_bdfdd []pathSection) rulingList {
	_bbede(_bdfdd)
	if _egc {
		_da.Log.Info("\u006da\u006b\u0065\u0046\u0069l\u006c\u0052\u0075\u006c\u0069n\u0067s\u003a \u0025\u0064\u0020\u0066\u0069\u006c\u006cs", len(_bdfdd))
	}
	var _efdff rulingList
	for _, _fceb := range _bdfdd {
		for _, _fddac := range _fceb._egfg {
			if !_fddac.isQuadrilateral() {
				if _egc {
					_da.Log.Error("!\u0069s\u0051\u0075\u0061\u0064\u0072\u0069\u006c\u0061t\u0065\u0072\u0061\u006c: \u0025\u0073", _fddac)
				}
				continue
			}
			if _aegff, _bdffdf := _fddac.makeRectRuling(_fceb.Color); _bdffdf {
				_efdff = append(_efdff, _aegff)
			} else {
				if _gddab {
					_da.Log.Error("\u0021\u006d\u0061\u006beR\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0025\u0073", _fddac)
				}
			}
		}
	}
	if _egc {
		_da.Log.Info("\u006d\u0061\u006b\u0065Fi\u006c\u006c\u0052\u0075\u006c\u0069\u006e\u0067\u0073\u003a\u0020\u0025\u0073", _efdff.String())
	}
	return _efdff
}

func (_egeb paraList) log(_cabdg string) {
	if !_bgada {
		return
	}
	_da.Log.Info("%\u0038\u0073\u003a\u0020\u0025\u0064 \u0070\u0061\u0072\u0061\u0073\u0020=\u003d\u003d\u003d\u003d\u003d\u003d\u002d-\u002d\u002d\u002d\u002d\u002d\u003d\u003d\u003d\u003d\u003d=\u003d", _cabdg, len(_egeb))
	for _edfc, _ceee := range _egeb {
		if _ceee == nil {
			continue
		}
		_ggac := _ceee.text()
		_abce := "\u0020\u0020"
		if _ceee._bcfcf != nil {
			_abce = _a.Sprintf("\u005b%\u0064\u0078\u0025\u0064\u005d", _ceee._bcfcf._cbcb, _ceee._bcfcf._faed)
		}
		_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0025s\u0020\u0025\u0071\u000a", _edfc, _ceee.PdfRectangle, _abce, _gggg(_ggac, 50))
	}
}

func _gggg(_aafa string, _deggd int) string {
	if len(_aafa) < _deggd {
		return _aafa
	}
	return _aafa[:_deggd]
}

func _cfgd(_dcfc _ab.PdfRectangle) *ruling {
	return &ruling{_ggcaf: _ceedg, _dbbce: _dcfc.Ury, _beadd: _dcfc.Llx, _dgbc: _dcfc.Urx}
}

func _cgdc(_defbg []*wordBag) []*wordBag {
	if len(_defbg) <= 1 {
		return _defbg
	}
	if _dabe {
		_da.Log.Info("\u006d\u0065\u0072\u0067\u0065\u0057\u006f\u0072\u0064B\u0061\u0067\u0073\u003a")
	}
	_cf.Slice(_defbg, func(_bceg, _cabbaf int) bool {
		_eeb, _cfff := _defbg[_bceg], _defbg[_cabbaf]
		_feg := _eeb.Width() * _eeb.Height()
		_ageg := _cfff.Width() * _cfff.Height()
		if _feg != _ageg {
			return _feg > _ageg
		}
		if _eeb.Height() != _cfff.Height() {
			return _eeb.Height() > _cfff.Height()
		}
		return _bceg < _cabbaf
	})
	var _aefc []*wordBag
	_debd := make(intSet)
	for _egfb := 0; _egfb < len(_defbg); _egfb++ {
		if _debd.has(_egfb) {
			continue
		}
		_dge := _defbg[_egfb]
		for _bdd := _egfb + 1; _bdd < len(_defbg); _bdd++ {
			if _debd.has(_egfb) {
				continue
			}
			_abgb := _defbg[_bdd]
			_gdec := _dge.PdfRectangle
			_gdec.Llx -= _dge._caag
			if _bgeg(_gdec, _abgb.PdfRectangle) {
				_dge.absorb(_abgb)
				_debd.add(_bdd)
			}
		}
		_aefc = append(_aefc, _dge)
	}
	if len(_defbg) != len(_aefc)+len(_debd) {
		_da.Log.Error("\u006d\u0065\u0072ge\u0057\u006f\u0072\u0064\u0042\u0061\u0067\u0073\u003a \u0025d\u2192%\u0064 \u0061\u0062\u0073\u006f\u0072\u0062\u0065\u0064\u003d\u0025\u0064", len(_defbg), len(_aefc), len(_debd))
	}
	return _aefc
}

func (_adfac *textTable) growTable() {
	_caadc := func(_fgegd paraList) {
		_adfac._faed++
		for _bfed := 0; _bfed < _adfac._cbcb; _bfed++ {
			_aefgc := _fgegd[_bfed]
			_adfac.put(_bfed, _adfac._faed-1, _aefgc)
		}
	}
	_fecab := func(_dcfec paraList) {
		_adfac._cbcb++
		for _gbcfg := 0; _gbcfg < _adfac._faed; _gbcfg++ {
			_gdfa := _dcfec[_gbcfg]
			_adfac.put(_adfac._cbcb-1, _gbcfg, _gdfa)
		}
	}
	if _cdegf {
		_adfac.log("\u0067r\u006f\u0077\u0054\u0061\u0062\u006ce")
	}
	for _ddbgg := 0; ; _ddbgg++ {
		_fdfcb := false
		_fadaaa := _adfac.getDown()
		_cbde := _adfac.getRight()
		if _cdegf {
			_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _ddbgg, _adfac)
			_a.Printf("\u0020\u0020 \u0020\u0020\u0020 \u0020\u0064\u006f\u0077\u006e\u003d\u0025\u0073\u000a", _fadaaa)
			_a.Printf("\u0020\u0020 \u0020\u0020\u0020 \u0072\u0069\u0067\u0068\u0074\u003d\u0025\u0073\u000a", _cbde)
		}
		if _fadaaa != nil && _cbde != nil {
			_fdcf := _fadaaa[len(_fadaaa)-1]
			if !_fdcf.taken() && _fdcf == _cbde[len(_cbde)-1] {
				_caadc(_fadaaa)
				if _cbde = _adfac.getRight(); _cbde != nil {
					_fecab(_cbde)
					_adfac.put(_adfac._cbcb-1, _adfac._faed-1, _fdcf)
				}
				_fdfcb = true
			}
		}
		if !_fdfcb && _fadaaa != nil {
			_caadc(_fadaaa)
			_fdfcb = true
		}
		if !_fdfcb && _cbde != nil {
			_fecab(_cbde)
			_fdfcb = true
		}
		if !_fdfcb {
			break
		}
	}
}
func (_efggf intSet) has(_bfebf int) bool { _, _gddff := _efggf[_bfebf]; return _gddff }

const (
	_fbag rulingKind = iota
	_ceedg
	_eccce
)

type markKind int

// String returns a human readable description of `path`.
func (_dgaf *subpath) String() string {
	_faaae := _dgaf._bfaa
	_egaaf := len(_faaae)
	if _egaaf <= 5 {
		return _a.Sprintf("\u0025d\u003a\u0020\u0025\u0036\u002e\u0032f", _egaaf, _faaae)
	}
	return _a.Sprintf("\u0025d\u003a\u0020\u0025\u0036.\u0032\u0066\u0020\u0025\u0036.\u0032f\u0020.\u002e\u002e\u0020\u0025\u0036\u002e\u0032f", _egaaf, _faaae[0], _faaae[1], _faaae[_egaaf-1])
}

func _cdcc(_eaecea int, _ecgf func(int, int) bool) []int {
	_ecaeb := make([]int, _eaecea)
	for _caadd := range _ecaeb {
		_ecaeb[_caadd] = _caadd
	}
	_cf.Slice(_ecaeb, func(_bbcee, _dgeea int) bool { return _ecgf(_ecaeb[_bbcee], _ecaeb[_dgeea]) })
	return _ecaeb
}

func (_fffg *wordBag) blocked(_bgfc *textWord) bool {
	if _bgfc.Urx < _fffg.Llx {
		_afb := _gdfge(_bgfc.PdfRectangle)
		_beda := _cebb(_fffg.PdfRectangle)
		if _fffg._bbcd.blocks(_afb, _beda) {
			if _dee {
				_da.Log.Info("\u0062\u006c\u006f\u0063ke\u0064\u0020\u2190\u0078\u003a\u0020\u0025\u0073\u0020\u0025\u0073", _bgfc, _fffg)
			}
			return true
		}
	} else if _fffg.Urx < _bgfc.Llx {
		_faebc := _gdfge(_fffg.PdfRectangle)
		_cdbc := _cebb(_bgfc.PdfRectangle)
		if _fffg._bbcd.blocks(_faebc, _cdbc) {
			if _dee {
				_da.Log.Info("b\u006co\u0063\u006b\u0065\u0064\u0020\u0078\u2192\u0020:\u0020\u0025\u0073\u0020%s", _bgfc, _fffg)
			}
			return true
		}
	}
	if _bgfc.Ury < _fffg.Lly {
		_deae := _cfgd(_bgfc.PdfRectangle)
		_fgde := _ccdb(_fffg.PdfRectangle)
		if _fffg._bae.blocks(_deae, _fgde) {
			if _dee {
				_da.Log.Info("\u0062\u006c\u006f\u0063ke\u0064\u0020\u2190\u0079\u003a\u0020\u0025\u0073\u0020\u0025\u0073", _bgfc, _fffg)
			}
			return true
		}
	} else if _fffg.Ury < _bgfc.Lly {
		_gdf := _cfgd(_fffg.PdfRectangle)
		_fbg := _ccdb(_bgfc.PdfRectangle)
		if _fffg._bae.blocks(_gdf, _fbg) {
			if _dee {
				_da.Log.Info("b\u006co\u0063\u006b\u0065\u0064\u0020\u0079\u2192\u0020:\u0020\u0025\u0073\u0020%s", _bgfc, _fffg)
			}
			return true
		}
	}
	return false
}
func (_dcfa *textMark) bbox() _ab.PdfRectangle { return _dcfa.PdfRectangle }
func _fdcg(_gadd map[int][]float64) string {
	_babed := _eaegb(_gadd)
	_bbgb := make([]string, len(_gadd))
	for _bbfcf, _ggad := range _babed {
		_bbgb[_bbfcf] = _a.Sprintf("\u0025\u0064\u003a\u0020\u0025\u002e\u0032\u0066", _ggad, _gadd[_ggad])
	}
	return _a.Sprintf("\u007b\u0025\u0073\u007d", _cb.Join(_bbgb, "\u002c\u0020"))
}

func _ecefg(_bgbd _fa.Matrix) _fa.Point {
	_accf, _cca := _bgbd.Translation()
	return _fa.Point{X: _accf, Y: _cca}
}

type compositeCell struct {
	_ab.PdfRectangle
	paraList
}

func (_gcad rulingList) aligned() bool {
	if len(_gcad) < 2 {
		return false
	}
	_efedf := make(map[*ruling]int)
	_efedf[_gcad[0]] = 0
	for _, _cdga := range _gcad[1:] {
		_gfbe := false
		for _gggc := range _efedf {
			if _cdga.gridIntersecting(_gggc) {
				_efedf[_gggc]++
				_gfbe = true
				break
			}
		}
		if !_gfbe {
			_efedf[_cdga] = 0
		}
	}
	_aecd := 0
	for _, _bdfa := range _efedf {
		if _bdfa == 0 {
			_aecd++
		}
	}
	_cgeg := float64(_aecd) / float64(len(_gcad))
	_acgef := _cgeg <= 1.0-_gfad
	if _egc {
		_da.Log.Info("\u0061\u006c\u0069\u0067\u006e\u0065\u0064\u003d\u0025\u0074\u0020\u0075\u006em\u0061\u0074\u0063\u0068\u0065\u0064=\u0025\u002e\u0032\u0066\u003d\u0025\u0064\u002f\u0025\u0064\u0020\u0076\u0065c\u0073\u003d\u0025\u0073", _acgef, _cgeg, _aecd, len(_gcad), _gcad.String())
	}
	return _acgef
}

func _agbcf(_bgdb, _dbbcea _fa.Point) bool {
	_cgegf := _g.Abs(_bgdb.X - _dbbcea.X)
	_fabb := _g.Abs(_bgdb.Y - _dbbcea.Y)
	return _dbdg(_fabb, _cgegf)
}

func (_ebdb *textPara) getListLines() []*textLine {
	var _fcgcb []*textLine
	_fdda := _ffea(_ebdb._fcbfe)
	for _, _ccag := range _ebdb._fcbfe {
		_dged := _ccag._gfed[0]._faaf[0]
		if _aabg(_dged) {
			_fcgcb = append(_fcgcb, _ccag)
		}
	}
	_fcgcb = append(_fcgcb, _fdda...)
	return _fcgcb
}

type bounded interface{ bbox() _ab.PdfRectangle }

func (_acbeg gridTile) contains(_fgace _ab.PdfRectangle) bool {
	if _acbeg.numBorders() < 3 {
		return false
	}
	if _acbeg._feaa && _fgace.Llx < _acbeg.Llx-_fdbf {
		return false
	}
	if _acbeg._gcde && _fgace.Urx > _acbeg.Urx+_fdbf {
		return false
	}
	if _acbeg._bdec && _fgace.Lly < _acbeg.Lly-_fdbf {
		return false
	}
	if _acbeg._dbgf && _fgace.Ury > _acbeg.Ury+_fdbf {
		return false
	}
	return true
}

func (_abagd *structTreeRoot) parseStructTreeRoot(_dcbf _eb.PdfObject) {
	if _dcbf != nil {
		_ffbb, _bcccg := _eb.GetDict(_dcbf)
		if !_bcccg {
			_da.Log.Debug("\u0070\u0061\u0072s\u0065\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u003a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006eo\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
		}
		K := _ffbb.Get("\u004b")
		_bccd := _ffbb.Get("\u0054\u0079\u0070\u0065").String()
		var _ddcfg *_eb.PdfObjectArray
		switch _adbe := K.(type) {
		case *_eb.PdfObjectArray:
			_ddcfg = _adbe
		case *_eb.PdfObjectReference:
			_ddcfg = _eb.MakeArray(K)
		}
		_bcfc := []structElement{}
		for _, _ecga := range _ddcfg.Elements() {
			_dcga := &structElement{}
			_dcga.parseStructElement(_ecga)
			_bcfc = append(_bcfc, *_dcga)
		}
		_abagd._eede = _bcfc
		_abagd._aabe = _bccd
	}
}

func (_adgg *wordBag) text() string {
	_eaef := _adgg.allWords()
	_dgef := make([]string, len(_eaef))
	for _gadc, _gcac := range _eaef {
		_dgef[_gadc] = _gcac._faaf
	}
	return _cb.Join(_dgef, "\u0020")
}

func (_dgab *subpath) removeDuplicates() {
	if len(_dgab._bfaa) == 0 {
		return
	}
	_fcgc := []_fa.Point{_dgab._bfaa[0]}
	for _, _geaf := range _dgab._bfaa[1:] {
		if !_gagg(_geaf, _fcgc[len(_fcgc)-1]) {
			_fcgc = append(_fcgc, _geaf)
		}
	}
	_dgab._bfaa = _fcgc
}

func _facad(_edgcg []*textWord, _edfg *textWord) []*textWord {
	for _bgaff, _ggcca := range _edgcg {
		if _ggcca == _edfg {
			return _becg(_edgcg, _bgaff)
		}
	}
	_da.Log.Error("\u0072\u0065\u006d\u006f\u0076e\u0057\u006f\u0072\u0064\u003a\u0020\u0077\u006f\u0072\u0064\u0073\u0020\u0064o\u0065\u0073\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0077\u006f\u0072\u0064\u003d\u0025\u0073", _edfg)
	return nil
}
func _fadg(_dedfb, _ebedec int) uint64 { return uint64(_dedfb)*0x1000000 + uint64(_ebedec) }
func (_bgfcb lineRuling) asRuling() (*ruling, bool) {
	_aeaag := ruling{_ggcaf: _bgfcb._agbc, Color: _bgfcb.Color, _bfccg: _bgcf}
	switch _bgfcb._agbc {
	case _eccce:
		_aeaag._dbbce = _bgfcb.xMean()
		_aeaag._beadd = _g.Min(_bgfcb._bbfe.Y, _bgfcb._eaega.Y)
		_aeaag._dgbc = _g.Max(_bgfcb._bbfe.Y, _bgfcb._eaega.Y)
	case _ceedg:
		_aeaag._dbbce = _bgfcb.yMean()
		_aeaag._beadd = _g.Min(_bgfcb._bbfe.X, _bgfcb._eaega.X)
		_aeaag._dgbc = _g.Max(_bgfcb._bbfe.X, _bgfcb._eaega.X)
	default:
		_da.Log.Error("\u0062\u0061\u0064\u0020pr\u0069\u006d\u0061\u0072\u0079\u0020\u006b\u0069\u006e\u0064\u003d\u0025\u0064", _bgfcb._agbc)
		return nil, false
	}
	return &_aeaag, true
}

func (_degd *stateStack) push(_ebc *textState) {
	_bbcg := *_ebc
	*_degd = append(*_degd, &_bbcg)
}

func (_abdag paraList) toTextMarks() []TextMark {
	_ecgcg := 0
	var _agbe []TextMark
	for _fdaf, _ggfbe := range _abdag {
		if _ggfbe._ccfeb {
			continue
		}
		_dedd := _ggfbe.toTextMarks(&_ecgcg)
		_agbe = append(_agbe, _dedd...)
		if _fdaf != len(_abdag)-1 {
			if _ceed(_ggfbe, _abdag[_fdaf+1]) {
				_agbe = _beec(_agbe, &_ecgcg, "\u0020")
			} else {
				_agbe = _beec(_agbe, &_ecgcg, "\u000a")
				_agbe = _beec(_agbe, &_ecgcg, "\u000a")
			}
		}
	}
	_agbe = _beec(_agbe, &_ecgcg, "\u000a")
	_agbe = _beec(_agbe, &_ecgcg, "\u000a")
	return _agbe
}

func (_gege paraList) topoOrder() []int {
	if _bgada {
		_da.Log.Info("\u0074\u006f\u0070\u006f\u004f\u0072\u0064\u0065\u0072\u003a")
	}
	_cedag := len(_gege)
	_bdag := make([]bool, _cedag)
	_defe := make([]int, 0, _cedag)
	_ffcfc := _gege.llyOrdering()
	var _bcada func(_gaaa int)
	_bcada = func(_aceb int) {
		_bdag[_aceb] = true
		for _aegc := 0; _aegc < _cedag; _aegc++ {
			if !_bdag[_aegc] {
				if _gege.readBefore(_ffcfc, _aceb, _aegc) {
					_bcada(_aegc)
				}
			}
		}
		_defe = append(_defe, _aceb)
	}
	for _ccdf := 0; _ccdf < _cedag; _ccdf++ {
		if !_bdag[_ccdf] {
			_bcada(_ccdf)
		}
	}
	return _beecb(_defe)
}

const _ccdac = 1.0 / 1000.0

// String returns a string descibing `i`.
func (_aegfff gridTile) String() string {
	_cbgbg := func(_cdeb bool, _dffcb string) string {
		if _cdeb {
			return _dffcb
		}
		return "\u005f"
	}
	return _a.Sprintf("\u00256\u002e2\u0066\u0020\u0025\u0031\u0073%\u0031\u0073%\u0031\u0073\u0025\u0031\u0073", _aegfff.PdfRectangle, _cbgbg(_aegfff._feaa, "\u004c"), _cbgbg(_aegfff._gcde, "\u0052"), _cbgbg(_aegfff._bdec, "\u0042"), _cbgbg(_aegfff._dbgf, "\u0054"))
}

type imageExtractContext struct {
	_daae []ImageMark
	_caa  int
	_cc   int
	_dde  int
	_abg  map[*_eb.PdfObjectStream]*cachedImage
	_bd   *ImageExtractOptions
	_bce  bool
}

func (_bbgff rulingList) primaries() []float64 {
	_ebbfd := make(map[float64]struct{}, len(_bbgff))
	for _, _ebdbf := range _bbgff {
		_ebbfd[_ebdbf._dbbce] = struct{}{}
	}
	_ddba := make([]float64, len(_ebbfd))
	_ddfb := 0
	for _cgcca := range _ebbfd {
		_ddba[_ddfb] = _cgcca
		_ddfb++
	}
	_cf.Float64s(_ddba)
	return _ddba
}

func (_bcdag *textTable) emptyCompositeColumn(_ebbe int) bool {
	for _ecgcgb := 0; _ecgcgb < _bcdag._faed; _ecgcgb++ {
		if _fedad, _befb := _bcdag._fcda[_fadg(_ebbe, _ecgcgb)]; _befb {
			if len(_fedad.paraList) > 0 {
				return false
			}
		}
	}
	return true
}

type rulingKind int

func (_fagcf *textPara) taken() bool { return _fagcf == nil || _fagcf._aagd }

// Font represents the font properties on a PDF page.
type Font struct {
	PdfFont *_ab.PdfFont

	// FontName represents Font Name from font properties.
	FontName string

	// FontType represents Font Subtype entry in the font dictionary inside page resources.
	// Examples : type0, Type1, MMType1, Type3, TrueType, CIDFont.
	FontType string

	// ToUnicode is true if font provides a `ToUnicode` mapping.
	ToUnicode bool

	// IsCID is true if underlying font is a composite font.
	// Composite font is represented by a font dictionary whose Subtype is `Type0`
	IsCID bool

	// IsSimple is true if font is simple font.
	// A simple font is limited to only 8 bit (255) character codes.
	IsSimple bool

	// FontData represents the raw data of the embedded font file.
	// It can have format TrueType (TTF), PostScript Font (PFB) or Compact Font Format (CCF).
	// FontData value can be indicates from `FontFile`, `FontFile2` or `FontFile3` inside Font Descriptor.
	// At most, only one of `FontFile`, `FontFile2` or `FontFile3` will be FontData value.
	FontData []byte

	// FontFileName is a name representing the font. it has format:
	// (Font Name) + (Font Type Extension), example: helvetica.ttf.
	FontFileName string

	// FontDescriptor represents metrics and other attributes inside font properties from PDF Structure (Font Descriptor).
	FontDescriptor *_ab.PdfFontDescriptor
}

func (_cfc *textObject) renderText(_facg _eb.PdfObject, _bed []byte, _gfdd int) error {
	if _cfc._bfcc {
		_da.Log.Debug("\u0072\u0065\u006e\u0064\u0065r\u0054\u0065\u0078\u0074\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0066\u006f\u006e\u0074\u002e\u0020\u004e\u006f\u0074\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u002e")
		return nil
	}
	_eag := _cfc.getCurrentFont()
	_egaa := _eag.BytesToCharcodes(_bed)
	_dff, _fdec, _eee := _eag.CharcodesToStrings(_egaa)
	if _eee > 0 {
		_da.Log.Debug("\u0072\u0065nd\u0065\u0072\u0054e\u0078\u0074\u003a\u0020num\u0043ha\u0072\u0073\u003d\u0025\u0064\u0020\u006eum\u004d\u0069\u0073\u0073\u0065\u0073\u003d%\u0064", _fdec, _eee)
	}
	_cfc._dbfd._cbee += _fdec
	_cfc._dbfd._agec += _eee
	_daed := _cfc._dbfd
	_dgd := _daed._egbc
	_efcf := _daed._fgc / 100.0
	_cgba := _ccdac
	if _eag.Subtype() == "\u0054\u0079\u0070e\u0033" {
		_cgba = 1
	}
	_aagc, _agef := _eag.GetRuneMetrics(' ')
	if !_agef {
		_aagc, _agef = _eag.GetCharMetrics(32)
	}
	if !_agef {
		_aagc, _ = _ab.DefaultFont().GetRuneMetrics(' ')
	}
	_gaf := _aagc.Wx * _cgba
	_da.Log.Trace("\u0073p\u0061\u0063e\u0057\u0069\u0064t\u0068\u003d\u0025\u002e\u0032\u0066\u0020t\u0065\u0078\u0074\u003d\u0025\u0071 \u0066\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0066\u006f\u006et\u0053\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066", _gaf, _dff, _eag, _dgd)
	_fcgef := _fa.NewMatrix(_dgd*_efcf, 0, 0, _dgd, 0, _daed._ded)
	if _bbf {
		_da.Log.Info("\u0072\u0065\u006e\u0064\u0065\u0072T\u0065\u0078\u0074\u003a\u0020\u0025\u0064\u0020\u0063\u006f\u0064\u0065\u0073=\u0025\u002b\u0076\u0020\u0074\u0065\u0078t\u0073\u003d\u0025\u0071", len(_egaa), _egaa, _dff)
	}
	_da.Log.Trace("\u0072\u0065\u006e\u0064\u0065\u0072T\u0065\u0078\u0074\u003a\u0020\u0025\u0064\u0020\u0063\u006f\u0064\u0065\u0073=\u0025\u002b\u0076\u0020\u0072\u0075\u006ee\u0073\u003d\u0025\u0071", len(_egaa), _egaa, len(_dff))
	_geaa := _cfc.getFillColor()
	_bece := _cfc.getStrokeColor()
	for _dbdb, _fedb := range _dff {
		_aga := []rune(_fedb)
		if len(_aga) == 1 && _aga[0] == '\x00' {
			continue
		}
		_eacg := _egaa[_dbdb]
		_acce := _cfc._fcfe.CTM.Mult(_cfc._dbd).Mult(_fcgef)
		_dac := 0.0
		if len(_aga) == 1 && _aga[0] == 32 {
			_dac = _daed._ebbfc
		}
		_aefg, _agga := _eag.GetCharMetrics(_eacg)
		if !_agga {
			_da.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u004e\u006f \u006d\u0065\u0074r\u0069\u0063\u0020\u0066\u006f\u0072\u0020\u0063\u006fde\u003d\u0025\u0064 \u0072\u003d0\u0078\u0025\u0030\u0034\u0078\u003d%\u002b\u0071 \u0025\u0073", _eacg, _aga, _aga, _eag)
			return _a.Errorf("\u006e\u006f\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073:\u0020f\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0063\u006f\u0064\u0065\u003d\u0025\u0064", _eag.String(), _eacg)
		}
		_gcc := _fa.Point{X: _aefg.Wx * _cgba, Y: _aefg.Wy * _cgba}
		_dcfb := _fa.Point{X: (_gcc.X*_dgd + _dac) * _efcf}
		_babe := _fa.Point{X: (_gcc.X*_dgd + _daed._edg + _dac) * _efcf}
		if _bbf {
			_da.Log.Info("\u0074\u0066\u0073\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0063\u003d\u0025\u002e\u0032f\u0020t\u0077\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0068\u003d\u0025\u002e\u0032\u0066", _dgd, _daed._edg, _daed._ebbfc, _efcf)
			_da.Log.Info("\u0064x\u002c\u0064\u0079\u003d%\u002e\u0033\u0066\u0020\u00740\u003d%\u002e3\u0066\u0020\u0074\u003d\u0025\u002e\u0033f", _gcc, _dcfb, _babe)
		}
		_gde := _gaae(_dcfb)
		_bgf := _gaae(_babe)
		_fbba := _cfc._fcfe.CTM.Mult(_cfc._dbd).Mult(_gde)
		if _afea {
			_da.Log.Info("e\u006e\u0064\u003a\u000a\tC\u0054M\u003d\u0025\u0073\u000a\u0009 \u0074\u006d\u003d\u0025\u0073\u000a"+"\u0009\u0020t\u0064\u003d\u0025s\u0020\u0078\u006c\u0061\u0074\u003d\u0025\u0073\u000a"+"\u0009t\u0064\u0030\u003d\u0025s\u000a\u0009\u0020\u0020\u2192 \u0025s\u0020x\u006c\u0061\u0074\u003d\u0025\u0073", _cfc._fcfe.CTM, _cfc._dbd, _bgf, _ecefg(_cfc._fcfe.CTM.Mult(_cfc._dbd).Mult(_bgf)), _gde, _fbba, _ecefg(_fbba))
		}
		_feb, _cgdd := _cfc.newTextMark(_ba.ExpandLigatures(_aga), _acce, _ecefg(_fbba), _g.Abs(_gaf*_acce.ScalingFactorX()), _eag, _cfc._dbfd._edg, _geaa, _bece, _facg, _dff, _dbdb, _gfdd)
		if !_cgdd {
			_da.Log.Debug("\u0054\u0065\u0078\u0074\u0020\u006d\u0061\u0072\u006b\u0020\u006f\u0075\u0074\u0073\u0069d\u0065 \u0070\u0061\u0067\u0065\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		if _eag == nil {
			_da.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u004e\u006f\u0020\u0066\u006f\u006e\u0074\u002e")
		} else if _eag.Encoder() == nil {
			_da.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020N\u006f\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006eg\u002e\u0020\u0066o\u006et\u003d\u0025\u0073", _eag)
		} else {
			if _bff, _ggg := _eag.Encoder().CharcodeToRune(_eacg); _ggg {
				_feb._cage = string(_bff)
			}
		}
		_da.Log.Trace("i\u003d\u0025\u0064\u0020\u0063\u006fd\u0065\u003d\u0025\u0064\u0020\u006d\u0061\u0072\u006b=\u0025\u0073\u0020t\u0072m\u003d\u0025\u0073", _dbdb, _eacg, _feb, _acce)
		_cfc._cecd = append(_cfc._cecd, &_feb)
		_cfc._dbd.Concat(_bgf)
	}
	return nil
}

func (_abggf *ruling) alignsPrimary(_agefb *ruling) bool {
	return _abggf._ggcaf == _agefb._ggcaf && _g.Abs(_abggf._dbbce-_agefb._dbbce) < _efbg*0.5
}
func (_aeg *wordBag) minDepth() float64 { return _aeg._bfgd - (_aeg.Ury - _aeg._caag) }
func (_acff rulingList) vertsHorzs() (rulingList, rulingList) {
	var _acfda, _gafg rulingList
	for _, _bbgca := range _acff {
		switch _bbgca._ggcaf {
		case _eccce:
			_acfda = append(_acfda, _bbgca)
		case _ceedg:
			_gafg = append(_gafg, _bbgca)
		}
	}
	return _acfda, _gafg
}

func (_gfgg rulingList) findPrimSec(_fcgg, _bgcg float64) *ruling {
	for _, _ebgb := range _gfgg {
		if _ccae(_ebgb._dbbce-_fcgg) && _ebgb._beadd-_gded <= _bgcg && _bgcg <= _ebgb._dgbc+_gded {
			return _ebgb
		}
	}
	return nil
}

func (_bedcd *textTable) logComposite(_cdbd string) {
	if !_cabbg {
		return
	}
	_da.Log.Info("\u007e~\u007eP\u0061\u0072\u0061\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0073", _bedcd._cbcb, _bedcd._faed, _cdbd)
	_a.Printf("\u0025\u0035\u0073 \u007c", "")
	for _befbc := 0; _befbc < _bedcd._cbcb; _befbc++ {
		_a.Printf("\u0025\u0033\u0064 \u007c", _befbc)
	}
	_a.Println("")
	_a.Printf("\u0025\u0035\u0073 \u002b", "")
	for _gefea := 0; _gefea < _bedcd._cbcb; _gefea++ {
		_a.Printf("\u0025\u0033\u0073 \u002b", "\u002d\u002d\u002d")
	}
	_a.Println("")
	for _afeb := 0; _afeb < _bedcd._faed; _afeb++ {
		_a.Printf("\u0025\u0035\u0064 \u007c", _afeb)
		for _agdcb := 0; _agdcb < _bedcd._cbcb; _agdcb++ {
			_eebc, _ := _bedcd._fcda[_fadg(_agdcb, _afeb)].parasBBox()
			_a.Printf("\u0025\u0033\u0064 \u007c", len(_eebc))
		}
		_a.Println("")
	}
	_da.Log.Info("\u007e~\u007eT\u0065\u0078\u0074\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0073", _bedcd._cbcb, _bedcd._faed, _cdbd)
	_a.Printf("\u0025\u0035\u0073 \u007c", "")
	for _gdde := 0; _gdde < _bedcd._cbcb; _gdde++ {
		_a.Printf("\u0025\u0031\u0032\u0064\u0020\u007c", _gdde)
	}
	_a.Println("")
	_a.Printf("\u0025\u0035\u0073 \u002b", "")
	for _cfbbf := 0; _cfbbf < _bedcd._cbcb; _cfbbf++ {
		_a.Print("\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d-\u002d\u002d\u002d\u002b")
	}
	_a.Println("")
	for _eagf := 0; _eagf < _bedcd._faed; _eagf++ {
		_a.Printf("\u0025\u0035\u0064 \u007c", _eagf)
		for _dfga := 0; _dfga < _bedcd._cbcb; _dfga++ {
			_cacag, _ := _bedcd._fcda[_fadg(_dfga, _eagf)].parasBBox()
			_ebedc := ""
			_dcgdc := _cacag.merge()
			if _dcgdc != nil {
				_ebedc = _dcgdc.text()
			}
			_ebedc = _a.Sprintf("\u0025\u0071", _gggg(_ebedc, 12))
			_ebedc = _ebedc[1 : len(_ebedc)-1]
			_a.Printf("\u0025\u0031\u0032\u0073\u0020\u007c", _ebedc)
		}
		_a.Println("")
	}
}

func (_gdbc *shapesState) drawRectangle(_agb, _bcc, _dfdf, _fabfed float64) {
	if _caad {
		_bgffg := _gdbc.devicePoint(_agb, _bcc)
		_ebe := _gdbc.devicePoint(_agb+_dfdf, _bcc+_fabfed)
		_affg := _ab.PdfRectangle{Llx: _bgffg.X, Lly: _bgffg.Y, Urx: _ebe.X, Ury: _ebe.Y}
		_da.Log.Info("d\u0072a\u0077\u0052\u0065\u0063\u0074\u0061\u006e\u0067l\u0065\u003a\u0020\u00256.\u0032\u0066", _affg)
	}
	_gdbc.newSubPath()
	_gdbc.moveTo(_agb, _bcc)
	_gdbc.lineTo(_agb+_dfdf, _bcc)
	_gdbc.lineTo(_agb+_dfdf, _bcc+_fabfed)
	_gdbc.lineTo(_agb, _bcc+_fabfed)
	_gdbc.closePath()
}

func _ggcg(_cggf float64, _fggg int) int {
	if _fggg == 0 {
		_fggg = 1
	}
	_gfee := float64(_fggg)
	return int(_g.Round(_cggf/_gfee) * _gfee)
}

// Len returns the number of TextMarks in `ma`.
func (_agac *TextMarkArray) Len() int {
	if _agac == nil {
		return 0
	}
	return len(_agac._beaf)
}

func (_dfa *shapesState) stroke(_feae *[]pathSection) {
	_beeg := pathSection{_egfg: _dfa._dagf, Color: _dfa._dfff.getStrokeColor()}
	*_feae = append(*_feae, _beeg)
	if _egc {
		_a.Printf("\u0020 \u0020\u0020S\u0054\u0052\u004fK\u0045\u003a\u0020\u0025\u0064\u0020\u0073t\u0072\u006f\u006b\u0065\u0073\u0020s\u0073\u003d\u0025\u0073\u0020\u0063\u006f\u006c\u006f\u0072\u003d%\u002b\u0076\u0020\u0025\u0036\u002e\u0032\u0066\u000a", len(*_feae), _dfa, _dfa._dfff.getStrokeColor(), _beeg.bbox())
		if _gcbad {
			for _cgdg, _egga := range _dfa._dagf {
				_a.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _cgdg, _egga)
				if _cgdg == 10 {
					break
				}
			}
		}
	}
}
func (_bdabd *textWord) bbox() _ab.PdfRectangle { return _bdabd.PdfRectangle }
func _ecfd(_dbbe []*textLine) map[float64][]*textLine {
	_cf.Slice(_dbbe, func(_gbdf, _gccab int) bool { return _dbbe[_gbdf]._defc < _dbbe[_gccab]._defc })
	_ffeg := map[float64][]*textLine{}
	for _, _gbgab := range _dbbe {
		_fedfg := _fcbf(_gbgab)
		_fedfg = _g.Round(_fedfg)
		_ffeg[_fedfg] = append(_ffeg[_fedfg], _gbgab)
	}
	return _ffeg
}

func (_cbda paraList) readBefore(_eaebb []int, _ddag, _dbade int) bool {
	_afadb, _egfcg := _cbda[_ddag], _cbda[_dbade]
	if _gecd(_afadb, _egfcg) && _afadb.Lly > _egfcg.Lly {
		return true
	}
	if !(_afadb._gedc.Urx < _egfcg._gedc.Llx) {
		return false
	}
	_cdea, _gece := _afadb.Lly, _egfcg.Lly
	if _cdea > _gece {
		_gece, _cdea = _cdea, _gece
	}
	_bcea := _g.Max(_afadb._gedc.Llx, _egfcg._gedc.Llx)
	_aedda := _g.Min(_afadb._gedc.Urx, _egfcg._gedc.Urx)
	_aeeac := _cbda.llyRange(_eaebb, _cdea, _gece)
	for _, _ddae := range _aeeac {
		if _ddae == _ddag || _ddae == _dbade {
			continue
		}
		_cgfe := _cbda[_ddae]
		if _cgfe._gedc.Llx <= _aedda && _bcea <= _cgfe._gedc.Urx {
			return false
		}
	}
	return true
}

type textState struct {
	_edg   float64
	_ebbfc float64
	_fgc   float64
	_abda  float64
	_egbc  float64
	_cdf   RenderMode
	_ded   float64
	_ggda  *_ab.PdfFont
	_cee   _ab.PdfRectangle
	_cbee  int
	_agec  int
}

func (_bdb *textLine) bbox() _ab.PdfRectangle { return _bdb.PdfRectangle }
func (_baeb *textMark) inDiacriticArea(_debg *textMark) bool {
	_cfa := _baeb.Llx - _debg.Llx
	_fcfc := _baeb.Urx - _debg.Urx
	_gdbgb := _baeb.Lly - _debg.Lly
	return _g.Abs(_cfa+_fcfc) < _baeb.Width()*_agdg && _g.Abs(_gdbgb) < _baeb.Height()*_agdg
}

func (_afee *textObject) showText(_dgc _eb.PdfObject, _ddf []byte, _faea int) error {
	return _afee.renderText(_dgc, _ddf, _faea)
}

func (_dfbg paraList) applyTables(_ggee []*textTable) paraList {
	var _fbddd paraList
	for _, _daaagg := range _ggee {
		_fbddd = append(_fbddd, _daaagg.newTablePara())
	}
	for _, _bgcac := range _dfbg {
		if _bgcac._aagd {
			continue
		}
		_fbddd = append(_fbddd, _bgcac)
	}
	return _fbddd
}

func (_eaac paraList) computeEBBoxes() {
	if _cfdb {
		_da.Log.Info("\u0063o\u006dp\u0075\u0074\u0065\u0045\u0042\u0042\u006f\u0078\u0065\u0073\u003a")
	}
	for _, _gbcg := range _eaac {
		_gbcg._gedc = _gbcg.PdfRectangle
	}
	_fedc := _eaac.yNeighbours(0)
	for _fdeg, _caea := range _eaac {
		_daga := _caea._gedc
		_baba, _feeca := -1.0e9, +1.0e9
		for _, _adab := range _fedc[_caea] {
			_bafa := _eaac[_adab]._gedc
			if _bafa.Urx < _daga.Llx {
				_baba = _g.Max(_baba, _bafa.Urx)
			} else if _daga.Urx < _bafa.Llx {
				_feeca = _g.Min(_feeca, _bafa.Llx)
			}
		}
		for _cadb, _gcdc := range _eaac {
			_ceba := _gcdc._gedc
			if _fdeg == _cadb || _ceba.Ury > _daga.Lly {
				continue
			}
			if _baba <= _ceba.Llx && _ceba.Llx < _daga.Llx {
				_daga.Llx = _ceba.Llx
			} else if _ceba.Urx <= _feeca && _daga.Urx < _ceba.Urx {
				_daga.Urx = _ceba.Urx
			}
		}
		if _cfdb {
			_a.Printf("\u0025\u0034\u0064\u003a %\u0036\u002e\u0032\u0066\u2192\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0071\u000a", _fdeg, _caea._gedc, _daga, _gggg(_caea.text(), 50))
		}
		_caea._gedc = _daga
	}
	if _fdfc {
		for _, _fbfd := range _eaac {
			_fbfd.PdfRectangle = _fbfd._gedc
		}
	}
}

type wordBag struct {
	_ab.PdfRectangle
	_caag       float64
	_bbcd, _bae rulingList
	_bfgd       float64
	_edgf       map[int][]*textWord
}

func _fefa(_afdb string) (string, bool) {
	_afaeb := []rune(_afdb)
	if len(_afaeb) != 1 {
		return "", false
	}
	_cgdbf, _fgggd := _gedg[_afaeb[0]]
	return _cgdbf, _fgggd
}

func _eabcb(_cfac []*textMark, _abad _ab.PdfRectangle) []*textWord {
	var _dafab []*textWord
	var _gbgf *textWord
	if _fbbg {
		_da.Log.Info("\u006d\u0061\u006beT\u0065\u0078\u0074\u0057\u006f\u0072\u0064\u0073\u003a\u0020\u0025\u0064\u0020\u006d\u0061\u0072\u006b\u0073", len(_cfac))
	}
	_bgcec := func() {
		if _gbgf != nil {
			_cgbff := _gbgf.computeText()
			if !_fcacf(_cgbff) {
				_gbgf._faaf = _cgbff
				_dafab = append(_dafab, _gbgf)
				if _fbbg {
					_da.Log.Info("\u0061\u0064\u0064Ne\u0077\u0057\u006f\u0072\u0064\u003a\u0020\u0025\u0064\u003a\u0020\u0077\u006f\u0072\u0064\u003d\u0025\u0073", len(_dafab)-1, _gbgf.String())
					for _dcffe, _faged := range _gbgf._bacfd {
						_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _dcffe, _faged.String())
					}
				}
			}
			_gbgf = nil
		}
	}
	for _, _bgbfe := range _cfac {
		if _becd && _gbgf != nil && len(_gbgf._bacfd) > 0 {
			_gfadc := _gbgf._bacfd[len(_gbgf._bacfd)-1]
			_baeca, _decf := _fefa(_bgbfe._bbeb)
			_fcacd, _efgef := _fefa(_gfadc._bbeb)
			if _decf && !_efgef && _gfadc.inDiacriticArea(_bgbfe) {
				_gbgf.addDiacritic(_baeca)
				continue
			}
			if _efgef && !_decf && _bgbfe.inDiacriticArea(_gfadc) {
				_gbgf._bacfd = _gbgf._bacfd[:len(_gbgf._bacfd)-1]
				_gbgf.appendMark(_bgbfe, _abad)
				_gbgf.addDiacritic(_fcacd)
				continue
			}
		}
		_abffd := _fcacf(_bgbfe._bbeb)
		if _abffd {
			_bgcec()
			continue
		}
		if _gbgf == nil && !_abffd {
			_gbgf = _adccb([]*textMark{_bgbfe}, _abad)
			continue
		}
		_acac := _gbgf._fcde
		_fabe := _g.Abs(_ddbg(_abad, _bgbfe)-_gbgf._faace) / _acac
		_cgeeg := _cfba(_bgbfe, _gbgf) / _acac
		if _cgeeg >= _aded || !(-_egdg <= _cgeeg && _fabe <= _aca) {
			_bgcec()
			_gbgf = _adccb([]*textMark{_bgbfe}, _abad)
			continue
		}
		_gbgf.appendMark(_bgbfe, _abad)
	}
	_bgcec()
	return _dafab
}

type list struct {
	_dbfeb []*textLine
	_dbga  string
	_ffaf  []*list
	_gfec  string
}

// RenderMode specifies the text rendering mode (Tmode), which determines whether showing text shall cause
// glyph outlines to be  stroked, filled, used as a clipping boundary, or some combination of the three.
// Stroking, filling, and clipping shall have the same effects for a text object as they do for a path object
// (see 8.5.3, "Path-Painting Operators" and 8.5.4, "Clipping Path Operators").
type RenderMode int

func _bbaa(_caagc, _ecgca bounded) float64 { return _gdeb(_caagc) - _gdeb(_ecgca) }

type pathSection struct {
	_egfg []*subpath
	_ca.Color
}

func (_ebgf *textPara) writeText(_ffafc _f.Writer) {
	if _ebgf._bcfcf == nil {
		_ebgf.writeCellText(_ffafc)
		return
	}
	for _eefba := 0; _eefba < _ebgf._bcfcf._faed; _eefba++ {
		for _ddcg := 0; _ddcg < _ebgf._bcfcf._cbcb; _ddcg++ {
			_ebgd := _ebgf._bcfcf.get(_ddcg, _eefba)
			if _ebgd == nil {
				_ffafc.Write([]byte("\u0009"))
			} else {
				_ebgd.writeCellText(_ffafc)
			}
			_ffafc.Write([]byte("\u0020"))
		}
		if _eefba < _ebgf._bcfcf._faed-1 {
			_ffafc.Write([]byte("\u000a"))
		}
	}
}

func (_afaea gridTile) numBorders() int {
	_cege := 0
	if _afaea._feaa {
		_cege++
	}
	if _afaea._gcde {
		_cege++
	}
	if _afaea._bdec {
		_cege++
	}
	if _afaea._dbgf {
		_cege++
	}
	return _cege
}

func (_ebae compositeCell) String() string {
	_afbc := ""
	if len(_ebae.paraList) > 0 {
		_afbc = _gggg(_ebae.paraList.merge().text(), 50)
	}
	return _a.Sprintf("\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0064\u0020\u0070\u0061\u0072a\u0073\u0020\u0025\u0071", _ebae.PdfRectangle, len(_ebae.paraList), _afbc)
}

func (_fed *textObject) showTextAdjusted(_bad *_eb.PdfObjectArray, _edcg int) error {
	_eeda := false
	for _, _fccf := range _bad.Elements() {
		switch _fccf.(type) {
		case *_eb.PdfObjectFloat, *_eb.PdfObjectInteger:
			_caae, _afd := _eb.GetNumberAsFloat(_fccf)
			if _afd != nil {
				_da.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0073\u0068\u006f\u0077\u0054\u0065\u0078t\u0041\u0064\u006a\u0075\u0073\u0074\u0065\u0064\u002e\u0020\u0042\u0061\u0064\u0020\u006e\u0075\u006d\u0065r\u0069\u0063\u0061\u006c\u0020a\u0072\u0067\u002e\u0020\u006f\u003d\u0025\u0073\u0020\u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _fccf, _bad)
				return _afd
			}
			_ebaa, _eae := -_caae*0.001*_fed._dbfd._egbc, 0.0
			if _eeda {
				_eae, _ebaa = _ebaa, _eae
			}
			_eacb := _gaae(_fa.Point{X: _ebaa, Y: _eae})
			_fed._dbd.Concat(_eacb)
		case *_eb.PdfObjectString:
			_bcbe := _eb.TraceToDirectObject(_fccf)
			_adg, _fbfg := _eb.GetStringBytes(_bcbe)
			if !_fbfg {
				_da.Log.Trace("s\u0068\u006f\u0077\u0054\u0065\u0078\u0074\u0041\u0064j\u0075\u0073\u0074\u0065\u0064\u003a\u0020Ba\u0064\u0020\u0073\u0074r\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u002e\u0020o=\u0025\u0073 \u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _fccf, _bad)
				return _eb.ErrTypeError
			}
			_fed.renderText(_bcbe, _adg, _edcg)
		default:
			_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0073\u0068\u006f\u0077\u0054\u0065\u0078\u0074A\u0064\u006a\u0075\u0073\u0074\u0065\u0064\u002e\u0020\u0055\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0028%T\u0029\u0020\u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _fccf, _bad)
			return _eb.ErrTypeError
		}
	}
	return nil
}

func (_gfaf *textObject) setTextRenderMode(_dgbb int) {
	if _gfaf == nil {
		return
	}
	_gfaf._dbfd._cdf = RenderMode(_dgbb)
}

func (_dfb rectRuling) asRuling() (*ruling, bool) {
	_gabf := ruling{_ggcaf: _dfb._fdba, Color: _dfb.Color, _bfccg: _dagaa}
	switch _dfb._fdba {
	case _eccce:
		_gabf._dbbce = 0.5 * (_dfb.Llx + _dfb.Urx)
		_gabf._beadd = _dfb.Lly
		_gabf._dgbc = _dfb.Ury
		_efea, _ddffc := _dfb.checkWidth(_dfb.Llx, _dfb.Urx)
		if !_ddffc {
			if _gddab {
				_da.Log.Error("\u0072\u0065\u0063\u0074\u0052\u0075l\u0069\u006e\u0067\u002e\u0061\u0073\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0072\u0075\u006c\u0069\u006e\u0067V\u0065\u0072\u0074\u0020\u0021\u0063\u0068\u0065\u0063\u006b\u0057\u0069\u0064\u0074h\u0020v\u003d\u0025\u002b\u0076", _dfb)
			}
			return nil, false
		}
		_gabf._fbfc = _efea
	case _ceedg:
		_gabf._dbbce = 0.5 * (_dfb.Lly + _dfb.Ury)
		_gabf._beadd = _dfb.Llx
		_gabf._dgbc = _dfb.Urx
		_eeccd, _gddgf := _dfb.checkWidth(_dfb.Lly, _dfb.Ury)
		if !_gddgf {
			if _gddab {
				_da.Log.Error("\u0072\u0065\u0063\u0074\u0052\u0075l\u0069\u006e\u0067\u002e\u0061\u0073\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0072\u0075\u006c\u0069\u006e\u0067H\u006f\u0072\u007a\u0020\u0021\u0063\u0068\u0065\u0063\u006b\u0057\u0069\u0064\u0074h\u0020v\u003d\u0025\u002b\u0076", _dfb)
			}
			return nil, false
		}
		_gabf._fbfc = _eeccd
	default:
		_da.Log.Error("\u0062\u0061\u0064\u0020pr\u0069\u006d\u0061\u0072\u0079\u0020\u006b\u0069\u006e\u0064\u003d\u0025\u0064", _dfb._fdba)
		return nil, false
	}
	return &_gabf, true
}

func (_dgeee rulingList) splitSec() []rulingList {
	_cf.Slice(_dgeee, func(_ffabc, _efff int) bool {
		_bfceg, _begg := _dgeee[_ffabc], _dgeee[_efff]
		if _bfceg._beadd != _begg._beadd {
			return _bfceg._beadd < _begg._beadd
		}
		return _bfceg._dgbc < _begg._dgbc
	})
	_effb := make(map[*ruling]struct{}, len(_dgeee))
	_gdfca := func(_ceafd *ruling) rulingList {
		_bfbbb := rulingList{_ceafd}
		_effb[_ceafd] = struct{}{}
		for _, _ffgfa := range _dgeee {
			if _, _acfbb := _effb[_ffgfa]; _acfbb {
				continue
			}
			for _, _cadf := range _bfbbb {
				if _ffgfa.alignsSec(_cadf) {
					_bfbbb = append(_bfbbb, _ffgfa)
					_effb[_ffgfa] = struct{}{}
					break
				}
			}
		}
		return _bfbbb
	}
	_bbdb := []rulingList{_gdfca(_dgeee[0])}
	for _, _gdcd := range _dgeee[1:] {
		if _, _agafa := _effb[_gdcd]; _agafa {
			continue
		}
		_bbdb = append(_bbdb, _gdfca(_gdcd))
	}
	return _bbdb
}

func (_feaf *textTable) getComposite(_cegff, _gfded int) (paraList, _ab.PdfRectangle) {
	_gfdeb, _aabga := _feaf._fcda[_fadg(_cegff, _gfded)]
	if _cabbg {
		_a.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0067\u0065\u0074\u0043\u006f\u006d\u0070o\u0073i\u0074\u0065\u0028\u0025\u0064\u002c\u0025\u0064\u0029\u002d\u003e\u0025\u0073\u000a", _cegff, _gfded, _gfdeb.String())
	}
	if !_aabga {
		return nil, _ab.PdfRectangle{}
	}
	return _gfdeb.parasBBox()
}

func (_fdg *shapesState) closePath() {
	if _fdg._caef {
		_fdg._dagf = append(_fdg._dagf, _cbc(_fdg._beaa))
		_fdg._caef = false
	} else if len(_fdg._dagf) == 0 {
		if _caad {
			_da.Log.Debug("\u0063\u006c\u006f\u0073eP\u0061\u0074\u0068\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f\u0020\u0070\u0061t\u0068")
		}
		_fdg._caef = false
		return
	}
	_fdg._dagf[len(_fdg._dagf)-1].close()
	if _caad {
		_da.Log.Info("\u0063\u006c\u006f\u0073\u0065\u0050\u0061\u0074\u0068\u003a\u0020\u0025\u0073", _fdg)
	}
}

func (_dgbe *structElement) parseStructElement(_fbgb _eb.PdfObject) {
	_gge, _ffd := _eb.GetDict(_fbgb)
	if !_ffd {
		_da.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u0053\u0074\u0072u\u0063\u0074\u0045le\u006d\u0065\u006e\u0074\u003a\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u002e")
		return
	}
	_aagf := _gge.Get("\u0053")
	_dbdd := _gge.Get("\u0050\u0067")
	_cgca := ""
	if _aagf != nil {
		_cgca = _aagf.String()
	}
	_dbcg := _gge.Get("\u004b")
	_dgbe._acfd = _cgca
	_dgbe._ffcf = _dbdd
	switch _abga := _dbcg.(type) {
	case *_eb.PdfObjectInteger:
		_dgbe._acfd = _cgca
		_dgbe._bbgg = int64(*_abga)
		_dgbe._ffcf = _dbdd
	case *_eb.PdfObjectReference:
		_gefe := *_eb.MakeArray(_abga)
		var _bege int64 = -1
		_dgbe._bbgg = _bege
		if _gefe.Len() == 1 {
			_bgfa := _gefe.Elements()[0]
			_bfcb, _egfc := _bgfa.(*_eb.PdfObjectInteger)
			if _egfc {
				_bege = int64(*_bfcb)
				_dgbe._bbgg = _bege
				_dgbe._acfd = _cgca
				_dgbe._ffcf = _dbdd
				return
			}
		}
		_beab := []structElement{}
		for _, _eefda := range _gefe.Elements() {
			_aacf, _bbgc := _eefda.(*_eb.PdfObjectInteger)
			if _bbgc {
				_bege = int64(*_aacf)
				_dgbe._bbgg = _bege
				_dgbe._acfd = _cgca
			} else {
				_afaeg := &structElement{}
				_afaeg.parseStructElement(_eefda)
				_beab = append(_beab, *_afaeg)
			}
			_bege = -1
		}
		_dgbe._baac = _beab
	case *_eb.PdfObjectArray:
		_gbbc := _dbcg.(*_eb.PdfObjectArray)
		var _gfgf int64 = -1
		_dgbe._bbgg = _gfgf
		if _gbbc.Len() == 1 {
			_bfce := _gbbc.Elements()[0]
			_gaadg, _dfaf := _bfce.(*_eb.PdfObjectInteger)
			if _dfaf {
				_gfgf = int64(*_gaadg)
				_dgbe._bbgg = _gfgf
				_dgbe._acfd = _cgca
				_dgbe._ffcf = _dbdd
				return
			}
		}
		_aeca := []structElement{}
		for _, _gggb := range _gbbc.Elements() {
			_bbde, _bfca := _gggb.(*_eb.PdfObjectInteger)
			if _bfca {
				_gfgf = int64(*_bbde)
				_dgbe._bbgg = _gfgf
				_dgbe._acfd = _cgca
				_dgbe._ffcf = _dbdd
			} else {
				_beag := &structElement{}
				_beag.parseStructElement(_gggb)
				_aeca = append(_aeca, *_beag)
			}
			_gfgf = -1
		}
		_dgbe._baac = _aeca
	}
}

// NewWithOptions an Extractor instance for extracting content from the input PDF page with options.
func NewWithOptions(page *_ab.PdfPage, options *Options) (*Extractor, error) {
	const _efa = "\u0065x\u0074\u0072\u0061\u0063\u0074\u006f\u0072\u002e\u004e\u0065\u0077W\u0069\u0074\u0068\u004f\u0070\u0074\u0069\u006f\u006e\u0073"
	_fae, _cff := page.GetAllContentStreams()
	if _cff != nil {
		return nil, _cff
	}
	_fb, _edc := page.GetStructTreeRoot()
	if !_edc {
		_da.Log.Info("T\u0068\u0065\u0020\u0070\u0064\u0066\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0074\u0061\u0067g\u0065d\u002e\u0020\u0053\u0074r\u0075\u0063t\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020\u0065\u0078\u0069\u0073\u0074\u002e")
	}
	_ee := page.GetContainingPdfObject()
	_de, _cff := page.GetMediaBox()
	if _cff != nil {
		return nil, _a.Errorf("\u0065\u0078\u0074r\u0061\u0063\u0074\u006fr\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u002e\u0020\u0025\u0076", _cff)
	}
	_daa := &Extractor{_ef: _fae, _fgf: page.Resources, _efe: *_de, _fcf: page.CropBox, _eg: map[string]fontEntry{}, _fag: map[string]textResult{}, _bbe: options, _egg: _fb, _gb: _ee}
	if _daa._efe.Llx > _daa._efe.Urx {
		_da.Log.Info("\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u0068\u0061\u0073\u0020\u0058\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0073\u0020r\u0065\u0076\u0065\u0072\u0073\u0065\u0064\u002e\u0020\u0025\u002e\u0032\u0066\u0020F\u0069x\u0069\u006e\u0067\u002e", _daa._efe)
		_daa._efe.Llx, _daa._efe.Urx = _daa._efe.Urx, _daa._efe.Llx
	}
	if _daa._efe.Lly > _daa._efe.Ury {
		_da.Log.Info("\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u0068\u0061\u0073\u0020\u0059\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0073\u0020r\u0065\u0076\u0065\u0072\u0073\u0065\u0064\u002e\u0020\u0025\u002e\u0032\u0066\u0020F\u0069x\u0069\u006e\u0067\u002e", _daa._efe)
		_daa._efe.Lly, _daa._efe.Ury = _daa._efe.Ury, _daa._efe.Lly
	}
	return _daa, nil
}

func (_ggba rulingList) intersections() map[int]intSet {
	var _dabd, _cbdc []int
	for _gdaf, _eddg := range _ggba {
		switch _eddg._ggcaf {
		case _eccce:
			_dabd = append(_dabd, _gdaf)
		case _ceedg:
			_cbdc = append(_cbdc, _gdaf)
		}
	}
	if len(_dabd) < _fafe+1 || len(_cbdc) < _ced+1 {
		return nil
	}
	if len(_dabd)+len(_cbdc) > _edcf {
		_da.Log.Debug("\u0069\u006e\u0074\u0065\u0072\u0073e\u0063\u0074\u0069\u006f\u006e\u0073\u003a\u0020\u0054\u004f\u004f\u0020\u004d\u0041\u004e\u0059\u0020\u0072\u0075\u006ci\u006e\u0067\u0073\u0020\u0076\u0065\u0063\u0073\u003d\u0025\u0064\u0020\u003d\u0020%\u0064 \u0078\u0020\u0025\u0064", len(_ggba), len(_dabd), len(_cbdc))
		return nil
	}
	_bacfa := make(map[int]intSet, len(_dabd)+len(_cbdc))
	for _, _dfad := range _dabd {
		for _, _gefee := range _cbdc {
			if _ggba[_dfad].intersects(_ggba[_gefee]) {
				if _, _cffc := _bacfa[_dfad]; !_cffc {
					_bacfa[_dfad] = make(intSet)
				}
				if _, _fafd := _bacfa[_gefee]; !_fafd {
					_bacfa[_gefee] = make(intSet)
				}
				_bacfa[_dfad].add(_gefee)
				_bacfa[_gefee].add(_dfad)
			}
		}
	}
	return _bacfa
}

func _agadc(_acdgc, _bcegf float64) string {
	_fadac := !_ccae(_acdgc - _bcegf)
	if _fadac {
		return "\u000a"
	}
	return "\u0020"
}

func (_acebb rulingList) toTilings() (rulingList, []gridTiling) {
	_acebb.log("\u0074o\u0054\u0069\u006c\u0069\u006e\u0067s")
	if len(_acebb) == 0 {
		return nil, nil
	}
	_acebb = _acebb.tidied("\u0061\u006c\u006c")
	_acebb.log("\u0074\u0069\u0064\u0069\u0065\u0064")
	_ffba := _acebb.toGrids()
	_caebd := make([]gridTiling, len(_ffba))
	for _gbgcf, _bcee := range _ffba {
		_caebd[_gbgcf] = _bcee.asTiling()
	}
	return _acebb, _caebd
}

func (_bdcfa compositeCell) split(_cadce, _aebf []float64) *textTable {
	_efba := len(_cadce) + 1
	_bbfa := len(_aebf) + 1
	if _cabbg {
		_da.Log.Info("\u0063\u006f\u006d\u0070\u006f\u0073\u0069t\u0065\u0043\u0065l\u006c\u002e\u0073\u0070l\u0069\u0074\u003a\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u000a\u0009\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u003d\u0025\u0073\u000a"+"\u0009\u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073=\u0025\u0036\u002e\u0032\u0066\u000a\t\u0063\u006f\u006c\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u003d%\u0036\u002e\u0032\u0066", _bbfa, _efba, _bdcfa, _cadce, _aebf)
		_a.Printf("\u0020\u0020\u0020\u0020\u0025\u0064\u0020\u0070\u0061\u0072\u0061\u0073\u000a", len(_bdcfa.paraList))
		for _aagb, _aaff := range _bdcfa.paraList {
			_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _aagb, _aaff.String())
		}
		_a.Printf("\u0020\u0020\u0020\u0020\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u000a", len(_bdcfa.lines()))
		for _dggg, _adef := range _bdcfa.lines() {
			_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _dggg, _adef)
		}
	}
	_cadce = _bfef(_cadce, _bdcfa.Ury, _bdcfa.Lly)
	_aebf = _bfef(_aebf, _bdcfa.Llx, _bdcfa.Urx)
	_gcbfb := make(map[uint64]*textPara, _bbfa*_efba)
	_fbd := textTable{_cbcb: _bbfa, _faed: _efba, _ebgbd: _gcbfb}
	_fcgdf := _bdcfa.paraList
	_cf.Slice(_fcgdf, func(_eeafd, _gccg int) bool {
		_agba, _bdgb := _fcgdf[_eeafd], _fcgdf[_gccg]
		_cebf, _efad := _agba.Lly, _bdgb.Lly
		if _cebf != _efad {
			return _cebf < _efad
		}
		return _agba.Llx < _bdgb.Llx
	})
	_dcgg := make(map[uint64]_ab.PdfRectangle, _bbfa*_efba)
	for _cagf, _aebbd := range _cadce[1:] {
		_eebdd := _cadce[_cagf]
		for _cdgc, _agcb := range _aebf[1:] {
			_geff := _aebf[_cdgc]
			_dcgg[_fadg(_cdgc, _cagf)] = _ab.PdfRectangle{Llx: _geff, Urx: _agcb, Lly: _aebbd, Ury: _eebdd}
		}
	}
	if _cabbg {
		_da.Log.Info("\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0043\u0065l\u006c\u002e\u0073\u0070\u006c\u0069\u0074\u003a\u0020\u0072e\u0063\u0074\u0073")
		_a.Printf("\u0020\u0020\u0020\u0020")
		for _bcafc := 0; _bcafc < _bbfa; _bcafc++ {
			_a.Printf("\u0025\u0033\u0030\u0064\u002c\u0020", _bcafc)
		}
		_a.Println()
		for _ddaa := 0; _ddaa < _efba; _ddaa++ {
			_a.Printf("\u0020\u0020\u0025\u0032\u0064\u003a", _ddaa)
			for _bdgbe := 0; _bdgbe < _bbfa; _bdgbe++ {
				_a.Printf("\u00256\u002e\u0032\u0066\u002c\u0020", _dcgg[_fadg(_bdgbe, _ddaa)])
			}
			_a.Println()
		}
	}
	_ccfab := func(_dedfe *textLine) (int, int) {
		for _afbgf := 0; _afbgf < _efba; _afbgf++ {
			for _aaac := 0; _aaac < _bbfa; _aaac++ {
				if _bgeg(_dcgg[_fadg(_aaac, _afbgf)], _dedfe.PdfRectangle) {
					return _aaac, _afbgf
				}
			}
		}
		return -1, -1
	}
	_adebb := make(map[uint64][]*textLine, _bbfa*_efba)
	for _, _fegdd := range _fcgdf.lines() {
		_abfec, _febe := _ccfab(_fegdd)
		if _abfec < 0 {
			continue
		}
		_adebb[_fadg(_abfec, _febe)] = append(_adebb[_fadg(_abfec, _febe)], _fegdd)
	}
	for _fddcg := 0; _fddcg < len(_cadce)-1; _fddcg++ {
		_acaf := _cadce[_fddcg]
		_daad := _cadce[_fddcg+1]
		for _cfed := 0; _cfed < len(_aebf)-1; _cfed++ {
			_fcfgb := _aebf[_cfed]
			_adea := _aebf[_cfed+1]
			_cfdff := _ab.PdfRectangle{Llx: _fcfgb, Urx: _adea, Lly: _daad, Ury: _acaf}
			_acfde := _adebb[_fadg(_cfed, _fddcg)]
			if len(_acfde) == 0 {
				continue
			}
			_dfceb := _eabgf(_cfdff, _acfde)
			_fbd.put(_cfed, _fddcg, _dfceb)
		}
	}
	return &_fbd
}

const (
	_deda markKind = iota
	_bgcf
	_dagaa
	_agbf
)

// TextMark represents extracted text on a page with information regarding both textual content,
// formatting (font and size) and positioning.
// It is the smallest unit of text on a PDF page, typically a single character.
//
// getBBox() in test_text.go shows how to compute bounding boxes of substrings of extracted text.
// The following code extracts the text on PDF page `page` into `text` then finds the bounding box
// `bbox` of substring `term` in `text`.
//
//	ex, _ := New(page)
//	// handle errors
//	pageText, _, _, err := ex.ExtractPageText()
//	// handle errors
//	text := pageText.Text()
//	textMarks := pageText.Marks()
//
//		start := strings.Index(text, term)
//	 end := start + len(term)
//	 spanMarks, err := textMarks.RangeOffset(start, end)
//	 // handle errors
//	 bbox, ok := spanMarks.BBox()
//	 // handle errors
type TextMark struct {
	// Text is the extracted text.
	Text string

	// Original is the text in the PDF. It has not been decoded like `Text`.
	Original string

	// BBox is the bounding box of the text.
	BBox _ab.PdfRectangle

	// Font is the font the text was drawn with.
	Font *_ab.PdfFont

	// FontSize is the font size the text was drawn with.
	FontSize float64

	// Offset is the offset of the start of TextMark.Text in the extracted text. If you do this
	//
	//	text, textMarks := pageText.Text(), pageText.Marks()
	//	marks := textMarks.Elements()
	//
	// then marks[i].Offset is the offset of marks[i].Text in text.
	Offset int

	// Meta is set true for spaces and line breaks that we insert in the extracted text. We insert
	// spaces (line breaks) when we see characters that are over a threshold horizontal (vertical)
	//
	//	distance  apart. See wordJoiner (lineJoiner) in PageText.computeViews().
	Meta bool

	// FillColor is the fill color of the text.
	// The color is nil for spaces and line breaks (i.e. the Meta field is true).
	FillColor _ca.Color

	// StrokeColor is the stroke color of the text.
	// The color is nil for spaces and line breaks (i.e. the Meta field is true).
	StrokeColor _ca.Color

	// Orientation is the text orientation
	Orientation int

	// DirectObject is the underlying PdfObject (Text Object) that represents the visible texts. This is introduced to get
	// a simple access to the TextObject in case editing or replacment of some text is needed. E.g during redaction.
	DirectObject _eb.PdfObject

	// ObjString is a decoded string operand of a text-showing operator. It has the same value as `Text` attribute except
	// when many glyphs are represented with the same Text Object that contains multiple length string operand in which case
	// ObjString spans more than one character string that falls in different TextMark objects.
	ObjString []string
	Tw        float64
	Th        float64
	Tc        float64
	Index     int
	_fddc     bool
	_fabfe    *TextTable
}
type fontEntry struct {
	_daba *_ab.PdfFont
	_dba  int64
}

// Elements returns the TextMarks in `ma`.
func (_dbe *TextMarkArray) Elements() []TextMark { return _dbe._beaf }

func (_faef *textObject) getFontDirect(_eecb string) (*_ab.PdfFont, error) {
	_baf, _edd := _faef.getFontDict(_eecb)
	if _edd != nil {
		return nil, _edd
	}
	_daag, _edd := _ab.NewPdfFontFromPdfObject(_baf)
	if _edd != nil {
		_da.Log.Debug("\u0067\u0065\u0074\u0046\u006f\u006e\u0074\u0044\u0069\u0072\u0065\u0063\u0074\u003a\u0020\u004e\u0065\u0077Pd\u0066F\u006f\u006e\u0074\u0046\u0072\u006f\u006d\u0050\u0064\u0066\u004f\u0062j\u0065\u0063\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u006e\u0061\u006d\u0065\u003d%\u0023\u0071\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _eecb, _edd)
	}
	return _daag, _edd
}

var _egebf = _c.MustCompile("\u005e\u005c\u0073\u002a\u0028\u005c\u0064\u002b\u005c\u002e\u003f|\u005b\u0049\u0069\u0076\u005d\u002b\u0029\u005c\u0073\u002a\\\u0029\u003f\u0024")

func (_gfbde *textWord) absorb(_eebg *textWord) {
	_gfbde.PdfRectangle = _abbb(_gfbde.PdfRectangle, _eebg.PdfRectangle)
	_gfbde._bacfd = append(_gfbde._bacfd, _eebg._bacfd...)
}

// Marks returns the TextMark collection for a page. It represents all the text on the page.
func (_cef PageText) Marks() *TextMarkArray               { return &TextMarkArray{_beaf: _cef._gddg} }
func _ddbg(_fgge _ab.PdfRectangle, _cbba bounded) float64 { return _fgge.Ury - _cbba.bbox().Lly }

// PageText represents the layout of text on a device page.
type PageText struct {
	_cbf  []*textMark
	_aeb  string
	_gddg []TextMark
	_aee  []TextTable
	_cfce _ab.PdfRectangle
	_dab  []pathSection
	_cdbf []pathSection
	_gcaf *_eb.PdfObject
	_fcff _eb.PdfObject
	_ffef *_bb.ContentStreamOperations
	_cbef PageTextOptions
}

func (_eedb rulingList) comp(_fdag, _gcbfbd int) bool {
	_ebad, _gdga := _eedb[_fdag], _eedb[_gcbfbd]
	_fcgce, _baead := _ebad._ggcaf, _gdga._ggcaf
	if _fcgce != _baead {
		return _fcgce > _baead
	}
	if _fcgce == _fbag {
		return false
	}
	_bdfgf := func(_bddc bool) bool {
		if _fcgce == _ceedg {
			return _bddc
		}
		return !_bddc
	}
	_caagb, _ffgg := _ebad._dbbce, _gdga._dbbce
	if _caagb != _ffgg {
		return _bdfgf(_caagb > _ffgg)
	}
	_caagb, _ffgg = _ebad._beadd, _gdga._beadd
	if _caagb != _ffgg {
		return _bdfgf(_caagb < _ffgg)
	}
	return _bdfgf(_ebad._dgbc < _gdga._dgbc)
}
func _fcbf(_gfdc *textLine) float64 { return _gfdc._gfed[0].Llx }
func (_dea *textObject) setWordSpacing(_fdabd float64) {
	if _dea == nil {
		return
	}
	_dea._dbfd._ebbfc = _fdabd
}

func _efefd(_ddbb, _gadfe int) int {
	if _ddbb > _gadfe {
		return _ddbb
	}
	return _gadfe
}

func (_bcdec paraList) eventNeighbours(_dgdf []event) map[*textPara][]int {
	_cf.Slice(_dgdf, func(_edafa, _dfag int) bool {
		_dcfcg, _aaeeg := _dgdf[_edafa], _dgdf[_dfag]
		_beecd, _cbgf := _dcfcg._dcfab, _aaeeg._dcfab
		if _beecd != _cbgf {
			return _beecd < _cbgf
		}
		if _dcfcg._ccbfba != _aaeeg._ccbfba {
			return _dcfcg._ccbfba
		}
		return _edafa < _dfag
	})
	_bbge := make(map[int]intSet)
	_egcge := make(intSet)
	for _, _eafdcg := range _dgdf {
		if _eafdcg._ccbfba {
			_bbge[_eafdcg._dbcaee] = make(intSet)
			for _aeeee := range _egcge {
				if _aeeee != _eafdcg._dbcaee {
					_bbge[_eafdcg._dbcaee].add(_aeeee)
					_bbge[_aeeee].add(_eafdcg._dbcaee)
				}
			}
			_egcge.add(_eafdcg._dbcaee)
		} else {
			_egcge.del(_eafdcg._dbcaee)
		}
	}
	_afbe := map[*textPara][]int{}
	for _bcafb, _deacc := range _bbge {
		_fbcf := _bcdec[_bcafb]
		if len(_deacc) == 0 {
			_afbe[_fbcf] = nil
			continue
		}
		_aaaeca := make([]int, len(_deacc))
		_cgdf := 0
		for _cgggb := range _deacc {
			_aaaeca[_cgdf] = _cgggb
			_cgdf++
		}
		_afbe[_fbcf] = _aaaeca
	}
	return _afbe
}

func _ccbfd(_fbbgb *list) []*textLine {
	for _, _efed := range _fbbgb._ffaf {
		switch _efed._dbga {
		case "\u004c\u0042\u006fd\u0079":
			if len(_efed._dbfeb) != 0 {
				return _efed._dbfeb
			}
			return _ccbfd(_efed)
		case "\u0053\u0070\u0061\u006e":
			return _efed._dbfeb
		case "I\u006e\u006c\u0069\u006e\u0065\u0053\u0068\u0061\u0070\u0065":
			return _efed._dbfeb
		}
	}
	return nil
}

func _dbadf(_dagce []_eb.PdfObject) (_ffde, _dgbeb float64, _cadddg error) {
	if len(_dagce) != 2 {
		return 0, 0, _a.Errorf("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0073\u003a \u0025\u0064", len(_dagce))
	}
	_bgebb, _cadddg := _eb.GetNumbersAsFloat(_dagce)
	if _cadddg != nil {
		return 0, 0, _cadddg
	}
	return _bgebb[0], _bgebb[1], nil
}

func _cdafg(_fggbd, _egaae _cfb.Image) _cfb.Image {
	_addbb, _gadab := _egaae.Bounds().Size(), _fggbd.Bounds().Size()
	_acba, _abbfe := _addbb.X, _addbb.Y
	if _gadab.X > _acba {
		_acba = _gadab.X
	}
	if _gadab.Y > _abbfe {
		_abbfe = _gadab.Y
	}
	_cddga := _cfb.Rect(0, 0, _acba, _abbfe)
	if _addbb.X != _acba || _addbb.Y != _abbfe {
		_bcac := _cfb.NewRGBA(_cddga)
		_b.BiLinear.Scale(_bcac, _cddga, _fggbd, _egaae.Bounds(), _b.Over, nil)
		_egaae = _bcac
	}
	if _gadab.X != _acba || _gadab.Y != _abbfe {
		_gaeg := _cfb.NewRGBA(_cddga)
		_b.BiLinear.Scale(_gaeg, _cddga, _fggbd, _fggbd.Bounds(), _b.Over, nil)
		_fggbd = _gaeg
	}
	_cgfd := _cfb.NewRGBA(_cddga)
	_b.DrawMask(_cgfd, _cddga, _fggbd, _cfb.Point{}, _egaae, _cfb.Point{}, _b.Over)
	return _cgfd
}

func (_cddd rectRuling) checkWidth(_fbeg, _bdcgb float64) (float64, bool) {
	_dfecg := _bdcgb - _fbeg
	_geda := _dfecg <= _efbg
	return _dfecg, _geda
}

func (_abac *textTable) emptyCompositeRow(_fbccd int) bool {
	for _ffdg := 0; _ffdg < _abac._cbcb; _ffdg++ {
		if _ddaf, _cgcg := _abac._fcda[_fadg(_ffdg, _fbccd)]; _cgcg {
			if len(_ddaf.paraList) > 0 {
				return false
			}
		}
	}
	return true
}

func (_affcg *textWord) appendMark(_egdeb *textMark, _abcc _ab.PdfRectangle) {
	_affcg._bacfd = append(_affcg._bacfd, _egdeb)
	_affcg.PdfRectangle = _abbb(_affcg.PdfRectangle, _egdeb.PdfRectangle)
	if _egdeb._acaa > _affcg._fcde {
		_affcg._fcde = _egdeb._acaa
	}
	_affcg._faace = _abcc.Ury - _affcg.PdfRectangle.Lly
}

func (_bddeg rulingList) asTiling() gridTiling {
	if _bcf {
		_da.Log.Info("r\u0075\u006ci\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0076\u0065\u0063s\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u002b\u002b\u002b\u0020\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d", len(_bddeg))
	}
	for _bgbdf, _bdce := range _bddeg[1:] {
		_ecag := _bddeg[_bgbdf]
		if _ecag.alignsPrimary(_bdce) && _ecag.alignsSec(_bdce) {
			_da.Log.Error("a\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0044\u0075\u0070\u006c\u0069\u0063\u0061\u0074\u0065 \u0072\u0075\u006c\u0069\u006e\u0067\u0073\u002e\u000a\u0009v=\u0025\u0073\u000a\t\u0077=\u0025\u0073", _bdce, _ecag)
		}
	}
	_bddeg.sortStrict()
	_bddeg.log("\u0073n\u0061\u0070\u0070\u0065\u0064")
	_fefe, _aabf := _bddeg.vertsHorzs()
	_abec := _fefe.primaries()
	_faead := _aabf.primaries()
	_cgbfc := len(_abec) - 1
	_cfae := len(_faead) - 1
	if _cgbfc == 0 || _cfae == 0 {
		return gridTiling{}
	}
	_badbc := _ab.PdfRectangle{Llx: _abec[0], Urx: _abec[_cgbfc], Lly: _faead[0], Ury: _faead[_cfae]}
	if _bcf {
		_da.Log.Info("\u0072\u0075l\u0069\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0076\u0065\u0072\u0074s=\u0025\u0064", len(_fefe))
		for _bada, _egeeb := range _fefe {
			_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _bada, _egeeb)
		}
		_da.Log.Info("\u0072\u0075l\u0069\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0068\u006f\u0072\u007as=\u0025\u0064", len(_aabf))
		for _aebd, _fgff := range _aabf {
			_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _aebd, _fgff)
		}
		_da.Log.Info("\u0072\u0075\u006c\u0069\u006eg\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067:\u0020\u0020\u0077\u0078\u0068\u003d\u0025\u0064\u0078\u0025\u0064\u000a\u0009\u006c\u006c\u0078\u003d\u0025\u002e\u0032\u0066\u000a\u0009\u006c\u006c\u0079\u003d\u0025\u002e\u0032f", _cgbfc, _cfae, _abec, _faead)
	}
	_adgee := make([]gridTile, _cgbfc*_cfae)
	for _gcfd := _cfae - 1; _gcfd >= 0; _gcfd-- {
		_afcd := _faead[_gcfd]
		_gggd := _faead[_gcfd+1]
		for _fedga := 0; _fedga < _cgbfc; _fedga++ {
			_afabe := _abec[_fedga]
			_cdbfa := _abec[_fedga+1]
			_ggcac := _fefe.findPrimSec(_afabe, _afcd)
			_bcefd := _fefe.findPrimSec(_cdbfa, _afcd)
			_fggde := _aabf.findPrimSec(_afcd, _afabe)
			_cfeed := _aabf.findPrimSec(_gggd, _afabe)
			_dgce := _ab.PdfRectangle{Llx: _afabe, Urx: _cdbfa, Lly: _afcd, Ury: _gggd}
			_feeb := _bfee(_dgce, _ggcac, _bcefd, _fggde, _cfeed)
			_adgee[_gcfd*_cgbfc+_fedga] = _feeb
			if _bcf {
				_a.Printf("\u0020\u0020\u0078\u003d\u0025\u0032\u0064\u0020\u0079\u003d\u0025\u0032\u0064\u003a\u0020%\u0073 \u0025\u0036\u002e\u0032\u0066\u0020\u0078\u0020\u0025\u0036\u002e\u0032\u0066\u000a", _fedga, _gcfd, _feeb.String(), _feeb.Width(), _feeb.Height())
			}
		}
	}
	if _bcf {
		_da.Log.Info("r\u0075\u006c\u0069\u006e\u0067\u004c\u0069\u0073\u0074.\u0061\u0073\u0054\u0069\u006c\u0069\u006eg:\u0020\u0063\u006f\u0061l\u0065\u0073\u0063\u0065\u0020\u0068\u006f\u0072\u0069zo\u006e\u0074a\u006c\u002e\u0020\u0025\u0036\u002e\u0032\u0066", _badbc)
	}
	_bcce := make([]map[float64]gridTile, _cfae)
	for _gfcdg := _cfae - 1; _gfcdg >= 0; _gfcdg-- {
		if _bcf {
			_a.Printf("\u0020\u0020\u0079\u003d\u0025\u0032\u0064\u000a", _gfcdg)
		}
		_bcce[_gfcdg] = make(map[float64]gridTile, _cgbfc)
		for _dgeg := 0; _dgeg < _cgbfc; _dgeg++ {
			_aedf := _adgee[_gfcdg*_cgbfc+_dgeg]
			if _bcf {
				_a.Printf("\u0020\u0020\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _dgeg, _aedf)
			}
			if !_aedf._feaa {
				continue
			}
			_bedc := _dgeg
			for _eabe := _dgeg + 1; !_aedf._gcde && _eabe < _cgbfc; _eabe++ {
				_cbea := _adgee[_gfcdg*_cgbfc+_eabe]
				_aedf.Urx = _cbea.Urx
				_aedf._dbgf = _aedf._dbgf || _cbea._dbgf
				_aedf._bdec = _aedf._bdec || _cbea._bdec
				_aedf._gcde = _cbea._gcde
				if _bcf {
					_a.Printf("\u0020 \u0020%\u0034\u0064\u003a\u0020\u0025s\u0020\u2192 \u0025\u0073\u000a", _eabe, _cbea, _aedf)
				}
				_bedc = _eabe
			}
			if _bcf {
				_a.Printf(" \u0020 \u0025\u0032\u0064\u0020\u002d\u0020\u0025\u0032d\u0020\u2192\u0020\u0025s\n", _dgeg, _bedc, _aedf)
			}
			_dgeg = _bedc
			_bcce[_gfcdg][_aedf.Llx] = _aedf
		}
	}
	_dbbg := make(map[float64]map[float64]gridTile, _cfae)
	_cgbaf := make(map[float64]map[float64]struct{}, _cfae)
	for _dgee := _cfae - 1; _dgee >= 0; _dgee-- {
		_eebdc := _adgee[_dgee*_cgbfc].Lly
		_dbbg[_eebdc] = make(map[float64]gridTile, _cgbfc)
		_cgbaf[_eebdc] = make(map[float64]struct{}, _cgbfc)
	}
	if _bcf {
		_da.Log.Info("\u0072u\u006c\u0069n\u0067\u004c\u0069s\u0074\u002e\u0061\u0073\u0054\u0069\u006ci\u006e\u0067\u003a\u0020\u0063\u006fa\u006c\u0065\u0073\u0063\u0065\u0020\u0076\u0065\u0072\u0074\u0069c\u0061\u006c\u002e\u0020\u0025\u0036\u002e\u0032\u0066", _badbc)
	}
	for _efadg := _cfae - 1; _efadg >= 0; _efadg-- {
		_adfa := _adgee[_efadg*_cgbfc].Lly
		_cdaf := _bcce[_efadg]
		if _bcf {
			_a.Printf("\u0020\u0020\u0079\u003d\u0025\u0032\u0064\u000a", _efadg)
		}
		for _, _aaea := range _dggbd(_cdaf) {
			if _, _cefdf := _cgbaf[_adfa][_aaea]; _cefdf {
				continue
			}
			_dfdd := _cdaf[_aaea]
			if _bcf {
				_a.Printf(" \u0020\u0020\u0020\u0020\u0076\u0030\u003d\u0025\u0073\u000a", _dfdd.String())
			}
			for _faeac := _efadg - 1; _faeac >= 0; _faeac-- {
				if _dfdd._bdec {
					break
				}
				_eadf := _bcce[_faeac]
				_dfecb, _ggea := _eadf[_aaea]
				if !_ggea {
					break
				}
				if _dfecb.Urx != _dfdd.Urx {
					break
				}
				_dfdd._bdec = _dfecb._bdec
				_dfdd.Lly = _dfecb.Lly
				if _bcf {
					_a.Printf("\u0020\u0020\u0020\u0020  \u0020\u0020\u0076\u003d\u0025\u0073\u0020\u0076\u0030\u003d\u0025\u0073\u000a", _dfecb.String(), _dfdd.String())
				}
				_cgbaf[_dfecb.Lly][_dfecb.Llx] = struct{}{}
			}
			if _efadg == 0 {
				_dfdd._bdec = true
			}
			if _dfdd.complete() {
				_dbbg[_adfa][_aaea] = _dfdd
			}
		}
	}
	_cdgg := gridTiling{PdfRectangle: _badbc, _cgeba: _gadb(_dbbg), _gfdac: _gbfc(_dbbg), _afce: _dbbg}
	_cdgg.log("\u0043r\u0065\u0061\u0074\u0065\u0064")
	return _cdgg
}

func (_afabc *textTable) compositeColCorridors() map[int][]float64 {
	_ccab := make(map[int][]float64, _afabc._cbcb)
	if _cabbg {
		_da.Log.Info("\u0063\u006f\u006d\u0070o\u0073\u0069\u0074\u0065\u0043\u006f\u006c\u0043\u006f\u0072r\u0069d\u006f\u0072\u0073\u003a\u0020\u0077\u003d%\u0064\u0020", _afabc._cbcb)
	}
	for _dcgcf := 0; _dcgcf < _afabc._cbcb; _dcgcf++ {
		_ccab[_dcgcf] = nil
	}
	return _ccab
}
func (_cgbg *textPara) bbox() _ab.PdfRectangle { return _cgbg.PdfRectangle }
func _dggbd(_eegab map[float64]gridTile) []float64 {
	_fcdd := make([]float64, 0, len(_eegab))
	for _afff := range _eegab {
		_fcdd = append(_fcdd, _afff)
	}
	_cf.Float64s(_fcdd)
	return _fcdd
}

func _bgeg(_aagg, _gcbb _ab.PdfRectangle) bool {
	return _aagg.Llx <= _gcbb.Llx && _gcbb.Urx <= _aagg.Urx && _aagg.Lly <= _gcbb.Lly && _gcbb.Ury <= _aagg.Ury
}

func _fggf(_gfcc string, _fffef []rulingList) {
	_da.Log.Info("\u0024\u0024 \u0025\u0064\u0020g\u0072\u0069\u0064\u0073\u0020\u002d\u0020\u0025\u0073", len(_fffef), _gfcc)
	for _fbbdd, _eaeeg := range _fffef {
		_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _fbbdd, _eaeeg.String())
	}
}

func _beecb(_aabd []int) []int {
	_fbcd := make([]int, len(_aabd))
	for _bceb, _ccbd := range _aabd {
		_fbcd[len(_aabd)-1-_bceb] = _ccbd
	}
	return _fbcd
}

func _eabgf(_eggd _ab.PdfRectangle, _fdbc []*textLine) *textPara {
	return &textPara{PdfRectangle: _eggd, _fcbfe: _fdbc}
}

type textWord struct {
	_ab.PdfRectangle
	_faace float64
	_faaf  string
	_bacfd []*textMark
	_fcde  float64
	_bace  bool
}

func (_gaef intSet) del(_edgd int) { delete(_gaef, _edgd) }

// PageFonts represents extracted fonts on a PDF page.
type PageFonts struct{ Fonts []Font }

func (_gcbe paraList) reorder(_gbcb []int) {
	_ccbfb := make(paraList, len(_gcbe))
	for _begbe, _feag := range _gbcb {
		_ccbfb[_begbe] = _gcbe[_feag]
	}
	copy(_gcbe, _ccbfb)
}

// String returns a description of `w`.
func (_eaebe *textWord) String() string {
	return _a.Sprintf("\u0025\u002e2\u0066\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\"%\u0073\u0022", _eaebe._faace, _eaebe.PdfRectangle, _eaebe._fcde, _eaebe._faaf)
}

func _eaea(_eda []*textLine, _dgfef map[float64][]*textLine, _baab []float64, _beef int, _ecge, _afad float64) []*list {
	_fafeg := []*list{}
	_bdbd := _beef
	_beef = _beef + 1
	_cdcf := _baab[_bdbd]
	_eceb := _dgfef[_cdcf]
	_cdag := _ccbe(_eceb, _afad, _ecge)
	for _bbbg, _cfbae := range _cdag {
		var _bggb float64
		_ecfg := []*list{}
		_dafde := _cfbae._defc
		_gce := _afad
		if _bbbg < len(_cdag)-1 {
			_gce = _cdag[_bbbg+1]._defc
		}
		if _beef < len(_baab) {
			_ecfg = _eaea(_eda, _dgfef, _baab, _beef, _dafde, _gce)
		}
		_bggb = _gce
		if len(_ecfg) > 0 {
			_gbaf := _ecfg[0]
			if len(_gbaf._dbfeb) > 0 {
				_bggb = _gbaf._dbfeb[0]._defc
			}
		}
		_baca := []*textLine{_cfbae}
		_eagg := _dcfd(_cfbae, _eda, _baab, _dafde, _bggb)
		_baca = append(_baca, _eagg...)
		_bdffd := _defbb(_baca, "\u0062\u0075\u006c\u006c\u0065\u0074", _ecfg)
		_bdffd._gfec = _aeag(_baca, "")
		_fafeg = append(_fafeg, _bdffd)
	}
	return _fafeg
}
func (_fcdf rulingList) sort() { _cf.Slice(_fcdf, _fcdf.comp) }
func (_gccef *ruling) intersects(_dggca *ruling) bool {
	_ecgae := (_gccef._ggcaf == _eccce && _dggca._ggcaf == _ceedg) || (_dggca._ggcaf == _eccce && _gccef._ggcaf == _ceedg)
	_acdga := func(_gbcf, _adbb *ruling) bool {
		return _gbcf._beadd-_gded <= _adbb._dbbce && _adbb._dbbce <= _gbcf._dgbc+_gded
	}
	_ecgg := _acdga(_gccef, _dggca)
	_bgde := _acdga(_dggca, _gccef)
	if _egc {
		_a.Printf("\u0020\u0020\u0020\u0020\u0069\u006e\u0074\u0065\u0072\u0073\u0065\u0063\u0074\u0073\u003a\u0020\u0020\u006fr\u0074\u0068\u006f\u0067\u006f\u006e\u0061l\u003d\u0025\u0074\u0020\u006f\u0031\u003d\u0025\u0074\u0020\u006f2\u003d\u0025\u0074\u0020\u2192\u0020\u0025\u0074\u000a"+"\u0020\u0020\u0020 \u0020\u0020\u0020\u0076\u003d\u0025\u0073\u000a"+" \u0020\u0020\u0020\u0020\u0020\u0077\u003d\u0025\u0073\u000a", _ecgae, _ecgg, _bgde, _ecgae && _ecgg && _bgde, _gccef, _dggca)
	}
	return _ecgae && _ecgg && _bgde
}

type lineRuling struct {
	_agbc rulingKind
	_fdcd markKind
	_ca.Color
	_bbfe, _eaega _fa.Point
}

func (_bbg *imageExtractContext) extractXObjectImage(_fcge *_eb.PdfObjectName, _ddd _bb.GraphicsState, _cgc *_ab.PdfPageResources) error {
	_ggc, _ := _cgc.GetXObjectByName(*_fcge)
	if _ggc == nil {
		return nil
	}
	_beae, _ffg := _bbg._abg[_ggc]
	if !_ffg {
		_cgd, _fabf := _cgc.GetXObjectImageByName(*_fcge)
		if _fabf != nil {
			return _fabf
		}
		if _cgd == nil {
			return nil
		}
		_dgg, _fabf := _cgd.ToImage()
		if _fabf != nil {
			return _fabf
		}
		var _ecd _cfb.Image
		if _cgd.Mask != nil {
			if _ecd, _fabf = _beeb(_cgd.Mask, _ca.Opaque); _fabf != nil {
				_da.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u0063\u006f\u0075\u006c\u0064 \u006eo\u0074\u0020\u0067\u0065\u0074\u0020\u0065\u0078\u0070\u006c\u0069\u0063\u0069\u0074\u0020\u0069\u006d\u0061\u0067e\u0020\u006d\u0061\u0073\u006b\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e")
			}
		} else if _cgd.SMask != nil {
			_ecd, _fabf = _ccfb(_cgd.SMask, _ca.Opaque)
			if _fabf != nil {
				_da.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0066\u0074\u0020\u0069\u006da\u0067e\u0020\u006d\u0061\u0073k\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
			}
		}
		if _ecd != nil {
			_afc, _fba := _dgg.ToGoImage()
			if _fba != nil {
				return _fba
			}
			_afc = _cdafg(_afc, _ecd)
			switch _cgd.ColorSpace.String() {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0049n\u0064\u0065\u0078\u0065\u0064":
				_dgg, _fba = _ab.ImageHandling.NewGrayImageFromGoImage(_afc)
				if _fba != nil {
					return _fba
				}
			default:
				_dgg, _fba = _ab.ImageHandling.NewImageFromGoImage(_afc)
				if _fba != nil {
					return _fba
				}
			}
		}
		_beae = &cachedImage{_daf: _dgg, _fcfg: _cgd.ColorSpace}
		_bbg._abg[_ggc] = _beae
	}
	_geb := _beae._daf
	_dcc := _beae._fcfg
	_fdc, _edbg := _dcc.ImageToRGB(*_geb)
	if _edbg != nil {
		return _edbg
	}
	_da.Log.Debug("@\u0044\u006f\u0020\u0043\u0054\u004d\u003a\u0020\u0025\u0073", _ddd.CTM.String())
	_ebb := ImageMark{Image: &_fdc, Width: _ddd.CTM.ScalingFactorX(), Height: _ddd.CTM.ScalingFactorY(), Angle: _ddd.CTM.Angle()}
	_ebb.X, _ebb.Y = _ddd.CTM.Translation()
	_bbg._daae = append(_bbg._daae, _ebb)
	_bbg._cc++
	return nil
}

var _ae = false

func (_fcffa gridTiling) complete() bool {
	for _, _gafde := range _fcffa._afce {
		for _, _gdgc := range _gafde {
			if !_gdgc.complete() {
				return false
			}
		}
	}
	return true
}

func (_dagb *textObject) moveTextSetLeading(_abba, _bcbg float64) {
	_dagb._dbfd._abda = -_bcbg
	_dagb.moveLP(_abba, _bcbg)
}

// String returns a human readable description of `vecs`.
func (_cdad rulingList) String() string {
	if len(_cdad) == 0 {
		return "\u007b \u0045\u004d\u0050\u0054\u0059\u0020}"
	}
	_gbbb, _dbadef := _cdad.vertsHorzs()
	_dbgd := len(_gbbb)
	_ddgg := len(_dbadef)
	if _dbgd == 0 || _ddgg == 0 {
		return _a.Sprintf("\u007b%\u0064\u0020\u0078\u0020\u0025\u0064}", _dbgd, _ddgg)
	}
	_fbfde := _ab.PdfRectangle{Llx: _gbbb[0]._dbbce, Urx: _gbbb[_dbgd-1]._dbbce, Lly: _dbadef[_ddgg-1]._dbbce, Ury: _dbadef[0]._dbbce}
	return _a.Sprintf("\u007b\u0025d\u0020\u0078\u0020%\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u007d", _dbgd, _ddgg, _fbfde)
}

func (_bcb *imageExtractContext) extractInlineImage(_fab *_bb.ContentStreamInlineImage, _dgf _bb.GraphicsState, _gff *_ab.PdfPageResources) error {
	_fgfe, _daaa := _fab.ToImage(_gff)
	if _daaa != nil {
		return _daaa
	}
	_cbd, _daaa := _fab.GetColorSpace(_gff)
	if _daaa != nil {
		return _daaa
	}
	if _cbd == nil {
		_cbd = _ab.NewPdfColorspaceDeviceGray()
	}
	_ecc, _daaa := _cbd.ImageToRGB(*_fgfe)
	if _daaa != nil {
		return _daaa
	}
	_agg := ImageMark{Image: &_ecc, Width: _dgf.CTM.ScalingFactorX(), Height: _dgf.CTM.ScalingFactorY(), Angle: _dgf.CTM.Angle()}
	_agg.X, _agg.Y = _dgf.CTM.Translation()
	_bcb._daae = append(_bcb._daae, _agg)
	_bcb._caa++
	return nil
}

const (
	_gccbd = 1.0e-6
	_dddg  = 1.0e-4
	_efgd  = 10
	_aegf  = 6
	_cbgd  = 0.5
	_aded  = 0.12
	_egdg  = 0.19
	_aca   = 0.04
	_aggcg = 0.04
	_fagc  = 1.0
	_ebbbc = 0.04
	_ddec  = 0.4
	_ccad  = 0.7
	_ggfb  = 1.0
	_cga   = 0.1
	_fffad = 1.4
	_aedd  = 0.46
	_gccd  = 0.02
	_dfgbg = 0.2
	_agdg  = 0.5
	_fafa  = 4
	_dbab  = 4.0
	_afdc  = 6
	_bfbd  = 0.3
	_daaag = 0.01
	_fdae  = 0.02
	_fafe  = 2
	_ced   = 2
	_edcf  = 500
	_addc  = 4.0
	_bcagg = 4.0
	_cfggb = 0.05
	_addcd = 0.1
	_gded  = 2.0
	_efbg  = 2.0
	_fdbf  = 1.5
	_agcf  = 3.0
	_gfad  = 0.25
)

func (_dbfdd *textWord) addDiacritic(_bcagb string) {
	_acgg := _dbfdd._bacfd[len(_dbfdd._bacfd)-1]
	_acgg._bbeb += _bcagb
	_acgg._bbeb = _cdb.NFKC.String(_acgg._bbeb)
}

type subpath struct {
	_bfaa []_fa.Point
	_gfde bool
}

const (
	_ff   = "\u0045\u0052R\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065"
	_eggg = "\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043a\u006e\u0027\u0074 g\u0065\u0074\u0020\u0066\u006f\u006et\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u002c\u0020\u0066\u006fn\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064"
	_ec   = "\u0045\u0052\u0052O\u0052\u003a\u0020\u0043\u0061\u006e\u0027\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002c\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065"
)

func (_gdb *stateStack) pop() *textState {
	if _gdb.empty() {
		return nil
	}
	_ebga := *(*_gdb)[len(*_gdb)-1]
	*_gdb = (*_gdb)[:len(*_gdb)-1]
	return &_ebga
}
func _faeff(_fafgb, _dedf bounded) float64 { return _fafgb.bbox().Llx - _dedf.bbox().Llx }
func (_bccca paraList) inTile(_gbef gridTile) paraList {
	var _gabcg paraList
	for _, _cgge := range _bccca {
		if _gbef.contains(_cgge.PdfRectangle) {
			_gabcg = append(_gabcg, _cgge)
		}
	}
	if _cabbg {
		_a.Printf("\u0020 \u0020\u0069\u006e\u0054i\u006c\u0065\u003a\u0020\u0020%\u0073 \u0069n\u0073\u0069\u0064\u0065\u003d\u0025\u0064\n", _gbef, len(_gabcg))
		for _bebd, _ecca := range _gabcg {
			_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _bebd, _ecca)
		}
		_a.Println("")
	}
	return _gabcg
}

func (_cdgb rulingList) connections(_dbdbg map[int]intSet, _adaa int) intSet {
	_bcgf := make(intSet)
	_gfagg := make(intSet)
	var _bgagc func(int)
	_bgagc = func(_fdgg int) {
		if !_gfagg.has(_fdgg) {
			_gfagg.add(_fdgg)
			for _bffbef := range _cdgb {
				if _dbdbg[_bffbef].has(_fdgg) {
					_bcgf.add(_bffbef)
				}
			}
			for _caee := range _cdgb {
				if _bcgf.has(_caee) {
					_bgagc(_caee)
				}
			}
		}
	}
	_bgagc(_adaa)
	return _bcgf
}

func (_gdg *textPara) writeCellText(_gacda _f.Writer) {
	for _dadg, _dabb := range _gdg._fcbfe {
		_gcfb := _dabb.text()
		_aeaf := _gacb && _dabb.endsInHyphen() && _dadg != len(_gdg._fcbfe)-1
		if _aeaf {
			_gcfb = _fdgf(_gcfb)
		}
		_gacda.Write([]byte(_gcfb))
		if !(_aeaf || _dadg == len(_gdg._fcbfe)-1) {
			_gacda.Write([]byte(_agadc(_dabb._defc, _gdg._fcbfe[_dadg+1]._defc)))
		}
	}
}

func (_ffcfg paraList) sortReadingOrder() {
	_da.Log.Trace("\u0073\u006fr\u0074\u0052\u0065\u0061\u0064i\u006e\u0067\u004f\u0072\u0064e\u0072\u003a\u0020\u0070\u0061\u0072\u0061\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_ffcfg))
	if len(_ffcfg) <= 1 {
		return
	}
	_ffcfg.computeEBBoxes()
	_cf.Slice(_ffcfg, func(_daaf, _bbaea int) bool { return _cega(_ffcfg[_daaf], _ffcfg[_bbaea]) <= 0 })
}

// NewFromContents creates a new extractor from contents and page resources.
func NewFromContents(contents string, resources *_ab.PdfPageResources) (*Extractor, error) {
	const _def = "\u0065x\u0074\u0072\u0061\u0063t\u006f\u0072\u002e\u004e\u0065w\u0046r\u006fm\u0043\u006f\u006e\u0074\u0065\u006e\u0074s"
	_bg := &Extractor{_ef: contents, _fgf: resources, _eg: map[string]fontEntry{}, _fag: map[string]textResult{}}
	return _bg, nil
}

func _beec(_fcgfe []TextMark, _ddfe *int, _bdac string) []TextMark {
	_cgcc := _cafa
	_cgcc.Text = _bdac
	return _egfcc(_fcgfe, _ddfe, _cgcc)
}

func _fcebe(_aaeb, _fbea _fa.Point) rulingKind {
	_eaed := _g.Abs(_aaeb.X - _fbea.X)
	_cgfeb := _g.Abs(_aaeb.Y - _fbea.Y)
	return _cddg(_eaed, _cgfeb, _cfggb)
}

type intSet map[int]struct{}

func (_ddb *shapesState) moveTo(_bfda, _eeeb float64) {
	_ddb._caef = true
	_ddb._beaa = _ddb.devicePoint(_bfda, _eeeb)
	if _caad {
		_da.Log.Info("\u006d\u006fv\u0065\u0054\u006f\u003a\u0020\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0020\u0064\u0065\u0076\u0069\u0063\u0065\u003d%.\u0032\u0066", _bfda, _eeeb, _ddb._beaa)
	}
}

var _cafa = TextMark{Text: "\u005b\u0058\u005d", Original: "\u0020", Meta: true, FillColor: _ca.White, StrokeColor: _ca.White}

const (
	_cfdb  = false
	_fbbg  = false
	_eagc  = false
	_afea  = false
	_caad  = false
	_bbf   = false
	_fcbg  = false
	_bgada = false
	_dabe  = false
	_bdcf  = _dabe && true
	_dfae  = _bdcf && false
	_abgg  = _dabe && true
	_cabbg = false
	_cdegf = _cabbg && false
	_adeb  = _cabbg && true
	_egc   = false
	_gcbad = _egc && false
	_gecc  = _egc && false
	_bcf   = _egc && true
	_gddab = _egc && false
	_dee   = _egc && false
)

func _cfgad(_gcfbb int, _degg map[int][]float64) ([]int, int) {
	_deec := make([]int, _gcfbb)
	_bgdcf := 0
	for _cbgc := 0; _cbgc < _gcfbb; _cbgc++ {
		_deec[_cbgc] = _bgdcf
		_bgdcf += len(_degg[_cbgc]) + 1
	}
	return _deec, _bgdcf
}

// Extractor stores and offers functionality for extracting content from PDF pages.
type Extractor struct {
	_ef  string
	_fgf *_ab.PdfPageResources
	_efe _ab.PdfRectangle
	_fcf *_ab.PdfRectangle
	_eg  map[string]fontEntry
	_fag map[string]textResult
	_bc  int64
	_ce  int
	_bbe *Options
	_egg *_eb.PdfObject
	_gb  _eb.PdfObject
}

func _eafe(_cbge _ab.PdfColorspace, _cdbag _ab.PdfColor) _ca.Color {
	if _cbge == nil || _cdbag == nil {
		return _ca.Black
	}
	_gbgfb, _ffad := _cbge.ColorToRGB(_cdbag)
	if _ffad != nil {
		_da.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006fu\u006c\u0064\u0020no\u0074\u0020\u0063\u006f\u006e\u0076e\u0072\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0025\u0076\u0020\u0028\u0025\u0076)\u0020\u0074\u006f\u0020\u0052\u0047\u0042\u003a \u0025\u0073", _cdbag, _cbge, _ffad)
		return _ca.Black
	}
	_ggeb, _gedab := _gbgfb.(*_ab.PdfColorDeviceRGB)
	if !_gedab {
		_da.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0065\u0064 \u0063\u006f\u006c\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0052\u0047\u0042\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0076", _gbgfb)
		return _ca.Black
	}
	return _ca.NRGBA{R: uint8(_ggeb.R() * 255), G: uint8(_ggeb.G() * 255), B: uint8(_ggeb.B() * 255), A: uint8(255)}
}

func _gbd(_aeab func(*wordBag, *textWord, float64) bool, _afec float64) func(*wordBag, *textWord) bool {
	return func(_afae *wordBag, _ecbg *textWord) bool { return _aeab(_afae, _ecbg, _afec) }
}

type cachedImage struct {
	_daf  *_ab.Image
	_fcfg _ab.PdfColorspace
}

var _gffc string = "\u0028\u003f\u0069\u0029\u005e\u0028\u004d\u007b\u0030\u002c\u0033\u007d\u0029\u0028\u0043\u0028?\u003a\u0044\u007cM\u0029\u007c\u0044\u003f\u0043{\u0030\u002c\u0033\u007d\u0029\u0028\u0058\u0028\u003f\u003a\u004c\u007c\u0043\u0029\u007cL\u003f\u0058\u007b\u0030\u002c\u0033}\u0029\u0028\u0049\u0028\u003f\u003a\u0056\u007c\u0058\u0029\u007c\u0056\u003f\u0049\u007b\u0030\u002c\u0033\u007d\u0029\u0028\u005c\u0029\u007c\u005c\u002e\u0029\u007c\u005e\u005c\u0028\u0028\u004d\u007b\u0030\u002c\u0033\u007d\u0029\u0028\u0043\u0028\u003f\u003aD\u007cM\u0029\u007c\u0044\u003f\u0043\u007b\u0030\u002c\u0033\u007d\u0029\u0028\u0058\u0028?\u003a\u004c\u007c\u0043\u0029\u007c\u004c?\u0058\u007b0\u002c\u0033\u007d\u0029(\u0049\u0028\u003f\u003a\u0056|\u0058\u0029\u007c\u0056\u003f\u0049\u007b\u0030\u002c\u0033\u007d\u0029\u005c\u0029"

type textResult struct {
	_ccf PageText
	_aec int
	_efg int
}

func _dcbe(_acfbe map[int]intSet) []int {
	_fddge := make([]int, 0, len(_acfbe))
	for _bgaf := range _acfbe {
		_fddge = append(_fddge, _bgaf)
	}
	_cf.Ints(_fddge)
	return _fddge
}

func _gbfc(_bfaad map[float64]map[float64]gridTile) []float64 {
	_baee := make([]float64, 0, len(_bfaad))
	for _abgbea := range _bfaad {
		_baee = append(_baee, _abgbea)
	}
	_cf.Float64s(_baee)
	_bbed := len(_baee)
	for _fedag := 0; _fedag < _bbed/2; _fedag++ {
		_baee[_fedag], _baee[_bbed-1-_fedag] = _baee[_bbed-1-_fedag], _baee[_fedag]
	}
	return _baee
}

func _cddg(_gbgac, _fgfc, _bfbc float64) rulingKind {
	if _gbgac >= _bfbc && _dbdg(_fgfc, _gbgac) {
		return _ceedg
	}
	if _fgfc >= _bfbc && _dbdg(_gbgac, _fgfc) {
		return _eccce
	}
	return _fbag
}

func (_fbeaa rulingList) primMinMax() (float64, float64) {
	_gceg, _dbgad := _fbeaa[0]._dbbce, _fbeaa[0]._dbbce
	for _, _ccefe := range _fbeaa[1:] {
		if _ccefe._dbbce < _gceg {
			_gceg = _ccefe._dbbce
		} else if _ccefe._dbbce > _dbgad {
			_dbgad = _ccefe._dbbce
		}
	}
	return _gceg, _dbgad
}

func (_abbf *shapesState) addPoint(_efda, _fdef float64) {
	_ggdaf := _abbf.establishSubpath()
	_adec := _abbf.devicePoint(_efda, _fdef)
	if _ggdaf == nil {
		_abbf._caef = true
		_abbf._beaa = _adec
	} else {
		_ggdaf.add(_adec)
	}
}

func _affa(_gfddd *wordBag, _cbfbg float64, _ccfa, _deea rulingList) []*wordBag {
	var _fggbb []*wordBag
	for _, _dcfbb := range _gfddd.depthIndexes() {
		_badb := false
		for !_gfddd.empty(_dcfbb) {
			_febd := _gfddd.firstReadingIndex(_dcfbb)
			_ecda := _gfddd.firstWord(_febd)
			_adaed := _daaga(_ecda, _cbfbg, _ccfa, _deea)
			_gfddd.removeWord(_ecda, _febd)
			if _fcbg {
				_da.Log.Info("\u0066\u0069\u0072\u0073\u0074\u0057\u006f\u0072\u0064\u0020\u005e\u005e^\u005e\u0020\u0025\u0073", _ecda.String())
			}
			for _bdde := true; _bdde; _bdde = _badb {
				_badb = false
				_cfad := _ggfb * _adaed._caag
				_fddae := _ddec * _adaed._caag
				_cffb := _fagc * _adaed._caag
				if _fcbg {
					_da.Log.Info("\u0070a\u0072a\u0057\u006f\u0072\u0064\u0073\u0020\u0064\u0065\u0070\u0074\u0068 \u0025\u002e\u0032\u0066 \u002d\u0020\u0025\u002e\u0032f\u0020\u006d\u0061\u0078\u0049\u006e\u0074\u0072\u0061\u0044\u0065\u0070\u0074\u0068\u0047\u0061\u0070\u003d\u0025\u002e\u0032\u0066\u0020\u006d\u0061\u0078\u0049\u006e\u0074\u0072\u0061R\u0065\u0061\u0064\u0069\u006e\u0067\u0047\u0061p\u003d\u0025\u002e\u0032\u0066", _adaed.minDepth(), _adaed.maxDepth(), _cffb, _fddae)
				}
				if _gfddd.scanBand("\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c", _adaed, _gbd(_ccgf, 0), _adaed.minDepth()-_cffb, _adaed.maxDepth()+_cffb, _ebbbc, false, false) > 0 {
					_badb = true
				}
				if _gfddd.scanBand("\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c", _adaed, _gbd(_ccgf, _fddae), _adaed.minDepth(), _adaed.maxDepth(), _ccad, false, false) > 0 {
					_badb = true
				}
				if _badb {
					continue
				}
				_edcgc := _gfddd.scanBand("", _adaed, _gbd(_bfcd, _cfad), _adaed.minDepth(), _adaed.maxDepth(), _cga, true, false)
				if _edcgc > 0 {
					_ddcc := (_adaed.maxDepth() - _adaed.minDepth()) / _adaed._caag
					if (_edcgc > 1 && float64(_edcgc) > 0.3*_ddcc) || _edcgc <= 10 {
						if _gfddd.scanBand("\u006f\u0074\u0068e\u0072", _adaed, _gbd(_bfcd, _cfad), _adaed.minDepth(), _adaed.maxDepth(), _cga, false, true) > 0 {
							_badb = true
						}
					}
				}
			}
			_fggbb = append(_fggbb, _adaed)
		}
	}
	return _fggbb
}

func (_gcff *wordBag) depthIndexes() []int {
	if len(_gcff._edgf) == 0 {
		return nil
	}
	_ebbb := make([]int, len(_gcff._edgf))
	_abfg := 0
	for _geab := range _gcff._edgf {
		_ebbb[_abfg] = _geab
		_abfg++
	}
	_cf.Ints(_ebbb)
	return _ebbb
}

// RangeOffset returns the TextMarks in `ma` that overlap text[start:end] in the extracted text.
// These are tm: `start` <= tm.Offset + len(tm.Text) && tm.Offset < `end` where
// `start` and `end` are offsets in the extracted text.
// NOTE: TextMarks can contain multiple characters. e.g. "ffi" for the ﬃ ligature so the first and
// last elements of the returned TextMarkArray may only partially overlap text[start:end].
func (_ddcd *TextMarkArray) RangeOffset(start, end int) (*TextMarkArray, error) {
	if _ddcd == nil {
		return nil, _e.New("\u006da\u003d\u003d\u006e\u0069\u006c")
	}
	if end < start {
		return nil, _a.Errorf("\u0065\u006e\u0064\u0020\u003c\u0020\u0073\u0074\u0061\u0072\u0074\u002e\u0020\u0052\u0061n\u0067\u0065\u004f\u0066\u0066\u0073\u0065\u0074\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064\u002e\u0020\u0073\u0074\u0061\u0072t=\u0025\u0064\u0020\u0065\u006e\u0064\u003d\u0025\u0064\u0020", start, end)
	}
	_bcef := len(_ddcd._beaf)
	if _bcef == 0 {
		return _ddcd, nil
	}
	if start < _ddcd._beaf[0].Offset {
		start = _ddcd._beaf[0].Offset
	}
	if end > _ddcd._beaf[_bcef-1].Offset+1 {
		end = _ddcd._beaf[_bcef-1].Offset + 1
	}
	_gbac := _cf.Search(_bcef, func(_gacd int) bool { return _ddcd._beaf[_gacd].Offset+len(_ddcd._beaf[_gacd].Text)-1 >= start })
	if !(0 <= _gbac && _gbac < _bcef) {
		_fgbe := _a.Errorf("\u004f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u002e\u0020\u0073\u0074\u0061\u0072\u0074\u003d%\u0064\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020\u006c\u0065\u006e\u003d\u0025\u0064\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025\u0076\u000a\u0009 \u006c\u0061\u0073\u0074\u003d%\u0076", start, _gbac, _bcef, _ddcd._beaf[0], _ddcd._beaf[_bcef-1])
		return nil, _fgbe
	}
	_aggd := _cf.Search(_bcef, func(_bebfb int) bool { return _ddcd._beaf[_bebfb].Offset > end-1 })
	if !(0 <= _aggd && _aggd < _bcef) {
		_bbag := _a.Errorf("\u004f\u0075\u0074\u0020\u006f\u0066\u0020r\u0061\u006e\u0067e\u002e\u0020\u0065n\u0064\u003d%\u0064\u0020\u0069\u0045\u006e\u0064=\u0025d \u006c\u0065\u006e\u003d\u0025\u0064\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025\u0076\u000a\u0009\u0020\u006c\u0061\u0073\u0074\u003d\u0025\u0076", end, _aggd, _bcef, _ddcd._beaf[0], _ddcd._beaf[_bcef-1])
		return nil, _bbag
	}
	if _aggd <= _gbac {
		return nil, _a.Errorf("\u0069\u0045\u006e\u0064\u0020\u003c=\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003a\u0020\u0073\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020\u0065\u006ed\u003d\u0025\u0064\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020i\u0045n\u0064\u003d\u0025\u0064", start, end, _gbac, _aggd)
	}
	return &TextMarkArray{_beaf: _ddcd._beaf[_gbac:_aggd]}, nil
}

func (_fcedf *textTable) subdivide() *textTable {
	_fcedf.logComposite("\u0073u\u0062\u0064\u0069\u0076\u0069\u0064e")
	_ceae := _fcedf.compositeRowCorridors()
	_dbgaa := _fcedf.compositeColCorridors()
	if _cabbg {
		_da.Log.Info("\u0073u\u0062\u0064i\u0076\u0069\u0064\u0065:\u000a\u0009\u0072o\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s=\u0025\u0073\u000a\t\u0063\u006fl\u0043\u006f\u0072\u0072\u0069\u0064o\u0072\u0073=\u0025\u0073", _fdcg(_ceae), _fdcg(_dbgaa))
	}
	if len(_ceae) == 0 || len(_dbgaa) == 0 {
		return _fcedf
	}
	_becf(_ceae)
	_becf(_dbgaa)
	if _cabbg {
		_da.Log.Info("\u0073\u0075\u0062\u0064\u0069\u0076\u0069\u0064\u0065\u0020\u0066\u0069\u0078\u0065\u0064\u003a\u000a\u0009r\u006f\u0077\u0043\u006f\u0072\u0072\u0069d\u006f\u0072\u0073\u003d\u0025\u0073\u000a\u0009\u0063\u006f\u006cC\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u003d\u0025\u0073", _fdcg(_ceae), _fdcg(_dbgaa))
	}
	_ggdee, _gfce := _cfgad(_fcedf._faed, _ceae)
	_gafdc, _fdgbe := _cfgad(_fcedf._cbcb, _dbgaa)
	_ecfge := make(map[uint64]*textPara, _fdgbe*_gfce)
	_bfdeb := &textTable{PdfRectangle: _fcedf.PdfRectangle, _cafe: _fcedf._cafe, _faed: _gfce, _cbcb: _fdgbe, _ebgbd: _ecfge}
	if _cabbg {
		_da.Log.Info("\u0073\u0075b\u0064\u0069\u0076\u0069\u0064\u0065\u003a\u0020\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u0020\u0063\u0065\u006c\u006c\u0073\u003d\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u000a"+"\u0009\u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s\u003d\u0025\u0073\u000a"+"\u0009\u0063\u006f\u006c\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s\u003d\u0025\u0073\u000a"+"\u0009\u0079\u004f\u0066\u0066\u0073\u0065\u0074\u0073=\u0025\u002b\u0076\u000a"+"\u0009\u0078\u004f\u0066\u0066\u0073\u0065\u0074\u0073\u003d\u0025\u002b\u0076", _fcedf._cbcb, _fcedf._faed, _fdgbe, _gfce, _fdcg(_ceae), _fdcg(_dbgaa), _ggdee, _gafdc)
	}
	for _acfff := 0; _acfff < _fcedf._faed; _acfff++ {
		_cbcfb := _ggdee[_acfff]
		for _bcebg := 0; _bcebg < _fcedf._cbcb; _bcebg++ {
			_ffbbb := _gafdc[_bcebg]
			if _cabbg {
				_a.Printf("\u0025\u0036\u0064\u002c %\u0032\u0064\u003a\u0020\u0078\u0030\u003d\u0025\u0064\u0020\u0079\u0030\u003d\u0025d\u000a", _bcebg, _acfff, _ffbbb, _cbcfb)
			}
			_aeagc, _edce := _fcedf._fcda[_fadg(_bcebg, _acfff)]
			if !_edce {
				continue
			}
			_eaag := _aeagc.split(_ceae[_acfff], _dbgaa[_bcebg])
			for _dabad := 0; _dabad < _eaag._faed; _dabad++ {
				for _bagb := 0; _bagb < _eaag._cbcb; _bagb++ {
					_gbcda := _eaag.get(_bagb, _dabad)
					_bfdeb.put(_ffbbb+_bagb, _cbcfb+_dabad, _gbcda)
					if _cabbg {
						_a.Printf("\u0025\u0038\u0064\u002c\u0020\u0025\u0032\u0064\u003a\u0020\u0025\u0073\u000a", _ffbbb+_bagb, _cbcfb+_dabad, _gbcda)
					}
				}
			}
		}
	}
	return _bfdeb
}

func _ccgf(_cbcd *wordBag, _aggf *textWord, _dcg float64) bool {
	return _aggf.Llx < _cbcd.Urx+_dcg && _cbcd.Llx-_dcg < _aggf.Urx
}
func _gega(_adcg, _eabcf _ab.PdfRectangle) bool { return _eedd(_adcg, _eabcf) && _gdac(_adcg, _eabcf) }

var _cadd *_c.Regexp = _c.MustCompile(_gffc + "\u007c" + _cfda)

type gridTile struct {
	_ab.PdfRectangle
	_dbgf, _feaa, _bdec, _gcde bool
}

func (_faac paraList) yNeighbours(_edbc float64) map[*textPara][]int {
	_ffbe := make([]event, 2*len(_faac))
	if _edbc == 0 {
		for _dgfba, _bdfec := range _faac {
			_ffbe[2*_dgfba] = event{_bdfec.Lly, true, _dgfba}
			_ffbe[2*_dgfba+1] = event{_bdfec.Ury, false, _dgfba}
		}
	} else {
		for _adfb, _adedc := range _faac {
			_ffbe[2*_adfb] = event{_adedc.Lly - _edbc*_adedc.fontsize(), true, _adfb}
			_ffbe[2*_adfb+1] = event{_adedc.Ury + _edbc*_adedc.fontsize(), false, _adfb}
		}
	}
	return _faac.eventNeighbours(_ffbe)
}

// Append appends `mark` to the mark array.
func (_adgb *TextMarkArray) Append(mark TextMark) { _adgb._beaf = append(_adgb._beaf, mark) }
func _gaae(_defb _fa.Point) _fa.Matrix            { return _fa.TranslationMatrix(_defb.X, _defb.Y) }
func (_gbcfb *textTable) isExportable() bool {
	if _gbcfb._cafe {
		return true
	}
	_fgcc := func(_caca int) bool {
		_bbdd := _gbcfb.get(0, _caca)
		if _bbdd == nil {
			return false
		}
		_dfaa := _bbdd.text()
		_ceaa := _fe.RuneCountInString(_dfaa)
		_dfgc := _egebf.MatchString(_dfaa)
		return _ceaa <= 1 || _dfgc
	}
	for _efeag := 0; _efeag < _gbcfb._faed; _efeag++ {
		if !_fgcc(_efeag) {
			return true
		}
	}
	return false
}

func _acad(_bgc []structElement, _gfbd map[int][]*textLine, _cdbcf _eb.PdfObject) []*list {
	_dgbfb := []*list{}
	for _, _gade := range _bgc {
		_dgafc := _gade._baac
		_gefc := int(_gade._bbgg)
		_gefec := _gade._acfd
		_afbg := []*textLine{}
		_bgca := []*list{}
		_gfacc := _gade._ffcf
		_cbfec, _bffbe := (_gfacc.(*_eb.PdfObjectReference))
		if !_bffbe {
			_da.Log.Debug("\u0066\u0061\u0069l\u0065\u0064\u0020\u006f\u0074\u0020\u0063\u0061\u0073\u0074\u0020\u0074\u006f\u0020\u002a\u0063\u006f\u0072\u0065\u002e\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		}
		if _gefc != -1 && _cbfec != nil {
			if _dgge, _gccbe := _gfbd[_gefc]; _gccbe {
				if _efdgb, _bedd := _cdbcf.(*_eb.PdfIndirectObject); _bedd {
					_ggcc := _efdgb.PdfObjectReference
					if _cd.DeepEqual(*_cbfec, _ggcc) {
						_afbg = _dgge
					}
				}
			}
		}
		if _dgafc != nil {
			_bgca = _acad(_dgafc, _gfbd, _cdbcf)
		}
		_dbdec := _defbb(_afbg, _gefec, _bgca)
		_dgbfb = append(_dgbfb, _dbdec)
	}
	return _dgbfb
}

func _acadc(_ddfeb []TextMark, _dcbfb *int) []TextMark {
	_cdgec := _ddfeb[len(_ddfeb)-1]
	_facab := []rune(_cdgec.Text)
	if len(_facab) == 1 {
		_ddfeb = _ddfeb[:len(_ddfeb)-1]
		_cbcf := _ddfeb[len(_ddfeb)-1]
		*_dcbfb = _cbcf.Offset + len(_cbcf.Text)
	} else {
		_gcfbd := _fdgf(_cdgec.Text)
		*_dcbfb += len(_gcfbd) - len(_cdgec.Text)
		_cdgec.Text = _gcfbd
	}
	return _ddfeb
}

var _agfc = map[markKind]string{_bgcf: "\u0073\u0074\u0072\u006f\u006b\u0065", _dagaa: "\u0066\u0069\u006c\u006c", _agbf: "\u0061u\u0067\u006d\u0065\u006e\u0074"}

func _ebede(_gcceg map[float64][]*textLine) []float64 {
	_fcbb := []float64{}
	for _baga := range _gcceg {
		_fcbb = append(_fcbb, _baga)
	}
	_cf.Float64s(_fcbb)
	return _fcbb
}

func (_bcec *textObject) getFillColor() _ca.Color {
	return _eafe(_bcec._fcfe.ColorspaceNonStroking, _bcec._fcfe.ColorNonStroking)
}

func _aeef(_bccc float64) int {
	var _gcag int
	if _bccc >= 0 {
		_gcag = int(_bccc / _aegf)
	} else {
		_gcag = int(_bccc/_aegf) - 1
	}
	return _gcag
}

func (_fddg *shapesState) lineTo(_dbgg, _cggb float64) {
	if _caad {
		_da.Log.Info("\u006c\u0069\u006eeT\u006f\u0028\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0020\u0070\u003d\u0025\u002e\u0032\u0066", _dbgg, _cggb, _fddg.devicePoint(_dbgg, _cggb))
	}
	_fddg.addPoint(_dbgg, _cggb)
}

func (_aeeeb *ruling) equals(_ccgc *ruling) bool {
	return _aeeeb._ggcaf == _ccgc._ggcaf && _ccfac(_aeeeb._dbbce, _ccgc._dbbce) && _ccfac(_aeeeb._beadd, _ccgc._beadd) && _ccfac(_aeeeb._dgbc, _ccgc._dgbc)
}

func (_bdcgf *textTable) log(_fcba string) {
	if !_cabbg {
		return
	}
	_da.Log.Info("~\u007e\u007e\u0020\u0025\u0073\u003a \u0025\u0064\u0020\u0078\u0020\u0025d\u0020\u0067\u0072\u0069\u0064\u003d\u0025t\u000a\u0020\u0020\u0020\u0020\u0020\u0020\u0025\u0036\u002e2\u0066", _fcba, _bdcgf._cbcb, _bdcgf._faed, _bdcgf._cafe, _bdcgf.PdfRectangle)
	for _bagca := 0; _bagca < _bdcgf._faed; _bagca++ {
		for _ffaac := 0; _ffaac < _bdcgf._cbcb; _ffaac++ {
			_afbga := _bdcgf.get(_ffaac, _bagca)
			if _afbga == nil {
				continue
			}
			_a.Printf("%\u0034\u0064\u0020\u00252d\u003a \u0025\u0036\u002e\u0032\u0066 \u0025\u0071\u0020\u0025\u0064\u000a", _ffaac, _bagca, _afbga.PdfRectangle, _gggg(_afbga.text(), 50), _fe.RuneCountInString(_afbga.text()))
		}
	}
}

func (_cbdg *shapesState) lastpointEstablished() (_fa.Point, bool) {
	if _cbdg._caef {
		return _cbdg._beaa, false
	}
	_gcfa := len(_cbdg._dagf)
	if _gcfa > 0 && _cbdg._dagf[_gcfa-1]._gfde {
		return _cbdg._dagf[_gcfa-1].last(), false
	}
	return _fa.Point{}, true
}

func (_begbb rulingList) isActualGrid() (rulingList, bool) {
	_caab, _befd := _begbb.augmentGrid()
	if !(len(_caab) >= _fafe+1 && len(_befd) >= _ced+1) {
		if _egc {
			_da.Log.Info("\u0069s\u0041\u0063t\u0075\u0061\u006c\u0047r\u0069\u0064\u003a \u004e\u006f\u0074\u0020\u0061\u006c\u0069\u0067\u006eed\u002e\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u003c\u0020\u0025d\u0020\u0078 \u0025\u0064", len(_caab), len(_befd), _fafe+1, _ced+1)
		}
		return nil, false
	}
	if _egc {
		_da.Log.Info("\u0069\u0073\u0041\u0063\u0074\u0075a\u006c\u0047\u0072\u0069\u0064\u003a\u0020\u0025\u0073\u0020\u003a\u0020\u0025t\u0020\u0026\u0020\u0025\u0074\u0020\u2192 \u0025\u0074", _begbb, len(_caab) >= 2, len(_befd) >= 2, len(_caab) >= 2 && len(_befd) >= 2)
		for _bdae, _abgbe := range _begbb {
			_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0076\u000a", _bdae, _abgbe)
		}
	}
	if _ccfd {
		_febf, _acegd := _caab[0], _caab[len(_caab)-1]
		_dagad, _abdg := _befd[0], _befd[len(_befd)-1]
		if !(_edcfb(_febf._dbbce-_dagad._beadd) && _edcfb(_acegd._dbbce-_dagad._dgbc) && _edcfb(_dagad._dbbce-_febf._dgbc) && _edcfb(_abdg._dbbce-_febf._beadd)) {
			if _egc {
				_da.Log.Info("\u0069\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069d\u003a\u0020\u0020N\u006f\u0074 \u0061\u006c\u0069\u0067\u006e\u0065d\u002e\n\t\u0076\u0030\u003d\u0025\u0073\u000a\u0009\u0076\u0031\u003d\u0025\u0073\u000a\u0009\u0068\u0030\u003d\u0025\u0073\u000a\u0009\u0068\u0031\u003d\u0025\u0073", _febf, _acegd, _dagad, _abdg)
			}
			return nil, false
		}
	} else {
		if !_caab.aligned() {
			if _gecc {
				_da.Log.Info("i\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069\u0064\u003a\u0020N\u006f\u0074\u0020\u0061\u006c\u0069\u0067n\u0065\u0064\u0020\u0076\u0065\u0072\u0074\u0073\u002e\u0020%\u0064", len(_caab))
			}
			return nil, false
		}
		if !_befd.aligned() {
			if _egc {
				_da.Log.Info("i\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069\u0064\u003a\u0020N\u006f\u0074\u0020\u0061\u006c\u0069\u0067n\u0065\u0064\u0020\u0068\u006f\u0072\u007a\u0073\u002e\u0020%\u0064", len(_befd))
			}
			return nil, false
		}
	}
	_ggde := append(_caab, _befd...)
	return _ggde, true
}
func _ccfac(_bcbf, _dccba float64) bool { return _g.Abs(_bcbf-_dccba) <= _gded }
func (_ega *imageExtractContext) extractContentStreamImages(_bdc string, _ea *_ab.PdfPageResources) error {
	_afa := _bb.NewContentStreamParser(_bdc)
	_bee, _gae := _afa.Parse()
	if _gae != nil {
		return _gae
	}
	if _ega._abg == nil {
		_ega._abg = map[*_eb.PdfObjectStream]*cachedImage{}
	}
	if _ega._bd == nil {
		_ega._bd = &ImageExtractOptions{}
	}
	_agc := _bb.NewContentStreamProcessor(*_bee)
	_agc.AddHandler(_bb.HandlerConditionEnumAllOperands, "", _ega.processOperand)
	return _agc.Process(_ea)
}

// String returns a string describing the current state of the textState stack.
func (_dbb *stateStack) String() string {
	_ccda := []string{_a.Sprintf("\u002d\u002d\u002d\u002d f\u006f\u006e\u0074\u0020\u0073\u0074\u0061\u0063\u006b\u003a\u0020\u0025\u0064", len(*_dbb))}
	for _ddgb, _gcfc := range *_dbb {
		_abd := "\u003c\u006e\u0069l\u003e"
		if _gcfc != nil {
			_abd = _gcfc.String()
		}
		_ccda = append(_ccda, _a.Sprintf("\u0009\u0025\u0032\u0064\u003a\u0020\u0025\u0073", _ddgb, _abd))
	}
	return _cb.Join(_ccda, "\u000a")
}

func _ccbe(_afgd []*textLine, _fcgf, _bcfe float64) []*textLine {
	var _eafd []*textLine
	for _, _adge := range _afgd {
		if _fcgf == -1 {
			if _adge._defc > _bcfe {
				_eafd = append(_eafd, _adge)
			}
		} else {
			if _adge._defc > _bcfe && _adge._defc < _fcgf {
				_eafd = append(_eafd, _adge)
			}
		}
	}
	return _eafd
}

func (_df *imageExtractContext) processOperand(_gf *_bb.ContentStreamOperation, _aag _bb.GraphicsState, _ad *_ab.PdfPageResources) error {
	if _gf.Operand == "\u0042\u0049" && len(_gf.Params) == 1 {
		_ebf, _dg := _gf.Params[0].(*_bb.ContentStreamInlineImage)
		if !_dg {
			return nil
		}
		if _eac, _add := _eb.GetBoolVal(_ebf.ImageMask); _add {
			if _eac && !_df._bd.IncludeInlineStencilMasks {
				return nil
			}
		}
		return _df.extractInlineImage(_ebf, _aag, _ad)
	} else if _gf.Operand == "\u0044\u006f" && len(_gf.Params) == 1 {
		_fggb, _gg := _eb.GetName(_gf.Params[0])
		if !_gg {
			_da.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0079\u0070\u0065")
			return _dag
		}
		_, _bf := _ad.GetXObjectByName(*_fggb)
		switch _bf {
		case _ab.XObjectTypeImage:
			return _df.extractXObjectImage(_fggb, _aag, _ad)
		case _ab.XObjectTypeForm:
			return _df.extractFormImages(_fggb, _aag, _ad)
		}
	} else if _df._bce && (_gf.Operand == "\u0073\u0063\u006e" || _gf.Operand == "\u0053\u0043\u004e") && len(_gf.Params) == 1 {
		_bfc, _gbc := _eb.GetName(_gf.Params[0])
		if !_gbc {
			_da.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0079\u0070\u0065")
			return _dag
		}
		_gcb, _gbc := _ad.GetPatternByName(*_bfc)
		if !_gbc {
			_da.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0074\u0074\u0065\u0072n\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
			return nil
		}
		if _gcb.IsTiling() {
			_ecae := _gcb.GetAsTilingPattern()
			_fcc, _fd := _ecae.GetContentStream()
			if _fd != nil {
				return _fd
			}
			_fd = _df.extractContentStreamImages(string(_fcc), _ecae.Resources)
			if _fd != nil {
				return _fd
			}
		}
	} else if (_gf.Operand == "\u0063\u0073" || _gf.Operand == "\u0043\u0053") && len(_gf.Params) >= 1 {
		_df._bce = _gf.Params[0].String() == "\u0050a\u0074\u0074\u0065\u0072\u006e"
	}
	return nil
}

func _ffea(_acfb []*textLine) []*textLine {
	_dafa := []*textLine{}
	for _, _bcbd := range _acfb {
		_bcbac := _bcbd.text()
		_bgdc := _cadd.Find([]byte(_bcbac))
		if _bgdc != nil {
			_dafa = append(_dafa, _bcbd)
		}
	}
	return _dafa
}

func (_cbeb *textLine) endsInHyphen() bool {
	_gfea := _cbeb._gfed[len(_cbeb._gfed)-1]
	_cea := _gfea._faaf
	_dcbg, _debbc := _fe.DecodeLastRuneInString(_cea)
	if _debbc <= 0 || !_fg.Is(_fg.Hyphen, _dcbg) {
		return false
	}
	if _gfea._bace && _faae(_cea) {
		return true
	}
	return _faae(_cbeb.text())
}
func (_ecgc *wordBag) maxDepth() float64 { return _ecgc._bfgd - _ecgc.Lly }
func (_fdd *textObject) reset() {
	_fdd._dbd = _fa.IdentityMatrix()
	_fdd._cdeg = _fa.IdentityMatrix()
	_fdd._cecd = nil
}

func (_afab *textObject) getFontDict(_fffa string) (_fdfde _eb.PdfObject, _fbaaa error) {
	_beedc := _afab._dafd
	if _beedc == nil {
		_da.Log.Debug("g\u0065\u0074\u0046\u006f\u006e\u0074D\u0069\u0063\u0074\u002e\u0020\u004eo\u0020\u0072\u0065\u0073\u006f\u0075\u0072c\u0065\u0073\u002e\u0020\u006e\u0061\u006d\u0065\u003d\u0025#\u0071", _fffa)
		return nil, nil
	}
	_fdfde, _dfc := _beedc.GetFontByName(_eb.PdfObjectName(_fffa))
	if !_dfc {
		_da.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0067\u0065t\u0046\u006f\u006et\u0044\u0069\u0063\u0074\u003a\u0020\u0046\u006f\u006et \u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020\u006e\u0061m\u0065\u003d%\u0023\u0071", _fffa)
		return nil, _e.New("f\u006f\u006e\u0074\u0020no\u0074 \u0069\u006e\u0020\u0072\u0065s\u006f\u0075\u0072\u0063\u0065\u0073")
	}
	return _fdfde, nil
}

func _affc(_egee *list, _bdff *_cb.Builder, _ggbe *string) {
	_cfef := _abfe(_egee, _ggbe)
	_bdff.WriteString(_cfef)
	for _, _cfee := range _egee._ffaf {
		_agad := *_ggbe + "\u0020\u0020\u0020"
		_affc(_cfee, _bdff, &_agad)
	}
}

func _gdfge(_facbe _ab.PdfRectangle) *ruling {
	return &ruling{_ggcaf: _eccce, _dbbce: _facbe.Urx, _beadd: _facbe.Lly, _dgbc: _facbe.Ury}
}

var _gedg = map[rune]string{0x0060: "\u0300", 0x02CB: "\u0300", 0x0027: "\u0301", 0x00B4: "\u0301", 0x02B9: "\u0301", 0x02CA: "\u0301", 0x005E: "\u0302", 0x02C6: "\u0302", 0x007E: "\u0303", 0x02DC: "\u0303", 0x00AF: "\u0304", 0x02C9: "\u0304", 0x02D8: "\u0306", 0x02D9: "\u0307", 0x00A8: "\u0308", 0x00B0: "\u030a", 0x02DA: "\u030a", 0x02BA: "\u030b", 0x02DD: "\u030b", 0x02C7: "\u030c", 0x02C8: "\u030d", 0x0022: "\u030e", 0x02BB: "\u0312", 0x02BC: "\u0313", 0x0486: "\u0313", 0x055A: "\u0313", 0x02BD: "\u0314", 0x0485: "\u0314", 0x0559: "\u0314", 0x02D4: "\u031d", 0x02D5: "\u031e", 0x02D6: "\u031f", 0x02D7: "\u0320", 0x02B2: "\u0321", 0x00B8: "\u0327", 0x02CC: "\u0329", 0x02B7: "\u032b", 0x02CD: "\u0331", 0x005F: "\u0332", 0x204E: "\u0359"}

func _aeag(_badd []*textLine, _cadg string) string {
	var _cbdge _cb.Builder
	_bbecg := 0.0
	for _efagf, _aefa := range _badd {
		_egde := _aefa.text()
		_eea := _aefa._defc
		if _efagf < len(_badd)-1 {
			_bbecg = _badd[_efagf+1]._defc
		} else {
			_bbecg = 0.0
		}
		_cbdge.WriteString(_cadg)
		_cbdge.WriteString(_egde)
		if _bbecg != _eea {
			_cbdge.WriteString("\u000a")
		} else {
			_cbdge.WriteString("\u0020")
		}
	}
	return _cbdge.String()
}

type paraList []*textPara

func (_dbag *wordBag) removeDuplicates() {
	if _abgg {
		_da.Log.Info("r\u0065m\u006f\u0076\u0065\u0044\u0075\u0070\u006c\u0069c\u0061\u0074\u0065\u0073: \u0025\u0071", _dbag.text())
	}
	for _, _fdegd := range _dbag.depthIndexes() {
		if len(_dbag._edgf[_fdegd]) == 0 {
			continue
		}
		_cecdc := _dbag._edgf[_fdegd][0]
		_bcfd := _dfgbg * _cecdc._fcde
		_gabc := _cecdc._faace
		for _, _facf := range _dbag.depthBand(_gabc, _gabc+_bcfd) {
			_ggab := map[*textWord]struct{}{}
			_acda := _dbag._edgf[_facf]
			for _, _edbgf := range _acda {
				if _, _bfafa := _ggab[_edbgf]; _bfafa {
					continue
				}
				for _, _gfgd := range _acda {
					if _, _abbg := _ggab[_gfgd]; _abbg {
						continue
					}
					if _gfgd != _edbgf && _gfgd._faaf == _edbgf._faaf && _g.Abs(_gfgd.Llx-_edbgf.Llx) < _bcfd && _g.Abs(_gfgd.Urx-_edbgf.Urx) < _bcfd && _g.Abs(_gfgd.Lly-_edbgf.Lly) < _bcfd && _g.Abs(_gfgd.Ury-_edbgf.Ury) < _bcfd {
						_ggab[_gfgd] = struct{}{}
					}
				}
			}
			if len(_ggab) > 0 {
				_defd := 0
				for _, _cfdf := range _acda {
					if _, _caac := _ggab[_cfdf]; !_caac {
						_acda[_defd] = _cfdf
						_defd++
					}
				}
				_dbag._edgf[_facf] = _acda[:len(_acda)-len(_ggab)]
				if len(_dbag._edgf[_facf]) == 0 {
					delete(_dbag._edgf, _facf)
				}
			}
		}
	}
}

func (_afcbf *textLine) text() string {
	var _ebfa []string
	for _, _caefd := range _afcbf._gfed {
		if _caefd._bace {
			_ebfa = append(_ebfa, "\u0020")
		}
		_ebfa = append(_ebfa, _caefd._faaf)
	}
	return _cb.Join(_ebfa, "")
}

// Text returns the extracted page text.
func (_dcd PageText) Text() string { return _dcd._aeb }

func _beeb(_dceea _eb.PdfObject, _aagcb _ca.Color) (_cfb.Image, error) {
	_dfed, _faeae := _eb.GetStream(_dceea)
	if !_faeae {
		return nil, nil
	}
	_bedbf, _dafaa := _ab.NewXObjectImageFromStream(_dfed)
	if _dafaa != nil {
		return nil, _dafaa
	}
	_dbdbge, _dafaa := _bedbf.ToImage()
	if _dafaa != nil {
		return nil, _dafaa
	}
	return _dfee(_dbdbge, _aagcb), nil
}

// String returns a human readable description of `ss`.
func (_dcda *shapesState) String() string {
	return _a.Sprintf("\u007b\u0025\u0064\u0020su\u0062\u0070\u0061\u0074\u0068\u0073\u0020\u0066\u0072\u0065\u0073\u0068\u003d\u0025t\u007d", len(_dcda._dagf), _dcda._caef)
}

func _gadb(_cgef map[float64]map[float64]gridTile) []float64 {
	_gdbgf := make([]float64, 0, len(_cgef))
	_fbdaf := make(map[float64]struct{}, len(_cgef))
	for _, _bade := range _cgef {
		for _fdbfb := range _bade {
			if _, _cedf := _fbdaf[_fdbfb]; _cedf {
				continue
			}
			_gdbgf = append(_gdbgf, _fdbfb)
			_fbdaf[_fdbfb] = struct{}{}
		}
	}
	_cf.Float64s(_gdbgf)
	return _gdbgf
}

func (_gfbf *TextMarkArray) getTextMarkAtOffset(_bbb int) *TextMark {
	for _, _abfb := range _gfbf._beaf {
		if _abfb.Offset == _bbb {
			return &_abfb
		}
	}
	return nil
}

// String returns a human readable description of `s`.
func (_gfef intSet) String() string {
	var _gegeb []int
	for _ccbgg := range _gfef {
		if _gfef.has(_ccbgg) {
			_gegeb = append(_gegeb, _ccbgg)
		}
	}
	_cf.Ints(_gegeb)
	return _a.Sprintf("\u0025\u002b\u0076", _gegeb)
}

// TableInfo gets table information of the textmark `tm`.
func (_cdfa *TextMark) TableInfo() (*TextTable, [][]int) {
	if !_cdfa._fddc {
		return nil, nil
	}
	_bcba := _cdfa._fabfe
	_dagc := _bcba.getCellInfo(*_cdfa)
	return _bcba, _dagc
}

func (_defba rulingList) secMinMax() (float64, float64) {
	_agadcd, _cece := _defba[0]._beadd, _defba[0]._dgbc
	for _, _eggf := range _defba[1:] {
		if _eggf._beadd < _agadcd {
			_agadcd = _eggf._beadd
		}
		if _eggf._dgbc > _cece {
			_cece = _eggf._dgbc
		}
	}
	return _agadcd, _cece
}

func _cegc(_cdd *Extractor, _gaa *_ab.PdfPageResources, _gcbd _bb.GraphicsState, _cabb *textState, _ebac *stateStack) *textObject {
	return &textObject{_bdg: _cdd, _dafd: _gaa, _fcfe: _gcbd, _beg: _ebac, _dbfd: _cabb, _dbd: _fa.IdentityMatrix(), _cdeg: _fa.IdentityMatrix()}
}

// BBox returns the smallest axis-aligned rectangle that encloses all the TextMarks in `ma`.
func (_ceeb *TextMarkArray) BBox() (_ab.PdfRectangle, bool) {
	var _cce _ab.PdfRectangle
	_efdc := false
	for _, _fgfed := range _ceeb._beaf {
		if _fgfed.Meta || _fcacf(_fgfed.Text) {
			continue
		}
		if _efdc {
			_cce = _abbb(_cce, _fgfed.BBox)
		} else {
			_cce = _fgfed.BBox
			_efdc = true
		}
	}
	return _cce, _efdc
}

type textObject struct {
	_bdg  *Extractor
	_dafd *_ab.PdfPageResources
	_fcfe _bb.GraphicsState
	_dbfd *textState
	_beg  *stateStack
	_dbd  _fa.Matrix
	_cdeg _fa.Matrix
	_cecd []*textMark
	_bfcc bool
}

func (_dfgb *textObject) moveLP(_afdd, _fce float64) {
	_dfgb._cdeg.Concat(_fa.NewMatrix(1, 0, 0, 1, _afdd, _fce))
	_dfgb._dbd = _dfgb._cdeg
}

// String returns a description of `tm`.
func (_cbfd *textMark) String() string {
	return _a.Sprintf("\u0025\u002e\u0032f \u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\u0022\u0025\u0073\u0022", _cbfd.PdfRectangle, _cbfd._acaa, _cbfd._bbeb)
}

func (_bag *textObject) getCurrentFont() *_ab.PdfFont {
	_dffc := _bag._dbfd._ggda
	if _dffc == nil {
		_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u002e\u0020U\u0073\u0069\u006e\u0067\u0020d\u0065\u0066a\u0075\u006c\u0074\u002e")
		return _ab.DefaultFont()
	}
	return _dffc
}

// String returns a string describing `pt`.
func (_afg PageText) String() string {
	_gddae := _a.Sprintf("P\u0061\u0067\u0065\u0054ex\u0074:\u0020\u0025\u0064\u0020\u0065l\u0065\u006d\u0065\u006e\u0074\u0073", len(_afg._cbf))
	_bcd := []string{"\u002d" + _gddae}
	for _, _fcca := range _afg._cbf {
		_bcd = append(_bcd, _fcca.String())
	}
	_bcd = append(_bcd, "\u002b"+_gddae)
	return _cb.Join(_bcd, "\u000a")
}

func _abag(_acg *wordBag, _aefd int) *textLine {
	_bbfb := _acg.firstWord(_aefd)
	_aggfa := textLine{PdfRectangle: _bbfb.PdfRectangle, _ead: _bbfb._fcde, _defc: _bbfb._faace}
	_aggfa.pullWord(_acg, _bbfb, _aefd)
	return &_aggfa
}

func _becf(_ecbe map[int][]float64) {
	if len(_ecbe) <= 1 {
		return
	}
	_bebaa := _eaegb(_ecbe)
	if _cabbg {
		_da.Log.Info("\u0066i\u0078C\u0065\u006c\u006c\u0073\u003a \u006b\u0065y\u0073\u003d\u0025\u002b\u0076", _bebaa)
	}
	var _cdcfa, _gfeaba int
	for _cdcfa, _gfeaba = range _bebaa {
		if _ecbe[_gfeaba] != nil {
			break
		}
	}
	for _agag, _cgccc := range _bebaa[_cdcfa:] {
		_aebaf := _ecbe[_cgccc]
		if _aebaf == nil {
			continue
		}
		if _cabbg {
			_a.Printf("\u0025\u0034\u0064\u003a\u0020\u006b\u0030\u003d\u0025\u0064\u0020\u006b1\u003d\u0025\u0064\u000a", _cdcfa+_agag, _gfeaba, _cgccc)
		}
		_ddgdd := _ecbe[_cgccc]
		if _ddgdd[len(_ddgdd)-1] > _aebaf[0] {
			_ddgdd[len(_ddgdd)-1] = _aebaf[0]
			_ecbe[_gfeaba] = _ddgdd
		}
		_gfeaba = _cgccc
	}
}
func (_aac *textObject) moveText(_bgg, _gdcf float64) { _aac.moveLP(_bgg, _gdcf) }
func (_fafb *textTable) newTablePara() *textPara {
	_bfae := _fafb.computeBbox()
	_cbfecg := &textPara{PdfRectangle: _bfae, _gedc: _bfae, _bcfcf: _fafb}
	if _cabbg {
		_da.Log.Info("\u006e\u0065w\u0054\u0061\u0062l\u0065\u0050\u0061\u0072\u0061\u003a\u0020\u0025\u0073", _cbfecg)
	}
	return _cbfecg
}

func (_beb *imageExtractContext) extractFormImages(_dfd *_eb.PdfObjectName, _agcc _bb.GraphicsState, _ebbf *_ab.PdfPageResources) error {
	_cge, _dgfd := _ebbf.GetXObjectFormByName(*_dfd)
	if _dgfd != nil {
		return _dgfd
	}
	if _cge == nil {
		return nil
	}
	_ecg, _dgfd := _cge.GetContentStream()
	if _dgfd != nil {
		return _dgfd
	}
	_bef := _cge.Resources
	if _bef == nil {
		_bef = _ebbf
	}
	_dgfd = _beb.extractContentStreamImages(string(_ecg), _bef)
	if _dgfd != nil {
		return _dgfd
	}
	_beb._dde++
	return nil
}
func _fdgf(_eeddc string) string { _ebgg := []rune(_eeddc); return string(_ebgg[:len(_ebgg)-1]) }
func _beadb(_aea, _caeb bounded) float64 {
	_dgdb := _faeff(_aea, _caeb)
	if !_ccae(_dgdb) {
		return _dgdb
	}
	return _bbaa(_aea, _caeb)
}
func _gecd(_bcgg, _eagb *textPara) bool { return _eedd(_bcgg._gedc, _eagb._gedc) }
func _efdd(_cg []Font, _eec string) bool {
	for _, _fff := range _cg {
		if _fff.FontName == _eec {
			return true
		}
	}
	return false
}

func (_gbee *textWord) toTextMarks(_dffge *int) []TextMark {
	var _ggae []TextMark
	for _, _aaef := range _gbee._bacfd {
		_ggae = _egfcc(_ggae, _dffge, _aaef.ToTextMark())
	}
	return _ggae
}

func (_baeeg paraList) findGridTables(_edaf []gridTiling) []*textTable {
	if _cabbg {
		_da.Log.Info("\u0066i\u006e\u0064\u0047\u0072\u0069\u0064\u0054\u0061\u0062\u006c\u0065s\u003a\u0020\u0025\u0064\u0020\u0070\u0061\u0072\u0061\u0073", len(_baeeg))
		for _deaed, _dfbd := range _baeeg {
			_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _deaed, _dfbd)
		}
	}
	var _aabae []*textTable
	for _decc, _cbbad := range _edaf {
		_fdecg, _bdecc := _baeeg.findTableGrid(_cbbad)
		if _fdecg != nil {
			_fdecg.log(_a.Sprintf("\u0066\u0069\u006e\u0064Ta\u0062\u006c\u0065\u0057\u0069\u0074\u0068\u0047\u0072\u0069\u0064\u0073\u003a\u0020%\u0064", _decc))
			_aabae = append(_aabae, _fdecg)
			_fdecg.markCells()
		}
		for _bgcfd := range _bdecc {
			_bgcfd._aagd = true
		}
	}
	if _cabbg {
		_da.Log.Info("\u0066i\u006e\u0064\u0047\u0072i\u0064\u0054\u0061\u0062\u006ce\u0073:\u0020%\u0064\u0020\u0074\u0061\u0062\u006c\u0065s", len(_aabae))
	}
	return _aabae
}

func _dbaf(_efgf []*textWord, _gagd float64, _cdc, _gbgeb rulingList) *wordBag {
	_bccg := _daaga(_efgf[0], _gagd, _cdc, _gbgeb)
	for _, _agca := range _efgf[1:] {
		_eafa := _aeef(_agca._faace)
		_bccg._edgf[_eafa] = append(_bccg._edgf[_eafa], _agca)
		_bccg.PdfRectangle = _abbb(_bccg.PdfRectangle, _agca.PdfRectangle)
	}
	_bccg.sort()
	return _bccg
}

func (_gfbeg paraList) findTextTables() []*textTable {
	var _gdgag []*textTable
	for _, _fafge := range _gfbeg {
		if _fafge.taken() || _fafge.Width() == 0 {
			continue
		}
		_aeada := _fafge.isAtom()
		if _aeada == nil {
			continue
		}
		_aeada.growTable()
		if _aeada._cbcb*_aeada._faed < _afdc {
			continue
		}
		_aeada.markCells()
		_aeada.log("\u0067\u0072\u006fw\u006e")
		_gdgag = append(_gdgag, _aeada)
	}
	return _gdgag
}

func (_fdbde paraList) addNeighbours() {
	_egag := func(_gdff []int, _cebae *textPara) ([]*textPara, []*textPara) {
		_cdba := make([]*textPara, 0, len(_gdff)-1)
		_gacdg := make([]*textPara, 0, len(_gdff)-1)
		for _, _agdff := range _gdff {
			_cefc := _fdbde[_agdff]
			if _cefc.Urx <= _cebae.Llx {
				_cdba = append(_cdba, _cefc)
			} else if _cefc.Llx >= _cebae.Urx {
				_gacdg = append(_gacdg, _cefc)
			}
		}
		return _cdba, _gacdg
	}
	_egbbc := func(_deffa []int, _fbgc *textPara) ([]*textPara, []*textPara) {
		_cddf := make([]*textPara, 0, len(_deffa)-1)
		_fggae := make([]*textPara, 0, len(_deffa)-1)
		for _, _eeabf := range _deffa {
			_effe := _fdbde[_eeabf]
			if _effe.Ury <= _fbgc.Lly {
				_fggae = append(_fggae, _effe)
			} else if _effe.Lly >= _fbgc.Ury {
				_cddf = append(_cddf, _effe)
			}
		}
		return _cddf, _fggae
	}
	_dccc := _fdbde.yNeighbours(_fdae)
	for _, _caaf := range _fdbde {
		_gcbac := _dccc[_caaf]
		if len(_gcbac) == 0 {
			continue
		}
		_fggcd, _fege := _egag(_gcbac, _caaf)
		if len(_fggcd) == 0 && len(_fege) == 0 {
			continue
		}
		if len(_fggcd) > 0 {
			_ggbb := _fggcd[0]
			for _, _ddgf := range _fggcd[1:] {
				if _ddgf.Urx >= _ggbb.Urx {
					_ggbb = _ddgf
				}
			}
			for _, _acfa := range _fggcd {
				if _acfa != _ggbb && _acfa.Urx > _ggbb.Llx {
					_ggbb = nil
					break
				}
			}
			if _ggbb != nil && _gdac(_caaf.PdfRectangle, _ggbb.PdfRectangle) {
				_caaf._fdccb = _ggbb
			}
		}
		if len(_fege) > 0 {
			_gdbf := _fege[0]
			for _, _ebcd := range _fege[1:] {
				if _ebcd.Llx <= _gdbf.Llx {
					_gdbf = _ebcd
				}
			}
			for _, _degec := range _fege {
				if _degec != _gdbf && _degec.Llx < _gdbf.Urx {
					_gdbf = nil
					break
				}
			}
			if _gdbf != nil && _gdac(_caaf.PdfRectangle, _gdbf.PdfRectangle) {
				_caaf._eefcd = _gdbf
			}
		}
	}
	_dccc = _fdbde.xNeighbours(_daaag)
	for _, _fged := range _fdbde {
		_cdff := _dccc[_fged]
		if len(_cdff) == 0 {
			continue
		}
		_dgeb, _ggbg := _egbbc(_cdff, _fged)
		if len(_dgeb) == 0 && len(_ggbg) == 0 {
			continue
		}
		if len(_ggbg) > 0 {
			_abdd := _ggbg[0]
			for _, _gafe := range _ggbg[1:] {
				if _gafe.Ury >= _abdd.Ury {
					_abdd = _gafe
				}
			}
			for _, _fddgad := range _ggbg {
				if _fddgad != _abdd && _fddgad.Ury > _abdd.Lly {
					_abdd = nil
					break
				}
			}
			if _abdd != nil && _eedd(_fged.PdfRectangle, _abdd.PdfRectangle) {
				_fged._afeeg = _abdd
			}
		}
		if len(_dgeb) > 0 {
			_gadgc := _dgeb[0]
			for _, _eage := range _dgeb[1:] {
				if _eage.Lly <= _gadgc.Lly {
					_gadgc = _eage
				}
			}
			for _, _bbfef := range _dgeb {
				if _bbfef != _gadgc && _bbfef.Lly < _gadgc.Ury {
					_gadgc = nil
					break
				}
			}
			if _gadgc != nil && _eedd(_fged.PdfRectangle, _gadgc.PdfRectangle) {
				_fged._aebbe = _gadgc
			}
		}
	}
	for _, _bagbg := range _fdbde {
		if _bagbg._fdccb != nil && _bagbg._fdccb._eefcd != _bagbg {
			_bagbg._fdccb = nil
		}
		if _bagbg._aebbe != nil && _bagbg._aebbe._afeeg != _bagbg {
			_bagbg._aebbe = nil
		}
		if _bagbg._eefcd != nil && _bagbg._eefcd._fdccb != _bagbg {
			_bagbg._eefcd = nil
		}
		if _bagbg._afeeg != nil && _bagbg._afeeg._aebbe != _bagbg {
			_bagbg._afeeg = nil
		}
	}
}

func _bfef(_eacf []float64, _bgbf, _bdeb float64) []float64 {
	_dcgd, _gage := _bgbf, _bdeb
	if _gage < _dcgd {
		_dcgd, _gage = _gage, _dcgd
	}
	_adefg := make([]float64, 0, len(_eacf)+2)
	_adefg = append(_adefg, _bgbf)
	for _, _ggdd := range _eacf {
		if _ggdd <= _dcgd {
			continue
		} else if _ggdd >= _gage {
			break
		}
		_adefg = append(_adefg, _ggdd)
	}
	_adefg = append(_adefg, _bdeb)
	return _adefg
}

// ExtractTextWithStats works like ExtractText but returns the number of characters in the output
// (`numChars`) and the number of characters that were not decoded (`numMisses`).
func (_dad *Extractor) ExtractTextWithStats() (_cffg string, _cab int, _eeg int, _ac error) {
	_afe, _cab, _eeg, _ac := _dad.ExtractPageText()
	if _ac != nil {
		return "", _cab, _eeg, _ac
	}
	return _afe.Text(), _cab, _eeg, nil
}

const (
	RenderModeStroke RenderMode = 1 << iota
	RenderModeFill
	RenderModeClip
)

func (_eeaa rulingList) snapToGroupsDirection() rulingList {
	_eeaa.sortStrict()
	_eccfe := make(map[*ruling]rulingList, len(_eeaa))
	_afdg := _eeaa[0]
	_dgfgd := func(_ccca *ruling) { _afdg = _ccca; _eccfe[_afdg] = rulingList{_ccca} }
	_dgfgd(_eeaa[0])
	for _, _effd := range _eeaa[1:] {
		if _effd._dbbce < _afdg._dbbce-_gccbd {
			_da.Log.Error("\u0073\u006e\u0061\u0070T\u006f\u0047\u0072\u006f\u0075\u0070\u0073\u0044\u0069r\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0057\u0072\u006f\u006e\u0067\u0020\u0070\u0072\u0069\u006da\u0072\u0079\u0020\u006f\u0072d\u0065\u0072\u002e\u000a\u0009\u0076\u0030\u003d\u0025\u0073\u000a\u0009\u0020\u0076\u003d\u0025\u0073", _afdg, _effd)
		}
		if _effd._dbbce > _afdg._dbbce+_efbg {
			_dgfgd(_effd)
		} else {
			_eccfe[_afdg] = append(_eccfe[_afdg], _effd)
		}
	}
	_ddgcc := make(map[*ruling]float64, len(_eccfe))
	_beac := make(map[*ruling]*ruling, len(_eeaa))
	for _agecf, _bccec := range _eccfe {
		_ddgcc[_agecf] = _bccec.mergePrimary()
		for _, _cfga := range _bccec {
			_beac[_cfga] = _agecf
		}
	}
	for _, _edbfd := range _eeaa {
		_edbfd._dbbce = _ddgcc[_beac[_edbfd]]
	}
	_ecac := make(rulingList, 0, len(_eeaa))
	for _, _dgde := range _eccfe {
		_gbab := _dgde.splitSec()
		for _cgefe, _febdb := range _gbab {
			_gfeab := _febdb.merge()
			if len(_ecac) > 0 {
				_edcd := _ecac[len(_ecac)-1]
				if _edcd.alignsPrimary(_gfeab) && _edcd.alignsSec(_gfeab) {
					_da.Log.Error("\u0073\u006e\u0061\u0070\u0054\u006fG\u0072\u006f\u0075\u0070\u0073\u0044\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0044\u0075\u0070\u006ci\u0063\u0061\u0074\u0065\u0020\u0069\u003d\u0025\u0064\u000a\u0009\u0077\u003d\u0025s\u000a\t\u0076\u003d\u0025\u0073", _cgefe, _edcd, _gfeab)
					continue
				}
			}
			_ecac = append(_ecac, _gfeab)
		}
	}
	_ecac.sortStrict()
	return _ecac
}

func (_bebf *textObject) setTextMatrix(_cda []float64) {
	if len(_cda) != 6 {
		_da.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u006c\u0065\u006e\u0028\u0066\u0029\u0020\u0021\u003d\u0020\u0036\u0020\u0028\u0025\u0064\u0029", len(_cda))
		return
	}
	_fdab, _ebbg, _ddc, _gba, _gcbc, _ffab := _cda[0], _cda[1], _cda[2], _cda[3], _cda[4], _cda[5]
	_bebf._dbd = _fa.NewMatrix(_fdab, _ebbg, _ddc, _gba, _gcbc, _ffab)
	_bebf._cdeg = _bebf._dbd
}

type structElement struct {
	_acfd string
	_baac []structElement
	_bbgg int64
	_ffcf _eb.PdfObject
}

func (_degb *shapesState) clearPath() {
	_degb._dagf = nil
	_degb._caef = false
	if _caad {
		_da.Log.Info("\u0043\u004c\u0045A\u0052\u003a\u0020\u0073\u0073\u003d\u0025\u0073", _degb)
	}
}

func (_agaff paraList) writeText(_gdbb _f.Writer) {
	for _eedg, _aegd := range _agaff {
		if _aegd._ccfeb {
			continue
		}
		_aegd.writeText(_gdbb)
		if _eedg != len(_agaff)-1 {
			if _ceed(_aegd, _agaff[_eedg+1]) {
				_gdbb.Write([]byte("\u0020"))
			} else {
				_gdbb.Write([]byte("\u000a"))
				_gdbb.Write([]byte("\u000a"))
			}
		}
	}
	_gdbb.Write([]byte("\u000a"))
	_gdbb.Write([]byte("\u000a"))
}

func _dfcg(_dbbd *list) []*list {
	var _ccbf []*list
	for _, _adad := range _dbbd._ffaf {
		switch _adad._dbga {
		case "\u004c\u0049":
			_begb := _ccbfd(_adad)
			_fefg := _dfcg(_adad)
			_gdad := _defbb(_begb, "\u0062\u0075\u006c\u006c\u0065\u0074", _fefg)
			_bfdg := _aeag(_begb, "")
			_gdad._gfec = _bfdg
			_ccbf = append(_ccbf, _gdad)
		case "\u004c\u0042\u006fd\u0079":
			return _dfcg(_adad)
		case "\u004c":
			_bcad := _dfcg(_adad)
			_ccbf = append(_ccbf, _bcad...)
			return _ccbf
		}
	}
	return _ccbf
}
func (_dbbea lineRuling) xMean() float64 { return 0.5 * (_dbbea._bbfe.X + _dbbea._eaega.X) }
func (_aagaa *textTable) markCells() {
	for _acbed := 0; _acbed < _aagaa._faed; _acbed++ {
		for _fgec := 0; _fgec < _aagaa._cbcb; _fgec++ {
			_gddfb := _aagaa.get(_fgec, _acbed)
			if _gddfb != nil {
				_gddfb._aagd = true
			}
		}
	}
}

// ApplyArea processes the page text only within the specified area `bbox`.
// Each time ApplyArea is called, it updates the result set in `pt`.
// Can be called multiple times in a row with different bounding boxes.
func (_ccbg *PageText) ApplyArea(bbox _ab.PdfRectangle) {
	_fef := make([]*textMark, 0, len(_ccbg._cbf))
	for _, _dacb := range _ccbg._cbf {
		if _gega(_dacb.bbox(), bbox) {
			_fef = append(_fef, _dacb)
		}
	}
	var _bda paraList
	_gcafb := len(_fef)
	for _ffag := 0; _ffag < 360 && _gcafb > 0; _ffag += 90 {
		_egdb := make([]*textMark, 0, len(_fef)-_gcafb)
		for _, _geba := range _fef {
			if _geba._fbfge == _ffag {
				_egdb = append(_egdb, _geba)
			}
		}
		if len(_egdb) > 0 {
			_dffg := _eefc(_egdb, _ccbg._cfce, nil, nil, _ccbg._cbef._dccbb)
			_bda = append(_bda, _dffg...)
			_gcafb -= len(_egdb)
		}
	}
	_fdfd := new(_gc.Buffer)
	_bda.writeText(_fdfd)
	_ccbg._aeb = _fdfd.String()
	_ccbg._gddg = _bda.toTextMarks()
	_ccbg._aee = _bda.tables()
}

func (_fcgea *textPara) text() string {
	_abcb := new(_gc.Buffer)
	_fcgea.writeText(_abcb)
	return _abcb.String()
}

func (_ffdgb paraList) findTableGrid(_cefa gridTiling) (*textTable, map[*textPara]struct{}) {
	_fcggb := len(_cefa._cgeba)
	_dbcab := len(_cefa._gfdac)
	_bcfcbe := textTable{_cafe: true, _cbcb: _fcggb, _faed: _dbcab, _ebgbd: make(map[uint64]*textPara, _fcggb*_dbcab), _fcda: make(map[uint64]compositeCell, _fcggb*_dbcab)}
	_bcfcbe.PdfRectangle = _cefa.PdfRectangle
	_eaga := make(map[*textPara]struct{})
	_bbdde := int((1.0 - _bfbd) * float64(_fcggb*_dbcab))
	_fbed := 0
	if _bcf {
		_da.Log.Info("\u0066\u0069\u006e\u0064Ta\u0062\u006c\u0065\u0047\u0072\u0069\u0064\u003a\u0020\u0025\u0064\u0020\u0078\u0020%\u0064", _fcggb, _dbcab)
	}
	for _ddfc, _eeag := range _cefa._gfdac {
		_dccbea, _cdac := _cefa._afce[_eeag]
		if !_cdac {
			continue
		}
		for _bcgfg, _addbd := range _cefa._cgeba {
			_aedg, _aace := _dccbea[_addbd]
			if !_aace {
				continue
			}
			_abbgb := _ffdgb.inTile(_aedg)
			if len(_abbgb) == 0 {
				_fbed++
				if _fbed > _bbdde {
					if _bcf {
						_da.Log.Info("\u0021\u006e\u0075m\u0045\u006d\u0070\u0074\u0079\u003d\u0025\u0064", _fbed)
					}
					return nil, nil
				}
			} else {
				_bcfcbe.putComposite(_bcgfg, _ddfc, _abbgb, _aedg.PdfRectangle)
				for _, _ccdfd := range _abbgb {
					_eaga[_ccdfd] = struct{}{}
				}
			}
		}
	}
	_dfda := 0
	for _acgeg := 0; _acgeg < _fcggb; _acgeg++ {
		_faaeg := _bcfcbe.get(_acgeg, 0)
		if _faaeg == nil || !_faaeg._ccfeb {
			_dfda++
		}
	}
	if _dfda == 0 {
		if _bcf {
			_da.Log.Info("\u0021\u006e\u0075m\u0048\u0065\u0061\u0064\u0065\u0072\u003d\u0030")
		}
		return nil, nil
	}
	_dbdgf := _bcfcbe.reduceTiling(_cefa, _agcf)
	_dbdgf = _dbdgf.subdivide()
	return _dbdgf, _eaga
}

func (_dbec *shapesState) quadraticTo(_aba, _efcad, _eabg, _fgda float64) {
	if _caad {
		_da.Log.Info("\u0071\u0075\u0061d\u0072\u0061\u0074\u0069\u0063\u0054\u006f\u003a")
	}
	_dbec.addPoint(_eabg, _fgda)
}

func (_degde paraList) extractTables(_bgddf []gridTiling) paraList {
	if _cabbg {
		_da.Log.Debug("\u0065\u0078\u0074r\u0061\u0063\u0074\u0054\u0061\u0062\u006c\u0065\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_degde))
	}
	if len(_degde) < _afdc {
		return _degde
	}
	_efffd := _degde.findTables(_bgddf)
	if _cabbg {
		_da.Log.Info("c\u006f\u006d\u0062\u0069\u006e\u0065d\u0020\u0074\u0061\u0062\u006c\u0065s\u0020\u0025\u0064\u0020\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d", len(_efffd))
		for _eeddg, _eabga := range _efffd {
			_eabga.log(_a.Sprintf("c\u006f\u006d\u0062\u0069\u006e\u0065\u0064\u0020\u0025\u0064", _eeddg))
		}
	}
	return _degde.applyTables(_efffd)
}

// Text returns the text content of the `bulletLists`.
func (_ffb *lists) Text() string {
	_gafd := &_cb.Builder{}
	for _, _gfgc := range *_ffb {
		_baag := _gfgc.Text()
		_gafd.WriteString(_baag)
	}
	return _gafd.String()
}

func (_gfafd *textLine) appendWord(_acfg *textWord) {
	_gfafd._gfed = append(_gfafd._gfed, _acfg)
	_gfafd.PdfRectangle = _abbb(_gfafd.PdfRectangle, _acfg.PdfRectangle)
	if _acfg._fcde > _gfafd._ead {
		_gfafd._ead = _acfg._fcde
	}
	if _acfg._faace > _gfafd._defc {
		_gfafd._defc = _acfg._faace
	}
}
func (_abbae lineRuling) yMean() float64 { return 0.5 * (_abbae._bbfe.Y + _abbae._eaega.Y) }
func (_abfd *wordBag) depthBand(_dbfe, _bged float64) []int {
	if len(_abfd._edgf) == 0 {
		return nil
	}
	return _abfd.depthRange(_abfd.getDepthIdx(_dbfe), _abfd.getDepthIdx(_bged))
}

var _ddbc = map[rulingKind]string{_fbag: "\u006e\u006f\u006e\u0065", _ceedg: "\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c", _eccce: "\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c"}

func (_debgg rulingList) tidied(_cfffa string) rulingList {
	_afgaa := _debgg.removeDuplicates()
	_afgaa.log("\u0075n\u0069\u0071\u0075\u0065\u0073")
	_ggcaa := _afgaa.snapToGroups()
	if _ggcaa == nil {
		return nil
	}
	_ggcaa.sort()
	if _egc {
		_da.Log.Info("\u0074\u0069\u0064i\u0065\u0064\u003a\u0020\u0025\u0071\u0020\u0076\u0065\u0063\u0073\u003d\u0025\u0064\u0020\u0075\u006e\u0069\u0071\u0075\u0065\u0073\u003d\u0025\u0064\u0020\u0063\u006f\u0061l\u0065\u0073\u0063\u0065\u0064\u003d\u0025\u0064", _cfffa, len(_debgg), len(_afgaa), len(_ggcaa))
	}
	_ggcaa.log("\u0063o\u0061\u006c\u0065\u0073\u0063\u0065d")
	return _ggcaa
}

func (_ccee *wordBag) depthRange(_ddddc, _fdca int) []int {
	var _eaeb []int
	for _cffdd := range _ccee._edgf {
		if _ddddc <= _cffdd && _cffdd <= _fdca {
			_eaeb = append(_eaeb, _cffdd)
		}
	}
	if len(_eaeb) == 0 {
		return nil
	}
	_cf.Ints(_eaeb)
	return _eaeb
}

func _gbccg(_eafdc, _gaaae _fa.Point, _cafac _ca.Color) (*ruling, bool) {
	_gfafa := lineRuling{_bbfe: _eafdc, _eaega: _gaaae, _agbc: _fbbae(_eafdc, _gaaae), Color: _cafac}
	if _gfafa._agbc == _fbag {
		return nil, false
	}
	return _gfafa.asRuling()
}

func _bbcdg(_eeec *textLine) bool {
	_eafcc := true
	_adgge := -1
	for _, _fdcc := range _eeec._gfed {
		for _, _ffbd := range _fdcc._bacfd {
			_fdabf := _ffbd._gbdg
			if _adgge == -1 {
				_adgge = _fdabf
			} else {
				if _adgge != _fdabf {
					_eafcc = false
					break
				}
			}
		}
	}
	return _eafcc
}

// PageImages represents extracted images on a PDF page with spatial information:
// display position and size.
type PageImages struct{ Images []ImageMark }

// ExtractText processes and extracts all text data in content streams and returns as a string.
// It takes into account character encodings in the PDF file, which are decoded by
// CharcodeBytesToUnicode.
// Characters that can't be decoded are replaced with MissingCodeRune ('\ufffd' = �).
func (_gfd *Extractor) ExtractText() (string, error) {
	_gfb, _, _, _aaa := _gfd.ExtractTextWithStats()
	return _gfb, _aaa
}

func _agdb(_bgab, _gbage _fa.Point) bool {
	_cdcbe := _g.Abs(_bgab.X - _gbage.X)
	_adgbd := _g.Abs(_bgab.Y - _gbage.Y)
	return _dbdg(_cdcbe, _adgbd)
}

func (_cfbe rulingList) bbox() _ab.PdfRectangle {
	var _ffaeg _ab.PdfRectangle
	if len(_cfbe) == 0 {
		_da.Log.Error("r\u0075\u006c\u0069\u006e\u0067\u004ci\u0073\u0074\u002e\u0062\u0062\u006f\u0078\u003a\u0020n\u006f\u0020\u0072u\u006ci\u006e\u0067\u0073")
		return _ab.PdfRectangle{}
	}
	if _cfbe[0]._ggcaf == _ceedg {
		_ffaeg.Llx, _ffaeg.Urx = _cfbe.secMinMax()
		_ffaeg.Lly, _ffaeg.Ury = _cfbe.primMinMax()
	} else {
		_ffaeg.Llx, _ffaeg.Urx = _cfbe.primMinMax()
		_ffaeg.Lly, _ffaeg.Ury = _cfbe.secMinMax()
	}
	return _ffaeg
}

func (_bdacc rulingList) log(_dgaa string) {
	if !_egc {
		return
	}
	_da.Log.Info("\u0023\u0023\u0023\u0020\u0025\u0031\u0030\u0073\u003a\u0020\u0076\u0065c\u0073\u003d\u0025\u0073", _dgaa, _bdacc.String())
	for _eefe, _agddb := range _bdacc {
		_a.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _eefe, _agddb.String())
	}
}

func _ccfb(_afde _eb.PdfObject, _aeeg _ca.Color) (_cfb.Image, error) {
	_gdadb, _dggdg := _eb.GetStream(_afde)
	if !_dggdg {
		return nil, nil
	}
	_abddf, _aaag := _ab.NewXObjectImageFromStream(_gdadb)
	if _aaag != nil {
		return nil, _aaag
	}
	_egad, _aaag := _abddf.ToImage()
	if _aaag != nil {
		return nil, _aaag
	}
	return _bgceg(_egad, _aeeg), nil
}
