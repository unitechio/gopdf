package segments

import (
	_bc "encoding/binary"
	_ag "errors"
	_be "fmt"
	_b "image"
	_de "io"
	_f "math"
	_bb "strings"
	_a "time"

	_gb "bitbucket.org/shenghui0779/gopdf/common"
	_g "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_cb "bitbucket.org/shenghui0779/gopdf/internal/jbig2/basic"
	_c "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_aa "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder/arithmetic"
	_aag "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder/huffman"
	_ae "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder/mmr"
	_cg "bitbucket.org/shenghui0779/gopdf/internal/jbig2/encoder/arithmetic"
	_ac "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
	_bca "bitbucket.org/shenghui0779/gopdf/internal/jbig2/internal"
	_fc "golang.org/x/xerrors"
)

func (_fcgc *Header) readHeaderLength(_aeac _g.StreamReader, _gfcdda int64) {
	_fcgc.HeaderLength = _aeac.StreamPosition() - _gfcdda
}
func (_egbb *Header) readHeaderFlags() error {
	const _bgc = "\u0072e\u0061d\u0048\u0065\u0061\u0064\u0065\u0072\u0046\u006c\u0061\u0067\u0073"
	_abdb, _gecg := _egbb.Reader.ReadBit()
	if _gecg != nil {
		return _ac.Wrap(_gecg, _bgc, "r\u0065\u0074\u0061\u0069\u006e\u0020\u0066\u006c\u0061\u0067")
	}
	if _abdb != 0 {
		_egbb.RetainFlag = true
	}
	_abdb, _gecg = _egbb.Reader.ReadBit()
	if _gecg != nil {
		return _ac.Wrap(_gecg, _bgc, "\u0070\u0061g\u0065\u0020\u0061s\u0073\u006f\u0063\u0069\u0061\u0074\u0069\u006f\u006e")
	}
	if _abdb != 0 {
		_egbb.PageAssociationFieldSize = true
	}
	_acc, _gecg := _egbb.Reader.ReadBits(6)
	if _gecg != nil {
		return _ac.Wrap(_gecg, _bgc, "\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065")
	}
	_egbb.Type = Type(int(_acc))
	return nil
}
func (_gccg *Header) String() string {
	_eafg := &_bb.Builder{}
	_eafg.WriteString("\u000a[\u0053E\u0047\u004d\u0045\u004e\u0054-\u0048\u0045A\u0044\u0045\u0052\u005d\u000a")
	_eafg.WriteString(_be.Sprintf("\t\u002d\u0020\u0053\u0065gm\u0065n\u0074\u004e\u0075\u006d\u0062e\u0072\u003a\u0020\u0025\u0076\u000a", _gccg.SegmentNumber))
	_eafg.WriteString(_be.Sprintf("\u0009\u002d\u0020T\u0079\u0070\u0065\u003a\u0020\u0025\u0076\u000a", _gccg.Type))
	_eafg.WriteString(_be.Sprintf("\u0009-\u0020R\u0065\u0074\u0061\u0069\u006eF\u006c\u0061g\u003a\u0020\u0025\u0076\u000a", _gccg.RetainFlag))
	_eafg.WriteString(_be.Sprintf("\u0009\u002d\u0020Pa\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _gccg.PageAssociation))
	_eafg.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0050\u0061\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061\u0074i\u006fn\u0046\u0069\u0065\u006c\u0064\u0053\u0069\u007a\u0065\u003a\u0020\u0025\u0076\u000a", _gccg.PageAssociationFieldSize))
	_eafg.WriteString("\u0009-\u0020R\u0054\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u0053\u003a\u000a")
	for _, _bcf := range _gccg.RTSNumbers {
		_eafg.WriteString(_be.Sprintf("\u0009\t\u002d\u0020\u0025\u0064\u000a", _bcf))
	}
	_eafg.WriteString(_be.Sprintf("\t\u002d \u0048\u0065\u0061\u0064\u0065\u0072\u004c\u0065n\u0067\u0074\u0068\u003a %\u0076\u000a", _gccg.HeaderLength))
	_eafg.WriteString(_be.Sprintf("\u0009-\u0020\u0053\u0065\u0067m\u0065\u006e\u0074\u0044\u0061t\u0061L\u0065n\u0067\u0074\u0068\u003a\u0020\u0025\u0076\n", _gccg.SegmentDataLength))
	_eafg.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074D\u0061\u0074\u0061\u0053\u0074\u0061\u0072t\u004f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0076\u000a", _gccg.SegmentDataStartOffset))
	return _eafg.String()
}
func (_def *PageInformationSegment) readResolution() error {
	_fcca, _fdbce := _def._fabba.ReadBits(32)
	if _fdbce != nil {
		return _fdbce
	}
	_def.ResolutionX = int(_fcca & _f.MaxInt32)
	_fcca, _fdbce = _def._fabba.ReadBits(32)
	if _fdbce != nil {
		return _fdbce
	}
	_def.ResolutionY = int(_fcca & _f.MaxInt32)
	return nil
}
func (_bfgd *TextRegion) encodeSymbols(_egacb _g.BinaryWriter) (_acf int, _bfbe error) {
	const _eabeb = "\u0065\u006e\u0063\u006f\u0064\u0065\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	_cefbf := make([]byte, 4)
	_bc.BigEndian.PutUint32(_cefbf, _bfgd.NumberOfSymbols)
	if _acf, _bfbe = _egacb.Write(_cefbf); _bfbe != nil {
		return _acf, _ac.Wrap(_bfbe, _eabeb, "\u004e\u0075\u006dbe\u0072\u004f\u0066\u0053\u0079\u006d\u0062\u006f\u006c\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073")
	}
	_ebefg, _bfbe := _c.NewClassedPoints(_bfgd._edded, _bfgd._babe)
	if _bfbe != nil {
		return 0, _ac.Wrap(_bfbe, _eabeb, "")
	}
	var _gccc, _bgbe int
	_dgbc := _cg.New()
	_dgbc.Init()
	if _bfbe = _dgbc.EncodeInteger(_cg.IADT, 0); _bfbe != nil {
		return _acf, _ac.Wrap(_bfbe, _eabeb, "\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0020\u0044\u0054")
	}
	_bfda, _bfbe := _ebefg.GroupByY()
	if _bfbe != nil {
		return 0, _ac.Wrap(_bfbe, _eabeb, "")
	}
	for _, _acff := range _bfda {
		_gdbb := int(_acff.YAtIndex(0))
		_fgee := _gdbb - _gccc
		if _bfbe = _dgbc.EncodeInteger(_cg.IADT, _fgee); _bfbe != nil {
			return _acf, _ac.Wrap(_bfbe, _eabeb, "")
		}
		var _ddeb int
		for _fcdb, _ecdd := range _acff.IntSlice {
			switch _fcdb {
			case 0:
				_aadbgf := int(_acff.XAtIndex(_fcdb)) - _bgbe
				if _bfbe = _dgbc.EncodeInteger(_cg.IAFS, _aadbgf); _bfbe != nil {
					return _acf, _ac.Wrap(_bfbe, _eabeb, "")
				}
				_bgbe += _aadbgf
				_ddeb = _bgbe
			default:
				_efgaf := int(_acff.XAtIndex(_fcdb)) - _ddeb
				if _bfbe = _dgbc.EncodeInteger(_cg.IADS, _efgaf); _bfbe != nil {
					return _acf, _ac.Wrap(_bfbe, _eabeb, "")
				}
				_ddeb += _efgaf
			}
			_fdbd, _dfaa := _bfgd._dcfb.Get(_ecdd)
			if _dfaa != nil {
				return _acf, _ac.Wrap(_dfaa, _eabeb, "")
			}
			_afbb, _ebfg := _bfgd._afcf[_fdbd]
			if !_ebfg {
				_afbb, _ebfg = _bfgd._dbe[_fdbd]
				if !_ebfg {
					return _acf, _ac.Errorf(_eabeb, "\u0053\u0079\u006d\u0062\u006f\u006c:\u0020\u0027\u0025d\u0027\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064 \u0069\u006e\u0020\u0067\u006cob\u0061\u006c\u0020\u0061\u006e\u0064\u0020\u006c\u006f\u0063\u0061\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0020\u006d\u0061\u0070", _fdbd)
				}
			}
			if _dfaa = _dgbc.EncodeIAID(_bfgd._fadfc, _afbb); _dfaa != nil {
				return _acf, _ac.Wrap(_dfaa, _eabeb, "")
			}
		}
		if _bfbe = _dgbc.EncodeOOB(_cg.IADS); _bfbe != nil {
			return _acf, _ac.Wrap(_bfbe, _eabeb, "")
		}
	}
	_dgbc.Final()
	_ddgf, _bfbe := _dgbc.WriteTo(_egacb)
	if _bfbe != nil {
		return _acf, _ac.Wrap(_bfbe, _eabeb, "")
	}
	_acf += int(_ddgf)
	return _acf, nil
}
func (_ffdd *SymbolDictionary) decodeHeightClassDeltaHeight() (int64, error) {
	if _ffdd.IsHuffmanEncoded {
		return _ffdd.decodeHeightClassDeltaHeightWithHuffman()
	}
	_eafe, _dgca := _ffdd._caaa.DecodeInt(_ffdd._fcbe)
	if _dgca != nil {
		return 0, _dgca
	}
	return int64(_eafe), nil
}

type Regioner interface {
	GetRegionBitmap() (*_c.Bitmap, error)
	GetRegionInfo() *RegionSegment
}

func (_eeaf *SymbolDictionary) readNumberOfExportedSymbols() error {
	_edac, _cgcc := _eeaf._aege.ReadBits(32)
	if _cgcc != nil {
		return _cgcc
	}
	_eeaf.NumberOfExportedSymbols = uint32(_edac & _f.MaxUint32)
	return nil
}
func (_edd *GenericRegion) copyLineAbove(_aba int) error {
	_acbc := _aba * _edd.Bitmap.RowStride
	_gfcd := _acbc - _edd.Bitmap.RowStride
	for _feac := 0; _feac < _edd.Bitmap.RowStride; _feac++ {
		_cbcg, _gdd := _edd.Bitmap.GetByte(_gfcd)
		if _gdd != nil {
			return _gdd
		}
		_gfcd++
		if _gdd = _edd.Bitmap.SetByte(_acbc, _cbcg); _gdd != nil {
			return _gdd
		}
		_acbc++
	}
	return nil
}
func (_eegc *PageInformationSegment) readContainsRefinement() error {
	_ffaf, _dcfc := _eegc._fabba.ReadBit()
	if _dcfc != nil {
		return _dcfc
	}
	if _ffaf == 1 {
		_eegc._adfb = true
	}
	return nil
}
func (_eddf *PageInformationSegment) readMaxStripeSize() error {
	_gbfc, _dcg := _eddf._fabba.ReadBits(15)
	if _dcg != nil {
		return _dcg
	}
	_eddf.MaxStripeSize = uint16(_gbfc & _f.MaxUint16)
	return nil
}
func (_ebe *GenericRefinementRegion) setParameters(_fec *_aa.DecoderStats, _aaab *_aa.Decoder, _fcbd int8, _dca, _cgb uint32, _bbec *_c.Bitmap, _bfde, _gff int32, _aeee bool, _aagc []int8, _dac []int8) {
	_gb.Log.Trace("\u005b\u0047\u0045NE\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052E\u0047I\u004fN\u005d \u0073\u0065\u0074\u0050\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	if _fec != nil {
		_ebe._ef = _fec
	}
	if _aaab != nil {
		_ebe._fb = _aaab
	}
	_ebe.TemplateID = _fcbd
	_ebe.RegionInfo.BitmapWidth = _dca
	_ebe.RegionInfo.BitmapHeight = _cgb
	_ebe.ReferenceBitmap = _bbec
	_ebe.ReferenceDX = _bfde
	_ebe.ReferenceDY = _gff
	_ebe.IsTPGROn = _aeee
	_ebe.GrAtX = _aagc
	_ebe.GrAtY = _dac
	_ebe.RegionBitmap = nil
	_gb.Log.Trace("[\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052E\u0046\u002d\u0052\u0045\u0047\u0049\u004fN]\u0020\u0073\u0065\u0074P\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073 f\u0069\u006ei\u0073\u0068\u0065\u0064\u002e\u0020\u0025\u0073", _ebe)
}
func (_cgbe *SymbolDictionary) readAtPixels(_acdg int) error {
	_cgbe.SdATX = make([]int8, _acdg)
	_cgbe.SdATY = make([]int8, _acdg)
	var (
		_efe   byte
		_cebab error
	)
	for _dffg := 0; _dffg < _acdg; _dffg++ {
		_efe, _cebab = _cgbe._aege.ReadByte()
		if _cebab != nil {
			return _cebab
		}
		_cgbe.SdATX[_dffg] = int8(_efe)
		_efe, _cebab = _cgbe._aege.ReadByte()
		if _cebab != nil {
			return _cebab
		}
		_cgbe.SdATY[_dffg] = int8(_efe)
	}
	return nil
}
func (_gabf *PatternDictionary) setGbAtPixels() {
	if _gabf.HDTemplate == 0 {
		_gabf.GBAtX = make([]int8, 4)
		_gabf.GBAtY = make([]int8, 4)
		_gabf.GBAtX[0] = -int8(_gabf.HdpWidth)
		_gabf.GBAtY[0] = 0
		_gabf.GBAtX[1] = -3
		_gabf.GBAtY[1] = -1
		_gabf.GBAtX[2] = 2
		_gabf.GBAtY[2] = -2
		_gabf.GBAtX[3] = -2
		_gabf.GBAtY[3] = -2
	} else {
		_gabf.GBAtX = []int8{-int8(_gabf.HdpWidth)}
		_gabf.GBAtY = []int8{0}
	}
}
func (_begd *HalftoneRegion) GetRegionBitmap() (*_c.Bitmap, error) {
	if _begd.HalftoneRegionBitmap != nil {
		return _begd.HalftoneRegionBitmap, nil
	}
	var _bgae error
	_begd.HalftoneRegionBitmap = _c.New(int(_begd.RegionSegment.BitmapWidth), int(_begd.RegionSegment.BitmapHeight))
	if _begd.Patterns == nil || len(_begd.Patterns) == 0 {
		_begd.Patterns, _bgae = _begd.GetPatterns()
		if _bgae != nil {
			return nil, _bgae
		}
	}
	if _begd.HDefaultPixel == 1 {
		_begd.HalftoneRegionBitmap.SetDefaultPixel()
	}
	_egb := _f.Ceil(_f.Log(float64(len(_begd.Patterns))) / _f.Log(2))
	_gbaf := int(_egb)
	var _afe [][]int
	_afe, _bgae = _begd.grayScaleDecoding(_gbaf)
	if _bgae != nil {
		return nil, _bgae
	}
	if _bgae = _begd.renderPattern(_afe); _bgae != nil {
		return nil, _bgae
	}
	return _begd.HalftoneRegionBitmap, nil
}
func (_gabb *TextRegion) decodeRdx() (int64, error) {
	const _faeb = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064x"
	if _gabb.IsHuffmanEncoded {
		if _gabb.SbHuffRDX == 3 {
			if _gabb._aafb == nil {
				var (
					_bbeca int
					_cfaf  error
				)
				if _gabb.SbHuffFS == 3 {
					_bbeca++
				}
				if _gabb.SbHuffDS == 3 {
					_bbeca++
				}
				if _gabb.SbHuffDT == 3 {
					_bbeca++
				}
				if _gabb.SbHuffRDWidth == 3 {
					_bbeca++
				}
				if _gabb.SbHuffRDHeight == 3 {
					_bbeca++
				}
				_gabb._aafb, _cfaf = _gabb.getUserTable(_bbeca)
				if _cfaf != nil {
					return 0, _ac.Wrap(_cfaf, _faeb, "")
				}
			}
			return _gabb._aafb.Decode(_gabb._ceegd)
		}
		_aeca, _abff := _aag.GetStandardTable(14 + int(_gabb.SbHuffRDX))
		if _abff != nil {
			return 0, _ac.Wrap(_abff, _faeb, "")
		}
		return _aeca.Decode(_gabb._ceegd)
	}
	_cceag, _gcgg := _gabb._bbcea.DecodeInt(_gabb._bfce)
	if _gcgg != nil {
		return 0, _ac.Wrap(_gcgg, _faeb, "")
	}
	return int64(_cceag), nil
}
func (_dedf *TextRegion) encodeFlags(_abcc _g.BinaryWriter) (_egfg int, _gcgec error) {
	const _bggdd = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _gcgec = _abcc.WriteBit(int(_dedf.SbrTemplate)); _gcgec != nil {
		return _egfg, _ac.Wrap(_gcgec, _bggdd, "s\u0062\u0072\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	if _, _gcgec = _abcc.WriteBits(uint64(_dedf.SbDsOffset), 5); _gcgec != nil {
		return _egfg, _ac.Wrap(_gcgec, _bggdd, "\u0073\u0062\u0044\u0073\u004f\u0066\u0066\u0073\u0065\u0074")
	}
	if _gcgec = _abcc.WriteBit(int(_dedf.DefaultPixel)); _gcgec != nil {
		return _egfg, _ac.Wrap(_gcgec, _bggdd, "\u0044\u0065\u0066a\u0075\u006c\u0074\u0050\u0069\u0078\u0065\u006c")
	}
	if _, _gcgec = _abcc.WriteBits(uint64(_dedf.CombinationOperator), 2); _gcgec != nil {
		return _egfg, _ac.Wrap(_gcgec, _bggdd, "\u0043\u006f\u006d\u0062in\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	if _gcgec = _abcc.WriteBit(int(_dedf.IsTransposed)); _gcgec != nil {
		return _egfg, _ac.Wrap(_gcgec, _bggdd, "\u0069\u0073\u0020\u0074\u0072\u0061\u006e\u0073\u0070\u006f\u0073\u0065\u0064")
	}
	if _, _gcgec = _abcc.WriteBits(uint64(_dedf.ReferenceCorner), 2); _gcgec != nil {
		return _egfg, _ac.Wrap(_gcgec, _bggdd, "\u0072\u0065f\u0065\u0072\u0065n\u0063\u0065\u0020\u0063\u006f\u0072\u006e\u0065\u0072")
	}
	if _, _gcgec = _abcc.WriteBits(uint64(_dedf.LogSBStrips), 2); _gcgec != nil {
		return _egfg, _ac.Wrap(_gcgec, _bggdd, "L\u006f\u0067\u0053\u0042\u0053\u0074\u0072\u0069\u0070\u0073")
	}
	var _gffe int
	if _dedf.UseRefinement {
		_gffe = 1
	}
	if _gcgec = _abcc.WriteBit(_gffe); _gcgec != nil {
		return _egfg, _ac.Wrap(_gcgec, _bggdd, "\u0075\u0073\u0065\u0020\u0072\u0065\u0066\u0069\u006ee\u006d\u0065\u006e\u0074")
	}
	_gffe = 0
	if _dedf.IsHuffmanEncoded {
		_gffe = 1
	}
	if _gcgec = _abcc.WriteBit(_gffe); _gcgec != nil {
		return _egfg, _ac.Wrap(_gcgec, _bggdd, "u\u0073\u0065\u0020\u0068\u0075\u0066\u0066\u006d\u0061\u006e")
	}
	_egfg = 2
	return _egfg, nil
}
func (_eeaeg *TextRegion) setParameters(_dcbdb *_aa.Decoder, _ada, _eadda bool, _gedbb, _fccdd uint32, _cgff uint32, _daaef int8, _deda uint32, _gdge int8, _gedg _c.CombinationOperator, _cafd int8, _gfcc int16, _efed, _bdfc, _gdgf, _ade, _abcb, _dcaef, _cfba, _cgcce, _agdc, _acffg int8, _cfad, _faad []int8, _dedcc []*_c.Bitmap, _eebd int8) {
	_eeaeg._bbcea = _dcbdb
	_eeaeg.IsHuffmanEncoded = _ada
	_eeaeg.UseRefinement = _eadda
	_eeaeg.RegionInfo.BitmapWidth = _gedbb
	_eeaeg.RegionInfo.BitmapHeight = _fccdd
	_eeaeg.NumberOfSymbolInstances = _cgff
	_eeaeg.SbStrips = _daaef
	_eeaeg.NumberOfSymbols = _deda
	_eeaeg.DefaultPixel = _gdge
	_eeaeg.CombinationOperator = _gedg
	_eeaeg.IsTransposed = _cafd
	_eeaeg.ReferenceCorner = _gfcc
	_eeaeg.SbDsOffset = _efed
	_eeaeg.SbHuffFS = _bdfc
	_eeaeg.SbHuffDS = _gdgf
	_eeaeg.SbHuffDT = _ade
	_eeaeg.SbHuffRDWidth = _abcb
	_eeaeg.SbHuffRDHeight = _dcaef
	_eeaeg.SbHuffRSize = _agdc
	_eeaeg.SbHuffRDX = _cfba
	_eeaeg.SbHuffRDY = _cgcce
	_eeaeg.SbrTemplate = _acffg
	_eeaeg.SbrATX = _cfad
	_eeaeg.SbrATY = _faad
	_eeaeg.Symbols = _dedcc
	_eeaeg._bfa = _eebd
}
func (_bgfb *TextRegion) blit(_afcce *_c.Bitmap, _bbdb int64) error {
	if _bgfb.IsTransposed == 0 && (_bgfb.ReferenceCorner == 2 || _bgfb.ReferenceCorner == 3) {
		_bgfb._dgee += int64(_afcce.Width - 1)
	} else if _bgfb.IsTransposed == 1 && (_bgfb.ReferenceCorner == 0 || _bgfb.ReferenceCorner == 2) {
		_bgfb._dgee += int64(_afcce.Height - 1)
	}
	_daea := _bgfb._dgee
	if _bgfb.IsTransposed == 1 {
		_daea, _bbdb = _bbdb, _daea
	}
	switch _bgfb.ReferenceCorner {
	case 0:
		_bbdb -= int64(_afcce.Height - 1)
	case 2:
		_bbdb -= int64(_afcce.Height - 1)
		_daea -= int64(_afcce.Width - 1)
	case 3:
		_daea -= int64(_afcce.Width - 1)
	}
	_bbcf := _c.Blit(_afcce, _bgfb.RegionBitmap, int(_daea), int(_bbdb), _bgfb.CombinationOperator)
	if _bbcf != nil {
		return _bbcf
	}
	if _bgfb.IsTransposed == 0 && (_bgfb.ReferenceCorner == 0 || _bgfb.ReferenceCorner == 1) {
		_bgfb._dgee += int64(_afcce.Width - 1)
	}
	if _bgfb.IsTransposed == 1 && (_bgfb.ReferenceCorner == 1 || _bgfb.ReferenceCorner == 3) {
		_bgfb._dgee += int64(_afcce.Height - 1)
	}
	return nil
}
func (_dedc *HalftoneRegion) parseHeader() error {
	if _egfd := _dedc.RegionSegment.parseHeader(); _egfd != nil {
		return _egfd
	}
	_abdf, _ecbe := _dedc._ebf.ReadBit()
	if _ecbe != nil {
		return _ecbe
	}
	_dedc.HDefaultPixel = int8(_abdf)
	_abg, _ecbe := _dedc._ebf.ReadBits(3)
	if _ecbe != nil {
		return _ecbe
	}
	_dedc.CombinationOperator = _c.CombinationOperator(_abg & 0xf)
	_abdf, _ecbe = _dedc._ebf.ReadBit()
	if _ecbe != nil {
		return _ecbe
	}
	if _abdf == 1 {
		_dedc.HSkipEnabled = true
	}
	_abg, _ecbe = _dedc._ebf.ReadBits(2)
	if _ecbe != nil {
		return _ecbe
	}
	_dedc.HTemplate = byte(_abg & 0xf)
	_abdf, _ecbe = _dedc._ebf.ReadBit()
	if _ecbe != nil {
		return _ecbe
	}
	if _abdf == 1 {
		_dedc.IsMMREncoded = true
	}
	_abg, _ecbe = _dedc._ebf.ReadBits(32)
	if _ecbe != nil {
		return _ecbe
	}
	_dedc.HGridWidth = uint32(_abg & _f.MaxUint32)
	_abg, _ecbe = _dedc._ebf.ReadBits(32)
	if _ecbe != nil {
		return _ecbe
	}
	_dedc.HGridHeight = uint32(_abg & _f.MaxUint32)
	_abg, _ecbe = _dedc._ebf.ReadBits(32)
	if _ecbe != nil {
		return _ecbe
	}
	_dedc.HGridX = int32(_abg & _f.MaxInt32)
	_abg, _ecbe = _dedc._ebf.ReadBits(32)
	if _ecbe != nil {
		return _ecbe
	}
	_dedc.HGridY = int32(_abg & _f.MaxInt32)
	_abg, _ecbe = _dedc._ebf.ReadBits(16)
	if _ecbe != nil {
		return _ecbe
	}
	_dedc.HRegionX = uint16(_abg & _f.MaxUint16)
	_abg, _ecbe = _dedc._ebf.ReadBits(16)
	if _ecbe != nil {
		return _ecbe
	}
	_dedc.HRegionY = uint16(_abg & _f.MaxUint16)
	if _ecbe = _dedc.computeSegmentDataStructure(); _ecbe != nil {
		return _ecbe
	}
	return _dedc.checkInput()
}
func (_dgffb *TextRegion) GetRegionInfo() *RegionSegment { return _dgffb.RegionInfo }

type HalftoneRegion struct {
	_ebf                 _g.StreamReader
	_bdeg                *Header
	DataHeaderOffset     int64
	DataHeaderLength     int64
	DataOffset           int64
	DataLength           int64
	RegionSegment        *RegionSegment
	HDefaultPixel        int8
	CombinationOperator  _c.CombinationOperator
	HSkipEnabled         bool
	HTemplate            byte
	IsMMREncoded         bool
	HGridWidth           uint32
	HGridHeight          uint32
	HGridX               int32
	HGridY               int32
	HRegionX             uint16
	HRegionY             uint16
	HalftoneRegionBitmap *_c.Bitmap
	Patterns             []*_c.Bitmap
}
type TextRegion struct {
	_ceegd                  _g.StreamReader
	RegionInfo              *RegionSegment
	SbrTemplate             int8
	SbDsOffset              int8
	DefaultPixel            int8
	CombinationOperator     _c.CombinationOperator
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
	_dgee                   int64
	SbStrips                int8
	NumberOfSymbols         uint32
	RegionBitmap            *_c.Bitmap
	Symbols                 []*_c.Bitmap
	_bbcea                  *_aa.Decoder
	_baea                   *GenericRefinementRegion
	_fdfbf                  *_aa.DecoderStats
	_dfge                   *_aa.DecoderStats
	_bcca                   *_aa.DecoderStats
	_gdgaa                  *_aa.DecoderStats
	_feaceb                 *_aa.DecoderStats
	_bbgdf                  *_aa.DecoderStats
	_agag                   *_aa.DecoderStats
	_bcfb                   *_aa.DecoderStats
	_bfce                   *_aa.DecoderStats
	_badb                   *_aa.DecoderStats
	_gfd                    *_aa.DecoderStats
	_bfa                    int8
	_gaead                  *_aag.FixedSizeTable
	Header                  *Header
	_cdcga                  _aag.Tabler
	_eabed                  _aag.Tabler
	_afeb                   _aag.Tabler
	_cdcgd                  _aag.Tabler
	_agba                   _aag.Tabler
	_aafb                   _aag.Tabler
	_ecgf                   _aag.Tabler
	_ddfff                  _aag.Tabler
	_afcf, _dbe             map[int]int
	_babe                   []int
	_edded                  *_c.Points
	_ebde                   *_c.Bitmaps
	_dcfb                   *_cb.IntSlice
	_afbf, _fadfc           int
	_bfab                   *_c.Boxes
}

func (_afff *Header) writeSegmentDataLength(_agbc _g.BinaryWriter) (_efae int, _cgce error) {
	_bgbb := make([]byte, 4)
	_bc.BigEndian.PutUint32(_bgbb, uint32(_afff.SegmentDataLength))
	if _efae, _cgce = _agbc.Write(_bgbb); _cgce != nil {
		return 0, _ac.Wrap(_cgce, "\u0048\u0065a\u0064\u0065\u0072\u002e\u0077\u0072\u0069\u0074\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0044\u0061\u0074\u0061\u004c\u0065ng\u0074\u0068", "")
	}
	return _efae, nil
}
func (_dgdg *HalftoneRegion) grayScaleDecoding(_aedf int) ([][]int, error) {
	var (
		_cgae []int8
		_eabd []int8
	)
	if !_dgdg.IsMMREncoded {
		_cgae = make([]int8, 4)
		_eabd = make([]int8, 4)
		if _dgdg.HTemplate <= 1 {
			_cgae[0] = 3
		} else if _dgdg.HTemplate >= 2 {
			_cgae[0] = 2
		}
		_eabd[0] = -1
		_cgae[1] = -3
		_eabd[1] = -1
		_cgae[2] = 2
		_eabd[2] = -2
		_cgae[3] = -2
		_eabd[3] = -2
	}
	_cbea := make([]*_c.Bitmap, _aedf)
	_baf := NewGenericRegion(_dgdg._ebf)
	_baf.setParametersMMR(_dgdg.IsMMREncoded, _dgdg.DataOffset, _dgdg.DataLength, _dgdg.HGridHeight, _dgdg.HGridWidth, _dgdg.HTemplate, false, _dgdg.HSkipEnabled, _cgae, _eabd)
	_gccb := _aedf - 1
	var _dgbf error
	_cbea[_gccb], _dgbf = _baf.GetRegionBitmap()
	if _dgbf != nil {
		return nil, _dgbf
	}
	for _gccb > 0 {
		_gccb--
		_baf.Bitmap = nil
		_cbea[_gccb], _dgbf = _baf.GetRegionBitmap()
		if _dgbf != nil {
			return nil, _dgbf
		}
		if _dgbf = _dgdg.combineGrayscalePlanes(_cbea, _gccb); _dgbf != nil {
			return nil, _dgbf
		}
	}
	return _dgdg.computeGrayScalePlanes(_cbea, _aedf)
}
func (_cfcgg *TextRegion) decodeStripT() (_baeb int64, _ebac error) {
	if _cfcgg.IsHuffmanEncoded {
		if _cfcgg.SbHuffDT == 3 {
			if _cfcgg._afeb == nil {
				var _bbfg int
				if _cfcgg.SbHuffFS == 3 {
					_bbfg++
				}
				if _cfcgg.SbHuffDS == 3 {
					_bbfg++
				}
				_cfcgg._afeb, _ebac = _cfcgg.getUserTable(_bbfg)
				if _ebac != nil {
					return 0, _ebac
				}
			}
			_baeb, _ebac = _cfcgg._afeb.Decode(_cfcgg._ceegd)
			if _ebac != nil {
				return 0, _ebac
			}
		} else {
			var _gcdf _aag.Tabler
			_gcdf, _ebac = _aag.GetStandardTable(11 + int(_cfcgg.SbHuffDT))
			if _ebac != nil {
				return 0, _ebac
			}
			_baeb, _ebac = _gcdf.Decode(_cfcgg._ceegd)
			if _ebac != nil {
				return 0, _ebac
			}
		}
	} else {
		var _abag int32
		_abag, _ebac = _cfcgg._bbcea.DecodeInt(_cfcgg._fdfbf)
		if _ebac != nil {
			return 0, _ebac
		}
		_baeb = int64(_abag)
	}
	_baeb *= int64(-_cfcgg.SbStrips)
	return _baeb, nil
}
func (_fcbb *PageInformationSegment) readCombinationOperator() error {
	_fdbf, _geece := _fcbb._fabba.ReadBits(2)
	if _geece != nil {
		return _geece
	}
	_fcbb._cba = _c.CombinationOperator(int(_fdbf))
	return nil
}
func (_ffda *GenericRegion) Size() int { return _ffda.RegionSegment.Size() + 1 + 2*len(_ffda.GBAtX) }

type template1 struct{}

func (_gbgd *RegionSegment) readCombinationOperator() error {
	_edaf, _effa := _gbgd._ffade.ReadBits(3)
	if _effa != nil {
		return _effa
	}
	_gbgd.CombinaionOperator = _c.CombinationOperator(_edaf & 0xF)
	return nil
}
func (_ebdc *Header) writeSegmentNumber(_eaade _g.BinaryWriter) (_aebbd int, _dfc error) {
	_ffbg := make([]byte, 4)
	_bc.BigEndian.PutUint32(_ffbg, _ebdc.SegmentNumber)
	if _aebbd, _dfc = _eaade.Write(_ffbg); _dfc != nil {
		return 0, _ac.Wrap(_dfc, "\u0048e\u0061\u0064\u0065\u0072.\u0077\u0072\u0069\u0074\u0065S\u0065g\u006de\u006e\u0074\u004e\u0075\u006d\u0062\u0065r", "")
	}
	return _aebbd, nil
}
func (_afb *SymbolDictionary) checkInput() error {
	if _afb.SdHuffDecodeHeightSelection == 2 {
		_gb.Log.Debug("\u0053\u0079\u006d\u0062\u006fl\u0020\u0044\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0048\u0065\u0069\u0067\u0068\u0074\u0020\u0053e\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0070\u0065r\u006d\u0069\u0074\u0074\u0065\u0064", _afb.SdHuffDecodeHeightSelection)
	}
	if _afb.SdHuffDecodeWidthSelection == 2 {
		_gb.Log.Debug("\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0057\u0069\u0064t\u0068\u0020\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064\u0020\u0076\u0061l\u0075\u0065\u0020\u006e\u006f\u0074 \u0070\u0065r\u006d\u0069t\u0074e\u0064", _afb.SdHuffDecodeWidthSelection)
	}
	if _afb.IsHuffmanEncoded {
		if _afb.SdTemplate != 0 {
			_gb.Log.Debug("\u0053\u0044T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0028\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e \u0030\u0029", _afb.SdTemplate)
		}
		if !_afb.UseRefinementAggregation {
			if !_afb.UseRefinementAggregation {
				if _afb._bebbf {
					_gb.Log.Debug("\u0049\u0073\u0043\u006f\u0064\u0069\u006e\u0067C\u006f\u006e\u0074ex\u0074\u0052\u0065\u0074\u0061\u0069n\u0065\u0064\u0020\u003d\u0020\u0074\u0072\u0075\u0065\u0020\u0028\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0066\u0061\u006cs\u0065\u0029")
					_afb._bebbf = false
				}
				if _afb._bcce {
					_gb.Log.Debug("\u0069s\u0043\u006fd\u0069\u006e\u0067\u0043o\u006e\u0074\u0065x\u0074\u0055\u0073\u0065\u0064\u0020\u003d\u0020\u0074ru\u0065\u0020\u0028s\u0068\u006fu\u006c\u0064\u0020\u0062\u0065\u0020f\u0061\u006cs\u0065\u0029")
					_afb._bcce = false
				}
			}
		}
	} else {
		if _afb.SdHuffBMSizeSelection != 0 {
			_gb.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066B\u004d\u0053\u0069\u007a\u0065\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0030")
			_afb.SdHuffBMSizeSelection = 0
		}
		if _afb.SdHuffDecodeWidthSelection != 0 {
			_gb.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0057\u0069\u0064\u0074\u0068\u0053\u0065\u006c\u0065\u0063\u0074\u0069o\u006e\u0020\u0073\u0068\u006fu\u006c\u0064 \u0062\u0065\u0020\u0030")
			_afb.SdHuffDecodeWidthSelection = 0
		}
		if _afb.SdHuffDecodeHeightSelection != 0 {
			_gb.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0048e\u0069\u0067\u0068\u0074\u0053\u0065l\u0065\u0063\u0074\u0069\u006f\u006e\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0030")
			_afb.SdHuffDecodeHeightSelection = 0
		}
	}
	if !_afb.UseRefinementAggregation {
		if _afb.SdrTemplate != 0 {
			_gb.Log.Debug("\u0053\u0044\u0052\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u003d\u0020\u0025\u0064\u0020\u0028s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030\u0029", _afb.SdrTemplate)
			_afb.SdrTemplate = 0
		}
	}
	if !_afb.IsHuffmanEncoded || !_afb.UseRefinementAggregation {
		if _afb.SdHuffAggInstanceSelection {
			_gb.Log.Debug("\u0053d\u0048\u0075f\u0066\u0041\u0067g\u0049\u006e\u0073\u0074\u0061\u006e\u0063e\u0053\u0065\u006c\u0065\u0063\u0074i\u006f\u006e\u0020\u003d\u0020\u0025\u0064\u0020\u0028\u0073\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030\u0029", _afb.SdHuffAggInstanceSelection)
		}
	}
	return nil
}

var (
	_ Regioner  = &TextRegion{}
	_ Segmenter = &TextRegion{}
)

func (_bdb *GenericRefinementRegion) overrideAtTemplate0(_cgaa, _fbda, _cdef, _fcb, _baec int) int {
	if _bdb._gbd[0] {
		_cgaa &= 0xfff7
		if _bdb.GrAtY[0] == 0 && int(_bdb.GrAtX[0]) >= -_baec {
			_cgaa |= (_fcb >> uint(7-(_baec+int(_bdb.GrAtX[0]))) & 0x1) << 3
		} else {
			_cgaa |= _bdb.getPixel(_bdb.RegionBitmap, _fbda+int(_bdb.GrAtX[0]), _cdef+int(_bdb.GrAtY[0])) << 3
		}
	}
	if _bdb._gbd[1] {
		_cgaa &= 0xefff
		if _bdb.GrAtY[1] == 0 && int(_bdb.GrAtX[1]) >= -_baec {
			_cgaa |= (_fcb >> uint(7-(_baec+int(_bdb.GrAtX[1]))) & 0x1) << 12
		} else {
			_cgaa |= _bdb.getPixel(_bdb.ReferenceBitmap, _fbda+int(_bdb.GrAtX[1]), _cdef+int(_bdb.GrAtY[1]))
		}
	}
	return _cgaa
}
func (_ced *SymbolDictionary) getUserTable(_ffgad int) (_aag.Tabler, error) {
	var _dgdd int
	for _, _degg := range _ced.Header.RTSegments {
		if _degg.Type == 53 {
			if _dgdd == _ffgad {
				_decd, _dacc := _degg.GetSegmentData()
				if _dacc != nil {
					return nil, _dacc
				}
				_fac := _decd.(_aag.BasicTabler)
				return _aag.NewEncodedTable(_fac)
			}
			_dgdd++
		}
	}
	return nil, nil
}
func (_caff *PageInformationSegment) CombinationOperator() _c.CombinationOperator { return _caff._cba }
func (_cfbe *GenericRegion) decodeTemplate0a(_ffgb, _gagc, _gbag int, _ecfg, _eece int) (_gea error) {
	const _bff = "\u0064\u0065c\u006f\u0064\u0065T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0030\u0061"
	var (
		_debg, _ffe int
		_aafc, _cbf int
		_aae        byte
		_ddaf       int
	)
	if _ffgb >= 1 {
		_aae, _gea = _cfbe.Bitmap.GetByte(_eece)
		if _gea != nil {
			return _ac.Wrap(_gea, _bff, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_aafc = int(_aae)
	}
	if _ffgb >= 2 {
		_aae, _gea = _cfbe.Bitmap.GetByte(_eece - _cfbe.Bitmap.RowStride)
		if _gea != nil {
			return _ac.Wrap(_gea, _bff, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_cbf = int(_aae) << 6
	}
	_debg = (_aafc & 0xf0) | (_cbf & 0x3800)
	for _bfff := 0; _bfff < _gbag; _bfff = _ddaf {
		var (
			_geaa byte
			_eda  int
		)
		_ddaf = _bfff + 8
		if _egcf := _gagc - _bfff; _egcf > 8 {
			_eda = 8
		} else {
			_eda = _egcf
		}
		if _ffgb > 0 {
			_aafc <<= 8
			if _ddaf < _gagc {
				_aae, _gea = _cfbe.Bitmap.GetByte(_eece + 1)
				if _gea != nil {
					return _ac.Wrap(_gea, _bff, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_aafc |= int(_aae)
			}
		}
		if _ffgb > 1 {
			_gbde := _eece - _cfbe.Bitmap.RowStride + 1
			_cbf <<= 8
			if _ddaf < _gagc {
				_aae, _gea = _cfbe.Bitmap.GetByte(_gbde)
				if _gea != nil {
					return _ac.Wrap(_gea, _bff, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_cbf |= int(_aae) << 6
			} else {
				_cbf |= 0
			}
		}
		for _ddeg := 0; _ddeg < _eda; _ddeg++ {
			_cacd := uint(7 - _ddeg)
			if _cfbe._aaaf {
				_ffe = _cfbe.overrideAtTemplate0a(_debg, _bfff+_ddeg, _ffgb, int(_geaa), _ddeg, int(_cacd))
				_cfbe._bcae.SetIndex(int32(_ffe))
			} else {
				_cfbe._bcae.SetIndex(int32(_debg))
			}
			var _fdd int
			_fdd, _gea = _cfbe._ga.DecodeBit(_cfbe._bcae)
			if _gea != nil {
				return _ac.Wrap(_gea, _bff, "")
			}
			_geaa |= byte(_fdd) << _cacd
			_debg = ((_debg & 0x7bf7) << 1) | _fdd | ((_aafc >> _cacd) & 0x10) | ((_cbf >> _cacd) & 0x800)
		}
		if _dgf := _cfbe.Bitmap.SetByte(_ecfg, _geaa); _dgf != nil {
			return _ac.Wrap(_dgf, _bff, "")
		}
		_ecfg++
		_eece++
	}
	return nil
}
func (_fcgb *TextRegion) readUseRefinement() error {
	if !_fcgb.UseRefinement || _fcgb.SbrTemplate != 0 {
		return nil
	}
	var (
		_cgba byte
		_ccda error
	)
	_fcgb.SbrATX = make([]int8, 2)
	_fcgb.SbrATY = make([]int8, 2)
	_cgba, _ccda = _fcgb._ceegd.ReadByte()
	if _ccda != nil {
		return _ccda
	}
	_fcgb.SbrATX[0] = int8(_cgba)
	_cgba, _ccda = _fcgb._ceegd.ReadByte()
	if _ccda != nil {
		return _ccda
	}
	_fcgb.SbrATY[0] = int8(_cgba)
	_cgba, _ccda = _fcgb._ceegd.ReadByte()
	if _ccda != nil {
		return _ccda
	}
	_fcgb.SbrATX[1] = int8(_cgba)
	_cgba, _ccda = _fcgb._ceegd.ReadByte()
	if _ccda != nil {
		return _ccda
	}
	_fcgb.SbrATY[1] = int8(_cgba)
	return nil
}
func (_ebagd *SymbolDictionary) readNumberOfNewSymbols() error {
	_ddcbcc, _aedfe := _ebagd._aege.ReadBits(32)
	if _aedfe != nil {
		return _aedfe
	}
	_ebagd.NumberOfNewSymbols = uint32(_ddcbcc & _f.MaxUint32)
	return nil
}
func (_dbaa *SymbolDictionary) encodeATFlags(_aede _g.BinaryWriter) (_bcge int, _ggca error) {
	const _cgadg = "\u0065\u006e\u0063\u006f\u0064\u0065\u0041\u0054\u0046\u006c\u0061\u0067\u0073"
	if _dbaa.IsHuffmanEncoded || _dbaa.SdTemplate != 0 {
		return 0, nil
	}
	_dcgce := 4
	if _dbaa.SdTemplate != 0 {
		_dcgce = 1
	}
	for _bgd := 0; _bgd < _dcgce; _bgd++ {
		if _ggca = _aede.WriteByte(byte(_dbaa.SdATX[_bgd])); _ggca != nil {
			return _bcge, _ac.Wrapf(_ggca, _cgadg, "\u0053d\u0041\u0054\u0058\u005b\u0025\u0064]", _bgd)
		}
		_bcge++
		if _ggca = _aede.WriteByte(byte(_dbaa.SdATY[_bgd])); _ggca != nil {
			return _bcge, _ac.Wrapf(_ggca, _cgadg, "\u0053d\u0041\u0054\u0059\u005b\u0025\u0064]", _bgd)
		}
		_bcge++
	}
	return _bcge, nil
}
func (_caec *Header) readDataStartOffset(_bffd _g.StreamReader, _fffe OrganizationType) {
	if _fffe == OSequential {
		_caec.SegmentDataStartOffset = uint64(_bffd.StreamPosition())
	}
}
func (_aged *GenericRegion) decodeLine(_beg, _gag, _gcg int) error {
	const _fgdg = "\u0064\u0065\u0063\u006f\u0064\u0065\u004c\u0069\u006e\u0065"
	_dgg := _aged.Bitmap.GetByteIndex(0, _beg)
	_gbb := _dgg - _aged.Bitmap.RowStride
	switch _aged.GBTemplate {
	case 0:
		if !_aged.UseExtTemplates {
			return _aged.decodeTemplate0a(_beg, _gag, _gcg, _dgg, _gbb)
		}
		return _aged.decodeTemplate0b(_beg, _gag, _gcg, _dgg, _gbb)
	case 1:
		return _aged.decodeTemplate1(_beg, _gag, _gcg, _dgg, _gbb)
	case 2:
		return _aged.decodeTemplate2(_beg, _gag, _gcg, _dgg, _gbb)
	case 3:
		return _aged.decodeTemplate3(_beg, _gag, _gcg, _dgg, _gbb)
	}
	return _ac.Errorf(_fgdg, "\u0069\u006e\u0076a\u006c\u0069\u0064\u0020G\u0042\u0054\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u003a\u0020\u0025\u0064", _aged.GBTemplate)
}
func (_ccbd *SymbolDictionary) GetDictionary() ([]*_c.Bitmap, error) {
	_gb.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c-\u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		_gb.Log.Trace("\u005b\u0053\u0059M\u0042\u004f\u004c\u002d\u0044\u0049\u0043\u0054\u0049\u004f\u004e\u0041\u0052\u0059\u005d\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
		_gb.Log.Trace("\u005b\u0053Y\u004d\u0042\u004f\u004c\u002dD\u0049\u0043\u0054\u0049\u004fN\u0041\u0052\u0059\u005d\u0020\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0020\u000a\u0045\u0078\u003a\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u000a\u006e\u0065\u0077\u003a\u0027\u0025\u0073\u0027", _ccbd._abab, _ccbd._gbeb)
	}()
	if _ccbd._abab == nil {
		var _gacde error
		if _ccbd.UseRefinementAggregation {
			_ccbd._ebafc = _ccbd.getSbSymCodeLen()
		}
		if !_ccbd.IsHuffmanEncoded {
			if _gacde = _ccbd.setCodingStatistics(); _gacde != nil {
				return nil, _gacde
			}
		}
		_ccbd._gbeb = make([]*_c.Bitmap, _ccbd.NumberOfNewSymbols)
		var _fcgf []int
		if _ccbd.IsHuffmanEncoded && !_ccbd.UseRefinementAggregation {
			_fcgf = make([]int, _ccbd.NumberOfNewSymbols)
		}
		if _gacde = _ccbd.setSymbolsArray(); _gacde != nil {
			return nil, _gacde
		}
		var _dcbd, _ffffe int64
		_ccbd._geba = 0
		for _ccbd._geba < _ccbd.NumberOfNewSymbols {
			_ffffe, _gacde = _ccbd.decodeHeightClassDeltaHeight()
			if _gacde != nil {
				return nil, _gacde
			}
			_dcbd += _ffffe
			var _gggc, _agef uint32
			_bccc := int64(_ccbd._geba)
			for {
				var _fbgc int64
				_fbgc, _gacde = _ccbd.decodeDifferenceWidth()
				if _fc.Is(_gacde, _bca.ErrOOB) {
					break
				}
				if _gacde != nil {
					return nil, _gacde
				}
				if _ccbd._geba >= _ccbd.NumberOfNewSymbols {
					break
				}
				_gggc += uint32(_fbgc)
				_agef += _gggc
				if !_ccbd.IsHuffmanEncoded || _ccbd.UseRefinementAggregation {
					if !_ccbd.UseRefinementAggregation {
						_gacde = _ccbd.decodeDirectlyThroughGenericRegion(_gggc, uint32(_dcbd))
						if _gacde != nil {
							return nil, _gacde
						}
					} else {
						_gacde = _ccbd.decodeAggregate(_gggc, uint32(_dcbd))
						if _gacde != nil {
							return nil, _gacde
						}
					}
				} else if _ccbd.IsHuffmanEncoded && !_ccbd.UseRefinementAggregation {
					_fcgf[_ccbd._geba] = int(_gggc)
				}
				_ccbd._geba++
			}
			if _ccbd.IsHuffmanEncoded && !_ccbd.UseRefinementAggregation {
				var _aafcf int64
				if _ccbd.SdHuffBMSizeSelection == 0 {
					var _cfaa _aag.Tabler
					_cfaa, _gacde = _aag.GetStandardTable(1)
					if _gacde != nil {
						return nil, _gacde
					}
					_aafcf, _gacde = _cfaa.Decode(_ccbd._aege)
					if _gacde != nil {
						return nil, _gacde
					}
				} else {
					_aafcf, _gacde = _ccbd.huffDecodeBmSize()
					if _gacde != nil {
						return nil, _gacde
					}
				}
				_ccbd._aege.Align()
				var _eacg *_c.Bitmap
				_eacg, _gacde = _ccbd.decodeHeightClassCollectiveBitmap(_aafcf, uint32(_dcbd), _agef)
				if _gacde != nil {
					return nil, _gacde
				}
				_gacde = _ccbd.decodeHeightClassBitmap(_eacg, _bccc, int(_dcbd), _fcgf)
				if _gacde != nil {
					return nil, _gacde
				}
			}
		}
		_cggcb, _gacde := _ccbd.getToExportFlags()
		if _gacde != nil {
			return nil, _gacde
		}
		_ccbd.setExportedSymbols(_cggcb)
	}
	return _ccbd._abab, nil
}
func (_fdgf *Header) Encode(w _g.BinaryWriter) (_eeae int, _ceef error) {
	const _acg = "\u0048\u0065\u0061d\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065"
	var _bgg _g.BinaryWriter
	_gb.Log.Trace("\u005b\u0053\u0045G\u004d\u0045\u004e\u0054-\u0048\u0045\u0041\u0044\u0045\u0052\u005d[\u0045\u004e\u0043\u004f\u0044\u0045\u005d\u0020\u0042\u0065\u0067\u0069\u006e\u0073")
	defer func() {
		if _ceef != nil {
			_gb.Log.Trace("[\u0053\u0045\u0047\u004d\u0045\u004eT\u002d\u0048\u0045\u0041\u0044\u0045R\u005d\u005b\u0045\u004e\u0043\u004f\u0044E\u005d\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u002e\u0020%\u0076", _ceef)
		} else {
			_gb.Log.Trace("\u005b\u0053\u0045\u0047ME\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0025\u0076", _fdgf)
			_gb.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045N\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u005b\u0045\u004e\u0043O\u0044\u0045\u005d\u0020\u0046\u0069\u006ei\u0073\u0068\u0065\u0064")
		}
	}()
	w.FinishByte()
	if _fdgf.SegmentData != nil {
		_edb, _aaca := _fdgf.SegmentData.(SegmentEncoder)
		if !_aaca {
			return 0, _ac.Errorf(_acg, "\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0054\u0020\u0064\u006f\u0065s\u006e\u0027\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074 \u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0045\u006e\u0063\u006f\u0064er\u0020\u0069\u006e\u0074\u0065\u0072\u0066\u0061\u0063\u0065", _fdgf.SegmentData)
		}
		_bgg = _g.BufferedMSB()
		_eeae, _ceef = _edb.Encode(_bgg)
		if _ceef != nil {
			return 0, _ac.Wrap(_ceef, _acg, "")
		}
		_fdgf.SegmentDataLength = uint64(_eeae)
	}
	if _fdgf.pageSize() == 4 {
		_fdgf.PageAssociationFieldSize = true
	}
	var _aca int
	_aca, _ceef = _fdgf.writeSegmentNumber(w)
	if _ceef != nil {
		return 0, _ac.Wrap(_ceef, _acg, "")
	}
	_eeae += _aca
	if _ceef = _fdgf.writeFlags(w); _ceef != nil {
		return _eeae, _ac.Wrap(_ceef, _acg, "")
	}
	_eeae++
	_aca, _ceef = _fdgf.writeReferredToCount(w)
	if _ceef != nil {
		return 0, _ac.Wrap(_ceef, _acg, "")
	}
	_eeae += _aca
	_aca, _ceef = _fdgf.writeReferredToSegments(w)
	if _ceef != nil {
		return 0, _ac.Wrap(_ceef, _acg, "")
	}
	_eeae += _aca
	_aca, _ceef = _fdgf.writeSegmentPageAssociation(w)
	if _ceef != nil {
		return 0, _ac.Wrap(_ceef, _acg, "")
	}
	_eeae += _aca
	_aca, _ceef = _fdgf.writeSegmentDataLength(w)
	if _ceef != nil {
		return 0, _ac.Wrap(_ceef, _acg, "")
	}
	_eeae += _aca
	_fdgf.HeaderLength = int64(_eeae) - int64(_fdgf.SegmentDataLength)
	if _bgg != nil {
		if _, _ceef = w.Write(_bgg.Data()); _ceef != nil {
			return _eeae, _ac.Wrap(_ceef, _acg, "\u0077r\u0069t\u0065\u0020\u0073\u0065\u0067m\u0065\u006et\u0020\u0064\u0061\u0074\u0061")
		}
	}
	return _eeae, nil
}

type Pager interface {
	GetSegment(int) (*Header, error)
	GetBitmap() (*_c.Bitmap, error)
}

func (_cege *TableSegment) HtPS() int32 { return _cege._gagg }
func (_effc *GenericRegion) decodeTemplate1(_cfdd, _aadb, _bced int, _bcdc, _cbb int) (_cdea error) {
	const _dege = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0031"
	var (
		_bdd, _cdgb int
		_bad, _fgg  int
		_cae        byte
		_fbg, _dccb int
	)
	if _cfdd >= 1 {
		_cae, _cdea = _effc.Bitmap.GetByte(_cbb)
		if _cdea != nil {
			return _ac.Wrap(_cdea, _dege, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_bad = int(_cae)
	}
	if _cfdd >= 2 {
		_cae, _cdea = _effc.Bitmap.GetByte(_cbb - _effc.Bitmap.RowStride)
		if _cdea != nil {
			return _ac.Wrap(_cdea, _dege, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_fgg = int(_cae) << 5
	}
	_bdd = ((_bad >> 1) & 0x1f8) | ((_fgg >> 1) & 0x1e00)
	for _aaff := 0; _aaff < _bced; _aaff = _fbg {
		var (
			_dfee byte
			_ccf  int
		)
		_fbg = _aaff + 8
		if _cgbg := _aadb - _aaff; _cgbg > 8 {
			_ccf = 8
		} else {
			_ccf = _cgbg
		}
		if _cfdd > 0 {
			_bad <<= 8
			if _fbg < _aadb {
				_cae, _cdea = _effc.Bitmap.GetByte(_cbb + 1)
				if _cdea != nil {
					return _ac.Wrap(_cdea, _dege, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_bad |= int(_cae)
			}
		}
		if _cfdd > 1 {
			_fgg <<= 8
			if _fbg < _aadb {
				_cae, _cdea = _effc.Bitmap.GetByte(_cbb - _effc.Bitmap.RowStride + 1)
				if _cdea != nil {
					return _ac.Wrap(_cdea, _dege, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_fgg |= int(_cae) << 5
			}
		}
		for _bgag := 0; _bgag < _ccf; _bgag++ {
			if _effc._aaaf {
				_cdgb = _effc.overrideAtTemplate1(_bdd, _aaff+_bgag, _cfdd, int(_dfee), _bgag)
				_effc._bcae.SetIndex(int32(_cdgb))
			} else {
				_effc._bcae.SetIndex(int32(_bdd))
			}
			_dccb, _cdea = _effc._ga.DecodeBit(_effc._bcae)
			if _cdea != nil {
				return _ac.Wrap(_cdea, _dege, "")
			}
			_dfee |= byte(_dccb) << uint(7-_bgag)
			_gecc := uint(8 - _bgag)
			_bdd = ((_bdd & 0xefb) << 1) | _dccb | ((_bad >> _gecc) & 0x8) | ((_fgg >> _gecc) & 0x200)
		}
		if _cdf := _effc.Bitmap.SetByte(_bcdc, _dfee); _cdf != nil {
			return _ac.Wrap(_cdf, _dege, "")
		}
		_bcdc++
		_cbb++
	}
	return nil
}
func (_feace *GenericRegion) readGBAtPixels(_gedb int) error {
	const _dfbg = "\u0072\u0065\u0061\u0064\u0047\u0042\u0041\u0074\u0050i\u0078\u0065\u006c\u0073"
	_feace.GBAtX = make([]int8, _gedb)
	_feace.GBAtY = make([]int8, _gedb)
	for _ccc := 0; _ccc < _gedb; _ccc++ {
		_debba, _bebb := _feace._degf.ReadByte()
		if _bebb != nil {
			return _ac.Wrapf(_bebb, _dfbg, "\u0058\u0020\u0061t\u0020\u0069\u003a\u0020\u0027\u0025\u0064\u0027", _ccc)
		}
		_feace.GBAtX[_ccc] = int8(_debba)
		_debba, _bebb = _feace._degf.ReadByte()
		if _bebb != nil {
			return _ac.Wrapf(_bebb, _dfbg, "\u0059\u0020\u0061t\u0020\u0069\u003a\u0020\u0027\u0025\u0064\u0027", _ccc)
		}
		_feace.GBAtY[_ccc] = int8(_debba)
	}
	return nil
}
func (_fbcce *TextRegion) decodeCurrentT() (int64, error) {
	if _fbcce.SbStrips != 1 {
		if _fbcce.IsHuffmanEncoded {
			_ddbc, _gcf := _fbcce._ceegd.ReadBits(byte(_fbcce.LogSBStrips))
			return int64(_ddbc), _gcf
		}
		_cdfb, _gbgg := _fbcce._bbcea.DecodeInt(_fbcce._gdgaa)
		if _gbgg != nil {
			return 0, _gbgg
		}
		return int64(_cdfb), nil
	}
	return 0, nil
}

type PatternDictionary struct {
	_fad             _g.StreamReader
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
	Patterns         []*_c.Bitmap
	GrayMax          uint32
}

func (_dgcc *TableSegment) StreamReader() _g.StreamReader { return _dgcc._fafc }
func (_gdgg *GenericRegion) updateOverrideFlags() error {
	const _fdca = "\u0075\u0070\u0064\u0061te\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _gdgg.GBAtX == nil || _gdgg.GBAtY == nil {
		return nil
	}
	if len(_gdgg.GBAtX) != len(_gdgg.GBAtY) {
		return _ac.Errorf(_fdca, "i\u006eco\u0073i\u0073t\u0065\u006e\u0074\u0020\u0041T\u0020\u0070\u0069x\u0065\u006c\u002e\u0020\u0041m\u006f\u0075\u006et\u0020\u006f\u0066\u0020\u0027\u0078\u0027\u0020\u0070\u0069\u0078e\u006c\u0073\u003a %d\u002c\u0020\u0041\u006d\u006f\u0075n\u0074\u0020\u006f\u0066\u0020\u0027\u0079\u0027\u0020\u0070\u0069\u0078e\u006cs\u003a\u0020\u0025\u0064", len(_gdgg.GBAtX), len(_gdgg.GBAtY))
	}
	_gdgg.GBAtOverride = make([]bool, len(_gdgg.GBAtX))
	switch _gdgg.GBTemplate {
	case 0:
		if !_gdgg.UseExtTemplates {
			if _gdgg.GBAtX[0] != 3 || _gdgg.GBAtY[0] != -1 {
				_gdgg.setOverrideFlag(0)
			}
			if _gdgg.GBAtX[1] != -3 || _gdgg.GBAtY[1] != -1 {
				_gdgg.setOverrideFlag(1)
			}
			if _gdgg.GBAtX[2] != 2 || _gdgg.GBAtY[2] != -2 {
				_gdgg.setOverrideFlag(2)
			}
			if _gdgg.GBAtX[3] != -2 || _gdgg.GBAtY[3] != -2 {
				_gdgg.setOverrideFlag(3)
			}
		} else {
			if _gdgg.GBAtX[0] != -2 || _gdgg.GBAtY[0] != 0 {
				_gdgg.setOverrideFlag(0)
			}
			if _gdgg.GBAtX[1] != 0 || _gdgg.GBAtY[1] != -2 {
				_gdgg.setOverrideFlag(1)
			}
			if _gdgg.GBAtX[2] != -2 || _gdgg.GBAtY[2] != -1 {
				_gdgg.setOverrideFlag(2)
			}
			if _gdgg.GBAtX[3] != -1 || _gdgg.GBAtY[3] != -2 {
				_gdgg.setOverrideFlag(3)
			}
			if _gdgg.GBAtX[4] != 1 || _gdgg.GBAtY[4] != -2 {
				_gdgg.setOverrideFlag(4)
			}
			if _gdgg.GBAtX[5] != 2 || _gdgg.GBAtY[5] != -1 {
				_gdgg.setOverrideFlag(5)
			}
			if _gdgg.GBAtX[6] != -3 || _gdgg.GBAtY[6] != 0 {
				_gdgg.setOverrideFlag(6)
			}
			if _gdgg.GBAtX[7] != -4 || _gdgg.GBAtY[7] != 0 {
				_gdgg.setOverrideFlag(7)
			}
			if _gdgg.GBAtX[8] != 2 || _gdgg.GBAtY[8] != -2 {
				_gdgg.setOverrideFlag(8)
			}
			if _gdgg.GBAtX[9] != 3 || _gdgg.GBAtY[9] != -1 {
				_gdgg.setOverrideFlag(9)
			}
			if _gdgg.GBAtX[10] != -2 || _gdgg.GBAtY[10] != -2 {
				_gdgg.setOverrideFlag(10)
			}
			if _gdgg.GBAtX[11] != -3 || _gdgg.GBAtY[11] != -1 {
				_gdgg.setOverrideFlag(11)
			}
		}
	case 1:
		if _gdgg.GBAtX[0] != 3 || _gdgg.GBAtY[0] != -1 {
			_gdgg.setOverrideFlag(0)
		}
	case 2:
		if _gdgg.GBAtX[0] != 2 || _gdgg.GBAtY[0] != -1 {
			_gdgg.setOverrideFlag(0)
		}
	case 3:
		if _gdgg.GBAtX[0] != 2 || _gdgg.GBAtY[0] != -1 {
			_gdgg.setOverrideFlag(0)
		}
	}
	return nil
}
func (_fgd *GenericRefinementRegion) decodeTypicalPredictedLine(_cbc, _dcd, _cac, _gf, _dda, _aaf int) error {
	_aaa := _cbc - int(_fgd.ReferenceDY)
	_gec := _fgd.ReferenceBitmap.GetByteIndex(0, _aaa)
	_bac := _fgd.RegionBitmap.GetByteIndex(0, _cbc)
	var _cbd error
	switch _fgd.TemplateID {
	case 0:
		_cbd = _fgd.decodeTypicalPredictedLineTemplate0(_cbc, _dcd, _cac, _gf, _dda, _aaf, _bac, _aaa, _gec)
	case 1:
		_cbd = _fgd.decodeTypicalPredictedLineTemplate1(_cbc, _dcd, _cac, _gf, _dda, _aaf, _bac, _aaa, _gec)
	}
	return _cbd
}
func (_aagg *PageInformationSegment) readIsStriped() error {
	_caacb, _ffde := _aagg._fabba.ReadBit()
	if _ffde != nil {
		return _ffde
	}
	if _caacb == 1 {
		_aagg.IsStripe = true
	}
	return nil
}
func (_fggf *GenericRegion) setParametersMMR(_gac bool, _cbfg, _ggac int64, _egcc, _acda uint32, _fgeg byte, _caa, _dacd bool, _fcf, _bbd []int8) {
	_fggf.DataOffset = _cbfg
	_fggf.DataLength = _ggac
	_fggf.RegionSegment = &RegionSegment{}
	_fggf.RegionSegment.BitmapHeight = _egcc
	_fggf.RegionSegment.BitmapWidth = _acda
	_fggf.GBTemplate = _fgeg
	_fggf.IsMMREncoded = _gac
	_fggf.IsTPGDon = _caa
	_fggf.GBAtX = _fcf
	_fggf.GBAtY = _bbd
}
func (_bdbgg *TextRegion) decodeRdh() (int64, error) {
	const _eeggg = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064h"
	if _bdbgg.IsHuffmanEncoded {
		if _bdbgg.SbHuffRDHeight == 3 {
			if _bdbgg._agba == nil {
				var (
					_fffa int
					_dadf error
				)
				if _bdbgg.SbHuffFS == 3 {
					_fffa++
				}
				if _bdbgg.SbHuffDS == 3 {
					_fffa++
				}
				if _bdbgg.SbHuffDT == 3 {
					_fffa++
				}
				if _bdbgg.SbHuffRDWidth == 3 {
					_fffa++
				}
				_bdbgg._agba, _dadf = _bdbgg.getUserTable(_fffa)
				if _dadf != nil {
					return 0, _ac.Wrap(_dadf, _eeggg, "")
				}
			}
			return _bdbgg._agba.Decode(_bdbgg._ceegd)
		}
		_fabg, _dcfe := _aag.GetStandardTable(14 + int(_bdbgg.SbHuffRDHeight))
		if _dcfe != nil {
			return 0, _ac.Wrap(_dcfe, _eeggg, "")
		}
		return _fabg.Decode(_bdbgg._ceegd)
	}
	_ddbd, _cdada := _bdbgg._bbcea.DecodeInt(_bdbgg._agag)
	if _cdada != nil {
		return 0, _ac.Wrap(_cdada, _eeggg, "")
	}
	return int64(_ddbd), nil
}
func (_cdbg *HalftoneRegion) computeSegmentDataStructure() error {
	_cdbg.DataOffset = _cdbg._ebf.StreamPosition()
	_cdbg.DataHeaderLength = _cdbg.DataOffset - _cdbg.DataHeaderOffset
	_cdbg.DataLength = int64(_cdbg._ebf.Length()) - _cdbg.DataHeaderLength
	return nil
}
func (_dcda *PatternDictionary) computeSegmentDataStructure() error {
	_dcda.DataOffset = _dcda._fad.StreamPosition()
	_dcda.DataHeaderLength = _dcda.DataOffset - _dcda.DataHeaderOffset
	_dcda.DataLength = int64(_dcda._fad.Length()) - _dcda.DataHeaderLength
	return nil
}
func (_cca *PageInformationSegment) Encode(w _g.BinaryWriter) (_dcae int, _degc error) {
	const _beae = "\u0050\u0061g\u0065\u0049\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u002e\u0045\u006eco\u0064\u0065"
	_dggc := make([]byte, 4)
	_bc.BigEndian.PutUint32(_dggc, uint32(_cca.PageBMWidth))
	_dcae, _degc = w.Write(_dggc)
	if _degc != nil {
		return _dcae, _ac.Wrap(_degc, _beae, "\u0077\u0069\u0064t\u0068")
	}
	_bc.BigEndian.PutUint32(_dggc, uint32(_cca.PageBMHeight))
	var _aaege int
	_aaege, _degc = w.Write(_dggc)
	if _degc != nil {
		return _aaege + _dcae, _ac.Wrap(_degc, _beae, "\u0068\u0065\u0069\u0067\u0068\u0074")
	}
	_dcae += _aaege
	_bc.BigEndian.PutUint32(_dggc, uint32(_cca.ResolutionX))
	_aaege, _degc = w.Write(_dggc)
	if _degc != nil {
		return _aaege + _dcae, _ac.Wrap(_degc, _beae, "\u0078\u0020\u0072e\u0073\u006f\u006c\u0075\u0074\u0069\u006f\u006e")
	}
	_dcae += _aaege
	_bc.BigEndian.PutUint32(_dggc, uint32(_cca.ResolutionY))
	if _aaege, _degc = w.Write(_dggc); _degc != nil {
		return _aaege + _dcae, _ac.Wrap(_degc, _beae, "\u0079\u0020\u0072e\u0073\u006f\u006c\u0075\u0074\u0069\u006f\u006e")
	}
	_dcae += _aaege
	if _degc = _cca.encodeFlags(w); _degc != nil {
		return _dcae, _ac.Wrap(_degc, _beae, "")
	}
	_dcae++
	if _aaege, _degc = _cca.encodeStripingInformation(w); _degc != nil {
		return _dcae, _ac.Wrap(_degc, _beae, "")
	}
	_dcae += _aaege
	return _dcae, nil
}
func (_cdfe *Header) referenceSize() uint {
	switch {
	case _cdfe.SegmentNumber <= 255:
		return 1
	case _cdfe.SegmentNumber <= 65535:
		return 2
	default:
		return 4
	}
}
func (_aaag *TableSegment) parseHeader() error {
	var (
		_eaeg int
		_efd  uint64
		_edeb error
	)
	_eaeg, _edeb = _aaag._fafc.ReadBit()
	if _edeb != nil {
		return _edeb
	}
	if _eaeg == 1 {
		return _be.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0061\u0062\u006c\u0065 \u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0064e\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u002e\u0020\u0042\u002e\u0032\u002e1\u0020\u0043\u006f\u0064\u0065\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0066\u006c\u0061\u0067\u0073\u003a\u0020\u0042\u0069\u0074\u0020\u0037\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u007a\u0065\u0072\u006f\u002e\u0020\u0057a\u0073\u003a \u0025\u0064", _eaeg)
	}
	if _efd, _edeb = _aaag._fafc.ReadBits(3); _edeb != nil {
		return _edeb
	}
	_aaag._fgfe = (int32(_efd) + 1) & 0xf
	if _efd, _edeb = _aaag._fafc.ReadBits(3); _edeb != nil {
		return _edeb
	}
	_aaag._gagg = (int32(_efd) + 1) & 0xf
	if _efd, _edeb = _aaag._fafc.ReadBits(32); _edeb != nil {
		return _edeb
	}
	_aaag._afgc = int32(_efd & _f.MaxInt32)
	if _efd, _edeb = _aaag._fafc.ReadBits(32); _edeb != nil {
		return _edeb
	}
	_aaag._bed = int32(_efd & _f.MaxInt32)
	return nil
}
func (_agedg *Header) readSegmentNumber(_dcbf _g.StreamReader) error {
	const _ebd = "\u0072\u0065\u0061\u0064\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u004eu\u006d\u0062\u0065\u0072"
	_fgfb := make([]byte, 4)
	_, _gdgge := _dcbf.Read(_fgfb)
	if _gdgge != nil {
		return _ac.Wrap(_gdgge, _ebd, "")
	}
	_agedg.SegmentNumber = _bc.BigEndian.Uint32(_fgfb)
	return nil
}
func (_fbfe *Header) pageSize() uint {
	if _fbfe.PageAssociation <= 255 {
		return 1
	}
	return 4
}
func NewRegionSegment(r _g.StreamReader) *RegionSegment { return &RegionSegment{_ffade: r} }
func (_bafg *HalftoneRegion) shiftAndFill(_dgda int) int {
	_dgda >>= 8
	if _dgda < 0 {
		_bagc := int(_f.Log(float64(_ffgc(_dgda))) / _f.Log(2))
		_deec := 31 - _bagc
		for _dedb := 1; _dedb < _deec; _dedb++ {
			_dgda |= 1 << uint(31-_dedb)
		}
	}
	return _dgda
}
func (_adcf *TextRegion) getUserTable(_bdegf int) (_aag.Tabler, error) {
	const _bddde = "\u0067\u0065\u0074U\u0073\u0065\u0072\u0054\u0061\u0062\u006c\u0065"
	var _gbdg int
	for _, _gabd := range _adcf.Header.RTSegments {
		if _gabd.Type == 53 {
			if _gbdg == _bdegf {
				_eccd, _gace := _gabd.GetSegmentData()
				if _gace != nil {
					return nil, _gace
				}
				_eefc, _dccbd := _eccd.(*TableSegment)
				if !_dccbd {
					_gb.Log.Debug(_be.Sprintf("\u0073\u0065\u0067\u006d\u0065\u006e\u0074 \u0077\u0069\u0074h\u0020\u0054\u0079p\u0065\u00205\u0033\u0020\u002d\u0020\u0061\u006ed\u0020in\u0064\u0065\u0078\u003a\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0054\u0061\u0062\u006c\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074", _gabd.SegmentNumber))
					return nil, _ac.Error(_bddde, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074 \u0077\u0069\u0074h\u0020\u0054\u0079\u0070e\u0020\u0035\u0033\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u002a\u0054\u0061\u0062\u006c\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
				}
				return _aag.NewEncodedTable(_eefc)
			}
			_gbdg++
		}
	}
	return nil, nil
}
func (_dga *GenericRegion) String() string {
	_acb := &_bb.Builder{}
	_acb.WriteString("\u000a[\u0047E\u004e\u0045\u0052\u0049\u0043 \u0052\u0045G\u0049\u004f\u004e\u005d\u000a")
	_acb.WriteString(_dga.RegionSegment.String() + "\u000a")
	_acb.WriteString(_be.Sprintf("\u0009\u002d\u0020Us\u0065\u0045\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073\u003a\u0020\u0025\u0076\u000a", _dga.UseExtTemplates))
	_acb.WriteString(_be.Sprintf("\u0009\u002d \u0049\u0073\u0054P\u0047\u0044\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _dga.IsTPGDon))
	_acb.WriteString(_be.Sprintf("\u0009-\u0020G\u0042\u0054\u0065\u006d\u0070l\u0061\u0074e\u003a\u0020\u0025\u0064\u000a", _dga.GBTemplate))
	_acb.WriteString(_be.Sprintf("\t\u002d \u0049\u0073\u004d\u004d\u0052\u0045\u006e\u0063o\u0064\u0065\u0064\u003a %\u0076\u000a", _dga.IsMMREncoded))
	_acb.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0047\u0042\u0041\u0074\u0058\u003a\u0020\u0025\u0076\u000a", _dga.GBAtX))
	_acb.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0047\u0042\u0041\u0074\u0059\u003a\u0020\u0025\u0076\u000a", _dga.GBAtY))
	_acb.WriteString(_be.Sprintf("\t\u002d \u0047\u0042\u0041\u0074\u004f\u0076\u0065\u0072r\u0069\u0064\u0065\u003a %\u0076\u000a", _dga.GBAtOverride))
	return _acb.String()
}
func _ffgc(_agac int) int {
	if _agac == 0 {
		return 0
	}
	_agac |= _agac >> 1
	_agac |= _agac >> 2
	_agac |= _agac >> 4
	_agac |= _agac >> 8
	_agac |= _agac >> 16
	return (_agac + 1) >> 1
}
func (_abbg *PageInformationSegment) checkInput() error {
	if _abbg.PageBMHeight == _f.MaxInt32 {
		if !_abbg.IsStripe {
			_gb.Log.Debug("P\u0061\u0067\u0065\u0049\u006e\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0053\u0065\u0067me\u006e\u0074\u002e\u0049s\u0053\u0074\u0072\u0069\u0070\u0065\u0020\u0073\u0068ou\u006c\u0064 \u0062\u0065\u0020\u0074\u0072\u0075\u0065\u002e")
		}
	}
	return nil
}
func (_gbdb *SymbolDictionary) Encode(w _g.BinaryWriter) (_fgbd int, _deef error) {
	const _caffg = "\u0053\u0079\u006dbo\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	if _gbdb == nil {
		return 0, _ac.Error(_caffg, "\u0073\u0079m\u0062\u006f\u006c\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	if _fgbd, _deef = _gbdb.encodeFlags(w); _deef != nil {
		return _fgbd, _ac.Wrap(_deef, _caffg, "")
	}
	_ecfd, _deef := _gbdb.encodeATFlags(w)
	if _deef != nil {
		return _fgbd, _ac.Wrap(_deef, _caffg, "")
	}
	_fgbd += _ecfd
	if _ecfd, _deef = _gbdb.encodeRefinementATFlags(w); _deef != nil {
		return _fgbd, _ac.Wrap(_deef, _caffg, "")
	}
	_fgbd += _ecfd
	if _ecfd, _deef = _gbdb.encodeNumSyms(w); _deef != nil {
		return _fgbd, _ac.Wrap(_deef, _caffg, "")
	}
	_fgbd += _ecfd
	if _ecfd, _deef = _gbdb.encodeSymbols(w); _deef != nil {
		return _fgbd, _ac.Wrap(_deef, _caffg, "")
	}
	_fgbd += _ecfd
	return _fgbd, nil
}
func (_abfg *GenericRegion) GetRegionBitmap() (_cfd *_c.Bitmap, _fecd error) {
	if _abfg.Bitmap != nil {
		return _abfg.Bitmap, nil
	}
	if _abfg.IsMMREncoded {
		if _abfg._cgdc == nil {
			_abfg._cgdc, _fecd = _ae.New(_abfg._degf, int(_abfg.RegionSegment.BitmapWidth), int(_abfg.RegionSegment.BitmapHeight), _abfg.DataOffset, _abfg.DataLength)
			if _fecd != nil {
				return nil, _fecd
			}
		}
		_abfg.Bitmap, _fecd = _abfg._cgdc.UncompressMMR()
		return _abfg.Bitmap, _fecd
	}
	if _fecd = _abfg.updateOverrideFlags(); _fecd != nil {
		return nil, _fecd
	}
	var _ffbf int
	if _abfg._ga == nil {
		_abfg._ga, _fecd = _aa.New(_abfg._degf)
		if _fecd != nil {
			return nil, _fecd
		}
	}
	if _abfg._bcae == nil {
		_abfg._bcae = _aa.NewStats(65536, 1)
	}
	_abfg.Bitmap = _c.New(int(_abfg.RegionSegment.BitmapWidth), int(_abfg.RegionSegment.BitmapHeight))
	_gca := int(uint32(_abfg.Bitmap.Width+7) & (^uint32(7)))
	for _ggeg := 0; _ggeg < _abfg.Bitmap.Height; _ggeg++ {
		if _abfg.IsTPGDon {
			var _cgc int
			_cgc, _fecd = _abfg.decodeSLTP()
			if _fecd != nil {
				return nil, _fecd
			}
			_ffbf ^= _cgc
		}
		if _ffbf == 1 {
			if _ggeg > 0 {
				if _fecd = _abfg.copyLineAbove(_ggeg); _fecd != nil {
					return nil, _fecd
				}
			}
		} else {
			if _fecd = _abfg.decodeLine(_ggeg, _abfg.Bitmap.Width, _gca); _fecd != nil {
				return nil, _fecd
			}
		}
	}
	return _abfg.Bitmap, nil
}
func (_aea *GenericRegion) overrideAtTemplate0b(_fecgg, _bec, _gfce, _bdbg, _ebab, _fcg int) int {
	if _aea.GBAtOverride[0] {
		_fecgg &= 0xFFFD
		if _aea.GBAtY[0] == 0 && _aea.GBAtX[0] >= -int8(_ebab) {
			_fecgg |= (_bdbg >> uint(int8(_fcg)-_aea.GBAtX[0]&0x1)) << 1
		} else {
			_fecgg |= int(_aea.getPixel(_bec+int(_aea.GBAtX[0]), _gfce+int(_aea.GBAtY[0]))) << 1
		}
	}
	if _aea.GBAtOverride[1] {
		_fecgg &= 0xDFFF
		if _aea.GBAtY[1] == 0 && _aea.GBAtX[1] >= -int8(_ebab) {
			_fecgg |= (_bdbg >> uint(int8(_fcg)-_aea.GBAtX[1]&0x1)) << 13
		} else {
			_fecgg |= int(_aea.getPixel(_bec+int(_aea.GBAtX[1]), _gfce+int(_aea.GBAtY[1]))) << 13
		}
	}
	if _aea.GBAtOverride[2] {
		_fecgg &= 0xFDFF
		if _aea.GBAtY[2] == 0 && _aea.GBAtX[2] >= -int8(_ebab) {
			_fecgg |= (_bdbg >> uint(int8(_fcg)-_aea.GBAtX[2]&0x1)) << 9
		} else {
			_fecgg |= int(_aea.getPixel(_bec+int(_aea.GBAtX[2]), _gfce+int(_aea.GBAtY[2]))) << 9
		}
	}
	if _aea.GBAtOverride[3] {
		_fecgg &= 0xBFFF
		if _aea.GBAtY[3] == 0 && _aea.GBAtX[3] >= -int8(_ebab) {
			_fecgg |= (_bdbg >> uint(int8(_fcg)-_aea.GBAtX[3]&0x1)) << 14
		} else {
			_fecgg |= int(_aea.getPixel(_bec+int(_aea.GBAtX[3]), _gfce+int(_aea.GBAtY[3]))) << 14
		}
	}
	if _aea.GBAtOverride[4] {
		_fecgg &= 0xEFFF
		if _aea.GBAtY[4] == 0 && _aea.GBAtX[4] >= -int8(_ebab) {
			_fecgg |= (_bdbg >> uint(int8(_fcg)-_aea.GBAtX[4]&0x1)) << 12
		} else {
			_fecgg |= int(_aea.getPixel(_bec+int(_aea.GBAtX[4]), _gfce+int(_aea.GBAtY[4]))) << 12
		}
	}
	if _aea.GBAtOverride[5] {
		_fecgg &= 0xFFDF
		if _aea.GBAtY[5] == 0 && _aea.GBAtX[5] >= -int8(_ebab) {
			_fecgg |= (_bdbg >> uint(int8(_fcg)-_aea.GBAtX[5]&0x1)) << 5
		} else {
			_fecgg |= int(_aea.getPixel(_bec+int(_aea.GBAtX[5]), _gfce+int(_aea.GBAtY[5]))) << 5
		}
	}
	if _aea.GBAtOverride[6] {
		_fecgg &= 0xFFFB
		if _aea.GBAtY[6] == 0 && _aea.GBAtX[6] >= -int8(_ebab) {
			_fecgg |= (_bdbg >> uint(int8(_fcg)-_aea.GBAtX[6]&0x1)) << 2
		} else {
			_fecgg |= int(_aea.getPixel(_bec+int(_aea.GBAtX[6]), _gfce+int(_aea.GBAtY[6]))) << 2
		}
	}
	if _aea.GBAtOverride[7] {
		_fecgg &= 0xFFF7
		if _aea.GBAtY[7] == 0 && _aea.GBAtX[7] >= -int8(_ebab) {
			_fecgg |= (_bdbg >> uint(int8(_fcg)-_aea.GBAtX[7]&0x1)) << 3
		} else {
			_fecgg |= int(_aea.getPixel(_bec+int(_aea.GBAtX[7]), _gfce+int(_aea.GBAtY[7]))) << 3
		}
	}
	if _aea.GBAtOverride[8] {
		_fecgg &= 0xF7FF
		if _aea.GBAtY[8] == 0 && _aea.GBAtX[8] >= -int8(_ebab) {
			_fecgg |= (_bdbg >> uint(int8(_fcg)-_aea.GBAtX[8]&0x1)) << 11
		} else {
			_fecgg |= int(_aea.getPixel(_bec+int(_aea.GBAtX[8]), _gfce+int(_aea.GBAtY[8]))) << 11
		}
	}
	if _aea.GBAtOverride[9] {
		_fecgg &= 0xFFEF
		if _aea.GBAtY[9] == 0 && _aea.GBAtX[9] >= -int8(_ebab) {
			_fecgg |= (_bdbg >> uint(int8(_fcg)-_aea.GBAtX[9]&0x1)) << 4
		} else {
			_fecgg |= int(_aea.getPixel(_bec+int(_aea.GBAtX[9]), _gfce+int(_aea.GBAtY[9]))) << 4
		}
	}
	if _aea.GBAtOverride[10] {
		_fecgg &= 0x7FFF
		if _aea.GBAtY[10] == 0 && _aea.GBAtX[10] >= -int8(_ebab) {
			_fecgg |= (_bdbg >> uint(int8(_fcg)-_aea.GBAtX[10]&0x1)) << 15
		} else {
			_fecgg |= int(_aea.getPixel(_bec+int(_aea.GBAtX[10]), _gfce+int(_aea.GBAtY[10]))) << 15
		}
	}
	if _aea.GBAtOverride[11] {
		_fecgg &= 0xFDFF
		if _aea.GBAtY[11] == 0 && _aea.GBAtX[11] >= -int8(_ebab) {
			_fecgg |= (_bdbg >> uint(int8(_fcg)-_aea.GBAtX[11]&0x1)) << 10
		} else {
			_fecgg |= int(_aea.getPixel(_bec+int(_aea.GBAtX[11]), _gfce+int(_aea.GBAtY[11]))) << 10
		}
	}
	return _fecgg
}
func (_afbe *SymbolDictionary) decodeHeightClassBitmap(_egbf *_c.Bitmap, _edce int64, _bgafe int, _ecgg []int) error {
	for _dfg := _edce; _dfg < int64(_afbe._geba); _dfg++ {
		var _bccd int
		for _fgdb := _edce; _fgdb <= _dfg-1; _fgdb++ {
			_bccd += _ecgg[_fgdb]
		}
		_efga := _b.Rect(_bccd, 0, _bccd+_ecgg[_dfg], _bgafe)
		_fccgd, _fgea := _c.Extract(_efga, _egbf)
		if _fgea != nil {
			return _fgea
		}
		_afbe._gbeb[_dfg] = _fccgd
		_afbe._deeb = append(_afbe._deeb, _fccgd)
	}
	return nil
}

type Type int

func (_bce *GenericRefinementRegion) GetRegionInfo() *RegionSegment { return _bce.RegionInfo }
func (_fffc *SymbolDictionary) huffDecodeRefAggNInst() (int64, error) {
	if !_fffc.SdHuffAggInstanceSelection {
		_dea, _afbg := _aag.GetStandardTable(1)
		if _afbg != nil {
			return 0, _afbg
		}
		return _dea.Decode(_fffc._aege)
	}
	if _fffc._ece == nil {
		var (
			_aafg int
			_bceg error
		)
		if _fffc.SdHuffDecodeHeightSelection == 3 {
			_aafg++
		}
		if _fffc.SdHuffDecodeWidthSelection == 3 {
			_aafg++
		}
		if _fffc.SdHuffBMSizeSelection == 3 {
			_aafg++
		}
		_fffc._ece, _bceg = _fffc.getUserTable(_aafg)
		if _bceg != nil {
			return 0, _bceg
		}
	}
	return _fffc._ece.Decode(_fffc._aege)
}
func (_afgb *SymbolDictionary) decodeDifferenceWidth() (int64, error) {
	if _afgb.IsHuffmanEncoded {
		switch _afgb.SdHuffDecodeWidthSelection {
		case 0:
			_aebbdd, _fbabb := _aag.GetStandardTable(2)
			if _fbabb != nil {
				return 0, _fbabb
			}
			return _aebbdd.Decode(_afgb._aege)
		case 1:
			_dbgf, _afddd := _aag.GetStandardTable(3)
			if _afddd != nil {
				return 0, _afddd
			}
			return _dbgf.Decode(_afgb._aege)
		case 3:
			if _afgb._dgff == nil {
				var _aebg int
				if _afgb.SdHuffDecodeHeightSelection == 3 {
					_aebg++
				}
				_aadbc, _gfad := _afgb.getUserTable(_aebg)
				if _gfad != nil {
					return 0, _gfad
				}
				_afgb._dgff = _aadbc
			}
			return _afgb._dgff.Decode(_afgb._aege)
		}
	} else {
		_gaf, _cfgb := _afgb._caaa.DecodeInt(_afgb._bba)
		if _cfgb != nil {
			return 0, _cfgb
		}
		return int64(_gaf), nil
	}
	return 0, nil
}
func (_cgg *PatternDictionary) checkInput() error {
	if _cgg.HdpHeight < 1 || _cgg.HdpWidth < 1 {
		return _ag.New("in\u0076\u0061l\u0069\u0064\u0020\u0048\u0065\u0061\u0064\u0065\u0072 \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0057\u0069\u0064\u0074\u0068\u002f\u0048\u0065\u0069\u0067\u0068\u0074\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020g\u0072e\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020z\u0065\u0072o")
	}
	if _cgg.IsMMREncoded {
		if _cgg.HDTemplate != 0 {
			_gb.Log.Debug("\u0076\u0061\u0072\u0069\u0061\u0062\u006c\u0065\u0020\u0048\u0044\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0030")
		}
	}
	return nil
}
func (_gddc *RegionSegment) parseHeader() error {
	const _gda = "p\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_gb.Log.Trace("\u005b\u0052\u0045\u0047I\u004f\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002dH\u0045A\u0044\u0045\u0052\u005d\u0020\u0042\u0065g\u0069\u006e")
	defer func() {
		_gb.Log.Trace("\u005b\u0052\u0045G\u0049\u004f\u004e\u005d[\u0050\u0041\u0052\u0053\u0045\u002d\u0048E\u0041\u0044\u0045\u0052\u005d\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
	}()
	_bdg, _ddgef := _gddc._ffade.ReadBits(32)
	if _ddgef != nil {
		return _ac.Wrap(_ddgef, _gda, "\u0077\u0069\u0064t\u0068")
	}
	_gddc.BitmapWidth = uint32(_bdg & _f.MaxUint32)
	_bdg, _ddgef = _gddc._ffade.ReadBits(32)
	if _ddgef != nil {
		return _ac.Wrap(_ddgef, _gda, "\u0068\u0065\u0069\u0067\u0068\u0074")
	}
	_gddc.BitmapHeight = uint32(_bdg & _f.MaxUint32)
	_bdg, _ddgef = _gddc._ffade.ReadBits(32)
	if _ddgef != nil {
		return _ac.Wrap(_ddgef, _gda, "\u0078\u0020\u006c\u006f\u0063\u0061\u0074\u0069\u006f\u006e")
	}
	_gddc.XLocation = uint32(_bdg & _f.MaxUint32)
	_bdg, _ddgef = _gddc._ffade.ReadBits(32)
	if _ddgef != nil {
		return _ac.Wrap(_ddgef, _gda, "\u0079\u0020\u006c\u006f\u0063\u0061\u0074\u0069\u006f\u006e")
	}
	_gddc.YLocation = uint32(_bdg & _f.MaxUint32)
	if _, _ddgef = _gddc._ffade.ReadBits(5); _ddgef != nil {
		return _ac.Wrap(_ddgef, _gda, "\u0064i\u0072\u0079\u0020\u0072\u0065\u0061d")
	}
	if _ddgef = _gddc.readCombinationOperator(); _ddgef != nil {
		return _ac.Wrap(_ddgef, _gda, "c\u006fm\u0062\u0069\u006e\u0061\u0074\u0069\u006f\u006e \u006f\u0070\u0065\u0072at\u006f\u0072")
	}
	return nil
}
func (_ecc *GenericRegion) parseHeader() (_dcf error) {
	_gb.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052I\u0043\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0050\u0061\u0072s\u0069\u006e\u0067\u0048\u0065\u0061\u0064e\u0072\u002e\u002e\u002e")
	defer func() {
		if _dcf != nil {
			_gb.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0047\u0049\u004f\u004e]\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065\u0061\u0064\u0065r\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u0020\u0077\u0069th\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020\u0025\u0076", _dcf)
		} else {
			_gb.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049C\u002d\u0052\u0045G\u0049\u004f\u004e]\u0020\u0050a\u0072\u0073\u0069\u006e\u0067\u0048e\u0061de\u0072\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u0020\u0053\u0075\u0063\u0063\u0065\u0073\u0073\u0066\u0075\u006c\u006c\u0079\u002e\u002e\u002e")
		}
	}()
	var (
		_eebc int
		_feg  uint64
	)
	if _dcf = _ecc.RegionSegment.parseHeader(); _dcf != nil {
		return _dcf
	}
	if _, _dcf = _ecc._degf.ReadBits(3); _dcf != nil {
		return _dcf
	}
	_eebc, _dcf = _ecc._degf.ReadBit()
	if _dcf != nil {
		return _dcf
	}
	if _eebc == 1 {
		_ecc.UseExtTemplates = true
	}
	_eebc, _dcf = _ecc._degf.ReadBit()
	if _dcf != nil {
		return _dcf
	}
	if _eebc == 1 {
		_ecc.IsTPGDon = true
	}
	_feg, _dcf = _ecc._degf.ReadBits(2)
	if _dcf != nil {
		return _dcf
	}
	_ecc.GBTemplate = byte(_feg & 0xf)
	_eebc, _dcf = _ecc._degf.ReadBit()
	if _dcf != nil {
		return _dcf
	}
	if _eebc == 1 {
		_ecc.IsMMREncoded = true
	}
	if !_ecc.IsMMREncoded {
		_degb := 1
		if _ecc.GBTemplate == 0 {
			_degb = 4
			if _ecc.UseExtTemplates {
				_degb = 12
			}
		}
		if _dcf = _ecc.readGBAtPixels(_degb); _dcf != nil {
			return _dcf
		}
	}
	if _dcf = _ecc.computeSegmentDataStructure(); _dcf != nil {
		return _dcf
	}
	_gb.Log.Trace("\u0025\u0073", _ecc)
	return nil
}

type templater interface {
	form(_adg, _bab, _bbea, _ffdf, _fgc int16) int16
	setIndex(_bda *_aa.DecoderStats)
}

func (_fgecb *SymbolDictionary) setCodingStatistics() error {
	if _fgecb._ddfd == nil {
		_fgecb._ddfd = _aa.NewStats(512, 1)
	}
	if _fgecb._fcbe == nil {
		_fgecb._fcbe = _aa.NewStats(512, 1)
	}
	if _fgecb._bba == nil {
		_fgecb._bba = _aa.NewStats(512, 1)
	}
	if _fgecb._efab == nil {
		_fgecb._efab = _aa.NewStats(512, 1)
	}
	if _fgecb._fddg == nil {
		_fgecb._fddg = _aa.NewStats(512, 1)
	}
	if _fgecb.UseRefinementAggregation && _fgecb._dcfa == nil {
		_fgecb._dcfa = _aa.NewStats(1<<uint(_fgecb._ebafc), 1)
		_fgecb._bcb = _aa.NewStats(512, 1)
		_fgecb._gebb = _aa.NewStats(512, 1)
	}
	if _fgecb._ebff == nil {
		_fgecb._ebff = _aa.NewStats(65536, 1)
	}
	if _fgecb._caaa == nil {
		var _gcbe error
		_fgecb._caaa, _gcbe = _aa.New(_fgecb._aege)
		if _gcbe != nil {
			return _gcbe
		}
	}
	return nil
}
func (_dbab *Header) writeSegmentPageAssociation(_gdba _g.BinaryWriter) (_faa int, _beab error) {
	const _gggf = "w\u0072\u0069\u0074\u0065\u0053\u0065g\u006d\u0065\u006e\u0074\u0050\u0061\u0067\u0065\u0041s\u0073\u006f\u0063i\u0061t\u0069\u006f\u006e"
	if _dbab.pageSize() != 4 {
		if _beab = _gdba.WriteByte(byte(_dbab.PageAssociation)); _beab != nil {
			return 0, _ac.Wrap(_beab, _gggf, "\u0070\u0061\u0067\u0065\u0053\u0069\u007a\u0065\u0020\u0021\u003d\u0020\u0034")
		}
		return 1, nil
	}
	_ccfa := make([]byte, 4)
	_bc.BigEndian.PutUint32(_ccfa, uint32(_dbab.PageAssociation))
	if _faa, _beab = _gdba.Write(_ccfa); _beab != nil {
		return 0, _ac.Wrap(_beab, _gggf, "\u0034 \u0062y\u0074\u0065\u0020\u0070\u0061g\u0065\u0020n\u0075\u006d\u0062\u0065\u0072")
	}
	return _faa, nil
}
func (_efbbbc *SymbolDictionary) decodeThroughTextRegion(_dcdab, _bcfe, _cbca uint32) error {
	if _efbbbc._cfcd == nil {
		_efbbbc._cfcd = _ggad(_efbbbc._aege, nil)
		_efbbbc._cfcd.setContexts(_efbbbc._ebff, _aa.NewStats(512, 1), _aa.NewStats(512, 1), _aa.NewStats(512, 1), _aa.NewStats(512, 1), _efbbbc._dcfa, _aa.NewStats(512, 1), _aa.NewStats(512, 1), _aa.NewStats(512, 1), _aa.NewStats(512, 1))
	}
	if _acbf := _efbbbc.setSymbolsArray(); _acbf != nil {
		return _acbf
	}
	_efbbbc._cfcd.setParameters(_efbbbc._caaa, _efbbbc.IsHuffmanEncoded, true, _dcdab, _bcfe, _cbca, 1, _efbbbc._fgfc+_efbbbc._geba, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, _efbbbc.SdrTemplate, _efbbbc.SdrATX, _efbbbc.SdrATY, _efbbbc._deeb, _efbbbc._ebafc)
	return _efbbbc.addSymbol(_efbbbc._cfcd)
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
	Reader                   _g.StreamReader
	SegmentData              Segmenter
	RTSNumbers               []int
	RetainBits               []uint8
}

func (_ccfb *Header) readNumberOfReferredToSegments(_gegf _g.StreamReader) (uint64, error) {
	const _eed = "\u0072\u0065\u0061\u0064\u004e\u0075\u006d\u0062\u0065\u0072O\u0066\u0052\u0065\u0066\u0065\u0072\u0072e\u0064\u0054\u006f\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073"
	_ggcc, _agce := _gegf.ReadBits(3)
	if _agce != nil {
		return 0, _ac.Wrap(_agce, _eed, "\u0063\u006f\u0075n\u0074\u0020\u006f\u0066\u0020\u0072\u0074\u0073")
	}
	_ggcc &= 0xf
	var _fga []byte
	if _ggcc <= 4 {
		_fga = make([]byte, 5)
		for _gecgg := 0; _gecgg <= 4; _gecgg++ {
			_ebad, _faea := _gegf.ReadBit()
			if _faea != nil {
				return 0, _ac.Wrap(_faea, _eed, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
			}
			_fga[_gecgg] = byte(_ebad)
		}
	} else {
		_ggcc, _agce = _gegf.ReadBits(29)
		if _agce != nil {
			return 0, _agce
		}
		_ggcc &= _f.MaxInt32
		_bagde := (_ggcc + 8) >> 3
		_bagde <<= 3
		_fga = make([]byte, _bagde)
		var _accf uint64
		for _accf = 0; _accf < _bagde; _accf++ {
			_ccbga, _ddba := _gegf.ReadBit()
			if _ddba != nil {
				return 0, _ac.Wrap(_ddba, _eed, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
			}
			_fga[_accf] = byte(_ccbga)
		}
	}
	return _ggcc, nil
}
func (_fcbc *PageInformationSegment) readCombinationOperatorOverrideAllowed() error {
	_fdda, _cbbd := _fcbc._fabba.ReadBit()
	if _cbbd != nil {
		return _cbbd
	}
	if _fdda == 1 {
		_fcbc._bddd = true
	}
	return nil
}
func (_dfgf *TextRegion) setContexts(_abge *_aa.DecoderStats, _gbeda *_aa.DecoderStats, _ebc *_aa.DecoderStats, _cfdf *_aa.DecoderStats, _eafb *_aa.DecoderStats, _cgfd *_aa.DecoderStats, _bcbg *_aa.DecoderStats, _gfb *_aa.DecoderStats, _fccc *_aa.DecoderStats, _gaeg *_aa.DecoderStats) {
	_dfgf._fdfbf = _gbeda
	_dfgf._dfge = _ebc
	_dfgf._bcca = _cfdf
	_dfgf._gdgaa = _eafb
	_dfgf._bbgdf = _bcbg
	_dfgf._agag = _gfb
	_dfgf._bcfb = _cgfd
	_dfgf._bfce = _fccc
	_dfgf._badb = _gaeg
	_dfgf._gfd = _abge
}
func (_ebace *TextRegion) readHuffmanFlags() error {
	var (
		_febg int
		_dgfc uint64
		_dagb error
	)
	_, _dagb = _ebace._ceegd.ReadBit()
	if _dagb != nil {
		return _dagb
	}
	_febg, _dagb = _ebace._ceegd.ReadBit()
	if _dagb != nil {
		return _dagb
	}
	_ebace.SbHuffRSize = int8(_febg)
	_dgfc, _dagb = _ebace._ceegd.ReadBits(2)
	if _dagb != nil {
		return _dagb
	}
	_ebace.SbHuffRDY = int8(_dgfc) & 0xf
	_dgfc, _dagb = _ebace._ceegd.ReadBits(2)
	if _dagb != nil {
		return _dagb
	}
	_ebace.SbHuffRDX = int8(_dgfc) & 0xf
	_dgfc, _dagb = _ebace._ceegd.ReadBits(2)
	if _dagb != nil {
		return _dagb
	}
	_ebace.SbHuffRDHeight = int8(_dgfc) & 0xf
	_dgfc, _dagb = _ebace._ceegd.ReadBits(2)
	if _dagb != nil {
		return _dagb
	}
	_ebace.SbHuffRDWidth = int8(_dgfc) & 0xf
	_dgfc, _dagb = _ebace._ceegd.ReadBits(2)
	if _dagb != nil {
		return _dagb
	}
	_ebace.SbHuffDT = int8(_dgfc) & 0xf
	_dgfc, _dagb = _ebace._ceegd.ReadBits(2)
	if _dagb != nil {
		return _dagb
	}
	_ebace.SbHuffDS = int8(_dgfc) & 0xf
	_dgfc, _dagb = _ebace._ceegd.ReadBits(2)
	if _dagb != nil {
		return _dagb
	}
	_ebace.SbHuffFS = int8(_dgfc) & 0xf
	return nil
}
func (_ffdb *GenericRefinementRegion) parseHeader() (_ebg error) {
	_gb.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072s\u0069\u006e\u0067\u0020\u0048e\u0061\u0064e\u0072\u002e\u002e\u002e")
	_cef := _a.Now()
	defer func() {
		if _ebg == nil {
			_gb.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045G\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020h\u0065\u0061\u0064\u0065\u0072\u0020\u0066\u0069\u006e\u0069\u0073\u0068id\u0020\u0069\u006e\u003a\u0020\u0025\u0064\u0020\u006e\u0073", _a.Since(_cef).Nanoseconds())
		} else {
			_gb.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0073", _ebg)
		}
	}()
	if _ebg = _ffdb.RegionInfo.parseHeader(); _ebg != nil {
		return _ebg
	}
	_, _ebg = _ffdb._ea.ReadBits(6)
	if _ebg != nil {
		return _ebg
	}
	_ffdb.IsTPGROn, _ebg = _ffdb._ea.ReadBool()
	if _ebg != nil {
		return _ebg
	}
	var _bbce int
	_bbce, _ebg = _ffdb._ea.ReadBit()
	if _ebg != nil {
		return _ebg
	}
	_ffdb.TemplateID = int8(_bbce)
	switch _ffdb.TemplateID {
	case 0:
		_ffdb.Template = _ffdb._fa
		if _ebg = _ffdb.readAtPixels(); _ebg != nil {
			return
		}
	case 1:
		_ffdb.Template = _ffdb._bcab
	}
	return nil
}
func (_ecfe *SymbolDictionary) Init(h *Header, r _g.StreamReader) error {
	_ecfe.Header = h
	_ecfe._aege = r
	return _ecfe.parseHeader()
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

func (_ddd *HalftoneRegion) Init(hd *Header, r _g.StreamReader) error {
	_ddd._ebf = r
	_ddd._bdeg = hd
	_ddd.RegionSegment = NewRegionSegment(r)
	return _ddd.parseHeader()
}
func (_bdce *TextRegion) computeSymbolCodeLength() error {
	if _bdce.IsHuffmanEncoded {
		return _bdce.symbolIDCodeLengths()
	}
	_bdce._bfa = int8(_f.Ceil(_f.Log(float64(_bdce.NumberOfSymbols)) / _f.Log(2)))
	return nil
}
func _agfb(_fcd _g.StreamReader, _aaae *Header) *GenericRefinementRegion {
	return &GenericRefinementRegion{_ea: _fcd, RegionInfo: NewRegionSegment(_fcd), _ff: _aaae, _fa: &template0{}, _bcab: &template1{}}
}
func (_cbe *GenericRegion) decodeSLTP() (int, error) {
	switch _cbe.GBTemplate {
	case 0:
		_cbe._bcae.SetIndex(0x9B25)
	case 1:
		_cbe._bcae.SetIndex(0x795)
	case 2:
		_cbe._bcae.SetIndex(0xE5)
	case 3:
		_cbe._bcae.SetIndex(0x195)
	}
	return _cbe._ga.DecodeBit(_cbe._bcae)
}
func (_fcddd *TextRegion) symbolIDCodeLengths() error {
	var (
		_agdf []*_aag.Code
		_dcgf uint64
		_ddfg _aag.Tabler
		_fbaf error
	)
	for _ffgg := 0; _ffgg < 35; _ffgg++ {
		_dcgf, _fbaf = _fcddd._ceegd.ReadBits(4)
		if _fbaf != nil {
			return _fbaf
		}
		_added := int(_dcgf & 0xf)
		if _added > 0 {
			_agdf = append(_agdf, _aag.NewCode(int32(_added), 0, int32(_ffgg), false))
		}
	}
	_ddfg, _fbaf = _aag.NewFixedSizeTable(_agdf)
	if _fbaf != nil {
		return _fbaf
	}
	var (
		_gbged int64
		_bafb  uint32
		_afgfc []*_aag.Code
		_bdbgb int64
	)
	for _bafb < _fcddd.NumberOfSymbols {
		_bdbgb, _fbaf = _ddfg.Decode(_fcddd._ceegd)
		if _fbaf != nil {
			return _fbaf
		}
		if _bdbgb < 32 {
			if _bdbgb > 0 {
				_afgfc = append(_afgfc, _aag.NewCode(int32(_bdbgb), 0, int32(_bafb), false))
			}
			_gbged = _bdbgb
			_bafb++
		} else {
			var _bggddb, _abce int64
			switch _bdbgb {
			case 32:
				_dcgf, _fbaf = _fcddd._ceegd.ReadBits(2)
				if _fbaf != nil {
					return _fbaf
				}
				_bggddb = 3 + int64(_dcgf)
				if _bafb > 0 {
					_abce = _gbged
				}
			case 33:
				_dcgf, _fbaf = _fcddd._ceegd.ReadBits(3)
				if _fbaf != nil {
					return _fbaf
				}
				_bggddb = 3 + int64(_dcgf)
			case 34:
				_dcgf, _fbaf = _fcddd._ceegd.ReadBits(7)
				if _fbaf != nil {
					return _fbaf
				}
				_bggddb = 11 + int64(_dcgf)
			}
			for _gbcc := 0; _gbcc < int(_bggddb); _gbcc++ {
				if _abce > 0 {
					_afgfc = append(_afgfc, _aag.NewCode(int32(_abce), 0, int32(_bafb), false))
				}
				_bafb++
			}
		}
	}
	_fcddd._ceegd.Align()
	_fcddd._gaead, _fbaf = _aag.NewFixedSizeTable(_afgfc)
	return _fbaf
}

type SegmentEncoder interface {
	Encode(_bcdb _g.BinaryWriter) (_acbb int, _bef error)
}

func (_fef *TextRegion) initSymbols() error {
	const _gdfe = "i\u006e\u0069\u0074\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	for _, _dbgd := range _fef.Header.RTSegments {
		if _dbgd == nil {
			return _ac.Error(_gdfe, "\u006e\u0069\u006c\u0020\u0073\u0065\u0067\u006de\u006e\u0074\u0020pr\u006f\u0076\u0069\u0064\u0065\u0064 \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0074\u0065\u0078\u0074\u0020\u0072\u0065g\u0069\u006f\u006e\u0020\u0053\u0079\u006d\u0062o\u006c\u0073")
		}
		if _dbgd.Type == 0 {
			_eddeg, _eagf := _dbgd.GetSegmentData()
			if _eagf != nil {
				return _ac.Wrap(_eagf, _gdfe, "")
			}
			_gbda, _feecb := _eddeg.(*SymbolDictionary)
			if !_feecb {
				return _ac.Error(_gdfe, "\u0072e\u0066\u0065r\u0072\u0065\u0064 \u0054\u006f\u0020\u0053\u0065\u0067\u006de\u006e\u0074\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0053\u0079\u006d\u0062\u006f\u006cD\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
			}
			_gbda._dcfa = _fef._bcfb
			_edgg, _eagf := _gbda.GetDictionary()
			if _eagf != nil {
				return _ac.Wrap(_eagf, _gdfe, "")
			}
			_fef.Symbols = append(_fef.Symbols, _edgg...)
		}
	}
	_fef.NumberOfSymbols = uint32(len(_fef.Symbols))
	return nil
}
func (_ccg *TextRegion) GetRegionBitmap() (*_c.Bitmap, error) {
	if _ccg.RegionBitmap != nil {
		return _ccg.RegionBitmap, nil
	}
	if !_ccg.IsHuffmanEncoded {
		if _aaagg := _ccg.setCodingStatistics(); _aaagg != nil {
			return nil, _aaagg
		}
	}
	if _gffa := _ccg.createRegionBitmap(); _gffa != nil {
		return nil, _gffa
	}
	if _cgge := _ccg.decodeSymbolInstances(); _cgge != nil {
		return nil, _cgge
	}
	return _ccg.RegionBitmap, nil
}
func (_fddde *TextRegion) setCodingStatistics() error {
	if _fddde._fdfbf == nil {
		_fddde._fdfbf = _aa.NewStats(512, 1)
	}
	if _fddde._dfge == nil {
		_fddde._dfge = _aa.NewStats(512, 1)
	}
	if _fddde._bcca == nil {
		_fddde._bcca = _aa.NewStats(512, 1)
	}
	if _fddde._gdgaa == nil {
		_fddde._gdgaa = _aa.NewStats(512, 1)
	}
	if _fddde._feaceb == nil {
		_fddde._feaceb = _aa.NewStats(512, 1)
	}
	if _fddde._bbgdf == nil {
		_fddde._bbgdf = _aa.NewStats(512, 1)
	}
	if _fddde._agag == nil {
		_fddde._agag = _aa.NewStats(512, 1)
	}
	if _fddde._bcfb == nil {
		_fddde._bcfb = _aa.NewStats(1<<uint(_fddde._bfa), 1)
	}
	if _fddde._bfce == nil {
		_fddde._bfce = _aa.NewStats(512, 1)
	}
	if _fddde._badb == nil {
		_fddde._badb = _aa.NewStats(512, 1)
	}
	if _fddde._bbcea == nil {
		var _fccd error
		_fddde._bbcea, _fccd = _aa.New(_fddde._ceegd)
		if _fccd != nil {
			return _fccd
		}
	}
	return nil
}
func (_edf *GenericRegion) setParametersWithAt(_cabc bool, _ggg byte, _daa, _ddad bool, _ccbg, _ded []int8, _fbf, _bgb uint32, _eadf *_aa.DecoderStats, _ffad *_aa.Decoder) {
	_edf.IsMMREncoded = _cabc
	_edf.GBTemplate = _ggg
	_edf.IsTPGDon = _daa
	_edf.GBAtX = _ccbg
	_edf.GBAtY = _ded
	_edf.RegionSegment.BitmapHeight = _bgb
	_edf.RegionSegment.BitmapWidth = _fbf
	_edf._cgdc = nil
	_edf.Bitmap = nil
	if _eadf != nil {
		_edf._bcae = _eadf
	}
	if _ffad != nil {
		_edf._ga = _ffad
	}
	_gb.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0047\u0049O\u004e\u005d\u0020\u0073\u0065\u0074P\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0053\u0044\u0041t\u003a\u0020\u0025\u0073", _edf)
}
func (_edfa *TextRegion) decodeDfs() (int64, error) {
	if _edfa.IsHuffmanEncoded {
		if _edfa.SbHuffFS == 3 {
			if _edfa._cdcga == nil {
				var _dbeb error
				_edfa._cdcga, _dbeb = _edfa.getUserTable(0)
				if _dbeb != nil {
					return 0, _dbeb
				}
			}
			return _edfa._cdcga.Decode(_edfa._ceegd)
		}
		_daeb, _ggbg := _aag.GetStandardTable(6 + int(_edfa.SbHuffFS))
		if _ggbg != nil {
			return 0, _ggbg
		}
		return _daeb.Decode(_edfa._ceegd)
	}
	_dggd, _cgea := _edfa._bbcea.DecodeInt(_edfa._dfge)
	if _cgea != nil {
		return 0, _cgea
	}
	return int64(_dggd), nil
}
func (_egfa *SymbolDictionary) decodeHeightClassDeltaHeightWithHuffman() (int64, error) {
	switch _egfa.SdHuffDecodeHeightSelection {
	case 0:
		_ffac, _badc := _aag.GetStandardTable(4)
		if _badc != nil {
			return 0, _badc
		}
		return _ffac.Decode(_egfa._aege)
	case 1:
		_fddf, _bdaa := _aag.GetStandardTable(5)
		if _bdaa != nil {
			return 0, _bdaa
		}
		return _fddf.Decode(_egfa._aege)
	case 3:
		if _egfa._bdgc == nil {
			_dgdc, _cgfc := _aag.GetStandardTable(0)
			if _cgfc != nil {
				return 0, _cgfc
			}
			_egfa._bdgc = _dgdc
		}
		return _egfa._bdgc.Decode(_egfa._aege)
	}
	return 0, nil
}
func (_abe *GenericRegion) GetRegionInfo() *RegionSegment { return _abe.RegionSegment }
func (_efbbb Type) String() string {
	switch _efbbb {
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
func (_cabf *PatternDictionary) parseHeader() error {
	_gb.Log.Trace("\u005b\u0050\u0041\u0054\u0054\u0045\u0052\u004e\u002d\u0044\u0049\u0043\u0054I\u004f\u004e\u0041\u0052\u0059\u005d[\u0070\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072\u005d\u0020b\u0065\u0067\u0069\u006e")
	defer func() {
		_gb.Log.Trace("\u005b\u0050\u0041T\u0054\u0045\u0052\u004e\u002d\u0044\u0049\u0043\u0054\u0049\u004f\u004e\u0041\u0052\u0059\u005d\u005b\u0070\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072\u005d \u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
	}()
	_, _efaea := _cabf._fad.ReadBits(5)
	if _efaea != nil {
		return _efaea
	}
	if _efaea = _cabf.readTemplate(); _efaea != nil {
		return _efaea
	}
	if _efaea = _cabf.readIsMMREncoded(); _efaea != nil {
		return _efaea
	}
	if _efaea = _cabf.readPatternWidthAndHeight(); _efaea != nil {
		return _efaea
	}
	if _efaea = _cabf.readGrayMax(); _efaea != nil {
		return _efaea
	}
	if _efaea = _cabf.computeSegmentDataStructure(); _efaea != nil {
		return _efaea
	}
	return _cabf.checkInput()
}
func (_dfdf *SymbolDictionary) huffDecodeBmSize() (int64, error) {
	if _dfdf._cacb == nil {
		var (
			_bcaf  int
			_bgagb error
		)
		if _dfdf.SdHuffDecodeHeightSelection == 3 {
			_bcaf++
		}
		if _dfdf.SdHuffDecodeWidthSelection == 3 {
			_bcaf++
		}
		_dfdf._cacb, _bgagb = _dfdf.getUserTable(_bcaf)
		if _bgagb != nil {
			return 0, _bgagb
		}
	}
	return _dfdf._cacb.Decode(_dfdf._aege)
}

type GenericRegion struct {
	_degf            _g.StreamReader
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
	_aaaf            bool
	Bitmap           *_c.Bitmap
	_ga              *_aa.Decoder
	_bcae            *_aa.DecoderStats
	_cgdc            *_ae.Decoder
}

func (_fea *template1) setIndex(_cee *_aa.DecoderStats) { _cee.SetIndex(0x080) }
func (_eggd *GenericRegion) writeGBAtPixels(_faf _g.BinaryWriter) (_fbag int, _cdc error) {
	const _aac = "\u0077r\u0069t\u0065\u0047\u0042\u0041\u0074\u0050\u0069\u0078\u0065\u006c\u0073"
	if _eggd.UseMMR {
		return 0, nil
	}
	_cdcg := 1
	if _eggd.GBTemplate == 0 {
		_cdcg = 4
	} else if _eggd.UseExtTemplates {
		_cdcg = 12
	}
	if len(_eggd.GBAtX) != _cdcg {
		return 0, _ac.Errorf(_aac, "\u0067\u0062\u0020\u0061\u0074\u0020\u0070\u0061\u0069\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020d\u006f\u0065\u0073\u006e\u0027\u0074\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u006f\u0020\u0047\u0042\u0041t\u0058\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u006c\u0065\u006e")
	}
	if len(_eggd.GBAtY) != _cdcg {
		return 0, _ac.Errorf(_aac, "\u0067\u0062\u0020\u0061\u0074\u0020\u0070\u0061\u0069\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020d\u006f\u0065\u0073\u006e\u0027\u0074\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u006f\u0020\u0047\u0042\u0041t\u0059\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u006c\u0065\u006e")
	}
	for _fce := 0; _fce < _cdcg; _fce++ {
		if _cdc = _faf.WriteByte(byte(_eggd.GBAtX[_fce])); _cdc != nil {
			return _fbag, _ac.Wrap(_cdc, _aac, "w\u0072\u0069\u0074\u0065\u0020\u0047\u0042\u0041\u0074\u0058")
		}
		_fbag++
		if _cdc = _faf.WriteByte(byte(_eggd.GBAtY[_fce])); _cdc != nil {
			return _fbag, _ac.Wrap(_cdc, _aac, "w\u0072\u0069\u0074\u0065\u0020\u0047\u0042\u0041\u0074\u0059")
		}
		_fbag++
	}
	return _fbag, nil
}
func (_cdcgc *TextRegion) decodeIb(_gceg, _fdddd int64) (*_c.Bitmap, error) {
	const _dabg = "\u0064\u0065\u0063\u006f\u0064\u0065\u0049\u0062"
	var (
		_gced  error
		_gbgeb *_c.Bitmap
	)
	if _gceg == 0 {
		if int(_fdddd) > len(_cdcgc.Symbols)-1 {
			return nil, _ac.Error(_dabg, "\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0049\u0042\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0020\u0069\u006e\u0064\u0065x\u0020\u006f\u0075\u0074\u0020o\u0066\u0020r\u0061\u006e\u0067\u0065")
		}
		return _cdcgc.Symbols[int(_fdddd)], nil
	}
	var _fega, _gbee, _bcgg, _bdgb int64
	_fega, _gced = _cdcgc.decodeRdw()
	if _gced != nil {
		return nil, _ac.Wrap(_gced, _dabg, "")
	}
	_gbee, _gced = _cdcgc.decodeRdh()
	if _gced != nil {
		return nil, _ac.Wrap(_gced, _dabg, "")
	}
	_bcgg, _gced = _cdcgc.decodeRdx()
	if _gced != nil {
		return nil, _ac.Wrap(_gced, _dabg, "")
	}
	_bdgb, _gced = _cdcgc.decodeRdy()
	if _gced != nil {
		return nil, _ac.Wrap(_gced, _dabg, "")
	}
	if _cdcgc.IsHuffmanEncoded {
		if _, _gced = _cdcgc.decodeSymInRefSize(); _gced != nil {
			return nil, _ac.Wrap(_gced, _dabg, "")
		}
		_cdcgc._ceegd.Align()
	}
	_bcaeb := _cdcgc.Symbols[_fdddd]
	_eeeg := uint32(_bcaeb.Width)
	_ecaa := uint32(_bcaeb.Height)
	_gedd := int32(uint32(_fega)>>1) + int32(_bcgg)
	_aefc := int32(uint32(_gbee)>>1) + int32(_bdgb)
	if _cdcgc._baea == nil {
		_cdcgc._baea = _agfb(_cdcgc._ceegd, nil)
	}
	_cdcgc._baea.setParameters(_cdcgc._gfd, _cdcgc._bbcea, _cdcgc.SbrTemplate, _eeeg+uint32(_fega), _ecaa+uint32(_gbee), _bcaeb, _gedd, _aefc, false, _cdcgc.SbrATX, _cdcgc.SbrATY)
	_gbgeb, _gced = _cdcgc._baea.GetRegionBitmap()
	if _gced != nil {
		return nil, _ac.Wrap(_gced, _dabg, "\u0067\u0072\u0066")
	}
	if _cdcgc.IsHuffmanEncoded {
		_cdcgc._ceegd.Align()
	}
	return _gbgeb, nil
}
func (_egd *RegionSegment) Encode(w _g.BinaryWriter) (_adcc int, _ebaf error) {
	const _fddd = "R\u0065g\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065n\u0074\u002e\u0045\u006eco\u0064\u0065"
	_cggc := make([]byte, 4)
	_bc.BigEndian.PutUint32(_cggc, _egd.BitmapWidth)
	_adcc, _ebaf = w.Write(_cggc)
	if _ebaf != nil {
		return 0, _ac.Wrap(_ebaf, _fddd, "\u0057\u0069\u0064t\u0068")
	}
	_bc.BigEndian.PutUint32(_cggc, _egd.BitmapHeight)
	var _gbbd int
	_gbbd, _ebaf = w.Write(_cggc)
	if _ebaf != nil {
		return 0, _ac.Wrap(_ebaf, _fddd, "\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_adcc += _gbbd
	_bc.BigEndian.PutUint32(_cggc, _egd.XLocation)
	_gbbd, _ebaf = w.Write(_cggc)
	if _ebaf != nil {
		return 0, _ac.Wrap(_ebaf, _fddd, "\u0058L\u006f\u0063\u0061\u0074\u0069\u006fn")
	}
	_adcc += _gbbd
	_bc.BigEndian.PutUint32(_cggc, _egd.YLocation)
	_gbbd, _ebaf = w.Write(_cggc)
	if _ebaf != nil {
		return 0, _ac.Wrap(_ebaf, _fddd, "\u0059L\u006f\u0063\u0061\u0074\u0069\u006fn")
	}
	_adcc += _gbbd
	if _ebaf = w.WriteByte(byte(_egd.CombinaionOperator) & 0x07); _ebaf != nil {
		return 0, _ac.Wrap(_ebaf, _fddd, "c\u006fm\u0062\u0069\u006e\u0061\u0074\u0069\u006f\u006e \u006f\u0070\u0065\u0072at\u006f\u0072")
	}
	_adcc++
	return _adcc, nil
}
func (_fedd *TextRegion) Encode(w _g.BinaryWriter) (_bebf int, _gcag error) {
	const _cdff = "\u0054\u0065\u0078\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u002e\u0045n\u0063\u006f\u0064\u0065"
	if _bebf, _gcag = _fedd.RegionInfo.Encode(w); _gcag != nil {
		return _bebf, _ac.Wrap(_gcag, _cdff, "")
	}
	var _eegg int
	if _eegg, _gcag = _fedd.encodeFlags(w); _gcag != nil {
		return _bebf, _ac.Wrap(_gcag, _cdff, "")
	}
	_bebf += _eegg
	if _eegg, _gcag = _fedd.encodeSymbols(w); _gcag != nil {
		return _bebf, _ac.Wrap(_gcag, _cdff, "")
	}
	_bebf += _eegg
	return _bebf, nil
}
func (_aee *GenericRefinementRegion) decodeTypicalPredictedLineTemplate0(_gcb, _ecf, _cd, _eg, _fbc, _ege, _dde, _cde, _eaf int) error {
	var (
		_bbc, _egef, _dgd, _bfd, _dfe, _gga int
		_bdf                                byte
		_dfdg                               error
	)
	if _gcb > 0 {
		_bdf, _dfdg = _aee.RegionBitmap.GetByte(_dde - _cd)
		if _dfdg != nil {
			return _dfdg
		}
		_dgd = int(_bdf)
	}
	if _cde > 0 && _cde <= _aee.ReferenceBitmap.Height {
		_bdf, _dfdg = _aee.ReferenceBitmap.GetByte(_eaf - _eg + _ege)
		if _dfdg != nil {
			return _dfdg
		}
		_bfd = int(_bdf) << 4
	}
	if _cde >= 0 && _cde < _aee.ReferenceBitmap.Height {
		_bdf, _dfdg = _aee.ReferenceBitmap.GetByte(_eaf + _ege)
		if _dfdg != nil {
			return _dfdg
		}
		_dfe = int(_bdf) << 1
	}
	if _cde > -2 && _cde < _aee.ReferenceBitmap.Height-1 {
		_bdf, _dfdg = _aee.ReferenceBitmap.GetByte(_eaf + _eg + _ege)
		if _dfdg != nil {
			return _dfdg
		}
		_gga = int(_bdf)
	}
	_bbc = ((_dgd >> 5) & 0x6) | ((_gga >> 2) & 0x30) | (_dfe & 0x180) | (_bfd & 0xc00)
	var _gfe int
	for _aff := 0; _aff < _fbc; _aff = _gfe {
		var _bde int
		_gfe = _aff + 8
		var _gef int
		if _gef = _ecf - _aff; _gef > 8 {
			_gef = 8
		}
		_ffg := _gfe < _ecf
		_agg := _gfe < _aee.ReferenceBitmap.Width
		_dfda := _ege + 1
		if _gcb > 0 {
			_bdf = 0
			if _ffg {
				_bdf, _dfdg = _aee.RegionBitmap.GetByte(_dde - _cd + 1)
				if _dfdg != nil {
					return _dfdg
				}
			}
			_dgd = (_dgd << 8) | int(_bdf)
		}
		if _cde > 0 && _cde <= _aee.ReferenceBitmap.Height {
			var _fbb int
			if _agg {
				_bdf, _dfdg = _aee.ReferenceBitmap.GetByte(_eaf - _eg + _dfda)
				if _dfdg != nil {
					return _dfdg
				}
				_fbb = int(_bdf) << 4
			}
			_bfd = (_bfd << 8) | _fbb
		}
		if _cde >= 0 && _cde < _aee.ReferenceBitmap.Height {
			var _abf int
			if _agg {
				_bdf, _dfdg = _aee.ReferenceBitmap.GetByte(_eaf + _dfda)
				if _dfdg != nil {
					return _dfdg
				}
				_abf = int(_bdf) << 1
			}
			_dfe = (_dfe << 8) | _abf
		}
		if _cde > -2 && _cde < (_aee.ReferenceBitmap.Height-1) {
			_bdf = 0
			if _agg {
				_bdf, _dfdg = _aee.ReferenceBitmap.GetByte(_eaf + _eg + _dfda)
				if _dfdg != nil {
					return _dfdg
				}
			}
			_gga = (_gga << 8) | int(_bdf)
		}
		for _fdb := 0; _fdb < _gef; _fdb++ {
			var _eab int
			_fbd := false
			_eac := (_bbc >> 4) & 0x1ff
			if _eac == 0x1ff {
				_fbd = true
				_eab = 1
			} else if _eac == 0x00 {
				_fbd = true
			}
			if !_fbd {
				if _aee._age {
					_egef = _aee.overrideAtTemplate0(_bbc, _aff+_fdb, _gcb, _bde, _fdb)
					_aee._ef.SetIndex(int32(_egef))
				} else {
					_aee._ef.SetIndex(int32(_bbc))
				}
				_eab, _dfdg = _aee._fb.DecodeBit(_aee._ef)
				if _dfdg != nil {
					return _dfdg
				}
			}
			_ecg := uint(7 - _fdb)
			_bde |= _eab << _ecg
			_bbc = ((_bbc & 0xdb6) << 1) | _eab | (_dgd>>_ecg+5)&0x002 | ((_gga>>_ecg + 2) & 0x010) | ((_dfe >> _ecg) & 0x080) | ((_bfd >> _ecg) & 0x400)
		}
		_dfdg = _aee.RegionBitmap.SetByte(_dde, byte(_bde))
		if _dfdg != nil {
			return _dfdg
		}
		_dde++
		_eaf++
	}
	return nil
}
func (_bffc *HalftoneRegion) combineGrayscalePlanes(_cbde []*_c.Bitmap, _cdeag int) error {
	_eadc := 0
	for _gcbc := 0; _gcbc < _cbde[_cdeag].Height; _gcbc++ {
		for _faff := 0; _faff < _cbde[_cdeag].Width; _faff += 8 {
			_gdga, _adfe := _cbde[_cdeag+1].GetByte(_eadc)
			if _adfe != nil {
				return _adfe
			}
			_ddcb, _adfe := _cbde[_cdeag].GetByte(_eadc)
			if _adfe != nil {
				return _adfe
			}
			_adfe = _cbde[_cdeag].SetByte(_eadc, _c.CombineBytes(_ddcb, _gdga, _c.CmbOpXor))
			if _adfe != nil {
				return _adfe
			}
			_eadc++
		}
	}
	return nil
}
func (_dbaf *GenericRefinementRegion) getPixel(_bfb *_c.Bitmap, _cdb, _fge int) int {
	if _cdb < 0 || _cdb >= _bfb.Width {
		return 0
	}
	if _fge < 0 || _fge >= _bfb.Height {
		return 0
	}
	if _bfb.GetPixel(_cdb, _fge) {
		return 1
	}
	return 0
}
func (_fdgc *SymbolDictionary) setRefinementAtPixels() error {
	if !_fdgc.UseRefinementAggregation || _fdgc.SdrTemplate != 0 {
		return nil
	}
	if _aecdf := _fdgc.readRefinementAtPixels(2); _aecdf != nil {
		return _aecdf
	}
	return nil
}
func (_abb *Header) readReferredToSegmentNumbers(_cbba _g.StreamReader, _fbdc int) ([]int, error) {
	const _dbagg = "\u0072\u0065\u0061\u0064R\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0054\u006f\u0053e\u0067m\u0065\u006e\u0074\u004e\u0075\u006d\u0062e\u0072\u0073"
	_aef := make([]int, _fbdc)
	if _fbdc > 0 {
		_abb.RTSegments = make([]*Header, _fbdc)
		var (
			_gefb uint64
			_cfa  error
		)
		for _dag := 0; _dag < _fbdc; _dag++ {
			_gefb, _cfa = _cbba.ReadBits(byte(_abb.referenceSize()) << 3)
			if _cfa != nil {
				return nil, _ac.Wrapf(_cfa, _dbagg, "\u0027\u0025\u0064\u0027 \u0072\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0020\u0073e\u0067m\u0065\u006e\u0074\u0020\u006e\u0075\u006db\u0065\u0072", _dag)
			}
			_aef[_dag] = int(_gefb & _f.MaxInt32)
		}
	}
	return _aef, nil
}
func (_aaafa *SymbolDictionary) encodeSymbols(_afcc _g.BinaryWriter) (_fdbe int, _dccf error) {
	const _deefg = "\u0065\u006e\u0063o\u0064\u0065\u0053\u0079\u006d\u0062\u006f\u006c"
	_eadd := _cg.New()
	_eadd.Init()
	_bcgf, _dccf := _aaafa._dfa.SelectByIndexes(_aaafa._eaac)
	if _dccf != nil {
		return 0, _ac.Wrap(_dccf, _deefg, "\u0069n\u0069\u0074\u0069\u0061\u006c")
	}
	_begg := map[*_c.Bitmap]int{}
	for _bgaf, _fddgc := range _bcgf.Values {
		_begg[_fddgc] = _bgaf
	}
	_bcgf.SortByHeight()
	var _fccf, _cdadf int
	_cdcb, _dccf := _bcgf.GroupByHeight()
	if _dccf != nil {
		return 0, _ac.Wrap(_dccf, _deefg, "")
	}
	for _, _fcee := range _cdcb.Values {
		_eeabb := _fcee.Values[0].Height
		_ffgbe := _eeabb - _fccf
		if _dccf = _eadd.EncodeInteger(_cg.IADH, _ffgbe); _dccf != nil {
			return 0, _ac.Wrapf(_dccf, _deefg, "\u0049\u0041\u0044\u0048\u0020\u0066\u006f\u0072\u0020\u0064\u0068\u003a \u0027\u0025\u0064\u0027", _ffgbe)
		}
		_fccf = _eeabb
		_ebfc, _fda := _fcee.GroupByWidth()
		if _fda != nil {
			return 0, _ac.Wrapf(_fda, _deefg, "\u0068\u0065\u0069g\u0068\u0074\u003a\u0020\u0027\u0025\u0064\u0027", _eeabb)
		}
		var _becg int
		for _, _bcea := range _ebfc.Values {
			for _, _aabff := range _bcea.Values {
				_ceba := _aabff.Width
				_ggb := _ceba - _becg
				if _fda = _eadd.EncodeInteger(_cg.IADW, _ggb); _fda != nil {
					return 0, _ac.Wrapf(_fda, _deefg, "\u0049\u0041\u0044\u0057\u0020\u0066\u006f\u0072\u0020\u0064\u0077\u003a \u0027\u0025\u0064\u0027", _ggb)
				}
				_becg += _ggb
				if _fda = _eadd.EncodeBitmap(_aabff, false); _fda != nil {
					return 0, _ac.Wrapf(_fda, _deefg, "H\u0065i\u0067\u0068\u0074\u003a\u0020\u0025\u0064\u0020W\u0069\u0064\u0074\u0068: \u0025\u0064", _eeabb, _ceba)
				}
				_daae := _begg[_aabff]
				_aaafa._faeae[_daae] = _cdadf
				_cdadf++
			}
		}
		if _fda = _eadd.EncodeOOB(_cg.IADW); _fda != nil {
			return 0, _ac.Wrap(_fda, _deefg, "\u0049\u0041\u0044\u0057")
		}
	}
	if _dccf = _eadd.EncodeInteger(_cg.IAEX, 0); _dccf != nil {
		return 0, _ac.Wrap(_dccf, _deefg, "\u0065\u0078p\u006f\u0072\u0074e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	if _dccf = _eadd.EncodeInteger(_cg.IAEX, len(_aaafa._eaac)); _dccf != nil {
		return 0, _ac.Wrap(_dccf, _deefg, "\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0073\u0079m\u0062\u006f\u006c\u0073")
	}
	_eadd.Final()
	_bfg, _dccf := _eadd.WriteTo(_afcc)
	if _dccf != nil {
		return 0, _ac.Wrap(_dccf, _deefg, "\u0077\u0072i\u0074\u0069\u006e\u0067 \u0065\u006ec\u006f\u0064\u0065\u0072\u0020\u0063\u006f\u006et\u0065\u0078\u0074\u0020\u0074\u006f\u0020\u0027\u0077\u0027\u0020\u0077r\u0069\u0074\u0065\u0072")
	}
	return int(_bfg), nil
}
func (_cad *SymbolDictionary) getSbSymCodeLen() int8 {
	_bfca := int8(_f.Ceil(_f.Log(float64(_cad._fgfc+_cad.NumberOfNewSymbols)) / _f.Log(2)))
	if _cad.IsHuffmanEncoded && _bfca < 1 {
		return 1
	}
	return _bfca
}
func (_cdbgb *TextRegion) decodeSymInRefSize() (int64, error) {
	const _cfgd = "\u0064e\u0063o\u0064\u0065\u0053\u0079\u006dI\u006e\u0052e\u0066\u0053\u0069\u007a\u0065"
	if _cdbgb.SbHuffRSize == 0 {
		_ggfd, _afec := _aag.GetStandardTable(1)
		if _afec != nil {
			return 0, _ac.Wrap(_afec, _cfgd, "")
		}
		return _ggfd.Decode(_cdbgb._ceegd)
	}
	if _cdbgb._ddfff == nil {
		var (
			_fbcdd int
			_ffgda error
		)
		if _cdbgb.SbHuffFS == 3 {
			_fbcdd++
		}
		if _cdbgb.SbHuffDS == 3 {
			_fbcdd++
		}
		if _cdbgb.SbHuffDT == 3 {
			_fbcdd++
		}
		if _cdbgb.SbHuffRDWidth == 3 {
			_fbcdd++
		}
		if _cdbgb.SbHuffRDHeight == 3 {
			_fbcdd++
		}
		if _cdbgb.SbHuffRDX == 3 {
			_fbcdd++
		}
		if _cdbgb.SbHuffRDY == 3 {
			_fbcdd++
		}
		_cdbgb._ddfff, _ffgda = _cdbgb.getUserTable(_fbcdd)
		if _ffgda != nil {
			return 0, _ac.Wrap(_ffgda, _cfgd, "")
		}
	}
	_cafbg, _fgdbc := _cdbgb._ddfff.Decode(_cdbgb._ceegd)
	if _fgdbc != nil {
		return 0, _ac.Wrap(_fgdbc, _cfgd, "")
	}
	return _cafbg, nil
}

type PageInformationSegment struct {
	_fabba            _g.StreamReader
	PageBMHeight      int
	PageBMWidth       int
	ResolutionX       int
	ResolutionY       int
	_bddd             bool
	_cba              _c.CombinationOperator
	_cabb             bool
	DefaultPixelValue uint8
	_adfb             bool
	IsLossless        bool
	IsStripe          bool
	MaxStripeSize     uint16
}

func (_dfbf *Header) parse(_gbbf Documenter, _bcfa _g.StreamReader, _bafc int64, _aega OrganizationType) (_egcfb error) {
	const _afc = "\u0070\u0061\u0072s\u0065"
	_gb.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u002d\u0048E\u0041\u0044\u0045\u0052\u005d\u005b\u0050A\u0052\u0053\u0045\u005d\u0020\u0042\u0065\u0067\u0069\u006e\u0073")
	defer func() {
		if _egcfb != nil {
			_gb.Log.Trace("\u005b\u0053\u0045GM\u0045\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045R\u005d[\u0050A\u0052S\u0045\u005d\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _egcfb)
		} else {
			_gb.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052]\u005bP\u0041\u0052\u0053\u0045\u005d\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
		}
	}()
	_, _egcfb = _bcfa.Seek(_bafc, _de.SeekStart)
	if _egcfb != nil {
		return _ac.Wrap(_egcfb, _afc, "\u0073\u0065\u0065\u006b\u0020\u0073\u0074\u0061\u0072\u0074")
	}
	if _egcfb = _dfbf.readSegmentNumber(_bcfa); _egcfb != nil {
		return _ac.Wrap(_egcfb, _afc, "")
	}
	if _egcfb = _dfbf.readHeaderFlags(); _egcfb != nil {
		return _ac.Wrap(_egcfb, _afc, "")
	}
	var _ebef uint64
	_ebef, _egcfb = _dfbf.readNumberOfReferredToSegments(_bcfa)
	if _egcfb != nil {
		return _ac.Wrap(_egcfb, _afc, "")
	}
	_dfbf.RTSNumbers, _egcfb = _dfbf.readReferredToSegmentNumbers(_bcfa, int(_ebef))
	if _egcfb != nil {
		return _ac.Wrap(_egcfb, _afc, "")
	}
	_egcfb = _dfbf.readSegmentPageAssociation(_gbbf, _bcfa, _ebef, _dfbf.RTSNumbers...)
	if _egcfb != nil {
		return _ac.Wrap(_egcfb, _afc, "")
	}
	if _dfbf.Type != TEndOfFile {
		if _egcfb = _dfbf.readSegmentDataLength(_bcfa); _egcfb != nil {
			return _ac.Wrap(_egcfb, _afc, "")
		}
	}
	_dfbf.readDataStartOffset(_bcfa, _aega)
	_dfbf.readHeaderLength(_bcfa, _bafc)
	_gb.Log.Trace("\u0025\u0073", _dfbf)
	return nil
}
func (_ebeg *SymbolDictionary) setSymbolsArray() error {
	if _ebeg._gdc == nil {
		if _edde := _ebeg.retrieveImportSymbols(); _edde != nil {
			return _edde
		}
	}
	if _ebeg._deeb == nil {
		_ebeg._deeb = append(_ebeg._deeb, _ebeg._gdc...)
	}
	return nil
}
func (_bcaed *Header) readSegmentPageAssociation(_degbb Documenter, _gbcg _g.StreamReader, _gbg uint64, _ceb ...int) (_dgbd error) {
	const _feeb = "\u0072\u0065\u0061\u0064\u0053\u0065\u0067\u006d\u0065\u006e\u0074P\u0061\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061t\u0069\u006f\u006e"
	if !_bcaed.PageAssociationFieldSize {
		_ebgb, _ddcbc := _gbcg.ReadBits(8)
		if _ddcbc != nil {
			return _ac.Wrap(_ddcbc, _feeb, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		_bcaed.PageAssociation = int(_ebgb & 0xFF)
	} else {
		_gaee, _eagef := _gbcg.ReadBits(32)
		if _eagef != nil {
			return _ac.Wrap(_eagef, _feeb, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		_bcaed.PageAssociation = int(_gaee & _f.MaxInt32)
	}
	if _gbg == 0 {
		return nil
	}
	if _bcaed.PageAssociation != 0 {
		_bccb, _ddff := _degbb.GetPage(_bcaed.PageAssociation)
		if _ddff != nil {
			return _ac.Wrap(_ddff, _feeb, "\u0061s\u0073\u006f\u0063\u0069a\u0074\u0065\u0064\u0020\u0070a\u0067e\u0020n\u006f\u0074\u0020\u0066\u006f\u0075\u006ed")
		}
		var _cdbc int
		for _gacc := uint64(0); _gacc < _gbg; _gacc++ {
			_cdbc = _ceb[_gacc]
			_bcaed.RTSegments[_gacc], _ddff = _bccb.GetSegment(_cdbc)
			if _ddff != nil {
				var _afdd error
				_bcaed.RTSegments[_gacc], _afdd = _degbb.GetGlobalSegment(_cdbc)
				if _afdd != nil {
					return _ac.Wrapf(_ddff, _feeb, "\u0072\u0065\u0066\u0065\u0072\u0065n\u0063\u0065\u0020s\u0065\u0067\u006de\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0020\u0061\u0074\u0020pa\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0072\u0020\u0069\u006e\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0073", _bcaed.PageAssociation)
				}
			}
		}
		return nil
	}
	for _ffeg := uint64(0); _ffeg < _gbg; _ffeg++ {
		_bcaed.RTSegments[_ffeg], _dgbd = _degbb.GetGlobalSegment(_ceb[_ffeg])
		if _dgbd != nil {
			return _ac.Wrapf(_dgbd, _feeb, "\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u003a\u0020\u0027\u0025d\u0027\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", _ceb[_ffeg])
		}
	}
	return nil
}
func (_debd *PageInformationSegment) readIsLossless() error {
	_cgbb, _dad := _debd._fabba.ReadBit()
	if _dad != nil {
		return _dad
	}
	if _cgbb == 1 {
		_debd.IsLossless = true
	}
	return nil
}
func (_bdgf *SymbolDictionary) setInSyms() error {
	if _bdgf.Header.RTSegments != nil {
		return _bdgf.retrieveImportSymbols()
	}
	_bdgf._gdc = make([]*_c.Bitmap, 0)
	return nil
}

type TableSegment struct {
	_fafc  _g.StreamReader
	_dgecf int32
	_gagg  int32
	_fgfe  int32
	_afgc  int32
	_bed   int32
}

func (_bbega *SymbolDictionary) parseHeader() (_cag error) {
	_gb.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c \u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		if _cag != nil {
			_gb.Log.Trace("\u005bS\u0059\u004dB\u004f\u004c\u0020\u0044I\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u005b\u0050\u0041\u0052SE\u002d\u0048\u0045A\u0044\u0045R\u005d\u0020\u0066\u0061\u0069\u006ce\u0064\u002e \u0025\u0076", _cag)
		} else {
			_gb.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c \u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		}
	}()
	if _cag = _bbega.readRegionFlags(); _cag != nil {
		return _cag
	}
	if _cag = _bbega.setAtPixels(); _cag != nil {
		return _cag
	}
	if _cag = _bbega.setRefinementAtPixels(); _cag != nil {
		return _cag
	}
	if _cag = _bbega.readNumberOfExportedSymbols(); _cag != nil {
		return _cag
	}
	if _cag = _bbega.readNumberOfNewSymbols(); _cag != nil {
		return _cag
	}
	if _cag = _bbega.setInSyms(); _cag != nil {
		return _cag
	}
	if _bbega._bcce {
		_dgec := _bbega.Header.RTSegments
		for _gbed := len(_dgec) - 1; _gbed >= 0; _gbed-- {
			if _dgec[_gbed].Type == 0 {
				_fbcc, _dbb := _dgec[_gbed].SegmentData.(*SymbolDictionary)
				if !_dbb {
					_cag = _be.Errorf("\u0072\u0065\u006c\u0061\u0074\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074:\u0020\u0025\u0076\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020S\u0079\u006d\u0062\u006f\u006c\u0020\u0044\u0069\u0063\u0074\u0069\u006fna\u0072\u0079\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074", _dgec[_gbed])
					return _cag
				}
				if _fbcc._bcce {
					_bbega.setRetainedCodingContexts(_fbcc)
				}
				break
			}
		}
	}
	if _cag = _bbega.checkInput(); _cag != nil {
		return _cag
	}
	return nil
}
func (_agdg *SymbolDictionary) retrieveImportSymbols() error {
	for _, _dbdg := range _agdg.Header.RTSegments {
		if _dbdg.Type == 0 {
			_fgec, _decc := _dbdg.GetSegmentData()
			if _decc != nil {
				return _decc
			}
			_gfac, _gcad := _fgec.(*SymbolDictionary)
			if !_gcad {
				return _be.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0044\u0061\u0074a\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0053\u0065\u0067m\u0065\u006e\u0074\u003a\u0020%\u0054", _fgec)
			}
			_gedba, _decc := _gfac.GetDictionary()
			if _decc != nil {
				return _be.Errorf("\u0072\u0065\u006c\u0061\u0074\u0065\u0064 \u0073\u0065\u0067m\u0065\u006e\u0074 \u0077\u0069t\u0068\u0020\u0069\u006e\u0064\u0065x\u003a %\u0064\u0020\u0067\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0073", _dbdg.SegmentNumber, _decc.Error())
			}
			_agdg._gdc = append(_agdg._gdc, _gedba...)
			_agdg._fgfc += _gfac.NumberOfExportedSymbols
		}
	}
	return nil
}
func (_cafb *TextRegion) decodeIds() (int64, error) {
	const _badca = "\u0064e\u0063\u006f\u0064\u0065\u0049\u0064s"
	if _cafb.IsHuffmanEncoded {
		if _cafb.SbHuffDS == 3 {
			if _cafb._eabed == nil {
				_bdbe := 0
				if _cafb.SbHuffFS == 3 {
					_bdbe++
				}
				var _ffgf error
				_cafb._eabed, _ffgf = _cafb.getUserTable(_bdbe)
				if _ffgf != nil {
					return 0, _ac.Wrap(_ffgf, _badca, "")
				}
			}
			return _cafb._eabed.Decode(_cafb._ceegd)
		}
		_fadd, _dccbc := _aag.GetStandardTable(8 + int(_cafb.SbHuffDS))
		if _dccbc != nil {
			return 0, _ac.Wrap(_dccbc, _badca, "")
		}
		return _fadd.Decode(_cafb._ceegd)
	}
	_fafe, _dgcb := _cafb._bbcea.DecodeInt(_cafb._bcca)
	if _dgcb != nil {
		return 0, _ac.Wrap(_dgcb, _badca, "\u0063\u0078\u0049\u0041\u0044\u0053")
	}
	return int64(_fafe), nil
}

type GenericRefinementRegion struct {
	_fa             templater
	_bcab           templater
	_ea             _g.StreamReader
	_ff             *Header
	RegionInfo      *RegionSegment
	IsTPGROn        bool
	TemplateID      int8
	Template        templater
	GrAtX           []int8
	GrAtY           []int8
	RegionBitmap    *_c.Bitmap
	ReferenceBitmap *_c.Bitmap
	ReferenceDX     int32
	ReferenceDY     int32
	_fb             *_aa.Decoder
	_ef             *_aa.DecoderStats
	_age            bool
	_gbd            []bool
}

func (_gg *GenericRefinementRegion) decodeOptimized(_gge, _cga, _fae, _ge, _bg, _bfc, _eec int) error {
	var (
		_eeb error
		_eae int
		_aad int
	)
	_ab := _gge - int(_gg.ReferenceDY)
	if _fg := int(-_gg.ReferenceDX); _fg > 0 {
		_eae = _fg
	}
	_dbd := _gg.ReferenceBitmap.GetByteIndex(_eae, _ab)
	if _gg.ReferenceDX > 0 {
		_aad = int(_gg.ReferenceDX)
	}
	_dc := _gg.RegionBitmap.GetByteIndex(_aad, _gge)
	switch _gg.TemplateID {
	case 0:
		_eeb = _gg.decodeTemplate(_gge, _cga, _fae, _ge, _bg, _bfc, _eec, _dc, _ab, _dbd, _gg._fa)
	case 1:
		_eeb = _gg.decodeTemplate(_gge, _cga, _fae, _ge, _bg, _bfc, _eec, _dc, _ab, _dbd, _gg._bcab)
	}
	return _eeb
}
func (_cccb *PageInformationSegment) Init(h *Header, r _g.StreamReader) (_gcd error) {
	_cccb._fabba = r
	if _gcd = _cccb.parseHeader(); _gcd != nil {
		return _ac.Wrap(_gcd, "P\u0061\u0067\u0065\u0049\u006e\u0066o\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065g\u006d\u0065\u006et\u002eI\u006e\u0069\u0074", "")
	}
	return nil
}

const (
	ORandom OrganizationType = iota
	OSequential
)

func (_fffg *SymbolDictionary) setRetainedCodingContexts(_bdad *SymbolDictionary) {
	_fffg._caaa = _bdad._caaa
	_fffg.IsHuffmanEncoded = _bdad.IsHuffmanEncoded
	_fffg.UseRefinementAggregation = _bdad.UseRefinementAggregation
	_fffg.SdTemplate = _bdad.SdTemplate
	_fffg.SdrTemplate = _bdad.SdrTemplate
	_fffg.SdATX = _bdad.SdATX
	_fffg.SdATY = _bdad.SdATY
	_fffg.SdrATX = _bdad.SdrATX
	_fffg.SdrATY = _bdad.SdrATY
	_fffg._ebff = _bdad._ebff
}
func (_cace *PatternDictionary) GetDictionary() ([]*_c.Bitmap, error) {
	if _cace.Patterns != nil {
		return _cace.Patterns, nil
	}
	if !_cace.IsMMREncoded {
		_cace.setGbAtPixels()
	}
	_geeca := NewGenericRegion(_cace._fad)
	_geeca.setParametersMMR(_cace.IsMMREncoded, _cace.DataOffset, _cace.DataLength, uint32(_cace.HdpHeight), (_cace.GrayMax+1)*uint32(_cace.HdpWidth), _cace.HDTemplate, false, false, _cace.GBAtX, _cace.GBAtY)
	_cebf, _bgaee := _geeca.GetRegionBitmap()
	if _bgaee != nil {
		return nil, _bgaee
	}
	if _bgaee = _cace.extractPatterns(_cebf); _bgaee != nil {
		return nil, _bgaee
	}
	return _cace.Patterns, nil
}
func (_gddag *SymbolDictionary) decodeAggregate(_daga, _eefb uint32) error {
	var (
		_aacf int64
		_agca error
	)
	if _gddag.IsHuffmanEncoded {
		_aacf, _agca = _gddag.huffDecodeRefAggNInst()
		if _agca != nil {
			return _agca
		}
	} else {
		_ebgbb, _bggd := _gddag._caaa.DecodeInt(_gddag._efab)
		if _bggd != nil {
			return _bggd
		}
		_aacf = int64(_ebgbb)
	}
	if _aacf > 1 {
		return _gddag.decodeThroughTextRegion(_daga, _eefb, uint32(_aacf))
	} else if _aacf == 1 {
		return _gddag.decodeRefinedSymbol(_daga, _eefb)
	}
	return nil
}

var _ _aag.BasicTabler = &TableSegment{}

func (_dbc *PageInformationSegment) parseHeader() (_bcaag error) {
	_gb.Log.Trace("\u005b\u0050\u0061\u0067\u0065I\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065\u0067m\u0065\u006e\u0074\u005d\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065\u0061\u0064\u0065\u0072\u002e\u002e\u002e")
	defer func() {
		var _bffdb = "[\u0050\u0061\u0067\u0065\u0049\u006e\u0066\u006f\u0072m\u0061\u0074\u0069\u006f\u006e\u0053\u0065gm\u0065\u006e\u0074\u005d \u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065ad\u0065\u0072 \u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064"
		if _bcaag != nil {
			_bffdb += "\u0020\u0077\u0069t\u0068\u0020\u0065\u0072\u0072\u006f\u0072\u0020" + _bcaag.Error()
		} else {
			_bffdb += "\u0020\u0073\u0075\u0063\u0063\u0065\u0073\u0073\u0066\u0075\u006c\u006c\u0079"
		}
		_gb.Log.Trace(_bffdb)
	}()
	if _bcaag = _dbc.readWidthAndHeight(); _bcaag != nil {
		return _bcaag
	}
	if _bcaag = _dbc.readResolution(); _bcaag != nil {
		return _bcaag
	}
	_, _bcaag = _dbc._fabba.ReadBit()
	if _bcaag != nil {
		return _bcaag
	}
	if _bcaag = _dbc.readCombinationOperatorOverrideAllowed(); _bcaag != nil {
		return _bcaag
	}
	if _bcaag = _dbc.readRequiresAuxiliaryBuffer(); _bcaag != nil {
		return _bcaag
	}
	if _bcaag = _dbc.readCombinationOperator(); _bcaag != nil {
		return _bcaag
	}
	if _bcaag = _dbc.readDefaultPixelValue(); _bcaag != nil {
		return _bcaag
	}
	if _bcaag = _dbc.readContainsRefinement(); _bcaag != nil {
		return _bcaag
	}
	if _bcaag = _dbc.readIsLossless(); _bcaag != nil {
		return _bcaag
	}
	if _bcaag = _dbc.readIsStriped(); _bcaag != nil {
		return _bcaag
	}
	if _bcaag = _dbc.readMaxStripeSize(); _bcaag != nil {
		return _bcaag
	}
	if _bcaag = _dbc.checkInput(); _bcaag != nil {
		return _bcaag
	}
	_gb.Log.Trace("\u0025\u0073", _dbc)
	return nil
}
func (_aab *GenericRefinementRegion) Init(header *Header, r _g.StreamReader) error {
	_aab._ff = header
	_aab._ea = r
	_aab.RegionInfo = NewRegionSegment(r)
	return _aab.parseHeader()
}
func (_abef *GenericRegion) Init(h *Header, r _g.StreamReader) error {
	_abef.RegionSegment = NewRegionSegment(r)
	_abef._degf = r
	return _abef.parseHeader()
}
func (_ggfgg *PageInformationSegment) readWidthAndHeight() error {
	_beef, _ffec := _ggfgg._fabba.ReadBits(32)
	if _ffec != nil {
		return _ffec
	}
	_ggfgg.PageBMWidth = int(_beef & _f.MaxInt32)
	_beef, _ffec = _ggfgg._fabba.ReadBits(32)
	if _ffec != nil {
		return _ffec
	}
	_ggfgg.PageBMHeight = int(_beef & _f.MaxInt32)
	return nil
}
func (_ddg *GenericRefinementRegion) decodeTemplate(_bcd, _dabd, _ed, _ecd, _beb, _ggf, _bfeb, _cfb, _bbf, _aed int, _cfg templater) (_ffd error) {
	var (
		_ecdc, _cgad, _gfc, _dec, _egc int16
		_gbc, _bbe, _bfcc, _dcc        int
		_bcaa                          byte
	)
	if _bbf >= 1 && (_bbf-1) < _ddg.ReferenceBitmap.Height {
		_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed - _ecd)
		if _ffd != nil {
			return
		}
		_gbc = int(_bcaa)
	}
	if _bbf >= 0 && (_bbf) < _ddg.ReferenceBitmap.Height {
		_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed)
		if _ffd != nil {
			return
		}
		_bbe = int(_bcaa)
	}
	if _bbf >= -1 && (_bbf+1) < _ddg.ReferenceBitmap.Height {
		_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed + _ecd)
		if _ffd != nil {
			return
		}
		_bfcc = int(_bcaa)
	}
	_aed++
	if _bcd >= 1 {
		_bcaa, _ffd = _ddg.RegionBitmap.GetByte(_cfb - _ed)
		if _ffd != nil {
			return
		}
		_dcc = int(_bcaa)
	}
	_cfb++
	_ddge := _ddg.ReferenceDX % 8
	_eca := 6 + _ddge
	_bae := _aed % _ecd
	if _eca >= 0 {
		if _eca < 8 {
			_ecdc = int16(_gbc>>uint(_eca)) & 0x07
		}
		if _eca < 8 {
			_cgad = int16(_bbe>>uint(_eca)) & 0x07
		}
		if _eca < 8 {
			_gfc = int16(_bfcc>>uint(_eca)) & 0x07
		}
		if _eca == 6 && _bae > 1 {
			if _bbf >= 1 && (_bbf-1) < _ddg.ReferenceBitmap.Height {
				_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed - _ecd - 2)
				if _ffd != nil {
					return _ffd
				}
				_ecdc |= int16(_bcaa<<2) & 0x04
			}
			if _bbf >= 0 && _bbf < _ddg.ReferenceBitmap.Height {
				_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed - 2)
				if _ffd != nil {
					return _ffd
				}
				_cgad |= int16(_bcaa<<2) & 0x04
			}
			if _bbf >= -1 && _bbf+1 < _ddg.ReferenceBitmap.Height {
				_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed + _ecd - 2)
				if _ffd != nil {
					return _ffd
				}
				_gfc |= int16(_bcaa<<2) & 0x04
			}
		}
		if _eca == 0 {
			_gbc = 0
			_bbe = 0
			_bfcc = 0
			if _bae < _ecd-1 {
				if _bbf >= 1 && _bbf-1 < _ddg.ReferenceBitmap.Height {
					_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed - _ecd)
					if _ffd != nil {
						return _ffd
					}
					_gbc = int(_bcaa)
				}
				if _bbf >= 0 && _bbf < _ddg.ReferenceBitmap.Height {
					_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed)
					if _ffd != nil {
						return _ffd
					}
					_bbe = int(_bcaa)
				}
				if _bbf >= -1 && _bbf+1 < _ddg.ReferenceBitmap.Height {
					_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed + _ecd)
					if _ffd != nil {
						return _ffd
					}
					_bfcc = int(_bcaa)
				}
			}
			_aed++
		}
	} else {
		_ecdc = int16(_gbc<<1) & 0x07
		_cgad = int16(_bbe<<1) & 0x07
		_gfc = int16(_bfcc<<1) & 0x07
		_gbc = 0
		_bbe = 0
		_bfcc = 0
		if _bae < _ecd-1 {
			if _bbf >= 1 && _bbf-1 < _ddg.ReferenceBitmap.Height {
				_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed - _ecd)
				if _ffd != nil {
					return _ffd
				}
				_gbc = int(_bcaa)
			}
			if _bbf >= 0 && _bbf < _ddg.ReferenceBitmap.Height {
				_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed)
				if _ffd != nil {
					return _ffd
				}
				_bbe = int(_bcaa)
			}
			if _bbf >= -1 && _bbf+1 < _ddg.ReferenceBitmap.Height {
				_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed + _ecd)
				if _ffd != nil {
					return _ffd
				}
				_bfcc = int(_bcaa)
			}
			_aed++
		}
		_ecdc |= int16((_gbc >> 7) & 0x07)
		_cgad |= int16((_bbe >> 7) & 0x07)
		_gfc |= int16((_bfcc >> 7) & 0x07)
	}
	_dec = int16(_dcc >> 6)
	_egc = 0
	_fbdd := (2 - _ddge) % 8
	_gbc <<= uint(_fbdd)
	_bbe <<= uint(_fbdd)
	_bfcc <<= uint(_fbdd)
	_dcc <<= 2
	var _ad int
	for _add := 0; _add < _dabd; _add++ {
		_cgf := _add & 0x07
		_agge := _cfg.form(_ecdc, _cgad, _gfc, _dec, _egc)
		if _ddg._age {
			_bcaa, _ffd = _ddg.RegionBitmap.GetByte(_ddg.RegionBitmap.GetByteIndex(_add, _bcd))
			if _ffd != nil {
				return _ffd
			}
			_ddg._ef.SetIndex(int32(_ddg.overrideAtTemplate0(int(_agge), _add, _bcd, int(_bcaa), _cgf)))
		} else {
			_ddg._ef.SetIndex(int32(_agge))
		}
		_ad, _ffd = _ddg._fb.DecodeBit(_ddg._ef)
		if _ffd != nil {
			return _ffd
		}
		if _ffd = _ddg.RegionBitmap.SetPixel(_add, _bcd, byte(_ad)); _ffd != nil {
			return _ffd
		}
		_ecdc = ((_ecdc << 1) | 0x01&int16(_gbc>>7)) & 0x07
		_cgad = ((_cgad << 1) | 0x01&int16(_bbe>>7)) & 0x07
		_gfc = ((_gfc << 1) | 0x01&int16(_bfcc>>7)) & 0x07
		_dec = ((_dec << 1) | 0x01&int16(_dcc>>7)) & 0x07
		_egc = int16(_ad)
		if (_add-int(_ddg.ReferenceDX))%8 == 5 {
			_gbc = 0
			_bbe = 0
			_bfcc = 0
			if ((_add-int(_ddg.ReferenceDX))/8)+1 < _ddg.ReferenceBitmap.RowStride {
				if _bbf >= 1 && (_bbf-1) < _ddg.ReferenceBitmap.Height {
					_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed - _ecd)
					if _ffd != nil {
						return _ffd
					}
					_gbc = int(_bcaa)
				}
				if _bbf >= 0 && _bbf < _ddg.ReferenceBitmap.Height {
					_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed)
					if _ffd != nil {
						return _ffd
					}
					_bbe = int(_bcaa)
				}
				if _bbf >= -1 && (_bbf+1) < _ddg.ReferenceBitmap.Height {
					_bcaa, _ffd = _ddg.ReferenceBitmap.GetByte(_aed + _ecd)
					if _ffd != nil {
						return _ffd
					}
					_bfcc = int(_bcaa)
				}
			}
			_aed++
		} else {
			_gbc <<= 1
			_bbe <<= 1
			_bfcc <<= 1
		}
		if _cgf == 5 && _bcd >= 1 {
			if ((_add >> 3) + 1) >= _ddg.RegionBitmap.RowStride {
				_dcc = 0
			} else {
				_bcaa, _ffd = _ddg.RegionBitmap.GetByte(_cfb - _ed)
				if _ffd != nil {
					return _ffd
				}
				_dcc = int(_bcaa)
			}
			_cfb++
		} else {
			_dcc <<= 1
		}
	}
	return nil
}

type EncodeInitializer interface{ InitEncode() }

func (_db *EndOfStripe) Init(h *Header, r _g.StreamReader) error {
	_db._fd = r
	return _db.parseHeader(h, r)
}
func (_bbge *RegionSegment) Size() int { return 17 }
func NewHeader(d Documenter, r _g.StreamReader, offset int64, organizationType OrganizationType) (*Header, error) {
	_ffgab := &Header{Reader: r}
	if _cgeg := _ffgab.parse(d, r, offset, organizationType); _cgeg != nil {
		return nil, _ac.Wrap(_cgeg, "\u004ee\u0077\u0048\u0065\u0061\u0064\u0065r", "")
	}
	return _ffgab, nil
}

type Segmenter interface {
	Init(_gbe *Header, _gedf _g.StreamReader) error
}

func (_cfee *TableSegment) HtHigh() int32 { return _cfee._bed }
func (_geed *TextRegion) readAmountOfSymbolInstances() error {
	_geddc, _bfcgc := _geed._ceegd.ReadBits(32)
	if _bfcgc != nil {
		return _bfcgc
	}
	_geed.NumberOfSymbolInstances = uint32(_geddc & _f.MaxUint32)
	_bdceg := _geed.RegionInfo.BitmapWidth * _geed.RegionInfo.BitmapHeight
	if _bdceg < _geed.NumberOfSymbolInstances {
		_gb.Log.Debug("\u004c\u0069\u006d\u0069t\u0069\u006e\u0067\u0020t\u0068\u0065\u0020n\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020d\u0065\u0063\u006f\u0064e\u0064\u0020\u0073\u0079m\u0062\u006f\u006c\u0020\u0069n\u0073\u0074\u0061\u006e\u0063\u0065\u0073 \u0074\u006f\u0020\u006f\u006ee\u0020\u0070\u0065\u0072\u0020\u0070\u0069\u0078\u0065l\u0020\u0028\u0020\u0025\u0064\u0020\u0069\u006e\u0073\u0074\u0065\u0061\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u0029", _bdceg, _geed.NumberOfSymbolInstances)
		_geed.NumberOfSymbolInstances = _bdceg
	}
	return nil
}
func (_ddfdd *SymbolDictionary) decodeRefinedSymbol(_cfca, _cacde uint32) error {
	var (
		_fcefe       int
		_ebag, _edfg int32
	)
	if _ddfdd.IsHuffmanEncoded {
		_gbfg, _ggga := _ddfdd._aege.ReadBits(byte(_ddfdd._ebafc))
		if _ggga != nil {
			return _ggga
		}
		_fcefe = int(_gbfg)
		_fcdd, _ggga := _aag.GetStandardTable(15)
		if _ggga != nil {
			return _ggga
		}
		_cbae, _ggga := _fcdd.Decode(_ddfdd._aege)
		if _ggga != nil {
			return _ggga
		}
		_ebag = int32(_cbae)
		_cbae, _ggga = _fcdd.Decode(_ddfdd._aege)
		if _ggga != nil {
			return _ggga
		}
		_edfg = int32(_cbae)
		_fcdd, _ggga = _aag.GetStandardTable(1)
		if _ggga != nil {
			return _ggga
		}
		if _, _ggga = _fcdd.Decode(_ddfdd._aege); _ggga != nil {
			return _ggga
		}
		_ddfdd._aege.Align()
	} else {
		_cgbgc, _cfbc := _ddfdd._caaa.DecodeIAID(uint64(_ddfdd._ebafc), _ddfdd._dcfa)
		if _cfbc != nil {
			return _cfbc
		}
		_fcefe = int(_cgbgc)
		_ebag, _cfbc = _ddfdd._caaa.DecodeInt(_ddfdd._bcb)
		if _cfbc != nil {
			return _cfbc
		}
		_edfg, _cfbc = _ddfdd._caaa.DecodeInt(_ddfdd._gebb)
		if _cfbc != nil {
			return _cfbc
		}
	}
	if _egcfbf := _ddfdd.setSymbolsArray(); _egcfbf != nil {
		return _egcfbf
	}
	_bdfgf := _ddfdd._deeb[_fcefe]
	if _egccb := _ddfdd.decodeNewSymbols(_cfca, _cacde, _bdfgf, _ebag, _edfg); _egccb != nil {
		return _egccb
	}
	if _ddfdd.IsHuffmanEncoded {
		_ddfdd._aege.Align()
	}
	return nil
}

var _ templater = &template1{}

func (_deb *EndOfStripe) LineNumber() int { return _deb._bf }
func (_dgdf *template0) form(_gbf, _debb, _abd, _gbaef, _geg int16) int16 {
	return (_gbf << 10) | (_debb << 7) | (_abd << 4) | (_gbaef << 1) | _geg
}
func (_dabe *GenericRegion) overrideAtTemplate3(_ageg, _eaad, _aeb, _eee, _dgb int) int {
	_ageg &= 0x3EF
	if _dabe.GBAtY[0] == 0 && _dabe.GBAtX[0] >= -int8(_dgb) {
		_ageg |= (_eee >> uint(7-(int8(_dgb)+_dabe.GBAtX[0])) & 0x1) << 4
	} else {
		_ageg |= int(_dabe.getPixel(_eaad+int(_dabe.GBAtX[0]), _aeb+int(_dabe.GBAtY[0]))) << 4
	}
	return _ageg
}
func (_ba *EndOfStripe) parseHeader(_e *Header, _af _g.StreamReader) error {
	_deg, _afd := _ba._fd.ReadBits(32)
	if _afd != nil {
		return _afd
	}
	_ba._bf = int(_deg & _f.MaxInt32)
	return nil
}

type OrganizationType uint8

func (_debf *TextRegion) checkInput() error {
	const _cbeef = "\u0063\u0068\u0065\u0063\u006b\u0049\u006e\u0070\u0075\u0074"
	if !_debf.UseRefinement {
		if _debf.SbrTemplate != 0 {
			_gb.Log.Debug("\u0053\u0062\u0072Te\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030")
			_debf.SbrTemplate = 0
		}
	}
	if _debf.SbHuffFS == 2 || _debf.SbHuffRDWidth == 2 || _debf.SbHuffRDHeight == 2 || _debf.SbHuffRDX == 2 || _debf.SbHuffRDY == 2 {
		return _ac.Error(_cbeef, "h\u0075\u0066\u0066\u006d\u0061\u006e \u0066\u006c\u0061\u0067\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064")
	}
	if !_debf.UseRefinement {
		if _debf.SbHuffRSize != 0 {
			_gb.Log.Debug("\u0053\u0062\u0048uf\u0066\u0052\u0053\u0069\u007a\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030")
			_debf.SbHuffRSize = 0
		}
		if _debf.SbHuffRDY != 0 {
			_gb.Log.Debug("S\u0062\u0048\u0075\u0066fR\u0044Y\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0030")
			_debf.SbHuffRDY = 0
		}
		if _debf.SbHuffRDX != 0 {
			_gb.Log.Debug("S\u0062\u0048\u0075\u0066fR\u0044X\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0030")
			_debf.SbHuffRDX = 0
		}
		if _debf.SbHuffRDWidth != 0 {
			_gb.Log.Debug("\u0053b\u0048\u0075\u0066\u0066R\u0044\u0057\u0069\u0064\u0074h\u0020s\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u00200")
			_debf.SbHuffRDWidth = 0
		}
		if _debf.SbHuffRDHeight != 0 {
			_gb.Log.Debug("\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0048\u0065\u0069g\u0068\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020b\u0065\u0020\u0030")
			_debf.SbHuffRDHeight = 0
		}
	}
	return nil
}

var (
	_bggb Segmenter
	_geb  = map[Type]func() Segmenter{TSymbolDictionary: func() Segmenter { return &SymbolDictionary{} }, TIntermediateTextRegion: func() Segmenter { return &TextRegion{} }, TImmediateTextRegion: func() Segmenter { return &TextRegion{} }, TImmediateLosslessTextRegion: func() Segmenter { return &TextRegion{} }, TPatternDictionary: func() Segmenter { return &PatternDictionary{} }, TIntermediateHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TImmediateHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TImmediateLosslessHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TIntermediateGenericRegion: func() Segmenter { return &GenericRegion{} }, TImmediateGenericRegion: func() Segmenter { return &GenericRegion{} }, TImmediateLosslessGenericRegion: func() Segmenter { return &GenericRegion{} }, TIntermediateGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TImmediateGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TImmediateLosslessGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TPageInformation: func() Segmenter { return &PageInformationSegment{} }, TEndOfPage: func() Segmenter { return _bggb }, TEndOfStrip: func() Segmenter { return &EndOfStripe{} }, TEndOfFile: func() Segmenter { return _bggb }, TProfiles: func() Segmenter { return _bggb }, TTables: func() Segmenter { return &TableSegment{} }, TExtension: func() Segmenter { return _bggb }, TBitmap: func() Segmenter { return _bggb }}
)

func (_feec *PatternDictionary) readPatternWidthAndHeight() error {
	_eega, _bdcg := _feec._fad.ReadByte()
	if _bdcg != nil {
		return _bdcg
	}
	_feec.HdpWidth = _eega
	_eega, _bdcg = _feec._fad.ReadByte()
	if _bdcg != nil {
		return _bdcg
	}
	_feec.HdpHeight = _eega
	return nil
}
func (_gdee *SymbolDictionary) decodeDirectlyThroughGenericRegion(_dbda, _dadg uint32) error {
	if _gdee._eaed == nil {
		_gdee._eaed = NewGenericRegion(_gdee._aege)
	}
	_gdee._eaed.setParametersWithAt(false, byte(_gdee.SdTemplate), false, false, _gdee.SdATX, _gdee.SdATY, _dbda, _dadg, _gdee._ebff, _gdee._caaa)
	return _gdee.addSymbol(_gdee._eaed)
}
func (_agfg *PageInformationSegment) encodeStripingInformation(_addb _g.BinaryWriter) (_dcaa int, _bbeda error) {
	const _ffgd = "\u0065n\u0063\u006f\u0064\u0065S\u0074\u0072\u0069\u0070\u0069n\u0067I\u006ef\u006f\u0072\u006d\u0061\u0074\u0069\u006fn"
	if !_agfg.IsStripe {
		if _dcaa, _bbeda = _addb.Write([]byte{0x00, 0x00}); _bbeda != nil {
			return 0, _ac.Wrap(_bbeda, _ffgd, "n\u006f\u0020\u0073\u0074\u0072\u0069\u0070\u0069\u006e\u0067")
		}
		return _dcaa, nil
	}
	_cbee := make([]byte, 2)
	_bc.BigEndian.PutUint16(_cbee, _agfg.MaxStripeSize|1<<15)
	if _dcaa, _bbeda = _addb.Write(_cbee); _bbeda != nil {
		return 0, _ac.Wrapf(_bbeda, _ffgd, "\u0073\u0074\u0072i\u0070\u0069\u006e\u0067\u003a\u0020\u0025\u0064", _agfg.MaxStripeSize)
	}
	return _dcaa, nil
}
func (_dee *HalftoneRegion) computeX(_edfd, _ede int) int {
	return _dee.shiftAndFill(int(_dee.HGridX) + _edfd*int(_dee.HRegionY) + _ede*int(_dee.HRegionX))
}
func (_defc *TextRegion) decodeRdw() (int64, error) {
	const _abaa = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064w"
	if _defc.IsHuffmanEncoded {
		if _defc.SbHuffRDWidth == 3 {
			if _defc._cdcgd == nil {
				var (
					_fffcg int
					_febe  error
				)
				if _defc.SbHuffFS == 3 {
					_fffcg++
				}
				if _defc.SbHuffDS == 3 {
					_fffcg++
				}
				if _defc.SbHuffDT == 3 {
					_fffcg++
				}
				_defc._cdcgd, _febe = _defc.getUserTable(_fffcg)
				if _febe != nil {
					return 0, _ac.Wrap(_febe, _abaa, "")
				}
			}
			return _defc._cdcgd.Decode(_defc._ceegd)
		}
		_eaab, _gecgb := _aag.GetStandardTable(14 + int(_defc.SbHuffRDWidth))
		if _gecgb != nil {
			return 0, _ac.Wrap(_gecgb, _abaa, "")
		}
		return _eaab.Decode(_defc._ceegd)
	}
	_bbdg, _ffbcd := _defc._bbcea.DecodeInt(_defc._bbgdf)
	if _ffbcd != nil {
		return 0, _ac.Wrap(_ffbcd, _abaa, "")
	}
	return int64(_bbdg), nil
}

var _ SegmentEncoder = &RegionSegment{}

func (_edc *Header) CleanSegmentData() {
	if _edc.SegmentData != nil {
		_edc.SegmentData = nil
	}
}
func (_dbfc *TextRegion) Init(header *Header, r _g.StreamReader) error {
	_dbfc.Header = header
	_dbfc._ceegd = r
	_dbfc.RegionInfo = NewRegionSegment(_dbfc._ceegd)
	return _dbfc.parseHeader()
}
func (_bagd *HalftoneRegion) renderPattern(_aebb [][]int) (_fdde error) {
	var _gfecd, _eefa int
	for _daff := 0; _daff < int(_bagd.HGridHeight); _daff++ {
		for _gfgb := 0; _gfgb < int(_bagd.HGridWidth); _gfgb++ {
			_gfecd = _bagd.computeX(_daff, _gfgb)
			_eefa = _bagd.computeY(_daff, _gfgb)
			_efg := _bagd.Patterns[_aebb[_daff][_gfgb]]
			if _fdde = _c.Blit(_efg, _bagd.HalftoneRegionBitmap, _gfecd+int(_bagd.HGridX), _eefa+int(_bagd.HGridY), _bagd.CombinationOperator); _fdde != nil {
				return _fdde
			}
		}
	}
	return nil
}
func (_dcgc *PatternDictionary) readIsMMREncoded() error {
	_gab, _ccd := _dcgc._fad.ReadBit()
	if _ccd != nil {
		return _ccd
	}
	if _gab != 0 {
		_dcgc.IsMMREncoded = true
	}
	return nil
}
func (_bdfg *SymbolDictionary) encodeRefinementATFlags(_fbcd _g.BinaryWriter) (_dffe int, _dacdf error) {
	const _bgdd = "\u0065\u006e\u0063od\u0065\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0041\u0054\u0046\u006c\u0061\u0067\u0073"
	if !_bdfg.UseRefinementAggregation || _bdfg.SdrTemplate != 0 {
		return 0, nil
	}
	for _dfac := 0; _dfac < 2; _dfac++ {
		if _dacdf = _fbcd.WriteByte(byte(_bdfg.SdrATX[_dfac])); _dacdf != nil {
			return _dffe, _ac.Wrapf(_dacdf, _bgdd, "\u0053\u0064\u0072\u0041\u0054\u0058\u005b\u0025\u0064\u005d", _dfac)
		}
		_dffe++
		if _dacdf = _fbcd.WriteByte(byte(_bdfg.SdrATY[_dfac])); _dacdf != nil {
			return _dffe, _ac.Wrapf(_dacdf, _bgdd, "\u0053\u0064\u0072\u0041\u0054\u0059\u005b\u0025\u0064\u005d", _dfac)
		}
		_dffe++
	}
	return _dffe, nil
}
func (_cge *GenericRegion) decodeTemplate3(_aadbg, _cabd, _dbag int, _beea, _ddc int) (_eef error) {
	const _bgf = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0033"
	var (
		_ged, _dcce   int
		_cbbc         int
		_fdg          byte
		_dfdac, _gddf int
	)
	if _aadbg >= 1 {
		_fdg, _eef = _cge.Bitmap.GetByte(_ddc)
		if _eef != nil {
			return _ac.Wrap(_eef, _bgf, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_cbbc = int(_fdg)
	}
	_ged = (_cbbc >> 1) & 0x70
	for _gdf := 0; _gdf < _dbag; _gdf = _dfdac {
		var (
			_fff  byte
			_gcbb int
		)
		_dfdac = _gdf + 8
		if _gaeb := _cabd - _gdf; _gaeb > 8 {
			_gcbb = 8
		} else {
			_gcbb = _gaeb
		}
		if _aadbg >= 1 {
			_cbbc <<= 8
			if _dfdac < _cabd {
				_fdg, _eef = _cge.Bitmap.GetByte(_ddc + 1)
				if _eef != nil {
					return _ac.Wrap(_eef, _bgf, "\u0069\u006e\u006e\u0065\u0072\u0020\u002d\u0020\u006c\u0069\u006e\u0065 \u003e\u003d\u0020\u0031")
				}
				_cbbc |= int(_fdg)
			}
		}
		for _fddb := 0; _fddb < _gcbb; _fddb++ {
			if _cge._aaaf {
				_dcce = _cge.overrideAtTemplate3(_ged, _gdf+_fddb, _aadbg, int(_fff), _fddb)
				_cge._bcae.SetIndex(int32(_dcce))
			} else {
				_cge._bcae.SetIndex(int32(_ged))
			}
			_gddf, _eef = _cge._ga.DecodeBit(_cge._bcae)
			if _eef != nil {
				return _ac.Wrap(_eef, _bgf, "")
			}
			_fff |= byte(_gddf) << byte(7-_fddb)
			_ged = ((_ged & 0x1f7) << 1) | _gddf | ((_cbbc >> uint(8-_fddb)) & 0x010)
		}
		if _ebb := _cge.Bitmap.SetByte(_beea, _fff); _ebb != nil {
			return _ac.Wrap(_ebb, _bgf, "")
		}
		_beea++
		_ddc++
	}
	return nil
}
func (_fcef *SymbolDictionary) String() string {
	_affff := &_bb.Builder{}
	_affff.WriteString("\n\u005b\u0053\u0059\u004dBO\u004c-\u0044\u0049\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u000a")
	_affff.WriteString(_be.Sprintf("\u0009-\u0020S\u0064\u0072\u0054\u0065\u006dp\u006c\u0061t\u0065\u0020\u0025\u0076\u000a", _fcef.SdrTemplate))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u0025\u0076\u000a", _fcef.SdTemplate))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0069\u0073\u0043\u006f\u0064\u0069\u006eg\u0043\u006f\u006e\u0074\u0065\u0078\u0074R\u0065\u0074\u0061\u0069\u006e\u0065\u0064\u0020\u0025\u0076\u000a", _fcef._bebbf))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0069\u0073\u0043\u006f\u0064\u0069\u006e\u0067C\u006f\u006e\u0074\u0065\u0078\u0074\u0055\u0073\u0065\u0064 \u0025\u0076\u000a", _fcef._bcce))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0048u\u0066\u0066\u0041\u0067\u0067\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065S\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0025\u0076\u000a", _fcef.SdHuffAggInstanceSelection))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053d\u0048\u0075\u0066\u0066\u0042\u004d\u0053\u0069\u007a\u0065S\u0065l\u0065\u0063\u0074\u0069\u006f\u006e\u0020%\u0076\u000a", _fcef.SdHuffBMSizeSelection))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0048u\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0057\u0069\u0064\u0074\u0068S\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0025\u0076\u000a", _fcef.SdHuffDecodeWidthSelection))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020Sd\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064e\u0048e\u0069g\u0068t\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0020\u0025\u0076\u000a", _fcef.SdHuffDecodeHeightSelection))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020U\u0073\u0065\u0052\u0065f\u0069\u006e\u0065\u006d\u0065\u006e\u0074A\u0067\u0067\u0072\u0065\u0067\u0061\u0074\u0069\u006f\u006e\u0020\u0025\u0076\u000a", _fcef.UseRefinementAggregation))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020is\u0048\u0075\u0066\u0066\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0065\u0064\u0020\u0025\u0076\u000a", _fcef.IsHuffmanEncoded))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020S\u0064\u0041\u0054\u0058\u0020\u0025\u0076\u000a", _fcef.SdATX))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020S\u0064\u0041\u0054\u0059\u0020\u0025\u0076\u000a", _fcef.SdATY))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0072\u0041\u0054\u0058\u0020\u0025\u0076\u000a", _fcef.SdrATX))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0072\u0041\u0054\u0059\u0020\u0025\u0076\u000a", _fcef.SdrATY))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u004ff\u0045\u0078\u0070\u006f\u0072\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0025\u0076\u000a", _fcef.NumberOfExportedSymbols))
	_affff.WriteString(_be.Sprintf("\u0009-\u0020\u004e\u0075\u006db\u0065\u0072\u004f\u0066\u004ee\u0077S\u0079m\u0062\u006f\u006c\u0073\u0020\u0025\u0076\n", _fcef.NumberOfNewSymbols))
	_affff.WriteString(_be.Sprintf("\u0009\u002d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u004ff\u0049\u006d\u0070\u006f\u0072\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0025\u0076\u000a", _fcef._fgfc))
	_affff.WriteString(_be.Sprintf("\u0009\u002d \u006e\u0075\u006d\u0062\u0065\u0072\u004f\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0064\u0053\u0079\u006d\u0062\u006f\u006c\u0073 %\u0076\u000a", _fcef._geba))
	return _affff.String()
}
func (_gfgc *GenericRefinementRegion) updateOverride() error {
	if _gfgc.GrAtX == nil || _gfgc.GrAtY == nil {
		return _ag.New("\u0041\u0054\u0020\u0070\u0069\u0078\u0065\u006c\u0073\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
	}
	if len(_gfgc.GrAtX) != len(_gfgc.GrAtY) {
		return _ag.New("A\u0054\u0020\u0070\u0069xe\u006c \u0069\u006e\u0063\u006f\u006es\u0069\u0073\u0074\u0065\u006e\u0074")
	}
	_gfgc._gbd = make([]bool, len(_gfgc.GrAtX))
	switch _gfgc.TemplateID {
	case 0:
		if _gfgc.GrAtX[0] != -1 && _gfgc.GrAtY[0] != -1 {
			_gfgc._gbd[0] = true
			_gfgc._age = true
		}
		if _gfgc.GrAtX[1] != -1 && _gfgc.GrAtY[1] != -1 {
			_gfgc._gbd[1] = true
			_gfgc._age = true
		}
	case 1:
		_gfgc._age = false
	}
	return nil
}
func (_abc *PageInformationSegment) String() string {
	_edad := &_bb.Builder{}
	_edad.WriteString("\u000a\u005b\u0050\u0041G\u0045\u002d\u0049\u004e\u0046\u004f\u0052\u004d\u0041\u0054I\u004fN\u002d\u0053\u0045\u0047\u004d\u0045\u004eT\u005d\u000a")
	_edad.WriteString(_be.Sprintf("\u0009\u002d \u0042\u004d\u0048e\u0069\u0067\u0068\u0074\u003a\u0020\u0025\u0064\u000a", _abc.PageBMHeight))
	_edad.WriteString(_be.Sprintf("\u0009-\u0020B\u004d\u0057\u0069\u0064\u0074\u0068\u003a\u0020\u0025\u0064\u000a", _abc.PageBMWidth))
	_edad.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0052es\u006f\u006c\u0075\u0074\u0069\u006f\u006e\u0058\u003a\u0020\u0025\u0064\u000a", _abc.ResolutionX))
	_edad.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0052es\u006f\u006c\u0075\u0074\u0069\u006f\u006e\u0059\u003a\u0020\u0025\u0064\u000a", _abc.ResolutionY))
	_edad.WriteString(_be.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065r\u0061\u0074\u006fr\u003a \u0025\u0073\u000a", _abc._cba))
	_edad.WriteString(_be.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006eO\u0070\u0065\u0072\u0061\u0074\u006f\u0072O\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u003a\u0020\u0025v\u000a", _abc._bddd))
	_edad.WriteString(_be.Sprintf("\u0009-\u0020I\u0073\u004c\u006f\u0073\u0073l\u0065\u0073s\u003a\u0020\u0025\u0076\u000a", _abc.IsLossless))
	_edad.WriteString(_be.Sprintf("\u0009\u002d\u0020R\u0065\u0071\u0075\u0069r\u0065\u0073\u0041\u0075\u0078\u0069\u006ci\u0061\u0072\u0079\u0042\u0075\u0066\u0066\u0065\u0072\u003a\u0020\u0025\u0076\u000a", _abc._cabb))
	_edad.WriteString(_be.Sprintf("\u0009\u002d\u0020M\u0069\u0067\u0068\u0074C\u006f\u006e\u0074\u0061\u0069\u006e\u0052e\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0076\u000a", _abc._adfb))
	_edad.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0049\u0073\u0053\u0074\u0072\u0069\u0070\u0065\u0064:\u0020\u0025\u0076\u000a", _abc.IsStripe))
	_edad.WriteString(_be.Sprintf("\t\u002d\u0020\u004d\u0061xS\u0074r\u0069\u0070\u0065\u0053\u0069z\u0065\u003a\u0020\u0025\u0076\u000a", _abc.MaxStripeSize))
	return _edad.String()
}
func (_aec *PageInformationSegment) readDefaultPixelValue() error {
	_adde, _ebae := _aec._fabba.ReadBit()
	if _ebae != nil {
		return _ebae
	}
	_aec.DefaultPixelValue = uint8(_adde & 0xf)
	return nil
}
func (_gcge *TableSegment) Init(h *Header, r _g.StreamReader) error {
	_gcge._fafc = r
	return _gcge.parseHeader()
}

type SymbolDictionary struct {
	_aege                       _g.StreamReader
	SdrTemplate                 int8
	SdTemplate                  int8
	_bebbf                      bool
	_bcce                       bool
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
	_fgfc                       uint32
	_gdc                        []*_c.Bitmap
	_geba                       uint32
	_gbeb                       []*_c.Bitmap
	_bdgc                       _aag.Tabler
	_dgff                       _aag.Tabler
	_cacb                       _aag.Tabler
	_ece                        _aag.Tabler
	_abab                       []*_c.Bitmap
	_deeb                       []*_c.Bitmap
	_caaa                       *_aa.Decoder
	_cfcd                       *TextRegion
	_eaed                       *GenericRegion
	_abdfc                      *GenericRefinementRegion
	_ebff                       *_aa.DecoderStats
	_fcbe                       *_aa.DecoderStats
	_bba                        *_aa.DecoderStats
	_efab                       *_aa.DecoderStats
	_fddg                       *_aa.DecoderStats
	_bcb                        *_aa.DecoderStats
	_gebb                       *_aa.DecoderStats
	_ddfd                       *_aa.DecoderStats
	_dcfa                       *_aa.DecoderStats
	_ebafc                      int8
	_dfa                        *_c.Bitmaps
	_eaac                       []int
	_faeae                      map[int]int
	_fadf                       bool
}

func (_fdea *SymbolDictionary) encodeFlags(_fed _g.BinaryWriter) (_agfa int, _eceg error) {
	const _agbg = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _eceg = _fed.SkipBits(3); _eceg != nil {
		return 0, _ac.Wrap(_eceg, _agbg, "\u0065\u006d\u0070\u0074\u0079\u0020\u0062\u0069\u0074\u0073")
	}
	var _fcff int
	if _fdea.SdrTemplate > 0 {
		_fcff = 1
	}
	if _eceg = _fed.WriteBit(_fcff); _eceg != nil {
		return _agfa, _ac.Wrap(_eceg, _agbg, "s\u0064\u0072\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_fcff = 0
	if _fdea.SdTemplate > 1 {
		_fcff = 1
	}
	if _eceg = _fed.WriteBit(_fcff); _eceg != nil {
		return _agfa, _ac.Wrap(_eceg, _agbg, "\u0073\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_fcff = 0
	if _fdea.SdTemplate == 1 || _fdea.SdTemplate == 3 {
		_fcff = 1
	}
	if _eceg = _fed.WriteBit(_fcff); _eceg != nil {
		return _agfa, _ac.Wrap(_eceg, _agbg, "\u0073\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_fcff = 0
	if _fdea._bebbf {
		_fcff = 1
	}
	if _eceg = _fed.WriteBit(_fcff); _eceg != nil {
		return _agfa, _ac.Wrap(_eceg, _agbg, "\u0063\u006f\u0064in\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0072\u0065\u0074\u0061\u0069\u006e\u0065\u0064")
	}
	_fcff = 0
	if _fdea._bcce {
		_fcff = 1
	}
	if _eceg = _fed.WriteBit(_fcff); _eceg != nil {
		return _agfa, _ac.Wrap(_eceg, _agbg, "\u0063\u006f\u0064\u0069ng\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0075\u0073\u0065\u0064")
	}
	_fcff = 0
	if _fdea.SdHuffAggInstanceSelection {
		_fcff = 1
	}
	if _eceg = _fed.WriteBit(_fcff); _eceg != nil {
		return _agfa, _ac.Wrap(_eceg, _agbg, "\u0073\u0064\u0048\u0075\u0066\u0066\u0041\u0067\u0067\u0049\u006e\u0073\u0074")
	}
	_fcff = int(_fdea.SdHuffBMSizeSelection)
	if _eceg = _fed.WriteBit(_fcff); _eceg != nil {
		return _agfa, _ac.Wrap(_eceg, _agbg, "\u0073\u0064\u0048u\u0066\u0066\u0042\u006d\u0053\u0069\u007a\u0065")
	}
	_fcff = 0
	if _fdea.SdHuffDecodeWidthSelection > 1 {
		_fcff = 1
	}
	if _eceg = _fed.WriteBit(_fcff); _eceg != nil {
		return _agfa, _ac.Wrap(_eceg, _agbg, "s\u0064\u0048\u0075\u0066\u0066\u0057\u0069\u0064\u0074\u0068")
	}
	_fcff = 0
	switch _fdea.SdHuffDecodeWidthSelection {
	case 1, 3:
		_fcff = 1
	}
	if _eceg = _fed.WriteBit(_fcff); _eceg != nil {
		return _agfa, _ac.Wrap(_eceg, _agbg, "s\u0064\u0048\u0075\u0066\u0066\u0057\u0069\u0064\u0074\u0068")
	}
	_fcff = 0
	if _fdea.SdHuffDecodeHeightSelection > 1 {
		_fcff = 1
	}
	if _eceg = _fed.WriteBit(_fcff); _eceg != nil {
		return _agfa, _ac.Wrap(_eceg, _agbg, "\u0073\u0064\u0048u\u0066\u0066\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_fcff = 0
	switch _fdea.SdHuffDecodeHeightSelection {
	case 1, 3:
		_fcff = 1
	}
	if _eceg = _fed.WriteBit(_fcff); _eceg != nil {
		return _agfa, _ac.Wrap(_eceg, _agbg, "\u0073\u0064\u0048u\u0066\u0066\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_fcff = 0
	if _fdea.UseRefinementAggregation {
		_fcff = 1
	}
	if _eceg = _fed.WriteBit(_fcff); _eceg != nil {
		return _agfa, _ac.Wrap(_eceg, _agbg, "\u0073\u0064\u0052\u0065\u0066\u0041\u0067\u0067")
	}
	_fcff = 0
	if _fdea.IsHuffmanEncoded {
		_fcff = 1
	}
	if _eceg = _fed.WriteBit(_fcff); _eceg != nil {
		return _agfa, _ac.Wrap(_eceg, _agbg, "\u0073\u0064\u0048\u0075\u0066\u0066")
	}
	return 2, nil
}
func (_bfcg *SymbolDictionary) readRefinementAtPixels(_gcbce int) error {
	_bfcg.SdrATX = make([]int8, _gcbce)
	_bfcg.SdrATY = make([]int8, _gcbce)
	var (
		_afbeg byte
		_acdc  error
	)
	for _fffd := 0; _fffd < _gcbce; _fffd++ {
		_afbeg, _acdc = _bfcg._aege.ReadByte()
		if _acdc != nil {
			return _acdc
		}
		_bfcg.SdrATX[_fffd] = int8(_afbeg)
		_afbeg, _acdc = _bfcg._aege.ReadByte()
		if _acdc != nil {
			return _acdc
		}
		_bfcg.SdrATY[_fffd] = int8(_afbeg)
	}
	return nil
}
func (_eaacc *TextRegion) readRegionFlags() error {
	var (
		_aadbcg int
		_fgab   uint64
		_ebage  error
	)
	_aadbcg, _ebage = _eaacc._ceegd.ReadBit()
	if _ebage != nil {
		return _ebage
	}
	_eaacc.SbrTemplate = int8(_aadbcg)
	_fgab, _ebage = _eaacc._ceegd.ReadBits(5)
	if _ebage != nil {
		return _ebage
	}
	_eaacc.SbDsOffset = int8(_fgab)
	if _eaacc.SbDsOffset > 0x0f {
		_eaacc.SbDsOffset -= 0x20
	}
	_aadbcg, _ebage = _eaacc._ceegd.ReadBit()
	if _ebage != nil {
		return _ebage
	}
	_eaacc.DefaultPixel = int8(_aadbcg)
	_fgab, _ebage = _eaacc._ceegd.ReadBits(2)
	if _ebage != nil {
		return _ebage
	}
	_eaacc.CombinationOperator = _c.CombinationOperator(int(_fgab) & 0x3)
	_aadbcg, _ebage = _eaacc._ceegd.ReadBit()
	if _ebage != nil {
		return _ebage
	}
	_eaacc.IsTransposed = int8(_aadbcg)
	_fgab, _ebage = _eaacc._ceegd.ReadBits(2)
	if _ebage != nil {
		return _ebage
	}
	_eaacc.ReferenceCorner = int16(_fgab) & 0x3
	_fgab, _ebage = _eaacc._ceegd.ReadBits(2)
	if _ebage != nil {
		return _ebage
	}
	_eaacc.LogSBStrips = int16(_fgab) & 0x3
	_eaacc.SbStrips = 1 << uint(_eaacc.LogSBStrips)
	_aadbcg, _ebage = _eaacc._ceegd.ReadBit()
	if _ebage != nil {
		return _ebage
	}
	if _aadbcg == 1 {
		_eaacc.UseRefinement = true
	}
	_aadbcg, _ebage = _eaacc._ceegd.ReadBit()
	if _ebage != nil {
		return _ebage
	}
	if _aadbcg == 1 {
		_eaacc.IsHuffmanEncoded = true
	}
	return nil
}
func (_ffb *GenericRefinementRegion) decodeSLTP() (int, error) {
	_ffb.Template.setIndex(_ffb._ef)
	return _ffb._fb.DecodeBit(_ffb._ef)
}
func (_ecfdd *TableSegment) HtLow() int32                                      { return _ecfdd._afgc }
func (_afag *PageInformationSegment) CombinationOperatorOverrideAllowed() bool { return _afag._bddd }

type template0 struct{}

func (_afae *SymbolDictionary) decodeHeightClassCollectiveBitmap(_gdag int64, _bbca, _cbab uint32) (*_c.Bitmap, error) {
	if _gdag == 0 {
		_gfgbd := _c.New(int(_cbab), int(_bbca))
		var (
			_edgb byte
			_baa  error
		)
		for _aebgf := 0; _aebgf < len(_gfgbd.Data); _aebgf++ {
			_edgb, _baa = _afae._aege.ReadByte()
			if _baa != nil {
				return nil, _baa
			}
			if _baa = _gfgbd.SetByte(_aebgf, _edgb); _baa != nil {
				return nil, _baa
			}
		}
		return _gfgbd, nil
	}
	if _afae._eaed == nil {
		_afae._eaed = NewGenericRegion(_afae._aege)
	}
	_afae._eaed.setParameters(true, _afae._aege.StreamPosition(), _gdag, _bbca, _cbab)
	_cgade, _cbfgd := _afae._eaed.GetRegionBitmap()
	if _cbfgd != nil {
		return nil, _cbfgd
	}
	return _cgade, nil
}
func (_cegef *TextRegion) decodeDT() (_gabe int64, _dcaaa error) {
	if _cegef.IsHuffmanEncoded {
		if _cegef.SbHuffDT == 3 {
			_gabe, _dcaaa = _cegef._afeb.Decode(_cegef._ceegd)
			if _dcaaa != nil {
				return 0, _dcaaa
			}
		} else {
			var _caab _aag.Tabler
			_caab, _dcaaa = _aag.GetStandardTable(11 + int(_cegef.SbHuffDT))
			if _dcaaa != nil {
				return 0, _dcaaa
			}
			_gabe, _dcaaa = _caab.Decode(_cegef._ceegd)
			if _dcaaa != nil {
				return 0, _dcaaa
			}
		}
	} else {
		var _dgde int32
		_dgde, _dcaaa = _cegef._bbcea.DecodeInt(_cegef._fdfbf)
		if _dcaaa != nil {
			return
		}
		_gabe = int64(_dgde)
	}
	_gabe *= int64(_cegef.SbStrips)
	return _gabe, nil
}

type Documenter interface {
	GetPage(int) (Pager, error)
	GetGlobalSegment(int) (*Header, error)
}

func (_gbafd *TextRegion) decodeID() (int64, error) {
	if _gbafd.IsHuffmanEncoded {
		if _gbafd._gaead == nil {
			_bbab, _ffegb := _gbafd._ceegd.ReadBits(byte(_gbafd._bfa))
			return int64(_bbab), _ffegb
		}
		return _gbafd._gaead.Decode(_gbafd._ceegd)
	}
	return _gbafd._bbcea.DecodeIAID(uint64(_gbafd._bfa), _gbafd._bcfb)
}

var _ templater = &template0{}

func (_cea *TableSegment) HtRS() int32 { return _cea._fgfe }
func (_dega *HalftoneRegion) computeY(_cgaf, _bdfb int) int {
	return _dega.shiftAndFill(int(_dega.HGridY) + _cgaf*int(_dega.HRegionX) - _bdfb*int(_dega.HRegionY))
}
func _ggad(_gbcf _g.StreamReader, _fccag *Header) *TextRegion {
	_edcd := &TextRegion{_ceegd: _gbcf, Header: _fccag, RegionInfo: NewRegionSegment(_gbcf)}
	return _edcd
}

type EndOfStripe struct {
	_fd _g.StreamReader
	_bf int
}

func (_eb *GenericRefinementRegion) getGrReference() (*_c.Bitmap, error) {
	segments := _eb._ff.RTSegments
	if len(segments) == 0 {
		return nil, _ag.New("\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0065\u0078is\u0074\u0073")
	}
	_df, _ca := segments[0].GetSegmentData()
	if _ca != nil {
		return nil, _ca
	}
	_dfd, _eeab := _df.(Regioner)
	if !_eeab {
		return nil, _be.Errorf("\u0072\u0065\u0066\u0065\u0072r\u0065\u0064\u0020\u0074\u006f\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0052\u0065\u0067\u0069\u006f\u006e\u0065\u0072\u003a\u0020\u0025\u0054", _df)
	}
	return _dfd.GetRegionBitmap()
}
func NewGenericRegion(r _g.StreamReader) *GenericRegion {
	return &GenericRegion{RegionSegment: NewRegionSegment(r), _degf: r}
}
func (_cagd *TextRegion) InitEncode(globalSymbolsMap, localSymbolsMap map[int]int, comps []int, inLL *_c.Points, symbols *_c.Bitmaps, classIDs *_cb.IntSlice, boxes *_c.Boxes, width, height, symBits int) {
	_cagd.RegionInfo = &RegionSegment{BitmapWidth: uint32(width), BitmapHeight: uint32(height)}
	_cagd._afcf = globalSymbolsMap
	_cagd._dbe = localSymbolsMap
	_cagd._babe = comps
	_cagd._edded = inLL
	_cagd._ebde = symbols
	_cagd._dcfb = classIDs
	_cagd._bfab = boxes
	_cagd._fadfc = symBits
}
func (_gbcgg *TextRegion) decodeRdy() (int64, error) {
	const _agdb = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064y"
	if _gbcgg.IsHuffmanEncoded {
		if _gbcgg.SbHuffRDY == 3 {
			if _gbcgg._ecgf == nil {
				var (
					_dggb int
					_aabc error
				)
				if _gbcgg.SbHuffFS == 3 {
					_dggb++
				}
				if _gbcgg.SbHuffDS == 3 {
					_dggb++
				}
				if _gbcgg.SbHuffDT == 3 {
					_dggb++
				}
				if _gbcgg.SbHuffRDWidth == 3 {
					_dggb++
				}
				if _gbcgg.SbHuffRDHeight == 3 {
					_dggb++
				}
				if _gbcgg.SbHuffRDX == 3 {
					_dggb++
				}
				_gbcgg._ecgf, _aabc = _gbcgg.getUserTable(_dggb)
				if _aabc != nil {
					return 0, _ac.Wrap(_aabc, _agdb, "")
				}
			}
			return _gbcgg._ecgf.Decode(_gbcgg._ceegd)
		}
		_cggg, _edgf := _aag.GetStandardTable(14 + int(_gbcgg.SbHuffRDY))
		if _edgf != nil {
			return 0, _edgf
		}
		return _cggg.Decode(_gbcgg._ceegd)
	}
	_cbg, _cbfgg := _gbcgg._bbcea.DecodeInt(_gbcgg._badb)
	if _cbfgg != nil {
		return 0, _ac.Wrap(_cbfgg, _agdb, "")
	}
	return int64(_cbg), nil
}
func (_dd *GenericRefinementRegion) GetRegionBitmap() (*_c.Bitmap, error) {
	var _ee error
	_gb.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0047\u0065\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		if _ee != nil {
			_gb.Log.Trace("[\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052E\u0046\u002d\u0052\u0045\u0047\u0049\u004fN]\u0020\u0047\u0065\u0074R\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006dap\u0020\u0066a\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _ee)
		} else {
			_gb.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0047\u0065\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		}
	}()
	if _dd.RegionBitmap != nil {
		return _dd.RegionBitmap, nil
	}
	_fag := 0
	if _dd.ReferenceBitmap == nil {
		_dd.ReferenceBitmap, _ee = _dd.getGrReference()
		if _ee != nil {
			return nil, _ee
		}
	}
	if _dd._fb == nil {
		_dd._fb, _ee = _aa.New(_dd._ea)
		if _ee != nil {
			return nil, _ee
		}
	}
	if _dd._ef == nil {
		_dd._ef = _aa.NewStats(8192, 1)
	}
	_dd.RegionBitmap = _c.New(int(_dd.RegionInfo.BitmapWidth), int(_dd.RegionInfo.BitmapHeight))
	if _dd.TemplateID == 0 {
		if _ee = _dd.updateOverride(); _ee != nil {
			return nil, _ee
		}
	}
	_eea := (_dd.RegionBitmap.Width + 7) & -8
	var _bd int
	if _dd.IsTPGROn {
		_bd = int(-_dd.ReferenceDY) * _dd.ReferenceBitmap.RowStride
	}
	_ec := _bd + 1
	for _dg := 0; _dg < _dd.RegionBitmap.Height; _dg++ {
		if _dd.IsTPGROn {
			_eff, _cf := _dd.decodeSLTP()
			if _cf != nil {
				return nil, _cf
			}
			_fag ^= _eff
		}
		if _fag == 0 {
			_ee = _dd.decodeOptimized(_dg, _dd.RegionBitmap.Width, _dd.RegionBitmap.RowStride, _dd.ReferenceBitmap.RowStride, _eea, _bd, _ec)
			if _ee != nil {
				return nil, _ee
			}
		} else {
			_ee = _dd.decodeTypicalPredictedLine(_dg, _dd.RegionBitmap.Width, _dd.RegionBitmap.RowStride, _dd.ReferenceBitmap.RowStride, _eea, _bd)
			if _ee != nil {
				return nil, _ee
			}
		}
	}
	return _dd.RegionBitmap, nil
}
func (_dcaf *GenericRegion) computeSegmentDataStructure() error {
	_dcaf.DataOffset = _dcaf._degf.StreamPosition()
	_dcaf.DataHeaderLength = _dcaf.DataOffset - _dcaf.DataHeaderOffset
	_dcaf.DataLength = int64(_dcaf._degf.Length()) - _dcaf.DataHeaderLength
	return nil
}
func (_eggg *TextRegion) decodeRI() (int64, error) {
	if !_eggg.UseRefinement {
		return 0, nil
	}
	if _eggg.IsHuffmanEncoded {
		_gfgbg, _feba := _eggg._ceegd.ReadBit()
		return int64(_gfgbg), _feba
	}
	_gebbd, _fcdc := _eggg._bbcea.DecodeInt(_eggg._feaceb)
	return int64(_gebbd), _fcdc
}
func (_ggc *HalftoneRegion) GetRegionInfo() *RegionSegment { return _ggc.RegionSegment }
func (_egcd *Header) GetSegmentData() (Segmenter, error) {
	var _fgfd Segmenter
	if _egcd.SegmentData != nil {
		_fgfd = _egcd.SegmentData
	}
	if _fgfd == nil {
		_fcfg, _cgee := _geb[_egcd.Type]
		if !_cgee {
			return nil, _be.Errorf("\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0073\u002f\u0020\u0025\u0064\u0020\u0063\u0072e\u0061t\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e\u0020", _egcd.Type, _egcd.Type)
		}
		_fgfd = _fcfg()
		_gb.Log.Trace("\u005b\u0053E\u0047\u004d\u0045\u004e\u0054-\u0048\u0045\u0041\u0044\u0045R\u005d\u005b\u0023\u0025\u0064\u005d\u0020\u0047\u0065\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u0074\u0020\u004f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0030\u0034\u0058", _egcd.SegmentNumber, _egcd.SegmentDataStartOffset)
		_aaeg, _aeg := _egcd.subInputReader()
		if _aeg != nil {
			return nil, _aeg
		}
		if _daba := _fgfd.Init(_egcd, _aaeg); _daba != nil {
			_gb.Log.Debug("\u0049\u006e\u0069\u0074 \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076 \u0066o\u0072\u0020\u0074\u0079\u0070\u0065\u003a \u0025\u0054", _daba, _fgfd)
			return nil, _daba
		}
		_egcd.SegmentData = _fgfd
	}
	return _fgfd, nil
}
func (_fccg *Header) subInputReader() (_g.StreamReader, error) {
	return _g.NewSubstreamReader(_fccg.Reader, _fccg.SegmentDataStartOffset, _fccg.SegmentDataLength)
}
func (_bede *TextRegion) decodeSymbolInstances() error {
	_gfcef, _fgcf := _bede.decodeStripT()
	if _fgcf != nil {
		return _fgcf
	}
	var (
		_fgca int64
		_cega uint32
	)
	for _cega < _bede.NumberOfSymbolInstances {
		_ffbc, _bdcea := _bede.decodeDT()
		if _bdcea != nil {
			return _bdcea
		}
		_gfcef += _ffbc
		var _cgdd int64
		_cgfa := true
		_bede._dgee = 0
		for {
			if _cgfa {
				_cgdd, _bdcea = _bede.decodeDfs()
				if _bdcea != nil {
					return _bdcea
				}
				_fgca += _cgdd
				_bede._dgee = _fgca
				_cgfa = false
			} else {
				_ggag, _begge := _bede.decodeIds()
				if _fc.Is(_begge, _bca.ErrOOB) {
					break
				}
				if _begge != nil {
					return _begge
				}
				if _cega >= _bede.NumberOfSymbolInstances {
					break
				}
				_bede._dgee += _ggag + int64(_bede.SbDsOffset)
			}
			_gdfa, _bbcg := _bede.decodeCurrentT()
			if _bbcg != nil {
				return _bbcg
			}
			_deee := _gfcef + _gdfa
			_ddfe, _bbcg := _bede.decodeID()
			if _bbcg != nil {
				return _bbcg
			}
			_abbge, _bbcg := _bede.decodeRI()
			if _bbcg != nil {
				return _bbcg
			}
			_bbcc, _bbcg := _bede.decodeIb(_abbge, _ddfe)
			if _bbcg != nil {
				return _bbcg
			}
			if _bbcg = _bede.blit(_bbcc, _deee); _bbcg != nil {
				return _bbcg
			}
			_cega++
		}
	}
	return nil
}
func (_fgb *HalftoneRegion) checkInput() error {
	if _fgb.IsMMREncoded {
		if _fgb.HTemplate != 0 {
			_gb.Log.Debug("\u0048\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0030", _fgb.HTemplate)
		}
		if _fgb.HSkipEnabled {
			_gb.Log.Debug("\u0048\u0053\u006b\u0069\u0070\u0045\u006e\u0061\u0062\u006c\u0065\u0064\u0020\u0030\u0020\u0025\u0076\u0020(\u0073\u0068\u006f\u0075\u006c\u0064\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u0029", _fgb.HSkipEnabled)
		}
	}
	return nil
}
func (_bbfc *SymbolDictionary) setAtPixels() error {
	if _bbfc.IsHuffmanEncoded {
		return nil
	}
	_aefe := 1
	if _bbfc.SdTemplate == 0 {
		_aefe = 4
	}
	if _fgbc := _bbfc.readAtPixels(_aefe); _fgbc != nil {
		return _fgbc
	}
	return nil
}

var _ SegmentEncoder = &GenericRegion{}

func (_dede *PatternDictionary) readTemplate() error {
	_ddcbe, _gbdf := _dede._fad.ReadBits(2)
	if _gbdf != nil {
		return _gbdf
	}
	_dede.HDTemplate = byte(_ddcbe)
	return nil
}
func (_bbcb *GenericRefinementRegion) decodeTypicalPredictedLineTemplate1(_cc, _cab, _aabf, _gee, _fbbd, _bfe, _fde, _da, _bga int) (_bee error) {
	var (
		_ce, _cda  int
		_bcee, _fe int
		_agb, _cce int
		_fba       byte
	)
	if _cc > 0 {
		_fba, _bee = _bbcb.RegionBitmap.GetByte(_fde - _aabf)
		if _bee != nil {
			return
		}
		_bcee = int(_fba)
	}
	if _da > 0 && _da <= _bbcb.ReferenceBitmap.Height {
		_fba, _bee = _bbcb.ReferenceBitmap.GetByte(_bga - _gee + _bfe)
		if _bee != nil {
			return
		}
		_fe = int(_fba) << 2
	}
	if _da >= 0 && _da < _bbcb.ReferenceBitmap.Height {
		_fba, _bee = _bbcb.ReferenceBitmap.GetByte(_bga + _bfe)
		if _bee != nil {
			return
		}
		_agb = int(_fba)
	}
	if _da > -2 && _da < _bbcb.ReferenceBitmap.Height-1 {
		_fba, _bee = _bbcb.ReferenceBitmap.GetByte(_bga + _gee + _bfe)
		if _bee != nil {
			return
		}
		_cce = int(_fba)
	}
	_ce = ((_bcee >> 5) & 0x6) | ((_cce >> 2) & 0x30) | (_agb & 0xc0) | (_fe & 0x200)
	_cda = ((_cce >> 2) & 0x70) | (_agb & 0xc0) | (_fe & 0x700)
	var _dba int
	for _acd := 0; _acd < _fbbd; _acd = _dba {
		var (
			_daf int
			_gfg int
		)
		_dba = _acd + 8
		if _daf = _cab - _acd; _daf > 8 {
			_daf = 8
		}
		_eaff := _dba < _cab
		_efa := _dba < _bbcb.ReferenceBitmap.Width
		_bea := _bfe + 1
		if _cc > 0 {
			_fba = 0
			if _eaff {
				_fba, _bee = _bbcb.RegionBitmap.GetByte(_fde - _aabf + 1)
				if _bee != nil {
					return
				}
			}
			_bcee = (_bcee << 8) | int(_fba)
		}
		if _da > 0 && _da <= _bbcb.ReferenceBitmap.Height {
			var _gba int
			if _efa {
				_fba, _bee = _bbcb.ReferenceBitmap.GetByte(_bga - _gee + _bea)
				if _bee != nil {
					return
				}
				_gba = int(_fba) << 2
			}
			_fe = (_fe << 8) | _gba
		}
		if _da >= 0 && _da < _bbcb.ReferenceBitmap.Height {
			_fba = 0
			if _efa {
				_fba, _bee = _bbcb.ReferenceBitmap.GetByte(_bga + _bea)
				if _bee != nil {
					return
				}
			}
			_agb = (_agb << 8) | int(_fba)
		}
		if _da > -2 && _da < (_bbcb.ReferenceBitmap.Height-1) {
			_fba = 0
			if _efa {
				_fba, _bee = _bbcb.ReferenceBitmap.GetByte(_bga + _gee + _bea)
				if _bee != nil {
					return
				}
			}
			_cce = (_cce << 8) | int(_fba)
		}
		for _gbae := 0; _gbae < _daf; _gbae++ {
			var _fdc int
			_agf := (_cda >> 4) & 0x1ff
			switch _agf {
			case 0x1ff:
				_fdc = 1
			case 0x00:
				_fdc = 0
			default:
				_bbcb._ef.SetIndex(int32(_ce))
				_fdc, _bee = _bbcb._fb.DecodeBit(_bbcb._ef)
				if _bee != nil {
					return
				}
			}
			_dab := uint(7 - _gbae)
			_gfg |= _fdc << _dab
			_ce = ((_ce & 0x0d6) << 1) | _fdc | (_bcee>>_dab+5)&0x002 | ((_cce>>_dab + 2) & 0x010) | ((_agb >> _dab) & 0x040) | ((_fe >> _dab) & 0x200)
			_cda = ((_cda & 0xdb) << 1) | ((_cce>>_dab + 2) & 0x010) | ((_agb >> _dab) & 0x080) | ((_fe >> _dab) & 0x400)
		}
		_bee = _bbcb.RegionBitmap.SetByte(_fde, byte(_gfg))
		if _bee != nil {
			return
		}
		_fde++
		_bga++
	}
	return nil
}
func (_agbd *HalftoneRegion) computeGrayScalePlanes(_gacd []*_c.Bitmap, _ddb int) ([][]int, error) {
	_bbed := make([][]int, _agbd.HGridHeight)
	for _bbda := 0; _bbda < len(_bbed); _bbda++ {
		_bbed[_bbda] = make([]int, _agbd.HGridWidth)
	}
	for _ebbf := 0; _ebbf < int(_agbd.HGridHeight); _ebbf++ {
		for _ceg := 0; _ceg < int(_agbd.HGridWidth); _ceg += 8 {
			var _abeg int
			if _agec := int(_agbd.HGridWidth) - _ceg; _agec > 8 {
				_abeg = 8
			} else {
				_abeg = _agec
			}
			_ffff := _gacd[0].GetByteIndex(_ceg, _ebbf)
			for _dcb := 0; _dcb < _abeg; _dcb++ {
				_fee := _dcb + _ceg
				_bbed[_ebbf][_fee] = 0
				for _dcee := 0; _dcee < _ddb; _dcee++ {
					_eage, _egac := _gacd[_dcee].GetByte(_ffff)
					if _egac != nil {
						return nil, _egac
					}
					_edda := _eage >> uint(7-_fee&7)
					_fceg := _edda & 1
					_gbbc := 1 << uint(_dcee)
					_bbb := int(_fceg) * _gbbc
					_bbed[_ebbf][_fee] += _bbb
				}
			}
		}
	}
	return _bbed, nil
}
func (_egbe *SymbolDictionary) getSymbol(_eggc int) (*_c.Bitmap, error) {
	const _bdga = "\u0067e\u0074\u0053\u0079\u006d\u0062\u006fl"
	_bbac, _ecda := _egbe._dfa.GetBitmap(_egbe._eaac[_eggc])
	if _ecda != nil {
		return nil, _ac.Wrap(_ecda, _bdga, "\u0063\u0061n\u0027\u0074\u0020g\u0065\u0074\u0020\u0073\u0079\u006d\u0062\u006f\u006c")
	}
	return _bbac, nil
}
func (_adfc *GenericRegion) decodeTemplate0b(_fagg, _ead, _effd int, _bfdd, _cfdg int) (_fca error) {
	const _gefc = "\u0064\u0065c\u006f\u0064\u0065T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0030\u0062"
	var (
		_eaa, _fbab int
		_dae, _adc  int
		_dcdd       byte
		_gfa        int
	)
	if _fagg >= 1 {
		_dcdd, _fca = _adfc.Bitmap.GetByte(_cfdg)
		if _fca != nil {
			return _ac.Wrap(_fca, _gefc, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_dae = int(_dcdd)
	}
	if _fagg >= 2 {
		_dcdd, _fca = _adfc.Bitmap.GetByte(_cfdg - _adfc.Bitmap.RowStride)
		if _fca != nil {
			return _ac.Wrap(_fca, _gefc, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_adc = int(_dcdd) << 6
	}
	_eaa = (_dae & 0xf0) | (_adc & 0x3800)
	for _agcc := 0; _agcc < _effd; _agcc = _gfa {
		var (
			_eba  byte
			_fdbc int
		)
		_gfa = _agcc + 8
		if _gdg := _ead - _agcc; _gdg > 8 {
			_fdbc = 8
		} else {
			_fdbc = _gdg
		}
		if _fagg > 0 {
			_dae <<= 8
			if _gfa < _ead {
				_dcdd, _fca = _adfc.Bitmap.GetByte(_cfdg + 1)
				if _fca != nil {
					return _ac.Wrap(_fca, _gefc, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_dae |= int(_dcdd)
			}
		}
		if _fagg > 1 {
			_adc <<= 8
			if _gfa < _ead {
				_dcdd, _fca = _adfc.Bitmap.GetByte(_cfdg - _adfc.Bitmap.RowStride + 1)
				if _fca != nil {
					return _ac.Wrap(_fca, _gefc, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_adc |= int(_dcdd) << 6
			}
		}
		for _afg := 0; _afg < _fdbc; _afg++ {
			_bcg := uint(7 - _afg)
			if _adfc._aaaf {
				_fbab = _adfc.overrideAtTemplate0b(_eaa, _agcc+_afg, _fagg, int(_eba), _afg, int(_bcg))
				_adfc._bcae.SetIndex(int32(_fbab))
			} else {
				_adfc._bcae.SetIndex(int32(_eaa))
			}
			var _eag int
			_eag, _fca = _adfc._ga.DecodeBit(_adfc._bcae)
			if _fca != nil {
				return _ac.Wrap(_fca, _gefc, "")
			}
			_eba |= byte(_eag << _bcg)
			_eaa = ((_eaa & 0x7bf7) << 1) | _eag | ((_dae >> _bcg) & 0x10) | ((_adc >> _bcg) & 0x800)
		}
		if _gcgd := _adfc.Bitmap.SetByte(_bfdd, _eba); _gcgd != nil {
			return _ac.Wrap(_gcgd, _gefc, "")
		}
		_bfdd++
		_cfdg++
	}
	return nil
}
func (_egefa *GenericRefinementRegion) String() string {
	_fcc := &_bb.Builder{}
	_fcc.WriteString("\u000a[\u0047E\u004e\u0045\u0052\u0049\u0043 \u0052\u0045G\u0049\u004f\u004e\u005d\u000a")
	_fcc.WriteString(_egefa.RegionInfo.String() + "\u000a")
	_fcc.WriteString(_be.Sprintf("\u0009\u002d \u0049\u0073\u0054P\u0047\u0052\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _egefa.IsTPGROn))
	_fcc.WriteString(_be.Sprintf("\u0009-\u0020T\u0065\u006d\u0070\u006c\u0061t\u0065\u0049D\u003a\u0020\u0025\u0076\u000a", _egefa.TemplateID))
	_fcc.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0047\u0072\u0041\u0074\u0058\u003a\u0020\u0025\u0076\u000a", _egefa.GrAtX))
	_fcc.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0047\u0072\u0041\u0074\u0059\u003a\u0020\u0025\u0076\u000a", _egefa.GrAtY))
	_fcc.WriteString(_be.Sprintf("\u0009-\u0020R\u0065\u0066\u0065\u0072\u0065n\u0063\u0065D\u0058\u0020\u0025\u0076\u000a", _egefa.ReferenceDX))
	_fcc.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0052ef\u0065\u0072\u0065\u006e\u0063\u0044\u0065\u0059\u003a\u0020\u0025\u0076\u000a", _egefa.ReferenceDY))
	return _fcc.String()
}
func (_bbeg *PageInformationSegment) encodeFlags(_afgg _g.BinaryWriter) (_caac error) {
	const _ecdb = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _caac = _afgg.SkipBits(1); _caac != nil {
		return _ac.Wrap(_caac, _ecdb, "\u0072\u0065\u0073e\u0072\u0076\u0065\u0064\u0020\u0062\u0069\u0074")
	}
	var _eeg int
	if _bbeg.CombinationOperatorOverrideAllowed() {
		_eeg = 1
	}
	if _caac = _afgg.WriteBit(_eeg); _caac != nil {
		return _ac.Wrap(_caac, _ecdb, "\u0063\u006f\u006db\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0070\u0065\u0072a\u0074\u006f\u0072\u0020\u006f\u0076\u0065\u0072\u0072\u0069\u0064\u0064\u0065\u006e")
	}
	_eeg = 0
	if _bbeg._cabb {
		_eeg = 1
	}
	if _caac = _afgg.WriteBit(_eeg); _caac != nil {
		return _ac.Wrap(_caac, _ecdb, "\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0061\u0075\u0078i\u006ci\u0061r\u0079\u0020\u0062\u0075\u0066\u0066\u0065r")
	}
	if _caac = _afgg.WriteBit((int(_bbeg._cba) >> 1) & 0x01); _caac != nil {
		return _ac.Wrap(_caac, _ecdb, "\u0063\u006f\u006d\u0062\u0069\u006e\u0061\u0074\u0069\u006fn\u0020\u006f\u0070\u0065\u0072\u0061\u0074o\u0072\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0062\u0069\u0074")
	}
	if _caac = _afgg.WriteBit(int(_bbeg._cba) & 0x01); _caac != nil {
		return _ac.Wrap(_caac, _ecdb, "\u0063\u006f\u006db\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0070\u0065\u0072a\u0074\u006f\u0072\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0062\u0069\u0074")
	}
	_eeg = int(_bbeg.DefaultPixelValue)
	if _caac = _afgg.WriteBit(_eeg); _caac != nil {
		return _ac.Wrap(_caac, _ecdb, "\u0064e\u0066\u0061\u0075\u006c\u0074\u0020\u0070\u0061\u0067\u0065\u0020p\u0069\u0078\u0065\u006c\u0020\u0076\u0061\u006c\u0075\u0065")
	}
	_eeg = 0
	if _bbeg._adfb {
		_eeg = 1
	}
	if _caac = _afgg.WriteBit(_eeg); _caac != nil {
		return _ac.Wrap(_caac, _ecdb, "\u0063\u006f\u006e\u0074ai\u006e\u0073\u0020\u0072\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074")
	}
	_eeg = 0
	if _bbeg.IsLossless {
		_eeg = 1
	}
	if _caac = _afgg.WriteBit(_eeg); _caac != nil {
		return _ac.Wrap(_caac, _ecdb, "p\u0061\u0067\u0065\u0020\u0069\u0073 \u0065\u0076\u0065\u006e\u0074\u0075\u0061\u006c\u006cy\u0020\u006c\u006fs\u0073l\u0065\u0073\u0073")
	}
	return nil
}
func (_ggff *PageInformationSegment) Size() int { return 19 }
func (_bbdbb *TextRegion) parseHeader() error {
	var _bfgb error
	_gb.Log.Trace("\u005b\u0054E\u0058\u0054\u0020\u0052E\u0047\u0049O\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045-\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0062\u0065\u0067\u0069n\u0073\u002e\u002e\u002e")
	defer func() {
		if _bfgb != nil {
			_gb.Log.Trace("\u005b\u0054\u0045\u0058\u0054\u0020\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044E\u0052\u005d\u0020\u0066\u0061i\u006c\u0065d\u002e\u0020\u0025\u0076", _bfgb)
		} else {
			_gb.Log.Trace("\u005b\u0054E\u0058\u0054\u0020\u0052E\u0047\u0049O\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045-\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0066\u0069\u006e\u0069s\u0068\u0065\u0064\u002e")
		}
	}()
	if _bfgb = _bbdbb.RegionInfo.parseHeader(); _bfgb != nil {
		return _bfgb
	}
	if _bfgb = _bbdbb.readRegionFlags(); _bfgb != nil {
		return _bfgb
	}
	if _bbdbb.IsHuffmanEncoded {
		if _bfgb = _bbdbb.readHuffmanFlags(); _bfgb != nil {
			return _bfgb
		}
	}
	if _bfgb = _bbdbb.readUseRefinement(); _bfgb != nil {
		return _bfgb
	}
	if _bfgb = _bbdbb.readAmountOfSymbolInstances(); _bfgb != nil {
		return _bfgb
	}
	if _bfgb = _bbdbb.getSymbols(); _bfgb != nil {
		return _bfgb
	}
	if _bfgb = _bbdbb.computeSymbolCodeLength(); _bfgb != nil {
		return _bfgb
	}
	if _bfgb = _bbdbb.checkInput(); _bfgb != nil {
		return _bfgb
	}
	_gb.Log.Trace("\u0025\u0073", _bbdbb.String())
	return nil
}
func (_bag *GenericRegion) overrideAtTemplate1(_cabg, _fab, _gfec, _aga, _bbgb int) int {
	_cabg &= 0x1FF7
	if _bag.GBAtY[0] == 0 && _bag.GBAtX[0] >= -int8(_bbgb) {
		_cabg |= (_aga >> uint(7-(int8(_bbgb)+_bag.GBAtX[0])) & 0x1) << 3
	} else {
		_cabg |= int(_bag.getPixel(_fab+int(_bag.GBAtX[0]), _gfec+int(_bag.GBAtY[0]))) << 3
	}
	return _cabg
}
func (_dbdb *HalftoneRegion) GetPatterns() ([]*_c.Bitmap, error) {
	var (
		_agea []*_c.Bitmap
		_afga error
	)
	for _, _cgfg := range _dbdb._bdeg.RTSegments {
		var _gffc Segmenter
		_gffc, _afga = _cgfg.GetSegmentData()
		if _afga != nil {
			_gb.Log.Debug("\u0047e\u0074\u0053\u0065\u0067m\u0065\u006e\u0074\u0044\u0061t\u0061 \u0066a\u0069\u006c\u0065\u0064\u003a\u0020\u0025v", _afga)
			return nil, _afga
		}
		_dbdf, _ebfe := _gffc.(*PatternDictionary)
		if !_ebfe {
			_afga = _be.Errorf("\u0072e\u006c\u0061t\u0065\u0064\u0020\u0073e\u0067\u006d\u0065n\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0070at\u0074\u0065\u0072n\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u003a \u0025\u0054", _gffc)
			return nil, _afga
		}
		var _fgf []*_c.Bitmap
		_fgf, _afga = _dbdf.GetDictionary()
		if _afga != nil {
			_gb.Log.Debug("\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _afga)
			return nil, _afga
		}
		_agea = append(_agea, _fgf...)
	}
	return _agea, nil
}
func (_bdcd *GenericRegion) overrideAtTemplate2(_aagb, _gce, _bcc, _fbde, _cgab int) int {
	_aagb &= 0x3FB
	if _bdcd.GBAtY[0] == 0 && _bdcd.GBAtX[0] >= -int8(_cgab) {
		_aagb |= (_fbde >> uint(7-(int8(_cgab)+_bdcd.GBAtX[0])) & 0x1) << 2
	} else {
		_aagb |= int(_bdcd.getPixel(_gce+int(_bdcd.GBAtX[0]), _bcc+int(_bdcd.GBAtY[0]))) << 2
	}
	return _aagb
}
func (_edg *PageInformationSegment) readRequiresAuxiliaryBuffer() error {
	_dffa, _bcgd := _edg._fabba.ReadBit()
	if _bcgd != nil {
		return _bcgd
	}
	if _dffa == 1 {
		_edg._cabb = true
	}
	return nil
}
func (_ceea *Header) readSegmentDataLength(_ggfg _g.StreamReader) (_fdf error) {
	_ceea.SegmentDataLength, _fdf = _ggfg.ReadBits(32)
	if _fdf != nil {
		return _fdf
	}
	_ceea.SegmentDataLength &= _f.MaxInt32
	return nil
}
func (_cdeg *SymbolDictionary) readRegionFlags() error {
	var (
		_abea uint64
		_cgfe int
	)
	_, _aedg := _cdeg._aege.ReadBits(3)
	if _aedg != nil {
		return _aedg
	}
	_cgfe, _aedg = _cdeg._aege.ReadBit()
	if _aedg != nil {
		return _aedg
	}
	_cdeg.SdrTemplate = int8(_cgfe)
	_abea, _aedg = _cdeg._aege.ReadBits(2)
	if _aedg != nil {
		return _aedg
	}
	_cdeg.SdTemplate = int8(_abea & 0xf)
	_cgfe, _aedg = _cdeg._aege.ReadBit()
	if _aedg != nil {
		return _aedg
	}
	if _cgfe == 1 {
		_cdeg._bebbf = true
	}
	_cgfe, _aedg = _cdeg._aege.ReadBit()
	if _aedg != nil {
		return _aedg
	}
	if _cgfe == 1 {
		_cdeg._bcce = true
	}
	_cgfe, _aedg = _cdeg._aege.ReadBit()
	if _aedg != nil {
		return _aedg
	}
	if _cgfe == 1 {
		_cdeg.SdHuffAggInstanceSelection = true
	}
	_cgfe, _aedg = _cdeg._aege.ReadBit()
	if _aedg != nil {
		return _aedg
	}
	_cdeg.SdHuffBMSizeSelection = int8(_cgfe)
	_abea, _aedg = _cdeg._aege.ReadBits(2)
	if _aedg != nil {
		return _aedg
	}
	_cdeg.SdHuffDecodeWidthSelection = int8(_abea & 0xf)
	_abea, _aedg = _cdeg._aege.ReadBits(2)
	if _aedg != nil {
		return _aedg
	}
	_cdeg.SdHuffDecodeHeightSelection = int8(_abea & 0xf)
	_cgfe, _aedg = _cdeg._aege.ReadBit()
	if _aedg != nil {
		return _aedg
	}
	if _cgfe == 1 {
		_cdeg.UseRefinementAggregation = true
	}
	_cgfe, _aedg = _cdeg._aege.ReadBit()
	if _aedg != nil {
		return _aedg
	}
	if _cgfe == 1 {
		_cdeg.IsHuffmanEncoded = true
	}
	return nil
}
func (_debef *TableSegment) HtOOB() int32 { return _debef._dgecf }
func (_dadag *SymbolDictionary) setExportedSymbols(_abba []int) {
	for _dcgcb := uint32(0); _dcgcb < _dadag._fgfc+_dadag.NumberOfNewSymbols; _dcgcb++ {
		if _abba[_dcgcb] == 1 {
			var _ddce *_c.Bitmap
			if _dcgcb < _dadag._fgfc {
				_ddce = _dadag._gdc[_dcgcb]
			} else {
				_ddce = _dadag._gbeb[_dcgcb-_dadag._fgfc]
			}
			_gb.Log.Trace("\u005bS\u0059\u004dB\u004f\u004c\u002d\u0044I\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u0020\u0041\u0064\u0064 E\u0078\u0070\u006fr\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c:\u0020\u0027%\u0073\u0027", _ddce)
			_dadag._abab = append(_dadag._abab, _ddce)
		}
	}
}
func (_afgf *GenericRegion) getPixel(_gdb, _gaed int) int8 {
	if _gdb < 0 || _gdb >= _afgf.Bitmap.Width {
		return 0
	}
	if _gaed < 0 || _gaed >= _afgf.Bitmap.Height {
		return 0
	}
	if _afgf.Bitmap.GetPixel(_gdb, _gaed) {
		return 1
	}
	return 0
}
func (_fcda *SymbolDictionary) encodeNumSyms(_ggce _g.BinaryWriter) (_fgbdc int, _ccea error) {
	const _gbge = "\u0065\u006e\u0063\u006f\u0064\u0065\u004e\u0075\u006d\u0053\u0079\u006d\u0073"
	_agcb := make([]byte, 4)
	_bc.BigEndian.PutUint32(_agcb, _fcda.NumberOfExportedSymbols)
	if _fgbdc, _ccea = _ggce.Write(_agcb); _ccea != nil {
		return _fgbdc, _ac.Wrap(_ccea, _gbge, "\u0065\u0078p\u006f\u0072\u0074e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	_bc.BigEndian.PutUint32(_agcb, _fcda.NumberOfNewSymbols)
	_dfea, _ccea := _ggce.Write(_agcb)
	if _ccea != nil {
		return _fgbdc, _ac.Wrap(_ccea, _gbge, "n\u0065\u0077\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	return _fgbdc + _dfea, nil
}

type RegionSegment struct {
	_ffade             _g.StreamReader
	BitmapWidth        uint32
	BitmapHeight       uint32
	XLocation          uint32
	YLocation          uint32
	CombinaionOperator _c.CombinationOperator
}

func (_agd *Header) writeFlags(_cdbd _g.BinaryWriter) (_bgac error) {
	const _egae = "\u0048\u0065\u0061\u0064\u0065\u0072\u002e\u0077\u0072\u0069\u0074\u0065F\u006c\u0061\u0067\u0073"
	_bge := byte(_agd.Type)
	if _bgac = _cdbd.WriteByte(_bge); _bgac != nil {
		return _ac.Wrap(_bgac, _egae, "\u0077\u0072\u0069ti\u006e\u0067\u0020\u0073\u0065\u0067\u006d\u0065\u006et\u0020t\u0079p\u0065 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	if !_agd.RetainFlag && !_agd.PageAssociationFieldSize {
		return nil
	}
	if _bgac = _cdbd.SkipBits(-8); _bgac != nil {
		return _ac.Wrap(_bgac, _egae, "\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0062\u0061\u0063\u006b\u0020\u0074\u0068\u0065\u0020\u0062\u0069\u0074\u0073\u0020\u0066\u0061il\u0065\u0064")
	}
	var _fdba int
	if _agd.RetainFlag {
		_fdba = 1
	}
	if _bgac = _cdbd.WriteBit(_fdba); _bgac != nil {
		return _ac.Wrap(_bgac, _egae, "\u0072\u0065\u0074\u0061in\u0020\u0072\u0065\u0074\u0061\u0069\u006e\u0020\u0066\u006c\u0061\u0067\u0073")
	}
	_fdba = 0
	if _agd.PageAssociationFieldSize {
		_fdba = 1
	}
	if _bgac = _cdbd.WriteBit(_fdba); _bgac != nil {
		return _ac.Wrap(_bgac, _egae, "p\u0061\u0067\u0065\u0020as\u0073o\u0063\u0069\u0061\u0074\u0069o\u006e\u0020\u0066\u006c\u0061\u0067")
	}
	_cdbd.FinishByte()
	return nil
}
func (_geea *GenericRegion) setOverrideFlag(_dfbc int) {
	_geea.GBAtOverride[_dfbc] = true
	_geea._aaaf = true
}
func (_ecb *template0) setIndex(_adf *_aa.DecoderStats) { _adf.SetIndex(0x100) }
func (_fabf *TextRegion) String() string {
	_daad := &_bb.Builder{}
	_daad.WriteString("\u000a[\u0054E\u0058\u0054\u0020\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u000a")
	_daad.WriteString(_fabf.RegionInfo.String() + "\u000a")
	_daad.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053br\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0025\u0076\u000a", _fabf.SbrTemplate))
	_daad.WriteString(_be.Sprintf("\u0009-\u0020S\u0062\u0044\u0073\u004f\u0066f\u0073\u0065t\u003a\u0020\u0025\u0076\u000a", _fabf.SbDsOffset))
	_daad.WriteString(_be.Sprintf("\t\u002d \u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0050i\u0078\u0065\u006c\u003a %\u0076\u000a", _fabf.DefaultPixel))
	_daad.WriteString(_be.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065r\u0061\u0074\u006fr\u003a \u0025\u0076\u000a", _fabf.CombinationOperator))
	_daad.WriteString(_be.Sprintf("\t\u002d \u0049\u0073\u0054\u0072\u0061\u006e\u0073\u0070o\u0073\u0065\u0064\u003a %\u0076\u000a", _fabf.IsTransposed))
	_daad.WriteString(_be.Sprintf("\u0009\u002d\u0020Re\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0043\u006f\u0072\u006e\u0065\u0072\u003a\u0020\u0025\u0076\u000a", _fabf.ReferenceCorner))
	_daad.WriteString(_be.Sprintf("\t\u002d\u0020\u0055\u0073eR\u0065f\u0069\u006e\u0065\u006d\u0065n\u0074\u003a\u0020\u0025\u0076\u000a", _fabf.UseRefinement))
	_daad.WriteString(_be.Sprintf("\u0009-\u0020\u0049\u0073\u0048\u0075\u0066\u0066\u006d\u0061\u006e\u0045n\u0063\u006f\u0064\u0065\u0064\u003a\u0020\u0025\u0076\u000a", _fabf.IsHuffmanEncoded))
	if _fabf.IsHuffmanEncoded {
		_daad.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053bH\u0075\u0066\u0066\u0052\u0053\u0069\u007a\u0065\u003a\u0020\u0025\u0076\u000a", _fabf.SbHuffRSize))
		_daad.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0059:\u0020\u0025\u0076\u000a", _fabf.SbHuffRDY))
		_daad.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0058:\u0020\u0025\u0076\u000a", _fabf.SbHuffRDX))
		_daad.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053bH\u0075\u0066\u0066\u0052\u0044\u0048\u0065\u0069\u0067\u0068\u0074\u003a\u0020\u0025v\u000a", _fabf.SbHuffRDHeight))
		_daad.WriteString(_be.Sprintf("\t\u002d\u0020\u0053\u0062Hu\u0066f\u0052\u0044\u0057\u0069\u0064t\u0068\u003a\u0020\u0025\u0076\u000a", _fabf.SbHuffRDWidth))
		_daad.WriteString(_be.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0044\u0054\u003a\u0020\u0025\u0076\u000a", _fabf.SbHuffDT))
		_daad.WriteString(_be.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0044\u0053\u003a\u0020\u0025\u0076\u000a", _fabf.SbHuffDS))
		_daad.WriteString(_be.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0046\u0053\u003a\u0020\u0025\u0076\u000a", _fabf.SbHuffFS))
	}
	_daad.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0058:\u0020\u0025\u0076\u000a", _fabf.SbrATX))
	_daad.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0059:\u0020\u0025\u0076\u000a", _fabf.SbrATY))
	_daad.WriteString(_be.Sprintf("\u0009\u002d\u0020N\u0075\u006d\u0062\u0065r\u004f\u0066\u0053\u0079\u006d\u0062\u006fl\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u003a\u0020\u0025\u0076\u000a", _fabf.NumberOfSymbolInstances))
	_daad.WriteString(_be.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0058:\u0020\u0025\u0076\u000a", _fabf.SbrATX))
	return _daad.String()
}
func (_aaec *SymbolDictionary) InitEncode(symbols *_c.Bitmaps, symbolList []int, symbolMap map[int]int, unborderSymbols bool) error {
	const _eaacf = "S\u0079\u006d\u0062\u006f\u006c\u0044i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002eI\u006e\u0069\u0074E\u006ec\u006f\u0064\u0065"
	_aaec.SdATX = []int8{3, -3, 2, -2}
	_aaec.SdATY = []int8{-1, -1, -2, -2}
	_aaec._dfa = symbols
	_aaec._eaac = make([]int, len(symbolList))
	copy(_aaec._eaac, symbolList)
	if len(_aaec._eaac) != _aaec._dfa.Size() {
		return _ac.Error(_eaacf, "s\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0061\u006e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u004ci\u0073\u0074\u0020\u006f\u0066\u0020\u0064\u0069\u0066\u0066er\u0065\u006e\u0074 \u0073i\u007a\u0065")
	}
	_aaec.NumberOfNewSymbols = uint32(symbols.Size())
	_aaec.NumberOfExportedSymbols = uint32(symbols.Size())
	_aaec._faeae = symbolMap
	_aaec._fadf = unborderSymbols
	return nil
}
func (_fcbab *GenericRegion) Encode(w _g.BinaryWriter) (_gae int, _bdc error) {
	const _egg = "G\u0065n\u0065\u0072\u0069\u0063\u0052\u0065\u0067\u0069o\u006e\u002e\u0045\u006eco\u0064\u0065"
	if _fcbab.Bitmap == nil {
		return 0, _ac.Error(_egg, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	_efbb, _bdc := _fcbab.RegionSegment.Encode(w)
	if _bdc != nil {
		return 0, _ac.Wrap(_bdc, _egg, "\u0052\u0065\u0067\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_gae += _efbb
	if _bdc = w.SkipBits(4); _bdc != nil {
		return _gae, _ac.Wrap(_bdc, _egg, "\u0073k\u0069p\u0020\u0072\u0065\u0073\u0065r\u0076\u0065d\u0020\u0062\u0069\u0074\u0073")
	}
	var _bdeb int
	if _fcbab.IsTPGDon {
		_bdeb = 1
	}
	if _bdc = w.WriteBit(_bdeb); _bdc != nil {
		return _gae, _ac.Wrap(_bdc, _egg, "\u0074\u0070\u0067\u0064\u006f\u006e")
	}
	_bdeb = 0
	if _bdc = w.WriteBit(int(_fcbab.GBTemplate>>1) & 0x01); _bdc != nil {
		return _gae, _ac.Wrap(_bdc, _egg, "f\u0069r\u0073\u0074\u0020\u0067\u0062\u0074\u0065\u006dp\u006c\u0061\u0074\u0065 b\u0069\u0074")
	}
	if _bdc = w.WriteBit(int(_fcbab.GBTemplate) & 0x01); _bdc != nil {
		return _gae, _ac.Wrap(_bdc, _egg, "s\u0065\u0063\u006f\u006ed \u0067b\u0074\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u0062\u0069\u0074")
	}
	if _fcbab.UseMMR {
		_bdeb = 1
	}
	if _bdc = w.WriteBit(_bdeb); _bdc != nil {
		return _gae, _ac.Wrap(_bdc, _egg, "u\u0073\u0065\u0020\u004d\u004d\u0052\u0020\u0062\u0069\u0074")
	}
	_gae++
	if _efbb, _bdc = _fcbab.writeGBAtPixels(w); _bdc != nil {
		return _gae, _ac.Wrap(_bdc, _egg, "")
	}
	_gae += _efbb
	_cdg := _cg.New()
	if _bdc = _cdg.EncodeBitmap(_fcbab.Bitmap, _fcbab.IsTPGDon); _bdc != nil {
		return _gae, _ac.Wrap(_bdc, _egg, "")
	}
	_cdg.Final()
	var _geec int64
	if _geec, _bdc = _cdg.WriteTo(w); _bdc != nil {
		return _gae, _ac.Wrap(_bdc, _egg, "")
	}
	_gae += int(_geec)
	return _gae, nil
}
func (_gebaf *SymbolDictionary) decodeNewSymbols(_fcaf, _cabeb uint32, _bcfc *_c.Bitmap, _bacf, _faga int32) error {
	if _gebaf._abdfc == nil {
		_gebaf._abdfc = _agfb(_gebaf._aege, nil)
		if _gebaf._caaa == nil {
			var _dddb error
			_gebaf._caaa, _dddb = _aa.New(_gebaf._aege)
			if _dddb != nil {
				return _dddb
			}
		}
		if _gebaf._ebff == nil {
			_gebaf._ebff = _aa.NewStats(65536, 1)
		}
	}
	_gebaf._abdfc.setParameters(_gebaf._ebff, _gebaf._caaa, _gebaf.SdrTemplate, _fcaf, _cabeb, _bcfc, _bacf, _faga, false, _gebaf.SdrATX, _gebaf.SdrATY)
	return _gebaf.addSymbol(_gebaf._abdfc)
}
func (_ecca *Header) writeReferredToCount(_egfe _g.BinaryWriter) (_feef int, _bddf error) {
	const _gde = "w\u0072i\u0074\u0065\u0052\u0065\u0066\u0065\u0072\u0072e\u0064\u0054\u006f\u0043ou\u006e\u0074"
	_ecca.RTSNumbers = make([]int, len(_ecca.RTSegments))
	for _fabb, _caca := range _ecca.RTSegments {
		_ecca.RTSNumbers[_fabb] = int(_caca.SegmentNumber)
	}
	if len(_ecca.RTSNumbers) <= 4 {
		var _cfcg byte
		if len(_ecca.RetainBits) >= 1 {
			_cfcg = _ecca.RetainBits[0]
		}
		_cfcg |= byte(len(_ecca.RTSNumbers)) << 5
		if _bddf = _egfe.WriteByte(_cfcg); _bddf != nil {
			return 0, _ac.Wrap(_bddf, _gde, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		return 1, nil
	}
	_cegd := uint32(len(_ecca.RTSNumbers))
	_abgb := make([]byte, 4+_cb.Ceil(len(_ecca.RTSNumbers)+1, 8))
	_cegd |= 0x7 << 29
	_bc.BigEndian.PutUint32(_abgb, _cegd)
	copy(_abgb[1:], _ecca.RetainBits)
	_feef, _bddf = _egfe.Write(_abgb)
	if _bddf != nil {
		return 0, _ac.Wrap(_bddf, _gde, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	return _feef, nil
}
func (_bdgba *TextRegion) getSymbols() error {
	if _bdgba.Header.RTSegments != nil {
		return _bdgba.initSymbols()
	}
	return nil
}
func (_ddf *template1) form(_agc, _fcba, _fbcf, _efc, _gcc int16) int16 {
	return ((_agc & 0x02) << 8) | (_fcba << 6) | ((_fbcf & 0x03) << 4) | (_efc << 1) | _gcc
}
func (_bdaf *SymbolDictionary) addSymbol(_eafc Regioner) error {
	_defa, _dada := _eafc.GetRegionBitmap()
	if _dada != nil {
		return _dada
	}
	_bdaf._gbeb[_bdaf._geba] = _defa
	_bdaf._deeb = append(_bdaf._deeb, _defa)
	_gb.Log.Trace("\u005b\u0053YM\u0042\u004f\u004c \u0044\u0049\u0043\u0054ION\u0041RY\u005d\u0020\u0041\u0064\u0064\u0065\u0064 s\u0079\u006d\u0062\u006f\u006c\u003a\u0020%\u0073", _defa)
	return nil
}
func (_babc *PatternDictionary) extractPatterns(_eded *_c.Bitmap) error {
	var _ffbe int
	_gbcgf := make([]*_c.Bitmap, _babc.GrayMax+1)
	for _ffbe <= int(_babc.GrayMax) {
		_gcba := int(_babc.HdpWidth) * _ffbe
		_cdad := _b.Rect(_gcba, 0, _gcba+int(_babc.HdpWidth), int(_babc.HdpHeight))
		_eeeb, _eegb := _c.Extract(_cdad, _eded)
		if _eegb != nil {
			return _eegb
		}
		_gbcgf[_ffbe] = _eeeb
		_ffbe++
	}
	_babc.Patterns = _gbcgf
	return nil
}
func (_gfcdd *GenericRegion) overrideAtTemplate0a(_dbf, _ffga, _adgc, _egf, _ccb, _dfb int) int {
	if _gfcdd.GBAtOverride[0] {
		_dbf &= 0xFFEF
		if _gfcdd.GBAtY[0] == 0 && _gfcdd.GBAtX[0] >= -int8(_ccb) {
			_dbf |= (_egf >> uint(int8(_dfb)-_gfcdd.GBAtX[0]&0x1)) << 4
		} else {
			_dbf |= int(_gfcdd.getPixel(_ffga+int(_gfcdd.GBAtX[0]), _adgc+int(_gfcdd.GBAtY[0]))) << 4
		}
	}
	if _gfcdd.GBAtOverride[1] {
		_dbf &= 0xFBFF
		if _gfcdd.GBAtY[1] == 0 && _gfcdd.GBAtX[1] >= -int8(_ccb) {
			_dbf |= (_egf >> uint(int8(_dfb)-_gfcdd.GBAtX[1]&0x1)) << 10
		} else {
			_dbf |= int(_gfcdd.getPixel(_ffga+int(_gfcdd.GBAtX[1]), _adgc+int(_gfcdd.GBAtY[1]))) << 10
		}
	}
	if _gfcdd.GBAtOverride[2] {
		_dbf &= 0xF7FF
		if _gfcdd.GBAtY[2] == 0 && _gfcdd.GBAtX[2] >= -int8(_ccb) {
			_dbf |= (_egf >> uint(int8(_dfb)-_gfcdd.GBAtX[2]&0x1)) << 11
		} else {
			_dbf |= int(_gfcdd.getPixel(_ffga+int(_gfcdd.GBAtX[2]), _adgc+int(_gfcdd.GBAtY[2]))) << 11
		}
	}
	if _gfcdd.GBAtOverride[3] {
		_dbf &= 0x7FFF
		if _gfcdd.GBAtY[3] == 0 && _gfcdd.GBAtX[3] >= -int8(_ccb) {
			_dbf |= (_egf >> uint(int8(_dfb)-_gfcdd.GBAtX[3]&0x1)) << 15
		} else {
			_dbf |= int(_gfcdd.getPixel(_ffga+int(_gfcdd.GBAtX[3]), _adgc+int(_gfcdd.GBAtY[3]))) << 15
		}
	}
	return _dbf
}
func (_gd *GenericRegion) InitEncode(bm *_c.Bitmap, xLoc, yLoc, template int, duplicateLineRemoval bool) error {
	const _caf = "\u0047e\u006e\u0065\u0072\u0069\u0063\u0052\u0065\u0067\u0069\u006f\u006e.\u0049\u006e\u0069\u0074\u0045\u006e\u0063\u006f\u0064\u0065"
	if bm == nil {
		return _ac.Error(_caf, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if xLoc < 0 || yLoc < 0 {
		return _ac.Error(_caf, "\u0078\u0020\u0061\u006e\u0064\u0020\u0079\u0020\u006c\u006f\u0063\u0061\u0074i\u006f\u006e\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	_gd.Bitmap = bm
	_gd.GBTemplate = byte(template)
	switch _gd.GBTemplate {
	case 0:
		_gd.GBAtX = []int8{3, -3, 2, -2}
		_gd.GBAtY = []int8{-1, -1, -2, -2}
	case 1:
		_gd.GBAtX = []int8{3}
		_gd.GBAtY = []int8{-1}
	case 2, 3:
		_gd.GBAtX = []int8{2}
		_gd.GBAtY = []int8{-1}
	default:
		return _ac.Errorf(_caf, "\u0070\u0072o\u0076\u0069\u0064\u0065\u0064 \u0074\u0065\u006d\u0070\u006ca\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0030\u002c\u0031\u002c\u0032\u002c\u0033\u007d", template)
	}
	_gd.RegionSegment = &RegionSegment{BitmapHeight: uint32(bm.Height), BitmapWidth: uint32(bm.Width), XLocation: uint32(xLoc), YLocation: uint32(yLoc)}
	_gd.IsTPGDon = duplicateLineRemoval
	return nil
}
func (_fafa *PatternDictionary) readGrayMax() error {
	_bgce, _bfba := _fafa._fad.ReadBits(32)
	if _bfba != nil {
		return _bfba
	}
	_fafa.GrayMax = uint32(_bgce & _f.MaxUint32)
	return nil
}
func (_cgeec *SymbolDictionary) getToExportFlags() ([]int, error) {
	var (
		_geda  int
		_decg  int32
		_gcgc  error
		_cebb  = int32(_cgeec._fgfc + _cgeec.NumberOfNewSymbols)
		_eecef = make([]int, _cebb)
	)
	for _ebade := int32(0); _ebade < _cebb; _ebade += _decg {
		if _cgeec.IsHuffmanEncoded {
			_cdfa, _eabe := _aag.GetStandardTable(1)
			if _eabe != nil {
				return nil, _eabe
			}
			_aedfb, _eabe := _cdfa.Decode(_cgeec._aege)
			if _eabe != nil {
				return nil, _eabe
			}
			_decg = int32(_aedfb)
		} else {
			_decg, _gcgc = _cgeec._caaa.DecodeInt(_cgeec._fddg)
			if _gcgc != nil {
				return nil, _gcgc
			}
		}
		if _decg != 0 {
			for _bcfg := _ebade; _bcfg < _ebade+_decg; _bcfg++ {
				_eecef[_bcfg] = _geda
			}
		}
		if _geda == 0 {
			_geda = 1
		} else {
			_geda = 0
		}
	}
	return _eecef, nil
}
func (_gfaf *RegionSegment) String() string {
	_ffc := &_bb.Builder{}
	_ffc.WriteString("\u0009[\u0052E\u0047\u0049\u004f\u004e\u0020S\u0045\u0047M\u0045\u004e\u0054\u005d\u000a")
	_ffc.WriteString(_be.Sprintf("\t\u0009\u002d\u0020\u0042\u0069\u0074m\u0061\u0070\u0020\u0028\u0077\u0069d\u0074\u0068\u002c\u0020\u0068\u0065\u0069g\u0068\u0074\u0029\u0020\u005b\u0025\u0064\u0078\u0025\u0064]\u000a", _gfaf.BitmapWidth, _gfaf.BitmapHeight))
	_ffc.WriteString(_be.Sprintf("\u0009\u0009\u002d\u0020L\u006f\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0028\u0078,\u0079)\u003a\u0020\u005b\u0025\u0064\u002c\u0025d\u005d\u000a", _gfaf.XLocation, _gfaf.YLocation))
	_ffc.WriteString(_be.Sprintf("\t\u0009\u002d\u0020\u0043\u006f\u006db\u0069\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0073", _gfaf.CombinaionOperator))
	return _ffc.String()
}
func (_gacg *PatternDictionary) Init(h *Header, r _g.StreamReader) error {
	_gacg._fad = r
	return _gacg.parseHeader()
}
func (_effdd *GenericRegion) setParameters(_gbbg bool, _cabe, _eecb int64, _bfed, _feb uint32) {
	_effdd.IsMMREncoded = _gbbg
	_effdd.DataOffset = _cabe
	_effdd.DataLength = _eecb
	_effdd.RegionSegment.BitmapHeight = _bfed
	_effdd.RegionSegment.BitmapWidth = _feb
	_effdd._cgdc = nil
	_effdd.Bitmap = nil
}
func (_dgcad *TextRegion) createRegionBitmap() error {
	_dgcad.RegionBitmap = _c.New(int(_dgcad.RegionInfo.BitmapWidth), int(_dgcad.RegionInfo.BitmapHeight))
	if _dgcad.DefaultPixel != 0 {
		_dgcad.RegionBitmap.SetDefaultPixel()
	}
	return nil
}
func (_aaad *GenericRefinementRegion) readAtPixels() error {
	_aaad.GrAtX = make([]int8, 2)
	_aaad.GrAtY = make([]int8, 2)
	_bead, _afa := _aaad._ea.ReadByte()
	if _afa != nil {
		return _afa
	}
	_aaad.GrAtX[0] = int8(_bead)
	_bead, _afa = _aaad._ea.ReadByte()
	if _afa != nil {
		return _afa
	}
	_aaad.GrAtY[0] = int8(_bead)
	_bead, _afa = _aaad._ea.ReadByte()
	if _afa != nil {
		return _afa
	}
	_aaad.GrAtX[1] = int8(_bead)
	_bead, _afa = _aaad._ea.ReadByte()
	if _afa != nil {
		return _afa
	}
	_aaad.GrAtY[1] = int8(_bead)
	return nil
}
func (_agga *Header) writeReferredToSegments(_gdda _g.BinaryWriter) (_ebgc int, _adge error) {
	const _fabc = "\u0077\u0072\u0069te\u0052\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0054\u006f\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073"
	var (
		_cefb uint16
		_cfe  uint32
	)
	_aefd := _agga.referenceSize()
	_cdbge := 1
	_fdfb := make([]byte, _aefd)
	for _, _bgga := range _agga.RTSNumbers {
		switch _aefd {
		case 4:
			_cfe = uint32(_bgga)
			_bc.BigEndian.PutUint32(_fdfb, _cfe)
			_cdbge, _adge = _gdda.Write(_fdfb)
			if _adge != nil {
				return 0, _ac.Wrap(_adge, _fabc, "u\u0069\u006e\u0074\u0033\u0032\u0020\u0073\u0069\u007a\u0065")
			}
		case 2:
			_cefb = uint16(_bgga)
			_bc.BigEndian.PutUint16(_fdfb, _cefb)
			_cdbge, _adge = _gdda.Write(_fdfb)
			if _adge != nil {
				return 0, _ac.Wrap(_adge, _fabc, "\u0075\u0069\u006e\u0074\u0031\u0036")
			}
		default:
			if _adge = _gdda.WriteByte(byte(_bgga)); _adge != nil {
				return 0, _ac.Wrap(_adge, _fabc, "\u0075\u0069\u006et\u0038")
			}
		}
		_ebgc += _cdbge
	}
	return _ebgc, nil
}
func (_cdefc *GenericRegion) decodeTemplate2(_cced, _bacg, _gfgg int, _bbg, _fecg int) (_fgdc error) {
	const _ecba = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0032"
	var (
		_debe, _dgfb int
		_dce, _bbgd  int
		_acbd        byte
		_egea, _ceeg int
	)
	if _cced >= 1 {
		_acbd, _fgdc = _cdefc.Bitmap.GetByte(_fecg)
		if _fgdc != nil {
			return _ac.Wrap(_fgdc, _ecba, "\u006ci\u006ee\u004e\u0075\u006d\u0062\u0065\u0072\u0020\u003e\u003d\u0020\u0031")
		}
		_dce = int(_acbd)
	}
	if _cced >= 2 {
		_acbd, _fgdc = _cdefc.Bitmap.GetByte(_fecg - _cdefc.Bitmap.RowStride)
		if _fgdc != nil {
			return _ac.Wrap(_fgdc, _ecba, "\u006ci\u006ee\u004e\u0075\u006d\u0062\u0065\u0072\u0020\u003e\u003d\u0020\u0032")
		}
		_bbgd = int(_acbd) << 4
	}
	_debe = (_dce >> 3 & 0x7c) | (_bbgd >> 3 & 0x380)
	for _cfc := 0; _cfc < _gfgg; _cfc = _egea {
		var (
			_ega  byte
			_geae int
		)
		_egea = _cfc + 8
		if _geab := _bacg - _cfc; _geab > 8 {
			_geae = 8
		} else {
			_geae = _geab
		}
		if _cced > 0 {
			_dce <<= 8
			if _egea < _bacg {
				_acbd, _fgdc = _cdefc.Bitmap.GetByte(_fecg + 1)
				if _fgdc != nil {
					return _ac.Wrap(_fgdc, _ecba, "\u006c\u0069\u006e\u0065\u004e\u0075\u006d\u0062\u0065r\u0020\u003e\u0020\u0030")
				}
				_dce |= int(_acbd)
			}
		}
		if _cced > 1 {
			_bbgd <<= 8
			if _egea < _bacg {
				_acbd, _fgdc = _cdefc.Bitmap.GetByte(_fecg - _cdefc.Bitmap.RowStride + 1)
				if _fgdc != nil {
					return _ac.Wrap(_fgdc, _ecba, "\u006c\u0069\u006e\u0065\u004e\u0075\u006d\u0062\u0065r\u0020\u003e\u0020\u0031")
				}
				_bbgd |= int(_acbd) << 4
			}
		}
		for _adb := 0; _adb < _geae; _adb++ {
			_fecb := uint(10 - _adb)
			if _cdefc._aaaf {
				_dgfb = _cdefc.overrideAtTemplate2(_debe, _cfc+_adb, _cced, int(_ega), _adb)
				_cdefc._bcae.SetIndex(int32(_dgfb))
			} else {
				_cdefc._bcae.SetIndex(int32(_debe))
			}
			_ceeg, _fgdc = _cdefc._ga.DecodeBit(_cdefc._bcae)
			if _fgdc != nil {
				return _ac.Wrap(_fgdc, _ecba, "")
			}
			_ega |= byte(_ceeg << uint(7-_adb))
			_debe = ((_debe & 0x1bd) << 1) | _ceeg | ((_dce >> _fecb) & 0x4) | ((_bbgd >> _fecb) & 0x80)
		}
		if _ffa := _cdefc.Bitmap.SetByte(_bbg, _ega); _ffa != nil {
			return _ac.Wrap(_ffa, _ecba, "")
		}
		_bbg++
		_fecg++
	}
	return nil
}
