package pdfa

import (
	_d "errors"
	_a "fmt"
	_c "image/color"
	_e "math"
	_af "sort"
	_ac "strings"
	_eg "time"

	_ec "bitbucket.org/shenghui0779/gopdf/common"
	_ff "bitbucket.org/shenghui0779/gopdf/contentstream"
	_cgd "bitbucket.org/shenghui0779/gopdf/core"
	_bd "bitbucket.org/shenghui0779/gopdf/internal/cmap"
	_cf "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_ca "bitbucket.org/shenghui0779/gopdf/internal/timeutils"
	_f "bitbucket.org/shenghui0779/gopdf/model"
	_dfa "bitbucket.org/shenghui0779/gopdf/model/internal/colorprofile"
	_gg "bitbucket.org/shenghui0779/gopdf/model/internal/docutil"
	_ecb "bitbucket.org/shenghui0779/gopdf/model/internal/fonts"
	_cg "bitbucket.org/shenghui0779/gopdf/model/xmputil"
	_df "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaextension"
	_da "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaid"
	_gc "github.com/adrg/sysfont"
	_be "github.com/trimmer-io/go-xmp/models/dc"
	_g "github.com/trimmer-io/go-xmp/models/pdf"
	_ee "github.com/trimmer-io/go-xmp/models/xmp_base"
	_bf "github.com/trimmer-io/go-xmp/models/xmp_mm"
	_ae "github.com/trimmer-io/go-xmp/models/xmp_rights"
	_aca "github.com/trimmer-io/go-xmp/xmp"
)

