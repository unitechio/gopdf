package extractor

import (
	_cf "bytes"
	_c "errors"
	_cfe "fmt"
	_gf "image"
	_eb "image/color"
	_a "io"
	_df "math"
	_g "reflect"
	_ca "regexp"
	_d "sort"
	_ag "strings"
	_ac "unicode"
	_cb "unicode/utf8"

	_ed "bitbucket.org/shenghui0779/gopdf/common"
	_aad "bitbucket.org/shenghui0779/gopdf/contentstream"
	_gg "bitbucket.org/shenghui0779/gopdf/core"
	_dd "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_cc "bitbucket.org/shenghui0779/gopdf/internal/transform"
	_gb "bitbucket.org/shenghui0779/gopdf/model"
	_dc "golang.org/x/image/draw"
	_f "golang.org/x/text/unicode/norm"
	_aa "golang.org/x/xerrors"
)

func _feee(_ebbc, _gbecb _gb.PdfRectangle) bool { return _cdcd(_ebbc, _gbecb) && _ebgc(_ebbc, _gbecb) }
func _bba(_ced _cc.Matrix) _cc.Point {
	_daef, _dfdcc := _ced.Translation()
	return _cc.Point{X: _daef, Y: _dfdcc}
}
func (_becb *shapesState) clearPath() {
	_becb._babc = nil
	_becb._dfff = false
	if _dgeg {
		_ed.Log.Info("\u0043\u004c\u0045A\u0052\u003a\u0020\u0073\u0073\u003d\u0025\u0073", _becb)
	}
}

var _gabd string = "\u005e\u005b\u0061\u002d\u007a\u0041\u002dZ\u005d\u0028\u005c)\u007c\u005c\u002e)\u007c\u005e[\u005c\u0064\u005d\u002b\u0028\u005c)\u007c\\.\u0029\u007c\u005e\u005c\u0028\u005b\u0061\u002d\u007a\u0041\u002d\u005a\u005d\u005c\u0029\u007c\u005e\u005c\u0028\u005b\u005c\u0064\u005d\u002b\u005c\u0029"

// TableCell is a cell in a TextTable.
type TableCell struct {
	_gb.PdfRectangle

	// Text is the extracted text.
	Text string

	// Marks returns the TextMarks corresponding to the text in Text.
	Marks TextMarkArray
}

func (_gcbgb gridTiling) complete() bool {
	for _, _fdeg := range _gcbgb._aaef {
		for _, _efeee := range _fdeg {
			if !_efeee.complete() {
				return false
			}
		}
	}
	return true
}
func (_bgcfc paraList) llyRange(_fbcc []int, _cgfab, _ggeb float64) []int {
	_baac := len(_bgcfc)
	if _ggeb < _bgcfc[_fbcc[0]].Lly || _cgfab > _bgcfc[_fbcc[_baac-1]].Lly {
		return nil
	}
	_ebed := _d.Search(_baac, func(_bgef int) bool { return _bgcfc[_fbcc[_bgef]].Lly >= _cgfab })
	_edbf := _d.Search(_baac, func(_fbfa int) bool { return _bgcfc[_fbcc[_fbfa]].Lly > _ggeb })
	return _fbcc[_ebed:_edbf]
}

type textLine struct {
	_gb.PdfRectangle
	_bdbg float64
	_bfca []*textWord
	_egbd float64
}

func (_dbadg paraList) inTile(_dbac gridTile) paraList {
	var _bcbcf paraList
	for _, _eaedg := range _dbadg {
		if _dbac.contains(_eaedg.PdfRectangle) {
			_bcbcf = append(_bcbcf, _eaedg)
		}
	}
	if _geff {
		_cfe.Printf("\u0020 \u0020\u0069\u006e\u0054i\u006c\u0065\u003a\u0020\u0020%\u0073 \u0069n\u0073\u0069\u0064\u0065\u003d\u0025\u0064\n", _dbac, len(_bcbcf))
		for _feffd, _cefdb := range _bcbcf {
			_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _feffd, _cefdb)
		}
		_cfe.Println("")
	}
	return _bcbcf
}
func (_efe *Extractor) extractPageText(_eg string, _gfe *_gb.PdfPageResources, _agb _cc.Matrix, _gfef int) (*PageText, int, int, error) {
	_ed.Log.Trace("\u0065x\u0074\u0072\u0061\u0063t\u0050\u0061\u0067\u0065\u0054e\u0078t\u003a \u006c\u0065\u0076\u0065\u006c\u003d\u0025d", _gfef)
	_ccc := &PageText{_ebae: _efe._eda, _ddbd: _efe._cd, _ddeg: _efe._dfd}
	_gdg := _beba(_efe._eda)
	var _bg stateStack
	_bcc := _gdag(_efe, _gfe, _aad.GraphicsState{}, &_gdg, &_bg)
	_eacc := shapesState{_begf: _agb, _cca: _cc.IdentityMatrix(), _geba: _bcc}
	var _edac bool
	_adb := -1
	if _gfef > _dgd {
		_edg := _c.New("\u0066\u006f\u0072\u006d s\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077")
		_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0065\u0078\u0074\u0072\u0061\u0063\u0074\u0050\u0061\u0067\u0065\u0054\u0065\u0078\u0074\u002e\u0020\u0072\u0065\u0063u\u0072\u0073\u0069\u006f\u006e\u0020\u006c\u0065\u0076\u0065\u006c\u003d\u0025\u0064 \u0065r\u0072\u003d\u0025\u0076", _gfef, _edg)
		return _ccc, _gdg._fdbd, _gdg._eec, _edg
	}
	_baf := _aad.NewContentStreamParser(_eg)
	_gbe, _fdgb := _baf.Parse()
	if _fdgb != nil {
		_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020e\u0078\u0074\u0072a\u0063\u0074\u0050\u0061g\u0065\u0054\u0065\u0078\u0074\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fdgb)
		return _ccc, _gdg._fdbd, _gdg._eec, _fdgb
	}
	_ccc._dff = _gbe
	_dcd := _aad.NewContentStreamProcessor(*_gbe)
	_dcd.AddHandler(_aad.HandlerConditionEnumAllOperands, "", func(_bbd *_aad.ContentStreamOperation, _gbeb _aad.GraphicsState, _ffc *_gb.PdfPageResources) error {
		_fda := _bbd.Operand
		if _adea {
			_ed.Log.Info("\u0026&\u0026\u0020\u006f\u0070\u003d\u0025s", _bbd)
		}
		switch _fda {
		case "\u0071":
			if _dgeg {
				_ed.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _eacc._cca)
			}
			_bg.push(&_gdg)
		case "\u0051":
			if !_bg.empty() {
				_gdg = *_bg.pop()
			}
			_eacc._cca = _gbeb.CTM
			if _dgeg {
				_ed.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _eacc._cca)
			}
		case "\u0042\u0044\u0043":
			_ddcd, _acde := _gg.GetDict(_bbd.Params[1])
			if !_acde {
				_ed.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0042D\u0043\u0020\u006f\u0070\u003d\u0025\u0073 \u0047\u0065\u0074\u0044\u0069\u0063\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _bbd)
				return _fdgb
			}
			_ccce := _ddcd.Get("\u004d\u0043\u0049\u0044")
			if _ccce != nil {
				_bbf, _ga := _gg.GetIntVal(_ccce)
				if !_ga {
					_ed.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0042\u0044C\u0020\u006f\u0070=\u0025\u0073\u002e\u0020\u0042\u0061\u0064\u0020\u006eum\u0065\u0072\u0069c\u0061\u006c \u006f\u0062\u006a\u0065\u0063\u0074.\u0020\u006f=\u0025\u0073", _bbd, _ccce)
				}
				_adb = _bbf
			} else {
				_adb = -1
			}
		case "\u0045\u004d\u0043":
			_adb = -1
		case "\u0042\u0054":
			if _edac {
				_ed.Log.Debug("\u0042\u0054\u0020\u0063\u0061\u006c\u006c\u0065\u0064\u0020\u0077\u0068\u0069\u006c\u0065 \u0069n\u0020\u0061\u0020\u0074\u0065\u0078\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
				_ccc._bagf = append(_ccc._bagf, _bcc._gafa...)
			}
			_edac = true
			_fadg := _gbeb
			_fadg.CTM = _agb.Mult(_fadg.CTM)
			_bcc = _gdag(_efe, _ffc, _fadg, &_gdg, &_bg)
			_eacc._geba = _bcc
		case "\u0045\u0054":
			if !_edac {
				_ed.Log.Debug("\u0045\u0054\u0020ca\u006c\u006c\u0065\u0064\u0020\u006f\u0075\u0074\u0073i\u0064e\u0020o\u0066 \u0061\u0020\u0074\u0065\u0078\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
			}
			_edac = false
			_ccc._bagf = append(_ccc._bagf, _bcc._gafa...)
			_bcc.reset()
		case "\u0054\u002a":
			_bcc.nextLine()
		case "\u0054\u0064":
			if _gdf, _ebd := _bcc.checkOp(_bbd, 2, true); !_gdf {
				_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ebd)
				return _ebd
			}
			_ffa, _ebfd, _bgb := _abee(_bbd.Params)
			if _bgb != nil {
				return _bgb
			}
			_bcc.moveText(_ffa, _ebfd)
		case "\u0054\u0044":
			if _feda, _dca := _bcc.checkOp(_bbd, 2, true); !_feda {
				_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dca)
				return _dca
			}
			_ebdd, _edf, _gcc := _abee(_bbd.Params)
			if _gcc != nil {
				_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gcc)
				return _gcc
			}
			_bcc.moveTextSetLeading(_ebdd, _edf)
		case "\u0054\u006a":
			if _abgb, _ebbf := _bcc.checkOp(_bbd, 1, true); !_abgb {
				_ed.Log.Debug("\u0045\u0052\u0052\u004fR:\u0020\u0054\u006a\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0065\u0072\u0072\u003d%\u0076", _bbd, _ebbf)
				return _ebbf
			}
			_cdf := _gg.TraceToDirectObject(_bbd.Params[0])
			_bcf, _dgcc := _gg.GetStringBytes(_cdf)
			if !_dgcc {
				_ed.Log.Debug("\u0045\u0052R\u004f\u0052\u003a\u0020T\u006a\u0020o\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074S\u0074\u0072\u0069\u006e\u0067\u0042\u0079\u0074\u0065\u0073\u0020\u0066a\u0069\u006c\u0065\u0064", _bbd)
				return _gg.ErrTypeError
			}
			return _bcc.showText(_cdf, _bcf, _adb)
		case "\u0054\u004a":
			if _dae, _fade := _bcc.checkOp(_bbd, 1, true); !_dae {
				_ed.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u004a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _fade)
				return _fade
			}
			_egg, _bag := _gg.GetArray(_bbd.Params[0])
			if !_bag {
				_ed.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0054\u004a\u0020\u006f\u0070\u003d\u0025s\u0020G\u0065t\u0041r\u0072\u0061\u0079\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _bbd)
				return _fdgb
			}
			return _bcc.showTextAdjusted(_egg, _adb)
		case "\u0027":
			if _agc, _dge := _bcc.checkOp(_bbd, 1, true); !_agc {
				_ed.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u0027\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dge)
				return _dge
			}
			_dbd := _gg.TraceToDirectObject(_bbd.Params[0])
			_gegec, _ccff := _gg.GetStringBytes(_dbd)
			if !_ccff {
				_ed.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020'\u0020\u006f\u0070\u003d%s \u0047et\u0053\u0074\u0072\u0069\u006e\u0067\u0042yt\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064", _bbd)
				return _gg.ErrTypeError
			}
			_bcc.nextLine()
			return _bcc.showText(_dbd, _gegec, _adb)
		case "\u0022":
			if _aec, _ggd := _bcc.checkOp(_bbd, 3, true); !_aec {
				_ed.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u0022\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ggd)
				return _ggd
			}
			_edb, _eafa, _dcf := _abee(_bbd.Params[:2])
			if _dcf != nil {
				return _dcf
			}
			_adaf := _gg.TraceToDirectObject(_bbd.Params[2])
			_becf, _gec := _gg.GetStringBytes(_adaf)
			if !_gec {
				_ed.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020\"\u0020\u006f\u0070\u003d%s \u0047et\u0053\u0074\u0072\u0069\u006e\u0067\u0042yt\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064", _bbd)
				return _gg.ErrTypeError
			}
			_bcc.setCharSpacing(_edb)
			_bcc.setWordSpacing(_eafa)
			_bcc.nextLine()
			return _bcc.showText(_adaf, _becf, _adb)
		case "\u0054\u004c":
			_egf, _egb := _bgc(_bbd)
			if _egb != nil {
				_ed.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u004c\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _egb)
				return _egb
			}
			_bcc.setTextLeading(_egf)
		case "\u0054\u0063":
			_fff, _adg := _bgc(_bbd)
			if _adg != nil {
				_ed.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0063\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _adg)
				return _adg
			}
			_bcc.setCharSpacing(_fff)
		case "\u0054\u0066":
			if _bdad, _gdbe := _bcc.checkOp(_bbd, 2, true); !_bdad {
				_ed.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0066\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gdbe)
				return _gdbe
			}
			_dbf, _cdc := _gg.GetNameVal(_bbd.Params[0])
			if !_cdc {
				_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _bbd)
				return _gg.ErrTypeError
			}
			_ggg, _abe := _gg.GetNumberAsFloat(_bbd.Params[1])
			if !_cdc {
				_ed.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u0046\u006c\u006f\u0061\u0074\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065d\u002e\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bbd, _abe)
				return _abe
			}
			_abe = _bcc.setFont(_dbf, _ggg)
			_bcc._efcbd = _aa.Is(_abe, _gg.ErrNotSupported)
			if _abe != nil && !_bcc._efcbd {
				return _abe
			}
		case "\u0054\u006d":
			if _cge, _dfe := _bcc.checkOp(_bbd, 6, true); !_cge {
				_ed.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u006d\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _dfe)
				return _dfe
			}
			_edc, _cfc := _gg.GetNumbersAsFloat(_bbd.Params)
			if _cfc != nil {
				_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _cfc)
				return _cfc
			}
			_bcc.setTextMatrix(_edc)
		case "\u0054\u0072":
			if _ecb, _cdcf := _bcc.checkOp(_bbd, 1, true); !_ecb {
				_ed.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0072\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _cdcf)
				return _cdcf
			}
			_cadc, _ega := _gg.GetIntVal(_bbd.Params[0])
			if !_ega {
				_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0072\u0020\u006f\u0070\u003d\u0025\u0073 \u0047e\u0074\u0049\u006e\u0074\u0056\u0061\u006c\u0020\u0066\u0061\u0069\u006c\u0065\u0064", _bbd)
				return _gg.ErrTypeError
			}
			_bcc.setTextRenderMode(_cadc)
		case "\u0054\u0073":
			if _gdc, _acdea := _bcc.checkOp(_bbd, 1, true); !_gdc {
				_ed.Log.Debug("\u0045\u0052R\u004f\u0052\u003a \u0054\u0073\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _acdea)
				return _acdea
			}
			_gae, _cba := _gg.GetNumberAsFloat(_bbd.Params[0])
			if _cba != nil {
				_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _cba)
				return _cba
			}
			_bcc.setTextRise(_gae)
		case "\u0054\u0077":
			if _afgb, _eefa := _bcc.checkOp(_bbd, 1, true); !_afgb {
				_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _eefa)
				return _eefa
			}
			_adc, _gaf := _gg.GetNumberAsFloat(_bbd.Params[0])
			if _gaf != nil {
				_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gaf)
				return _gaf
			}
			_bcc.setWordSpacing(_adc)
		case "\u0054\u007a":
			if _fec, _gcb := _bcc.checkOp(_bbd, 1, true); !_fec {
				_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _gcb)
				return _gcb
			}
			_ccdd, _ecfc := _gg.GetNumberAsFloat(_bbd.Params[0])
			if _ecfc != nil {
				_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _ecfc)
				return _ecfc
			}
			_bcc.setHorizScaling(_ccdd)
		case "\u0063\u006d":
			_eacc._cca = _gbeb.CTM
			if _eacc._cca.Singular() {
				_fgb := _cc.IdentityMatrix().Translate(_eacc._cca.Translation())
				_ed.Log.Debug("S\u0069n\u0067\u0075\u006c\u0061\u0072\u0020\u0063\u0074m\u003d\u0025\u0073\u2192%s", _eacc._cca, _fgb)
				_eacc._cca = _fgb
			}
			if _dgeg {
				_ed.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _eacc._cca)
			}
		case "\u006d":
			if len(_bbd.Params) != 2 {
				_ed.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006d\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _be)
				return nil
			}
			_gcg, _bdac := _gg.GetNumbersAsFloat(_bbd.Params)
			if _bdac != nil {
				return _bdac
			}
			_eacc.moveTo(_gcg[0], _gcg[1])
		case "\u006c":
			if len(_bbd.Params) != 2 {
				_ed.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0065\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u0020\u0060\u006c\u0060\u0020o\u0070\u0065r\u0061\u0074o\u0072\u003a\u0020\u0025\u0076\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074 m\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e", _be)
				return nil
			}
			_agad, _dee := _gg.GetNumbersAsFloat(_bbd.Params)
			if _dee != nil {
				return _dee
			}
			_eacc.lineTo(_agad[0], _agad[1])
		case "\u0063":
			if len(_bbd.Params) != 6 {
				return _be
			}
			_bfd, _aef := _gg.GetNumbersAsFloat(_bbd.Params)
			if _aef != nil {
				return _aef
			}
			_ed.Log.Debug("\u0043u\u0062\u0069\u0063\u0020b\u0065\u007a\u0069\u0065\u0072 \u0070a\u0072a\u006d\u0073\u003a\u0020\u0025\u002e\u0032f", _bfd)
			_eacc.cubicTo(_bfd[0], _bfd[1], _bfd[2], _bfd[3], _bfd[4], _bfd[5])
		case "\u0076", "\u0079":
			if len(_bbd.Params) != 4 {
				return _be
			}
			_bdb, _fcf := _gg.GetNumbersAsFloat(_bbd.Params)
			if _fcf != nil {
				return _fcf
			}
			_ed.Log.Debug("\u0043u\u0062\u0069\u0063\u0020b\u0065\u007a\u0069\u0065\u0072 \u0070a\u0072a\u006d\u0073\u003a\u0020\u0025\u002e\u0032f", _bdb)
			_eacc.quadraticTo(_bdb[0], _bdb[1], _bdb[2], _bdb[3])
		case "\u0068":
			_eacc.closePath()
		case "\u0072\u0065":
			if len(_bbd.Params) != 4 {
				return _be
			}
			_ebbfc, _aee := _gg.GetNumbersAsFloat(_bbd.Params)
			if _aee != nil {
				return _aee
			}
			_eacc.drawRectangle(_ebbfc[0], _ebbfc[1], _ebbfc[2], _ebbfc[3])
			_eacc.closePath()
		case "\u0053":
			_eacc.stroke(&_ccc._eff)
			_eacc.clearPath()
		case "\u0073":
			_eacc.closePath()
			_eacc.stroke(&_ccc._eff)
			_eacc.clearPath()
		case "\u0046":
			_eacc.fill(&_ccc._efbe)
			_eacc.clearPath()
		case "\u0066", "\u0066\u002a":
			_eacc.closePath()
			_eacc.fill(&_ccc._efbe)
			_eacc.clearPath()
		case "\u0042", "\u0042\u002a":
			_eacc.fill(&_ccc._efbe)
			_eacc.stroke(&_ccc._eff)
			_eacc.clearPath()
		case "\u0062", "\u0062\u002a":
			_eacc.closePath()
			_eacc.fill(&_ccc._efbe)
			_eacc.stroke(&_ccc._eff)
			_eacc.clearPath()
		case "\u006e":
			_eacc.clearPath()
		case "\u0044\u006f":
			if len(_bbd.Params) == 0 {
				_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0058\u004fbj\u0065c\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0070\u0065\u0072\u0061n\u0064\u0020\u0066\u006f\u0072\u0020\u0044\u006f\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072.\u0020\u0047\u006f\u0074\u0020\u0025\u002b\u0076\u002e", _bbd.Params)
				return _gg.ErrRangeError
			}
			_baa, _ebc := _gg.GetName(_bbd.Params[0])
			if !_ebc {
				_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0044\u006f\u0020\u006f\u0070e\u0072a\u0074\u006f\u0072\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006fp\u0065\u0072\u0061\u006e\u0064\u003a\u0020\u0025\u002b\u0076\u002e", _bbd.Params[0])
				return _gg.ErrTypeError
			}
			_, _aecd := _ffc.GetXObjectByName(*_baa)
			if _aecd != _gb.XObjectTypeForm {
				break
			}
			_bfcc, _ebc := _efe._db[_baa.String()]
			if !_ebc {
				_gbb, _ead := _ffc.GetXObjectFormByName(*_baa)
				if _ead != nil {
					_ed.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ead)
					return _ead
				}
				_eefc, _ead := _gbb.GetContentStream()
				if _ead != nil {
					_ed.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ead)
					return _ead
				}
				_fce := _gbb.Resources
				if _fce == nil {
					_fce = _ffc
				}
				_cbc := _gbeb.CTM
				if _bac, _cgd := _gg.GetArray(_gbb.Matrix); _cgd {
					_deed, _fdfb := _bac.GetAsFloat64Slice()
					if _fdfb != nil {
						return _fdfb
					}
					if len(_deed) != 6 {
						return _be
					}
					_bad := _cc.NewMatrix(_deed[0], _deed[1], _deed[2], _deed[3], _deed[4], _deed[5])
					_cbc = _gbeb.CTM.Mult(_bad)
				}
				_ecdb, _cef, _fgg, _ead := _efe.extractPageText(string(_eefc), _fce, _agb.Mult(_cbc), _gfef+1)
				if _ead != nil {
					_ed.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ead)
					return _ead
				}
				_bfcc = textResult{*_ecdb, _cef, _fgg}
				_efe._db[_baa.String()] = _bfcc
			}
			_eacc._cca = _gbeb.CTM
			if _dgeg {
				_ed.Log.Info("\u0063\u0074\u006d\u003d\u0025\u0073", _eacc._cca)
			}
			_ccc._bagf = append(_ccc._bagf, _bfcc._gga._bagf...)
			_ccc._eff = append(_ccc._eff, _bfcc._gga._eff...)
			_ccc._efbe = append(_ccc._efbe, _bfcc._gga._efbe...)
			_gdg._fdbd += _bfcc._bdef
			_gdg._eec += _bfcc._bee
		case "\u0072\u0067", "\u0067", "\u006b", "\u0063\u0073", "\u0073\u0063", "\u0073\u0063\u006e":
			_bcc._ccg.ColorspaceNonStroking = _gbeb.ColorspaceNonStroking
			_bcc._ccg.ColorNonStroking = _gbeb.ColorNonStroking
		case "\u0052\u0047", "\u0047", "\u004b", "\u0043\u0053", "\u0053\u0043", "\u0053\u0043\u004e":
			_bcc._ccg.ColorspaceStroking = _gbeb.ColorspaceStroking
			_bcc._ccg.ColorStroking = _gbeb.ColorStroking
		}
		return nil
	})
	_fdgb = _dcd.Process(_gfe)
	return _ccc, _gdg._fdbd, _gdg._eec, _fdgb
}
func (_cfbc *ruling) alignsSec(_eceab *ruling) bool {
	const _cafe = _dbgc + 1.0
	return _cfbc._daag-_cafe <= _eceab._fcff && _eceab._daag-_cafe <= _cfbc._fcff
}
func _cdcd(_acfg, _ecfd _gb.PdfRectangle) bool {
	return _ecfd.Llx <= _acfg.Urx && _acfg.Llx <= _ecfd.Urx
}

// String returns a description of `state`.
func (_bada *textState) String() string {
	_dbb := "\u005bN\u004f\u0054\u0020\u0053\u0045\u0054]"
	if _bada._bdg != nil {
		_dbb = _bada._bdg.BaseFont()
	}
	return _cfe.Sprintf("\u0074\u0063\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0077\u003d\u0025\u002e\u0032\u0066 \u0074f\u0073\u003d\u0025\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u003d\u0025\u0071", _bada._bfdf, _bada._bfa, _bada._agge, _dbb)
}

// String returns a description of `p`.
func (_gdbd *textPara) String() string {
	if _gdbd._befd {
		return _cfe.Sprintf("\u0025\u0036\u002e\u0032\u0066\u0020\u005b\u0045\u004d\u0050\u0054\u0059\u005d", _gdbd.PdfRectangle)
	}
	_acfcd := ""
	if _gdbd._bedfe != nil {
		_acfcd = _cfe.Sprintf("\u005b\u0025\u0064\u0078\u0025\u0064\u005d\u0020", _gdbd._bedfe._fccgee, _gdbd._bedfe._dege)
	}
	return _cfe.Sprintf("\u0025\u0036\u002e\u0032f \u0025\u0073\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u0020\u0025\u0071", _gdbd.PdfRectangle, _acfcd, len(_gdbd._abbe), _fabf(_gdbd.text(), 50))
}
func _dfba(_adff, _cgbc _gb.PdfRectangle) (_gb.PdfRectangle, bool) {
	if !_feee(_adff, _cgbc) {
		return _gb.PdfRectangle{}, false
	}
	return _gb.PdfRectangle{Llx: _df.Max(_adff.Llx, _cgbc.Llx), Urx: _df.Min(_adff.Urx, _cgbc.Urx), Lly: _df.Max(_adff.Lly, _cgbc.Lly), Ury: _df.Min(_adff.Ury, _cgbc.Ury)}, true
}
func _dga(_efee *textLine) float64 { return _efee._bfca[0].Llx }
func (_aebe compositeCell) String() string {
	_ebecf := ""
	if len(_aebe.paraList) > 0 {
		_ebecf = _fabf(_aebe.paraList.merge().text(), 50)
	}
	return _cfe.Sprintf("\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0064\u0020\u0070\u0061\u0072a\u0073\u0020\u0025\u0071", _aebe.PdfRectangle, len(_aebe.paraList), _ebecf)
}
func (_fef *stateStack) size() int { return len(*_fef) }
func _bedf(_dgdf float64) int {
	var _bfdc int
	if _dgdf >= 0 {
		_bfdc = int(_dgdf / _edagd)
	} else {
		_bfdc = int(_dgdf/_edagd) - 1
	}
	return _bfdc
}

// ExtractFonts returns all font information from the page extractor, including
// font name, font type, the raw data of the embedded font file (if embedded), font descriptor and more.
//
// The argument `previousPageFonts` is used when trying to build a complete font catalog for multiple pages or the entire document.
// The entries from `previousPageFonts` are added to the returned result unless already included in the page, i.e. no duplicate entries.
//
// NOTE: If previousPageFonts is nil, all fonts from the page will be returned. Use it when building up a full list of fonts for a document or page range.
func (_ge *Extractor) ExtractFonts(previousPageFonts *PageFonts) (*PageFonts, error) {
	_beb := PageFonts{}
	_gfb := _beb.extractPageResourcesToFont(_ge._bd)
	if _gfb != nil {
		return nil, _gfb
	}
	if previousPageFonts != nil {
		for _, _fd := range previousPageFonts.Fonts {
			if !_eaf(_beb.Fonts, _fd.FontName) {
				_beb.Fonts = append(_beb.Fonts, _fd)
			}
		}
	}
	return &PageFonts{Fonts: _beb.Fonts}, nil
}

type ruling struct {
	_cffe rulingKind
	_bbge markKind
	_eb.Color
	_abfg float64
	_daag float64
	_fcff float64
	_fdde float64
}

func (_bgbaa *textPara) depth() float64 {
	if _bgbaa._befd {
		return -1.0
	}
	if len(_bgbaa._abbe) > 0 {
		return _bgbaa._abbe[0]._bdbg
	}
	return _bgbaa._bedfe.depth()
}
func _fggcc(_cfbd *list) []*list {
	var _acgc []*list
	for _, _dedb := range _cfbd._gbcab {
		switch _dedb._feec {
		case "\u004c\u0049":
			_ccfgg := _accb(_dedb)
			_ddff := _fggcc(_dedb)
			_eecg := _gbcag(_ccfgg, "\u0062\u0075\u006c\u006c\u0065\u0074", _ddff)
			_afdb := _adag(_ccfgg, "")
			_eecg._ebabg = _afdb
			_acgc = append(_acgc, _eecg)
		case "\u004c\u0042\u006fd\u0079":
			return _fggcc(_dedb)
		case "\u004c":
			_bbfea := _fggcc(_dedb)
			_acgc = append(_acgc, _bbfea...)
			return _acgc
		}
	}
	return _acgc
}
func (_bdce *wordBag) maxDepth() float64 { return _bdce._gebga - _bdce.Lly }
func (_ffee *textObject) setHorizScaling(_eca float64) {
	if _ffee == nil {
		return
	}
	_ffee._cbag._fceg = _eca
}

// RenderMode specifies the text rendering mode (Tmode), which determines whether showing text shall cause
// glyph outlines to be  stroked, filled, used as a clipping boundary, or some combination of the three.
// Stroking, filling, and clipping shall have the same effects for a text object as they do for a path object
// (see 8.5.3, "Path-Painting Operators" and 8.5.4, "Clipping Path Operators").
type RenderMode int

func (_fgd *PageText) getParagraphs() paraList {
	var _ebee rulingList
	if _gfaf {
		_bbcg := _dbff(_fgd._eff)
		_ebee = append(_ebee, _bbcg...)
	}
	if _bdacb {
		_edag := _bdcec(_fgd._efbe)
		_ebee = append(_ebee, _edag...)
	}
	_ebee, _bebb := _ebee.toTilings()
	var _fcdf paraList
	_ddf := len(_fgd._bagf)
	for _ddcc := 0; _ddcc < 360 && _ddf > 0; _ddcc += 90 {
		_geee := make([]*textMark, 0, len(_fgd._bagf)-_ddf)
		for _, _deg := range _fgd._bagf {
			if _deg._cggbc == _ddcc {
				_geee = append(_geee, _deg)
			}
		}
		if len(_geee) > 0 {
			_dcee := _acgef(_geee, _fgd._ebae, _ebee, _bebb, _fgd._ege._ebff)
			_fcdf = append(_fcdf, _dcee...)
			_ddf -= len(_geee)
		}
	}
	return _fcdf
}
func (_defg *wordBag) arrangeText() *textPara {
	_defg.sort()
	if _gdggd {
		_defg.removeDuplicates()
	}
	var _dgga []*textLine
	for _, _gecac := range _defg.depthIndexes() {
		for !_defg.empty(_gecac) {
			_fgfaa := _defg.firstReadingIndex(_gecac)
			_fgfd := _defg.firstWord(_fgfaa)
			_dfgf := _dabb(_defg, _fgfaa)
			_agbb := _fgfd._fgdbd
			_abedb := _fgfd._cbfcee - _dgfg*_agbb
			_facc := _fgfd._cbfcee + _dgfg*_agbb
			_abaac := _dedd * _agbb
			_dead := _ccfe * _agbb
		_fdeb:
			for {
				var _ffcb *textWord
				_cbcg := 0
				for _, _bcbcd := range _defg.depthBand(_abedb, _facc) {
					_dbaf := _defg.highestWord(_bcbcd, _abedb, _facc)
					if _dbaf == nil {
						continue
					}
					_fedcg := _daefa(_dbaf, _dfgf._bfca[len(_dfgf._bfca)-1])
					if _fedcg < -_dead {
						break _fdeb
					}
					if _fedcg > _abaac {
						continue
					}
					if _ffcb != nil && _fdd(_dbaf, _ffcb) >= 0 {
						continue
					}
					_ffcb = _dbaf
					_cbcg = _bcbcd
				}
				if _ffcb == nil {
					break
				}
				_dfgf.pullWord(_defg, _ffcb, _cbcg)
			}
			_dfgf.markWordBoundaries()
			_dgga = append(_dgga, _dfgf)
		}
	}
	if len(_dgga) == 0 {
		return nil
	}
	_d.Slice(_dgga, func(_fagfe, _bdcb int) bool { return _ecdf(_dgga[_fagfe], _dgga[_bdcb]) < 0 })
	_gdgb := _daae(_defg.PdfRectangle, _dgga)
	if _gcga {
		_ed.Log.Info("\u0061\u0072\u0072an\u0067\u0065\u0054\u0065\u0078\u0074\u0020\u0021\u0021\u0021\u0020\u0070\u0061\u0072\u0061\u003d\u0025\u0073", _gdgb.String())
		if _abgd {
			for _gebd, _bdag := range _gdgb._abbe {
				_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _gebd, _bdag.String())
				if _befe {
					for _bfdfa, _cbgc := range _bdag._bfca {
						_cfe.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _bfdfa, _cbgc.String())
						for _fdbbe, _fcec := range _cbgc._fdgbc {
							_cfe.Printf("\u00251\u0032\u0064\u003a\u0020\u0025\u0073\n", _fdbbe, _fcec.String())
						}
					}
				}
			}
		}
	}
	return _gdgb
}

type lineRuling struct {
	_ecee rulingKind
	_fbcd markKind
	_eb.Color
	_efffg, _agae _cc.Point
}
type compositeCell struct {
	_gb.PdfRectangle
	paraList
}

func (_geg *imageExtractContext) extractXObjectImage(_gee *_gg.PdfObjectName, _ecfa _aad.GraphicsState, _aae *_gb.PdfPageResources) error {
	_fdf, _ := _aae.GetXObjectByName(*_gee)
	if _fdf == nil {
		return nil
	}
	_gdb, _dcg := _geg._eae[_fdf]
	if !_dcg {
		_agf, _eea := _aae.GetXObjectImageByName(*_gee)
		if _eea != nil {
			return _eea
		}
		if _agf == nil {
			return nil
		}
		_def, _eea := _agf.ToImage()
		if _eea != nil {
			return _eea
		}
		var _gea _gf.Image
		if _agf.Mask != nil {
			if _gea, _eea = _ccbcb(_agf.Mask, _eb.Opaque); _eea != nil {
				_ed.Log.Debug("\u0057\u0041\u0052\u004e\u003a \u0063\u006f\u0075\u006c\u0064 \u006eo\u0074\u0020\u0067\u0065\u0074\u0020\u0065\u0078\u0070\u006c\u0069\u0063\u0069\u0074\u0020\u0069\u006d\u0061\u0067e\u0020\u006d\u0061\u0073\u006b\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063o\u0072\u0072\u0065\u0063\u0074\u002e")
			}
		} else if _agf.SMask != nil {
			_gea, _eea = _agbca(_agf.SMask, _eb.Opaque)
			if _eea != nil {
				_ed.Log.Debug("W\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0066\u0074\u0020\u0069\u006da\u0067e\u0020\u006d\u0061\u0073k\u002e\u0020O\u0075\u0074\u0070\u0075\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u002e")
			}
		}
		if _gea != nil {
			_gcd, _bfc := _def.ToGoImage()
			if _bfc != nil {
				return _bfc
			}
			_gcd = _cadba(_gcd, _gea)
			switch _agf.ColorSpace.String() {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079", "\u0049n\u0064\u0065\u0078\u0065\u0064":
				_def, _bfc = _gb.ImageHandling.NewGrayImageFromGoImage(_gcd)
				if _bfc != nil {
					return _bfc
				}
			default:
				_def, _bfc = _gb.ImageHandling.NewImageFromGoImage(_gcd)
				if _bfc != nil {
					return _bfc
				}
			}
		}
		_gdb = &cachedImage{_gde: _def, _ecg: _agf.ColorSpace}
		_geg._eae[_fdf] = _gdb
	}
	_feg := _gdb._gde
	_fg := _gdb._ecg
	_fba, _bc := _fg.ImageToRGB(*_feg)
	if _bc != nil {
		return _bc
	}
	_ed.Log.Debug("@\u0044\u006f\u0020\u0043\u0054\u004d\u003a\u0020\u0025\u0073", _ecfa.CTM.String())
	_bcg := ImageMark{Image: &_fba, Width: _ecfa.CTM.ScalingFactorX(), Height: _ecfa.CTM.ScalingFactorY(), Angle: _ecfa.CTM.Angle()}
	_bcg.X, _bcg.Y = _ecfa.CTM.Translation()
	_geg._cdb = append(_geg._cdb, _bcg)
	_geg._ce++
	return nil
}
func _gecd(_feae *textLine, _fbaa []*textLine, _bcff []float64) float64 {
	var _agafc float64 = -1
	for _, _gagd := range _fbaa {
		if _gagd._bdbg > _feae._bdbg {
			if _df.Round(_gagd.Llx) >= _df.Round(_feae.Llx) {
				_agafc = _gagd._bdbg
			} else {
				break
			}
		}
	}
	return _agafc
}
func (_fgfdb rulingList) log(_agdbf string) {
	if !_geaeg {
		return
	}
	_ed.Log.Info("\u0023\u0023\u0023\u0020\u0025\u0031\u0030\u0073\u003a\u0020\u0076\u0065c\u0073\u003d\u0025\u0073", _agdbf, _fgfdb.String())
	for _cface, _fgae := range _fgfdb {
		_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _cface, _fgae.String())
	}
}
func (_bfggd *textTable) getComposite(_dfbd, _fbac int) (paraList, _gb.PdfRectangle) {
	_fecd, _aecg := _bfggd._faeg[_aafb(_dfbd, _fbac)]
	if _geff {
		_cfe.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0067\u0065\u0074\u0043\u006f\u006d\u0070o\u0073i\u0074\u0065\u0028\u0025\u0064\u002c\u0025\u0064\u0029\u002d\u003e\u0025\u0073\u000a", _dfbd, _fbac, _fecd.String())
	}
	if !_aecg {
		return nil, _gb.PdfRectangle{}
	}
	return _fecd.parasBBox()
}
func (_fgf *textObject) showText(_bbdb _gg.PdfObject, _fcea []byte, _abdg int) error {
	return _fgf.renderText(_bbdb, _fcea, _abdg)
}

var _aaeg *_ca.Regexp = _ca.MustCompile(_fdfgd + "\u007c" + _gabd)

