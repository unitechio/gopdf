package optimize

import (
	_af "bytes"
	_be "crypto/md5"
	_cf "errors"
	_d "fmt"
	_ff "math"
	_bc "strings"

	_f "bitbucket.org/shenghui0779/gopdf/common"
	_fa "bitbucket.org/shenghui0779/gopdf/contentstream"
	_g "bitbucket.org/shenghui0779/gopdf/core"
	_e "bitbucket.org/shenghui0779/gopdf/extractor"
	_ce "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_a "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_ag "bitbucket.org/shenghui0779/gopdf/model"
	_fd "github.com/unidoc/unitype"
	_c "golang.org/x/image/draw"
)

func _agf(_dcfb string, _eaea []string) bool {
	for _, _ccg := range _eaea {
		if _dcfb == _ccg {
			return true
		}
	}
	return false
}
func _gaa(_bcd []_g.PdfObject) (_fde map[*_g.PdfObjectStream]struct{}, _deff error) {
	_fde = map[*_g.PdfObjectStream]struct{}{}
	_afgc := map[*_ag.PdfFont]struct{}{}
	_bgf := _gfab(_bcd)
	for _, _dee := range _bgf._cebg {
		_cb, _dff := _g.GetDict(_dee.PdfObject)
		if !_dff {
			continue
		}
		_abb, _dff := _g.GetDict(_cb.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
		if !_dff {
			continue
		}
		_gab, _ := _agdb(_cb.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		_fdc, _gca := _ag.NewPdfPageResourcesFromDict(_abb)
		if _gca != nil {
			return nil, _gca
		}
		_dda := []content{{_beae: _gab, _daee: _fdc}}
		_ea := _fgf(_cb.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _ea != nil {
			_dda = append(_dda, _ea...)
		}
		for _, _ddab := range _dda {
			_eaf, _ccff := _e.NewFromContents(_ddab._beae, _ddab._daee)
			if _ccff != nil {
				return nil, _ccff
			}
			_cgd, _, _, _ccff := _eaf.ExtractPageText()
			if _ccff != nil {
				return nil, _ccff
			}
			for _, _cef := range _cgd.Marks().Elements() {
				if _cef.Font == nil {
					continue
				}
				if _, _ccd := _afgc[_cef.Font]; !_ccd {
					_afgc[_cef.Font] = struct{}{}
				}
			}
		}
	}
	_dfd := map[*_g.PdfObjectStream][]*_ag.PdfFont{}
	for _gf := range _afgc {
		_ef := _gf.FontDescriptor()
		if _ef == nil || _ef.FontFile2 == nil {
			continue
		}
		_aaa, _dec := _g.GetStream(_ef.FontFile2)
		if !_dec {
			continue
		}
		_dfd[_aaa] = append(_dfd[_aaa], _gf)
	}
	for _abbb := range _dfd {
		var _aga []rune
		var _fgb []_fd.GlyphIndex
		for _, _fee := range _dfd[_abbb] {
			switch _bbc := _fee.Encoder().(type) {
			case *_a.IdentityEncoder:
				_bgfe := _bbc.RegisteredRunes()
				_bce := make([]_fd.GlyphIndex, len(_bgfe))
				for _dag, _bbf := range _bgfe {
					_bce[_dag] = _fd.GlyphIndex(_bbf)
				}
				_fgb = append(_fgb, _bce...)
			case *_a.TrueTypeFontEncoder:
				_ffa := _bbc.RegisteredRunes()
				_aga = append(_aga, _ffa...)
			case _a.SimpleEncoder:
				_dcd := _bbc.Charcodes()
				for _, _ead := range _dcd {
					_fgbb, _fcb := _bbc.CharcodeToRune(_ead)
					if !_fcb {
						_f.Log.Debug("\u0043\u0068a\u0072\u0063\u006f\u0064\u0065\u003c\u002d\u003e\u0072\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064: \u0025\u0064", _ead)
						continue
					}
					_aga = append(_aga, _fgbb)
				}
			}
		}
		_deff = _efe(_abbb, _aga, _fgb)
		if _deff != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006eg\u0020f\u006f\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _deff)
			return nil, _deff
		}
		_fde[_abbb] = struct{}{}
	}
	return _fde, nil
}
func _afg(_dg *_fa.ContentStreamOperations) *_fa.ContentStreamOperations {
	if _dg == nil {
		return nil
	}
	_ab := _fa.ContentStreamOperations{}
	for _, _ed := range *_dg {
		switch _ed.Operand {
		case "\u0042\u0044\u0043", "\u0042\u004d\u0043", "\u0045\u004d\u0043":
			continue
		case "\u0054\u006d":
			if len(_ed.Params) == 6 {
				if _gd, _acf := _g.GetNumbersAsFloat(_ed.Params); _acf == nil {
					if _gd[0] == 1 && _gd[1] == 0 && _gd[2] == 0 && _gd[3] == 1 {
						_ed = &_fa.ContentStreamOperation{Params: []_g.PdfObject{_ed.Params[4], _ed.Params[5]}, Operand: "\u0054\u0064"}
					}
				}
			}
		}
		_ab = append(_ab, _ed)
	}
	return &_ab
}

// CleanContentstream cleans up redundant operands in content streams, including Page and XObject Form
// contents. This process includes:
// 1. Marked content operators are removed.
// 2. Some operands are simplified (shorter form).
// TODO: Add more reduction methods and improving the methods for identifying unnecessary operands.
type CleanContentstream struct{}

func _agdb(_cgba _g.PdfObject) (_eeac string, _eeaf []_g.PdfObject) {
	var _becd _af.Buffer
	switch _gbef := _cgba.(type) {
	case *_g.PdfIndirectObject:
		_eeaf = append(_eeaf, _gbef)
		_cgba = _gbef.PdfObject
	}
	switch _fdec := _cgba.(type) {
	case *_g.PdfObjectStream:
		if _cbb, _cda := _g.DecodeStream(_fdec); _cda == nil {
			_becd.Write(_cbb)
			_eeaf = append(_eeaf, _fdec)
		}
	case *_g.PdfObjectArray:
		for _, _baca := range _fdec.Elements() {
			switch _gcdc := _baca.(type) {
			case *_g.PdfObjectStream:
				if _cbc, _edga := _g.DecodeStream(_gcdc); _edga == nil {
					_becd.Write(_cbc)
					_eeaf = append(_eeaf, _gcdc)
				}
			}
		}
	}
	return _becd.String(), _eeaf
}

type content struct {
	_beae string
	_daee *_ag.PdfPageResources
}

// Optimize implements Optimizer interface.
func (_agc *CleanUnusedResources) Optimize(objects []_g.PdfObject) (_ecge []_g.PdfObject, _gbga error) {
	_abc, _gbga := _ffe(objects)
	if _gbga != nil {
		return nil, _gbga
	}
	_aef := []_g.PdfObject{}
	for _, _gge := range objects {
		_, _abd := _abc[_gge]
		if _abd {
			continue
		}
		_aef = append(_aef, _gge)
	}
	return _aef, nil
}

// ImagePPI optimizes images by scaling images such that the PPI (pixels per inch) is never higher than ImageUpperPPI.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type ImagePPI struct{ ImageUpperPPI float64 }

func _ffe(_ffad []_g.PdfObject) (map[_g.PdfObject]struct{}, error) {
	_dbg := _gfab(_ffad)
	_dgad := _dbg._cebg
	_bcgb := make(map[_g.PdfObject]struct{})
	_abeg := _ddde(_dgad)
	for _, _bac := range _dgad {
		_acg, _abg := _g.GetDict(_bac.PdfObject)
		if !_abg {
			continue
		}
		_gdd, _abg := _g.GetDict(_acg.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
		if !_abg {
			continue
		}
		_aca := _abeg["\u0058O\u0062\u006a\u0065\u0063\u0074"]
		_cefa, _abg := _g.GetDict(_gdd.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
		if _abg {
			_cff := _bcbf(_cefa)
			for _, _gfed := range _cff {
				if _agf(_gfed, _aca) {
					continue
				}
				_add := *_g.MakeName(_gfed)
				_cae := _cefa.Get(_add)
				_bcgb[_cae] = struct{}{}
				_cefa.Remove(_add)
				_eddd := _bec(_cae, _bcgb)
				if _eddd != nil {
					_f.Log.Debug("\u0066\u0061\u0069\u006ce\u0064\u0020\u0074\u006f\u0020\u0074\u0072\u0061\u0076\u0065r\u0073e\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0025\u0076", _cae)
				}
			}
		}
		_gba, _abg := _g.GetDict(_gdd.Get("\u0046\u006f\u006e\u0074"))
		_eadg := _abeg["\u0046\u006f\u006e\u0074"]
		if _abg {
			_dad := _bcbf(_gba)
			for _, _gcd := range _dad {
				if _agf(_gcd, _eadg) {
					continue
				}
				_cee := *_g.MakeName(_gcd)
				_gcfe := _gba.Get(_cee)
				_bcgb[_gcfe] = struct{}{}
				_gba.Remove(_cee)
				_dbc := _bec(_gcfe, _bcgb)
				if _dbc != nil {
					_f.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074 %\u0076\u000a", _gcfe)
				}
			}
		}
		_aadaa, _abg := _g.GetDict(_gdd.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
		if _abg {
			_fgg := _bcbf(_aadaa)
			_fgga := _abeg["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"]
			for _, _acad := range _fgg {
				if _agf(_acad, _fgga) {
					continue
				}
				_eaed := *_g.MakeName(_acad)
				_ebc := _aadaa.Get(_eaed)
				_bcgb[_ebc] = struct{}{}
				_aadaa.Remove(_eaed)
				_dbb := _bec(_ebc, _bcgb)
				if _dbb != nil {
					_f.Log.Debug("\u0066\u0061i\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074 %\u0076\u000a", _ebc)
				}
			}
		}
	}
	return _bcgb, nil
}

// New creates a optimizers chain from options.
func New(options Options) *Chain {
	_ggfa := new(Chain)
	if options.CleanFonts || options.SubsetFonts {
		_ggfa.Append(&CleanFonts{Subset: options.SubsetFonts})
	}
	if options.CleanContentstream {
		_ggfa.Append(new(CleanContentstream))
	}
	if options.ImageUpperPPI > 0 {
		_abbg := new(ImagePPI)
		_abbg.ImageUpperPPI = options.ImageUpperPPI
		_ggfa.Append(_abbg)
	}
	if options.ImageQuality > 0 {
		_acbg := new(Image)
		_acbg.ImageQuality = options.ImageQuality
		_ggfa.Append(_acbg)
	}
	if options.CombineDuplicateDirectObjects {
		_ggfa.Append(new(CombineDuplicateDirectObjects))
	}
	if options.CombineDuplicateStreams {
		_ggfa.Append(new(CombineDuplicateStreams))
	}
	if options.CombineIdenticalIndirectObjects {
		_ggfa.Append(new(CombineIdenticalIndirectObjects))
	}
	if options.UseObjectStreams {
		_ggfa.Append(new(ObjectStreams))
	}
	if options.CompressStreams {
		_ggfa.Append(new(CompressStreams))
	}
	if options.CleanUnusedResources {
		_ggfa.Append(new(CleanUnusedResources))
	}
	return _ggfa
}

// CombineDuplicateStreams combines duplicated streams by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateStreams struct{}
type imageInfo struct {
	BitsPerComponent int
	ColorComponents  int
	Width            int
	Height           int
	Stream           *_g.PdfObjectStream
	PPI              float64
}

// CombineDuplicateDirectObjects combines duplicated direct objects by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateDirectObjects struct{}

// Optimize optimizes PDF objects to decrease PDF size.
func (_ecc *CombineDuplicateDirectObjects) Optimize(objects []_g.PdfObject) (_fdce []_g.PdfObject, _bccd error) {
	_gdc(objects)
	_aff := make(map[string][]*_g.PdfObjectDictionary)
	var _fea func(_gbe *_g.PdfObjectDictionary)
	_fea = func(_dfe *_g.PdfObjectDictionary) {
		for _, _bgge := range _dfe.Keys() {
			_gcce := _dfe.Get(_bgge)
			if _bcbfd, _cefb := _gcce.(*_g.PdfObjectDictionary); _cefb {
				_cfg := _be.New()
				_cfg.Write([]byte(_bcbfd.WriteString()))
				_gda := string(_cfg.Sum(nil))
				_aff[_gda] = append(_aff[_gda], _bcbfd)
				_fea(_bcbfd)
			}
		}
	}
	for _, _eac := range objects {
		_fcf, _cgg := _eac.(*_g.PdfIndirectObject)
		if !_cgg {
			continue
		}
		if _feb, _faa := _fcf.PdfObject.(*_g.PdfObjectDictionary); _faa {
			_fea(_feb)
		}
	}
	_fba := make([]_g.PdfObject, 0, len(_aff))
	_caf := make(map[_g.PdfObject]_g.PdfObject)
	for _, _dagd := range _aff {
		if len(_dagd) < 2 {
			continue
		}
		_eea := _g.MakeDict()
		_eea.Merge(_dagd[0])
		_aee := _g.MakeIndirectObject(_eea)
		_fba = append(_fba, _aee)
		for _cgga := 0; _cgga < len(_dagd); _cgga++ {
			_efc := _dagd[_cgga]
			_caf[_efc] = _aee
		}
	}
	_fdce = make([]_g.PdfObject, len(objects))
	copy(_fdce, objects)
	_fdce = append(_fba, _fdce...)
	_daf(_fdce, _caf)
	return _fdce, nil
}

type objectStructure struct {
	_bddf *_g.PdfObjectDictionary
	_eccc *_g.PdfObjectDictionary
	_cebg []*_g.PdfIndirectObject
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
	CleanUnusedResources            bool
}

// Chain allows to use sequence of optimizers.
// It implements interface model.Optimizer.
type Chain struct{ _dc []_ag.Optimizer }

// Optimize optimizes PDF objects to decrease PDF size.
func (_faf *Chain) Optimize(objects []_g.PdfObject) (_ge []_g.PdfObject, _ae error) {
	_ec := objects
	for _, _ac := range _faf._dc {
		_bd, _bg := _ac.Optimize(_ec)
		if _bg != nil {
			_f.Log.Debug("\u0045\u0052\u0052OR\u0020\u004f\u0070\u0074\u0069\u006d\u0069\u007a\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u002b\u0076", _bg)
			continue
		}
		_ec = _bd
	}
	return _ec, nil
}
func _aefc(_afbf *_ag.Image, _cagd float64) (*_ag.Image, error) {
	_ecfc, _eafe := _afbf.ToGoImage()
	if _eafe != nil {
		return nil, _eafe
	}
	var _cce _ce.Image
	_efcg, _gaef := _ecfc.(*_ce.Monochrome)
	if _gaef {
		if _eafe = _efcg.ResolveDecode(); _eafe != nil {
			return nil, _eafe
		}
		_cce, _eafe = _efcg.Scale(_cagd)
		if _eafe != nil {
			return nil, _eafe
		}
	} else {
		_gcdd := int(_ff.RoundToEven(float64(_afbf.Width) * _cagd))
		_caa := int(_ff.RoundToEven(float64(_afbf.Height) * _cagd))
		_cce, _eafe = _ce.NewImage(_gcdd, _caa, int(_afbf.BitsPerComponent), _afbf.ColorComponents, nil, nil, nil)
		if _eafe != nil {
			return nil, _eafe
		}
		_c.CatmullRom.Scale(_cce, _cce.Bounds(), _ecfc, _ecfc.Bounds(), _c.Over, &_c.Options{})
	}
	_edgd := _cce.Base()
	_gaba := &_ag.Image{Width: int64(_edgd.Width), Height: int64(_edgd.Height), BitsPerComponent: int64(_edgd.BitsPerComponent), ColorComponents: _edgd.ColorComponents, Data: _edgd.Data}
	_gaba.SetDecode(_edgd.Decode)
	_gaba.SetAlpha(_edgd.Alpha)
	return _gaba, nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_gaff *ImagePPI) Optimize(objects []_g.PdfObject) (_baf []_g.PdfObject, _abbbc error) {
	if _gaff.ImageUpperPPI <= 0 {
		return objects, nil
	}
	_ecd := _gbgg(objects)
	if len(_ecd) == 0 {
		return objects, nil
	}
	_adee := make(map[_g.PdfObject]struct{})
	for _, _bcec := range _ecd {
		_aea := _bcec.Stream.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b")
		_adee[_aea] = struct{}{}
	}
	_faef := make(map[*_g.PdfObjectStream]*imageInfo)
	for _, _ffdf := range _ecd {
		_faef[_ffdf.Stream] = _ffdf
	}
	var _gdad *_g.PdfObjectDictionary
	for _, _cad := range objects {
		if _bfda, _gga := _g.GetDict(_cad); _gdad == nil && _gga {
			if _efa, _ccda := _g.GetName(_bfda.Get("\u0054\u0079\u0070\u0065")); _ccda && *_efa == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
				_gdad = _bfda
			}
		}
	}
	if _gdad == nil {
		return objects, nil
	}
	_fed, _gbea := _g.GetDict(_gdad.Get("\u0050\u0061\u0067e\u0073"))
	if !_gbea {
		return objects, nil
	}
	_gfad, _bbcbg := _g.GetArray(_fed.Get("\u004b\u0069\u0064\u0073"))
	if !_bbcbg {
		return objects, nil
	}
	for _, _bfgg := range _gfad.Elements() {
		_feg := make(map[string]*imageInfo)
		_febd, _fdea := _g.GetDict(_bfgg)
		if !_fdea {
			continue
		}
		_ecb, _ := _agdb(_febd.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		if len(_ecb) == 0 {
			continue
		}
		_bab, _agg := _g.GetDict(_febd.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
		if !_agg {
			continue
		}
		_egf, _bdea := _ag.NewPdfPageResourcesFromDict(_bab)
		if _bdea != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020-\u0020\u0069\u0067\u006e\u006fr\u0069\u006eg\u003a\u0020\u0025\u0076", _bdea)
			continue
		}
		_ggee, _efca := _g.GetDict(_bab.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
		if !_efca {
			continue
		}
		_abbd := _ggee.Keys()
		for _, _bfbd := range _abbd {
			if _eceb, _gcbc := _g.GetStream(_ggee.Get(_bfbd)); _gcbc {
				if _facd, _eef := _faef[_eceb]; _eef {
					_feg[string(_bfbd)] = _facd
				}
			}
		}
		_aaeg := _fa.NewContentStreamParser(_ecb)
		_febg, _bdea := _aaeg.Parse()
		if _bdea != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _bdea)
			continue
		}
		_addd := _fa.NewContentStreamProcessor(*_febg)
		_addd.AddHandler(_fa.HandlerConditionEnumAllOperands, "", func(_dfg *_fa.ContentStreamOperation, _dfb _fa.GraphicsState, _ebcc *_ag.PdfPageResources) error {
			switch _dfg.Operand {
			case "\u0044\u006f":
				if len(_dfg.Params) != 1 {
					_f.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020w\u0069\u0074\u0068\u0020\u006c\u0065\u006e\u0028\u0070\u0061ra\u006d\u0073\u0029 \u0021=\u0020\u0031")
					return nil
				}
				_ggff, _abec := _g.GetName(_dfg.Params[0])
				if !_abec {
					_f.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f\u006e\u0020\u004e\u0061\u006d\u0065\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072")
					return nil
				}
				if _bfgc, _cdda := _feg[string(*_ggff)]; _cdda {
					_bgd := _dfb.CTM.ScalingFactorX()
					_feeb := _dfb.CTM.ScalingFactorY()
					_gfadf, _cecd := _bgd/72.0, _feeb/72.0
					_edf, _faefe := float64(_bfgc.Width)/_gfadf, float64(_bfgc.Height)/_cecd
					if _gfadf == 0 || _cecd == 0 {
						_edf = 72.0
						_faefe = 72.0
					}
					_bfgc.PPI = _ff.Max(_bfgc.PPI, _edf)
					_bfgc.PPI = _ff.Max(_bfgc.PPI, _faefe)
				}
			}
			return nil
		})
		_bdea = _addd.Process(_egf)
		if _bdea != nil {
			_f.Log.Debug("E\u0052\u0052\u004f\u0052 p\u0072o\u0063\u0065\u0073\u0073\u0069n\u0067\u003a\u0020\u0025\u002b\u0076", _bdea)
			continue
		}
	}
	for _, _baba := range _ecd {
		if _, _fedb := _adee[_baba.Stream]; _fedb {
			continue
		}
		if _baba.PPI <= _gaff.ImageUpperPPI {
			continue
		}
		_aadd, _aabfa := _ag.NewXObjectImageFromStream(_baba.Stream)
		if _aabfa != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _aabfa)
			continue
		}
		var _caea imageModifications
		_caea.Scale = _gaff.ImageUpperPPI / _baba.PPI
		if _baba.BitsPerComponent == 1 && _baba.ColorComponents == 1 {
			_bgcd := _ff.Round(_baba.PPI / _gaff.ImageUpperPPI)
			_fdf := _ce.NextPowerOf2(uint(_bgcd))
			if _ce.InDelta(float64(_fdf), 1/_caea.Scale, 0.3) {
				_caea.Scale = float64(1) / float64(_fdf)
			}
			if _, _eff := _aadd.Filter.(*_g.JBIG2Encoder); !_eff {
				_caea.Encoding = _g.NewJBIG2Encoder()
			}
		}
		if _aabfa = _fbc(_aadd, _caea); _aabfa != nil {
			_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072 \u0073\u0063\u0061\u006c\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u006be\u0065\u0070\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _aabfa)
			continue
		}
		_caea.Encoding = nil
		if _fagc, _gcbg := _g.GetStream(_baba.Stream.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b")); _gcbg {
			_cefc, _ded := _ag.NewXObjectImageFromStream(_fagc)
			if _ded != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _ded)
				continue
			}
			if _ded = _fbc(_cefc, _caea); _ded != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _ded)
				continue
			}
		}
	}
	return objects, nil
}
func _gbgg(_eddda []_g.PdfObject) []*imageInfo {
	_fcfb := _g.PdfObjectName("\u0053u\u0062\u0074\u0079\u0070\u0065")
	_efb := make(map[*_g.PdfObjectStream]struct{})
	var _baa []*imageInfo
	for _, _edde := range _eddda {
		_fgcb, _bfb := _g.GetStream(_edde)
		if !_bfb {
			continue
		}
		if _, _ceb := _efb[_fgcb]; _ceb {
			continue
		}
		_efb[_fgcb] = struct{}{}
		_afb := _fgcb.PdfObjectDictionary.Get(_fcfb)
		_eda, _bfb := _g.GetName(_afb)
		if !_bfb || string(*_eda) != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		_bbfe := &imageInfo{Stream: _fgcb, BitsPerComponent: 8}
		if _cbdb, _cffc := _g.GetIntVal(_fgcb.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _cffc {
			_bbfe.BitsPerComponent = _cbdb
		}
		if _fgbe, _adb := _g.GetIntVal(_fgcb.Get("\u0057\u0069\u0064t\u0068")); _adb {
			_bbfe.Width = _fgbe
		}
		if _edbg, _geda := _g.GetIntVal(_fgcb.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _geda {
			_bbfe.Height = _edbg
		}
		_ffb, _ggc := _ag.NewPdfColorspaceFromPdfObject(_fgcb.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"))
		if _ggc != nil {
			_f.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ggc)
			continue
		}
		if _ffb == nil {
			_dgff, _fgd := _g.GetName(_fgcb.Get("\u0046\u0069\u006c\u0074\u0065\u0072"))
			if _fgd {
				switch _dgff.String() {
				case "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065", "J\u0042\u0049\u0047\u0032\u0044\u0065\u0063\u006f\u0064\u0065":
					_ffb = _ag.NewPdfColorspaceDeviceGray()
					_bbfe.BitsPerComponent = 1
				}
			}
		}
		switch _caeb := _ffb.(type) {
		case *_ag.PdfColorspaceDeviceRGB:
			_bbfe.ColorComponents = 3
		case *_ag.PdfColorspaceDeviceGray:
			_bbfe.ColorComponents = 1
		default:
			_f.Log.Debug("\u004f\u0070\u0074\u0069\u006d\u0069\u007aa\u0074\u0069\u006fn\u0020\u0069\u0073 \u006e\u006ft\u0020\u0073\u0075\u0070\u0070\u006fr\u0074ed\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u006b\u0069\u0070", _caeb)
			continue
		}
		_baa = append(_baa, _bbfe)
	}
	return _baa
}
func _fgf(_aab _g.PdfObject) []content {
	if _aab == nil {
		return nil
	}
	_abe, _ecg := _g.GetArray(_aab)
	if !_ecg {
		_f.Log.Debug("\u0041\u006e\u006e\u006fts\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
		return nil
	}
	var _ece []content
	for _, _bge := range _abe.Elements() {
		_eae, _gcf := _g.GetDict(_bge)
		if !_gcf {
			_f.Log.Debug("I\u0067\u006e\u006f\u0072\u0069\u006eg\u0020\u006e\u006f\u006e\u002d\u0064i\u0063\u0074\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006e\u0020\u0041\u006e\u006e\u006ft\u0073")
			continue
		}
		_fdbc, _gcf := _g.GetDict(_eae.Get("\u0041\u0050"))
		if !_gcf {
			_f.Log.Debug("\u004e\u006f\u0020\u0041P \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067")
			continue
		}
		_bff := _g.TraceToDirectObject(_fdbc.Get("\u004e"))
		if _bff == nil {
			_f.Log.Debug("N\u006f\u0020\u004e\u0020en\u0074r\u0079\u0020\u002d\u0020\u0073k\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		var _eadc *_g.PdfObjectStream
		switch _fbb := _bff.(type) {
		case *_g.PdfObjectDictionary:
			_edg, _aabf := _g.GetName(_eae.Get("\u0041\u0053"))
			if !_aabf {
				_f.Log.Debug("\u004e\u006f\u0020\u0041S \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067")
				continue
			}
			_eadc, _aabf = _g.GetStream(_fbb.Get(*_edg))
			if !_aabf {
				_f.Log.Debug("\u0046o\u0072\u006d\u0020\u006eo\u0074\u0020\u0066\u006f\u0075n\u0064 \u002d \u0073\u006b\u0069\u0070\u0070\u0069\u006eg")
				continue
			}
		case *_g.PdfObjectStream:
			_eadc = _fbb
		}
		if _eadc == nil {
			_f.Log.Debug("\u0046\u006f\u0072m\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0028n\u0069\u006c\u0029\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		_afgb, _ddd := _ag.NewXObjectFormFromStream(_eadc)
		if _ddd != nil {
			_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _ddd)
			continue
		}
		_cdd, _ddd := _afgb.GetContentStream()
		if _ddd != nil {
			_f.Log.Debug("E\u0072\u0072\u006f\u0072\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0063\u006fn\u0074\u0065\u006et\u0073:\u0020\u0025\u0076", _ddd)
			continue
		}
		_ece = append(_ece, content{_beae: string(_cdd), _daee: _afgb.Resources})
	}
	return _ece
}

// Image optimizes images by rewrite images into JPEG format with quality equals to ImageQuality.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type Image struct{ ImageQuality int }

func _gdc(_ffcbb []_g.PdfObject) {
	for _cafbc, _caba := range _ffcbb {
		switch _babe := _caba.(type) {
		case *_g.PdfIndirectObject:
			_babe.ObjectNumber = int64(_cafbc + 1)
			_babe.GenerationNumber = 0
		case *_g.PdfObjectStream:
			_babe.ObjectNumber = int64(_cafbc + 1)
			_babe.GenerationNumber = 0
		case *_g.PdfObjectStreams:
			_babe.ObjectNumber = int64(_cafbc + 1)
			_babe.GenerationNumber = 0
		}
	}
}
func _efe(_dfde *_g.PdfObjectStream, _bdc []rune, _gcb []_fd.GlyphIndex) error {
	_dfde, _cab := _g.GetStream(_dfde)
	if !_cab {
		_f.Log.Debug("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u002d\u002d\u0020\u0041\u0042\u004f\u0052T\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006e\u0067")
		return _cf.New("\u0066\u006f\u006e\u0074fi\u006c\u0065\u0032\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_efd, _bccg := _g.DecodeStream(_dfde)
	if _bccg != nil {
		_f.Log.Debug("\u0044\u0065c\u006f\u0064\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _bccg)
		return _bccg
	}
	_cec, _bccg := _fd.Parse(_af.NewReader(_efd))
	if _bccg != nil {
		_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0020f\u006f\u006e\u0074", len(_dfde.Stream))
		return _bccg
	}
	_gaf := _gcb
	if len(_bdc) > 0 {
		_cd := _cec.LookupRunes(_bdc)
		_gaf = append(_gaf, _cd...)
	}
	_cec, _bccg = _cec.SubsetKeepIndices(_gaf)
	if _bccg != nil {
		_f.Log.Debug("\u0045R\u0052\u004f\u0052\u0020s\u0075\u0062\u0073\u0065\u0074t\u0069n\u0067 \u0066\u006f\u006e\u0074\u003a\u0020\u0025v", _bccg)
		return _bccg
	}
	var _cba _af.Buffer
	_bccg = _cec.Write(&_cba)
	if _bccg != nil {
		_f.Log.Debug("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _bccg)
		return _bccg
	}
	if _cba.Len() > len(_efd) {
		_f.Log.Debug("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070")
		return nil
	}
	_bfd, _bccg := _g.MakeStream(_cba.Bytes(), _g.NewFlateEncoder())
	if _bccg != nil {
		_f.Log.Debug("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _bccg)
		return _bccg
	}
	*_dfde = *_bfd
	_dfde.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _g.MakeInteger(int64(_cba.Len())))
	return nil
}
func _gfab(_dbe []_g.PdfObject) objectStructure {
	_beag := objectStructure{}
	_ccgc := false
	for _, _baag := range _dbe {
		switch _dde := _baag.(type) {
		case *_g.PdfIndirectObject:
			_gefg, _ffg := _g.GetDict(_dde)
			if !_ffg {
				continue
			}
			_ddaag, _ffg := _g.GetName(_gefg.Get("\u0054\u0079\u0070\u0065"))
			if !_ffg {
				continue
			}
			switch _ddaag.String() {
			case "\u0043a\u0074\u0061\u006c\u006f\u0067":
				_beag._bddf = _gefg
				_ccgc = true
			}
		}
		if _ccgc {
			break
		}
	}
	if !_ccgc {
		return _beag
	}
	_baagg, _ecca := _g.GetDict(_beag._bddf.Get("\u0050\u0061\u0067e\u0073"))
	if !_ecca {
		return _beag
	}
	_beag._eccc = _baagg
	_bcdf, _ecca := _g.GetArray(_baagg.Get("\u004b\u0069\u0064\u0073"))
	if !_ecca {
		return _beag
	}
	for _, _bfgd := range _bcdf.Elements() {
		_fabf, _dbcf := _g.GetIndirect(_bfgd)
		if !_dbcf {
			break
		}
		_beag._cebg = append(_beag._cebg, _fabf)
	}
	return _beag
}

// CleanUnusedResources represents an optimizer used to clean unused resources.
type CleanUnusedResources struct{}

func _bec(_gddc _g.PdfObject, _fgc map[_g.PdfObject]struct{}) error {
	if _gag, _dgcd := _gddc.(*_g.PdfIndirectObject); _dgcd {
		_fgc[_gddc] = struct{}{}
		_bca := _bec(_gag.PdfObject, _fgc)
		if _bca != nil {
			return _bca
		}
		return nil
	}
	if _bag, _bfee := _gddc.(*_g.PdfObjectStream); _bfee {
		_fgc[_bag] = struct{}{}
		_bed := _bec(_bag.PdfObjectDictionary, _fgc)
		if _bed != nil {
			return _bed
		}
		return nil
	}
	if _gabg, _gefe := _gddc.(*_g.PdfObjectDictionary); _gefe {
		for _, _aefe := range _gabg.Keys() {
			_bgg := _gabg.Get(_aefe)
			_ = _bgg
			if _dfdg, _gcc := _bgg.(*_g.PdfObjectReference); _gcc {
				_bgg = _dfdg.Resolve()
				_gabg.Set(_aefe, _bgg)
			}
			if _aefe != "\u0050\u0061\u0072\u0065\u006e\u0074" {
				if _gfa := _bec(_bgg, _fgc); _gfa != nil {
					return _gfa
				}
			}
		}
		return nil
	}
	if _dgfd, _ggg := _gddc.(*_g.PdfObjectArray); _ggg {
		if _dgfd == nil {
			return _cf.New("\u0061\u0072\u0072a\u0079\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		}
		for _gae, _gff := range _dgfd.Elements() {
			if _ced, _abga := _gff.(*_g.PdfObjectReference); _abga {
				_gff = _ced.Resolve()
				_dgfd.Set(_gae, _gff)
			}
			if _dbce := _bec(_gff, _fgc); _dbce != nil {
				return _dbce
			}
		}
		return nil
	}
	return nil
}

// GetOptimizers gets the list of optimizers in chain `c`.
func (_gc *Chain) GetOptimizers() []_ag.Optimizer { return _gc._dc }

type imageModifications struct {
	Scale    float64
	Encoding _g.StreamEncoder
}

func _aa(_bcc *_g.PdfObjectStream) error {
	_deg, _ca := _g.DecodeStream(_bcc)
	if _ca != nil {
		return _ca
	}
	_fe := _fa.NewContentStreamParser(string(_deg))
	_def, _ca := _fe.Parse()
	if _ca != nil {
		return _ca
	}
	_def = _afg(_def)
	_dd := _def.Bytes()
	if len(_dd) >= len(_deg) {
		return nil
	}
	_cfc, _ca := _g.MakeStream(_def.Bytes(), _g.NewFlateEncoder())
	if _ca != nil {
		return _ca
	}
	_bcc.Stream = _cfc.Stream
	_bcc.Merge(_cfc.PdfObjectDictionary)
	return nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_fcd *CombineIdenticalIndirectObjects) Optimize(objects []_g.PdfObject) (_dbd []_g.PdfObject, _dcec error) {
	_gdc(objects)
	_bfg := make(map[_g.PdfObject]_g.PdfObject)
	_ddaa := make(map[_g.PdfObject]struct{})
	_bdd := make(map[string][]*_g.PdfIndirectObject)
	for _, _cgda := range objects {
		_bad, _agbf := _cgda.(*_g.PdfIndirectObject)
		if !_agbf {
			continue
		}
		if _bcdb, _ggf := _bad.PdfObject.(*_g.PdfObjectDictionary); _ggf {
			if _deea, _abef := _bcdb.Get("\u0054\u0079\u0070\u0065").(*_g.PdfObjectName); _abef && *_deea == "\u0050\u0061\u0067\u0065" {
				continue
			}
			_ffcb := _be.New()
			_ffcb.Write([]byte(_bcdb.WriteString()))
			_efg := string(_ffcb.Sum(nil))
			_bdd[_efg] = append(_bdd[_efg], _bad)
		}
	}
	for _, _dcb := range _bdd {
		if len(_dcb) < 2 {
			continue
		}
		_cdf := _dcb[0]
		for _dfeb := 1; _dfeb < len(_dcb); _dfeb++ {
			_cgf := _dcb[_dfeb]
			_bfg[_cgf] = _cdf
			_ddaa[_cgf] = struct{}{}
		}
	}
	_dbd = make([]_g.PdfObject, 0, len(objects)-len(_ddaa))
	for _, _caeg := range objects {
		if _, _bcdc := _ddaa[_caeg]; _bcdc {
			continue
		}
		_dbd = append(_dbd, _caeg)
	}
	_daf(_dbd, _bfg)
	return _dbd, nil
}

// CombineIdenticalIndirectObjects combines identical indirect objects.
// It implements interface model.Optimizer.
type CombineIdenticalIndirectObjects struct{}

// Optimize optimizes PDF objects to decrease PDF size.
func (_fae *Image) Optimize(objects []_g.PdfObject) (_ebff []_g.PdfObject, _gaedf error) {
	if _fae.ImageQuality <= 0 {
		return objects, nil
	}
	_gbgd := _gbgg(objects)
	if len(_gbgd) == 0 {
		return objects, nil
	}
	_cddb := make(map[_g.PdfObject]_g.PdfObject)
	_abda := make(map[_g.PdfObject]struct{})
	for _, _cecf := range _gbgd {
		_gcfed := _cecf.Stream.Get("\u0053\u004d\u0061s\u006b")
		_abda[_gcfed] = struct{}{}
	}
	for _gfg, _bcde := range _gbgd {
		_defa := _bcde.Stream
		if _, _fcdg := _abda[_defa]; _fcdg {
			continue
		}
		_fce, _aaf := _ag.NewXObjectImageFromStream(_defa)
		if _aaf != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _aaf)
			continue
		}
		switch _fce.Filter.(type) {
		case *_g.JBIG2Encoder:
			continue
		case *_g.CCITTFaxEncoder:
			continue
		}
		_edc, _aaf := _fce.ToImage()
		if _aaf != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _aaf)
			continue
		}
		_dgac := _g.NewDCTEncoder()
		_dgac.ColorComponents = _edc.ColorComponents
		_dgac.Quality = _fae.ImageQuality
		_dgac.BitsPerComponent = _bcde.BitsPerComponent
		_dgac.Width = _bcde.Width
		_dgac.Height = _bcde.Height
		_dddc, _aaf := _dgac.EncodeBytes(_edc.Data)
		if _aaf != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _aaf)
			continue
		}
		var _gad _g.StreamEncoder
		_gad = _dgac
		{
			_ebfe := _g.NewFlateEncoder()
			_acgc := _g.NewMultiEncoder()
			_acgc.AddEncoder(_ebfe)
			_acgc.AddEncoder(_dgac)
			_cdc, _fdcea := _acgc.EncodeBytes(_edc.Data)
			if _fdcea != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _fdcea)
				continue
			}
			if len(_cdc) < len(_dddc) {
				_f.Log.Trace("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063\u0020\u0069\u006d\u0070\u0072\u006f\u0076\u0065\u0073\u003a\u0020\u0025\u0064\u0020\u0074o\u0020\u0025\u0064\u0020\u0028o\u0072\u0069g\u0020\u0025\u0064\u0029", len(_dddc), len(_cdc), len(_defa.Stream))
				_dddc = _cdc
				_gad = _acgc
			}
		}
		_gbc := len(_defa.Stream)
		if _gbc < len(_dddc) {
			continue
		}
		_ecf := &_g.PdfObjectStream{Stream: _dddc}
		_ecf.PdfObjectReference = _defa.PdfObjectReference
		_ecf.PdfObjectDictionary = _g.MakeDict()
		_ecf.Merge(_defa.PdfObjectDictionary)
		_ecf.Merge(_gad.MakeStreamDict())
		_ecf.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _g.MakeInteger(int64(len(_dddc))))
		_cddb[_defa] = _ecf
		_gbgd[_gfg].Stream = _ecf
	}
	_ebff = make([]_g.PdfObject, len(objects))
	copy(_ebff, objects)
	_daf(_ebff, _cddb)
	return _ebff, nil
}
func _ddde(_dgc []*_g.PdfIndirectObject) map[string][]string {
	_afgf := map[string][]string{}
	for _, _aed := range _dgc {
		_cffe, _ffaf := _g.GetDict(_aed.PdfObject)
		if !_ffaf {
			continue
		}
		_beb := _cffe.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073")
		_fgbf := _g.TraceToDirectObject(_beb)
		var _gafe []string
		if _bcce, _cfdf := _fgbf.(*_g.PdfObjectArray); _cfdf {
			for _, _eba := range _bcce.Elements() {
				_dce, _eed := _dcff(_eba)
				if _eed != nil {
					continue
				}
				_gafe = append(_gafe, _dce)
			}
		}
		_gbb := _bc.Join(_gafe, "\u0020")
		_gbac := _fa.NewContentStreamParser(_gbb)
		_efdf, _bbcb := _gbac.Parse()
		if _bbcb != nil {
			continue
		}
		for _, _fgba := range *_efdf {
			_deca := _fgba.Operand
			_cag := _fgba.Params
			switch _deca {
			case "\u0044\u006f":
				_bbd := _cag[0].String()
				if _, _gde := _afgf["\u0058O\u0062\u006a\u0065\u0063\u0074"]; !_gde {
					_afgf["\u0058O\u0062\u006a\u0065\u0063\u0074"] = []string{_bbd}
				} else {
					_afgf["\u0058O\u0062\u006a\u0065\u0063\u0074"] = append(_afgf["\u0058O\u0062\u006a\u0065\u0063\u0074"], _bbd)
				}
			case "\u0054\u0066":
				_ecea := _cag[0].String()
				if _, _cgb := _afgf["\u0046\u006f\u006e\u0074"]; !_cgb {
					_afgf["\u0046\u006f\u006e\u0074"] = []string{_ecea}
				} else {
					_afgf["\u0046\u006f\u006e\u0074"] = append(_afgf["\u0046\u006f\u006e\u0074"], _ecea)
				}
			case "\u0067\u0073":
				_bbdg := _cag[0].String()
				if _, _bged := _afgf["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"]; !_bged {
					_afgf["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"] = []string{_bbdg}
				} else {
					_afgf["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"] = append(_afgf["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"], _bbdg)
				}
			}
		}
	}
	return _afgf
}

// ObjectStreams groups PDF objects to object streams.
// It implements interface model.Optimizer.
type ObjectStreams struct{}

// Append appends optimizers to the chain.
func (_de *Chain) Append(optimizers ..._ag.Optimizer) { _de._dc = append(_de._dc, optimizers...) }

// Optimize optimizes PDF objects to decrease PDF size.
func (_aba *CleanContentstream) Optimize(objects []_g.PdfObject) (_cfe []_g.PdfObject, _fc error) {
	_fag := map[*_g.PdfObjectStream]struct{}{}
	var _eb []*_g.PdfObjectStream
	_ad := func(_gbf *_g.PdfObjectStream) {
		if _, _age := _fag[_gbf]; !_age {
			_fag[_gbf] = struct{}{}
			_eb = append(_eb, _gbf)
		}
	}
	_bea := map[_g.PdfObject]bool{}
	_aad := map[_g.PdfObject]bool{}
	for _, _bb := range objects {
		switch _beg := _bb.(type) {
		case *_g.PdfIndirectObject:
			switch _fac := _beg.PdfObject.(type) {
			case *_g.PdfObjectDictionary:
				if _aadc, _aec := _g.GetName(_fac.Get("\u0054\u0079\u0070\u0065")); !_aec || _aadc.String() != "\u0050\u0061\u0067\u0065" {
					continue
				}
				if _eca, _gbff := _g.GetStream(_fac.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073")); _gbff {
					_ad(_eca)
				} else if _ebf, _agb := _g.GetArray(_fac.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073")); _agb {
					var _begg []*_g.PdfObjectStream
					for _, _da := range _ebf.Elements() {
						if _fab, _ddf := _g.GetStream(_da); _ddf {
							_begg = append(_begg, _fab)
						}
					}
					if len(_begg) > 0 {
						var _dae _af.Buffer
						for _, _defd := range _begg {
							if _aada, _cg := _g.DecodeStream(_defd); _cg == nil {
								_dae.Write(_aada)
							}
							_bea[_defd] = true
						}
						_fdb, _bgc := _g.MakeStream(_dae.Bytes(), _g.NewFlateEncoder())
						if _bgc != nil {
							return nil, _bgc
						}
						_aad[_fdb] = true
						_fac.Set("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _fdb)
						_ad(_fdb)
					}
				}
			}
		case *_g.PdfObjectStream:
			if _ccf, _ga := _g.GetName(_beg.Get("\u0054\u0079\u0070\u0065")); !_ga || _ccf.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
				continue
			}
			if _bde, _deb := _g.GetName(_beg.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); !_deb || _bde.String() != "\u0046\u006f\u0072\u006d" {
				continue
			}
			_ad(_beg)
		}
	}
	for _, _bgb := range _eb {
		_fc = _aa(_bgb)
		if _fc != nil {
			return nil, _fc
		}
	}
	_cfe = nil
	for _, _bf := range objects {
		if _bea[_bf] {
			continue
		}
		_cfe = append(_cfe, _bf)
	}
	for _df := range _aad {
		_cfe = append(_cfe, _df)
	}
	return _cfe, nil
}
func _fbc(_dgeb *_ag.XObjectImage, _bcab imageModifications) error {
	_eg, _gagd := _dgeb.ToImage()
	if _gagd != nil {
		return _gagd
	}
	if _bcab.Scale != 0 {
		_eg, _gagd = _aefc(_eg, _bcab.Scale)
		if _gagd != nil {
			return _gagd
		}
	}
	if _bcab.Encoding != nil {
		_dgeb.Filter = _bcab.Encoding
	}
	_dgeb.Decode = nil
	switch _bbb := _dgeb.Filter.(type) {
	case *_g.FlateEncoder:
		if _bbb.Predictor != 1 && _bbb.Predictor != 11 {
			_bbb.Predictor = 1
		}
	}
	if _gagd = _dgeb.SetImage(_eg, nil); _gagd != nil {
		_f.Log.Debug("\u0045\u0072\u0072or\u0020\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0076", _gagd)
		return _gagd
	}
	_dgeb.ToPdfObject()
	return nil
}
func _dcff(_adf _g.PdfObject) (string, error) {
	_bffa := _g.TraceToDirectObject(_adf)
	switch _gbfg := _bffa.(type) {
	case *_g.PdfObjectString:
		return _gbfg.Str(), nil
	case *_g.PdfObjectStream:
		_fcg, _dcg := _g.DecodeStream(_gbfg)
		if _dcg != nil {
			return "", _dcg
		}
		return string(_fcg), nil
	}
	return "", _d.Errorf("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072e\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0068\u006f\u006c\u0064\u0065\u0072\u0020\u0028\u0025\u0054\u0029", _bffa)
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_dgd *CompressStreams) Optimize(objects []_g.PdfObject) (_beca []_g.PdfObject, _bagf error) {
	_beca = make([]_g.PdfObject, len(objects))
	copy(_beca, objects)
	for _, _gdbe := range objects {
		_acaf, _abdd := _g.GetStream(_gdbe)
		if !_abdd {
			continue
		}
		if _eabd := _acaf.Get("\u0046\u0069\u006c\u0074\u0065\u0072"); _eabd != nil {
			if _, _dea := _g.GetName(_eabd); _dea {
				continue
			}
			if _dabg, _fge := _g.GetArray(_eabd); _fge && _dabg.Len() > 0 {
				continue
			}
		}
		_fcdf := _g.NewFlateEncoder()
		var _dfa []byte
		_dfa, _bagf = _fcdf.EncodeBytes(_acaf.Stream)
		if _bagf != nil {
			return _beca, _bagf
		}
		_gfb := _fcdf.MakeStreamDict()
		if len(_dfa)+len(_gfb.WriteString()) < len(_acaf.Stream) {
			_acaf.Stream = _dfa
			_acaf.PdfObjectDictionary.Merge(_gfb)
			_acaf.PdfObjectDictionary.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _g.MakeInteger(int64(len(_acaf.Stream))))
		}
	}
	return _beca, nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_babd *ObjectStreams) Optimize(objects []_g.PdfObject) (_dedg []_g.PdfObject, _bbe error) {
	_ebfb := &_g.PdfObjectStreams{}
	_acb := make([]_g.PdfObject, 0, len(objects))
	for _, _aeg := range objects {
		if _fgfa, _gccef := _aeg.(*_g.PdfIndirectObject); _gccef && _fgfa.GenerationNumber == 0 {
			_ebfb.Append(_aeg)
		} else {
			_acb = append(_acb, _aeg)
		}
	}
	if _ebfb.Len() == 0 {
		return _acb, nil
	}
	_dedg = make([]_g.PdfObject, 0, len(_acb)+_ebfb.Len()+1)
	if _ebfb.Len() > 1 {
		_dedg = append(_dedg, _ebfb)
	}
	_dedg = append(_dedg, _ebfb.Elements()...)
	_dedg = append(_dedg, _acb...)
	return _dedg, nil
}
func _daf(_fef []_g.PdfObject, _gbab map[_g.PdfObject]_g.PdfObject) {
	if len(_gbab) == 0 {
		return
	}
	for _cafb, _edfc := range _fef {
		if _eedd, _egg := _gbab[_edfc]; _egg {
			_fef[_cafb] = _eedd
			continue
		}
		_gbab[_edfc] = _edfc
		switch _ggae := _edfc.(type) {
		case *_g.PdfObjectArray:
			_febb := make([]_g.PdfObject, _ggae.Len())
			copy(_febb, _ggae.Elements())
			_daf(_febb, _gbab)
			for _gfeb, _afca := range _febb {
				_ggae.Set(_gfeb, _afca)
			}
		case *_g.PdfObjectStreams:
			_daf(_ggae.Elements(), _gbab)
		case *_g.PdfObjectStream:
			_bgfef := []_g.PdfObject{_ggae.PdfObjectDictionary}
			_daf(_bgfef, _gbab)
			_ggae.PdfObjectDictionary = _bgfef[0].(*_g.PdfObjectDictionary)
		case *_g.PdfObjectDictionary:
			_dabc := _ggae.Keys()
			_egfd := make([]_g.PdfObject, len(_dabc))
			for _bfgca, _gccea := range _dabc {
				_egfd[_bfgca] = _ggae.Get(_gccea)
			}
			_daf(_egfd, _gbab)
			for _cafe, _fdfe := range _dabc {
				_ggae.Set(_fdfe, _egfd[_cafe])
			}
		case *_g.PdfIndirectObject:
			_cbe := []_g.PdfObject{_ggae.PdfObject}
			_daf(_cbe, _gbab)
			_ggae.PdfObject = _cbe[0]
		}
	}
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_ccba *CombineDuplicateStreams) Optimize(objects []_g.PdfObject) (_efeg []_g.PdfObject, _aae error) {
	_cde := make(map[_g.PdfObject]_g.PdfObject)
	_dcfc := make(map[_g.PdfObject]struct{})
	_ffd := make(map[string][]*_g.PdfObjectStream)
	for _, _dgca := range objects {
		if _acge, _afff := _dgca.(*_g.PdfObjectStream); _afff {
			_ged := _be.New()
			_ged.Write(_acge.Stream)
			_ged.Write([]byte(_acge.PdfObjectDictionary.WriteString()))
			_debe := string(_ged.Sum(nil))
			_ffd[_debe] = append(_ffd[_debe], _acge)
		}
	}
	for _, _ade := range _ffd {
		if len(_ade) < 2 {
			continue
		}
		_fad := _ade[0]
		for _adc := 1; _adc < len(_ade); _adc++ {
			_fbba := _ade[_adc]
			_cde[_fbba] = _fad
			_dcfc[_fbba] = struct{}{}
		}
	}
	_efeg = make([]_g.PdfObject, 0, len(objects)-len(_dcfc))
	for _, _dab := range objects {
		if _, _cea := _dcfc[_dab]; _cea {
			continue
		}
		_efeg = append(_efeg, _dab)
	}
	_daf(_efeg, _cde)
	return _efeg, nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_ba *CleanFonts) Optimize(objects []_g.PdfObject) (_afa []_g.PdfObject, _ddg error) {
	var _ccb map[*_g.PdfObjectStream]struct{}
	if _ba.Subset {
		var _gaac error
		_ccb, _gaac = _gaa(objects)
		if _gaac != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0073u\u0062s\u0065\u0074\u0074\u0069\u006e\u0067\u003a \u0025\u0076", _gaac)
			return nil, _gaac
		}
	}
	for _, _cbf := range objects {
		_afc, _aag := _g.GetStream(_cbf)
		if !_aag {
			continue
		}
		if _, _fb := _ccb[_afc]; _fb {
			continue
		}
		_gcaf, _gdb := _g.NewEncoderFromStream(_afc)
		if _gdb != nil {
			_f.Log.Debug("\u0045\u0052RO\u0052\u0020\u0067e\u0074\u0074\u0069\u006eg e\u006eco\u0064\u0065\u0072\u003a\u0020\u0025\u0076 -\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067", _gdb)
			continue
		}
		_dga, _gdb := _gcaf.DecodeStream(_afc)
		if _gdb != nil {
			_f.Log.Debug("\u0044\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065r\u0072\u006f\u0072\u0020\u003a\u0020\u0025v\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067", _gdb)
			continue
		}
		if len(_dga) < 4 {
			continue
		}
		_gbg := string(_dga[:4])
		if _gbg == "\u004f\u0054\u0054\u004f" {
			continue
		}
		if _gbg != "\u0000\u0001\u0000\u0000" && _gbg != "\u0074\u0072\u0075\u0065" {
			continue
		}
		_gfe, _gdb := _fd.Parse(_af.NewReader(_dga))
		if _gdb != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020P\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _gdb)
			continue
		}
		_gdb = _gfe.Optimize()
		if _gdb != nil {
			_f.Log.Debug("\u0045\u0052RO\u0052\u0020\u004fp\u0074\u0069\u006d\u0069zin\u0067 f\u006f\u006e\u0074\u003a\u0020\u0025\u0076 -\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067", _gdb)
			continue
		}
		var _dgf _af.Buffer
		_gdb = _gfe.Write(&_dgf)
		if _gdb != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020W\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _gdb)
			continue
		}
		if _dgf.Len() > len(_dga) {
			_f.Log.Debug("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070")
			continue
		}
		_cdb, _gdb := _g.MakeStream(_dgf.Bytes(), _g.NewFlateEncoder())
		if _gdb != nil {
			continue
		}
		*_afc = *_cdb
		_afc.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _g.MakeInteger(int64(_dgf.Len())))
	}
	return objects, nil
}

// CompressStreams compresses uncompressed streams.
// It implements interface model.Optimizer.
type CompressStreams struct{}

// CleanFonts cleans up embedded fonts, reducing font sizes.
type CleanFonts struct {

	// Subset embedded fonts if encountered (if true).
	// Otherwise attempts to reduce the font program.
	Subset bool
}

func _bcbf(_ebcf *_g.PdfObjectDictionary) []string {
	_gddg := []string{}
	for _, _bcf := range _ebcf.Keys() {
		_gddg = append(_gddg, _bcf.String())
	}
	return _gddg
}
