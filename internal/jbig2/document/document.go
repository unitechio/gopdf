package document

import (
	_b "encoding/binary"
	_g "fmt"
	_fe "io"
	_f "math"
	_c "runtime/debug"

	_cb "unitechio/gopdf/gopdf/common"
	_a "unitechio/gopdf/gopdf/internal/bitwise"
	_eg "unitechio/gopdf/gopdf/internal/jbig2/basic"
	_ff "unitechio/gopdf/gopdf/internal/jbig2/bitmap"
	_ca "unitechio/gopdf/gopdf/internal/jbig2/document/segments"
	_d "unitechio/gopdf/gopdf/internal/jbig2/encoder/classer"
	_gd "unitechio/gopdf/gopdf/internal/jbig2/errors"
)

func _gad(_dee *Document, _fde int) *Page {
	return &Page{Document: _dee, PageNumber: _fde, Segments: []*_ca.Header{}}
}

func (_gaa *Document) determineRandomDataOffsets(_eee []*_ca.Header, _ggb uint64) {
	if _gaa.OrganizationType != _ca.ORandom {
		return
	}
	for _, _gbc := range _eee {
		_gbc.SegmentDataStartOffset = _ggb
		_ggb += _gbc.SegmentDataLength
	}
}

func InitEncodeDocument(fullHeaders bool) *Document {
	return &Document{FullHeaders: fullHeaders, _fc: _a.BufferedMSB(), Pages: map[int]*Page{}, _egf: map[int][]int{}, _ad: map[int]int{}, _ac: map[int][]int{}}
}

func (_fag *Page) GetSegment(number int) (*_ca.Header, error) {
	const _cgg = "\u0050a\u0067e\u002e\u0047\u0065\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074"
	for _, _edd := range _fag.Segments {
		if _edd.SegmentNumber == uint32(number) {
			return _edd, nil
		}
	}
	_daf := make([]uint32, len(_fag.Segments))
	for _bdfb, _ccd := range _fag.Segments {
		_daf[_bdfb] = _ccd.SegmentNumber
	}
	return nil, _gd.Errorf(_cgg, "\u0073e\u0067\u006d\u0065n\u0074\u0020\u0077i\u0074h \u006e\u0075\u006d\u0062\u0065\u0072\u003a \u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0067\u0065\u003a\u0020'%\u0064'\u002e\u0020\u004b\u006e\u006f\u0077n\u0020\u0073\u0065\u0067\u006de\u006e\u0074\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0073\u003a \u0025\u0076", number, _fag.PageNumber, _daf)
}

func (_agdd *Page) AddGenericRegion(bm *_ff.Bitmap, xloc, yloc, template int, tp _ca.Type, duplicateLineRemoval bool) error {
	const _cca = "P\u0061\u0067\u0065\u002eAd\u0064G\u0065\u006e\u0065\u0072\u0069c\u0052\u0065\u0067\u0069\u006f\u006e"
	_cbd := &_ca.GenericRegion{}
	if _feg := _cbd.InitEncode(bm, xloc, yloc, template, duplicateLineRemoval); _feg != nil {
		return _gd.Wrap(_feg, _cca, "")
	}
	_afg := &_ca.Header{Type: _ca.TImmediateGenericRegion, PageAssociation: _agdd.PageNumber, SegmentData: _cbd}
	_agdd.Segments = append(_agdd.Segments, _afg)
	return nil
}

func (_ecgce *Page) getPageInformationSegment() *_ca.Header {
	for _, _egg := range _ecgce.Segments {
		if _egg.Type == _ca.TPageInformation {
			return _egg
		}
	}
	_cb.Log.Debug("\u0050\u0061\u0067\u0065\u0020\u0069\u006e\u0066o\u0072\u006d\u0061ti\u006f\u006e\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0066o\u0072\u0020\u0070\u0061\u0067\u0065\u003a\u0020%\u0073\u002e", _ecgce)
	return nil
}