func _fgecc(_ffbf *_f.CompliancePdfReader) (_dage []ViolatedRule) {
	var _aadg, _ggae, _ebd, _adgf, _eede, _acde bool
	_dbcc := map[*_cgd.PdfObjectStream]struct{}{}
	for _, _gaag := range _ffbf.GetObjectNums() {
		if _aadg && _ggae && _eede && _ebd && _adgf && _acde {
			return _dage
		}
		_fbfd, _caa := _ffbf.GetIndirectObjectByNumber(_gaag)
		if _caa != nil {
			continue
		}
		_beade, _gdbef := _cgd.GetStream(_fbfd)
		if !_gdbef {
			continue
		}
		if _, _gdbef = _dbcc[_beade]; _gdbef {
			continue
		}
		_dbcc[_beade] = struct{}{}
		_afde, _gdbef := _cgd.GetName(_beade.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_gdbef {
			continue
		}
		if !_adgf {
			if _beade.Get("\u0052\u0065\u0066") != nil {
				_dage = append(_dage, _cc("\u0036.\u0032\u002e\u0036\u002d\u0031", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e"))
				_adgf = true
			}
		}
		if _afde.String() == "\u0050\u0053" {
			if !_acde {
				_dage = append(_dage, _cc("\u0036.\u0032\u002e\u0037\u002d\u0031", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073."))
				_acde = true
				continue
			}
		}
		if _afde.String() == "\u0046\u006f\u0072\u006d" {
			if _ggae && _ebd && _adgf {
				continue
			}
			if !_ggae && _beade.Get("\u004f\u0050\u0049") != nil {
				_dage = append(_dage, _cc("\u0036.\u0032\u002e\u0034\u002d\u0032", "\u0041\u006e\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u0028\u0049\u006d\u0061\u0067\u0065\u0020\u006f\u0072\u0020\u0046\u006f\u0072\u006d\u0029\u0020\u0073\u0068\u0061\u006cl\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u004fP\u0049\u0020\u006b\u0065\u0079\u002e"))
				_ggae = true
			}
			if !_ebd {
				if _beade.Get("\u0050\u0053") != nil {
					_ebd = true
				}
				if _fddf := _beade.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _fddf != nil && !_ebd {
					if _dfgb, _cfgg := _cgd.GetName(_fddf); _cfgg && *_dfgb == "\u0050\u0053" {
						_ebd = true
					}
				}
				if _ebd {
					_dage = append(_dage, _cc("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e"))
				}
			}
			continue
		}
		if _afde.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if !_aadg && _beade.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
			_dage = append(_dage, _cc("\u0036.\u0032\u002e\u0034\u002d\u0031", "\u0041\u006e\u0020\u0049m\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
			_aadg = true
		}
		if !_eede && _beade.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065") != nil {
			_dbg, _gaec := _cgd.GetBool(_beade.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _gaec && bool(*_dbg) {
				continue
			}
			_dage = append(_dage, _cc("\u0036.\u0032\u002e\u0034\u002d\u0033", "\u0049\u0066 a\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0063o\u006e\u0074\u0061\u0069n\u0073\u0020\u0074\u0068e \u0049\u006et\u0065r\u0070\u006f\u006c\u0061\u0074\u0065 \u006b\u0065\u0079,\u0020\u0069t\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020b\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
			_eede = true
		}
	}
	return _dage
}
func (_ef *documentImages) hasOnlyDeviceCMYK() bool { return _ef._bg && !_ef._fde && !_ef._bgd }

type pageColorspaceOptimizeFunc func(_aaeb *_gg.Document, _edgf *_gg.Page, _gddf []*_gg.Image) error

func _gf() standardType { return standardType{_ba: 2, _fa: "\u0055"} }
func _caag(_acbcf *_f.CompliancePdfReader) (_gcgb []ViolatedRule) {
	_daac := true
	_efag, _eddf := _acbcf.GetCatalogMarkInfo()
	if !_eddf {
		_daac = false
	} else {
		_gceeb, _eeeba := _cgd.GetDict(_efag)
		if _eeeba {
			_gdeb, _adeg := _cgd.GetBool(_gceeb.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
			if !bool(*_gdeb) || !_adeg {
				_daac = false
			}
		} else {
			_daac = false
		}
	}
	if !_daac {
		_gcgb = append(_gcgb, _cc("\u0036.\u0038\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006cog\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020M\u0061r\u006b\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061 \u004d\u0061\u0072\u006b\u0065\u0064\u0020\u0065\u006et\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074,\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0076\u0061lu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0072\u0075\u0065"))
	}
	_ffffd, _eddf := _acbcf.GetCatalogStructTreeRoot()
	if !_eddf {
		_gcgb = append(_gcgb, _cc("\u0036.\u0038\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075r\u0065\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065d \u0062\u0079\u0020a\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072e\u0020\u0068\u0069\u0065\u0072\u0061\u0072\u0063\u0068\u0079\u0020\u0072\u006f\u006ft\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0065\u006e\u0074r\u0079\u0020\u006f\u0066\u0020\u0074h\u0065\u0020d\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061t\u0061\u006c\u006fg \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0036\u002e"))
	}
	_dedc, _eddf := _cgd.GetDict(_ffffd)
	if _eddf {
		_egbd, _efdf := _cgd.GetName(_dedc.Get("\u0052o\u006c\u0065\u004d\u0061\u0070"))
		if _efdf {
			_cdcg, _agag := _cgd.GetDict(_egbd)
			if _agag {
				for _, _bdddd := range _cdcg.Keys() {
					_eceb := _cdcg.Get(_bdddd)
					if _eceb == nil {
						_gcgb = append(_gcgb, _cc("\u0036.\u0038\u002e\u0033\u002d\u0032", "\u0041\u006c\u006c\u0020\u006eo\u006e\u002ds\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0073t\u0072\u0075\u0063\u0074ure\u0020\u0074\u0079\u0070\u0065s\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u006d\u0061\u0070\u0070\u0065d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020n\u0065\u0061\u0072\u0065\u0073\u0074\u0020\u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0073\u0074a\u006ed\u0061r\u0064\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065re\u006e\u0063e\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020i\u006e\u0020\u0074\u0068e\u0020\u0072\u006fl\u0065\u0020\u006d\u0061p \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066 \u0074h\u0065\u0020\u0073\u0074\u0072\u0075c\u0074\u0075r\u0065\u0020\u0074\u0072e\u0065\u0020\u0072\u006f\u006ft\u002e"))
					}
				}
			}
		}
	}
	return _gcgb
}
func (_baac *documentImages) hasOnlyDeviceGray() bool { return _baac._bgd && !_baac._fde && !_baac._bg }
func _dac(_dffg *_gg.Document, _aga int) error {
	_dga := map[*_cgd.PdfObjectStream]struct{}{}
	for _, _cgca := range _dffg.Objects {
		_bcd, _feef := _cgd.GetStream(_cgca)
		if !_feef {
			continue
		}
		if _, _feef = _dga[_bcd]; _feef {
			continue
		}
		_dga[_bcd] = struct{}{}
		_cdd, _feef := _cgd.GetName(_bcd.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_feef {
			continue
		}
		if _bcd.Get("\u0052\u0065\u0066") != nil {
			_bcd.Remove("\u0052\u0065\u0066")
		}
		if _cdd.String() == "\u0050\u0053" {
			_bcd.Remove("\u0050\u0053")
			continue
		}
		if _cdd.String() == "\u0046\u006f\u0072\u006d" {
			if _bcd.Get("\u004f\u0050\u0049") != nil {
				_bcd.Remove("\u004f\u0050\u0049")
			}
			if _bcd.Get("\u0050\u0053") != nil {
				_bcd.Remove("\u0050\u0053")
			}
			if _cbba := _bcd.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _cbba != nil {
				if _gcaab, _geebe := _cgd.GetName(_cbba); _geebe && *_gcaab == "\u0050\u0053" {
					_bcd.Remove("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032")
				}
			}
			continue
		}
		if _cdd.String() == "\u0049\u006d\u0061g\u0065" {
			_bfe, _gbac := _cgd.GetBool(_bcd.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _gbac && bool(*_bfe) {
				_bcd.Set("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065", _cgd.MakeBool(false))
			}
			if _aga == 2 {
				if _bcd.Get("\u004f\u0050\u0049") != nil {
					_bcd.Remove("\u004f\u0050\u0049")
				}
			}
			if _bcd.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
				_bcd.Remove("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073")
			}
			continue
		}
	}
	return nil
}
func _bdge(_gbfee *_f.CompliancePdfReader) (_fdbg []ViolatedRule) {
	var _ccfge, _gadg, _gbga, _bagc, _dcfde, _cfcfe, _fbdcd bool
	_fgbg := func() bool { return _ccfge && _gadg && _gbga && _bagc && _dcfde && _cfcfe && _fbdcd }
	for _, _bdgcc := range _gbfee.PageList {
		_befd, _dgggc := _bdgcc.GetAnnotations()
		if _dgggc != nil {
			_ec.Log.Trace("\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006es\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _dgggc)
			continue
		}
		for _, _cccag := range _befd {
			if !_ccfge {
				switch _cccag.GetContext().(type) {
				case *_f.PdfAnnotationScreen, *_f.PdfAnnotation3D, *_f.PdfAnnotationSound, *_f.PdfAnnotationMovie, nil:
					_fdbg = append(_fdbg, _cc("\u0036.\u0033\u002e\u0031\u002d\u0031", "\u0041nn\u006f\u0074\u0061\u0074i\u006f\u006e t\u0079\u0070\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065r\u006d\u0069t\u0074\u0065\u0064\u002e\u0020\u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0033\u0044\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u0053\u0063\u0072\u0065\u0065\u006e\u0020\u0061n\u0064\u0020\u004d\u006f\u0076\u0069\u0065\u0020\u0074\u0079\u0070\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_ccfge = true
					if _fgbg() {
						return _fdbg
					}
				}
			}
			_cbga, _eaad := _cgd.GetDict(_cccag.GetContainingPdfObject())
			if !_eaad {
				continue
			}
			_, _addag := _cccag.GetContext().(*_f.PdfAnnotationPopup)
			if !_addag && !_gadg {
				_, _cbgfe := _cgd.GetIntVal(_cbga.Get("\u0046"))
				if !_cbgfe {
					_fdbg = append(_fdbg, _cc("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0045\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072i\u0065\u0073\u0020\u0077\u0068\u006fs\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0069\u0073\u0020\u0050\u006f\u0070u\u0070\u002c\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0046 \u006b\u0065y."))
					_gadg = true
					if _fgbg() {
						return _fdbg
					}
				}
			}
			if !_gbga {
				_fbed, _aaeea := _cgd.GetIntVal(_cbga.Get("\u0046"))
				if _aaeea && !(_fbed&4 == 4 && _fbed&1 == 0 && _fbed&2 == 0 && _fbed&32 == 0 && _fbed&256 == 0) {
					_fdbg = append(_fdbg, _cc("\u0036.\u0033\u002e\u0032\u002d\u0032", "I\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u0020\u0046 \u006b\u0065\u0079\u0027\u0073\u0020\u0050\u0072\u0069\u006e\u0074\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0020\u0073\u0068\u0061l\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0074\u0020\u0074\u006f\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0048\u0069\u0064\u0064\u0065\u006e\u002c\u0020\u0049\u006e\u0076\u0069\u0073\u0069\u0062\u006c\u0065\u002c\u0020\u0054\u006f\u0067\u0067\u006c\u0065\u004e\u006f\u0056\u0069\u0065\u0077\u002c\u0020\u0061\u006e\u0064 \u004eo\u0056\u0069\u0065\u0077\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020s\u0065\u0074\u0020t\u006f\u0020\u0030."))
					_gbga = true
					if _fgbg() {
						return _fdbg
					}
				}
			}
			_, _edaaa := _cccag.GetContext().(*_f.PdfAnnotationText)
			if _edaaa && !_bagc {
				_bbdbd, _babdg := _cgd.GetIntVal(_cbga.Get("\u0046"))
				if _babdg && !(_bbdbd&8 == 8 && _bbdbd&16 == 16) {
					_fdbg = append(_fdbg, _cc("\u0036.\u0033\u002e\u0032\u002d\u0033", "\u0054\u0065\u0078\u0074\u0020a\u006e\u006e\u006f\u0074\u0061t\u0069o\u006e\u0020\u0068\u0061\u0073\u0020\u006f\u006e\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006ca\u0067\u0073\u0020\u004e\u006f\u005a\u006f\u006f\u006d\u0020\u006f\u0072\u0020\u004e\u006f\u0052\u006f\u0074\u0061\u0074\u0065\u0020\u0073\u0065t\u0020\u0074\u006f\u0020\u0030\u002e"))
					_bagc = true
					if _fgbg() {
						return _fdbg
					}
				}
			}
			if !_dcfde {
				_abgg, _gacg := _cgd.GetDict(_cbga.Get("\u0041\u0050"))
				if _gacg {
					_ebcf := _abgg.Get("\u004e")
					if _ebcf == nil || len(_abgg.Keys()) > 1 {
						_fdbg = append(_fdbg, _cc("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_dcfde = true
						if _fgbg() {
							return _fdbg
						}
						continue
					}
					_, _bcea := _cccag.GetContext().(*_f.PdfAnnotationWidget)
					if _bcea {
						_bgag, _beefg := _cgd.GetName(_cbga.Get("\u0046\u0054"))
						if _beefg && *_bgag == "\u0042\u0074\u006e" {
							if _, _fbddb := _cgd.GetDict(_ebcf); !_fbddb {
								_fdbg = append(_fdbg, _cc("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
								_dcfde = true
								if _fgbg() {
									return _fdbg
								}
								continue
							}
						}
					}
					_, _dcge := _cgd.GetStream(_ebcf)
					if !_dcge {
						_fdbg = append(_fdbg, _cc("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_dcfde = true
						if _fgbg() {
							return _fdbg
						}
						continue
					}
				}
			}
			_fedac, _acccg := _cccag.GetContext().(*_f.PdfAnnotationWidget)
			if !_acccg {
				continue
			}
			if !_cfcfe {
				if _fedac.A != nil {
					_fdbg = append(_fdbg, _cc("\u0036.\u0034\u002e\u0031\u002d\u0031", "A \u0057\u0069d\u0067\u0065\u0074\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006ec\u006cu\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0020e\u006et\u0072\u0079."))
					_cfcfe = true
					if _fgbg() {
						return _fdbg
					}
				}
			}
			if !_fbdcd {
				if _fedac.AA != nil {
					_fdbg = append(_fdbg, _cc("\u0036.\u0034\u002e\u0031\u002d\u0031", "\u0041\u0020\u0057\u0069\u0064\u0067\u0065\u0074\u0020\u0061\u006e\u006eo\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073h\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					_fbdcd = true
					if _fgbg() {
						return _fdbg
					}
				}
			}
		}
	}
	return _fdbg
}
func _gfcfc(_afcb *_f.CompliancePdfReader) (_aaff []ViolatedRule) {
	var _eagg, _dbbfe, _daea, _eace bool
	_gefg := func() bool { return _eagg && _dbbfe && _daea && _eace }
	_fcag, _acfdg := _gdgaf(_afcb)
	var _cfddd _dfa.ProfileHeader
	if _acfdg {
		_cfddd, _ = _dfa.ParseHeader(_fcag.DestOutputProfile)
	}
	_gfgg := map[_cgd.PdfObject]struct{}{}
	var _bdeg func(_bbfd _f.PdfColorspace) bool
	_bdeg = func(_cdeba _f.PdfColorspace) bool {
		switch _dbcb := _cdeba.(type) {
		case *_f.PdfColorspaceDeviceGray:
			if !_eagg {
				if !_acfdg {
					_aaff = append(_aaff, _cc("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d4", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006f\u006e\u006c\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064 \u0069\u0066\u0020\u0061\u0020\u0064\u0065v\u0069\u0063\u0065\u0020\u0069\u006e\u0064\u0065p\u0065\u006e\u0064\u0065\u006e\u0074\u0020\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0047\u0072\u0061\u0079\u0020\u0063\u006f\u006c\u006f\u0075r \u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061\u0073\u0020\u0062\u0065\u0065\u006e \u0073\u0065\u0074\u0020\u0077\u0068\u0065n \u0074\u0068\u0065\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072a\u0079\u0020\u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u002c o\u0072\u0020\u0069\u0066\u0020\u0061\u0020\u0050\u0044\u0046\u002fA\u0020\u004f\u0075tp\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u002e"))
					_eagg = true
					if _gefg() {
						return true
					}
				}
			}
		case *_f.PdfColorspaceDeviceRGB:
			if !_dbbfe {
				if !_acfdg || _cfddd.ColorSpace != _dfa.ColorSpaceRGB {
					_aaff = append(_aaff, _cc("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d2", "\u0044\u0065\u0076\u0069c\u0065\u0052\u0047\u0042\u0020\u0073\u0068\u0061\u006cl\u0020\u006f\u006e\u006c\u0079\u0020\u0062e\u0020\u0075\u0073\u0065\u0064\u0020\u0069f\u0020\u0061\u0020\u0064\u0065\u0076\u0069\u0063e\u0020\u0069n\u0064\u0065\u0070e\u006e\u0064\u0065\u006et \u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0052\u0047\u0042\u0020\u0063\u006fl\u006f\u0075r\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061\u0073\u0020b\u0065\u0065\u006e\u0020s\u0065\u0074 \u0077\u0068\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0052\u0047\u0042\u0020c\u006flou\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020i\u0073\u0020\u0075\u0073\u0065\u0064\u002c\u0020\u006f\u0072\u0020if\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044F\u002f\u0041\u0020\u004fut\u0070\u0075\u0074\u0049\u006e\u0074\u0065n\u0074\u0020t\u0068\u0061t\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0061\u006e\u0020\u0052\u0047\u0042\u0020\u0064\u0065\u0073\u0074\u0069\u006e\u0061\u0074io\u006e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u002e"))
					_dbbfe = true
					if _gefg() {
						return true
					}
				}
			}
		case *_f.PdfColorspaceDeviceCMYK:
			if !_daea {
				if !_acfdg || _cfddd.ColorSpace != _dfa.ColorSpaceCMYK {
					_aaff = append(_aaff, _cc("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d3", "\u0044e\u0076\u0069c\u0065\u0043\u004d\u0059\u004b\u0020\u0073hal\u006c\u0020\u006f\u006e\u006c\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u0066\u0020\u0061\u0020\u0064\u0065\u0076\u0069\u0063\u0065\u0020\u0069\u006e\u0064\u0065\u0070\u0065\u006e\u0064\u0065\u006e\u0074\u0020\u0044ef\u0061\u0075\u006c\u0074\u0043\u004d\u0059K\u0020\u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061s\u0020\u0062\u0065\u0065\u006e \u0073\u0065\u0074\u0020\u006fr \u0069\u0066\u0020\u0061\u0020\u0044e\u0076\u0069\u0063\u0065\u004e\u002d\u0062\u0061\u0073\u0065\u0064\u0020\u0044\u0065f\u0061\u0075\u006c\u0074\u0043\u004d\u0059\u004b\u0020c\u006f\u006c\u006f\u0075r\u0020\u0073\u0070\u0061\u0063e\u0020\u0068\u0061\u0073\u0020\u0062\u0065\u0065\u006e\u0020\u0073\u0065\u0074\u0020\u0077\u0068\u0065\u006e\u0020\u0074h\u0065\u0020\u0044\u0065\u0076\u0069c\u0065\u0043\u004d\u0059\u004b\u0020c\u006f\u006c\u006fu\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u0072\u0020t\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0061\u0020\u0043\u004d\u0059\u004b\u0020d\u0065\u0073\u0074\u0069\u006e\u0061t\u0069\u006f\u006e\u0020\u0070r\u006f\u0066\u0069\u006c\u0065\u002e"))
					_daea = true
					if _gefg() {
						return true
					}
				}
			}
		case *_f.PdfColorspaceICCBased:
			if !_eace {
				_bdab, _bgaf := _dfa.ParseHeader(_dbcb.Data)
				if _bgaf != nil {
					_ec.Log.Debug("\u0070\u0061\u0072si\u006e\u0067\u0020\u0049\u0043\u0043\u0042\u0061\u0073e\u0064 \u0068e\u0061d\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _bgaf)
					_aaff = append(_aaff, func() ViolatedRule {
						return _cc("\u0036.\u0032\u002e\u0034\u002e\u0032\u002d1", "\u0054\u0068e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0074\u0068\u0061\u0074\u0020\u0066o\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0073\u0074r\u0065\u0061\u006d o\u0066\u0020\u0061\u006e\u0020\u0049C\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006fl\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0049\u0043\u0043.\u0031\u003a\u0031\u0039\u0039\u0038-\u0030\u0039,\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0031\u002d\u00312\u002c\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0033\u002d\u0030\u0039\u0020\u006f\u0072\u0020I\u0053\u004f\u0020\u0031\u0035\u0030\u0037\u0036\u002d\u0031\u002e")
					}())
					_eace = true
					if _gefg() {
						return true
					}
				}
				if !_eace {
					var _abeb, _adefe bool
					switch _bdab.DeviceClass {
					case _dfa.DeviceClassPRTR, _dfa.DeviceClassMNTR, _dfa.DeviceClassSCNR, _dfa.DeviceClassSPAC:
					default:
						_abeb = true
					}
					switch _bdab.ColorSpace {
					case _dfa.ColorSpaceRGB, _dfa.ColorSpaceCMYK, _dfa.ColorSpaceGRAY, _dfa.ColorSpaceLAB:
					default:
						_adefe = true
					}
					if _abeb || _adefe {
						_aaff = append(_aaff, _cc("\u0036.\u0032\u002e\u0034\u002e\u0032\u002d1", "\u0054\u0068e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0074\u0068\u0061\u0074\u0020\u0066o\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0073\u0074r\u0065\u0061\u006d o\u0066\u0020\u0061\u006e\u0020\u0049C\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006fl\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0049\u0043\u0043.\u0031\u003a\u0031\u0039\u0039\u0038-\u0030\u0039,\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0031\u002d\u00312\u002c\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0033\u002d\u0030\u0039\u0020\u006f\u0072\u0020I\u0053\u004f\u0020\u0031\u0035\u0030\u0037\u0036\u002d\u0031\u002e"))
						_eace = true
						if _gefg() {
							return true
						}
					}
				}
			}
			if _dbcb.Alternate != nil {
				return _bdeg(_dbcb.Alternate)
			}
		}
		return false
	}
	for _, _dedbe := range _afcb.GetObjectNums() {
		_cgddf, _befa := _afcb.GetIndirectObjectByNumber(_dedbe)
		if _befa != nil {
			continue
		}
		_fgcc, _gfebbd := _cgd.GetStream(_cgddf)
		if !_gfebbd {
			continue
		}
		_efcda, _gfebbd := _cgd.GetName(_fgcc.Get("\u0054\u0079\u0070\u0065"))
		if !_gfebbd || _efcda.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_bggf, _gfebbd := _cgd.GetName(_fgcc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_gfebbd {
			continue
		}
		_gfgg[_fgcc] = struct{}{}
		switch _bggf.String() {
		case "\u0049\u006d\u0061g\u0065":
			_abecc, _fffg := _f.NewXObjectImageFromStream(_fgcc)
			if _fffg != nil {
				continue
			}
			_gfgg[_fgcc] = struct{}{}
			if _bdeg(_abecc.ColorSpace) {
				return _aaff
			}
		case "\u0046\u006f\u0072\u006d":
			_aed, _ccced := _cgd.GetDict(_fgcc.Get("\u0047\u0072\u006fu\u0070"))
			if !_ccced {
				continue
			}
			_cgae := _aed.Get("\u0043\u0053")
			if _cgae == nil {
				continue
			}
			_dffab, _geacd := _f.NewPdfColorspaceFromPdfObject(_cgae)
			if _geacd != nil {
				continue
			}
			if _bdeg(_dffab) {
				return _aaff
			}
		}
	}
	for _, _cbgcc := range _afcb.PageList {
		_egba, _bffc := _cbgcc.GetContentStreams()
		if _bffc != nil {
			continue
		}
		for _, _aggg := range _egba {
			_dadfa, _fcaea := _ff.NewContentStreamParser(_aggg).Parse()
			if _fcaea != nil {
				continue
			}
			for _, _beda := range *_dadfa {
				if len(_beda.Params) > 1 {
					continue
				}
				switch _beda.Operand {
				case "\u0042\u0049":
					_ecaf, _bgeg := _beda.Params[0].(*_ff.ContentStreamInlineImage)
					if !_bgeg {
						continue
					}
					_ffbd, _gebd := _ecaf.GetColorSpace(_cbgcc.Resources)
					if _gebd != nil {
						continue
					}
					if _bdeg(_ffbd) {
						return _aaff
					}
				case "\u0044\u006f":
					_dcdc, _ccbdg := _cgd.GetName(_beda.Params[0])
					if !_ccbdg {
						continue
					}
					_abcff, _fcfcd := _cbgcc.Resources.GetXObjectByName(*_dcdc)
					if _, _ccgcg := _gfgg[_abcff]; _ccgcg {
						continue
					}
					switch _fcfcd {
					case _f.XObjectTypeImage:
						_feca, _cddc := _f.NewXObjectImageFromStream(_abcff)
						if _cddc != nil {
							continue
						}
						_gfgg[_abcff] = struct{}{}
						if _bdeg(_feca.ColorSpace) {
							return _aaff
						}
					case _f.XObjectTypeForm:
						_edfc, _agac := _cgd.GetDict(_abcff.Get("\u0047\u0072\u006fu\u0070"))
						if !_agac {
							continue
						}
						_abcad, _agac := _cgd.GetName(_edfc.Get("\u0043\u0053"))
						if !_agac {
							continue
						}
						_eaff, _afff := _f.NewPdfColorspaceFromPdfObject(_abcad)
						if _afff != nil {
							continue
						}
						_gfgg[_abcff] = struct{}{}
						if _bdeg(_eaff) {
							return _aaff
						}
					}
				}
			}
		}
	}
	return _aaff
}
func _gfbff(_befc *_f.CompliancePdfReader) ViolatedRule {
	_bcec, _dcd := _befc.GetTrailer()
	if _dcd != nil {
		_ec.Log.Debug("\u0043\u0061\u006en\u006f\u0074\u0020\u0067e\u0074\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0076", _dcd)
		return _cb
	}
	_gefad, _baea := _bcec.Get("\u0052\u006f\u006f\u0074").(*_cgd.PdfObjectReference)
	if !_baea {
		_ec.Log.Debug("\u0043a\u006e\u006e\u006f\u0074 \u0066\u0069\u006e\u0064\u0020d\u006fc\u0075m\u0065\u006e\u0074\u0020\u0072\u006f\u006ft")
		return _cb
	}
	_cbcc, _baea := _cgd.GetDict(_cgd.ResolveReference(_gefad))
	if !_baea {
		_ec.Log.Debug("\u0063\u0061\u006e\u006e\u006f\u0074 \u0072\u0065\u0073\u006f\u006c\u0076\u0065\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		return _cb
	}
	if _cbcc.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073") != nil {
		return _cc("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063u\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u004f\u0043\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073")
	}
	return _cb
}

type profile1 struct {
	_acfg standardType
	_beb  Profile1Options
}

func _dbe(_gbad, _efba, _gbf, _dfbc string) (string, bool) {
	_eeed := _ac.Index(_gbad, _efba)
	if _eeed == -1 {
		return "", false
	}
	_daeeg := _ac.Index(_gbad, _gbf)
	if _daeeg == -1 {
		return "", false
	}
	if _daeeg < _eeed {
		return "", false
	}
	return _gbad[:_eeed] + _efba + _dfbc + _gbad[_daeeg:], true
}
func _cffa(_fgbfg *_f.PdfFont, _cbge *_cgd.PdfObjectDictionary, _cadf bool) ViolatedRule {
	const (
		_gcacdc = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0034\u002d\u0031"
		_bdbg   = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006ds\u0020\u0066\u006fr\u0020\u0061\u006c\u006c\u0020f\u006f\u006e\u0074\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0072e\u006e\u0064\u0065\u0072\u0069\u006eg\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020w\u0069t\u0068\u0069\u006e\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u002c \u0061\u0073\u0020\u0064\u0065\u0066\u0069n\u0065\u0064 \u0069\u006e\u0020\u0049S\u004f\u0020\u0033\u0032\u00300\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0039\u002e"
	)
	if _cadf {
		return _cb
	}
	_abbbb := _fgbfg.FontDescriptor()
	var _cdfb string
	if _ecfd, _decfe := _cgd.GetName(_cbge.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _decfe {
		_cdfb = _ecfd.String()
	}
	switch _cdfb {
	case "\u0054\u0079\u0070e\u0031":
		if _abbbb.FontFile == nil {
			return _cc(_gcacdc, _bdbg)
		}
	case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		if _abbbb.FontFile2 == nil {
			return _cc(_gcacdc, _bdbg)
		}
	case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0033":
	default:
		if _abbbb.FontFile3 == nil {
			return _cc(_gcacdc, _bdbg)
		}
	}
	return _cb
}
func _bdgda(_ebdf *_f.CompliancePdfReader) ViolatedRule {
	for _, _fgdf := range _ebdf.PageList {
		_fgga, _fded := _fgdf.GetContentStreams()
		if _fded != nil {
			continue
		}
		for _, _gefd := range _fgga {
			_cdagf := _ff.NewContentStreamParser(_gefd)
			_, _fded = _cdagf.Parse()
			if _fded != nil {
				return _cc("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0063onten\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079 \u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u0020\u006e\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0065\u0076\u0065\u006e\u0020\u0069\u0066\u0020s\u0075\u0063\u0068\u0020\u006f\u0070\u0065r\u0061\u0074\u006f\u0072\u0073\u0020\u0061\u0072\u0065\u0020\u0062\u0072\u0061\u0063\u006b\u0065\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0074\u0068\u0065\u0020\u0042\u0058\u002f\u0045\u0058\u0020\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u002e")
			}
		}
	}
	return _cb
}

// Profile2Options are the options that changes the way how optimizer may try to adapt document into PDF/A standard.
type Profile2Options struct {

	// CMYKDefaultColorSpace is an option that refers PDF/A
	CMYKDefaultColorSpace bool

	// Now is a function that returns current time.
	Now func() _eg.Time

	// Xmp is the xmp options information.
	Xmp XmpOptions
}

// Profile2B is the implementation of the PDF/A-2B standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2B struct{ profile2 }

// Part gets the PDF/A version level.
func (_bfff *profile2) Part() int { return _bfff._geg._ba }

// Profile1A is the implementation of the PDF/A-1A standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile1A struct{ profile1 }

func _adec(_feg *_f.PdfInfo, _dfec func() _eg.Time) error {
	var _fbbc *_f.PdfDate
	if _feg.CreationDate == nil {
		_eeag, _ecba := _f.NewPdfDateFromTime(_dfec())
		if _ecba != nil {
			return _ecba
		}
		_fbbc = &_eeag
		_feg.CreationDate = _fbbc
	}
	if _feg.ModifiedDate == nil {
		if _fbbc != nil {
			_afcgb, _cdea := _f.NewPdfDateFromTime(_dfec())
			if _cdea != nil {
				return _cdea
			}
			_fbbc = &_afcgb
		}
		_feg.ModifiedDate = _fbbc
	}
	return nil
}

// String gets a string representation of the violated rule.
func (_aef ViolatedRule) String() string {
	return _a.Sprintf("\u0025\u0073\u003a\u0020\u0025\u0073", _aef.RuleNo, _aef.Detail)
}

// Error implements error interface.
func (_gfe VerificationError) Error() string {
	_fac := _ac.Builder{}
	_fac.WriteString("\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u003a\u0020")
	_fac.WriteString(_a.Sprintf("\u0050\u0044\u0046\u002f\u0041\u002d\u0025\u0064\u0025\u0073", _gfe.ConformanceLevel, _gfe.ConformanceVariant))
	_fac.WriteString("\u0020\u0056\u0069\u006f\u006c\u0061\u0074\u0065\u0064\u0020\u0072\u0075l\u0065\u0073\u003a\u0020")
	for _ge, _fg := range _gfe.ViolatedRules {
		_fac.WriteString(_fg.String())
		if _ge != len(_gfe.ViolatedRules)-1 {
			_fac.WriteRune('\n')
		}
	}
	return _fac.String()
}
func _acbf(_faedge *_cgd.PdfObjectStream, _bgbg map[*_cgd.PdfObjectStream][]byte, _feab map[*_cgd.PdfObjectStream]*_bd.CMap) (*_bd.CMap, error) {
	_abce, _cdaa := _feab[_faedge]
	if !_cdaa {
		var _fcfd error
		_ffdd, _ceddc := _bgbg[_faedge]
		if !_ceddc {
			_ffdd, _fcfd = _cgd.DecodeStream(_faedge)
			if _fcfd != nil {
				_ec.Log.Debug("\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0073\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _fcfd)
				return nil, _fcfd
			}
			_bgbg[_faedge] = _ffdd
		}
		_abce, _fcfd = _bd.LoadCmapFromData(_ffdd, false)
		if _fcfd != nil {
			return nil, _fcfd
		}
		_feab[_faedge] = _abce
	}
	return _abce, nil
}
func _daefe(_cbgc *_f.CompliancePdfReader, _fddd standardType) (_bdfa []ViolatedRule) {
	var _ceea, _gcaef, _ebff, _cdec, _cacd, _aec, _fgac, _ffgec, _ebffd, _ddc, _ddaa bool
	_eddg := func() bool {
		return _ceea && _gcaef && _ebff && _cdec && _cacd && _aec && _fgac && _ffgec && _ebffd && _ddc && _ddaa
	}
	_eeda := map[*_cgd.PdfObjectStream]*_bd.CMap{}
	_cegga := map[*_cgd.PdfObjectStream][]byte{}
	_cdfg := map[_cgd.PdfObject]*_f.PdfFont{}
	for _, _facac := range _cbgc.GetObjectNums() {
		_cccd, _degc := _cbgc.GetIndirectObjectByNumber(_facac)
		if _degc != nil {
			continue
		}
		_ggbgb, _cfce := _cgd.GetDict(_cccd)
		if !_cfce {
			continue
		}
		_abff, _cfce := _cgd.GetName(_ggbgb.Get("\u0054\u0079\u0070\u0065"))
		if !_cfce {
			continue
		}
		if *_abff != "\u0046\u006f\u006e\u0074" {
			continue
		}
		_cdddd, _degc := _f.NewPdfFontFromPdfObject(_ggbgb)
		if _degc != nil {
			_ec.Log.Debug("g\u0065\u0074\u0074\u0069\u006e\u0067 \u0066\u006f\u006e\u0074\u0020\u0066r\u006f\u006d\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0076", _degc)
			continue
		}
		_cdfg[_ggbgb] = _cdddd
	}
	for _, _ggdf := range _cbgc.PageList {
		_ffef, _agbf := _ggdf.GetContentStreams()
		if _agbf != nil {
			_ec.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0070\u0061\u0067\u0065\u0020\u0063o\u006e\u0074\u0065\u006e\u0074\u0020\u0073t\u0072\u0065\u0061\u006d\u0073\u0020\u0066\u0061\u0069\u006ce\u0064")
			continue
		}
		for _, _gcfgc := range _ffef {
			_cgbgc := _ff.NewContentStreamParser(_gcfgc)
			_gdga, _addd := _cgbgc.Parse()
			if _addd != nil {
				_ec.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _addd)
				continue
			}
			var _gbagg bool
			for _, _dffgf := range *_gdga {
				if _dffgf.Operand != "\u0054\u0072" {
					continue
				}
				if len(_dffgf.Params) != 1 {
					_ec.Log.Debug("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054\u0072\u0027\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0065\u0064\u0020\u0027\u0031\u0027\u0020\u0062\u0075\u0074 \u0069\u0073\u003a\u0020\u0027\u0025d\u0027", len(_dffgf.Params))
					continue
				}
				_cfcd, _bffff := _cgd.GetIntVal(_dffgf.Params[0])
				if !_bffff {
					_ec.Log.Debug("\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020i\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
					continue
				}
				if _cfcd == 3 {
					_gbagg = true
					break
				}
			}
			for _, _fdfa := range *_gdga {
				if _fdfa.Operand != "\u0054\u0066" {
					continue
				}
				if len(_fdfa.Params) != 2 {
					_ec.Log.Debug("i\u006eva\u006ci\u0064 \u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066 \u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072s\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054f\u0027\u0020\u006fper\u0061\u006e\u0064\u002c\u0020\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0032\u0027\u0020\u0069s\u003a \u0027\u0025\u0064\u0027", len(_fdfa.Params))
					continue
				}
				_ecdd, _geccb := _cgd.GetName(_fdfa.Params[0])
				if !_geccb {
					_ec.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _fdfa)
					continue
				}
				_adde, _dceed := _ggdf.Resources.GetFontByName(*_ecdd)
				if !_dceed {
					_ec.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
					continue
				}
				_fbaae, _geccb := _cgd.GetDict(_adde)
				if !_geccb {
					_ec.Log.Debug("\u0066\u006f\u006e\u0074 d\u0069\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
					continue
				}
				_fcdd, _geccb := _cdfg[_fbaae]
				if !_geccb {
					var _bdeef error
					_fcdd, _bdeef = _f.NewPdfFontFromPdfObject(_fbaae)
					if _bdeef != nil {
						_ec.Log.Debug("\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0072o\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _bdeef)
						continue
					}
					_cdfg[_fbaae] = _fcdd
				}
				if !_ceea {
					_ffbgf := _gfbb(_fbaae, _cegga, _eeda)
					if _ffbgf != _cb {
						_bdfa = append(_bdfa, _ffbgf)
						_ceea = true
						if _eddg() {
							return _bdfa
						}
					}
				}
				if !_gcaef {
					_bagdb := _edef(_fbaae)
					if _bagdb != _cb {
						_bdfa = append(_bdfa, _bagdb)
						_gcaef = true
						if _eddg() {
							return _bdfa
						}
					}
				}
				if !_ebff {
					_bgead := _bebd(_fbaae, _cegga, _eeda)
					if _bgead != _cb {
						_bdfa = append(_bdfa, _bgead)
						_ebff = true
						if _eddg() {
							return _bdfa
						}
					}
				}
				if !_cdec {
					_dbgc := _bgdc(_fbaae, _cegga, _eeda)
					if _dbgc != _cb {
						_bdfa = append(_bdfa, _dbgc)
						_cdec = true
						if _eddg() {
							return _bdfa
						}
					}
				}
				if !_cacd {
					_edda := _ddff(_fcdd, _fbaae, _gbagg)
					if _edda != _cb {
						_cacd = true
						_bdfa = append(_bdfa, _edda)
						if _eddg() {
							return _bdfa
						}
					}
				}
				if !_aec {
					_bbfa := _fbddg(_fcdd, _fbaae)
					if _bbfa != _cb {
						_aec = true
						_bdfa = append(_bdfa, _bbfa)
						if _eddg() {
							return _bdfa
						}
					}
				}
				if !_fgac {
					_ecga := _gdcd(_fcdd, _fbaae)
					if _ecga != _cb {
						_fgac = true
						_bdfa = append(_bdfa, _ecga)
						if _eddg() {
							return _bdfa
						}
					}
				}
				if !_ffgec {
					_agce := _gcce(_fcdd, _fbaae)
					if _agce != _cb {
						_ffgec = true
						_bdfa = append(_bdfa, _agce)
						if _eddg() {
							return _bdfa
						}
					}
				}
				if !_ebffd {
					_ffgd := _ggea(_fcdd, _fbaae)
					if _ffgd != _cb {
						_ebffd = true
						_bdfa = append(_bdfa, _ffgd)
						if _eddg() {
							return _bdfa
						}
					}
				}
				if !_ddc {
					_geff := _bagdd(_fcdd, _fbaae)
					if _geff != _cb {
						_ddc = true
						_bdfa = append(_bdfa, _geff)
						if _eddg() {
							return _bdfa
						}
					}
				}
				if !_ddaa && _fddd._fa == "\u0041" {
					_gegd := _deaed(_fbaae, _cegga, _eeda)
					if _gegd != _cb {
						_ddaa = true
						_bdfa = append(_bdfa, _gegd)
						if _eddg() {
							return _bdfa
						}
					}
				}
			}
		}
	}
	return _bdfa
}

type documentColorspaceOptimizeFunc func(_gdag *_gg.Document, _ddd []*_gg.Image) error

// Profile1B is the implementation of the PDF/A-1B standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile1B struct{ profile1 }

func _bagdd(_ffed *_f.PdfFont, _dffdb *_cgd.PdfObjectDictionary) ViolatedRule {
	const (
		_efbb  = "\u0036.\u0033\u002e\u0037\u002d\u0033"
		_baadb = "\u0046\u006f\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006d\u0073\u0027\u0020\u0022\u0063\u006d\u0061\u0070\u0022\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u0020\u0066\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0073\u0079\u006d\u0062o\u006c\u0069c\u0020\u0054\u0072\u0075e\u0054\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073 \u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0065\u0078\u0061\u0063\u0074\u006cy\u0020\u006f\u006ee\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u002e"
	)
	var _cegf string
	if _aegb, _bggb := _cgd.GetName(_dffdb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _bggb {
		_cegf = _aegb.String()
	}
	if _cegf != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _cb
	}
	_cdead := _ffed.FontDescriptor()
	_faccf, _cbgdc := _cgd.GetIntVal(_cdead.Flags)
	if !_cbgdc {
		_ec.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _cc(_efbb, _baadb)
	}
	_afee := (uint32(_faccf) >> 3) != 0
	if !_afee {
		return _cb
	}
	return _cb
}
func _efcb(_fcgbc *_f.CompliancePdfReader) (_gafce []ViolatedRule) {
	if _fcgbc.ParserMetadata().HasOddLengthHexStrings() {
		_gafce = append(_gafce, _cc("\u0036.\u0031\u002e\u0036\u002d\u0031", "\u0068\u0065\u0078a\u0064\u0065\u0063\u0069\u006d\u0061\u006c\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006f\u0066\u0020e\u0076\u0065\u006e\u0020\u0073\u0069\u007a\u0065"))
	}
	if _fcgbc.ParserMetadata().HasOddLengthHexStrings() {
		_gafce = append(_gafce, _cc("\u0036.\u0031\u002e\u0036\u002d\u0032", "\u0068\u0065\u0078\u0061\u0064\u0065\u0063\u0069\u006da\u006c\u0020s\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068o\u0075\u006c\u0064\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u006f\u006e\u006c\u0079\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0066\u0072\u006f\u006d\u0020\u0072\u0061n\u0067\u0065\u0020[\u0030\u002d\u0039\u003b\u0041\u002d\u0046\u003b\u0061\u002d\u0066\u005d"))
	}
	return _gafce
}

// Profile2U is the implementation of the PDF/A-2U standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2U struct{ profile2 }

func _caec(_cbbd *_f.CompliancePdfReader, _edfg bool) (_eagb []ViolatedRule) {
	var _abdf, _bcab, _dbfd, _afced, _deeff, _ecda, _efebe bool
	_fedga := func() bool { return _abdf && _bcab && _dbfd && _afced && _deeff && _ecda && _efebe }
	_baec, _cbag := _gdgaf(_cbbd)
	var _gecfd _dfa.ProfileHeader
	if _cbag {
		_gecfd, _ = _dfa.ParseHeader(_baec.DestOutputProfile)
	}
	var _dbdb bool
	_cfebf := map[_cgd.PdfObject]struct{}{}
	var _dfaae func(_fadf _f.PdfColorspace) bool
	_dfaae = func(_gdfcb _f.PdfColorspace) bool {
		switch _bcbf := _gdfcb.(type) {
		case *_f.PdfColorspaceDeviceGray:
			if !_ecda {
				if !_cbag {
					_dbdb = true
					_eagb = append(_eagb, _cc("\u0036.\u0032\u002e\u0033\u002d\u0034", "\u0044\u0065\u0076\u0069\u0063\u0065G\u0072\u0061\u0079\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0075s\u0065\u0064\u0020\u006f\u006el\u0079\u0020\u0069\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006ce\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020O\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u002e"))
					_ecda = true
					if _fedga() {
						return true
					}
				}
			}
		case *_f.PdfColorspaceDeviceRGB:
			if !_afced {
				if !_cbag || _gecfd.ColorSpace != _dfa.ColorSpaceRGB {
					_dbdb = true
					_eagb = append(_eagb, _cc("\u0036.\u0032\u002e\u0033\u002d\u0032", "\u0044\u0065\u0076\u0069\u0063\u0065\u0052\u0047\u0042\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u006e\u006c\u0079\u0020\u0069\u0066\u0020\u0074\u0068\u0065 \u0066\u0069\u006c\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074In\u0074\u0065\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020u\u0073es\u0020a\u006e\u0020\u0052\u0047\u0042\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u002e"))
					_afced = true
					if _fedga() {
						return true
					}
				}
			}
		case *_f.PdfColorspaceDeviceCMYK:
			if !_deeff {
				if !_cbag || _gecfd.ColorSpace != _dfa.ColorSpaceCMYK {
					_dbdb = true
					_eagb = append(_eagb, _cc("\u0036.\u0032\u002e\u0033\u002d\u0033", "\u0044\u0065\u0076\u0069\u0063e\u0043\u004d\u0059\u004b \u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u006e\u006c\u0079\u0020\u0069\u0066\u0020\u0074h\u0065\u0020\u0066\u0069\u006ce \u0068\u0061\u0073\u0020\u0061 \u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0068a\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0061\u006e \u0043\u004d\u0059\u004b\u0020\u0063\u006f\u006c\u006f\u0072\u0020s\u0070\u0061\u0063e\u002e"))
					_deeff = true
					if _fedga() {
						return true
					}
				}
			}
		case *_f.PdfColorspaceICCBased:
			if !_dbfd || !_efebe {
				_gfcd, _fafa := _dfa.ParseHeader(_bcbf.Data)
				if _fafa != nil {
					_ec.Log.Debug("\u0070\u0061\u0072si\u006e\u0067\u0020\u0049\u0043\u0043\u0042\u0061\u0073e\u0064 \u0068e\u0061d\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _fafa)
					_eagb = append(_eagb, func() ViolatedRule {
						return _cc("\u0036.\u0032\u002e\u0033\u002d\u0031", "\u0041\u006cl \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006co\u0072\u0020\u0073\u0070a\u0063e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0061\u0073\u0020\u0049\u0043\u0043 \u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0073 \u0061\u0073\u0020d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0034\u002e\u0035")
					}())
					_dbfd = true
					if _fedga() {
						return true
					}
				}
				if !_dbfd {
					var _ffabf, _faaa bool
					switch _gfcd.DeviceClass {
					case _dfa.DeviceClassPRTR, _dfa.DeviceClassMNTR, _dfa.DeviceClassSCNR, _dfa.DeviceClassSPAC:
					default:
						_ffabf = true
					}
					switch _gfcd.ColorSpace {
					case _dfa.ColorSpaceRGB, _dfa.ColorSpaceCMYK, _dfa.ColorSpaceGRAY, _dfa.ColorSpaceLAB:
					default:
						_faaa = true
					}
					if _ffabf || _faaa {
						_eagb = append(_eagb, _cc("\u0036.\u0032\u002e\u0033\u002d\u0031", "\u0041\u006cl \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006co\u0072\u0020\u0073\u0070a\u0063e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0061\u0073\u0020\u0049\u0043\u0043 \u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0073 \u0061\u0073\u0020d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0034\u002e\u0035"))
						_dbfd = true
						if _fedga() {
							return true
						}
					}
				}
				if !_efebe {
					_fcac, _ := _cgd.GetStream(_bcbf.GetContainingPdfObject())
					if _fcac.Get("\u004e") == nil || (_bcbf.N == 1 && _gfcd.ColorSpace != _dfa.ColorSpaceGRAY) || (_bcbf.N == 3 && !(_gfcd.ColorSpace == _dfa.ColorSpaceRGB || _gfcd.ColorSpace == _dfa.ColorSpaceLAB)) || (_bcbf.N == 4 && _gfcd.ColorSpace != _dfa.ColorSpaceCMYK) {
						_eagb = append(_eagb, _cc("\u0036.\u0032\u002e\u0033\u002d\u0035", "\u0049\u0066\u0020a\u006e\u0020u\u006e\u0063\u0061\u006c\u0069\u0062\u0072a\u0074\u0065\u0064\u0020\u0063\u006fl\u006f\u0072 \u0073\u0070\u0061c\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u0069\u006c\u0065 \u0074\u0068\u0065\u006e \u0074\u0068\u0061\u0074 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041-\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065d\u0020\u0069\u006e\u0020\u0036\u002e\u0032\u002e\u0032\u002e"))
						_efebe = true
						if _fedga() {
							return true
						}
					}
				}
			}
			if _bcbf.Alternate != nil {
				return _dfaae(_bcbf.Alternate)
			}
		}
		return false
	}
	for _, _aeeea := range _cbbd.GetObjectNums() {
		_badf, _bggg := _cbbd.GetIndirectObjectByNumber(_aeeea)
		if _bggg != nil {
			continue
		}
		_fbagd, _ccbb := _cgd.GetStream(_badf)
		if !_ccbb {
			continue
		}
		_gbfe, _ccbb := _cgd.GetName(_fbagd.Get("\u0054\u0079\u0070\u0065"))
		if !_ccbb || _gbfe.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_dfca, _ccbb := _cgd.GetName(_fbagd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_ccbb {
			continue
		}
		_cfebf[_fbagd] = struct{}{}
		switch _dfca.String() {
		case "\u0049\u006d\u0061g\u0065":
			_efac, _cfgae := _f.NewXObjectImageFromStream(_fbagd)
			if _cfgae != nil {
				continue
			}
			_cfebf[_fbagd] = struct{}{}
			if _dfaae(_efac.ColorSpace) {
				return _eagb
			}
		case "\u0046\u006f\u0072\u006d":
			_gdcb, _deccb := _cgd.GetDict(_fbagd.Get("\u0047\u0072\u006fu\u0070"))
			if !_deccb {
				continue
			}
			_befg := _gdcb.Get("\u0043\u0053")
			if _befg == nil {
				continue
			}
			_dffa, _fcfea := _f.NewPdfColorspaceFromPdfObject(_befg)
			if _fcfea != nil {
				continue
			}
			if _dfaae(_dffa) {
				return _eagb
			}
		}
	}
	for _, _dcea := range _cbbd.PageList {
		_gadb, _gbbce := _dcea.GetContentStreams()
		if _gbbce != nil {
			continue
		}
		for _, _aagc := range _gadb {
			_cadc, _abbf := _ff.NewContentStreamParser(_aagc).Parse()
			if _abbf != nil {
				continue
			}
			for _, _cddeg := range *_cadc {
				if len(_cddeg.Params) > 1 {
					continue
				}
				switch _cddeg.Operand {
				case "\u0042\u0049":
					_defbf, _agb := _cddeg.Params[0].(*_ff.ContentStreamInlineImage)
					if !_agb {
						continue
					}
					_egbb, _dfdb := _defbf.GetColorSpace(_dcea.Resources)
					if _dfdb != nil {
						continue
					}
					if _dfaae(_egbb) {
						return _eagb
					}
				case "\u0044\u006f":
					_bead, _fdf := _cgd.GetName(_cddeg.Params[0])
					if !_fdf {
						continue
					}
					_ccgc, _aafcf := _dcea.Resources.GetXObjectByName(*_bead)
					if _, _cfbga := _cfebf[_ccgc]; _cfbga {
						continue
					}
					switch _aafcf {
					case _f.XObjectTypeImage:
						_fbdd, _faaaa := _f.NewXObjectImageFromStream(_ccgc)
						if _faaaa != nil {
							continue
						}
						_cfebf[_ccgc] = struct{}{}
						if _dfaae(_fbdd.ColorSpace) {
							return _eagb
						}
					case _f.XObjectTypeForm:
						_dgag, _gcfg := _cgd.GetDict(_ccgc.Get("\u0047\u0072\u006fu\u0070"))
						if !_gcfg {
							continue
						}
						_egee, _gcfg := _cgd.GetName(_dgag.Get("\u0043\u0053"))
						if !_gcfg {
							continue
						}
						_gadba, _dcdf := _f.NewPdfColorspaceFromPdfObject(_egee)
						if _dcdf != nil {
							continue
						}
						_cfebf[_ccgc] = struct{}{}
						if _dfaae(_gadba) {
							return _eagb
						}
					}
				}
			}
		}
	}
	if !_dbdb {
		return _eagb
	}
	if (_gecfd.DeviceClass == _dfa.DeviceClassPRTR || _gecfd.DeviceClass == _dfa.DeviceClassMNTR) && (_gecfd.ColorSpace == _dfa.ColorSpaceRGB || _gecfd.ColorSpace == _dfa.ColorSpaceCMYK || _gecfd.ColorSpace == _dfa.ColorSpaceGRAY) {
		return _eagb
	}
	if !_edfg {
		return _eagb
	}
	_dbbdb, _bccb := _ecgd(_cbbd)
	if !_bccb {
		return _eagb
	}
	_bcgg, _bccb := _cgd.GetArray(_dbbdb.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_bccb {
		_eagb = append(_eagb, _cc("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u0020\u0069\u0073\u0020a\u006e \u004f\u0075\u0074\u0070\u0075\u0074\u0049n\u0074\u0065\u006e\u0074\u0020\u0064i\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0062y\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0031\u0030.4\u002c\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073 \u0069\u006e\u0063\u006c\u0075\u0064e\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0027\u0073\u0020O\u0075\u0074p\u0075\u0074I\u006e\u0074\u0065\u006e\u0074\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020a\u006e\u0064\u0020h\u0061\u0073\u0020\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0041\u0031\u0020\u0061\u0073 \u0074\u0068\u0065\u0020\u0076a\u006c\u0075e\u0020\u006f\u0066\u0020i\u0074\u0073 \u0053\u0020\u006b\u0065\u0079\u0020\u0061\u006e\u0064\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020I\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006ce\u0020s\u0074\u0072\u0065\u0061\u006d \u0061\u0073\u0020\u0074h\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0074\u0073\u0020\u0044\u0065\u0073t\u004f\u0075t\u0070\u0075\u0074P\u0072\u006f\u0066\u0069\u006c\u0065 \u006b\u0065\u0079\u002e"), _cc("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e"))
		return _eagb
	}
	if _bcgg.Len() > 1 {
		_bgcef := map[*_cgd.PdfObjectDictionary]struct{}{}
		for _dba := 0; _dba < _bcgg.Len(); _dba++ {
			_effa, _gddg := _cgd.GetDict(_bcgg.Get(_dba))
			if !_gddg {
				continue
			}
			if _dba == 0 {
				_bgcef[_effa] = struct{}{}
				continue
			}
			if _, _cged := _bgcef[_effa]; !_cged {
				_eagb = append(_eagb, _cc("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e"))
				break
			}
		}
	}
	return _eagb
}

// ApplyStandard tries to change the content of the writer to match the PDF/A-1 standard.
// Implements model.StandardApplier.
func (_edgd *profile1) ApplyStandard(document *_gg.Document) (_fcaee error) {
	_fff(document, 4)
	if _fcaee = _bbe(document, _edgd._beb.Now); _fcaee != nil {
		return _fcaee
	}
	if _fcaee = _cgbf(document); _fcaee != nil {
		return _fcaee
	}
	_ebce, _bacc := _eag(_edgd._beb.CMYKDefaultColorSpace, _edgd._acfg)
	_fcaee = _ecd(document, []pageColorspaceOptimizeFunc{_efge, _ebce}, []documentColorspaceOptimizeFunc{_bacc})
	if _fcaee != nil {
		return _fcaee
	}
	_gde(document)
	if _fcaee = _dac(document, _edgd._acfg._ba); _fcaee != nil {
		return _fcaee
	}
	if _fcaee = _bdfc(document); _fcaee != nil {
		return _fcaee
	}
	if _fcaee = _gcbf(document); _fcaee != nil {
		return _fcaee
	}
	if _fcaee = _gfbf(document); _fcaee != nil {
		return _fcaee
	}
	if _fcaee = _bdd(document); _fcaee != nil {
		return _fcaee
	}
	if _edgd._acfg._fa == "\u0041" {
		_cfeg(document)
	}
	if _fcaee = _bde(document, _edgd._acfg._ba); _fcaee != nil {
		return _fcaee
	}
	if _fcaee = _gfbg(document); _fcaee != nil {
		return _fcaee
	}
	if _eggff := _fede(document, _edgd._acfg, _edgd._beb.Xmp); _eggff != nil {
		return _eggff
	}
	if _edgd._acfg == _gb() {
		if _fcaee = _age(document); _fcaee != nil {
			return _fcaee
		}
	}
	if _fcaee = _cac(document); _fcaee != nil {
		return _fcaee
	}
	return nil
}

type profile2 struct {
	_geg standardType
	_ddb Profile2Options
}

func _beaa(_fcea *Profile1Options) {
	if _fcea.Now == nil {
		_fcea.Now = _eg.Now
	}
}

// Profile1Options are the options that changes the way how optimizer may try to adapt document into PDF/A standard.
type Profile1Options struct {

	// CMYKDefaultColorSpace is an option that refers PDF/A-1
	CMYKDefaultColorSpace bool

	// Now is a function that returns current time.
	Now func() _eg.Time

	// Xmp is the xmp options information.
	Xmp XmpOptions
}

// DefaultProfile1Options are the default options for the Profile1.
func DefaultProfile1Options() *Profile1Options {
	return &Profile1Options{Now: _eg.Now, Xmp: XmpOptions{MarshalIndent: "\u0009"}}
}
func _eeaf(_bdgfe *_f.CompliancePdfReader) (_bgebb []ViolatedRule) {
	var (
		_befcb, _ccgce, _deaab, _afagc, _bceda bool
		_dgdg                                  func(_cgd.PdfObject)
	)
	_dgdg = func(_bgecc _cgd.PdfObject) {
		switch _dbba := _bgecc.(type) {
		case *_cgd.PdfObjectInteger:
			if !_befcb && (int64(*_dbba) > _e.MaxInt32 || int64(*_dbba) < -_e.MaxInt32) {
				_bgebb = append(_bgebb, _cc("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0031", "L\u0061\u0072\u0067e\u0073\u0074\u0020\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0032\u002c\u0031\u0034\u0037,\u0034\u0038\u0033,\u0036\u0034\u0037\u002e\u0020\u0053\u006d\u0061\u006c\u006c\u0065\u0073\u0074 \u0069\u006e\u0074\u0065g\u0065\u0072\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u002d\u0032\u002c\u0031\u0034\u0037\u002c\u0034\u0038\u0033,\u0036\u0034\u0038\u002e"))
				_befcb = true
			}
		case *_cgd.PdfObjectFloat:
			if !_ccgce && (_e.Abs(float64(*_dbba)) > _e.MaxFloat32) {
				_bgebb = append(_bgebb, _cc("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0032", "\u0041 \u0063\u006f\u006e\u0066orm\u0069\u006e\u0067\u0020f\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0061\u006c\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u006f\u0075\u0074\u0073\u0069de\u0020\u0074\u0068e\u0020\u0072\u0061\u006e\u0067e\u0020o\u0066\u0020\u002b\u002f\u002d\u0033\u002e\u0034\u00303\u0020\u0078\u0020\u0031\u0030\u005e\u0033\u0038\u002e"))
			}
		case *_cgd.PdfObjectString:
			if !_deaab && len([]byte(_dbba.Str())) > 32767 {
				_bgebb = append(_bgebb, _cc("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0033", "M\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006c\u0065n\u0067\u0074\u0068\u0020\u006f\u0066\u0020a \u0073\u0074\u0072\u0069n\u0067\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074es\u0029\u0020i\u0073\u0020\u0033\u0032\u0037\u0036\u0037\u002e"))
				_deaab = true
			}
		case *_cgd.PdfObjectName:
			if !_afagc && len([]byte(*_dbba)) > 127 {
				_bgebb = append(_bgebb, _cc("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0034", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006c\u0065\u006eg\u0074\u0068\u0020\u006ff\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0029\u0020\u0069\u0073\u0020\u0031\u0032\u0037\u002e"))
				_afagc = true
			}
		case *_cgd.PdfObjectArray:
			for _, _gcgag := range _dbba.Elements() {
				_dgdg(_gcgag)
			}
			if !_bceda && (_dbba.Len() == 4 || _dbba.Len() == 5) {
				_ceec, _bbgd := _cgd.GetName(_dbba.Get(0))
				if !_bbgd {
					return
				}
				if *_ceec != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
					return
				}
				_aegbc := _dbba.Get(1)
				_aegbc = _cgd.TraceToDirectObject(_aegbc)
				_fgdaa, _bbgd := _cgd.GetArray(_aegbc)
				if !_bbgd {
					return
				}
				if _fgdaa.Len() > 32 {
					_bgebb = append(_bgebb, _cc("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0039", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006e\u0075\u006db\u0065\u0072\u0020\u006ff\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0069\u0073\u0020\u0033\u0032\u002e"))
					_bceda = true
				}
			}
		case *_cgd.PdfObjectDictionary:
			_ddcf := _dbba.Keys()
			for _fgeb, _aadbc := range _ddcf {
				_dgdg(&_ddcf[_fgeb])
				_dgdg(_dbba.Get(_aadbc))
			}
		case *_cgd.PdfObjectStream:
			_dgdg(_dbba.PdfObjectDictionary)
		case *_cgd.PdfObjectStreams:
			for _, _cadb := range _dbba.Elements() {
				_dgdg(_cadb)
			}
		case *_cgd.PdfObjectReference:
			_dgdg(_dbba.Resolve())
		}
	}
	_cdbac := _bdgfe.GetObjectNums()
	if len(_cdbac) > 8388607 {
		_bgebb = append(_bgebb, _cc("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0037", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020in\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073 \u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0069\u0073\u00208\u002c\u0033\u0038\u0038\u002c\u0036\u0030\u0037\u002e"))
	}
	for _, _abcg := range _cdbac {
		_dfcaf, _adfa := _bdgfe.GetIndirectObjectByNumber(_abcg)
		if _adfa != nil {
			continue
		}
		_bfegb := _cgd.TraceToDirectObject(_dfcaf)
		_dgdg(_bfegb)
	}
	return _bgebb
}
func _ged(_feebg standardType, _ggeb *_gg.OutputIntents) error {
	_gff, _cgce := _dfa.NewISOCoatedV2Gray1CBasOutputIntent(_feebg.outputIntentSubtype())
	if _cgce != nil {
		return _cgce
	}
	if _cgce = _ggeb.Add(_gff.ToPdfObject()); _cgce != nil {
		return _cgce
	}
	return nil
}
func _egcf(_gcee *_f.CompliancePdfReader) (_cdgdf []ViolatedRule) {
	_dbec := _gcee.ParserMetadata()
	if _dbec.HasInvalidSubsectionHeader() {
		_cdgdf = append(_cdgdf, _cc("\u0036.\u0031\u002e\u0034\u002d\u0031", "\u006e\u0020\u0061\u0020\u0063\u0072\u006f\u0073\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0073\u0065c\u0074\u0069\u006f\u006e\u0020h\u0065a\u0064\u0065\u0072\u0020t\u0068\u0065\u0020\u0073\u0074\u0061\u0072t\u0069\u006e\u0067\u0020\u006fb\u006a\u0065\u0063\u0074 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0072\u0061n\u0067e\u0020s\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020s\u0069\u006e\u0067\u006c\u0065\u0020\u0053\u0050\u0041C\u0045\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0020\u0028\u0032\u0030\u0068\u0029\u002e"))
	}
	if _dbec.HasInvalidSeparationAfterXRef() {
		_cdgdf = append(_cdgdf, _cc("\u0036.\u0031\u002e\u0034\u002d\u0032", "\u0054\u0068\u0065 \u0078\u0072\u0065\u0066\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0063\u0072\u006f\u0073s\u0020\u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065 s\u0075b\u0073\u0065\u0063ti\u006f\u006e\u0020\u0068\u0065\u0061\u0064e\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079 \u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e"))
	}
	return _cdgdf
}

var _ Profile = (*Profile2B)(nil)
var _cb = ViolatedRule{}

func _aaga(_fgbe *_gg.Document, _fbd bool) error {
	_cec, _gaf := _fgbe.GetPages()
	if !_gaf {
		return nil
	}
	for _, _gad := range _cec {
		_dbbf, _gdagb := _gad.GetContents()
		if !_gdagb {
			continue
		}
		var _cfeb *_f.PdfPageResources
		_aafe, _gdagb := _gad.GetResources()
		if _gdagb {
			_cfeb, _ = _f.NewPdfPageResourcesFromDict(_aafe)
		}
		for _bbd, _cbafe := range _dbbf {
			_efeg, _ffe := _cbafe.GetData()
			if _ffe != nil {
				continue
			}
			_gbde := _ff.NewContentStreamParser(string(_efeg))
			_dfd, _ffe := _gbde.Parse()
			if _ffe != nil {
				continue
			}
			_abab, _ffe := _defc(_cfeb, _dfd, _fbd)
			if _ffe != nil {
				return _ffe
			}
			if _abab == nil {
				continue
			}
			if _ffe = (&_dbbf[_bbd]).SetData(_abab); _ffe != nil {
				return _ffe
			}
		}
	}
	return nil
}
func _gcbf(_cee *_gg.Document) error {
	_bgfg, _gab := _cee.GetPages()
	if !_gab {
		return nil
	}
	for _, _cdb := range _bgfg {
		_adff, _afbc := _cgd.GetArray(_cdb.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_afbc {
			continue
		}
		for _, _adbd := range _adff.Elements() {
			_adbd = _cgd.ResolveReference(_adbd)
			if _, _bbdc := _adbd.(*_cgd.PdfObjectNull); _bbdc {
				continue
			}
			_bddg, _agga := _cgd.GetDict(_adbd)
			if !_agga {
				continue
			}
			_cbdg, _ := _cgd.GetIntVal(_bddg.Get("\u0046"))
			_cbdg &= ^(1 << 0)
			_cbdg &= ^(1 << 1)
			_cbdg &= ^(1 << 5)
			_cbdg |= 1 << 2
			_bddg.Set("\u0046", _cgd.MakeInteger(int64(_cbdg)))
			_ffg := false
			if _fcga := _bddg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _fcga != nil {
				_aeeg, _acaa := _cgd.GetName(_fcga)
				if _acaa && _aeeg.String() == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					_ffg = true
					if _bddg.Get("\u0041\u0041") != nil {
						_bddg.Remove("\u0041\u0041")
					}
				}
			}
			if _bddg.Get("\u0043") != nil || _bddg.Get("\u0049\u0043") != nil {
				_fage, _cggc := _febb(_cee)
				if !_cggc {
					_bddg.Remove("\u0043")
					_bddg.Remove("\u0049\u0043")
				} else {
					_bgg, _fcge := _cgd.GetIntVal(_fage.Get("\u004e"))
					if !_fcge || _bgg != 3 {
						_bddg.Remove("\u0043")
						_bddg.Remove("\u0049\u0043")
					}
				}
			}
			_bdee, _agga := _cgd.GetDict(_bddg.Get("\u0041\u0050"))
			if _agga {
				_bdgd := _bdee.Get("\u004e")
				if _bdgd == nil {
					continue
				}
				if len(_bdee.Keys()) > 1 {
					_bdee.Clear()
					_bdee.Set("\u004e", _bdgd)
				}
				if _ffg {
					_ccdfga, _cgee := _cgd.GetName(_bddg.Get("\u0046\u0054"))
					if _cgee && *_ccdfga == "\u0042\u0074\u006e" {
						continue
					}
				}
			}
		}
	}
	return nil
}
func _eged(_aaea *_f.CompliancePdfReader) ViolatedRule { return _cb }
func _fggc(_bbgf *_f.CompliancePdfReader) ViolatedRule {
	_cdgc, _eccb := _bbgf.PdfReader.GetTrailer()
	if _eccb != nil {
		return _cc("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u006d\u0069\u0073s\u0069\u006e\u0067\u0020t\u0072\u0061\u0069\u006c\u0065\u0072\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	}
	if _cdgc.Get("\u0049\u0044") == nil {
		return _cc("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068e\u0020\u0027\u0049\u0044\u0027\u0020k\u0065\u0079\u0077o\u0072\u0064")
	}
	if _cdgc.Get("\u0045n\u0063\u0072\u0079\u0070\u0074") != nil {
		return _cc("\u0036.\u0031\u002e\u0033\u002d\u0032", "\u0054\u0068\u0065\u0020\u006b\u0065y\u0077\u006f\u0072\u0064\u0020'\u0045\u006e\u0063\u0072\u0079\u0070t\u0027\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0075\u0073\u0065d\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u002e\u0020")
	}
	return _cb
}
func _efge(_dda *_gg.Document, _aafc *_gg.Page, _aafca []*_gg.Image) error {
	for _, _eefe := range _aafca {
		if _eefe.SMask == nil {
			continue
		}
		_begc, _dgeg := _f.NewXObjectImageFromStream(_eefe.Stream)
		if _dgeg != nil {
			return _dgeg
		}
		_bab, _dgeg := _begc.ToImage()
		if _dgeg != nil {
			return _dgeg
		}
		_bba, _dgeg := _bab.ToGoImage()
		if _dgeg != nil {
			return _dgeg
		}
		_ffc, _dgeg := _cf.RGBAConverter.Convert(_bba)
		if _dgeg != nil {
			return _dgeg
		}
		_dagf := _ffc.Base()
		_bgea := &_f.Image{Width: int64(_dagf.Width), Height: int64(_dagf.Height), BitsPerComponent: int64(_dagf.BitsPerComponent), ColorComponents: _dagf.ColorComponents, Data: _dagf.Data}
		_bgea.SetDecode(_dagf.Decode)
		_bgea.SetAlpha(_dagf.Alpha)
		if _dgeg = _begc.SetImage(_bgea, nil); _dgeg != nil {
			return _dgeg
		}
		_begc.SMask = _cgd.MakeNull()
		var _cfgce _cgd.PdfObject
		_baf := -1
		for _baf, _cfgce = range _dda.Objects {
			if _cfgce == _eefe.SMask.Stream {
				break
			}
		}
		if _baf != -1 {
			_dda.Objects = append(_dda.Objects[:_baf], _dda.Objects[_baf+1:]...)
		}
		_eefe.SMask = nil
		_begc.ToPdfObject()
	}
	return nil
}
func _gaee(_abgae *_f.CompliancePdfReader) ViolatedRule { return _cb }
func _cbff(_faff *_f.CompliancePdfReader) (_bdgc []ViolatedRule) {
	var _dede, _fcead, _acaab, _aggc, _fdae, _eegcd, _cdgac bool
	_geef := func() bool { return _dede && _fcead && _acaab && _aggc && _fdae && _eegcd && _cdgac }
	_cgdb := func(_fgebg *_cgd.PdfObjectDictionary) bool {
		if !_dede && _fgebg.Get("\u0054\u0052") != nil {
			_dede = true
			_bdgc = append(_bdgc, _cc("\u0036.\u0032\u002e\u0035\u002d\u0031", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0054\u0052\u0020\u006b\u0065\u0079\u002e"))
		}
		if _dcad := _fgebg.Get("\u0054\u0052\u0032"); !_fcead && _dcad != nil {
			_dcfd, _cggac := _cgd.GetName(_dcad)
			if !_cggac || (_cggac && *_dcfd != "\u0044e\u0066\u0061\u0075\u006c\u0074") {
				_fcead = true
				_bdgc = append(_bdgc, _cc("\u0036.\u0032\u002e\u0035\u002d\u0032", "\u0041\u006e \u0045\u0078\u0074G\u0053\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074a\u0069n\u0020\u0074\u0068\u0065\u0020\u0054R2 \u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076al\u0075e\u0020\u006f\u0074\u0068e\u0072 \u0074h\u0061\u006e \u0044\u0065fa\u0075\u006c\u0074\u002e"))
				if _geef() {
					return true
				}
			}
		}
		if !_acaab && _fgebg.Get("\u0048\u0054\u0050") != nil {
			_acaab = true
			_bdgc = append(_bdgc, _cc("\u0036.\u0032\u002e\u0035\u002d\u0033", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020th\u0065\u0020\u0048\u0054\u0050\u0020\u006b\u0065\u0079\u002e"))
		}
		_dbcce, _edbf := _cgd.GetDict(_fgebg.Get("\u0048\u0054"))
		if _edbf {
			if _ggega := _dbcce.Get("\u0048\u0061\u006cf\u0074\u006f\u006e\u0065\u0054\u0079\u0070\u0065"); !_aggc && _ggega != nil {
				_ccdb, _cfba := _cgd.GetInt(_ggega)
				if !_cfba || (_cfba && !(*_ccdb == 1 || *_ccdb == 5)) {
					_bdgc = append(_bdgc, _cc("\u0020\u0036\u002e\u0032\u002e\u0035\u002d\u0034", "\u0041\u006c\u006c\u0020\u0068\u0061\u006c\u0066\u0074\u006f\u006e\u0065\u0073\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006ce\u0020\u0073h\u0061\u006c\u006c\u0020h\u0061\u0076\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0031\u0020\u006f\u0072\u0020\u0035 \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0048\u0061l\u0066\u0074\u006fn\u0065\u0054\u0079\u0070\u0065\u0020\u006be\u0079\u002e"))
					if _geef() {
						return true
					}
				}
			}
			if _gbeec := _dbcce.Get("\u0048\u0061\u006cf\u0074\u006f\u006e\u0065\u004e\u0061\u006d\u0065"); !_fdae && _gbeec != nil {
				_fdae = true
				_bdgc = append(_bdgc, _cc("\u0036.\u0032\u002e\u0035\u002d\u0035", "\u0048\u0061\u006c\u0066\u0074o\u006e\u0065\u0073\u0020\u0069\u006e\u0020a\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020\u0061\u0020\u0048\u0061\u006c\u0066\u0074\u006f\u006e\u0065N\u0061\u006d\u0065\u0020\u006b\u0065y\u002e"))
				if _geef() {
					return true
				}
			}
		}
		_, _fbdde := _gdgaf(_faff)
		var _bgab bool
		_ada, _edbf := _cgd.GetDict(_fgebg.Get("\u0047\u0072\u006fu\u0070"))
		if _edbf {
			_, _ceab := _cgd.GetName(_ada.Get("\u0043\u0053"))
			if _ceab {
				_bgab = true
			}
		}
		if _caea := _fgebg.Get("\u0042\u004d"); !_eegcd && !_cdgac && _caea != nil {
			_dcgf, _agacf := _cgd.GetName(_caea)
			if _agacf {
				switch _dcgf.String() {
				case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065", "\u004d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "\u0053\u0063\u0072\u0065\u0065\u006e", "\u004fv\u0065\u0072\u006c\u0061\u0079", "\u0044\u0061\u0072\u006b\u0065\u006e", "\u004ci\u0067\u0068\u0074\u0065\u006e", "\u0043\u006f\u006c\u006f\u0072\u0044\u006f\u0064\u0067\u0065", "\u0043o\u006c\u006f\u0072\u0042\u0075\u0072n", "\u0048a\u0072\u0064\u004c\u0069\u0067\u0068t", "\u0053o\u0066\u0074\u004c\u0069\u0067\u0068t", "\u0044\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065", "\u0045x\u0063\u006c\u0075\u0073\u0069\u006fn", "\u0048\u0075\u0065", "\u0053\u0061\u0074\u0075\u0072\u0061\u0074\u0069\u006f\u006e", "\u0043\u006f\u006co\u0072", "\u004c\u0075\u006d\u0069\u006e\u006f\u0073\u0069\u0074\u0079":
				default:
					_eegcd = true
					_bdgc = append(_bdgc, _cc("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0031", "\u004f\u006el\u0079\u0020\u0062\u006c\u0065\u006e\u0064\u0020\u006d\u006f\u0064\u0065\u0073\u0020\u0074h\u0061\u0074\u0020\u0061\u0072\u0065\u0020\u0073\u0070\u0065c\u0069\u0066\u0069ed\u0020\u0069\u006e\u0020\u0049\u0053O\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a2\u0030\u0030\u0038\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075e\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0042M\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0065\u0078t\u0065\u006e\u0064\u0065\u0064\u0020\u0067\u0072\u0061\u0070\u0068\u0069\u0063\u0020\u0073\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					if _geef() {
						return true
					}
				}
				if _dcgf.String() != "\u004e\u006f\u0072\u006d\u0061\u006c" && !_fbdde && !_bgab {
					_cdgac = true
					_bdgc = append(_bdgc, _cc("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
					if _geef() {
						return true
					}
				}
			}
		}
		if _, _edbf = _cgd.GetDict(_fgebg.Get("\u0053\u004d\u0061s\u006b")); !_cdgac && _edbf && !_fbdde && !_bgab {
			_cdgac = true
			_bdgc = append(_bdgc, _cc("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
			if _geef() {
				return true
			}
		}
		if _bbgg := _fgebg.Get("\u0043\u0041"); !_cdgac && _bbgg != nil && !_fbdde && !_bgab {
			_eebd, _bggfe := _cgd.GetNumberAsFloat(_bbgg)
			if _bggfe == nil && _eebd < 1.0 {
				_cdgac = true
				_bdgc = append(_bdgc, _cc("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
				if _geef() {
					return true
				}
			}
		}
		if _aefc := _fgebg.Get("\u0063\u0061"); !_cdgac && _aefc != nil && !_fbdde && !_bgab {
			_eaae, _dfbdf := _cgd.GetNumberAsFloat(_aefc)
			if _dfbdf == nil && _eaae < 1.0 {
				_cdgac = true
				_bdgc = append(_bdgc, _cc("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
				if _geef() {
					return true
				}
			}
		}
		return false
	}
	for _, _eecca := range _faff.PageList {
		_ddcfc := _eecca.Resources
		if _ddcfc == nil {
			continue
		}
		if _ddcfc.ExtGState == nil {
			continue
		}
		_bcbe, _aadf := _cgd.GetDict(_ddcfc.ExtGState)
		if !_aadf {
			continue
		}
		_cedf := _bcbe.Keys()
		for _, _gced := range _cedf {
			_bfcb, _cbaad := _cgd.GetDict(_bcbe.Get(_gced))
			if !_cbaad {
				continue
			}
			if _cgdb(_bfcb) {
				return _bdgc
			}
		}
	}
	for _, _acec := range _faff.PageList {
		_bcce := _acec.Resources
		if _bcce == nil {
			continue
		}
		_ggagf, _gebdd := _cgd.GetDict(_bcce.XObject)
		if !_gebdd {
			continue
		}
		for _, _gdcc := range _ggagf.Keys() {
			_daeac, _gcdb := _cgd.GetStream(_ggagf.Get(_gdcc))
			if !_gcdb {
				continue
			}
			_ddbg, _gcdb := _cgd.GetDict(_daeac.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_gcdb {
				continue
			}
			_egge, _gcdb := _cgd.GetDict(_ddbg.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
			if !_gcdb {
				continue
			}
			for _, _eefed := range _egge.Keys() {
				_ffbb, _fgfg := _cgd.GetDict(_egge.Get(_eefed))
				if !_fgfg {
					continue
				}
				if _cgdb(_ffbb) {
					return _bdgc
				}
			}
		}
	}
	return _bdgc
}
func _gcfgd(_cgfe *_f.CompliancePdfReader) (_cebd []ViolatedRule) {
	var _deag, _eggaf, _accc, _effd, _eca, _dafa bool
	_dbfb := func() bool { return _deag && _eggaf && _accc && _effd && _eca && _dafa }
	_bbab := func(_cacf *_cgd.PdfObjectDictionary) bool {
		if !_deag && _cacf.Get("\u0054\u0052") != nil {
			_deag = true
			_cebd = append(_cebd, _cc("\u0036.\u0032\u002e\u0038\u002d\u0031", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0054\u0052\u0020\u006b\u0065\u0079\u002e"))
		}
		if _fdcd := _cacf.Get("\u0054\u0052\u0032"); !_eggaf && _fdcd != nil {
			_acdec, _edcd := _cgd.GetName(_fdcd)
			if !_edcd || (_edcd && *_acdec != "\u0044e\u0066\u0061\u0075\u006c\u0074") {
				_eggaf = true
				_cebd = append(_cebd, _cc("\u0036.\u0032\u002e\u0038\u002d\u0032", "\u0041\u006e \u0045\u0078\u0074G\u0053\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074a\u0069n\u0020\u0074\u0068\u0065\u0020\u0054R2 \u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076al\u0075e\u0020\u006f\u0074\u0068e\u0072 \u0074h\u0061\u006e \u0044\u0065fa\u0075\u006c\u0074\u002e"))
				if _dbfb() {
					return true
				}
			}
		}
		if _cdbe := _cacf.Get("\u0053\u004d\u0061s\u006b"); !_accc && _cdbe != nil {
			_cfage, _fbcec := _cgd.GetName(_cdbe)
			if !_fbcec || (_fbcec && *_cfage != "\u004e\u006f\u006e\u0065") {
				_accc = true
				_cebd = append(_cebd, _cc("\u0036\u002e\u0034-\u0031", "\u0049\u0066\u0020\u0061\u006e \u0053\u004d\u0061\u0073\u006b\u0020\u006be\u0079\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0069\u0074s\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u004e\u006f\u006ee\u002e"))
				if _dbfb() {
					return true
				}
			}
		}
		if _dcgcg := _cacf.Get("\u0043\u0041"); !_eca && _dcgcg != nil {
			_cfbgc, _ccbd := _cgd.GetNumberAsFloat(_dcgcg)
			if _ccbd == nil && _cfbgc != 1.0 {
				_eca = true
				_cebd = append(_cebd, _cc("\u0036\u002e\u0034-\u0035", "\u0054\u0068\u0065\u0020\u0066ol\u006c\u006fw\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0073\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078t\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068a\u006c\u006c\u0020\u0068\u0061v\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0073 \u0073h\u006f\u0077\u006e\u003a\u0020\u0043\u0041 \u002d\u0020\u0031\u002e\u0030\u002e"))
				if _dbfb() {
					return true
				}
			}
		}
		if _aabbb := _cacf.Get("\u0063\u0061"); !_dafa && _aabbb != nil {
			_ddda, _abga := _cgd.GetNumberAsFloat(_aabbb)
			if _abga == nil && _ddda != 1.0 {
				_dafa = true
				_cebd = append(_cebd, _cc("\u0036\u002e\u0034-\u0036", "\u0054\u0068\u0065\u0020\u0066ol\u006c\u006fw\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0073\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078t\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068a\u006c\u006c\u0020\u0068\u0061v\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0073 \u0073h\u006f\u0077\u006e\u003a\u0020\u0063\u0061 \u002d\u0020\u0031\u002e\u0030\u002e"))
				if _dbfb() {
					return true
				}
			}
		}
		if _dgbe := _cacf.Get("\u0042\u004d"); !_effd && _dgbe != nil {
			_edfec, _dgfg := _cgd.GetName(_dgbe)
			if _dgfg {
				switch _edfec.String() {
				case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065":
				default:
					_effd = true
					_cebd = append(_cebd, _cc("\u0036\u002e\u0034-\u0034", "T\u0068\u0065\u0020\u0066\u006f\u006cl\u006f\u0077\u0069\u006e\u0067 \u006b\u0065y\u0073\u002c\u0020\u0069\u0066 \u0070res\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078\u0074\u0047S\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065 \u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0073\u0068\u006f\u0077n\u003a\u0020\u0042\u004d\u0020\u002d\u0020\u004e\u006f\u0072m\u0061\u006c\u0020\u006f\u0072\u0020\u0043\u006f\u006d\u0070\u0061t\u0069\u0062\u006c\u0065\u002e"))
					if _dbfb() {
						return true
					}
				}
			}
		}
		return false
	}
	for _, _dedd := range _cgfe.PageList {
		_dcgce := _dedd.Resources
		if _dcgce == nil {
			continue
		}
		if _dcgce.ExtGState == nil {
			continue
		}
		_befgb, _aeac := _cgd.GetDict(_dcgce.ExtGState)
		if !_aeac {
			continue
		}
		_cfabf := _befgb.Keys()
		for _, _ffd := range _cfabf {
			_efgg, _edcb := _cgd.GetDict(_befgb.Get(_ffd))
			if !_edcb {
				continue
			}
			if _bbab(_efgg) {
				return _cebd
			}
		}
	}
	for _, _gceg := range _cgfe.PageList {
		_gbfdg := _gceg.Resources
		if _gbfdg == nil {
			continue
		}
		_ebea, _geab := _cgd.GetDict(_gbfdg.XObject)
		if !_geab {
			continue
		}
		for _, _eagf := range _ebea.Keys() {
			_cfae, _fcddc := _cgd.GetStream(_ebea.Get(_eagf))
			if !_fcddc {
				continue
			}
			_dgae, _fcddc := _cgd.GetDict(_cfae.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_fcddc {
				continue
			}
			_abde, _fcddc := _cgd.GetDict(_dgae.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
			if !_fcddc {
				continue
			}
			for _, _gbgd := range _abde.Keys() {
				_eaed, _deda := _cgd.GetDict(_abde.Get(_gbgd))
				if !_deda {
					continue
				}
				if _bbab(_eaed) {
					return _cebd
				}
			}
		}
	}
	return _cebd
}
func _acac() standardType { return standardType{_ba: 2, _fa: "\u0042"} }
func _adee(_fgdg *_f.CompliancePdfReader) (_fedgb ViolatedRule) {
	for _, _dbee := range _fgdg.GetObjectNums() {
		_cea, _fcee := _fgdg.GetIndirectObjectByNumber(_dbee)
		if _fcee != nil {
			continue
		}
		_eggd, _gaad := _cgd.GetStream(_cea)
		if !_gaad {
			continue
		}
		_gffd, _gaad := _cgd.GetName(_eggd.Get("\u0054\u0079\u0070\u0065"))
		if !_gaad {
			continue
		}
		if *_gffd != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_fcfa, _gaad := _cgd.GetName(_eggd.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"))
		if !_gaad {
			continue
		}
		if *_fcfa == "\u0050\u0053" {
			return _cc("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		if _eggd.Get("\u0050\u0053") != nil {
			return _cc("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _fedgb
}

// NewProfile2B creates a new Profile2B with the given options.
func NewProfile2B(options *Profile2Options) *Profile2B {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_baab(options)
	return &Profile2B{profile2{_ddb: *options, _geg: _acac()}}
}
func _gcbc(_edabf *_f.CompliancePdfReader) (_deaf []ViolatedRule) { return _deaf }
func _gfbg(_fed *_gg.Document) error {
	_edf, _bgf := _fed.FindCatalog()
	if !_bgf {
		return _d.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_, _bgf = _cgd.GetDict(_edf.Object.Get("\u0041\u0041"))
	if !_bgf {
		return nil
	}
	_edf.Object.Remove("\u0041\u0041")
	return nil
}
func _cdef(_fdaa *_f.CompliancePdfReader) ViolatedRule { return _cb }
func _gdcd(_ebgbf *_f.PdfFont, _dcefa *_cgd.PdfObjectDictionary) ViolatedRule {
	const (
		_eage = "\u0036.\u0033\u002e\u0035\u002d\u0033"
		_bdcb = "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0073 \u0072e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0077i\u0074\u0068\u0069n\u0020\u0061\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069l\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006et\u0020\u0064\u0065s\u0063\u0072\u0069\u0070\u0074\u006f\u0072\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020\u0043\u0049\u0044\u0053\u0065\u0074\u0020s\u0074\u0072\u0065\u0061\u006d\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0079\u0069\u006eg\u0020\u0077\u0068i\u0063\u0068\u0020\u0043\u0049\u0044\u0073 \u0061\u0072e\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e \u0074\u0068\u0065\u0020\u0065\u006d\u0062\u0065\u0064d\u0065\u0064\u0020\u0043\u0049D\u0046\u006f\u006e\u0074\u0020\u0066\u0069l\u0065,\u0020\u0061\u0073 \u0064\u0065\u0073\u0063\u0072\u0069b\u0065\u0064 \u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063e\u0020\u0054ab\u006c\u0065\u0020\u0035.\u00320\u002e"
	)
	var _ggee string
	if _bfge, _cacg := _cgd.GetName(_dcefa.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cacg {
		_ggee = _bfge.String()
	}
	switch _ggee {
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_bcabe := _ebgbf.FontDescriptor()
		if _bcabe.CIDSet == nil {
			return _cc(_eage, _bdcb)
		}
		return _cb
	default:
		return _cb
	}
}

// Profile is the model.StandardImplementer enhanced by the information about the profile conformance level.
type Profile interface {
	_f.StandardImplementer
	Conformance() string
	Part() int
}

func _aafg(_bgge *_f.CompliancePdfReader) []ViolatedRule { return nil }
func _adcg(_gcae *_gg.Document, _efc bool) error {
	_deca, _gaff := _gcae.GetPages()
	if !_gaff {
		return nil
	}
	for _, _cfff := range _deca {
		_abed := _cfff.FindXObjectForms()
		for _, _bfaf := range _abed {
			_fagd, _becg := _f.NewXObjectFormFromStream(_bfaf)
			if _becg != nil {
				return _becg
			}
			_gbdc, _becg := _fagd.GetContentStream()
			if _becg != nil {
				return _becg
			}
			_aaed := _ff.NewContentStreamParser(string(_gbdc))
			_gcaa, _becg := _aaed.Parse()
			if _becg != nil {
				return _becg
			}
			_ccf, _becg := _defc(_fagd.Resources, _gcaa, _efc)
			if _becg != nil {
				return _becg
			}
			if len(_ccf) == 0 {
				continue
			}
			if _becg = _fagd.SetContentStream(_ccf, _cgd.NewFlateEncoder()); _becg != nil {
				return _becg
			}
			_fagd.ToPdfObject()
		}
	}
	return nil
}
func _ecd(_acae *_gg.Document, _fbaa []pageColorspaceOptimizeFunc, _accg []documentColorspaceOptimizeFunc) error {
	_efdd, _fbca := _acae.GetPages()
	if !_fbca {
		return nil
	}
	var _cab []*_gg.Image
	for _gec, _cbdb := range _efdd {
		_fgb, _ggac := _cbdb.FindXObjectImages()
		if _ggac != nil {
			return _ggac
		}
		for _, _eeca := range _fbaa {
			if _ggac = _eeca(_acae, &_efdd[_gec], _fgb); _ggac != nil {
				return _ggac
			}
		}
		_cab = append(_cab, _fgb...)
	}
	for _, _cgc := range _accg {
		if _dcef := _cgc(_acae, _cab); _dcef != nil {
			return _dcef
		}
	}
	return nil
}
func _ecgd(_bbbf *_f.CompliancePdfReader) (*_cgd.PdfObjectDictionary, bool) {
	_cfgde, _dfaec := _bbbf.GetTrailer()
	if _dfaec != nil {
		_ec.Log.Debug("\u0043\u0061\u006en\u006f\u0074\u0020\u0067e\u0074\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0076", _dfaec)
		return nil, false
	}
	_ggeef, _beaad := _cfgde.Get("\u0052\u006f\u006f\u0074").(*_cgd.PdfObjectReference)
	if !_beaad {
		_ec.Log.Debug("\u0043a\u006e\u006e\u006f\u0074 \u0066\u0069\u006e\u0064\u0020d\u006fc\u0075m\u0065\u006e\u0074\u0020\u0072\u006f\u006ft")
		return nil, false
	}
	_eeae, _beaad := _cgd.GetDict(_cgd.ResolveReference(_ggeef))
	if !_beaad {
		_ec.Log.Debug("\u0063\u0061\u006e\u006e\u006f\u0074 \u0072\u0065\u0073\u006f\u006c\u0076\u0065\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		return nil, false
	}
	return _eeae, true
}

type imageInfo struct {
	ColorSpace       _cgd.PdfObjectName
	BitsPerComponent int
	ColorComponents  int
	Width            int
	Height           int
	Stream           *_cgd.PdfObjectStream
	_bac             bool
}

func _afdg(_ega *_f.CompliancePdfReader) (_gcd []ViolatedRule) {
	_bcdg, _ebcdc := _ecgd(_ega)
	if !_ebcdc {
		return _gcd
	}
	_dca := _cc("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u0020\u0069\u0073\u0020a\u006e \u004f\u0075\u0074\u0070\u0075\u0074\u0049n\u0074\u0065\u006e\u0074\u0020\u0064i\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0062y\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0031\u0030.4\u002c\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073 \u0069\u006e\u0063\u006c\u0075\u0064e\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0027\u0073\u0020O\u0075\u0074p\u0075\u0074I\u006e\u0074\u0065\u006e\u0074\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020a\u006e\u0064\u0020h\u0061\u0073\u0020\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0041\u0031\u0020\u0061\u0073 \u0074\u0068\u0065\u0020\u0076a\u006c\u0075e\u0020\u006f\u0066\u0020i\u0074\u0073 \u0053\u0020\u006b\u0065\u0079\u0020\u0061\u006e\u0064\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020I\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006ce\u0020s\u0074\u0072\u0065\u0061\u006d \u0061\u0073\u0020\u0074h\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0074\u0073\u0020\u0044\u0065\u0073t\u004f\u0075t\u0070\u0075\u0074P\u0072\u006f\u0066\u0069\u006c\u0065 \u006b\u0065\u0079\u002e")
	_ded, _ebcdc := _cgd.GetArray(_bcdg.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_ebcdc {
		_gcd = append(_gcd, _dca)
		return _gcd
	}
	_ccag := _cc("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e")
	if _ded.Len() > 1 {
		_cecc := map[*_cgd.PdfObjectDictionary]struct{}{}
		for _dfee := 0; _dfee < _ded.Len(); _dfee++ {
			_bdgb, _gfc := _cgd.GetDict(_ded.Get(_dfee))
			if !_gfc {
				_gcd = append(_gcd, _dca)
				return _gcd
			}
			if _dfee == 0 {
				_cecc[_bdgb] = struct{}{}
				continue
			}
			if _, _eggfb := _cecc[_bdgb]; !_eggfb {
				_gcd = append(_gcd, _ccag)
				break
			}
		}
	} else if _ded.Len() == 0 {
		_gcd = append(_gcd, _dca)
		return _gcd
	}
	_gcaeg, _ebcdc := _cgd.GetDict(_ded.Get(0))
	if !_ebcdc {
		_gcd = append(_gcd, _dca)
		return _gcd
	}
	if _daec, _fdgg := _cgd.GetName(_gcaeg.Get("\u0053")); !_fdgg || (*_daec) != "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411" {
		_gcd = append(_gcd, _dca)
		return _gcd
	}
	_cbfg, _cdf := _f.NewPdfOutputIntentFromPdfObject(_gcaeg)
	if _cdf != nil {
		_ec.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020i\u006et\u0065\u006e\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _cdf)
		return _gcd
	}
	_ffcd, _cdf := _dfa.ParseHeader(_cbfg.DestOutputProfile)
	if _cdf != nil {
		_ec.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061d\u0065\u0072\u0020\u0066\u0061i\u006c\u0065d\u003a\u0020\u0025\u0076", _cdf)
		return _gcd
	}
	if (_ffcd.DeviceClass == _dfa.DeviceClassPRTR || _ffcd.DeviceClass == _dfa.DeviceClassMNTR) && (_ffcd.ColorSpace == _dfa.ColorSpaceRGB || _ffcd.ColorSpace == _dfa.ColorSpaceCMYK || _ffcd.ColorSpace == _dfa.ColorSpaceGRAY) {
		return _gcd
	}
	_gcd = append(_gcd, _dca)
	return _gcd
}

// Conformance gets the PDF/A conformance.
func (_cdag *profile2) Conformance() string { return _cdag._geg._fa }

// NewProfile2A creates a new Profile2A with given options.
func NewProfile2A(options *Profile2Options) *Profile2A {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_baab(options)
	return &Profile2A{profile2{_ddb: *options, _geg: _fae()}}
}
func _cegc(_fcbd *_f.CompliancePdfReader) (_dded []ViolatedRule) {
	var _dffe, _addde, _dceee, _aceg, _ebcg, _gdaf, _gffb bool
	_fefd := map[*_cgd.PdfObjectStream]struct{}{}
	for _, _gdbcd := range _fcbd.GetObjectNums() {
		if _dffe && _addde && _ebcg && _dceee && _aceg && _gdaf && _gffb {
			return _dded
		}
		_bfaea, _aaaaa := _fcbd.GetIndirectObjectByNumber(_gdbcd)
		if _aaaaa != nil {
			continue
		}
		_fdea, _aecge := _cgd.GetStream(_bfaea)
		if !_aecge {
			continue
		}
		if _, _aecge = _fefd[_fdea]; _aecge {
			continue
		}
		_fefd[_fdea] = struct{}{}
		_abbg, _aecge := _cgd.GetName(_fdea.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_aecge {
			continue
		}
		if !_aceg {
			if _fdea.Get("\u0052\u0065\u0066") != nil {
				_dded = append(_dded, _cc("\u0036.\u0032\u002e\u0039\u002d\u0032", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e"))
				_aceg = true
			}
		}
		if _abbg.String() == "\u0050\u0053" {
			if !_gdaf {
				_dded = append(_dded, _cc("\u0036.\u0032\u002e\u0039\u002d\u0033", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073."))
				_gdaf = true
				continue
			}
		}
		if _abbg.String() == "\u0046\u006f\u0072\u006d" {
			if _addde && _dceee && _aceg {
				continue
			}
			if !_addde && _fdea.Get("\u004f\u0050\u0049") != nil {
				_dded = append(_dded, _cc("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072\u006d \u0058\u004f\u0062j\u0065\u0063\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004f\u0050\u0049\u0020\u006b\u0065\u0079\u002e"))
				_addde = true
			}
			if !_dceee {
				if _fdea.Get("\u0050\u0053") != nil {
					_dceee = true
				}
				if _acccd := _fdea.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _acccd != nil && !_dceee {
					if _fbdcf, _aeeee := _cgd.GetName(_acccd); _aeeee && *_fbdcf == "\u0050\u0053" {
						_dceee = true
					}
				}
				if _dceee {
					_dded = append(_dded, _cc("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065y \u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076\u0061\u006cu\u0065 o\u0066 \u0050\u0053\u0020\u0061\u006e\u0064\u0020t\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e"))
				}
			}
			continue
		}
		if _abbg.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if !_dffe && _fdea.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
			_dded = append(_dded, _cc("\u0036.\u0032\u002e\u0038\u002d\u0031", "\u0041\u006e\u0020\u0049m\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
			_dffe = true
		}
		if !_gffb && _fdea.Get("\u004f\u0050\u0049") != nil {
			_dded = append(_dded, _cc("\u0036.\u0032\u002e\u0038\u002d\u0032", "\u0041\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0068\u0065\u0020\u004f\u0050\u0049\u0020\u006b\u0065\u0079\u002e"))
			_gffb = true
		}
		if !_ebcg && _fdea.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065") != nil {
			_aeaf, _gabg := _cgd.GetBool(_fdea.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _gabg && bool(*_aeaf) {
				continue
			}
			_dded = append(_dded, _cc("\u0036.\u0032\u002e\u0038\u002d\u0033", "\u0049\u0066 a\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0063o\u006e\u0074\u0061\u0069n\u0073\u0020\u0074\u0068e \u0049\u006et\u0065r\u0070\u006f\u006c\u0061\u0074\u0065 \u006b\u0065\u0079,\u0020\u0069t\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020b\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
			_ebcg = true
		}
	}
	return _dded
}
func _cgef(_gbee *_gg.Document) error {
	_ceba, _cgba := _gbee.FindCatalog()
	if !_cgba {
		return _d.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_abf, _cgba := _cgd.GetDict(_ceba.Object.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"))
	if !_cgba {
		return nil
	}
	_bfce, _cgba := _cgd.GetDict(_abf.Get("\u0044"))
	if _cgba {
		if _bfce.Get("\u0041\u0053") != nil {
			_bfce.Remove("\u0041\u0053")
		}
	}
	_eddc, _cgba := _cgd.GetArray(_abf.Get("\u0043o\u006e\u0066\u0069\u0067\u0073"))
	if _cgba {
		for _cbgd := 0; _cbgd < _eddc.Len(); _cbgd++ {
			_fdaf, _bfcg := _cgd.GetDict(_eddc.Get(_cbgd))
			if !_bfcg {
				continue
			}
			if _fdaf.Get("\u0041\u0053") != nil {
				_fdaf.Remove("\u0041\u0053")
			}
		}
	}
	return nil
}
func (_baa standardType) String() string {
	return _a.Sprintf("\u0050\u0044\u0046\u002f\u0041\u002d\u0025\u0064\u0025\u0073", _baa._ba, _baa._fa)
}
func _baab(_gbbgg *Profile2Options) {
	if _gbbgg.Now == nil {
		_gbbgg.Now = _eg.Now
	}
}
func _dbdf(_gbfc *_f.CompliancePdfReader) (_aege []ViolatedRule) {
	if _gbfc.ParserMetadata().HasOddLengthHexStrings() {
		_aege = append(_aege, _cc("\u0036.\u0031\u002e\u0036\u002d\u0031", "\u0068\u0065\u0078a\u0064\u0065\u0063\u0069\u006d\u0061\u006c\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006f\u0066\u0020e\u0076\u0065\u006e\u0020\u0073\u0069\u007a\u0065"))
	}
	if _gbfc.ParserMetadata().HasOddLengthHexStrings() {
		_aege = append(_aege, _cc("\u0036.\u0031\u002e\u0036\u002d\u0032", "\u0068\u0065\u0078\u0061\u0064\u0065\u0063\u0069\u006da\u006c\u0020s\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068o\u0075\u006c\u0064\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u006f\u006e\u006c\u0079\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0066\u0072\u006f\u006d\u0020\u0072\u0061n\u0067\u0065\u0020[\u0030\u002d\u0039\u003b\u0041\u002d\u0046\u003b\u0061\u002d\u0066\u005d"))
	}
	return _aege
}
func _cgbf(_dgc *_gg.Document) error {
	_fgec, _cde := _dgc.FindCatalog()
	if !_cde {
		return _d.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_fgec.SetVersion()
	return nil
}

var _ Profile = (*Profile2U)(nil)

func (_bgdf *documentImages) hasOnlyDeviceRGB() bool { return _bgdf._fde && !_bgdf._bg && !_bgdf._bgd }
func _bdfc(_ebec *_gg.Document) error {
	_eaeg, _abb := _ebec.GetPages()
	if !_abb {
		return nil
	}
	for _, _eba := range _eaeg {
		_aead := _eba.FindXObjectForms()
		for _, _edgg := range _aead {
			_ggga, _eccg := _cgd.GetDict(_edgg.Get("\u0047\u0072\u006fu\u0070"))
			if _eccg {
				if _gcea := _ggga.Get("\u0053"); _gcea != nil {
					_gaaf, _dfcb := _cgd.GetName(_gcea)
					if _dfcb && _gaaf.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						_edgg.Remove("\u0047\u0072\u006fu\u0070")
					}
				}
			}
		}
		_aced, _cgfb := _eba.GetResourcesXObject()
		if _cgfb {
			_cge, _cbab := _cgd.GetDict(_aced.Get("\u0047\u0072\u006fu\u0070"))
			if _cbab {
				_cdga := _cge.Get("\u0053")
				if _cdga != nil {
					_cfgd, _facc := _cgd.GetName(_cdga)
					if _facc && _cfgd.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						_aced.Remove("\u0047\u0072\u006fu\u0070")
					}
				}
			}
		}
		_aegg, _fdd := _cgd.GetDict(_eba.Object.Get("\u0047\u0072\u006fu\u0070"))
		if _fdd {
			_dfff := _aegg.Get("\u0053")
			if _dfff != nil {
				_ceca, _dafe := _cgd.GetName(_dfff)
				if _dafe && _ceca.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
					_eba.Object.Remove("\u0047\u0072\u006fu\u0070")
				}
			}
		}
	}
	return nil
}
func _adeeb(_acfe *_f.CompliancePdfReader) (_fbeed ViolatedRule) {
	_ffce, _baeeb := _ecgd(_acfe)
	if !_baeeb {
		return _cb
	}
	if _ffce.Get("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073") != nil {
		return _cc("\u0036\u002e\u0031\u0031\u002d\u0031", "Th\u0065\u0020d\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063a\u0074\u0061\u006c\u006f\u0067\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020R\u0065q\u0075\u0069\u0072\u0065\u006d\u0065\u006e\u0074s\u0020k\u0065\u0079.")
	}
	return _cb
}

var _ Profile = (*Profile1A)(nil)

func _ddff(_bacba *_f.PdfFont, _fgdc *_cgd.PdfObjectDictionary, _dcbcg bool) ViolatedRule {
	const (
		_efca = "\u0036.\u0033\u002e\u0034\u002d\u0031"
		_gfeb = "\u0054\u0068\u0065\u0020\u0066\u006f\u006et\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006d\u0073\u0020\u0066\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0069\u006e \u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069l\u0065\u0020s\u0068\u0061\u006cl\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0077\u0069\u0074\u0068i\u006e\u0020\u0074h\u0061\u0074\u0020\u0066\u0069\u006ce\u002c\u0020a\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052e\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0035\u002e\u0038\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0077h\u0065\u006e\u0020\u0074\u0068\u0065 \u0066\u006f\u006e\u0074\u0073\u0020\u0061\u0072\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0065\u0078\u0063\u006cu\u0073i\u0076\u0065\u006c\u0079\u0020\u0077\u0069t\u0068\u0020\u0074\u0065\u0078\u0074\u0020\u0072e\u006ed\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020\u0033\u002e"
	)
	if _dcbcg {
		return _cb
	}
	_dgegc := _bacba.FontDescriptor()
	var _ddecb string
	if _ggadb, _ccdg := _cgd.GetName(_fgdc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _ccdg {
		_ddecb = _ggadb.String()
	}
	switch _ddecb {
	case "\u0054\u0079\u0070e\u0031":
		if _dgegc.FontFile == nil {
			return _cc(_efca, _gfeb)
		}
	case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		if _dgegc.FontFile2 == nil {
			return _cc(_efca, _gfeb)
		}
	case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0033":
	default:
		if _dgegc.FontFile3 == nil {
			return _cc(_efca, _gfeb)
		}
	}
	return _cb
}
func _dcbe(_edaaf *_f.CompliancePdfReader) ViolatedRule {
	if _edaaf.ParserMetadata().HeaderPosition() != 0 {
		return _cc("\u0036.\u0031\u002e\u0032\u002d\u0031", "h\u0065\u0061\u0064\u0065\u0072\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0074\u0020\u0074\u0068\u0065\u0020fi\u0072\u0073\u0074 \u0062y\u0074\u0065")
	}
	if _edaaf.PdfVersion().Major != 1 {
		return _cc("\u0036.\u0031\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0065\u0061\u0064e\u0072 \u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0073\u0069s\u0074 \u006f\u0066\u0020\u201c%\u0050\u0044\u0046\u002d\u0031\u002e\u006e\u201d\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065 \u0045\u004f\u004c\u0020ma\u0072\u006b\u0065\u0072\u002c \u0077\u0068\u0065\u0072\u0065\u0020\u0027\u006e\u0027\u0020\u0069s\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0064\u0069\u0067\u0069t\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0030\u0020(\u0033\u0030h\u0029\u0020\u0061\u006e\u0064\u0020\u0037\u0020\u0028\u0033\u0037\u0068\u0029")
	}
	if _edaaf.PdfVersion().Minor < 0 || _edaaf.PdfVersion().Minor > 7 {
		return _cc("\u0036.\u0031\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0065\u0061\u0064e\u0072 \u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0073\u0069s\u0074 \u006f\u0066\u0020\u201c%\u0050\u0044\u0046\u002d\u0031\u002e\u006e\u201d\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065 \u0045\u004f\u004c\u0020ma\u0072\u006b\u0065\u0072\u002c \u0077\u0068\u0065\u0072\u0065\u0020\u0027\u006e\u0027\u0020\u0069s\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0064\u0069\u0067\u0069t\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0030\u0020(\u0033\u0030h\u0029\u0020\u0061\u006e\u0064\u0020\u0037\u0020\u0028\u0033\u0037\u0068\u0029")
	}
	return _cb
}
func _dgcc(_efee *_f.CompliancePdfReader) ViolatedRule {
	_aafce := _efee.ParserMetadata().HeaderCommentBytes()
	if _aafce[0] > 127 && _aafce[1] > 127 && _aafce[2] > 127 && _aafce[3] > 127 {
		return _cb
	}
	return _cc("\u0036.\u0031\u002e\u0032\u002d\u0032", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u006c\u0069\u006e\u0065\u0020\u0073\u0068\u0061\u006c\u006c b\u0065\u0020i\u006d\u006d\u0065\u0064\u0069a\u0074\u0065\u006c\u0079 \u0066\u006f\u006c\u006co\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0063\u006f\u006d\u006d\u0065n\u0074\u0020\u0063\u006f\u006e\u0073\u0069s\u0074\u0069\u006e\u0067\u0020o\u0066\u0020\u0061\u0020\u0025\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006fwe\u0064\u0020\u0062y\u0020a\u0074\u0009\u006c\u0065a\u0073\u0074\u0020f\u006f\u0075\u0072\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u002c\u0020e\u0061\u0063\u0068\u0020\u006f\u0066\u0020\u0077\u0068\u006f\u0073\u0065 \u0065\u006e\u0063\u006f\u0064e\u0064\u0020\u0062\u0079\u0074e\u0020\u0076\u0061\u006c\u0075\u0065s\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0064e\u0063\u0069\u006d\u0061\u006c \u0076\u0061\u006c\u0075\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0031\u0032\u0037\u002e")
}
func _ggeg(_fafgb *_f.CompliancePdfReader) (_baba ViolatedRule) {
	_dafb, _effe := _ecgd(_fafgb)
	if !_effe {
		return _cb
	}
	if _dafb.Get("\u0041\u0041") != nil {
		return _cc("\u0036.\u0036\u002e\u0032\u002d\u0033", "\u0054\u0068e\u0020\u0064\u006f\u0063\u0075\u006d\u0065n\u0074 \u0063\u0061\u0074a\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065n\u0074r\u0079 \u0066\u006f\u0072 \u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063\u0074i\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
	}
	return _cb
}
func _edef(_fdada *_cgd.PdfObjectDictionary) ViolatedRule {
	const (
		_fcbbc  = "\u0036.\u0033\u002e\u0033\u002d\u0032"
		_dfaaed = "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0054y\u0070\u0065\u0020\u0032\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072\u0065\u0020\u0075\u0073\u0065\u0064\u0020f\u006f\u0072 \u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067,\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0046\u006fn\u0074\u0020\u0064\u0069c\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c \u0063\u006f\u006e\u0074\u0061i\u006e\u0020\u0061\u0020\u0043\u0049\u0044\u0054\u006f\u0047\u0049D\u004d\u0061\u0070\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020a\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006d\u0061\u0070\u0070\u0069\u006e\u0067\u0020\u0066\u0072\u006f\u006d\u0020\u0043\u0049\u0044\u0073\u0020\u0074\u006f\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u0069\u006e\u0064\u0069c\u0065\u0073\u0020\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u0049d\u0065\u006e\u0074\u0069\u0074\u0079\u002c\u0020\u0061s d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020P\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0054a\u0062\u006c\u0065\u0020\u0035\u002e\u00313"
	)
	var _fcfg string
	if _dab, _cceg := _cgd.GetName(_fdada.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cceg {
		_fcfg = _dab.String()
	}
	if _fcfg != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		return _cb
	}
	if _fdada.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070") == nil {
		return _cc(_fcbbc, _dfaaed)
	}
	return _cb
}
func _cadce(_cafc *_f.CompliancePdfReader) (_gdbfb ViolatedRule) {
	for _, _fcgd := range _cafc.GetObjectNums() {
		_acbg, _cbagb := _cafc.GetIndirectObjectByNumber(_fcgd)
		if _cbagb != nil {
			continue
		}
		_cbbcc, _cefd := _cgd.GetStream(_acbg)
		if !_cefd {
			continue
		}
		_fgba, _cefd := _cgd.GetName(_cbbcc.Get("\u0054\u0079\u0070\u0065"))
		if !_cefd {
			continue
		}
		if *_fgba != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		if _cbbcc.Get("\u0052\u0065\u0066") != nil {
			return _cc("\u0036.\u0032\u002e\u0036\u002d\u0031", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e")
		}
	}
	return _gdbfb
}
func _cabc(_dfge *_f.CompliancePdfReader) ViolatedRule {
	_ccga := _dfge.ParserMetadata()
	if _ccga.HasInvalidSeparationAfterXRef() {
		return _cc("\u0036.\u0031\u002e\u0034\u002d\u0032", "\u0054\u0068\u0065 \u0078\u0072\u0065\u0066\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0063\u0072\u006f\u0073s\u0020\u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065 s\u0075b\u0073\u0065\u0063ti\u006f\u006e\u0020\u0068\u0065\u0061\u0064e\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079 \u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e")
	}
	return _cb
}
func _age(_faeaf *_gg.Document) error {
	_dgd, _gfa := _faeaf.FindCatalog()
	if !_gfa {
		return nil
	}
	_, _gfa = _cgd.GetDict(_dgd.Object.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074"))
	if !_gfa {
		_bbc := _cgd.MakeDict()
		_bbc.Set("\u0054\u0079\u0070\u0065", _cgd.MakeName("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074"))
		_dgd.Object.Set("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074", _bbc)
	}
	return nil
}

// NewProfile2U creates a new Profile2U with the given options.
func NewProfile2U(options *Profile2Options) *Profile2U {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_baab(options)
	return &Profile2U{profile2{_ddb: *options, _geg: _gf()}}
}
func _cfaf(_dcgc *_f.CompliancePdfReader) ViolatedRule {
	_gfdd, _cced := _dcgc.PdfReader.GetTrailer()
	if _cced != nil {
		return _cc("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u006d\u0069\u0073s\u0069\u006e\u0067\u0020t\u0072\u0061\u0069\u006c\u0065\u0072\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	}
	if _gfdd.Get("\u0049\u0044") == nil {
		return _cc("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068e\u0020\u0027\u0049\u0044\u0027\u0020k\u0065\u0079\u0077o\u0072\u0064")
	}
	if _gfdd.Get("\u0045n\u0063\u0072\u0079\u0070\u0074") != nil {
		return _cc("\u0036.\u0031\u002e\u0033\u002d\u0032", "\u0054\u0068\u0065\u0020\u006b\u0065y\u0077\u006f\u0072\u0064\u0020'\u0045\u006e\u0063\u0072\u0079\u0070t\u0027\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0075\u0073\u0065d\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u002e\u0020")
	}
	return _cb
}
func _bgdc(_ggef *_cgd.PdfObjectDictionary, _fdeba map[*_cgd.PdfObjectStream][]byte, _bfeb map[*_cgd.PdfObjectStream]*_bd.CMap) ViolatedRule {
	const (
		_bedg = "\u0036.\u0033\u002e\u0033\u002d\u0034"
		_afdc = "\u0046\u006f\u0072\u0020\u0074\u0068\u006fs\u0065\u0020\u0043\u004d\u0061\u0070\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072e\u0020\u0065m\u0062\u0065\u0064de\u0064\u002c\u0020\u0074\u0068\u0065\u0020\u0069\u006et\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0057\u004d\u006f\u0064\u0065\u0020\u0065\u006e\u0074r\u0079\u0020i\u006e t\u0068\u0065\u0020CM\u0061\u0070\u0020\u0064\u0069\u0063\u0074\u0069o\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u0020\u0074\u006f \u0074h\u0065\u0020\u0057\u004d\u006f\u0064e\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064ed\u0020\u0043\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"
	)
	var _cgbb string
	if _affac, _deff := _cgd.GetName(_ggef.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _deff {
		_cgbb = _affac.String()
	}
	if _cgbb != "\u0054\u0079\u0070e\u0030" {
		return _cb
	}
	_fgedg := _ggef.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _, _ffbaa := _cgd.GetName(_fgedg); _ffbaa {
		return _cb
	}
	_gacc, _dggg := _cgd.GetStream(_fgedg)
	if !_dggg {
		return _cc(_bedg, _afdc)
	}
	_ceef, _fdbe := _acbf(_gacc, _fdeba, _bfeb)
	if _fdbe != nil {
		return _cc(_bedg, _afdc)
	}
	_gafd, _fbcc := _cgd.GetIntVal(_gacc.Get("\u0057\u004d\u006fd\u0065"))
	_gfdc, _dbbfc := _ceef.WMode()
	if _fbcc && _dbbfc {
		if _gfdc != _gafd {
			return _cc(_bedg, _afdc)
		}
	}
	if (_fbcc && !_dbbfc) || (!_fbcc && _dbbfc) {
		return _cc(_bedg, _afdc)
	}
	return _cb
}
func _bgeccf(_bada *_f.CompliancePdfReader) (_gefaf []ViolatedRule) {
	_efga := true
	_agccd, _gfdf := _bada.GetCatalogMarkInfo()
	if !_gfdf {
		_efga = false
	} else {
		_gafff, _bcga := _cgd.GetDict(_agccd)
		if _bcga {
			_cdbdc, _cccec := _cgd.GetBool(_gafff.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
			if !bool(*_cdbdc) || !_cccec {
				_efga = false
			}
		} else {
			_efga = false
		}
	}
	if !_efga {
		_gefaf = append(_gefaf, _cc("\u0036.\u0037\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006cog\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020M\u0061r\u006b\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061 \u004d\u0061\u0072\u006b\u0065\u0064\u0020\u0065\u006et\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074,\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0076\u0061lu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0072\u0075\u0065"))
	}
	_fbbg, _gfdf := _bada.GetCatalogStructTreeRoot()
	if !_gfdf {
		_gefaf = append(_gefaf, _cc("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075r\u0065\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065d \u0062\u0079\u0020a\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072e\u0020\u0068\u0069\u0065\u0072\u0061\u0072\u0063\u0068\u0079\u0020\u0072\u006f\u006ft\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0065\u006e\u0074r\u0079\u0020\u006f\u0066\u0020\u0074h\u0065\u0020d\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061t\u0061\u006c\u006fg \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0036\u002e"))
	}
	_eabe, _gfdf := _cgd.GetDict(_fbbg)
	if _gfdf {
		_fffb, _fbagf := _cgd.GetName(_eabe.Get("\u0052o\u006c\u0065\u004d\u0061\u0070"))
		if _fbagf {
			_bgaa, _gaced := _cgd.GetDict(_fffb)
			if _gaced {
				for _, _cfcfa := range _bgaa.Keys() {
					_deagd := _bgaa.Get(_cfcfa)
					if _deagd == nil {
						_gefaf = append(_gefaf, _cc("\u0036.\u0037\u002e\u0033\u002d\u0032", "\u0041\u006c\u006c\u0020\u006eo\u006e\u002ds\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0073t\u0072\u0075\u0063\u0074ure\u0020\u0074\u0079\u0070\u0065s\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u006d\u0061\u0070\u0070\u0065d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020n\u0065\u0061\u0072\u0065\u0073\u0074\u0020\u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0073\u0074a\u006ed\u0061r\u0064\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065re\u006e\u0063e\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020i\u006e\u0020\u0074\u0068e\u0020\u0072\u006fl\u0065\u0020\u006d\u0061p \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066 \u0074h\u0065\u0020\u0073\u0074\u0072\u0075c\u0074\u0075r\u0065\u0020\u0074\u0072e\u0065\u0020\u0072\u006f\u006ft\u002e"))
					}
				}
			}
		}
	}
	return _gefaf
}
func _ea(_ecf []*_gg.Image, _ggf bool) error {
	_gga := _cgd.PdfObjectName("\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B")
	if _ggf {
		_gga = "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
	}
	for _, _gfb := range _ecf {
		if _gfb.Colorspace == _gga {
			continue
		}
		_faea, _acd := _f.NewXObjectImageFromStream(_gfb.Stream)
		if _acd != nil {
			return _acd
		}
		_bdf, _acd := _faea.ToImage()
		if _acd != nil {
			return _acd
		}
		_gbb, _acd := _bdf.ToGoImage()
		if _acd != nil {
			return _acd
		}
		var _eac _f.PdfColorspace
		if _ggf {
			_eac = _f.NewPdfColorspaceDeviceCMYK()
			_gbb, _acd = _cf.CMYKConverter.Convert(_gbb)
		} else {
			_eac = _f.NewPdfColorspaceDeviceRGB()
			_gbb, _acd = _cf.NRGBAConverter.Convert(_gbb)
		}
		if _acd != nil {
			return _acd
		}
		_cfa, _aa := _gbb.(_cf.Image)
		if !_aa {
			return _d.New("\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074 \u0069\u006d\u0070\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002eI\u006d\u0061\u0067\u0065")
		}
		_efd := _cfa.Base()
		_gbe := &_f.Image{Width: int64(_efd.Width), Height: int64(_efd.Height), BitsPerComponent: int64(_efd.BitsPerComponent), ColorComponents: _efd.ColorComponents, Data: _efd.Data}
		_gbe.SetDecode(_efd.Decode)
		_gbe.SetAlpha(_efd.Alpha)
		if _acd = _faea.SetImage(_gbe, _eac); _acd != nil {
			return _acd
		}
		_faea.ToPdfObject()
		_gfb.ColorComponents = _efd.ColorComponents
		_gfb.Colorspace = _gga
	}
	return nil
}

type imageModifications struct {
	_bge *colorspaceModification
	_ebf _cgd.StreamEncoder
}

func _eaedd(_gaga *_f.CompliancePdfReader) (_geda []ViolatedRule) {
	var _fggf, _eafe, _fdebe bool
	if _gaga.ParserMetadata().HasNonConformantStream() {
		_geda = []ViolatedRule{_cc("\u0036.\u0031\u002e\u0037\u002d\u0032", "T\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020f\u006f\u006cl\u006fw\u0065\u0064\u0020e\u0069\u0074h\u0065\u0072\u0020\u0062\u0079\u0020\u0061 \u0043\u0041\u0052\u0052I\u0041\u0047\u0045\u0020\u0052E\u0054\u0055\u0052\u004e\u0020\u00280\u0044\u0068\u0029\u0020\u0061\u006e\u0064\u0020\u004c\u0049\u004e\u0045\u0020F\u0045\u0045\u0044\u0020\u0028\u0030\u0041\u0068\u0029\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065\u0020o\u0072\u0020\u0062\u0079\u0020\u0061 \u0073\u0069ng\u006c\u0065\u0020\u004cIN\u0045 \u0046\u0045\u0045\u0044 \u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u002e\u0020T\u0068\u0065\u0020e\u006e\u0064\u0073\u0074r\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020p\u0072\u0065\u0063\u0065\u0064\u0065\u0064\u0020\u0062\u0079\u0020\u0061n\u0020\u0045\u004f\u004c \u006d\u0061\u0072\u006b\u0065\u0072\u002e")}
	}
	for _, _adeea := range _gaga.GetObjectNums() {
		_gfedg, _ := _gaga.GetIndirectObjectByNumber(_adeea)
		if _gfedg == nil {
			continue
		}
		_bcdb, _fdfb := _cgd.GetStream(_gfedg)
		if !_fdfb {
			continue
		}
		if !_fggf {
			_gbgg := _bcdb.Get("\u004c\u0065\u006e\u0067\u0074\u0068")
			if _gbgg == nil {
				_geda = append(_geda, _cc("\u0036.\u0031\u002e\u0037\u002d\u0031", "\u006e\u006f\u0020'\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074"))
				_fggf = true
			} else {
				_cegab, _fbcf := _cgd.GetIntVal(_gbgg)
				if !_fbcf {
					_geda = append(_geda, _cc("\u0036.\u0031\u002e\u0037\u002d\u0031", "s\u0074\u0072\u0065\u0061\u006d\u0020\u0027\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020an\u0020\u0069\u006et\u0065g\u0065\u0072"))
					_fggf = true
				} else {
					if len(_bcdb.Stream) != _cegab {
						_geda = append(_geda, _cc("\u0036.\u0031\u002e\u0037\u002d\u0031", "\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006c\u0065\u006e\u0067th\u0020v\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d"))
						_fggf = true
					}
				}
			}
		}
		if !_eafe {
			if _bcdb.Get("\u0046") != nil {
				_eafe = true
				_geda = append(_geda, _cc("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
			}
			if _bcdb.Get("\u0046F\u0069\u006c\u0074\u0065\u0072") != nil && !_eafe {
				_eafe = true
				_geda = append(_geda, _cc("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
			if _bcdb.Get("\u0046\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u0061\u006d\u0073") != nil && !_eafe {
				_eafe = true
				_geda = append(_geda, _cc("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
		}
		if !_fdebe {
			_fdca, _bcfd := _cgd.GetName(_cgd.TraceToDirectObject(_bcdb.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
			if !_bcfd {
				continue
			}
			if *_fdca == _cgd.StreamEncodingFilterNameLZW {
				_fdebe = true
				_geda = append(_geda, _cc("\u0036.\u0031\u002e\u0037\u002d\u0034", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e"))
			}
		}
	}
	return _geda
}

// Profile2A is the implementation of the PDF/A-2A standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2A struct{ profile2 }

func _agdb(_ggad *_f.CompliancePdfReader) []ViolatedRule { return nil }

// ValidateStandard checks if provided input CompliancePdfReader matches rules that conforms PDF/A-1 standard.
func (_bbac *profile1) ValidateStandard(r *_f.CompliancePdfReader) error {
	_dgef := VerificationError{ConformanceLevel: _bbac._acfg._ba, ConformanceVariant: _bbac._acfg._fa}
	if _cdge := _bbba(r); _cdge != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _cdge)
	}
	if _bbfe := _dfbbg(r); _bbfe != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _bbfe)
	}
	if _eeeg := _cfaf(r); _eeeg != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _eeeg)
	}
	if _fcabc := _cggdd(r); _fcabc != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _fcabc)
	}
	if _edgb := _fggb(r); _edgb != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _edgb)
	}
	if _bcf := _egcf(r); len(_bcf) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _bcf...)
	}
	if _aaee := _cdef(r); _aaee != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _aaee)
	}
	if _eebe := _dbdf(r); len(_eebe) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _eebe...)
	}
	if _bgfdb := _gbbge(r); len(_bgfdb) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _bgfdb...)
	}
	if _gfaf := _agdb(r); len(_gfaf) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _gfaf...)
	}
	if _baag := _decad(r); _baag != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _baag)
	}
	if _dead := _dcee(r); len(_dead) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _dead...)
	}
	if _aaba := _debd(r); len(_aaba) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _aaba...)
	}
	if _edfea := _gfbff(r); _edfea != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _edfea)
	}
	if _gdeg := _caec(r, false); len(_gdeg) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _gdeg...)
	}
	if _eegb := _fgecc(r); len(_eegb) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _eegb...)
	}
	if _cbda := _adee(r); _cbda != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _cbda)
	}
	if _cabfa := _cadce(r); _cabfa != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _cabfa)
	}
	if _ggaf := _bcbc(r); _ggaf != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _ggaf)
	}
	if _eefd := _aabc(r); _eefd != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _eefd)
	}
	if _aadb := _cbaa(r); _aadb != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _aadb)
	}
	if _gdf := _eeedb(r); len(_gdf) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _gdf...)
	}
	if _ddfd := _daefe(r, _bbac._acfg); len(_ddfd) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _ddfd...)
	}
	if _edgbf := _gcfgd(r); len(_edgbf) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _edgbf...)
	}
	if _efeb := _cccbc(r); _efeb != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _efeb)
	}
	if _egfe := _gegdd(r); _egfe != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _egfe)
	}
	if _gfd := _deba(r); len(_gfd) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _gfd...)
	}
	if _cbe := _agbbe(r); len(_cbe) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _cbe...)
	}
	if _dcb := _agfdf(r); _dcb != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _dcb)
	}
	if _beeg := _ggeg(r); _beeg != _cb {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _beeg)
	}
	if _fega := _fdba(r, _bbac._acfg, false); len(_fega) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _fega...)
	}
	if _bbac._acfg == _gb() {
		if _fddaf := _caag(r); len(_fddaf) != 0 {
			_dgef.ViolatedRules = append(_dgef.ViolatedRules, _fddaf...)
		}
	}
	if _adfc := _cagd(r); len(_adfc) != 0 {
		_dgef.ViolatedRules = append(_dgef.ViolatedRules, _adfc...)
	}
	if len(_dgef.ViolatedRules) > 0 {
		_af.Slice(_dgef.ViolatedRules, func(_dagb, _edde int) bool {
			return _dgef.ViolatedRules[_dagb].RuleNo < _dgef.ViolatedRules[_edde].RuleNo
		})
		return _dgef
	}
	return nil
}
func _fecdd(_aebe string, _decf string, _cbed string) (string, bool) {
	_babb := _ac.Index(_aebe, _decf)
	if _babb == -1 {
		return "", false
	}
	_babb += len(_decf)
	_befb := _ac.Index(_aebe[_babb:], _cbed)
	if _befb == -1 {
		return "", false
	}
	_befb = _babb + _befb
	return _aebe[_babb:_befb], true
}
func _gb() standardType { return standardType{_ba: 1, _fa: "\u0041"} }

// ViolatedRule is the structure that defines violated PDF/A rule.
type ViolatedRule struct {
	RuleNo string
	Detail string
}

func _deaed(_bbaca *_cgd.PdfObjectDictionary, _dgeb map[*_cgd.PdfObjectStream][]byte, _geee map[*_cgd.PdfObjectStream]*_bd.CMap) ViolatedRule {
	const (
		_gbeb  = "\u0036.\u0033\u002e\u0038\u002d\u0031"
		_beeag = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006cl\u0020\u0069\u006e\u0063l\u0075\u0064e\u0020\u0061 \u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020w\u0068\u006f\u0073\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073 \u0061\u0020\u0043M\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u006d\u0061p\u0073\u0020\u0063\u0068\u0061\u0072ac\u0074\u0065\u0072\u0020\u0063\u006fd\u0065s\u0020\u0074\u006f\u0020\u0055\u006e\u0069\u0063\u006f\u0064e \u0076a\u006c\u0075\u0065\u0073,\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063r\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020P\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035.\u0039\u002c\u0020\u0075\u006e\u006ce\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006d\u0065\u0065\u0074\u0073 \u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0063\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u003a\u000a\u0020\u002d\u0020\u0066o\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0073\u0020M\u0061\u0063\u0052o\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074E\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041n\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u006f\u0072\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020t\u0068\u0065\u0020\u0070\u0072\u0065d\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0056\u0020C\u004d\u0061\u0070s\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0073\u0020a\u0072\u0065 \u0074\u0061k\u0065\u006e\u0020\u0066\u0072\u006f\u006d\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u0020\u0073\u0074\u0061n\u0064\u0061\u0072\u0064\u0020L\u0061t\u0069\u006e\u0020\u0063\u0068a\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0074\u0020\u006fr\u0020\u0074\u0068\u0065 \u0073\u0065\u0074\u0020\u006f\u0066 \u006e\u0061\u006d\u0065\u0064\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0066\u006f\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046 \u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0041\u0070\u0070\u0065\u006e\u0064\u0069\u0078 \u0044\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020w\u0068\u006f\u0073e\u0020d\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u0047B\u0031\u002c\u0020\u0041\u0064\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002c\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031\u0020\u006f\u0072\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006fr\u0065\u0061\u0031\u0020\u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u006c\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0073\u002e"
	)
	_begfc, _bcdgd := _cgd.GetStream(_bbaca.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e"))
	if _bcdgd {
		_, _gcgd := _acbf(_begfc, _dgeb, _geee)
		if _gcgd != nil {
			return _cc(_gbeb, _beeag)
		}
		return _cb
	}
	_aegd, _bcdgd := _cgd.GetName(_bbaca.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_bcdgd {
		return _cc(_gbeb, _beeag)
	}
	switch _aegd.String() {
	case "\u0054\u0079\u0070e\u0031":
		return _cb
	}
	return _cc(_gbeb, _beeag)
}
func _fae() standardType { return standardType{_ba: 2, _fa: "\u0041"} }

// NewProfile1A creates a new Profile1A with given options.
func NewProfile1A(options *Profile1Options) *Profile1A {
	if options == nil {
		options = DefaultProfile1Options()
	}
	_beaa(options)
	return &Profile1A{profile1{_beb: *options, _acfg: _gb()}}
}
func _gegdd(_fbde *_f.CompliancePdfReader) ViolatedRule {
	_faaf := map[*_cgd.PdfObjectStream]struct{}{}
	for _, _gcda := range _fbde.PageList {
		if _gcda.Resources == nil && _gcda.Contents == nil {
			continue
		}
		if _dfgd := _gcda.GetPageDict(); _dfgd != nil {
			_aecg, _dfde := _cgd.GetDict(_dfgd.Get("\u0047\u0072\u006fu\u0070"))
			if _dfde {
				if _adea := _aecg.Get("\u0053"); _adea != nil {
					_aeca, _efbf := _cgd.GetName(_adea)
					if _efbf && _aeca.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						return _cc("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
					}
				}
			}
		}
		if _gcda.Resources != nil {
			if _cecca, _dedb := _cgd.GetDict(_gcda.Resources.XObject); _dedb {
				for _, _adda := range _cecca.Keys() {
					_gggc, _ffff := _cgd.GetStream(_cecca.Get(_adda))
					if !_ffff {
						continue
					}
					if _, _facfd := _faaf[_gggc]; _facfd {
						continue
					}
					_bgda, _ffff := _cgd.GetDict(_gggc.Get("\u0047\u0072\u006fu\u0070"))
					if !_ffff {
						_faaf[_gggc] = struct{}{}
						continue
					}
					_dcgca := _bgda.Get("\u0053")
					if _dcgca != nil {
						_afba, _bdgbc := _cgd.GetName(_dcgca)
						if _bdgbc && _afba.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
							return _cc("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
						}
					}
					_faaf[_gggc] = struct{}{}
					continue
				}
			}
		}
		if _gcda.Contents != nil {
			_dfaga, _dadf := _gcda.GetContentStreams()
			if _dadf != nil {
				continue
			}
			for _, _fced := range _dfaga {
				_caae, _daaa := _ff.NewContentStreamParser(_fced).Parse()
				if _daaa != nil {
					continue
				}
				for _, _ggadbd := range *_caae {
					if len(_ggadbd.Params) == 0 {
						continue
					}
					_gcfe, _ffbfd := _cgd.GetName(_ggadbd.Params[0])
					if !_ffbfd {
						continue
					}
					_bddd, _cfcef := _gcda.Resources.GetXObjectByName(*_gcfe)
					if _cfcef != _f.XObjectTypeForm {
						continue
					}
					if _, _dfgbd := _faaf[_bddd]; _dfgbd {
						continue
					}
					_bcad, _ffbfd := _cgd.GetDict(_bddd.Get("\u0047\u0072\u006fu\u0070"))
					if !_ffbfd {
						_faaf[_bddd] = struct{}{}
						continue
					}
					_gcbfa := _bcad.Get("\u0053")
					if _gcbfa != nil {
						_cffe, _ebcc := _cgd.GetName(_gcbfa)
						if _ebcc && _cffe.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
							return _cc("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
						}
					}
					_faaf[_bddd] = struct{}{}
				}
			}
		}
	}
	return _cb
}
func _dbdfa(_ggec *_cgd.PdfObjectDictionary, _acab map[*_cgd.PdfObjectStream][]byte, _gcacd map[*_cgd.PdfObjectStream]*_bd.CMap) ViolatedRule {
	const (
		_fcede = "\u0046\u006f\u0072\u0020\u0061\u006e\u0079\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0063\u006f\u006d\u0070o\u0073\u0069\u0074e\u0020\u0028\u0054\u0079\u0070\u0065\u0020\u0030\u0029 \u0066\u006fn\u0074\u0020\u0077\u0069\u0074\u0068\u0069\u006e \u0061\u0020\u0063\u006fn\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f \u0065\u006e\u0074\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074\u0073 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020\u0074\u0068\u0065\u0020\u0066\u006fl\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0072\u0065l\u0061t\u0069\u006f\u006e\u0073\u0068\u0069\u0070. \u0049\u0066\u0020\u0074\u0068\u0065\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006b\u0065\u0079 \u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0054\u0079\u0070\u0065\u0020\u0030 \u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0069\u0073\u0020I\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056\u002c\u0020\u0061\u006e\u0079\u0020v\u0061\u006c\u0075\u0065\u0073\u0020\u006f\u0066\u0020\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u002c\u0020\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u002c\u0020\u0061\u006e\u0064\u0020\u0053up\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0069n\u0020\u0074h\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u0020\u0065\u006e\u0074r\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044F\u006f\u006e\u0074\u002e\u0020\u004f\u0074\u0068\u0065\u0072\u0077\u0069\u0073\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0063\u006f\u0072\u0072\u0065\u0073\u0070\u006f\u006e\u0064\u0069\u006e\u0067\u0020\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u0020a\u006e\u0064\u0020\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u0020s\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0069\u006e\u0020\u0062\u006f\u0074h\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065m\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006cl\u0020\u0062\u0065\u0020i\u0064en\u0074\u0069\u0063\u0061\u006c\u002c \u0061n\u0064\u0020\u0074\u0068\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0053\u0075\u0070\u0070l\u0065\u006d\u0065\u006e\u0074 \u006b\u0065\u0079\u0020\u0069\u006e\u0020t\u0068\u0065\u0020\u0043I\u0044S\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066o\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0067re\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f t\u0068\u0065\u0020\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066o\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043M\u0061p\u002e"
		_ddaeg = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0031"
	)
	var _cacega string
	if _cefg, _dfaeb := _cgd.GetName(_ggec.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _dfaeb {
		_cacega = _cefg.String()
	}
	if _cacega != "\u0054\u0079\u0070e\u0030" {
		return _cb
	}
	_eefaf := _ggec.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _bfca, _cacegaf := _cgd.GetName(_eefaf); _cacegaf {
		switch _bfca.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _cb
		}
		_gaaga, _dbbe := _bd.LoadPredefinedCMap(_bfca.String())
		if _dbbe != nil {
			return _cc(_ddaeg, _fcede)
		}
		_gcfc := _gaaga.CIDSystemInfo()
		if _gcfc.Ordering != _gcfc.Registry {
			return _cc(_ddaeg, _fcede)
		}
		return _cb
	}
	_dgaeb, _bcbce := _cgd.GetStream(_eefaf)
	if !_bcbce {
		return _cc(_ddaeg, _fcede)
	}
	_fabfg, _cdacb := _acbf(_dgaeb, _acab, _gcacd)
	if _cdacb != nil {
		return _cc(_ddaeg, _fcede)
	}
	_dccd := _fabfg.CIDSystemInfo()
	if _dccd.Ordering != _dccd.Registry {
		return _cc(_ddaeg, _fcede)
	}
	return _cb
}
func _bdagc(_gaac *_f.CompliancePdfReader, _gfdca standardType) (_cgabd []ViolatedRule) {
	var _abbbf, _edee, _cdab, _fabf, _caaa, _aebef, _bgdgb bool
	_adbc := func() bool { return _abbbf && _edee && _cdab && _fabf && _caaa && _aebef && _bgdgb }
	_fbba := map[*_cgd.PdfObjectStream]*_bd.CMap{}
	_dcaf := map[*_cgd.PdfObjectStream][]byte{}
	_cdgef := map[_cgd.PdfObject]*_f.PdfFont{}
	for _, _beag := range _gaac.GetObjectNums() {
		_fgce, _beffa := _gaac.GetIndirectObjectByNumber(_beag)
		if _beffa != nil {
			continue
		}
		_ddeg, _befbd := _cgd.GetDict(_fgce)
		if !_befbd {
			continue
		}
		_deaag, _befbd := _cgd.GetName(_ddeg.Get("\u0054\u0079\u0070\u0065"))
		if !_befbd {
			continue
		}
		if *_deaag != "\u0046\u006f\u006e\u0074" {
			continue
		}
		_aeaec, _beffa := _f.NewPdfFontFromPdfObject(_ddeg)
		if _beffa != nil {
			_ec.Log.Debug("g\u0065\u0074\u0074\u0069\u006e\u0067 \u0066\u006f\u006e\u0074\u0020\u0066r\u006f\u006d\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0076", _beffa)
			continue
		}
		_cdgef[_ddeg] = _aeaec
	}
	for _, _eeff := range _gaac.PageList {
		_fedgbb, _fcebd := _eeff.GetContentStreams()
		if _fcebd != nil {
			_ec.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0070\u0061\u0067\u0065\u0020\u0063o\u006e\u0074\u0065\u006e\u0074\u0020\u0073t\u0072\u0065\u0061\u006d\u0073\u0020\u0066\u0061\u0069\u006ce\u0064")
			continue
		}
		for _, _dccf := range _fedgbb {
			_deefg := _ff.NewContentStreamParser(_dccf)
			_cgbe, _feec := _deefg.Parse()
			if _feec != nil {
				_ec.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _feec)
				continue
			}
			var _fagc bool
			for _, _fgffg := range *_cgbe {
				if _fgffg.Operand != "\u0054\u0072" {
					continue
				}
				if len(_fgffg.Params) != 1 {
					_ec.Log.Debug("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054\u0072\u0027\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0065\u0064\u0020\u0027\u0031\u0027\u0020\u0062\u0075\u0074 \u0069\u0073\u003a\u0020\u0027\u0025d\u0027", len(_fgffg.Params))
					continue
				}
				_fecbg, _gcebb := _cgd.GetIntVal(_fgffg.Params[0])
				if !_gcebb {
					_ec.Log.Debug("\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020i\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
					continue
				}
				if _fecbg == 3 {
					_fagc = true
					break
				}
			}
			for _, _cceee := range *_cgbe {
				if _cceee.Operand != "\u0054\u0066" {
					continue
				}
				if len(_cceee.Params) != 2 {
					_ec.Log.Debug("i\u006eva\u006ci\u0064 \u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066 \u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072s\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054f\u0027\u0020\u006fper\u0061\u006e\u0064\u002c\u0020\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0032\u0027\u0020\u0069s\u003a \u0027\u0025\u0064\u0027", len(_cceee.Params))
					continue
				}
				_bcbca, _gcac := _cgd.GetName(_cceee.Params[0])
				if !_gcac {
					_ec.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _cceee)
					continue
				}
				_adffe, _fgabed := _eeff.Resources.GetFontByName(*_bcbca)
				if !_fgabed {
					_ec.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
					continue
				}
				_ebgc, _gcac := _cgd.GetDict(_adffe)
				if !_gcac {
					_ec.Log.Debug("\u0066\u006f\u006e\u0074 d\u0069\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
					continue
				}
				_dgdfa, _gcac := _cdgef[_ebgc]
				if !_gcac {
					var _gaef error
					_dgdfa, _gaef = _f.NewPdfFontFromPdfObject(_ebgc)
					if _gaef != nil {
						_ec.Log.Debug("\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0072o\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _gaef)
						continue
					}
					_cdgef[_ebgc] = _dgdfa
				}
				if !_abbbf {
					_dccc := _dbdfa(_ebgc, _dcaf, _fbba)
					if _dccc != _cb {
						_cgabd = append(_cgabd, _dccc)
						_abbbf = true
						if _adbc() {
							return _cgabd
						}
					}
				}
				if !_edee {
					_afdf := _dbbc(_ebgc)
					if _afdf != _cb {
						_cgabd = append(_cgabd, _afdf)
						_edee = true
						if _adbc() {
							return _cgabd
						}
					}
				}
				if !_cdab {
					_dafd := _dfagaa(_ebgc, _dcaf, _fbba)
					if _dafd != _cb {
						_cgabd = append(_cgabd, _dafd)
						_cdab = true
						if _adbc() {
							return _cgabd
						}
					}
				}
				if !_fabf {
					_gddff := _eeged(_ebgc, _dcaf, _fbba)
					if _gddff != _cb {
						_cgabd = append(_cgabd, _gddff)
						_fabf = true
						if _adbc() {
							return _cgabd
						}
					}
				}
				if !_caaa {
					_cdeae := _cffa(_dgdfa, _ebgc, _fagc)
					if _cdeae != _cb {
						_caaa = true
						_cgabd = append(_cgabd, _cdeae)
						if _adbc() {
							return _cgabd
						}
					}
				}
				if !_aebef {
					_ecgf := _fggfa(_dgdfa, _ebgc)
					if _ecgf != _cb {
						_aebef = true
						_cgabd = append(_cgabd, _ecgf)
						if _adbc() {
							return _cgabd
						}
					}
				}
				if !_bgdgb && (_gfdca._fa == "\u0041" || _gfdca._fa == "\u0055") {
					_adefd := _gdcf(_ebgc, _dcaf, _fbba)
					if _adefd != _cb {
						_bgdgb = true
						_cgabd = append(_cgabd, _adefd)
						if _adbc() {
							return _cgabd
						}
					}
				}
			}
		}
	}
	return _cgabd
}

// Part gets the PDF/A version level.
func (_ebef *profile1) Part() int { return _ebef._acfg._ba }
func _gfbb(_bebb *_cgd.PdfObjectDictionary, _accd map[*_cgd.PdfObjectStream][]byte, _efddf map[*_cgd.PdfObjectStream]*_bd.CMap) ViolatedRule {
	const (
		_bdaa = "\u0046\u006f\u0072 \u0061\u006e\u0079\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0020\u0028\u0054\u0079\u0070\u0065\u0020\u0030\u0029\u0020\u0066\u006f\u006et \u0072\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0064 \u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006fn\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0043I\u0044\u0053y\u0073\u0074\u0065\u006d\u0049nf\u006f\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u006f\u0066\u0020i\u0074\u0073\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0061\u006e\u0064 \u0043\u004d\u0061\u0070 \u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0063\u006f\u006d\u0070\u0061\u0074i\u0062\u006c\u0065\u002e\u0020\u0049\u006e\u0020o\u0074\u0068\u0065\u0072\u0020\u0077\u006f\u0072\u0064\u0073\u002c\u0020\u0074\u0068\u0065\u0020R\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u0020a\u006e\u0064\u0020\u004fr\u0064\u0065\u0072\u0069\u006e\u0067 \u0073\u0074\u0072i\u006e\u0067\u0073\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0066\u006f\u0072 \u0074\u0068\u0061\u0074\u0020\u0066o\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u002c\u0020u\u006el\u0065ss \u0074\u0068\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006eg\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074h\u0065 \u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073 \u0049\u0064\u0065\u006e\u0074\u0069t\u0079\u002d\u0048\u0020o\u0072\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074y\u002dV\u002e"
		_fbaf = "\u0036.\u0033\u002e\u0033\u002d\u0031"
	)
	var _dgffb string
	if _ceeb, _acff := _cgd.GetName(_bebb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _acff {
		_dgffb = _ceeb.String()
	}
	if _dgffb != "\u0054\u0079\u0070e\u0030" {
		return _cb
	}
	_dagc := _bebb.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _bcgd, _acfgg := _cgd.GetName(_dagc); _acfgg {
		switch _bcgd.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _cb
		}
		_fadc, _cedb := _bd.LoadPredefinedCMap(_bcgd.String())
		if _cedb != nil {
			return _cc(_fbaf, _bdaa)
		}
		_ebgf := _fadc.CIDSystemInfo()
		if _ebgf.Ordering != _ebgf.Registry {
			return _cc(_fbaf, _bdaa)
		}
		return _cb
	}
	_gffg, _fdcb := _cgd.GetStream(_dagc)
	if !_fdcb {
		return _cc(_fbaf, _bdaa)
	}
	_ggfb, _gcga := _acbf(_gffg, _accd, _efddf)
	if _gcga != nil {
		return _cc(_fbaf, _bdaa)
	}
	_ccdcg := _ggfb.CIDSystemInfo()
	if _ccdcg.Ordering != _ccdcg.Registry {
		return _cc(_fbaf, _bdaa)
	}
	return _cb
}
func _acb(_agc standardType, _caebe *_gg.OutputIntents) error {
	_cace, _dgcge := _dfa.NewCmykIsoCoatedV2OutputIntent(_agc.outputIntentSubtype())
	if _dgcge != nil {
		return _dgcge
	}
	if _dgcge = _caebe.Add(_cace.ToPdfObject()); _dgcge != nil {
		return _dgcge
	}
	return nil
}
func _cagd(_dbgcf *_f.CompliancePdfReader) (_fgbaa []ViolatedRule) {
	for _, _dabc := range _dbgcf.GetObjectNums() {
		_ggcd, _fdagg := _dbgcf.GetIndirectObjectByNumber(_dabc)
		if _fdagg != nil {
			continue
		}
		_eggdd, _bbde := _cgd.GetDict(_ggcd)
		if !_bbde {
			continue
		}
		_caba, _bbde := _cgd.GetName(_eggdd.Get("\u0054\u0079\u0070\u0065"))
		if !_bbde {
			continue
		}
		if _caba.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_faedg, _bbde := _cgd.GetBool(_eggdd.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if !_bbde {
			return _fgbaa
		}
		if bool(*_faedg) {
			_fgbaa = append(_fgbaa, _cc("\u0036\u002e\u0039-\u0031", "\u0054\u0068\u0065\u0020\u004e\u0065e\u0064\u0041\u0070\u0070\u0065a\u0072\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006c\u0061\u0067\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0069\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076e\u0020\u0066\u006f\u0072\u006d \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u006e\u006f\u0074\u0020b\u0065\u0020\u0070\u0072\u0065se\u006e\u0074\u0020\u006f\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
		}
	}
	return _fgbaa
}

type documentImages struct {
	_fde, _bg, _bgd bool
	_bff            map[_cgd.PdfObject]struct{}
	_fe             []*imageInfo
}

// ApplyStandard tries to change the content of the writer to match the PDF/A-2 standard.
// Implements model.StandardApplier.
func (_fafg *profile2) ApplyStandard(document *_gg.Document) (_dbfc error) {
	_fff(document, 7)
	if _dbfc = _bbe(document, _fafg._ddb.Now); _dbfc != nil {
		return _dbfc
	}
	if _dbfc = _cgbf(document); _dbfc != nil {
		return _dbfc
	}
	_faag, _dgfc := _eag(_fafg._ddb.CMYKDefaultColorSpace, _fafg._geg)
	_dbfc = _ecd(document, []pageColorspaceOptimizeFunc{_faag}, []documentColorspaceOptimizeFunc{_dgfc})
	if _dbfc != nil {
		return _dbfc
	}
	_gde(document)
	if _dbfc = _bege(document); _dbfc != nil {
		return _dbfc
	}
	if _dbfc = _dac(document, _fafg._geg._ba); _dbfc != nil {
		return _dbfc
	}
	if _dbfc = _febg(document); _dbfc != nil {
		return _dbfc
	}
	if _dbfc = _fgaa(document); _dbfc != nil {
		return _dbfc
	}
	if _dbfc = _bdd(document); _dbfc != nil {
		return _dbfc
	}
	if _dbfc = _cabd(document); _dbfc != nil {
		return _dbfc
	}
	if _fafg._geg._fa == "\u0041" {
		_cfeg(document)
	}
	if _dbfc = _bde(document, _fafg._geg._ba); _dbfc != nil {
		return _dbfc
	}
	if _dbfc = _gfbg(document); _dbfc != nil {
		return _dbfc
	}
	if _edec := _fede(document, _fafg._geg, _fafg._ddb.Xmp); _edec != nil {
		return _edec
	}
	if _fafg._geg == _fae() {
		if _dbfc = _age(document); _dbfc != nil {
			return _dbfc
		}
	}
	if _dbfc = _cgef(document); _dbfc != nil {
		return _dbfc
	}
	if _dbfc = _aaebe(document); _dbfc != nil {
		return _dbfc
	}
	if _dbfc = _gcad(document); _dbfc != nil {
		return _dbfc
	}
	return nil
}
func _dgca(_gbgc *_f.PdfInfo, _fafgg *_cg.Document) bool {
	_bebba, _agdf := _fafgg.GetPdfInfo()
	if !_agdf {
		return false
	}
	if _bebba.InfoDict == nil {
		return false
	}
	_cfdfb, _gecfb := _f.NewPdfInfoFromObject(_bebba.InfoDict)
	if _gecfb != nil {
		return false
	}
	if _gbgc.Creator != nil {
		if _cfdfb.Creator == nil || _cfdfb.Creator.String() != _gbgc.Creator.String() {
			return false
		}
	}
	if _gbgc.CreationDate != nil {
		if _cfdfb.CreationDate == nil || !_cfdfb.CreationDate.ToGoTime().Equal(_gbgc.CreationDate.ToGoTime()) {
			return false
		}
	}
	if _gbgc.ModifiedDate != nil {
		if _cfdfb.ModifiedDate == nil || !_cfdfb.ModifiedDate.ToGoTime().Equal(_gbgc.ModifiedDate.ToGoTime()) {
			return false
		}
	}
	if _gbgc.Producer != nil {
		if _cfdfb.Producer == nil || _cfdfb.Producer.String() != _gbgc.Producer.String() {
			return false
		}
	}
	if _gbgc.Keywords != nil {
		if _cfdfb.Keywords == nil || _cfdfb.Keywords.String() != _gbgc.Keywords.String() {
			return false
		}
	}
	if _gbgc.Trapped != nil {
		if _cfdfb.Trapped == nil {
			return false
		}
		switch _gbgc.Trapped.String() {
		case "\u0054\u0072\u0075\u0065":
			if _cfdfb.Trapped.String() != "\u0054\u0072\u0075\u0065" {
				return false
			}
		case "\u0046\u0061\u006cs\u0065":
			if _cfdfb.Trapped.String() != "\u0046\u0061\u006cs\u0065" {
				return false
			}
		default:
			if _cfdfb.Trapped.String() != "\u0046\u0061\u006cs\u0065" {
				return false
			}
		}
	}
	if _gbgc.Title != nil {
		if _cfdfb.Title == nil || _cfdfb.Title.String() != _gbgc.Title.String() {
			return false
		}
	}
	if _gbgc.Subject != nil {
		if _cfdfb.Subject == nil || _cfdfb.Subject.String() != _gbgc.Subject.String() {
			return false
		}
	}
	return true
}
func _bdce(_gfadda *_f.CompliancePdfReader) (_deab []ViolatedRule) {
	var _baadc, _bebf bool
	_ffae := func() bool { return _baadc && _bebf }
	for _, _gfeg := range _gfadda.GetObjectNums() {
		_dfdbg, _caeg := _gfadda.GetIndirectObjectByNumber(_gfeg)
		if _caeg != nil {
			_ec.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0025\u0064\u0020fa\u0069\u006c\u0065d\u003a \u0025\u0076", _gfeg, _caeg)
			continue
		}
		_bgcc, _eeab := _cgd.GetDict(_dfdbg)
		if !_eeab {
			continue
		}
		_eggg, _eeab := _cgd.GetName(_bgcc.Get("\u0054\u0079\u0070\u0065"))
		if !_eeab {
			continue
		}
		if *_eggg != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_ggaedg, _eeab := _cgd.GetName(_bgcc.Get("\u0053"))
		if !_eeab {
			if !_baadc {
				_deab = append(_deab, _cc("\u0036.\u0035\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004caun\u0063\u0068\u002c\u0020S\u006f\u0075\u006e\u0064,\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046\u006f\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044a\u0074\u0061,\u0020\u0048\u0069\u0064\u0065\u002c\u0020\u0053\u0065\u0074\u004f\u0043\u0047\u0053\u0074\u0061\u0074\u0065\u002c\u0020\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u002c\u0020T\u0072\u0061\u006e\u0073\u002c\u0020\u0047o\u0054\u006f\u0033\u0044\u0056\u0069\u0065\u0077\u0020\u0061\u006e\u0064\u0020\u004a\u0061v\u0061Sc\u0072\u0069p\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020\u0070\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e \u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020t\u0068\u0065\u0020\u0064\u0065\u0070\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020\u0073\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u006f\u0070\u0020\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070e\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_baadc = true
				if _ffae() {
					return _deab
				}
			}
			continue
		}
		switch _f.PdfActionType(*_ggaedg) {
		case _f.ActionTypeLaunch, _f.ActionTypeSound, _f.ActionTypeMovie, _f.ActionTypeResetForm, _f.ActionTypeImportData, _f.ActionTypeJavaScript, _f.ActionTypeHide, _f.ActionTypeSetOCGState, _f.ActionTypeRendition, _f.ActionTypeTrans, _f.ActionTypeGoTo3DView:
			if !_baadc {
				_deab = append(_deab, _cc("\u0036.\u0035\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004caun\u0063\u0068\u002c\u0020S\u006f\u0075\u006e\u0064,\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046\u006f\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044a\u0074\u0061,\u0020\u0048\u0069\u0064\u0065\u002c\u0020\u0053\u0065\u0074\u004f\u0043\u0047\u0053\u0074\u0061\u0074\u0065\u002c\u0020\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u002c\u0020T\u0072\u0061\u006e\u0073\u002c\u0020\u0047o\u0054\u006f\u0033\u0044\u0056\u0069\u0065\u0077\u0020\u0061\u006e\u0064\u0020\u004a\u0061v\u0061Sc\u0072\u0069p\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020\u0070\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e \u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020t\u0068\u0065\u0020\u0064\u0065\u0070\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020\u0073\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u006f\u0070\u0020\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070e\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_baadc = true
				if _ffae() {
					return _deab
				}
			}
			continue
		case _f.ActionTypeNamed:
			if !_bebf {
				_fefb, _bacdg := _cgd.GetName(_bgcc.Get("\u004e"))
				if !_bacdg {
					_deab = append(_deab, _cc("\u0036.\u0035\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_bebf = true
					if _ffae() {
						return _deab
					}
					continue
				}
				switch *_fefb {
				case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
				default:
					_deab = append(_deab, _cc("\u0036.\u0035\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_bebf = true
					if _ffae() {
						return _deab
					}
					continue
				}
			}
		}
	}
	return _deab
}

var _ Profile = (*Profile2A)(nil)

func _fcaf(_decc *_gg.Document, _fgd standardType, _ccdfg *_gg.OutputIntents) error {
	var (
		_egf  *_f.PdfOutputIntent
		_bgac error
	)
	if _decc.Version.Minor <= 7 {
		_egf, _bgac = _dfa.NewSRGBv2OutputIntent(_fgd.outputIntentSubtype())
	} else {
		_egf, _bgac = _dfa.NewSRGBv4OutputIntent(_fgd.outputIntentSubtype())
	}
	if _bgac != nil {
		return _bgac
	}
	if _bgac = _ccdfg.Add(_egf.ToPdfObject()); _bgac != nil {
		return _bgac
	}
	return nil
}
func _agbbe(_dfae *_f.CompliancePdfReader) (_cgcg []ViolatedRule) {
	var _ceadg, _bcac bool
	_cgcc := func() bool { return _ceadg && _bcac }
	for _, _caff := range _dfae.GetObjectNums() {
		_dbda, _fbg := _dfae.GetIndirectObjectByNumber(_caff)
		if _fbg != nil {
			_ec.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0025\u0064\u0020fa\u0069\u006c\u0065d\u003a \u0025\u0076", _caff, _fbg)
			continue
		}
		_acagf, _abadb := _cgd.GetDict(_dbda)
		if !_abadb {
			continue
		}
		_ceefe, _abadb := _cgd.GetName(_acagf.Get("\u0054\u0079\u0070\u0065"))
		if !_abadb {
			continue
		}
		if *_ceefe != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_bcba, _abadb := _cgd.GetName(_acagf.Get("\u0053"))
		if !_abadb {
			if !_ceadg {
				_cgcg = append(_cgcg, _cc("\u0036.\u0036\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004c\u0061\u0075\u006e\u0063\u0068\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046o\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u006e\u0064 \u004a\u0061\u0076a\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020s\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e \u0041\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020th\u0065\u0020\u0064\u0065p\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020s\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u002d\u006f\u0070\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062e\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e\u0020T\u0068\u0065\u0020\u0048\u0069\u0064\u0065\u0020a\u0063\u0074\u0069\u006f\u006e \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_ceadg = true
				if _cgcc() {
					return _cgcg
				}
			}
			continue
		}
		switch _f.PdfActionType(*_bcba) {
		case _f.ActionTypeLaunch, _f.ActionTypeSound, _f.ActionTypeMovie, _f.ActionTypeResetForm, _f.ActionTypeImportData, _f.ActionTypeJavaScript:
			if !_ceadg {
				_cgcg = append(_cgcg, _cc("\u0036.\u0036\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004c\u0061\u0075\u006e\u0063\u0068\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046o\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u006e\u0064 \u004a\u0061\u0076a\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020s\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e \u0041\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020th\u0065\u0020\u0064\u0065p\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020s\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u002d\u006f\u0070\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062e\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e\u0020T\u0068\u0065\u0020\u0048\u0069\u0064\u0065\u0020a\u0063\u0074\u0069\u006f\u006e \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_ceadg = true
				if _cgcc() {
					return _cgcg
				}
			}
			continue
		case _f.ActionTypeNamed:
			if !_bcac {
				_fcdg, _fcfb := _cgd.GetName(_acagf.Get("\u004e"))
				if !_fcfb {
					_cgcg = append(_cgcg, _cc("\u0036.\u0036\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_bcac = true
					if _cgcc() {
						return _cgcg
					}
					continue
				}
				switch *_fcdg {
				case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
				default:
					_cgcg = append(_cgcg, _cc("\u0036.\u0036\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_bcac = true
					if _cgcc() {
						return _cgcg
					}
					continue
				}
			}
		}
	}
	return _cgcg
}
func _eab(_bdda *_gg.Document, _ccce bool) error {
	_fdad, _eebb := _bdda.GetPages()
	if !_eebb {
		return nil
	}
	for _, _bad := range _fdad {
		_faca, _bgc := _cgd.GetArray(_bad.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_bgc {
			continue
		}
		for _, _bdeb := range _faca.Elements() {
			_ddfg, _bece := _cgd.GetDict(_bdeb)
			if !_bece {
				continue
			}
			_egg := _ddfg.Get("\u0043")
			if _egg == nil {
				continue
			}
			_eggf, _bece := _cgd.GetArray(_egg)
			if !_bece {
				continue
			}
			_dffd, _bcaac := _eggf.GetAsFloat64Slice()
			if _bcaac != nil {
				return _bcaac
			}
			switch _eggf.Len() {
			case 0, 1:
				if _ccce {
					_ddfg.Set("\u0043", _cgd.MakeArrayFromIntegers([]int{1, 1, 1, 1}))
				} else {
					_ddfg.Set("\u0043", _cgd.MakeArrayFromIntegers([]int{1, 1, 1}))
				}
			case 3:
				if _ccce {
					_adbb, _fffc, _gba, _cdg := _c.RGBToCMYK(uint8(_dffd[0]*255), uint8(_dffd[1]*255), uint8(_dffd[2]*255))
					_ddfg.Set("\u0043", _cgd.MakeArrayFromFloats([]float64{float64(_adbb) / 255, float64(_fffc) / 255, float64(_gba) / 255, float64(_cdg) / 255}))
				}
			case 4:
				if !_ccce {
					_cgg, _dbb, _fadaf := _c.CMYKToRGB(uint8(_dffd[0]*255), uint8(_dffd[1]*255), uint8(_dffd[2]*255), uint8(_dffd[3]*255))
					_ddfg.Set("\u0043", _cgd.MakeArrayFromFloats([]float64{float64(_cgg) / 255, float64(_dbb) / 255, float64(_fadaf) / 255}))
				}
			}
		}
	}
	return nil
}
func _gdcf(_ceabf *_cgd.PdfObjectDictionary, _aadd map[*_cgd.PdfObjectStream][]byte, _cdbed map[*_cgd.PdfObjectStream]*_bd.CMap) ViolatedRule {
	const (
		_aebg = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0037\u002d\u0031"
		_eada = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006cl\u0020\u0069\u006e\u0063l\u0075\u0064e\u0020\u0061 \u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020w\u0068\u006f\u0073\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073 \u0061\u0020\u0043M\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u006d\u0061p\u0073\u0020\u0063\u0068\u0061\u0072ac\u0074\u0065\u0072\u0020\u0063\u006fd\u0065s\u0020\u0074\u006f\u0020\u0055\u006e\u0069\u0063\u006f\u0064e \u0076a\u006c\u0075\u0065\u0073,\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063r\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020P\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035.\u0039\u002c\u0020\u0075\u006e\u006ce\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006d\u0065\u0065\u0074\u0073 \u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0063\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u003a\u000a\u0020\u002d\u0020\u0066o\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0073\u0020M\u0061\u0063\u0052o\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074E\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041n\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u006f\u0072\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020t\u0068\u0065\u0020\u0070\u0072\u0065d\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0056\u0020C\u004d\u0061\u0070s\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0073\u0020a\u0072\u0065 \u0074\u0061k\u0065\u006e\u0020\u0066\u0072\u006f\u006d\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u0020\u0073\u0074\u0061n\u0064\u0061\u0072\u0064\u0020L\u0061t\u0069\u006e\u0020\u0063\u0068a\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0074\u0020\u006fr\u0020\u0074\u0068\u0065 \u0073\u0065\u0074\u0020\u006f\u0066 \u006e\u0061\u006d\u0065\u0064\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0066\u006f\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046 \u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0041\u0070\u0070\u0065\u006e\u0064\u0069\u0078 \u0044\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020w\u0068\u006f\u0073e\u0020d\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u0047B\u0031\u002c\u0020\u0041\u0064\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002c\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031\u0020\u006f\u0072\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006fr\u0065\u0061\u0031\u0020\u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u006c\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0073\u002e"
	)
	_fcafd, _afadc := _cgd.GetStream(_ceabf.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e"))
	if _afadc {
		_, _cdgea := _acbf(_fcafd, _aadd, _cdbed)
		if _cdgea != nil {
			return _cc(_aebg, _eada)
		}
		return _cb
	}
	_gfcb, _afadc := _cgd.GetName(_ceabf.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_afadc {
		return _cc(_aebg, _eada)
	}
	switch _gfcb.String() {
	case "\u0054\u0079\u0070e\u0031":
		return _cb
	}
	return _cc(_aebg, _eada)
}
func _dfagaa(_gcba *_cgd.PdfObjectDictionary, _bafg map[*_cgd.PdfObjectStream][]byte, _egcd map[*_cgd.PdfObjectStream]*_bd.CMap) ViolatedRule {
	const (
		_bdfaa = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0033"
		_afcef = "\u0041\u006c\u006c \u0043\u004d\u0061\u0070s\u0020\u0075\u0073ed\u0020\u0077\u0069\u0074\u0068i\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074 th\u006f\u0073\u0065\u0020\u006ci\u0073\u0074\u0065\u0064\u0020i\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u00320\u00300\u002d1\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0037\u002e\u0035\u002e\u0032\u002c\u0020\u0054\u0061\u0062\u006c\u0065 \u0031\u00318,\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0069\u006e \u0074\u0068\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u0020\u0061\u0073\u0020\u0064e\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u0032\u00300\u0030-\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u00209\u002e\u0037\u002e\u0035\u002e"
	)
	var _dbed string
	if _dgac, _fgfab := _cgd.GetName(_gcba.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _fgfab {
		_dbed = _dgac.String()
	}
	if _dbed != "\u0054\u0079\u0070e\u0030" {
		return _cb
	}
	_gebg := _gcba.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _eabff, _cdfea := _cgd.GetName(_gebg); _cdfea {
		switch _eabff.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _cb
		default:
			return _cc(_bdfaa, _afcef)
		}
	}
	_gdafb, _dgfgf := _cgd.GetStream(_gebg)
	if !_dgfgf {
		return _cc(_bdfaa, _afcef)
	}
	_, _ebbfa := _acbf(_gdafb, _bafg, _egcd)
	if _ebbfa != nil {
		return _cc(_bdfaa, _afcef)
	}
	return _cb
}
func _decad(_aagf *_f.CompliancePdfReader) ViolatedRule {
	for _, _cccb := range _aagf.PageList {
		_bbbe := _cccb.GetContentStreamObjs()
		for _, _aggb := range _bbbe {
			_aggb = _cgd.TraceToDirectObject(_aggb)
			var _dcbc string
			switch _efcd := _aggb.(type) {
			case *_cgd.PdfObjectString:
				_dcbc = _efcd.Str()
			case *_cgd.PdfObjectStream:
				_eeaga, _bcbg := _cgd.GetName(_cgd.TraceToDirectObject(_efcd.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
				if _bcbg {
					if *_eeaga == _cgd.StreamEncodingFilterNameLZW {
						return _cc("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0032", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e")
					}
				}
				_bbdb, _cgab := _cgd.DecodeStream(_efcd)
				if _cgab != nil {
					_ec.Log.Debug("\u0045r\u0072\u003a\u0020\u0025\u0076", _cgab)
					continue
				}
				_dcbc = string(_bbdb)
			default:
				_ec.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063t\u003a\u0020\u0025\u0054", _aggb)
				continue
			}
			_eadc := _ff.NewContentStreamParser(_dcbc)
			_bed, _faeg := _eadc.Parse()
			if _faeg != nil {
				_ec.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d:\u0020\u0025\u0076", _faeg)
				continue
			}
			for _, _dbfe := range *_bed {
				if !(_dbfe.Operand == "\u0042\u0049" && len(_dbfe.Params) == 1) {
					continue
				}
				_fcfe, _aega := _dbfe.Params[0].(*_ff.ContentStreamInlineImage)
				if !_aega {
					continue
				}
				_fged, _fbfb := _fcfe.GetEncoder()
				if _fbfb != nil {
					_ec.Log.Debug("\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006e\u006c\u0069\u006ee\u0020\u0069\u006d\u0061\u0067\u0065 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _fbfb)
					continue
				}
				if _fged.GetFilterName() == _cgd.StreamEncodingFilterNameLZW {
					return _cc("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0032", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e")
				}
			}
		}
	}
	return _cb
}
func _dcee(_cfgf *_f.CompliancePdfReader) (_effg []ViolatedRule) {
	_cdded := _cfgf.GetObjectNums()
	for _, _cggdg := range _cdded {
		_gbfd, _aeae := _cfgf.GetIndirectObjectByNumber(_cggdg)
		if _aeae != nil {
			continue
		}
		_gade, _efbc := _cgd.GetDict(_gbfd)
		if !_efbc {
			continue
		}
		_gdfc, _efbc := _cgd.GetName(_gade.Get("\u0054\u0079\u0070\u0065"))
		if !_efbc {
			continue
		}
		if _gdfc.String() != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
			continue
		}
		if _gade.Get("\u0045\u0046") != nil {
			_effg = append(_effg, _cc("\u0036\u002e\u0031\u002e\u0031\u0031\u002d\u0031", "\u0041 \u0066\u0069\u006c\u0065 \u0073p\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066i\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046 \u0033\u002e\u0031\u0030\u002e\u0032\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0045\u0046 \u006be\u0079\u002e"))
			break
		}
	}
	_accag, _bfaec := _ecgd(_cfgf)
	if !_bfaec {
		return _effg
	}
	_gcf, _bfaec := _cgd.GetDict(_accag.Get("\u004e\u0061\u006de\u0073"))
	if !_bfaec {
		return _effg
	}
	if _gcf.Get("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0046\u0069\u006c\u0065\u0073") != nil {
		_effg = append(_effg, _cc("\u0036\u002e\u0031\u002e\u0031\u0031\u002d\u0032", "\u0041\u0020\u0066i\u006c\u0065\u0027\u0073\u0020\u006e\u0061\u006d\u0065\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020d\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020PD\u0046 \u0052\u0065\u0066er\u0065\u006e\u0063\u0065\u0020\u0033\u002e6\u002e\u0033\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0045m\u0062\u0065\u0064\u0064\u0065\u0064\u0046i\u006c\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
	}
	return _effg
}
func _agfdf(_abge *_f.CompliancePdfReader) (_bcef ViolatedRule) {
	_cded, _dggd := _ecgd(_abge)
	if !_dggd {
		return _cb
	}
	_bgcf, _dggd := _cgd.GetDict(_cded.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
	if !_dggd {
		return _cb
	}
	_gafcd, _dggd := _cgd.GetArray(_bgcf.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_dggd {
		return _cb
	}
	for _egddb := 0; _egddb < _gafcd.Len(); _egddb++ {
		_ccfg, _baee := _cgd.GetDict(_gafcd.Get(_egddb))
		if !_baee {
			continue
		}
		if _ccfg.Get("\u0041\u0041") != nil {
			return _cc("\u0036.\u0036\u002e\u0032\u002d\u0032", "\u0041\u0020F\u0069\u0065\u006cd\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079 s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061n\u0020A\u0041\u0020\u0065\u006e\u0074\u0072y f\u006f\u0072\u0020\u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069on\u0061l\u002d\u0061\u0063\u0074i\u006fn\u0073 \u0064\u0069c\u0074\u0069on\u0061\u0072\u0079\u002e")
		}
	}
	return _cb
}
func (_fd standardType) outputIntentSubtype() _f.PdfOutputIntentType {
	switch _fd._ba {
	case 1:
		return _f.PdfOutputIntentTypeA1
	case 2:
		return _f.PdfOutputIntentTypeA2
	case 3:
		return _f.PdfOutputIntentTypeA3
	case 4:
		return _f.PdfOutputIntentTypeA4
	default:
		return 0
	}
}
func _fggb(_dfg *_f.CompliancePdfReader) ViolatedRule { return _cb }
func _bde(_eeeb *_gg.Document, _bca int) error {
	for _, _cfbe := range _eeeb.Objects {
		_eae, _dfaf := _cgd.GetDict(_cfbe)
		if !_dfaf {
			continue
		}
		_cfcf := _eae.Get("\u0054\u0079\u0070\u0065")
		if _cfcf == nil {
			continue
		}
		if _fag, _dbd := _cgd.GetName(_cfcf); _dbd && _fag.String() != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_dec, _fgfa := _cgd.GetName(_eae.Get("\u0053"))
		if !_fgfa {
			continue
		}
		switch _f.PdfActionType(*_dec) {
		case _f.ActionTypeLaunch, _f.ActionTypeSound, _f.ActionTypeMovie, _f.ActionTypeResetForm, _f.ActionTypeImportData, _f.ActionTypeJavaScript:
			_eae.Remove("\u0053")
		case _f.ActionTypeHide, _f.ActionTypeSetOCGState, _f.ActionTypeRendition, _f.ActionTypeTrans, _f.ActionTypeGoTo3DView:
			if _bca == 2 {
				_eae.Remove("\u0053")
			}
		case _f.ActionTypeNamed:
			_fgff, _fba := _cgd.GetName(_eae.Get("\u004e"))
			if !_fba {
				continue
			}
			switch *_fgff {
			case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
			default:
				_eae.Remove("\u004e")
			}
		}
	}
	return nil
}

// DefaultProfile2Options are the default options for the Profile2.
func DefaultProfile2Options() *Profile2Options {
	return &Profile2Options{Now: _eg.Now, Xmp: XmpOptions{MarshalIndent: "\u0009"}}
}
func _cccbc(_bafe *_f.CompliancePdfReader) ViolatedRule {
	for _, _ffgdf := range _bafe.GetObjectNums() {
		_ggc, _gdbc := _bafe.GetIndirectObjectByNumber(_ffgdf)
		if _gdbc != nil {
			continue
		}
		_agff, _fece := _cgd.GetStream(_ggc)
		if !_fece {
			continue
		}
		_badc, _fece := _cgd.GetName(_agff.Get("\u0054\u0079\u0070\u0065"))
		if !_fece {
			continue
		}
		if *_badc != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		if _agff.Get("\u0053\u004d\u0061s\u006b") != nil {
			return _cc("\u0036\u002e\u0034-\u0032", "\u0041\u006e\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068e \u0053\u004d\u0061\u0073\u006b\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _cb
}
func _debd(_eddb *_f.CompliancePdfReader) (_cgabc []ViolatedRule) {
	var (
		_caceg, _afdd, _bcge, _gfddb, _cebe, _afad, _afae bool
		_ddfe                                             func(_cgd.PdfObject)
	)
	_ddfe = func(_ccee _cgd.PdfObject) {
		switch _affc := _ccee.(type) {
		case *_cgd.PdfObjectInteger:
			if !_caceg && (int64(*_affc) > _e.MaxInt32 || int64(*_affc) < -_e.MaxInt32) {
				_cgabc = append(_cgabc, _cc("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0031", "L\u0061\u0072\u0067e\u0073\u0074\u0020\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0032\u002c\u0031\u0034\u0037,\u0034\u0038\u0033,\u0036\u0034\u0037\u002e\u0020\u0053\u006d\u0061\u006c\u006c\u0065\u0073\u0074 \u0069\u006e\u0074\u0065g\u0065\u0072\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u002d\u0032\u002c\u0031\u0034\u0037\u002c\u0034\u0038\u0033,\u0036\u0034\u0038\u002e"))
				_caceg = true
			}
		case *_cgd.PdfObjectFloat:
			if !_afdd && (_e.Abs(float64(*_affc)) > 32767.0) {
				_cgabc = append(_cgabc, _cc("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0032", "\u0041\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0020\u0072\u0065\u0061\u006c\u0020\u0076\u0061\u006c\u0075\u0065\u0020m\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u006c\u0065s\u0073\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075a\u006c\u0020\u0074\u006f\u0020\u00332\u0037\u0036\u0037.\u0030\u002e"))
			}
		case *_cgd.PdfObjectString:
			if !_bcge && len([]byte(_affc.Str())) > 65535 {
				_cgabc = append(_cgabc, _cc("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0033", "M\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006c\u0065n\u0067\u0074\u0068\u0020\u006f\u0066\u0020a \u0073\u0074\u0072\u0069n\u0067\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074es\u0029\u0020i\u0073\u0020\u0036\u0035\u0035\u0033\u0035\u002e"))
				_bcge = true
			}
		case *_cgd.PdfObjectName:
			if !_gfddb && len([]byte(*_affc)) > 127 {
				_cgabc = append(_cgabc, _cc("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0034", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006c\u0065\u006eg\u0074\u0068\u0020\u006ff\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0029\u0020\u0069\u0073\u0020\u0031\u0032\u0037\u002e"))
				_gfddb = true
			}
		case *_cgd.PdfObjectArray:
			if !_cebe && _affc.Len() > 8191 {
				_cgabc = append(_cgabc, _cc("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0035", "\u004d\u0061\u0078\u0069\u006d\u0075m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006f\u0066\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020(\u0069\u006e\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0029\u0020\u0069s\u00208\u0031\u0039\u0031\u002e"))
				_cebe = true
			}
			for _, _geac := range _affc.Elements() {
				_ddfe(_geac)
			}
			if !_afae && (_affc.Len() == 4 || _affc.Len() == 5) {
				_agdd, _gdbf := _cgd.GetName(_affc.Get(0))
				if !_gdbf {
					return
				}
				if *_agdd != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
					return
				}
				_dgb := _affc.Get(1)
				_dgb = _cgd.TraceToDirectObject(_dgb)
				_ffab, _gdbf := _cgd.GetArray(_dgb)
				if !_gdbf {
					return
				}
				if _ffab.Len() > 8 {
					_cgabc = append(_cgabc, _cc("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0039", "\u004d\u0061\u0078i\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0069\u0073\u0020\u0038\u002e"))
					_afae = true
				}
			}
		case *_cgd.PdfObjectDictionary:
			_eafg := _affc.Keys()
			if !_afad && len(_eafg) > 4095 {
				_cgabc = append(_cgabc, _cc("\u0036.\u0031\u002e\u0031\u0032\u002d\u00311", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u0063\u0061\u0070\u0061\u0063\u0069\u0074y\u0020\u006f\u0066\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0028\u0069\u006e\u0020\u0065\u006e\u0074\u0072\u0069es\u0029\u0020\u0069\u0073\u0020\u0034\u0030\u0039\u0035\u002e"))
				_afad = true
			}
			for _efeba, _cceef := range _eafg {
				_ddfe(&_eafg[_efeba])
				_ddfe(_affc.Get(_cceef))
			}
		case *_cgd.PdfObjectStream:
			_ddfe(_affc.PdfObjectDictionary)
		case *_cgd.PdfObjectStreams:
			for _, _bgce := range _affc.Elements() {
				_ddfe(_bgce)
			}
		case *_cgd.PdfObjectReference:
			_ddfe(_affc.Resolve())
		}
	}
	_eead := _eddb.GetObjectNums()
	if len(_eead) > 8388607 {
		_cgabc = append(_cgabc, _cc("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0037", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020in\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073 \u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0069\u0073\u00208\u002c\u0033\u0038\u0038\u002c\u0036\u0030\u0037\u002e"))
	}
	for _, _agfd := range _eead {
		_aeag, _dcca := _eddb.GetIndirectObjectByNumber(_agfd)
		if _dcca != nil {
			continue
		}
		_gcebg := _cgd.TraceToDirectObject(_aeag)
		_ddfe(_gcebg)
	}
	return _cgabc
}
func _ggea(_aecf *_f.PdfFont, _acfd *_cgd.PdfObjectDictionary) ViolatedRule {
	const (
		_abbe = "\u0036.\u0033\u002e\u0037\u002d\u0032"
		_ccaa = "\u0041l\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072u\u0065\u0054\u0079p\u0065\u0020\u0066\u006f\u006e\u0074s\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065n\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068e\u0020\u0066\u006f\u006e\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"
	)
	var _fcff string
	if _bcfa, _fbfe := _cgd.GetName(_acfd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _fbfe {
		_fcff = _bcfa.String()
	}
	if _fcff != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _cb
	}
	_cead := _aecf.FontDescriptor()
	_dcefe, _ccgcf := _cgd.GetIntVal(_cead.Flags)
	if !_ccgcf {
		_ec.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _cc(_abbe, _ccaa)
	}
	_ecca := (uint32(_dcefe) >> 3) & 1
	_dbeb := _ecca != 0
	if !_dbeb {
		return _cb
	}
	if _acfd.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067") != nil {
		return _cc(_abbe, _ccaa)
	}
	return _cb
}
func _bbaba(_gfcf *_f.CompliancePdfReader) (*_cgd.PdfObjectDictionary, bool) {
	_cfbb, _afbg := _ecgd(_gfcf)
	if !_afbg {
		return nil, false
	}
	_aggab, _afbg := _cgd.GetArray(_cfbb.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_afbg {
		return nil, false
	}
	if _aggab.Len() == 0 {
		return nil, false
	}
	return _cgd.GetDict(_aggab.Get(0))
}
func _gcad(_edddg *_gg.Document) error {
	_gdbe, _gacf := _edddg.FindCatalog()
	if !_gacf {
		return _d.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _gdbe.Object.Get("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073") != nil {
		_gdbe.Object.Remove("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073")
	}
	return nil
}
func _aaebe(_acaaf *_gg.Document) error {
	_fcd, _ecg := _acaaf.FindCatalog()
	if !_ecg {
		return _d.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_fbag, _ecg := _cgd.GetDict(_fcd.Object.Get("\u004e\u0061\u006de\u0073"))
	if !_ecg {
		return nil
	}
	if _fbag.Get("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073") != nil {
		_fbag.Remove("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073")
	}
	return nil
}
func _acfc(_fcba *_f.CompliancePdfReader, _bgfc standardType, _aaeda bool) (_cgbba []ViolatedRule) {
	_abgf, _bffg := _ecgd(_fcba)
	if !_bffg {
		return []ViolatedRule{_cc("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0063a\u0074a\u006c\u006f\u0067\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064\u002e")}
	}
	_bgdfc := _abgf.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	if _bgdfc == nil {
		return []ViolatedRule{_cc("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0054\u0068\u0065\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u006f\u0066\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074ai\u006e\u0020\u0074\u0068\u0065\u0020\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0068\u006f\u0073\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0061\u0020m\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020s\u0074\u0072\u0065\u0061\u006d")}
	}
	_cdcb, _bffg := _cgd.GetStream(_bgdfc)
	if !_bffg {
		return []ViolatedRule{_cc("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0054\u0068\u0065\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u006f\u0066\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074ai\u006e\u0020\u0074\u0068\u0065\u0020\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0068\u006f\u0073\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0061\u0020m\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020s\u0074\u0072\u0065\u0061\u006d")}
	}
	_cbfdb, _cbca := _cg.LoadDocument(_cdcb.Stream)
	if _cbca != nil {
		return []ViolatedRule{_cc("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d4", "\u0041\u006c\u006c\u0020\u006de\u0074\u0061\u0064a\u0074\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020i\u006e \u0074\u0068\u0065\u0020\u0050\u0044\u0046 \u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0066\u006f\u0072\u006d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065ci\u0066\u0069\u0063\u0061\u0074\u0069\u006fn\u002e\u0020\u0041\u006c\u006c\u0020c\u006fn\u0074\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0061\u006c\u006c\u0020\u0058\u004d\u0050\u0020p\u0061\u0063\u006b\u0065\u0074\u0073 \u0073h\u0061\u006c\u006c \u0062\u0065\u0020\u0077\u0065\u006c\u006c\u002d\u0066o\u0072\u006de\u0064")}
	}
	_eadee := _cbfdb.GetGoXmpDocument()
	var _gebc []*_aca.Namespace
	for _, _bbfdf := range _eadee.Namespaces() {
		switch _bbfdf.Name {
		case _be.NsDc.Name, _g.NsPDF.Name, _ee.NsXmp.Name, _ae.NsXmpRights.Name, _da.Namespace.Name, _df.Namespace.Name, _bf.NsXmpMM.Name, _df.FieldNS.Name, _df.SchemaNS.Name, _df.PropertyNS.Name, "\u0073\u0074\u0045v\u0074", "\u0073\u0074\u0056e\u0072", "\u0073\u0074\u0052e\u0066", "\u0073\u0074\u0044i\u006d", "\u0078a\u0070\u0047\u0049\u006d\u0067", "\u0073\u0074\u004ao\u0062", "\u0078\u006d\u0070\u0069\u0064\u0071":
			continue
		}
		_gebc = append(_gebc, _bbfdf)
	}
	_gacce := true
	_cefe, _cbca := _cbfdb.GetPdfaExtensionSchemas()
	if _cbca == nil {
		for _, _dcfc := range _gebc {
			var _cbagc bool
			for _ggefa := range _cefe {
				if _dcfc.URI == _cefe[_ggefa].NamespaceURI {
					_cbagc = true
					break
				}
			}
			if !_cbagc {
				_gacce = false
				break
			}
		}
	} else {
		_gacce = false
	}
	if !_gacce {
		_cgbba = append(_cgbba, _cc("\u0036.\u0036\u002e\u0032\u002e\u0033\u002d7", "\u0041\u006c\u006c\u0020\u0070\u0072\u006f\u0070e\u0072\u0074\u0069e\u0073\u0020\u0073\u0070\u0065\u0063i\u0066\u0069\u0065\u0064\u0020\u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072m\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0075s\u0065\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0073\u0063he\u006da\u0073 \u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002c\u0020\u0049\u0053\u004f\u0020\u0031\u00390\u0030\u0035-\u0031\u0020\u006f\u0072\u0020\u0074h\u0069s\u0020\u0070\u0061\u0072\u0074\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u0020\u0031\u0039\u0030\u0030\u0035\u002c\u0020o\u0072\u0020\u0061\u006e\u0079\u0020e\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073c\u0068\u0065\u006das\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006fm\u0070\u006c\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0036\u002e\u0036\u002e\u0032.\u0033\u002e\u0032\u002e"))
	}
	_ggeea, _bffg := _cbfdb.GetPdfAID()
	if !_bffg {
		_cgbba = append(_cgbba, _cc("\u0036.\u0036\u002e\u0034\u002d\u0031", "\u0054\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0061n\u0064\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020\u006c\u0065\u0076\u0065l\u0020\u006f\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0070e\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073\u0063h\u0065\u006da."))
	} else {
		if _ggeea.Part != _bgfc._ba {
			_cgbba = append(_cgbba, _cc("\u0036.\u0036\u002e\u0034\u002d\u0032", "\u0054h\u0065\u0020\u0076\u0061lue\u0020\u006f\u0066\u0020p\u0064\u0066\u0061\u0069\u0064\u003a\u0070\u0061\u0072\u0074 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0072\u0074\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u002019\u0030\u0030\u0035 \u0074\u006f\u0020\u0077\u0068i\u0063h\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065 \u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0073\u002e"))
		}
		if _bgfc._fa == "\u0041" && _ggeea.Conformance != "\u0041" {
			_cgbba = append(_cgbba, _cc("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		} else if _bgfc._fa == "\u0055" && (_ggeea.Conformance != "\u0041" && _ggeea.Conformance != "\u0055") {
			_cgbba = append(_cgbba, _cc("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		} else if _bgfc._fa == "\u0042" && (_ggeea.Conformance != "\u0041" && _ggeea.Conformance != "\u0042" && _ggeea.Conformance != "\u0055") {
			_cgbba = append(_cgbba, _cc("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		}
	}
	return _cgbba
}
func _fggfa(_beaf *_f.PdfFont, _ccgdd *_cgd.PdfObjectDictionary) ViolatedRule {
	const (
		_cfgaf = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0036\u002d\u0033"
		_eefae = "\u0041l\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072u\u0065\u0054\u0079p\u0065\u0020\u0066\u006f\u006e\u0074s\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065n\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068e\u0020\u0066\u006f\u006e\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"
	)
	var _fcfbg string
	if _acfdgc, _bdga := _cgd.GetName(_ccgdd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _bdga {
		_fcfbg = _acfdgc.String()
	}
	if _fcfbg != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _cb
	}
	_bdbd := _beaf.FontDescriptor()
	_gefgf, _cedbb := _cgd.GetIntVal(_bdbd.Flags)
	if !_cedbb {
		_ec.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _cc(_cfgaf, _eefae)
	}
	_caagg := (uint32(_gefgf) >> 3) & 1
	_gaed := _caagg != 0
	if !_gaed {
		return _cb
	}
	if _ccgdd.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067") != nil {
		return _cc(_cfgaf, _eefae)
	}
	return _cb
}
func _gcce(_cgfg *_f.PdfFont, _bfeg *_cgd.PdfObjectDictionary) ViolatedRule {
	const (
		_bfab = "\u0036.\u0033\u002e\u0037\u002d\u0031"
		_dcgd = "\u0041\u006cl \u006e\u006f\u006e\u002d\u0073\u0079\u006db\u006f\u006c\u0069\u0063\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065\u0020\u0066o\u006e\u0074s\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020e\u0069\u0074h\u0065\u0072\u0020\u004d\u0061\u0063\u0052\u006f\u006d\u0061\u006e\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041\u006e\u0073i\u0045n\u0063\u006f\u0064\u0069n\u0067\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0066o\u0072\u0020t\u0068\u0065 \u0045n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006b\u0065\u0079 \u0069\u006e\u0020t\u0068e\u0020\u0046o\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0072\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0066\u006f\u0072 \u0074\u0068\u0065\u0020\u0042\u0061\u0073\u0065\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074\u0068\u0065 \u0064i\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0077\u0068\u0069\u0063\u0068\u0020\u0069s\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0074\u0068e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006be\u0079\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0046\u006f\u006e\u0074 \u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u002e\u0020\u0049\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e, \u006eo\u0020n\u006f\u006e\u002d\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072\u0075\u0065\u0054\u0079p\u0065 \u0066\u006f\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0020\u0061\u0020\u0044\u0069\u0066\u0066e\u0072\u0065\u006e\u0063\u0065\u0073\u0020a\u0072\u0072\u0061\u0079\u0020\u0075n\u006c\u0065s\u0073\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0074h\u0065\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u006e\u0061\u006d\u0065\u0073 \u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0044\u0069f\u0066\u0065\u0072\u0065\u006ec\u0065\u0073\u0020a\u0072\u0072\u0061\u0079\u0020\u0061\u0072\u0065\u0020\u006c\u0069\u0073\u0074\u0065\u0064 \u0069\u006e \u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065 G\u006c\u0079\u0070\u0068\u0020\u004c\u0069\u0073t\u0020\u0061\u006e\u0064\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066o\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072a\u006d\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073t\u0020\u0074\u0068\u0065\u0020\u004d\u0069\u0063\u0072o\u0073o\u0066\u0074\u0020\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0028\u0033\u002c\u0031 \u2013 P\u006c\u0061\u0074\u0066\u006f\u0072\u006d\u0020I\u0044\u003d\u0033\u002c\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067 I\u0044\u003d\u0031\u0029\u0020\u0065\u006e\u0063\u006f\u0064i\u006e\u0067 \u0069\u006e\u0020t\u0068\u0065\u0020'\u0063\u006d\u0061\u0070\u0027\u0020\u0074\u0061\u0062\u006c\u0065\u002e"
	)
	var _ebeb string
	if _aage, _dffca := _cgd.GetName(_bfeg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _dffca {
		_ebeb = _aage.String()
	}
	if _ebeb != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _cb
	}
	_ddcd := _cgfg.FontDescriptor()
	_eeee, _gafde := _cgd.GetIntVal(_ddcd.Flags)
	if !_gafde {
		_ec.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _cc(_bfab, _dcgd)
	}
	_cgdd := (uint32(_eeee) >> 3) != 0
	if _cgdd {
		return _cb
	}
	_fadgd, _gafde := _cgd.GetName(_bfeg.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	if !_gafde {
		return _cc(_bfab, _dcgd)
	}
	switch _fadgd.String() {
	case "\u004d\u0061c\u0052\u006f\u006da\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067":
		return _cb
	default:
		return _cc(_bfab, _dcgd)
	}
}
func _aabc(_agcd *_f.CompliancePdfReader) ViolatedRule { return _cb }
func _dce(_fge *_f.XObjectImage, _ccd imageModifications) error {
	_cae, _afcg := _fge.ToImage()
	if _afcg != nil {
		return _afcg
	}
	if _ccd._ebf != nil {
		_fge.Filter = _ccd._ebf
	}
	_geeb := _cgd.MakeDict()
	_geeb.Set("\u0051u\u0061\u006c\u0069\u0074\u0079", _cgd.MakeInteger(100))
	_geeb.Set("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr", _cgd.MakeInteger(1))
	_fge.Decode = nil
	if _afcg = _fge.SetImage(_cae, nil); _afcg != nil {
		return _afcg
	}
	_fge.ToPdfObject()
	return nil
}

// ValidateStandard checks if provided input CompliancePdfReader matches rules that conforms PDF/A-2 standard.
func (_gfbfa *profile2) ValidateStandard(r *_f.CompliancePdfReader) error {
	_cga := VerificationError{ConformanceLevel: _gfbfa._geg._ba, ConformanceVariant: _gfbfa._geg._fa}
	if _edaa := _dcbe(r); _edaa != _cb {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _edaa)
	}
	if _agf := _dgcc(r); _agf != _cb {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _agf)
	}
	if _feda := _fggc(r); _feda != _cb {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _feda)
	}
	if _gceb := _eecc(r); _gceb != _cb {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _gceb)
	}
	if _cbbc := _cabc(r); _cbbc != _cb {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _cbbc)
	}
	if _abd := _efcb(r); len(_abd) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _abd...)
	}
	if _cfad := _eaedd(r); len(_cfad) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _cfad...)
	}
	if _eafa := _aafg(r); len(_eafa) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _eafa...)
	}
	if _gbdcfa := _gaee(r); _gbdcfa != _cb {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _gbdcfa)
	}
	if _gfadd := _affcf(r); len(_gfadd) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _gfadd...)
	}
	if _eeecd := _eeaf(r); len(_eeecd) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _eeecd...)
	}
	if _cca := _bdgda(r); _cca != _cb {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _cca)
	}
	if _fgg := _gfcfc(r); len(_fgg) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _fgg...)
	}
	if _ccdc := _cbff(r); len(_ccdc) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _ccdc...)
	}
	if _fcc := _eged(r); _fcc != _cb {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _fcc)
	}
	if _afgg := _cegc(r); len(_afgg) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _afgg...)
	}
	if _gdfa := _eabg(r); len(_gdfa) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _gdfa...)
	}
	if _caee := _egeea(r); _caee != _cb {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _caee)
	}
	if _fgdag := _ccbf(r); len(_fgdag) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _fgdag...)
	}
	if _fccc := _bdagc(r, _gfbfa._geg); len(_fccc) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _fccc...)
	}
	if _fcbe := _bdge(r); len(_fcbe) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _fcbe...)
	}
	if _gbca := _acbd(r); len(_gbca) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _gbca...)
	}
	if _gdg := _gcbc(r); len(_gdg) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _gdg...)
	}
	if _gbag := _fdga(r); _gbag != _cb {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _gbag)
	}
	if _fbfc := _bdce(r); len(_fbfc) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _fbfc...)
	}
	if _acca := _cddgb(r); _acca != _cb {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _acca)
	}
	if _fgaac := _acfc(r, _gfbfa._geg, false); len(_fgaac) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _fgaac...)
	}
	if _gfbfa._geg == _fae() {
		if _eadec := _bgeccf(r); len(_eadec) != 0 {
			_cga.ViolatedRules = append(_cga.ViolatedRules, _eadec...)
		}
	}
	if _eacb := _gebdb(r); len(_eacb) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _eacb...)
	}
	if _fcgg := _ebced(r); len(_fcgg) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _fcgg...)
	}
	if _fdbfe := _gcdd(r); len(_fdbfe) != 0 {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _fdbfe...)
	}
	if _fbdc := _adeeb(r); _fbdc != _cb {
		_cga.ViolatedRules = append(_cga.ViolatedRules, _fbdc)
	}
	if len(_cga.ViolatedRules) > 0 {
		_af.Slice(_cga.ViolatedRules, func(_bbdf, _fgae int) bool {
			return _cga.ViolatedRules[_bbdf].RuleNo < _cga.ViolatedRules[_fgae].RuleNo
		})
		return _cga
	}
	return nil
}
func _gfbf(_bef *_gg.Document) error {
	_cbf := func(_ecbe *_cgd.PdfObjectDictionary) error {
		if _eed := _ecbe.Get("\u0053\u004d\u0061s\u006b"); _eed != nil {
			_ecbe.Set("\u0053\u004d\u0061s\u006b", _cgd.MakeName("\u004e\u006f\u006e\u0065"))
		}
		_dcg := _ecbe.Get("\u0043\u0041")
		if _dcg != nil {
			_gfgb, _cce := _cgd.GetNumberAsFloat(_dcg)
			if _cce != nil {
				_ec.Log.Debug("\u0045x\u0074\u0047S\u0074\u0061\u0074\u0065 \u006f\u0062\u006ae\u0063\u0074\u0020\u0043\u0041\u0020\u0076\u0061\u006cue\u0020\u0069\u0073 \u006e\u006ft\u0020\u0061\u0020\u0066\u006c\u006fa\u0074\u003a \u0025\u0076", _cce)
				_gfgb = 0
			}
			if _gfgb != 1.0 {
				_ecbe.Set("\u0043\u0041", _cgd.MakeFloat(1.0))
			}
		}
		_dcg = _ecbe.Get("\u0063\u0061")
		if _dcg != nil {
			_bcb, _gcc := _cgd.GetNumberAsFloat(_dcg)
			if _gcc != nil {
				_ec.Log.Debug("\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0027\u0063\u0061\u0027\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _gcc)
				_bcb = 0
			}
			if _bcb != 1.0 {
				_ecbe.Set("\u0063\u0061", _cgd.MakeFloat(1.0))
			}
		}
		_ed := _ecbe.Get("\u0042\u004d")
		if _ed != nil {
			_abc, _eea := _cgd.GetName(_ed)
			if !_eea {
				_ec.Log.Debug("E\u0078\u0074\u0047\u0053\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0027\u0042\u004d\u0027\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061m\u0065")
				_abc = _cgd.MakeName("")
			}
			_aae := _abc.String()
			switch _aae {
			case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065":
			default:
				_ecbe.Set("\u0042\u004d", _cgd.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
			}
		}
		_ebb := _ecbe.Get("\u0054\u0052")
		if _ebb != nil {
			_ec.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0054\u0052\u0020\u006b\u0065\u0079")
			_ecbe.Remove("\u0054\u0052")
		}
		_fgf := _ecbe.Get("\u0054\u0052\u0032")
		if _fgf != nil {
			_bbg := _fgf.String()
			if _bbg != "\u0044e\u0066\u0061\u0075\u006c\u0074" {
				_ec.Log.Debug("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074\u0065 o\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073 \u0054\u00522\u0020\u006b\u0065y\u0020\u0077\u0069\u0074\u0068\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065r\u0020\u0074ha\u006e\u0020\u0044e\u0066\u0061\u0075\u006c\u0074")
				_ecbe.Set("\u0054\u0052\u0032", _cgd.MakeName("\u0044e\u0066\u0061\u0075\u006c\u0074"))
			}
		}
		return nil
	}
	_daef, _fec := _bef.GetPages()
	if !_fec {
		return nil
	}
	for _, _bbb := range _daef {
		_cba, _dcf := _bbb.GetResources()
		if !_dcf {
			continue
		}
		_daee, _ce := _cgd.GetDict(_cba.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
		if !_ce {
			return nil
		}
		_gfba := _daee.Keys()
		for _, _ceg := range _gfba {
			_ag, _cgb := _cgd.GetDict(_daee.Get(_ceg))
			if !_cgb {
				continue
			}
			_ebe := _cbf(_ag)
			if _ebe != nil {
				continue
			}
		}
	}
	for _, _faa := range _daef {
		_aag, _cgf := _faa.GetContents()
		if !_cgf {
			return nil
		}
		for _, _add := range _aag {
			_dfb, _bgec := _add.GetData()
			if _bgec != nil {
				continue
			}
			_fdag := _ff.NewContentStreamParser(string(_dfb))
			_cbaf, _bgec := _fdag.Parse()
			if _bgec != nil {
				continue
			}
			for _, _eacg := range *_cbaf {
				if len(_eacg.Params) == 0 {
					continue
				}
				_, _gfef := _cgd.GetName(_eacg.Params[0])
				if !_gfef {
					continue
				}
				_aeg, _efg := _faa.GetResourcesXObject()
				if !_efg {
					continue
				}
				for _, _fbfa := range _aeg.Keys() {
					_dg, _fga := _cgd.GetStream(_aeg.Get(_fbfa))
					if !_fga {
						continue
					}
					_fgfd, _fga := _cgd.GetDict(_dg.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
					if !_fga {
						continue
					}
					_fafc, _fga := _cgd.GetDict(_fgfd.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
					if !_fga {
						continue
					}
					for _, _daf := range _fafc.Keys() {
						_ccb, _ade := _cgd.GetDict(_fafc.Get(_daf))
						if !_ade {
							continue
						}
						_fdab := _cbf(_ccb)
						if _fdab != nil {
							continue
						}
					}
				}
			}
		}
	}
	return nil
}

type colorspaceModification struct {
	_gaa _cf.ColorConverter
	_bc  _f.PdfColorspace
}

func _fff(_ecc *_gg.Document, _efe int) {
	if _ecc.Version.Major == 0 {
		_ecc.Version.Major = 1
	}
	if _ecc.Version.Minor < _efe {
		_ecc.Version.Minor = _efe
	}
}
func _dbbc(_aabd *_cgd.PdfObjectDictionary) ViolatedRule {
	const (
		_acef  = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0032"
		_egbca = "IS\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u00209\u002e\u0037\u002e\u0034\u002c\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0031\u0031\u0037\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068a\u0074\u0020\u0061\u006c\u006c\u0020\u0065m\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0032\u0020\u0043\u0049\u0044\u0046\u006fn\u0074\u0073\u0020\u0069n\u0020t\u0068e\u0020\u0043\u0049D\u0046\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0043\u0049\u0044\u0054\u006fG\u0049\u0044M\u0061\u0070\u0020\u0065\u006e\u0074\u0072\u0079 \u0074\u0068\u0061\u0074\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020\u0061\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u006d\u0061\u0070p\u0069\u006e\u0067 f\u0072\u006f\u006d \u0043\u0049\u0044\u0073\u0020\u0074\u006f\u0020\u0067\u006c\u0079p\u0068 \u0069\u006e\u0064\u0069c\u0065\u0073\u0020\u006fr\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u0049d\u0065\u006e\u0074\u0069\u0074\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0031\u0031\u0037\u002e"
	)
	var _bbag string
	if _ebfed, _ebad := _cgd.GetName(_aabd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _ebad {
		_bbag = _ebfed.String()
	}
	if _bbag != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		return _cb
	}
	if _aabd.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070") == nil {
		return _cc(_acef, _egbca)
	}
	return _cb
}
func _eabg(_bgde *_f.CompliancePdfReader) (_accde []ViolatedRule) { return _accde }

type standardType struct {
	_ba int
	_fa string
}

// XmpOptions are the options used by the optimization of the XMP metadata.
type XmpOptions struct {

	// Copyright information.
	Copyright string

	// OriginalDocumentID is the original document identifier.
	// By default, if this field is empty the value is extracted from the XMP Metadata or generated UUID.
	OriginalDocumentID string

	// DocumentID is the original document identifier.
	// By default, if this field is empty the value is extracted from the XMP Metadata or generated UUID.
	DocumentID string

	// InstanceID is the original document identifier.
	// By default, if this field is empty the value is set to generated UUID.
	InstanceID string

	// NewDocumentVersion is a flag that defines if a document was overwritten.
	// If the new document was created this should be true. On changing given document file, and overwriting it it should be true.
	NewDocumentVersion bool

	// MarshalIndent defines marshaling indent of the XMP metadata.
	MarshalIndent string

	// MarshalPrefix defines marshaling prefix of the XMP metadata.
	MarshalPrefix string
}

// Validate checks if provided input document reader matches given PDF/A profile.
func Validate(d *_f.CompliancePdfReader, profile Profile) error { return profile.ValidateStandard(d) }
func _cc(_bb string, _gee string) ViolatedRule                  { return ViolatedRule{RuleNo: _bb, Detail: _gee} }
func _cac(_gcaf *_gg.Document) error {
	for _, _egd := range _gcaf.Objects {
		_eeb, _bea := _cgd.GetDict(_egd)
		if !_bea {
			continue
		}
		_bfdd := _eeb.Get("\u0054\u0079\u0070\u0065")
		if _bfdd == nil {
			continue
		}
		if _agec, _ggb := _cgd.GetName(_bfdd); _ggb && _agec.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_bce, _gcb := _cgd.GetBool(_eeb.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _gcb {
			if bool(*_bce) {
				_eeb.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _cgd.MakeBool(false))
			}
		}
		_fee := _eeb.Get("\u0041")
		if _fee != nil {
			_eeb.Remove("\u0041")
		}
		_bacg, _gcb := _cgd.GetArray(_eeb.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if _gcb {
			for _bfg := 0; _bfg < _bacg.Len(); _bfg++ {
				_aee, _gbba := _cgd.GetDict(_bacg.Get(_bfg))
				if !_gbba {
					continue
				}
				if _aee.Get("\u0041\u0041") != nil {
					_aee.Remove("\u0041\u0041")
				}
			}
		}
	}
	return nil
}
func _defc(_cega *_f.PdfPageResources, _fafd *_ff.ContentStreamOperations, _acee bool) ([]byte, error) {
	var _egde bool
	for _, _cfg := range *_fafd {
	_bfde:
		switch _cfg.Operand {
		case "\u0042\u0049":
			_fcg, _bagd := _cfg.Params[0].(*_ff.ContentStreamInlineImage)
			if !_bagd {
				break
			}
			_adeb, _gccc := _fcg.GetColorSpace(_cega)
			if _gccc != nil {
				return nil, _gccc
			}
			switch _adeb.(type) {
			case *_f.PdfColorspaceDeviceCMYK:
				if _acee {
					break _bfde
				}
			case *_f.PdfColorspaceDeviceGray:
			case *_f.PdfColorspaceDeviceRGB:
				if !_acee {
					break _bfde
				}
			default:
				break _bfde
			}
			_egde = true
			_dfaa, _gccc := _fcg.ToImage(_cega)
			if _gccc != nil {
				return nil, _gccc
			}
			_bgfd, _gccc := _dfaa.ToGoImage()
			if _gccc != nil {
				return nil, _gccc
			}
			if _acee {
				_bgfd, _gccc = _cf.CMYKConverter.Convert(_bgfd)
			} else {
				_bgfd, _gccc = _cf.NRGBAConverter.Convert(_bgfd)
			}
			if _gccc != nil {
				return nil, _gccc
			}
			_daa, _bagd := _bgfd.(_cf.Image)
			if !_bagd {
				return nil, _d.New("\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074 \u0069\u006d\u0070\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002eI\u006d\u0061\u0067\u0065")
			}
			_efege := _daa.Base()
			_caf := _f.Image{Width: int64(_efege.Width), Height: int64(_efege.Height), BitsPerComponent: int64(_efege.BitsPerComponent), ColorComponents: _efege.ColorComponents, Data: _efege.Data}
			_caf.SetDecode(_efege.Decode)
			_caf.SetAlpha(_efege.Alpha)
			_dfcd, _gccc := _fcg.GetEncoder()
			if _gccc != nil {
				_dfcd = _cgd.NewFlateEncoder()
			}
			_afb, _gccc := _ff.NewInlineImageFromImage(_caf, _dfcd)
			if _gccc != nil {
				return nil, _gccc
			}
			_cfg.Params[0] = _afb
		case "\u0047", "\u0067":
			if len(_cfg.Params) != 1 {
				break
			}
			_cebg, _bdcf := _cgd.GetNumberAsFloat(_cfg.Params[0])
			if _bdcf != nil {
				break
			}
			if _acee {
				_cfg.Params = []_cgd.PdfObject{_cgd.MakeFloat(0), _cgd.MakeFloat(0), _cgd.MakeFloat(0), _cgd.MakeFloat(1 - _cebg)}
				_cfgc := "\u004b"
				if _cfg.Operand == "\u0067" {
					_cfgc = "\u006b"
				}
				_cfg.Operand = _cfgc
			} else {
				_cfg.Params = []_cgd.PdfObject{_cgd.MakeFloat(_cebg), _cgd.MakeFloat(_cebg), _cgd.MakeFloat(_cebg)}
				_dgcg := "\u0052\u0047"
				if _cfg.Operand == "\u0067" {
					_dgcg = "\u0072\u0067"
				}
				_cfg.Operand = _dgcg
			}
			_egde = true
		case "\u0052\u0047", "\u0072\u0067":
			if !_acee {
				break
			}
			if len(_cfg.Params) != 3 {
				break
			}
			_abee, _aefb := _cgd.GetNumbersAsFloat(_cfg.Params)
			if _aefb != nil {
				break
			}
			_egde = true
			_geec, _fgab, _deae := _abee[0], _abee[1], _abee[2]
			_cafa, _cbcge, _eafb, _gce := _c.RGBToCMYK(uint8(_geec*255), uint8(_fgab*255), uint8(255*_deae))
			_cfg.Params = []_cgd.PdfObject{_cgd.MakeFloat(float64(_cafa) / 255), _cgd.MakeFloat(float64(_cbcge) / 255), _cgd.MakeFloat(float64(_eafb) / 255), _cgd.MakeFloat(float64(_gce) / 255)}
			_aaaa := "\u004b"
			if _cfg.Operand == "\u0072\u0067" {
				_aaaa = "\u006b"
			}
			_cfg.Operand = _aaaa
		case "\u004b", "\u006b":
			if _acee {
				break
			}
			if len(_cfg.Params) != 4 {
				break
			}
			_gace, _cdeb := _cgd.GetNumbersAsFloat(_cfg.Params)
			if _cdeb != nil {
				break
			}
			_abg, _eff, _bdg, _fce := _gace[0], _gace[1], _gace[2], _gace[3]
			_dfbd, _fdeb, _bga := _c.CMYKToRGB(uint8(255*_abg), uint8(255*_eff), uint8(255*_bdg), uint8(255*_fce))
			_cfg.Params = []_cgd.PdfObject{_cgd.MakeFloat(float64(_dfbd) / 255), _cgd.MakeFloat(float64(_fdeb) / 255), _cgd.MakeFloat(float64(_bga) / 255)}
			_bbga := "\u0052\u0047"
			if _cfg.Operand == "\u006b" {
				_bbga = "\u0072\u0067"
			}
			_cfg.Operand = _bbga
			_egde = true
		}
	}
	if !_egde {
		return nil, nil
	}
	_gbdcf := _ff.NewContentCreator()
	for _, _fagb := range *_fafd {
		_gbdcf.AddOperand(*_fagb)
	}
	_aabg := _gbdcf.Bytes()
	return _aabg, nil
}

// Conformance gets the PDF/A conformance.
func (_bdag *profile1) Conformance() string { return _bdag._acfg._fa }
func _gd() standardType                     { return standardType{_ba: 1, _fa: "\u0042"} }
func _bdd(_ffbg *_gg.Document) error {
	_acag := map[string]*_cgd.PdfObjectDictionary{}
	_cff := _gc.NewFinder(&_gc.FinderOpts{Extensions: []string{"\u002e\u0074\u0074\u0066"}})
	_geed := map[_cgd.PdfObject]struct{}{}
	_ccdf := map[_cgd.PdfObject]struct{}{}
	for _, _efgc := range _ffbg.Objects {
		_cd, _gbbg := _cgd.GetDict(_efgc)
		if !_gbbg {
			continue
		}
		_de := _cd.Get("\u0054\u0079\u0070\u0065")
		if _de == nil {
			continue
		}
		if _fecb, _cbgg := _cgd.GetName(_de); _cbgg && _fecb.String() != "\u0046\u006f\u006e\u0074" {
			continue
		}
		if _, _deg := _geed[_efgc]; _deg {
			continue
		}
		_edc, _edg := _f.NewPdfFontFromPdfObject(_cd)
		if _edg != nil {
			_ec.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066\u0072\u006fm\u0020\u006fb\u006a\u0065\u0063\u0074")
			return _edg
		}
		_ebg, _edg := _edc.GetFontDescriptor()
		if _edg != nil {
			return _edg
		}
		if _ebg != nil && (_ebg.FontFile != nil || _ebg.FontFile2 != nil || _ebg.FontFile3 != nil) {
			continue
		}
		_ace := _edc.BaseFont()
		if _ace == "" {
			return _a.Errorf("\u006f\u006e\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0073\u0020\u0073\u0079\u006e\u0074\u0061\u0078\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069d\u0020\u002d\u0020\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u0075\u006ed\u0065\u0066\u0069n\u0065\u0064\u003a\u0020\u0025\u0073", _cd.String())
		}
		_aff, _ead := _acag[_ace]
		if !_ead {
			if len(_ace) > 7 && _ace[6] == '+' {
				_ace = _ace[7:]
			}
			_agd := []string{_ace, "\u0054i\u006de\u0073\u0020\u004e\u0065\u0077\u0020\u0052\u006f\u006d\u0061\u006e", "\u0041\u0072\u0069a\u006c", "D\u0065\u006a\u0061\u0056\u0075\u0020\u0053\u0061\u006e\u0073"}
			for _, _geb := range _agd {
				_ec.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a \u0073\u0065\u0061\u0072\u0063\u0068\u0069\u006e\u0067\u0020\u0073\u0079\u0073t\u0065\u006d\u0020\u0066\u006f\u006e\u0074 \u0060\u0025\u0073\u0060", _geb)
				if _aff, _ead = _acag[_geb]; _ead {
					break
				}
				_cfb := _cff.Match(_geb)
				if _cfb == nil {
					_ec.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _geb)
					continue
				}
				_gac, _cda := _f.NewPdfFontFromTTFFile(_cfb.Filename)
				if _cda != nil {
					return _cda
				}
				_dfe := _gac.FontDescriptor()
				if _dfe.FontFile != nil {
					if _, _ead = _ccdf[_dfe.FontFile]; !_ead {
						_ffbg.Objects = append(_ffbg.Objects, _dfe.FontFile)
						_ccdf[_dfe.FontFile] = struct{}{}
					}
				}
				if _dfe.FontFile2 != nil {
					if _, _ead = _ccdf[_dfe.FontFile2]; !_ead {
						_ffbg.Objects = append(_ffbg.Objects, _dfe.FontFile2)
						_ccdf[_dfe.FontFile2] = struct{}{}
					}
				}
				if _dfe.FontFile3 != nil {
					if _, _ead = _ccdf[_dfe.FontFile3]; !_ead {
						_ffbg.Objects = append(_ffbg.Objects, _dfe.FontFile3)
						_ccdf[_dfe.FontFile3] = struct{}{}
					}
				}
				_eaf, _fca := _gac.ToPdfObject().(*_cgd.PdfIndirectObject)
				if !_fca {
					_ec.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
					continue
				}
				_baae, _fca := _eaf.PdfObject.(*_cgd.PdfObjectDictionary)
				if !_fca {
					_ec.Log.Debug("\u0046\u006fn\u0074\u0020\u0074\u0079p\u0065\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
					continue
				}
				_acag[_geb] = _baae
				_aff = _baae
				break
			}
			if _aff == nil {
				_ec.Log.Debug("\u004e\u006f\u0020\u006d\u0061\u0074\u0063\u0068\u0069\u006eg\u0020\u0066\u006f\u006e\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0066\u006f\u0072\u003a\u0020\u0025\u0073", _edc.BaseFont())
				return _d.New("\u006e\u006f m\u0061\u0074\u0063h\u0069\u006e\u0067\u0020fon\u0074 f\u006f\u0075\u006e\u0064\u0020\u0069\u006e t\u0068\u0065\u0020\u0073\u0079\u0073\u0074e\u006d")
			}
		}
		for _, _eec := range _aff.Keys() {
			_cd.Set(_eec, _aff.Get(_eec))
		}
		_gdc := _aff.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
		if _gdc != nil {
			if _, _ead = _ccdf[_gdc]; !_ead {
				_ffbg.Objects = append(_ffbg.Objects, _gdc)
				_ccdf[_gdc] = struct{}{}
			}
		}
		_geed[_efgc] = struct{}{}
		_bae := _cd.Get("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072")
		if _bae != nil {
			_ffbg.Objects = append(_ffbg.Objects, _bae)
			_ccdf[_bae] = struct{}{}
		}
	}
	return nil
}
func _bbba(_beca *_f.CompliancePdfReader) ViolatedRule {
	if _beca.ParserMetadata().HeaderPosition() != 0 {
		return _cc("\u0036.\u0031\u002e\u0032\u002d\u0031", "h\u0065\u0061\u0064\u0065\u0072\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0074\u0020\u0074\u0068\u0065\u0020fi\u0072\u0073\u0074 \u0062y\u0074\u0065")
	}
	return _cb
}
func _cabd(_dgad *_gg.Document) error {
	for _, _baad := range _dgad.Objects {
		_edddc, _fgef := _cgd.GetDict(_baad)
		if !_fgef {
			continue
		}
		_edea := _edddc.Get("\u0054\u0079\u0070\u0065")
		if _edea == nil {
			continue
		}
		if _ebcd, _dde := _cgd.GetName(_edea); _dde && _ebcd.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_bdad, _cgcf := _cgd.GetBool(_edddc.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _cgcf && bool(*_bdad) {
			_edddc.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _cgd.MakeBool(false))
		}
		if _edddc.Get("\u0058\u0046\u0041") != nil {
			_edddc.Remove("\u0058\u0046\u0041")
		}
	}
	_ccca, _efgd := _dgad.FindCatalog()
	if !_efgd {
		return _d.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _ccca.Object.Get("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067") != nil {
		_ccca.Object.Remove("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067")
	}
	return nil
}
func _fbddg(_edbb *_f.PdfFont, _ebecb *_cgd.PdfObjectDictionary) ViolatedRule {
	const (
		_fdec = "\u0036.\u0033\u002e\u0035\u002d\u0032"
		_gfea = "\u0046\u006f\u0072\u0020\u0061l\u006c\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074 \u0073\u0075bs\u0065\u0074\u0073 \u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074he\u0020f\u006f\u006e\u0074\u0020\u0064\u0065s\u0063r\u0069\u0070\u0074o\u0072\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006ec\u006c\u0075\u0064e\u0020\u0061\u0020\u0043\u0068\u0061\u0072\u0053\u0065\u0074\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u006c\u0069\u0073\u0074\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0063\u0068\u0061\u0072a\u0063\u0074\u0065\u0072 \u006e\u0061\u006d\u0065\u0073\u0020d\u0065\u0066i\u006e\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020f\u006f\u006e\u0074\u0020s\u0075\u0062\u0073\u0065\u0074, \u0061\u0073 \u0064\u0065s\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e \u0050\u0044\u0046\u0020\u0052e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0054\u0061\u0062\u006ce\u0020\u0035\u002e1\u0038\u002e"
	)
	var _acbc string
	if _fdbeb, _gbadf := _cgd.GetName(_ebecb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _gbadf {
		_acbc = _fdbeb.String()
	}
	if _acbc != "\u0054\u0079\u0070e\u0031" {
		return _cb
	}
	if _ecb.IsStdFont(_ecb.StdFontName(_edbb.BaseFont())) {
		return _cb
	}
	_ggdg := _edbb.FontDescriptor()
	if _ggdg.CharSet == nil {
		return _cc(_fdec, _gfea)
	}
	return _cb
}
func _fdga(_cdafe *_f.CompliancePdfReader) (_gbgge ViolatedRule) {
	_eecb, _agcb := _ecgd(_cdafe)
	if !_agcb {
		return _cb
	}
	_bcfge, _agcb := _cgd.GetDict(_eecb.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
	if !_agcb {
		return _cb
	}
	_fadga, _agcb := _cgd.GetArray(_bcfge.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_agcb {
		return _cb
	}
	for _baef := 0; _baef < _fadga.Len(); _baef++ {
		_becca, _ggdfg := _cgd.GetDict(_fadga.Get(_baef))
		if !_ggdfg {
			continue
		}
		if _becca.Get("\u0041") != nil {
			return _cc("\u0036.\u0034\u002e\u0031\u002d\u0032", "\u0041\u0020\u0046\u0069\u0065\u006c\u0064\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0041 o\u0072\u0020\u0041\u0041\u0020\u006b\u0065\u0079\u0073\u002e")
		}
		if _becca.Get("\u0041\u0041") != nil {
			return _cc("\u0036.\u0034\u002e\u0031\u002d\u0032", "\u0041\u0020\u0046\u0069\u0065\u006c\u0064\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0041 o\u0072\u0020\u0041\u0041\u0020\u006b\u0065\u0079\u0073\u002e")
		}
	}
	return _cb
}
func _febb(_gbd *_gg.Document) (*_cgd.PdfObjectDictionary, bool) {
	_aaf, _edfe := _gbd.FindCatalog()
	if !_edfe {
		return nil, false
	}
	_deaa, _edfe := _cgd.GetArray(_aaf.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_edfe {
		return nil, false
	}
	if _deaa.Len() == 0 {
		return nil, false
	}
	return _cgd.GetDict(_deaa.Get(0))
}
func _cfeg(_fgc *_gg.Document) {
	_gaaa, _dcc := _fgc.FindCatalog()
	if !_dcc {
		return
	}
	_ddec, _dcc := _gaaa.GetMarkInfo()
	if !_dcc {
		_ddec = _cgd.MakeDict()
	}
	_bfdeb, _dcc := _cgd.GetBool(_ddec.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
	if !_dcc || !bool(*_bfdeb) {
		_ddec.Set("\u004d\u0061\u0072\u006b\u0065\u0064", _cgd.MakeBool(true))
		_gaaa.SetMarkInfo(_ddec)
	}
}
func (_fad *documentImages) hasUncalibratedImages() bool { return _fad._fde || _fad._bg || _fad._bgd }

// StandardName gets the name of the standard.
func (_eabb *profile2) StandardName() string {
	return _a.Sprintf("\u0050D\u0046\u002f\u0041\u002d\u0032\u0025s", _eabb._geg._fa)
}
func _deba(_dgfge *_f.CompliancePdfReader) (_fffd []ViolatedRule) {
	var _dgce, _bgca, _gcdf, _abec, _gbdcd, _bdeeg, _bedcd bool
	_cdce := func() bool { return _dgce && _bgca && _gcdf && _abec && _gbdcd && _bdeeg && _bedcd }
	for _, _ddae := range _dgfge.PageList {
		_fbcg, _accdc := _ddae.GetAnnotations()
		if _accdc != nil {
			_ec.Log.Trace("\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006es\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _accdc)
			continue
		}
		for _, _addg := range _fbcg {
			if !_dgce {
				switch _addg.GetContext().(type) {
				case *_f.PdfAnnotationFileAttachment, *_f.PdfAnnotationSound, *_f.PdfAnnotationMovie, nil:
					_fffd = append(_fffd, _cc("\u0036.\u0035\u002e\u0032\u002d\u0031", "\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0079\u0070\u0065\u0073\u0020\u006e\u006f\u0074 \u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020p\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e\u0020\u0041d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020\u0074\u0068\u0065\u0020F\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063\u0068\u006de\u006e\u0074\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u0020\u0061\u006e\u0064\u0020\u004d\u006f\u0076\u0069e\u0020\u0074\u0079\u0070\u0065s \u0073ha\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_dgce = true
					if _cdce() {
						return _fffd
					}
				}
			}
			_cfaa, _aegc := _cgd.GetDict(_addg.GetContainingPdfObject())
			if !_aegc {
				continue
			}
			if !_bgca {
				_degb, _bafd := _cgd.GetFloatVal(_cfaa.Get("\u0043\u0041"))
				if _bafd && _degb != 1.0 {
					_fffd = append(_fffd, _cc("\u0036.\u0035\u002e\u0033\u002d\u0031", "\u0041\u006e\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073h\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0043\u0041\u0020\u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0031\u002e\u0030\u002e"))
					_bgca = true
					if _cdce() {
						return _fffd
					}
				}
			}
			if !_gcdf {
				_fdff, _fcaec := _cgd.GetIntVal(_cfaa.Get("\u0046"))
				if !(_fcaec && _fdff&4 == 4 && _fdff&1 == 0 && _fdff&2 == 0 && _fdff&32 == 0) {
					_fffd = append(_fffd, _cc("\u0036.\u0035\u002e\u0033\u002d\u0032", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0074\u0068\u0065\u0020\u0046\u0020\u006b\u0065\u0079\u002e\u0020\u0054\u0068\u0065\u0020\u0046\u0020\u006b\u0065\u0079\u0027\u0073\u0020\u0050\u0072\u0069\u006e\u0074\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065 s\u0065\u0074\u0020\u0074\u006f\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0048\u0069\u0064\u0064\u0065\u006e\u002c\u0020I\u006e\u0076\u0069\u0073\u0069\u0062\u006c\u0065\u0020\u0061\u006e\u0064\u0020\u004e\u006f\u0056\u0069\u0065\u0077\u0020\u0066\u006c\u0061\u0067\u0020b\u0069\u0074\u0073 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073e\u0074\u0020t\u006f\u0020\u0030\u002e"))
					_gcdf = true
					if _cdce() {
						return _fffd
					}
				}
			}
			if !_abec {
				_bfaecd, _abad := _cgd.GetDict(_cfaa.Get("\u0041\u0050"))
				if _abad {
					_ebbf := _bfaecd.Get("\u004e")
					if _ebbf == nil || len(_bfaecd.Keys()) > 1 {
						_fffd = append(_fffd, _cc("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_abec = true
						if _cdce() {
							return _fffd
						}
						continue
					}
					_, _efde := _addg.GetContext().(*_f.PdfAnnotationWidget)
					if _efde {
						_afgf, _efgeg := _cgd.GetName(_cfaa.Get("\u0046\u0054"))
						if _efgeg && *_afgf == "\u0042\u0074\u006e" {
							if _, _gcbg := _cgd.GetDict(_ebbf); !_gcbg {
								_fffd = append(_fffd, _cc("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
								_abec = true
								if _cdce() {
									return _fffd
								}
								continue
							}
						}
					}
					_, _abbfg := _cgd.GetStream(_ebbf)
					if !_abbfg {
						_fffd = append(_fffd, _cc("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_abec = true
						if _cdce() {
							return _fffd
						}
						continue
					}
				}
			}
			if !_gbdcd {
				if _cfaa.Get("\u0043") != nil || _cfaa.Get("\u0049\u0043") != nil {
					_ecddf, _eaa := _bbaba(_dgfge)
					if !_eaa {
						_fffd = append(_fffd, _cc("\u0036.\u0035\u002e\u0033\u002d\u0033", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074a\u0069\u006e\u0020t\u0068e\u0020\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0072\u0020\u0074\u0068e\u0020\u0049\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0075\u006e\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006ff\u0069\u006ce\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069n\u0020\u0036\u002e\u0032\u002e2\u002c\u0020\u0069\u0073\u0020\u0052\u0047\u0042."))
						_gbdcd = true
						if _cdce() {
							return _fffd
						}
					} else {
						_daeb, _dfage := _cgd.GetIntVal(_ecddf.Get("\u004e"))
						if !_dfage || _daeb != 3 {
							_fffd = append(_fffd, _cc("\u0036.\u0035\u002e\u0033\u002d\u0033", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074a\u0069\u006e\u0020t\u0068e\u0020\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0072\u0020\u0074\u0068e\u0020\u0049\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0075\u006e\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006ff\u0069\u006ce\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069n\u0020\u0036\u002e\u0032\u002e2\u002c\u0020\u0069\u0073\u0020\u0052\u0047\u0042."))
							_gbdcd = true
							if _cdce() {
								return _fffd
							}
						}
					}
				}
			}
			_ebde, _defd := _addg.GetContext().(*_f.PdfAnnotationWidget)
			if !_defd {
				continue
			}
			if !_bdeeg {
				if _ebde.A != nil {
					_fffd = append(_fffd, _cc("\u0036.\u0036\u002e\u0031\u002d\u0033", "A \u0057\u0069d\u0067\u0065\u0074\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006ec\u006cu\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0020e\u006et\u0072\u0079."))
					_bdeeg = true
					if _cdce() {
						return _fffd
					}
				}
			}
			if !_bedcd {
				if _ebde.AA != nil {
					_fffd = append(_fffd, _cc("\u0036.\u0036\u002e\u0032\u002d\u0031", "\u0041\u0020\u0057\u0069\u0064\u0067\u0065\u0074\u0020\u0061\u006e\u006eo\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073h\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					_bedcd = true
					if _cdce() {
						return _fffd
					}
				}
			}
		}
	}
	return _fffd
}
func _fdba(_dgaf *_f.CompliancePdfReader, _aeb standardType, _adcf bool) (_bgcb []ViolatedRule) {
	_deadf, _gcdc := _ecgd(_dgaf)
	if !_gcdc {
		return []ViolatedRule{_cc("\u0036.\u0037\u002e\u0032\u002d\u0031", "\u0063a\u0074a\u006c\u006f\u0067\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064\u002e")}
	}
	_gadfg := _deadf.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	if _gadfg == nil {
		return []ViolatedRule{_cc("\u0036.\u0037\u002e\u0032\u002d\u0031", "\u006e\u006f\u0020\u0027\u004d\u0065\u0074\u0061d\u0061\u0074\u0061' \u006b\u0065\u0079\u0020\u0066\u006fu\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u002e"), _cc("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e")}
	}
	_efbaf, _gcdc := _cgd.GetStream(_gadfg)
	if !_gcdc {
		return []ViolatedRule{_cc("\u0036.\u0037\u002e\u0032\u002d\u0032", "\u0063\u0061\u0074a\u006c\u006f\u0067\u0020\u0027\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0027\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020a\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"), _cc("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e")}
	}
	if _efbaf.Get("\u0046\u0069\u006c\u0074\u0065\u0072") != nil {
		_bgcb = append(_bgcb, _cc("\u0036.\u0037\u002e\u0032\u002d\u0032", "M\u0065\u0074a\u0064\u0061\u0074\u0061\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u006b\u0065y\u002e"))
	}
	_bgggf, _ggcb := _cg.LoadDocument(_efbaf.Stream)
	if _ggcb != nil {
		return []ViolatedRule{_cc("\u0036.\u0037\u002e\u0039\u002d\u0031", "The\u0020\u006d\u0065\u0074a\u0064\u0061t\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0061\u006e\u0064\u0020\u0077\u0065\u006c\u006c\u0020\u0066\u006f\u0072\u006de\u0064\u0020\u0050\u0044\u0046\u0041\u0045\u0078\u0074e\u006e\u0073\u0069\u006f\u006e\u0020\u0053\u0063\u0068\u0065\u006da\u0020\u0066\u006fr\u0020\u0061\u006c\u006c\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073\u002e")}
	}
	_bfgcd := _bgggf.GetGoXmpDocument()
	var _fab []*_aca.Namespace
	for _, _eabf := range _bfgcd.Namespaces() {
		switch _eabf.Name {
		case _be.NsDc.Name, _g.NsPDF.Name, _ee.NsXmp.Name, _ae.NsXmpRights.Name, _da.Namespace.Name, _df.Namespace.Name, _bf.NsXmpMM.Name, _df.FieldNS.Name, _df.SchemaNS.Name, _df.PropertyNS.Name, "\u0073\u0074\u0045v\u0074", "\u0073\u0074\u0056e\u0072", "\u0073\u0074\u0052e\u0066", "\u0073\u0074\u0044i\u006d", "\u0078a\u0070\u0047\u0049\u006d\u0067", "\u0073\u0074\u004ao\u0062", "\u0078\u006d\u0070\u0069\u0064\u0071":
			continue
		}
		_fab = append(_fab, _eabf)
	}
	_dbfg := true
	_baabg, _ggcb := _bgggf.GetPdfaExtensionSchemas()
	if _ggcb == nil {
		for _, _bbfg := range _fab {
			var _feae bool
			for _bagded := range _baabg {
				if _bbfg.URI == _baabg[_bagded].NamespaceURI {
					_feae = true
					break
				}
			}
			if !_feae {
				_dbfg = false
				break
			}
		}
	} else {
		_dbfg = false
	}
	if !_dbfg {
		_bgcb = append(_bgcb, _cc("\u0036.\u0037\u002e\u0039\u002d\u0032", "\u0050\u0072\u006f\u0070\u0065\u0072\u0074i\u0065\u0073 \u0073\u0070\u0065\u0063\u0069\u0066\u0069ed\u0020\u0069\u006e\u0020\u0058M\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0073\u0068\u0061\u006cl\u0020\u0075\u0073\u0065\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073 \u0064\u0065\u0066i\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006fn\u002c\u0020\u006f\u0072\u0020\u0065\u0078\u0074\u0065ns\u0069\u006f\u006e\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u0074\u0068\u0061\u0074 \u0063\u006f\u006d\u0070\u006c\u0079\u0020\u0077\u0069\u0074h\u0020\u0058\u004d\u0050\u0020\u0053\u0070e\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002e"))
	}
	_abcf, _ggcb := _dgaf.GetPdfInfo()
	if _ggcb == nil {
		if !_dgca(_abcf, _bgggf) {
			_bgcb = append(_bgcb, _cc("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e"))
		}
	} else if _, _eeege := _bgggf.GetMediaManagement(); _eeege {
		_bgcb = append(_bgcb, _cc("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e"))
	}
	_cdecc, _gcdc := _bgggf.GetPdfAID()
	if !_gcdc {
		_bgcb = append(_bgcb, _cc("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0061n\u0064\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020\u006c\u0065\u0076\u0065l\u0020\u006f\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0070e\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073\u0063h\u0065\u006da."))
	} else {
		if _cdecc.Part != _aeb._ba {
			_bgcb = append(_bgcb, _cc("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0032", "\u0054h\u0065\u0020\u0076\u0061lue\u0020\u006f\u0066\u0020p\u0064\u0066\u0061\u0069\u0064\u003a\u0070\u0061\u0072\u0074 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0072\u0074\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u002019\u0030\u0030\u0035 \u0074\u006f\u0020\u0077\u0068i\u0063h\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065 \u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0073\u002e"))
		}
		if _aeb._fa == "\u0041" && _cdecc.Conformance != "\u0041" {
			_bgcb = append(_bgcb, _cc("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063i\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063o\u006e\u0066\u006fr\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0041\u002e\u0020\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0042\u0020\u0063\u006f\u006e\u0066o\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e"))
		} else if _aeb._fa == "\u0042" && (_cdecc.Conformance != "\u0041" && _cdecc.Conformance != "\u0042") {
			_bgcb = append(_bgcb, _cc("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063i\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063o\u006e\u0066\u006fr\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0041\u002e\u0020\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0042\u0020\u0063\u006f\u006e\u0066o\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e"))
		}
	}
	return _bgcb
}
func _eag(_fcbb bool, _eadf standardType) (pageColorspaceOptimizeFunc, documentColorspaceOptimizeFunc) {
	var _abe, _gea, _dag bool
	_bgb := func(_feeb *_gg.Document, _cbb *_gg.Page, _caeb []*_gg.Image) error {
		for _, _aab := range _caeb {
			switch _aab.Colorspace {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
				_gea = true
			case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
				_abe = true
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				_dag = true
			}
		}
		_ggba, _bdc := _cbb.GetContents()
		if !_bdc {
			return nil
		}
		for _, _dge := range _ggba {
			_aba, _cbac := _dge.GetData()
			if _cbac != nil {
				continue
			}
			_adf := _ff.NewContentStreamParser(string(_aba))
			_edb, _cbac := _adf.Parse()
			if _cbac != nil {
				continue
			}
			for _, _gecf := range *_edb {
				switch _gecf.Operand {
				case "\u0047", "\u0067":
					_gea = true
				case "\u0052\u0047", "\u0072\u0067":
					_abe = true
				case "\u004b", "\u006b":
					_dag = true
				case "\u0043\u0053", "\u0063\u0073":
					if len(_gecf.Params) == 0 {
						continue
					}
					_ecdg, _febf := _cgd.GetName(_gecf.Params[0])
					if !_febf {
						continue
					}
					switch _ecdg.String() {
					case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
						_abe = true
					case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
						_gea = true
					case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
						_dag = true
					}
				}
			}
		}
		_cbcg := _cbb.FindXObjectForms()
		for _, _cgfd := range _cbcg {
			_aea := _ff.NewContentStreamParser(string(_cgfd.Stream))
			_ceb, _ggg := _aea.Parse()
			if _ggg != nil {
				continue
			}
			for _, _gfad := range *_ceb {
				switch _gfad.Operand {
				case "\u0047", "\u0067":
					_gea = true
				case "\u0052\u0047", "\u0072\u0067":
					_abe = true
				case "\u004b", "\u006b":
					_dag = true
				case "\u0043\u0053", "\u0063\u0073":
					if len(_gfad.Params) == 0 {
						continue
					}
					_adb, _dbc := _cgd.GetName(_gfad.Params[0])
					if !_dbc {
						continue
					}
					switch _adb.String() {
					case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
						_abe = true
					case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
						_gea = true
					case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
						_dag = true
					}
				}
			}
			_bcaa, _fef := _cgd.GetArray(_cbb.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
			if !_fef {
				return nil
			}
			for _, _edcg := range _bcaa.Elements() {
				_agecd, _gag := _cgd.GetDict(_edcg)
				if !_gag {
					continue
				}
				_dff := _agecd.Get("\u0043")
				if _dff == nil {
					continue
				}
				_eebc, _gag := _cgd.GetArray(_dff)
				if !_gag {
					continue
				}
				switch _eebc.Len() {
				case 0:
				case 1:
					_gea = true
				case 3:
					_abe = true
				case 4:
					_dag = true
				}
			}
		}
		return nil
	}
	_afgd := func(_gbc *_gg.Document, _bee []*_gg.Image) error {
		_cfd, _eda := _gbc.FindCatalog()
		if !_eda {
			return nil
		}
		_def, _eda := _cfd.GetOutputIntents()
		if _eda && _def.Len() > 0 {
			return nil
		}
		if !_eda {
			_def = _cfd.NewOutputIntents()
		}
		if !(_abe || _dag || _gea) {
			return nil
		}
		defer _cfd.SetOutputIntents(_def)
		if _abe && !_dag && !_gea {
			return _fcaf(_gbc, _eadf, _def)
		}
		if _dag && !_abe && !_gea {
			return _acb(_eadf, _def)
		}
		if _gea && !_abe && !_dag {
			return _ged(_eadf, _def)
		}
		if _abe && _dag {
			if _abcc := _ea(_bee, _fcbb); _abcc != nil {
				return _abcc
			}
			if _gfec := _aaga(_gbc, _fcbb); _gfec != nil {
				return _gfec
			}
			if _cef := _adcg(_gbc, _fcbb); _cef != nil {
				return _cef
			}
			if _afd := _eab(_gbc, _fcbb); _afd != nil {
				return _afd
			}
			if _fcbb {
				return _acb(_eadf, _def)
			}
			return _fcaf(_gbc, _eadf, _def)
		}
		return nil
	}
	return _bgb, _afgd
}
func _gde(_dfbb *_gg.Document) {
	if _dfbb.ID[0] != "" && _dfbb.ID[1] != "" {
		return
	}
	_dfbb.UseHashBasedID = true
}

// NewProfile1B creates a new Profile1B with the given options.
func NewProfile1B(options *Profile1Options) *Profile1B {
	if options == nil {
		options = DefaultProfile1Options()
	}
	_beaa(options)
	return &Profile1B{profile1{_beb: *options, _acfg: _gd()}}
}
func _gebdb(_eaba *_f.CompliancePdfReader) (_cdbedc []ViolatedRule) {
	_agdbc := _eaba.GetObjectNums()
	for _, _acg := range _agdbc {
		_bcdf, _dgbc := _eaba.GetIndirectObjectByNumber(_acg)
		if _dgbc != nil {
			continue
		}
		_fafb, _gacd := _cgd.GetDict(_bcdf)
		if !_gacd {
			continue
		}
		_eabfg, _gacd := _cgd.GetName(_fafb.Get("\u0054\u0079\u0070\u0065"))
		if !_gacd {
			continue
		}
		if _eabfg.String() != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
			continue
		}
		if _fafb.Get("\u0045\u0046") != nil {
			if _fafb.Get("\u0046") == nil || _fafb.Get("\u0045\u0046") == nil {
				_cdbedc = append(_cdbedc, _cc("\u0036\u002e\u0038-\u0032", "\u0054h\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066i\u0063\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063t\u0069\u006fn\u0061\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020t\u0068\u0065\u0020\u0046\u0020a\u006e\u0064\u0020\u0055\u0046\u0020\u006b\u0065\u0079\u0073\u002e"))
			}
			if _fafb.Get("\u0041\u0046\u0052\u0065\u006c\u0061\u0074\u0069\u006fn\u0073\u0068\u0069\u0070") == nil {
				_cdbedc = append(_cdbedc, _cc("\u0036\u002e\u0038-\u0033", "\u0049\u006e\u0020\u006f\u0072d\u0065\u0072\u0020\u0074\u006f\u0020\u0065\u006e\u0061\u0062\u006c\u0065\u0020i\u0064\u0065nt\u0069\u0066\u0069c\u0061\u0074\u0069o\u006e\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0074\u0068\u0065\u0020fi\u006ce\u0020\u0073\u0070\u0065\u0063\u0069f\u0069c\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020c\u006f\u006e\u0074e\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0069\u0074\u002c\u0020\u0061\u0020\u006e\u0065\u0077\u0020(\u0072\u0065\u0071\u0075i\u0072\u0065\u0064\u0029\u0020\u006be\u0079\u0020h\u0061\u0073\u0020\u0062e\u0065\u006e\u0020\u0064\u0065\u0066i\u006e\u0065\u0064\u0020a\u006e\u0064\u0020\u0069\u0074s \u0070\u0072e\u0073\u0065n\u0063\u0065\u0020\u0028\u0069\u006e\u0020\u0074\u0068e\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u0029\u0020\u0069\u0073\u0020\u0072\u0065q\u0075\u0069\u0072e\u0064\u002e"))
			}
			break
		}
	}
	return _cdbedc
}

// StandardName gets the name of the standard.
func (_daaf *profile1) StandardName() string {
	return _a.Sprintf("\u0050D\u0046\u002f\u0041\u002d\u0031\u0025s", _daaf._acfg._fa)
}
func _cggdd(_geebc *_f.CompliancePdfReader) ViolatedRule {
	if _geebc.ParserMetadata().HasDataAfterEOF() {
		return _cc("\u0036.\u0031\u002e\u0033\u002d\u0033", "\u004e\u006f\u0020\u0064\u0061ta\u0020\u0073h\u0061\u006c\u006c\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0020\u0074\u0068\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0065\u006e\u0064\u002d\u006f\u0066\u002d\u0066\u0069l\u0065\u0020\u006da\u0072\u006b\u0065\u0072\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0061 \u0073\u0069\u006e\u0067\u006ce\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061\u006c \u0065\u006ed\u002do\u0066\u002d\u006c\u0069\u006e\u0065\u0020m\u0061\u0072\u006b\u0065\u0072\u002e")
	}
	return _cb
}
func _ccbf(_bbggg *_f.CompliancePdfReader) (_ddbd []ViolatedRule) {
	var _gbed, _geege, _aac, _eedd, _babg, _bdfg bool
	_cfbd := func() bool { return _gbed && _geege && _aac && _eedd && _babg && _bdfg }
	for _, _fafcb := range _bbggg.PageList {
		if _fafcb.Resources == nil {
			continue
		}
		_gcdaa, _efdcd := _cgd.GetDict(_fafcb.Resources.Font)
		if !_efdcd {
			continue
		}
		for _, _daed := range _gcdaa.Keys() {
			_ababb, _fcbc := _cgd.GetDict(_gcdaa.Get(_daed))
			if !_fcbc {
				if !_gbed {
					_ddbd = append(_ddbd, _cc("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0031", "\u0041\u006c\u006c\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006e\u0064\u0020\u0066on\u0074 \u0070\u0072\u006fg\u0072\u0061\u006ds\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072mi\u006e\u0067\u0020\u0066\u0069\u006ce\u002c\u0020\u0072\u0065\u0067\u0061\u0072\u0064\u006c\u0065s\u0073\u0020\u006f\u0066\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006eg m\u006f\u0064\u0065\u0020\u0075\u0073\u0061\u0067\u0065\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0074\u0068e\u0020\u0070\u0072o\u0076\u0069\u0073\u0069\u006f\u006e\u0073\u0020\u0069\u006e \u0049\u0053\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031:\u0032\u0030\u0030\u0038\u002c \u0039\u002e\u0036\u0020a\u006e\u0064\u0020\u0039.\u0037\u002e"))
					_gbed = true
					if _cfbd() {
						return _ddbd
					}
				}
				continue
			}
			if _adffc, _ddad := _cgd.GetName(_ababb.Get("\u0054\u0079\u0070\u0065")); !_gbed && (!_ddad || _adffc.String() != "\u0046\u006f\u006e\u0074") {
				_ddbd = append(_ddbd, _cc("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0031", "\u0054\u0079\u0070e\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029 Th\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066 \u0050\u0044\u0046\u0020\u006fbj\u0065\u0063\u0074\u0020\u0074\u0068\u0061t\u0020\u0074\u0068\u0069s\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0064\u0065\u0073c\u0072\u0069\u0062\u0065\u0073\u003b\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0046\u006f\u006e\u0074\u0020\u0066\u006fr\u0020\u0061\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
				_gbed = true
				if _cfbd() {
					return _ddbd
				}
			}
			_eeef, _ddbgc := _f.NewPdfFontFromPdfObject(_ababb)
			if _ddbgc != nil {
				continue
			}
			var _fgcd string
			if _ecege, _bfffa := _cgd.GetName(_ababb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _bfffa {
				_fgcd = _ecege.String()
			}
			if !_geege {
				switch _fgcd {
				case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
				default:
					_geege = true
					_ddbd = append(_ddbd, _cc("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0032", "\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065d\u0029\u0020\u0054\u0068e \u0074\u0079\u0070\u0065 \u006f\u0066\u0020\u0066\u006f\u006et\u003b\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u0022\u0054\u0079\u0070\u0065\u0031\u0022\u0020f\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020f\u006f\u006e\u0074\u0073\u002c\u0020\u0022\u004d\u004d\u0054\u0079\u0070\u0065\u0031\u0022\u0020\u0066\u006f\u0072\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006da\u0073\u0074e\u0072\u0020\u0066\u006f\u006e\u0074s\u002c\u0020\u0022\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0022\u0054\u0079\u0070\u0065\u0033\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070e\u0020\u0033\u0020\u0066\u006f\u006e\u0074\u0073\u002c\u0020\"\u0054\u0079\u0070\u0065\u0030\"\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006ed\u0020\u0022\u0043\u0049\u0044\u0046\u006fn\u0074\u0054\u0079\u0070\u0065\u0030\u0022 \u006f\u0072\u0020\u0022\u0043\u0049\u0044\u0046\u006f\u006e\u0074T\u0079\u0070e\u0032\u0022\u0020\u0066\u006f\u0072\u0020\u0043\u0049\u0044\u0020\u0066\u006f\u006e\u0074\u0073\u002e"))
					if _cfbd() {
						return _ddbd
					}
				}
			}
			if !_aac {
				if _fgcd != "\u0054\u0079\u0070e\u0033" {
					_cddgf, _efaa := _cgd.GetName(_ababb.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
					if !_efaa || _cddgf.String() == "" {
						_ddbd = append(_ddbd, _cc("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0033", "B\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064)\u0020T\u0068\u0065\u0020\u0050o\u0073\u0074S\u0063\u0072\u0069\u0070\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e"))
						_aac = true
						if _cfbd() {
							return _ddbd
						}
					}
				}
			}
			if _fgcd != "\u0054\u0079\u0070e\u0031" {
				continue
			}
			_cdfe := _ecb.IsStdFont(_ecb.StdFontName(_eeef.BaseFont()))
			if _cdfe {
				continue
			}
			_eacgf, _gffa := _cgd.GetIntVal(_ababb.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
			if !_gffa && !_eedd {
				_ddbd = append(_ddbd, _cc("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0034", "\u0046\u0069r\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072d\u0020\u0031\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u0029\u0020\u0054\u0068\u0065\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064e\u0020\u0064\u0065\u0066i\u006ee\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057i\u0064\u0074\u0068\u0073 \u0061r\u0072\u0061y\u002e"))
				_eedd = true
				if _cfbd() {
					return _ddbd
				}
			}
			_effc, _cbgf := _cgd.GetIntVal(_ababb.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
			if !_cbgf && !_babg {
				_ddbd = append(_ddbd, _cc("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0035", "\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069n\u0074\u0065\u0067e\u0072 \u002d\u0020\u0028\u0052\u0065\u0071u\u0069\u0072\u0065d\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020t\u0068\u0065 s\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0031\u0034\u0020\u0066\u006f\u006ets\u0029\u0020\u0054\u0068\u0065\u0020\u006c\u0061\u0073t\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057\u0069\u0064\u0074h\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u002e"))
				_babg = true
				if _cfbd() {
					return _ddbd
				}
			}
			if !_bdfg {
				_ggaed, _fabd := _cgd.GetArray(_ababb.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
				if !_fabd || !_gffa || !_cbgf || _ggaed.Len() != _effc-_eacgf+1 {
					_ddbd = append(_ddbd, _cc("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0036", "\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u002d a\u0072\u0072\u0061y \u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0074a\u006e\u0064a\u0072\u0064\u00201\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u003b\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0070\u0072\u0065\u0066e\u0072\u0072e\u0064\u0029\u0020\u0041\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0066\u0020\u0028\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u2212 F\u0069\u0072\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u002b\u00201\u0029\u0020\u0077\u0069\u0064\u0074\u0068\u0073."))
					_bdfg = true
					if _cfbd() {
						return _ddbd
					}
				}
			}
		}
	}
	return _ddbd
}
func _gdgaf(_ebag *_f.CompliancePdfReader) (*_f.PdfOutputIntent, bool) {
	_egbc, _dcgdg := _bbaba(_ebag)
	if !_dcgdg {
		return nil, false
	}
	_cagc, _abbfga := _f.NewPdfOutputIntentFromPdfObject(_egbc)
	if _abbfga != nil {
		return nil, false
	}
	return _cagc, true
}
func _ebced(_dbge *_f.CompliancePdfReader) (_efeef []ViolatedRule) {
	_fcgeb := func(_bccbf *_cgd.PdfObjectDictionary, _feebgd *[]string, _bgfce *[]ViolatedRule) error {
		_fadb := _bccbf.Get("\u004e\u0061\u006d\u0065")
		if _fadb == nil || len(_fadb.String()) == 0 {
			*_bgfce = append(*_bgfce, _cc("\u0036\u002e\u0039-\u0031", "\u0045\u0061\u0063\u0068\u0020o\u0070\u0074\u0069\u006f\u006e\u0061l\u0020\u0063\u006f\u006e\u0074\u0065\u006et\u0020\u0063\u006fn\u0066\u0069\u0067\u0075r\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004e\u0061\u006d\u0065\u0020\u006b\u0065\u0079\u002e"))
		}
		for _, _bgbd := range *_feebgd {
			if _bgbd == _fadb.String() {
				*_bgfce = append(*_bgfce, _cc("\u0036\u002e\u0039-\u0032", "\u0045\u0061\u0063\u0068\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061l\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0063\u006f\u006e\u0066\u0069\u0067\u0075\u0072a\u0074\u0069\u006fn\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020N\u0061\u006d\u0065\u0020\u006b\u0065\u0079\u002c w\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020s\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0075ni\u0071\u0075\u0065 \u0061\u006d\u006f\u006e\u0067\u0073\u0074\u0020\u0061\u006c\u006c\u0020o\u0070\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0063\u006fn\u0074\u0065\u006e\u0074 \u0063\u006f\u006e\u0066\u0069\u0067u\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074i\u006fn\u0061\u0072\u0069\u0065\u0073\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0074\u0068e\u0020\u0050\u0044\u0046\u002fA\u002d\u0032\u0020\u0066\u0069l\u0065\u002e"))
			} else {
				*_feebgd = append(*_feebgd, _fadb.String())
			}
		}
		if _bccbf.Get("\u0041\u0053") != nil {
			*_bgfce = append(*_bgfce, _cc("\u0036\u002e\u0039-\u0034", "Th\u0065\u0020\u0041\u0053\u0020\u006b\u0065y \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0070\u0070\u0065\u0061r\u0020\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0020\u0063\u006fn\u0066\u0069\u0067\u0075\u0072\u0061\u0074\u0069\u006fn\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
		}
		return nil
	}
	_efff, _aaeac := _ecgd(_dbge)
	if !_aaeac {
		return _efeef
	}
	_gbda, _aaeac := _cgd.GetDict(_efff.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"))
	if !_aaeac {
		return _efeef
	}
	var _begdf []string
	_edbdg, _aaeac := _cgd.GetDict(_gbda.Get("\u0044"))
	if _aaeac {
		_fcgeb(_edbdg, &_begdf, &_efeef)
	}
	_eeagae, _aaeac := _cgd.GetArray(_gbda.Get("\u0043o\u006e\u0066\u0069\u0067\u0073"))
	if _aaeac {
		for _eggfg := 0; _eggfg < _eeagae.Len(); _eggfg++ {
			_gdaga, _ceggc := _cgd.GetDict(_eeagae.Get(_eggfg))
			if !_ceggc {
				continue
			}
			_fcgeb(_gdaga, &_begdf, &_efeef)
		}
	}
	return _efeef
}
func _cddgb(_cdbd *_f.CompliancePdfReader) (_dcce ViolatedRule) {
	_edgfa, _cgaf := _ecgd(_cdbd)
	if !_cgaf {
		return _cb
	}
	if _edgfa.Get("\u0041\u0041") != nil {
		return _cc("\u0036.\u0035\u002e\u0032\u002d\u0031", "\u0054h\u0065\u0020\u0064\u006fc\u0075m\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020a\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079 \u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079\u002e")
	}
	return _cb
}
func _dfbbg(_adg *_f.CompliancePdfReader) ViolatedRule {
	_dad := _adg.ParserMetadata().HeaderCommentBytes()
	if _dad[0] > 127 && _dad[1] > 127 && _dad[2] > 127 && _dad[3] > 127 {
		return _cb
	}
	return _cc("\u0036.\u0031\u002e\u0032\u002d\u0032", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u006c\u0069\u006e\u0065\u0020\u0073\u0068\u0061\u006c\u006c b\u0065\u0020i\u006d\u006d\u0065\u0064\u0069a\u0074\u0065\u006c\u0079 \u0066\u006f\u006c\u006co\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0063\u006f\u006d\u006d\u0065n\u0074\u0020\u0063\u006f\u006e\u0073\u0069s\u0074\u0069\u006e\u0067\u0020o\u0066\u0020\u0061\u0020\u0025\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006fwe\u0064\u0020\u0062y\u0020a\u0074\u0009\u006c\u0065a\u0073\u0074\u0020f\u006f\u0075\u0072\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u002c\u0020e\u0061\u0063\u0068\u0020\u006f\u0066\u0020\u0077\u0068\u006f\u0073\u0065 \u0065\u006e\u0063\u006f\u0064e\u0064\u0020\u0062\u0079\u0074e\u0020\u0076\u0061\u006c\u0075\u0065s\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0064e\u0063\u0069\u006d\u0061\u006c \u0076\u0061\u006c\u0075\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0031\u0032\u0037\u002e")
}
func _fc(_fea []_cgd.PdfObject) (*documentImages, error) {
	_ga := _cgd.PdfObjectName("\u0053u\u0062\u0074\u0079\u0070\u0065")
	_dae := make(map[*_cgd.PdfObjectStream]struct{})
	_gfg := make(map[_cgd.PdfObject]struct{})
	var (
		_ab, _ad, _ffb bool
		_fadg          []*imageInfo
		_dfc           error
	)
	for _, _gge := range _fea {
		_fb, _dc := _cgd.GetStream(_gge)
		if !_dc {
			continue
		}
		if _, _cgde := _dae[_fb]; _cgde {
			continue
		}
		_dae[_fb] = struct{}{}
		_fbc := _fb.PdfObjectDictionary.Get(_ga)
		_cgdf, _dc := _cgd.GetName(_fbc)
		if !_dc || string(*_cgdf) != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if _fbf := _fb.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b"); _fbf != nil {
			_gfg[_fbf] = struct{}{}
		}
		_cbg := imageInfo{BitsPerComponent: 8, Stream: _fb}
		_cbg.ColorSpace, _dfc = _f.DetermineColorspaceNameFromPdfObject(_fb.PdfObjectDictionary.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"))
		if _dfc != nil {
			return nil, _dfc
		}
		if _egb, _ege := _cgd.GetIntVal(_fb.PdfObjectDictionary.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _ege {
			_cbg.BitsPerComponent = _egb
		}
		if _faf, _fbb := _cgd.GetIntVal(_fb.PdfObjectDictionary.Get("\u0057\u0069\u0064t\u0068")); _fbb {
			_cbg.Width = _faf
		}
		if _feb, _afc := _cgd.GetIntVal(_fb.PdfObjectDictionary.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _afc {
			_cbg.Height = _feb
		}
		switch _cbg.ColorSpace {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			_ffb = true
			_cbg.ColorComponents = 1
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			_ab = true
			_cbg.ColorComponents = 3
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			_ad = true
			_cbg.ColorComponents = 4
		default:
			_cbg._bac = true
		}
		_fadg = append(_fadg, &_cbg)
	}
	if len(_gfg) > 0 {
		if len(_gfg) == len(_fadg) {
			_fadg = nil
		} else {
			_eb := make([]*imageInfo, len(_fadg)-len(_gfg))
			var _eeg int
			for _, _gfgd := range _fadg {
				if _, _bfd := _gfg[_gfgd.Stream]; _bfd {
					continue
				}
				_eb[_eeg] = _gfgd
				_eeg++
			}
			_fadg = _eb
		}
	}
	return &documentImages{_fde: _ab, _bg: _ad, _bgd: _ffb, _bff: _gfg, _fe: _fadg}, nil
}
func _bebd(_geecb *_cgd.PdfObjectDictionary, _feee map[*_cgd.PdfObjectStream][]byte, _gccg map[*_cgd.PdfObjectStream]*_bd.CMap) ViolatedRule {
	const (
		_daff  = "\u0036.\u0033\u002e\u0033\u002d\u0033"
		_efegf = "\u0041\u006cl \u0043\u004d\u0061\u0070\u0073\u0020\u0075\u0073e\u0064 \u0077i\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072m\u0069n\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020a\u006e\u0064\u0020\u0049\u0064\u0065\u006et\u0069\u0074\u0079-\u0056\u002c\u0020\u0073\u0068a\u006c\u006c \u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0074h\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u0020\u0061\u0073\u0020\u0064es\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u00205\u002e\u0036\u002e\u0034\u002e"
	)
	var _aabcg string
	if _fcgb, _cdeag := _cgd.GetName(_geecb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cdeag {
		_aabcg = _fcgb.String()
	}
	if _aabcg != "\u0054\u0079\u0070e\u0030" {
		return _cb
	}
	_cceed := _geecb.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _cedd, _aabb := _cgd.GetName(_cceed); _aabb {
		switch _cedd.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _cb
		default:
			return _cc(_daff, _efegf)
		}
	}
	_dcbb, _aade := _cgd.GetStream(_cceed)
	if !_aade {
		return _cc(_daff, _efegf)
	}
	_, _dgdf := _acbf(_dcbb, _feee, _gccg)
	if _dgdf != nil {
		return _cc(_daff, _efegf)
	}
	return _cb
}
func _gbbge(_eccef *_f.CompliancePdfReader) (_ggbaf []ViolatedRule) {
	var _afce, _gddb, _ffa bool
	if _eccef.ParserMetadata().HasNonConformantStream() {
		_ggbaf = []ViolatedRule{_cc("\u0036.\u0031\u002e\u0037\u002d\u0031", "T\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020f\u006f\u006cl\u006fw\u0065\u0064\u0020e\u0069\u0074h\u0065\u0072\u0020\u0062\u0079\u0020\u0061 \u0043\u0041\u0052\u0052I\u0041\u0047\u0045\u0020\u0052E\u0054\u0055\u0052\u004e\u0020\u00280\u0044\u0068\u0029\u0020\u0061\u006e\u0064\u0020\u004c\u0049\u004e\u0045\u0020F\u0045\u0045\u0044\u0020\u0028\u0030\u0041\u0068\u0029\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065\u0020o\u0072\u0020\u0062\u0079\u0020\u0061 \u0073\u0069ng\u006c\u0065\u0020\u004cIN\u0045 \u0046\u0045\u0045\u0044 \u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u002e\u0020T\u0068\u0065\u0020e\u006e\u0064\u0073\u0074r\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020p\u0072\u0065\u0063\u0065\u0064\u0065\u0064\u0020\u0062\u0079\u0020\u0061n\u0020\u0045\u004f\u004c \u006d\u0061\u0072\u006b\u0065\u0072\u002e")}
	}
	for _, _egdd := range _eccef.GetObjectNums() {
		_eded, _ := _eccef.GetIndirectObjectByNumber(_egdd)
		if _eded == nil {
			continue
		}
		_gdca, _beea := _cgd.GetStream(_eded)
		if !_beea {
			continue
		}
		if !_afce {
			_bgeb := _gdca.Get("\u004c\u0065\u006e\u0067\u0074\u0068")
			if _bgeb == nil {
				_ggbaf = append(_ggbaf, _cc("\u0036.\u0031\u002e\u0037\u002d\u0032", "\u006e\u006f\u0020'\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074"))
				_afce = true
			} else {
				_cdde, _facf := _cgd.GetIntVal(_bgeb)
				if !_facf {
					_ggbaf = append(_ggbaf, _cc("\u0036.\u0031\u002e\u0037\u002d\u0032", "s\u0074\u0072\u0065\u0061\u006d\u0020\u0027\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020an\u0020\u0069\u006et\u0065g\u0065\u0072"))
					_afce = true
				} else {
					if len(_gdca.Stream) != _cdde {
						_ggbaf = append(_ggbaf, _cc("\u0036.\u0031\u002e\u0037\u002d\u0032", "\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006c\u0065\u006e\u0067th\u0020v\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d"))
						_afce = true
					}
				}
			}
		}
		if !_gddb {
			if _gdca.Get("\u0046") != nil {
				_gddb = true
				_ggbaf = append(_ggbaf, _cc("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
			}
			if _gdca.Get("\u0046F\u0069\u006c\u0074\u0065\u0072") != nil && !_gddb {
				_gddb = true
				_ggbaf = append(_ggbaf, _cc("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
			if _gdca.Get("\u0046\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u0061\u006d\u0073") != nil && !_gddb {
				_gddb = true
				_ggbaf = append(_ggbaf, _cc("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
		}
		if !_ffa {
			_cfga, _ggag := _cgd.GetName(_cgd.TraceToDirectObject(_gdca.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
			if !_ggag {
				continue
			}
			if *_cfga == _cgd.StreamEncodingFilterNameLZW {
				_ffa = true
				_ggbaf = append(_ggbaf, _cc("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0031", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e"))
			}
		}
	}
	return _ggbaf
}
func _eeedb(_ggbg *_f.CompliancePdfReader) (_egfa []ViolatedRule) {
	var _cddd, _beed, _caca, _defg, _aegge, _cebc, _cgbg bool
	_abba := func() bool { return _cddd && _beed && _caca && _defg && _aegge && _cebc && _cgbg }
	for _, _ffee := range _ggbg.PageList {
		if _ffee.Resources == nil {
			continue
		}
		_bfgc, _agcc := _cgd.GetDict(_ffee.Resources.Font)
		if !_agcc {
			continue
		}
		for _, _fafe := range _bfgc.Keys() {
			_ced, _agda := _cgd.GetDict(_bfgc.Get(_fafe))
			if !_agda {
				if !_cddd {
					_egfa = append(_egfa, _cc("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0041\u006c\u006c\u0020\u0066\u006fn\u0074\u0073\u0020\u0075\u0073e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0066\u006f\u0072m\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0073\u0020d\u0065\u0066\u0069\u006e\u0065d \u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035\u002e\u0035\u002e"))
					_cddd = true
					if _abba() {
						return _egfa
					}
				}
				continue
			}
			if _bbcb, _eaca := _cgd.GetName(_ced.Get("\u0054\u0079\u0070\u0065")); !_cddd && (!_eaca || _bbcb.String() != "\u0046\u006f\u006e\u0074") {
				_egfa = append(_egfa, _cc("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0054\u0079\u0070e\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029 Th\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066 \u0050\u0044\u0046\u0020\u006fbj\u0065\u0063\u0074\u0020\u0074\u0068\u0061t\u0020\u0074\u0068\u0069s\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0064\u0065\u0073c\u0072\u0069\u0062\u0065\u0073\u003b\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0046\u006f\u006e\u0074\u0020\u0066\u006fr\u0020\u0061\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
				_cddd = true
				if _abba() {
					return _egfa
				}
			}
			_ddbc, _fdc := _f.NewPdfFontFromPdfObject(_ced)
			if _fdc != nil {
				continue
			}
			var _babd string
			if _eege, _cfda := _cgd.GetName(_ced.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cfda {
				_babd = _eege.String()
			}
			if !_beed {
				switch _babd {
				case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
				default:
					_beed = true
					_egfa = append(_egfa, _cc("\u0036.\u0033\u002e\u0032\u002d\u0032", "\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065d\u0029\u0020\u0054\u0068e \u0074\u0079\u0070\u0065 \u006f\u0066\u0020\u0066\u006f\u006et\u003b\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u0022\u0054\u0079\u0070\u0065\u0031\u0022\u0020f\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020f\u006f\u006e\u0074\u0073\u002c\u0020\u0022\u004d\u004d\u0054\u0079\u0070\u0065\u0031\u0022\u0020\u0066\u006f\u0072\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006da\u0073\u0074e\u0072\u0020\u0066\u006f\u006e\u0074s\u002c\u0020\u0022\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0022\u0054\u0079\u0070\u0065\u0033\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070e\u0020\u0033\u0020\u0066\u006f\u006e\u0074\u0073\u002c\u0020\"\u0054\u0079\u0070\u0065\u0030\"\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006ed\u0020\u0022\u0043\u0049\u0044\u0046\u006fn\u0074\u0054\u0079\u0070\u0065\u0030\u0022 \u006f\u0072\u0020\u0022\u0043\u0049\u0044\u0046\u006f\u006e\u0074T\u0079\u0070e\u0032\u0022\u0020\u0066\u006f\u0072\u0020\u0043\u0049\u0044\u0020\u0066\u006f\u006e\u0074\u0073\u002e"))
					if _abba() {
						return _egfa
					}
				}
			}
			if !_caca {
				if _babd != "\u0054\u0079\u0070e\u0033" {
					_adef, _begd := _cgd.GetName(_ced.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
					if !_begd || _adef.String() == "" {
						_egfa = append(_egfa, _cc("\u0036.\u0033\u002e\u0032\u002d\u0033", "B\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064)\u0020T\u0068\u0065\u0020\u0050o\u0073\u0074S\u0063\u0072\u0069\u0070\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e"))
						_caca = true
						if _abba() {
							return _egfa
						}
					}
				}
			}
			if _babd != "\u0054\u0079\u0070e\u0031" {
				continue
			}
			_gafc := _ecb.IsStdFont(_ecb.StdFontName(_ddbc.BaseFont()))
			if _gafc {
				continue
			}
			_gbec, _afe := _cgd.GetIntVal(_ced.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
			if !_afe && !_defg {
				_egfa = append(_egfa, _cc("\u0036.\u0033\u002e\u0032\u002d\u0034", "\u0046\u0069r\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072d\u0020\u0031\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u0029\u0020\u0054\u0068\u0065\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064e\u0020\u0064\u0065\u0066i\u006ee\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057i\u0064\u0074\u0068\u0073 \u0061r\u0072\u0061y\u002e"))
				_defg = true
				if _abba() {
					return _egfa
				}
			}
			_bbdfd, _cdefd := _cgd.GetIntVal(_ced.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
			if !_cdefd && !_aegge {
				_egfa = append(_egfa, _cc("\u0036.\u0033\u002e\u0032\u002d\u0035", "\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069n\u0074\u0065\u0067e\u0072 \u002d\u0020\u0028\u0052\u0065\u0071u\u0069\u0072\u0065d\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020t\u0068\u0065 s\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0031\u0034\u0020\u0066\u006f\u006ets\u0029\u0020\u0054\u0068\u0065\u0020\u006c\u0061\u0073t\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057\u0069\u0064\u0074h\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u002e"))
				_aegge = true
				if _abba() {
					return _egfa
				}
			}
			if !_cebc {
				_deee, _agbe := _cgd.GetArray(_ced.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
				if !_agbe || !_afe || !_cdefd || _deee.Len() != _bbdfd-_gbec+1 {
					_egfa = append(_egfa, _cc("\u0036.\u0033\u002e\u0032\u002d\u0036", "\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u002d a\u0072\u0072\u0061y \u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0074a\u006e\u0064a\u0072\u0064\u00201\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u003b\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0070\u0072\u0065\u0066e\u0072\u0072e\u0064\u0029\u0020\u0041\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0066\u0020\u0028\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u2212 F\u0069\u0072\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u002b\u00201\u0029\u0020\u0077\u0069\u0064\u0074\u0068\u0073."))
					_cebc = true
					if _abba() {
						return _egfa
					}
				}
			}
		}
	}
	return _egfa
}
func _bege(_fedg *_gg.Document) error {
	_bcc, _bfae := _fedg.FindCatalog()
	if !_bfae {
		return _d.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_fbcd, _bfae := _cgd.GetDict(_bcc.Object.Get("\u0050\u0065\u0072m\u0073"))
	if _bfae {
		_feded := _cgd.MakeDict()
		_bdgf := _fbcd.Keys()
		for _, _acaac := range _bdgf {
			if _acaac.String() == "\u0055\u0052\u0033" || _acaac.String() == "\u0044\u006f\u0063\u004d\u0044\u0050" {
				_feded.Set(_acaac, _fbcd.Get(_acaac))
			}
		}
		_bcc.Object.Set("\u0050\u0065\u0072m\u0073", _feded)
	}
	return nil
}
func _fgaa(_begg *_gg.Document) error {
	_dfcdf := func(_bcaaf *_cgd.PdfObjectDictionary) error {
		if _bcaaf.Get("\u0054\u0052") != nil {
			_ec.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0054\u0052\u0020\u006b\u0065\u0079")
			_bcaaf.Remove("\u0054\u0052")
		}
		_dbbd := _bcaaf.Get("\u0054\u0052\u0032")
		if _dbbd != nil {
			_eebce := _dbbd.String()
			if _eebce != "\u0044e\u0066\u0061\u0075\u006c\u0074" {
				_ec.Log.Debug("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074\u0065 o\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073 \u0054\u00522\u0020\u006b\u0065y\u0020\u0077\u0069\u0074\u0068\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065r\u0020\u0074ha\u006e\u0020\u0044e\u0066\u0061\u0075\u006c\u0074")
				_bcaaf.Set("\u0054\u0052\u0032", _cgd.MakeName("\u0044e\u0066\u0061\u0075\u006c\u0074"))
			}
		}
		if _bcaaf.Get("\u0048\u0054\u0050") != nil {
			_ec.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0073\u0020\u0048\u0054P\u0020\u006b\u0065\u0079")
			_bcaaf.Remove("\u0048\u0054\u0050")
		}
		_cfcc := _bcaaf.Get("\u0042\u004d")
		if _cfcc != nil {
			_fdda, _fgabe := _cgd.GetName(_cfcc)
			if !_fgabe {
				_ec.Log.Debug("E\u0078\u0074\u0047\u0053\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0027\u0042\u004d\u0027\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061m\u0065")
				_fdda = _cgd.MakeName("")
			}
			_cegg := _fdda.String()
			switch _cegg {
			case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065", "\u004d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "\u0053\u0063\u0072\u0065\u0065\u006e", "\u004fv\u0065\u0072\u006c\u0061\u0079", "\u0044\u0061\u0072\u006b\u0065\u006e", "\u004ci\u0067\u0068\u0074\u0065\u006e", "\u0043\u006f\u006c\u006f\u0072\u0044\u006f\u0064\u0067\u0065", "\u0043o\u006c\u006f\u0072\u0042\u0075\u0072n", "\u0048a\u0072\u0064\u004c\u0069\u0067\u0068t", "\u0053o\u0066\u0074\u004c\u0069\u0067\u0068t", "\u0044\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065", "\u0045x\u0063\u006c\u0075\u0073\u0069\u006fn", "\u0048\u0075\u0065", "\u0053\u0061\u0074\u0075\u0072\u0061\u0074\u0069\u006f\u006e", "\u0043\u006f\u006co\u0072", "\u004c\u0075\u006d\u0069\u006e\u006f\u0073\u0069\u0074\u0079":
			default:
				_bcaaf.Set("\u0042\u004d", _cgd.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
			}
		}
		return nil
	}
	_dgg, _fdg := _begg.GetPages()
	if !_fdg {
		return nil
	}
	for _, _gdcg := range _dgg {
		_bda, _gbbc := _gdcg.GetResources()
		if !_gbbc {
			continue
		}
		_dbf, _becf := _cgd.GetDict(_bda.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
		if !_becf {
			return nil
		}
		_fcfc := _dbf.Keys()
		for _, _deec := range _fcfc {
			_bbbd, _ebfe := _cgd.GetDict(_dbf.Get(_deec))
			if !_ebfe {
				continue
			}
			_efbe := _dfcdf(_bbbd)
			if _efbe != nil {
				continue
			}
		}
	}
	for _, _cag := range _dgg {
		_ede, _gae := _cag.GetContents()
		if !_gae {
			return nil
		}
		for _, _egga := range _ede {
			_dffb, _cabf := _egga.GetData()
			if _cabf != nil {
				continue
			}
			_beff := _ff.NewContentStreamParser(string(_dffb))
			_dgff, _cabf := _beff.Parse()
			if _cabf != nil {
				continue
			}
			for _, _gefa := range *_dgff {
				if len(_gefa.Params) == 0 {
					continue
				}
				_, _acacb := _cgd.GetName(_gefa.Params[0])
				if !_acacb {
					continue
				}
				_bdcc, _eeec := _cag.GetResourcesXObject()
				if !_eeec {
					continue
				}
				for _, _fceb := range _bdcc.Keys() {
					_eacgd, _fdbf := _cgd.GetStream(_bdcc.Get(_fceb))
					if !_fdbf {
						continue
					}
					_bceg, _fdbf := _cgd.GetDict(_eacgd.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
					if !_fdbf {
						continue
					}
					_edab, _fdbf := _cgd.GetDict(_bceg.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
					if !_fdbf {
						continue
					}
					for _, _begf := range _edab.Keys() {
						_cdaf, _cdac := _cgd.GetDict(_edab.Get(_begf))
						if !_cdac {
							continue
						}
						_edbd := _dfcdf(_cdaf)
						if _edbd != nil {
							continue
						}
					}
				}
			}
		}
	}
	return nil
}
func _bbe(_ebeg *_gg.Document, _gcg func() _eg.Time) error {
	_acf, _acc := _f.NewPdfInfoFromObject(_ebeg.Info)
	if _acc != nil {
		return _acc
	}
	if _dfag := _adec(_acf, _gcg); _dfag != nil {
		return _dfag
	}
	_ebeg.Info = _acf.ToPdfObject()
	return nil
}
func _eecc(_bbae *_f.CompliancePdfReader) ViolatedRule {
	if _bbae.ParserMetadata().HasDataAfterEOF() {
		return _cc("\u0036.\u0031\u002e\u0033\u002d\u0033", "\u004e\u006f\u0020\u0064\u0061ta\u0020\u0073h\u0061\u006c\u006c\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0020\u0074\u0068\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0065\u006e\u0064\u002d\u006f\u0066\u002d\u0066\u0069l\u0065\u0020\u006da\u0072\u006b\u0065\u0072\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0061 \u0073\u0069\u006e\u0067\u006ce\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061\u006c \u0065\u006ed\u002do\u0066\u002d\u006c\u0069\u006e\u0065\u0020m\u0061\u0072\u006b\u0065\u0072\u002e")
	}
	return _cb
}
func _fede(_efa *_gg.Document, _cfe standardType, _gda XmpOptions) error {
	_dea, _adc := _efa.FindCatalog()
	if !_adc {
		return nil
	}
	var _fcf *_cg.Document
	_fcb, _adc := _dea.GetMetadata()
	if !_adc {
		_fcf = _cg.NewDocument()
	} else {
		var _dege error
		_fcf, _dege = _cg.LoadDocument(_fcb.Stream)
		if _dege != nil {
			return _dege
		}
	}
	_ecce := _cg.PdfInfoOptions{InfoDict: _efa.Info, PdfVersion: _a.Sprintf("\u0025\u0064\u002e%\u0064", _efa.Version.Major, _efa.Version.Minor), Copyright: _gda.Copyright, Overwrite: true}
	_bgdg, _adc := _dea.GetMarkInfo()
	if _adc {
		_bgdge, _deb := _cgd.GetBool(_bgdg.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
		if _deb && bool(*_bgdge) {
			_ecce.Marked = true
		}
	}
	if _afa := _fcf.SetPdfInfo(&_ecce); _afa != nil {
		return _afa
	}
	if _gfbd := _fcf.SetPdfAID(_cfe._ba, _cfe._fa); _gfbd != nil {
		return _gfbd
	}
	_ggd := _cg.MediaManagementOptions{OriginalDocumentID: _gda.OriginalDocumentID, DocumentID: _gda.DocumentID, InstanceID: _gda.InstanceID, NewDocumentID: !_gda.NewDocumentVersion, ModifyComment: "O\u0070\u0074\u0069\u006d\u0069\u007ae\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u006f\u0020\u0050D\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064"}
	_afag, _adc := _cgd.GetDict(_efa.Info)
	if _adc {
		if _eedc, _edd := _cgd.GetString(_afag.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _edd && _eedc.String() != "" {
			_fcab, _fbce := _ca.ParsePdfTime(_eedc.String())
			if _fbce != nil {
				return _a.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _fbce)
			}
			_ggd.ModifyDate = _fcab
		}
	}
	if _cfab := _fcf.SetMediaManagement(&_ggd); _cfab != nil {
		return _cfab
	}
	if _cbc := _fcf.SetPdfAExtension(); _cbc != nil {
		return _cbc
	}
	_bfa, _dgf := _fcf.MarshalIndent(_gda.MarshalPrefix, _gda.MarshalIndent)
	if _dgf != nil {
		return _dgf
	}
	if _bag := _dea.SetMetadata(_bfa); _bag != nil {
		return _bag
	}
	return nil
}
func _acbd(_ggda *_f.CompliancePdfReader) (_ddg []ViolatedRule) {
	for _, _affae := range _ggda.GetObjectNums() {
		_fegc, _bggc := _ggda.GetIndirectObjectByNumber(_affae)
		if _bggc != nil {
			continue
		}
		_aecd, _efegb := _cgd.GetDict(_fegc)
		if !_efegb {
			continue
		}
		_aeeb, _efegb := _cgd.GetName(_aecd.Get("\u0054\u0079\u0070\u0065"))
		if !_efegb {
			continue
		}
		if _aeeb.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_edcba, _efegb := _cgd.GetBool(_aecd.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _efegb && bool(*_edcba) {
			_ddg = append(_ddg, _cc("\u0036.\u0034\u002e\u0031\u002d\u0033", "\u0054\u0068\u0065\u0020\u004e\u0065e\u0064\u0041\u0070\u0070\u0065a\u0072\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006c\u0061\u0067\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0069\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076e\u0020\u0066\u006f\u0072\u006d \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u006e\u006f\u0074\u0020b\u0065\u0020\u0070\u0072\u0065se\u006e\u0074\u0020\u006f\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
		}
		if _aecd.Get("\u0058\u0046\u0041") != nil {
			_ddg = append(_ddg, _cc("\u0036.\u0034\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020i\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020t\u0068\u0061\u0074\u0020f\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065 \u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d \u006b\u0065\u0079\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006f\u0066 \u0061 \u0050\u0044F\u002fA\u002d\u0032\u0020\u0066ile\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065n\u0074\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0058\u0046\u0041\u0020\u006b\u0065y."))
		}
	}
	_gcff, _bfdaa := _ecgd(_ggda)
	if _bfdaa && _gcff.Get("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067") != nil {
		_ddg = append(_ddg, _cc("\u0036.\u0034\u002e\u0032\u002d\u0032", "\u0041\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020\u0043\u0061\u0074\u0061\u006cog\u0020s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u002e"))
	}
	return _ddg
}
func _eeged(_agfc *_cgd.PdfObjectDictionary, _dgfa map[*_cgd.PdfObjectStream][]byte, _ffad map[*_cgd.PdfObjectStream]*_bd.CMap) ViolatedRule {
	const (
		_bdca  = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0034"
		_ddfec = "\u0046\u006f\u0072\u0020\u0074\u0068\u006fs\u0065\u0020\u0043\u004d\u0061\u0070\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072e\u0020\u0065m\u0062\u0065\u0064de\u0064\u002c\u0020\u0074\u0068\u0065\u0020\u0069\u006et\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0057\u004d\u006f\u0064\u0065\u0020\u0065\u006e\u0074r\u0079\u0020i\u006e t\u0068\u0065\u0020CM\u0061\u0070\u0020\u0064\u0069\u0063\u0074\u0069o\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u0020\u0074\u006f \u0074h\u0065\u0020\u0057\u004d\u006f\u0064e\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064ed\u0020\u0043\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"
	)
	var _bdb string
	if _fgee, _fgdge := _cgd.GetName(_agfc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _fgdge {
		_bdb = _fgee.String()
	}
	if _bdb != "\u0054\u0079\u0070e\u0030" {
		return _cb
	}
	_gdfaa := _agfc.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _, _edeb := _cgd.GetName(_gdfaa); _edeb {
		return _cb
	}
	_cacad, _ccgd := _cgd.GetStream(_gdfaa)
	if !_ccgd {
		return _cc(_bdca, _ddfec)
	}
	_fdaef, _bfffb := _acbf(_cacad, _dgfa, _ffad)
	if _bfffb != nil {
		return _cc(_bdca, _ddfec)
	}
	_gbfa, _fgbf := _cgd.GetIntVal(_cacad.Get("\u0057\u004d\u006fd\u0065"))
	_geabg, _bebbac := _fdaef.WMode()
	if _fgbf && _bebbac {
		if _geabg != _gbfa {
			return _cc(_bdca, _ddfec)
		}
	}
	if (_fgbf && !_bebbac) || (!_fgbf && _bebbac) {
		return _cc(_bdca, _ddfec)
	}
	return _cb
}
func _bcbc(_geecc *_f.CompliancePdfReader) (_ecff ViolatedRule) {
	for _, _edac := range _geecc.GetObjectNums() {
		_ecee, _gecc := _geecc.GetIndirectObjectByNumber(_edac)
		if _gecc != nil {
			continue
		}
		_aefa, _adfg := _cgd.GetStream(_ecee)
		if !_adfg {
			continue
		}
		_caga, _adfg := _cgd.GetName(_aefa.Get("\u0054\u0079\u0070\u0065"))
		if !_adfg {
			continue
		}
		if *_caga != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_bbacg, _adfg := _cgd.GetName(_aefa.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_adfg {
			continue
		}
		if *_bbacg == "\u0050\u0053" {
			return _cc("\u0036.\u0032\u002e\u0037\u002d\u0031", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073.")
		}
	}
	return _ecff
}
func _affcf(_dedaa *_f.CompliancePdfReader) (_gcfb []ViolatedRule) {
	_gdbea, _eebbc := _ecgd(_dedaa)
	if !_eebbc {
		return _gcfb
	}
	_aeadd, _eebbc := _cgd.GetDict(_gdbea.Get("\u0050\u0065\u0072m\u0073"))
	if !_eebbc {
		return _gcfb
	}
	_bdfb := _aeadd.Keys()
	for _, _eefa := range _bdfb {
		if _eefa.String() != "\u0055\u0052\u0033" && _eefa.String() != "\u0044\u006f\u0063\u004d\u0044\u0050" {
			_gcfb = append(_gcfb, _cc("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0031", "\u004e\u006f\u0020\u006b\u0065\u0079\u0073 \u006f\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0055\u0052\u0033 \u0061n\u0064\u0020\u0044\u006f\u0063\u004dD\u0050\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0065\u0072\u006d\u0069\u0073\u0073i\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u002e"))
		}
	}
	return _gcfb
}

var _ Profile = (*Profile1B)(nil)

func _febg(_cfdd *_gg.Document) error {
	_cacc, _fecbe := _cfdd.GetPages()
	if !_fecbe {
		return nil
	}
	for _, _gdb := range _cacc {
		_bacd, _eecae := _cgd.GetArray(_gdb.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_eecae {
			continue
		}
		for _, _ffgb := range _bacd.Elements() {
			_ffgb = _cgd.ResolveReference(_ffgb)
			if _, _fefc := _ffgb.(*_cgd.PdfObjectNull); _fefc {
				continue
			}
			_febfe, _ceggf := _cgd.GetDict(_ffgb)
			if !_ceggf {
				continue
			}
			_gddfa, _ := _cgd.GetIntVal(_febfe.Get("\u0046"))
			_gddfa &= ^(1 << 0)
			_gddfa &= ^(1 << 1)
			_gddfa &= ^(1 << 5)
			_gddfa &= ^(1 << 8)
			_gddfa |= 1 << 2
			_febfe.Set("\u0046", _cgd.MakeInteger(int64(_gddfa)))
			_adfff := false
			if _gadf := _febfe.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _gadf != nil {
				_bbf, _dcfg := _cgd.GetName(_gadf)
				if _dcfg && _bbf.String() == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					_adfff = true
					if _febfe.Get("\u0041\u0041") != nil {
						_febfe.Remove("\u0041\u0041")
					}
					if _febfe.Get("\u0041") != nil {
						_febfe.Remove("\u0041")
					}
				}
				if _dcfg && _bbf.String() == "\u0054\u0065\u0078\u0074" {
					_daga, _ := _cgd.GetIntVal(_febfe.Get("\u0046"))
					_daga |= 1 << 3
					_daga |= 1 << 4
					_febfe.Set("\u0046", _cgd.MakeInteger(int64(_daga)))
				}
			}
			_cddg, _ceggf := _cgd.GetDict(_febfe.Get("\u0041\u0050"))
			if _ceggf {
				_gcca := _cddg.Get("\u004e")
				if _gcca == nil {
					continue
				}
				if len(_cddg.Keys()) > 1 {
					_cddg.Clear()
					_cddg.Set("\u004e", _gcca)
				}
				if _adfff {
					_daag, _aad := _cgd.GetName(_febfe.Get("\u0046\u0054"))
					if _aad && *_daag == "\u0042\u0074\u006e" {
						continue
					}
				}
			}
		}
	}
	return nil
}
func _gcdd(_aacg *_f.CompliancePdfReader) (_aagad []ViolatedRule) {
	_bfcbd, _eddga := _ecgd(_aacg)
	if !_eddga {
		return _aagad
	}
	_cgafa, _eddga := _cgd.GetDict(_bfcbd.Get("\u004e\u0061\u006de\u0073"))
	if !_eddga {
		return _aagad
	}
	if _cgafa.Get("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073") != nil {
		_aagad = append(_aagad, _cc("\u0036\u002e\u0031\u0030\u002d\u0031", "T\u0068\u0065\u0072e\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u006e\u006f\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0050\u0072\u0065s\u0065\u006e\u0074a\u0074\u0069\u006f\u006e\u0073\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u0027\u0073\u0020\u006e\u0061\u006d\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002e"))
	}
	return _aagad
}
func _cbaa(_facg *_f.CompliancePdfReader) ViolatedRule {
	for _, _cgge := range _facg.PageList {
		_dgba, _gacfa := _cgge.GetContentStreams()
		if _gacfa != nil {
			continue
		}
		for _, _bbbc := range _dgba {
			_cgea := _ff.NewContentStreamParser(_bbbc)
			_, _gacfa = _cgea.Parse()
			if _gacfa != nil {
				return _cc("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0031", "\u0041\u0020\u0063onten\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079 \u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u0020\u006e\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0065\u0076\u0065\u006e\u0020\u0069\u0066\u0020s\u0075\u0063\u0068\u0020\u006f\u0070\u0065r\u0061\u0074\u006f\u0072\u0073\u0020\u0061\u0072\u0065\u0020\u0062\u0072\u0061\u0063\u006b\u0065\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0074\u0068\u0065\u0020\u0042\u0058\u002f\u0045\u0058\u0020\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u002e")
			}
		}
	}
	return _cb
}
func _egeea(_bdea *_f.CompliancePdfReader) (_dgefd ViolatedRule) {
	for _, _efdc := range _bdea.GetObjectNums() {
		_bbbb, _fbfec := _bdea.GetIndirectObjectByNumber(_efdc)
		if _fbfec != nil {
			continue
		}
		_dcdb, _aece := _cgd.GetStream(_bbbb)
		if !_aece {
			continue
		}
		_bfda, _aece := _cgd.GetName(_dcdb.Get("\u0054\u0079\u0070\u0065"))
		if !_aece {
			continue
		}
		if *_bfda != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_, _aece = _cgd.GetName(_dcdb.Get("\u004f\u0050\u0049"))
		if _aece {
			return _cc("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		_bbcc, _aece := _cgd.GetName(_dcdb.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"))
		if !_aece {
			continue
		}
		if *_bbcc == "\u0050\u0053" {
			return _cc("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		if _dcdb.Get("\u0050\u0053") != nil {
			return _cc("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _dgefd
}

// VerificationError is the PDF/A verification error structure, that contains all violated rules.
type VerificationError struct {

	// ViolatedRules are the rules that were violated during error verification.
	ViolatedRules []ViolatedRule

	// ConformanceLevel defines the standard on verification failed.
	ConformanceLevel int

	// ConformanceVariant is the standard variant used on verification.
	ConformanceVariant string
}