func _bafe(_gbda, _cbdd bounded) float64 { return _ggbc(_gbda) - _ggbc(_cbdd) }
func (_ecba *textObject) getFontDict(_baag string) (_gbca _gg.PdfObject, _cbbg error) {
	_cffd := _ecba._edacg
	if _cffd == nil {
		_ed.Log.Debug("g\u0065\u0074\u0046\u006f\u006e\u0074D\u0069\u0063\u0074\u002e\u0020\u004eo\u0020\u0072\u0065\u0073\u006f\u0075\u0072c\u0065\u0073\u002e\u0020\u006e\u0061\u006d\u0065\u003d\u0025#\u0071", _baag)
		return nil, nil
	}
	_gbca, _agea := _cffd.GetFontByName(_gg.PdfObjectName(_baag))
	if !_agea {
		_ed.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u0067\u0065t\u0046\u006f\u006et\u0044\u0069\u0063\u0074\u003a\u0020\u0046\u006f\u006et \u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020\u006e\u0061m\u0065\u003d%\u0023\u0071", _baag)
		return nil, _c.New("f\u006f\u006e\u0074\u0020no\u0074 \u0069\u006e\u0020\u0072\u0065s\u006f\u0075\u0072\u0063\u0065\u0073")
	}
	return _gbca, nil
}
func (_acg *textObject) setTextRise(_fcd float64) {
	if _acg == nil {
		return
	}
	_acg._cbag._gce = _fcd
}
func (_gebg *TextMarkArray) exists(_eedc TextMark) bool {
	for _, _ddfb := range _gebg.Elements() {
		if _g.DeepEqual(_eedc.DirectObject, _ddfb.DirectObject) && _g.DeepEqual(_eedc.BBox, _ddfb.BBox) && _ddfb.Text == _eedc.Text {
			return true
		}
	}
	return false
}
func _gbcbg(_adfe map[int]intSet) []int {
	_gecaf := make([]int, 0, len(_adfe))
	for _cbcce := range _adfe {
		_gecaf = append(_gecaf, _cbcce)
	}
	_d.Ints(_gecaf)
	return _gecaf
}
func (_abdf *shapesState) closePath() {
	if _abdf._dfff {
		_abdf._babc = append(_abdf._babc, _dcab(_abdf._bcb))
		_abdf._dfff = false
	} else if len(_abdf._babc) == 0 {
		if _dgeg {
			_ed.Log.Debug("\u0063\u006c\u006f\u0073eP\u0061\u0074\u0068\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f\u0020\u0070\u0061t\u0068")
		}
		_abdf._dfff = false
		return
	}
	_abdf._babc[len(_abdf._babc)-1].close()
	if _dgeg {
		_ed.Log.Info("\u0063\u006c\u006f\u0073\u0065\u0050\u0061\u0074\u0068\u003a\u0020\u0025\u0073", _abdf)
	}
}

// ApplyArea processes the page text only within the specified area `bbox`.
// Each time ApplyArea is called, it updates the result set in `pt`.
// Can be called multiple times in a row with different bounding boxes.
func (_bfcff *PageText) ApplyArea(bbox _gb.PdfRectangle) {
	_ceda := make([]*textMark, 0, len(_bfcff._bagf))
	for _, _dfdcb := range _bfcff._bagf {
		if _feee(_dfdcb.bbox(), bbox) {
			_ceda = append(_ceda, _dfdcb)
		}
	}
	var _abfd paraList
	_fbg := len(_ceda)
	for _fac := 0; _fac < 360 && _fbg > 0; _fac += 90 {
		_aag := make([]*textMark, 0, len(_ceda)-_fbg)
		for _, _gffe := range _ceda {
			if _gffe._cggbc == _fac {
				_aag = append(_aag, _gffe)
			}
		}
		if len(_aag) > 0 {
			_deb := _acgef(_aag, _bfcff._ebae, nil, nil, _bfcff._ege._ebff)
			_abfd = append(_abfd, _deb...)
			_fbg -= len(_aag)
		}
	}
	_ggb := new(_cf.Buffer)
	_abfd.writeText(_ggb)
	_bfcff._ccbd = _ggb.String()
	_bfcff._afc = _abfd.toTextMarks()
	_bfcff._ded = _abfd.tables()
}
func (_bbab *textTable) put(_egbg, _gbba int, _ccfa *textPara) {
	_bbab._adda[_aafb(_egbg, _gbba)] = _ccfa
}
func (_fdddd rulingList) vertsHorzs() (rulingList, rulingList) {
	var _gdbg, _eadb rulingList
	for _, _cedgb := range _fdddd {
		switch _cedgb._cffe {
		case _gfafc:
			_gdbg = append(_gdbg, _cedgb)
		case _bdfd:
			_eadb = append(_eadb, _cedgb)
		}
	}
	return _gdbg, _eadb
}
func (_ebdf *textObject) newTextMark(_gcge string, _aeff _cc.Matrix, _ebea _cc.Point, _eagg float64, _dbef *_gb.PdfFont, _dbbc float64, _dbbd, _cbce _eb.Color, _bcbg _gg.PdfObject, _badeb []string, _gfafb int, _edacd int) (textMark, bool) {
	_fgbd := _aeff.Angle()
	_gabg := _edeeg(_fgbd, _ecce)
	var _fafcg float64
	if _gabg%180 != 90 {
		_fafcg = _aeff.ScalingFactorY()
	} else {
		_fafcg = _aeff.ScalingFactorX()
	}
	_edagcg := _bba(_aeff)
	_abbb := _gb.PdfRectangle{Llx: _edagcg.X, Lly: _edagcg.Y, Urx: _ebea.X, Ury: _ebea.Y}
	switch _gabg % 360 {
	case 90:
		_abbb.Urx -= _fafcg
	case 180:
		_abbb.Ury -= _fafcg
	case 270:
		_abbb.Urx += _fafcg
	case 0:
		_abbb.Ury += _fafcg
	default:
		_gabg = 0
		_abbb.Ury += _fafcg
	}
	if _abbb.Llx > _abbb.Urx {
		_abbb.Llx, _abbb.Urx = _abbb.Urx, _abbb.Llx
	}
	if _abbb.Lly > _abbb.Ury {
		_abbb.Lly, _abbb.Ury = _abbb.Ury, _abbb.Lly
	}
	_abed := true
	if _ebdf._cgf._eda.Width() > 0 {
		_ffd, _abec := _dfba(_abbb, _ebdf._cgf._eda)
		if !_abec {
			_abed = false
			_ed.Log.Debug("\u0054\u0065\u0078\u0074\u0020m\u0061\u0072\u006b\u0020\u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0070a\u0067\u0065\u002e\u0020\u0062\u0062\u006f\u0078\u003d\u0025\u0067\u0020\u006d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u0067\u0020\u0074\u0065\u0078\u0074\u003d\u0025q", _abbb, _ebdf._cgf._eda, _gcge)
		}
		_abbb = _ffd
	}
	_dccg := _abbb
	_baba := _ebdf._cgf._eda
	switch _gabg % 360 {
	case 90:
		_baba.Urx, _baba.Ury = _baba.Ury, _baba.Urx
		_dccg = _gb.PdfRectangle{Llx: _baba.Urx - _abbb.Ury, Urx: _baba.Urx - _abbb.Lly, Lly: _abbb.Llx, Ury: _abbb.Urx}
	case 180:
		_dccg = _gb.PdfRectangle{Llx: _baba.Urx - _abbb.Llx, Urx: _baba.Urx - _abbb.Urx, Lly: _baba.Ury - _abbb.Lly, Ury: _baba.Ury - _abbb.Ury}
	case 270:
		_baba.Urx, _baba.Ury = _baba.Ury, _baba.Urx
		_dccg = _gb.PdfRectangle{Llx: _abbb.Ury, Urx: _abbb.Lly, Lly: _baba.Ury - _abbb.Llx, Ury: _baba.Ury - _abbb.Urx}
	}
	if _dccg.Llx > _dccg.Urx {
		_dccg.Llx, _dccg.Urx = _dccg.Urx, _dccg.Llx
	}
	if _dccg.Lly > _dccg.Ury {
		_dccg.Lly, _dccg.Ury = _dccg.Ury, _dccg.Lly
	}
	_afbc := textMark{_ceca: _gcge, PdfRectangle: _dccg, _aagcg: _abbb, _gcfc: _dbef, _gded: _fafcg, _ggaf: _dbbc, _dddf: _aeff, _fdda: _ebea, _cggbc: _gabg, _egea: _dbbd, _cdfbb: _cbce, _dgebc: _bcbg, _bfdcd: _badeb, Th: _ebdf._cbag._fceg, Tw: _ebdf._cbag._bfa, _afae: _edacd, _gegcc: _gfafb}
	if _gbbcc {
		_ed.Log.Info("n\u0065\u0077\u0054\u0065\u0078\u0074M\u0061\u0072\u006b\u003a\u0020\u0073t\u0061\u0072\u0074\u003d\u0025\u002e\u0032f\u0020\u0065\u006e\u0064\u003d\u0025\u002e\u0032\u0066\u0020%\u0073", _edagcg, _ebea, _afbc.String())
	}
	return _afbc, _abed
}
func (_cbfe *shapesState) devicePoint(_fca, _fgeg float64) _cc.Point {
	_ecfb := _cbfe._begf.Mult(_cbfe._cca)
	_fca, _fgeg = _ecfb.Transform(_fca, _fgeg)
	return _cc.NewPoint(_fca, _fgeg)
}
func (_ggee rulingList) splitSec() []rulingList {
	_d.Slice(_ggee, func(_cgfgd, _fceeg int) bool {
		_dfaa, _ggdee := _ggee[_cgfgd], _ggee[_fceeg]
		if _dfaa._daag != _ggdee._daag {
			return _dfaa._daag < _ggdee._daag
		}
		return _dfaa._fcff < _ggdee._fcff
	})
	_agac := make(map[*ruling]struct{}, len(_ggee))
	_abdgg := func(_fdga *ruling) rulingList {
		_fbfe := rulingList{_fdga}
		_agac[_fdga] = struct{}{}
		for _, _ccebc := range _ggee {
			if _, _agefg := _agac[_ccebc]; _agefg {
				continue
			}
			for _, _egfd := range _fbfe {
				if _ccebc.alignsSec(_egfd) {
					_fbfe = append(_fbfe, _ccebc)
					_agac[_ccebc] = struct{}{}
					break
				}
			}
		}
		return _fbfe
	}
	_ecaf := []rulingList{_abdgg(_ggee[0])}
	for _, _fcafc := range _ggee[1:] {
		if _, _caefc := _agac[_fcafc]; _caefc {
			continue
		}
		_ecaf = append(_ecaf, _abdgg(_fcafc))
	}
	return _ecaf
}

type list struct {
	_ggbee []*textLine
	_feec  string
	_gbcab []*list
	_ebabg string
}

func (_cbb *textObject) getFillColor() _eb.Color {
	return _fbfab(_cbb._ccg.ColorspaceNonStroking, _cbb._ccg.ColorNonStroking)
}
func (_ececce paraList) findGridTables(_gedfb []gridTiling) []*textTable {
	if _geff {
		_ed.Log.Info("\u0066i\u006e\u0064\u0047\u0072\u0069\u0064\u0054\u0061\u0062\u006c\u0065s\u003a\u0020\u0025\u0064\u0020\u0070\u0061\u0072\u0061\u0073", len(_ececce))
		for _feaf, _eddd := range _ececce {
			_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _feaf, _eddd)
		}
	}
	var _eddg []*textTable
	for _ffacg, _bgdgf := range _gedfb {
		_dgede, _fbfbc := _ececce.findTableGrid(_bgdgf)
		if _dgede != nil {
			_dgede.log(_cfe.Sprintf("\u0066\u0069\u006e\u0064Ta\u0062\u006c\u0065\u0057\u0069\u0074\u0068\u0047\u0072\u0069\u0064\u0073\u003a\u0020%\u0064", _ffacg))
			_eddg = append(_eddg, _dgede)
			_dgede.markCells()
		}
		for _edgcd := range _fbfbc {
			_edgcd._ggdd = true
		}
	}
	if _geff {
		_ed.Log.Info("\u0066i\u006e\u0064\u0047\u0072i\u0064\u0054\u0061\u0062\u006ce\u0073:\u0020%\u0064\u0020\u0074\u0061\u0062\u006c\u0065s", len(_eddg))
	}
	return _eddg
}
func _adcb(_fdfda float64) bool { return _df.Abs(_fdfda) < _dbgc }
func _fgdf(_ggef string) (string, bool) {
	_gfdgg := []rune(_ggef)
	if len(_gfdgg) != 1 {
		return "", false
	}
	_dbbfg, _ccec := _edfcef[_gfdgg[0]]
	return _dbbfg, _ccec
}
func _cbcc(_abae _gb.PdfRectangle) *ruling {
	return &ruling{_cffe: _bdfd, _abfg: _abae.Ury, _daag: _abae.Llx, _fcff: _abae.Urx}
}
func _ggff(_aged []*wordBag) []*wordBag {
	if len(_aged) <= 1 {
		return _aged
	}
	if _gcga {
		_ed.Log.Info("\u006d\u0065\u0072\u0067\u0065\u0057\u006f\u0072\u0064B\u0061\u0067\u0073\u003a")
	}
	_d.Slice(_aged, func(_cffgg, _ebag int) bool {
		_cffge, _dddg := _aged[_cffgg], _aged[_ebag]
		_egee := _cffge.Width() * _cffge.Height()
		_aggb := _dddg.Width() * _dddg.Height()
		if _egee != _aggb {
			return _egee > _aggb
		}
		if _cffge.Height() != _dddg.Height() {
			return _cffge.Height() > _dddg.Height()
		}
		return _cffgg < _ebag
	})
	var _gecef []*wordBag
	_ecc := make(intSet)
	for _bgff := 0; _bgff < len(_aged); _bgff++ {
		if _ecc.has(_bgff) {
			continue
		}
		_cdg := _aged[_bgff]
		for _fgfb := _bgff + 1; _fgfb < len(_aged); _fgfb++ {
			if _ecc.has(_bgff) {
				continue
			}
			_dcff := _aged[_fgfb]
			_fdbb := _cdg.PdfRectangle
			_fdbb.Llx -= _cdg._bcef
			if _egcd(_fdbb, _dcff.PdfRectangle) {
				_cdg.absorb(_dcff)
				_ecc.add(_fgfb)
			}
		}
		_gecef = append(_gecef, _cdg)
	}
	if len(_aged) != len(_gecef)+len(_ecc) {
		_ed.Log.Error("\u006d\u0065\u0072ge\u0057\u006f\u0072\u0064\u0042\u0061\u0067\u0073\u003a \u0025d\u2192%\u0064 \u0061\u0062\u0073\u006f\u0072\u0062\u0065\u0064\u003d\u0025\u0064", len(_aged), len(_gecef), len(_ecc))
	}
	return _gecef
}
func (_gdga rulingList) sort() { _d.Slice(_gdga, _gdga.comp) }
func (_caef *stateStack) top() *textState {
	if _caef.empty() {
		return nil
	}
	return (*_caef)[_caef.size()-1]
}
func _ccbcb(_fdbc _gg.PdfObject, _abcf _eb.Color) (_gf.Image, error) {
	_fccb, _cdfda := _gg.GetStream(_fdbc)
	if !_cdfda {
		return nil, nil
	}
	_abdca, _ccac := _gb.NewXObjectImageFromStream(_fccb)
	if _ccac != nil {
		return nil, _ccac
	}
	_feead, _ccac := _abdca.ToImage()
	if _ccac != nil {
		return nil, _ccac
	}
	return _fbaff(_feead, _abcf), nil
}

// Text returns the text content of the `bulletLists`.
func (_ceeb *lists) Text() string {
	_ccfc := &_ag.Builder{}
	for _, _gedbg := range *_ceeb {
		_bgde := _gedbg.Text()
		_ccfc.WriteString(_bgde)
	}
	return _ccfc.String()
}
func (_gfag *structElement) parseStructElement(_eebae _gg.PdfObject) {
	_bace, _gdef := _gg.GetDict(_eebae)
	if !_gdef {
		_ed.Log.Debug("\u0070\u0061\u0072\u0073\u0065\u0053\u0074\u0072u\u0063\u0074\u0045le\u006d\u0065\u006e\u0074\u003a\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u002e")
		return
	}
	_ecfbf := _bace.Get("\u0053")
	_ebged := _bace.Get("\u0050\u0067")
	_bgab := ""
	if _ecfbf != nil {
		_bgab = _ecfbf.String()
	}
	_agda := _bace.Get("\u004b")
	_gfag._gacb = _bgab
	_gfag._gfeff = _ebged
	switch _fdcad := _agda.(type) {
	case *_gg.PdfObjectInteger:
		_gfag._gacb = _bgab
		_gfag._adfb = int64(*_fdcad)
		_gfag._gfeff = _ebged
	case *_gg.PdfObjectReference:
		_egec := *_gg.MakeArray(_fdcad)
		var _dfcf int64 = -1
		_gfag._adfb = _dfcf
		if _egec.Len() == 1 {
			_cgfd := _egec.Elements()[0]
			_ecdaca, _eded := _cgfd.(*_gg.PdfObjectInteger)
			if _eded {
				_dfcf = int64(*_ecdaca)
				_gfag._adfb = _dfcf
				_gfag._gacb = _bgab
				_gfag._gfeff = _ebged
				return
			}
		}
		_bfgd := []structElement{}
		for _, _ccfd := range _egec.Elements() {
			_dedae, _ddae := _ccfd.(*_gg.PdfObjectInteger)
			if _ddae {
				_dfcf = int64(*_dedae)
				_gfag._adfb = _dfcf
				_gfag._gacb = _bgab
			} else {
				_faec := &structElement{}
				_faec.parseStructElement(_ccfd)
				_bfgd = append(_bfgd, *_faec)
			}
			_dfcf = -1
		}
		_gfag._ddcfe = _bfgd
	case *_gg.PdfObjectArray:
		_bbbd := _agda.(*_gg.PdfObjectArray)
		var _adfg int64 = -1
		_gfag._adfb = _adfg
		if _bbbd.Len() == 1 {
			_acfef := _bbbd.Elements()[0]
			_afed, _eefad := _acfef.(*_gg.PdfObjectInteger)
			if _eefad {
				_adfg = int64(*_afed)
				_gfag._adfb = _adfg
				_gfag._gacb = _bgab
				_gfag._gfeff = _ebged
				return
			}
		}
		_gced := []structElement{}
		for _, _eeac := range _bbbd.Elements() {
			_bfde, _agbg := _eeac.(*_gg.PdfObjectInteger)
			if _agbg {
				_adfg = int64(*_bfde)
				_gfag._adfb = _adfg
				_gfag._gacb = _bgab
				_gfag._gfeff = _ebged
			} else {
				_dgbb := &structElement{}
				_dgbb.parseStructElement(_eeac)
				_gced = append(_gced, *_dgbb)
			}
			_adfg = -1
		}
		_gfag._ddcfe = _gced
	}
}
func _fdedf(_egead, _dbbab _cc.Point) bool {
	_fecf := _df.Abs(_egead.X - _dbbab.X)
	_egbdb := _df.Abs(_egead.Y - _dbbab.Y)
	return _dffe(_fecf, _egbdb)
}

type textState struct {
	_bfdf  float64
	_bfa   float64
	_fceg  float64
	_fea   float64
	_agge  float64
	_eacgf RenderMode
	_gce   float64
	_bdg   *_gb.PdfFont
	_afd   _gb.PdfRectangle
	_fdbd  int
	_eec   int
}

const _baee = 10

func (_cgga rectRuling) asRuling() (*ruling, bool) {
	_dgcdd := ruling{_cffe: _cgga._gcaad, Color: _cgga.Color, _bbge: _cbcd}
	switch _cgga._gcaad {
	case _gfafc:
		_dgcdd._abfg = 0.5 * (_cgga.Llx + _cgga.Urx)
		_dgcdd._daag = _cgga.Lly
		_dgcdd._fcff = _cgga.Ury
		_dfed, _gddef := _cgga.checkWidth(_cgga.Llx, _cgga.Urx)
		if !_gddef {
			if _afbec {
				_ed.Log.Error("\u0072\u0065\u0063\u0074\u0052\u0075l\u0069\u006e\u0067\u002e\u0061\u0073\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0072\u0075\u006c\u0069\u006e\u0067V\u0065\u0072\u0074\u0020\u0021\u0063\u0068\u0065\u0063\u006b\u0057\u0069\u0064\u0074h\u0020v\u003d\u0025\u002b\u0076", _cgga)
			}
			return nil, false
		}
		_dgcdd._fdde = _dfed
	case _bdfd:
		_dgcdd._abfg = 0.5 * (_cgga.Lly + _cgga.Ury)
		_dgcdd._daag = _cgga.Llx
		_dgcdd._fcff = _cgga.Urx
		_bdbfb, _dbfcf := _cgga.checkWidth(_cgga.Lly, _cgga.Ury)
		if !_dbfcf {
			if _afbec {
				_ed.Log.Error("\u0072\u0065\u0063\u0074\u0052\u0075l\u0069\u006e\u0067\u002e\u0061\u0073\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0072\u0075\u006c\u0069\u006e\u0067H\u006f\u0072\u007a\u0020\u0021\u0063\u0068\u0065\u0063\u006b\u0057\u0069\u0064\u0074h\u0020v\u003d\u0025\u002b\u0076", _cgga)
			}
			return nil, false
		}
		_dgcdd._fdde = _bdbfb
	default:
		_ed.Log.Error("\u0062\u0061\u0064\u0020pr\u0069\u006d\u0061\u0072\u0079\u0020\u006b\u0069\u006e\u0064\u003d\u0025\u0064", _cgga._gcaad)
		return nil, false
	}
	return &_dgcdd, true
}
func (_adad *textLine) text() string {
	var _dcgd []string
	for _, _bead := range _adad._bfca {
		if _bead._cacca {
			_dcgd = append(_dcgd, "\u0020")
		}
		_dcgd = append(_dcgd, _bead._bcfea)
	}
	return _ag.Join(_dcgd, "")
}
func _dcab(_eece _cc.Point) *subpath { return &subpath{_ccgf: []_cc.Point{_eece}} }
func (_gbee *textObject) setTextRenderMode(_gfed int) {
	if _gbee == nil {
		return
	}
	_gbee._cbag._eacgf = RenderMode(_gfed)
}

type rulingKind int

func _dffe(_bcdga, _ecceb float64) bool { return _bcdga/_df.Max(_cccd, _ecceb) < _gdde }

type textMark struct {
	_gb.PdfRectangle
	_cggbc int
	_ceca  string
	_fbgee string
	_gcfc  *_gb.PdfFont
	_gded  float64
	_ggaf  float64
	_dddf  _cc.Matrix
	_fdda  _cc.Point
	_aagcg _gb.PdfRectangle
	_egea  _eb.Color
	_cdfbb _eb.Color
	_dgebc _gg.PdfObject
	_bfdcd []string
	Tw     float64
	Th     float64
	_afae  int
	_gegcc int
}

func (_beae *shapesState) moveTo(_bdgf, _aafe float64) {
	_beae._dfff = true
	_beae._bcb = _beae.devicePoint(_bdgf, _aafe)
	if _dgeg {
		_ed.Log.Info("\u006d\u006fv\u0065\u0054\u006f\u003a\u0020\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0020\u0064\u0065\u0076\u0069\u0063\u0065\u003d%.\u0032\u0066", _bdgf, _aafe, _beae._bcb)
	}
}

const (
	_eefaf rulingKind = iota
	_bdfd
	_gfafc
)

func _gede(_fgecb, _ddef bounded) float64 {
	_gdgf := _fdd(_fgecb, _ddef)
	if !_cdgd(_gdgf) {
		return _gdgf
	}
	return _bafe(_fgecb, _ddef)
}

// NewWithOptions an Extractor instance for extracting content from the input PDF page with options.
func NewWithOptions(page *_gb.PdfPage, options *Options) (*Extractor, error) {
	_eeb, _ff := page.GetAllContentStreams()
	if _ff != nil {
		return nil, _ff
	}
	_ec, _ef := page.GetStructTreeRoot()
	if !_ef {
		_ed.Log.Info("T\u0068\u0065\u0020\u0070\u0064\u0066\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0074\u0061\u0067g\u0065d\u002e\u0020\u0053\u0074r\u0075\u0063t\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020\u0065\u0078\u0069\u0073\u0074\u002e")
	}
	_fe := page.GetContainingPdfObject()
	_cfg, _ff := page.GetMediaBox()
	if _ff != nil {
		return nil, _cfe.Errorf("\u0065\u0078\u0074r\u0061\u0063\u0074\u006fr\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u002e\u0020\u0025\u0076", _ff)
	}
	_bf := &Extractor{_ee: _eeb, _bd: page.Resources, _eda: *_cfg, _gfg: page.CropBox, _ebb: map[string]fontEntry{}, _db: map[string]textResult{}, _dfb: options, _cd: _ec, _dfd: _fe}
	if _bf._eda.Llx > _bf._eda.Urx {
		_ed.Log.Info("\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u0068\u0061\u0073\u0020\u0058\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0073\u0020r\u0065\u0076\u0065\u0072\u0073\u0065\u0064\u002e\u0020\u0025\u002e\u0032\u0066\u0020F\u0069x\u0069\u006e\u0067\u002e", _bf._eda)
		_bf._eda.Llx, _bf._eda.Urx = _bf._eda.Urx, _bf._eda.Llx
	}
	if _bf._eda.Lly > _bf._eda.Ury {
		_ed.Log.Info("\u004d\u0065\u0064\u0069\u0061\u0042o\u0078\u0020\u0068\u0061\u0073\u0020\u0059\u0020\u0063\u006f\u006f\u0072\u0064\u0069\u006e\u0061\u0074\u0065\u0073\u0020r\u0065\u0076\u0065\u0072\u0073\u0065\u0064\u002e\u0020\u0025\u002e\u0032\u0066\u0020F\u0069x\u0069\u006e\u0067\u002e", _bf._eda)
		_bf._eda.Lly, _bf._eda.Ury = _bf._eda.Ury, _bf._eda.Lly
	}
	return _bf, nil
}
func _dfcc(_fcfd *textWord, _gdgg float64, _fbge, _gccc rulingList) *wordBag {
	_gebbd := _bedf(_fcfd._cbfcee)
	_cedac := []*textWord{_fcfd}
	_affe := wordBag{_gegc: map[int][]*textWord{_gebbd: _cedac}, PdfRectangle: _fcfd.PdfRectangle, _bcef: _fcfd._fgdbd, _gebga: _gdgg, _cbfg: _fbge, _gaeg: _gccc}
	return &_affe
}
func (_fbaf *wordBag) applyRemovals(_fefd map[int]map[*textWord]struct{}) {
	for _ade, _bdgc := range _fefd {
		if len(_bdgc) == 0 {
			continue
		}
		_agaa := _fbaf._gegc[_ade]
		_bbae := len(_agaa) - len(_bdgc)
		if _bbae == 0 {
			delete(_fbaf._gegc, _ade)
			continue
		}
		_gacf := make([]*textWord, _bbae)
		_fbfc := 0
		for _, _efae := range _agaa {
			if _, _bcfc := _bdgc[_efae]; !_bcfc {
				_gacf[_fbfc] = _efae
				_fbfc++
			}
		}
		_fbaf._gegc[_ade] = _gacf
	}
}
func _aadb(_efffe, _eafe float64) string {
	_fegg := !_cdgd(_efffe - _eafe)
	if _fegg {
		return "\u000a"
	}
	return "\u0020"
}
func (_eggg *textTable) bbox() _gb.PdfRectangle { return _eggg.PdfRectangle }

var _bcbc = []string{"\u2756", "\u27a2", "\u2713", "\u2022", "\uf0a7", "\u25a1", "\u2212", "\u25a0", "\u25aa", "\u006f"}

func _fdddc(_adagg []pathSection) {
	if _bcdg < 0.0 {
		return
	}
	if _geaeg {
		_ed.Log.Info("\u0067\u0072\u0061\u006e\u0075\u006c\u0061\u0072\u0069\u007a\u0065\u003a\u0020\u0025\u0064 \u0073u\u0062\u0070\u0061\u0074\u0068\u0020\u0073\u0065\u0063\u0074\u0069\u006f\u006e\u0073", len(_adagg))
	}
	for _fbfd, _gaaa := range _adagg {
		for _ggdg, _eaeae := range _gaaa._ggfc {
			for _ageff, _cdba := range _eaeae._ccgf {
				_eaeae._ccgf[_ageff] = _cc.Point{X: _bfddb(_cdba.X), Y: _bfddb(_cdba.Y)}
				if _geaeg {
					_eggb := _eaeae._ccgf[_ageff]
					if !_ggadd(_cdba, _eggb) {
						_cfce := _cc.Point{X: _eggb.X - _cdba.X, Y: _eggb.Y - _cdba.Y}
						_cfe.Printf("\u0025\u0034d \u002d\u0020\u00254\u0064\u0020\u002d\u0020%4d\u003a %\u002e\u0032\u0066\u0020\u2192\u0020\u0025.2\u0066\u0020\u0028\u0025\u0067\u0029\u000a", _fbfd, _ggdg, _ageff, _cdba, _eggb, _cfce)
					}
				}
			}
		}
	}
}

// Len returns the number of TextMarks in `ma`.
func (_dfac *TextMarkArray) Len() int {
	if _dfac == nil {
		return 0
	}
	return len(_dfac._ggcf)
}

// TableInfo gets table information of the textmark `tm`.
func (_cega *TextMark) TableInfo() (*TextTable, [][]int) {
	if !_cega._fefe {
		return nil, nil
	}
	_cbbf := _cega._egc
	_cggb := _cbbf.getCellInfo(*_cega)
	return _cbbf, _cggb
}

const (
	_adab  = false
	_gbbcc = false
	_adea  = false
	_abb   = false
	_dgeg  = false
	_geafa = false
	_fdec  = false
	_baff  = false
	_gcga  = false
	_abgd  = _gcga && true
	_befe  = _abgd && false
	_fccg  = _gcga && true
	_geff  = false
	_abef  = _geff && false
	_cbadb = _geff && true
	_geaeg = false
	_faab  = _geaeg && false
	_dgbge = _geaeg && false
	_aaga  = _geaeg && true
	_afbec = _geaeg && false
	_aagad = _geaeg && false
)

func (_dfee *subpath) last() _cc.Point { return _dfee._ccgf[len(_dfee._ccgf)-1] }
func _fde(_cbbe, _faaa _gb.PdfRectangle) _gb.PdfRectangle {
	return _gb.PdfRectangle{Llx: _df.Min(_cbbe.Llx, _faaa.Llx), Lly: _df.Min(_cbbe.Lly, _faaa.Lly), Urx: _df.Max(_cbbe.Urx, _faaa.Urx), Ury: _df.Max(_cbbe.Ury, _faaa.Ury)}
}

// BBox returns the smallest axis-aligned rectangle that encloses all the TextMarks in `ma`.
func (_deeg *TextMarkArray) BBox() (_gb.PdfRectangle, bool) {
	var _bfcd _gb.PdfRectangle
	_ggbe := false
	for _, _cffg := range _deeg._ggcf {
		if _cffg.Meta || _edefac(_cffg.Text) {
			continue
		}
		if _ggbe {
			_bfcd = _fde(_bfcd, _cffg.BBox)
		} else {
			_bfcd = _cffg.BBox
			_ggbe = true
		}
	}
	return _bfcd, _ggbe
}
func (_efbed rectRuling) checkWidth(_cegbd, _ceggf float64) (float64, bool) {
	_badba := _ceggf - _cegbd
	_cece := _badba <= _dbgc
	return _badba, _cece
}
func (_ecdg *textTable) newTablePara() *textPara {
	_abea := _ecdg.computeBbox()
	_bgag := &textPara{PdfRectangle: _abea, _cedf: _abea, _bedfe: _ecdg}
	if _geff {
		_ed.Log.Info("\u006e\u0065w\u0054\u0061\u0062l\u0065\u0050\u0061\u0072\u0061\u003a\u0020\u0025\u0073", _bgag)
	}
	return _bgag
}
func (_egaf *textObject) setFont(_fdc string, _fgef float64) error {
	if _egaf == nil {
		return nil
	}
	_egaf._cbag._agge = _fgef
	_cfac, _aaad := _egaf.getFont(_fdc)
	if _aaad != nil {
		return _aaad
	}
	_egaf._cbag._bdg = _cfac
	return nil
}

type cachedImage struct {
	_gde *_gb.Image
	_ecg _gb.PdfColorspace
}

func _fdgd(_cfdc map[float64][]*textLine) []float64 {
	_afda := []float64{}
	for _bgcf := range _cfdc {
		_afda = append(_afda, _bgcf)
	}
	_d.Float64s(_afda)
	return _afda
}

// String returns a human readable description of `vecs`.
func (_adade rulingList) String() string {
	if len(_adade) == 0 {
		return "\u007b \u0045\u004d\u0050\u0054\u0059\u0020}"
	}
	_gfggd, _fdbe := _adade.vertsHorzs()
	_fcde := len(_gfggd)
	_gabe := len(_fdbe)
	if _fcde == 0 || _gabe == 0 {
		return _cfe.Sprintf("\u007b%\u0064\u0020\u0078\u0020\u0025\u0064}", _fcde, _gabe)
	}
	_bcbd := _gb.PdfRectangle{Llx: _gfggd[0]._abfg, Urx: _gfggd[_fcde-1]._abfg, Lly: _fdbe[_gabe-1]._abfg, Ury: _fdbe[0]._abfg}
	return _cfe.Sprintf("\u007b\u0025d\u0020\u0078\u0020%\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u007d", _fcde, _gabe, _bcbd)
}

// RangeOffset returns the TextMarks in `ma` that overlap text[start:end] in the extracted text.
// These are tm: `start` <= tm.Offset + len(tm.Text) && tm.Offset < `end` where
// `start` and `end` are offsets in the extracted text.
// NOTE: TextMarks can contain multiple characters. e.g. "ffi" for the ﬃ ligature so the first and
// last elements of the returned TextMarkArray may only partially overlap text[start:end].
func (_cce *TextMarkArray) RangeOffset(start, end int) (*TextMarkArray, error) {
	if _cce == nil {
		return nil, _c.New("\u006da\u003d\u003d\u006e\u0069\u006c")
	}
	if end < start {
		return nil, _cfe.Errorf("\u0065\u006e\u0064\u0020\u003c\u0020\u0073\u0074\u0061\u0072\u0074\u002e\u0020\u0052\u0061n\u0067\u0065\u004f\u0066\u0066\u0073\u0065\u0074\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064\u002e\u0020\u0073\u0074\u0061\u0072t=\u0025\u0064\u0020\u0065\u006e\u0064\u003d\u0025\u0064\u0020", start, end)
	}
	_fbf := len(_cce._ggcf)
	if _fbf == 0 {
		return _cce, nil
	}
	if start < _cce._ggcf[0].Offset {
		start = _cce._ggcf[0].Offset
	}
	if end > _cce._ggcf[_fbf-1].Offset+1 {
		end = _cce._ggcf[_fbf-1].Offset + 1
	}
	_ecdac := _d.Search(_fbf, func(_daec int) bool { return _cce._ggcf[_daec].Offset+len(_cce._ggcf[_daec].Text)-1 >= start })
	if !(0 <= _ecdac && _ecdac < _fbf) {
		_gad := _cfe.Errorf("\u004f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u002e\u0020\u0073\u0074\u0061\u0072\u0074\u003d%\u0064\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020\u006c\u0065\u006e\u003d\u0025\u0064\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025\u0076\u000a\u0009 \u006c\u0061\u0073\u0074\u003d%\u0076", start, _ecdac, _fbf, _cce._ggcf[0], _cce._ggcf[_fbf-1])
		return nil, _gad
	}
	_gfgf := _d.Search(_fbf, func(_fcdg int) bool { return _cce._ggcf[_fcdg].Offset > end-1 })
	if !(0 <= _gfgf && _gfgf < _fbf) {
		_aggd := _cfe.Errorf("\u004f\u0075\u0074\u0020\u006f\u0066\u0020r\u0061\u006e\u0067e\u002e\u0020\u0065n\u0064\u003d%\u0064\u0020\u0069\u0045\u006e\u0064=\u0025d \u006c\u0065\u006e\u003d\u0025\u0064\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025\u0076\u000a\u0009\u0020\u006c\u0061\u0073\u0074\u003d\u0025\u0076", end, _gfgf, _fbf, _cce._ggcf[0], _cce._ggcf[_fbf-1])
		return nil, _aggd
	}
	if _gfgf <= _ecdac {
		return nil, _cfe.Errorf("\u0069\u0045\u006e\u0064\u0020\u003c=\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003a\u0020\u0073\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020\u0065\u006ed\u003d\u0025\u0064\u0020\u0069\u0053\u0074\u0061\u0072\u0074\u003d\u0025\u0064\u0020i\u0045n\u0064\u003d\u0025\u0064", start, end, _ecdac, _gfgf)
	}
	return &TextMarkArray{_ggcf: _cce._ggcf[_ecdac:_gfgf]}, nil
}

// String returns a description of `b`.
func (_aece *wordBag) String() string {
	var _fgec []string
	for _, _fdac := range _aece.depthIndexes() {
		_bdfff := _aece._gegc[_fdac]
		for _, _aagc := range _bdfff {
			_fgec = append(_fgec, _aagc._bcfea)
		}
	}
	return _cfe.Sprintf("\u0025.\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065=\u0025\u002e\u0032\u0066\u0020\u0025\u0064\u0020\u0025\u0071", _aece.PdfRectangle, _aece._bcef, len(_fgec), _fgec)
}
func (_cceag intSet) del(_cacg int) { delete(_cceag, _cacg) }
func _aaffa(_cfdf map[float64]gridTile) []float64 {
	_fabb := make([]float64, 0, len(_cfdf))
	for _begge := range _cfdf {
		_fabb = append(_fabb, _begge)
	}
	_d.Float64s(_fabb)
	return _fabb
}
func (_ffgb *textObject) moveText(_gdec, _cab float64) { _ffgb.moveLP(_gdec, _cab) }
func _ggadd(_aaag, _gbge _cc.Point) bool               { return _aaag.X == _gbge.X && _aaag.Y == _gbge.Y }
func _ffbbc(_beec []*textWord, _gffadc int) []*textWord {
	_bdgdc := len(_beec)
	copy(_beec[_gffadc:], _beec[_gffadc+1:])
	return _beec[:_bdgdc-1]
}
func _eddb(_bdfb []*textLine, _dfga map[float64][]*textLine, _dcebcb []float64, _cdgge int, _fga, _fbdf float64) []*list {
	_acga := []*list{}
	_cbfb := _cdgge
	_cdgge = _cdgge + 1
	_gddb := _dcebcb[_cbfb]
	_caccf := _dfga[_gddb]
	_gdccc := _feed(_caccf, _fbdf, _fga)
	for _bagg, _beee := range _gdccc {
		var _dgad float64
		_cbeb := []*list{}
		_edcf := _beee._bdbg
		_aaea := _fbdf
		if _bagg < len(_gdccc)-1 {
			_aaea = _gdccc[_bagg+1]._bdbg
		}
		if _cdgge < len(_dcebcb) {
			_cbeb = _eddb(_bdfb, _dfga, _dcebcb, _cdgge, _edcf, _aaea)
		}
		_dgad = _aaea
		if len(_cbeb) > 0 {
			_dbba := _cbeb[0]
			if len(_dbba._ggbee) > 0 {
				_dgad = _dbba._ggbee[0]._bdbg
			}
		}
		_cfff := []*textLine{_beee}
		_ffceg := _bbfg(_beee, _bdfb, _dcebcb, _edcf, _dgad)
		_cfff = append(_cfff, _ffceg...)
		_bbfa := _gbcag(_cfff, "\u0062\u0075\u006c\u006c\u0065\u0074", _cbeb)
		_bbfa._ebabg = _adag(_cfff, "")
		_acga = append(_acga, _bbfa)
	}
	return _acga
}

