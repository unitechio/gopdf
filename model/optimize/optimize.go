package optimize

import (
	_bg "bytes"
	_afb "crypto/md5"
	_c "errors"
	_af "math"

	_f "bitbucket.org/shenghui0779/gopdf/common"
	_bb "bitbucket.org/shenghui0779/gopdf/contentstream"
	_cg "bitbucket.org/shenghui0779/gopdf/core"
	_ab "bitbucket.org/shenghui0779/gopdf/extractor"
	_aa "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_ac "bitbucket.org/shenghui0779/gopdf/internal/textencoding"
	_e "bitbucket.org/shenghui0779/gopdf/model"
	_b "github.com/unidoc/unitype"
	_g "golang.org/x/image/draw"
)

// CompressStreams compresses uncompressed streams.
// It implements interface model.Optimizer.
type CompressStreams struct{}

// Optimize optimizes PDF objects to decrease PDF size.
func (_dee *ImagePPI) Optimize(objects []_cg.PdfObject) (_cdgc []_cg.PdfObject, _adgcf error) {
	if _dee.ImageUpperPPI <= 0 {
		return objects, nil
	}
	_efg := _gbc(objects)
	if len(_efg) == 0 {
		return objects, nil
	}
	_egga := make(map[_cg.PdfObject]struct{})
	for _, _abc := range _efg {
		_ebbg := _abc.Stream.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b")
		_egga[_ebbg] = struct{}{}
	}
	_cgb := make(map[*_cg.PdfObjectStream]*imageInfo)
	for _, _cad := range _efg {
		_cgb[_cad.Stream] = _cad
	}
	var _beafc *_cg.PdfObjectDictionary
	for _, _ggd := range objects {
		if _aef, _aab := _cg.GetDict(_ggd); _beafc == nil && _aab {
			if _bdbg, _gdc := _cg.GetName(_aef.Get("\u0054\u0079\u0070\u0065")); _gdc && *_bdbg == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
				_beafc = _aef
			}
		}
	}
	if _beafc == nil {
		return objects, nil
	}
	_aaae, _dbcd := _cg.GetDict(_beafc.Get("\u0050\u0061\u0067e\u0073"))
	if !_dbcd {
		return objects, nil
	}
	_bcc, _fddd := _cg.GetArray(_aaae.Get("\u004b\u0069\u0064\u0073"))
	if !_fddd {
		return objects, nil
	}
	for _, _cde := range _bcc.Elements() {
		_dbfg := make(map[string]*imageInfo)
		_cfe, _dgb := _cg.GetDict(_cde)
		if !_dgb {
			continue
		}
		_aefb, _ := _eagd(_cfe.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		if len(_aefb) == 0 {
			continue
		}
		_ddc, _ddgd := _cg.GetDict(_cfe.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
		if !_ddgd {
			continue
		}
		_afe, _dfc := _e.NewPdfPageResourcesFromDict(_ddc)
		if _dfc != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020-\u0020\u0069\u0067\u006e\u006fr\u0069\u006eg\u003a\u0020\u0025\u0076", _dfc)
			continue
		}
		_ecc, _egdd := _cg.GetDict(_ddc.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
		if !_egdd {
			continue
		}
		_bbdg := _ecc.Keys()
		for _, _agcc := range _bbdg {
			if _bafe, _egdc := _cg.GetStream(_ecc.Get(_agcc)); _egdc {
				if _gefc, _acgc := _cgb[_bafe]; _acgc {
					_dbfg[string(_agcc)] = _gefc
				}
			}
		}
		_abac := _bb.NewContentStreamParser(_aefb)
		_dfca, _dfc := _abac.Parse()
		if _dfc != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _dfc)
			continue
		}
		_gffg := _bb.NewContentStreamProcessor(*_dfca)
		_gffg.AddHandler(_bb.HandlerConditionEnumAllOperands, "", func(_agfb *_bb.ContentStreamOperation, _bbde _bb.GraphicsState, _cgd *_e.PdfPageResources) error {
			switch _agfb.Operand {
			case "\u0044\u006f":
				if len(_agfb.Params) != 1 {
					_f.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020w\u0069\u0074\u0068\u0020\u006c\u0065\u006e\u0028\u0070\u0061ra\u006d\u0073\u0029 \u0021=\u0020\u0031")
					return nil
				}
				_bgde, _agca := _cg.GetName(_agfb.Params[0])
				if !_agca {
					_f.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f\u006e\u0020\u004e\u0061\u006d\u0065\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072")
					return nil
				}
				if _cbfb, _gdg := _dbfg[string(*_bgde)]; _gdg {
					_cadg := _bbde.CTM.ScalingFactorX()
					_gage := _bbde.CTM.ScalingFactorY()
					_geed, _dde := _cadg/72.0, _gage/72.0
					_abg, _bdaf := float64(_cbfb.Width)/_geed, float64(_cbfb.Height)/_dde
					if _geed == 0 || _dde == 0 {
						_abg = 72.0
						_bdaf = 72.0
					}
					_cbfb.PPI = _af.Max(_cbfb.PPI, _abg)
					_cbfb.PPI = _af.Max(_cbfb.PPI, _bdaf)
				}
			}
			return nil
		})
		_dfc = _gffg.Process(_afe)
		if _dfc != nil {
			_f.Log.Debug("E\u0052\u0052\u004f\u0052 p\u0072o\u0063\u0065\u0073\u0073\u0069n\u0067\u003a\u0020\u0025\u002b\u0076", _dfc)
			continue
		}
	}
	for _, _edca := range _efg {
		if _, _deee := _egga[_edca.Stream]; _deee {
			continue
		}
		if _edca.PPI <= _dee.ImageUpperPPI {
			continue
		}
		_bac, _bccb := _e.NewXObjectImageFromStream(_edca.Stream)
		if _bccb != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _bccb)
			continue
		}
		var _fbe imageModifications
		_fbe.Scale = _dee.ImageUpperPPI / _edca.PPI
		if _edca.BitsPerComponent == 1 && _edca.ColorComponents == 1 {
			_gfab := _af.Round(_edca.PPI / _dee.ImageUpperPPI)
			_abea := _aa.NextPowerOf2(uint(_gfab))
			if _aa.InDelta(float64(_abea), 1/_fbe.Scale, 0.3) {
				_fbe.Scale = float64(1) / float64(_abea)
			}
			if _, _bcg := _bac.Filter.(*_cg.JBIG2Encoder); !_bcg {
				_fbe.Encoding = _cg.NewJBIG2Encoder()
			}
		}
		if _bccb = _ecgf(_bac, _fbe); _bccb != nil {
			_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072 \u0073\u0063\u0061\u006c\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u006be\u0065\u0070\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _bccb)
			continue
		}
		_fbe.Encoding = nil
		if _gbf, _egeba := _cg.GetStream(_edca.Stream.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b")); _egeba {
			_baee, _cdae := _e.NewXObjectImageFromStream(_gbf)
			if _cdae != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _cdae)
				continue
			}
			if _cdae = _ecgf(_baee, _fbe); _cdae != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _cdae)
				continue
			}
		}
	}
	return objects, nil
}

