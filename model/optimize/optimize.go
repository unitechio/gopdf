package optimize

import (
	_df "bytes"
	_ce "crypto/md5"
	_e "errors"
	_ac "fmt"
	_b "math"
	_a "strings"

	_f "unitechio/gopdf/gopdf/common"
	_bf "unitechio/gopdf/gopdf/contentstream"
	_ba "unitechio/gopdf/gopdf/core"
	_d "unitechio/gopdf/gopdf/extractor"
	_fg "unitechio/gopdf/gopdf/internal/imageutil"
	_cf "unitechio/gopdf/gopdf/internal/textencoding"
	_cff "unitechio/gopdf/gopdf/model"
	_gc "github.com/unidoc/unitype"
	_c "golang.org/x/image/draw"
)

// Optimize optimizes PDF objects to decrease PDF size.
func (_efeb *CleanFonts) Optimize(objects []_ba.PdfObject) (_ee []_ba.PdfObject, _cbbf error) {
	var _gec map[*_ba.PdfObjectStream]struct{}
	if _efeb.Subset {
		var _gfga error
		_gec, _gfga = _egf(objects)
		if _gfga != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004fR\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0073u\u0062s\u0065\u0074\u0074\u0069\u006e\u0067\u003a \u0025\u0076", _gfga)
			return nil, _gfga
		}
	}
	for _, _fa := range objects {
		_cgg, _gbg := _ba.GetStream(_fa)
		if !_gbg {
			continue
		}
		if _, _gcgd := _gec[_cgg]; _gcgd {
			continue
		}
		_bde, _eee := _ba.NewEncoderFromStream(_cgg)
		if _eee != nil {
			_f.Log.Debug("\u0045\u0052RO\u0052\u0020\u0067e\u0074\u0074\u0069\u006eg e\u006eco\u0064\u0065\u0072\u003a\u0020\u0025\u0076 -\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067", _eee)
			continue
		}
		_dae, _eee := _bde.DecodeStream(_cgg)
		if _eee != nil {
			_f.Log.Debug("\u0044\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065r\u0072\u006f\u0072\u0020\u003a\u0020\u0025v\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067", _eee)
			continue
		}
		if len(_dae) < 4 {
			continue
		}
		_cea := string(_dae[:4])
		if _cea == "\u004f\u0054\u0054\u004f" {
			continue
		}
		if _cea != "\u0000\u0001\u0000\u0000" && _cea != "\u0074\u0072\u0075\u0065" {
			continue
		}
		_dac, _eee := _gc.Parse(_df.NewReader(_dae))
		if _eee != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020P\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _eee)
			continue
		}
		_eee = _dac.Optimize()
		if _eee != nil {
			_f.Log.Debug("\u0045\u0052RO\u0052\u0020\u004fp\u0074\u0069\u006d\u0069zin\u0067 f\u006f\u006e\u0074\u003a\u0020\u0025\u0076 -\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067", _eee)
			continue
		}
		var _ae _df.Buffer
		_eee = _dac.Write(&_ae)
		if _eee != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020W\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _eee)
			continue
		}
		if _ae.Len() > len(_dae) {
			_f.Log.Debug("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070")
			continue
		}
		_dca, _eee := _ba.MakeStream(_ae.Bytes(), _ba.NewFlateEncoder())
		if _eee != nil {
			continue
		}
		*_cgg = *_dca
		_cgg.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _ba.MakeInteger(int64(_ae.Len())))
	}
	return objects, nil
}

// CombineIdenticalIndirectObjects combines identical indirect objects.
// It implements interface model.Optimizer.
type CombineIdenticalIndirectObjects struct{}

// Chain allows to use sequence of optimizers.
// It implements interface model.Optimizer.
type (
	Chain           struct{ _fb []_cff.Optimizer }
	objectStructure struct {
		_eefd *_ba.PdfObjectDictionary
		_fcdc *_ba.PdfObjectDictionary
		_ceg  []*_ba.PdfIndirectObject
	}
)

// CleanUnusedResources represents an optimizer used to clean unused resources.
type (
	CleanUnusedResources struct{}
	imageInfo            struct {
		BitsPerComponent int
		ColorComponents  int
		Width            int
		Height           int
		Stream           *_ba.PdfObjectStream
		PPI              float64
	}
)

// CombineDuplicateStreams combines duplicated streams by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateStreams struct{}

// ObjectStreams groups PDF objects to object streams.
// It implements interface model.Optimizer.
type ObjectStreams struct{}

// Optimize optimizes PDF objects to decrease PDF size.
func (_cb *CleanContentstream) Optimize(objects []_ba.PdfObject) (_gg []_ba.PdfObject, _cgb error) {
	_cba := map[*_ba.PdfObjectStream]struct{}{}
	var _abe []*_ba.PdfObjectStream
	_ea := func(_dd *_ba.PdfObjectStream) {
		if _, _bbea := _cba[_dd]; !_bbea {
			_cba[_dd] = struct{}{}
			_abe = append(_abe, _dd)
		}
	}
	_gce := map[_ba.PdfObject]bool{}
	_af := map[_ba.PdfObject]bool{}
	for _, _fbe := range objects {
		switch _bd := _fbe.(type) {
		case *_ba.PdfIndirectObject:
			switch _fc := _bd.PdfObject.(type) {
			case *_ba.PdfObjectDictionary:
				if _bfa, _fef := _ba.GetName(_fc.Get("\u0054\u0079\u0070\u0065")); !_fef || _bfa.String() != "\u0050\u0061\u0067\u0065" {
					continue
				}
				if _fbeg, _ddg := _ba.GetStream(_fc.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073")); _ddg {
					_ea(_fbeg)
				} else if _efe, _dad := _ba.GetArray(_fc.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073")); _dad {
					var _ag []*_ba.PdfObjectStream
					for _, _eff := range _efe.Elements() {
						if _bc, _baa := _ba.GetStream(_eff); _baa {
							_ag = append(_ag, _bc)
						}
					}
					if len(_ag) > 0 {
						var _fda _df.Buffer
						for _, _gga := range _ag {
							if _eg, _cbb := _ba.DecodeStream(_gga); _cbb == nil {
								_fda.Write(_eg)
							}
							_gce[_gga] = true
						}
						_cbc, _aag := _ba.MakeStream(_fda.Bytes(), _ba.NewFlateEncoder())
						if _aag != nil {
							return nil, _aag
						}
						_af[_cbc] = true
						_fc.Set("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073", _cbc)
						_ea(_cbc)
					}
				}
			}
		case *_ba.PdfObjectStream:
			if _adb, _acb := _ba.GetName(_bd.Get("\u0054\u0079\u0070\u0065")); !_acb || _adb.String() != "\u0058O\u0062\u006a\u0065\u0063\u0074" {
				continue
			}
			if _bcd, _gb := _ba.GetName(_bd.Get("\u0053u\u0062\u0074\u0079\u0070\u0065")); !_gb || _bcd.String() != "\u0046\u006f\u0072\u006d" {
				continue
			}
			_ea(_bd)
		}
	}
	for _, _cef := range _abe {
		_cgb = _dc(_cef)
		if _cgb != nil {
			return nil, _cgb
		}
	}
	_gg = nil
	for _, _gbe := range objects {
		if _gce[_gbe] {
			continue
		}
		_gg = append(_gg, _gbe)
	}
	for _eab := range _af {
		_gg = append(_gg, _eab)
	}
	return _gg, nil
}

