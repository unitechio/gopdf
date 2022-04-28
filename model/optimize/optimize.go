package optimize

import (
	_ad "bytes"
	_b "crypto/md5"
	_fa "errors"
	_a "math"

	_f "bitbucket.org/shenghui0779/gopdf/common"
	_e "bitbucket.org/shenghui0779/gopdf/contentstream"
	_cg "bitbucket.org/shenghui0779/gopdf/core"
	_af "bitbucket.org/shenghui0779/gopdf/extractor"
	_d "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_ae "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_dc "bitbucket.org/shenghui0779/gopdf/model"
	_aa "github.com/unidoc/unitype"
	_ff "golang.org/x/image/draw"
)

// ObjectStreams groups PDF objects to object streams.
// It implements interface model.Optimizer.
type ObjectStreams struct{}

func _ggbd(_cgca []_cg.PdfObject) []*imageInfo {
	_dbf := _cg.PdfObjectName("\u0053u\u0062\u0074\u0079\u0070\u0065")
	_fefd := make(map[*_cg.PdfObjectStream]struct{})
	var _fg []*imageInfo
	for _, _fbd := range _cgca {
		_dce, _edf := _cg.GetStream(_fbd)
		if !_edf {
			continue
		}
		if _, _babe := _fefd[_dce]; _babe {
			continue
		}
		_fefd[_dce] = struct{}{}
		_faec := _dce.PdfObjectDictionary.Get(_dbf)
		_ggda, _edf := _cg.GetName(_faec)
		if !_edf || string(*_ggda) != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		_abbe := &imageInfo{Stream: _dce, BitsPerComponent: 8}
		if _dge, _cgad := _cg.GetIntVal(_dce.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _cgad {
			_abbe.BitsPerComponent = _dge
		}
		if _dccg, _cfeb := _cg.GetIntVal(_dce.Get("\u0057\u0069\u0064t\u0068")); _cfeb {
			_abbe.Width = _dccg
		}
		if _ccca, _fbbb := _cg.GetIntVal(_dce.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _fbbb {
			_abbe.Height = _ccca
		}
		_edcg, _gfg := _dc.NewPdfColorspaceFromPdfObject(_dce.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"))
		if _gfg != nil {
			_f.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gfg)
			continue
		}
		if _edcg == nil {
			_dea, _eba := _cg.GetName(_dce.Get("\u0046\u0069\u006c\u0074\u0065\u0072"))
			if _eba {
				switch _dea.String() {
				case "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065", "J\u0042\u0049\u0047\u0032\u0044\u0065\u0063\u006f\u0064\u0065":
					_edcg = _dc.NewPdfColorspaceDeviceGray()
					_abbe.BitsPerComponent = 1
				}
			}
		}
		switch _gcgc := _edcg.(type) {
		case *_dc.PdfColorspaceDeviceRGB:
			_abbe.ColorComponents = 3
		case *_dc.PdfColorspaceDeviceGray:
			_abbe.ColorComponents = 1
		default:
			_f.Log.Debug("\u004f\u0070\u0074\u0069\u006d\u0069\u007aa\u0074\u0069\u006fn\u0020\u0069\u0073 \u006e\u006ft\u0020\u0073\u0075\u0070\u0070\u006fr\u0074ed\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u006b\u0069\u0070", _gcgc)
			continue
		}
		_fg = append(_fg, _abbe)
	}
	return _fg
}

// CompressStreams compresses uncompressed streams.
// It implements interface model.Optimizer.
type CompressStreams struct{}

// Optimize optimizes PDF objects to decrease PDF size.
func (_aaa *CombineDuplicateDirectObjects) Optimize(objects []_cg.PdfObject) (_bag []_cg.PdfObject, _aaf error) {
	_cbff(objects)
	_gfe := make(map[string][]*_cg.PdfObjectDictionary)
	var _cda func(_cbca *_cg.PdfObjectDictionary)
	_cda = func(_edc *_cg.PdfObjectDictionary) {
		for _, _gfc := range _edc.Keys() {
			_cbd := _edc.Get(_gfc)
			if _egc, _gaae := _cbd.(*_cg.PdfObjectDictionary); _gaae {
				_dbe := _b.New()
				_dbe.Write([]byte(_egc.WriteString()))
				_cfd := string(_dbe.Sum(nil))
				_gfe[_cfd] = append(_gfe[_cfd], _egc)
				_cda(_egc)
			}
		}
	}
	for _, _baad := range objects {
		_ggb, _cae := _baad.(*_cg.PdfIndirectObject)
		if !_cae {
			continue
		}
		if _dcdf, _defg := _ggb.PdfObject.(*_cg.PdfObjectDictionary); _defg {
			_cda(_dcdf)
		}
	}
	_abc := make([]_cg.PdfObject, 0, len(_gfe))
	_fdf := make(map[_cg.PdfObject]_cg.PdfObject)
	for _, _fce := range _gfe {
		if len(_fce) < 2 {
			continue
		}
		_dga := _cg.MakeDict()
		_dga.Merge(_fce[0])
		_gab := _cg.MakeIndirectObject(_dga)
		_abc = append(_abc, _gab)
		for _gdef := 0; _gdef < len(_fce); _gdef++ {
			_bace := _fce[_gdef]
			_fdf[_bace] = _gab
		}
	}
	_bag = make([]_cg.PdfObject, len(objects))
	copy(_bag, objects)
	_bag = append(_abc, _bag...)
	_cacg(_bag, _fdf)
	return _bag, nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_afe *CombineDuplicateStreams) Optimize(objects []_cg.PdfObject) (_ddc []_cg.PdfObject, _dba error) {
	_cgcb := make(map[_cg.PdfObject]_cg.PdfObject)
	_agb := make(map[_cg.PdfObject]struct{})
	_bafe := make(map[string][]*_cg.PdfObjectStream)
	for _, _bbf := range objects {
		if _fea, _ebf := _bbf.(*_cg.PdfObjectStream); _ebf {
			_ece := _b.New()
			_ece.Write(_fea.Stream)
			_ece.Write([]byte(_fea.PdfObjectDictionary.WriteString()))
			_faad := string(_ece.Sum(nil))
			_bafe[_faad] = append(_bafe[_faad], _fea)
		}
	}
	for _, _affa := range _bafe {
		if len(_affa) < 2 {
			continue
		}
		_bga := _affa[0]
		for _cff := 1; _cff < len(_affa); _cff++ {
			_fcb := _affa[_cff]
			_cgcb[_fcb] = _bga
			_agb[_fcb] = struct{}{}
		}
	}
	_ddc = make([]_cg.PdfObject, 0, len(objects)-len(_agb))
	for _, _eec := range objects {
		if _, _eeg := _agb[_eec]; _eeg {
			continue
		}
		_ddc = append(_ddc, _eec)
	}
	_cacg(_ddc, _cgcb)
	return _ddc, nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_bbed *Image) Optimize(objects []_cg.PdfObject) (_eegb []_cg.PdfObject, _adc error) {
	if _bbed.ImageQuality <= 0 {
		return objects, nil
	}
	_bbg := _ggbd(objects)
	if len(_bbg) == 0 {
		return objects, nil
	}
	_abba := make(map[_cg.PdfObject]_cg.PdfObject)
	_bcf := make(map[_cg.PdfObject]struct{})
	for _, _ddcg := range _bbg {
		_egg := _ddcg.Stream.Get("\u0053\u004d\u0061s\u006b")
		_bcf[_egg] = struct{}{}
	}
	for _bage, _cbde := range _bbg {
		_egce := _cbde.Stream
		if _, _bagg := _bcf[_egce]; _bagg {
			continue
		}
		_dcgg, _ggde := _dc.NewXObjectImageFromStream(_egce)
		if _ggde != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _ggde)
			continue
		}
		switch _dcgg.Filter.(type) {
		case *_cg.JBIG2Encoder:
			continue
		case *_cg.CCITTFaxEncoder:
			continue
		}
		_badg, _ggde := _dcgg.ToImage()
		if _ggde != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _ggde)
			continue
		}
		_dad := _cg.NewDCTEncoder()
		_dad.ColorComponents = _badg.ColorComponents
		_dad.Quality = _bbed.ImageQuality
		_dad.BitsPerComponent = _cbde.BitsPerComponent
		_dad.Width = _cbde.Width
		_dad.Height = _cbde.Height
		_gccd, _ggde := _dad.EncodeBytes(_badg.Data)
		if _ggde != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _ggde)
			continue
		}
		var _dfc _cg.StreamEncoder
		_dfc = _dad
		{
			_dccgd := _cg.NewFlateEncoder()
			_ffbg := _cg.NewMultiEncoder()
			_ffbg.AddEncoder(_dccgd)
			_ffbg.AddEncoder(_dad)
			_dcca, _ccf := _ffbg.EncodeBytes(_badg.Data)
			if _ccf != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _ccf)
				continue
			}
			if len(_dcca) < len(_gccd) {
				_f.Log.Trace("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063\u0020\u0069\u006d\u0070\u0072\u006f\u0076\u0065\u0073\u003a\u0020\u0025\u0064\u0020\u0074o\u0020\u0025\u0064\u0020\u0028o\u0072\u0069g\u0020\u0025\u0064\u0029", len(_gccd), len(_dcca), len(_egce.Stream))
				_gccd = _dcca
				_dfc = _ffbg
			}
		}
		_feg := len(_egce.Stream)
		if _feg < len(_gccd) {
			continue
		}
		_bfcf := &_cg.PdfObjectStream{Stream: _gccd}
		_bfcf.PdfObjectReference = _egce.PdfObjectReference
		_bfcf.PdfObjectDictionary = _cg.MakeDict()
		_bfcf.Merge(_egce.PdfObjectDictionary)
		_bfcf.Merge(_dfc.MakeStreamDict())
		_bfcf.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _cg.MakeInteger(int64(len(_gccd))))
		_abba[_egce] = _bfcf
		_bbg[_bage].Stream = _bfcf
	}
	_eegb = make([]_cg.PdfObject, len(objects))
	copy(_eegb, objects)
	_cacg(_eegb, _abba)
	return _eegb, nil
}
func _ffg(_fcaf _cg.PdfObject) (_ceea string, _eaff []_cg.PdfObject) {
	var _agc _ad.Buffer
	switch _acg := _fcaf.(type) {
	case *_cg.PdfIndirectObject:
		_eaff = append(_eaff, _acg)
		_fcaf = _acg.PdfObject
	}
	switch _egef := _fcaf.(type) {
	case *_cg.PdfObjectStream:
		if _edb, _gfbd := _cg.DecodeStream(_egef); _gfbd == nil {
			_agc.Write(_edb)
			_eaff = append(_eaff, _egef)
		}
	case *_cg.PdfObjectArray:
		for _, _fgd := range _egef.Elements() {
			switch _daa := _fgd.(type) {
			case *_cg.PdfObjectStream:
				if _ecd, _agea := _cg.DecodeStream(_daa); _agea == nil {
					_agc.Write(_ecd)
					_eaff = append(_eaff, _daa)
				}
			}
		}
	}
	return _agc.String(), _eaff
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_dcbb *ImagePPI) Optimize(objects []_cg.PdfObject) (_caa []_cg.PdfObject, _caad error) {
	if _dcbb.ImageUpperPPI <= 0 {
		return objects, nil
	}
	_gadb := _ggbd(objects)
	if len(_gadb) == 0 {
		return objects, nil
	}
	_dfcc := make(map[_cg.PdfObject]struct{})
	for _, _cgf := range _gadb {
		_aeg := _cgf.Stream.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b")
		_dfcc[_aeg] = struct{}{}
	}
	_cbgc := make(map[*_cg.PdfObjectStream]*imageInfo)
	for _, _dead := range _gadb {
		_cbgc[_dead.Stream] = _dead
	}
	var _ebba *_cg.PdfObjectDictionary
	for _, _aee := range objects {
		if _egd, _ccg := _cg.GetDict(_aee); _ebba == nil && _ccg {
			if _ffe, _eggb := _cg.GetName(_egd.Get("\u0054\u0079\u0070\u0065")); _eggb && *_ffe == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
				_ebba = _egd
			}
		}
	}
	if _ebba == nil {
		return objects, nil
	}
	_gebf, _afb := _cg.GetDict(_ebba.Get("\u0050\u0061\u0067e\u0073"))
	if !_afb {
		return objects, nil
	}
	_fefde, _fece := _cg.GetArray(_gebf.Get("\u004b\u0069\u0064\u0073"))
	if !_fece {
		return objects, nil
	}
	for _, _gedb := range _fefde.Elements() {
		_bef := make(map[string]*imageInfo)
		_dbeb, _ebe := _cg.GetDict(_gedb)
		if !_ebe {
			continue
		}
		_bgd, _ := _ffg(_dbeb.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		if len(_bgd) == 0 {
			continue
		}
		_dbac, _egec := _cg.GetDict(_dbeb.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
		if !_egec {
			continue
		}
		_adbe, _agac := _dc.NewPdfPageResourcesFromDict(_dbac)
		if _agac != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020-\u0020\u0069\u0067\u006e\u006fr\u0069\u006eg\u003a\u0020\u0025\u0076", _agac)
			continue
		}
		_gcf, _cac := _cg.GetDict(_dbac.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
		if !_cac {
			continue
		}
		_agee := _gcf.Keys()
		for _, _gef := range _agee {
			if _defb, _ebg := _cg.GetStream(_gcf.Get(_gef)); _ebg {
				if _acde, _gaag := _cbgc[_defb]; _gaag {
					_bef[string(_gef)] = _acde
				}
			}
		}
		_egcf := _e.NewContentStreamParser(_bgd)
		_cgb, _agac := _egcf.Parse()
		if _agac != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _agac)
			continue
		}
		_egde := _e.NewContentStreamProcessor(*_cgb)
		_egde.AddHandler(_e.HandlerConditionEnumAllOperands, "", func(_faf *_e.ContentStreamOperation, _cdg _e.GraphicsState, _cef *_dc.PdfPageResources) error {
			switch _faf.Operand {
			case "\u0044\u006f":
				if len(_faf.Params) != 1 {
					_f.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020w\u0069\u0074\u0068\u0020\u006c\u0065\u006e\u0028\u0070\u0061ra\u006d\u0073\u0029 \u0021=\u0020\u0031")
					return nil
				}
				_ccge, _eda := _cg.GetName(_faf.Params[0])
				if !_eda {
					_f.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f\u006e\u0020\u004e\u0061\u006d\u0065\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072")
					return nil
				}
				if _cee, _gfaf := _bef[string(*_ccge)]; _gfaf {
					_badb := _cdg.CTM.ScalingFactorX()
					_abaa := _cdg.CTM.ScalingFactorY()
					_ebc, _ffa := _badb/72.0, _abaa/72.0
					_caee, _feee := float64(_cee.Width)/_ebc, float64(_cee.Height)/_ffa
					if _ebc == 0 || _ffa == 0 {
						_caee = 72.0
						_feee = 72.0
					}
					_cee.PPI = _a.Max(_cee.PPI, _caee)
					_cee.PPI = _a.Max(_cee.PPI, _feee)
				}
			}
			return nil
		})
		_agac = _egde.Process(_adbe)
		if _agac != nil {
			_f.Log.Debug("E\u0052\u0052\u004f\u0052 p\u0072o\u0063\u0065\u0073\u0073\u0069n\u0067\u003a\u0020\u0025\u002b\u0076", _agac)
			continue
		}
	}
	for _, _add := range _gadb {
		if _, _bdd := _dfcc[_add.Stream]; _bdd {
			continue
		}
		if _add.PPI <= _dcbb.ImageUpperPPI {
			continue
		}
		_aaca, _fegd := _dc.NewXObjectImageFromStream(_add.Stream)
		if _fegd != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _fegd)
			continue
		}
		var _egdef imageModifications
		_egdef.Scale = _dcbb.ImageUpperPPI / _add.PPI
		if _add.BitsPerComponent == 1 && _add.ColorComponents == 1 {
			_agf := _a.Round(_add.PPI / _dcbb.ImageUpperPPI)
			_bdf := _d.NextPowerOf2(uint(_agf))
			if _d.InDelta(float64(_bdf), 1/_egdef.Scale, 0.3) {
				_egdef.Scale = float64(1) / float64(_bdf)
			}
			if _, _fabd := _aaca.Filter.(*_cg.JBIG2Encoder); !_fabd {
				_egdef.Encoding = _cg.NewJBIG2Encoder()
			}
		}
		if _fegd = _eag(_aaca, _egdef); _fegd != nil {
			_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072 \u0073\u0063\u0061\u006c\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u006be\u0065\u0070\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _fegd)
			continue
		}
		_egdef.Encoding = nil
		if _ega, _cecf := _cg.GetStream(_add.Stream.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b")); _cecf {
			_afcg, _dfd := _dc.NewXObjectImageFromStream(_ega)
			if _dfd != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _dfd)
				continue
			}
			if _dfd = _eag(_afcg, _egdef); _dfd != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _dfd)
				continue
			}
		}
	}
	return objects, nil
}
func _db(_ge *_cg.PdfObjectStream) error {
	_ac, _ga := _cg.DecodeStream(_ge)
	if _ga != nil {
		return _ga
	}
	_bg := _e.NewContentStreamParser(string(_ac))
	_bb, _ga := _bg.Parse()
	if _ga != nil {
		return _ga
	}
	_bb = _ef(_bb)
	_ba := _bb.Bytes()
	if len(_ba) >= len(_ac) {
		return nil
	}
	_be, _ga := _cg.MakeStream(_bb.Bytes(), _cg.NewFlateEncoder())
	if _ga != nil {
		return _ga
	}
	_ge.Stream = _be.Stream
	_ge.Merge(_be.PdfObjectDictionary)
	return nil
}