type imageInfo struct {
	BitsPerComponent int
	ColorComponents  int
	Width            int
	Height           int
	Stream           *_cg.PdfObjectStream
	PPI              float64
}

func _ecgf(_bca *_e.XObjectImage, _eaee imageModifications) error {
	_bbga, _efa := _bca.ToImage()
	if _efa != nil {
		return _efa
	}
	if _eaee.Scale != 0 {
		_bbga, _efa = _ccc(_bbga, _eaee.Scale)
		if _efa != nil {
			return _efa
		}
	}
	if _eaee.Encoding != nil {
		_bca.Filter = _eaee.Encoding
	}
	_bca.Decode = nil
	switch _aebg := _bca.Filter.(type) {
	case *_cg.FlateEncoder:
		if _aebg.Predictor != 1 && _aebg.Predictor != 11 {
			_aebg.Predictor = 1
		}
	}
	if _efa = _bca.SetImage(_bbga, nil); _efa != nil {
		_f.Log.Debug("\u0045\u0072\u0072or\u0020\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0076", _efa)
		return _efa
	}
	_bca.ToPdfObject()
	return nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_agf *CombineIdenticalIndirectObjects) Optimize(objects []_cg.PdfObject) (_agba []_cg.PdfObject, _abae error) {
	_dda(objects)
	_cdg := make(map[_cg.PdfObject]_cg.PdfObject)
	_fdc := make(map[_cg.PdfObject]struct{})
	_gbg := make(map[string][]*_cg.PdfIndirectObject)
	for _, _fbc := range objects {
		_dfa, _gbag := _fbc.(*_cg.PdfIndirectObject)
		if !_gbag {
			continue
		}
		if _gfbe, _bfce := _dfa.PdfObject.(*_cg.PdfObjectDictionary); _bfce {
			if _cca, _cac := _gfbe.Get("\u0054\u0079\u0070\u0065").(*_cg.PdfObjectName); _cac && *_cca == "\u0050\u0061\u0067\u0065" {
				continue
			}
			_ggg := _afb.New()
			_ggg.Write([]byte(_gfbe.WriteString()))
			_agfg := string(_ggg.Sum(nil))
			_gbg[_agfg] = append(_gbg[_agfg], _dfa)
		}
	}
	for _, _agc := range _gbg {
		if len(_agc) < 2 {
			continue
		}
		_bge := _agc[0]
		for _beaf := 1; _beaf < len(_agc); _beaf++ {
			_fbcc := _agc[_beaf]
			_cdg[_fbcc] = _bge
			_fdc[_fbcc] = struct{}{}
		}
	}
	_agba = make([]_cg.PdfObject, 0, len(objects)-len(_fdc))
	for _, _dgd := range objects {
		if _, _feef := _fdc[_dgd]; _feef {
			continue
		}
		_agba = append(_agba, _dgd)
	}
	_gac(_agba, _cdg)
	return _agba, nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_fgaa *ObjectStreams) Optimize(objects []_cg.PdfObject) (_acdd []_cg.PdfObject, _bbdc error) {
	_fbed := &_cg.PdfObjectStreams{}
	_gcag := make([]_cg.PdfObject, 0, len(objects))
	for _, _fce := range objects {
		if _dgff, _bcag := _fce.(*_cg.PdfIndirectObject); _bcag && _dgff.GenerationNumber == 0 {
			_fbed.Append(_fce)
		} else {
			_gcag = append(_gcag, _fce)
		}
	}
	if _fbed.Len() == 0 {
		return _gcag, nil
	}
	_acdd = make([]_cg.PdfObject, 0, len(_gcag)+_fbed.Len()+1)
	if _fbed.Len() > 1 {
		_acdd = append(_acdd, _fbed)
	}
	_acdd = append(_acdd, _fbed.Elements()...)
	_acdd = append(_acdd, _gcag...)
	return _acdd, nil
}