func _fbdb(_dga *_cff.Image, _feb float64) (*_cff.Image, error) {
	_eac, _ceed := _dga.ToGoImage()
	if _ceed != nil {
		return nil, _ceed
	}
	var _fbbd _fg.Image
	_fceed, _agd := _eac.(*_fg.Monochrome)
	if _agd {
		if _ceed = _fceed.ResolveDecode(); _ceed != nil {
			return nil, _ceed
		}
		_fbbd, _ceed = _fceed.Scale(_feb)
		if _ceed != nil {
			return nil, _ceed
		}
	} else {
		_gbb := int(_b.RoundToEven(float64(_dga.Width) * _feb))
		_deggg := int(_b.RoundToEven(float64(_dga.Height) * _feb))
		_fbbd, _ceed = _fg.NewImage(_gbb, _deggg, int(_dga.BitsPerComponent), _dga.ColorComponents, nil, nil, nil)
		if _ceed != nil {
			return nil, _ceed
		}
		_c.CatmullRom.Scale(_fbbd, _fbbd.Bounds(), _eac, _eac.Bounds(), _c.Over, &_c.Options{})
	}
	_dfgg := _fbbd.Base()
	_cbce := &_cff.Image{Width: int64(_dfgg.Width), Height: int64(_dfgg.Height), BitsPerComponent: int64(_dfgg.BitsPerComponent), ColorComponents: _dfgg.ColorComponents, Data: _dfgg.Data}
	_cbce.SetDecode(_dfgg.Decode)
	_cbce.SetAlpha(_dfgg.Alpha)
	return _cbce, nil
}

func _dc(_fe *_ba.PdfObjectStream) error {
	_de, _da := _ba.DecodeStream(_fe)
	if _da != nil {
		return _da
	}
	_gd := _bf.NewContentStreamParser(string(_de))
	_baf, _da := _gd.Parse()
	if _da != nil {
		return _da
	}
	_baf = _gee(_baf)
	_bbe := _baf.Bytes()
	if len(_bbe) >= len(_de) {
		return nil
	}
	_ff, _da := _ba.MakeStream(_baf.Bytes(), _ba.NewFlateEncoder())
	if _da != nil {
		return _da
	}
	_fe.Stream = _ff.Stream
	_fe.Merge(_ff.PdfObjectDictionary)
	return nil
}

func _bfd(_fage []_ba.PdfObject) objectStructure {
	_ccdb := objectStructure{}
	_fffgc := false
	for _, _geef := range _fage {
		switch _eggg := _geef.(type) {
		case *_ba.PdfIndirectObject:
			_gagc, _feae := _ba.GetDict(_eggg)
			if !_feae {
				continue
			}
			_geaf, _feae := _ba.GetName(_gagc.Get("\u0054\u0079\u0070\u0065"))
			if !_feae {
				continue
			}
			switch _geaf.String() {
			case "\u0043a\u0074\u0061\u006c\u006f\u0067":
				_ccdb._eefd = _gagc
				_fffgc = true
			}
		}
		if _fffgc {
			break
		}
	}
	if !_fffgc {
		return _ccdb
	}
	_dbf, _aff := _ba.GetDict(_ccdb._eefd.Get("\u0050\u0061\u0067e\u0073"))
	if !_aff {
		return _ccdb
	}
	_ccdb._fcdc = _dbf
	_cadb, _aff := _ba.GetArray(_dbf.Get("\u004b\u0069\u0064\u0073"))
	if !_aff {
		return _ccdb
	}
	for _, _bcbf := range _cadb.Elements() {
		_ceaf, _bage := _ba.GetIndirect(_bcbf)
		if !_bage {
			break
		}
		_ccdb._ceg = append(_ccdb._ceg, _ceaf)
	}
	return _ccdb
}

// CleanFonts cleans up embedded fonts, reducing font sizes.
type CleanFonts struct {
	// Subset embedded fonts if encountered (if true).
	// Otherwise attempts to reduce the font program.
	Subset bool
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_edbd *CombineIdenticalIndirectObjects) Optimize(objects []_ba.PdfObject) (_dcf []_ba.PdfObject, _ddab error) {
	_cfgcg(objects)
	_ega := make(map[_ba.PdfObject]_ba.PdfObject)
	_egg := make(map[_ba.PdfObject]struct{})
	_cbbd := make(map[string][]*_ba.PdfIndirectObject)
	for _, _acbc := range objects {
		_acda, _cfc := _acbc.(*_ba.PdfIndirectObject)
		if !_cfc {
			continue
		}
		if _fac, _bdf := _acda.PdfObject.(*_ba.PdfObjectDictionary); _bdf {
			if _eeef, _cgc := _fac.Get("\u0054\u0079\u0070\u0065").(*_ba.PdfObjectName); _cgc && *_eeef == "\u0050\u0061\u0067\u0065" {
				continue
			}
			_eaa := _ce.New()
			_eaa.Write([]byte(_fac.WriteString()))
			_agfb := string(_eaa.Sum(nil))
			_cbbd[_agfb] = append(_cbbd[_agfb], _acda)
		}
	}
	for _, _aed := range _cbbd {
		if len(_aed) < 2 {
			continue
		}
		_gfc := _aed[0]
		for _effc := 1; _effc < len(_aed); _effc++ {
			_bba := _aed[_effc]
			_ega[_bba] = _gfc
			_egg[_bba] = struct{}{}
		}
	}
	_dcf = make([]_ba.PdfObject, 0, len(objects)-len(_egg))
	for _, _abgd := range objects {
		if _, _abec := _egg[_abgd]; _abec {
			continue
		}
		_dcf = append(_dcf, _abgd)
	}
	_gdgd(_dcf, _ega)
	return _dcf, nil
}

// ImagePPI optimizes images by scaling images such that the PPI (pixels per inch) is never higher than ImageUpperPPI.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type ImagePPI struct{ ImageUpperPPI float64 }

// CombineDuplicateDirectObjects combines duplicated direct objects by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateDirectObjects struct{}

// Image optimizes images by rewrite images into JPEG format with quality equals to ImageQuality.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type Image struct{ ImageQuality int }

func _gee(_ad *_bf.ContentStreamOperations) *_bf.ContentStreamOperations {
	if _ad == nil {
		return nil
	}
	_fd := _bf.ContentStreamOperations{}
	for _, _gfg := range *_ad {
		switch _gfg.Operand {
		case "\u0042\u0044\u0043", "\u0042\u004d\u0043", "\u0045\u004d\u0043":
			continue
		case "\u0054\u006d":
			if len(_gfg.Params) == 6 {
				if _fde, _aa := _ba.GetNumbersAsFloat(_gfg.Params); _aa == nil {
					if _fde[0] == 1 && _fde[1] == 0 && _fde[2] == 0 && _fde[3] == 1 {
						_gfg = &_bf.ContentStreamOperation{Params: []_ba.PdfObject{_gfg.Params[4], _gfg.Params[5]}, Operand: "\u0054\u0064"}
					}
				}
			}
		}
		_fd = append(_fd, _gfg)
	}
	return &_fd
}