// PageFonts represents extracted fonts on a PDF page.
type PageFonts struct{ Fonts []Font }

func _ggbc(_deff bounded) float64 { return -_deff.bbox().Lly }

// String returns a human readable description of `path`.
func (_cfeb *subpath) String() string {
	_adf := _cfeb._ccgf
	_dfca := len(_adf)
	if _dfca <= 5 {
		return _cfe.Sprintf("\u0025d\u003a\u0020\u0025\u0036\u002e\u0032f", _dfca, _adf)
	}
	return _cfe.Sprintf("\u0025d\u003a\u0020\u0025\u0036.\u0032\u0066\u0020\u0025\u0036.\u0032f\u0020.\u002e\u002e\u0020\u0025\u0036\u002e\u0032f", _dfca, _adf[0], _adf[1], _adf[_dfca-1])
}
func _bbg(_caega *textLine) bool {
	_bgbe := true
	_dfcag := -1
	for _, _dddeff := range _caega._bfca {
		for _, _aage := range _dddeff._fdgbc {
			_faef := _aage._afae
			if _dfcag == -1 {
				_dfcag = _faef
			} else {
				if _dfcag != _faef {
					_bgbe = false
					break
				}
			}
		}
	}
	return _bgbe
}

// String returns a string descibing `i`.
func (_befc gridTile) String() string {
	_cecc := func(_gdaa bool, _gddbg string) string {
		if _gdaa {
			return _gddbg
		}
		return "\u005f"
	}
	return _cfe.Sprintf("\u00256\u002e2\u0066\u0020\u0025\u0031\u0073%\u0031\u0073%\u0031\u0073\u0025\u0031\u0073", _befc.PdfRectangle, _cecc(_befc._bdcfac, "\u004c"), _cecc(_befc._babe, "\u0052"), _cecc(_befc._debb, "\u0042"), _cecc(_befc._cccb, "\u0054"))
}
func (_addce gridTile) contains(_eedd _gb.PdfRectangle) bool {
	if _addce.numBorders() < 3 {
		return false
	}
	if _addce._bdcfac && _eedd.Llx < _addce.Llx-_caeg {
		return false
	}
	if _addce._babe && _eedd.Urx > _addce.Urx+_caeg {
		return false
	}
	if _addce._debb && _eedd.Lly < _addce.Lly-_caeg {
		return false
	}
	if _addce._cccb && _eedd.Ury > _addce.Ury+_caeg {
		return false
	}
	return true
}
func (_aaafg *textTable) log(_bfdd string) {
	if !_geff {
		return
	}
	_ed.Log.Info("~\u007e\u007e\u0020\u0025\u0073\u003a \u0025\u0064\u0020\u0078\u0020\u0025d\u0020\u0067\u0072\u0069\u0064\u003d\u0025t\u000a\u0020\u0020\u0020\u0020\u0020\u0020\u0025\u0036\u002e2\u0066", _bfdd, _aaafg._fccgee, _aaafg._dege, _aaafg._baccf, _aaafg.PdfRectangle)
	for _adedf := 0; _adedf < _aaafg._dege; _adedf++ {
		for _bffg := 0; _bffg < _aaafg._fccgee; _bffg++ {
			_efcag := _aaafg.get(_bffg, _adedf)
			if _efcag == nil {
				continue
			}
			_cfe.Printf("%\u0034\u0064\u0020\u00252d\u003a \u0025\u0036\u002e\u0032\u0066 \u0025\u0071\u0020\u0025\u0064\u000a", _bffg, _adedf, _efcag.PdfRectangle, _fabf(_efcag.text(), 50), _cb.RuneCountInString(_efcag.text()))
		}
	}
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

func (_fgeb *wordBag) firstWord(_cdcg int) *textWord { return _fgeb._gegc[_cdcg][0] }
func (_cacb rulingList) connections(_edbfd map[int]intSet, _gadd int) intSet {
	_ddcg := make(intSet)
	_agcag := make(intSet)
	var _faaee func(int)
	_faaee = func(_fegee int) {
		if !_agcag.has(_fegee) {
			_agcag.add(_fegee)
			for _debc := range _cacb {
				if _edbfd[_debc].has(_fegee) {
					_ddcg.add(_debc)
				}
			}
			for _cbebd := range _cacb {
				if _ddcg.has(_cbebd) {
					_faaee(_cbebd)
				}
			}
		}
	}
	_faaee(_gadd)
	return _ddcg
}
func (_ddcdb *wordBag) sort() {
	for _, _acaf := range _ddcdb._gegc {
		_d.Slice(_acaf, func(_ccfg, _dab int) bool { return _fdd(_acaf[_ccfg], _acaf[_dab]) < 0 })
	}
}
func (_aggg *textTable) getDown() paraList {
	_acbdc := make(paraList, _aggg._fccgee)
	for _cbff := 0; _cbff < _aggg._fccgee; _cbff++ {
		_efgfc := _aggg.get(_cbff, _aggg._dege-1)._cgfe
		if _efgfc.taken() {
			return nil
		}
		_acbdc[_cbff] = _efgfc
	}
	for _bded := 0; _bded < _aggg._fccgee-1; _bded++ {
		if _acbdc[_bded]._ccdg != _acbdc[_bded+1] {
			return nil
		}
	}
	return _acbdc
}

type rectRuling struct {
	_gcaad rulingKind
	_abgg  markKind
	_eb.Color
	_gb.PdfRectangle
}

func _bdda(_agga *_gb.Image, _cegfa _eb.Color) _gf.Image {
	_bdcg, _eebf := int(_agga.Width), int(_agga.Height)
	_geefb := _gf.NewRGBA(_gf.Rect(0, 0, _bdcg, _eebf))
	for _abcbe := 0; _abcbe < _eebf; _abcbe++ {
		for _gcfg := 0; _gcfg < _bdcg; _gcfg++ {
			_geade, _dcaf := _agga.ColorAt(_gcfg, _abcbe)
			if _dcaf != nil {
				_ed.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _gcfg, _abcbe)
				continue
			}
			_febf, _abagd, _ecgcb, _ := _geade.RGBA()
			var _eaaa _eb.Color
			if _febf+_abagd+_ecgcb == 0 {
				_eaaa = _eb.Transparent
			} else {
				_eaaa = _cegfa
			}
			_geefb.Set(_gcfg, _abcbe, _eaaa)
		}
	}
	return _geefb
}
func (_bdgfd rulingList) primMinMax() (float64, float64) {
	_aeadd, _ccad := _bdgfd[0]._abfg, _bdgfd[0]._abfg
	for _, _gfda := range _bdgfd[1:] {
		if _gfda._abfg < _aeadd {
			_aeadd = _gfda._abfg
		} else if _gfda._abfg > _ccad {
			_ccad = _gfda._abfg
		}
	}
	return _aeadd, _ccad
}
func (_dedga rulingList) toTilings() (rulingList, []gridTiling) {
	_dedga.log("\u0074o\u0054\u0069\u006c\u0069\u006e\u0067s")
	if len(_dedga) == 0 {
		return nil, nil
	}
	_dedga = _dedga.tidied("\u0061\u006c\u006c")
	_dedga.log("\u0074\u0069\u0064\u0069\u0065\u0064")
	_fbbab := _dedga.toGrids()
	_bdfba := make([]gridTiling, len(_fbbab))
	for _bebf, _bfcdb := range _fbbab {
		_bdfba[_bebf] = _bfcdb.asTiling()
	}
	return _dedga, _bdfba
}
func (_fdcg rulingList) intersections() map[int]intSet {
	var _eaeea, _gacg []int
	for _bgffe, _edbc := range _fdcg {
		switch _edbc._cffe {
		case _gfafc:
			_eaeea = append(_eaeea, _bgffe)
		case _bdfd:
			_gacg = append(_gacg, _bgffe)
		}
	}
	if len(_eaeea) < _edeec+1 || len(_gacg) < _fegbc+1 {
		return nil
	}
	if len(_eaeea)+len(_gacg) > _aaaa {
		_ed.Log.Debug("\u0069\u006e\u0074\u0065\u0072\u0073e\u0063\u0074\u0069\u006f\u006e\u0073\u003a\u0020\u0054\u004f\u004f\u0020\u004d\u0041\u004e\u0059\u0020\u0072\u0075\u006ci\u006e\u0067\u0073\u0020\u0076\u0065\u0063\u0073\u003d\u0025\u0064\u0020\u003d\u0020%\u0064 \u0078\u0020\u0025\u0064", len(_fdcg), len(_eaeea), len(_gacg))
		return nil
	}
	_cddg := make(map[int]intSet, len(_eaeea)+len(_gacg))
	for _, _gggeg := range _eaeea {
		for _, _afbg := range _gacg {
			if _fdcg[_gggeg].intersects(_fdcg[_afbg]) {
				if _, _acgb := _cddg[_gggeg]; !_acgb {
					_cddg[_gggeg] = make(intSet)
				}
				if _, _fdcb := _cddg[_afbg]; !_fdcb {
					_cddg[_afbg] = make(intSet)
				}
				_cddg[_gggeg].add(_afbg)
				_cddg[_afbg].add(_gggeg)
			}
		}
	}
	return _cddg
}

var _gegca = map[rulingKind]string{_eefaf: "\u006e\u006f\u006e\u0065", _bdfd: "\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c", _gfafc: "\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c"}

func (_beeb *textPara) toCellTextMarks(_bbdbe *int) []TextMark {
	var _edefc []TextMark
	for _gba, _eabag := range _beeb._abbe {
		_bceb := _eabag.toTextMarks(_bbdbe)
		_gbaf := _ddfg && _eabag.endsInHyphen() && _gba != len(_beeb._abbe)-1
		if _gbaf {
			_bceb = _aege(_bceb, _bbdbe)
		}
		_edefc = append(_edefc, _bceb...)
		if !(_gbaf || _gba == len(_beeb._abbe)-1) {
			_edefc = _dgccag(_edefc, _bbdbe, _aadb(_eabag._bdbg, _beeb._abbe[_gba+1]._bdbg))
		}
	}
	return _edefc
}
func _aafb(_deac, _bfbgb int) uint64 { return uint64(_deac)*0x1000000 + uint64(_bfbgb) }
func (_dgcd *wordBag) makeRemovals() map[int]map[*textWord]struct{} {
	_fdbf := make(map[int]map[*textWord]struct{}, len(_dgcd._gegc))
	for _gbfe := range _dgcd._gegc {
		_fdbf[_gbfe] = make(map[*textWord]struct{})
	}
	return _fdbf
}
func (_feaae paraList) sortTopoOrder() {
	_adfbe := _feaae.topoOrder()
	_feaae.reorder(_adfbe)
}
func _ecdf(_cabc, _eccc bounded) float64 {
	_bbaea := _bafe(_cabc, _eccc)
	if !_cdgd(_bbaea) {
		return _bbaea
	}
	return _fdd(_cabc, _eccc)
}
func _gcda(_beegg func(*wordBag, *textWord, float64) bool, _eabd float64) func(*wordBag, *textWord) bool {
	return func(_edfb *wordBag, _aaed *textWord) bool { return _beegg(_edfb, _aaed, _eabd) }
}
func (_bcfd *wordBag) minDepth() float64 { return _bcfd._gebga - (_bcfd.Ury - _bcfd._bcef) }
func _eaf(_cbe []Font, _bdc string) bool {
	for _, _ebbd := range _cbe {
		if _ebbd.FontName == _bdc {
			return true
		}
	}
	return false
}
func (_dbccd gridTile) numBorders() int {
	_cagc := 0
	if _dbccd._bdcfac {
		_cagc++
	}
	if _dbccd._babe {
		_cagc++
	}
	if _dbccd._debb {
		_cagc++
	}
	if _dbccd._cccb {
		_cagc++
	}
	return _cagc
}
func (_gbd *wordBag) depthRange(_bgf, _caeb int) []int {
	var _ddab []int
	for _eaea := range _gbd._gegc {
		if _bgf <= _eaea && _eaea <= _caeb {
			_ddab = append(_ddab, _eaea)
		}
	}
	if len(_ddab) == 0 {
		return nil
	}
	_d.Ints(_ddab)
	return _ddab
}

// ImageExtractOptions contains options for controlling image extraction from
// PDF pages.
type ImageExtractOptions struct{ IncludeInlineStencilMasks bool }

func _beba(_bbc _gb.PdfRectangle) textState {
	return textState{_fceg: 100, _eacgf: RenderModeFill, _afd: _bbc}
}
func _fdd(_cegbb, _bfcgc bounded) float64 { return _cegbb.bbox().Llx - _bfcgc.bbox().Llx }
func (_cgdc rulingList) snapToGroupsDirection() rulingList {
	_cgdc.sortStrict()
	_dbcb := make(map[*ruling]rulingList, len(_cgdc))
	_ecfcg := _cgdc[0]
	_decdg := func(_fgbb *ruling) { _ecfcg = _fgbb; _dbcb[_ecfcg] = rulingList{_fgbb} }
	_decdg(_cgdc[0])
	for _, _efgc := range _cgdc[1:] {
		if _efgc._abfg < _ecfcg._abfg-_dafe {
			_ed.Log.Error("\u0073\u006e\u0061\u0070T\u006f\u0047\u0072\u006f\u0075\u0070\u0073\u0044\u0069r\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0057\u0072\u006f\u006e\u0067\u0020\u0070\u0072\u0069\u006da\u0072\u0079\u0020\u006f\u0072d\u0065\u0072\u002e\u000a\u0009\u0076\u0030\u003d\u0025\u0073\u000a\u0009\u0020\u0076\u003d\u0025\u0073", _ecfcg, _efgc)
		}
		if _efgc._abfg > _ecfcg._abfg+_dbgc {
			_decdg(_efgc)
		} else {
			_dbcb[_ecfcg] = append(_dbcb[_ecfcg], _efgc)
		}
	}
	_gagae := make(map[*ruling]float64, len(_dbcb))
	_efedc := make(map[*ruling]*ruling, len(_cgdc))
	for _ddbe, _fbag := range _dbcb {
		_gagae[_ddbe] = _fbag.mergePrimary()
		for _, _ebeb := range _fbag {
			_efedc[_ebeb] = _ddbe
		}
	}
	for _, _fbgg := range _cgdc {
		_fbgg._abfg = _gagae[_efedc[_fbgg]]
	}
	_agab := make(rulingList, 0, len(_cgdc))
	for _, _agag := range _dbcb {
		_aacc := _agag.splitSec()
		for _fdgcb, _ebcg := range _aacc {
			_cegga := _ebcg.merge()
			if len(_agab) > 0 {
				_ececc := _agab[len(_agab)-1]
				if _ececc.alignsPrimary(_cegga) && _ececc.alignsSec(_cegga) {
					_ed.Log.Error("\u0073\u006e\u0061\u0070\u0054\u006fG\u0072\u006f\u0075\u0070\u0073\u0044\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0044\u0075\u0070\u006ci\u0063\u0061\u0074\u0065\u0020\u0069\u003d\u0025\u0064\u000a\u0009\u0077\u003d\u0025s\u000a\t\u0076\u003d\u0025\u0073", _fdgcb, _ececc, _cegga)
					continue
				}
			}
			_agab = append(_agab, _cegga)
		}
	}
	_agab.sortStrict()
	return _agab
}
func _fabf(_bdea string, _feafa int) string {
	if len(_bdea) < _feafa {
		return _bdea
	}
	return _bdea[:_feafa]
}

type rulingList []*ruling

func (_ggecc *ruling) intersects(_abfe *ruling) bool {
	_cbcge := (_ggecc._cffe == _gfafc && _abfe._cffe == _bdfd) || (_abfe._cffe == _gfafc && _ggecc._cffe == _bdfd)
	_cdbfb := func(_cdcgg, _edbb *ruling) bool {
		return _cdcgg._daag-_egaa <= _edbb._abfg && _edbb._abfg <= _cdcgg._fcff+_egaa
	}
	_facgb := _cdbfb(_ggecc, _abfe)
	_baffg := _cdbfb(_abfe, _ggecc)
	if _geaeg {
		_cfe.Printf("\u0020\u0020\u0020\u0020\u0069\u006e\u0074\u0065\u0072\u0073\u0065\u0063\u0074\u0073\u003a\u0020\u0020\u006fr\u0074\u0068\u006f\u0067\u006f\u006e\u0061l\u003d\u0025\u0074\u0020\u006f\u0031\u003d\u0025\u0074\u0020\u006f2\u003d\u0025\u0074\u0020\u2192\u0020\u0025\u0074\u000a"+"\u0020\u0020\u0020 \u0020\u0020\u0020\u0076\u003d\u0025\u0073\u000a"+" \u0020\u0020\u0020\u0020\u0020\u0077\u003d\u0025\u0073\u000a", _cbcge, _facgb, _baffg, _cbcge && _facgb && _baffg, _ggecc, _abfe)
	}
	return _cbcge && _facgb && _baffg
}

type bounded interface{ bbox() _gb.PdfRectangle }

func (_gag *subpath) removeDuplicates() {
	if len(_gag._ccgf) == 0 {
		return
	}
	_egbc := []_cc.Point{_gag._ccgf[0]}
	for _, _dfeea := range _gag._ccgf[1:] {
		if !_ggadd(_dfeea, _egbc[len(_egbc)-1]) {
			_egbc = append(_egbc, _dfeea)
		}
	}
	_gag._ccgf = _egbc
}
func (_acda *textWord) appendMark(_fadc *textMark, _bfdgg _gb.PdfRectangle) {
	_acda._fdgbc = append(_acda._fdgbc, _fadc)
	_acda.PdfRectangle = _fde(_acda.PdfRectangle, _fadc.PdfRectangle)
	if _fadc._gded > _acda._fgdbd {
		_acda._fgdbd = _fadc._gded
	}
	_acda._cbfcee = _bfdgg.Ury - _acda.PdfRectangle.Lly
}
func _cebdd(_gdcgb []compositeCell) []float64 {
	var _cgfbe []*textLine
	_gceda := 0
	for _, _gffad := range _gdcgb {
		_gceda += len(_gffad.paraList)
		_cgfbe = append(_cgfbe, _gffad.lines()...)
	}
	_d.Slice(_cgfbe, func(_gdbf, _ddgb int) bool {
		_ddecd, _fefdd := _cgfbe[_gdbf], _cgfbe[_ddgb]
		_adcg, _gadea := _ddecd._bdbg, _fefdd._bdbg
		if !_cdgd(_adcg - _gadea) {
			return _adcg < _gadea
		}
		return _ddecd.Llx < _fefdd.Llx
	})
	if _geff {
		_cfe.Printf("\u0020\u0020\u0020 r\u006f\u0077\u0042\u006f\u0072\u0064\u0065\u0072\u0073:\u0020%\u0064 \u0070a\u0072\u0061\u0073\u0020\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u000a", _gceda, len(_cgfbe))
		for _ddege, _gddc := range _cgfbe {
			_cfe.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _ddege, _gddc)
		}
	}
	var _gbbdb []float64
	_cebga := _cgfbe[0]
	var _edbgf [][]*textLine
	_gcef := []*textLine{_cebga}
	for _feca, _aafd := range _cgfbe[1:] {
		if _aafd.Ury < _cebga.Lly {
			_ccbad := 0.5 * (_aafd.Ury + _cebga.Lly)
			if _geff {
				_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u003c\u0020\u0025\u0036.\u0032f\u0020\u0062\u006f\u0072\u0064\u0065\u0072\u003d\u0025\u0036\u002e\u0032\u0066\u000a"+"\u0009\u0020\u0071\u003d\u0025\u0073\u000a\u0009\u0020p\u003d\u0025\u0073\u000a", _feca, _aafd.Ury, _cebga.Lly, _ccbad, _cebga, _aafd)
			}
			_gbbdb = append(_gbbdb, _ccbad)
			_edbgf = append(_edbgf, _gcef)
			_gcef = nil
		}
		_gcef = append(_gcef, _aafd)
		if _aafd.Lly < _cebga.Lly {
			_cebga = _aafd
		}
	}
	if len(_gcef) > 0 {
		_edbgf = append(_edbgf, _gcef)
	}
	if _geff {
		_cfe.Printf(" \u0020\u0020\u0020\u0020\u0020\u0020 \u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064o\u0072\u0073\u003d%\u0036.\u0032\u0066\u000a", _gbbdb)
	}
	if _geff {
		_ed.Log.Info("\u0072\u006f\u0077\u003d\u0025\u0064", len(_gdcgb))
		for _acca, _gebc := range _gdcgb {
			_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _acca, _gebc)
		}
		_ed.Log.Info("\u0067r\u006f\u0075\u0070\u0073\u003d\u0025d", len(_edbgf))
		for _dfbfg, _efdd := range _edbgf {
			_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0064\u000a", _dfbfg, len(_efdd))
			for _agdfgg, _eaaf := range _efdd {
				_cfe.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _agdfgg, _eaaf)
			}
		}
	}
	_gfcf := true
	for _acdb, _ccdf := range _edbgf {
		_dfdcbe := true
		for _ddged, _agdg := range _gdcgb {
			if _geff {
				_cfe.Printf("\u0020\u0020\u0020\u007e\u007e\u007e\u0067\u0072\u006f\u0075\u0070\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u0020\u0063\u0065\u006cl\u0020\u0025\u0064\u0020\u006ff\u0020\u0025d\u0020\u0025\u0073\u000a", _acdb, len(_edbgf), _ddged, len(_gdcgb), _agdg)
			}
			if !_agdg.hasLines(_ccdf) {
				if _geff {
					_cfe.Printf("\u0020\u0020\u0020\u0021\u0021\u0021\u0067\u0072\u006f\u0075\u0070\u0020\u0025d\u0020\u006f\u0066\u0020\u0025\u0064 \u0063\u0065\u006c\u006c\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064 \u004f\u0055\u0054\u000a", _acdb, len(_edbgf), _ddged, len(_gdcgb))
				}
				_dfdcbe = false
				break
			}
		}
		if !_dfdcbe {
			_gfcf = false
			break
		}
	}
	if !_gfcf {
		if _geff {
			_ed.Log.Info("\u0072\u006f\u0077\u0020\u0063o\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u0020\u0064\u006f\u006e\u0027\u0074 \u0073\u0070\u0061\u006e\u0020\u0061\u006c\u006c\u0020\u0063\u0065\u006c\u006c\u0073\u0020\u0069\u006e\u0020\u0072\u006f\u0077\u002e\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006eg")
		}
		_gbbdb = nil
	}
	if _geff && _gbbdb != nil {
		_cfe.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u002a\u002a*\u0072\u006f\u0077\u0043\u006f\u0072\u0072i\u0064\u006f\u0072\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u000a", _gbbdb)
	}
	return _gbbdb
}
func (_cedg *textPara) writeCellText(_ecgce _a.Writer) {
	for _bcgf, _bedd := range _cedg._abbe {
		_ageada := _bedd.text()
		_fcca := _ddfg && _bedd.endsInHyphen() && _bcgf != len(_cedg._abbe)-1
		if _fcca {
			_ageada = _aabe(_ageada)
		}
		_ecgce.Write([]byte(_ageada))
		if !(_fcca || _bcgf == len(_cedg._abbe)-1) {
			_ecgce.Write([]byte(_aadb(_bedd._bdbg, _cedg._abbe[_bcgf+1]._bdbg)))
		}
	}
}
func (_ddeac *wordBag) getDepthIdx(_dbgac float64) int {
	_affg := _ddeac.depthIndexes()
	_fee := _bedf(_dbgac)
	if _fee < _affg[0] {
		return _affg[0]
	}
	if _fee > _affg[len(_affg)-1] {
		return _affg[len(_affg)-1]
	}
	return _fee
}

type paraList []*textPara
type fontEntry struct {
	_cegb *_gb.PdfFont
	_dgca int64
}

func (_dagbg paraList) topoOrder() []int {
	if _baff {
		_ed.Log.Info("\u0074\u006f\u0070\u006f\u004f\u0072\u0064\u0065\u0072\u003a")
	}
	_dbec := len(_dagbg)
	_eabca := make([]bool, _dbec)
	_ecfdb := make([]int, 0, _dbec)
	_edede := _dagbg.llyOrdering()
	var _ebgb func(_cgbcb int)
	_ebgb = func(_cbfgfb int) {
		_eabca[_cbfgfb] = true
		for _fddd := 0; _fddd < _dbec; _fddd++ {
			if !_eabca[_fddd] {
				if _dagbg.readBefore(_edede, _cbfgfb, _fddd) {
					_ebgb(_fddd)
				}
			}
		}
		_ecfdb = append(_ecfdb, _cbfgfb)
	}
	for _geeeg := 0; _geeeg < _dbec; _geeeg++ {
		if !_eabca[_geeeg] {
			_ebgb(_geeeg)
		}
	}
	return _edde(_ecfdb)
}

type event struct {
	_bggcc  float64
	_aagcb  bool
	_caccdd int
}

const (
	_dafe   = 1.0e-6
	_bcdg   = 1.0e-4
	_ecce   = 10
	_edagd  = 6
	_dgfg   = 0.5
	_cgc    = 0.12
	_gbdc   = 0.19
	_cedace = 0.04
	_fcga   = 0.04
	_dccd   = 1.0
	_aeg    = 0.04
	_debd   = 0.4
	_cdcgd  = 0.7
	_cbfc   = 1.0
	_cdce   = 0.1
	_dedd   = 1.4
	_ccfe   = 0.46
	_bacc   = 0.02
	_gcca   = 0.2
	_fabc   = 0.5
	_ddcf   = 4
	_cfab   = 4.0
	_dgea   = 6
	_dgcca  = 0.3
	_agcd   = 0.01
	_gbea   = 0.02
	_edeec  = 2
	_fegbc  = 2
	_aaaa   = 500
	_caf    = 4.0
	_eee    = 4.0
	_gdde   = 0.05
	_cccd   = 0.1
	_egaa   = 2.0
	_dbgc   = 2.0
	_caeg   = 1.5
	_fbgec  = 3.0
	_ecdaa  = 0.25
)

func (_abecc *textWord) absorb(_edbfa *textWord) {
	_abecc.PdfRectangle = _fde(_abecc.PdfRectangle, _edbfa.PdfRectangle)
	_abecc._fdgbc = append(_abecc._fdgbc, _edbfa._fdgbc...)
}

// Tables returns the tables extracted from the page.
func (_gdba PageText) Tables() []TextTable {
	if _geff {
		_ed.Log.Info("\u0054\u0061\u0062\u006c\u0065\u0073\u003a\u0020\u0025\u0064", len(_gdba._ded))
	}
	return _gdba._ded
}
func _dbfg(_aeeg, _caaa _cc.Point) rulingKind {
	_eced := _df.Abs(_aeeg.X - _caaa.X)
	_gggd := _df.Abs(_aeeg.Y - _caaa.Y)
	return _gaba(_eced, _gggd, _gdde)
}
func (_dffc *textPara) getListLines() []*textLine {
	var _cgca []*textLine
	_cegd := _eaag(_dffc._abbe)
	for _, _agd := range _dffc._abbe {
		_gcaa := _agd._bfca[0]._bcfea[0]
		if _cacc(_gcaa) {
			_cgca = append(_cgca, _agd)
		}
	}
	_cgca = append(_cgca, _cegd...)
	return _cgca
}

type shapesState struct {
	_cca  _cc.Matrix
	_begf _cc.Matrix
	_babc []*subpath
	_dfff bool
	_bcb  _cc.Point
	_geba *textObject
}

func (_ffg *PageFonts) extractPageResourcesToFont(_ada *_gb.PdfPageResources) error {
	_dfc, _gd := _gg.GetDict(_ada.Font)
	if !_gd {
		return _c.New(_cfd)
	}
	for _, _dcc := range _dfc.Keys() {
		var (
			_ecd = true
			_aac []byte
			_eac string
		)
		_dce, _ebe := _ada.GetFontByName(_dcc)
		if !_ebe {
			return _c.New(_cad)
		}
		_efc, _eba := _gb.NewPdfFontFromPdfObject(_dce)
		if _eba != nil {
			return _eba
		}
		_eacd := _efc.FontDescriptor()
		_ddc := _efc.FontDescriptor().FontName.String()
		_ddd := _efc.Subtype()
		if _eaf(_ffg.Fonts, _ddc) {
			continue
		}
		if len(_efc.ToUnicode()) == 0 {
			_ecd = false
		}
		if _eacd.FontFile != nil {
			if _ebbb, _dgg := _gg.GetStream(_eacd.FontFile); _dgg {
				_aac, _eba = _gg.DecodeStream(_ebbb)
				if _eba != nil {
					return _eba
				}
				_eac = _ddc + "\u002e\u0070\u0066\u0062"
			}
		} else if _eacd.FontFile2 != nil {
			if _de, _ab := _gg.GetStream(_eacd.FontFile2); _ab {
				_aac, _eba = _gg.DecodeStream(_de)
				if _eba != nil {
					return _eba
				}
				_eac = _ddc + "\u002e\u0074\u0074\u0066"
			}
		} else if _eacd.FontFile3 != nil {
			if _fdg, _bec := _gg.GetStream(_eacd.FontFile3); _bec {
				_aac, _eba = _gg.DecodeStream(_fdg)
				if _eba != nil {
					return _eba
				}
				_eac = _ddc + "\u002e\u0063\u0066\u0066"
			}
		}
		if len(_eac) < 1 {
			_ed.Log.Debug(_ea)
		}
		_aga := Font{FontName: _ddc, PdfFont: _efc, IsCID: _efc.IsCID(), IsSimple: _efc.IsSimple(), ToUnicode: _ecd, FontType: _ddd, FontData: _aac, FontFileName: _eac, FontDescriptor: _eacd}
		_ffg.Fonts = append(_ffg.Fonts, _aga)
	}
	return nil
}
func (_egeea paraList) findTables(_eaac []gridTiling) []*textTable {
	_egeea.addNeighbours()
	_d.Slice(_egeea, func(_cade, _ebabgc int) bool { return _gede(_egeea[_cade], _egeea[_ebabgc]) < 0 })
	var _eecgd []*textTable
	if _ecbbf {
		_faba := _egeea.findGridTables(_eaac)
		_eecgd = append(_eecgd, _faba...)
	}
	if _cbaf {
		_gbdfb := _egeea.findTextTables()
		_eecgd = append(_eecgd, _gbdfb...)
	}
	return _eecgd
}
func _dbff(_bbdc []pathSection) rulingList {
	_fdddc(_bbdc)
	if _geaeg {
		_ed.Log.Info("\u006d\u0061k\u0065\u0053\u0074\u0072\u006f\u006b\u0065\u0052\u0075\u006c\u0069\u006e\u0067\u0073\u003a\u0020\u0025\u0064\u0020\u0073\u0074\u0072ok\u0065\u0073", len(_bbdc))
	}
	var _bgdg rulingList
	for _, _cdbc := range _bbdc {
		for _, _fedb := range _cdbc._ggfc {
			if len(_fedb._ccgf) < 2 {
				continue
			}
			_badgb := _fedb._ccgf[0]
			for _, _faad := range _fedb._ccgf[1:] {
				if _ageac, _aeda := _caff(_badgb, _faad, _cdbc.Color); _aeda {
					_bgdg = append(_bgdg, _ageac)
				}
				_badgb = _faad
			}
		}
	}
	if _geaeg {
		_ed.Log.Info("m\u0061\u006b\u0065\u0053tr\u006fk\u0065\u0052\u0075\u006c\u0069n\u0067\u0073\u003a\u0020\u0025\u0073", _bgdg)
	}
	return _bgdg
}
func _cacc(_gedee byte) bool {
	for _, _cfba := range _bcbc {
		if []byte(_cfba)[0] == _gedee {
			return true
		}
	}
	return false
}
func (_ffca *textTable) computeBbox() _gb.PdfRectangle {
	var _gcbbf _gb.PdfRectangle
	_efffae := false
	for _bfggc := 0; _bfggc < _ffca._dege; _bfggc++ {
		for _feacfd := 0; _feacfd < _ffca._fccgee; _feacfd++ {
			_eedgb := _ffca.get(_feacfd, _bfggc)
			if _eedgb == nil {
				continue
			}
			if !_efffae {
				_gcbbf = _eedgb.PdfRectangle
				_efffae = true
			} else {
				_gcbbf = _fde(_gcbbf, _eedgb.PdfRectangle)
			}
		}
	}
	return _gcbbf
}
func (_cgbdb rulingList) asTiling() gridTiling {
	if _aaga {
		_ed.Log.Info("r\u0075\u006ci\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0076\u0065\u0063s\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u002b\u002b\u002b\u0020\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d", len(_cgbdb))
	}
	for _gaga, _efbee := range _cgbdb[1:] {
		_dabba := _cgbdb[_gaga]
		if _dabba.alignsPrimary(_efbee) && _dabba.alignsSec(_efbee) {
			_ed.Log.Error("a\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0044\u0075\u0070\u006c\u0069\u0063\u0061\u0074\u0065 \u0072\u0075\u006c\u0069\u006e\u0067\u0073\u002e\u000a\u0009v=\u0025\u0073\u000a\t\u0077=\u0025\u0073", _efbee, _dabba)
		}
	}
	_cgbdb.sortStrict()
	_cgbdb.log("\u0073n\u0061\u0070\u0070\u0065\u0064")
	_ggfa, _feeb := _cgbdb.vertsHorzs()
	_bdba := _ggfa.primaries()
	_efga := _feeb.primaries()
	_degf := len(_bdba) - 1
	_ecae := len(_efga) - 1
	if _degf == 0 || _ecae == 0 {
		return gridTiling{}
	}
	_egce := _gb.PdfRectangle{Llx: _bdba[0], Urx: _bdba[_degf], Lly: _efga[0], Ury: _efga[_ecae]}
	if _aaga {
		_ed.Log.Info("\u0072\u0075l\u0069\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0076\u0065\u0072\u0074s=\u0025\u0064", len(_ggfa))
		for _dcbca, _ddfe := range _ggfa {
			_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _dcbca, _ddfe)
		}
		_ed.Log.Info("\u0072\u0075l\u0069\u006e\u0067\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0068\u006f\u0072\u007as=\u0025\u0064", len(_feeb))
		for _abggf, _ccfgb := range _feeb {
			_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _abggf, _ccfgb)
		}
		_ed.Log.Info("\u0072\u0075\u006c\u0069\u006eg\u004c\u0069\u0073\u0074\u002e\u0061\u0073\u0054\u0069\u006c\u0069\u006e\u0067:\u0020\u0020\u0077\u0078\u0068\u003d\u0025\u0064\u0078\u0025\u0064\u000a\u0009\u006c\u006c\u0078\u003d\u0025\u002e\u0032\u0066\u000a\u0009\u006c\u006c\u0079\u003d\u0025\u002e\u0032f", _degf, _ecae, _bdba, _efga)
	}
	_cggd := make([]gridTile, _degf*_ecae)
	for _gdgc := _ecae - 1; _gdgc >= 0; _gdgc-- {
		_bgbebf := _efga[_gdgc]
		_dbdg := _efga[_gdgc+1]
		for _ggebe := 0; _ggebe < _degf; _ggebe++ {
			_ddaeg := _bdba[_ggebe]
			_efbc := _bdba[_ggebe+1]
			_aagf := _ggfa.findPrimSec(_ddaeg, _bgbebf)
			_ddffa := _ggfa.findPrimSec(_efbc, _bgbebf)
			_bbdg := _feeb.findPrimSec(_bgbebf, _ddaeg)
			_adeb := _feeb.findPrimSec(_dbdg, _ddaeg)
			_ddade := _gb.PdfRectangle{Llx: _ddaeg, Urx: _efbc, Lly: _bgbebf, Ury: _dbdg}
			_cffag := _fggg(_ddade, _aagf, _ddffa, _bbdg, _adeb)
			_cggd[_gdgc*_degf+_ggebe] = _cffag
			if _aaga {
				_cfe.Printf("\u0020\u0020\u0078\u003d\u0025\u0032\u0064\u0020\u0079\u003d\u0025\u0032\u0064\u003a\u0020%\u0073 \u0025\u0036\u002e\u0032\u0066\u0020\u0078\u0020\u0025\u0036\u002e\u0032\u0066\u000a", _ggebe, _gdgc, _cffag.String(), _cffag.Width(), _cffag.Height())
			}
		}
	}
	if _aaga {
		_ed.Log.Info("r\u0075\u006c\u0069\u006e\u0067\u004c\u0069\u0073\u0074.\u0061\u0073\u0054\u0069\u006c\u0069\u006eg:\u0020\u0063\u006f\u0061l\u0065\u0073\u0063\u0065\u0020\u0068\u006f\u0072\u0069zo\u006e\u0074a\u006c\u002e\u0020\u0025\u0036\u002e\u0032\u0066", _egce)
	}
	_deade := make([]map[float64]gridTile, _ecae)
	for _aabb := _ecae - 1; _aabb >= 0; _aabb-- {
		if _aaga {
			_cfe.Printf("\u0020\u0020\u0079\u003d\u0025\u0032\u0064\u000a", _aabb)
		}
		_deade[_aabb] = make(map[float64]gridTile, _degf)
		for _edad := 0; _edad < _degf; _edad++ {
			_daea := _cggd[_aabb*_degf+_edad]
			if _aaga {
				_cfe.Printf("\u0020\u0020\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _edad, _daea)
			}
			if !_daea._bdcfac {
				continue
			}
			_cddfdg := _edad
			for _dbgca := _edad + 1; !_daea._babe && _dbgca < _degf; _dbgca++ {
				_baae := _cggd[_aabb*_degf+_dbgca]
				_daea.Urx = _baae.Urx
				_daea._cccb = _daea._cccb || _baae._cccb
				_daea._debb = _daea._debb || _baae._debb
				_daea._babe = _baae._babe
				if _aaga {
					_cfe.Printf("\u0020 \u0020%\u0034\u0064\u003a\u0020\u0025s\u0020\u2192 \u0025\u0073\u000a", _dbgca, _baae, _daea)
				}
				_cddfdg = _dbgca
			}
			if _aaga {
				_cfe.Printf(" \u0020 \u0025\u0032\u0064\u0020\u002d\u0020\u0025\u0032d\u0020\u2192\u0020\u0025s\n", _edad, _cddfdg, _daea)
			}
			_edad = _cddfdg
			_deade[_aabb][_daea.Llx] = _daea
		}
	}
	_dgcad := make(map[float64]map[float64]gridTile, _ecae)
	_ggcad := make(map[float64]map[float64]struct{}, _ecae)
	for _dafdb := _ecae - 1; _dafdb >= 0; _dafdb-- {
		_edea := _cggd[_dafdb*_degf].Lly
		_dgcad[_edea] = make(map[float64]gridTile, _degf)
		_ggcad[_edea] = make(map[float64]struct{}, _degf)
	}
	if _aaga {
		_ed.Log.Info("\u0072u\u006c\u0069n\u0067\u004c\u0069s\u0074\u002e\u0061\u0073\u0054\u0069\u006ci\u006e\u0067\u003a\u0020\u0063\u006fa\u006c\u0065\u0073\u0063\u0065\u0020\u0076\u0065\u0072\u0074\u0069c\u0061\u006c\u002e\u0020\u0025\u0036\u002e\u0032\u0066", _egce)
	}
	for _gfdg := _ecae - 1; _gfdg >= 0; _gfdg-- {
		_gagdf := _cggd[_gfdg*_degf].Lly
		_ceaa := _deade[_gfdg]
		if _aaga {
			_cfe.Printf("\u0020\u0020\u0079\u003d\u0025\u0032\u0064\u000a", _gfdg)
		}
		for _, _acffd := range _aaffa(_ceaa) {
			if _, _cgff := _ggcad[_gagdf][_acffd]; _cgff {
				continue
			}
			_dgbgb := _ceaa[_acffd]
			if _aaga {
				_cfe.Printf(" \u0020\u0020\u0020\u0020\u0076\u0030\u003d\u0025\u0073\u000a", _dgbgb.String())
			}
			for _dbefc := _gfdg - 1; _dbefc >= 0; _dbefc-- {
				if _dgbgb._debb {
					break
				}
				_gcecg := _deade[_dbefc]
				_gdcgac, _fgdd := _gcecg[_acffd]
				if !_fgdd {
					break
				}
				if _gdcgac.Urx != _dgbgb.Urx {
					break
				}
				_dgbgb._debb = _gdcgac._debb
				_dgbgb.Lly = _gdcgac.Lly
				if _aaga {
					_cfe.Printf("\u0020\u0020\u0020\u0020  \u0020\u0020\u0076\u003d\u0025\u0073\u0020\u0076\u0030\u003d\u0025\u0073\u000a", _gdcgac.String(), _dgbgb.String())
				}
				_ggcad[_gdcgac.Lly][_gdcgac.Llx] = struct{}{}
			}
			if _gfdg == 0 {
				_dgbgb._debb = true
			}
			if _dgbgb.complete() {
				_dgcad[_gagdf][_acffd] = _dgbgb
			}
		}
	}
	_acfb := gridTiling{PdfRectangle: _egce, _edfg: _cgdeg(_dgcad), _abbf: _eccca(_dgcad), _aaef: _dgcad}
	_acfb.log("\u0043r\u0065\u0061\u0074\u0065\u0064")
	return _acfb
}
func (_ecgc *textLine) bbox() _gb.PdfRectangle { return _ecgc.PdfRectangle }
func _dcga(_bfeg *wordBag, _abce *textWord, _dddd float64) bool {
	return _abce.Llx < _bfeg.Urx+_dddd && _bfeg.Llx-_dddd < _abce.Urx
}
func _dcfc(_effb int, _bcdca map[int][]float64) ([]int, int) {
	_acba := make([]int, _effb)
	_fgab := 0
	for _febc := 0; _febc < _effb; _febc++ {
		_acba[_febc] = _fgab
		_fgab += len(_bcdca[_febc]) + 1
	}
	return _acba, _fgab
}

