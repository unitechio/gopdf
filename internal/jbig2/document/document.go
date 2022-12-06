package document

import (
	_c "encoding/binary"
	_cc "fmt"
	_ea "io"
	_ec "math"
	_ccb "runtime/debug"

	_f "bitbucket.org/shenghui0779/gopdf/common"
	_g "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_a "bitbucket.org/shenghui0779/gopdf/internal/jbig2/basic"
	_ee "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_ecd "bitbucket.org/shenghui0779/gopdf/internal/jbig2/document/segments"
	_d "bitbucket.org/shenghui0779/gopdf/internal/jbig2/encoder/classer"
	_ab "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

func (_gdge *Document) determineRandomDataOffsets(_cce []*_ecd.Header, _bfc uint64) {
	if _gdge.OrganizationType != _ecd.ORandom {
		return
	}
	for _, _fabb := range _cce {
		_fabb.SegmentDataStartOffset = _bfc
		_bfc += _fabb.SegmentDataLength
	}
}

type Globals struct{ Segments []*_ecd.Header }

func (_cag *Page) String() string {
	return _cc.Sprintf("\u0050\u0061\u0067\u0065\u0020\u0023\u0025\u0064", _cag.PageNumber)
}
func (_gda *Document) produceClassifiedPage(_ccg *Page, _de *_ecd.Header) (_aba error) {
	const _ebd = "p\u0072\u006f\u0064\u0075ce\u0043l\u0061\u0073\u0073\u0069\u0066i\u0065\u0064\u0050\u0061\u0067\u0065"
	var _edb map[int]int
	_gb := _gda._db
	_ddf := []*_ecd.Header{_de}
	if len(_gda._dc[_ccg.PageNumber]) > 0 {
		_edb = map[int]int{}
		_eed, _ae := _gda.addSymbolDictionary(_ccg.PageNumber, _gda.Classer.UndilatedTemplates, _gda._dc[_ccg.PageNumber], _edb, false)
		if _ae != nil {
			return _ab.Wrap(_ae, _ebd, "")
		}
		_ddf = append(_ddf, _eed)
		_gb += len(_gda._dc[_ccg.PageNumber])
	}
	_dca := _gda._da[_ccg.PageNumber]
	_f.Log.Debug("P\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020c\u006f\u006d\u0070\u0073: \u0025\u0076", _ccg.PageNumber, _dca)
	_ccg.addTextRegionSegment(_ddf, _gda._dgf, _edb, _gda._da[_ccg.PageNumber], _gda.Classer.PtaLL, _gda.Classer.UndilatedTemplates, _gda.Classer.ClassIDs, nil, _ad(_gb), len(_gda._da[_ccg.PageNumber]))
	return nil
}

const (
	GenericEM EncodingMethod = iota
	CorrelationEM
	RankHausEM
)

func (_bdd *Document) GetPage(pageNumber int) (_ecd.Pager, error) {
	const _cd = "\u0044\u006fc\u0075\u006d\u0065n\u0074\u002e\u0047\u0065\u0074\u0050\u0061\u0067\u0065"
	if pageNumber < 0 {
		_f.Log.Debug("\u004a\u0042\u0049\u00472\u0020\u0050\u0061\u0067\u0065\u0020\u002d\u0020\u0047e\u0074\u0050\u0061\u0067\u0065\u003a\u0020\u0025\u0064\u002e\u0020\u0050\u0061\u0067\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062e\u0020\u006c\u006f\u0077\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030\u002e\u0020\u0025\u0073", pageNumber, _ccb.Stack())
		return nil, _ab.Errorf(_cd, "\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032\u0020d\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u002d\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064 \u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u003a\u0020\u0025\u0064", pageNumber)
	}
	if pageNumber > len(_bdd.Pages) {
		_f.Log.Debug("\u0050\u0061\u0067\u0065 n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0064\u002e\u0020%\u0073", pageNumber, _ccb.Stack())
		return nil, _ab.Error(_cd, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u002d\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_dgd, _deb := _bdd.Pages[pageNumber]
	if !_deb {
		_f.Log.Debug("\u0050\u0061\u0067\u0065 n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0064\u002e\u0020%\u0073", pageNumber, _ccb.Stack())
		return nil, _ab.Errorf(_cd, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u002d\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	return _dgd, nil
}
func (_fdgd *Page) getResolutionX() (int, error) {
	const _afda = "\u0067\u0065\u0074\u0052\u0065\u0073\u006f\u006c\u0075t\u0069\u006f\u006e\u0058"
	if _fdgd.ResolutionX != 0 {
		return _fdgd.ResolutionX, nil
	}
	_fada := _fdgd.getPageInformationSegment()
	if _fada == nil {
		return 0, _ab.Error(_afda, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_cfbb, _eefg := _fada.GetSegmentData()
	if _eefg != nil {
		return 0, _ab.Wrap(_eefg, _afda, "")
	}
	_aea, _dcfa := _cfbb.(*_ecd.PageInformationSegment)
	if !_dcfa {
		return 0, _ab.Errorf(_afda, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _cfbb)
	}
	_fdgd.ResolutionX = _aea.ResolutionX
	return _fdgd.ResolutionX, nil
}
func _ad(_bgc int) int {
	_faf := 0
	_abf := (_bgc & (_bgc - 1)) == 0
	_bgc >>= 1
	for ; _bgc != 0; _bgc >>= 1 {
		_faf++
	}
	if _abf {
		return _faf
	}
	return _faf + 1
}
func (_ebaa *Document) completeSymbols() (_edd error) {
	const _ca = "\u0063o\u006dp\u006c\u0065\u0074\u0065\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	if _ebaa.Classer == nil {
		return nil
	}
	if _ebaa.Classer.UndilatedTemplates == nil {
		return _ab.Error(_ca, "\u006e\u006f t\u0065\u006d\u0070l\u0061\u0074\u0065\u0073 de\u0066in\u0065\u0064\u0020\u0066\u006f\u0072\u0020th\u0065\u0020\u0063\u006c\u0061\u0073\u0073e\u0072")
	}
	_caa := len(_ebaa.Pages) == 1
	_fab := make([]int, _ebaa.Classer.UndilatedTemplates.Size())
	var _bef int
	for _ccde := 0; _ccde < _ebaa.Classer.ClassIDs.Size(); _ccde++ {
		_bef, _edd = _ebaa.Classer.ClassIDs.Get(_ccde)
		if _edd != nil {
			return _ab.Wrap(_edd, _ca, "\u0063\u006c\u0061\u0073\u0073\u0020\u0049\u0044\u0027\u0073")
		}
		_fab[_bef]++
	}
	var _ga []int
	for _ba := 0; _ba < _ebaa.Classer.UndilatedTemplates.Size(); _ba++ {
		if _fab[_ba] == 0 {
			return _ab.Error(_ca, "\u006eo\u0020\u0073y\u006d\u0062\u006f\u006cs\u0020\u0069\u006es\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006fun\u0064\u0020\u0066o\u0072\u0020g\u0069\u0076\u0065\u006e\u0020\u0063l\u0061\u0073s\u003f\u0020")
		}
		if _fab[_ba] > 1 || _caa {
			_ga = append(_ga, _ba)
		}
	}
	_ebaa._db = len(_ga)
	var _gba, _cf int
	for _ade := 0; _ade < _ebaa.Classer.ComponentPageNumbers.Size(); _ade++ {
		_gba, _edd = _ebaa.Classer.ComponentPageNumbers.Get(_ade)
		if _edd != nil {
			return _ab.Wrapf(_edd, _ca, "p\u0061\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027 \u006e\u006f\u0074\u0020\u0066\u006f\u0075nd\u0020\u0069\u006e\u0020t\u0068\u0065\u0020\u0063\u006c\u0061\u0073\u0073\u0065r \u0070\u0061g\u0065\u006e\u0075\u006d\u0062\u0065\u0072\u0073", _ade)
		}
		_cf, _edd = _ebaa.Classer.ClassIDs.Get(_ade)
		if _edd != nil {
			return _ab.Wrapf(_edd, _ca, "\u0063\u0061\u006e\u0027\u0074\u0020\u0067e\u0074\u0020\u0073y\u006d\u0062\u006f\u006c \u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0027\u0025\u0064\u0027\u0020\u0066\u0072\u006f\u006d\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072", _gba)
		}
		if _fab[_cf] == 1 && !_caa {
			_ebaa._dc[_gba] = append(_ebaa._dc[_gba], _cf)
		}
	}
	if _edd = _ebaa.Classer.ComputeLLCorners(); _edd != nil {
		return _ab.Wrap(_edd, _ca, "")
	}
	return nil
}
func (_febg *Page) addTextRegionSegment(_cfb []*_ecd.Header, _dea, _ecde map[int]int, _dgdc []int, _cgd *_ee.Points, _abgb *_ee.Bitmaps, _dcc *_a.IntSlice, _gbf *_ee.Boxes, _abe, _cdce int) {
	_ggcc := &_ecd.TextRegion{NumberOfSymbols: uint32(_cdce)}
	_ggcc.InitEncode(_dea, _ecde, _dgdc, _cgd, _abgb, _dcc, _gbf, _febg.FinalWidth, _febg.FinalHeight, _abe)
	_gbg := &_ecd.Header{RTSegments: _cfb, SegmentData: _ggcc, PageAssociation: _febg.PageNumber, Type: _ecd.TImmediateTextRegion}
	_gcf := _ecd.TPageInformation
	if _ecde != nil {
		_gcf = _ecd.TSymbolDictionary
	}
	var _egf int
	for ; _egf < len(_febg.Segments); _egf++ {
		if _febg.Segments[_egf].Type == _gcf {
			_egf++
			break
		}
	}
	_febg.Segments = append(_febg.Segments, nil)
	copy(_febg.Segments[_egf+1:], _febg.Segments[_egf:])
	_febg.Segments[_egf] = _gbg
}
func _bfe(_gfa _g.StreamReader, _abd *Globals) (*Document, error) {
	_dfc := &Document{Pages: make(map[int]*Page), InputStream: _gfa, OrganizationType: _ecd.OSequential, NumberOfPagesUnknown: true, GlobalSegments: _abd, _gd: 9}
	if _dfc.GlobalSegments == nil {
		_dfc.GlobalSegments = &Globals{}
	}
	if _ebfd := _dfc.mapData(); _ebfd != nil {
		return nil, _ebfd
	}
	return _dfc, nil
}
func (_cfd *Page) composePageBitmap() error {
	const _adg = "\u0063\u006f\u006d\u0070\u006f\u0073\u0065\u0050\u0061\u0067\u0065\u0042i\u0074\u006d\u0061\u0070"
	if _cfd.PageNumber == 0 {
		return nil
	}
	_dbg := _cfd.getPageInformationSegment()
	if _dbg == nil {
		return _ab.Error(_adg, "\u0070\u0061\u0067e \u0069\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069o\u006e \u0073e\u0067m\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_ggcb, _cba := _dbg.GetSegmentData()
	if _cba != nil {
		return _cba
	}
	_afcf, _ecg := _ggcb.(*_ecd.PageInformationSegment)
	if !_ecg {
		return _ab.Error(_adg, "\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006da\u0074\u0069\u006f\u006e \u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065")
	}
	if _cba = _cfd.createPage(_afcf); _cba != nil {
		return _ab.Wrap(_cba, _adg, "")
	}
	_cfd.clearSegmentData()
	return nil
}
func (_babec *Page) GetBitmap() (_gge *_ee.Bitmap, _gfgf error) {
	_f.Log.Trace(_cc.Sprintf("\u005b\u0050\u0041G\u0045\u005d\u005b\u0023%\u0064\u005d\u0020\u0047\u0065\u0074\u0042i\u0074\u006d\u0061\u0070\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e", _babec.PageNumber))
	defer func() {
		if _gfgf != nil {
			_f.Log.Trace(_cc.Sprintf("\u005b\u0050\u0041\u0047\u0045\u005d\u005b\u0023\u0025\u0064\u005d\u0020\u0047\u0065\u0074B\u0069t\u006d\u0061\u0070\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _babec.PageNumber, _gfgf))
		} else {
			_f.Log.Trace(_cc.Sprintf("\u005b\u0050\u0041\u0047\u0045\u005d\u005b\u0023\u0025\u0064]\u0020\u0047\u0065\u0074\u0042\u0069\u0074m\u0061\u0070\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064", _babec.PageNumber))
		}
	}()
	if _babec.Bitmap != nil {
		return _babec.Bitmap, nil
	}
	_gfgf = _babec.composePageBitmap()
	if _gfgf != nil {
		return nil, _gfgf
	}
	return _babec.Bitmap, nil
}
func (_efe *Document) nextSegmentNumber() uint32 {
	_agd := _efe.CurrentSegmentNumber
	_efe.CurrentSegmentNumber++
	return _agd
}
func (_becb *Document) mapData() error {
	const _cda = "\u006da\u0070\u0044\u0061\u0074\u0061"
	var (
		_caf []*_ecd.Header
		_gag int64
		_aff _ecd.Type
	)
	_cfe, _bgee := _becb.isFileHeaderPresent()
	if _bgee != nil {
		return _ab.Wrap(_bgee, _cda, "")
	}
	if _cfe {
		if _bgee = _becb.parseFileHeader(); _bgee != nil {
			return _ab.Wrap(_bgee, _cda, "")
		}
		_gag += int64(_becb._gd)
		_becb.FullHeaders = true
	}
	var (
		_dcg *Page
		_fgb bool
	)
	for _aff != 51 && !_fgb {
		_bfbf, _fef := _ecd.NewHeader(_becb, _becb.InputStream, _gag, _becb.OrganizationType)
		if _fef != nil {
			return _ab.Wrap(_fef, _cda, "")
		}
		_f.Log.Trace("\u0044\u0065c\u006f\u0064\u0069\u006eg\u0020\u0073e\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0025\u0064\u002c\u0020\u0054\u0079\u0070e\u003a\u0020\u0025\u0073", _bfbf.SegmentNumber, _bfbf.Type)
		_aff = _bfbf.Type
		if _aff != _ecd.TEndOfFile {
			if _bfbf.PageAssociation != 0 {
				_dcg = _becb.Pages[_bfbf.PageAssociation]
				if _dcg == nil {
					_dcg = _dgc(_becb, _bfbf.PageAssociation)
					_becb.Pages[_bfbf.PageAssociation] = _dcg
					if _becb.NumberOfPagesUnknown {
						_becb.NumberOfPages++
					}
				}
				_dcg.Segments = append(_dcg.Segments, _bfbf)
			} else {
				_becb.GlobalSegments.AddSegment(_bfbf)
			}
		}
		_caf = append(_caf, _bfbf)
		_gag = _becb.InputStream.StreamPosition()
		if _becb.OrganizationType == _ecd.OSequential {
			_gag += int64(_bfbf.SegmentDataLength)
		}
		_fgb, _fef = _becb.reachedEOF(_gag)
		if _fef != nil {
			_f.Log.Debug("\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0072\u0065\u0061\u0063h\u0065\u0064\u0020\u0045\u004f\u0046\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _fef)
			return _ab.Wrap(_fef, _cda, "")
		}
	}
	_becb.determineRandomDataOffsets(_caf, uint64(_gag))
	return nil
}
func (_fad *Document) encodeSegment(_dbc *_ecd.Header, _eca *int) error {
	const _dbb = "\u0065\u006e\u0063\u006f\u0064\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074"
	_dbc.SegmentNumber = _fad.nextSegmentNumber()
	_dec, _gf := _dbc.Encode(_fad._eec)
	if _gf != nil {
		return _ab.Wrapf(_gf, _dbb, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", _dbc.SegmentNumber)
	}
	*_eca += _dec
	return nil
}
func (_fgea *Page) getWidth() (int, error) {
	const _faa = "\u0067\u0065\u0074\u0057\u0069\u0064\u0074\u0068"
	if _fgea.FinalWidth != 0 {
		return _fgea.FinalWidth, nil
	}
	_bgd := _fgea.getPageInformationSegment()
	if _bgd == nil {
		return 0, _ab.Error(_faa, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_dcee, _bce := _bgd.GetSegmentData()
	if _bce != nil {
		return 0, _ab.Wrap(_bce, _faa, "")
	}
	_fca, _afd := _dcee.(*_ecd.PageInformationSegment)
	if !_afd {
		return 0, _ab.Errorf(_faa, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _dcee)
	}
	_fgea.FinalWidth = _fca.PageBMWidth
	return _fgea.FinalWidth, nil
}
func (_gff *Page) createStripedPage(_ffc *_ecd.PageInformationSegment) error {
	const _eedf = "\u0063\u0072\u0065\u0061\u0074\u0065\u0053\u0074\u0072\u0069\u0070\u0065d\u0050\u0061\u0067\u0065"
	_ge, _gfc := _gff.collectPageStripes()
	if _gfc != nil {
		return _ab.Wrap(_gfc, _eedf, "")
	}
	var _aeda int
	for _, _adgf := range _ge {
		if _bdf, _bbg := _adgf.(*_ecd.EndOfStripe); _bbg {
			_aeda = _bdf.LineNumber() + 1
		} else {
			_ece := _adgf.(_ecd.Regioner)
			_feff := _ece.GetRegionInfo()
			_cga := _gff.getCombinationOperator(_ffc, _feff.CombinaionOperator)
			_cfbd, _aaf := _ece.GetRegionBitmap()
			if _aaf != nil {
				return _ab.Wrap(_aaf, _eedf, "")
			}
			_aaf = _ee.Blit(_cfbd, _gff.Bitmap, int(_feff.XLocation), _aeda, _cga)
			if _aaf != nil {
				return _ab.Wrap(_aaf, _eedf, "")
			}
		}
	}
	return nil
}
func (_dece *Document) GetGlobalSegment(i int) (*_ecd.Header, error) {
	_bge, _fbf := _dece.GlobalSegments.GetSegment(i)
	if _fbf != nil {
		return nil, _ab.Wrap(_fbf, "\u0047\u0065t\u0047\u006c\u006fb\u0061\u006c\u0053\u0065\u0067\u006d\u0065\u006e\u0074", "")
	}
	return _bge, nil
}
func InitEncodeDocument(fullHeaders bool) *Document {
	return &Document{FullHeaders: fullHeaders, _eec: _g.BufferedMSB(), Pages: map[int]*Page{}, _dc: map[int][]int{}, _dgf: map[int]int{}, _da: map[int][]int{}}
}
func (_beca *Page) createPage(_dcff *_ecd.PageInformationSegment) error {
	var _fbd error
	if !_dcff.IsStripe || _dcff.PageBMHeight != -1 {
		_fbd = _beca.createNormalPage(_dcff)
	} else {
		_fbd = _beca.createStripedPage(_dcff)
	}
	return _fbd
}
func (_gdd *Document) nextPageNumber() uint32 { _gdd.NumberOfPages++; return _gdd.NumberOfPages }
func (_ggcf *Document) GetNumberOfPages() (uint32, error) {
	if _ggcf.NumberOfPagesUnknown || _ggcf.NumberOfPages == 0 {
		if len(_ggcf.Pages) == 0 {
			if _bab := _ggcf.mapData(); _bab != nil {
				return 0, _ab.Wrap(_bab, "\u0044o\u0063\u0075\u006d\u0065n\u0074\u002e\u0047\u0065\u0074N\u0075m\u0062e\u0072\u004f\u0066\u0050\u0061\u0067\u0065s", "")
			}
		}
		return uint32(len(_ggcf.Pages)), nil
	}
	return _ggcf.NumberOfPages, nil
}
func DecodeDocument(input _g.StreamReader, globals *Globals) (*Document, error) {
	return _bfe(input, globals)
}

var _dg = []byte{0x97, 0x4A, 0x42, 0x32, 0x0D, 0x0A, 0x1A, 0x0A}

func (_ded *Page) nextSegmentNumber() uint32 { return _ded.Document.nextSegmentNumber() }
func (_fdd *Page) AddPageInformationSegment() {
	_dag := &_ecd.PageInformationSegment{PageBMWidth: _fdd.FinalWidth, PageBMHeight: _fdd.FinalHeight, ResolutionX: _fdd.ResolutionX, ResolutionY: _fdd.ResolutionY, IsLossless: _fdd.IsLossless}
	if _fdd.BlackIsOne {
		_dag.DefaultPixelValue = uint8(0x1)
	}
	_dce := &_ecd.Header{PageAssociation: _fdd.PageNumber, SegmentDataLength: uint64(_dag.Size()), SegmentData: _dag, Type: _ecd.TPageInformation}
	_fdd.Segments = append(_fdd.Segments, _dce)
}
func (_bg *Document) AddClassifiedPage(bm *_ee.Bitmap, method _d.Method) (_af error) {
	const _eb = "\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0041\u0064d\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064P\u0061\u0067\u0065"
	if !_bg.FullHeaders && _bg.NumberOfPages != 0 {
		return _ab.Error(_eb, "\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0061\u006c\u0072\u0065a\u0064\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0070\u0061\u0067\u0065\u002e\u0020\u0046\u0069\u006c\u0065\u004d\u006f\u0064\u0065\u0020\u0064\u0069\u0073\u0061\u006c\u006c\u006f\u0077\u0073\u0020\u0061\u0064\u0064i\u006e\u0067\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e \u006f\u006e\u0065\u0020\u0070\u0061g\u0065")
	}
	if _bg.Classer == nil {
		if _bg.Classer, _af = _d.Init(_d.DefaultSettings()); _af != nil {
			return _ab.Wrap(_af, _eb, "")
		}
	}
	_bf := int(_bg.nextPageNumber())
	_ce := &Page{Segments: []*_ecd.Header{}, Bitmap: bm, Document: _bg, FinalHeight: bm.Height, FinalWidth: bm.Width, PageNumber: _bf}
	_bg.Pages[_bf] = _ce
	switch method {
	case _d.RankHaus:
		_ce.EncodingMethod = RankHausEM
	case _d.Correlation:
		_ce.EncodingMethod = CorrelationEM
	}
	_ce.AddPageInformationSegment()
	if _af = _bg.Classer.AddPage(bm, _bf, method); _af != nil {
		return _ab.Wrap(_af, _eb, "")
	}
	if _bg.FullHeaders {
		_ce.AddEndOfPageSegment()
	}
	return nil
}
func (_decb *Page) GetSegment(number int) (*_ecd.Header, error) {
	const _bfgc = "\u0050a\u0067e\u002e\u0047\u0065\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074"
	for _, _faeac := range _decb.Segments {
		if _faeac.SegmentNumber == uint32(number) {
			return _faeac, nil
		}
	}
	_adee := make([]uint32, len(_decb.Segments))
	for _fbb, _beb := range _decb.Segments {
		_adee[_fbb] = _beb.SegmentNumber
	}
	return nil, _ab.Errorf(_bfgc, "\u0073e\u0067\u006d\u0065n\u0074\u0020\u0077i\u0074h \u006e\u0075\u006d\u0062\u0065\u0072\u003a \u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0067\u0065\u003a\u0020'%\u0064'\u002e\u0020\u004b\u006e\u006f\u0077n\u0020\u0073\u0065\u0067\u006de\u006e\u0074\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0073\u003a \u0025\u0076", number, _decb.PageNumber, _adee)
}
func (_bda *Page) collectPageStripes() (_gfd []_ecd.Segmenter, _fbfg error) {
	const _fdb = "\u0063o\u006cl\u0065\u0063\u0074\u0050\u0061g\u0065\u0053t\u0072\u0069\u0070\u0065\u0073"
	var _fce _ecd.Segmenter
	for _, _ceg := range _bda.Segments {
		switch _ceg.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_fce, _fbfg = _ceg.GetSegmentData()
			if _fbfg != nil {
				return nil, _ab.Wrap(_fbfg, _fdb, "")
			}
			_gfd = append(_gfd, _fce)
		case 50:
			_fce, _fbfg = _ceg.GetSegmentData()
			if _fbfg != nil {
				return nil, _fbfg
			}
			_edbg, _bfeg := _fce.(*_ecd.EndOfStripe)
			if !_bfeg {
				return nil, _ab.Errorf(_fdb, "\u0045\u006e\u0064\u004f\u0066\u0053\u0074\u0072\u0069\u0070\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u006f\u0066\u0020\u0076\u0061l\u0069\u0064\u0020\u0074\u0079p\u0065\u003a \u0027\u0025\u0054\u0027", _fce)
			}
			_gfd = append(_gfd, _edbg)
			_bda.FinalHeight = _edbg.LineNumber()
		}
	}
	return _gfd, nil
}
func (_ddd *Document) encodeFileHeader(_ccc _g.BinaryWriter) (_cea int, _abg error) {
	const _ac = "\u0065\u006ec\u006f\u0064\u0065F\u0069\u006c\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_cea, _abg = _ccc.Write(_dg)
	if _abg != nil {
		return _cea, _ab.Wrap(_abg, _ac, "\u0069\u0064")
	}
	if _abg = _ccc.WriteByte(0x01); _abg != nil {
		return _cea, _ab.Wrap(_abg, _ac, "\u0066\u006c\u0061g\u0073")
	}
	_cea++
	_aed := make([]byte, 4)
	_c.BigEndian.PutUint32(_aed, _ddd.NumberOfPages)
	_aad, _abg := _ccc.Write(_aed)
	if _abg != nil {
		return _aad, _ab.Wrap(_abg, _ac, "p\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_cea += _aad
	return _cea, nil
}
func (_afa *Page) GetWidth() (int, error)       { return _afa.getWidth() }
func (_dcf *Page) GetResolutionX() (int, error) { return _dcf.getResolutionX() }
func (_cace *Page) Encode(w _g.BinaryWriter) (_ffa int, _cdca error) {
	const _ffg = "P\u0061\u0067\u0065\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	var _fgg int
	for _, _gbfg := range _cace.Segments {
		if _fgg, _cdca = _gbfg.Encode(w); _cdca != nil {
			return _ffa, _ab.Wrap(_cdca, _ffg, "")
		}
		_ffa += _fgg
	}
	return _ffa, nil
}
func (_cef *Globals) AddSegment(segment *_ecd.Header) { _cef.Segments = append(_cef.Segments, segment) }
func (_eda *Document) isFileHeaderPresent() (bool, error) {
	_eda.InputStream.Mark()
	for _, _daef := range _dg {
		_gbdg, _fag := _eda.InputStream.ReadByte()
		if _fag != nil {
			return false, _fag
		}
		if _daef != _gbdg {
			_eda.InputStream.Reset()
			return false, nil
		}
	}
	_eda.InputStream.Reset()
	return true, nil
}
func (_aae *Page) clearSegmentData() {
	for _ddg := range _aae.Segments {
		_aae.Segments[_ddg].CleanSegmentData()
	}
}
func (_agdc *Page) GetHeight() (int, error) { return _agdc.getHeight() }

type Page struct {
	Segments           []*_ecd.Header
	PageNumber         int
	Bitmap             *_ee.Bitmap
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

func (_dba *Document) produceClassifiedPages() (_gg error) {
	const _ccd = "\u0070\u0072\u006f\u0064uc\u0065\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0050\u0061\u0067e\u0073"
	if _dba.Classer == nil {
		return nil
	}
	var (
		_dab *Page
		_ed  bool
		_ggc *_ecd.Header
	)
	for _feb := 1; _feb <= int(_dba.NumberOfPages); _feb++ {
		if _dab, _ed = _dba.Pages[_feb]; !_ed {
			return _ab.Errorf(_ccd, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _feb)
		}
		if _dab.EncodingMethod == GenericEM {
			continue
		}
		if _ggc == nil {
			if _ggc, _gg = _dba.GlobalSegments.GetSymbolDictionary(); _gg != nil {
				return _ab.Wrap(_gg, _ccd, "")
			}
		}
		if _gg = _dba.produceClassifiedPage(_dab, _ggc); _gg != nil {
			return _ab.Wrapf(_gg, _ccd, "\u0070\u0061\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _feb)
		}
	}
	return nil
}
func (_eag *Document) completeClassifiedPages() (_bfb error) {
	const _ff = "\u0063\u006f\u006dpl\u0065\u0074\u0065\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0050\u0061\u0067\u0065\u0073"
	if _eag.Classer == nil {
		return nil
	}
	_eag._fa = make([]int, _eag.Classer.UndilatedTemplates.Size())
	for _dae := 0; _dae < _eag.Classer.ClassIDs.Size(); _dae++ {
		_eg, _bfg := _eag.Classer.ClassIDs.Get(_dae)
		if _bfg != nil {
			return _ab.Wrapf(_bfg, _ff, "\u0063\u006c\u0061\u0073s \u0077\u0069\u0074\u0068\u0020\u0069\u0064\u003a\u0020\u0027\u0025\u0064\u0027", _dae)
		}
		_eag._fa[_eg]++
	}
	var _eba []int
	for _dd := 0; _dd < _eag.Classer.UndilatedTemplates.Size(); _dd++ {
		if _eag.NumberOfPages == 1 || _eag._fa[_dd] > 1 {
			_eba = append(_eba, _dd)
		}
	}
	var (
		_bb *Page
		_cb bool
	)
	for _cg, _fe := range *_eag.Classer.ComponentPageNumbers {
		if _bb, _cb = _eag.Pages[_fe]; !_cb {
			return _ab.Errorf(_ff, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _cg)
		}
		if _bb.EncodingMethod == GenericEM {
			_f.Log.Error("\u0047\u0065\u006e\u0065\u0072\u0069c\u0020\u0070\u0061g\u0065\u0020\u0077i\u0074\u0068\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u003a \u0027\u0025\u0064\u0027\u0020ma\u0070\u0070\u0065\u0064\u0020\u0061\u0073\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0020\u0070\u0061\u0067\u0065", _cg)
			continue
		}
		_eag._da[_fe] = append(_eag._da[_fe], _cg)
		_ege, _dge := _eag.Classer.ClassIDs.Get(_cg)
		if _dge != nil {
			return _ab.Wrapf(_dge, _ff, "\u006e\u006f\u0020\u0073uc\u0068\u0020\u0063\u006c\u0061\u0073\u0073\u0049\u0044\u003a\u0020\u0025\u0064", _cg)
		}
		if _eag._fa[_ege] == 1 && _eag.NumberOfPages != 1 {
			_gc := append(_eag._dc[_fe], _ege)
			_eag._dc[_fe] = _gc
		}
	}
	if _bfb = _eag.Classer.ComputeLLCorners(); _bfb != nil {
		return _ab.Wrap(_bfb, _ff, "")
	}
	if _, _bfb = _eag.addSymbolDictionary(0, _eag.Classer.UndilatedTemplates, _eba, _eag._dgf, false); _bfb != nil {
		return _ab.Wrap(_bfb, _ff, "")
	}
	return nil
}
func (_eef *Page) GetResolutionY() (int, error) { return _eef.getResolutionY() }
func (_egd *Page) getResolutionY() (int, error) {
	const _cabf = "\u0067\u0065\u0074\u0052\u0065\u0073\u006f\u006c\u0075t\u0069\u006f\u006e\u0059"
	if _egd.ResolutionY != 0 {
		return _egd.ResolutionY, nil
	}
	_gbe := _egd.getPageInformationSegment()
	if _gbe == nil {
		return 0, _ab.Error(_cabf, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_eae, _cgdg := _gbe.GetSegmentData()
	if _cgdg != nil {
		return 0, _ab.Wrap(_cgdg, _cabf, "")
	}
	_bbc, _fbe := _eae.(*_ecd.PageInformationSegment)
	if !_fbe {
		return 0, _ab.Errorf(_cabf, "\u0070\u0061\u0067\u0065\u0020\u0069\u006e\u0066o\u0072\u006d\u0061ti\u006f\u006e\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0074\u0079\u0070\u0065\u003a\u0027%\u0054\u0027", _eae)
	}
	_egd.ResolutionY = _bbc.ResolutionY
	return _egd.ResolutionY, nil
}
func (_cbb *Page) getPageInformationSegment() *_ecd.Header {
	for _, _deac := range _cbb.Segments {
		if _deac.Type == _ecd.TPageInformation {
			return _deac
		}
	}
	_f.Log.Debug("\u0050\u0061\u0067\u0065\u0020\u0069\u006e\u0066o\u0072\u006d\u0061ti\u006f\u006e\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0066o\u0072\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0073\u002e", _cbb)
	return nil
}
func (_dbcc *Document) encodeEOFHeader(_bcd _g.BinaryWriter) (_daa int, _gad error) {
	_fgd := &_ecd.Header{SegmentNumber: _dbcc.nextSegmentNumber(), Type: _ecd.TEndOfFile}
	if _daa, _gad = _fgd.Encode(_bcd); _gad != nil {
		return 0, _ab.Wrap(_gad, "\u0065n\u0063o\u0064\u0065\u0045\u004f\u0046\u0048\u0065\u0061\u0064\u0065\u0072", "")
	}
	return _daa, nil
}
func (_fbc *Globals) GetSegmentByIndex(index int) (*_ecd.Header, error) {
	const _eedg = "\u0047l\u006f\u0062\u0061\u006cs\u002e\u0047\u0065\u0074\u0053e\u0067m\u0065n\u0074\u0042\u0079\u0049\u006e\u0064\u0065x"
	if _fbc == nil {
		return nil, _ab.Error(_eedg, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_fbc.Segments) == 0 {
		return nil, _ab.Error(_eedg, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	if index > len(_fbc.Segments)-1 {
		return nil, _ab.Error(_eedg, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	return _fbc.Segments[index], nil
}

type Document struct {
	Pages                map[int]*Page
	NumberOfPagesUnknown bool
	NumberOfPages        uint32
	GBUseExtTemplate     bool
	InputStream          _g.StreamReader
	GlobalSegments       *Globals
	OrganizationType     _ecd.OrganizationType
	Classer              *_d.Classer
	XRes, YRes           int
	FullHeaders          bool
	CurrentSegmentNumber uint32
	AverageTemplates     *_ee.Bitmaps
	BaseIndexes          []int
	Refinement           bool
	RefineLevel          int
	_gd                  uint8
	_eec                 *_g.BufferedWriter
	EncodeGlobals        bool
	_db                  int
	_dc                  map[int][]int
	_da                  map[int][]int
	_fa                  []int
	_dgf                 map[int]int
}

func (_cgf *Document) Encode() (_bbf []byte, _fdg error) {
	const _cab = "\u0044o\u0063u\u006d\u0065\u006e\u0074\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	var _bec, _gbd int
	if _cgf.FullHeaders {
		if _bec, _fdg = _cgf.encodeFileHeader(_cgf._eec); _fdg != nil {
			return nil, _ab.Wrap(_fdg, _cab, "")
		}
	}
	var (
		_efg bool
		_dda *_ecd.Header
		_ag  *Page
	)
	if _fdg = _cgf.completeClassifiedPages(); _fdg != nil {
		return nil, _ab.Wrap(_fdg, _cab, "")
	}
	if _fdg = _cgf.produceClassifiedPages(); _fdg != nil {
		return nil, _ab.Wrap(_fdg, _cab, "")
	}
	if _cgf.GlobalSegments != nil {
		for _, _dda = range _cgf.GlobalSegments.Segments {
			if _fdg = _cgf.encodeSegment(_dda, &_bec); _fdg != nil {
				return nil, _ab.Wrap(_fdg, _cab, "")
			}
		}
	}
	for _bc := 1; _bc <= int(_cgf.NumberOfPages); _bc++ {
		if _ag, _efg = _cgf.Pages[_bc]; !_efg {
			return nil, _ab.Errorf(_cab, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _bc)
		}
		for _, _dda = range _ag.Segments {
			if _fdg = _cgf.encodeSegment(_dda, &_bec); _fdg != nil {
				return nil, _ab.Wrap(_fdg, _cab, "")
			}
		}
	}
	if _cgf.FullHeaders {
		if _gbd, _fdg = _cgf.encodeEOFHeader(_cgf._eec); _fdg != nil {
			return nil, _ab.Wrap(_fdg, _cab, "")
		}
		_bec += _gbd
	}
	_bbf = _cgf._eec.Data()
	if len(_bbf) != _bec {
		_f.Log.Debug("\u0042\u0079\u0074\u0065\u0073 \u0077\u0072\u0069\u0074\u0074\u0065\u006e \u0028\u006e\u0029\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0064\u0061\u0074\u0061\u0020\u0065\u006e\u0063\u006fd\u0065\u0064\u003a\u0020\u0027\u0025d\u0027", _bec, len(_bbf))
	}
	return _bbf, nil
}
func (_gaf *Page) countRegions() int {
	var _bfga int
	for _, _aade := range _gaf.Segments {
		switch _aade.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_bfga++
		}
	}
	return _bfga
}

type EncodingMethod int

func (_gfg *Page) AddEndOfPageSegment() {
	_ace := &_ecd.Header{Type: _ecd.TEndOfPage, PageAssociation: _gfg.PageNumber}
	_gfg.Segments = append(_gfg.Segments, _ace)
}
func (_edg *Document) reachedEOF(_faea int64) (bool, error) {
	const _acc = "\u0072\u0065\u0061\u0063\u0068\u0065\u0064\u0045\u004f\u0046"
	_, _gcc := _edg.InputStream.Seek(_faea, _ea.SeekStart)
	if _gcc != nil {
		_f.Log.Debug("\u0072\u0065\u0061c\u0068\u0065\u0064\u0045\u004f\u0046\u0020\u002d\u0020\u0064\u002e\u0049\u006e\u0070\u0075\u0074\u0053\u0074\u0072\u0065\u0061\u006d\u002e\u0053\u0065\u0065\u006b\u0020\u0066a\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _gcc)
		return false, _ab.Wrap(_gcc, _acc, "\u0069n\u0070\u0075\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020s\u0065\u0065\u006b\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_, _gcc = _edg.InputStream.ReadBits(32)
	if _gcc == _ea.EOF {
		return true, nil
	} else if _gcc != nil {
		return false, _ab.Wrap(_gcc, _acc, "")
	}
	return false, nil
}
func (_eab *Page) createNormalPage(_gab *_ecd.PageInformationSegment) error {
	const _efd = "\u0063\u0072e\u0061\u0074\u0065N\u006f\u0072\u006d\u0061\u006c\u0050\u0061\u0067\u0065"
	_eab.Bitmap = _ee.New(_gab.PageBMWidth, _gab.PageBMHeight)
	if _gab.DefaultPixelValue != 0 {
		_eab.Bitmap.SetDefaultPixel()
	}
	for _, _bad := range _eab.Segments {
		switch _bad.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_f.Log.Trace("\u0047\u0065\u0074\u0074in\u0067\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0064", _bad.SegmentNumber)
			_gadg, _fac := _bad.GetSegmentData()
			if _fac != nil {
				return _fac
			}
			_fabg, _gbcg := _gadg.(_ecd.Regioner)
			if !_gbcg {
				_f.Log.Debug("\u0053\u0065g\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0054\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0052\u0065\u0067\u0069on\u0065\u0072", _gadg)
				return _ab.Errorf(_efd, "i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062i\u0067\u0032\u0020\u0073\u0065\u0067\u006den\u0074\u0020\u0074\u0079p\u0065\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061 R\u0065\u0067i\u006f\u006e\u0065\u0072\u003a\u0020\u0025\u0054", _gadg)
			}
			_cfdf, _fac := _fabg.GetRegionBitmap()
			if _fac != nil {
				return _ab.Wrap(_fac, _efd, "")
			}
			if _eab.fitsPage(_gab, _cfdf) {
				_eab.Bitmap = _cfdf
			} else {
				_cefd := _fabg.GetRegionInfo()
				_bebg := _eab.getCombinationOperator(_gab, _cefd.CombinaionOperator)
				_fac = _ee.Blit(_cfdf, _eab.Bitmap, int(_cefd.XLocation), int(_cefd.YLocation), _bebg)
				if _fac != nil {
					return _ab.Wrap(_fac, _efd, "")
				}
			}
		}
	}
	return nil
}
func (_df *Document) AddGenericPage(bm *_ee.Bitmap, duplicateLineRemoval bool) (_gdg error) {
	const _b = "\u0044\u006f\u0063um\u0065\u006e\u0074\u002e\u0041\u0064\u0064\u0047\u0065\u006e\u0065\u0072\u0069\u0063\u0050\u0061\u0067\u0065"
	if !_df.FullHeaders && _df.NumberOfPages != 0 {
		return _ab.Error(_b, "\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0061\u006c\u0072\u0065a\u0064\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0070\u0061\u0067\u0065\u002e\u0020\u0046\u0069\u006c\u0065\u004d\u006f\u0064\u0065\u0020\u0064\u0069\u0073\u0061\u006c\u006c\u006f\u0077\u0073\u0020\u0061\u0064\u0064i\u006e\u0067\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e \u006f\u006e\u0065\u0020\u0070\u0061g\u0065")
	}
	_ef := &Page{Segments: []*_ecd.Header{}, Bitmap: bm, Document: _df, FinalHeight: bm.Height, FinalWidth: bm.Width, IsLossless: true, BlackIsOne: bm.Color == _ee.Chocolate}
	_ef.PageNumber = int(_df.nextPageNumber())
	_df.Pages[_ef.PageNumber] = _ef
	bm.InverseData()
	_ef.AddPageInformationSegment()
	if _gdg = _ef.AddGenericRegion(bm, 0, 0, 0, _ecd.TImmediateGenericRegion, duplicateLineRemoval); _gdg != nil {
		return _ab.Wrap(_gdg, _b, "")
	}
	if _df.FullHeaders {
		_ef.AddEndOfPageSegment()
	}
	return nil
}
func (_dbf *Page) getCombinationOperator(_ffd *_ecd.PageInformationSegment, _edgf _ee.CombinationOperator) _ee.CombinationOperator {
	if _ffd.CombinationOperatorOverrideAllowed() {
		return _edgf
	}
	return _ffd.CombinationOperator()
}
func _dgc(_ggg *Document, _agb int) *Page {
	return &Page{Document: _ggg, PageNumber: _agb, Segments: []*_ecd.Header{}}
}
func (_eafg *Page) AddGenericRegion(bm *_ee.Bitmap, xloc, yloc, template int, tp _ecd.Type, duplicateLineRemoval bool) error {
	const _fbg = "P\u0061\u0067\u0065\u002eAd\u0064G\u0065\u006e\u0065\u0072\u0069c\u0052\u0065\u0067\u0069\u006f\u006e"
	_gga := &_ecd.GenericRegion{}
	if _accg := _gga.InitEncode(bm, xloc, yloc, template, duplicateLineRemoval); _accg != nil {
		return _ab.Wrap(_accg, _fbg, "")
	}
	_aag := &_ecd.Header{Type: _ecd.TImmediateGenericRegion, PageAssociation: _eafg.PageNumber, SegmentData: _gga}
	_eafg.Segments = append(_eafg.Segments, _aag)
	return nil
}
func (_gec *Page) getHeight() (int, error) {
	const _ccga = "\u0067e\u0074\u0048\u0065\u0069\u0067\u0068t"
	if _gec.FinalHeight != 0 {
		return _gec.FinalHeight, nil
	}
	_ebac := _gec.getPageInformationSegment()
	if _ebac == nil {
		return 0, _ab.Error(_ccga, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_adfg, _faed := _ebac.GetSegmentData()
	if _faed != nil {
		return 0, _ab.Wrap(_faed, _ccga, "")
	}
	_gbb, _daee := _adfg.(*_ecd.PageInformationSegment)
	if !_daee {
		return 0, _ab.Errorf(_ccga, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _adfg)
	}
	if _gbb.PageBMHeight == _ec.MaxInt32 {
		_, _faed = _gec.GetBitmap()
		if _faed != nil {
			return 0, _ab.Wrap(_faed, _ccga, "")
		}
	} else {
		_gec.FinalHeight = _gbb.PageBMHeight
	}
	return _gec.FinalHeight, nil
}
func (_bag *Globals) GetSegment(segmentNumber int) (*_ecd.Header, error) {
	const _daba = "\u0047l\u006fb\u0061\u006c\u0073\u002e\u0047e\u0074\u0053e\u0067\u006d\u0065\u006e\u0074"
	if _bag == nil {
		return nil, _ab.Error(_daba, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_bag.Segments) == 0 {
		return nil, _ab.Error(_daba, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	var _afc *_ecd.Header
	for _, _afc = range _bag.Segments {
		if _afc.SegmentNumber == uint32(segmentNumber) {
			break
		}
	}
	if _afc == nil {
		return nil, _ab.Error(_daba, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	return _afc, nil
}
func (_gea *Page) lastSegmentNumber() (_dgfcg uint32, _gbeg error) {
	const _dabe = "\u006c\u0061\u0073\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u004eu\u006d\u0062\u0065\u0072"
	if len(_gea.Segments) == 0 {
		return _dgfcg, _ab.Errorf(_dabe, "\u006e\u006f\u0020se\u0067\u006d\u0065\u006e\u0074\u0073\u0020\u0066\u006fu\u006ed\u0020i\u006e \u0074\u0068\u0065\u0020\u0070\u0061\u0067\u0065\u0020\u0027\u0025\u0064\u0027", _gea.PageNumber)
	}
	return _gea.Segments[len(_gea.Segments)-1].SegmentNumber, nil
}
func (_gbc *Document) parseFileHeader() error {
	const _cge = "\u0070a\u0072s\u0065\u0046\u0069\u006c\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_, _bae := _gbc.InputStream.Seek(8, _ea.SeekStart)
	if _bae != nil {
		return _ab.Wrap(_bae, _cge, "\u0069\u0064")
	}
	_, _bae = _gbc.InputStream.ReadBits(5)
	if _bae != nil {
		return _ab.Wrap(_bae, _cge, "\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0064\u0020\u0062\u0069\u0074\u0073")
	}
	_babe, _bae := _gbc.InputStream.ReadBit()
	if _bae != nil {
		return _ab.Wrap(_bae, _cge, "\u0065x\u0074e\u006e\u0064\u0065\u0064\u0020t\u0065\u006dp\u006c\u0061\u0074\u0065\u0073")
	}
	if _babe == 1 {
		_gbc.GBUseExtTemplate = true
	}
	_babe, _bae = _gbc.InputStream.ReadBit()
	if _bae != nil {
		return _ab.Wrap(_bae, _cge, "\u0075\u006e\u006b\u006eow\u006e\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	if _babe != 1 {
		_gbc.NumberOfPagesUnknown = false
	}
	_babe, _bae = _gbc.InputStream.ReadBit()
	if _bae != nil {
		return _ab.Wrap(_bae, _cge, "\u006f\u0072\u0067\u0061\u006e\u0069\u007a\u0061\u0074\u0069\u006f\u006e \u0074\u0079\u0070\u0065")
	}
	_gbc.OrganizationType = _ecd.OrganizationType(_babe)
	if !_gbc.NumberOfPagesUnknown {
		_gbc.NumberOfPages, _bae = _gbc.InputStream.ReadUint32()
		if _bae != nil {
			return _ab.Wrap(_bae, _cge, "\u006eu\u006db\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0061\u0067\u0065\u0073")
		}
		_gbc._gd = 13
	}
	return nil
}
func (_ggd *Page) fitsPage(_gafe *_ecd.PageInformationSegment, _gddg *_ee.Bitmap) bool {
	return _ggd.countRegions() == 1 && _gafe.DefaultPixelValue == 0 && _gafe.PageBMWidth == _gddg.Width && _gafe.PageBMHeight == _gddg.Height
}
func (_be *Document) addSymbolDictionary(_dcd int, _fb *_ee.Bitmaps, _fae []int, _fd map[int]int, _bd bool) (*_ecd.Header, error) {
	const _egg = "\u0061\u0064\u0064\u0053ym\u0062\u006f\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079"
	_eaf := &_ecd.SymbolDictionary{}
	if _dgfc := _eaf.InitEncode(_fb, _fae, _fd, _bd); _dgfc != nil {
		return nil, _dgfc
	}
	_adf := &_ecd.Header{Type: _ecd.TSymbolDictionary, PageAssociation: _dcd, SegmentData: _eaf}
	if _dcd == 0 {
		if _be.GlobalSegments == nil {
			_be.GlobalSegments = &Globals{}
		}
		_be.GlobalSegments.AddSegment(_adf)
		return _adf, nil
	}
	_aa, _dbd := _be.Pages[_dcd]
	if !_dbd {
		return nil, _ab.Errorf(_egg, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _dcd)
	}
	var (
		_dfg int
		_ebf *_ecd.Header
	)
	for _dfg, _ebf = range _aa.Segments {
		if _ebf.Type == _ecd.TPageInformation {
			break
		}
	}
	_dfg++
	_aa.Segments = append(_aa.Segments, nil)
	copy(_aa.Segments[_dfg+1:], _aa.Segments[_dfg:])
	_aa.Segments[_dfg] = _adf
	return _adf, nil
}
func (_fdf *Globals) GetSymbolDictionary() (*_ecd.Header, error) {
	const _fge = "G\u006c\u006f\u0062\u0061\u006c\u0073.\u0047\u0065\u0074\u0053\u0079\u006d\u0062\u006f\u006cD\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079"
	if _fdf == nil {
		return nil, _ab.Error(_fge, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_fdf.Segments) == 0 {
		return nil, _ab.Error(_fge, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	for _, _cee := range _fdf.Segments {
		if _cee.Type == _ecd.TSymbolDictionary {
			return _cee, nil
		}
	}
	return nil, _ab.Error(_fge, "\u0067\u006c\u006fba\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0020d\u0069c\u0074i\u006fn\u0061\u0072\u0079\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}
