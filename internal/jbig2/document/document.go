package document

import (
	_d "encoding/binary"
	_cb "fmt"
	_g "io"
	_ca "math"
	_f "runtime/debug"

	_a "bitbucket.org/shenghui0779/gopdf/common"
	_ac "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_gf "bitbucket.org/shenghui0779/gopdf/internal/jbig2/basic"
	_ag "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_ae "bitbucket.org/shenghui0779/gopdf/internal/jbig2/document/segments"
	_b "bitbucket.org/shenghui0779/gopdf/internal/jbig2/encoder/classer"
	_df "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

func (_cfg *Document) encodeEOFHeader(_cdc _ac.BinaryWriter) (_fdd int, _ge error) {
	_abdb := &_ae.Header{SegmentNumber: _cfg.nextSegmentNumber(), Type: _ae.TEndOfFile}
	if _fdd, _ge = _abdb.Encode(_cdc); _ge != nil {
		return 0, _df.Wrap(_ge, "\u0065n\u0063o\u0064\u0065\u0045\u004f\u0046\u0048\u0065\u0061\u0064\u0065\u0072", "")
	}
	return _fdd, nil
}
func (_ccf *Page) GetResolutionX() (int, error) { return _ccf.getResolutionX() }
func (_fgca *Page) String() string {
	return _cb.Sprintf("\u0050\u0061\u0067\u0065\u0020\u0023\u0025\u0064", _fgca.PageNumber)
}
func (_gd *Document) completeSymbols() (_eda error) {
	const _eg = "\u0063o\u006dp\u006c\u0065\u0074\u0065\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	if _gd.Classer == nil {
		return nil
	}
	if _gd.Classer.UndilatedTemplates == nil {
		return _df.Error(_eg, "\u006e\u006f t\u0065\u006d\u0070l\u0061\u0074\u0065\u0073 de\u0066in\u0065\u0064\u0020\u0066\u006f\u0072\u0020th\u0065\u0020\u0063\u006c\u0061\u0073\u0073e\u0072")
	}
	_aga := len(_gd.Pages) == 1
	_cf := make([]int, _gd.Classer.UndilatedTemplates.Size())
	var _bac int
	for _cdf := 0; _cdf < _gd.Classer.ClassIDs.Size(); _cdf++ {
		_bac, _eda = _gd.Classer.ClassIDs.Get(_cdf)
		if _eda != nil {
			return _df.Wrap(_eda, _eg, "\u0063\u006c\u0061\u0073\u0073\u0020\u0049\u0044\u0027\u0073")
		}
		_cf[_bac]++
	}
	var _bae []int
	for _db := 0; _db < _gd.Classer.UndilatedTemplates.Size(); _db++ {
		if _cf[_db] == 0 {
			return _df.Error(_eg, "\u006eo\u0020\u0073y\u006d\u0062\u006f\u006cs\u0020\u0069\u006es\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006fun\u0064\u0020\u0066o\u0072\u0020g\u0069\u0076\u0065\u006e\u0020\u0063l\u0061\u0073s\u003f\u0020")
		}
		if _cf[_db] > 1 || _aga {
			_bae = append(_bae, _db)
		}
	}
	_gd._cbc = len(_bae)
	var _acc, _agd int
	for _cc := 0; _cc < _gd.Classer.ComponentPageNumbers.Size(); _cc++ {
		_acc, _eda = _gd.Classer.ComponentPageNumbers.Get(_cc)
		if _eda != nil {
			return _df.Wrapf(_eda, _eg, "p\u0061\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027 \u006e\u006f\u0074\u0020\u0066\u006f\u0075nd\u0020\u0069\u006e\u0020t\u0068\u0065\u0020\u0063\u006c\u0061\u0073\u0073\u0065r \u0070\u0061g\u0065\u006e\u0075\u006d\u0062\u0065\u0072\u0073", _cc)
		}
		_agd, _eda = _gd.Classer.ClassIDs.Get(_cc)
		if _eda != nil {
			return _df.Wrapf(_eda, _eg, "\u0063\u0061\u006e\u0027\u0074\u0020\u0067e\u0074\u0020\u0073y\u006d\u0062\u006f\u006c \u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0027\u0025\u0064\u0027\u0020\u0066\u0072\u006f\u006d\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072", _acc)
		}
		if _cf[_agd] == 1 && !_aga {
			_gd._gb[_acc] = append(_gd._gb[_acc], _agd)
		}
	}
	if _eda = _gd.Classer.ComputeLLCorners(); _eda != nil {
		return _df.Wrap(_eda, _eg, "")
	}
	return nil
}
func (_gbf *Page) getWidth() (int, error) {
	const _fed = "\u0067\u0065\u0074\u0057\u0069\u0064\u0074\u0068"
	if _gbf.FinalWidth != 0 {
		return _gbf.FinalWidth, nil
	}
	_afb := _gbf.getPageInformationSegment()
	if _afb == nil {
		return 0, _df.Error(_fed, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_dcb, _gaga := _afb.GetSegmentData()
	if _gaga != nil {
		return 0, _df.Wrap(_gaga, _fed, "")
	}
	_aaedc, _fee := _dcb.(*_ae.PageInformationSegment)
	if !_fee {
		return 0, _df.Errorf(_fed, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _dcb)
	}
	_gbf.FinalWidth = _aaedc.PageBMWidth
	return _gbf.FinalWidth, nil
}
func (_geg *Page) AddEndOfPageSegment() {
	_cfc := &_ae.Header{Type: _ae.TEndOfPage, PageAssociation: _geg.PageNumber}
	_geg.Segments = append(_geg.Segments, _cfc)
}
func (_bdg *Page) GetHeight() (int, error) { return _bdg.getHeight() }
func (_ce *Document) completeClassifiedPages() (_fg error) {
	const _eb = "\u0063\u006f\u006dpl\u0065\u0074\u0065\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0050\u0061\u0067\u0065\u0073"
	if _ce.Classer == nil {
		return nil
	}
	_ce._ab = make([]int, _ce.Classer.UndilatedTemplates.Size())
	for _gbg := 0; _gbg < _ce.Classer.ClassIDs.Size(); _gbg++ {
		_cba, _bd := _ce.Classer.ClassIDs.Get(_gbg)
		if _bd != nil {
			return _df.Wrapf(_bd, _eb, "\u0063\u006c\u0061\u0073s \u0077\u0069\u0074\u0068\u0020\u0069\u0064\u003a\u0020\u0027\u0025\u0064\u0027", _gbg)
		}
		_ce._ab[_cba]++
	}
	var _fa []int
	for _dfg := 0; _dfg < _ce.Classer.UndilatedTemplates.Size(); _dfg++ {
		if _ce.NumberOfPages == 1 || _ce._ab[_dfg] > 1 {
			_fa = append(_fa, _dfg)
		}
	}
	var (
		_fad *Page
		_acg bool
	)
	for _feb, _ee := range *_ce.Classer.ComponentPageNumbers {
		if _fad, _acg = _ce.Pages[_ee]; !_acg {
			return _df.Errorf(_eb, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _feb)
		}
		if _fad.EncodingMethod == GenericEM {
			_a.Log.Error("\u0047\u0065\u006e\u0065\u0072\u0069c\u0020\u0070\u0061g\u0065\u0020\u0077i\u0074\u0068\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u003a \u0027\u0025\u0064\u0027\u0020ma\u0070\u0070\u0065\u0064\u0020\u0061\u0073\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0020\u0070\u0061\u0067\u0065", _feb)
			continue
		}
		_ce._da[_ee] = append(_ce._da[_ee], _feb)
		_gga, _gbb := _ce.Classer.ClassIDs.Get(_feb)
		if _gbb != nil {
			return _df.Wrapf(_gbb, _eb, "\u006e\u006f\u0020\u0073uc\u0068\u0020\u0063\u006c\u0061\u0073\u0073\u0049\u0044\u003a\u0020\u0025\u0064", _feb)
		}
		if _ce._ab[_gga] == 1 && _ce.NumberOfPages != 1 {
			_bg := append(_ce._gb[_ee], _gga)
			_ce._gb[_ee] = _bg
		}
	}
	if _fg = _ce.Classer.ComputeLLCorners(); _fg != nil {
		return _df.Wrap(_fg, _eb, "")
	}
	if _, _fg = _ce.addSymbolDictionary(0, _ce.Classer.UndilatedTemplates, _fa, _ce._ea, false); _fg != nil {
		return _df.Wrap(_fg, _eb, "")
	}
	return nil
}
func (_ggc *Page) addTextRegionSegment(_eef []*_ae.Header, _aaa, _dcd map[int]int, _ecc []int, _cge *_ag.Points, _fga *_ag.Bitmaps, _abed *_gf.IntSlice, _abee *_ag.Boxes, _aff, _ddd int) {
	_gbeag := &_ae.TextRegion{NumberOfSymbols: uint32(_ddd)}
	_gbeag.InitEncode(_aaa, _dcd, _ecc, _cge, _fga, _abed, _abee, _ggc.FinalWidth, _ggc.FinalHeight, _aff)
	_eefb := &_ae.Header{RTSegments: _eef, SegmentData: _gbeag, PageAssociation: _ggc.PageNumber, Type: _ae.TImmediateTextRegion}
	_cec := _ae.TPageInformation
	if _dcd != nil {
		_cec = _ae.TSymbolDictionary
	}
	var _bdd int
	for ; _bdd < len(_ggc.Segments); _bdd++ {
		if _ggc.Segments[_bdd].Type == _cec {
			_bdd++
			break
		}
	}
	_ggc.Segments = append(_ggc.Segments, nil)
	copy(_ggc.Segments[_bdd+1:], _ggc.Segments[_bdd:])
	_ggc.Segments[_bdd] = _eefb
}
func (_acgf *Page) fitsPage(_fcg *_ae.PageInformationSegment, _aaed *_ag.Bitmap) bool {
	return _acgf.countRegions() == 1 && _fcg.DefaultPixelValue == 0 && _fcg.PageBMWidth == _aaed.Width && _fcg.PageBMHeight == _aaed.Height
}
func (_def *Document) determineRandomDataOffsets(_dbf []*_ae.Header, _fdab uint64) {
	if _def.OrganizationType != _ae.ORandom {
		return
	}
	for _, _bce := range _dbf {
		_bce.SegmentDataStartOffset = _fdab
		_fdab += _bce.SegmentDataLength
	}
}
func (_cbd *Page) getHeight() (int, error) {
	const _bfee = "\u0067e\u0074\u0048\u0065\u0069\u0067\u0068t"
	if _cbd.FinalHeight != 0 {
		return _cbd.FinalHeight, nil
	}
	_dcdc := _cbd.getPageInformationSegment()
	if _dcdc == nil {
		return 0, _df.Error(_bfee, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_gagf, _aeec := _dcdc.GetSegmentData()
	if _aeec != nil {
		return 0, _df.Wrap(_aeec, _bfee, "")
	}
	_ged, _fbf := _gagf.(*_ae.PageInformationSegment)
	if !_fbf {
		return 0, _df.Errorf(_bfee, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _gagf)
	}
	if _ged.PageBMHeight == _ca.MaxInt32 {
		_, _aeec = _cbd.GetBitmap()
		if _aeec != nil {
			return 0, _df.Wrap(_aeec, _bfee, "")
		}
	} else {
		_cbd.FinalHeight = _ged.PageBMHeight
	}
	return _cbd.FinalHeight, nil
}
func _gegb(_fafb *Document, _ced int) *Page {
	return &Page{Document: _fafb, PageNumber: _ced, Segments: []*_ae.Header{}}
}
func (_bab *Page) Encode(w _ac.BinaryWriter) (_dddb int, _fadd error) {
	const _ggg = "P\u0061\u0067\u0065\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	var _gaf int
	for _, _feg := range _bab.Segments {
		if _gaf, _fadd = _feg.Encode(w); _fadd != nil {
			return _dddb, _df.Wrap(_fadd, _ggg, "")
		}
		_dddb += _gaf
	}
	return _dddb, nil
}
func (_faf *Document) GetGlobalSegment(i int) (*_ae.Header, error) {
	_fdea, _bcg := _faf.GlobalSegments.GetSegment(i)
	if _bcg != nil {
		return nil, _df.Wrap(_bcg, "\u0047\u0065t\u0047\u006c\u006fb\u0061\u006c\u0053\u0065\u0067\u006d\u0065\u006e\u0074", "")
	}
	return _fdea, nil
}
func (_dcf *Page) GetWidth() (int, error) { return _dcf.getWidth() }
func (_gfc *Page) composePageBitmap() error {
	const _egb = "\u0063\u006f\u006d\u0070\u006f\u0073\u0065\u0050\u0061\u0067\u0065\u0042i\u0074\u006d\u0061\u0070"
	if _gfc.PageNumber == 0 {
		return nil
	}
	_fded := _gfc.getPageInformationSegment()
	if _fded == nil {
		return _df.Error(_egb, "\u0070\u0061\u0067e \u0069\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069o\u006e \u0073e\u0067m\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_fecg, _cbbg := _fded.GetSegmentData()
	if _cbbg != nil {
		return _cbbg
	}
	_fdfg, _cae := _fecg.(*_ae.PageInformationSegment)
	if !_cae {
		return _df.Error(_egb, "\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006da\u0074\u0069\u006f\u006e \u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065")
	}
	if _cbbg = _gfc.createPage(_fdfg); _cbbg != nil {
		return _df.Wrap(_cbbg, _egb, "")
	}
	_gfc.clearSegmentData()
	return nil
}
func (_fgg *Page) getCombinationOperator(_debb *_ae.PageInformationSegment, _ecf _ag.CombinationOperator) _ag.CombinationOperator {
	if _debb.CombinationOperatorOverrideAllowed() {
		return _ecf
	}
	return _debb.CombinationOperator()
}
func (_dcbg *Page) nextSegmentNumber() uint32 { return _dcbg.Document.nextSegmentNumber() }
func _dbg(_cdfd *_ac.Reader, _fcf *Globals) (*Document, error) {
	_aad := &Document{Pages: make(map[int]*Page), InputStream: _cdfd, OrganizationType: _ae.OSequential, NumberOfPagesUnknown: true, GlobalSegments: _fcf, _e: 9}
	if _aad.GlobalSegments == nil {
		_aad.GlobalSegments = &Globals{}
	}
	if _accg := _aad.mapData(); _accg != nil {
		return nil, _accg
	}
	return _aad, nil
}
func (_abg *Globals) GetSymbolDictionary() (*_ae.Header, error) {
	const _gbeg = "G\u006c\u006f\u0062\u0061\u006c\u0073.\u0047\u0065\u0074\u0053\u0079\u006d\u0062\u006f\u006cD\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079"
	if _abg == nil {
		return nil, _df.Error(_gbeg, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_abg.Segments) == 0 {
		return nil, _df.Error(_gbeg, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	for _, _ddf := range _abg.Segments {
		if _ddf.Type == _ae.TSymbolDictionary {
			return _ddf, nil
		}
	}
	return nil, _df.Error(_gbeg, "\u0067\u006c\u006fba\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0020d\u0069c\u0074i\u006fn\u0061\u0072\u0079\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}
func (_bfa *Page) GetBitmap() (_gcf *_ag.Bitmap, _edg error) {
	_a.Log.Trace(_cb.Sprintf("\u005b\u0050\u0041G\u0045\u005d\u005b\u0023%\u0064\u005d\u0020\u0047\u0065\u0074\u0042i\u0074\u006d\u0061\u0070\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e", _bfa.PageNumber))
	defer func() {
		if _edg != nil {
			_a.Log.Trace(_cb.Sprintf("\u005b\u0050\u0041\u0047\u0045\u005d\u005b\u0023\u0025\u0064\u005d\u0020\u0047\u0065\u0074B\u0069t\u006d\u0061\u0070\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _bfa.PageNumber, _edg))
		} else {
			_a.Log.Trace(_cb.Sprintf("\u005b\u0050\u0041\u0047\u0045\u005d\u005b\u0023\u0025\u0064]\u0020\u0047\u0065\u0074\u0042\u0069\u0074m\u0061\u0070\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064", _bfa.PageNumber))
		}
	}()
	if _bfa.Bitmap != nil {
		return _bfa.Bitmap, nil
	}
	_edg = _bfa.composePageBitmap()
	if _edg != nil {
		return nil, _edg
	}
	return _bfa.Bitmap, nil
}
func (_efa *Document) produceClassifiedPage(_ba *Page, _ggb *_ae.Header) (_bgg error) {
	const _fd = "p\u0072\u006f\u0064\u0075ce\u0043l\u0061\u0073\u0073\u0069\u0066i\u0065\u0064\u0050\u0061\u0067\u0065"
	var _ebb map[int]int
	_bgf := _efa._cbc
	_bad := []*_ae.Header{_ggb}
	if len(_efa._gb[_ba.PageNumber]) > 0 {
		_ebb = map[int]int{}
		_fdf, _efd := _efa.addSymbolDictionary(_ba.PageNumber, _efa.Classer.UndilatedTemplates, _efa._gb[_ba.PageNumber], _ebb, false)
		if _efd != nil {
			return _df.Wrap(_efd, _fd, "")
		}
		_bad = append(_bad, _fdf)
		_bgf += len(_efa._gb[_ba.PageNumber])
	}
	_cga := _efa._da[_ba.PageNumber]
	_a.Log.Debug("P\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020c\u006f\u006d\u0070\u0073: \u0025\u0076", _ba.PageNumber, _cga)
	_ba.addTextRegionSegment(_bad, _efa._ea, _ebb, _efa._da[_ba.PageNumber], _efa.Classer.PtaLL, _efa.Classer.UndilatedTemplates, _efa.Classer.ClassIDs, nil, _dc(_bgf), len(_efa._da[_ba.PageNumber]))
	return nil
}

type Page struct {
	Segments           []*_ae.Header
	PageNumber         int
	Bitmap             *_ag.Bitmap
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

func (_ebf *Page) getResolutionY() (int, error) {
	const _aadd = "\u0067\u0065\u0074\u0052\u0065\u0073\u006f\u006c\u0075t\u0069\u006f\u006e\u0059"
	if _ebf.ResolutionY != 0 {
		return _ebf.ResolutionY, nil
	}
	_bfae := _ebf.getPageInformationSegment()
	if _bfae == nil {
		return 0, _df.Error(_aadd, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_cdga, _bee := _bfae.GetSegmentData()
	if _bee != nil {
		return 0, _df.Wrap(_bee, _aadd, "")
	}
	_ecfg, _ebg := _cdga.(*_ae.PageInformationSegment)
	if !_ebg {
		return 0, _df.Errorf(_aadd, "\u0070\u0061\u0067\u0065\u0020\u0069\u006e\u0066o\u0072\u006d\u0061ti\u006f\u006e\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0074\u0079\u0070\u0065\u003a\u0027%\u0054\u0027", _cdga)
	}
	_ebf.ResolutionY = _ecfg.ResolutionY
	return _ebf.ResolutionY, nil
}
func (_gcb *Page) createStripedPage(_efc *_ae.PageInformationSegment) error {
	const _dfgc = "\u0063\u0072\u0065\u0061\u0074\u0065\u0053\u0074\u0072\u0069\u0070\u0065d\u0050\u0061\u0067\u0065"
	_ade, _ggce := _gcb.collectPageStripes()
	if _ggce != nil {
		return _df.Wrap(_ggce, _dfgc, "")
	}
	var _bbe int
	for _, _aadf := range _ade {
		if _add, _aee := _aadf.(*_ae.EndOfStripe); _aee {
			_bbe = _add.LineNumber() + 1
		} else {
			_edga := _aadf.(_ae.Regioner)
			_eec := _edga.GetRegionInfo()
			_abdbg := _gcb.getCombinationOperator(_efc, _eec.CombinaionOperator)
			_gcad, _bdcc := _edga.GetRegionBitmap()
			if _bdcc != nil {
				return _df.Wrap(_bdcc, _dfgc, "")
			}
			_bdcc = _ag.Blit(_gcad, _gcb.Bitmap, int(_eec.XLocation), _bbe, _abdbg)
			if _bdcc != nil {
				return _df.Wrap(_bdcc, _dfgc, "")
			}
		}
	}
	return nil
}
func (_gaa *Document) parseFileHeader() error {
	const _gfe = "\u0070a\u0072s\u0065\u0046\u0069\u006c\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_, _ceb := _gaa.InputStream.Seek(8, _g.SeekStart)
	if _ceb != nil {
		return _df.Wrap(_ceb, _gfe, "\u0069\u0064")
	}
	_, _ceb = _gaa.InputStream.ReadBits(5)
	if _ceb != nil {
		return _df.Wrap(_ceb, _gfe, "\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0064\u0020\u0062\u0069\u0074\u0073")
	}
	_fecd, _ceb := _gaa.InputStream.ReadBit()
	if _ceb != nil {
		return _df.Wrap(_ceb, _gfe, "\u0065x\u0074e\u006e\u0064\u0065\u0064\u0020t\u0065\u006dp\u006c\u0061\u0074\u0065\u0073")
	}
	if _fecd == 1 {
		_gaa.GBUseExtTemplate = true
	}
	_fecd, _ceb = _gaa.InputStream.ReadBit()
	if _ceb != nil {
		return _df.Wrap(_ceb, _gfe, "\u0075\u006e\u006b\u006eow\u006e\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	if _fecd != 1 {
		_gaa.NumberOfPagesUnknown = false
	}
	_fecd, _ceb = _gaa.InputStream.ReadBit()
	if _ceb != nil {
		return _df.Wrap(_ceb, _gfe, "\u006f\u0072\u0067\u0061\u006e\u0069\u007a\u0061\u0074\u0069\u006f\u006e \u0074\u0079\u0070\u0065")
	}
	_gaa.OrganizationType = _ae.OrganizationType(_fecd)
	if !_gaa.NumberOfPagesUnknown {
		_gaa.NumberOfPages, _ceb = _gaa.InputStream.ReadUint32()
		if _ceb != nil {
			return _df.Wrap(_ceb, _gfe, "\u006eu\u006db\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0061\u0067\u0065\u0073")
		}
		_gaa._e = 13
	}
	return nil
}
func (_fcd *Page) createNormalPage(_deb *_ae.PageInformationSegment) error {
	const _gbbd = "\u0063\u0072e\u0061\u0074\u0065N\u006f\u0072\u006d\u0061\u006c\u0050\u0061\u0067\u0065"
	_fcd.Bitmap = _ag.New(_deb.PageBMWidth, _deb.PageBMHeight)
	if _deb.DefaultPixelValue != 0 {
		_fcd.Bitmap.SetDefaultPixel()
	}
	for _, _gca := range _fcd.Segments {
		switch _gca.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_a.Log.Trace("\u0047\u0065\u0074\u0074in\u0067\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0064", _gca.SegmentNumber)
			_gad, _caa := _gca.GetSegmentData()
			if _caa != nil {
				return _caa
			}
			_afdc, _bcf := _gad.(_ae.Regioner)
			if !_bcf {
				_a.Log.Debug("\u0053\u0065g\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0054\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0052\u0065\u0067\u0069on\u0065\u0072", _gad)
				return _df.Errorf(_gbbd, "i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062i\u0067\u0032\u0020\u0073\u0065\u0067\u006den\u0074\u0020\u0074\u0079p\u0065\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061 R\u0065\u0067i\u006f\u006e\u0065\u0072\u003a\u0020\u0025\u0054", _gad)
			}
			_ece, _caa := _afdc.GetRegionBitmap()
			if _caa != nil {
				return _df.Wrap(_caa, _gbbd, "")
			}
			if _fcd.fitsPage(_deb, _ece) {
				_fcd.Bitmap = _ece
			} else {
				_fafd := _afdc.GetRegionInfo()
				_gagg := _fcd.getCombinationOperator(_deb, _fafd.CombinaionOperator)
				_caa = _ag.Blit(_ece, _fcd.Bitmap, int(_fafd.XLocation), int(_fafd.YLocation), _gagg)
				if _caa != nil {
					return _df.Wrap(_caa, _gbbd, "")
				}
			}
		}
	}
	return nil
}
func (_gfec *Globals) AddSegment(segment *_ae.Header) {
	_gfec.Segments = append(_gfec.Segments, segment)
}
func (_eab *Globals) GetSegmentByIndex(index int) (*_ae.Header, error) {
	const _bed = "\u0047l\u006f\u0062\u0061\u006cs\u002e\u0047\u0065\u0074\u0053e\u0067m\u0065n\u0074\u0042\u0079\u0049\u006e\u0064\u0065x"
	if _eab == nil {
		return nil, _df.Error(_bed, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_eab.Segments) == 0 {
		return nil, _df.Error(_bed, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	if index > len(_eab.Segments)-1 {
		return nil, _df.Error(_bed, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	return _eab.Segments[index], nil
}
func (_cg *Document) produceClassifiedPages() (_dd error) {
	const _abd = "\u0070\u0072\u006f\u0064uc\u0065\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0050\u0061\u0067e\u0073"
	if _cg.Classer == nil {
		return nil
	}
	var (
		_bdc *Page
		_bgd bool
		_bf  *_ae.Header
	)
	for _bda := 1; _bda <= int(_cg.NumberOfPages); _bda++ {
		if _bdc, _bgd = _cg.Pages[_bda]; !_bgd {
			return _df.Errorf(_abd, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _bda)
		}
		if _bdc.EncodingMethod == GenericEM {
			continue
		}
		if _bf == nil {
			if _bf, _dd = _cg.GlobalSegments.GetSymbolDictionary(); _dd != nil {
				return _df.Wrap(_dd, _abd, "")
			}
		}
		if _dd = _cg.produceClassifiedPage(_bdc, _bf); _dd != nil {
			return _df.Wrapf(_dd, _abd, "\u0070\u0061\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _bda)
		}
	}
	return nil
}
func (_dea *Page) AddPageInformationSegment() {
	_gff := &_ae.PageInformationSegment{PageBMWidth: _dea.FinalWidth, PageBMHeight: _dea.FinalHeight, ResolutionX: _dea.ResolutionX, ResolutionY: _dea.ResolutionY, IsLossless: _dea.IsLossless}
	if _dea.BlackIsOne {
		_gff.DefaultPixelValue = uint8(0x1)
	}
	_cdbc := &_ae.Header{PageAssociation: _dea.PageNumber, SegmentDataLength: uint64(_gff.Size()), SegmentData: _gff, Type: _ae.TPageInformation}
	_dea.Segments = append(_dea.Segments, _cdbc)
}
func (_gggc *Page) GetSegment(number int) (*_ae.Header, error) {
	const _ebd = "\u0050a\u0067e\u002e\u0047\u0065\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074"
	for _, _ede := range _gggc.Segments {
		if _ede.SegmentNumber == uint32(number) {
			return _ede, nil
		}
	}
	_cce := make([]uint32, len(_gggc.Segments))
	for _afe, _cfgd := range _gggc.Segments {
		_cce[_afe] = _cfgd.SegmentNumber
	}
	return nil, _df.Errorf(_ebd, "\u0073e\u0067\u006d\u0065n\u0074\u0020\u0077i\u0074h \u006e\u0075\u006d\u0062\u0065\u0072\u003a \u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0067\u0065\u003a\u0020'%\u0064'\u002e\u0020\u004b\u006e\u006f\u0077n\u0020\u0073\u0065\u0067\u006de\u006e\u0074\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0073\u003a \u0025\u0076", number, _gggc.PageNumber, _cce)
}
func (_fda *Document) GetNumberOfPages() (uint32, error) {
	if _fda.NumberOfPagesUnknown || _fda.NumberOfPages == 0 {
		if len(_fda.Pages) == 0 {
			if _cag := _fda.mapData(); _cag != nil {
				return 0, _df.Wrap(_cag, "\u0044o\u0063\u0075\u006d\u0065n\u0074\u002e\u0047\u0065\u0074N\u0075m\u0062e\u0072\u004f\u0066\u0050\u0061\u0067\u0065s", "")
			}
		}
		return uint32(len(_fda.Pages)), nil
	}
	return _fda.NumberOfPages, nil
}
func (_cdb *Globals) GetSegment(segmentNumber int) (*_ae.Header, error) {
	const _fbc = "\u0047l\u006fb\u0061\u006c\u0073\u002e\u0047e\u0074\u0053e\u0067\u006d\u0065\u006e\u0074"
	if _cdb == nil {
		return nil, _df.Error(_fbc, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_cdb.Segments) == 0 {
		return nil, _df.Error(_fbc, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	var _aae *_ae.Header
	for _, _aae = range _cdb.Segments {
		if _aae.SegmentNumber == uint32(segmentNumber) {
			break
		}
	}
	if _aae == nil {
		return nil, _df.Error(_fbc, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	return _aae, nil
}
func (_bdf *Page) getPageInformationSegment() *_ae.Header {
	for _, _adg := range _bdf.Segments {
		if _adg.Type == _ae.TPageInformation {
			return _adg
		}
	}
	_a.Log.Debug("\u0050\u0061\u0067\u0065\u0020\u0069\u006e\u0066o\u0072\u006d\u0061ti\u006f\u006e\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0066o\u0072\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0073\u002e", _bdf)
	return nil
}
func (_dgf *Document) Encode() (_agf []byte, _cbb error) {
	const _bbg = "\u0044o\u0063u\u006d\u0065\u006e\u0074\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	var _aed, _agc int
	if _dgf.FullHeaders {
		if _aed, _cbb = _dgf.encodeFileHeader(_dgf._dfa); _cbb != nil {
			return nil, _df.Wrap(_cbb, _bbg, "")
		}
	}
	var (
		_aeff bool
		_efb  *_ae.Header
		_gbea *Page
	)
	if _cbb = _dgf.completeClassifiedPages(); _cbb != nil {
		return nil, _df.Wrap(_cbb, _bbg, "")
	}
	if _cbb = _dgf.produceClassifiedPages(); _cbb != nil {
		return nil, _df.Wrap(_cbb, _bbg, "")
	}
	if _dgf.GlobalSegments != nil {
		for _, _efb = range _dgf.GlobalSegments.Segments {
			if _cbb = _dgf.encodeSegment(_efb, &_aed); _cbb != nil {
				return nil, _df.Wrap(_cbb, _bbg, "")
			}
		}
	}
	for _gba := 1; _gba <= int(_dgf.NumberOfPages); _gba++ {
		if _gbea, _aeff = _dgf.Pages[_gba]; !_aeff {
			return nil, _df.Errorf(_bbg, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _gba)
		}
		for _, _efb = range _gbea.Segments {
			if _cbb = _dgf.encodeSegment(_efb, &_aed); _cbb != nil {
				return nil, _df.Wrap(_cbb, _bbg, "")
			}
		}
	}
	if _dgf.FullHeaders {
		if _agc, _cbb = _dgf.encodeEOFHeader(_dgf._dfa); _cbb != nil {
			return nil, _df.Wrap(_cbb, _bbg, "")
		}
		_aed += _agc
	}
	_agf = _dgf._dfa.Data()
	if len(_agf) != _aed {
		_a.Log.Debug("\u0042\u0079\u0074\u0065\u0073 \u0077\u0072\u0069\u0074\u0074\u0065\u006e \u0028\u006e\u0029\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0064\u0061\u0074\u0061\u0020\u0065\u006e\u0063\u006fd\u0065\u0064\u003a\u0020\u0027\u0025d\u0027", _aed, len(_agf))
	}
	return _agf, nil
}
func (_ec *Document) reachedEOF(_aa int64) (bool, error) {
	const _dbfa = "\u0072\u0065\u0061\u0063\u0068\u0065\u0064\u0045\u004f\u0046"
	_, _dfbd := _ec.InputStream.Seek(_aa, _g.SeekStart)
	if _dfbd != nil {
		_a.Log.Debug("\u0072\u0065\u0061c\u0068\u0065\u0064\u0045\u004f\u0046\u0020\u002d\u0020\u0064\u002e\u0049\u006e\u0070\u0075\u0074\u0053\u0074\u0072\u0065\u0061\u006d\u002e\u0053\u0065\u0065\u006b\u0020\u0066a\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _dfbd)
		return false, _df.Wrap(_dfbd, _dbfa, "\u0069n\u0070\u0075\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020s\u0065\u0065\u006b\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_, _dfbd = _ec.InputStream.ReadBits(32)
	if _dfbd == _g.EOF {
		return true, nil
	} else if _dfbd != nil {
		return false, _df.Wrap(_dfbd, _dbfa, "")
	}
	return false, nil
}

type Globals struct{ Segments []*_ae.Header }

func InitEncodeDocument(fullHeaders bool) *Document {
	return &Document{FullHeaders: fullHeaders, _dfa: _ac.BufferedMSB(), Pages: map[int]*Page{}, _gb: map[int][]int{}, _ea: map[int]int{}, _da: map[int][]int{}}
}
func (_bdgd *Page) getResolutionX() (int, error) {
	const _bgdc = "\u0067\u0065\u0074\u0052\u0065\u0073\u006f\u006c\u0075t\u0069\u006f\u006e\u0058"
	if _bdgd.ResolutionX != 0 {
		return _bdgd.ResolutionX, nil
	}
	_ggcd := _bdgd.getPageInformationSegment()
	if _ggcd == nil {
		return 0, _df.Error(_bgdc, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_baa, _cagg := _ggcd.GetSegmentData()
	if _cagg != nil {
		return 0, _df.Wrap(_cagg, _bgdc, "")
	}
	_dca, _bedg := _baa.(*_ae.PageInformationSegment)
	if !_bedg {
		return 0, _df.Errorf(_bgdc, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _baa)
	}
	_bdgd.ResolutionX = _dca.ResolutionX
	return _bdgd.ResolutionX, nil
}
func (_ccd *Document) mapData() error {
	const _fgc = "\u006da\u0070\u0044\u0061\u0074\u0061"
	var (
		_abdc []*_ae.Header
		_cfb  int64
		_bbfg _ae.Type
	)
	_dde, _fb := _ccd.isFileHeaderPresent()
	if _fb != nil {
		return _df.Wrap(_fb, _fgc, "")
	}
	if _dde {
		if _fb = _ccd.parseFileHeader(); _fb != nil {
			return _df.Wrap(_fb, _fgc, "")
		}
		_cfb += int64(_ccd._e)
		_ccd.FullHeaders = true
	}
	var (
		_be   *Page
		_defd bool
	)
	for _bbfg != 51 && !_defd {
		_efbb, _bcb := _ae.NewHeader(_ccd, _ccd.InputStream, _cfb, _ccd.OrganizationType)
		if _bcb != nil {
			return _df.Wrap(_bcb, _fgc, "")
		}
		_a.Log.Trace("\u0044\u0065c\u006f\u0064\u0069\u006eg\u0020\u0073e\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0025\u0064\u002c\u0020\u0054\u0079\u0070e\u003a\u0020\u0025\u0073", _efbb.SegmentNumber, _efbb.Type)
		_bbfg = _efbb.Type
		if _bbfg != _ae.TEndOfFile {
			if _efbb.PageAssociation != 0 {
				_be = _ccd.Pages[_efbb.PageAssociation]
				if _be == nil {
					_be = _gegb(_ccd, _efbb.PageAssociation)
					_ccd.Pages[_efbb.PageAssociation] = _be
					if _ccd.NumberOfPagesUnknown {
						_ccd.NumberOfPages++
					}
				}
				_be.Segments = append(_be.Segments, _efbb)
			} else {
				_ccd.GlobalSegments.AddSegment(_efbb)
			}
		}
		_abdc = append(_abdc, _efbb)
		_cfb = _ccd.InputStream.AbsolutePosition()
		if _ccd.OrganizationType == _ae.OSequential {
			_cfb += int64(_efbb.SegmentDataLength)
		}
		_defd, _bcb = _ccd.reachedEOF(_cfb)
		if _bcb != nil {
			_a.Log.Debug("\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0072\u0065\u0061\u0063h\u0065\u0064\u0020\u0045\u004f\u0046\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _bcb)
			return _df.Wrap(_bcb, _fgc, "")
		}
	}
	_ccd.determineRandomDataOffsets(_abdc, uint64(_cfb))
	return nil
}

type EncodingMethod int

func (_bcd *Page) collectPageStripes() (_cdba []_ae.Segmenter, _cecg error) {
	const _dfe = "\u0063o\u006cl\u0065\u0063\u0074\u0050\u0061g\u0065\u0053t\u0072\u0069\u0070\u0065\u0073"
	var _dgc _ae.Segmenter
	for _, _eabd := range _bcd.Segments {
		switch _eabd.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_dgc, _cecg = _eabd.GetSegmentData()
			if _cecg != nil {
				return nil, _df.Wrap(_cecg, _dfe, "")
			}
			_cdba = append(_cdba, _dgc)
		case 50:
			_dgc, _cecg = _eabd.GetSegmentData()
			if _cecg != nil {
				return nil, _cecg
			}
			_bbd, _fgf := _dgc.(*_ae.EndOfStripe)
			if !_fgf {
				return nil, _df.Errorf(_dfe, "\u0045\u006e\u0064\u004f\u0066\u0053\u0074\u0072\u0069\u0070\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u006f\u0066\u0020\u0076\u0061l\u0069\u0064\u0020\u0074\u0079p\u0065\u003a \u0027\u0025\u0054\u0027", _dgc)
			}
			_cdba = append(_cdba, _bbd)
			_bcd.FinalHeight = _bbd.LineNumber()
		}
	}
	return _cdba, nil
}
func (_fdeg *Document) nextPageNumber() uint32 { _fdeg.NumberOfPages++; return _fdeg.NumberOfPages }
func (_cd *Document) AddClassifiedPage(bm *_ag.Bitmap, method _b.Method) (_bb error) {
	const _ef = "\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0041\u0064d\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064P\u0061\u0067\u0065"
	if !_cd.FullHeaders && _cd.NumberOfPages != 0 {
		return _df.Error(_ef, "\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0061\u006c\u0072\u0065a\u0064\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0070\u0061\u0067\u0065\u002e\u0020\u0046\u0069\u006c\u0065\u004d\u006f\u0064\u0065\u0020\u0064\u0069\u0073\u0061\u006c\u006c\u006f\u0077\u0073\u0020\u0061\u0064\u0064i\u006e\u0067\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e \u006f\u006e\u0065\u0020\u0070\u0061g\u0065")
	}
	if _cd.Classer == nil {
		if _cd.Classer, _bb = _b.Init(_b.DefaultSettings()); _bb != nil {
			return _df.Wrap(_bb, _ef, "")
		}
	}
	_aec := int(_cd.nextPageNumber())
	_fe := &Page{Segments: []*_ae.Header{}, Bitmap: bm, Document: _cd, FinalHeight: bm.Height, FinalWidth: bm.Width, PageNumber: _aec}
	_cd.Pages[_aec] = _fe
	switch method {
	case _b.RankHaus:
		_fe.EncodingMethod = RankHausEM
	case _b.Correlation:
		_fe.EncodingMethod = CorrelationEM
	}
	_fe.AddPageInformationSegment()
	if _bb = _cd.Classer.AddPage(bm, _aec, method); _bb != nil {
		return _df.Wrap(_bb, _ef, "")
	}
	if _cd.FullHeaders {
		_fe.AddEndOfPageSegment()
	}
	return nil
}
func (_ed *Document) addSymbolDictionary(_agb int, _fgb *_ag.Bitmaps, _bfb []int, _ga map[int]int, _bc bool) (*_ae.Header, error) {
	const _dfb = "\u0061\u0064\u0064\u0053ym\u0062\u006f\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079"
	_fc := &_ae.SymbolDictionary{}
	if _caf := _fc.InitEncode(_fgb, _bfb, _ga, _bc); _caf != nil {
		return nil, _caf
	}
	_gbe := &_ae.Header{Type: _ae.TSymbolDictionary, PageAssociation: _agb, SegmentData: _fc}
	if _agb == 0 {
		if _ed.GlobalSegments == nil {
			_ed.GlobalSegments = &Globals{}
		}
		_ed.GlobalSegments.AddSegment(_gbe)
		return _gbe, nil
	}
	_gae, _cdg := _ed.Pages[_agb]
	if !_cdg {
		return nil, _df.Errorf(_dfb, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _agb)
	}
	var (
		_dff int
		_bfd *_ae.Header
	)
	for _dff, _bfd = range _gae.Segments {
		if _bfd.Type == _ae.TPageInformation {
			break
		}
	}
	_dff++
	_gae.Segments = append(_gae.Segments, nil)
	copy(_gae.Segments[_dff+1:], _gae.Segments[_dff:])
	_gae.Segments[_dff] = _gbe
	return _gbe, nil
}
func (_bfe *Page) GetResolutionY() (int, error) { return _bfe.getResolutionY() }
func (_dbea *Page) createPage(_ggd *_ae.PageInformationSegment) error {
	var _aag error
	if !_ggd.IsStripe || _ggd.PageBMHeight != -1 {
		_aag = _dbea.createNormalPage(_ggd)
	} else {
		_aag = _dbea.createStripedPage(_ggd)
	}
	return _aag
}
func (_defb *Page) lastSegmentNumber() (_ccde uint32, _egf error) {
	const _fce = "\u006c\u0061\u0073\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u004eu\u006d\u0062\u0065\u0072"
	if len(_defb.Segments) == 0 {
		return _ccde, _df.Errorf(_fce, "\u006e\u006f\u0020se\u0067\u006d\u0065\u006e\u0074\u0073\u0020\u0066\u006fu\u006ed\u0020i\u006e \u0074\u0068\u0065\u0020\u0070\u0061\u0067\u0065\u0020\u0027\u0025\u0064\u0027", _defb.PageNumber)
	}
	return _defb.Segments[len(_defb.Segments)-1].SegmentNumber, nil
}
func (_ceg *Document) encodeSegment(_bde *_ae.Header, _cee *int) error {
	const _ead = "\u0065\u006e\u0063\u006f\u0064\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074"
	_bde.SegmentNumber = _ceg.nextSegmentNumber()
	_ad, _ebe := _bde.Encode(_ceg._dfa)
	if _ebe != nil {
		return _df.Wrapf(_ebe, _ead, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", _bde.SegmentNumber)
	}
	*_cee += _ad
	return nil
}

type Document struct {
	Pages                map[int]*Page
	NumberOfPagesUnknown bool
	NumberOfPages        uint32
	GBUseExtTemplate     bool
	InputStream          *_ac.Reader
	GlobalSegments       *Globals
	OrganizationType     _ae.OrganizationType
	Classer              *_b.Classer
	XRes, YRes           int
	FullHeaders          bool
	CurrentSegmentNumber uint32
	AverageTemplates     *_ag.Bitmaps
	BaseIndexes          []int
	Refinement           bool
	RefineLevel          int
	_e                   uint8
	_dfa                 *_ac.BufferedWriter
	EncodeGlobals        bool
	_cbc                 int
	_gb                  map[int][]int
	_da                  map[int][]int
	_ab                  []int
	_ea                  map[int]int
}

const (
	GenericEM EncodingMethod = iota
	CorrelationEM
	RankHausEM
)

func (_cdcf *Page) countRegions() int {
	var _bdac int
	for _, _efae := range _cdcf.Segments {
		switch _efae.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_bdac++
		}
	}
	return _bdac
}
func (_gag *Document) nextSegmentNumber() uint32 {
	_dba := _gag.CurrentSegmentNumber
	_gag.CurrentSegmentNumber++
	return _dba
}
func (_dbe *Page) AddGenericRegion(bm *_ag.Bitmap, xloc, yloc, template int, tp _ae.Type, duplicateLineRemoval bool) error {
	const _acgb = "P\u0061\u0067\u0065\u002eAd\u0064G\u0065\u006e\u0065\u0072\u0069c\u0052\u0065\u0067\u0069\u006f\u006e"
	_cegb := &_ae.GenericRegion{}
	if _aede := _cegb.InitEncode(bm, xloc, yloc, template, duplicateLineRemoval); _aede != nil {
		return _df.Wrap(_aede, _acgb, "")
	}
	_cgd := &_ae.Header{Type: _ae.TImmediateGenericRegion, PageAssociation: _dbe.PageNumber, SegmentData: _cegb}
	_dbe.Segments = append(_dbe.Segments, _cgd)
	return nil
}
func DecodeDocument(input *_ac.Reader, globals *Globals) (*Document, error) {
	return _dbg(input, globals)
}

var _gg = []byte{0x97, 0x4A, 0x42, 0x32, 0x0D, 0x0A, 0x1A, 0x0A}

func (_dg *Document) AddGenericPage(bm *_ag.Bitmap, duplicateLineRemoval bool) (_aef error) {
	const _gc = "\u0044\u006f\u0063um\u0065\u006e\u0074\u002e\u0041\u0064\u0064\u0047\u0065\u006e\u0065\u0072\u0069\u0063\u0050\u0061\u0067\u0065"
	if !_dg.FullHeaders && _dg.NumberOfPages != 0 {
		return _df.Error(_gc, "\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0061\u006c\u0072\u0065a\u0064\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0070\u0061\u0067\u0065\u002e\u0020\u0046\u0069\u006c\u0065\u004d\u006f\u0064\u0065\u0020\u0064\u0069\u0073\u0061\u006c\u006c\u006f\u0077\u0073\u0020\u0061\u0064\u0064i\u006e\u0067\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e \u006f\u006e\u0065\u0020\u0070\u0061g\u0065")
	}
	_abe := &Page{Segments: []*_ae.Header{}, Bitmap: bm, Document: _dg, FinalHeight: bm.Height, FinalWidth: bm.Width, IsLossless: true, BlackIsOne: bm.Color == _ag.Chocolate}
	_abe.PageNumber = int(_dg.nextPageNumber())
	_dg.Pages[_abe.PageNumber] = _abe
	bm.InverseData()
	_abe.AddPageInformationSegment()
	if _aef = _abe.AddGenericRegion(bm, 0, 0, 0, _ae.TImmediateGenericRegion, duplicateLineRemoval); _aef != nil {
		return _df.Wrap(_aef, _gc, "")
	}
	if _dg.FullHeaders {
		_abe.AddEndOfPageSegment()
	}
	return nil
}
func (_af *Document) GetPage(pageNumber int) (_ae.Pager, error) {
	const _cfd = "\u0044\u006fc\u0075\u006d\u0065n\u0074\u002e\u0047\u0065\u0074\u0050\u0061\u0067\u0065"
	if pageNumber < 0 {
		_a.Log.Debug("\u004a\u0042\u0049\u00472\u0020\u0050\u0061\u0067\u0065\u0020\u002d\u0020\u0047e\u0074\u0050\u0061\u0067\u0065\u003a\u0020\u0025\u0064\u002e\u0020\u0050\u0061\u0067\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062e\u0020\u006c\u006f\u0077\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030\u002e\u0020\u0025\u0073", pageNumber, _f.Stack())
		return nil, _df.Errorf(_cfd, "\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032\u0020d\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u002d\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064 \u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u003a\u0020\u0025\u0064", pageNumber)
	}
	if pageNumber > len(_af.Pages) {
		_a.Log.Debug("\u0050\u0061\u0067\u0065 n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0064\u002e\u0020%\u0073", pageNumber, _f.Stack())
		return nil, _df.Error(_cfd, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u002d\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_gfb, _caff := _af.Pages[pageNumber]
	if !_caff {
		_a.Log.Debug("\u0050\u0061\u0067\u0065 n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0064\u002e\u0020%\u0073", pageNumber, _f.Stack())
		return nil, _df.Errorf(_cfd, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u002d\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	return _gfb, nil
}
func (_bfg *Document) encodeFileHeader(_efac _ac.BinaryWriter) (_dbb int, _ada error) {
	const _aba = "\u0065\u006ec\u006f\u0064\u0065F\u0069\u006c\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_dbb, _ada = _efac.Write(_gg)
	if _ada != nil {
		return _dbb, _df.Wrap(_ada, _aba, "\u0069\u0064")
	}
	if _ada = _efac.WriteByte(0x01); _ada != nil {
		return _dbb, _df.Wrap(_ada, _aba, "\u0066\u006c\u0061g\u0073")
	}
	_dbb++
	_dgd := make([]byte, 4)
	_d.BigEndian.PutUint32(_dgd, _bfg.NumberOfPages)
	_ebbg, _ada := _efac.Write(_dgd)
	if _ada != nil {
		return _ebbg, _df.Wrap(_ada, _aba, "p\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_dbb += _ebbg
	return _dbb, nil
}
func (_bbf *Document) isFileHeaderPresent() (bool, error) {
	_bbf.InputStream.Mark()
	for _, _ebbgc := range _gg {
		_bfdc, _eeg := _bbf.InputStream.ReadByte()
		if _eeg != nil {
			return false, _eeg
		}
		if _ebbgc != _bfdc {
			_bbf.InputStream.Reset()
			return false, nil
		}
	}
	_bbf.InputStream.Reset()
	return true, nil
}
func (_aaf *Page) clearSegmentData() {
	for _ggcg := range _aaf.Segments {
		_aaf.Segments[_ggcg].CleanSegmentData()
	}
}
func _dc(_fde int) int {
	_dac := 0
	_de := (_fde & (_fde - 1)) == 0
	_fde >>= 1
	for ; _fde != 0; _fde >>= 1 {
		_dac++
	}
	if _de {
		return _dac
	}
	return _dac + 1
}
