package segments

import (
	_ag "encoding/binary"
	_b "errors"
	_af "fmt"
	_a "image"
	_ed "io"
	_e "math"
	_d "strings"
	_cf "time"

	_aae "bitbucket.org/shenghui0779/gopdf/common"
	_g "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_bg "bitbucket.org/shenghui0779/gopdf/internal/jbig2/basic"
	_gf "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_df "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder/arithmetic"
	_fa "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder/huffman"
	_eb "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder/mmr"
	_afe "bitbucket.org/shenghui0779/gopdf/internal/jbig2/encoder/arithmetic"
	_edb "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
	_f "bitbucket.org/shenghui0779/gopdf/internal/jbig2/internal"
	_aa "golang.org/x/xerrors"
)

func (_gdgc *TextRegion) decodeStripT() (_gace int64, _cfdf error) {
	if _gdgc.IsHuffmanEncoded {
		if _gdgc.SbHuffDT == 3 {
			if _gdgc._bbfg == nil {
				var _abgcc int
				if _gdgc.SbHuffFS == 3 {
					_abgcc++
				}
				if _gdgc.SbHuffDS == 3 {
					_abgcc++
				}
				_gdgc._bbfg, _cfdf = _gdgc.getUserTable(_abgcc)
				if _cfdf != nil {
					return 0, _cfdf
				}
			}
			_gace, _cfdf = _gdgc._bbfg.Decode(_gdgc._ecc)
			if _cfdf != nil {
				return 0, _cfdf
			}
		} else {
			var _eefe _fa.Tabler
			_eefe, _cfdf = _fa.GetStandardTable(11 + int(_gdgc.SbHuffDT))
			if _cfdf != nil {
				return 0, _cfdf
			}
			_gace, _cfdf = _eefe.Decode(_gdgc._ecc)
			if _cfdf != nil {
				return 0, _cfdf
			}
		}
	} else {
		var _cegf int32
		_cegf, _cfdf = _gdgc._egdc.DecodeInt(_gdgc._gede)
		if _cfdf != nil {
			return 0, _cfdf
		}
		_gace = int64(_cegf)
	}
	_gace *= int64(-_gdgc.SbStrips)
	return _gace, nil
}
func (_eeac *SymbolDictionary) setRefinementAtPixels() error {
	if !_eeac.UseRefinementAggregation || _eeac.SdrTemplate != 0 {
		return nil
	}
	if _aeac := _eeac.readRefinementAtPixels(2); _aeac != nil {
		return _aeac
	}
	return nil
}
func (_aede *SymbolDictionary) retrieveImportSymbols() error {
	for _, _beacd := range _aede.Header.RTSegments {
		if _beacd.Type == 0 {
			_defbf, _cbabg := _beacd.GetSegmentData()
			if _cbabg != nil {
				return _cbabg
			}
			_gcg, _cee := _defbf.(*SymbolDictionary)
			if !_cee {
				return _af.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0044\u0061\u0074a\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0053\u0065\u0067m\u0065\u006e\u0074\u003a\u0020%\u0054", _defbf)
			}
			_dafgc, _cbabg := _gcg.GetDictionary()
			if _cbabg != nil {
				return _af.Errorf("\u0072\u0065\u006c\u0061\u0074\u0065\u0064 \u0073\u0065\u0067m\u0065\u006e\u0074 \u0077\u0069t\u0068\u0020\u0069\u006e\u0064\u0065x\u003a %\u0064\u0020\u0067\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0073", _beacd.SegmentNumber, _cbabg.Error())
			}
			_aede._bgbd = append(_aede._bgbd, _dafgc...)
			_aede._gfg += _gcg.NumberOfExportedSymbols
		}
	}
	return nil
}
func NewRegionSegment(r *_g.Reader) *RegionSegment { return &RegionSegment{_ebecg: r} }

type template0 struct{}

func (_aace *Header) writeReferredToCount(_fbab _g.BinaryWriter) (_aeae int, _ggdd error) {
	const _dea = "w\u0072i\u0074\u0065\u0052\u0065\u0066\u0065\u0072\u0072e\u0064\u0054\u006f\u0043ou\u006e\u0074"
	_aace.RTSNumbers = make([]int, len(_aace.RTSegments))
	for _dfca, _bffc := range _aace.RTSegments {
		_aace.RTSNumbers[_dfca] = int(_bffc.SegmentNumber)
	}
	if len(_aace.RTSNumbers) <= 4 {
		var _abge byte
		if len(_aace.RetainBits) >= 1 {
			_abge = _aace.RetainBits[0]
		}
		_abge |= byte(len(_aace.RTSNumbers)) << 5
		if _ggdd = _fbab.WriteByte(_abge); _ggdd != nil {
			return 0, _edb.Wrap(_ggdd, _dea, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		return 1, nil
	}
	_gfee := uint32(len(_aace.RTSNumbers))
	_cega := make([]byte, 4+_bg.Ceil(len(_aace.RTSNumbers)+1, 8))
	_gfee |= 0x7 << 29
	_ag.BigEndian.PutUint32(_cega, _gfee)
	copy(_cega[1:], _aace.RetainBits)
	_aeae, _ggdd = _fbab.Write(_cega)
	if _ggdd != nil {
		return 0, _edb.Wrap(_ggdd, _dea, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	return _aeae, nil
}
func (_eedg *SymbolDictionary) huffDecodeRefAggNInst() (int64, error) {
	if !_eedg.SdHuffAggInstanceSelection {
		_cbgb, _efcg := _fa.GetStandardTable(1)
		if _efcg != nil {
			return 0, _efcg
		}
		return _cbgb.Decode(_eedg._caba)
	}
	if _eedg._ecef == nil {
		var (
			_becdd int
			_affe  error
		)
		if _eedg.SdHuffDecodeHeightSelection == 3 {
			_becdd++
		}
		if _eedg.SdHuffDecodeWidthSelection == 3 {
			_becdd++
		}
		if _eedg.SdHuffBMSizeSelection == 3 {
			_becdd++
		}
		_eedg._ecef, _affe = _eedg.getUserTable(_becdd)
		if _affe != nil {
			return 0, _affe
		}
	}
	return _eedg._ecef.Decode(_eedg._caba)
}
func (_age *GenericRegion) overrideAtTemplate0b(_aed, _dedg, _acf, _egd, _ebegb, _aaed int) int {
	if _age.GBAtOverride[0] {
		_aed &= 0xFFFD
		if _age.GBAtY[0] == 0 && _age.GBAtX[0] >= -int8(_ebegb) {
			_aed |= (_egd >> uint(int8(_aaed)-_age.GBAtX[0]&0x1)) << 1
		} else {
			_aed |= int(_age.getPixel(_dedg+int(_age.GBAtX[0]), _acf+int(_age.GBAtY[0]))) << 1
		}
	}
	if _age.GBAtOverride[1] {
		_aed &= 0xDFFF
		if _age.GBAtY[1] == 0 && _age.GBAtX[1] >= -int8(_ebegb) {
			_aed |= (_egd >> uint(int8(_aaed)-_age.GBAtX[1]&0x1)) << 13
		} else {
			_aed |= int(_age.getPixel(_dedg+int(_age.GBAtX[1]), _acf+int(_age.GBAtY[1]))) << 13
		}
	}
	if _age.GBAtOverride[2] {
		_aed &= 0xFDFF
		if _age.GBAtY[2] == 0 && _age.GBAtX[2] >= -int8(_ebegb) {
			_aed |= (_egd >> uint(int8(_aaed)-_age.GBAtX[2]&0x1)) << 9
		} else {
			_aed |= int(_age.getPixel(_dedg+int(_age.GBAtX[2]), _acf+int(_age.GBAtY[2]))) << 9
		}
	}
	if _age.GBAtOverride[3] {
		_aed &= 0xBFFF
		if _age.GBAtY[3] == 0 && _age.GBAtX[3] >= -int8(_ebegb) {
			_aed |= (_egd >> uint(int8(_aaed)-_age.GBAtX[3]&0x1)) << 14
		} else {
			_aed |= int(_age.getPixel(_dedg+int(_age.GBAtX[3]), _acf+int(_age.GBAtY[3]))) << 14
		}
	}
	if _age.GBAtOverride[4] {
		_aed &= 0xEFFF
		if _age.GBAtY[4] == 0 && _age.GBAtX[4] >= -int8(_ebegb) {
			_aed |= (_egd >> uint(int8(_aaed)-_age.GBAtX[4]&0x1)) << 12
		} else {
			_aed |= int(_age.getPixel(_dedg+int(_age.GBAtX[4]), _acf+int(_age.GBAtY[4]))) << 12
		}
	}
	if _age.GBAtOverride[5] {
		_aed &= 0xFFDF
		if _age.GBAtY[5] == 0 && _age.GBAtX[5] >= -int8(_ebegb) {
			_aed |= (_egd >> uint(int8(_aaed)-_age.GBAtX[5]&0x1)) << 5
		} else {
			_aed |= int(_age.getPixel(_dedg+int(_age.GBAtX[5]), _acf+int(_age.GBAtY[5]))) << 5
		}
	}
	if _age.GBAtOverride[6] {
		_aed &= 0xFFFB
		if _age.GBAtY[6] == 0 && _age.GBAtX[6] >= -int8(_ebegb) {
			_aed |= (_egd >> uint(int8(_aaed)-_age.GBAtX[6]&0x1)) << 2
		} else {
			_aed |= int(_age.getPixel(_dedg+int(_age.GBAtX[6]), _acf+int(_age.GBAtY[6]))) << 2
		}
	}
	if _age.GBAtOverride[7] {
		_aed &= 0xFFF7
		if _age.GBAtY[7] == 0 && _age.GBAtX[7] >= -int8(_ebegb) {
			_aed |= (_egd >> uint(int8(_aaed)-_age.GBAtX[7]&0x1)) << 3
		} else {
			_aed |= int(_age.getPixel(_dedg+int(_age.GBAtX[7]), _acf+int(_age.GBAtY[7]))) << 3
		}
	}
	if _age.GBAtOverride[8] {
		_aed &= 0xF7FF
		if _age.GBAtY[8] == 0 && _age.GBAtX[8] >= -int8(_ebegb) {
			_aed |= (_egd >> uint(int8(_aaed)-_age.GBAtX[8]&0x1)) << 11
		} else {
			_aed |= int(_age.getPixel(_dedg+int(_age.GBAtX[8]), _acf+int(_age.GBAtY[8]))) << 11
		}
	}
	if _age.GBAtOverride[9] {
		_aed &= 0xFFEF
		if _age.GBAtY[9] == 0 && _age.GBAtX[9] >= -int8(_ebegb) {
			_aed |= (_egd >> uint(int8(_aaed)-_age.GBAtX[9]&0x1)) << 4
		} else {
			_aed |= int(_age.getPixel(_dedg+int(_age.GBAtX[9]), _acf+int(_age.GBAtY[9]))) << 4
		}
	}
	if _age.GBAtOverride[10] {
		_aed &= 0x7FFF
		if _age.GBAtY[10] == 0 && _age.GBAtX[10] >= -int8(_ebegb) {
			_aed |= (_egd >> uint(int8(_aaed)-_age.GBAtX[10]&0x1)) << 15
		} else {
			_aed |= int(_age.getPixel(_dedg+int(_age.GBAtX[10]), _acf+int(_age.GBAtY[10]))) << 15
		}
	}
	if _age.GBAtOverride[11] {
		_aed &= 0xFDFF
		if _age.GBAtY[11] == 0 && _age.GBAtX[11] >= -int8(_ebegb) {
			_aed |= (_egd >> uint(int8(_aaed)-_age.GBAtX[11]&0x1)) << 10
		} else {
			_aed |= int(_age.getPixel(_dedg+int(_age.GBAtX[11]), _acf+int(_age.GBAtY[11]))) << 10
		}
	}
	return _aed
}
func (_fgbf *HalftoneRegion) combineGrayscalePlanes(_aec []*_gf.Bitmap, _ebdd int) error {
	_dbe := 0
	for _bacd := 0; _bacd < _aec[_ebdd].Height; _bacd++ {
		for _faece := 0; _faece < _aec[_ebdd].Width; _faece += 8 {
			_bbab, _aeeg := _aec[_ebdd+1].GetByte(_dbe)
			if _aeeg != nil {
				return _aeeg
			}
			_dcf, _aeeg := _aec[_ebdd].GetByte(_dbe)
			if _aeeg != nil {
				return _aeeg
			}
			_aeeg = _aec[_ebdd].SetByte(_dbe, _gf.CombineBytes(_dcf, _bbab, _gf.CmbOpXor))
			if _aeeg != nil {
				return _aeeg
			}
			_dbe++
		}
	}
	return nil
}
func (_dfdae *PatternDictionary) extractPatterns(_defa *_gf.Bitmap) error {
	var _eefg int
	_effd := make([]*_gf.Bitmap, _dfdae.GrayMax+1)
	for _eefg <= int(_dfdae.GrayMax) {
		_ggff := int(_dfdae.HdpWidth) * _eefg
		_dfdag := _a.Rect(_ggff, 0, _ggff+int(_dfdae.HdpWidth), int(_dfdae.HdpHeight))
		_aaae, _gfda := _gf.Extract(_dfdag, _defa)
		if _gfda != nil {
			return _gfda
		}
		_effd[_eefg] = _aaae
		_eefg++
	}
	_dfdae.Patterns = _effd
	return nil
}
func (_ebae *HalftoneRegion) Init(hd *Header, r *_g.Reader) error {
	_ebae._fagd = r
	_ebae._ceag = hd
	_ebae.RegionSegment = NewRegionSegment(r)
	return _ebae.parseHeader()
}

type GenericRefinementRegion struct {
	_fea            templater
	_ad             templater
	_dd             *_g.Reader
	_cb             *Header
	RegionInfo      *RegionSegment
	IsTPGROn        bool
	TemplateID      int8
	Template        templater
	GrAtX           []int8
	GrAtY           []int8
	RegionBitmap    *_gf.Bitmap
	ReferenceBitmap *_gf.Bitmap
	ReferenceDX     int32
	ReferenceDY     int32
	_dc             *_df.Decoder
	_cd             *_df.DecoderStats
	_fb             bool
	_ce             []bool
}

func (_gffg *TextRegion) encodeFlags(_cbdg _g.BinaryWriter) (_cbed int, _bbbg error) {
	const _gbccg = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _bbbg = _cbdg.WriteBit(int(_gffg.SbrTemplate)); _bbbg != nil {
		return _cbed, _edb.Wrap(_bbbg, _gbccg, "s\u0062\u0072\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	if _, _bbbg = _cbdg.WriteBits(uint64(_gffg.SbDsOffset), 5); _bbbg != nil {
		return _cbed, _edb.Wrap(_bbbg, _gbccg, "\u0073\u0062\u0044\u0073\u004f\u0066\u0066\u0073\u0065\u0074")
	}
	if _bbbg = _cbdg.WriteBit(int(_gffg.DefaultPixel)); _bbbg != nil {
		return _cbed, _edb.Wrap(_bbbg, _gbccg, "\u0044\u0065\u0066a\u0075\u006c\u0074\u0050\u0069\u0078\u0065\u006c")
	}
	if _, _bbbg = _cbdg.WriteBits(uint64(_gffg.CombinationOperator), 2); _bbbg != nil {
		return _cbed, _edb.Wrap(_bbbg, _gbccg, "\u0043\u006f\u006d\u0062in\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	if _bbbg = _cbdg.WriteBit(int(_gffg.IsTransposed)); _bbbg != nil {
		return _cbed, _edb.Wrap(_bbbg, _gbccg, "\u0069\u0073\u0020\u0074\u0072\u0061\u006e\u0073\u0070\u006f\u0073\u0065\u0064")
	}
	if _, _bbbg = _cbdg.WriteBits(uint64(_gffg.ReferenceCorner), 2); _bbbg != nil {
		return _cbed, _edb.Wrap(_bbbg, _gbccg, "\u0072\u0065f\u0065\u0072\u0065n\u0063\u0065\u0020\u0063\u006f\u0072\u006e\u0065\u0072")
	}
	if _, _bbbg = _cbdg.WriteBits(uint64(_gffg.LogSBStrips), 2); _bbbg != nil {
		return _cbed, _edb.Wrap(_bbbg, _gbccg, "L\u006f\u0067\u0053\u0042\u0053\u0074\u0072\u0069\u0070\u0073")
	}
	var _bddbe int
	if _gffg.UseRefinement {
		_bddbe = 1
	}
	if _bbbg = _cbdg.WriteBit(_bddbe); _bbbg != nil {
		return _cbed, _edb.Wrap(_bbbg, _gbccg, "\u0075\u0073\u0065\u0020\u0072\u0065\u0066\u0069\u006ee\u006d\u0065\u006e\u0074")
	}
	_bddbe = 0
	if _gffg.IsHuffmanEncoded {
		_bddbe = 1
	}
	if _bbbg = _cbdg.WriteBit(_bddbe); _bbbg != nil {
		return _cbed, _edb.Wrap(_bbbg, _gbccg, "u\u0073\u0065\u0020\u0068\u0075\u0066\u0066\u006d\u0061\u006e")
	}
	_cbed = 2
	return _cbed, nil
}
func (_bee *PageInformationSegment) readMaxStripeSize() error {
	_becd, _egfe := _bee._faef.ReadBits(15)
	if _egfe != nil {
		return _egfe
	}
	_bee.MaxStripeSize = uint16(_becd & _e.MaxUint16)
	return nil
}
func (_bfbg *RegionSegment) Encode(w _g.BinaryWriter) (_dbfb int, _ebg error) {
	const _aedc = "R\u0065g\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065n\u0074\u002e\u0045\u006eco\u0064\u0065"
	_abbg := make([]byte, 4)
	_ag.BigEndian.PutUint32(_abbg, _bfbg.BitmapWidth)
	_dbfb, _ebg = w.Write(_abbg)
	if _ebg != nil {
		return 0, _edb.Wrap(_ebg, _aedc, "\u0057\u0069\u0064t\u0068")
	}
	_ag.BigEndian.PutUint32(_abbg, _bfbg.BitmapHeight)
	var _ffbd int
	_ffbd, _ebg = w.Write(_abbg)
	if _ebg != nil {
		return 0, _edb.Wrap(_ebg, _aedc, "\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_dbfb += _ffbd
	_ag.BigEndian.PutUint32(_abbg, _bfbg.XLocation)
	_ffbd, _ebg = w.Write(_abbg)
	if _ebg != nil {
		return 0, _edb.Wrap(_ebg, _aedc, "\u0058L\u006f\u0063\u0061\u0074\u0069\u006fn")
	}
	_dbfb += _ffbd
	_ag.BigEndian.PutUint32(_abbg, _bfbg.YLocation)
	_ffbd, _ebg = w.Write(_abbg)
	if _ebg != nil {
		return 0, _edb.Wrap(_ebg, _aedc, "\u0059L\u006f\u0063\u0061\u0074\u0069\u006fn")
	}
	_dbfb += _ffbd
	if _ebg = w.WriteByte(byte(_bfbg.CombinaionOperator) & 0x07); _ebg != nil {
		return 0, _edb.Wrap(_ebg, _aedc, "c\u006fm\u0062\u0069\u006e\u0061\u0074\u0069\u006f\u006e \u006f\u0070\u0065\u0072at\u006f\u0072")
	}
	_dbfb++
	return _dbfb, nil
}

var (
	_ebba Segmenter
	_ggdf = map[Type]func() Segmenter{TSymbolDictionary: func() Segmenter { return &SymbolDictionary{} }, TIntermediateTextRegion: func() Segmenter { return &TextRegion{} }, TImmediateTextRegion: func() Segmenter { return &TextRegion{} }, TImmediateLosslessTextRegion: func() Segmenter { return &TextRegion{} }, TPatternDictionary: func() Segmenter { return &PatternDictionary{} }, TIntermediateHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TImmediateHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TImmediateLosslessHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TIntermediateGenericRegion: func() Segmenter { return &GenericRegion{} }, TImmediateGenericRegion: func() Segmenter { return &GenericRegion{} }, TImmediateLosslessGenericRegion: func() Segmenter { return &GenericRegion{} }, TIntermediateGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TImmediateGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TImmediateLosslessGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TPageInformation: func() Segmenter { return &PageInformationSegment{} }, TEndOfPage: func() Segmenter { return _ebba }, TEndOfStrip: func() Segmenter { return &EndOfStripe{} }, TEndOfFile: func() Segmenter { return _ebba }, TProfiles: func() Segmenter { return _ebba }, TTables: func() Segmenter { return &TableSegment{} }, TExtension: func() Segmenter { return _ebba }, TBitmap: func() Segmenter { return _ebba }}
)

type PageInformationSegment struct {
	_faef             *_g.Reader
	PageBMHeight      int
	PageBMWidth       int
	ResolutionX       int
	ResolutionY       int
	_fbbge            bool
	_dbcd             _gf.CombinationOperator
	_cbdc             bool
	DefaultPixelValue uint8
	_ccdd             bool
	IsLossless        bool
	IsStripe          bool
	MaxStripeSize     uint16
}

func (_fbaf *PageInformationSegment) readResolution() error {
	_aaff, _cdge := _fbaf._faef.ReadBits(32)
	if _cdge != nil {
		return _cdge
	}
	_fbaf.ResolutionX = int(_aaff & _e.MaxInt32)
	_aaff, _cdge = _fbaf._faef.ReadBits(32)
	if _cdge != nil {
		return _cdge
	}
	_fbaf.ResolutionY = int(_aaff & _e.MaxInt32)
	return nil
}
func (_egb *SymbolDictionary) Encode(w _g.BinaryWriter) (_edda int, _cgda error) {
	const _aeca = "\u0053\u0079\u006dbo\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	if _egb == nil {
		return 0, _edb.Error(_aeca, "\u0073\u0079m\u0062\u006f\u006c\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	if _edda, _cgda = _egb.encodeFlags(w); _cgda != nil {
		return _edda, _edb.Wrap(_cgda, _aeca, "")
	}
	_fega, _cgda := _egb.encodeATFlags(w)
	if _cgda != nil {
		return _edda, _edb.Wrap(_cgda, _aeca, "")
	}
	_edda += _fega
	if _fega, _cgda = _egb.encodeRefinementATFlags(w); _cgda != nil {
		return _edda, _edb.Wrap(_cgda, _aeca, "")
	}
	_edda += _fega
	if _fega, _cgda = _egb.encodeNumSyms(w); _cgda != nil {
		return _edda, _edb.Wrap(_cgda, _aeca, "")
	}
	_edda += _fega
	if _fega, _cgda = _egb.encodeSymbols(w); _cgda != nil {
		return _edda, _edb.Wrap(_cgda, _aeca, "")
	}
	_edda += _fega
	return _edda, nil
}
func (_cfc *GenericRegion) computeSegmentDataStructure() error {
	_cfc.DataOffset = _cfc._efeg.AbsolutePosition()
	_cfc.DataHeaderLength = _cfc.DataOffset - _cfc.DataHeaderOffset
	_cfc.DataLength = int64(_cfc._efeg.AbsoluteLength()) - _cfc.DataHeaderLength
	return nil
}
func (_fbbg *Header) readDataStartOffset(_acdb *_g.Reader, _fbfb OrganizationType) {
	if _fbfb == OSequential {
		_fbbg.SegmentDataStartOffset = uint64(_acdb.AbsolutePosition())
	}
}
func (_gec *HalftoneRegion) grayScaleDecoding(_baff int) ([][]int, error) {
	var (
		_dfbd []int8
		_ged  []int8
	)
	if !_gec.IsMMREncoded {
		_dfbd = make([]int8, 4)
		_ged = make([]int8, 4)
		if _gec.HTemplate <= 1 {
			_dfbd[0] = 3
		} else if _gec.HTemplate >= 2 {
			_dfbd[0] = 2
		}
		_ged[0] = -1
		_dfbd[1] = -3
		_ged[1] = -1
		_dfbd[2] = 2
		_ged[2] = -2
		_dfbd[3] = -2
		_ged[3] = -2
	}
	_ceca := make([]*_gf.Bitmap, _baff)
	_dggc := NewGenericRegion(_gec._fagd)
	_dggc.setParametersMMR(_gec.IsMMREncoded, _gec.DataOffset, _gec.DataLength, _gec.HGridHeight, _gec.HGridWidth, _gec.HTemplate, false, _gec.HSkipEnabled, _dfbd, _ged)
	_adca := _baff - 1
	var _ddb error
	_ceca[_adca], _ddb = _dggc.GetRegionBitmap()
	if _ddb != nil {
		return nil, _ddb
	}
	for _adca > 0 {
		_adca--
		_dggc.Bitmap = nil
		_ceca[_adca], _ddb = _dggc.GetRegionBitmap()
		if _ddb != nil {
			return nil, _ddb
		}
		if _ddb = _gec.combineGrayscalePlanes(_ceca, _adca); _ddb != nil {
			return nil, _ddb
		}
	}
	return _gec.computeGrayScalePlanes(_ceca, _baff)
}
func (_cba *GenericRegion) setParameters(_cffg bool, _ebb, _agd int64, _debf, _aacf uint32) {
	_cba.IsMMREncoded = _cffg
	_cba.DataOffset = _ebb
	_cba.DataLength = _agd
	_cba.RegionSegment.BitmapHeight = _debf
	_cba.RegionSegment.BitmapWidth = _aacf
	_cba._gdf = nil
	_cba.Bitmap = nil
}
func (_fgd *GenericRefinementRegion) getPixel(_fcd *_gf.Bitmap, _bfaa, _ded int) int {
	if _bfaa < 0 || _bfaa >= _fcd.Width {
		return 0
	}
	if _ded < 0 || _ded >= _fcd.Height {
		return 0
	}
	if _fcd.GetPixel(_bfaa, _ded) {
		return 1
	}
	return 0
}
func (_ggfg *Header) writeSegmentPageAssociation(_adaa _g.BinaryWriter) (_ddea int, _gad error) {
	const _dbaa = "w\u0072\u0069\u0074\u0065\u0053\u0065g\u006d\u0065\u006e\u0074\u0050\u0061\u0067\u0065\u0041s\u0073\u006f\u0063i\u0061t\u0069\u006f\u006e"
	if _ggfg.pageSize() != 4 {
		if _gad = _adaa.WriteByte(byte(_ggfg.PageAssociation)); _gad != nil {
			return 0, _edb.Wrap(_gad, _dbaa, "\u0070\u0061\u0067\u0065\u0053\u0069\u007a\u0065\u0020\u0021\u003d\u0020\u0034")
		}
		return 1, nil
	}
	_ebbe := make([]byte, 4)
	_ag.BigEndian.PutUint32(_ebbe, uint32(_ggfg.PageAssociation))
	if _ddea, _gad = _adaa.Write(_ebbe); _gad != nil {
		return 0, _edb.Wrap(_gad, _dbaa, "\u0034 \u0062y\u0074\u0065\u0020\u0070\u0061g\u0065\u0020n\u0075\u006d\u0062\u0065\u0072")
	}
	return _ddea, nil
}
func (_edgb *PatternDictionary) readPatternWidthAndHeight() error {
	_fgcb, _cgeg := _edgb._dbec.ReadByte()
	if _cgeg != nil {
		return _cgeg
	}
	_edgb.HdpWidth = _fgcb
	_fgcb, _cgeg = _edgb._dbec.ReadByte()
	if _cgeg != nil {
		return _cgeg
	}
	_edgb.HdpHeight = _fgcb
	return nil
}
func (_acde *PageInformationSegment) CombinationOperator() _gf.CombinationOperator {
	return _acde._dbcd
}
func (_ccd *HalftoneRegion) GetRegionBitmap() (*_gf.Bitmap, error) {
	if _ccd.HalftoneRegionBitmap != nil {
		return _ccd.HalftoneRegionBitmap, nil
	}
	var _agec error
	_ccd.HalftoneRegionBitmap = _gf.New(int(_ccd.RegionSegment.BitmapWidth), int(_ccd.RegionSegment.BitmapHeight))
	if _ccd.Patterns == nil || len(_ccd.Patterns) == 0 {
		_ccd.Patterns, _agec = _ccd.GetPatterns()
		if _agec != nil {
			return nil, _agec
		}
	}
	if _ccd.HDefaultPixel == 1 {
		_ccd.HalftoneRegionBitmap.SetDefaultPixel()
	}
	_bdf := _e.Ceil(_e.Log(float64(len(_ccd.Patterns))) / _e.Log(2))
	_aade := int(_bdf)
	var _edcb [][]int
	_edcb, _agec = _ccd.grayScaleDecoding(_aade)
	if _agec != nil {
		return nil, _agec
	}
	if _agec = _ccd.renderPattern(_edcb); _agec != nil {
		return nil, _agec
	}
	return _ccd.HalftoneRegionBitmap, nil
}
func (_dagf *PatternDictionary) readTemplate() error {
	_gcda, _aaac := _dagf._dbec.ReadBits(2)
	if _aaac != nil {
		return _aaac
	}
	_dagf.HDTemplate = byte(_gcda)
	return nil
}
func (_cdeg *TextRegion) readUseRefinement() error {
	if !_cdeg.UseRefinement || _cdeg.SbrTemplate != 0 {
		return nil
	}
	var (
		_aeag byte
		_bcfb error
	)
	_cdeg.SbrATX = make([]int8, 2)
	_cdeg.SbrATY = make([]int8, 2)
	_aeag, _bcfb = _cdeg._ecc.ReadByte()
	if _bcfb != nil {
		return _bcfb
	}
	_cdeg.SbrATX[0] = int8(_aeag)
	_aeag, _bcfb = _cdeg._ecc.ReadByte()
	if _bcfb != nil {
		return _bcfb
	}
	_cdeg.SbrATY[0] = int8(_aeag)
	_aeag, _bcfb = _cdeg._ecc.ReadByte()
	if _bcfb != nil {
		return _bcfb
	}
	_cdeg.SbrATX[1] = int8(_aeag)
	_aeag, _bcfb = _cdeg._ecc.ReadByte()
	if _bcfb != nil {
		return _bcfb
	}
	_cdeg.SbrATY[1] = int8(_aeag)
	return nil
}
func (_aea *HalftoneRegion) shiftAndFill(_bffg int) int {
	_bffg >>= 8
	if _bffg < 0 {
		_abbc := int(_e.Log(float64(_ffbc(_bffg))) / _e.Log(2))
		_agff := 31 - _abbc
		for _dfgaf := 1; _dfgaf < _agff; _dfgaf++ {
			_bffg |= 1 << uint(31-_dfgaf)
		}
	}
	return _bffg
}
func (_ceba *SymbolDictionary) setAtPixels() error {
	if _ceba.IsHuffmanEncoded {
		return nil
	}
	_gggb := 1
	if _ceba.SdTemplate == 0 {
		_gggb = 4
	}
	if _dbda := _ceba.readAtPixels(_gggb); _dbda != nil {
		return _dbda
	}
	return nil
}
func (_fcef *TextRegion) decodeCurrentT() (int64, error) {
	if _fcef.SbStrips != 1 {
		if _fcef.IsHuffmanEncoded {
			_daagb, _ecgf := _fcef._ecc.ReadBits(byte(_fcef.LogSBStrips))
			return int64(_daagb), _ecgf
		}
		_dfbc, _ccff := _fcef._egdc.DecodeInt(_fcef._fff)
		if _ccff != nil {
			return 0, _ccff
		}
		return int64(_dfbc), nil
	}
	return 0, nil
}
func (_cbge *PageInformationSegment) encodeStripingInformation(_geccg _g.BinaryWriter) (_cgc int, _aggc error) {
	const _feef = "\u0065n\u0063\u006f\u0064\u0065S\u0074\u0072\u0069\u0070\u0069n\u0067I\u006ef\u006f\u0072\u006d\u0061\u0074\u0069\u006fn"
	if !_cbge.IsStripe {
		if _cgc, _aggc = _geccg.Write([]byte{0x00, 0x00}); _aggc != nil {
			return 0, _edb.Wrap(_aggc, _feef, "n\u006f\u0020\u0073\u0074\u0072\u0069\u0070\u0069\u006e\u0067")
		}
		return _cgc, nil
	}
	_babe := make([]byte, 2)
	_ag.BigEndian.PutUint16(_babe, _cbge.MaxStripeSize|1<<15)
	if _cgc, _aggc = _geccg.Write(_babe); _aggc != nil {
		return 0, _edb.Wrapf(_aggc, _feef, "\u0073\u0074\u0072i\u0070\u0069\u006e\u0067\u003a\u0020\u0025\u0064", _cbge.MaxStripeSize)
	}
	return _cgc, nil
}
func (_be *GenericRefinementRegion) GetRegionInfo() *RegionSegment { return _be.RegionInfo }
func (_eeacb *TextRegion) setParameters(_cecg *_df.Decoder, _cbae, _cbddb bool, _dbaag, _deab uint32, _fbcfg uint32, _cbadc int8, _bagf uint32, _gfgc int8, _dabf _gf.CombinationOperator, _dcef int8, _dfgag int16, _gfgcb, _efa, _eece, _ecdb, _bdbb, _cgbd, _bfgec, _aafg, _abgb, _aabd int8, _eafa, _acece []int8, _adbc []*_gf.Bitmap, _ecfe int8) {
	_eeacb._egdc = _cecg
	_eeacb.IsHuffmanEncoded = _cbae
	_eeacb.UseRefinement = _cbddb
	_eeacb.RegionInfo.BitmapWidth = _dbaag
	_eeacb.RegionInfo.BitmapHeight = _deab
	_eeacb.NumberOfSymbolInstances = _fbcfg
	_eeacb.SbStrips = _cbadc
	_eeacb.NumberOfSymbols = _bagf
	_eeacb.DefaultPixel = _gfgc
	_eeacb.CombinationOperator = _dabf
	_eeacb.IsTransposed = _dcef
	_eeacb.ReferenceCorner = _dfgag
	_eeacb.SbDsOffset = _gfgcb
	_eeacb.SbHuffFS = _efa
	_eeacb.SbHuffDS = _eece
	_eeacb.SbHuffDT = _ecdb
	_eeacb.SbHuffRDWidth = _bdbb
	_eeacb.SbHuffRDHeight = _cgbd
	_eeacb.SbHuffRSize = _abgb
	_eeacb.SbHuffRDX = _bfgec
	_eeacb.SbHuffRDY = _aafg
	_eeacb.SbrTemplate = _aabd
	_eeacb.SbrATX = _eafa
	_eeacb.SbrATY = _acece
	_eeacb.Symbols = _adbc
	_eeacb._efeb = _ecfe
}
func (_gafe *TableSegment) HtOOB() int32 { return _gafe._cbad }
func (_ccef *SymbolDictionary) decodeDifferenceWidth() (int64, error) {
	if _ccef.IsHuffmanEncoded {
		switch _ccef.SdHuffDecodeWidthSelection {
		case 0:
			_bccf, _eaga := _fa.GetStandardTable(2)
			if _eaga != nil {
				return 0, _eaga
			}
			return _bccf.Decode(_ccef._caba)
		case 1:
			_eedf, _dfcf := _fa.GetStandardTable(3)
			if _dfcf != nil {
				return 0, _dfcf
			}
			return _eedf.Decode(_ccef._caba)
		case 3:
			if _ccef._accg == nil {
				var _ffef int
				if _ccef.SdHuffDecodeHeightSelection == 3 {
					_ffef++
				}
				_cacc, _ddccc := _ccef.getUserTable(_ffef)
				if _ddccc != nil {
					return 0, _ddccc
				}
				_ccef._accg = _cacc
			}
			return _ccef._accg.Decode(_ccef._caba)
		}
	} else {
		_eace, _eca := _ccef._dgcb.DecodeInt(_ccef._bcbb)
		if _eca != nil {
			return 0, _eca
		}
		return int64(_eace), nil
	}
	return 0, nil
}
func (_bea *GenericRegion) decodeSLTP() (int, error) {
	switch _bea.GBTemplate {
	case 0:
		_bea._fac.SetIndex(0x9B25)
	case 1:
		_bea._fac.SetIndex(0x795)
	case 2:
		_bea._fac.SetIndex(0xE5)
	case 3:
		_bea._fac.SetIndex(0x195)
	}
	return _bea._ddga.DecodeBit(_bea._fac)
}
func (_dggf *SymbolDictionary) decodeNewSymbols(_cagc, _daca uint32, _daab *_gf.Bitmap, _ebgg, _cbeb int32) error {
	if _dggf._ebca == nil {
		_dggf._ebca = _agf(_dggf._caba, nil)
		if _dggf._dgcb == nil {
			var _gefec error
			_dggf._dgcb, _gefec = _df.New(_dggf._caba)
			if _gefec != nil {
				return _gefec
			}
		}
		if _dggf._afbd == nil {
			_dggf._afbd = _df.NewStats(65536, 1)
		}
	}
	_dggf._ebca.setParameters(_dggf._afbd, _dggf._dgcb, _dggf.SdrTemplate, _cagc, _daca, _daab, _ebgg, _cbeb, false, _dggf.SdrATX, _dggf.SdrATY)
	return _dggf.addSymbol(_dggf._ebca)
}
func (_ffc *template0) setIndex(_afgd *_df.DecoderStats) { _afgd.SetIndex(0x100) }
func (_fgdg *SymbolDictionary) String() string {
	_debea := &_d.Builder{}
	_debea.WriteString("\n\u005b\u0053\u0059\u004dBO\u004c-\u0044\u0049\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u000a")
	_debea.WriteString(_af.Sprintf("\u0009-\u0020S\u0064\u0072\u0054\u0065\u006dp\u006c\u0061t\u0065\u0020\u0025\u0076\u000a", _fgdg.SdrTemplate))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u0025\u0076\u000a", _fgdg.SdTemplate))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0069\u0073\u0043\u006f\u0064\u0069\u006eg\u0043\u006f\u006e\u0074\u0065\u0078\u0074R\u0065\u0074\u0061\u0069\u006e\u0065\u0064\u0020\u0025\u0076\u000a", _fgdg._dedee))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0069\u0073\u0043\u006f\u0064\u0069\u006e\u0067C\u006f\u006e\u0074\u0065\u0078\u0074\u0055\u0073\u0065\u0064 \u0025\u0076\u000a", _fgdg._bddf))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0048u\u0066\u0066\u0041\u0067\u0067\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065S\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0025\u0076\u000a", _fgdg.SdHuffAggInstanceSelection))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053d\u0048\u0075\u0066\u0066\u0042\u004d\u0053\u0069\u007a\u0065S\u0065l\u0065\u0063\u0074\u0069\u006f\u006e\u0020%\u0076\u000a", _fgdg.SdHuffBMSizeSelection))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0048u\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0057\u0069\u0064\u0074\u0068S\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0025\u0076\u000a", _fgdg.SdHuffDecodeWidthSelection))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020Sd\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064e\u0048e\u0069g\u0068t\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0020\u0025\u0076\u000a", _fgdg.SdHuffDecodeHeightSelection))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020U\u0073\u0065\u0052\u0065f\u0069\u006e\u0065\u006d\u0065\u006e\u0074A\u0067\u0067\u0072\u0065\u0067\u0061\u0074\u0069\u006f\u006e\u0020\u0025\u0076\u000a", _fgdg.UseRefinementAggregation))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020is\u0048\u0075\u0066\u0066\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0065\u0064\u0020\u0025\u0076\u000a", _fgdg.IsHuffmanEncoded))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020S\u0064\u0041\u0054\u0058\u0020\u0025\u0076\u000a", _fgdg.SdATX))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020S\u0064\u0041\u0054\u0059\u0020\u0025\u0076\u000a", _fgdg.SdATY))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0072\u0041\u0054\u0058\u0020\u0025\u0076\u000a", _fgdg.SdrATX))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0072\u0041\u0054\u0059\u0020\u0025\u0076\u000a", _fgdg.SdrATY))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u004ff\u0045\u0078\u0070\u006f\u0072\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0025\u0076\u000a", _fgdg.NumberOfExportedSymbols))
	_debea.WriteString(_af.Sprintf("\u0009-\u0020\u004e\u0075\u006db\u0065\u0072\u004f\u0066\u004ee\u0077S\u0079m\u0062\u006f\u006c\u0073\u0020\u0025\u0076\n", _fgdg.NumberOfNewSymbols))
	_debea.WriteString(_af.Sprintf("\u0009\u002d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u004ff\u0049\u006d\u0070\u006f\u0072\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0025\u0076\u000a", _fgdg._gfg))
	_debea.WriteString(_af.Sprintf("\u0009\u002d \u006e\u0075\u006d\u0062\u0065\u0072\u004f\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0064\u0053\u0079\u006d\u0062\u006f\u006c\u0073 %\u0076\u000a", _fgdg._acga))
	return _debea.String()
}
func (_cgee Type) String() string {
	switch _cgee {
	case TSymbolDictionary:
		return "\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0044\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079"
	case TIntermediateTextRegion:
		return "\u0049n\u0074\u0065\u0072\u006d\u0065\u0064\u0069\u0061\u0074\u0065\u0020T\u0065\u0078\u0074\u0020\u0052\u0065\u0067\u0069\u006f\u006e"
	case TImmediateTextRegion:
		return "I\u006d\u006d\u0065\u0064ia\u0074e\u0020\u0054\u0065\u0078\u0074 \u0052\u0065\u0067\u0069\u006f\u006e"
	case TImmediateLosslessTextRegion:
		return "\u0049\u006d\u006d\u0065\u0064\u0069\u0061\u0074\u0065\u0020L\u006f\u0073\u0073\u006c\u0065\u0073\u0073 \u0054\u0065\u0078\u0074\u0020\u0052\u0065\u0067\u0069\u006f\u006e"
	case TPatternDictionary:
		return "\u0050a\u0074t\u0065\u0072\u006e\u0020\u0044i\u0063\u0074i\u006f\u006e\u0061\u0072\u0079"
	case TIntermediateHalftoneRegion:
		return "\u0049\u006e\u0074\u0065r\u006d\u0065\u0064\u0069\u0061\u0074\u0065\u0020\u0048\u0061l\u0066t\u006f\u006e\u0065\u0020\u0052\u0065\u0067i\u006f\u006e"
	case TImmediateHalftoneRegion:
		return "\u0049m\u006d\u0065\u0064\u0069a\u0074\u0065\u0020\u0048\u0061l\u0066t\u006fn\u0065\u0020\u0052\u0065\u0067\u0069\u006fn"
	case TImmediateLosslessHalftoneRegion:
		return "\u0049\u006d\u006ded\u0069\u0061\u0074\u0065\u0020\u004c\u006f\u0073\u0073l\u0065s\u0073 \u0048a\u006c\u0066\u0074\u006f\u006e\u0065\u0020\u0052\u0065\u0067\u0069\u006f\u006e"
	case TIntermediateGenericRegion:
		return "I\u006e\u0074\u0065\u0072\u006d\u0065d\u0069\u0061\u0074\u0065\u0020\u0047\u0065\u006e\u0065r\u0069\u0063\u0020R\u0065g\u0069\u006f\u006e"
	case TImmediateGenericRegion:
		return "\u0049m\u006d\u0065\u0064\u0069\u0061\u0074\u0065\u0020\u0047\u0065\u006ee\u0072\u0069\u0063\u0020\u0052\u0065\u0067\u0069\u006f\u006e"
	case TImmediateLosslessGenericRegion:
		return "\u0049\u006d\u006d\u0065\u0064\u0069a\u0074\u0065\u0020\u004c\u006f\u0073\u0073\u006c\u0065\u0073\u0073\u0020\u0047e\u006e\u0065\u0072\u0069\u0063\u0020\u0052e\u0067\u0069\u006f\u006e"
	case TIntermediateGenericRefinementRegion:
		return "\u0049\u006e\u0074\u0065\u0072\u006d\u0065\u0064\u0069\u0061\u0074\u0065\u0020\u0047\u0065\u006e\u0065\u0072\u0069\u0063\u0020\u0052\u0065\u0066i\u006e\u0065\u006d\u0065\u006et\u0020\u0052e\u0067\u0069\u006f\u006e"
	case TImmediateGenericRefinementRegion:
		return "I\u006d\u006d\u0065\u0064\u0069\u0061t\u0065\u0020\u0047\u0065\u006e\u0065r\u0069\u0063\u0020\u0052\u0065\u0066\u0069n\u0065\u006d\u0065\u006e\u0074\u0020\u0052\u0065\u0067\u0069o\u006e"
	case TImmediateLosslessGenericRefinementRegion:
		return "\u0049m\u006d\u0065d\u0069\u0061\u0074\u0065 \u004c\u006f\u0073s\u006c\u0065\u0073\u0073\u0020\u0047\u0065\u006e\u0065ri\u0063\u0020\u0052e\u0066\u0069n\u0065\u006d\u0065\u006e\u0074\u0020R\u0065\u0067i\u006f\u006e"
	case TPageInformation:
		return "\u0050\u0061g\u0065\u0020\u0049n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e"
	case TEndOfPage:
		return "E\u006e\u0064\u0020\u004f\u0066\u0020\u0050\u0061\u0067\u0065"
	case TEndOfStrip:
		return "\u0045\u006e\u0064 \u004f\u0066\u0020\u0053\u0074\u0072\u0069\u0070"
	case TEndOfFile:
		return "E\u006e\u0064\u0020\u004f\u0066\u0020\u0046\u0069\u006c\u0065"
	case TProfiles:
		return "\u0050\u0072\u006f\u0066\u0069\u006c\u0065\u0073"
	case TTables:
		return "\u0054\u0061\u0062\u006c\u0065\u0073"
	case TExtension:
		return "\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"
	case TBitmap:
		return "\u0042\u0069\u0074\u006d\u0061\u0070"
	}
	return "I\u006ev\u0061\u006c\u0069\u0064\u0020\u0053\u0065\u0067m\u0065\u006e\u0074\u0020Ki\u006e\u0064"
}
func (_dbgd *GenericRegion) overrideAtTemplate1(_bbg, _gag, _aac, _ddcc, _cfcb int) int {
	_bbg &= 0x1FF7
	if _dbgd.GBAtY[0] == 0 && _dbgd.GBAtX[0] >= -int8(_cfcb) {
		_bbg |= (_ddcc >> uint(7-(int8(_cfcb)+_dbgd.GBAtX[0])) & 0x1) << 3
	} else {
		_bbg |= int(_dbgd.getPixel(_gag+int(_dbgd.GBAtX[0]), _aac+int(_dbgd.GBAtY[0]))) << 3
	}
	return _bbg
}
func (_bffb *SymbolDictionary) readNumberOfExportedSymbols() error {
	_fagf, _cbag := _bffb._caba.ReadBits(32)
	if _cbag != nil {
		return _cbag
	}
	_bffb.NumberOfExportedSymbols = uint32(_fagf & _e.MaxUint32)
	return nil
}
func (_bddb *GenericRegion) decodeTemplate2(_efc, _aeg, _gef int, _ddfg, _fde int) (_aadg error) {
	const _efge = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0032"
	var (
		_bfdf, _gbbc int
		_fbdg, _fab  int
		_caaec       byte
		_agfg, _bcd  int
	)
	if _efc >= 1 {
		_caaec, _aadg = _bddb.Bitmap.GetByte(_fde)
		if _aadg != nil {
			return _edb.Wrap(_aadg, _efge, "\u006ci\u006ee\u004e\u0075\u006d\u0062\u0065\u0072\u0020\u003e\u003d\u0020\u0031")
		}
		_fbdg = int(_caaec)
	}
	if _efc >= 2 {
		_caaec, _aadg = _bddb.Bitmap.GetByte(_fde - _bddb.Bitmap.RowStride)
		if _aadg != nil {
			return _edb.Wrap(_aadg, _efge, "\u006ci\u006ee\u004e\u0075\u006d\u0062\u0065\u0072\u0020\u003e\u003d\u0020\u0032")
		}
		_fab = int(_caaec) << 4
	}
	_bfdf = (_fbdg >> 3 & 0x7c) | (_fab >> 3 & 0x380)
	for _efgg := 0; _efgg < _gef; _efgg = _agfg {
		var (
			_fafe byte
			_fgga int
		)
		_agfg = _efgg + 8
		if _aegb := _aeg - _efgg; _aegb > 8 {
			_fgga = 8
		} else {
			_fgga = _aegb
		}
		if _efc > 0 {
			_fbdg <<= 8
			if _agfg < _aeg {
				_caaec, _aadg = _bddb.Bitmap.GetByte(_fde + 1)
				if _aadg != nil {
					return _edb.Wrap(_aadg, _efge, "\u006c\u0069\u006e\u0065\u004e\u0075\u006d\u0062\u0065r\u0020\u003e\u0020\u0030")
				}
				_fbdg |= int(_caaec)
			}
		}
		if _efc > 1 {
			_fab <<= 8
			if _agfg < _aeg {
				_caaec, _aadg = _bddb.Bitmap.GetByte(_fde - _bddb.Bitmap.RowStride + 1)
				if _aadg != nil {
					return _edb.Wrap(_aadg, _efge, "\u006c\u0069\u006e\u0065\u004e\u0075\u006d\u0062\u0065r\u0020\u003e\u0020\u0031")
				}
				_fab |= int(_caaec) << 4
			}
		}
		for _afad := 0; _afad < _fgga; _afad++ {
			_bgag := uint(10 - _afad)
			if _bddb._cab {
				_gbbc = _bddb.overrideAtTemplate2(_bfdf, _efgg+_afad, _efc, int(_fafe), _afad)
				_bddb._fac.SetIndex(int32(_gbbc))
			} else {
				_bddb._fac.SetIndex(int32(_bfdf))
			}
			_bcd, _aadg = _bddb._ddga.DecodeBit(_bddb._fac)
			if _aadg != nil {
				return _edb.Wrap(_aadg, _efge, "")
			}
			_fafe |= byte(_bcd << uint(7-_afad))
			_bfdf = ((_bfdf & 0x1bd) << 1) | _bcd | ((_fbdg >> _bgag) & 0x4) | ((_fab >> _bgag) & 0x80)
		}
		if _aaee := _bddb.Bitmap.SetByte(_ddfg, _fafe); _aaee != nil {
			return _edb.Wrap(_aaee, _efge, "")
		}
		_ddfg++
		_fde++
	}
	return nil
}