type textTable struct {
	_gb.PdfRectangle
	_fccgee, _dege int
	_baccf         bool
	_adda          map[uint64]*textPara
	_faeg          map[uint64]compositeCell
}

func _gggb(_fae _cc.Point) _cc.Matrix { return _cc.TranslationMatrix(_fae.X, _fae.Y) }

// ExtractText processes and extracts all text data in content streams and returns as a string.
// It takes into account character encodings in the PDF file, which are decoded by
// CharcodeBytesToUnicode.
// Characters that can't be decoded are replaced with MissingCodeRune ('\ufffd' = �).
func (_geeb *Extractor) ExtractText() (string, error) {
	_agg, _, _, _dbg := _geeb.ExtractTextWithStats()
	return _agg, _dbg
}
func (_bebed *textTable) reduce() *textTable {
	_eggec := make([]int, 0, _bebed._dege)
	_efead := make([]int, 0, _bebed._fccgee)
	for _feedd := 0; _feedd < _bebed._dege; _feedd++ {
		if !_bebed.emptyCompositeRow(_feedd) {
			_eggec = append(_eggec, _feedd)
		}
	}
	for _cdfd := 0; _cdfd < _bebed._fccgee; _cdfd++ {
		if !_bebed.emptyCompositeColumn(_cdfd) {
			_efead = append(_efead, _cdfd)
		}
	}
	if len(_eggec) == _bebed._dege && len(_efead) == _bebed._fccgee {
		return _bebed
	}
	_dbffe := textTable{_baccf: _bebed._baccf, _fccgee: len(_efead), _dege: len(_eggec), _adda: make(map[uint64]*textPara, len(_efead)*len(_eggec))}
	if _geff {
		_ed.Log.Info("\u0072\u0065\u0064\u0075ce\u003a\u0020\u0025\u0064\u0078\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0064\u0078%\u0064", _bebed._fccgee, _bebed._dege, len(_efead), len(_eggec))
		_ed.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0043\u006f\u006c\u0073\u003a\u0020\u0025\u002b\u0076", _efead)
		_ed.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0052\u006f\u0077\u0073\u003a\u0020\u0025\u002b\u0076", _eggec)
	}
	for _cbfce, _ceba := range _eggec {
		for _egfe, _ffafb := range _efead {
			_beade, _bbacf := _bebed.getComposite(_ffafb, _ceba)
			if _beade == nil {
				continue
			}
			if _geff {
				_cfe.Printf("\u0020 \u0025\u0032\u0064\u002c \u0025\u0032\u0064\u0020\u0028%\u0032d\u002c \u0025\u0032\u0064\u0029\u0020\u0025\u0071\n", _egfe, _cbfce, _ffafb, _ceba, _fabf(_beade.merge().text(), 50))
			}
			_dbffe.putComposite(_egfe, _cbfce, _beade, _bbacf)
		}
	}
	return &_dbffe
}
func (_cedab *textObject) getFontDirect(_bfe string) (*_gb.PdfFont, error) {
	_debe, _abff := _cedab.getFontDict(_bfe)
	if _abff != nil {
		return nil, _abff
	}
	_aaaf, _abff := _gb.NewPdfFontFromPdfObject(_debe)
	if _abff != nil {
		_ed.Log.Debug("\u0067\u0065\u0074\u0046\u006f\u006e\u0074\u0044\u0069\u0072\u0065\u0063\u0074\u003a\u0020\u004e\u0065\u0077Pd\u0066F\u006f\u006e\u0074\u0046\u0072\u006f\u006d\u0050\u0064\u0066\u004f\u0062j\u0065\u0063\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u006e\u0061\u006d\u0065\u003d%\u0023\u0071\u0020\u0065\u0072\u0072\u003d\u0025\u0076", _bfe, _abff)
	}
	return _aaaf, _abff
}
func (_faee *textLine) pullWord(_fadb *wordBag, _cefa *textWord, _caea int) {
	_faee.appendWord(_cefa)
	_fadb.removeWord(_cefa, _caea)
}
func (_abefg paraList) sortReadingOrder() {
	_ed.Log.Trace("\u0073\u006fr\u0074\u0052\u0065\u0061\u0064i\u006e\u0067\u004f\u0072\u0064e\u0072\u003a\u0020\u0070\u0061\u0072\u0061\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_abefg))
	if len(_abefg) <= 1 {
		return
	}
	_abefg.computeEBBoxes()
	_d.Slice(_abefg, func(_gcgdb, _bgcff int) bool { return _ecdf(_abefg[_gcgdb], _abefg[_bgcff]) <= 0 })
}
func (_faaag *textPara) fontsize() float64 { return _faaag._abbe[0]._egbd }

var _eebbd = map[markKind]string{_fdacf: "\u0073\u0074\u0072\u006f\u006b\u0065", _cbcd: "\u0066\u0069\u006c\u006c", _ecdcf: "\u0061u\u0067\u006d\u0065\u006e\u0074"}

func _gdcc(_bgda []structElement, _gcbb map[int][]*textLine, _fdae _gg.PdfObject) []*list {
	_bfge := []*list{}
	for _, _abcc := range _bgda {
		_ecec := _abcc._ddcfe
		_fcdff := int(_abcc._adfb)
		_deeea := _abcc._gacb
		_cgdeb := []*textLine{}
		_egbcg := []*list{}
		_aacd := _abcc._gfeff
		_eabc, _gbfee := (_aacd.(*_gg.PdfObjectReference))
		if !_gbfee {
			_ed.Log.Debug("\u0066\u0061\u0069l\u0065\u0064\u0020\u006f\u0074\u0020\u0063\u0061\u0073\u0074\u0020\u0074\u006f\u0020\u002a\u0063\u006f\u0072\u0065\u002e\u0050\u0064\u0066\u004f\u0062\u006a\u0065\u0063\u0074R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		}
		if _fcdff != -1 && _eabc != nil {
			if _abeb, _bgbb := _gcbb[_fcdff]; _bgbb {
				if _gegee, _cefc := _fdae.(*_gg.PdfIndirectObject); _cefc {
					_geeaa := _gegee.PdfObjectReference
					if _g.DeepEqual(*_eabc, _geeaa) {
						_cgdeb = _abeb
					}
				}
			}
		}
		if _ecec != nil {
			_egbcg = _gdcc(_ecec, _gcbb, _fdae)
		}
		_dcda := _gbcag(_cgdeb, _deeea, _egbcg)
		_bfge = append(_bfge, _dcda)
	}
	return _bfge
}

// ExtractPageText returns the text contents of `e` (an Extractor for a page) as a PageText.
// TODO(peterwilliams97): The stats complicate this function signature and aren't very useful.
//
//	Replace with a function like Extract() (*PageText, error)
func (_bde *Extractor) ExtractPageText() (*PageText, int, int, error) {
	_cg, _cgg, _eef, _fag := _bde.extractPageText(_bde._ee, _bde._bd, _cc.IdentityMatrix(), 0)
	if _fag != nil && _fag != _gb.ErrColorOutOfRange {
		return nil, 0, 0, _fag
	}
	if _bde._dfb != nil {
		_cg._ege._ebff = _bde._dfb.UseSimplerExtractionProcess
	}
	_cg.computeViews()
	if _bde._dfb != nil {
		if _bde._dfb.ApplyCropBox && _bde._gfg != nil {
			_cg.ApplyArea(*_bde._gfg)
		}
		_cg._ege._gdeb = _bde._dfb.DisableDocumentTags
	}
	return _cg, _cgg, _eef, nil
}
func (_gdd *subpath) close() {
	if !_ggadd(_gdd._ccgf[0], _gdd.last()) {
		_gdd.add(_gdd._ccgf[0])
	}
	_gdd._beac = true
	_gdd.removeDuplicates()
}
func (_adfbf *textTable) isExportable() bool {
	if _adfbf._baccf {
		return true
	}
	_eggea := func(_bdbcb int) bool {
		_fcafa := _adfbf.get(0, _bdbcb)
		if _fcafa == nil {
			return false
		}
		_cecb := _fcafa.text()
		_fbbg := _cb.RuneCountInString(_cecb)
		_beega := _dfbag.MatchString(_cecb)
		return _fbbg <= 1 || _beega
	}
	for _gccae := 0; _gccae < _adfbf._dege; _gccae++ {
		if !_eggea(_gccae) {
			return true
		}
	}
	return false
}
func _ccee(_gaff, _efeae *textPara) bool {
	if _gaff._befd || _efeae._befd {
		return true
	}
	return _cdgd(_gaff.depth() - _efeae.depth())
}
func (_efffa gridTile) complete() bool { return _efffa.numBorders() == 4 }
func (_affed *textPara) isAtom() *textTable {
	_eadba := _affed
	_dgab := _affed._ccdg
	_cegf := _affed._cgfe
	if _dgab.taken() || _cegf.taken() {
		return nil
	}
	_fbef := _dgab._cgfe
	if _fbef.taken() || _fbef != _cegf._ccdg {
		return nil
	}
	return _dfcee(_eadba, _dgab, _cegf, _fbef)
}

type wordBag struct {
	_gb.PdfRectangle
	_bcef        float64
	_cbfg, _gaeg rulingList
	_gebga       float64
	_gegc        map[int][]*textWord
}

// String returns a description of `k`.
func (_cgbecc rulingKind) String() string {
	_dccgd, _acec := _gegca[_cgbecc]
	if !_acec {
		return _cfe.Sprintf("\u004e\u006ft\u0020\u0061\u0020r\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0025\u0064", _cgbecc)
	}
	return _dccgd
}
func (_cddf *wordBag) text() string {
	_ddda := _cddf.allWords()
	_daefb := make([]string, len(_ddda))
	for _ffbd, _acb := range _ddda {
		_daefb[_ffbd] = _acb._bcfea
	}
	return _ag.Join(_daefb, "\u0020")
}

const (
	_ddfg  = true
	_gdggd = true
	_dbfc  = true
	_bcba  = false
	_cec   = false
	_dbae  = 6
	_dcba  = 3.0
	_dagb  = 200
	_ecbbf = true
	_cbaf  = true
	_gfaf  = true
	_bdacb = true
	_abdfd = false
)

func (_geegc intSet) has(_beadc int) bool { _, _ccfag := _geegc[_beadc]; return _ccfag }
func (_feaee *textWord) addDiacritic(_afbcd string) {
	_gdfg := _feaee._fdgbc[len(_feaee._fdgbc)-1]
	_gdfg._ceca += _afbcd
	_gdfg._ceca = _f.NFKC.String(_gdfg._ceca)
}
func _bdbf(_egef _gb.PdfRectangle) *ruling {
	return &ruling{_cffe: _gfafc, _abfg: _egef.Urx, _daag: _egef.Lly, _fcff: _egef.Ury}
}
func (_gcae *structTreeRoot) buildList(_aegg map[int][]*textLine, _ceff _gg.PdfObject) []*list {
	if _gcae == nil {
		_ed.Log.Debug("\u0062\u0075\u0069\u006c\u0064\u004c\u0069\u0073\u0074\u003a\u0020t\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0069\u0073 \u006e\u0069\u006c")
		return nil
	}
	var _bdaa *structElement
	_ddfbg := []structElement{}
	if len(_gcae._gafdd) == 1 {
		_egeg := _gcae._gafdd[0]._gacb
		if _egeg == "\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074" || _egeg == "\u0053\u0065\u0063\u0074" || _egeg == "\u0050\u0061\u0072\u0074" || _egeg == "\u0044\u0069\u0076" || _egeg == "\u0041\u0072\u0074" {
			_bdaa = &_gcae._gafdd[0]
		}
	} else {
		_bdaa = &structElement{_ddcfe: _gcae._gafdd, _gacb: _gcae._bgee}
	}
	if _bdaa == nil {
		_ed.Log.Debug("\u0062\u0075\u0069\u006cd\u004c\u0069\u0073\u0074\u003a\u0020\u0074\u006f\u0070\u0045l\u0065m\u0065\u006e\u0074\u0020\u0069\u0073\u0020n\u0069\u006c")
		return nil
	}
	for _, _ccdb := range _bdaa._ddcfe {
		if _ccdb._gacb == "\u004c" {
			_ddfbg = append(_ddfbg, _ccdb)
		} else if _ccdb._gacb == "\u0054\u0061\u0062l\u0065" {
			_fdeff := _bbfe(_ccdb)
			_ddfbg = append(_ddfbg, _fdeff...)
		}
	}
	_ccde := _gdcc(_ddfbg, _aegg, _ceff)
	var _caccc []*list
	for _, _afbecb := range _ccde {
		_geed := _fggcc(_afbecb)
		_caccc = append(_caccc, _geed...)
	}
	return _caccc
}
func (_gfge *TextMarkArray) getTextMarkAtOffset(_gef int) *TextMark {
	for _, _feaa := range _gfge._ggcf {
		if _feaa.Offset == _gef {
			return &_feaa
		}
	}
	return nil
}

// String returns a string describing `pt`.
func (_gbg PageText) String() string {
	_adbb := _cfe.Sprintf("P\u0061\u0067\u0065\u0054ex\u0074:\u0020\u0025\u0064\u0020\u0065l\u0065\u006d\u0065\u006e\u0074\u0073", len(_gbg._bagf))
	_egag := []string{"\u002d" + _adbb}
	for _, _cac := range _gbg._bagf {
		_egag = append(_egag, _cac.String())
	}
	_egag = append(_egag, "\u002b"+_adbb)
	return _ag.Join(_egag, "\u000a")
}

// String returns a string describing `tm`.
func (_bce TextMark) String() string {
	_aaffg := _bce.BBox
	var _eab string
	if _bce.Font != nil {
		_eab = _bce.Font.String()
		if len(_eab) > 50 {
			_eab = _eab[:50] + "\u002e\u002e\u002e"
		}
	}
	var _fggc string
	if _bce.Meta {
		_fggc = "\u0020\u002a\u004d\u002a"
	}
	return _cfe.Sprintf("\u007b\u0054\u0065\u0078t\u004d\u0061\u0072\u006b\u003a\u0020\u0025\u0064\u0020%\u0071\u003d\u0025\u0030\u0032\u0078\u0020\u0028\u0025\u0036\u002e\u0032\u0066\u002c\u0020\u0025\u0036\u002e2\u0066\u0029\u0020\u0028\u00256\u002e\u0032\u0066\u002c\u0020\u0025\u0036\u002e\u0032\u0066\u0029\u0020\u0025\u0073\u0025\u0073\u007d", _bce.Offset, _bce.Text, []rune(_bce.Text), _aaffg.Llx, _aaffg.Lly, _aaffg.Urx, _aaffg.Ury, _eab, _fggc)
}
func (_gcab *textObject) moveTextSetLeading(_gbc, _baad float64) {
	_gcab._cbag._fea = -_baad
	_gcab.moveLP(_gbc, _baad)
}

// String returns a human readable description of `s`.
func (_afacf intSet) String() string {
	var _ecdec []int
	for _gecf := range _afacf {
		if _afacf.has(_gecf) {
			_ecdec = append(_ecdec, _gecf)
		}
	}
	_d.Ints(_ecdec)
	return _cfe.Sprintf("\u0025\u002b\u0076", _ecdec)
}
func (_ecdd *ruling) alignsPrimary(_daed *ruling) bool {
	return _ecdd._cffe == _daed._cffe && _df.Abs(_ecdd._abfg-_daed._abfg) < _dbgc*0.5
}
func (_cfcb paraList) lines() []*textLine {
	var _bdcde []*textLine
	for _, _bbaa := range _cfcb {
		_bdcde = append(_bdcde, _bbaa._abbe...)
	}
	return _bdcde
}

func (_gffa *wordBag) scanBand(_daeb string, _cfed *wordBag, _cag func(_dgbg *wordBag, _fgecc *textWord) bool, _efge, _bbb, _egfg float64, _fdca, _cage bool) int {
	_cedb := _cfed._bcef
	var _cfebf map[int]map[*textWord]struct{}
	if !_fdca {
		_cfebf = _gffa.makeRemovals()
	}
	_dgfd := _dgfg * _cedb
	_cdd := 0
	for _, _gab := range _gffa.depthBand(_efge-_dgfd, _bbb+_dgfd) {
		if len(_gffa._gegc[_gab]) == 0 {
			continue
		}
		for _, _gega := range _gffa._gegc[_gab] {
			if !(_efge-_dgfd <= _gega._cbfcee && _gega._cbfcee <= _bbb+_dgfd) {
				continue
			}
			if !_cag(_cfed, _gega) {
				continue
			}
			_adge := 2.0 * _df.Abs(_gega._fgdbd-_cfed._bcef) / (_gega._fgdbd + _cfed._bcef)
			_edef := _df.Max(_gega._fgdbd/_cfed._bcef, _cfed._bcef/_gega._fgdbd)
			_bbba := _df.Min(_adge, _edef)
			if _egfg > 0 && _bbba > _egfg {
				continue
			}
			if _cfed.blocked(_gega) {
				continue
			}
			if !_fdca {
				_cfed.pullWord(_gega, _gab, _cfebf)
			}
			_cdd++
			if !_cage {
				if _gega._cbfcee < _efge {
					_efge = _gega._cbfcee
				}
				if _gega._cbfcee > _bbb {
					_bbb = _gega._cbfcee
				}
			}
			if _fdca {
				break
			}
		}
	}
	if !_fdca {
		_gffa.applyRemovals(_cfebf)
	}
	return _cdd
}
func _bdcec(_ggbg []pathSection) rulingList {
	_fdddc(_ggbg)
	if _geaeg {
		_ed.Log.Info("\u006da\u006b\u0065\u0046\u0069l\u006c\u0052\u0075\u006c\u0069n\u0067s\u003a \u0025\u0064\u0020\u0066\u0069\u006c\u006cs", len(_ggbg))
	}
	var _gbece rulingList
	for _, _acgea := range _ggbg {
		for _, _ebca := range _acgea._ggfc {
			if !_ebca.isQuadrilateral() {
				if _geaeg {
					_ed.Log.Error("!\u0069s\u0051\u0075\u0061\u0064\u0072\u0069\u006c\u0061t\u0065\u0072\u0061\u006c: \u0025\u0073", _ebca)
				}
				continue
			}
			if _ggfcd, _feggg := _ebca.makeRectRuling(_acgea.Color); _feggg {
				_gbece = append(_gbece, _ggfcd)
			} else {
				if _afbec {
					_ed.Log.Error("\u0021\u006d\u0061\u006beR\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0025\u0073", _ebca)
				}
			}
		}
	}
	if _geaeg {
		_ed.Log.Info("\u006d\u0061\u006b\u0065Fi\u006c\u006c\u0052\u0075\u006c\u0069\u006e\u0067\u0073\u003a\u0020\u0025\u0073", _gbece.String())
	}
	return _gbece
}
func (_ggf TextTable) getCellInfo(_gfgee TextMark) [][]int {
	for _beed, _bdca := range _ggf.Cells {
		for _fegb, _fcc := range _bdca {
			_acc := &_fcc.Marks
			if _acc.exists(_gfgee) {
				return [][]int{{_beed}, {_fegb}}
			}
		}
	}
	return nil
}
func (_dgbfa paraList) extractTables(_ceebb []gridTiling) paraList {
	if _geff {
		_ed.Log.Debug("\u0065\u0078\u0074r\u0061\u0063\u0074\u0054\u0061\u0062\u006c\u0065\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_dgbfa))
	}
	if len(_dgbfa) < _dgea {
		return _dgbfa
	}
	_daede := _dgbfa.findTables(_ceebb)
	if _geff {
		_ed.Log.Info("c\u006f\u006d\u0062\u0069\u006e\u0065d\u0020\u0074\u0061\u0062\u006c\u0065s\u0020\u0025\u0064\u0020\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d=\u003d", len(_daede))
		for _dbgg, _eeag := range _daede {
			_eeag.log(_cfe.Sprintf("c\u006f\u006d\u0062\u0069\u006e\u0065\u0064\u0020\u0025\u0064", _dbgg))
		}
	}
	return _dgbfa.applyTables(_daede)
}
func (_agbc *textPara) text() string {
	_dgedb := new(_cf.Buffer)
	_agbc.writeText(_dgedb)
	return _dgedb.String()
}
func _gaee(_adae *wordBag, _efca float64, _bege, _aegb rulingList) []*wordBag {
	var _cbfgf []*wordBag
	for _, _edcg := range _adae.depthIndexes() {
		_bcgag := false
		for !_adae.empty(_edcg) {
			_fcgb := _adae.firstReadingIndex(_edcg)
			_eagc := _adae.firstWord(_fcgb)
			_ebcc := _dfcc(_eagc, _efca, _bege, _aegb)
			_adae.removeWord(_eagc, _fcgb)
			if _fdec {
				_ed.Log.Info("\u0066\u0069\u0072\u0073\u0074\u0057\u006f\u0072\u0064\u0020\u005e\u005e^\u005e\u0020\u0025\u0073", _eagc.String())
			}
			for _baef := true; _baef; _baef = _bcgag {
				_bcgag = false
				_eeed := _cbfc * _ebcc._bcef
				_fadf := _debd * _ebcc._bcef
				_dgfe := _dccd * _ebcc._bcef
				if _fdec {
					_ed.Log.Info("\u0070a\u0072a\u0057\u006f\u0072\u0064\u0073\u0020\u0064\u0065\u0070\u0074\u0068 \u0025\u002e\u0032\u0066 \u002d\u0020\u0025\u002e\u0032f\u0020\u006d\u0061\u0078\u0049\u006e\u0074\u0072\u0061\u0044\u0065\u0070\u0074\u0068\u0047\u0061\u0070\u003d\u0025\u002e\u0032\u0066\u0020\u006d\u0061\u0078\u0049\u006e\u0074\u0072\u0061R\u0065\u0061\u0064\u0069\u006e\u0067\u0047\u0061p\u003d\u0025\u002e\u0032\u0066", _ebcc.minDepth(), _ebcc.maxDepth(), _dgfe, _fadf)
				}
				if _adae.scanBand("\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c", _ebcc, _gcda(_dcga, 0), _ebcc.minDepth()-_dgfe, _ebcc.maxDepth()+_dgfe, _aeg, false, false) > 0 {
					_bcgag = true
				}
				if _adae.scanBand("\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c", _ebcc, _gcda(_dcga, _fadf), _ebcc.minDepth(), _ebcc.maxDepth(), _cdcgd, false, false) > 0 {
					_bcgag = true
				}
				if _bcgag {
					continue
				}
				_gfcc := _adae.scanBand("", _ebcc, _gcda(_dbbg, _eeed), _ebcc.minDepth(), _ebcc.maxDepth(), _cdce, true, false)
				if _gfcc > 0 {
					_efcad := (_ebcc.maxDepth() - _ebcc.minDepth()) / _ebcc._bcef
					if (_gfcc > 1 && float64(_gfcc) > 0.3*_efcad) || _gfcc <= 10 {
						if _adae.scanBand("\u006f\u0074\u0068e\u0072", _ebcc, _gcda(_dbbg, _eeed), _ebcc.minDepth(), _ebcc.maxDepth(), _cdce, false, true) > 0 {
							_bcgag = true
						}
					}
				}
			}
			_cbfgf = append(_cbfgf, _ebcc)
		}
	}
	return _cbfgf
}

// String returns a string describing `ma`.
func (_deec TextMarkArray) String() string {
	_fbb := len(_deec._ggcf)
	if _fbb == 0 {
		return "\u0045\u004d\u0050T\u0059"
	}
	_dfbb := _deec._ggcf[0]
	_ecdc := _deec._ggcf[_fbb-1]
	return _cfe.Sprintf("\u007b\u0054\u0045\u0058\u0054\u004d\u0041\u0052K\u0041\u0052\u0052AY\u003a\u0020\u0025\u0064\u0020\u0065l\u0065\u006d\u0065\u006e\u0074\u0073\u000a\u0009\u0066\u0069\u0072\u0073\u0074\u003d\u0025s\u000a\u0009\u0020\u006c\u0061\u0073\u0074\u003d%\u0073\u007d", _fbb, _dfbb, _ecdc)
}

var _fdfgd string = "\u0028\u003f\u0069\u0029\u005e\u0028\u004d\u007b\u0030\u002c\u0033\u007d\u0029\u0028\u0043\u0028?\u003a\u0044\u007cM\u0029\u007c\u0044\u003f\u0043{\u0030\u002c\u0033\u007d\u0029\u0028\u0058\u0028\u003f\u003a\u004c\u007c\u0043\u0029\u007cL\u003f\u0058\u007b\u0030\u002c\u0033}\u0029\u0028\u0049\u0028\u003f\u003a\u0056\u007c\u0058\u0029\u007c\u0056\u003f\u0049\u007b\u0030\u002c\u0033\u007d\u0029\u0028\u005c\u0029\u007c\u005c\u002e\u0029\u007c\u005e\u005c\u0028\u0028\u004d\u007b\u0030\u002c\u0033\u007d\u0029\u0028\u0043\u0028\u003f\u003aD\u007cM\u0029\u007c\u0044\u003f\u0043\u007b\u0030\u002c\u0033\u007d\u0029\u0028\u0058\u0028?\u003a\u004c\u007c\u0043\u0029\u007c\u004c?\u0058\u007b0\u002c\u0033\u007d\u0029(\u0049\u0028\u003f\u003a\u0056|\u0058\u0029\u007c\u0056\u003f\u0049\u007b\u0030\u002c\u0033\u007d\u0029\u005c\u0029"

func (_aeb *shapesState) newSubPath() {
	_aeb.clearPath()
	if _dgeg {
		_ed.Log.Info("\u006e\u0065\u0077\u0053\u0075\u0062\u0050\u0061\u0074h\u003a\u0020\u0025\u0073", _aeb)
	}
}
func (_deag *shapesState) establishSubpath() *subpath {
	_fgfa, _cfcd := _deag.lastpointEstablished()
	if !_cfcd {
		_deag._babc = append(_deag._babc, _dcab(_fgfa))
	}
	if len(_deag._babc) == 0 {
		return nil
	}
	_deag._dfff = false
	return _deag._babc[len(_deag._babc)-1]
}

// Elements returns the TextMarks in `ma`.
func (_edff *TextMarkArray) Elements() []TextMark { return _edff._ggcf }
func (_bfbd *textLine) toTextMarks(_daaf *int) []TextMark {
	var _bgg []TextMark
	for _, _aggee := range _bfbd._bfca {
		if _aggee._cacca {
			_bgg = _dgccag(_bgg, _daaf, "\u0020")
		}
		_efef := _aggee.toTextMarks(_daaf)
		_bgg = append(_bgg, _efef...)
	}
	return _bgg
}
func _agbca(_faefb _gg.PdfObject, _cagd _eb.Color) (_gf.Image, error) {
	_fdbcd, _eaeef := _gg.GetStream(_faefb)
	if !_eaeef {
		return nil, nil
	}
	_ffgg, _bcbe := _gb.NewXObjectImageFromStream(_fdbcd)
	if _bcbe != nil {
		return nil, _bcbe
	}
	_adaa, _bcbe := _ffgg.ToImage()
	if _bcbe != nil {
		return nil, _bcbe
	}
	return _bdda(_adaa, _cagd), nil
}
func _caff(_cbfgfc, _gfeb _cc.Point, _afca _eb.Color) (*ruling, bool) {
	_ebaf := lineRuling{_efffg: _cbfgfc, _agae: _gfeb, _ecee: _cccea(_cbfgfc, _gfeb), Color: _afca}
	if _ebaf._ecee == _eefaf {
		return nil, false
	}
	return _ebaf.asRuling()
}

// TextMarkArray is a collection of TextMarks.
type TextMarkArray struct{ _ggcf []TextMark }

const (
	RenderModeStroke RenderMode = 1 << iota
	RenderModeFill
	RenderModeClip
)

func (_gbeeb *wordBag) allWords() []*textWord {
	var _eadf []*textWord
	for _, _bbec := range _gbeeb._gegc {
		_eadf = append(_eadf, _bbec...)
	}
	return _eadf
}
func (_efab paraList) list() []*list {
	var _edeg []*textLine
	var _gfbd []*textLine
	for _, _eaba := range _efab {
		_dcebg := _eaba.getListLines()
		_edeg = append(_edeg, _dcebg...)
		_gfbd = append(_gfbd, _eaba._abbe...)
	}
	_dgcg := _gbga(_edeg)
	_aaca := _cea(_gfbd, _dgcg)
	return _aaca
}
func (_faae *textObject) reset() {
	_faae._ffb = _cc.IdentityMatrix()
	_faae._ddge = _cc.IdentityMatrix()
	_faae._gafa = nil
}
func (_cfcagg *textWord) toTextMarks(_efgab *int) []TextMark {
	var _gfafe []TextMark
	for _, _gaddb := range _cfcagg._fdgbc {
		_gfafe = _gegag(_gfafe, _efgab, _gaddb.ToTextMark())
	}
	return _gfafe
}
func (_afe *shapesState) lineTo(_dbdd, _ecbb float64) {
	if _dgeg {
		_ed.Log.Info("\u006c\u0069\u006eeT\u006f\u0028\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0020\u0070\u003d\u0025\u002e\u0032\u0066", _dbdd, _ecbb, _afe.devicePoint(_dbdd, _ecbb))
	}
	_afe.addPoint(_dbdd, _ecbb)
}
func _abebe(_geagc []rulingList) (rulingList, rulingList) {
	var _edgc rulingList
	for _, _egga := range _geagc {
		_edgc = append(_edgc, _egga...)
	}
	return _edgc.vertsHorzs()
}
func (_abcg *textTable) emptyCompositeColumn(_bfdcag int) bool {
	for _bdbde := 0; _bdbde < _abcg._dege; _bdbde++ {
		if _dada, _ecbd := _abcg._faeg[_aafb(_bfdcag, _bdbde)]; _ecbd {
			if len(_dada.paraList) > 0 {
				return false
			}
		}
	}
	return true
}
func _gdfe(_gfgb _gb.PdfRectangle) *ruling {
	return &ruling{_cffe: _gfafc, _abfg: _gfgb.Llx, _daag: _gfgb.Lly, _fcff: _gfgb.Ury}
}
func _dcge(_bcbaf []TextMark, _ebddb *TextTable) []TextMark {
	var _cbfbe []TextMark
	for _, _ebccc := range _bcbaf {
		_ebccc._fefe = true
		_ebccc._egc = _ebddb
		_cbfbe = append(_cbfbe, _ebccc)
	}
	return _cbfbe
}

// Extractor stores and offers functionality for extracting content from PDF pages.
type Extractor struct {
	_ee  string
	_bd  *_gb.PdfPageResources
	_eda _gb.PdfRectangle
	_gfg *_gb.PdfRectangle
	_ebb map[string]fontEntry
	_db  map[string]textResult
	_bda int64
	_dg  int
	_dfb *Options
	_cd  *_gg.PdfObject
	_dfd _gg.PdfObject
}

func (_fecge paraList) readBefore(_agcb []int, _dcbb, _ddad int) bool {
	_ccbb, _begee := _fecge[_dcbb], _fecge[_ddad]
	if _cgfgb(_ccbb, _begee) && _ccbb.Lly > _begee.Lly {
		return true
	}
	if !(_ccbb._cedf.Urx < _begee._cedf.Llx) {
		return false
	}
	_gbfc, _gefa := _ccbb.Lly, _begee.Lly
	if _gbfc > _gefa {
		_gefa, _gbfc = _gbfc, _gefa
	}
	_abadc := _df.Max(_ccbb._cedf.Llx, _begee._cedf.Llx)
	_eebb := _df.Min(_ccbb._cedf.Urx, _begee._cedf.Urx)
	_gfggb := _fecge.llyRange(_agcb, _gbfc, _gefa)
	for _, _acbc := range _gfggb {
		if _acbc == _dcbb || _acbc == _ddad {
			continue
		}
		_ddba := _fecge[_acbc]
		if _ddba._cedf.Llx <= _eebb && _abadc <= _ddba._cedf.Urx {
			return false
		}
	}
	return true
}
func (_acdg *textMark) inDiacriticArea(_aecff *textMark) bool {
	_dged := _acdg.Llx - _aecff.Llx
	_bcdb := _acdg.Urx - _aecff.Urx
	_egcf := _acdg.Lly - _aecff.Lly
	return _df.Abs(_dged+_bcdb) < _acdg.Width()*_fabc && _df.Abs(_egcf) < _acdg.Height()*_fabc
}

// ToText returns the page text as a single string.
// Deprecated: This function is deprecated and will be removed in a future major version. Please use
// Text() instead.
func (_fede PageText) ToText() string { return _fede.Text() }
func _dgce(_dafea map[int][]float64) []int {
	_fafg := make([]int, len(_dafea))
	_dfeb := 0
	for _dceda := range _dafea {
		_fafg[_dfeb] = _dceda
		_dfeb++
	}
	_d.Ints(_fafg)
	return _fafg
}
func _ebgc(_abca, _eabf _gb.PdfRectangle) bool {
	return _abca.Lly <= _eabf.Ury && _eabf.Lly <= _abca.Ury
}

type textPara struct {
	_gb.PdfRectangle
	_cedf  _gb.PdfRectangle
	_abbe  []*textLine
	_bedfe *textTable
	_ggdd  bool
	_befd  bool
	_dgdb  *textPara
	_ccdg  *textPara
	_dbecf *textPara
	_cgfe  *textPara
	_fega  []list
}

func (_aab *imageExtractContext) extractContentStreamImages(_cbf string, _cfa *_gb.PdfPageResources) error {
	_aaf := _aad.NewContentStreamParser(_cbf)
	_fa, _feb := _aaf.Parse()
	if _feb != nil {
		return _feb
	}
	if _aab._eae == nil {
		_aab._eae = map[*_gg.PdfObjectStream]*cachedImage{}
	}
	if _aab._eag == nil {
		_aab._eag = &ImageExtractOptions{}
	}
	_agaf := _aad.NewContentStreamProcessor(*_fa)
	_agaf.AddHandler(_aad.HandlerConditionEnumAllOperands, "", _aab.processOperand)
	return _agaf.Process(_cfa)
}
func (_dcaca *textTable) getRight() paraList {
	_bbbab := make(paraList, _dcaca._dege)
	for _eaed := 0; _eaed < _dcaca._dege; _eaed++ {
		_fefb := _dcaca.get(_dcaca._fccgee-1, _eaed)._ccdg
		if _fefb.taken() {
			return nil
		}
		_bbbab[_eaed] = _fefb
	}
	for _dgadf := 0; _dgadf < _dcaca._dege-1; _dgadf++ {
		if _bbbab[_dgadf]._cgfe != _bbbab[_dgadf+1] {
			return nil
		}
	}
	return _bbbab
}
func (_afaa rulingList) sortStrict() {
	_d.Slice(_afaa, func(_dfcfc, _cgfb int) bool {
		_cdfg, _abgbbb := _afaa[_dfcfc], _afaa[_cgfb]
		_fcdb, _feacf := _cdfg._cffe, _abgbbb._cffe
		if _fcdb != _feacf {
			return _fcdb > _feacf
		}
		_cebd, _ddfgc := _cdfg._abfg, _abgbbb._abfg
		if !_cdgd(_cebd - _ddfgc) {
			return _cebd < _ddfgc
		}
		_cebd, _ddfgc = _cdfg._daag, _abgbbb._daag
		if _cebd != _ddfgc {
			return _cebd < _ddfgc
		}
		return _cdfg._fcff < _abgbbb._fcff
	})
}
func (_dbfe paraList) toTextMarks() []TextMark {
	_bbfd := 0
	var _dbbge []TextMark
	for _cgee, _bgba := range _dbfe {
		if _bgba._befd {
			continue
		}
		_beeee := _bgba.toTextMarks(&_bbfd)
		_dbbge = append(_dbbge, _beeee...)
		if _cgee != len(_dbfe)-1 {
			if _ccee(_bgba, _dbfe[_cgee+1]) {
				_dbbge = _dgccag(_dbbge, &_bbfd, "\u0020")
			} else {
				_dbbge = _dgccag(_dbbge, &_bbfd, "\u000a")
				_dbbge = _dgccag(_dbbge, &_bbfd, "\u000a")
			}
		}
	}
	_dbbge = _dgccag(_dbbge, &_bbfd, "\u000a")
	_dbbge = _dgccag(_dbbge, &_bbfd, "\u000a")
	return _dbbge
}
func (_gaab rulingList) aligned() bool {
	if len(_gaab) < 2 {
		return false
	}
	_bcbgg := make(map[*ruling]int)
	_bcbgg[_gaab[0]] = 0
	for _, _gcdc := range _gaab[1:] {
		_cgab := false
		for _fcag := range _bcbgg {
			if _gcdc.gridIntersecting(_fcag) {
				_bcbgg[_fcag]++
				_cgab = true
				break
			}
		}
		if !_cgab {
			_bcbgg[_gcdc] = 0
		}
	}
	_edfec := 0
	for _, _bbea := range _bcbgg {
		if _bbea == 0 {
			_edfec++
		}
	}
	_bgbad := float64(_edfec) / float64(len(_gaab))
	_dac := _bgbad <= 1.0-_ecdaa
	if _geaeg {
		_ed.Log.Info("\u0061\u006c\u0069\u0067\u006e\u0065\u0064\u003d\u0025\u0074\u0020\u0075\u006em\u0061\u0074\u0063\u0068\u0065\u0064=\u0025\u002e\u0032\u0066\u003d\u0025\u0064\u002f\u0025\u0064\u0020\u0076\u0065c\u0073\u003d\u0025\u0073", _dac, _bgbad, _edfec, len(_gaab), _gaab.String())
	}
	return _dac
}
func _aege(_ccggb []TextMark, _ggag *int) []TextMark {
	_eegb := _ccggb[len(_ccggb)-1]
	_egac := []rune(_eegb.Text)
	if len(_egac) == 1 {
		_ccggb = _ccggb[:len(_ccggb)-1]
		_cgbag := _ccggb[len(_ccggb)-1]
		*_ggag = _cgbag.Offset + len(_cgbag.Text)
	} else {
		_gceb := _aabe(_eegb.Text)
		*_ggag += len(_gceb) - len(_eegb.Text)
		_eegb.Text = _gceb
	}
	return _ccggb
}
func (_bcdgd *textPara) toTextMarks(_fgbe *int) []TextMark {
	if _bcdgd._bedfe == nil {
		return _bcdgd.toCellTextMarks(_fgbe)
	}
	var _fedg []TextMark
	for _gdcga := 0; _gdcga < _bcdgd._bedfe._dege; _gdcga++ {
		for _gcddg := 0; _gcddg < _bcdgd._bedfe._fccgee; _gcddg++ {
			_badf := _bcdgd._bedfe.get(_gcddg, _gdcga)
			if _badf == nil {
				_fedg = _dgccag(_fedg, _fgbe, "\u0009")
			} else {
				_bcgb := _badf.toCellTextMarks(_fgbe)
				_fedg = append(_fedg, _bcgb...)
			}
			_fedg = _dgccag(_fedg, _fgbe, "\u0020")
		}
		if _gdcga < _bcdgd._bedfe._dege-1 {
			_fedg = _dgccag(_fedg, _fgbe, "\u000a")
		}
	}
	_cfcdg := _bcdgd._bedfe
	if _cfcdg.isExportable() {
		_gfgc := _cfcdg.toTextTable()
		_fedg = _dcge(_fedg, &_gfgc)
	}
	return _fedg
}
func _cdgd(_ggac float64) bool { return _df.Abs(_ggac) < _dafe }

// String returns a human readable description of `ss`.
func (_bgce *shapesState) String() string {
	return _cfe.Sprintf("\u007b\u0025\u0064\u0020su\u0062\u0070\u0061\u0074\u0068\u0073\u0020\u0066\u0072\u0065\u0073\u0068\u003d\u0025t\u007d", len(_bgce._babc), _bgce._dfff)
}
func (_bdeg paraList) applyTables(_edfdg []*textTable) paraList {
	var _gcdg paraList
	for _, _eafac := range _edfdg {
		_gcdg = append(_gcdg, _eafac.newTablePara())
	}
	for _, _cfca := range _bdeg {
		if _cfca._ggdd {
			continue
		}
		_gcdg = append(_gcdg, _cfca)
	}
	return _gcdg
}
func (_gebb *imageExtractContext) extractFormImages(_dceb *_gg.PdfObjectName, _eaee _aad.GraphicsState, _gege *_gb.PdfPageResources) error {
	_faf, _aaff := _gege.GetXObjectFormByName(*_dceb)
	if _aaff != nil {
		return _aaff
	}
	if _faf == nil {
		return nil
	}
	_cae, _aaff := _faf.GetContentStream()
	if _aaff != nil {
		return _aaff
	}
	_ccf := _faf.Resources
	if _ccf == nil {
		_ccf = _gege
	}
	_aaff = _gebb.extractContentStreamImages(string(_cae), _ccf)
	if _aaff != nil {
		return _aaff
	}
	_gebb._ebf++
	return nil
}

// String returns a string describing the current state of the textState stack.
func (_gece *stateStack) String() string {
	_cgb := []string{_cfe.Sprintf("\u002d\u002d\u002d\u002d f\u006f\u006e\u0074\u0020\u0073\u0074\u0061\u0063\u006b\u003a\u0020\u0025\u0064", len(*_gece))}
	for _dba, _age := range *_gece {
		_ged := "\u003c\u006e\u0069l\u003e"
		if _age != nil {
			_ged = _age.String()
		}
		_cgb = append(_cgb, _cfe.Sprintf("\u0009\u0025\u0032\u0064\u003a\u0020\u0025\u0073", _dba, _ged))
	}
	return _ag.Join(_cgb, "\u000a")
}
func (_dbeb *ruling) encloses(_fagc, _cfadb float64) bool {
	return _dbeb._daag-_egaa <= _fagc && _cfadb <= _dbeb._fcff+_egaa
}

// ImageMark represents an image drawn on a page and its position in device coordinates.
// All coordinates are in device coordinates.
type ImageMark struct {
	Image *_gb.Image

	// Dimensions of the image as displayed in the PDF.
	Width  float64
	Height float64

	// Position of the image in PDF coordinates (lower left corner).
	X float64
	Y float64

	// Angle in degrees, if rotated.
	Angle float64
}

func _befg(_adcdc *list, _gcdd *_ag.Builder, _gbff *string) {
	_fbcf := _efeb(_adcdc, _gbff)
	_gcdd.WriteString(_fbcf)
	for _, _ffce := range _adcdc._gbcab {
		_gcec := *_gbff + "\u0020\u0020\u0020"
		_befg(_ffce, _gcdd, &_gcec)
	}
}

var (
	_ccb = _c.New("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072")
	_be  = _c.New("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072")
)

type textWord struct {
	_gb.PdfRectangle
	_cbfcee float64
	_bcfea  string
	_fdgbc  []*textMark
	_fgdbd  float64
	_cacca  bool
}

func (_eddea rulingList) bbox() _gb.PdfRectangle {
	var _gbcbf _gb.PdfRectangle
	if len(_eddea) == 0 {
		_ed.Log.Error("r\u0075\u006c\u0069\u006e\u0067\u004ci\u0073\u0074\u002e\u0062\u0062\u006f\u0078\u003a\u0020n\u006f\u0020\u0072u\u006ci\u006e\u0067\u0073")
		return _gb.PdfRectangle{}
	}
	if _eddea[0]._cffe == _bdfd {
		_gbcbf.Llx, _gbcbf.Urx = _eddea.secMinMax()
		_gbcbf.Lly, _gbcbf.Ury = _eddea.primMinMax()
	} else {
		_gbcbf.Llx, _gbcbf.Urx = _eddea.primMinMax()
		_gbcbf.Lly, _gbcbf.Ury = _eddea.secMinMax()
	}
	return _gbcbf
}
func (_gbebg *shapesState) fill(_dbcd *[]pathSection) {
	_ffeg := pathSection{_ggfc: _gbebg._babc, Color: _gbebg._geba.getFillColor()}
	*_dbcd = append(*_dbcd, _ffeg)
	if _geaeg {
		_fdfc := _ffeg.bbox()
		_cfe.Printf("\u0020 \u0020\u0020\u0046\u0049\u004c\u004c\u003a %\u0032\u0064\u0020\u0066\u0069\u006c\u006c\u0073\u0020\u0028\u0025\u0064\u0020\u006ee\u0077\u0029 \u0073\u0073\u003d%\u0073\u0020\u0063\u006f\u006c\u006f\u0072\u003d\u0025\u0033\u0076\u0020\u0025\u0036\u002e\u0032f\u003d\u00256.\u0032\u0066\u0078%\u0036\u002e\u0032\u0066\u000a", len(*_dbcd), len(_ffeg._ggfc), _gbebg, _ffeg.Color, _fdfc, _fdfc.Width(), _fdfc.Height())
		if _faab {
			for _gbf, _fdfd := range _ffeg._ggfc {
				_cfe.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _gbf, _fdfd)
				if _gbf == 10 {
					break
				}
			}
		}
	}
}
func (_cbfa *wordBag) blocked(_bca *textWord) bool {
	if _bca.Urx < _cbfa.Llx {
		_agead := _bdbf(_bca.PdfRectangle)
		_fgge := _gdfe(_cbfa.PdfRectangle)
		if _cbfa._cbfg.blocks(_agead, _fgge) {
			if _aagad {
				_ed.Log.Info("\u0062\u006c\u006f\u0063ke\u0064\u0020\u2190\u0078\u003a\u0020\u0025\u0073\u0020\u0025\u0073", _bca, _cbfa)
			}
			return true
		}
	} else if _cbfa.Urx < _bca.Llx {
		_dfgd := _bdbf(_cbfa.PdfRectangle)
		_efbf := _gdfe(_bca.PdfRectangle)
		if _cbfa._cbfg.blocks(_dfgd, _efbf) {
			if _aagad {
				_ed.Log.Info("b\u006co\u0063\u006b\u0065\u0064\u0020\u0078\u2192\u0020:\u0020\u0025\u0073\u0020%s", _bca, _cbfa)
			}
			return true
		}
	}
	if _bca.Ury < _cbfa.Lly {
		_cee := _cbcc(_bca.PdfRectangle)
		_abfc := _bdd(_cbfa.PdfRectangle)
		if _cbfa._gaeg.blocks(_cee, _abfc) {
			if _aagad {
				_ed.Log.Info("\u0062\u006c\u006f\u0063ke\u0064\u0020\u2190\u0079\u003a\u0020\u0025\u0073\u0020\u0025\u0073", _bca, _cbfa)
			}
			return true
		}
	} else if _cbfa.Ury < _bca.Lly {
		_bccce := _cbcc(_cbfa.PdfRectangle)
		_gebe := _bdd(_bca.PdfRectangle)
		if _cbfa._gaeg.blocks(_bccce, _gebe) {
			if _aagad {
				_ed.Log.Info("b\u006co\u0063\u006b\u0065\u0064\u0020\u0079\u2192\u0020:\u0020\u0025\u0073\u0020%s", _bca, _cbfa)
			}
			return true
		}
	}
	return false
}
func (_gcf *shapesState) drawRectangle(_feaaa, _aabf, _afbe, _bge float64) {
	if _dgeg {
		_ecea := _gcf.devicePoint(_feaaa, _aabf)
		_ffbg := _gcf.devicePoint(_feaaa+_afbe, _aabf+_bge)
		_edgb := _gb.PdfRectangle{Llx: _ecea.X, Lly: _ecea.Y, Urx: _ffbg.X, Ury: _ffbg.Y}
		_ed.Log.Info("d\u0072a\u0077\u0052\u0065\u0063\u0074\u0061\u006e\u0067l\u0065\u003a\u0020\u00256.\u0032\u0066", _edgb)
	}
	_gcf.newSubPath()
	_gcf.moveTo(_feaaa, _aabf)
	_gcf.lineTo(_feaaa+_afbe, _aabf)
	_gcf.lineTo(_feaaa+_afbe, _aabf+_bge)
	_gcf.lineTo(_feaaa, _aabf+_bge)
	_gcf.closePath()
}

// String returns a description of `w`.
func (_gabdf *textWord) String() string {
	return _cfe.Sprintf("\u0025\u002e2\u0066\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\"%\u0073\u0022", _gabdf._cbfcee, _gabdf.PdfRectangle, _gabdf._fgdbd, _gabdf._bcfea)
}
func _cea(_ebfad []*textLine, _efaed map[float64][]*textLine) []*list {
	_dgfb := _fdgd(_efaed)
	_ccfdd := []*list{}
	if len(_dgfb) == 0 {
		return _ccfdd
	}
	_babg := _dgfb[0]
	_gdce := 1
	_dgcaf := _efaed[_babg]
	for _cdcgb, _abaa := range _dgcaf {
		var _cgac float64
		_badg := []*list{}
		_efcg := _abaa._bdbg
		_ecfg := -1.0
		if _cdcgb < len(_dgcaf)-1 {
			_ecfg = _dgcaf[_cdcgb+1]._bdbg
		}
		if _gdce < len(_dgfb) {
			_badg = _eddb(_ebfad, _efaed, _dgfb, _gdce, _efcg, _ecfg)
		}
		_cgac = _ecfg
		if len(_badg) > 0 {
			_aagg := _badg[0]
			if len(_aagg._ggbee) > 0 {
				_cgac = _aagg._ggbee[0]._bdbg
			}
		}
		_fbfg := []*textLine{_abaa}
		_dfaf := _bbfg(_abaa, _ebfad, _dgfb, _efcg, _cgac)
		_fbfg = append(_fbfg, _dfaf...)
		_bbde := _gbcag(_fbfg, "\u0062\u0075\u006c\u006c\u0065\u0074", _badg)
		_bbde._ebabg = _adag(_fbfg, "")
		_ccfdd = append(_ccfdd, _bbde)
	}
	return _ccfdd
}

type gridTiling struct {
	_gb.PdfRectangle
	_edfg []float64
	_abbf []float64
	_aaef map[float64]map[float64]gridTile
}

func _bdd(_fcadf _gb.PdfRectangle) *ruling {
	return &ruling{_cffe: _bdfd, _abfg: _fcadf.Lly, _daag: _fcadf.Llx, _fcff: _fcadf.Urx}
}

type textObject struct {
	_cgf   *Extractor
	_edacg *_gb.PdfPageResources
	_ccg   _aad.GraphicsState
	_cbag  *textState
	_bab   *stateStack
	_ffb   _cc.Matrix
	_ddge  _cc.Matrix
	_gafa  []*textMark
	_efcbd bool
}

func _cgdeg(_gecaa map[float64]map[float64]gridTile) []float64 {
	_caccd := make([]float64, 0, len(_gecaa))
	_aadc := make(map[float64]struct{}, len(_gecaa))
	for _, _eaec := range _gecaa {
		for _fbdb := range _eaec {
			if _, _adcee := _aadc[_fbdb]; _adcee {
				continue
			}
			_caccd = append(_caccd, _fbdb)
			_aadc[_fbdb] = struct{}{}
		}
	}
	_d.Float64s(_caccd)
	return _caccd
}

// PageImages represents extracted images on a PDF page with spatial information:
// display position and size.
type PageImages struct{ Images []ImageMark }

func (_cdcc *stateStack) pop() *textState {
	if _cdcc.empty() {
		return nil
	}
	_ddde := *(*_cdcc)[len(*_cdcc)-1]
	*_cdcc = (*_cdcc)[:len(*_cdcc)-1]
	return &_ddde
}
func (_efbfe paraList) writeText(_cbdc _a.Writer) {
	for _abad, _babd := range _efbfe {
		if _babd._befd {
			continue
		}
		_babd.writeText(_cbdc)
		if _abad != len(_efbfe)-1 {
			if _ccee(_babd, _efbfe[_abad+1]) {
				_cbdc.Write([]byte("\u0020"))
			} else {
				_cbdc.Write([]byte("\u000a"))
				_cbdc.Write([]byte("\u000a"))
			}
		}
	}
	_cbdc.Write([]byte("\u000a"))
	_cbdc.Write([]byte("\u000a"))
}
func _bfddb(_bbga float64) float64 { return _bcdg * _df.Round(_bbga/_bcdg) }
func (_geag pathSection) bbox() _gb.PdfRectangle {
	_acf := _geag._ggfc[0]._ccgf[0]
	_geae := _gb.PdfRectangle{Llx: _acf.X, Urx: _acf.X, Lly: _acf.Y, Ury: _acf.Y}
	_gccd := func(_dcag _cc.Point) {
		if _dcag.X < _geae.Llx {
			_geae.Llx = _dcag.X
		} else if _dcag.X > _geae.Urx {
			_geae.Urx = _dcag.X
		}
		if _dcag.Y < _geae.Lly {
			_geae.Lly = _dcag.Y
		} else if _dcag.Y > _geae.Ury {
			_geae.Ury = _dcag.Y
		}
	}
	for _, _edee := range _geag._ggfc[0]._ccgf[1:] {
		_gccd(_edee)
	}
	for _, _degb := range _geag._ggfc[1:] {
		for _, _afba := range _degb._ccgf {
			_gccd(_afba)
		}
	}
	return _geae
}
func (_cfb *wordBag) absorb(_bgeb *wordBag) {
	_eaca := _bgeb.makeRemovals()
	for _ddgg, _fab := range _bgeb._gegc {
		for _, _ebab := range _fab {
			_cfb.pullWord(_ebab, _ddgg, _eaca)
		}
	}
	_bgeb.applyRemovals(_eaca)
}
func (_cdfbd *textTable) growTable() {
	_fbbf := func(_bagc paraList) {
		_cdfbd._dege++
		for _bdgg := 0; _bdgg < _cdfbd._fccgee; _bdgg++ {
			_acad := _bagc[_bdgg]
			_cdfbd.put(_bdgg, _cdfbd._dege-1, _acad)
		}
	}
	_fecgf := func(_gbacf paraList) {
		_cdfbd._fccgee++
		for _bfdg := 0; _bfdg < _cdfbd._dege; _bfdg++ {
			_baaec := _gbacf[_bfdg]
			_cdfbd.put(_cdfbd._fccgee-1, _bfdg, _baaec)
		}
	}
	if _abef {
		_cdfbd.log("\u0067r\u006f\u0077\u0054\u0061\u0062\u006ce")
	}
	for _gbcf := 0; ; _gbcf++ {
		_bdbd := false
		_dadgf := _cdfbd.getDown()
		_eggf := _cdfbd.getRight()
		if _abef {
			_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _gbcf, _cdfbd)
			_cfe.Printf("\u0020\u0020 \u0020\u0020\u0020 \u0020\u0064\u006f\u0077\u006e\u003d\u0025\u0073\u000a", _dadgf)
			_cfe.Printf("\u0020\u0020 \u0020\u0020\u0020 \u0072\u0069\u0067\u0068\u0074\u003d\u0025\u0073\u000a", _eggf)
		}
		if _dadgf != nil && _eggf != nil {
			_ccaa := _dadgf[len(_dadgf)-1]
			if !_ccaa.taken() && _ccaa == _eggf[len(_eggf)-1] {
				_fbbf(_dadgf)
				if _eggf = _cdfbd.getRight(); _eggf != nil {
					_fecgf(_eggf)
					_cdfbd.put(_cdfbd._fccgee-1, _cdfbd._dege-1, _ccaa)
				}
				_bdbd = true
			}
		}
		if !_bdbd && _dadgf != nil {
			_fbbf(_dadgf)
			_bdbd = true
		}
		if !_bdbd && _eggf != nil {
			_fecgf(_eggf)
			_bdbd = true
		}
		if !_bdbd {
			break
		}
	}
}

type markKind int

func (_ggba *textTable) emptyCompositeRow(_ageag int) bool {
	for _cdcb := 0; _cdcb < _ggba._fccgee; _cdcb++ {
		if _eaeec, _acbde := _ggba._faeg[_aafb(_cdcb, _ageag)]; _acbde {
			if len(_eaeec.paraList) > 0 {
				return false
			}
		}
	}
	return true
}
func (_bbgg compositeCell) hasLines(_gcbg []*textLine) bool {
	for _dcdaf, _dagd := range _gcbg {
		_cefb := _feee(_bbgg.PdfRectangle, _dagd.PdfRectangle)
		if _geff {
			_cfe.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u005e\u005e\u005e\u0069\u006e\u0074\u0065\u0072\u0073e\u0063t\u0073\u003d\u0025\u0074\u0020\u0025\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u000a", _cefb, _dcdaf, len(_gcbg))
			_cfe.Printf("\u0020\u0020\u0020\u0020  \u005e\u005e\u005e\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u003d\u0025s\u000a", _bbgg)
			_cfe.Printf("\u0020 \u0020 \u0020\u0020\u0020\u006c\u0069\u006e\u0065\u003d\u0025\u0073\u000a", _dagd)
		}
		if _cefb {
			return true
		}
	}
	return false
}

// ExtractTextWithStats works like ExtractText but returns the number of characters in the output
// (`numChars`) and the number of characters that were not decoded (`numMisses`).
func (_abg *Extractor) ExtractTextWithStats() (_ddee string, _fed int, _ffe int, _fc error) {
	_fge, _fed, _ffe, _fc := _abg.ExtractPageText()
	if _fc != nil {
		return "", _fed, _ffe, _fc
	}
	return _fge.Text(), _fed, _ffe, nil
}
func _efeb(_acdd *list, _fgcb *string) string {
	_edagc := _ag.Split(_acdd._ebabg, "\u000a")
	_cddfd := &_ag.Builder{}
	for _, _gcdb := range _edagc {
		if _gcdb != "" {
			_cddfd.WriteString(*_fgcb)
			_cddfd.WriteString(_gcdb)
			_cddfd.WriteString("\u000a")
		}
	}
	return _cddfd.String()
}
func (_fbdg paraList) tables() []TextTable {
	var _cfcf []TextTable
	if _geff {
		_ed.Log.Info("\u0070\u0061\u0072\u0061\u0073\u002e\u0074\u0061\u0062\u006c\u0065\u0073\u003a")
	}
	for _, _fdgc := range _fbdg {
		_fedc := _fdgc._bedfe
		if _fedc != nil && _fedc.isExportable() {
			_cfcf = append(_cfcf, _fedc.toTextTable())
		}
	}
	return _cfcf
}
func _feed(_eceg []*textLine, _bdec, _cde float64) []*textLine {
	var _dbe []*textLine
	for _, _fcbb := range _eceg {
		if _bdec == -1 {
			if _fcbb._bdbg > _cde {
				_dbe = append(_dbe, _fcbb)
			}
		} else {
			if _fcbb._bdbg > _cde && _fcbb._bdbg < _bdec {
				_dbe = append(_dbe, _fcbb)
			}
		}
	}
	return _dbe
}
func (_bef *wordBag) depthIndexes() []int {
	if len(_bef._gegc) == 0 {
		return nil
	}
	_gfgd := make([]int, len(_bef._gegc))
	_aedf := 0
	for _degbg := range _bef._gegc {
		_gfgd[_aedf] = _degbg
		_aedf++
	}
	_d.Ints(_gfgd)
	return _gfgd
}
func _dgccag(_ccda []TextMark, _cfaa *int, _afdbb string) []TextMark {
	_ecegb := _gdcf
	_ecegb.Text = _afdbb
	return _gegag(_ccda, _cfaa, _ecegb)
}
func (_fgda *textTable) markCells() {
	for _fdcd := 0; _fdcd < _fgda._dege; _fdcd++ {
		for _afdfc := 0; _afdfc < _fgda._fccgee; _afdfc++ {
			_gaeb := _fgda.get(_afdfc, _fdcd)
			if _gaeb != nil {
				_gaeb._ggdd = true
			}
		}
	}
}
func (_edgde *textTable) depth() float64 {
	_fagfd := 1e10
	for _adcbe := 0; _adcbe < _edgde._fccgee; _adcbe++ {
		_cgfde := _edgde.get(_adcbe, 0)
		if _cgfde == nil || _cgfde._befd {
			continue
		}
		_fagfd = _df.Min(_fagfd, _cgfde.depth())
	}
	return _fagfd
}
func (_cdad *textTable) reduceTiling(_ebbgd gridTiling, _gbefb float64) *textTable {
	_dgaba := make([]int, 0, _cdad._dege)
	_cdcgc := make([]int, 0, _cdad._fccgee)
	_ageaf := _ebbgd._edfg
	_egae := _ebbgd._abbf
	for _ccbf := 0; _ccbf < _cdad._dege; _ccbf++ {
		_ebef := _ccbf > 0 && _df.Abs(_egae[_ccbf-1]-_egae[_ccbf]) < _gbefb && _cdad.emptyCompositeRow(_ccbf)
		if !_ebef {
			_dgaba = append(_dgaba, _ccbf)
		}
	}
	for _dgbgg := 0; _dgbgg < _cdad._fccgee; _dgbgg++ {
		_gcdbe := _dgbgg < _cdad._fccgee-1 && _df.Abs(_ageaf[_dgbgg+1]-_ageaf[_dgbgg]) < _gbefb && _cdad.emptyCompositeColumn(_dgbgg)
		if !_gcdbe {
			_cdcgc = append(_cdcgc, _dgbgg)
		}
	}
	if len(_dgaba) == _cdad._dege && len(_cdcgc) == _cdad._fccgee {
		return _cdad
	}
	_gebda := textTable{_baccf: _cdad._baccf, _fccgee: len(_cdcgc), _dege: len(_dgaba), _faeg: make(map[uint64]compositeCell, len(_cdcgc)*len(_dgaba))}
	if _geff {
		_ed.Log.Info("\u0072\u0065\u0064\u0075c\u0065\u0054\u0069\u006c\u0069\u006e\u0067\u003a\u0020\u0025d\u0078%\u0064\u0020\u002d\u003e\u0020\u0025\u0064x\u0025\u0064", _cdad._fccgee, _cdad._dege, len(_cdcgc), len(_dgaba))
		_ed.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0043\u006f\u006c\u0073\u003a\u0020\u0025\u002b\u0076", _cdcgc)
		_ed.Log.Info("\u0072\u0065d\u0075\u0063\u0065d\u0052\u006f\u0077\u0073\u003a\u0020\u0025\u002b\u0076", _dgaba)
	}
	for _gcbe, _ddbaab := range _dgaba {
		for _bcbag, _gbaa := range _cdcgc {
			_beab, _bbdcd := _cdad.getComposite(_gbaa, _ddbaab)
			if len(_beab) == 0 {
				continue
			}
			if _geff {
				_cfe.Printf("\u0020 \u0025\u0032\u0064\u002c \u0025\u0032\u0064\u0020\u0028%\u0032d\u002c \u0025\u0032\u0064\u0029\u0020\u0025\u0071\n", _bcbag, _gcbe, _gbaa, _ddbaab, _fabf(_beab.merge().text(), 50))
			}
			_gebda.putComposite(_bcbag, _gcbe, _beab, _bbdcd)
		}
	}
	return &_gebda
}
func (_bggd *textTable) compositeColCorridors() map[int][]float64 {
	_gadg := make(map[int][]float64, _bggd._fccgee)
	if _geff {
		_ed.Log.Info("\u0063\u006f\u006d\u0070o\u0073\u0069\u0074\u0065\u0043\u006f\u006c\u0043\u006f\u0072r\u0069d\u006f\u0072\u0073\u003a\u0020\u0077\u003d%\u0064\u0020", _bggd._fccgee)
	}
	for _baggb := 0; _baggb < _bggd._fccgee; _baggb++ {
		_gadg[_baggb] = nil
	}
	return _gadg
}
func _gbcag(_bbeb []*textLine, _gade string, _defef []*list) *list {
	return &list{_ggbee: _bbeb, _feec: _gade, _gbcab: _defef}
}
func _bbfe(_fccf structElement) []structElement {
	_gedf := []structElement{}
	for _, _cedag := range _fccf._ddcfe {
		for _, _agefb := range _cedag._ddcfe {
			for _, _abceg := range _agefb._ddcfe {
				if _abceg._gacb == "\u004c" {
					_gedf = append(_gedf, _abceg)
				}
			}
		}
	}
	return _gedf
}

type lists []*list

func _fbfab(_affgb _gb.PdfColorspace, _aacg _gb.PdfColor) _eb.Color {
	if _affgb == nil || _aacg == nil {
		return _eb.Black
	}
	_ceab, _ebgdc := _affgb.ColorToRGB(_aacg)
	if _ebgdc != nil {
		_ed.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006fu\u006c\u0064\u0020no\u0074\u0020\u0063\u006f\u006e\u0076e\u0072\u0074\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0025\u0076\u0020\u0028\u0025\u0076)\u0020\u0074\u006f\u0020\u0052\u0047\u0042\u003a \u0025\u0073", _aacg, _affgb, _ebgdc)
		return _eb.Black
	}
	_gfdd, _bgdc := _ceab.(*_gb.PdfColorDeviceRGB)
	if !_bgdc {
		_ed.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0065\u0064 \u0063\u006f\u006c\u006f\u0072\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0052\u0047\u0042\u0020\u0063\u006flo\u0072\u0073\u0070\u0061\u0063\u0065\u003a\u0020\u0025\u0076", _ceab)
		return _eb.Black
	}
	return _eb.NRGBA{R: uint8(_gfdd.R() * 255), G: uint8(_gfdd.G() * 255), B: uint8(_gfdd.B() * 255), A: uint8(255)}
}
func (_cfbag *textPara) writeText(_fcef _a.Writer) {
	if _cfbag._bedfe == nil {
		_cfbag.writeCellText(_fcef)
		return
	}
	for _gfgfa := 0; _gfgfa < _cfbag._bedfe._dege; _gfgfa++ {
		for _affb := 0; _affb < _cfbag._bedfe._fccgee; _affb++ {
			_facg := _cfbag._bedfe.get(_affb, _gfgfa)
			if _facg == nil {
				_fcef.Write([]byte("\u0009"))
			} else {
				_facg.writeCellText(_fcef)
			}
			_fcef.Write([]byte("\u0020"))
		}
		if _gfgfa < _cfbag._bedfe._dege-1 {
			_fcef.Write([]byte("\u000a"))
		}
	}
}
func (_dec compositeCell) split(_gebec, _fbeg []float64) *textTable {
	_gggf := len(_gebec) + 1
	_ccfec := len(_fbeg) + 1
	if _geff {
		_ed.Log.Info("\u0063\u006f\u006d\u0070\u006f\u0073\u0069t\u0065\u0043\u0065l\u006c\u002e\u0073\u0070l\u0069\u0074\u003a\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u000a\u0009\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u003d\u0025\u0073\u000a"+"\u0009\u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073=\u0025\u0036\u002e\u0032\u0066\u000a\t\u0063\u006f\u006c\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u003d%\u0036\u002e\u0032\u0066", _ccfec, _gggf, _dec, _gebec, _fbeg)
		_cfe.Printf("\u0020\u0020\u0020\u0020\u0025\u0064\u0020\u0070\u0061\u0072\u0061\u0073\u000a", len(_dec.paraList))
		for _afea, _gdbc := range _dec.paraList {
			_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _afea, _gdbc.String())
		}
		_cfe.Printf("\u0020\u0020\u0020\u0020\u0025\u0064\u0020\u006c\u0069\u006e\u0065\u0073\u000a", len(_dec.lines()))
		for _fddad, _ffac := range _dec.lines() {
			_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _fddad, _ffac)
		}
	}
	_gebec = _bcaef(_gebec, _dec.Ury, _dec.Lly)
	_fbeg = _bcaef(_fbeg, _dec.Llx, _dec.Urx)
	_cgfga := make(map[uint64]*textPara, _ccfec*_gggf)
	_afaca := textTable{_fccgee: _ccfec, _dege: _gggf, _adda: _cgfga}
	_aaae := _dec.paraList
	_d.Slice(_aaae, func(_fcgba, _fedad int) bool {
		_begg, _adabf := _aaae[_fcgba], _aaae[_fedad]
		_ggffg, _cgbec := _begg.Lly, _adabf.Lly
		if _ggffg != _cgbec {
			return _ggffg < _cgbec
		}
		return _begg.Llx < _adabf.Llx
	})
	_fded := make(map[uint64]_gb.PdfRectangle, _ccfec*_gggf)
	for _fcfb, _gdbce := range _gebec[1:] {
		_efabf := _gebec[_fcfb]
		for _aedcb, _efaf := range _fbeg[1:] {
			_gfga := _fbeg[_aedcb]
			_fded[_aafb(_aedcb, _fcfb)] = _gb.PdfRectangle{Llx: _gfga, Urx: _efaf, Lly: _gdbce, Ury: _efabf}
		}
	}
	if _geff {
		_ed.Log.Info("\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0043\u0065l\u006c\u002e\u0073\u0070\u006c\u0069\u0074\u003a\u0020\u0072e\u0063\u0074\u0073")
		_cfe.Printf("\u0020\u0020\u0020\u0020")
		for _ceefc := 0; _ceefc < _ccfec; _ceefc++ {
			_cfe.Printf("\u0025\u0033\u0030\u0064\u002c\u0020", _ceefc)
		}
		_cfe.Println()
		for _cbced := 0; _cbced < _gggf; _cbced++ {
			_cfe.Printf("\u0020\u0020\u0025\u0032\u0064\u003a", _cbced)
			for _fcee := 0; _fcee < _ccfec; _fcee++ {
				_cfe.Printf("\u00256\u002e\u0032\u0066\u002c\u0020", _fded[_aafb(_fcee, _cbced)])
			}
			_cfe.Println()
		}
	}
	_afcb := func(_ccae *textLine) (int, int) {
		for _effd := 0; _effd < _gggf; _effd++ {
			for _dedg := 0; _dedg < _ccfec; _dedg++ {
				if _egcd(_fded[_aafb(_dedg, _effd)], _ccae.PdfRectangle) {
					return _dedg, _effd
				}
			}
		}
		return -1, -1
	}
	_ccfb := make(map[uint64][]*textLine, _ccfec*_gggf)
	for _, _aafg := range _aaae.lines() {
		_dfacg, _efbg := _afcb(_aafg)
		if _dfacg < 0 {
			continue
		}
		_ccfb[_aafb(_dfacg, _efbg)] = append(_ccfb[_aafb(_dfacg, _efbg)], _aafg)
	}
	for _fgga := 0; _fgga < len(_gebec)-1; _fgga++ {
		_cdda := _gebec[_fgga]
		_adgg := _gebec[_fgga+1]
		for _bggc := 0; _bggc < len(_fbeg)-1; _bggc++ {
			_dfcb := _fbeg[_bggc]
			_fdgce := _fbeg[_bggc+1]
			_cbage := _gb.PdfRectangle{Llx: _dfcb, Urx: _fdgce, Lly: _adgg, Ury: _cdda}
			_adce := _ccfb[_aafb(_bggc, _fgga)]
			if len(_adce) == 0 {
				continue
			}
			_egbbb := _daae(_cbage, _adce)
			_afaca.put(_bggc, _fgga, _egbbb)
		}
	}
	return &_afaca
}
func (_dddbc rulingList) comp(_bafg, _efbga int) bool {
	_gdcab, _egge := _dddbc[_bafg], _dddbc[_efbga]
	_bccccg, _fgba := _gdcab._cffe, _egge._cffe
	if _bccccg != _fgba {
		return _bccccg > _fgba
	}
	if _bccccg == _eefaf {
		return false
	}
	_gfad := func(_cdga bool) bool {
		if _bccccg == _bdfd {
			return _cdga
		}
		return !_cdga
	}
	_gaeec, _ebgdg := _gdcab._abfg, _egge._abfg
	if _gaeec != _ebgdg {
		return _gfad(_gaeec > _ebgdg)
	}
	_gaeec, _ebgdg = _gdcab._daag, _egge._daag
	if _gaeec != _ebgdg {
		return _gfad(_gaeec < _ebgdg)
	}
	return _gfad(_gdcab._fcff < _egge._fcff)
}
func _fggg(_fbged _gb.PdfRectangle, _begd, _gfgcb, _beaa, _ffdg *ruling) gridTile {
	_cegeg := _fbged.Llx
	_efgf := _fbged.Urx
	_bcdc := _fbged.Lly
	_abda := _fbged.Ury
	return gridTile{PdfRectangle: _fbged, _bdcfac: _begd != nil && _begd.encloses(_bcdc, _abda), _babe: _gfgcb != nil && _gfgcb.encloses(_bcdc, _abda), _debb: _beaa != nil && _beaa.encloses(_cegeg, _efgf), _cccb: _ffdg != nil && _ffdg.encloses(_cegeg, _efgf)}
}