// CleanContentstream cleans up redundant operands in content streams, including Page and XObject Form
// contents. This process includes:
// 1. Marked content operators are removed.
// 2. Some operands are simplified (shorter form).
// TODO: Add more reduction methods and improving the methods for identifying unnecessary operands.
type CleanContentstream struct{}

// Optimize optimizes PDF objects to decrease PDF size.
func (_fae *CompressStreams) Optimize(objects []_ba.PdfObject) (_agab []_ba.PdfObject, _cdfg error) {
	_agab = make([]_ba.PdfObject, len(objects))
	copy(_agab, objects)
	for _, _ddgf := range objects {
		_cefe, _gaga := _ba.GetStream(_ddgf)
		if !_gaga {
			continue
		}
		if _agaa := _cefe.Get("\u0046\u0069\u006c\u0074\u0065\u0072"); _agaa != nil {
			if _, _dgcb := _ba.GetName(_agaa); _dgcb {
				continue
			}
			if _ebf, _ggcc := _ba.GetArray(_agaa); _ggcc && _ebf.Len() > 0 {
				continue
			}
		}
		_dag := _ba.NewFlateEncoder()
		var _ccdf []byte
		_ccdf, _cdfg = _dag.EncodeBytes(_cefe.Stream)
		if _cdfg != nil {
			return _agab, _cdfg
		}
		_ebb := _dag.MakeStreamDict()
		if len(_ccdf)+len(_ebb.WriteString()) < len(_cefe.Stream) {
			_cefe.Stream = _ccdf
			_cefe.PdfObjectDictionary.Merge(_ebb)
			_cefe.PdfObjectDictionary.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _ba.MakeInteger(int64(len(_cefe.Stream))))
		}
	}
	return _agab, nil
}

func _gdac(_bdc _ba.PdfObject) (_afdg string, _bfbb []_ba.PdfObject) {
	var _aeafb _df.Buffer
	switch _eefb := _bdc.(type) {
	case *_ba.PdfIndirectObject:
		_bfbb = append(_bfbb, _eefb)
		_bdc = _eefb.PdfObject
	}
	switch _bccb := _bdc.(type) {
	case *_ba.PdfObjectStream:
		if _acbe, _gdc := _ba.DecodeStream(_bccb); _gdc == nil {
			_aeafb.Write(_acbe)
			_bfbb = append(_bfbb, _bccb)
		}
	case *_ba.PdfObjectArray:
		for _, _cadc := range _bccb.Elements() {
			switch _abgb := _cadc.(type) {
			case *_ba.PdfObjectStream:
				if _eeed, _bfeeg := _ba.DecodeStream(_abgb); _bfeeg == nil {
					_aeafb.Write(_eeed)
					_bfbb = append(_bfbb, _abgb)
				}
			}
		}
	}
	return _aeafb.String(), _bfbb
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_eb *Chain) Optimize(objects []_ba.PdfObject) (_ed []_ba.PdfObject, _gf error) {
	_cg := objects
	for _, _cd := range _eb._fb {
		_dfb, _cdf := _cd.Optimize(_cg)
		if _cdf != nil {
			_f.Log.Debug("\u0045\u0052\u0052OR\u0020\u004f\u0070\u0074\u0069\u006d\u0069\u007a\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u002b\u0076", _cdf)
			continue
		}
		_cg = _dfb
	}
	return _cg, nil
}

