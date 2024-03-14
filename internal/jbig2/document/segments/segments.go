package segments

import (
	_ab "encoding/binary"
	_g "errors"
	_f "fmt"
	_c "image"
	_d "io"
	_b "math"
	_cf "strings"
	_dd "time"

	_eg "bitbucket.org/shenghui0779/gopdf/common"
	_e "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_dc "bitbucket.org/shenghui0779/gopdf/internal/jbig2/basic"
	_ea "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_ag "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder/arithmetic"
	_cg "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder/huffman"
	_eae "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder/mmr"
	_bb "bitbucket.org/shenghui0779/gopdf/internal/jbig2/encoder/arithmetic"
	_da "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
	_abf "bitbucket.org/shenghui0779/gopdf/internal/jbig2/internal"
	_ae "golang.org/x/xerrors"
)

func (_fdda *TextRegion) decodeRdx() (int64, error) {
	const _fffa = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064x"
	if _fdda.IsHuffmanEncoded {
		if _fdda.SbHuffRDX == 3 {
			if _fdda._gcgaga == nil {
				var (
					_eegf int
					_ccgf error
				)
				if _fdda.SbHuffFS == 3 {
					_eegf++
				}
				if _fdda.SbHuffDS == 3 {
					_eegf++
				}
				if _fdda.SbHuffDT == 3 {
					_eegf++
				}
				if _fdda.SbHuffRDWidth == 3 {
					_eegf++
				}
				if _fdda.SbHuffRDHeight == 3 {
					_eegf++
				}
				_fdda._gcgaga, _ccgf = _fdda.getUserTable(_eegf)
				if _ccgf != nil {
					return 0, _da.Wrap(_ccgf, _fffa, "")
				}
			}
			return _fdda._gcgaga.Decode(_fdda._gdbbd)
		}
		_dcab, _bdbd := _cg.GetStandardTable(14 + int(_fdda.SbHuffRDX))
		if _bdbd != nil {
			return 0, _da.Wrap(_bdbd, _fffa, "")
		}
		return _dcab.Decode(_fdda._gdbbd)
	}
	_ebdf, _dbeef := _fdda._bggc.DecodeInt(_fdda._dcfba)
	if _dbeef != nil {
		return 0, _da.Wrap(_dbeef, _fffa, "")
	}
	return int64(_ebdf), nil
}

func (_gecg *PageInformationSegment) encodeStripingInformation(_ccad _e.BinaryWriter) (_dagec int, _baag error) {
	const _edad = "\u0065n\u0063\u006f\u0064\u0065S\u0074\u0072\u0069\u0070\u0069n\u0067I\u006ef\u006f\u0072\u006d\u0061\u0074\u0069\u006fn"
	if !_gecg.IsStripe {
		if _dagec, _baag = _ccad.Write([]byte{0x00, 0x00}); _baag != nil {
			return 0, _da.Wrap(_baag, _edad, "n\u006f\u0020\u0073\u0074\u0072\u0069\u0070\u0069\u006e\u0067")
		}
		return _dagec, nil
	}
	_gfff := make([]byte, 2)
	_ab.BigEndian.PutUint16(_gfff, _gecg.MaxStripeSize|1<<15)
	if _dagec, _baag = _ccad.Write(_gfff); _baag != nil {
		return 0, _da.Wrapf(_baag, _edad, "\u0073\u0074\u0072i\u0070\u0069\u006e\u0067\u003a\u0020\u0025\u0064", _gecg.MaxStripeSize)
	}
	return _dagec, nil
}

func (_cbae *TextRegion) decodeIb(_gdbf, _cgge int64) (*_ea.Bitmap, error) {
	const _eddc = "\u0064\u0065\u0063\u006f\u0064\u0065\u0049\u0062"
	var (
		_fgab  error
		_eaece *_ea.Bitmap
	)
	if _gdbf == 0 {
		if int(_cgge) > len(_cbae.Symbols)-1 {
			return nil, _da.Error(_eddc, "\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0049\u0042\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0020\u0069\u006e\u0064\u0065x\u0020\u006f\u0075\u0074\u0020o\u0066\u0020r\u0061\u006e\u0067\u0065")
		}
		return _cbae.Symbols[int(_cgge)], nil
	}
	var _efcdc, _eage, _fgad, _fdae int64
	_efcdc, _fgab = _cbae.decodeRdw()
	if _fgab != nil {
		return nil, _da.Wrap(_fgab, _eddc, "")
	}
	_eage, _fgab = _cbae.decodeRdh()
	if _fgab != nil {
		return nil, _da.Wrap(_fgab, _eddc, "")
	}
	_fgad, _fgab = _cbae.decodeRdx()
	if _fgab != nil {
		return nil, _da.Wrap(_fgab, _eddc, "")
	}
	_fdae, _fgab = _cbae.decodeRdy()
	if _fgab != nil {
		return nil, _da.Wrap(_fgab, _eddc, "")
	}
	if _cbae.IsHuffmanEncoded {
		if _, _fgab = _cbae.decodeSymInRefSize(); _fgab != nil {
			return nil, _da.Wrap(_fgab, _eddc, "")
		}
		_cbae._gdbbd.Align()
	}
	_dgfgf := _cbae.Symbols[_cgge]
	_ffeg := uint32(_dgfgf.Width)
	_fagbd := uint32(_dgfgf.Height)
	_bdabd := int32(uint32(_efcdc)>>1) + int32(_fgad)
	_aaeg := int32(uint32(_eage)>>1) + int32(_fdae)
	if _cbae._eca == nil {
		_cbae._eca = _dbb(_cbae._gdbbd, nil)
	}
	_cbae._eca.setParameters(_cbae._ecaa, _cbae._bggc, _cbae.SbrTemplate, _ffeg+uint32(_efcdc), _fagbd+uint32(_eage), _dgfgf, _bdabd, _aaeg, false, _cbae.SbrATX, _cbae.SbrATY)
	_eaece, _fgab = _cbae._eca.GetRegionBitmap()
	if _fgab != nil {
		return nil, _da.Wrap(_fgab, _eddc, "\u0067\u0072\u0066")
	}
	if _cbae.IsHuffmanEncoded {
		_cbae._gdbbd.Align()
	}
	return _eaece, nil
}

type PatternDictionary struct {
	_ccca            *_e.Reader
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
	Patterns         []*_ea.Bitmap
	GrayMax          uint32
}

func (_cddf *PageInformationSegment) encodeFlags(_fef _e.BinaryWriter) (_fbdcgb error) {
	const _egb = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _fbdcgb = _fef.SkipBits(1); _fbdcgb != nil {
		return _da.Wrap(_fbdcgb, _egb, "\u0072\u0065\u0073e\u0072\u0076\u0065\u0064\u0020\u0062\u0069\u0074")
	}
	var _dbdf int
	if _cddf.CombinationOperatorOverrideAllowed() {
		_dbdf = 1
	}
	if _fbdcgb = _fef.WriteBit(_dbdf); _fbdcgb != nil {
		return _da.Wrap(_fbdcgb, _egb, "\u0063\u006f\u006db\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0070\u0065\u0072a\u0074\u006f\u0072\u0020\u006f\u0076\u0065\u0072\u0072\u0069\u0064\u0064\u0065\u006e")
	}
	_dbdf = 0
	if _cddf._dcfb {
		_dbdf = 1
	}
	if _fbdcgb = _fef.WriteBit(_dbdf); _fbdcgb != nil {
		return _da.Wrap(_fbdcgb, _egb, "\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0061\u0075\u0078i\u006ci\u0061r\u0079\u0020\u0062\u0075\u0066\u0066\u0065r")
	}
	if _fbdcgb = _fef.WriteBit((int(_cddf._ccde) >> 1) & 0x01); _fbdcgb != nil {
		return _da.Wrap(_fbdcgb, _egb, "\u0063\u006f\u006d\u0062\u0069\u006e\u0061\u0074\u0069\u006fn\u0020\u006f\u0070\u0065\u0072\u0061\u0074o\u0072\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0062\u0069\u0074")
	}
	if _fbdcgb = _fef.WriteBit(int(_cddf._ccde) & 0x01); _fbdcgb != nil {
		return _da.Wrap(_fbdcgb, _egb, "\u0063\u006f\u006db\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0070\u0065\u0072a\u0074\u006f\u0072\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0062\u0069\u0074")
	}
	_dbdf = int(_cddf.DefaultPixelValue)
	if _fbdcgb = _fef.WriteBit(_dbdf); _fbdcgb != nil {
		return _da.Wrap(_fbdcgb, _egb, "\u0064e\u0066\u0061\u0075\u006c\u0074\u0020\u0070\u0061\u0067\u0065\u0020p\u0069\u0078\u0065\u006c\u0020\u0076\u0061\u006c\u0075\u0065")
	}
	_dbdf = 0
	if _cddf._fagf {
		_dbdf = 1
	}
	if _fbdcgb = _fef.WriteBit(_dbdf); _fbdcgb != nil {
		return _da.Wrap(_fbdcgb, _egb, "\u0063\u006f\u006e\u0074ai\u006e\u0073\u0020\u0072\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074")
	}
	_dbdf = 0
	if _cddf.IsLossless {
		_dbdf = 1
	}
	if _fbdcgb = _fef.WriteBit(_dbdf); _fbdcgb != nil {
		return _da.Wrap(_fbdcgb, _egb, "p\u0061\u0067\u0065\u0020\u0069\u0073 \u0065\u0076\u0065\u006e\u0074\u0075\u0061\u006c\u006cy\u0020\u006c\u006fs\u0073l\u0065\u0073\u0073")
	}
	return nil
}

func (_adc *GenericRegion) overrideAtTemplate1(_faed, _agfb, _eaccc, _dcde, _aafa int) int {
	_faed &= 0x1FF7
	if _adc.GBAtY[0] == 0 && _adc.GBAtX[0] >= -int8(_aafa) {
		_faed |= (_dcde >> uint(7-(int8(_aafa)+_adc.GBAtX[0])) & 0x1) << 3
	} else {
		_faed |= int(_adc.getPixel(_agfb+int(_adc.GBAtX[0]), _eaccc+int(_adc.GBAtY[0]))) << 3
	}
	return _faed
}

func (_gbge *SymbolDictionary) decodeNewSymbols(_fggf, _cecd uint32, _afea *_ea.Bitmap, _ggda, _ddcc int32) error {
	if _gbge._eec == nil {
		_gbge._eec = _dbb(_gbge._abege, nil)
		if _gbge._agbbf == nil {
			var _adfa error
			_gbge._agbbf, _adfa = _ag.New(_gbge._abege)
			if _adfa != nil {
				return _adfa
			}
		}
		if _gbge._ccgbe == nil {
			_gbge._ccgbe = _ag.NewStats(65536, 1)
		}
	}
	_gbge._eec.setParameters(_gbge._ccgbe, _gbge._agbbf, _gbge.SdrTemplate, _fggf, _cecd, _afea, _ggda, _ddcc, false, _gbge.SdrATX, _gbge.SdrATY)
	return _gbge.addSymbol(_gbge._eec)
}

func (_caee *SymbolDictionary) checkInput() error {
	if _caee.SdHuffDecodeHeightSelection == 2 {
		_eg.Log.Debug("\u0053\u0079\u006d\u0062\u006fl\u0020\u0044\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0048\u0065\u0069\u0067\u0068\u0074\u0020\u0053e\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0070\u0065r\u006d\u0069\u0074\u0074\u0065\u0064", _caee.SdHuffDecodeHeightSelection)
	}
	if _caee.SdHuffDecodeWidthSelection == 2 {
		_eg.Log.Debug("\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0057\u0069\u0064t\u0068\u0020\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064\u0020\u0076\u0061l\u0075\u0065\u0020\u006e\u006f\u0074 \u0070\u0065r\u006d\u0069t\u0074e\u0064", _caee.SdHuffDecodeWidthSelection)
	}
	if _caee.IsHuffmanEncoded {
		if _caee.SdTemplate != 0 {
			_eg.Log.Debug("\u0053\u0044T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0028\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e \u0030\u0029", _caee.SdTemplate)
		}
		if !_caee.UseRefinementAggregation {
			if !_caee.UseRefinementAggregation {
				if _caee._cba {
					_eg.Log.Debug("\u0049\u0073\u0043\u006f\u0064\u0069\u006e\u0067C\u006f\u006e\u0074ex\u0074\u0052\u0065\u0074\u0061\u0069n\u0065\u0064\u0020\u003d\u0020\u0074\u0072\u0075\u0065\u0020\u0028\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0066\u0061\u006cs\u0065\u0029")
					_caee._cba = false
				}
				if _caee._bdgf {
					_eg.Log.Debug("\u0069s\u0043\u006fd\u0069\u006e\u0067\u0043o\u006e\u0074\u0065x\u0074\u0055\u0073\u0065\u0064\u0020\u003d\u0020\u0074ru\u0065\u0020\u0028s\u0068\u006fu\u006c\u0064\u0020\u0062\u0065\u0020f\u0061\u006cs\u0065\u0029")
					_caee._bdgf = false
				}
			}
		}
	} else {
		if _caee.SdHuffBMSizeSelection != 0 {
			_eg.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066B\u004d\u0053\u0069\u007a\u0065\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0030")
			_caee.SdHuffBMSizeSelection = 0
		}
		if _caee.SdHuffDecodeWidthSelection != 0 {
			_eg.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0057\u0069\u0064\u0074\u0068\u0053\u0065\u006c\u0065\u0063\u0074\u0069o\u006e\u0020\u0073\u0068\u006fu\u006c\u0064 \u0062\u0065\u0020\u0030")
			_caee.SdHuffDecodeWidthSelection = 0
		}
		if _caee.SdHuffDecodeHeightSelection != 0 {
			_eg.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0048e\u0069\u0067\u0068\u0074\u0053\u0065l\u0065\u0063\u0074\u0069\u006f\u006e\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0030")
			_caee.SdHuffDecodeHeightSelection = 0
		}
	}
	if !_caee.UseRefinementAggregation {
		if _caee.SdrTemplate != 0 {
			_eg.Log.Debug("\u0053\u0044\u0052\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u003d\u0020\u0025\u0064\u0020\u0028s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030\u0029", _caee.SdrTemplate)
			_caee.SdrTemplate = 0
		}
	}
	if !_caee.IsHuffmanEncoded || !_caee.UseRefinementAggregation {
		if _caee.SdHuffAggInstanceSelection {
			_eg.Log.Debug("\u0053d\u0048\u0075f\u0066\u0041\u0067g\u0049\u006e\u0073\u0074\u0061\u006e\u0063e\u0053\u0065\u006c\u0065\u0063\u0074i\u006f\u006e\u0020\u003d\u0020\u0025\u0064\u0020\u0028\u0073\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030\u0029", _caee.SdHuffAggInstanceSelection)
		}
	}
	return nil
}
func (_cdfb *PageInformationSegment) CombinationOperatorOverrideAllowed() bool { return _cdfb._efgad }
func (_ede *GenericRegion) readGBAtPixels(_aecc int) error {
	const _cbga = "\u0072\u0065\u0061\u0064\u0047\u0042\u0041\u0074\u0050i\u0078\u0065\u006c\u0073"
	_ede.GBAtX = make([]int8, _aecc)
	_ede.GBAtY = make([]int8, _aecc)
	for _caf := 0; _caf < _aecc; _caf++ {
		_acac, _bgcb := _ede._fec.ReadByte()
		if _bgcb != nil {
			return _da.Wrapf(_bgcb, _cbga, "\u0058\u0020\u0061t\u0020\u0069\u003a\u0020\u0027\u0025\u0064\u0027", _caf)
		}
		_ede.GBAtX[_caf] = int8(_acac)
		_acac, _bgcb = _ede._fec.ReadByte()
		if _bgcb != nil {
			return _da.Wrapf(_bgcb, _cbga, "\u0059\u0020\u0061t\u0020\u0069\u003a\u0020\u0027\u0025\u0064\u0027", _caf)
		}
		_ede.GBAtY[_caf] = int8(_acac)
	}
	return nil
}

var _ templater = &template0{}

func (_geaa *SymbolDictionary) decodeThroughTextRegion(_bcbe, _aacb, _beae uint32) error {
	if _geaa._dddd == nil {
		_geaa._dddd = _cefb(_geaa._abege, nil)
		_geaa._dddd.setContexts(_geaa._ccgbe, _ag.NewStats(512, 1), _ag.NewStats(512, 1), _ag.NewStats(512, 1), _ag.NewStats(512, 1), _geaa._gaac, _ag.NewStats(512, 1), _ag.NewStats(512, 1), _ag.NewStats(512, 1), _ag.NewStats(512, 1))
	}
	if _dfag := _geaa.setSymbolsArray(); _dfag != nil {
		return _dfag
	}
	_geaa._dddd.setParameters(_geaa._agbbf, _geaa.IsHuffmanEncoded, true, _bcbe, _aacb, _beae, 1, _geaa._afag+_geaa._addfb, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, _geaa.SdrTemplate, _geaa.SdrATX, _geaa.SdrATY, _geaa._ffbb, _geaa._cbe)
	return _geaa.addSymbol(_geaa._dddd)
}

