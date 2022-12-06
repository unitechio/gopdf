package classer

import (
	_e "image"
	_c "math"

	_b "bitbucket.org/shenghui0779/gopdf/common"
	_ea "bitbucket.org/shenghui0779/gopdf/internal/jbig2/basic"
	_f "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_g "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

func (_fg *Classer) addPageComponents(_bde *_f.Bitmap, _ee *_f.Boxes, _gf *_f.Bitmaps, _ce int, _bcg Method) error {
	const _ffe = "\u0043l\u0061\u0073\u0073\u0065r\u002e\u0041\u0064\u0064\u0050a\u0067e\u0043o\u006d\u0070\u006f\u006e\u0065\u006e\u0074s"
	if _bde == nil {
		return _g.Error(_ffe, "\u006e\u0069\u006c\u0020\u0069\u006e\u0070\u0075\u0074 \u0070\u0061\u0067\u0065")
	}
	if _ee == nil || _gf == nil || len(*_ee) == 0 {
		_b.Log.Trace("\u0041\u0064\u0064P\u0061\u0067\u0065\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0073\u002e\u0020\u004e\u006f\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064", _bde)
		return nil
	}
	var _ag error
	switch _bcg {
	case RankHaus:
		_ag = _fg.classifyRankHaus(_ee, _gf, _ce)
	case Correlation:
		_ag = _fg.classifyCorrelation(_ee, _gf, _ce)
	default:
		_b.Log.Debug("\u0055\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u003a\u0020'%\u0076\u0027", _bcg)
		return _g.Error(_ffe, "\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064")
	}
	if _ag != nil {
		return _g.Wrap(_ag, _ffe, "")
	}
	if _ag = _fg.getULCorners(_bde, _ee); _ag != nil {
		return _g.Wrap(_ag, _ffe, "")
	}
	_eba := len(*_ee)
	_fg.BaseIndex += _eba
	if _ag = _fg.ComponentsNumber.Add(_eba); _ag != nil {
		return _g.Wrap(_ag, _ffe, "")
	}
	return nil
}
func (_dde Settings) Validate() error {
	const _ccb = "\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0073\u002e\u0056\u0061\u006ci\u0064\u0061\u0074\u0065"
	if _dde.Thresh < 0.4 || _dde.Thresh > 0.98 {
		return _g.Error(_ccb, "\u006a\u0062i\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0034\u0020\u002d\u0020\u0030\u002e\u0039\u0038\u005d")
	}
	if _dde.WeightFactor < 0.0 || _dde.WeightFactor > 1.0 {
		return _g.Error(_ccb, "\u006a\u0062i\u0067\u0032\u0020\u0065\u006ec\u006f\u0064\u0065\u0072\u0020w\u0065\u0069\u0067\u0068\u0074\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0030\u0020\u002d\u0020\u0031\u002e\u0030\u005d")
	}
	if _dde.RankHaus < 0.5 || _dde.RankHaus > 1.0 {
		return _g.Error(_ccb, "\u006a\u0062\u0069\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065 \u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 [\u0030\u002e\u0035\u0020\u002d\u0020\u0031\u002e\u0030\u005d")
	}
	if _dde.SizeHaus < 1 || _dde.SizeHaus > 10 {
		return _g.Error(_ccb, "\u006a\u0062\u0069\u0067\u0032 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0073\u0069\u007a\u0065\u0020h\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0031\u0020\u002d\u0020\u0031\u0030]")
	}
	switch _dde.Components {
	case _f.ComponentConn, _f.ComponentCharacters, _f.ComponentWords:
	default:
		return _g.Error(_ccb, "\u0069n\u0076\u0061\u006c\u0069d\u0020\u0063\u006c\u0061\u0073s\u0065r\u0020c\u006f\u006d\u0070\u006f\u006e\u0065\u006et")
	}
	return nil
}

type similarTemplatesFinder struct {
	Classer        *Classer
	Width          int
	Height         int
	Index          int
	CurrentNumbers []int
	N              int
}

func (_edf *Classer) classifyRankHouseOne(_ceae *_f.Boxes, _baf, _egb, _cae *_f.Bitmaps, _bbg *_f.Points, _efg int) (_cdbc error) {
	const _cfb = "\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065"
	var (
		_bgfd, _gba, _ccf, _abb        float32
		_bdg                           int
		_ecc, _fbgad, _eca, _bca, _ecf *_f.Bitmap
		_fgg, _efgd                    bool
	)
	for _ace := 0; _ace < len(_baf.Values); _ace++ {
		_fbgad = _egb.Values[_ace]
		_eca = _cae.Values[_ace]
		_bgfd, _gba, _cdbc = _bbg.GetGeometry(_ace)
		if _cdbc != nil {
			return _g.Wrapf(_cdbc, _cfb, "\u0066\u0069\u0072\u0073\u0074\u0020\u0067\u0065\u006fm\u0065\u0074\u0072\u0079")
		}
		_bga := len(_edf.UndilatedTemplates.Values)
		_fgg = false
		_aggb := _fdaa(_edf, _fbgad)
		for _bdg = _aggb.Next(); _bdg > -1; {
			_bca, _cdbc = _edf.UndilatedTemplates.GetBitmap(_bdg)
			if _cdbc != nil {
				return _g.Wrap(_cdbc, _cfb, "\u0062\u006d\u0033")
			}
			_ecf, _cdbc = _edf.DilatedTemplates.GetBitmap(_bdg)
			if _cdbc != nil {
				return _g.Wrap(_cdbc, _cfb, "\u0062\u006d\u0034")
			}
			_ccf, _abb, _cdbc = _edf.CentroidPointsTemplates.GetGeometry(_bdg)
			if _cdbc != nil {
				return _g.Wrap(_cdbc, _cfb, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0073")
			}
			_efgd, _cdbc = _f.HausTest(_fbgad, _eca, _bca, _ecf, _bgfd-_ccf, _gba-_abb, MaxDiffWidth, MaxDiffHeight)
			if _cdbc != nil {
				return _g.Wrap(_cdbc, _cfb, "")
			}
			if _efgd {
				_fgg = true
				if _cdbc = _edf.ClassIDs.Add(_bdg); _cdbc != nil {
					return _g.Wrap(_cdbc, _cfb, "")
				}
				if _cdbc = _edf.ComponentPageNumbers.Add(_efg); _cdbc != nil {
					return _g.Wrap(_cdbc, _cfb, "")
				}
				if _edf.Settings.KeepClassInstances {
					_gdd, _cbc := _edf.ClassInstances.GetBitmaps(_bdg)
					if _cbc != nil {
						return _g.Wrap(_cbc, _cfb, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_ecc, _cbc = _baf.GetBitmap(_ace)
					if _cbc != nil {
						return _g.Wrap(_cbc, _cfb, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_gdd.AddBitmap(_ecc)
					_aga, _cbc := _ceae.Get(_ace)
					if _cbc != nil {
						return _g.Wrap(_cbc, _cfb, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_gdd.AddBox(_aga)
				}
				break
			}
		}
		if !_fgg {
			if _cdbc = _edf.ClassIDs.Add(_bga); _cdbc != nil {
				return _g.Wrap(_cdbc, _cfb, "")
			}
			if _cdbc = _edf.ComponentPageNumbers.Add(_efg); _cdbc != nil {
				return _g.Wrap(_cdbc, _cfb, "")
			}
			_fde := &_f.Bitmaps{}
			_ecc, _cdbc = _baf.GetBitmap(_ace)
			if _cdbc != nil {
				return _g.Wrap(_cdbc, _cfb, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_fde.Values = append(_fde.Values, _ecc)
			_dc, _bfg := _ecc.Width, _ecc.Height
			_edf.TemplatesSize.Add(uint64(_bfg)*uint64(_dc), _bga)
			_dgb, _dff := _ceae.Get(_ace)
			if _dff != nil {
				return _g.Wrap(_dff, _cfb, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_fde.AddBox(_dgb)
			_edf.ClassInstances.AddBitmaps(_fde)
			_edf.CentroidPointsTemplates.AddPoint(_bgfd, _gba)
			_edf.UndilatedTemplates.AddBitmap(_fbgad)
			_edf.DilatedTemplates.AddBitmap(_eca)
		}
	}
	return nil
}

var _bge bool

const (
	MaxConnCompWidth = 350
	MaxCharCompWidth = 350
	MaxWordCompWidth = 1000
	MaxCompHeight    = 120
)

func (_bddb *Classer) classifyRankHaus(_gfb *_f.Boxes, _fdbb *_f.Bitmaps, _agef int) error {
	const _ade = "\u0063\u006ca\u0073\u0073\u0069f\u0079\u0052\u0061\u006e\u006b\u0048\u0061\u0075\u0073"
	if _gfb == nil {
		return _g.Error(_ade, "\u0062\u006fx\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _fdbb == nil {
		return _g.Error(_ade, "\u0070\u0069x\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	_cdc := len(_fdbb.Values)
	if _cdc == 0 {
		return _g.Error(_ade, "e\u006dp\u0074\u0079\u0020\u006e\u0065\u0077\u0020\u0063o\u006d\u0070\u006f\u006een\u0074\u0073")
	}
	_cde := _fdbb.CountPixels()
	_ec := _bddb.Settings.SizeHaus
	_gff := _f.SelCreateBrick(_ec, _ec, _ec/2, _ec/2, _f.SelHit)
	_adaa := &_f.Bitmaps{Values: make([]*_f.Bitmap, _cdc)}
	_bag := &_f.Bitmaps{Values: make([]*_f.Bitmap, _cdc)}
	var (
		_bbb, _fbga, _bba *_f.Bitmap
		_cfe              error
	)
	for _fdbe := 0; _fdbe < _cdc; _fdbe++ {
		_bbb, _cfe = _fdbb.GetBitmap(_fdbe)
		if _cfe != nil {
			return _g.Wrap(_cfe, _ade, "")
		}
		_fbga, _cfe = _bbb.AddBorderGeneral(JbAddedPixels, JbAddedPixels, JbAddedPixels, JbAddedPixels, 0)
		if _cfe != nil {
			return _g.Wrap(_cfe, _ade, "")
		}
		_bba, _cfe = _f.Dilate(nil, _fbga, _gff)
		if _cfe != nil {
			return _g.Wrap(_cfe, _ade, "")
		}
		_adaa.Values[_cdc] = _fbga
		_bag.Values[_cdc] = _bba
	}
	_eec, _cfe := _f.Centroids(_adaa.Values)
	if _cfe != nil {
		return _g.Wrap(_cfe, _ade, "")
	}
	if _cfe = _eec.Add(_bddb.CentroidPoints); _cfe != nil {
		_b.Log.Trace("\u004e\u006f\u0020\u0063en\u0074\u0072\u006f\u0069\u0064\u0073\u0020\u0074\u006f\u0020\u0061\u0064\u0064")
	}
	if _bddb.Settings.RankHaus == 1.0 {
		_cfe = _bddb.classifyRankHouseOne(_gfb, _fdbb, _adaa, _bag, _eec, _agef)
	} else {
		_cfe = _bddb.classifyRankHouseNonOne(_gfb, _fdbb, _adaa, _bag, _eec, _cde, _agef)
	}
	if _cfe != nil {
		return _g.Wrap(_cfe, _ade, "")
	}
	return nil
}
func (_bd *Classer) ComputeLLCorners() (_ca error) {
	const _ad = "\u0043l\u0061\u0073\u0073\u0065\u0072\u002e\u0043\u006f\u006d\u0070\u0075t\u0065\u004c\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073"
	if _bd.PtaUL == nil {
		return _g.Error(_ad, "\u0055\u004c\u0020\u0043or\u006e\u0065\u0072\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	_ffg := len(*_bd.PtaUL)
	_bd.PtaLL = &_f.Points{}
	var (
		_cd, _ge float32
		_eb, _cb int
		_bc      *_f.Bitmap
	)
	for _cc := 0; _cc < _ffg; _cc++ {
		_cd, _ge, _ca = _bd.PtaUL.GetGeometry(_cc)
		if _ca != nil {
			_b.Log.Debug("\u0047e\u0074\u0074\u0069\u006e\u0067\u0020\u0050\u0074\u0061\u0055\u004c \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ca)
			return _g.Wrap(_ca, _ad, "\u0050\u0074\u0061\u0055\u004c\u0020\u0047\u0065\u006fm\u0065\u0074\u0072\u0079")
		}
		_eb, _ca = _bd.ClassIDs.Get(_cc)
		if _ca != nil {
			_b.Log.Debug("\u0047\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0043\u006c\u0061s\u0073\u0049\u0044\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _ca)
			return _g.Wrap(_ca, _ad, "\u0043l\u0061\u0073\u0073\u0049\u0044")
		}
		_bc, _ca = _bd.UndilatedTemplates.GetBitmap(_eb)
		if _ca != nil {
			_b.Log.Debug("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0055\u006ed\u0069\u006c\u0061\u0074\u0065\u0064\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _ca)
			return _g.Wrap(_ca, _ad, "\u0055\u006e\u0064\u0069la\u0074\u0065\u0064\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073")
		}
		_cb = _bc.Height
		_bd.PtaLL.AddPoint(_cd, _ge+float32(_cb))
	}
	return nil
}
func (_fce *Settings) SetDefault() {
	if _fce.MaxCompWidth == 0 {
		switch _fce.Components {
		case _f.ComponentConn:
			_fce.MaxCompWidth = MaxConnCompWidth
		case _f.ComponentCharacters:
			_fce.MaxCompWidth = MaxCharCompWidth
		case _f.ComponentWords:
			_fce.MaxCompWidth = MaxWordCompWidth
		}
	}
	if _fce.MaxCompHeight == 0 {
		_fce.MaxCompHeight = MaxCompHeight
	}
	if _fce.Thresh == 0.0 {
		_fce.Thresh = 0.9
	}
	if _fce.WeightFactor == 0.0 {
		_fce.WeightFactor = 0.75
	}
	if _fce.RankHaus == 0.0 {
		_fce.RankHaus = 0.97
	}
	if _fce.SizeHaus == 0 {
		_fce.SizeHaus = 2
	}
}

const (
	MaxDiffWidth  = 2
	MaxDiffHeight = 2
)

func (_bgd *Classer) verifyMethod(_cgb Method) error {
	if _cgb != RankHaus && _cgb != Correlation {
		return _g.Error("\u0076\u0065\u0072i\u0066\u0079\u004d\u0065\u0074\u0068\u006f\u0064", "\u0069\u006e\u0076\u0061li\u0064\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072\u0020\u006d\u0065\u0074\u0068o\u0064")
	}
	return nil
}
func Init(settings Settings) (*Classer, error) {
	const _ff = "\u0063\u006c\u0061s\u0073\u0065\u0072\u002e\u0049\u006e\u0069\u0074"
	_ga := &Classer{Settings: settings, Widths: map[int]int{}, Heights: map[int]int{}, TemplatesSize: _ea.IntsMap{}, TemplateAreas: &_ea.IntSlice{}, ComponentPageNumbers: &_ea.IntSlice{}, ClassIDs: &_ea.IntSlice{}, ComponentsNumber: &_ea.IntSlice{}, CentroidPoints: &_f.Points{}, CentroidPointsTemplates: &_f.Points{}, UndilatedTemplates: &_f.Bitmaps{}, DilatedTemplates: &_f.Bitmaps{}, ClassInstances: &_f.BitmapsArray{}, FgTemplates: &_ea.NumSlice{}}
	if _gg := _ga.Settings.Validate(); _gg != nil {
		return nil, _g.Wrap(_gg, _ff, "")
	}
	return _ga, nil
}

type Method int

const (
	RankHaus Method = iota
	Correlation
)

func (_fe *Classer) classifyCorrelation(_fdge *_f.Boxes, _fcf *_f.Bitmaps, _ege int) error {
	const _dgf = "\u0063\u006c\u0061\u0073si\u0066\u0079\u0043\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e"
	if _fdge == nil {
		return _g.Error(_dgf, "\u006e\u0065\u0077\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062o\u0078\u0065\u0073\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _fcf == nil {
		return _g.Error(_dgf, "\u006e\u0065wC\u006f\u006d\u0070o\u006e\u0065\u006e\u0074s b\u0069tm\u0061\u0070\u0020\u0061\u0072\u0072\u0061y \u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
	}
	_gegb := len(_fcf.Values)
	if _gegb == 0 {
		_b.Log.Debug("\u0063l\u0061\u0073s\u0069\u0066\u0079C\u006f\u0072\u0072\u0065\u006c\u0061\u0074i\u006f\u006e\u0020\u002d\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0070\u0069\u0078\u0061s\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079")
		return nil
	}
	var (
		_fbe, _af *_f.Bitmap
		_egec     error
	)
	_caa := &_f.Bitmaps{Values: make([]*_f.Bitmap, _gegb)}
	for _gdc, _cab := range _fcf.Values {
		_af, _egec = _cab.AddBorderGeneral(JbAddedPixels, JbAddedPixels, JbAddedPixels, JbAddedPixels, 0)
		if _egec != nil {
			return _g.Wrap(_egec, _dgf, "")
		}
		_caa.Values[_gdc] = _af
	}
	_ggf := _fe.FgTemplates
	_dga := _f.MakePixelSumTab8()
	_gea := _f.MakePixelCentroidTab8()
	_cdd := make([]int, _gegb)
	_acd := make([][]int, _gegb)
	_ceec := _f.Points(make([]_f.Point, _gegb))
	_bgf := &_ceec
	var (
		_gge, _cbf       int
		_cfg, _fge, _aea int
		_bef, _de        int
		_bdc             byte
	)
	for _dfd, _bdd := range _caa.Values {
		_acd[_dfd] = make([]int, _bdd.Height)
		_gge = 0
		_cbf = 0
		_fge = (_bdd.Height - 1) * _bdd.RowStride
		_cfg = 0
		for _de = _bdd.Height - 1; _de >= 0; _de, _fge = _de-1, _fge-_bdd.RowStride {
			_acd[_dfd][_de] = _cfg
			_aea = 0
			for _bef = 0; _bef < _bdd.RowStride; _bef++ {
				_bdc = _bdd.Data[_fge+_bef]
				_aea += _dga[_bdc]
				_gge += _gea[_bdc] + _bef*8*_dga[_bdc]
			}
			_cfg += _aea
			_cbf += _aea * _de
		}
		_cdd[_dfd] = _cfg
		if _cfg > 0 {
			(*_bgf)[_dfd] = _f.Point{X: float32(_gge) / float32(_cfg), Y: float32(_cbf) / float32(_cfg)}
		} else {
			(*_bgf)[_dfd] = _f.Point{X: float32(_bdd.Width) / float32(2), Y: float32(_bdd.Height) / float32(2)}
		}
	}
	if _egec = _fe.CentroidPoints.Add(_bgf); _egec != nil {
		return _g.Wrap(_egec, _dgf, "\u0063\u0065\u006et\u0072\u006f\u0069\u0064\u0020\u0061\u0064\u0064")
	}
	var (
		_eab, _bdcf, _bcd      int
		_fdb                   float64
		_egg, _ggb, _ada, _add float32
		_cgf, _gab             _f.Point
		_age                   bool
		_gdce                  *similarTemplatesFinder
		_eaa                   int
		_dab                   *_f.Bitmap
		_gc                    *_e.Rectangle
		_cdb                   *_f.Bitmaps
	)
	for _eaa, _af = range _caa.Values {
		_bdcf = _cdd[_eaa]
		if _egg, _ggb, _egec = _bgf.GetGeometry(_eaa); _egec != nil {
			return _g.Wrap(_egec, _dgf, "\u0070t\u0061\u0020\u002d\u0020\u0069")
		}
		_age = false
		_cac := len(_fe.UndilatedTemplates.Values)
		_gdce = _fdaa(_fe, _af)
		for _gee := _gdce.Next(); _gee > -1; {
			if _dab, _egec = _fe.UndilatedTemplates.GetBitmap(_gee); _egec != nil {
				return _g.Wrap(_egec, _dgf, "\u0075\u006e\u0069dl\u0061\u0074\u0065\u0064\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0062\u006d\u0032")
			}
			if _bcd, _egec = _ggf.GetInt(_gee); _egec != nil {
				_b.Log.Trace("\u0046\u0047\u0020T\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u005b\u0069\u0063\u006ca\u0073\u0073\u005d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _egec)
			}
			if _ada, _add, _egec = _fe.CentroidPointsTemplates.GetGeometry(_gee); _egec != nil {
				return _g.Wrap(_egec, _dgf, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074T\u0065\u006d\u0070\u006c\u0061\u0074e\u0073\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u00782\u002c\u0079\u0032\u0020")
			}
			if _fe.Settings.WeightFactor > 0.0 {
				if _eab, _egec = _fe.TemplateAreas.Get(_gee); _egec != nil {
					_b.Log.Trace("\u0054\u0065\u006dp\u006c\u0061\u0074\u0065A\u0072\u0065\u0061\u0073\u005b\u0069\u0063l\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0061\u0072\u0065\u0061\u0020\u0025\u0076", _egec)
				}
				_fdb = _fe.Settings.Thresh + (1.0-_fe.Settings.Thresh)*_fe.Settings.WeightFactor*float64(_bcd)/float64(_eab)
			} else {
				_fdb = _fe.Settings.Thresh
			}
			_caaf, _bcdd := _f.CorrelationScoreThresholded(_af, _dab, _bdcf, _bcd, _cgf.X-_gab.X, _cgf.Y-_gab.Y, MaxDiffWidth, MaxDiffHeight, _dga, _acd[_eaa], float32(_fdb))
			if _bcdd != nil {
				return _g.Wrap(_bcdd, _dgf, "")
			}
			if _bge {
				var (
					_ged, _abc float64
					_ffc, _dbf int
				)
				_ged, _bcdd = _f.CorrelationScore(_af, _dab, _bdcf, _bcd, _egg-_ada, _ggb-_add, MaxDiffWidth, MaxDiffHeight, _dga)
				if _bcdd != nil {
					return _g.Wrap(_bcdd, _dgf, "d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065")
				}
				_abc, _bcdd = _f.CorrelationScoreSimple(_af, _dab, _bdcf, _bcd, _egg-_ada, _ggb-_add, MaxDiffWidth, MaxDiffHeight, _dga)
				if _bcdd != nil {
					return _g.Wrap(_bcdd, _dgf, "d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065")
				}
				_ffc = int(_c.Sqrt(_ged * float64(_bdcf) * float64(_bcd)))
				_dbf = int(_c.Sqrt(_abc * float64(_bdcf) * float64(_bcd)))
				if (_ged >= _fdb) != (_abc >= _fdb) {
					return _g.Errorf(_dgf, "\u0064\u0065\u0062\u0075\u0067\u0020\u0043\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u006f\u0072\u0065\u0020\u006d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020-\u0020\u0025d\u0028\u00250\u002e\u0034\u0066\u002c\u0020\u0025\u0076\u0029\u0020\u0076\u0073\u0020\u0025d(\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0076)\u0020\u0025\u0030\u002e\u0034\u0066", _ffc, _ged, _ged >= _fdb, _dbf, _abc, _abc >= _fdb, _ged-_abc)
				}
				if _ged >= _fdb != _caaf {
					return _g.Errorf(_dgf, "\u0064\u0065\u0062\u0075\u0067\u0020\u0043o\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e \u0073\u0063\u006f\u0072\u0065 \u004d\u0069\u0073\u006d\u0061t\u0063\u0068 \u0062\u0065\u0074w\u0065\u0065\u006e\u0020\u0063\u006frr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020/\u0020\u0074\u0068\u0072\u0065s\u0068\u006f\u006c\u0064\u002e\u0020\u0043\u006f\u006dpa\u0072\u0069\u0073\u006f\u006e:\u0020\u0025\u0030\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0064\u0029\u0020\u003e\u003d\u0020\u00250\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u0029\u0020\u0076\u0073\u0020\u0025\u0076", _ged, _ged*float64(_bdcf)*float64(_bcd), _ffc, _fdb, float32(_fdb)*float32(_bdcf)*float32(_bcd), _caaf)
				}
			}
			if _caaf {
				_age = true
				if _bcdd = _fe.ClassIDs.Add(_gee); _bcdd != nil {
					return _g.Wrap(_bcdd, _dgf, "\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064")
				}
				if _bcdd = _fe.ComponentPageNumbers.Add(_ege); _bcdd != nil {
					return _g.Wrap(_bcdd, _dgf, "\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064")
				}
				if _fe.Settings.KeepClassInstances {
					if _fbe, _bcdd = _fcf.GetBitmap(_eaa); _bcdd != nil {
						return _g.Wrap(_bcdd, _dgf, "\u004b\u0065\u0065\u0070Cl\u0061\u0073\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u002d \u0069")
					}
					if _cdb, _bcdd = _fe.ClassInstances.GetBitmaps(_gee); _bcdd != nil {
						return _g.Wrap(_bcdd, _dgf, "K\u0065\u0065\u0070\u0043\u006c\u0061s\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065s\u0020\u002d\u0020i\u0043l\u0061\u0073\u0073")
					}
					_cdb.AddBitmap(_fbe)
					if _gc, _bcdd = _fdge.Get(_eaa); _bcdd != nil {
						return _g.Wrap(_bcdd, _dgf, "\u004be\u0065p\u0043\u006c\u0061\u0073\u0073I\u006e\u0073t\u0061\u006e\u0063\u0065\u0073")
					}
					_cdb.AddBox(_gc)
				}
				break
			}
		}
		if !_age {
			if _egec = _fe.ClassIDs.Add(_cac); _egec != nil {
				return _g.Wrap(_egec, _dgf, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			if _egec = _fe.ComponentPageNumbers.Add(_ege); _egec != nil {
				return _g.Wrap(_egec, _dgf, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_cdb = &_f.Bitmaps{}
			if _fbe, _egec = _fcf.GetBitmap(_eaa); _egec != nil {
				return _g.Wrap(_egec, _dgf, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_cdb.AddBitmap(_fbe)
			_cga, _addd := _fbe.Width, _fbe.Height
			_eaf := uint64(_addd) * uint64(_cga)
			_fe.TemplatesSize.Add(_eaf, _cac)
			if _gc, _egec = _fdge.Get(_eaa); _egec != nil {
				return _g.Wrap(_egec, _dgf, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_cdb.AddBox(_gc)
			_fe.ClassInstances.AddBitmaps(_cdb)
			_fe.CentroidPointsTemplates.AddPoint(_egg, _ggb)
			_fe.FgTemplates.AddInt(_bdcf)
			_fe.UndilatedTemplates.AddBitmap(_fbe)
			_eab = (_af.Width - 2*JbAddedPixels) * (_af.Height - 2*JbAddedPixels)
			if _egec = _fe.TemplateAreas.Add(_eab); _egec != nil {
				return _g.Wrap(_egec, _dgf, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
		}
	}
	_fe.NumberOfClasses = len(_fe.UndilatedTemplates.Values)
	return nil
}
func DefaultSettings() Settings { _cffc := &Settings{}; _cffc.SetDefault(); return *_cffc }

var TwoByTwoWalk = []int{0, 0, 0, 1, -1, 0, 0, -1, 1, 0, -1, 1, 1, 1, -1, -1, 1, -1, 0, -2, 2, 0, 0, 2, -2, 0, -1, -2, 1, -2, 2, -1, 2, 1, 1, 2, -1, 2, -2, 1, -2, -1, -2, -2, 2, -2, 2, 2, -2, 2}

type Settings struct {
	MaxCompWidth       int
	MaxCompHeight      int
	SizeHaus           int
	RankHaus           float64
	Thresh             float64
	WeightFactor       float64
	KeepClassInstances bool
	Components         _f.Component
	Method             Method
}

func (_ffcg *similarTemplatesFinder) Next() int {
	var (
		_dad, _ggbb, _cce, _baa int
		_fca                    bool
		_ceea                   *_f.Bitmap
		_ebc                    error
	)
	for {
		if _ffcg.Index >= 25 {
			return -1
		}
		_ggbb = _ffcg.Width + TwoByTwoWalk[2*_ffcg.Index]
		_dad = _ffcg.Height + TwoByTwoWalk[2*_ffcg.Index+1]
		if _dad < 1 || _ggbb < 1 {
			_ffcg.Index++
			continue
		}
		if len(_ffcg.CurrentNumbers) == 0 {
			_ffcg.CurrentNumbers, _fca = _ffcg.Classer.TemplatesSize.GetSlice(uint64(_ggbb) * uint64(_dad))
			if !_fca {
				_ffcg.Index++
				continue
			}
			_ffcg.N = 0
		}
		_cce = len(_ffcg.CurrentNumbers)
		for ; _ffcg.N < _cce; _ffcg.N++ {
			_baa = _ffcg.CurrentNumbers[_ffcg.N]
			_ceea, _ebc = _ffcg.Classer.DilatedTemplates.GetBitmap(_baa)
			if _ebc != nil {
				_b.Log.Debug("\u0046\u0069\u006e\u0064\u004e\u0065\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0074\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020")
				return 0
			}
			if _ceea.Width-2*JbAddedPixels == _ggbb && _ceea.Height-2*JbAddedPixels == _dad {
				return _baa
			}
		}
		_ffcg.Index++
		_ffcg.CurrentNumbers = nil
	}
}

type Classer struct {
	BaseIndex               int
	Settings                Settings
	ComponentsNumber        *_ea.IntSlice
	TemplateAreas           *_ea.IntSlice
	Widths                  map[int]int
	Heights                 map[int]int
	NumberOfClasses         int
	ClassInstances          *_f.BitmapsArray
	UndilatedTemplates      *_f.Bitmaps
	DilatedTemplates        *_f.Bitmaps
	TemplatesSize           _ea.IntsMap
	FgTemplates             *_ea.NumSlice
	CentroidPoints          *_f.Points
	CentroidPointsTemplates *_f.Points
	ClassIDs                *_ea.IntSlice
	ComponentPageNumbers    *_ea.IntSlice
	PtaUL                   *_f.Points
	PtaLL                   *_f.Points
}

func _fdaa(_afe *Classer, _dea *_f.Bitmap) *similarTemplatesFinder {
	return &similarTemplatesFinder{Width: _dea.Width, Height: _dea.Height, Classer: _afe}
}
func (_ced *Classer) classifyRankHouseNonOne(_fa *_f.Boxes, _bfcf, _dfg, _aef *_f.Bitmaps, _agf *_f.Points, _deb *_ea.NumSlice, _ega int) (_fda error) {
	const _fdf = "\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065"
	var (
		_dd, _egea, _fgb, _fgd          float32
		_gabg, _bdf, _adg               int
		_gca, _egae, _egee, _gae, _adda *_f.Bitmap
		_bdge, _gad                     bool
	)
	_aec := _f.MakePixelSumTab8()
	for _dbg := 0; _dbg < len(_bfcf.Values); _dbg++ {
		if _egae, _fda = _dfg.GetBitmap(_dbg); _fda != nil {
			return _g.Wrap(_fda, _fdf, "b\u006d\u0073\u0031\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
		}
		if _gabg, _fda = _deb.GetInt(_dbg); _fda != nil {
			_b.Log.Trace("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0046\u0047T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073 \u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _dbg, _fda)
		}
		if _egee, _fda = _aef.GetBitmap(_dbg); _fda != nil {
			return _g.Wrap(_fda, _fdf, "b\u006d\u0073\u0032\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
		}
		if _dd, _egea, _fda = _agf.GetGeometry(_dbg); _fda != nil {
			return _g.Wrapf(_fda, _fdf, "\u0070t\u0061[\u0069\u005d\u002e\u0047\u0065\u006f\u006d\u0065\u0074\u0072\u0079")
		}
		_bbge := len(_ced.UndilatedTemplates.Values)
		_bdge = false
		_afb := _fdaa(_ced, _egae)
		for _adg = _afb.Next(); _adg > -1; {
			if _gae, _fda = _ced.UndilatedTemplates.GetBitmap(_adg); _fda != nil {
				return _g.Wrap(_fda, _fdf, "\u0070\u0069\u0078\u0061\u0074\u002e\u005b\u0069\u0043l\u0061\u0073\u0073\u005d")
			}
			if _bdf, _fda = _ced.FgTemplates.GetInt(_adg); _fda != nil {
				_b.Log.Trace("\u0047\u0065\u0074\u0074\u0069\u006eg\u0020\u0046\u0047\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u005b\u0025d\u005d\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _adg, _fda)
			}
			if _adda, _fda = _ced.DilatedTemplates.GetBitmap(_adg); _fda != nil {
				return _g.Wrap(_fda, _fdf, "\u0070\u0069\u0078\u0061\u0074\u0064\u005b\u0069\u0043l\u0061\u0073\u0073\u005d")
			}
			if _fgb, _fgd, _fda = _ced.CentroidPointsTemplates.GetGeometry(_adg); _fda != nil {
				return _g.Wrap(_fda, _fdf, "\u0043\u0065\u006et\u0072\u006f\u0069\u0064P\u006f\u0069\u006e\u0074\u0073\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u005b\u0069\u0043\u006c\u0061\u0073\u0073\u005d")
			}
			_gad, _fda = _f.RankHausTest(_egae, _egee, _gae, _adda, _dd-_fgb, _egea-_fgd, MaxDiffWidth, MaxDiffHeight, _gabg, _bdf, float32(_ced.Settings.RankHaus), _aec)
			if _fda != nil {
				return _g.Wrap(_fda, _fdf, "")
			}
			if _gad {
				_bdge = true
				if _fda = _ced.ClassIDs.Add(_adg); _fda != nil {
					return _g.Wrap(_fda, _fdf, "")
				}
				if _fda = _ced.ComponentPageNumbers.Add(_ega); _fda != nil {
					return _g.Wrap(_fda, _fdf, "")
				}
				if _ced.Settings.KeepClassInstances {
					_aca, _cffg := _ced.ClassInstances.GetBitmaps(_adg)
					if _cffg != nil {
						return _g.Wrap(_cffg, _fdf, "\u0063\u002e\u0050\u0069\u0078\u0061\u0061\u002e\u0047\u0065\u0074B\u0069\u0074\u006d\u0061\u0070\u0073\u0028\u0069\u0043\u006ca\u0073\u0073\u0029")
					}
					if _gca, _cffg = _bfcf.GetBitmap(_dbg); _cffg != nil {
						return _g.Wrap(_cffg, _fdf, "\u0070i\u0078\u0061\u005b\u0069\u005d")
					}
					_aca.Values = append(_aca.Values, _gca)
					_cdbf, _cffg := _fa.Get(_dbg)
					if _cffg != nil {
						return _g.Wrap(_cffg, _fdf, "b\u006f\u0078\u0061\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
					}
					_aca.Boxes = append(_aca.Boxes, _cdbf)
				}
				break
			}
		}
		if !_bdge {
			if _fda = _ced.ClassIDs.Add(_bbge); _fda != nil {
				return _g.Wrap(_fda, _fdf, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			if _fda = _ced.ComponentPageNumbers.Add(_ega); _fda != nil {
				return _g.Wrap(_fda, _fdf, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_gdcd := &_f.Bitmaps{}
			_gca = _bfcf.Values[_dbg]
			_gdcd.AddBitmap(_gca)
			_fcd, _eged := _gca.Width, _gca.Height
			_ced.TemplatesSize.Add(uint64(_fcd)*uint64(_eged), _bbge)
			_gdg, _cfee := _fa.Get(_dbg)
			if _cfee != nil {
				return _g.Wrap(_cfee, _fdf, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_gdcd.AddBox(_gdg)
			_ced.ClassInstances.AddBitmaps(_gdcd)
			_ced.CentroidPointsTemplates.AddPoint(_dd, _egea)
			_ced.UndilatedTemplates.AddBitmap(_egae)
			_ced.DilatedTemplates.AddBitmap(_egee)
			_ced.FgTemplates.AddInt(_gabg)
		}
	}
	_ced.NumberOfClasses = len(_ced.UndilatedTemplates.Values)
	return nil
}

const JbAddedPixels = 6

func (_abf *Classer) getULCorners(_fc *_f.Bitmap, _bg *_f.Boxes) error {
	const _ggc = "\u0067\u0065\u0074U\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073"
	if _fc == nil {
		return _g.Error(_ggc, "\u006e\u0069l\u0020\u0069\u006da\u0067\u0065\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _bg == nil {
		return _g.Error(_ggc, "\u006e\u0069\u006c\u0020\u0062\u006f\u0075\u006e\u0064\u0073")
	}
	if _abf.PtaUL == nil {
		_abf.PtaUL = &_f.Points{}
	}
	_cg := len(*_bg)
	var (
		_ebb, _ae, _eed, _cf int
		_d, _fdg, _agg, _cfc float32
		_ef                  error
		_fgc                 *_e.Rectangle
		_geg                 *_f.Bitmap
		_ba                  _e.Point
	)
	for _dg := 0; _dg < _cg; _dg++ {
		_ebb = _abf.BaseIndex + _dg
		if _d, _fdg, _ef = _abf.CentroidPoints.GetGeometry(_ebb); _ef != nil {
			return _g.Wrap(_ef, _ggc, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050o\u0069\u006e\u0074\u0073")
		}
		if _ae, _ef = _abf.ClassIDs.Get(_ebb); _ef != nil {
			return _g.Wrap(_ef, _ggc, "\u0043\u006c\u0061s\u0073\u0049\u0044\u0073\u002e\u0047\u0065\u0074")
		}
		if _agg, _cfc, _ef = _abf.CentroidPointsTemplates.GetGeometry(_ae); _ef != nil {
			return _g.Wrap(_ef, _ggc, "\u0043\u0065\u006etr\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073")
		}
		_eee := _agg - _d
		_ebab := _cfc - _fdg
		if _eee >= 0 {
			_eed = int(_eee + 0.5)
		} else {
			_eed = int(_eee - 0.5)
		}
		if _ebab >= 0 {
			_cf = int(_ebab + 0.5)
		} else {
			_cf = int(_ebab - 0.5)
		}
		if _fgc, _ef = _bg.Get(_dg); _ef != nil {
			return _g.Wrap(_ef, _ggc, "")
		}
		_fbg, _cgc := _fgc.Min.X, _fgc.Min.Y
		_geg, _ef = _abf.UndilatedTemplates.GetBitmap(_ae)
		if _ef != nil {
			return _g.Wrap(_ef, _ggc, "\u0055\u006e\u0064\u0069\u006c\u0061\u0074\u0065\u0064\u0054e\u006d\u0070\u006c\u0061\u0074\u0065\u0073.\u0047\u0065\u0074\u0028\u0069\u0043\u006c\u0061\u0073\u0073\u0029")
		}
		_ba, _ef = _efc(_fc, _fbg, _cgc, _eed, _cf, _geg)
		if _ef != nil {
			return _g.Wrap(_ef, _ggc, "")
		}
		_abf.PtaUL.AddPoint(float32(_fbg-_eed+_ba.X), float32(_cgc-_cf+_ba.Y))
	}
	return nil
}
func _efc(_db *_f.Bitmap, _cea, _df, _bea, _aee int, _fbb *_f.Bitmap) (_da _e.Point, _bf error) {
	const _bfc = "\u0066i\u006e\u0061\u006c\u0041l\u0069\u0067\u006e\u006d\u0065n\u0074P\u006fs\u0069\u0074\u0069\u006f\u006e\u0069\u006eg"
	if _db == nil {
		return _da, _g.Error(_bfc, "\u0073\u006f\u0075\u0072ce\u0020\u006e\u006f\u0074\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	if _fbb == nil {
		return _da, _g.Error(_bfc, "t\u0065\u006d\u0070\u006cat\u0065 \u006e\u006f\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_bb, _eg := _fbb.Width, _fbb.Height
	_cff, _cge := _cea-_bea-JbAddedPixels, _df-_aee-JbAddedPixels
	_b.Log.Trace("\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0077\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0068\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0078\u003a\u0020\u0027\u0025d'\u002c\u0020\u0062\u0079\u003a\u0020\u0027\u0025\u0064\u0027", _cea, _df, _bb, _eg, _cff, _cge)
	_cbe, _bf := _f.Rect(_cff, _cge, _bb, _eg)
	if _bf != nil {
		return _da, _g.Wrap(_bf, _bfc, "")
	}
	_cda, _, _bf := _db.ClipRectangle(_cbe)
	if _bf != nil {
		_b.Log.Error("\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069\u0070\u0020\u0072\u0065c\u0074\u0061\u006e\u0067\u006c\u0065\u003a\u0020\u0025\u0076", _cbe)
		return _da, _g.Wrap(_bf, _bfc, "")
	}
	_cee := _f.New(_cda.Width, _cda.Height)
	_bgc := _c.MaxInt32
	var _gb, _ed, _ffd, _gd, _fga int
	for _gb = -1; _gb <= 1; _gb++ {
		for _ed = -1; _ed <= 1; _ed++ {
			if _, _bf = _f.Copy(_cee, _cda); _bf != nil {
				return _da, _g.Wrap(_bf, _bfc, "")
			}
			if _bf = _cee.RasterOperation(_ed, _gb, _bb, _eg, _f.PixSrcXorDst, _fbb, 0, 0); _bf != nil {
				return _da, _g.Wrap(_bf, _bfc, "")
			}
			_ffd = _cee.CountPixels()
			if _ffd < _bgc {
				_gd = _ed
				_fga = _gb
				_bgc = _ffd
			}
		}
	}
	_da.X = _gd
	_da.Y = _fga
	return _da, nil
}
func (_fb *Classer) AddPage(inputPage *_f.Bitmap, pageNumber int, method Method) (_fd error) {
	const _ab = "\u0043l\u0061s\u0073\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065"
	_fb.Widths[pageNumber] = inputPage.Width
	_fb.Heights[pageNumber] = inputPage.Height
	if _fd = _fb.verifyMethod(method); _fd != nil {
		return _g.Wrap(_fd, _ab, "")
	}
	_ac, _be, _fd := inputPage.GetComponents(_fb.Settings.Components, _fb.Settings.MaxCompWidth, _fb.Settings.MaxCompHeight)
	if _fd != nil {
		return _g.Wrap(_fd, _ab, "")
	}
	_b.Log.Debug("\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _ac)
	if _fd = _fb.addPageComponents(inputPage, _be, _ac, pageNumber, method); _fd != nil {
		return _g.Wrap(_fd, _ab, "")
	}
	return nil
}