// ToTextMark returns the public view of `tm`.
func (_cfdga *textMark) ToTextMark() TextMark {
	return TextMark{Text: _cfdga._ceca, Original: _cfdga._fbgee, BBox: _cfdga._aagcg, Font: _cfdga._gcfc, FontSize: _cfdga._gded, FillColor: _cfdga._egea, StrokeColor: _cfdga._cdfbb, Orientation: _cfdga._cggbc, DirectObject: _cfdga._dgebc, ObjString: _cfdga._bfdcd, Tw: _cfdga.Tw, Th: _cfdga.Th, Tc: _cfdga._ggaf, Index: _cfdga._gegcc}
}
func (_gfggg lineRuling) xMean() float64 { return 0.5 * (_gfggg._efffg.X + _gfggg._agae.X) }
func _edde(_aecde []int) []int {
	_fedeb := make([]int, len(_aecde))
	for _cgdb, _cfecf := range _aecde {
		_fedeb[len(_aecde)-1-_cgdb] = _cfecf
	}
	return _fedeb
}
func (_bfggg rulingList) blocks(_fcbc, _ffea *ruling) bool {
	if _fcbc._daag > _ffea._fcff || _ffea._daag > _fcbc._fcff {
		return false
	}
	_gdca := _df.Max(_fcbc._daag, _ffea._daag)
	_fdgbb := _df.Min(_fcbc._fcff, _ffea._fcff)
	if _fcbc._abfg > _ffea._abfg {
		_fcbc, _ffea = _ffea, _fcbc
	}
	for _, _aedad := range _bfggg {
		if _fcbc._abfg <= _aedad._abfg+_dbgc && _aedad._abfg <= _ffea._abfg+_dbgc && _aedad._daag <= _fdgbb && _gdca <= _aedad._fcff {
			return true
		}
	}
	return false
}
func (_daf *shapesState) stroke(_ggec *[]pathSection) {
	_fbca := pathSection{_ggfc: _daf._babc, Color: _daf._geba.getStrokeColor()}
	*_ggec = append(*_ggec, _fbca)
	if _geaeg {
		_cfe.Printf("\u0020 \u0020\u0020S\u0054\u0052\u004fK\u0045\u003a\u0020\u0025\u0064\u0020\u0073t\u0072\u006f\u006b\u0065\u0073\u0020s\u0073\u003d\u0025\u0073\u0020\u0063\u006f\u006c\u006f\u0072\u003d%\u002b\u0076\u0020\u0025\u0036\u002e\u0032\u0066\u000a", len(*_ggec), _daf, _daf._geba.getStrokeColor(), _fbca.bbox())
		if _faab {
			for _gbec, _egcb := range _daf._babc {
				_cfe.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _gbec, _egcb)
				if _gbec == 10 {
					break
				}
			}
		}
	}
}