func (_gefe *TextRegion) String() string {
	_cbgd := &_cf.Builder{}
	_cbgd.WriteString("\u000a[\u0054E\u0058\u0054\u0020\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u000a")
	_cbgd.WriteString(_gefe.RegionInfo.String() + "\u000a")
	_cbgd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053br\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0025\u0076\u000a", _gefe.SbrTemplate))
	_cbgd.WriteString(_f.Sprintf("\u0009-\u0020S\u0062\u0044\u0073\u004f\u0066f\u0073\u0065t\u003a\u0020\u0025\u0076\u000a", _gefe.SbDsOffset))
	_cbgd.WriteString(_f.Sprintf("\t\u002d \u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0050i\u0078\u0065\u006c\u003a %\u0076\u000a", _gefe.DefaultPixel))
	_cbgd.WriteString(_f.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065r\u0061\u0074\u006fr\u003a \u0025\u0076\u000a", _gefe.CombinationOperator))
	_cbgd.WriteString(_f.Sprintf("\t\u002d \u0049\u0073\u0054\u0072\u0061\u006e\u0073\u0070o\u0073\u0065\u0064\u003a %\u0076\u000a", _gefe.IsTransposed))
	_cbgd.WriteString(_f.Sprintf("\u0009\u002d\u0020Re\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0043\u006f\u0072\u006e\u0065\u0072\u003a\u0020\u0025\u0076\u000a", _gefe.ReferenceCorner))
	_cbgd.WriteString(_f.Sprintf("\t\u002d\u0020\u0055\u0073eR\u0065f\u0069\u006e\u0065\u006d\u0065n\u0074\u003a\u0020\u0025\u0076\u000a", _gefe.UseRefinement))
	_cbgd.WriteString(_f.Sprintf("\u0009-\u0020\u0049\u0073\u0048\u0075\u0066\u0066\u006d\u0061\u006e\u0045n\u0063\u006f\u0064\u0065\u0064\u003a\u0020\u0025\u0076\u000a", _gefe.IsHuffmanEncoded))
	if _gefe.IsHuffmanEncoded {
		_cbgd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053bH\u0075\u0066\u0066\u0052\u0053\u0069\u007a\u0065\u003a\u0020\u0025\u0076\u000a", _gefe.SbHuffRSize))
		_cbgd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0059:\u0020\u0025\u0076\u000a", _gefe.SbHuffRDY))
		_cbgd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0058:\u0020\u0025\u0076\u000a", _gefe.SbHuffRDX))
		_cbgd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053bH\u0075\u0066\u0066\u0052\u0044\u0048\u0065\u0069\u0067\u0068\u0074\u003a\u0020\u0025v\u000a", _gefe.SbHuffRDHeight))
		_cbgd.WriteString(_f.Sprintf("\t\u002d\u0020\u0053\u0062Hu\u0066f\u0052\u0044\u0057\u0069\u0064t\u0068\u003a\u0020\u0025\u0076\u000a", _gefe.SbHuffRDWidth))
		_cbgd.WriteString(_f.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0044\u0054\u003a\u0020\u0025\u0076\u000a", _gefe.SbHuffDT))
		_cbgd.WriteString(_f.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0044\u0053\u003a\u0020\u0025\u0076\u000a", _gefe.SbHuffDS))
		_cbgd.WriteString(_f.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0046\u0053\u003a\u0020\u0025\u0076\u000a", _gefe.SbHuffFS))
	}
	_cbgd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0058:\u0020\u0025\u0076\u000a", _gefe.SbrATX))
	_cbgd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0059:\u0020\u0025\u0076\u000a", _gefe.SbrATY))
	_cbgd.WriteString(_f.Sprintf("\u0009\u002d\u0020N\u0075\u006d\u0062\u0065r\u004f\u0066\u0053\u0079\u006d\u0062\u006fl\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u003a\u0020\u0025\u0076\u000a", _gefe.NumberOfSymbolInstances))
	_cbgd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0058:\u0020\u0025\u0076\u000a", _gefe.SbrATX))
	return _cbgd.String()
}

func (_fdfe *PageInformationSegment) readWidthAndHeight() error {
	_efdgd, _agfg := _fdfe._dfdf.ReadBits(32)
	if _agfg != nil {
		return _agfg
	}
	_fdfe.PageBMWidth = int(_efdgd & _b.MaxInt32)
	_efdgd, _agfg = _fdfe._dfdf.ReadBits(32)
	if _agfg != nil {
		return _agfg
	}
	_fdfe.PageBMHeight = int(_efdgd & _b.MaxInt32)
	return nil
}

func (_fae *GenericRegion) decodeTemplate0a(_gfc, _faa, _bdd int, _bcg, _dgfg int) (_fab error) {
	const _gfb = "\u0064\u0065c\u006f\u0064\u0065T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0030\u0061"
	var (
		_baee, _cga   int
		_dcggd, _ffge int
		_caec         byte
		_dafb         int
	)
	if _gfc >= 1 {
		_caec, _fab = _fae.Bitmap.GetByte(_dgfg)
		if _fab != nil {
			return _da.Wrap(_fab, _gfb, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_dcggd = int(_caec)
	}
	if _gfc >= 2 {
		_caec, _fab = _fae.Bitmap.GetByte(_dgfg - _fae.Bitmap.RowStride)
		if _fab != nil {
			return _da.Wrap(_fab, _gfb, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_ffge = int(_caec) << 6
	}
	_baee = (_dcggd & 0xf0) | (_ffge & 0x3800)
	for _fdff := 0; _fdff < _bdd; _fdff = _dafb {
		var (
			_aeba byte
			_fbe  int
		)
		_dafb = _fdff + 8
		if _egce := _faa - _fdff; _egce > 8 {
			_fbe = 8
		} else {
			_fbe = _egce
		}
		if _gfc > 0 {
			_dcggd <<= 8
			if _dafb < _faa {
				_caec, _fab = _fae.Bitmap.GetByte(_dgfg + 1)
				if _fab != nil {
					return _da.Wrap(_fab, _gfb, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_dcggd |= int(_caec)
			}
		}
		if _gfc > 1 {
			_abce := _dgfg - _fae.Bitmap.RowStride + 1
			_ffge <<= 8
			if _dafb < _faa {
				_caec, _fab = _fae.Bitmap.GetByte(_abce)
				if _fab != nil {
					return _da.Wrap(_fab, _gfb, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_ffge |= int(_caec) << 6
			} else {
				_ffge |= 0
			}
		}
		for _fdgd := 0; _fdgd < _fbe; _fdgd++ {
			_beb := uint(7 - _fdgd)
			if _fae._aaf {
				_cga = _fae.overrideAtTemplate0a(_baee, _fdff+_fdgd, _gfc, int(_aeba), _fdgd, int(_beb))
				_fae._eced.SetIndex(int32(_cga))
			} else {
				_fae._eced.SetIndex(int32(_baee))
			}
			var _gag int
			_gag, _fab = _fae._bda.DecodeBit(_fae._eced)
			if _fab != nil {
				return _da.Wrap(_fab, _gfb, "")
			}
			_aeba |= byte(_gag) << _beb
			_baee = ((_baee & 0x7bf7) << 1) | _gag | ((_dcggd >> _beb) & 0x10) | ((_ffge >> _beb) & 0x800)
		}
		if _bdf := _fae.Bitmap.SetByte(_bcg, _aeba); _bdf != nil {
			return _da.Wrap(_bdf, _gfb, "")
		}
		_bcg++
		_dgfg++
	}
	return nil
}
func (_dgee *GenericRegion) GetRegionInfo() *RegionSegment { return _dgee.RegionSegment }
func (_cafg *GenericRegion) setParametersWithAt(_dgdf bool, _fggd byte, _baga, _edc bool, _ddda, _gbe []int8, _bgb, _efdg uint32, _acef *_ag.DecoderStats, _decc *_ag.Decoder) {
	_cafg.IsMMREncoded = _dgdf
	_cafg.GBTemplate = _fggd
	_cafg.IsTPGDon = _baga
	_cafg.GBAtX = _ddda
	_cafg.GBAtY = _gbe
	_cafg.RegionSegment.BitmapHeight = _efdg
	_cafg.RegionSegment.BitmapWidth = _bgb
	_cafg._fdg = nil
	_cafg.Bitmap = nil
	if _acef != nil {
		_cafg._eced = _acef
	}
	if _decc != nil {
		_cafg._bda = _decc
	}
	_eg.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0047\u0049O\u004e\u005d\u0020\u0073\u0065\u0074P\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0053\u0044\u0041t\u003a\u0020\u0025\u0073", _cafg)
}

var _ SegmentEncoder = &GenericRegion{}

func (_bdag *GenericRegion) setParameters(_feg bool, _ceb, _ecff int64, _edee, _afbe uint32) {
	_bdag.IsMMREncoded = _feg
	_bdag.DataOffset = _ceb
	_bdag.DataLength = _ecff
	_bdag.RegionSegment.BitmapHeight = _edee
	_bdag.RegionSegment.BitmapWidth = _afbe
	_bdag._fdg = nil
	_bdag.Bitmap = nil
}

func (_edgc *SymbolDictionary) getSbSymCodeLen() int8 {
	_eagfc := int8(_b.Ceil(_b.Log(float64(_edgc._afag+_edgc.NumberOfNewSymbols)) / _b.Log(2)))
	if _edgc.IsHuffmanEncoded && _eagfc < 1 {
		return 1
	}
	return _eagfc
}

func (_fecc *TextRegion) decodeStripT() (_fbcc int64, _bfaa error) {
	if _fecc.IsHuffmanEncoded {
		if _fecc.SbHuffDT == 3 {
			if _fecc._accc == nil {
				var _adbf int
				if _fecc.SbHuffFS == 3 {
					_adbf++
				}
				if _fecc.SbHuffDS == 3 {
					_adbf++
				}
				_fecc._accc, _bfaa = _fecc.getUserTable(_adbf)
				if _bfaa != nil {
					return 0, _bfaa
				}
			}
			_fbcc, _bfaa = _fecc._accc.Decode(_fecc._gdbbd)
			if _bfaa != nil {
				return 0, _bfaa
			}
		} else {
			var _fecgb _cg.Tabler
			_fecgb, _bfaa = _cg.GetStandardTable(11 + int(_fecc.SbHuffDT))
			if _bfaa != nil {
				return 0, _bfaa
			}
			_fbcc, _bfaa = _fecgb.Decode(_fecc._gdbbd)
			if _bfaa != nil {
				return 0, _bfaa
			}
		}
	} else {
		var _ffea int32
		_ffea, _bfaa = _fecc._bggc.DecodeInt(_fecc._ffae)
		if _bfaa != nil {
			return 0, _bfaa
		}
		_fbcc = int64(_ffea)
	}
	_fbcc *= int64(-_fecc.SbStrips)
	return _fbcc, nil
}

func (_cag *HalftoneRegion) renderPattern(_adgf [][]int) (_dcdb error) {
	var _cfbbg, _ddcb int
	for _faec := 0; _faec < int(_cag.HGridHeight); _faec++ {
		for _dcbg := 0; _dcbg < int(_cag.HGridWidth); _dcbg++ {
			_cfbbg = _cag.computeX(_faec, _dcbg)
			_ddcb = _cag.computeY(_faec, _dcbg)
			_baff := _cag.Patterns[_adgf[_faec][_dcbg]]
			if _dcdb = _ea.Blit(_baff, _cag.HalftoneRegionBitmap, _cfbbg+int(_cag.HGridX), _ddcb+int(_cag.HGridY), _cag.CombinationOperator); _dcdb != nil {
				return _dcdb
			}
		}
	}
	return nil
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

func (_gaaa *PageInformationSegment) Init(h *Header, r *_e.Reader) (_cdb error) {
	_gaaa._dfdf = r
	if _cdb = _gaaa.parseHeader(); _cdb != nil {
		return _da.Wrap(_cdb, "P\u0061\u0067\u0065\u0049\u006e\u0066o\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065g\u006d\u0065\u006et\u002eI\u006e\u0069\u0074", "")
	}
	return nil
}

func (_aga *GenericRegion) decodeSLTP() (int, error) {
	switch _aga.GBTemplate {
	case 0:
		_aga._eced.SetIndex(0x9B25)
	case 1:
		_aga._eced.SetIndex(0x795)
	case 2:
		_aga._eced.SetIndex(0xE5)
	case 3:
		_aga._eced.SetIndex(0x195)
	}
	return _aga._bda.DecodeBit(_aga._eced)
}

func (_gggg *GenericRegion) decodeLine(_cgea, _efga, _dbddc int) error {
	const _gfg = "\u0064\u0065\u0063\u006f\u0064\u0065\u004c\u0069\u006e\u0065"
	_acd := _gggg.Bitmap.GetByteIndex(0, _cgea)
	_dgd := _acd - _gggg.Bitmap.RowStride
	switch _gggg.GBTemplate {
	case 0:
		if !_gggg.UseExtTemplates {
			return _gggg.decodeTemplate0a(_cgea, _efga, _dbddc, _acd, _dgd)
		}
		return _gggg.decodeTemplate0b(_cgea, _efga, _dbddc, _acd, _dgd)
	case 1:
		return _gggg.decodeTemplate1(_cgea, _efga, _dbddc, _acd, _dgd)
	case 2:
		return _gggg.decodeTemplate2(_cgea, _efga, _dbddc, _acd, _dgd)
	case 3:
		return _gggg.decodeTemplate3(_cgea, _efga, _dbddc, _acd, _dgd)
	}
	return _da.Errorf(_gfg, "\u0069\u006e\u0076a\u006c\u0069\u0064\u0020G\u0042\u0054\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u003a\u0020\u0025\u0064", _gggg.GBTemplate)
}

func (_acgef *SymbolDictionary) setRefinementAtPixels() error {
	if !_acgef.UseRefinementAggregation || _acgef.SdrTemplate != 0 {
		return nil
	}
	if _bagc := _acgef.readRefinementAtPixels(2); _bagc != nil {
		return _bagc
	}
	return nil
}
func (_cdg *HalftoneRegion) GetRegionInfo() *RegionSegment { return _cdg.RegionSegment }
func (_dgaff *TableSegment) HtRS() int32                   { return _dgaff._gbda }
func (_aba *template1) setIndex(_cfc *_ag.DecoderStats)    { _cfc.SetIndex(0x080) }
func (_adgb *GenericRegion) overrideAtTemplate3(_feaa, _dfec, _cfbb, _dddg, _afbf int) int {
	_feaa &= 0x3EF
	if _adgb.GBAtY[0] == 0 && _adgb.GBAtX[0] >= -int8(_afbf) {
		_feaa |= (_dddg >> uint(7-(int8(_afbf)+_adgb.GBAtX[0])) & 0x1) << 4
	} else {
		_feaa |= int(_adgb.getPixel(_dfec+int(_adgb.GBAtX[0]), _cfbb+int(_adgb.GBAtY[0]))) << 4
	}
	return _feaa
}

func _cefb(_ecca *_e.Reader, _cafa *Header) *TextRegion {
	_gcffb := &TextRegion{_gdbbd: _ecca, Header: _cafa, RegionInfo: NewRegionSegment(_ecca)}
	return _gcffb
}

type GenericRegion struct {
	_fec             *_e.Reader
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
	_aaf             bool
	Bitmap           *_ea.Bitmap
	_bda             *_ag.Decoder
	_eced            *_ag.DecoderStats
	_fdg             *_eae.Decoder
}

var _ templater = &template1{}

func (_bc *GenericRefinementRegion) decodeSLTP() (int, error) {
	_bc.Template.setIndex(_bc._gbf)
	return _bc._cc.DecodeBit(_bc._gbf)
}

func (_ffd *Header) readHeaderLength(_ccefg *_e.Reader, _abbfb int64) {
	_ffd.HeaderLength = _ccefg.AbsolutePosition() - _abbfb
}

func (_aeebb *TextRegion) blit(_bffc *_ea.Bitmap, _bfea int64) error {
	if _aeebb.IsTransposed == 0 && (_aeebb.ReferenceCorner == 2 || _aeebb.ReferenceCorner == 3) {
		_aeebb._afcg += int64(_bffc.Width - 1)
	} else if _aeebb.IsTransposed == 1 && (_aeebb.ReferenceCorner == 0 || _aeebb.ReferenceCorner == 2) {
		_aeebb._afcg += int64(_bffc.Height - 1)
	}
	_fddc := _aeebb._afcg
	if _aeebb.IsTransposed == 1 {
		_fddc, _bfea = _bfea, _fddc
	}
	switch _aeebb.ReferenceCorner {
	case 0:
		_bfea -= int64(_bffc.Height - 1)
	case 2:
		_bfea -= int64(_bffc.Height - 1)
		_fddc -= int64(_bffc.Width - 1)
	case 3:
		_fddc -= int64(_bffc.Width - 1)
	}
	_gebd := _ea.Blit(_bffc, _aeebb.RegionBitmap, int(_fddc), int(_bfea), _aeebb.CombinationOperator)
	if _gebd != nil {
		return _gebd
	}
	if _aeebb.IsTransposed == 0 && (_aeebb.ReferenceCorner == 0 || _aeebb.ReferenceCorner == 1) {
		_aeebb._afcg += int64(_bffc.Width - 1)
	}
	if _aeebb.IsTransposed == 1 && (_aeebb.ReferenceCorner == 1 || _aeebb.ReferenceCorner == 3) {
		_aeebb._afcg += int64(_bffc.Height - 1)
	}
	return nil
}
func (_ed *EndOfStripe) Init(h *Header, r *_e.Reader) error { _ed._agb = r; return _ed.parseHeader() }
func (_gcfc *TextRegion) checkInput() error {
	const _dccff = "\u0063\u0068\u0065\u0063\u006b\u0049\u006e\u0070\u0075\u0074"
	if !_gcfc.UseRefinement {
		if _gcfc.SbrTemplate != 0 {
			_eg.Log.Debug("\u0053\u0062\u0072Te\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030")
			_gcfc.SbrTemplate = 0
		}
	}
	if _gcfc.SbHuffFS == 2 || _gcfc.SbHuffRDWidth == 2 || _gcfc.SbHuffRDHeight == 2 || _gcfc.SbHuffRDX == 2 || _gcfc.SbHuffRDY == 2 {
		return _da.Error(_dccff, "h\u0075\u0066\u0066\u006d\u0061\u006e \u0066\u006c\u0061\u0067\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064")
	}
	if !_gcfc.UseRefinement {
		if _gcfc.SbHuffRSize != 0 {
			_eg.Log.Debug("\u0053\u0062\u0048uf\u0066\u0052\u0053\u0069\u007a\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030")
			_gcfc.SbHuffRSize = 0
		}
		if _gcfc.SbHuffRDY != 0 {
			_eg.Log.Debug("S\u0062\u0048\u0075\u0066fR\u0044Y\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0030")
			_gcfc.SbHuffRDY = 0
		}
		if _gcfc.SbHuffRDX != 0 {
			_eg.Log.Debug("S\u0062\u0048\u0075\u0066fR\u0044X\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0030")
			_gcfc.SbHuffRDX = 0
		}
		if _gcfc.SbHuffRDWidth != 0 {
			_eg.Log.Debug("\u0053b\u0048\u0075\u0066\u0066R\u0044\u0057\u0069\u0064\u0074h\u0020s\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u00200")
			_gcfc.SbHuffRDWidth = 0
		}
		if _gcfc.SbHuffRDHeight != 0 {
			_eg.Log.Debug("\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0048\u0065\u0069g\u0068\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020b\u0065\u0020\u0030")
			_gcfc.SbHuffRDHeight = 0
		}
	}
	return nil
}

func (_bdab *HalftoneRegion) checkInput() error {
	if _bdab.IsMMREncoded {
		if _bdab.HTemplate != 0 {
			_eg.Log.Debug("\u0048\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0030", _bdab.HTemplate)
		}
		if _bdab.HSkipEnabled {
			_eg.Log.Debug("\u0048\u0053\u006b\u0069\u0070\u0045\u006e\u0061\u0062\u006c\u0065\u0064\u0020\u0030\u0020\u0025\u0076\u0020(\u0073\u0068\u006f\u0075\u006c\u0064\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u0029", _bdab.HSkipEnabled)
		}
	}
	return nil
}

func (_bdc *GenericRegion) String() string {
	_dbfg := &_cf.Builder{}
	_dbfg.WriteString("\u000a[\u0047E\u004e\u0045\u0052\u0049\u0043 \u0052\u0045G\u0049\u004f\u004e\u005d\u000a")
	_dbfg.WriteString(_bdc.RegionSegment.String() + "\u000a")
	_dbfg.WriteString(_f.Sprintf("\u0009\u002d\u0020Us\u0065\u0045\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073\u003a\u0020\u0025\u0076\u000a", _bdc.UseExtTemplates))
	_dbfg.WriteString(_f.Sprintf("\u0009\u002d \u0049\u0073\u0054P\u0047\u0044\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _bdc.IsTPGDon))
	_dbfg.WriteString(_f.Sprintf("\u0009-\u0020G\u0042\u0054\u0065\u006d\u0070l\u0061\u0074e\u003a\u0020\u0025\u0064\u000a", _bdc.GBTemplate))
	_dbfg.WriteString(_f.Sprintf("\t\u002d \u0049\u0073\u004d\u004d\u0052\u0045\u006e\u0063o\u0064\u0065\u0064\u003a %\u0076\u000a", _bdc.IsMMREncoded))
	_dbfg.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0047\u0042\u0041\u0074\u0058\u003a\u0020\u0025\u0076\u000a", _bdc.GBAtX))
	_dbfg.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0047\u0042\u0041\u0074\u0059\u003a\u0020\u0025\u0076\u000a", _bdc.GBAtY))
	_dbfg.WriteString(_f.Sprintf("\t\u002d \u0047\u0042\u0041\u0074\u004f\u0076\u0065\u0072r\u0069\u0064\u0065\u003a %\u0076\u000a", _bdc.GBAtOverride))
	return _dbfg.String()
}

func (_dffgf *SymbolDictionary) decodeHeightClassDeltaHeight() (int64, error) {
	if _dffgf.IsHuffmanEncoded {
		return _dffgf.decodeHeightClassDeltaHeightWithHuffman()
	}
	_fagc, _dged := _dffgf._agbbf.DecodeInt(_dffgf._abdd)
	if _dged != nil {
		return 0, _dged
	}
	return int64(_fagc), nil
}

func (_dfgde *TextRegion) getUserTable(_gdege int) (_cg.Tabler, error) {
	const _fbec = "\u0067\u0065\u0074U\u0073\u0065\u0072\u0054\u0061\u0062\u006c\u0065"
	var _cfad int
	for _, _fggfd := range _dfgde.Header.RTSegments {
		if _fggfd.Type == 53 {
			if _cfad == _gdege {
				_acce, _agab := _fggfd.GetSegmentData()
				if _agab != nil {
					return nil, _agab
				}
				_cabe, _gfce := _acce.(*TableSegment)
				if !_gfce {
					_eg.Log.Debug(_f.Sprintf("\u0073\u0065\u0067\u006d\u0065\u006e\u0074 \u0077\u0069\u0074h\u0020\u0054\u0079p\u0065\u00205\u0033\u0020\u002d\u0020\u0061\u006ed\u0020in\u0064\u0065\u0078\u003a\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0054\u0061\u0062\u006c\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074", _fggfd.SegmentNumber))
					return nil, _da.Error(_fbec, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074 \u0077\u0069\u0074h\u0020\u0054\u0079\u0070e\u0020\u0035\u0033\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u002a\u0054\u0061\u0062\u006c\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
				}
				return _cg.NewEncodedTable(_cabe)
			}
			_cfad++
		}
	}
	return nil, nil
}

type templater interface {
	form(_gafd, _dafa, _caa, _ffb, _eba int16) int16
	setIndex(_aag *_ag.DecoderStats)
}

func (_bafg *PageInformationSegment) checkInput() error {
	if _bafg.PageBMHeight == _b.MaxInt32 {
		if !_bafg.IsStripe {
			_eg.Log.Debug("P\u0061\u0067\u0065\u0049\u006e\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0053\u0065\u0067me\u006e\u0074\u002e\u0049s\u0053\u0074\u0072\u0069\u0070\u0065\u0020\u0073\u0068ou\u006c\u0064 \u0062\u0065\u0020\u0074\u0072\u0075\u0065\u002e")
		}
	}
	return nil
}

func (_bbb *Header) readSegmentDataLength(_effc *_e.Reader) (_dcede error) {
	_bbb.SegmentDataLength, _dcede = _effc.ReadBits(32)
	if _dcede != nil {
		return _dcede
	}
	_bbb.SegmentDataLength &= _b.MaxInt32
	return nil
}

func (_beag *PatternDictionary) checkInput() error {
	if _beag.HdpHeight < 1 || _beag.HdpWidth < 1 {
		return _g.New("in\u0076\u0061l\u0069\u0064\u0020\u0048\u0065\u0061\u0064\u0065\u0072 \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0057\u0069\u0064\u0074\u0068\u002f\u0048\u0065\u0069\u0067\u0068\u0074\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020g\u0072e\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020z\u0065\u0072o")
	}
	if _beag.IsMMREncoded {
		if _beag.HDTemplate != 0 {
			_eg.Log.Debug("\u0076\u0061\u0072\u0069\u0061\u0062\u006c\u0065\u0020\u0048\u0044\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0030")
		}
	}
	return nil
}

func (_ecd *GenericRefinementRegion) parseHeader() (_bee error) {
	_eg.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072s\u0069\u006e\u0067\u0020\u0048e\u0061\u0064e\u0072\u002e\u002e\u002e")
	_eab := _dd.Now()
	defer func() {
		if _bee == nil {
			_eg.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045G\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020h\u0065\u0061\u0064\u0065\u0072\u0020\u0066\u0069\u006e\u0069\u0073\u0068id\u0020\u0069\u006e\u003a\u0020\u0025\u0064\u0020\u006e\u0073", _dd.Since(_eab).Nanoseconds())
		} else {
			_eg.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0073", _bee)
		}
	}()
	if _bee = _ecd.RegionInfo.parseHeader(); _bee != nil {
		return _bee
	}
	_, _bee = _ecd._gbc.ReadBits(6)
	if _bee != nil {
		return _bee
	}
	_ecd.IsTPGROn, _bee = _ecd._gbc.ReadBool()
	if _bee != nil {
		return _bee
	}
	var _aac int
	_aac, _bee = _ecd._gbc.ReadBit()
	if _bee != nil {
		return _bee
	}
	_ecd.TemplateID = int8(_aac)
	switch _ecd.TemplateID {
	case 0:
		_ecd.Template = _ecd._agbg
		if _bee = _ecd.readAtPixels(); _bee != nil {
			return
		}
	case 1:
		_ecd.Template = _ecd._ad
	}
	return nil
}

func (_ffga *SymbolDictionary) getToExportFlags() ([]int, error) {
	var (
		_gfadb int
		_ebcf  int32
		_bgff  error
		_babc  = int32(_ffga._afag + _ffga.NumberOfNewSymbols)
		_cbda  = make([]int, _babc)
	)
	for _bgec := int32(0); _bgec < _babc; _bgec += _ebcf {
		if _ffga.IsHuffmanEncoded {
			_aegc, _abcg := _cg.GetStandardTable(1)
			if _abcg != nil {
				return nil, _abcg
			}
			_egcb, _abcg := _aegc.Decode(_ffga._abege)
			if _abcg != nil {
				return nil, _abcg
			}
			_ebcf = int32(_egcb)
		} else {
			_ebcf, _bgff = _ffga._agbbf.DecodeInt(_ffga._dbcca)
			if _bgff != nil {
				return nil, _bgff
			}
		}
		if _ebcf != 0 {
			if _bgec+_ebcf > _babc {
				return nil, _da.Error("\u0053\u0079\u006d\u0062\u006f\u006cD\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0067\u0065\u0074T\u006f\u0045\u0078\u0070\u006f\u0072\u0074F\u006c\u0061\u0067\u0073", "\u006d\u0061\u006c\u0066\u006f\u0072m\u0065\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0064\u0061\u0074\u0061\u0020\u0070\u0072\u006f\u0076\u0069\u0064e\u0064\u002e\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u006f\u0075\u0074\u0020\u006ff\u0020r\u0061\u006e\u0067\u0065")
			}
			for _abad := _bgec; _abad < _bgec+_ebcf; _abad++ {
				_cbda[_abad] = _gfadb
			}
		}
		if _gfadb == 0 {
			_gfadb = 1
		} else {
			_gfadb = 0
		}
	}
	return _cbda, nil
}

func (_cfd *GenericRefinementRegion) String() string {
	_abe := &_cf.Builder{}
	_abe.WriteString("\u000a[\u0047E\u004e\u0045\u0052\u0049\u0043 \u0052\u0045G\u0049\u004f\u004e\u005d\u000a")
	_abe.WriteString(_cfd.RegionInfo.String() + "\u000a")
	_abe.WriteString(_f.Sprintf("\u0009\u002d \u0049\u0073\u0054P\u0047\u0052\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _cfd.IsTPGROn))
	_abe.WriteString(_f.Sprintf("\u0009-\u0020T\u0065\u006d\u0070\u006c\u0061t\u0065\u0049D\u003a\u0020\u0025\u0076\u000a", _cfd.TemplateID))
	_abe.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0047\u0072\u0041\u0074\u0058\u003a\u0020\u0025\u0076\u000a", _cfd.GrAtX))
	_abe.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0047\u0072\u0041\u0074\u0059\u003a\u0020\u0025\u0076\u000a", _cfd.GrAtY))
	_abe.WriteString(_f.Sprintf("\u0009-\u0020R\u0065\u0066\u0065\u0072\u0065n\u0063\u0065D\u0058\u0020\u0025\u0076\u000a", _cfd.ReferenceDX))
	_abe.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0052ef\u0065\u0072\u0065\u006e\u0063\u0044\u0065\u0059\u003a\u0020\u0025\u0076\u000a", _cfd.ReferenceDY))
	return _abe.String()
}

type HalftoneRegion struct {
	_bfbe                *_e.Reader
	_aadb                *Header
	DataHeaderOffset     int64
	DataHeaderLength     int64
	DataOffset           int64
	DataLength           int64
	RegionSegment        *RegionSegment
	HDefaultPixel        int8
	CombinationOperator  _ea.CombinationOperator
	HSkipEnabled         bool
	HTemplate            byte
	IsMMREncoded         bool
	HGridWidth           uint32
	HGridHeight          uint32
	HGridX               int32
	HGridY               int32
	HRegionX             uint16
	HRegionY             uint16
	HalftoneRegionBitmap *_ea.Bitmap
	Patterns             []*_ea.Bitmap
}

func (_bgc *GenericRefinementRegion) GetRegionBitmap() (*_ea.Bitmap, error) {
	var _fee error
	_eg.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0047\u0065\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		if _fee != nil {
			_eg.Log.Trace("[\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052E\u0046\u002d\u0052\u0045\u0047\u0049\u004fN]\u0020\u0047\u0065\u0074R\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006dap\u0020\u0066a\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _fee)
		} else {
			_eg.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0047\u0065\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		}
	}()
	if _bgc.RegionBitmap != nil {
		return _bgc.RegionBitmap, nil
	}
	_gd := 0
	if _bgc.ReferenceBitmap == nil {
		_bgc.ReferenceBitmap, _fee = _bgc.getGrReference()
		if _fee != nil {
			return nil, _fee
		}
	}
	if _bgc._cc == nil {
		_bgc._cc, _fee = _ag.New(_bgc._gbc)
		if _fee != nil {
			return nil, _fee
		}
	}
	if _bgc._gbf == nil {
		_bgc._gbf = _ag.NewStats(8192, 1)
	}
	_bgc.RegionBitmap = _ea.New(int(_bgc.RegionInfo.BitmapWidth), int(_bgc.RegionInfo.BitmapHeight))
	if _bgc.TemplateID == 0 {
		if _fee = _bgc.updateOverride(); _fee != nil {
			return nil, _fee
		}
	}
	_gg := (_bgc.RegionBitmap.Width + 7) & -8
	var _bbe int
	if _bgc.IsTPGROn {
		_bbe = int(-_bgc.ReferenceDY) * _bgc.ReferenceBitmap.RowStride
	}
	_ba := _bbe + 1
	for _ff := 0; _ff < _bgc.RegionBitmap.Height; _ff++ {
		if _bgc.IsTPGROn {
			_gef, _ca := _bgc.decodeSLTP()
			if _ca != nil {
				return nil, _ca
			}
			_gd ^= _gef
		}
		if _gd == 0 {
			_fee = _bgc.decodeOptimized(_ff, _bgc.RegionBitmap.Width, _bgc.RegionBitmap.RowStride, _bgc.ReferenceBitmap.RowStride, _gg, _bbe, _ba)
			if _fee != nil {
				return nil, _fee
			}
		} else {
			_fee = _bgc.decodeTypicalPredictedLine(_ff, _bgc.RegionBitmap.Width, _bgc.RegionBitmap.RowStride, _bgc.ReferenceBitmap.RowStride, _gg, _bbe)
			if _fee != nil {
				return nil, _fee
			}
		}
	}
	return _bgc.RegionBitmap, nil
}

func (_efdc *SymbolDictionary) setRetainedCodingContexts(_babe *SymbolDictionary) {
	_efdc._agbbf = _babe._agbbf
	_efdc.IsHuffmanEncoded = _babe.IsHuffmanEncoded
	_efdc.UseRefinementAggregation = _babe.UseRefinementAggregation
	_efdc.SdTemplate = _babe.SdTemplate
	_efdc.SdrTemplate = _babe.SdrTemplate
	_efdc.SdATX = _babe.SdATX
	_efdc.SdATY = _babe.SdATY
	_efdc.SdrATX = _babe.SdrATX
	_efdc.SdrATY = _babe.SdrATY
	_efdc._ccgbe = _babe._ccgbe
}

func (_bca *GenericRefinementRegion) decodeOptimized(_aec, _edd, _ccf, _ccc, _ef, _aeg, _fd int) error {
	var (
		_de  error
		_efe int
		_gba int
	)
	_gc := _aec - int(_bca.ReferenceDY)
	if _cae := int(-_bca.ReferenceDX); _cae > 0 {
		_efe = _cae
	}
	_dda := _bca.ReferenceBitmap.GetByteIndex(_efe, _gc)
	if _bca.ReferenceDX > 0 {
		_gba = int(_bca.ReferenceDX)
	}
	_af := _bca.RegionBitmap.GetByteIndex(_gba, _aec)
	switch _bca.TemplateID {
	case 0:
		_de = _bca.decodeTemplate(_aec, _edd, _ccf, _ccc, _ef, _aeg, _fd, _af, _gc, _dda, _bca._agbg)
	case 1:
		_de = _bca.decodeTemplate(_aec, _edd, _ccf, _ccc, _ef, _aeg, _fd, _af, _gc, _dda, _bca._ad)
	}
	return _de
}

func (_bbfc *SymbolDictionary) retrieveImportSymbols() error {
	for _, _fffb := range _bbfc.Header.RTSegments {
		if _fffb.Type == 0 {
			_fbeg, _aacce := _fffb.GetSegmentData()
			if _aacce != nil {
				return _aacce
			}
			_gbca, _ggbf := _fbeg.(*SymbolDictionary)
			if !_ggbf {
				return _f.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0044\u0061\u0074a\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0053\u0065\u0067m\u0065\u006e\u0074\u003a\u0020%\u0054", _fbeg)
			}
			_fcec, _aacce := _gbca.GetDictionary()
			if _aacce != nil {
				return _f.Errorf("\u0072\u0065\u006c\u0061\u0074\u0065\u0064 \u0073\u0065\u0067m\u0065\u006e\u0074 \u0077\u0069t\u0068\u0020\u0069\u006e\u0064\u0065x\u003a %\u0064\u0020\u0067\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0073", _fffb.SegmentNumber, _aacce.Error())
			}
			_bbfc._dcgb = append(_bbfc._dcgb, _fcec...)
			_bbfc._afag += _gbca.NumberOfExportedSymbols
		}
	}
	return nil
}

func (_gbac *GenericRegion) overrideAtTemplate0b(_gcb, _geff, _eeef, _cffa, _dec, _cgbc int) int {
	if _gbac.GBAtOverride[0] {
		_gcb &= 0xFFFD
		if _gbac.GBAtY[0] == 0 && _gbac.GBAtX[0] >= -int8(_dec) {
			_gcb |= (_cffa >> uint(int8(_cgbc)-_gbac.GBAtX[0]&0x1)) << 1
		} else {
			_gcb |= int(_gbac.getPixel(_geff+int(_gbac.GBAtX[0]), _eeef+int(_gbac.GBAtY[0]))) << 1
		}
	}
	if _gbac.GBAtOverride[1] {
		_gcb &= 0xDFFF
		if _gbac.GBAtY[1] == 0 && _gbac.GBAtX[1] >= -int8(_dec) {
			_gcb |= (_cffa >> uint(int8(_cgbc)-_gbac.GBAtX[1]&0x1)) << 13
		} else {
			_gcb |= int(_gbac.getPixel(_geff+int(_gbac.GBAtX[1]), _eeef+int(_gbac.GBAtY[1]))) << 13
		}
	}
	if _gbac.GBAtOverride[2] {
		_gcb &= 0xFDFF
		if _gbac.GBAtY[2] == 0 && _gbac.GBAtX[2] >= -int8(_dec) {
			_gcb |= (_cffa >> uint(int8(_cgbc)-_gbac.GBAtX[2]&0x1)) << 9
		} else {
			_gcb |= int(_gbac.getPixel(_geff+int(_gbac.GBAtX[2]), _eeef+int(_gbac.GBAtY[2]))) << 9
		}
	}
	if _gbac.GBAtOverride[3] {
		_gcb &= 0xBFFF
		if _gbac.GBAtY[3] == 0 && _gbac.GBAtX[3] >= -int8(_dec) {
			_gcb |= (_cffa >> uint(int8(_cgbc)-_gbac.GBAtX[3]&0x1)) << 14
		} else {
			_gcb |= int(_gbac.getPixel(_geff+int(_gbac.GBAtX[3]), _eeef+int(_gbac.GBAtY[3]))) << 14
		}
	}
	if _gbac.GBAtOverride[4] {
		_gcb &= 0xEFFF
		if _gbac.GBAtY[4] == 0 && _gbac.GBAtX[4] >= -int8(_dec) {
			_gcb |= (_cffa >> uint(int8(_cgbc)-_gbac.GBAtX[4]&0x1)) << 12
		} else {
			_gcb |= int(_gbac.getPixel(_geff+int(_gbac.GBAtX[4]), _eeef+int(_gbac.GBAtY[4]))) << 12
		}
	}
	if _gbac.GBAtOverride[5] {
		_gcb &= 0xFFDF
		if _gbac.GBAtY[5] == 0 && _gbac.GBAtX[5] >= -int8(_dec) {
			_gcb |= (_cffa >> uint(int8(_cgbc)-_gbac.GBAtX[5]&0x1)) << 5
		} else {
			_gcb |= int(_gbac.getPixel(_geff+int(_gbac.GBAtX[5]), _eeef+int(_gbac.GBAtY[5]))) << 5
		}
	}
	if _gbac.GBAtOverride[6] {
		_gcb &= 0xFFFB
		if _gbac.GBAtY[6] == 0 && _gbac.GBAtX[6] >= -int8(_dec) {
			_gcb |= (_cffa >> uint(int8(_cgbc)-_gbac.GBAtX[6]&0x1)) << 2
		} else {
			_gcb |= int(_gbac.getPixel(_geff+int(_gbac.GBAtX[6]), _eeef+int(_gbac.GBAtY[6]))) << 2
		}
	}
	if _gbac.GBAtOverride[7] {
		_gcb &= 0xFFF7
		if _gbac.GBAtY[7] == 0 && _gbac.GBAtX[7] >= -int8(_dec) {
			_gcb |= (_cffa >> uint(int8(_cgbc)-_gbac.GBAtX[7]&0x1)) << 3
		} else {
			_gcb |= int(_gbac.getPixel(_geff+int(_gbac.GBAtX[7]), _eeef+int(_gbac.GBAtY[7]))) << 3
		}
	}
	if _gbac.GBAtOverride[8] {
		_gcb &= 0xF7FF
		if _gbac.GBAtY[8] == 0 && _gbac.GBAtX[8] >= -int8(_dec) {
			_gcb |= (_cffa >> uint(int8(_cgbc)-_gbac.GBAtX[8]&0x1)) << 11
		} else {
			_gcb |= int(_gbac.getPixel(_geff+int(_gbac.GBAtX[8]), _eeef+int(_gbac.GBAtY[8]))) << 11
		}
	}
	if _gbac.GBAtOverride[9] {
		_gcb &= 0xFFEF
		if _gbac.GBAtY[9] == 0 && _gbac.GBAtX[9] >= -int8(_dec) {
			_gcb |= (_cffa >> uint(int8(_cgbc)-_gbac.GBAtX[9]&0x1)) << 4
		} else {
			_gcb |= int(_gbac.getPixel(_geff+int(_gbac.GBAtX[9]), _eeef+int(_gbac.GBAtY[9]))) << 4
		}
	}
	if _gbac.GBAtOverride[10] {
		_gcb &= 0x7FFF
		if _gbac.GBAtY[10] == 0 && _gbac.GBAtX[10] >= -int8(_dec) {
			_gcb |= (_cffa >> uint(int8(_cgbc)-_gbac.GBAtX[10]&0x1)) << 15
		} else {
			_gcb |= int(_gbac.getPixel(_geff+int(_gbac.GBAtX[10]), _eeef+int(_gbac.GBAtY[10]))) << 15
		}
	}
	if _gbac.GBAtOverride[11] {
		_gcb &= 0xFDFF
		if _gbac.GBAtY[11] == 0 && _gbac.GBAtX[11] >= -int8(_dec) {
			_gcb |= (_cffa >> uint(int8(_cgbc)-_gbac.GBAtX[11]&0x1)) << 10
		} else {
			_gcb |= int(_gbac.getPixel(_geff+int(_gbac.GBAtX[11]), _eeef+int(_gbac.GBAtY[11]))) << 10
		}
	}
	return _gcb
}

func (_dab *GenericRegion) overrideAtTemplate2(_bacb, _bdg, _fcba, _aade, _ddd int) int {
	_bacb &= 0x3FB
	if _dab.GBAtY[0] == 0 && _dab.GBAtX[0] >= -int8(_ddd) {
		_bacb |= (_aade >> uint(7-(int8(_ddd)+_dab.GBAtX[0])) & 0x1) << 2
	} else {
		_bacb |= int(_dab.getPixel(_bdg+int(_dab.GBAtX[0]), _fcba+int(_dab.GBAtY[0]))) << 2
	}
	return _bacb
}

func (_eee *GenericRegion) overrideAtTemplate0a(_acge, _cca, _fdbf, _aaca, _agad, _dcbf int) int {
	if _eee.GBAtOverride[0] {
		_acge &= 0xFFEF
		if _eee.GBAtY[0] == 0 && _eee.GBAtX[0] >= -int8(_agad) {
			_acge |= (_aaca >> uint(int8(_dcbf)-_eee.GBAtX[0]&0x1)) << 4
		} else {
			_acge |= int(_eee.getPixel(_cca+int(_eee.GBAtX[0]), _fdbf+int(_eee.GBAtY[0]))) << 4
		}
	}
	if _eee.GBAtOverride[1] {
		_acge &= 0xFBFF
		if _eee.GBAtY[1] == 0 && _eee.GBAtX[1] >= -int8(_agad) {
			_acge |= (_aaca >> uint(int8(_dcbf)-_eee.GBAtX[1]&0x1)) << 10
		} else {
			_acge |= int(_eee.getPixel(_cca+int(_eee.GBAtX[1]), _fdbf+int(_eee.GBAtY[1]))) << 10
		}
	}
	if _eee.GBAtOverride[2] {
		_acge &= 0xF7FF
		if _eee.GBAtY[2] == 0 && _eee.GBAtX[2] >= -int8(_agad) {
			_acge |= (_aaca >> uint(int8(_dcbf)-_eee.GBAtX[2]&0x1)) << 11
		} else {
			_acge |= int(_eee.getPixel(_cca+int(_eee.GBAtX[2]), _fdbf+int(_eee.GBAtY[2]))) << 11
		}
	}
	if _eee.GBAtOverride[3] {
		_acge &= 0x7FFF
		if _eee.GBAtY[3] == 0 && _eee.GBAtX[3] >= -int8(_agad) {
			_acge |= (_aaca >> uint(int8(_dcbf)-_eee.GBAtX[3]&0x1)) << 15
		} else {
			_acge |= int(_eee.getPixel(_cca+int(_eee.GBAtX[3]), _fdbf+int(_eee.GBAtY[3]))) << 15
		}
	}
	return _acge
}

func (_fgfd *TextRegion) decodeDfs() (int64, error) {
	if _fgfd.IsHuffmanEncoded {
		if _fgfd.SbHuffFS == 3 {
			if _fgfd._bcbc == nil {
				var _fcecd error
				_fgfd._bcbc, _fcecd = _fgfd.getUserTable(0)
				if _fcecd != nil {
					return 0, _fcecd
				}
			}
			return _fgfd._bcbc.Decode(_fgfd._gdbbd)
		}
		_bbbe, _edec := _cg.GetStandardTable(6 + int(_fgfd.SbHuffFS))
		if _edec != nil {
			return 0, _edec
		}
		return _bbbe.Decode(_fgfd._gdbbd)
	}
	_gddb, _gced := _fgfd._bggc.DecodeInt(_fgfd._ccfc)
	if _gced != nil {
		return 0, _gced
	}
	return int64(_gddb), nil
}

func (_efa *GenericRefinementRegion) updateOverride() error {
	if _efa.GrAtX == nil || _efa.GrAtY == nil {
		return _g.New("\u0041\u0054\u0020\u0070\u0069\u0078\u0065\u006c\u0073\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
	}
	if len(_efa.GrAtX) != len(_efa.GrAtY) {
		return _g.New("A\u0054\u0020\u0070\u0069xe\u006c \u0069\u006e\u0063\u006f\u006es\u0069\u0073\u0074\u0065\u006e\u0074")
	}
	_efa._ge = make([]bool, len(_efa.GrAtX))
	switch _efa.TemplateID {
	case 0:
		if _efa.GrAtX[0] != -1 && _efa.GrAtY[0] != -1 {
			_efa._ge[0] = true
			_efa._bba = true
		}
		if _efa.GrAtX[1] != -1 && _efa.GrAtY[1] != -1 {
			_efa._ge[1] = true
			_efa._bba = true
		}
	case 1:
		_efa._bba = false
	}
	return nil
}

func (_fecd *PageInformationSegment) Encode(w _e.BinaryWriter) (_dcbca int, _ceba error) {
	const _adef = "\u0050\u0061g\u0065\u0049\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u002e\u0045\u006eco\u0064\u0065"
	_agdf := make([]byte, 4)
	_ab.BigEndian.PutUint32(_agdf, uint32(_fecd.PageBMWidth))
	_dcbca, _ceba = w.Write(_agdf)
	if _ceba != nil {
		return _dcbca, _da.Wrap(_ceba, _adef, "\u0077\u0069\u0064t\u0068")
	}
	_ab.BigEndian.PutUint32(_agdf, uint32(_fecd.PageBMHeight))
	var _cadbf int
	_cadbf, _ceba = w.Write(_agdf)
	if _ceba != nil {
		return _cadbf + _dcbca, _da.Wrap(_ceba, _adef, "\u0068\u0065\u0069\u0067\u0068\u0074")
	}
	_dcbca += _cadbf
	_ab.BigEndian.PutUint32(_agdf, uint32(_fecd.ResolutionX))
	_cadbf, _ceba = w.Write(_agdf)
	if _ceba != nil {
		return _cadbf + _dcbca, _da.Wrap(_ceba, _adef, "\u0078\u0020\u0072e\u0073\u006f\u006c\u0075\u0074\u0069\u006f\u006e")
	}
	_dcbca += _cadbf
	_ab.BigEndian.PutUint32(_agdf, uint32(_fecd.ResolutionY))
	if _cadbf, _ceba = w.Write(_agdf); _ceba != nil {
		return _cadbf + _dcbca, _da.Wrap(_ceba, _adef, "\u0079\u0020\u0072e\u0073\u006f\u006c\u0075\u0074\u0069\u006f\u006e")
	}
	_dcbca += _cadbf
	if _ceba = _fecd.encodeFlags(w); _ceba != nil {
		return _dcbca, _da.Wrap(_ceba, _adef, "")
	}
	_dcbca++
	if _cadbf, _ceba = _fecd.encodeStripingInformation(w); _ceba != nil {
		return _dcbca, _da.Wrap(_ceba, _adef, "")
	}
	_dcbca += _cadbf
	return _dcbca, nil
}

func (_gab *RegionSegment) String() string {
	_aged := &_cf.Builder{}
	_aged.WriteString("\u0009[\u0052E\u0047\u0049\u004f\u004e\u0020S\u0045\u0047M\u0045\u004e\u0054\u005d\u000a")
	_aged.WriteString(_f.Sprintf("\t\u0009\u002d\u0020\u0042\u0069\u0074m\u0061\u0070\u0020\u0028\u0077\u0069d\u0074\u0068\u002c\u0020\u0068\u0065\u0069g\u0068\u0074\u0029\u0020\u005b\u0025\u0064\u0078\u0025\u0064]\u000a", _gab.BitmapWidth, _gab.BitmapHeight))
	_aged.WriteString(_f.Sprintf("\u0009\u0009\u002d\u0020L\u006f\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0028\u0078,\u0079)\u003a\u0020\u005b\u0025\u0064\u002c\u0025d\u005d\u000a", _gab.XLocation, _gab.YLocation))
	_aged.WriteString(_f.Sprintf("\t\u0009\u002d\u0020\u0043\u006f\u006db\u0069\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0073", _gab.CombinaionOperator))
	return _aged.String()
}

func (_dgebg *TextRegion) decodeRdy() (int64, error) {
	const _bcagb = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064y"
	if _dgebg.IsHuffmanEncoded {
		if _dgebg.SbHuffRDY == 3 {
			if _dgebg._acfd == nil {
				var (
					_cfaf int
					_feef error
				)
				if _dgebg.SbHuffFS == 3 {
					_cfaf++
				}
				if _dgebg.SbHuffDS == 3 {
					_cfaf++
				}
				if _dgebg.SbHuffDT == 3 {
					_cfaf++
				}
				if _dgebg.SbHuffRDWidth == 3 {
					_cfaf++
				}
				if _dgebg.SbHuffRDHeight == 3 {
					_cfaf++
				}
				if _dgebg.SbHuffRDX == 3 {
					_cfaf++
				}
				_dgebg._acfd, _feef = _dgebg.getUserTable(_cfaf)
				if _feef != nil {
					return 0, _da.Wrap(_feef, _bcagb, "")
				}
			}
			return _dgebg._acfd.Decode(_dgebg._gdbbd)
		}
		_feda, _dcffg := _cg.GetStandardTable(14 + int(_dgebg.SbHuffRDY))
		if _dcffg != nil {
			return 0, _dcffg
		}
		return _feda.Decode(_dgebg._gdbbd)
	}
	_egbg, _faga := _dgebg._bggc.DecodeInt(_dgebg._feab)
	if _faga != nil {
		return 0, _da.Wrap(_faga, _bcagb, "")
	}
	return int64(_egbg), nil
}

func (_adge *PageInformationSegment) readCombinationOperator() error {
	_cabg, _gaab := _adge._dfdf.ReadBits(2)
	if _gaab != nil {
		return _gaab
	}
	_adge._ccde = _ea.CombinationOperator(int(_cabg))
	return nil
}

func (_bdff *SymbolDictionary) addSymbol(_eede Regioner) error {
	_gagcb, _deac := _eede.GetRegionBitmap()
	if _deac != nil {
		return _deac
	}
	_bdff._abed[_bdff._addfb] = _gagcb
	_bdff._ffbb = append(_bdff._ffbb, _gagcb)
	_eg.Log.Trace("\u005b\u0053YM\u0042\u004f\u004c \u0044\u0049\u0043\u0054ION\u0041RY\u005d\u0020\u0041\u0064\u0064\u0065\u0064 s\u0079\u006d\u0062\u006f\u006c\u003a\u0020%\u0073", _gagcb)
	return nil
}

func (_aagg *TextRegion) decodeRI() (int64, error) {
	if !_aagg.UseRefinement {
		return 0, nil
	}
	if _aagg.IsHuffmanEncoded {
		_edgg, _fbag := _aagg._gdbbd.ReadBit()
		return int64(_edgg), _fbag
	}
	_debg, _baed := _aagg._bggc.DecodeInt(_aagg._eacf)
	return int64(_debg), _baed
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
	Reader                   *_e.Reader
	SegmentData              Segmenter
	RTSNumbers               []int
	RetainBits               []uint8
}

func (_affg *Header) subInputReader() (*_e.Reader, error) {
	_gefg := int(_affg.SegmentDataLength)
	if _affg.SegmentDataLength == _b.MaxInt32 {
		_gefg = -1
	}
	return _affg.Reader.NewPartialReader(int(_affg.SegmentDataStartOffset), _gefg, false)
}

func (_bade *HalftoneRegion) computeX(_bbag, _cbfg int) int {
	return _bade.shiftAndFill(int(_bade.HGridX) + _bbag*int(_bade.HRegionY) + _cbfg*int(_bade.HRegionX))
}

func (_bcgb *SymbolDictionary) setInSyms() error {
	if _bcgb.Header.RTSegments != nil {
		return _bcgb.retrieveImportSymbols()
	}
	_bcgb._dcgb = make([]*_ea.Bitmap, 0)
	return nil
}

func (_fggdc *Header) writeFlags(_befg _e.BinaryWriter) (_dfdb error) {
	const _deec = "\u0048\u0065\u0061\u0064\u0065\u0072\u002e\u0077\u0072\u0069\u0074\u0065F\u006c\u0061\u0067\u0073"
	_ggaa := byte(_fggdc.Type)
	if _dfdb = _befg.WriteByte(_ggaa); _dfdb != nil {
		return _da.Wrap(_dfdb, _deec, "\u0077\u0072\u0069ti\u006e\u0067\u0020\u0073\u0065\u0067\u006d\u0065\u006et\u0020t\u0079p\u0065 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	if !_fggdc.RetainFlag && !_fggdc.PageAssociationFieldSize {
		return nil
	}
	if _dfdb = _befg.SkipBits(-8); _dfdb != nil {
		return _da.Wrap(_dfdb, _deec, "\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0062\u0061\u0063\u006b\u0020\u0074\u0068\u0065\u0020\u0062\u0069\u0074\u0073\u0020\u0066\u0061il\u0065\u0064")
	}
	var _gfa int
	if _fggdc.RetainFlag {
		_gfa = 1
	}
	if _dfdb = _befg.WriteBit(_gfa); _dfdb != nil {
		return _da.Wrap(_dfdb, _deec, "\u0072\u0065\u0074\u0061in\u0020\u0072\u0065\u0074\u0061\u0069\u006e\u0020\u0066\u006c\u0061\u0067\u0073")
	}
	_gfa = 0
	if _fggdc.PageAssociationFieldSize {
		_gfa = 1
	}
	if _dfdb = _befg.WriteBit(_gfa); _dfdb != nil {
		return _da.Wrap(_dfdb, _deec, "p\u0061\u0067\u0065\u0020as\u0073o\u0063\u0069\u0061\u0074\u0069o\u006e\u0020\u0066\u006c\u0061\u0067")
	}
	_befg.FinishByte()
	return nil
}

func (_beefd *TextRegion) decodeCurrentT() (int64, error) {
	if _beefd.SbStrips != 1 {
		if _beefd.IsHuffmanEncoded {
			_efba, _adfe := _beefd._gdbbd.ReadBits(byte(_beefd.LogSBStrips))
			return int64(_efba), _adfe
		}
		_adff, _fcgga := _beefd._bggc.DecodeInt(_beefd._efgab)
		if _fcgga != nil {
			return 0, _fcgga
		}
		return int64(_adff), nil
	}
	return 0, nil
}

func (_gaec *PageInformationSegment) readDefaultPixelValue() error {
	_gcbc, _cagg := _gaec._dfdf.ReadBit()
	if _cagg != nil {
		return _cagg
	}
	_gaec.DefaultPixelValue = uint8(_gcbc & 0xf)
	return nil
}

type TextRegion struct {
	_gdbbd                  *_e.Reader
	RegionInfo              *RegionSegment
	SbrTemplate             int8
	SbDsOffset              int8
	DefaultPixel            int8
	CombinationOperator     _ea.CombinationOperator
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
	_afcg                   int64
	SbStrips                int8
	NumberOfSymbols         uint32
	RegionBitmap            *_ea.Bitmap
	Symbols                 []*_ea.Bitmap
	_bggc                   *_ag.Decoder
	_eca                    *GenericRefinementRegion
	_ffae                   *_ag.DecoderStats
	_ccfc                   *_ag.DecoderStats
	_fbfg                   *_ag.DecoderStats
	_efgab                  *_ag.DecoderStats
	_eacf                   *_ag.DecoderStats
	_bgcbb                  *_ag.DecoderStats
	_dcgcb                  *_ag.DecoderStats
	_cgfd                   *_ag.DecoderStats
	_dcfba                  *_ag.DecoderStats
	_feab                   *_ag.DecoderStats
	_ecaa                   *_ag.DecoderStats
	_ggggf                  int8
	_adcf                   *_cg.FixedSizeTable
	Header                  *Header
	_bcbc                   _cg.Tabler
	_begc                   _cg.Tabler
	_accc                   _cg.Tabler
	_acccd                  _cg.Tabler
	_dgfeg                  _cg.Tabler
	_gcgaga                 _cg.Tabler
	_acfd                   _cg.Tabler
	_ecdf                   _cg.Tabler
	_dffgd, _edbd           map[int]int
	_bfec                   []int
	_ecfg                   *_ea.Points
	_bbagb                  *_ea.Bitmaps
	_afegf                  *_dc.IntSlice
	_dbec, _beffc           int
	_fagbc                  *_ea.Boxes
}
type Regioner interface {
	GetRegionBitmap() (*_ea.Bitmap, error)
	GetRegionInfo() *RegionSegment
}

func (_ecce *TableSegment) HtOOB() int32             { return _ecce._caff }
func (_gffb *TableSegment) StreamReader() *_e.Reader { return _gffb._agda }
func (_cecf *SymbolDictionary) Encode(w _e.BinaryWriter) (_fgd int, _abcfb error) {
	const _fbae = "\u0053\u0079\u006dbo\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	if _cecf == nil {
		return 0, _da.Error(_fbae, "\u0073\u0079m\u0062\u006f\u006c\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	if _fgd, _abcfb = _cecf.encodeFlags(w); _abcfb != nil {
		return _fgd, _da.Wrap(_abcfb, _fbae, "")
	}
	_eeea, _abcfb := _cecf.encodeATFlags(w)
	if _abcfb != nil {
		return _fgd, _da.Wrap(_abcfb, _fbae, "")
	}
	_fgd += _eeea
	if _eeea, _abcfb = _cecf.encodeRefinementATFlags(w); _abcfb != nil {
		return _fgd, _da.Wrap(_abcfb, _fbae, "")
	}
	_fgd += _eeea
	if _eeea, _abcfb = _cecf.encodeNumSyms(w); _abcfb != nil {
		return _fgd, _da.Wrap(_abcfb, _fbae, "")
	}
	_fgd += _eeea
	if _eeea, _abcfb = _cecf.encodeSymbols(w); _abcfb != nil {
		return _fgd, _da.Wrap(_abcfb, _fbae, "")
	}
	_fgd += _eeea
	return _fgd, nil
}
func (_eeaed *RegionSegment) Size() int { return 17 }
func (_bcag *TableSegment) Init(h *Header, r *_e.Reader) error {
	_bcag._agda = r
	return _bcag.parseHeader()
}

func (_aadg *HalftoneRegion) combineGrayscalePlanes(_dcc []*_ea.Bitmap, _faf int) error {
	_ffa := 0
	for _gdd := 0; _gdd < _dcc[_faf].Height; _gdd++ {
		for _defc := 0; _defc < _dcc[_faf].Width; _defc += 8 {
			_dee, _cdc := _dcc[_faf+1].GetByte(_ffa)
			if _cdc != nil {
				return _cdc
			}
			_dfa, _cdc := _dcc[_faf].GetByte(_ffa)
			if _cdc != nil {
				return _cdc
			}
			_cdc = _dcc[_faf].SetByte(_ffa, _ea.CombineBytes(_dfa, _dee, _ea.CmbOpXor))
			if _cdc != nil {
				return _cdc
			}
			_ffa++
		}
	}
	return nil
}

func (_bace *TableSegment) parseHeader() error {
	var (
		_ggdf  int
		_efbgc uint64
		_cbeb  error
	)
	_ggdf, _cbeb = _bace._agda.ReadBit()
	if _cbeb != nil {
		return _cbeb
	}
	if _ggdf == 1 {
		return _f.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0061\u0062\u006c\u0065 \u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0064e\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u002e\u0020\u0042\u002e\u0032\u002e1\u0020\u0043\u006f\u0064\u0065\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0066\u006c\u0061\u0067\u0073\u003a\u0020\u0042\u0069\u0074\u0020\u0037\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u007a\u0065\u0072\u006f\u002e\u0020\u0057a\u0073\u003a \u0025\u0064", _ggdf)
	}
	if _efbgc, _cbeb = _bace._agda.ReadBits(3); _cbeb != nil {
		return _cbeb
	}
	_bace._gbda = (int32(_efbgc) + 1) & 0xf
	if _efbgc, _cbeb = _bace._agda.ReadBits(3); _cbeb != nil {
		return _cbeb
	}
	_bace._cggf = (int32(_efbgc) + 1) & 0xf
	if _efbgc, _cbeb = _bace._agda.ReadBits(32); _cbeb != nil {
		return _cbeb
	}
	_bace._aegdd = int32(_efbgc & _b.MaxInt32)
	if _efbgc, _cbeb = _bace._agda.ReadBits(32); _cbeb != nil {
		return _cbeb
	}
	_bace._ccfd = int32(_efbgc & _b.MaxInt32)
	return nil
}

func (_badfg *TextRegion) Encode(w _e.BinaryWriter) (_cbgbb int, _gdada error) {
	const _adec = "\u0054\u0065\u0078\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u002e\u0045n\u0063\u006f\u0064\u0065"
	if _cbgbb, _gdada = _badfg.RegionInfo.Encode(w); _gdada != nil {
		return _cbgbb, _da.Wrap(_gdada, _adec, "")
	}
	var _agfd int
	if _agfd, _gdada = _badfg.encodeFlags(w); _gdada != nil {
		return _cbgbb, _da.Wrap(_gdada, _adec, "")
	}
	_cbgbb += _agfd
	if _agfd, _gdada = _badfg.encodeSymbols(w); _gdada != nil {
		return _cbgbb, _da.Wrap(_gdada, _adec, "")
	}
	_cbgbb += _agfd
	return _cbgbb, nil
}

func (_gdfae *PageInformationSegment) readIsStriped() error {
	_aefa, _egfb := _gdfae._dfdf.ReadBit()
	if _egfb != nil {
		return _egfb
	}
	if _aefa == 1 {
		_gdfae.IsStripe = true
	}
	return nil
}

func (_adg *GenericRefinementRegion) setParameters(_bfb *_ag.DecoderStats, _gce *_ag.Decoder, _baf int8, _gcde, _ffg uint32, _dae *_ea.Bitmap, _ecf, _gaf int32, _bgad bool, _ead []int8, _dg []int8) {
	_eg.Log.Trace("\u005b\u0047\u0045NE\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052E\u0047I\u004fN\u005d \u0073\u0065\u0074\u0050\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	if _bfb != nil {
		_adg._gbf = _bfb
	}
	if _gce != nil {
		_adg._cc = _gce
	}
	_adg.TemplateID = _baf
	_adg.RegionInfo.BitmapWidth = _gcde
	_adg.RegionInfo.BitmapHeight = _ffg
	_adg.ReferenceBitmap = _dae
	_adg.ReferenceDX = _ecf
	_adg.ReferenceDY = _gaf
	_adg.IsTPGROn = _bgad
	_adg.GrAtX = _ead
	_adg.GrAtY = _dg
	_adg.RegionBitmap = nil
	_eg.Log.Trace("[\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052E\u0046\u002d\u0052\u0045\u0047\u0049\u004fN]\u0020\u0073\u0065\u0074P\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073 f\u0069\u006ei\u0073\u0068\u0065\u0064\u002e\u0020\u0025\u0073", _adg)
}

func (_feea *PatternDictionary) readTemplate() error {
	_abeg, _gcdg := _feea._ccca.ReadBits(2)
	if _gcdg != nil {
		return _gcdg
	}
	_feea.HDTemplate = byte(_abeg)
	return nil
}

func (_fdd *SymbolDictionary) encodeFlags(_deadf _e.BinaryWriter) (_ddcdg int, _afaab error) {
	const _gdbb = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _afaab = _deadf.SkipBits(3); _afaab != nil {
		return 0, _da.Wrap(_afaab, _gdbb, "\u0065\u006d\u0070\u0074\u0079\u0020\u0062\u0069\u0074\u0073")
	}
	var _ccec int
	if _fdd.SdrTemplate > 0 {
		_ccec = 1
	}
	if _afaab = _deadf.WriteBit(_ccec); _afaab != nil {
		return _ddcdg, _da.Wrap(_afaab, _gdbb, "s\u0064\u0072\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_ccec = 0
	if _fdd.SdTemplate > 1 {
		_ccec = 1
	}
	if _afaab = _deadf.WriteBit(_ccec); _afaab != nil {
		return _ddcdg, _da.Wrap(_afaab, _gdbb, "\u0073\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_ccec = 0
	if _fdd.SdTemplate == 1 || _fdd.SdTemplate == 3 {
		_ccec = 1
	}
	if _afaab = _deadf.WriteBit(_ccec); _afaab != nil {
		return _ddcdg, _da.Wrap(_afaab, _gdbb, "\u0073\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_ccec = 0
	if _fdd._cba {
		_ccec = 1
	}
	if _afaab = _deadf.WriteBit(_ccec); _afaab != nil {
		return _ddcdg, _da.Wrap(_afaab, _gdbb, "\u0063\u006f\u0064in\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0072\u0065\u0074\u0061\u0069\u006e\u0065\u0064")
	}
	_ccec = 0
	if _fdd._bdgf {
		_ccec = 1
	}
	if _afaab = _deadf.WriteBit(_ccec); _afaab != nil {
		return _ddcdg, _da.Wrap(_afaab, _gdbb, "\u0063\u006f\u0064\u0069ng\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0075\u0073\u0065\u0064")
	}
	_ccec = 0
	if _fdd.SdHuffAggInstanceSelection {
		_ccec = 1
	}
	if _afaab = _deadf.WriteBit(_ccec); _afaab != nil {
		return _ddcdg, _da.Wrap(_afaab, _gdbb, "\u0073\u0064\u0048\u0075\u0066\u0066\u0041\u0067\u0067\u0049\u006e\u0073\u0074")
	}
	_ccec = int(_fdd.SdHuffBMSizeSelection)
	if _afaab = _deadf.WriteBit(_ccec); _afaab != nil {
		return _ddcdg, _da.Wrap(_afaab, _gdbb, "\u0073\u0064\u0048u\u0066\u0066\u0042\u006d\u0053\u0069\u007a\u0065")
	}
	_ccec = 0
	if _fdd.SdHuffDecodeWidthSelection > 1 {
		_ccec = 1
	}
	if _afaab = _deadf.WriteBit(_ccec); _afaab != nil {
		return _ddcdg, _da.Wrap(_afaab, _gdbb, "s\u0064\u0048\u0075\u0066\u0066\u0057\u0069\u0064\u0074\u0068")
	}
	_ccec = 0
	switch _fdd.SdHuffDecodeWidthSelection {
	case 1, 3:
		_ccec = 1
	}
	if _afaab = _deadf.WriteBit(_ccec); _afaab != nil {
		return _ddcdg, _da.Wrap(_afaab, _gdbb, "s\u0064\u0048\u0075\u0066\u0066\u0057\u0069\u0064\u0074\u0068")
	}
	_ccec = 0
	if _fdd.SdHuffDecodeHeightSelection > 1 {
		_ccec = 1
	}
	if _afaab = _deadf.WriteBit(_ccec); _afaab != nil {
		return _ddcdg, _da.Wrap(_afaab, _gdbb, "\u0073\u0064\u0048u\u0066\u0066\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_ccec = 0
	switch _fdd.SdHuffDecodeHeightSelection {
	case 1, 3:
		_ccec = 1
	}
	if _afaab = _deadf.WriteBit(_ccec); _afaab != nil {
		return _ddcdg, _da.Wrap(_afaab, _gdbb, "\u0073\u0064\u0048u\u0066\u0066\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_ccec = 0
	if _fdd.UseRefinementAggregation {
		_ccec = 1
	}
	if _afaab = _deadf.WriteBit(_ccec); _afaab != nil {
		return _ddcdg, _da.Wrap(_afaab, _gdbb, "\u0073\u0064\u0052\u0065\u0066\u0041\u0067\u0067")
	}
	_ccec = 0
	if _fdd.IsHuffmanEncoded {
		_ccec = 1
	}
	if _afaab = _deadf.WriteBit(_ccec); _afaab != nil {
		return _ddcdg, _da.Wrap(_afaab, _gdbb, "\u0073\u0064\u0048\u0075\u0066\u0066")
	}
	return 2, nil
}

type OrganizationType uint8

func (_gfca *TextRegion) GetRegionBitmap() (*_ea.Bitmap, error) {
	if _gfca.RegionBitmap != nil {
		return _gfca.RegionBitmap, nil
	}
	if !_gfca.IsHuffmanEncoded {
		if _ffbe := _gfca.setCodingStatistics(); _ffbe != nil {
			return nil, _ffbe
		}
	}
	if _dedg := _gfca.createRegionBitmap(); _dedg != nil {
		return nil, _dedg
	}
	if _dgag := _gfca.decodeSymbolInstances(); _dgag != nil {
		return nil, _dgag
	}
	return _gfca.RegionBitmap, nil
}

func (_ggga *PageInformationSegment) readResolution() error {
	_beef, _degg := _ggga._dfdf.ReadBits(32)
	if _degg != nil {
		return _degg
	}
	_ggga.ResolutionX = int(_beef & _b.MaxInt32)
	_beef, _degg = _ggga._dfdf.ReadBits(32)
	if _degg != nil {
		return _degg
	}
	_ggga.ResolutionY = int(_beef & _b.MaxInt32)
	return nil
}

func (_aabb *SymbolDictionary) setSymbolsArray() error {
	if _aabb._dcgb == nil {
		if _bgbf := _aabb.retrieveImportSymbols(); _bgbf != nil {
			return _bgbf
		}
	}
	if _aabb._ffbb == nil {
		_aabb._ffbb = append(_aabb._ffbb, _aabb._dcgb...)
	}
	return nil
}

func (_faff *Header) readSegmentPageAssociation(_fgffd Documenter, _dgaf *_e.Reader, _bff uint64, _dcbc ...int) (_ffca error) {
	const _fegb = "\u0072\u0065\u0061\u0064\u0053\u0065\u0067\u006d\u0065\u006e\u0074P\u0061\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061t\u0069\u006f\u006e"
	if !_faff.PageAssociationFieldSize {
		_bcf, _efeg := _dgaf.ReadBits(8)
		if _efeg != nil {
			return _da.Wrap(_efeg, _fegb, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		_faff.PageAssociation = int(_bcf & 0xFF)
	} else {
		_eaad, _gfed := _dgaf.ReadBits(32)
		if _gfed != nil {
			return _da.Wrap(_gfed, _fegb, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		_faff.PageAssociation = int(_eaad & _b.MaxInt32)
	}
	if _bff == 0 {
		return nil
	}
	if _faff.PageAssociation != 0 {
		_fgc, _dceb := _fgffd.GetPage(_faff.PageAssociation)
		if _dceb != nil {
			return _da.Wrap(_dceb, _fegb, "\u0061s\u0073\u006f\u0063\u0069a\u0074\u0065\u0064\u0020\u0070a\u0067e\u0020n\u006f\u0074\u0020\u0066\u006f\u0075\u006ed")
		}
		var _dffa int
		for _acb := uint64(0); _acb < _bff; _acb++ {
			_dffa = _dcbc[_acb]
			_faff.RTSegments[_acb], _dceb = _fgc.GetSegment(_dffa)
			if _dceb != nil {
				var _dedf error
				_faff.RTSegments[_acb], _dedf = _fgffd.GetGlobalSegment(_dffa)
				if _dedf != nil {
					return _da.Wrapf(_dceb, _fegb, "\u0072\u0065\u0066\u0065\u0072\u0065n\u0063\u0065\u0020s\u0065\u0067\u006de\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0020\u0061\u0074\u0020pa\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0072\u0020\u0069\u006e\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0073", _faff.PageAssociation)
				}
			}
		}
		return nil
	}
	for _aecbg := uint64(0); _aecbg < _bff; _aecbg++ {
		_faff.RTSegments[_aecbg], _ffca = _fgffd.GetGlobalSegment(_dcbc[_aecbg])
		if _ffca != nil {
			return _da.Wrapf(_ffca, _fegb, "\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u003a\u0020\u0027\u0025d\u0027\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", _dcbc[_aecbg])
		}
	}
	return nil
}

func (_dgcc *TextRegion) readRegionFlags() error {
	var (
		_bfdb  int
		_fbcee uint64
		_adaeb error
	)
	_bfdb, _adaeb = _dgcc._gdbbd.ReadBit()
	if _adaeb != nil {
		return _adaeb
	}
	_dgcc.SbrTemplate = int8(_bfdb)
	_fbcee, _adaeb = _dgcc._gdbbd.ReadBits(5)
	if _adaeb != nil {
		return _adaeb
	}
	_dgcc.SbDsOffset = int8(_fbcee)
	if _dgcc.SbDsOffset > 0x0f {
		_dgcc.SbDsOffset -= 0x20
	}
	_bfdb, _adaeb = _dgcc._gdbbd.ReadBit()
	if _adaeb != nil {
		return _adaeb
	}
	_dgcc.DefaultPixel = int8(_bfdb)
	_fbcee, _adaeb = _dgcc._gdbbd.ReadBits(2)
	if _adaeb != nil {
		return _adaeb
	}
	_dgcc.CombinationOperator = _ea.CombinationOperator(int(_fbcee) & 0x3)
	_bfdb, _adaeb = _dgcc._gdbbd.ReadBit()
	if _adaeb != nil {
		return _adaeb
	}
	_dgcc.IsTransposed = int8(_bfdb)
	_fbcee, _adaeb = _dgcc._gdbbd.ReadBits(2)
	if _adaeb != nil {
		return _adaeb
	}
	_dgcc.ReferenceCorner = int16(_fbcee) & 0x3
	_fbcee, _adaeb = _dgcc._gdbbd.ReadBits(2)
	if _adaeb != nil {
		return _adaeb
	}
	_dgcc.LogSBStrips = int16(_fbcee) & 0x3
	_dgcc.SbStrips = 1 << uint(_dgcc.LogSBStrips)
	_bfdb, _adaeb = _dgcc._gdbbd.ReadBit()
	if _adaeb != nil {
		return _adaeb
	}
	if _bfdb == 1 {
		_dgcc.UseRefinement = true
	}
	_bfdb, _adaeb = _dgcc._gdbbd.ReadBit()
	if _adaeb != nil {
		return _adaeb
	}
	if _bfdb == 1 {
		_dgcc.IsHuffmanEncoded = true
	}
	return nil
}

func (_bdaa *PatternDictionary) Init(h *Header, r *_e.Reader) error {
	_bdaa._ccca = r
	return _bdaa.parseHeader()
}

func _dffg(_gdeg int) int {
	if _gdeg == 0 {
		return 0
	}
	_gdeg |= _gdeg >> 1
	_gdeg |= _gdeg >> 2
	_gdeg |= _gdeg >> 4
	_gdeg |= _gdeg >> 8
	_gdeg |= _gdeg >> 16
	return (_gdeg + 1) >> 1
}

func (_fca *GenericRegion) writeGBAtPixels(_ecbd _e.BinaryWriter) (_adae int, _abea error) {
	const _bafd = "\u0077r\u0069t\u0065\u0047\u0042\u0041\u0074\u0050\u0069\u0078\u0065\u006c\u0073"
	if _fca.UseMMR {
		return 0, nil
	}
	_ecda := 1
	if _fca.GBTemplate == 0 {
		_ecda = 4
	} else if _fca.UseExtTemplates {
		_ecda = 12
	}
	if len(_fca.GBAtX) != _ecda {
		return 0, _da.Errorf(_bafd, "\u0067\u0062\u0020\u0061\u0074\u0020\u0070\u0061\u0069\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020d\u006f\u0065\u0073\u006e\u0027\u0074\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u006f\u0020\u0047\u0042\u0041t\u0058\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u006c\u0065\u006e")
	}
	if len(_fca.GBAtY) != _ecda {
		return 0, _da.Errorf(_bafd, "\u0067\u0062\u0020\u0061\u0074\u0020\u0070\u0061\u0069\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020d\u006f\u0065\u0073\u006e\u0027\u0074\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u006f\u0020\u0047\u0042\u0041t\u0059\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u006c\u0065\u006e")
	}
	for _gec := 0; _gec < _ecda; _gec++ {
		if _abea = _ecbd.WriteByte(byte(_fca.GBAtX[_gec])); _abea != nil {
			return _adae, _da.Wrap(_abea, _bafd, "w\u0072\u0069\u0074\u0065\u0020\u0047\u0042\u0041\u0074\u0058")
		}
		_adae++
		if _abea = _ecbd.WriteByte(byte(_fca.GBAtY[_gec])); _abea != nil {
			return _adae, _da.Wrap(_abea, _bafd, "w\u0072\u0069\u0074\u0065\u0020\u0047\u0042\u0041\u0074\u0059")
		}
		_adae++
	}
	return _adae, nil
}

func (_bgfd *SymbolDictionary) parseHeader() (_cef error) {
	_eg.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c \u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		if _cef != nil {
			_eg.Log.Trace("\u005bS\u0059\u004dB\u004f\u004c\u0020\u0044I\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u005b\u0050\u0041\u0052SE\u002d\u0048\u0045A\u0044\u0045R\u005d\u0020\u0066\u0061\u0069\u006ce\u0064\u002e \u0025\u0076", _cef)
		} else {
			_eg.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c \u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		}
	}()
	if _cef = _bgfd.readRegionFlags(); _cef != nil {
		return _cef
	}
	if _cef = _bgfd.setAtPixels(); _cef != nil {
		return _cef
	}
	if _cef = _bgfd.setRefinementAtPixels(); _cef != nil {
		return _cef
	}
	if _cef = _bgfd.readNumberOfExportedSymbols(); _cef != nil {
		return _cef
	}
	if _cef = _bgfd.readNumberOfNewSymbols(); _cef != nil {
		return _cef
	}
	if _cef = _bgfd.setInSyms(); _cef != nil {
		return _cef
	}
	if _bgfd._bdgf {
		_fcbac := _bgfd.Header.RTSegments
		for _gdbd := len(_fcbac) - 1; _gdbd >= 0; _gdbd-- {
			if _fcbac[_gdbd].Type == 0 {
				_beff, _afad := _fcbac[_gdbd].SegmentData.(*SymbolDictionary)
				if !_afad {
					_cef = _f.Errorf("\u0072\u0065\u006c\u0061\u0074\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074:\u0020\u0025\u0076\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020S\u0079\u006d\u0062\u006f\u006c\u0020\u0044\u0069\u0063\u0074\u0069\u006fna\u0072\u0079\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074", _fcbac[_gdbd])
					return _cef
				}
				if _beff._bdgf {
					_bgfd.setRetainedCodingContexts(_beff)
				}
				break
			}
		}
	}
	if _cef = _bgfd.checkInput(); _cef != nil {
		return _cef
	}
	return nil
}

func (_bgcd *GenericRegion) InitEncode(bm *_ea.Bitmap, xLoc, yLoc, template int, duplicateLineRemoval bool) error {
	const _fcgg = "\u0047e\u006e\u0065\u0072\u0069\u0063\u0052\u0065\u0067\u0069\u006f\u006e.\u0049\u006e\u0069\u0074\u0045\u006e\u0063\u006f\u0064\u0065"
	if bm == nil {
		return _da.Error(_fcgg, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if xLoc < 0 || yLoc < 0 {
		return _da.Error(_fcgg, "\u0078\u0020\u0061\u006e\u0064\u0020\u0079\u0020\u006c\u006f\u0063\u0061\u0074i\u006f\u006e\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	_bgcd.Bitmap = bm
	_bgcd.GBTemplate = byte(template)
	switch _bgcd.GBTemplate {
	case 0:
		_bgcd.GBAtX = []int8{3, -3, 2, -2}
		_bgcd.GBAtY = []int8{-1, -1, -2, -2}
	case 1:
		_bgcd.GBAtX = []int8{3}
		_bgcd.GBAtY = []int8{-1}
	case 2, 3:
		_bgcd.GBAtX = []int8{2}
		_bgcd.GBAtY = []int8{-1}
	default:
		return _da.Errorf(_fcgg, "\u0070\u0072o\u0076\u0069\u0064\u0065\u0064 \u0074\u0065\u006d\u0070\u006ca\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0030\u002c\u0031\u002c\u0032\u002c\u0033\u007d", template)
	}
	_bgcd.RegionSegment = &RegionSegment{BitmapHeight: uint32(bm.Height), BitmapWidth: uint32(bm.Width), XLocation: uint32(xLoc), YLocation: uint32(yLoc)}
	_bgcd.IsTPGDon = duplicateLineRemoval
	return nil
}

func (_dabb *SymbolDictionary) InitEncode(symbols *_ea.Bitmaps, symbolList []int, symbolMap map[int]int, unborderSymbols bool) error {
	const _fedg = "S\u0079\u006d\u0062\u006f\u006c\u0044i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002eI\u006e\u0069\u0074E\u006ec\u006f\u0064\u0065"
	_dabb.SdATX = []int8{3, -3, 2, -2}
	_dabb.SdATY = []int8{-1, -1, -2, -2}
	_dabb._efbc = symbols
	_dabb._decd = make([]int, len(symbolList))
	copy(_dabb._decd, symbolList)
	if len(_dabb._decd) != _dabb._efbc.Size() {
		return _da.Error(_fedg, "s\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0061\u006e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u004ci\u0073\u0074\u0020\u006f\u0066\u0020\u0064\u0069\u0066\u0066er\u0065\u006e\u0074 \u0073i\u007a\u0065")
	}
	_dabb.NumberOfNewSymbols = uint32(symbols.Size())
	_dabb.NumberOfExportedSymbols = uint32(symbols.Size())
	_dabb._fcae = symbolMap
	_dabb._ecffe = unborderSymbols
	return nil
}

func (_facg *SymbolDictionary) GetDictionary() ([]*_ea.Bitmap, error) {
	_eg.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c-\u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		_eg.Log.Trace("\u005b\u0053\u0059M\u0042\u004f\u004c\u002d\u0044\u0049\u0043\u0054\u0049\u004f\u004e\u0041\u0052\u0059\u005d\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
		_eg.Log.Trace("\u005b\u0053Y\u004d\u0042\u004f\u004c\u002dD\u0049\u0043\u0054\u0049\u004fN\u0041\u0052\u0059\u005d\u0020\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0020\u000a\u0045\u0078\u003a\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u000a\u006e\u0065\u0077\u003a\u0027\u0025\u0073\u0027", _facg._cfbdb, _facg._abed)
	}()
	if _facg._cfbdb == nil {
		var _bccd error
		if _facg.UseRefinementAggregation {
			_facg._cbe = _facg.getSbSymCodeLen()
		}
		if !_facg.IsHuffmanEncoded {
			if _bccd = _facg.setCodingStatistics(); _bccd != nil {
				return nil, _bccd
			}
		}
		_facg._abed = make([]*_ea.Bitmap, _facg.NumberOfNewSymbols)
		var _bdcg []int
		if _facg.IsHuffmanEncoded && !_facg.UseRefinementAggregation {
			_bdcg = make([]int, _facg.NumberOfNewSymbols)
		}
		if _bccd = _facg.setSymbolsArray(); _bccd != nil {
			return nil, _bccd
		}
		var _deggd, _agage int64
		_facg._addfb = 0
		for _facg._addfb < _facg.NumberOfNewSymbols {
			_agage, _bccd = _facg.decodeHeightClassDeltaHeight()
			if _bccd != nil {
				return nil, _bccd
			}
			_deggd += _agage
			var _bcga, _dadf uint32
			_caga := int64(_facg._addfb)
			for {
				var _cgcb int64
				_cgcb, _bccd = _facg.decodeDifferenceWidth()
				if _ae.Is(_bccd, _abf.ErrOOB) {
					break
				}
				if _bccd != nil {
					return nil, _bccd
				}
				if _facg._addfb >= _facg.NumberOfNewSymbols {
					break
				}
				_bcga += uint32(_cgcb)
				_dadf += _bcga
				if !_facg.IsHuffmanEncoded || _facg.UseRefinementAggregation {
					if !_facg.UseRefinementAggregation {
						_bccd = _facg.decodeDirectlyThroughGenericRegion(_bcga, uint32(_deggd))
						if _bccd != nil {
							return nil, _bccd
						}
					} else {
						_bccd = _facg.decodeAggregate(_bcga, uint32(_deggd))
						if _bccd != nil {
							return nil, _bccd
						}
					}
				} else if _facg.IsHuffmanEncoded && !_facg.UseRefinementAggregation {
					_bdcg[_facg._addfb] = int(_bcga)
				}
				_facg._addfb++
			}
			if _facg.IsHuffmanEncoded && !_facg.UseRefinementAggregation {
				var _dgg int64
				if _facg.SdHuffBMSizeSelection == 0 {
					var _bcgab _cg.Tabler
					_bcgab, _bccd = _cg.GetStandardTable(1)
					if _bccd != nil {
						return nil, _bccd
					}
					_dgg, _bccd = _bcgab.Decode(_facg._abege)
					if _bccd != nil {
						return nil, _bccd
					}
				} else {
					_dgg, _bccd = _facg.huffDecodeBmSize()
					if _bccd != nil {
						return nil, _bccd
					}
				}
				_facg._abege.Align()
				var _bdda *_ea.Bitmap
				_bdda, _bccd = _facg.decodeHeightClassCollectiveBitmap(_dgg, uint32(_deggd), _dadf)
				if _bccd != nil {
					return nil, _bccd
				}
				_bccd = _facg.decodeHeightClassBitmap(_bdda, _caga, int(_deggd), _bdcg)
				if _bccd != nil {
					return nil, _bccd
				}
			}
		}
		_efbf, _bccd := _facg.getToExportFlags()
		if _bccd != nil {
			return nil, _bccd
		}
		_facg.setExportedSymbols(_efbf)
	}
	return _facg._cfbdb, nil
}

func (_dgfe *Header) referenceSize() uint {
	switch {
	case _dgfe.SegmentNumber <= 255:
		return 1
	case _dgfe.SegmentNumber <= 65535:
		return 2
	default:
		return 4
	}
}

func (_bdgd *SymbolDictionary) encodeATFlags(_fagb _e.BinaryWriter) (_efbg int, _dfgc error) {
	const _acec = "\u0065\u006e\u0063\u006f\u0064\u0065\u0041\u0054\u0046\u006c\u0061\u0067\u0073"
	if _bdgd.IsHuffmanEncoded || _bdgd.SdTemplate != 0 {
		return 0, nil
	}
	_dfaa := 4
	if _bdgd.SdTemplate != 0 {
		_dfaa = 1
	}
	for _dbee := 0; _dbee < _dfaa; _dbee++ {
		if _dfgc = _fagb.WriteByte(byte(_bdgd.SdATX[_dbee])); _dfgc != nil {
			return _efbg, _da.Wrapf(_dfgc, _acec, "\u0053d\u0041\u0054\u0058\u005b\u0025\u0064]", _dbee)
		}
		_efbg++
		if _dfgc = _fagb.WriteByte(byte(_bdgd.SdATY[_dbee])); _dfgc != nil {
			return _efbg, _da.Wrapf(_dfgc, _acec, "\u0053d\u0041\u0054\u0059\u005b\u0025\u0064]", _dbee)
		}
		_efbg++
	}
	return _efbg, nil
}

func (_ccfa *Header) readHeaderFlags() error {
	const _bedd = "\u0072e\u0061d\u0048\u0065\u0061\u0064\u0065\u0072\u0046\u006c\u0061\u0067\u0073"
	_bcbf, _bafb := _ccfa.Reader.ReadBit()
	if _bafb != nil {
		return _da.Wrap(_bafb, _bedd, "r\u0065\u0074\u0061\u0069\u006e\u0020\u0066\u006c\u0061\u0067")
	}
	if _bcbf != 0 {
		_ccfa.RetainFlag = true
	}
	_bcbf, _bafb = _ccfa.Reader.ReadBit()
	if _bafb != nil {
		return _da.Wrap(_bafb, _bedd, "\u0070\u0061g\u0065\u0020\u0061s\u0073\u006f\u0063\u0069\u0061\u0074\u0069\u006f\u006e")
	}
	if _bcbf != 0 {
		_ccfa.PageAssociationFieldSize = true
	}
	_ggaf, _bafb := _ccfa.Reader.ReadBits(6)
	if _bafb != nil {
		return _da.Wrap(_bafb, _bedd, "\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065")
	}
	_ccfa.Type = Type(int(_ggaf))
	return nil
}
func (_cgf *template0) setIndex(_ece *_ag.DecoderStats) { _ece.SetIndex(0x100) }
func (_dedd *GenericRegion) parseHeader() (_cgd error) {
	_eg.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052I\u0043\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0050\u0061\u0072s\u0069\u006e\u0067\u0048\u0065\u0061\u0064e\u0072\u002e\u002e\u002e")
	defer func() {
		if _cgd != nil {
			_eg.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0047\u0049\u004f\u004e]\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065\u0061\u0064\u0065r\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u0020\u0077\u0069th\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020\u0025\u0076", _cgd)
		} else {
			_eg.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049C\u002d\u0052\u0045G\u0049\u004f\u004e]\u0020\u0050a\u0072\u0073\u0069\u006e\u0067\u0048e\u0061de\u0072\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u0020\u0053\u0075\u0063\u0063\u0065\u0073\u0073\u0066\u0075\u006c\u006c\u0079\u002e\u002e\u002e")
		}
	}()
	var (
		_gdc  int
		_gdfe uint64
	)
	if _cgd = _dedd.RegionSegment.parseHeader(); _cgd != nil {
		return _cgd
	}
	if _, _cgd = _dedd._fec.ReadBits(3); _cgd != nil {
		return _cgd
	}
	_gdc, _cgd = _dedd._fec.ReadBit()
	if _cgd != nil {
		return _cgd
	}
	if _gdc == 1 {
		_dedd.UseExtTemplates = true
	}
	_gdc, _cgd = _dedd._fec.ReadBit()
	if _cgd != nil {
		return _cgd
	}
	if _gdc == 1 {
		_dedd.IsTPGDon = true
	}
	_gdfe, _cgd = _dedd._fec.ReadBits(2)
	if _cgd != nil {
		return _cgd
	}
	_dedd.GBTemplate = byte(_gdfe & 0xf)
	_gdc, _cgd = _dedd._fec.ReadBit()
	if _cgd != nil {
		return _cgd
	}
	if _gdc == 1 {
		_dedd.IsMMREncoded = true
	}
	if !_dedd.IsMMREncoded {
		_ggae := 1
		if _dedd.GBTemplate == 0 {
			_ggae = 4
			if _dedd.UseExtTemplates {
				_ggae = 12
			}
		}
		if _cgd = _dedd.readGBAtPixels(_ggae); _cgd != nil {
			return _cgd
		}
	}
	if _cgd = _dedd.computeSegmentDataStructure(); _cgd != nil {
		return _cgd
	}
	_eg.Log.Trace("\u0025\u0073", _dedd)
	return nil
}

func (_begd *PageInformationSegment) readRequiresAuxiliaryBuffer() error {
	_fgcc, _effeg := _begd._dfdf.ReadBit()
	if _effeg != nil {
		return _effeg
	}
	if _fgcc == 1 {
		_begd._dcfb = true
	}
	return nil
}

type Segmenter interface {
	Init(_ggbb *Header, _defcd *_e.Reader) error
}

func (_cacb *PageInformationSegment) CombinationOperator() _ea.CombinationOperator {
	return _cacb._ccde
}

func (_eefc *PatternDictionary) extractPatterns(_cgcf *_ea.Bitmap) error {
	var _cea int
	_adag := make([]*_ea.Bitmap, _eefc.GrayMax+1)
	for _cea <= int(_eefc.GrayMax) {
		_ceg := int(_eefc.HdpWidth) * _cea
		_gdb := _c.Rect(_ceg, 0, _ceg+int(_eefc.HdpWidth), int(_eefc.HdpHeight))
		_gega, _aea := _ea.Extract(_gdb, _cgcf)
		if _aea != nil {
			return _aea
		}
		_adag[_cea] = _gega
		_cea++
	}
	_eefc.Patterns = _adag
	return nil
}

func (_gbg *GenericRefinementRegion) readAtPixels() error {
	_gbg.GrAtX = make([]int8, 2)
	_gbg.GrAtY = make([]int8, 2)
	_ccg, _cgcgd := _gbg._gbc.ReadByte()
	if _cgcgd != nil {
		return _cgcgd
	}
	_gbg.GrAtX[0] = int8(_ccg)
	_ccg, _cgcgd = _gbg._gbc.ReadByte()
	if _cgcgd != nil {
		return _cgcgd
	}
	_gbg.GrAtY[0] = int8(_ccg)
	_ccg, _cgcgd = _gbg._gbc.ReadByte()
	if _cgcgd != nil {
		return _cgcgd
	}
	_gbg.GrAtX[1] = int8(_ccg)
	_ccg, _cgcgd = _gbg._gbc.ReadByte()
	if _cgcgd != nil {
		return _cgcgd
	}
	_gbg.GrAtY[1] = int8(_ccg)
	return nil
}

func (_dgea *GenericRegion) decodeTemplate3(_ace, _eeg, _dfd int, _ecbc, _effe int) (_cab error) {
	const _cbf = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0033"
	var (
		_dcd, _cfda  int
		_dff         int
		_eabe        byte
		_degc, _affc int
	)
	if _ace >= 1 {
		_eabe, _cab = _dgea.Bitmap.GetByte(_effe)
		if _cab != nil {
			return _da.Wrap(_cab, _cbf, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_dff = int(_eabe)
	}
	_dcd = (_dff >> 1) & 0x70
	for _edff := 0; _edff < _dfd; _edff = _degc {
		var (
			_efdb byte
			_afgc int
		)
		_degc = _edff + 8
		if _gfbc := _eeg - _edff; _gfbc > 8 {
			_afgc = 8
		} else {
			_afgc = _gfbc
		}
		if _ace >= 1 {
			_dff <<= 8
			if _degc < _eeg {
				_eabe, _cab = _dgea.Bitmap.GetByte(_effe + 1)
				if _cab != nil {
					return _da.Wrap(_cab, _cbf, "\u0069\u006e\u006e\u0065\u0072\u0020\u002d\u0020\u006c\u0069\u006e\u0065 \u003e\u003d\u0020\u0031")
				}
				_dff |= int(_eabe)
			}
		}
		for _dafg := 0; _dafg < _afgc; _dafg++ {
			if _dgea._aaf {
				_cfda = _dgea.overrideAtTemplate3(_dcd, _edff+_dafg, _ace, int(_efdb), _dafg)
				_dgea._eced.SetIndex(int32(_cfda))
			} else {
				_dgea._eced.SetIndex(int32(_dcd))
			}
			_affc, _cab = _dgea._bda.DecodeBit(_dgea._eced)
			if _cab != nil {
				return _da.Wrap(_cab, _cbf, "")
			}
			_efdb |= byte(_affc) << byte(7-_dafg)
			_dcd = ((_dcd & 0x1f7) << 1) | _affc | ((_dff >> uint(8-_dafg)) & 0x010)
		}
		if _eacc := _dgea.Bitmap.SetByte(_ecbc, _efdb); _eacc != nil {
			return _da.Wrap(_eacc, _cbf, "")
		}
		_ecbc++
		_effe++
	}
	return nil
}

func NewHeader(d Documenter, r *_e.Reader, offset int64, organizationType OrganizationType) (*Header, error) {
	_efbd := &Header{Reader: r}
	if _bdb := _efbd.parse(d, r, offset, organizationType); _bdb != nil {
		return nil, _da.Wrap(_bdb, "\u004ee\u0077\u0048\u0065\u0061\u0064\u0065r", "")
	}
	return _efbd, nil
}

var (
	_eeag Segmenter
	_egd  = map[Type]func() Segmenter{TSymbolDictionary: func() Segmenter { return &SymbolDictionary{} }, TIntermediateTextRegion: func() Segmenter { return &TextRegion{} }, TImmediateTextRegion: func() Segmenter { return &TextRegion{} }, TImmediateLosslessTextRegion: func() Segmenter { return &TextRegion{} }, TPatternDictionary: func() Segmenter { return &PatternDictionary{} }, TIntermediateHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TImmediateHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TImmediateLosslessHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TIntermediateGenericRegion: func() Segmenter { return &GenericRegion{} }, TImmediateGenericRegion: func() Segmenter { return &GenericRegion{} }, TImmediateLosslessGenericRegion: func() Segmenter { return &GenericRegion{} }, TIntermediateGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TImmediateGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TImmediateLosslessGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TPageInformation: func() Segmenter { return &PageInformationSegment{} }, TEndOfPage: func() Segmenter { return _eeag }, TEndOfStrip: func() Segmenter { return &EndOfStripe{} }, TEndOfFile: func() Segmenter { return _eeag }, TProfiles: func() Segmenter { return _eeag }, TTables: func() Segmenter { return &TableSegment{} }, TExtension: func() Segmenter { return _eeag }, TBitmap: func() Segmenter { return _eeag }}
)

func (_dac *Header) readSegmentNumber(_dbg *_e.Reader) error {
	const _fbbc = "\u0072\u0065\u0061\u0064\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u004eu\u006d\u0062\u0065\u0072"
	_gebcc := make([]byte, 4)
	_, _gddg := _dbg.Read(_gebcc)
	if _gddg != nil {
		return _da.Wrap(_gddg, _fbbc, "")
	}
	_dac.SegmentNumber = _ab.BigEndian.Uint32(_gebcc)
	return nil
}

func (_ageb *TextRegion) decodeID() (int64, error) {
	if _ageb.IsHuffmanEncoded {
		if _ageb._adcf == nil {
			_baec, _dcbb := _ageb._gdbbd.ReadBits(byte(_ageb._ggggf))
			return int64(_baec), _dcbb
		}
		return _ageb._adcf.Decode(_ageb._gdbbd)
	}
	return _ageb._bggc.DecodeIAID(uint64(_ageb._ggggf), _ageb._cgfd)
}

type template0 struct{}

func (_eaeg *SymbolDictionary) getSymbol(_feba int) (*_ea.Bitmap, error) {
	const _bdbf = "\u0067e\u0074\u0053\u0079\u006d\u0062\u006fl"
	_gbcd, _ecdaf := _eaeg._efbc.GetBitmap(_eaeg._decd[_feba])
	if _ecdaf != nil {
		return nil, _da.Wrap(_ecdaf, _bdbf, "\u0063\u0061n\u0027\u0074\u0020g\u0065\u0074\u0020\u0073\u0079\u006d\u0062\u006f\u006c")
	}
	return _gbcd, nil
}

var _ SegmentEncoder = &RegionSegment{}

func (_gfde *TextRegion) decodeSymbolInstances() error {
	_fcab, _eddg := _gfde.decodeStripT()
	if _eddg != nil {
		return _eddg
	}
	var (
		_bfd  int64
		_bfbb uint32
	)
	for _bfbb < _gfde.NumberOfSymbolInstances {
		_bgbe, _agef := _gfde.decodeDT()
		if _agef != nil {
			return _agef
		}
		_fcab += _bgbe
		var _daac int64
		_afafg := true
		_gfde._afcg = 0
		for {
			if _afafg {
				_daac, _agef = _gfde.decodeDfs()
				if _agef != nil {
					return _agef
				}
				_bfd += _daac
				_gfde._afcg = _bfd
				_afafg = false
			} else {
				_aagd, _acda := _gfde.decodeIds()
				if _ae.Is(_acda, _abf.ErrOOB) {
					break
				}
				if _acda != nil {
					return _acda
				}
				if _bfbb >= _gfde.NumberOfSymbolInstances {
					break
				}
				_gfde._afcg += _aagd + int64(_gfde.SbDsOffset)
			}
			_cece, _cbdf := _gfde.decodeCurrentT()
			if _cbdf != nil {
				return _cbdf
			}
			_dabf := _fcab + _cece
			_dgbc, _cbdf := _gfde.decodeID()
			if _cbdf != nil {
				return _cbdf
			}
			_gdge, _cbdf := _gfde.decodeRI()
			if _cbdf != nil {
				return _cbdf
			}
			_ccce, _cbdf := _gfde.decodeIb(_gdge, _dgbc)
			if _cbdf != nil {
				return _cbdf
			}
			if _cbdf = _gfde.blit(_ccce, _dabf); _cbdf != nil {
				return _cbdf
			}
			_bfbb++
		}
	}
	return nil
}

func (_egff *SymbolDictionary) huffDecodeBmSize() (int64, error) {
	if _egff._gdfd == nil {
		var (
			_eaac int
			_eece error
		)
		if _egff.SdHuffDecodeHeightSelection == 3 {
			_eaac++
		}
		if _egff.SdHuffDecodeWidthSelection == 3 {
			_eaac++
		}
		_egff._gdfd, _eece = _egff.getUserTable(_eaac)
		if _eece != nil {
			return 0, _eece
		}
	}
	return _egff._gdfd.Decode(_egff._abege)
}

func (_ecba *SymbolDictionary) Init(h *Header, r *_e.Reader) error {
	_ecba.Header = h
	_ecba._abege = r
	return _ecba.parseHeader()
}
func (_gbbe *TableSegment) HtHigh() int32 { return _gbbe._ccfd }
func (_dad *EndOfStripe) LineNumber() int { return _dad._fb }
func (_daec *TextRegion) decodeRdh() (int64, error) {
	const _daga = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064h"
	if _daec.IsHuffmanEncoded {
		if _daec.SbHuffRDHeight == 3 {
			if _daec._dgfeg == nil {
				var (
					_bec   int
					_agbda error
				)
				if _daec.SbHuffFS == 3 {
					_bec++
				}
				if _daec.SbHuffDS == 3 {
					_bec++
				}
				if _daec.SbHuffDT == 3 {
					_bec++
				}
				if _daec.SbHuffRDWidth == 3 {
					_bec++
				}
				_daec._dgfeg, _agbda = _daec.getUserTable(_bec)
				if _agbda != nil {
					return 0, _da.Wrap(_agbda, _daga, "")
				}
			}
			return _daec._dgfeg.Decode(_daec._gdbbd)
		}
		_dcff, _caaa := _cg.GetStandardTable(14 + int(_daec.SbHuffRDHeight))
		if _caaa != nil {
			return 0, _da.Wrap(_caaa, _daga, "")
		}
		return _dcff.Decode(_daec._gdbbd)
	}
	_bcffa, _gffd := _daec._bggc.DecodeInt(_daec._dcgcb)
	if _gffd != nil {
		return 0, _da.Wrap(_gffd, _daga, "")
	}
	return int64(_bcffa), nil
}

func (_agbgg *GenericRegion) getPixel(_bcc, _fbdc int) int8 {
	if _bcc < 0 || _bcc >= _agbgg.Bitmap.Width {
		return 0
	}
	if _fbdc < 0 || _fbdc >= _agbgg.Bitmap.Height {
		return 0
	}
	if _agbgg.Bitmap.GetPixel(_bcc, _fbdc) {
		return 1
	}
	return 0
}

func (_begf *GenericRegion) decodeTemplate2(_dfe, _eabb, _acga int, _fdb, _gda int) (_bbcf error) {
	const _fed = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0032"
	var (
		_fead, _ecb  int
		_ccbb, _abdg int
		_eed         byte
		_ggb, _bgf   int
	)
	if _dfe >= 1 {
		_eed, _bbcf = _begf.Bitmap.GetByte(_gda)
		if _bbcf != nil {
			return _da.Wrap(_bbcf, _fed, "\u006ci\u006ee\u004e\u0075\u006d\u0062\u0065\u0072\u0020\u003e\u003d\u0020\u0031")
		}
		_ccbb = int(_eed)
	}
	if _dfe >= 2 {
		_eed, _bbcf = _begf.Bitmap.GetByte(_gda - _begf.Bitmap.RowStride)
		if _bbcf != nil {
			return _da.Wrap(_bbcf, _fed, "\u006ci\u006ee\u004e\u0075\u006d\u0062\u0065\u0072\u0020\u003e\u003d\u0020\u0032")
		}
		_abdg = int(_eed) << 4
	}
	_fead = (_ccbb >> 3 & 0x7c) | (_abdg >> 3 & 0x380)
	for _agbd := 0; _agbd < _acga; _agbd = _ggb {
		var (
			_ccef byte
			_dde  int
		)
		_ggb = _agbd + 8
		if _ddf := _eabb - _agbd; _ddf > 8 {
			_dde = 8
		} else {
			_dde = _ddf
		}
		if _dfe > 0 {
			_ccbb <<= 8
			if _ggb < _eabb {
				_eed, _bbcf = _begf.Bitmap.GetByte(_gda + 1)
				if _bbcf != nil {
					return _da.Wrap(_bbcf, _fed, "\u006c\u0069\u006e\u0065\u004e\u0075\u006d\u0062\u0065r\u0020\u003e\u0020\u0030")
				}
				_ccbb |= int(_eed)
			}
		}
		if _dfe > 1 {
			_abdg <<= 8
			if _ggb < _eabb {
				_eed, _bbcf = _begf.Bitmap.GetByte(_gda - _begf.Bitmap.RowStride + 1)
				if _bbcf != nil {
					return _da.Wrap(_bbcf, _fed, "\u006c\u0069\u006e\u0065\u004e\u0075\u006d\u0062\u0065r\u0020\u003e\u0020\u0031")
				}
				_abdg |= int(_eed) << 4
			}
		}
		for _ced := 0; _ced < _dde; _ced++ {
			_cggd := uint(10 - _ced)
			if _begf._aaf {
				_ecb = _begf.overrideAtTemplate2(_fead, _agbd+_ced, _dfe, int(_ccef), _ced)
				_begf._eced.SetIndex(int32(_ecb))
			} else {
				_begf._eced.SetIndex(int32(_fead))
			}
			_bgf, _bbcf = _begf._bda.DecodeBit(_begf._eced)
			if _bbcf != nil {
				return _da.Wrap(_bbcf, _fed, "")
			}
			_ccef |= byte(_bgf << uint(7-_ced))
			_fead = ((_fead & 0x1bd) << 1) | _bgf | ((_ccbb >> _cggd) & 0x4) | ((_abdg >> _cggd) & 0x80)
		}
		if _cggb := _begf.Bitmap.SetByte(_fdb, _ccef); _cggb != nil {
			return _da.Wrap(_cggb, _fed, "")
		}
		_fdb++
		_gda++
	}
	return nil
}

func (_degb *RegionSegment) Encode(w _e.BinaryWriter) (_faebg int, _gfad error) {
	const _babb = "R\u0065g\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065n\u0074\u002e\u0045\u006eco\u0064\u0065"
	_ccbe := make([]byte, 4)
	_ab.BigEndian.PutUint32(_ccbe, _degb.BitmapWidth)
	_faebg, _gfad = w.Write(_ccbe)
	if _gfad != nil {
		return 0, _da.Wrap(_gfad, _babb, "\u0057\u0069\u0064t\u0068")
	}
	_ab.BigEndian.PutUint32(_ccbe, _degb.BitmapHeight)
	var _cecb int
	_cecb, _gfad = w.Write(_ccbe)
	if _gfad != nil {
		return 0, _da.Wrap(_gfad, _babb, "\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_faebg += _cecb
	_ab.BigEndian.PutUint32(_ccbe, _degb.XLocation)
	_cecb, _gfad = w.Write(_ccbe)
	if _gfad != nil {
		return 0, _da.Wrap(_gfad, _babb, "\u0058L\u006f\u0063\u0061\u0074\u0069\u006fn")
	}
	_faebg += _cecb
	_ab.BigEndian.PutUint32(_ccbe, _degb.YLocation)
	_cecb, _gfad = w.Write(_ccbe)
	if _gfad != nil {
		return 0, _da.Wrap(_gfad, _babb, "\u0059L\u006f\u0063\u0061\u0074\u0069\u006fn")
	}
	_faebg += _cecb
	if _gfad = w.WriteByte(byte(_degb.CombinaionOperator) & 0x07); _gfad != nil {
		return 0, _da.Wrap(_gfad, _babb, "c\u006fm\u0062\u0069\u006e\u0061\u0074\u0069\u006f\u006e \u006f\u0070\u0065\u0072at\u006f\u0072")
	}
	_faebg++
	return _faebg, nil
}

func (_agbb *PatternDictionary) setGbAtPixels() {
	if _agbb.HDTemplate == 0 {
		_agbb.GBAtX = make([]int8, 4)
		_agbb.GBAtY = make([]int8, 4)
		_agbb.GBAtX[0] = -int8(_agbb.HdpWidth)
		_agbb.GBAtY[0] = 0
		_agbb.GBAtX[1] = -3
		_agbb.GBAtY[1] = -1
		_agbb.GBAtX[2] = 2
		_agbb.GBAtY[2] = -2
		_agbb.GBAtX[3] = -2
		_agbb.GBAtY[3] = -2
	} else {
		_agbb.GBAtX = []int8{-int8(_agbb.HdpWidth)}
		_agbb.GBAtY = []int8{0}
	}
}

func (_eagbf *Header) CleanSegmentData() {
	if _eagbf.SegmentData != nil {
		_eagbf.SegmentData = nil
	}
}

func (_gccf *TextRegion) createRegionBitmap() error {
	_gccf.RegionBitmap = _ea.New(int(_gccf.RegionInfo.BitmapWidth), int(_gccf.RegionInfo.BitmapHeight))
	if _gccf.DefaultPixel != 0 {
		_gccf.RegionBitmap.SetDefaultPixel()
	}
	return nil
}

func (_bbcfbg *PatternDictionary) computeSegmentDataStructure() error {
	_bbcfbg.DataOffset = _bbcfbg._ccca.AbsolutePosition()
	_bbcfbg.DataHeaderLength = _bbcfbg.DataOffset - _bbcfbg.DataHeaderOffset
	_bbcfbg.DataLength = int64(_bbcfbg._ccca.AbsoluteLength()) - _bbcfbg.DataHeaderLength
	return nil
}

func (_agec *HalftoneRegion) parseHeader() error {
	if _cbd := _agec.RegionSegment.parseHeader(); _cbd != nil {
		return _cbd
	}
	_fegd, _fdgg := _agec._bfbe.ReadBit()
	if _fdgg != nil {
		return _fdgg
	}
	_agec.HDefaultPixel = int8(_fegd)
	_gfd, _fdgg := _agec._bfbe.ReadBits(3)
	if _fdgg != nil {
		return _fdgg
	}
	_agec.CombinationOperator = _ea.CombinationOperator(_gfd & 0xf)
	_fegd, _fdgg = _agec._bfbe.ReadBit()
	if _fdgg != nil {
		return _fdgg
	}
	if _fegd == 1 {
		_agec.HSkipEnabled = true
	}
	_gfd, _fdgg = _agec._bfbe.ReadBits(2)
	if _fdgg != nil {
		return _fdgg
	}
	_agec.HTemplate = byte(_gfd & 0xf)
	_fegd, _fdgg = _agec._bfbe.ReadBit()
	if _fdgg != nil {
		return _fdgg
	}
	if _fegd == 1 {
		_agec.IsMMREncoded = true
	}
	_gfd, _fdgg = _agec._bfbe.ReadBits(32)
	if _fdgg != nil {
		return _fdgg
	}
	_agec.HGridWidth = uint32(_gfd & _b.MaxUint32)
	_gfd, _fdgg = _agec._bfbe.ReadBits(32)
	if _fdgg != nil {
		return _fdgg
	}
	_agec.HGridHeight = uint32(_gfd & _b.MaxUint32)
	_gfd, _fdgg = _agec._bfbe.ReadBits(32)
	if _fdgg != nil {
		return _fdgg
	}
	_agec.HGridX = int32(_gfd & _b.MaxInt32)
	_gfd, _fdgg = _agec._bfbe.ReadBits(32)
	if _fdgg != nil {
		return _fdgg
	}
	_agec.HGridY = int32(_gfd & _b.MaxInt32)
	_gfd, _fdgg = _agec._bfbe.ReadBits(16)
	if _fdgg != nil {
		return _fdgg
	}
	_agec.HRegionX = uint16(_gfd & _b.MaxUint16)
	_gfd, _fdgg = _agec._bfbe.ReadBits(16)
	if _fdgg != nil {
		return _fdgg
	}
	_agec.HRegionY = uint16(_gfd & _b.MaxUint16)
	if _fdgg = _agec.computeSegmentDataStructure(); _fdgg != nil {
		return _fdgg
	}
	return _agec.checkInput()
}

func (_afgf *Header) readReferredToSegmentNumbers(_eggg *_e.Reader, _bgbb int) ([]int, error) {
	const _gegb = "\u0072\u0065\u0061\u0064R\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0054\u006f\u0053e\u0067m\u0065\u006e\u0074\u004e\u0075\u006d\u0062e\u0072\u0073"
	_dccf := make([]int, _bgbb)
	if _bgbb > 0 {
		_afgf.RTSegments = make([]*Header, _bgbb)
		var (
			_dfgd uint64
			_fdgc error
		)
		for _dbgg := 0; _dbgg < _bgbb; _dbgg++ {
			_dfgd, _fdgc = _eggg.ReadBits(byte(_afgf.referenceSize()) << 3)
			if _fdgc != nil {
				return nil, _da.Wrapf(_fdgc, _gegb, "\u0027\u0025\u0064\u0027 \u0072\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0020\u0073e\u0067m\u0065\u006e\u0074\u0020\u006e\u0075\u006db\u0065\u0072", _dbgg)
			}
			_dccf[_dbgg] = int(_dfgd & _b.MaxInt32)
		}
	}
	return _dccf, nil
}

func (_agga *SymbolDictionary) readAtPixels(_ebfg int) error {
	_agga.SdATX = make([]int8, _ebfg)
	_agga.SdATY = make([]int8, _ebfg)
	var (
		_fagfb byte
		_aadeb error
	)
	for _ggeb := 0; _ggeb < _ebfg; _ggeb++ {
		_fagfb, _aadeb = _agga._abege.ReadByte()
		if _aadeb != nil {
			return _aadeb
		}
		_agga.SdATX[_ggeb] = int8(_fagfb)
		_fagfb, _aadeb = _agga._abege.ReadByte()
		if _aadeb != nil {
			return _aadeb
		}
		_agga.SdATY[_ggeb] = int8(_fagfb)
	}
	return nil
}
func (_gbgf *GenericRegion) Size() int { return _gbgf.RegionSegment.Size() + 1 + 2*len(_gbgf.GBAtX) }
func (_aafg *GenericRegion) decodeTemplate1(_gff, _ggd, _cgeb int, _fce, _fdffc int) (_dga error) {
	const _caca = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0031"
	var (
		_gfbd, _efdf int
		_geaf, _gffe int
		_afa         byte
		_bbcd, _ade  int
	)
	if _gff >= 1 {
		_afa, _dga = _aafg.Bitmap.GetByte(_fdffc)
		if _dga != nil {
			return _da.Wrap(_dga, _caca, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_geaf = int(_afa)
	}
	if _gff >= 2 {
		_afa, _dga = _aafg.Bitmap.GetByte(_fdffc - _aafg.Bitmap.RowStride)
		if _dga != nil {
			return _da.Wrap(_dga, _caca, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_gffe = int(_afa) << 5
	}
	_gfbd = ((_geaf >> 1) & 0x1f8) | ((_gffe >> 1) & 0x1e00)
	for _fecg := 0; _fecg < _cgeb; _fecg = _bbcd {
		var (
			_cbg byte
			_fea int
		)
		_bbcd = _fecg + 8
		if _fff := _ggd - _fecg; _fff > 8 {
			_fea = 8
		} else {
			_fea = _fff
		}
		if _gff > 0 {
			_geaf <<= 8
			if _bbcd < _ggd {
				_afa, _dga = _aafg.Bitmap.GetByte(_fdffc + 1)
				if _dga != nil {
					return _da.Wrap(_dga, _caca, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_geaf |= int(_afa)
			}
		}
		if _gff > 1 {
			_gffe <<= 8
			if _bbcd < _ggd {
				_afa, _dga = _aafg.Bitmap.GetByte(_fdffc - _aafg.Bitmap.RowStride + 1)
				if _dga != nil {
					return _da.Wrap(_dga, _caca, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_gffe |= int(_afa) << 5
			}
		}
		for _ebfd := 0; _ebfd < _fea; _ebfd++ {
			if _aafg._aaf {
				_efdf = _aafg.overrideAtTemplate1(_gfbd, _fecg+_ebfd, _gff, int(_cbg), _ebfd)
				_aafg._eced.SetIndex(int32(_efdf))
			} else {
				_aafg._eced.SetIndex(int32(_gfbd))
			}
			_ade, _dga = _aafg._bda.DecodeBit(_aafg._eced)
			if _dga != nil {
				return _da.Wrap(_dga, _caca, "")
			}
			_cbg |= byte(_ade) << uint(7-_ebfd)
			_febgf := uint(8 - _ebfd)
			_gfbd = ((_gfbd & 0xefb) << 1) | _ade | ((_geaf >> _febgf) & 0x8) | ((_gffe >> _febgf) & 0x200)
		}
		if _cgg := _aafg.Bitmap.SetByte(_fce, _cbg); _cgg != nil {
			return _da.Wrap(_cgg, _caca, "")
		}
		_fce++
		_fdffc++
	}
	return nil
}

func (_eeb *Header) Encode(w _e.BinaryWriter) (_ddbb int, _cfgf error) {
	const _add = "\u0048\u0065\u0061d\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065"
	var _fbb _e.BinaryWriter
	_eg.Log.Trace("\u005b\u0053\u0045G\u004d\u0045\u004e\u0054-\u0048\u0045\u0041\u0044\u0045\u0052\u005d[\u0045\u004e\u0043\u004f\u0044\u0045\u005d\u0020\u0042\u0065\u0067\u0069\u006e\u0073")
	defer func() {
		if _cfgf != nil {
			_eg.Log.Trace("[\u0053\u0045\u0047\u004d\u0045\u004eT\u002d\u0048\u0045\u0041\u0044\u0045R\u005d\u005b\u0045\u004e\u0043\u004f\u0044E\u005d\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u002e\u0020%\u0076", _cfgf)
		} else {
			_eg.Log.Trace("\u005b\u0053\u0045\u0047ME\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0025\u0076", _eeb)
			_eg.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045N\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u005b\u0045\u004e\u0043O\u0044\u0045\u005d\u0020\u0046\u0069\u006ei\u0073\u0068\u0065\u0064")
		}
	}()
	w.FinishByte()
	if _eeb.SegmentData != nil {
		_gagg, _ggdc := _eeb.SegmentData.(SegmentEncoder)
		if !_ggdc {
			return 0, _da.Errorf(_add, "\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0054\u0020\u0064\u006f\u0065s\u006e\u0027\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074 \u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0045\u006e\u0063\u006f\u0064er\u0020\u0069\u006e\u0074\u0065\u0072\u0066\u0061\u0063\u0065", _eeb.SegmentData)
		}
		_fbb = _e.BufferedMSB()
		_ddbb, _cfgf = _gagg.Encode(_fbb)
		if _cfgf != nil {
			return 0, _da.Wrap(_cfgf, _add, "")
		}
		_eeb.SegmentDataLength = uint64(_ddbb)
	}
	if _eeb.pageSize() == 4 {
		_eeb.PageAssociationFieldSize = true
	}
	var _cace int
	_cace, _cfgf = _eeb.writeSegmentNumber(w)
	if _cfgf != nil {
		return 0, _da.Wrap(_cfgf, _add, "")
	}
	_ddbb += _cace
	if _cfgf = _eeb.writeFlags(w); _cfgf != nil {
		return _ddbb, _da.Wrap(_cfgf, _add, "")
	}
	_ddbb++
	_cace, _cfgf = _eeb.writeReferredToCount(w)
	if _cfgf != nil {
		return 0, _da.Wrap(_cfgf, _add, "")
	}
	_ddbb += _cace
	_cace, _cfgf = _eeb.writeReferredToSegments(w)
	if _cfgf != nil {
		return 0, _da.Wrap(_cfgf, _add, "")
	}
	_ddbb += _cace
	_cace, _cfgf = _eeb.writeSegmentPageAssociation(w)
	if _cfgf != nil {
		return 0, _da.Wrap(_cfgf, _add, "")
	}
	_ddbb += _cace
	_cace, _cfgf = _eeb.writeSegmentDataLength(w)
	if _cfgf != nil {
		return 0, _da.Wrap(_cfgf, _add, "")
	}
	_ddbb += _cace
	_eeb.HeaderLength = int64(_ddbb) - int64(_eeb.SegmentDataLength)
	if _fbb != nil {
		if _, _cfgf = w.Write(_fbb.Data()); _cfgf != nil {
			return _ddbb, _da.Wrap(_cfgf, _add, "\u0077r\u0069t\u0065\u0020\u0073\u0065\u0067m\u0065\u006et\u0020\u0064\u0061\u0074\u0061")
		}
	}
	return _ddbb, nil
}

func (_bbg *TextRegion) decodeDT() (_gecd int64, _gceg error) {
	if _bbg.IsHuffmanEncoded {
		if _bbg.SbHuffDT == 3 {
			_gecd, _gceg = _bbg._accc.Decode(_bbg._gdbbd)
			if _gceg != nil {
				return 0, _gceg
			}
		} else {
			var _fdfc _cg.Tabler
			_fdfc, _gceg = _cg.GetStandardTable(11 + int(_bbg.SbHuffDT))
			if _gceg != nil {
				return 0, _gceg
			}
			_gecd, _gceg = _fdfc.Decode(_bbg._gdbbd)
			if _gceg != nil {
				return 0, _gceg
			}
		}
	} else {
		var _dbbf int32
		_dbbf, _gceg = _bbg._bggc.DecodeInt(_bbg._ffae)
		if _gceg != nil {
			return
		}
		_gecd = int64(_dbbf)
	}
	_gecd *= int64(_bbg.SbStrips)
	return _gecd, nil
}

func (_ecgb *SymbolDictionary) encodeRefinementATFlags(_fbba _e.BinaryWriter) (_gdad int, _dbbe error) {
	const _eccg = "\u0065\u006e\u0063od\u0065\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0041\u0054\u0046\u006c\u0061\u0067\u0073"
	if !_ecgb.UseRefinementAggregation || _ecgb.SdrTemplate != 0 {
		return 0, nil
	}
	for _bbf := 0; _bbf < 2; _bbf++ {
		if _dbbe = _fbba.WriteByte(byte(_ecgb.SdrATX[_bbf])); _dbbe != nil {
			return _gdad, _da.Wrapf(_dbbe, _eccg, "\u0053\u0064\u0072\u0041\u0054\u0058\u005b\u0025\u0064\u005d", _bbf)
		}
		_gdad++
		if _dbbe = _fbba.WriteByte(byte(_ecgb.SdrATY[_bbf])); _dbbe != nil {
			return _gdad, _da.Wrapf(_dbbe, _eccg, "\u0053\u0064\u0072\u0041\u0054\u0059\u005b\u0025\u0064\u005d", _bbf)
		}
		_gdad++
	}
	return _gdad, nil
}

func (_dgde *GenericRegion) setOverrideFlag(_fgb int) {
	_dgde.GBAtOverride[_fgb] = true
	_dgde._aaf = true
}

func (_dbfde *TextRegion) parseHeader() error {
	var _eagg error
	_eg.Log.Trace("\u005b\u0054E\u0058\u0054\u0020\u0052E\u0047\u0049O\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045-\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0062\u0065\u0067\u0069n\u0073\u002e\u002e\u002e")
	defer func() {
		if _eagg != nil {
			_eg.Log.Trace("\u005b\u0054\u0045\u0058\u0054\u0020\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044E\u0052\u005d\u0020\u0066\u0061i\u006c\u0065d\u002e\u0020\u0025\u0076", _eagg)
		} else {
			_eg.Log.Trace("\u005b\u0054E\u0058\u0054\u0020\u0052E\u0047\u0049O\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045-\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0066\u0069\u006e\u0069s\u0068\u0065\u0064\u002e")
		}
	}()
	if _eagg = _dbfde.RegionInfo.parseHeader(); _eagg != nil {
		return _eagg
	}
	if _eagg = _dbfde.readRegionFlags(); _eagg != nil {
		return _eagg
	}
	if _dbfde.IsHuffmanEncoded {
		if _eagg = _dbfde.readHuffmanFlags(); _eagg != nil {
			return _eagg
		}
	}
	if _eagg = _dbfde.readUseRefinement(); _eagg != nil {
		return _eagg
	}
	if _eagg = _dbfde.readAmountOfSymbolInstances(); _eagg != nil {
		return _eagg
	}
	if _eagg = _dbfde.getSymbols(); _eagg != nil {
		return _eagg
	}
	if _eagg = _dbfde.computeSymbolCodeLength(); _eagg != nil {
		return _eagg
	}
	if _eagg = _dbfde.checkInput(); _eagg != nil {
		return _eagg
	}
	_eg.Log.Trace("\u0025\u0073", _dbfde.String())
	return nil
}

func (_gdgbe *Header) writeSegmentPageAssociation(_dcga _e.BinaryWriter) (_ccefd int, _efcd error) {
	const _eeae = "w\u0072\u0069\u0074\u0065\u0053\u0065g\u006d\u0065\u006e\u0074\u0050\u0061\u0067\u0065\u0041s\u0073\u006f\u0063i\u0061t\u0069\u006f\u006e"
	if _gdgbe.pageSize() != 4 {
		if _efcd = _dcga.WriteByte(byte(_gdgbe.PageAssociation)); _efcd != nil {
			return 0, _da.Wrap(_efcd, _eeae, "\u0070\u0061\u0067\u0065\u0053\u0069\u007a\u0065\u0020\u0021\u003d\u0020\u0034")
		}
		return 1, nil
	}
	_cdfa := make([]byte, 4)
	_ab.BigEndian.PutUint32(_cdfa, uint32(_gdgbe.PageAssociation))
	if _ccefd, _efcd = _dcga.Write(_cdfa); _efcd != nil {
		return 0, _da.Wrap(_efcd, _eeae, "\u0034 \u0062y\u0074\u0065\u0020\u0070\u0061g\u0065\u0020n\u0075\u006d\u0062\u0065\u0072")
	}
	return _ccefd, nil
}
func (_daeb *TableSegment) HtLow() int32 { return _daeb._aegdd }
func (_addf *PatternDictionary) readPatternWidthAndHeight() error {
	_acc, _aeccc := _addf._ccca.ReadByte()
	if _aeccc != nil {
		return _aeccc
	}
	_addf.HdpWidth = _acc
	_acc, _aeccc = _addf._ccca.ReadByte()
	if _aeccc != nil {
		return _aeccc
	}
	_addf.HdpHeight = _acc
	return nil
}

func (_dge *template0) form(_bef, _dbdd, _eff, _aff, _ddg int16) int16 {
	return (_bef << 10) | (_dbdd << 7) | (_eff << 4) | (_aff << 1) | _ddg
}

func (_afgb *Header) writeReferredToCount(_dcf _e.BinaryWriter) (_bfc int, _dfgb error) {
	const _eef = "w\u0072i\u0074\u0065\u0052\u0065\u0066\u0065\u0072\u0072e\u0064\u0054\u006f\u0043ou\u006e\u0074"
	_afgb.RTSNumbers = make([]int, len(_afgb.RTSegments))
	for _cdeg, _dada := range _afgb.RTSegments {
		_afgb.RTSNumbers[_cdeg] = int(_dada.SegmentNumber)
	}
	if len(_afgb.RTSNumbers) <= 4 {
		var _aegd byte
		if len(_afgb.RetainBits) >= 1 {
			_aegd = _afgb.RetainBits[0]
		}
		_aegd |= byte(len(_afgb.RTSNumbers)) << 5
		if _dfgb = _dcf.WriteByte(_aegd); _dfgb != nil {
			return 0, _da.Wrap(_dfgb, _eef, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		return 1, nil
	}
	_ffdf := uint32(len(_afgb.RTSNumbers))
	_badf := make([]byte, 4+_dc.Ceil(len(_afgb.RTSNumbers)+1, 8))
	_ffdf |= 0x7 << 29
	_ab.BigEndian.PutUint32(_badf, _ffdf)
	copy(_badf[1:], _afgb.RetainBits)
	_bfc, _dfgb = _dcf.Write(_badf)
	if _dfgb != nil {
		return 0, _da.Wrap(_dfgb, _eef, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	return _bfc, nil
}

func (_gggag *SymbolDictionary) decodeAggregate(_afaf, _babd uint32) error {
	var (
		_eaef  int64
		_fbaeg error
	)
	if _gggag.IsHuffmanEncoded {
		_eaef, _fbaeg = _gggag.huffDecodeRefAggNInst()
		if _fbaeg != nil {
			return _fbaeg
		}
	} else {
		_baffb, _faedg := _gggag._agbbf.DecodeInt(_gggag._gaed)
		if _faedg != nil {
			return _faedg
		}
		_eaef = int64(_baffb)
	}
	if _eaef > 1 {
		return _gggag.decodeThroughTextRegion(_afaf, _babd, uint32(_eaef))
	} else if _eaef == 1 {
		return _gggag.decodeRefinedSymbol(_afaf, _babd)
	}
	return nil
}

func (_ffcd *TextRegion) symbolIDCodeLengths() error {
	var (
		_gege []*_cg.Code
		_cced uint64
		_edag _cg.Tabler
		_edfc error
	)
	for _bfeaa := 0; _bfeaa < 35; _bfeaa++ {
		_cced, _edfc = _ffcd._gdbbd.ReadBits(4)
		if _edfc != nil {
			return _edfc
		}
		_ffbdc := int(_cced & 0xf)
		if _ffbdc > 0 {
			_gege = append(_gege, _cg.NewCode(int32(_ffbdc), 0, int32(_bfeaa), false))
		}
	}
	_edag, _edfc = _cg.NewFixedSizeTable(_gege)
	if _edfc != nil {
		return _edfc
	}
	var (
		_bggg  int64
		_fdffd uint32
		_bdae  []*_cg.Code
		_fgea  int64
	)
	for _fdffd < _ffcd.NumberOfSymbols {
		_fgea, _edfc = _edag.Decode(_ffcd._gdbbd)
		if _edfc != nil {
			return _edfc
		}
		if _fgea < 32 {
			if _fgea > 0 {
				_bdae = append(_bdae, _cg.NewCode(int32(_fgea), 0, int32(_fdffd), false))
			}
			_bggg = _fgea
			_fdffd++
		} else {
			var _fdef, _ebbc int64
			switch _fgea {
			case 32:
				_cced, _edfc = _ffcd._gdbbd.ReadBits(2)
				if _edfc != nil {
					return _edfc
				}
				_fdef = 3 + int64(_cced)
				if _fdffd > 0 {
					_ebbc = _bggg
				}
			case 33:
				_cced, _edfc = _ffcd._gdbbd.ReadBits(3)
				if _edfc != nil {
					return _edfc
				}
				_fdef = 3 + int64(_cced)
			case 34:
				_cced, _edfc = _ffcd._gdbbd.ReadBits(7)
				if _edfc != nil {
					return _edfc
				}
				_fdef = 11 + int64(_cced)
			}
			for _cbfc := 0; _cbfc < int(_fdef); _cbfc++ {
				if _ebbc > 0 {
					_bdae = append(_bdae, _cg.NewCode(int32(_ebbc), 0, int32(_fdffd), false))
				}
				_fdffd++
			}
		}
	}
	_ffcd._gdbbd.Align()
	_ffcd._adcf, _edfc = _cg.NewFixedSizeTable(_bdae)
	return _edfc
}

func NewGenericRegion(r *_e.Reader) *GenericRegion {
	return &GenericRegion{RegionSegment: NewRegionSegment(r), _fec: r}
}

func (_abc *GenericRefinementRegion) decodeTypicalPredictedLineTemplate0(_eaec, _efb, _egc, _abb, _aca, _fc, _df, _dfg, _cad int) error {
	var (
		_ee, _bad, _ded, _fa, _fg, _ddab int
		_afg                             byte
		_ec                              error
	)
	if _eaec > 0 {
		_afg, _ec = _abc.RegionBitmap.GetByte(_df - _egc)
		if _ec != nil {
			return _ec
		}
		_ded = int(_afg)
	}
	if _dfg > 0 && _dfg <= _abc.ReferenceBitmap.Height {
		_afg, _ec = _abc.ReferenceBitmap.GetByte(_cad - _abb + _fc)
		if _ec != nil {
			return _ec
		}
		_fa = int(_afg) << 4
	}
	if _dfg >= 0 && _dfg < _abc.ReferenceBitmap.Height {
		_afg, _ec = _abc.ReferenceBitmap.GetByte(_cad + _fc)
		if _ec != nil {
			return _ec
		}
		_fg = int(_afg) << 1
	}
	if _dfg > -2 && _dfg < _abc.ReferenceBitmap.Height-1 {
		_afg, _ec = _abc.ReferenceBitmap.GetByte(_cad + _abb + _fc)
		if _ec != nil {
			return _ec
		}
		_ddab = int(_afg)
	}
	_ee = ((_ded >> 5) & 0x6) | ((_ddab >> 2) & 0x30) | (_fg & 0x180) | (_fa & 0xc00)
	var _cd int
	for _deg := 0; _deg < _aca; _deg = _cd {
		var _aee int
		_cd = _deg + 8
		var _ddc int
		if _ddc = _efb - _deg; _ddc > 8 {
			_ddc = 8
		}
		_cfb := _cd < _efb
		_abbf := _cd < _abc.ReferenceBitmap.Width
		_gbag := _fc + 1
		if _eaec > 0 {
			_afg = 0
			if _cfb {
				_afg, _ec = _abc.RegionBitmap.GetByte(_df - _egc + 1)
				if _ec != nil {
					return _ec
				}
			}
			_ded = (_ded << 8) | int(_afg)
		}
		if _dfg > 0 && _dfg <= _abc.ReferenceBitmap.Height {
			var _bag int
			if _abbf {
				_afg, _ec = _abc.ReferenceBitmap.GetByte(_cad - _abb + _gbag)
				if _ec != nil {
					return _ec
				}
				_bag = int(_afg) << 4
			}
			_fa = (_fa << 8) | _bag
		}
		if _dfg >= 0 && _dfg < _abc.ReferenceBitmap.Height {
			var _bd int
			if _abbf {
				_afg, _ec = _abc.ReferenceBitmap.GetByte(_cad + _gbag)
				if _ec != nil {
					return _ec
				}
				_bd = int(_afg) << 1
			}
			_fg = (_fg << 8) | _bd
		}
		if _dfg > -2 && _dfg < (_abc.ReferenceBitmap.Height-1) {
			_afg = 0
			if _abbf {
				_afg, _ec = _abc.ReferenceBitmap.GetByte(_cad + _abb + _gbag)
				if _ec != nil {
					return _ec
				}
			}
			_ddab = (_ddab << 8) | int(_afg)
		}
		for _dfga := 0; _dfga < _ddc; _dfga++ {
			var _be int
			_fgf := false
			_cac := (_ee >> 4) & 0x1ff
			if _cac == 0x1ff {
				_fgf = true
				_be = 1
			} else if _cac == 0x00 {
				_fgf = true
			}
			if !_fgf {
				if _abc._bba {
					_bad = _abc.overrideAtTemplate0(_ee, _deg+_dfga, _eaec, _aee, _dfga)
					_abc._gbf.SetIndex(int32(_bad))
				} else {
					_abc._gbf.SetIndex(int32(_ee))
				}
				_be, _ec = _abc._cc.DecodeBit(_abc._gbf)
				if _ec != nil {
					return _ec
				}
			}
			_ggf := uint(7 - _dfga)
			_aee |= _be << _ggf
			_ee = ((_ee & 0xdb6) << 1) | _be | (_ded>>_ggf+5)&0x002 | ((_ddab>>_ggf + 2) & 0x010) | ((_fg >> _ggf) & 0x080) | ((_fa >> _ggf) & 0x400)
		}
		_ec = _abc.RegionBitmap.SetByte(_df, byte(_aee))
		if _ec != nil {
			return _ec
		}
		_df++
		_cad++
	}
	return nil
}

func (_dbfd Type) String() string {
	switch _dbfd {
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

func _dbb(_gea *_e.Reader, _dfb *Header) *GenericRefinementRegion {
	return &GenericRefinementRegion{_gbc: _gea, RegionInfo: NewRegionSegment(_gea), _bg: _dfb, _agbg: &template0{}, _ad: &template1{}}
}

func (_bab *PageInformationSegment) readContainsRefinement() error {
	_afeg, _defa := _bab._dfdf.ReadBit()
	if _defa != nil {
		return _defa
	}
	if _afeg == 1 {
		_bab._fagf = true
	}
	return nil
}

func (_eaa *GenericRefinementRegion) decodeTemplate(_aefg, _bbca, _beg, _cec, _cgc, _abcf, _bcb, _gcd, _dbd, _abfa int, _dba templater) (_fbd error) {
	var (
		_gcf, _gde, _adb, _fgg, _fac int16
		_cgcg, _dadb, _efbe, _bcbd   int
		_gdff                        byte
	)
	if _dbd >= 1 && (_dbd-1) < _eaa.ReferenceBitmap.Height {
		_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa - _cec)
		if _fbd != nil {
			return
		}
		_cgcg = int(_gdff)
	}
	if _dbd >= 0 && (_dbd) < _eaa.ReferenceBitmap.Height {
		_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa)
		if _fbd != nil {
			return
		}
		_dadb = int(_gdff)
	}
	if _dbd >= -1 && (_dbd+1) < _eaa.ReferenceBitmap.Height {
		_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa + _cec)
		if _fbd != nil {
			return
		}
		_efbe = int(_gdff)
	}
	_abfa++
	if _aefg >= 1 {
		_gdff, _fbd = _eaa.RegionBitmap.GetByte(_gcd - _beg)
		if _fbd != nil {
			return
		}
		_bcbd = int(_gdff)
	}
	_gcd++
	_cdf := _eaa.ReferenceDX % 8
	_bce := 6 + _cdf
	_aega := _abfa % _cec
	if _bce >= 0 {
		if _bce < 8 {
			_gcf = int16(_cgcg>>uint(_bce)) & 0x07
		}
		if _bce < 8 {
			_gde = int16(_dadb>>uint(_bce)) & 0x07
		}
		if _bce < 8 {
			_adb = int16(_efbe>>uint(_bce)) & 0x07
		}
		if _bce == 6 && _aega > 1 {
			if _dbd >= 1 && (_dbd-1) < _eaa.ReferenceBitmap.Height {
				_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa - _cec - 2)
				if _fbd != nil {
					return _fbd
				}
				_gcf |= int16(_gdff<<2) & 0x04
			}
			if _dbd >= 0 && _dbd < _eaa.ReferenceBitmap.Height {
				_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa - 2)
				if _fbd != nil {
					return _fbd
				}
				_gde |= int16(_gdff<<2) & 0x04
			}
			if _dbd >= -1 && _dbd+1 < _eaa.ReferenceBitmap.Height {
				_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa + _cec - 2)
				if _fbd != nil {
					return _fbd
				}
				_adb |= int16(_gdff<<2) & 0x04
			}
		}
		if _bce == 0 {
			_cgcg = 0
			_dadb = 0
			_efbe = 0
			if _aega < _cec-1 {
				if _dbd >= 1 && _dbd-1 < _eaa.ReferenceBitmap.Height {
					_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa - _cec)
					if _fbd != nil {
						return _fbd
					}
					_cgcg = int(_gdff)
				}
				if _dbd >= 0 && _dbd < _eaa.ReferenceBitmap.Height {
					_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa)
					if _fbd != nil {
						return _fbd
					}
					_dadb = int(_gdff)
				}
				if _dbd >= -1 && _dbd+1 < _eaa.ReferenceBitmap.Height {
					_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa + _cec)
					if _fbd != nil {
						return _fbd
					}
					_efbe = int(_gdff)
				}
			}
			_abfa++
		}
	} else {
		_gcf = int16(_cgcg<<1) & 0x07
		_gde = int16(_dadb<<1) & 0x07
		_adb = int16(_efbe<<1) & 0x07
		_cgcg = 0
		_dadb = 0
		_efbe = 0
		if _aega < _cec-1 {
			if _dbd >= 1 && _dbd-1 < _eaa.ReferenceBitmap.Height {
				_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa - _cec)
				if _fbd != nil {
					return _fbd
				}
				_cgcg = int(_gdff)
			}
			if _dbd >= 0 && _dbd < _eaa.ReferenceBitmap.Height {
				_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa)
				if _fbd != nil {
					return _fbd
				}
				_dadb = int(_gdff)
			}
			if _dbd >= -1 && _dbd+1 < _eaa.ReferenceBitmap.Height {
				_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa + _cec)
				if _fbd != nil {
					return _fbd
				}
				_efbe = int(_gdff)
			}
			_abfa++
		}
		_gcf |= int16((_cgcg >> 7) & 0x07)
		_gde |= int16((_dadb >> 7) & 0x07)
		_adb |= int16((_efbe >> 7) & 0x07)
	}
	_fgg = int16(_bcbd >> 6)
	_fac = 0
	_dbc := (2 - _cdf) % 8
	_cgcg <<= uint(_dbc)
	_dadb <<= uint(_dbc)
	_efbe <<= uint(_dbc)
	_bcbd <<= 2
	var _dbf int
	for _gaae := 0; _gaae < _bbca; _gaae++ {
		_cb := _gaae & 0x07
		_dbea := _dba.form(_gcf, _gde, _adb, _fgg, _fac)
		if _eaa._bba {
			_gdff, _fbd = _eaa.RegionBitmap.GetByte(_eaa.RegionBitmap.GetByteIndex(_gaae, _aefg))
			if _fbd != nil {
				return _fbd
			}
			_eaa._gbf.SetIndex(int32(_eaa.overrideAtTemplate0(int(_dbea), _gaae, _aefg, int(_gdff), _cb)))
		} else {
			_eaa._gbf.SetIndex(int32(_dbea))
		}
		_dbf, _fbd = _eaa._cc.DecodeBit(_eaa._gbf)
		if _fbd != nil {
			return _fbd
		}
		if _fbd = _eaa.RegionBitmap.SetPixel(_gaae, _aefg, byte(_dbf)); _fbd != nil {
			return _fbd
		}
		_gcf = ((_gcf << 1) | 0x01&int16(_cgcg>>7)) & 0x07
		_gde = ((_gde << 1) | 0x01&int16(_dadb>>7)) & 0x07
		_adb = ((_adb << 1) | 0x01&int16(_efbe>>7)) & 0x07
		_fgg = ((_fgg << 1) | 0x01&int16(_bcbd>>7)) & 0x07
		_fac = int16(_dbf)
		if (_gaae-int(_eaa.ReferenceDX))%8 == 5 {
			_cgcg = 0
			_dadb = 0
			_efbe = 0
			if ((_gaae-int(_eaa.ReferenceDX))/8)+1 < _eaa.ReferenceBitmap.RowStride {
				if _dbd >= 1 && (_dbd-1) < _eaa.ReferenceBitmap.Height {
					_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa - _cec)
					if _fbd != nil {
						return _fbd
					}
					_cgcg = int(_gdff)
				}
				if _dbd >= 0 && _dbd < _eaa.ReferenceBitmap.Height {
					_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa)
					if _fbd != nil {
						return _fbd
					}
					_dadb = int(_gdff)
				}
				if _dbd >= -1 && (_dbd+1) < _eaa.ReferenceBitmap.Height {
					_gdff, _fbd = _eaa.ReferenceBitmap.GetByte(_abfa + _cec)
					if _fbd != nil {
						return _fbd
					}
					_efbe = int(_gdff)
				}
			}
			_abfa++
		} else {
			_cgcg <<= 1
			_dadb <<= 1
			_efbe <<= 1
		}
		if _cb == 5 && _aefg >= 1 {
			if ((_gaae >> 3) + 1) >= _eaa.RegionBitmap.RowStride {
				_bcbd = 0
			} else {
				_gdff, _fbd = _eaa.RegionBitmap.GetByte(_gcd - _beg)
				if _fbd != nil {
					return _fbd
				}
				_bcbd = int(_gdff)
			}
			_gcd++
		} else {
			_bcbd <<= 1
		}
	}
	return nil
}

func (_fe *EndOfStripe) parseHeader() error {
	_gb, _ga := _fe._agb.ReadBits(32)
	if _ga != nil {
		return _ga
	}
	_fe._fb = int(_gb & _b.MaxInt32)
	return nil
}

func (_ebbf *SymbolDictionary) readNumberOfExportedSymbols() error {
	_dbdg, _fabb := _ebbf._abege.ReadBits(32)
	if _fabb != nil {
		return _fabb
	}
	_ebbf.NumberOfExportedSymbols = uint32(_dbdg & _b.MaxUint32)
	return nil
}

func (_fda *SymbolDictionary) setExportedSymbols(_bgd []int) {
	for _gfgf := uint32(0); _gfgf < _fda._afag+_fda.NumberOfNewSymbols; _gfgf++ {
		if _bgd[_gfgf] == 1 {
			var _acdf *_ea.Bitmap
			if _gfgf < _fda._afag {
				_acdf = _fda._dcgb[_gfgf]
			} else {
				_acdf = _fda._abed[_gfgf-_fda._afag]
			}
			_eg.Log.Trace("\u005bS\u0059\u004dB\u004f\u004c\u002d\u0044I\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u0020\u0041\u0064\u0064 E\u0078\u0070\u006fr\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c:\u0020\u0027%\u0073\u0027", _acdf)
			_fda._cfbdb = append(_fda._cfbdb, _acdf)
		}
	}
}
func (_bcbge *TextRegion) GetRegionInfo() *RegionSegment { return _bcbge.RegionInfo }
func NewRegionSegment(r *_e.Reader) *RegionSegment       { return &RegionSegment{_dfbgc: r} }
func (_ega *PageInformationSegment) readIsLossless() error {
	_bfeb, _gcgg := _ega._dfdf.ReadBit()
	if _gcgg != nil {
		return _gcgg
	}
	if _bfeb == 1 {
		_ega.IsLossless = true
	}
	return nil
}

func (_ggc *GenericRegion) decodeTemplate0b(_cfg, _ged, _cgb int, _dfbg, _dgff int) (_dbfe error) {
	const _dafe = "\u0064\u0065c\u006f\u0064\u0065T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0030\u0062"
	var (
		_agf, _geac  int
		_aecb, _beeg int
		_efee        byte
		_dcb         int
	)
	if _cfg >= 1 {
		_efee, _dbfe = _ggc.Bitmap.GetByte(_dgff)
		if _dbfe != nil {
			return _da.Wrap(_dbfe, _dafe, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_aecb = int(_efee)
	}
	if _cfg >= 2 {
		_efee, _dbfe = _ggc.Bitmap.GetByte(_dgff - _ggc.Bitmap.RowStride)
		if _dbfe != nil {
			return _da.Wrap(_dbfe, _dafe, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_beeg = int(_efee) << 6
	}
	_agf = (_aecb & 0xf0) | (_beeg & 0x3800)
	for _bebd := 0; _bebd < _cgb; _bebd = _dcb {
		var (
			_bacd byte
			_acff int
		)
		_dcb = _bebd + 8
		if _bebb := _ged - _bebd; _bebb > 8 {
			_acff = 8
		} else {
			_acff = _bebb
		}
		if _cfg > 0 {
			_aecb <<= 8
			if _dcb < _ged {
				_efee, _dbfe = _ggc.Bitmap.GetByte(_dgff + 1)
				if _dbfe != nil {
					return _da.Wrap(_dbfe, _dafe, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_aecb |= int(_efee)
			}
		}
		if _cfg > 1 {
			_beeg <<= 8
			if _dcb < _ged {
				_efee, _dbfe = _ggc.Bitmap.GetByte(_dgff - _ggc.Bitmap.RowStride + 1)
				if _dbfe != nil {
					return _da.Wrap(_dbfe, _dafe, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_beeg |= int(_efee) << 6
			}
		}
		for _gcg := 0; _gcg < _acff; _gcg++ {
			_afc := uint(7 - _gcg)
			if _ggc._aaf {
				_geac = _ggc.overrideAtTemplate0b(_agf, _bebd+_gcg, _cfg, int(_bacd), _gcg, int(_afc))
				_ggc._eced.SetIndex(int32(_geac))
			} else {
				_ggc._eced.SetIndex(int32(_agf))
			}
			var _gegc int
			_gegc, _dbfe = _ggc._bda.DecodeBit(_ggc._eced)
			if _dbfe != nil {
				return _da.Wrap(_dbfe, _dafe, "")
			}
			_bacd |= byte(_gegc << _afc)
			_agf = ((_agf & 0x7bf7) << 1) | _gegc | ((_aecb >> _afc) & 0x10) | ((_beeg >> _afc) & 0x800)
		}
		if _aad := _ggc.Bitmap.SetByte(_dfbg, _bacd); _aad != nil {
			return _da.Wrap(_aad, _dafe, "")
		}
		_dfbg++
		_dgff++
	}
	return nil
}

func (_eaag *PatternDictionary) parseHeader() error {
	_eg.Log.Trace("\u005b\u0050\u0041\u0054\u0054\u0045\u0052\u004e\u002d\u0044\u0049\u0043\u0054I\u004f\u004e\u0041\u0052\u0059\u005d[\u0070\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072\u005d\u0020b\u0065\u0067\u0069\u006e")
	defer func() {
		_eg.Log.Trace("\u005b\u0050\u0041T\u0054\u0045\u0052\u004e\u002d\u0044\u0049\u0043\u0054\u0049\u004f\u004e\u0041\u0052\u0059\u005d\u005b\u0070\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072\u005d \u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
	}()
	_, _ffef := _eaag._ccca.ReadBits(5)
	if _ffef != nil {
		return _ffef
	}
	if _ffef = _eaag.readTemplate(); _ffef != nil {
		return _ffef
	}
	if _ffef = _eaag.readIsMMREncoded(); _ffef != nil {
		return _ffef
	}
	if _ffef = _eaag.readPatternWidthAndHeight(); _ffef != nil {
		return _ffef
	}
	if _ffef = _eaag.readGrayMax(); _ffef != nil {
		return _ffef
	}
	if _ffef = _eaag.computeSegmentDataStructure(); _ffef != nil {
		return _ffef
	}
	return _eaag.checkInput()
}

func (_cfcb *TextRegion) setCodingStatistics() error {
	if _cfcb._ffae == nil {
		_cfcb._ffae = _ag.NewStats(512, 1)
	}
	if _cfcb._ccfc == nil {
		_cfcb._ccfc = _ag.NewStats(512, 1)
	}
	if _cfcb._fbfg == nil {
		_cfcb._fbfg = _ag.NewStats(512, 1)
	}
	if _cfcb._efgab == nil {
		_cfcb._efgab = _ag.NewStats(512, 1)
	}
	if _cfcb._eacf == nil {
		_cfcb._eacf = _ag.NewStats(512, 1)
	}
	if _cfcb._bgcbb == nil {
		_cfcb._bgcbb = _ag.NewStats(512, 1)
	}
	if _cfcb._dcgcb == nil {
		_cfcb._dcgcb = _ag.NewStats(512, 1)
	}
	if _cfcb._cgfd == nil {
		_cfcb._cgfd = _ag.NewStats(1<<uint(_cfcb._ggggf), 1)
	}
	if _cfcb._dcfba == nil {
		_cfcb._dcfba = _ag.NewStats(512, 1)
	}
	if _cfcb._feab == nil {
		_cfcb._feab = _ag.NewStats(512, 1)
	}
	if _cfcb._bggc == nil {
		var _ccfe error
		_cfcb._bggc, _ccfe = _ag.New(_cfcb._gdbbd)
		if _ccfe != nil {
			return _ccfe
		}
	}
	return nil
}

func (_edg *SymbolDictionary) decodeHeightClassBitmap(_fged *_ea.Bitmap, _efdgc int64, _fgce int, _efbea []int) error {
	for _egee := _efdgc; _egee < int64(_edg._addfb); _egee++ {
		var _cafb int
		for _fega := _efdgc; _fega <= _egee-1; _fega++ {
			_cafb += _efbea[_fega]
		}
		_agaed := _c.Rect(_cafb, 0, _cafb+_efbea[_egee], _fgce)
		_eecg, _eccc := _ea.Extract(_agaed, _fged)
		if _eccc != nil {
			return _eccc
		}
		_edg._abed[_egee] = _eecg
		_edg._ffbb = append(_edg._ffbb, _eecg)
	}
	return nil
}

func (_fbeb *Header) parse(_ege Documenter, _fcc *_e.Reader, _afd int64, _feag OrganizationType) (_adfc error) {
	const _cde = "\u0070\u0061\u0072s\u0065"
	_eg.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u002d\u0048E\u0041\u0044\u0045\u0052\u005d\u005b\u0050A\u0052\u0053\u0045\u005d\u0020\u0042\u0065\u0067\u0069\u006e\u0073")
	defer func() {
		if _adfc != nil {
			_eg.Log.Trace("\u005b\u0053\u0045GM\u0045\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045R\u005d[\u0050A\u0052S\u0045\u005d\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _adfc)
		} else {
			_eg.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052]\u005bP\u0041\u0052\u0053\u0045\u005d\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
		}
	}()
	_, _adfc = _fcc.Seek(_afd, _d.SeekStart)
	if _adfc != nil {
		return _da.Wrap(_adfc, _cde, "\u0073\u0065\u0065\u006b\u0020\u0073\u0074\u0061\u0072\u0074")
	}
	if _adfc = _fbeb.readSegmentNumber(_fcc); _adfc != nil {
		return _da.Wrap(_adfc, _cde, "")
	}
	if _adfc = _fbeb.readHeaderFlags(); _adfc != nil {
		return _da.Wrap(_adfc, _cde, "")
	}
	var _egcf uint64
	_egcf, _adfc = _fbeb.readNumberOfReferredToSegments(_fcc)
	if _adfc != nil {
		return _da.Wrap(_adfc, _cde, "")
	}
	_fbeb.RTSNumbers, _adfc = _fbeb.readReferredToSegmentNumbers(_fcc, int(_egcf))
	if _adfc != nil {
		return _da.Wrap(_adfc, _cde, "")
	}
	_adfc = _fbeb.readSegmentPageAssociation(_ege, _fcc, _egcf, _fbeb.RTSNumbers...)
	if _adfc != nil {
		return _da.Wrap(_adfc, _cde, "")
	}
	if _fbeb.Type != TEndOfFile {
		if _adfc = _fbeb.readSegmentDataLength(_fcc); _adfc != nil {
			return _da.Wrap(_adfc, _cde, "")
		}
	}
	_fbeb.readDataStartOffset(_fcc, _feag)
	_fbeb.readHeaderLength(_fcc, _afd)
	_eg.Log.Trace("\u0025\u0073", _fbeb)
	return nil
}

func (_ddcd *Header) readDataStartOffset(_geag *_e.Reader, _dbae OrganizationType) {
	if _dbae == OSequential {
		_ddcd.SegmentDataStartOffset = uint64(_geag.AbsolutePosition())
	}
}

func (_dfbd *PatternDictionary) GetDictionary() ([]*_ea.Bitmap, error) {
	if _dfbd.Patterns != nil {
		return _dfbd.Patterns, nil
	}
	if !_dfbd.IsMMREncoded {
		_dfbd.setGbAtPixels()
	}
	_bbac := NewGenericRegion(_dfbd._ccca)
	_bbac.setParametersMMR(_dfbd.IsMMREncoded, _dfbd.DataOffset, _dfbd.DataLength, uint32(_dfbd.HdpHeight), (_dfbd.GrayMax+1)*uint32(_dfbd.HdpWidth), _dfbd.HDTemplate, false, false, _dfbd.GBAtX, _dfbd.GBAtY)
	_gge, _acaca := _bbac.GetRegionBitmap()
	if _acaca != nil {
		return nil, _acaca
	}
	if _acaca = _dfbd.extractPatterns(_gge); _acaca != nil {
		return nil, _acaca
	}
	return _dfbd.Patterns, nil
}

func (_cgac *TextRegion) encodeFlags(_dafgc _e.BinaryWriter) (_dece int, _bagf error) {
	const _fdgb = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _bagf = _dafgc.WriteBit(int(_cgac.SbrTemplate)); _bagf != nil {
		return _dece, _da.Wrap(_bagf, _fdgb, "s\u0062\u0072\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	if _, _bagf = _dafgc.WriteBits(uint64(_cgac.SbDsOffset), 5); _bagf != nil {
		return _dece, _da.Wrap(_bagf, _fdgb, "\u0073\u0062\u0044\u0073\u004f\u0066\u0066\u0073\u0065\u0074")
	}
	if _bagf = _dafgc.WriteBit(int(_cgac.DefaultPixel)); _bagf != nil {
		return _dece, _da.Wrap(_bagf, _fdgb, "\u0044\u0065\u0066a\u0075\u006c\u0074\u0050\u0069\u0078\u0065\u006c")
	}
	if _, _bagf = _dafgc.WriteBits(uint64(_cgac.CombinationOperator), 2); _bagf != nil {
		return _dece, _da.Wrap(_bagf, _fdgb, "\u0043\u006f\u006d\u0062in\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	if _bagf = _dafgc.WriteBit(int(_cgac.IsTransposed)); _bagf != nil {
		return _dece, _da.Wrap(_bagf, _fdgb, "\u0069\u0073\u0020\u0074\u0072\u0061\u006e\u0073\u0070\u006f\u0073\u0065\u0064")
	}
	if _, _bagf = _dafgc.WriteBits(uint64(_cgac.ReferenceCorner), 2); _bagf != nil {
		return _dece, _da.Wrap(_bagf, _fdgb, "\u0072\u0065f\u0065\u0072\u0065n\u0063\u0065\u0020\u0063\u006f\u0072\u006e\u0065\u0072")
	}
	if _, _bagf = _dafgc.WriteBits(uint64(_cgac.LogSBStrips), 2); _bagf != nil {
		return _dece, _da.Wrap(_bagf, _fdgb, "L\u006f\u0067\u0053\u0042\u0053\u0074\u0072\u0069\u0070\u0073")
	}
	var _bdad int
	if _cgac.UseRefinement {
		_bdad = 1
	}
	if _bagf = _dafgc.WriteBit(_bdad); _bagf != nil {
		return _dece, _da.Wrap(_bagf, _fdgb, "\u0075\u0073\u0065\u0020\u0072\u0065\u0066\u0069\u006ee\u006d\u0065\u006e\u0074")
	}
	_bdad = 0
	if _cgac.IsHuffmanEncoded {
		_bdad = 1
	}
	if _bagf = _dafgc.WriteBit(_bdad); _bagf != nil {
		return _dece, _da.Wrap(_bagf, _fdgb, "u\u0073\u0065\u0020\u0068\u0075\u0066\u0066\u006d\u0061\u006e")
	}
	_dece = 2
	return _dece, nil
}

func (_faac *GenericRegion) setParametersMMR(_dfgf bool, _aede, _eea int64, _ddae, _cfa uint32, _bgbc byte, _bgga, _agff bool, _dced, _ffc []int8) {
	_faac.DataOffset = _aede
	_faac.DataLength = _eea
	_faac.RegionSegment = &RegionSegment{}
	_faac.RegionSegment.BitmapHeight = _ddae
	_faac.RegionSegment.BitmapWidth = _cfa
	_faac.GBTemplate = _bgbc
	_faac.IsMMREncoded = _dfgf
	_faac.IsTPGDon = _bgga
	_faac.GBAtX = _dced
	_faac.GBAtY = _ffc
}

func (_afga *GenericRefinementRegion) decodeTypicalPredictedLineTemplate1(_egg, _ce, _edfd, _agg, _eda, _gad, _gaa, _efg, _eag int) (_eagb error) {
	var (
		_afe, _age int
		_aed, _bbc int
		_efd, _gga int
		_efde      byte
	)
	if _egg > 0 {
		_efde, _eagb = _afga.RegionBitmap.GetByte(_gaa - _edfd)
		if _eagb != nil {
			return
		}
		_aed = int(_efde)
	}
	if _efg > 0 && _efg <= _afga.ReferenceBitmap.Height {
		_efde, _eagb = _afga.ReferenceBitmap.GetByte(_eag - _agg + _gad)
		if _eagb != nil {
			return
		}
		_bbc = int(_efde) << 2
	}
	if _efg >= 0 && _efg < _afga.ReferenceBitmap.Height {
		_efde, _eagb = _afga.ReferenceBitmap.GetByte(_eag + _gad)
		if _eagb != nil {
			return
		}
		_efd = int(_efde)
	}
	if _efg > -2 && _efg < _afga.ReferenceBitmap.Height-1 {
		_efde, _eagb = _afga.ReferenceBitmap.GetByte(_eag + _agg + _gad)
		if _eagb != nil {
			return
		}
		_gga = int(_efde)
	}
	_afe = ((_aed >> 5) & 0x6) | ((_gga >> 2) & 0x30) | (_efd & 0xc0) | (_bbc & 0x200)
	_age = ((_gga >> 2) & 0x70) | (_efd & 0xc0) | (_bbc & 0x700)
	var _cfeb int
	for _daf := 0; _daf < _eda; _daf = _cfeb {
		var (
			_acf int
			_ebf int
		)
		_cfeb = _daf + 8
		if _acf = _ce - _daf; _acf > 8 {
			_acf = 8
		}
		_ccd := _cfeb < _ce
		_agd := _cfeb < _afga.ReferenceBitmap.Width
		_dbe := _gad + 1
		if _egg > 0 {
			_efde = 0
			if _ccd {
				_efde, _eagb = _afga.RegionBitmap.GetByte(_gaa - _edfd + 1)
				if _eagb != nil {
					return
				}
			}
			_aed = (_aed << 8) | int(_efde)
		}
		if _efg > 0 && _efg <= _afga.ReferenceBitmap.Height {
			var _aa int
			if _agd {
				_efde, _eagb = _afga.ReferenceBitmap.GetByte(_eag - _agg + _dbe)
				if _eagb != nil {
					return
				}
				_aa = int(_efde) << 2
			}
			_bbc = (_bbc << 8) | _aa
		}
		if _efg >= 0 && _efg < _afga.ReferenceBitmap.Height {
			_efde = 0
			if _agd {
				_efde, _eagb = _afga.ReferenceBitmap.GetByte(_eag + _dbe)
				if _eagb != nil {
					return
				}
			}
			_efd = (_efd << 8) | int(_efde)
		}
		if _efg > -2 && _efg < (_afga.ReferenceBitmap.Height-1) {
			_efde = 0
			if _agd {
				_efde, _eagb = _afga.ReferenceBitmap.GetByte(_eag + _agg + _dbe)
				if _eagb != nil {
					return
				}
			}
			_gga = (_gga << 8) | int(_efde)
		}
		for _bgg := 0; _bgg < _acf; _bgg++ {
			var _feb int
			_ggfc := (_age >> 4) & 0x1ff
			switch _ggfc {
			case 0x1ff:
				_feb = 1
			case 0x00:
				_feb = 0
			default:
				_afga._gbf.SetIndex(int32(_afe))
				_feb, _eagb = _afga._cc.DecodeBit(_afga._gbf)
				if _eagb != nil {
					return
				}
			}
			_cff := uint(7 - _bgg)
			_ebf |= _feb << _cff
			_afe = ((_afe & 0x0d6) << 1) | _feb | (_aed>>_cff+5)&0x002 | ((_gga>>_cff + 2) & 0x010) | ((_efd >> _cff) & 0x040) | ((_bbc >> _cff) & 0x200)
			_age = ((_age & 0xdb) << 1) | ((_gga>>_cff + 2) & 0x010) | ((_efd >> _cff) & 0x080) | ((_bbc >> _cff) & 0x400)
		}
		_eagb = _afga.RegionBitmap.SetByte(_gaa, byte(_ebf))
		if _eagb != nil {
			return
		}
		_gaa++
		_eag++
	}
	return nil
}

func (_cgee *Header) String() string {
	_bfbg := &_cf.Builder{}
	_bfbg.WriteString("\u000a[\u0053E\u0047\u004d\u0045\u004e\u0054-\u0048\u0045A\u0044\u0045\u0052\u005d\u000a")
	_bfbg.WriteString(_f.Sprintf("\t\u002d\u0020\u0053\u0065gm\u0065n\u0074\u004e\u0075\u006d\u0062e\u0072\u003a\u0020\u0025\u0076\u000a", _cgee.SegmentNumber))
	_bfbg.WriteString(_f.Sprintf("\u0009\u002d\u0020T\u0079\u0070\u0065\u003a\u0020\u0025\u0076\u000a", _cgee.Type))
	_bfbg.WriteString(_f.Sprintf("\u0009-\u0020R\u0065\u0074\u0061\u0069\u006eF\u006c\u0061g\u003a\u0020\u0025\u0076\u000a", _cgee.RetainFlag))
	_bfbg.WriteString(_f.Sprintf("\u0009\u002d\u0020Pa\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _cgee.PageAssociation))
	_bfbg.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0050\u0061\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061\u0074i\u006fn\u0046\u0069\u0065\u006c\u0064\u0053\u0069\u007a\u0065\u003a\u0020\u0025\u0076\u000a", _cgee.PageAssociationFieldSize))
	_bfbg.WriteString("\u0009-\u0020R\u0054\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u0053\u003a\u000a")
	for _, _ccbbb := range _cgee.RTSNumbers {
		_bfbg.WriteString(_f.Sprintf("\u0009\t\u002d\u0020\u0025\u0064\u000a", _ccbbb))
	}
	_bfbg.WriteString(_f.Sprintf("\t\u002d \u0048\u0065\u0061\u0064\u0065\u0072\u004c\u0065n\u0067\u0074\u0068\u003a %\u0076\u000a", _cgee.HeaderLength))
	_bfbg.WriteString(_f.Sprintf("\u0009-\u0020\u0053\u0065\u0067m\u0065\u006e\u0074\u0044\u0061t\u0061L\u0065n\u0067\u0074\u0068\u003a\u0020\u0025\u0076\n", _cgee.SegmentDataLength))
	_bfbg.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074D\u0061\u0074\u0061\u0053\u0074\u0061\u0072t\u004f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0076\u000a", _cgee.SegmentDataStartOffset))
	return _bfbg.String()
}

func (_ddaf *SymbolDictionary) encodeSymbols(_acgb _e.BinaryWriter) (_bebbb int, _egeg error) {
	const _efad = "\u0065\u006e\u0063o\u0064\u0065\u0053\u0079\u006d\u0062\u006f\u006c"
	_ceaa := _bb.New()
	_ceaa.Init()
	_bfbgc, _egeg := _ddaf._efbc.SelectByIndexes(_ddaf._decd)
	if _egeg != nil {
		return 0, _da.Wrap(_egeg, _efad, "\u0069n\u0069\u0074\u0069\u0061\u006c")
	}
	_ddde := map[*_ea.Bitmap]int{}
	for _eadg, _gbaf := range _bfbgc.Values {
		_ddde[_gbaf] = _eadg
	}
	_bfbgc.SortByHeight()
	var _faae, _fccc int
	_caef, _egeg := _bfbgc.GroupByHeight()
	if _egeg != nil {
		return 0, _da.Wrap(_egeg, _efad, "")
	}
	for _, _eagf := range _caef.Values {
		_cegb := _eagf.Values[0].Height
		_aaea := _cegb - _faae
		if _egeg = _ceaa.EncodeInteger(_bb.IADH, _aaea); _egeg != nil {
			return 0, _da.Wrapf(_egeg, _efad, "\u0049\u0041\u0044\u0048\u0020\u0066\u006f\u0072\u0020\u0064\u0068\u003a \u0027\u0025\u0064\u0027", _aaea)
		}
		_faae = _cegb
		_dcgag, _cacag := _eagf.GroupByWidth()
		if _cacag != nil {
			return 0, _da.Wrapf(_cacag, _efad, "\u0068\u0065\u0069g\u0068\u0074\u003a\u0020\u0027\u0025\u0064\u0027", _cegb)
		}
		var _gdcf int
		for _, _afggg := range _dcgag.Values {
			for _, _ffab := range _afggg.Values {
				_gadc := _ffab.Width
				_cagc := _gadc - _gdcf
				if _cacag = _ceaa.EncodeInteger(_bb.IADW, _cagc); _cacag != nil {
					return 0, _da.Wrapf(_cacag, _efad, "\u0049\u0041\u0044\u0057\u0020\u0066\u006f\u0072\u0020\u0064\u0077\u003a \u0027\u0025\u0064\u0027", _cagc)
				}
				_gdcf += _cagc
				if _cacag = _ceaa.EncodeBitmap(_ffab, false); _cacag != nil {
					return 0, _da.Wrapf(_cacag, _efad, "H\u0065i\u0067\u0068\u0074\u003a\u0020\u0025\u0064\u0020W\u0069\u0064\u0074\u0068: \u0025\u0064", _cegb, _gadc)
				}
				_aacf := _ddde[_ffab]
				_ddaf._fcae[_aacf] = _fccc
				_fccc++
			}
		}
		if _cacag = _ceaa.EncodeOOB(_bb.IADW); _cacag != nil {
			return 0, _da.Wrap(_cacag, _efad, "\u0049\u0041\u0044\u0057")
		}
	}
	if _egeg = _ceaa.EncodeInteger(_bb.IAEX, 0); _egeg != nil {
		return 0, _da.Wrap(_egeg, _efad, "\u0065\u0078p\u006f\u0072\u0074e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	if _egeg = _ceaa.EncodeInteger(_bb.IAEX, len(_ddaf._decd)); _egeg != nil {
		return 0, _da.Wrap(_egeg, _efad, "\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0073\u0079m\u0062\u006f\u006c\u0073")
	}
	_ceaa.Final()
	_bbda, _egeg := _ceaa.WriteTo(_acgb)
	if _egeg != nil {
		return 0, _da.Wrap(_egeg, _efad, "\u0077\u0072i\u0074\u0069\u006e\u0067 \u0065\u006ec\u006f\u0064\u0065\u0072\u0020\u0063\u006f\u006et\u0065\u0078\u0074\u0020\u0074\u006f\u0020\u0027\u0077\u0027\u0020\u0077r\u0069\u0074\u0065\u0072")
	}
	return int(_bbda), nil
}

func (_febg *GenericRefinementRegion) overrideAtTemplate0(_bf, _eac, _fgfc, _gae, _cfbd int) int {
	if _febg._ge[0] {
		_bf &= 0xfff7
		if _febg.GrAtY[0] == 0 && int(_febg.GrAtX[0]) >= -_cfbd {
			_bf |= (_gae >> uint(7-(_cfbd+int(_febg.GrAtX[0]))) & 0x1) << 3
		} else {
			_bf |= _febg.getPixel(_febg.RegionBitmap, _eac+int(_febg.GrAtX[0]), _fgfc+int(_febg.GrAtY[0])) << 3
		}
	}
	if _febg._ge[1] {
		_bf &= 0xefff
		if _febg.GrAtY[1] == 0 && int(_febg.GrAtX[1]) >= -_cfbd {
			_bf |= (_gae >> uint(7-(_cfbd+int(_febg.GrAtX[1]))) & 0x1) << 12
		} else {
			_bf |= _febg.getPixel(_febg.ReferenceBitmap, _eac+int(_febg.GrAtX[1]), _fgfc+int(_febg.GrAtY[1]))
		}
	}
	return _bf
}

func (_abfb *Header) pageSize() uint {
	if _abfb.PageAssociation <= 255 {
		return 1
	}
	return 4
}

func (_cadb *GenericRegion) computeSegmentDataStructure() error {
	_cadb.DataOffset = _cadb._fec.AbsolutePosition()
	_cadb.DataHeaderLength = _cadb.DataOffset - _cadb.DataHeaderOffset
	_cadb.DataLength = int64(_cadb._fec.AbsoluteLength()) - _cadb.DataHeaderLength
	return nil
}

func (_bceg *PageInformationSegment) parseHeader() (_bcgg error) {
	_eg.Log.Trace("\u005b\u0050\u0061\u0067\u0065I\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065\u0067m\u0065\u006e\u0074\u005d\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065\u0061\u0064\u0065\u0072\u002e\u002e\u002e")
	defer func() {
		_baebb := "[\u0050\u0061\u0067\u0065\u0049\u006e\u0066\u006f\u0072m\u0061\u0074\u0069\u006f\u006e\u0053\u0065gm\u0065\u006e\u0074\u005d \u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065ad\u0065\u0072 \u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064"
		if _bcgg != nil {
			_baebb += "\u0020\u0077\u0069t\u0068\u0020\u0065\u0072\u0072\u006f\u0072\u0020" + _bcgg.Error()
		} else {
			_baebb += "\u0020\u0073\u0075\u0063\u0063\u0065\u0073\u0073\u0066\u0075\u006c\u006c\u0079"
		}
		_eg.Log.Trace(_baebb)
	}()
	if _bcgg = _bceg.readWidthAndHeight(); _bcgg != nil {
		return _bcgg
	}
	if _bcgg = _bceg.readResolution(); _bcgg != nil {
		return _bcgg
	}
	_, _bcgg = _bceg._dfdf.ReadBit()
	if _bcgg != nil {
		return _bcgg
	}
	if _bcgg = _bceg.readCombinationOperatorOverrideAllowed(); _bcgg != nil {
		return _bcgg
	}
	if _bcgg = _bceg.readRequiresAuxiliaryBuffer(); _bcgg != nil {
		return _bcgg
	}
	if _bcgg = _bceg.readCombinationOperator(); _bcgg != nil {
		return _bcgg
	}
	if _bcgg = _bceg.readDefaultPixelValue(); _bcgg != nil {
		return _bcgg
	}
	if _bcgg = _bceg.readContainsRefinement(); _bcgg != nil {
		return _bcgg
	}
	if _bcgg = _bceg.readIsLossless(); _bcgg != nil {
		return _bcgg
	}
	if _bcgg = _bceg.readIsStriped(); _bcgg != nil {
		return _bcgg
	}
	if _bcgg = _bceg.readMaxStripeSize(); _bcgg != nil {
		return _bcgg
	}
	if _bcgg = _bceg.checkInput(); _bcgg != nil {
		return _bcgg
	}
	_eg.Log.Trace("\u0025\u0073", _bceg)
	return nil
}

func (_deca *SymbolDictionary) decodeDirectlyThroughGenericRegion(_edb, _dafeb uint32) error {
	if _deca._dadbb == nil {
		_deca._dadbb = NewGenericRegion(_deca._abege)
	}
	_deca._dadbb.setParametersWithAt(false, byte(_deca.SdTemplate), false, false, _deca.SdATX, _deca.SdATY, _edb, _dafeb, _deca._ccgbe, _deca._agbbf)
	return _deca.addSymbol(_deca._dadbb)
}

func (_dbag *TextRegion) Init(header *Header, r *_e.Reader) error {
	_dbag.Header = header
	_dbag._gdbbd = r
	_dbag.RegionInfo = NewRegionSegment(_dbag._gdbbd)
	return _dbag.parseHeader()
}

func (_ggg *GenericRefinementRegion) getGrReference() (*_ea.Bitmap, error) {
	segments := _ggg._bg.RTSegments
	if len(segments) == 0 {
		return nil, _g.New("\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0065\u0078is\u0074\u0073")
	}
	_geb, _eb := segments[0].GetSegmentData()
	if _eb != nil {
		return nil, _eb
	}
	_ac, _dce := _geb.(Regioner)
	if !_dce {
		return nil, _f.Errorf("\u0072\u0065\u0066\u0065\u0072r\u0065\u0064\u0020\u0074\u006f\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0052\u0065\u0067\u0069\u006f\u006e\u0065\u0072\u003a\u0020\u0025\u0054", _geb)
	}
	return _ac.GetRegionBitmap()
}

func (_gdcg *SymbolDictionary) huffDecodeRefAggNInst() (int64, error) {
	if !_gdcg.SdHuffAggInstanceSelection {
		_fefc, _eadcg := _cg.GetStandardTable(1)
		if _eadcg != nil {
			return 0, _eadcg
		}
		return _fefc.Decode(_gdcg._abege)
	}
	if _gdcg._febf == nil {
		var (
			_fdcc  int
			_eggdd error
		)
		if _gdcg.SdHuffDecodeHeightSelection == 3 {
			_fdcc++
		}
		if _gdcg.SdHuffDecodeWidthSelection == 3 {
			_fdcc++
		}
		if _gdcg.SdHuffBMSizeSelection == 3 {
			_fdcc++
		}
		_gdcg._febf, _eggdd = _gdcg.getUserTable(_fdcc)
		if _eggdd != nil {
			return 0, _eggdd
		}
	}
	return _gdcg._febf.Decode(_gdcg._abege)
}

func (_ffde *PageInformationSegment) readMaxStripeSize() error {
	_fabf, _ffgef := _ffde._dfdf.ReadBits(15)
	if _ffgef != nil {
		return _ffgef
	}
	_ffde.MaxStripeSize = uint16(_fabf & _b.MaxUint16)
	return nil
}

type template1 struct{}

func (_edfgb *SymbolDictionary) setCodingStatistics() error {
	if _edfgb._edaa == nil {
		_edfgb._edaa = _ag.NewStats(512, 1)
	}
	if _edfgb._abdd == nil {
		_edfgb._abdd = _ag.NewStats(512, 1)
	}
	if _edfgb._eaecb == nil {
		_edfgb._eaecb = _ag.NewStats(512, 1)
	}
	if _edfgb._gaed == nil {
		_edfgb._gaed = _ag.NewStats(512, 1)
	}
	if _edfgb._dbcca == nil {
		_edfgb._dbcca = _ag.NewStats(512, 1)
	}
	if _edfgb.UseRefinementAggregation && _edfgb._gaac == nil {
		_edfgb._gaac = _ag.NewStats(1<<uint(_edfgb._cbe), 1)
		_edfgb._ddea = _ag.NewStats(512, 1)
		_edfgb._fad = _ag.NewStats(512, 1)
	}
	if _edfgb._ccgbe == nil {
		_edfgb._ccgbe = _ag.NewStats(65536, 1)
	}
	if _edfgb._agbbf == nil {
		var _fbg error
		_edfgb._agbbf, _fbg = _ag.New(_edfgb._abege)
		if _fbg != nil {
			return _fbg
		}
	}
	return nil
}

func (_dgb *HalftoneRegion) Init(hd *Header, r *_e.Reader) error {
	_dgb._bfbe = r
	_dgb._aadb = hd
	_dgb.RegionSegment = NewRegionSegment(r)
	return _dgb.parseHeader()
}

func (_gffg *SymbolDictionary) decodeHeightClassDeltaHeightWithHuffman() (int64, error) {
	switch _gffg.SdHuffDecodeHeightSelection {
	case 0:
		_dgeb, _bfa := _cg.GetStandardTable(4)
		if _bfa != nil {
			return 0, _bfa
		}
		return _dgeb.Decode(_gffg._abege)
	case 1:
		_dffd, _cfdd := _cg.GetStandardTable(5)
		if _cfdd != nil {
			return 0, _cfdd
		}
		return _dffd.Decode(_gffg._abege)
	case 3:
		if _gffg._gefb == nil {
			_ceafe, _eecgc := _cg.GetStandardTable(0)
			if _eecgc != nil {
				return 0, _eecgc
			}
			_gffg._gefb = _ceafe
		}
		return _gffg._gefb.Decode(_gffg._abege)
	}
	return 0, nil
}

type Documenter interface {
	GetPage(int) (Pager, error)
	GetGlobalSegment(int) (*Header, error)
}
type EncodeInitializer interface {
	InitEncode()
}

func (_afcgc *TextRegion) readUseRefinement() error {
	if !_afcgc.UseRefinement || _afcgc.SbrTemplate != 0 {
		return nil
	}
	var (
		_eacb byte
		_gagf error
	)
	_afcgc.SbrATX = make([]int8, 2)
	_afcgc.SbrATY = make([]int8, 2)
	_eacb, _gagf = _afcgc._gdbbd.ReadByte()
	if _gagf != nil {
		return _gagf
	}
	_afcgc.SbrATX[0] = int8(_eacb)
	_eacb, _gagf = _afcgc._gdbbd.ReadByte()
	if _gagf != nil {
		return _gagf
	}
	_afcgc.SbrATY[0] = int8(_eacb)
	_eacb, _gagf = _afcgc._gdbbd.ReadByte()
	if _gagf != nil {
		return _gagf
	}
	_afcgc.SbrATX[1] = int8(_eacb)
	_eacb, _gagf = _afcgc._gdbbd.ReadByte()
	if _gagf != nil {
		return _gagf
	}
	_afcgc.SbrATY[1] = int8(_eacb)
	return nil
}

var _ _cg.BasicTabler = &TableSegment{}

func (_gagd *Header) writeSegmentNumber(_ebc _e.BinaryWriter) (_cbgb int, _gbgb error) {
	_dcgc := make([]byte, 4)
	_ab.BigEndian.PutUint32(_dcgc, _gagd.SegmentNumber)
	if _cbgb, _gbgb = _ebc.Write(_dcgc); _gbgb != nil {
		return 0, _da.Wrap(_gbgb, "\u0048e\u0061\u0064\u0065\u0072.\u0077\u0072\u0069\u0074\u0065S\u0065g\u006de\u006e\u0074\u004e\u0075\u006d\u0062\u0065r", "")
	}
	return _cbgb, nil
}

func (_baa *GenericRefinementRegion) getPixel(_cfff *_ea.Bitmap, _bga, _cce int) int {
	if _bga < 0 || _bga >= _cfff.Width {
		return 0
	}
	if _cce < 0 || _cce >= _cfff.Height {
		return 0
	}
	if _cfff.GetPixel(_bga, _cce) {
		return 1
	}
	return 0
}

var (
	_ Regioner  = &TextRegion{}
	_ Segmenter = &TextRegion{}
)

func (_afgcb *Header) GetSegmentData() (Segmenter, error) {
	var _abfd Segmenter
	if _afgcb.SegmentData != nil {
		_abfd = _afgcb.SegmentData
	}
	if _abfd == nil {
		_ecc, _gcef := _egd[_afgcb.Type]
		if !_gcef {
			return nil, _f.Errorf("\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0073\u002f\u0020\u0025\u0064\u0020\u0063\u0072e\u0061t\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e\u0020", _afgcb.Type, _afgcb.Type)
		}
		_abfd = _ecc()
		_eg.Log.Trace("\u005b\u0053E\u0047\u004d\u0045\u004e\u0054-\u0048\u0045\u0041\u0044\u0045R\u005d\u005b\u0023\u0025\u0064\u005d\u0020\u0047\u0065\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u0074\u0020\u004f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0030\u0034\u0058", _afgcb.SegmentNumber, _afgcb.SegmentDataStartOffset)
		_fbdcg, _dbcc := _afgcb.subInputReader()
		if _dbcc != nil {
			return nil, _dbcc
		}
		if _fcaf := _abfd.Init(_afgcb, _fbdcg); _fcaf != nil {
			_eg.Log.Debug("\u0049\u006e\u0069\u0074 \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076 \u0066o\u0072\u0020\u0074\u0079\u0070\u0065\u003a \u0025\u0054", _fcaf, _abfd)
			return nil, _fcaf
		}
		_afgcb.SegmentData = _abfd
	}
	return _abfd, nil
}

func (_ddbf *GenericRegion) updateOverrideFlags() error {
	const _bbd = "\u0075\u0070\u0064\u0061te\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _ddbf.GBAtX == nil || _ddbf.GBAtY == nil {
		return nil
	}
	if len(_ddbf.GBAtX) != len(_ddbf.GBAtY) {
		return _da.Errorf(_bbd, "i\u006eco\u0073i\u0073t\u0065\u006e\u0074\u0020\u0041T\u0020\u0070\u0069x\u0065\u006c\u002e\u0020\u0041m\u006f\u0075\u006et\u0020\u006f\u0066\u0020\u0027\u0078\u0027\u0020\u0070\u0069\u0078e\u006c\u0073\u003a %d\u002c\u0020\u0041\u006d\u006f\u0075n\u0074\u0020\u006f\u0066\u0020\u0027\u0079\u0027\u0020\u0070\u0069\u0078e\u006cs\u003a\u0020\u0025\u0064", len(_ddbf.GBAtX), len(_ddbf.GBAtY))
	}
	_ddbf.GBAtOverride = make([]bool, len(_ddbf.GBAtX))
	switch _ddbf.GBTemplate {
	case 0:
		if !_ddbf.UseExtTemplates {
			if _ddbf.GBAtX[0] != 3 || _ddbf.GBAtY[0] != -1 {
				_ddbf.setOverrideFlag(0)
			}
			if _ddbf.GBAtX[1] != -3 || _ddbf.GBAtY[1] != -1 {
				_ddbf.setOverrideFlag(1)
			}
			if _ddbf.GBAtX[2] != 2 || _ddbf.GBAtY[2] != -2 {
				_ddbf.setOverrideFlag(2)
			}
			if _ddbf.GBAtX[3] != -2 || _ddbf.GBAtY[3] != -2 {
				_ddbf.setOverrideFlag(3)
			}
		} else {
			if _ddbf.GBAtX[0] != -2 || _ddbf.GBAtY[0] != 0 {
				_ddbf.setOverrideFlag(0)
			}
			if _ddbf.GBAtX[1] != 0 || _ddbf.GBAtY[1] != -2 {
				_ddbf.setOverrideFlag(1)
			}
			if _ddbf.GBAtX[2] != -2 || _ddbf.GBAtY[2] != -1 {
				_ddbf.setOverrideFlag(2)
			}
			if _ddbf.GBAtX[3] != -1 || _ddbf.GBAtY[3] != -2 {
				_ddbf.setOverrideFlag(3)
			}
			if _ddbf.GBAtX[4] != 1 || _ddbf.GBAtY[4] != -2 {
				_ddbf.setOverrideFlag(4)
			}
			if _ddbf.GBAtX[5] != 2 || _ddbf.GBAtY[5] != -1 {
				_ddbf.setOverrideFlag(5)
			}
			if _ddbf.GBAtX[6] != -3 || _ddbf.GBAtY[6] != 0 {
				_ddbf.setOverrideFlag(6)
			}
			if _ddbf.GBAtX[7] != -4 || _ddbf.GBAtY[7] != 0 {
				_ddbf.setOverrideFlag(7)
			}
			if _ddbf.GBAtX[8] != 2 || _ddbf.GBAtY[8] != -2 {
				_ddbf.setOverrideFlag(8)
			}
			if _ddbf.GBAtX[9] != 3 || _ddbf.GBAtY[9] != -1 {
				_ddbf.setOverrideFlag(9)
			}
			if _ddbf.GBAtX[10] != -2 || _ddbf.GBAtY[10] != -2 {
				_ddbf.setOverrideFlag(10)
			}
			if _ddbf.GBAtX[11] != -3 || _ddbf.GBAtY[11] != -1 {
				_ddbf.setOverrideFlag(11)
			}
		}
	case 1:
		if _ddbf.GBAtX[0] != 3 || _ddbf.GBAtY[0] != -1 {
			_ddbf.setOverrideFlag(0)
		}
	case 2:
		if _ddbf.GBAtX[0] != 2 || _ddbf.GBAtY[0] != -1 {
			_ddbf.setOverrideFlag(0)
		}
	case 3:
		if _ddbf.GBAtX[0] != 2 || _ddbf.GBAtY[0] != -1 {
			_ddbf.setOverrideFlag(0)
		}
	}
	return nil
}

type Type int

func (_bffa *TextRegion) decodeIds() (int64, error) {
	const _dgce = "\u0064e\u0063\u006f\u0064\u0065\u0049\u0064s"
	if _bffa.IsHuffmanEncoded {
		if _bffa.SbHuffDS == 3 {
			if _bffa._begc == nil {
				_ffbbe := 0
				if _bffa.SbHuffFS == 3 {
					_ffbbe++
				}
				var _dca error
				_bffa._begc, _dca = _bffa.getUserTable(_ffbbe)
				if _dca != nil {
					return 0, _da.Wrap(_dca, _dgce, "")
				}
			}
			return _bffa._begc.Decode(_bffa._gdbbd)
		}
		_gdda, _faaef := _cg.GetStandardTable(8 + int(_bffa.SbHuffDS))
		if _faaef != nil {
			return 0, _da.Wrap(_faaef, _dgce, "")
		}
		return _gdda.Decode(_bffa._gdbbd)
	}
	_dcag, _edgd := _bffa._bggc.DecodeInt(_bffa._fbfg)
	if _edgd != nil {
		return 0, _da.Wrap(_edgd, _dgce, "\u0063\u0078\u0049\u0041\u0044\u0053")
	}
	return int64(_dcag), nil
}

func (_ebd *HalftoneRegion) GetPatterns() ([]*_ea.Bitmap, error) {
	var (
		_dgae []*_ea.Bitmap
		_afae error
	)
	for _, _afgg := range _ebd._aadb.RTSegments {
		var _ccda Segmenter
		_ccda, _afae = _afgg.GetSegmentData()
		if _afae != nil {
			_eg.Log.Debug("\u0047e\u0074\u0053\u0065\u0067m\u0065\u006e\u0074\u0044\u0061t\u0061 \u0066a\u0069\u006c\u0065\u0064\u003a\u0020\u0025v", _afae)
			return nil, _afae
		}
		_gcff, _eggd := _ccda.(*PatternDictionary)
		if !_eggd {
			_afae = _f.Errorf("\u0072e\u006c\u0061t\u0065\u0064\u0020\u0073e\u0067\u006d\u0065n\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0070at\u0074\u0065\u0072n\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u003a \u0025\u0054", _ccda)
			return nil, _afae
		}
		var _gfgg []*_ea.Bitmap
		_gfgg, _afae = _gcff.GetDictionary()
		if _afae != nil {
			_eg.Log.Debug("\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _afae)
			return nil, _afae
		}
		_dgae = append(_dgae, _gfgg...)
	}
	return _dgae, nil
}

type TableSegment struct {
	_agda  *_e.Reader
	_caff  int32
	_cggf  int32
	_gbda  int32
	_aegdd int32
	_ccfd  int32
}

func (_baeed *TextRegion) encodeSymbols(_gefed _e.BinaryWriter) (_fbfd int, _gddf error) {
	const _dgaec = "\u0065\u006e\u0063\u006f\u0064\u0065\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	_baba := make([]byte, 4)
	_ab.BigEndian.PutUint32(_baba, _baeed.NumberOfSymbols)
	if _fbfd, _gddf = _gefed.Write(_baba); _gddf != nil {
		return _fbfd, _da.Wrap(_gddf, _dgaec, "\u004e\u0075\u006dbe\u0072\u004f\u0066\u0053\u0079\u006d\u0062\u006f\u006c\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073")
	}
	_cfcg, _gddf := _ea.NewClassedPoints(_baeed._ecfg, _baeed._bfec)
	if _gddf != nil {
		return 0, _da.Wrap(_gddf, _dgaec, "")
	}
	var _bdbg, _gdgf int
	_bdgbb := _bb.New()
	_bdgbb.Init()
	if _gddf = _bdgbb.EncodeInteger(_bb.IADT, 0); _gddf != nil {
		return _fbfd, _da.Wrap(_gddf, _dgaec, "\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0020\u0044\u0054")
	}
	_cbcf, _gddf := _cfcg.GroupByY()
	if _gddf != nil {
		return 0, _da.Wrap(_gddf, _dgaec, "")
	}
	for _, _face := range _cbcf {
		_gcgd := int(_face.YAtIndex(0))
		_fceb := _gcgd - _bdbg
		if _gddf = _bdgbb.EncodeInteger(_bb.IADT, _fceb); _gddf != nil {
			return _fbfd, _da.Wrap(_gddf, _dgaec, "")
		}
		var _dadfc int
		for _agca, _ceca := range _face.IntSlice {
			switch _agca {
			case 0:
				_fggc := int(_face.XAtIndex(_agca)) - _gdgf
				if _gddf = _bdgbb.EncodeInteger(_bb.IAFS, _fggc); _gddf != nil {
					return _fbfd, _da.Wrap(_gddf, _dgaec, "")
				}
				_gdgf += _fggc
				_dadfc = _gdgf
			default:
				_efeb := int(_face.XAtIndex(_agca)) - _dadfc
				if _gddf = _bdgbb.EncodeInteger(_bb.IADS, _efeb); _gddf != nil {
					return _fbfd, _da.Wrap(_gddf, _dgaec, "")
				}
				_dadfc += _efeb
			}
			_cbcb, _cdfd := _baeed._afegf.Get(_ceca)
			if _cdfd != nil {
				return _fbfd, _da.Wrap(_cdfd, _dgaec, "")
			}
			_bggce, _efgf := _baeed._dffgd[_cbcb]
			if !_efgf {
				_bggce, _efgf = _baeed._edbd[_cbcb]
				if !_efgf {
					return _fbfd, _da.Errorf(_dgaec, "\u0053\u0079\u006d\u0062\u006f\u006c:\u0020\u0027\u0025d\u0027\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064 \u0069\u006e\u0020\u0067\u006cob\u0061\u006c\u0020\u0061\u006e\u0064\u0020\u006c\u006f\u0063\u0061\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0020\u006d\u0061\u0070", _cbcb)
				}
			}
			if _cdfd = _bdgbb.EncodeIAID(_baeed._beffc, _bggce); _cdfd != nil {
				return _fbfd, _da.Wrap(_cdfd, _dgaec, "")
			}
		}
		if _gddf = _bdgbb.EncodeOOB(_bb.IADS); _gddf != nil {
			return _fbfd, _da.Wrap(_gddf, _dgaec, "")
		}
	}
	_bdgbb.Final()
	_dacf, _gddf := _bdgbb.WriteTo(_gefed)
	if _gddf != nil {
		return _fbfd, _da.Wrap(_gddf, _dgaec, "")
	}
	_fbfd += int(_dacf)
	return _fbfd, nil
}

func (_fbaef *TextRegion) InitEncode(globalSymbolsMap, localSymbolsMap map[int]int, comps []int, inLL *_ea.Points, symbols *_ea.Bitmaps, classIDs *_dc.IntSlice, boxes *_ea.Boxes, width, height, symBits int) {
	_fbaef.RegionInfo = &RegionSegment{BitmapWidth: uint32(width), BitmapHeight: uint32(height)}
	_fbaef._dffgd = globalSymbolsMap
	_fbaef._edbd = localSymbolsMap
	_fbaef._bfec = comps
	_fbaef._ecfg = inLL
	_fbaef._bbagb = symbols
	_fbaef._afegf = classIDs
	_fbaef._fagbc = boxes
	_fbaef._beffc = symBits
}

func (_eeba *TextRegion) decodeSymInRefSize() (int64, error) {
	const _bcgabb = "\u0064e\u0063o\u0064\u0065\u0053\u0079\u006dI\u006e\u0052e\u0066\u0053\u0069\u007a\u0065"
	if _eeba.SbHuffRSize == 0 {
		_aegae, _gefa := _cg.GetStandardTable(1)
		if _gefa != nil {
			return 0, _da.Wrap(_gefa, _bcgabb, "")
		}
		return _aegae.Decode(_eeba._gdbbd)
	}
	if _eeba._ecdf == nil {
		var (
			_ddcf  int
			_ebfdc error
		)
		if _eeba.SbHuffFS == 3 {
			_ddcf++
		}
		if _eeba.SbHuffDS == 3 {
			_ddcf++
		}
		if _eeba.SbHuffDT == 3 {
			_ddcf++
		}
		if _eeba.SbHuffRDWidth == 3 {
			_ddcf++
		}
		if _eeba.SbHuffRDHeight == 3 {
			_ddcf++
		}
		if _eeba.SbHuffRDX == 3 {
			_ddcf++
		}
		if _eeba.SbHuffRDY == 3 {
			_ddcf++
		}
		_eeba._ecdf, _ebfdc = _eeba.getUserTable(_ddcf)
		if _ebfdc != nil {
			return 0, _da.Wrap(_ebfdc, _bcgabb, "")
		}
	}
	_egfba, _gdgef := _eeba._ecdf.Decode(_eeba._gdbbd)
	if _gdgef != nil {
		return 0, _da.Wrap(_gdgef, _bcgabb, "")
	}
	return _egfba, nil
}

func (_cegf *SymbolDictionary) readNumberOfNewSymbols() error {
	_gbb, _gaedc := _cegf._abege.ReadBits(32)
	if _gaedc != nil {
		return _gaedc
	}
	_cegf.NumberOfNewSymbols = uint32(_gbb & _b.MaxUint32)
	return nil
}

func (_aggc *TextRegion) computeSymbolCodeLength() error {
	if _aggc.IsHuffmanEncoded {
		return _aggc.symbolIDCodeLengths()
	}
	_aggc._ggggf = int8(_b.Ceil(_b.Log(float64(_aggc.NumberOfSymbols)) / _b.Log(2)))
	return nil
}

const (
	ORandom OrganizationType = iota
	OSequential
)

func (_gdg *HalftoneRegion) computeY(_cgdg, _dead int) int {
	return _gdg.shiftAndFill(int(_gdg.HGridY) + _cgdg*int(_gdg.HRegionX) - _dead*int(_gdg.HRegionY))
}
func (_gdf *GenericRefinementRegion) GetRegionInfo() *RegionSegment { return _gdf.RegionInfo }

type Pager interface {
	GetSegment(int) (*Header, error)
	GetBitmap() (*_ea.Bitmap, error)
}

func (_gdca *SymbolDictionary) decodeHeightClassCollectiveBitmap(_daee int64, _decbe, _afegg uint32) (*_ea.Bitmap, error) {
	if _daee == 0 {
		_abfc := _ea.New(int(_afegg), int(_decbe))
		var (
			_befd  byte
			_dbfgg error
		)
		for _agcf := 0; _agcf < len(_abfc.Data); _agcf++ {
			_befd, _dbfgg = _gdca._abege.ReadByte()
			if _dbfgg != nil {
				return nil, _dbfgg
			}
			if _dbfgg = _abfc.SetByte(_agcf, _befd); _dbfgg != nil {
				return nil, _dbfgg
			}
		}
		return _abfc, nil
	}
	if _gdca._dadbb == nil {
		_gdca._dadbb = NewGenericRegion(_gdca._abege)
	}
	_gdca._dadbb.setParameters(true, _gdca._abege.AbsolutePosition(), _daee, _decbe, _afegg)
	_eefa, _fdc := _gdca._dadbb.GetRegionBitmap()
	if _fdc != nil {
		return nil, _fdc
	}
	return _eefa, nil
}

func (_ccaf *TextRegion) getSymbols() error {
	if _ccaf.Header.RTSegments != nil {
		return _ccaf.initSymbols()
	}
	return nil
}

type SymbolDictionary struct {
	_abege                      *_e.Reader
	SdrTemplate                 int8
	SdTemplate                  int8
	_cba                        bool
	_bdgf                       bool
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
	_afag                       uint32
	_dcgb                       []*_ea.Bitmap
	_addfb                      uint32
	_abed                       []*_ea.Bitmap
	_gefb                       _cg.Tabler
	_aceb                       _cg.Tabler
	_gdfd                       _cg.Tabler
	_febf                       _cg.Tabler
	_cfbdb                      []*_ea.Bitmap
	_ffbb                       []*_ea.Bitmap
	_agbbf                      *_ag.Decoder
	_dddd                       *TextRegion
	_dadbb                      *GenericRegion
	_eec                        *GenericRefinementRegion
	_ccgbe                      *_ag.DecoderStats
	_abdd                       *_ag.DecoderStats
	_eaecb                      *_ag.DecoderStats
	_gaed                       *_ag.DecoderStats
	_dbcca                      *_ag.DecoderStats
	_ddea                       *_ag.DecoderStats
	_fad                        *_ag.DecoderStats
	_edaa                       *_ag.DecoderStats
	_gaac                       *_ag.DecoderStats
	_cbe                        int8
	_efbc                       *_ea.Bitmaps
	_decd                       []int
	_fcae                       map[int]int
	_ecffe                      bool
}

func (_fga *HalftoneRegion) computeGrayScalePlanes(_cggde []*_ea.Bitmap, _ccgb int) ([][]int, error) {
	_bfg := make([][]int, _fga.HGridHeight)
	for _bbcfb := 0; _bbcfb < len(_bfg); _bbcfb++ {
		_bfg[_bbcfb] = make([]int, _fga.HGridWidth)
	}
	for _bagg := 0; _bagg < int(_fga.HGridHeight); _bagg++ {
		for _aeca := 0; _aeca < int(_fga.HGridWidth); _aeca += 8 {
			var _gdef int
			if _ffce := int(_fga.HGridWidth) - _aeca; _ffce > 8 {
				_gdef = 8
			} else {
				_gdef = _ffce
			}
			_bcad := _cggde[0].GetByteIndex(_aeca, _bagg)
			for _dbeg := 0; _dbeg < _gdef; _dbeg++ {
				_fcf := _dbeg + _aeca
				_bfg[_bagg][_fcf] = 0
				for _deb := 0; _deb < _ccgb; _deb++ {
					_abg, _deba := _cggde[_deb].GetByte(_bcad)
					if _deba != nil {
						return nil, _deba
					}
					_eceb := _abg >> uint(7-_fcf&7)
					_dea := _eceb & 1
					_fag := 1 << uint(_deb)
					_gfgc := int(_dea) * _fag
					_bfg[_bagg][_fcf] += _gfgc
				}
			}
		}
	}
	return _bfg, nil
}

func (_bacdb *PageInformationSegment) String() string {
	_cdd := &_cf.Builder{}
	_cdd.WriteString("\u000a\u005b\u0050\u0041G\u0045\u002d\u0049\u004e\u0046\u004f\u0052\u004d\u0041\u0054I\u004fN\u002d\u0053\u0045\u0047\u004d\u0045\u004eT\u005d\u000a")
	_cdd.WriteString(_f.Sprintf("\u0009\u002d \u0042\u004d\u0048e\u0069\u0067\u0068\u0074\u003a\u0020\u0025\u0064\u000a", _bacdb.PageBMHeight))
	_cdd.WriteString(_f.Sprintf("\u0009-\u0020B\u004d\u0057\u0069\u0064\u0074\u0068\u003a\u0020\u0025\u0064\u000a", _bacdb.PageBMWidth))
	_cdd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0052es\u006f\u006c\u0075\u0074\u0069\u006f\u006e\u0058\u003a\u0020\u0025\u0064\u000a", _bacdb.ResolutionX))
	_cdd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0052es\u006f\u006c\u0075\u0074\u0069\u006f\u006e\u0059\u003a\u0020\u0025\u0064\u000a", _bacdb.ResolutionY))
	_cdd.WriteString(_f.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065r\u0061\u0074\u006fr\u003a \u0025\u0073\u000a", _bacdb._ccde))
	_cdd.WriteString(_f.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006eO\u0070\u0065\u0072\u0061\u0074\u006f\u0072O\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u003a\u0020\u0025v\u000a", _bacdb._efgad))
	_cdd.WriteString(_f.Sprintf("\u0009-\u0020I\u0073\u004c\u006f\u0073\u0073l\u0065\u0073s\u003a\u0020\u0025\u0076\u000a", _bacdb.IsLossless))
	_cdd.WriteString(_f.Sprintf("\u0009\u002d\u0020R\u0065\u0071\u0075\u0069r\u0065\u0073\u0041\u0075\u0078\u0069\u006ci\u0061\u0072\u0079\u0042\u0075\u0066\u0066\u0065\u0072\u003a\u0020\u0025\u0076\u000a", _bacdb._dcfb))
	_cdd.WriteString(_f.Sprintf("\u0009\u002d\u0020M\u0069\u0067\u0068\u0074C\u006f\u006e\u0074\u0061\u0069\u006e\u0052e\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0076\u000a", _bacdb._fagf))
	_cdd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0049\u0073\u0053\u0074\u0072\u0069\u0070\u0065\u0064:\u0020\u0025\u0076\u000a", _bacdb.IsStripe))
	_cdd.WriteString(_f.Sprintf("\t\u002d\u0020\u004d\u0061xS\u0074r\u0069\u0070\u0065\u0053\u0069z\u0065\u003a\u0020\u0025\u0076\u000a", _bacdb.MaxStripeSize))
	return _cdd.String()
}

func (_eadc *PageInformationSegment) readCombinationOperatorOverrideAllowed() error {
	_afaa, _cbc := _eadc._dfdf.ReadBit()
	if _cbc != nil {
		return _cbc
	}
	if _afaa == 1 {
		_eadc._efgad = true
	}
	return nil
}

type EndOfStripe struct {
	_agb *_e.Reader
	_fb  int
}

func (_gfef *SymbolDictionary) encodeNumSyms(_ebba _e.BinaryWriter) (_gagde int, _ffbdb error) {
	const _bddb = "\u0065\u006e\u0063\u006f\u0064\u0065\u004e\u0075\u006d\u0053\u0079\u006d\u0073"
	_cgdd := make([]byte, 4)
	_ab.BigEndian.PutUint32(_cgdd, _gfef.NumberOfExportedSymbols)
	if _gagde, _ffbdb = _ebba.Write(_cgdd); _ffbdb != nil {
		return _gagde, _da.Wrap(_ffbdb, _bddb, "\u0065\u0078p\u006f\u0072\u0074e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	_ab.BigEndian.PutUint32(_cgdd, _gfef.NumberOfNewSymbols)
	_caaf, _ffbdb := _ebba.Write(_cgdd)
	if _ffbdb != nil {
		return _gagde, _da.Wrap(_ffbdb, _bddb, "n\u0065\u0077\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	return _gagde + _caaf, nil
}

func (_gbaff *TextRegion) setParameters(_aeae *_ag.Decoder, _dbbg, _fgeb bool, _egaf, _adgca uint32, _afed uint32, _dgfd int8, _fbceb uint32, _ggge int8, _accf _ea.CombinationOperator, _becg int8, _cbcd int16, _bdcgg, _fdgad, _gcge, _ggdb, _gcbb, _efff, _fdgf, _abga, _cgff, _ffaf int8, _fcdg, _cacac []int8, _fdbb []*_ea.Bitmap, _fgaa int8) {
	_gbaff._bggc = _aeae
	_gbaff.IsHuffmanEncoded = _dbbg
	_gbaff.UseRefinement = _fgeb
	_gbaff.RegionInfo.BitmapWidth = _egaf
	_gbaff.RegionInfo.BitmapHeight = _adgca
	_gbaff.NumberOfSymbolInstances = _afed
	_gbaff.SbStrips = _dgfd
	_gbaff.NumberOfSymbols = _fbceb
	_gbaff.DefaultPixel = _ggge
	_gbaff.CombinationOperator = _accf
	_gbaff.IsTransposed = _becg
	_gbaff.ReferenceCorner = _cbcd
	_gbaff.SbDsOffset = _bdcgg
	_gbaff.SbHuffFS = _fdgad
	_gbaff.SbHuffDS = _gcge
	_gbaff.SbHuffDT = _ggdb
	_gbaff.SbHuffRDWidth = _gcbb
	_gbaff.SbHuffRDHeight = _efff
	_gbaff.SbHuffRSize = _cgff
	_gbaff.SbHuffRDX = _fdgf
	_gbaff.SbHuffRDY = _abga
	_gbaff.SbrTemplate = _ffaf
	_gbaff.SbrATX = _fcdg
	_gbaff.SbrATY = _cacac
	_gbaff.Symbols = _fdbb
	_gbaff._ggggf = _fgaa
}

type RegionSegment struct {
	_dfbgc             *_e.Reader
	BitmapWidth        uint32
	BitmapHeight       uint32
	XLocation          uint32
	YLocation          uint32
	CombinaionOperator _ea.CombinationOperator
}

func (_bbeb *RegionSegment) parseHeader() error {
	const _fbce = "p\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_eg.Log.Trace("\u005b\u0052\u0045\u0047I\u004f\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002dH\u0045A\u0044\u0045\u0052\u005d\u0020\u0042\u0065g\u0069\u006e")
	defer func() {
		_eg.Log.Trace("\u005b\u0052\u0045G\u0049\u004f\u004e\u005d[\u0050\u0041\u0052\u0053\u0045\u002d\u0048E\u0041\u0044\u0045\u0052\u005d\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
	}()
	_feae, _gaag := _bbeb._dfbgc.ReadBits(32)
	if _gaag != nil {
		return _da.Wrap(_gaag, _fbce, "\u0077\u0069\u0064t\u0068")
	}
	_bbeb.BitmapWidth = uint32(_feae & _b.MaxUint32)
	_feae, _gaag = _bbeb._dfbgc.ReadBits(32)
	if _gaag != nil {
		return _da.Wrap(_gaag, _fbce, "\u0068\u0065\u0069\u0067\u0068\u0074")
	}
	_bbeb.BitmapHeight = uint32(_feae & _b.MaxUint32)
	_feae, _gaag = _bbeb._dfbgc.ReadBits(32)
	if _gaag != nil {
		return _da.Wrap(_gaag, _fbce, "\u0078\u0020\u006c\u006f\u0063\u0061\u0074\u0069\u006f\u006e")
	}
	_bbeb.XLocation = uint32(_feae & _b.MaxUint32)
	_feae, _gaag = _bbeb._dfbgc.ReadBits(32)
	if _gaag != nil {
		return _da.Wrap(_gaag, _fbce, "\u0079\u0020\u006c\u006f\u0063\u0061\u0074\u0069\u006f\u006e")
	}
	_bbeb.YLocation = uint32(_feae & _b.MaxUint32)
	if _, _gaag = _bbeb._dfbgc.ReadBits(5); _gaag != nil {
		return _da.Wrap(_gaag, _fbce, "\u0064i\u0072\u0079\u0020\u0072\u0065\u0061d")
	}
	if _gaag = _bbeb.readCombinationOperator(); _gaag != nil {
		return _da.Wrap(_gaag, _fbce, "c\u006fm\u0062\u0069\u006e\u0061\u0074\u0069\u006f\u006e \u006f\u0070\u0065\u0072at\u006f\u0072")
	}
	return nil
}

func (_afgbe *SymbolDictionary) getUserTable(_dfae int) (_cg.Tabler, error) {
	var _fcea int
	for _, _dgbba := range _afgbe.Header.RTSegments {
		if _dgbba.Type == 53 {
			if _fcea == _dfae {
				_eeab, _begg := _dgbba.GetSegmentData()
				if _begg != nil {
					return nil, _begg
				}
				_dagb := _eeab.(_cg.BasicTabler)
				return _cg.NewEncodedTable(_dagb)
			}
			_fcea++
		}
	}
	return nil, nil
}

type PageInformationSegment struct {
	_dfdf             *_e.Reader
	PageBMHeight      int
	PageBMWidth       int
	ResolutionX       int
	ResolutionY       int
	_efgad            bool
	_ccde             _ea.CombinationOperator
	_dcfb             bool
	DefaultPixelValue uint8
	_fagf             bool
	IsLossless        bool
	IsStripe          bool
	MaxStripeSize     uint16
}

func (_fbcd *TextRegion) readHuffmanFlags() error {
	var (
		_bdgg int
		_fbcg uint64
		_bgeb error
	)
	_, _bgeb = _fbcd._gdbbd.ReadBit()
	if _bgeb != nil {
		return _bgeb
	}
	_bdgg, _bgeb = _fbcd._gdbbd.ReadBit()
	if _bgeb != nil {
		return _bgeb
	}
	_fbcd.SbHuffRSize = int8(_bdgg)
	_fbcg, _bgeb = _fbcd._gdbbd.ReadBits(2)
	if _bgeb != nil {
		return _bgeb
	}
	_fbcd.SbHuffRDY = int8(_fbcg) & 0xf
	_fbcg, _bgeb = _fbcd._gdbbd.ReadBits(2)
	if _bgeb != nil {
		return _bgeb
	}
	_fbcd.SbHuffRDX = int8(_fbcg) & 0xf
	_fbcg, _bgeb = _fbcd._gdbbd.ReadBits(2)
	if _bgeb != nil {
		return _bgeb
	}
	_fbcd.SbHuffRDHeight = int8(_fbcg) & 0xf
	_fbcg, _bgeb = _fbcd._gdbbd.ReadBits(2)
	if _bgeb != nil {
		return _bgeb
	}
	_fbcd.SbHuffRDWidth = int8(_fbcg) & 0xf
	_fbcg, _bgeb = _fbcd._gdbbd.ReadBits(2)
	if _bgeb != nil {
		return _bgeb
	}
	_fbcd.SbHuffDT = int8(_fbcg) & 0xf
	_fbcg, _bgeb = _fbcd._gdbbd.ReadBits(2)
	if _bgeb != nil {
		return _bgeb
	}
	_fbcd.SbHuffDS = int8(_fbcg) & 0xf
	_fbcg, _bgeb = _fbcd._gdbbd.ReadBits(2)
	if _bgeb != nil {
		return _bgeb
	}
	_fbcd.SbHuffFS = int8(_fbcg) & 0xf
	return nil
}

func (_aadgg *SymbolDictionary) readRegionFlags() error {
	var (
		_cgeba uint64
		_gdde  int
	)
	_, _gbae := _aadgg._abege.ReadBits(3)
	if _gbae != nil {
		return _gbae
	}
	_gdde, _gbae = _aadgg._abege.ReadBit()
	if _gbae != nil {
		return _gbae
	}
	_aadgg.SdrTemplate = int8(_gdde)
	_cgeba, _gbae = _aadgg._abege.ReadBits(2)
	if _gbae != nil {
		return _gbae
	}
	_aadgg.SdTemplate = int8(_cgeba & 0xf)
	_gdde, _gbae = _aadgg._abege.ReadBit()
	if _gbae != nil {
		return _gbae
	}
	if _gdde == 1 {
		_aadgg._cba = true
	}
	_gdde, _gbae = _aadgg._abege.ReadBit()
	if _gbae != nil {
		return _gbae
	}
	if _gdde == 1 {
		_aadgg._bdgf = true
	}
	_gdde, _gbae = _aadgg._abege.ReadBit()
	if _gbae != nil {
		return _gbae
	}
	if _gdde == 1 {
		_aadgg.SdHuffAggInstanceSelection = true
	}
	_gdde, _gbae = _aadgg._abege.ReadBit()
	if _gbae != nil {
		return _gbae
	}
	_aadgg.SdHuffBMSizeSelection = int8(_gdde)
	_cgeba, _gbae = _aadgg._abege.ReadBits(2)
	if _gbae != nil {
		return _gbae
	}
	_aadgg.SdHuffDecodeWidthSelection = int8(_cgeba & 0xf)
	_cgeba, _gbae = _aadgg._abege.ReadBits(2)
	if _gbae != nil {
		return _gbae
	}
	_aadgg.SdHuffDecodeHeightSelection = int8(_cgeba & 0xf)
	_gdde, _gbae = _aadgg._abege.ReadBit()
	if _gbae != nil {
		return _gbae
	}
	if _gdde == 1 {
		_aadgg.UseRefinementAggregation = true
	}
	_gdde, _gbae = _aadgg._abege.ReadBit()
	if _gbae != nil {
		return _gbae
	}
	if _gdde == 1 {
		_aadgg.IsHuffmanEncoded = true
	}
	return nil
}

func (_eadf *PatternDictionary) readGrayMax() error {
	_bcfg, _ffbda := _eadf._ccca.ReadBits(32)
	if _ffbda != nil {
		return _ffbda
	}
	_eadf.GrayMax = uint32(_bcfg & _b.MaxUint32)
	return nil
}

func (_bdgb *HalftoneRegion) shiftAndFill(_gbdb int) int {
	_gbdb >>= 8
	if _gbdb < 0 {
		_gdfa := int(_b.Log(float64(_dffg(_gbdb))) / _b.Log(2))
		_fge := 31 - _gdfa
		for _gebc := 1; _gebc < _fge; _gebc++ {
			_gbdb |= 1 << uint(31-_gebc)
		}
	}
	return _gbdb
}

func (_defb *TextRegion) setContexts(_bffag *_ag.DecoderStats, _ggea *_ag.DecoderStats, _fbfc *_ag.DecoderStats, _gcggc *_ag.DecoderStats, _fde *_ag.DecoderStats, _fggg *_ag.DecoderStats, _feggg *_ag.DecoderStats, _cgad *_ag.DecoderStats, _eedb *_ag.DecoderStats, _gabd *_ag.DecoderStats) {
	_defb._ffae = _ggea
	_defb._ccfc = _fbfc
	_defb._fbfg = _gcggc
	_defb._efgab = _fde
	_defb._bgcbb = _feggg
	_defb._dcgcb = _cgad
	_defb._cgfd = _fggg
	_defb._dcfba = _eedb
	_defb._feab = _gabd
	_defb._ecaa = _bffag
}

func (_gfec *GenericRegion) copyLineAbove(_abd int) error {
	_aeeb := _abd * _gfec.Bitmap.RowStride
	_ebaa := _aeeb - _gfec.Bitmap.RowStride
	for _fcge := 0; _fcge < _gfec.Bitmap.RowStride; _fcge++ {
		_adga, _cccg := _gfec.Bitmap.GetByte(_ebaa)
		if _cccg != nil {
			return _cccg
		}
		_ebaa++
		if _cccg = _gfec.Bitmap.SetByte(_aeeb, _adga); _cccg != nil {
			return _cccg
		}
		_aeeb++
	}
	return nil
}

type SegmentEncoder interface {
	Encode(_cgdgf _e.BinaryWriter) (_agae int, _gagcd error)
}

func (_cda *Header) writeReferredToSegments(_ggaab _e.BinaryWriter) (_efec int, _adca error) {
	const _cdae = "\u0077\u0072\u0069te\u0052\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0054\u006f\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073"
	var (
		_bacdc uint16
		_gbec  uint32
	)
	_bcbg := _cda.referenceSize()
	_gac := 1
	_bdfd := make([]byte, _bcbg)
	for _, _faeb := range _cda.RTSNumbers {
		switch _bcbg {
		case 4:
			_gbec = uint32(_faeb)
			_ab.BigEndian.PutUint32(_bdfd, _gbec)
			_gac, _adca = _ggaab.Write(_bdfd)
			if _adca != nil {
				return 0, _da.Wrap(_adca, _cdae, "u\u0069\u006e\u0074\u0033\u0032\u0020\u0073\u0069\u007a\u0065")
			}
		case 2:
			_bacdc = uint16(_faeb)
			_ab.BigEndian.PutUint16(_bdfd, _bacdc)
			_gac, _adca = _ggaab.Write(_bdfd)
			if _adca != nil {
				return 0, _da.Wrap(_adca, _cdae, "\u0075\u0069\u006e\u0074\u0031\u0036")
			}
		default:
			if _adca = _ggaab.WriteByte(byte(_faeb)); _adca != nil {
				return 0, _da.Wrap(_adca, _cdae, "\u0075\u0069\u006et\u0038")
			}
		}
		_efec += _gac
	}
	return _efec, nil
}

func (_gee *SymbolDictionary) decodeRefinedSymbol(_afec, _bdge uint32) error {
	var (
		_acdg        int
		_bddg, _eedf int32
	)
	if _gee.IsHuffmanEncoded {
		_fcaed, _faebd := _gee._abege.ReadBits(byte(_gee._cbe))
		if _faebd != nil {
			return _faebd
		}
		_acdg = int(_fcaed)
		_faaca, _faebd := _cg.GetStandardTable(15)
		if _faebd != nil {
			return _faebd
		}
		_dgbb, _faebd := _faaca.Decode(_gee._abege)
		if _faebd != nil {
			return _faebd
		}
		_bddg = int32(_dgbb)
		_dgbb, _faebd = _faaca.Decode(_gee._abege)
		if _faebd != nil {
			return _faebd
		}
		_eedf = int32(_dgbb)
		_faaca, _faebd = _cg.GetStandardTable(1)
		if _faebd != nil {
			return _faebd
		}
		if _, _faebd = _faaca.Decode(_gee._abege); _faebd != nil {
			return _faebd
		}
		_gee._abege.Align()
	} else {
		_defaa, _agdg := _gee._agbbf.DecodeIAID(uint64(_gee._cbe), _gee._gaac)
		if _agdg != nil {
			return _agdg
		}
		_acdg = int(_defaa)
		_bddg, _agdg = _gee._agbbf.DecodeInt(_gee._ddea)
		if _agdg != nil {
			return _agdg
		}
		_eedf, _agdg = _gee._agbbf.DecodeInt(_gee._fad)
		if _agdg != nil {
			return _agdg
		}
	}
	if _ddgb := _gee.setSymbolsArray(); _ddgb != nil {
		return _ddgb
	}
	_dedda := _gee._ffbb[_acdg]
	if _eccf := _gee.decodeNewSymbols(_afec, _bdge, _dedda, _bddg, _eedf); _eccf != nil {
		return _eccf
	}
	if _gee.IsHuffmanEncoded {
		_gee._abege.Align()
	}
	return nil
}

func (_gbd *GenericRefinementRegion) decodeTypicalPredictedLine(_adf, _afb, _cfe, _edf, _ada, _ddb int) error {
	_ebb := _adf - int(_gbd.ReferenceDY)
	_db := _gbd.ReferenceBitmap.GetByteIndex(0, _ebb)
	_gf := _gbd.RegionBitmap.GetByteIndex(0, _adf)
	var _fba error
	switch _gbd.TemplateID {
	case 0:
		_fba = _gbd.decodeTypicalPredictedLineTemplate0(_adf, _afb, _cfe, _edf, _ada, _ddb, _gf, _ebb, _db)
	case 1:
		_fba = _gbd.decodeTypicalPredictedLineTemplate1(_adf, _afb, _cfe, _edf, _ada, _ddb, _gf, _ebb, _db)
	}
	return _fba
}

func (_dcbfe *TextRegion) decodeRdw() (int64, error) {
	const _aadbf = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064w"
	if _dcbfe.IsHuffmanEncoded {
		if _dcbfe.SbHuffRDWidth == 3 {
			if _dcbfe._acccd == nil {
				var (
					_cabd int
					_fbea error
				)
				if _dcbfe.SbHuffFS == 3 {
					_cabd++
				}
				if _dcbfe.SbHuffDS == 3 {
					_cabd++
				}
				if _dcbfe.SbHuffDT == 3 {
					_cabd++
				}
				_dcbfe._acccd, _fbea = _dcbfe.getUserTable(_cabd)
				if _fbea != nil {
					return 0, _da.Wrap(_fbea, _aadbf, "")
				}
			}
			return _dcbfe._acccd.Decode(_dcbfe._gdbbd)
		}
		_ggad, _ggff := _cg.GetStandardTable(14 + int(_dcbfe.SbHuffRDWidth))
		if _ggff != nil {
			return 0, _da.Wrap(_ggff, _aadbf, "")
		}
		return _ggad.Decode(_dcbfe._gdbbd)
	}
	_ggfg, _fafa := _dcbfe._bggc.DecodeInt(_dcbfe._bgcbb)
	if _fafa != nil {
		return 0, _da.Wrap(_fafa, _aadbf, "")
	}
	return int64(_ggfg), nil
}

func (_edfg *RegionSegment) readCombinationOperator() error {
	_decb, _eaf := _edfg._dfbgc.ReadBits(3)
	if _eaf != nil {
		return _eaf
	}
	_edfg.CombinaionOperator = _ea.CombinationOperator(_decb & 0xF)
	return nil
}

func (_bed *GenericRegion) Init(h *Header, r *_e.Reader) error {
	_bed.RegionSegment = NewRegionSegment(r)
	_bed._fec = r
	return _bed.parseHeader()
}

func (_fdge *GenericRegion) GetRegionBitmap() (_agc *_ea.Bitmap, _ffe error) {
	if _fdge.Bitmap != nil {
		return _fdge.Bitmap, nil
	}
	if _fdge.IsMMREncoded {
		if _fdge._fdg == nil {
			_fdge._fdg, _ffe = _eae.New(_fdge._fec, int(_fdge.RegionSegment.BitmapWidth), int(_fdge.RegionSegment.BitmapHeight), _fdge.DataOffset, _fdge.DataLength)
			if _ffe != nil {
				return nil, _ffe
			}
		}
		_fdge.Bitmap, _ffe = _fdge._fdg.UncompressMMR()
		return _fdge.Bitmap, _ffe
	}
	if _ffe = _fdge.updateOverrideFlags(); _ffe != nil {
		return nil, _ffe
	}
	var _fdf int
	if _fdge._bda == nil {
		_fdge._bda, _ffe = _ag.New(_fdge._fec)
		if _ffe != nil {
			return nil, _ffe
		}
	}
	if _fdge._eced == nil {
		_fdge._eced = _ag.NewStats(65536, 1)
	}
	_fdge.Bitmap = _ea.New(int(_fdge.RegionSegment.BitmapWidth), int(_fdge.RegionSegment.BitmapHeight))
	_aeb := int(uint32(_fdge.Bitmap.Width+7) & (^uint32(7)))
	for _bac := 0; _bac < _fdge.Bitmap.Height; _bac++ {
		if _fdge.IsTPGDon {
			var _ecdd int
			_ecdd, _ffe = _fdge.decodeSLTP()
			if _ffe != nil {
				return nil, _ffe
			}
			_fdf ^= _ecdd
		}
		if _fdf == 1 {
			if _bac > 0 {
				if _ffe = _fdge.copyLineAbove(_bac); _ffe != nil {
					return nil, _ffe
				}
			}
		} else {
			if _ffe = _fdge.decodeLine(_bac, _fdge.Bitmap.Width, _aeb); _ffe != nil {
				return nil, _ffe
			}
		}
	}
	return _fdge.Bitmap, nil
}

func (_faea *SymbolDictionary) decodeDifferenceWidth() (int64, error) {
	if _faea.IsHuffmanEncoded {
		switch _faea.SdHuffDecodeWidthSelection {
		case 0:
			_aab, _fdga := _cg.GetStandardTable(2)
			if _fdga != nil {
				return 0, _fdga
			}
			return _aab.Decode(_faea._abege)
		case 1:
			_bcff, _afce := _cg.GetStandardTable(3)
			if _afce != nil {
				return 0, _afce
			}
			return _bcff.Decode(_faea._abege)
		case 3:
			if _faea._aceb == nil {
				var _gdcd int
				if _faea.SdHuffDecodeHeightSelection == 3 {
					_gdcd++
				}
				_dgaef, _ceaf := _faea.getUserTable(_gdcd)
				if _ceaf != nil {
					return 0, _ceaf
				}
				_faea._aceb = _dgaef
			}
			return _faea._aceb.Decode(_faea._abege)
		}
	} else {
		_beagc, _cged := _faea._agbbf.DecodeInt(_faea._eaecb)
		if _cged != nil {
			return 0, _cged
		}
		return int64(_beagc), nil
	}
	return 0, nil
}

func (_fcd *HalftoneRegion) grayScaleDecoding(_dcgf int) ([][]int, error) {
	var (
		_aae  []int8
		_gagc []int8
	)
	if !_fcd.IsMMREncoded {
		_aae = make([]int8, 4)
		_gagc = make([]int8, 4)
		if _fcd.HTemplate <= 1 {
			_aae[0] = 3
		} else if _fcd.HTemplate >= 2 {
			_aae[0] = 2
		}
		_gagc[0] = -1
		_aae[1] = -3
		_gagc[1] = -1
		_aae[2] = 2
		_gagc[2] = -2
		_aae[3] = -2
		_gagc[3] = -2
	}
	_eaccg := make([]*_ea.Bitmap, _dcgf)
	_fcbf := NewGenericRegion(_fcd._bfbe)
	_fcbf.setParametersMMR(_fcd.IsMMREncoded, _fcd.DataOffset, _fcd.DataLength, _fcd.HGridHeight, _fcd.HGridWidth, _fcd.HTemplate, false, _fcd.HSkipEnabled, _aae, _gagc)
	_cee := _dcgf - 1
	var _ggab error
	_eaccg[_cee], _ggab = _fcbf.GetRegionBitmap()
	if _ggab != nil {
		return nil, _ggab
	}
	for _cee > 0 {
		_cee--
		_fcbf.Bitmap = nil
		_eaccg[_cee], _ggab = _fcbf.GetRegionBitmap()
		if _ggab != nil {
			return nil, _ggab
		}
		if _ggab = _fcd.combineGrayscalePlanes(_eaccg, _cee); _ggab != nil {
			return nil, _ggab
		}
	}
	return _fcd.computeGrayScalePlanes(_eaccg, _dcgf)
}

func (_fabc *SymbolDictionary) String() string {
	_fggdd := &_cf.Builder{}
	_fggdd.WriteString("\n\u005b\u0053\u0059\u004dBO\u004c-\u0044\u0049\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u000a")
	_fggdd.WriteString(_f.Sprintf("\u0009-\u0020S\u0064\u0072\u0054\u0065\u006dp\u006c\u0061t\u0065\u0020\u0025\u0076\u000a", _fabc.SdrTemplate))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u0025\u0076\u000a", _fabc.SdTemplate))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0069\u0073\u0043\u006f\u0064\u0069\u006eg\u0043\u006f\u006e\u0074\u0065\u0078\u0074R\u0065\u0074\u0061\u0069\u006e\u0065\u0064\u0020\u0025\u0076\u000a", _fabc._cba))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0069\u0073\u0043\u006f\u0064\u0069\u006e\u0067C\u006f\u006e\u0074\u0065\u0078\u0074\u0055\u0073\u0065\u0064 \u0025\u0076\u000a", _fabc._bdgf))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0048u\u0066\u0066\u0041\u0067\u0067\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065S\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0025\u0076\u000a", _fabc.SdHuffAggInstanceSelection))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053d\u0048\u0075\u0066\u0066\u0042\u004d\u0053\u0069\u007a\u0065S\u0065l\u0065\u0063\u0074\u0069\u006f\u006e\u0020%\u0076\u000a", _fabc.SdHuffBMSizeSelection))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0048u\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0057\u0069\u0064\u0074\u0068S\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0025\u0076\u000a", _fabc.SdHuffDecodeWidthSelection))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020Sd\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064e\u0048e\u0069g\u0068t\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0020\u0025\u0076\u000a", _fabc.SdHuffDecodeHeightSelection))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020U\u0073\u0065\u0052\u0065f\u0069\u006e\u0065\u006d\u0065\u006e\u0074A\u0067\u0067\u0072\u0065\u0067\u0061\u0074\u0069\u006f\u006e\u0020\u0025\u0076\u000a", _fabc.UseRefinementAggregation))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020is\u0048\u0075\u0066\u0066\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0065\u0064\u0020\u0025\u0076\u000a", _fabc.IsHuffmanEncoded))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020S\u0064\u0041\u0054\u0058\u0020\u0025\u0076\u000a", _fabc.SdATX))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020S\u0064\u0041\u0054\u0059\u0020\u0025\u0076\u000a", _fabc.SdATY))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0072\u0041\u0054\u0058\u0020\u0025\u0076\u000a", _fabc.SdrATX))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0072\u0041\u0054\u0059\u0020\u0025\u0076\u000a", _fabc.SdrATY))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u004ff\u0045\u0078\u0070\u006f\u0072\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0025\u0076\u000a", _fabc.NumberOfExportedSymbols))
	_fggdd.WriteString(_f.Sprintf("\u0009-\u0020\u004e\u0075\u006db\u0065\u0072\u004f\u0066\u004ee\u0077S\u0079m\u0062\u006f\u006c\u0073\u0020\u0025\u0076\n", _fabc.NumberOfNewSymbols))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u004ff\u0049\u006d\u0070\u006f\u0072\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0025\u0076\u000a", _fabc._afag))
	_fggdd.WriteString(_f.Sprintf("\u0009\u002d \u006e\u0075\u006d\u0062\u0065\u0072\u004f\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0064\u0053\u0079\u006d\u0062\u006f\u006c\u0073 %\u0076\u000a", _fabc._addfb))
	return _fggdd.String()
}

func (_gcgag *SymbolDictionary) setAtPixels() error {
	if _gcgag.IsHuffmanEncoded {
		return nil
	}
	_edab := 1
	if _gcgag.SdTemplate == 0 {
		_edab = 4
	}
	if _aage := _gcgag.readAtPixels(_edab); _aage != nil {
		return _aage
	}
	return nil
}

func (_cadc *TextRegion) initSymbols() error {
	const _fegg = "i\u006e\u0069\u0074\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	for _, _ceac := range _cadc.Header.RTSegments {
		if _ceac == nil {
			return _da.Error(_fegg, "\u006e\u0069\u006c\u0020\u0073\u0065\u0067\u006de\u006e\u0074\u0020pr\u006f\u0076\u0069\u0064\u0065\u0064 \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0074\u0065\u0078\u0074\u0020\u0072\u0065g\u0069\u006f\u006e\u0020\u0053\u0079\u006d\u0062o\u006c\u0073")
		}
		if _ceac.Type == 0 {
			_agefb, _bccg := _ceac.GetSegmentData()
			if _bccg != nil {
				return _da.Wrap(_bccg, _fegg, "")
			}
			_fcggg, _bffcg := _agefb.(*SymbolDictionary)
			if !_bffcg {
				return _da.Error(_fegg, "\u0072e\u0066\u0065r\u0072\u0065\u0064 \u0054\u006f\u0020\u0053\u0065\u0067\u006de\u006e\u0074\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0053\u0079\u006d\u0062\u006f\u006cD\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
			}
			_fcggg._gaac = _cadc._cgfd
			_bdaf, _bccg := _fcggg.GetDictionary()
			if _bccg != nil {
				return _da.Wrap(_bccg, _fegg, "")
			}
			_cadc.Symbols = append(_cadc.Symbols, _bdaf...)
		}
	}
	_cadc.NumberOfSymbols = uint32(len(_cadc.Symbols))
	return nil
}

type GenericRefinementRegion struct {
	_agbg           templater
	_ad             templater
	_gbc            *_e.Reader
	_bg             *Header
	RegionInfo      *RegionSegment
	IsTPGROn        bool
	TemplateID      int8
	Template        templater
	GrAtX           []int8
	GrAtY           []int8
	RegionBitmap    *_ea.Bitmap
	ReferenceBitmap *_ea.Bitmap
	ReferenceDX     int32
	ReferenceDY     int32
	_cc             *_ag.Decoder
	_gbf            *_ag.DecoderStats
	_bba            bool
	_ge             []bool
}

func (_abgc *HalftoneRegion) computeSegmentDataStructure() error {
	_abgc.DataOffset = _abgc._bfbe.AbsolutePosition()
	_abgc.DataHeaderLength = _abgc.DataOffset - _abgc.DataHeaderOffset
	_abgc.DataLength = int64(_abgc._bfbe.AbsoluteLength()) - _abgc.DataHeaderLength
	return nil
}

func (_fbf *SymbolDictionary) readRefinementAtPixels(_bdgde int) error {
	_fbf.SdrATX = make([]int8, _bdgde)
	_fbf.SdrATY = make([]int8, _bdgde)
	var (
		_daa  byte
		_gcga error
	)
	for _dgc := 0; _dgc < _bdgde; _dgc++ {
		_daa, _gcga = _fbf._abege.ReadByte()
		if _gcga != nil {
			return _gcga
		}
		_fbf.SdrATX[_dgc] = int8(_daa)
		_daa, _gcga = _fbf._abege.ReadByte()
		if _gcga != nil {
			return _gcga
		}
		_fbf.SdrATY[_dgc] = int8(_daa)
	}
	return nil
}
func (_bfe *PageInformationSegment) Size() int { return 19 }
func (_dbgb *Header) writeSegmentDataLength(_dcdg _e.BinaryWriter) (_gdgb int, _baca error) {
	_agag := make([]byte, 4)
	_ab.BigEndian.PutUint32(_agag, uint32(_dbgb.SegmentDataLength))
	if _gdgb, _baca = _dcdg.Write(_agag); _baca != nil {
		return 0, _da.Wrap(_baca, "\u0048\u0065a\u0064\u0065\u0072\u002e\u0077\u0072\u0069\u0074\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0044\u0061\u0074\u0061\u004c\u0065ng\u0074\u0068", "")
	}
	return _gdgb, nil
}

func (_cecef *TextRegion) readAmountOfSymbolInstances() error {
	_eabf, _eeaf := _cecef._gdbbd.ReadBits(32)
	if _eeaf != nil {
		return _eeaf
	}
	_cecef.NumberOfSymbolInstances = uint32(_eabf & _b.MaxUint32)
	_edbg := _cecef.RegionInfo.BitmapWidth * _cecef.RegionInfo.BitmapHeight
	if _edbg < _cecef.NumberOfSymbolInstances {
		_eg.Log.Debug("\u004c\u0069\u006d\u0069t\u0069\u006e\u0067\u0020t\u0068\u0065\u0020n\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020d\u0065\u0063\u006f\u0064e\u0064\u0020\u0073\u0079m\u0062\u006f\u006c\u0020\u0069n\u0073\u0074\u0061\u006e\u0063\u0065\u0073 \u0074\u006f\u0020\u006f\u006ee\u0020\u0070\u0065\u0072\u0020\u0070\u0069\u0078\u0065l\u0020\u0028\u0020\u0025\u0064\u0020\u0069\u006e\u0073\u0074\u0065\u0061\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u0029", _edbg, _cecef.NumberOfSymbolInstances)
		_cecef.NumberOfSymbolInstances = _edbg
	}
	return nil
}

func (_cfgc *HalftoneRegion) GetRegionBitmap() (*_ea.Bitmap, error) {
	if _cfgc.HalftoneRegionBitmap != nil {
		return _cfgc.HalftoneRegionBitmap, nil
	}
	var _ecg error
	_cfgc.HalftoneRegionBitmap = _ea.New(int(_cfgc.RegionSegment.BitmapWidth), int(_cfgc.RegionSegment.BitmapHeight))
	if _cfgc.Patterns == nil || len(_cfgc.Patterns) == 0 {
		_cfgc.Patterns, _ecg = _cfgc.GetPatterns()
		if _ecg != nil {
			return nil, _ecg
		}
	}
	if _cfgc.HDefaultPixel == 1 {
		_cfgc.HalftoneRegionBitmap.SetDefaultPixel()
	}
	_ccae := _b.Ceil(_b.Log(float64(len(_cfgc.Patterns))) / _b.Log(2))
	_fbc := int(_ccae)
	var _aeef [][]int
	_aeef, _ecg = _cfgc.grayScaleDecoding(_fbc)
	if _ecg != nil {
		return nil, _ecg
	}
	if _ecg = _cfgc.renderPattern(_aeef); _ecg != nil {
		return nil, _ecg
	}
	return _cfgc.HalftoneRegionBitmap, nil
}

func (_ffbd *Header) readNumberOfReferredToSegments(_ffec *_e.Reader) (uint64, error) {
	const _baeb = "\u0072\u0065\u0061\u0064\u004e\u0075\u006d\u0062\u0065\u0072O\u0066\u0052\u0065\u0066\u0065\u0072\u0072e\u0064\u0054\u006f\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073"
	_faeg, _ebab := _ffec.ReadBits(3)
	if _ebab != nil {
		return 0, _da.Wrap(_ebab, _baeb, "\u0063\u006f\u0075n\u0074\u0020\u006f\u0066\u0020\u0072\u0074\u0073")
	}
	_faeg &= 0xf
	var _ecgd []byte
	if _faeg <= 4 {
		_ecgd = make([]byte, 5)
		for _bge := 0; _bge <= 4; _bge++ {
			_bea, _dccc := _ffec.ReadBit()
			if _dccc != nil {
				return 0, _da.Wrap(_dccc, _baeb, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
			}
			_ecgd[_bge] = byte(_bea)
		}
	} else {
		_faeg, _ebab = _ffec.ReadBits(29)
		if _ebab != nil {
			return 0, _ebab
		}
		_faeg &= _b.MaxInt32
		_ffad := (_faeg + 8) >> 3
		_ffad <<= 3
		_ecgd = make([]byte, _ffad)
		var _efc uint64
		for _efc = 0; _efc < _ffad; _efc++ {
			_cffd, _abdf := _ffec.ReadBit()
			if _abdf != nil {
				return 0, _da.Wrap(_abdf, _baeb, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
			}
			_ecgd[_efc] = byte(_cffd)
		}
	}
	return _faeg, nil
}

func (_deee *PatternDictionary) readIsMMREncoded() error {
	_ddbe, _dgbe := _deee._ccca.ReadBit()
	if _dgbe != nil {
		return _dgbe
	}
	if _ddbe != 0 {
		_deee.IsMMREncoded = true
	}
	return nil
}

func (_bcaa *template1) form(_bcd, _bae, _dcgg, _dgf, _ccb int16) int16 {
	return ((_bcd & 0x02) << 8) | (_bae << 6) | ((_dcgg & 0x03) << 4) | (_dgf << 1) | _ccb
}
func (_dagbc *TableSegment) HtPS() int32 { return _dagbc._cggf }
func (_cge *GenericRefinementRegion) Init(header *Header, r *_e.Reader) error {
	_cge._bg = header
	_cge._gbc = r
	_cge.RegionInfo = NewRegionSegment(r)
	return _cge.parseHeader()
}

func (_dage *GenericRegion) Encode(w _e.BinaryWriter) (_abcc int, _fcg error) {
	const _acg = "G\u0065n\u0065\u0072\u0069\u0063\u0052\u0065\u0067\u0069o\u006e\u002e\u0045\u006eco\u0064\u0065"
	if _dage.Bitmap == nil {
		return 0, _da.Error(_acg, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	_fgff, _fcg := _dage.RegionSegment.Encode(w)
	if _fcg != nil {
		return 0, _da.Wrap(_fcg, _acg, "\u0052\u0065\u0067\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_abcc += _fgff
	if _fcg = w.SkipBits(4); _fcg != nil {
		return _abcc, _da.Wrap(_fcg, _acg, "\u0073k\u0069p\u0020\u0072\u0065\u0073\u0065r\u0076\u0065d\u0020\u0062\u0069\u0074\u0073")
	}
	var _geg int
	if _dage.IsTPGDon {
		_geg = 1
	}
	if _fcg = w.WriteBit(_geg); _fcg != nil {
		return _abcc, _da.Wrap(_fcg, _acg, "\u0074\u0070\u0067\u0064\u006f\u006e")
	}
	_geg = 0
	if _fcg = w.WriteBit(int(_dage.GBTemplate>>1) & 0x01); _fcg != nil {
		return _abcc, _da.Wrap(_fcg, _acg, "f\u0069r\u0073\u0074\u0020\u0067\u0062\u0074\u0065\u006dp\u006c\u0061\u0074\u0065 b\u0069\u0074")
	}
	if _fcg = w.WriteBit(int(_dage.GBTemplate) & 0x01); _fcg != nil {
		return _abcc, _da.Wrap(_fcg, _acg, "s\u0065\u0063\u006f\u006ed \u0067b\u0074\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u0062\u0069\u0074")
	}
	if _dage.UseMMR {
		_geg = 1
	}
	if _fcg = w.WriteBit(_geg); _fcg != nil {
		return _abcc, _da.Wrap(_fcg, _acg, "u\u0073\u0065\u0020\u004d\u004d\u0052\u0020\u0062\u0069\u0074")
	}
	_abcc++
	if _fgff, _fcg = _dage.writeGBAtPixels(w); _fcg != nil {
		return _abcc, _da.Wrap(_fcg, _acg, "")
	}
	_abcc += _fgff
	_cadf := _bb.New()
	if _fcg = _cadf.EncodeBitmap(_dage.Bitmap, _dage.IsTPGDon); _fcg != nil {
		return _abcc, _da.Wrap(_fcg, _acg, "")
	}
	_cadf.Final()
	var _bced int64
	if _bced, _fcg = _cadf.WriteTo(w); _fcg != nil {
		return _abcc, _da.Wrap(_fcg, _acg, "")
	}
	_abcc += int(_bced)
	return _abcc, nil
}
