package classer

import (
	_e "image"
	_ec "math"

	_d "unitechio/gopdf/gopdf/common"
	_g "unitechio/gopdf/gopdf/internal/jbig2/basic"
	_ad "unitechio/gopdf/gopdf/internal/jbig2/bitmap"
	_eb "unitechio/gopdf/gopdf/internal/jbig2/errors"
)

const (
	MaxDiffWidth  = 2
	MaxDiffHeight = 2
)

func (_ggd *Classer) verifyMethod(_ege Method) error {
	if _ege != RankHaus && _ege != Correlation {
		return _eb.Error("\u0076\u0065\u0072i\u0066\u0079\u004d\u0065\u0074\u0068\u006f\u0064", "\u0069\u006e\u0076\u0061li\u0064\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072\u0020\u006d\u0065\u0074\u0068o\u0064")
	}
	return nil
}

func (_aba *Classer) getULCorners(_ecaa *_ad.Bitmap, _cg *_ad.Boxes) error {
	const _gd = "\u0067\u0065\u0074U\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073"
	if _ecaa == nil {
		return _eb.Error(_gd, "\u006e\u0069l\u0020\u0069\u006da\u0067\u0065\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _cg == nil {
		return _eb.Error(_gd, "\u006e\u0069\u006c\u0020\u0062\u006f\u0075\u006e\u0064\u0073")
	}
	if _aba.PtaUL == nil {
		_aba.PtaUL = &_ad.Points{}
	}
	_fd := len(*_cg)
	var (
		_ga, _bc, _dfd, _ba    int
		_edgf, _eg, _ade, _afe float32
		_bcf                   error
		_ca                    *_e.Rectangle
		_be                    *_ad.Bitmap
		_gg                    _e.Point
	)
	for _fdd := 0; _fdd < _fd; _fdd++ {
		_ga = _aba.BaseIndex + _fdd
		if _edgf, _eg, _bcf = _aba.CentroidPoints.GetGeometry(_ga); _bcf != nil {
			return _eb.Wrap(_bcf, _gd, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050o\u0069\u006e\u0074\u0073")
		}
		if _bc, _bcf = _aba.ClassIDs.Get(_ga); _bcf != nil {
			return _eb.Wrap(_bcf, _gd, "\u0043\u006c\u0061s\u0073\u0049\u0044\u0073\u002e\u0047\u0065\u0074")
		}
		if _ade, _afe, _bcf = _aba.CentroidPointsTemplates.GetGeometry(_bc); _bcf != nil {
			return _eb.Wrap(_bcf, _gd, "\u0043\u0065\u006etr\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073")
		}
		_bb := _ade - _edgf
		_fde := _afe - _eg
		if _bb >= 0 {
			_dfd = int(_bb + 0.5)
		} else {
			_dfd = int(_bb - 0.5)
		}
		if _fde >= 0 {
			_ba = int(_fde + 0.5)
		} else {
			_ba = int(_fde - 0.5)
		}
		if _ca, _bcf = _cg.Get(_fdd); _bcf != nil {
			return _eb.Wrap(_bcf, _gd, "")
		}
		_cc, _bd := _ca.Min.X, _ca.Min.Y
		_be, _bcf = _aba.UndilatedTemplates.GetBitmap(_bc)
		if _bcf != nil {
			return _eb.Wrap(_bcf, _gd, "\u0055\u006e\u0064\u0069\u006c\u0061\u0074\u0065\u0064\u0054e\u006d\u0070\u006c\u0061\u0074\u0065\u0073.\u0047\u0065\u0074\u0028\u0069\u0043\u006c\u0061\u0073\u0073\u0029")
		}
		_gg, _bcf = _gec(_ecaa, _cc, _bd, _dfd, _ba, _be)
		if _bcf != nil {
			return _eb.Wrap(_bcf, _gd, "")
		}
		_aba.PtaUL.AddPoint(float32(_cc-_dfd+_gg.X), float32(_bd-_ba+_gg.Y))
	}
	return nil
}

func (_da *Classer) addPageComponents(_ab *_ad.Bitmap, _ee *_ad.Boxes, _ecf *_ad.Bitmaps, _ebb int, _eca Method) error {
	const _gcb = "\u0043l\u0061\u0073\u0073\u0065r\u002e\u0041\u0064\u0064\u0050a\u0067e\u0043o\u006d\u0070\u006f\u006e\u0065\u006e\u0074s"
	if _ab == nil {
		return _eb.Error(_gcb, "\u006e\u0069\u006c\u0020\u0069\u006e\u0070\u0075\u0074 \u0070\u0061\u0067\u0065")
	}
	if _ee == nil || _ecf == nil || len(*_ee) == 0 {
		_d.Log.Trace("\u0041\u0064\u0064P\u0061\u0067\u0065\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0073\u002e\u0020\u004e\u006f\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064", _ab)
		return nil
	}
	var _ge error
	switch _eca {
	case RankHaus:
		_ge = _da.classifyRankHaus(_ee, _ecf, _ebb)
	case Correlation:
		_ge = _da.classifyCorrelation(_ee, _ecf, _ebb)
	default:
		_d.Log.Debug("\u0055\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u003a\u0020'%\u0076\u0027", _eca)
		return _eb.Error(_gcb, "\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064")
	}
	if _ge != nil {
		return _eb.Wrap(_ge, _gcb, "")
	}
	if _ge = _da.getULCorners(_ab, _ee); _ge != nil {
		return _eb.Wrap(_ge, _gcb, "")
	}
	_c := len(*_ee)
	_da.BaseIndex += _c
	if _ge = _da.ComponentsNumber.Add(_c); _ge != nil {
		return _eb.Wrap(_ge, _gcb, "")
	}
	return nil
}

const (
	MaxConnCompWidth = 350
	MaxCharCompWidth = 350
	MaxWordCompWidth = 1000
	MaxCompHeight    = 120
)

func (_gcbg *Classer) classifyRankHouseNonOne(_eba *_ad.Boxes, _daaa, _feg, _dbc *_ad.Bitmaps, _daff *_ad.Points, _egb *_g.NumSlice, _cfb int) (_eegg error) {
	const _gcg = "\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065"
	var (
		_dde, _dfb, _edfe, _gbc          float32
		_eggf, _bcd, _cef                int
		_gbe, _feag, _gcbc, _cadd, _dbcc *_ad.Bitmap
		_cca, _ccbd                      bool
	)
	_dbd := _ad.MakePixelSumTab8()
	for _dea := 0; _dea < len(_daaa.Values); _dea++ {
		if _feag, _eegg = _feg.GetBitmap(_dea); _eegg != nil {
			return _eb.Wrap(_eegg, _gcg, "b\u006d\u0073\u0031\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
		}
		if _eggf, _eegg = _egb.GetInt(_dea); _eegg != nil {
			_d.Log.Trace("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0046\u0047T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073 \u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _dea, _eegg)
		}
		if _gcbc, _eegg = _dbc.GetBitmap(_dea); _eegg != nil {
			return _eb.Wrap(_eegg, _gcg, "b\u006d\u0073\u0032\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
		}
		if _dde, _dfb, _eegg = _daff.GetGeometry(_dea); _eegg != nil {
			return _eb.Wrapf(_eegg, _gcg, "\u0070t\u0061[\u0069\u005d\u002e\u0047\u0065\u006f\u006d\u0065\u0074\u0072\u0079")
		}
		_aedd := len(_gcbg.UndilatedTemplates.Values)
		_cca = false
		_faf := _cda(_gcbg, _feag)
		for _cef = _faf.Next(); _cef > -1; {
			if _cadd, _eegg = _gcbg.UndilatedTemplates.GetBitmap(_cef); _eegg != nil {
				return _eb.Wrap(_eegg, _gcg, "\u0070\u0069\u0078\u0061\u0074\u002e\u005b\u0069\u0043l\u0061\u0073\u0073\u005d")
			}
			if _bcd, _eegg = _gcbg.FgTemplates.GetInt(_cef); _eegg != nil {
				_d.Log.Trace("\u0047\u0065\u0074\u0074\u0069\u006eg\u0020\u0046\u0047\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u005b\u0025d\u005d\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _cef, _eegg)
			}
			if _dbcc, _eegg = _gcbg.DilatedTemplates.GetBitmap(_cef); _eegg != nil {
				return _eb.Wrap(_eegg, _gcg, "\u0070\u0069\u0078\u0061\u0074\u0064\u005b\u0069\u0043l\u0061\u0073\u0073\u005d")
			}
			if _edfe, _gbc, _eegg = _gcbg.CentroidPointsTemplates.GetGeometry(_cef); _eegg != nil {
				return _eb.Wrap(_eegg, _gcg, "\u0043\u0065\u006et\u0072\u006f\u0069\u0064P\u006f\u0069\u006e\u0074\u0073\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u005b\u0069\u0043\u006c\u0061\u0073\u0073\u005d")
			}
			_ccbd, _eegg = _ad.RankHausTest(_feag, _gcbc, _cadd, _dbcc, _dde-_edfe, _dfb-_gbc, MaxDiffWidth, MaxDiffHeight, _eggf, _bcd, float32(_gcbg.Settings.RankHaus), _dbd)
			if _eegg != nil {
				return _eb.Wrap(_eegg, _gcg, "")
			}
			if _ccbd {
				_cca = true
				if _eegg = _gcbg.ClassIDs.Add(_cef); _eegg != nil {
					return _eb.Wrap(_eegg, _gcg, "")
				}
				if _eegg = _gcbg.ComponentPageNumbers.Add(_cfb); _eegg != nil {
					return _eb.Wrap(_eegg, _gcg, "")
				}
				if _gcbg.Settings.KeepClassInstances {
					_cgc, _gaf := _gcbg.ClassInstances.GetBitmaps(_cef)
					if _gaf != nil {
						return _eb.Wrap(_gaf, _gcg, "\u0063\u002e\u0050\u0069\u0078\u0061\u0061\u002e\u0047\u0065\u0074B\u0069\u0074\u006d\u0061\u0070\u0073\u0028\u0069\u0043\u006ca\u0073\u0073\u0029")
					}
					if _gbe, _gaf = _daaa.GetBitmap(_dea); _gaf != nil {
						return _eb.Wrap(_gaf, _gcg, "\u0070i\u0078\u0061\u005b\u0069\u005d")
					}
					_cgc.Values = append(_cgc.Values, _gbe)
					_agf, _gaf := _eba.Get(_dea)
					if _gaf != nil {
						return _eb.Wrap(_gaf, _gcg, "b\u006f\u0078\u0061\u002e\u0047\u0065\u0074\u0028\u0069\u0029")
					}
					_cgc.Boxes = append(_cgc.Boxes, _agf)
				}
				break
			}
		}
		if !_cca {
			if _eegg = _gcbg.ClassIDs.Add(_aedd); _eegg != nil {
				return _eb.Wrap(_eegg, _gcg, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			if _eegg = _gcbg.ComponentPageNumbers.Add(_cfb); _eegg != nil {
				return _eb.Wrap(_eegg, _gcg, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_fba := &_ad.Bitmaps{}
			_gbe = _daaa.Values[_dea]
			_fba.AddBitmap(_gbe)
			_dgb, _aeg := _gbe.Width, _gbe.Height
			_gcbg.TemplatesSize.Add(uint64(_dgb)*uint64(_aeg), _aedd)
			_ffc, _bdd := _eba.Get(_dea)
			if _bdd != nil {
				return _eb.Wrap(_bdd, _gcg, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_fba.AddBox(_ffc)
			_gcbg.ClassInstances.AddBitmaps(_fba)
			_gcbg.CentroidPointsTemplates.AddPoint(_dde, _dfb)
			_gcbg.UndilatedTemplates.AddBitmap(_feag)
			_gcbg.DilatedTemplates.AddBitmap(_gcbc)
			_gcbg.FgTemplates.AddInt(_eggf)
		}
	}
	_gcbg.NumberOfClasses = len(_gcbg.UndilatedTemplates.Values)
	return nil
}

const JbAddedPixels = 6
const (
	RankHaus Method = iota
	Correlation
)

func (_ffcf Settings) Validate() error {
	const _bfb = "\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0073\u002e\u0056\u0061\u006ci\u0064\u0061\u0074\u0065"
	if _ffcf.Thresh < 0.4 || _ffcf.Thresh > 0.98 {
		return _eb.Error(_bfb, "\u006a\u0062i\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0034\u0020\u002d\u0020\u0030\u002e\u0039\u0038\u005d")
	}
	if _ffcf.WeightFactor < 0.0 || _ffcf.WeightFactor > 1.0 {
		return _eb.Error(_bfb, "\u006a\u0062i\u0067\u0032\u0020\u0065\u006ec\u006f\u0064\u0065\u0072\u0020w\u0065\u0069\u0067\u0068\u0074\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0030\u0020\u002d\u0020\u0031\u002e\u0030\u005d")
	}
	if _ffcf.RankHaus < 0.5 || _ffcf.RankHaus > 1.0 {
		return _eb.Error(_bfb, "\u006a\u0062\u0069\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065 \u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 [\u0030\u002e\u0035\u0020\u002d\u0020\u0031\u002e\u0030\u005d")
	}
	if _ffcf.SizeHaus < 1 || _ffcf.SizeHaus > 10 {
		return _eb.Error(_bfb, "\u006a\u0062\u0069\u0067\u0032 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0073\u0069\u007a\u0065\u0020h\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0031\u0020\u002d\u0020\u0031\u0030]")
	}
	switch _ffcf.Components {
	case _ad.ComponentConn, _ad.ComponentCharacters, _ad.ComponentWords:
	default:
		return _eb.Error(_bfb, "\u0069n\u0076\u0061\u006c\u0069d\u0020\u0063\u006c\u0061\u0073s\u0065r\u0020c\u006f\u006d\u0070\u006f\u006e\u0065\u006et")
	}
	return nil
}

func (_bfe *Classer) ComputeLLCorners() (_ed error) {
	const _fb = "\u0043l\u0061\u0073\u0073\u0065\u0072\u002e\u0043\u006f\u006d\u0070\u0075t\u0065\u004c\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073"
	if _bfe.PtaUL == nil {
		return _eb.Error(_fb, "\u0055\u004c\u0020\u0043or\u006e\u0065\u0072\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	_fa := len(*_bfe.PtaUL)
	_bfe.PtaLL = &_ad.Points{}
	var (
		_fe, _ea  float32
		_efa, _df int
		_ae       *_ad.Bitmap
	)
	for _edg := 0; _edg < _fa; _edg++ {
		_fe, _ea, _ed = _bfe.PtaUL.GetGeometry(_edg)
		if _ed != nil {
			_d.Log.Debug("\u0047e\u0074\u0074\u0069\u006e\u0067\u0020\u0050\u0074\u0061\u0055\u004c \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ed)
			return _eb.Wrap(_ed, _fb, "\u0050\u0074\u0061\u0055\u004c\u0020\u0047\u0065\u006fm\u0065\u0074\u0072\u0079")
		}
		_efa, _ed = _bfe.ClassIDs.Get(_edg)
		if _ed != nil {
			_d.Log.Debug("\u0047\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0043\u006c\u0061s\u0073\u0049\u0044\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076", _ed)
			return _eb.Wrap(_ed, _fb, "\u0043l\u0061\u0073\u0073\u0049\u0044")
		}
		_ae, _ed = _bfe.UndilatedTemplates.GetBitmap(_efa)
		if _ed != nil {
			_d.Log.Debug("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0055\u006ed\u0069\u006c\u0061\u0074\u0065\u0064\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076", _ed)
			return _eb.Wrap(_ed, _fb, "\u0055\u006e\u0064\u0069la\u0074\u0065\u0064\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073")
		}
		_df = _ae.Height
		_bfe.PtaLL.AddPoint(_fe, _ea+float32(_df))
	}
	return nil
}

func _gec(_fgg *_ad.Bitmap, _bed, _eed, _gca, _caf int, _dg *_ad.Bitmap) (_fag _e.Point, _fbg error) {
	const _dgc = "\u0066i\u006e\u0061\u006c\u0041l\u0069\u0067\u006e\u006d\u0065n\u0074P\u006fs\u0069\u0074\u0069\u006f\u006e\u0069\u006eg"
	if _fgg == nil {
		return _fag, _eb.Error(_dgc, "\u0073\u006f\u0075\u0072ce\u0020\u006e\u006f\u0074\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	if _dg == nil {
		return _fag, _eb.Error(_dgc, "t\u0065\u006d\u0070\u006cat\u0065 \u006e\u006f\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_ebbg, _cad := _dg.Width, _dg.Height
	_cd, _gag := _bed-_gca-JbAddedPixels, _eed-_caf-JbAddedPixels
	_d.Log.Trace("\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0077\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0068\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0078\u003a\u0020\u0027\u0025d'\u002c\u0020\u0062\u0079\u003a\u0020\u0027\u0025\u0064\u0027", _bed, _eed, _ebbg, _cad, _cd, _gag)
	_dab, _fbg := _ad.Rect(_cd, _gag, _ebbg, _cad)
	if _fbg != nil {
		return _fag, _eb.Wrap(_fbg, _dgc, "")
	}
	_ggb, _, _fbg := _fgg.ClipRectangle(_dab)
	if _fbg != nil {
		_d.Log.Error("\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069\u0070\u0020\u0072\u0065c\u0074\u0061\u006e\u0067\u006c\u0065\u003a\u0020\u0025\u0076", _dab)
		return _fag, _eb.Wrap(_fbg, _dgc, "")
	}
	_fc := _ad.New(_ggb.Width, _ggb.Height)
	_cdg := _ec.MaxInt32
	var _ecg, _adef, _eea, _fgd, _eag int
	for _ecg = -1; _ecg <= 1; _ecg++ {
		for _adef = -1; _adef <= 1; _adef++ {
			if _, _fbg = _ad.Copy(_fc, _ggb); _fbg != nil {
				return _fag, _eb.Wrap(_fbg, _dgc, "")
			}
			if _fbg = _fc.RasterOperation(_adef, _ecg, _ebbg, _cad, _ad.PixSrcXorDst, _dg, 0, 0); _fbg != nil {
				return _fag, _eb.Wrap(_fbg, _dgc, "")
			}
			_eea = _fc.CountPixels()
			if _eea < _cdg {
				_fgd = _adef
				_eag = _ecg
				_cdg = _eea
			}
		}
	}
	_fag.X = _fgd
	_fag.Y = _eag
	return _fag, nil
}

func Init(settings Settings) (*Classer, error) {
	const _f = "\u0063\u006c\u0061s\u0073\u0065\u0072\u002e\u0049\u006e\u0069\u0074"
	_b := &Classer{Settings: settings, Widths: map[int]int{}, Heights: map[int]int{}, TemplatesSize: _g.IntsMap{}, TemplateAreas: &_g.IntSlice{}, ComponentPageNumbers: &_g.IntSlice{}, ClassIDs: &_g.IntSlice{}, ComponentsNumber: &_g.IntSlice{}, CentroidPoints: &_ad.Points{}, CentroidPointsTemplates: &_ad.Points{}, UndilatedTemplates: &_ad.Bitmaps{}, DilatedTemplates: &_ad.Bitmaps{}, ClassInstances: &_ad.BitmapsArray{}, FgTemplates: &_g.NumSlice{}}
	if _ef := _b.Settings.Validate(); _ef != nil {
		return nil, _eb.Wrap(_ef, _f, "")
	}
	return _b, nil
}

func (_ce *Classer) classifyCorrelation(_geb *_ad.Boxes, _daa *_ad.Bitmaps, _dc int) error {
	const _bee = "\u0063\u006c\u0061\u0073si\u0066\u0079\u0043\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e"
	if _geb == nil {
		return _eb.Error(_bee, "\u006e\u0065\u0077\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062o\u0078\u0065\u0073\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064")
	}
	if _daa == nil {
		return _eb.Error(_bee, "\u006e\u0065wC\u006f\u006d\u0070o\u006e\u0065\u006e\u0074s b\u0069tm\u0061\u0070\u0020\u0061\u0072\u0072\u0061y \u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
	}
	_gef := len(_daa.Values)
	if _gef == 0 {
		_d.Log.Debug("\u0063l\u0061\u0073s\u0069\u0066\u0079C\u006f\u0072\u0072\u0065\u006c\u0061\u0074i\u006f\u006e\u0020\u002d\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0070\u0069\u0078\u0061s\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079")
		return nil
	}
	var (
		_de, _dca *_ad.Bitmap
		_ceb      error
	)
	_bfd := &_ad.Bitmaps{Values: make([]*_ad.Bitmap, _gef)}
	for _dge, _adb := range _daa.Values {
		_dca, _ceb = _adb.AddBorderGeneral(JbAddedPixels, JbAddedPixels, JbAddedPixels, JbAddedPixels, 0)
		if _ceb != nil {
			return _eb.Wrap(_ceb, _bee, "")
		}
		_bfd.Values[_dge] = _dca
	}
	_edc := _ce.FgTemplates
	_cea := _ad.MakePixelSumTab8()
	_adc := _ad.MakePixelCentroidTab8()
	_afdg := make([]int, _gef)
	_egg := make([][]int, _gef)
	_ag := _ad.Points(make([]_ad.Point, _gef))
	_dad := &_ag
	var (
		_geg, _gegd       int
		_bg, _afee, _eeda int
		_daf, _gaa        int
		_afb              byte
	)
	for _fgde, _gdg := range _bfd.Values {
		_egg[_fgde] = make([]int, _gdg.Height)
		_geg = 0
		_gegd = 0
		_afee = (_gdg.Height - 1) * _gdg.RowStride
		_bg = 0
		for _gaa = _gdg.Height - 1; _gaa >= 0; _gaa, _afee = _gaa-1, _afee-_gdg.RowStride {
			_egg[_fgde][_gaa] = _bg
			_eeda = 0
			for _daf = 0; _daf < _gdg.RowStride; _daf++ {
				_afb = _gdg.Data[_afee+_daf]
				_eeda += _cea[_afb]
				_geg += _adc[_afb] + _daf*8*_cea[_afb]
			}
			_bg += _eeda
			_gegd += _eeda * _gaa
		}
		_afdg[_fgde] = _bg
		if _bg > 0 {
			(*_dad)[_fgde] = _ad.Point{X: float32(_geg) / float32(_bg), Y: float32(_gegd) / float32(_bg)}
		} else {
			(*_dad)[_fgde] = _ad.Point{X: float32(_gdg.Width) / float32(2), Y: float32(_gdg.Height) / float32(2)}
		}
	}
	if _ceb = _ce.CentroidPoints.Add(_dad); _ceb != nil {
		return _eb.Wrap(_ceb, _bee, "\u0063\u0065\u006et\u0072\u006f\u0069\u0064\u0020\u0061\u0064\u0064")
	}
	var (
		_fea, _ac, _db        int
		_aga                  float64
		_cgd, _fef, _aa, _bba float32
		_fec, _beec           _ad.Point
		_abc                  bool
		_eeg                  *similarTemplatesFinder
		_bgg                  int
		_eab                  *_ad.Bitmap
		_cde                  *_e.Rectangle
		_bbe                  *_ad.Bitmaps
	)
	for _bgg, _dca = range _bfd.Values {
		_ac = _afdg[_bgg]
		if _cgd, _fef, _ceb = _dad.GetGeometry(_bgg); _ceb != nil {
			return _eb.Wrap(_ceb, _bee, "\u0070t\u0061\u0020\u002d\u0020\u0069")
		}
		_abc = false
		_fac := len(_ce.UndilatedTemplates.Values)
		_eeg = _cda(_ce, _dca)
		for _fgdf := _eeg.Next(); _fgdf > -1; {
			if _eab, _ceb = _ce.UndilatedTemplates.GetBitmap(_fgdf); _ceb != nil {
				return _eb.Wrap(_ceb, _bee, "\u0075\u006e\u0069dl\u0061\u0074\u0065\u0064\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0062\u006d\u0032")
			}
			if _db, _ceb = _edc.GetInt(_fgdf); _ceb != nil {
				_d.Log.Trace("\u0046\u0047\u0020T\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u005b\u0069\u0063\u006ca\u0073\u0073\u005d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _ceb)
			}
			if _aa, _bba, _ceb = _ce.CentroidPointsTemplates.GetGeometry(_fgdf); _ceb != nil {
				return _eb.Wrap(_ceb, _bee, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074T\u0065\u006d\u0070\u006c\u0061\u0074e\u0073\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u00782\u002c\u0079\u0032\u0020")
			}
			if _ce.Settings.WeightFactor > 0.0 {
				if _fea, _ceb = _ce.TemplateAreas.Get(_fgdf); _ceb != nil {
					_d.Log.Trace("\u0054\u0065\u006dp\u006c\u0061\u0074\u0065A\u0072\u0065\u0061\u0073\u005b\u0069\u0063l\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0061\u0072\u0065\u0061\u0020\u0025\u0076", _ceb)
				}
				_aga = _ce.Settings.Thresh + (1.0-_ce.Settings.Thresh)*_ce.Settings.WeightFactor*float64(_db)/float64(_fea)
			} else {
				_aga = _ce.Settings.Thresh
			}
			_ff, _deb := _ad.CorrelationScoreThresholded(_dca, _eab, _ac, _db, _fec.X-_beec.X, _fec.Y-_beec.Y, MaxDiffWidth, MaxDiffHeight, _cea, _egg[_bgg], float32(_aga))
			if _deb != nil {
				return _eb.Wrap(_deb, _bee, "")
			}
			if _fcg {
				var (
					_eef, _ggg float64
					_aed, _cf  int
				)
				_eef, _deb = _ad.CorrelationScore(_dca, _eab, _ac, _db, _cgd-_aa, _fef-_bba, MaxDiffWidth, MaxDiffHeight, _cea)
				if _deb != nil {
					return _eb.Wrap(_deb, _bee, "d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065")
				}
				_ggg, _deb = _ad.CorrelationScoreSimple(_dca, _eab, _ac, _db, _cgd-_aa, _fef-_bba, MaxDiffWidth, MaxDiffHeight, _cea)
				if _deb != nil {
					return _eb.Wrap(_deb, _bee, "d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065")
				}
				_aed = int(_ec.Sqrt(_eef * float64(_ac) * float64(_db)))
				_cf = int(_ec.Sqrt(_ggg * float64(_ac) * float64(_db)))
				if (_eef >= _aga) != (_ggg >= _aga) {
					return _eb.Errorf(_bee, "\u0064\u0065\u0062\u0075\u0067\u0020\u0043\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u006f\u0072\u0065\u0020\u006d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020-\u0020\u0025d\u0028\u00250\u002e\u0034\u0066\u002c\u0020\u0025\u0076\u0029\u0020\u0076\u0073\u0020\u0025d(\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0076)\u0020\u0025\u0030\u002e\u0034\u0066", _aed, _eef, _eef >= _aga, _cf, _ggg, _ggg >= _aga, _eef-_ggg)
				}
				if _eef >= _aga != _ff {
					return _eb.Errorf(_bee, "\u0064\u0065\u0062\u0075\u0067\u0020\u0043o\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e \u0073\u0063\u006f\u0072\u0065 \u004d\u0069\u0073\u006d\u0061t\u0063\u0068 \u0062\u0065\u0074w\u0065\u0065\u006e\u0020\u0063\u006frr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020/\u0020\u0074\u0068\u0072\u0065s\u0068\u006f\u006c\u0064\u002e\u0020\u0043\u006f\u006dpa\u0072\u0069\u0073\u006f\u006e:\u0020\u0025\u0030\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0064\u0029\u0020\u003e\u003d\u0020\u00250\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u0029\u0020\u0076\u0073\u0020\u0025\u0076", _eef, _eef*float64(_ac)*float64(_db), _aed, _aga, float32(_aga)*float32(_ac)*float32(_db), _ff)
				}
			}
			if _ff {
				_abc = true
				if _deb = _ce.ClassIDs.Add(_fgdf); _deb != nil {
					return _eb.Wrap(_deb, _bee, "\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064")
				}
				if _deb = _ce.ComponentPageNumbers.Add(_dc); _deb != nil {
					return _eb.Wrap(_deb, _bee, "\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064")
				}
				if _ce.Settings.KeepClassInstances {
					if _de, _deb = _daa.GetBitmap(_bgg); _deb != nil {
						return _eb.Wrap(_deb, _bee, "\u004b\u0065\u0065\u0070Cl\u0061\u0073\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u002d \u0069")
					}
					if _bbe, _deb = _ce.ClassInstances.GetBitmaps(_fgdf); _deb != nil {
						return _eb.Wrap(_deb, _bee, "K\u0065\u0065\u0070\u0043\u006c\u0061s\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065s\u0020\u002d\u0020i\u0043l\u0061\u0073\u0073")
					}
					_bbe.AddBitmap(_de)
					if _cde, _deb = _geb.Get(_bgg); _deb != nil {
						return _eb.Wrap(_deb, _bee, "\u004be\u0065p\u0043\u006c\u0061\u0073\u0073I\u006e\u0073t\u0061\u006e\u0063\u0065\u0073")
					}
					_bbe.AddBox(_cde)
				}
				break
			}
		}
		if !_abc {
			if _ceb = _ce.ClassIDs.Add(_fac); _ceb != nil {
				return _eb.Wrap(_ceb, _bee, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			if _ceb = _ce.ComponentPageNumbers.Add(_dc); _ceb != nil {
				return _eb.Wrap(_ceb, _bee, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_bbe = &_ad.Bitmaps{}
			if _de, _ceb = _daa.GetBitmap(_bgg); _ceb != nil {
				return _eb.Wrap(_ceb, _bee, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_bbe.AddBitmap(_de)
			_bda, _aaf := _de.Width, _de.Height
			_dgg := uint64(_aaf) * uint64(_bda)
			_ce.TemplatesSize.Add(_dgg, _fac)
			if _cde, _ceb = _geb.Get(_bgg); _ceb != nil {
				return _eb.Wrap(_ceb, _bee, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_bbe.AddBox(_cde)
			_ce.ClassInstances.AddBitmaps(_bbe)
			_ce.CentroidPointsTemplates.AddPoint(_cgd, _fef)
			_ce.FgTemplates.AddInt(_ac)
			_ce.UndilatedTemplates.AddBitmap(_de)
			_fea = (_dca.Width - 2*JbAddedPixels) * (_dca.Height - 2*JbAddedPixels)
			if _ceb = _ce.TemplateAreas.Add(_fea); _ceb != nil {
				return _eb.Wrap(_ceb, _bee, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
		}
	}
	_ce.NumberOfClasses = len(_ce.UndilatedTemplates.Values)
	return nil
}

func (_gb *Classer) classifyRankHouseOne(_bbec *_ad.Boxes, _abg, _dfc, _gefg *_ad.Bitmaps, _cdb *_ad.Points, _efad int) (_abe error) {
	const _afbe = "\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065"
	var (
		_edcb, _gea, _ccb, _edb         float32
		_fddg                           int
		_cae, _ebe, _caea, _abab, _abcg *_ad.Bitmap
		_edf, _abb                      bool
	)
	for _geaa := 0; _geaa < len(_abg.Values); _geaa++ {
		_ebe = _dfc.Values[_geaa]
		_caea = _gefg.Values[_geaa]
		_edcb, _gea, _abe = _cdb.GetGeometry(_geaa)
		if _abe != nil {
			return _eb.Wrapf(_abe, _afbe, "\u0066\u0069\u0072\u0073\u0074\u0020\u0067\u0065\u006fm\u0065\u0074\u0072\u0079")
		}
		_gee := len(_gb.UndilatedTemplates.Values)
		_edf = false
		_febg := _cda(_gb, _ebe)
		for _fddg = _febg.Next(); _fddg > -1; {
			_abab, _abe = _gb.UndilatedTemplates.GetBitmap(_fddg)
			if _abe != nil {
				return _eb.Wrap(_abe, _afbe, "\u0062\u006d\u0033")
			}
			_abcg, _abe = _gb.DilatedTemplates.GetBitmap(_fddg)
			if _abe != nil {
				return _eb.Wrap(_abe, _afbe, "\u0062\u006d\u0034")
			}
			_ccb, _edb, _abe = _gb.CentroidPointsTemplates.GetGeometry(_fddg)
			if _abe != nil {
				return _eb.Wrap(_abe, _afbe, "\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0073")
			}
			_abb, _abe = _ad.HausTest(_ebe, _caea, _abab, _abcg, _edcb-_ccb, _gea-_edb, MaxDiffWidth, MaxDiffHeight)
			if _abe != nil {
				return _eb.Wrap(_abe, _afbe, "")
			}
			if _abb {
				_edf = true
				if _abe = _gb.ClassIDs.Add(_fddg); _abe != nil {
					return _eb.Wrap(_abe, _afbe, "")
				}
				if _abe = _gb.ComponentPageNumbers.Add(_efad); _abe != nil {
					return _eb.Wrap(_abe, _afbe, "")
				}
				if _gb.Settings.KeepClassInstances {
					_adf, _ggdc := _gb.ClassInstances.GetBitmaps(_fddg)
					if _ggdc != nil {
						return _eb.Wrap(_ggdc, _afbe, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_cae, _ggdc = _abg.GetBitmap(_geaa)
					if _ggdc != nil {
						return _eb.Wrap(_ggdc, _afbe, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_adf.AddBitmap(_cae)
					_dd, _ggdc := _bbec.Get(_geaa)
					if _ggdc != nil {
						return _eb.Wrap(_ggdc, _afbe, "\u004be\u0065\u0070\u0050\u0069\u0078\u0061a")
					}
					_adf.AddBox(_dd)
				}
				break
			}
		}
		if !_edf {
			if _abe = _gb.ClassIDs.Add(_gee); _abe != nil {
				return _eb.Wrap(_abe, _afbe, "")
			}
			if _abe = _gb.ComponentPageNumbers.Add(_efad); _abe != nil {
				return _eb.Wrap(_abe, _afbe, "")
			}
			_ceae := &_ad.Bitmaps{}
			_cae, _abe = _abg.GetBitmap(_geaa)
			if _abe != nil {
				return _eb.Wrap(_abe, _afbe, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_ceae.Values = append(_ceae.Values, _cae)
			_bfed, _ddc := _cae.Width, _cae.Height
			_gb.TemplatesSize.Add(uint64(_ddc)*uint64(_bfed), _gee)
			_ccg, _bdc := _bbec.Get(_geaa)
			if _bdc != nil {
				return _eb.Wrap(_bdc, _afbe, "\u0021\u0066\u006f\u0075\u006e\u0064")
			}
			_ceae.AddBox(_ccg)
			_gb.ClassInstances.AddBitmaps(_ceae)
			_gb.CentroidPointsTemplates.AddPoint(_edcb, _gea)
			_gb.UndilatedTemplates.AddBitmap(_ebe)
			_gb.DilatedTemplates.AddBitmap(_caea)
		}
	}
	return nil
}

type Classer struct {
	BaseIndex               int
	Settings                Settings
	ComponentsNumber        *_g.IntSlice
	TemplateAreas           *_g.IntSlice
	Widths                  map[int]int
	Heights                 map[int]int
	NumberOfClasses         int
	ClassInstances          *_ad.BitmapsArray
	UndilatedTemplates      *_ad.Bitmaps
	DilatedTemplates        *_ad.Bitmaps
	TemplatesSize           _g.IntsMap
	FgTemplates             *_g.NumSlice
	CentroidPoints          *_ad.Points
	CentroidPointsTemplates *_ad.Points
	ClassIDs                *_g.IntSlice
	ComponentPageNumbers    *_g.IntSlice
	PtaUL                   *_ad.Points
	PtaLL                   *_ad.Points
}

func (_dbg *Classer) classifyRankHaus(_ega *_ad.Boxes, _eeaf *_ad.Bitmaps, _efc int) error {
	const _caa = "\u0063\u006ca\u0073\u0073\u0069f\u0079\u0052\u0061\u006e\u006b\u0048\u0061\u0075\u0073"
	if _ega == nil {
		return _eb.Error(_caa, "\u0062\u006fx\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _eeaf == nil {
		return _eb.Error(_caa, "\u0070\u0069x\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	_cfg := len(_eeaf.Values)
	if _cfg == 0 {
		return _eb.Error(_caa, "e\u006dp\u0074\u0079\u0020\u006e\u0065\u0077\u0020\u0063o\u006d\u0070\u006f\u006een\u0074\u0073")
	}
	_afba := _eeaf.CountPixels()
	_fbd := _dbg.Settings.SizeHaus
	_fbf := _ad.SelCreateBrick(_fbd, _fbd, _fbd/2, _fbd/2, _ad.SelHit)
	_feb := &_ad.Bitmaps{Values: make([]*_ad.Bitmap, _cfg)}
	_afeg := &_ad.Bitmaps{Values: make([]*_ad.Bitmap, _cfg)}
	var (
		_edcd, _fcd, _dff *_ad.Bitmap
		_fee              error
	)
	for _cgdb := 0; _cgdb < _cfg; _cgdb++ {
		_edcd, _fee = _eeaf.GetBitmap(_cgdb)
		if _fee != nil {
			return _eb.Wrap(_fee, _caa, "")
		}
		_fcd, _fee = _edcd.AddBorderGeneral(JbAddedPixels, JbAddedPixels, JbAddedPixels, JbAddedPixels, 0)
		if _fee != nil {
			return _eb.Wrap(_fee, _caa, "")
		}
		_dff, _fee = _ad.Dilate(nil, _fcd, _fbf)
		if _fee != nil {
			return _eb.Wrap(_fee, _caa, "")
		}
		_feb.Values[_cfg] = _fcd
		_afeg.Values[_cfg] = _dff
	}
	_gf, _fee := _ad.Centroids(_feb.Values)
	if _fee != nil {
		return _eb.Wrap(_fee, _caa, "")
	}
	if _fee = _gf.Add(_dbg.CentroidPoints); _fee != nil {
		_d.Log.Trace("\u004e\u006f\u0020\u0063en\u0074\u0072\u006f\u0069\u0064\u0073\u0020\u0074\u006f\u0020\u0061\u0064\u0064")
	}
	if _dbg.Settings.RankHaus == 1.0 {
		_fee = _dbg.classifyRankHouseOne(_ega, _eeaf, _feb, _afeg, _gf, _efc)
	} else {
		_fee = _dbg.classifyRankHouseNonOne(_ega, _eeaf, _feb, _afeg, _gf, _afba, _efc)
	}
	if _fee != nil {
		return _eb.Wrap(_fee, _caa, "")
	}
	return nil
}
func DefaultSettings() Settings { _geab := &Settings{}; _geab.SetDefault(); return *_geab }
func (_gc *Classer) AddPage(inputPage *_ad.Bitmap, pageNumber int, method Method) (_fg error) {
	const _bf = "\u0043l\u0061s\u0073\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065"
	_gc.Widths[pageNumber] = inputPage.Width
	_gc.Heights[pageNumber] = inputPage.Height
	if _fg = _gc.verifyMethod(method); _fg != nil {
		return _eb.Wrap(_fg, _bf, "")
	}
	_af, _afd, _fg := inputPage.GetComponents(_gc.Settings.Components, _gc.Settings.MaxCompWidth, _gc.Settings.MaxCompHeight)
	if _fg != nil {
		return _eb.Wrap(_fg, _bf, "")
	}
	_d.Log.Debug("\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074s\u003a\u0020\u0025\u0076", _af)
	if _fg = _gc.addPageComponents(inputPage, _afd, _af, pageNumber, method); _fg != nil {
		return _eb.Wrap(_fg, _bf, "")
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

var TwoByTwoWalk = []int{0, 0, 0, 1, -1, 0, 0, -1, 1, 0, -1, 1, 1, 1, -1, -1, 1, -1, 0, -2, 2, 0, 0, 2, -2, 0, -1, -2, 1, -2, 2, -1, 2, 1, 1, 2, -1, 2, -2, 1, -2, -1, -2, -2, 2, -2, 2, 2, -2, 2}

func (_dcf *Settings) SetDefault() {
	if _dcf.MaxCompWidth == 0 {
		switch _dcf.Components {
		case _ad.ComponentConn:
			_dcf.MaxCompWidth = MaxConnCompWidth
		case _ad.ComponentCharacters:
			_dcf.MaxCompWidth = MaxCharCompWidth
		case _ad.ComponentWords:
			_dcf.MaxCompWidth = MaxWordCompWidth
		}
	}
	if _dcf.MaxCompHeight == 0 {
		_dcf.MaxCompHeight = MaxCompHeight
	}
	if _dcf.Thresh == 0.0 {
		_dcf.Thresh = 0.9
	}
	if _dcf.WeightFactor == 0.0 {
		_dcf.WeightFactor = 0.75
	}
	if _dcf.RankHaus == 0.0 {
		_dcf.RankHaus = 0.97
	}
	if _dcf.SizeHaus == 0 {
		_dcf.SizeHaus = 2
	}
}

func _cda(_afbg *Classer, _gge *_ad.Bitmap) *similarTemplatesFinder {
	return &similarTemplatesFinder{Width: _gge.Width, Height: _gge.Height, Classer: _afbg}
}

type (
	Method   int
	Settings struct {
		MaxCompWidth       int
		MaxCompHeight      int
		SizeHaus           int
		RankHaus           float64
		Thresh             float64
		WeightFactor       float64
		KeepClassInstances bool
		Components         _ad.Component
		Method             Method
	}
)

func (_baf *similarTemplatesFinder) Next() int {
	var (
		_bfbc, _egf, _ede, _gbf int
		_dec                    bool
		_ccd                    *_ad.Bitmap
		_gff                    error
	)
	for {
		if _baf.Index >= 25 {
			return -1
		}
		_egf = _baf.Width + TwoByTwoWalk[2*_baf.Index]
		_bfbc = _baf.Height + TwoByTwoWalk[2*_baf.Index+1]
		if _bfbc < 1 || _egf < 1 {
			_baf.Index++
			continue
		}
		if len(_baf.CurrentNumbers) == 0 {
			_baf.CurrentNumbers, _dec = _baf.Classer.TemplatesSize.GetSlice(uint64(_egf) * uint64(_bfbc))
			if !_dec {
				_baf.Index++
				continue
			}
			_baf.N = 0
		}
		_ede = len(_baf.CurrentNumbers)
		for ; _baf.N < _ede; _baf.N++ {
			_gbf = _baf.CurrentNumbers[_baf.N]
			_ccd, _gff = _baf.Classer.DilatedTemplates.GetBitmap(_gbf)
			if _gff != nil {
				_d.Log.Debug("\u0046\u0069\u006e\u0064\u004e\u0065\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0074\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020")
				return 0
			}
			if _ccd.Width-2*JbAddedPixels == _egf && _ccd.Height-2*JbAddedPixels == _bfbc {
				return _gbf
			}
		}
		_baf.Index++
		_baf.CurrentNumbers = nil
	}
}

var _fcg bool