func (_ec *Document) parseFileHeader() error {
	const _bcfd = "\u0070a\u0072s\u0065\u0046\u0069\u006c\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_, _gbba := _ec.InputStream.Seek(8, _fe.SeekStart)
	if _gbba != nil {
		return _gd.Wrap(_gbba, _bcfd, "\u0069\u0064")
	}
	_, _gbba = _ec.InputStream.ReadBits(5)
	if _gbba != nil {
		return _gd.Wrap(_gbba, _bcfd, "\u0072\u0065\u0073\u0065\u0072\u0076\u0065\u0064\u0020\u0062\u0069\u0074\u0073")
	}
	_fbe, _gbba := _ec.InputStream.ReadBit()
	if _gbba != nil {
		return _gd.Wrap(_gbba, _bcfd, "\u0065x\u0074e\u006e\u0064\u0065\u0064\u0020t\u0065\u006dp\u006c\u0061\u0074\u0065\u0073")
	}
	if _fbe == 1 {
		_ec.GBUseExtTemplate = true
	}
	_fbe, _gbba = _ec.InputStream.ReadBit()
	if _gbba != nil {
		return _gd.Wrap(_gbba, _bcfd, "\u0075\u006e\u006b\u006eow\u006e\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	if _fbe != 1 {
		_ec.NumberOfPagesUnknown = false
	}
	_fbe, _gbba = _ec.InputStream.ReadBit()
	if _gbba != nil {
		return _gd.Wrap(_gbba, _bcfd, "\u006f\u0072\u0067\u0061\u006e\u0069\u007a\u0061\u0074\u0069\u006f\u006e \u0074\u0079\u0070\u0065")
	}
	_ec.OrganizationType = _ca.OrganizationType(_fbe)
	if !_ec.NumberOfPagesUnknown {
		_ec.NumberOfPages, _gbba = _ec.InputStream.ReadUint32()
		if _gbba != nil {
			return _gd.Wrap(_gbba, _bcfd, "\u006eu\u006db\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0061\u0067\u0065\u0073")
		}
		_ec._ggg = 13
	}
	return nil
}

func (_ffcdb *Page) Encode(w _a.BinaryWriter) (_ebb int, _eabb error) {
	const _acg = "P\u0061\u0067\u0065\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	var _eaa int
	for _, _eba := range _ffcdb.Segments {
		if _eaa, _eabb = _eba.Encode(w); _eabb != nil {
			return _ebb, _gd.Wrap(_eabb, _acg, "")
		}
		_ebb += _eaa
	}
	return _ebb, nil
}

func (_efaa *Document) encodeFileHeader(_de _a.BinaryWriter) (_fbd int, _fba error) {
	const _dad = "\u0065\u006ec\u006f\u0064\u0065F\u0069\u006c\u0065\u0048\u0065\u0061\u0064\u0065\u0072"
	_fbd, _fba = _de.Write(_gg)
	if _fba != nil {
		return _fbd, _gd.Wrap(_fba, _dad, "\u0069\u0064")
	}
	if _fba = _de.WriteByte(0x01); _fba != nil {
		return _fbd, _gd.Wrap(_fba, _dad, "\u0066\u006c\u0061g\u0073")
	}
	_fbd++
	_gfc := make([]byte, 4)
	_b.BigEndian.PutUint32(_gfc, _efaa.NumberOfPages)
	_acc, _fba := _de.Write(_gfc)
	if _fba != nil {
		return _acc, _gd.Wrap(_fba, _dad, "p\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_fbd += _acc
	return _fbd, nil
}

func (_aca *Document) completeClassifiedPages() (_ab error) {
	const _bg = "\u0063\u006f\u006dpl\u0065\u0074\u0065\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0050\u0061\u0067\u0065\u0073"
	if _aca.Classer == nil {
		return nil
	}
	_aca._ffa = make([]int, _aca.Classer.UndilatedTemplates.Size())
	for _be := 0; _be < _aca.Classer.ClassIDs.Size(); _be++ {
		_fcb, _ee := _aca.Classer.ClassIDs.Get(_be)
		if _ee != nil {
			return _gd.Wrapf(_ee, _bg, "\u0063\u006c\u0061\u0073s \u0077\u0069\u0074\u0068\u0020\u0069\u0064\u003a\u0020\u0027\u0025\u0064\u0027", _be)
		}
		_aca._ffa[_fcb]++
	}
	var _feb []int
	for _bf := 0; _bf < _aca.Classer.UndilatedTemplates.Size(); _bf++ {
		if _aca.NumberOfPages == 1 || _aca._ffa[_bf] > 1 {
			_feb = append(_feb, _bf)
		}
	}
	var (
		_cg *Page
		_df bool
	)
	for _bcf, _ef := range *_aca.Classer.ComponentPageNumbers {
		if _cg, _df = _aca.Pages[_ef]; !_df {
			return _gd.Errorf(_bg, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _bcf)
		}
		if _cg.EncodingMethod == GenericEM {
			_cb.Log.Error("\u0047\u0065\u006e\u0065\u0072\u0069c\u0020\u0070\u0061g\u0065\u0020\u0077i\u0074\u0068\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u003a \u0027\u0025\u0064\u0027\u0020ma\u0070\u0070\u0065\u0064\u0020\u0061\u0073\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0020\u0070\u0061\u0067\u0065", _bcf)
			continue
		}
		_aca._ac[_ef] = append(_aca._ac[_ef], _bcf)
		_bdf, _dfd := _aca.Classer.ClassIDs.Get(_bcf)
		if _dfd != nil {
			return _gd.Wrapf(_dfd, _bg, "\u006e\u006f\u0020\u0073uc\u0068\u0020\u0063\u006c\u0061\u0073\u0073\u0049\u0044\u003a\u0020\u0025\u0064", _bcf)
		}
		if _aca._ffa[_bdf] == 1 && _aca.NumberOfPages != 1 {
			_cfe := append(_aca._egf[_ef], _bdf)
			_aca._egf[_ef] = _cfe
		}
	}
	if _ab = _aca.Classer.ComputeLLCorners(); _ab != nil {
		return _gd.Wrap(_ab, _bg, "")
	}
	if _, _ab = _aca.addSymbolDictionary(0, _aca.Classer.UndilatedTemplates, _feb, _aca._ad, false); _ab != nil {
		return _gd.Wrap(_ab, _bg, "")
	}
	return nil
}

func (_ccb *Document) Encode() (_fdc []byte, _dga error) {
	const _bbg = "\u0044o\u0063u\u006d\u0065\u006e\u0074\u002e\u0045\u006e\u0063\u006f\u0064\u0065"
	var _aga, _eaf int
	if _ccb.FullHeaders {
		if _aga, _dga = _ccb.encodeFileHeader(_ccb._fc); _dga != nil {
			return nil, _gd.Wrap(_dga, _bbg, "")
		}
	}
	var (
		_fga bool
		_bfe *_ca.Header
		_gab *Page
	)
	if _dga = _ccb.completeClassifiedPages(); _dga != nil {
		return nil, _gd.Wrap(_dga, _bbg, "")
	}
	if _dga = _ccb.produceClassifiedPages(); _dga != nil {
		return nil, _gd.Wrap(_dga, _bbg, "")
	}
	if _ccb.GlobalSegments != nil {
		for _, _bfe = range _ccb.GlobalSegments.Segments {
			if _dga = _ccb.encodeSegment(_bfe, &_aga); _dga != nil {
				return nil, _gd.Wrap(_dga, _bbg, "")
			}
		}
	}
	for _bad := 1; _bad <= int(_ccb.NumberOfPages); _bad++ {
		if _gab, _fga = _ccb.Pages[_bad]; !_fga {
			return nil, _gd.Errorf(_bbg, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _bad)
		}
		for _, _bfe = range _gab.Segments {
			if _dga = _ccb.encodeSegment(_bfe, &_aga); _dga != nil {
				return nil, _gd.Wrap(_dga, _bbg, "")
			}
		}
	}
	if _ccb.FullHeaders {
		if _eaf, _dga = _ccb.encodeEOFHeader(_ccb._fc); _dga != nil {
			return nil, _gd.Wrap(_dga, _bbg, "")
		}
		_aga += _eaf
	}
	_fdc = _ccb._fc.Data()
	if len(_fdc) != _aga {
		_cb.Log.Debug("\u0042\u0079\u0074\u0065\u0073 \u0077\u0072\u0069\u0074\u0074\u0065\u006e \u0028\u006e\u0029\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u0064\u0061\u0074\u0061\u0020\u0065\u006e\u0063\u006fd\u0065\u0064\u003a\u0020\u0027\u0025d\u0027", _aga, len(_fdc))
	}
	return _fdc, nil
}

func (_gded *Page) getCombinationOperator(_ebag *_ca.PageInformationSegment, _dadc _ff.CombinationOperator) _ff.CombinationOperator {
	if _ebag.CombinationOperatorOverrideAllowed() {
		return _dadc
	}
	return _ebag.CombinationOperator()
}
func (_gdd *Page) GetResolutionX() (int, error) { return _gdd.getResolutionX() }
func (_gbdb *Page) GetHeight() (int, error)     { return _gbdb.getHeight() }

type Page struct {
	Segments           []*_ca.Header
	PageNumber         int
	Bitmap             *_ff.Bitmap
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

func (_ege *Document) encodeSegment(_cdb *_ca.Header, _bce *int) error {
	const _aa = "\u0065\u006e\u0063\u006f\u0064\u0065\u0053\u0065\u0067\u006d\u0065\u006e\u0074"
	_cdb.SegmentNumber = _ege.nextSegmentNumber()
	_dac, _dgc := _cdb.Encode(_ege._fc)
	if _dgc != nil {
		return _gd.Wrapf(_dgc, _aa, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0027\u0025\u0064\u0027", _cdb.SegmentNumber)
	}
	*_bce += _dac
	return nil
}
func (_ddd *Page) nextSegmentNumber() uint32 { return _ddd.Document.nextSegmentNumber() }
func (_ecgf *Page) fitsPage(_fef *_ca.PageInformationSegment, _baea *_ff.Bitmap) bool {
	return _ecgf.countRegions() == 1 && _fef.DefaultPixelValue == 0 && _fef.PageBMWidth == _baea.Width && _fef.PageBMHeight == _baea.Height
}

func (_gcf *Document) mapData() error {
	const _ccf = "\u006da\u0070\u0044\u0061\u0074\u0061"
	var (
		_cfc []*_ca.Header
		_gbd int64
		_eab _ca.Type
	)
	_eag, _dfb := _gcf.isFileHeaderPresent()
	if _dfb != nil {
		return _gd.Wrap(_dfb, _ccf, "")
	}
	if _eag {
		if _dfb = _gcf.parseFileHeader(); _dfb != nil {
			return _gd.Wrap(_dfb, _ccf, "")
		}
		_gbd += int64(_gcf._ggg)
		_gcf.FullHeaders = true
	}
	var (
		_eabg *Page
		_ffgb bool
	)
	for _eab != 51 && !_ffgb {
		_ffgbd, _fdf := _ca.NewHeader(_gcf, _gcf.InputStream, _gbd, _gcf.OrganizationType)
		if _fdf != nil {
			return _gd.Wrap(_fdf, _ccf, "")
		}
		_cb.Log.Trace("\u0044\u0065c\u006f\u0064\u0069\u006eg\u0020\u0073e\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0025\u0064\u002c\u0020\u0054\u0079\u0070e\u003a\u0020\u0025\u0073", _ffgbd.SegmentNumber, _ffgbd.Type)
		_eab = _ffgbd.Type
		if _eab != _ca.TEndOfFile {
			if _ffgbd.PageAssociation != 0 {
				_eabg = _gcf.Pages[_ffgbd.PageAssociation]
				if _eabg == nil {
					_eabg = _gad(_gcf, _ffgbd.PageAssociation)
					_gcf.Pages[_ffgbd.PageAssociation] = _eabg
					if _gcf.NumberOfPagesUnknown {
						_gcf.NumberOfPages++
					}
				}
				_eabg.Segments = append(_eabg.Segments, _ffgbd)
			} else {
				_gcf.GlobalSegments.AddSegment(_ffgbd)
			}
		}
		_cfc = append(_cfc, _ffgbd)
		_gbd = _gcf.InputStream.AbsolutePosition()
		if _gcf.OrganizationType == _ca.OSequential {
			_gbd += int64(_ffgbd.SegmentDataLength)
		}
		_ffgb, _fdf = _gcf.reachedEOF(_gbd)
		if _fdf != nil {
			_cb.Log.Debug("\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0072\u0065\u0061\u0063h\u0065\u0064\u0020\u0045\u004f\u0046\u0020\u0077\u0069\u0074\u0068\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076", _fdf)
			return _gd.Wrap(_fdf, _ccf, "")
		}
	}
	_gcf.determineRandomDataOffsets(_cfc, uint64(_gbd))
	return nil
}

func (_cee *Page) String() string {
	return _g.Sprintf("\u0050\u0061\u0067\u0065\u0020\u0023\u0025\u0064", _cee.PageNumber)
}

func (_dfbe *Globals) GetSymbolDictionary() (*_ca.Header, error) {
	const _gda = "G\u006c\u006f\u0062\u0061\u006c\u0073.\u0047\u0065\u0074\u0053\u0079\u006d\u0062\u006f\u006cD\u0069\u0063\u0074i\u006fn\u0061\u0072\u0079"
	if _dfbe == nil {
		return nil, _gd.Error(_gda, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_dfbe.Segments) == 0 {
		return nil, _gd.Error(_gda, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	for _, _bef := range _dfbe.Segments {
		if _bef.Type == _ca.TSymbolDictionary {
			return _bef, nil
		}
	}
	return nil, _gd.Error(_gda, "\u0067\u006c\u006fba\u006c\u0020\u0073\u0079\u006d\u0062\u006f\u006c\u0020d\u0069c\u0074i\u006fn\u0061\u0072\u0079\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}

func (_dde *Page) createStripedPage(_ceee *_ca.PageInformationSegment) error {
	const _fad = "\u0063\u0072\u0065\u0061\u0074\u0065\u0053\u0074\u0072\u0069\u0070\u0065d\u0050\u0061\u0067\u0065"
	_ffgc, _ead := _dde.collectPageStripes()
	if _ead != nil {
		return _gd.Wrap(_ead, _fad, "")
	}
	var _dbb int
	for _, _fggg := range _ffgc {
		if _dgba, _gcc := _fggg.(*_ca.EndOfStripe); _gcc {
			_dbb = _dgba.LineNumber() + 1
		} else {
			_cdc := _fggg.(_ca.Regioner)
			_abce := _cdc.GetRegionInfo()
			_cgc := _dde.getCombinationOperator(_ceee, _abce.CombinaionOperator)
			_cbgb, _dfe := _cdc.GetRegionBitmap()
			if _dfe != nil {
				return _gd.Wrap(_dfe, _fad, "")
			}
			_dfe = _ff.Blit(_cbgb, _dde.Bitmap, int(_abce.XLocation), _dbb, _cgc)
			if _dfe != nil {
				return _gd.Wrap(_dfe, _fad, "")
			}
		}
	}
	return nil
}

func (_acd *Document) isFileHeaderPresent() (bool, error) {
	_acd.InputStream.Mark()
	for _, _bfb := range _gg {
		_eed, _bdd := _acd.InputStream.ReadByte()
		if _bdd != nil {
			return false, _bdd
		}
		if _bfb != _eed {
			_acd.InputStream.Reset()
			return false, nil
		}
	}
	_acd.InputStream.Reset()
	return true, nil
}

func (_ded *Globals) GetSegment(segmentNumber int) (*_ca.Header, error) {
	const _eeg = "\u0047l\u006fb\u0061\u006c\u0073\u002e\u0047e\u0074\u0053e\u0067\u006d\u0065\u006e\u0074"
	if _ded == nil {
		return nil, _gd.Error(_eeg, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_ded.Segments) == 0 {
		return nil, _gd.Error(_eeg, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	var _ffac *_ca.Header
	for _, _ffac = range _ded.Segments {
		if _ffac.SegmentNumber == uint32(segmentNumber) {
			break
		}
	}
	if _ffac == nil {
		return nil, _gd.Error(_eeg, "\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	return _ffac, nil
}
func (_baaa *Page) GetWidth() (int, error) { return _baaa.getWidth() }
func (_dcd *Globals) GetSegmentByIndex(index int) (*_ca.Header, error) {
	const _cbf = "\u0047l\u006f\u0062\u0061\u006cs\u002e\u0047\u0065\u0074\u0053e\u0067m\u0065n\u0074\u0042\u0079\u0049\u006e\u0064\u0065x"
	if _dcd == nil {
		return nil, _gd.Error(_cbf, "\u0067\u006c\u006f\u0062al\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(_dcd.Segments) == 0 {
		return nil, _gd.Error(_cbf, "\u0067\u006c\u006f\u0062\u0061\u006c\u0073\u0020\u0061\u0072\u0065\u0020e\u006d\u0070\u0074\u0079")
	}
	if index > len(_dcd.Segments)-1 {
		return nil, _gd.Error(_cbf, "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	return _dcd.Segments[index], nil
}

func (_ffg *Document) produceClassifiedPage(_bedc *Page, _bff *_ca.Header) (_eda error) {
	const _gb = "p\u0072\u006f\u0064\u0075ce\u0043l\u0061\u0073\u0073\u0069\u0066i\u0065\u0064\u0050\u0061\u0067\u0065"
	var _ffc map[int]int
	_gbb := _ffg._gf
	_da := []*_ca.Header{_bff}
	if len(_ffg._egf[_bedc.PageNumber]) > 0 {
		_ffc = map[int]int{}
		_cce, _bab := _ffg.addSymbolDictionary(_bedc.PageNumber, _ffg.Classer.UndilatedTemplates, _ffg._egf[_bedc.PageNumber], _ffc, false)
		if _bab != nil {
			return _gd.Wrap(_bab, _gb, "")
		}
		_da = append(_da, _cce)
		_gbb += len(_ffg._egf[_bedc.PageNumber])
	}
	_fae := _ffg._ac[_bedc.PageNumber]
	_cb.Log.Debug("P\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020c\u006f\u006d\u0070\u0073: \u0025\u0076", _bedc.PageNumber, _fae)
	_bedc.addTextRegionSegment(_da, _ffg._ad, _ffc, _ffg._ac[_bedc.PageNumber], _ffg.Classer.PtaLL, _ffg.Classer.UndilatedTemplates, _ffg.Classer.ClassIDs, nil, _ea(_gbb), len(_ffg._ac[_bedc.PageNumber]))
	return nil
}

func (_gcg *Document) addSymbolDictionary(_gcgg int, _ge *_ff.Bitmaps, _bea []int, _cbb map[int]int, _eb bool) (*_ca.Header, error) {
	const _ga = "\u0061\u0064\u0064\u0053ym\u0062\u006f\u006c\u0044\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079"
	_abc := &_ca.SymbolDictionary{}
	if _cd := _abc.InitEncode(_ge, _bea, _cbb, _eb); _cd != nil {
		return nil, _cd
	}
	_fg := &_ca.Header{Type: _ca.TSymbolDictionary, PageAssociation: _gcgg, SegmentData: _abc}
	if _gcgg == 0 {
		if _gcg.GlobalSegments == nil {
			_gcg.GlobalSegments = &Globals{}
		}
		_gcg.GlobalSegments.AddSegment(_fg)
		return _fg, nil
	}
	_dc, _abf := _gcg.Pages[_gcgg]
	if !_abf {
		return nil, _gd.Errorf(_ga, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _gcgg)
	}
	var (
		_ace int
		_baa *_ca.Header
	)
	for _ace, _baa = range _dc.Segments {
		if _baa.Type == _ca.TPageInformation {
			break
		}
	}
	_ace++
	_dc.Segments = append(_dc.Segments, nil)
	copy(_dc.Segments[_ace+1:], _dc.Segments[_ace:])
	_dc.Segments[_ace] = _fg
	return _fg, nil
}

func (_fac *Document) encodeEOFHeader(_bde _a.BinaryWriter) (_ffcd int, _bfdd error) {
	_ged := &_ca.Header{SegmentNumber: _fac.nextSegmentNumber(), Type: _ca.TEndOfFile}
	if _ffcd, _bfdd = _ged.Encode(_bde); _bfdd != nil {
		return 0, _gd.Wrap(_bfdd, "\u0065n\u0063o\u0064\u0065\u0045\u004f\u0046\u0048\u0065\u0061\u0064\u0065\u0072", "")
	}
	return _ffcd, nil
}

type (
	EncodingMethod int
	Globals        struct{ Segments []*_ca.Header }
)

var _gg = []byte{0x97, 0x4A, 0x42, 0x32, 0x0D, 0x0A, 0x1A, 0x0A}

func (_efa *Document) completeSymbols() (_af error) {
	const _ebg = "\u0063o\u006dp\u006c\u0065\u0074\u0065\u0053\u0079\u006d\u0062\u006f\u006c\u0073"
	if _efa.Classer == nil {
		return nil
	}
	if _efa.Classer.UndilatedTemplates == nil {
		return _gd.Error(_ebg, "\u006e\u006f t\u0065\u006d\u0070l\u0061\u0074\u0065\u0073 de\u0066in\u0065\u0064\u0020\u0066\u006f\u0072\u0020th\u0065\u0020\u0063\u006c\u0061\u0073\u0073e\u0072")
	}
	_gge := len(_efa.Pages) == 1
	_adb := make([]int, _efa.Classer.UndilatedTemplates.Size())
	var _ce int
	for _abg := 0; _abg < _efa.Classer.ClassIDs.Size(); _abg++ {
		_ce, _af = _efa.Classer.ClassIDs.Get(_abg)
		if _af != nil {
			return _gd.Wrap(_af, _ebg, "\u0063\u006c\u0061\u0073\u0073\u0020\u0049\u0044\u0027\u0073")
		}
		_adb[_ce]++
	}
	var _bb []int
	for _bgd := 0; _bgd < _efa.Classer.UndilatedTemplates.Size(); _bgd++ {
		if _adb[_bgd] == 0 {
			return _gd.Error(_ebg, "\u006eo\u0020\u0073y\u006d\u0062\u006f\u006cs\u0020\u0069\u006es\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u0066\u006fun\u0064\u0020\u0066o\u0072\u0020g\u0069\u0076\u0065\u006e\u0020\u0063l\u0061\u0073s\u003f\u0020")
		}
		if _adb[_bgd] > 1 || _gge {
			_bb = append(_bb, _bgd)
		}
	}
	_efa._gf = len(_bb)
	var _fca, _bda int
	for _bcc := 0; _bcc < _efa.Classer.ComponentPageNumbers.Size(); _bcc++ {
		_fca, _af = _efa.Classer.ComponentPageNumbers.Get(_bcc)
		if _af != nil {
			return _gd.Wrapf(_af, _ebg, "p\u0061\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027 \u006e\u006f\u0074\u0020\u0066\u006f\u0075nd\u0020\u0069\u006e\u0020t\u0068\u0065\u0020\u0063\u006c\u0061\u0073\u0073\u0065r \u0070\u0061g\u0065\u006e\u0075\u006d\u0062\u0065\u0072\u0073", _bcc)
		}
		_bda, _af = _efa.Classer.ClassIDs.Get(_bcc)
		if _af != nil {
			return _gd.Wrapf(_af, _ebg, "\u0063\u0061\u006e\u0027\u0074\u0020\u0067e\u0074\u0020\u0073y\u006d\u0062\u006f\u006c \u0066\u006f\u0072\u0020\u0070\u0061\u0067\u0065\u0020\u0027\u0025\u0064\u0027\u0020\u0066\u0072\u006f\u006d\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072", _fca)
		}
		if _adb[_bda] == 1 && !_gge {
			_efa._egf[_fca] = append(_efa._egf[_fca], _bda)
		}
	}
	if _af = _efa.Classer.ComputeLLCorners(); _af != nil {
		return _gd.Wrap(_af, _ebg, "")
	}
	return nil
}

func (_bba *Document) nextPageNumber() uint32 {
	_bba.NumberOfPages++
	return _bba.NumberOfPages
}
func (_agec *Page) GetResolutionY() (int, error) { return _agec.getResolutionY() }
func (_dge *Document) nextSegmentNumber() uint32 {
	_ebd := _dge.CurrentSegmentNumber
	_dge.CurrentSegmentNumber++
	return _ebd
}

func (_daa *Page) collectPageStripes() (_gabd []_ca.Segmenter, _cgfg error) {
	const _dbdg = "\u0063o\u006cl\u0065\u0063\u0074\u0050\u0061g\u0065\u0053t\u0072\u0069\u0070\u0065\u0073"
	var _ecd _ca.Segmenter
	for _, _ebab := range _daa.Segments {
		switch _ebab.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_ecd, _cgfg = _ebab.GetSegmentData()
			if _cgfg != nil {
				return nil, _gd.Wrap(_cgfg, _dbdg, "")
			}
			_gabd = append(_gabd, _ecd)
		case 50:
			_ecd, _cgfg = _ebab.GetSegmentData()
			if _cgfg != nil {
				return nil, _cgfg
			}
			_bbgb, _gfcc := _ecd.(*_ca.EndOfStripe)
			if !_gfcc {
				return nil, _gd.Errorf(_dbdg, "\u0045\u006e\u0064\u004f\u0066\u0053\u0074\u0072\u0069\u0070\u0065\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u006f\u0066\u0020\u0076\u0061l\u0069\u0064\u0020\u0074\u0079p\u0065\u003a \u0027\u0025\u0054\u0027", _ecd)
			}
			_gabd = append(_gabd, _bbgb)
			_daa.FinalHeight = _bbgb.LineNumber()
		}
	}
	return _gabd, nil
}

func (_cbde *Page) AddPageInformationSegment() {
	_fggc := &_ca.PageInformationSegment{PageBMWidth: _cbde.FinalWidth, PageBMHeight: _cbde.FinalHeight, ResolutionX: _cbde.ResolutionX, ResolutionY: _cbde.ResolutionY, IsLossless: _cbde.IsLossless}
	if _cbde.BlackIsOne {
		_fggc.DefaultPixelValue = uint8(0x1)
	}
	_efag := &_ca.Header{PageAssociation: _cbde.PageNumber, SegmentDataLength: uint64(_fggc.Size()), SegmentData: _fggc, Type: _ca.TPageInformation}
	_cbde.Segments = append(_cbde.Segments, _efag)
}

func (_beef *Page) getResolutionX() (int, error) {
	const _geg = "\u0067\u0065\u0074\u0052\u0065\u0073\u006f\u006c\u0075t\u0069\u006f\u006e\u0058"
	if _beef.ResolutionX != 0 {
		return _beef.ResolutionX, nil
	}
	_fgbf := _beef.getPageInformationSegment()
	if _fgbf == nil {
		return 0, _gd.Error(_geg, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_abd, _fge := _fgbf.GetSegmentData()
	if _fge != nil {
		return 0, _gd.Wrap(_fge, _geg, "")
	}
	_bbef, _gfb := _abd.(*_ca.PageInformationSegment)
	if !_gfb {
		return 0, _gd.Errorf(_geg, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _abd)
	}
	_beef.ResolutionX = _bbef.ResolutionX
	return _beef.ResolutionX, nil
}

func (_ag *Document) produceClassifiedPages() (_cfg error) {
	const _bed = "\u0070\u0072\u006f\u0064uc\u0065\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064\u0050\u0061\u0067e\u0073"
	if _ag.Classer == nil {
		return nil
	}
	var (
		_bae  *Page
		_fa   bool
		_baeg *_ca.Header
	)
	for _gggf := 1; _gggf <= int(_ag.NumberOfPages); _gggf++ {
		if _bae, _fa = _ag.Pages[_gggf]; !_fa {
			return _gd.Errorf(_bed, "p\u0061g\u0065\u003a\u0020\u0027\u0025\u0064\u0027\u0020n\u006f\u0074\u0020\u0066ou\u006e\u0064", _gggf)
		}
		if _bae.EncodingMethod == GenericEM {
			continue
		}
		if _baeg == nil {
			if _baeg, _cfg = _ag.GlobalSegments.GetSymbolDictionary(); _cfg != nil {
				return _gd.Wrap(_cfg, _bed, "")
			}
		}
		if _cfg = _ag.produceClassifiedPage(_bae, _baeg); _cfg != nil {
			return _gd.Wrapf(_cfg, _bed, "\u0070\u0061\u0067\u0065\u003a\u0020\u0027\u0025\u0064\u0027", _gggf)
		}
	}
	return nil
}

func (_ae *Page) addTextRegionSegment(_dda []*_ca.Header, _babf, _gaaf map[int]int, _begg []int, _bfgg *_ff.Points, _bbgd *_ff.Bitmaps, _gfg *_eg.IntSlice, _aec *_ff.Boxes, _aeg, _dea int) {
	_ecg := &_ca.TextRegion{NumberOfSymbols: uint32(_dea)}
	_ecg.InitEncode(_babf, _gaaf, _begg, _bfgg, _bbgd, _gfg, _aec, _ae.FinalWidth, _ae.FinalHeight, _aeg)
	_abcc := &_ca.Header{RTSegments: _dda, SegmentData: _ecg, PageAssociation: _ae.PageNumber, Type: _ca.TImmediateTextRegion}
	_agc := _ca.TPageInformation
	if _gaaf != nil {
		_agc = _ca.TSymbolDictionary
	}
	var _faf int
	for ; _faf < len(_ae.Segments); _faf++ {
		if _ae.Segments[_faf].Type == _agc {
			_faf++
			break
		}
	}
	_ae.Segments = append(_ae.Segments, nil)
	copy(_ae.Segments[_faf+1:], _ae.Segments[_faf:])
	_ae.Segments[_faf] = _abcc
}

func _edc(_dd *_a.Reader, _bcg *Globals) (*Document, error) {
	_def := &Document{Pages: make(map[int]*Page), InputStream: _dd, OrganizationType: _ca.OSequential, NumberOfPagesUnknown: true, GlobalSegments: _bcg, _ggg: 9}
	if _def.GlobalSegments == nil {
		_def.GlobalSegments = &Globals{}
	}
	if _edad := _def.mapData(); _edad != nil {
		return nil, _edad
	}
	return _def, nil
}

func (_dgac *Page) getWidth() (int, error) {
	const _edb = "\u0067\u0065\u0074\u0057\u0069\u0064\u0074\u0068"
	if _dgac.FinalWidth != 0 {
		return _dgac.FinalWidth, nil
	}
	_bfc := _dgac.getPageInformationSegment()
	if _bfc == nil {
		return 0, _gd.Error(_edb, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_aed, _bfgb := _bfc.GetSegmentData()
	if _bfgb != nil {
		return 0, _gd.Wrap(_bfgb, _edb, "")
	}
	_gddg, _befc := _aed.(*_ca.PageInformationSegment)
	if !_befc {
		return 0, _gd.Errorf(_edb, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _aed)
	}
	_dgac.FinalWidth = _gddg.PageBMWidth
	return _dgac.FinalWidth, nil
}

func (_bbba *Page) clearSegmentData() {
	for _ceed := range _bbba.Segments {
		_bbba.Segments[_ceed].CleanSegmentData()
	}
}

func (_bdaa *Document) reachedEOF(_bbe int64) (bool, error) {
	const _bgcc = "\u0072\u0065\u0061\u0063\u0068\u0065\u0064\u0045\u004f\u0046"
	_, _ggf := _bdaa.InputStream.Seek(_bbe, _fe.SeekStart)
	if _ggf != nil {
		_cb.Log.Debug("\u0072\u0065\u0061c\u0068\u0065\u0064\u0045\u004f\u0046\u0020\u002d\u0020\u0064\u002e\u0049\u006e\u0070\u0075\u0074\u0053\u0074\u0072\u0065\u0061\u006d\u002e\u0053\u0065\u0065\u006b\u0020\u0066a\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ggf)
		return false, _gd.Wrap(_ggf, _bgcc, "\u0069n\u0070\u0075\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020s\u0065\u0065\u006b\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	_, _ggf = _bdaa.InputStream.ReadBits(32)
	if _ggf == _fe.EOF {
		return true, nil
	} else if _ggf != nil {
		return false, _gd.Wrap(_ggf, _bgcc, "")
	}
	return false, nil
}

func (_bdeb *Page) createPage(_cag *_ca.PageInformationSegment) error {
	var _cgf error
	if !_cag.IsStripe || _cag.PageBMHeight != -1 {
		_cgf = _bdeb.createNormalPage(_cag)
	} else {
		_cgf = _bdeb.createStripedPage(_cag)
	}
	return _cgf
}

func (_bgae *Page) createNormalPage(_bac *_ca.PageInformationSegment) error {
	const _ecgc = "\u0063\u0072e\u0061\u0074\u0065N\u006f\u0072\u006d\u0061\u006c\u0050\u0061\u0067\u0065"
	_bgae.Bitmap = _ff.New(_bac.PageBMWidth, _bac.PageBMHeight)
	if _bac.DefaultPixelValue != 0 {
		_bgae.Bitmap.SetDefaultPixel()
	}
	for _, _gdde := range _bgae.Segments {
		switch _gdde.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_cb.Log.Trace("\u0047\u0065\u0074\u0074in\u0067\u0020\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0064", _gdde.SegmentNumber)
			_deb, _fgb := _gdde.GetSegmentData()
			if _fgb != nil {
				return _fgb
			}
			_ggeb, _gea := _deb.(_ca.Regioner)
			if !_gea {
				_cb.Log.Debug("\u0053\u0065g\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0054\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0052\u0065\u0067\u0069on\u0065\u0072", _deb)
				return _gd.Errorf(_ecgc, "i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062i\u0067\u0032\u0020\u0073\u0065\u0067\u006den\u0074\u0020\u0074\u0079p\u0065\u0020\u002d\u0020\u006e\u006f\u0074\u0020\u0061 R\u0065\u0067i\u006f\u006e\u0065\u0072\u003a\u0020\u0025\u0054", _deb)
			}
			_faa, _fgb := _ggeb.GetRegionBitmap()
			if _fgb != nil {
				return _gd.Wrap(_fgb, _ecgc, "")
			}
			if _bgae.fitsPage(_bac, _faa) {
				_bgae.Bitmap = _faa
			} else {
				_cae := _ggeb.GetRegionInfo()
				_cbg := _bgae.getCombinationOperator(_bac, _cae.CombinaionOperator)
				_fgb = _ff.Blit(_faa, _bgae.Bitmap, int(_cae.XLocation), int(_cae.YLocation), _cbg)
				if _fgb != nil {
					return _gd.Wrap(_fgb, _ecgc, "")
				}
			}
		}
	}
	return nil
}

func (_beff *Page) AddEndOfPageSegment() {
	_ffca := &_ca.Header{Type: _ca.TEndOfPage, PageAssociation: _beff.PageNumber}
	_beff.Segments = append(_beff.Segments, _ffca)
}

func (_cga *Page) getHeight() (int, error) {
	const _cdbf = "\u0067e\u0074\u0048\u0065\u0069\u0067\u0068t"
	if _cga.FinalHeight != 0 {
		return _cga.FinalHeight, nil
	}
	_edac := _cga.getPageInformationSegment()
	if _edac == nil {
		return 0, _gd.Error(_cdbf, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_dca, _bfba := _edac.GetSegmentData()
	if _bfba != nil {
		return 0, _gd.Wrap(_bfba, _cdbf, "")
	}
	_ffaf, _dcgc := _dca.(*_ca.PageInformationSegment)
	if !_dcgc {
		return 0, _gd.Errorf(_cdbf, "\u0070\u0061\u0067\u0065\u0020\u0069n\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073 \u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070e\u003a \u0027\u0025\u0054\u0027", _dca)
	}
	if _ffaf.PageBMHeight == _f.MaxInt32 {
		_, _bfba = _cga.GetBitmap()
		if _bfba != nil {
			return 0, _gd.Wrap(_bfba, _cdbf, "")
		}
	} else {
		_cga.FinalHeight = _ffaf.PageBMHeight
	}
	return _cga.FinalHeight, nil
}

func (_cf *Document) AddGenericPage(bm *_ff.Bitmap, duplicateLineRemoval bool) (_bc error) {
	const _dg = "\u0044\u006f\u0063um\u0065\u006e\u0074\u002e\u0041\u0064\u0064\u0047\u0065\u006e\u0065\u0072\u0069\u0063\u0050\u0061\u0067\u0065"
	if !_cf.FullHeaders && _cf.NumberOfPages != 0 {
		return _gd.Error(_dg, "\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0061\u006c\u0072\u0065a\u0064\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0070\u0061\u0067\u0065\u002e\u0020\u0046\u0069\u006c\u0065\u004d\u006f\u0064\u0065\u0020\u0064\u0069\u0073\u0061\u006c\u006c\u006f\u0077\u0073\u0020\u0061\u0064\u0064i\u006e\u0067\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e \u006f\u006e\u0065\u0020\u0070\u0061g\u0065")
	}
	_ba := &Page{Segments: []*_ca.Header{}, Bitmap: bm, Document: _cf, FinalHeight: bm.Height, FinalWidth: bm.Width, IsLossless: true, BlackIsOne: bm.Color == _ff.Chocolate}
	_ba.PageNumber = int(_cf.nextPageNumber())
	_cf.Pages[_ba.PageNumber] = _ba
	bm.InverseData()
	_ba.AddPageInformationSegment()
	if _bc = _ba.AddGenericRegion(bm, 0, 0, 0, _ca.TImmediateGenericRegion, duplicateLineRemoval); _bc != nil {
		return _gd.Wrap(_bc, _dg, "")
	}
	if _cf.FullHeaders {
		_ba.AddEndOfPageSegment()
	}
	return nil
}

func (_agag *Page) composePageBitmap() error {
	const _cgd = "\u0063\u006f\u006d\u0070\u006f\u0073\u0065\u0050\u0061\u0067\u0065\u0042i\u0074\u006d\u0061\u0070"
	if _agag.PageNumber == 0 {
		return nil
	}
	_bgg := _agag.getPageInformationSegment()
	if _bgg == nil {
		return _gd.Error(_cgd, "\u0070\u0061\u0067e \u0069\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069o\u006e \u0073e\u0067m\u0065\u006e\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_ddc, _bca := _bgg.GetSegmentData()
	if _bca != nil {
		return _bca
	}
	_gde, _abgd := _ddc.(*_ca.PageInformationSegment)
	if !_abgd {
		return _gd.Error(_cgd, "\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006da\u0074\u0069\u006f\u006e \u0073\u0065\u0067\u006d\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0074\u0079\u0070\u0065")
	}
	if _bca = _agag.createPage(_gde); _bca != nil {
		return _gd.Wrap(_bca, _cgd, "")
	}
	_agag.clearSegmentData()
	return nil
}

func (_bee *Document) GetPage(pageNumber int) (_ca.Pager, error) {
	const _dgce = "\u0044\u006fc\u0075\u006d\u0065n\u0074\u002e\u0047\u0065\u0074\u0050\u0061\u0067\u0065"
	if pageNumber < 0 {
		_cb.Log.Debug("\u004a\u0042\u0049\u00472\u0020\u0050\u0061\u0067\u0065\u0020\u002d\u0020\u0047e\u0074\u0050\u0061\u0067\u0065\u003a\u0020\u0025\u0064\u002e\u0020\u0050\u0061\u0067\u0065\u0020\u0063\u0061n\u006e\u006f\u0074\u0020\u0062e\u0020\u006c\u006f\u0077\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030\u002e\u0020\u0025\u0073", pageNumber, _c.Stack())
		return nil, _gd.Errorf(_dgce, "\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032\u0020d\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u002d\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064 \u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u003a\u0020\u0025\u0064", pageNumber)
	}
	if pageNumber > len(_bee.Pages) {
		_cb.Log.Debug("\u0050\u0061\u0067\u0065 n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0064\u002e\u0020%\u0073", pageNumber, _c.Stack())
		return nil, _gd.Error(_dgce, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u002d\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	_afd, _bga := _bee.Pages[pageNumber]
	if !_bga {
		_cb.Log.Debug("\u0050\u0061\u0067\u0065 n\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u003a\u0020\u0025\u0064\u002e\u0020%\u0073", pageNumber, _c.Stack())
		return nil, _gd.Errorf(_dgce, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006a\u0062\u0069\u0067\u0032 \u0064\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u002d\u0020\u0070\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	return _afd, nil
}

func _ea(_agd int) int {
	_bfg := 0
	_fcf := (_agd & (_agd - 1)) == 0
	_agd >>= 1
	for ; _agd != 0; _agd >>= 1 {
		_bfg++
	}
	if _fcf {
		return _bfg
	}
	return _bfg + 1
}

func DecodeDocument(input *_a.Reader, globals *Globals) (*Document, error) {
	return _edc(input, globals)
}

func (_fb *Document) AddClassifiedPage(bm *_ff.Bitmap, method _d.Method) (_ed error) {
	const _bd = "\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0041\u0064d\u0043\u006c\u0061\u0073\u0073\u0069\u0066\u0069\u0065\u0064P\u0061\u0067\u0065"
	if !_fb.FullHeaders && _fb.NumberOfPages != 0 {
		return _gd.Error(_bd, "\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0061\u006c\u0072\u0065a\u0064\u0079\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0070\u0061\u0067\u0065\u002e\u0020\u0046\u0069\u006c\u0065\u004d\u006f\u0064\u0065\u0020\u0064\u0069\u0073\u0061\u006c\u006c\u006f\u0077\u0073\u0020\u0061\u0064\u0064i\u006e\u0067\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e \u006f\u006e\u0065\u0020\u0070\u0061g\u0065")
	}
	if _fb.Classer == nil {
		if _fb.Classer, _ed = _d.Init(_d.DefaultSettings()); _ed != nil {
			return _gd.Wrap(_ed, _bd, "")
		}
	}
	_fd := int(_fb.nextPageNumber())
	_cc := &Page{Segments: []*_ca.Header{}, Bitmap: bm, Document: _fb, FinalHeight: bm.Height, FinalWidth: bm.Width, PageNumber: _fd}
	_fb.Pages[_fd] = _cc
	switch method {
	case _d.RankHaus:
		_cc.EncodingMethod = RankHausEM
	case _d.Correlation:
		_cc.EncodingMethod = CorrelationEM
	}
	_cc.AddPageInformationSegment()
	if _ed = _fb.Classer.AddPage(bm, _fd, method); _ed != nil {
		return _gd.Wrap(_ed, _bd, "")
	}
	if _fb.FullHeaders {
		_cc.AddEndOfPageSegment()
	}
	return nil
}

const (
	GenericEM EncodingMethod = iota
	CorrelationEM
	RankHausEM
)

func (_ccfd *Page) GetBitmap() (_aaf *_ff.Bitmap, _fgd error) {
	_cb.Log.Trace(_g.Sprintf("\u005b\u0050\u0041G\u0045\u005d\u005b\u0023%\u0064\u005d\u0020\u0047\u0065\u0074\u0042i\u0074\u006d\u0061\u0070\u0020\u0062\u0065\u0067\u0069\u006e\u0073\u002e\u002e\u002e", _ccfd.PageNumber))
	defer func() {
		if _fgd != nil {
			_cb.Log.Trace(_g.Sprintf("\u005b\u0050\u0041\u0047\u0045\u005d\u005b\u0023\u0025\u0064\u005d\u0020\u0047\u0065\u0074B\u0069t\u006d\u0061\u0070\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076", _ccfd.PageNumber, _fgd))
		} else {
			_cb.Log.Trace(_g.Sprintf("\u005b\u0050\u0041\u0047\u0045\u005d\u005b\u0023\u0025\u0064]\u0020\u0047\u0065\u0074\u0042\u0069\u0074m\u0061\u0070\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064", _ccfd.PageNumber))
		}
	}()
	if _ccfd.Bitmap != nil {
		return _ccfd.Bitmap, nil
	}
	_fgd = _ccfd.composePageBitmap()
	if _fgd != nil {
		return nil, _fgd
	}
	return _ccfd.Bitmap, nil
}

func (_abb *Document) GetNumberOfPages() (uint32, error) {
	if _abb.NumberOfPagesUnknown || _abb.NumberOfPages == 0 {
		if len(_abb.Pages) == 0 {
			if _cda := _abb.mapData(); _cda != nil {
				return 0, _gd.Wrap(_cda, "\u0044o\u0063\u0075\u006d\u0065n\u0074\u002e\u0047\u0065\u0074N\u0075m\u0062e\u0072\u004f\u0066\u0050\u0061\u0067\u0065s", "")
			}
		}
		return uint32(len(_abb.Pages)), nil
	}
	return _abb.NumberOfPages, nil
}

func (_gef *Page) lastSegmentNumber() (_gfe uint32, _ceb error) {
	const _ebc = "\u006c\u0061\u0073\u0074\u0053\u0065\u0067\u006d\u0065\u006e\u0074\u004eu\u006d\u0062\u0065\u0072"
	if len(_gef.Segments) == 0 {
		return _gfe, _gd.Errorf(_ebc, "\u006e\u006f\u0020se\u0067\u006d\u0065\u006e\u0074\u0073\u0020\u0066\u006fu\u006ed\u0020i\u006e \u0074\u0068\u0065\u0020\u0070\u0061\u0067\u0065\u0020\u0027\u0025\u0064\u0027", _gef.PageNumber)
	}
	return _gef.Segments[len(_gef.Segments)-1].SegmentNumber, nil
}
func (_age *Globals) AddSegment(segment *_ca.Header) { _age.Segments = append(_age.Segments, segment) }
func (_bfgd *Page) getResolutionY() (int, error) {
	const _deg = "\u0067\u0065\u0074\u0052\u0065\u0073\u006f\u006c\u0075t\u0069\u006f\u006e\u0059"
	if _bfgd.ResolutionY != 0 {
		return _bfgd.ResolutionY, nil
	}
	_bfa := _bfgd.getPageInformationSegment()
	if _bfa == nil {
		return 0, _gd.Error(_deg, "n\u0069l\u0020\u0070\u0061\u0067\u0065\u0020\u0069\u006ef\u006f\u0072\u006d\u0061ti\u006f\u006e")
	}
	_ffb, _eaga := _bfa.GetSegmentData()
	if _eaga != nil {
		return 0, _gd.Wrap(_eaga, _deg, "")
	}
	_ecdb, _ecf := _ffb.(*_ca.PageInformationSegment)
	if !_ecf {
		return 0, _gd.Errorf(_deg, "\u0070\u0061\u0067\u0065\u0020\u0069\u006e\u0066o\u0072\u006d\u0061ti\u006f\u006e\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0020\u0069\u0073\u0020\u006f\u0066\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0074\u0079\u0070\u0065\u003a\u0027%\u0054\u0027", _ffb)
	}
	_bfgd.ResolutionY = _ecdb.ResolutionY
	return _bfgd.ResolutionY, nil
}

func (_afe *Document) GetGlobalSegment(i int) (*_ca.Header, error) {
	_bgc, _dcg := _afe.GlobalSegments.GetSegment(i)
	if _dcg != nil {
		return nil, _gd.Wrap(_dcg, "\u0047\u0065t\u0047\u006c\u006fb\u0061\u006c\u0053\u0065\u0067\u006d\u0065\u006e\u0074", "")
	}
	return _bgc, nil
}

func (_agb *Page) countRegions() int {
	var _eafe int
	for _, _fbb := range _agb.Segments {
		switch _fbb.Type {
		case 6, 7, 22, 23, 38, 39, 42, 43:
			_eafe++
		}
	}
	return _eafe
}

type Document struct {
	Pages                map[int]*Page
	NumberOfPagesUnknown bool
	NumberOfPages        uint32
	GBUseExtTemplate     bool
	InputStream          *_a.Reader
	GlobalSegments       *Globals
	OrganizationType     _ca.OrganizationType
	Classer              *_d.Classer
	XRes, YRes           int
	FullHeaders          bool
	CurrentSegmentNumber uint32
	AverageTemplates     *_ff.Bitmaps
	BaseIndexes          []int
	Refinement           bool
	RefineLevel          int
	_ggg                 uint8
	_fc                  *_a.BufferedWriter
	EncodeGlobals        bool
	_gf                  int
	_egf                 map[int][]int
	_ac                  map[int][]int
	_ffa                 []int
	_ad                  map[int]int
}