type content struct {
	_gaab string
	_ec   *_dc.PdfPageResources
}

func _ef(_ce *_e.ContentStreamOperations) *_e.ContentStreamOperations {
	if _ce == nil {
		return nil
	}
	_de := _e.ContentStreamOperations{}
	for _, _da := range *_ce {
		switch _da.Operand {
		case "\u0042\u0044\u0043", "\u0042\u004d\u0043", "\u0045\u004d\u0043":
			continue
		case "\u0054\u006d":
			if len(_da.Params) == 6 {
				if _dd, _gf := _cg.GetNumbersAsFloat(_da.Params); _gf == nil {
					if _dd[0] == 1 && _dd[1] == 0 && _dd[2] == 0 && _dd[3] == 1 {
						_da = &_e.ContentStreamOperation{Params: []_cg.PdfObject{_da.Params[4], _da.Params[5]}, Operand: "\u0054\u0064"}
					}
				}
			}
		}
		_de = append(_de, _da)
	}
	return &_de
}

// ImagePPI optimizes images by scaling images such that the PPI (pixels per inch) is never higher than ImageUpperPPI.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type ImagePPI struct{ ImageUpperPPI float64 }

// Optimize optimizes PDF objects to decrease PDF size.
func (_gee *ObjectStreams) Optimize(objects []_cg.PdfObject) (_cfbb []_cg.PdfObject, _cadd error) {
	_cece := &_cg.PdfObjectStreams{}
	_beg := make([]_cg.PdfObject, 0, len(objects))
	for _, _bgb := range objects {
		if _abdf, _bfdd := _bgb.(*_cg.PdfIndirectObject); _bfdd && _abdf.GenerationNumber == 0 {
			_cece.Append(_bgb)
		} else {
			_beg = append(_beg, _bgb)
		}
	}
	if _cece.Len() == 0 {
		return _beg, nil
	}
	_cfbb = make([]_cg.PdfObject, 0, len(_beg)+_cece.Len()+1)
	if _cece.Len() > 1 {
		_cfbb = append(_cfbb, _cece)
	}
	_cfbb = append(_cfbb, _cece.Elements()...)
	_cfbb = append(_cfbb, _beg...)
	return _cfbb, nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_cfe *CleanFonts) Optimize(objects []_cg.PdfObject) (_aacb []_cg.PdfObject, _bce error) {
	var _cec map[*_cg.PdfObjectStream]struct{}
	if _cfe.Subset {
		var _fccd error
		_cec, _fccd = _eac(objects)
		if _fccd != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0073u\u0062s\u0065\u0074\u0074\u0069\u006e\u0067\u003a \u0025\u0076", _fccd)
			return nil, _fccd
		}
	}
	for _, _fcg := range objects {
		_fef, _gge := _cg.GetStream(_fcg)
		if !_gge {
			continue
		}
		if _, _fcde := _cec[_fef]; _fcde {
			continue
		}
		_cd, _aae := _cg.NewEncoderFromStream(_fef)
		if _aae != nil {
			_f.Log.Debug("\u0045\u0052RO\u0052\u0020\u0067e\u0074\u0074\u0069\u006eg e\u006eco\u0064\u0065\u0072\u003a\u0020\u0025\u0076 -\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067", _aae)
			continue
		}
		_ffbb, _aae := _cd.DecodeStream(_fef)
		if _aae != nil {
			_f.Log.Debug("\u0044\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065r\u0072\u006f\u0072\u0020\u003a\u0020\u0025v\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067", _aae)
			continue
		}
		if len(_ffbb) < 4 {
			continue
		}
		_edg := string(_ffbb[:4])
		if _edg == "\u004f\u0054\u0054\u004f" {
			continue
		}
		if _edg != "\u0000\u0001\u0000\u0000" && _edg != "\u0074\u0072\u0075\u0065" {
			continue
		}
		_bac, _aae := _aa.Parse(_ad.NewReader(_ffbb))
		if _aae != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020P\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _aae)
			continue
		}
		_aae = _bac.Optimize()
		if _aae != nil {
			_f.Log.Debug("\u0045\u0052RO\u0052\u0020\u004fp\u0074\u0069\u006d\u0069zin\u0067 f\u006f\u006e\u0074\u003a\u0020\u0025\u0076 -\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067", _aae)
			continue
		}
		var _dcgf _ad.Buffer
		_aae = _bac.Write(&_dcgf)
		if _aae != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020W\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _aae)
			continue
		}
		if _dcgf.Len() > len(_ffbb) {
			_f.Log.Debug("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070")
			continue
		}
		_gad, _aae := _cg.MakeStream(_dcgf.Bytes(), _cg.NewFlateEncoder())
		if _aae != nil {
			continue
		}
		*_fef = *_gad
		_fef.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _cg.MakeInteger(int64(_dcgf.Len())))
	}
	return objects, nil
}
func _cbff(_dcdd []_cg.PdfObject) {
	for _efg, _bbedb := range _dcdd {
		switch _bca := _bbedb.(type) {
		case *_cg.PdfIndirectObject:
			_bca.ObjectNumber = int64(_efg + 1)
			_bca.GenerationNumber = 0
		case *_cg.PdfObjectStream:
			_bca.ObjectNumber = int64(_efg + 1)
			_bca.GenerationNumber = 0
		case *_cg.PdfObjectStreams:
			_bca.ObjectNumber = int64(_efg + 1)
			_bca.GenerationNumber = 0
		}
	}
}
func _cacg(_fed []_cg.PdfObject, _cecfb map[_cg.PdfObject]_cg.PdfObject) {
	if len(_cecfb) == 0 {
		return
	}
	for _eae, _dadfg := range _fed {
		if _cccc, _affdb := _cecfb[_dadfg]; _affdb {
			_fed[_eae] = _cccc
			continue
		}
		_cecfb[_dadfg] = _dadfg
		switch _eebd := _dadfg.(type) {
		case *_cg.PdfObjectArray:
			_gbd := make([]_cg.PdfObject, _eebd.Len())
			copy(_gbd, _eebd.Elements())
			_cacg(_gbd, _cecfb)
			for _cbfe, _aebd := range _gbd {
				_eebd.Set(_cbfe, _aebd)
			}
		case *_cg.PdfObjectStreams:
			_cacg(_eebd.Elements(), _cecfb)
		case *_cg.PdfObjectStream:
			_dadc := []_cg.PdfObject{_eebd.PdfObjectDictionary}
			_cacg(_dadc, _cecfb)
			_eebd.PdfObjectDictionary = _dadc[0].(*_cg.PdfObjectDictionary)
		case *_cg.PdfObjectDictionary:
			_afcd := _eebd.Keys()
			_caaa := make([]_cg.PdfObject, len(_afcd))
			for _gcgf, _eef := range _afcd {
				_caaa[_gcgf] = _eebd.Get(_eef)
			}
			_cacg(_caaa, _cecfb)
			for _adca, _egf := range _afcd {
				_eebd.Set(_egf, _caaa[_adca])
			}
		case *_cg.PdfIndirectObject:
			_eeff := []_cg.PdfObject{_eebd.PdfObject}
			_cacg(_eeff, _cecfb)
			_eebd.PdfObject = _eeff[0]
		}
	}
}
func _fbc(_gade _cg.PdfObject) []content {
	if _gade == nil {
		return nil
	}
	_fd, _adg := _cg.GetArray(_gade)
	if !_adg {
		_f.Log.Debug("\u0041\u006e\u006e\u006fts\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
		return nil
	}
	var _bafd []content
	for _, _efc := range _fd.Elements() {
		_beb, _faa := _cg.GetDict(_efc)
		if !_faa {
			_f.Log.Debug("I\u0067\u006e\u006f\u0072\u0069\u006eg\u0020\u006e\u006f\u006e\u002d\u0064i\u0063\u0074\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006e\u0020\u0041\u006e\u006e\u006ft\u0073")
			continue
		}
		_efad, _faa := _cg.GetDict(_beb.Get("\u0041\u0050"))
		if !_faa {
			_f.Log.Debug("\u004e\u006f\u0020\u0041P \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067")
			continue
		}
		_bed := _cg.TraceToDirectObject(_efad.Get("\u004e"))
		if _bed == nil {
			_f.Log.Debug("N\u006f\u0020\u004e\u0020en\u0074r\u0079\u0020\u002d\u0020\u0073k\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		var _bbe *_cg.PdfObjectStream
		switch _aba := _bed.(type) {
		case *_cg.PdfObjectDictionary:
			_facg, _dec := _cg.GetName(_beb.Get("\u0041\u0053"))
			if !_dec {
				_f.Log.Debug("\u004e\u006f\u0020\u0041S \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067")
				continue
			}
			_bbe, _dec = _cg.GetStream(_aba.Get(*_facg))
			if !_dec {
				_f.Log.Debug("\u0046o\u0072\u006d\u0020\u006eo\u0074\u0020\u0066\u006f\u0075n\u0064 \u002d \u0073\u006b\u0069\u0070\u0070\u0069\u006eg")
				continue
			}
		case *_cg.PdfObjectStream:
			_bbe = _aba
		}
		if _bbe == nil {
			_f.Log.Debug("\u0046\u006f\u0072m\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0028n\u0069\u006c\u0029\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		_cdb, _abfb := _dc.NewXObjectFormFromStream(_bbe)
		if _abfb != nil {
			_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _abfb)
			continue
		}
		_bbc, _abfb := _cdb.GetContentStream()
		if _abfb != nil {
			_f.Log.Debug("E\u0072\u0072\u006f\u0072\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0063\u006fn\u0074\u0065\u006et\u0073:\u0020\u0025\u0076", _abfb)
			continue
		}
		_bafd = append(_bafd, content{_gaab: string(_bbc), _ec: _cdb.Resources})
	}
	return _bafd
}

type imageModifications struct {
	Scale    float64
	Encoding _cg.StreamEncoder
}

// CleanFonts cleans up embedded fonts, reducing font sizes.
type CleanFonts struct {

	// Subset embedded fonts if encountered (if true).
	// Otherwise attempts to reduce the font program.
	Subset bool
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_bf *CleanContentstream) Optimize(objects []_cg.PdfObject) (_bc []_cg.PdfObject, _dcc error) {
	_adf := map[*_cg.PdfObjectStream]struct{}{}
	var _bea []*_cg.PdfObjectStream
	_geb := func(_ag *_cg.PdfObjectStream) {
		if _, _gfb := _adf[_ag]; !_gfb {
			_adf[_ag] = struct{}{}
			_bea = append(_bea, _ag)
		}
	}
	_fb := map[_cg.PdfObject]bool{}
	_ed := map[_cg.PdfObject]bool{}
	for _, _ceg := range objects {
		switch _ab := _ceg.(type) {
		case *_cg.PdfIndirectObject:
			switch _gcc := _ab.PdfObject.(type) {
			case *_cg.PdfObjectDictionary:
				if _aac, _baa := _cg.GetName(_gcc.Get("\u0054\u0079\u0070\u0065")); !_baa || _aac.String() != "\u0050\u0061\u0067\u0065" {
					continue
				}
				if _dae, _dac := _cg.GetStream(_gcc.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073")); _dac {
					_geb(_dae)
				} else if _cb, _fc := _cg.GetArray(_gcc.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073")); _fc {
					var _efd []*_cg.PdfObjectStream
					for _, _ebb := range _cb.Elements() {
						if _gag, _fac := _cg.GetStream(_ebb); _fac {
							_efd = append(_efd, _gag)
						}
					}
					if len(_efd) > 0 {
						var _aff _ad.Buffer
						for _, _cbb := range _efd {
							if _cba, _dcd := _cg.DecodeStream(_cbb); _dcd == nil {
								_aff.Write(_cba)
							}
							_fb[_cbb] = true
						}
						_beaa, _ged := _cg.MakeStream(_aff.Bytes(), _cg.NewFlateEncoder())
						if _ged != nil {
							return nil, _ged
						}
						_ed[_beaa] = true
						_gcc.Set("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _beaa)
						_geb(_beaa)
					}
				}
			}
		case *_cg.PdfObjectStream:
			if _ea, _bfg := _cg.GetName(_ab.Get("\u0054\u0079\u0070\u0065")); !_bfg || _ea.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
				continue
			}
			if _bfd, _afc := _cg.GetName(_ab.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); !_afc || _bfd.String() != "\u0046\u006f\u0072\u006d" {
				continue
			}
			_geb(_ab)
		}
	}
	for _, _gde := range _bea {
		_dcc = _db(_gde)
		if _dcc != nil {
			return nil, _dcc
		}
	}
	_bc = nil
	for _, _gff := range objects {
		if _fb[_gff] {
			continue
		}
		_bc = append(_bc, _gff)
	}
	for _cc := range _ed {
		_bc = append(_bc, _cc)
	}
	return _bc, nil
}
func _eac(_eaf []_cg.PdfObject) (_aeb map[*_cg.PdfObjectStream]struct{}, _dcg error) {
	_aeb = map[*_cg.PdfObjectStream]struct{}{}
	_ggf := map[*_dc.PdfFont]struct{}{}
	_bab := _dde(_eaf)
	for _, _bdga := range _bab._gdff {
		_dg, _gaf := _cg.GetDict(_bdga.PdfObject)
		if !_gaf {
			continue
		}
		_gdf, _gaf := _cg.GetDict(_dg.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
		if !_gaf {
			continue
		}
		_ddg, _ := _ffg(_dg.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		_dgc, _ccb := _dc.NewPdfPageResourcesFromDict(_gdf)
		if _ccb != nil {
			return nil, _ccb
		}
		_fcf := []content{{_gaab: _ddg, _ec: _dgc}}
		_fbb := _fbc(_dg.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _fbb != nil {
			_fcf = append(_fcf, _fbb...)
		}
		for _, _fcd := range _fcf {
			_ccbc, _abf := _af.NewFromContents(_fcd._gaab, _fcd._ec)
			if _abf != nil {
				return nil, _abf
			}
			_fab, _, _, _abf := _ccbc.ExtractPageText()
			if _abf != nil {
				return nil, _abf
			}
			for _, _fae := range _fab.Marks().Elements() {
				if _fae.Font == nil {
					continue
				}
				if _, _gcg := _ggf[_fae.Font]; !_gcg {
					_ggf[_fae.Font] = struct{}{}
				}
			}
		}
	}
	_gb := map[*_cg.PdfObjectStream][]*_dc.PdfFont{}
	for _bgf := range _ggf {
		_afcc := _bgf.FontDescriptor()
		if _afcc == nil || _afcc.FontFile2 == nil {
			continue
		}
		_gbg, _gfa := _cg.GetStream(_afcc.FontFile2)
		if !_gfa {
			continue
		}
		_gb[_gbg] = append(_gb[_gbg], _bgf)
	}
	for _eg := range _gb {
		var _dcdc []rune
		var _bge []_aa.GlyphIndex
		for _, _faee := range _gb[_eg] {
			switch _acd := _faee.Encoder().(type) {
			case *_ae.IdentityEncoder:
				_dddb := _acd.RegisteredRunes()
				_ee := make([]_aa.GlyphIndex, len(_dddb))
				for _dee, _cgd := range _dddb {
					_ee[_dee] = _aa.GlyphIndex(_cgd)
				}
				_bge = append(_bge, _ee...)
			case *_ae.TrueTypeFontEncoder:
				_cgaf := _acd.RegisteredRunes()
				_dcdc = append(_dcdc, _cgaf...)
			case _ae.SimpleEncoder:
				_ege := _acd.Charcodes()
				for _, _afa := range _ege {
					_cea, _fe := _acd.CharcodeToRune(_afa)
					if !_fe {
						_f.Log.Debug("\u0043\u0068a\u0072\u0063\u006f\u0064\u0065\u003c\u002d\u003e\u0072\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064: \u0025\u0064", _afa)
						continue
					}
					_dcdc = append(_dcdc, _cea)
				}
			}
		}
		_dcg = _ggfe(_eg, _dcdc, _bge)
		if _dcg != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006eg\u0020f\u006f\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _dcg)
			return nil, _dcg
		}
		_aeb[_eg] = struct{}{}
	}
	return _aeb, nil
}

// Options describes PDF optimization parameters.
type Options struct {
	CombineDuplicateStreams         bool
	CombineDuplicateDirectObjects   bool
	ImageUpperPPI                   float64
	ImageQuality                    int
	UseObjectStreams                bool
	CombineIdenticalIndirectObjects bool
	CompressStreams                 bool
	CleanFonts                      bool
	SubsetFonts                     bool
	CleanContentstream              bool
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_g *Chain) Optimize(objects []_cg.PdfObject) (_gc []_cg.PdfObject, _gg error) {
	_gd := objects
	for _, _cge := range _g._ffb {
		_eb, _cgc := _cge.Optimize(_gd)
		if _cgc != nil {
			_f.Log.Debug("\u0045\u0052\u0052OR\u0020\u004f\u0070\u0074\u0069\u006d\u0069\u007a\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u002b\u0076", _cgc)
			continue
		}
		_gd = _eb
	}
	return _gd, nil
}
func _bagb(_dcfa *_dc.Image, _fag float64) (*_dc.Image, error) {
	_ffbe, _fdfe := _dcfa.ToGoImage()
	if _fdfe != nil {
		return nil, _fdfe
	}
	var _cdc _d.Image
	_dadf, _fffc := _ffbe.(*_d.Monochrome)
	if _fffc {
		if _fdfe = _dadf.ResolveDecode(); _fdfe != nil {
			return nil, _fdfe
		}
		_cdc, _fdfe = _dadf.Scale(_fag)
		if _fdfe != nil {
			return nil, _fdfe
		}
	} else {
		_dgb := int(_a.RoundToEven(float64(_dcfa.Width) * _fag))
		_cffe := int(_a.RoundToEven(float64(_dcfa.Height) * _fag))
		_cdc, _fdfe = _d.NewImage(_dgb, _cffe, int(_dcfa.BitsPerComponent), _dcfa.ColorComponents, nil, nil, nil)
		if _fdfe != nil {
			return nil, _fdfe
		}
		_ff.CatmullRom.Scale(_cdc, _cdc.Bounds(), _ffbe, _ffbe.Bounds(), _ff.Over, &_ff.Options{})
	}
	_dfg := _cdc.Base()
	_dgbc := &_dc.Image{Width: int64(_dfg.Width), Height: int64(_dfg.Height), BitsPerComponent: int64(_dfg.BitsPerComponent), ColorComponents: _dfg.ColorComponents, Data: _dfg.Data}
	_dgbc.SetDecode(_dfg.Decode)
	_dgbc.SetAlpha(_dfg.Alpha)
	return _dgbc, nil
}

// CleanContentstream cleans up redundant operands in content streams, including Page and XObject Form
// contents. This process includes:
// 1. Marked content operators are removed.
// 2. Some operands are simplified (shorter form).
// TODO: Add more reduction methods and improving the methods for identifying unnecessary operands.
type CleanContentstream struct{}

func _eag(_gbe *_dc.XObjectImage, _bcbf imageModifications) error {
	_cgcf, _bafdg := _gbe.ToImage()
	if _bafdg != nil {
		return _bafdg
	}
	if _bcbf.Scale != 0 {
		_cgcf, _bafdg = _bagb(_cgcf, _bcbf.Scale)
		if _bafdg != nil {
			return _bafdg
		}
	}
	if _bcbf.Encoding != nil {
		_gbe.Filter = _bcbf.Encoding
	}
	_gbe.Decode = nil
	switch _dgd := _gbe.Filter.(type) {
	case *_cg.FlateEncoder:
		if _dgd.Predictor != 1 && _dgd.Predictor != 11 {
			_dgd.Predictor = 1
		}
	}
	if _bafdg = _gbe.SetImage(_cgcf, nil); _bafdg != nil {
		_f.Log.Debug("\u0045\u0072\u0072or\u0020\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0076", _bafdg)
		return _bafdg
	}
	_gbe.ToPdfObject()
	return nil
}

// CombineDuplicateDirectObjects combines duplicated direct objects by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateDirectObjects struct{}

// Optimize optimizes PDF objects to decrease PDF size.
func (_bcbe *CompressStreams) Optimize(objects []_cg.PdfObject) (_cad []_cg.PdfObject, _gbb error) {
	_cad = make([]_cg.PdfObject, len(objects))
	copy(_cad, objects)
	for _, _gead := range objects {
		_geg, _fdb := _cg.GetStream(_gead)
		if !_fdb {
			continue
		}
		if _fff := _geg.Get("\u0046\u0069\u006c\u0074\u0065\u0072"); _fff != nil {
			if _, _abd := _cg.GetName(_fff); _abd {
				continue
			}
			if _bgff, _fca := _cg.GetArray(_fff); _fca && _bgff.Len() > 0 {
				continue
			}
		}
		_ggd := _cg.NewFlateEncoder()
		var _ead []byte
		_ead, _gbb = _ggd.EncodeBytes(_geg.Stream)
		if _gbb != nil {
			return _cad, _gbb
		}
		_affg := _ggd.MakeStreamDict()
		if len(_ead)+len(_affg.WriteString()) < len(_geg.Stream) {
			_geg.Stream = _ead
			_geg.PdfObjectDictionary.Merge(_affg)
			_geg.PdfObjectDictionary.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _cg.MakeInteger(int64(len(_geg.Stream))))
		}
	}
	return _cad, nil
}

type imageInfo struct {
	BitsPerComponent int
	ColorComponents  int
	Width            int
	Height           int
	Stream           *_cg.PdfObjectStream
	PPI              float64
}

// Image optimizes images by rewrite images into JPEG format with quality equals to ImageQuality.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type Image struct{ ImageQuality int }

// CombineDuplicateStreams combines duplicated streams by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateStreams struct{}

func _ggfe(_dfa *_cg.PdfObjectStream, _def []rune, _fcc []_aa.GlyphIndex) error {
	_dfa, _cfb := _cg.GetStream(_dfa)
	if !_cfb {
		_f.Log.Debug("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u002d\u002d\u0020\u0041\u0042\u004f\u0052T\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006e\u0067")
		return _fa.New("\u0066\u006f\u006e\u0074fi\u006c\u0065\u0032\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_cfc, _dcb := _cg.DecodeStream(_dfa)
	if _dcb != nil {
		_f.Log.Debug("\u0044\u0065c\u006f\u0064\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _dcb)
		return _dcb
	}
	_gaa, _dcb := _aa.Parse(_ad.NewReader(_cfc))
	if _dcb != nil {
		_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0020f\u006f\u006e\u0074", len(_dfa.Stream))
		return _dcb
	}
	_gea := _fcc
	if len(_def) > 0 {
		_cfba := _gaa.LookupRunes(_def)
		_gea = append(_gea, _cfba...)
	}
	_gaa, _dcb = _gaa.SubsetKeepIndices(_gea)
	if _dcb != nil {
		_f.Log.Debug("\u0045R\u0052\u004f\u0052\u0020s\u0075\u0062\u0073\u0065\u0074t\u0069n\u0067 \u0066\u006f\u006e\u0074\u003a\u0020\u0025v", _dcb)
		return _dcb
	}
	var _ddb _ad.Buffer
	_dcb = _gaa.Write(&_ddb)
	if _dcb != nil {
		_f.Log.Debug("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _dcb)
		return _dcb
	}
	if _ddb.Len() > len(_cfc) {
		_f.Log.Debug("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070")
		return nil
	}
	_gbf, _dcb := _cg.MakeStream(_ddb.Bytes(), _cg.NewFlateEncoder())
	if _dcb != nil {
		_f.Log.Debug("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _dcb)
		return _dcb
	}
	*_dfa = *_gbf
	_dfa.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _cg.MakeInteger(int64(_ddb.Len())))
	return nil
}
func _dde(_bacg []_cg.PdfObject) objectStructure {
	_efgg := objectStructure{}
	_babee := false
	for _, _deg := range _bacg {
		switch _egecc := _deg.(type) {
		case *_cg.PdfIndirectObject:
			_bbea, _gfcc := _cg.GetDict(_egecc)
			if !_gfcc {
				continue
			}
			_fcbe, _gfcc := _cg.GetName(_bbea.Get("\u0054\u0079\u0070\u0065"))
			if !_gfcc {
				continue
			}
			switch _fcbe.String() {
			case "\u0043a\u0074\u0061\u006c\u006f\u0067":
				_efgg._gadec = _bbea
				_babee = true
			}
		}
		if _babee {
			break
		}
	}
	if !_babee {
		return _efgg
	}
	_dgdd, _cab := _cg.GetDict(_efgg._gadec.Get("\u0050\u0061\u0067e\u0073"))
	if !_cab {
		return _efgg
	}
	_efgg._dag = _dgdd
	_agg, _cab := _cg.GetArray(_dgdd.Get("\u004b\u0069\u0064\u0073"))
	if !_cab {
		return _efgg
	}
	for _, _ecf := range _agg.Elements() {
		_bfca, _cca := _cg.GetIndirect(_ecf)
		if !_cca {
			break
		}
		_efgg._gdff = append(_efgg._gdff, _bfca)
	}
	return _efgg
}

// Append appends optimizers to the chain.
func (_bd *Chain) Append(optimizers ..._dc.Optimizer) { _bd._ffb = append(_bd._ffb, optimizers...) }

// Optimize optimizes PDF objects to decrease PDF size.
func (_fccg *CombineIdenticalIndirectObjects) Optimize(objects []_cg.PdfObject) (_ddgd []_cg.PdfObject, _bda error) {
	_cbff(objects)
	_fcda := make(map[_cg.PdfObject]_cg.PdfObject)
	_fee := make(map[_cg.PdfObject]struct{})
	_gffg := make(map[string][]*_cg.PdfIndirectObject)
	for _, _decg := range objects {
		_edd, _age := _decg.(*_cg.PdfIndirectObject)
		if !_age {
			continue
		}
		if _bdb, _fefg := _edd.PdfObject.(*_cg.PdfObjectDictionary); _fefg {
			if _daec, _bae := _bdb.Get("\u0054\u0079\u0070\u0065").(*_cg.PdfObjectName); _bae && *_daec == "\u0050\u0061\u0067\u0065" {
				continue
			}
			_eeb := _b.New()
			_eeb.Write([]byte(_bdb.WriteString()))
			_gec := string(_eeb.Sum(nil))
			_gffg[_gec] = append(_gffg[_gec], _edd)
		}
	}
	for _, _bebd := range _gffg {
		if len(_bebd) < 2 {
			continue
		}
		_fec := _bebd[0]
		for _ccc := 1; _ccc < len(_bebd); _ccc++ {
			_eee := _bebd[_ccc]
			_fcda[_eee] = _fec
			_fee[_eee] = struct{}{}
		}
	}
	_ddgd = make([]_cg.PdfObject, 0, len(objects)-len(_fee))
	for _, _egb := range objects {
		if _, _afd := _fee[_egb]; _afd {
			continue
		}
		_ddgd = append(_ddgd, _egb)
	}
	_cacg(_ddgd, _fcda)
	return _ddgd, nil
}

// Chain allows to use sequence of optimizers.
// It implements interface model.Optimizer.
type Chain struct{ _ffb []_dc.Optimizer }
type objectStructure struct {
	_gadec *_cg.PdfObjectDictionary
	_dag   *_cg.PdfObjectDictionary
	_gdff  []*_cg.PdfIndirectObject
}

// New creates a optimizers chain from options.
func New(options Options) *Chain {
	_gedba := new(Chain)
	if options.CleanFonts || options.SubsetFonts {
		_gedba.Append(&CleanFonts{Subset: options.SubsetFonts})
	}
	if options.CleanContentstream {
		_gedba.Append(new(CleanContentstream))
	}
	if options.ImageUpperPPI > 0 {
		_ada := new(ImagePPI)
		_ada.ImageUpperPPI = options.ImageUpperPPI
		_gedba.Append(_ada)
	}
	if options.ImageQuality > 0 {
		_cdaa := new(Image)
		_cdaa.ImageQuality = options.ImageQuality
		_gedba.Append(_cdaa)
	}
	if options.CombineDuplicateDirectObjects {
		_gedba.Append(new(CombineDuplicateDirectObjects))
	}
	if options.CombineDuplicateStreams {
		_gedba.Append(new(CombineDuplicateStreams))
	}
	if options.CombineIdenticalIndirectObjects {
		_gedba.Append(new(CombineIdenticalIndirectObjects))
	}
	if options.UseObjectStreams {
		_gedba.Append(new(ObjectStreams))
	}
	if options.CompressStreams {
		_gedba.Append(new(CompressStreams))
	}
	return _gedba
}

// CombineIdenticalIndirectObjects combines identical indirect objects.
// It implements interface model.Optimizer.
type CombineIdenticalIndirectObjects struct{}