// CleanFonts cleans up embedded fonts, reducing font sizes.
type CleanFonts struct {

	// Subset embedded fonts if encountered (if true).
	// Otherwise attempts to reduce the font program.
	Subset bool
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_dd *CombineDuplicateStreams) Optimize(objects []_cg.PdfObject) (_dbg []_cg.PdfObject, _egeb error) {
	_baa := make(map[_cg.PdfObject]_cg.PdfObject)
	_bdg := make(map[_cg.PdfObject]struct{})
	_feff := make(map[string][]*_cg.PdfObjectStream)
	for _, _gff := range objects {
		if _dfe, _egb := _gff.(*_cg.PdfObjectStream); _egb {
			_ccfa := _afb.New()
			_ccfa.Write(_dfe.Stream)
			_ccfa.Write([]byte(_dfe.PdfObjectDictionary.WriteString()))
			_dae := string(_ccfa.Sum(nil))
			_feff[_dae] = append(_feff[_dae], _dfe)
		}
	}
	for _, _efc := range _feff {
		if len(_efc) < 2 {
			continue
		}
		_cbf := _efc[0]
		for _fgf := 1; _fgf < len(_efc); _fgf++ {
			_afgb := _efc[_fgf]
			_baa[_afgb] = _cbf
			_bdg[_afgb] = struct{}{}
		}
	}
	_dbg = make([]_cg.PdfObject, 0, len(objects)-len(_bdg))
	for _, _bfe := range objects {
		if _, _abee := _bdg[_bfe]; _abee {
			continue
		}
		_dbg = append(_dbg, _bfe)
	}
	_gac(_dbg, _baa)
	return _dbg, nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_eda *CleanContentstream) Optimize(objects []_cg.PdfObject) (_cbb []_cg.PdfObject, _fc error) {
	_bd := map[*_cg.PdfObjectStream]struct{}{}
	var _fbb []*_cg.PdfObjectStream
	_be := func(_ca *_cg.PdfObjectStream) {
		if _, _ag := _bd[_ca]; !_ag {
			_bd[_ca] = struct{}{}
			_fbb = append(_fbb, _ca)
		}
	}
	_ff := map[_cg.PdfObject]bool{}
	_fa := map[_cg.PdfObject]bool{}
	for _, _bdd := range objects {
		switch _ceb := _bdd.(type) {
		case *_cg.PdfIndirectObject:
			switch _fg := _ceb.PdfObject.(type) {
			case *_cg.PdfObjectDictionary:
				if _ffe, _gg := _cg.GetName(_fg.Get("\u0054\u0079\u0070\u0065")); !_gg || _ffe.String() != "\u0050\u0061\u0067\u0065" {
					continue
				}
				if _gcc, _eee := _cg.GetStream(_fg.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073")); _eee {
					_be(_gcc)
				} else if _cbd, _abef := _cg.GetArray(_fg.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073")); _abef {
					var _aad []*_cg.PdfObjectStream
					for _, _ecb := range _cbd.Elements() {
						if _bgf, _adg := _cg.GetStream(_ecb); _adg {
							_aad = append(_aad, _bgf)
						}
					}
					if len(_aad) > 0 {
						var _gccg _bg.Buffer
						for _, _fbf := range _aad {
							if _dc, _fef := _cg.DecodeStream(_fbf); _fef == nil {
								_gccg.Write(_dc)
							}
							_ff[_fbf] = true
						}
						_edf, _gfg := _cg.MakeStream(_gccg.Bytes(), _cg.NewFlateEncoder())
						if _gfg != nil {
							return nil, _gfg
						}
						_fa[_edf] = true
						_fg.Set("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _edf)
						_be(_edf)
					}
				}
			}
		case *_cg.PdfObjectStream:
			if _cfa, _ccd := _cg.GetName(_ceb.Get("\u0054\u0079\u0070\u0065")); !_ccd || _cfa.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
				continue
			}
			if _bdc, _fda := _cg.GetName(_ceb.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); !_fda || _bdc.String() != "\u0046\u006f\u0072\u006d" {
				continue
			}
			_be(_ceb)
		}
	}
	for _, _gba := range _fbb {
		_fc = _eg(_gba)
		if _fc != nil {
			return nil, _fc
		}
	}
	_cbb = nil
	for _, _cfc := range objects {
		if _ff[_cfc] {
			continue
		}
		_cbb = append(_cbb, _cfc)
	}
	for _aadc := range _fa {
		_cbb = append(_cbb, _aadc)
	}
	return _cbb, nil
}
func _abe(_ea *_bb.ContentStreamOperations) *_bb.ContentStreamOperations {
	if _ea == nil {
		return nil
	}
	_ba := _bb.ContentStreamOperations{}
	for _, _cgc := range *_ea {
		switch _cgc.Operand {
		case "\u0042\u0044\u0043", "\u0042\u004d\u0043", "\u0045\u004d\u0043":
			continue
		case "\u0054\u006d":
			if len(_cgc.Params) == 6 {
				if _cb, _cc := _cg.GetNumbersAsFloat(_cgc.Params); _cc == nil {
					if _cb[0] == 1 && _cb[1] == 0 && _cb[2] == 0 && _cb[3] == 1 {
						_cgc = &_bb.ContentStreamOperation{Params: []_cg.PdfObject{_cgc.Params[4], _cgc.Params[5]}, Operand: "\u0054\u0064"}
					}
				}
			}
		}
		_ba = append(_ba, _cgc)
	}
	return &_ba
}

// ObjectStreams groups PDF objects to object streams.
// It implements interface model.Optimizer.
type ObjectStreams struct{}

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

// Append appends optimizers to the chain.
func (_ee *Chain) Append(optimizers ..._e.Optimizer) { _ee._gb = append(_ee._gb, optimizers...) }