func _edg(_bcc *_ba.PdfObjectStream, _dfg []rune, _ggg []_gc.GlyphIndex) error {
	_bcc, _caf := _ba.GetStream(_bcc)
	if !_caf {
		_f.Log.Debug("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u002d\u002d\u0020\u0041\u0042\u004f\u0052T\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006e\u0067")
		return _e.New("\u0066\u006f\u006e\u0074fi\u006c\u0065\u0032\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_cgf, _bbd := _ba.DecodeStream(_bcc)
	if _bbd != nil {
		_f.Log.Debug("\u0044\u0065c\u006f\u0064\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _bbd)
		return _bbd
	}
	_cbg, _bbd := _gc.Parse(_df.NewReader(_cgf))
	if _bbd != nil {
		_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0020f\u006f\u006e\u0074", len(_bcc.Stream))
		return _bbd
	}
	_cfg := _ggg
	if len(_dfg) > 0 {
		_gbc := _cbg.LookupRunes(_dfg)
		_cfg = append(_cfg, _gbc...)
	}
	_cbg, _bbd = _cbg.SubsetKeepIndices(_cfg)
	if _bbd != nil {
		_f.Log.Debug("\u0045R\u0052\u004f\u0052\u0020s\u0075\u0062\u0073\u0065\u0074t\u0069n\u0067 \u0066\u006f\u006e\u0074\u003a\u0020\u0025v", _bbd)
		return _bbd
	}
	var _gda _df.Buffer
	_bbd = _cbg.Write(&_gda)
	if _bbd != nil {
		_f.Log.Debug("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _bbd)
		return _bbd
	}
	if _gda.Len() > len(_cgf) {
		_f.Log.Debug("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070")
		return nil
	}
	_gfb, _bbd := _ba.MakeStream(_gda.Bytes(), _ba.NewFlateEncoder())
	if _bbd != nil {
		_f.Log.Debug("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076", _bbd)
		return _bbd
	}
	*_bcc = *_gfb
	_bcc.Set("\u004ce\u006e\u0067\u0074\u0068\u0031", _ba.MakeInteger(int64(_gda.Len())))
	return nil
}

// New creates a optimizers chain from options.
func New(options Options) *Chain {
	_gbdb := new(Chain)
	if options.CleanFonts || options.SubsetFonts {
		_gbdb.Append(&CleanFonts{Subset: options.SubsetFonts})
	}
	if options.CleanContentstream {
		_gbdb.Append(new(CleanContentstream))
	}
	if options.ImageUpperPPI > 0 {
		_aedf := new(ImagePPI)
		_aedf.ImageUpperPPI = options.ImageUpperPPI
		_gbdb.Append(_aedf)
	}
	if options.ImageQuality > 0 {
		_bbeg := new(Image)
		_bbeg.ImageQuality = options.ImageQuality
		_gbdb.Append(_bbeg)
	}
	if options.CombineDuplicateDirectObjects {
		_gbdb.Append(new(CombineDuplicateDirectObjects))
	}
	if options.CombineDuplicateStreams {
		_gbdb.Append(new(CombineDuplicateStreams))
	}
	if options.CombineIdenticalIndirectObjects {
		_gbdb.Append(new(CombineIdenticalIndirectObjects))
	}
	if options.UseObjectStreams {
		_gbdb.Append(new(ObjectStreams))
	}
	if options.CompressStreams {
		_gbdb.Append(new(CompressStreams))
	}
	if options.CleanUnusedResources {
		_gbdb.Append(new(CleanUnusedResources))
	}
	return _gbdb
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_ceaa *CombineDuplicateDirectObjects) Optimize(objects []_ba.PdfObject) (_adc []_ba.PdfObject, _ffc error) {
	_cfgcg(objects)
	_aea := make(map[string][]*_ba.PdfObjectDictionary)
	var _gdab func(_efg *_ba.PdfObjectDictionary)
	_gdab = func(_gdff *_ba.PdfObjectDictionary) {
		for _, _fffg := range _gdff.Keys() {
			_daf := _gdff.Get(_fffg)
			if _dda, _ffd := _daf.(*_ba.PdfObjectDictionary); _ffd {
				_bcf := _ce.New()
				_bcf.Write([]byte(_dda.WriteString()))
				_dee := string(_bcf.Sum(nil))
				_aea[_dee] = append(_aea[_dee], _dda)
				_gdab(_dda)
			}
		}
	}
	for _, _eec := range objects {
		_agbf, _gcfc := _eec.(*_ba.PdfIndirectObject)
		if !_gcfc {
			continue
		}
		if _daca, _afb := _agbf.PdfObject.(*_ba.PdfObjectDictionary); _afb {
			_gdab(_daca)
		}
	}
	_bfee := make([]_ba.PdfObject, 0, len(_aea))
	_ded := make(map[_ba.PdfObject]_ba.PdfObject)
	for _, _aeb := range _aea {
		if len(_aeb) < 2 {
			continue
		}
		_ggbf := _ba.MakeDict()
		_ggbf.Merge(_aeb[0])
		_ada := _ba.MakeIndirectObject(_ggbf)
		_bfee = append(_bfee, _ada)
		for _bcff := 0; _bcff < len(_aeb); _bcff++ {
			_ggcf := _aeb[_bcff]
			_ded[_ggcf] = _ada
		}
	}
	_adc = make([]_ba.PdfObject, len(objects))
	copy(_adc, objects)
	_adc = append(_bfee, _adc...)
	_gdgd(_adc, _ded)
	return _adc, nil
}

func _cad(_bffe string, _feg []string) bool {
	for _, _gdg := range _feg {
		if _bffe == _gdg {
			return true
		}
	}
	return false
}

func _eabg(_aeg []*_ba.PdfIndirectObject) map[string][]string {
	_bdga := map[string][]string{}
	for _, _ffe := range _aeg {
		_dgg, _gdaa := _ba.GetDict(_ffe.PdfObject)
		if !_gdaa {
			continue
		}
		_efea := _dgg.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073")
		_cag := _ba.TraceToDirectObject(_efea)
		var _bgb []string
		if _eeb, _agc := _cag.(*_ba.PdfObjectArray); _agc {
			for _, _ffeg := range _eeb.Elements() {
				_afe, _ffa := _eeg(_ffeg)
				if _ffa != nil {
					continue
				}
				_bgb = append(_bgb, _afe)
			}
		}
		_gfe := _a.Join(_bgb, "\u0020")
		_beg := _bf.NewContentStreamParser(_gfe)
		_geb, _feeb := _beg.Parse()
		if _feeb != nil {
			continue
		}
		for _, _ebe := range *_geb {
			_gage := _ebe.Operand
			_cbad := _ebe.Params
			switch _gage {
			case "\u0044\u006f":
				_ggf := _cbad[0].String()
				if _, _dcaa := _bdga["\u0058O\u0062\u006a\u0065\u0063\u0074"]; !_dcaa {
					_bdga["\u0058O\u0062\u006a\u0065\u0063\u0074"] = []string{_ggf}
				} else {
					_bdga["\u0058O\u0062\u006a\u0065\u0063\u0074"] = append(_bdga["\u0058O\u0062\u006a\u0065\u0063\u0074"], _ggf)
				}
			case "\u0054\u0066":
				_gcda := _cbad[0].String()
				if _, _abf := _bdga["\u0046\u006f\u006e\u0074"]; !_abf {
					_bdga["\u0046\u006f\u006e\u0074"] = []string{_gcda}
				} else {
					_bdga["\u0046\u006f\u006e\u0074"] = append(_bdga["\u0046\u006f\u006e\u0074"], _gcda)
				}
			case "\u0067\u0073":
				_ecgb := _cbad[0].String()
				if _, _dccd := _bdga["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"]; !_dccd {
					_bdga["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"] = []string{_ecgb}
				} else {
					_bdga["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"] = append(_bdga["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"], _ecgb)
				}
			}
		}
	}
	return _bdga
}

func _egf(_gde []_ba.PdfObject) (_fca map[*_ba.PdfObjectStream]struct{}, _aagg error) {
	_fca = map[*_ba.PdfObjectStream]struct{}{}
	_fff := map[*_cff.PdfFont]struct{}{}
	_fbd := _bfd(_gde)
	for _, _aga := range _fbd._ceg {
		_ffb, _efee := _ba.GetDict(_aga.PdfObject)
		if !_efee {
			continue
		}
		_cdb, _efee := _ba.GetDict(_ffb.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
		if !_efee {
			continue
		}
		_agb, _ := _gdac(_ffb.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		_egfg, _be := _cff.NewPdfPageResourcesFromDict(_cdb)
		if _be != nil {
			return nil, _be
		}
		_bee := []content{{_gab: _agb, _fga: _egfg}}
		_fce := _ccf(_ffb.Get("\u0041\u006e\u006e\u006f\u0074\u0073"))
		if _fce != nil {
			_bee = append(_bee, _fce...)
		}
		for _, _gcf := range _bee {
			_egff, _bae := _d.NewFromContents(_gcf._gab, _gcf._fga)
			if _bae != nil {
				return nil, _bae
			}
			_cgd, _, _, _bae := _egff.ExtractPageText()
			if _bae != nil {
				return nil, _bae
			}
			for _, _fcc := range _cgd.Marks().Elements() {
				if _fcc.Font == nil {
					continue
				}
				if _, _db := _fff[_fcc.Font]; !_db {
					_fff[_fcc.Font] = struct{}{}
				}
			}
		}
	}
	_fee := map[*_ba.PdfObjectStream][]*_cff.PdfFont{}
	for _ecb := range _fff {
		_cc := _ecb.FontDescriptor()
		if _cc == nil || _cc.FontFile2 == nil {
			continue
		}
		_fceg, _bcb := _ba.GetStream(_cc.FontFile2)
		if !_bcb {
			continue
		}
		_fee[_fceg] = append(_fee[_fceg], _ecb)
	}
	for _gcg := range _fee {
		var _dcd []rune
		var _abc []_gc.GlyphIndex
		for _, _gcee := range _fee[_gcg] {
			switch _aab := _gcee.Encoder().(type) {
			case *_cf.IdentityEncoder:
				_cee := _aab.RegisteredRunes()
				_dbb := make([]_gc.GlyphIndex, len(_cee))
				for _bfe, _bab := range _cee {
					_dbb[_bfe] = _gc.GlyphIndex(_bab)
				}
				_abc = append(_abc, _dbb...)
			case *_cf.TrueTypeFontEncoder:
				_ccd := _aab.RegisteredRunes()
				_dcd = append(_dcd, _ccd...)
			case _cf.SimpleEncoder:
				_cbaf := _aab.Charcodes()
				for _, _ca := range _cbaf {
					_ddc, _egd := _aab.CharcodeToRune(_ca)
					if !_egd {
						_f.Log.Debug("\u0043\u0068a\u0072\u0063\u006f\u0064\u0065\u003c\u002d\u003e\u0072\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064: \u0025\u0064", _ca)
						continue
					}
					_dcd = append(_dcd, _ddc)
				}
			}
		}
		_aagg = _edg(_gcg, _dcd, _abc)
		if _aagg != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006eg\u0020f\u006f\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076", _aagg)
			return nil, _aagg
		}
		_fca[_gcg] = struct{}{}
	}
	return _fca, nil
}

// Optimize implements Optimizer interface.
func (_ged *CleanUnusedResources) Optimize(objects []_ba.PdfObject) (_bgg []_ba.PdfObject, _bac error) {
	_ddcd, _bac := _ade(objects)
	if _bac != nil {
		return nil, _bac
	}
	_aba := []_ba.PdfObject{}
	for _, _dgd := range objects {
		_, _bfac := _ddcd[_dgd]
		if _bfac {
			continue
		}
		_aba = append(_aba, _dgd)
	}
	return _aba, nil
}

func _edb(_gca *_ba.PdfObjectDictionary) []string {
	_fgee := []string{}
	for _, _dbgb := range _gca.Keys() {
		_fgee = append(_fgee, _dbgb.String())
	}
	return _fgee
}

func _ebaa(_dcg []_ba.PdfObject) []*imageInfo {
	_fgg := _ba.PdfObjectName("\u0053u\u0062\u0074\u0079\u0070\u0065")
	_ffea := make(map[*_ba.PdfObjectStream]struct{})
	var _babd []*imageInfo
	for _, _aebg := range _dcg {
		_bfg, _eaac := _ba.GetStream(_aebg)
		if !_eaac {
			continue
		}
		if _, _aeac := _ffea[_bfg]; _aeac {
			continue
		}
		_ffea[_bfg] = struct{}{}
		_ddb := _bfg.PdfObjectDictionary.Get(_fgg)
		_gbgd, _eaac := _ba.GetName(_ddb)
		if !_eaac || string(*_gbgd) != "\u0049\u006d\u0061g\u0065" {
			continue
		}
		_ebec := &imageInfo{Stream: _bfg, BitsPerComponent: 8}
		if _bbdd, _gbeb := _ba.GetIntVal(_bfg.Get("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074")); _gbeb {
			_ebec.BitsPerComponent = _bbdd
		}
		if _gagb, _bdb := _ba.GetIntVal(_bfg.Get("\u0057\u0069\u0064t\u0068")); _bdb {
			_ebec.Width = _gagb
		}
		if _bbbf, _cbd := _ba.GetIntVal(_bfg.Get("\u0048\u0065\u0069\u0067\u0068\u0074")); _cbd {
			_ebec.Height = _bbbf
		}
		_gabed, _ggbd := _cff.NewPdfColorspaceFromPdfObject(_bfg.Get("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"))
		if _ggbd != nil {
			_f.Log.Debug("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ggbd)
			continue
		}
		if _gabed == nil {
			_ccg, _ecdg := _ba.GetName(_bfg.Get("\u0046\u0069\u006c\u0074\u0065\u0072"))
			if _ecdg {
				switch _ccg.String() {
				case "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065", "J\u0042\u0049\u0047\u0032\u0044\u0065\u0063\u006f\u0064\u0065":
					_gabed = _cff.NewPdfColorspaceDeviceGray()
					_ebec.BitsPerComponent = 1
				}
			}
		}
		switch _ebac := _gabed.(type) {
		case *_cff.PdfColorspaceDeviceRGB:
			_ebec.ColorComponents = 3
		case *_cff.PdfColorspaceDeviceGray:
			_ebec.ColorComponents = 1
		default:
			_f.Log.Debug("\u004f\u0070\u0074\u0069\u006d\u0069\u007aa\u0074\u0069\u006fn\u0020\u0069\u0073 \u006e\u006ft\u0020\u0073\u0075\u0070\u0070\u006fr\u0074ed\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u006b\u0069\u0070", _ebac)
			continue
		}
		_babd = append(_babd, _ebec)
	}
	return _babd
}

func _ade(_cbgb []_ba.PdfObject) (map[_ba.PdfObject]struct{}, error) {
	_cae := _bfd(_cbgb)
	_aaa := _cae._ceg
	_fba := make(map[_ba.PdfObject]struct{})
	_gcc := _eabg(_aaa)
	for _, _cfbf := range _aaa {
		_bdg, _cdfd := _ba.GetDict(_cfbf.PdfObject)
		if !_cdfd {
			continue
		}
		_bfae, _cdfd := _ba.GetDict(_bdg.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
		if !_cdfd {
			continue
		}
		_cfd := _gcc["\u0058O\u0062\u006a\u0065\u0063\u0074"]
		_cdd, _cdfd := _ba.GetDict(_bfae.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
		if _cdfd {
			_dbg := _edb(_cdd)
			for _, _dbe := range _dbg {
				if _cad(_dbe, _cfd) {
					continue
				}
				_gdb := *_ba.MakeName(_dbe)
				_acc := _cdd.Get(_gdb)
				_fba[_acc] = struct{}{}
				_cdd.Remove(_gdb)
				_fbaa := _gcdb(_acc, _fba)
				if _fbaa != nil {
					_f.Log.Debug("\u0066\u0061\u0069\u006ce\u0064\u0020\u0074\u006f\u0020\u0074\u0072\u0061\u0076\u0065r\u0073e\u0020\u006f\u0062\u006a\u0065\u0063\u0074 \u0025\u0076", _acc)
				}
			}
		}
		_dab, _cdfd := _ba.GetDict(_bfae.Get("\u0046\u006f\u006e\u0074"))
		_fcf := _gcc["\u0046\u006f\u006e\u0074"]
		if _cdfd {
			_eeee := _edb(_dab)
			for _, _gac := range _eeee {
				if _cad(_gac, _fcf) {
					continue
				}
				_fddc := *_ba.MakeName(_gac)
				_ggc := _dab.Get(_fddc)
				_fba[_ggc] = struct{}{}
				_dab.Remove(_fddc)
				_geda := _gcdb(_ggc, _fba)
				if _geda != nil {
					_f.Log.Debug("\u0046\u0061i\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074 %\u0076\u000a", _ggc)
				}
			}
		}
		_cdc, _cdfd := _ba.GetDict(_bfae.Get("\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"))
		if _cdfd {
			_bda := _edb(_cdc)
			_cbe := _gcc["\u0045x\u0074\u0047\u0053\u0074\u0061\u0074e"]
			for _, _dcba := range _bda {
				if _cad(_dcba, _cbe) {
					continue
				}
				_gefg := *_ba.MakeName(_dcba)
				_gdf := _cdc.Get(_gefg)
				_fba[_gdf] = struct{}{}
				_cdc.Remove(_gefg)
				_gbcb := _gcdb(_gdf, _fba)
				if _gbcb != nil {
					_f.Log.Debug("\u0066\u0061i\u006c\u0065\u0064\u0020\u0074\u006f\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0065\u0020\u006f\u0062\u006a\u0065\u0063\u0074 %\u0076\u000a", _gdf)
				}
			}
		}
	}
	return _fba, nil
}

func _eeg(_age _ba.PdfObject) (string, error) {
	_bdgaf := _ba.TraceToDirectObject(_age)
	switch _afa := _bdgaf.(type) {
	case *_ba.PdfObjectString:
		return _afa.Str(), nil
	case *_ba.PdfObjectStream:
		_bfb, _fge := _ba.DecodeStream(_afa)
		if _fge != nil {
			return "", _fge
		}
		return string(_bfb), nil
	}
	return "", _ac.Errorf("\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0020\u0073\u0074\u0072e\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0068\u006f\u006c\u0064\u0065\u0072\u0020\u0028\u0025\u0054\u0029", _bdgaf)
}

func _cfgcg(_fccgb []_ba.PdfObject) {
	for _ggbfb, _acf := range _fccgb {
		switch _geeb := _acf.(type) {
		case *_ba.PdfIndirectObject:
			_geeb.ObjectNumber = int64(_ggbfb + 1)
			_geeb.GenerationNumber = 0
		case *_ba.PdfObjectStream:
			_geeb.ObjectNumber = int64(_ggbfb + 1)
			_geeb.GenerationNumber = 0
		case *_ba.PdfObjectStreams:
			_geeb.ObjectNumber = int64(_ggbfb + 1)
			_geeb.GenerationNumber = 0
		}
	}
}

type imageModifications struct {
	Scale    float64
	Encoding _ba.StreamEncoder
}

func _ccf(_bec _ba.PdfObject) []content {
	if _bec == nil {
		return nil
	}
	_ecd, _gag := _ba.GetArray(_bec)
	if !_gag {
		_f.Log.Debug("\u0041\u006e\u006e\u006fts\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
		return nil
	}
	var _fdd []content
	for _, _eef := range _ecd.Elements() {
		_baaf, _dccb := _ba.GetDict(_eef)
		if !_dccb {
			_f.Log.Debug("I\u0067\u006e\u006f\u0072\u0069\u006eg\u0020\u006e\u006f\u006e\u002d\u0064i\u0063\u0074\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006e\u0020\u0041\u006e\u006e\u006ft\u0073")
			continue
		}
		_bgc, _dccb := _ba.GetDict(_baaf.Get("\u0041\u0050"))
		if !_dccb {
			_f.Log.Debug("\u004e\u006f\u0020\u0041P \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067")
			continue
		}
		_efc := _ba.TraceToDirectObject(_bgc.Get("\u004e"))
		if _efc == nil {
			_f.Log.Debug("N\u006f\u0020\u004e\u0020en\u0074r\u0079\u0020\u002d\u0020\u0073k\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		var _dg *_ba.PdfObjectStream
		switch _eda := _efc.(type) {
		case *_ba.PdfObjectDictionary:
			_deg, _fea := _ba.GetName(_baaf.Get("\u0041\u0053"))
			if !_fea {
				_f.Log.Debug("\u004e\u006f\u0020\u0041S \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067")
				continue
			}
			_dg, _fea = _ba.GetStream(_eda.Get(*_deg))
			if !_fea {
				_f.Log.Debug("\u0046o\u0072\u006d\u0020\u006eo\u0074\u0020\u0066\u006f\u0075n\u0064 \u002d \u0073\u006b\u0069\u0070\u0070\u0069\u006eg")
				continue
			}
		case *_ba.PdfObjectStream:
			_dg = _eda
		}
		if _dg == nil {
			_f.Log.Debug("\u0046\u006f\u0072m\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0028n\u0069\u006c\u0029\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067")
			continue
		}
		_ecg, _eabd := _cff.NewXObjectFormFromStream(_dg)
		if _eabd != nil {
			_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067", _eabd)
			continue
		}
		_ace, _eabd := _ecg.GetContentStream()
		if _eabd != nil {
			_f.Log.Debug("E\u0072\u0072\u006f\u0072\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0063\u006fn\u0074\u0065\u006et\u0073:\u0020\u0025\u0076", _eabd)
			continue
		}
		_fdd = append(_fdd, content{_gab: string(_ace), _fga: _ecg.Resources})
	}
	return _fdd
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_eaaa *ImagePPI) Optimize(objects []_ba.PdfObject) (_fcaa []_ba.PdfObject, _fbcf error) {
	if _eaaa.ImageUpperPPI <= 0 {
		return objects, nil
	}
	_cbadc := _ebaa(objects)
	if len(_cbadc) == 0 {
		return objects, nil
	}
	_fgac := make(map[_ba.PdfObject]struct{})
	for _, _gggc := range _cbadc {
		_bag := _gggc.Stream.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b")
		_fgac[_bag] = struct{}{}
	}
	_cgca := make(map[*_ba.PdfObjectStream]*imageInfo)
	for _, _ddbd := range _cbadc {
		_cgca[_ddbd.Stream] = _ddbd
	}
	var _efd *_ba.PdfObjectDictionary
	for _, _dbde := range objects {
		if _fdea, _gdgb := _ba.GetDict(_dbde); _efd == nil && _gdgb {
			if _bccg, _aaaa := _ba.GetName(_fdea.Get("\u0054\u0079\u0070\u0065")); _aaaa && *_bccg == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
				_efd = _fdea
			}
		}
	}
	if _efd == nil {
		return objects, nil
	}
	_gabg, _efb := _ba.GetDict(_efd.Get("\u0050\u0061\u0067e\u0073"))
	if !_efb {
		return objects, nil
	}
	_eeeea, _eecf := _ba.GetArray(_gabg.Get("\u004b\u0069\u0064\u0073"))
	if !_eecf {
		return objects, nil
	}
	for _, _gece := range _eeeea.Elements() {
		_ddd := make(map[string]*imageInfo)
		_aad, _deb := _ba.GetDict(_gece)
		if !_deb {
			continue
		}
		_fbg, _ := _gdac(_aad.Get("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"))
		if len(_fbg) == 0 {
			continue
		}
		_effb, _cge := _ba.GetDict(_aad.Get("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"))
		if !_cge {
			continue
		}
		_ebdb, _gcac := _cff.NewPdfPageResourcesFromDict(_effb)
		if _gcac != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020-\u0020\u0069\u0067\u006e\u006fr\u0069\u006eg\u003a\u0020\u0025\u0076", _gcac)
			continue
		}
		_gabb, _bacg := _ba.GetDict(_effb.Get("\u0058O\u0062\u006a\u0065\u0063\u0074"))
		if !_bacg {
			continue
		}
		_dba := _gabb.Keys()
		for _, _gbf := range _dba {
			if _efebf, _bca := _ba.GetStream(_gabb.Get(_gbf)); _bca {
				if _defc, _gbgg := _cgca[_efebf]; _gbgg {
					_ddd[string(_gbf)] = _defc
				}
			}
		}
		_dgdg := _bf.NewContentStreamParser(_fbg)
		_afbc, _gcac := _dgdg.Parse()
		if _gcac != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _gcac)
			continue
		}
		_deee := _bf.NewContentStreamProcessor(*_afbc)
		_deee.AddHandler(_bf.HandlerConditionEnumAllOperands, "", func(_bdff *_bf.ContentStreamOperation, _bece _bf.GraphicsState, _acca *_cff.PdfPageResources) error {
			switch _bdff.Operand {
			case "\u0044\u006f":
				if len(_bdff.Params) != 1 {
					_f.Log.Debug("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020w\u0069\u0074\u0068\u0020\u006c\u0065\u006e\u0028\u0070\u0061ra\u006d\u0073\u0029 \u0021=\u0020\u0031")
					return nil
				}
				_bacgc, _cfa := _ba.GetName(_bdff.Params[0])
				if !_cfa {
					_f.Log.Debug("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f\u006e\u0020\u004e\u0061\u006d\u0065\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072")
					return nil
				}
				if _eaeg, _cfcg := _ddd[string(*_bacgc)]; _cfcg {
					_gdd := _bece.CTM.ScalingFactorX()
					_faf := _bece.CTM.ScalingFactorY()
					_fgd, _bge := _gdd/72.0, _faf/72.0
					_gfab, _gebb := float64(_eaeg.Width)/_fgd, float64(_eaeg.Height)/_bge
					if _fgd == 0 || _bge == 0 {
						_gfab = 72.0
						_gebb = 72.0
					}
					_eaeg.PPI = _b.Max(_eaeg.PPI, _gfab)
					_eaeg.PPI = _b.Max(_eaeg.PPI, _gebb)
				}
			}
			return nil
		})
		_gcac = _deee.Process(_ebdb)
		if _gcac != nil {
			_f.Log.Debug("E\u0052\u0052\u004f\u0052 p\u0072o\u0063\u0065\u0073\u0073\u0069n\u0067\u003a\u0020\u0025\u002b\u0076", _gcac)
			continue
		}
	}
	for _, _cffa := range _cbadc {
		if _, _fccbe := _fgac[_cffa.Stream]; _fccbe {
			continue
		}
		if _cffa.PPI <= _eaaa.ImageUpperPPI {
			continue
		}
		_gge, _aebe := _cff.NewXObjectImageFromStream(_cffa.Stream)
		if _aebe != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _aebe)
			continue
		}
		var _fag imageModifications
		_fag.Scale = _eaaa.ImageUpperPPI / _cffa.PPI
		if _cffa.BitsPerComponent == 1 && _cffa.ColorComponents == 1 {
			_bdgaa := _b.Round(_cffa.PPI / _eaaa.ImageUpperPPI)
			_cfgc := _fg.NextPowerOf2(uint(_bdgaa))
			if _fg.InDelta(float64(_cfgc), 1/_fag.Scale, 0.3) {
				_fag.Scale = float64(1) / float64(_cfgc)
			}
			if _, _abfa := _gge.Filter.(*_ba.JBIG2Encoder); !_abfa {
				_fag.Encoding = _ba.NewJBIG2Encoder()
			}
		}
		if _aebe = _gbd(_gge, _fag); _aebe != nil {
			_f.Log.Debug("\u0045\u0072\u0072\u006f\u0072 \u0073\u0063\u0061\u006c\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u006be\u0065\u0070\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _aebe)
			continue
		}
		_fag.Encoding = nil
		if _aebb, _dacg := _ba.GetStream(_cffa.Stream.PdfObjectDictionary.Get("\u0053\u004d\u0061s\u006b")); _dacg {
			_gff, _gfae := _cff.NewXObjectImageFromStream(_aebb)
			if _gfae != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _gfae)
				continue
			}
			if _gfae = _gbd(_gff, _fag); _gfae != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _gfae)
				continue
			}
		}
	}
	return objects, nil
}

func _gdgd(_caeg []_ba.PdfObject, _ecc map[_ba.PdfObject]_ba.PdfObject) {
	if len(_ecc) == 0 {
		return
	}
	for _gccb, _ceeg := range _caeg {
		if _gfag, _fgeda := _ecc[_ceeg]; _fgeda {
			_caeg[_gccb] = _gfag
			continue
		}
		_ecc[_ceeg] = _ceeg
		switch _bbfdg := _ceeg.(type) {
		case *_ba.PdfObjectArray:
			_aeaf := make([]_ba.PdfObject, _bbfdg.Len())
			copy(_aeaf, _bbfdg.Elements())
			_gdgd(_aeaf, _ecc)
			for _adg, _gcef := range _aeaf {
				_bbfdg.Set(_adg, _gcef)
			}
		case *_ba.PdfObjectStreams:
			_gdgd(_bbfdg.Elements(), _ecc)
		case *_ba.PdfObjectStream:
			_dfc := []_ba.PdfObject{_bbfdg.PdfObjectDictionary}
			_gdgd(_dfc, _ecc)
			_bbfdg.PdfObjectDictionary = _dfc[0].(*_ba.PdfObjectDictionary)
		case *_ba.PdfObjectDictionary:
			_bdaa := _bbfdg.Keys()
			_ecf := make([]_ba.PdfObject, len(_bdaa))
			for _bgce, _ccgd := range _bdaa {
				_ecf[_bgce] = _bbfdg.Get(_ccgd)
			}
			_gdgd(_ecf, _ecc)
			for _ebecd, _cfeb := range _bdaa {
				_bbfdg.Set(_cfeb, _ecf[_ebecd])
			}
		case *_ba.PdfIndirectObject:
			_defa := []_ba.PdfObject{_bbfdg.PdfObject}
			_gdgd(_defa, _ecc)
			_bbfdg.PdfObject = _defa[0]
		}
	}
}

// CompressStreams compresses uncompressed streams.
// It implements interface model.Optimizer.
type CompressStreams struct{}

func _gcdb(_gedaf _ba.PdfObject, _ceff map[_ba.PdfObject]struct{}) error {
	if _eba, _eae := _gedaf.(*_ba.PdfIndirectObject); _eae {
		_ceff[_gedaf] = struct{}{}
		_gafe := _gcdb(_eba.PdfObject, _ceff)
		if _gafe != nil {
			return _gafe
		}
		return nil
	}
	if _egb, _gagf := _gedaf.(*_ba.PdfObjectStream); _gagf {
		_ceff[_egb] = struct{}{}
		_agf := _gcdb(_egb.PdfObjectDictionary, _ceff)
		if _agf != nil {
			return _agf
		}
		return nil
	}
	if _dbd, _gfa := _gedaf.(*_ba.PdfObjectDictionary); _gfa {
		for _, _bdge := range _dbd.Keys() {
			_eefa := _dbd.Get(_bdge)
			_ = _eefa
			if _cggd, _bff := _eefa.(*_ba.PdfObjectReference); _bff {
				_eefa = _cggd.Resolve()
				_dbd.Set(_bdge, _eefa)
			}
			if _bdge != "\u0050\u0061\u0072\u0065\u006e\u0074" {
				if _geg := _gcdb(_eefa, _ceff); _geg != nil {
					return _geg
				}
			}
		}
		return nil
	}
	if _gdeb, _bace := _gedaf.(*_ba.PdfObjectArray); _bace {
		if _gdeb == nil {
			return _e.New("\u0061\u0072\u0072a\u0079\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
		}
		for _ggfa, _degg := range _gdeb.Elements() {
			if _ebc, _bbc := _degg.(*_ba.PdfObjectReference); _bbc {
				_degg = _ebc.Resolve()
				_gdeb.Set(_ggfa, _degg)
			}
			if _gfbf := _gcdb(_degg, _ceff); _gfbf != nil {
				return _gfbf
			}
		}
		return nil
	}
	return nil
}

// GetOptimizers gets the list of optimizers in chain `c`.
func (_bb *Chain) GetOptimizers() []_cff.Optimizer { return _bb._fb }

func _gbd(_dea *_cff.XObjectImage, _ffg imageModifications) error {
	_deeg, _gbgb := _dea.ToImage()
	if _gbgb != nil {
		return _gbgb
	}
	if _ffg.Scale != 0 {
		_deeg, _gbgb = _fbdb(_deeg, _ffg.Scale)
		if _gbgb != nil {
			return _gbgb
		}
	}
	if _ffg.Encoding != nil {
		_dea.Filter = _ffg.Encoding
	}
	_dea.Decode = nil
	switch _afgf := _dea.Filter.(type) {
	case *_ba.FlateEncoder:
		if _afgf.Predictor != 1 && _afgf.Predictor != 11 {
			_afgf.Predictor = 1
		}
	}
	if _gbgb = _dea.SetImage(_deeg, nil); _gbgb != nil {
		_f.Log.Debug("\u0045\u0072\u0072or\u0020\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0076", _gbgb)
		return _gbgb
	}
	_dea.ToPdfObject()
	return nil
}

// Append appends optimizers to the chain.
func (_ef *Chain) Append(optimizers ..._cff.Optimizer) { _ef._fb = append(_ef._fb, optimizers...) }

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
type content struct {
	_gab string
	_fga *_cff.PdfPageResources
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_fccf *ObjectStreams) Optimize(objects []_ba.PdfObject) (_gedb []_ba.PdfObject, _cacea error) {
	_fegg := &_ba.PdfObjectStreams{}
	_beca := make([]_ba.PdfObject, 0, len(objects))
	for _, _aeba := range objects {
		if _afc, _efdb := _aeba.(*_ba.PdfIndirectObject); _efdb && _afc.GenerationNumber == 0 {
			_fegg.Append(_aeba)
		} else {
			_beca = append(_beca, _aeba)
		}
	}
	if _fegg.Len() == 0 {
		return _beca, nil
	}
	_gedb = make([]_ba.PdfObject, 0, len(_beca)+_fegg.Len()+1)
	if _fegg.Len() > 1 {
		_gedb = append(_gedb, _fegg)
	}
	_gedb = append(_gedb, _fegg.Elements()...)
	_gedb = append(_gedb, _beca...)
	return _gedb, nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_fccd *CombineDuplicateStreams) Optimize(objects []_ba.PdfObject) (_gcea []_ba.PdfObject, _gcfa error) {
	_bbf := make(map[_ba.PdfObject]_ba.PdfObject)
	_cadg := make(map[_ba.PdfObject]struct{})
	_fged := make(map[string][]*_ba.PdfObjectStream)
	for _, _ffaa := range objects {
		if _aee, _agg := _ffaa.(*_ba.PdfObjectStream); _agg {
			_cfbd := _ce.New()
			_cfbd.Write(_aee.Stream)
			_cfbd.Write([]byte(_aee.PdfObjectDictionary.WriteString()))
			_gabe := string(_cfbd.Sum(nil))
			_fged[_gabe] = append(_fged[_gabe], _aee)
		}
	}
	for _, _beaa := range _fged {
		if len(_beaa) < 2 {
			continue
		}
		_adaa := _beaa[0]
		for _abg := 1; _abg < len(_beaa); _abg++ {
			_eaf := _beaa[_abg]
			_bbf[_eaf] = _adaa
			_cadg[_eaf] = struct{}{}
		}
	}
	_gcea = make([]_ba.PdfObject, 0, len(objects)-len(_cadg))
	for _, _aca := range objects {
		if _, _bgcg := _cadg[_aca]; _bgcg {
			continue
		}
		_gcea = append(_gcea, _aca)
	}
	_gdgd(_gcea, _bbf)
	return _gcea, nil
}

// Optimize optimizes PDF objects to decrease PDF size.
func (_fgae *Image) Optimize(objects []_ba.PdfObject) (_dcgb []_ba.PdfObject, _dfgda error) {
	if _fgae.ImageQuality <= 0 {
		return objects, nil
	}
	_fcge := _ebaa(objects)
	if len(_fcge) == 0 {
		return objects, nil
	}
	_edbf := make(map[_ba.PdfObject]_ba.PdfObject)
	_def := make(map[_ba.PdfObject]struct{})
	for _, _ebcc := range _fcge {
		_abea := _ebcc.Stream.Get("\u0053\u004d\u0061s\u006b")
		_def[_abea] = struct{}{}
	}
	for _ccgc, _ebd := range _fcge {
		_adeb := _ebd.Stream
		if _, _faaf := _def[_adeb]; _faaf {
			continue
		}
		_ebfa, _bafe := _cff.NewXObjectImageFromStream(_adeb)
		if _bafe != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _bafe)
			continue
		}
		switch _ebfa.Filter.(type) {
		case *_ba.JBIG2Encoder:
			continue
		case *_ba.CCITTFaxEncoder:
			continue
		}
		_afd, _bafe := _ebfa.ToImage()
		if _bafe != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _bafe)
			continue
		}
		_cace := _ba.NewDCTEncoder()
		_cace.ColorComponents = _afd.ColorComponents
		_cace.Quality = _fgae.ImageQuality
		_cace.BitsPerComponent = _ebd.BitsPerComponent
		_cace.Width = _ebd.Width
		_cace.Height = _ebd.Height
		_cfed, _bafe := _cace.EncodeBytes(_afd.Data)
		if _bafe != nil {
			_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _bafe)
			continue
		}
		var _ccgb _ba.StreamEncoder
		_ccgb = _cace
		{
			_bcbg := _ba.NewFlateEncoder()
			_fcgf := _ba.NewMultiEncoder()
			_fcgf.AddEncoder(_bcbg)
			_fcgf.AddEncoder(_cace)
			_caed, _bbfd := _fcgf.EncodeBytes(_afd.Data)
			if _bbfd != nil {
				_f.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076", _bbfd)
				continue
			}
			if len(_caed) < len(_cfed) {
				_f.Log.Trace("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063\u0020\u0069\u006d\u0070\u0072\u006f\u0076\u0065\u0073\u003a\u0020\u0025\u0064\u0020\u0074o\u0020\u0025\u0064\u0020\u0028o\u0072\u0069g\u0020\u0025\u0064\u0029", len(_cfed), len(_caed), len(_adeb.Stream))
				_cfed = _caed
				_ccgb = _fcgf
			}
		}
		_eecg := len(_adeb.Stream)
		if _eecg < len(_cfed) {
			continue
		}
		_ebff := &_ba.PdfObjectStream{Stream: _cfed}
		_ebff.PdfObjectReference = _adeb.PdfObjectReference
		_ebff.PdfObjectDictionary = _ba.MakeDict()
		_ebff.Merge(_adeb.PdfObjectDictionary)
		_ebff.Merge(_ccgb.MakeStreamDict())
		_ebff.Set("\u004c\u0065\u006e\u0067\u0074\u0068", _ba.MakeInteger(int64(len(_cfed))))
		_edbf[_adeb] = _ebff
		_fcge[_ccgc].Stream = _ebff
	}
	_dcgb = make([]_ba.PdfObject, len(objects))
	copy(_dcgb, objects)
	_gdgd(_dcgb, _edbf)
	return _dcgb, nil
}