// String returns a description of `tm`.
func (_cgba *textMark) String() string {
	return _cfe.Sprintf("\u0025\u002e\u0032f \u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\u0022\u0025\u0073\u0022", _cgba.PdfRectangle, _cgba._gded, _cgba._ceca)
}
func _aabe(_bgdea string) string {
	_gbfb := []rune(_bgdea)
	return string(_gbfb[:len(_gbfb)-1])
}

type textResult struct {
	_gga  PageText
	_bdef int
	_bee  int
}

func (_bcbf rulingList) toGrids() []rulingList {
	if _geaeg {
		_ed.Log.Info("t\u006f\u0047\u0072\u0069\u0064\u0073\u003a\u0020\u0025\u0073", _bcbf)
	}
	_feggd := _bcbf.intersections()
	if _geaeg {
		_ed.Log.Info("\u0074\u006f\u0047r\u0069\u0064\u0073\u003a \u0076\u0065\u0063\u0073\u003d\u0025\u0064 \u0069\u006e\u0074\u0065\u0072\u0073\u0065\u0063\u0074\u0073\u003d\u0025\u0064\u0020", len(_bcbf), len(_feggd))
		for _, _gbbdg := range _gbcbg(_feggd) {
			_cfe.Printf("\u00254\u0064\u003a\u0020\u0025\u002b\u0076\n", _gbbdg, _feggd[_gbbdg])
		}
	}
	_dfdcce := make(map[int]intSet, len(_bcbf))
	for _bggg := range _bcbf {
		_gegbd := _bcbf.connections(_feggd, _bggg)
		if len(_gegbd) > 0 {
			_dfdcce[_bggg] = _gegbd
		}
	}
	if _geaeg {
		_ed.Log.Info("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u0063\u006fn\u006e\u0065\u0063\u0074s=\u0025\u0064", len(_dfdcce))
		for _, _gfdc := range _gbcbg(_dfdcce) {
			_cfe.Printf("\u00254\u0064\u003a\u0020\u0025\u002b\u0076\n", _gfdc, _dfdcce[_gfdc])
		}
	}
	_gbbcg := _edfbf(len(_bcbf), func(_aacab, _eedf int) bool {
		_bcgfc, _cedc := len(_dfdcce[_aacab]), len(_dfdcce[_eedf])
		if _bcgfc != _cedc {
			return _bcgfc > _cedc
		}
		return _bcbf.comp(_aacab, _eedf)
	})
	if _geaeg {
		_ed.Log.Info("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u006f\u0072d\u0065\u0072\u0069\u006eg=\u0025\u0076", _gbbcg)
	}
	_ebbg := [][]int{{_gbbcg[0]}}
_bgffg:
	for _, _fcae := range _gbbcg[1:] {
		for _bgdee, _eeaa := range _ebbg {
			for _, _ccge := range _eeaa {
				if _dfdcce[_ccge].has(_fcae) {
					_ebbg[_bgdee] = append(_eeaa, _fcae)
					continue _bgffg
				}
			}
		}
		_ebbg = append(_ebbg, []int{_fcae})
	}
	if _geaeg {
		_ed.Log.Info("\u0074o\u0047r\u0069\u0064\u0073\u003a\u0020i\u0067\u0072i\u0064\u0073\u003d\u0025\u0076", _ebbg)
	}
	_d.SliceStable(_ebbg, func(_abfcc, _addc int) bool { return len(_ebbg[_abfcc]) > len(_ebbg[_addc]) })
	for _, _cfcg := range _ebbg {
		_d.Slice(_cfcg, func(_ebaea, _decb int) bool { return _bcbf.comp(_cfcg[_ebaea], _cfcg[_decb]) })
	}
	_dffbc := make([]rulingList, len(_ebbg))
	for _efec, _facb := range _ebbg {
		_fdbbf := make(rulingList, len(_facb))
		for _bcca, _eadeb := range _facb {
			_fdbbf[_bcca] = _bcbf[_eadeb]
		}
		_dffbc[_efec] = _fdbbf
	}
	if _geaeg {
		_ed.Log.Info("\u0074o\u0047r\u0069\u0064\u0073\u003a\u0020g\u0072\u0069d\u0073\u003d\u0025\u002b\u0076", _dffbc)
	}
	var _ebcb []rulingList
	for _, _cdag := range _dffbc {
		if _fcfg, _deafa := _cdag.isActualGrid(); _deafa {
			_cdag = _fcfg
			_cdag = _cdag.snapToGroups()
			_ebcb = append(_ebcb, _cdag)
		}
	}
	if _geaeg {
		_afbf("t\u006fG\u0072\u0069\u0064\u0073\u003a\u0020\u0061\u0063t\u0075\u0061\u006c\u0047ri\u0064\u0073", _ebcb)
		_ed.Log.Info("\u0074\u006f\u0047\u0072\u0069\u0064\u0073\u003a\u0020\u0067\u0072\u0069\u0064\u0073\u003d%\u0064 \u0061\u0063\u0074\u0075\u0061\u006c\u0047\u0072\u0069\u0064\u0073\u003d\u0025\u0064", len(_dffbc), len(_ebcb))
	}
	return _ebcb
}

// TextTable represents a table.
// Cells are ordered top-to-bottom, left-to-right.
// Cells[y] is the (0-offset) y'th row in the table.
// Cells[y][x] is the (0-offset) x'th column in the table.
type TextTable struct {
	_gb.PdfRectangle
	W, H  int
	Cells [][]TableCell
}

const _cegg = 1.0 / 1000.0

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
func (_aefb PageText) List() lists {
	_ggbd := !_aefb._ege._gdeb
	_adcd := _aefb.getParagraphs()
	_fdaf := true
	if _aefb._ddbd == nil || *_aefb._ddbd == nil {
		_fdaf = false
	}
	_bgebc := _adcd.list()
	if _fdaf && _ggbd {
		_cfacf := _gfeg(&_adcd)
		_ebcf := &structTreeRoot{}
		_ebcf.parseStructTreeRoot(*_aefb._ddbd)
		if _ebcf._gafdd == nil {
			_ed.Log.Debug("\u004c\u0069\u0073\u0074\u003a\u0020\u0073t\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0064\u006f\u0065\u0073\u006e'\u0074\u0020\u0068\u0061\u0076e\u0020\u0061\u006e\u0079\u0020\u0063\u006f\u006e\u0074e\u006e\u0074\u002c\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0074\u0065\u0078\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0069\u006e\u0067\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u0020\u0069\u006e\u0073\u0074\u0065\u0061\u0064\u002e")
			return _bgebc
		}
		_bgebc = _ebcf.buildList(_cfacf, _aefb._ddeg)
	}
	return _bgebc
}

type gridTile struct {
	_gb.PdfRectangle
	_cccb, _bdcfac, _debb, _babe bool
}

func (_afg *imageExtractContext) processOperand(_faa *_aad.ContentStreamOperation, _bebe _aad.GraphicsState, _cadf *_gb.PdfPageResources) error {
	if _faa.Operand == "\u0042\u0049" && len(_faa.Params) == 1 {
		_gfc, _aaa := _faa.Params[0].(*_aad.ContentStreamInlineImage)
		if !_aaa {
			return nil
		}
		if _bfb, _ecf := _gg.GetBoolVal(_gfc.ImageMask); _ecf {
			if _bfb && !_afg._eag.IncludeInlineStencilMasks {
				return nil
			}
		}
		return _afg.extractInlineImage(_gfc, _bebe, _cadf)
	} else if _faa.Operand == "\u0044\u006f" && len(_faa.Params) == 1 {
		_gff, _geb := _gg.GetName(_faa.Params[0])
		if !_geb {
			_ed.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0079\u0070\u0065")
			return _ccb
		}
		_, _ae := _cadf.GetXObjectByName(*_gff)
		switch _ae {
		case _gb.XObjectTypeImage:
			return _afg.extractXObjectImage(_gff, _bebe, _cadf)
		case _gb.XObjectTypeForm:
			return _afg.extractFormImages(_gff, _bebe, _cadf)
		}
	} else if _afg._cbg && (_faa.Operand == "\u0073\u0063\u006e" || _faa.Operand == "\u0053\u0043\u004e") && len(_faa.Params) == 1 {
		_ddb, _gda := _gg.GetName(_faa.Params[0])
		if !_gda {
			_ed.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0054\u0079\u0070\u0065")
			return _ccb
		}
		_acd, _gda := _cadf.GetPatternByName(*_ddb)
		if !_gda {
			_ed.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0050\u0061\u0074\u0074\u0065\u0072n\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
			return nil
		}
		if _acd.IsTiling() {
			_gc := _acd.GetAsTilingPattern()
			_aca, _bea := _gc.GetContentStream()
			if _bea != nil {
				return _bea
			}
			_bea = _afg.extractContentStreamImages(string(_aca), _gc.Resources)
			if _bea != nil {
				return _bea
			}
		}
	} else if (_faa.Operand == "\u0063\u0073" || _faa.Operand == "\u0043\u0053") && len(_faa.Params) >= 1 {
		_afg._cbg = _faa.Params[0].String() == "\u0050a\u0074\u0074\u0065\u0072\u006e"
	}
	return nil
}
func (_defa *wordBag) depthBand(_ebec, _dfgcd float64) []int {
	if len(_defa._gegc) == 0 {
		return nil
	}
	return _defa.depthRange(_defa.getDepthIdx(_ebec), _defa.getDepthIdx(_dfgcd))
}
func (_cgbd lineRuling) yMean() float64 { return 0.5 * (_cgbd._efffg.Y + _cgbd._agae.Y) }
func (_egcfg rulingList) secMinMax() (float64, float64) {
	_ggca, _edegf := _egcfg[0]._daag, _egcfg[0]._fcff
	for _, _ccea := range _egcfg[1:] {
		if _ccea._daag < _ggca {
			_ggca = _ccea._daag
		}
		if _ccea._fcff > _edegf {
			_edegf = _ccea._fcff
		}
	}
	return _ggca, _edegf
}
func (_ffbb paraList) merge() *textPara {
	_ed.Log.Trace("\u006d\u0065\u0072\u0067\u0065:\u0020\u0070\u0061\u0072\u0061\u0073\u003d\u0025\u0064\u0020\u003d\u003d\u003d=\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u0078\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d\u003d", len(_ffbb))
	if len(_ffbb) == 0 {
		return nil
	}
	_ffbb.sortReadingOrder()
	_eeaba := _ffbb[0].PdfRectangle
	_ddddc := _ffbb[0]._abbe
	for _, _cefe := range _ffbb[1:] {
		_eeaba = _fde(_eeaba, _cefe.PdfRectangle)
		_ddddc = append(_ddddc, _cefe._abbe...)
	}
	return _daae(_eeaba, _ddddc)
}

// String returns a description of `t`.
func (_bcabc *textTable) String() string {
	return _cfe.Sprintf("\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u0020\u0025\u0074", _bcabc._fccgee, _bcabc._dege, _bcabc._baccf)
}
func _fbaff(_ggbgf *_gb.Image, _aeega _eb.Color) _gf.Image {
	_dcbcb, _ecff := int(_ggbgf.Width), int(_ggbgf.Height)
	_abdd := _gf.NewRGBA(_gf.Rect(0, 0, _dcbcb, _ecff))
	for _effa := 0; _effa < _ecff; _effa++ {
		for _aabc := 0; _aabc < _dcbcb; _aabc++ {
			_ddfa, _beaec := _ggbgf.ColorAt(_aabc, _effa)
			if _beaec != nil {
				_ed.Log.Debug("\u0057\u0041\u0052\u004e\u003a\u0020\u0063o\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0072\u0065\u0074\u0072\u0069\u0065v\u0065 \u0069\u006d\u0061\u0067\u0065\u0020m\u0061\u0073\u006b\u0020\u0076\u0061\u006cu\u0065\u0020\u0061\u0074\u0020\u0028\u0025\u0064\u002c\u0020\u0025\u0064\u0029\u002e\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0069\u006e\u0063\u006f\u0072\u0072\u0065\u0063t\u002e", _aabc, _effa)
				continue
			}
			_adcba, _afeg, _gfce, _ := _ddfa.RGBA()
			var _bebbe _eb.Color
			if _adcba+_afeg+_gfce == 0 {
				_bebbe = _aeega
			} else {
				_bebbe = _eb.Transparent
			}
			_abdd.Set(_aabc, _effa, _bebbe)
		}
	}
	return _abdd
}
func (_beca *textObject) setTextLeading(_ffef float64) {
	if _beca == nil {
		return
	}
	_beca._cbag._fea = _ffef
}
func _cbd(_bfef []*textWord, _cbagd float64, _eadd, _cbef rulingList) *wordBag {
	_geaf := _dfcc(_bfef[0], _cbagd, _eadd, _cbef)
	for _, _daa := range _bfef[1:] {
		_gead := _bedf(_daa._cbfcee)
		_geaf._gegc[_gead] = append(_geaf._gegc[_gead], _daa)
		_geaf.PdfRectangle = _fde(_geaf.PdfRectangle, _daa.PdfRectangle)
	}
	_geaf.sort()
	return _geaf
}
func _aedg(_gfadg, _cafdd int) int {
	if _gfadg > _cafdd {
		return _gfadg
	}
	return _cafdd
}
func (_gggc compositeCell) parasBBox() (paraList, _gb.PdfRectangle) {
	return _gggc.paraList, _gggc.PdfRectangle
}

type intSet map[int]struct{}

const (
	_cfd = "\u0045\u0052R\u004f\u0052\u003a\u0020\u0043\u0061\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065"
	_cad = "\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0043a\u006e\u0027\u0074 g\u0065\u0074\u0020\u0066\u006f\u006et\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u002c\u0020\u0066\u006fn\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064"
	_ea  = "\u0045\u0052\u0052O\u0052\u003a\u0020\u0043\u0061\u006e\u0027\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002c\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065"
)

func (_fgcg *textMark) bbox() _gb.PdfRectangle { return _fgcg.PdfRectangle }
func _accb(_cefd *list) []*textLine {
	for _, _cfefc := range _cefd._gbcab {
		switch _cfefc._feec {
		case "\u004c\u0042\u006fd\u0079":
			if len(_cfefc._ggbee) != 0 {
				return _cfefc._ggbee
			}
			return _accb(_cfefc)
		case "\u0053\u0070\u0061\u006e":
			return _cfefc._ggbee
		case "I\u006e\u006c\u0069\u006e\u0065\u0053\u0068\u0061\u0070\u0065":
			return _cfefc._ggbee
		}
	}
	return nil
}
func _cafd(_cbgae, _ggga int) int {
	if _cbgae < _ggga {
		return _cbgae
	}
	return _ggga
}
func (_ccba rulingList) findPrimSec(_dbeff, _dcdafe float64) *ruling {
	for _, _afad := range _ccba {
		if _cdgd(_afad._abfg-_dbeff) && _afad._daag-_egaa <= _dcdafe && _dcdafe <= _afad._fcff+_egaa {
			return _afad
		}
	}
	return nil
}
func _gaba(_gbab, _ddbaa, _cbba float64) rulingKind {
	if _gbab >= _cbba && _dffe(_ddbaa, _gbab) {
		return _bdfd
	}
	if _ddbaa >= _cbba && _dffe(_gbab, _ddbaa) {
		return _gfafc
	}
	return _eefaf
}
func _cadba(_cbae, _bcfec _gf.Image) _gf.Image {
	_eceb, _gfee := _bcfec.Bounds().Size(), _cbae.Bounds().Size()
	_gegbb, _eabbf := _eceb.X, _eceb.Y
	if _gfee.X > _gegbb {
		_gegbb = _gfee.X
	}
	if _gfee.Y > _eabbf {
		_eabbf = _gfee.Y
	}
	_bacaa := _gf.Rect(0, 0, _gegbb, _eabbf)
	if _eceb.X != _gegbb || _eceb.Y != _eabbf {
		_fdcde := _gf.NewRGBA(_bacaa)
		_dc.BiLinear.Scale(_fdcde, _bacaa, _cbae, _bcfec.Bounds(), _dc.Over, nil)
		_bcfec = _fdcde
	}
	if _gfee.X != _gegbb || _gfee.Y != _eabbf {
		_dgdg := _gf.NewRGBA(_bacaa)
		_dc.BiLinear.Scale(_dgdg, _bacaa, _cbae, _cbae.Bounds(), _dc.Over, nil)
		_cbae = _dgdg
	}
	_gbcbd := _gf.NewRGBA(_bacaa)
	_dc.DrawMask(_gbcbd, _bacaa, _cbae, _gf.Point{}, _bcfec, _gf.Point{}, _dc.Over)
	return _gbcbd
}
func (_ccaec *textTable) get(_ceae, _ddeb int) *textPara { return _ccaec._adda[_aafb(_ceae, _ddeb)] }
func (_feab *textWord) computeText() string {
	_cgea := make([]string, len(_feab._fdgbc))
	for _agcbdb, _cgbb := range _feab._fdgbc {
		_cgea[_agcbdb] = _cgbb._ceca
	}
	return _ag.Join(_cgea, "")
}
func (_fcgd *shapesState) lastpointEstablished() (_cc.Point, bool) {
	if _fcgd._dfff {
		return _fcgd._bcb, false
	}
	_bfcgd := len(_fcgd._babc)
	if _bfcgd > 0 && _fcgd._babc[_bfcgd-1]._beac {
		return _fcgd._babc[_bfcgd-1].last(), false
	}
	return _cc.Point{}, true
}
func (_eeg *textObject) nextLine() { _eeg.moveLP(0, -_eeg._cbag._fea) }
func (_fbcb *PageText) computeViews() {
	_gbbb := _fbcb.getParagraphs()
	_gaaf := new(_cf.Buffer)
	_gbbb.writeText(_gaaf)
	_fbcb._ccbd = _gaaf.String()
	_fbcb._afc = _gbbb.toTextMarks()
	_fbcb._ded = _gbbb.tables()
	if _geff {
		_ed.Log.Info("\u0063\u006f\u006dpu\u0074\u0065\u0056\u0069\u0065\u0077\u0073\u003a\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u003d\u0025\u0064", len(_fbcb._ded))
	}
}
func (_fead *shapesState) cubicTo(_fdbdb, _cbad, _caee, _aeed, _bccc, _bcfed float64) {
	if _dgeg {
		_ed.Log.Info("\u0063\u0075\u0062\u0069\u0063\u0054\u006f\u003a")
	}
	_fead.addPoint(_bccc, _bcfed)
}
func (_bbe *textObject) renderText(_cfefe _gg.PdfObject, _ccbc []byte, _gbbd int) error {
	if _bbe._efcbd {
		_ed.Log.Debug("\u0072\u0065\u006e\u0064\u0065r\u0054\u0065\u0078\u0074\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064 \u0066\u006f\u006e\u0074\u002e\u0020\u004e\u006f\u0074\u0020\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067\u002e")
		return nil
	}
	_dgee := _bbe.getCurrentFont()
	_bcga := _dgee.BytesToCharcodes(_ccbc)
	_ddec, _beg, _gge := _dgee.CharcodesToStrings(_bcga)
	if _gge > 0 {
		_ed.Log.Debug("\u0072\u0065nd\u0065\u0072\u0054e\u0078\u0074\u003a\u0020num\u0043ha\u0072\u0073\u003d\u0025\u0064\u0020\u006eum\u004d\u0069\u0073\u0073\u0065\u0073\u003d%\u0064", _beg, _gge)
	}
	_bbe._cbag._fdbd += _beg
	_bbe._cbag._eec += _gge
	_feff := _bbe._cbag
	_bfcf := _feff._agge
	_bae := _feff._fceg / 100.0
	_cffb := _cegg
	if _dgee.Subtype() == "\u0054\u0079\u0070e\u0033" {
		_cffb = 1
	}
	_gafd, _ceb := _dgee.GetRuneMetrics(' ')
	if !_ceb {
		_gafd, _ceb = _dgee.GetCharMetrics(32)
	}
	if !_ceb {
		_gafd, _ = _gb.DefaultFont().GetRuneMetrics(' ')
	}
	_gedb := _gafd.Wx * _cffb
	_ed.Log.Trace("\u0073p\u0061\u0063e\u0057\u0069\u0064t\u0068\u003d\u0025\u002e\u0032\u0066\u0020t\u0065\u0078\u0074\u003d\u0025\u0071 \u0066\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0066\u006f\u006et\u0053\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066", _gedb, _ddec, _dgee, _bfcf)
	_egfa := _cc.NewMatrix(_bfcf*_bae, 0, 0, _bfcf, 0, _feff._gce)
	if _geafa {
		_ed.Log.Info("\u0072\u0065\u006e\u0064\u0065\u0072T\u0065\u0078\u0074\u003a\u0020\u0025\u0064\u0020\u0063\u006f\u0064\u0065\u0073=\u0025\u002b\u0076\u0020\u0074\u0065\u0078t\u0073\u003d\u0025\u0071", len(_bcga), _bcga, _ddec)
	}
	_ed.Log.Trace("\u0072\u0065\u006e\u0064\u0065\u0072T\u0065\u0078\u0074\u003a\u0020\u0025\u0064\u0020\u0063\u006f\u0064\u0065\u0073=\u0025\u002b\u0076\u0020\u0072\u0075\u006ee\u0073\u003d\u0025\u0071", len(_bcga), _bcga, len(_ddec))
	_afac := _bbe.getFillColor()
	_aadg := _bbe.getStrokeColor()
	for _afdd, _cga := range _ddec {
		_bcd := []rune(_cga)
		if len(_bcd) == 1 && _bcd[0] == '\x00' {
			continue
		}
		_abc := _bcga[_afdd]
		_bfcg := _bbe._ccg.CTM.Mult(_bbe._ffb).Mult(_egfa)
		_efa := 0.0
		if len(_bcd) == 1 && _bcd[0] == 32 {
			_efa = _feff._bfa
		}
		_baed, _eebg := _dgee.GetCharMetrics(_abc)
		if !_eebg {
			_ed.Log.Debug("\u0045R\u0052\u004fR\u003a\u0020\u004e\u006f \u006d\u0065\u0074r\u0069\u0063\u0020\u0066\u006f\u0072\u0020\u0063\u006fde\u003d\u0025\u0064 \u0072\u003d0\u0078\u0025\u0030\u0034\u0078\u003d%\u002b\u0071 \u0025\u0073", _abc, _bcd, _bcd, _dgee)
			return _cfe.Errorf("\u006e\u006f\u0020\u0063\u0068\u0061\u0072\u0020\u006d\u0065\u0074\u0072\u0069\u0063\u0073:\u0020f\u006f\u006e\u0074\u003d\u0025\u0073\u0020\u0063\u006f\u0064\u0065\u003d\u0025\u0064", _dgee.String(), _abc)
		}
		_dfde := _cc.Point{X: _baed.Wx * _cffb, Y: _baed.Wy * _cffb}
		_gdaf := _cc.Point{X: (_dfde.X*_bfcf + _efa) * _bae}
		_dcb := _cc.Point{X: (_dfde.X*_bfcf + _feff._bfdf + _efa) * _bae}
		if _geafa {
			_ed.Log.Info("\u0074\u0066\u0073\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0063\u003d\u0025\u002e\u0032f\u0020t\u0077\u003d\u0025\u002e\u0032\u0066\u0020\u0074\u0068\u003d\u0025\u002e\u0032\u0066", _bfcf, _feff._bfdf, _feff._bfa, _bae)
			_ed.Log.Info("\u0064x\u002c\u0064\u0079\u003d%\u002e\u0033\u0066\u0020\u00740\u003d%\u002e3\u0066\u0020\u0074\u003d\u0025\u002e\u0033f", _dfde, _gdaf, _dcb)
		}
		_cebc := _gggb(_gdaf)
		_ebddc := _gggb(_dcb)
		_eed := _bbe._ccg.CTM.Mult(_bbe._ffb).Mult(_cebc)
		if _abb {
			_ed.Log.Info("e\u006e\u0064\u003a\u000a\tC\u0054M\u003d\u0025\u0073\u000a\u0009 \u0074\u006d\u003d\u0025\u0073\u000a"+"\u0009\u0020t\u0064\u003d\u0025s\u0020\u0078\u006c\u0061\u0074\u003d\u0025\u0073\u000a"+"\u0009t\u0064\u0030\u003d\u0025s\u000a\u0009\u0020\u0020\u2192 \u0025s\u0020x\u006c\u0061\u0074\u003d\u0025\u0073", _bbe._ccg.CTM, _bbe._ffb, _ebddc, _bba(_bbe._ccg.CTM.Mult(_bbe._ffb).Mult(_ebddc)), _cebc, _eed, _bba(_eed))
		}
		_bcfe, _fgc := _bbe.newTextMark(_dd.ExpandLigatures(_bcd), _bfcg, _bba(_eed), _df.Abs(_gedb*_bfcg.ScalingFactorX()), _dgee, _bbe._cbag._bfdf, _afac, _aadg, _cfefe, _ddec, _afdd, _gbbd)
		if !_fgc {
			_ed.Log.Debug("\u0054\u0065\u0078\u0074\u0020\u006d\u0061\u0072\u006b\u0020\u006f\u0075\u0074\u0073\u0069d\u0065 \u0070\u0061\u0067\u0065\u002e\u0020\u0053\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		if _dgee == nil {
			_ed.Log.Debug("\u0045R\u0052O\u0052\u003a\u0020\u004e\u006f\u0020\u0066\u006f\u006e\u0074\u002e")
		} else if _dgee.Encoder() == nil {
			_ed.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020N\u006f\u0020\u0065\u006e\u0063\u006f\u0064\u0069\u006eg\u002e\u0020\u0066o\u006et\u003d\u0025\u0073", _dgee)
		} else {
			if _afb, _dda := _dgee.Encoder().CharcodeToRune(_abc); _dda {
				_bcfe._fbgee = string(_afb)
			}
		}
		_ed.Log.Trace("i\u003d\u0025\u0064\u0020\u0063\u006fd\u0065\u003d\u0025\u0064\u0020\u006d\u0061\u0072\u006b=\u0025\u0073\u0020t\u0072m\u003d\u0025\u0073", _afdd, _abc, _bcfe, _bfcg)
		_bbe._gafa = append(_bbe._gafa, &_bcfe)
		_bbe._ffb.Concat(_ebddc)
	}
	return nil
}
func (_baffb *textTable) compositeRowCorridors() map[int][]float64 {
	_cffef := make(map[int][]float64, _baffb._dege)
	if _geff {
		_ed.Log.Info("c\u006f\u006d\u0070\u006f\u0073\u0069t\u0065\u0052\u006f\u0077\u0043\u006f\u0072\u0072\u0069d\u006f\u0072\u0073:\u0020h\u003d\u0025\u0064", _baffb._dege)
	}
	for _gdfd := 1; _gdfd < _baffb._dege; _gdfd++ {
		var _deafad []compositeCell
		for _faegb := 0; _faegb < _baffb._fccgee; _faegb++ {
			if _dgfa, _gbcbe := _baffb._faeg[_aafb(_faegb, _gdfd)]; _gbcbe {
				_deafad = append(_deafad, _dgfa)
			}
		}
		if len(_deafad) == 0 {
			continue
		}
		_fdefb := _cebdd(_deafad)
		_cffef[_gdfd] = _fdefb
		if _geff {
			_cfe.Printf("\u0020\u0020\u0020\u0025\u0032\u0064\u003a\u0020\u00256\u002e\u0032\u0066\u000a", _gdfd, _fdefb)
		}
	}
	return _cffef
}

// Text returns the extracted page text.
func (_fcg PageText) Text() string { return _fcg._ccbd }
func (_cege gridTiling) log(_ccded string) {
	if !_aaga {
		return
	}
	_ed.Log.Info("\u0074i\u006ci\u006e\u0067\u003a\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0071", len(_cege._edfg), len(_cege._abbf), _ccded)
	_cfe.Printf("\u0020\u0020\u0020l\u006c\u0078\u003d\u0025\u002e\u0032\u0066\u000a", _cege._edfg)
	_cfe.Printf("\u0020\u0020\u0020l\u006c\u0079\u003d\u0025\u002e\u0032\u0066\u000a", _cege._abbf)
	for _cadbc, _bdee := range _cege._abbf {
		_degc, _gcbff := _cege._aaef[_bdee]
		if !_gcbff {
			continue
		}
		_cfe.Printf("%\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u000a", _cadbc, _bdee)
		for _bcdf, _decd := range _cege._edfg {
			_ebdg, _cgcc := _degc[_decd]
			if !_cgcc {
				continue
			}
			_cfe.Printf("\u0025\u0038\u0064\u003a\u0020\u0025\u0073\u000a", _bcdf, _ebdg.String())
		}
	}
}

// GetContentStreamOps returns the contentStreamOps field of `pt`.
func (_ede *PageText) GetContentStreamOps() *_aad.ContentStreamOperations { return _ede._dff }

type imageExtractContext struct {
	_cdb []ImageMark
	_fbd int
	_ce  int
	_ebf int
	_eae map[*_gg.PdfObjectStream]*cachedImage
	_eag *ImageExtractOptions
	_cbg bool
}

func _abee(_fddeg []_gg.PdfObject) (_caggf, _bacee float64, _dcbd error) {
	if len(_fddeg) != 2 {
		return 0, 0, _cfe.Errorf("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066 \u0070\u0061\u0072\u0061\u006d\u0073\u003a \u0025\u0064", len(_fddeg))
	}
	_agbf, _dcbd := _gg.GetNumbersAsFloat(_fddeg)
	if _dcbd != nil {
		return 0, 0, _dcbd
	}
	return _agbf[0], _agbf[1], nil
}
func _edeeg(_fbcfe float64, _eccd int) int {
	if _eccd == 0 {
		_eccd = 1
	}
	_bfec := float64(_eccd)
	return int(_df.Round(_fbcfe/_bfec) * _bfec)
}
func (_ebbbb *textObject) getCurrentFont() *_gb.PdfFont {
	_cgfg := _ebbbb._cbag._bdg
	if _cgfg == nil {
		_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u004e\u006f\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u002e\u0020U\u0073\u0069\u006e\u0067\u0020d\u0065\u0066a\u0075\u006c\u0074\u002e")
		return _gb.DefaultFont()
	}
	return _cgfg
}
func (_egcbg paraList) xNeighbours(_gadb float64) map[*textPara][]int {
	_geef := make([]event, 2*len(_egcbg))
	if _gadb == 0 {
		for _gafdbd, _dcebb := range _egcbg {
			_geef[2*_gafdbd] = event{_dcebb.Llx, true, _gafdbd}
			_geef[2*_gafdbd+1] = event{_dcebb.Urx, false, _gafdbd}
		}
	} else {
		for _edagdf, _cefca := range _egcbg {
			_geef[2*_edagdf] = event{_cefca.Llx - _gadb*_cefca.fontsize(), true, _edagdf}
			_geef[2*_edagdf+1] = event{_cefca.Urx + _gadb*_cefca.fontsize(), false, _edagdf}
		}
	}
	return _egcbg.eventNeighbours(_geef)
}
func (_ccbde *textObject) getFont(_aaec string) (*_gb.PdfFont, error) {
	if _ccbde._cgf._ebb != nil {
		_gbcd, _acce := _ccbde.getFontDict(_aaec)
		if _acce != nil {
			_ed.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u0067\u0065\u0074\u0046\u006f\u006e\u0074:\u0020n\u0061m\u0065=\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073", _aaec, _acce.Error())
			return nil, _acce
		}
		_ccbde._cgf._bda++
		_bfdb, _fafc := _ccbde._cgf._ebb[_gbcd.String()]
		if _fafc {
			_bfdb._dgca = _ccbde._cgf._bda
			return _bfdb._cegb, nil
		}
	}
	_deeda, _deaa := _ccbde.getFontDict(_aaec)
	if _deaa != nil {
		return nil, _deaa
	}
	_ffcd, _deaa := _ccbde.getFontDirect(_aaec)
	if _deaa != nil {
		return nil, _deaa
	}
	if _ccbde._cgf._ebb != nil {
		_edd := fontEntry{_ffcd, _ccbde._cgf._bda}
		if len(_ccbde._cgf._ebb) >= _baee {
			var _gefc []string
			for _edbg := range _ccbde._cgf._ebb {
				_gefc = append(_gefc, _edbg)
			}
			_d.Slice(_gefc, func(_dgf, _adgc int) bool {
				return _ccbde._cgf._ebb[_gefc[_dgf]]._dgca < _ccbde._cgf._ebb[_gefc[_adgc]]._dgca
			})
			delete(_ccbde._cgf._ebb, _gefc[0])
		}
		_ccbde._cgf._ebb[_deeda.String()] = _edd
	}
	return _ffcd, nil
}
func _adag(_aedfg []*textLine, _befa string) string {
	var _cbfd _ag.Builder
	_egege := 0.0
	for _dgeb, _bcab := range _aedfg {
		_bbce := _bcab.text()
		_afdf := _bcab._bdbg
		if _dgeb < len(_aedfg)-1 {
			_egege = _aedfg[_dgeb+1]._bdbg
		} else {
			_egege = 0.0
		}
		_cbfd.WriteString(_befa)
		_cbfd.WriteString(_bbce)
		if _egege != _afdf {
			_cbfd.WriteString("\u000a")
		} else {
			_cbfd.WriteString("\u0020")
		}
	}
	return _cbfd.String()
}
func _acgef(_acae []*textMark, _gbgd _gb.PdfRectangle, _gegb rulingList, _gdee []gridTiling, _gdcb bool) paraList {
	_ed.Log.Trace("\u006d\u0061\u006b\u0065\u0054\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u003a \u0025\u0064\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0073\u0020\u0070\u0061\u0067\u0065\u0053\u0069\u007a\u0065=\u0025\u002e\u0032\u0066", len(_acae), _gbgd)
	if len(_acae) == 0 {
		return nil
	}
	_cfbdb := _fgege(_acae, _gbgd)
	if len(_cfbdb) == 0 {
		return nil
	}
	_gegb.log("\u006d\u0061\u006be\u0054\u0065\u0078\u0074\u0050\u0061\u0067\u0065")
	_ffge, _fcdc := _gegb.vertsHorzs()
	_ddcb := _cbd(_cfbdb, _gbgd.Ury, _ffge, _fcdc)
	_agca := _gaee(_ddcb, _gbgd.Ury, _ffge, _fcdc)
	_agca = _ggff(_agca)
	_baca := make(paraList, 0, len(_agca))
	for _, _edbgg := range _agca {
		_bfbg := _edbgg.arrangeText()
		if _bfbg != nil {
			_baca = append(_baca, _bfbg)
		}
	}
	if !_gdcb && len(_baca) >= _dgea {
		_baca = _baca.extractTables(_gdee)
	}
	_baca.sortReadingOrder()
	if !_gdcb {
		_baca.sortTopoOrder()
	}
	_baca.log("\u0073\u006f\u0072te\u0064\u0020\u0069\u006e\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0064\u0065\u0072")
	return _baca
}
func _fgag(_bdae map[int][]float64) string {
	_ddbab := _dgce(_bdae)
	_dbad := make([]string, len(_bdae))
	for _caebe, _cceab := range _ddbab {
		_dbad[_caebe] = _cfe.Sprintf("\u0025\u0064\u003a\u0020\u0025\u002e\u0032\u0066", _cceab, _bdae[_cceab])
	}
	return _cfe.Sprintf("\u007b\u0025\u0073\u007d", _ag.Join(_dbad, "\u002c\u0020"))
}

const (
	_gbeab markKind = iota
	_fdacf
	_cbcd
	_ecdcf
)

func _dabb(_cabe *wordBag, _dfbg int) *textLine {
	_cafb := _cabe.firstWord(_dfbg)
	_ecde := textLine{PdfRectangle: _cafb.PdfRectangle, _egbd: _cafb._fgdbd, _bdbg: _cafb._cbfcee}
	_ecde.pullWord(_cabe, _cafb, _dfbg)
	return &_ecde
}

// String returns a description of `l`.
func (_acfd *textLine) String() string {
	return _cfe.Sprintf("\u0025\u002e2\u0066\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0066\u006f\u006e\u0074\u0073\u0069\u007a\u0065\u003d\u0025\u002e\u0032\u0066\u0020\"%\u0073\u0022", _acfd._bdbg, _acfd.PdfRectangle, _acfd._egbd, _acfd.text())
}

var _gdcf = TextMark{Text: "\u005b\u0058\u005d", Original: "\u0020", Meta: true, FillColor: _eb.White, StrokeColor: _eb.White}

func (_aade paraList) eventNeighbours(_eege []event) map[*textPara][]int {
	_d.Slice(_eege, func(_cefg, _ddgba int) bool {
		_cacf, _edaa := _eege[_cefg], _eege[_ddgba]
		_cfcag, _efcf := _cacf._bggcc, _edaa._bggcc
		if _cfcag != _efcf {
			return _cfcag < _efcf
		}
		if _cacf._aagcb != _edaa._aagcb {
			return _cacf._aagcb
		}
		return _cefg < _ddgba
	})
	_acgbc := make(map[int]intSet)
	_aabd := make(intSet)
	for _, _gefcc := range _eege {
		if _gefcc._aagcb {
			_acgbc[_gefcc._caccdd] = make(intSet)
			for _faeca := range _aabd {
				if _faeca != _gefcc._caccdd {
					_acgbc[_gefcc._caccdd].add(_faeca)
					_acgbc[_faeca].add(_gefcc._caccdd)
				}
			}
			_aabd.add(_gefcc._caccdd)
		} else {
			_aabd.del(_gefcc._caccdd)
		}
	}
	_bdbb := map[*textPara][]int{}
	for _edgdc, _dddba := range _acgbc {
		_ebgg := _aade[_edgdc]
		if len(_dddba) == 0 {
			_bdbb[_ebgg] = nil
			continue
		}
		_fbgb := make([]int, len(_dddba))
		_acbcc := 0
		for _bdbgf := range _dddba {
			_fbgb[_acbcc] = _bdbgf
			_acbcc++
		}
		_bdbb[_ebgg] = _fbgb
	}
	return _bdbb
}
func (_gfbe *subpath) makeRectRuling(_bdaga _eb.Color) (*ruling, bool) {
	if _afbec {
		_ed.Log.Info("\u006d\u0061\u006beR\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006e\u0067\u003a\u0020\u0070\u0061\u0074\u0068\u003d\u0025\u0076", _gfbe)
	}
	_cfefd := _gfbe._ccgf[:4]
	_abffc := make(map[int]rulingKind, len(_cfefd))
	for _dcea, _gdccg := range _cfefd {
		_fadbg := _gfbe._ccgf[(_dcea+1)%4]
		_abffc[_dcea] = _dbfg(_gdccg, _fadbg)
		if _afbec {
			_cfe.Printf("\u0025\u0034\u0064: \u0025\u0073\u0020\u003d\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u002d\u0020\u0025\u0036\u002e\u0032\u0066", _dcea, _abffc[_dcea], _gdccg, _fadbg)
		}
	}
	if _afbec {
		_cfe.Printf("\u0020\u0020\u0020\u006b\u0069\u006e\u0064\u0073\u003d\u0025\u002b\u0076\u000a", _abffc)
	}
	var _fgdb, _gdeda []int
	for _eedg, _bfea := range _abffc {
		switch _bfea {
		case _bdfd:
			_gdeda = append(_gdeda, _eedg)
		case _gfafc:
			_fgdb = append(_fgdb, _eedg)
		}
	}
	if _afbec {
		_cfe.Printf("\u0020\u0020 \u0068\u006f\u0072z\u0073\u003d\u0025\u0064\u0020\u0025\u002b\u0076\u000a", len(_gdeda), _gdeda)
		_cfe.Printf("\u0020\u0020 \u0076\u0065\u0072t\u0073\u003d\u0025\u0064\u0020\u0025\u002b\u0076\u000a", len(_fgdb), _fgdb)
	}
	_dfcg := (len(_gdeda) == 2 && len(_fgdb) == 2) || (len(_gdeda) == 2 && len(_fgdb) == 0 && _ggbdb(_cfefd[_gdeda[0]], _cfefd[_gdeda[1]])) || (len(_fgdb) == 2 && len(_gdeda) == 0 && _fdedf(_cfefd[_fgdb[0]], _cfefd[_fgdb[1]]))
	if _afbec {
		_cfe.Printf(" \u0020\u0020\u0068\u006f\u0072\u007as\u003d\u0025\u0064\u0020\u0076\u0065\u0072\u0074\u0073=\u0025\u0064\u0020o\u006b=\u0025\u0074\u000a", len(_gdeda), len(_fgdb), _dfcg)
	}
	if !_dfcg {
		if _afbec {
			_ed.Log.Error("\u0021!\u006d\u0061\u006b\u0065R\u0065\u0063\u0074\u0052\u0075l\u0069n\u0067:\u0020\u0070\u0061\u0074\u0068\u003d\u0025v", _gfbe)
			_cfe.Printf(" \u0020\u0020\u0068\u006f\u0072\u007as\u003d\u0025\u0064\u0020\u0076\u0065\u0072\u0074\u0073=\u0025\u0064\u0020o\u006b=\u0025\u0074\u000a", len(_gdeda), len(_fgdb), _dfcg)
		}
		return &ruling{}, false
	}
	if len(_fgdb) == 0 {
		for _gbac, _cfece := range _abffc {
			if _cfece != _bdfd {
				_fgdb = append(_fgdb, _gbac)
			}
		}
	}
	if len(_gdeda) == 0 {
		for _cecg, _fgbdd := range _abffc {
			if _fgbdd != _gfafc {
				_gdeda = append(_gdeda, _cecg)
			}
		}
	}
	if _afbec {
		_ed.Log.Info("\u006da\u006b\u0065R\u0065\u0063\u0074\u0052u\u006c\u0069\u006eg\u003a\u0020\u0068\u006f\u0072\u007a\u0073\u003d\u0025d \u0076\u0065\u0072t\u0073\u003d%\u0064\u0020\u0070\u006f\u0069\u006et\u0073\u003d%\u0064\u000a"+"\u0009\u0020\u0068o\u0072\u007a\u0073\u003d\u0025\u002b\u0076\u000a"+"\u0009\u0020\u0076e\u0072\u0074\u0073\u003d\u0025\u002b\u0076\u000a"+"\t\u0070\u006f\u0069\u006e\u0074\u0073\u003d\u0025\u002b\u0076", len(_gdeda), len(_fgdb), len(_cfefd), _gdeda, _fgdb, _cfefd)
	}
	var _befgb, _deadg, _dcbc, _agdf _cc.Point
	if _cfefd[_gdeda[0]].Y > _cfefd[_gdeda[1]].Y {
		_dcbc, _agdf = _cfefd[_gdeda[0]], _cfefd[_gdeda[1]]
	} else {
		_dcbc, _agdf = _cfefd[_gdeda[1]], _cfefd[_gdeda[0]]
	}
	if _cfefd[_fgdb[0]].X > _cfefd[_fgdb[1]].X {
		_befgb, _deadg = _cfefd[_fgdb[0]], _cfefd[_fgdb[1]]
	} else {
		_befgb, _deadg = _cfefd[_fgdb[1]], _cfefd[_fgdb[0]]
	}
	_dddb := _gb.PdfRectangle{Llx: _befgb.X, Urx: _deadg.X, Lly: _agdf.Y, Ury: _dcbc.Y}
	if _dddb.Llx > _dddb.Urx {
		_dddb.Llx, _dddb.Urx = _dddb.Urx, _dddb.Llx
	}
	if _dddb.Lly > _dddb.Ury {
		_dddb.Lly, _dddb.Ury = _dddb.Ury, _dddb.Lly
	}
	_dcdf := rectRuling{PdfRectangle: _dddb, _gcaad: _ebccd(_dddb), Color: _bdaga}
	if _dcdf._gcaad == _eefaf {
		if _afbec {
			_ed.Log.Error("\u006da\u006b\u0065\u0052\u0065\u0063\u0074\u0052\u0075\u006c\u0069\u006eg\u003a\u0020\u006b\u0069\u006e\u0064\u003d\u006e\u0069\u006c")
		}
		return nil, false
	}
	_bgbeb, _gcfb := _dcdf.asRuling()
	if !_gcfb {
		if _afbec {
			_ed.Log.Error("\u006da\u006b\u0065\u0052\u0065c\u0074\u0052\u0075\u006c\u0069n\u0067:\u0020!\u0069\u0073\u0052\u0075\u006c\u0069\u006eg")
		}
		return nil, false
	}
	if _geaeg {
		_cfe.Printf("\u0020\u0020\u0020\u0072\u003d\u0025\u0073\u000a", _bgbeb.String())
	}
	return _bgbeb, true
}
func (_cbec rulingList) primaries() []float64 {
	_eeafd := make(map[float64]struct{}, len(_cbec))
	for _, _cgag := range _cbec {
		_eeafd[_cgag._abfg] = struct{}{}
	}
	_beeac := make([]float64, len(_eeafd))
	_edda := 0
	for _caec := range _eeafd {
		_beeac[_edda] = _caec
		_edda++
	}
	_d.Float64s(_beeac)
	return _beeac
}

// ExtractPageImages returns the image contents of the page extractor, including data
// and position, size information for each image.
// A set of options to control page image extraction can be passed in. The options
// parameter can be nil for the default options. By default, inline stencil masks
// are not extracted.
func (_ebg *Extractor) ExtractPageImages(options *ImageExtractOptions) (*PageImages, error) {
	_becc := &imageExtractContext{_eag: options}
	_afa := _becc.extractContentStreamImages(_ebg._ee, _ebg._bd)
	if _afa != nil {
		return nil, _afa
	}
	return &PageImages{Images: _becc._cdb}, nil
}
func _bbfg(_ecdfc *textLine, _dgegc []*textLine, _gegf []float64, _agdb, _bdcae float64) []*textLine {
	_gcde := []*textLine{}
	for _, _gcgd := range _dgegc {
		if _gcgd._bdbg >= _agdb {
			if _bdcae != -1 && _gcgd._bdbg < _bdcae {
				if _gcgd.text() != _ecdfc.text() {
					if _df.Round(_gcgd.Llx) < _df.Round(_ecdfc.Llx) {
						break
					}
					_gcde = append(_gcde, _gcgd)
				}
			} else if _bdcae == -1 {
				if _gcgd._bdbg == _ecdfc._bdbg {
					if _gcgd.text() != _ecdfc.text() {
						_gcde = append(_gcde, _gcgd)
					}
					continue
				}
				_agcdg := _gecd(_ecdfc, _dgegc, _gegf)
				if _agcdg != -1 && _gcgd._bdbg <= _agcdg {
					_gcde = append(_gcde, _gcgd)
				}
			}
		}
	}
	return _gcde
}

// Append appends `mark` to the mark array.
func (_ccbdb *TextMarkArray) Append(mark TextMark) { _ccbdb._ggcf = append(_ccbdb._ggcf, mark) }
func _gegag(_aead []TextMark, _dfef *int, _eggc TextMark) []TextMark {
	_eggc.Offset = *_dfef
	_aead = append(_aead, _eggc)
	*_dfef += len(_eggc.Text)
	return _aead
}
func _ccefg(_gfgbb, _egcc float64) bool { return _df.Abs(_gfgbb-_egcc) <= _egaa }

// Marks returns the TextMark collection for a page. It represents all the text on the page.
func (_dfce PageText) Marks() *TextMarkArray { return &TextMarkArray{_ggcf: _dfce._afc} }

// Text gets the extracted text contained in `l`.
func (_fbfb *list) Text() string {
	_cffc := &_ag.Builder{}
	_eacf := ""
	_befg(_fbfb, _cffc, &_eacf)
	return _cffc.String()
}
func (_bbac *textTable) toTextTable() TextTable {
	if _geff {
		_ed.Log.Info("t\u006fT\u0065\u0078\u0074\u0054\u0061\u0062\u006c\u0065:\u0020\u0025\u0064\u0020x \u0025\u0064", _bbac._fccgee, _bbac._dege)
	}
	_ggffb := make([][]TableCell, _bbac._dege)
	for _beda := 0; _beda < _bbac._dege; _beda++ {
		_ggffb[_beda] = make([]TableCell, _bbac._fccgee)
		for _acbf := 0; _acbf < _bbac._fccgee; _acbf++ {
			_fbbb := _bbac.get(_acbf, _beda)
			if _fbbb == nil {
				continue
			}
			if _geff {
				_cfe.Printf("\u0025\u0034\u0064 \u0025\u0032\u0064\u003a\u0020\u0025\u0073\u000a", _acbf, _beda, _fbbb)
			}
			_ggffb[_beda][_acbf].Text = _fbbb.text()
			_adca := 0
			_ggffb[_beda][_acbf].Marks._ggcf = _fbbb.toTextMarks(&_adca)
		}
	}
	_fgcc := TextTable{W: _bbac._fccgee, H: _bbac._dege, Cells: _ggffb}
	_fgcc.PdfRectangle = _bbac.bbox()
	return _fgcc
}
func _dgdc(_fdff _gb.PdfRectangle, _aaffb bounded) float64 { return _fdff.Ury - _aaffb.bbox().Lly }
func (_eeaf *stateStack) push(_abf *textState)             { _gbeg := *_abf; *_eeaf = append(*_eeaf, &_gbeg) }

var _dfbag = _ca.MustCompile("\u005e\u005c\u0073\u002a\u0028\u005c\u0064\u002b\u005c\u002e\u003f|\u005b\u0049\u0069\u0076\u005d\u002b\u0029\u005c\u0073\u002a\\\u0029\u003f\u0024")

func _fgege(_gbegb []*textMark, _efabfb _gb.PdfRectangle) []*textWord {
	var _aeeddb []*textWord
	var _gcccd *textWord
	if _gbbcc {
		_ed.Log.Info("\u006d\u0061\u006beT\u0065\u0078\u0074\u0057\u006f\u0072\u0064\u0073\u003a\u0020\u0025\u0064\u0020\u006d\u0061\u0072\u006b\u0073", len(_gbegb))
	}
	_bgaa := func() {
		if _gcccd != nil {
			_gddbb := _gcccd.computeText()
			if !_edefac(_gddbb) {
				_gcccd._bcfea = _gddbb
				_aeeddb = append(_aeeddb, _gcccd)
				if _gbbcc {
					_ed.Log.Info("\u0061\u0064\u0064Ne\u0077\u0057\u006f\u0072\u0064\u003a\u0020\u0025\u0064\u003a\u0020\u0077\u006f\u0072\u0064\u003d\u0025\u0073", len(_aeeddb)-1, _gcccd.String())
					for _aacf, _gdda := range _gcccd._fdgbc {
						_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _aacf, _gdda.String())
					}
				}
			}
			_gcccd = nil
		}
	}
	for _, _bffd := range _gbegb {
		if _dbfc && _gcccd != nil && len(_gcccd._fdgbc) > 0 {
			_cebf := _gcccd._fdgbc[len(_gcccd._fdgbc)-1]
			_bbfgd, _gbebb := _fgdf(_bffd._ceca)
			_daefc, _ffege := _fgdf(_cebf._ceca)
			if _gbebb && !_ffege && _cebf.inDiacriticArea(_bffd) {
				_gcccd.addDiacritic(_bbfgd)
				continue
			}
			if _ffege && !_gbebb && _bffd.inDiacriticArea(_cebf) {
				_gcccd._fdgbc = _gcccd._fdgbc[:len(_gcccd._fdgbc)-1]
				_gcccd.appendMark(_bffd, _efabfb)
				_gcccd.addDiacritic(_daefc)
				continue
			}
		}
		_gfba := _edefac(_bffd._ceca)
		if _gfba {
			_bgaa()
			continue
		}
		if _gcccd == nil && !_gfba {
			_gcccd = _bbcee([]*textMark{_bffd}, _efabfb)
			continue
		}
		_accbf := _gcccd._fgdbd
		_afgf := _df.Abs(_dgdc(_efabfb, _bffd)-_gcccd._cbfcee) / _accbf
		_begb := _daefa(_bffd, _gcccd) / _accbf
		if _begb >= _cgc || !(-_gbdc <= _begb && _afgf <= _cedace) {
			_bgaa()
			_gcccd = _bbcee([]*textMark{_bffd}, _efabfb)
			continue
		}
		_gcccd.appendMark(_bffd, _efabfb)
	}
	_bgaa()
	return _aeeddb
}
func (_abbbf *textTable) logComposite(_gdff string) {
	if !_geff {
		return
	}
	_ed.Log.Info("\u007e~\u007eP\u0061\u0072\u0061\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0073", _abbbf._fccgee, _abbbf._dege, _gdff)
	_cfe.Printf("\u0025\u0035\u0073 \u007c", "")
	for _bedfa := 0; _bedfa < _abbbf._fccgee; _bedfa++ {
		_cfe.Printf("\u0025\u0033\u0064 \u007c", _bedfa)
	}
	_cfe.Println("")
	_cfe.Printf("\u0025\u0035\u0073 \u002b", "")
	for _fged := 0; _fged < _abbbf._fccgee; _fged++ {
		_cfe.Printf("\u0025\u0033\u0073 \u002b", "\u002d\u002d\u002d")
	}
	_cfe.Println("")
	for _bagd := 0; _bagd < _abbbf._dege; _bagd++ {
		_cfe.Printf("\u0025\u0035\u0064 \u007c", _bagd)
		for _fabe := 0; _fabe < _abbbf._fccgee; _fabe++ {
			_eecd, _ := _abbbf._faeg[_aafb(_fabe, _bagd)].parasBBox()
			_cfe.Printf("\u0025\u0033\u0064 \u007c", len(_eecd))
		}
		_cfe.Println("")
	}
	_ed.Log.Info("\u007e~\u007eT\u0065\u0078\u0074\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u0025\u0073", _abbbf._fccgee, _abbbf._dege, _gdff)
	_cfe.Printf("\u0025\u0035\u0073 \u007c", "")
	for _cfgb := 0; _cfgb < _abbbf._fccgee; _cfgb++ {
		_cfe.Printf("\u0025\u0031\u0032\u0064\u0020\u007c", _cfgb)
	}
	_cfe.Println("")
	_cfe.Printf("\u0025\u0035\u0073 \u002b", "")
	for _bgga := 0; _bgga < _abbbf._fccgee; _bgga++ {
		_cfe.Print("\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d\u002d-\u002d\u002d\u002d\u002b")
	}
	_cfe.Println("")
	for _ecgb := 0; _ecgb < _abbbf._dege; _ecgb++ {
		_cfe.Printf("\u0025\u0035\u0064 \u007c", _ecgb)
		for _ggfe := 0; _ggfe < _abbbf._fccgee; _ggfe++ {
			_agdgf, _ := _abbbf._faeg[_aafb(_ggfe, _ecgb)].parasBBox()
			_eceda := ""
			_adee := _agdgf.merge()
			if _adee != nil {
				_eceda = _adee.text()
			}
			_eceda = _cfe.Sprintf("\u0025\u0071", _fabf(_eceda, 12))
			_eceda = _eceda[1 : len(_eceda)-1]
			_cfe.Printf("\u0025\u0031\u0032\u0073\u0020\u007c", _eceda)
		}
		_cfe.Println("")
	}
}
func _gbga(_becba []*textLine) map[float64][]*textLine {
	_d.Slice(_becba, func(_deeb, _bgcd int) bool { return _becba[_deeb]._bdbg < _becba[_bgcd]._bdbg })
	_dfbbb := map[float64][]*textLine{}
	for _, _edfce := range _becba {
		_bgced := _dga(_edfce)
		_bgced = _df.Round(_bgced)
		_dfbbb[_bgced] = append(_dfbbb[_bgced], _edfce)
	}
	return _dfbbb
}
func (_gccad *textWord) bbox() _gb.PdfRectangle { return _gccad.PdfRectangle }