// New creates a optimizers chain from options.
func New(options Options) *Chain {
	_afba := new(Chain)
	if options.CleanFonts || options.SubsetFonts {
		_afba.Append(&CleanFonts{Subset: options.SubsetFonts})
	}
	if options.CleanContentstream {
		_afba.Append(new(CleanContentstream))
	}
	if options.ImageUpperPPI > 0 {
		_cabb := new(ImagePPI)
		_cabb.ImageUpperPPI = options.ImageUpperPPI
		_afba.Append(_cabb)
	}
	if options.ImageQuality > 0 {
		_caca := new(Image)
		_caca.ImageQuality = options.ImageQuality
		_afba.Append(_caca)
	}
	if options.CombineDuplicateDirectObjects {
		_afba.Append(new(CombineDuplicateDirectObjects))
	}
	if options.CombineDuplicateStreams {
		_afba.Append(new(CombineDuplicateStreams))
	}
	if options.CombineIdenticalIndirectObjects {
		_afba.Append(new(CombineIdenticalIndirectObjects))
	}
	if options.UseObjectStreams {
		_afba.Append(new(ObjectStreams))
	}
	if options.CompressStreams {
		_afba.Append(new(CompressStreams))
	}
	return _afba
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_ebc *CleanFonts) Optimize(objects []_cg.PdfObject) (_bbc []_cg.PdfObject, _dec error) {
	var _deg map[*_cg.PdfObjectStream]struct{}
	if _ebc.Subset {
		var _bgd error
		_deg, _bgd = _de(objects)
		if _bgd != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0073u\u0062s\u0065\u0074\u0074\u0069\u006e\u0067\u003a \u0025\u0076", _bgd)
			return nil, _bgd
		}
	}
	for _, _ccb := range objects {
		_cda, _bdb := _cg.GetStream(_ccb)
		if !_bdb {
			continue
		}
		if _, _eafe := _deg[_cda]; _eafe {
			continue
		}
		_adgc, _ffcd := _cg.NewEncoderFromStream(_cda)
		if _ffcd != nil {
			_f.Log.Debug("\u0045\u0052RO\u0052\u0020\u0067e\u0074\u0074\u0069\u006eg e\u006eco\u0064\u0065\u0072\u003a\u0020\u0025\u0076 -\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067", _ffcd)
			continue
		}
		_gcfd, _ffcd := _adgc.DecodeStream(_cda)
		if _ffcd != nil {
			_f.Log.Debug("\u0044\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065r\u0072\u006f\u0072\u0020\u003a\u0020\u0025v\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067", _ffcd)
			continue
		}
		if len(_gcfd) < 4 {
			continue
		}
		_ecd := string(_gcfd[:4])
		if _ecd == "\u004f\u0054\u0054\u004f" {
			continue
		}
		if _ecd != "\u0000\u0001\u0000\u0000" && _ecd != "\u0074\u0072\u0075\u0065" {
			continue
		}
		_bda, _ffcd := _b.Parse(_bg.NewReader(_gcfd))
		if _ffcd != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020P\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _ffcd)
			continue
		}
		_ffcd = _bda.Optimize()
		if _ffcd != nil {
			_f.Log.Debug("\u0045\u0052RO\u0052\u0020\u004fp\u0074\u0069\u006d\u0069zin\u0067 f\u006f\u006e\u0074\u003a\u0020\u0025\u0076 -\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067", _ffcd)
			continue
		}
		var _def _bg.Buffer
		_ffcd = _bda.Write(&_def)
		if _ffcd != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020W\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _ffcd)
			continue
		}
		if _def.Len() > len(_gcfd) {
			_f.Log.Debug("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070")
			continue
		}
		_dbd, _ffcd := _cg.MakeStream(_def.Bytes(), _cg.NewFlateEncoder())
		if _ffcd != nil {
			continue
		}
		*_cda = *_dbd
		_cda.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _cg.MakeInteger(int64(_def.Len())))
	}
	return objects, nil
}

// CombineDuplicateStreams combines duplicated streams by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateStreams struct{}
type objectStructure struct {
	_ffd  *_cg.PdfObjectDictionary
	_gddf *_cg.PdfObjectDictionary
	_gfd  []*_cg.PdfIndirectObject
}

// CleanContentstream cleans up redundant operands in content streams, including Page and XObject Form
// contents. This process includes:
// 1. Marked content operators are removed.
// 2. Some operands are simplified (shorter form).
// TODO: Add more reduction methods and improving the methods for identifying unnecessary operands.
type CleanContentstream struct{}

// CombineIdenticalIndirectObjects combines identical indirect objects.
// It implements interface model.Optimizer.
type CombineIdenticalIndirectObjects struct{}

// CombineDuplicateDirectObjects combines duplicated direct objects by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateDirectObjects struct{}

// ImagePPI optimizes images by scaling images such that the PPI (pixels per inch) is never higher than ImageUpperPPI.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type ImagePPI struct{ ImageUpperPPI float64 }