type template1 struct{}

func (_bbabb *HalftoneRegion) parseHeader() error {
	if _ggag := _bbabb.RegionSegment.parseHeader(); _ggag != nil {
		return _ggag
	}
	_gecg, _eaeg := _bbabb._fagd.ReadBit()
	if _eaeg != nil {
		return _eaeg
	}
	_bbabb.HDefaultPixel = int8(_gecg)
	_degd, _eaeg := _bbabb._fagd.ReadBits(3)
	if _eaeg != nil {
		return _eaeg
	}
	_bbabb.CombinationOperator = _gf.CombinationOperator(_degd & 0xf)
	_gecg, _eaeg = _bbabb._fagd.ReadBit()
	if _eaeg != nil {
		return _eaeg
	}
	if _gecg == 1 {
		_bbabb.HSkipEnabled = true
	}
	_degd, _eaeg = _bbabb._fagd.ReadBits(2)
	if _eaeg != nil {
		return _eaeg
	}
	_bbabb.HTemplate = byte(_degd & 0xf)
	_gecg, _eaeg = _bbabb._fagd.ReadBit()
	if _eaeg != nil {
		return _eaeg
	}
	if _gecg == 1 {
		_bbabb.IsMMREncoded = true
	}
	_degd, _eaeg = _bbabb._fagd.ReadBits(32)
	if _eaeg != nil {
		return _eaeg
	}
	_bbabb.HGridWidth = uint32(_degd & _e.MaxUint32)
	_degd, _eaeg = _bbabb._fagd.ReadBits(32)
	if _eaeg != nil {
		return _eaeg
	}
	_bbabb.HGridHeight = uint32(_degd & _e.MaxUint32)
	_degd, _eaeg = _bbabb._fagd.ReadBits(32)
	if _eaeg != nil {
		return _eaeg
	}
	_bbabb.HGridX = int32(_degd & _e.MaxInt32)
	_degd, _eaeg = _bbabb._fagd.ReadBits(32)
	if _eaeg != nil {
		return _eaeg
	}
	_bbabb.HGridY = int32(_degd & _e.MaxInt32)
	_degd, _eaeg = _bbabb._fagd.ReadBits(16)
	if _eaeg != nil {
		return _eaeg
	}
	_bbabb.HRegionX = uint16(_degd & _e.MaxUint16)
	_degd, _eaeg = _bbabb._fagd.ReadBits(16)
	if _eaeg != nil {
		return _eaeg
	}
	_bbabb.HRegionY = uint16(_degd & _e.MaxUint16)
	if _eaeg = _bbabb.computeSegmentDataStructure(); _eaeg != nil {
		return _eaeg
	}
	return _bbabb.checkInput()
}
func (_bcgd *Header) CleanSegmentData() {
	if _bcgd.SegmentData != nil {
		_bcgd.SegmentData = nil
	}
}
func (_dabbc *SymbolDictionary) decodeDirectlyThroughGenericRegion(_eec, _bfdad uint32) error {
	if _dabbc._ddcgf == nil {
		_dabbc._ddcgf = NewGenericRegion(_dabbc._caba)
	}
	_dabbc._ddcgf.setParametersWithAt(false, byte(_dabbc.SdTemplate), false, false, _dabbc.SdATX, _dabbc.SdATY, _eec, _bfdad, _dabbc._afbd, _dabbc._dgcb)
	return _dabbc.addSymbol(_dabbc._ddcgf)
}
func (_ddd *template1) form(_cced, _acd, _bgfg, _dcc, _agc int16) int16 {
	return ((_cced & 0x02) << 8) | (_acd << 6) | ((_bgfg & 0x03) << 4) | (_dcc << 1) | _agc
}
func (_caad *PageInformationSegment) Init(h *Header, r *_g.Reader) (_fdcd error) {
	_caad._faef = r
	if _fdcd = _caad.parseHeader(); _fdcd != nil {
		return _edb.Wrap(_fdcd, "P\u0061\u0067\u0065\u0049\u006e\u0066o\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065g\u006d\u0065\u006et\u002eI\u006e\u0069\u0074", "")
	}
	return nil
}
func (_egeb *PatternDictionary) readIsMMREncoded() error {
	_gbeb, _afga := _egeb._dbec.ReadBit()
	if _afga != nil {
		return _afga
	}
	if _gbeb != 0 {
		_egeb.IsMMREncoded = true
	}
	return nil
}
func (_defb *SymbolDictionary) encodeATFlags(_cgddc _g.BinaryWriter) (_dfad int, _egc error) {
	const _gfeeb = "\u0065\u006e\u0063\u006f\u0064\u0065\u0041\u0054\u0046\u006c\u0061\u0067\u0073"
	if _defb.IsHuffmanEncoded || _defb.SdTemplate != 0 {
		return 0, nil
	}
	_ccga := 4
	if _defb.SdTemplate != 0 {
		_ccga = 1
	}
	for _ffaa := 0; _ffaa < _ccga; _ffaa++ {
		if _egc = _cgddc.WriteByte(byte(_defb.SdATX[_ffaa])); _egc != nil {
			return _dfad, _edb.Wrapf(_egc, _gfeeb, "\u0053d\u0041\u0054\u0058\u005b\u0025\u0064]", _ffaa)
		}
		_dfad++
		if _egc = _cgddc.WriteByte(byte(_defb.SdATY[_ffaa])); _egc != nil {
			return _dfad, _edb.Wrapf(_egc, _gfeeb, "\u0053d\u0041\u0054\u0059\u005b\u0025\u0064]", _ffaa)
		}
		_dfad++
	}
	return _dfad, nil
}

type Header struct {
	SegmentNumber            uint32
	Type                     Type
	RetainFlag               bool
	PageAssociation          int
	PageAssociationFieldSize bool
	RTSegments               []*Header
	HeaderLength             int64
	SegmentDataLength        uint64
	SegmentDataStartOffset   uint64
	Reader                   *_g.Reader
	SegmentData              Segmenter
	RTSNumbers               []int
	RetainBits               []uint8
}