var (
	_edfcef = map[rune]string{0x0060: "\u0300", 0x02CB: "\u0300", 0x0027: "\u0301", 0x00B4: "\u0301", 0x02B9: "\u0301", 0x02CA: "\u0301", 0x005E: "\u0302", 0x02C6: "\u0302", 0x007E: "\u0303", 0x02DC: "\u0303", 0x00AF: "\u0304", 0x02C9: "\u0304", 0x02D8: "\u0306", 0x02D9: "\u0307", 0x00A8: "\u0308", 0x00B0: "\u030a", 0x02DA: "\u030a", 0x02BA: "\u030b", 0x02DD: "\u030b", 0x02C7: "\u030c", 0x02C8: "\u030d", 0x0022: "\u030e", 0x02BB: "\u0312", 0x02BC: "\u0313", 0x0486: "\u0313", 0x055A: "\u0313", 0x02BD: "\u0314", 0x0485: "\u0314", 0x0559: "\u0314", 0x02D4: "\u031d", 0x02D5: "\u031e", 0x02D6: "\u031f", 0x02D7: "\u0320", 0x02B2: "\u0321", 0x00B8: "\u0327", 0x02CC: "\u0329", 0x02B7: "\u032b", 0x02CD: "\u0331", 0x005F: "\u0332", 0x204E: "\u0359"}
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
	BBox _gb.PdfRectangle

	// Font is the font the text was drawn with.
	Font *_gb.PdfFont

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
	FillColor _eb.Color

	// StrokeColor is the stroke color of the text.
	// The color is nil for spaces and line breaks (i.e. the Meta field is true).
	StrokeColor _eb.Color

	// Orientation is the text orientation
	Orientation int

	// DirectObject is the underlying PdfObject (Text Object) that represents the visible texts. This is introduced to get
	// a simple access to the TextObject in case editing or replacment of some text is needed. E.g during redaction.
	DirectObject _gg.PdfObject

	// ObjString is a decoded string operand of a text-showing operator. It has the same value as `Text` attribute except
	// when many glyphs are represented with the same Text Object that contains multiple length string operand in which case
	// ObjString spans more than one character string that falls in different TextMark objects.
	ObjString []string
	Tw        float64
	Th        float64
	Tc        float64
	Index     int
	_fefe     bool
	_egc      *TextTable
}

func (_aba *textLine) markWordBoundaries() {
	_cceb := _bacc * _aba._egbd
	for _cgcb, _ceef := range _aba._bfca[1:] {
		if _daefa(_ceef, _aba._bfca[_cgcb]) >= _cceb {
			_ceef._cacca = true
		}
	}
}
func (_ebge *wordBag) empty(_cccg int) bool { _, _fege := _ebge._gegc[_cccg]; return !_fege }
func _dcebc(_cede string) bool {
	if _cb.RuneCountInString(_cede) < _ddcf {
		return false
	}
	_gacfg, _deaf := _cb.DecodeLastRuneInString(_cede)
	if _deaf <= 0 || !_ac.Is(_ac.Hyphen, _gacfg) {
		return false
	}
	_gacfg, _deaf = _cb.DecodeLastRuneInString(_cede[:len(_cede)-_deaf])
	return _deaf > 0 && !_ac.IsSpace(_gacfg)
}
func (_gfgac *ruling) gridIntersecting(_bega *ruling) bool {
	return _ccefg(_gfgac._daag, _bega._daag) && _ccefg(_gfgac._fcff, _bega._fcff)
}
func _egcd(_ddaa, _caeeg _gb.PdfRectangle) bool {
	return _ddaa.Llx <= _caeeg.Llx && _caeeg.Urx <= _ddaa.Urx && _ddaa.Lly <= _caeeg.Lly && _caeeg.Ury <= _ddaa.Ury
}
func _edfbf(_bgcg int, _caad func(int, int) bool) []int {
	_fdea := make([]int, _bgcg)
	for _eedb := range _fdea {
		_fdea[_eedb] = _eedb
	}
	_d.Slice(_fdea, func(_eadc, _dcfg int) bool { return _caad(_fdea[_eadc], _fdea[_dcfg]) })
	return _fdea
}
func (_defc *wordBag) firstReadingIndex(_abgbb int) int {
	_ccgg := _defc.firstWord(_abgbb)._fgdbd
	_cagf := float64(_abgbb+1) * _edagd
	_dagc := _cagf + _cfab*_ccgg
	_eeca := _abgbb
	for _, _gefd := range _defc.depthBand(_cagf, _dagc) {
		if _fdd(_defc.firstWord(_gefd), _defc.firstWord(_eeca)) < 0 {
			_eeca = _gefd
		}
	}
	return _eeca
}

// String returns a description of `k`.
func (_cadb markKind) String() string {
	_bcaf, _ccfba := _eebbd[_cadb]
	if !_ccfba {
		return _cfe.Sprintf("\u004e\u006f\u0074\u0020\u0061\u0020\u006d\u0061\u0072k\u003a\u0020\u0025\u0064", _cadb)
	}
	return _bcaf
}
func (_faggb rulingList) mergePrimary() float64 {
	_fcac := _faggb[0]._abfg
	for _, _gdae := range _faggb[1:] {
		_fcac += _gdae._abfg
	}
	return _fcac / float64(len(_faggb))
}
func (_cda *textObject) moveLP(_gdfa, _gafc float64) {
	_cda._ddge.Concat(_cc.NewMatrix(1, 0, 0, 1, _gdfa, _gafc))
	_cda._ffb = _cda._ddge
}
func _afbf(_agdc string, _eacag []rulingList) {
	_ed.Log.Info("\u0024\u0024 \u0025\u0064\u0020g\u0072\u0069\u0064\u0073\u0020\u002d\u0020\u0025\u0073", len(_eacag), _agdc)
	for _gafdf, _ddfga := range _eacag {
		_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _gafdf, _ddfga.String())
	}
}

// New returns an Extractor instance for extracting content from the input PDF page.
func New(page *_gb.PdfPage) (*Extractor, error) { return NewWithOptions(page, nil) }
func (_ccef rulingList) augmentGrid() (rulingList, rulingList) {
	_ggeg, _dcac := _ccef.vertsHorzs()
	if len(_ggeg) == 0 || len(_dcac) == 0 {
		return _ggeg, _dcac
	}
	_efabcf, _aded := _ggeg, _dcac
	_bfee := _ggeg.bbox()
	_agcbd := _dcac.bbox()
	if _geaeg {
		_ed.Log.Info("\u0061u\u0067\u006d\u0065\u006e\u0074\u0047\u0072\u0069\u0064\u003a\u0020b\u0062\u006f\u0078\u0056\u003d\u0025\u0036\u002e\u0032\u0066", _bfee)
		_ed.Log.Info("\u0061u\u0067\u006d\u0065\u006e\u0074\u0047\u0072\u0069\u0064\u003a\u0020b\u0062\u006f\u0078\u0048\u003d\u0025\u0036\u002e\u0032\u0066", _agcbd)
	}
	var _gegage, _faeegg, _baec, _becbag *ruling
	if _agcbd.Llx < _bfee.Llx-_egaa {
		_gegage = &ruling{_bbge: _ecdcf, _cffe: _gfafc, _abfg: _agcbd.Llx, _daag: _bfee.Lly, _fcff: _bfee.Ury}
		_ggeg = append(rulingList{_gegage}, _ggeg...)
	}
	if _agcbd.Urx > _bfee.Urx+_egaa {
		_faeegg = &ruling{_bbge: _ecdcf, _cffe: _gfafc, _abfg: _agcbd.Urx, _daag: _bfee.Lly, _fcff: _bfee.Ury}
		_ggeg = append(_ggeg, _faeegg)
	}
	if _bfee.Lly < _agcbd.Lly-_egaa {
		_baec = &ruling{_bbge: _ecdcf, _cffe: _bdfd, _abfg: _bfee.Lly, _daag: _agcbd.Llx, _fcff: _agcbd.Urx}
		_dcac = append(rulingList{_baec}, _dcac...)
	}
	if _bfee.Ury > _agcbd.Ury+_egaa {
		_becbag = &ruling{_bbge: _ecdcf, _cffe: _bdfd, _abfg: _bfee.Ury, _daag: _agcbd.Llx, _fcff: _agcbd.Urx}
		_dcac = append(_dcac, _becbag)
	}
	if len(_ggeg)+len(_dcac) == len(_ccef) {
		return _efabcf, _aded
	}
	_cagg := append(_ggeg, _dcac...)
	_ccef.log("u\u006e\u0061\u0075\u0067\u006d\u0065\u006e\u0074\u0065\u0064")
	_cagg.log("\u0061u\u0067\u006d\u0065\u006e\u0074\u0065d")
	return _ggeg, _dcac
}
func _edefac(_deafac string) bool {
	for _, _dbed := range _deafac {
		if !_ac.IsSpace(_dbed) {
			return false
		}
	}
	return true
}
func _ggbdb(_efed, _befad _cc.Point) bool {
	_fadd := _df.Abs(_efed.X - _befad.X)
	_ffae := _df.Abs(_efed.Y - _befad.Y)
	return _dffe(_ffae, _fadd)
}
func _dfcee(_ffbga, _eaae, _acgba, _gaabf *textPara) *textTable {
	_bacd := &textTable{_fccgee: 2, _dege: 2, _adda: make(map[uint64]*textPara, 4)}
	_bacd.put(0, 0, _ffbga)
	_bacd.put(1, 0, _eaae)
	_bacd.put(0, 1, _acgba)
	_bacd.put(1, 1, _gaabf)
	return _bacd
}
func (_bff *subpath) isQuadrilateral() bool {
	if len(_bff._ccgf) < 4 || len(_bff._ccgf) > 5 {
		return false
	}
	if len(_bff._ccgf) == 5 {
		_fgfdbc := _bff._ccgf[0]
		_fage := _bff._ccgf[4]
		if _fgfdbc.X != _fage.X || _fgfdbc.Y != _fage.Y {
			return false
		}
	}
	return true
}
func (_fcdgd *shapesState) addPoint(_ddbg, _dbcc float64) {
	_efbeb := _fcdgd.establishSubpath()
	_cgde := _fcdgd.devicePoint(_ddbg, _dbcc)
	if _efbeb == nil {
		_fcdgd._dfff = true
		_fcdgd._bcb = _cgde
	} else {
		_efbeb.add(_cgde)
	}
}
func (_ccgee rulingList) snapToGroups() rulingList {
	_ggfg, _dccgb := _ccgee.vertsHorzs()
	if len(_ggfg) > 0 {
		_ggfg = _ggfg.snapToGroupsDirection()
	}
	if len(_dccgb) > 0 {
		_dccgb = _dccgb.snapToGroupsDirection()
	}
	_ggdad := append(_ggfg, _dccgb...)
	_ggdad.log("\u0073\u006e\u0061p\u0054\u006f\u0047\u0072\u006f\u0075\u0070\u0073")
	return _ggdad
}
func _bgc(_gbcb *_aad.ContentStreamOperation) (float64, error) {
	if len(_gbcb.Params) != 1 {
		_ddg := _c.New("\u0069n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0070\u0061r\u0061m\u0065t\u0065\u0072\u0020\u0063\u006f\u0075\u006et")
		_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u0023\u0071\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020h\u0061\u0076\u0065\u0020\u0025\u0064\u0020i\u006e\u0070\u0075\u0074\u0020\u0070\u0061\u0072\u0061\u006d\u0073,\u0020\u0067\u006f\u0074\u0020\u0025\u0064\u0020\u0025\u002b\u0076", _gbcb.Operand, 1, len(_gbcb.Params), _gbcb.Params)
		return 0.0, _ddg
	}
	return _gg.GetNumberAsFloat(_gbcb.Params[0])
}

type structTreeRoot struct {
	_gafdd []structElement
	_bgee  string
}

func (_gfa *subpath) add(_cffa ..._cc.Point) { _gfa._ccgf = append(_gfa._ccgf, _cffa...) }
func (_aed *textObject) showTextAdjusted(_gcaf *_gg.PdfObjectArray, _edfc int) error {
	_afge := false
	for _, _ceg := range _gcaf.Elements() {
		switch _ceg.(type) {
		case *_gg.PdfObjectFloat, *_gg.PdfObjectInteger:
			_dgb, _aecc := _gg.GetNumberAsFloat(_ceg)
			if _aecc != nil {
				_ed.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0073\u0068\u006f\u0077\u0054\u0065\u0078t\u0041\u0064\u006a\u0075\u0073\u0074\u0065\u0064\u002e\u0020\u0042\u0061\u0064\u0020\u006e\u0075\u006d\u0065r\u0069\u0063\u0061\u006c\u0020a\u0072\u0067\u002e\u0020\u006f\u003d\u0025\u0073\u0020\u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _ceg, _gcaf)
				return _aecc
			}
			_dfg, _badb := -_dgb*0.001*_aed._cbag._agge, 0.0
			if _afge {
				_badb, _dfg = _dfg, _badb
			}
			_cbgg := _gggb(_cc.Point{X: _dfg, Y: _badb})
			_aed._ffb.Concat(_cbgg)
		case *_gg.PdfObjectString:
			_ecda := _gg.TraceToDirectObject(_ceg)
			_efb, _cbed := _gg.GetStringBytes(_ecda)
			if !_cbed {
				_ed.Log.Trace("s\u0068\u006f\u0077\u0054\u0065\u0078\u0074\u0041\u0064j\u0075\u0073\u0074\u0065\u0064\u003a\u0020Ba\u0064\u0020\u0073\u0074r\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u002e\u0020o=\u0025\u0073 \u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _ceg, _gcaf)
				return _gg.ErrTypeError
			}
			_aed.renderText(_ecda, _efb, _edfc)
		default:
			_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0073\u0068\u006f\u0077\u0054\u0065\u0078\u0074A\u0064\u006a\u0075\u0073\u0074\u0065\u0064\u002e\u0020\u0055\u006e\u0065\u0078p\u0065\u0063\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0028%T\u0029\u0020\u0061\u0072\u0067\u0073\u003d\u0025\u002b\u0076", _ceg, _gcaf)
			return _gg.ErrTypeError
		}
	}
	return nil
}
func (_ffad *textObject) checkOp(_eeba *_aad.ContentStreamOperation, _cggc int, _dbc bool) (_dbga bool, _dad error) {
	if _ffad == nil {
		var _dfdc []_gg.PdfObject
		if _cggc > 0 {
			_dfdc = _eeba.Params
			if len(_dfdc) > _cggc {
				_dfdc = _dfdc[:_cggc]
			}
		}
		_ed.Log.Debug("\u0025\u0023q \u006f\u0070\u0065r\u0061\u006e\u0064\u0020out\u0073id\u0065\u0020\u0074\u0065\u0078\u0074\u002e p\u0061\u0072\u0061\u006d\u0073\u003d\u0025+\u0076", _eeba.Operand, _dfdc)
	}
	if _cggc >= 0 {
		if len(_eeba.Params) != _cggc {
			if _dbc {
				_dad = _c.New("\u0069n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0070\u0061r\u0061m\u0065t\u0065\u0072\u0020\u0063\u006f\u0075\u006et")
			}
			_ed.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u0023\u0071\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020h\u0061\u0076\u0065\u0020\u0025\u0064\u0020i\u006e\u0070\u0075\u0074\u0020\u0070\u0061\u0072\u0061\u006d\u0073,\u0020\u0067\u006f\u0074\u0020\u0025\u0064\u0020\u0025\u002b\u0076", _eeba.Operand, _cggc, len(_eeba.Params), _eeba.Params)
			return false, _dad
		}
	}
	return true, nil
}
func (_bcabd lineRuling) asRuling() (*ruling, bool) {
	_eadg := ruling{_cffe: _bcabd._ecee, Color: _bcabd.Color, _bbge: _fdacf}
	switch _bcabd._ecee {
	case _gfafc:
		_eadg._abfg = _bcabd.xMean()
		_eadg._daag = _df.Min(_bcabd._efffg.Y, _bcabd._agae.Y)
		_eadg._fcff = _df.Max(_bcabd._efffg.Y, _bcabd._agae.Y)
	case _bdfd:
		_eadg._abfg = _bcabd.yMean()
		_eadg._daag = _df.Min(_bcabd._efffg.X, _bcabd._agae.X)
		_eadg._fcff = _df.Max(_bcabd._efffg.X, _bcabd._agae.X)
	default:
		_ed.Log.Error("\u0062\u0061\u0064\u0020pr\u0069\u006d\u0061\u0072\u0079\u0020\u006b\u0069\u006e\u0064\u003d\u0025\u0064", _bcabd._ecee)
		return nil, false
	}
	return &_eadg, true
}
func (_bbgc intSet) add(_afdc int) { _bbgc[_afdc] = struct{}{} }

// PageText represents the layout of text on a device page.
type PageText struct {
	_bagf []*textMark
	_ccbd string
	_afc  []TextMark
	_ded  []TextTable
	_ebae _gb.PdfRectangle
	_eff  []pathSection
	_efbe []pathSection
	_ddbd *_gg.PdfObject
	_ddeg _gg.PdfObject
	_dff  *_aad.ContentStreamOperations
	_ege  PageTextOptions
}
type pathSection struct {
	_ggfc []*subpath
	_eb.Color
}

