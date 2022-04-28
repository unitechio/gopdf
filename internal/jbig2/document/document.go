package document

import (
	_g "encoding/binary"
	_dg "fmt"
	_c "io"
	_dc "math"
	_df "runtime/debug"

	_dff "bitbucket.org/shenghui0779/gopdf/common"
	_f "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_da "bitbucket.org/shenghui0779/gopdf/internal/jbig2/basic"
	_ca "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_ea "bitbucket.org/shenghui0779/gopdf/internal/jbig2/document/segments"
	_e "bitbucket.org/shenghui0779/gopdf/internal/jbig2/encoder/classer"
	_fb "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

func (_dccg *Page) createNormalPage(_acb *_ea.PageInformationSegment) error {
	const _fgg = "\u0063\u0072e\u0061\u0074\u0065N\u006f\u0072\u006d\u0061\u006c\u0050\u0061\u0067\u0065"
	_dccg.Bitmap = _ca.New(_acb.PageBMWidth, _acb.PageBMHeight)
	if _acb.DefaultPixelValue != 0 {
		_dccg.Bitmap.SetDefaultPixel()
	}
	for _, _adec := range _dccg.Segments {
		switch _adec.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_dff.Log.Trace("\u0047\u0065\u0074\u0074in\u0067\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0064", _adec.SegmentNumber)
			_fggd, _fdf := _adec.GetSegmentData()
			if _fdf != nil {
				return _fdf
			}
			_bbe, _gaf := _fggd.(_ea.Regioner)
			if !_gaf {
				_dff.Log.Debug("\u0053\u0065g\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0054\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0052\u0065\u0067\u0069on\u0065\u0072", _fggd)
				return _fb.Errorf(_fgg, "i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062i\u0067\u0032\u0020\u0073\u0065\u0067\u006den\u0074\u0020\u0074\u0079p\u0065\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061 R\u0065\u0067i\u006f\u006e\u0065\u0072\u003a\u0020\u0025\u0054", _fggd)
			}
			_dbcc, _fdf := _bbe.GetRegionBitmap()
			if _fdf != nil {
				return _fb.Wrap(_fdf, _fgg, "")
			}
			if _dccg.fitsPage(_acb, _dbcc) {
				_dccg.Bitmap = _dbcc
			} else {
				_ecc := _bbe.GetRegionInfo()
				_gfb := _dccg.getCombinationOperator(_acb, _ecc.CombinaionOperator)
				_fdf = _ca.Blit(_dbcc, _dccg.Bitmap, int(_ecc.XLocation), int(_ecc.YLocation), _gfb)
				if _fdf != nil {
					return _fb.Wrap(_fdf, _fgg, "")
				}
			}
		}
	}
	return nil
}
func (_fabe *Page) countRegions() int {
	var _cdd int
	for _, _dbe := range _fabe.Segments {
		switch _dbe.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_cdd++
		}
	}
	return _cdd
}
func (_bcc *Page) GetResolutionX() (int, error) { return _bcc.getResolutionX() }
func (_aca *Page) composePageBitmap() error {
	const _bac = "\u0063\u006f\u006d\u0070\u006f\u0073\u0065\u0050\u0061\u0067\u0065\u0042i\u0074\u006d\u0061\u0070"
	if _aca.PageNumber == 0 {
		return nil
	}
	_ddd := _aca.getPageInformationSegment()
	if _ddd == nil {
		return _fb.Error(_bac, "\u0070\u0061\u0067e \u0069\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069o\u006e \u0073e\u0067m\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_dfd, _daa := _ddd.GetSegmentData()
	if _daa != nil {
		return _daa
	}
	_cgd, _gdgg := _dfd.(*_ea.PageInformationSegment)
	if !_gdgg {
		return _fb.Error(_bac, "\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006da\u0074\u0069\u006f\u006e \u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065")
	}
	if _daa = _aca.createPage(_cgd); _daa != nil {
		return _fb.Wrap(_daa, _bac, "")
	}
	_aca.clearSegmentData()
	return nil
}
func _gda(_ead int) int {
	_ec := 0
	_ac := (_ead & (_ead - 1)) == 0
	_ead >>= 1
	for ; _ead != 0; _ead >>= 1 {
		_ec++
	}
	if _ac {
		return _ec
	}
	return _ec + 1
}
func (_afa *Document) nextPageNumber() uint32 { _afa.NumberOfPages++; return _afa.NumberOfPages }
func (_gba *Page) Encode(w _f.BinaryWriter) (_dcc int, _cgfd error) {
	const _gge = "P\u0061\u0067\u0065\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	var _effb int
	for _, _aded := range _gba.Segments {
		if _effb, _cgfd = _aded.Encode(w); _cgfd != nil {
			return _dcc, _fb.Wrap(_cgfd, _gge, "")
		}
		_dcc += _effb
	}
	return _dcc, nil
}
func (_fbb *Document) parseFileHeader() error {
	const _geg = "\u0070a\u0072s\u0065\u0046\u0069\u006c\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_, _abe := _fbb.InputStream.Seek(8, _c.SeekStart)
	if _abe != nil {
		return _fb.Wrap(_abe, _geg, "\u0069\u0064")
	}
	_, _abe = _fbb.InputStream.ReadBits(5)
	if _abe != nil {
		return _fb.Wrap(_abe, _geg, "\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0064\u0020\u0062\u0069\u0074\u0073")
	}
	_cfg, _abe := _fbb.InputStream.ReadBit()
	if _abe != nil {
		return _fb.Wrap(_abe, _geg, "\u0065x\u0074e\u006e\u0064\u0065\u0064\u0020t\u0065\u006dp\u006c\u0061\u0074\u0065\u0073")
	}
	if _cfg == 1 {
		_fbb.GBUseExtTemplate = true
	}
	_cfg, _abe = _fbb.InputStream.ReadBit()
	if _abe != nil {
		return _fb.Wrap(_abe, _geg, "\u0075\u006e\u006b\u006eow\u006e\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	if _cfg != 1 {
		_fbb.NumberOfPagesUnknown = false
	}
	_cfg, _abe = _fbb.InputStream.ReadBit()
	if _abe != nil {
		return _fb.Wrap(_abe, _geg, "\u006f\u0072\u0067\u0061\u006e\u0069\u007a\u0061\u0074\u0069\u006f\u006e \u0074\u0079\u0070\u0065")
	}
	_fbb.OrganizationType = _ea.OrganizationType(_cfg)
	if !_fbb.NumberOfPagesUnknown {
		_fbb.NumberOfPages, _abe = _fbb.InputStream.ReadUint32()
		if _abe != nil {
			return _fb.Wrap(_abe, _geg, "\u006eu\u006db\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0061\u0067\u0065\u0073")
		}
		_fbb._a = 13
	}
	return nil
}
func (_fdd *Document) mapData() error {
	const _eece = "\u006da\u0070\u0044\u0061\u0074\u0061"
	var (
		_eaf []*_ea.Header
		_dcf int64
		_dbc _ea.Type
	)
	_fcg, _gcd := _fdd.isFileHeaderPresent()
	if _gcd != nil {
		return _fb.Wrap(_gcd, _eece, "")
	}
	if _fcg {
		if _gcd = _fdd.parseFileHeader(); _gcd != nil {
			return _fb.Wrap(_gcd, _eece, "")
		}
		_dcf += int64(_fdd._a)
		_fdd.FullHeaders = true
	}
	var (
		_fddd *Page
		_cd   bool
	)
	for _dbc != 51 && !_cd {
		_fef, _bce := _ea.NewHeader(_fdd, _fdd.InputStream, _dcf, _fdd.OrganizationType)
		if _bce != nil {
			return _fb.Wrap(_bce, _eece, "")
		}
		_dff.Log.Trace("\u0044\u0065c\u006f\u0064\u0069\u006eg\u0020\u0073e\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0025\u0064\u002c\u0020\u0054\u0079\u0070e\u003a\u0020\u0025\u0073", _fef.SegmentNumber, _fef.Type)
		_dbc = _fef.Type
		if _dbc != _ea.TEndOfFile {
			if _fef.PageAssociation != 0 {
				_fddd = _fdd.Pages[_fef.PageAssociation]
				if _fddd == nil {
					_fddd = _bgge(_fdd, _fef.PageAssociation)
					_fdd.Pages[_fef.PageAssociation] = _fddd
					if _fdd.NumberOfPagesUnknown {
						_fdd.NumberOfPages++
					}
				}
				_fddd.Segments = append(_fddd.Segments, _fef)
			} else {
				_fdd.GlobalSegments.AddSegment(_fef)
			}
		}
		_eaf = append(_eaf, _fef)
		_dcf = _fdd.InputStream.StreamPosition()
		if _fdd.OrganizationType == _ea.OSequential {
			_dcf += int64(_fef.SegmentDataLength)
		}
		_cd, _bce = _fdd.reachedEOF(_dcf)
		if _bce != nil {
			_dff.Log.Debug("\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0072\u0065\u0061\u0063h\u0065\u0064\u0020\u0045\u004f\u0046\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _bce)
			return _fb.Wrap(_bce, _eece, "")
		}
	}
	_fdd.determineRandomDataOffsets(_eaf, uint64(_dcf))
	return nil
}
func (_ffg *Document) addSymbolDictionary(_ccg int, _de *_ca.Bitmaps, _bff []int, _ece map[int]int, _fba bool) (*_ea.Header, error) {
	const _ccd = "\u0061\u0064\u0064\u0053ym\u0062\u006f\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079"
	_cba := &_ea.SymbolDictionary{}
	if _fgf := _cba.InitEncode(_de, _bff, _ece, _fba); _fgf != nil {
		return nil, _fgf
	}
	_dee := &_ea.Header{Type: _ea.TSymbolDictionary, PageAssociation: _ccg, SegmentData: _cba}
	if _ccg == 0 {
		if _ffg.GlobalSegments == nil {
			_ffg.GlobalSegments = &Globals{}
		}
		_ffg.GlobalSegments.AddSegment(_dee)
		return _dee, nil
	}
	_ad, _adc := _ffg.Pages[_ccg]
	if !_adc {
		return nil, _fb.Errorf(_ccd, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _ccg)
	}
	var (
		_fcf int
		_ee  *_ea.Header
	)
	for _fcf, _ee = range _ad.Segments {
		if _ee.Type == _ea.TPageInformation {
			break
		}
	}
	_fcf++
	_ad.Segments = append(_ad.Segments, nil)
	copy(_ad.Segments[_fcf+1:], _ad.Segments[_fcf:])
	_ad.Segments[_fcf] = _dee
	return _dee, nil
}

var _gc = []byte{0x97, 0x4A, 0x42, 0x32, 0x0D, 0x0A, 0x1A, 0x0A}

func (_bgfe *Document) reachedEOF(_cdg int64) (bool, error) {
	const _gcb = "\u0072\u0065\u0061\u0063\u0068\u0065\u0064\u0045\u004f\u0046"
	_, _gcbe := _bgfe.InputStream.Seek(_cdg, _c.SeekStart)
	if _gcbe != nil {
		_dff.Log.Debug("\u0072\u0065\u0061c\u0068\u0065\u0064\u0045\u004f\u0046\u0020\u002d\u0020\u0064\u002e\u0049\u006e\u0070\u0075\u0074\u0053\u0074\u0072\u0065\u0061\u006d\u002e\u0053\u0065\u0065\u006b\u0020\u0066a\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _gcbe)
		return false, _fb.Wrap(_gcbe, _gcb, "\u0069n\u0070\u0075\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020s\u0065\u0065\u006b\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_, _gcbe = _bgfe.InputStream.ReadBits(32)
	if _gcbe == _c.EOF {
		return true, nil
	} else if _gcbe != nil {
		return false, _fb.Wrap(_gcbe, _gcb, "")
	}
	return false, nil
}

type EncodingMethod int

func (_cee *Page) collectPageStripes() (_ccff []_ea.Segmenter, _fbae error) {
	const _cgdg = "\u0063o\u006cl\u0065\u0063\u0074\u0050\u0061g\u0065\u0053t\u0072\u0069\u0070\u0065\u0073"
	var _aebae _ea.Segmenter
	for _, _aeg := range _cee.Segments {
		switch _aeg.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_aebae, _fbae = _aeg.GetSegmentData()
			if _fbae != nil {
				return nil, _fb.Wrap(_fbae, _cgdg, "")
			}
			_ccff = append(_ccff, _aebae)
		case 50:
			_aebae, _fbae = _aeg.GetSegmentData()
			if _fbae != nil {
				return nil, _fbae
			}
			_dffb, _dad := _aebae.(*_ea.EndOfStripe)
			if !_dad {
				return nil, _fb.Errorf(_cgdg, "\u0045\u006e\u0064\u004f\u0066\u0053\u0074\u0072\u0069\u0070\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u006f\u0066\u0020\u0076\u0061l\u0069\u0064\u0020\u0074\u0079p\u0065\u003a \u0027\u0025\u0054\u0027", _aebae)
			}
			_ccff = append(_ccff, _dffb)
			_cee.FinalHeight = _dffb.LineNumber()
		}
	}
	return _ccff, nil
}
func (_ffb *Document) nextSegmentNumber() uint32 {
	_ddg := _ffb.CurrentSegmentNumber
	_ffb.CurrentSegmentNumber++
	return _ddg
}
func DecodeDocument(input _f.StreamReader, globals *Globals) (*Document, error) {
	return _gag(input, globals)
}

const (
	GenericEM EncodingMethod = iota
	CorrelationEM
	RankHausEM
)

func (_edcd *Page) createStripedPage(_gcf *_ea.PageInformationSegment) error {
	const _fac = "\u0063\u0072\u0065\u0061\u0074\u0065\u0053\u0074\u0072\u0069\u0070\u0065d\u0050\u0061\u0067\u0065"
	_cgdc, _debc := _edcd.collectPageStripes()
	if _debc != nil {
		return _fb.Wrap(_debc, _fac, "")
	}
	var _eed int
	for _, _cec := range _cgdc {
		if _bdc, _ffbd := _cec.(*_ea.EndOfStripe); _ffbd {
			_eed = _bdc.LineNumber() + 1
		} else {
			_gegg := _cec.(_ea.Regioner)
			_gfc := _gegg.GetRegionInfo()
			_ded := _edcd.getCombinationOperator(_gcf, _gfc.CombinaionOperator)
			_gdf, _bgfea := _gegg.GetRegionBitmap()
			if _bgfea != nil {
				return _fb.Wrap(_bgfea, _fac, "")
			}
			_bgfea = _ca.Blit(_gdf, _edcd.Bitmap, int(_gfc.XLocation), _eed, _ded)
			if _bgfea != nil {
				return _fb.Wrap(_bgfea, _fac, "")
			}
		}
	}
	return nil
}
func (_ade *Document) completeSymbols() (_ggda error) {
	const _cbac = "\u0063o\u006dp\u006c\u0065\u0074\u0065\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	if _ade.Classer == nil {
		return nil
	}
	if _ade.Classer.UndilatedTemplates == nil {
		return _fb.Error(_cbac, "\u006e\u006f t\u0065\u006d\u0070l\u0061\u0074\u0065\u0073 de\u0066in\u0065\u0064\u0020\u0066\u006f\u0072\u0020th\u0065\u0020\u0063\u006c\u0061\u0073\u0073e\u0072")
	}
	_be := len(_ade.Pages) == 1
	_bdf := make([]int, _ade.Classer.UndilatedTemplates.Size())
	var _ab int
	for _ggg := 0; _ggg < _ade.Classer.ClassIDs.Size(); _ggg++ {
		_ab, _ggda = _ade.Classer.ClassIDs.Get(_ggg)
		if _ggda != nil {
			return _fb.Wrap(_ggda, _cbac, "\u0063\u006c\u0061\u0073\u0073\u0020\u0049\u0044\u0027\u0073")
		}
		_bdf[_ab]++
	}
	var _ffa []int
	for _fgb := 0; _fgb < _ade.Classer.UndilatedTemplates.Size(); _fgb++ {
		if _bdf[_fgb] == 0 {
			return _fb.Error(_cbac, "\u006eo\u0020\u0073y\u006d\u0062\u006f\u006cs\u0020\u0069\u006es\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006fun\u0064\u0020\u0066o\u0072\u0020g\u0069\u0076\u0065\u006e\u0020\u0063l\u0061\u0073s\u003f\u0020")
		}
		if _bdf[_fgb] > 1 || _be {
			_ffa = append(_ffa, _fgb)
		}
	}
	_ade._bd = len(_ffa)
	var _dag, _ffd int
	for _ba := 0; _ba < _ade.Classer.ComponentPageNumbers.Size(); _ba++ {
		_dag, _ggda = _ade.Classer.ComponentPageNumbers.Get(_ba)
		if _ggda != nil {
			return _fb.Wrapf(_ggda, _cbac, "p\u0061\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027 \u006e\u006f\u0074\u0020\u0066\u006f\u0075nd\u0020\u0069\u006e\u0020t\u0068\u0065\u0020\u0063\u006c\u0061\u0073\u0073\u0065r \u0070\u0061g\u0065\u006e\u0075\u006d\u0062\u0065\u0072\u0073", _ba)
		}
		_ffd, _ggda = _ade.Classer.ClassIDs.Get(_ba)
		if _ggda != nil {
			return _fb.Wrapf(_ggda, _cbac, "\u0063\u0061\u006e\u0027\u0074\u0020\u0067e\u0074\u0020\u0073y\u006d\u0062\u006f\u006c \u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0027\u0025\u0064\u0027\u0020\u0066\u0072\u006f\u006d\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072", _dag)
		}
		if _bdf[_ffd] == 1 && !_be {
			_ade._ef[_dag] = append(_ade._ef[_dag], _ffd)
		}
	}
	if _ggda = _ade.Classer.ComputeLLCorners(); _ggda != nil {
		return _fb.Wrap(_ggda, _cbac, "")
	}
	return nil
}
func (_baa *Document) encodeFileHeader(_bgf _f.BinaryWriter) (_gacf int, _faf error) {
	const _fccc = "\u0065\u006ec\u006f\u0064\u0065F\u0069\u006c\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_gacf, _faf = _bgf.Write(_gc)
	if _faf != nil {
		return _gacf, _fb.Wrap(_faf, _fccc, "\u0069\u0064")
	}
	if _faf = _bgf.WriteByte(0x01); _faf != nil {
		return _gacf, _fb.Wrap(_faf, _fccc, "\u0066\u006c\u0061g\u0073")
	}
	_gacf++
	_bbg := make([]byte, 4)
	_g.BigEndian.PutUint32(_bbg, _baa.NumberOfPages)
	_bed, _faf := _bgf.Write(_bbg)
	if _faf != nil {
		return _bed, _fb.Wrap(_faf, _fccc, "p\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_gacf += _bed
	return _gacf, nil
}
func (_ggd *Document) produceClassifiedPages() (_fg error) {
	const _ce = "\u0070\u0072\u006f\u0064uc\u0065\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0050\u0061\u0067e\u0073"
	if _ggd.Classer == nil {
		return nil
	}
	var (
		_db  *Page
		_dba bool
		_cfa *_ea.Header
	)
	for _fgc := 1; _fgc <= int(_ggd.NumberOfPages); _fgc++ {
		if _db, _dba = _ggd.Pages[_fgc]; !_dba {
			return _fb.Errorf(_ce, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _fgc)
		}
		if _db.EncodingMethod == GenericEM {
			continue
		}
		if _cfa == nil {
			if _cfa, _fg = _ggd.GlobalSegments.GetSymbolDictionary(); _fg != nil {
				return _fb.Wrap(_fg, _ce, "")
			}
		}
		if _fg = _ggd.produceClassifiedPage(_db, _cfa); _fg != nil {
			return _fb.Wrapf(_fg, _ce, "\u0070\u0061\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _fgc)
		}
	}
	return nil
}
func (_agf *Page) addTextRegionSegment(_ggdd []*_ea.Header, _aad, _gfe map[int]int, _acc []int, _bcecf *_ca.Points, _dbag *_ca.Bitmaps, _abfb *_da.IntSlice, _dbfa *_ca.Boxes, _fdc, _eceg int) {
	_bgg := &_ea.TextRegion{NumberOfSymbols: uint32(_eceg)}
	_bgg.InitEncode(_aad, _gfe, _acc, _bcecf, _dbag, _abfb, _dbfa, _agf.FinalWidth, _agf.FinalHeight, _fdc)
	_bbc := &_ea.Header{RTSegments: _ggdd, SegmentData: _bgg, PageAssociation: _agf.PageNumber, Type: _ea.TImmediateTextRegion}
	_eac := _ea.TPageInformation
	if _gfe != nil {
		_eac = _ea.TSymbolDictionary
	}
	var _ceb int
	for ; _ceb < len(_agf.Segments); _ceb++ {
		if _agf.Segments[_ceb].Type == _eac {
			_ceb++
			break
		}
	}
	_agf.Segments = append(_agf.Segments, nil)
	copy(_agf.Segments[_ceb+1:], _agf.Segments[_ceb:])
	_agf.Segments[_ceb] = _bbc
}

type Page struct {
	Segments           []*_ea.Header
	PageNumber         int
	Bitmap             *_ca.Bitmap
	FinalHeight        int
	FinalWidth         int
	ResolutionX        int
	ResolutionY        int
	IsLossless         bool
	Document           *Document
	FirstSegmentNumber int
	EncodingMethod     EncodingMethod
	BlackIsOne         bool
}

func (_aee *Page) GetSegment(number int) (*_ea.Header, error) {
	const _deec = "\u0050a\u0067e\u002e\u0047\u0065\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074"
	for _, _edc := range _aee.Segments {
		if _edc.SegmentNumber == uint32(number) {
			return _edc, nil
		}
	}
	_bgbb := make([]uint32, len(_aee.Segments))
	for _edbg, _gbe := range _aee.Segments {
		_bgbb[_edbg] = _gbe.SegmentNumber
	}
	return nil, _fb.Errorf(_deec, "\u0073e\u0067\u006d\u0065n\u0074\u0020\u0077i\u0074h \u006e\u0075\u006d\u0062\u0065\u0072\u003a \u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0067\u0065\u003a\u0020'%\u0064'\u002e\u0020\u004b\u006e\u006f\u0077n\u0020\u0073\u0065\u0067\u006de\u006e\u0074\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0073\u003a \u0025\u0076", number, _aee.PageNumber, _bgbb)
}
func (_bfc *Page) getCombinationOperator(_fgfd *_ea.PageInformationSegment, _bbge _ca.CombinationOperator) _ca.CombinationOperator {
	if _fgfd.CombinationOperatorOverrideAllowed() {
		return _bbge
	}
	return _fgfd.CombinationOperator()
}

type Globals struct {
	Segments []*_ea.Header
}

func (_ccaf *Page) nextSegmentNumber() uint32 { return _ccaf.Document.nextSegmentNumber() }
func (_ccf *Page) AddEndOfPageSegment() {
	_cdf := &_ea.Header{Type: _ea.TEndOfPage, PageAssociation: _ccf.PageNumber}
	_ccf.Segments = append(_ccf.Segments, _cdf)
}
func (_ed *Document) GetGlobalSegment(i int) (*_ea.Header, error) {
	_bcd, _gfg := _ed.GlobalSegments.GetSegment(i)
	if _gfg != nil {
		return nil, _fb.Wrap(_gfg, "\u0047\u0065t\u0047\u006c\u006fb\u0061\u006c\u0053\u0065\u0067\u006d\u0065\u006e\u0074", "")
	}
	return _bcd, nil
}
func _gag(_eca _f.StreamReader, _cgfb *Globals) (*Document, error) {
	_fbee := &Document{Pages: make(map[int]*Page), InputStream: _eca, OrganizationType: _ea.OSequential, NumberOfPagesUnknown: true, GlobalSegments: _cgfb, _a: 9}
	if _fbee.GlobalSegments == nil {
		_fbee.GlobalSegments = &Globals{}
	}
	if _ccdg := _fbee.mapData(); _ccdg != nil {
		return nil, _ccdg
	}
	return _fbee, nil
}
func _bgge(_fab *Document, _cbge int) *Page {
	return &Page{Document: _fab, PageNumber: _cbge, Segments: []*_ea.Header{}}
}
func (_gb *Document) isFileHeaderPresent() (bool, error) {
	_gb.InputStream.Mark()
	for _, _cfb := range _gc {
		_aeba, _ecb := _gb.InputStream.ReadByte()
		if _ecb != nil {
			return false, _ecb
		}
		if _cfb != _aeba {
			_gb.InputStream.Reset()
			return false, nil
		}
	}
	_gb.InputStream.Reset()
	return true, nil
}
func (_gf *Document) AddGenericPage(bm *_ca.Bitmap, duplicateLineRemoval bool) (_dfb error) {
	const _eb = "\u0044\u006f\u0063um\u0065\u006e\u0074\u002e\u0041\u0064\u0064\u0047\u0065\u006e\u0065\u0072\u0069\u0063\u0050\u0061\u0067\u0065"
	if !_gf.FullHeaders && _gf.NumberOfPages != 0 {
		return _fb.Error(_eb, "\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0061\u006c\u0072\u0065a\u0064\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0070\u0061\u0067\u0065\u002e\u0020\u0046\u0069\u006c\u0065\u004d\u006f\u0064\u0065\u0020\u0064\u0069\u0073\u0061\u006c\u006c\u006f\u0077\u0073\u0020\u0061\u0064\u0064i\u006e\u0067\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e \u006f\u006e\u0065\u0020\u0070\u0061g\u0065")
	}
	_bb := &Page{Segments: []*_ea.Header{}, Bitmap: bm, Document: _gf, FinalHeight: bm.Height, FinalWidth: bm.Width, IsLossless: true, BlackIsOne: bm.Color == _ca.Chocolate}
	_bb.PageNumber = int(_gf.nextPageNumber())
	_gf.Pages[_bb.PageNumber] = _bb
	bm.InverseData()
	_bb.AddPageInformationSegment()
	if _dfb = _bb.AddGenericRegion(bm, 0, 0, 0, _ea.TImmediateGenericRegion, duplicateLineRemoval); _dfb != nil {
		return _fb.Wrap(_dfb, _eb, "")
	}
	if _gf.FullHeaders {
		_bb.AddEndOfPageSegment()
	}
	return nil
}
func (_fbeg *Page) getResolutionX() (int, error) {
	const _bca = "\u0067\u0065\u0074\u0052\u0065\u0073\u006f\u006c\u0075t\u0069\u006f\u006e\u0058"
	if _fbeg.ResolutionX != 0 {
		return _fbeg.ResolutionX, nil
	}
	_cca := _fbeg.getPageInformationSegment()
	if _cca == nil {
		return 0, _fb.Error(_bca, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_acg, _fge := _cca.GetSegmentData()
	if _fge != nil {
		return 0, _fb.Wrap(_fge, _bca, "")
	}
	_bga, _cecf := _acg.(*_ea.PageInformationSegment)
	if !_cecf {
		return 0, _fb.Errorf(_bca, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _acg)
	}
	_fbeg.ResolutionX = _bga.ResolutionX
	return _fbeg.ResolutionX, nil
}
func (_eaae *Page) createPage(_bde *_ea.PageInformationSegment) error {
	var _adee error
	if !_bde.IsStripe || _bde.PageBMHeight != -1 {
		_adee = _eaae.createNormalPage(_bde)
	} else {
		_adee = _eaae.createStripedPage(_bde)
	}
	return _adee
}
func (_cgfa *Page) fitsPage(_cfd *_ea.PageInformationSegment, _gee *_ca.Bitmap) bool {
	return _cgfa.countRegions() == 1 && _cfd.DefaultPixelValue == 0 && _cfd.PageBMWidth == _gee.Width && _cfd.PageBMHeight == _gee.Height
}
func (_gac *Document) Encode() (_cbae []byte, _ege error) {
	const _cg = "\u0044o\u0063u\u006d\u0065\u006e\u0074\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	var _adcc, _dd int
	if _gac.FullHeaders {
		if _adcc, _ege = _gac.encodeFileHeader(_gac._b); _ege != nil {
			return nil, _fb.Wrap(_ege, _cg, "")
		}
	}
	var (
		_cgf bool
		_ced *_ea.Header
		_af  *Page
	)
	if _ege = _gac.completeClassifiedPages(); _ege != nil {
		return nil, _fb.Wrap(_ege, _cg, "")
	}
	if _ege = _gac.produceClassifiedPages(); _ege != nil {
		return nil, _fb.Wrap(_ege, _cg, "")
	}
	if _gac.GlobalSegments != nil {
		for _, _ced = range _gac.GlobalSegments.Segments {
			if _ege = _gac.encodeSegment(_ced, &_adcc); _ege != nil {
				return nil, _fb.Wrap(_ege, _cg, "")
			}
		}
	}
	for _egf := 1; _egf <= int(_gac.NumberOfPages); _egf++ {
		if _af, _cgf = _gac.Pages[_egf]; !_cgf {
			return nil, _fb.Errorf(_cg, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _egf)
		}
		for _, _ced = range _af.Segments {
			if _ege = _gac.encodeSegment(_ced, &_adcc); _ege != nil {
				return nil, _fb.Wrap(_ege, _cg, "")
			}
		}
	}
	if _gac.FullHeaders {
		if _dd, _ege = _gac.encodeEOFHeader(_gac._b); _ege != nil {
			return nil, _fb.Wrap(_ege, _cg, "")
		}
		_adcc += _dd
	}
	_cbae = _gac._b.Data()
	if len(_cbae) != _adcc {
		_dff.Log.Debug("\u0042\u0079\u0074\u0065\u0073 \u0077\u0072\u0069\u0074\u0074\u0065\u006e \u0028\u006e\u0029\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0064\u0061\u0074\u0061\u0020\u0065\u006e\u0063\u006fd\u0065\u0064\u003a\u0020\u0027\u0025d\u0027", _adcc, len(_cbae))
	}
	return _cbae, nil
}
func (_cbgc *Globals) GetSegment(segmentNumber int) (*_ea.Header, error) {
	const _cgb = "\u0047l\u006fb\u0061\u006c\u0073\u002e\u0047e\u0074\u0053e\u0067\u006d\u0065\u006e\u0074"
	if _cbgc == nil {
		return nil, _fb.Error(_cgb, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_cbgc.Segments) == 0 {
		return nil, _fb.Error(_cgb, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	var _bec *_ea.Header
	for _, _bec = range _cbgc.Segments {
		if _bec.SegmentNumber == uint32(segmentNumber) {
			break
		}
	}
	if _bec == nil {
		return nil, _fb.Error(_cgb, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	return _bec, nil
}
func (_abc *Page) getResolutionY() (int, error) {
	const _fad = "\u0067\u0065\u0074\u0052\u0065\u0073\u006f\u006c\u0075t\u0069\u006f\u006e\u0059"
	if _abc.ResolutionY != 0 {
		return _abc.ResolutionY, nil
	}
	_gbg := _abc.getPageInformationSegment()
	if _gbg == nil {
		return 0, _fb.Error(_fad, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_bcgc, _aba := _gbg.GetSegmentData()
	if _aba != nil {
		return 0, _fb.Wrap(_aba, _fad, "")
	}
	_ccgf, _eedc := _bcgc.(*_ea.PageInformationSegment)
	if !_eedc {
		return 0, _fb.Errorf(_fad, "\u0070\u0061\u0067\u0065\u0020\u0069\u006e\u0066o\u0072\u006d\u0061ti\u006f\u006e\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0074\u0079\u0070\u0065\u003a\u0027%\u0054\u0027", _bcgc)
	}
	_abc.ResolutionY = _ccgf.ResolutionY
	return _abc.ResolutionY, nil
}
func (_cfe *Page) getWidth() (int, error) {
	const _bda = "\u0067\u0065\u0074\u0057\u0069\u0064\u0074\u0068"
	if _cfe.FinalWidth != 0 {
		return _cfe.FinalWidth, nil
	}
	_dge := _cfe.getPageInformationSegment()
	if _dge == nil {
		return 0, _fb.Error(_bda, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_ddf, _faca := _dge.GetSegmentData()
	if _faca != nil {
		return 0, _fb.Wrap(_faca, _bda, "")
	}
	_afab, _fcb := _ddf.(*_ea.PageInformationSegment)
	if !_fcb {
		return 0, _fb.Errorf(_bda, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _ddf)
	}
	_cfe.FinalWidth = _afab.PageBMWidth
	return _cfe.FinalWidth, nil
}
func (_fa *Document) AddClassifiedPage(bm *_ca.Bitmap, method _e.Method) (_cb error) {
	const _bc = "\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0041\u0064d\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064P\u0061\u0067\u0065"
	if !_fa.FullHeaders && _fa.NumberOfPages != 0 {
		return _fb.Error(_bc, "\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0061\u006c\u0072\u0065a\u0064\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0070\u0061\u0067\u0065\u002e\u0020\u0046\u0069\u006c\u0065\u004d\u006f\u0064\u0065\u0020\u0064\u0069\u0073\u0061\u006c\u006c\u006f\u0077\u0073\u0020\u0061\u0064\u0064i\u006e\u0067\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e \u006f\u006e\u0065\u0020\u0070\u0061g\u0065")
	}
	if _fa.Classer == nil {
		if _fa.Classer, _cb = _e.Init(_e.DefaultSettings()); _cb != nil {
			return _fb.Wrap(_cb, _bc, "")
		}
	}
	_cf := int(_fa.nextPageNumber())
	_eag := &Page{Segments: []*_ea.Header{}, Bitmap: bm, Document: _fa, FinalHeight: bm.Height, FinalWidth: bm.Width, PageNumber: _cf}
	_fa.Pages[_cf] = _eag
	switch method {
	case _e.RankHaus:
		_eag.EncodingMethod = RankHausEM
	case _e.Correlation:
		_eag.EncodingMethod = CorrelationEM
	}
	_eag.AddPageInformationSegment()
	if _cb = _fa.Classer.AddPage(bm, _cf, method); _cb != nil {
		return _fb.Wrap(_cb, _bc, "")
	}
	if _fa.FullHeaders {
		_eag.AddEndOfPageSegment()
	}
	return nil
}
func (_bfa *Page) String() string {
	return _dg.Sprintf("\u0050\u0061\u0067\u0065\u0020\u0023\u0025\u0064", _bfa.PageNumber)
}
func (_dcbe *Page) GetBitmap() (_bfe *_ca.Bitmap, _edb error) {
	_dff.Log.Trace(_dg.Sprintf("\u005b\u0050\u0041G\u0045\u005d\u005b\u0023%\u0064\u005d\u0020\u0047\u0065\u0074\u0042i\u0074\u006d\u0061\u0070\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e", _dcbe.PageNumber))
	defer func() {
		if _edb != nil {
			_dff.Log.Trace(_dg.Sprintf("\u005b\u0050\u0041\u0047\u0045\u005d\u005b\u0023\u0025\u0064\u005d\u0020\u0047\u0065\u0074B\u0069t\u006d\u0061\u0070\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _dcbe.PageNumber, _edb))
		} else {
			_dff.Log.Trace(_dg.Sprintf("\u005b\u0050\u0041\u0047\u0045\u005d\u005b\u0023\u0025\u0064]\u0020\u0047\u0065\u0074\u0042\u0069\u0074m\u0061\u0070\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064", _dcbe.PageNumber))
		}
	}()
	if _dcbe.Bitmap != nil {
		return _dcbe.Bitmap, nil
	}
	_edb = _dcbe.composePageBitmap()
	if _edb != nil {
		return nil, _edb
	}
	return _dcbe.Bitmap, nil
}
func (_adg *Document) determineRandomDataOffsets(_efff []*_ea.Header, _eaa uint64) {
	if _adg.OrganizationType != _ea.ORandom {
		return
	}
	for _, _fea := range _efff {
		_fea.SegmentDataStartOffset = _eaa
		_eaa += _fea.SegmentDataLength
	}
}
func (_acf *Page) clearSegmentData() {
	for _gfaf := range _acf.Segments {
		_acf.Segments[_gfaf].CleanSegmentData()
	}
}
func (_aae *Document) GetNumberOfPages() (uint32, error) {
	if _aae.NumberOfPagesUnknown || _aae.NumberOfPages == 0 {
		if len(_aae.Pages) == 0 {
			if _edg := _aae.mapData(); _edg != nil {
				return 0, _fb.Wrap(_edg, "\u0044o\u0063\u0075\u006d\u0065n\u0074\u002e\u0047\u0065\u0074N\u0075m\u0062e\u0072\u004f\u0066\u0050\u0061\u0067\u0065s", "")
			}
		}
		return uint32(len(_aae.Pages)), nil
	}
	return _aae.NumberOfPages, nil
}
func (_cge *Page) AddPageInformationSegment() {
	_gef := &_ea.PageInformationSegment{PageBMWidth: _cge.FinalWidth, PageBMHeight: _cge.FinalHeight, ResolutionX: _cge.ResolutionX, ResolutionY: _cge.ResolutionY, IsLossless: _cge.IsLossless}
	if _cge.BlackIsOne {
		_gef.DefaultPixelValue = uint8(0x1)
	}
	_gdg := &_ea.Header{PageAssociation: _cge.PageNumber, SegmentDataLength: uint64(_gef.Size()), SegmentData: _gef, Type: _ea.TPageInformation}
	_cge.Segments = append(_cge.Segments, _gdg)
}
func (_deb *Globals) AddSegment(segment *_ea.Header) { _deb.Segments = append(_deb.Segments, segment) }
func (_fde *Document) encodeEOFHeader(_cgc _f.BinaryWriter) (_gfgc int, _gce error) {
	_fbde := &_ea.Header{SegmentNumber: _fde.nextSegmentNumber(), Type: _ea.TEndOfFile}
	if _gfgc, _gce = _fbde.Encode(_cgc); _gce != nil {
		return 0, _fb.Wrap(_gce, "\u0065n\u0063o\u0064\u0065\u0045\u004f\u0046\u0048\u0065\u0061\u0064\u0065\u0072", "")
	}
	return _gfgc, nil
}

type Document struct {
	Pages                map[int]*Page
	NumberOfPagesUnknown bool
	NumberOfPages        uint32
	GBUseExtTemplate     bool
	InputStream          _f.StreamReader
	GlobalSegments       *Globals
	OrganizationType     _ea.OrganizationType
	Classer              *_e.Classer
	XRes, YRes           int
	FullHeaders          bool
	CurrentSegmentNumber uint32
	AverageTemplates     *_ca.Bitmaps
	BaseIndexes          []int
	Refinement           bool
	RefineLevel          int
	_a                   uint8
	_b                   *_f.BufferedWriter
	EncodeGlobals        bool
	_bd                  int
	_ef                  map[int][]int
	_ag                  map[int][]int
	_dffa                []int
	_eg                  map[int]int
}

func (_afaa *Page) GetHeight() (int, error) { return _afaa.getHeight() }
func (_bgb *Page) AddGenericRegion(bm *_ca.Bitmap, xloc, yloc, template int, tp _ea.Type, duplicateLineRemoval bool) error {
	const _bcb = "P\u0061\u0067\u0065\u002eAd\u0064G\u0065\u006e\u0065\u0072\u0069c\u0052\u0065\u0067\u0069\u006f\u006e"
	_abb := &_ea.GenericRegion{}
	if _bcec := _abb.InitEncode(bm, xloc, yloc, template, duplicateLineRemoval); _bcec != nil {
		return _fb.Wrap(_bcec, _bcb, "")
	}
	_ecaa := &_ea.Header{Type: _ea.TImmediateGenericRegion, PageAssociation: _bgb.PageNumber, SegmentData: _abb}
	_bgb.Segments = append(_bgb.Segments, _ecaa)
	return nil
}
func (_cfc *Document) completeClassifiedPages() (_gfa error) {
	const _bba = "\u0063\u006f\u006dpl\u0065\u0074\u0065\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0050\u0061\u0067\u0065\u0073"
	if _cfc.Classer == nil {
		return nil
	}
	_cfc._dffa = make([]int, _cfc.Classer.UndilatedTemplates.Size())
	for _ga := 0; _ga < _cfc.Classer.ClassIDs.Size(); _ga++ {
		_gg, _dgf := _cfc.Classer.ClassIDs.Get(_ga)
		if _dgf != nil {
			return _fb.Wrapf(_dgf, _bba, "\u0063\u006c\u0061\u0073s \u0077\u0069\u0074\u0068\u0020\u0069\u0064\u003a\u0020\u0027\u0025\u0064\u0027", _ga)
		}
		_cfc._dffa[_gg]++
	}
	var _cbg []int
	for _bg := 0; _bg < _cfc.Classer.UndilatedTemplates.Size(); _bg++ {
		if _cfc.NumberOfPages == 1 || _cfc._dffa[_bg] > 1 {
			_cbg = append(_cbg, _bg)
		}
	}
	var (
		_fe *Page
		_cc bool
	)
	for _fbe, _ge := range *_cfc.Classer.ComponentPageNumbers {
		if _fe, _cc = _cfc.Pages[_ge]; !_cc {
			return _fb.Errorf(_bba, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _fbe)
		}
		if _fe.EncodingMethod == GenericEM {
			_dff.Log.Error("\u0047\u0065\u006e\u0065\u0072\u0069c\u0020\u0070\u0061g\u0065\u0020\u0077i\u0074\u0068\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u003a \u0027\u0025\u0064\u0027\u0020ma\u0070\u0070\u0065\u0064\u0020\u0061\u0073\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0020\u0070\u0061\u0067\u0065", _fbe)
			continue
		}
		_cfc._ag[_ge] = append(_cfc._ag[_ge], _fbe)
		_gcg, _cbe := _cfc.Classer.ClassIDs.Get(_fbe)
		if _cbe != nil {
			return _fb.Wrapf(_cbe, _bba, "\u006e\u006f\u0020\u0073uc\u0068\u0020\u0063\u006c\u0061\u0073\u0073\u0049\u0044\u003a\u0020\u0025\u0064", _fbe)
		}
		if _cfc._dffa[_gcg] == 1 && _cfc.NumberOfPages != 1 {
			_dgd := append(_cfc._ef[_ge], _gcg)
			_cfc._ef[_ge] = _dgd
		}
	}
	if _gfa = _cfc.Classer.ComputeLLCorners(); _gfa != nil {
		return _fb.Wrap(_gfa, _bba, "")
	}
	if _, _gfa = _cfc.addSymbolDictionary(0, _cfc.Classer.UndilatedTemplates, _cbg, _cfc._eg, false); _gfa != nil {
		return _fb.Wrap(_gfa, _bba, "")
	}
	return nil
}
func (_eeb *Page) GetWidth() (int, error) { return _eeb.getWidth() }
func (_fee *Page) getPageInformationSegment() *_ea.Header {
	for _, _gaff := range _fee.Segments {
		if _gaff.Type == _ea.TPageInformation {
			return _gaff
		}
	}
	_dff.Log.Debug("\u0050\u0061\u0067\u0065\u0020\u0069\u006e\u0066o\u0072\u006d\u0061ti\u006f\u006e\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0066o\u0072\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0073\u002e", _fee)
	return nil
}
func InitEncodeDocument(fullHeaders bool) *Document {
	return &Document{FullHeaders: fullHeaders, _b: _f.BufferedMSB(), Pages: map[int]*Page{}, _ef: map[int][]int{}, _eg: map[int]int{}, _ag: map[int][]int{}}
}
func (_daf *Document) produceClassifiedPage(_ff *Page, _gd *_ea.Header) (_bf error) {
	const _aa = "p\u0072\u006f\u0064\u0075ce\u0043l\u0061\u0073\u0073\u0069\u0066i\u0065\u0064\u0050\u0061\u0067\u0065"
	var _gab map[int]int
	_fc := _daf._bd
	_dcb := []*_ea.Header{_gd}
	if len(_daf._ef[_ff.PageNumber]) > 0 {
		_gab = map[int]int{}
		_ae, _eff := _daf.addSymbolDictionary(_ff.PageNumber, _daf.Classer.UndilatedTemplates, _daf._ef[_ff.PageNumber], _gab, false)
		if _eff != nil {
			return _fb.Wrap(_eff, _aa, "")
		}
		_dcb = append(_dcb, _ae)
		_fc += len(_daf._ef[_ff.PageNumber])
	}
	_bbaa := _daf._ag[_ff.PageNumber]
	_dff.Log.Debug("P\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020c\u006f\u006d\u0070\u0073: \u0025\u0076", _ff.PageNumber, _bbaa)
	_ff.addTextRegionSegment(_dcb, _daf._eg, _gab, _daf._ag[_ff.PageNumber], _daf.Classer.PtaLL, _daf.Classer.UndilatedTemplates, _daf.Classer.ClassIDs, nil, _gda(_fc), len(_daf._ag[_ff.PageNumber]))
	return nil
}
func (_cbef *Page) lastSegmentNumber() (_gfea uint32, _deg error) {
	const _gdd = "\u006c\u0061\u0073\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u004eu\u006d\u0062\u0065\u0072"
	if len(_cbef.Segments) == 0 {
		return _gfea, _fb.Errorf(_gdd, "\u006e\u006f\u0020se\u0067\u006d\u0065\u006e\u0074\u0073\u0020\u0066\u006fu\u006ed\u0020i\u006e \u0074\u0068\u0065\u0020\u0070\u0061\u0067\u0065\u0020\u0027\u0025\u0064\u0027", _cbef.PageNumber)
	}
	return _cbef.Segments[len(_cbef.Segments)-1].SegmentNumber, nil
}
func (_fbd *Document) encodeSegment(_ggf *_ea.Header, _ged *int) error {
	const _dfbf = "\u0065\u006e\u0063\u006f\u0064\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074"
	_ggf.SegmentNumber = _fbd.nextSegmentNumber()
	_ebf, _aeb := _ggf.Encode(_fbd._b)
	if _aeb != nil {
		return _fb.Wrapf(_aeb, _dfbf, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", _ggf.SegmentNumber)
	}
	*_ged += _ebf
	return nil
}
func (_fddc *Page) GetResolutionY() (int, error) { return _fddc.getResolutionY() }
func (_dfe *Document) GetPage(pageNumber int) (_ea.Pager, error) {
	const _adf = "\u0044\u006fc\u0075\u006d\u0065n\u0074\u002e\u0047\u0065\u0074\u0050\u0061\u0067\u0065"
	if pageNumber < 0 {
		_dff.Log.Debug("\u004a\u0042\u0049\u00472\u0020\u0050\u0061\u0067\u0065\u0020\u002d\u0020\u0047e\u0074\u0050\u0061\u0067\u0065\u003a\u0020\u0025\u0064\u002e\u0020\u0050\u0061\u0067\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062e\u0020\u006c\u006f\u0077\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030\u002e\u0020\u0025\u0073", pageNumber, _df.Stack())
		return nil, _fb.Errorf(_adf, "\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032\u0020d\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u002d\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064 \u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u003a\u0020\u0025\u0064", pageNumber)
	}
	if pageNumber > len(_dfe.Pages) {
		_dff.Log.Debug("\u0050\u0061\u0067\u0065 n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0064\u002e\u0020%\u0073", pageNumber, _df.Stack())
		return nil, _fb.Error(_adf, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u002d\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_fd, _gcgb := _dfe.Pages[pageNumber]
	if !_gcgb {
		_dff.Log.Debug("\u0050\u0061\u0067\u0065 n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0064\u002e\u0020%\u0073", pageNumber, _df.Stack())
		return nil, _fb.Errorf(_adf, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u002d\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	return _fd, nil
}
func (_bdbe *Globals) GetSymbolDictionary() (*_ea.Header, error) {
	const _ggdf = "G\u006c\u006f\u0062\u0061\u006c\u0073.\u0047\u0065\u0074\u0053\u0079\u006d\u0062\u006f\u006cD\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079"
	if _bdbe == nil {
		return nil, _fb.Error(_ggdf, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_bdbe.Segments) == 0 {
		return nil, _fb.Error(_ggdf, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	for _, _aaeb := range _bdbe.Segments {
		if _aaeb.Type == _ea.TSymbolDictionary {
			return _aaeb, nil
		}
	}
	return nil, _fb.Error(_ggdf, "\u0067\u006c\u006fba\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0020d\u0069c\u0074i\u006fn\u0061\u0072\u0079\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}
func (_bcg *Page) getHeight() (int, error) {
	const _cedc = "\u0067e\u0074\u0048\u0065\u0069\u0067\u0068t"
	if _bcg.FinalHeight != 0 {
		return _bcg.FinalHeight, nil
	}
	_cga := _bcg.getPageInformationSegment()
	if _cga == nil {
		return 0, _fb.Error(_cedc, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_aec, _ggb := _cga.GetSegmentData()
	if _ggb != nil {
		return 0, _fb.Wrap(_ggb, _cedc, "")
	}
	_agc, _bcgf := _aec.(*_ea.PageInformationSegment)
	if !_bcgf {
		return 0, _fb.Errorf(_cedc, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _aec)
	}
	if _agc.PageBMHeight == _dc.MaxInt32 {
		_, _ggb = _bcg.GetBitmap()
		if _ggb != nil {
			return 0, _fb.Wrap(_ggb, _cedc, "")
		}
	} else {
		_bcg.FinalHeight = _agc.PageBMHeight
	}
	return _bcg.FinalHeight, nil
}
func (_eagb *Globals) GetSegmentByIndex(index int) (*_ea.Header, error) {
	const _fafb = "\u0047l\u006f\u0062\u0061\u006cs\u002e\u0047\u0065\u0074\u0053e\u0067m\u0065n\u0074\u0042\u0079\u0049\u006e\u0064\u0065x"
	if _eagb == nil {
		return nil, _fb.Error(_fafb, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_eagb.Segments) == 0 {
		return nil, _fb.Error(_fafb, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	if index > len(_eagb.Segments)-1 {
		return nil, _fb.Error(_fafb, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	return _eagb.Segments[index], nil
}