func (_gaef *SymbolDictionary) getSbSymCodeLen() int8 {
	_fabe := int8(_e.Ceil(_e.Log(float64(_gaef._gfg+_gaef.NumberOfNewSymbols)) / _e.Log(2)))
	if _gaef.IsHuffmanEncoded && _fabe < 1 {
		return 1
	}
	return _fabe
}
func (_bacf *SymbolDictionary) encodeNumSyms(_aegg _g.BinaryWriter) (_fgcbc int, _gefe error) {
	const _fdffe = "\u0065\u006e\u0063\u006f\u0064\u0065\u004e\u0075\u006d\u0053\u0079\u006d\u0073"
	_gfdbg := make([]byte, 4)
	_ag.BigEndian.PutUint32(_gfdbg, _bacf.NumberOfExportedSymbols)
	if _fgcbc, _gefe = _aegg.Write(_gfdbg); _gefe != nil {
		return _fgcbc, _edb.Wrap(_gefe, _fdffe, "\u0065\u0078p\u006f\u0072\u0074e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	_ag.BigEndian.PutUint32(_gfdbg, _bacf.NumberOfNewSymbols)
	_ggdg, _gefe := _aegg.Write(_gfdbg)
	if _gefe != nil {
		return _fgcbc, _edb.Wrap(_gefe, _fdffe, "n\u0065\u0077\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	return _fgcbc + _ggdg, nil
}
func (_caf *HalftoneRegion) GetRegionInfo() *RegionSegment { return _caf.RegionSegment }
func (_daba *PageInformationSegment) readCombinationOperator() error {
	_fdg, _acbf := _daba._faef.ReadBits(2)
	if _acbf != nil {
		return _acbf
	}
	_daba._dbcd = _gf.CombinationOperator(int(_fdg))
	return nil
}
func (_eff *template1) setIndex(_bgb *_df.DecoderStats) { _bgb.SetIndex(0x080) }
func (_dgda *TextRegion) Encode(w _g.BinaryWriter) (_fffb int, _bcgda error) {
	const _affb = "\u0054\u0065\u0078\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u002e\u0045n\u0063\u006f\u0064\u0065"
	if _fffb, _bcgda = _dgda.RegionInfo.Encode(w); _bcgda != nil {
		return _fffb, _edb.Wrap(_bcgda, _affb, "")
	}
	var _eccd int
	if _eccd, _bcgda = _dgda.encodeFlags(w); _bcgda != nil {
		return _fffb, _edb.Wrap(_bcgda, _affb, "")
	}
	_fffb += _eccd
	if _eccd, _bcgda = _dgda.encodeSymbols(w); _bcgda != nil {
		return _fffb, _edb.Wrap(_bcgda, _affb, "")
	}
	_fffb += _eccd
	return _fffb, nil
}
func (_gae *GenericRegion) writeGBAtPixels(_ccaf _g.BinaryWriter) (_cdf int, _acbe error) {
	const _cagg = "\u0077r\u0069t\u0065\u0047\u0042\u0041\u0074\u0050\u0069\u0078\u0065\u006c\u0073"
	if _gae.UseMMR {
		return 0, nil
	}
	_caac := 1
	if _gae.GBTemplate == 0 {
		_caac = 4
	} else if _gae.UseExtTemplates {
		_caac = 12
	}
	if len(_gae.GBAtX) != _caac {
		return 0, _edb.Errorf(_cagg, "\u0067\u0062\u0020\u0061\u0074\u0020\u0070\u0061\u0069\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020d\u006f\u0065\u0073\u006e\u0027\u0074\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u006f\u0020\u0047\u0042\u0041t\u0058\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u006c\u0065\u006e")
	}
	if len(_gae.GBAtY) != _caac {
		return 0, _edb.Errorf(_cagg, "\u0067\u0062\u0020\u0061\u0074\u0020\u0070\u0061\u0069\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020d\u006f\u0065\u0073\u006e\u0027\u0074\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u006f\u0020\u0047\u0042\u0041t\u0059\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u006c\u0065\u006e")
	}
	for _eed := 0; _eed < _caac; _eed++ {
		if _acbe = _ccaf.WriteByte(byte(_gae.GBAtX[_eed])); _acbe != nil {
			return _cdf, _edb.Wrap(_acbe, _cagg, "w\u0072\u0069\u0074\u0065\u0020\u0047\u0042\u0041\u0074\u0058")
		}
		_cdf++
		if _acbe = _ccaf.WriteByte(byte(_gae.GBAtY[_eed])); _acbe != nil {
			return _cdf, _edb.Wrap(_acbe, _cagg, "w\u0072\u0069\u0074\u0065\u0020\u0047\u0042\u0041\u0074\u0059")
		}
		_cdf++
	}
	return _cdf, nil
}
func (_fec *GenericRegion) Size() int { return _fec.RegionSegment.Size() + 1 + 2*len(_fec.GBAtX) }
func (_aacb *HalftoneRegion) checkInput() error {
	if _aacb.IsMMREncoded {
		if _aacb.HTemplate != 0 {
			_aae.Log.Debug("\u0048\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0030", _aacb.HTemplate)
		}
		if _aacb.HSkipEnabled {
			_aae.Log.Debug("\u0048\u0053\u006b\u0069\u0070\u0045\u006e\u0061\u0062\u006c\u0065\u0064\u0020\u0030\u0020\u0025\u0076\u0020(\u0073\u0068\u006f\u0075\u006c\u0064\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u0029", _aacb.HSkipEnabled)
		}
	}
	return nil
}
func (_adgd *TextRegion) decodeSymInRefSize() (int64, error) {
	const _dacac = "\u0064e\u0063o\u0064\u0065\u0053\u0079\u006dI\u006e\u0052e\u0066\u0053\u0069\u007a\u0065"
	if _adgd.SbHuffRSize == 0 {
		_aedbd, _bfga := _fa.GetStandardTable(1)
		if _bfga != nil {
			return 0, _edb.Wrap(_bfga, _dacac, "")
		}
		return _aedbd.Decode(_adgd._ecc)
	}
	if _adgd._efgefe == nil {
		var (
			_ccdf int
			_efba error
		)
		if _adgd.SbHuffFS == 3 {
			_ccdf++
		}
		if _adgd.SbHuffDS == 3 {
			_ccdf++
		}
		if _adgd.SbHuffDT == 3 {
			_ccdf++
		}
		if _adgd.SbHuffRDWidth == 3 {
			_ccdf++
		}
		if _adgd.SbHuffRDHeight == 3 {
			_ccdf++
		}
		if _adgd.SbHuffRDX == 3 {
			_ccdf++
		}
		if _adgd.SbHuffRDY == 3 {
			_ccdf++
		}
		_adgd._efgefe, _efba = _adgd.getUserTable(_ccdf)
		if _efba != nil {
			return 0, _edb.Wrap(_efba, _dacac, "")
		}
	}
	_dddbb, _gbce := _adgd._efgefe.Decode(_adgd._ecc)
	if _gbce != nil {
		return 0, _edb.Wrap(_gbce, _dacac, "")
	}
	return _dddbb, nil
}
func (_bae *SymbolDictionary) huffDecodeBmSize() (int64, error) {
	if _bae._ccbc == nil {
		var (
			_faca int
			_abag error
		)
		if _bae.SdHuffDecodeHeightSelection == 3 {
			_faca++
		}
		if _bae.SdHuffDecodeWidthSelection == 3 {
			_faca++
		}
		_bae._ccbc, _abag = _bae.getUserTable(_faca)
		if _abag != nil {
			return 0, _abag
		}
	}
	return _bae._ccbc.Decode(_bae._caba)
}
func (_ggf *GenericRegion) updateOverrideFlags() error {
	const _ege = "\u0075\u0070\u0064\u0061te\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _ggf.GBAtX == nil || _ggf.GBAtY == nil {
		return nil
	}
	if len(_ggf.GBAtX) != len(_ggf.GBAtY) {
		return _edb.Errorf(_ege, "i\u006eco\u0073i\u0073t\u0065\u006e\u0074\u0020\u0041T\u0020\u0070\u0069x\u0065\u006c\u002e\u0020\u0041m\u006f\u0075\u006et\u0020\u006f\u0066\u0020\u0027\u0078\u0027\u0020\u0070\u0069\u0078e\u006c\u0073\u003a %d\u002c\u0020\u0041\u006d\u006f\u0075n\u0074\u0020\u006f\u0066\u0020\u0027\u0079\u0027\u0020\u0070\u0069\u0078e\u006cs\u003a\u0020\u0025\u0064", len(_ggf.GBAtX), len(_ggf.GBAtY))
	}
	_ggf.GBAtOverride = make([]bool, len(_ggf.GBAtX))
	switch _ggf.GBTemplate {
	case 0:
		if !_ggf.UseExtTemplates {
			if _ggf.GBAtX[0] != 3 || _ggf.GBAtY[0] != -1 {
				_ggf.setOverrideFlag(0)
			}
			if _ggf.GBAtX[1] != -3 || _ggf.GBAtY[1] != -1 {
				_ggf.setOverrideFlag(1)
			}
			if _ggf.GBAtX[2] != 2 || _ggf.GBAtY[2] != -2 {
				_ggf.setOverrideFlag(2)
			}
			if _ggf.GBAtX[3] != -2 || _ggf.GBAtY[3] != -2 {
				_ggf.setOverrideFlag(3)
			}
		} else {
			if _ggf.GBAtX[0] != -2 || _ggf.GBAtY[0] != 0 {
				_ggf.setOverrideFlag(0)
			}
			if _ggf.GBAtX[1] != 0 || _ggf.GBAtY[1] != -2 {
				_ggf.setOverrideFlag(1)
			}
			if _ggf.GBAtX[2] != -2 || _ggf.GBAtY[2] != -1 {
				_ggf.setOverrideFlag(2)
			}
			if _ggf.GBAtX[3] != -1 || _ggf.GBAtY[3] != -2 {
				_ggf.setOverrideFlag(3)
			}
			if _ggf.GBAtX[4] != 1 || _ggf.GBAtY[4] != -2 {
				_ggf.setOverrideFlag(4)
			}
			if _ggf.GBAtX[5] != 2 || _ggf.GBAtY[5] != -1 {
				_ggf.setOverrideFlag(5)
			}
			if _ggf.GBAtX[6] != -3 || _ggf.GBAtY[6] != 0 {
				_ggf.setOverrideFlag(6)
			}
			if _ggf.GBAtX[7] != -4 || _ggf.GBAtY[7] != 0 {
				_ggf.setOverrideFlag(7)
			}
			if _ggf.GBAtX[8] != 2 || _ggf.GBAtY[8] != -2 {
				_ggf.setOverrideFlag(8)
			}
			if _ggf.GBAtX[9] != 3 || _ggf.GBAtY[9] != -1 {
				_ggf.setOverrideFlag(9)
			}
			if _ggf.GBAtX[10] != -2 || _ggf.GBAtY[10] != -2 {
				_ggf.setOverrideFlag(10)
			}
			if _ggf.GBAtX[11] != -3 || _ggf.GBAtY[11] != -1 {
				_ggf.setOverrideFlag(11)
			}
		}
	case 1:
		if _ggf.GBAtX[0] != 3 || _ggf.GBAtY[0] != -1 {
			_ggf.setOverrideFlag(0)
		}
	case 2:
		if _ggf.GBAtX[0] != 2 || _ggf.GBAtY[0] != -1 {
			_ggf.setOverrideFlag(0)
		}
	case 3:
		if _ggf.GBAtX[0] != 2 || _ggf.GBAtY[0] != -1 {
			_ggf.setOverrideFlag(0)
		}
	}
	return nil
}
func (_cgce *PageInformationSegment) readRequiresAuxiliaryBuffer() error {
	_gcc, _bfbd := _cgce._faef.ReadBit()
	if _bfbd != nil {
		return _bfbd
	}
	if _gcc == 1 {
		_cgce._cbdc = true
	}
	return nil
}

var _ _fa.BasicTabler = &TableSegment{}

func (_fdbg *TableSegment) parseHeader() error {
	var (
		_bcbg int
		_fgbb uint64
		_gbac error
	)
	_bcbg, _gbac = _fdbg._dgdg.ReadBit()
	if _gbac != nil {
		return _gbac
	}
	if _bcbg == 1 {
		return _af.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0061\u0062\u006c\u0065 \u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0064e\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u002e\u0020\u0042\u002e\u0032\u002e1\u0020\u0043\u006f\u0064\u0065\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0066\u006c\u0061\u0067\u0073\u003a\u0020\u0042\u0069\u0074\u0020\u0037\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u007a\u0065\u0072\u006f\u002e\u0020\u0057a\u0073\u003a \u0025\u0064", _bcbg)
	}
	if _fgbb, _gbac = _fdbg._dgdg.ReadBits(3); _gbac != nil {
		return _gbac
	}
	_fdbg._efee = (int32(_fgbb) + 1) & 0xf
	if _fgbb, _gbac = _fdbg._dgdg.ReadBits(3); _gbac != nil {
		return _gbac
	}
	_fdbg._facaf = (int32(_fgbb) + 1) & 0xf
	if _fgbb, _gbac = _fdbg._dgdg.ReadBits(32); _gbac != nil {
		return _gbac
	}
	_fdbg._effba = int32(_fgbb & _e.MaxInt32)
	if _fgbb, _gbac = _fdbg._dgdg.ReadBits(32); _gbac != nil {
		return _gbac
	}
	_fdbg._dgaf = int32(_fgbb & _e.MaxInt32)
	return nil
}
func (_effg *SymbolDictionary) getUserTable(_aeff int) (_fa.Tabler, error) {
	var _effc int
	for _, _bbf := range _effg.Header.RTSegments {
		if _bbf.Type == 53 {
			if _effc == _aeff {
				_dcfb, _eddad := _bbf.GetSegmentData()
				if _eddad != nil {
					return nil, _eddad
				}
				_eecb := _dcfb.(_fa.BasicTabler)
				return _fa.NewEncodedTable(_eecb)
			}
			_effc++
		}
	}
	return nil, nil
}
func (_bcab *TextRegion) getSymbols() error {
	if _bcab.Header.RTSegments != nil {
		return _bcab.initSymbols()
	}
	return nil
}
func (_edga *Header) GetSegmentData() (Segmenter, error) {
	var _edfe Segmenter
	if _edga.SegmentData != nil {
		_edfe = _edga.SegmentData
	}
	if _edfe == nil {
		_acff, _dfcc := _ggdf[_edga.Type]
		if !_dfcc {
			return nil, _af.Errorf("\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0073\u002f\u0020\u0025\u0064\u0020\u0063\u0072e\u0061t\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e\u0020", _edga.Type, _edga.Type)
		}
		_edfe = _acff()
		_aae.Log.Trace("\u005b\u0053E\u0047\u004d\u0045\u004e\u0054-\u0048\u0045\u0041\u0044\u0045R\u005d\u005b\u0023\u0025\u0064\u005d\u0020\u0047\u0065\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u0074\u0020\u004f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0030\u0034\u0058", _edga.SegmentNumber, _edga.SegmentDataStartOffset)
		_gbd, _cgec := _edga.subInputReader()
		if _cgec != nil {
			return nil, _cgec
		}
		if _cbde := _edfe.Init(_edga, _gbd); _cbde != nil {
			_aae.Log.Debug("\u0049\u006e\u0069\u0074 \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076 \u0066o\u0072\u0020\u0074\u0079\u0070\u0065\u003a \u0025\u0054", _cbde, _edfe)
			return nil, _cbde
		}
		_edga.SegmentData = _edfe
	}
	return _edfe, nil
}
func (_acgg *SymbolDictionary) encodeFlags(_cadb _g.BinaryWriter) (_bfacd int, _aadgc error) {
	const _faff = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _aadgc = _cadb.SkipBits(3); _aadgc != nil {
		return 0, _edb.Wrap(_aadgc, _faff, "\u0065\u006d\u0070\u0074\u0079\u0020\u0062\u0069\u0074\u0073")
	}
	var _afd int
	if _acgg.SdrTemplate > 0 {
		_afd = 1
	}
	if _aadgc = _cadb.WriteBit(_afd); _aadgc != nil {
		return _bfacd, _edb.Wrap(_aadgc, _faff, "s\u0064\u0072\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_afd = 0
	if _acgg.SdTemplate > 1 {
		_afd = 1
	}
	if _aadgc = _cadb.WriteBit(_afd); _aadgc != nil {
		return _bfacd, _edb.Wrap(_aadgc, _faff, "\u0073\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_afd = 0
	if _acgg.SdTemplate == 1 || _acgg.SdTemplate == 3 {
		_afd = 1
	}
	if _aadgc = _cadb.WriteBit(_afd); _aadgc != nil {
		return _bfacd, _edb.Wrap(_aadgc, _faff, "\u0073\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_afd = 0
	if _acgg._dedee {
		_afd = 1
	}
	if _aadgc = _cadb.WriteBit(_afd); _aadgc != nil {
		return _bfacd, _edb.Wrap(_aadgc, _faff, "\u0063\u006f\u0064in\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0072\u0065\u0074\u0061\u0069\u006e\u0065\u0064")
	}
	_afd = 0
	if _acgg._bddf {
		_afd = 1
	}
	if _aadgc = _cadb.WriteBit(_afd); _aadgc != nil {
		return _bfacd, _edb.Wrap(_aadgc, _faff, "\u0063\u006f\u0064\u0069ng\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0075\u0073\u0065\u0064")
	}
	_afd = 0
	if _acgg.SdHuffAggInstanceSelection {
		_afd = 1
	}
	if _aadgc = _cadb.WriteBit(_afd); _aadgc != nil {
		return _bfacd, _edb.Wrap(_aadgc, _faff, "\u0073\u0064\u0048\u0075\u0066\u0066\u0041\u0067\u0067\u0049\u006e\u0073\u0074")
	}
	_afd = int(_acgg.SdHuffBMSizeSelection)
	if _aadgc = _cadb.WriteBit(_afd); _aadgc != nil {
		return _bfacd, _edb.Wrap(_aadgc, _faff, "\u0073\u0064\u0048u\u0066\u0066\u0042\u006d\u0053\u0069\u007a\u0065")
	}
	_afd = 0
	if _acgg.SdHuffDecodeWidthSelection > 1 {
		_afd = 1
	}
	if _aadgc = _cadb.WriteBit(_afd); _aadgc != nil {
		return _bfacd, _edb.Wrap(_aadgc, _faff, "s\u0064\u0048\u0075\u0066\u0066\u0057\u0069\u0064\u0074\u0068")
	}
	_afd = 0
	switch _acgg.SdHuffDecodeWidthSelection {
	case 1, 3:
		_afd = 1
	}
	if _aadgc = _cadb.WriteBit(_afd); _aadgc != nil {
		return _bfacd, _edb.Wrap(_aadgc, _faff, "s\u0064\u0048\u0075\u0066\u0066\u0057\u0069\u0064\u0074\u0068")
	}
	_afd = 0
	if _acgg.SdHuffDecodeHeightSelection > 1 {
		_afd = 1
	}
	if _aadgc = _cadb.WriteBit(_afd); _aadgc != nil {
		return _bfacd, _edb.Wrap(_aadgc, _faff, "\u0073\u0064\u0048u\u0066\u0066\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_afd = 0
	switch _acgg.SdHuffDecodeHeightSelection {
	case 1, 3:
		_afd = 1
	}
	if _aadgc = _cadb.WriteBit(_afd); _aadgc != nil {
		return _bfacd, _edb.Wrap(_aadgc, _faff, "\u0073\u0064\u0048u\u0066\u0066\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_afd = 0
	if _acgg.UseRefinementAggregation {
		_afd = 1
	}
	if _aadgc = _cadb.WriteBit(_afd); _aadgc != nil {
		return _bfacd, _edb.Wrap(_aadgc, _faff, "\u0073\u0064\u0052\u0065\u0066\u0041\u0067\u0067")
	}
	_afd = 0
	if _acgg.IsHuffmanEncoded {
		_afd = 1
	}
	if _aadgc = _cadb.WriteBit(_afd); _aadgc != nil {
		return _bfacd, _edb.Wrap(_aadgc, _faff, "\u0073\u0064\u0048\u0075\u0066\u0066")
	}
	return 2, nil
}
func (_fcfa *TextRegion) readHuffmanFlags() error {
	var (
		_eggd  int
		_ffgd  uint64
		_aggfc error
	)
	_, _aggfc = _fcfa._ecc.ReadBit()
	if _aggfc != nil {
		return _aggfc
	}
	_eggd, _aggfc = _fcfa._ecc.ReadBit()
	if _aggfc != nil {
		return _aggfc
	}
	_fcfa.SbHuffRSize = int8(_eggd)
	_ffgd, _aggfc = _fcfa._ecc.ReadBits(2)
	if _aggfc != nil {
		return _aggfc
	}
	_fcfa.SbHuffRDY = int8(_ffgd) & 0xf
	_ffgd, _aggfc = _fcfa._ecc.ReadBits(2)
	if _aggfc != nil {
		return _aggfc
	}
	_fcfa.SbHuffRDX = int8(_ffgd) & 0xf
	_ffgd, _aggfc = _fcfa._ecc.ReadBits(2)
	if _aggfc != nil {
		return _aggfc
	}
	_fcfa.SbHuffRDHeight = int8(_ffgd) & 0xf
	_ffgd, _aggfc = _fcfa._ecc.ReadBits(2)
	if _aggfc != nil {
		return _aggfc
	}
	_fcfa.SbHuffRDWidth = int8(_ffgd) & 0xf
	_ffgd, _aggfc = _fcfa._ecc.ReadBits(2)
	if _aggfc != nil {
		return _aggfc
	}
	_fcfa.SbHuffDT = int8(_ffgd) & 0xf
	_ffgd, _aggfc = _fcfa._ecc.ReadBits(2)
	if _aggfc != nil {
		return _aggfc
	}
	_fcfa.SbHuffDS = int8(_ffgd) & 0xf
	_ffgd, _aggfc = _fcfa._ecc.ReadBits(2)
	if _aggfc != nil {
		return _aggfc
	}
	_fcfa.SbHuffFS = int8(_ffgd) & 0xf
	return nil
}
func (_gaae *Header) String() string {
	_aaab := &_d.Builder{}
	_aaab.WriteString("\u000a[\u0053E\u0047\u004d\u0045\u004e\u0054-\u0048\u0045A\u0044\u0045\u0052\u005d\u000a")
	_aaab.WriteString(_af.Sprintf("\t\u002d\u0020\u0053\u0065gm\u0065n\u0074\u004e\u0075\u006d\u0062e\u0072\u003a\u0020\u0025\u0076\u000a", _gaae.SegmentNumber))
	_aaab.WriteString(_af.Sprintf("\u0009\u002d\u0020T\u0079\u0070\u0065\u003a\u0020\u0025\u0076\u000a", _gaae.Type))
	_aaab.WriteString(_af.Sprintf("\u0009-\u0020R\u0065\u0074\u0061\u0069\u006eF\u006c\u0061g\u003a\u0020\u0025\u0076\u000a", _gaae.RetainFlag))
	_aaab.WriteString(_af.Sprintf("\u0009\u002d\u0020Pa\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _gaae.PageAssociation))
	_aaab.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0050\u0061\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061\u0074i\u006fn\u0046\u0069\u0065\u006c\u0064\u0053\u0069\u007a\u0065\u003a\u0020\u0025\u0076\u000a", _gaae.PageAssociationFieldSize))
	_aaab.WriteString("\u0009-\u0020R\u0054\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u0053\u003a\u000a")
	for _, _acag := range _gaae.RTSNumbers {
		_aaab.WriteString(_af.Sprintf("\u0009\t\u002d\u0020\u0025\u0064\u000a", _acag))
	}
	_aaab.WriteString(_af.Sprintf("\t\u002d \u0048\u0065\u0061\u0064\u0065\u0072\u004c\u0065n\u0067\u0074\u0068\u003a %\u0076\u000a", _gaae.HeaderLength))
	_aaab.WriteString(_af.Sprintf("\u0009-\u0020\u0053\u0065\u0067m\u0065\u006e\u0074\u0044\u0061t\u0061L\u0065n\u0067\u0074\u0068\u003a\u0020\u0025\u0076\n", _gaae.SegmentDataLength))
	_aaab.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074D\u0061\u0074\u0061\u0053\u0074\u0061\u0072t\u004f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0076\u000a", _gaae.SegmentDataStartOffset))
	return _aaab.String()
}

type Pager interface {
	GetSegment(int) (*Header, error)
	GetBitmap() (*_gf.Bitmap, error)
}

func (_eef *GenericRegion) parseHeader() (_fca error) {
	_aae.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052I\u0043\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0050\u0061\u0072s\u0069\u006e\u0067\u0048\u0065\u0061\u0064e\u0072\u002e\u002e\u002e")
	defer func() {
		if _fca != nil {
			_aae.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0047\u0049\u004f\u004e]\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065\u0061\u0064\u0065r\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u0020\u0077\u0069th\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020\u0025\u0076", _fca)
		} else {
			_aae.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049C\u002d\u0052\u0045G\u0049\u004f\u004e]\u0020\u0050a\u0072\u0073\u0069\u006e\u0067\u0048e\u0061de\u0072\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u0020\u0053\u0075\u0063\u0063\u0065\u0073\u0073\u0066\u0075\u006c\u006c\u0079\u002e\u002e\u002e")
		}
	}()
	var (
		_bgac int
		_gba  uint64
	)
	if _fca = _eef.RegionSegment.parseHeader(); _fca != nil {
		return _fca
	}
	if _, _fca = _eef._efeg.ReadBits(3); _fca != nil {
		return _fca
	}
	_bgac, _fca = _eef._efeg.ReadBit()
	if _fca != nil {
		return _fca
	}
	if _bgac == 1 {
		_eef.UseExtTemplates = true
	}
	_bgac, _fca = _eef._efeg.ReadBit()
	if _fca != nil {
		return _fca
	}
	if _bgac == 1 {
		_eef.IsTPGDon = true
	}
	_gba, _fca = _eef._efeg.ReadBits(2)
	if _fca != nil {
		return _fca
	}
	_eef.GBTemplate = byte(_gba & 0xf)
	_bgac, _fca = _eef._efeg.ReadBit()
	if _fca != nil {
		return _fca
	}
	if _bgac == 1 {
		_eef.IsMMREncoded = true
	}
	if !_eef.IsMMREncoded {
		_dee := 1
		if _eef.GBTemplate == 0 {
			_dee = 4
			if _eef.UseExtTemplates {
				_dee = 12
			}
		}
		if _fca = _eef.readGBAtPixels(_dee); _fca != nil {
			return _fca
		}
	}
	if _fca = _eef.computeSegmentDataStructure(); _fca != nil {
		return _fca
	}
	_aae.Log.Trace("\u0025\u0073", _eef)
	return nil
}
func (_bfe *PatternDictionary) Init(h *Header, r *_g.Reader) error {
	_bfe._dbec = r
	return _bfe.parseHeader()
}
func (_fbdce *SymbolDictionary) decodeHeightClassCollectiveBitmap(_cbfc int64, _dbgcf, _dcde uint32) (*_gf.Bitmap, error) {
	if _cbfc == 0 {
		_fgcc := _gf.New(int(_dcde), int(_dbgcf))
		var (
			_dbfgc byte
			_ddbg  error
		)
		for _baaa := 0; _baaa < len(_fgcc.Data); _baaa++ {
			_dbfgc, _ddbg = _fbdce._caba.ReadByte()
			if _ddbg != nil {
				return nil, _ddbg
			}
			if _ddbg = _fgcc.SetByte(_baaa, _dbfgc); _ddbg != nil {
				return nil, _ddbg
			}
		}
		return _fgcc, nil
	}
	if _fbdce._ddcgf == nil {
		_fbdce._ddcgf = NewGenericRegion(_fbdce._caba)
	}
	_fbdce._ddcgf.setParameters(true, _fbdce._caba.AbsolutePosition(), _cbfc, _dbgcf, _dcde)
	_eggf, _dedga := _fbdce._ddcgf.GetRegionBitmap()
	if _dedga != nil {
		return nil, _dedga
	}
	return _eggf, nil
}

type SegmentEncoder interface {
	Encode(_dbce _g.BinaryWriter) (_afadd int, _aege error)
}

func (_eea *PageInformationSegment) readIsStriped() error {
	_ddbc, _agdg := _eea._faef.ReadBit()
	if _agdg != nil {
		return _agdg
	}
	if _ddbc == 1 {
		_eea.IsStripe = true
	}
	return nil
}
func (_dccg *Header) readHeaderLength(_dba *_g.Reader, _cageb int64) {
	_dccg.HeaderLength = _dba.AbsolutePosition() - _cageb
}
func (_fcaf *TextRegion) decodeDfs() (int64, error) {
	if _fcaf.IsHuffmanEncoded {
		if _fcaf.SbHuffFS == 3 {
			if _fcaf._ddba == nil {
				var _acggc error
				_fcaf._ddba, _acggc = _fcaf.getUserTable(0)
				if _acggc != nil {
					return 0, _acggc
				}
			}
			return _fcaf._ddba.Decode(_fcaf._ecc)
		}
		_cbgd, _ccgab := _fa.GetStandardTable(6 + int(_fcaf.SbHuffFS))
		if _ccgab != nil {
			return 0, _ccgab
		}
		return _cbgd.Decode(_fcaf._ecc)
	}
	_ggde, _aaedc := _fcaf._egdc.DecodeInt(_fcaf._cgdb)
	if _aaedc != nil {
		return 0, _aaedc
	}
	return int64(_ggde), nil
}
func (_cacde *SymbolDictionary) GetDictionary() ([]*_gf.Bitmap, error) {
	_aae.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c-\u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		_aae.Log.Trace("\u005b\u0053\u0059M\u0042\u004f\u004c\u002d\u0044\u0049\u0043\u0054\u0049\u004f\u004e\u0041\u0052\u0059\u005d\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
		_aae.Log.Trace("\u005b\u0053Y\u004d\u0042\u004f\u004c\u002dD\u0049\u0043\u0054\u0049\u004fN\u0041\u0052\u0059\u005d\u0020\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0020\u000a\u0045\u0078\u003a\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u000a\u006e\u0065\u0077\u003a\u0027\u0025\u0073\u0027", _cacde._aebc, _cacde._gedb)
	}()
	if _cacde._aebc == nil {
		var _egda error
		if _cacde.UseRefinementAggregation {
			_cacde._abcgg = _cacde.getSbSymCodeLen()
		}
		if !_cacde.IsHuffmanEncoded {
			if _egda = _cacde.setCodingStatistics(); _egda != nil {
				return nil, _egda
			}
		}
		_cacde._gedb = make([]*_gf.Bitmap, _cacde.NumberOfNewSymbols)
		var _feggf []int
		if _cacde.IsHuffmanEncoded && !_cacde.UseRefinementAggregation {
			_feggf = make([]int, _cacde.NumberOfNewSymbols)
		}
		if _egda = _cacde.setSymbolsArray(); _egda != nil {
			return nil, _egda
		}
		var _ccgf, _ebaa int64
		_cacde._acga = 0
		for _cacde._acga < _cacde.NumberOfNewSymbols {
			_ebaa, _egda = _cacde.decodeHeightClassDeltaHeight()
			if _egda != nil {
				return nil, _egda
			}
			_ccgf += _ebaa
			var _ebgb, _ffega uint32
			_gegfc := int64(_cacde._acga)
			for {
				var _gfc int64
				_gfc, _egda = _cacde.decodeDifferenceWidth()
				if _aa.Is(_egda, _f.ErrOOB) {
					break
				}
				if _egda != nil {
					return nil, _egda
				}
				if _cacde._acga >= _cacde.NumberOfNewSymbols {
					break
				}
				_ebgb += uint32(_gfc)
				_ffega += _ebgb
				if !_cacde.IsHuffmanEncoded || _cacde.UseRefinementAggregation {
					if !_cacde.UseRefinementAggregation {
						_egda = _cacde.decodeDirectlyThroughGenericRegion(_ebgb, uint32(_ccgf))
						if _egda != nil {
							return nil, _egda
						}
					} else {
						_egda = _cacde.decodeAggregate(_ebgb, uint32(_ccgf))
						if _egda != nil {
							return nil, _egda
						}
					}
				} else if _cacde.IsHuffmanEncoded && !_cacde.UseRefinementAggregation {
					_feggf[_cacde._acga] = int(_ebgb)
				}
				_cacde._acga++
			}
			if _cacde.IsHuffmanEncoded && !_cacde.UseRefinementAggregation {
				var _fbfc int64
				if _cacde.SdHuffBMSizeSelection == 0 {
					var _bcbd _fa.Tabler
					_bcbd, _egda = _fa.GetStandardTable(1)
					if _egda != nil {
						return nil, _egda
					}
					_fbfc, _egda = _bcbd.Decode(_cacde._caba)
					if _egda != nil {
						return nil, _egda
					}
				} else {
					_fbfc, _egda = _cacde.huffDecodeBmSize()
					if _egda != nil {
						return nil, _egda
					}
				}
				_cacde._caba.Align()
				var _abaf *_gf.Bitmap
				_abaf, _egda = _cacde.decodeHeightClassCollectiveBitmap(_fbfc, uint32(_ccgf), _ffega)
				if _egda != nil {
					return nil, _egda
				}
				_egda = _cacde.decodeHeightClassBitmap(_abaf, _gegfc, int(_ccgf), _feggf)
				if _egda != nil {
					return nil, _egda
				}
			}
		}
		_cbffa, _egda := _cacde.getToExportFlags()
		if _egda != nil {
			return nil, _egda
		}
		_cacde.setExportedSymbols(_cbffa)
	}
	return _cacde._aebc, nil
}
func (_dbdfa *TableSegment) HtLow() int32 { return _dbdfa._effba }
func (_fdac *TextRegion) InitEncode(globalSymbolsMap, localSymbolsMap map[int]int, comps []int, inLL *_gf.Points, symbols *_gf.Bitmaps, classIDs *_bg.IntSlice, boxes *_gf.Boxes, width, height, symBits int) {
	_fdac.RegionInfo = &RegionSegment{BitmapWidth: uint32(width), BitmapHeight: uint32(height)}
	_fdac._geef = globalSymbolsMap
	_fdac._afdaa = localSymbolsMap
	_fdac._bcfg = comps
	_fdac._bdeea = inLL
	_fdac._debb = symbols
	_fdac._cegd = classIDs
	_fdac._bfea = boxes
	_fdac._geff = symBits
}

type TextRegion struct {
	_ecc                    *_g.Reader
	RegionInfo              *RegionSegment
	SbrTemplate             int8
	SbDsOffset              int8
	DefaultPixel            int8
	CombinationOperator     _gf.CombinationOperator
	IsTransposed            int8
	ReferenceCorner         int16
	LogSBStrips             int16
	UseRefinement           bool
	IsHuffmanEncoded        bool
	SbHuffRSize             int8
	SbHuffRDY               int8
	SbHuffRDX               int8
	SbHuffRDHeight          int8
	SbHuffRDWidth           int8
	SbHuffDT                int8
	SbHuffDS                int8
	SbHuffFS                int8
	SbrATX                  []int8
	SbrATY                  []int8
	NumberOfSymbolInstances uint32
	_abca                   int64
	SbStrips                int8
	NumberOfSymbols         uint32
	RegionBitmap            *_gf.Bitmap
	Symbols                 []*_gf.Bitmap
	_egdc                   *_df.Decoder
	_dagd                   *GenericRefinementRegion
	_gede                   *_df.DecoderStats
	_cgdb                   *_df.DecoderStats
	_dacg                   *_df.DecoderStats
	_fff                    *_df.DecoderStats
	_feeg                   *_df.DecoderStats
	_gbcc                   *_df.DecoderStats
	_acecd                  *_df.DecoderStats
	_bbae                   *_df.DecoderStats
	_ffab                   *_df.DecoderStats
	_fbeg                   *_df.DecoderStats
	_bada                   *_df.DecoderStats
	_efeb                   int8
	_dbcbe                  *_fa.FixedSizeTable
	Header                  *Header
	_ddba                   _fa.Tabler
	_gadb                   _fa.Tabler
	_bbfg                   _fa.Tabler
	_beed                   _fa.Tabler
	_cceff                  _fa.Tabler
	_adgga                  _fa.Tabler
	_bfde                   _fa.Tabler
	_efgefe                 _fa.Tabler
	_geef, _afdaa           map[int]int
	_bcfg                   []int
	_bdeea                  *_gf.Points
	_debb                   *_gf.Bitmaps
	_cegd                   *_bg.IntSlice
	_ddfa, _geff            int
	_bfea                   *_gf.Boxes
}

func (_bcba *SymbolDictionary) checkInput() error {
	if _bcba.SdHuffDecodeHeightSelection == 2 {
		_aae.Log.Debug("\u0053\u0079\u006d\u0062\u006fl\u0020\u0044\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0048\u0065\u0069\u0067\u0068\u0074\u0020\u0053e\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0070\u0065r\u006d\u0069\u0074\u0074\u0065\u0064", _bcba.SdHuffDecodeHeightSelection)
	}
	if _bcba.SdHuffDecodeWidthSelection == 2 {
		_aae.Log.Debug("\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0057\u0069\u0064t\u0068\u0020\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064\u0020\u0076\u0061l\u0075\u0065\u0020\u006e\u006f\u0074 \u0070\u0065r\u006d\u0069t\u0074e\u0064", _bcba.SdHuffDecodeWidthSelection)
	}
	if _bcba.IsHuffmanEncoded {
		if _bcba.SdTemplate != 0 {
			_aae.Log.Debug("\u0053\u0044T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0028\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e \u0030\u0029", _bcba.SdTemplate)
		}
		if !_bcba.UseRefinementAggregation {
			if !_bcba.UseRefinementAggregation {
				if _bcba._dedee {
					_aae.Log.Debug("\u0049\u0073\u0043\u006f\u0064\u0069\u006e\u0067C\u006f\u006e\u0074ex\u0074\u0052\u0065\u0074\u0061\u0069n\u0065\u0064\u0020\u003d\u0020\u0074\u0072\u0075\u0065\u0020\u0028\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0066\u0061\u006cs\u0065\u0029")
					_bcba._dedee = false
				}
				if _bcba._bddf {
					_aae.Log.Debug("\u0069s\u0043\u006fd\u0069\u006e\u0067\u0043o\u006e\u0074\u0065x\u0074\u0055\u0073\u0065\u0064\u0020\u003d\u0020\u0074ru\u0065\u0020\u0028s\u0068\u006fu\u006c\u0064\u0020\u0062\u0065\u0020f\u0061\u006cs\u0065\u0029")
					_bcba._bddf = false
				}
			}
		}
	} else {
		if _bcba.SdHuffBMSizeSelection != 0 {
			_aae.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066B\u004d\u0053\u0069\u007a\u0065\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0030")
			_bcba.SdHuffBMSizeSelection = 0
		}
		if _bcba.SdHuffDecodeWidthSelection != 0 {
			_aae.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0057\u0069\u0064\u0074\u0068\u0053\u0065\u006c\u0065\u0063\u0074\u0069o\u006e\u0020\u0073\u0068\u006fu\u006c\u0064 \u0062\u0065\u0020\u0030")
			_bcba.SdHuffDecodeWidthSelection = 0
		}
		if _bcba.SdHuffDecodeHeightSelection != 0 {
			_aae.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0048e\u0069\u0067\u0068\u0074\u0053\u0065l\u0065\u0063\u0074\u0069\u006f\u006e\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0030")
			_bcba.SdHuffDecodeHeightSelection = 0
		}
	}
	if !_bcba.UseRefinementAggregation {
		if _bcba.SdrTemplate != 0 {
			_aae.Log.Debug("\u0053\u0044\u0052\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u003d\u0020\u0025\u0064\u0020\u0028s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030\u0029", _bcba.SdrTemplate)
			_bcba.SdrTemplate = 0
		}
	}
	if !_bcba.IsHuffmanEncoded || !_bcba.UseRefinementAggregation {
		if _bcba.SdHuffAggInstanceSelection {
			_aae.Log.Debug("\u0053d\u0048\u0075f\u0066\u0041\u0067g\u0049\u006e\u0073\u0074\u0061\u006e\u0063e\u0053\u0065\u006c\u0065\u0063\u0074i\u006f\u006e\u0020\u003d\u0020\u0025\u0064\u0020\u0028\u0073\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030\u0029", _bcba.SdHuffAggInstanceSelection)
		}
	}
	return nil
}
func (_dgge *TextRegion) decodeRI() (int64, error) {
	if !_dgge.UseRefinement {
		return 0, nil
	}
	if _dgge.IsHuffmanEncoded {
		_dged, _aaaed := _dgge._ecc.ReadBit()
		return int64(_dged), _aaaed
	}
	_bacce, _gadca := _dgge._egdc.DecodeInt(_dgge._feeg)
	return int64(_bacce), _gadca
}
func (_fcf *SymbolDictionary) setRetainedCodingContexts(_bcf *SymbolDictionary) {
	_fcf._dgcb = _bcf._dgcb
	_fcf.IsHuffmanEncoded = _bcf.IsHuffmanEncoded
	_fcf.UseRefinementAggregation = _bcf.UseRefinementAggregation
	_fcf.SdTemplate = _bcf.SdTemplate
	_fcf.SdrTemplate = _bcf.SdrTemplate
	_fcf.SdATX = _bcf.SdATX
	_fcf.SdATY = _bcf.SdATY
	_fcf.SdrATX = _bcf.SdrATX
	_fcf.SdrATY = _bcf.SdrATY
	_fcf._afbd = _bcf._afbd
}

type PatternDictionary struct {
	_dbec            *_g.Reader
	DataHeaderOffset int64
	DataHeaderLength int64
	DataOffset       int64
	DataLength       int64
	GBAtX            []int8
	GBAtY            []int8
	IsMMREncoded     bool
	HDTemplate       byte
	HdpWidth         byte
	HdpHeight        byte
	Patterns         []*_gf.Bitmap
	GrayMax          uint32
}

func (_afcf *TextRegion) computeSymbolCodeLength() error {
	if _afcf.IsHuffmanEncoded {
		return _afcf.symbolIDCodeLengths()
	}
	_afcf._efeb = int8(_e.Ceil(_e.Log(float64(_afcf.NumberOfSymbols)) / _e.Log(2)))
	return nil
}
func (_dgdb *SymbolDictionary) decodeHeightClassBitmap(_aadgg *_gf.Bitmap, _dff int64, _cfee int, _effa []int) error {
	for _dfe := _dff; _dfe < int64(_dgdb._acga); _dfe++ {
		var _edgd int
		for _edea := _dff; _edea <= _dfe-1; _edea++ {
			_edgd += _effa[_edea]
		}
		_cbcd := _a.Rect(_edgd, 0, _edgd+_effa[_dfe], _cfee)
		_gbcb, _cgga := _gf.Extract(_cbcd, _aadgg)
		if _cgga != nil {
			return _cgga
		}
		_dgdb._gedb[_dfe] = _gbcb
		_dgdb._ebfb = append(_dgdb._ebfb, _gbcb)
	}
	return nil
}
func (_dca *PageInformationSegment) readDefaultPixelValue() error {
	_gegf, _dcda := _dca._faef.ReadBit()
	if _dcda != nil {
		return _dcda
	}
	_dca.DefaultPixelValue = uint8(_gegf & 0xf)
	return nil
}
func (_dbfg *GenericRefinementRegion) decodeTypicalPredictedLineTemplate1(_aef, _def, _caa, _ebe, _cddc, _fbg, _bf, _eabc, _fed int) (_bdd error) {
	var (
		_efe, _aaa int
		_cdg, _fdc int
		_eda, _dgd int
		_dge       byte
	)
	if _aef > 0 {
		_dge, _bdd = _dbfg.RegionBitmap.GetByte(_bf - _caa)
		if _bdd != nil {
			return
		}
		_cdg = int(_dge)
	}
	if _eabc > 0 && _eabc <= _dbfg.ReferenceBitmap.Height {
		_dge, _bdd = _dbfg.ReferenceBitmap.GetByte(_fed - _ebe + _fbg)
		if _bdd != nil {
			return
		}
		_fdc = int(_dge) << 2
	}
	if _eabc >= 0 && _eabc < _dbfg.ReferenceBitmap.Height {
		_dge, _bdd = _dbfg.ReferenceBitmap.GetByte(_fed + _fbg)
		if _bdd != nil {
			return
		}
		_eda = int(_dge)
	}
	if _eabc > -2 && _eabc < _dbfg.ReferenceBitmap.Height-1 {
		_dge, _bdd = _dbfg.ReferenceBitmap.GetByte(_fed + _ebe + _fbg)
		if _bdd != nil {
			return
		}
		_dgd = int(_dge)
	}
	_efe = ((_cdg >> 5) & 0x6) | ((_dgd >> 2) & 0x30) | (_eda & 0xc0) | (_fdc & 0x200)
	_aaa = ((_dgd >> 2) & 0x70) | (_eda & 0xc0) | (_fdc & 0x700)
	var _cce int
	for _agg := 0; _agg < _cddc; _agg = _cce {
		var (
			_dac int
			_gge int
		)
		_cce = _agg + 8
		if _dac = _def - _agg; _dac > 8 {
			_dac = 8
		}
		_acb := _cce < _def
		_fee := _cce < _dbfg.ReferenceBitmap.Width
		_afa := _fbg + 1
		if _aef > 0 {
			_dge = 0
			if _acb {
				_dge, _bdd = _dbfg.RegionBitmap.GetByte(_bf - _caa + 1)
				if _bdd != nil {
					return
				}
			}
			_cdg = (_cdg << 8) | int(_dge)
		}
		if _eabc > 0 && _eabc <= _dbfg.ReferenceBitmap.Height {
			var _ebc int
			if _fee {
				_dge, _bdd = _dbfg.ReferenceBitmap.GetByte(_fed - _ebe + _afa)
				if _bdd != nil {
					return
				}
				_ebc = int(_dge) << 2
			}
			_fdc = (_fdc << 8) | _ebc
		}
		if _eabc >= 0 && _eabc < _dbfg.ReferenceBitmap.Height {
			_dge = 0
			if _fee {
				_dge, _bdd = _dbfg.ReferenceBitmap.GetByte(_fed + _afa)
				if _bdd != nil {
					return
				}
			}
			_eda = (_eda << 8) | int(_dge)
		}
		if _eabc > -2 && _eabc < (_dbfg.ReferenceBitmap.Height-1) {
			_dge = 0
			if _fee {
				_dge, _bdd = _dbfg.ReferenceBitmap.GetByte(_fed + _ebe + _afa)
				if _bdd != nil {
					return
				}
			}
			_dgd = (_dgd << 8) | int(_dge)
		}
		for _bbad := 0; _bbad < _dac; _bbad++ {
			var _cgd int
			_cfa := (_aaa >> 4) & 0x1ff
			switch _cfa {
			case 0x1ff:
				_cgd = 1
			case 0x00:
				_cgd = 0
			default:
				_dbfg._cd.SetIndex(int32(_efe))
				_cgd, _bdd = _dbfg._dc.DecodeBit(_dbfg._cd)
				if _bdd != nil {
					return
				}
			}
			_baf := uint(7 - _bbad)
			_gge |= _cgd << _baf
			_efe = ((_efe & 0x0d6) << 1) | _cgd | (_cdg>>_baf+5)&0x002 | ((_dgd>>_baf + 2) & 0x010) | ((_eda >> _baf) & 0x040) | ((_fdc >> _baf) & 0x200)
			_aaa = ((_aaa & 0xdb) << 1) | ((_dgd>>_baf + 2) & 0x010) | ((_eda >> _baf) & 0x080) | ((_fdc >> _baf) & 0x400)
		}
		_bdd = _dbfg.RegionBitmap.SetByte(_bf, byte(_gge))
		if _bdd != nil {
			return
		}
		_bf++
		_fed++
	}
	return nil
}
func (_fbaa *RegionSegment) String() string {
	_gacb := &_d.Builder{}
	_gacb.WriteString("\u0009[\u0052E\u0047\u0049\u004f\u004e\u0020S\u0045\u0047M\u0045\u004e\u0054\u005d\u000a")
	_gacb.WriteString(_af.Sprintf("\t\u0009\u002d\u0020\u0042\u0069\u0074m\u0061\u0070\u0020\u0028\u0077\u0069d\u0074\u0068\u002c\u0020\u0068\u0065\u0069g\u0068\u0074\u0029\u0020\u005b\u0025\u0064\u0078\u0025\u0064]\u000a", _fbaa.BitmapWidth, _fbaa.BitmapHeight))
	_gacb.WriteString(_af.Sprintf("\u0009\u0009\u002d\u0020L\u006f\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0028\u0078,\u0079)\u003a\u0020\u005b\u0025\u0064\u002c\u0025d\u005d\u000a", _fbaa.XLocation, _fbaa.YLocation))
	_gacb.WriteString(_af.Sprintf("\t\u0009\u002d\u0020\u0043\u006f\u006db\u0069\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0073", _fbaa.CombinaionOperator))
	return _gacb.String()
}

type EncodeInitializer interface{ InitEncode() }

func (_efb *PatternDictionary) readGrayMax() error {
	_abcg, _cebb := _efb._dbec.ReadBits(32)
	if _cebb != nil {
		return _cebb
	}
	_efb.GrayMax = uint32(_abcg & _e.MaxUint32)
	return nil
}

type Type int

func (_daad *TextRegion) encodeSymbols(_fbga _g.BinaryWriter) (_efgc int, _daff error) {
	const _gebc = "\u0065\u006e\u0063\u006f\u0064\u0065\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	_fecdb := make([]byte, 4)
	_ag.BigEndian.PutUint32(_fecdb, _daad.NumberOfSymbols)
	if _efgc, _daff = _fbga.Write(_fecdb); _daff != nil {
		return _efgc, _edb.Wrap(_daff, _gebc, "\u004e\u0075\u006dbe\u0072\u004f\u0066\u0053\u0079\u006d\u0062\u006f\u006c\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073")
	}
	_cgdabg, _daff := _gf.NewClassedPoints(_daad._bdeea, _daad._bcfg)
	if _daff != nil {
		return 0, _edb.Wrap(_daff, _gebc, "")
	}
	var _cdbf, _eaab int
	_dcba := _afe.New()
	_dcba.Init()
	if _daff = _dcba.EncodeInteger(_afe.IADT, 0); _daff != nil {
		return _efgc, _edb.Wrap(_daff, _gebc, "\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0020\u0044\u0054")
	}
	_fdd, _daff := _cgdabg.GroupByY()
	if _daff != nil {
		return 0, _edb.Wrap(_daff, _gebc, "")
	}
	for _, _dfcd := range _fdd {
		_dbcbg := int(_dfcd.YAtIndex(0))
		_fbdb := _dbcbg - _cdbf
		if _daff = _dcba.EncodeInteger(_afe.IADT, _fbdb); _daff != nil {
			return _efgc, _edb.Wrap(_daff, _gebc, "")
		}
		var _gagg int
		for _egbd, _bdgf := range _dfcd.IntSlice {
			switch _egbd {
			case 0:
				_aecf := int(_dfcd.XAtIndex(_egbd)) - _eaab
				if _daff = _dcba.EncodeInteger(_afe.IAFS, _aecf); _daff != nil {
					return _efgc, _edb.Wrap(_daff, _gebc, "")
				}
				_eaab += _aecf
				_gagg = _eaab
			default:
				_ccbe := int(_dfcd.XAtIndex(_egbd)) - _gagg
				if _daff = _dcba.EncodeInteger(_afe.IADS, _ccbe); _daff != nil {
					return _efgc, _edb.Wrap(_daff, _gebc, "")
				}
				_gagg += _ccbe
			}
			_afdg, _cdda := _daad._cegd.Get(_bdgf)
			if _cdda != nil {
				return _efgc, _edb.Wrap(_cdda, _gebc, "")
			}
			_gbag, _daac := _daad._geef[_afdg]
			if !_daac {
				_gbag, _daac = _daad._afdaa[_afdg]
				if !_daac {
					return _efgc, _edb.Errorf(_gebc, "\u0053\u0079\u006d\u0062\u006f\u006c:\u0020\u0027\u0025d\u0027\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064 \u0069\u006e\u0020\u0067\u006cob\u0061\u006c\u0020\u0061\u006e\u0064\u0020\u006c\u006f\u0063\u0061\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0020\u006d\u0061\u0070", _afdg)
				}
			}
			if _cdda = _dcba.EncodeIAID(_daad._geff, _gbag); _cdda != nil {
				return _efgc, _edb.Wrap(_cdda, _gebc, "")
			}
		}
		if _daff = _dcba.EncodeOOB(_afe.IADS); _daff != nil {
			return _efgc, _edb.Wrap(_daff, _gebc, "")
		}
	}
	_dcba.Final()
	_gecf, _daff := _dcba.WriteTo(_fbga)
	if _daff != nil {
		return _efgc, _edb.Wrap(_daff, _gebc, "")
	}
	_efgc += int(_gecf)
	return _efgc, nil
}
func (_ccaa *TextRegion) initSymbols() error {
	const _geae = "i\u006e\u0069\u0074\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	for _, _aaggd := range _ccaa.Header.RTSegments {
		if _aaggd == nil {
			return _edb.Error(_geae, "\u006e\u0069\u006c\u0020\u0073\u0065\u0067\u006de\u006e\u0074\u0020pr\u006f\u0076\u0069\u0064\u0065\u0064 \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0074\u0065\u0078\u0074\u0020\u0072\u0065g\u0069\u006f\u006e\u0020\u0053\u0079\u006d\u0062o\u006c\u0073")
		}
		if _aaggd.Type == 0 {
			_daaa, _fbad := _aaggd.GetSegmentData()
			if _fbad != nil {
				return _edb.Wrap(_fbad, _geae, "")
			}
			_bbdce, _aefe := _daaa.(*SymbolDictionary)
			if !_aefe {
				return _edb.Error(_geae, "\u0072e\u0066\u0065r\u0072\u0065\u0064 \u0054\u006f\u0020\u0053\u0065\u0067\u006de\u006e\u0074\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0053\u0079\u006d\u0062\u006f\u006cD\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
			}
			_bbdce._ggc = _ccaa._bbae
			_aafe, _fbad := _bbdce.GetDictionary()
			if _fbad != nil {
				return _edb.Wrap(_fbad, _geae, "")
			}
			_ccaa.Symbols = append(_ccaa.Symbols, _aafe...)
		}
	}
	_ccaa.NumberOfSymbols = uint32(len(_ccaa.Symbols))
	return nil
}

type Regioner interface {
	GetRegionBitmap() (*_gf.Bitmap, error)
	GetRegionInfo() *RegionSegment
}

func (_aaf *Header) writeSegmentDataLength(_acdf _g.BinaryWriter) (_bgd int, _ada error) {
	_gbef := make([]byte, 4)
	_ag.BigEndian.PutUint32(_gbef, uint32(_aaf.SegmentDataLength))
	if _bgd, _ada = _acdf.Write(_gbef); _ada != nil {
		return 0, _edb.Wrap(_ada, "\u0048\u0065a\u0064\u0065\u0072\u002e\u0077\u0072\u0069\u0074\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0044\u0061\u0074\u0061\u004c\u0065ng\u0074\u0068", "")
	}
	return _bgd, nil
}
func (_abgdb *SymbolDictionary) readRefinementAtPixels(_bbgab int) error {
	_abgdb.SdrATX = make([]int8, _bbgab)
	_abgdb.SdrATY = make([]int8, _bbgab)
	var (
		_gfce byte
		_fgab error
	)
	for _acaf := 0; _acaf < _bbgab; _acaf++ {
		_gfce, _fgab = _abgdb._caba.ReadByte()
		if _fgab != nil {
			return _fgab
		}
		_abgdb.SdrATX[_acaf] = int8(_gfce)
		_gfce, _fgab = _abgdb._caba.ReadByte()
		if _fgab != nil {
			return _fgab
		}
		_abgdb.SdrATY[_acaf] = int8(_gfce)
	}
	return nil
}
func (_de *GenericRefinementRegion) decodeTypicalPredictedLine(_abb, _cc, _eab, _edgg, _ac, _afc int) error {
	_fad := _abb - int(_de.ReferenceDY)
	_cbf := _de.ReferenceBitmap.GetByteIndex(0, _fad)
	_ec := _de.RegionBitmap.GetByteIndex(0, _abb)
	var _eg error
	switch _de.TemplateID {
	case 0:
		_eg = _de.decodeTypicalPredictedLineTemplate0(_abb, _cc, _eab, _edgg, _ac, _afc, _ec, _fad, _cbf)
	case 1:
		_eg = _de.decodeTypicalPredictedLineTemplate1(_abb, _cc, _eab, _edgg, _ac, _afc, _ec, _fad, _cbf)
	}
	return _eg
}
func (_aded *Header) referenceSize() uint {
	switch {
	case _aded.SegmentNumber <= 255:
		return 1
	case _aded.SegmentNumber <= 65535:
		return 2
	default:
		return 4
	}
}
func (_gecc *Header) readReferredToSegmentNumbers(_acbae *_g.Reader, _eagc int) ([]int, error) {
	const _cad = "\u0072\u0065\u0061\u0064R\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0054\u006f\u0053e\u0067m\u0065\u006e\u0074\u004e\u0075\u006d\u0062e\u0072\u0073"
	_bbgd := make([]int, _eagc)
	if _eagc > 0 {
		_gecc.RTSegments = make([]*Header, _eagc)
		var (
			_efea uint64
			_begg error
		)
		for _dgga := 0; _dgga < _eagc; _dgga++ {
			_efea, _begg = _acbae.ReadBits(byte(_gecc.referenceSize()) << 3)
			if _begg != nil {
				return nil, _edb.Wrapf(_begg, _cad, "\u0027\u0025\u0064\u0027 \u0072\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0020\u0073e\u0067m\u0065\u006e\u0074\u0020\u006e\u0075\u006db\u0065\u0072", _dgga)
			}
			_bbgd[_dgga] = int(_efea & _e.MaxInt32)
		}
	}
	return _bbgd, nil
}
func (_bgab *PageInformationSegment) readContainsRefinement() error {
	_gfba, _aegeb := _bgab._faef.ReadBit()
	if _aegeb != nil {
		return _aegeb
	}
	if _gfba == 1 {
		_bgab._ccdd = true
	}
	return nil
}
func (_aefcd *PatternDictionary) computeSegmentDataStructure() error {
	_aefcd.DataOffset = _aefcd._dbec.AbsolutePosition()
	_aefcd.DataHeaderLength = _aefcd.DataOffset - _aefcd.DataHeaderOffset
	_aefcd.DataLength = int64(_aefcd._dbec.AbsoluteLength()) - _aefcd.DataHeaderLength
	return nil
}

type OrganizationType uint8

func (_edeb *GenericRegion) setParametersWithAt(_eae bool, _eegb byte, _cge, _eaf bool, _bcebe, _aeda []int8, _acdd, _bfca uint32, _fccg *_df.DecoderStats, _cfe *_df.Decoder) {
	_edeb.IsMMREncoded = _eae
	_edeb.GBTemplate = _eegb
	_edeb.IsTPGDon = _cge
	_edeb.GBAtX = _bcebe
	_edeb.GBAtY = _aeda
	_edeb.RegionSegment.BitmapHeight = _bfca
	_edeb.RegionSegment.BitmapWidth = _acdd
	_edeb._gdf = nil
	_edeb.Bitmap = nil
	if _fccg != nil {
		_edeb._fac = _fccg
	}
	if _cfe != nil {
		_edeb._ddga = _cfe
	}
	_aae.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0047\u0049O\u004e\u005d\u0020\u0073\u0065\u0074P\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0053\u0044\u0041t\u003a\u0020\u0025\u0073", _edeb)
}
func (_fbcf *Header) Encode(w _g.BinaryWriter) (_ggedc int, _afea error) {
	const _dafg = "\u0048\u0065\u0061d\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065"
	var _agbe _g.BinaryWriter
	_aae.Log.Trace("\u005b\u0053\u0045G\u004d\u0045\u004e\u0054-\u0048\u0045\u0041\u0044\u0045\u0052\u005d[\u0045\u004e\u0043\u004f\u0044\u0045\u005d\u0020\u0042\u0065\u0067\u0069\u006e\u0073")
	defer func() {
		if _afea != nil {
			_aae.Log.Trace("[\u0053\u0045\u0047\u004d\u0045\u004eT\u002d\u0048\u0045\u0041\u0044\u0045R\u005d\u005b\u0045\u004e\u0043\u004f\u0044E\u005d\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u002e\u0020%\u0076", _afea)
		} else {
			_aae.Log.Trace("\u005b\u0053\u0045\u0047ME\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0025\u0076", _fbcf)
			_aae.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045N\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u005b\u0045\u004e\u0043O\u0044\u0045\u005d\u0020\u0046\u0069\u006ei\u0073\u0068\u0065\u0064")
		}
	}()
	w.FinishByte()
	if _fbcf.SegmentData != nil {
		_fbbc, _fecg := _fbcf.SegmentData.(SegmentEncoder)
		if !_fecg {
			return 0, _edb.Errorf(_dafg, "\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0054\u0020\u0064\u006f\u0065s\u006e\u0027\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074 \u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0045\u006e\u0063\u006f\u0064er\u0020\u0069\u006e\u0074\u0065\u0072\u0066\u0061\u0063\u0065", _fbcf.SegmentData)
		}
		_agbe = _g.BufferedMSB()
		_ggedc, _afea = _fbbc.Encode(_agbe)
		if _afea != nil {
			return 0, _edb.Wrap(_afea, _dafg, "")
		}
		_fbcf.SegmentDataLength = uint64(_ggedc)
	}
	if _fbcf.pageSize() == 4 {
		_fbcf.PageAssociationFieldSize = true
	}
	var _fecc int
	_fecc, _afea = _fbcf.writeSegmentNumber(w)
	if _afea != nil {
		return 0, _edb.Wrap(_afea, _dafg, "")
	}
	_ggedc += _fecc
	if _afea = _fbcf.writeFlags(w); _afea != nil {
		return _ggedc, _edb.Wrap(_afea, _dafg, "")
	}
	_ggedc++
	_fecc, _afea = _fbcf.writeReferredToCount(w)
	if _afea != nil {
		return 0, _edb.Wrap(_afea, _dafg, "")
	}
	_ggedc += _fecc
	_fecc, _afea = _fbcf.writeReferredToSegments(w)
	if _afea != nil {
		return 0, _edb.Wrap(_afea, _dafg, "")
	}
	_ggedc += _fecc
	_fecc, _afea = _fbcf.writeSegmentPageAssociation(w)
	if _afea != nil {
		return 0, _edb.Wrap(_afea, _dafg, "")
	}
	_ggedc += _fecc
	_fecc, _afea = _fbcf.writeSegmentDataLength(w)
	if _afea != nil {
		return 0, _edb.Wrap(_afea, _dafg, "")
	}
	_ggedc += _fecc
	_fbcf.HeaderLength = int64(_ggedc) - int64(_fbcf.SegmentDataLength)
	if _agbe != nil {
		if _, _afea = w.Write(_agbe.Data()); _afea != nil {
			return _ggedc, _edb.Wrap(_afea, _dafg, "\u0077r\u0069t\u0065\u0020\u0073\u0065\u0067m\u0065\u006et\u0020\u0064\u0061\u0074\u0061")
		}
	}
	return _ggedc, nil
}
func (_bdff *SymbolDictionary) readAtPixels(_afgg int) error {
	_bdff.SdATX = make([]int8, _afgg)
	_bdff.SdATY = make([]int8, _afgg)
	var (
		_dfee byte
		_effb error
	)
	for _cddf := 0; _cddf < _afgg; _cddf++ {
		_dfee, _effb = _bdff._caba.ReadByte()
		if _effb != nil {
			return _effb
		}
		_bdff.SdATX[_cddf] = int8(_dfee)
		_dfee, _effb = _bdff._caba.ReadByte()
		if _effb != nil {
			return _effb
		}
		_bdff.SdATY[_cddf] = int8(_dfee)
	}
	return nil
}
func (_fe *EndOfStripe) LineNumber() int { return _fe._bc }
func (_ffefg *TextRegion) decodeIb(_cfaac, _agba int64) (*_gf.Bitmap, error) {
	const _ddecd = "\u0064\u0065\u0063\u006f\u0064\u0065\u0049\u0062"
	var (
		_bfge  error
		_ebgbc *_gf.Bitmap
	)
	if _cfaac == 0 {
		if int(_agba) > len(_ffefg.Symbols)-1 {
			return nil, _edb.Error(_ddecd, "\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0049\u0042\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0020\u0069\u006e\u0064\u0065x\u0020\u006f\u0075\u0074\u0020o\u0066\u0020r\u0061\u006e\u0067\u0065")
		}
		return _ffefg.Symbols[int(_agba)], nil
	}
	var _bfeg, _gab, _gffae, _eee int64
	_bfeg, _bfge = _ffefg.decodeRdw()
	if _bfge != nil {
		return nil, _edb.Wrap(_bfge, _ddecd, "")
	}
	_gab, _bfge = _ffefg.decodeRdh()
	if _bfge != nil {
		return nil, _edb.Wrap(_bfge, _ddecd, "")
	}
	_gffae, _bfge = _ffefg.decodeRdx()
	if _bfge != nil {
		return nil, _edb.Wrap(_bfge, _ddecd, "")
	}
	_eee, _bfge = _ffefg.decodeRdy()
	if _bfge != nil {
		return nil, _edb.Wrap(_bfge, _ddecd, "")
	}
	if _ffefg.IsHuffmanEncoded {
		if _, _bfge = _ffefg.decodeSymInRefSize(); _bfge != nil {
			return nil, _edb.Wrap(_bfge, _ddecd, "")
		}
		_ffefg._ecc.Align()
	}
	_abd := _ffefg.Symbols[_agba]
	_cedgc := uint32(_abd.Width)
	_bedae := uint32(_abd.Height)
	_aefb := int32(uint32(_bfeg)>>1) + int32(_gffae)
	_adga := int32(uint32(_gab)>>1) + int32(_eee)
	if _ffefg._dagd == nil {
		_ffefg._dagd = _agf(_ffefg._ecc, nil)
	}
	_ffefg._dagd.setParameters(_ffefg._bada, _ffefg._egdc, _ffefg.SbrTemplate, _cedgc+uint32(_bfeg), _bedae+uint32(_gab), _abd, _aefb, _adga, false, _ffefg.SbrATX, _ffefg.SbrATY)
	_ebgbc, _bfge = _ffefg._dagd.GetRegionBitmap()
	if _bfge != nil {
		return nil, _edb.Wrap(_bfge, _ddecd, "\u0067\u0072\u0066")
	}
	if _ffefg.IsHuffmanEncoded {
		_ffefg._ecc.Align()
	}
	return _ebgbc, nil
}
func (_eegf *RegionSegment) parseHeader() error {
	const _caea = "p\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_aae.Log.Trace("\u005b\u0052\u0045\u0047I\u004f\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002dH\u0045A\u0044\u0045\u0052\u005d\u0020\u0042\u0065g\u0069\u006e")
	defer func() {
		_aae.Log.Trace("\u005b\u0052\u0045G\u0049\u004f\u004e\u005d[\u0050\u0041\u0052\u0053\u0045\u002d\u0048E\u0041\u0044\u0045\u0052\u005d\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
	}()
	_bgc, _gefb := _eegf._ebecg.ReadBits(32)
	if _gefb != nil {
		return _edb.Wrap(_gefb, _caea, "\u0077\u0069\u0064t\u0068")
	}
	_eegf.BitmapWidth = uint32(_bgc & _e.MaxUint32)
	_bgc, _gefb = _eegf._ebecg.ReadBits(32)
	if _gefb != nil {
		return _edb.Wrap(_gefb, _caea, "\u0068\u0065\u0069\u0067\u0068\u0074")
	}
	_eegf.BitmapHeight = uint32(_bgc & _e.MaxUint32)
	_bgc, _gefb = _eegf._ebecg.ReadBits(32)
	if _gefb != nil {
		return _edb.Wrap(_gefb, _caea, "\u0078\u0020\u006c\u006f\u0063\u0061\u0074\u0069\u006f\u006e")
	}
	_eegf.XLocation = uint32(_bgc & _e.MaxUint32)
	_bgc, _gefb = _eegf._ebecg.ReadBits(32)
	if _gefb != nil {
		return _edb.Wrap(_gefb, _caea, "\u0079\u0020\u006c\u006f\u0063\u0061\u0074\u0069\u006f\u006e")
	}
	_eegf.YLocation = uint32(_bgc & _e.MaxUint32)
	if _, _gefb = _eegf._ebecg.ReadBits(5); _gefb != nil {
		return _edb.Wrap(_gefb, _caea, "\u0064i\u0072\u0079\u0020\u0072\u0065\u0061d")
	}
	if _gefb = _eegf.readCombinationOperator(); _gefb != nil {
		return _edb.Wrap(_gefb, _caea, "c\u006fm\u0062\u0069\u006e\u0061\u0074\u0069\u006f\u006e \u006f\u0070\u0065\u0072at\u006f\u0072")
	}
	return nil
}
func (_bgff *Header) parse(_agga Documenter, _bgagd *_g.Reader, _ceac int64, _cbbd OrganizationType) (_aggf error) {
	const _gffa = "\u0070\u0061\u0072s\u0065"
	_aae.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u002d\u0048E\u0041\u0044\u0045\u0052\u005d\u005b\u0050A\u0052\u0053\u0045\u005d\u0020\u0042\u0065\u0067\u0069\u006e\u0073")
	defer func() {
		if _aggf != nil {
			_aae.Log.Trace("\u005b\u0053\u0045GM\u0045\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045R\u005d[\u0050A\u0052S\u0045\u005d\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _aggf)
		} else {
			_aae.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052]\u005bP\u0041\u0052\u0053\u0045\u005d\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
		}
	}()
	_, _aggf = _bgagd.Seek(_ceac, _ed.SeekStart)
	if _aggf != nil {
		return _edb.Wrap(_aggf, _gffa, "\u0073\u0065\u0065\u006b\u0020\u0073\u0074\u0061\u0072\u0074")
	}
	if _aggf = _bgff.readSegmentNumber(_bgagd); _aggf != nil {
		return _edb.Wrap(_aggf, _gffa, "")
	}
	if _aggf = _bgff.readHeaderFlags(); _aggf != nil {
		return _edb.Wrap(_aggf, _gffa, "")
	}
	var _acddg uint64
	_acddg, _aggf = _bgff.readNumberOfReferredToSegments(_bgagd)
	if _aggf != nil {
		return _edb.Wrap(_aggf, _gffa, "")
	}
	_bgff.RTSNumbers, _aggf = _bgff.readReferredToSegmentNumbers(_bgagd, int(_acddg))
	if _aggf != nil {
		return _edb.Wrap(_aggf, _gffa, "")
	}
	_aggf = _bgff.readSegmentPageAssociation(_agga, _bgagd, _acddg, _bgff.RTSNumbers...)
	if _aggf != nil {
		return _edb.Wrap(_aggf, _gffa, "")
	}
	if _bgff.Type != TEndOfFile {
		if _aggf = _bgff.readSegmentDataLength(_bgagd); _aggf != nil {
			return _edb.Wrap(_aggf, _gffa, "")
		}
	}
	_bgff.readDataStartOffset(_bgagd, _cbbd)
	_bgff.readHeaderLength(_bgagd, _ceac)
	_aae.Log.Trace("\u0025\u0073", _bgff)
	return nil
}
func (_fcce *PageInformationSegment) CombinationOperatorOverrideAllowed() bool { return _fcce._fbbge }
func (_fdga *TextRegion) readRegionFlags() error {
	var (
		_gedg int
		_bdaf uint64
		_cebg error
	)
	_gedg, _cebg = _fdga._ecc.ReadBit()
	if _cebg != nil {
		return _cebg
	}
	_fdga.SbrTemplate = int8(_gedg)
	_bdaf, _cebg = _fdga._ecc.ReadBits(5)
	if _cebg != nil {
		return _cebg
	}
	_fdga.SbDsOffset = int8(_bdaf)
	if _fdga.SbDsOffset > 0x0f {
		_fdga.SbDsOffset -= 0x20
	}
	_gedg, _cebg = _fdga._ecc.ReadBit()
	if _cebg != nil {
		return _cebg
	}
	_fdga.DefaultPixel = int8(_gedg)
	_bdaf, _cebg = _fdga._ecc.ReadBits(2)
	if _cebg != nil {
		return _cebg
	}
	_fdga.CombinationOperator = _gf.CombinationOperator(int(_bdaf) & 0x3)
	_gedg, _cebg = _fdga._ecc.ReadBit()
	if _cebg != nil {
		return _cebg
	}
	_fdga.IsTransposed = int8(_gedg)
	_bdaf, _cebg = _fdga._ecc.ReadBits(2)
	if _cebg != nil {
		return _cebg
	}
	_fdga.ReferenceCorner = int16(_bdaf) & 0x3
	_bdaf, _cebg = _fdga._ecc.ReadBits(2)
	if _cebg != nil {
		return _cebg
	}
	_fdga.LogSBStrips = int16(_bdaf) & 0x3
	_fdga.SbStrips = 1 << uint(_fdga.LogSBStrips)
	_gedg, _cebg = _fdga._ecc.ReadBit()
	if _cebg != nil {
		return _cebg
	}
	if _gedg == 1 {
		_fdga.UseRefinement = true
	}
	_gedg, _cebg = _fdga._ecc.ReadBit()
	if _cebg != nil {
		return _cebg
	}
	if _gedg == 1 {
		_fdga.IsHuffmanEncoded = true
	}
	return nil
}
func (_afbe *Header) readSegmentPageAssociation(_fbea Documenter, _afada *_g.Reader, _baba uint64, _cdc ...int) (_cdb error) {
	const _aadea = "\u0072\u0065\u0061\u0064\u0053\u0065\u0067\u006d\u0065\u006e\u0074P\u0061\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061t\u0069\u006f\u006e"
	if !_afbe.PageAssociationFieldSize {
		_gde, _acfg := _afada.ReadBits(8)
		if _acfg != nil {
			return _edb.Wrap(_acfg, _aadea, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		_afbe.PageAssociation = int(_gde & 0xFF)
	} else {
		_edde, _cdbb := _afada.ReadBits(32)
		if _cdbb != nil {
			return _edb.Wrap(_cdbb, _aadea, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		_afbe.PageAssociation = int(_edde & _e.MaxInt32)
	}
	if _baba == 0 {
		return nil
	}
	if _afbe.PageAssociation != 0 {
		_ebee, _dcgd := _fbea.GetPage(_afbe.PageAssociation)
		if _dcgd != nil {
			return _edb.Wrap(_dcgd, _aadea, "\u0061s\u0073\u006f\u0063\u0069a\u0074\u0065\u0064\u0020\u0070a\u0067e\u0020n\u006f\u0074\u0020\u0066\u006f\u0075\u006ed")
		}
		var _defg int
		for _bcee := uint64(0); _bcee < _baba; _bcee++ {
			_defg = _cdc[_bcee]
			_afbe.RTSegments[_bcee], _dcgd = _ebee.GetSegment(_defg)
			if _dcgd != nil {
				var _fcda error
				_afbe.RTSegments[_bcee], _fcda = _fbea.GetGlobalSegment(_defg)
				if _fcda != nil {
					return _edb.Wrapf(_dcgd, _aadea, "\u0072\u0065\u0066\u0065\u0072\u0065n\u0063\u0065\u0020s\u0065\u0067\u006de\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0020\u0061\u0074\u0020pa\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0072\u0020\u0069\u006e\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0073", _afbe.PageAssociation)
				}
			}
		}
		return nil
	}
	for _dfbf := uint64(0); _dfbf < _baba; _dfbf++ {
		_afbe.RTSegments[_dfbf], _cdb = _fbea.GetGlobalSegment(_cdc[_dfbf])
		if _cdb != nil {
			return _edb.Wrapf(_cdb, _aadea, "\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u003a\u0020\u0027\u0025d\u0027\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", _cdc[_dfbf])
		}
	}
	return nil
}
func (_efg *GenericRefinementRegion) decodeTypicalPredictedLineTemplate0(_bce, _dec, _gg, _ede, _bd, _dgf, _ddc, _fce, _gbg int) error {
	var (
		_aag, _bad, _fcea, _ceg, _aga, _dbf int
		_dgc                                byte
		_bga                                error
	)
	if _bce > 0 {
		_dgc, _bga = _efg.RegionBitmap.GetByte(_ddc - _gg)
		if _bga != nil {
			return _bga
		}
		_fcea = int(_dgc)
	}
	if _fce > 0 && _fce <= _efg.ReferenceBitmap.Height {
		_dgc, _bga = _efg.ReferenceBitmap.GetByte(_gbg - _ede + _dgf)
		if _bga != nil {
			return _bga
		}
		_ceg = int(_dgc) << 4
	}
	if _fce >= 0 && _fce < _efg.ReferenceBitmap.Height {
		_dgc, _bga = _efg.ReferenceBitmap.GetByte(_gbg + _dgf)
		if _bga != nil {
			return _bga
		}
		_aga = int(_dgc) << 1
	}
	if _fce > -2 && _fce < _efg.ReferenceBitmap.Height-1 {
		_dgc, _bga = _efg.ReferenceBitmap.GetByte(_gbg + _ede + _dgf)
		if _bga != nil {
			return _bga
		}
		_dbf = int(_dgc)
	}
	_aag = ((_fcea >> 5) & 0x6) | ((_dbf >> 2) & 0x30) | (_aga & 0x180) | (_ceg & 0xc00)
	var _bcc int
	for _cag := 0; _cag < _bd; _cag = _bcc {
		var _cdd int
		_bcc = _cag + 8
		var _ae int
		if _ae = _dec - _cag; _ae > 8 {
			_ae = 8
		}
		_faf := _bcc < _dec
		_faea := _bcc < _efg.ReferenceBitmap.Width
		_eag := _dgf + 1
		if _bce > 0 {
			_dgc = 0
			if _faf {
				_dgc, _bga = _efg.RegionBitmap.GetByte(_ddc - _gg + 1)
				if _bga != nil {
					return _bga
				}
			}
			_fcea = (_fcea << 8) | int(_dgc)
		}
		if _fce > 0 && _fce <= _efg.ReferenceBitmap.Height {
			var _bba int
			if _faea {
				_dgc, _bga = _efg.ReferenceBitmap.GetByte(_gbg - _ede + _eag)
				if _bga != nil {
					return _bga
				}
				_bba = int(_dgc) << 4
			}
			_ceg = (_ceg << 8) | _bba
		}
		if _fce >= 0 && _fce < _efg.ReferenceBitmap.Height {
			var _ade int
			if _faea {
				_dgc, _bga = _efg.ReferenceBitmap.GetByte(_gbg + _eag)
				if _bga != nil {
					return _bga
				}
				_ade = int(_dgc) << 1
			}
			_aga = (_aga << 8) | _ade
		}
		if _fce > -2 && _fce < (_efg.ReferenceBitmap.Height-1) {
			_dgc = 0
			if _faea {
				_dgc, _bga = _efg.ReferenceBitmap.GetByte(_gbg + _ede + _eag)
				if _bga != nil {
					return _bga
				}
			}
			_dbf = (_dbf << 8) | int(_dgc)
		}
		for _ccf := 0; _ccf < _ae; _ccf++ {
			var _edc int
			_ace := false
			_dcg := (_aag >> 4) & 0x1ff
			if _dcg == 0x1ff {
				_ace = true
				_edc = 1
			} else if _dcg == 0x00 {
				_ace = true
			}
			if !_ace {
				if _efg._fb {
					_bad = _efg.overrideAtTemplate0(_aag, _cag+_ccf, _bce, _cdd, _ccf)
					_efg._cd.SetIndex(int32(_bad))
				} else {
					_efg._cd.SetIndex(int32(_aag))
				}
				_edc, _bga = _efg._dc.DecodeBit(_efg._cd)
				if _bga != nil {
					return _bga
				}
			}
			_acc := uint(7 - _ccf)
			_cdd |= _edc << _acc
			_aag = ((_aag & 0xdb6) << 1) | _edc | (_fcea>>_acc+5)&0x002 | ((_dbf>>_acc + 2) & 0x010) | ((_aga >> _acc) & 0x080) | ((_ceg >> _acc) & 0x400)
		}
		_bga = _efg.RegionBitmap.SetByte(_ddc, byte(_cdd))
		if _bga != nil {
			return _bga
		}
		_ddc++
		_gbg++
	}
	return nil
}
func (_agcc *TextRegion) checkInput() error {
	const _gaee = "\u0063\u0068\u0065\u0063\u006b\u0049\u006e\u0070\u0075\u0074"
	if !_agcc.UseRefinement {
		if _agcc.SbrTemplate != 0 {
			_aae.Log.Debug("\u0053\u0062\u0072Te\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030")
			_agcc.SbrTemplate = 0
		}
	}
	if _agcc.SbHuffFS == 2 || _agcc.SbHuffRDWidth == 2 || _agcc.SbHuffRDHeight == 2 || _agcc.SbHuffRDX == 2 || _agcc.SbHuffRDY == 2 {
		return _edb.Error(_gaee, "h\u0075\u0066\u0066\u006d\u0061\u006e \u0066\u006c\u0061\u0067\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064")
	}
	if !_agcc.UseRefinement {
		if _agcc.SbHuffRSize != 0 {
			_aae.Log.Debug("\u0053\u0062\u0048uf\u0066\u0052\u0053\u0069\u007a\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030")
			_agcc.SbHuffRSize = 0
		}
		if _agcc.SbHuffRDY != 0 {
			_aae.Log.Debug("S\u0062\u0048\u0075\u0066fR\u0044Y\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0030")
			_agcc.SbHuffRDY = 0
		}
		if _agcc.SbHuffRDX != 0 {
			_aae.Log.Debug("S\u0062\u0048\u0075\u0066fR\u0044X\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0030")
			_agcc.SbHuffRDX = 0
		}
		if _agcc.SbHuffRDWidth != 0 {
			_aae.Log.Debug("\u0053b\u0048\u0075\u0066\u0066R\u0044\u0057\u0069\u0064\u0074h\u0020s\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u00200")
			_agcc.SbHuffRDWidth = 0
		}
		if _agcc.SbHuffRDHeight != 0 {
			_aae.Log.Debug("\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0048\u0065\u0069g\u0068\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020b\u0065\u0020\u0030")
			_agcc.SbHuffRDHeight = 0
		}
	}
	return nil
}
func (_abfa *HalftoneRegion) renderPattern(_dcb [][]int) (_agbda error) {
	var _afge, _bccd int
	for _gagbb := 0; _gagbb < int(_abfa.HGridHeight); _gagbb++ {
		for _gee := 0; _gee < int(_abfa.HGridWidth); _gee++ {
			_afge = _abfa.computeX(_gagbb, _gee)
			_bccd = _abfa.computeY(_gagbb, _gee)
			_fgdb := _abfa.Patterns[_dcb[_gagbb][_gee]]
			if _agbda = _gf.Blit(_fgdb, _abfa.HalftoneRegionBitmap, _afge+int(_abfa.HGridX), _bccd+int(_abfa.HGridY), _abfa.CombinationOperator); _agbda != nil {
				return _agbda
			}
		}
	}
	return nil
}

const (
	ORandom OrganizationType = iota
	OSequential
)

func (_fedf *GenericRefinementRegion) decodeTemplate(_cca, _bfc, _ccg, _gbgg, _dfg, _fgg, _fbc, _dcd, _afg, _ecb int, _faec templater) (_bac error) {
	var (
		_bfa, _feg, _dce, _cec, _fafd int16
		_daf, _bbd, _gged, _dbb       int
		_ega                          byte
	)
	if _afg >= 1 && (_afg-1) < _fedf.ReferenceBitmap.Height {
		_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb - _gbgg)
		if _bac != nil {
			return
		}
		_daf = int(_ega)
	}
	if _afg >= 0 && (_afg) < _fedf.ReferenceBitmap.Height {
		_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb)
		if _bac != nil {
			return
		}
		_bbd = int(_ega)
	}
	if _afg >= -1 && (_afg+1) < _fedf.ReferenceBitmap.Height {
		_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb + _gbgg)
		if _bac != nil {
			return
		}
		_gged = int(_ega)
	}
	_ecb++
	if _cca >= 1 {
		_ega, _bac = _fedf.RegionBitmap.GetByte(_dcd - _ccg)
		if _bac != nil {
			return
		}
		_dbb = int(_ega)
	}
	_dcd++
	_gda := _fedf.ReferenceDX % 8
	_fafb := 6 + _gda
	_gfd := _ecb % _gbgg
	if _fafb >= 0 {
		if _fafb < 8 {
			_bfa = int16(_daf>>uint(_fafb)) & 0x07
		}
		if _fafb < 8 {
			_feg = int16(_bbd>>uint(_fafb)) & 0x07
		}
		if _fafb < 8 {
			_dce = int16(_gged>>uint(_fafb)) & 0x07
		}
		if _fafb == 6 && _gfd > 1 {
			if _afg >= 1 && (_afg-1) < _fedf.ReferenceBitmap.Height {
				_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb - _gbgg - 2)
				if _bac != nil {
					return _bac
				}
				_bfa |= int16(_ega<<2) & 0x04
			}
			if _afg >= 0 && _afg < _fedf.ReferenceBitmap.Height {
				_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb - 2)
				if _bac != nil {
					return _bac
				}
				_feg |= int16(_ega<<2) & 0x04
			}
			if _afg >= -1 && _afg+1 < _fedf.ReferenceBitmap.Height {
				_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb + _gbgg - 2)
				if _bac != nil {
					return _bac
				}
				_dce |= int16(_ega<<2) & 0x04
			}
		}
		if _fafb == 0 {
			_daf = 0
			_bbd = 0
			_gged = 0
			if _gfd < _gbgg-1 {
				if _afg >= 1 && _afg-1 < _fedf.ReferenceBitmap.Height {
					_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb - _gbgg)
					if _bac != nil {
						return _bac
					}
					_daf = int(_ega)
				}
				if _afg >= 0 && _afg < _fedf.ReferenceBitmap.Height {
					_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb)
					if _bac != nil {
						return _bac
					}
					_bbd = int(_ega)
				}
				if _afg >= -1 && _afg+1 < _fedf.ReferenceBitmap.Height {
					_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb + _gbgg)
					if _bac != nil {
						return _bac
					}
					_gged = int(_ega)
				}
			}
			_ecb++
		}
	} else {
		_bfa = int16(_daf<<1) & 0x07
		_feg = int16(_bbd<<1) & 0x07
		_dce = int16(_gged<<1) & 0x07
		_daf = 0
		_bbd = 0
		_gged = 0
		if _gfd < _gbgg-1 {
			if _afg >= 1 && _afg-1 < _fedf.ReferenceBitmap.Height {
				_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb - _gbgg)
				if _bac != nil {
					return _bac
				}
				_daf = int(_ega)
			}
			if _afg >= 0 && _afg < _fedf.ReferenceBitmap.Height {
				_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb)
				if _bac != nil {
					return _bac
				}
				_bbd = int(_ega)
			}
			if _afg >= -1 && _afg+1 < _fedf.ReferenceBitmap.Height {
				_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb + _gbgg)
				if _bac != nil {
					return _bac
				}
				_gged = int(_ega)
			}
			_ecb++
		}
		_bfa |= int16((_daf >> 7) & 0x07)
		_feg |= int16((_bbd >> 7) & 0x07)
		_dce |= int16((_gged >> 7) & 0x07)
	}
	_cec = int16(_dbb >> 6)
	_fafd = 0
	_ddf := (2 - _gda) % 8
	_daf <<= uint(_ddf)
	_bbd <<= uint(_ddf)
	_gged <<= uint(_ddf)
	_dbb <<= 2
	var _dfgg int
	for _cfde := 0; _cfde < _bfc; _cfde++ {
		_cac := _cfde & 0x07
		_bdb := _faec.form(_bfa, _feg, _dce, _cec, _fafd)
		if _fedf._fb {
			_ega, _bac = _fedf.RegionBitmap.GetByte(_fedf.RegionBitmap.GetByteIndex(_cfde, _cca))
			if _bac != nil {
				return _bac
			}
			_fedf._cd.SetIndex(int32(_fedf.overrideAtTemplate0(int(_bdb), _cfde, _cca, int(_ega), _cac)))
		} else {
			_fedf._cd.SetIndex(int32(_bdb))
		}
		_dfgg, _bac = _fedf._dc.DecodeBit(_fedf._cd)
		if _bac != nil {
			return _bac
		}
		if _bac = _fedf.RegionBitmap.SetPixel(_cfde, _cca, byte(_dfgg)); _bac != nil {
			return _bac
		}
		_bfa = ((_bfa << 1) | 0x01&int16(_daf>>7)) & 0x07
		_feg = ((_feg << 1) | 0x01&int16(_bbd>>7)) & 0x07
		_dce = ((_dce << 1) | 0x01&int16(_gged>>7)) & 0x07
		_cec = ((_cec << 1) | 0x01&int16(_dbb>>7)) & 0x07
		_fafd = int16(_dfgg)
		if (_cfde-int(_fedf.ReferenceDX))%8 == 5 {
			_daf = 0
			_bbd = 0
			_gged = 0
			if ((_cfde-int(_fedf.ReferenceDX))/8)+1 < _fedf.ReferenceBitmap.RowStride {
				if _afg >= 1 && (_afg-1) < _fedf.ReferenceBitmap.Height {
					_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb - _gbgg)
					if _bac != nil {
						return _bac
					}
					_daf = int(_ega)
				}
				if _afg >= 0 && _afg < _fedf.ReferenceBitmap.Height {
					_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb)
					if _bac != nil {
						return _bac
					}
					_bbd = int(_ega)
				}
				if _afg >= -1 && (_afg+1) < _fedf.ReferenceBitmap.Height {
					_ega, _bac = _fedf.ReferenceBitmap.GetByte(_ecb + _gbgg)
					if _bac != nil {
						return _bac
					}
					_gged = int(_ega)
				}
			}
			_ecb++
		} else {
			_daf <<= 1
			_bbd <<= 1
			_gged <<= 1
		}
		if _cac == 5 && _cca >= 1 {
			if ((_cfde >> 3) + 1) >= _fedf.RegionBitmap.RowStride {
				_dbb = 0
			} else {
				_ega, _bac = _fedf.RegionBitmap.GetByte(_dcd - _ccg)
				if _bac != nil {
					return _bac
				}
				_dbb = int(_ega)
			}
			_dcd++
		} else {
			_dbb <<= 1
		}
	}
	return nil
}
func (_ddfb *SymbolDictionary) setExportedSymbols(_daag []int) {
	for _begd := uint32(0); _begd < _ddfb._gfg+_ddfb.NumberOfNewSymbols; _begd++ {
		if _daag[_begd] == 1 {
			var _fbaac *_gf.Bitmap
			if _begd < _ddfb._gfg {
				_fbaac = _ddfb._bgbd[_begd]
			} else {
				_fbaac = _ddfb._gedb[_begd-_ddfb._gfg]
			}
			_aae.Log.Trace("\u005bS\u0059\u004dB\u004f\u004c\u002d\u0044I\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u0020\u0041\u0064\u0064 E\u0078\u0070\u006fr\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c:\u0020\u0027%\u0073\u0027", _fbaac)
			_ddfb._aebc = append(_ddfb._aebc, _fbaac)
		}
	}
}
func (_fbbcf *TextRegion) blit(_addd *_gf.Bitmap, _dcbe int64) error {
	if _fbbcf.IsTransposed == 0 && (_fbbcf.ReferenceCorner == 2 || _fbbcf.ReferenceCorner == 3) {
		_fbbcf._abca += int64(_addd.Width - 1)
	} else if _fbbcf.IsTransposed == 1 && (_fbbcf.ReferenceCorner == 0 || _fbbcf.ReferenceCorner == 2) {
		_fbbcf._abca += int64(_addd.Height - 1)
	}
	_cggc := _fbbcf._abca
	if _fbbcf.IsTransposed == 1 {
		_cggc, _dcbe = _dcbe, _cggc
	}
	switch _fbbcf.ReferenceCorner {
	case 0:
		_dcbe -= int64(_addd.Height - 1)
	case 2:
		_dcbe -= int64(_addd.Height - 1)
		_cggc -= int64(_addd.Width - 1)
	case 3:
		_cggc -= int64(_addd.Width - 1)
	}
	_eegea := _gf.Blit(_addd, _fbbcf.RegionBitmap, int(_cggc), int(_dcbe), _fbbcf.CombinationOperator)
	if _eegea != nil {
		return _eegea
	}
	if _fbbcf.IsTransposed == 0 && (_fbbcf.ReferenceCorner == 0 || _fbbcf.ReferenceCorner == 1) {
		_fbbcf._abca += int64(_addd.Width - 1)
	}
	if _fbbcf.IsTransposed == 1 && (_fbbcf.ReferenceCorner == 1 || _fbbcf.ReferenceCorner == 3) {
		_fbbcf._abca += int64(_addd.Height - 1)
	}
	return nil
}
func (_ffa *GenericRegion) getPixel(_bceb, _decc int) int8 {
	if _bceb < 0 || _bceb >= _ffa.Bitmap.Width {
		return 0
	}
	if _decc < 0 || _decc >= _ffa.Bitmap.Height {
		return 0
	}
	if _ffa.Bitmap.GetPixel(_bceb, _decc) {
		return 1
	}
	return 0
}
func (_cda *PageInformationSegment) String() string {
	_beced := &_d.Builder{}
	_beced.WriteString("\u000a\u005b\u0050\u0041G\u0045\u002d\u0049\u004e\u0046\u004f\u0052\u004d\u0041\u0054I\u004fN\u002d\u0053\u0045\u0047\u004d\u0045\u004eT\u005d\u000a")
	_beced.WriteString(_af.Sprintf("\u0009\u002d \u0042\u004d\u0048e\u0069\u0067\u0068\u0074\u003a\u0020\u0025\u0064\u000a", _cda.PageBMHeight))
	_beced.WriteString(_af.Sprintf("\u0009-\u0020B\u004d\u0057\u0069\u0064\u0074\u0068\u003a\u0020\u0025\u0064\u000a", _cda.PageBMWidth))
	_beced.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0052es\u006f\u006c\u0075\u0074\u0069\u006f\u006e\u0058\u003a\u0020\u0025\u0064\u000a", _cda.ResolutionX))
	_beced.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0052es\u006f\u006c\u0075\u0074\u0069\u006f\u006e\u0059\u003a\u0020\u0025\u0064\u000a", _cda.ResolutionY))
	_beced.WriteString(_af.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065r\u0061\u0074\u006fr\u003a \u0025\u0073\u000a", _cda._dbcd))
	_beced.WriteString(_af.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006eO\u0070\u0065\u0072\u0061\u0074\u006f\u0072O\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u003a\u0020\u0025v\u000a", _cda._fbbge))
	_beced.WriteString(_af.Sprintf("\u0009-\u0020I\u0073\u004c\u006f\u0073\u0073l\u0065\u0073s\u003a\u0020\u0025\u0076\u000a", _cda.IsLossless))
	_beced.WriteString(_af.Sprintf("\u0009\u002d\u0020R\u0065\u0071\u0075\u0069r\u0065\u0073\u0041\u0075\u0078\u0069\u006ci\u0061\u0072\u0079\u0042\u0075\u0066\u0066\u0065\u0072\u003a\u0020\u0025\u0076\u000a", _cda._cbdc))
	_beced.WriteString(_af.Sprintf("\u0009\u002d\u0020M\u0069\u0067\u0068\u0074C\u006f\u006e\u0074\u0061\u0069\u006e\u0052e\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0076\u000a", _cda._ccdd))
	_beced.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0049\u0073\u0053\u0074\u0072\u0069\u0070\u0065\u0064:\u0020\u0025\u0076\u000a", _cda.IsStripe))
	_beced.WriteString(_af.Sprintf("\t\u002d\u0020\u004d\u0061xS\u0074r\u0069\u0070\u0065\u0053\u0069z\u0065\u003a\u0020\u0025\u0076\u000a", _cda.MaxStripeSize))
	return _beced.String()
}
func (_dcbb *PatternDictionary) setGbAtPixels() {
	if _dcbb.HDTemplate == 0 {
		_dcbb.GBAtX = make([]int8, 4)
		_dcbb.GBAtY = make([]int8, 4)
		_dcbb.GBAtX[0] = -int8(_dcbb.HdpWidth)
		_dcbb.GBAtY[0] = 0
		_dcbb.GBAtX[1] = -3
		_dcbb.GBAtY[1] = -1
		_dcbb.GBAtX[2] = 2
		_dcbb.GBAtY[2] = -2
		_dcbb.GBAtX[3] = -2
		_dcbb.GBAtY[3] = -2
	} else {
		_dcbb.GBAtX = []int8{-int8(_dcbb.HdpWidth)}
		_dcbb.GBAtY = []int8{0}
	}
}
func (_gbe *GenericRefinementRegion) updateOverride() error {
	if _gbe.GrAtX == nil || _gbe.GrAtY == nil {
		return _b.New("\u0041\u0054\u0020\u0070\u0069\u0078\u0065\u006c\u0073\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
	}
	if len(_gbe.GrAtX) != len(_gbe.GrAtY) {
		return _b.New("A\u0054\u0020\u0070\u0069xe\u006c \u0069\u006e\u0063\u006f\u006es\u0069\u0073\u0074\u0065\u006e\u0074")
	}
	_gbe._ce = make([]bool, len(_gbe.GrAtX))
	switch _gbe.TemplateID {
	case 0:
		if _gbe.GrAtX[0] != -1 && _gbe.GrAtY[0] != -1 {
			_gbe._ce[0] = true
			_gbe._fb = true
		}
		if _gbe.GrAtX[1] != -1 && _gbe.GrAtY[1] != -1 {
			_gbe._ce[1] = true
			_gbe._fb = true
		}
	case 1:
		_gbe._fb = false
	}
	return nil
}
func (_da *EndOfStripe) parseHeader() error {
	_cfd, _dg := _da._gb.ReadBits(32)
	if _dg != nil {
		return _dg
	}
	_da._bc = int(_cfd & _e.MaxInt32)
	return nil
}

type EndOfStripe struct {
	_gb *_g.Reader
	_bc int
}

func (_aedb *PageInformationSegment) parseHeader() (_fegge error) {
	_aae.Log.Trace("\u005b\u0050\u0061\u0067\u0065I\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065\u0067m\u0065\u006e\u0074\u005d\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065\u0061\u0064\u0065\u0072\u002e\u002e\u002e")
	defer func() {
		var _fbdc = "[\u0050\u0061\u0067\u0065\u0049\u006e\u0066\u006f\u0072m\u0061\u0074\u0069\u006f\u006e\u0053\u0065gm\u0065\u006e\u0074\u005d \u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065ad\u0065\u0072 \u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064"
		if _fegge != nil {
			_fbdc += "\u0020\u0077\u0069t\u0068\u0020\u0065\u0072\u0072\u006f\u0072\u0020" + _fegge.Error()
		} else {
			_fbdc += "\u0020\u0073\u0075\u0063\u0063\u0065\u0073\u0073\u0066\u0075\u006c\u006c\u0079"
		}
		_aae.Log.Trace(_fbdc)
	}()
	if _fegge = _aedb.readWidthAndHeight(); _fegge != nil {
		return _fegge
	}
	if _fegge = _aedb.readResolution(); _fegge != nil {
		return _fegge
	}
	_, _fegge = _aedb._faef.ReadBit()
	if _fegge != nil {
		return _fegge
	}
	if _fegge = _aedb.readCombinationOperatorOverrideAllowed(); _fegge != nil {
		return _fegge
	}
	if _fegge = _aedb.readRequiresAuxiliaryBuffer(); _fegge != nil {
		return _fegge
	}
	if _fegge = _aedb.readCombinationOperator(); _fegge != nil {
		return _fegge
	}
	if _fegge = _aedb.readDefaultPixelValue(); _fegge != nil {
		return _fegge
	}
	if _fegge = _aedb.readContainsRefinement(); _fegge != nil {
		return _fegge
	}
	if _fegge = _aedb.readIsLossless(); _fegge != nil {
		return _fegge
	}
	if _fegge = _aedb.readIsStriped(); _fegge != nil {
		return _fegge
	}
	if _fegge = _aedb.readMaxStripeSize(); _fegge != nil {
		return _fegge
	}
	if _fegge = _aedb.checkInput(); _fegge != nil {
		return _fegge
	}
	_aae.Log.Trace("\u0025\u0073", _aedb)
	return nil
}
func (_dbbdg *SymbolDictionary) setSymbolsArray() error {
	if _dbbdg._bgbd == nil {
		if _afdb := _dbbdg.retrieveImportSymbols(); _afdb != nil {
			return _afdb
		}
	}
	if _dbbdg._ebfb == nil {
		_dbbdg._ebfb = append(_dbbdg._ebfb, _dbbdg._bgbd...)
	}
	return nil
}
func (_fef *SymbolDictionary) readNumberOfNewSymbols() error {
	_gfefc, _cdad := _fef._caba.ReadBits(32)
	if _cdad != nil {
		return _cdad
	}
	_fef.NumberOfNewSymbols = uint32(_gfefc & _e.MaxUint32)
	return nil
}
func (_bda *SymbolDictionary) parseHeader() (_ebeb error) {
	_aae.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c \u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		if _ebeb != nil {
			_aae.Log.Trace("\u005bS\u0059\u004dB\u004f\u004c\u0020\u0044I\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u005b\u0050\u0041\u0052SE\u002d\u0048\u0045A\u0044\u0045R\u005d\u0020\u0066\u0061\u0069\u006ce\u0064\u002e \u0025\u0076", _ebeb)
		} else {
			_aae.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c \u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		}
	}()
	if _ebeb = _bda.readRegionFlags(); _ebeb != nil {
		return _ebeb
	}
	if _ebeb = _bda.setAtPixels(); _ebeb != nil {
		return _ebeb
	}
	if _ebeb = _bda.setRefinementAtPixels(); _ebeb != nil {
		return _ebeb
	}
	if _ebeb = _bda.readNumberOfExportedSymbols(); _ebeb != nil {
		return _ebeb
	}
	if _ebeb = _bda.readNumberOfNewSymbols(); _ebeb != nil {
		return _ebeb
	}
	if _ebeb = _bda.setInSyms(); _ebeb != nil {
		return _ebeb
	}
	if _bda._bddf {
		_eafg := _bda.Header.RTSegments
		for _cecd := len(_eafg) - 1; _cecd >= 0; _cecd-- {
			if _eafg[_cecd].Type == 0 {
				_deba, _ccae := _eafg[_cecd].SegmentData.(*SymbolDictionary)
				if !_ccae {
					_ebeb = _af.Errorf("\u0072\u0065\u006c\u0061\u0074\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074:\u0020\u0025\u0076\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020S\u0079\u006d\u0062\u006f\u006c\u0020\u0044\u0069\u0063\u0074\u0069\u006fna\u0072\u0079\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074", _eafg[_cecd])
					return _ebeb
				}
				if _deba._bddf {
					_bda.setRetainedCodingContexts(_deba)
				}
				break
			}
		}
	}
	if _ebeb = _bda.checkInput(); _ebeb != nil {
		return _ebeb
	}
	return nil
}
func (_eac *GenericRefinementRegion) overrideAtTemplate0(_cde, _abc, _fbb, _feaab, _aagc int) int {
	if _eac._ce[0] {
		_cde &= 0xfff7
		if _eac.GrAtY[0] == 0 && int(_eac.GrAtX[0]) >= -_aagc {
			_cde |= (_feaab >> uint(7-(_aagc+int(_eac.GrAtX[0]))) & 0x1) << 3
		} else {
			_cde |= _eac.getPixel(_eac.RegionBitmap, _abc+int(_eac.GrAtX[0]), _fbb+int(_eac.GrAtY[0])) << 3
		}
	}
	if _eac._ce[1] {
		_cde &= 0xefff
		if _eac.GrAtY[1] == 0 && int(_eac.GrAtX[1]) >= -_aagc {
			_cde |= (_feaab >> uint(7-(_aagc+int(_eac.GrAtX[1]))) & 0x1) << 12
		} else {
			_cde |= _eac.getPixel(_eac.ReferenceBitmap, _abc+int(_eac.GrAtX[1]), _fbb+int(_eac.GrAtY[1]))
		}
	}
	return _cde
}

var _ SegmentEncoder = &RegionSegment{}

func (_abg *GenericRefinementRegion) readAtPixels() error {
	_abg.GrAtX = make([]int8, 2)
	_abg.GrAtY = make([]int8, 2)
	_cage, _bec := _abg._dd.ReadByte()
	if _bec != nil {
		return _bec
	}
	_abg.GrAtX[0] = int8(_cage)
	_cage, _bec = _abg._dd.ReadByte()
	if _bec != nil {
		return _bec
	}
	_abg.GrAtY[0] = int8(_cage)
	_cage, _bec = _abg._dd.ReadByte()
	if _bec != nil {
		return _bec
	}
	_abg.GrAtX[1] = int8(_cage)
	_cage, _bec = _abg._dd.ReadByte()
	if _bec != nil {
		return _bec
	}
	_abg.GrAtY[1] = int8(_cage)
	return nil
}
func (_abafe *SymbolDictionary) decodeThroughTextRegion(_cfba, _cceba, _beda uint32) error {
	if _abafe._cged == nil {
		_abafe._cged = _adda(_abafe._caba, nil)
		_abafe._cged.setContexts(_abafe._afbd, _df.NewStats(512, 1), _df.NewStats(512, 1), _df.NewStats(512, 1), _df.NewStats(512, 1), _abafe._ggc, _df.NewStats(512, 1), _df.NewStats(512, 1), _df.NewStats(512, 1), _df.NewStats(512, 1))
	}
	if _afda := _abafe.setSymbolsArray(); _afda != nil {
		return _afda
	}
	_abafe._cged.setParameters(_abafe._dgcb, _abafe.IsHuffmanEncoded, true, _cfba, _cceba, _beda, 1, _abafe._gfg+_abafe._acga, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, _abafe.SdrTemplate, _abafe.SdrATX, _abafe.SdrATY, _abafe._ebfb, _abafe._abcgg)
	return _abafe.addSymbol(_abafe._cged)
}
func (_ecfc *GenericRegion) decodeTemplate3(_bag, _debe, _gfb int, _fegg, _beg int) (_acg error) {
	const _fcba = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0033"
	var (
		_dfc, _dabb int
		_cdgf       int
		_gefg       byte
		_ga, _gfdc  int
	)
	if _bag >= 1 {
		_gefg, _acg = _ecfc.Bitmap.GetByte(_beg)
		if _acg != nil {
			return _edb.Wrap(_acg, _fcba, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_cdgf = int(_gefg)
	}
	_dfc = (_cdgf >> 1) & 0x70
	for _edd := 0; _edd < _gfb; _edd = _ga {
		var (
			_fafee byte
			_dfaa  int
		)
		_ga = _edd + 8
		if _dfae := _debe - _edd; _dfae > 8 {
			_dfaa = 8
		} else {
			_dfaa = _dfae
		}
		if _bag >= 1 {
			_cdgf <<= 8
			if _ga < _debe {
				_gefg, _acg = _ecfc.Bitmap.GetByte(_beg + 1)
				if _acg != nil {
					return _edb.Wrap(_acg, _fcba, "\u0069\u006e\u006e\u0065\u0072\u0020\u002d\u0020\u006c\u0069\u006e\u0065 \u003e\u003d\u0020\u0031")
				}
				_cdgf |= int(_gefg)
			}
		}
		for _gfa := 0; _gfa < _dfaa; _gfa++ {
			if _ecfc._cab {
				_dabb = _ecfc.overrideAtTemplate3(_dfc, _edd+_gfa, _bag, int(_fafee), _gfa)
				_ecfc._fac.SetIndex(int32(_dabb))
			} else {
				_ecfc._fac.SetIndex(int32(_dfc))
			}
			_gfdc, _acg = _ecfc._ddga.DecodeBit(_ecfc._fac)
			if _acg != nil {
				return _edb.Wrap(_acg, _fcba, "")
			}
			_fafee |= byte(_gfdc) << byte(7-_gfa)
			_dfc = ((_dfc & 0x1f7) << 1) | _gfdc | ((_cdgf >> uint(8-_gfa)) & 0x010)
		}
		if _egag := _ecfc.Bitmap.SetByte(_fegg, _fafee); _egag != nil {
			return _edb.Wrap(_egag, _fcba, "")
		}
		_fegg++
		_beg++
	}
	return nil
}

var _ templater = &template0{}

func (_adec *GenericRefinementRegion) String() string {
	_ced := &_d.Builder{}
	_ced.WriteString("\u000a[\u0047E\u004e\u0045\u0052\u0049\u0043 \u0052\u0045G\u0049\u004f\u004e\u005d\u000a")
	_ced.WriteString(_adec.RegionInfo.String() + "\u000a")
	_ced.WriteString(_af.Sprintf("\u0009\u002d \u0049\u0073\u0054P\u0047\u0052\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _adec.IsTPGROn))
	_ced.WriteString(_af.Sprintf("\u0009-\u0020T\u0065\u006d\u0070\u006c\u0061t\u0065\u0049D\u003a\u0020\u0025\u0076\u000a", _adec.TemplateID))
	_ced.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0047\u0072\u0041\u0074\u0058\u003a\u0020\u0025\u0076\u000a", _adec.GrAtX))
	_ced.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0047\u0072\u0041\u0074\u0059\u003a\u0020\u0025\u0076\u000a", _adec.GrAtY))
	_ced.WriteString(_af.Sprintf("\u0009-\u0020R\u0065\u0066\u0065\u0072\u0065n\u0063\u0065D\u0058\u0020\u0025\u0076\u000a", _adec.ReferenceDX))
	_ced.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0052ef\u0065\u0072\u0065\u006e\u0063\u0044\u0065\u0059\u003a\u0020\u0025\u0076\u000a", _adec.ReferenceDY))
	return _ced.String()
}

type RegionSegment struct {
	_ebecg             *_g.Reader
	BitmapWidth        uint32
	BitmapHeight       uint32
	XLocation          uint32
	YLocation          uint32
	CombinaionOperator _gf.CombinationOperator
}
type SymbolDictionary struct {
	_caba                       *_g.Reader
	SdrTemplate                 int8
	SdTemplate                  int8
	_dedee                      bool
	_bddf                       bool
	SdHuffAggInstanceSelection  bool
	SdHuffBMSizeSelection       int8
	SdHuffDecodeWidthSelection  int8
	SdHuffDecodeHeightSelection int8
	UseRefinementAggregation    bool
	IsHuffmanEncoded            bool
	SdATX                       []int8
	SdATY                       []int8
	SdrATX                      []int8
	SdrATY                      []int8
	NumberOfExportedSymbols     uint32
	NumberOfNewSymbols          uint32
	Header                      *Header
	_gfg                        uint32
	_bgbd                       []*_gf.Bitmap
	_acga                       uint32
	_gedb                       []*_gf.Bitmap
	_cfgd                       _fa.Tabler
	_accg                       _fa.Tabler
	_ccbc                       _fa.Tabler
	_ecef                       _fa.Tabler
	_aebc                       []*_gf.Bitmap
	_ebfb                       []*_gf.Bitmap
	_dgcb                       *_df.Decoder
	_cged                       *TextRegion
	_ddcgf                      *GenericRegion
	_ebca                       *GenericRefinementRegion
	_afbd                       *_df.DecoderStats
	_ggae                       *_df.DecoderStats
	_bcbb                       *_df.DecoderStats
	_aagg                       *_df.DecoderStats
	_fdfd                       *_df.DecoderStats
	_fagc                       *_df.DecoderStats
	_fccf                       *_df.DecoderStats
	_adba                       *_df.DecoderStats
	_ggc                        *_df.DecoderStats
	_abcgg                      int8
	_cfcf                       *_gf.Bitmaps
	_bcgg                       []int
	_efgef                      map[int]int
	_gdc                        bool
}
type HalftoneRegion struct {
	_fagd                *_g.Reader
	_ceag                *Header
	DataHeaderOffset     int64
	DataHeaderLength     int64
	DataOffset           int64
	DataLength           int64
	RegionSegment        *RegionSegment
	HDefaultPixel        int8
	CombinationOperator  _gf.CombinationOperator
	HSkipEnabled         bool
	HTemplate            byte
	IsMMREncoded         bool
	HGridWidth           uint32
	HGridHeight          uint32
	HGridX               int32
	HGridY               int32
	HRegionX             uint16
	HRegionY             uint16
	HalftoneRegionBitmap *_gf.Bitmap
	Patterns             []*_gf.Bitmap
}

func (_bead *Header) writeFlags(_aadd _g.BinaryWriter) (_afbc error) {
	const _fgf = "\u0048\u0065\u0061\u0064\u0065\u0072\u002e\u0077\u0072\u0069\u0074\u0065F\u006c\u0061\u0067\u0073"
	_aada := byte(_bead.Type)
	if _afbc = _aadd.WriteByte(_aada); _afbc != nil {
		return _edb.Wrap(_afbc, _fgf, "\u0077\u0072\u0069ti\u006e\u0067\u0020\u0073\u0065\u0067\u006d\u0065\u006et\u0020t\u0079p\u0065 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	if !_bead.RetainFlag && !_bead.PageAssociationFieldSize {
		return nil
	}
	if _afbc = _aadd.SkipBits(-8); _afbc != nil {
		return _edb.Wrap(_afbc, _fgf, "\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0062\u0061\u0063\u006b\u0020\u0074\u0068\u0065\u0020\u0062\u0069\u0074\u0073\u0020\u0066\u0061il\u0065\u0064")
	}
	var _dad int
	if _bead.RetainFlag {
		_dad = 1
	}
	if _afbc = _aadd.WriteBit(_dad); _afbc != nil {
		return _edb.Wrap(_afbc, _fgf, "\u0072\u0065\u0074\u0061in\u0020\u0072\u0065\u0074\u0061\u0069\u006e\u0020\u0066\u006c\u0061\u0067\u0073")
	}
	_dad = 0
	if _bead.PageAssociationFieldSize {
		_dad = 1
	}
	if _afbc = _aadd.WriteBit(_dad); _afbc != nil {
		return _edb.Wrap(_afbc, _fgf, "p\u0061\u0067\u0065\u0020as\u0073o\u0063\u0069\u0061\u0074\u0069o\u006e\u0020\u0066\u006c\u0061\u0067")
	}
	_aadd.FinishByte()
	return nil
}

type templater interface {
	form(_aba, _dfa, _agae, _fgb, _aefd int16) int16
	setIndex(_dgfc *_df.DecoderStats)
}

func (_abe *TextRegion) setContexts(_dcfc *_df.DecoderStats, _eeeg *_df.DecoderStats, _cfgg *_df.DecoderStats, _fagfg *_df.DecoderStats, _babbg *_df.DecoderStats, _aaedg *_df.DecoderStats, _gfbg *_df.DecoderStats, _ebgd *_df.DecoderStats, _aaaa *_df.DecoderStats, _ggeee *_df.DecoderStats) {
	_abe._gede = _eeeg
	_abe._cgdb = _cfgg
	_abe._dacg = _fagfg
	_abe._fff = _babbg
	_abe._gbcc = _gfbg
	_abe._acecd = _ebgd
	_abe._bbae = _aaedg
	_abe._ffab = _aaaa
	_abe._fbeg = _ggeee
	_abe._bada = _dcfc
}
func (_beag *GenericRegion) overrideAtTemplate3(_eage, _cae, _gac, _cbc, _cbfg int) int {
	_eage &= 0x3EF
	if _beag.GBAtY[0] == 0 && _beag.GBAtX[0] >= -int8(_cbfg) {
		_eage |= (_cbc >> uint(7-(int8(_cbfg)+_beag.GBAtX[0])) & 0x1) << 4
	} else {
		_eage |= int(_beag.getPixel(_cae+int(_beag.GBAtX[0]), _gac+int(_beag.GBAtY[0]))) << 4
	}
	return _eage
}
func (_gbcd *TextRegion) symbolIDCodeLengths() error {
	var (
		_fcfe []*_fa.Code
		_agdd uint64
		_cedf _fa.Tabler
		_dcea error
	)
	for _abec := 0; _abec < 35; _abec++ {
		_agdd, _dcea = _gbcd._ecc.ReadBits(4)
		if _dcea != nil {
			return _dcea
		}
		_gaab := int(_agdd & 0xf)
		if _gaab > 0 {
			_fcfe = append(_fcfe, _fa.NewCode(int32(_gaab), 0, int32(_abec), false))
		}
	}
	_cedf, _dcea = _fa.NewFixedSizeTable(_fcfe)
	if _dcea != nil {
		return _dcea
	}
	var (
		_daefg int64
		_dbbgb uint32
		_gaag  []*_fa.Code
		_beggf int64
	)
	for _dbbgb < _gbcd.NumberOfSymbols {
		_beggf, _dcea = _cedf.Decode(_gbcd._ecc)
		if _dcea != nil {
			return _dcea
		}
		if _beggf < 32 {
			if _beggf > 0 {
				_gaag = append(_gaag, _fa.NewCode(int32(_beggf), 0, int32(_dbbgb), false))
			}
			_daefg = _beggf
			_dbbgb++
		} else {
			var _cdefg, _ggcd int64
			switch _beggf {
			case 32:
				_agdd, _dcea = _gbcd._ecc.ReadBits(2)
				if _dcea != nil {
					return _dcea
				}
				_cdefg = 3 + int64(_agdd)
				if _dbbgb > 0 {
					_ggcd = _daefg
				}
			case 33:
				_agdd, _dcea = _gbcd._ecc.ReadBits(3)
				if _dcea != nil {
					return _dcea
				}
				_cdefg = 3 + int64(_agdd)
			case 34:
				_agdd, _dcea = _gbcd._ecc.ReadBits(7)
				if _dcea != nil {
					return _dcea
				}
				_cdefg = 11 + int64(_agdd)
			}
			for _dead := 0; _dead < int(_cdefg); _dead++ {
				if _ggcd > 0 {
					_gaag = append(_gaag, _fa.NewCode(int32(_ggcd), 0, int32(_dbbgb), false))
				}
				_dbbgb++
			}
		}
	}
	_gbcd._ecc.Align()
	_gbcd._dbcbe, _dcea = _fa.NewFixedSizeTable(_gaag)
	return _dcea
}
func (_feb *TextRegion) decodeRdw() (int64, error) {
	const _gdgcg = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064w"
	if _feb.IsHuffmanEncoded {
		if _feb.SbHuffRDWidth == 3 {
			if _feb._beed == nil {
				var (
					_bgbdd int
					_efeab error
				)
				if _feb.SbHuffFS == 3 {
					_bgbdd++
				}
				if _feb.SbHuffDS == 3 {
					_bgbdd++
				}
				if _feb.SbHuffDT == 3 {
					_bgbdd++
				}
				_feb._beed, _efeab = _feb.getUserTable(_bgbdd)
				if _efeab != nil {
					return 0, _edb.Wrap(_efeab, _gdgcg, "")
				}
			}
			return _feb._beed.Decode(_feb._ecc)
		}
		_efcb, _bgdf := _fa.GetStandardTable(14 + int(_feb.SbHuffRDWidth))
		if _bgdf != nil {
			return 0, _edb.Wrap(_bgdf, _gdgcg, "")
		}
		return _efcb.Decode(_feb._ecc)
	}
	_ebea, _ecde := _feb._egdc.DecodeInt(_feb._gbcc)
	if _ecde != nil {
		return 0, _edb.Wrap(_ecde, _gdgcg, "")
	}
	return int64(_ebea), nil
}
func (_bbe *Header) subInputReader() (*_g.Reader, error) {
	_gfab := int(_bbe.SegmentDataLength)
	if _bbe.SegmentDataLength == _e.MaxInt32 {
		_gfab = -1
	}
	return _bbe.Reader.NewPartialReader(int(_bbe.SegmentDataStartOffset), _gfab, false)
}
func (_gccb *SymbolDictionary) addSymbol(_eafc Regioner) error {
	_gbee, _cfge := _eafc.GetRegionBitmap()
	if _cfge != nil {
		return _cfge
	}
	_gccb._gedb[_gccb._acga] = _gbee
	_gccb._ebfb = append(_gccb._ebfb, _gbee)
	_aae.Log.Trace("\u005b\u0053YM\u0042\u004f\u004c \u0044\u0049\u0043\u0054ION\u0041RY\u005d\u0020\u0041\u0064\u0064\u0065\u0064 s\u0079\u006d\u0062\u006f\u006c\u003a\u0020%\u0073", _gbee)
	return nil
}
func (_dcdf *GenericRegion) decodeTemplate1(_bacc, _ebad, _cfb int, _dabe, _fecd int) (_dcga error) {
	const _eeg = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0031"
	var (
		_ggee, _fbf  int
		_gea, _gdbd  int
		_ccgc        byte
		_aggg, _abcf int
	)
	if _bacc >= 1 {
		_ccgc, _dcga = _dcdf.Bitmap.GetByte(_fecd)
		if _dcga != nil {
			return _edb.Wrap(_dcga, _eeg, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_gea = int(_ccgc)
	}
	if _bacc >= 2 {
		_ccgc, _dcga = _dcdf.Bitmap.GetByte(_fecd - _dcdf.Bitmap.RowStride)
		if _dcga != nil {
			return _edb.Wrap(_dcga, _eeg, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_gdbd = int(_ccgc) << 5
	}
	_ggee = ((_gea >> 1) & 0x1f8) | ((_gdbd >> 1) & 0x1e00)
	for _cbb := 0; _cbb < _cfb; _cbb = _aggg {
		var (
			_cddee byte
			_agge  int
		)
		_aggg = _cbb + 8
		if _geac := _ebad - _cbb; _geac > 8 {
			_agge = 8
		} else {
			_agge = _geac
		}
		if _bacc > 0 {
			_gea <<= 8
			if _aggg < _ebad {
				_ccgc, _dcga = _dcdf.Bitmap.GetByte(_fecd + 1)
				if _dcga != nil {
					return _edb.Wrap(_dcga, _eeg, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_gea |= int(_ccgc)
			}
		}
		if _bacc > 1 {
			_gdbd <<= 8
			if _aggg < _ebad {
				_ccgc, _dcga = _dcdf.Bitmap.GetByte(_fecd - _dcdf.Bitmap.RowStride + 1)
				if _dcga != nil {
					return _edb.Wrap(_dcga, _eeg, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_gdbd |= int(_ccgc) << 5
			}
		}
		for _bdee := 0; _bdee < _agge; _bdee++ {
			if _dcdf._cab {
				_fbf = _dcdf.overrideAtTemplate1(_ggee, _cbb+_bdee, _bacc, int(_cddee), _bdee)
				_dcdf._fac.SetIndex(int32(_fbf))
			} else {
				_dcdf._fac.SetIndex(int32(_ggee))
			}
			_abcf, _dcga = _dcdf._ddga.DecodeBit(_dcdf._fac)
			if _dcga != nil {
				return _edb.Wrap(_dcga, _eeg, "")
			}
			_cddee |= byte(_abcf) << uint(7-_bdee)
			_ccgcg := uint(8 - _bdee)
			_ggee = ((_ggee & 0xefb) << 1) | _abcf | ((_gea >> _ccgcg) & 0x8) | ((_gdbd >> _ccgcg) & 0x200)
		}
		if _acba := _dcdf.Bitmap.SetByte(_dabe, _cddee); _acba != nil {
			return _edb.Wrap(_acba, _eeg, "")
		}
		_dabe++
		_fecd++
	}
	return nil
}

type GenericRegion struct {
	_efeg            *_g.Reader
	DataHeaderOffset int64
	DataHeaderLength int64
	DataOffset       int64
	DataLength       int64
	RegionSegment    *RegionSegment
	UseExtTemplates  bool
	IsTPGDon         bool
	GBTemplate       byte
	IsMMREncoded     bool
	UseMMR           bool
	GBAtX            []int8
	GBAtY            []int8
	GBAtOverride     []bool
	_cab             bool
	Bitmap           *_gf.Bitmap
	_ddga            *_df.Decoder
	_fac             *_df.DecoderStats
	_gdf             *_eb.Decoder
}

func (_abbb *TextRegion) decodeRdh() (int64, error) {
	const _bdae = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064h"
	if _abbb.IsHuffmanEncoded {
		if _abbb.SbHuffRDHeight == 3 {
			if _abbb._cceff == nil {
				var (
					_dbdfb int
					_gefa  error
				)
				if _abbb.SbHuffFS == 3 {
					_dbdfb++
				}
				if _abbb.SbHuffDS == 3 {
					_dbdfb++
				}
				if _abbb.SbHuffDT == 3 {
					_dbdfb++
				}
				if _abbb.SbHuffRDWidth == 3 {
					_dbdfb++
				}
				_abbb._cceff, _gefa = _abbb.getUserTable(_dbdfb)
				if _gefa != nil {
					return 0, _edb.Wrap(_gefa, _bdae, "")
				}
			}
			return _abbb._cceff.Decode(_abbb._ecc)
		}
		_baab, _bdbac := _fa.GetStandardTable(14 + int(_abbb.SbHuffRDHeight))
		if _bdbac != nil {
			return 0, _edb.Wrap(_bdbac, _bdae, "")
		}
		return _baab.Decode(_abbb._ecc)
	}
	_ffec, _afaf := _abbb._egdc.DecodeInt(_abbb._acecd)
	if _afaf != nil {
		return 0, _edb.Wrap(_afaf, _bdae, "")
	}
	return int64(_ffec), nil
}
func (_eafb *SymbolDictionary) decodeRefinedSymbol(_dace, _dfcg uint32) error {
	var (
		_fbecd         int
		_cfdef, _aggad int32
	)
	if _eafb.IsHuffmanEncoded {
		_adeda, _cdfa := _eafb._caba.ReadBits(byte(_eafb._abcgg))
		if _cdfa != nil {
			return _cdfa
		}
		_fbecd = int(_adeda)
		_befa, _cdfa := _fa.GetStandardTable(15)
		if _cdfa != nil {
			return _cdfa
		}
		_bbga, _cdfa := _befa.Decode(_eafb._caba)
		if _cdfa != nil {
			return _cdfa
		}
		_cfdef = int32(_bbga)
		_bbga, _cdfa = _befa.Decode(_eafb._caba)
		if _cdfa != nil {
			return _cdfa
		}
		_aggad = int32(_bbga)
		_befa, _cdfa = _fa.GetStandardTable(1)
		if _cdfa != nil {
			return _cdfa
		}
		if _, _cdfa = _befa.Decode(_eafb._caba); _cdfa != nil {
			return _cdfa
		}
		_eafb._caba.Align()
	} else {
		_gbea, _gfef := _eafb._dgcb.DecodeIAID(uint64(_eafb._abcgg), _eafb._ggc)
		if _gfef != nil {
			return _gfef
		}
		_fbecd = int(_gbea)
		_cfdef, _gfef = _eafb._dgcb.DecodeInt(_eafb._fagc)
		if _gfef != nil {
			return _gfef
		}
		_aggad, _gfef = _eafb._dgcb.DecodeInt(_eafb._fccf)
		if _gfef != nil {
			return _gfef
		}
	}
	if _acagg := _eafb.setSymbolsArray(); _acagg != nil {
		return _acagg
	}
	_dfccb := _eafb._ebfb[_fbecd]
	if _fdfc := _eafb.decodeNewSymbols(_dace, _dfcg, _dfccb, _cfdef, _aggad); _fdfc != nil {
		return _fdfc
	}
	if _eafb.IsHuffmanEncoded {
		_eafb._caba.Align()
	}
	return nil
}
func (_fege *TextRegion) decodeSymbolInstances() error {
	_ebbb, _aeaf := _fege.decodeStripT()
	if _aeaf != nil {
		return _aeaf
	}
	var (
		_dbbg int64
		_fece uint32
	)
	for _fece < _fege.NumberOfSymbolInstances {
		_bcef, _cgedc := _fege.decodeDT()
		if _cgedc != nil {
			return _cgedc
		}
		_ebbb += _bcef
		var _cfca int64
		_bfbgd := true
		_fege._abca = 0
		for {
			if _bfbgd {
				_cfca, _cgedc = _fege.decodeDfs()
				if _cgedc != nil {
					return _cgedc
				}
				_dbbg += _cfca
				_fege._abca = _dbbg
				_bfbgd = false
			} else {
				_ddbgc, _cbdd := _fege.decodeIds()
				if _aa.Is(_cbdd, _f.ErrOOB) {
					break
				}
				if _cbdd != nil {
					return _cbdd
				}
				if _fece >= _fege.NumberOfSymbolInstances {
					break
				}
				_fege._abca += _ddbgc + int64(_fege.SbDsOffset)
			}
			_ccca, _cagf := _fege.decodeCurrentT()
			if _cagf != nil {
				return _cagf
			}
			_bbb := _ebbb + _ccca
			_dfff, _cagf := _fege.decodeID()
			if _cagf != nil {
				return _cagf
			}
			_eefd, _cagf := _fege.decodeRI()
			if _cagf != nil {
				return _cagf
			}
			_eede, _cagf := _fege.decodeIb(_eefd, _dfff)
			if _cagf != nil {
				return _cagf
			}
			if _cagf = _fege.blit(_eede, _bbb); _cagf != nil {
				return _cagf
			}
			_fece++
		}
	}
	return nil
}
func (_cafd *PageInformationSegment) Size() int { return 19 }
func (_agef *HalftoneRegion) computeGrayScalePlanes(_bab []*_gf.Bitmap, _bef int) ([][]int, error) {
	_fda := make([][]int, _agef.HGridHeight)
	for _gdbb := 0; _gdbb < len(_fda); _gdbb++ {
		_fda[_gdbb] = make([]int, _agef.HGridWidth)
	}
	for _dfag := 0; _dfag < int(_agef.HGridHeight); _dfag++ {
		for _gce := 0; _gce < int(_agef.HGridWidth); _gce += 8 {
			var _eaaa int
			if _cbd := int(_agef.HGridWidth) - _gce; _cbd > 8 {
				_eaaa = 8
			} else {
				_eaaa = _cbd
			}
			_cacd := _bab[0].GetByteIndex(_gce, _dfag)
			for _gagb := 0; _gagb < _eaaa; _gagb++ {
				_gfe := _gagb + _gce
				_fda[_dfag][_gfe] = 0
				for _ecgbb := 0; _ecgbb < _bef; _ecgbb++ {
					_ffb, _babb := _bab[_ecgbb].GetByte(_cacd)
					if _babb != nil {
						return nil, _babb
					}
					_efd := _ffb >> uint(7-_gfe&7)
					_afed := _efd & 1
					_ggd := 1 << uint(_ecgbb)
					_cgdd := int(_afed) * _ggd
					_fda[_dfag][_gfe] += _cgdd
				}
			}
		}
	}
	return _fda, nil
}
func (_gga *GenericRegion) decodeTemplate0b(_ddec, _ecgb, _dedd int, _bdg, _bafa int) (_fga error) {
	const _dede = "\u0064\u0065c\u006f\u0064\u0065T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0030\u0062"
	var (
		_ebeg, _fbed int
		_bbce, _ebd  int
		_fcb         byte
		_dae         int
	)
	if _ddec >= 1 {
		_fcb, _fga = _gga.Bitmap.GetByte(_bafa)
		if _fga != nil {
			return _edb.Wrap(_fga, _dede, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_bbce = int(_fcb)
	}
	if _ddec >= 2 {
		_fcb, _fga = _gga.Bitmap.GetByte(_bafa - _gga.Bitmap.RowStride)
		if _fga != nil {
			return _edb.Wrap(_fga, _dede, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_ebd = int(_fcb) << 6
	}
	_ebeg = (_bbce & 0xf0) | (_ebd & 0x3800)
	for _fgdd := 0; _fgdd < _dedd; _fgdd = _dae {
		var (
			_edf byte
			_bed int
		)
		_dae = _fgdd + 8
		if _cbg := _ecgb - _fgdd; _cbg > 8 {
			_bed = 8
		} else {
			_bed = _cbg
		}
		if _ddec > 0 {
			_bbce <<= 8
			if _dae < _ecgb {
				_fcb, _fga = _gga.Bitmap.GetByte(_bafa + 1)
				if _fga != nil {
					return _edb.Wrap(_fga, _dede, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_bbce |= int(_fcb)
			}
		}
		if _ddec > 1 {
			_ebd <<= 8
			if _dae < _ecgb {
				_fcb, _fga = _gga.Bitmap.GetByte(_bafa - _gga.Bitmap.RowStride + 1)
				if _fga != nil {
					return _edb.Wrap(_fga, _dede, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_ebd |= int(_fcb) << 6
			}
		}
		for _aebd := 0; _aebd < _bed; _aebd++ {
			_eba := uint(7 - _aebd)
			if _gga._cab {
				_fbed = _gga.overrideAtTemplate0b(_ebeg, _fgdd+_aebd, _ddec, int(_edf), _aebd, int(_eba))
				_gga._fac.SetIndex(int32(_fbed))
			} else {
				_gga._fac.SetIndex(int32(_ebeg))
			}
			var _beac int
			_beac, _fga = _gga._ddga.DecodeBit(_gga._fac)
			if _fga != nil {
				return _edb.Wrap(_fga, _dede, "")
			}
			_edf |= byte(_beac << _eba)
			_ebeg = ((_ebeg & 0x7bf7) << 1) | _beac | ((_bbce >> _eba) & 0x10) | ((_ebd >> _eba) & 0x800)
		}
		if _dbcb := _gga.Bitmap.SetByte(_bdg, _edf); _dbcb != nil {
			return _edb.Wrap(_dbcb, _dede, "")
		}
		_bdg++
		_bafa++
	}
	return nil
}
func (_ffg *HalftoneRegion) computeX(_dbgc, _egab int) int {
	return _ffg.shiftAndFill(int(_ffg.HGridX) + _dbgc*int(_ffg.HRegionY) + _egab*int(_ffg.HRegionX))
}
func (_faacd *SymbolDictionary) setCodingStatistics() error {
	if _faacd._adba == nil {
		_faacd._adba = _df.NewStats(512, 1)
	}
	if _faacd._ggae == nil {
		_faacd._ggae = _df.NewStats(512, 1)
	}
	if _faacd._bcbb == nil {
		_faacd._bcbb = _df.NewStats(512, 1)
	}
	if _faacd._aagg == nil {
		_faacd._aagg = _df.NewStats(512, 1)
	}
	if _faacd._fdfd == nil {
		_faacd._fdfd = _df.NewStats(512, 1)
	}
	if _faacd.UseRefinementAggregation && _faacd._ggc == nil {
		_faacd._ggc = _df.NewStats(1<<uint(_faacd._abcgg), 1)
		_faacd._fagc = _df.NewStats(512, 1)
		_faacd._fccf = _df.NewStats(512, 1)
	}
	if _faacd._afbd == nil {
		_faacd._afbd = _df.NewStats(65536, 1)
	}
	if _faacd._dgcb == nil {
		var _eege error
		_faacd._dgcb, _eege = _df.New(_faacd._caba)
		if _eege != nil {
			return _eege
		}
	}
	return nil
}
func (_bgbb *SymbolDictionary) encodeRefinementATFlags(_ebaga _g.BinaryWriter) (_gecgg int, _cedb error) {
	const _edcg = "\u0065\u006e\u0063od\u0065\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0041\u0054\u0046\u006c\u0061\u0067\u0073"
	if !_bgbb.UseRefinementAggregation || _bgbb.SdrTemplate != 0 {
		return 0, nil
	}
	for _gbae := 0; _gbae < 2; _gbae++ {
		if _cedb = _ebaga.WriteByte(byte(_bgbb.SdrATX[_gbae])); _cedb != nil {
			return _gecgg, _edb.Wrapf(_cedb, _edcg, "\u0053\u0064\u0072\u0041\u0054\u0058\u005b\u0025\u0064\u005d", _gbae)
		}
		_gecgg++
		if _cedb = _ebaga.WriteByte(byte(_bgbb.SdrATY[_gbae])); _cedb != nil {
			return _gecgg, _edb.Wrapf(_cedb, _edcg, "\u0053\u0064\u0072\u0041\u0054\u0059\u005b\u0025\u0064\u005d", _gbae)
		}
		_gecgg++
	}
	return _gecgg, nil
}
func (_bceeb *PatternDictionary) checkInput() error {
	if _bceeb.HdpHeight < 1 || _bceeb.HdpWidth < 1 {
		return _b.New("in\u0076\u0061l\u0069\u0064\u0020\u0048\u0065\u0061\u0064\u0065\u0072 \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0057\u0069\u0064\u0074\u0068\u002f\u0048\u0065\u0069\u0067\u0068\u0074\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020g\u0072e\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020z\u0065\u0072o")
	}
	if _bceeb.IsMMREncoded {
		if _bceeb.HDTemplate != 0 {
			_aae.Log.Debug("\u0076\u0061\u0072\u0069\u0061\u0062\u006c\u0065\u0020\u0048\u0044\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0030")
		}
	}
	return nil
}
func (_ddag *GenericRegion) copyLineAbove(_agbd int) error {
	_afcc := _agbd * _ddag.Bitmap.RowStride
	_bfff := _afcc - _ddag.Bitmap.RowStride
	for _bbc := 0; _bbc < _ddag.Bitmap.RowStride; _bbc++ {
		_feee, _afec := _ddag.Bitmap.GetByte(_bfff)
		if _afec != nil {
			return _afec
		}
		_bfff++
		if _afec = _ddag.Bitmap.SetByte(_afcc, _feee); _afec != nil {
			return _afec
		}
		_afcc++
	}
	return nil
}

type TableSegment struct {
	_dgdg  *_g.Reader
	_cbad  int32
	_facaf int32
	_efee  int32
	_effba int32
	_dgaf  int32
}

func (_gfdbe *GenericRegion) InitEncode(bm *_gf.Bitmap, xLoc, yLoc, template int, duplicateLineRemoval bool) error {
	const _geg = "\u0047e\u006e\u0065\u0072\u0069\u0063\u0052\u0065\u0067\u0069\u006f\u006e.\u0049\u006e\u0069\u0074\u0045\u006e\u0063\u006f\u0064\u0065"
	if bm == nil {
		return _edb.Error(_geg, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if xLoc < 0 || yLoc < 0 {
		return _edb.Error(_geg, "\u0078\u0020\u0061\u006e\u0064\u0020\u0079\u0020\u006c\u006f\u0063\u0061\u0074i\u006f\u006e\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	_gfdbe.Bitmap = bm
	_gfdbe.GBTemplate = byte(template)
	switch _gfdbe.GBTemplate {
	case 0:
		_gfdbe.GBAtX = []int8{3, -3, 2, -2}
		_gfdbe.GBAtY = []int8{-1, -1, -2, -2}
	case 1:
		_gfdbe.GBAtX = []int8{3}
		_gfdbe.GBAtY = []int8{-1}
	case 2, 3:
		_gfdbe.GBAtX = []int8{2}
		_gfdbe.GBAtY = []int8{-1}
	default:
		return _edb.Errorf(_geg, "\u0070\u0072o\u0076\u0069\u0064\u0065\u0064 \u0074\u0065\u006d\u0070\u006ca\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0030\u002c\u0031\u002c\u0032\u002c\u0033\u007d", template)
	}
	_gfdbe.RegionSegment = &RegionSegment{BitmapHeight: uint32(bm.Height), BitmapWidth: uint32(bm.Width), XLocation: uint32(xLoc), YLocation: uint32(yLoc)}
	_gfdbe.IsTPGDon = duplicateLineRemoval
	return nil
}
func (_egfa *GenericRegion) overrideAtTemplate2(_efcd, _dbdb, _faeg, _aee, _cceg int) int {
	_efcd &= 0x3FB
	if _egfa.GBAtY[0] == 0 && _egfa.GBAtX[0] >= -int8(_cceg) {
		_efcd |= (_aee >> uint(7-(int8(_cceg)+_egfa.GBAtX[0])) & 0x1) << 2
	} else {
		_efcd |= int(_egfa.getPixel(_dbdb+int(_egfa.GBAtX[0]), _faeg+int(_egfa.GBAtY[0]))) << 2
	}
	return _efcd
}
func (_bcg *GenericRefinementRegion) Init(header *Header, r *_g.Reader) error {
	_bcg._cb = header
	_bcg._dd = r
	_bcg.RegionInfo = NewRegionSegment(r)
	return _bcg.parseHeader()
}

var _ SegmentEncoder = &GenericRegion{}

func (_bcfd *TableSegment) StreamReader() *_g.Reader { return _bcfd._dgdg }
func (_fdb *RegionSegment) Size() int                { return 17 }

type Documenter interface {
	GetPage(int) (Pager, error)
	GetGlobalSegment(int) (*Header, error)
}

func (_aabc *TextRegion) String() string {
	_edded := &_d.Builder{}
	_edded.WriteString("\u000a[\u0054E\u0058\u0054\u0020\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u000a")
	_edded.WriteString(_aabc.RegionInfo.String() + "\u000a")
	_edded.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053br\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0025\u0076\u000a", _aabc.SbrTemplate))
	_edded.WriteString(_af.Sprintf("\u0009-\u0020S\u0062\u0044\u0073\u004f\u0066f\u0073\u0065t\u003a\u0020\u0025\u0076\u000a", _aabc.SbDsOffset))
	_edded.WriteString(_af.Sprintf("\t\u002d \u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0050i\u0078\u0065\u006c\u003a %\u0076\u000a", _aabc.DefaultPixel))
	_edded.WriteString(_af.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065r\u0061\u0074\u006fr\u003a \u0025\u0076\u000a", _aabc.CombinationOperator))
	_edded.WriteString(_af.Sprintf("\t\u002d \u0049\u0073\u0054\u0072\u0061\u006e\u0073\u0070o\u0073\u0065\u0064\u003a %\u0076\u000a", _aabc.IsTransposed))
	_edded.WriteString(_af.Sprintf("\u0009\u002d\u0020Re\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0043\u006f\u0072\u006e\u0065\u0072\u003a\u0020\u0025\u0076\u000a", _aabc.ReferenceCorner))
	_edded.WriteString(_af.Sprintf("\t\u002d\u0020\u0055\u0073eR\u0065f\u0069\u006e\u0065\u006d\u0065n\u0074\u003a\u0020\u0025\u0076\u000a", _aabc.UseRefinement))
	_edded.WriteString(_af.Sprintf("\u0009-\u0020\u0049\u0073\u0048\u0075\u0066\u0066\u006d\u0061\u006e\u0045n\u0063\u006f\u0064\u0065\u0064\u003a\u0020\u0025\u0076\u000a", _aabc.IsHuffmanEncoded))
	if _aabc.IsHuffmanEncoded {
		_edded.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053bH\u0075\u0066\u0066\u0052\u0053\u0069\u007a\u0065\u003a\u0020\u0025\u0076\u000a", _aabc.SbHuffRSize))
		_edded.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0059:\u0020\u0025\u0076\u000a", _aabc.SbHuffRDY))
		_edded.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0058:\u0020\u0025\u0076\u000a", _aabc.SbHuffRDX))
		_edded.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053bH\u0075\u0066\u0066\u0052\u0044\u0048\u0065\u0069\u0067\u0068\u0074\u003a\u0020\u0025v\u000a", _aabc.SbHuffRDHeight))
		_edded.WriteString(_af.Sprintf("\t\u002d\u0020\u0053\u0062Hu\u0066f\u0052\u0044\u0057\u0069\u0064t\u0068\u003a\u0020\u0025\u0076\u000a", _aabc.SbHuffRDWidth))
		_edded.WriteString(_af.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0044\u0054\u003a\u0020\u0025\u0076\u000a", _aabc.SbHuffDT))
		_edded.WriteString(_af.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0044\u0053\u003a\u0020\u0025\u0076\u000a", _aabc.SbHuffDS))
		_edded.WriteString(_af.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0046\u0053\u003a\u0020\u0025\u0076\u000a", _aabc.SbHuffFS))
	}
	_edded.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0058:\u0020\u0025\u0076\u000a", _aabc.SbrATX))
	_edded.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0059:\u0020\u0025\u0076\u000a", _aabc.SbrATY))
	_edded.WriteString(_af.Sprintf("\u0009\u002d\u0020N\u0075\u006d\u0062\u0065r\u004f\u0066\u0053\u0079\u006d\u0062\u006fl\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u003a\u0020\u0025\u0076\u000a", _aabc.NumberOfSymbolInstances))
	_edded.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0058:\u0020\u0025\u0076\u000a", _aabc.SbrATX))
	return _edded.String()
}
func (_ece *GenericRegion) GetRegionBitmap() (_ddge *_gf.Bitmap, _gfdb error) {
	if _ece.Bitmap != nil {
		return _ece.Bitmap, nil
	}
	if _ece.IsMMREncoded {
		if _ece._gdf == nil {
			_ece._gdf, _gfdb = _eb.New(_ece._efeg, int(_ece.RegionSegment.BitmapWidth), int(_ece.RegionSegment.BitmapHeight), _ece.DataOffset, _ece.DataLength)
			if _gfdb != nil {
				return nil, _gfdb
			}
		}
		_ece.Bitmap, _gfdb = _ece._gdf.UncompressMMR()
		return _ece.Bitmap, _gfdb
	}
	if _gfdb = _ece.updateOverrideFlags(); _gfdb != nil {
		return nil, _gfdb
	}
	var _bdeg int
	if _ece._ddga == nil {
		_ece._ddga, _gfdb = _df.New(_ece._efeg)
		if _gfdb != nil {
			return nil, _gfdb
		}
	}
	if _ece._fac == nil {
		_ece._fac = _df.NewStats(65536, 1)
	}
	_ece.Bitmap = _gf.New(int(_ece.RegionSegment.BitmapWidth), int(_ece.RegionSegment.BitmapHeight))
	_dedf := int(uint32(_ece.Bitmap.Width+7) & (^uint32(7)))
	for _cbff := 0; _cbff < _ece.Bitmap.Height; _cbff++ {
		if _ece.IsTPGDon {
			var _fead int
			_fead, _gfdb = _ece.decodeSLTP()
			if _gfdb != nil {
				return nil, _gfdb
			}
			_bdeg ^= _fead
		}
		if _bdeg == 1 {
			if _cbff > 0 {
				if _gfdb = _ece.copyLineAbove(_cbff); _gfdb != nil {
					return nil, _gfdb
				}
			}
		} else {
			if _gfdb = _ece.decodeLine(_cbff, _ece.Bitmap.Width, _dedf); _gfdb != nil {
				return nil, _gfdb
			}
		}
	}
	return _ece.Bitmap, nil
}
func (_dagb *SymbolDictionary) encodeSymbols(_ddef _g.BinaryWriter) (_cdfd int, _fgbe error) {
	const _gbab = "\u0065\u006e\u0063o\u0064\u0065\u0053\u0079\u006d\u0062\u006f\u006c"
	_agaa := _afe.New()
	_agaa.Init()
	_dbbd, _fgbe := _dagb._cfcf.SelectByIndexes(_dagb._bcgg)
	if _fgbe != nil {
		return 0, _edb.Wrap(_fgbe, _gbab, "\u0069n\u0069\u0074\u0069\u0061\u006c")
	}
	_fbec := map[*_gf.Bitmap]int{}
	for _fgff, _fdec := range _dbbd.Values {
		_fbec[_fdec] = _fgff
	}
	_dbbd.SortByHeight()
	var _bbeb, _eaca int
	_befb, _fgbe := _dbbd.GroupByHeight()
	if _fgbe != nil {
		return 0, _edb.Wrap(_fgbe, _gbab, "")
	}
	for _, _babf := range _befb.Values {
		_cfdc := _babf.Values[0].Height
		_cbe := _cfdc - _bbeb
		if _fgbe = _agaa.EncodeInteger(_afe.IADH, _cbe); _fgbe != nil {
			return 0, _edb.Wrapf(_fgbe, _gbab, "\u0049\u0041\u0044\u0048\u0020\u0066\u006f\u0072\u0020\u0064\u0068\u003a \u0027\u0025\u0064\u0027", _cbe)
		}
		_bbeb = _cfdc
		_dcff, _fdfe := _babf.GroupByWidth()
		if _fdfe != nil {
			return 0, _edb.Wrapf(_fdfe, _gbab, "\u0068\u0065\u0069g\u0068\u0074\u003a\u0020\u0027\u0025\u0064\u0027", _cfdc)
		}
		var _geaf int
		for _, _aceca := range _dcff.Values {
			for _, _gdbdc := range _aceca.Values {
				_cbab := _gdbdc.Width
				_gcfe := _cbab - _geaf
				if _fdfe = _agaa.EncodeInteger(_afe.IADW, _gcfe); _fdfe != nil {
					return 0, _edb.Wrapf(_fdfe, _gbab, "\u0049\u0041\u0044\u0057\u0020\u0066\u006f\u0072\u0020\u0064\u0077\u003a \u0027\u0025\u0064\u0027", _gcfe)
				}
				_geaf += _gcfe
				if _fdfe = _agaa.EncodeBitmap(_gdbdc, false); _fdfe != nil {
					return 0, _edb.Wrapf(_fdfe, _gbab, "H\u0065i\u0067\u0068\u0074\u003a\u0020\u0025\u0064\u0020W\u0069\u0064\u0074\u0068: \u0025\u0064", _cfdc, _cbab)
				}
				_dgce := _fbec[_gdbdc]
				_dagb._efgef[_dgce] = _eaca
				_eaca++
			}
		}
		if _fdfe = _agaa.EncodeOOB(_afe.IADW); _fdfe != nil {
			return 0, _edb.Wrap(_fdfe, _gbab, "\u0049\u0041\u0044\u0057")
		}
	}
	if _fgbe = _agaa.EncodeInteger(_afe.IAEX, 0); _fgbe != nil {
		return 0, _edb.Wrap(_fgbe, _gbab, "\u0065\u0078p\u006f\u0072\u0074e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	if _fgbe = _agaa.EncodeInteger(_afe.IAEX, len(_dagb._bcgg)); _fgbe != nil {
		return 0, _edb.Wrap(_fgbe, _gbab, "\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0073\u0079m\u0062\u006f\u006c\u0073")
	}
	_agaa.Final()
	_abcef, _fgbe := _agaa.WriteTo(_ddef)
	if _fgbe != nil {
		return 0, _edb.Wrap(_fgbe, _gbab, "\u0077\u0072i\u0074\u0069\u006e\u0067 \u0065\u006ec\u006f\u0064\u0065\u0072\u0020\u0063\u006f\u006et\u0065\u0078\u0074\u0020\u0074\u006f\u0020\u0027\u0077\u0027\u0020\u0077r\u0069\u0074\u0065\u0072")
	}
	return int(_abcef), nil
}
func (_fgc *HalftoneRegion) GetPatterns() ([]*_gf.Bitmap, error) {
	var (
		_cafa []*_gf.Bitmap
		_acec error
	)
	for _, _eaea := range _fgc._ceag.RTSegments {
		var _gdfa Segmenter
		_gdfa, _acec = _eaea.GetSegmentData()
		if _acec != nil {
			_aae.Log.Debug("\u0047e\u0074\u0053\u0065\u0067m\u0065\u006e\u0074\u0044\u0061t\u0061 \u0066a\u0069\u006c\u0065\u0064\u003a\u0020\u0025v", _acec)
			return nil, _acec
		}
		_gcd, _feea := _gdfa.(*PatternDictionary)
		if !_feea {
			_acec = _af.Errorf("\u0072e\u006c\u0061t\u0065\u0064\u0020\u0073e\u0067\u006d\u0065n\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0070at\u0074\u0065\u0072n\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u003a \u0025\u0054", _gdfa)
			return nil, _acec
		}
		var _acbc []*_gf.Bitmap
		_acbc, _acec = _gcd.GetDictionary()
		if _acec != nil {
			_aae.Log.Debug("\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _acec)
			return nil, _acec
		}
		_cafa = append(_cafa, _acbc...)
	}
	return _cafa, nil
}
func (_eddg *Header) readHeaderFlags() error {
	const _gdfaf = "\u0072e\u0061d\u0048\u0065\u0061\u0064\u0065\u0072\u0046\u006c\u0061\u0067\u0073"
	_bgec, _dggg := _eddg.Reader.ReadBit()
	if _dggg != nil {
		return _edb.Wrap(_dggg, _gdfaf, "r\u0065\u0074\u0061\u0069\u006e\u0020\u0066\u006c\u0061\u0067")
	}
	if _bgec != 0 {
		_eddg.RetainFlag = true
	}
	_bgec, _dggg = _eddg.Reader.ReadBit()
	if _dggg != nil {
		return _edb.Wrap(_dggg, _gdfaf, "\u0070\u0061g\u0065\u0020\u0061s\u0073\u006f\u0063\u0069\u0061\u0074\u0069\u006f\u006e")
	}
	if _bgec != 0 {
		_eddg.PageAssociationFieldSize = true
	}
	_daa, _dggg := _eddg.Reader.ReadBits(6)
	if _dggg != nil {
		return _edb.Wrap(_dggg, _gdfaf, "\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065")
	}
	_eddg.Type = Type(int(_daa))
	return nil
}
func (_bff *GenericRegion) String() string {
	_faac := &_d.Builder{}
	_faac.WriteString("\u000a[\u0047E\u004e\u0045\u0052\u0049\u0043 \u0052\u0045G\u0049\u004f\u004e\u005d\u000a")
	_faac.WriteString(_bff.RegionSegment.String() + "\u000a")
	_faac.WriteString(_af.Sprintf("\u0009\u002d\u0020Us\u0065\u0045\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073\u003a\u0020\u0025\u0076\u000a", _bff.UseExtTemplates))
	_faac.WriteString(_af.Sprintf("\u0009\u002d \u0049\u0073\u0054P\u0047\u0044\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _bff.IsTPGDon))
	_faac.WriteString(_af.Sprintf("\u0009-\u0020G\u0042\u0054\u0065\u006d\u0070l\u0061\u0074e\u003a\u0020\u0025\u0064\u000a", _bff.GBTemplate))
	_faac.WriteString(_af.Sprintf("\t\u002d \u0049\u0073\u004d\u004d\u0052\u0045\u006e\u0063o\u0064\u0065\u0064\u003a %\u0076\u000a", _bff.IsMMREncoded))
	_faac.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0047\u0042\u0041\u0074\u0058\u003a\u0020\u0025\u0076\u000a", _bff.GBAtX))
	_faac.WriteString(_af.Sprintf("\u0009\u002d\u0020\u0047\u0042\u0041\u0074\u0059\u003a\u0020\u0025\u0076\u000a", _bff.GBAtY))
	_faac.WriteString(_af.Sprintf("\t\u002d \u0047\u0042\u0041\u0074\u004f\u0076\u0065\u0072r\u0069\u0064\u0065\u003a %\u0076\u000a", _bff.GBAtOverride))
	return _faac.String()
}
func (_ffe *GenericRefinementRegion) decodeOptimized(_ab, _ba, _cg, _bb, _fc, _ge, _fae int) error {
	var (
		_feag error
		_fdf  int
		_gd   int
	)
	_ca := _ab - int(_ffe.ReferenceDY)
	if _dbg := int(-_ffe.ReferenceDX); _dbg > 0 {
		_fdf = _dbg
	}
	_cgg := _ffe.ReferenceBitmap.GetByteIndex(_fdf, _ca)
	if _ffe.ReferenceDX > 0 {
		_gd = int(_ffe.ReferenceDX)
	}
	_bge := _ffe.RegionBitmap.GetByteIndex(_gd, _ab)
	switch _ffe.TemplateID {
	case 0:
		_feag = _ffe.decodeTemplate(_ab, _ba, _cg, _bb, _fc, _ge, _fae, _bge, _ca, _cgg, _ffe._fea)
	case 1:
		_feag = _ffe.decodeTemplate(_ab, _ba, _cg, _bb, _fc, _ge, _fae, _bge, _ca, _cgg, _ffe._ad)
	}
	return _feag
}
func (_gacc *HalftoneRegion) computeSegmentDataStructure() error {
	_gacc.DataOffset = _gacc._fagd.AbsolutePosition()
	_gacc.DataHeaderLength = _gacc.DataOffset - _gacc.DataHeaderOffset
	_gacc.DataLength = int64(_gacc._fagd.AbsoluteLength()) - _gacc.DataHeaderLength
	return nil
}
func (_bgagg *TextRegion) readAmountOfSymbolInstances() error {
	_cbfge, _gdfe := _bgagg._ecc.ReadBits(32)
	if _gdfe != nil {
		return _gdfe
	}
	_bgagg.NumberOfSymbolInstances = uint32(_cbfge & _e.MaxUint32)
	_gdcfc := _bgagg.RegionInfo.BitmapWidth * _bgagg.RegionInfo.BitmapHeight
	if _gdcfc < _bgagg.NumberOfSymbolInstances {
		_aae.Log.Debug("\u004c\u0069\u006d\u0069t\u0069\u006e\u0067\u0020t\u0068\u0065\u0020n\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020d\u0065\u0063\u006f\u0064e\u0064\u0020\u0073\u0079m\u0062\u006f\u006c\u0020\u0069n\u0073\u0074\u0061\u006e\u0063\u0065\u0073 \u0074\u006f\u0020\u006f\u006ee\u0020\u0070\u0065\u0072\u0020\u0070\u0069\u0078\u0065l\u0020\u0028\u0020\u0025\u0064\u0020\u0069\u006e\u0073\u0074\u0065\u0061\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u0029", _gdcfc, _bgagg.NumberOfSymbolInstances)
		_bgagg.NumberOfSymbolInstances = _gdcfc
	}
	return nil
}
func NewHeader(d Documenter, r *_g.Reader, offset int64, organizationType OrganizationType) (*Header, error) {
	_dag := &Header{Reader: r}
	if _accad := _dag.parse(d, r, offset, organizationType); _accad != nil {
		return nil, _edb.Wrap(_accad, "\u004ee\u0077\u0048\u0065\u0061\u0064\u0065r", "")
	}
	return _dag, nil
}
func (_dfgc *SymbolDictionary) Init(h *Header, r *_g.Reader) error {
	_dfgc.Header = h
	_dfgc._caba = r
	return _dfgc.parseHeader()
}
func (_fgda *Header) readNumberOfReferredToSegments(_abce *_g.Reader) (uint64, error) {
	const _cccg = "\u0072\u0065\u0061\u0064\u004e\u0075\u006d\u0062\u0065\u0072O\u0066\u0052\u0065\u0066\u0065\u0072\u0072e\u0064\u0054\u006f\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073"
	_eaff, _eagb := _abce.ReadBits(3)
	if _eagb != nil {
		return 0, _edb.Wrap(_eagb, _cccg, "\u0063\u006f\u0075n\u0074\u0020\u006f\u0066\u0020\u0072\u0074\u0073")
	}
	_eaff &= 0xf
	var _bbde []byte
	if _eaff <= 4 {
		_bbde = make([]byte, 5)
		for _afce := 0; _afce <= 4; _afce++ {
			_cfdg, _dafga := _abce.ReadBit()
			if _dafga != nil {
				return 0, _edb.Wrap(_dafga, _cccg, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
			}
			_bbde[_afce] = byte(_cfdg)
		}
	} else {
		_eaff, _eagb = _abce.ReadBits(29)
		if _eagb != nil {
			return 0, _eagb
		}
		_eaff &= _e.MaxInt32
		_eaae := (_eaff + 8) >> 3
		_eaae <<= 3
		_bbde = make([]byte, _eaae)
		var _cdeb uint64
		for _cdeb = 0; _cdeb < _eaae; _cdeb++ {
			_edad, _gceb := _abce.ReadBit()
			if _gceb != nil {
				return 0, _edb.Wrap(_gceb, _cccg, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
			}
			_bbde[_cdeb] = byte(_edad)
		}
	}
	return _eaff, nil
}
func (_dbdf *GenericRegion) GetRegionInfo() *RegionSegment { return _dbdf.RegionSegment }
func (_efgee *TextRegion) GetRegionBitmap() (*_gf.Bitmap, error) {
	if _efgee.RegionBitmap != nil {
		return _efgee.RegionBitmap, nil
	}
	if !_efgee.IsHuffmanEncoded {
		if _gfca := _efgee.setCodingStatistics(); _gfca != nil {
			return nil, _gfca
		}
	}
	if _gadc := _efgee.createRegionBitmap(); _gadc != nil {
		return nil, _gadc
	}
	if _bbdc := _efgee.decodeSymbolInstances(); _bbdc != nil {
		return nil, _bbdc
	}
	return _efgee.RegionBitmap, nil
}
func (_ggdae *TextRegion) setCodingStatistics() error {
	if _ggdae._gede == nil {
		_ggdae._gede = _df.NewStats(512, 1)
	}
	if _ggdae._cgdb == nil {
		_ggdae._cgdb = _df.NewStats(512, 1)
	}
	if _ggdae._dacg == nil {
		_ggdae._dacg = _df.NewStats(512, 1)
	}
	if _ggdae._fff == nil {
		_ggdae._fff = _df.NewStats(512, 1)
	}
	if _ggdae._feeg == nil {
		_ggdae._feeg = _df.NewStats(512, 1)
	}
	if _ggdae._gbcc == nil {
		_ggdae._gbcc = _df.NewStats(512, 1)
	}
	if _ggdae._acecd == nil {
		_ggdae._acecd = _df.NewStats(512, 1)
	}
	if _ggdae._bbae == nil {
		_ggdae._bbae = _df.NewStats(1<<uint(_ggdae._efeb), 1)
	}
	if _ggdae._ffab == nil {
		_ggdae._ffab = _df.NewStats(512, 1)
	}
	if _ggdae._fbeg == nil {
		_ggdae._fbeg = _df.NewStats(512, 1)
	}
	if _ggdae._egdc == nil {
		var _cbdf error
		_ggdae._egdc, _cbdf = _df.New(_ggdae._ecc)
		if _cbdf != nil {
			return _cbdf
		}
	}
	return nil
}
func (_defc *GenericRegion) setOverrideFlag(_dafbf int) {
	_defc.GBAtOverride[_dafbf] = true
	_defc._cab = true
}
func (_cadg *TextRegion) Init(header *Header, r *_g.Reader) error {
	_cadg.Header = header
	_cadg._ecc = r
	_cadg.RegionInfo = NewRegionSegment(_cadg._ecc)
	return _cadg.parseHeader()
}
func (_gdd *SymbolDictionary) setInSyms() error {
	if _gdd.Header.RTSegments != nil {
		return _gdd.retrieveImportSymbols()
	}
	_gdd._bgbd = make([]*_gf.Bitmap, 0)
	return nil
}
func (_gdb *GenericRegion) decodeLine(_ccb, _afb, _fdff int) error {
	const _dfga = "\u0064\u0065\u0063\u006f\u0064\u0065\u004c\u0069\u006e\u0065"
	_cceb := _gdb.Bitmap.GetByteIndex(0, _ccb)
	_beaf := _cceb - _gdb.Bitmap.RowStride
	switch _gdb.GBTemplate {
	case 0:
		if !_gdb.UseExtTemplates {
			return _gdb.decodeTemplate0a(_ccb, _afb, _fdff, _cceb, _beaf)
		}
		return _gdb.decodeTemplate0b(_ccb, _afb, _fdff, _cceb, _beaf)
	case 1:
		return _gdb.decodeTemplate1(_ccb, _afb, _fdff, _cceb, _beaf)
	case 2:
		return _gdb.decodeTemplate2(_ccb, _afb, _fdff, _cceb, _beaf)
	case 3:
		return _gdb.decodeTemplate3(_ccb, _afb, _fdff, _cceb, _beaf)
	}
	return _edb.Errorf(_dfga, "\u0069\u006e\u0076a\u006c\u0069\u0064\u0020G\u0042\u0054\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u003a\u0020\u0025\u0064", _gdb.GBTemplate)
}
func (_cfg *PageInformationSegment) encodeFlags(_gecb _g.BinaryWriter) (_bca error) {
	const _dga = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _bca = _gecb.SkipBits(1); _bca != nil {
		return _edb.Wrap(_bca, _dga, "\u0072\u0065\u0073e\u0072\u0076\u0065\u0064\u0020\u0062\u0069\u0074")
	}
	var _dfdc int
	if _cfg.CombinationOperatorOverrideAllowed() {
		_dfdc = 1
	}
	if _bca = _gecb.WriteBit(_dfdc); _bca != nil {
		return _edb.Wrap(_bca, _dga, "\u0063\u006f\u006db\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0070\u0065\u0072a\u0074\u006f\u0072\u0020\u006f\u0076\u0065\u0072\u0072\u0069\u0064\u0064\u0065\u006e")
	}
	_dfdc = 0
	if _cfg._cbdc {
		_dfdc = 1
	}
	if _bca = _gecb.WriteBit(_dfdc); _bca != nil {
		return _edb.Wrap(_bca, _dga, "\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0061\u0075\u0078i\u006ci\u0061r\u0079\u0020\u0062\u0075\u0066\u0066\u0065r")
	}
	if _bca = _gecb.WriteBit((int(_cfg._dbcd) >> 1) & 0x01); _bca != nil {
		return _edb.Wrap(_bca, _dga, "\u0063\u006f\u006d\u0062\u0069\u006e\u0061\u0074\u0069\u006fn\u0020\u006f\u0070\u0065\u0072\u0061\u0074o\u0072\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0062\u0069\u0074")
	}
	if _bca = _gecb.WriteBit(int(_cfg._dbcd) & 0x01); _bca != nil {
		return _edb.Wrap(_bca, _dga, "\u0063\u006f\u006db\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0070\u0065\u0072a\u0074\u006f\u0072\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0062\u0069\u0074")
	}
	_dfdc = int(_cfg.DefaultPixelValue)
	if _bca = _gecb.WriteBit(_dfdc); _bca != nil {
		return _edb.Wrap(_bca, _dga, "\u0064e\u0066\u0061\u0075\u006c\u0074\u0020\u0070\u0061\u0067\u0065\u0020p\u0069\u0078\u0065\u006c\u0020\u0076\u0061\u006c\u0075\u0065")
	}
	_dfdc = 0
	if _cfg._ccdd {
		_dfdc = 1
	}
	if _bca = _gecb.WriteBit(_dfdc); _bca != nil {
		return _edb.Wrap(_bca, _dga, "\u0063\u006f\u006e\u0074ai\u006e\u0073\u0020\u0072\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074")
	}
	_dfdc = 0
	if _cfg.IsLossless {
		_dfdc = 1
	}
	if _bca = _gecb.WriteBit(_dfdc); _bca != nil {
		return _edb.Wrap(_bca, _dga, "p\u0061\u0067\u0065\u0020\u0069\u0073 \u0065\u0076\u0065\u006e\u0074\u0075\u0061\u006c\u006cy\u0020\u006c\u006fs\u0073l\u0065\u0073\u0073")
	}
	return nil
}
func (_ddcg *RegionSegment) readCombinationOperator() error {
	_ecd, _fge := _ddcg._ebecg.ReadBits(3)
	if _fge != nil {
		return _fge
	}
	_ddcg.CombinaionOperator = _gf.CombinationOperator(_ecd & 0xF)
	return nil
}
func (_edg *GenericRefinementRegion) decodeSLTP() (int, error) {
	_edg.Template.setIndex(_edg._cd)
	return _edg._dc.DecodeBit(_edg._cd)
}
func (_acbac *PageInformationSegment) Encode(w _g.BinaryWriter) (_agbbf int, _agefe error) {
	const _ddagf = "\u0050\u0061g\u0065\u0049\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u002e\u0045\u006eco\u0064\u0065"
	_bedc := make([]byte, 4)
	_ag.BigEndian.PutUint32(_bedc, uint32(_acbac.PageBMWidth))
	_agbbf, _agefe = w.Write(_bedc)
	if _agefe != nil {
		return _agbbf, _edb.Wrap(_agefe, _ddagf, "\u0077\u0069\u0064t\u0068")
	}
	_ag.BigEndian.PutUint32(_bedc, uint32(_acbac.PageBMHeight))
	var _accac int
	_accac, _agefe = w.Write(_bedc)
	if _agefe != nil {
		return _accac + _agbbf, _edb.Wrap(_agefe, _ddagf, "\u0068\u0065\u0069\u0067\u0068\u0074")
	}
	_agbbf += _accac
	_ag.BigEndian.PutUint32(_bedc, uint32(_acbac.ResolutionX))
	_accac, _agefe = w.Write(_bedc)
	if _agefe != nil {
		return _accac + _agbbf, _edb.Wrap(_agefe, _ddagf, "\u0078\u0020\u0072e\u0073\u006f\u006c\u0075\u0074\u0069\u006f\u006e")
	}
	_agbbf += _accac
	_ag.BigEndian.PutUint32(_bedc, uint32(_acbac.ResolutionY))
	if _accac, _agefe = w.Write(_bedc); _agefe != nil {
		return _accac + _agbbf, _edb.Wrap(_agefe, _ddagf, "\u0079\u0020\u0072e\u0073\u006f\u006c\u0075\u0074\u0069\u006f\u006e")
	}
	_agbbf += _accac
	if _agefe = _acbac.encodeFlags(w); _agefe != nil {
		return _agbbf, _edb.Wrap(_agefe, _ddagf, "")
	}
	_agbbf++
	if _accac, _agefe = _acbac.encodeStripingInformation(w); _agefe != nil {
		return _agbbf, _edb.Wrap(_agefe, _ddagf, "")
	}
	_agbbf += _accac
	return _agbbf, nil
}
func (_edaee *PageInformationSegment) checkInput() error {
	if _edaee.PageBMHeight == _e.MaxInt32 {
		if !_edaee.IsStripe {
			_aae.Log.Debug("P\u0061\u0067\u0065\u0049\u006e\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0053\u0065\u0067me\u006e\u0074\u002e\u0049s\u0053\u0074\u0072\u0069\u0070\u0065\u0020\u0073\u0068ou\u006c\u0064 \u0062\u0065\u0020\u0074\u0072\u0075\u0065\u002e")
		}
	}
	return nil
}
func _adda(_caedd *_g.Reader, _fbafe *Header) *TextRegion {
	_fdacb := &TextRegion{_ecc: _caedd, Header: _fbafe, RegionInfo: NewRegionSegment(_caedd)}
	return _fdacb
}
func (_fbgg *SymbolDictionary) decodeHeightClassDeltaHeight() (int64, error) {
	if _fbgg.IsHuffmanEncoded {
		return _fbgg.decodeHeightClassDeltaHeightWithHuffman()
	}
	_caee, _bdde := _fbgg._dgcb.DecodeInt(_fbgg._ggae)
	if _bdde != nil {
		return 0, _bdde
	}
	return int64(_caee), nil
}
func (_bgbdb *TextRegion) GetRegionInfo() *RegionSegment { return _bgbdb.RegionInfo }
func (_fbba *TableSegment) HtHigh() int32                { return _fbba._dgaf }
func (_fcg *GenericRegion) setParametersMMR(_fbbe bool, _bede, _abgd int64, _gebd, _adg uint32, _aca byte, _gaa, _acad bool, _aegf, _ggbg []int8) {
	_fcg.DataOffset = _bede
	_fcg.DataLength = _abgd
	_fcg.RegionSegment = &RegionSegment{}
	_fcg.RegionSegment.BitmapHeight = _gebd
	_fcg.RegionSegment.BitmapWidth = _adg
	_fcg.GBTemplate = _aca
	_fcg.IsMMREncoded = _fbbe
	_fcg.IsTPGDon = _gaa
	_fcg.GBAtX = _aegf
	_fcg.GBAtY = _ggbg
}
func (_bafb *GenericRegion) Encode(w _g.BinaryWriter) (_aab int, _ggb error) {
	const _egf = "G\u0065n\u0065\u0072\u0069\u0063\u0052\u0065\u0067\u0069o\u006e\u002e\u0045\u006eco\u0064\u0065"
	if _bafb.Bitmap == nil {
		return 0, _edb.Error(_egf, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	_ggeg, _ggb := _bafb.RegionSegment.Encode(w)
	if _ggb != nil {
		return 0, _edb.Wrap(_ggb, _egf, "\u0052\u0065\u0067\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_aab += _ggeg
	if _ggb = w.SkipBits(4); _ggb != nil {
		return _aab, _edb.Wrap(_ggb, _egf, "\u0073k\u0069p\u0020\u0072\u0065\u0073\u0065r\u0076\u0065d\u0020\u0062\u0069\u0074\u0073")
	}
	var _bde int
	if _bafb.IsTPGDon {
		_bde = 1
	}
	if _ggb = w.WriteBit(_bde); _ggb != nil {
		return _aab, _edb.Wrap(_ggb, _egf, "\u0074\u0070\u0067\u0064\u006f\u006e")
	}
	_bde = 0
	if _ggb = w.WriteBit(int(_bafb.GBTemplate>>1) & 0x01); _ggb != nil {
		return _aab, _edb.Wrap(_ggb, _egf, "f\u0069r\u0073\u0074\u0020\u0067\u0062\u0074\u0065\u006dp\u006c\u0061\u0074\u0065 b\u0069\u0074")
	}
	if _ggb = w.WriteBit(int(_bafb.GBTemplate) & 0x01); _ggb != nil {
		return _aab, _edb.Wrap(_ggb, _egf, "s\u0065\u0063\u006f\u006ed \u0067b\u0074\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u0062\u0069\u0074")
	}
	if _bafb.UseMMR {
		_bde = 1
	}
	if _ggb = w.WriteBit(_bde); _ggb != nil {
		return _aab, _edb.Wrap(_ggb, _egf, "u\u0073\u0065\u0020\u004d\u004d\u0052\u0020\u0062\u0069\u0074")
	}
	_aab++
	if _ggeg, _ggb = _bafb.writeGBAtPixels(w); _ggb != nil {
		return _aab, _edb.Wrap(_ggb, _egf, "")
	}
	_aab += _ggeg
	_abgc := _afe.New()
	if _ggb = _abgc.EncodeBitmap(_bafb.Bitmap, _bafb.IsTPGDon); _ggb != nil {
		return _aab, _edb.Wrap(_ggb, _egf, "")
	}
	_abgc.Final()
	var _becg int64
	if _becg, _ggb = _abgc.WriteTo(w); _ggb != nil {
		return _aab, _edb.Wrap(_ggb, _egf, "")
	}
	_aab += int(_becg)
	return _aab, nil
}
func (_gdbbg *TextRegion) decodeDT() (_fbbb int64, _acdeg error) {
	if _gdbbg.IsHuffmanEncoded {
		if _gdbbg.SbHuffDT == 3 {
			_fbbb, _acdeg = _gdbbg._bbfg.Decode(_gdbbg._ecc)
			if _acdeg != nil {
				return 0, _acdeg
			}
		} else {
			var _gbeeb _fa.Tabler
			_gbeeb, _acdeg = _fa.GetStandardTable(11 + int(_gdbbg.SbHuffDT))
			if _acdeg != nil {
				return 0, _acdeg
			}
			_fbbb, _acdeg = _gbeeb.Decode(_gdbbg._ecc)
			if _acdeg != nil {
				return 0, _acdeg
			}
		}
	} else {
		var _bfdb int32
		_bfdb, _acdeg = _gdbbg._egdc.DecodeInt(_gdbbg._gede)
		if _acdeg != nil {
			return
		}
		_fbbb = int64(_bfdb)
	}
	_fbbb *= int64(_gdbbg.SbStrips)
	return _fbbb, nil
}
func (_dfb *EndOfStripe) Init(h *Header, r *_g.Reader) error { _dfb._gb = r; return _dfb.parseHeader() }
func (_bcga *TextRegion) decodeID() (int64, error) {
	if _bcga.IsHuffmanEncoded {
		if _bcga._dbcbe == nil {
			_gfbd, _ccec := _bcga._ecc.ReadBits(byte(_bcga._efeb))
			return int64(_gfbd), _ccec
		}
		return _bcga._dbcbe.Decode(_bcga._ecc)
	}
	return _bcga._egdc.DecodeIAID(uint64(_bcga._efeb), _bcga._bbae)
}
func (_daeb *PageInformationSegment) readCombinationOperatorOverrideAllowed() error {
	_dfda, _dddb := _daeb._faef.ReadBit()
	if _dddb != nil {
		return _dddb
	}
	if _dfda == 1 {
		_daeb._fbbge = true
	}
	return nil
}
func _agf(_ccc *_g.Reader, _aaef *Header) *GenericRefinementRegion {
	return &GenericRefinementRegion{_dd: _ccc, RegionInfo: NewRegionSegment(_ccc), _cb: _aaef, _fea: &template0{}, _ad: &template1{}}
}
func (_cabb *TableSegment) HtPS() int32 { return _cabb._facaf }
func (_eabgd *SymbolDictionary) getSymbol(_fbge int) (*_gf.Bitmap, error) {
	const _fbfa = "\u0067e\u0074\u0053\u0079\u006d\u0062\u006fl"
	_geab, _dgdc := _eabgd._cfcf.GetBitmap(_eabgd._bcgg[_fbge])
	if _dgdc != nil {
		return nil, _edb.Wrap(_dgdc, _fbfa, "\u0063\u0061n\u0027\u0074\u0020g\u0065\u0074\u0020\u0073\u0079\u006d\u0062\u006f\u006c")
	}
	return _geab, nil
}
func (_gbde *TextRegion) decodeRdx() (int64, error) {
	const _feda = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064x"
	if _gbde.IsHuffmanEncoded {
		if _gbde.SbHuffRDX == 3 {
			if _gbde._adgga == nil {
				var (
					_aaeb int
					_eeeb error
				)
				if _gbde.SbHuffFS == 3 {
					_aaeb++
				}
				if _gbde.SbHuffDS == 3 {
					_aaeb++
				}
				if _gbde.SbHuffDT == 3 {
					_aaeb++
				}
				if _gbde.SbHuffRDWidth == 3 {
					_aaeb++
				}
				if _gbde.SbHuffRDHeight == 3 {
					_aaeb++
				}
				_gbde._adgga, _eeeb = _gbde.getUserTable(_aaeb)
				if _eeeb != nil {
					return 0, _edb.Wrap(_eeeb, _feda, "")
				}
			}
			return _gbde._adgga.Decode(_gbde._ecc)
		}
		_egebe, _cbga := _fa.GetStandardTable(14 + int(_gbde.SbHuffRDX))
		if _cbga != nil {
			return 0, _edb.Wrap(_cbga, _feda, "")
		}
		return _egebe.Decode(_gbde._ecc)
	}
	_cdadc, _ebac := _gbde._egdc.DecodeInt(_gbde._ffab)
	if _ebac != nil {
		return 0, _edb.Wrap(_ebac, _feda, "")
	}
	return int64(_cdadc), nil
}
func _ffbc(_ddeca int) int {
	if _ddeca == 0 {
		return 0
	}
	_ddeca |= _ddeca >> 1
	_ddeca |= _ddeca >> 2
	_ddeca |= _ddeca >> 4
	_ddeca |= _ddeca >> 8
	_ddeca |= _ddeca >> 16
	return (_ddeca + 1) >> 1
}
func (_bcb *HalftoneRegion) computeY(_acbca, _edeg int) int {
	return _bcb.shiftAndFill(int(_bcb.HGridY) + _acbca*int(_bcb.HRegionX) - _edeg*int(_bcb.HRegionY))
}
func (_abbf *PageInformationSegment) readIsLossless() error {
	_add, _aced := _abbf._faef.ReadBit()
	if _aced != nil {
		return _aced
	}
	if _add == 1 {
		_abbf.IsLossless = true
	}
	return nil
}
func (_gc *GenericRefinementRegion) GetRegionBitmap() (*_gf.Bitmap, error) {
	var _ea error
	_aae.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0047\u0065\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		if _ea != nil {
			_aae.Log.Trace("[\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052E\u0046\u002d\u0052\u0045\u0047\u0049\u004fN]\u0020\u0047\u0065\u0074R\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006dap\u0020\u0066a\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _ea)
		} else {
			_aae.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0047\u0065\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		}
	}()
	if _gc.RegionBitmap != nil {
		return _gc.RegionBitmap, nil
	}
	_ef := 0
	if _gc.ReferenceBitmap == nil {
		_gc.ReferenceBitmap, _ea = _gc.getGrReference()
		if _ea != nil {
			return nil, _ea
		}
	}
	if _gc._dc == nil {
		_gc._dc, _ea = _df.New(_gc._dd)
		if _ea != nil {
			return nil, _ea
		}
	}
	if _gc._cd == nil {
		_gc._cd = _df.NewStats(8192, 1)
	}
	_gc.RegionBitmap = _gf.New(int(_gc.RegionInfo.BitmapWidth), int(_gc.RegionInfo.BitmapHeight))
	if _gc.TemplateID == 0 {
		if _ea = _gc.updateOverride(); _ea != nil {
			return nil, _ea
		}
	}
	_ebf := (_gc.RegionBitmap.Width + 7) & -8
	var _dde int
	if _gc.IsTPGROn {
		_dde = int(-_gc.ReferenceDY) * _gc.ReferenceBitmap.RowStride
	}
	_fd := _dde + 1
	for _fg := 0; _fg < _gc.RegionBitmap.Height; _fg++ {
		if _gc.IsTPGROn {
			_db, _aad := _gc.decodeSLTP()
			if _aad != nil {
				return nil, _aad
			}
			_ef ^= _db
		}
		if _ef == 0 {
			_ea = _gc.decodeOptimized(_fg, _gc.RegionBitmap.Width, _gc.RegionBitmap.RowStride, _gc.ReferenceBitmap.RowStride, _ebf, _dde, _fd)
			if _ea != nil {
				return nil, _ea
			}
		} else {
			_ea = _gc.decodeTypicalPredictedLine(_fg, _gc.RegionBitmap.Width, _gc.RegionBitmap.RowStride, _gc.ReferenceBitmap.RowStride, _ebf, _dde)
			if _ea != nil {
				return nil, _ea
			}
		}
	}
	return _gc.RegionBitmap, nil
}
func (_gfefd *TextRegion) decodeRdy() (int64, error) {
	const _dagc = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064y"
	if _gfefd.IsHuffmanEncoded {
		if _gfefd.SbHuffRDY == 3 {
			if _gfefd._bfde == nil {
				var (
					_cggd  int
					_aeffa error
				)
				if _gfefd.SbHuffFS == 3 {
					_cggd++
				}
				if _gfefd.SbHuffDS == 3 {
					_cggd++
				}
				if _gfefd.SbHuffDT == 3 {
					_cggd++
				}
				if _gfefd.SbHuffRDWidth == 3 {
					_cggd++
				}
				if _gfefd.SbHuffRDHeight == 3 {
					_cggd++
				}
				if _gfefd.SbHuffRDX == 3 {
					_cggd++
				}
				_gfefd._bfde, _aeffa = _gfefd.getUserTable(_cggd)
				if _aeffa != nil {
					return 0, _edb.Wrap(_aeffa, _dagc, "")
				}
			}
			return _gfefd._bfde.Decode(_gfefd._ecc)
		}
		_dabc, _ffdg := _fa.GetStandardTable(14 + int(_gfefd.SbHuffRDY))
		if _ffdg != nil {
			return 0, _ffdg
		}
		return _dabc.Decode(_gfefd._ecc)
	}
	_fdad, _dbac := _gfefd._egdc.DecodeInt(_gfefd._fbeg)
	if _dbac != nil {
		return 0, _edb.Wrap(_dbac, _dagc, "")
	}
	return int64(_fdad), nil
}
func (_egfd *Header) readSegmentDataLength(_ffeg *_g.Reader) (_cfbd error) {
	_egfd.SegmentDataLength, _cfbd = _ffeg.ReadBits(32)
	if _cfbd != nil {
		return _cfbd
	}
	_egfd.SegmentDataLength &= _e.MaxInt32
	return nil
}
func (_fbd *GenericRefinementRegion) parseHeader() (_cgb error) {
	_aae.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072s\u0069\u006e\u0067\u0020\u0048e\u0061\u0064e\u0072\u002e\u002e\u002e")
	_dda := _cf.Now()
	defer func() {
		if _cgb == nil {
			_aae.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045G\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020h\u0065\u0061\u0064\u0065\u0072\u0020\u0066\u0069\u006e\u0069\u0073\u0068id\u0020\u0069\u006e\u003a\u0020\u0025\u0064\u0020\u006e\u0073", _cf.Since(_dda).Nanoseconds())
		} else {
			_aae.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0073", _cgb)
		}
	}()
	if _cgb = _fbd.RegionInfo.parseHeader(); _cgb != nil {
		return _cgb
	}
	_, _cgb = _fbd._dd.ReadBits(6)
	if _cgb != nil {
		return _cgb
	}
	_fbd.IsTPGROn, _cgb = _fbd._dd.ReadBool()
	if _cgb != nil {
		return _cgb
	}
	var _gbc int
	_gbc, _cgb = _fbd._dd.ReadBit()
	if _cgb != nil {
		return _cgb
	}
	_fbd.TemplateID = int8(_gbc)
	switch _fbd.TemplateID {
	case 0:
		_fbd.Template = _fbd._fea
		if _cgb = _fbd.readAtPixels(); _cgb != nil {
			return
		}
	case 1:
		_fbd.Template = _fbd._ad
	}
	return nil
}
func (_adbab *TextRegion) parseHeader() error {
	var _cgbf error
	_aae.Log.Trace("\u005b\u0054E\u0058\u0054\u0020\u0052E\u0047\u0049O\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045-\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0062\u0065\u0067\u0069n\u0073\u002e\u002e\u002e")
	defer func() {
		if _cgbf != nil {
			_aae.Log.Trace("\u005b\u0054\u0045\u0058\u0054\u0020\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044E\u0052\u005d\u0020\u0066\u0061i\u006c\u0065d\u002e\u0020\u0025\u0076", _cgbf)
		} else {
			_aae.Log.Trace("\u005b\u0054E\u0058\u0054\u0020\u0052E\u0047\u0049O\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045-\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0066\u0069\u006e\u0069s\u0068\u0065\u0064\u002e")
		}
	}()
	if _cgbf = _adbab.RegionInfo.parseHeader(); _cgbf != nil {
		return _cgbf
	}
	if _cgbf = _adbab.readRegionFlags(); _cgbf != nil {
		return _cgbf
	}
	if _adbab.IsHuffmanEncoded {
		if _cgbf = _adbab.readHuffmanFlags(); _cgbf != nil {
			return _cgbf
		}
	}
	if _cgbf = _adbab.readUseRefinement(); _cgbf != nil {
		return _cgbf
	}
	if _cgbf = _adbab.readAmountOfSymbolInstances(); _cgbf != nil {
		return _cgbf
	}
	if _cgbf = _adbab.getSymbols(); _cgbf != nil {
		return _cgbf
	}
	if _cgbf = _adbab.computeSymbolCodeLength(); _cgbf != nil {
		return _cgbf
	}
	if _cgbf = _adbab.checkInput(); _cgbf != nil {
		return _cgbf
	}
	_aae.Log.Trace("\u0025\u0073", _adbab.String())
	return nil
}
func NewGenericRegion(r *_g.Reader) *GenericRegion {
	return &GenericRegion{RegionSegment: NewRegionSegment(r), _efeg: r}
}
func (_gcf *PatternDictionary) parseHeader() error {
	_aae.Log.Trace("\u005b\u0050\u0041\u0054\u0054\u0045\u0052\u004e\u002d\u0044\u0049\u0043\u0054I\u004f\u004e\u0041\u0052\u0059\u005d[\u0070\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072\u005d\u0020b\u0065\u0067\u0069\u006e")
	defer func() {
		_aae.Log.Trace("\u005b\u0050\u0041T\u0054\u0045\u0052\u004e\u002d\u0044\u0049\u0043\u0054\u0049\u004f\u004e\u0041\u0052\u0059\u005d\u005b\u0070\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072\u005d \u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
	}()
	_, _fccb := _gcf._dbec.ReadBits(5)
	if _fccb != nil {
		return _fccb
	}
	if _fccb = _gcf.readTemplate(); _fccb != nil {
		return _fccb
	}
	if _fccb = _gcf.readIsMMREncoded(); _fccb != nil {
		return _fccb
	}
	if _fccb = _gcf.readPatternWidthAndHeight(); _fccb != nil {
		return _fccb
	}
	if _fccb = _gcf.readGrayMax(); _fccb != nil {
		return _fccb
	}
	if _fccb = _gcf.computeSegmentDataStructure(); _fccb != nil {
		return _fccb
	}
	return _gcf.checkInput()
}
func (_aff *Header) pageSize() uint {
	if _aff.PageAssociation <= 255 {
		return 1
	}
	return 4
}

const (
	TSymbolDictionary                         Type = 0
	TIntermediateTextRegion                   Type = 4
	TImmediateTextRegion                      Type = 6
	TImmediateLosslessTextRegion              Type = 7
	TPatternDictionary                        Type = 16
	TIntermediateHalftoneRegion               Type = 20
	TImmediateHalftoneRegion                  Type = 22
	TImmediateLosslessHalftoneRegion          Type = 23
	TIntermediateGenericRegion                Type = 36
	TImmediateGenericRegion                   Type = 38
	TImmediateLosslessGenericRegion           Type = 39
	TIntermediateGenericRefinementRegion      Type = 40
	TImmediateGenericRefinementRegion         Type = 42
	TImmediateLosslessGenericRefinementRegion Type = 43
	TPageInformation                          Type = 48
	TEndOfPage                                Type = 49
	TEndOfStrip                               Type = 50
	TEndOfFile                                Type = 51
	TProfiles                                 Type = 52
	TTables                                   Type = 53
	TExtension                                Type = 62
	TBitmap                                   Type = 70
)

func (_aefg *Header) writeSegmentNumber(_bece _g.BinaryWriter) (_becc int, _adgg error) {
	_fafgd := make([]byte, 4)
	_ag.BigEndian.PutUint32(_fafgd, _aefg.SegmentNumber)
	if _becc, _adgg = _bece.Write(_fafgd); _adgg != nil {
		return 0, _edb.Wrap(_adgg, "\u0048e\u0061\u0064\u0065\u0072.\u0077\u0072\u0069\u0074\u0065S\u0065g\u006de\u006e\u0074\u004e\u0075\u006d\u0062\u0065r", "")
	}
	return _becc, nil
}
func (_dacb *GenericRegion) Init(h *Header, r *_g.Reader) error {
	_dacb.RegionSegment = NewRegionSegment(r)
	_dacb._efeg = r
	return _dacb.parseHeader()
}

var (
	_ Regioner  = &TextRegion{}
	_ Segmenter = &TextRegion{}
)

func (_ffce *Header) readSegmentNumber(_bfac *_g.Reader) error {
	const _eeb = "\u0072\u0065\u0061\u0064\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u004eu\u006d\u0062\u0065\u0072"
	_efda := make([]byte, 4)
	_, _gagf := _bfac.Read(_efda)
	if _gagf != nil {
		return _edb.Wrap(_gagf, _eeb, "")
	}
	_ffce.SegmentNumber = _ag.BigEndian.Uint32(_efda)
	return nil
}
func (_faa *template0) form(_cddg, _bgf, _bfb, _dab, _dbd int16) int16 {
	return (_cddg << 10) | (_bgf << 7) | (_bfb << 4) | (_dab << 1) | _dbd
}
func (_gagd *SymbolDictionary) InitEncode(symbols *_gf.Bitmaps, symbolList []int, symbolMap map[int]int, unborderSymbols bool) error {
	const _efec = "S\u0079\u006d\u0062\u006f\u006c\u0044i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002eI\u006e\u0069\u0074E\u006ec\u006f\u0064\u0065"
	_gagd.SdATX = []int8{3, -3, 2, -2}
	_gagd.SdATY = []int8{-1, -1, -2, -2}
	_gagd._cfcf = symbols
	_gagd._bcgg = make([]int, len(symbolList))
	copy(_gagd._bcgg, symbolList)
	if len(_gagd._bcgg) != _gagd._cfcf.Size() {
		return _edb.Error(_efec, "s\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0061\u006e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u004ci\u0073\u0074\u0020\u006f\u0066\u0020\u0064\u0069\u0066\u0066er\u0065\u006e\u0074 \u0073i\u007a\u0065")
	}
	_gagd.NumberOfNewSymbols = uint32(symbols.Size())
	_gagd.NumberOfExportedSymbols = uint32(symbols.Size())
	_gagd._efgef = symbolMap
	_gagd._gdc = unborderSymbols
	return nil
}
func (_ead *TextRegion) getUserTable(_agfb int) (_fa.Tabler, error) {
	const _edab = "\u0067\u0065\u0074U\u0073\u0065\u0072\u0054\u0061\u0062\u006c\u0065"
	var _ddbb int
	for _, _cef := range _ead.Header.RTSegments {
		if _cef.Type == 53 {
			if _ddbb == _agfb {
				_eafd, _caed := _cef.GetSegmentData()
				if _caed != nil {
					return nil, _caed
				}
				_cebfb, _agggb := _eafd.(*TableSegment)
				if !_agggb {
					_aae.Log.Debug(_af.Sprintf("\u0073\u0065\u0067\u006d\u0065\u006e\u0074 \u0077\u0069\u0074h\u0020\u0054\u0079p\u0065\u00205\u0033\u0020\u002d\u0020\u0061\u006ed\u0020in\u0064\u0065\u0078\u003a\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0054\u0061\u0062\u006c\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074", _cef.SegmentNumber))
					return nil, _edb.Error(_edab, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074 \u0077\u0069\u0074h\u0020\u0054\u0079\u0070e\u0020\u0035\u0033\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u002a\u0054\u0061\u0062\u006c\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
				}
				return _fa.NewEncodedTable(_cebfb)
			}
			_ddbb++
		}
	}
	return nil, nil
}
func (_agbb *GenericRegion) decodeTemplate0a(_cfaa, _dafb, _bfbf int, _adc, _cecc int) (_gbb error) {
	const _caae = "\u0064\u0065c\u006f\u0064\u0065T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0030\u0061"
	var (
		_acca, _cdef int
		_ceb, _fba   int
		_gdg         byte
		_fbgc        int
	)
	if _cfaa >= 1 {
		_gdg, _gbb = _agbb.Bitmap.GetByte(_cecc)
		if _gbb != nil {
			return _edb.Wrap(_gbb, _caae, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_ceb = int(_gdg)
	}
	if _cfaa >= 2 {
		_gdg, _gbb = _agbb.Bitmap.GetByte(_cecc - _agbb.Bitmap.RowStride)
		if _gbb != nil {
			return _edb.Wrap(_gbb, _caae, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_fba = int(_gdg) << 6
	}
	_acca = (_ceb & 0xf0) | (_fba & 0x3800)
	for _efff := 0; _efff < _bfbf; _efff = _fbgc {
		var (
			_bfda byte
			_abcd int
		)
		_fbgc = _efff + 8
		if _bcgb := _dafb - _efff; _bcgb > 8 {
			_abcd = 8
		} else {
			_abcd = _bcgb
		}
		if _cfaa > 0 {
			_ceb <<= 8
			if _fbgc < _dafb {
				_gdg, _gbb = _agbb.Bitmap.GetByte(_cecc + 1)
				if _gbb != nil {
					return _edb.Wrap(_gbb, _caae, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_ceb |= int(_gdg)
			}
		}
		if _cfaa > 1 {
			_deb := _cecc - _agbb.Bitmap.RowStride + 1
			_fba <<= 8
			if _fbgc < _dafb {
				_gdg, _gbb = _agbb.Bitmap.GetByte(_deb)
				if _gbb != nil {
					return _edb.Wrap(_gbb, _caae, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_fba |= int(_gdg) << 6
			} else {
				_fba |= 0
			}
		}
		for _cdde := 0; _cdde < _abcd; _cdde++ {
			_cgba := uint(7 - _cdde)
			if _agbb._cab {
				_cdef = _agbb.overrideAtTemplate0a(_acca, _efff+_cdde, _cfaa, int(_bfda), _cdde, int(_cgba))
				_agbb._fac.SetIndex(int32(_cdef))
			} else {
				_agbb._fac.SetIndex(int32(_acca))
			}
			var _eaa int
			_eaa, _gbb = _agbb._ddga.DecodeBit(_agbb._fac)
			if _gbb != nil {
				return _edb.Wrap(_gbb, _caae, "")
			}
			_bfda |= byte(_eaa) << _cgba
			_acca = ((_acca & 0x7bf7) << 1) | _eaa | ((_ceb >> _cgba) & 0x10) | ((_fba >> _cgba) & 0x800)
		}
		if _geb := _agbb.Bitmap.SetByte(_adc, _bfda); _geb != nil {
			return _edb.Wrap(_geb, _caae, "")
		}
		_adc++
		_cecc++
	}
	return nil
}
func (_gdga *Header) writeReferredToSegments(_ebfc _g.BinaryWriter) (_cedg int, _fegd error) {
	const _gbca = "\u0077\u0072\u0069te\u0052\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0054\u006f\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073"
	var (
		_dfd  uint16
		_feae uint32
	)
	_fggg := _gdga.referenceSize()
	_afaa := 1
	_gedd := make([]byte, _fggg)
	for _, _bdba := range _gdga.RTSNumbers {
		switch _fggg {
		case 4:
			_feae = uint32(_bdba)
			_ag.BigEndian.PutUint32(_gedd, _feae)
			_afaa, _fegd = _ebfc.Write(_gedd)
			if _fegd != nil {
				return 0, _edb.Wrap(_fegd, _gbca, "u\u0069\u006e\u0074\u0033\u0032\u0020\u0073\u0069\u007a\u0065")
			}
		case 2:
			_dfd = uint16(_bdba)
			_ag.BigEndian.PutUint16(_gedd, _dfd)
			_afaa, _fegd = _ebfc.Write(_gedd)
			if _fegd != nil {
				return 0, _edb.Wrap(_fegd, _gbca, "\u0075\u0069\u006e\u0074\u0031\u0036")
			}
		default:
			if _fegd = _ebfc.WriteByte(byte(_bdba)); _fegd != nil {
				return 0, _edb.Wrap(_fegd, _gbca, "\u0075\u0069\u006et\u0038")
			}
		}
		_cedg += _afaa
	}
	return _cedg, nil
}
func (_gece *SymbolDictionary) getToExportFlags() ([]int, error) {
	var (
		_fabb int
		_beb  int32
		_eecf error
		_cgdf = int32(_gece._gfg + _gece.NumberOfNewSymbols)
		_ffd  = make([]int, _cgdf)
	)
	for _gaea := int32(0); _gaea < _cgdf; _gaea += _beb {
		if _gece.IsHuffmanEncoded {
			_gca, _facc := _fa.GetStandardTable(1)
			if _facc != nil {
				return nil, _facc
			}
			_ceaf, _facc := _gca.Decode(_gece._caba)
			if _facc != nil {
				return nil, _facc
			}
			_beb = int32(_ceaf)
		} else {
			_beb, _eecf = _gece._dgcb.DecodeInt(_gece._fdfd)
			if _eecf != nil {
				return nil, _eecf
			}
		}
		if _beb != 0 {
			if _gaea+_beb > _cgdf {
				return nil, _edb.Error("\u0053\u0079\u006d\u0062\u006f\u006cD\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0067\u0065\u0074T\u006f\u0045\u0078\u0070\u006f\u0072\u0074F\u006c\u0061\u0067\u0073", "\u006d\u0061\u006c\u0066\u006f\u0072m\u0065\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0064\u0061\u0074\u0061\u0020\u0070\u0072\u006f\u0076\u0069\u0064e\u0064\u002e\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u006f\u0075\u0074\u0020\u006ff\u0020r\u0061\u006e\u0067\u0065")
			}
			for _dddd := _gaea; _dddd < _gaea+_beb; _dddd++ {
				_ffd[_dddd] = _fabb
			}
		}
		if _fabb == 0 {
			_fabb = 1
		} else {
			_fabb = 0
		}
	}
	return _ffd, nil
}
func (_cff *GenericRefinementRegion) setParameters(_agb *_df.DecoderStats, _abf *_df.Decoder, _fbe int8, _ecf, _ddg uint32, _fdcc *_gf.Bitmap, _ecg, _edae int32, _ebec bool, _bfd []int8, _bced []int8) {
	_aae.Log.Trace("\u005b\u0047\u0045NE\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052E\u0047I\u004fN\u005d \u0073\u0065\u0074\u0050\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	if _agb != nil {
		_cff._cd = _agb
	}
	if _abf != nil {
		_cff._dc = _abf
	}
	_cff.TemplateID = _fbe
	_cff.RegionInfo.BitmapWidth = _ecf
	_cff.RegionInfo.BitmapHeight = _ddg
	_cff.ReferenceBitmap = _fdcc
	_cff.ReferenceDX = _ecg
	_cff.ReferenceDY = _edae
	_cff.IsTPGROn = _ebec
	_cff.GrAtX = _bfd
	_cff.GrAtY = _bced
	_cff.RegionBitmap = nil
	_aae.Log.Trace("[\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052E\u0046\u002d\u0052\u0045\u0047\u0049\u004fN]\u0020\u0073\u0065\u0074P\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073 f\u0069\u006ei\u0073\u0068\u0065\u0064\u002e\u0020\u0025\u0073", _cff)
}
func (_feade *GenericRegion) overrideAtTemplate0a(_eabg, _adb, _afcb, _cead, _gdbf, _deg int) int {
	if _feade.GBAtOverride[0] {
		_eabg &= 0xFFEF
		if _feade.GBAtY[0] == 0 && _feade.GBAtX[0] >= -int8(_gdbf) {
			_eabg |= (_cead >> uint(int8(_deg)-_feade.GBAtX[0]&0x1)) << 4
		} else {
			_eabg |= int(_feade.getPixel(_adb+int(_feade.GBAtX[0]), _afcb+int(_feade.GBAtY[0]))) << 4
		}
	}
	if _feade.GBAtOverride[1] {
		_eabg &= 0xFBFF
		if _feade.GBAtY[1] == 0 && _feade.GBAtX[1] >= -int8(_gdbf) {
			_eabg |= (_cead >> uint(int8(_deg)-_feade.GBAtX[1]&0x1)) << 10
		} else {
			_eabg |= int(_feade.getPixel(_adb+int(_feade.GBAtX[1]), _afcb+int(_feade.GBAtY[1]))) << 10
		}
	}
	if _feade.GBAtOverride[2] {
		_eabg &= 0xF7FF
		if _feade.GBAtY[2] == 0 && _feade.GBAtX[2] >= -int8(_gdbf) {
			_eabg |= (_cead >> uint(int8(_deg)-_feade.GBAtX[2]&0x1)) << 11
		} else {
			_eabg |= int(_feade.getPixel(_adb+int(_feade.GBAtX[2]), _afcb+int(_feade.GBAtY[2]))) << 11
		}
	}
	if _feade.GBAtOverride[3] {
		_eabg &= 0x7FFF
		if _feade.GBAtY[3] == 0 && _feade.GBAtX[3] >= -int8(_gdbf) {
			_eabg |= (_cead >> uint(int8(_deg)-_feade.GBAtX[3]&0x1)) << 15
		} else {
			_eabg |= int(_feade.getPixel(_adb+int(_feade.GBAtX[3]), _afcb+int(_feade.GBAtY[3]))) << 15
		}
	}
	return _eabg
}
func (_bgde *TextRegion) decodeIds() (int64, error) {
	const _gfcg = "\u0064e\u0063\u006f\u0064\u0065\u0049\u0064s"
	if _bgde.IsHuffmanEncoded {
		if _bgde.SbHuffDS == 3 {
			if _bgde._gadb == nil {
				_adgge := 0
				if _bgde.SbHuffFS == 3 {
					_adgge++
				}
				var _caca error
				_bgde._gadb, _caca = _bgde.getUserTable(_adgge)
				if _caca != nil {
					return 0, _edb.Wrap(_caca, _gfcg, "")
				}
			}
			return _bgde._gadb.Decode(_bgde._ecc)
		}
		_badb, _cggb := _fa.GetStandardTable(8 + int(_bgde.SbHuffDS))
		if _cggb != nil {
			return 0, _edb.Wrap(_cggb, _gfcg, "")
		}
		return _badb.Decode(_bgde._ecc)
	}
	_gabc, _cbce := _bgde._egdc.DecodeInt(_bgde._dacg)
	if _cbce != nil {
		return 0, _edb.Wrap(_cbce, _gfcg, "\u0063\u0078\u0049\u0041\u0044\u0053")
	}
	return int64(_gabc), nil
}

type Segmenter interface {
	Init(_fabg *Header, _bedd *_g.Reader) error
}

var _ templater = &template1{}

func (_ff *GenericRefinementRegion) getGrReference() (*_gf.Bitmap, error) {
	segments := _ff._cb.RTSegments
	if len(segments) == 0 {
		return nil, _b.New("\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0065\u0078is\u0074\u0073")
	}
	_cea, _fag := segments[0].GetSegmentData()
	if _fag != nil {
		return nil, _fag
	}
	_feaa, _dbc := _cea.(Regioner)
	if !_dbc {
		return nil, _af.Errorf("\u0072\u0065\u0066\u0065\u0072r\u0065\u0064\u0020\u0074\u006f\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0052\u0065\u0067\u0069\u006f\u006e\u0065\u0072\u003a\u0020\u0025\u0054", _cea)
	}
	return _feaa.GetRegionBitmap()
}
func (_fcc *GenericRegion) readGBAtPixels(_ggac int) error {
	const _aefc = "\u0072\u0065\u0061\u0064\u0047\u0042\u0041\u0074\u0050i\u0078\u0065\u006c\u0073"
	_fcc.GBAtX = make([]int8, _ggac)
	_fcc.GBAtY = make([]int8, _ggac)
	for _agac := 0; _agac < _ggac; _agac++ {
		_fafg, _cebf := _fcc._efeg.ReadByte()
		if _cebf != nil {
			return _edb.Wrapf(_cebf, _aefc, "\u0058\u0020\u0061t\u0020\u0069\u003a\u0020\u0027\u0025\u0064\u0027", _agac)
		}
		_fcc.GBAtX[_agac] = int8(_fafg)
		_fafg, _cebf = _fcc._efeg.ReadByte()
		if _cebf != nil {
			return _edb.Wrapf(_cebf, _aefc, "\u0059\u0020\u0061t\u0020\u0069\u003a\u0020\u0027\u0025\u0064\u0027", _agac)
		}
		_fcc.GBAtY[_agac] = int8(_fafg)
	}
	return nil
}
func (_aafff *PatternDictionary) GetDictionary() ([]*_gf.Bitmap, error) {
	if _aafff.Patterns != nil {
		return _aafff.Patterns, nil
	}
	if !_aafff.IsMMREncoded {
		_aafff.setGbAtPixels()
	}
	_egg := NewGenericRegion(_aafff._dbec)
	_egg.setParametersMMR(_aafff.IsMMREncoded, _aafff.DataOffset, _aafff.DataLength, uint32(_aafff.HdpHeight), (_aafff.GrayMax+1)*uint32(_aafff.HdpWidth), _aafff.HDTemplate, false, false, _aafff.GBAtX, _aafff.GBAtY)
	_bfg, _gcdd := _egg.GetRegionBitmap()
	if _gcdd != nil {
		return nil, _gcdd
	}
	if _gcdd = _aafff.extractPatterns(_bfg); _gcdd != nil {
		return nil, _gcdd
	}
	return _aafff.Patterns, nil
}
func (_ccag *SymbolDictionary) decodeAggregate(_faegc, _dcaf uint32) error {
	var (
		_gegd  int64
		_adaad error
	)
	if _ccag.IsHuffmanEncoded {
		_gegd, _adaad = _ccag.huffDecodeRefAggNInst()
		if _adaad != nil {
			return _adaad
		}
	} else {
		_dffe, _gdcf := _ccag._dgcb.DecodeInt(_ccag._aagg)
		if _gdcf != nil {
			return _gdcf
		}
		_gegd = int64(_dffe)
	}
	if _gegd > 1 {
		return _ccag.decodeThroughTextRegion(_faegc, _dcaf, uint32(_gegd))
	} else if _gegd == 1 {
		return _ccag.decodeRefinedSymbol(_faegc, _dcaf)
	}
	return nil
}
func (_agee *SymbolDictionary) readRegionFlags() error {
	var (
		_gdbff uint64
		_badd  int
	)
	_, _ggg := _agee._caba.ReadBits(3)
	if _ggg != nil {
		return _ggg
	}
	_badd, _ggg = _agee._caba.ReadBit()
	if _ggg != nil {
		return _ggg
	}
	_agee.SdrTemplate = int8(_badd)
	_gdbff, _ggg = _agee._caba.ReadBits(2)
	if _ggg != nil {
		return _ggg
	}
	_agee.SdTemplate = int8(_gdbff & 0xf)
	_badd, _ggg = _agee._caba.ReadBit()
	if _ggg != nil {
		return _ggg
	}
	if _badd == 1 {
		_agee._dedee = true
	}
	_badd, _ggg = _agee._caba.ReadBit()
	if _ggg != nil {
		return _ggg
	}
	if _badd == 1 {
		_agee._bddf = true
	}
	_badd, _ggg = _agee._caba.ReadBit()
	if _ggg != nil {
		return _ggg
	}
	if _badd == 1 {
		_agee.SdHuffAggInstanceSelection = true
	}
	_badd, _ggg = _agee._caba.ReadBit()
	if _ggg != nil {
		return _ggg
	}
	_agee.SdHuffBMSizeSelection = int8(_badd)
	_gdbff, _ggg = _agee._caba.ReadBits(2)
	if _ggg != nil {
		return _ggg
	}
	_agee.SdHuffDecodeWidthSelection = int8(_gdbff & 0xf)
	_gdbff, _ggg = _agee._caba.ReadBits(2)
	if _ggg != nil {
		return _ggg
	}
	_agee.SdHuffDecodeHeightSelection = int8(_gdbff & 0xf)
	_badd, _ggg = _agee._caba.ReadBit()
	if _ggg != nil {
		return _ggg
	}
	if _badd == 1 {
		_agee.UseRefinementAggregation = true
	}
	_badd, _ggg = _agee._caba.ReadBit()
	if _ggg != nil {
		return _ggg
	}
	if _badd == 1 {
		_agee.IsHuffmanEncoded = true
	}
	return nil
}
func (_aeacc *TableSegment) Init(h *Header, r *_g.Reader) error {
	_aeacc._dgdg = r
	return _aeacc.parseHeader()
}
func (_aggef *PageInformationSegment) readWidthAndHeight() error {
	_babba, _geaa := _aggef._faef.ReadBits(32)
	if _geaa != nil {
		return _geaa
	}
	_aggef.PageBMWidth = int(_babba & _e.MaxInt32)
	_babba, _geaa = _aggef._faef.ReadBits(32)
	if _geaa != nil {
		return _geaa
	}
	_aggef.PageBMHeight = int(_babba & _e.MaxInt32)
	return nil
}
func (_ggeb *TextRegion) createRegionBitmap() error {
	_ggeb.RegionBitmap = _gf.New(int(_ggeb.RegionInfo.BitmapWidth), int(_ggeb.RegionInfo.BitmapHeight))
	if _ggeb.DefaultPixel != 0 {
		_ggeb.RegionBitmap.SetDefaultPixel()
	}
	return nil
}
func (_cbcg *TableSegment) HtRS() int32 { return _cbcg._efee }
func (_ffbde *SymbolDictionary) decodeHeightClassDeltaHeightWithHuffman() (int64, error) {
	switch _ffbde.SdHuffDecodeHeightSelection {
	case 0:
		_cgdab, _gbgb := _fa.GetStandardTable(4)
		if _gbgb != nil {
			return 0, _gbgb
		}
		return _cgdab.Decode(_ffbde._caba)
	case 1:
		_aggcb, _dade := _fa.GetStandardTable(5)
		if _dade != nil {
			return 0, _dade
		}
		return _aggcb.Decode(_ffbde._caba)
	case 3:
		if _ffbde._cfgd == nil {
			_accab, _cga := _fa.GetStandardTable(0)
			if _cga != nil {
				return 0, _cga
			}
			_ffbde._cfgd = _accab
		}
		return _ffbde._cfgd.Decode(_ffbde._caba)
	}
	return 0, nil
}