func _ccc(_efef *_e.Image, _adde float64) (*_e.Image, error) {
	_bfec, _afda := _efef.ToGoImage()
	if _afda != nil {
		return nil, _afda
	}
	var _cfb _aa.Image
	_dgc, _ggaf := _bfec.(*_aa.Monochrome)
	if _ggaf {
		if _afda = _dgc.ResolveDecode(); _afda != nil {
			return nil, _afda
		}
		_cfb, _afda = _dgc.Scale(_adde)
		if _afda != nil {
			return nil, _afda
		}
	} else {
		_daf := int(_af.RoundToEven(float64(_efef.Width) * _adde))
		_gbcf := int(_af.RoundToEven(float64(_efef.Height) * _adde))
		_cfb, _afda = _aa.NewImage(_daf, _gbcf, int(_efef.BitsPerComponent), _efef.ColorComponents, nil, nil, nil)
		if _afda != nil {
			return nil, _afda
		}
		_g.CatmullRom.Scale(_cfb, _cfb.Bounds(), _bfec, _bfec.Bounds(), _g.Over, &_g.Options{})
	}
	_ceec := _cfb.Base()
	_bfaa := &_e.Image{Width: int64(_ceec.Width), Height: int64(_ceec.Height), BitsPerComponent: int64(_ceec.BitsPerComponent), ColorComponents: _ceec.ColorComponents, Data: _ceec.Data}
	_bfaa.SetDecode(_ceec.Decode)
	_bfaa.SetAlpha(_ceec.Alpha)
	return _bfaa, nil
}
func _gac(_dfaf []_cg.PdfObject, _dafc map[_cg.PdfObject]_cg.PdfObject) {
	if len(_dafc) == 0 {
		return
	}
	for _fge, _ffb := range _dfaf {
		if _caaa, _dbfe := _dafc[_ffb]; _dbfe {
			_dfaf[_fge] = _caaa
			continue
		}
		_dafc[_ffb] = _ffb
		switch _feefd := _ffb.(type) {
		case *_cg.PdfObjectArray:
			_bcaa := make([]_cg.PdfObject, _feefd.Len())
			copy(_bcaa, _feefd.Elements())
			_gac(_bcaa, _dafc)
			for _ega, _aaaeg := range _bcaa {
				_feefd.Set(_ega, _aaaeg)
			}
		case *_cg.PdfObjectStreams:
			_gac(_feefd.Elements(), _dafc)
		case *_cg.PdfObjectStream:
			_abeb := []_cg.PdfObject{_feefd.PdfObjectDictionary}
			_gac(_abeb, _dafc)
			_feefd.PdfObjectDictionary = _abeb[0].(*_cg.PdfObjectDictionary)
		case *_cg.PdfObjectDictionary:
			_aadf := _feefd.Keys()
			_gbfa := make([]_cg.PdfObject, len(_aadf))
			for _cdbd, _bga := range _aadf {
				_gbfa[_cdbd] = _feefd.Get(_bga)
			}
			_gac(_gbfa, _dafc)
			for _bccg, _abf := range _aadf {
				_feefd.Set(_abf, _gbfa[_bccg])
			}
		case *_cg.PdfIndirectObject:
			_ggbe := []_cg.PdfObject{_feefd.PdfObject}
			_gac(_ggbe, _dafc)
			_feefd.PdfObject = _ggbe[0]
		}
	}
}
func _de(_aadcf []_cg.PdfObject) (_bad map[*_cg.PdfObjectStream]struct{}, _cbe error) {
	_bad = map[*_cg.PdfObjectStream]struct{}{}
	_ecg := map[*_e.PdfFont]struct{}{}
	_afc := _ebdc(_aadcf)
	for _, _fcg := range _afc._gfd {
		_cgca, _db := _cg.GetDict(_fcg.PdfObject)
		if !_db {
			continue
		}
		_dg, _db := _cg.GetDict(_cgca.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
		if !_db {
			continue
		}
		_afbc, _ := _eagd(_cgca.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		_eaf, _cae := _e.NewPdfPageResourcesFromDict(_dg)
		if _cae != nil {
			return nil, _cae
		}
		_dge := []content{{_ded: _afbc, _dce: _eaf}}
		_bea := _aba(_cgca.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _bea != nil {
			_dge = append(_dge, _bea...)
		}
		for _, _edc := range _dge {
			_ccf, _fdgb := _ab.NewFromContents(_edc._ded, _edc._dce)
			if _fdgb != nil {
				return nil, _fdgb
			}
			_ffc, _, _, _fdgb := _ccf.ExtractPageText()
			if _fdgb != nil {
				return nil, _fdgb
			}
			for _, _gcb := range _ffc.Marks().Elements() {
				if _gcb.Font == nil {
					continue
				}
				if _, _fcf := _ecg[_gcb.Font]; !_fcf {
					_ecg[_gcb.Font] = struct{}{}
				}
			}
		}
	}
	_gfc := map[*_cg.PdfObjectStream][]*_e.PdfFont{}
	for _gcg := range _ecg {
		_gea := _gcg.FontDescriptor()
		if _gea == nil || _gea.FontFile2 == nil {
			continue
		}
		_fca, _faf := _cg.GetStream(_gea.FontFile2)
		if !_faf {
			continue
		}
		_gfc[_fca] = append(_gfc[_fca], _gcg)
	}
	for _ae := range _gfc {
		var _acc []rune
		var _fad []_b.GlyphIndex
		for _, _acd := range _gfc[_ae] {
			switch _egc := _acd.Encoder().(type) {
			case *_ac.IdentityEncoder:
				_dbc := _egc.RegisteredRunes()
				_edg := make([]_b.GlyphIndex, len(_dbc))
				for _ef, _aac := range _dbc {
					_edg[_ef] = _b.GlyphIndex(_aac)
				}
				_fad = append(_fad, _edg...)
			case *_ac.TrueTypeFontEncoder:
				_edfa := _egc.RegisteredRunes()
				_acc = append(_acc, _edfa...)
			case _ac.SimpleEncoder:
				_acg := _egc.Charcodes()
				for _, _gfa := range _acg {
					_afg, _aae := _egc.CharcodeToRune(_gfa)
					if !_aae {
						_f.Log.Debug("\u0043\u0068a\u0072\u0063\u006f\u0064\u0065\u003c\u002d\u003e\u0072\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064: \u0025\u0064", _gfa)
						continue
					}
					_acc = append(_acc, _afg)
				}
			}
		}
		_cbe = _fff(_ae, _acc, _fad)
		if _cbe != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006eg\u0020f\u006f\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _cbe)
			return nil, _cbe
		}
		_bad[_ae] = struct{}{}
	}
	return _bad, nil
}
func _dda(_bcgd []_cg.PdfObject) {
	for _aaec, _ebbca := range _bcgd {
		switch _gcgeg := _ebbca.(type) {
		case *_cg.PdfIndirectObject:
			_gcgeg.ObjectNumber = int64(_aaec + 1)
			_gcgeg.GenerationNumber = 0
		case *_cg.PdfObjectStream:
			_gcgeg.ObjectNumber = int64(_aaec + 1)
			_gcgeg.GenerationNumber = 0
		case *_cg.PdfObjectStreams:
			_gcgeg.ObjectNumber = int64(_aaec + 1)
			_gcgeg.GenerationNumber = 0
		}
	}
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_ebcf *CombineDuplicateDirectObjects) Optimize(objects []_cg.PdfObject) (_gga []_cg.PdfObject, _gced error) {
	_dda(objects)
	_dgaa := make(map[string][]*_cg.PdfObjectDictionary)
	var _fdd func(_bfa *_cg.PdfObjectDictionary)
	_fdd = func(_fbg *_cg.PdfObjectDictionary) {
		for _, _baf := range _fbg.Keys() {
			_adb := _fbg.Get(_baf)
			if _ebb, _ede := _adb.(*_cg.PdfObjectDictionary); _ede {
				_bce := _afb.New()
				_bce.Write([]byte(_ebb.WriteString()))
				_decf := string(_bce.Sum(nil))
				_dgaa[_decf] = append(_dgaa[_decf], _ebb)
				_fdd(_ebb)
			}
		}
	}
	for _, _acag := range objects {
		_dgag, _ga := _acag.(*_cg.PdfIndirectObject)
		if !_ga {
			continue
		}
		if _cab, _ade := _dgag.PdfObject.(*_cg.PdfObjectDictionary); _ade {
			_fdd(_cab)
		}
	}
	_fada := make([]_cg.PdfObject, 0, len(_dgaa))
	_caef := make(map[_cg.PdfObject]_cg.PdfObject)
	for _, _abbf := range _dgaa {
		if len(_abbf) < 2 {
			continue
		}
		_eca := _cg.MakeDict()
		_eca.Merge(_abbf[0])
		_dba := _cg.MakeIndirectObject(_eca)
		_fada = append(_fada, _dba)
		for _ebd := 0; _ebd < len(_abbf); _ebd++ {
			_ebf := _abbf[_ebd]
			_caef[_ebf] = _dba
		}
	}
	_gga = make([]_cg.PdfObject, len(objects))
	copy(_gga, objects)
	_gga = append(_fada, _gga...)
	_gac(_gga, _caef)
	return _gga, nil
}

type content struct {
	_ded string
	_dce *_e.PdfPageResources
}

func _eagd(_beg _cg.PdfObject) (_dbfd string, _ccdd []_cg.PdfObject) {
	var _abefd _bg.Buffer
	switch _cge := _beg.(type) {
	case *_cg.PdfIndirectObject:
		_ccdd = append(_ccdd, _cge)
		_beg = _cge.PdfObject
	}
	switch _gcad := _beg.(type) {
	case *_cg.PdfObjectStream:
		if _eaef, _fbcce := _cg.DecodeStream(_gcad); _fbcce == nil {
			_abefd.Write(_eaef)
			_ccdd = append(_ccdd, _gcad)
		}
	case *_cg.PdfObjectArray:
		for _, _dbbf := range _gcad.Elements() {
			switch _ccce := _dbbf.(type) {
			case *_cg.PdfObjectStream:
				if _ddec, _cfbd := _cg.DecodeStream(_ccce); _cfbd == nil {
					_abefd.Write(_ddec)
					_ccdd = append(_ccdd, _ccce)
				}
			}
		}
	}
	return _abefd.String(), _ccdd
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_bgdb *CompressStreams) Optimize(objects []_cg.PdfObject) (_ggc []_cg.PdfObject, _dag error) {
	_ggc = make([]_cg.PdfObject, len(objects))
	copy(_ggc, objects)
	for _, _cedb := range objects {
		_dbdg, _ggb := _cg.GetStream(_cedb)
		if !_ggb {
			continue
		}
		if _gfbd := _dbdg.Get("\u0046\u0069\u006c\u0074\u0065\u0072"); _gfbd != nil {
			if _, _dab := _cg.GetName(_gfbd); _dab {
				continue
			}
			if _bgdf, _caea := _cg.GetArray(_gfbd); _caea && _bgdf.Len() > 0 {
				continue
			}
		}
		_edfe := _cg.NewFlateEncoder()
		var _cce []byte
		_cce, _dag = _edfe.EncodeBytes(_dbdg.Stream)
		if _dag != nil {
			return _ggc, _dag
		}
		_gec := _edfe.MakeStreamDict()
		if len(_cce)+len(_gec.WriteString()) < len(_dbdg.Stream) {
			_dbdg.Stream = _cce
			_dbdg.PdfObjectDictionary.Merge(_gec)
			_dbdg.PdfObjectDictionary.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _cg.MakeInteger(int64(len(_dbdg.Stream))))
		}
	}
	return _ggc, nil
}
func _ebdc(_cadb []_cg.PdfObject) objectStructure {
	_eab := objectStructure{}
	_ccfad := false
	for _, _ccagb := range _cadb {
		switch _bcac := _ccagb.(type) {
		case *_cg.PdfIndirectObject:
			_gdf, _aed := _cg.GetDict(_bcac)
			if !_aed {
				continue
			}
			_gfgf, _aed := _cg.GetName(_gdf.Get("\u0054\u0079\u0070\u0065"))
			if !_aed {
				continue
			}
			switch _gfgf.String() {
			case "\u0043a\u0074\u0061\u006c\u006f\u0067":
				_eab._ffd = _gdf
				_ccfad = true
			}
		}
		if _ccfad {
			break
		}
	}
	if !_ccfad {
		return _eab
	}
	_cegc, _fbge := _cg.GetDict(_eab._ffd.Get("\u0050\u0061\u0067e\u0073"))
	if !_fbge {
		return _eab
	}
	_eab._gddf = _cegc
	_gbgc, _fbge := _cg.GetArray(_cegc.Get("\u004b\u0069\u0064\u0073"))
	if !_fbge {
		return _eab
	}
	for _, _gbcb := range _gbgc.Elements() {
		_eeb, _fcfg := _cg.GetIndirect(_gbcb)
		if !_fcfg {
			break
		}
		_eab._gfd = append(_eab._gfd, _eeb)
	}
	return _eab
}
func _eg(_fb *_cg.PdfObjectStream) error {
	_cd, _ed := _cg.DecodeStream(_fb)
	if _ed != nil {
		return _ed
	}
	_ad := _bb.NewContentStreamParser(string(_cd))
	_aaa, _ed := _ad.Parse()
	if _ed != nil {
		return _ed
	}
	_aaa = _abe(_aaa)
	_d := _aaa.Bytes()
	if len(_d) >= len(_cd) {
		return nil
	}
	_fdg, _ed := _cg.MakeStream(_aaa.Bytes(), _cg.NewFlateEncoder())
	if _ed != nil {
		return _ed
	}
	_fb.Stream = _fdg.Stream
	_fb.Merge(_fdg.PdfObjectDictionary)
	return nil
}
func _fff(_fee *_cg.PdfObjectStream, _aca []rune, _fcff []_b.GlyphIndex) error {
	_fee, _gce := _cg.GetStream(_fee)
	if !_gce {
		_f.Log.Debug("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u002d\u002d\u0020\u0041\u0042\u004f\u0052T\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006e\u0067")
		return _c.New("\u0066\u006f\u006e\u0074fi\u006c\u0065\u0032\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_cgg, _ccgd := _cg.DecodeStream(_fee)
	if _ccgd != nil {
		_f.Log.Debug("\u0044\u0065c\u006f\u0064\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _ccgd)
		return _ccgd
	}
	_bf, _ccgd := _b.Parse(_bg.NewReader(_cgg))
	if _ccgd != nil {
		_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0020f\u006f\u006e\u0074", len(_fee.Stream))
		return _ccgd
	}
	_dga := _fcff
	if len(_aca) > 0 {
		_cbde := _bf.LookupRunes(_aca)
		_dga = append(_dga, _cbde...)
	}
	_bf, _ccgd = _bf.SubsetKeepIndices(_dga)
	if _ccgd != nil {
		_f.Log.Debug("\u0045R\u0052\u004f\u0052\u0020s\u0075\u0062\u0073\u0065\u0074t\u0069n\u0067 \u0066\u006f\u006e\u0074\u003a\u0020\u0025v", _ccgd)
		return _ccgd
	}
	var _gbe _bg.Buffer
	_ccgd = _bf.Write(&_gbe)
	if _ccgd != nil {
		_f.Log.Debug("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _ccgd)
		return _ccgd
	}
	if _gbe.Len() > len(_cgg) {
		_f.Log.Debug("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070")
		return nil
	}
	_bfc, _ccgd := _cg.MakeStream(_gbe.Bytes(), _cg.NewFlateEncoder())
	if _ccgd != nil {
		_f.Log.Debug("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _ccgd)
		return _ccgd
	}
	*_fee = *_bfc
	_fee.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _cg.MakeInteger(int64(_gbe.Len())))
	return nil
}

// Image optimizes images by rewrite images into JPEG format with quality equals to ImageQuality.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type Image struct{ ImageQuality int }

// Chain allows to use sequence of optimizers.
// It implements interface model.Optimizer.
type Chain struct{ _gb []_e.Optimizer }

func _gbc(_aaee []_cg.PdfObject) []*imageInfo {
	_abbg := _cg.PdfObjectName("\u0053u\u0062\u0074\u0079\u0070\u0065")
	_fbbd := make(map[*_cg.PdfObjectStream]struct{})
	var _cbbg []*imageInfo
	for _, _ada := range _aaee {
		_fba, _gag := _cg.GetStream(_ada)
		if !_gag {
			continue
		}
		if _, _ebbc := _fbbd[_fba]; _ebbc {
			continue
		}
		_fbbd[_fba] = struct{}{}
		_fdcg := _fba.PdfObjectDictionary.Get(_abbg)
		_ecab, _gag := _cg.GetName(_fdcg)
		if !_gag || string(*_ecab) != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		_fgff := &imageInfo{Stream: _fba, BitsPerComponent: 8}
		if _dfg, _bgfb := _cg.GetIntVal(_fba.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _bgfb {
			_fgff.BitsPerComponent = _dfg
		}
		if _afd, _adee := _cg.GetIntVal(_fba.Get("\u0057\u0069\u0064t\u0068")); _adee {
			_fgff.Width = _afd
		}
		if _beag, _dbfb := _cg.GetIntVal(_fba.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _dbfb {
			_fgff.Height = _beag
		}
		_agbd, _fgab := _e.NewPdfColorspaceFromPdfObject(_fba.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"))
		if _fgab != nil {
			_f.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fgab)
			continue
		}
		if _agbd == nil {
			_geg, _aafe := _cg.GetName(_fba.Get("\u0046\u0069\u006c\u0074\u0065\u0072"))
			if _aafe {
				switch _geg.String() {
				case "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065", "J\u0042\u0049\u0047\u0032\u0044\u0065\u0063\u006f\u0064\u0065":
					_agbd = _e.NewPdfColorspaceDeviceGray()
					_fgff.BitsPerComponent = 1
				}
			}
		}
		switch _eeed := _agbd.(type) {
		case *_e.PdfColorspaceDeviceRGB:
			_fgff.ColorComponents = 3
		case *_e.PdfColorspaceDeviceGray:
			_fgff.ColorComponents = 1
		default:
			_f.Log.Debug("\u004f\u0070\u0074\u0069\u006d\u0069\u007aa\u0074\u0069\u006fn\u0020\u0069\u0073 \u006e\u006ft\u0020\u0073\u0075\u0070\u0070\u006fr\u0074ed\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u006b\u0069\u0070", _eeed)
			continue
		}
		_cbbg = append(_cbbg, _fgff)
	}
	return _cbbg
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_gf *Chain) Optimize(objects []_cg.PdfObject) (_gc []_cg.PdfObject, _gd error) {
	_gcf := objects
	for _, _eef := range _gf._gb {
		_ce, _fd := _eef.Optimize(_gcf)
		if _fd != nil {
			_f.Log.Debug("\u0045\u0052\u0052OR\u0020\u004f\u0070\u0074\u0069\u006d\u0069\u007a\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u002b\u0076", _fd)
			continue
		}
		_gcf = _ce
	}
	return _gcf, nil
}

type imageModifications struct {
	Scale    float64
	Encoding _cg.StreamEncoder
}

func _aba(_fcc _cg.PdfObject) []content {
	if _fcc == nil {
		return nil
	}
	_ace, _ecdd := _cg.GetArray(_fcc)
	if !_ecdd {
		_f.Log.Debug("\u0041\u006e\u006e\u006fts\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
		return nil
	}
	var _ecde []content
	for _, _add := range _ace.Elements() {
		_gfb, _gcge := _cg.GetDict(_add)
		if !_gcge {
			_f.Log.Debug("I\u0067\u006e\u006f\u0072\u0069\u006eg\u0020\u006e\u006f\u006e\u002d\u0064i\u0063\u0074\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006e\u0020\u0041\u006e\u006e\u006ft\u0073")
			continue
		}
		_bbb, _gcge := _cg.GetDict(_gfb.Get("\u0041\u0050"))
		if !_gcge {
			_f.Log.Debug("\u004e\u006f\u0020\u0041P \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067")
			continue
		}
		_bfd := _cg.TraceToDirectObject(_bbb.Get("\u004e"))
		if _bfd == nil {
			_f.Log.Debug("N\u006f\u0020\u004e\u0020en\u0074r\u0079\u0020\u002d\u0020\u0073k\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		var _aeb *_cg.PdfObjectStream
		switch _efe := _bfd.(type) {
		case *_cg.PdfObjectDictionary:
			_caa, _fga := _cg.GetName(_gfb.Get("\u0041\u0053"))
			if !_fga {
				_f.Log.Debug("\u004e\u006f\u0020\u0041S \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067")
				continue
			}
			_aeb, _fga = _cg.GetStream(_efe.Get(*_caa))
			if !_fga {
				_f.Log.Debug("\u0046o\u0072\u006d\u0020\u006eo\u0074\u0020\u0066\u006f\u0075n\u0064 \u002d \u0073\u006b\u0069\u0070\u0070\u0069\u006eg")
				continue
			}
		case *_cg.PdfObjectStream:
			_aeb = _efe
		}
		if _aeb == nil {
			_f.Log.Debug("\u0046\u006f\u0072m\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0028n\u0069\u006c\u0029\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		_abb, _eea := _e.NewXObjectFormFromStream(_aeb)
		if _eea != nil {
			_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _eea)
			continue
		}
		_df, _eea := _abb.GetContentStream()
		if _eea != nil {
			_f.Log.Debug("E\u0072\u0072\u006f\u0072\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0063\u006fn\u0074\u0065\u006et\u0073:\u0020\u0025\u0076", _eea)
			continue
		}
		_ecde = append(_ecde, content{_ded: string(_df), _dce: _abb.Resources})
	}
	return _ecde
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_ddg *Image) Optimize(objects []_cg.PdfObject) (_eeec []_cg.PdfObject, _gfbg error) {
	if _ddg.ImageQuality <= 0 {
		return objects, nil
	}
	_gca := _gbc(objects)
	if len(_gca) == 0 {
		return objects, nil
	}
	_bfef := make(map[_cg.PdfObject]_cg.PdfObject)
	_eae := make(map[_cg.PdfObject]struct{})
	for _, _egd := range _gca {
		_cdb := _egd.Stream.Get("\u0053\u004d\u0061s\u006b")
		_eae[_cdb] = struct{}{}
	}
	for _eag, _fcb := range _gca {
		_fae := _fcb.Stream
		if _, _cffa := _eae[_fae]; _cffa {
			continue
		}
		_eged, _egg := _e.NewXObjectImageFromStream(_fae)
		if _egg != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _egg)
			continue
		}
		switch _eged.Filter.(type) {
		case *_cg.JBIG2Encoder:
			continue
		case *_cg.CCITTFaxEncoder:
			continue
		}
		_dbb, _egg := _eged.ToImage()
		if _egg != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _egg)
			continue
		}
		_edgb := _cg.NewDCTEncoder()
		_edgb.ColorComponents = _dbb.ColorComponents
		_edgb.Quality = _ddg.ImageQuality
		_edgb.BitsPerComponent = _fcb.BitsPerComponent
		_edgb.Width = _fcb.Width
		_edgb.Height = _fcb.Height
		_dgf, _egg := _edgb.EncodeBytes(_dbb.Data)
		if _egg != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _egg)
			continue
		}
		var _gef _cg.StreamEncoder
		_gef = _edgb
		{
			_gfff := _cg.NewFlateEncoder()
			_bbf := _cg.NewMultiEncoder()
			_bbf.AddEncoder(_gfff)
			_bbf.AddEncoder(_edgb)
			_bdcf, _gae := _bbf.EncodeBytes(_dbb.Data)
			if _gae != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _gae)
				continue
			}
			if len(_bdcf) < len(_dgf) {
				_f.Log.Trace("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063\u0020\u0069\u006d\u0070\u0072\u006f\u0076\u0065\u0073\u003a\u0020\u0025\u0064\u0020\u0074o\u0020\u0025\u0064\u0020\u0028o\u0072\u0069g\u0020\u0025\u0064\u0029", len(_dgf), len(_bdcf), len(_fae.Stream))
				_dgf = _bdcf
				_gef = _bbf
			}
		}
		_ccag := len(_fae.Stream)
		if _ccag < len(_dgf) {
			continue
		}
		_acaa := &_cg.PdfObjectStream{Stream: _dgf}
		_acaa.PdfObjectReference = _fae.PdfObjectReference
		_acaa.PdfObjectDictionary = _cg.MakeDict()
		_acaa.Merge(_fae.PdfObjectDictionary)
		_acaa.Merge(_gef.MakeStreamDict())
		_acaa.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _cg.MakeInteger(int64(len(_dgf))))
		_bfef[_fae] = _acaa
		_gca[_eag].Stream = _acaa
	}
	_eeec = make([]_cg.PdfObject, len(objects))
	copy(_eeec, objects)
	_gac(_eeec, _bfef)
	return _eeec, nil
}