func (_ecbf paraList) llyOrdering() []int {
	_cgfa := make([]int, len(_ecbf))
	for _ebbcc := range _ecbf {
		_cgfa[_ebbcc] = _ebbcc
	}
	_d.SliceStable(_cgfa, func(_afgg, _adga int) bool {
		_bebab, _dgbf := _cgfa[_afgg], _cgfa[_adga]
		return _ecbf[_bebab].Lly < _ecbf[_dgbf].Lly
	})
	return _cgfa
}
func (_afgd *wordBag) removeWord(_ecga *textWord, _fecg int) {
	_gedg := _afgd._gegc[_fecg]
	_gedg = _abdc(_gedg, _ecga)
	if len(_gedg) == 0 {
		delete(_afgd._gegc, _fecg)
	} else {
		_afgd._gegc[_fecg] = _gedg
	}
}
func (_edacdf *ruling) equals(_fddeb *ruling) bool {
	return _edacdf._cffe == _fddeb._cffe && _ccefg(_edacdf._abfg, _fddeb._abfg) && _ccefg(_edacdf._daag, _fddeb._daag) && _ccefg(_edacdf._fcff, _fddeb._fcff)
}
func (_bcfde *textTable) putComposite(_ecfcaa, _ggfd int, _abbef paraList, _eagf _gb.PdfRectangle) {
	if len(_abbef) == 0 {
		_ed.Log.Error("\u0074\u0065xt\u0054\u0061\u0062l\u0065\u0029\u0020\u0070utC\u006fmp\u006f\u0073\u0069\u0074\u0065\u003a\u0020em\u0070\u0074\u0079\u0020\u0070\u0061\u0072a\u0073")
		return
	}
	_efde := compositeCell{PdfRectangle: _eagf, paraList: _abbef}
	if _geff {
		_cfe.Printf("\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0020\u0070\u0075\u0074\u0043\u006f\u006d\u0070o\u0073i\u0074\u0065\u0028\u0025\u0064\u002c\u0025\u0064\u0029\u003c\u002d\u0025\u0073\u000a", _ecfcaa, _ggfd, _efde.String())
	}
	_efde.updateBBox()
	_bcfde._faeg[_aafb(_ecfcaa, _ggfd)] = _efde
}
func (_eada *textObject) getStrokeColor() _eb.Color {
	return _fbfab(_eada._ccg.ColorspaceStroking, _eada._ccg.ColorStroking)
}
func (_daac paraList) addNeighbours() {
	_ecfad := func(_ffeb []int, _bbbc *textPara) ([]*textPara, []*textPara) {
		_gdac := make([]*textPara, 0, len(_ffeb)-1)
		_dgaa := make([]*textPara, 0, len(_ffeb)-1)
		for _, _cbbc := range _ffeb {
			_gfadb := _daac[_cbbc]
			if _gfadb.Urx <= _bbbc.Llx {
				_gdac = append(_gdac, _gfadb)
			} else if _gfadb.Llx >= _bbbc.Urx {
				_dgaa = append(_dgaa, _gfadb)
			}
		}
		return _gdac, _dgaa
	}
	_dcbed := func(_bcdgdc []int, _gagaa *textPara) ([]*textPara, []*textPara) {
		_egaad := make([]*textPara, 0, len(_bcdgdc)-1)
		_cbcef := make([]*textPara, 0, len(_bcdgdc)-1)
		for _, _cgace := range _bcdgdc {
			_cfcff := _daac[_cgace]
			if _cfcff.Ury <= _gagaa.Lly {
				_cbcef = append(_cbcef, _cfcff)
			} else if _cfcff.Lly >= _gagaa.Ury {
				_egaad = append(_egaad, _cfcff)
			}
		}
		return _egaad, _cbcef
	}
	_fcce := _daac.yNeighbours(_gbea)
	for _, _eccdd := range _daac {
		_fbfaa := _fcce[_eccdd]
		if len(_fbfaa) == 0 {
			continue
		}
		_bgfb, _becddb := _ecfad(_fbfaa, _eccdd)
		if len(_bgfb) == 0 && len(_becddb) == 0 {
			continue
		}
		if len(_bgfb) > 0 {
			_gcdgd := _bgfb[0]
			for _, _dfdce := range _bgfb[1:] {
				if _dfdce.Urx >= _gcdgd.Urx {
					_gcdgd = _dfdce
				}
			}
			for _, _efbcd := range _bgfb {
				if _efbcd != _gcdgd && _efbcd.Urx > _gcdgd.Llx {
					_gcdgd = nil
					break
				}
			}
			if _gcdgd != nil && _ebgc(_eccdd.PdfRectangle, _gcdgd.PdfRectangle) {
				_eccdd._dgdb = _gcdgd
			}
		}
		if len(_becddb) > 0 {
			_bfegd := _becddb[0]
			for _, _fcece := range _becddb[1:] {
				if _fcece.Llx <= _bfegd.Llx {
					_bfegd = _fcece
				}
			}
			for _, _gedd := range _becddb {
				if _gedd != _bfegd && _gedd.Llx < _bfegd.Urx {
					_bfegd = nil
					break
				}
			}
			if _bfegd != nil && _ebgc(_eccdd.PdfRectangle, _bfegd.PdfRectangle) {
				_eccdd._ccdg = _bfegd
			}
		}
	}
	_fcce = _daac.xNeighbours(_agcd)
	for _, _cgdg := range _daac {
		_ggdde := _fcce[_cgdg]
		if len(_ggdde) == 0 {
			continue
		}
		_bced, _gfgfc := _dcbed(_ggdde, _cgdg)
		if len(_bced) == 0 && len(_gfgfc) == 0 {
			continue
		}
		if len(_gfgfc) > 0 {
			_eace := _gfgfc[0]
			for _, _bbgee := range _gfgfc[1:] {
				if _bbgee.Ury >= _eace.Ury {
					_eace = _bbgee
				}
			}
			for _, _bdefc := range _gfgfc {
				if _bdefc != _eace && _bdefc.Ury > _eace.Lly {
					_eace = nil
					break
				}
			}
			if _eace != nil && _cdcd(_cgdg.PdfRectangle, _eace.PdfRectangle) {
				_cgdg._cgfe = _eace
			}
		}
		if len(_bced) > 0 {
			_abdff := _bced[0]
			for _, _gcggf := range _bced[1:] {
				if _gcggf.Lly <= _abdff.Lly {
					_abdff = _gcggf
				}
			}
			for _, _aeegf := range _bced {
				if _aeegf != _abdff && _aeegf.Lly < _abdff.Ury {
					_abdff = nil
					break
				}
			}
			if _abdff != nil && _cdcd(_cgdg.PdfRectangle, _abdff.PdfRectangle) {
				_cgdg._dbecf = _abdff
			}
		}
	}
	for _, _cgfc := range _daac {
		if _cgfc._dgdb != nil && _cgfc._dgdb._ccdg != _cgfc {
			_cgfc._dgdb = nil
		}
		if _cgfc._dbecf != nil && _cgfc._dbecf._cgfe != _cgfc {
			_cgfc._dbecf = nil
		}
		if _cgfc._ccdg != nil && _cgfc._ccdg._dgdb != _cgfc {
			_cgfc._ccdg = nil
		}
		if _cgfc._cgfe != nil && _cgfc._cgfe._dbecf != _cgfc {
			_cgfc._cgfe = nil
		}
	}
}
func _daae(_dafd _gb.PdfRectangle, _adbe []*textLine) *textPara {
	return &textPara{PdfRectangle: _dafd, _abbe: _adbe}
}
func (_gaac paraList) findTableGrid(_egbe gridTiling) (*textTable, map[*textPara]struct{}) {
	_fdcga := len(_egbe._edfg)
	_ceee := len(_egbe._abbf)
	_bagbf := textTable{_baccf: true, _fccgee: _fdcga, _dege: _ceee, _adda: make(map[uint64]*textPara, _fdcga*_ceee), _faeg: make(map[uint64]compositeCell, _fdcga*_ceee)}
	_bagbf.PdfRectangle = _egbe.PdfRectangle
	_dfgg := make(map[*textPara]struct{})
	_fcdcc := int((1.0 - _dgcca) * float64(_fdcga*_ceee))
	_ggcfc := 0
	if _aaga {
		_ed.Log.Info("\u0066\u0069\u006e\u0064Ta\u0062\u006c\u0065\u0047\u0072\u0069\u0064\u003a\u0020\u0025\u0064\u0020\u0078\u0020%\u0064", _fdcga, _ceee)
	}
	for _bcdd, _baaeg := range _egbe._abbf {
		_gbdb, _daedc := _egbe._aaef[_baaeg]
		if !_daedc {
			continue
		}
		for _acbbe, _ecab := range _egbe._edfg {
			_gcbbfg, _fgad := _gbdb[_ecab]
			if !_fgad {
				continue
			}
			_dggac := _gaac.inTile(_gcbbfg)
			if len(_dggac) == 0 {
				_ggcfc++
				if _ggcfc > _fcdcc {
					if _aaga {
						_ed.Log.Info("\u0021\u006e\u0075m\u0045\u006d\u0070\u0074\u0079\u003d\u0025\u0064", _ggcfc)
					}
					return nil, nil
				}
			} else {
				_bagbf.putComposite(_acbbe, _bcdd, _dggac, _gcbbfg.PdfRectangle)
				for _, _efffd := range _dggac {
					_dfgg[_efffd] = struct{}{}
				}
			}
		}
	}
	_fgbde := 0
	for _cfcc := 0; _cfcc < _fdcga; _cfcc++ {
		_aegf := _bagbf.get(_cfcc, 0)
		if _aegf == nil || !_aegf._befd {
			_fgbde++
		}
	}
	if _fgbde == 0 {
		if _aaga {
			_ed.Log.Info("\u0021\u006e\u0075m\u0048\u0065\u0061\u0064\u0065\u0072\u003d\u0030")
		}
		return nil, nil
	}
	_cdcfa := _bagbf.reduceTiling(_egbe, _fbgec)
	_cdcfa = _cdcfa.subdivide()
	return _cdcfa, _dfgg
}
func (_cedbf *textPara) taken() bool { return _cedbf == nil || _cedbf._ggdd }
func _eccca(_cgfbg map[float64]map[float64]gridTile) []float64 {
	_cegcg := make([]float64, 0, len(_cgfbg))
	for _dded := range _cgfbg {
		_cegcg = append(_cegcg, _dded)
	}
	_d.Float64s(_cegcg)
	_abag := len(_cegcg)
	for _dcgb := 0; _dcgb < _abag/2; _dcgb++ {
		_cegcg[_dcgb], _cegcg[_abag-1-_dcgb] = _cegcg[_abag-1-_dcgb], _cegcg[_dcgb]
	}
	return _cegcg
}
func _ebccd(_degbf _gb.PdfRectangle) rulingKind {
	_dffb := _degbf.Width()
	_afeb := _degbf.Height()
	if _dffb > _afeb {
		if _dffb >= _caf {
			return _bdfd
		}
	} else {
		if _afeb >= _caf {
			return _gfafc
		}
	}
	return _eefaf
}
func _bcaef(_babb []float64, _gbgc, _aeedd float64) []float64 {
	_gbfcd, _caag := _gbgc, _aeedd
	if _caag < _gbfcd {
		_gbfcd, _caag = _caag, _gbfcd
	}
	_fegeg := make([]float64, 0, len(_babb)+2)
	_fegeg = append(_fegeg, _gbgc)
	for _, _cfaf := range _babb {
		if _cfaf <= _gbfcd {
			continue
		} else if _cfaf >= _caag {
			break
		}
		_fegeg = append(_fegeg, _cfaf)
	}
	_fegeg = append(_fegeg, _aeedd)
	return _fegeg
}
func (_ebfa *textObject) setWordSpacing(_bga float64) {
	if _ebfa == nil {
		return
	}
	_ebfa._cbag._bfa = _bga
}

type stateStack []*textState

func (_cgeb rulingList) merge() *ruling {
	_bedc := _cgeb[0]._abfg
	_geeg := _cgeb[0]._daag
	_bcbcb := _cgeb[0]._fcff
	for _, _cdaca := range _cgeb[1:] {
		_bedc += _cdaca._abfg
		if _cdaca._daag < _geeg {
			_geeg = _cdaca._daag
		}
		if _cdaca._fcff > _bcbcb {
			_bcbcb = _cdaca._fcff
		}
	}
	_gfgce := &ruling{_cffe: _cgeb[0]._cffe, _bbge: _cgeb[0]._bbge, Color: _cgeb[0].Color, _abfg: _bedc / float64(len(_cgeb)), _daag: _geeg, _fcff: _bcbcb}
	if _dgbge {
		_ed.Log.Info("\u006de\u0072g\u0065\u003a\u0020\u0025\u0032d\u0020\u0076e\u0063\u0073\u0020\u0025\u0073", len(_cgeb), _gfgce)
		for _adbd, _eabe := range _cgeb {
			_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0073\u000a", _adbd, _eabe)
		}
	}
	return _gfgce
}
func (_aefa paraList) yNeighbours(_cbde float64) map[*textPara][]int {
	_cbcf := make([]event, 2*len(_aefa))
	if _cbde == 0 {
		for _aaecc, _bfdedc := range _aefa {
			_cbcf[2*_aaecc] = event{_bfdedc.Lly, true, _aaecc}
			_cbcf[2*_aaecc+1] = event{_bfdedc.Ury, false, _aaecc}
		}
	} else {
		for _edgbg, _bbff := range _aefa {
			_cbcf[2*_edgbg] = event{_bbff.Lly - _cbde*_bbff.fontsize(), true, _edgbg}
			_cbcf[2*_edgbg+1] = event{_bbff.Ury + _cbde*_bbff.fontsize(), false, _edgbg}
		}
	}
	return _aefa.eventNeighbours(_cbcf)
}
func (_ebgd *wordBag) removeDuplicates() {
	if _fccg {
		_ed.Log.Info("r\u0065m\u006f\u0076\u0065\u0044\u0075\u0070\u006c\u0069c\u0061\u0074\u0065\u0073: \u0025\u0071", _ebgd.text())
	}
	for _, _faeeg := range _ebgd.depthIndexes() {
		if len(_ebgd._gegc[_faeeg]) == 0 {
			continue
		}
		_egbb := _ebgd._gegc[_faeeg][0]
		_efd := _gcca * _egbb._fgdbd
		_eadfb := _egbb._cbfcee
		for _, _deab := range _ebgd.depthBand(_eadfb, _eadfb+_efd) {
			_ececd := map[*textWord]struct{}{}
			_cgcf := _ebgd._gegc[_deab]
			for _, _bbbae := range _cgcf {
				if _, _ggbdf := _ececd[_bbbae]; _ggbdf {
					continue
				}
				for _, _cabb := range _cgcf {
					if _, _afce := _ececd[_cabb]; _afce {
						continue
					}
					if _cabb != _bbbae && _cabb._bcfea == _bbbae._bcfea && _df.Abs(_cabb.Llx-_bbbae.Llx) < _efd && _df.Abs(_cabb.Urx-_bbbae.Urx) < _efd && _df.Abs(_cabb.Lly-_bbbae.Lly) < _efd && _df.Abs(_cabb.Ury-_bbbae.Ury) < _efd {
						_ececd[_cabb] = struct{}{}
					}
				}
			}
			if len(_ececd) > 0 {
				_aadga := 0
				for _, _ggde := range _cgcf {
					if _, _bfed := _ececd[_ggde]; !_bfed {
						_cgcf[_aadga] = _ggde
						_aadga++
					}
				}
				_ebgd._gegc[_deab] = _cgcf[:len(_cgcf)-len(_ececd)]
				if len(_ebgd._gegc[_deab]) == 0 {
					delete(_ebgd._gegc, _deab)
				}
			}
		}
	}
}
func (_bdcfb paraList) computeEBBoxes() {
	if _adab {
		_ed.Log.Info("\u0063o\u006dp\u0075\u0074\u0065\u0045\u0042\u0042\u006f\u0078\u0065\u0073\u003a")
	}
	for _, _gbef := range _bdcfb {
		_gbef._cedf = _gbef.PdfRectangle
	}
	_aggdb := _bdcfb.yNeighbours(0)
	for _bdcfa, _bdab := range _bdcfb {
		_affa := _bdab._cedf
		_cfbad, _cfad := -1.0e9, +1.0e9
		for _, _beea := range _aggdb[_bdab] {
			_ggbb := _bdcfb[_beea]._cedf
			if _ggbb.Urx < _affa.Llx {
				_cfbad = _df.Max(_cfbad, _ggbb.Urx)
			} else if _affa.Urx < _ggbb.Llx {
				_cfad = _df.Min(_cfad, _ggbb.Llx)
			}
		}
		for _feadf, _cagee := range _bdcfb {
			_cafg := _cagee._cedf
			if _bdcfa == _feadf || _cafg.Ury > _affa.Lly {
				continue
			}
			if _cfbad <= _cafg.Llx && _cafg.Llx < _affa.Llx {
				_affa.Llx = _cafg.Llx
			} else if _cafg.Urx <= _cfad && _affa.Urx < _cafg.Urx {
				_affa.Urx = _cafg.Urx
			}
		}
		if _adab {
			_cfe.Printf("\u0025\u0034\u0064\u003a %\u0036\u002e\u0032\u0066\u2192\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0071\u000a", _bdcfa, _bdab._cedf, _affa, _fabf(_bdab.text(), 50))
		}
		_bdab._cedf = _affa
	}
	if _bcba {
		for _, _bbed := range _bdcfb {
			_bbed.PdfRectangle = _bbed._cedf
		}
	}
}
func _gfeg(_cdac *paraList) map[int][]*textLine {
	_egd := map[int][]*textLine{}
	for _, _fbcg := range *_cdac {
		for _, _gagb := range _fbcg._abbe {
			if !_bbg(_gagb) {
				_ed.Log.Debug("g\u0072\u006f\u0075p\u004c\u0069\u006e\u0065\u0073\u003a\u0020\u0054\u0068\u0065\u0020\u0074\u0065\u0078\u0074\u0020\u006c\u0069\u006e\u0065\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0073 \u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u006e\u0065 \u006d\u0063\u0069\u0064 \u006e\u0075\u006d\u0062e\u0072\u002e\u0020\u0049\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0073p\u006c\u0069\u0074\u002e")
				continue
			}
			_badbe := _gagb._bfca[0]._fdgbc[0]._afae
			_egd[_badbe] = append(_egd[_badbe], _gagb)
		}
		if _fbcg._bedfe != nil {
			_acaa := _fbcg._bedfe._adda
			for _, _bfce := range _acaa {
				for _, _eccf := range _bfce._abbe {
					if !_bbg(_eccf) {
						_ed.Log.Debug("g\u0072\u006f\u0075p\u004c\u0069\u006e\u0065\u0073\u003a\u0020\u0054\u0068\u0065\u0020\u0074\u0065\u0078\u0074\u0020\u006c\u0069\u006e\u0065\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0073 \u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u006e\u0065 \u006d\u0063\u0069\u0064 \u006e\u0075\u006d\u0062e\u0072\u002e\u0020\u0049\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0073p\u006c\u0069\u0074\u002e")
						continue
					}
					_bfded := _eccf._bfca[0]._fdgbc[0]._afae
					_egd[_bfded] = append(_egd[_bfded], _eccf)
				}
			}
		}
	}
	return _egd
}
func _cgfgb(_efabc, _dfec *textPara) bool { return _cdcd(_efabc._cedf, _dfec._cedf) }
func _abdc(_abfb []*textWord, _aaaac *textWord) []*textWord {
	for _dbgf, _deba := range _abfb {
		if _deba == _aaaac {
			return _ffbbc(_abfb, _dbgf)
		}
	}
	_ed.Log.Error("\u0072\u0065\u006d\u006f\u0076e\u0057\u006f\u0072\u0064\u003a\u0020\u0077\u006f\u0072\u0064\u0073\u0020\u0064o\u0065\u0073\u006e\u0027\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0077\u006f\u0072\u0064\u003d\u0025\u0073", _aaaac)
	return nil
}
func (_ddea *textObject) setCharSpacing(_dfa float64) {
	if _ddea == nil {
		return
	}
	_ddea._cbag._bfdf = _dfa
	if _geafa {
		_ed.Log.Info("\u0073\u0065t\u0043\u0068\u0061\u0072\u0053\u0070\u0061\u0063\u0069\u006e\u0067\u003a\u0020\u0025\u002e\u0032\u0066\u0020\u0073\u0074\u0061\u0074e=\u0025\u0073", _dfa, _ddea._cbag.String())
	}
}

// PageTextOptions holds various options available in extraction process.
type PageTextOptions struct {
	_gdeb bool
	_ebff bool
}

func (_fbc *imageExtractContext) extractInlineImage(_bfg *_aad.ContentStreamInlineImage, _dde _aad.GraphicsState, _fdb *_gb.PdfPageResources) error {
	_gca, _fad := _bfg.ToImage(_fdb)
	if _fad != nil {
		return _fad
	}
	_eacg, _fad := _bfg.GetColorSpace(_fdb)
	if _fad != nil {
		return _fad
	}
	if _eacg == nil {
		_eacg = _gb.NewPdfColorspaceDeviceGray()
	}
	_abd, _fad := _eacg.ImageToRGB(*_gca)
	if _fad != nil {
		return _fad
	}
	_ba := ImageMark{Image: &_abd, Width: _dde.CTM.ScalingFactorX(), Height: _dde.CTM.ScalingFactorY(), Angle: _dde.CTM.Angle()}
	_ba.X, _ba.Y = _dde.CTM.Translation()
	_fbc._cdb = append(_fbc._cdb, _ba)
	_fbc._fbd++
	return nil
}

const _dgd = 20

func (_gfegf rulingList) removeDuplicates() rulingList {
	if len(_gfegf) == 0 {
		return nil
	}
	_gfegf.sort()
	_fbad := rulingList{_gfegf[0]}
	for _, _acgd := range _gfegf[1:] {
		if _acgd.equals(_fbad[len(_fbad)-1]) {
			continue
		}
		_fbad = append(_fbad, _acgd)
	}
	return _fbad
}

// String returns a description of `v`.
func (_eead *ruling) String() string {
	if _eead._cffe == _eefaf {
		return "\u004e\u004f\u0054\u0020\u0052\u0055\u004c\u0049\u004e\u0047"
	}
	_dcce, _bcce := "\u0078", "\u0079"
	if _eead._cffe == _bdfd {
		_dcce, _bcce = "\u0079", "\u0078"
	}
	_ababb := ""
	if _eead._fdde != 0.0 {
		_ababb = _cfe.Sprintf(" \u0077\u0069\u0064\u0074\u0068\u003d\u0025\u002e\u0032\u0066", _eead._fdde)
	}
	return _cfe.Sprintf("\u0025\u00310\u0073\u0020\u0025\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u0020\u0025\u0073\u003d\u0025\u0036\u002e\u0032\u0066\u0020\u002d\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0028\u0025\u0036\u002e\u0032\u0066\u0029\u0020\u0025\u0073\u0020\u0025\u0076\u0025\u0073", _eead._cffe, _dcce, _eead._abfg, _bcce, _eead._daag, _eead._fcff, _eead._fcff-_eead._daag, _eead._bbge, _eead.Color, _ababb)
}
func (_ddfff *textTable) subdivide() *textTable {
	_ddfff.logComposite("\u0073u\u0062\u0064\u0069\u0076\u0069\u0064e")
	_ebaec := _ddfff.compositeRowCorridors()
	_dfab := _ddfff.compositeColCorridors()
	if _geff {
		_ed.Log.Info("\u0073u\u0062\u0064i\u0076\u0069\u0064\u0065:\u000a\u0009\u0072o\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s=\u0025\u0073\u000a\t\u0063\u006fl\u0043\u006f\u0072\u0072\u0069\u0064o\u0072\u0073=\u0025\u0073", _fgag(_ebaec), _fgag(_dfab))
	}
	if len(_ebaec) == 0 || len(_dfab) == 0 {
		return _ddfff
	}
	_bcgfca(_ebaec)
	_bcgfca(_dfab)
	if _geff {
		_ed.Log.Info("\u0073\u0075\u0062\u0064\u0069\u0076\u0069\u0064\u0065\u0020\u0066\u0069\u0078\u0065\u0064\u003a\u000a\u0009r\u006f\u0077\u0043\u006f\u0072\u0072\u0069d\u006f\u0072\u0073\u003d\u0025\u0073\u000a\u0009\u0063\u006f\u006cC\u006f\u0072\u0072\u0069\u0064\u006f\u0072\u0073\u003d\u0025\u0073", _fgag(_ebaec), _fgag(_dfab))
	}
	_cccf, _fdcae := _dcfc(_ddfff._dege, _ebaec)
	_adafb, _aceg := _dcfc(_ddfff._fccgee, _dfab)
	_edefa := make(map[uint64]*textPara, _aceg*_fdcae)
	_dced := &textTable{PdfRectangle: _ddfff.PdfRectangle, _baccf: _ddfff._baccf, _dege: _fdcae, _fccgee: _aceg, _adda: _edefa}
	if _geff {
		_ed.Log.Info("\u0073\u0075b\u0064\u0069\u0076\u0069\u0064\u0065\u003a\u0020\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u0020\u0063\u0065\u006c\u006c\u0073\u003d\u0020\u0025\u0064\u0020\u0078\u0020\u0025\u0064\u000a"+"\u0009\u0072\u006f\u0077\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s\u003d\u0025\u0073\u000a"+"\u0009\u0063\u006f\u006c\u0043\u006f\u0072\u0072\u0069\u0064\u006f\u0072s\u003d\u0025\u0073\u000a"+"\u0009\u0079\u004f\u0066\u0066\u0073\u0065\u0074\u0073=\u0025\u002b\u0076\u000a"+"\u0009\u0078\u004f\u0066\u0066\u0073\u0065\u0074\u0073\u003d\u0025\u002b\u0076", _ddfff._fccgee, _ddfff._dege, _aceg, _fdcae, _fgag(_ebaec), _fgag(_dfab), _cccf, _adafb)
	}
	for _cfbb := 0; _cfbb < _ddfff._dege; _cfbb++ {
		_aedaa := _cccf[_cfbb]
		for _fafe := 0; _fafe < _ddfff._fccgee; _fafe++ {
			_ecfca := _adafb[_fafe]
			if _geff {
				_cfe.Printf("\u0025\u0036\u0064\u002c %\u0032\u0064\u003a\u0020\u0078\u0030\u003d\u0025\u0064\u0020\u0079\u0030\u003d\u0025d\u000a", _fafe, _cfbb, _ecfca, _aedaa)
			}
			_dggc, _eecf := _ddfff._faeg[_aafb(_fafe, _cfbb)]
			if !_eecf {
				continue
			}
			_gfbef := _dggc.split(_ebaec[_cfbb], _dfab[_fafe])
			for _dcbe := 0; _dcbe < _gfbef._dege; _dcbe++ {
				for _bfdca := 0; _bfdca < _gfbef._fccgee; _bfdca++ {
					_abcb := _gfbef.get(_bfdca, _dcbe)
					_dced.put(_ecfca+_bfdca, _aedaa+_dcbe, _abcb)
					if _geff {
						_cfe.Printf("\u0025\u0038\u0064\u002c\u0020\u0025\u0032\u0064\u003a\u0020\u0025\u0073\u000a", _ecfca+_bfdca, _aedaa+_dcbe, _abcb)
					}
				}
			}
		}
	}
	return _dced
}
func (_dea *textObject) setTextMatrix(_dag []float64) {
	if len(_dag) != 6 {
		_ed.Log.Debug("\u0045\u0052\u0052OR\u003a\u0020\u006c\u0065\u006e\u0028\u0066\u0029\u0020\u0021\u003d\u0020\u0036\u0020\u0028\u0025\u0064\u0029", len(_dag))
		return
	}
	_ggc, _edfd, _ace, _gbbc, _aecf, _gac := _dag[0], _dag[1], _dag[2], _dag[3], _dag[4], _dag[5]
	_dea._ffb = _cc.NewMatrix(_ggc, _edfd, _ace, _gbbc, _aecf, _gac)
	_dea._ddge = _dea._ffb
}
func (_cdgg *textLine) endsInHyphen() bool {
	_ggad := _cdgg._bfca[len(_cdgg._bfca)-1]
	_cbdb := _ggad._bcfea
	_eacb, _efff := _cb.DecodeLastRuneInString(_cbdb)
	if _efff <= 0 || !_ac.Is(_ac.Hyphen, _eacb) {
		return false
	}
	if _ggad._cacca && _dcebc(_cbdb) {
		return true
	}
	return _dcebc(_cdgg.text())
}

// Font represents the font properties on a PDF page.
type Font struct {
	PdfFont *_gb.PdfFont

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
	FontDescriptor *_gb.PdfFontDescriptor
}

func (_bcae *wordBag) highestWord(_bdgd int, _acge, _bed float64) *textWord {
	for _, _cbga := range _bcae._gegc[_bdgd] {
		if _acge <= _cbga._cbfcee && _cbga._cbfcee <= _bed {
			return _cbga
		}
	}
	return nil
}
func _bcgfca(_ffba map[int][]float64) {
	if len(_ffba) <= 1 {
		return
	}
	_ecfab := _dgce(_ffba)
	if _geff {
		_ed.Log.Info("\u0066i\u0078C\u0065\u006c\u006c\u0073\u003a \u006b\u0065y\u0073\u003d\u0025\u002b\u0076", _ecfab)
	}
	var _dacdf, _egdb int
	for _dacdf, _egdb = range _ecfab {
		if _ffba[_egdb] != nil {
			break
		}
	}
	for _gegbe, _becdd := range _ecfab[_dacdf:] {
		_ggcg := _ffba[_becdd]
		if _ggcg == nil {
			continue
		}
		if _geff {
			_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u006b\u0030\u003d\u0025\u0064\u0020\u006b1\u003d\u0025\u0064\u000a", _dacdf+_gegbe, _egdb, _becdd)
		}
		_cbbab := _ffba[_becdd]
		if _cbbab[len(_cbbab)-1] > _ggcg[0] {
			_cbbab[len(_cbbab)-1] = _ggcg[0]
			_ffba[_egdb] = _cbbab
		}
		_egdb = _becdd
	}
}
func _eaag(_bdcf []*textLine) []*textLine {
	_fdef := []*textLine{}
	for _, _bcccc := range _bdcf {
		_ccbe := _bcccc.text()
		_dadg := _aaeg.Find([]byte(_ccbe))
		if _dadg != nil {
			_fdef = append(_fdef, _bcccc)
		}
	}
	return _fdef
}
func (_egeab rulingList) tidied(_eeedd string) rulingList {
	_daga := _egeab.removeDuplicates()
	_daga.log("\u0075n\u0069\u0071\u0075\u0065\u0073")
	_adfc := _daga.snapToGroups()
	if _adfc == nil {
		return nil
	}
	_adfc.sort()
	if _geaeg {
		_ed.Log.Info("\u0074\u0069\u0064i\u0065\u0064\u003a\u0020\u0025\u0071\u0020\u0076\u0065\u0063\u0073\u003d\u0025\u0064\u0020\u0075\u006e\u0069\u0071\u0075\u0065\u0073\u003d\u0025\u0064\u0020\u0063\u006f\u0061l\u0065\u0073\u0063\u0065\u0064\u003d\u0025\u0064", _eeedd, len(_egeab), len(_daga), len(_adfc))
	}
	_adfc.log("\u0063o\u0061\u006c\u0065\u0073\u0063\u0065d")
	return _adfc
}
func (_ggdb *textLine) appendWord(_feea *textWord) {
	_ggdb._bfca = append(_ggdb._bfca, _feea)
	_ggdb.PdfRectangle = _fde(_ggdb.PdfRectangle, _feea.PdfRectangle)
	if _feea._fgdbd > _ggdb._egbd {
		_ggdb._egbd = _feea._fgdbd
	}
	if _feea._cbfcee > _ggdb._bdbg {
		_ggdb._bdbg = _feea._cbfcee
	}
}
func _bbcee(_eddac []*textMark, _ddcef _gb.PdfRectangle) *textWord {
	_affee := _eddac[0].PdfRectangle
	_dfbbbf := _eddac[0]._gded
	for _, _egda := range _eddac[1:] {
		_affee = _fde(_affee, _egda.PdfRectangle)
		if _egda._gded > _dfbbbf {
			_dfbbbf = _egda._gded
		}
	}
	return &textWord{PdfRectangle: _affee, _fdgbc: _eddac, _cbfcee: _ddcef.Ury - _affee.Lly, _fgdbd: _dfbbbf}
}
func (_dbaa paraList) findTextTables() []*textTable {
	var _cdcge []*textTable
	for _, _agebe := range _dbaa {
		if _agebe.taken() || _agebe.Width() == 0 {
			continue
		}
		_ggdc := _agebe.isAtom()
		if _ggdc == nil {
			continue
		}
		_ggdc.growTable()
		if _ggdc._fccgee*_ggdc._dege < _dgea {
			continue
		}
		_ggdc.markCells()
		_ggdc.log("\u0067\u0072\u006fw\u006e")
		_cdcge = append(_cdcge, _ggdc)
	}
	return _cdcge
}
func _dbbg(_ebaef *wordBag, _beeg *textWord, _bdcd float64) bool {
	return _ebaef.Urx <= _beeg.Llx && _beeg.Llx < _ebaef.Urx+_bdcd
}
func (_fgcf *shapesState) quadraticTo(_ffaa, _dfbf, _eeab, _fdfg float64) {
	if _dgeg {
		_ed.Log.Info("\u0071\u0075\u0061d\u0072\u0061\u0074\u0069\u0063\u0054\u006f\u003a")
	}
	_fgcf.addPoint(_eeab, _fdfg)
}
func (_eaab *compositeCell) updateBBox() {
	for _, _fagg := range _eaab.paraList {
		_eaab.PdfRectangle = _fde(_eaab.PdfRectangle, _fagg.PdfRectangle)
	}
}
func (_ecef paraList) reorder(_aagcd []int) {
	_caccb := make(paraList, len(_ecef))
	for _acbb, _agfe := range _aagcd {
		_caccb[_acbb] = _ecef[_agfe]
	}
	copy(_ecef, _caccb)
}
func _daefa(_gcbf, _cegc bounded) float64 { return _gcbf.bbox().Llx - _cegc.bbox().Urx }

// NewFromContents creates a new extractor from contents and page resources.
func NewFromContents(contents string, resources *_gb.PdfPageResources) (*Extractor, error) {
	_ece := &Extractor{_ee: contents, _bd: resources, _ebb: map[string]fontEntry{}, _db: map[string]textResult{}}
	return _ece, nil
}
func (_gdcfg *subpath) clear()                  { *_gdcfg = subpath{} }
func (_ddbga *textPara) bbox() _gb.PdfRectangle { return _ddbga.PdfRectangle }
func (_agfc *wordBag) pullWord(_fbe *textWord, _acee int, _gacfe map[int]map[*textWord]struct{}) {
	_agfc.PdfRectangle = _fde(_agfc.PdfRectangle, _fbe.PdfRectangle)
	if _fbe._fgdbd > _agfc._bcef {
		_agfc._bcef = _fbe._fgdbd
	}
	_agfc._gegc[_acee] = append(_agfc._gegc[_acee], _fbe)
	_gacfe[_acee][_fbe] = struct{}{}
}
func (_dfccd rulingList) isActualGrid() (rulingList, bool) {
	_deddg, _affec := _dfccd.augmentGrid()
	if !(len(_deddg) >= _edeec+1 && len(_affec) >= _fegbc+1) {
		if _geaeg {
			_ed.Log.Info("\u0069s\u0041\u0063t\u0075\u0061\u006c\u0047r\u0069\u0064\u003a \u004e\u006f\u0074\u0020\u0061\u006c\u0069\u0067\u006eed\u002e\u0020\u0025d\u0020\u0078 \u0025\u0064\u0020\u003c\u0020\u0025d\u0020\u0078 \u0025\u0064", len(_deddg), len(_affec), _edeec+1, _fegbc+1)
		}
		return nil, false
	}
	if _geaeg {
		_ed.Log.Info("\u0069\u0073\u0041\u0063\u0074\u0075a\u006c\u0047\u0072\u0069\u0064\u003a\u0020\u0025\u0073\u0020\u003a\u0020\u0025t\u0020\u0026\u0020\u0025\u0074\u0020\u2192 \u0025\u0074", _dfccd, len(_deddg) >= 2, len(_affec) >= 2, len(_deddg) >= 2 && len(_affec) >= 2)
		for _fcaea, _bfgg := range _dfccd {
			_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0076\u000a", _fcaea, _bfgg)
		}
	}
	if _abdfd {
		_efdg, _fced := _deddg[0], _deddg[len(_deddg)-1]
		_bgdeb, _fddc := _affec[0], _affec[len(_affec)-1]
		if !(_adcb(_efdg._abfg-_bgdeb._daag) && _adcb(_fced._abfg-_bgdeb._fcff) && _adcb(_bgdeb._abfg-_efdg._fcff) && _adcb(_fddc._abfg-_efdg._daag)) {
			if _geaeg {
				_ed.Log.Info("\u0069\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069d\u003a\u0020\u0020N\u006f\u0074 \u0061\u006c\u0069\u0067\u006e\u0065d\u002e\n\t\u0076\u0030\u003d\u0025\u0073\u000a\u0009\u0076\u0031\u003d\u0025\u0073\u000a\u0009\u0068\u0030\u003d\u0025\u0073\u000a\u0009\u0068\u0031\u003d\u0025\u0073", _efdg, _fced, _bgdeb, _fddc)
			}
			return nil, false
		}
	} else {
		if !_deddg.aligned() {
			if _dgbge {
				_ed.Log.Info("i\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069\u0064\u003a\u0020N\u006f\u0074\u0020\u0061\u006c\u0069\u0067n\u0065\u0064\u0020\u0076\u0065\u0072\u0074\u0073\u002e\u0020%\u0064", len(_deddg))
			}
			return nil, false
		}
		if !_affec.aligned() {
			if _geaeg {
				_ed.Log.Info("i\u0073\u0041\u0063\u0074\u0075\u0061l\u0047\u0072\u0069\u0064\u003a\u0020N\u006f\u0074\u0020\u0061\u006c\u0069\u0067n\u0065\u0064\u0020\u0068\u006f\u0072\u007a\u0073\u002e\u0020%\u0064", len(_affec))
			}
			return nil, false
		}
	}
	_dccdb := append(_deddg, _affec...)
	return _dccdb, true
}
func (_efcb *stateStack) empty() bool { return len(*_efcb) == 0 }
func _cccea(_dbfec, _bdbe _cc.Point) rulingKind {
	_gebdd := _df.Abs(_dbfec.X - _bdbe.X)
	_fbfce := _df.Abs(_dbfec.Y - _bdbe.Y)
	return _gaba(_gebdd, _fbfce, _caf)
}

type structElement struct {
	_gacb  string
	_ddcfe []structElement
	_adfb  int64
	_gfeff _gg.PdfObject
}

func (_gafdb paraList) log(_abfde string) {
	if !_baff {
		return
	}
	_ed.Log.Info("%\u0038\u0073\u003a\u0020\u0025\u0064 \u0070\u0061\u0072\u0061\u0073\u0020=\u003d\u003d\u003d\u003d\u003d\u003d\u002d-\u002d\u002d\u002d\u002d\u002d\u003d\u003d\u003d\u003d\u003d=\u003d", _abfde, len(_gafdb))
	for _ggdba, _cgbe := range _gafdb {
		if _cgbe == nil {
			continue
		}
		_cgef := _cgbe.text()
		_edfe := "\u0020\u0020"
		if _cgbe._bedfe != nil {
			_edfe = _cfe.Sprintf("\u005b%\u0064\u0078\u0025\u0064\u005d", _cgbe._bedfe._fccgee, _cgbe._bedfe._dege)
		}
		_cfe.Printf("\u0025\u0034\u0064\u003a\u0020\u0025\u0036\u002e\u0032\u0066\u0020\u0025s\u0020\u0025\u0071\u000a", _ggdba, _cgbe.PdfRectangle, _edfe, _fabf(_cgef, 50))
	}
}
func (_ggge *structTreeRoot) parseStructTreeRoot(_deee _gg.PdfObject) {
	if _deee != nil {
		_abfcf, _cced := _gg.GetDict(_deee)
		if !_cced {
			_ed.Log.Debug("\u0070\u0061\u0072s\u0065\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u003a\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006eo\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e")
		}
		K := _abfcf.Get("\u004b")
		_fagf := _abfcf.Get("\u0054\u0079\u0070\u0065").String()
		var _caa *_gg.PdfObjectArray
		switch _bade := K.(type) {
		case *_gg.PdfObjectArray:
			_caa = _bade
		case *_gg.PdfObjectReference:
			_caa = _gg.MakeArray(K)
		}
		_cfec := []structElement{}
		for _, _dddef := range _caa.Elements() {
			_baedg := &structElement{}
			_baedg.parseStructElement(_dddef)
			_cfec = append(_cfec, *_baedg)
		}
		_ggge._gafdd = _cfec
		_ggge._bgee = _fagf
	}
}
func _gdag(_defe *Extractor, _gaa *_gb.PdfPageResources, _gacc _aad.GraphicsState, _geea *textState, _egfc *stateStack) *textObject {
	return &textObject{_cgf: _defe, _edacg: _gaa, _ccg: _gacc, _bab: _egfc, _cbag: _geea, _ffb: _cc.IdentityMatrix(), _ddge: _cc.IdentityMatrix()}
}

type subpath struct {
	_ccgf []_cc.Point
	_beac bool
}
