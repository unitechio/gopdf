package classer

import (
	_e "image"
	_a "math"

	_d "bitbucket.org/shenghui0779/gopdf/common"
	_ce "bitbucket.org/shenghui0779/gopdf/internal/jbig2/basic"
	_ed "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_aa "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

func (_fbg *Settings) SetDefault() {
	if _fbg.MaxCompWidth == 0 {
		switch _fbg.Components {
		case _ed.ComponentConn:
			_fbg.MaxCompWidth = MaxConnCompWidth
		case _ed.ComponentCharacters:
			_fbg.MaxCompWidth = MaxCharCompWidth
		case _ed.ComponentWords:
			_fbg.MaxCompWidth = MaxWordCompWidth
		}
	}
	if _fbg.MaxCompHeight == 0 {
		_fbg.MaxCompHeight = MaxCompHeight
	}
	if _fbg.Thresh == 0.0 {
		_fbg.Thresh = 0.9
	}
	if _fbg.WeightFactor == 0.0 {
		_fbg.WeightFactor = 0.75
	}
	if _fbg.RankHaus == 0.0 {
		_fbg.RankHaus = 0.97
	}
	if _fbg.SizeHaus == 0 {
		_fbg.SizeHaus = 2
	}
}

const (
	RankHaus Method = iota
	Correlation
)

func _ec(_eag *_ed.Bitmap, _eg, _cg, _aed, _bfd int, _cc *_ed.Bitmap) (_edb _e.Point, _ffe error) {
	const _add = "\u0066i\u006e\u0061\u006c\u0041l\u0069\u0067\u006e\u006d\u0065n\u0074P\u006fs\u0069\u0074\u0069\u006f\u006e\u0069\u006eg"
	if _eag == nil {
		return _edb, _aa.Error(_add, "\u0073\u006f\u0075\u0072ce\u0020\u006e\u006f\u0074\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	if _cc == nil {
		return _edb, _aa.Error(_add, "t\u0065\u006d\u0070\u006cat\u0065 \u006e\u006f\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_eb, _eddf := _cc.Width, _cc.Height
	_bdc, _aca := _eg-_aed-JbAddedPixels, _cg-_bfd-JbAddedPixels
	_d.Log.Trace("\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0077\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0068\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0078\u003a\u0020\u0027\u0025d'\u002c\u0020\u0062\u0079\u003a\u0020\u0027\u0025\u0064\u0027", _eg, _cg, _eb, _eddf, _bdc, _aca)
	_eae, _ffe := _ed.Rect(_bdc, _aca, _eb, _eddf)
	if _ffe != nil {
		return _edb, _aa.Wrap(_ffe, _add, "")
	}
	_efc, _, _ffe := _eag.ClipRectangle(_eae)
	if _ffe != nil {
		_d.Log.Error("\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069\u0070\u0020\u0072\u0065c\u0074\u0061\u006e\u0067\u006c\u0065\u003a\u0020\u0025\u0076", _eae)
		return _edb, _aa.Wrap(_ffe, _add, "")
	}
	_bad := _ed.New(_efc.Width, _efc.Height)
	_dd := _a.MaxInt32
	var _ega, _dbg, _cdd, _dfg, _af int
	for _ega = -1; _ega <= 1; _ega++ {
		for _dbg = -1; _dbg <= 1; _dbg++ {
			if _, _ffe = _ed.Copy(_bad, _efc); _ffe != nil {
				return _edb, _aa.Wrap(_ffe, _add, "")
			}
			if _ffe = _bad.RasterOperation(_dbg, _ega, _eb, _eddf, _ed.PixSrcXorDst, _cc, 0, 0); _ffe != nil {
				return _edb, _aa.Wrap(_ffe, _add, "")
			}
			_cdd = _bad.CountPixels()
			if _cdd < _dd {
				_dfg = _dbg
				_af = _ega
				_dd = _cdd
			}
		}
	}
	_edb.X = _dfg
	_edb.Y = _af
	return _edb, nil
}
func (_cf *Classer) addPageComponents(_fda *_ed.Bitmap, _gf *_ed.Boxes, _bg *_ed.Bitmaps, _cac int, _fg Method) error {
	const _db = "\u0043l\u0061\u0073\u0073\u0065r\u002e\u0041\u0064\u0064\u0050a\u0067e\u0043o\u006d\u0070\u006f\u006e\u0065\u006e\u0074s"
	if _fda == nil {
		return _aa.Error(_db, "\u006e\u0069\u006c\u0020\u0069\u006e\u0070\u0075\u0074 \u0070\u0061\u0067\u0065")
	}
	if _gf == nil || _bg == nil || len(*_gf) == 0 {
		_d.Log.Trace("\u0041\u0064\u0064P\u0061\u0067\u0065\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0073\u002e\u0020\u004e\u006f\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064", _fda)
		return nil
	}
	var _cef error
	switch _fg {
	case RankHaus:
		_cef = _cf.classifyRankHaus(_gf, _bg, _cac)
	case Correlation:
		_cef = _cf.classifyCorrelation(_gf, _bg, _cac)
	default:
		_d.Log.Debug("\u0055\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u003a\u0020'%\u0076\u0027", _fg)
		return _aa.Error(_db, "\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064")
	}
	if _cef != nil {
		return _aa.Wrap(_cef, _db, "")
	}
	if _cef = _cf.getULCorners(_fda, _gf); _cef != nil {
		return _aa.Wrap(_cef, _db, "")
	}
	_gd := len(*_gf)
	_cf.BaseIndex += _gd
	if _cef = _cf.ComponentsNumber.Add(_gd); _cef != nil {
		return _aa.Wrap(_cef, _db, "")
	}
	return nil
}
func (_agg *Classer) getULCorners(_aeb *_ed.Bitmap, _gg *_ed.Boxes) error {
	const _ac = "\u0067\u0065\u0074U\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073"
	if _aeb == nil {
		return _aa.Error(_ac, "\u006e\u0069l\u0020\u0069\u006da\u0067\u0065\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _gg == nil {
		return _aa.Error(_ac, "\u006e\u0069\u006c\u0020\u0062\u006f\u0075\u006e\u0064\u0073")
	}
	if _agg.PtaUL == nil {
		_agg.PtaUL = &_ed.Points{}
	}
	_bfb := len(*_gg)
	var (
		_daf, _ef, _cd, _ga   int
		_cdg, _fc, _bd, _aebf float32
		_bc                   error
		_fgc                  *_e.Rectangle
		_ea                   *_ed.Bitmap
		_fgb                  _e.Point
	)
	for _ba := 0; _ba < _bfb; _ba++ {
		_daf = _agg.BaseIndex + _ba
		if _cdg, _fc, _bc = _agg.CentroidPoints.GetGeometry(_daf); _bc != nil {
			return _aa.Wrap(_bc, _ac, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050o\u0069\u006e\u0074\u0073")
		}
		if _ef, _bc = _agg.ClassIDs.Get(_daf); _bc != nil {
			return _aa.Wrap(_bc, _ac, "\u0043\u006c\u0061s\u0073\u0049\u0044\u0073\u002e\u0047\u0065\u0074")
		}
		if _bd, _aebf, _bc = _agg.CentroidPointsTemplates.GetGeometry(_ef); _bc != nil {
			return _aa.Wrap(_bc, _ac, "\u0043\u0065\u006etr\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073")
		}
		_caa := _bd - _cdg
		_be := _aebf - _fc
		if _caa >= 0 {
			_cd = int(_caa + 0.5)
		} else {
			_cd = int(_caa - 0.5)
		}
		if _be >= 0 {
			_ga = int(_be + 0.5)
		} else {
			_ga = int(_be - 0.5)
		}
		if _fgc, _bc = _gg.Get(_ba); _bc != nil {
			return _aa.Wrap(_bc, _ac, "")
		}
		_efd, _baa := _fgc.Min.X, _fgc.Min.Y
		_ea, _bc = _agg.UndilatedTemplates.GetBitmap(_ef)
		if _bc != nil {
			return _aa.Wrap(_bc, _ac, "\u0055\u006e\u0064\u0069\u006c\u0061\u0074\u0065\u0064\u0054e\u006d\u0070\u006c\u0061\u0074\u0065\u0073.\u0047\u0065\u0074\u0028\u0069\u0043\u006c\u0061\u0073\u0073\u0029")
		}
		_fgb, _bc = _ec(_aeb, _efd, _baa, _cd, _ga, _ea)
		if _bc != nil {
			return _aa.Wrap(_bc, _ac, "")
		}
		_agg.PtaUL.AddPoint(float32(_efd-_cd+_fgb.X), float32(_baa-_ga+_fgb.Y))
	}
	return nil
}

type Settings struct {
	MaxCompWidth       int
	MaxCompHeight      int
	SizeHaus           int
	RankHaus           float64
	Thresh             float64
	WeightFactor       float64
	KeepClassInstances bool
	Components         _ed.Component
	Method             Method
}

const (
	MaxConnCompWidth = 350
	MaxCharCompWidth = 350
	MaxWordCompWidth = 1000
	MaxCompHeight    = 120
)

func (_eac *Classer) classifyCorrelation(_baf *_ed.Boxes, _cfg *_ed.Bitmaps, _gb int) error {
	const _bgf = "\u0063\u006c\u0061\u0073si\u0066\u0079\u0043\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e"
	if _baf == nil {
		return _aa.Error(_bgf, "\u006e\u0065\u0077\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062o\u0078\u0065\u0073\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _cfg == nil {
		return _aa.Error(_bgf, "\u006e\u0065wC\u006f\u006d\u0070o\u006e\u0065\u006e\u0074s b\u0069tm\u0061\u0070\u0020\u0061\u0072\u0072\u0061y \u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
	}
	_bgb := len(_cfg.Values)
	if _bgb == 0 {
		_d.Log.Debug("\u0063l\u0061\u0073s\u0069\u0066\u0079C\u006f\u0072\u0072\u0065\u006c\u0061\u0074i\u006f\u006e\u0020\u002d\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0070\u0069\u0078\u0061s\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079")
		return nil
	}
	var (
		_abc, _gaa *_ed.Bitmap
		_gca       error
	)
	_egc := &_ed.Bitmaps{Values: make([]*_ed.Bitmap, _bgb)}
	for _ecd, _badb := range _cfg.Values {
		_gaa, _gca = _badb.AddBorderGeneral(JbAddedPixels, JbAddedPixels, JbAddedPixels, JbAddedPixels, 0)
		if _gca != nil {
			return _aa.Wrap(_gca, _bgf, "")
		}
		_egc.Values[_ecd] = _gaa
	}
	_bda := _eac.FgTemplates
	_cec := _ed.MakePixelSumTab8()
	_agd := _ed.MakePixelCentroidTab8()
	_abg := make([]int, _bgb)
	_cfgb := make([][]int, _bgb)
	_ffee := _ed.Points(make([]_ed.Point, _bgb))
	_ccb := &_ffee
	var (
		_gcf, _aaa     int
		_dc, _fb, _bfa int
		_bfe, _ade     int
		_bgc           byte
	)
	for _dcd, _bea := range _egc.Values {
		_cfgb[_dcd] = make([]int, _bea.Height)
		_gcf = 0
		_aaa = 0
		_fb = (_bea.Height - 1) * _bea.RowStride
		_dc = 0
		for _ade = _bea.Height - 1; _ade >= 0; _ade, _fb = _ade-1, _fb-_bea.RowStride {
			_cfgb[_dcd][_ade] = _dc
			_bfa = 0
			for _bfe = 0; _bfe < _bea.RowStride; _bfe++ {
				_bgc = _bea.Data[_fb+_bfe]
				_bfa += _cec[_bgc]
				_gcf += _agd[_bgc] + _bfe*8*_cec[_bgc]
			}
			_dc += _bfa
			_aaa += _bfa * _ade
		}
		_abg[_dcd] = _dc
		if _dc > 0 {
			(*_ccb)[_dcd] = _ed.Point{X: float32(_gcf) / float32(_dc), Y: float32(_aaa) / float32(_dc)}
		} else {
			(*_ccb)[_dcd] = _ed.Point{X: float32(_bea.Width) / float32(2), Y: float32(_bea.Height) / float32(2)}
		}
	}
	if _gca = _eac.CentroidPoints.Add(_ccb); _gca != nil {
		return _aa.Wrap(_gca, _bgf, "\u0063\u0065\u006et\u0072\u006f\u0069\u0064\u0020\u0061\u0064\u0064")
	}
	var (
		_eaf, _ddb, _dae       int
		_ggc                   float64
		_cb, _eafb, _aec, _bdg float32
		_dgc, _fdd             _ed.Point
		_ebf                   bool
		_gec                   *similarTemplatesFinder
		_ecdc                  int
		_aef                   *_ed.Bitmap
		_cab                   *_e.Rectangle
		_cgb                   *_ed.Bitmaps
	)
	for _ecdc, _gaa = range _egc.Values {
		_ddb = _abg[_ecdc]
		if _cb, _eafb, _gca = _ccb.GetGeometry(_ecdc); _gca != nil {
			return _aa.Wrap(_gca, _bgf, "\u0070t\u0061\u0020\u002d\u0020\u0069")
		}
		_ebf = false
		_gcae := len(_eac.UndilatedTemplates.Values)
		_gec = _dggd(_eac, _gaa)
		for _fge := _gec.Next(); _fge > -1; {
			if _aef, _gca = _eac.UndilatedTemplates.GetBitmap(_fge); _gca != nil {
				return _aa.Wrap(_gca, _bgf, "\u0075\u006e\u0069dl\u0061\u0074\u0065\u0064\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0062\u006d\u0032")
			}
			if _dae, _gca = _bda.GetInt(_fge); _gca != nil {
				_d.Log.Trace("\u0046\u0047\u0020T\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u005b\u0069\u0063\u006ca\u0073\u0073\u005d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _gca)
			}
			if _aec, _bdg, _gca = _eac.CentroidPointsTemplates.GetGeometry(_fge); _gca != nil {
				return _aa.Wrap(_gca, _bgf, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074T\u0065\u006d\u0070\u006c\u0061\u0074e\u0073\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u00782\u002c\u0079\u0032\u0020")
			}
			if _eac.Settings.WeightFactor > 0.0 {
				if _eaf, _gca = _eac.TemplateAreas.Get(_fge); _gca != nil {
					_d.Log.Trace("\u0054\u0065\u006dp\u006c\u0061\u0074\u0065A\u0072\u0065\u0061\u0073\u005b\u0069\u0063l\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0061\u0072\u0065\u0061\u0020\u0025\u0076", _gca)
				}
				_ggc = _eac.Settings.Thresh + (1.0-_eac.Settings.Thresh)*_eac.Settings.WeightFactor*float64(_dae)/float64(_eaf)
			} else {
				_ggc = _eac.Settings.Thresh
			}
			_gbd, _fbb := _ed.CorrelationScoreThresholded(_gaa, _aef, _ddb, _dae, _dgc.X-_fdd.X, _dgc.Y-_fdd.Y, MaxDiffWidth, MaxDiffHeight, _cec, _cfgb[_ecdc], float32(_ggc))
			if _fbb != nil {
				return _aa.Wrap(_fbb, _bgf, "")
			}
			if _dge {
				var (
					_ccd, _bdgf float64
					_fa, _efce  int
				)
				_ccd, _fbb = _ed.CorrelationScore(_gaa, _aef, _ddb, _dae, _cb-_aec, _eafb-_bdg, MaxDiffWidth, MaxDiffHeight, _cec)
				if _fbb != nil {
					return _aa.Wrap(_fbb, _bgf, "d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065")
				}
				_bdgf, _fbb = _ed.CorrelationScoreSimple(_gaa, _aef, _ddb, _dae, _cb-_aec, _eafb-_bdg, MaxDiffWidth, MaxDiffHeight, _cec)
				if _fbb != nil {
					return _aa.Wrap(_fbb, _bgf, "d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065")
				}
				_fa = int(_a.Sqrt(_ccd * float64(_ddb) * float64(_dae)))
				_efce = int(_a.Sqrt(_bdgf * float64(_ddb) * float64(_dae)))
				if (_ccd >= _ggc) != (_bdgf >= _ggc) {
					return _aa.Errorf(_bgf, "\u0064\u0065\u0062\u0075\u0067\u0020\u0043\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u006f\u0072\u0065\u0020\u006d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020-\u0020\u0025d\u0028\u00250\u002e\u0034\u0066\u002c\u0020\u0025\u0076\u0029\u0020\u0076\u0073\u0020\u0025d(\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0076)\u0020\u0025\u0030\u002e\u0034\u0066", _fa, _ccd, _ccd >= _ggc, _efce, _bdgf, _bdgf >= _ggc, _ccd-_bdgf)
				}
				if _ccd >= _ggc != _gbd {
					return _aa.Errorf(_bgf, "\u0064\u0065\u0062\u0075\u0067\u0020\u0043o\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e \u0073\u0063\u006f\u0072\u0065 \u004d\u0069\u0073\u006d\u0061t\u0063\u0068 \u0062\u0065\u0074w\u0065\u0065\u006e\u0020\u0063\u006frr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020/\u0020\u0074\u0068\u0072\u0065s\u0068\u006f\u006c\u0064\u002e\u0020\u0043\u006f\u006dpa\u0072\u0069\u0073\u006f\u006e:\u0020\u0025\u0030\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0064\u0029\u0020\u003e\u003d\u0020\u00250\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u0029\u0020\u0076\u0073\u0020\u0025\u0076", _ccd, _ccd*float64(_ddb)*float64(_dae), _fa, _ggc, float32(_ggc)*float32(_ddb)*float32(_dae), _gbd)
				}
			}
			if _gbd {
				_ebf = true
				if _fbb = _eac.ClassIDs.Add(_fge); _fbb != nil {
					return _aa.Wrap(_fbb, _bgf, "\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064")
				}
				if _fbb = _eac.ComponentPageNumbers.Add(_gb); _fbb != nil {
					return _aa.Wrap(_fbb, _bgf, "\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064")
				}
				if _eac.Settings.KeepClassInstances {
					if _abc, _fbb = _cfg.GetBitmap(_ecdc); _fbb != nil {
						return _aa.Wrap(_fbb, _bgf, "\u004b\u0065\u0065\u0070Cl\u0061\u0073\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u002d \u0069")
					}
					if _cgb, _fbb = _eac.ClassInstances.GetBitmaps(_fge); _fbb != nil {
						return _aa.Wrap(_fbb, _bgf, "K\u0065\u0065\u0070\u0043\u006c\u0061s\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065s\u0020\u002d\u0020i\u0043l\u0061\u0073\u0073")
					}
					_cgb.AddBitmap(_abc)
					if _cab, _fbb = _baf.Get(_ecdc); _fbb != nil {
						return _aa.Wrap(_fbb, _bgf, "\u004be\u0065p\u0043\u006c\u0061\u0073\u0073I\u006e\u0073t\u0061\u006e\u0063\u0065\u0073")
					}
					_cgb.AddBox(_cab)
				}
				break
			}
		}
		if !_ebf {
			if _gca = _eac.ClassIDs.Add(_gcae); _gca != nil {
				return _aa.Wrap(_gca, _bgf, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			if _gca = _eac.ComponentPageNumbers.Add(_gb); _gca != nil {
				return _aa.Wrap(_gca, _bgf, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_cgb = &_ed.Bitmaps{}
			if _abc, _gca = _cfg.GetBitmap(_ecdc); _gca != nil {
				return _aa.Wrap(_gca, _bgf, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_cgb.AddBitmap(_abc)
			_fdf, _def := _abc.Width, _abc.Height
			_bge := uint64(_def) * uint64(_fdf)
			_eac.TemplatesSize.Add(_bge, _gcae)
			if _cab, _gca = _baf.Get(_ecdc); _gca != nil {
				return _aa.Wrap(_gca, _bgf, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_cgb.AddBox(_cab)
			_eac.ClassInstances.AddBitmaps(_cgb)
			_eac.CentroidPointsTemplates.AddPoint(_cb, _eafb)
			_eac.FgTemplates.AddInt(_ddb)
			_eac.UndilatedTemplates.AddBitmap(_abc)
			_eaf = (_gaa.Width - 2*JbAddedPixels) * (_gaa.Height - 2*JbAddedPixels)
			if _gca = _eac.TemplateAreas.Add(_eaf); _gca != nil {
				return _aa.Wrap(_gca, _bgf, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
		}
	}
	_eac.NumberOfClasses = len(_eac.UndilatedTemplates.Values)
	return nil
}
func Init(settings Settings) (*Classer, error) {
	const _f = "\u0063\u006c\u0061s\u0073\u0065\u0072\u002e\u0049\u006e\u0069\u0074"
	_b := &Classer{Settings: settings, Widths: map[int]int{}, Heights: map[int]int{}, TemplatesSize: _ce.IntsMap{}, TemplateAreas: &_ce.IntSlice{}, ComponentPageNumbers: &_ce.IntSlice{}, ClassIDs: &_ce.IntSlice{}, ComponentsNumber: &_ce.IntSlice{}, CentroidPoints: &_ed.Points{}, CentroidPointsTemplates: &_ed.Points{}, UndilatedTemplates: &_ed.Bitmaps{}, DilatedTemplates: &_ed.Bitmaps{}, ClassInstances: &_ed.BitmapsArray{}, FgTemplates: &_ce.NumSlice{}}
	if _g := _b.Settings.Validate(); _g != nil {
		return nil, _aa.Wrap(_g, _f, "")
	}
	return _b, nil
}
func (_gaad *Classer) classifyRankHouseNonOne(_ace *_ed.Boxes, _aaf, _gcef, _bef *_ed.Bitmaps, _cfb *_ed.Points, _geg *_ce.NumSlice, _dfa int) (_cce error) {
	const _edc = "\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065"
	var (
		_dgg, _eddd, _ecc, _ggf        float32
		_bffg, _dad, _fbf              int
		_eab, _ggb, _cabe, _ggbd, _efa *_ed.Bitmap
		_eggd, _fba                    bool
	)
	_dbgg := _ed.MakePixelSumTab8()
	for _aag := 0; _aag < len(_aaf.Values); _aag++ {
		if _ggb, _cce = _gcef.GetBitmap(_aag); _cce != nil {
			return _aa.Wrap(_cce, _edc, "b\u006d\u0073\u0031\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
		}
		if _bffg, _cce = _geg.GetInt(_aag); _cce != nil {
			_d.Log.Trace("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0046\u0047T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073 \u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _aag, _cce)
		}
		if _cabe, _cce = _bef.GetBitmap(_aag); _cce != nil {
			return _aa.Wrap(_cce, _edc, "b\u006d\u0073\u0032\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
		}
		if _dgg, _eddd, _cce = _cfb.GetGeometry(_aag); _cce != nil {
			return _aa.Wrapf(_cce, _edc, "\u0070t\u0061[\u0069\u005d\u002e\u0047\u0065\u006f\u006d\u0065\u0074\u0072\u0079")
		}
		_cefg := len(_gaad.UndilatedTemplates.Values)
		_eggd = false
		_aebg := _dggd(_gaad, _ggb)
		for _fbf = _aebg.Next(); _fbf > -1; {
			if _ggbd, _cce = _gaad.UndilatedTemplates.GetBitmap(_fbf); _cce != nil {
				return _aa.Wrap(_cce, _edc, "\u0070\u0069\u0078\u0061\u0074\u002e\u005b\u0069\u0043l\u0061\u0073\u0073\u005d")
			}
			if _dad, _cce = _gaad.FgTemplates.GetInt(_fbf); _cce != nil {
				_d.Log.Trace("\u0047\u0065\u0074\u0074\u0069\u006eg\u0020\u0046\u0047\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u005b\u0025d\u005d\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _fbf, _cce)
			}
			if _efa, _cce = _gaad.DilatedTemplates.GetBitmap(_fbf); _cce != nil {
				return _aa.Wrap(_cce, _edc, "\u0070\u0069\u0078\u0061\u0074\u0064\u005b\u0069\u0043l\u0061\u0073\u0073\u005d")
			}
			if _ecc, _ggf, _cce = _gaad.CentroidPointsTemplates.GetGeometry(_fbf); _cce != nil {
				return _aa.Wrap(_cce, _edc, "\u0043\u0065\u006et\u0072\u006f\u0069\u0064P\u006f\u0069\u006e\u0074\u0073\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u005b\u0069\u0043\u006c\u0061\u0073\u0073\u005d")
			}
			_fba, _cce = _ed.RankHausTest(_ggb, _cabe, _ggbd, _efa, _dgg-_ecc, _eddd-_ggf, MaxDiffWidth, MaxDiffHeight, _bffg, _dad, float32(_gaad.Settings.RankHaus), _dbgg)
			if _cce != nil {
				return _aa.Wrap(_cce, _edc, "")
			}
			if _fba {
				_eggd = true
				if _cce = _gaad.ClassIDs.Add(_fbf); _cce != nil {
					return _aa.Wrap(_cce, _edc, "")
				}
				if _cce = _gaad.ComponentPageNumbers.Add(_dfa); _cce != nil {
					return _aa.Wrap(_cce, _edc, "")
				}
				if _gaad.Settings.KeepClassInstances {
					_gge, _fged := _gaad.ClassInstances.GetBitmaps(_fbf)
					if _fged != nil {
						return _aa.Wrap(_fged, _edc, "\u0063\u002e\u0050\u0069\u0078\u0061\u0061\u002e\u0047\u0065\u0074B\u0069\u0074\u006d\u0061\u0070\u0073\u0028\u0069\u0043\u006ca\u0073\u0073\u0029")
					}
					if _eab, _fged = _aaf.GetBitmap(_aag); _fged != nil {
						return _aa.Wrap(_fged, _edc, "\u0070i\u0078\u0061\u005b\u0069\u005d")
					}
					_gge.Values = append(_gge.Values, _eab)
					_aedc, _fged := _ace.Get(_aag)
					if _fged != nil {
						return _aa.Wrap(_fged, _edc, "b\u006f\u0078\u0061\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
					}
					_gge.Boxes = append(_gge.Boxes, _aedc)
				}
				break
			}
		}
		if !_eggd {
			if _cce = _gaad.ClassIDs.Add(_cefg); _cce != nil {
				return _aa.Wrap(_cce, _edc, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			if _cce = _gaad.ComponentPageNumbers.Add(_dfa); _cce != nil {
				return _aa.Wrap(_cce, _edc, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_fad := &_ed.Bitmaps{}
			_eab = _aaf.Values[_aag]
			_fad.AddBitmap(_eab)
			_aae, _gdg := _eab.Width, _eab.Height
			_gaad.TemplatesSize.Add(uint64(_aae)*uint64(_gdg), _cefg)
			_fbdg, _acb := _ace.Get(_aag)
			if _acb != nil {
				return _aa.Wrap(_acb, _edc, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_fad.AddBox(_fbdg)
			_gaad.ClassInstances.AddBitmaps(_fad)
			_gaad.CentroidPointsTemplates.AddPoint(_dgg, _eddd)
			_gaad.UndilatedTemplates.AddBitmap(_ggb)
			_gaad.DilatedTemplates.AddBitmap(_cabe)
			_gaad.FgTemplates.AddInt(_bffg)
		}
	}
	_gaad.NumberOfClasses = len(_gaad.UndilatedTemplates.Values)
	return nil
}
func (_acd *Classer) classifyRankHaus(_gef *_ed.Boxes, _gce *_ed.Bitmaps, _cbe int) error {
	const _gfc = "\u0063\u006ca\u0073\u0073\u0069f\u0079\u0052\u0061\u006e\u006b\u0048\u0061\u0075\u0073"
	if _gef == nil {
		return _aa.Error(_gfc, "\u0062\u006fx\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _gce == nil {
		return _aa.Error(_gfc, "\u0070\u0069x\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	_egag := len(_gce.Values)
	if _egag == 0 {
		return _aa.Error(_gfc, "e\u006dp\u0074\u0079\u0020\u006e\u0065\u0077\u0020\u0063o\u006d\u0070\u006f\u006een\u0074\u0073")
	}
	_dea := _gce.CountPixels()
	_dgf := _acd.Settings.SizeHaus
	_beb := _ed.SelCreateBrick(_dgf, _dgf, _dgf/2, _dgf/2, _ed.SelHit)
	_fde := &_ed.Bitmaps{Values: make([]*_ed.Bitmap, _egag)}
	_fbd := &_ed.Bitmaps{Values: make([]*_ed.Bitmap, _egag)}
	var (
		_afe, _abe, _bdaa *_ed.Bitmap
		_adf              error
	)
	for _cacd := 0; _cacd < _egag; _cacd++ {
		_afe, _adf = _gce.GetBitmap(_cacd)
		if _adf != nil {
			return _aa.Wrap(_adf, _gfc, "")
		}
		_abe, _adf = _afe.AddBorderGeneral(JbAddedPixels, JbAddedPixels, JbAddedPixels, JbAddedPixels, 0)
		if _adf != nil {
			return _aa.Wrap(_adf, _gfc, "")
		}
		_bdaa, _adf = _ed.Dilate(nil, _abe, _beb)
		if _adf != nil {
			return _aa.Wrap(_adf, _gfc, "")
		}
		_fde.Values[_egag] = _abe
		_fbd.Values[_egag] = _bdaa
	}
	_bb, _adf := _ed.Centroids(_fde.Values)
	if _adf != nil {
		return _aa.Wrap(_adf, _gfc, "")
	}
	if _adf = _bb.Add(_acd.CentroidPoints); _adf != nil {
		_d.Log.Trace("\u004e\u006f\u0020\u0063en\u0074\u0072\u006f\u0069\u0064\u0073\u0020\u0074\u006f\u0020\u0061\u0064\u0064")
	}
	if _acd.Settings.RankHaus == 1.0 {
		_adf = _acd.classifyRankHouseOne(_gef, _gce, _fde, _fbd, _bb, _cbe)
	} else {
		_adf = _acd.classifyRankHouseNonOne(_gef, _gce, _fde, _fbd, _bb, _dea, _cbe)
	}
	if _adf != nil {
		return _aa.Wrap(_adf, _gfc, "")
	}
	return nil
}

type Classer struct {
	BaseIndex               int
	Settings                Settings
	ComponentsNumber        *_ce.IntSlice
	TemplateAreas           *_ce.IntSlice
	Widths                  map[int]int
	Heights                 map[int]int
	NumberOfClasses         int
	ClassInstances          *_ed.BitmapsArray
	UndilatedTemplates      *_ed.Bitmaps
	DilatedTemplates        *_ed.Bitmaps
	TemplatesSize           _ce.IntsMap
	FgTemplates             *_ce.NumSlice
	CentroidPoints          *_ed.Points
	CentroidPointsTemplates *_ed.Points
	ClassIDs                *_ce.IntSlice
	ComponentPageNumbers    *_ce.IntSlice
	PtaUL                   *_ed.Points
	PtaLL                   *_ed.Points
}

func (_ad *Classer) verifyMethod(_bdb Method) error {
	if _bdb != RankHaus && _bdb != Correlation {
		return _aa.Error("\u0076\u0065\u0072i\u0066\u0079\u004d\u0065\u0074\u0068\u006f\u0064", "\u0069\u006e\u0076\u0061li\u0064\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072\u0020\u006d\u0065\u0074\u0068o\u0064")
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

func (_fd *Classer) ComputeLLCorners() (_ee error) {
	const _da = "\u0043l\u0061\u0073\u0073\u0065\u0072\u002e\u0043\u006f\u006d\u0070\u0075t\u0065\u004c\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073"
	if _fd.PtaUL == nil {
		return _aa.Error(_da, "\u0055\u004c\u0020\u0043or\u006e\u0065\u0072\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	_ca := len(*_fd.PtaUL)
	_fd.PtaLL = &_ed.Points{}
	var (
		_ab, _ge float32
		_dg, _df int
		_bfc     *_ed.Bitmap
	)
	for _eeg := 0; _eeg < _ca; _eeg++ {
		_ab, _ge, _ee = _fd.PtaUL.GetGeometry(_eeg)
		if _ee != nil {
			_d.Log.Debug("\u0047e\u0074\u0074\u0069\u006e\u0067\u0020\u0050\u0074\u0061\u0055\u004c \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ee)
			return _aa.Wrap(_ee, _da, "\u0050\u0074\u0061\u0055\u004c\u0020\u0047\u0065\u006fm\u0065\u0074\u0072\u0079")
		}
		_dg, _ee = _fd.ClassIDs.Get(_eeg)
		if _ee != nil {
			_d.Log.Debug("\u0047\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0043\u006c\u0061s\u0073\u0049\u0044\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _ee)
			return _aa.Wrap(_ee, _da, "\u0043l\u0061\u0073\u0073\u0049\u0044")
		}
		_bfc, _ee = _fd.UndilatedTemplates.GetBitmap(_dg)
		if _ee != nil {
			_d.Log.Debug("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0055\u006ed\u0069\u006c\u0061\u0074\u0065\u0064\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _ee)
			return _aa.Wrap(_ee, _da, "\u0055\u006e\u0064\u0069la\u0074\u0065\u0064\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073")
		}
		_df = _bfc.Height
		_fd.PtaLL.AddPoint(_ab, _ge+float32(_df))
	}
	return nil
}
func DefaultSettings() Settings {
	_fdfg := &Settings{}
	_fdfg.SetDefault()
	return *_fdfg
}

var _dge bool
var TwoByTwoWalk = []int{0, 0, 0, 1, -1, 0, 0, -1, 1, 0, -1, 1, 1, 1, -1, -1, 1, -1, 0, -2, 2, 0, 0, 2, -2, 0, -1, -2, 1, -2, 2, -1, 2, 1, 1, 2, -1, 2, -2, 1, -2, -1, -2, -2, 2, -2, 2, 2, -2, 2}

func (_edd *Classer) AddPage(inputPage *_ed.Bitmap, pageNumber int, method Method) (_ff error) {
	const _ag = "\u0043l\u0061s\u0073\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065"
	_edd.Widths[pageNumber] = inputPage.Width
	_edd.Heights[pageNumber] = inputPage.Height
	if _ff = _edd.verifyMethod(method); _ff != nil {
		return _aa.Wrap(_ff, _ag, "")
	}
	_bf, _ae, _ff := inputPage.GetComponents(_edd.Settings.Components, _edd.Settings.MaxCompWidth, _edd.Settings.MaxCompHeight)
	if _ff != nil {
		return _aa.Wrap(_ff, _ag, "")
	}
	_d.Log.Debug("\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _bf)
	if _ff = _edd.addPageComponents(inputPage, _ae, _bf, pageNumber, method); _ff != nil {
		return _aa.Wrap(_ff, _ag, "")
	}
	return nil
}

type Method int

func (_bfg *Classer) classifyRankHouseOne(_dca *_ed.Boxes, _bcd, _cff, _gaf *_ed.Bitmaps, _bbg *_ed.Points, _bdbf int) (_egd error) {
	const _dff = "\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065"
	var (
		_bff, _dfb, _gcb, _ccf         float32
		_bfgg                          int
		_fcb, _agb, _adb, _ddbg, _gafd *_ed.Bitmap
		_ccg, _aad                     bool
	)
	for _bdd := 0; _bdd < len(_bcd.Values); _bdd++ {
		_agb = _cff.Values[_bdd]
		_adb = _gaf.Values[_bdd]
		_bff, _dfb, _egd = _bbg.GetGeometry(_bdd)
		if _egd != nil {
			return _aa.Wrapf(_egd, _dff, "\u0066\u0069\u0072\u0073\u0074\u0020\u0067\u0065\u006fm\u0065\u0074\u0072\u0079")
		}
		_egg := len(_bfg.UndilatedTemplates.Values)
		_ccg = false
		_gea := _dggd(_bfg, _agb)
		for _bfgg = _gea.Next(); _bfgg > -1; {
			_ddbg, _egd = _bfg.UndilatedTemplates.GetBitmap(_bfgg)
			if _egd != nil {
				return _aa.Wrap(_egd, _dff, "\u0062\u006d\u0033")
			}
			_gafd, _egd = _bfg.DilatedTemplates.GetBitmap(_bfgg)
			if _egd != nil {
				return _aa.Wrap(_egd, _dff, "\u0062\u006d\u0034")
			}
			_gcb, _ccf, _egd = _bfg.CentroidPointsTemplates.GetGeometry(_bfgg)
			if _egd != nil {
				return _aa.Wrap(_egd, _dff, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0073")
			}
			_aad, _egd = _ed.HausTest(_agb, _adb, _ddbg, _gafd, _bff-_gcb, _dfb-_ccf, MaxDiffWidth, MaxDiffHeight)
			if _egd != nil {
				return _aa.Wrap(_egd, _dff, "")
			}
			if _aad {
				_ccg = true
				if _egd = _bfg.ClassIDs.Add(_bfgg); _egd != nil {
					return _aa.Wrap(_egd, _dff, "")
				}
				if _egd = _bfg.ComponentPageNumbers.Add(_bdbf); _egd != nil {
					return _aa.Wrap(_egd, _dff, "")
				}
				if _bfg.Settings.KeepClassInstances {
					_caf, _gbdd := _bfg.ClassInstances.GetBitmaps(_bfgg)
					if _gbdd != nil {
						return _aa.Wrap(_gbdd, _dff, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_fcb, _gbdd = _bcd.GetBitmap(_bdd)
					if _gbdd != nil {
						return _aa.Wrap(_gbdd, _dff, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_caf.AddBitmap(_fcb)
					_gbf, _gbdd := _dca.Get(_bdd)
					if _gbdd != nil {
						return _aa.Wrap(_gbdd, _dff, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_caf.AddBox(_gbf)
				}
				break
			}
		}
		if !_ccg {
			if _egd = _bfg.ClassIDs.Add(_egg); _egd != nil {
				return _aa.Wrap(_egd, _dff, "")
			}
			if _egd = _bfg.ComponentPageNumbers.Add(_bdbf); _egd != nil {
				return _aa.Wrap(_egd, _dff, "")
			}
			_fbbe := &_ed.Bitmaps{}
			_fcb, _egd = _bcd.GetBitmap(_bdd)
			if _egd != nil {
				return _aa.Wrap(_egd, _dff, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_fbbe.Values = append(_fbbe.Values, _fcb)
			_bgce, _ada := _fcb.Width, _fcb.Height
			_bfg.TemplatesSize.Add(uint64(_ada)*uint64(_bgce), _egg)
			_ebc, _afec := _dca.Get(_bdd)
			if _afec != nil {
				return _aa.Wrap(_afec, _dff, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_fbbe.AddBox(_ebc)
			_bfg.ClassInstances.AddBitmaps(_fbbe)
			_bfg.CentroidPointsTemplates.AddPoint(_bff, _dfb)
			_bfg.UndilatedTemplates.AddBitmap(_agb)
			_bfg.DilatedTemplates.AddBitmap(_adb)
		}
	}
	return nil
}
func _dggd(_cdc *Classer, _eccg *_ed.Bitmap) *similarTemplatesFinder {
	return &similarTemplatesFinder{Width: _eccg.Width, Height: _eccg.Height, Classer: _cdc}
}
func (_ccfc Settings) Validate() error {
	const _gfg = "\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0073\u002e\u0056\u0061\u006ci\u0064\u0061\u0074\u0065"
	if _ccfc.Thresh < 0.4 || _ccfc.Thresh > 0.98 {
		return _aa.Error(_gfg, "\u006a\u0062i\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0034\u0020\u002d\u0020\u0030\u002e\u0039\u0038\u005d")
	}
	if _ccfc.WeightFactor < 0.0 || _ccfc.WeightFactor > 1.0 {
		return _aa.Error(_gfg, "\u006a\u0062i\u0067\u0032\u0020\u0065\u006ec\u006f\u0064\u0065\u0072\u0020w\u0065\u0069\u0067\u0068\u0074\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0030\u0020\u002d\u0020\u0031\u002e\u0030\u005d")
	}
	if _ccfc.RankHaus < 0.5 || _ccfc.RankHaus > 1.0 {
		return _aa.Error(_gfg, "\u006a\u0062\u0069\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065 \u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 [\u0030\u002e\u0035\u0020\u002d\u0020\u0031\u002e\u0030\u005d")
	}
	if _ccfc.SizeHaus < 1 || _ccfc.SizeHaus > 10 {
		return _aa.Error(_gfg, "\u006a\u0062\u0069\u0067\u0032 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0073\u0069\u007a\u0065\u0020h\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0031\u0020\u002d\u0020\u0031\u0030]")
	}
	switch _ccfc.Components {
	case _ed.ComponentConn, _ed.ComponentCharacters, _ed.ComponentWords:
	default:
		return _aa.Error(_gfg, "\u0069n\u0076\u0061\u006c\u0069d\u0020\u0063\u006c\u0061\u0073s\u0065r\u0020c\u006f\u006d\u0070\u006f\u006e\u0065\u006et")
	}
	return nil
}

const (
	MaxDiffWidth  = 2
	MaxDiffHeight = 2
)

func (_bbc *similarTemplatesFinder) Next() int {
	var (
		_bae, _bgg, _dggb, _fce int
		_eaff                   bool
		_edg                    *_ed.Bitmap
		_gbc                    error
	)
	for {
		if _bbc.Index >= 25 {
			return -1
		}
		_bgg = _bbc.Width + TwoByTwoWalk[2*_bbc.Index]
		_bae = _bbc.Height + TwoByTwoWalk[2*_bbc.Index+1]
		if _bae < 1 || _bgg < 1 {
			_bbc.Index++
			continue
		}
		if len(_bbc.CurrentNumbers) == 0 {
			_bbc.CurrentNumbers, _eaff = _bbc.Classer.TemplatesSize.GetSlice(uint64(_bgg) * uint64(_bae))
			if !_eaff {
				_bbc.Index++
				continue
			}
			_bbc.N = 0
		}
		_dggb = len(_bbc.CurrentNumbers)
		for ; _bbc.N < _dggb; _bbc.N++ {
			_fce = _bbc.CurrentNumbers[_bbc.N]
			_edg, _gbc = _bbc.Classer.DilatedTemplates.GetBitmap(_fce)
			if _gbc != nil {
				_d.Log.Debug("\u0046\u0069\u006e\u0064\u004e\u0065\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0074\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020")
				return 0
			}
			if _edg.Width-2*JbAddedPixels == _bgg && _edg.Height-2*JbAddedPixels == _bae {
				return _fce
			}
		}
		_bbc.Index++
		_bbc.CurrentNumbers = nil
	}
}

const JbAddedPixels = 6
