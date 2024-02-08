package classer

import (
	_d "image"
	_ca "math"

	_b "bitbucket.org/shenghui0779/gopdf/common"
	_f "bitbucket.org/shenghui0779/gopdf/internal/jbig2/basic"
	_cf "bitbucket.org/shenghui0779/gopdf/internal/jbig2/bitmap"
	_e "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
)

func (_ggdd *Classer) classifyRankHaus(_dgg *_cf.Boxes, _eeb *_cf.Bitmaps, _egc int) error {
	const _afa = "\u0063\u006ca\u0073\u0073\u0069f\u0079\u0052\u0061\u006e\u006b\u0048\u0061\u0075\u0073"
	if _dgg == nil {
		return _e.Error(_afa, "\u0062\u006fx\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _eeb == nil {
		return _e.Error(_afa, "\u0070\u0069x\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	_gbd := len(_eeb.Values)
	if _gbd == 0 {
		return _e.Error(_afa, "e\u006dp\u0074\u0079\u0020\u006e\u0065\u0077\u0020\u0063o\u006d\u0070\u006f\u006een\u0074\u0073")
	}
	_ggf := _eeb.CountPixels()
	_eee := _ggdd.Settings.SizeHaus
	_ade := _cf.SelCreateBrick(_eee, _eee, _eee/2, _eee/2, _cf.SelHit)
	_fcgfa := &_cf.Bitmaps{Values: make([]*_cf.Bitmap, _gbd)}
	_ead := &_cf.Bitmaps{Values: make([]*_cf.Bitmap, _gbd)}
	var (
		_cbe, _dgde, _gef *_cf.Bitmap
		_ace              error
	)
	for _dfc := 0; _dfc < _gbd; _dfc++ {
		_cbe, _ace = _eeb.GetBitmap(_dfc)
		if _ace != nil {
			return _e.Wrap(_ace, _afa, "")
		}
		_dgde, _ace = _cbe.AddBorderGeneral(JbAddedPixels, JbAddedPixels, JbAddedPixels, JbAddedPixels, 0)
		if _ace != nil {
			return _e.Wrap(_ace, _afa, "")
		}
		_gef, _ace = _cf.Dilate(nil, _dgde, _ade)
		if _ace != nil {
			return _e.Wrap(_ace, _afa, "")
		}
		_fcgfa.Values[_gbd] = _dgde
		_ead.Values[_gbd] = _gef
	}
	_agd, _ace := _cf.Centroids(_fcgfa.Values)
	if _ace != nil {
		return _e.Wrap(_ace, _afa, "")
	}
	if _ace = _agd.Add(_ggdd.CentroidPoints); _ace != nil {
		_b.Log.Trace("\u004e\u006f\u0020\u0063en\u0074\u0072\u006f\u0069\u0064\u0073\u0020\u0074\u006f\u0020\u0061\u0064\u0064")
	}
	if _ggdd.Settings.RankHaus == 1.0 {
		_ace = _ggdd.classifyRankHouseOne(_dgg, _eeb, _fcgfa, _ead, _agd, _egc)
	} else {
		_ace = _ggdd.classifyRankHouseNonOne(_dgg, _eeb, _fcgfa, _ead, _agd, _ggf, _egc)
	}
	if _ace != nil {
		return _e.Wrap(_ace, _afa, "")
	}
	return nil
}

const (
	MaxDiffWidth  = 2
	MaxDiffHeight = 2
)

func (_fe *Classer) classifyCorrelation(_adc *_cf.Boxes, _ffad *_cf.Bitmaps, _fed int) error {
	const _ffg = "\u0063\u006c\u0061\u0073si\u0066\u0079\u0043\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e"
	if _adc == nil {
		return _e.Error(_ffg, "\u006e\u0065\u0077\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062o\u0078\u0065\u0073\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _ffad == nil {
		return _e.Error(_ffg, "\u006e\u0065wC\u006f\u006d\u0070o\u006e\u0065\u006e\u0074s b\u0069tm\u0061\u0070\u0020\u0061\u0072\u0072\u0061y \u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
	}
	_gbb := len(_ffad.Values)
	if _gbb == 0 {
		_b.Log.Debug("\u0063l\u0061\u0073s\u0069\u0066\u0079C\u006f\u0072\u0072\u0065\u006c\u0061\u0074i\u006f\u006e\u0020\u002d\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0070\u0069\u0078\u0061s\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079")
		return nil
	}
	var (
		_dfa, _gab *_cf.Bitmap
		_efd       error
	)
	_acad := &_cf.Bitmaps{Values: make([]*_cf.Bitmap, _gbb)}
	for _fga, _bfg := range _ffad.Values {
		_gab, _efd = _bfg.AddBorderGeneral(JbAddedPixels, JbAddedPixels, JbAddedPixels, JbAddedPixels, 0)
		if _efd != nil {
			return _e.Wrap(_efd, _ffg, "")
		}
		_acad.Values[_fga] = _gab
	}
	_fcd := _fe.FgTemplates
	_cba := _cf.MakePixelSumTab8()
	_fcb := _cf.MakePixelCentroidTab8()
	_fcg := make([]int, _gbb)
	_ccg := make([][]int, _gbb)
	_ece := _cf.Points(make([]_cf.Point, _gbb))
	_babc := &_ece
	var (
		_ae, _degf      int
		_dca, _fee, _ag int
		_defc, _efb     int
		_ffd            byte
	)
	for _cga, _cfa := range _acad.Values {
		_ccg[_cga] = make([]int, _cfa.Height)
		_ae = 0
		_degf = 0
		_fee = (_cfa.Height - 1) * _cfa.RowStride
		_dca = 0
		for _efb = _cfa.Height - 1; _efb >= 0; _efb, _fee = _efb-1, _fee-_cfa.RowStride {
			_ccg[_cga][_efb] = _dca
			_ag = 0
			for _defc = 0; _defc < _cfa.RowStride; _defc++ {
				_ffd = _cfa.Data[_fee+_defc]
				_ag += _cba[_ffd]
				_ae += _fcb[_ffd] + _defc*8*_cba[_ffd]
			}
			_dca += _ag
			_degf += _ag * _efb
		}
		_fcg[_cga] = _dca
		if _dca > 0 {
			(*_babc)[_cga] = _cf.Point{X: float32(_ae) / float32(_dca), Y: float32(_degf) / float32(_dca)}
		} else {
			(*_babc)[_cga] = _cf.Point{X: float32(_cfa.Width) / float32(2), Y: float32(_cfa.Height) / float32(2)}
		}
	}
	if _efd = _fe.CentroidPoints.Add(_babc); _efd != nil {
		return _e.Wrap(_efd, _ffg, "\u0063\u0065\u006et\u0072\u006f\u0069\u0064\u0020\u0061\u0064\u0064")
	}
	var (
		_db, _fcgf, _bdge      int
		_gba                   float64
		_bdbd, _gga, _gd, _gdc float32
		_fd, _bcc              _cf.Point
		_fcc                   bool
		_dgdf                  *similarTemplatesFinder
		_daa                   int
		_dega                  *_cf.Bitmap
		_dgb                   *_d.Rectangle
		_fgc                   *_cf.Bitmaps
	)
	for _daa, _gab = range _acad.Values {
		_fcgf = _fcg[_daa]
		if _bdbd, _gga, _efd = _babc.GetGeometry(_daa); _efd != nil {
			return _e.Wrap(_efd, _ffg, "\u0070t\u0061\u0020\u002d\u0020\u0069")
		}
		_fcc = false
		_ecc := len(_fe.UndilatedTemplates.Values)
		_dgdf = _caeg(_fe, _gab)
		for _dcg := _dgdf.Next(); _dcg > -1; {
			if _dega, _efd = _fe.UndilatedTemplates.GetBitmap(_dcg); _efd != nil {
				return _e.Wrap(_efd, _ffg, "\u0075\u006e\u0069dl\u0061\u0074\u0065\u0064\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0062\u006d\u0032")
			}
			if _bdge, _efd = _fcd.GetInt(_dcg); _efd != nil {
				_b.Log.Trace("\u0046\u0047\u0020T\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u005b\u0069\u0063\u006ca\u0073\u0073\u005d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _efd)
			}
			if _gd, _gdc, _efd = _fe.CentroidPointsTemplates.GetGeometry(_dcg); _efd != nil {
				return _e.Wrap(_efd, _ffg, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074T\u0065\u006d\u0070\u006c\u0061\u0074e\u0073\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u00782\u002c\u0079\u0032\u0020")
			}
			if _fe.Settings.WeightFactor > 0.0 {
				if _db, _efd = _fe.TemplateAreas.Get(_dcg); _efd != nil {
					_b.Log.Trace("\u0054\u0065\u006dp\u006c\u0061\u0074\u0065A\u0072\u0065\u0061\u0073\u005b\u0069\u0063l\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0061\u0072\u0065\u0061\u0020\u0025\u0076", _efd)
				}
				_gba = _fe.Settings.Thresh + (1.0-_fe.Settings.Thresh)*_fe.Settings.WeightFactor*float64(_bdge)/float64(_db)
			} else {
				_gba = _fe.Settings.Thresh
			}
			_fdc, _ceag := _cf.CorrelationScoreThresholded(_gab, _dega, _fcgf, _bdge, _fd.X-_bcc.X, _fd.Y-_bcc.Y, MaxDiffWidth, MaxDiffHeight, _cba, _ccg[_daa], float32(_gba))
			if _ceag != nil {
				return _e.Wrap(_ceag, _ffg, "")
			}
			if _gc {
				var (
					_eef, _fcgd float64
					_deb, _daf  int
				)
				_eef, _ceag = _cf.CorrelationScore(_gab, _dega, _fcgf, _bdge, _bdbd-_gd, _gga-_gdc, MaxDiffWidth, MaxDiffHeight, _cba)
				if _ceag != nil {
					return _e.Wrap(_ceag, _ffg, "d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065")
				}
				_fcgd, _ceag = _cf.CorrelationScoreSimple(_gab, _dega, _fcgf, _bdge, _bdbd-_gd, _gga-_gdc, MaxDiffWidth, MaxDiffHeight, _cba)
				if _ceag != nil {
					return _e.Wrap(_ceag, _ffg, "d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065")
				}
				_deb = int(_ca.Sqrt(_eef * float64(_fcgf) * float64(_bdge)))
				_daf = int(_ca.Sqrt(_fcgd * float64(_fcgf) * float64(_bdge)))
				if (_eef >= _gba) != (_fcgd >= _gba) {
					return _e.Errorf(_ffg, "\u0064\u0065\u0062\u0075\u0067\u0020\u0043\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u006f\u0072\u0065\u0020\u006d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020-\u0020\u0025d\u0028\u00250\u002e\u0034\u0066\u002c\u0020\u0025\u0076\u0029\u0020\u0076\u0073\u0020\u0025d(\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0076)\u0020\u0025\u0030\u002e\u0034\u0066", _deb, _eef, _eef >= _gba, _daf, _fcgd, _fcgd >= _gba, _eef-_fcgd)
				}
				if _eef >= _gba != _fdc {
					return _e.Errorf(_ffg, "\u0064\u0065\u0062\u0075\u0067\u0020\u0043o\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e \u0073\u0063\u006f\u0072\u0065 \u004d\u0069\u0073\u006d\u0061t\u0063\u0068 \u0062\u0065\u0074w\u0065\u0065\u006e\u0020\u0063\u006frr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020/\u0020\u0074\u0068\u0072\u0065s\u0068\u006f\u006c\u0064\u002e\u0020\u0043\u006f\u006dpa\u0072\u0069\u0073\u006f\u006e:\u0020\u0025\u0030\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0064\u0029\u0020\u003e\u003d\u0020\u00250\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u0029\u0020\u0076\u0073\u0020\u0025\u0076", _eef, _eef*float64(_fcgf)*float64(_bdge), _deb, _gba, float32(_gba)*float32(_fcgf)*float32(_bdge), _fdc)
				}
			}
			if _fdc {
				_fcc = true
				if _ceag = _fe.ClassIDs.Add(_dcg); _ceag != nil {
					return _e.Wrap(_ceag, _ffg, "\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064")
				}
				if _ceag = _fe.ComponentPageNumbers.Add(_fed); _ceag != nil {
					return _e.Wrap(_ceag, _ffg, "\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064")
				}
				if _fe.Settings.KeepClassInstances {
					if _dfa, _ceag = _ffad.GetBitmap(_daa); _ceag != nil {
						return _e.Wrap(_ceag, _ffg, "\u004b\u0065\u0065\u0070Cl\u0061\u0073\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u002d \u0069")
					}
					if _fgc, _ceag = _fe.ClassInstances.GetBitmaps(_dcg); _ceag != nil {
						return _e.Wrap(_ceag, _ffg, "K\u0065\u0065\u0070\u0043\u006c\u0061s\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065s\u0020\u002d\u0020i\u0043l\u0061\u0073\u0073")
					}
					_fgc.AddBitmap(_dfa)
					if _dgb, _ceag = _adc.Get(_daa); _ceag != nil {
						return _e.Wrap(_ceag, _ffg, "\u004be\u0065p\u0043\u006c\u0061\u0073\u0073I\u006e\u0073t\u0061\u006e\u0063\u0065\u0073")
					}
					_fgc.AddBox(_dgb)
				}
				break
			}
		}
		if !_fcc {
			if _efd = _fe.ClassIDs.Add(_ecc); _efd != nil {
				return _e.Wrap(_efd, _ffg, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			if _efd = _fe.ComponentPageNumbers.Add(_fed); _efd != nil {
				return _e.Wrap(_efd, _ffg, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_fgc = &_cf.Bitmaps{}
			if _dfa, _efd = _ffad.GetBitmap(_daa); _efd != nil {
				return _e.Wrap(_efd, _ffg, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_fgc.AddBitmap(_dfa)
			_aadd, _dege := _dfa.Width, _dfa.Height
			_be := uint64(_dege) * uint64(_aadd)
			_fe.TemplatesSize.Add(_be, _ecc)
			if _dgb, _efd = _adc.Get(_daa); _efd != nil {
				return _e.Wrap(_efd, _ffg, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_fgc.AddBox(_dgb)
			_fe.ClassInstances.AddBitmaps(_fgc)
			_fe.CentroidPointsTemplates.AddPoint(_bdbd, _gga)
			_fe.FgTemplates.AddInt(_fcgf)
			_fe.UndilatedTemplates.AddBitmap(_dfa)
			_db = (_gab.Width - 2*JbAddedPixels) * (_gab.Height - 2*JbAddedPixels)
			if _efd = _fe.TemplateAreas.Add(_db); _efd != nil {
				return _e.Wrap(_efd, _ffg, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
		}
	}
	_fe.NumberOfClasses = len(_fe.UndilatedTemplates.Values)
	return nil
}
func _gec(_bc *_cf.Bitmap, _cb, _eab, _ged, _cdbe int, _ffc *_cf.Bitmap) (_ecg _d.Point, _gb error) {
	const _ebg = "\u0066i\u006e\u0061\u006c\u0041l\u0069\u0067\u006e\u006d\u0065n\u0074P\u006fs\u0069\u0074\u0069\u006f\u006e\u0069\u006eg"
	if _bc == nil {
		return _ecg, _e.Error(_ebg, "\u0073\u006f\u0075\u0072ce\u0020\u006e\u006f\u0074\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	if _ffc == nil {
		return _ecg, _e.Error(_ebg, "t\u0065\u006d\u0070\u006cat\u0065 \u006e\u006f\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_ce, _eea := _ffc.Width, _ffc.Height
	_dc, _ebb := _cb-_ged-JbAddedPixels, _eab-_cdbe-JbAddedPixels
	_b.Log.Trace("\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0077\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0068\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0078\u003a\u0020\u0027\u0025d'\u002c\u0020\u0062\u0079\u003a\u0020\u0027\u0025\u0064\u0027", _cb, _eab, _ce, _eea, _dc, _ebb)
	_acg, _gb := _cf.Rect(_dc, _ebb, _ce, _eea)
	if _gb != nil {
		return _ecg, _e.Wrap(_gb, _ebg, "")
	}
	_faa, _, _gb := _bc.ClipRectangle(_acg)
	if _gb != nil {
		_b.Log.Error("\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069\u0070\u0020\u0072\u0065c\u0074\u0061\u006e\u0067\u006c\u0065\u003a\u0020\u0025\u0076", _acg)
		return _ecg, _e.Wrap(_gb, _ebg, "")
	}
	_dd := _cf.New(_faa.Width, _faa.Height)
	_cfd := _ca.MaxInt32
	var _ffb, _ggd, _bbf, _cea, _cgf int
	for _ffb = -1; _ffb <= 1; _ffb++ {
		for _ggd = -1; _ggd <= 1; _ggd++ {
			if _, _gb = _cf.Copy(_dd, _faa); _gb != nil {
				return _ecg, _e.Wrap(_gb, _ebg, "")
			}
			if _gb = _dd.RasterOperation(_ggd, _ffb, _ce, _eea, _cf.PixSrcXorDst, _ffc, 0, 0); _gb != nil {
				return _ecg, _e.Wrap(_gb, _ebg, "")
			}
			_bbf = _dd.CountPixels()
			if _bbf < _cfd {
				_cea = _ggd
				_cgf = _ffb
				_cfd = _bbf
			}
		}
	}
	_ecg.X = _cea
	_ecg.Y = _cgf
	return _ecg, nil
}

const (
	MaxConnCompWidth = 350
	MaxCharCompWidth = 350
	MaxWordCompWidth = 1000
	MaxCompHeight    = 120
)

func DefaultSettings() Settings { _baa := &Settings{}; _baa.SetDefault(); return *_baa }
func (_ffa *Classer) getULCorners(_aca *_cf.Bitmap, _ec *_cf.Boxes) error {
	const _ea = "\u0067\u0065\u0074U\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073"
	if _aca == nil {
		return _e.Error(_ea, "\u006e\u0069l\u0020\u0069\u006da\u0067\u0065\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _ec == nil {
		return _e.Error(_ea, "\u006e\u0069\u006c\u0020\u0062\u006f\u0075\u006e\u0064\u0073")
	}
	if _ffa.PtaUL == nil {
		_ffa.PtaUL = &_cf.Points{}
	}
	_ef := len(*_ec)
	var (
		_cdb, _cc, _cag, _eaf int
		_bdb, _bdg, _fc, _ga  float32
		_eg                   error
		_def                  *_d.Rectangle
		_ccd                  *_cf.Bitmap
		_af                   _d.Point
	)
	for _ad := 0; _ad < _ef; _ad++ {
		_cdb = _ffa.BaseIndex + _ad
		if _bdb, _bdg, _eg = _ffa.CentroidPoints.GetGeometry(_cdb); _eg != nil {
			return _e.Wrap(_eg, _ea, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050o\u0069\u006e\u0074\u0073")
		}
		if _cc, _eg = _ffa.ClassIDs.Get(_cdb); _eg != nil {
			return _e.Wrap(_eg, _ea, "\u0043\u006c\u0061s\u0073\u0049\u0044\u0073\u002e\u0047\u0065\u0074")
		}
		if _fc, _ga, _eg = _ffa.CentroidPointsTemplates.GetGeometry(_cc); _eg != nil {
			return _e.Wrap(_eg, _ea, "\u0043\u0065\u006etr\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073")
		}
		_ge := _fc - _bdb
		_cg := _ga - _bdg
		if _ge >= 0 {
			_cag = int(_ge + 0.5)
		} else {
			_cag = int(_ge - 0.5)
		}
		if _cg >= 0 {
			_eaf = int(_cg + 0.5)
		} else {
			_eaf = int(_cg - 0.5)
		}
		if _def, _eg = _ec.Get(_ad); _eg != nil {
			return _e.Wrap(_eg, _ea, "")
		}
		_fab, _cge := _def.Min.X, _def.Min.Y
		_ccd, _eg = _ffa.UndilatedTemplates.GetBitmap(_cc)
		if _eg != nil {
			return _e.Wrap(_eg, _ea, "\u0055\u006e\u0064\u0069\u006c\u0061\u0074\u0065\u0064\u0054e\u006d\u0070\u006c\u0061\u0074\u0065\u0073.\u0047\u0065\u0074\u0028\u0069\u0043\u006c\u0061\u0073\u0073\u0029")
		}
		_af, _eg = _gec(_aca, _fab, _cge, _cag, _eaf, _ccd)
		if _eg != nil {
			return _e.Wrap(_eg, _ea, "")
		}
		_ffa.PtaUL.AddPoint(float32(_fab-_cag+_af.X), float32(_cge-_eaf+_af.Y))
	}
	return nil
}

type Method int

func (_ac *Classer) ComputeLLCorners() (_g error) {
	const _bf = "\u0043l\u0061\u0073\u0073\u0065\u0072\u002e\u0043\u006f\u006d\u0070\u0075t\u0065\u004c\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073"
	if _ac.PtaUL == nil {
		return _e.Error(_bf, "\u0055\u004c\u0020\u0043or\u006e\u0065\u0072\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	_ffe := len(*_ac.PtaUL)
	_ac.PtaLL = &_cf.Points{}
	var (
		_eb, _bfb float32
		_deg, _cd int
		_dg       *_cf.Bitmap
	)
	for _fb := 0; _fb < _ffe; _fb++ {
		_eb, _bfb, _g = _ac.PtaUL.GetGeometry(_fb)
		if _g != nil {
			_b.Log.Debug("\u0047e\u0074\u0074\u0069\u006e\u0067\u0020\u0050\u0074\u0061\u0055\u004c \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _g)
			return _e.Wrap(_g, _bf, "\u0050\u0074\u0061\u0055\u004c\u0020\u0047\u0065\u006fm\u0065\u0074\u0072\u0079")
		}
		_deg, _g = _ac.ClassIDs.Get(_fb)
		if _g != nil {
			_b.Log.Debug("\u0047\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0043\u006c\u0061s\u0073\u0049\u0044\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _g)
			return _e.Wrap(_g, _bf, "\u0043l\u0061\u0073\u0073\u0049\u0044")
		}
		_dg, _g = _ac.UndilatedTemplates.GetBitmap(_deg)
		if _g != nil {
			_b.Log.Debug("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0055\u006ed\u0069\u006c\u0061\u0074\u0065\u0064\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _g)
			return _e.Wrap(_g, _bf, "\u0055\u006e\u0064\u0069la\u0074\u0065\u0064\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073")
		}
		_cd = _dg.Height
		_ac.PtaLL.AddPoint(_eb, _bfb+float32(_cd))
	}
	return nil
}
func (_df *Classer) AddPage(inputPage *_cf.Bitmap, pageNumber int, method Method) (_a error) {
	const _cac = "\u0043l\u0061s\u0073\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065"
	_df.Widths[pageNumber] = inputPage.Width
	_df.Heights[pageNumber] = inputPage.Height
	if _a = _df.verifyMethod(method); _a != nil {
		return _e.Wrap(_a, _cac, "")
	}
	_de, _ff, _a := inputPage.GetComponents(_df.Settings.Components, _df.Settings.MaxCompWidth, _df.Settings.MaxCompHeight)
	if _a != nil {
		return _e.Wrap(_a, _cac, "")
	}
	_b.Log.Debug("\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _de)
	if _a = _df.addPageComponents(inputPage, _ff, _de, pageNumber, method); _a != nil {
		return _e.Wrap(_a, _cac, "")
	}
	return nil
}
func (_degc *Settings) SetDefault() {
	if _degc.MaxCompWidth == 0 {
		switch _degc.Components {
		case _cf.ComponentConn:
			_degc.MaxCompWidth = MaxConnCompWidth
		case _cf.ComponentCharacters:
			_degc.MaxCompWidth = MaxCharCompWidth
		case _cf.ComponentWords:
			_degc.MaxCompWidth = MaxWordCompWidth
		}
	}
	if _degc.MaxCompHeight == 0 {
		_degc.MaxCompHeight = MaxCompHeight
	}
	if _degc.Thresh == 0.0 {
		_degc.Thresh = 0.9
	}
	if _degc.WeightFactor == 0.0 {
		_degc.WeightFactor = 0.75
	}
	if _degc.RankHaus == 0.0 {
		_degc.RankHaus = 0.97
	}
	if _degc.SizeHaus == 0 {
		_degc.SizeHaus = 2
	}
}
func (_cde *Classer) classifyRankHouseOne(_ced *_cf.Boxes, _eccc, _ed, _bbb *_cf.Bitmaps, _gecc *_cf.Points, _aadg int) (_dgbe error) {
	const _gcg = "\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065"
	var (
		_ega, _dbf, _adb, _bg         float32
		_adcg                         int
		_ecf, _bgd, _gge, _fedb, _cda *_cf.Bitmap
		_eag, _bdgg                   bool
	)
	for _gf := 0; _gf < len(_eccc.Values); _gf++ {
		_bgd = _ed.Values[_gf]
		_gge = _bbb.Values[_gf]
		_ega, _dbf, _dgbe = _gecc.GetGeometry(_gf)
		if _dgbe != nil {
			return _e.Wrapf(_dgbe, _gcg, "\u0066\u0069\u0072\u0073\u0074\u0020\u0067\u0065\u006fm\u0065\u0074\u0072\u0079")
		}
		_gcd := len(_cde.UndilatedTemplates.Values)
		_eag = false
		_ebgd := _caeg(_cde, _bgd)
		for _adcg = _ebgd.Next(); _adcg > -1; {
			_fedb, _dgbe = _cde.UndilatedTemplates.GetBitmap(_adcg)
			if _dgbe != nil {
				return _e.Wrap(_dgbe, _gcg, "\u0062\u006d\u0033")
			}
			_cda, _dgbe = _cde.DilatedTemplates.GetBitmap(_adcg)
			if _dgbe != nil {
				return _e.Wrap(_dgbe, _gcg, "\u0062\u006d\u0034")
			}
			_adb, _bg, _dgbe = _cde.CentroidPointsTemplates.GetGeometry(_adcg)
			if _dgbe != nil {
				return _e.Wrap(_dgbe, _gcg, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0073")
			}
			_bdgg, _dgbe = _cf.HausTest(_bgd, _gge, _fedb, _cda, _ega-_adb, _dbf-_bg, MaxDiffWidth, MaxDiffHeight)
			if _dgbe != nil {
				return _e.Wrap(_dgbe, _gcg, "")
			}
			if _bdgg {
				_eag = true
				if _dgbe = _cde.ClassIDs.Add(_adcg); _dgbe != nil {
					return _e.Wrap(_dgbe, _gcg, "")
				}
				if _dgbe = _cde.ComponentPageNumbers.Add(_aadg); _dgbe != nil {
					return _e.Wrap(_dgbe, _gcg, "")
				}
				if _cde.Settings.KeepClassInstances {
					_ab, _gac := _cde.ClassInstances.GetBitmaps(_adcg)
					if _gac != nil {
						return _e.Wrap(_gac, _gcg, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_ecf, _gac = _eccc.GetBitmap(_gf)
					if _gac != nil {
						return _e.Wrap(_gac, _gcg, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_ab.AddBitmap(_ecf)
					_eca, _gac := _ced.Get(_gf)
					if _gac != nil {
						return _e.Wrap(_gac, _gcg, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_ab.AddBox(_eca)
				}
				break
			}
		}
		if !_eag {
			if _dgbe = _cde.ClassIDs.Add(_gcd); _dgbe != nil {
				return _e.Wrap(_dgbe, _gcg, "")
			}
			if _dgbe = _cde.ComponentPageNumbers.Add(_aadg); _dgbe != nil {
				return _e.Wrap(_dgbe, _gcg, "")
			}
			_gaf := &_cf.Bitmaps{}
			_ecf, _dgbe = _eccc.GetBitmap(_gf)
			if _dgbe != nil {
				return _e.Wrap(_dgbe, _gcg, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_gaf.Values = append(_gaf.Values, _ecf)
			_caca, _cbb := _ecf.Width, _ecf.Height
			_cde.TemplatesSize.Add(uint64(_cbb)*uint64(_caca), _gcd)
			_eadc, _cef := _ced.Get(_gf)
			if _cef != nil {
				return _e.Wrap(_cef, _gcg, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_gaf.AddBox(_eadc)
			_cde.ClassInstances.AddBitmaps(_gaf)
			_cde.CentroidPointsTemplates.AddPoint(_ega, _dbf)
			_cde.UndilatedTemplates.AddBitmap(_bgd)
			_cde.DilatedTemplates.AddBitmap(_gge)
		}
	}
	return nil
}

var _gc bool

func (_bb *Classer) verifyMethod(_ba Method) error {
	if _ba != RankHaus && _ba != Correlation {
		return _e.Error("\u0076\u0065\u0072i\u0066\u0079\u004d\u0065\u0074\u0068\u006f\u0064", "\u0069\u006e\u0076\u0061li\u0064\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072\u0020\u006d\u0065\u0074\u0068o\u0064")
	}
	return nil
}
func (_gg *Classer) addPageComponents(_fge *_cf.Bitmap, _aa *_cf.Boxes, _fa *_cf.Bitmaps, _da int, _ee Method) error {
	const _aaf = "\u0043l\u0061\u0073\u0073\u0065r\u002e\u0041\u0064\u0064\u0050a\u0067e\u0043o\u006d\u0070\u006f\u006e\u0065\u006e\u0074s"
	if _fge == nil {
		return _e.Error(_aaf, "\u006e\u0069\u006c\u0020\u0069\u006e\u0070\u0075\u0074 \u0070\u0061\u0067\u0065")
	}
	if _aa == nil || _fa == nil || len(*_aa) == 0 {
		_b.Log.Trace("\u0041\u0064\u0064P\u0061\u0067\u0065\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0073\u002e\u0020\u004e\u006f\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064", _fge)
		return nil
	}
	var _aad error
	switch _ee {
	case RankHaus:
		_aad = _gg.classifyRankHaus(_aa, _fa, _da)
	case Correlation:
		_aad = _gg.classifyCorrelation(_aa, _fa, _da)
	default:
		_b.Log.Debug("\u0055\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u003a\u0020'%\u0076\u0027", _ee)
		return _e.Error(_aaf, "\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064")
	}
	if _aad != nil {
		return _e.Wrap(_aad, _aaf, "")
	}
	if _aad = _gg.getULCorners(_fge, _aa); _aad != nil {
		return _e.Wrap(_aad, _aaf, "")
	}
	_dfb := len(*_aa)
	_gg.BaseIndex += _dfb
	if _aad = _gg.ComponentsNumber.Add(_dfb); _aad != nil {
		return _e.Wrap(_aad, _aaf, "")
	}
	return nil
}
func (_eefg Settings) Validate() error {
	const _edf = "\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0073\u002e\u0056\u0061\u006ci\u0064\u0061\u0074\u0065"
	if _eefg.Thresh < 0.4 || _eefg.Thresh > 0.98 {
		return _e.Error(_edf, "\u006a\u0062i\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0034\u0020\u002d\u0020\u0030\u002e\u0039\u0038\u005d")
	}
	if _eefg.WeightFactor < 0.0 || _eefg.WeightFactor > 1.0 {
		return _e.Error(_edf, "\u006a\u0062i\u0067\u0032\u0020\u0065\u006ec\u006f\u0064\u0065\u0072\u0020w\u0065\u0069\u0067\u0068\u0074\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0030\u0020\u002d\u0020\u0031\u002e\u0030\u005d")
	}
	if _eefg.RankHaus < 0.5 || _eefg.RankHaus > 1.0 {
		return _e.Error(_edf, "\u006a\u0062\u0069\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065 \u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 [\u0030\u002e\u0035\u0020\u002d\u0020\u0031\u002e\u0030\u005d")
	}
	if _eefg.SizeHaus < 1 || _eefg.SizeHaus > 10 {
		return _e.Error(_edf, "\u006a\u0062\u0069\u0067\u0032 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0073\u0069\u007a\u0065\u0020h\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0031\u0020\u002d\u0020\u0031\u0030]")
	}
	switch _eefg.Components {
	case _cf.ComponentConn, _cf.ComponentCharacters, _cf.ComponentWords:
	default:
		return _e.Error(_edf, "\u0069n\u0076\u0061\u006c\u0069d\u0020\u0063\u006c\u0061\u0073s\u0065r\u0020c\u006f\u006d\u0070\u006f\u006e\u0065\u006et")
	}
	return nil
}

type Classer struct {
	BaseIndex               int
	Settings                Settings
	ComponentsNumber        *_f.IntSlice
	TemplateAreas           *_f.IntSlice
	Widths                  map[int]int
	Heights                 map[int]int
	NumberOfClasses         int
	ClassInstances          *_cf.BitmapsArray
	UndilatedTemplates      *_cf.Bitmaps
	DilatedTemplates        *_cf.Bitmaps
	TemplatesSize           _f.IntsMap
	FgTemplates             *_f.NumSlice
	CentroidPoints          *_cf.Points
	CentroidPointsTemplates *_cf.Points
	ClassIDs                *_f.IntSlice
	ComponentPageNumbers    *_f.IntSlice
	PtaUL                   *_cf.Points
	PtaLL                   *_cf.Points
}

func (_aae *Classer) classifyRankHouseNonOne(_ggc *_cf.Boxes, _dee, _gdf, _eagd *_cf.Bitmaps, _faaa *_cf.Points, _bfd *_f.NumSlice, _aafc int) (_gbab error) {
	const _gff = "\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065"
	var (
		_aaa, _cacag, _gcc, _gbba     float32
		_agf, _cae, _acf              int
		_cgd, _afd, _gfd, _egce, _fbe *_cf.Bitmap
		_fbc, _fabb                   bool
	)
	_dggd := _cf.MakePixelSumTab8()
	for _dga := 0; _dga < len(_dee.Values); _dga++ {
		if _afd, _gbab = _gdf.GetBitmap(_dga); _gbab != nil {
			return _e.Wrap(_gbab, _gff, "b\u006d\u0073\u0031\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
		}
		if _agf, _gbab = _bfd.GetInt(_dga); _gbab != nil {
			_b.Log.Trace("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0046\u0047T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073 \u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _dga, _gbab)
		}
		if _gfd, _gbab = _eagd.GetBitmap(_dga); _gbab != nil {
			return _e.Wrap(_gbab, _gff, "b\u006d\u0073\u0032\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
		}
		if _aaa, _cacag, _gbab = _faaa.GetGeometry(_dga); _gbab != nil {
			return _e.Wrapf(_gbab, _gff, "\u0070t\u0061[\u0069\u005d\u002e\u0047\u0065\u006f\u006d\u0065\u0074\u0072\u0079")
		}
		_aab := len(_aae.UndilatedTemplates.Values)
		_fbc = false
		_dggf := _caeg(_aae, _afd)
		for _acf = _dggf.Next(); _acf > -1; {
			if _egce, _gbab = _aae.UndilatedTemplates.GetBitmap(_acf); _gbab != nil {
				return _e.Wrap(_gbab, _gff, "\u0070\u0069\u0078\u0061\u0074\u002e\u005b\u0069\u0043l\u0061\u0073\u0073\u005d")
			}
			if _cae, _gbab = _aae.FgTemplates.GetInt(_acf); _gbab != nil {
				_b.Log.Trace("\u0047\u0065\u0074\u0074\u0069\u006eg\u0020\u0046\u0047\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u005b\u0025d\u005d\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _acf, _gbab)
			}
			if _fbe, _gbab = _aae.DilatedTemplates.GetBitmap(_acf); _gbab != nil {
				return _e.Wrap(_gbab, _gff, "\u0070\u0069\u0078\u0061\u0074\u0064\u005b\u0069\u0043l\u0061\u0073\u0073\u005d")
			}
			if _gcc, _gbba, _gbab = _aae.CentroidPointsTemplates.GetGeometry(_acf); _gbab != nil {
				return _e.Wrap(_gbab, _gff, "\u0043\u0065\u006et\u0072\u006f\u0069\u0064P\u006f\u0069\u006e\u0074\u0073\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u005b\u0069\u0043\u006c\u0061\u0073\u0073\u005d")
			}
			_fabb, _gbab = _cf.RankHausTest(_afd, _gfd, _egce, _fbe, _aaa-_gcc, _cacag-_gbba, MaxDiffWidth, MaxDiffHeight, _agf, _cae, float32(_aae.Settings.RankHaus), _dggd)
			if _gbab != nil {
				return _e.Wrap(_gbab, _gff, "")
			}
			if _fabb {
				_fbc = true
				if _gbab = _aae.ClassIDs.Add(_acf); _gbab != nil {
					return _e.Wrap(_gbab, _gff, "")
				}
				if _gbab = _aae.ComponentPageNumbers.Add(_aafc); _gbab != nil {
					return _e.Wrap(_gbab, _gff, "")
				}
				if _aae.Settings.KeepClassInstances {
					_bca, _egceg := _aae.ClassInstances.GetBitmaps(_acf)
					if _egceg != nil {
						return _e.Wrap(_egceg, _gff, "\u0063\u002e\u0050\u0069\u0078\u0061\u0061\u002e\u0047\u0065\u0074B\u0069\u0074\u006d\u0061\u0070\u0073\u0028\u0069\u0043\u006ca\u0073\u0073\u0029")
					}
					if _cgd, _egceg = _dee.GetBitmap(_dga); _egceg != nil {
						return _e.Wrap(_egceg, _gff, "\u0070i\u0078\u0061\u005b\u0069\u005d")
					}
					_bca.Values = append(_bca.Values, _cgd)
					_cdee, _egceg := _ggc.Get(_dga)
					if _egceg != nil {
						return _e.Wrap(_egceg, _gff, "b\u006f\u0078\u0061\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
					}
					_bca.Boxes = append(_bca.Boxes, _cdee)
				}
				break
			}
		}
		if !_fbc {
			if _gbab = _aae.ClassIDs.Add(_aab); _gbab != nil {
				return _e.Wrap(_gbab, _gff, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			if _gbab = _aae.ComponentPageNumbers.Add(_aafc); _gbab != nil {
				return _e.Wrap(_gbab, _gff, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_cgc := &_cf.Bitmaps{}
			_cgd = _dee.Values[_dga]
			_cgc.AddBitmap(_cgd)
			_eabf, _fca := _cgd.Width, _cgd.Height
			_aae.TemplatesSize.Add(uint64(_eabf)*uint64(_fca), _aab)
			_babg, _fgeb := _ggc.Get(_dga)
			if _fgeb != nil {
				return _e.Wrap(_fgeb, _gff, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_cgc.AddBox(_babg)
			_aae.ClassInstances.AddBitmaps(_cgc)
			_aae.CentroidPointsTemplates.AddPoint(_aaa, _cacag)
			_aae.UndilatedTemplates.AddBitmap(_afd)
			_aae.DilatedTemplates.AddBitmap(_gfd)
			_aae.FgTemplates.AddInt(_agf)
		}
	}
	_aae.NumberOfClasses = len(_aae.UndilatedTemplates.Values)
	return nil
}
func Init(settings Settings) (*Classer, error) {
	const _cad = "\u0063\u006c\u0061s\u0073\u0065\u0072\u002e\u0049\u006e\u0069\u0074"
	_bd := &Classer{Settings: settings, Widths: map[int]int{}, Heights: map[int]int{}, TemplatesSize: _f.IntsMap{}, TemplateAreas: &_f.IntSlice{}, ComponentPageNumbers: &_f.IntSlice{}, ClassIDs: &_f.IntSlice{}, ComponentsNumber: &_f.IntSlice{}, CentroidPoints: &_cf.Points{}, CentroidPointsTemplates: &_cf.Points{}, UndilatedTemplates: &_cf.Bitmaps{}, DilatedTemplates: &_cf.Bitmaps{}, ClassInstances: &_cf.BitmapsArray{}, FgTemplates: &_f.NumSlice{}}
	if _fg := _bd.Settings.Validate(); _fg != nil {
		return nil, _e.Wrap(_fg, _cad, "")
	}
	return _bd, nil
}
func _caeg(_dea *Classer, _fedc *_cf.Bitmap) *similarTemplatesFinder {
	return &similarTemplatesFinder{Width: _fedc.Width, Height: _fedc.Height, Classer: _dea}
}

const (
	RankHaus Method = iota
	Correlation
)

type similarTemplatesFinder struct {
	Classer        *Classer
	Width          int
	Height         int
	Index          int
	CurrentNumbers []int
	N              int
}

const JbAddedPixels = 6

var TwoByTwoWalk = []int{0, 0, 0, 1, -1, 0, 0, -1, 1, 0, -1, 1, 1, 1, -1, -1, 1, -1, 0, -2, 2, 0, 0, 2, -2, 0, -1, -2, 1, -2, 2, -1, 2, 1, 1, 2, -1, 2, -2, 1, -2, -1, -2, -2, 2, -2, 2, 2, -2, 2}

type Settings struct {
	MaxCompWidth       int
	MaxCompHeight      int
	SizeHaus           int
	RankHaus           float64
	Thresh             float64
	WeightFactor       float64
	KeepClassInstances bool
	Components         _cf.Component
	Method             Method
}

func (_eeab *similarTemplatesFinder) Next() int {
	var (
		_bccf, _eba, _ffbg, _dec int
		_dfcd                    bool
		_acd                     *_cf.Bitmap
		_degce                   error
	)
	for {
		if _eeab.Index >= 25 {
			return -1
		}
		_eba = _eeab.Width + TwoByTwoWalk[2*_eeab.Index]
		_bccf = _eeab.Height + TwoByTwoWalk[2*_eeab.Index+1]
		if _bccf < 1 || _eba < 1 {
			_eeab.Index++
			continue
		}
		if len(_eeab.CurrentNumbers) == 0 {
			_eeab.CurrentNumbers, _dfcd = _eeab.Classer.TemplatesSize.GetSlice(uint64(_eba) * uint64(_bccf))
			if !_dfcd {
				_eeab.Index++
				continue
			}
			_eeab.N = 0
		}
		_ffbg = len(_eeab.CurrentNumbers)
		for ; _eeab.N < _ffbg; _eeab.N++ {
			_dec = _eeab.CurrentNumbers[_eeab.N]
			_acd, _degce = _eeab.Classer.DilatedTemplates.GetBitmap(_dec)
			if _degce != nil {
				_b.Log.Debug("\u0046\u0069\u006e\u0064\u004e\u0065\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0074\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020")
				return 0
			}
			if _acd.Width-2*JbAddedPixels == _eba && _acd.Height-2*JbAddedPixels == _bccf {
				return _dec
			}
		}
		_eeab.Index++
		_eeab.CurrentNumbers = nil
	}
}
