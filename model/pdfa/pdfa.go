package pdfa

import (
	_gg "errors"
	_d "fmt"
	_g "image/color"
	_e "math"
	_cg "sort"
	_eg "strings"
	_a "time"

	_ca "bitbucket.org/shenghui0779/gopdf/common"
	_cad "bitbucket.org/shenghui0779/gopdf/contentstream"
	_ad "bitbucket.org/shenghui0779/gopdf/core"
	_b "bitbucket.org/shenghui0779/gopdf/internal/cmap"
	_de "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_fe "bitbucket.org/shenghui0779/gopdf/internal/timeutils"
	_da "bitbucket.org/shenghui0779/gopdf/model"
	_ac "bitbucket.org/shenghui0779/gopdf/model/internal/colorprofile"
	_ea "bitbucket.org/shenghui0779/gopdf/model/internal/docutil"
	_cac "bitbucket.org/shenghui0779/gopdf/model/internal/fonts"
	_gf "bitbucket.org/shenghui0779/gopdf/model/xmputil"
	_be "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaextension"
	_ebe "bitbucket.org/shenghui0779/gopdf/model/xmputil/pdfaid"
	_dad "github.com/adrg/sysfont"
	_cb "github.com/trimmer-io/go-xmp/models/dc"
	_f "github.com/trimmer-io/go-xmp/models/pdf"
	_ef "github.com/trimmer-io/go-xmp/models/xmp_base"
	_cf "github.com/trimmer-io/go-xmp/models/xmp_mm"
	_eb "github.com/trimmer-io/go-xmp/models/xmp_rights"
	_ga "github.com/trimmer-io/go-xmp/xmp"
)

func _decc(_ecdfb *_da.CompliancePdfReader, _adgbc standardType) (_ebag []ViolatedRule) {
	var _ddfe, _dgbe, _ggea, _bgbe, _dfgd, _fege, _gedg, _fbgc, _egce, _bgdbg, _gaadf bool
	_aada := func() bool {
		return _ddfe && _dgbe && _ggea && _bgbe && _dfgd && _fege && _gedg && _fbgc && _egce && _bgdbg && _gaadf
	}
	_aaad := map[*_ad.PdfObjectStream]*_b.CMap{}
	_faccc := map[*_ad.PdfObjectStream][]byte{}
	_gbcc := map[_ad.PdfObject]*_da.PdfFont{}
	for _, _bdcd := range _ecdfb.GetObjectNums() {
		_dcdb, _ccagg := _ecdfb.GetIndirectObjectByNumber(_bdcd)
		if _ccagg != nil {
			continue
		}
		_afab, _gebc := _ad.GetDict(_dcdb)
		if !_gebc {
			continue
		}
		_fdab, _gebc := _ad.GetName(_afab.Get("\u0054\u0079\u0070\u0065"))
		if !_gebc {
			continue
		}
		if *_fdab != "\u0046\u006f\u006e\u0074" {
			continue
		}
		_ffda, _ccagg := _da.NewPdfFontFromPdfObject(_afab)
		if _ccagg != nil {
			_ca.Log.Debug("g\u0065\u0074\u0074\u0069\u006e\u0067 \u0066\u006f\u006e\u0074\u0020\u0066r\u006f\u006d\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0076", _ccagg)
			continue
		}
		_gbcc[_afab] = _ffda
	}
	for _, _bgfca := range _ecdfb.PageList {
		_bbbg, _gfbd := _bgfca.GetContentStreams()
		if _gfbd != nil {
			_ca.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0070\u0061\u0067\u0065\u0020\u0063o\u006e\u0074\u0065\u006e\u0074\u0020\u0073t\u0072\u0065\u0061\u006d\u0073\u0020\u0066\u0061\u0069\u006ce\u0064")
			continue
		}
		for _, _cbg := range _bbbg {
			_daca := _cad.NewContentStreamParser(_cbg)
			_dbece, _ecdfe := _daca.Parse()
			if _ecdfe != nil {
				_ca.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ecdfe)
				continue
			}
			var _acf bool
			for _, _fdgce := range *_dbece {
				if _fdgce.Operand != "\u0054\u0072" {
					continue
				}
				if len(_fdgce.Params) != 1 {
					_ca.Log.Debug("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054\u0072\u0027\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0065\u0064\u0020\u0027\u0031\u0027\u0020\u0062\u0075\u0074 \u0069\u0073\u003a\u0020\u0027\u0025d\u0027", len(_fdgce.Params))
					continue
				}
				_dabf, _cdfc := _ad.GetIntVal(_fdgce.Params[0])
				if !_cdfc {
					_ca.Log.Debug("\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020i\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
					continue
				}
				if _dabf == 3 {
					_acf = true
					break
				}
			}
			for _, _eebee := range *_dbece {
				if _eebee.Operand != "\u0054\u0066" {
					continue
				}
				if len(_eebee.Params) != 2 {
					_ca.Log.Debug("i\u006eva\u006ci\u0064 \u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066 \u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072s\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054f\u0027\u0020\u006fper\u0061\u006e\u0064\u002c\u0020\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0032\u0027\u0020\u0069s\u003a \u0027\u0025\u0064\u0027", len(_eebee.Params))
					continue
				}
				_beda, _dcbe := _ad.GetName(_eebee.Params[0])
				if !_dcbe {
					_ca.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _eebee)
					continue
				}
				_cefc, _efbe := _bgfca.Resources.GetFontByName(*_beda)
				if !_efbe {
					_ca.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
					continue
				}
				_bedfa, _dcbe := _ad.GetDict(_cefc)
				if !_dcbe {
					_ca.Log.Debug("\u0066\u006f\u006e\u0074 d\u0069\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
					continue
				}
				_ggde, _dcbe := _gbcc[_bedfa]
				if !_dcbe {
					var _agbf error
					_ggde, _agbf = _da.NewPdfFontFromPdfObject(_bedfa)
					if _agbf != nil {
						_ca.Log.Debug("\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0072o\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _agbf)
						continue
					}
					_gbcc[_bedfa] = _ggde
				}
				if !_ddfe {
					_edee := _fgafc(_bedfa, _faccc, _aaad)
					if _edee != _bg {
						_ebag = append(_ebag, _edee)
						_ddfe = true
						if _aada() {
							return _ebag
						}
					}
				}
				if !_dgbe {
					_fbaa := _dabb(_bedfa)
					if _fbaa != _bg {
						_ebag = append(_ebag, _fbaa)
						_dgbe = true
						if _aada() {
							return _ebag
						}
					}
				}
				if !_ggea {
					_fddb := _abgc(_bedfa, _faccc, _aaad)
					if _fddb != _bg {
						_ebag = append(_ebag, _fddb)
						_ggea = true
						if _aada() {
							return _ebag
						}
					}
				}
				if !_bgbe {
					_dgcf := _dfdffa(_bedfa, _faccc, _aaad)
					if _dgcf != _bg {
						_ebag = append(_ebag, _dgcf)
						_bgbe = true
						if _aada() {
							return _ebag
						}
					}
				}
				if !_dfgd {
					_bdcc := _dgbf(_ggde, _bedfa, _acf)
					if _bdcc != _bg {
						_dfgd = true
						_ebag = append(_ebag, _bdcc)
						if _aada() {
							return _ebag
						}
					}
				}
				if !_fege {
					_ddcc := _ddccc(_ggde, _bedfa)
					if _ddcc != _bg {
						_fege = true
						_ebag = append(_ebag, _ddcc)
						if _aada() {
							return _ebag
						}
					}
				}
				if !_gedg {
					_ceab := _faea(_ggde, _bedfa)
					if _ceab != _bg {
						_gedg = true
						_ebag = append(_ebag, _ceab)
						if _aada() {
							return _ebag
						}
					}
				}
				if !_fbgc {
					_ceaa := _fggce(_ggde, _bedfa)
					if _ceaa != _bg {
						_fbgc = true
						_ebag = append(_ebag, _ceaa)
						if _aada() {
							return _ebag
						}
					}
				}
				if !_egce {
					_fcab := _afda(_ggde, _bedfa)
					if _fcab != _bg {
						_egce = true
						_ebag = append(_ebag, _fcab)
						if _aada() {
							return _ebag
						}
					}
				}
				if !_bgdbg {
					_fgaf := _faga(_ggde, _bedfa)
					if _fgaf != _bg {
						_bgdbg = true
						_ebag = append(_ebag, _fgaf)
						if _aada() {
							return _ebag
						}
					}
				}
				if !_gaadf && _adgbc._ee == "\u0041" {
					_fecc := _fefe(_bedfa, _faccc, _aaad)
					if _fecc != _bg {
						_gaadf = true
						_ebag = append(_ebag, _fecc)
						if _aada() {
							return _ebag
						}
					}
				}
			}
		}
	}
	return _ebag
}
func _aca(_gccb *_ea.Document, _bfb standardType, _eac XmpOptions) error {
	_eda, _cfe := _gccb.FindCatalog()
	if !_cfe {
		return nil
	}
	var _dbec *_gf.Document
	_ecfb, _cfe := _eda.GetMetadata()
	if !_cfe {
		_dbec = _gf.NewDocument()
	} else {
		var _eebb error
		_dbec, _eebb = _gf.LoadDocument(_ecfb.Stream)
		if _eebb != nil {
			return _eebb
		}
	}
	_bdg := _gf.PdfInfoOptions{InfoDict: _gccb.Info, PdfVersion: _d.Sprintf("\u0025\u0064\u002e%\u0064", _gccb.Version.Major, _gccb.Version.Minor), Copyright: _eac.Copyright, Overwrite: true}
	_acbc, _cfe := _eda.GetMarkInfo()
	if _cfe {
		_dcb, _ggff := _ad.GetBool(_acbc.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
		if _ggff && bool(*_dcb) {
			_bdg.Marked = true
		}
	}
	if _adgg := _dbec.SetPdfInfo(&_bdg); _adgg != nil {
		return _adgg
	}
	if _gdde := _dbec.SetPdfAID(_bfb._efe, _bfb._ee); _gdde != nil {
		return _gdde
	}
	_ffb := _gf.MediaManagementOptions{OriginalDocumentID: _eac.OriginalDocumentID, DocumentID: _eac.DocumentID, InstanceID: _eac.InstanceID, NewDocumentID: !_eac.NewDocumentVersion, ModifyComment: "O\u0070\u0074\u0069\u006d\u0069\u007ae\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u006f\u0020\u0050D\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064"}
	_bdb, _cfe := _ad.GetDict(_gccb.Info)
	if _cfe {
		if _eada, _efc := _ad.GetString(_bdb.Get("\u004do\u0064\u0044\u0061\u0074\u0065")); _efc && _eada.String() != "" {
			_efgf, _aff := _fe.ParsePdfTime(_eada.String())
			if _aff != nil {
				return _d.Errorf("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w", _aff)
			}
			_ffb.ModifyDate = _efgf
		}
	}
	if _dafd := _dbec.SetMediaManagement(&_ffb); _dafd != nil {
		return _dafd
	}
	if _aafb := _dbec.SetPdfAExtension(); _aafb != nil {
		return _aafb
	}
	_fbg, _dec := _dbec.MarshalIndent(_eac.MarshalPrefix, _eac.MarshalIndent)
	if _dec != nil {
		return _dec
	}
	if _fge := _eda.SetMetadata(_fbg); _fge != nil {
		return _fge
	}
	return nil
}
func _cecec(_dcedc *_da.CompliancePdfReader) []ViolatedRule { return nil }
func (_dac *documentImages) hasUncalibratedImages() bool    { return _dac._eea || _dac._fc || _dac._eeb }
func _gga(_fcef string, _ebfg string, _dfac string) (string, bool) {
	_cgad := _eg.Index(_fcef, _ebfg)
	if _cgad == -1 {
		return "", false
	}
	_cgad += len(_ebfg)
	_gbae := _eg.Index(_fcef[_cgad:], _dfac)
	if _gbae == -1 {
		return "", false
	}
	_gbae = _cgad + _gbae
	return _fcef[_cgad:_gbae], true
}
func _gde(_eead *_da.XObjectImage, _dg imageModifications) error {
	_fba, _gbfa := _eead.ToImage()
	if _gbfa != nil {
		return _gbfa
	}
	if _dg._fgg != nil {
		_eead.Filter = _dg._fgg
	}
	_fcg := _ad.MakeDict()
	_fcg.Set("\u0051u\u0061\u006c\u0069\u0074\u0079", _ad.MakeInteger(100))
	_fcg.Set("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr", _ad.MakeInteger(1))
	_eead.Decode = nil
	if _gbfa = _eead.SetImage(_fba, nil); _gbfa != nil {
		return _gbfa
	}
	_eead.ToPdfObject()
	return nil
}
func _bedf(_geaa *_da.CompliancePdfReader) (_gbcf []ViolatedRule) {
	var (
		_aecf, _gcf, _bgdd, _aaac, _dcad, _ecebc, _dceag bool
		_fedd                                            func(_ad.PdfObject)
	)
	_fedd = func(_bfbf _ad.PdfObject) {
		switch _dde := _bfbf.(type) {
		case *_ad.PdfObjectInteger:
			if !_aecf && (int64(*_dde) > _e.MaxInt32 || int64(*_dde) < -_e.MaxInt32) {
				_gbcf = append(_gbcf, _ba("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0031", "L\u0061\u0072\u0067e\u0073\u0074\u0020\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0032\u002c\u0031\u0034\u0037,\u0034\u0038\u0033,\u0036\u0034\u0037\u002e\u0020\u0053\u006d\u0061\u006c\u006c\u0065\u0073\u0074 \u0069\u006e\u0074\u0065g\u0065\u0072\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u002d\u0032\u002c\u0031\u0034\u0037\u002c\u0034\u0038\u0033,\u0036\u0034\u0038\u002e"))
				_aecf = true
			}
		case *_ad.PdfObjectFloat:
			if !_gcf && (_e.Abs(float64(*_dde)) > 32767.0) {
				_gbcf = append(_gbcf, _ba("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0032", "\u0041\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0020\u0072\u0065\u0061\u006c\u0020\u0076\u0061\u006c\u0075\u0065\u0020m\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u006c\u0065s\u0073\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075a\u006c\u0020\u0074\u006f\u0020\u00332\u0037\u0036\u0037.\u0030\u002e"))
			}
		case *_ad.PdfObjectString:
			if !_bgdd && len([]byte(_dde.Str())) > 65535 {
				_gbcf = append(_gbcf, _ba("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0033", "M\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006c\u0065n\u0067\u0074\u0068\u0020\u006f\u0066\u0020a \u0073\u0074\u0072\u0069n\u0067\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074es\u0029\u0020i\u0073\u0020\u0036\u0035\u0035\u0033\u0035\u002e"))
				_bgdd = true
			}
		case *_ad.PdfObjectName:
			if !_aaac && len([]byte(*_dde)) > 127 {
				_gbcf = append(_gbcf, _ba("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0034", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006c\u0065\u006eg\u0074\u0068\u0020\u006ff\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0029\u0020\u0069\u0073\u0020\u0031\u0032\u0037\u002e"))
				_aaac = true
			}
		case *_ad.PdfObjectArray:
			if !_dcad && _dde.Len() > 8191 {
				_gbcf = append(_gbcf, _ba("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0035", "\u004d\u0061\u0078\u0069\u006d\u0075m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006f\u0066\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079\u0020(\u0069\u006e\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0029\u0020\u0069s\u00208\u0031\u0039\u0031\u002e"))
				_dcad = true
			}
			for _, _cfca := range _dde.Elements() {
				_fedd(_cfca)
			}
			if !_dceag && (_dde.Len() == 4 || _dde.Len() == 5) {
				_dfcca, _dcf := _ad.GetName(_dde.Get(0))
				if !_dcf {
					return
				}
				if *_dfcca != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
					return
				}
				_cfdd := _dde.Get(1)
				_cfdd = _ad.TraceToDirectObject(_cfdd)
				_bcbaa, _dcf := _ad.GetArray(_cfdd)
				if !_dcf {
					return
				}
				if _bcbaa.Len() > 8 {
					_gbcf = append(_gbcf, _ba("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0039", "\u004d\u0061\u0078i\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0069\u0073\u0020\u0038\u002e"))
					_dceag = true
				}
			}
		case *_ad.PdfObjectDictionary:
			_afdf := _dde.Keys()
			if !_ecebc && len(_afdf) > 4095 {
				_gbcf = append(_gbcf, _ba("\u0036.\u0031\u002e\u0031\u0032\u002d\u00311", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u0063\u0061\u0070\u0061\u0063\u0069\u0074y\u0020\u006f\u0066\u0020\u0061\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0028\u0069\u006e\u0020\u0065\u006e\u0074\u0072\u0069es\u0029\u0020\u0069\u0073\u0020\u0034\u0030\u0039\u0035\u002e"))
				_ecebc = true
			}
			for _gfdb, _fead := range _afdf {
				_fedd(&_afdf[_gfdb])
				_fedd(_dde.Get(_fead))
			}
		case *_ad.PdfObjectStream:
			_fedd(_dde.PdfObjectDictionary)
		case *_ad.PdfObjectStreams:
			for _, _eccff := range _dde.Elements() {
				_fedd(_eccff)
			}
		case *_ad.PdfObjectReference:
			_fedd(_dde.Resolve())
		}
	}
	_ebdd := _geaa.GetObjectNums()
	if len(_ebdd) > 8388607 {
		_gbcf = append(_gbcf, _ba("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0037", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020in\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073 \u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0069\u0073\u00208\u002c\u0033\u0038\u0038\u002c\u0036\u0030\u0037\u002e"))
	}
	for _, _bcgd := range _ebdd {
		_cded, _bgfea := _geaa.GetIndirectObjectByNumber(_bcgd)
		if _bgfea != nil {
			continue
		}
		_fadga := _ad.TraceToDirectObject(_cded)
		_fedd(_fadga)
	}
	return _gbcf
}
func _dcc(_dfa *_ea.Document) error {
	for _, _ddg := range _dfa.Objects {
		_gegb, _ggbc := _ad.GetDict(_ddg)
		if !_ggbc {
			continue
		}
		_ddf := _gegb.Get("\u0054\u0079\u0070\u0065")
		if _ddf == nil {
			continue
		}
		if _gfb, _deff := _ad.GetName(_ddf); _deff && _gfb.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_ab, _eaag := _ad.GetBool(_gegb.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _eaag {
			if bool(*_ab) {
				_gegb.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _ad.MakeBool(false))
			}
		}
		_gge := _gegb.Get("\u0041")
		if _gge != nil {
			_gegb.Remove("\u0041")
		}
		_cce, _eaag := _ad.GetArray(_gegb.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
		if _eaag {
			for _eaad := 0; _eaad < _cce.Len(); _eaad++ {
				_fda, _bgfe := _ad.GetDict(_cce.Get(_eaad))
				if !_bgfe {
					continue
				}
				if _fda.Get("\u0041\u0041") != nil {
					_fda.Remove("\u0041\u0041")
				}
			}
		}
	}
	return nil
}
func _dge(_dba *_da.PdfPageResources, _cgc *_cad.ContentStreamOperations, _bfae bool) ([]byte, error) {
	var _agd bool
	for _, _efcg := range *_cgc {
	_fgfe:
		switch _efcg.Operand {
		case "\u0042\u0049":
			_aegc, _egfg := _efcg.Params[0].(*_cad.ContentStreamInlineImage)
			if !_egfg {
				break
			}
			_fdaec, _deec := _aegc.GetColorSpace(_dba)
			if _deec != nil {
				return nil, _deec
			}
			switch _fdaec.(type) {
			case *_da.PdfColorspaceDeviceCMYK:
				if _bfae {
					break _fgfe
				}
			case *_da.PdfColorspaceDeviceGray:
			case *_da.PdfColorspaceDeviceRGB:
				if !_bfae {
					break _fgfe
				}
			default:
				break _fgfe
			}
			_agd = true
			_bef, _deec := _aegc.ToImage(_dba)
			if _deec != nil {
				return nil, _deec
			}
			_gebe, _deec := _bef.ToGoImage()
			if _deec != nil {
				return nil, _deec
			}
			if _bfae {
				_gebe, _deec = _de.CMYKConverter.Convert(_gebe)
			} else {
				_gebe, _deec = _de.NRGBAConverter.Convert(_gebe)
			}
			if _deec != nil {
				return nil, _deec
			}
			_cgf, _egfg := _gebe.(_de.Image)
			if !_egfg {
				return nil, _gg.New("\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074 \u0069\u006d\u0070\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002eI\u006d\u0061\u0067\u0065")
			}
			_cabd := _cgf.Base()
			_fdc := _da.Image{Width: int64(_cabd.Width), Height: int64(_cabd.Height), BitsPerComponent: int64(_cabd.BitsPerComponent), ColorComponents: _cabd.ColorComponents, Data: _cabd.Data}
			_fdc.SetDecode(_cabd.Decode)
			_fdc.SetAlpha(_cabd.Alpha)
			_bfgg, _deec := _aegc.GetEncoder()
			if _deec != nil {
				_bfgg = _ad.NewFlateEncoder()
			}
			_gdcc, _deec := _cad.NewInlineImageFromImage(_fdc, _bfgg)
			if _deec != nil {
				return nil, _deec
			}
			_efcg.Params[0] = _gdcc
		case "\u0047", "\u0067":
			if len(_efcg.Params) != 1 {
				break
			}
			_acdd, _acda := _ad.GetNumberAsFloat(_efcg.Params[0])
			if _acda != nil {
				break
			}
			if _bfae {
				_efcg.Params = []_ad.PdfObject{_ad.MakeFloat(0), _ad.MakeFloat(0), _ad.MakeFloat(0), _ad.MakeFloat(1 - _acdd)}
				_fbfc := "\u004b"
				if _efcg.Operand == "\u0067" {
					_fbfc = "\u006b"
				}
				_efcg.Operand = _fbfc
			} else {
				_efcg.Params = []_ad.PdfObject{_ad.MakeFloat(_acdd), _ad.MakeFloat(_acdd), _ad.MakeFloat(_acdd)}
				_fcbb := "\u0052\u0047"
				if _efcg.Operand == "\u0067" {
					_fcbb = "\u0072\u0067"
				}
				_efcg.Operand = _fcbb
			}
			_agd = true
		case "\u0052\u0047", "\u0072\u0067":
			if !_bfae {
				break
			}
			if len(_efcg.Params) != 3 {
				break
			}
			_ceacg, _cada := _ad.GetNumbersAsFloat(_efcg.Params)
			if _cada != nil {
				break
			}
			_agd = true
			_cbbf, _cdgd, _aeae := _ceacg[0], _ceacg[1], _ceacg[2]
			_gacd, _fcde, _bec, _edea := _g.RGBToCMYK(uint8(_cbbf*255), uint8(_cdgd*255), uint8(255*_aeae))
			_efcg.Params = []_ad.PdfObject{_ad.MakeFloat(float64(_gacd) / 255), _ad.MakeFloat(float64(_fcde) / 255), _ad.MakeFloat(float64(_bec) / 255), _ad.MakeFloat(float64(_edea) / 255)}
			_efac := "\u004b"
			if _efcg.Operand == "\u0072\u0067" {
				_efac = "\u006b"
			}
			_efcg.Operand = _efac
		case "\u004b", "\u006b":
			if _bfae {
				break
			}
			if len(_efcg.Params) != 4 {
				break
			}
			_ccbg, _decg := _ad.GetNumbersAsFloat(_efcg.Params)
			if _decg != nil {
				break
			}
			_bfag, _fddg, _bcfb, _fgef := _ccbg[0], _ccbg[1], _ccbg[2], _ccbg[3]
			_dfe, _eddd, _fbbf := _g.CMYKToRGB(uint8(255*_bfag), uint8(255*_fddg), uint8(255*_bcfb), uint8(255*_fgef))
			_efcg.Params = []_ad.PdfObject{_ad.MakeFloat(float64(_dfe) / 255), _ad.MakeFloat(float64(_eddd) / 255), _ad.MakeFloat(float64(_fbbf) / 255)}
			_dcbd := "\u0052\u0047"
			if _efcg.Operand == "\u006b" {
				_dcbd = "\u0072\u0067"
			}
			_efcg.Operand = _dcbd
			_agd = true
		}
	}
	if !_agd {
		return nil, nil
	}
	_gagf := _cad.NewContentCreator()
	for _, _aadb := range *_cgc {
		_gagf.AddOperand(*_aadb)
	}
	_age := _gagf.Bytes()
	return _age, nil
}
func _agge(_efbb *_ea.Document) error {
	_gfbe, _gfda := _efbb.FindCatalog()
	if !_gfda {
		return _gg.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_bgab, _gfda := _ad.GetDict(_gfbe.Object.Get("\u0050\u0065\u0072m\u0073"))
	if _gfda {
		_fdgf := _ad.MakeDict()
		_fedg := _bgab.Keys()
		for _, _egac := range _fedg {
			if _egac.String() == "\u0055\u0052\u0033" || _egac.String() == "\u0044\u006f\u0063\u004d\u0044\u0050" {
				_fdgf.Set(_egac, _bgab.Get(_egac))
			}
		}
		_gfbe.Object.Set("\u0050\u0065\u0072m\u0073", _fdgf)
	}
	return nil
}
func _faaa(_dcdd *_ea.Document) error {
	for _, _cfafb := range _dcdd.Objects {
		_ddab, _eddg := _ad.GetDict(_cfafb)
		if !_eddg {
			continue
		}
		_cddc := _ddab.Get("\u0054\u0079\u0070\u0065")
		if _cddc == nil {
			continue
		}
		if _aeac, _gdge := _ad.GetName(_cddc); _gdge && _aeac.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_gfgg, _gdgef := _ad.GetBool(_ddab.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _gdgef && bool(*_gfgg) {
			_ddab.Set("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073", _ad.MakeBool(false))
		}
		if _ddab.Get("\u0058\u0046\u0041") != nil {
			_ddab.Remove("\u0058\u0046\u0041")
		}
	}
	_bdaa, _ffega := _dcdd.FindCatalog()
	if !_ffega {
		return _gg.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _bdaa.Object.Get("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067") != nil {
		_bdaa.Object.Remove("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067")
	}
	return nil
}
func _acdf(_edcb *_ea.Document, _eca standardType, _fagb *_ea.OutputIntents) error {
	var (
		_dfgc *_da.PdfOutputIntent
		_dcde error
	)
	if _edcb.Version.Minor <= 7 {
		_dfgc, _dcde = _ac.NewSRGBv2OutputIntent(_eca.outputIntentSubtype())
	} else {
		_dfgc, _dcde = _ac.NewSRGBv4OutputIntent(_eca.outputIntentSubtype())
	}
	if _dcde != nil {
		return _dcde
	}
	if _dcde = _fagb.Add(_dfgc.ToPdfObject()); _dcde != nil {
		return _dcde
	}
	return nil
}
func _aggbf(_daedc *_da.CompliancePdfReader) (_fdabd []ViolatedRule) {
	_deefcg, _dgda := _acgab(_daedc)
	if !_dgda {
		return _fdabd
	}
	_bdca, _dgda := _ad.GetDict(_deefcg.Get("\u0050\u0065\u0072m\u0073"))
	if !_dgda {
		return _fdabd
	}
	_cfcefd := _bdca.Keys()
	for _, _accf := range _cfcefd {
		if _accf.String() != "\u0055\u0052\u0033" && _accf.String() != "\u0044\u006f\u0063\u004d\u0044\u0050" {
			_fdabd = append(_fdabd, _ba("\u0036\u002e\u0031\u002e\u0031\u0032\u002d\u0031", "\u004e\u006f\u0020\u006b\u0065\u0079\u0073 \u006f\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0055\u0052\u0033 \u0061n\u0064\u0020\u0044\u006f\u0063\u004dD\u0050\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0065\u0072\u006d\u0069\u0073\u0073i\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u002e"))
		}
	}
	return _fdabd
}
func _fdae(_bgbd *_ea.Document, _edbb int) {
	if _bgbd.Version.Major == 0 {
		_bgbd.Version.Major = 1
	}
	if _bgbd.Version.Minor < _edbb {
		_bgbd.Version.Minor = _edbb
	}
}
func _egdd(_ccdd *_da.CompliancePdfReader) ViolatedRule {
	_effe, _dbea := _ccdd.GetTrailer()
	if _dbea != nil {
		_ca.Log.Debug("\u0043\u0061\u006en\u006f\u0074\u0020\u0067e\u0074\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0076", _dbea)
		return _bg
	}
	_bacg, _bgeea := _effe.Get("\u0052\u006f\u006f\u0074").(*_ad.PdfObjectReference)
	if !_bgeea {
		_ca.Log.Debug("\u0043a\u006e\u006e\u006f\u0074 \u0066\u0069\u006e\u0064\u0020d\u006fc\u0075m\u0065\u006e\u0074\u0020\u0072\u006f\u006ft")
		return _bg
	}
	_aafe, _bgeea := _ad.GetDict(_ad.ResolveReference(_bacg))
	if !_bgeea {
		_ca.Log.Debug("\u0063\u0061\u006e\u006e\u006f\u0074 \u0072\u0065\u0073\u006f\u006c\u0076\u0065\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		return _bg
	}
	if _aafe.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073") != nil {
		return _ba("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063u\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u004f\u0043\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073")
	}
	return _bg
}

// ValidateStandard checks if provided input CompliancePdfReader matches rules that conforms PDF/A-1 standard.
func (_dafe *profile1) ValidateStandard(r *_da.CompliancePdfReader) error {
	_afdbb := VerificationError{ConformanceLevel: _dafe._efde._efe, ConformanceVariant: _dafe._efde._ee}
	if _afc := _beb(r); _afc != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _afc)
	}
	if _befb := _efgg(r); _befb != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _befb)
	}
	if _gdad := _gbdf(r); _gdad != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _gdad)
	}
	if _bfdf := _ceaf(r); _bfdf != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _bfdf)
	}
	if _bgee := _fdef(r); _bgee != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _bgee)
	}
	if _debb := _ceg(r); len(_debb) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _debb...)
	}
	if _dcea := _abdb(r); _dcea != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _dcea)
	}
	if _effd := _aedfg(r); len(_effd) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _effd...)
	}
	if _gcca := _bebg(r); len(_gcca) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _gcca...)
	}
	if _edag := _cecec(r); len(_edag) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _edag...)
	}
	if _gbff := _gbfd(r); _gbff != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _gbff)
	}
	if _bgdf := _eege(r); len(_bgdf) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _bgdf...)
	}
	if _bfbb := _bedf(r); len(_bfbb) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _bfbb...)
	}
	if _gage := _egdd(r); _gage != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _gage)
	}
	if _dcdg := _faabd(r, false); len(_dcdg) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _dcdg...)
	}
	if _bcbf := _gbcfg(r); len(_bcbf) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _bcbf...)
	}
	if _cafd := _dadf(r); _cafd != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _cafd)
	}
	if _eeag := _edfb(r); _eeag != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _eeag)
	}
	if _cca := _adde(r); _cca != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _cca)
	}
	if _ffcd := _bcdf(r); _ffcd != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _ffcd)
	}
	if _gcda := _feda(r); _gcda != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _gcda)
	}
	if _gacg := _bbac(r); len(_gacg) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _gacg...)
	}
	if _edca := _decc(r, _dafe._efde); len(_edca) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _edca...)
	}
	if _fdde := _fgeb(r); len(_fdde) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _fdde...)
	}
	if _fcbf := _gaae(r); _fcbf != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _fcbf)
	}
	if _baec := _bceg(r); _baec != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _baec)
	}
	if _eaf := _afcg(r); len(_eaf) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _eaf...)
	}
	if _bde := _ebge(r); len(_bde) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _bde...)
	}
	if _gdgeb := _dgdb(r); _gdgeb != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _gdgeb)
	}
	if _cdde := _ecff(r); _cdde != _bg {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _cdde)
	}
	if _gggc := _fgab(r, _dafe._efde, false); len(_gggc) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _gggc...)
	}
	if _dafe._efde == _df() {
		if _dgfa := _afdbea(r); len(_dgfa) != 0 {
			_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _dgfa...)
		}
	}
	if _cade := _bdbeb(r); len(_cade) != 0 {
		_afdbb.ViolatedRules = append(_afdbb.ViolatedRules, _cade...)
	}
	if len(_afdbb.ViolatedRules) > 0 {
		_cg.Slice(_afdbb.ViolatedRules, func(_dacd, _bcgc int) bool {
			return _afdbb.ViolatedRules[_dacd].RuleNo < _afdbb.ViolatedRules[_bcgc].RuleNo
		})
		return _afdbb
	}
	return nil
}
func _dbfde(_afcgb *_da.CompliancePdfReader, _aedb standardType) (_cbag []ViolatedRule) {
	var _ebde, _bfedd, _gfgb, _accgf, _dbfa, _aabe, _abde bool
	_dfgcc := func() bool { return _ebde && _bfedd && _gfgb && _accgf && _dbfa && _aabe && _abde }
	_gbadg := map[*_ad.PdfObjectStream]*_b.CMap{}
	_faadb := map[*_ad.PdfObjectStream][]byte{}
	_aebc := map[_ad.PdfObject]*_da.PdfFont{}
	for _, _fabeed := range _afcgb.GetObjectNums() {
		_bfddb, _ceebd := _afcgb.GetIndirectObjectByNumber(_fabeed)
		if _ceebd != nil {
			continue
		}
		_gagc, _beaea := _ad.GetDict(_bfddb)
		if !_beaea {
			continue
		}
		_ebbeb, _beaea := _ad.GetName(_gagc.Get("\u0054\u0079\u0070\u0065"))
		if !_beaea {
			continue
		}
		if *_ebbeb != "\u0046\u006f\u006e\u0074" {
			continue
		}
		_ffed, _ceebd := _da.NewPdfFontFromPdfObject(_gagc)
		if _ceebd != nil {
			_ca.Log.Debug("g\u0065\u0074\u0074\u0069\u006e\u0067 \u0066\u006f\u006e\u0074\u0020\u0066r\u006f\u006d\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0076", _ceebd)
			continue
		}
		_aebc[_gagc] = _ffed
	}
	for _, _bcbge := range _afcgb.PageList {
		_bbaf, _ccecc := _bcbge.GetContentStreams()
		if _ccecc != nil {
			_ca.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0070\u0061\u0067\u0065\u0020\u0063o\u006e\u0074\u0065\u006e\u0074\u0020\u0073t\u0072\u0065\u0061\u006d\u0073\u0020\u0066\u0061\u0069\u006ce\u0064")
			continue
		}
		for _, _edbc := range _bbaf {
			_cgff := _cad.NewContentStreamParser(_edbc)
			_egag, _fedfb := _cgff.Parse()
			if _fedfb != nil {
				_ca.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074s\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _fedfb)
				continue
			}
			var _cdgce bool
			for _, _efedg := range *_egag {
				if _efedg.Operand != "\u0054\u0072" {
					continue
				}
				if len(_efedg.Params) != 1 {
					_ca.Log.Debug("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054\u0072\u0027\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u002c\u0020\u0065\u0078\u0070e\u0063\u0074\u0065\u0064\u0020\u0027\u0031\u0027\u0020\u0062\u0075\u0074 \u0069\u0073\u003a\u0020\u0027\u0025d\u0027", len(_efedg.Params))
					continue
				}
				_ccdb, _cgba := _ad.GetIntVal(_efedg.Params[0])
				if !_cgba {
					_ca.Log.Debug("\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020i\u0073 \u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072")
					continue
				}
				if _ccdb == 3 {
					_cdgce = true
					break
				}
			}
			for _, _bdaabc := range *_egag {
				if _bdaabc.Operand != "\u0054\u0066" {
					continue
				}
				if len(_bdaabc.Params) != 2 {
					_ca.Log.Debug("i\u006eva\u006ci\u0064 \u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066 \u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072s\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0027\u0054f\u0027\u0020\u006fper\u0061\u006e\u0064\u002c\u0020\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0032\u0027\u0020\u0069s\u003a \u0027\u0025\u0064\u0027", len(_bdaabc.Params))
					continue
				}
				_ceceb, _efbecb := _ad.GetName(_bdaabc.Params[0])
				if !_efbecb {
					_ca.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a \u0054\u0066\u0020\u006f\u0070\u003d\u0025\u0073\u0020\u0047\u0065\u0074\u004ea\u006d\u0065\u0056\u0061\u006c\u0020\u0066a\u0069\u006c\u0065\u0064", _bdaabc)
					continue
				}
				_eeffe, _ggfca := _bcbge.Resources.GetFontByName(*_ceceb)
				if !_ggfca {
					_ca.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
					continue
				}
				_cgfe, _efbecb := _ad.GetDict(_eeffe)
				if !_efbecb {
					_ca.Log.Debug("\u0066\u006f\u006e\u0074 d\u0069\u0063\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
					continue
				}
				_bgca, _efbecb := _aebc[_cgfe]
				if !_efbecb {
					var _gcdg error
					_bgca, _gcdg = _da.NewPdfFontFromPdfObject(_cgfe)
					if _gcdg != nil {
						_ca.Log.Debug("\u0067\u0065\u0074\u0074i\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0072o\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u003a \u0025\u0076", _gcdg)
						continue
					}
					_aebc[_cgfe] = _bgca
				}
				if !_ebde {
					_bfde := _dgce(_cgfe, _faadb, _gbadg)
					if _bfde != _bg {
						_cbag = append(_cbag, _bfde)
						_ebde = true
						if _dfgcc() {
							return _cbag
						}
					}
				}
				if !_bfedd {
					_efef := _dgag(_cgfe)
					if _efef != _bg {
						_cbag = append(_cbag, _efef)
						_bfedd = true
						if _dfgcc() {
							return _cbag
						}
					}
				}
				if !_gfgb {
					_ebda := _aaabb(_cgfe, _faadb, _gbadg)
					if _ebda != _bg {
						_cbag = append(_cbag, _ebda)
						_gfgb = true
						if _dfgcc() {
							return _cbag
						}
					}
				}
				if !_accgf {
					_edfe := _dcdec(_cgfe, _faadb, _gbadg)
					if _edfe != _bg {
						_cbag = append(_cbag, _edfe)
						_accgf = true
						if _dfgcc() {
							return _cbag
						}
					}
				}
				if !_dbfa {
					_edcbe := _gdaca(_bgca, _cgfe, _cdgce)
					if _edcbe != _bg {
						_dbfa = true
						_cbag = append(_cbag, _edcbe)
						if _dfgcc() {
							return _cbag
						}
					}
				}
				if !_aabe {
					_cdafe := _dgcea(_bgca, _cgfe)
					if _cdafe != _bg {
						_aabe = true
						_cbag = append(_cbag, _cdafe)
						if _dfgcc() {
							return _cbag
						}
					}
				}
				if !_abde && (_aedb._ee == "\u0041" || _aedb._ee == "\u0055") {
					_ggfcc := _eecad(_cgfe, _faadb, _gbadg)
					if _ggfcc != _bg {
						_abde = true
						_cbag = append(_cbag, _ggfcc)
						if _dfgcc() {
							return _cbag
						}
					}
				}
			}
		}
	}
	return _cbag
}
func _dabb(_cfdae *_ad.PdfObjectDictionary) ViolatedRule {
	const (
		_bgbg = "\u0036.\u0033\u002e\u0033\u002d\u0032"
		_dcaa = "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0054y\u0070\u0065\u0020\u0032\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072\u0065\u0020\u0075\u0073\u0065\u0064\u0020f\u006f\u0072 \u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006e\u0067,\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0046\u006fn\u0074\u0020\u0064\u0069c\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c \u0063\u006f\u006e\u0074\u0061i\u006e\u0020\u0061\u0020\u0043\u0049\u0044\u0054\u006f\u0047\u0049D\u004d\u0061\u0070\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020a\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006d\u0061\u0070\u0070\u0069\u006e\u0067\u0020\u0066\u0072\u006f\u006d\u0020\u0043\u0049\u0044\u0073\u0020\u0074\u006f\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u0069\u006e\u0064\u0069c\u0065\u0073\u0020\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u0049d\u0065\u006e\u0074\u0069\u0074\u0079\u002c\u0020\u0061s d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020P\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0054a\u0062\u006c\u0065\u0020\u0035\u002e\u00313"
	)
	var _abfe string
	if _fdggg, _cfbe := _ad.GetName(_cfdae.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cfbe {
		_abfe = _fdggg.String()
	}
	if _abfe != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		return _bg
	}
	if _cfdae.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070") == nil {
		return _ba(_bgbg, _dcaa)
	}
	return _bg
}
func _edfb(_dfag *_da.CompliancePdfReader) (_dbac ViolatedRule) {
	for _, _ccbd := range _dfag.GetObjectNums() {
		_cgcgb, _dbbf := _dfag.GetIndirectObjectByNumber(_ccbd)
		if _dbbf != nil {
			continue
		}
		_dfcf, _fagf := _ad.GetStream(_cgcgb)
		if !_fagf {
			continue
		}
		_aege, _fagf := _ad.GetName(_dfcf.Get("\u0054\u0079\u0070\u0065"))
		if !_fagf {
			continue
		}
		if *_aege != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		if _dfcf.Get("\u0052\u0065\u0066") != nil {
			return _ba("\u0036.\u0032\u002e\u0036\u002d\u0031", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e")
		}
	}
	return _dbac
}
func _fefc(_aaef *_da.CompliancePdfReader) (_daeb ViolatedRule) {
	for _, _fefea := range _aaef.GetObjectNums() {
		_cgfc, _dbdaf := _aaef.GetIndirectObjectByNumber(_fefea)
		if _dbdaf != nil {
			continue
		}
		_dfef, _bfffc := _ad.GetStream(_cgfc)
		if !_bfffc {
			continue
		}
		_eegaf, _bfffc := _ad.GetName(_dfef.Get("\u0054\u0079\u0070\u0065"))
		if !_bfffc {
			continue
		}
		if *_eegaf != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_, _bfffc = _ad.GetName(_dfef.Get("\u004f\u0050\u0049"))
		if _bfffc {
			return _ba("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		_febf, _bfffc := _ad.GetName(_dfef.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"))
		if !_bfffc {
			continue
		}
		if *_febf == "\u0050\u0053" {
			return _ba("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		if _dfef.Get("\u0050\u0053") != nil {
			return _ba("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072m\u0020\u0058\u004f\u0062\u006a\u0065c\u0074\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u003a \u002d\u0020\u0074\u0068\u0065\u0020O\u0050\u0049\u0020\u006b\u0065\u0079\u003b \u002d\u0020\u0074\u0068e \u0053u\u0062\u0074\u0079\u0070\u0065\u0032 ke\u0079 \u0077\u0069t\u0068\u0020\u0061\u0020\u0076\u0061l\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u003b\u0020\u002d \u0074\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _daeb
}
func _eface(_bcfe, _efae, _fad, _bfge string) (string, bool) {
	_eab := _eg.Index(_bcfe, _efae)
	if _eab == -1 {
		return "", false
	}
	_gggbc := _eg.Index(_bcfe, _fad)
	if _gggbc == -1 {
		return "", false
	}
	if _gggbc < _eab {
		return "", false
	}
	return _bcfe[:_eab] + _efae + _bfge + _bcfe[_gggbc:], true
}
func _ggef(_cbfe *_da.CompliancePdfReader) (_cbcg []ViolatedRule) {
	var _aced, _egdg, _gfbc, _dagcc, _aaea, _degaa, _adggb bool
	_bdbec := map[*_ad.PdfObjectStream]struct{}{}
	for _, _ceaad := range _cbfe.GetObjectNums() {
		if _aced && _egdg && _aaea && _gfbc && _dagcc && _degaa && _adggb {
			return _cbcg
		}
		_bade, _gadca := _cbfe.GetIndirectObjectByNumber(_ceaad)
		if _gadca != nil {
			continue
		}
		_gaga, _bgbf := _ad.GetStream(_bade)
		if !_bgbf {
			continue
		}
		if _, _bgbf = _bdbec[_gaga]; _bgbf {
			continue
		}
		_bdbec[_gaga] = struct{}{}
		_fedea, _bgbf := _ad.GetName(_gaga.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_bgbf {
			continue
		}
		if !_dagcc {
			if _gaga.Get("\u0052\u0065\u0066") != nil {
				_cbcg = append(_cbcg, _ba("\u0036.\u0032\u002e\u0039\u002d\u0032", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e"))
				_dagcc = true
			}
		}
		if _fedea.String() == "\u0050\u0053" {
			if !_degaa {
				_cbcg = append(_cbcg, _ba("\u0036.\u0032\u002e\u0039\u002d\u0033", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073."))
				_degaa = true
				continue
			}
		}
		if _fedea.String() == "\u0046\u006f\u0072\u006d" {
			if _egdg && _gfbc && _dagcc {
				continue
			}
			if !_egdg && _gaga.Get("\u004f\u0050\u0049") != nil {
				_cbcg = append(_cbcg, _ba("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072\u006d \u0058\u004f\u0062j\u0065\u0063\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004f\u0050\u0049\u0020\u006b\u0065\u0079\u002e"))
				_egdg = true
			}
			if !_gfbc {
				if _gaga.Get("\u0050\u0053") != nil {
					_gfbc = true
				}
				if _egafb := _gaga.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _egafb != nil && !_gfbc {
					if _babe, _fedbc := _ad.GetName(_egafb); _fedbc && *_babe == "\u0050\u0053" {
						_gfbc = true
					}
				}
				if _gfbc {
					_cbcg = append(_cbcg, _ba("\u0036.\u0032\u002e\u0039\u002d\u0031", "\u0041\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065y \u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076\u0061\u006cu\u0065 o\u0066 \u0050\u0053\u0020\u0061\u006e\u0064\u0020t\u0068\u0065\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e"))
				}
			}
			continue
		}
		if _fedea.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if !_aced && _gaga.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
			_cbcg = append(_cbcg, _ba("\u0036.\u0032\u002e\u0038\u002d\u0031", "\u0041\u006e\u0020\u0049m\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
			_aced = true
		}
		if !_adggb && _gaga.Get("\u004f\u0050\u0049") != nil {
			_cbcg = append(_cbcg, _ba("\u0036.\u0032\u002e\u0038\u002d\u0032", "\u0041\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0068\u0065\u0020\u004f\u0050\u0049\u0020\u006b\u0065\u0079\u002e"))
			_adggb = true
		}
		if !_aaea && _gaga.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065") != nil {
			_degc, _eedc := _ad.GetBool(_gaga.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _eedc && bool(*_degc) {
				continue
			}
			_cbcg = append(_cbcg, _ba("\u0036.\u0032\u002e\u0038\u002d\u0033", "\u0049\u0066 a\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0063o\u006e\u0074\u0061\u0069n\u0073\u0020\u0074\u0068e \u0049\u006et\u0065r\u0070\u006f\u006c\u0061\u0074\u0065 \u006b\u0065\u0079,\u0020\u0069t\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020b\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
			_aaea = true
		}
	}
	return _cbcg
}
func _cdae(_fab *_ea.Document, _edc func() _a.Time) error {
	_ddac, _cfc := _da.NewPdfInfoFromObject(_fab.Info)
	if _cfc != nil {
		return _cfc
	}
	if _efg := _cfbb(_ddac, _edc); _efg != nil {
		return _efg
	}
	_fab.Info = _ddac.ToPdfObject()
	return nil
}
func _beb(_ecee *_da.CompliancePdfReader) ViolatedRule {
	if _ecee.ParserMetadata().HeaderPosition() != 0 {
		return _ba("\u0036.\u0031\u002e\u0032\u002d\u0031", "h\u0065\u0061\u0064\u0065\u0072\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0074\u0020\u0074\u0068\u0065\u0020fi\u0072\u0073\u0074 \u0062y\u0074\u0065")
	}
	return _bg
}
func _ceeeb(_ecba *_da.CompliancePdfReader) (_aceb []ViolatedRule) {
	_cdfg, _dgfaf := _acgab(_ecba)
	if !_dgfaf {
		return _aceb
	}
	_dffgd, _dgfaf := _ad.GetDict(_cdfg.Get("\u004e\u0061\u006de\u0073"))
	if !_dgfaf {
		return _aceb
	}
	if _dffgd.Get("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073") != nil {
		_aceb = append(_aceb, _ba("\u0036\u002e\u0031\u0030\u002d\u0031", "T\u0068\u0065\u0072e\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u006e\u006f\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0050\u0072\u0065s\u0065\u006e\u0074a\u0074\u0069\u006f\u006e\u0073\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u0027\u0073\u0020\u006e\u0061\u006d\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002e"))
	}
	return _aceb
}
func (_ebb *documentImages) hasOnlyDeviceCMYK() bool { return _ebb._fc && !_ebb._eea && !_ebb._eeb }
func _acag(_cdea *_da.CompliancePdfReader) (_geac []ViolatedRule) {
	_ggda, _ccbc := _acgab(_cdea)
	if !_ccbc {
		return _geac
	}
	_cbbd := _ba("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u0020\u0069\u0073\u0020a\u006e \u004f\u0075\u0074\u0070\u0075\u0074\u0049n\u0074\u0065\u006e\u0074\u0020\u0064i\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0062y\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0031\u0030.4\u002c\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073 \u0069\u006e\u0063\u006c\u0075\u0064e\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0027\u0073\u0020O\u0075\u0074p\u0075\u0074I\u006e\u0074\u0065\u006e\u0074\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020a\u006e\u0064\u0020h\u0061\u0073\u0020\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0041\u0031\u0020\u0061\u0073 \u0074\u0068\u0065\u0020\u0076a\u006c\u0075e\u0020\u006f\u0066\u0020i\u0074\u0073 \u0053\u0020\u006b\u0065\u0079\u0020\u0061\u006e\u0064\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020I\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006ce\u0020s\u0074\u0072\u0065\u0061\u006d \u0061\u0073\u0020\u0074h\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0074\u0073\u0020\u0044\u0065\u0073t\u004f\u0075t\u0070\u0075\u0074P\u0072\u006f\u0066\u0069\u006c\u0065 \u006b\u0065\u0079\u002e")
	_gcfg, _ccbc := _ad.GetArray(_ggda.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_ccbc {
		_geac = append(_geac, _cbbd)
		return _geac
	}
	_aece := _ba("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e")
	if _gcfg.Len() > 1 {
		_eade := map[*_ad.PdfObjectDictionary]struct{}{}
		for _fabd := 0; _fabd < _gcfg.Len(); _fabd++ {
			_fabb, _ccec := _ad.GetDict(_gcfg.Get(_fabd))
			if !_ccec {
				_geac = append(_geac, _cbbd)
				return _geac
			}
			if _fabd == 0 {
				_eade[_fabb] = struct{}{}
				continue
			}
			if _, _ffcda := _eade[_fabb]; !_ffcda {
				_geac = append(_geac, _aece)
				break
			}
		}
	} else if _gcfg.Len() == 0 {
		_geac = append(_geac, _cbbd)
		return _geac
	}
	_bgcf, _ccbc := _ad.GetDict(_gcfg.Get(0))
	if !_ccbc {
		_geac = append(_geac, _cbbd)
		return _geac
	}
	if _fccd, _bfgc := _ad.GetName(_bgcf.Get("\u0053")); !_bfgc || (*_fccd) != "\u0047T\u0053\u005f\u0050\u0044\u0046\u00411" {
		_geac = append(_geac, _cbbd)
		return _geac
	}
	_cebg, _bcaa := _da.NewPdfOutputIntentFromPdfObject(_bgcf)
	if _bcaa != nil {
		_ca.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020i\u006et\u0065\u006e\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _bcaa)
		return _geac
	}
	_bcef, _bcaa := _ac.ParseHeader(_cebg.DestOutputProfile)
	if _bcaa != nil {
		_ca.Log.Debug("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063\u006f\u006c\u006f\u0072\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061d\u0065\u0072\u0020\u0066\u0061i\u006c\u0065d\u003a\u0020\u0025\u0076", _bcaa)
		return _geac
	}
	if (_bcef.DeviceClass == _ac.DeviceClassPRTR || _bcef.DeviceClass == _ac.DeviceClassMNTR) && (_bcef.ColorSpace == _ac.ColorSpaceRGB || _bcef.ColorSpace == _ac.ColorSpaceCMYK || _bcef.ColorSpace == _ac.ColorSpaceGRAY) {
		return _geac
	}
	_geac = append(_geac, _cbbd)
	return _geac
}
func _acb(_ggb *_ea.Document) error {
	_gda := func(_bagc *_ad.PdfObjectDictionary) error {
		if _ed := _bagc.Get("\u0053\u004d\u0061s\u006b"); _ed != nil {
			_bagc.Set("\u0053\u004d\u0061s\u006b", _ad.MakeName("\u004e\u006f\u006e\u0065"))
		}
		_ddc := _bagc.Get("\u0043\u0041")
		if _ddc != nil {
			_ccg, _daf := _ad.GetNumberAsFloat(_ddc)
			if _daf != nil {
				_ca.Log.Debug("\u0045x\u0074\u0047S\u0074\u0061\u0074\u0065 \u006f\u0062\u006ae\u0063\u0074\u0020\u0043\u0041\u0020\u0076\u0061\u006cue\u0020\u0069\u0073 \u006e\u006ft\u0020\u0061\u0020\u0066\u006c\u006fa\u0074\u003a \u0025\u0076", _daf)
				_ccg = 0
			}
			if _ccg != 1.0 {
				_bagc.Set("\u0043\u0041", _ad.MakeFloat(1.0))
			}
		}
		_ddc = _bagc.Get("\u0063\u0061")
		if _ddc != nil {
			_cda, _agba := _ad.GetNumberAsFloat(_ddc)
			if _agba != nil {
				_ca.Log.Debug("\u0045\u0078t\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0027\u0063\u0061\u0027\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0066\u006c\u006f\u0061\u0074\u003a\u0020\u0025\u0076", _agba)
				_cda = 0
			}
			if _cda != 1.0 {
				_bagc.Set("\u0063\u0061", _ad.MakeFloat(1.0))
			}
		}
		_ce := _bagc.Get("\u0042\u004d")
		if _ce != nil {
			_gbd, _eaa := _ad.GetName(_ce)
			if !_eaa {
				_ca.Log.Debug("E\u0078\u0074\u0047\u0053\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0027\u0042\u004d\u0027\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061m\u0065")
				_gbd = _ad.MakeName("")
			}
			_gfeg := _gbd.String()
			switch _gfeg {
			case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065":
			default:
				_bagc.Set("\u0042\u004d", _ad.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
			}
		}
		_edd := _bagc.Get("\u0054\u0052")
		if _edd != nil {
			_ca.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0054\u0052\u0020\u006b\u0065\u0079")
			_bagc.Remove("\u0054\u0052")
		}
		_egc := _bagc.Get("\u0054\u0052\u0032")
		if _egc != nil {
			_dc := _egc.String()
			if _dc != "\u0044e\u0066\u0061\u0075\u006c\u0074" {
				_ca.Log.Debug("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074\u0065 o\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073 \u0054\u00522\u0020\u006b\u0065y\u0020\u0077\u0069\u0074\u0068\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065r\u0020\u0074ha\u006e\u0020\u0044e\u0066\u0061\u0075\u006c\u0074")
				_bagc.Set("\u0054\u0052\u0032", _ad.MakeName("\u0044e\u0066\u0061\u0075\u006c\u0074"))
			}
		}
		return nil
	}
	_aad, _efd := _ggb.GetPages()
	if !_efd {
		return nil
	}
	for _, _dbe := range _aad {
		_fgd, _bbf := _dbe.GetResources()
		if !_bbf {
			continue
		}
		_fgc, _feab := _ad.GetDict(_fgd.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
		if !_feab {
			return nil
		}
		_bead := _fgc.Keys()
		for _, _ebcc := range _bead {
			_cfa, _fggg := _ad.GetDict(_fgc.Get(_ebcc))
			if !_fggg {
				continue
			}
			_fbb := _gda(_cfa)
			if _fbb != nil {
				continue
			}
		}
	}
	for _, _gdc := range _aad {
		_dade, _acg := _gdc.GetContents()
		if !_acg {
			return nil
		}
		for _, _ggg := range _dade {
			_gdcg, _cde := _ggg.GetData()
			if _cde != nil {
				continue
			}
			_cbc := _cad.NewContentStreamParser(string(_gdcg))
			_eeca, _cde := _cbc.Parse()
			if _cde != nil {
				continue
			}
			for _, _fdf := range *_eeca {
				if len(_fdf.Params) == 0 {
					continue
				}
				_, _efb := _ad.GetName(_fdf.Params[0])
				if !_efb {
					continue
				}
				_dda, _ge := _gdc.GetResourcesXObject()
				if !_ge {
					continue
				}
				for _, _agf := range _dda.Keys() {
					_bcb, _bcf := _ad.GetStream(_dda.Get(_agf))
					if !_bcf {
						continue
					}
					_bfg, _bcf := _ad.GetDict(_bcb.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
					if !_bcf {
						continue
					}
					_adg, _bcf := _ad.GetDict(_bfg.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
					if !_bcf {
						continue
					}
					for _, _cdc := range _adg.Keys() {
						_gcc, _cee := _ad.GetDict(_adg.Get(_cdc))
						if !_cee {
							continue
						}
						_gbb := _gda(_gcc)
						if _gbb != nil {
							continue
						}
					}
				}
			}
		}
	}
	return nil
}
func _aaabb(_dbee *_ad.PdfObjectDictionary, _bdag map[*_ad.PdfObjectStream][]byte, _fadb map[*_ad.PdfObjectStream]*_b.CMap) ViolatedRule {
	const (
		_edcd  = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0033"
		_efgfe = "\u0041\u006c\u006c \u0043\u004d\u0061\u0070s\u0020\u0075\u0073ed\u0020\u0077\u0069\u0074\u0068i\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074 th\u006f\u0073\u0065\u0020\u006ci\u0073\u0074\u0065\u0064\u0020i\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u00320\u00300\u002d1\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0037\u002e\u0035\u002e\u0032\u002c\u0020\u0054\u0061\u0062\u006c\u0065 \u0031\u00318,\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0069\u006e \u0074\u0068\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u0020\u0061\u0073\u0020\u0064e\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u0032\u00300\u0030-\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u00209\u002e\u0037\u002e\u0035\u002e"
	)
	var _abec string
	if _cbgc, _edgaa := _ad.GetName(_dbee.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _edgaa {
		_abec = _cbgc.String()
	}
	if _abec != "\u0054\u0079\u0070e\u0030" {
		return _bg
	}
	_gfdca := _dbee.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _agdg, _bfbfd := _ad.GetName(_gfdca); _bfbfd {
		switch _agdg.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _bg
		default:
			return _ba(_edcd, _efgfe)
		}
	}
	_bdee, _edgda := _ad.GetStream(_gfdca)
	if !_edgda {
		return _ba(_edcd, _efgfe)
	}
	_, _cgbab := _ggag(_bdee, _bdag, _fadb)
	if _cgbab != nil {
		return _ba(_edcd, _efgfe)
	}
	return _bg
}
func _dfgdgd(_gfcg *_da.CompliancePdfReader) (_eged []ViolatedRule) {
	_egbe := true
	_aaefd, _cadf := _gfcg.GetCatalogMarkInfo()
	if !_cadf {
		_egbe = false
	} else {
		_aegd, _cagc := _ad.GetDict(_aaefd)
		if _cagc {
			_aaedb, _fbbb := _ad.GetBool(_aegd.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
			if !bool(*_aaedb) || !_fbbb {
				_egbe = false
			}
		} else {
			_egbe = false
		}
	}
	if !_egbe {
		_eged = append(_eged, _ba("\u0036.\u0037\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006cog\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020M\u0061r\u006b\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061 \u004d\u0061\u0072\u006b\u0065\u0064\u0020\u0065\u006et\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074,\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0076\u0061lu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0072\u0075\u0065"))
	}
	_gfbfd, _cadf := _gfcg.GetCatalogStructTreeRoot()
	if !_cadf {
		_eged = append(_eged, _ba("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075r\u0065\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065d \u0062\u0079\u0020a\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072e\u0020\u0068\u0069\u0065\u0072\u0061\u0072\u0063\u0068\u0079\u0020\u0072\u006f\u006ft\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0065\u006e\u0074r\u0079\u0020\u006f\u0066\u0020\u0074h\u0065\u0020d\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061t\u0061\u006c\u006fg \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0036\u002e"))
	}
	_bffb, _cadf := _ad.GetDict(_gfbfd)
	if _cadf {
		_cagfcg, _efca := _ad.GetName(_bffb.Get("\u0052o\u006c\u0065\u004d\u0061\u0070"))
		if _efca {
			_cdeg, _bbbe := _ad.GetDict(_cagfcg)
			if _bbbe {
				for _, _ggaa := range _cdeg.Keys() {
					_bdgb := _cdeg.Get(_ggaa)
					if _bdgb == nil {
						_eged = append(_eged, _ba("\u0036.\u0037\u002e\u0033\u002d\u0032", "\u0041\u006c\u006c\u0020\u006eo\u006e\u002ds\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0073t\u0072\u0075\u0063\u0074ure\u0020\u0074\u0079\u0070\u0065s\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u006d\u0061\u0070\u0070\u0065d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020n\u0065\u0061\u0072\u0065\u0073\u0074\u0020\u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0073\u0074a\u006ed\u0061r\u0064\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065re\u006e\u0063e\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020i\u006e\u0020\u0074\u0068e\u0020\u0072\u006fl\u0065\u0020\u006d\u0061p \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066 \u0074h\u0065\u0020\u0073\u0074\u0072\u0075c\u0074\u0075r\u0065\u0020\u0074\u0072e\u0065\u0020\u0072\u006f\u006ft\u002e"))
					}
				}
			}
		}
	}
	return _eged
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

func _daae(_dddb *_da.CompliancePdfReader) (_eece ViolatedRule) {
	_dfggd, _dead := _acgab(_dddb)
	if !_dead {
		return _bg
	}
	_agaa, _dead := _ad.GetDict(_dfggd.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
	if !_dead {
		return _bg
	}
	_geea, _dead := _ad.GetArray(_agaa.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_dead {
		return _bg
	}
	for _ffgcf := 0; _ffgcf < _geea.Len(); _ffgcf++ {
		_dedgb, _ffbf := _ad.GetDict(_geea.Get(_ffgcf))
		if !_ffbf {
			continue
		}
		if _dedgb.Get("\u0041") != nil {
			return _ba("\u0036.\u0034\u002e\u0031\u002d\u0032", "\u0041\u0020\u0046\u0069\u0065\u006c\u0064\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0041 o\u0072\u0020\u0041\u0041\u0020\u006b\u0065\u0079\u0073\u002e")
		}
		if _dedgb.Get("\u0041\u0041") != nil {
			return _ba("\u0036.\u0034\u002e\u0031\u002d\u0032", "\u0041\u0020\u0046\u0069\u0065\u006c\u0064\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0041 o\u0072\u0020\u0041\u0041\u0020\u006b\u0065\u0079\u0073\u002e")
		}
	}
	return _bg
}
func _ebaec(_abfg *_da.CompliancePdfReader) (_dcafa []ViolatedRule) {
	for _, _befd := range _abfg.GetObjectNums() {
		_ggfg, _ggac := _abfg.GetIndirectObjectByNumber(_befd)
		if _ggac != nil {
			continue
		}
		_gdaeg, _fgac := _ad.GetDict(_ggfg)
		if !_fgac {
			continue
		}
		_agffg, _fgac := _ad.GetName(_gdaeg.Get("\u0054\u0079\u0070\u0065"))
		if !_fgac {
			continue
		}
		if _agffg.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_bbec, _fgac := _ad.GetBool(_gdaeg.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if _fgac && bool(*_bbec) {
			_dcafa = append(_dcafa, _ba("\u0036.\u0034\u002e\u0031\u002d\u0033", "\u0054\u0068\u0065\u0020\u004e\u0065e\u0064\u0041\u0070\u0070\u0065a\u0072\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006c\u0061\u0067\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0069\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076e\u0020\u0066\u006f\u0072\u006d \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u006e\u006f\u0074\u0020b\u0065\u0020\u0070\u0072\u0065se\u006e\u0074\u0020\u006f\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
		}
		if _gdaeg.Get("\u0058\u0046\u0041") != nil {
			_dcafa = append(_dcafa, _ba("\u0036.\u0034\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020i\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076\u0065\u0020\u0066\u006f\u0072\u006d\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020t\u0068\u0061\u0074\u0020f\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065 \u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d \u006b\u0065\u0079\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006f\u0066 \u0061 \u0050\u0044F\u002fA\u002d\u0032\u0020\u0066ile\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065n\u0074\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0058\u0046\u0041\u0020\u006b\u0065y."))
		}
	}
	_aded, _ccccg := _acgab(_abfg)
	if _ccccg && _aded.Get("\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067") != nil {
		_dcafa = append(_dcafa, _ba("\u0036.\u0034\u002e\u0032\u002d\u0032", "\u0041\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020\u0043\u0061\u0074\u0061\u006cog\u0020s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004e\u0065\u0065\u0064\u0073\u0052\u0065\u006e\u0064e\u0072\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u002e"))
	}
	return _dcafa
}
func _abgc(_cbge *_ad.PdfObjectDictionary, _abgd map[*_ad.PdfObjectStream][]byte, _ffdfe map[*_ad.PdfObjectStream]*_b.CMap) ViolatedRule {
	const (
		_gefdb = "\u0036.\u0033\u002e\u0033\u002d\u0033"
		_bbde  = "\u0041\u006cl \u0043\u004d\u0061\u0070\u0073\u0020\u0075\u0073e\u0064 \u0077i\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072m\u0069n\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020a\u006e\u0064\u0020\u0049\u0064\u0065\u006et\u0069\u0074\u0079-\u0056\u002c\u0020\u0073\u0068a\u006c\u006c \u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0074h\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u0020\u0061\u0073\u0020\u0064es\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u00205\u002e\u0036\u002e\u0034\u002e"
	)
	var _gfffb string
	if _egbb, _aceec := _ad.GetName(_cbge.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _aceec {
		_gfffb = _egbb.String()
	}
	if _gfffb != "\u0054\u0079\u0070e\u0030" {
		return _bg
	}
	_ecfc := _cbge.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _efbcf, _dgfe := _ad.GetName(_ecfc); _dgfe {
		switch _efbcf.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _bg
		default:
			return _ba(_gefdb, _bbde)
		}
	}
	_cgeg, _dgde := _ad.GetStream(_ecfc)
	if !_dgde {
		return _ba(_gefdb, _bbde)
	}
	_, _effb := _ggag(_cgeg, _abgd, _ffdfe)
	if _effb != nil {
		return _ba(_gefdb, _bbde)
	}
	return _bg
}
func (_cd *documentImages) hasOnlyDeviceRGB() bool { return _cd._eea && !_cd._fc && !_cd._eeb }

var _ Profile = (*Profile2B)(nil)

// StandardName gets the name of the standard.
func (_agcc *profile2) StandardName() string {
	return _d.Sprintf("\u0050D\u0046\u002f\u0041\u002d\u0032\u0025s", _agcc._fcdd._ee)
}
func _ggag(_afcb *_ad.PdfObjectStream, _facd map[*_ad.PdfObjectStream][]byte, _gabde map[*_ad.PdfObjectStream]*_b.CMap) (*_b.CMap, error) {
	_aaee, _gdcb := _gabde[_afcb]
	if !_gdcb {
		var _dcff error
		_egad, _ggdf := _facd[_afcb]
		if !_ggdf {
			_egad, _dcff = _ad.DecodeStream(_afcb)
			if _dcff != nil {
				_ca.Log.Debug("\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0073\u0074r\u0065\u0061\u006d\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _dcff)
				return nil, _dcff
			}
			_facd[_afcb] = _egad
		}
		_aaee, _dcff = _b.LoadCmapFromData(_egad, false)
		if _dcff != nil {
			return nil, _dcff
		}
		_gabde[_afcb] = _aaee
	}
	return _aaee, nil
}
func _df() standardType { return standardType{_efe: 1, _ee: "\u0041"} }
func _fdbf(_efaf *_da.CompliancePdfReader) ViolatedRule {
	if _efaf.ParserMetadata().HeaderPosition() != 0 {
		return _ba("\u0036.\u0031\u002e\u0032\u002d\u0031", "h\u0065\u0061\u0064\u0065\u0072\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0074\u0020\u0074\u0068\u0065\u0020fi\u0072\u0073\u0074 \u0062y\u0074\u0065")
	}
	if _efaf.PdfVersion().Major != 1 {
		return _ba("\u0036.\u0031\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0065\u0061\u0064e\u0072 \u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0073\u0069s\u0074 \u006f\u0066\u0020\u201c%\u0050\u0044\u0046\u002d\u0031\u002e\u006e\u201d\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065 \u0045\u004f\u004c\u0020ma\u0072\u006b\u0065\u0072\u002c \u0077\u0068\u0065\u0072\u0065\u0020\u0027\u006e\u0027\u0020\u0069s\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0064\u0069\u0067\u0069t\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0030\u0020(\u0033\u0030h\u0029\u0020\u0061\u006e\u0064\u0020\u0037\u0020\u0028\u0033\u0037\u0068\u0029")
	}
	if _efaf.PdfVersion().Minor < 0 || _efaf.PdfVersion().Minor > 7 {
		return _ba("\u0036.\u0031\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0065\u0061\u0064e\u0072 \u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0073\u0069s\u0074 \u006f\u0066\u0020\u201c%\u0050\u0044\u0046\u002d\u0031\u002e\u006e\u201d\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065 \u0045\u004f\u004c\u0020ma\u0072\u006b\u0065\u0072\u002c \u0077\u0068\u0065\u0072\u0065\u0020\u0027\u006e\u0027\u0020\u0069s\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0064\u0069\u0067\u0069t\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0030\u0020(\u0033\u0030h\u0029\u0020\u0061\u006e\u0064\u0020\u0037\u0020\u0028\u0033\u0037\u0068\u0029")
	}
	return _bg
}
func _edga(_ddfd standardType, _bfagc *_ea.OutputIntents) error {
	_bcde, _fcfdf := _ac.NewCmykIsoCoatedV2OutputIntent(_ddfd.outputIntentSubtype())
	if _fcfdf != nil {
		return _fcfdf
	}
	if _fcfdf = _bfagc.Add(_bcde.ToPdfObject()); _fcfdf != nil {
		return _fcfdf
	}
	return nil
}
func _geddd(_edda *_da.CompliancePdfReader) ViolatedRule { return _bg }
func _agc(_ced *_ea.Document) error {
	_dae, _dbfc := _ced.FindCatalog()
	if !_dbfc {
		return _gg.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_dae.SetVersion()
	return nil
}
func _gdaca(_efaa *_da.PdfFont, _ebfed *_ad.PdfObjectDictionary, _dbead bool) ViolatedRule {
	const (
		_ceegg = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0034\u002d\u0031"
		_bcffg = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006ds\u0020\u0066\u006fr\u0020\u0061\u006c\u006c\u0020f\u006f\u006e\u0074\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0072e\u006e\u0064\u0065\u0072\u0069\u006eg\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020w\u0069t\u0068\u0069\u006e\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u0069\u006c\u0065\u002c \u0061\u0073\u0020\u0064\u0065\u0066\u0069n\u0065\u0064 \u0069\u006e\u0020\u0049S\u004f\u0020\u0033\u0032\u00300\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0039\u002e"
	)
	if _dbead {
		return _bg
	}
	_ecfe := _efaa.FontDescriptor()
	var _cbced string
	if _fbae, _adba := _ad.GetName(_ebfed.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _adba {
		_cbced = _fbae.String()
	}
	switch _cbced {
	case "\u0054\u0079\u0070e\u0031":
		if _ecfe.FontFile == nil {
			return _ba(_ceegg, _bcffg)
		}
	case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		if _ecfe.FontFile2 == nil {
			return _ba(_ceegg, _bcffg)
		}
	case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0033":
	default:
		if _ecfe.FontFile3 == nil {
			return _ba(_ceegg, _bcffg)
		}
	}
	return _bg
}

// NewProfile2U creates a new Profile2U with the given options.
func NewProfile2U(options *Profile2Options) *Profile2U {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_cdcc(options)
	return &Profile2U{profile2{_eceb: *options, _fcdd: _eec()}}
}

type profile1 struct {
	_efde standardType
	_fcce Profile1Options
}

func _cfbb(_ddfb *_da.PdfInfo, _dbdd func() _a.Time) error {
	var _cbce *_da.PdfDate
	if _ddfb.CreationDate == nil {
		_ceee, _defa := _da.NewPdfDateFromTime(_dbdd())
		if _defa != nil {
			return _defa
		}
		_cbce = &_ceee
		_ddfb.CreationDate = _cbce
	}
	if _ddfb.ModifiedDate == nil {
		if _cbce != nil {
			_baab, _dccd := _da.NewPdfDateFromTime(_dbdd())
			if _dccd != nil {
				return _dccd
			}
			_cbce = &_baab
		}
		_ddfb.ModifiedDate = _cbce
	}
	return nil
}

// Conformance gets the PDF/A conformance.
func (_gbge *profile2) Conformance() string { return _gbge._fcdd._ee }

// NewProfile1B creates a new Profile1B with the given options.
func NewProfile1B(options *Profile1Options) *Profile1B {
	if options == nil {
		options = DefaultProfile1Options()
	}
	_aegab(options)
	return &Profile1B{profile1{_fcce: *options, _efde: _ebc()}}
}

type documentImages struct {
	_eea, _fc, _eeb bool
	_dbf            map[_ad.PdfObject]struct{}
	_fca            []*imageInfo
}

func _feda(_gaceb *_da.CompliancePdfReader) ViolatedRule {
	for _, _ddfbe := range _gaceb.PageList {
		_eeed, _agfbb := _ddfbe.GetContentStreams()
		if _agfbb != nil {
			continue
		}
		for _, _cgfd := range _eeed {
			_gggcb := _cad.NewContentStreamParser(_cgfd)
			_, _agfbb = _gggcb.Parse()
			if _agfbb != nil {
				return _ba("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0031", "\u0041\u0020\u0063onten\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079 \u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u0020\u006e\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0065\u0076\u0065\u006e\u0020\u0069\u0066\u0020s\u0075\u0063\u0068\u0020\u006f\u0070\u0065r\u0061\u0074\u006f\u0072\u0073\u0020\u0061\u0072\u0065\u0020\u0062\u0072\u0061\u0063\u006b\u0065\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0074\u0068\u0065\u0020\u0042\u0058\u002f\u0045\u0058\u0020\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u002e")
			}
		}
	}
	return _bg
}
func _bdbeb(_egea *_da.CompliancePdfReader) (_dcdbe []ViolatedRule) {
	for _, _gbgc := range _egea.GetObjectNums() {
		_fcaa, _cdec := _egea.GetIndirectObjectByNumber(_gbgc)
		if _cdec != nil {
			continue
		}
		_fagea, _fgcg := _ad.GetDict(_fcaa)
		if !_fgcg {
			continue
		}
		_badb, _fgcg := _ad.GetName(_fagea.Get("\u0054\u0079\u0070\u0065"))
		if !_fgcg {
			continue
		}
		if _badb.String() != "\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d" {
			continue
		}
		_fegf, _fgcg := _ad.GetBool(_fagea.Get("\u004ee\u0065d\u0041\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0073"))
		if !_fgcg {
			return _dcdbe
		}
		if bool(*_fegf) {
			_dcdbe = append(_dcdbe, _ba("\u0036\u002e\u0039-\u0031", "\u0054\u0068\u0065\u0020\u004e\u0065e\u0064\u0041\u0070\u0070\u0065a\u0072\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006c\u0061\u0067\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0069\u006e\u0074\u0065\u0072\u0061\u0063\u0074\u0069\u0076e\u0020\u0066\u006f\u0072\u006d \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u006e\u006f\u0074\u0020b\u0065\u0020\u0070\u0072\u0065se\u006e\u0074\u0020\u006f\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
		}
	}
	return _dcdbe
}
func _gdga(_bega *_ea.Document, _gaa *_ea.Page, _bbdd []*_ea.Image) error {
	for _, _eded := range _bbdd {
		if _eded.SMask == nil {
			continue
		}
		_eccf, _dca := _da.NewXObjectImageFromStream(_eded.Stream)
		if _dca != nil {
			return _dca
		}
		_gbac, _dca := _eccf.ToImage()
		if _dca != nil {
			return _dca
		}
		_fbee, _dca := _gbac.ToGoImage()
		if _dca != nil {
			return _dca
		}
		_agfb, _dca := _de.RGBAConverter.Convert(_fbee)
		if _dca != nil {
			return _dca
		}
		_affa := _agfb.Base()
		_gccg := &_da.Image{Width: int64(_affa.Width), Height: int64(_affa.Height), BitsPerComponent: int64(_affa.BitsPerComponent), ColorComponents: _affa.ColorComponents, Data: _affa.Data}
		_gccg.SetDecode(_affa.Decode)
		_gccg.SetAlpha(_affa.Alpha)
		if _dca = _eccf.SetImage(_gccg, nil); _dca != nil {
			return _dca
		}
		_eccf.SMask = _ad.MakeNull()
		var _ffbg _ad.PdfObject
		_dfae := -1
		for _dfae, _ffbg = range _bega.Objects {
			if _ffbg == _eded.SMask.Stream {
				break
			}
		}
		if _dfae != -1 {
			_bega.Objects = append(_bega.Objects[:_dfae], _bega.Objects[_dfae+1:]...)
		}
		_eded.SMask = nil
		_eccf.ToPdfObject()
	}
	return nil
}
func _fgafc(_dcfb *_ad.PdfObjectDictionary, _ddgfd map[*_ad.PdfObjectStream][]byte, _adf map[*_ad.PdfObjectStream]*_b.CMap) ViolatedRule {
	const (
		_fdcff = "\u0046\u006f\u0072 \u0061\u006e\u0079\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0063\u006f\u006d\u0070\u006f\u0073\u0069\u0074\u0065\u0020\u0028\u0054\u0079\u0070\u0065\u0020\u0030\u0029\u0020\u0066\u006f\u006et \u0072\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0064 \u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006fn\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0043I\u0044\u0053y\u0073\u0074\u0065\u006d\u0049nf\u006f\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u006f\u0066\u0020i\u0074\u0073\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0061\u006e\u0064 \u0043\u004d\u0061\u0070 \u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0063\u006f\u006d\u0070\u0061\u0074i\u0062\u006c\u0065\u002e\u0020\u0049\u006e\u0020o\u0074\u0068\u0065\u0072\u0020\u0077\u006f\u0072\u0064\u0073\u002c\u0020\u0074\u0068\u0065\u0020R\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u0020a\u006e\u0064\u0020\u004fr\u0064\u0065\u0072\u0069\u006e\u0067 \u0073\u0074\u0072i\u006e\u0067\u0073\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0066\u006f\u0072 \u0074\u0068\u0061\u0074\u0020\u0066o\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u002c\u0020u\u006el\u0065ss \u0074\u0068\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006eg\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074h\u0065 \u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0069\u0073 \u0049\u0064\u0065\u006e\u0074\u0069t\u0079\u002d\u0048\u0020o\u0072\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074y\u002dV\u002e"
		_gggd  = "\u0036.\u0033\u002e\u0033\u002d\u0031"
	)
	var _ebfab string
	if _cgbe, _faed := _ad.GetName(_dcfb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _faed {
		_ebfab = _cgbe.String()
	}
	if _ebfab != "\u0054\u0079\u0070e\u0030" {
		return _bg
	}
	_cdbfc := _dcfb.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _caedg, _aeda := _ad.GetName(_cdbfc); _aeda {
		switch _caedg.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _bg
		}
		_ccef, _bfff := _b.LoadPredefinedCMap(_caedg.String())
		if _bfff != nil {
			return _ba(_gggd, _fdcff)
		}
		_gcab := _ccef.CIDSystemInfo()
		if _gcab.Ordering != _gcab.Registry {
			return _ba(_gggd, _fdcff)
		}
		return _bg
	}
	_bgdbgb, _adga := _ad.GetStream(_cdbfc)
	if !_adga {
		return _ba(_gggd, _fdcff)
	}
	_fbbeb, _ecbd := _ggag(_bgdbgb, _ddgfd, _adf)
	if _ecbd != nil {
		return _ba(_gggd, _fdcff)
	}
	_eafc := _fbbeb.CIDSystemInfo()
	if _eafc.Ordering != _eafc.Registry {
		return _ba(_gggd, _fdcff)
	}
	return _bg
}
func _eadd(_ebg *_ea.Document, _fde bool) error {
	_dece, _deef := _ebg.GetPages()
	if !_deef {
		return nil
	}
	for _, _eeg := range _dece {
		_gebd, _eadf := _eeg.GetContents()
		if !_eadf {
			continue
		}
		var _bda *_da.PdfPageResources
		_dga, _eadf := _eeg.GetResources()
		if _eadf {
			_bda, _ = _da.NewPdfPageResourcesFromDict(_dga)
		}
		for _gfc, _fag := range _gebd {
			_ceac, _edec := _fag.GetData()
			if _edec != nil {
				continue
			}
			_bff := _cad.NewContentStreamParser(string(_ceac))
			_aceg, _edec := _bff.Parse()
			if _edec != nil {
				continue
			}
			_fcfd, _edec := _dge(_bda, _aceg, _fde)
			if _edec != nil {
				return _edec
			}
			if _fcfd == nil {
				continue
			}
			if _edec = (&_gebd[_gfc]).SetData(_fcfd); _edec != nil {
				return _edec
			}
		}
	}
	return nil
}

// NewProfile1A creates a new Profile1A with given options.
func NewProfile1A(options *Profile1Options) *Profile1A {
	if options == nil {
		options = DefaultProfile1Options()
	}
	_aegab(options)
	return &Profile1A{profile1{_fcce: *options, _efde: _df()}}
}
func _ddccc(_eadc *_da.PdfFont, _abdg *_ad.PdfObjectDictionary) ViolatedRule {
	const (
		_dgeb = "\u0036.\u0033\u002e\u0035\u002d\u0032"
		_caeb = "\u0046\u006f\u0072\u0020\u0061l\u006c\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074 \u0073\u0075bs\u0065\u0074\u0073 \u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0061\u0020\u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074he\u0020f\u006f\u006e\u0074\u0020\u0064\u0065s\u0063r\u0069\u0070\u0074o\u0072\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006ec\u006c\u0075\u0064e\u0020\u0061\u0020\u0043\u0068\u0061\u0072\u0053\u0065\u0074\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u006c\u0069\u0073\u0074\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0063\u0068\u0061\u0072a\u0063\u0074\u0065\u0072 \u006e\u0061\u006d\u0065\u0073\u0020d\u0065\u0066i\u006e\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020f\u006f\u006e\u0074\u0020s\u0075\u0062\u0073\u0065\u0074, \u0061\u0073 \u0064\u0065s\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e \u0050\u0044\u0046\u0020\u0052e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0054\u0061\u0062\u006ce\u0020\u0035\u002e1\u0038\u002e"
	)
	var _bebd string
	if _cdee, _cagec := _ad.GetName(_abdg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cagec {
		_bebd = _cdee.String()
	}
	if _bebd != "\u0054\u0079\u0070e\u0031" {
		return _bg
	}
	if _cac.IsStdFont(_cac.StdFontName(_eadc.BaseFont())) {
		return _bg
	}
	_ffag := _eadc.FontDescriptor()
	if _ffag.CharSet == nil {
		return _ba(_dgeb, _caeb)
	}
	return _bg
}
func _faea(_dcgfd *_da.PdfFont, _ggce *_ad.PdfObjectDictionary) ViolatedRule {
	const (
		_edadc = "\u0036.\u0033\u002e\u0035\u002d\u0033"
		_gedd  = "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0073 \u0072e\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0077i\u0074\u0068\u0069n\u0020\u0061\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069l\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006et\u0020\u0064\u0065s\u0063\u0072\u0069\u0070\u0074\u006f\u0072\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020\u0043\u0049\u0044\u0053\u0065\u0074\u0020s\u0074\u0072\u0065\u0061\u006d\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0079\u0069\u006eg\u0020\u0077\u0068i\u0063\u0068\u0020\u0043\u0049\u0044\u0073 \u0061\u0072e\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e \u0074\u0068\u0065\u0020\u0065\u006d\u0062\u0065\u0064d\u0065\u0064\u0020\u0043\u0049D\u0046\u006f\u006e\u0074\u0020\u0066\u0069l\u0065,\u0020\u0061\u0073 \u0064\u0065\u0073\u0063\u0072\u0069b\u0065\u0064 \u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063e\u0020\u0054ab\u006c\u0065\u0020\u0035.\u00320\u002e"
	)
	var _gfcd string
	if _aeca, _gafg := _ad.GetName(_ggce.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _gafg {
		_gfcd = _aeca.String()
	}
	switch _gfcd {
	case "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
		_egdb := _dcgfd.FontDescriptor()
		if _egdb.CIDSet == nil {
			return _ba(_edadc, _gedd)
		}
		return _bg
	default:
		return _bg
	}
}
func _afcg(_gcgd *_da.CompliancePdfReader) (_cdga []ViolatedRule) {
	var _ebga, _ceegd, _cadg, _bdac, _bbfa, _fceg, _bfeb bool
	_cacd := func() bool { return _ebga && _ceegd && _cadg && _bdac && _bbfa && _fceg && _bfeb }
	for _, _bafge := range _gcgd.PageList {
		_cgcd, _ebaad := _bafge.GetAnnotations()
		if _ebaad != nil {
			_ca.Log.Trace("\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006es\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ebaad)
			continue
		}
		for _, _bdde := range _cgcd {
			if !_ebga {
				switch _bdde.GetContext().(type) {
				case *_da.PdfAnnotationFileAttachment, *_da.PdfAnnotationSound, *_da.PdfAnnotationMovie, nil:
					_cdga = append(_cdga, _ba("\u0036.\u0035\u002e\u0032\u002d\u0031", "\u0041\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0079\u0070\u0065\u0073\u0020\u006e\u006f\u0074 \u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020p\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e\u0020\u0041d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020\u0074\u0068\u0065\u0020F\u0069\u006c\u0065\u0041\u0074\u0074\u0061\u0063\u0068\u006de\u006e\u0074\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u0020\u0061\u006e\u0064\u0020\u004d\u006f\u0076\u0069e\u0020\u0074\u0079\u0070\u0065s \u0073ha\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_ebga = true
					if _cacd() {
						return _cdga
					}
				}
			}
			_dgdc, _adec := _ad.GetDict(_bdde.GetContainingPdfObject())
			if !_adec {
				continue
			}
			if !_ceegd {
				_cegb, _dceg := _ad.GetFloatVal(_dgdc.Get("\u0043\u0041"))
				if _dceg && _cegb != 1.0 {
					_cdga = append(_cdga, _ba("\u0036.\u0035\u002e\u0033\u002d\u0031", "\u0041\u006e\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073h\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0043\u0041\u0020\u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0031\u002e\u0030\u002e"))
					_ceegd = true
					if _cacd() {
						return _cdga
					}
				}
			}
			if !_cadg {
				_fbcd, _eggb := _ad.GetIntVal(_dgdc.Get("\u0046"))
				if !(_eggb && _fbcd&4 == 4 && _fbcd&1 == 0 && _fbcd&2 == 0 && _fbcd&32 == 0) {
					_cdga = append(_cdga, _ba("\u0036.\u0035\u002e\u0033\u002d\u0032", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0074\u0068\u0065\u0020\u0046\u0020\u006b\u0065\u0079\u002e\u0020\u0054\u0068\u0065\u0020\u0046\u0020\u006b\u0065\u0079\u0027\u0073\u0020\u0050\u0072\u0069\u006e\u0074\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065 s\u0065\u0074\u0020\u0074\u006f\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0048\u0069\u0064\u0064\u0065\u006e\u002c\u0020I\u006e\u0076\u0069\u0073\u0069\u0062\u006c\u0065\u0020\u0061\u006e\u0064\u0020\u004e\u006f\u0056\u0069\u0065\u0077\u0020\u0066\u006c\u0061\u0067\u0020b\u0069\u0074\u0073 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073e\u0074\u0020t\u006f\u0020\u0030\u002e"))
					_cadg = true
					if _cacd() {
						return _cdga
					}
				}
			}
			if !_bdac {
				_afdcg, _addc := _ad.GetDict(_dgdc.Get("\u0041\u0050"))
				if _addc {
					_fdcg := _afdcg.Get("\u004e")
					if _fdcg == nil || len(_afdcg.Keys()) > 1 {
						_cdga = append(_cdga, _ba("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_bdac = true
						if _cacd() {
							return _cdga
						}
						continue
					}
					_, _dcdgc := _bdde.GetContext().(*_da.PdfAnnotationWidget)
					if _dcdgc {
						_becd, _fgff := _ad.GetName(_dgdc.Get("\u0046\u0054"))
						if _fgff && *_becd == "\u0042\u0074\u006e" {
							if _, _debg := _ad.GetDict(_fdcg); !_debg {
								_cdga = append(_cdga, _ba("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
								_bdac = true
								if _cacd() {
									return _cdga
								}
								continue
							}
						}
					}
					_, _dgcb := _ad.GetStream(_fdcg)
					if !_dgcb {
						_cdga = append(_cdga, _ba("\u0036.\u0035\u002e\u0033\u002d\u0034", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_bdac = true
						if _cacd() {
							return _cdga
						}
						continue
					}
				}
			}
			if !_bbfa {
				if _dgdc.Get("\u0043") != nil || _dgdc.Get("\u0049\u0043") != nil {
					_gbgf, _efcgc := _caee(_gcgd)
					if !_efcgc {
						_cdga = append(_cdga, _ba("\u0036.\u0035\u002e\u0033\u002d\u0033", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074a\u0069\u006e\u0020t\u0068e\u0020\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0072\u0020\u0074\u0068e\u0020\u0049\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0075\u006e\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006ff\u0069\u006ce\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069n\u0020\u0036\u002e\u0032\u002e2\u002c\u0020\u0069\u0073\u0020\u0052\u0047\u0042."))
						_bbfa = true
						if _cacd() {
							return _cdga
						}
					} else {
						_cbad, _dbagc := _ad.GetIntVal(_gbgf.Get("\u004e"))
						if !_dbagc || _cbad != 3 {
							_cdga = append(_cdga, _ba("\u0036.\u0035\u002e\u0033\u002d\u0033", "\u0041\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006fn\u0074a\u0069\u006e\u0020t\u0068e\u0020\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0072\u0020\u0074\u0068e\u0020\u0049\u0043\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0075\u006e\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006ff\u0069\u006ce\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069n\u0020\u0036\u002e\u0032\u002e2\u002c\u0020\u0069\u0073\u0020\u0052\u0047\u0042."))
							_bbfa = true
							if _cacd() {
								return _cdga
							}
						}
					}
				}
			}
			_gefe, _fffg := _bdde.GetContext().(*_da.PdfAnnotationWidget)
			if !_fffg {
				continue
			}
			if !_fceg {
				if _gefe.A != nil {
					_cdga = append(_cdga, _ba("\u0036.\u0036\u002e\u0031\u002d\u0033", "A \u0057\u0069d\u0067\u0065\u0074\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006ec\u006cu\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0020e\u006et\u0072\u0079."))
					_fceg = true
					if _cacd() {
						return _cdga
					}
				}
			}
			if !_bfeb {
				if _gefe.AA != nil {
					_cdga = append(_cdga, _ba("\u0036.\u0036\u002e\u0032\u002d\u0031", "\u0041\u0020\u0057\u0069\u0064\u0067\u0065\u0074\u0020\u0061\u006e\u006eo\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073h\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					_bfeb = true
					if _cacd() {
						return _cdga
					}
				}
			}
		}
	}
	return _cdga
}

// StandardName gets the name of the standard.
func (_bgfc *profile1) StandardName() string {
	return _d.Sprintf("\u0050D\u0046\u002f\u0041\u002d\u0031\u0025s", _bgfc._efde._ee)
}
func _facca(_cadad *_ea.Document) error {
	_gfa, _ecda := _cadad.GetPages()
	if !_ecda {
		return nil
	}
	for _, _dbcd := range _gfa {
		_fbbe, _bcff := _ad.GetArray(_dbcd.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_bcff {
			continue
		}
		for _, _fgb := range _fbbe.Elements() {
			_fgb = _ad.ResolveReference(_fgb)
			if _, _bdab := _fgb.(*_ad.PdfObjectNull); _bdab {
				continue
			}
			_cbab, _cef := _ad.GetDict(_fgb)
			if !_cef {
				continue
			}
			_eefb, _ := _ad.GetIntVal(_cbab.Get("\u0046"))
			_eefb &= ^(1 << 0)
			_eefb &= ^(1 << 1)
			_eefb &= ^(1 << 5)
			_eefb &= ^(1 << 8)
			_eefb |= 1 << 2
			_cbab.Set("\u0046", _ad.MakeInteger(int64(_eefb)))
			_cggc := false
			if _aed := _cbab.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _aed != nil {
				_fgfb, _gfef := _ad.GetName(_aed)
				if _gfef && _fgfb.String() == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					_cggc = true
					if _cbab.Get("\u0041\u0041") != nil {
						_cbab.Remove("\u0041\u0041")
					}
					if _cbab.Get("\u0041") != nil {
						_cbab.Remove("\u0041")
					}
				}
				if _gfef && _fgfb.String() == "\u0054\u0065\u0078\u0074" {
					_dfbf, _ := _ad.GetIntVal(_cbab.Get("\u0046"))
					_dfbf |= 1 << 3
					_dfbf |= 1 << 4
					_cbab.Set("\u0046", _ad.MakeInteger(int64(_dfbf)))
				}
			}
			_dbddg, _cef := _ad.GetDict(_cbab.Get("\u0041\u0050"))
			if _cef {
				_edde := _dbddg.Get("\u004e")
				if _edde == nil {
					continue
				}
				if len(_dbddg.Keys()) > 1 {
					_dbddg.Clear()
					_dbddg.Set("\u004e", _edde)
				}
				if _cggc {
					_cdce, _ddacb := _ad.GetName(_cbab.Get("\u0046\u0054"))
					if _ddacb && *_cdce == "\u0042\u0074\u006e" {
						continue
					}
				}
			}
		}
	}
	return nil
}

// DefaultProfile1Options are the default options for the Profile1.
func DefaultProfile1Options() *Profile1Options {
	return &Profile1Options{Now: _a.Now, Xmp: XmpOptions{MarshalIndent: "\u0009"}}
}
func _ff(_dfg []_ad.PdfObject) (*documentImages, error) {
	_aac := _ad.PdfObjectName("\u0053u\u0062\u0074\u0079\u0070\u0065")
	_bgb := make(map[*_ad.PdfObjectStream]struct{})
	_dbc := make(map[_ad.PdfObject]struct{})
	var (
		_dadg, _def, _ag bool
		_caf             []*imageInfo
		_dee             error
	)
	for _, _fa := range _dfg {
		_gad, _ecf := _ad.GetStream(_fa)
		if !_ecf {
			continue
		}
		if _, _gadb := _bgb[_gad]; _gadb {
			continue
		}
		_bgb[_gad] = struct{}{}
		_fcf := _gad.PdfObjectDictionary.Get(_aac)
		_gd, _ecf := _ad.GetName(_fcf)
		if !_ecf || string(*_gd) != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if _bag := _gad.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b"); _bag != nil {
			_dbc[_bag] = struct{}{}
		}
		_cge := imageInfo{BitsPerComponent: 8, Stream: _gad}
		_cge.ColorSpace, _dee = _da.DetermineColorspaceNameFromPdfObject(_gad.PdfObjectDictionary.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"))
		if _dee != nil {
			return nil, _dee
		}
		if _dd, _cdb := _ad.GetIntVal(_gad.PdfObjectDictionary.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _cdb {
			_cge.BitsPerComponent = _dd
		}
		if _bd, _agb := _ad.GetIntVal(_gad.PdfObjectDictionary.Get("\u0057\u0069\u0064t\u0068")); _agb {
			_cge.Width = _bd
		}
		if _gab, _dfc := _ad.GetIntVal(_gad.PdfObjectDictionary.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _dfc {
			_cge.Height = _gab
		}
		switch _cge.ColorSpace {
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
			_ag = true
			_cge.ColorComponents = 1
		case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
			_dadg = true
			_cge.ColorComponents = 3
		case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
			_def = true
			_cge.ColorComponents = 4
		default:
			_cge._dfb = true
		}
		_caf = append(_caf, &_cge)
	}
	if len(_dbc) > 0 {
		if len(_dbc) == len(_caf) {
			_caf = nil
		} else {
			_bf := make([]*imageInfo, len(_caf)-len(_dbc))
			var _fd int
			for _, _bgf := range _caf {
				if _, _gaf := _dbc[_bgf.Stream]; _gaf {
					continue
				}
				_bf[_fd] = _bgf
				_fd++
			}
			_caf = _bf
		}
	}
	return &documentImages{_eea: _dadg, _fc: _def, _eeb: _ag, _dbf: _dbc, _fca: _caf}, nil
}
func _dfdffa(_abaff *_ad.PdfObjectDictionary, _edbfa map[*_ad.PdfObjectStream][]byte, _edcad map[*_ad.PdfObjectStream]*_b.CMap) ViolatedRule {
	const (
		_daab = "\u0036.\u0033\u002e\u0033\u002d\u0034"
		_dbad = "\u0046\u006f\u0072\u0020\u0074\u0068\u006fs\u0065\u0020\u0043\u004d\u0061\u0070\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072e\u0020\u0065m\u0062\u0065\u0064de\u0064\u002c\u0020\u0074\u0068\u0065\u0020\u0069\u006et\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0057\u004d\u006f\u0064\u0065\u0020\u0065\u006e\u0074r\u0079\u0020i\u006e t\u0068\u0065\u0020CM\u0061\u0070\u0020\u0064\u0069\u0063\u0074\u0069o\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u0020\u0074\u006f \u0074h\u0065\u0020\u0057\u004d\u006f\u0064e\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064ed\u0020\u0043\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"
	)
	var _adgbcc string
	if _adce, _fcdda := _ad.GetName(_abaff.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _fcdda {
		_adgbcc = _adce.String()
	}
	if _adgbcc != "\u0054\u0079\u0070e\u0030" {
		return _bg
	}
	_aedfa := _abaff.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _, _ffcbf := _ad.GetName(_aedfa); _ffcbf {
		return _bg
	}
	_fgag, _ade := _ad.GetStream(_aedfa)
	if !_ade {
		return _ba(_daab, _dbad)
	}
	_acbcd, _agee := _ggag(_fgag, _edbfa, _edcad)
	if _agee != nil {
		return _ba(_daab, _dbad)
	}
	_gdfc, _feeab := _ad.GetIntVal(_fgag.Get("\u0057\u004d\u006fd\u0065"))
	_adge, _dbbg := _acbcd.WMode()
	if _feeab && _dbbg {
		if _adge != _gdfc {
			return _ba(_daab, _dbad)
		}
	}
	if (_feeab && !_dbbg) || (!_feeab && _dbbg) {
		return _ba(_daab, _dbad)
	}
	return _bg
}
func _cdcc(_gfff *Profile2Options) {
	if _gfff.Now == nil {
		_gfff.Now = _a.Now
	}
}

type profile2 struct {
	_fcdd standardType
	_eceb Profile2Options
}

func (_cfd standardType) outputIntentSubtype() _da.PdfOutputIntentType {
	switch _cfd._efe {
	case 1:
		return _da.PdfOutputIntentTypeA1
	case 2:
		return _da.PdfOutputIntentTypeA2
	case 3:
		return _da.PdfOutputIntentTypeA3
	case 4:
		return _da.PdfOutputIntentTypeA4
	default:
		return 0
	}
}
func _cafe(_febe *_da.CompliancePdfReader) (_afdbeb ViolatedRule) {
	_adcd, _bedg := _acgab(_febe)
	if !_bedg {
		return _bg
	}
	if _adcd.Get("\u0041\u0041") != nil {
		return _ba("\u0036.\u0035\u002e\u0032\u002d\u0031", "\u0054h\u0065\u0020\u0064\u006fc\u0075m\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020a\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079 \u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079\u002e")
	}
	return _bg
}
func _gbdf(_gade *_da.CompliancePdfReader) ViolatedRule {
	_dfad, _egee := _gade.PdfReader.GetTrailer()
	if _egee != nil {
		return _ba("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u006d\u0069\u0073s\u0069\u006e\u0067\u0020t\u0072\u0061\u0069\u006c\u0065\u0072\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	}
	if _dfad.Get("\u0049\u0044") == nil {
		return _ba("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068e\u0020\u0027\u0049\u0044\u0027\u0020k\u0065\u0079\u0077o\u0072\u0064")
	}
	if _dfad.Get("\u0045n\u0063\u0072\u0079\u0070\u0074") != nil {
		return _ba("\u0036.\u0031\u002e\u0033\u002d\u0032", "\u0054\u0068\u0065\u0020\u006b\u0065y\u0077\u006f\u0072\u0064\u0020'\u0045\u006e\u0063\u0072\u0079\u0070t\u0027\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0075\u0073\u0065d\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u002e\u0020")
	}
	return _bg
}

// DefaultProfile2Options are the default options for the Profile2.
func DefaultProfile2Options() *Profile2Options {
	return &Profile2Options{Now: _a.Now, Xmp: XmpOptions{MarshalIndent: "\u0009"}}
}

type imageInfo struct {
	ColorSpace       _ad.PdfObjectName
	BitsPerComponent int
	ColorComponents  int
	Width            int
	Height           int
	Stream           *_ad.PdfObjectStream
	_dfb             bool
}
type imageModifications struct {
	_cdbd *colorspaceModification
	_fgg  _ad.StreamEncoder
}

func _bea(_bc []*_ea.Image, _dfcc bool) error {
	_fea := _ad.PdfObjectName("\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B")
	if _dfcc {
		_fea = "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b"
	}
	for _, _bcg := range _bc {
		if _bcg.Colorspace == _fea {
			continue
		}
		_eecf, _fbd := _da.NewXObjectImageFromStream(_bcg.Stream)
		if _fbd != nil {
			return _fbd
		}
		_bge, _fbd := _eecf.ToImage()
		if _fbd != nil {
			return _fbd
		}
		_ggf, _fbd := _bge.ToGoImage()
		if _fbd != nil {
			return _fbd
		}
		var _gac _da.PdfColorspace
		if _dfcc {
			_gac = _da.NewPdfColorspaceDeviceCMYK()
			_ggf, _fbd = _de.CMYKConverter.Convert(_ggf)
		} else {
			_gac = _da.NewPdfColorspaceDeviceRGB()
			_ggf, _fbd = _de.NRGBAConverter.Convert(_ggf)
		}
		if _fbd != nil {
			return _fbd
		}
		_cgga, _ffg := _ggf.(_de.Image)
		if !_ffg {
			return _gg.New("\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074 \u0069\u006d\u0070\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002eI\u006d\u0061\u0067\u0065")
		}
		_dff := _cgga.Base()
		_gcd := &_da.Image{Width: int64(_dff.Width), Height: int64(_dff.Height), BitsPerComponent: int64(_dff.BitsPerComponent), ColorComponents: _dff.ColorComponents, Data: _dff.Data}
		_gcd.SetDecode(_dff.Decode)
		_gcd.SetAlpha(_dff.Alpha)
		if _fbd = _eecf.SetImage(_gcd, _gac); _fbd != nil {
			return _fbd
		}
		_eecf.ToPdfObject()
		_bcg.ColorComponents = _dff.ColorComponents
		_bcg.Colorspace = _fea
	}
	return nil
}
func _gbcfg(_deeca *_da.CompliancePdfReader) (_deeg []ViolatedRule) {
	var _bdgf, _cbaf, _feadg, _dfgb, _gaeg, _bedfd bool
	_cceg := map[*_ad.PdfObjectStream]struct{}{}
	for _, _eae := range _deeca.GetObjectNums() {
		if _bdgf && _cbaf && _gaeg && _feadg && _dfgb && _bedfd {
			return _deeg
		}
		_eebe, _bgaa := _deeca.GetIndirectObjectByNumber(_eae)
		if _bgaa != nil {
			continue
		}
		_ggdc, _dbgdb := _ad.GetStream(_eebe)
		if !_dbgdb {
			continue
		}
		if _, _dbgdb = _cceg[_ggdc]; _dbgdb {
			continue
		}
		_cceg[_ggdc] = struct{}{}
		_bbee, _dbgdb := _ad.GetName(_ggdc.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_dbgdb {
			continue
		}
		if !_dfgb {
			if _ggdc.Get("\u0052\u0065\u0066") != nil {
				_deeg = append(_deeg, _ba("\u0036.\u0032\u002e\u0036\u002d\u0031", "\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0058O\u0062\u006a\u0065\u0063\u0074s\u002e"))
				_dfgb = true
			}
		}
		if _bbee.String() == "\u0050\u0053" {
			if !_bedfd {
				_deeg = append(_deeg, _ba("\u0036.\u0032\u002e\u0037\u002d\u0031", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073."))
				_bedfd = true
				continue
			}
		}
		if _bbee.String() == "\u0046\u006f\u0072\u006d" {
			if _cbaf && _feadg && _dfgb {
				continue
			}
			if !_cbaf && _ggdc.Get("\u004f\u0050\u0049") != nil {
				_deeg = append(_deeg, _ba("\u0036.\u0032\u002e\u0034\u002d\u0032", "\u0041\u006e\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u0028\u0049\u006d\u0061\u0067\u0065\u0020\u006f\u0072\u0020\u0046\u006f\u0072\u006d\u0029\u0020\u0073\u0068\u0061\u006cl\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u004fP\u0049\u0020\u006b\u0065\u0079\u002e"))
				_cbaf = true
			}
			if !_feadg {
				if _ggdc.Get("\u0050\u0053") != nil {
					_feadg = true
				}
				if _eccb := _ggdc.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _eccb != nil && !_feadg {
					if _faccg, _abb := _ad.GetName(_eccb); _abb && *_faccg == "\u0050\u0053" {
						_feadg = true
					}
				}
				if _feadg {
					_deeg = append(_deeg, _ba("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e"))
				}
			}
			continue
		}
		if _bbee.String() != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		if !_bdgf && _ggdc.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
			_deeg = append(_deeg, _ba("\u0036.\u0032\u002e\u0034\u002d\u0031", "\u0041\u006e\u0020\u0049m\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
			_bdgf = true
		}
		if !_gaeg && _ggdc.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065") != nil {
			_gdcf, _adae := _ad.GetBool(_ggdc.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _adae && bool(*_gdcf) {
				continue
			}
			_deeg = append(_deeg, _ba("\u0036.\u0032\u002e\u0034\u002d\u0033", "\u0049\u0066 a\u006e\u0020\u0049\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0063o\u006e\u0074\u0061\u0069n\u0073\u0020\u0074\u0068e \u0049\u006et\u0065r\u0070\u006f\u006c\u0061\u0074\u0065 \u006b\u0065\u0079,\u0020\u0069t\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020b\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u002e"))
			_gaeg = true
		}
	}
	return _deeg
}

var _ Profile = (*Profile1B)(nil)

func (_gbf *documentImages) hasOnlyDeviceGray() bool { return _gbf._eeb && !_gbf._eea && !_gbf._fc }
func _fec(_beg *_ea.Document) error {
	_ddca := map[string]*_ad.PdfObjectDictionary{}
	_fcb := _dad.NewFinder(&_dad.FinderOpts{Extensions: []string{"\u002e\u0074\u0074\u0066"}})
	_ccd := map[_ad.PdfObject]struct{}{}
	_bce := map[_ad.PdfObject]struct{}{}
	for _, _efa := range _beg.Objects {
		_eef, _fee := _ad.GetDict(_efa)
		if !_fee {
			continue
		}
		_ceca := _eef.Get("\u0054\u0079\u0070\u0065")
		if _ceca == nil {
			continue
		}
		if _fgf, _ffc := _ad.GetName(_ceca); _ffc && _fgf.String() != "\u0046\u006f\u006e\u0074" {
			continue
		}
		if _, _fbbg := _ccd[_efa]; _fbbg {
			continue
		}
		_geg, _eaae := _da.NewPdfFontFromPdfObject(_eef)
		if _eaae != nil {
			_ca.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006fn\u0074\u0020\u0066\u0072\u006fm\u0020\u006fb\u006a\u0065\u0063\u0074")
			return _eaae
		}
		_dab, _eaae := _geg.GetFontDescriptor()
		if _eaae != nil {
			return _eaae
		}
		if _dab != nil && (_dab.FontFile != nil || _dab.FontFile2 != nil || _dab.FontFile3 != nil) {
			continue
		}
		_gdb := _geg.BaseFont()
		if _gdb == "" {
			return _d.Errorf("\u006f\u006e\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0073\u0020\u0073\u0079\u006e\u0074\u0061\u0078\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069d\u0020\u002d\u0020\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u0075\u006ed\u0065\u0066\u0069n\u0065\u0064\u003a\u0020\u0025\u0073", _eef.String())
		}
		_eag, _gec := _ddca[_gdb]
		if !_gec {
			if len(_gdb) > 7 && _gdb[6] == '+' {
				_gdb = _gdb[7:]
			}
			_gggb := []string{_gdb, "\u0054i\u006de\u0073\u0020\u004e\u0065\u0077\u0020\u0052\u006f\u006d\u0061\u006e", "\u0041\u0072\u0069a\u006c", "D\u0065\u006a\u0061\u0056\u0075\u0020\u0053\u0061\u006e\u0073"}
			for _, _aeg := range _gggb {
				_ca.Log.Debug("\u0044\u0045\u0042\u0055\u0047\u003a \u0073\u0065\u0061\u0072\u0063\u0068\u0069\u006e\u0067\u0020\u0073\u0079\u0073t\u0065\u006d\u0020\u0066\u006f\u006e\u0074 \u0060\u0025\u0073\u0060", _aeg)
				if _eag, _gec = _ddca[_aeg]; _gec {
					break
				}
				_bfe := _fcb.Match(_aeg)
				if _bfe == nil {
					_ca.Log.Debug("c\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0066\u006fn\u0074\u0020\u0066i\u006ce\u0020\u0025\u0073", _aeg)
					continue
				}
				_fbf, _gdd := _da.NewPdfFontFromTTFFile(_bfe.Filename)
				if _gdd != nil {
					return _gdd
				}
				_dgf := _fbf.FontDescriptor()
				if _dgf.FontFile != nil {
					if _, _gec = _bce[_dgf.FontFile]; !_gec {
						_beg.Objects = append(_beg.Objects, _dgf.FontFile)
						_bce[_dgf.FontFile] = struct{}{}
					}
				}
				if _dgf.FontFile2 != nil {
					if _, _gec = _bce[_dgf.FontFile2]; !_gec {
						_beg.Objects = append(_beg.Objects, _dgf.FontFile2)
						_bce[_dgf.FontFile2] = struct{}{}
					}
				}
				if _dgf.FontFile3 != nil {
					if _, _gec = _bce[_dgf.FontFile3]; !_gec {
						_beg.Objects = append(_beg.Objects, _dgf.FontFile3)
						_bce[_dgf.FontFile3] = struct{}{}
					}
				}
				_ecc, _bcd := _fbf.ToPdfObject().(*_ad.PdfIndirectObject)
				if !_bcd {
					_ca.Log.Debug("\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006e\u006ft\u0020\u0061\u006e\u0020\u0069\u006e\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074")
					continue
				}
				_daa, _bcd := _ecc.PdfObject.(*_ad.PdfObjectDictionary)
				if !_bcd {
					_ca.Log.Debug("\u0046\u006fn\u0074\u0020\u0074\u0079p\u0065\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u006e \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
					continue
				}
				_ddca[_aeg] = _daa
				_eag = _daa
				break
			}
			if _eag == nil {
				_ca.Log.Debug("\u004e\u006f\u0020\u006d\u0061\u0074\u0063\u0068\u0069\u006eg\u0020\u0066\u006f\u006e\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0066\u006f\u0072\u003a\u0020\u0025\u0073", _geg.BaseFont())
				return _gg.New("\u006e\u006f m\u0061\u0074\u0063h\u0069\u006e\u0067\u0020fon\u0074 f\u006f\u0075\u006e\u0064\u0020\u0069\u006e t\u0068\u0065\u0020\u0073\u0079\u0073\u0074e\u006d")
			}
		}
		for _, _edb := range _eag.Keys() {
			_eef.Set(_edb, _eag.Get(_edb))
		}
		_fac := _eag.Get("\u0057\u0069\u0064\u0074\u0068\u0073")
		if _fac != nil {
			if _, _gec = _bce[_fac]; !_gec {
				_beg.Objects = append(_beg.Objects, _fac)
				_bce[_fac] = struct{}{}
			}
		}
		_ccd[_efa] = struct{}{}
		_bcdd := _eef.Get("\u0046\u006f\u006e\u0074\u0044\u0065\u0073\u0063\u0072i\u0070\u0074\u006f\u0072")
		if _bcdd != nil {
			_beg.Objects = append(_beg.Objects, _bcdd)
			_bce[_bcdd] = struct{}{}
		}
	}
	return nil
}

// Part gets the PDF/A version level.
func (_fffb *profile2) Part() int { return _fffb._fcdd._efe }
func _ffdf(_aedf *_ea.Document) error {
	_deeb, _ffa := _aedf.FindCatalog()
	if !_ffa {
		return _gg.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _deeb.Object.Get("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073") != nil {
		_deeb.Object.Remove("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073")
	}
	return nil
}
func _gbfd(_fdbc *_da.CompliancePdfReader) ViolatedRule {
	for _, _cfag := range _fdbc.PageList {
		_dacdf := _cfag.GetContentStreamObjs()
		for _, _gdae := range _dacdf {
			_gdae = _ad.TraceToDirectObject(_gdae)
			var _daed string
			switch _dgcc := _gdae.(type) {
			case *_ad.PdfObjectString:
				_daed = _dgcc.Str()
			case *_ad.PdfObjectStream:
				_fgeac, _adbc := _ad.GetName(_ad.TraceToDirectObject(_dgcc.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
				if _adbc {
					if *_fgeac == _ad.StreamEncodingFilterNameLZW {
						return _ba("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0032", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e")
					}
				}
				_eaca, _cdff := _ad.DecodeStream(_dgcc)
				if _cdff != nil {
					_ca.Log.Debug("\u0045r\u0072\u003a\u0020\u0025\u0076", _cdff)
					continue
				}
				_daed = string(_eaca)
			default:
				_ca.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063t\u003a\u0020\u0025\u0054", _gdae)
				continue
			}
			_dfed := _cad.NewContentStreamParser(_daed)
			_ebad, _ccbb := _dfed.Parse()
			if _ccbb != nil {
				_ca.Log.Debug("\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d:\u0020\u0025\u0076", _ccbb)
				continue
			}
			for _, _ceec := range *_ebad {
				if !(_ceec.Operand == "\u0042\u0049" && len(_ceec.Params) == 1) {
					continue
				}
				_faaag, _ceed := _ceec.Params[0].(*_cad.ContentStreamInlineImage)
				if !_ceed {
					continue
				}
				_bgef, _bgc := _faaag.GetEncoder()
				if _bgc != nil {
					_ca.Log.Debug("\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006e\u006c\u0069\u006ee\u0020\u0069\u006d\u0061\u0067\u0065 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _bgc)
					continue
				}
				if _bgef.GetFilterName() == _ad.StreamEncodingFilterNameLZW {
					return _ba("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0032", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e")
				}
			}
		}
	}
	return _bg
}

var _ Profile = (*Profile2U)(nil)

func (_db standardType) String() string {
	return _d.Sprintf("\u0050\u0044\u0046\u002f\u0041\u002d\u0025\u0064\u0025\u0073", _db._efe, _db._ee)
}

// Profile2B is the implementation of the PDF/A-2B standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2B struct{ profile2 }

var _ Profile = (*Profile1A)(nil)

func _acgab(_gdgeff *_da.CompliancePdfReader) (*_ad.PdfObjectDictionary, bool) {
	_bedc, _bfddf := _gdgeff.GetTrailer()
	if _bfddf != nil {
		_ca.Log.Debug("\u0043\u0061\u006en\u006f\u0074\u0020\u0067e\u0074\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u003a\u0020\u0025\u0076", _bfddf)
		return nil, false
	}
	_dbdf, _baca := _bedc.Get("\u0052\u006f\u006f\u0074").(*_ad.PdfObjectReference)
	if !_baca {
		_ca.Log.Debug("\u0043a\u006e\u006e\u006f\u0074 \u0066\u0069\u006e\u0064\u0020d\u006fc\u0075m\u0065\u006e\u0074\u0020\u0072\u006f\u006ft")
		return nil, false
	}
	_fgad, _baca := _ad.GetDict(_ad.ResolveReference(_dbdf))
	if !_baca {
		_ca.Log.Debug("\u0063\u0061\u006e\u006e\u006f\u0074 \u0072\u0065\u0073\u006f\u006c\u0076\u0065\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079")
		return nil, false
	}
	return _fgad, true
}

// Error implements error interface.
func (_ec VerificationError) Error() string {
	_gb := _eg.Builder{}
	_gb.WriteString("\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u003a\u0020")
	_gb.WriteString(_d.Sprintf("\u0050\u0044\u0046\u002f\u0041\u002d\u0025\u0064\u0025\u0073", _ec.ConformanceLevel, _ec.ConformanceVariant))
	_gb.WriteString("\u0020\u0056\u0069\u006f\u006c\u0061\u0074\u0065\u0064\u0020\u0072\u0075l\u0065\u0073\u003a\u0020")
	for _aa, _af := range _ec.ViolatedRules {
		_gb.WriteString(_af.String())
		if _aa != len(_ec.ViolatedRules)-1 {
			_gb.WriteRune('\n')
		}
	}
	return _gb.String()
}
func _eecad(_ccea *_ad.PdfObjectDictionary, _eedg map[*_ad.PdfObjectStream][]byte, _fcgbb map[*_ad.PdfObjectStream]*_b.CMap) ViolatedRule {
	const (
		_egfd  = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0037\u002d\u0031"
		_ccagf = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006cl\u0020\u0069\u006e\u0063l\u0075\u0064e\u0020\u0061 \u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020w\u0068\u006f\u0073\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073 \u0061\u0020\u0043M\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u006d\u0061p\u0073\u0020\u0063\u0068\u0061\u0072ac\u0074\u0065\u0072\u0020\u0063\u006fd\u0065s\u0020\u0074\u006f\u0020\u0055\u006e\u0069\u0063\u006f\u0064e \u0076a\u006c\u0075\u0065\u0073,\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063r\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020P\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035.\u0039\u002c\u0020\u0075\u006e\u006ce\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006d\u0065\u0065\u0074\u0073 \u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0063\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u003a\u000a\u0020\u002d\u0020\u0066o\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0073\u0020M\u0061\u0063\u0052o\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074E\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041n\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u006f\u0072\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020t\u0068\u0065\u0020\u0070\u0072\u0065d\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0056\u0020C\u004d\u0061\u0070s\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0073\u0020a\u0072\u0065 \u0074\u0061k\u0065\u006e\u0020\u0066\u0072\u006f\u006d\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u0020\u0073\u0074\u0061n\u0064\u0061\u0072\u0064\u0020L\u0061t\u0069\u006e\u0020\u0063\u0068a\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0074\u0020\u006fr\u0020\u0074\u0068\u0065 \u0073\u0065\u0074\u0020\u006f\u0066 \u006e\u0061\u006d\u0065\u0064\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0066\u006f\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046 \u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0041\u0070\u0070\u0065\u006e\u0064\u0069\u0078 \u0044\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020w\u0068\u006f\u0073e\u0020d\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u0047B\u0031\u002c\u0020\u0041\u0064\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002c\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031\u0020\u006f\u0072\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006fr\u0065\u0061\u0031\u0020\u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u006c\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0073\u002e"
	)
	_bgbge, _fcbfd := _ad.GetStream(_ccea.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e"))
	if _fcbfd {
		_, _cedf := _ggag(_bgbge, _eedg, _fcgbb)
		if _cedf != nil {
			return _ba(_egfd, _ccagf)
		}
		return _bg
	}
	_acbe, _fcbfd := _ad.GetName(_ccea.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_fcbfd {
		return _ba(_egfd, _ccagf)
	}
	switch _acbe.String() {
	case "\u0054\u0079\u0070e\u0031":
		return _bg
	}
	return _ba(_egfd, _ccagf)
}
func _fdef(_ecdff *_da.CompliancePdfReader) ViolatedRule { return _bg }
func _eec() standardType                                 { return standardType{_efe: 2, _ee: "\u0055"} }
func _dgag(_cabc *_ad.PdfObjectDictionary) ViolatedRule {
	const (
		_ccad = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0032"
		_gbde = "IS\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u00209\u002e\u0037\u002e\u0034\u002c\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0031\u0031\u0037\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068a\u0074\u0020\u0061\u006c\u006c\u0020\u0065m\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0054\u0079\u0070\u0065\u0020\u0032\u0020\u0043\u0049\u0044\u0046\u006fn\u0074\u0073\u0020\u0069n\u0020t\u0068e\u0020\u0043\u0049D\u0046\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0043\u0049\u0044\u0054\u006fG\u0049\u0044M\u0061\u0070\u0020\u0065\u006e\u0074\u0072\u0079 \u0074\u0068\u0061\u0074\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020\u0061\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u006d\u0061\u0070p\u0069\u006e\u0067 f\u0072\u006f\u006d \u0043\u0049\u0044\u0073\u0020\u0074\u006f\u0020\u0067\u006c\u0079p\u0068 \u0069\u006e\u0064\u0069c\u0065\u0073\u0020\u006fr\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0020\u0049d\u0065\u006e\u0074\u0069\u0074\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0049\u0053\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a\u0032\u0030\u0030\u0038\u002c\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0031\u0031\u0037\u002e"
	)
	var _bgag string
	if _fdee, _geae := _ad.GetName(_cabc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _geae {
		_bgag = _fdee.String()
	}
	if _bgag != "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032" {
		return _bg
	}
	if _cabc.Get("C\u0049\u0044\u0054\u006f\u0047\u0049\u0044\u004d\u0061\u0070") == nil {
		return _ba(_ccad, _gbde)
	}
	return _bg
}
func _faabd(_bbe *_da.CompliancePdfReader, _gadd bool) (_geeg []ViolatedRule) {
	var _fffee, _aafeg, _ceeg, _fcgb, _cbba, _egcb, _ffgf bool
	_cfef := func() bool { return _fffee && _aafeg && _ceeg && _fcgb && _cbba && _egcb && _ffgf }
	_gfgd, _ggfc := _faad(_bbe)
	var _fdbcf _ac.ProfileHeader
	if _ggfc {
		_fdbcf, _ = _ac.ParseHeader(_gfgd.DestOutputProfile)
	}
	var _afcd bool
	_beadc := map[_ad.PdfObject]struct{}{}
	var _bdcb func(_ffgc _da.PdfColorspace) bool
	_bdcb = func(_fdea _da.PdfColorspace) bool {
		switch _ecdfd := _fdea.(type) {
		case *_da.PdfColorspaceDeviceGray:
			if !_egcb {
				if !_ggfc {
					_afcd = true
					_geeg = append(_geeg, _ba("\u0036.\u0032\u002e\u0033\u002d\u0034", "\u0044\u0065\u0076\u0069\u0063\u0065G\u0072\u0061\u0079\u0020\u006da\u0079\u0020\u0062\u0065\u0020\u0075s\u0065\u0064\u0020\u006f\u006el\u0079\u0020\u0069\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006ce\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020O\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u002e"))
					_egcb = true
					if _cfef() {
						return true
					}
				}
			}
		case *_da.PdfColorspaceDeviceRGB:
			if !_fcgb {
				if !_ggfc || _fdbcf.ColorSpace != _ac.ColorSpaceRGB {
					_afcd = true
					_geeg = append(_geeg, _ba("\u0036.\u0032\u002e\u0033\u002d\u0032", "\u0044\u0065\u0076\u0069\u0063\u0065\u0052\u0047\u0042\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u006e\u006c\u0079\u0020\u0069\u0066\u0020\u0074\u0068\u0065 \u0066\u0069\u006c\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074In\u0074\u0065\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020u\u0073es\u0020a\u006e\u0020\u0052\u0047\u0042\u0020\u0063o\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u002e"))
					_fcgb = true
					if _cfef() {
						return true
					}
				}
			}
		case *_da.PdfColorspaceDeviceCMYK:
			if !_cbba {
				if !_ggfc || _fdbcf.ColorSpace != _ac.ColorSpaceCMYK {
					_afcd = true
					_geeg = append(_geeg, _ba("\u0036.\u0032\u002e\u0033\u002d\u0033", "\u0044\u0065\u0076\u0069\u0063e\u0043\u004d\u0059\u004b \u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u006e\u006c\u0079\u0020\u0069\u0066\u0020\u0074h\u0065\u0020\u0066\u0069\u006ce \u0068\u0061\u0073\u0020\u0061 \u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0068a\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0061\u006e \u0043\u004d\u0059\u004b\u0020\u0063\u006f\u006c\u006f\u0072\u0020s\u0070\u0061\u0063e\u002e"))
					_cbba = true
					if _cfef() {
						return true
					}
				}
			}
		case *_da.PdfColorspaceICCBased:
			if !_ceeg || !_ffgf {
				_ccaa, _ebba := _ac.ParseHeader(_ecdfd.Data)
				if _ebba != nil {
					_ca.Log.Debug("\u0070\u0061\u0072si\u006e\u0067\u0020\u0049\u0043\u0043\u0042\u0061\u0073e\u0064 \u0068e\u0061d\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ebba)
					_geeg = append(_geeg, func() ViolatedRule {
						return _ba("\u0036.\u0032\u002e\u0033\u002d\u0031", "\u0041\u006cl \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006co\u0072\u0020\u0073\u0070a\u0063e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0061\u0073\u0020\u0049\u0043\u0043 \u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0073 \u0061\u0073\u0020d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0034\u002e\u0035")
					}())
					_ceeg = true
					if _cfef() {
						return true
					}
				}
				if !_ceeg {
					var _cfbg, _gdfbg bool
					switch _ccaa.DeviceClass {
					case _ac.DeviceClassPRTR, _ac.DeviceClassMNTR, _ac.DeviceClassSCNR, _ac.DeviceClassSPAC:
					default:
						_cfbg = true
					}
					switch _ccaa.ColorSpace {
					case _ac.ColorSpaceRGB, _ac.ColorSpaceCMYK, _ac.ColorSpaceGRAY, _ac.ColorSpaceLAB:
					default:
						_gdfbg = true
					}
					if _cfbg || _gdfbg {
						_geeg = append(_geeg, _ba("\u0036.\u0032\u002e\u0033\u002d\u0031", "\u0041\u006cl \u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006f\u006co\u0072\u0020\u0073\u0070a\u0063e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0061\u0073\u0020\u0049\u0043\u0043 \u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0073 \u0061\u0073\u0020d\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0034\u002e\u0035"))
						_ceeg = true
						if _cfef() {
							return true
						}
					}
				}
				if !_ffgf {
					_geab, _ := _ad.GetStream(_ecdfd.GetContainingPdfObject())
					if _geab.Get("\u004e") == nil || (_ecdfd.N == 1 && _ccaa.ColorSpace != _ac.ColorSpaceGRAY) || (_ecdfd.N == 3 && !(_ccaa.ColorSpace == _ac.ColorSpaceRGB || _ccaa.ColorSpace == _ac.ColorSpaceLAB)) || (_ecdfd.N == 4 && _ccaa.ColorSpace != _ac.ColorSpaceCMYK) {
						_geeg = append(_geeg, _ba("\u0036.\u0032\u002e\u0033\u002d\u0035", "\u0049\u0066\u0020a\u006e\u0020u\u006e\u0063\u0061\u006c\u0069\u0062\u0072a\u0074\u0065\u0064\u0020\u0063\u006fl\u006f\u0072 \u0073\u0070\u0061c\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u0069\u006c\u0065 \u0074\u0068\u0065\u006e \u0074\u0068\u0061\u0074 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041-\u0031\u0020\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065d\u0020\u0069\u006e\u0020\u0036\u002e\u0032\u002e\u0032\u002e"))
						_ffgf = true
						if _cfef() {
							return true
						}
					}
				}
			}
			if _ecdfd.Alternate != nil {
				return _bdcb(_ecdfd.Alternate)
			}
		}
		return false
	}
	for _, _efceg := range _bbe.GetObjectNums() {
		_abdbe, _geff := _bbe.GetIndirectObjectByNumber(_efceg)
		if _geff != nil {
			continue
		}
		_fccc, _bcdc := _ad.GetStream(_abdbe)
		if !_bcdc {
			continue
		}
		_ggdb, _bcdc := _ad.GetName(_fccc.Get("\u0054\u0079\u0070\u0065"))
		if !_bcdc || _ggdb.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_cdgf, _bcdc := _ad.GetName(_fccc.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_bcdc {
			continue
		}
		_beadc[_fccc] = struct{}{}
		switch _cdgf.String() {
		case "\u0049\u006d\u0061g\u0065":
			_cfdaf, _ccf := _da.NewXObjectImageFromStream(_fccc)
			if _ccf != nil {
				continue
			}
			_beadc[_fccc] = struct{}{}
			if _bdcb(_cfdaf.ColorSpace) {
				return _geeg
			}
		case "\u0046\u006f\u0072\u006d":
			_gcbf, _aef := _ad.GetDict(_fccc.Get("\u0047\u0072\u006fu\u0070"))
			if !_aef {
				continue
			}
			_abgb := _gcbf.Get("\u0043\u0053")
			if _abgb == nil {
				continue
			}
			_bdd, _abaa := _da.NewPdfColorspaceFromPdfObject(_abgb)
			if _abaa != nil {
				continue
			}
			if _bdcb(_bdd) {
				return _geeg
			}
		}
	}
	for _, _cbcee := range _bbe.PageList {
		_cdda, _baad := _cbcee.GetContentStreams()
		if _baad != nil {
			continue
		}
		for _, _fbge := range _cdda {
			_aab, _aggf := _cad.NewContentStreamParser(_fbge).Parse()
			if _aggf != nil {
				continue
			}
			for _, _aaba := range *_aab {
				if len(_aaba.Params) > 1 {
					continue
				}
				switch _aaba.Operand {
				case "\u0042\u0049":
					_eaaf, _edgc := _aaba.Params[0].(*_cad.ContentStreamInlineImage)
					if !_edgc {
						continue
					}
					_cfcae, _eccae := _eaaf.GetColorSpace(_cbcee.Resources)
					if _eccae != nil {
						continue
					}
					if _bdcb(_cfcae) {
						return _geeg
					}
				case "\u0044\u006f":
					_bged, _fdgc := _ad.GetName(_aaba.Params[0])
					if !_fdgc {
						continue
					}
					_geegg, _beebc := _cbcee.Resources.GetXObjectByName(*_bged)
					if _, _feef := _beadc[_geegg]; _feef {
						continue
					}
					switch _beebc {
					case _da.XObjectTypeImage:
						_eegg, _affb := _da.NewXObjectImageFromStream(_geegg)
						if _affb != nil {
							continue
						}
						_beadc[_geegg] = struct{}{}
						if _bdcb(_eegg.ColorSpace) {
							return _geeg
						}
					case _da.XObjectTypeForm:
						_cgfg, _cbec := _ad.GetDict(_geegg.Get("\u0047\u0072\u006fu\u0070"))
						if !_cbec {
							continue
						}
						_cccd, _cbec := _ad.GetName(_cgfg.Get("\u0043\u0053"))
						if !_cbec {
							continue
						}
						_ccac, _cgec := _da.NewPdfColorspaceFromPdfObject(_cccd)
						if _cgec != nil {
							continue
						}
						_beadc[_geegg] = struct{}{}
						if _bdcb(_ccac) {
							return _geeg
						}
					}
				}
			}
		}
	}
	if !_afcd {
		return _geeg
	}
	if (_fdbcf.DeviceClass == _ac.DeviceClassPRTR || _fdbcf.DeviceClass == _ac.DeviceClassMNTR) && (_fdbcf.ColorSpace == _ac.ColorSpaceRGB || _fdbcf.ColorSpace == _ac.ColorSpaceCMYK || _fdbcf.ColorSpace == _ac.ColorSpaceGRAY) {
		return _geeg
	}
	if !_gadd {
		return _geeg
	}
	_ggffe, _fbgg := _acgab(_bbe)
	if !_fbgg {
		return _geeg
	}
	_ebaf, _fbgg := _ad.GetArray(_ggffe.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_fbgg {
		_geeg = append(_geeg, _ba("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0031\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074e\u006e\u0074\u0020\u0069\u0073\u0020a\u006e \u004f\u0075\u0074\u0070\u0075\u0074\u0049n\u0074\u0065\u006e\u0074\u0020\u0064i\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0062y\u0020\u0050\u0044F\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0031\u0030.4\u002c\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073 \u0069\u006e\u0063\u006c\u0075\u0064e\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0027\u0073\u0020O\u0075\u0074p\u0075\u0074I\u006e\u0074\u0065\u006e\u0074\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020a\u006e\u0064\u0020h\u0061\u0073\u0020\u0047\u0054\u0053\u005f\u0050\u0044\u0046\u0041\u0031\u0020\u0061\u0073 \u0074\u0068\u0065\u0020\u0076a\u006c\u0075e\u0020\u006f\u0066\u0020i\u0074\u0073 \u0053\u0020\u006b\u0065\u0079\u0020\u0061\u006e\u0064\u0020\u0061\u0020\u0076\u0061\u006c\u0069\u0064\u0020I\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006ce\u0020s\u0074\u0072\u0065\u0061\u006d \u0061\u0073\u0020\u0074h\u0065\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0074\u0073\u0020\u0044\u0065\u0073t\u004f\u0075t\u0070\u0075\u0074P\u0072\u006f\u0066\u0069\u006c\u0065 \u006b\u0065\u0079\u002e"), _ba("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e"))
		return _geeg
	}
	if _ebaf.Len() > 1 {
		_ddfbc := map[*_ad.PdfObjectDictionary]struct{}{}
		for _ddb := 0; _ddb < _ebaf.Len(); _ddb++ {
			_ddgf, _aeb := _ad.GetDict(_ebaf.Get(_ddb))
			if !_aeb {
				continue
			}
			if _ddb == 0 {
				_ddfbc[_ddgf] = struct{}{}
				continue
			}
			if _, _cfec := _ddfbc[_ddgf]; !_cfec {
				_geeg = append(_geeg, _ba("\u0036.\u0032\u002e\u0032\u002d\u0032", "\u0049\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065's\u0020O\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073 \u0061\u0072\u0072a\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u006f\u006ee\u0020\u0065\u006e\u0074\u0072\u0079\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0065n\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e a \u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006cl\u0020\u0068\u0061\u0076\u0065 \u0061\u0073\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068a\u0074\u0020\u006b\u0065\u0079 \u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0069\u006e\u0064\u0069\u0072\u0065c\u0074\u0020\u006fb\u006ae\u0063t\u002c\u0020\u0077h\u0069\u0063\u0068\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u0061\u0020\u0076\u0061\u006c\u0069d\u0020\u0049\u0043\u0043\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0073\u0074r\u0065\u0061m\u002e"))
				break
			}
		}
	}
	return _geeg
}

// Part gets the PDF/A version level.
func (_bbbc *profile1) Part() int { return _bbbc._efde._efe }

// ApplyStandard tries to change the content of the writer to match the PDF/A-1 standard.
// Implements model.StandardApplier.
func (_cfcfe *profile1) ApplyStandard(document *_ea.Document) (_ccc error) {
	_fdae(document, 4)
	if _ccc = _cdae(document, _cfcfe._fcce.Now); _ccc != nil {
		return _ccc
	}
	if _ccc = _agc(document); _ccc != nil {
		return _ccc
	}
	_bcdb, _acea := _fdg(_cfcfe._fcce.CMYKDefaultColorSpace, _cfcfe._efde)
	_ccc = _fae(document, []pageColorspaceOptimizeFunc{_gdga, _bcdb}, []documentColorspaceOptimizeFunc{_acea})
	if _ccc != nil {
		return _ccc
	}
	_dbb(document)
	if _ccc = _bae(document, _cfcfe._efde._efe); _ccc != nil {
		return _ccc
	}
	if _ccc = _dccb(document); _ccc != nil {
		return _ccc
	}
	if _ccc = _dccbc(document); _ccc != nil {
		return _ccc
	}
	if _ccc = _acb(document); _ccc != nil {
		return _ccc
	}
	if _ccc = _fec(document); _ccc != nil {
		return _ccc
	}
	if _cfcfe._efde._ee == "\u0041" {
		_ecgd(document)
	}
	if _ccc = _geb(document, _cfcfe._efde._efe); _ccc != nil {
		return _ccc
	}
	if _ccc = _dbba(document); _ccc != nil {
		return _ccc
	}
	if _bgde := _aca(document, _cfcfe._efde, _cfcfe._fcce.Xmp); _bgde != nil {
		return _bgde
	}
	if _cfcfe._efde == _df() {
		if _ccc = _bgd(document); _ccc != nil {
			return _ccc
		}
	}
	if _ccc = _dcc(document); _ccc != nil {
		return _ccc
	}
	return nil
}
func _abdb(_bdfg *_da.CompliancePdfReader) ViolatedRule { return _bg }

// ApplyStandard tries to change the content of the writer to match the PDF/A-2 standard.
// Implements model.StandardApplier.
func (_fbfce *profile2) ApplyStandard(document *_ea.Document) (_acgae error) {
	_fdae(document, 7)
	if _acgae = _cdae(document, _fbfce._eceb.Now); _acgae != nil {
		return _acgae
	}
	if _acgae = _agc(document); _acgae != nil {
		return _acgae
	}
	_caba, _bfec := _fdg(_fbfce._eceb.CMYKDefaultColorSpace, _fbfce._fcdd)
	_acgae = _fae(document, []pageColorspaceOptimizeFunc{_caba}, []documentColorspaceOptimizeFunc{_bfec})
	if _acgae != nil {
		return _acgae
	}
	_dbb(document)
	if _acgae = _agge(document); _acgae != nil {
		return _acgae
	}
	if _acgae = _bae(document, _fbfce._fcdd._efe); _acgae != nil {
		return _acgae
	}
	if _acgae = _facca(document); _acgae != nil {
		return _acgae
	}
	if _acgae = _baafg(document); _acgae != nil {
		return _acgae
	}
	if _acgae = _fec(document); _acgae != nil {
		return _acgae
	}
	if _acgae = _faaa(document); _acgae != nil {
		return _acgae
	}
	if _fbfce._fcdd._ee == "\u0041" {
		_ecgd(document)
	}
	if _acgae = _geb(document, _fbfce._fcdd._efe); _acgae != nil {
		return _acgae
	}
	if _acgae = _dbba(document); _acgae != nil {
		return _acgae
	}
	if _dgc := _aca(document, _fbfce._fcdd, _fbfce._eceb.Xmp); _dgc != nil {
		return _dgc
	}
	if _fbfce._fcdd == _bb() {
		if _acgae = _bgd(document); _acgae != nil {
			return _acgae
		}
	}
	if _acgae = _dcaf(document); _acgae != nil {
		return _acgae
	}
	if _acgae = _gcdca(document); _acgae != nil {
		return _acgae
	}
	if _acgae = _ffdf(document); _acgae != nil {
		return _acgae
	}
	return nil
}
func _dbba(_baaf *_ea.Document) error {
	_agg, _efbc := _baaf.FindCatalog()
	if !_efbc {
		return _gg.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_, _efbc = _ad.GetDict(_agg.Object.Get("\u0041\u0041"))
	if !_efbc {
		return nil
	}
	_agg.Object.Remove("\u0041\u0041")
	return nil
}
func _geb(_egf *_ea.Document, _ebd int) error {
	for _, _cae := range _egf.Objects {
		_bbc, _dce := _ad.GetDict(_cae)
		if !_dce {
			continue
		}
		_bcfd := _bbc.Get("\u0054\u0079\u0070\u0065")
		if _bcfd == nil {
			continue
		}
		if _bfa, _gea := _ad.GetName(_bcfd); _gea && _bfa.String() != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_baa, _cbf := _ad.GetName(_bbc.Get("\u0053"))
		if !_cbf {
			continue
		}
		switch _da.PdfActionType(*_baa) {
		case _da.ActionTypeLaunch, _da.ActionTypeSound, _da.ActionTypeMovie, _da.ActionTypeResetForm, _da.ActionTypeImportData, _da.ActionTypeJavaScript:
			_bbc.Remove("\u0053")
		case _da.ActionTypeHide, _da.ActionTypeSetOCGState, _da.ActionTypeRendition, _da.ActionTypeTrans, _da.ActionTypeGoTo3DView:
			if _ebd == 2 {
				_bbc.Remove("\u0053")
			}
		case _da.ActionTypeNamed:
			_ccdf, _aafc := _ad.GetName(_bbc.Get("\u004e"))
			if !_aafc {
				continue
			}
			switch *_ccdf {
			case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
			default:
				_bbc.Remove("\u004e")
			}
		}
	}
	return nil
}

// ViolatedRule is the structure that defines violated PDF/A rule.
type ViolatedRule struct {
	RuleNo string
	Detail string
}

func _eege(_gabfd *_da.CompliancePdfReader) (_bdbe []ViolatedRule) {
	_adc := _gabfd.GetObjectNums()
	for _, _fbac := range _adc {
		_fgdc, _dabe := _gabfd.GetIndirectObjectByNumber(_fbac)
		if _dabe != nil {
			continue
		}
		_cgaf, _acbg := _ad.GetDict(_fgdc)
		if !_acbg {
			continue
		}
		_cfgf, _acbg := _ad.GetName(_cgaf.Get("\u0054\u0079\u0070\u0065"))
		if !_acbg {
			continue
		}
		if _cfgf.String() != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
			continue
		}
		if _cgaf.Get("\u0045\u0046") != nil {
			_bdbe = append(_bdbe, _ba("\u0036\u002e\u0031\u002e\u0031\u0031\u002d\u0031", "\u0041 \u0066\u0069\u006c\u0065 \u0073p\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006fn\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066i\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046 \u0033\u002e\u0031\u0030\u002e\u0032\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0045\u0046 \u006be\u0079\u002e"))
			break
		}
	}
	_edac, _abaf := _acgab(_gabfd)
	if !_abaf {
		return _bdbe
	}
	_fdcb, _abaf := _ad.GetDict(_edac.Get("\u004e\u0061\u006de\u0073"))
	if !_abaf {
		return _bdbe
	}
	if _fdcb.Get("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0046\u0069\u006c\u0065\u0073") != nil {
		_bdbe = append(_bdbe, _ba("\u0036\u002e\u0031\u002e\u0031\u0031\u002d\u0032", "\u0041\u0020\u0066i\u006c\u0065\u0027\u0073\u0020\u006e\u0061\u006d\u0065\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020d\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020PD\u0046 \u0052\u0065\u0066er\u0065\u006e\u0063\u0065\u0020\u0033\u002e6\u002e\u0033\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0045m\u0062\u0065\u0064\u0064\u0065\u0064\u0046i\u006c\u0065\u0073\u0020\u006b\u0065\u0079\u002e"))
	}
	return _bdbe
}
func _adde(_dfgg *_da.CompliancePdfReader) (_cfed ViolatedRule) {
	for _, _fega := range _dfgg.GetObjectNums() {
		_cage, _gfbeb := _dfgg.GetIndirectObjectByNumber(_fega)
		if _gfbeb != nil {
			continue
		}
		_cagf, _cgcc := _ad.GetStream(_cage)
		if !_cgcc {
			continue
		}
		_ggbf, _cgcc := _ad.GetName(_cagf.Get("\u0054\u0079\u0070\u0065"))
		if !_cgcc {
			continue
		}
		if *_ggbf != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_afdbe, _cgcc := _ad.GetName(_cagf.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_cgcc {
			continue
		}
		if *_afdbe == "\u0050\u0053" {
			return _ba("\u0036.\u0032\u002e\u0037\u002d\u0031", "A \u0063\u006fn\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066i\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0050\u006f\u0073t\u0053c\u0072\u0069\u0070\u0074\u0020\u0058\u004f\u0062j\u0065c\u0074\u0073.")
		}
	}
	return _cfed
}
func _dccbc(_fdfa *_ea.Document) error {
	_gcge, _edddg := _fdfa.GetPages()
	if !_edddg {
		return nil
	}
	for _, _cbae := range _gcge {
		_caaf, _aged := _ad.GetArray(_cbae.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_aged {
			continue
		}
		for _, _dfbg := range _caaf.Elements() {
			_dfbg = _ad.ResolveReference(_dfbg)
			if _, _gefd := _dfbg.(*_ad.PdfObjectNull); _gefd {
				continue
			}
			_gecg, _deecc := _ad.GetDict(_dfbg)
			if !_deecc {
				continue
			}
			_fdaeg, _ := _ad.GetIntVal(_gecg.Get("\u0046"))
			_fdaeg &= ^(1 << 0)
			_fdaeg &= ^(1 << 1)
			_fdaeg &= ^(1 << 5)
			_fdaeg |= 1 << 2
			_gecg.Set("\u0046", _ad.MakeInteger(int64(_fdaeg)))
			_ecdf := false
			if _gdbf := _gecg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"); _gdbf != nil {
				_fbfa, _acegd := _ad.GetName(_gdbf)
				if _acegd && _fbfa.String() == "\u0057\u0069\u0064\u0067\u0065\u0074" {
					_ecdf = true
					if _gecg.Get("\u0041\u0041") != nil {
						_gecg.Remove("\u0041\u0041")
					}
				}
			}
			if _gecg.Get("\u0043") != nil || _gecg.Get("\u0049\u0043") != nil {
				_fffe, _ega := _gff(_fdfa)
				if !_ega {
					_gecg.Remove("\u0043")
					_gecg.Remove("\u0049\u0043")
				} else {
					_ebaa, _fcc := _ad.GetIntVal(_fffe.Get("\u004e"))
					if !_fcc || _ebaa != 3 {
						_gecg.Remove("\u0043")
						_gecg.Remove("\u0049\u0043")
					}
				}
			}
			_gfbf, _deecc := _ad.GetDict(_gecg.Get("\u0041\u0050"))
			if _deecc {
				_cddg := _gfbf.Get("\u004e")
				if _cddg == nil {
					continue
				}
				if len(_gfbf.Keys()) > 1 {
					_gfbf.Clear()
					_gfbf.Set("\u004e", _cddg)
				}
				if _ecdf {
					_ffea, _baf := _ad.GetName(_gecg.Get("\u0046\u0054"))
					if _baf && *_ffea == "\u0042\u0074\u006e" {
						continue
					}
				}
			}
		}
	}
	return nil
}
func _bcdf(_afaf *_da.CompliancePdfReader) ViolatedRule { return _bg }

// Profile2A is the implementation of the PDF/A-2A standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2A struct{ profile2 }

func _efgg(_aadf *_da.CompliancePdfReader) ViolatedRule {
	_bgec := _aadf.ParserMetadata().HeaderCommentBytes()
	if _bgec[0] > 127 && _bgec[1] > 127 && _bgec[2] > 127 && _bgec[3] > 127 {
		return _bg
	}
	return _ba("\u0036.\u0031\u002e\u0032\u002d\u0032", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u006c\u0069\u006e\u0065\u0020\u0073\u0068\u0061\u006c\u006c b\u0065\u0020i\u006d\u006d\u0065\u0064\u0069a\u0074\u0065\u006c\u0079 \u0066\u006f\u006c\u006co\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0063\u006f\u006d\u006d\u0065n\u0074\u0020\u0063\u006f\u006e\u0073\u0069s\u0074\u0069\u006e\u0067\u0020o\u0066\u0020\u0061\u0020\u0025\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006fwe\u0064\u0020\u0062y\u0020a\u0074\u0009\u006c\u0065a\u0073\u0074\u0020f\u006f\u0075\u0072\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u002c\u0020e\u0061\u0063\u0068\u0020\u006f\u0066\u0020\u0077\u0068\u006f\u0073\u0065 \u0065\u006e\u0063\u006f\u0064e\u0064\u0020\u0062\u0079\u0074e\u0020\u0076\u0061\u006c\u0075\u0065s\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0064e\u0063\u0069\u006d\u0061\u006c \u0076\u0061\u006c\u0075\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0031\u0032\u0037\u002e")
}

type documentColorspaceOptimizeFunc func(_ffd *_ea.Document, _cdd []*_ea.Image) error

// Profile1B is the implementation of the PDF/A-1B standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile1B struct{ profile1 }

func _fggce(_cfdf *_da.PdfFont, _fggb *_ad.PdfObjectDictionary) ViolatedRule {
	const (
		_fdge = "\u0036.\u0033\u002e\u0037\u002d\u0031"
		_fdda = "\u0041\u006cl \u006e\u006f\u006e\u002d\u0073\u0079\u006db\u006f\u006c\u0069\u0063\u0020\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065\u0020\u0066o\u006e\u0074s\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020e\u0069\u0074h\u0065\u0072\u0020\u004d\u0061\u0063\u0052\u006f\u006d\u0061\u006e\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041\u006e\u0073i\u0045n\u0063\u006f\u0064\u0069n\u0067\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0066o\u0072\u0020t\u0068\u0065 \u0045n\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006b\u0065\u0079 \u0069\u006e\u0020t\u0068e\u0020\u0046o\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0072\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0066\u006f\u0072 \u0074\u0068\u0065\u0020\u0042\u0061\u0073\u0065\u0045\u006e\u0063\u006fd\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074\u0068\u0065 \u0064i\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0077\u0068\u0069\u0063\u0068\u0020\u0069s\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0074\u0068e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006be\u0079\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0046\u006f\u006e\u0074 \u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u002e\u0020\u0049\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e, \u006eo\u0020n\u006f\u006e\u002d\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072\u0075\u0065\u0054\u0079p\u0065 \u0066\u006f\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0020\u0061\u0020\u0044\u0069\u0066\u0066e\u0072\u0065\u006e\u0063\u0065\u0073\u0020a\u0072\u0072\u0061\u0079\u0020\u0075n\u006c\u0065s\u0073\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0074h\u0065\u0020\u0067\u006c\u0079\u0070\u0068\u0020\u006e\u0061\u006d\u0065\u0073 \u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0044\u0069f\u0066\u0065\u0072\u0065\u006ec\u0065\u0073\u0020a\u0072\u0072\u0061\u0079\u0020\u0061\u0072\u0065\u0020\u006c\u0069\u0073\u0074\u0065\u0064 \u0069\u006e \u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065 G\u006c\u0079\u0070\u0068\u0020\u004c\u0069\u0073t\u0020\u0061\u006e\u0064\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066o\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072a\u006d\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073t\u0020\u0074\u0068\u0065\u0020\u004d\u0069\u0063\u0072o\u0073o\u0066\u0074\u0020\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0028\u0033\u002c\u0031 \u2013 P\u006c\u0061\u0074\u0066\u006f\u0072\u006d\u0020I\u0044\u003d\u0033\u002c\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067 I\u0044\u003d\u0031\u0029\u0020\u0065\u006e\u0063\u006f\u0064i\u006e\u0067 \u0069\u006e\u0020t\u0068\u0065\u0020'\u0063\u006d\u0061\u0070\u0027\u0020\u0074\u0061\u0062\u006c\u0065\u002e"
	)
	var _efgag string
	if _fgbb, _cbgeg := _ad.GetName(_fggb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cbgeg {
		_efgag = _fgbb.String()
	}
	if _efgag != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _bg
	}
	_eeec := _cfdf.FontDescriptor()
	_geeggb, _bfdb := _ad.GetIntVal(_eeec.Flags)
	if !_bfdb {
		_ca.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _ba(_fdge, _fdda)
	}
	_cfba := (uint32(_geeggb) >> 3) != 0
	if _cfba {
		return _bg
	}
	_cfcac, _bfdb := _ad.GetName(_fggb.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067"))
	if !_bfdb {
		return _ba(_fdge, _fdda)
	}
	switch _cfcac.String() {
	case "\u004d\u0061c\u0052\u006f\u006da\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067", "\u0057i\u006eA\u006e\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067":
		return _bg
	default:
		return _ba(_fdge, _fdda)
	}
}
func _faga(_gadc *_da.PdfFont, _fbdbe *_ad.PdfObjectDictionary) ViolatedRule {
	const (
		_cecg = "\u0036.\u0033\u002e\u0037\u002d\u0033"
		_bfgb = "\u0046\u006f\u006e\u0074\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006d\u0073\u0027\u0020\u0022\u0063\u006d\u0061\u0070\u0022\u0020\u0074\u0061\u0062\u006c\u0065\u0073\u0020\u0066\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0073\u0079\u006d\u0062o\u006c\u0069c\u0020\u0054\u0072\u0075e\u0054\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073 \u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0020\u0065\u0078\u0061\u0063\u0074\u006cy\u0020\u006f\u006ee\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u002e"
	)
	var _acgb string
	if _aaaca, _cedce := _ad.GetName(_fbdbe.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cedce {
		_acgb = _aaaca.String()
	}
	if _acgb != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _bg
	}
	_daggg := _gadc.FontDescriptor()
	_cfbad, _geda := _ad.GetIntVal(_daggg.Flags)
	if !_geda {
		_ca.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _ba(_cecg, _bfgb)
	}
	_edfa := (uint32(_cfbad) >> 3) != 0
	if !_edfa {
		return _bg
	}
	return _bg
}
func _ebc() standardType { return standardType{_efe: 1, _ee: "\u0042"} }

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

func _afdbea(_acgd *_da.CompliancePdfReader) (_edcc []ViolatedRule) {
	_cecaa := true
	_acdb, _bceb := _acgd.GetCatalogMarkInfo()
	if !_bceb {
		_cecaa = false
	} else {
		_edeg, _dabbc := _ad.GetDict(_acdb)
		if _dabbc {
			_bdaab, _aeefb := _ad.GetBool(_edeg.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
			if !bool(*_bdaab) || !_aeefb {
				_cecaa = false
			}
		} else {
			_cecaa = false
		}
	}
	if !_cecaa {
		_edcc = append(_edcc, _ba("\u0036.\u0038\u002e\u0032\u002d\u0031", "\u0054\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006cog\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u0020M\u0061r\u006b\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061 \u004d\u0061\u0072\u006b\u0065\u0064\u0020\u0065\u006et\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074,\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0076\u0061lu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0072\u0075\u0065"))
	}
	_bcc, _bceb := _acgd.GetCatalogStructTreeRoot()
	if !_bceb {
		_edcc = append(_edcc, _ba("\u0036.\u0038\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u006c\u006f\u0067\u0069\u0063\u0061\u006c\u0020\u0073\u0074\u0072\u0075\u0063\u0074\u0075r\u0065\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067 \u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065d \u0062\u0079\u0020a\u0020s\u0074\u0072\u0075\u0063\u0074\u0075\u0072e\u0020\u0068\u0069\u0065\u0072\u0061\u0072\u0063\u0068\u0079\u0020\u0072\u006f\u006ft\u0065\u0064\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065\u0065\u0052\u006f\u006f\u0074\u0020\u0065\u006e\u0074r\u0079\u0020\u006f\u0066\u0020\u0074h\u0065\u0020d\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u0061t\u0061\u006c\u006fg \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0062\u0065\u0064\u0020\u0069n\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0039\u002e\u0036\u002e"))
	}
	_deaf, _bceb := _ad.GetDict(_bcc)
	if _bceb {
		_aedg, _baef := _ad.GetName(_deaf.Get("\u0052o\u006c\u0065\u004d\u0061\u0070"))
		if _baef {
			_dedgf, _cfedf := _ad.GetDict(_aedg)
			if _cfedf {
				for _, _ggee := range _dedgf.Keys() {
					_gead := _dedgf.Get(_ggee)
					if _gead == nil {
						_edcc = append(_edcc, _ba("\u0036.\u0038\u002e\u0033\u002d\u0032", "\u0041\u006c\u006c\u0020\u006eo\u006e\u002ds\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0073t\u0072\u0075\u0063\u0074ure\u0020\u0074\u0079\u0070\u0065s\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u006d\u0061\u0070\u0070\u0065d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020n\u0065\u0061\u0072\u0065\u0073\u0074\u0020\u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0073\u0074a\u006ed\u0061r\u0064\u0020\u0074\u0079\u0070\u0065\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006ee\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065re\u006e\u0063e\u0020\u0039\u002e\u0037\u002e\u0034\u002c\u0020i\u006e\u0020\u0074\u0068e\u0020\u0072\u006fl\u0065\u0020\u006d\u0061p \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006f\u0066 \u0074h\u0065\u0020\u0073\u0074\u0072\u0075c\u0074\u0075r\u0065\u0020\u0074\u0072e\u0065\u0020\u0072\u006f\u006ft\u002e"))
					}
				}
			}
		}
	}
	return _edcc
}
func _badd(_cbabb *_da.CompliancePdfReader, _eefa standardType, _bdfgc bool) (_dace []ViolatedRule) {
	_eaea, _bbbgd := _acgab(_cbabb)
	if !_bbbgd {
		return []ViolatedRule{_ba("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0063a\u0074a\u006c\u006f\u0067\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064\u002e")}
	}
	_egfe := _eaea.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	if _egfe == nil {
		return []ViolatedRule{_ba("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0054\u0068\u0065\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u006f\u0066\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074ai\u006e\u0020\u0074\u0068\u0065\u0020\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0068\u006f\u0073\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0061\u0020m\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020s\u0074\u0072\u0065\u0061\u006d")}
	}
	_eddgd, _bbbgd := _ad.GetStream(_egfe)
	if !_bbbgd {
		return []ViolatedRule{_ba("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d1", "\u0054\u0068\u0065\u0020\u0043\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072y\u0020\u006f\u0066\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074ai\u006e\u0020\u0074\u0068\u0065\u0020\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u006b\u0065\u0079\u0020\u0077\u0068\u006f\u0073\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0061\u0020m\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020s\u0074\u0072\u0065\u0061\u006d")}
	}
	_edbbg, _ddef := _gf.LoadDocument(_eddgd.Stream)
	if _ddef != nil {
		return []ViolatedRule{_ba("\u0036.\u0036\u002e\u0032\u002e\u0031\u002d4", "\u0041\u006c\u006c\u0020\u006de\u0074\u0061\u0064a\u0074\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020i\u006e \u0074\u0068\u0065\u0020\u0050\u0044\u0046 \u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0066\u006f\u0072\u006d\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065ci\u0066\u0069\u0063\u0061\u0074\u0069\u006fn\u002e\u0020\u0041\u006c\u006c\u0020c\u006fn\u0074\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0061\u006c\u006c\u0020\u0058\u004d\u0050\u0020p\u0061\u0063\u006b\u0065\u0074\u0073 \u0073h\u0061\u006c\u006c \u0062\u0065\u0020\u0077\u0065\u006c\u006c\u002d\u0066o\u0072\u006de\u0064")}
	}
	_egab := _edbbg.GetGoXmpDocument()
	var _dbdfg []*_ga.Namespace
	for _, _gfdf := range _egab.Namespaces() {
		switch _gfdf.Name {
		case _cb.NsDc.Name, _f.NsPDF.Name, _ef.NsXmp.Name, _eb.NsXmpRights.Name, _ebe.Namespace.Name, _be.Namespace.Name, _cf.NsXmpMM.Name, _be.FieldNS.Name, _be.SchemaNS.Name, _be.PropertyNS.Name, "\u0073\u0074\u0045v\u0074", "\u0073\u0074\u0056e\u0072", "\u0073\u0074\u0052e\u0066", "\u0073\u0074\u0044i\u006d", "\u0078a\u0070\u0047\u0049\u006d\u0067", "\u0073\u0074\u004ao\u0062", "\u0078\u006d\u0070\u0069\u0064\u0071":
			continue
		}
		_dbdfg = append(_dbdfg, _gfdf)
	}
	_cfddc := true
	_eddcd, _ddef := _edbbg.GetPdfaExtensionSchemas()
	if _ddef == nil {
		for _, _dedbb := range _dbdfg {
			var _adgac bool
			for _fecgf := range _eddcd {
				if _dedbb.URI == _eddcd[_fecgf].NamespaceURI {
					_adgac = true
					break
				}
			}
			if !_adgac {
				_cfddc = false
				break
			}
		}
	} else {
		_cfddc = false
	}
	if !_cfddc {
		_dace = append(_dace, _ba("\u0036.\u0036\u002e\u0032\u002e\u0033\u002d7", "\u0041\u006c\u006c\u0020\u0070\u0072\u006f\u0070e\u0072\u0074\u0069e\u0073\u0020\u0073\u0070\u0065\u0063i\u0066\u0069\u0065\u0064\u0020\u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072m\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0075s\u0065\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0073\u0063he\u006da\u0073 \u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002c\u0020\u0049\u0053\u004f\u0020\u0031\u00390\u0030\u0035-\u0031\u0020\u006f\u0072\u0020\u0074h\u0069s\u0020\u0070\u0061\u0072\u0074\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u0020\u0031\u0039\u0030\u0030\u0035\u002c\u0020o\u0072\u0020\u0061\u006e\u0079\u0020e\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073c\u0068\u0065\u006das\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006fm\u0070\u006c\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0036\u002e\u0036\u002e\u0032.\u0033\u002e\u0032\u002e"))
	}
	_ggdfe, _bbbgd := _edbbg.GetPdfAID()
	if !_bbbgd {
		_dace = append(_dace, _ba("\u0036.\u0036\u002e\u0034\u002d\u0031", "\u0054\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0061n\u0064\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020\u006c\u0065\u0076\u0065l\u0020\u006f\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0070e\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073\u0063h\u0065\u006da."))
	} else {
		if _ggdfe.Part != _eefa._efe {
			_dace = append(_dace, _ba("\u0036.\u0036\u002e\u0034\u002d\u0032", "\u0054h\u0065\u0020\u0076\u0061lue\u0020\u006f\u0066\u0020p\u0064\u0066\u0061\u0069\u0064\u003a\u0070\u0061\u0072\u0074 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0072\u0074\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u002019\u0030\u0030\u0035 \u0074\u006f\u0020\u0077\u0068i\u0063h\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065 \u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0073\u002e"))
		}
		if _eefa._ee == "\u0041" && _ggdfe.Conformance != "\u0041" {
			_dace = append(_dace, _ba("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		} else if _eefa._ee == "\u0055" && (_ggdfe.Conformance != "\u0041" && _ggdfe.Conformance != "\u0055") {
			_dace = append(_dace, _ba("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		} else if _eefa._ee == "\u0042" && (_ggdfe.Conformance != "\u0041" && _ggdfe.Conformance != "\u0042" && _ggdfe.Conformance != "\u0055") {
			_dace = append(_dace, _ba("\u0036.\u0036\u002e\u0034\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076\u0065\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061l\u006c\u0020\u0073\u0070ec\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020as\u0020\u0041\u002e\u0020\u0041 \u004c\u0065v\u0065\u006c\u0020\u0042\u0020c\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061lu\u0065\u0020o\u0066 \u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e\u0020\u0041\u0020\u004c\u0065\u0076\u0065\u006c \u0055\u0020\u0063\u006f\u006e\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063\u0069\u0066\u0079 \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff\u0020\u0070\u0064f\u0061i\u0064\u003ac\u006fn\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065 \u0061\u0073\u0020\u0055."))
		}
	}
	return _dace
}
func _aegab(_agga *Profile1Options) {
	if _agga.Now == nil {
		_agga.Now = _a.Now
	}
}

// Profile1A is the implementation of the PDF/A-1A standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile1A struct{ profile1 }
type standardType struct {
	_efe int
	_ee  string
}

func _gfege(_dcg *_ea.Document, _caed bool) error {
	_dbde, _eccc := _dcg.GetPages()
	if !_eccc {
		return nil
	}
	for _, _facc := range _dbde {
		_fcae := _facc.FindXObjectForms()
		for _, _dbgd := range _fcae {
			_dcgf, _egd := _da.NewXObjectFormFromStream(_dbgd)
			if _egd != nil {
				return _egd
			}
			_gef, _egd := _dcgf.GetContentStream()
			if _egd != nil {
				return _egd
			}
			_agfd := _cad.NewContentStreamParser(string(_gef))
			_gcg, _egd := _agfd.Parse()
			if _egd != nil {
				return _egd
			}
			_acd, _egd := _dge(_dcgf.Resources, _gcg, _caed)
			if _egd != nil {
				return _egd
			}
			if len(_acd) == 0 {
				continue
			}
			if _egd = _dcgf.SetContentStream(_acd, _ad.NewFlateEncoder()); _egd != nil {
				return _egd
			}
			_dcgf.ToPdfObject()
		}
	}
	return nil
}
func _gff(_abd *_ea.Document) (*_ad.PdfObjectDictionary, bool) {
	_ffe, _bbg := _abd.FindCatalog()
	if !_bbg {
		return nil, false
	}
	_fdd, _bbg := _ad.GetArray(_ffe.Object.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_bbg {
		return nil, false
	}
	if _fdd.Len() == 0 {
		return nil, false
	}
	return _ad.GetDict(_fdd.Get(0))
}
func _bceg(_fefd *_da.CompliancePdfReader) ViolatedRule {
	_caege := map[*_ad.PdfObjectStream]struct{}{}
	for _, _fbaab := range _fefd.PageList {
		if _fbaab.Resources == nil && _fbaab.Contents == nil {
			continue
		}
		if _ffage := _fbaab.GetPageDict(); _ffage != nil {
			_cbfd, _bffa := _ad.GetDict(_ffage.Get("\u0047\u0072\u006fu\u0070"))
			if _bffa {
				if _gfccb := _cbfd.Get("\u0053"); _gfccb != nil {
					_fcfdfe, _gegd := _ad.GetName(_gfccb)
					if _gegd && _fcfdfe.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						return _ba("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
					}
				}
			}
		}
		if _fbaab.Resources != nil {
			if _gfdc, _baee := _ad.GetDict(_fbaab.Resources.XObject); _baee {
				for _, _dfgf := range _gfdc.Keys() {
					_efcec, _bab := _ad.GetStream(_gfdc.Get(_dfgf))
					if !_bab {
						continue
					}
					if _, _dedg := _caege[_efcec]; _dedg {
						continue
					}
					_babd, _bab := _ad.GetDict(_efcec.Get("\u0047\u0072\u006fu\u0070"))
					if !_bab {
						_caege[_efcec] = struct{}{}
						continue
					}
					_aefg := _babd.Get("\u0053")
					if _aefg != nil {
						_bage, _bcfbc := _ad.GetName(_aefg)
						if _bcfbc && _bage.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
							return _ba("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
						}
					}
					_caege[_efcec] = struct{}{}
					continue
				}
			}
		}
		if _fbaab.Contents != nil {
			_faeg, _bebgg := _fbaab.GetContentStreams()
			if _bebgg != nil {
				continue
			}
			for _, _dfdd := range _faeg {
				_fcbba, _dfcfg := _cad.NewContentStreamParser(_dfdd).Parse()
				if _dfcfg != nil {
					continue
				}
				for _, _dcdde := range *_fcbba {
					if len(_dcdde.Params) == 0 {
						continue
					}
					_dacaf, _cfcef := _ad.GetName(_dcdde.Params[0])
					if !_cfcef {
						continue
					}
					_bagcd, _ffebd := _fbaab.Resources.GetXObjectByName(*_dacaf)
					if _ffebd != _da.XObjectTypeForm {
						continue
					}
					if _, _bad := _caege[_bagcd]; _bad {
						continue
					}
					_cabb, _cfcef := _ad.GetDict(_bagcd.Get("\u0047\u0072\u006fu\u0070"))
					if !_cfcef {
						_caege[_bagcd] = struct{}{}
						continue
					}
					_dccdd := _cabb.Get("\u0053")
					if _dccdd != nil {
						_fcdf, _cdbdg := _ad.GetName(_dccdd)
						if _cdbdg && _fcdf.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
							return _ba("\u0036\u002e\u0034-\u0033", "\u0041\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020\u0053\u0020\u0078Ob\u006a\u0065c\u0074\u0020\u0077\u0069\u0074h\u0020\u0061\u0020\u0076a\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0066\u006f\u0072\u006d\u0020\u0058\u004f\u0062je\u0063\u0074\u002e\n\u0041 \u0047\u0072\u006f\u0075p\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u006e\u0020S\u0020\u0078\u004fb\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020v\u0061\u006c\u0075\u0065\u0020o\u0066\u0020\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006ec\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020i\u006e\u0063\u006c\u0075\u0064e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0070\u0061\u0067\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
						}
					}
					_caege[_bagcd] = struct{}{}
				}
			}
		}
	}
	return _bg
}
func _ecff(_bddb *_da.CompliancePdfReader) (_cfbgb ViolatedRule) {
	_fbaf, _gdgc := _acgab(_bddb)
	if !_gdgc {
		return _bg
	}
	if _fbaf.Get("\u0041\u0041") != nil {
		return _ba("\u0036.\u0036\u002e\u0032\u002d\u0033", "\u0054\u0068e\u0020\u0064\u006f\u0063\u0075\u006d\u0065n\u0074 \u0063\u0061\u0074a\u006c\u006f\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065n\u0074r\u0079 \u0066\u006f\u0072 \u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063\u0074i\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e")
	}
	return _bg
}
func _afda(_afb *_da.PdfFont, _cgfa *_ad.PdfObjectDictionary) ViolatedRule {
	const (
		_adab = "\u0036.\u0033\u002e\u0037\u002d\u0032"
		_acac = "\u0041l\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072u\u0065\u0054\u0079p\u0065\u0020\u0066\u006f\u006e\u0074s\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065n\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068e\u0020\u0066\u006f\u006e\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"
	)
	var _dcfe string
	if _aebd, _bbcc := _ad.GetName(_cgfa.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _bbcc {
		_dcfe = _aebd.String()
	}
	if _dcfe != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _bg
	}
	_egeg := _afb.FontDescriptor()
	_ccbdb, _edff := _ad.GetIntVal(_egeg.Flags)
	if !_edff {
		_ca.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _ba(_adab, _acac)
	}
	_cfdg := (uint32(_ccbdb) >> 3) & 1
	_edeee := _cfdg != 0
	if !_edeee {
		return _bg
	}
	if _cgfa.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067") != nil {
		return _ba(_adab, _acac)
	}
	return _bg
}
func _fae(_bagd *_ea.Document, _gbc []pageColorspaceOptimizeFunc, _bga []documentColorspaceOptimizeFunc) error {
	_ebf, _dbff := _bagd.GetPages()
	if !_dbff {
		return nil
	}
	var _aggg []*_ea.Image
	for _cdg, _deb := range _ebf {
		_efga, _aacb := _deb.FindXObjectImages()
		if _aacb != nil {
			return _aacb
		}
		for _, _ggd := range _gbc {
			if _aacb = _ggd(_bagd, &_ebf[_cdg], _efga); _aacb != nil {
				return _aacb
			}
		}
		_aggg = append(_aggg, _efga...)
	}
	for _, _ggfe := range _bga {
		if _cga := _ggfe(_bagd, _aggg); _cga != nil {
			return _cga
		}
	}
	return nil
}
func _faad(_aeff *_da.CompliancePdfReader) (*_da.PdfOutputIntent, bool) {
	_ceafe, _bacec := _caee(_aeff)
	if !_bacec {
		return nil, false
	}
	_dafa, _dgea := _da.NewPdfOutputIntentFromPdfObject(_ceafe)
	if _dgea != nil {
		return nil, false
	}
	return _dafa, true
}

// Profile is the model.StandardImplementer enhanced by the information about the profile conformance level.
type Profile interface {
	_da.StandardImplementer
	Conformance() string
	Part() int
}

func _ecgd(_caae *_ea.Document) {
	_gafb, _dabd := _caae.FindCatalog()
	if !_dabd {
		return
	}
	_bfdc, _dabd := _gafb.GetMarkInfo()
	if !_dabd {
		_bfdc = _ad.MakeDict()
	}
	_dgdf, _dabd := _ad.GetBool(_bfdc.Get("\u004d\u0061\u0072\u006b\u0065\u0064"))
	if !_dabd || !bool(*_dgdf) {
		_bfdc.Set("\u004d\u0061\u0072\u006b\u0065\u0064", _ad.MakeBool(true))
		_gafb.SetMarkInfo(_bfdc)
	}
}
func _fdg(_ede bool, _dced standardType) (pageColorspaceOptimizeFunc, documentColorspaceOptimizeFunc) {
	var _bgad, _aea, _efda bool
	_facb := func(_ece *_ea.Document, _fbgd *_ea.Page, _gdac []*_ea.Image) error {
		for _, _gebg := range _gdac {
			switch _gebg.Colorspace {
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
				_aea = true
			case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
				_bgad = true
			case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
				_efda = true
			}
		}
		_gbda, _ddd := _fbgd.GetContents()
		if !_ddd {
			return nil
		}
		for _, _bgfed := range _gbda {
			_dffb, _fef := _bgfed.GetData()
			if _fef != nil {
				continue
			}
			_cfb := _cad.NewContentStreamParser(string(_dffb))
			_cbb, _fef := _cfb.Parse()
			if _fef != nil {
				continue
			}
			for _, _dfd := range *_cbb {
				switch _dfd.Operand {
				case "\u0047", "\u0067":
					_aea = true
				case "\u0052\u0047", "\u0072\u0067":
					_bgad = true
				case "\u004b", "\u006b":
					_efda = true
				case "\u0043\u0053", "\u0063\u0073":
					if len(_dfd.Params) == 0 {
						continue
					}
					_adgd, _ded := _ad.GetName(_dfd.Params[0])
					if !_ded {
						continue
					}
					switch _adgd.String() {
					case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
						_bgad = true
					case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
						_aea = true
					case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
						_efda = true
					}
				}
			}
		}
		_gdca := _fbgd.FindXObjectForms()
		for _, _gdbb := range _gdca {
			_cafa := _cad.NewContentStreamParser(string(_gdbb.Stream))
			_ggeg, _feag := _cafa.Parse()
			if _feag != nil {
				continue
			}
			for _, _gdg := range *_ggeg {
				switch _gdg.Operand {
				case "\u0047", "\u0067":
					_aea = true
				case "\u0052\u0047", "\u0072\u0067":
					_bgad = true
				case "\u004b", "\u006b":
					_efda = true
				case "\u0043\u0053", "\u0063\u0073":
					if len(_gdg.Params) == 0 {
						continue
					}
					_gfg, _ffeb := _ad.GetName(_gdg.Params[0])
					if !_ffeb {
						continue
					}
					switch _gfg.String() {
					case "\u0052\u0047\u0042", "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":
						_bgad = true
					case "\u0047", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":
						_aea = true
					case "\u0043\u004d\u0059\u004b", "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":
						_efda = true
					}
				}
			}
			_faa, _bfc := _ad.GetArray(_fbgd.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
			if !_bfc {
				return nil
			}
			for _, _cfbc := range _faa.Elements() {
				_aec, _debd := _ad.GetDict(_cfbc)
				if !_debd {
					continue
				}
				_ebfe := _aec.Get("\u0043")
				if _ebfe == nil {
					continue
				}
				_gag, _debd := _ad.GetArray(_ebfe)
				if !_debd {
					continue
				}
				switch _gag.Len() {
				case 0:
				case 1:
					_aea = true
				case 3:
					_bgad = true
				case 4:
					_efda = true
				}
			}
		}
		return nil
	}
	_facf := func(_aaff *_ea.Document, _dbbd []*_ea.Image) error {
		_cfcf, _cgb := _aaff.FindCatalog()
		if !_cgb {
			return nil
		}
		_facbg, _cgb := _cfcf.GetOutputIntents()
		if _cgb && _facbg.Len() > 0 {
			return nil
		}
		if !_cgb {
			_facbg = _cfcf.NewOutputIntents()
		}
		if !(_bgad || _efda || _aea) {
			return nil
		}
		defer _cfcf.SetOutputIntents(_facbg)
		if _bgad && !_efda && !_aea {
			return _acdf(_aaff, _dced, _facbg)
		}
		if _efda && !_bgad && !_aea {
			return _edga(_dced, _facbg)
		}
		if _aea && !_bgad && !_efda {
			return _aaa(_dced, _facbg)
		}
		if _bgad && _efda {
			if _aba := _bea(_dbbd, _ede); _aba != nil {
				return _aba
			}
			if _afd := _eadd(_aaff, _ede); _afd != nil {
				return _afd
			}
			if _bed := _gfege(_aaff, _ede); _bed != nil {
				return _bed
			}
			if _ebfa := _ggffc(_aaff, _ede); _ebfa != nil {
				return _ebfa
			}
			if _ede {
				return _edga(_dced, _facbg)
			}
			return _acdf(_aaff, _dced, _facbg)
		}
		return nil
	}
	return _facb, _facf
}
func _bgd(_gggbd *_ea.Document) error {
	_bdf, _aega := _gggbd.FindCatalog()
	if !_aega {
		return nil
	}
	_, _aega = _ad.GetDict(_bdf.Object.Get("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074"))
	if !_aega {
		_aaf := _ad.MakeDict()
		_aaf.Set("\u0054\u0079\u0070\u0065", _ad.MakeName("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074"))
		_bdf.Object.Set("\u0053\u0074\u0072\u0075\u0063\u0074\u0054\u0072\u0065e\u0052\u006f\u006f\u0074", _aaf)
	}
	return nil
}
func _gcdca(_eadb *_ea.Document) error {
	_cedc, _acab := _eadb.FindCatalog()
	if !_acab {
		return _gg.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_faab, _acab := _ad.GetDict(_cedc.Object.Get("\u004e\u0061\u006de\u0073"))
	if !_acab {
		return nil
	}
	if _faab.Get("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073") != nil {
		_faab.Remove("\u0041\u006c\u0074\u0065rn\u0061\u0074\u0065\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006fn\u0073")
	}
	return nil
}

var _bg = ViolatedRule{}

func _dadf(_afa *_da.CompliancePdfReader) (_befa ViolatedRule) {
	for _, _fbgb := range _afa.GetObjectNums() {
		_dfff, _gbcfb := _afa.GetIndirectObjectByNumber(_fbgb)
		if _gbcfb != nil {
			continue
		}
		_cccf, _acc := _ad.GetStream(_dfff)
		if !_acc {
			continue
		}
		_aabf, _acc := _ad.GetName(_cccf.Get("\u0054\u0079\u0070\u0065"))
		if !_acc {
			continue
		}
		if *_aabf != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_ecb, _acc := _ad.GetName(_cccf.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"))
		if !_acc {
			continue
		}
		if *_ecb == "\u0050\u0053" {
			return _ba("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
		if _cccf.Get("\u0050\u0053") != nil {
			return _ba("\u0036.\u0032\u002e\u0035\u002d\u0031", "A\u0020\u0066\u006fr\u006d\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006ft\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032\u0020\u006b\u0065\u0079 \u0077\u0069\u0074\u0068\u0020a\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0050\u0053\u0020o\u0072\u0020\u0074\u0068e\u0020\u0050\u0053\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _befa
}
func _ebge(_dcbda *_da.CompliancePdfReader) (_caeeg []ViolatedRule) {
	var _dfgdg, _bgdfa bool
	_gfaa := func() bool { return _dfgdg && _bgdfa }
	for _, _adbcg := range _dcbda.GetObjectNums() {
		_geba, _bfed := _dcbda.GetIndirectObjectByNumber(_adbcg)
		if _bfed != nil {
			_ca.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0025\u0064\u0020fa\u0069\u006c\u0065d\u003a \u0025\u0076", _adbcg, _bfed)
			continue
		}
		_dfbgd, _daff := _ad.GetDict(_geba)
		if !_daff {
			continue
		}
		_dbda, _daff := _ad.GetName(_dfbgd.Get("\u0054\u0079\u0070\u0065"))
		if !_daff {
			continue
		}
		if *_dbda != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_aaadg, _daff := _ad.GetName(_dfbgd.Get("\u0053"))
		if !_daff {
			if !_dfgdg {
				_caeeg = append(_caeeg, _ba("\u0036.\u0036\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004c\u0061\u0075\u006e\u0063\u0068\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046o\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u006e\u0064 \u004a\u0061\u0076a\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020s\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e \u0041\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020th\u0065\u0020\u0064\u0065p\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020s\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u002d\u006f\u0070\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062e\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e\u0020T\u0068\u0065\u0020\u0048\u0069\u0064\u0065\u0020a\u0063\u0074\u0069\u006f\u006e \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_dfgdg = true
				if _gfaa() {
					return _caeeg
				}
			}
			continue
		}
		switch _da.PdfActionType(*_aaadg) {
		case _da.ActionTypeLaunch, _da.ActionTypeSound, _da.ActionTypeMovie, _da.ActionTypeResetForm, _da.ActionTypeImportData, _da.ActionTypeJavaScript:
			if !_dfgdg {
				_caeeg = append(_caeeg, _ba("\u0036.\u0036\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004c\u0061\u0075\u006e\u0063\u0068\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046o\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u006e\u0064 \u004a\u0061\u0076a\u0053\u0063\u0072\u0069\u0070\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020s\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e \u0041\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020th\u0065\u0020\u0064\u0065p\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020s\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u002d\u006f\u0070\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068a\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062e\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064\u002e\u0020T\u0068\u0065\u0020\u0048\u0069\u0064\u0065\u0020a\u0063\u0074\u0069\u006f\u006e \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_dfgdg = true
				if _gfaa() {
					return _caeeg
				}
			}
			continue
		case _da.ActionTypeNamed:
			if !_bgdfa {
				_dfaff, _cfaa := _ad.GetName(_dfbgd.Get("\u004e"))
				if !_cfaa {
					_caeeg = append(_caeeg, _ba("\u0036.\u0036\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_bgdfa = true
					if _gfaa() {
						return _caeeg
					}
					continue
				}
				switch *_dfaff {
				case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
				default:
					_caeeg = append(_caeeg, _ba("\u0036.\u0036\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_bgdfa = true
					if _gfaa() {
						return _caeeg
					}
					continue
				}
			}
		}
	}
	return _caeeg
}
func _gcgef(_fgagf *_da.CompliancePdfReader) (_agfc []ViolatedRule) {
	var _cgcdb, _fcda bool
	_gafgf := func() bool { return _cgcdb && _fcda }
	for _, _ccbgb := range _fgagf.GetObjectNums() {
		_eddc, _fagaf := _fgagf.GetIndirectObjectByNumber(_ccbgb)
		if _fagaf != nil {
			_ca.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0077\u0069\u0074\u0068 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0025\u0064\u0020fa\u0069\u006c\u0065d\u003a \u0025\u0076", _ccbgb, _fagaf)
			continue
		}
		_abbb, _feada := _ad.GetDict(_eddc)
		if !_feada {
			continue
		}
		_ccfe, _feada := _ad.GetName(_abbb.Get("\u0054\u0079\u0070\u0065"))
		if !_feada {
			continue
		}
		if *_ccfe != "\u0041\u0063\u0074\u0069\u006f\u006e" {
			continue
		}
		_ffff, _feada := _ad.GetName(_abbb.Get("\u0053"))
		if !_feada {
			if !_cgcdb {
				_agfc = append(_agfc, _ba("\u0036.\u0035\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004caun\u0063\u0068\u002c\u0020S\u006f\u0075\u006e\u0064,\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046\u006f\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044a\u0074\u0061,\u0020\u0048\u0069\u0064\u0065\u002c\u0020\u0053\u0065\u0074\u004f\u0043\u0047\u0053\u0074\u0061\u0074\u0065\u002c\u0020\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u002c\u0020T\u0072\u0061\u006e\u0073\u002c\u0020\u0047o\u0054\u006f\u0033\u0044\u0056\u0069\u0065\u0077\u0020\u0061\u006e\u0064\u0020\u004a\u0061v\u0061Sc\u0072\u0069p\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020\u0070\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e \u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020t\u0068\u0065\u0020\u0064\u0065\u0070\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020\u0073\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u006f\u0070\u0020\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070e\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_cgcdb = true
				if _gafgf() {
					return _agfc
				}
			}
			continue
		}
		switch _da.PdfActionType(*_ffff) {
		case _da.ActionTypeLaunch, _da.ActionTypeSound, _da.ActionTypeMovie, _da.ActionTypeResetForm, _da.ActionTypeImportData, _da.ActionTypeJavaScript, _da.ActionTypeHide, _da.ActionTypeSetOCGState, _da.ActionTypeRendition, _da.ActionTypeTrans, _da.ActionTypeGoTo3DView:
			if !_cgcdb {
				_agfc = append(_agfc, _ba("\u0036.\u0035\u002e\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u004caun\u0063\u0068\u002c\u0020S\u006f\u0075\u006e\u0064,\u0020\u004d\u006f\u0076\u0069\u0065\u002c\u0020\u0052\u0065\u0073\u0065\u0074\u0046\u006f\u0072\u006d\u002c\u0020\u0049\u006d\u0070\u006f\u0072\u0074\u0044a\u0074\u0061,\u0020\u0048\u0069\u0064\u0065\u002c\u0020\u0053\u0065\u0074\u004f\u0043\u0047\u0053\u0074\u0061\u0074\u0065\u002c\u0020\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u002c\u0020T\u0072\u0061\u006e\u0073\u002c\u0020\u0047o\u0054\u006f\u0033\u0044\u0056\u0069\u0065\u0077\u0020\u0061\u006e\u0064\u0020\u004a\u0061v\u0061Sc\u0072\u0069p\u0074\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074 \u0062\u0065\u0020\u0070\u0065\u0072m\u0069\u0074\u0074\u0065\u0064\u002e \u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020t\u0068\u0065\u0020\u0064\u0065\u0070\u0072\u0065\u0063\u0061\u0074\u0065\u0064\u0020\u0073\u0065\u0074\u002d\u0073\u0074\u0061\u0074\u0065\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u006f\u0070\u0020\u0061c\u0074\u0069\u006f\u006e\u0073\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070e\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
				_cgcdb = true
				if _gafgf() {
					return _agfc
				}
			}
			continue
		case _da.ActionTypeNamed:
			if !_fcda {
				_abce, _fadcb := _ad.GetName(_abbb.Get("\u004e"))
				if !_fadcb {
					_agfc = append(_agfc, _ba("\u0036.\u0035\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_fcda = true
					if _gafgf() {
						return _agfc
					}
					continue
				}
				switch *_abce {
				case "\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065", "\u0050\u0072\u0065\u0076\u0050\u0061\u0067\u0065", "\u0046i\u0072\u0073\u0074\u0050\u0061\u0067e", "\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065":
				default:
					_agfc = append(_agfc, _ba("\u0036.\u0035\u002e\u0031\u002d\u0032", "N\u0061\u006d\u0065\u0064\u0020\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u006f\u0074\u0068e\u0072\u0020\u0074h\u0061\u006e\u0020\u004e\u0065\u0078\u0074\u0050\u0061\u0067\u0065\u002c\u0020P\u0072\u0065v\u0050\u0061\u0067\u0065\u002c\u0020\u0046\u0069\u0072\u0073\u0074\u0050a\u0067e\u002c\u0020\u0061\u006e\u0064\u0020\u004c\u0061\u0073\u0074\u0050\u0061\u0067\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_fcda = true
					if _gafgf() {
						return _agfc
					}
					continue
				}
			}
		}
	}
	return _agfc
}
func _dgbf(_ebaae *_da.PdfFont, _cggb *_ad.PdfObjectDictionary, _aeee bool) ViolatedRule {
	const (
		_gdbd  = "\u0036.\u0033\u002e\u0034\u002d\u0031"
		_eceaa = "\u0054\u0068\u0065\u0020\u0066\u006f\u006et\u0020\u0070\u0072\u006f\u0067\u0072\u0061\u006d\u0073\u0020\u0066\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0069\u006e \u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069l\u0065\u0020s\u0068\u0061\u006cl\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0077\u0069\u0074\u0068i\u006e\u0020\u0074h\u0061\u0074\u0020\u0066\u0069\u006ce\u002c\u0020a\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052e\u0066\u0065\u0072\u0065\u006e\u0063\u0065 \u0035\u002e\u0038\u002c\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0077h\u0065\u006e\u0020\u0074\u0068\u0065 \u0066\u006f\u006e\u0074\u0073\u0020\u0061\u0072\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0065\u0078\u0063\u006cu\u0073i\u0076\u0065\u006c\u0079\u0020\u0077\u0069t\u0068\u0020\u0074\u0065\u0078\u0074\u0020\u0072e\u006ed\u0065\u0072\u0069\u006e\u0067\u0020\u006d\u006f\u0064\u0065\u0020\u0033\u002e"
	)
	if _aeee {
		return _bg
	}
	_dcgb := _ebaae.FontDescriptor()
	var _ccda string
	if _beaa, _fdgcg := _ad.GetName(_cggb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _fdgcg {
		_ccda = _beaa.String()
	}
	switch _ccda {
	case "\u0054\u0079\u0070e\u0031":
		if _dcgb.FontFile == nil {
			return _ba(_gdbd, _eceaa)
		}
	case "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065":
		if _dcgb.FontFile2 == nil {
			return _ba(_gdbd, _eceaa)
		}
	case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0033":
	default:
		if _dcgb.FontFile3 == nil {
			return _ba(_gdbd, _eceaa)
		}
	}
	return _bg
}
func _ggffc(_abe *_ea.Document, _eacd bool) error {
	_fga, _fdgd := _abe.GetPages()
	if !_fdgd {
		return nil
	}
	for _, _cfbcc := range _fga {
		_edg, _gdf := _ad.GetArray(_cfbcc.Object.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if !_gdf {
			continue
		}
		for _, _ecdc := range _edg.Elements() {
			_eee, _afdc := _ad.GetDict(_ecdc)
			if !_afdc {
				continue
			}
			_gabg := _eee.Get("\u0043")
			if _gabg == nil {
				continue
			}
			_cag, _afdc := _ad.GetArray(_gabg)
			if !_afdc {
				continue
			}
			_fgea, _ffcb := _cag.GetAsFloat64Slice()
			if _ffcb != nil {
				return _ffcb
			}
			switch _cag.Len() {
			case 0, 1:
				if _eacd {
					_eee.Set("\u0043", _ad.MakeArrayFromIntegers([]int{1, 1, 1, 1}))
				} else {
					_eee.Set("\u0043", _ad.MakeArrayFromIntegers([]int{1, 1, 1}))
				}
			case 3:
				if _eacd {
					_eddf, _fcdg, _fed, _gadg := _g.RGBToCMYK(uint8(_fgea[0]*255), uint8(_fgea[1]*255), uint8(_fgea[2]*255))
					_eee.Set("\u0043", _ad.MakeArrayFromFloats([]float64{float64(_eddf) / 255, float64(_fcdg) / 255, float64(_fed) / 255, float64(_gadg) / 255}))
				}
			case 4:
				if !_eacd {
					_deba, _cece, _cfaf := _g.CMYKToRGB(uint8(_fgea[0]*255), uint8(_fgea[1]*255), uint8(_fgea[2]*255), uint8(_fgea[3]*255))
					_eee.Set("\u0043", _ad.MakeArrayFromFloats([]float64{float64(_deba) / 255, float64(_cece) / 255, float64(_cfaf) / 255}))
				}
			}
		}
	}
	return nil
}
func _bcfde(_ccdaf *_da.CompliancePdfReader) ViolatedRule {
	_aegf := _ccdaf.ParserMetadata()
	if _aegf.HasInvalidSeparationAfterXRef() {
		return _ba("\u0036.\u0031\u002e\u0034\u002d\u0032", "\u0054\u0068\u0065 \u0078\u0072\u0065\u0066\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0063\u0072\u006f\u0073s\u0020\u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065 s\u0075b\u0073\u0065\u0063ti\u006f\u006e\u0020\u0068\u0065\u0061\u0064e\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079 \u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e")
	}
	return _bg
}
func _dgce(_bcbfd *_ad.PdfObjectDictionary, _cbdd map[*_ad.PdfObjectStream][]byte, _fgfg map[*_ad.PdfObjectStream]*_b.CMap) ViolatedRule {
	const (
		_cbgf  = "\u0046\u006f\u0072\u0020\u0061\u006e\u0079\u0020\u0067\u0069\u0076\u0065\u006e\u0020\u0063\u006f\u006d\u0070o\u0073\u0069\u0074e\u0020\u0028\u0054\u0079\u0070\u0065\u0020\u0030\u0029 \u0066\u006fn\u0074\u0020\u0077\u0069\u0074\u0068\u0069\u006e \u0061\u0020\u0063\u006fn\u0066\u006fr\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f \u0065\u006e\u0074\u0072\u0079\u0020\u0069\u006e\u0020\u0069\u0074\u0073 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020\u0074\u0068\u0065\u0020\u0066\u006fl\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0072\u0065l\u0061t\u0069\u006f\u006e\u0073\u0068\u0069\u0070. \u0049\u0066\u0020\u0074\u0068\u0065\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006b\u0065\u0079 \u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0054\u0079\u0070\u0065\u0020\u0030 \u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0069\u0073\u0020I\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056\u002c\u0020\u0061\u006e\u0079\u0020v\u0061\u006c\u0075\u0065\u0073\u0020\u006f\u0066\u0020\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u002c\u0020\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u002c\u0020\u0061\u006e\u0064\u0020\u0053up\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006d\u0061\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0069n\u0020\u0074h\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066\u006f\u0020\u0065\u006e\u0074r\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044F\u006f\u006e\u0074\u002e\u0020\u004f\u0074\u0068\u0065\u0072\u0077\u0069\u0073\u0065\u002c\u0020\u0074\u0068\u0065\u0020\u0063\u006f\u0072\u0072\u0065\u0073\u0070\u006f\u006e\u0064\u0069\u006e\u0067\u0020\u0052\u0065\u0067\u0069\u0073\u0074\u0072\u0079\u0020a\u006e\u0064\u0020\u004f\u0072\u0064\u0065\u0072\u0069\u006e\u0067\u0020s\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0069\u006e\u0020\u0062\u006f\u0074h\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065m\u0049\u006e\u0066\u006f\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006cl\u0020\u0062\u0065\u0020i\u0064en\u0074\u0069\u0063\u0061\u006c\u002c \u0061n\u0064\u0020\u0074\u0068\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0053\u0075\u0070\u0070l\u0065\u006d\u0065\u006e\u0074 \u006b\u0065\u0079\u0020\u0069\u006e\u0020t\u0068\u0065\u0020\u0043I\u0044S\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066o\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0067re\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f t\u0068\u0065\u0020\u0053\u0075\u0070\u0070\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0043\u0049\u0044\u0053\u0079\u0073\u0074\u0065\u006d\u0049\u006e\u0066o\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0043M\u0061p\u002e"
		_gfggc = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0031"
	)
	var _edgg string
	if _eccbg, _dfdeg := _ad.GetName(_bcbfd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _dfdeg {
		_edgg = _eccbg.String()
	}
	if _edgg != "\u0054\u0079\u0070e\u0030" {
		return _bg
	}
	_deda := _bcbfd.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _bdeag, _gfcad := _ad.GetName(_deda); _gfcad {
		switch _bdeag.String() {
		case "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048", "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0056":
			return _bg
		}
		_degg, _ecbg := _b.LoadPredefinedCMap(_bdeag.String())
		if _ecbg != nil {
			return _ba(_gfggc, _cbgf)
		}
		_caaa := _degg.CIDSystemInfo()
		if _caaa.Ordering != _caaa.Registry {
			return _ba(_gfggc, _cbgf)
		}
		return _bg
	}
	_aag, _afbg := _ad.GetStream(_deda)
	if !_afbg {
		return _ba(_gfggc, _cbgf)
	}
	_cbdg, _dfaa := _ggag(_aag, _cbdd, _fgfg)
	if _dfaa != nil {
		return _ba(_gfggc, _cbgf)
	}
	_dcfg := _cbdg.CIDSystemInfo()
	if _dcfg.Ordering != _dcfg.Registry {
		return _ba(_gfggc, _cbgf)
	}
	return _bg
}
func _caee(_fdaef *_da.CompliancePdfReader) (*_ad.PdfObjectDictionary, bool) {
	_bdfd, _bebdf := _acgab(_fdaef)
	if !_bebdf {
		return nil, false
	}
	_ccfg, _bebdf := _ad.GetArray(_bdfd.Get("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073"))
	if !_bebdf {
		return nil, false
	}
	if _ccfg.Len() == 0 {
		return nil, false
	}
	return _ad.GetDict(_ccfg.Get(0))
}
func _cagb(_cdaf *_da.PdfInfo, _feb *_gf.Document) bool {
	_bdbed, _gafgd := _feb.GetPdfInfo()
	if !_gafgd {
		return false
	}
	if _bdbed.InfoDict == nil {
		return false
	}
	_ccge, _geee := _da.NewPdfInfoFromObject(_bdbed.InfoDict)
	if _geee != nil {
		return false
	}
	if _cdaf.Creator != nil {
		if _ccge.Creator == nil || _ccge.Creator.String() != _cdaf.Creator.String() {
			return false
		}
	}
	if _cdaf.CreationDate != nil {
		if _ccge.CreationDate == nil || !_ccge.CreationDate.ToGoTime().Equal(_cdaf.CreationDate.ToGoTime()) {
			return false
		}
	}
	if _cdaf.ModifiedDate != nil {
		if _ccge.ModifiedDate == nil || !_ccge.ModifiedDate.ToGoTime().Equal(_cdaf.ModifiedDate.ToGoTime()) {
			return false
		}
	}
	if _cdaf.Producer != nil {
		if _ccge.Producer == nil || _ccge.Producer.String() != _cdaf.Producer.String() {
			return false
		}
	}
	if _cdaf.Keywords != nil {
		if _ccge.Keywords == nil || _ccge.Keywords.String() != _cdaf.Keywords.String() {
			return false
		}
	}
	if _cdaf.Trapped != nil {
		if _ccge.Trapped == nil {
			return false
		}
		switch _cdaf.Trapped.String() {
		case "\u0054\u0072\u0075\u0065":
			if _ccge.Trapped.String() != "\u0054\u0072\u0075\u0065" {
				return false
			}
		case "\u0046\u0061\u006cs\u0065":
			if _ccge.Trapped.String() != "\u0046\u0061\u006cs\u0065" {
				return false
			}
		default:
			if _ccge.Trapped.String() != "\u0046\u0061\u006cs\u0065" {
				return false
			}
		}
	}
	if _cdaf.Title != nil {
		if _ccge.Title == nil || _ccge.Title.String() != _cdaf.Title.String() {
			return false
		}
	}
	if _cdaf.Subject != nil {
		if _ccge.Subject == nil || _ccge.Subject.String() != _cdaf.Subject.String() {
			return false
		}
	}
	return true
}
func _dcaf(_cffa *_ea.Document) error {
	_beeb, _ebbe := _cffa.FindCatalog()
	if !_ebbe {
		return _gg.New("\u0063\u0061\u0074\u0061\u006c\u006f\u0067\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_feg, _ebbe := _ad.GetDict(_beeb.Object.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"))
	if !_ebbe {
		return nil
	}
	_gce, _ebbe := _ad.GetDict(_feg.Get("\u0044"))
	if _ebbe {
		if _gce.Get("\u0041\u0053") != nil {
			_gce.Remove("\u0041\u0053")
		}
	}
	_ebec, _ebbe := _ad.GetArray(_feg.Get("\u0043o\u006e\u0066\u0069\u0067\u0073"))
	if _ebbe {
		for _eaaag := 0; _eaaag < _ebec.Len(); _eaaag++ {
			_fce, _gecgb := _ad.GetDict(_ebec.Get(_eaaag))
			if !_gecgb {
				continue
			}
			if _fce.Get("\u0041\u0053") != nil {
				_fce.Remove("\u0041\u0053")
			}
		}
	}
	return nil
}
func _fbed(_cfea *_da.CompliancePdfReader) (_dgead []ViolatedRule) { return _dgead }
func _gaae(_gcdf *_da.CompliancePdfReader) ViolatedRule {
	for _, _ccaaf := range _gcdf.GetObjectNums() {
		_eaed, _dafg := _gcdf.GetIndirectObjectByNumber(_ccaaf)
		if _dafg != nil {
			continue
		}
		_efed, _gfec := _ad.GetStream(_eaed)
		if !_gfec {
			continue
		}
		_ebgbb, _gfec := _ad.GetName(_efed.Get("\u0054\u0079\u0070\u0065"))
		if !_gfec {
			continue
		}
		if *_ebgbb != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		if _efed.Get("\u0053\u004d\u0061s\u006b") != nil {
			return _ba("\u0036\u002e\u0034-\u0032", "\u0041\u006e\u0020\u0058\u004f\u0062\u006a\u0065\u0063\u0074\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068e \u0053\u004d\u0061\u0073\u006b\u0020\u006b\u0065\u0079\u002e")
		}
	}
	return _bg
}
func _baafg(_decd *_ea.Document) error {
	_dea := func(_dbeg *_ad.PdfObjectDictionary) error {
		if _dbeg.Get("\u0054\u0052") != nil {
			_ca.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0054\u0052\u0020\u006b\u0065\u0079")
			_dbeg.Remove("\u0054\u0052")
		}
		_gbce := _dbeg.Get("\u0054\u0052\u0032")
		if _gbce != nil {
			_dfdf := _gbce.String()
			if _dfdf != "\u0044e\u0066\u0061\u0075\u006c\u0074" {
				_ca.Log.Debug("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074\u0065 o\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073 \u0054\u00522\u0020\u006b\u0065y\u0020\u0077\u0069\u0074\u0068\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0074\u0068\u0065r\u0020\u0074ha\u006e\u0020\u0044e\u0066\u0061\u0075\u006c\u0074")
				_dbeg.Set("\u0054\u0052\u0032", _ad.MakeName("\u0044e\u0066\u0061\u0075\u006c\u0074"))
			}
		}
		if _dbeg.Get("\u0048\u0054\u0050") != nil {
			_ca.Log.Debug("\u0045\u0078\u0074\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0063\u006f\u006et\u0061\u0069\u006e\u0073\u0020\u0048\u0054P\u0020\u006b\u0065\u0079")
			_dbeg.Remove("\u0048\u0054\u0050")
		}
		_eadaf := _dbeg.Get("\u0042\u004d")
		if _eadaf != nil {
			_cbe, _fggd := _ad.GetName(_eadaf)
			if !_fggd {
				_ca.Log.Debug("E\u0078\u0074\u0047\u0053\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u0027\u0042\u004d\u0027\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u004e\u0061m\u0065")
				_cbe = _ad.MakeName("")
			}
			_dgef := _cbe.String()
			switch _dgef {
			case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065", "\u004d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "\u0053\u0063\u0072\u0065\u0065\u006e", "\u004fv\u0065\u0072\u006c\u0061\u0079", "\u0044\u0061\u0072\u006b\u0065\u006e", "\u004ci\u0067\u0068\u0074\u0065\u006e", "\u0043\u006f\u006c\u006f\u0072\u0044\u006f\u0064\u0067\u0065", "\u0043o\u006c\u006f\u0072\u0042\u0075\u0072n", "\u0048a\u0072\u0064\u004c\u0069\u0067\u0068t", "\u0053o\u0066\u0074\u004c\u0069\u0067\u0068t", "\u0044\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065", "\u0045x\u0063\u006c\u0075\u0073\u0069\u006fn", "\u0048\u0075\u0065", "\u0053\u0061\u0074\u0075\u0072\u0061\u0074\u0069\u006f\u006e", "\u0043\u006f\u006co\u0072", "\u004c\u0075\u006d\u0069\u006e\u006f\u0073\u0069\u0074\u0079":
			default:
				_dbeg.Set("\u0042\u004d", _ad.MakeName("\u004e\u006f\u0072\u006d\u0061\u006c"))
			}
		}
		return nil
	}
	_adgb, _cdf := _decd.GetPages()
	if !_cdf {
		return nil
	}
	for _, _caafa := range _adgb {
		_deg, _egaf := _caafa.GetResources()
		if !_egaf {
			continue
		}
		_cgag, _cbdb := _ad.GetDict(_deg.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
		if !_cbdb {
			return nil
		}
		_beca := _cgag.Keys()
		for _, _afg := range _beca {
			_ebce, _cfcg := _ad.GetDict(_cgag.Get(_afg))
			if !_cfcg {
				continue
			}
			_facg := _dea(_ebce)
			if _facg != nil {
				continue
			}
		}
	}
	for _, _fcfag := range _adgb {
		_dfdff, _fcaea := _fcfag.GetContents()
		if !_fcaea {
			return nil
		}
		for _, _eaaa := range _dfdff {
			_aaed, _cebb := _eaaa.GetData()
			if _cebb != nil {
				continue
			}
			_ada := _cad.NewContentStreamParser(string(_aaed))
			_gaebf, _cebb := _ada.Parse()
			if _cebb != nil {
				continue
			}
			for _, _cbbfg := range *_gaebf {
				if len(_cbbfg.Params) == 0 {
					continue
				}
				_, _egfgf := _ad.GetName(_cbbfg.Params[0])
				if !_egfgf {
					continue
				}
				_cff, _baabb := _fcfag.GetResourcesXObject()
				if !_baabb {
					continue
				}
				for _, _gddb := range _cff.Keys() {
					_gfca, _dabc := _ad.GetStream(_cff.Get(_gddb))
					if !_dabc {
						continue
					}
					_egg, _dabc := _ad.GetDict(_gfca.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
					if !_dabc {
						continue
					}
					_cabf, _dabc := _ad.GetDict(_egg.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
					if !_dabc {
						continue
					}
					for _, _gecb := range _cabf.Keys() {
						_acaf, _aaag := _ad.GetDict(_cabf.Get(_gecb))
						if !_aaag {
							continue
						}
						_effg := _dea(_acaf)
						if _effg != nil {
							continue
						}
					}
				}
			}
		}
	}
	return nil
}
func _dcdec(_adcg *_ad.PdfObjectDictionary, _ggdfg map[*_ad.PdfObjectStream][]byte, _cdbe map[*_ad.PdfObjectStream]*_b.CMap) ViolatedRule {
	const (
		_gefdd = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0033\u002d\u0034"
		_aadfe = "\u0046\u006f\u0072\u0020\u0074\u0068\u006fs\u0065\u0020\u0043\u004d\u0061\u0070\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0061\u0072e\u0020\u0065m\u0062\u0065\u0064de\u0064\u002c\u0020\u0074\u0068\u0065\u0020\u0069\u006et\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0057\u004d\u006f\u0064\u0065\u0020\u0065\u006e\u0074r\u0079\u0020i\u006e t\u0068\u0065\u0020CM\u0061\u0070\u0020\u0064\u0069\u0063\u0074\u0069o\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0069\u0064\u0065\u006e\u0074\u0069\u0063\u0061\u006c\u0020\u0074\u006f \u0074h\u0065\u0020\u0057\u004d\u006f\u0064e\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0074h\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064ed\u0020\u0043\u004d\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"
	)
	var _defe string
	if _egfad, _bgecfd := _ad.GetName(_adcg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _bgecfd {
		_defe = _egfad.String()
	}
	if _defe != "\u0054\u0079\u0070e\u0030" {
		return _bg
	}
	_afgdc := _adcg.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")
	if _, _ffaf := _ad.GetName(_afgdc); _ffaf {
		return _bg
	}
	_adbf, _afgg := _ad.GetStream(_afgdc)
	if !_afgg {
		return _ba(_gefdd, _aadfe)
	}
	_fgcge, _dfgba := _ggag(_adbf, _ggdfg, _cdbe)
	if _dfgba != nil {
		return _ba(_gefdd, _aadfe)
	}
	_fcebc, _feabe := _ad.GetIntVal(_adbf.Get("\u0057\u004d\u006fd\u0065"))
	_ecafb, _abc := _fgcge.WMode()
	if _feabe && _abc {
		if _ecafb != _fcebc {
			return _ba(_gefdd, _aadfe)
		}
	}
	if (_feabe && !_abc) || (!_feabe && _abc) {
		return _ba(_gefdd, _aadfe)
	}
	return _bg
}

// String gets a string representation of the violated rule.
func (_add ViolatedRule) String() string {
	return _d.Sprintf("\u0025\u0073\u003a\u0020\u0025\u0073", _add.RuleNo, _add.Detail)
}
func _bfaae(_fdbff *_da.CompliancePdfReader) (_agcfb []ViolatedRule) {
	var _ceffg, _aga, _abdaf, _facfb, _ecfab, _ffdd, _gedb bool
	_fcaaf := func() bool { return _ceffg && _aga && _abdaf && _facfb && _ecfab && _ffdd && _gedb }
	for _, _ecfeb := range _fdbff.PageList {
		_begaf, _ecdbg := _ecfeb.GetAnnotations()
		if _ecdbg != nil {
			_ca.Log.Trace("\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0061\u006en\u006f\u0074\u0061\u0074\u0069\u006f\u006es\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ecdbg)
			continue
		}
		for _, _afggc := range _begaf {
			if !_ceffg {
				switch _afggc.GetContext().(type) {
				case *_da.PdfAnnotationScreen, *_da.PdfAnnotation3D, *_da.PdfAnnotationSound, *_da.PdfAnnotationMovie, nil:
					_agcfb = append(_agcfb, _ba("\u0036.\u0033\u002e\u0031\u002d\u0031", "\u0041nn\u006f\u0074\u0061\u0074i\u006f\u006e t\u0079\u0070\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072e\u006e\u0063\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065r\u006d\u0069t\u0074\u0065\u0064\u002e\u0020\u0041\u0064d\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u006c\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0033\u0044\u002c\u0020\u0053\u006f\u0075\u006e\u0064\u002c\u0020\u0053\u0063\u0072\u0065\u0065\u006e\u0020\u0061n\u0064\u0020\u004d\u006f\u0076\u0069\u0065\u0020\u0074\u0079\u0070\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064\u002e"))
					_ceffg = true
					if _fcaaf() {
						return _agcfb
					}
				}
			}
			_dbaa, _fgefe := _ad.GetDict(_afggc.GetContainingPdfObject())
			if !_fgefe {
				continue
			}
			_, _egafe := _afggc.GetContext().(*_da.PdfAnnotationPopup)
			if !_egafe && !_aga {
				_, _febg := _ad.GetIntVal(_dbaa.Get("\u0046"))
				if !_febg {
					_agcfb = append(_agcfb, _ba("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0045\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072i\u0065\u0073\u0020\u0077\u0068\u006fs\u0065\u0020\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0069\u0073\u0020\u0050\u006f\u0070u\u0070\u002c\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0046 \u006b\u0065y."))
					_aga = true
					if _fcaaf() {
						return _agcfb
					}
				}
			}
			if !_abdaf {
				_deeba, _edcgf := _ad.GetIntVal(_dbaa.Get("\u0046"))
				if _edcgf && !(_deeba&4 == 4 && _deeba&1 == 0 && _deeba&2 == 0 && _deeba&32 == 0 && _deeba&256 == 0) {
					_agcfb = append(_agcfb, _ba("\u0036.\u0033\u002e\u0032\u002d\u0032", "I\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u0020\u0046 \u006b\u0065\u0079\u0027\u0073\u0020\u0050\u0072\u0069\u006e\u0074\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0020\u0073\u0068\u0061l\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0074\u0020\u0074\u006f\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0069\u0074\u0073\u0020\u0048\u0069\u0064\u0064\u0065\u006e\u002c\u0020\u0049\u006e\u0076\u0069\u0073\u0069\u0062\u006c\u0065\u002c\u0020\u0054\u006f\u0067\u0067\u006c\u0065\u004e\u006f\u0056\u0069\u0065\u0077\u002c\u0020\u0061\u006e\u0064 \u004eo\u0056\u0069\u0065\u0077\u0020\u0066\u006c\u0061\u0067\u0020\u0062\u0069\u0074\u0073\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020s\u0065\u0074\u0020t\u006f\u0020\u0030."))
					_abdaf = true
					if _fcaaf() {
						return _agcfb
					}
				}
			}
			_, _gbdec := _afggc.GetContext().(*_da.PdfAnnotationText)
			if _gbdec && !_facfb {
				_addd, _dccda := _ad.GetIntVal(_dbaa.Get("\u0046"))
				if _dccda && !(_addd&8 == 8 && _addd&16 == 16) {
					_agcfb = append(_agcfb, _ba("\u0036.\u0033\u002e\u0032\u002d\u0033", "\u0054\u0065\u0078\u0074\u0020a\u006e\u006e\u006f\u0074\u0061t\u0069o\u006e\u0020\u0068\u0061\u0073\u0020\u006f\u006e\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006ca\u0067\u0073\u0020\u004e\u006f\u005a\u006f\u006f\u006d\u0020\u006f\u0072\u0020\u004e\u006f\u0052\u006f\u0074\u0061\u0074\u0065\u0020\u0073\u0065t\u0020\u0074\u006f\u0020\u0030\u002e"))
					_facfb = true
					if _fcaaf() {
						return _agcfb
					}
				}
			}
			if !_ecfab {
				_ecaa, _gedc := _ad.GetDict(_dbaa.Get("\u0041\u0050"))
				if _gedc {
					_gcbe := _ecaa.Get("\u004e")
					if _gcbe == nil || len(_ecaa.Keys()) > 1 {
						_agcfb = append(_agcfb, _ba("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_ecfab = true
						if _fcaaf() {
							return _agcfb
						}
						continue
					}
					_, _efefc := _afggc.GetContext().(*_da.PdfAnnotationWidget)
					if _efefc {
						_adbd, _ddbe := _ad.GetName(_dbaa.Get("\u0046\u0054"))
						if _ddbe && *_adbd == "\u0042\u0074\u006e" {
							if _, _fgcc := _ad.GetDict(_gcbe); !_fgcc {
								_agcfb = append(_agcfb, _ba("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
								_ecfab = true
								if _fcaaf() {
									return _agcfb
								}
								continue
							}
						}
					}
					_, _eaaee := _ad.GetStream(_gcbe)
					if !_eaaee {
						_agcfb = append(_agcfb, _ba("\u0036.\u0033\u002e\u0033\u002d\u0032", "\u0046\u006f\u0072\u0020\u0061\u006c\u006c\u0020\u0061\u006e\u006e\u006ft\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063t\u0069\u006f\u006ea\u0072\u0069\u0065\u0073 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0069\u006e\u0067\u0020\u0061\u006e\u0020\u0041\u0050 \u006b\u0065\u0079\u002c\u0020\u0074\u0068\u0065\u0020\u0061p\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0073\u0020\u0061\u0073\u0020it\u0073\u0020\u0076\u0061\u006cu\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020o\u006e\u006c\u0079\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u002e\u0020\u0049\u0066\u0020\u0061\u006e\u0020\u0061\u006e\u006e\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0020\u0064i\u0063\u0074\u0069o\u006e\u0061\u0072\u0079\u0027\u0073\u0020\u0053\u0075\u0062ty\u0070\u0065\u0020\u006b\u0065\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0057\u0069\u0064g\u0065\u0074\u0020\u0061\u006e\u0064\u0020\u0069\u0074s\u0020\u0046\u0054 \u006be\u0079\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020B\u0074\u006e,\u0020\u0074he \u0076a\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u004e\u0020\u006b\u0065\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0064\u0069\u0063\u0074\u0069\u006fn\u0061r\u0079; \u006f\u0074\u0068\u0065\u0072\u0077\u0069s\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020th\u0065\u0020N\u0020\u006b\u0065y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062e\u0020\u0061\u006e\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0061n\u0063\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"))
						_ecfab = true
						if _fcaaf() {
							return _agcfb
						}
						continue
					}
				}
			}
			_gafc, _gdff := _afggc.GetContext().(*_da.PdfAnnotationWidget)
			if !_gdff {
				continue
			}
			if !_ffdd {
				if _gafc.A != nil {
					_agcfb = append(_agcfb, _ba("\u0036.\u0034\u002e\u0031\u002d\u0031", "A \u0057\u0069d\u0067\u0065\u0074\u0020\u0061\u006e\u006e\u006f\u0074a\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006ec\u006cu\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0020e\u006et\u0072\u0079."))
					_ffdd = true
					if _fcaaf() {
						return _agcfb
					}
				}
			}
			if !_gedb {
				if _gafc.AA != nil {
					_agcfb = append(_agcfb, _ba("\u0036.\u0034\u002e\u0031\u002d\u0031", "\u0041\u0020\u0057\u0069\u0064\u0067\u0065\u0074\u0020\u0061\u006e\u006eo\u0074\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u0020\u0073h\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061\u006e\u0020\u0041\u0041\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0061d\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u002d\u0061\u0063t\u0069\u006f\u006e\u0073\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					_gedb = true
					if _fcaaf() {
						return _agcfb
					}
				}
			}
		}
	}
	return _agcfb
}
func _cdac(_baeaa *_da.CompliancePdfReader) ViolatedRule { return _bg }

// Profile1Options are the options that changes the way how optimizer may try to adapt document into PDF/A standard.
type Profile1Options struct {

	// CMYKDefaultColorSpace is an option that refers PDF/A-1
	CMYKDefaultColorSpace bool

	// Now is a function that returns current time.
	Now func() _a.Time

	// Xmp is the xmp options information.
	Xmp XmpOptions
}

func _dddf(_efdb *_da.CompliancePdfReader) (_fgdca []ViolatedRule) {
	_dfbc := func(_eafg *_ad.PdfObjectDictionary, _cgcgf *[]string, _gebeb *[]ViolatedRule) error {
		_cfcgb := _eafg.Get("\u004e\u0061\u006d\u0065")
		if _cfcgb == nil || len(_cfcgb.String()) == 0 {
			*_gebeb = append(*_gebeb, _ba("\u0036\u002e\u0039-\u0031", "\u0045\u0061\u0063\u0068\u0020o\u0070\u0074\u0069\u006f\u006e\u0061l\u0020\u0063\u006f\u006e\u0074\u0065\u006et\u0020\u0063\u006fn\u0066\u0069\u0067\u0075r\u0061\u0074\u0069\u006f\u006e\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u004e\u0061\u006d\u0065\u0020\u006b\u0065\u0079\u002e"))
		}
		for _, _acbca := range *_cgcgf {
			if _acbca == _cfcgb.String() {
				*_gebeb = append(*_gebeb, _ba("\u0036\u002e\u0039-\u0032", "\u0045\u0061\u0063\u0068\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061l\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0063\u006f\u006e\u0066\u0069\u0067\u0075\u0072a\u0074\u0069\u006fn\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020N\u0061\u006d\u0065\u0020\u006b\u0065\u0079\u002c w\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020s\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0075ni\u0071\u0075\u0065 \u0061\u006d\u006f\u006e\u0067\u0073\u0074\u0020\u0061\u006c\u006c\u0020o\u0070\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0063\u006fn\u0074\u0065\u006e\u0074 \u0063\u006f\u006e\u0066\u0069\u0067u\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074i\u006fn\u0061\u0072\u0069\u0065\u0073\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0074\u0068e\u0020\u0050\u0044\u0046\u002fA\u002d\u0032\u0020\u0066\u0069l\u0065\u002e"))
			} else {
				*_cgcgf = append(*_cgcgf, _cfcgb.String())
			}
		}
		if _eafg.Get("\u0041\u0053") != nil {
			*_gebeb = append(*_gebeb, _ba("\u0036\u002e\u0039-\u0034", "Th\u0065\u0020\u0041\u0053\u0020\u006b\u0065y \u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0061\u0070\u0070\u0065\u0061r\u0020\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0063\u006f\u006et\u0065\u006e\u0074\u0020\u0063\u006fn\u0066\u0069\u0067\u0075\u0072\u0061\u0074\u0069\u006fn\u0020d\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
		}
		return nil
	}
	_gedbc, _ddge := _acgab(_efdb)
	if !_ddge {
		return _fgdca
	}
	_bgdeg, _ddge := _ad.GetDict(_gedbc.Get("\u004f\u0043\u0050r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073"))
	if !_ddge {
		return _fgdca
	}
	var _gdbef []string
	_cfdc, _ddge := _ad.GetDict(_bgdeg.Get("\u0044"))
	if _ddge {
		_dfbc(_cfdc, &_gdbef, &_fgdca)
	}
	_efbee, _ddge := _ad.GetArray(_bgdeg.Get("\u0043o\u006e\u0066\u0069\u0067\u0073"))
	if _ddge {
		for _babdb := 0; _babdb < _efbee.Len(); _babdb++ {
			_agcgf, _fafb := _ad.GetDict(_efbee.Get(_babdb))
			if !_fafb {
				continue
			}
			_dfbc(_agcgf, &_gdbef, &_fgdca)
		}
	}
	return _fgdca
}
func _ba(_ace string, _bee string) ViolatedRule { return ViolatedRule{RuleNo: _ace, Detail: _bee} }
func _ceg(_fcgf *_da.CompliancePdfReader) (_dbdg []ViolatedRule) {
	_gbad := _fcgf.ParserMetadata()
	if _gbad.HasInvalidSubsectionHeader() {
		_dbdg = append(_dbdg, _ba("\u0036.\u0031\u002e\u0034\u002d\u0031", "\u006e\u0020\u0061\u0020\u0063\u0072\u006f\u0073\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0073\u0075\u0062\u0073\u0065c\u0074\u0069\u006f\u006e\u0020h\u0065a\u0064\u0065\u0072\u0020t\u0068\u0065\u0020\u0073\u0074\u0061\u0072t\u0069\u006e\u0067\u0020\u006fb\u006a\u0065\u0063\u0074 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0072\u0061n\u0067e\u0020s\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020s\u0069\u006e\u0067\u006c\u0065\u0020\u0053\u0050\u0041C\u0045\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0020\u0028\u0032\u0030\u0068\u0029\u002e"))
	}
	if _gbad.HasInvalidSeparationAfterXRef() {
		_dbdg = append(_dbdg, _ba("\u0036.\u0031\u002e\u0034\u002d\u0032", "\u0054\u0068\u0065 \u0078\u0072\u0065\u0066\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0063\u0072\u006f\u0073s\u0020\u0072\u0065\u0066e\u0072\u0065\u006e\u0063\u0065 s\u0075b\u0073\u0065\u0063ti\u006f\u006e\u0020\u0068\u0065\u0061\u0064e\u0072\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u0065\u0064\u0020\u0062\u0079 \u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u002e"))
	}
	return _dbdg
}
func _fgeb(_eeba *_da.CompliancePdfReader) (_cagfc []ViolatedRule) {
	var _aaeda, _gcde, _dbffd, _edadd, _dcca, _ffcdg bool
	_baeb := func() bool { return _aaeda && _gcde && _dbffd && _edadd && _dcca && _ffcdg }
	_cdca := func(_eccd *_ad.PdfObjectDictionary) bool {
		if !_aaeda && _eccd.Get("\u0054\u0052") != nil {
			_aaeda = true
			_cagfc = append(_cagfc, _ba("\u0036.\u0032\u002e\u0038\u002d\u0031", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0054\u0052\u0020\u006b\u0065\u0079\u002e"))
		}
		if _ddee := _eccd.Get("\u0054\u0052\u0032"); !_gcde && _ddee != nil {
			_cddd, _eagb := _ad.GetName(_ddee)
			if !_eagb || (_eagb && *_cddd != "\u0044e\u0066\u0061\u0075\u006c\u0074") {
				_gcde = true
				_cagfc = append(_cagfc, _ba("\u0036.\u0032\u002e\u0038\u002d\u0032", "\u0041\u006e \u0045\u0078\u0074G\u0053\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074a\u0069n\u0020\u0074\u0068\u0065\u0020\u0054R2 \u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076al\u0075e\u0020\u006f\u0074\u0068e\u0072 \u0074h\u0061\u006e \u0044\u0065fa\u0075\u006c\u0074\u002e"))
				if _baeb() {
					return true
				}
			}
		}
		if _fbff := _eccd.Get("\u0053\u004d\u0061s\u006b"); !_dbffd && _fbff != nil {
			_fadge, _edcga := _ad.GetName(_fbff)
			if !_edcga || (_edcga && *_fadge != "\u004e\u006f\u006e\u0065") {
				_dbffd = true
				_cagfc = append(_cagfc, _ba("\u0036\u002e\u0034-\u0031", "\u0049\u0066\u0020\u0061\u006e \u0053\u004d\u0061\u0073\u006b\u0020\u006be\u0079\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002c\u0020\u0069\u0074s\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020\u004e\u006f\u006ee\u002e"))
				if _baeb() {
					return true
				}
			}
		}
		if _eeeb := _eccd.Get("\u0043\u0041"); !_dcca && _eeeb != nil {
			_ecdd, _dbddb := _ad.GetNumberAsFloat(_eeeb)
			if _dbddb == nil && _ecdd != 1.0 {
				_dcca = true
				_cagfc = append(_cagfc, _ba("\u0036\u002e\u0034-\u0035", "\u0054\u0068\u0065\u0020\u0066ol\u006c\u006fw\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0073\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078t\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068a\u006c\u006c\u0020\u0068\u0061v\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0073 \u0073h\u006f\u0077\u006e\u003a\u0020\u0043\u0041 \u002d\u0020\u0031\u002e\u0030\u002e"))
				if _baeb() {
					return true
				}
			}
		}
		if _eeagb := _eccd.Get("\u0063\u0061"); !_ffcdg && _eeagb != nil {
			_dedb, _dffbg := _ad.GetNumberAsFloat(_eeagb)
			if _dffbg == nil && _dedb != 1.0 {
				_ffcdg = true
				_cagfc = append(_cagfc, _ba("\u0036\u002e\u0034-\u0036", "\u0054\u0068\u0065\u0020\u0066ol\u006c\u006fw\u0069\u006e\u0067\u0020\u006b\u0065\u0079\u0073\u002c\u0020\u0069\u0066\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078t\u0047\u0053\u0074a\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068a\u006c\u006c\u0020\u0068\u0061v\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0073 \u0073h\u006f\u0077\u006e\u003a\u0020\u0063\u0061 \u002d\u0020\u0031\u002e\u0030\u002e"))
				if _baeb() {
					return true
				}
			}
		}
		if _ecgg := _eccd.Get("\u0042\u004d"); !_edadd && _ecgg != nil {
			_cede, _dabca := _ad.GetName(_ecgg)
			if _dabca {
				switch _cede.String() {
				case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065":
				default:
					_edadd = true
					_cagfc = append(_cagfc, _ba("\u0036\u002e\u0034-\u0034", "T\u0068\u0065\u0020\u0066\u006f\u006cl\u006f\u0077\u0069\u006e\u0067 \u006b\u0065y\u0073\u002c\u0020\u0069\u0066 \u0070res\u0065\u006e\u0074\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0045\u0078\u0074\u0047S\u0074\u0061t\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065 \u0074\u0068\u0065 \u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0073\u0068\u006f\u0077n\u003a\u0020\u0042\u004d\u0020\u002d\u0020\u004e\u006f\u0072m\u0061\u006c\u0020\u006f\u0072\u0020\u0043\u006f\u006d\u0070\u0061t\u0069\u0062\u006c\u0065\u002e"))
					if _baeb() {
						return true
					}
				}
			}
		}
		return false
	}
	for _, _ffdad := range _eeba.PageList {
		_bgcb := _ffdad.Resources
		if _bgcb == nil {
			continue
		}
		if _bgcb.ExtGState == nil {
			continue
		}
		_fbcb, _beed := _ad.GetDict(_bgcb.ExtGState)
		if !_beed {
			continue
		}
		_aaeb := _fbcb.Keys()
		for _, _dgae := range _aaeb {
			_bfdd, _efea := _ad.GetDict(_fbcb.Get(_dgae))
			if !_efea {
				continue
			}
			if _cdca(_bfdd) {
				return _cagfc
			}
		}
	}
	for _, _gfcaf := range _eeba.PageList {
		_dgee := _gfcaf.Resources
		if _dgee == nil {
			continue
		}
		_eed, _bgae := _ad.GetDict(_dgee.XObject)
		if !_bgae {
			continue
		}
		for _, _agca := range _eed.Keys() {
			_faabg, _eebae := _ad.GetStream(_eed.Get(_agca))
			if !_eebae {
				continue
			}
			_cefe, _eebae := _ad.GetDict(_faabg.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_eebae {
				continue
			}
			_cffd, _eebae := _ad.GetDict(_cefe.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
			if !_eebae {
				continue
			}
			for _, _baabc := range _cffd.Keys() {
				_baeg, _feaf := _ad.GetDict(_cffd.Get(_baabc))
				if !_feaf {
					continue
				}
				if _cdca(_baeg) {
					return _cagfc
				}
			}
		}
	}
	return _cagfc
}
func _aaa(_gee standardType, _bdbd *_ea.OutputIntents) error {
	_efeb, _dgb := _ac.NewISOCoatedV2Gray1CBasOutputIntent(_gee.outputIntentSubtype())
	if _dgb != nil {
		return _dgb
	}
	if _dgb = _bdbd.Add(_efeb.ToPdfObject()); _dgb != nil {
		return _dgb
	}
	return nil
}

type colorspaceModification struct {
	_cc  _de.ColorConverter
	_ecd _da.PdfColorspace
}

func _adgfe(_ccae *_da.CompliancePdfReader) (_baecc []ViolatedRule) {
	var _cgbg, _deac, _ccgc, _bege bool
	_gfegc := func() bool { return _cgbg && _deac && _ccgc && _bege }
	_dddd, _cfbeb := _faad(_ccae)
	var _bgddb _ac.ProfileHeader
	if _cfbeb {
		_bgddb, _ = _ac.ParseHeader(_dddd.DestOutputProfile)
	}
	_affbe := map[_ad.PdfObject]struct{}{}
	var _afgd func(_feca _da.PdfColorspace) bool
	_afgd = func(_bgdg _da.PdfColorspace) bool {
		switch _afae := _bgdg.(type) {
		case *_da.PdfColorspaceDeviceGray:
			if !_cgbg {
				if !_cfbeb {
					_baecc = append(_baecc, _ba("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d4", "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079 \u0073\u0068\u0061\u006c\u006c\u0020\u006f\u006e\u006c\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064 \u0069\u0066\u0020\u0061\u0020\u0064\u0065v\u0069\u0063\u0065\u0020\u0069\u006e\u0064\u0065p\u0065\u006e\u0064\u0065\u006e\u0074\u0020\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0047\u0072\u0061\u0079\u0020\u0063\u006f\u006c\u006f\u0075r \u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061\u0073\u0020\u0062\u0065\u0065\u006e \u0073\u0065\u0074\u0020\u0077\u0068\u0065n \u0074\u0068\u0065\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072a\u0079\u0020\u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u002c o\u0072\u0020\u0069\u0066\u0020\u0061\u0020\u0050\u0044\u0046\u002fA\u0020\u004f\u0075tp\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u002e"))
					_cgbg = true
					if _gfegc() {
						return true
					}
				}
			}
		case *_da.PdfColorspaceDeviceRGB:
			if !_deac {
				if !_cfbeb || _bgddb.ColorSpace != _ac.ColorSpaceRGB {
					_baecc = append(_baecc, _ba("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d2", "\u0044\u0065\u0076\u0069c\u0065\u0052\u0047\u0042\u0020\u0073\u0068\u0061\u006cl\u0020\u006f\u006e\u006c\u0079\u0020\u0062e\u0020\u0075\u0073\u0065\u0064\u0020\u0069f\u0020\u0061\u0020\u0064\u0065\u0076\u0069\u0063e\u0020\u0069n\u0064\u0065\u0070e\u006e\u0064\u0065\u006et \u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0052\u0047\u0042\u0020\u0063\u006fl\u006f\u0075r\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061\u0073\u0020b\u0065\u0065\u006e\u0020s\u0065\u0074 \u0077\u0068\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u0052\u0047\u0042\u0020c\u006flou\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020i\u0073\u0020\u0075\u0073\u0065\u0064\u002c\u0020\u006f\u0072\u0020if\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044F\u002f\u0041\u0020\u004fut\u0070\u0075\u0074\u0049\u006e\u0074\u0065n\u0074\u0020t\u0068\u0061t\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0061\u006e\u0020\u0052\u0047\u0042\u0020\u0064\u0065\u0073\u0074\u0069\u006e\u0061\u0074io\u006e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u002e"))
					_deac = true
					if _gfegc() {
						return true
					}
				}
			}
		case *_da.PdfColorspaceDeviceCMYK:
			if !_ccgc {
				if !_cfbeb || _bgddb.ColorSpace != _ac.ColorSpaceCMYK {
					_baecc = append(_baecc, _ba("\u0036.\u0032\u002e\u0034\u002e\u0033\u002d3", "\u0044e\u0076\u0069c\u0065\u0043\u004d\u0059\u004b\u0020\u0073hal\u006c\u0020\u006f\u006e\u006c\u0079\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u0066\u0020\u0061\u0020\u0064\u0065\u0076\u0069\u0063\u0065\u0020\u0069\u006e\u0064\u0065\u0070\u0065\u006e\u0064\u0065\u006e\u0074\u0020\u0044ef\u0061\u0075\u006c\u0074\u0043\u004d\u0059K\u0020\u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0068\u0061s\u0020\u0062\u0065\u0065\u006e \u0073\u0065\u0074\u0020\u006fr \u0069\u0066\u0020\u0061\u0020\u0044e\u0076\u0069\u0063\u0065\u004e\u002d\u0062\u0061\u0073\u0065\u0064\u0020\u0044\u0065f\u0061\u0075\u006c\u0074\u0043\u004d\u0059\u004b\u0020c\u006f\u006c\u006f\u0075r\u0020\u0073\u0070\u0061\u0063e\u0020\u0068\u0061\u0073\u0020\u0062\u0065\u0065\u006e\u0020\u0073\u0065\u0074\u0020\u0077\u0068\u0065\u006e\u0020\u0074h\u0065\u0020\u0044\u0065\u0076\u0069c\u0065\u0043\u004d\u0059\u004b\u0020c\u006f\u006c\u006fu\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0069\u0073\u0020\u0075\u0073\u0065\u0064\u0020\u006f\u0072\u0020t\u0068\u0065\u0020\u0066\u0069l\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074p\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0061\u0020\u0043\u004d\u0059\u004b\u0020d\u0065\u0073\u0074\u0069\u006e\u0061t\u0069\u006f\u006e\u0020\u0070r\u006f\u0066\u0069\u006c\u0065\u002e"))
					_ccgc = true
					if _gfegc() {
						return true
					}
				}
			}
		case *_da.PdfColorspaceICCBased:
			if !_bege {
				_dbgg, _fgbc := _ac.ParseHeader(_afae.Data)
				if _fgbc != nil {
					_ca.Log.Debug("\u0070\u0061\u0072si\u006e\u0067\u0020\u0049\u0043\u0043\u0042\u0061\u0073e\u0064 \u0068e\u0061d\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _fgbc)
					_baecc = append(_baecc, func() ViolatedRule {
						return _ba("\u0036.\u0032\u002e\u0034\u002e\u0032\u002d1", "\u0054\u0068e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0074\u0068\u0061\u0074\u0020\u0066o\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0073\u0074r\u0065\u0061\u006d o\u0066\u0020\u0061\u006e\u0020\u0049C\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006fl\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0049\u0043\u0043.\u0031\u003a\u0031\u0039\u0039\u0038-\u0030\u0039,\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0031\u002d\u00312\u002c\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0033\u002d\u0030\u0039\u0020\u006f\u0072\u0020I\u0053\u004f\u0020\u0031\u0035\u0030\u0037\u0036\u002d\u0031\u002e")
					}())
					_bege = true
					if _gfegc() {
						return true
					}
				}
				if !_bege {
					var _gfbfg, _cefb bool
					switch _dbgg.DeviceClass {
					case _ac.DeviceClassPRTR, _ac.DeviceClassMNTR, _ac.DeviceClassSCNR, _ac.DeviceClassSPAC:
					default:
						_gfbfg = true
					}
					switch _dbgg.ColorSpace {
					case _ac.ColorSpaceRGB, _ac.ColorSpaceCMYK, _ac.ColorSpaceGRAY, _ac.ColorSpaceLAB:
					default:
						_cefb = true
					}
					if _gfbfg || _cefb {
						_baecc = append(_baecc, _ba("\u0036.\u0032\u002e\u0034\u002e\u0032\u002d1", "\u0054\u0068e\u0020\u0070\u0072\u006f\u0066\u0069\u006c\u0065\u0020\u0074\u0068\u0061\u0074\u0020\u0066o\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0073\u0074r\u0065\u0061\u006d o\u0066\u0020\u0061\u006e\u0020\u0049C\u0043\u0042\u0061\u0073\u0065\u0064\u0020\u0063\u006fl\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0049\u0043\u0043.\u0031\u003a\u0031\u0039\u0039\u0038-\u0030\u0039,\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0031\u002d\u00312\u002c\u0020\u0049\u0043\u0043\u002e\u0031\u003a\u0032\u0030\u0030\u0033\u002d\u0030\u0039\u0020\u006f\u0072\u0020I\u0053\u004f\u0020\u0031\u0035\u0030\u0037\u0036\u002d\u0031\u002e"))
						_bege = true
						if _gfegc() {
							return true
						}
					}
				}
			}
			if _afae.Alternate != nil {
				return _afgd(_afae.Alternate)
			}
		}
		return false
	}
	for _, _acdda := range _ccae.GetObjectNums() {
		_cgac, _dfacc := _ccae.GetIndirectObjectByNumber(_acdda)
		if _dfacc != nil {
			continue
		}
		_cdfbb, _gcfe := _ad.GetStream(_cgac)
		if !_gcfe {
			continue
		}
		_cadgd, _gcfe := _ad.GetName(_cdfbb.Get("\u0054\u0079\u0070\u0065"))
		if !_gcfe || _cadgd.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
			continue
		}
		_aecfg, _gcfe := _ad.GetName(_cdfbb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
		if !_gcfe {
			continue
		}
		_affbe[_cdfbb] = struct{}{}
		switch _aecfg.String() {
		case "\u0049\u006d\u0061g\u0065":
			_fccda, _abae := _da.NewXObjectImageFromStream(_cdfbb)
			if _abae != nil {
				continue
			}
			_affbe[_cdfbb] = struct{}{}
			if _afgd(_fccda.ColorSpace) {
				return _baecc
			}
		case "\u0046\u006f\u0072\u006d":
			_agea, _dfadc := _ad.GetDict(_cdfbb.Get("\u0047\u0072\u006fu\u0070"))
			if !_dfadc {
				continue
			}
			_dcegg := _agea.Get("\u0043\u0053")
			if _dcegg == nil {
				continue
			}
			_cgbge, _egcc := _da.NewPdfColorspaceFromPdfObject(_dcegg)
			if _egcc != nil {
				continue
			}
			if _afgd(_cgbge) {
				return _baecc
			}
		}
	}
	for _, _fgbbf := range _ccae.PageList {
		_gbe, _gaca := _fgbbf.GetContentStreams()
		if _gaca != nil {
			continue
		}
		for _, _bgea := range _gbe {
			_dgbfd, _aefa := _cad.NewContentStreamParser(_bgea).Parse()
			if _aefa != nil {
				continue
			}
			for _, _accg := range *_dgbfd {
				if len(_accg.Params) > 1 {
					continue
				}
				switch _accg.Operand {
				case "\u0042\u0049":
					_edeab, _baea := _accg.Params[0].(*_cad.ContentStreamInlineImage)
					if !_baea {
						continue
					}
					_eeeg, _fbab := _edeab.GetColorSpace(_fgbbf.Resources)
					if _fbab != nil {
						continue
					}
					if _afgd(_eeeg) {
						return _baecc
					}
				case "\u0044\u006f":
					_eedf, _efdf := _ad.GetName(_accg.Params[0])
					if !_efdf {
						continue
					}
					_bfef, _fedf := _fgbbf.Resources.GetXObjectByName(*_eedf)
					if _, _dbcdf := _affbe[_bfef]; _dbcdf {
						continue
					}
					switch _fedf {
					case _da.XObjectTypeImage:
						_gagg, _ebgd := _da.NewXObjectImageFromStream(_bfef)
						if _ebgd != nil {
							continue
						}
						_affbe[_bfef] = struct{}{}
						if _afgd(_gagg.ColorSpace) {
							return _baecc
						}
					case _da.XObjectTypeForm:
						_fgbe, _fcddf := _ad.GetDict(_bfef.Get("\u0047\u0072\u006fu\u0070"))
						if !_fcddf {
							continue
						}
						_eacda, _fcddf := _ad.GetName(_fgbe.Get("\u0043\u0053"))
						if !_fcddf {
							continue
						}
						_dgdg, _accge := _da.NewPdfColorspaceFromPdfObject(_eacda)
						if _accge != nil {
							continue
						}
						_affbe[_bfef] = struct{}{}
						if _afgd(_dgdg) {
							return _baecc
						}
					}
				}
			}
		}
	}
	return _baecc
}
func _ebbag(_dbefe *_da.CompliancePdfReader) (_accc ViolatedRule) {
	_bdgc, _bgcc := _acgab(_dbefe)
	if !_bgcc {
		return _bg
	}
	if _bdgc.Get("\u0052\u0065\u0071u\u0069\u0072\u0065\u006d\u0065\u006e\u0074\u0073") != nil {
		return _ba("\u0036\u002e\u0031\u0031\u002d\u0031", "Th\u0065\u0020d\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063a\u0074\u0061\u006c\u006f\u0067\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020R\u0065q\u0075\u0069\u0072\u0065\u006d\u0065\u006e\u0074s\u0020k\u0065\u0079.")
	}
	return _bg
}
func _cfgd(_dcdc *_da.CompliancePdfReader) (_ceede []ViolatedRule) {
	var _edadcg, _abda, _aggc bool
	if _dcdc.ParserMetadata().HasNonConformantStream() {
		_ceede = []ViolatedRule{_ba("\u0036.\u0031\u002e\u0037\u002d\u0032", "T\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020f\u006f\u006cl\u006fw\u0065\u0064\u0020e\u0069\u0074h\u0065\u0072\u0020\u0062\u0079\u0020\u0061 \u0043\u0041\u0052\u0052I\u0041\u0047\u0045\u0020\u0052E\u0054\u0055\u0052\u004e\u0020\u00280\u0044\u0068\u0029\u0020\u0061\u006e\u0064\u0020\u004c\u0049\u004e\u0045\u0020F\u0045\u0045\u0044\u0020\u0028\u0030\u0041\u0068\u0029\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065\u0020o\u0072\u0020\u0062\u0079\u0020\u0061 \u0073\u0069ng\u006c\u0065\u0020\u004cIN\u0045 \u0046\u0045\u0045\u0044 \u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u002e\u0020T\u0068\u0065\u0020e\u006e\u0064\u0073\u0074r\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020p\u0072\u0065\u0063\u0065\u0064\u0065\u0064\u0020\u0062\u0079\u0020\u0061n\u0020\u0045\u004f\u004c \u006d\u0061\u0072\u006b\u0065\u0072\u002e")}
	}
	for _, _agedg := range _dcdc.GetObjectNums() {
		_ggcf, _ := _dcdc.GetIndirectObjectByNumber(_agedg)
		if _ggcf == nil {
			continue
		}
		_aabd, _edgd := _ad.GetStream(_ggcf)
		if !_edgd {
			continue
		}
		if !_edadcg {
			_cecf := _aabd.Get("\u004c\u0065\u006e\u0067\u0074\u0068")
			if _cecf == nil {
				_ceede = append(_ceede, _ba("\u0036.\u0031\u002e\u0037\u002d\u0031", "\u006e\u006f\u0020'\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074"))
				_edadcg = true
			} else {
				_bcbd, _faaf := _ad.GetIntVal(_cecf)
				if !_faaf {
					_ceede = append(_ceede, _ba("\u0036.\u0031\u002e\u0037\u002d\u0031", "s\u0074\u0072\u0065\u0061\u006d\u0020\u0027\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020an\u0020\u0069\u006et\u0065g\u0065\u0072"))
					_edadcg = true
				} else {
					if len(_aabd.Stream) != _bcbd {
						_ceede = append(_ceede, _ba("\u0036.\u0031\u002e\u0037\u002d\u0031", "\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006c\u0065\u006e\u0067th\u0020v\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d"))
						_edadcg = true
					}
				}
			}
		}
		if !_abda {
			if _aabd.Get("\u0046") != nil {
				_abda = true
				_ceede = append(_ceede, _ba("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
			}
			if _aabd.Get("\u0046F\u0069\u006c\u0074\u0065\u0072") != nil && !_abda {
				_abda = true
				_ceede = append(_ceede, _ba("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
			if _aabd.Get("\u0046\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u0061\u006d\u0073") != nil && !_abda {
				_abda = true
				_ceede = append(_ceede, _ba("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
		}
		if !_aggc {
			_bfdbe, _ecge := _ad.GetName(_ad.TraceToDirectObject(_aabd.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
			if !_ecge {
				continue
			}
			if *_bfdbe == _ad.StreamEncodingFilterNameLZW {
				_aggc = true
				_ceede = append(_ceede, _ba("\u0036.\u0031\u002e\u0037\u002d\u0034", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e"))
			}
		}
	}
	return _ceede
}
func _acegf(_gggde *_da.CompliancePdfReader) (_fgfc []ViolatedRule) {
	_acca := _gggde.GetObjectNums()
	for _, _ebgea := range _acca {
		_geaf, _fcfae := _gggde.GetIndirectObjectByNumber(_ebgea)
		if _fcfae != nil {
			continue
		}
		_accac, _egegf := _ad.GetDict(_geaf)
		if !_egegf {
			continue
		}
		_bfcd, _egegf := _ad.GetName(_accac.Get("\u0054\u0079\u0070\u0065"))
		if !_egegf {
			continue
		}
		if _bfcd.String() != "\u0046\u0069\u006c\u0065\u0073\u0070\u0065\u0063" {
			continue
		}
		if _accac.Get("\u0045\u0046") != nil {
			if _accac.Get("\u0046") == nil || _accac.Get("\u0045\u0046") == nil {
				_fgfc = append(_fgfc, _ba("\u0036\u002e\u0038-\u0032", "\u0054h\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0070\u0065\u0063\u0069\u0066i\u0063\u0061\u0074i\u006f\u006e\u0020\u0064\u0069\u0063t\u0069\u006fn\u0061\u0072\u0079\u0020\u0066\u006f\u0072\u0020\u0061\u006e\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006cl\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020t\u0068\u0065\u0020\u0046\u0020a\u006e\u0064\u0020\u0055\u0046\u0020\u006b\u0065\u0079\u0073\u002e"))
			}
			if _accac.Get("\u0041\u0046\u0052\u0065\u006c\u0061\u0074\u0069\u006fn\u0073\u0068\u0069\u0070") == nil {
				_fgfc = append(_fgfc, _ba("\u0036\u002e\u0038-\u0033", "\u0049\u006e\u0020\u006f\u0072d\u0065\u0072\u0020\u0074\u006f\u0020\u0065\u006e\u0061\u0062\u006c\u0065\u0020i\u0064\u0065nt\u0069\u0066\u0069c\u0061\u0074\u0069o\u006e\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0074\u0068\u0065\u0020fi\u006ce\u0020\u0073\u0070\u0065\u0063\u0069f\u0069c\u0061\u0074\u0069o\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020c\u006f\u006e\u0074e\u006e\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u0069\u0073\u0020\u0072\u0065\u0066\u0065\u0072\u0072\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0069\u0074\u002c\u0020\u0061\u0020\u006e\u0065\u0077\u0020(\u0072\u0065\u0071\u0075i\u0072\u0065\u0064\u0029\u0020\u006be\u0079\u0020h\u0061\u0073\u0020\u0062e\u0065\u006e\u0020\u0064\u0065\u0066i\u006e\u0065\u0064\u0020a\u006e\u0064\u0020\u0069\u0074s \u0070\u0072e\u0073\u0065n\u0063\u0065\u0020\u0028\u0069\u006e\u0020\u0074\u0068e\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079\u0029\u0020\u0069\u0073\u0020\u0072\u0065q\u0075\u0069\u0072e\u0064\u002e"))
			}
			break
		}
	}
	return _fgfc
}
func _cgcdc(_aacbd *_da.CompliancePdfReader) (_cege []ViolatedRule) {
	var _aaedf, _cdbdd, _edeb, _egbd, _ggbg, _egca, _bedaf bool
	_dceaa := func() bool { return _aaedf && _cdbdd && _edeb && _egbd && _ggbg && _egca && _bedaf }
	_ggad := func(_cbef *_ad.PdfObjectDictionary) bool {
		if !_aaedf && _cbef.Get("\u0054\u0052") != nil {
			_aaedf = true
			_cege = append(_cege, _ba("\u0036.\u0032\u002e\u0035\u002d\u0031", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072y\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0054\u0052\u0020\u006b\u0065\u0079\u002e"))
		}
		if _dbef := _cbef.Get("\u0054\u0052\u0032"); !_cdbdd && _dbef != nil {
			_ecfg, _cecag := _ad.GetName(_dbef)
			if !_cecag || (_cecag && *_ecfg != "\u0044e\u0066\u0061\u0075\u006c\u0074") {
				_cdbdd = true
				_cege = append(_cege, _ba("\u0036.\u0032\u002e\u0035\u002d\u0032", "\u0041\u006e \u0045\u0078\u0074G\u0053\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069on\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074a\u0069n\u0020\u0074\u0068\u0065\u0020\u0054R2 \u006b\u0065\u0079\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0076al\u0075e\u0020\u006f\u0074\u0068e\u0072 \u0074h\u0061\u006e \u0044\u0065fa\u0075\u006c\u0074\u002e"))
				if _dceaa() {
					return true
				}
			}
		}
		if !_edeb && _cbef.Get("\u0048\u0054\u0050") != nil {
			_edeb = true
			_cege = append(_cege, _ba("\u0036.\u0032\u002e\u0035\u002d\u0033", "\u0041\u006e\u0020\u0045\u0078\u0074\u0047\u0053\u0074\u0061\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068\u0061\u006c\u006c \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020th\u0065\u0020\u0048\u0054\u0050\u0020\u006b\u0065\u0079\u002e"))
		}
		_acegdb, _cabbe := _ad.GetDict(_cbef.Get("\u0048\u0054"))
		if _cabbe {
			if _afge := _acegdb.Get("\u0048\u0061\u006cf\u0074\u006f\u006e\u0065\u0054\u0079\u0070\u0065"); !_egbd && _afge != nil {
				_dbbaf, _fagab := _ad.GetInt(_afge)
				if !_fagab || (_fagab && !(*_dbbaf == 1 || *_dbbaf == 5)) {
					_cege = append(_cege, _ba("\u0020\u0036\u002e\u0032\u002e\u0035\u002d\u0034", "\u0041\u006c\u006c\u0020\u0068\u0061\u006c\u0066\u0074\u006f\u006e\u0065\u0073\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006ce\u0020\u0073h\u0061\u006c\u006c\u0020h\u0061\u0076\u0065\u0020\u0074\u0068\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0031\u0020\u006f\u0072\u0020\u0035 \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0048\u0061l\u0066\u0074\u006fn\u0065\u0054\u0079\u0070\u0065\u0020\u006be\u0079\u002e"))
					if _dceaa() {
						return true
					}
				}
			}
			if _cagbg := _acegdb.Get("\u0048\u0061\u006cf\u0074\u006f\u006e\u0065\u004e\u0061\u006d\u0065"); !_ggbg && _cagbg != nil {
				_ggbg = true
				_cege = append(_cege, _ba("\u0036.\u0032\u002e\u0035\u002d\u0035", "\u0048\u0061\u006c\u0066\u0074o\u006e\u0065\u0073\u0020\u0069\u006e\u0020a\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u002f\u0041\u002d\u0032\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061i\u006e\u0020\u0061\u0020\u0048\u0061\u006c\u0066\u0074\u006f\u006e\u0065N\u0061\u006d\u0065\u0020\u006b\u0065y\u002e"))
				if _dceaa() {
					return true
				}
			}
		}
		_, _agcf := _faad(_aacbd)
		var _bgdc bool
		_decf, _cabbe := _ad.GetDict(_cbef.Get("\u0047\u0072\u006fu\u0070"))
		if _cabbe {
			_, _ecbf := _ad.GetName(_decf.Get("\u0043\u0053"))
			if _ecbf {
				_bgdc = true
			}
		}
		if _dfgfc := _cbef.Get("\u0042\u004d"); !_egca && !_bedaf && _dfgfc != nil {
			_debgf, _fbada := _ad.GetName(_dfgfc)
			if _fbada {
				switch _debgf.String() {
				case "\u004e\u006f\u0072\u006d\u0061\u006c", "\u0043\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u006c\u0065", "\u004d\u0075\u006c\u0074\u0069\u0070\u006c\u0079", "\u0053\u0063\u0072\u0065\u0065\u006e", "\u004fv\u0065\u0072\u006c\u0061\u0079", "\u0044\u0061\u0072\u006b\u0065\u006e", "\u004ci\u0067\u0068\u0074\u0065\u006e", "\u0043\u006f\u006c\u006f\u0072\u0044\u006f\u0064\u0067\u0065", "\u0043o\u006c\u006f\u0072\u0042\u0075\u0072n", "\u0048a\u0072\u0064\u004c\u0069\u0067\u0068t", "\u0053o\u0066\u0074\u004c\u0069\u0067\u0068t", "\u0044\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0063\u0065", "\u0045x\u0063\u006c\u0075\u0073\u0069\u006fn", "\u0048\u0075\u0065", "\u0053\u0061\u0074\u0075\u0072\u0061\u0074\u0069\u006f\u006e", "\u0043\u006f\u006co\u0072", "\u004c\u0075\u006d\u0069\u006e\u006f\u0073\u0069\u0074\u0079":
				default:
					_egca = true
					_cege = append(_cege, _ba("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0031", "\u004f\u006el\u0079\u0020\u0062\u006c\u0065\u006e\u0064\u0020\u006d\u006f\u0064\u0065\u0073\u0020\u0074h\u0061\u0074\u0020\u0061\u0072\u0065\u0020\u0073\u0070\u0065c\u0069\u0066\u0069ed\u0020\u0069\u006e\u0020\u0049\u0053O\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031\u003a2\u0030\u0030\u0038\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0075\u0073\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065 \u0076\u0061\u006c\u0075e\u0020\u006f\u0066\u0020\u0074\u0068e\u0020\u0042M\u0020\u006b\u0065\u0079\u0020\u0069\u006e\u0020\u0061\u006e\u0020\u0065\u0078t\u0065\u006e\u0064\u0065\u0064\u0020\u0067\u0072\u0061\u0070\u0068\u0069\u0063\u0020\u0073\u0074\u0061\u0074\u0065 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
					if _dceaa() {
						return true
					}
				}
				if _debgf.String() != "\u004e\u006f\u0072\u006d\u0061\u006c" && !_agcf && !_bgdc {
					_bedaf = true
					_cege = append(_cege, _ba("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
					if _dceaa() {
						return true
					}
				}
			}
		}
		if _, _cabbe = _ad.GetDict(_cbef.Get("\u0053\u004d\u0061s\u006b")); !_bedaf && _cabbe && !_agcf && !_bgdc {
			_bedaf = true
			_cege = append(_cege, _ba("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
			if _dceaa() {
				return true
			}
		}
		if _abfa := _cbef.Get("\u0043\u0041"); !_bedaf && _abfa != nil && !_agcf && !_bgdc {
			_agbc, _bcgbf := _ad.GetNumberAsFloat(_abfa)
			if _bcgbf == nil && _agbc < 1.0 {
				_bedaf = true
				_cege = append(_cege, _ba("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
				if _dceaa() {
					return true
				}
			}
		}
		if _faec := _cbef.Get("\u0063\u0061"); !_bedaf && _faec != nil && !_agcf && !_bgdc {
			_ffab, _aaeec := _ad.GetNumberAsFloat(_faec)
			if _aaeec == nil && _ffab < 1.0 {
				_bedaf = true
				_cege = append(_cege, _ba("\u0036\u002e\u0032\u002e\u0031\u0030\u002d\u0032", "\u0049\u0066\u0020\u0074\u0068\u0065 \u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020P\u0044\u0046\u002f\u0041\u0020\u004f\u0075\u0074\u0070u\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u002c\u0020\u0074\u0068\u0065\u006e\u0020\u0061\u006c\u006c\u0020\u0050\u0061\u0067\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0073\u0020\u0074\u0068a\u0074 \u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020t\u0072\u0061\u006e\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0069\u006e\u0063l\u0075\u0064\u0065\u0020\u0074\u0068\u0065\u0020\u0047\u0072\u006f\u0075\u0070\u0020\u006b\u0065y\u002c a\u006e\u0064\u0020\u0074\u0068\u0065\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u0064\u0069c\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0074\u0068\u0061\u0074\u0020\u0066\u006f\u0072\u006d\u0073\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006cu\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0061\u0074\u0020\u0047\u0072\u006fu\u0070\u0020\u006b\u0065y\u0020sh\u0061\u006c\u006c\u0020\u0069\u006e\u0063\u006c\u0075d\u0065\u0020\u0061\u0020\u0043\u0053\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0077\u0068\u006fs\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065 \u0075\u0073\u0065\u0064\u0020\u0061\u0073\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u0020\u0062\u006c\u0065\u006e\u0064\u0069n\u0067 \u0063\u006f\u006c\u006f\u0075\u0072\u0020\u0073p\u0061\u0063\u0065\u002e"))
				if _dceaa() {
					return true
				}
			}
		}
		return false
	}
	for _, _bbceg := range _aacbd.PageList {
		_eabc := _bbceg.Resources
		if _eabc == nil {
			continue
		}
		if _eabc.ExtGState == nil {
			continue
		}
		_deafe, _ecdgc := _ad.GetDict(_eabc.ExtGState)
		if !_ecdgc {
			continue
		}
		_bdaf := _deafe.Keys()
		for _, _gfecg := range _bdaf {
			_aebb, _gfgdb := _ad.GetDict(_deafe.Get(_gfecg))
			if !_gfgdb {
				continue
			}
			if _ggad(_aebb) {
				return _cege
			}
		}
	}
	for _, _adea := range _aacbd.PageList {
		_afaa := _adea.Resources
		if _afaa == nil {
			continue
		}
		_bdcf, _cagbd := _ad.GetDict(_afaa.XObject)
		if !_cagbd {
			continue
		}
		for _, _bddc := range _bdcf.Keys() {
			_ddde, _edfbg := _ad.GetStream(_bdcf.Get(_bddc))
			if !_edfbg {
				continue
			}
			_efgae, _edfbg := _ad.GetDict(_ddde.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
			if !_edfbg {
				continue
			}
			_gbgeb, _edfbg := _ad.GetDict(_efgae.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
			if !_edfbg {
				continue
			}
			for _, _ffec := range _gbgeb.Keys() {
				_dbfb, _fcdb := _ad.GetDict(_gbgeb.Get(_ffec))
				if !_fcdb {
					continue
				}
				if _ggad(_dbfb) {
					return _cege
				}
			}
		}
	}
	return _cege
}

// Validate checks if provided input document reader matches given PDF/A profile.
func Validate(d *_da.CompliancePdfReader, profile Profile) error { return profile.ValidateStandard(d) }
func _bae(_gacdg *_ea.Document, _dccf int) error {
	_egfa := map[*_ad.PdfObjectStream]struct{}{}
	for _, _acga := range _gacdg.Objects {
		_dbffc, _gfcc := _ad.GetStream(_acga)
		if !_gfcc {
			continue
		}
		if _, _gfcc = _egfa[_dbffc]; _gfcc {
			continue
		}
		_egfa[_dbffc] = struct{}{}
		_ecg, _gfcc := _ad.GetName(_dbffc.Get("\u0053u\u0062\u0054\u0079\u0070\u0065"))
		if !_gfcc {
			continue
		}
		if _dbffc.Get("\u0052\u0065\u0066") != nil {
			_dbffc.Remove("\u0052\u0065\u0066")
		}
		if _ecg.String() == "\u0050\u0053" {
			_dbffc.Remove("\u0050\u0053")
			continue
		}
		if _ecg.String() == "\u0046\u006f\u0072\u006d" {
			if _dbffc.Get("\u004f\u0050\u0049") != nil {
				_dbffc.Remove("\u004f\u0050\u0049")
			}
			if _dbffc.Get("\u0050\u0053") != nil {
				_dbffc.Remove("\u0050\u0053")
			}
			if _abg := _dbffc.Get("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032"); _abg != nil {
				if _acee, _fdcf := _ad.GetName(_abg); _fdcf && *_acee == "\u0050\u0053" {
					_dbffc.Remove("\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0032")
				}
			}
			continue
		}
		if _ecg.String() == "\u0049\u006d\u0061g\u0065" {
			_fcge, _afdb := _ad.GetBool(_dbffc.Get("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065"))
			if _afdb && bool(*_fcge) {
				_dbffc.Set("I\u006e\u0074\u0065\u0072\u0070\u006f\u006c\u0061\u0074\u0065", _ad.MakeBool(false))
			}
			if _dccf == 2 {
				if _dbffc.Get("\u004f\u0050\u0049") != nil {
					_dbffc.Remove("\u004f\u0050\u0049")
				}
			}
			if _dbffc.Get("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073") != nil {
				_dbffc.Remove("\u0041\u006c\u0074\u0065\u0072\u006e\u0061\u0074\u0065\u0073")
			}
			continue
		}
	}
	return nil
}
func _bfba(_gagfa *_da.CompliancePdfReader) ViolatedRule {
	for _, _eabd := range _gagfa.PageList {
		_gdeg, _dbcc := _eabd.GetContentStreams()
		if _dbcc != nil {
			continue
		}
		for _, _fecd := range _gdeg {
			_gbdc := _cad.NewContentStreamParser(_fecd)
			_, _dbcc = _gbdc.Parse()
			if _dbcc != nil {
				return _ba("\u0036.\u0032\u002e\u0032\u002d\u0031", "\u0041\u0020\u0063onten\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c n\u006f\u0074\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079 \u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u0020\u006e\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0065\u0076\u0065\u006e\u0020\u0069\u0066\u0020s\u0075\u0063\u0068\u0020\u006f\u0070\u0065r\u0061\u0074\u006f\u0072\u0073\u0020\u0061\u0072\u0065\u0020\u0062\u0072\u0061\u0063\u006b\u0065\u0074\u0065\u0064\u0020\u0062\u0079\u0020\u0074\u0068\u0065\u0020\u0042\u0058\u002f\u0045\u0058\u0020\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062i\u006c\u0069\u0074\u0079\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072\u0073\u002e")
			}
		}
	}
	return _bg
}
func _dbb(_dag *_ea.Document) {
	if _dag.ID[0] != "" && _dag.ID[1] != "" {
		return
	}
	_dag.UseHashBasedID = true
}

// Conformance gets the PDF/A conformance.
func (_cfce *profile1) Conformance() string { return _cfce._efde._ee }
func _aedfg(_gcad *_da.CompliancePdfReader) (_acge []ViolatedRule) {
	if _gcad.ParserMetadata().HasOddLengthHexStrings() {
		_acge = append(_acge, _ba("\u0036.\u0031\u002e\u0036\u002d\u0031", "\u0068\u0065\u0078a\u0064\u0065\u0063\u0069\u006d\u0061\u006c\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006f\u0066\u0020e\u0076\u0065\u006e\u0020\u0073\u0069\u007a\u0065"))
	}
	if _gcad.ParserMetadata().HasOddLengthHexStrings() {
		_acge = append(_acge, _ba("\u0036.\u0031\u002e\u0036\u002d\u0032", "\u0068\u0065\u0078\u0061\u0064\u0065\u0063\u0069\u006da\u006c\u0020s\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068o\u0075\u006c\u0064\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u006f\u006e\u006c\u0079\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0066\u0072\u006f\u006d\u0020\u0072\u0061n\u0067\u0065\u0020[\u0030\u002d\u0039\u003b\u0041\u002d\u0046\u003b\u0061\u002d\u0066\u005d"))
	}
	return _acge
}
func _fabe(_fddd *_da.CompliancePdfReader) []ViolatedRule { return nil }

// ValidateStandard checks if provided input CompliancePdfReader matches rules that conforms PDF/A-2 standard.
func (_ccbe *profile2) ValidateStandard(r *_da.CompliancePdfReader) error {
	_efce := VerificationError{ConformanceLevel: _ccbe._fcdd._efe, ConformanceVariant: _ccbe._fcdd._ee}
	if _dffg := _fdbf(r); _dffg != _bg {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _dffg)
	}
	if _cbfa := _gdgae(r); _cbfa != _bg {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _cbfa)
	}
	if _fgec := _edcgb(r); _fgec != _bg {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _fgec)
	}
	if _dega := _eeda(r); _dega != _bg {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _dega)
	}
	if _efacg := _bcfde(r); _efacg != _bg {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _efacg)
	}
	if _edcg := _dgccb(r); len(_edcg) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _edcg...)
	}
	if _bba := _cfgd(r); len(_bba) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _bba...)
	}
	if _bcba := _fabe(r); len(_bcba) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _bcba...)
	}
	if _dbce := _geddd(r); _dbce != _bg {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _dbce)
	}
	if _cfceg := _aggbf(r); len(_cfceg) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _cfceg...)
	}
	if _caeg := _eecg(r); len(_caeg) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _caeg...)
	}
	if _bca := _bfba(r); _bca != _bg {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _bca)
	}
	if _cfbba := _adgfe(r); len(_cfbba) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _cfbba...)
	}
	if _afe := _cgcdc(r); len(_afe) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _afe...)
	}
	if _dagcg := _cdac(r); _dagcg != _bg {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _dagcg)
	}
	if _feabg := _ggef(r); len(_feabg) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _feabg...)
	}
	if _caef := _fbed(r); len(_caef) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _caef...)
	}
	if _gccd := _fefc(r); _gccd != _bg {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _gccd)
	}
	if _fbc := _dagb(r); len(_fbc) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _fbc...)
	}
	if _fgaa := _dbfde(r, _ccbe._fcdd); len(_fgaa) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _fgaa...)
	}
	if _feea := _bfaae(r); len(_feea) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _feea...)
	}
	if _bgfa := _ebaec(r); len(_bgfa) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _bgfa...)
	}
	if _gddbb := _bfdfa(r); len(_gddbb) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _gddbb...)
	}
	if _edf := _daae(r); _edf != _bg {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _edf)
	}
	if _bgeef := _gcgef(r); len(_bgeef) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _bgeef...)
	}
	if _cdgbe := _cafe(r); _cdgbe != _bg {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _cdgbe)
	}
	if _fgdd := _badd(r, _ccbe._fcdd, false); len(_fgdd) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _fgdd...)
	}
	if _ccbe._fcdd == _bb() {
		if _efff := _dfgdgd(r); len(_efff) != 0 {
			_efce.ViolatedRules = append(_efce.ViolatedRules, _efff...)
		}
	}
	if _dcba := _acegf(r); len(_dcba) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _dcba...)
	}
	if _fggdc := _dddf(r); len(_fggdc) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _fggdc...)
	}
	if _fdad := _ceeeb(r); len(_fdad) != 0 {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _fdad...)
	}
	if _bfcf := _ebbag(r); _bfcf != _bg {
		_efce.ViolatedRules = append(_efce.ViolatedRules, _bfcf)
	}
	if len(_efce.ViolatedRules) > 0 {
		_cg.Slice(_efce.ViolatedRules, func(_ebae, _bgdb int) bool {
			return _efce.ViolatedRules[_ebae].RuleNo < _efce.ViolatedRules[_bgdb].RuleNo
		})
		return _efce
	}
	return nil
}
func _dgcea(_fffbf *_da.PdfFont, _agcgb *_ad.PdfObjectDictionary) ViolatedRule {
	const (
		_begb = "\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0036\u002d\u0033"
		_gcgg = "\u0041l\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0069\u0063\u0020\u0054\u0072u\u0065\u0054\u0079p\u0065\u0020\u0066\u006f\u006e\u0074s\u0020\u0073h\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0020\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065n\u0074\u0072\u0079\u0020\u0069n\u0020\u0074\u0068e\u0020\u0066\u006f\u006e\u0074 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e"
	)
	var _cabfgb string
	if _dbdge, _cdad := _ad.GetName(_agcgb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _cdad {
		_cabfgb = _dbdge.String()
	}
	if _cabfgb != "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065" {
		return _bg
	}
	_bbcb := _fffbf.FontDescriptor()
	_fegec, _ecdcg := _ad.GetIntVal(_bbcb.Flags)
	if !_ecdcg {
		_ca.Log.Debug("\u0066\u006c\u0061\u0067\u0073 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u006f\u0072")
		return _ba(_begb, _gcgg)
	}
	_ccba := (uint32(_fegec) >> 3) & 1
	_fcebb := _ccba != 0
	if !_fcebb {
		return _bg
	}
	if _agcgb.Get("\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067") != nil {
		return _ba(_begb, _gcgg)
	}
	return _bg
}
func _bbac(_dbfcf *_da.CompliancePdfReader) (_gffc []ViolatedRule) {
	var _ecdb, _ggdbb, _eaef, _cgd, _gdbg, _geef, _abf bool
	_eagd := func() bool { return _ecdb && _ggdbb && _eaef && _cgd && _gdbg && _geef && _abf }
	for _, _ebdf := range _dbfcf.PageList {
		if _ebdf.Resources == nil {
			continue
		}
		_edgcd, _fcbe := _ad.GetDict(_ebdf.Resources.Font)
		if !_fcbe {
			continue
		}
		for _, _beada := range _edgcd.Keys() {
			_edce, _fede := _ad.GetDict(_edgcd.Get(_beada))
			if !_fede {
				if !_ecdb {
					_gffc = append(_gffc, _ba("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0041\u006c\u006c\u0020\u0066\u006fn\u0074\u0073\u0020\u0075\u0073e\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020c\u006f\u006e\u0066\u006f\u0072m\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0073\u0020d\u0065\u0066\u0069\u006e\u0065d \u0069\u006e\u0020\u0050\u0044\u0046\u0020\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035\u002e\u0035\u002e"))
					_ecdb = true
					if _eagd() {
						return _gffc
					}
				}
				continue
			}
			if _ddce, _cbee := _ad.GetName(_edce.Get("\u0054\u0079\u0070\u0065")); !_ecdb && (!_cbee || _ddce.String() != "\u0046\u006f\u006e\u0074") {
				_gffc = append(_gffc, _ba("\u0036.\u0033\u002e\u0032\u002d\u0031", "\u0054\u0079\u0070e\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029 Th\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066 \u0050\u0044\u0046\u0020\u006fbj\u0065\u0063\u0074\u0020\u0074\u0068\u0061t\u0020\u0074\u0068\u0069s\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0064\u0065\u0073c\u0072\u0069\u0062\u0065\u0073\u003b\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0046\u006f\u006e\u0074\u0020\u0066\u006fr\u0020\u0061\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
				_ecdb = true
				if _eagd() {
					return _gffc
				}
			}
			_dgbb, _feaa := _da.NewPdfFontFromPdfObject(_edce)
			if _feaa != nil {
				continue
			}
			var _dagg string
			if _cead, _ecaf := _ad.GetName(_edce.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _ecaf {
				_dagg = _cead.String()
			}
			if !_ggdbb {
				switch _dagg {
				case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
				default:
					_ggdbb = true
					_gffc = append(_gffc, _ba("\u0036.\u0033\u002e\u0032\u002d\u0032", "\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065d\u0029\u0020\u0054\u0068e \u0074\u0079\u0070\u0065 \u006f\u0066\u0020\u0066\u006f\u006et\u003b\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u0022\u0054\u0079\u0070\u0065\u0031\u0022\u0020f\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020f\u006f\u006e\u0074\u0073\u002c\u0020\u0022\u004d\u004d\u0054\u0079\u0070\u0065\u0031\u0022\u0020\u0066\u006f\u0072\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006da\u0073\u0074e\u0072\u0020\u0066\u006f\u006e\u0074s\u002c\u0020\u0022\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0022\u0054\u0079\u0070\u0065\u0033\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070e\u0020\u0033\u0020\u0066\u006f\u006e\u0074\u0073\u002c\u0020\"\u0054\u0079\u0070\u0065\u0030\"\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006ed\u0020\u0022\u0043\u0049\u0044\u0046\u006fn\u0074\u0054\u0079\u0070\u0065\u0030\u0022 \u006f\u0072\u0020\u0022\u0043\u0049\u0044\u0046\u006f\u006e\u0074T\u0079\u0070e\u0032\u0022\u0020\u0066\u006f\u0072\u0020\u0043\u0049\u0044\u0020\u0066\u006f\u006e\u0074\u0073\u002e"))
					if _eagd() {
						return _gffc
					}
				}
			}
			if !_eaef {
				if _dagg != "\u0054\u0079\u0070e\u0033" {
					_faae, _cdbf := _ad.GetName(_edce.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
					if !_cdbf || _faae.String() == "" {
						_gffc = append(_gffc, _ba("\u0036.\u0033\u002e\u0032\u002d\u0033", "B\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064)\u0020T\u0068\u0065\u0020\u0050o\u0073\u0074S\u0063\u0072\u0069\u0070\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e"))
						_eaef = true
						if _eagd() {
							return _gffc
						}
					}
				}
			}
			if _dagg != "\u0054\u0079\u0070e\u0031" {
				continue
			}
			_ddff := _cac.IsStdFont(_cac.StdFontName(_dgbb.BaseFont()))
			if _ddff {
				continue
			}
			_abea, _ggbd := _ad.GetIntVal(_edce.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
			if !_ggbd && !_cgd {
				_gffc = append(_gffc, _ba("\u0036.\u0033\u002e\u0032\u002d\u0034", "\u0046\u0069r\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072d\u0020\u0031\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u0029\u0020\u0054\u0068\u0065\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064e\u0020\u0064\u0065\u0066i\u006ee\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057i\u0064\u0074\u0068\u0073 \u0061r\u0072\u0061y\u002e"))
				_cgd = true
				if _eagd() {
					return _gffc
				}
			}
			_eaddf, _ccaca := _ad.GetIntVal(_edce.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
			if !_ccaca && !_gdbg {
				_gffc = append(_gffc, _ba("\u0036.\u0033\u002e\u0032\u002d\u0035", "\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069n\u0074\u0065\u0067e\u0072 \u002d\u0020\u0028\u0052\u0065\u0071u\u0069\u0072\u0065d\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020t\u0068\u0065 s\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0031\u0034\u0020\u0066\u006f\u006ets\u0029\u0020\u0054\u0068\u0065\u0020\u006c\u0061\u0073t\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057\u0069\u0064\u0074h\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u002e"))
				_gdbg = true
				if _eagd() {
					return _gffc
				}
			}
			if !_geef {
				_ccag, _fcbed := _ad.GetArray(_edce.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
				if !_fcbed || !_ggbd || !_ccaca || _ccag.Len() != _eaddf-_abea+1 {
					_gffc = append(_gffc, _ba("\u0036.\u0033\u002e\u0032\u002d\u0036", "\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u002d a\u0072\u0072\u0061y \u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0074a\u006e\u0064a\u0072\u0064\u00201\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u003b\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0070\u0072\u0065\u0066e\u0072\u0072e\u0064\u0029\u0020\u0041\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0066\u0020\u0028\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u2212 F\u0069\u0072\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u002b\u00201\u0029\u0020\u0077\u0069\u0064\u0074\u0068\u0073."))
					_geef = true
					if _eagd() {
						return _gffc
					}
				}
			}
		}
	}
	return _gffc
}
func _eecg(_ffdab *_da.CompliancePdfReader) (_fccf []ViolatedRule) {
	var (
		_fddda, _agccg, _fabee, _adabb, _aeec bool
		_fceb                                 func(_ad.PdfObject)
	)
	_fceb = func(_gcdb _ad.PdfObject) {
		switch _dcegd := _gcdb.(type) {
		case *_ad.PdfObjectInteger:
			if !_fddda && (int64(*_dcegd) > _e.MaxInt32 || int64(*_dcegd) < -_e.MaxInt32) {
				_fccf = append(_fccf, _ba("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0031", "L\u0061\u0072\u0067e\u0073\u0074\u0020\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u0032\u002c\u0031\u0034\u0037,\u0034\u0038\u0033,\u0036\u0034\u0037\u002e\u0020\u0053\u006d\u0061\u006c\u006c\u0065\u0073\u0074 \u0069\u006e\u0074\u0065g\u0065\u0072\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020\u002d\u0032\u002c\u0031\u0034\u0037\u002c\u0034\u0038\u0033,\u0036\u0034\u0038\u002e"))
				_fddda = true
			}
		case *_ad.PdfObjectFloat:
			if !_agccg && (_e.Abs(float64(*_dcegd)) > _e.MaxFloat32) {
				_fccf = append(_fccf, _ba("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0032", "\u0041 \u0063\u006f\u006e\u0066orm\u0069\u006e\u0067\u0020f\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020n\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u006e\u0079\u0020\u0072\u0065\u0061\u006c\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u006f\u0075\u0074\u0073\u0069de\u0020\u0074\u0068e\u0020\u0072\u0061\u006e\u0067e\u0020o\u0066\u0020\u002b\u002f\u002d\u0033\u002e\u0034\u00303\u0020\u0078\u0020\u0031\u0030\u005e\u0033\u0038\u002e"))
			}
		case *_ad.PdfObjectString:
			if !_fabee && len([]byte(_dcegd.Str())) > 32767 {
				_fccf = append(_fccf, _ba("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0033", "M\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006c\u0065n\u0067\u0074\u0068\u0020\u006f\u0066\u0020a \u0073\u0074\u0072\u0069n\u0067\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074es\u0029\u0020i\u0073\u0020\u0033\u0032\u0037\u0036\u0037\u002e"))
				_fabee = true
			}
		case *_ad.PdfObjectName:
			if !_adabb && len([]byte(*_dcegd)) > 127 {
				_fccf = append(_fccf, _ba("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0034", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006c\u0065\u006eg\u0074\u0068\u0020\u006ff\u0020\u0061\u0020\u006e\u0061\u006d\u0065\u0020\u0028\u0069\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0029\u0020\u0069\u0073\u0020\u0031\u0032\u0037\u002e"))
				_adabb = true
			}
		case *_ad.PdfObjectArray:
			for _, _gbdg := range _dcegd.Elements() {
				_fceb(_gbdg)
			}
			if !_aeec && (_dcegd.Len() == 4 || _dcegd.Len() == 5) {
				_cabfg, _fdbd := _ad.GetName(_dcegd.Get(0))
				if !_fdbd {
					return
				}
				if *_cabfg != "\u0044e\u0076\u0069\u0063\u0065\u004e" {
					return
				}
				_ddfee := _dcegd.Get(1)
				_ddfee = _ad.TraceToDirectObject(_ddfee)
				_cbgg, _fdbd := _ad.GetArray(_ddfee)
				if !_fdbd {
					return
				}
				if _cbgg.Len() > 32 {
					_fccf = append(_fccf, _ba("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0039", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d \u006e\u0075\u006db\u0065\u0072\u0020\u006ff\u0020\u0044\u0065\u0076\u0069\u0063\u0065\u004e\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0069\u0073\u0020\u0033\u0032\u002e"))
					_aeec = true
				}
			}
		case *_ad.PdfObjectDictionary:
			_dcffa := _dcegd.Keys()
			for _fecg, _bebb := range _dcffa {
				_fceb(&_dcffa[_fecg])
				_fceb(_dcegd.Get(_bebb))
			}
		case *_ad.PdfObjectStream:
			_fceb(_dcegd.PdfObjectDictionary)
		case *_ad.PdfObjectStreams:
			for _, _gfcf := range _dcegd.Elements() {
				_fceb(_gfcf)
			}
		case *_ad.PdfObjectReference:
			_fceb(_dcegd.Resolve())
		}
	}
	_egfc := _ffdab.GetObjectNums()
	if len(_egfc) > 8388607 {
		_fccf = append(_fccf, _ba("\u0036\u002e\u0031\u002e\u0031\u0033\u002d\u0037", "\u004d\u0061\u0078\u0069\u006d\u0075\u006d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020in\u0064i\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0073 \u0069\u006e\u0020\u0061\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0069\u0073\u00208\u002c\u0033\u0038\u0038\u002c\u0036\u0030\u0037\u002e"))
	}
	for _, _ffcc := range _egfc {
		_cgbd, _caag := _ffdab.GetIndirectObjectByNumber(_ffcc)
		if _caag != nil {
			continue
		}
		_adef := _ad.TraceToDirectObject(_cgbd)
		_fceb(_adef)
	}
	return _fccf
}

// NewProfile2A creates a new Profile2A with given options.
func NewProfile2A(options *Profile2Options) *Profile2A {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_cdcc(options)
	return &Profile2A{profile2{_eceb: *options, _fcdd: _bb()}}
}
func _dgdb(_dfgdf *_da.CompliancePdfReader) (_dfde ViolatedRule) {
	_cbcb, _eegb := _acgab(_dfgdf)
	if !_eegb {
		return _bg
	}
	_bgbc, _eegb := _ad.GetDict(_cbcb.Get("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"))
	if !_eegb {
		return _bg
	}
	_acae, _eegb := _ad.GetArray(_bgbc.Get("\u0046\u0069\u0065\u006c\u0064\u0073"))
	if !_eegb {
		return _bg
	}
	for _cgegc := 0; _cgegc < _acae.Len(); _cgegc++ {
		_efcgg, _dcddf := _ad.GetDict(_acae.Get(_cgegc))
		if !_dcddf {
			continue
		}
		if _efcgg.Get("\u0041\u0041") != nil {
			return _ba("\u0036.\u0036\u002e\u0032\u002d\u0032", "\u0041\u0020F\u0069\u0065\u006cd\u0020\u0064\u0069\u0063\u0074i\u006f\u006e\u0061\u0072\u0079 s\u0068\u0061\u006c\u006c\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0063\u006c\u0075\u0064\u0065\u0020\u0061n\u0020A\u0041\u0020\u0065\u006e\u0074\u0072y f\u006f\u0072\u0020\u0061\u006e\u0020\u0061\u0064\u0064\u0069\u0074\u0069on\u0061l\u002d\u0061\u0063\u0074i\u006fn\u0073 \u0064\u0069c\u0074\u0069on\u0061\u0072\u0079\u002e")
		}
	}
	return _bg
}
func _dccb(_fbe *_ea.Document) error {
	_dgab, _eba := _fbe.GetPages()
	if !_eba {
		return nil
	}
	for _, _ecfa := range _dgab {
		_gabd := _ecfa.FindXObjectForms()
		for _, _dffbc := range _gabd {
			_ebcf, _gfd := _ad.GetDict(_dffbc.Get("\u0047\u0072\u006fu\u0070"))
			if _gfd {
				if _cba := _ebcf.Get("\u0053"); _cba != nil {
					_bbfb, _aggb := _ad.GetName(_cba)
					if _aggb && _bbfb.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						_dffbc.Remove("\u0047\u0072\u006fu\u0070")
					}
				}
			}
		}
		_dgac, _cbd := _ecfa.GetResourcesXObject()
		if _cbd {
			_bdae, _dgd := _ad.GetDict(_dgac.Get("\u0047\u0072\u006fu\u0070"))
			if _dgd {
				_adb := _bdae.Get("\u0053")
				if _adb != nil {
					_ffeg, _caca := _ad.GetName(_adb)
					if _caca && _ffeg.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
						_dgac.Remove("\u0047\u0072\u006fu\u0070")
					}
				}
			}
		}
		_fgdf, _cdgb := _ad.GetDict(_ecfa.Object.Get("\u0047\u0072\u006fu\u0070"))
		if _cdgb {
			_gaeb := _fgdf.Get("\u0053")
			if _gaeb != nil {
				_dacc, _ffegd := _ad.GetName(_gaeb)
				if _ffegd && _dacc.String() == "\u0054\u0072\u0061n\u0073\u0070\u0061\u0072\u0065\u006e\u0063\u0079" {
					_ecfa.Object.Remove("\u0047\u0072\u006fu\u0070")
				}
			}
		}
	}
	return nil
}
func _bfdfa(_cgeee *_da.CompliancePdfReader) (_efgab []ViolatedRule) { return _efgab }

// Profile2U is the implementation of the PDF/A-2U standard profile.
// Implements model.StandardImplementer, Profile interfaces.
type Profile2U struct{ profile2 }

func _bebg(_dcab *_da.CompliancePdfReader) (_fdb []ViolatedRule) {
	var _fbde, _fadg, _fegg bool
	if _dcab.ParserMetadata().HasNonConformantStream() {
		_fdb = []ViolatedRule{_ba("\u0036.\u0031\u002e\u0037\u002d\u0031", "T\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062\u0065\u0020f\u006f\u006cl\u006fw\u0065\u0064\u0020e\u0069\u0074h\u0065\u0072\u0020\u0062\u0079\u0020\u0061 \u0043\u0041\u0052\u0052I\u0041\u0047\u0045\u0020\u0052E\u0054\u0055\u0052\u004e\u0020\u00280\u0044\u0068\u0029\u0020\u0061\u006e\u0064\u0020\u004c\u0049\u004e\u0045\u0020F\u0045\u0045\u0044\u0020\u0028\u0030\u0041\u0068\u0029\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065\u0020o\u0072\u0020\u0062\u0079\u0020\u0061 \u0073\u0069ng\u006c\u0065\u0020\u004cIN\u0045 \u0046\u0045\u0045\u0044 \u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u002e\u0020T\u0068\u0065\u0020e\u006e\u0064\u0073\u0074r\u0065\u0061\u006d\u0020\u006b\u0065\u0079\u0077\u006fr\u0064\u0020\u0073\u0068\u0061\u006c\u006c \u0062e\u0020p\u0072\u0065\u0063\u0065\u0064\u0065\u0064\u0020\u0062\u0079\u0020\u0061n\u0020\u0045\u004f\u004c \u006d\u0061\u0072\u006b\u0065\u0072\u002e")}
	}
	for _, _dbag := range _dcab.GetObjectNums() {
		_gaad, _ := _dcab.GetIndirectObjectByNumber(_dbag)
		if _gaad == nil {
			continue
		}
		_eaddc, _dbaf := _ad.GetStream(_gaad)
		if !_dbaf {
			continue
		}
		if !_fbde {
			_bac := _eaddc.Get("\u004c\u0065\u006e\u0067\u0074\u0068")
			if _bac == nil {
				_fdb = append(_fdb, _ba("\u0036.\u0031\u002e\u0037\u002d\u0032", "\u006e\u006f\u0020'\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0073\u0074\u0072\u0065a\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074"))
				_fbde = true
			} else {
				_aggee, _dffe := _ad.GetIntVal(_bac)
				if !_dffe {
					_fdb = append(_fdb, _ba("\u0036.\u0031\u002e\u0037\u002d\u0032", "s\u0074\u0072\u0065\u0061\u006d\u0020\u0027\u004c\u0065\u006e\u0067\u0074\u0068\u0027\u0020\u006b\u0065\u0079 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020an\u0020\u0069\u006et\u0065g\u0065\u0072"))
					_fbde = true
				} else {
					if len(_eaddc.Stream) != _aggee {
						_fdb = append(_fdb, _ba("\u0036.\u0031\u002e\u0037\u002d\u0032", "\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006c\u0065\u006e\u0067th\u0020v\u0061\u006c\u0075\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u0068\u0065\u0020\u0073\u0069\u007a\u0065\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d"))
						_fbde = true
					}
				}
			}
		}
		if !_fadg {
			if _eaddc.Get("\u0046") != nil {
				_fadg = true
				_fdb = append(_fdb, _ba("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
			}
			if _eaddc.Get("\u0046F\u0069\u006c\u0074\u0065\u0072") != nil && !_fadg {
				_fadg = true
				_fdb = append(_fdb, _ba("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
			if _eaddc.Get("\u0046\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0072\u0061\u006d\u0073") != nil && !_fadg {
				_fadg = true
				_fdb = append(_fdb, _ba("\u0036.\u0031\u002e\u0037\u002d\u0033", "\u0073\u0074r\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006eo\u0074\u0020\u0063\u006f\u006e\u0074a\u0069\u006e\u0020\u0027\u0046\u0027\u002c\u0027\u0046\u0046\u0069\u006c\u0074\u0065r\u0027\u002c'\u0046\u0044\u0065\u0063o\u0064\u0065\u0050\u0061\u0072a\u006d\u0073\u0027\u0020\u006b\u0065\u0079"))
				continue
			}
		}
		if !_fegg {
			_ecea, _bfaa := _ad.GetName(_ad.TraceToDirectObject(_eaddc.Get("\u0046\u0069\u006c\u0074\u0065\u0072")))
			if !_bfaa {
				continue
			}
			if *_ecea == _ad.StreamEncodingFilterNameLZW {
				_fegg = true
				_fdb = append(_fdb, _ba("\u0036\u002e\u0031\u002e\u0031\u0030\u002d\u0031", "\u0054h\u0065\u0020L\u005a\u0057\u0044\u0065c\u006f\u0064\u0065 \u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0073\u0068al\u006c\u0020\u006eo\u0074\u0020b\u0065\u0020\u0070\u0065\u0072\u006di\u0074\u0074e\u0064\u002e"))
			}
		}
	}
	return _fdb
}
func _fefe(_ggbb *_ad.PdfObjectDictionary, _eacf map[*_ad.PdfObjectStream][]byte, _gdbe map[*_ad.PdfObjectStream]*_b.CMap) ViolatedRule {
	const (
		_dfbfc = "\u0036.\u0033\u002e\u0038\u002d\u0031"
		_dabee = "\u0054\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006ea\u0072\u0079\u0020\u0073\u0068\u0061\u006cl\u0020\u0069\u006e\u0063l\u0075\u0064e\u0020\u0061 \u0054\u006f\u0055\u006e\u0069\u0063\u006f\u0064\u0065\u0020\u0065\u006e\u0074\u0072\u0079\u0020w\u0068\u006f\u0073\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u0073 \u0061\u0020\u0043M\u0061\u0070\u0020\u0073\u0074\u0072\u0065\u0061\u006d \u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0074\u0068\u0061\u0074\u0020\u006d\u0061p\u0073\u0020\u0063\u0068\u0061\u0072ac\u0074\u0065\u0072\u0020\u0063\u006fd\u0065s\u0020\u0074\u006f\u0020\u0055\u006e\u0069\u0063\u006f\u0064e \u0076a\u006c\u0075\u0065\u0073,\u0020\u0061\u0073\u0020\u0064\u0065\u0073\u0063r\u0069\u0062\u0065\u0064\u0020\u0069\u006e\u0020P\u0044\u0046\u0020\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0035.\u0039\u002c\u0020\u0075\u006e\u006ce\u0073\u0073\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0020\u006d\u0065\u0065\u0074\u0073 \u0061\u006e\u0079\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0063\u006f\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u003a\u000a\u0020\u002d\u0020\u0066o\u006e\u0074\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0065\u006e\u0063\u006f\u0064\u0069n\u0067\u0073\u0020M\u0061\u0063\u0052o\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u004d\u0061\u0063\u0045\u0078\u0070\u0065\u0072\u0074E\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u006f\u0072\u0020\u0057\u0069\u006e\u0041n\u0073\u0069\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067\u002c\u0020\u006f\u0072\u0020\u0074\u0068\u0061\u0074\u0020\u0075\u0073\u0065\u0020t\u0068\u0065\u0020\u0070\u0072\u0065d\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079\u002d\u0048\u0020\u006f\u0072\u0020\u0049\u0064\u0065n\u0074\u0069\u0074\u0079\u002d\u0056\u0020C\u004d\u0061\u0070s\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0077\u0068\u006f\u0073\u0065\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0073\u0020a\u0072\u0065 \u0074\u0061k\u0065\u006e\u0020\u0066\u0072\u006f\u006d\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u0020\u0073\u0074\u0061n\u0064\u0061\u0072\u0064\u0020L\u0061t\u0069\u006e\u0020\u0063\u0068a\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0073\u0065\u0074\u0020\u006fr\u0020\u0074\u0068\u0065 \u0073\u0065\u0074\u0020\u006f\u0066 \u006e\u0061\u006d\u0065\u0064\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0066\u006f\u006e\u0074\u002c\u0020\u0061\u0073\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020i\u006e\u0020\u0050\u0044\u0046 \u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0041\u0070\u0070\u0065\u006e\u0064\u0069\u0078 \u0044\u003b\u000a\u0020\u002d\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020w\u0068\u006f\u0073e\u0020d\u0065\u0073\u0063\u0065n\u0064\u0061\u006e\u0074 \u0043\u0049\u0044\u0046\u006f\u006e\u0074\u0020\u0075\u0073\u0065\u0073\u0020\u0074\u0068\u0065\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u0047B\u0031\u002c\u0020\u0041\u0064\u006fb\u0065\u002d\u0043\u004e\u0053\u0031\u002c\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004a\u0061\u0070\u0061\u006e\u0031\u0020\u006f\u0072\u0020\u0041\u0064\u006f\u0062\u0065\u002d\u004b\u006fr\u0065\u0061\u0031\u0020\u0063\u0068\u0061r\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u006c\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0073\u002e"
	)
	_dede, _ggfb := _ad.GetStream(_ggbb.Get("\u0054o\u0055\u006e\u0069\u0063\u006f\u0064e"))
	if _ggfb {
		_, _acgef := _ggag(_dede, _eacf, _gdbe)
		if _acgef != nil {
			return _ba(_dfbfc, _dabee)
		}
		return _bg
	}
	_dgaf, _ggfb := _ad.GetName(_ggbb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065"))
	if !_ggfb {
		return _ba(_dfbfc, _dabee)
	}
	switch _dgaf.String() {
	case "\u0054\u0079\u0070e\u0031":
		return _bg
	}
	return _ba(_dfbfc, _dabee)
}
func _dagb(_cdffd *_da.CompliancePdfReader) (_aebf []ViolatedRule) {
	var _gaddb, _gafbf, _gafa, _eecc, _eadg, _dgaed bool
	_ebaag := func() bool { return _gaddb && _gafbf && _gafa && _eecc && _eadg && _dgaed }
	for _, _fccfb := range _cdffd.PageList {
		if _fccfb.Resources == nil {
			continue
		}
		_gaedg, _caeeb := _ad.GetDict(_fccfb.Resources.Font)
		if !_caeeb {
			continue
		}
		for _, _eegd := range _gaedg.Keys() {
			_caefe, _ceabe := _ad.GetDict(_gaedg.Get(_eegd))
			if !_ceabe {
				if !_gaddb {
					_aebf = append(_aebf, _ba("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0031", "\u0041\u006c\u006c\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006e\u0064\u0020\u0066on\u0074 \u0070\u0072\u006fg\u0072\u0061\u006ds\u0020\u0075\u0073\u0065\u0064\u0020\u0069\u006e\u0020\u0061\u0020\u0063\u006f\u006e\u0066\u006f\u0072mi\u006e\u0067\u0020\u0066\u0069\u006ce\u002c\u0020\u0072\u0065\u0067\u0061\u0072\u0064\u006c\u0065s\u0073\u0020\u006f\u0066\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0069\u006eg m\u006f\u0064\u0065\u0020\u0075\u0073\u0061\u0067\u0065\u002c\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0074\u0068e\u0020\u0070\u0072o\u0076\u0069\u0073\u0069\u006f\u006e\u0073\u0020\u0069\u006e \u0049\u0053\u004f\u0020\u0033\u0032\u0030\u0030\u0030\u002d\u0031:\u0032\u0030\u0030\u0038\u002c \u0039\u002e\u0036\u0020a\u006e\u0064\u0020\u0039.\u0037\u002e"))
					_gaddb = true
					if _ebaag() {
						return _aebf
					}
				}
				continue
			}
			if _cafb, _fegee := _ad.GetName(_caefe.Get("\u0054\u0079\u0070\u0065")); !_gaddb && (!_fegee || _cafb.String() != "\u0046\u006f\u006e\u0074") {
				_aebf = append(_aebf, _ba("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0031", "\u0054\u0079\u0070e\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0029 Th\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u006f\u0066 \u0050\u0044\u0046\u0020\u006fbj\u0065\u0063\u0074\u0020\u0074\u0068\u0061t\u0020\u0074\u0068\u0069s\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0064\u0065\u0073c\u0072\u0069\u0062\u0065\u0073\u003b\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0046\u006f\u006e\u0074\u0020\u0066\u006fr\u0020\u0061\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e"))
				_gaddb = true
				if _ebaag() {
					return _aebf
				}
			}
			_eecac, _eedb := _da.NewPdfFontFromPdfObject(_caefe)
			if _eedb != nil {
				continue
			}
			var _befaa string
			if _abead, _dacce := _ad.GetName(_caefe.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); _dacce {
				_befaa = _abead.String()
			}
			if !_gafbf {
				switch _befaa {
				case "\u0054\u0079\u0070e\u0030", "\u0054\u0079\u0070e\u0031", "\u004dM\u0054\u0079\u0070\u0065\u0031", "\u0054\u0072\u0075\u0065\u0054\u0079\u0070\u0065", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0030", "\u0043\u0049\u0044F\u006f\u006e\u0074\u0054\u0079\u0070\u0065\u0032":
				default:
					_gafbf = true
					_aebf = append(_aebf, _ba("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0032", "\u0053\u0075\u0062\u0074\u0079\u0070\u0065\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065d\u0029\u0020\u0054\u0068e \u0074\u0079\u0070\u0065 \u006f\u0066\u0020\u0066\u006f\u006et\u003b\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u0022\u0054\u0079\u0070\u0065\u0031\u0022\u0020f\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0031\u0020f\u006f\u006e\u0074\u0073\u002c\u0020\u0022\u004d\u004d\u0054\u0079\u0070\u0065\u0031\u0022\u0020\u0066\u006f\u0072\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006da\u0073\u0074e\u0072\u0020\u0066\u006f\u006e\u0074s\u002c\u0020\u0022\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0072\u0075\u0065T\u0079\u0070\u0065\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0022\u0054\u0079\u0070\u0065\u0033\u0022\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070e\u0020\u0033\u0020\u0066\u006f\u006e\u0074\u0073\u002c\u0020\"\u0054\u0079\u0070\u0065\u0030\"\u0020\u0066\u006f\u0072\u0020\u0054\u0079\u0070\u0065\u0020\u0030\u0020\u0066\u006f\u006e\u0074\u0073\u0020\u0061\u006ed\u0020\u0022\u0043\u0049\u0044\u0046\u006fn\u0074\u0054\u0079\u0070\u0065\u0030\u0022 \u006f\u0072\u0020\u0022\u0043\u0049\u0044\u0046\u006f\u006e\u0074T\u0079\u0070e\u0032\u0022\u0020\u0066\u006f\u0072\u0020\u0043\u0049\u0044\u0020\u0066\u006f\u006e\u0074\u0073\u002e"))
					if _ebaag() {
						return _aebf
					}
				}
			}
			if !_gafa {
				if _befaa != "\u0054\u0079\u0070e\u0033" {
					_bbffc, _bacd := _ad.GetName(_caefe.Get("\u0042\u0061\u0073\u0065\u0046\u006f\u006e\u0074"))
					if !_bacd || _bbffc.String() == "" {
						_aebf = append(_aebf, _ba("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0033", "B\u0061\u0073\u0065\u0046\u006f\u006e\u0074\u0020\u002d\u0020\u006e\u0061\u006d\u0065\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064)\u0020T\u0068\u0065\u0020\u0050o\u0073\u0074S\u0063\u0072\u0069\u0070\u0074\u0020\u006e\u0061\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u002e"))
						_gafa = true
						if _ebaag() {
							return _aebf
						}
					}
				}
			}
			if _befaa != "\u0054\u0079\u0070e\u0031" {
				continue
			}
			_gecd := _cac.IsStdFont(_cac.StdFontName(_eecac.BaseFont()))
			if _gecd {
				continue
			}
			_ddeaa, _beae := _ad.GetIntVal(_caefe.Get("\u0046i\u0072\u0073\u0074\u0043\u0068\u0061r"))
			if !_beae && !_eecc {
				_aebf = append(_aebf, _ba("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0034", "\u0046\u0069r\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u002d\u0020\u0028\u0052\u0065\u0071\u0075i\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074h\u0065\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u0072d\u0020\u0031\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u0029\u0020\u0054\u0068\u0065\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064e\u0020\u0064\u0065\u0066i\u006ee\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057i\u0064\u0074\u0068\u0073 \u0061r\u0072\u0061y\u002e"))
				_eecc = true
				if _ebaag() {
					return _aebf
				}
			}
			_efbec, _egbde := _ad.GetIntVal(_caefe.Get("\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072"))
			if !_egbde && !_eadg {
				_aebf = append(_aebf, _ba("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0035", "\u004c\u0061\u0073t\u0043\u0068\u0061\u0072\u0020\u002d\u0020\u0069n\u0074\u0065\u0067e\u0072 \u002d\u0020\u0028\u0052\u0065\u0071u\u0069\u0072\u0065d\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0066\u006f\u0072\u0020t\u0068\u0065 s\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0020\u0031\u0034\u0020\u0066\u006f\u006ets\u0029\u0020\u0054\u0068\u0065\u0020\u006c\u0061\u0073t\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u006e\u0074\u0027\u0073\u0020\u0057\u0069\u0064\u0074h\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u002e"))
				_eadg = true
				if _ebaag() {
					return _aebf
				}
			}
			if !_dgaed {
				_cgef, _eddab := _ad.GetArray(_caefe.Get("\u0057\u0069\u0064\u0074\u0068\u0073"))
				if !_eddab || !_beae || !_egbde || _cgef.Len() != _efbec-_ddeaa+1 {
					_aebf = append(_aebf, _ba("\u0036\u002e\u0032\u002e\u0031\u0031\u002e\u0032\u002d\u0036", "\u0057\u0069\u0064\u0074\u0068\u0073\u0020\u002d a\u0072\u0072\u0061y \u002d\u0020\u0028\u0052\u0065\u0071\u0075\u0069\u0072\u0065\u0064\u0020\u0065\u0078\u0063\u0065\u0070t\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0073\u0074a\u006e\u0064a\u0072\u0064\u00201\u0034\u0020\u0066\u006f\u006e\u0074\u0073\u003b\u0020\u0069\u006ed\u0069\u0072\u0065\u0063\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0070\u0072\u0065\u0066e\u0072\u0072e\u0064\u0029\u0020\u0041\u006e \u0061\u0072\u0072\u0061\u0079\u0020\u006f\u0066\u0020\u0028\u004c\u0061\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u2212 F\u0069\u0072\u0073\u0074\u0043\u0068\u0061\u0072\u0020\u002b\u00201\u0029\u0020\u0077\u0069\u0064\u0074\u0068\u0073."))
					_dgaed = true
					if _ebaag() {
						return _aebf
					}
				}
			}
		}
	}
	return _aebf
}

type pageColorspaceOptimizeFunc func(_dagc *_ea.Document, _fcfa *_ea.Page, _caa []*_ea.Image) error

func _bb() standardType { return standardType{_efe: 2, _ee: "\u0041"} }
func _fgab(_gfac *_da.CompliancePdfReader, _cfcc standardType, _cdab bool) (_dbdgg []ViolatedRule) {
	_deefc, _aead := _acgab(_gfac)
	if !_aead {
		return []ViolatedRule{_ba("\u0036.\u0037\u002e\u0032\u002d\u0031", "\u0063a\u0074a\u006c\u006f\u0067\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064\u002e")}
	}
	_efcd := _deefc.Get("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061")
	if _efcd == nil {
		return []ViolatedRule{_ba("\u0036.\u0037\u002e\u0032\u002d\u0031", "\u006e\u006f\u0020\u0027\u004d\u0065\u0074\u0061d\u0061\u0074\u0061' \u006b\u0065\u0079\u0020\u0066\u006fu\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0063\u0061\u0074\u0061\u006co\u0067\u002e"), _ba("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e")}
	}
	_bdea, _aead := _ad.GetStream(_efcd)
	if !_aead {
		return []ViolatedRule{_ba("\u0036.\u0037\u002e\u0032\u002d\u0032", "\u0063\u0061\u0074a\u006c\u006f\u0067\u0020\u0027\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0027\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020a\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u002e"), _ba("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e")}
	}
	if _bdea.Get("\u0046\u0069\u006c\u0074\u0065\u0072") != nil {
		_dbdgg = append(_dbdgg, _ba("\u0036.\u0037\u002e\u0032\u002d\u0032", "M\u0065\u0074a\u0064\u0061\u0074\u0061\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0069\u0065\u0073\u0020\u0073\u0068\u0061\u006c\u006c \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074h\u0065\u0020\u0046\u0069\u006c\u0074\u0065\u0072\u0020\u006b\u0065y\u002e"))
	}
	_bbff, _bbcg := _gf.LoadDocument(_bdea.Stream)
	if _bbcg != nil {
		return []ViolatedRule{_ba("\u0036.\u0037\u002e\u0039\u002d\u0031", "The\u0020\u006d\u0065\u0074a\u0064\u0061t\u0061\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0063o\u006e\u0066\u006f\u0072\u006d\u0020\u0074o\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0061\u006e\u0064\u0020\u0077\u0065\u006c\u006c\u0020\u0066\u006f\u0072\u006de\u0064\u0020\u0050\u0044\u0046\u0041\u0045\u0078\u0074e\u006e\u0073\u0069\u006f\u006e\u0020\u0053\u0063\u0068\u0065\u006da\u0020\u0066\u006fr\u0020\u0061\u006c\u006c\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0073\u002e")}
	}
	_dcdba := _bbff.GetGoXmpDocument()
	var _cebc []*_ga.Namespace
	for _, _daffa := range _dcdba.Namespaces() {
		switch _daffa.Name {
		case _cb.NsDc.Name, _f.NsPDF.Name, _ef.NsXmp.Name, _eb.NsXmpRights.Name, _ebe.Namespace.Name, _be.Namespace.Name, _cf.NsXmpMM.Name, _be.FieldNS.Name, _be.SchemaNS.Name, _be.PropertyNS.Name, "\u0073\u0074\u0045v\u0074", "\u0073\u0074\u0056e\u0072", "\u0073\u0074\u0052e\u0066", "\u0073\u0074\u0044i\u006d", "\u0078a\u0070\u0047\u0049\u006d\u0067", "\u0073\u0074\u004ao\u0062", "\u0078\u006d\u0070\u0069\u0064\u0071":
			continue
		}
		_cebc = append(_cebc, _daffa)
	}
	_dgbec := true
	_fdbb, _bbcg := _bbff.GetPdfaExtensionSchemas()
	if _bbcg == nil {
		for _, _dfgbe := range _cebc {
			var _bdacc bool
			for _fcdfg := range _fdbb {
				if _dfgbe.URI == _fdbb[_fcdfg].NamespaceURI {
					_bdacc = true
					break
				}
			}
			if !_bdacc {
				_dgbec = false
				break
			}
		}
	} else {
		_dgbec = false
	}
	if !_dgbec {
		_dbdgg = append(_dbdgg, _ba("\u0036.\u0037\u002e\u0039\u002d\u0032", "\u0050\u0072\u006f\u0070\u0065\u0072\u0074i\u0065\u0073 \u0073\u0070\u0065\u0063\u0069\u0066\u0069ed\u0020\u0069\u006e\u0020\u0058M\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0073\u0068\u0061\u006cl\u0020\u0075\u0073\u0065\u0020\u0065\u0069\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0065\u0020\u0070\u0072\u0065\u0064\u0065\u0066\u0069\u006e\u0065\u0064\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073 \u0064\u0065\u0066i\u006e\u0065\u0064\u0020\u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0053\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006fn\u002c\u0020\u006f\u0072\u0020\u0065\u0078\u0074\u0065ns\u0069\u006f\u006e\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u0074\u0068\u0061\u0074 \u0063\u006f\u006d\u0070\u006c\u0079\u0020\u0077\u0069\u0074h\u0020\u0058\u004d\u0050\u0020\u0053\u0070e\u0063\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002e"))
	}
	_gfccf, _bbcg := _gfac.GetPdfInfo()
	if _bbcg == nil {
		if !_cagb(_gfccf, _bbff) {
			_dbdgg = append(_dbdgg, _ba("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e"))
		}
	} else if _, _fggf := _bbff.GetMediaManagement(); _fggf {
		_dbdgg = append(_dbdgg, _ba("\u0036.\u0037\u002e\u0033\u002d\u0031", "\u0049\u0066\u0020\u005b\u0061\u0020\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u0020\u0069\u006e\u0066o\u0072\u006d\u0061t\u0069\u006f\u006e\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0061\u0070p\u0065\u0061r\u0073\u0020\u0069n\u0020\u0061 \u0064\u006f\u0063um\u0065\u006e\u0074\u005d\u002c\u0020\u0074\u0068\u0065n\u0020\u0061\u006c\u006c\u0020\u006f\u0066\u0020\u0069\u0074\u0073\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u006e\u0061\u006c\u006f\u0067\u006fu\u0073\u0020\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073 \u0069\u006e\u0020\u0070\u0072\u0065\u0064e\u0066\u0069\u006e\u0065\u0064\u0020\u0058\u004d\u0050\u0020\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u0020\u2026 \u0073\u0068\u0061\u006c\u006c\u0020\u0061\u006c\u0073\u006f\u0020\u0062\u0065\u0020\u0065\u006d\u0062\u0065\u0064\u0064\u0065d\u0020\u0069\u006e\u0020\u0074he\u0020\u0066i\u006c\u0065 \u0069\u006e\u0020\u0058\u004d\u0050\u0020\u0066\u006f\u0072\u006d\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0071\u0075\u0069\u0076\u0061\u006c\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u002e"))
	}
	_babg, _aead := _bbff.GetPdfAID()
	if !_aead {
		_dbdgg = append(_dbdgg, _ba("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0031", "\u0054\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0061n\u0064\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006ec\u0065\u0020\u006c\u0065\u0076\u0065l\u0020\u006f\u0066\u0020\u0061\u0020\u0066\u0069\u006c\u0065\u0020\u0073h\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0073\u0070e\u0063\u0069\u0066\u0069\u0065\u0064\u0020\u0075\u0073\u0069\u006e\u0067\u0020\u0074\u0068\u0065\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0049\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e\u0020\u0073\u0063h\u0065\u006da."))
	} else {
		if _babg.Part != _cfcc._efe {
			_dbdgg = append(_dbdgg, _ba("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0032", "\u0054h\u0065\u0020\u0076\u0061lue\u0020\u006f\u0066\u0020p\u0064\u0066\u0061\u0069\u0064\u003a\u0070\u0061\u0072\u0074 \u0073\u0068\u0061\u006c\u006c\u0020\u0062\u0065\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0072\u0074\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u006f\u0066\u0020\u0049\u0053\u004f\u002019\u0030\u0030\u0035 \u0074\u006f\u0020\u0077\u0068i\u0063h\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065 \u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0073\u002e"))
		}
		if _cfcc._ee == "\u0041" && _babg.Conformance != "\u0041" {
			_dbdgg = append(_dbdgg, _ba("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063i\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063o\u006e\u0066\u006fr\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0041\u002e\u0020\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0042\u0020\u0063\u006f\u006e\u0066o\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e"))
		} else if _cfcc._ee == "\u0042" && (_babg.Conformance != "\u0041" && _babg.Conformance != "\u0042") {
			_dbdgg = append(_dbdgg, _ba("\u0036\u002e\u0037\u002e\u0031\u0031\u002d\u0033", "\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0041\u0020\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065 \u0073\u0068\u0061\u006c\u006c\u0020s\u0070\u0065\u0063i\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069\u0064\u003a\u0063o\u006e\u0066\u006fr\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0041\u002e\u0020\u0041\u0020\u004c\u0065\u0076e\u006c\u0020\u0042\u0020\u0063\u006f\u006e\u0066o\u0072\u006d\u0069\u006e\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020\u0070\u0064\u0066\u0061\u0069d\u003a\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065\u0020\u0061\u0073\u0020\u0042\u002e"))
		}
	}
	return _dbdgg
}

// NewProfile2B creates a new Profile2B with the given options.
func NewProfile2B(options *Profile2Options) *Profile2B {
	if options == nil {
		options = DefaultProfile2Options()
	}
	_cdcc(options)
	return &Profile2B{profile2{_eceb: *options, _fcdd: _fb()}}
}
func _ceaf(_caec *_da.CompliancePdfReader) ViolatedRule {
	if _caec.ParserMetadata().HasDataAfterEOF() {
		return _ba("\u0036.\u0031\u002e\u0033\u002d\u0033", "\u004e\u006f\u0020\u0064\u0061ta\u0020\u0073h\u0061\u006c\u006c\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0020\u0074\u0068\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0065\u006e\u0064\u002d\u006f\u0066\u002d\u0066\u0069l\u0065\u0020\u006da\u0072\u006b\u0065\u0072\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0061 \u0073\u0069\u006e\u0067\u006ce\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061\u006c \u0065\u006ed\u002do\u0066\u002d\u006c\u0069\u006e\u0065\u0020m\u0061\u0072\u006b\u0065\u0072\u002e")
	}
	return _bg
}
func _eeda(_fbad *_da.CompliancePdfReader) ViolatedRule {
	if _fbad.ParserMetadata().HasDataAfterEOF() {
		return _ba("\u0036.\u0031\u002e\u0033\u002d\u0033", "\u004e\u006f\u0020\u0064\u0061ta\u0020\u0073h\u0061\u006c\u006c\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0020\u0074\u0068\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0065\u006e\u0064\u002d\u006f\u0066\u002d\u0066\u0069l\u0065\u0020\u006da\u0072\u006b\u0065\u0072\u0020\u0065\u0078\u0063\u0065\u0070\u0074\u0020\u0061 \u0073\u0069\u006e\u0067\u006ce\u0020\u006f\u0070\u0074\u0069\u006f\u006e\u0061\u006c \u0065\u006ed\u002do\u0066\u002d\u006c\u0069\u006e\u0065\u0020m\u0061\u0072\u006b\u0065\u0072\u002e")
	}
	return _bg
}
func _fb() standardType { return standardType{_efe: 2, _ee: "\u0042"} }

var _ Profile = (*Profile2A)(nil)

// Profile2Options are the options that changes the way how optimizer may try to adapt document into PDF/A standard.
type Profile2Options struct {

	// CMYKDefaultColorSpace is an option that refers PDF/A
	CMYKDefaultColorSpace bool

	// Now is a function that returns current time.
	Now func() _a.Time

	// Xmp is the xmp options information.
	Xmp XmpOptions
}

func _edcgb(_eacb *_da.CompliancePdfReader) ViolatedRule {
	_gebgb, _bcda := _eacb.PdfReader.GetTrailer()
	if _bcda != nil {
		return _ba("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u006d\u0069\u0073s\u0069\u006e\u0067\u0020t\u0072\u0061\u0069\u006c\u0065\u0072\u0020i\u006e\u0020\u0074\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	}
	if _gebgb.Get("\u0049\u0044") == nil {
		return _ba("\u0036.\u0031\u002e\u0033\u002d\u0031", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0073\u0068a\u006c\u006c\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068e\u0020\u0027\u0049\u0044\u0027\u0020k\u0065\u0079\u0077o\u0072\u0064")
	}
	if _gebgb.Get("\u0045n\u0063\u0072\u0079\u0070\u0074") != nil {
		return _ba("\u0036.\u0031\u002e\u0033\u002d\u0032", "\u0054\u0068\u0065\u0020\u006b\u0065y\u0077\u006f\u0072\u0064\u0020'\u0045\u006e\u0063\u0072\u0079\u0070t\u0027\u0020\u0073\u0068\u0061l\u006c\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0075\u0073\u0065d\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0074\u0072\u0061\u0069\u006c\u0065\u0072 \u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u002e\u0020")
	}
	return _bg
}
func _gdgae(_fafg *_da.CompliancePdfReader) ViolatedRule {
	_cfbd := _fafg.ParserMetadata().HeaderCommentBytes()
	if _cfbd[0] > 127 && _cfbd[1] > 127 && _cfbd[2] > 127 && _cfbd[3] > 127 {
		return _bg
	}
	return _ba("\u0036.\u0031\u002e\u0032\u002d\u0032", "\u0054\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u006c\u0069\u006e\u0065\u0020\u0073\u0068\u0061\u006c\u006c b\u0065\u0020i\u006d\u006d\u0065\u0064\u0069a\u0074\u0065\u006c\u0079 \u0066\u006f\u006c\u006co\u0077\u0065\u0064\u0020\u0062\u0079\u0020\u0061\u0020\u0063\u006f\u006d\u006d\u0065n\u0074\u0020\u0063\u006f\u006e\u0073\u0069s\u0074\u0069\u006e\u0067\u0020o\u0066\u0020\u0061\u0020\u0025\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006fwe\u0064\u0020\u0062y\u0020a\u0074\u0009\u006c\u0065a\u0073\u0074\u0020f\u006f\u0075\u0072\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065r\u0073\u002c\u0020e\u0061\u0063\u0068\u0020\u006f\u0066\u0020\u0077\u0068\u006f\u0073\u0065 \u0065\u006e\u0063\u006f\u0064e\u0064\u0020\u0062\u0079\u0074e\u0020\u0076\u0061\u006c\u0075\u0065s\u0020\u0073\u0068\u0061\u006c\u006c\u0020\u0068\u0061\u0076\u0065\u0020\u0061\u0020\u0064e\u0063\u0069\u006d\u0061\u006c \u0076\u0061\u006c\u0075\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0031\u0032\u0037\u002e")
}
func _dgccb(_bcebc *_da.CompliancePdfReader) (_gbba []ViolatedRule) {
	if _bcebc.ParserMetadata().HasOddLengthHexStrings() {
		_gbba = append(_gbba, _ba("\u0036.\u0031\u002e\u0036\u002d\u0031", "\u0068\u0065\u0078a\u0064\u0065\u0063\u0069\u006d\u0061\u006c\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006f\u0066\u0020e\u0076\u0065\u006e\u0020\u0073\u0069\u007a\u0065"))
	}
	if _bcebc.ParserMetadata().HasOddLengthHexStrings() {
		_gbba = append(_gbba, _ba("\u0036.\u0031\u002e\u0036\u002d\u0032", "\u0068\u0065\u0078\u0061\u0064\u0065\u0063\u0069\u006da\u006c\u0020s\u0074\u0072\u0069\u006e\u0067\u0073\u0020\u0073\u0068o\u0075\u006c\u0064\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u006f\u006e\u006c\u0079\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0066\u0072\u006f\u006d\u0020\u0072\u0061n\u0067\u0065\u0020[\u0030\u002d\u0039\u003b\u0041\u002d\u0046\u003b\u0061\u002d\u0066\u005d"))
	}
	return _gbba
}
