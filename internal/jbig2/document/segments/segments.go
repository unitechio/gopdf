package segments

import (
	_fe "encoding/binary"
	_ge "errors"
	_e "fmt"
	_c "image"
	_gc "io"
	_f "math"
	_cc "strings"
	_a "time"

	_da "bitbucket.org/shenghui0779/gopdf/common"
	_fg "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_ca "bitbucket.org/shenghui0779/gopdf/internal/jbig2/basic"
	_d "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_b "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder/arithmetic"
	_ba "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder/huffman"
	_ad "bitbucket.org/shenghui0779/gopdf/internal/jbig2/decoder/mmr"
	_ga "bitbucket.org/shenghui0779/gopdf/internal/jbig2/encoder/arithmetic"
	_fd "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
	_gd "bitbucket.org/shenghui0779/gopdf/internal/jbig2/internal"
	_cd "golang.org/x/xerrors"
)

type TableSegment struct {
	_defgf _fg.StreamReader
	_eaed  int32
	_fbdc  int32
	_cgeg  int32
	_cccb  int32
	_efgd  int32
}

func (_cad *GenericRegion) decodeTemplate0b(_cadb, _cdeg, _dade int, _gbbf, _fcd int) (_cge error) {
	const _gefg = "\u0064\u0065c\u006f\u0064\u0065T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0030\u0062"
	var (
		_gfg, _egga  int
		_aac, _eebba int
		_dadec       byte
		_agd         int
	)
	if _cadb >= 1 {
		_dadec, _cge = _cad.Bitmap.GetByte(_fcd)
		if _cge != nil {
			return _fd.Wrap(_cge, _gefg, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_aac = int(_dadec)
	}
	if _cadb >= 2 {
		_dadec, _cge = _cad.Bitmap.GetByte(_fcd - _cad.Bitmap.RowStride)
		if _cge != nil {
			return _fd.Wrap(_cge, _gefg, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_eebba = int(_dadec) << 6
	}
	_gfg = (_aac & 0xf0) | (_eebba & 0x3800)
	for _afg := 0; _afg < _dade; _afg = _agd {
		var (
			_agg byte
			_gfe int
		)
		_agd = _afg + 8
		if _fdgb := _cdeg - _afg; _fdgb > 8 {
			_gfe = 8
		} else {
			_gfe = _fdgb
		}
		if _cadb > 0 {
			_aac <<= 8
			if _agd < _cdeg {
				_dadec, _cge = _cad.Bitmap.GetByte(_fcd + 1)
				if _cge != nil {
					return _fd.Wrap(_cge, _gefg, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_aac |= int(_dadec)
			}
		}
		if _cadb > 1 {
			_eebba <<= 8
			if _agd < _cdeg {
				_dadec, _cge = _cad.Bitmap.GetByte(_fcd - _cad.Bitmap.RowStride + 1)
				if _cge != nil {
					return _fd.Wrap(_cge, _gefg, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_eebba |= int(_dadec) << 6
			}
		}
		for _fabc := 0; _fabc < _gfe; _fabc++ {
			_fgca := uint(7 - _fabc)
			if _cad._fgf {
				_egga = _cad.overrideAtTemplate0b(_gfg, _afg+_fabc, _cadb, int(_agg), _fabc, int(_fgca))
				_cad._aae.SetIndex(int32(_egga))
			} else {
				_cad._aae.SetIndex(int32(_gfg))
			}
			var _bbbge int
			_bbbge, _cge = _cad._acb.DecodeBit(_cad._aae)
			if _cge != nil {
				return _fd.Wrap(_cge, _gefg, "")
			}
			_agg |= byte(_bbbge << _fgca)
			_gfg = ((_gfg & 0x7bf7) << 1) | _bbbge | ((_aac >> _fgca) & 0x10) | ((_eebba >> _fgca) & 0x800)
		}
		if _gbfg := _cad.Bitmap.SetByte(_gbbf, _agg); _gbfg != nil {
			return _fd.Wrap(_gbfg, _gefg, "")
		}
		_gbbf++
		_fcd++
	}
	return nil
}

var _ templater = &template0{}

func (_eabb *TextRegion) createRegionBitmap() error {
	_eabb.RegionBitmap = _d.New(int(_eabb.RegionInfo.BitmapWidth), int(_eabb.RegionInfo.BitmapHeight))
	if _eabb.DefaultPixel != 0 {
		_eabb.RegionBitmap.SetDefaultPixel()
	}
	return nil
}
func (_fede *GenericRegion) decodeTemplate0a(_dbbc, _aab, _egaa int, _gafff, _acfb int) (_gggg error) {
	const _badf = "\u0064\u0065c\u006f\u0064\u0065T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0030\u0061"
	var (
		_ebfg, _egd int
		_deea, _gec int
		_gafc       byte
		_fdgc       int
	)
	if _dbbc >= 1 {
		_gafc, _gggg = _fede.Bitmap.GetByte(_acfb)
		if _gggg != nil {
			return _fd.Wrap(_gggg, _badf, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_deea = int(_gafc)
	}
	if _dbbc >= 2 {
		_gafc, _gggg = _fede.Bitmap.GetByte(_acfb - _fede.Bitmap.RowStride)
		if _gggg != nil {
			return _fd.Wrap(_gggg, _badf, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_gec = int(_gafc) << 6
	}
	_ebfg = (_deea & 0xf0) | (_gec & 0x3800)
	for _dgc := 0; _dgc < _egaa; _dgc = _fdgc {
		var (
			_ecfeg byte
			_bdb   int
		)
		_fdgc = _dgc + 8
		if _fba := _aab - _dgc; _fba > 8 {
			_bdb = 8
		} else {
			_bdb = _fba
		}
		if _dbbc > 0 {
			_deea <<= 8
			if _fdgc < _aab {
				_gafc, _gggg = _fede.Bitmap.GetByte(_acfb + 1)
				if _gggg != nil {
					return _fd.Wrap(_gggg, _badf, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_deea |= int(_gafc)
			}
		}
		if _dbbc > 1 {
			_dab := _acfb - _fede.Bitmap.RowStride + 1
			_gec <<= 8
			if _fdgc < _aab {
				_gafc, _gggg = _fede.Bitmap.GetByte(_dab)
				if _gggg != nil {
					return _fd.Wrap(_gggg, _badf, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_gec |= int(_gafc) << 6
			} else {
				_gec |= 0
			}
		}
		for _eagfg := 0; _eagfg < _bdb; _eagfg++ {
			_dcba := uint(7 - _eagfg)
			if _fede._fgf {
				_egd = _fede.overrideAtTemplate0a(_ebfg, _dgc+_eagfg, _dbbc, int(_ecfeg), _eagfg, int(_dcba))
				_fede._aae.SetIndex(int32(_egd))
			} else {
				_fede._aae.SetIndex(int32(_ebfg))
			}
			var _bbbg int
			_bbbg, _gggg = _fede._acb.DecodeBit(_fede._aae)
			if _gggg != nil {
				return _fd.Wrap(_gggg, _badf, "")
			}
			_ecfeg |= byte(_bbbg) << _dcba
			_ebfg = ((_ebfg & 0x7bf7) << 1) | _bbbg | ((_deea >> _dcba) & 0x10) | ((_gec >> _dcba) & 0x800)
		}
		if _fea := _fede.Bitmap.SetByte(_gafff, _ecfeg); _fea != nil {
			return _fd.Wrap(_fea, _badf, "")
		}
		_gafff++
		_acfb++
	}
	return nil
}
func (_agcdc *TableSegment) HtPS() int32 { return _agcdc._fbdc }
func (_acge *PatternDictionary) readIsMMREncoded() error {
	_cebf, _eefd := _acge._fbdd.ReadBit()
	if _eefd != nil {
		return _eefd
	}
	if _cebf != 0 {
		_acge.IsMMREncoded = true
	}
	return nil
}
func (_bbdf *RegionSegment) parseHeader() error {
	const _dbaf = "p\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_da.Log.Trace("\u005b\u0052\u0045\u0047I\u004f\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002dH\u0045A\u0044\u0045\u0052\u005d\u0020\u0042\u0065g\u0069\u006e")
	defer func() {
		_da.Log.Trace("\u005b\u0052\u0045G\u0049\u004f\u004e\u005d[\u0050\u0041\u0052\u0053\u0045\u002d\u0048E\u0041\u0044\u0045\u0052\u005d\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
	}()
	_fccb, _ggdg := _bbdf._dbge.ReadBits(32)
	if _ggdg != nil {
		return _fd.Wrap(_ggdg, _dbaf, "\u0077\u0069\u0064t\u0068")
	}
	_bbdf.BitmapWidth = uint32(_fccb & _f.MaxUint32)
	_fccb, _ggdg = _bbdf._dbge.ReadBits(32)
	if _ggdg != nil {
		return _fd.Wrap(_ggdg, _dbaf, "\u0068\u0065\u0069\u0067\u0068\u0074")
	}
	_bbdf.BitmapHeight = uint32(_fccb & _f.MaxUint32)
	_fccb, _ggdg = _bbdf._dbge.ReadBits(32)
	if _ggdg != nil {
		return _fd.Wrap(_ggdg, _dbaf, "\u0078\u0020\u006c\u006f\u0063\u0061\u0074\u0069\u006f\u006e")
	}
	_bbdf.XLocation = uint32(_fccb & _f.MaxUint32)
	_fccb, _ggdg = _bbdf._dbge.ReadBits(32)
	if _ggdg != nil {
		return _fd.Wrap(_ggdg, _dbaf, "\u0079\u0020\u006c\u006f\u0063\u0061\u0074\u0069\u006f\u006e")
	}
	_bbdf.YLocation = uint32(_fccb & _f.MaxUint32)
	if _, _ggdg = _bbdf._dbge.ReadBits(5); _ggdg != nil {
		return _fd.Wrap(_ggdg, _dbaf, "\u0064i\u0072\u0079\u0020\u0072\u0065\u0061d")
	}
	if _ggdg = _bbdf.readCombinationOperator(); _ggdg != nil {
		return _fd.Wrap(_ggdg, _dbaf, "c\u006fm\u0062\u0069\u006e\u0061\u0074\u0069\u006f\u006e \u006f\u0070\u0065\u0072at\u006f\u0072")
	}
	return nil
}
func (_bcbfe *TableSegment) HtLow() int32 { return _bcbfe._cccb }
func (_daec *GenericRegion) decodeTemplate1(_fdca, _dadc, _bfcc int, _faa, _adc int) (_afga error) {
	const _gbe = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0031"
	var (
		_aedd, _defg  int
		_bgfab, _adca int
		_bagg         byte
		_bfe, _fbd    int
	)
	if _fdca >= 1 {
		_bagg, _afga = _daec.Bitmap.GetByte(_adc)
		if _afga != nil {
			return _fd.Wrap(_afga, _gbe, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_bgfab = int(_bagg)
	}
	if _fdca >= 2 {
		_bagg, _afga = _daec.Bitmap.GetByte(_adc - _daec.Bitmap.RowStride)
		if _afga != nil {
			return _fd.Wrap(_afga, _gbe, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00202")
		}
		_adca = int(_bagg) << 5
	}
	_aedd = ((_bgfab >> 1) & 0x1f8) | ((_adca >> 1) & 0x1e00)
	for _cgb := 0; _cgb < _bfcc; _cgb = _bfe {
		var (
			_ddda byte
			_bbc  int
		)
		_bfe = _cgb + 8
		if _cea := _dadc - _cgb; _cea > 8 {
			_bbc = 8
		} else {
			_bbc = _cea
		}
		if _fdca > 0 {
			_bgfab <<= 8
			if _bfe < _dadc {
				_bagg, _afga = _daec.Bitmap.GetByte(_adc + 1)
				if _afga != nil {
					return _fd.Wrap(_afga, _gbe, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0030")
				}
				_bgfab |= int(_bagg)
			}
		}
		if _fdca > 1 {
			_adca <<= 8
			if _bfe < _dadc {
				_bagg, _afga = _daec.Bitmap.GetByte(_adc - _daec.Bitmap.RowStride + 1)
				if _afga != nil {
					return _fd.Wrap(_afga, _gbe, "\u006c\u0069\u006e\u0065\u0020\u003e\u0020\u0031")
				}
				_adca |= int(_bagg) << 5
			}
		}
		for _cgdf := 0; _cgdf < _bbc; _cgdf++ {
			if _daec._fgf {
				_defg = _daec.overrideAtTemplate1(_aedd, _cgb+_cgdf, _fdca, int(_ddda), _cgdf)
				_daec._aae.SetIndex(int32(_defg))
			} else {
				_daec._aae.SetIndex(int32(_aedd))
			}
			_fbd, _afga = _daec._acb.DecodeBit(_daec._aae)
			if _afga != nil {
				return _fd.Wrap(_afga, _gbe, "")
			}
			_ddda |= byte(_fbd) << uint(7-_cgdf)
			_eebc := uint(8 - _cgdf)
			_aedd = ((_aedd & 0xefb) << 1) | _fbd | ((_bgfab >> _eebc) & 0x8) | ((_adca >> _eebc) & 0x200)
		}
		if _ede := _daec.Bitmap.SetByte(_faa, _ddda); _ede != nil {
			return _fd.Wrap(_ede, _gbe, "")
		}
		_faa++
		_adc++
	}
	return nil
}
func (_ddeba *PageInformationSegment) readIsLossless() error {
	_dff, _dabca := _ddeba._ace.ReadBit()
	if _dabca != nil {
		return _dabca
	}
	if _dff == 1 {
		_ddeba.IsLossless = true
	}
	return nil
}

type HalftoneRegion struct {
	_gdfe                _fg.StreamReader
	_fdag                *Header
	DataHeaderOffset     int64
	DataHeaderLength     int64
	DataOffset           int64
	DataLength           int64
	RegionSegment        *RegionSegment
	HDefaultPixel        int8
	CombinationOperator  _d.CombinationOperator
	HSkipEnabled         bool
	HTemplate            byte
	IsMMREncoded         bool
	HGridWidth           uint32
	HGridHeight          uint32
	HGridX               int32
	HGridY               int32
	HRegionX             uint16
	HRegionY             uint16
	HalftoneRegionBitmap *_d.Bitmap
	Patterns             []*_d.Bitmap
}

func (_ffcc *SymbolDictionary) decodeHeightClassDeltaHeight() (int64, error) {
	if _ffcc.IsHuffmanEncoded {
		return _ffcc.decodeHeightClassDeltaHeightWithHuffman()
	}
	_ggdc, _ebba := _ffcc._fffd.DecodeInt(_ffcc._aabg)
	if _ebba != nil {
		return 0, _ebba
	}
	return int64(_ggdc), nil
}
func (_bbe *GenericRegion) overrideAtTemplate0b(_bgfcf, _fcb, _fgg, _bcag, _gcea, _cada int) int {
	if _bbe.GBAtOverride[0] {
		_bgfcf &= 0xFFFD
		if _bbe.GBAtY[0] == 0 && _bbe.GBAtX[0] >= -int8(_gcea) {
			_bgfcf |= (_bcag >> uint(int8(_cada)-_bbe.GBAtX[0]&0x1)) << 1
		} else {
			_bgfcf |= int(_bbe.getPixel(_fcb+int(_bbe.GBAtX[0]), _fgg+int(_bbe.GBAtY[0]))) << 1
		}
	}
	if _bbe.GBAtOverride[1] {
		_bgfcf &= 0xDFFF
		if _bbe.GBAtY[1] == 0 && _bbe.GBAtX[1] >= -int8(_gcea) {
			_bgfcf |= (_bcag >> uint(int8(_cada)-_bbe.GBAtX[1]&0x1)) << 13
		} else {
			_bgfcf |= int(_bbe.getPixel(_fcb+int(_bbe.GBAtX[1]), _fgg+int(_bbe.GBAtY[1]))) << 13
		}
	}
	if _bbe.GBAtOverride[2] {
		_bgfcf &= 0xFDFF
		if _bbe.GBAtY[2] == 0 && _bbe.GBAtX[2] >= -int8(_gcea) {
			_bgfcf |= (_bcag >> uint(int8(_cada)-_bbe.GBAtX[2]&0x1)) << 9
		} else {
			_bgfcf |= int(_bbe.getPixel(_fcb+int(_bbe.GBAtX[2]), _fgg+int(_bbe.GBAtY[2]))) << 9
		}
	}
	if _bbe.GBAtOverride[3] {
		_bgfcf &= 0xBFFF
		if _bbe.GBAtY[3] == 0 && _bbe.GBAtX[3] >= -int8(_gcea) {
			_bgfcf |= (_bcag >> uint(int8(_cada)-_bbe.GBAtX[3]&0x1)) << 14
		} else {
			_bgfcf |= int(_bbe.getPixel(_fcb+int(_bbe.GBAtX[3]), _fgg+int(_bbe.GBAtY[3]))) << 14
		}
	}
	if _bbe.GBAtOverride[4] {
		_bgfcf &= 0xEFFF
		if _bbe.GBAtY[4] == 0 && _bbe.GBAtX[4] >= -int8(_gcea) {
			_bgfcf |= (_bcag >> uint(int8(_cada)-_bbe.GBAtX[4]&0x1)) << 12
		} else {
			_bgfcf |= int(_bbe.getPixel(_fcb+int(_bbe.GBAtX[4]), _fgg+int(_bbe.GBAtY[4]))) << 12
		}
	}
	if _bbe.GBAtOverride[5] {
		_bgfcf &= 0xFFDF
		if _bbe.GBAtY[5] == 0 && _bbe.GBAtX[5] >= -int8(_gcea) {
			_bgfcf |= (_bcag >> uint(int8(_cada)-_bbe.GBAtX[5]&0x1)) << 5
		} else {
			_bgfcf |= int(_bbe.getPixel(_fcb+int(_bbe.GBAtX[5]), _fgg+int(_bbe.GBAtY[5]))) << 5
		}
	}
	if _bbe.GBAtOverride[6] {
		_bgfcf &= 0xFFFB
		if _bbe.GBAtY[6] == 0 && _bbe.GBAtX[6] >= -int8(_gcea) {
			_bgfcf |= (_bcag >> uint(int8(_cada)-_bbe.GBAtX[6]&0x1)) << 2
		} else {
			_bgfcf |= int(_bbe.getPixel(_fcb+int(_bbe.GBAtX[6]), _fgg+int(_bbe.GBAtY[6]))) << 2
		}
	}
	if _bbe.GBAtOverride[7] {
		_bgfcf &= 0xFFF7
		if _bbe.GBAtY[7] == 0 && _bbe.GBAtX[7] >= -int8(_gcea) {
			_bgfcf |= (_bcag >> uint(int8(_cada)-_bbe.GBAtX[7]&0x1)) << 3
		} else {
			_bgfcf |= int(_bbe.getPixel(_fcb+int(_bbe.GBAtX[7]), _fgg+int(_bbe.GBAtY[7]))) << 3
		}
	}
	if _bbe.GBAtOverride[8] {
		_bgfcf &= 0xF7FF
		if _bbe.GBAtY[8] == 0 && _bbe.GBAtX[8] >= -int8(_gcea) {
			_bgfcf |= (_bcag >> uint(int8(_cada)-_bbe.GBAtX[8]&0x1)) << 11
		} else {
			_bgfcf |= int(_bbe.getPixel(_fcb+int(_bbe.GBAtX[8]), _fgg+int(_bbe.GBAtY[8]))) << 11
		}
	}
	if _bbe.GBAtOverride[9] {
		_bgfcf &= 0xFFEF
		if _bbe.GBAtY[9] == 0 && _bbe.GBAtX[9] >= -int8(_gcea) {
			_bgfcf |= (_bcag >> uint(int8(_cada)-_bbe.GBAtX[9]&0x1)) << 4
		} else {
			_bgfcf |= int(_bbe.getPixel(_fcb+int(_bbe.GBAtX[9]), _fgg+int(_bbe.GBAtY[9]))) << 4
		}
	}
	if _bbe.GBAtOverride[10] {
		_bgfcf &= 0x7FFF
		if _bbe.GBAtY[10] == 0 && _bbe.GBAtX[10] >= -int8(_gcea) {
			_bgfcf |= (_bcag >> uint(int8(_cada)-_bbe.GBAtX[10]&0x1)) << 15
		} else {
			_bgfcf |= int(_bbe.getPixel(_fcb+int(_bbe.GBAtX[10]), _fgg+int(_bbe.GBAtY[10]))) << 15
		}
	}
	if _bbe.GBAtOverride[11] {
		_bgfcf &= 0xFDFF
		if _bbe.GBAtY[11] == 0 && _bbe.GBAtX[11] >= -int8(_gcea) {
			_bgfcf |= (_bcag >> uint(int8(_cada)-_bbe.GBAtX[11]&0x1)) << 10
		} else {
			_bgfcf |= int(_bbe.getPixel(_fcb+int(_bbe.GBAtX[11]), _fgg+int(_bbe.GBAtY[11]))) << 10
		}
	}
	return _bgfcf
}
func (_fcega *SymbolDictionary) encodeSymbols(_bceb _fg.BinaryWriter) (_bbde int, _gcgd error) {
	const _efcc = "\u0065\u006e\u0063o\u0064\u0065\u0053\u0079\u006d\u0062\u006f\u006c"
	_bcbf := _ga.New()
	_bcbf.Init()
	_cfea, _gcgd := _fcega._dbag.SelectByIndexes(_fcega._fagdd)
	if _gcgd != nil {
		return 0, _fd.Wrap(_gcgd, _efcc, "\u0069n\u0069\u0074\u0069\u0061\u006c")
	}
	_bgdf := map[*_d.Bitmap]int{}
	for _bfda, _dgcg := range _cfea.Values {
		_bgdf[_dgcg] = _bfda
	}
	_cfea.SortByHeight()
	var _ecbg, _gebd int
	_dgad, _gcgd := _cfea.GroupByHeight()
	if _gcgd != nil {
		return 0, _fd.Wrap(_gcgd, _efcc, "")
	}
	for _, _bfdc := range _dgad.Values {
		_cdce := _bfdc.Values[0].Height
		_ecfd := _cdce - _ecbg
		if _gcgd = _bcbf.EncodeInteger(_ga.IADH, _ecfd); _gcgd != nil {
			return 0, _fd.Wrapf(_gcgd, _efcc, "\u0049\u0041\u0044\u0048\u0020\u0066\u006f\u0072\u0020\u0064\u0068\u003a \u0027\u0025\u0064\u0027", _ecfd)
		}
		_ecbg = _cdce
		_afad, _edgg := _bfdc.GroupByWidth()
		if _edgg != nil {
			return 0, _fd.Wrapf(_edgg, _efcc, "\u0068\u0065\u0069g\u0068\u0074\u003a\u0020\u0027\u0025\u0064\u0027", _cdce)
		}
		var _aabba int
		for _, _gfec := range _afad.Values {
			for _, _cefb := range _gfec.Values {
				_cdae := _cefb.Width
				_cffd := _cdae - _aabba
				if _edgg = _bcbf.EncodeInteger(_ga.IADW, _cffd); _edgg != nil {
					return 0, _fd.Wrapf(_edgg, _efcc, "\u0049\u0041\u0044\u0057\u0020\u0066\u006f\u0072\u0020\u0064\u0077\u003a \u0027\u0025\u0064\u0027", _cffd)
				}
				_aabba += _cffd
				if _edgg = _bcbf.EncodeBitmap(_cefb, false); _edgg != nil {
					return 0, _fd.Wrapf(_edgg, _efcc, "H\u0065i\u0067\u0068\u0074\u003a\u0020\u0025\u0064\u0020W\u0069\u0064\u0074\u0068: \u0025\u0064", _cdce, _cdae)
				}
				_abfb := _bgdf[_cefb]
				_fcega._dcbf[_abfb] = _gebd
				_gebd++
			}
		}
		if _edgg = _bcbf.EncodeOOB(_ga.IADW); _edgg != nil {
			return 0, _fd.Wrap(_edgg, _efcc, "\u0049\u0041\u0044\u0057")
		}
	}
	if _gcgd = _bcbf.EncodeInteger(_ga.IAEX, 0); _gcgd != nil {
		return 0, _fd.Wrap(_gcgd, _efcc, "\u0065\u0078p\u006f\u0072\u0074e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	if _gcgd = _bcbf.EncodeInteger(_ga.IAEX, len(_fcega._fagdd)); _gcgd != nil {
		return 0, _fd.Wrap(_gcgd, _efcc, "\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0073\u0079m\u0062\u006f\u006c\u0073")
	}
	_bcbf.Final()
	_ccfa, _gcgd := _bcbf.WriteTo(_bceb)
	if _gcgd != nil {
		return 0, _fd.Wrap(_gcgd, _efcc, "\u0077\u0072i\u0074\u0069\u006e\u0067 \u0065\u006ec\u006f\u0064\u0065\u0072\u0020\u0063\u006f\u006et\u0065\u0078\u0074\u0020\u0074\u006f\u0020\u0027\u0077\u0027\u0020\u0077r\u0069\u0074\u0065\u0072")
	}
	return int(_ccfa), nil
}
func _dfe(_eeba int) int {
	if _eeba == 0 {
		return 0
	}
	_eeba |= _eeba >> 1
	_eeba |= _eeba >> 2
	_eeba |= _eeba >> 4
	_eeba |= _eeba >> 8
	_eeba |= _eeba >> 16
	return (_eeba + 1) >> 1
}
func (_aeaga *SymbolDictionary) setAtPixels() error {
	if _aeaga.IsHuffmanEncoded {
		return nil
	}
	_aafe := 1
	if _aeaga.SdTemplate == 0 {
		_aafe = 4
	}
	if _agcd := _aeaga.readAtPixels(_aafe); _agcd != nil {
		return _agcd
	}
	return nil
}
func NewGenericRegion(r _fg.StreamReader) *GenericRegion {
	return &GenericRegion{RegionSegment: NewRegionSegment(r), _baf: r}
}
func (_fafdf *SymbolDictionary) setCodingStatistics() error {
	if _fafdf._eegg == nil {
		_fafdf._eegg = _b.NewStats(512, 1)
	}
	if _fafdf._aabg == nil {
		_fafdf._aabg = _b.NewStats(512, 1)
	}
	if _fafdf._aafa == nil {
		_fafdf._aafa = _b.NewStats(512, 1)
	}
	if _fafdf._eagcd == nil {
		_fafdf._eagcd = _b.NewStats(512, 1)
	}
	if _fafdf._gabe == nil {
		_fafdf._gabe = _b.NewStats(512, 1)
	}
	if _fafdf.UseRefinementAggregation && _fafdf._fac == nil {
		_fafdf._fac = _b.NewStats(1<<uint(_fafdf._cff), 1)
		_fafdf._abag = _b.NewStats(512, 1)
		_fafdf._dfb = _b.NewStats(512, 1)
	}
	if _fafdf._bagf == nil {
		_fafdf._bagf = _b.NewStats(65536, 1)
	}
	if _fafdf._fffd == nil {
		var _aefgg error
		_fafdf._fffd, _aefgg = _b.New(_fafdf._bcagf)
		if _aefgg != nil {
			return _aefgg
		}
	}
	return nil
}
func (_egcb *PageInformationSegment) parseHeader() (_dbaa error) {
	_da.Log.Trace("\u005b\u0050\u0061\u0067\u0065I\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065\u0067m\u0065\u006e\u0074\u005d\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065\u0061\u0064\u0065\u0072\u002e\u002e\u002e")
	defer func() {
		var _dbga = "[\u0050\u0061\u0067\u0065\u0049\u006e\u0066\u006f\u0072m\u0061\u0074\u0069\u006f\u006e\u0053\u0065gm\u0065\u006e\u0074\u005d \u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065ad\u0065\u0072 \u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064"
		if _dbaa != nil {
			_dbga += "\u0020\u0077\u0069t\u0068\u0020\u0065\u0072\u0072\u006f\u0072\u0020" + _dbaa.Error()
		} else {
			_dbga += "\u0020\u0073\u0075\u0063\u0063\u0065\u0073\u0073\u0066\u0075\u006c\u006c\u0079"
		}
		_da.Log.Trace(_dbga)
	}()
	if _dbaa = _egcb.readWidthAndHeight(); _dbaa != nil {
		return _dbaa
	}
	if _dbaa = _egcb.readResolution(); _dbaa != nil {
		return _dbaa
	}
	_, _dbaa = _egcb._ace.ReadBit()
	if _dbaa != nil {
		return _dbaa
	}
	if _dbaa = _egcb.readCombinationOperatorOverrideAllowed(); _dbaa != nil {
		return _dbaa
	}
	if _dbaa = _egcb.readRequiresAuxiliaryBuffer(); _dbaa != nil {
		return _dbaa
	}
	if _dbaa = _egcb.readCombinationOperator(); _dbaa != nil {
		return _dbaa
	}
	if _dbaa = _egcb.readDefaultPixelValue(); _dbaa != nil {
		return _dbaa
	}
	if _dbaa = _egcb.readContainsRefinement(); _dbaa != nil {
		return _dbaa
	}
	if _dbaa = _egcb.readIsLossless(); _dbaa != nil {
		return _dbaa
	}
	if _dbaa = _egcb.readIsStriped(); _dbaa != nil {
		return _dbaa
	}
	if _dbaa = _egcb.readMaxStripeSize(); _dbaa != nil {
		return _dbaa
	}
	if _dbaa = _egcb.checkInput(); _dbaa != nil {
		return _dbaa
	}
	_da.Log.Trace("\u0025\u0073", _egcb)
	return nil
}
func (_abc *template1) form(_gbf, _dge, _gbd, _fegd, _efe int16) int16 {
	return ((_gbf & 0x02) << 8) | (_dge << 6) | ((_gbd & 0x03) << 4) | (_fegd << 1) | _efe
}
func (_gggf *TextRegion) decodeRdx() (int64, error) {
	const _ccgea = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064x"
	if _gggf.IsHuffmanEncoded {
		if _gggf.SbHuffRDX == 3 {
			if _gggf._agag == nil {
				var (
					_edbb  int
					_cabfa error
				)
				if _gggf.SbHuffFS == 3 {
					_edbb++
				}
				if _gggf.SbHuffDS == 3 {
					_edbb++
				}
				if _gggf.SbHuffDT == 3 {
					_edbb++
				}
				if _gggf.SbHuffRDWidth == 3 {
					_edbb++
				}
				if _gggf.SbHuffRDHeight == 3 {
					_edbb++
				}
				_gggf._agag, _cabfa = _gggf.getUserTable(_edbb)
				if _cabfa != nil {
					return 0, _fd.Wrap(_cabfa, _ccgea, "")
				}
			}
			return _gggf._agag.Decode(_gggf._effb)
		}
		_fcff, _gfb := _ba.GetStandardTable(14 + int(_gggf.SbHuffRDX))
		if _gfb != nil {
			return 0, _fd.Wrap(_gfb, _ccgea, "")
		}
		return _fcff.Decode(_gggf._effb)
	}
	_fdeed, _adge := _gggf._dffe.DecodeInt(_gggf._beda)
	if _adge != nil {
		return 0, _fd.Wrap(_adge, _ccgea, "")
	}
	return int64(_fdeed), nil
}
func (_daecb *GenericRegion) writeGBAtPixels(_fecd _fg.BinaryWriter) (_eed int, _fggg error) {
	const _dadd = "\u0077r\u0069t\u0065\u0047\u0042\u0041\u0074\u0050\u0069\u0078\u0065\u006c\u0073"
	if _daecb.UseMMR {
		return 0, nil
	}
	_edg := 1
	if _daecb.GBTemplate == 0 {
		_edg = 4
	} else if _daecb.UseExtTemplates {
		_edg = 12
	}
	if len(_daecb.GBAtX) != _edg {
		return 0, _fd.Errorf(_dadd, "\u0067\u0062\u0020\u0061\u0074\u0020\u0070\u0061\u0069\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020d\u006f\u0065\u0073\u006e\u0027\u0074\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u006f\u0020\u0047\u0042\u0041t\u0058\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u006c\u0065\u006e")
	}
	if len(_daecb.GBAtY) != _edg {
		return 0, _fd.Errorf(_dadd, "\u0067\u0062\u0020\u0061\u0074\u0020\u0070\u0061\u0069\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020d\u006f\u0065\u0073\u006e\u0027\u0074\u0020m\u0061\u0074\u0063\u0068\u0020\u0074\u006f\u0020\u0047\u0042\u0041t\u0059\u0020\u0073\u006c\u0069\u0063\u0065\u0020\u006c\u0065\u006e")
	}
	for _acacc := 0; _acacc < _edg; _acacc++ {
		if _fggg = _fecd.WriteByte(byte(_daecb.GBAtX[_acacc])); _fggg != nil {
			return _eed, _fd.Wrap(_fggg, _dadd, "w\u0072\u0069\u0074\u0065\u0020\u0047\u0042\u0041\u0074\u0058")
		}
		_eed++
		if _fggg = _fecd.WriteByte(byte(_daecb.GBAtY[_acacc])); _fggg != nil {
			return _eed, _fd.Wrap(_fggg, _dadd, "w\u0072\u0069\u0074\u0065\u0020\u0047\u0042\u0041\u0074\u0059")
		}
		_eed++
	}
	return _eed, nil
}

type template0 struct{}

func NewHeader(d Documenter, r _fg.StreamReader, offset int64, organizationType OrganizationType) (*Header, error) {
	_ged := &Header{Reader: r}
	if _ccab := _ged.parse(d, r, offset, organizationType); _ccab != nil {
		return nil, _fd.Wrap(_ccab, "\u004ee\u0077\u0048\u0065\u0061\u0064\u0065r", "")
	}
	return _ged, nil
}
func (_ff *EndOfStripe) LineNumber() int { return _ff._dc }
func (_fb *GenericRegion) Size() int     { return _fb.RegionSegment.Size() + 1 + 2*len(_fb.GBAtX) }
func (_dbggg *PageInformationSegment) readContainsRefinement() error {
	_dcf, _egbf := _dbggg._ace.ReadBit()
	if _egbf != nil {
		return _egbf
	}
	if _dcf == 1 {
		_dbggg._bgeb = true
	}
	return nil
}

var _ SegmentEncoder = &RegionSegment{}

func (_cgeb *PageInformationSegment) encodeFlags(_ddfb _fg.BinaryWriter) (_ggcad error) {
	const _bace = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _ggcad = _ddfb.SkipBits(1); _ggcad != nil {
		return _fd.Wrap(_ggcad, _bace, "\u0072\u0065\u0073e\u0072\u0076\u0065\u0064\u0020\u0062\u0069\u0074")
	}
	var _dfdb int
	if _cgeb.CombinationOperatorOverrideAllowed() {
		_dfdb = 1
	}
	if _ggcad = _ddfb.WriteBit(_dfdb); _ggcad != nil {
		return _fd.Wrap(_ggcad, _bace, "\u0063\u006f\u006db\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0070\u0065\u0072a\u0074\u006f\u0072\u0020\u006f\u0076\u0065\u0072\u0072\u0069\u0064\u0064\u0065\u006e")
	}
	_dfdb = 0
	if _cgeb._bfgd {
		_dfdb = 1
	}
	if _ggcad = _ddfb.WriteBit(_dfdb); _ggcad != nil {
		return _fd.Wrap(_ggcad, _bace, "\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0061\u0075\u0078i\u006ci\u0061r\u0079\u0020\u0062\u0075\u0066\u0066\u0065r")
	}
	if _ggcad = _ddfb.WriteBit((int(_cgeb._aefgc) >> 1) & 0x01); _ggcad != nil {
		return _fd.Wrap(_ggcad, _bace, "\u0063\u006f\u006d\u0062\u0069\u006e\u0061\u0074\u0069\u006fn\u0020\u006f\u0070\u0065\u0072\u0061\u0074o\u0072\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0062\u0069\u0074")
	}
	if _ggcad = _ddfb.WriteBit(int(_cgeb._aefgc) & 0x01); _ggcad != nil {
		return _fd.Wrap(_ggcad, _bace, "\u0063\u006f\u006db\u0069\u006e\u0061\u0074i\u006f\u006e\u0020\u006f\u0070\u0065\u0072a\u0074\u006f\u0072\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0062\u0069\u0074")
	}
	_dfdb = int(_cgeb.DefaultPixelValue)
	if _ggcad = _ddfb.WriteBit(_dfdb); _ggcad != nil {
		return _fd.Wrap(_ggcad, _bace, "\u0064e\u0066\u0061\u0075\u006c\u0074\u0020\u0070\u0061\u0067\u0065\u0020p\u0069\u0078\u0065\u006c\u0020\u0076\u0061\u006c\u0075\u0065")
	}
	_dfdb = 0
	if _cgeb._bgeb {
		_dfdb = 1
	}
	if _ggcad = _ddfb.WriteBit(_dfdb); _ggcad != nil {
		return _fd.Wrap(_ggcad, _bace, "\u0063\u006f\u006e\u0074ai\u006e\u0073\u0020\u0072\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074")
	}
	_dfdb = 0
	if _cgeb.IsLossless {
		_dfdb = 1
	}
	if _ggcad = _ddfb.WriteBit(_dfdb); _ggcad != nil {
		return _fd.Wrap(_ggcad, _bace, "p\u0061\u0067\u0065\u0020\u0069\u0073 \u0065\u0076\u0065\u006e\u0074\u0075\u0061\u006c\u006cy\u0020\u006c\u006fs\u0073l\u0065\u0073\u0073")
	}
	return nil
}
func (_dedd *Header) readSegmentDataLength(_fgcb _fg.StreamReader) (_dgag error) {
	_dedd.SegmentDataLength, _dgag = _fgcb.ReadBits(32)
	if _dgag != nil {
		return _dgag
	}
	_dedd.SegmentDataLength &= _f.MaxInt32
	if _dedd.SegmentDataLength > _fgcb.Length() {
		_dedd.SegmentDataLength = _fgcb.Length()
	}
	return nil
}
func (_eggd *TableSegment) HtOOB() int32 { return _eggd._eaed }
func (_ggdd *HalftoneRegion) combineGrayscalePlanes(_eeeg []*_d.Bitmap, _dfa int) error {
	_bgb := 0
	for _bbd := 0; _bbd < _eeeg[_dfa].Height; _bbd++ {
		for _ffbe := 0; _ffbe < _eeeg[_dfa].Width; _ffbe += 8 {
			_gfggc, _feb := _eeeg[_dfa+1].GetByte(_bgb)
			if _feb != nil {
				return _feb
			}
			_ccg, _feb := _eeeg[_dfa].GetByte(_bgb)
			if _feb != nil {
				return _feb
			}
			_feb = _eeeg[_dfa].SetByte(_bgb, _d.CombineBytes(_ccg, _gfggc, _d.CmbOpXor))
			if _feb != nil {
				return _feb
			}
			_bgb++
		}
	}
	return nil
}
func (_cadf *TableSegment) parseHeader() error {
	var (
		_gcafc int
		_bacc  uint64
		_ddebf error
	)
	_gcafc, _ddebf = _cadf._defgf.ReadBit()
	if _ddebf != nil {
		return _ddebf
	}
	if _gcafc == 1 {
		return _e.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0061\u0062\u006c\u0065 \u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0064e\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u002e\u0020\u0042\u002e\u0032\u002e1\u0020\u0043\u006f\u0064\u0065\u0020\u0054\u0061\u0062\u006c\u0065\u0020\u0066\u006c\u0061\u0067\u0073\u003a\u0020\u0042\u0069\u0074\u0020\u0037\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u007a\u0065\u0072\u006f\u002e\u0020\u0057a\u0073\u003a \u0025\u0064", _gcafc)
	}
	if _bacc, _ddebf = _cadf._defgf.ReadBits(3); _ddebf != nil {
		return _ddebf
	}
	_cadf._cgeg = (int32(_bacc) + 1) & 0xf
	if _bacc, _ddebf = _cadf._defgf.ReadBits(3); _ddebf != nil {
		return _ddebf
	}
	_cadf._fbdc = (int32(_bacc) + 1) & 0xf
	if _bacc, _ddebf = _cadf._defgf.ReadBits(32); _ddebf != nil {
		return _ddebf
	}
	_cadf._cccb = int32(_bacc & _f.MaxInt32)
	if _bacc, _ddebf = _cadf._defgf.ReadBits(32); _ddebf != nil {
		return _ddebf
	}
	_cadf._efgd = int32(_bacc & _f.MaxInt32)
	return nil
}
func (_ceec *TextRegion) GetRegionBitmap() (*_d.Bitmap, error) {
	if _ceec.RegionBitmap != nil {
		return _ceec.RegionBitmap, nil
	}
	if !_ceec.IsHuffmanEncoded {
		if _baff := _ceec.setCodingStatistics(); _baff != nil {
			return nil, _baff
		}
	}
	if _cebd := _ceec.createRegionBitmap(); _cebd != nil {
		return nil, _cebd
	}
	if _dadg := _ceec.decodeSymbolInstances(); _dadg != nil {
		return nil, _dadg
	}
	return _ceec.RegionBitmap, nil
}

var (
	_ Regioner  = &TextRegion{}
	_ Segmenter = &TextRegion{}
)

func (_aaed *GenericRegion) computeSegmentDataStructure() error {
	_aaed.DataOffset = _aaed._baf.StreamPosition()
	_aaed.DataHeaderLength = _aaed.DataOffset - _aaed.DataHeaderOffset
	_aaed.DataLength = int64(_aaed._baf.Length()) - _aaed.DataHeaderLength
	return nil
}
func (_badd *Header) pageSize() uint {
	if _badd.PageAssociation <= 255 {
		return 1
	}
	return 4
}
func (_bfac *TextRegion) decodeDfs() (int64, error) {
	if _bfac.IsHuffmanEncoded {
		if _bfac.SbHuffFS == 3 {
			if _bfac._bcf == nil {
				var _agee error
				_bfac._bcf, _agee = _bfac.getUserTable(0)
				if _agee != nil {
					return 0, _agee
				}
			}
			return _bfac._bcf.Decode(_bfac._effb)
		}
		_fcfg, _aagg := _ba.GetStandardTable(6 + int(_bfac.SbHuffFS))
		if _aagg != nil {
			return 0, _aagg
		}
		return _fcfg.Decode(_bfac._effb)
	}
	_ddge, _gdaf := _bfac._dffe.DecodeInt(_bfac._geaa)
	if _gdaf != nil {
		return 0, _gdaf
	}
	return int64(_ddge), nil
}
func _cda(_bcef _fg.StreamReader, _geged *Header) *GenericRefinementRegion {
	return &GenericRefinementRegion{_geg: _bcef, RegionInfo: NewRegionSegment(_bcef), _cdg: _geged, _dae: &template0{}, _dec: &template1{}}
}
func (_aaff *SymbolDictionary) getToExportFlags() ([]int, error) {
	var (
		_cded  int
		_cdde  int32
		_febc  error
		_abbff = int32(_aaff._gbbg + _aaff.NumberOfNewSymbols)
		_fcdb  = make([]int, _abbff)
	)
	for _effa := int32(0); _effa < _abbff; _effa += _cdde {
		if _aaff.IsHuffmanEncoded {
			_fbbc, _bdd := _ba.GetStandardTable(1)
			if _bdd != nil {
				return nil, _bdd
			}
			_gbgg, _bdd := _fbbc.Decode(_aaff._bcagf)
			if _bdd != nil {
				return nil, _bdd
			}
			_cdde = int32(_gbgg)
		} else {
			_cdde, _febc = _aaff._fffd.DecodeInt(_aaff._gabe)
			if _febc != nil {
				return nil, _febc
			}
		}
		if _cdde != 0 {
			for _eadb := _effa; _eadb < _effa+_cdde; _eadb++ {
				_fcdb[_eadb] = _cded
			}
		}
		if _cded == 0 {
			_cded = 1
		} else {
			_cded = 0
		}
	}
	return _fcdb, nil
}
func (_fccgf *SymbolDictionary) huffDecodeRefAggNInst() (int64, error) {
	if !_fccgf.SdHuffAggInstanceSelection {
		_ddca, _afcg := _ba.GetStandardTable(1)
		if _afcg != nil {
			return 0, _afcg
		}
		return _ddca.Decode(_fccgf._bcagf)
	}
	if _fccgf._fggec == nil {
		var (
			_gfde int
			_abfa error
		)
		if _fccgf.SdHuffDecodeHeightSelection == 3 {
			_gfde++
		}
		if _fccgf.SdHuffDecodeWidthSelection == 3 {
			_gfde++
		}
		if _fccgf.SdHuffBMSizeSelection == 3 {
			_gfde++
		}
		_fccgf._fggec, _abfa = _fccgf.getUserTable(_gfde)
		if _abfa != nil {
			return 0, _abfa
		}
	}
	return _fccgf._fggec.Decode(_fccgf._bcagf)
}
func (_gegda *TextRegion) String() string {
	_cdaag := &_cc.Builder{}
	_cdaag.WriteString("\u000a[\u0054E\u0058\u0054\u0020\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u000a")
	_cdaag.WriteString(_gegda.RegionInfo.String() + "\u000a")
	_cdaag.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053br\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0025\u0076\u000a", _gegda.SbrTemplate))
	_cdaag.WriteString(_e.Sprintf("\u0009-\u0020S\u0062\u0044\u0073\u004f\u0066f\u0073\u0065t\u003a\u0020\u0025\u0076\u000a", _gegda.SbDsOffset))
	_cdaag.WriteString(_e.Sprintf("\t\u002d \u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0050i\u0078\u0065\u006c\u003a %\u0076\u000a", _gegda.DefaultPixel))
	_cdaag.WriteString(_e.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065r\u0061\u0074\u006fr\u003a \u0025\u0076\u000a", _gegda.CombinationOperator))
	_cdaag.WriteString(_e.Sprintf("\t\u002d \u0049\u0073\u0054\u0072\u0061\u006e\u0073\u0070o\u0073\u0065\u0064\u003a %\u0076\u000a", _gegda.IsTransposed))
	_cdaag.WriteString(_e.Sprintf("\u0009\u002d\u0020Re\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0043\u006f\u0072\u006e\u0065\u0072\u003a\u0020\u0025\u0076\u000a", _gegda.ReferenceCorner))
	_cdaag.WriteString(_e.Sprintf("\t\u002d\u0020\u0055\u0073eR\u0065f\u0069\u006e\u0065\u006d\u0065n\u0074\u003a\u0020\u0025\u0076\u000a", _gegda.UseRefinement))
	_cdaag.WriteString(_e.Sprintf("\u0009-\u0020\u0049\u0073\u0048\u0075\u0066\u0066\u006d\u0061\u006e\u0045n\u0063\u006f\u0064\u0065\u0064\u003a\u0020\u0025\u0076\u000a", _gegda.IsHuffmanEncoded))
	if _gegda.IsHuffmanEncoded {
		_cdaag.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053bH\u0075\u0066\u0066\u0052\u0053\u0069\u007a\u0065\u003a\u0020\u0025\u0076\u000a", _gegda.SbHuffRSize))
		_cdaag.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0059:\u0020\u0025\u0076\u000a", _gegda.SbHuffRDY))
		_cdaag.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0058:\u0020\u0025\u0076\u000a", _gegda.SbHuffRDX))
		_cdaag.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053bH\u0075\u0066\u0066\u0052\u0044\u0048\u0065\u0069\u0067\u0068\u0074\u003a\u0020\u0025v\u000a", _gegda.SbHuffRDHeight))
		_cdaag.WriteString(_e.Sprintf("\t\u002d\u0020\u0053\u0062Hu\u0066f\u0052\u0044\u0057\u0069\u0064t\u0068\u003a\u0020\u0025\u0076\u000a", _gegda.SbHuffRDWidth))
		_cdaag.WriteString(_e.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0044\u0054\u003a\u0020\u0025\u0076\u000a", _gegda.SbHuffDT))
		_cdaag.WriteString(_e.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0044\u0053\u003a\u0020\u0025\u0076\u000a", _gegda.SbHuffDS))
		_cdaag.WriteString(_e.Sprintf("\u0009\u002d \u0053\u0062\u0048u\u0066\u0066\u0046\u0053\u003a\u0020\u0025\u0076\u000a", _gegda.SbHuffFS))
	}
	_cdaag.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0058:\u0020\u0025\u0076\u000a", _gegda.SbrATX))
	_cdaag.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0059:\u0020\u0025\u0076\u000a", _gegda.SbrATY))
	_cdaag.WriteString(_e.Sprintf("\u0009\u002d\u0020N\u0075\u006d\u0062\u0065r\u004f\u0066\u0053\u0079\u006d\u0062\u006fl\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u003a\u0020\u0025\u0076\u000a", _gegda.NumberOfSymbolInstances))
	_cdaag.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053\u0062\u0072\u0041\u0054\u0058:\u0020\u0025\u0076\u000a", _gegda.SbrATX))
	return _cdaag.String()
}
func (_dfee Type) String() string {
	switch _dfee {
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

const (
	ORandom OrganizationType = iota
	OSequential
)

func (_acee *SymbolDictionary) String() string {
	_bccb := &_cc.Builder{}
	_bccb.WriteString("\n\u005b\u0053\u0059\u004dBO\u004c-\u0044\u0049\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u000a")
	_bccb.WriteString(_e.Sprintf("\u0009-\u0020S\u0064\u0072\u0054\u0065\u006dp\u006c\u0061t\u0065\u0020\u0025\u0076\u000a", _acee.SdrTemplate))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u0025\u0076\u000a", _acee.SdTemplate))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0069\u0073\u0043\u006f\u0064\u0069\u006eg\u0043\u006f\u006e\u0074\u0065\u0078\u0074R\u0065\u0074\u0061\u0069\u006e\u0065\u0064\u0020\u0025\u0076\u000a", _acee._gdae))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0069\u0073\u0043\u006f\u0064\u0069\u006e\u0067C\u006f\u006e\u0074\u0065\u0078\u0074\u0055\u0073\u0065\u0064 \u0025\u0076\u000a", _acee._gacg))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0048u\u0066\u0066\u0041\u0067\u0067\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065S\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0025\u0076\u000a", _acee.SdHuffAggInstanceSelection))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053d\u0048\u0075\u0066\u0066\u0042\u004d\u0053\u0069\u007a\u0065S\u0065l\u0065\u0063\u0074\u0069\u006f\u006e\u0020%\u0076\u000a", _acee.SdHuffBMSizeSelection))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0048u\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0057\u0069\u0064\u0074\u0068S\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0025\u0076\u000a", _acee.SdHuffDecodeWidthSelection))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020Sd\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064e\u0048e\u0069g\u0068t\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u0020\u0025\u0076\u000a", _acee.SdHuffDecodeHeightSelection))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020U\u0073\u0065\u0052\u0065f\u0069\u006e\u0065\u006d\u0065\u006e\u0074A\u0067\u0067\u0072\u0065\u0067\u0061\u0074\u0069\u006f\u006e\u0020\u0025\u0076\u000a", _acee.UseRefinementAggregation))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020is\u0048\u0075\u0066\u0066\u006d\u0061\u006e\u0045\u006e\u0063\u006f\u0064\u0065\u0064\u0020\u0025\u0076\u000a", _acee.IsHuffmanEncoded))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020S\u0064\u0041\u0054\u0058\u0020\u0025\u0076\u000a", _acee.SdATX))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020S\u0064\u0041\u0054\u0059\u0020\u0025\u0076\u000a", _acee.SdATY))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0072\u0041\u0054\u0058\u0020\u0025\u0076\u000a", _acee.SdrATX))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053\u0064\u0072\u0041\u0054\u0059\u0020\u0025\u0076\u000a", _acee.SdrATY))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u004ff\u0045\u0078\u0070\u006f\u0072\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0025\u0076\u000a", _acee.NumberOfExportedSymbols))
	_bccb.WriteString(_e.Sprintf("\u0009-\u0020\u004e\u0075\u006db\u0065\u0072\u004f\u0066\u004ee\u0077S\u0079m\u0062\u006f\u006c\u0073\u0020\u0025\u0076\n", _acee.NumberOfNewSymbols))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u004ff\u0049\u006d\u0070\u006f\u0072\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0025\u0076\u000a", _acee._gbbg))
	_bccb.WriteString(_e.Sprintf("\u0009\u002d \u006e\u0075\u006d\u0062\u0065\u0072\u004f\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0064\u0053\u0079\u006d\u0062\u006f\u006c\u0073 %\u0076\u000a", _acee._ffa))
	return _bccb.String()
}
func (_dgca *SymbolDictionary) Init(h *Header, r _fg.StreamReader) error {
	_dgca.Header = h
	_dgca._bcagf = r
	return _dgca.parseHeader()
}
func (_cdaa *Header) String() string {
	_gfd := &_cc.Builder{}
	_gfd.WriteString("\u000a[\u0053E\u0047\u004d\u0045\u004e\u0054-\u0048\u0045A\u0044\u0045\u0052\u005d\u000a")
	_gfd.WriteString(_e.Sprintf("\t\u002d\u0020\u0053\u0065gm\u0065n\u0074\u004e\u0075\u006d\u0062e\u0072\u003a\u0020\u0025\u0076\u000a", _cdaa.SegmentNumber))
	_gfd.WriteString(_e.Sprintf("\u0009\u002d\u0020T\u0079\u0070\u0065\u003a\u0020\u0025\u0076\u000a", _cdaa.Type))
	_gfd.WriteString(_e.Sprintf("\u0009-\u0020R\u0065\u0074\u0061\u0069\u006eF\u006c\u0061g\u003a\u0020\u0025\u0076\u000a", _cdaa.RetainFlag))
	_gfd.WriteString(_e.Sprintf("\u0009\u002d\u0020Pa\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _cdaa.PageAssociation))
	_gfd.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0050\u0061\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061\u0074i\u006fn\u0046\u0069\u0065\u006c\u0064\u0053\u0069\u007a\u0065\u003a\u0020\u0025\u0076\u000a", _cdaa.PageAssociationFieldSize))
	_gfd.WriteString("\u0009-\u0020R\u0054\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u0053\u003a\u000a")
	for _, _daegg := range _cdaa.RTSNumbers {
		_gfd.WriteString(_e.Sprintf("\u0009\t\u002d\u0020\u0025\u0064\u000a", _daegg))
	}
	_gfd.WriteString(_e.Sprintf("\t\u002d \u0048\u0065\u0061\u0064\u0065\u0072\u004c\u0065n\u0067\u0074\u0068\u003a %\u0076\u000a", _cdaa.HeaderLength))
	_gfd.WriteString(_e.Sprintf("\u0009-\u0020\u0053\u0065\u0067m\u0065\u006e\u0074\u0044\u0061t\u0061L\u0065n\u0067\u0074\u0068\u003a\u0020\u0025\u0076\n", _cdaa.SegmentDataLength))
	_gfd.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074D\u0061\u0074\u0061\u0053\u0074\u0061\u0072t\u004f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0076\u000a", _cdaa.SegmentDataStartOffset))
	return _gfd.String()
}
func (_fdf *SymbolDictionary) decodeHeightClassBitmap(_fbad *_d.Bitmap, _gdfc int64, _aded int, _caecc []int) error {
	for _abaa := _gdfc; _abaa < int64(_fdf._ffa); _abaa++ {
		var _feeg int
		for _gdgg := _gdfc; _gdgg <= _abaa-1; _gdgg++ {
			_feeg += _caecc[_gdgg]
		}
		_gecd := _c.Rect(_feeg, 0, _feeg+_caecc[_abaa], _aded)
		_dgfab, _egab := _d.Extract(_gecd, _fbad)
		if _egab != nil {
			return _egab
		}
		_fdf._gbga[_abaa] = _dgfab
		_fdf._bded = append(_fdf._bded, _dgfab)
	}
	return nil
}
func (_egdd *SymbolDictionary) readRegionFlags() error {
	var (
		_abdf uint64
		_cbcb int
	)
	_, _fcag := _egdd._bcagf.ReadBits(3)
	if _fcag != nil {
		return _fcag
	}
	_cbcb, _fcag = _egdd._bcagf.ReadBit()
	if _fcag != nil {
		return _fcag
	}
	_egdd.SdrTemplate = int8(_cbcb)
	_abdf, _fcag = _egdd._bcagf.ReadBits(2)
	if _fcag != nil {
		return _fcag
	}
	_egdd.SdTemplate = int8(_abdf & 0xf)
	_cbcb, _fcag = _egdd._bcagf.ReadBit()
	if _fcag != nil {
		return _fcag
	}
	if _cbcb == 1 {
		_egdd._gdae = true
	}
	_cbcb, _fcag = _egdd._bcagf.ReadBit()
	if _fcag != nil {
		return _fcag
	}
	if _cbcb == 1 {
		_egdd._gacg = true
	}
	_cbcb, _fcag = _egdd._bcagf.ReadBit()
	if _fcag != nil {
		return _fcag
	}
	if _cbcb == 1 {
		_egdd.SdHuffAggInstanceSelection = true
	}
	_cbcb, _fcag = _egdd._bcagf.ReadBit()
	if _fcag != nil {
		return _fcag
	}
	_egdd.SdHuffBMSizeSelection = int8(_cbcb)
	_abdf, _fcag = _egdd._bcagf.ReadBits(2)
	if _fcag != nil {
		return _fcag
	}
	_egdd.SdHuffDecodeWidthSelection = int8(_abdf & 0xf)
	_abdf, _fcag = _egdd._bcagf.ReadBits(2)
	if _fcag != nil {
		return _fcag
	}
	_egdd.SdHuffDecodeHeightSelection = int8(_abdf & 0xf)
	_cbcb, _fcag = _egdd._bcagf.ReadBit()
	if _fcag != nil {
		return _fcag
	}
	if _cbcb == 1 {
		_egdd.UseRefinementAggregation = true
	}
	_cbcb, _fcag = _egdd._bcagf.ReadBit()
	if _fcag != nil {
		return _fcag
	}
	if _cbcb == 1 {
		_egdd.IsHuffmanEncoded = true
	}
	return nil
}
func (_aca *GenericRegion) decodeTemplate2(_fga, _ggc, _bcc int, _cgg, _afgf int) (_agf error) {
	const _fegf = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0032"
	var (
		_gfgg, _acac int
		_affd, _dcc  int
		_ffg         byte
		_ccd, _efee  int
	)
	if _fga >= 1 {
		_ffg, _agf = _aca.Bitmap.GetByte(_afgf)
		if _agf != nil {
			return _fd.Wrap(_agf, _fegf, "\u006ci\u006ee\u004e\u0075\u006d\u0062\u0065\u0072\u0020\u003e\u003d\u0020\u0031")
		}
		_affd = int(_ffg)
	}
	if _fga >= 2 {
		_ffg, _agf = _aca.Bitmap.GetByte(_afgf - _aca.Bitmap.RowStride)
		if _agf != nil {
			return _fd.Wrap(_agf, _fegf, "\u006ci\u006ee\u004e\u0075\u006d\u0062\u0065\u0072\u0020\u003e\u003d\u0020\u0032")
		}
		_dcc = int(_ffg) << 4
	}
	_gfgg = (_affd >> 3 & 0x7c) | (_dcc >> 3 & 0x380)
	for _egcfd := 0; _egcfd < _bcc; _egcfd = _ccd {
		var (
			_ecdd byte
			_gdcc int
		)
		_ccd = _egcfd + 8
		if _fec := _ggc - _egcfd; _fec > 8 {
			_gdcc = 8
		} else {
			_gdcc = _fec
		}
		if _fga > 0 {
			_affd <<= 8
			if _ccd < _ggc {
				_ffg, _agf = _aca.Bitmap.GetByte(_afgf + 1)
				if _agf != nil {
					return _fd.Wrap(_agf, _fegf, "\u006c\u0069\u006e\u0065\u004e\u0075\u006d\u0062\u0065r\u0020\u003e\u0020\u0030")
				}
				_affd |= int(_ffg)
			}
		}
		if _fga > 1 {
			_dcc <<= 8
			if _ccd < _ggc {
				_ffg, _agf = _aca.Bitmap.GetByte(_afgf - _aca.Bitmap.RowStride + 1)
				if _agf != nil {
					return _fd.Wrap(_agf, _fegf, "\u006c\u0069\u006e\u0065\u004e\u0075\u006d\u0062\u0065r\u0020\u003e\u0020\u0031")
				}
				_dcc |= int(_ffg) << 4
			}
		}
		for _cace := 0; _cace < _gdcc; _cace++ {
			_fda := uint(10 - _cace)
			if _aca._fgf {
				_acac = _aca.overrideAtTemplate2(_gfgg, _egcfd+_cace, _fga, int(_ecdd), _cace)
				_aca._aae.SetIndex(int32(_acac))
			} else {
				_aca._aae.SetIndex(int32(_gfgg))
			}
			_efee, _agf = _aca._acb.DecodeBit(_aca._aae)
			if _agf != nil {
				return _fd.Wrap(_agf, _fegf, "")
			}
			_ecdd |= byte(_efee << uint(7-_cace))
			_gfgg = ((_gfgg & 0x1bd) << 1) | _efee | ((_affd >> _fda) & 0x4) | ((_dcc >> _fda) & 0x80)
		}
		if _ddfc := _aca.Bitmap.SetByte(_cgg, _ecdd); _ddfc != nil {
			return _fd.Wrap(_ddfc, _fegf, "")
		}
		_cgg++
		_afgf++
	}
	return nil
}
func (_afb *TextRegion) decodeRdw() (int64, error) {
	const _bcfd = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064w"
	if _afb.IsHuffmanEncoded {
		if _afb.SbHuffRDWidth == 3 {
			if _afb._gcba == nil {
				var (
					_gdfeb int
					_egae  error
				)
				if _afb.SbHuffFS == 3 {
					_gdfeb++
				}
				if _afb.SbHuffDS == 3 {
					_gdfeb++
				}
				if _afb.SbHuffDT == 3 {
					_gdfeb++
				}
				_afb._gcba, _egae = _afb.getUserTable(_gdfeb)
				if _egae != nil {
					return 0, _fd.Wrap(_egae, _bcfd, "")
				}
			}
			return _afb._gcba.Decode(_afb._effb)
		}
		_bgbcb, _ebgf := _ba.GetStandardTable(14 + int(_afb.SbHuffRDWidth))
		if _ebgf != nil {
			return 0, _fd.Wrap(_ebgf, _bcfd, "")
		}
		return _bgbcb.Decode(_afb._effb)
	}
	_gfgga, _fbbe := _afb._dffe.DecodeInt(_afb._egcff)
	if _fbbe != nil {
		return 0, _fd.Wrap(_fbbe, _bcfd, "")
	}
	return int64(_gfgga), nil
}
func (_dea *HalftoneRegion) GetRegionInfo() *RegionSegment { return _dea.RegionSegment }

var (
	_bfd  Segmenter
	_ffdg = map[Type]func() Segmenter{TSymbolDictionary: func() Segmenter { return &SymbolDictionary{} }, TIntermediateTextRegion: func() Segmenter { return &TextRegion{} }, TImmediateTextRegion: func() Segmenter { return &TextRegion{} }, TImmediateLosslessTextRegion: func() Segmenter { return &TextRegion{} }, TPatternDictionary: func() Segmenter { return &PatternDictionary{} }, TIntermediateHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TImmediateHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TImmediateLosslessHalftoneRegion: func() Segmenter { return &HalftoneRegion{} }, TIntermediateGenericRegion: func() Segmenter { return &GenericRegion{} }, TImmediateGenericRegion: func() Segmenter { return &GenericRegion{} }, TImmediateLosslessGenericRegion: func() Segmenter { return &GenericRegion{} }, TIntermediateGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TImmediateGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TImmediateLosslessGenericRefinementRegion: func() Segmenter { return &GenericRefinementRegion{} }, TPageInformation: func() Segmenter { return &PageInformationSegment{} }, TEndOfPage: func() Segmenter { return _bfd }, TEndOfStrip: func() Segmenter { return &EndOfStripe{} }, TEndOfFile: func() Segmenter { return _bfd }, TProfiles: func() Segmenter { return _bfd }, TTables: func() Segmenter { return &TableSegment{} }, TExtension: func() Segmenter { return _bfd }, TBitmap: func() Segmenter { return _bfd }}
)

func (_caef *SymbolDictionary) readNumberOfExportedSymbols() error {
	_caed, _fcac := _caef._bcagf.ReadBits(32)
	if _fcac != nil {
		return _fcac
	}
	_caef.NumberOfExportedSymbols = uint32(_caed & _f.MaxUint32)
	return nil
}
func (_fdae *HalftoneRegion) checkInput() error {
	if _fdae.IsMMREncoded {
		if _fdae.HTemplate != 0 {
			_da.Log.Debug("\u0048\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0030", _fdae.HTemplate)
		}
		if _fdae.HSkipEnabled {
			_da.Log.Debug("\u0048\u0053\u006b\u0069\u0070\u0045\u006e\u0061\u0062\u006c\u0065\u0064\u0020\u0030\u0020\u0025\u0076\u0020(\u0073\u0068\u006f\u0075\u006c\u0064\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0074\u0068\u0065\u0020v\u0061\u006c\u0075\u0065\u0020\u0066\u0061\u006c\u0073\u0065\u0029", _fdae.HSkipEnabled)
		}
	}
	return nil
}
func (_geb *GenericRegion) overrideAtTemplate0a(_gfgb, _fcc, _gfeb, _bbcb, _bedg, _dddd int) int {
	if _geb.GBAtOverride[0] {
		_gfgb &= 0xFFEF
		if _geb.GBAtY[0] == 0 && _geb.GBAtX[0] >= -int8(_bedg) {
			_gfgb |= (_bbcb >> uint(int8(_dddd)-_geb.GBAtX[0]&0x1)) << 4
		} else {
			_gfgb |= int(_geb.getPixel(_fcc+int(_geb.GBAtX[0]), _gfeb+int(_geb.GBAtY[0]))) << 4
		}
	}
	if _geb.GBAtOverride[1] {
		_gfgb &= 0xFBFF
		if _geb.GBAtY[1] == 0 && _geb.GBAtX[1] >= -int8(_bedg) {
			_gfgb |= (_bbcb >> uint(int8(_dddd)-_geb.GBAtX[1]&0x1)) << 10
		} else {
			_gfgb |= int(_geb.getPixel(_fcc+int(_geb.GBAtX[1]), _gfeb+int(_geb.GBAtY[1]))) << 10
		}
	}
	if _geb.GBAtOverride[2] {
		_gfgb &= 0xF7FF
		if _geb.GBAtY[2] == 0 && _geb.GBAtX[2] >= -int8(_bedg) {
			_gfgb |= (_bbcb >> uint(int8(_dddd)-_geb.GBAtX[2]&0x1)) << 11
		} else {
			_gfgb |= int(_geb.getPixel(_fcc+int(_geb.GBAtX[2]), _gfeb+int(_geb.GBAtY[2]))) << 11
		}
	}
	if _geb.GBAtOverride[3] {
		_gfgb &= 0x7FFF
		if _geb.GBAtY[3] == 0 && _geb.GBAtX[3] >= -int8(_bedg) {
			_gfgb |= (_bbcb >> uint(int8(_dddd)-_geb.GBAtX[3]&0x1)) << 15
		} else {
			_gfgb |= int(_geb.getPixel(_fcc+int(_geb.GBAtX[3]), _gfeb+int(_geb.GBAtY[3]))) << 15
		}
	}
	return _gfgb
}
func (_ddcg *PageInformationSegment) Encode(w _fg.BinaryWriter) (_cddf int, _fadc error) {
	const _adebe = "\u0050\u0061g\u0065\u0049\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u002e\u0045\u006eco\u0064\u0065"
	_degbc := make([]byte, 4)
	_fe.BigEndian.PutUint32(_degbc, uint32(_ddcg.PageBMWidth))
	_cddf, _fadc = w.Write(_degbc)
	if _fadc != nil {
		return _cddf, _fd.Wrap(_fadc, _adebe, "\u0077\u0069\u0064t\u0068")
	}
	_fe.BigEndian.PutUint32(_degbc, uint32(_ddcg.PageBMHeight))
	var _caec int
	_caec, _fadc = w.Write(_degbc)
	if _fadc != nil {
		return _caec + _cddf, _fd.Wrap(_fadc, _adebe, "\u0068\u0065\u0069\u0067\u0068\u0074")
	}
	_cddf += _caec
	_fe.BigEndian.PutUint32(_degbc, uint32(_ddcg.ResolutionX))
	_caec, _fadc = w.Write(_degbc)
	if _fadc != nil {
		return _caec + _cddf, _fd.Wrap(_fadc, _adebe, "\u0078\u0020\u0072e\u0073\u006f\u006c\u0075\u0074\u0069\u006f\u006e")
	}
	_cddf += _caec
	_fe.BigEndian.PutUint32(_degbc, uint32(_ddcg.ResolutionY))
	if _caec, _fadc = w.Write(_degbc); _fadc != nil {
		return _caec + _cddf, _fd.Wrap(_fadc, _adebe, "\u0079\u0020\u0072e\u0073\u006f\u006c\u0075\u0074\u0069\u006f\u006e")
	}
	_cddf += _caec
	if _fadc = _ddcg.encodeFlags(w); _fadc != nil {
		return _cddf, _fd.Wrap(_fadc, _adebe, "")
	}
	_cddf++
	if _caec, _fadc = _ddcg.encodeStripingInformation(w); _fadc != nil {
		return _cddf, _fd.Wrap(_fadc, _adebe, "")
	}
	_cddf += _caec
	return _cddf, nil
}
func (_ddg *GenericRefinementRegion) GetRegionInfo() *RegionSegment { return _ddg.RegionInfo }
func (_ded *HalftoneRegion) GetPatterns() ([]*_d.Bitmap, error) {
	var (
		_cgbe []*_d.Bitmap
		_gee  error
	)
	for _, _gddd := range _ded._fdag.RTSegments {
		var _deg Segmenter
		_deg, _gee = _gddd.GetSegmentData()
		if _gee != nil {
			_da.Log.Debug("\u0047e\u0074\u0053\u0065\u0067m\u0065\u006e\u0074\u0044\u0061t\u0061 \u0066a\u0069\u006c\u0065\u0064\u003a\u0020\u0025v", _gee)
			return nil, _gee
		}
		_cfd, _aefg := _deg.(*PatternDictionary)
		if !_aefg {
			_gee = _e.Errorf("\u0072e\u006c\u0061t\u0065\u0064\u0020\u0073e\u0067\u006d\u0065n\u0074\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0070at\u0074\u0065\u0072n\u0020\u0064i\u0063\u0074\u0069\u006f\u006e\u0061r\u0079\u003a \u0025\u0054", _deg)
			return nil, _gee
		}
		var _dfca []*_d.Bitmap
		_dfca, _gee = _cfd.GetDictionary()
		if _gee != nil {
			_da.Log.Debug("\u0070\u0061\u0074\u0074\u0065\u0072\u006e\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _gee)
			return nil, _gee
		}
		_cgbe = append(_cgbe, _dfca...)
	}
	return _cgbe, nil
}
func (_ecb *template0) form(_gcb, _eea, _aged, _edc, _gae int16) int16 {
	return (_gcb << 10) | (_eea << 7) | (_aged << 4) | (_edc << 1) | _gae
}
func (_eafa *GenericRegion) setParametersMMR(_gcfg bool, _bea, _fcaa int64, _daca, _badff uint32, _ccag byte, _afc, _gcff bool, _cbff, _gffc []int8) {
	_eafa.DataOffset = _bea
	_eafa.DataLength = _fcaa
	_eafa.RegionSegment = &RegionSegment{}
	_eafa.RegionSegment.BitmapHeight = _daca
	_eafa.RegionSegment.BitmapWidth = _badff
	_eafa.GBTemplate = _ccag
	_eafa.IsMMREncoded = _gcfg
	_eafa.IsTPGDon = _afc
	_eafa.GBAtX = _cbff
	_eafa.GBAtY = _gffc
}
func (_efg *template1) setIndex(_agec *_b.DecoderStats) { _agec.SetIndex(0x080) }
func (_bad *GenericRegion) String() string {
	_effg := &_cc.Builder{}
	_effg.WriteString("\u000a[\u0047E\u004e\u0045\u0052\u0049\u0043 \u0052\u0045G\u0049\u004f\u004e\u005d\u000a")
	_effg.WriteString(_bad.RegionSegment.String() + "\u000a")
	_effg.WriteString(_e.Sprintf("\u0009\u002d\u0020Us\u0065\u0045\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073\u003a\u0020\u0025\u0076\u000a", _bad.UseExtTemplates))
	_effg.WriteString(_e.Sprintf("\u0009\u002d \u0049\u0073\u0054P\u0047\u0044\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _bad.IsTPGDon))
	_effg.WriteString(_e.Sprintf("\u0009-\u0020G\u0042\u0054\u0065\u006d\u0070l\u0061\u0074e\u003a\u0020\u0025\u0064\u000a", _bad.GBTemplate))
	_effg.WriteString(_e.Sprintf("\t\u002d \u0049\u0073\u004d\u004d\u0052\u0045\u006e\u0063o\u0064\u0065\u0064\u003a %\u0076\u000a", _bad.IsMMREncoded))
	_effg.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0047\u0042\u0041\u0074\u0058\u003a\u0020\u0025\u0076\u000a", _bad.GBAtX))
	_effg.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0047\u0042\u0041\u0074\u0059\u003a\u0020\u0025\u0076\u000a", _bad.GBAtY))
	_effg.WriteString(_e.Sprintf("\t\u002d \u0047\u0042\u0041\u0074\u004f\u0076\u0065\u0072r\u0069\u0064\u0065\u003a %\u0076\u000a", _bad.GBAtOverride))
	return _effg.String()
}
func (_dca *SymbolDictionary) decodeDifferenceWidth() (int64, error) {
	if _dca.IsHuffmanEncoded {
		switch _dca.SdHuffDecodeWidthSelection {
		case 0:
			_agab, _fdeg := _ba.GetStandardTable(2)
			if _fdeg != nil {
				return 0, _fdeg
			}
			return _agab.Decode(_dca._bcagf)
		case 1:
			_aefe, _ffde := _ba.GetStandardTable(3)
			if _ffde != nil {
				return 0, _ffde
			}
			return _aefe.Decode(_dca._bcagf)
		case 3:
			if _dca._agdg == nil {
				var _fgcbb int
				if _dca.SdHuffDecodeHeightSelection == 3 {
					_fgcbb++
				}
				_fbdb, _cfga := _dca.getUserTable(_fgcbb)
				if _cfga != nil {
					return 0, _cfga
				}
				_dca._agdg = _fbdb
			}
			return _dca._agdg.Decode(_dca._bcagf)
		}
	} else {
		_fegga, _cfge := _dca._fffd.DecodeInt(_dca._aafa)
		if _cfge != nil {
			return 0, _cfge
		}
		return int64(_fegga), nil
	}
	return 0, nil
}
func (_eab *GenericRegion) decodeTemplate3(_addg, _fafc, _gaa int, _aef, _adea int) (_bfa error) {
	const _bccg = "\u0064e\u0063o\u0064\u0065\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0033"
	var (
		_cdbf, _fcg int
		_gfea       int
		_ffe        byte
		_fddc, _fbe int
	)
	if _addg >= 1 {
		_ffe, _bfa = _eab.Bitmap.GetByte(_adea)
		if _bfa != nil {
			return _fd.Wrap(_bfa, _bccg, "\u006ci\u006e\u0065\u0020\u003e\u003d\u00201")
		}
		_gfea = int(_ffe)
	}
	_cdbf = (_gfea >> 1) & 0x70
	for _abbg := 0; _abbg < _gaa; _abbg = _fddc {
		var (
			_ege  byte
			_fedb int
		)
		_fddc = _abbg + 8
		if _afa := _fafc - _abbg; _afa > 8 {
			_fedb = 8
		} else {
			_fedb = _afa
		}
		if _addg >= 1 {
			_gfea <<= 8
			if _fddc < _fafc {
				_ffe, _bfa = _eab.Bitmap.GetByte(_adea + 1)
				if _bfa != nil {
					return _fd.Wrap(_bfa, _bccg, "\u0069\u006e\u006e\u0065\u0072\u0020\u002d\u0020\u006c\u0069\u006e\u0065 \u003e\u003d\u0020\u0031")
				}
				_gfea |= int(_ffe)
			}
		}
		for _fde := 0; _fde < _fedb; _fde++ {
			if _eab._fgf {
				_fcg = _eab.overrideAtTemplate3(_cdbf, _abbg+_fde, _addg, int(_ege), _fde)
				_eab._aae.SetIndex(int32(_fcg))
			} else {
				_eab._aae.SetIndex(int32(_cdbf))
			}
			_fbe, _bfa = _eab._acb.DecodeBit(_eab._aae)
			if _bfa != nil {
				return _fd.Wrap(_bfa, _bccg, "")
			}
			_ege |= byte(_fbe) << byte(7-_fde)
			_cdbf = ((_cdbf & 0x1f7) << 1) | _fbe | ((_gfea >> uint(8-_fde)) & 0x010)
		}
		if _gceg := _eab.Bitmap.SetByte(_aef, _ege); _gceg != nil {
			return _fd.Wrap(_gceg, _bccg, "")
		}
		_aef++
		_adea++
	}
	return nil
}
func (_eeg *GenericRegion) copyLineAbove(_cec int) error {
	_daeg := _cec * _eeg.Bitmap.RowStride
	_bef := _daeg - _eeg.Bitmap.RowStride
	for _fbf := 0; _fbf < _eeg.Bitmap.RowStride; _fbf++ {
		_ecbb, _fff := _eeg.Bitmap.GetByte(_bef)
		if _fff != nil {
			return _fff
		}
		_bef++
		if _fff = _eeg.Bitmap.SetByte(_daeg, _ecbb); _fff != nil {
			return _fff
		}
		_daeg++
	}
	return nil
}
func (_gaee *TextRegion) decodeSymInRefSize() (int64, error) {
	const _fcbd = "\u0064e\u0063o\u0064\u0065\u0053\u0079\u006dI\u006e\u0052e\u0066\u0053\u0069\u007a\u0065"
	if _gaee.SbHuffRSize == 0 {
		_gbbfc, _cgfd := _ba.GetStandardTable(1)
		if _cgfd != nil {
			return 0, _fd.Wrap(_cgfd, _fcbd, "")
		}
		return _gbbfc.Decode(_gaee._effb)
	}
	if _gaee._fbeac == nil {
		var (
			_bege int
			_fgaf error
		)
		if _gaee.SbHuffFS == 3 {
			_bege++
		}
		if _gaee.SbHuffDS == 3 {
			_bege++
		}
		if _gaee.SbHuffDT == 3 {
			_bege++
		}
		if _gaee.SbHuffRDWidth == 3 {
			_bege++
		}
		if _gaee.SbHuffRDHeight == 3 {
			_bege++
		}
		if _gaee.SbHuffRDX == 3 {
			_bege++
		}
		if _gaee.SbHuffRDY == 3 {
			_bege++
		}
		_gaee._fbeac, _fgaf = _gaee.getUserTable(_bege)
		if _fgaf != nil {
			return 0, _fd.Wrap(_fgaf, _fcbd, "")
		}
	}
	_affc, _cgdef := _gaee._fbeac.Decode(_gaee._effb)
	if _cgdef != nil {
		return 0, _fd.Wrap(_cgdef, _fcbd, "")
	}
	return _affc, nil
}
func (_ecfa *HalftoneRegion) grayScaleDecoding(_ggge int) ([][]int, error) {
	var (
		_adce []int8
		_abbc []int8
	)
	if !_ecfa.IsMMREncoded {
		_adce = make([]int8, 4)
		_abbc = make([]int8, 4)
		if _ecfa.HTemplate <= 1 {
			_adce[0] = 3
		} else if _ecfa.HTemplate >= 2 {
			_adce[0] = 2
		}
		_abbc[0] = -1
		_adce[1] = -3
		_abbc[1] = -1
		_adce[2] = 2
		_abbc[2] = -2
		_adce[3] = -2
		_abbc[3] = -2
	}
	_gddg := make([]*_d.Bitmap, _ggge)
	_eced := NewGenericRegion(_ecfa._gdfe)
	_eced.setParametersMMR(_ecfa.IsMMREncoded, _ecfa.DataOffset, _ecfa.DataLength, _ecfa.HGridHeight, _ecfa.HGridWidth, _ecfa.HTemplate, false, _ecfa.HSkipEnabled, _adce, _abbc)
	_fgfg := _ggge - 1
	var _ggf error
	_gddg[_fgfg], _ggf = _eced.GetRegionBitmap()
	if _ggf != nil {
		return nil, _ggf
	}
	for _fgfg > 0 {
		_fgfg--
		_eced.Bitmap = nil
		_gddg[_fgfg], _ggf = _eced.GetRegionBitmap()
		if _ggf != nil {
			return nil, _ggf
		}
		if _ggf = _ecfa.combineGrayscalePlanes(_gddg, _fgfg); _ggf != nil {
			return nil, _ggf
		}
	}
	return _ecfa.computeGrayScalePlanes(_gddg, _ggge)
}
func (_gad *SymbolDictionary) encodeATFlags(_abeg _fg.BinaryWriter) (_ebbb int, _agdb error) {
	const _accb = "\u0065\u006e\u0063\u006f\u0064\u0065\u0041\u0054\u0046\u006c\u0061\u0067\u0073"
	if _gad.IsHuffmanEncoded || _gad.SdTemplate != 0 {
		return 0, nil
	}
	_fccg := 4
	if _gad.SdTemplate != 0 {
		_fccg = 1
	}
	for _ecbf := 0; _ecbf < _fccg; _ecbf++ {
		if _agdb = _abeg.WriteByte(byte(_gad.SdATX[_ecbf])); _agdb != nil {
			return _ebbb, _fd.Wrapf(_agdb, _accb, "\u0053d\u0041\u0054\u0058\u005b\u0025\u0064]", _ecbf)
		}
		_ebbb++
		if _agdb = _abeg.WriteByte(byte(_gad.SdATY[_ecbf])); _agdb != nil {
			return _ebbb, _fd.Wrapf(_agdb, _accb, "\u0053d\u0041\u0054\u0059\u005b\u0025\u0064]", _ecbf)
		}
		_ebbb++
	}
	return _ebbb, nil
}

type PageInformationSegment struct {
	_ace              _fg.StreamReader
	PageBMHeight      int
	PageBMWidth       int
	ResolutionX       int
	ResolutionY       int
	_eae              bool
	_aefgc            _d.CombinationOperator
	_bfgd             bool
	DefaultPixelValue uint8
	_bgeb             bool
	IsLossless        bool
	IsStripe          bool
	MaxStripeSize     uint16
}

func (_ddeb *PageInformationSegment) readResolution() error {
	_dbgg, _fbde := _ddeb._ace.ReadBits(32)
	if _fbde != nil {
		return _fbde
	}
	_ddeb.ResolutionX = int(_dbgg & _f.MaxInt32)
	_dbgg, _fbde = _ddeb._ace.ReadBits(32)
	if _fbde != nil {
		return _fbde
	}
	_ddeb.ResolutionY = int(_dbgg & _f.MaxInt32)
	return nil
}

type Pager interface {
	GetSegment(int) (*Header, error)
	GetBitmap() (*_d.Bitmap, error)
}

func (_cadc *RegionSegment) String() string {
	_aada := &_cc.Builder{}
	_aada.WriteString("\u0009[\u0052E\u0047\u0049\u004f\u004e\u0020S\u0045\u0047M\u0045\u004e\u0054\u005d\u000a")
	_aada.WriteString(_e.Sprintf("\t\u0009\u002d\u0020\u0042\u0069\u0074m\u0061\u0070\u0020\u0028\u0077\u0069d\u0074\u0068\u002c\u0020\u0068\u0065\u0069g\u0068\u0074\u0029\u0020\u005b\u0025\u0064\u0078\u0025\u0064]\u000a", _cadc.BitmapWidth, _cadc.BitmapHeight))
	_aada.WriteString(_e.Sprintf("\u0009\u0009\u002d\u0020L\u006f\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0028\u0078,\u0079)\u003a\u0020\u005b\u0025\u0064\u002c\u0025d\u005d\u000a", _cadc.XLocation, _cadc.YLocation))
	_aada.WriteString(_e.Sprintf("\t\u0009\u002d\u0020\u0043\u006f\u006db\u0069\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0073", _cadc.CombinaionOperator))
	return _aada.String()
}
func (_gg *GenericRefinementRegion) decodeTypicalPredictedLineTemplate0(_ac, _bcg, _eda, _cdge, _cca, _gb, _fgc, _bcd, _fdg int) error {
	var (
		_gfc, _ddb, _af, _bg, _fdd, _ggb int
		_dcgc                            byte
		_fgb                             error
	)
	if _ac > 0 {
		_dcgc, _fgb = _gg.RegionBitmap.GetByte(_fgc - _eda)
		if _fgb != nil {
			return _fgb
		}
		_af = int(_dcgc)
	}
	if _bcd > 0 && _bcd <= _gg.ReferenceBitmap.Height {
		_dcgc, _fgb = _gg.ReferenceBitmap.GetByte(_fdg - _cdge + _gb)
		if _fgb != nil {
			return _fgb
		}
		_bg = int(_dcgc) << 4
	}
	if _bcd >= 0 && _bcd < _gg.ReferenceBitmap.Height {
		_dcgc, _fgb = _gg.ReferenceBitmap.GetByte(_fdg + _gb)
		if _fgb != nil {
			return _fgb
		}
		_fdd = int(_dcgc) << 1
	}
	if _bcd > -2 && _bcd < _gg.ReferenceBitmap.Height-1 {
		_dcgc, _fgb = _gg.ReferenceBitmap.GetByte(_fdg + _cdge + _gb)
		if _fgb != nil {
			return _fgb
		}
		_ggb = int(_dcgc)
	}
	_gfc = ((_af >> 5) & 0x6) | ((_ggb >> 2) & 0x30) | (_fdd & 0x180) | (_bg & 0xc00)
	var _bgd int
	for _ea := 0; _ea < _cca; _ea = _bgd {
		var _ecfe int
		_bgd = _ea + 8
		var _dbbg int
		if _dbbg = _bcg - _ea; _dbbg > 8 {
			_dbbg = 8
		}
		_deb := _bgd < _bcg
		_dad := _bgd < _gg.ReferenceBitmap.Width
		_cgc := _gb + 1
		if _ac > 0 {
			_dcgc = 0
			if _deb {
				_dcgc, _fgb = _gg.RegionBitmap.GetByte(_fgc - _eda + 1)
				if _fgb != nil {
					return _fgb
				}
			}
			_af = (_af << 8) | int(_dcgc)
		}
		if _bcd > 0 && _bcd <= _gg.ReferenceBitmap.Height {
			var _bf int
			if _dad {
				_dcgc, _fgb = _gg.ReferenceBitmap.GetByte(_fdg - _cdge + _cgc)
				if _fgb != nil {
					return _fgb
				}
				_bf = int(_dcgc) << 4
			}
			_bg = (_bg << 8) | _bf
		}
		if _bcd >= 0 && _bcd < _gg.ReferenceBitmap.Height {
			var _eaf int
			if _dad {
				_dcgc, _fgb = _gg.ReferenceBitmap.GetByte(_fdg + _cgc)
				if _fgb != nil {
					return _fgb
				}
				_eaf = int(_dcgc) << 1
			}
			_fdd = (_fdd << 8) | _eaf
		}
		if _bcd > -2 && _bcd < (_gg.ReferenceBitmap.Height-1) {
			_dcgc = 0
			if _dad {
				_dcgc, _fgb = _gg.ReferenceBitmap.GetByte(_fdg + _cdge + _cgc)
				if _fgb != nil {
					return _fgb
				}
			}
			_ggb = (_ggb << 8) | int(_dcgc)
		}
		for _dbbgg := 0; _dbbgg < _dbbg; _dbbgg++ {
			var _ggg int
			_bce := false
			_def := (_gfc >> 4) & 0x1ff
			if _def == 0x1ff {
				_bce = true
				_ggg = 1
			} else if _def == 0x00 {
				_bce = true
			}
			if !_bce {
				if _gg._dag {
					_ddb = _gg.overrideAtTemplate0(_gfc, _ea+_dbbgg, _ac, _ecfe, _dbbgg)
					_gg._cdb.SetIndex(int32(_ddb))
				} else {
					_gg._cdb.SetIndex(int32(_gfc))
				}
				_ggg, _fgb = _gg._gf.DecodeBit(_gg._cdb)
				if _fgb != nil {
					return _fgb
				}
			}
			_gdc := uint(7 - _dbbgg)
			_ecfe |= _ggg << _gdc
			_gfc = ((_gfc & 0xdb6) << 1) | _ggg | (_af>>_gdc+5)&0x002 | ((_ggb>>_gdc + 2) & 0x010) | ((_fdd >> _gdc) & 0x080) | ((_bg >> _gdc) & 0x400)
		}
		_fgb = _gg.RegionBitmap.SetByte(_fgc, byte(_ecfe))
		if _fgb != nil {
			return _fgb
		}
		_fgc++
		_fdg++
	}
	return nil
}
func (_ddde *Header) Encode(w _fg.BinaryWriter) (_cdfa int, _dgec error) {
	const _bgbf = "\u0048\u0065\u0061d\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065"
	var _bgcc _fg.BinaryWriter
	_da.Log.Trace("\u005b\u0053\u0045G\u004d\u0045\u004e\u0054-\u0048\u0045\u0041\u0044\u0045\u0052\u005d[\u0045\u004e\u0043\u004f\u0044\u0045\u005d\u0020\u0042\u0065\u0067\u0069\u006e\u0073")
	defer func() {
		if _dgec != nil {
			_da.Log.Trace("[\u0053\u0045\u0047\u004d\u0045\u004eT\u002d\u0048\u0045\u0041\u0044\u0045R\u005d\u005b\u0045\u004e\u0043\u004f\u0044E\u005d\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u002e\u0020%\u0076", _dgec)
		} else {
			_da.Log.Trace("\u005b\u0053\u0045\u0047ME\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0025\u0076", _ddde)
			_da.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045N\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u005b\u0045\u004e\u0043O\u0044\u0045\u005d\u0020\u0046\u0069\u006ei\u0073\u0068\u0065\u0064")
		}
	}()
	w.FinishByte()
	if _ddde.SegmentData != nil {
		_ebfgc, _ccf := _ddde.SegmentData.(SegmentEncoder)
		if !_ccf {
			return 0, _fd.Errorf(_bgbf, "\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0054\u0020\u0064\u006f\u0065s\u006e\u0027\u0074\u0020\u0069\u006d\u0070\u006c\u0065\u006d\u0065\u006e\u0074 \u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0045\u006e\u0063\u006f\u0064er\u0020\u0069\u006e\u0074\u0065\u0072\u0066\u0061\u0063\u0065", _ddde.SegmentData)
		}
		_bgcc = _fg.BufferedMSB()
		_cdfa, _dgec = _ebfgc.Encode(_bgcc)
		if _dgec != nil {
			return 0, _fd.Wrap(_dgec, _bgbf, "")
		}
		_ddde.SegmentDataLength = uint64(_cdfa)
	}
	if _ddde.pageSize() == 4 {
		_ddde.PageAssociationFieldSize = true
	}
	var _aeba int
	_aeba, _dgec = _ddde.writeSegmentNumber(w)
	if _dgec != nil {
		return 0, _fd.Wrap(_dgec, _bgbf, "")
	}
	_cdfa += _aeba
	if _dgec = _ddde.writeFlags(w); _dgec != nil {
		return _cdfa, _fd.Wrap(_dgec, _bgbf, "")
	}
	_cdfa++
	_aeba, _dgec = _ddde.writeReferredToCount(w)
	if _dgec != nil {
		return 0, _fd.Wrap(_dgec, _bgbf, "")
	}
	_cdfa += _aeba
	_aeba, _dgec = _ddde.writeReferredToSegments(w)
	if _dgec != nil {
		return 0, _fd.Wrap(_dgec, _bgbf, "")
	}
	_cdfa += _aeba
	_aeba, _dgec = _ddde.writeSegmentPageAssociation(w)
	if _dgec != nil {
		return 0, _fd.Wrap(_dgec, _bgbf, "")
	}
	_cdfa += _aeba
	_aeba, _dgec = _ddde.writeSegmentDataLength(w)
	if _dgec != nil {
		return 0, _fd.Wrap(_dgec, _bgbf, "")
	}
	_cdfa += _aeba
	_ddde.HeaderLength = int64(_cdfa) - int64(_ddde.SegmentDataLength)
	if _bgcc != nil {
		if _, _dgec = w.Write(_bgcc.Data()); _dgec != nil {
			return _cdfa, _fd.Wrap(_dgec, _bgbf, "\u0077r\u0069t\u0065\u0020\u0073\u0065\u0067m\u0065\u006et\u0020\u0064\u0061\u0074\u0061")
		}
	}
	return _cdfa, nil
}
func (_ceg *RegionSegment) readCombinationOperator() error {
	_defd, _ddae := _ceg._dbge.ReadBits(3)
	if _ddae != nil {
		return _ddae
	}
	_ceg.CombinaionOperator = _d.CombinationOperator(_defd & 0xF)
	return nil
}
func (_degc *TextRegion) encodeFlags(_faae _fg.BinaryWriter) (_caded int, _aggb error) {
	const _fddb = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _aggb = _faae.WriteBit(int(_degc.SbrTemplate)); _aggb != nil {
		return _caded, _fd.Wrap(_aggb, _fddb, "s\u0062\u0072\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	if _, _aggb = _faae.WriteBits(uint64(_degc.SbDsOffset), 5); _aggb != nil {
		return _caded, _fd.Wrap(_aggb, _fddb, "\u0073\u0062\u0044\u0073\u004f\u0066\u0066\u0073\u0065\u0074")
	}
	if _aggb = _faae.WriteBit(int(_degc.DefaultPixel)); _aggb != nil {
		return _caded, _fd.Wrap(_aggb, _fddb, "\u0044\u0065\u0066a\u0075\u006c\u0074\u0050\u0069\u0078\u0065\u006c")
	}
	if _, _aggb = _faae.WriteBits(uint64(_degc.CombinationOperator), 2); _aggb != nil {
		return _caded, _fd.Wrap(_aggb, _fddb, "\u0043\u006f\u006d\u0062in\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	if _aggb = _faae.WriteBit(int(_degc.IsTransposed)); _aggb != nil {
		return _caded, _fd.Wrap(_aggb, _fddb, "\u0069\u0073\u0020\u0074\u0072\u0061\u006e\u0073\u0070\u006f\u0073\u0065\u0064")
	}
	if _, _aggb = _faae.WriteBits(uint64(_degc.ReferenceCorner), 2); _aggb != nil {
		return _caded, _fd.Wrap(_aggb, _fddb, "\u0072\u0065f\u0065\u0072\u0065n\u0063\u0065\u0020\u0063\u006f\u0072\u006e\u0065\u0072")
	}
	if _, _aggb = _faae.WriteBits(uint64(_degc.LogSBStrips), 2); _aggb != nil {
		return _caded, _fd.Wrap(_aggb, _fddb, "L\u006f\u0067\u0053\u0042\u0053\u0074\u0072\u0069\u0070\u0073")
	}
	var _daddb int
	if _degc.UseRefinement {
		_daddb = 1
	}
	if _aggb = _faae.WriteBit(_daddb); _aggb != nil {
		return _caded, _fd.Wrap(_aggb, _fddb, "\u0075\u0073\u0065\u0020\u0072\u0065\u0066\u0069\u006ee\u006d\u0065\u006e\u0074")
	}
	_daddb = 0
	if _degc.IsHuffmanEncoded {
		_daddb = 1
	}
	if _aggb = _faae.WriteBit(_daddb); _aggb != nil {
		return _caded, _fd.Wrap(_aggb, _fddb, "u\u0073\u0065\u0020\u0068\u0075\u0066\u0066\u006d\u0061\u006e")
	}
	_caded = 2
	return _caded, nil
}
func (_cfec *SymbolDictionary) decodeThroughTextRegion(_gggae, _aebg, _gdfaf uint32) error {
	if _cfec._begf == nil {
		_cfec._begf = _daae(_cfec._bcagf, nil)
		_cfec._begf.setContexts(_cfec._bagf, _b.NewStats(512, 1), _b.NewStats(512, 1), _b.NewStats(512, 1), _b.NewStats(512, 1), _cfec._fac, _b.NewStats(512, 1), _b.NewStats(512, 1), _b.NewStats(512, 1), _b.NewStats(512, 1))
	}
	if _dfbg := _cfec.setSymbolsArray(); _dfbg != nil {
		return _dfbg
	}
	_cfec._begf.setParameters(_cfec._fffd, _cfec.IsHuffmanEncoded, true, _gggae, _aebg, _gdfaf, 1, _cfec._gbbg+_cfec._ffa, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, _cfec.SdrTemplate, _cfec.SdrATX, _cfec.SdrATY, _cfec._bded, _cfec._cff)
	return _cfec.addSymbol(_cfec._begf)
}
func (_dccd *Header) readHeaderLength(_dba _fg.StreamReader, _acc int64) {
	_dccd.HeaderLength = _dba.StreamPosition() - _acc
}
func (_effd *GenericRegion) updateOverrideFlags() error {
	const _eddf = "\u0075\u0070\u0064\u0061te\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _effd.GBAtX == nil || _effd.GBAtY == nil {
		return nil
	}
	if len(_effd.GBAtX) != len(_effd.GBAtY) {
		return _fd.Errorf(_eddf, "i\u006eco\u0073i\u0073t\u0065\u006e\u0074\u0020\u0041T\u0020\u0070\u0069x\u0065\u006c\u002e\u0020\u0041m\u006f\u0075\u006et\u0020\u006f\u0066\u0020\u0027\u0078\u0027\u0020\u0070\u0069\u0078e\u006c\u0073\u003a %d\u002c\u0020\u0041\u006d\u006f\u0075n\u0074\u0020\u006f\u0066\u0020\u0027\u0079\u0027\u0020\u0070\u0069\u0078e\u006cs\u003a\u0020\u0025\u0064", len(_effd.GBAtX), len(_effd.GBAtY))
	}
	_effd.GBAtOverride = make([]bool, len(_effd.GBAtX))
	switch _effd.GBTemplate {
	case 0:
		if !_effd.UseExtTemplates {
			if _effd.GBAtX[0] != 3 || _effd.GBAtY[0] != -1 {
				_effd.setOverrideFlag(0)
			}
			if _effd.GBAtX[1] != -3 || _effd.GBAtY[1] != -1 {
				_effd.setOverrideFlag(1)
			}
			if _effd.GBAtX[2] != 2 || _effd.GBAtY[2] != -2 {
				_effd.setOverrideFlag(2)
			}
			if _effd.GBAtX[3] != -2 || _effd.GBAtY[3] != -2 {
				_effd.setOverrideFlag(3)
			}
		} else {
			if _effd.GBAtX[0] != -2 || _effd.GBAtY[0] != 0 {
				_effd.setOverrideFlag(0)
			}
			if _effd.GBAtX[1] != 0 || _effd.GBAtY[1] != -2 {
				_effd.setOverrideFlag(1)
			}
			if _effd.GBAtX[2] != -2 || _effd.GBAtY[2] != -1 {
				_effd.setOverrideFlag(2)
			}
			if _effd.GBAtX[3] != -1 || _effd.GBAtY[3] != -2 {
				_effd.setOverrideFlag(3)
			}
			if _effd.GBAtX[4] != 1 || _effd.GBAtY[4] != -2 {
				_effd.setOverrideFlag(4)
			}
			if _effd.GBAtX[5] != 2 || _effd.GBAtY[5] != -1 {
				_effd.setOverrideFlag(5)
			}
			if _effd.GBAtX[6] != -3 || _effd.GBAtY[6] != 0 {
				_effd.setOverrideFlag(6)
			}
			if _effd.GBAtX[7] != -4 || _effd.GBAtY[7] != 0 {
				_effd.setOverrideFlag(7)
			}
			if _effd.GBAtX[8] != 2 || _effd.GBAtY[8] != -2 {
				_effd.setOverrideFlag(8)
			}
			if _effd.GBAtX[9] != 3 || _effd.GBAtY[9] != -1 {
				_effd.setOverrideFlag(9)
			}
			if _effd.GBAtX[10] != -2 || _effd.GBAtY[10] != -2 {
				_effd.setOverrideFlag(10)
			}
			if _effd.GBAtX[11] != -3 || _effd.GBAtY[11] != -1 {
				_effd.setOverrideFlag(11)
			}
		}
	case 1:
		if _effd.GBAtX[0] != 3 || _effd.GBAtY[0] != -1 {
			_effd.setOverrideFlag(0)
		}
	case 2:
		if _effd.GBAtX[0] != 2 || _effd.GBAtY[0] != -1 {
			_effd.setOverrideFlag(0)
		}
	case 3:
		if _effd.GBAtX[0] != 2 || _effd.GBAtY[0] != -1 {
			_effd.setOverrideFlag(0)
		}
	}
	return nil
}
func (_bbea *PatternDictionary) GetDictionary() ([]*_d.Bitmap, error) {
	if _bbea.Patterns != nil {
		return _bbea.Patterns, nil
	}
	if !_bbea.IsMMREncoded {
		_bbea.setGbAtPixels()
	}
	_cgba := NewGenericRegion(_bbea._fbdd)
	_cgba.setParametersMMR(_bbea.IsMMREncoded, _bbea.DataOffset, _bbea.DataLength, uint32(_bbea.HdpHeight), (_bbea.GrayMax+1)*uint32(_bbea.HdpWidth), _bbea.HDTemplate, false, false, _bbea.GBAtX, _bbea.GBAtY)
	_gdge, _cfdf := _cgba.GetRegionBitmap()
	if _cfdf != nil {
		return nil, _cfdf
	}
	if _cfdf = _bbea.extractPatterns(_gdge); _cfdf != nil {
		return nil, _cfdf
	}
	return _bbea.Patterns, nil
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

func (_faff *HalftoneRegion) renderPattern(_dcda [][]int) (_ageg error) {
	var _eege, _gde int
	for _fdee := 0; _fdee < int(_faff.HGridHeight); _fdee++ {
		for _fge := 0; _fge < int(_faff.HGridWidth); _fge++ {
			_eege = _faff.computeX(_fdee, _fge)
			_gde = _faff.computeY(_fdee, _fge)
			_dbc := _faff.Patterns[_dcda[_fdee][_fge]]
			if _ageg = _d.Blit(_dbc, _faff.HalftoneRegionBitmap, _eege+int(_faff.HGridX), _gde+int(_faff.HGridY), _faff.CombinationOperator); _ageg != nil {
				return _ageg
			}
		}
	}
	return nil
}
func (_fadb *TextRegion) decodeRdy() (int64, error) {
	const _fbg = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064y"
	if _fadb.IsHuffmanEncoded {
		if _fadb.SbHuffRDY == 3 {
			if _fadb._bggb == nil {
				var (
					_cfaad int
					_dbged error
				)
				if _fadb.SbHuffFS == 3 {
					_cfaad++
				}
				if _fadb.SbHuffDS == 3 {
					_cfaad++
				}
				if _fadb.SbHuffDT == 3 {
					_cfaad++
				}
				if _fadb.SbHuffRDWidth == 3 {
					_cfaad++
				}
				if _fadb.SbHuffRDHeight == 3 {
					_cfaad++
				}
				if _fadb.SbHuffRDX == 3 {
					_cfaad++
				}
				_fadb._bggb, _dbged = _fadb.getUserTable(_cfaad)
				if _dbged != nil {
					return 0, _fd.Wrap(_dbged, _fbg, "")
				}
			}
			return _fadb._bggb.Decode(_fadb._effb)
		}
		_eaac, _bgacg := _ba.GetStandardTable(14 + int(_fadb.SbHuffRDY))
		if _bgacg != nil {
			return 0, _bgacg
		}
		return _eaac.Decode(_fadb._effb)
	}
	_dcef, _fdbd := _fadb._dffe.DecodeInt(_fadb._bfbb)
	if _fdbd != nil {
		return 0, _fd.Wrap(_fdbd, _fbg, "")
	}
	return int64(_dcef), nil
}

var _ templater = &template1{}

func NewRegionSegment(r _fg.StreamReader) *RegionSegment { return &RegionSegment{_dbge: r} }

type GenericRefinementRegion struct {
	_dae            templater
	_dec            templater
	_geg            _fg.StreamReader
	_cdg            *Header
	RegionInfo      *RegionSegment
	IsTPGROn        bool
	TemplateID      int8
	Template        templater
	GrAtX           []int8
	GrAtY           []int8
	RegionBitmap    *_d.Bitmap
	ReferenceBitmap *_d.Bitmap
	ReferenceDX     int32
	ReferenceDY     int32
	_gf             *_b.Decoder
	_cdb            *_b.DecoderStats
	_dag            bool
	_db             []bool
}

func (_dgbc *SymbolDictionary) getSbSymCodeLen() int8 {
	_fcbc := int8(_f.Ceil(_f.Log(float64(_dgbc._gbbg+_dgbc.NumberOfNewSymbols)) / _f.Log(2)))
	if _dgbc.IsHuffmanEncoded && _fcbc < 1 {
		return 1
	}
	return _fcbc
}
func (_fgfe *SymbolDictionary) setSymbolsArray() error {
	if _fgfe._cgfbb == nil {
		if _dcad := _fgfe.retrieveImportSymbols(); _dcad != nil {
			return _dcad
		}
	}
	if _fgfe._bded == nil {
		_fgfe._bded = append(_fgfe._bded, _fgfe._cgfbb...)
	}
	return nil
}
func (_ccaa *SymbolDictionary) setRetainedCodingContexts(_bfdac *SymbolDictionary) {
	_ccaa._fffd = _bfdac._fffd
	_ccaa.IsHuffmanEncoded = _bfdac.IsHuffmanEncoded
	_ccaa.UseRefinementAggregation = _bfdac.UseRefinementAggregation
	_ccaa.SdTemplate = _bfdac.SdTemplate
	_ccaa.SdrTemplate = _bfdac.SdrTemplate
	_ccaa.SdATX = _bfdac.SdATX
	_ccaa.SdATY = _bfdac.SdATY
	_ccaa.SdrATX = _bfdac.SdrATX
	_ccaa.SdrATY = _bfdac.SdrATY
	_ccaa._bagf = _bfdac._bagf
}
func (_cb *GenericRefinementRegion) overrideAtTemplate0(_fegc, _acf, _bgfc, _aeb, _fagg int) int {
	if _cb._db[0] {
		_fegc &= 0xfff7
		if _cb.GrAtY[0] == 0 && int(_cb.GrAtX[0]) >= -_fagg {
			_fegc |= (_aeb >> uint(7-(_fagg+int(_cb.GrAtX[0]))) & 0x1) << 3
		} else {
			_fegc |= _cb.getPixel(_cb.RegionBitmap, _acf+int(_cb.GrAtX[0]), _bgfc+int(_cb.GrAtY[0])) << 3
		}
	}
	if _cb._db[1] {
		_fegc &= 0xefff
		if _cb.GrAtY[1] == 0 && int(_cb.GrAtX[1]) >= -_fagg {
			_fegc |= (_aeb >> uint(7-(_fagg+int(_cb.GrAtX[1]))) & 0x1) << 12
		} else {
			_fegc |= _cb.getPixel(_cb.ReferenceBitmap, _acf+int(_cb.GrAtX[1]), _bgfc+int(_cb.GrAtY[1]))
		}
	}
	return _fegc
}
func (_ggbf *PageInformationSegment) readWidthAndHeight() error {
	_gdfag, _cagg := _ggbf._ace.ReadBits(32)
	if _cagg != nil {
		return _cagg
	}
	_ggbf.PageBMWidth = int(_gdfag & _f.MaxInt32)
	_gdfag, _cagg = _ggbf._ace.ReadBits(32)
	if _cagg != nil {
		return _cagg
	}
	_ggbf.PageBMHeight = int(_gdfag & _f.MaxInt32)
	return nil
}
func (_gce *GenericRefinementRegion) updateOverride() error {
	if _gce.GrAtX == nil || _gce.GrAtY == nil {
		return _ge.New("\u0041\u0054\u0020\u0070\u0069\u0078\u0065\u006c\u0073\u0020\u006e\u006ft\u0020\u0073\u0065\u0074")
	}
	if len(_gce.GrAtX) != len(_gce.GrAtY) {
		return _ge.New("A\u0054\u0020\u0070\u0069xe\u006c \u0069\u006e\u0063\u006f\u006es\u0069\u0073\u0074\u0065\u006e\u0074")
	}
	_gce._db = make([]bool, len(_gce.GrAtX))
	switch _gce.TemplateID {
	case 0:
		if _gce.GrAtX[0] != -1 && _gce.GrAtY[0] != -1 {
			_gce._db[0] = true
			_gce._dag = true
		}
		if _gce.GrAtX[1] != -1 && _gce.GrAtY[1] != -1 {
			_gce._db[1] = true
			_gce._dag = true
		}
	case 1:
		_gce._dag = false
	}
	return nil
}
func (_ccc *GenericRefinementRegion) getPixel(_daed *_d.Bitmap, _age, _acg int) int {
	if _age < 0 || _age >= _daed.Width {
		return 0
	}
	if _acg < 0 || _acg >= _daed.Height {
		return 0
	}
	if _daed.GetPixel(_age, _acg) {
		return 1
	}
	return 0
}
func (_agfc *SymbolDictionary) decodeNewSymbols(_fbecc, _bacg uint32, _eec *_d.Bitmap, _gfce, _cfbb int32) error {
	if _agfc._debeg == nil {
		_agfc._debeg = _cda(_agfc._bcagf, nil)
		if _agfc._fffd == nil {
			var _badc error
			_agfc._fffd, _badc = _b.New(_agfc._bcagf)
			if _badc != nil {
				return _badc
			}
		}
		if _agfc._bagf == nil {
			_agfc._bagf = _b.NewStats(65536, 1)
		}
	}
	_agfc._debeg.setParameters(_agfc._bagf, _agfc._fffd, _agfc.SdrTemplate, _fbecc, _bacg, _eec, _gfce, _cfbb, false, _agfc.SdrATX, _agfc.SdrATY)
	return _agfc.addSymbol(_agfc._debeg)
}
func (_ebd *HalftoneRegion) computeSegmentDataStructure() error {
	_ebd.DataOffset = _ebd._gdfe.StreamPosition()
	_ebd.DataHeaderLength = _ebd.DataOffset - _ebd.DataHeaderOffset
	_ebd.DataLength = int64(_ebd._gdfe.Length()) - _ebd.DataHeaderLength
	return nil
}

type TextRegion struct {
	_effb                   _fg.StreamReader
	RegionInfo              *RegionSegment
	SbrTemplate             int8
	SbDsOffset              int8
	DefaultPixel            int8
	CombinationOperator     _d.CombinationOperator
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
	_fbdbc                  int64
	SbStrips                int8
	NumberOfSymbols         uint32
	RegionBitmap            *_d.Bitmap
	Symbols                 []*_d.Bitmap
	_dffe                   *_b.Decoder
	_edge                   *GenericRefinementRegion
	_bgad                   *_b.DecoderStats
	_geaa                   *_b.DecoderStats
	_dafb                   *_b.DecoderStats
	_dace                   *_b.DecoderStats
	_gbfge                  *_b.DecoderStats
	_egcff                  *_b.DecoderStats
	_acgg                   *_b.DecoderStats
	_beabb                  *_b.DecoderStats
	_beda                   *_b.DecoderStats
	_bfbb                   *_b.DecoderStats
	_aeed                   *_b.DecoderStats
	_ebcb                   int8
	_gegc                   *_ba.FixedSizeTable
	Header                  *Header
	_bcf                    _ba.Tabler
	_cadeg                  _ba.Tabler
	_gfad                   _ba.Tabler
	_gcba                   _ba.Tabler
	_fbdda                  _ba.Tabler
	_agag                   _ba.Tabler
	_bggb                   _ba.Tabler
	_fbeac                  _ba.Tabler
	_fegdc, _cfaa           map[int]int
	_ddfbb                  []int
	_fbc                    *_d.Points
	_aeee                   *_d.Bitmaps
	_faed                   *_ca.IntSlice
	_eagb, _aagc            int
	_ffbc                   *_d.Boxes
}

func (_ecfeb *SymbolDictionary) decodeHeightClassCollectiveBitmap(_egfg int64, _cefg, _geff uint32) (*_d.Bitmap, error) {
	if _egfg == 0 {
		_fgbc := _d.New(int(_geff), int(_cefg))
		var (
			_cabc byte
			_dcga error
		)
		for _ggfe := 0; _ggfe < len(_fgbc.Data); _ggfe++ {
			_cabc, _dcga = _ecfeb._bcagf.ReadByte()
			if _dcga != nil {
				return nil, _dcga
			}
			if _dcga = _fgbc.SetByte(_ggfe, _cabc); _dcga != nil {
				return nil, _dcga
			}
		}
		return _fgbc, nil
	}
	if _ecfeb._eaeb == nil {
		_ecfeb._eaeb = NewGenericRegion(_ecfeb._bcagf)
	}
	_ecfeb._eaeb.setParameters(true, _ecfeb._bcagf.StreamPosition(), _egfg, _cefg, _geff)
	_dcgca, _gagf := _ecfeb._eaeb.GetRegionBitmap()
	if _gagf != nil {
		return nil, _gagf
	}
	return _dcgca, nil
}
func (_efbd *TextRegion) encodeSymbols(_fegdd _fg.BinaryWriter) (_gbbe int, _ggfd error) {
	const _bdfga = "\u0065\u006e\u0063\u006f\u0064\u0065\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	_eggf := make([]byte, 4)
	_fe.BigEndian.PutUint32(_eggf, _efbd.NumberOfSymbols)
	if _gbbe, _ggfd = _fegdd.Write(_eggf); _ggfd != nil {
		return _gbbe, _fd.Wrap(_ggfd, _bdfga, "\u004e\u0075\u006dbe\u0072\u004f\u0066\u0053\u0079\u006d\u0062\u006f\u006c\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073")
	}
	_geed, _ggfd := _d.NewClassedPoints(_efbd._fbc, _efbd._ddfbb)
	if _ggfd != nil {
		return 0, _fd.Wrap(_ggfd, _bdfga, "")
	}
	var _bbg, _eeagc int
	_dbfb := _ga.New()
	_dbfb.Init()
	if _ggfd = _dbfb.EncodeInteger(_ga.IADT, 0); _ggfd != nil {
		return _gbbe, _fd.Wrap(_ggfd, _bdfga, "\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0020\u0044\u0054")
	}
	_ffdc, _ggfd := _geed.GroupByY()
	if _ggfd != nil {
		return 0, _fd.Wrap(_ggfd, _bdfga, "")
	}
	for _, _bdge := range _ffdc {
		_debg := int(_bdge.YAtIndex(0))
		_bbca := _debg - _bbg
		if _ggfd = _dbfb.EncodeInteger(_ga.IADT, _bbca); _ggfd != nil {
			return _gbbe, _fd.Wrap(_ggfd, _bdfga, "")
		}
		var _cfafb int
		for _fdef, _ddead := range _bdge.IntSlice {
			switch _fdef {
			case 0:
				_afgg := int(_bdge.XAtIndex(_fdef)) - _eeagc
				if _ggfd = _dbfb.EncodeInteger(_ga.IAFS, _afgg); _ggfd != nil {
					return _gbbe, _fd.Wrap(_ggfd, _bdfga, "")
				}
				_eeagc += _afgg
				_cfafb = _eeagc
			default:
				_ceeg := int(_bdge.XAtIndex(_fdef)) - _cfafb
				if _ggfd = _dbfb.EncodeInteger(_ga.IADS, _ceeg); _ggfd != nil {
					return _gbbe, _fd.Wrap(_ggfd, _bdfga, "")
				}
				_cfafb += _ceeg
			}
			_eefc, _acga := _efbd._faed.Get(_ddead)
			if _acga != nil {
				return _gbbe, _fd.Wrap(_acga, _bdfga, "")
			}
			_acafe, _fefg := _efbd._fegdc[_eefc]
			if !_fefg {
				_acafe, _fefg = _efbd._cfaa[_eefc]
				if !_fefg {
					return _gbbe, _fd.Errorf(_bdfga, "\u0053\u0079\u006d\u0062\u006f\u006c:\u0020\u0027\u0025d\u0027\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064 \u0069\u006e\u0020\u0067\u006cob\u0061\u006c\u0020\u0061\u006e\u0064\u0020\u006c\u006f\u0063\u0061\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0020\u006d\u0061\u0070", _eefc)
				}
			}
			if _acga = _dbfb.EncodeIAID(_efbd._aagc, _acafe); _acga != nil {
				return _gbbe, _fd.Wrap(_acga, _bdfga, "")
			}
		}
		if _ggfd = _dbfb.EncodeOOB(_ga.IADS); _ggfd != nil {
			return _gbbe, _fd.Wrap(_ggfd, _bdfga, "")
		}
	}
	_dbfb.Final()
	_acdd, _ggfd := _dbfb.WriteTo(_fegdd)
	if _ggfd != nil {
		return _gbbe, _fd.Wrap(_ggfd, _bdfga, "")
	}
	_gbbe += int(_acdd)
	return _gbbe, nil
}
func (_gggb *SymbolDictionary) retrieveImportSymbols() error {
	for _, _ggbe := range _gggb.Header.RTSegments {
		if _ggbe.Type == 0 {
			_cfaf, _gcdfa := _ggbe.GetSegmentData()
			if _gcdfa != nil {
				return _gcdfa
			}
			_dcaa, _bdgd := _cfaf.(*SymbolDictionary)
			if !_bdgd {
				return _e.Errorf("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0044\u0061\u0074a\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0053\u0079\u006d\u0062\u006f\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0053\u0065\u0067m\u0065\u006e\u0074\u003a\u0020%\u0054", _cfaf)
			}
			_fdgd, _gcdfa := _dcaa.GetDictionary()
			if _gcdfa != nil {
				return _e.Errorf("\u0072\u0065\u006c\u0061\u0074\u0065\u0064 \u0073\u0065\u0067m\u0065\u006e\u0074 \u0077\u0069t\u0068\u0020\u0069\u006e\u0064\u0065x\u003a %\u0064\u0020\u0067\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0073", _ggbe.SegmentNumber, _gcdfa.Error())
			}
			_gggb._cgfbb = append(_gggb._cgfbb, _fdgd...)
			_gggb._gbbg += _dcaa.NumberOfExportedSymbols
		}
	}
	return nil
}
func (_caag *PageInformationSegment) String() string {
	_gbc := &_cc.Builder{}
	_gbc.WriteString("\u000a\u005b\u0050\u0041G\u0045\u002d\u0049\u004e\u0046\u004f\u0052\u004d\u0041\u0054I\u004fN\u002d\u0053\u0045\u0047\u004d\u0045\u004eT\u005d\u000a")
	_gbc.WriteString(_e.Sprintf("\u0009\u002d \u0042\u004d\u0048e\u0069\u0067\u0068\u0074\u003a\u0020\u0025\u0064\u000a", _caag.PageBMHeight))
	_gbc.WriteString(_e.Sprintf("\u0009-\u0020B\u004d\u0057\u0069\u0064\u0074\u0068\u003a\u0020\u0025\u0064\u000a", _caag.PageBMWidth))
	_gbc.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0052es\u006f\u006c\u0075\u0074\u0069\u006f\u006e\u0058\u003a\u0020\u0025\u0064\u000a", _caag.ResolutionX))
	_gbc.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0052es\u006f\u006c\u0075\u0074\u0069\u006f\u006e\u0059\u003a\u0020\u0025\u0064\u000a", _caag.ResolutionY))
	_gbc.WriteString(_e.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006e\u004f\u0070\u0065r\u0061\u0074\u006fr\u003a \u0025\u0073\u000a", _caag._aefgc))
	_gbc.WriteString(_e.Sprintf("\t\u002d\u0020\u0043\u006f\u006d\u0062i\u006e\u0061\u0074\u0069\u006f\u006eO\u0070\u0065\u0072\u0061\u0074\u006f\u0072O\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u003a\u0020\u0025v\u000a", _caag._eae))
	_gbc.WriteString(_e.Sprintf("\u0009-\u0020I\u0073\u004c\u006f\u0073\u0073l\u0065\u0073s\u003a\u0020\u0025\u0076\u000a", _caag.IsLossless))
	_gbc.WriteString(_e.Sprintf("\u0009\u002d\u0020R\u0065\u0071\u0075\u0069r\u0065\u0073\u0041\u0075\u0078\u0069\u006ci\u0061\u0072\u0079\u0042\u0075\u0066\u0066\u0065\u0072\u003a\u0020\u0025\u0076\u000a", _caag._bfgd))
	_gbc.WriteString(_e.Sprintf("\u0009\u002d\u0020M\u0069\u0067\u0068\u0074C\u006f\u006e\u0074\u0061\u0069\u006e\u0052e\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0076\u000a", _caag._bgeb))
	_gbc.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0049\u0073\u0053\u0074\u0072\u0069\u0070\u0065\u0064:\u0020\u0025\u0076\u000a", _caag.IsStripe))
	_gbc.WriteString(_e.Sprintf("\t\u002d\u0020\u004d\u0061xS\u0074r\u0069\u0070\u0065\u0053\u0069z\u0065\u003a\u0020\u0025\u0076\u000a", _caag.MaxStripeSize))
	return _gbc.String()
}

type EncodeInitializer interface{ InitEncode() }

func (_fee *GenericRegion) Init(h *Header, r _fg.StreamReader) error {
	_fee.RegionSegment = NewRegionSegment(r)
	_fee._baf = r
	return _fee.parseHeader()
}
func (_bbecf *TextRegion) decodeDT() (_ggdgf int64, _egbc error) {
	if _bbecf.IsHuffmanEncoded {
		if _bbecf.SbHuffDT == 3 {
			_ggdgf, _egbc = _bbecf._gfad.Decode(_bbecf._effb)
			if _egbc != nil {
				return 0, _egbc
			}
		} else {
			var _bgge _ba.Tabler
			_bgge, _egbc = _ba.GetStandardTable(11 + int(_bbecf.SbHuffDT))
			if _egbc != nil {
				return 0, _egbc
			}
			_ggdgf, _egbc = _bgge.Decode(_bbecf._effb)
			if _egbc != nil {
				return 0, _egbc
			}
		}
	} else {
		var _dgfgd int32
		_dgfgd, _egbc = _bbecf._dffe.DecodeInt(_bbecf._bgad)
		if _egbc != nil {
			return
		}
		_ggdgf = int64(_dgfgd)
	}
	_ggdgf *= int64(_bbecf.SbStrips)
	return _ggdgf, nil
}
func (_ef *GenericRefinementRegion) decodeOptimized(_cf, _gea, _egc, _cg, _gaf, _ade, _gcc int) error {
	var (
		_efb  error
		_egcf int
		_fa   int
	)
	_ed := _cf - int(_ef.ReferenceDY)
	if _bag := int(-_ef.ReferenceDX); _bag > 0 {
		_egcf = _bag
	}
	_fab := _ef.ReferenceBitmap.GetByteIndex(_egcf, _ed)
	if _ef.ReferenceDX > 0 {
		_fa = int(_ef.ReferenceDX)
	}
	_gcg := _ef.RegionBitmap.GetByteIndex(_fa, _cf)
	switch _ef.TemplateID {
	case 0:
		_efb = _ef.decodeTemplate(_cf, _gea, _egc, _cg, _gaf, _ade, _gcc, _gcg, _ed, _fab, _ef._dae)
	case 1:
		_efb = _ef.decodeTemplate(_cf, _gea, _egc, _cg, _gaf, _ade, _gcc, _gcg, _ed, _fab, _ef._dec)
	}
	return _efb
}
func (_gccb *PageInformationSegment) readCombinationOperator() error {
	_baga, _bgbg := _gccb._ace.ReadBits(2)
	if _bgbg != nil {
		return _bgbg
	}
	_gccb._aefgc = _d.CombinationOperator(int(_baga))
	return nil
}
func (_eba *template0) setIndex(_bcec *_b.DecoderStats) { _bcec.SetIndex(0x100) }
func (_ccgd *TextRegion) decodeRdh() (int64, error) {
	const _fada = "\u0064e\u0063\u006f\u0064\u0065\u0052\u0064h"
	if _ccgd.IsHuffmanEncoded {
		if _ccgd.SbHuffRDHeight == 3 {
			if _ccgd._fbdda == nil {
				var (
					_fgae int
					_bfdf error
				)
				if _ccgd.SbHuffFS == 3 {
					_fgae++
				}
				if _ccgd.SbHuffDS == 3 {
					_fgae++
				}
				if _ccgd.SbHuffDT == 3 {
					_fgae++
				}
				if _ccgd.SbHuffRDWidth == 3 {
					_fgae++
				}
				_ccgd._fbdda, _bfdf = _ccgd.getUserTable(_fgae)
				if _bfdf != nil {
					return 0, _fd.Wrap(_bfdf, _fada, "")
				}
			}
			return _ccgd._fbdda.Decode(_ccgd._effb)
		}
		_bbdef, _dcgcd := _ba.GetStandardTable(14 + int(_ccgd.SbHuffRDHeight))
		if _dcgcd != nil {
			return 0, _fd.Wrap(_dcgcd, _fada, "")
		}
		return _bbdef.Decode(_ccgd._effb)
	}
	_agef, _ffag := _ccgd._dffe.DecodeInt(_ccgd._acgg)
	if _ffag != nil {
		return 0, _fd.Wrap(_ffag, _fada, "")
	}
	return int64(_agef), nil
}
func (_gfdg *TextRegion) readAmountOfSymbolInstances() error {
	_dcca, _acba := _gfdg._effb.ReadBits(32)
	if _acba != nil {
		return _acba
	}
	_gfdg.NumberOfSymbolInstances = uint32(_dcca & _f.MaxUint32)
	_fefa := _gfdg.RegionInfo.BitmapWidth * _gfdg.RegionInfo.BitmapHeight
	if _fefa < _gfdg.NumberOfSymbolInstances {
		_da.Log.Debug("\u004c\u0069\u006d\u0069t\u0069\u006e\u0067\u0020t\u0068\u0065\u0020n\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020d\u0065\u0063\u006f\u0064e\u0064\u0020\u0073\u0079m\u0062\u006f\u006c\u0020\u0069n\u0073\u0074\u0061\u006e\u0063\u0065\u0073 \u0074\u006f\u0020\u006f\u006ee\u0020\u0070\u0065\u0072\u0020\u0070\u0069\u0078\u0065l\u0020\u0028\u0020\u0025\u0064\u0020\u0069\u006e\u0073\u0074\u0065\u0061\u0064\u0020\u006f\u0066\u0020\u0025\u0064\u0029", _fefa, _gfdg.NumberOfSymbolInstances)
		_gfdg.NumberOfSymbolInstances = _fefa
	}
	return nil
}
func (_eaa *GenericRefinementRegion) decodeTypicalPredictedLineTemplate1(_bca, _ddc, _df, _ab, _dgd, _gaff, _ae, _eag, _cfg int) (_dbf error) {
	var (
		_fdgf, _dac int
		_caf, _dagb int
		_faf, _ead  int
		_aff        byte
	)
	if _bca > 0 {
		_aff, _dbf = _eaa.RegionBitmap.GetByte(_ae - _df)
		if _dbf != nil {
			return
		}
		_caf = int(_aff)
	}
	if _eag > 0 && _eag <= _eaa.ReferenceBitmap.Height {
		_aff, _dbf = _eaa.ReferenceBitmap.GetByte(_cfg - _ab + _gaff)
		if _dbf != nil {
			return
		}
		_dagb = int(_aff) << 2
	}
	if _eag >= 0 && _eag < _eaa.ReferenceBitmap.Height {
		_aff, _dbf = _eaa.ReferenceBitmap.GetByte(_cfg + _gaff)
		if _dbf != nil {
			return
		}
		_faf = int(_aff)
	}
	if _eag > -2 && _eag < _eaa.ReferenceBitmap.Height-1 {
		_aff, _dbf = _eaa.ReferenceBitmap.GetByte(_cfg + _ab + _gaff)
		if _dbf != nil {
			return
		}
		_ead = int(_aff)
	}
	_fdgf = ((_caf >> 5) & 0x6) | ((_ead >> 2) & 0x30) | (_faf & 0xc0) | (_dagb & 0x200)
	_dac = ((_ead >> 2) & 0x70) | (_faf & 0xc0) | (_dagb & 0x700)
	var _fgcf int
	for _abb := 0; _abb < _dgd; _abb = _fgcf {
		var (
			_bd    int
			_fgcfb int
		)
		_fgcf = _abb + 8
		if _bd = _ddc - _abb; _bd > 8 {
			_bd = 8
		}
		_fag := _fgcf < _ddc
		_bgf := _fgcf < _eaa.ReferenceBitmap.Width
		_fad := _gaff + 1
		if _bca > 0 {
			_aff = 0
			if _fag {
				_aff, _dbf = _eaa.RegionBitmap.GetByte(_ae - _df + 1)
				if _dbf != nil {
					return
				}
			}
			_caf = (_caf << 8) | int(_aff)
		}
		if _eag > 0 && _eag <= _eaa.ReferenceBitmap.Height {
			var _be int
			if _bgf {
				_aff, _dbf = _eaa.ReferenceBitmap.GetByte(_cfg - _ab + _fad)
				if _dbf != nil {
					return
				}
				_be = int(_aff) << 2
			}
			_dagb = (_dagb << 8) | _be
		}
		if _eag >= 0 && _eag < _eaa.ReferenceBitmap.Height {
			_aff = 0
			if _bgf {
				_aff, _dbf = _eaa.ReferenceBitmap.GetByte(_cfg + _fad)
				if _dbf != nil {
					return
				}
			}
			_faf = (_faf << 8) | int(_aff)
		}
		if _eag > -2 && _eag < (_eaa.ReferenceBitmap.Height-1) {
			_aff = 0
			if _bgf {
				_aff, _dbf = _eaa.ReferenceBitmap.GetByte(_cfg + _ab + _fad)
				if _dbf != nil {
					return
				}
			}
			_ead = (_ead << 8) | int(_aff)
		}
		for _cgd := 0; _cgd < _bd; _cgd++ {
			var _abbf int
			_efa := (_dac >> 4) & 0x1ff
			switch _efa {
			case 0x1ff:
				_abbf = 1
			case 0x00:
				_abbf = 0
			default:
				_eaa._cdb.SetIndex(int32(_fdgf))
				_abbf, _dbf = _eaa._gf.DecodeBit(_eaa._cdb)
				if _dbf != nil {
					return
				}
			}
			_ebb := uint(7 - _cgd)
			_fgcfb |= _abbf << _ebb
			_fdgf = ((_fdgf & 0x0d6) << 1) | _abbf | (_caf>>_ebb+5)&0x002 | ((_ead>>_ebb + 2) & 0x010) | ((_faf >> _ebb) & 0x040) | ((_dagb >> _ebb) & 0x200)
			_dac = ((_dac & 0xdb) << 1) | ((_ead>>_ebb + 2) & 0x010) | ((_faf >> _ebb) & 0x080) | ((_dagb >> _ebb) & 0x400)
		}
		_dbf = _eaa.RegionBitmap.SetByte(_ae, byte(_fgcfb))
		if _dbf != nil {
			return
		}
		_ae++
		_cfg++
	}
	return nil
}

type Regioner interface {
	GetRegionBitmap() (*_d.Bitmap, error)
	GetRegionInfo() *RegionSegment
}

func (_bgdb *TextRegion) setParameters(_aebc *_b.Decoder, _fcgf, _acdf bool, _ccdga, _cdfe uint32, _gdadb uint32, _agadc int8, _abed uint32, _edbcbd int8, _fddd _d.CombinationOperator, _cbea int8, _egfgc int16, _ebea, _adbc, _aeac, _eagga, _fcgc, _caga, _bdda, _ebef, _gfeg, _abgg int8, _gfgc, _fecb []int8, _bcaf []*_d.Bitmap, _cdaaa int8) {
	_bgdb._dffe = _aebc
	_bgdb.IsHuffmanEncoded = _fcgf
	_bgdb.UseRefinement = _acdf
	_bgdb.RegionInfo.BitmapWidth = _ccdga
	_bgdb.RegionInfo.BitmapHeight = _cdfe
	_bgdb.NumberOfSymbolInstances = _gdadb
	_bgdb.SbStrips = _agadc
	_bgdb.NumberOfSymbols = _abed
	_bgdb.DefaultPixel = _edbcbd
	_bgdb.CombinationOperator = _fddd
	_bgdb.IsTransposed = _cbea
	_bgdb.ReferenceCorner = _egfgc
	_bgdb.SbDsOffset = _ebea
	_bgdb.SbHuffFS = _adbc
	_bgdb.SbHuffDS = _aeac
	_bgdb.SbHuffDT = _eagga
	_bgdb.SbHuffRDWidth = _fcgc
	_bgdb.SbHuffRDHeight = _caga
	_bgdb.SbHuffRSize = _gfeg
	_bgdb.SbHuffRDX = _bdda
	_bgdb.SbHuffRDY = _ebef
	_bgdb.SbrTemplate = _abgg
	_bgdb.SbrATX = _gfgc
	_bgdb.SbrATY = _fecb
	_bgdb.Symbols = _bcaf
	_bgdb._ebcb = _cdaaa
}
func (_caceb *GenericRegion) overrideAtTemplate2(_bbbd, _fgba, _gacb, _fagd, _agb int) int {
	_bbbd &= 0x3FB
	if _caceb.GBAtY[0] == 0 && _caceb.GBAtX[0] >= -int8(_agb) {
		_bbbd |= (_fagd >> uint(7-(int8(_agb)+_caceb.GBAtX[0])) & 0x1) << 2
	} else {
		_bbbd |= int(_caceb.getPixel(_fgba+int(_caceb.GBAtX[0]), _gacb+int(_caceb.GBAtY[0]))) << 2
	}
	return _bbbd
}
func (_deba *Header) writeSegmentPageAssociation(_efead _fg.BinaryWriter) (_efbe int, _ecef error) {
	const _dada = "w\u0072\u0069\u0074\u0065\u0053\u0065g\u006d\u0065\u006e\u0074\u0050\u0061\u0067\u0065\u0041s\u0073\u006f\u0063i\u0061t\u0069\u006f\u006e"
	if _deba.pageSize() != 4 {
		if _ecef = _efead.WriteByte(byte(_deba.PageAssociation)); _ecef != nil {
			return 0, _fd.Wrap(_ecef, _dada, "\u0070\u0061\u0067\u0065\u0053\u0069\u007a\u0065\u0020\u0021\u003d\u0020\u0034")
		}
		return 1, nil
	}
	_cfgf := make([]byte, 4)
	_fe.BigEndian.PutUint32(_cfgf, uint32(_deba.PageAssociation))
	if _efbe, _ecef = _efead.Write(_cfgf); _ecef != nil {
		return 0, _fd.Wrap(_ecef, _dada, "\u0034 \u0062y\u0074\u0065\u0020\u0070\u0061g\u0065\u0020n\u0075\u006d\u0062\u0065\u0072")
	}
	return _efbe, nil
}
func (_eac *GenericRegion) parseHeader() (_cbf error) {
	_da.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052I\u0043\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0050\u0061\u0072s\u0069\u006e\u0067\u0048\u0065\u0061\u0064e\u0072\u002e\u002e\u002e")
	defer func() {
		if _cbf != nil {
			_da.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0047\u0049\u004f\u004e]\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0048\u0065\u0061\u0064\u0065r\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u0020\u0077\u0069th\u0020\u0065\u0072\u0072\u006f\u0072\u002e\u0020\u0025\u0076", _cbf)
		} else {
			_da.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049C\u002d\u0052\u0045G\u0049\u004f\u004e]\u0020\u0050a\u0072\u0073\u0069\u006e\u0067\u0048e\u0061de\u0072\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u0020\u0053\u0075\u0063\u0063\u0065\u0073\u0073\u0066\u0075\u006c\u006c\u0079\u002e\u002e\u002e")
		}
	}()
	var (
		_bbb int
		_aec uint64
	)
	if _cbf = _eac.RegionSegment.parseHeader(); _cbf != nil {
		return _cbf
	}
	if _, _cbf = _eac._baf.ReadBits(3); _cbf != nil {
		return _cbf
	}
	_bbb, _cbf = _eac._baf.ReadBit()
	if _cbf != nil {
		return _cbf
	}
	if _bbb == 1 {
		_eac.UseExtTemplates = true
	}
	_bbb, _cbf = _eac._baf.ReadBit()
	if _cbf != nil {
		return _cbf
	}
	if _bbb == 1 {
		_eac.IsTPGDon = true
	}
	_aec, _cbf = _eac._baf.ReadBits(2)
	if _cbf != nil {
		return _cbf
	}
	_eac.GBTemplate = byte(_aec & 0xf)
	_bbb, _cbf = _eac._baf.ReadBit()
	if _cbf != nil {
		return _cbf
	}
	if _bbb == 1 {
		_eac.IsMMREncoded = true
	}
	if !_eac.IsMMREncoded {
		_ebf := 1
		if _eac.GBTemplate == 0 {
			_ebf = 4
			if _eac.UseExtTemplates {
				_ebf = 12
			}
		}
		if _cbf = _eac.readGBAtPixels(_ebf); _cbf != nil {
			return _cbf
		}
	}
	if _cbf = _eac.computeSegmentDataStructure(); _cbf != nil {
		return _cbf
	}
	_da.Log.Trace("\u0025\u0073", _eac)
	return nil
}
func (_dgf *GenericRefinementRegion) getGrReference() (*_d.Bitmap, error) {
	segments := _dgf._cdg.RTSegments
	if len(segments) == 0 {
		return nil, _ge.New("\u0052\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0065\u0078is\u0074\u0073")
	}
	_eg, _dcg := segments[0].GetSegmentData()
	if _dcg != nil {
		return nil, _dcg
	}
	_dce, _adg := _eg.(Regioner)
	if !_adg {
		return nil, _e.Errorf("\u0072\u0065\u0066\u0065\u0072r\u0065\u0064\u0020\u0074\u006f\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0052\u0065\u0067\u0069\u006f\u006e\u0065\u0072\u003a\u0020\u0025\u0054", _eg)
	}
	return _dce.GetRegionBitmap()
}
func (_dgfd *SymbolDictionary) encodeFlags(_dafe _fg.BinaryWriter) (_acgf int, _eggg error) {
	const _gdda = "e\u006e\u0063\u006f\u0064\u0065\u0046\u006c\u0061\u0067\u0073"
	if _eggg = _dafe.SkipBits(3); _eggg != nil {
		return 0, _fd.Wrap(_eggg, _gdda, "\u0065\u006d\u0070\u0074\u0079\u0020\u0062\u0069\u0074\u0073")
	}
	var _dceb int
	if _dgfd.SdrTemplate > 0 {
		_dceb = 1
	}
	if _eggg = _dafe.WriteBit(_dceb); _eggg != nil {
		return _acgf, _fd.Wrap(_eggg, _gdda, "s\u0064\u0072\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_dceb = 0
	if _dgfd.SdTemplate > 1 {
		_dceb = 1
	}
	if _eggg = _dafe.WriteBit(_dceb); _eggg != nil {
		return _acgf, _fd.Wrap(_eggg, _gdda, "\u0073\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_dceb = 0
	if _dgfd.SdTemplate == 1 || _dgfd.SdTemplate == 3 {
		_dceb = 1
	}
	if _eggg = _dafe.WriteBit(_dceb); _eggg != nil {
		return _acgf, _fd.Wrap(_eggg, _gdda, "\u0073\u0064\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	_dceb = 0
	if _dgfd._gdae {
		_dceb = 1
	}
	if _eggg = _dafe.WriteBit(_dceb); _eggg != nil {
		return _acgf, _fd.Wrap(_eggg, _gdda, "\u0063\u006f\u0064in\u0067\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0072\u0065\u0074\u0061\u0069\u006e\u0065\u0064")
	}
	_dceb = 0
	if _dgfd._gacg {
		_dceb = 1
	}
	if _eggg = _dafe.WriteBit(_dceb); _eggg != nil {
		return _acgf, _fd.Wrap(_eggg, _gdda, "\u0063\u006f\u0064\u0069ng\u0020\u0063\u006f\u006e\u0074\u0065\u0078\u0074\u0020\u0075\u0073\u0065\u0064")
	}
	_dceb = 0
	if _dgfd.SdHuffAggInstanceSelection {
		_dceb = 1
	}
	if _eggg = _dafe.WriteBit(_dceb); _eggg != nil {
		return _acgf, _fd.Wrap(_eggg, _gdda, "\u0073\u0064\u0048\u0075\u0066\u0066\u0041\u0067\u0067\u0049\u006e\u0073\u0074")
	}
	_dceb = int(_dgfd.SdHuffBMSizeSelection)
	if _eggg = _dafe.WriteBit(_dceb); _eggg != nil {
		return _acgf, _fd.Wrap(_eggg, _gdda, "\u0073\u0064\u0048u\u0066\u0066\u0042\u006d\u0053\u0069\u007a\u0065")
	}
	_dceb = 0
	if _dgfd.SdHuffDecodeWidthSelection > 1 {
		_dceb = 1
	}
	if _eggg = _dafe.WriteBit(_dceb); _eggg != nil {
		return _acgf, _fd.Wrap(_eggg, _gdda, "s\u0064\u0048\u0075\u0066\u0066\u0057\u0069\u0064\u0074\u0068")
	}
	_dceb = 0
	switch _dgfd.SdHuffDecodeWidthSelection {
	case 1, 3:
		_dceb = 1
	}
	if _eggg = _dafe.WriteBit(_dceb); _eggg != nil {
		return _acgf, _fd.Wrap(_eggg, _gdda, "s\u0064\u0048\u0075\u0066\u0066\u0057\u0069\u0064\u0074\u0068")
	}
	_dceb = 0
	if _dgfd.SdHuffDecodeHeightSelection > 1 {
		_dceb = 1
	}
	if _eggg = _dafe.WriteBit(_dceb); _eggg != nil {
		return _acgf, _fd.Wrap(_eggg, _gdda, "\u0073\u0064\u0048u\u0066\u0066\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_dceb = 0
	switch _dgfd.SdHuffDecodeHeightSelection {
	case 1, 3:
		_dceb = 1
	}
	if _eggg = _dafe.WriteBit(_dceb); _eggg != nil {
		return _acgf, _fd.Wrap(_eggg, _gdda, "\u0073\u0064\u0048u\u0066\u0066\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_dceb = 0
	if _dgfd.UseRefinementAggregation {
		_dceb = 1
	}
	if _eggg = _dafe.WriteBit(_dceb); _eggg != nil {
		return _acgf, _fd.Wrap(_eggg, _gdda, "\u0073\u0064\u0052\u0065\u0066\u0041\u0067\u0067")
	}
	_dceb = 0
	if _dgfd.IsHuffmanEncoded {
		_dceb = 1
	}
	if _eggg = _dafe.WriteBit(_dceb); _eggg != nil {
		return _acgf, _fd.Wrap(_eggg, _gdda, "\u0073\u0064\u0048\u0075\u0066\u0066")
	}
	return 2, nil
}
func (_ebbd *SymbolDictionary) readRefinementAtPixels(_gacgb int) error {
	_ebbd.SdrATX = make([]int8, _gacgb)
	_ebbd.SdrATY = make([]int8, _gacgb)
	var (
		_fafg byte
		_gcaa error
	)
	for _dcdd := 0; _dcdd < _gacgb; _dcdd++ {
		_fafg, _gcaa = _ebbd._bcagf.ReadByte()
		if _gcaa != nil {
			return _gcaa
		}
		_ebbd.SdrATX[_dcdd] = int8(_fafg)
		_fafg, _gcaa = _ebbd._bcagf.ReadByte()
		if _gcaa != nil {
			return _gcaa
		}
		_ebbd.SdrATY[_dcdd] = int8(_fafg)
	}
	return nil
}
func (_ecc *GenericRefinementRegion) decodeTemplate(_gege, _egg, _ee, _add, _dfc, _fgbf, _eagf, _eadc, _dga, _feg int, _ce templater) (_fdc error) {
	var (
		_edb, _fedf, _adeb, _cfa, _fddg int16
		_dbfd, _gge, _gac, _dgdb        int
		_dda                            byte
	)
	if _dga >= 1 && (_dga-1) < _ecc.ReferenceBitmap.Height {
		_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg - _add)
		if _fdc != nil {
			return
		}
		_dbfd = int(_dda)
	}
	if _dga >= 0 && (_dga) < _ecc.ReferenceBitmap.Height {
		_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg)
		if _fdc != nil {
			return
		}
		_gge = int(_dda)
	}
	if _dga >= -1 && (_dga+1) < _ecc.ReferenceBitmap.Height {
		_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg + _add)
		if _fdc != nil {
			return
		}
		_gac = int(_dda)
	}
	_feg++
	if _gege >= 1 {
		_dda, _fdc = _ecc.RegionBitmap.GetByte(_eadc - _ee)
		if _fdc != nil {
			return
		}
		_dgdb = int(_dda)
	}
	_eadc++
	_bgfa := _ecc.ReferenceDX % 8
	_cga := 6 + _bgfa
	_ega := _feg % _add
	if _cga >= 0 {
		if _cga < 8 {
			_edb = int16(_dbfd>>uint(_cga)) & 0x07
		}
		if _cga < 8 {
			_fedf = int16(_gge>>uint(_cga)) & 0x07
		}
		if _cga < 8 {
			_adeb = int16(_gac>>uint(_cga)) & 0x07
		}
		if _cga == 6 && _ega > 1 {
			if _dga >= 1 && (_dga-1) < _ecc.ReferenceBitmap.Height {
				_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg - _add - 2)
				if _fdc != nil {
					return _fdc
				}
				_edb |= int16(_dda<<2) & 0x04
			}
			if _dga >= 0 && _dga < _ecc.ReferenceBitmap.Height {
				_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg - 2)
				if _fdc != nil {
					return _fdc
				}
				_fedf |= int16(_dda<<2) & 0x04
			}
			if _dga >= -1 && _dga+1 < _ecc.ReferenceBitmap.Height {
				_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg + _add - 2)
				if _fdc != nil {
					return _fdc
				}
				_adeb |= int16(_dda<<2) & 0x04
			}
		}
		if _cga == 0 {
			_dbfd = 0
			_gge = 0
			_gac = 0
			if _ega < _add-1 {
				if _dga >= 1 && _dga-1 < _ecc.ReferenceBitmap.Height {
					_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg - _add)
					if _fdc != nil {
						return _fdc
					}
					_dbfd = int(_dda)
				}
				if _dga >= 0 && _dga < _ecc.ReferenceBitmap.Height {
					_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg)
					if _fdc != nil {
						return _fdc
					}
					_gge = int(_dda)
				}
				if _dga >= -1 && _dga+1 < _ecc.ReferenceBitmap.Height {
					_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg + _add)
					if _fdc != nil {
						return _fdc
					}
					_gac = int(_dda)
				}
			}
			_feg++
		}
	} else {
		_edb = int16(_dbfd<<1) & 0x07
		_fedf = int16(_gge<<1) & 0x07
		_adeb = int16(_gac<<1) & 0x07
		_dbfd = 0
		_gge = 0
		_gac = 0
		if _ega < _add-1 {
			if _dga >= 1 && _dga-1 < _ecc.ReferenceBitmap.Height {
				_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg - _add)
				if _fdc != nil {
					return _fdc
				}
				_dbfd = int(_dda)
			}
			if _dga >= 0 && _dga < _ecc.ReferenceBitmap.Height {
				_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg)
				if _fdc != nil {
					return _fdc
				}
				_gge = int(_dda)
			}
			if _dga >= -1 && _dga+1 < _ecc.ReferenceBitmap.Height {
				_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg + _add)
				if _fdc != nil {
					return _fdc
				}
				_gac = int(_dda)
			}
			_feg++
		}
		_edb |= int16((_dbfd >> 7) & 0x07)
		_fedf |= int16((_gge >> 7) & 0x07)
		_adeb |= int16((_gac >> 7) & 0x07)
	}
	_cfa = int16(_dgdb >> 6)
	_fddg = 0
	_cde := (2 - _bgfa) % 8
	_dbfd <<= uint(_cde)
	_gge <<= uint(_cde)
	_gac <<= uint(_cde)
	_dgdb <<= 2
	var _ece int
	for _eeb := 0; _eeb < _egg; _eeb++ {
		_dgg := _eeb & 0x07
		_efbg := _ce.form(_edb, _fedf, _adeb, _cfa, _fddg)
		if _ecc._dag {
			_dda, _fdc = _ecc.RegionBitmap.GetByte(_ecc.RegionBitmap.GetByteIndex(_eeb, _gege))
			if _fdc != nil {
				return _fdc
			}
			_ecc._cdb.SetIndex(int32(_ecc.overrideAtTemplate0(int(_efbg), _eeb, _gege, int(_dda), _dgg)))
		} else {
			_ecc._cdb.SetIndex(int32(_efbg))
		}
		_ece, _fdc = _ecc._gf.DecodeBit(_ecc._cdb)
		if _fdc != nil {
			return _fdc
		}
		if _fdc = _ecc.RegionBitmap.SetPixel(_eeb, _gege, byte(_ece)); _fdc != nil {
			return _fdc
		}
		_edb = ((_edb << 1) | 0x01&int16(_dbfd>>7)) & 0x07
		_fedf = ((_fedf << 1) | 0x01&int16(_gge>>7)) & 0x07
		_adeb = ((_adeb << 1) | 0x01&int16(_gac>>7)) & 0x07
		_cfa = ((_cfa << 1) | 0x01&int16(_dgdb>>7)) & 0x07
		_fddg = int16(_ece)
		if (_eeb-int(_ecc.ReferenceDX))%8 == 5 {
			_dbfd = 0
			_gge = 0
			_gac = 0
			if ((_eeb-int(_ecc.ReferenceDX))/8)+1 < _ecc.ReferenceBitmap.RowStride {
				if _dga >= 1 && (_dga-1) < _ecc.ReferenceBitmap.Height {
					_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg - _add)
					if _fdc != nil {
						return _fdc
					}
					_dbfd = int(_dda)
				}
				if _dga >= 0 && _dga < _ecc.ReferenceBitmap.Height {
					_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg)
					if _fdc != nil {
						return _fdc
					}
					_gge = int(_dda)
				}
				if _dga >= -1 && (_dga+1) < _ecc.ReferenceBitmap.Height {
					_dda, _fdc = _ecc.ReferenceBitmap.GetByte(_feg + _add)
					if _fdc != nil {
						return _fdc
					}
					_gac = int(_dda)
				}
			}
			_feg++
		} else {
			_dbfd <<= 1
			_gge <<= 1
			_gac <<= 1
		}
		if _dgg == 5 && _gege >= 1 {
			if ((_eeb >> 3) + 1) >= _ecc.RegionBitmap.RowStride {
				_dgdb = 0
			} else {
				_dda, _fdc = _ecc.RegionBitmap.GetByte(_eadc - _ee)
				if _fdc != nil {
					return _fdc
				}
				_dgdb = int(_dda)
			}
			_eadc++
		} else {
			_dgdb <<= 1
		}
	}
	return nil
}
func (_bcgf *HalftoneRegion) computeGrayScalePlanes(_fecde []*_d.Bitmap, _cddb int) ([][]int, error) {
	_dfcb := make([][]int, _bcgf.HGridHeight)
	for _bbeg := 0; _bbeg < len(_dfcb); _bbeg++ {
		_dfcb[_bbeg] = make([]int, _bcgf.HGridWidth)
	}
	for _cfe := 0; _cfe < int(_bcgf.HGridHeight); _cfe++ {
		for _agebf := 0; _agebf < int(_bcgf.HGridWidth); _agebf += 8 {
			var _bcb int
			if _ccdf := int(_bcgf.HGridWidth) - _agebf; _ccdf > 8 {
				_bcb = 8
			} else {
				_bcb = _ccdf
			}
			_cag := _fecde[0].GetByteIndex(_agebf, _cfe)
			for _feaf := 0; _feaf < _bcb; _feaf++ {
				_cgfb := _feaf + _agebf
				_dfcb[_cfe][_cgfb] = 0
				for _bcde := 0; _bcde < _cddb; _bcde++ {
					_ceba, _addf := _fecde[_bcde].GetByte(_cag)
					if _addf != nil {
						return nil, _addf
					}
					_dagg := _ceba >> uint(7-_cgfb&7)
					_dgfe := _dagg & 1
					_bba := 1 << uint(_bcde)
					_dgeb := int(_dgfe) * _bba
					_dfcb[_cfe][_cgfb] += _dgeb
				}
			}
		}
	}
	return _dfcb, nil
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
	Reader                   _fg.StreamReader
	SegmentData              Segmenter
	RTSNumbers               []int
	RetainBits               []uint8
}

func (_fcde *TextRegion) blit(_dcbde *_d.Bitmap, _ddgb int64) error {
	if _fcde.IsTransposed == 0 && (_fcde.ReferenceCorner == 2 || _fcde.ReferenceCorner == 3) {
		_fcde._fbdbc += int64(_dcbde.Width - 1)
	} else if _fcde.IsTransposed == 1 && (_fcde.ReferenceCorner == 0 || _fcde.ReferenceCorner == 2) {
		_fcde._fbdbc += int64(_dcbde.Height - 1)
	}
	_agcgc := _fcde._fbdbc
	if _fcde.IsTransposed == 1 {
		_agcgc, _ddgb = _ddgb, _agcgc
	}
	switch _fcde.ReferenceCorner {
	case 0:
		_ddgb -= int64(_dcbde.Height - 1)
	case 2:
		_ddgb -= int64(_dcbde.Height - 1)
		_agcgc -= int64(_dcbde.Width - 1)
	case 3:
		_agcgc -= int64(_dcbde.Width - 1)
	}
	_bfaa := _d.Blit(_dcbde, _fcde.RegionBitmap, int(_agcgc), int(_ddgb), _fcde.CombinationOperator)
	if _bfaa != nil {
		return _bfaa
	}
	if _fcde.IsTransposed == 0 && (_fcde.ReferenceCorner == 0 || _fcde.ReferenceCorner == 1) {
		_fcde._fbdbc += int64(_dcbde.Width - 1)
	}
	if _fcde.IsTransposed == 1 && (_fcde.ReferenceCorner == 1 || _fcde.ReferenceCorner == 3) {
		_fcde._fbdbc += int64(_dcbde.Height - 1)
	}
	return nil
}
func (_baed *SymbolDictionary) readNumberOfNewSymbols() error {
	_edcb, _agad := _baed._bcagf.ReadBits(32)
	if _agad != nil {
		return _agad
	}
	_baed.NumberOfNewSymbols = uint32(_edcb & _f.MaxUint32)
	return nil
}
func (_dcbd *SymbolDictionary) GetDictionary() ([]*_d.Bitmap, error) {
	_da.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c-\u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		_da.Log.Trace("\u005b\u0053\u0059M\u0042\u004f\u004c\u002d\u0044\u0049\u0043\u0054\u0049\u004f\u004e\u0041\u0052\u0059\u005d\u0020\u0047\u0065\u0074\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
		_da.Log.Trace("\u005b\u0053Y\u004d\u0042\u004f\u004c\u002dD\u0049\u0043\u0054\u0049\u004fN\u0041\u0052\u0059\u005d\u0020\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0020\u000a\u0045\u0078\u003a\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u000a\u006e\u0065\u0077\u003a\u0027\u0025\u0073\u0027", _dcbd._cba, _dcbd._gbga)
	}()
	if _dcbd._cba == nil {
		var _cbgd error
		if _dcbd.UseRefinementAggregation {
			_dcbd._cff = _dcbd.getSbSymCodeLen()
		}
		if !_dcbd.IsHuffmanEncoded {
			if _cbgd = _dcbd.setCodingStatistics(); _cbgd != nil {
				return nil, _cbgd
			}
		}
		_dcbd._gbga = make([]*_d.Bitmap, _dcbd.NumberOfNewSymbols)
		var _dfaa []int
		if _dcbd.IsHuffmanEncoded && !_dcbd.UseRefinementAggregation {
			_dfaa = make([]int, _dcbd.NumberOfNewSymbols)
		}
		if _cbgd = _dcbd.setSymbolsArray(); _cbgd != nil {
			return nil, _cbgd
		}
		var _cee, _dgede int64
		_dcbd._ffa = 0
		for _dcbd._ffa < _dcbd.NumberOfNewSymbols {
			_dgede, _cbgd = _dcbd.decodeHeightClassDeltaHeight()
			if _cbgd != nil {
				return nil, _cbgd
			}
			_cee += _dgede
			var _daag, _fcdc uint32
			_eeed := int64(_dcbd._ffa)
			for {
				var _fgfc int64
				_fgfc, _cbgd = _dcbd.decodeDifferenceWidth()
				if _cd.Is(_cbgd, _gd.ErrOOB) {
					break
				}
				if _cbgd != nil {
					return nil, _cbgd
				}
				if _dcbd._ffa >= _dcbd.NumberOfNewSymbols {
					break
				}
				_daag += uint32(_fgfc)
				_fcdc += _daag
				if !_dcbd.IsHuffmanEncoded || _dcbd.UseRefinementAggregation {
					if !_dcbd.UseRefinementAggregation {
						_cbgd = _dcbd.decodeDirectlyThroughGenericRegion(_daag, uint32(_cee))
						if _cbgd != nil {
							return nil, _cbgd
						}
					} else {
						_cbgd = _dcbd.decodeAggregate(_daag, uint32(_cee))
						if _cbgd != nil {
							return nil, _cbgd
						}
					}
				} else if _dcbd.IsHuffmanEncoded && !_dcbd.UseRefinementAggregation {
					_dfaa[_dcbd._ffa] = int(_daag)
				}
				_dcbd._ffa++
			}
			if _dcbd.IsHuffmanEncoded && !_dcbd.UseRefinementAggregation {
				var _fcfb int64
				if _dcbd.SdHuffBMSizeSelection == 0 {
					var _cfce _ba.Tabler
					_cfce, _cbgd = _ba.GetStandardTable(1)
					if _cbgd != nil {
						return nil, _cbgd
					}
					_fcfb, _cbgd = _cfce.Decode(_dcbd._bcagf)
					if _cbgd != nil {
						return nil, _cbgd
					}
				} else {
					_fcfb, _cbgd = _dcbd.huffDecodeBmSize()
					if _cbgd != nil {
						return nil, _cbgd
					}
				}
				_dcbd._bcagf.Align()
				var _ced *_d.Bitmap
				_ced, _cbgd = _dcbd.decodeHeightClassCollectiveBitmap(_fcfb, uint32(_cee), _fcdc)
				if _cbgd != nil {
					return nil, _cbgd
				}
				_cbgd = _dcbd.decodeHeightClassBitmap(_ced, _eeed, int(_cee), _dfaa)
				if _cbgd != nil {
					return nil, _cbgd
				}
			}
		}
		_cbde, _cbgd := _dcbd.getToExportFlags()
		if _cbgd != nil {
			return nil, _cbgd
		}
		_dcbd.setExportedSymbols(_cbde)
	}
	return _dcbd._cba, nil
}
func (_dgff *TextRegion) Init(header *Header, r _fg.StreamReader) error {
	_dgff.Header = header
	_dgff._effb = r
	_dgff.RegionInfo = NewRegionSegment(_dgff._effb)
	return _dgff.parseHeader()
}
func (_abbbd *SymbolDictionary) setRefinementAtPixels() error {
	if !_abbbd.UseRefinementAggregation || _abbbd.SdrTemplate != 0 {
		return nil
	}
	if _ggbd := _abbbd.readRefinementAtPixels(2); _ggbd != nil {
		return _ggbd
	}
	return nil
}
func (_bafg *SymbolDictionary) decodeRefinedSymbol(_gdccb, _fadcg uint32) error {
	var (
		_cdbe         int
		_bgac, _cacfc int32
	)
	if _bafg.IsHuffmanEncoded {
		_edab, _faag := _bafg._bcagf.ReadBits(byte(_bafg._cff))
		if _faag != nil {
			return _faag
		}
		_cdbe = int(_edab)
		_egff, _faag := _ba.GetStandardTable(15)
		if _faag != nil {
			return _faag
		}
		_ddfcb, _faag := _egff.Decode(_bafg._bcagf)
		if _faag != nil {
			return _faag
		}
		_bgac = int32(_ddfcb)
		_ddfcb, _faag = _egff.Decode(_bafg._bcagf)
		if _faag != nil {
			return _faag
		}
		_cacfc = int32(_ddfcb)
		_egff, _faag = _ba.GetStandardTable(1)
		if _faag != nil {
			return _faag
		}
		if _, _faag = _egff.Decode(_bafg._bcagf); _faag != nil {
			return _faag
		}
		_bafg._bcagf.Align()
	} else {
		_gaab, _aeag := _bafg._fffd.DecodeIAID(uint64(_bafg._cff), _bafg._fac)
		if _aeag != nil {
			return _aeag
		}
		_cdbe = int(_gaab)
		_bgac, _aeag = _bafg._fffd.DecodeInt(_bafg._abag)
		if _aeag != nil {
			return _aeag
		}
		_cacfc, _aeag = _bafg._fffd.DecodeInt(_bafg._dfb)
		if _aeag != nil {
			return _aeag
		}
	}
	if _aefgf := _bafg.setSymbolsArray(); _aefgf != nil {
		return _aefgf
	}
	_adbf := _bafg._bded[_cdbe]
	if _gbfb := _bafg.decodeNewSymbols(_gdccb, _fadcg, _adbf, _bgac, _cacfc); _gbfb != nil {
		return _gbfb
	}
	if _bafg.IsHuffmanEncoded {
		_bafg._bcagf.Align()
	}
	return nil
}
func (_bfef *HalftoneRegion) GetRegionBitmap() (*_d.Bitmap, error) {
	if _bfef.HalftoneRegionBitmap != nil {
		return _bfef.HalftoneRegionBitmap, nil
	}
	var _abfe error
	_bfef.HalftoneRegionBitmap = _d.New(int(_bfef.RegionSegment.BitmapWidth), int(_bfef.RegionSegment.BitmapHeight))
	if _bfef.Patterns == nil || len(_bfef.Patterns) == 0 {
		_bfef.Patterns, _abfe = _bfef.GetPatterns()
		if _abfe != nil {
			return nil, _abfe
		}
	}
	if _bfef.HDefaultPixel == 1 {
		_bfef.HalftoneRegionBitmap.SetDefaultPixel()
	}
	_bdff := _f.Ceil(_f.Log(float64(len(_bfef.Patterns))) / _f.Log(2))
	_ccdc := int(_bdff)
	var _bedb [][]int
	_bedb, _abfe = _bfef.grayScaleDecoding(_ccdc)
	if _abfe != nil {
		return nil, _abfe
	}
	if _abfe = _bfef.renderPattern(_bedb); _abfe != nil {
		return nil, _abfe
	}
	return _bfef.HalftoneRegionBitmap, nil
}
func (_baaa *PageInformationSegment) readRequiresAuxiliaryBuffer() error {
	_cgcb, _efgg := _baaa._ace.ReadBit()
	if _efgg != nil {
		return _efgg
	}
	if _cgcb == 1 {
		_baaa._bfgd = true
	}
	return nil
}
func (_aad *HalftoneRegion) computeY(_abbe, _beg int) int {
	return _aad.shiftAndFill(int(_aad.HGridY) + _abbe*int(_aad.HRegionX) - _beg*int(_aad.HRegionY))
}

type Segmenter interface {
	Init(_aag *Header, _feea _fg.StreamReader) error
}

func (_ageb *GenericRegion) setParametersWithAt(_abf bool, _gbac byte, _fbea, _abgc bool, _bdf, _aabb []int8, _bfg, _fgga uint32, _eadd *_b.DecoderStats, _fcdg *_b.Decoder) {
	_ageb.IsMMREncoded = _abf
	_ageb.GBTemplate = _gbac
	_ageb.IsTPGDon = _fbea
	_ageb.GBAtX = _bdf
	_ageb.GBAtY = _aabb
	_ageb.RegionSegment.BitmapHeight = _fgga
	_ageb.RegionSegment.BitmapWidth = _bfg
	_ageb._gcd = nil
	_ageb.Bitmap = nil
	if _eadd != nil {
		_ageb._aae = _eadd
	}
	if _fcdg != nil {
		_ageb._acb = _fcdg
	}
	_da.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0047\u0049O\u004e\u005d\u0020\u0073\u0065\u0074P\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0053\u0044\u0041t\u003a\u0020\u0025\u0073", _ageb)
}
func (_aea *PatternDictionary) computeSegmentDataStructure() error {
	_aea.DataOffset = _aea._fbdd.StreamPosition()
	_aea.DataHeaderLength = _aea.DataOffset - _aea.DataHeaderOffset
	_aea.DataLength = int64(_aea._fbdd.Length()) - _aea.DataHeaderLength
	return nil
}
func (_gfff *Header) readNumberOfReferredToSegments(_bcbd _fg.StreamReader) (uint64, error) {
	const _fcfe = "\u0072\u0065\u0061\u0064\u004e\u0075\u006d\u0062\u0065\u0072O\u0066\u0052\u0065\u0066\u0065\u0072\u0072e\u0064\u0054\u006f\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073"
	_abdc, _daege := _bcbd.ReadBits(3)
	if _daege != nil {
		return 0, _fd.Wrap(_daege, _fcfe, "\u0063\u006f\u0075n\u0074\u0020\u006f\u0066\u0020\u0072\u0074\u0073")
	}
	_abdc &= 0xf
	var _dadef []byte
	if _abdc <= 4 {
		_dadef = make([]byte, 5)
		for _fgcae := 0; _fgcae <= 4; _fgcae++ {
			_beae, _ggead := _bcbd.ReadBit()
			if _ggead != nil {
				return 0, _fd.Wrap(_ggead, _fcfe, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
			}
			_dadef[_fgcae] = byte(_beae)
		}
	} else {
		_abdc, _daege = _bcbd.ReadBits(29)
		if _daege != nil {
			return 0, _daege
		}
		_abdc &= _f.MaxInt32
		_gedd := (_abdc + 8) >> 3
		_gedd <<= 3
		_dadef = make([]byte, _gedd)
		var _bgcd uint64
		for _bgcd = 0; _bgcd < _gedd; _bgcd++ {
			_ecda, _bgbc := _bcbd.ReadBit()
			if _bgbc != nil {
				return 0, _fd.Wrap(_bgbc, _fcfe, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
			}
			_dadef[_bgcd] = byte(_ecda)
		}
	}
	return _abdc, nil
}
func (_bgec *SymbolDictionary) decodeDirectlyThroughGenericRegion(_dgbf, _ebbe uint32) error {
	if _bgec._eaeb == nil {
		_bgec._eaeb = NewGenericRegion(_bgec._bcagf)
	}
	_bgec._eaeb.setParametersWithAt(false, byte(_bgec.SdTemplate), false, false, _bgec.SdATX, _bgec.SdATY, _dgbf, _ebbe, _bgec._bagf, _bgec._fffd)
	return _bgec.addSymbol(_bgec._eaeb)
}
func (_fage *SymbolDictionary) setExportedSymbols(_dfdg []int) {
	for _fdgce := uint32(0); _fdgce < _fage._gbbg+_fage.NumberOfNewSymbols; _fdgce++ {
		if _dfdg[_fdgce] == 1 {
			var _ggee *_d.Bitmap
			if _fdgce < _fage._gbbg {
				_ggee = _fage._cgfbb[_fdgce]
			} else {
				_ggee = _fage._gbga[_fdgce-_fage._gbbg]
			}
			_da.Log.Trace("\u005bS\u0059\u004dB\u004f\u004c\u002d\u0044I\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u0020\u0041\u0064\u0064 E\u0078\u0070\u006fr\u0074\u0065d\u0053\u0079\u006d\u0062\u006f\u006c:\u0020\u0027%\u0073\u0027", _ggee)
			_fage._cba = append(_fage._cba, _ggee)
		}
	}
}
func (_cddba *Header) readDataStartOffset(_edfd _fg.StreamReader, _egbg OrganizationType) {
	if _egbg == OSequential {
		_cddba.SegmentDataStartOffset = uint64(_edfd.StreamPosition())
	}
}
func (_ebc *GenericRefinementRegion) decodeTypicalPredictedLine(_gdf, _cfb, _ffb, _fedd, _cdf, _gda int) error {
	_gdg := _gdf - int(_ebc.ReferenceDY)
	_dbg := _ebc.ReferenceBitmap.GetByteIndex(0, _gdg)
	_dcd := _ebc.RegionBitmap.GetByteIndex(0, _gdf)
	var _gef error
	switch _ebc.TemplateID {
	case 0:
		_gef = _ebc.decodeTypicalPredictedLineTemplate0(_gdf, _cfb, _ffb, _fedd, _cdf, _gda, _dcd, _gdg, _dbg)
	case 1:
		_gef = _ebc.decodeTypicalPredictedLineTemplate1(_gdf, _cfb, _ffb, _fedd, _cdf, _gda, _dcd, _gdg, _dbg)
	}
	return _gef
}
func (_aed *GenericRegion) GetRegionInfo() *RegionSegment { return _aed.RegionSegment }
func (_fbec *Header) writeReferredToCount(_efea _fg.BinaryWriter) (_dddg int, _gegb error) {
	const _dged = "w\u0072i\u0074\u0065\u0052\u0065\u0066\u0065\u0072\u0072e\u0064\u0054\u006f\u0043ou\u006e\u0074"
	_fbec.RTSNumbers = make([]int, len(_fbec.RTSegments))
	for _feee, _gbg := range _fbec.RTSegments {
		_fbec.RTSNumbers[_feee] = int(_gbg.SegmentNumber)
	}
	if len(_fbec.RTSNumbers) <= 4 {
		var _eacd byte
		if len(_fbec.RetainBits) >= 1 {
			_eacd = _fbec.RetainBits[0]
		}
		_eacd |= byte(len(_fbec.RTSNumbers)) << 5
		if _gegb = _efea.WriteByte(_eacd); _gegb != nil {
			return 0, _fd.Wrap(_gegb, _dged, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		return 1, nil
	}
	_bac := uint32(len(_fbec.RTSNumbers))
	_ceff := make([]byte, 4+_ca.Ceil(len(_fbec.RTSNumbers)+1, 8))
	_bac |= 0x7 << 29
	_fe.BigEndian.PutUint32(_ceff, _bac)
	copy(_ceff[1:], _fbec.RetainBits)
	_dddg, _gegb = _efea.Write(_ceff)
	if _gegb != nil {
		return 0, _fd.Wrap(_gegb, _dged, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	return _dddg, nil
}

type templater interface {
	form(_eebb, _fabb, _gcge, _cac, _gff int16) int16
	setIndex(_dfd *_b.DecoderStats)
}

func (_ggef *TextRegion) decodeStripT() (_deeaf int64, _ccfd error) {
	if _ggef.IsHuffmanEncoded {
		if _ggef.SbHuffDT == 3 {
			if _ggef._gfad == nil {
				var _cafg int
				if _ggef.SbHuffFS == 3 {
					_cafg++
				}
				if _ggef.SbHuffDS == 3 {
					_cafg++
				}
				_ggef._gfad, _ccfd = _ggef.getUserTable(_cafg)
				if _ccfd != nil {
					return 0, _ccfd
				}
			}
			_deeaf, _ccfd = _ggef._gfad.Decode(_ggef._effb)
			if _ccfd != nil {
				return 0, _ccfd
			}
		} else {
			var _deeg _ba.Tabler
			_deeg, _ccfd = _ba.GetStandardTable(11 + int(_ggef.SbHuffDT))
			if _ccfd != nil {
				return 0, _ccfd
			}
			_deeaf, _ccfd = _deeg.Decode(_ggef._effb)
			if _ccfd != nil {
				return 0, _ccfd
			}
		}
	} else {
		var _eddda int32
		_eddda, _ccfd = _ggef._dffe.DecodeInt(_ggef._bgad)
		if _ccfd != nil {
			return 0, _ccfd
		}
		_deeaf = int64(_eddda)
	}
	_deeaf *= int64(-_ggef.SbStrips)
	return _deeaf, nil
}
func (_bebb *TextRegion) decodeRI() (int64, error) {
	if !_bebb.UseRefinement {
		return 0, nil
	}
	if _bebb.IsHuffmanEncoded {
		_eaaf, _bagfb := _bebb._effb.ReadBit()
		return int64(_eaaf), _bagfb
	}
	_ecge, _gcfb := _bebb._dffe.DecodeInt(_bebb._gbfge)
	return int64(_ecge), _gcfb
}
func (_edff *PageInformationSegment) encodeStripingInformation(_ddbe _fg.BinaryWriter) (_cfc int, _efgf error) {
	const _fgee = "\u0065n\u0063\u006f\u0064\u0065S\u0074\u0072\u0069\u0070\u0069n\u0067I\u006ef\u006f\u0072\u006d\u0061\u0074\u0069\u006fn"
	if !_edff.IsStripe {
		if _cfc, _efgf = _ddbe.Write([]byte{0x00, 0x00}); _efgf != nil {
			return 0, _fd.Wrap(_efgf, _fgee, "n\u006f\u0020\u0073\u0074\u0072\u0069\u0070\u0069\u006e\u0067")
		}
		return _cfc, nil
	}
	_cbd := make([]byte, 2)
	_fe.BigEndian.PutUint16(_cbd, _edff.MaxStripeSize|1<<15)
	if _cfc, _efgf = _ddbe.Write(_cbd); _efgf != nil {
		return 0, _fd.Wrapf(_efgf, _fgee, "\u0073\u0074\u0072i\u0070\u0069\u006e\u0067\u003a\u0020\u0025\u0064", _edff.MaxStripeSize)
	}
	return _cfc, nil
}

type template1 struct{}

func (_effdd *PageInformationSegment) Init(h *Header, r _fg.StreamReader) (_fdec error) {
	_effdd._ace = r
	if _fdec = _effdd.parseHeader(); _fdec != nil {
		return _fd.Wrap(_fdec, "P\u0061\u0067\u0065\u0049\u006e\u0066o\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0053\u0065g\u006d\u0065\u006et\u002eI\u006e\u0069\u0074", "")
	}
	return nil
}
func (_gcadd *SymbolDictionary) checkInput() error {
	if _gcadd.SdHuffDecodeHeightSelection == 2 {
		_da.Log.Debug("\u0053\u0079\u006d\u0062\u006fl\u0020\u0044\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0048\u0065\u0069\u0067\u0068\u0074\u0020\u0053e\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0070\u0065r\u006d\u0069\u0074\u0074\u0065\u0064", _gcadd.SdHuffDecodeHeightSelection)
	}
	if _gcadd.SdHuffDecodeWidthSelection == 2 {
		_da.Log.Debug("\u0053\u0079\u006d\u0062\u006f\u006c\u0020\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079 \u0044\u0065\u0063\u006f\u0064\u0065\u0020\u0057\u0069\u0064t\u0068\u0020\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064\u0020\u0076\u0061l\u0075\u0065\u0020\u006e\u006f\u0074 \u0070\u0065r\u006d\u0069t\u0074e\u0064", _gcadd.SdHuffDecodeWidthSelection)
	}
	if _gcadd.IsHuffmanEncoded {
		if _gcadd.SdTemplate != 0 {
			_da.Log.Debug("\u0053\u0044T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u003d\u0020\u0025\u0064\u0020\u0028\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e \u0030\u0029", _gcadd.SdTemplate)
		}
		if !_gcadd.UseRefinementAggregation {
			if !_gcadd.UseRefinementAggregation {
				if _gcadd._gdae {
					_da.Log.Debug("\u0049\u0073\u0043\u006f\u0064\u0069\u006e\u0067C\u006f\u006e\u0074ex\u0074\u0052\u0065\u0074\u0061\u0069n\u0065\u0064\u0020\u003d\u0020\u0074\u0072\u0075\u0065\u0020\u0028\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0066\u0061\u006cs\u0065\u0029")
					_gcadd._gdae = false
				}
				if _gcadd._gacg {
					_da.Log.Debug("\u0069s\u0043\u006fd\u0069\u006e\u0067\u0043o\u006e\u0074\u0065x\u0074\u0055\u0073\u0065\u0064\u0020\u003d\u0020\u0074ru\u0065\u0020\u0028s\u0068\u006fu\u006c\u0064\u0020\u0062\u0065\u0020f\u0061\u006cs\u0065\u0029")
					_gcadd._gacg = false
				}
			}
		}
	} else {
		if _gcadd.SdHuffBMSizeSelection != 0 {
			_da.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066B\u004d\u0053\u0069\u007a\u0065\u0053\u0065\u006c\u0065\u0063\u0074\u0069\u006fn\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0030")
			_gcadd.SdHuffBMSizeSelection = 0
		}
		if _gcadd.SdHuffDecodeWidthSelection != 0 {
			_da.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0057\u0069\u0064\u0074\u0068\u0053\u0065\u006c\u0065\u0063\u0074\u0069o\u006e\u0020\u0073\u0068\u006fu\u006c\u0064 \u0062\u0065\u0020\u0030")
			_gcadd.SdHuffDecodeWidthSelection = 0
		}
		if _gcadd.SdHuffDecodeHeightSelection != 0 {
			_da.Log.Debug("\u0053\u0064\u0048\u0075\u0066\u0066\u0044\u0065\u0063\u006f\u0064\u0065\u0048e\u0069\u0067\u0068\u0074\u0053\u0065l\u0065\u0063\u0074\u0069\u006f\u006e\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0030")
			_gcadd.SdHuffDecodeHeightSelection = 0
		}
	}
	if !_gcadd.UseRefinementAggregation {
		if _gcadd.SdrTemplate != 0 {
			_da.Log.Debug("\u0053\u0044\u0052\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u003d\u0020\u0025\u0064\u0020\u0028s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030\u0029", _gcadd.SdrTemplate)
			_gcadd.SdrTemplate = 0
		}
	}
	if !_gcadd.IsHuffmanEncoded || !_gcadd.UseRefinementAggregation {
		if _gcadd.SdHuffAggInstanceSelection {
			_da.Log.Debug("\u0053d\u0048\u0075f\u0066\u0041\u0067g\u0049\u006e\u0073\u0074\u0061\u006e\u0063e\u0053\u0065\u006c\u0065\u0063\u0074i\u006f\u006e\u0020\u003d\u0020\u0025\u0064\u0020\u0028\u0073\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030\u0029", _gcadd.SdHuffAggInstanceSelection)
		}
	}
	return nil
}
func (_cecb *TextRegion) GetRegionInfo() *RegionSegment { return _cecb.RegionInfo }
func (_gebe *TextRegion) parseHeader() error {
	var _bebd error
	_da.Log.Trace("\u005b\u0054E\u0058\u0054\u0020\u0052E\u0047\u0049O\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045-\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0062\u0065\u0067\u0069n\u0073\u002e\u002e\u002e")
	defer func() {
		if _bebd != nil {
			_da.Log.Trace("\u005b\u0054\u0045\u0058\u0054\u0020\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044E\u0052\u005d\u0020\u0066\u0061i\u006c\u0065d\u002e\u0020\u0025\u0076", _bebd)
		} else {
			_da.Log.Trace("\u005b\u0054E\u0058\u0054\u0020\u0052E\u0047\u0049O\u004e\u005d\u005b\u0050\u0041\u0052\u0053\u0045-\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0066\u0069\u006e\u0069s\u0068\u0065\u0064\u002e")
		}
	}()
	if _bebd = _gebe.RegionInfo.parseHeader(); _bebd != nil {
		return _bebd
	}
	if _bebd = _gebe.readRegionFlags(); _bebd != nil {
		return _bebd
	}
	if _gebe.IsHuffmanEncoded {
		if _bebd = _gebe.readHuffmanFlags(); _bebd != nil {
			return _bebd
		}
	}
	if _bebd = _gebe.readUseRefinement(); _bebd != nil {
		return _bebd
	}
	if _bebd = _gebe.readAmountOfSymbolInstances(); _bebd != nil {
		return _bebd
	}
	if _bebd = _gebe.getSymbols(); _bebd != nil {
		return _bebd
	}
	if _bebd = _gebe.computeSymbolCodeLength(); _bebd != nil {
		return _bebd
	}
	if _bebd = _gebe.checkInput(); _bebd != nil {
		return _bebd
	}
	_da.Log.Trace("\u0025\u0073", _gebe.String())
	return nil
}
func (_dgbe *TextRegion) setContexts(_cdbee *_b.DecoderStats, _gadb *_b.DecoderStats, _cafbg *_b.DecoderStats, _fcgg *_b.DecoderStats, _cccf *_b.DecoderStats, _bgged *_b.DecoderStats, _cadega *_b.DecoderStats, _dcgfd *_b.DecoderStats, _fgdc *_b.DecoderStats, _cbe *_b.DecoderStats) {
	_dgbe._bgad = _gadb
	_dgbe._geaa = _cafbg
	_dgbe._dafb = _fcgg
	_dgbe._dace = _cccf
	_dgbe._egcff = _cadega
	_dgbe._acgg = _dcgfd
	_dgbe._beabb = _bgged
	_dgbe._beda = _fgdc
	_dgbe._bfbb = _cbe
	_dgbe._aeed = _cdbee
}
func (_abae *TextRegion) computeSymbolCodeLength() error {
	if _abae.IsHuffmanEncoded {
		return _abae.symbolIDCodeLengths()
	}
	_abae._ebcb = int8(_f.Ceil(_f.Log(float64(_abae.NumberOfSymbols)) / _f.Log(2)))
	return nil
}
func (_cbdf *TextRegion) readRegionFlags() error {
	var (
		_eace  int
		_ffdcg uint64
		_fcdd  error
	)
	_eace, _fcdd = _cbdf._effb.ReadBit()
	if _fcdd != nil {
		return _fcdd
	}
	_cbdf.SbrTemplate = int8(_eace)
	_ffdcg, _fcdd = _cbdf._effb.ReadBits(5)
	if _fcdd != nil {
		return _fcdd
	}
	_cbdf.SbDsOffset = int8(_ffdcg)
	if _cbdf.SbDsOffset > 0x0f {
		_cbdf.SbDsOffset -= 0x20
	}
	_eace, _fcdd = _cbdf._effb.ReadBit()
	if _fcdd != nil {
		return _fcdd
	}
	_cbdf.DefaultPixel = int8(_eace)
	_ffdcg, _fcdd = _cbdf._effb.ReadBits(2)
	if _fcdd != nil {
		return _fcdd
	}
	_cbdf.CombinationOperator = _d.CombinationOperator(int(_ffdcg) & 0x3)
	_eace, _fcdd = _cbdf._effb.ReadBit()
	if _fcdd != nil {
		return _fcdd
	}
	_cbdf.IsTransposed = int8(_eace)
	_ffdcg, _fcdd = _cbdf._effb.ReadBits(2)
	if _fcdd != nil {
		return _fcdd
	}
	_cbdf.ReferenceCorner = int16(_ffdcg) & 0x3
	_ffdcg, _fcdd = _cbdf._effb.ReadBits(2)
	if _fcdd != nil {
		return _fcdd
	}
	_cbdf.LogSBStrips = int16(_ffdcg) & 0x3
	_cbdf.SbStrips = 1 << uint(_cbdf.LogSBStrips)
	_eace, _fcdd = _cbdf._effb.ReadBit()
	if _fcdd != nil {
		return _fcdd
	}
	if _eace == 1 {
		_cbdf.UseRefinement = true
	}
	_eace, _fcdd = _cbdf._effb.ReadBit()
	if _fcdd != nil {
		return _fcdd
	}
	if _eace == 1 {
		_cbdf.IsHuffmanEncoded = true
	}
	return nil
}
func (_edbc *SymbolDictionary) encodeNumSyms(_acd _fg.BinaryWriter) (_cfgdd int, _cefa error) {
	const _fceg = "\u0065\u006e\u0063\u006f\u0064\u0065\u004e\u0075\u006d\u0053\u0079\u006d\u0073"
	_eeag := make([]byte, 4)
	_fe.BigEndian.PutUint32(_eeag, _edbc.NumberOfExportedSymbols)
	if _cfgdd, _cefa = _acd.Write(_eeag); _cefa != nil {
		return _cfgdd, _fd.Wrap(_cefa, _fceg, "\u0065\u0078p\u006f\u0072\u0074e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	_fe.BigEndian.PutUint32(_eeag, _edbc.NumberOfNewSymbols)
	_dcgf, _cefa := _acd.Write(_eeag)
	if _cefa != nil {
		return _cfgdd, _fd.Wrap(_cefa, _fceg, "n\u0065\u0077\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0073")
	}
	return _cfgdd + _dcgf, nil
}
func (_cdag *PageInformationSegment) readDefaultPixelValue() error {
	_gfgbf, _cbc := _cdag._ace.ReadBit()
	if _cbc != nil {
		return _cbc
	}
	_cdag.DefaultPixelValue = uint8(_gfgbf & 0xf)
	return nil
}
func (_bda *SymbolDictionary) getUserTable(_dbdf int) (_ba.Tabler, error) {
	var _abfg int
	for _, _afec := range _bda.Header.RTSegments {
		if _afec.Type == 53 {
			if _abfg == _dbdf {
				_fbab, _gdeg := _afec.GetSegmentData()
				if _gdeg != nil {
					return nil, _gdeg
				}
				_fcdce := _fbab.(_ba.BasicTabler)
				return _ba.NewEncodedTable(_fcdce)
			}
			_abfg++
		}
	}
	return nil, nil
}
func (_dbcb *SymbolDictionary) addSymbol(_gdad Regioner) error {
	_fcea, _bfdb := _gdad.GetRegionBitmap()
	if _bfdb != nil {
		return _bfdb
	}
	_dbcb._gbga[_dbcb._ffa] = _fcea
	_dbcb._bded = append(_dbcb._bded, _fcea)
	_da.Log.Trace("\u005b\u0053YM\u0042\u004f\u004c \u0044\u0049\u0043\u0054ION\u0041RY\u005d\u0020\u0041\u0064\u0064\u0065\u0064 s\u0079\u006d\u0062\u006f\u006c\u003a\u0020%\u0073", _fcea)
	return nil
}
func (_dcdg *Header) readHeaderFlags() error {
	const _fgef = "\u0072e\u0061d\u0048\u0065\u0061\u0064\u0065\u0072\u0046\u006c\u0061\u0067\u0073"
	_cfeb, _gdfa := _dcdg.Reader.ReadBit()
	if _gdfa != nil {
		return _fd.Wrap(_gdfa, _fgef, "r\u0065\u0074\u0061\u0069\u006e\u0020\u0066\u006c\u0061\u0067")
	}
	if _cfeb != 0 {
		_dcdg.RetainFlag = true
	}
	_cfeb, _gdfa = _dcdg.Reader.ReadBit()
	if _gdfa != nil {
		return _fd.Wrap(_gdfa, _fgef, "\u0070\u0061g\u0065\u0020\u0061s\u0073\u006f\u0063\u0069\u0061\u0074\u0069\u006f\u006e")
	}
	if _cfeb != 0 {
		_dcdg.PageAssociationFieldSize = true
	}
	_dbfc, _gdfa := _dcdg.Reader.ReadBits(6)
	if _gdfa != nil {
		return _fd.Wrap(_gdfa, _fgef, "\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065")
	}
	_dcdg.Type = Type(int(_dbfc))
	return nil
}
func (_fbfg *SymbolDictionary) Encode(w _fg.BinaryWriter) (_bab int, _eaag error) {
	const _bagae = "\u0053\u0079\u006dbo\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	if _fbfg == nil {
		return 0, _fd.Error(_bagae, "\u0073\u0079m\u0062\u006f\u006c\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	if _bab, _eaag = _fbfg.encodeFlags(w); _eaag != nil {
		return _bab, _fd.Wrap(_eaag, _bagae, "")
	}
	_gcad, _eaag := _fbfg.encodeATFlags(w)
	if _eaag != nil {
		return _bab, _fd.Wrap(_eaag, _bagae, "")
	}
	_bab += _gcad
	if _gcad, _eaag = _fbfg.encodeRefinementATFlags(w); _eaag != nil {
		return _bab, _fd.Wrap(_eaag, _bagae, "")
	}
	_bab += _gcad
	if _gcad, _eaag = _fbfg.encodeNumSyms(w); _eaag != nil {
		return _bab, _fd.Wrap(_eaag, _bagae, "")
	}
	_bab += _gcad
	if _gcad, _eaag = _fbfg.encodeSymbols(w); _eaag != nil {
		return _bab, _fd.Wrap(_eaag, _bagae, "")
	}
	_bab += _gcad
	return _bab, nil
}
func (_edf *GenericRegion) overrideAtTemplate1(_afgag, _faga, _dabc, _eddd, _bde int) int {
	_afgag &= 0x1FF7
	if _edf.GBAtY[0] == 0 && _edf.GBAtX[0] >= -int8(_bde) {
		_afgag |= (_eddd >> uint(7-(int8(_bde)+_edf.GBAtX[0])) & 0x1) << 3
	} else {
		_afgag |= int(_edf.getPixel(_faga+int(_edf.GBAtX[0]), _dabc+int(_edf.GBAtY[0]))) << 3
	}
	return _afgag
}
func (_aaca *SymbolDictionary) huffDecodeBmSize() (int64, error) {
	if _aaca._afea == nil {
		var (
			_feca int
			_faef error
		)
		if _aaca.SdHuffDecodeHeightSelection == 3 {
			_feca++
		}
		if _aaca.SdHuffDecodeWidthSelection == 3 {
			_feca++
		}
		_aaca._afea, _faef = _aaca.getUserTable(_feca)
		if _faef != nil {
			return 0, _faef
		}
	}
	return _aaca._afea.Decode(_aaca._bcagf)
}
func (_ebagf *TextRegion) symbolIDCodeLengths() error {
	var (
		_ggff []*_ba.Code
		_dbff uint64
		_bcgc _ba.Tabler
		_eafc error
	)
	for _bdcd := 0; _bdcd < 35; _bdcd++ {
		_dbff, _eafc = _ebagf._effb.ReadBits(4)
		if _eafc != nil {
			return _eafc
		}
		_egeg := int(_dbff & 0xf)
		if _egeg > 0 {
			_ggff = append(_ggff, _ba.NewCode(int32(_egeg), 0, int32(_bdcd), false))
		}
	}
	_bcgc, _eafc = _ba.NewFixedSizeTable(_ggff)
	if _eafc != nil {
		return _eafc
	}
	var (
		_bbaa  int64
		_acff  uint32
		_edbbd []*_ba.Code
		_degcc int64
	)
	for _acff < _ebagf.NumberOfSymbols {
		_degcc, _eafc = _bcgc.Decode(_ebagf._effb)
		if _eafc != nil {
			return _eafc
		}
		if _degcc < 32 {
			if _degcc > 0 {
				_edbbd = append(_edbbd, _ba.NewCode(int32(_degcc), 0, int32(_acff), false))
			}
			_bbaa = _degcc
			_acff++
		} else {
			var _eacf, _cgef int64
			switch _degcc {
			case 32:
				_dbff, _eafc = _ebagf._effb.ReadBits(2)
				if _eafc != nil {
					return _eafc
				}
				_eacf = 3 + int64(_dbff)
				if _acff > 0 {
					_cgef = _bbaa
				}
			case 33:
				_dbff, _eafc = _ebagf._effb.ReadBits(3)
				if _eafc != nil {
					return _eafc
				}
				_eacf = 3 + int64(_dbff)
			case 34:
				_dbff, _eafc = _ebagf._effb.ReadBits(7)
				if _eafc != nil {
					return _eafc
				}
				_eacf = 11 + int64(_dbff)
			}
			for _gbfe := 0; _gbfe < int(_eacf); _gbfe++ {
				if _cgef > 0 {
					_edbbd = append(_edbbd, _ba.NewCode(int32(_cgef), 0, int32(_acff), false))
				}
				_acff++
			}
		}
	}
	_ebagf._effb.Align()
	_ebagf._gegc, _eafc = _ba.NewFixedSizeTable(_edbbd)
	return _eafc
}
func (_eagg *TextRegion) decodeID() (int64, error) {
	if _eagg.IsHuffmanEncoded {
		if _eagg._gegc == nil {
			_gbbfd, _gcged := _eagg._effb.ReadBits(byte(_eagg._ebcb))
			return int64(_gbbfd), _gcged
		}
		return _eagg._gegc.Decode(_eagg._effb)
	}
	return _eagg._dffe.DecodeIAID(uint64(_eagg._ebcb), _eagg._beabb)
}

type EndOfStripe struct {
	_dg _fg.StreamReader
	_dc int
}

func (_baba *SymbolDictionary) setInSyms() error {
	if _baba.Header.RTSegments != nil {
		return _baba.retrieveImportSymbols()
	}
	_baba._cgfbb = make([]*_d.Bitmap, 0)
	return nil
}
func (_baa *GenericRefinementRegion) Init(header *Header, r _fg.StreamReader) error {
	_baa._cdg = header
	_baa._geg = r
	_baa.RegionInfo = NewRegionSegment(r)
	return _baa.parseHeader()
}
func (_ggedd *PatternDictionary) parseHeader() error {
	_da.Log.Trace("\u005b\u0050\u0041\u0054\u0054\u0045\u0052\u004e\u002d\u0044\u0049\u0043\u0054I\u004f\u004e\u0041\u0052\u0059\u005d[\u0070\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072\u005d\u0020b\u0065\u0067\u0069\u006e")
	defer func() {
		_da.Log.Trace("\u005b\u0050\u0041T\u0054\u0045\u0052\u004e\u002d\u0044\u0049\u0043\u0054\u0049\u004f\u004e\u0041\u0052\u0059\u005d\u005b\u0070\u0061\u0072\u0073\u0065\u0048\u0065\u0061\u0064\u0065\u0072\u005d \u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
	}()
	_, _fbfd := _ggedd._fbdd.ReadBits(5)
	if _fbfd != nil {
		return _fbfd
	}
	if _fbfd = _ggedd.readTemplate(); _fbfd != nil {
		return _fbfd
	}
	if _fbfd = _ggedd.readIsMMREncoded(); _fbfd != nil {
		return _fbfd
	}
	if _fbfd = _ggedd.readPatternWidthAndHeight(); _fbfd != nil {
		return _fbfd
	}
	if _fbfd = _ggedd.readGrayMax(); _fbfd != nil {
		return _fbfd
	}
	if _fbfd = _ggedd.computeSegmentDataStructure(); _fbfd != nil {
		return _fbfd
	}
	return _ggedd.checkInput()
}
func (_fegg *Header) subInputReader() (_fg.StreamReader, error) {
	return _fg.NewSubstreamReader(_fegg.Reader, _fegg.SegmentDataStartOffset, _fegg.SegmentDataLength)
}
func (_eef *GenericRegion) Encode(w _fg.BinaryWriter) (_cgag int, _fce error) {
	const _bdc = "G\u0065n\u0065\u0072\u0069\u0063\u0052\u0065\u0067\u0069o\u006e\u002e\u0045\u006eco\u0064\u0065"
	if _eef.Bitmap == nil {
		return 0, _fd.Error(_bdc, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	_cgf, _fce := _eef.RegionSegment.Encode(w)
	if _fce != nil {
		return 0, _fd.Wrap(_fce, _bdc, "\u0052\u0065\u0067\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_cgag += _cgf
	if _fce = w.SkipBits(4); _fce != nil {
		return _cgag, _fd.Wrap(_fce, _bdc, "\u0073k\u0069p\u0020\u0072\u0065\u0073\u0065r\u0076\u0065d\u0020\u0062\u0069\u0074\u0073")
	}
	var _abd int
	if _eef.IsTPGDon {
		_abd = 1
	}
	if _fce = w.WriteBit(_abd); _fce != nil {
		return _cgag, _fd.Wrap(_fce, _bdc, "\u0074\u0070\u0067\u0064\u006f\u006e")
	}
	_abd = 0
	if _fce = w.WriteBit(int(_eef.GBTemplate>>1) & 0x01); _fce != nil {
		return _cgag, _fd.Wrap(_fce, _bdc, "f\u0069r\u0073\u0074\u0020\u0067\u0062\u0074\u0065\u006dp\u006c\u0061\u0074\u0065 b\u0069\u0074")
	}
	if _fce = w.WriteBit(int(_eef.GBTemplate) & 0x01); _fce != nil {
		return _cgag, _fd.Wrap(_fce, _bdc, "s\u0065\u0063\u006f\u006ed \u0067b\u0074\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u0062\u0069\u0074")
	}
	if _eef.UseMMR {
		_abd = 1
	}
	if _fce = w.WriteBit(_abd); _fce != nil {
		return _cgag, _fd.Wrap(_fce, _bdc, "u\u0073\u0065\u0020\u004d\u004d\u0052\u0020\u0062\u0069\u0074")
	}
	_cgag++
	if _cgf, _fce = _eef.writeGBAtPixels(w); _fce != nil {
		return _cgag, _fd.Wrap(_fce, _bdc, "")
	}
	_cgag += _cgf
	_gbbd := _ga.New()
	if _fce = _gbbd.EncodeBitmap(_eef.Bitmap, _eef.IsTPGDon); _fce != nil {
		return _cgag, _fd.Wrap(_fce, _bdc, "")
	}
	_gbbd.Final()
	var _ecd int64
	if _ecd, _fce = _gbbd.WriteTo(w); _fce != nil {
		return _cgag, _fd.Wrap(_fce, _bdc, "")
	}
	_cgag += int(_ecd)
	return _cgag, nil
}
func (_gbb *GenericRefinementRegion) String() string {
	_dgfad := &_cc.Builder{}
	_dgfad.WriteString("\u000a[\u0047E\u004e\u0045\u0052\u0049\u0043 \u0052\u0045G\u0049\u004f\u004e\u005d\u000a")
	_dgfad.WriteString(_gbb.RegionInfo.String() + "\u000a")
	_dgfad.WriteString(_e.Sprintf("\u0009\u002d \u0049\u0073\u0054P\u0047\u0052\u006f\u006e\u003a\u0020\u0025\u0076\u000a", _gbb.IsTPGROn))
	_dgfad.WriteString(_e.Sprintf("\u0009-\u0020T\u0065\u006d\u0070\u006c\u0061t\u0065\u0049D\u003a\u0020\u0025\u0076\u000a", _gbb.TemplateID))
	_dgfad.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0047\u0072\u0041\u0074\u0058\u003a\u0020\u0025\u0076\u000a", _gbb.GrAtX))
	_dgfad.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0047\u0072\u0041\u0074\u0059\u003a\u0020\u0025\u0076\u000a", _gbb.GrAtY))
	_dgfad.WriteString(_e.Sprintf("\u0009-\u0020R\u0065\u0066\u0065\u0072\u0065n\u0063\u0065D\u0058\u0020\u0025\u0076\u000a", _gbb.ReferenceDX))
	_dgfad.WriteString(_e.Sprintf("\u0009\u002d\u0020\u0052ef\u0065\u0072\u0065\u006e\u0063\u0044\u0065\u0059\u003a\u0020\u0025\u0076\u000a", _gbb.ReferenceDY))
	return _dgfad.String()
}
func (_bed *GenericRefinementRegion) setParameters(_dgfa *_b.DecoderStats, _bb *_b.Decoder, _edd int8, _bae, _afe uint32, _gba *_d.Bitmap, _aaa, _daf int32, _caa bool, _cdd []int8, _bagd []int8) {
	_da.Log.Trace("\u005b\u0047\u0045NE\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052E\u0047I\u004fN\u005d \u0073\u0065\u0074\u0050\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073")
	if _dgfa != nil {
		_bed._cdb = _dgfa
	}
	if _bb != nil {
		_bed._gf = _bb
	}
	_bed.TemplateID = _edd
	_bed.RegionInfo.BitmapWidth = _bae
	_bed.RegionInfo.BitmapHeight = _afe
	_bed.ReferenceBitmap = _gba
	_bed.ReferenceDX = _aaa
	_bed.ReferenceDY = _daf
	_bed.IsTPGROn = _caa
	_bed.GrAtX = _cdd
	_bed.GrAtY = _bagd
	_bed.RegionBitmap = nil
	_da.Log.Trace("[\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052E\u0046\u002d\u0052\u0045\u0047\u0049\u004fN]\u0020\u0073\u0065\u0074P\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073 f\u0069\u006ei\u0073\u0068\u0065\u0064\u002e\u0020\u0025\u0073", _bed)
}
func (_ccb *SymbolDictionary) encodeRefinementATFlags(_cfac _fg.BinaryWriter) (_eca int, _gbgc error) {
	const _beb = "\u0065\u006e\u0063od\u0065\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0041\u0054\u0046\u006c\u0061\u0067\u0073"
	if !_ccb.UseRefinementAggregation || _ccb.SdrTemplate != 0 {
		return 0, nil
	}
	for _acea := 0; _acea < 2; _acea++ {
		if _gbgc = _cfac.WriteByte(byte(_ccb.SdrATX[_acea])); _gbgc != nil {
			return _eca, _fd.Wrapf(_gbgc, _beb, "\u0053\u0064\u0072\u0041\u0054\u0058\u005b\u0025\u0064\u005d", _acea)
		}
		_eca++
		if _gbgc = _cfac.WriteByte(byte(_ccb.SdrATY[_acea])); _gbgc != nil {
			return _eca, _fd.Wrapf(_gbgc, _beb, "\u0053\u0064\u0072\u0041\u0054\u0059\u005b\u0025\u0064\u005d", _acea)
		}
		_eca++
	}
	return _eca, nil
}
func (_aebga *TextRegion) InitEncode(globalSymbolsMap, localSymbolsMap map[int]int, comps []int, inLL *_d.Points, symbols *_d.Bitmaps, classIDs *_ca.IntSlice, boxes *_d.Boxes, width, height, symBits int) {
	_aebga.RegionInfo = &RegionSegment{BitmapWidth: uint32(width), BitmapHeight: uint32(height)}
	_aebga._fegdc = globalSymbolsMap
	_aebga._cfaa = localSymbolsMap
	_aebga._ddfbb = comps
	_aebga._fbc = inLL
	_aebga._aeee = symbols
	_aebga._faed = classIDs
	_aebga._ffbc = boxes
	_aebga._aagc = symBits
}
func (_daeaf *TextRegion) initSymbols() error {
	const _faeb = "i\u006e\u0069\u0074\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	for _, _befd := range _daeaf.Header.RTSegments {
		if _befd == nil {
			return _fd.Error(_faeb, "\u006e\u0069\u006c\u0020\u0073\u0065\u0067\u006de\u006e\u0074\u0020pr\u006f\u0076\u0069\u0064\u0065\u0064 \u0066\u006f\u0072\u0020\u0074\u0068\u0065\u0020\u0074\u0065\u0078\u0074\u0020\u0072\u0065g\u0069\u006f\u006e\u0020\u0053\u0079\u006d\u0062o\u006c\u0073")
		}
		if _befd.Type == 0 {
			_ebfe, _edae := _befd.GetSegmentData()
			if _edae != nil {
				return _fd.Wrap(_edae, _faeb, "")
			}
			_cgcc, _cbfd := _ebfe.(*SymbolDictionary)
			if !_cbfd {
				return _fd.Error(_faeb, "\u0072e\u0066\u0065r\u0072\u0065\u0064 \u0054\u006f\u0020\u0053\u0065\u0067\u006de\u006e\u0074\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0061\u0020\u0053\u0079\u006d\u0062\u006f\u006cD\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079")
			}
			_cgcc._fac = _daeaf._beabb
			_eefe, _edae := _cgcc.GetDictionary()
			if _edae != nil {
				return _fd.Wrap(_edae, _faeb, "")
			}
			_daeaf.Symbols = append(_daeaf.Symbols, _eefe...)
		}
	}
	_daeaf.NumberOfSymbols = uint32(len(_daeaf.Symbols))
	return nil
}
func (_cfgd *PatternDictionary) extractPatterns(_dgeg *_d.Bitmap) error {
	var _cab int
	_afcb := make([]*_d.Bitmap, _cfgd.GrayMax+1)
	for _cab <= int(_cfgd.GrayMax) {
		_egcg := int(_cfgd.HdpWidth) * _cab
		_afee := _c.Rect(_egcg, 0, _egcg+int(_cfgd.HdpWidth), int(_cfgd.HdpHeight))
		_adad, _dgbd := _d.Extract(_afee, _dgeg)
		if _dgbd != nil {
			return _dgbd
		}
		_afcb[_cab] = _adad
		_cab++
	}
	_cfgd.Patterns = _afcb
	return nil
}
func (_bdbg *TableSegment) Init(h *Header, r _fg.StreamReader) error {
	_bdbg._defgf = r
	return _bdbg.parseHeader()
}
func (_aecd *PageInformationSegment) readMaxStripeSize() error {
	_dcce, _ebfd := _aecd._ace.ReadBits(15)
	if _ebfd != nil {
		return _ebfd
	}
	_aecd.MaxStripeSize = uint16(_dcce & _f.MaxUint16)
	return nil
}

type SegmentEncoder interface {
	Encode(_decc _fg.BinaryWriter) (_fcbb int, _bdg error)
}

func (_bdec *PageInformationSegment) checkInput() error {
	if _bdec.PageBMHeight == _f.MaxInt32 {
		if !_bdec.IsStripe {
			_da.Log.Debug("P\u0061\u0067\u0065\u0049\u006e\u0066\u006f\u0072\u006da\u0074\u0069\u006f\u006e\u0053\u0065\u0067me\u006e\u0074\u002e\u0049s\u0053\u0074\u0072\u0069\u0070\u0065\u0020\u0073\u0068ou\u006c\u0064 \u0062\u0065\u0020\u0074\u0072\u0075\u0065\u002e")
		}
	}
	return nil
}
func (_fdb *Header) parse(_ffd Documenter, _edgb _fg.StreamReader, _cebe int64, _gcgf OrganizationType) (_cade error) {
	const _gede = "\u0070\u0061\u0072s\u0065"
	_da.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u002d\u0048E\u0041\u0044\u0045\u0052\u005d\u005b\u0050A\u0052\u0053\u0045\u005d\u0020\u0042\u0065\u0067\u0069\u006e\u0073")
	defer func() {
		if _cade != nil {
			_da.Log.Trace("\u005b\u0053\u0045GM\u0045\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045R\u005d[\u0050A\u0052S\u0045\u005d\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _cade)
		} else {
			_da.Log.Trace("\u005b\u0053\u0045\u0047\u004d\u0045\u004e\u0054\u002d\u0048\u0045\u0041\u0044\u0045\u0052]\u005bP\u0041\u0052\u0053\u0045\u005d\u0020\u0046\u0069\u006e\u0069\u0073\u0068\u0065\u0064")
		}
	}()
	_, _cade = _edgb.Seek(_cebe, _gc.SeekStart)
	if _cade != nil {
		return _fd.Wrap(_cade, _gede, "\u0073\u0065\u0065\u006b\u0020\u0073\u0074\u0061\u0072\u0074")
	}
	if _cade = _fdb.readSegmentNumber(_edgb); _cade != nil {
		return _fd.Wrap(_cade, _gede, "")
	}
	if _cade = _fdb.readHeaderFlags(); _cade != nil {
		return _fd.Wrap(_cade, _gede, "")
	}
	var _dbd uint64
	_dbd, _cade = _fdb.readNumberOfReferredToSegments(_edgb)
	if _cade != nil {
		return _fd.Wrap(_cade, _gede, "")
	}
	_fdb.RTSNumbers, _cade = _fdb.readReferredToSegmentNumbers(_edgb, int(_dbd))
	if _cade != nil {
		return _fd.Wrap(_cade, _gede, "")
	}
	_cade = _fdb.readSegmentPageAssociation(_ffd, _edgb, _dbd, _fdb.RTSNumbers...)
	if _cade != nil {
		return _fd.Wrap(_cade, _gede, "")
	}
	if _fdb.Type != TEndOfFile {
		if _cade = _fdb.readSegmentDataLength(_edgb); _cade != nil {
			return _fd.Wrap(_cade, _gede, "")
		}
	}
	_fdb.readDataStartOffset(_edgb, _gcgf)
	_fdb.readHeaderLength(_edgb, _cebe)
	_da.Log.Trace("\u0025\u0073", _fdb)
	return nil
}
func (_adde *TableSegment) HtRS() int32 { return _adde._cgeg }
func (_eff *GenericRegion) InitEncode(bm *_d.Bitmap, xLoc, yLoc, template int, duplicateLineRemoval bool) error {
	const _eee = "\u0047e\u006e\u0065\u0072\u0069\u0063\u0052\u0065\u0067\u0069\u006f\u006e.\u0049\u006e\u0069\u0074\u0045\u006e\u0063\u006f\u0064\u0065"
	if bm == nil {
		return _fd.Error(_eee, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if xLoc < 0 || yLoc < 0 {
		return _fd.Error(_eee, "\u0078\u0020\u0061\u006e\u0064\u0020\u0079\u0020\u006c\u006f\u0063\u0061\u0074i\u006f\u006e\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0030")
	}
	_eff.Bitmap = bm
	_eff.GBTemplate = byte(template)
	switch _eff.GBTemplate {
	case 0:
		_eff.GBAtX = []int8{3, -3, 2, -2}
		_eff.GBAtY = []int8{-1, -1, -2, -2}
	case 1:
		_eff.GBAtX = []int8{3}
		_eff.GBAtY = []int8{-1}
	case 2, 3:
		_eff.GBAtX = []int8{2}
		_eff.GBAtY = []int8{-1}
	default:
		return _fd.Errorf(_eee, "\u0070\u0072o\u0076\u0069\u0064\u0065\u0064 \u0074\u0065\u006d\u0070\u006ca\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0030\u002c\u0031\u002c\u0032\u002c\u0033\u007d", template)
	}
	_eff.RegionSegment = &RegionSegment{BitmapHeight: uint32(bm.Height), BitmapWidth: uint32(bm.Width), XLocation: uint32(xLoc), YLocation: uint32(yLoc)}
	_eff.IsTPGDon = duplicateLineRemoval
	return nil
}
func (_gaffc *HalftoneRegion) shiftAndFill(_cef int) int {
	_cef >>= 8
	if _cef < 0 {
		_ggea := int(_f.Log(float64(_dfe(_cef))) / _f.Log(2))
		_beaf := 31 - _ggea
		for _gcaf := 1; _gcaf < _beaf; _gcaf++ {
			_cef |= 1 << uint(31-_gcaf)
		}
	}
	return _cef
}

type GenericRegion struct {
	_baf             _fg.StreamReader
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
	_fgf             bool
	Bitmap           *_d.Bitmap
	_acb             *_b.Decoder
	_aae             *_b.DecoderStats
	_gcd             *_ad.Decoder
}

func (_fcfed *TableSegment) HtHigh() int32 { return _fcfed._efgd }
func (_gcf *GenericRefinementRegion) parseHeader() (_aaf error) {
	_da.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072s\u0069\u006e\u0067\u0020\u0048e\u0061\u0064e\u0072\u002e\u002e\u002e")
	_dcb := _a.Now()
	defer func() {
		if _aaf == nil {
			_da.Log.Trace("\u005b\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045G\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020h\u0065\u0061\u0064\u0065\u0072\u0020\u0066\u0069\u006e\u0069\u0073\u0068id\u0020\u0069\u006e\u003a\u0020\u0025\u0064\u0020\u006e\u0073", _a.Since(_dcb).Nanoseconds())
		} else {
			_da.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0073", _aaf)
		}
	}()
	if _aaf = _gcf.RegionInfo.parseHeader(); _aaf != nil {
		return _aaf
	}
	_, _aaf = _gcf._geg.ReadBits(6)
	if _aaf != nil {
		return _aaf
	}
	_gcf.IsTPGROn, _aaf = _gcf._geg.ReadBool()
	if _aaf != nil {
		return _aaf
	}
	var _daa int
	_daa, _aaf = _gcf._geg.ReadBit()
	if _aaf != nil {
		return _aaf
	}
	_gcf.TemplateID = int8(_daa)
	switch _gcf.TemplateID {
	case 0:
		_gcf.Template = _gcf._dae
		if _aaf = _gcf.readAtPixels(); _aaf != nil {
			return
		}
	case 1:
		_gcf.Template = _gcf._dec
	}
	return nil
}
func (_caab *PageInformationSegment) readCombinationOperatorOverrideAllowed() error {
	_efd, _dgfg := _caab._ace.ReadBit()
	if _dgfg != nil {
		return _dgfg
	}
	if _efd == 1 {
		_caab._eae = true
	}
	return nil
}
func (_bgecd *SymbolDictionary) readAtPixels(_afce int) error {
	_bgecd.SdATX = make([]int8, _afce)
	_bgecd.SdATY = make([]int8, _afce)
	var (
		_ccdg byte
		_fbda error
	)
	for _edbcb := 0; _edbcb < _afce; _edbcb++ {
		_ccdg, _fbda = _bgecd._bcagf.ReadByte()
		if _fbda != nil {
			return _fbda
		}
		_bgecd.SdATX[_edbcb] = int8(_ccdg)
		_ccdg, _fbda = _bgecd._bcagf.ReadByte()
		if _fbda != nil {
			return _fbda
		}
		_bgecd.SdATY[_edbcb] = int8(_ccdg)
	}
	return nil
}

type Type int

func (_dfda *PageInformationSegment) CombinationOperator() _d.CombinationOperator {
	return _dfda._aefgc
}
func (_bdcef *GenericRegion) getPixel(_fffe, _caae int) int8 {
	if _fffe < 0 || _fffe >= _bdcef.Bitmap.Width {
		return 0
	}
	if _caae < 0 || _caae >= _bdcef.Bitmap.Height {
		return 0
	}
	if _bdcef.Bitmap.GetPixel(_fffe, _caae) {
		return 1
	}
	return 0
}
func (_ffee *Header) readReferredToSegmentNumbers(_cacf _fg.StreamReader, _eeab int) ([]int, error) {
	const _dcdc = "\u0072\u0065\u0061\u0064R\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0054\u006f\u0053e\u0067m\u0065\u006e\u0074\u004e\u0075\u006d\u0062e\u0072\u0073"
	_dgae := make([]int, _eeab)
	if _eeab > 0 {
		_ffee.RTSegments = make([]*Header, _eeab)
		var (
			_dbbf uint64
			_cdfg error
		)
		for _geaf := 0; _geaf < _eeab; _geaf++ {
			_dbbf, _cdfg = _cacf.ReadBits(byte(_ffee.referenceSize()) << 3)
			if _cdfg != nil {
				return nil, _fd.Wrapf(_cdfg, _dcdc, "\u0027\u0025\u0064\u0027 \u0072\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0020\u0073e\u0067m\u0065\u006e\u0074\u0020\u006e\u0075\u006db\u0065\u0072", _geaf)
			}
			_dgae[_geaf] = int(_dbbf & _f.MaxInt32)
		}
	}
	return _dgae, nil
}
func (_abdd *SymbolDictionary) InitEncode(symbols *_d.Bitmaps, symbolList []int, symbolMap map[int]int, unborderSymbols bool) error {
	const _bcefg = "S\u0079\u006d\u0062\u006f\u006c\u0044i\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079\u002eI\u006e\u0069\u0074E\u006ec\u006f\u0064\u0065"
	_abdd.SdATX = []int8{3, -3, 2, -2}
	_abdd.SdATY = []int8{-1, -1, -2, -2}
	_abdd._dbag = symbols
	_abdd._fagdd = make([]int, len(symbolList))
	copy(_abdd._fagdd, symbolList)
	if len(_abdd._fagdd) != _abdd._dbag.Size() {
		return _fd.Error(_bcefg, "s\u0079\u006d\u0062\u006f\u006c\u0073\u0020\u0061\u006e\u0064\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u004ci\u0073\u0074\u0020\u006f\u0066\u0020\u0064\u0069\u0066\u0066er\u0065\u006e\u0074 \u0073i\u007a\u0065")
	}
	_abdd.NumberOfNewSymbols = uint32(symbols.Size())
	_abdd.NumberOfExportedSymbols = uint32(symbols.Size())
	_abdd._dcbf = symbolMap
	_abdd._cgbf = unborderSymbols
	return nil
}
func (_bdfb *TextRegion) decodeSymbolInstances() error {
	_cfgb, _daea := _bdfb.decodeStripT()
	if _daea != nil {
		return _daea
	}
	var (
		_gfab int64
		_adfb uint32
	)
	for _adfb < _bdfb.NumberOfSymbolInstances {
		_cffe, _gebb := _bdfb.decodeDT()
		if _gebb != nil {
			return _gebb
		}
		_cfgb += _cffe
		var _cdcb int64
		_egcd := true
		_bdfb._fbdbc = 0
		for {
			if _egcd {
				_cdcb, _gebb = _bdfb.decodeDfs()
				if _gebb != nil {
					return _gebb
				}
				_gfab += _cdcb
				_bdfb._fbdbc = _gfab
				_egcd = false
			} else {
				_ddea, _dadca := _bdfb.decodeIds()
				if _cd.Is(_dadca, _gd.ErrOOB) {
					break
				}
				if _dadca != nil {
					return _dadca
				}
				if _adfb >= _bdfb.NumberOfSymbolInstances {
					break
				}
				_bdfb._fbdbc += _ddea + int64(_bdfb.SbDsOffset)
			}
			_dggg, _agecd := _bdfb.decodeCurrentT()
			if _agecd != nil {
				return _agecd
			}
			_gecg := _cfgb + _dggg
			_fged, _agecd := _bdfb.decodeID()
			if _agecd != nil {
				return _agecd
			}
			_fcba, _agecd := _bdfb.decodeRI()
			if _agecd != nil {
				return _agecd
			}
			_dffd, _agecd := _bdfb.decodeIb(_fcba, _fged)
			if _agecd != nil {
				return _agecd
			}
			if _agecd = _bdfb.blit(_dffd, _gecg); _agecd != nil {
				return _agecd
			}
			_adfb++
		}
	}
	return nil
}
func (_acae *TextRegion) Encode(w _fg.BinaryWriter) (_dagc int, _ebgc error) {
	const _efef = "\u0054\u0065\u0078\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u002e\u0045n\u0063\u006f\u0064\u0065"
	if _dagc, _ebgc = _acae.RegionInfo.Encode(w); _ebgc != nil {
		return _dagc, _fd.Wrap(_ebgc, _efef, "")
	}
	var _aefa int
	if _aefa, _ebgc = _acae.encodeFlags(w); _ebgc != nil {
		return _dagc, _fd.Wrap(_ebgc, _efef, "")
	}
	_dagc += _aefa
	if _aefa, _ebgc = _acae.encodeSymbols(w); _ebgc != nil {
		return _dagc, _fd.Wrap(_ebgc, _efef, "")
	}
	_dagc += _aefa
	return _dagc, nil
}
func (_gged *GenericRegion) setParameters(_cdfb bool, _ada, _badg int64, _gffa, _ffc uint32) {
	_gged.IsMMREncoded = _cdfb
	_gged.DataOffset = _ada
	_gged.DataLength = _badg
	_gged.RegionSegment.BitmapHeight = _gffa
	_gged.RegionSegment.BitmapWidth = _ffc
	_gged._gcd = nil
	_gged.Bitmap = nil
}
func (_gafa *PatternDictionary) checkInput() error {
	if _gafa.HdpHeight < 1 || _gafa.HdpWidth < 1 {
		return _ge.New("in\u0076\u0061l\u0069\u0064\u0020\u0048\u0065\u0061\u0064\u0065\u0072 \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0057\u0069\u0064\u0074\u0068\u002f\u0048\u0065\u0069\u0067\u0068\u0074\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020g\u0072e\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020z\u0065\u0072o")
	}
	if _gafa.IsMMREncoded {
		if _gafa.HDTemplate != 0 {
			_da.Log.Debug("\u0076\u0061\u0072\u0069\u0061\u0062\u006c\u0065\u0020\u0048\u0044\u0054\u0065\u006d\u0070\u006c\u0061\u0074e\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u006e\u006f\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e \u0074\u0068\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0030")
		}
	}
	return nil
}
func (_ffcb *TextRegion) readHuffmanFlags() error {
	var (
		_ecde int
		_fecc uint64
		_bee  error
	)
	_, _bee = _ffcb._effb.ReadBit()
	if _bee != nil {
		return _bee
	}
	_ecde, _bee = _ffcb._effb.ReadBit()
	if _bee != nil {
		return _bee
	}
	_ffcb.SbHuffRSize = int8(_ecde)
	_fecc, _bee = _ffcb._effb.ReadBits(2)
	if _bee != nil {
		return _bee
	}
	_ffcb.SbHuffRDY = int8(_fecc) & 0xf
	_fecc, _bee = _ffcb._effb.ReadBits(2)
	if _bee != nil {
		return _bee
	}
	_ffcb.SbHuffRDX = int8(_fecc) & 0xf
	_fecc, _bee = _ffcb._effb.ReadBits(2)
	if _bee != nil {
		return _bee
	}
	_ffcb.SbHuffRDHeight = int8(_fecc) & 0xf
	_fecc, _bee = _ffcb._effb.ReadBits(2)
	if _bee != nil {
		return _bee
	}
	_ffcb.SbHuffRDWidth = int8(_fecc) & 0xf
	_fecc, _bee = _ffcb._effb.ReadBits(2)
	if _bee != nil {
		return _bee
	}
	_ffcb.SbHuffDT = int8(_fecc) & 0xf
	_fecc, _bee = _ffcb._effb.ReadBits(2)
	if _bee != nil {
		return _bee
	}
	_ffcb.SbHuffDS = int8(_fecc) & 0xf
	_fecc, _bee = _ffcb._effb.ReadBits(2)
	if _bee != nil {
		return _bee
	}
	_ffcb.SbHuffFS = int8(_fecc) & 0xf
	return nil
}

type PatternDictionary struct {
	_fbdd            _fg.StreamReader
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
	Patterns         []*_d.Bitmap
	GrayMax          uint32
}

func (_daff *PatternDictionary) readTemplate() error {
	_dbfg, _dgbb := _daff._fbdd.ReadBits(2)
	if _dgbb != nil {
		return _dgbb
	}
	_daff.HDTemplate = byte(_dbfg)
	return nil
}
func (_gaec *PatternDictionary) readGrayMax() error {
	_aced, _abde := _gaec._fbdd.ReadBits(32)
	if _abde != nil {
		return _abde
	}
	_gaec.GrayMax = uint32(_aced & _f.MaxUint32)
	return nil
}
func (_gbgb *PatternDictionary) readPatternWidthAndHeight() error {
	_fbac, _daddf := _gbgb._fbdd.ReadByte()
	if _daddf != nil {
		return _daddf
	}
	_gbgb.HdpWidth = _fbac
	_fbac, _daddf = _gbgb._fbdd.ReadByte()
	if _daddf != nil {
		return _daddf
	}
	_gbgb.HdpHeight = _fbac
	return nil
}
func (_gbed *TableSegment) StreamReader() _fg.StreamReader { return _gbed._defgf }
func (_gdcd *TextRegion) checkInput() error {
	const _bbec = "\u0063\u0068\u0065\u0063\u006b\u0049\u006e\u0070\u0075\u0074"
	if !_gdcd.UseRefinement {
		if _gdcd.SbrTemplate != 0 {
			_da.Log.Debug("\u0053\u0062\u0072Te\u006d\u0070\u006c\u0061\u0074\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030")
			_gdcd.SbrTemplate = 0
		}
	}
	if _gdcd.SbHuffFS == 2 || _gdcd.SbHuffRDWidth == 2 || _gdcd.SbHuffRDHeight == 2 || _gdcd.SbHuffRDX == 2 || _gdcd.SbHuffRDY == 2 {
		return _fd.Error(_bbec, "h\u0075\u0066\u0066\u006d\u0061\u006e \u0066\u006c\u0061\u0067\u0020\u0076a\u006c\u0075\u0065\u0020\u0069\u0073\u0020n\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074e\u0064")
	}
	if !_gdcd.UseRefinement {
		if _gdcd.SbHuffRSize != 0 {
			_da.Log.Debug("\u0053\u0062\u0048uf\u0066\u0052\u0053\u0069\u007a\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0030")
			_gdcd.SbHuffRSize = 0
		}
		if _gdcd.SbHuffRDY != 0 {
			_da.Log.Debug("S\u0062\u0048\u0075\u0066fR\u0044Y\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0030")
			_gdcd.SbHuffRDY = 0
		}
		if _gdcd.SbHuffRDX != 0 {
			_da.Log.Debug("S\u0062\u0048\u0075\u0066fR\u0044X\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0030")
			_gdcd.SbHuffRDX = 0
		}
		if _gdcd.SbHuffRDWidth != 0 {
			_da.Log.Debug("\u0053b\u0048\u0075\u0066\u0066R\u0044\u0057\u0069\u0064\u0074h\u0020s\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u00200")
			_gdcd.SbHuffRDWidth = 0
		}
		if _gdcd.SbHuffRDHeight != 0 {
			_da.Log.Debug("\u0053\u0062\u0048\u0075\u0066\u0066\u0052\u0044\u0048\u0065\u0069g\u0068\u0074\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020b\u0065\u0020\u0030")
			_gdcd.SbHuffRDHeight = 0
		}
	}
	return nil
}
func (_gafe *Header) writeReferredToSegments(_gagd _fg.BinaryWriter) (_cdegd int, _feec error) {
	const _abbb = "\u0077\u0072\u0069te\u0052\u0065\u0066\u0065\u0072\u0072\u0065\u0064\u0054\u006f\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0073"
	var (
		_egf  uint16
		_dfdc uint32
	)
	_cbg := _gafe.referenceSize()
	_eaae := 1
	_bgg := make([]byte, _cbg)
	for _, _gcgff := range _gafe.RTSNumbers {
		switch _cbg {
		case 4:
			_dfdc = uint32(_gcgff)
			_fe.BigEndian.PutUint32(_bgg, _dfdc)
			_eaae, _feec = _gagd.Write(_bgg)
			if _feec != nil {
				return 0, _fd.Wrap(_feec, _abbb, "u\u0069\u006e\u0074\u0033\u0032\u0020\u0073\u0069\u007a\u0065")
			}
		case 2:
			_egf = uint16(_gcgff)
			_fe.BigEndian.PutUint16(_bgg, _egf)
			_eaae, _feec = _gagd.Write(_bgg)
			if _feec != nil {
				return 0, _fd.Wrap(_feec, _abbb, "\u0075\u0069\u006e\u0074\u0031\u0036")
			}
		default:
			if _feec = _gagd.WriteByte(byte(_gcgff)); _feec != nil {
				return 0, _fd.Wrap(_feec, _abbb, "\u0075\u0069\u006et\u0038")
			}
		}
		_cdegd += _eaae
	}
	return _cdegd, nil
}
func (_cae *Header) readSegmentNumber(_bbf _fg.StreamReader) error {
	const _ebcg = "\u0072\u0065\u0061\u0064\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u004eu\u006d\u0062\u0065\u0072"
	_bbbb := make([]byte, 4)
	_, _gddf := _bbf.Read(_bbbb)
	if _gddf != nil {
		return _fd.Wrap(_gddf, _ebcg, "")
	}
	_cae.SegmentNumber = _fe.BigEndian.Uint32(_bbbb)
	return nil
}
func (_cafb *Header) writeSegmentNumber(_agcc _fg.BinaryWriter) (_bcda int, _eccc error) {
	_ggeb := make([]byte, 4)
	_fe.BigEndian.PutUint32(_ggeb, _cafb.SegmentNumber)
	if _bcda, _eccc = _agcc.Write(_ggeb); _eccc != nil {
		return 0, _fd.Wrap(_eccc, "\u0048e\u0061\u0064\u0065\u0072.\u0077\u0072\u0069\u0074\u0065S\u0065g\u006de\u006e\u0074\u004e\u0075\u006d\u0062\u0065r", "")
	}
	return _bcda, nil
}
func (_fbddg *TextRegion) decodeIds() (int64, error) {
	const _defa = "\u0064e\u0063\u006f\u0064\u0065\u0049\u0064s"
	if _fbddg.IsHuffmanEncoded {
		if _fbddg.SbHuffDS == 3 {
			if _fbddg._cadeg == nil {
				_cagc := 0
				if _fbddg.SbHuffFS == 3 {
					_cagc++
				}
				var _ccaf error
				_fbddg._cadeg, _ccaf = _fbddg.getUserTable(_cagc)
				if _ccaf != nil {
					return 0, _fd.Wrap(_ccaf, _defa, "")
				}
			}
			return _fbddg._cadeg.Decode(_fbddg._effb)
		}
		_ccagd, _bcgfc := _ba.GetStandardTable(8 + int(_fbddg.SbHuffDS))
		if _bcgfc != nil {
			return 0, _fd.Wrap(_bcgfc, _defa, "")
		}
		return _ccagd.Decode(_fbddg._effb)
	}
	_fafcb, _degg := _fbddg._dffe.DecodeInt(_fbddg._dafb)
	if _degg != nil {
		return 0, _fd.Wrap(_degg, _defa, "\u0063\u0078\u0049\u0041\u0044\u0053")
	}
	return int64(_fafcb), nil
}
func (_ddd *GenericRefinementRegion) decodeSLTP() (int, error) {
	_ddd.Template.setIndex(_ddd._cdb)
	return _ddd._gf.DecodeBit(_ddd._cdb)
}
func (_ccgb *RegionSegment) Encode(w _fg.BinaryWriter) (_fddf int, _ffcd error) {
	const _ebdg = "R\u0065g\u0069\u006f\u006e\u0053\u0065\u0067\u006d\u0065n\u0074\u002e\u0045\u006eco\u0064\u0065"
	_cdegb := make([]byte, 4)
	_fe.BigEndian.PutUint32(_cdegb, _ccgb.BitmapWidth)
	_fddf, _ffcd = w.Write(_cdegb)
	if _ffcd != nil {
		return 0, _fd.Wrap(_ffcd, _ebdg, "\u0057\u0069\u0064t\u0068")
	}
	_fe.BigEndian.PutUint32(_cdegb, _ccgb.BitmapHeight)
	var _fgge int
	_fgge, _ffcd = w.Write(_cdegb)
	if _ffcd != nil {
		return 0, _fd.Wrap(_ffcd, _ebdg, "\u0048\u0065\u0069\u0067\u0068\u0074")
	}
	_fddf += _fgge
	_fe.BigEndian.PutUint32(_cdegb, _ccgb.XLocation)
	_fgge, _ffcd = w.Write(_cdegb)
	if _ffcd != nil {
		return 0, _fd.Wrap(_ffcd, _ebdg, "\u0058L\u006f\u0063\u0061\u0074\u0069\u006fn")
	}
	_fddf += _fgge
	_fe.BigEndian.PutUint32(_cdegb, _ccgb.YLocation)
	_fgge, _ffcd = w.Write(_cdegb)
	if _ffcd != nil {
		return 0, _fd.Wrap(_ffcd, _ebdg, "\u0059L\u006f\u0063\u0061\u0074\u0069\u006fn")
	}
	_fddf += _fgge
	if _ffcd = w.WriteByte(byte(_ccgb.CombinaionOperator) & 0x07); _ffcd != nil {
		return 0, _fd.Wrap(_ffcd, _ebdg, "c\u006fm\u0062\u0069\u006e\u0061\u0074\u0069\u006f\u006e \u006f\u0070\u0065\u0072at\u006f\u0072")
	}
	_fddf++
	return _fddf, nil
}
func (_ffgf *RegionSegment) Size() int { return 17 }

type Documenter interface {
	GetPage(int) (Pager, error)
	GetGlobalSegment(int) (*Header, error)
}

func (_daddc *TextRegion) readUseRefinement() error {
	if !_daddc.UseRefinement || _daddc.SbrTemplate != 0 {
		return nil
	}
	var (
		_cgcba byte
		_geea  error
	)
	_daddc.SbrATX = make([]int8, 2)
	_daddc.SbrATY = make([]int8, 2)
	_cgcba, _geea = _daddc._effb.ReadByte()
	if _geea != nil {
		return _geea
	}
	_daddc.SbrATX[0] = int8(_cgcba)
	_cgcba, _geea = _daddc._effb.ReadByte()
	if _geea != nil {
		return _geea
	}
	_daddc.SbrATY[0] = int8(_cgcba)
	_cgcba, _geea = _daddc._effb.ReadByte()
	if _geea != nil {
		return _geea
	}
	_daddc.SbrATX[1] = int8(_cgcba)
	_cgcba, _geea = _daddc._effb.ReadByte()
	if _geea != nil {
		return _geea
	}
	_daddc.SbrATY[1] = int8(_cgcba)
	return nil
}
func (_eadg *PatternDictionary) setGbAtPixels() {
	if _eadg.HDTemplate == 0 {
		_eadg.GBAtX = make([]int8, 4)
		_eadg.GBAtY = make([]int8, 4)
		_eadg.GBAtX[0] = -int8(_eadg.HdpWidth)
		_eadg.GBAtY[0] = 0
		_eadg.GBAtX[1] = -3
		_eadg.GBAtY[1] = -1
		_eadg.GBAtX[2] = 2
		_eadg.GBAtY[2] = -2
		_eadg.GBAtX[3] = -2
		_eadg.GBAtY[3] = -2
	} else {
		_eadg.GBAtX = []int8{-int8(_eadg.HdpWidth)}
		_eadg.GBAtY = []int8{0}
	}
}
func (_gag *GenericRegion) decodeLine(_abe, _fafd, _bfc int) error {
	const _ecdg = "\u0064\u0065\u0063\u006f\u0064\u0065\u004c\u0069\u006e\u0065"
	_ggga := _gag.Bitmap.GetByteIndex(0, _abe)
	_ebg := _ggga - _gag.Bitmap.RowStride
	switch _gag.GBTemplate {
	case 0:
		if !_gag.UseExtTemplates {
			return _gag.decodeTemplate0a(_abe, _fafd, _bfc, _ggga, _ebg)
		}
		return _gag.decodeTemplate0b(_abe, _fafd, _bfc, _ggga, _ebg)
	case 1:
		return _gag.decodeTemplate1(_abe, _fafd, _bfc, _ggga, _ebg)
	case 2:
		return _gag.decodeTemplate2(_abe, _fafd, _bfc, _ggga, _ebg)
	case 3:
		return _gag.decodeTemplate3(_abe, _fafd, _bfc, _ggga, _ebg)
	}
	return _fd.Errorf(_ecdg, "\u0069\u006e\u0076a\u006c\u0069\u0064\u0020G\u0042\u0054\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u003a\u0020\u0025\u0064", _gag.GBTemplate)
}
func (_de *EndOfStripe) Init(h *Header, r _fg.StreamReader) error {
	_de._dg = r
	return _de.parseHeader(h, r)
}
func (_ecbbg *PageInformationSegment) CombinationOperatorOverrideAllowed() bool { return _ecbbg._eae }
func (_gdfb *Header) readSegmentPageAssociation(_aga Documenter, _dde _fg.StreamReader, _eddg uint64, _agcg ...int) (_cggd error) {
	const _gdfbf = "\u0072\u0065\u0061\u0064\u0053\u0065\u0067\u006d\u0065\u006e\u0074P\u0061\u0067\u0065\u0041\u0073\u0073\u006f\u0063\u0069\u0061t\u0069\u006f\u006e"
	if !_gdfb.PageAssociationFieldSize {
		_bedc, _adda := _dde.ReadBits(8)
		if _adda != nil {
			return _fd.Wrap(_adda, _gdfbf, "\u0073\u0068\u006fr\u0074\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		_gdfb.PageAssociation = int(_bedc & 0xFF)
	} else {
		_aba, _ffce := _dde.ReadBits(32)
		if _ffce != nil {
			return _fd.Wrap(_ffce, _gdfbf, "l\u006f\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
		}
		_gdfb.PageAssociation = int(_aba & _f.MaxInt32)
	}
	if _eddg == 0 {
		return nil
	}
	if _gdfb.PageAssociation != 0 {
		_gecb, _degb := _aga.GetPage(_gdfb.PageAssociation)
		if _degb != nil {
			return _fd.Wrap(_degb, _gdfbf, "\u0061s\u0073\u006f\u0063\u0069a\u0074\u0065\u0064\u0020\u0070a\u0067e\u0020n\u006f\u0074\u0020\u0066\u006f\u0075\u006ed")
		}
		var _eagd int
		for _bfb := uint64(0); _bfb < _eddg; _bfb++ {
			_eagd = _agcg[_bfb]
			_gdfb.RTSegments[_bfb], _degb = _gecb.GetSegment(_eagd)
			if _degb != nil {
				var _eegc error
				_gdfb.RTSegments[_bfb], _eegc = _aga.GetGlobalSegment(_eagd)
				if _eegc != nil {
					return _fd.Wrapf(_degb, _gdfbf, "\u0072\u0065\u0066\u0065\u0072\u0065n\u0063\u0065\u0020s\u0065\u0067\u006de\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0020\u0061\u0074\u0020pa\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0072\u0020\u0069\u006e\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0073", _gdfb.PageAssociation)
				}
			}
		}
		return nil
	}
	for _gaed := uint64(0); _gaed < _eddg; _gaed++ {
		_gdfb.RTSegments[_gaed], _cggd = _aga.GetGlobalSegment(_agcg[_gaed])
		if _cggd != nil {
			return _fd.Wrapf(_cggd, _gdfbf, "\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u003a\u0020\u0027\u0025d\u0027\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064", _agcg[_gaed])
		}
	}
	return nil
}
func (_faefd *TextRegion) getSymbols() error {
	if _faefd.Header.RTSegments != nil {
		return _faefd.initSymbols()
	}
	return nil
}

type SymbolDictionary struct {
	_bcagf                      _fg.StreamReader
	SdrTemplate                 int8
	SdTemplate                  int8
	_gdae                       bool
	_gacg                       bool
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
	_gbbg                       uint32
	_cgfbb                      []*_d.Bitmap
	_ffa                        uint32
	_gbga                       []*_d.Bitmap
	_dgecb                      _ba.Tabler
	_agdg                       _ba.Tabler
	_afea                       _ba.Tabler
	_fggec                      _ba.Tabler
	_cba                        []*_d.Bitmap
	_bded                       []*_d.Bitmap
	_fffd                       *_b.Decoder
	_begf                       *TextRegion
	_eaeb                       *GenericRegion
	_debeg                      *GenericRefinementRegion
	_bagf                       *_b.DecoderStats
	_aabg                       *_b.DecoderStats
	_aafa                       *_b.DecoderStats
	_eagcd                      *_b.DecoderStats
	_gabe                       *_b.DecoderStats
	_abag                       *_b.DecoderStats
	_dfb                        *_b.DecoderStats
	_eegg                       *_b.DecoderStats
	_fac                        *_b.DecoderStats
	_cff                        int8
	_dbag                       *_d.Bitmaps
	_fagdd                      []int
	_dcbf                       map[int]int
	_cgbf                       bool
}

func (_bcgg *PageInformationSegment) readIsStriped() error {
	_ggfb, _cdc := _bcgg._ace.ReadBit()
	if _cdc != nil {
		return _cdc
	}
	if _ggfb == 1 {
		_bcgg.IsStripe = true
	}
	return nil
}
func (_ggca *GenericRegion) overrideAtTemplate3(_aee, _edbg, _gca, _gacf, _fgbfc int) int {
	_aee &= 0x3EF
	if _ggca.GBAtY[0] == 0 && _ggca.GBAtX[0] >= -int8(_fgbfc) {
		_aee |= (_gacf >> uint(7-(int8(_fgbfc)+_ggca.GBAtX[0])) & 0x1) << 4
	} else {
		_aee |= int(_ggca.getPixel(_edbg+int(_ggca.GBAtX[0]), _gca+int(_ggca.GBAtY[0]))) << 4
	}
	return _aee
}
func (_dfaf *HalftoneRegion) computeX(_gcdf, _bbed int) int {
	return _dfaf.shiftAndFill(int(_dfaf.HGridX) + _gcdf*int(_dfaf.HRegionY) + _bbed*int(_dfaf.HRegionX))
}

type RegionSegment struct {
	_dbge              _fg.StreamReader
	BitmapWidth        uint32
	BitmapHeight       uint32
	XLocation          uint32
	YLocation          uint32
	CombinaionOperator _d.CombinationOperator
}

func (_acfbd *SymbolDictionary) parseHeader() (_dfcbf error) {
	_da.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c \u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		if _dfcbf != nil {
			_da.Log.Trace("\u005bS\u0059\u004dB\u004f\u004c\u0020\u0044I\u0043\u0054\u0049O\u004e\u0041\u0052\u0059\u005d\u005b\u0050\u0041\u0052SE\u002d\u0048\u0045A\u0044\u0045R\u005d\u0020\u0066\u0061\u0069\u006ce\u0064\u002e \u0025\u0076", _dfcbf)
		} else {
			_da.Log.Trace("\u005b\u0053\u0059\u004d\u0042\u004f\u004c \u0044\u0049\u0043T\u0049\u004f\u004e\u0041R\u0059\u005d\u005b\u0050\u0041\u0052\u0053\u0045\u002d\u0048\u0045\u0041\u0044\u0045\u0052\u005d\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		}
	}()
	if _dfcbf = _acfbd.readRegionFlags(); _dfcbf != nil {
		return _dfcbf
	}
	if _dfcbf = _acfbd.setAtPixels(); _dfcbf != nil {
		return _dfcbf
	}
	if _dfcbf = _acfbd.setRefinementAtPixels(); _dfcbf != nil {
		return _dfcbf
	}
	if _dfcbf = _acfbd.readNumberOfExportedSymbols(); _dfcbf != nil {
		return _dfcbf
	}
	if _dfcbf = _acfbd.readNumberOfNewSymbols(); _dfcbf != nil {
		return _dfcbf
	}
	if _dfcbf = _acfbd.setInSyms(); _dfcbf != nil {
		return _dfcbf
	}
	if _acfbd._gacg {
		_eeac := _acfbd.Header.RTSegments
		for _fcfa := len(_eeac) - 1; _fcfa >= 0; _fcfa-- {
			if _eeac[_fcfa].Type == 0 {
				_gbba, _cgcf := _eeac[_fcfa].SegmentData.(*SymbolDictionary)
				if !_cgcf {
					_dfcbf = _e.Errorf("\u0072\u0065\u006c\u0061\u0074\u0065\u0064\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074:\u0020\u0025\u0076\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020S\u0079\u006d\u0062\u006f\u006c\u0020\u0044\u0069\u0063\u0074\u0069\u006fna\u0072\u0079\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074", _eeac[_fcfa])
					return _dfcbf
				}
				if _gbba._gacg {
					_acfbd.setRetainedCodingContexts(_gbba)
				}
				break
			}
		}
	}
	if _dfcbf = _acfbd.checkInput(); _dfcbf != nil {
		return _dfcbf
	}
	return nil
}
func (_bbcf *SymbolDictionary) getSymbol(_edfde int) (*_d.Bitmap, error) {
	const _cgde = "\u0067e\u0074\u0053\u0079\u006d\u0062\u006fl"
	_ebaa, _cbcc := _bbcf._dbag.GetBitmap(_bbcf._fagdd[_edfde])
	if _cbcc != nil {
		return nil, _fd.Wrap(_cbcc, _cgde, "\u0063\u0061n\u0027\u0074\u0020g\u0065\u0074\u0020\u0073\u0079\u006d\u0062\u006f\u006c")
	}
	return _ebaa, nil
}
func (_cafa *GenericRefinementRegion) readAtPixels() error {
	_cafa.GrAtX = make([]int8, 2)
	_cafa.GrAtY = make([]int8, 2)
	_ggd, _ceb := _cafa._geg.ReadByte()
	if _ceb != nil {
		return _ceb
	}
	_cafa.GrAtX[0] = int8(_ggd)
	_ggd, _ceb = _cafa._geg.ReadByte()
	if _ceb != nil {
		return _ceb
	}
	_cafa.GrAtY[0] = int8(_ggd)
	_ggd, _ceb = _cafa._geg.ReadByte()
	if _ceb != nil {
		return _ceb
	}
	_cafa.GrAtX[1] = int8(_ggd)
	_ggd, _ceb = _cafa._geg.ReadByte()
	if _ceb != nil {
		return _ceb
	}
	_cafa.GrAtY[1] = int8(_ggd)
	return nil
}
func (_dd *EndOfStripe) parseHeader(_fc *Header, _aa _fg.StreamReader) error {
	_adb, _gab := _dd._dg.ReadBits(32)
	if _gab != nil {
		return _gab
	}
	_dd._dc = int(_adb & _f.MaxInt32)
	return nil
}
func (_gbcb *TextRegion) decodeCurrentT() (int64, error) {
	if _gbcb.SbStrips != 1 {
		if _gbcb.IsHuffmanEncoded {
			_abcf, _faca := _gbcb._effb.ReadBits(byte(_gbcb.LogSBStrips))
			return int64(_abcf), _faca
		}
		_fef, _ecg := _gbcb._dffe.DecodeInt(_gbcb._dace)
		if _ecg != nil {
			return 0, _ecg
		}
		return int64(_fef), nil
	}
	return 0, nil
}
func (_fgbfg *HalftoneRegion) parseHeader() error {
	if _efeb := _fgbfg.RegionSegment.parseHeader(); _efeb != nil {
		return _efeb
	}
	_ddac, _edeg := _fgbfg._gdfe.ReadBit()
	if _edeg != nil {
		return _edeg
	}
	_fgbfg.HDefaultPixel = int8(_ddac)
	_gegd, _edeg := _fgbfg._gdfe.ReadBits(3)
	if _edeg != nil {
		return _edeg
	}
	_fgbfg.CombinationOperator = _d.CombinationOperator(_gegd & 0xf)
	_ddac, _edeg = _fgbfg._gdfe.ReadBit()
	if _edeg != nil {
		return _edeg
	}
	if _ddac == 1 {
		_fgbfg.HSkipEnabled = true
	}
	_gegd, _edeg = _fgbfg._gdfe.ReadBits(2)
	if _edeg != nil {
		return _edeg
	}
	_fgbfg.HTemplate = byte(_gegd & 0xf)
	_ddac, _edeg = _fgbfg._gdfe.ReadBit()
	if _edeg != nil {
		return _edeg
	}
	if _ddac == 1 {
		_fgbfg.IsMMREncoded = true
	}
	_gegd, _edeg = _fgbfg._gdfe.ReadBits(32)
	if _edeg != nil {
		return _edeg
	}
	_fgbfg.HGridWidth = uint32(_gegd & _f.MaxUint32)
	_gegd, _edeg = _fgbfg._gdfe.ReadBits(32)
	if _edeg != nil {
		return _edeg
	}
	_fgbfg.HGridHeight = uint32(_gegd & _f.MaxUint32)
	_gegd, _edeg = _fgbfg._gdfe.ReadBits(32)
	if _edeg != nil {
		return _edeg
	}
	_fgbfg.HGridX = int32(_gegd & _f.MaxInt32)
	_gegd, _edeg = _fgbfg._gdfe.ReadBits(32)
	if _edeg != nil {
		return _edeg
	}
	_fgbfg.HGridY = int32(_gegd & _f.MaxInt32)
	_gegd, _edeg = _fgbfg._gdfe.ReadBits(16)
	if _edeg != nil {
		return _edeg
	}
	_fgbfg.HRegionX = uint16(_gegd & _f.MaxUint16)
	_gegd, _edeg = _fgbfg._gdfe.ReadBits(16)
	if _edeg != nil {
		return _edeg
	}
	_fgbfg.HRegionY = uint16(_gegd & _f.MaxUint16)
	if _edeg = _fgbfg.computeSegmentDataStructure(); _edeg != nil {
		return _edeg
	}
	return _fgbfg.checkInput()
}
func (_eagc *Header) referenceSize() uint {
	switch {
	case _eagc.SegmentNumber <= 255:
		return 1
	case _eagc.SegmentNumber <= 65535:
		return 2
	default:
		return 4
	}
}
func (_efga *TextRegion) getUserTable(_ggbb int) (_ba.Tabler, error) {
	const _dbcc = "\u0067\u0065\u0074U\u0073\u0065\u0072\u0054\u0061\u0062\u006c\u0065"
	var _ccga int
	for _, _fbdac := range _efga.Header.RTSegments {
		if _fbdac.Type == 53 {
			if _ccga == _ggbb {
				_cege, _eaca := _fbdac.GetSegmentData()
				if _eaca != nil {
					return nil, _eaca
				}
				_abac, _aceg := _cege.(*TableSegment)
				if !_aceg {
					_da.Log.Debug(_e.Sprintf("\u0073\u0065\u0067\u006d\u0065\u006e\u0074 \u0077\u0069\u0074h\u0020\u0054\u0079p\u0065\u00205\u0033\u0020\u002d\u0020\u0061\u006ed\u0020in\u0064\u0065\u0078\u003a\u0020\u0025\u0064\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0054\u0061\u0062\u006c\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074", _fbdac.SegmentNumber))
					return nil, _fd.Error(_dbcc, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074 \u0077\u0069\u0074h\u0020\u0054\u0079\u0070e\u0020\u0035\u0033\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u002a\u0054\u0061\u0062\u006c\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
				}
				return _ba.NewEncodedTable(_abac)
			}
			_ccga++
		}
	}
	return nil, nil
}

var _ SegmentEncoder = &GenericRegion{}

func (_abbgc *GenericRegion) readGBAtPixels(_fca int) error {
	const _edad = "\u0072\u0065\u0061\u0064\u0047\u0042\u0041\u0074\u0050i\u0078\u0065\u006c\u0073"
	_abbgc.GBAtX = make([]int8, _fca)
	_abbgc.GBAtY = make([]int8, _fca)
	for _gaeg := 0; _gaeg < _fca; _gaeg++ {
		_fcf, _gacfe := _abbgc._baf.ReadByte()
		if _gacfe != nil {
			return _fd.Wrapf(_gacfe, _edad, "\u0058\u0020\u0061t\u0020\u0069\u003a\u0020\u0027\u0025\u0064\u0027", _gaeg)
		}
		_abbgc.GBAtX[_gaeg] = int8(_fcf)
		_fcf, _gacfe = _abbgc._baf.ReadByte()
		if _gacfe != nil {
			return _fd.Wrapf(_gacfe, _edad, "\u0059\u0020\u0061t\u0020\u0069\u003a\u0020\u0027\u0025\u0064\u0027", _gaeg)
		}
		_abbgc.GBAtY[_gaeg] = int8(_fcf)
	}
	return nil
}
func (_fae *Header) CleanSegmentData() {
	if _fae.SegmentData != nil {
		_fae.SegmentData = nil
	}
}
func (_daba *TextRegion) decodeIb(_cddd, _egec int64) (*_d.Bitmap, error) {
	const _gcbb = "\u0064\u0065\u0063\u006f\u0064\u0065\u0049\u0062"
	var (
		_bdfbe error
		_cceb  *_d.Bitmap
	)
	if _cddd == 0 {
		if int(_egec) > len(_daba.Symbols)-1 {
			return nil, _fd.Error(_gcbb, "\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0049\u0042\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0020\u0069\u006e\u0064\u0065x\u0020\u006f\u0075\u0074\u0020o\u0066\u0020r\u0061\u006e\u0067\u0065")
		}
		return _daba.Symbols[int(_egec)], nil
	}
	var _baae, _eaaga, _acaf, _eggab int64
	_baae, _bdfbe = _daba.decodeRdw()
	if _bdfbe != nil {
		return nil, _fd.Wrap(_bdfbe, _gcbb, "")
	}
	_eaaga, _bdfbe = _daba.decodeRdh()
	if _bdfbe != nil {
		return nil, _fd.Wrap(_bdfbe, _gcbb, "")
	}
	_acaf, _bdfbe = _daba.decodeRdx()
	if _bdfbe != nil {
		return nil, _fd.Wrap(_bdfbe, _gcbb, "")
	}
	_eggab, _bdfbe = _daba.decodeRdy()
	if _bdfbe != nil {
		return nil, _fd.Wrap(_bdfbe, _gcbb, "")
	}
	if _daba.IsHuffmanEncoded {
		if _, _bdfbe = _daba.decodeSymInRefSize(); _bdfbe != nil {
			return nil, _fd.Wrap(_bdfbe, _gcbb, "")
		}
		_daba._effb.Align()
	}
	_dgbfe := _daba.Symbols[_egec]
	_effac := uint32(_dgbfe.Width)
	_dfbgb := uint32(_dgbfe.Height)
	_ccdd := int32(uint32(_baae)>>1) + int32(_acaf)
	_efba := int32(uint32(_eaaga)>>1) + int32(_eggab)
	if _daba._edge == nil {
		_daba._edge = _cda(_daba._effb, nil)
	}
	_daba._edge.setParameters(_daba._aeed, _daba._dffe, _daba.SbrTemplate, _effac+uint32(_baae), _dfbgb+uint32(_eaaga), _dgbfe, _ccdd, _efba, false, _daba.SbrATX, _daba.SbrATY)
	_cceb, _bdfbe = _daba._edge.GetRegionBitmap()
	if _bdfbe != nil {
		return nil, _fd.Wrap(_bdfbe, _gcbb, "\u0067\u0072\u0066")
	}
	if _daba.IsHuffmanEncoded {
		_daba._effb.Align()
	}
	return _cceb, nil
}
func (_ec *GenericRefinementRegion) GetRegionBitmap() (*_d.Bitmap, error) {
	var _bc error
	_da.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0047\u0065\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e")
	defer func() {
		if _bc != nil {
			_da.Log.Trace("[\u0047\u0045\u004e\u0045\u0052\u0049\u0043\u002d\u0052E\u0046\u002d\u0052\u0045\u0047\u0049\u004fN]\u0020\u0047\u0065\u0074R\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006dap\u0020\u0066a\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _bc)
		} else {
			_da.Log.Trace("\u005b\u0047E\u004e\u0045\u0052\u0049\u0043\u002d\u0052\u0045\u0046\u002d\u0052\u0045\u0047\u0049\u004f\u004e\u005d\u0020\u0047\u0065\u0074\u0052\u0065\u0067\u0069\u006f\u006e\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e")
		}
	}()
	if _ec.RegionBitmap != nil {
		return _ec.RegionBitmap, nil
	}
	_dgb := 0
	if _ec.ReferenceBitmap == nil {
		_ec.ReferenceBitmap, _bc = _ec.getGrReference()
		if _bc != nil {
			return nil, _bc
		}
	}
	if _ec._gf == nil {
		_ec._gf, _bc = _b.New(_ec._geg)
		if _bc != nil {
			return nil, _bc
		}
	}
	if _ec._cdb == nil {
		_ec._cdb = _b.NewStats(8192, 1)
	}
	_ec.RegionBitmap = _d.New(int(_ec.RegionInfo.BitmapWidth), int(_ec.RegionInfo.BitmapHeight))
	if _ec.TemplateID == 0 {
		if _bc = _ec.updateOverride(); _bc != nil {
			return nil, _bc
		}
	}
	_fed := (_ec.RegionBitmap.Width + 7) & -8
	var _ag int
	if _ec.IsTPGROn {
		_ag = int(-_ec.ReferenceDY) * _ec.ReferenceBitmap.RowStride
	}
	_dbb := _ag + 1
	for _dee := 0; _dee < _ec.RegionBitmap.Height; _dee++ {
		if _ec.IsTPGROn {
			_eb, _agc := _ec.decodeSLTP()
			if _agc != nil {
				return nil, _agc
			}
			_dgb ^= _eb
		}
		if _dgb == 0 {
			_bc = _ec.decodeOptimized(_dee, _ec.RegionBitmap.Width, _ec.RegionBitmap.RowStride, _ec.ReferenceBitmap.RowStride, _fed, _ag, _dbb)
			if _bc != nil {
				return nil, _bc
			}
		} else {
			_bc = _ec.decodeTypicalPredictedLine(_dee, _ec.RegionBitmap.Width, _ec.RegionBitmap.RowStride, _ec.ReferenceBitmap.RowStride, _fed, _ag)
			if _bc != nil {
				return nil, _bc
			}
		}
	}
	return _ec.RegionBitmap, nil
}
func (_gffg *HalftoneRegion) Init(hd *Header, r _fg.StreamReader) error {
	_gffg._gdfe = r
	_gffg._fdag = hd
	_gffg.RegionSegment = NewRegionSegment(r)
	return _gffg.parseHeader()
}

type OrganizationType uint8

func (_ebgb *SymbolDictionary) decodeHeightClassDeltaHeightWithHuffman() (int64, error) {
	switch _ebgb.SdHuffDecodeHeightSelection {
	case 0:
		_gedc, _ebag := _ba.GetStandardTable(4)
		if _ebag != nil {
			return 0, _ebag
		}
		return _gedc.Decode(_ebgb._bcagf)
	case 1:
		_edfg, _affdf := _ba.GetStandardTable(5)
		if _affdf != nil {
			return 0, _affdf
		}
		return _edfg.Decode(_ebgb._bcagf)
	case 3:
		if _ebgb._dgecb == nil {
			_afgd, _fcbbd := _ba.GetStandardTable(0)
			if _fcbbd != nil {
				return 0, _fcbbd
			}
			_ebgb._dgecb = _afgd
		}
		return _ebgb._dgecb.Decode(_ebgb._bcagf)
	}
	return 0, nil
}
func (_effe *GenericRegion) decodeSLTP() (int, error) {
	switch _effe.GBTemplate {
	case 0:
		_effe._aae.SetIndex(0x9B25)
	case 1:
		_effe._aae.SetIndex(0x795)
	case 2:
		_effe._aae.SetIndex(0xE5)
	case 3:
		_effe._aae.SetIndex(0x195)
	}
	return _effe._acb.DecodeBit(_effe._aae)
}
func (_dbe *TextRegion) setCodingStatistics() error {
	if _dbe._bgad == nil {
		_dbe._bgad = _b.NewStats(512, 1)
	}
	if _dbe._geaa == nil {
		_dbe._geaa = _b.NewStats(512, 1)
	}
	if _dbe._dafb == nil {
		_dbe._dafb = _b.NewStats(512, 1)
	}
	if _dbe._dace == nil {
		_dbe._dace = _b.NewStats(512, 1)
	}
	if _dbe._gbfge == nil {
		_dbe._gbfge = _b.NewStats(512, 1)
	}
	if _dbe._egcff == nil {
		_dbe._egcff = _b.NewStats(512, 1)
	}
	if _dbe._acgg == nil {
		_dbe._acgg = _b.NewStats(512, 1)
	}
	if _dbe._beabb == nil {
		_dbe._beabb = _b.NewStats(1<<uint(_dbe._ebcb), 1)
	}
	if _dbe._beda == nil {
		_dbe._beda = _b.NewStats(512, 1)
	}
	if _dbe._bfbb == nil {
		_dbe._bfbb = _b.NewStats(512, 1)
	}
	if _dbe._dffe == nil {
		var _ddcf error
		_dbe._dffe, _ddcf = _b.New(_dbe._effb)
		if _ddcf != nil {
			return _ddcf
		}
	}
	return nil
}
func (_fgd *Header) GetSegmentData() (Segmenter, error) {
	var _dagf Segmenter
	if _fgd.SegmentData != nil {
		_dagf = _fgd.SegmentData
	}
	if _dagf == nil {
		_bdfg, _ccge := _ffdg[_fgd.Type]
		if !_ccge {
			return nil, _e.Errorf("\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0073\u002f\u0020\u0025\u0064\u0020\u0063\u0072e\u0061t\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e\u0020", _fgd.Type, _fgd.Type)
		}
		_dagf = _bdfg()
		_da.Log.Trace("\u005b\u0053E\u0047\u004d\u0045\u004e\u0054-\u0048\u0045\u0041\u0044\u0045R\u005d\u005b\u0023\u0025\u0064\u005d\u0020\u0047\u0065\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0044\u0061\u0074\u0061\u0020\u0061\u0074\u0020\u004f\u0066\u0066\u0073\u0065\u0074\u003a\u0020\u0025\u0030\u0034\u0058", _fgd.SegmentNumber, _fgd.SegmentDataStartOffset)
		_ffca, _ebe := _fgd.subInputReader()
		if _ebe != nil {
			return nil, _ebe
		}
		if _beab := _dagf.Init(_fgd, _ffca); _beab != nil {
			_da.Log.Debug("\u0049\u006e\u0069\u0074 \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076 \u0066o\u0072\u0020\u0074\u0079\u0070\u0065\u003a \u0025\u0054", _beab, _dagf)
			return nil, _beab
		}
		_fgd.SegmentData = _dagf
	}
	return _dagf, nil
}
func (_addgc *GenericRegion) setOverrideFlag(_agbc int) {
	_addgc.GBAtOverride[_agbc] = true
	_addgc._fgf = true
}
func (_ebfgb *Header) writeFlags(_dbbfd _fg.BinaryWriter) (_gacc error) {
	const _debe = "\u0048\u0065\u0061\u0064\u0065\u0072\u002e\u0077\u0072\u0069\u0074\u0065F\u006c\u0061\u0067\u0073"
	_gcbf := byte(_ebfgb.Type)
	if _gacc = _dbbfd.WriteByte(_gcbf); _gacc != nil {
		return _fd.Wrap(_gacc, _debe, "\u0077\u0072\u0069ti\u006e\u0067\u0020\u0073\u0065\u0067\u006d\u0065\u006et\u0020t\u0079p\u0065 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	if !_ebfgb.RetainFlag && !_ebfgb.PageAssociationFieldSize {
		return nil
	}
	if _gacc = _dbbfd.SkipBits(-8); _gacc != nil {
		return _fd.Wrap(_gacc, _debe, "\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0062\u0061\u0063\u006b\u0020\u0074\u0068\u0065\u0020\u0062\u0069\u0074\u0073\u0020\u0066\u0061il\u0065\u0064")
	}
	var _dfg int
	if _ebfgb.RetainFlag {
		_dfg = 1
	}
	if _gacc = _dbbfd.WriteBit(_dfg); _gacc != nil {
		return _fd.Wrap(_gacc, _debe, "\u0072\u0065\u0074\u0061in\u0020\u0072\u0065\u0074\u0061\u0069\u006e\u0020\u0066\u006c\u0061\u0067\u0073")
	}
	_dfg = 0
	if _ebfgb.PageAssociationFieldSize {
		_dfg = 1
	}
	if _gacc = _dbbfd.WriteBit(_dfg); _gacc != nil {
		return _fd.Wrap(_gacc, _debe, "p\u0061\u0067\u0065\u0020as\u0073o\u0063\u0069\u0061\u0074\u0069o\u006e\u0020\u0066\u006c\u0061\u0067")
	}
	_dbbfd.FinishByte()
	return nil
}
func (_gbbgb *SymbolDictionary) decodeAggregate(_fdgbg, _badfff uint32) error {
	var (
		_gbea int64
		_dbdb error
	)
	if _gbbgb.IsHuffmanEncoded {
		_gbea, _dbdb = _gbbgb.huffDecodeRefAggNInst()
		if _dbdb != nil {
			return _dbdb
		}
	} else {
		_gagb, _fccgg := _gbbgb._fffd.DecodeInt(_gbbgb._eagcd)
		if _fccgg != nil {
			return _fccgg
		}
		_gbea = int64(_gagb)
	}
	if _gbea > 1 {
		return _gbbgb.decodeThroughTextRegion(_fdgbg, _badfff, uint32(_gbea))
	} else if _gbea == 1 {
		return _gbbgb.decodeRefinedSymbol(_fdgbg, _badfff)
	}
	return nil
}

var _ _ba.BasicTabler = &TableSegment{}

func _daae(_dbed _fg.StreamReader, _cdfgf *Header) *TextRegion {
	_bfdff := &TextRegion{_effb: _dbed, Header: _cdfgf, RegionInfo: NewRegionSegment(_dbed)}
	return _bfdff
}
func (_debab *PatternDictionary) Init(h *Header, r _fg.StreamReader) error {
	_debab._fbdd = r
	return _debab.parseHeader()
}
func (_cgca *Header) writeSegmentDataLength(_aggc _fg.BinaryWriter) (_ccgg int, _fgeg error) {
	_fbb := make([]byte, 4)
	_fe.BigEndian.PutUint32(_fbb, uint32(_cgca.SegmentDataLength))
	if _ccgg, _fgeg = _aggc.Write(_fbb); _fgeg != nil {
		return 0, _fd.Wrap(_fgeg, "\u0048\u0065a\u0064\u0065\u0072\u002e\u0077\u0072\u0069\u0074\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u0044\u0061\u0074\u0061\u004c\u0065ng\u0074\u0068", "")
	}
	return _ccgg, nil
}
func (_aaef *PageInformationSegment) Size() int { return 19 }
func (_bdce *GenericRegion) GetRegionBitmap() (_efc *_d.Bitmap, _efcb error) {
	if _bdce.Bitmap != nil {
		return _bdce.Bitmap, nil
	}
	if _bdce.IsMMREncoded {
		if _bdce._gcd == nil {
			_bdce._gcd, _efcb = _ad.New(_bdce._baf, int(_bdce.RegionSegment.BitmapWidth), int(_bdce.RegionSegment.BitmapHeight), _bdce.DataOffset, _bdce.DataLength)
			if _efcb != nil {
				return nil, _efcb
			}
		}
		_bdce.Bitmap, _efcb = _bdce._gcd.UncompressMMR()
		return _bdce.Bitmap, _efcb
	}
	if _efcb = _bdce.updateOverrideFlags(); _efcb != nil {
		return nil, _efcb
	}
	var _adf int
	if _bdce._acb == nil {
		_bdce._acb, _efcb = _b.New(_bdce._baf)
		if _efcb != nil {
			return nil, _efcb
		}
	}
	if _bdce._aae == nil {
		_bdce._aae = _b.NewStats(65536, 1)
	}
	_bdce.Bitmap = _d.New(int(_bdce.RegionSegment.BitmapWidth), int(_bdce.RegionSegment.BitmapHeight))
	_bge := int(uint32(_bdce.Bitmap.Width+7) & (^uint32(7)))
	for _abg := 0; _abg < _bdce.Bitmap.Height; _abg++ {
		if _bdce.IsTPGDon {
			var _egb int
			_egb, _efcb = _bdce.decodeSLTP()
			if _efcb != nil {
				return nil, _efcb
			}
			_adf ^= _egb
		}
		if _adf == 1 {
			if _abg > 0 {
				if _efcb = _bdce.copyLineAbove(_abg); _efcb != nil {
					return nil, _efcb
				}
			}
		} else {
			if _efcb = _bdce.decodeLine(_abg, _bdce.Bitmap.Width, _bge); _efcb != nil {
				return nil, _efcb
			}
		}
	}
	return _bdce.Bitmap, nil
}
