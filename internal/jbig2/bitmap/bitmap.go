package bitmap

import (
	_gc "encoding/binary"
	_ac "image"
	_ce "math"
	_e "sort"
	_fb "strings"
	_f "testing"

	_db "bitbucket.org/shenghui0779/gopdf/common"
	_c "bitbucket.org/shenghui0779/gopdf/internal/bitwise"
	_d "bitbucket.org/shenghui0779/gopdf/internal/imageutil"
	_fc "bitbucket.org/shenghui0779/gopdf/internal/jbig2/basic"
	_a "bitbucket.org/shenghui0779/gopdf/internal/jbig2/errors"
	_g "github.com/stretchr/testify/require"
)

func (_abg *Bitmap) Equivalent(s *Bitmap) bool { return _abg.equivalent(s) }
func _agce(_eecdb, _eede *Bitmap, _bbgae, _bgfg, _ceaf, _egbb, _cgd, _deada, _aaa, _aea int, _egcb CombinationOperator) error {
	var _aeg int
	_dbaa := func() { _aeg++; _ceaf += _eede.RowStride; _egbb += _eecdb.RowStride; _cgd += _eecdb.RowStride }
	for _aeg = _bbgae; _aeg < _bgfg; _dbaa() {
		var _aeca uint16
		_ged := _ceaf
		for _ece := _egbb; _ece <= _cgd; _ece++ {
			_gecg, _ffcd := _eede.GetByte(_ged)
			if _ffcd != nil {
				return _ffcd
			}
			_fgf, _ffcd := _eecdb.GetByte(_ece)
			if _ffcd != nil {
				return _ffcd
			}
			_aeca = (_aeca | uint16(_fgf)) << uint(_aea)
			_fgf = byte(_aeca >> 8)
			if _ece == _cgd {
				_fgf = _fge(uint(_deada), _fgf)
			}
			if _ffcd = _eede.SetByte(_ged, _bgbg(_gecg, _fgf, _egcb)); _ffcd != nil {
				return _ffcd
			}
			_ged++
			_aeca <<= uint(_aaa)
		}
	}
	return nil
}
func (_ecf *Bitmap) SetByte(index int, v byte) error {
	if index > len(_ecf.Data)-1 || index < 0 {
		return _a.Errorf("\u0053e\u0074\u0042\u0079\u0074\u0065", "\u0069\u006e\u0064\u0065x \u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020%\u0064", index)
	}
	_ecf.Data[index] = v
	return nil
}
func Centroid(bm *Bitmap, centTab, sumTab []int) (Point, error) { return bm.centroid(centTab, sumTab) }
func _bggbg(_edb, _dcaf *Bitmap) (*Bitmap, error) {
	if _dcaf == nil {
		return nil, _a.Error("\u0063\u006f\u0070\u0079\u0042\u0069\u0074\u006d\u0061\u0070", "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _dcaf == _edb {
		return _edb, nil
	}
	if _edb == nil {
		_edb = _dcaf.createTemplate()
		copy(_edb.Data, _dcaf.Data)
		return _edb, nil
	}
	_dbdd := _edb.resizeImageData(_dcaf)
	if _dbdd != nil {
		return nil, _a.Wrap(_dbdd, "\u0063\u006f\u0070\u0079\u0042\u0069\u0074\u006d\u0061\u0070", "")
	}
	_edb.Text = _dcaf.Text
	copy(_edb.Data, _dcaf.Data)
	return _edb, nil
}
func _afb() (_ffc [256]uint64) {
	for _fdc := 0; _fdc < 256; _fdc++ {
		if _fdc&0x01 != 0 {
			_ffc[_fdc] |= 0xff
		}
		if _fdc&0x02 != 0 {
			_ffc[_fdc] |= 0xff00
		}
		if _fdc&0x04 != 0 {
			_ffc[_fdc] |= 0xff0000
		}
		if _fdc&0x08 != 0 {
			_ffc[_fdc] |= 0xff000000
		}
		if _fdc&0x10 != 0 {
			_ffc[_fdc] |= 0xff00000000
		}
		if _fdc&0x20 != 0 {
			_ffc[_fdc] |= 0xff0000000000
		}
		if _fdc&0x40 != 0 {
			_ffc[_fdc] |= 0xff000000000000
		}
		if _fdc&0x80 != 0 {
			_ffc[_fdc] |= 0xff00000000000000
		}
	}
	return _ffc
}

type BitmapsArray struct {
	Values []*Bitmaps
	Boxes  []*_ac.Rectangle
}

func Centroids(bms []*Bitmap) (*Points, error) {
	_eea := make([]Point, len(bms))
	_eeaf := _caed()
	_ccfg := _cdgge()
	var _cdbb error
	for _adcd, _ddg := range bms {
		_eea[_adcd], _cdbb = _ddg.centroid(_eeaf, _ccfg)
		if _cdbb != nil {
			return nil, _cdbb
		}
	}
	_gdcd := Points(_eea)
	return &_gdcd, nil
}
func (_bec *Bitmap) CreateTemplate() *Bitmap { return _bec.createTemplate() }

type Component int
type fillSegment struct {
	_eeee  int
	_fdddg int
	_dabg  int
	_bege  int
}

func ClipBoxToRectangle(box *_ac.Rectangle, wi, hi int) (_agba *_ac.Rectangle, _fbgf error) {
	const _aeae = "\u0043l\u0069p\u0042\u006f\u0078\u0054\u006fR\u0065\u0063t\u0061\u006e\u0067\u006c\u0065"
	if box == nil {
		return nil, _a.Error(_aeae, "\u0027\u0062\u006f\u0078\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if box.Min.X >= wi || box.Min.Y >= hi || box.Max.X <= 0 || box.Max.Y <= 0 {
		return nil, _a.Error(_aeae, "\u0027\u0062\u006fx'\u0020\u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u0072\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065")
	}
	_gbb := *box
	_agba = &_gbb
	if _agba.Min.X < 0 {
		_agba.Max.X += _agba.Min.X
		_agba.Min.X = 0
	}
	if _agba.Min.Y < 0 {
		_agba.Max.Y += _agba.Min.Y
		_agba.Min.Y = 0
	}
	if _agba.Max.X > wi {
		_agba.Max.X = wi
	}
	if _agba.Max.Y > hi {
		_agba.Max.Y = hi
	}
	return _agba, nil
}
func MakePixelCentroidTab8() []int { return _caed() }
func _bcaa(_bgff *Bitmap, _ggba *_fc.Stack, _fcce, _ffcde int) (_agbf *_ac.Rectangle, _cgec error) {
	const _gggfb = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _bgff == nil {
		return nil, _a.Error(_gggfb, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _ggba == nil {
		return nil, _a.Error(_gggfb, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	_cddf, _ggec := _bgff.Width, _bgff.Height
	_ddgg := _cddf - 1
	_ecgaa := _ggec - 1
	if _fcce < 0 || _fcce > _ddgg || _ffcde < 0 || _ffcde > _ecgaa || !_bgff.GetPixel(_fcce, _ffcde) {
		return nil, nil
	}
	_gcea := _ac.Rect(100000, 100000, 0, 0)
	if _cgec = _gbdf(_ggba, _fcce, _fcce, _ffcde, 1, _ecgaa, &_gcea); _cgec != nil {
		return nil, _a.Wrap(_cgec, _gggfb, "\u0069\u006e\u0069t\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	if _cgec = _gbdf(_ggba, _fcce, _fcce, _ffcde+1, -1, _ecgaa, &_gcea); _cgec != nil {
		return nil, _a.Wrap(_cgec, _gggfb, "\u0032\u006ed\u0020\u0069\u006ei\u0074\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	_gcea.Min.X, _gcea.Max.X = _fcce, _fcce
	_gcea.Min.Y, _gcea.Max.Y = _ffcde, _ffcde
	var (
		_cfda  *fillSegment
		_decbe int
	)
	for _ggba.Len() > 0 {
		if _cfda, _cgec = _dabec(_ggba); _cgec != nil {
			return nil, _a.Wrap(_cgec, _gggfb, "")
		}
		_ffcde = _cfda._dabg
		for _fcce = _cfda._eeee - 1; _fcce >= 0 && _bgff.GetPixel(_fcce, _ffcde); _fcce-- {
			if _cgec = _bgff.SetPixel(_fcce, _ffcde, 0); _cgec != nil {
				return nil, _a.Wrap(_cgec, _gggfb, "\u0031s\u0074\u0020\u0073\u0065\u0074")
			}
		}
		if _fcce >= _cfda._eeee-1 {
			for {
				for _fcce++; _fcce <= _cfda._fdddg+1 && _fcce <= _ddgg && !_bgff.GetPixel(_fcce, _ffcde); _fcce++ {
				}
				_decbe = _fcce
				if !(_fcce <= _cfda._fdddg+1 && _fcce <= _ddgg) {
					break
				}
				for ; _fcce <= _ddgg && _bgff.GetPixel(_fcce, _ffcde); _fcce++ {
					if _cgec = _bgff.SetPixel(_fcce, _ffcde, 0); _cgec != nil {
						return nil, _a.Wrap(_cgec, _gggfb, "\u0032n\u0064\u0020\u0073\u0065\u0074")
					}
				}
				if _cgec = _gbdf(_ggba, _decbe, _fcce-1, _cfda._dabg, _cfda._bege, _ecgaa, &_gcea); _cgec != nil {
					return nil, _a.Wrap(_cgec, _gggfb, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
				}
				if _fcce > _cfda._fdddg {
					if _cgec = _gbdf(_ggba, _cfda._fdddg+1, _fcce-1, _cfda._dabg, -_cfda._bege, _ecgaa, &_gcea); _cgec != nil {
						return nil, _a.Wrap(_cgec, _gggfb, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
					}
				}
			}
			continue
		}
		_decbe = _fcce + 1
		if _decbe < _cfda._eeee {
			if _cgec = _gbdf(_ggba, _decbe, _cfda._eeee-1, _cfda._dabg, -_cfda._bege, _ecgaa, &_gcea); _cgec != nil {
				return nil, _a.Wrap(_cgec, _gggfb, "\u006c\u0065\u0061\u006b\u0020\u006f\u006e\u0020\u006c\u0065\u0066\u0074 \u0073\u0069\u0064\u0065")
			}
		}
		_fcce = _cfda._eeee
		for {
			for ; _fcce <= _ddgg && _bgff.GetPixel(_fcce, _ffcde); _fcce++ {
				if _cgec = _bgff.SetPixel(_fcce, _ffcde, 0); _cgec != nil {
					return nil, _a.Wrap(_cgec, _gggfb, "\u0032n\u0064\u0020\u0073\u0065\u0074")
				}
			}
			if _cgec = _gbdf(_ggba, _decbe, _fcce-1, _cfda._dabg, _cfda._bege, _ecgaa, &_gcea); _cgec != nil {
				return nil, _a.Wrap(_cgec, _gggfb, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
			}
			if _fcce > _cfda._fdddg {
				if _cgec = _gbdf(_ggba, _cfda._fdddg+1, _fcce-1, _cfda._dabg, -_cfda._bege, _ecgaa, &_gcea); _cgec != nil {
					return nil, _a.Wrap(_cgec, _gggfb, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
				}
			}
			for _fcce++; _fcce <= _cfda._fdddg+1 && _fcce <= _ddgg && !_bgff.GetPixel(_fcce, _ffcde); _fcce++ {
			}
			_decbe = _fcce
			if !(_fcce <= _cfda._fdddg+1 && _fcce <= _ddgg) {
				break
			}
		}
	}
	_gcea.Max.X++
	_gcea.Max.Y++
	return &_gcea, nil
}

type byWidth Bitmaps

const (
	ComponentConn Component = iota
	ComponentCharacters
	ComponentWords
)

func (_ebbe *Bitmaps) HeightSorter() func(_fbadg, _cceb int) bool {
	return func(_acbag, _cgfac int) bool {
		_ecce := _ebbe.Values[_acbag].Height < _ebbe.Values[_cgfac].Height
		_db.Log.Debug("H\u0065i\u0067\u0068\u0074\u003a\u0020\u0025\u0076\u0020<\u0020\u0025\u0076\u0020= \u0025\u0076", _ebbe.Values[_acbag].Height, _ebbe.Values[_cgfac].Height, _ecce)
		return _ecce
	}
}
func init() {
	for _gea := 0; _gea < 256; _gea++ {
		_afdc[_gea] = uint8(_gea&0x1) + (uint8(_gea>>1) & 0x1) + (uint8(_gea>>2) & 0x1) + (uint8(_gea>>3) & 0x1) + (uint8(_gea>>4) & 0x1) + (uint8(_gea>>5) & 0x1) + (uint8(_gea>>6) & 0x1) + (uint8(_gea>>7) & 0x1)
	}
}
func (_acbeg *Bitmap) Copy() *Bitmap {
	_fdg := make([]byte, len(_acbeg.Data))
	copy(_fdg, _acbeg.Data)
	return &Bitmap{Width: _acbeg.Width, Height: _acbeg.Height, RowStride: _acbeg.RowStride, Data: _fdg, Color: _acbeg.Color, Text: _acbeg.Text, BitmapNumber: _acbeg.BitmapNumber, Special: _acbeg.Special}
}
func TstFrameBitmapData() []byte { return _eefb.Data }
func (_bcfa *Bitmap) connComponentsBitmapsBB(_bebee *Bitmaps, _eabc int) (_bgfge *Boxes, _dcdb error) {
	const _bcgce = "\u0063\u006f\u006enC\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0042\u0069\u0074\u006d\u0061\u0070\u0073\u0042\u0042"
	if _eabc != 4 && _eabc != 8 {
		return nil, _a.Error(_bcgce, "\u0063\u006f\u006e\u006e\u0065\u0063t\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0027\u0034\u0027\u0020\u006fr\u0020\u0027\u0038\u0027")
	}
	if _bebee == nil {
		return nil, _a.Error(_bcgce, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0042\u0069\u0074ma\u0070\u0073")
	}
	if len(_bebee.Values) > 0 {
		return nil, _a.Error(_bcgce, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u006fn\u002d\u0065\u006d\u0070\u0074\u0079\u0020\u0042\u0069\u0074m\u0061\u0070\u0073")
	}
	if _bcfa.Zero() {
		return &Boxes{}, nil
	}
	var (
		_ccd, _ceffe, _fbdb, _aabd *Bitmap
	)
	_bcfa.setPadBits(0)
	if _ccd, _dcdb = _bggbg(nil, _bcfa); _dcdb != nil {
		return nil, _a.Wrap(_dcdb, _bcgce, "\u0062\u006d\u0031")
	}
	if _ceffe, _dcdb = _bggbg(nil, _bcfa); _dcdb != nil {
		return nil, _a.Wrap(_dcdb, _bcgce, "\u0062\u006d\u0032")
	}
	_aedd := &_fc.Stack{}
	_aedd.Aux = &_fc.Stack{}
	_bgfge = &Boxes{}
	var (
		_fbac, _efcf int
		_fcad        _ac.Point
		_gfgc        bool
		_ecgb        *_ac.Rectangle
	)
	for {
		if _fcad, _gfgc, _dcdb = _ccd.nextOnPixel(_fbac, _efcf); _dcdb != nil {
			return nil, _a.Wrap(_dcdb, _bcgce, "")
		}
		if !_gfgc {
			break
		}
		if _ecgb, _dcdb = _bebaf(_ccd, _aedd, _fcad.X, _fcad.Y, _eabc); _dcdb != nil {
			return nil, _a.Wrap(_dcdb, _bcgce, "")
		}
		if _dcdb = _bgfge.Add(_ecgb); _dcdb != nil {
			return nil, _a.Wrap(_dcdb, _bcgce, "")
		}
		if _fbdb, _dcdb = _ccd.clipRectangle(_ecgb, nil); _dcdb != nil {
			return nil, _a.Wrap(_dcdb, _bcgce, "\u0062\u006d\u0033")
		}
		if _aabd, _dcdb = _ceffe.clipRectangle(_ecgb, nil); _dcdb != nil {
			return nil, _a.Wrap(_dcdb, _bcgce, "\u0062\u006d\u0034")
		}
		if _, _dcdb = _gege(_fbdb, _fbdb, _aabd); _dcdb != nil {
			return nil, _a.Wrap(_dcdb, _bcgce, "\u0062m\u0033\u0020\u005e\u0020\u0062\u006d4")
		}
		if _dcdb = _ceffe.RasterOperation(_ecgb.Min.X, _ecgb.Min.Y, _ecgb.Dx(), _ecgb.Dy(), PixSrcXorDst, _fbdb, 0, 0); _dcdb != nil {
			return nil, _a.Wrap(_dcdb, _bcgce, "\u0062\u006d\u0032\u0020\u002d\u0058\u004f\u0052\u002d>\u0020\u0062\u006d\u0033")
		}
		_bebee.AddBitmap(_fbdb)
		_fbac = _fcad.X
		_efcf = _fcad.Y
	}
	_bebee.Boxes = *_bgfge
	return _bgfge, nil
}
func _feag(_ebacg, _dfdfg *Bitmap, _edcd, _cedg int) (*Bitmap, error) {
	const _ebegd = "\u006fp\u0065\u006e\u0042\u0072\u0069\u0063k"
	if _dfdfg == nil {
		return nil, _a.Error(_ebegd, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _edcd < 1 && _cedg < 1 {
		return nil, _a.Error(_ebegd, "\u0068\u0053\u0069\u007ae \u003c\u0020\u0031\u0020\u0026\u0026\u0020\u0076\u0053\u0069\u007a\u0065\u0020\u003c \u0031")
	}
	if _edcd == 1 && _cedg == 1 {
		return _dfdfg.Copy(), nil
	}
	if _edcd == 1 || _cedg == 1 {
		var _dgec error
		_fggec := SelCreateBrick(_cedg, _edcd, _cedg/2, _edcd/2, SelHit)
		_ebacg, _dgec = _bccf(_ebacg, _dfdfg, _fggec)
		if _dgec != nil {
			return nil, _a.Wrap(_dgec, _ebegd, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _ebacg, nil
	}
	_efacd := SelCreateBrick(1, _edcd, 0, _edcd/2, SelHit)
	_gaaa := SelCreateBrick(_cedg, 1, _cedg/2, 0, SelHit)
	_fgef, _dbdcd := _abbc(nil, _dfdfg, _efacd)
	if _dbdcd != nil {
		return nil, _a.Wrap(_dbdcd, _ebegd, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	_ebacg, _dbdcd = _abbc(_ebacg, _fgef, _gaaa)
	if _dbdcd != nil {
		return nil, _a.Wrap(_dbdcd, _ebegd, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	_, _dbdcd = _egcg(_fgef, _ebacg, _efacd)
	if _dbdcd != nil {
		return nil, _a.Wrap(_dbdcd, _ebegd, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	_, _dbdcd = _egcg(_ebacg, _fgef, _gaaa)
	if _dbdcd != nil {
		return nil, _a.Wrap(_dbdcd, _ebegd, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	return _ebacg, nil
}
func (_cfaa *Bitmap) inverseData() {
	if _efe := _cfaa.RasterOperation(0, 0, _cfaa.Width, _cfaa.Height, PixNotDst, nil, 0, 0); _efe != nil {
		_db.Log.Debug("\u0049n\u0076\u0065\u0072\u0073e\u0020\u0064\u0061\u0074\u0061 \u0066a\u0069l\u0065\u0064\u003a\u0020\u0027\u0025\u0076'", _efe)
	}
	if _cfaa.Color == Chocolate {
		_cfaa.Color = Vanilla
	} else {
		_cfaa.Color = Chocolate
	}
}
func (_bbf *Bitmap) And(s *Bitmap) (_adg *Bitmap, _bdfg error) {
	const _gcc = "\u0042\u0069\u0074\u006d\u0061\u0070\u002e\u0041\u006e\u0064"
	if _bbf == nil {
		return nil, _a.Error(_gcc, "\u0027b\u0069t\u006d\u0061\u0070\u0020\u0027b\u0027\u0020i\u0073\u0020\u006e\u0069\u006c")
	}
	if s == nil {
		return nil, _a.Error(_gcc, "\u0062\u0069\u0074\u006d\u0061\u0070\u0020\u0027\u0073\u0027\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if !_bbf.SizesEqual(s) {
		_db.Log.Debug("\u0025\u0073\u0020-\u0020\u0042\u0069\u0074\u006d\u0061\u0070\u0020\u0027\u0073\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0073\u0069\u007a\u0065 \u0077\u0069\u0074\u0068\u0020\u0027\u0062\u0027", _gcc)
	}
	if _adg, _bdfg = _bggbg(_adg, _bbf); _bdfg != nil {
		return nil, _a.Wrap(_bdfg, _gcc, "\u0063\u0061\u006e't\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020\u0027\u0064\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _bdfg = _adg.RasterOperation(0, 0, _adg.Width, _adg.Height, PixSrcAndDst, s, 0, 0); _bdfg != nil {
		return nil, _a.Wrap(_bdfg, _gcc, "")
	}
	return _adg, nil
}
func _cfdc(_bfgf *Bitmap, _egfgd, _fdag, _bbb, _gadc int, _bceb RasterOperator) {
	if _egfgd < 0 {
		_bbb += _egfgd
		_egfgd = 0
	}
	_gfd := _egfgd + _bbb - _bfgf.Width
	if _gfd > 0 {
		_bbb -= _gfd
	}
	if _fdag < 0 {
		_gadc += _fdag
		_fdag = 0
	}
	_aagfd := _fdag + _gadc - _bfgf.Height
	if _aagfd > 0 {
		_gadc -= _aagfd
	}
	if _bbb <= 0 || _gadc <= 0 {
		return
	}
	if (_egfgd & 7) == 0 {
		_fed(_bfgf, _egfgd, _fdag, _bbb, _gadc, _bceb)
	} else {
		_bfcdc(_bfgf, _egfgd, _fdag, _bbb, _gadc, _bceb)
	}
}
func (_aff *Bitmap) CountPixels() int { return _aff.countPixels() }

type Boxes []*_ac.Rectangle

func (_gaec *Bitmap) addBorderGeneral(_decc, _bcf, _faac, _eda int, _gcb int) (*Bitmap, error) {
	const _geg = "\u0061\u0064d\u0042\u006f\u0072d\u0065\u0072\u0047\u0065\u006e\u0065\u0072\u0061\u006c"
	if _decc < 0 || _bcf < 0 || _faac < 0 || _eda < 0 {
		return nil, _a.Error(_geg, "n\u0065\u0067\u0061\u0074iv\u0065 \u0062\u006f\u0072\u0064\u0065r\u0020\u0061\u0064\u0064\u0065\u0064")
	}
	_gac, _aaeb := _gaec.Width, _gaec.Height
	_ggd := _gac + _decc + _bcf
	_egg := _aaeb + _faac + _eda
	_gdag := New(_ggd, _egg)
	_gdag.Color = _gaec.Color
	_bgeg := PixClr
	if _gcb > 0 {
		_bgeg = PixSet
	}
	_fgc := _gdag.RasterOperation(0, 0, _decc, _egg, _bgeg, nil, 0, 0)
	if _fgc != nil {
		return nil, _a.Wrap(_fgc, _geg, "\u006c\u0065\u0066\u0074")
	}
	_fgc = _gdag.RasterOperation(_ggd-_bcf, 0, _bcf, _egg, _bgeg, nil, 0, 0)
	if _fgc != nil {
		return nil, _a.Wrap(_fgc, _geg, "\u0072\u0069\u0067h\u0074")
	}
	_fgc = _gdag.RasterOperation(0, 0, _ggd, _faac, _bgeg, nil, 0, 0)
	if _fgc != nil {
		return nil, _a.Wrap(_fgc, _geg, "\u0074\u006f\u0070")
	}
	_fgc = _gdag.RasterOperation(0, _egg-_eda, _ggd, _eda, _bgeg, nil, 0, 0)
	if _fgc != nil {
		return nil, _a.Wrap(_fgc, _geg, "\u0062\u006f\u0074\u0074\u006f\u006d")
	}
	_fgc = _gdag.RasterOperation(_decc, _faac, _gac, _aaeb, PixSrc, _gaec, 0, 0)
	if _fgc != nil {
		return nil, _a.Wrap(_fgc, _geg, "\u0063\u006f\u0070\u0079")
	}
	return _gdag, nil
}

type CombinationOperator int

func _cbed(_dgba *Bitmap, _fdca ...MorphProcess) (_ababg *Bitmap, _gafc error) {
	const _cbae = "\u006d\u006f\u0072\u0070\u0068\u0053\u0065\u0071\u0075\u0065\u006e\u0063\u0065"
	if _dgba == nil {
		return nil, _a.Error(_cbae, "\u006d\u006f\u0072\u0070\u0068\u0053\u0065\u0071\u0075\u0065\u006e\u0063\u0065 \u0073\u006f\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if len(_fdca) == 0 {
		return nil, _a.Error(_cbae, "m\u006f\u0072\u0070\u0068\u0053\u0065q\u0075\u0065\u006e\u0063\u0065\u002c \u0073\u0065\u0071\u0075\u0065\u006e\u0063e\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	if _gafc = _gaecg(_fdca...); _gafc != nil {
		return nil, _a.Wrap(_gafc, _cbae, "")
	}
	var _dbaad, _agdbd, _abdb int
	_ababg = _dgba.Copy()
	for _, _dgae := range _fdca {
		switch _dgae.Operation {
		case MopDilation:
			_dbaad, _agdbd = _dgae.getWidthHeight()
			_ababg, _gafc = DilateBrick(nil, _ababg, _dbaad, _agdbd)
			if _gafc != nil {
				return nil, _a.Wrap(_gafc, _cbae, "")
			}
		case MopErosion:
			_dbaad, _agdbd = _dgae.getWidthHeight()
			_ababg, _gafc = _fcee(nil, _ababg, _dbaad, _agdbd)
			if _gafc != nil {
				return nil, _a.Wrap(_gafc, _cbae, "")
			}
		case MopOpening:
			_dbaad, _agdbd = _dgae.getWidthHeight()
			_ababg, _gafc = _feag(nil, _ababg, _dbaad, _agdbd)
			if _gafc != nil {
				return nil, _a.Wrap(_gafc, _cbae, "")
			}
		case MopClosing:
			_dbaad, _agdbd = _dgae.getWidthHeight()
			_ababg, _gafc = _dbdcge(nil, _ababg, _dbaad, _agdbd)
			if _gafc != nil {
				return nil, _a.Wrap(_gafc, _cbae, "")
			}
		case MopRankBinaryReduction:
			_ababg, _gafc = _cec(_ababg, _dgae.Arguments...)
			if _gafc != nil {
				return nil, _a.Wrap(_gafc, _cbae, "")
			}
		case MopReplicativeBinaryExpansion:
			_ababg, _gafc = _cgae(_ababg, _dgae.Arguments[0])
			if _gafc != nil {
				return nil, _a.Wrap(_gafc, _cbae, "")
			}
		case MopAddBorder:
			_abdb = _dgae.Arguments[0]
			_ababg, _gafc = _ababg.AddBorder(_abdb, 0)
			if _gafc != nil {
				return nil, _a.Wrap(_gafc, _cbae, "")
			}
		default:
			return nil, _a.Error(_cbae, "i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006d\u006fr\u0070\u0068\u004f\u0070\u0065\u0072\u0061ti\u006f\u006e\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0074\u006f t\u0068\u0065 \u0073\u0065\u0071\u0075\u0065\u006e\u0063\u0065")
		}
	}
	if _abdb > 0 {
		_ababg, _gafc = _ababg.RemoveBorder(_abdb)
		if _gafc != nil {
			return nil, _a.Wrap(_gafc, _cbae, "\u0062\u006f\u0072\u0064\u0065\u0072\u0020\u003e\u0020\u0030")
		}
	}
	return _ababg, nil
}
func _eggb(_cccg *Bitmap, _bgbgd, _aaaac, _aabab, _dgfb int, _aeb RasterOperator, _dcec *Bitmap, _acbc, _daacd int) error {
	var (
		_efcd          bool
		_bcfg          bool
		_edcb          int
		_feea          int
		_fcfcc         int
		_eafc          bool
		_eege          byte
		_aefg          int
		_dcbc          int
		_dgbdg         int
		_bgfaa, _ccgcc int
	)
	_babd := 8 - (_bgbgd & 7)
	_efad := _gded[_babd]
	_dabb := _cccg.RowStride*_aaaac + (_bgbgd >> 3)
	_agab := _dcec.RowStride*_daacd + (_acbc >> 3)
	if _aabab < _babd {
		_efcd = true
		_efad &= _fcgba[8-_babd+_aabab]
	}
	if !_efcd {
		_edcb = (_aabab - _babd) >> 3
		if _edcb > 0 {
			_bcfg = true
			_feea = _dabb + 1
			_fcfcc = _agab + 1
		}
	}
	_aefg = (_bgbgd + _aabab) & 7
	if !(_efcd || _aefg == 0) {
		_eafc = true
		_eege = _fcgba[_aefg]
		_dcbc = _dabb + 1 + _edcb
		_dgbdg = _agab + 1 + _edcb
	}
	switch _aeb {
	case PixSrc:
		for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
			_cccg.Data[_dabb] = _cbcaa(_cccg.Data[_dabb], _dcec.Data[_agab], _efad)
			_dabb += _cccg.RowStride
			_agab += _dcec.RowStride
		}
		if _bcfg {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				for _ccgcc = 0; _ccgcc < _edcb; _ccgcc++ {
					_cccg.Data[_feea+_ccgcc] = _dcec.Data[_fcfcc+_ccgcc]
				}
				_feea += _cccg.RowStride
				_fcfcc += _dcec.RowStride
			}
		}
		if _eafc {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				_cccg.Data[_dcbc] = _cbcaa(_cccg.Data[_dcbc], _dcec.Data[_dgbdg], _eege)
				_dcbc += _cccg.RowStride
				_dgbdg += _dcec.RowStride
			}
		}
	case PixNotSrc:
		for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
			_cccg.Data[_dabb] = _cbcaa(_cccg.Data[_dabb], ^_dcec.Data[_agab], _efad)
			_dabb += _cccg.RowStride
			_agab += _dcec.RowStride
		}
		if _bcfg {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				for _ccgcc = 0; _ccgcc < _edcb; _ccgcc++ {
					_cccg.Data[_feea+_ccgcc] = ^_dcec.Data[_fcfcc+_ccgcc]
				}
				_feea += _cccg.RowStride
				_fcfcc += _dcec.RowStride
			}
		}
		if _eafc {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				_cccg.Data[_dcbc] = _cbcaa(_cccg.Data[_dcbc], ^_dcec.Data[_dgbdg], _eege)
				_dcbc += _cccg.RowStride
				_dgbdg += _dcec.RowStride
			}
		}
	case PixSrcOrDst:
		for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
			_cccg.Data[_dabb] = _cbcaa(_cccg.Data[_dabb], _dcec.Data[_agab]|_cccg.Data[_dabb], _efad)
			_dabb += _cccg.RowStride
			_agab += _dcec.RowStride
		}
		if _bcfg {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				for _ccgcc = 0; _ccgcc < _edcb; _ccgcc++ {
					_cccg.Data[_feea+_ccgcc] |= _dcec.Data[_fcfcc+_ccgcc]
				}
				_feea += _cccg.RowStride
				_fcfcc += _dcec.RowStride
			}
		}
		if _eafc {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				_cccg.Data[_dcbc] = _cbcaa(_cccg.Data[_dcbc], _dcec.Data[_dgbdg]|_cccg.Data[_dcbc], _eege)
				_dcbc += _cccg.RowStride
				_dgbdg += _dcec.RowStride
			}
		}
	case PixSrcAndDst:
		for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
			_cccg.Data[_dabb] = _cbcaa(_cccg.Data[_dabb], _dcec.Data[_agab]&_cccg.Data[_dabb], _efad)
			_dabb += _cccg.RowStride
			_agab += _dcec.RowStride
		}
		if _bcfg {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				for _ccgcc = 0; _ccgcc < _edcb; _ccgcc++ {
					_cccg.Data[_feea+_ccgcc] &= _dcec.Data[_fcfcc+_ccgcc]
				}
				_feea += _cccg.RowStride
				_fcfcc += _dcec.RowStride
			}
		}
		if _eafc {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				_cccg.Data[_dcbc] = _cbcaa(_cccg.Data[_dcbc], _dcec.Data[_dgbdg]&_cccg.Data[_dcbc], _eege)
				_dcbc += _cccg.RowStride
				_dgbdg += _dcec.RowStride
			}
		}
	case PixSrcXorDst:
		for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
			_cccg.Data[_dabb] = _cbcaa(_cccg.Data[_dabb], _dcec.Data[_agab]^_cccg.Data[_dabb], _efad)
			_dabb += _cccg.RowStride
			_agab += _dcec.RowStride
		}
		if _bcfg {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				for _ccgcc = 0; _ccgcc < _edcb; _ccgcc++ {
					_cccg.Data[_feea+_ccgcc] ^= _dcec.Data[_fcfcc+_ccgcc]
				}
				_feea += _cccg.RowStride
				_fcfcc += _dcec.RowStride
			}
		}
		if _eafc {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				_cccg.Data[_dcbc] = _cbcaa(_cccg.Data[_dcbc], _dcec.Data[_dgbdg]^_cccg.Data[_dcbc], _eege)
				_dcbc += _cccg.RowStride
				_dgbdg += _dcec.RowStride
			}
		}
	case PixNotSrcOrDst:
		for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
			_cccg.Data[_dabb] = _cbcaa(_cccg.Data[_dabb], ^(_dcec.Data[_agab])|_cccg.Data[_dabb], _efad)
			_dabb += _cccg.RowStride
			_agab += _dcec.RowStride
		}
		if _bcfg {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				for _ccgcc = 0; _ccgcc < _edcb; _ccgcc++ {
					_cccg.Data[_feea+_ccgcc] |= ^(_dcec.Data[_fcfcc+_ccgcc])
				}
				_feea += _cccg.RowStride
				_fcfcc += _dcec.RowStride
			}
		}
		if _eafc {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				_cccg.Data[_dcbc] = _cbcaa(_cccg.Data[_dcbc], ^(_dcec.Data[_dgbdg])|_cccg.Data[_dcbc], _eege)
				_dcbc += _cccg.RowStride
				_dgbdg += _dcec.RowStride
			}
		}
	case PixNotSrcAndDst:
		for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
			_cccg.Data[_dabb] = _cbcaa(_cccg.Data[_dabb], ^(_dcec.Data[_agab])&_cccg.Data[_dabb], _efad)
			_dabb += _cccg.RowStride
			_agab += _dcec.RowStride
		}
		if _bcfg {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				for _ccgcc = 0; _ccgcc < _edcb; _ccgcc++ {
					_cccg.Data[_feea+_ccgcc] &= ^_dcec.Data[_fcfcc+_ccgcc]
				}
				_feea += _cccg.RowStride
				_fcfcc += _dcec.RowStride
			}
		}
		if _eafc {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				_cccg.Data[_dcbc] = _cbcaa(_cccg.Data[_dcbc], ^(_dcec.Data[_dgbdg])&_cccg.Data[_dcbc], _eege)
				_dcbc += _cccg.RowStride
				_dgbdg += _dcec.RowStride
			}
		}
	case PixSrcOrNotDst:
		for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
			_cccg.Data[_dabb] = _cbcaa(_cccg.Data[_dabb], _dcec.Data[_agab]|^(_cccg.Data[_dabb]), _efad)
			_dabb += _cccg.RowStride
			_agab += _dcec.RowStride
		}
		if _bcfg {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				for _ccgcc = 0; _ccgcc < _edcb; _ccgcc++ {
					_cccg.Data[_feea+_ccgcc] = _dcec.Data[_fcfcc+_ccgcc] | ^(_cccg.Data[_feea+_ccgcc])
				}
				_feea += _cccg.RowStride
				_fcfcc += _dcec.RowStride
			}
		}
		if _eafc {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				_cccg.Data[_dcbc] = _cbcaa(_cccg.Data[_dcbc], _dcec.Data[_dgbdg]|^(_cccg.Data[_dcbc]), _eege)
				_dcbc += _cccg.RowStride
				_dgbdg += _dcec.RowStride
			}
		}
	case PixSrcAndNotDst:
		for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
			_cccg.Data[_dabb] = _cbcaa(_cccg.Data[_dabb], _dcec.Data[_agab]&^(_cccg.Data[_dabb]), _efad)
			_dabb += _cccg.RowStride
			_agab += _dcec.RowStride
		}
		if _bcfg {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				for _ccgcc = 0; _ccgcc < _edcb; _ccgcc++ {
					_cccg.Data[_feea+_ccgcc] = _dcec.Data[_fcfcc+_ccgcc] &^ (_cccg.Data[_feea+_ccgcc])
				}
				_feea += _cccg.RowStride
				_fcfcc += _dcec.RowStride
			}
		}
		if _eafc {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				_cccg.Data[_dcbc] = _cbcaa(_cccg.Data[_dcbc], _dcec.Data[_dgbdg]&^(_cccg.Data[_dcbc]), _eege)
				_dcbc += _cccg.RowStride
				_dgbdg += _dcec.RowStride
			}
		}
	case PixNotPixSrcOrDst:
		for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
			_cccg.Data[_dabb] = _cbcaa(_cccg.Data[_dabb], ^(_dcec.Data[_agab] | _cccg.Data[_dabb]), _efad)
			_dabb += _cccg.RowStride
			_agab += _dcec.RowStride
		}
		if _bcfg {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				for _ccgcc = 0; _ccgcc < _edcb; _ccgcc++ {
					_cccg.Data[_feea+_ccgcc] = ^(_dcec.Data[_fcfcc+_ccgcc] | _cccg.Data[_feea+_ccgcc])
				}
				_feea += _cccg.RowStride
				_fcfcc += _dcec.RowStride
			}
		}
		if _eafc {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				_cccg.Data[_dcbc] = _cbcaa(_cccg.Data[_dcbc], ^(_dcec.Data[_dgbdg] | _cccg.Data[_dcbc]), _eege)
				_dcbc += _cccg.RowStride
				_dgbdg += _dcec.RowStride
			}
		}
	case PixNotPixSrcAndDst:
		for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
			_cccg.Data[_dabb] = _cbcaa(_cccg.Data[_dabb], ^(_dcec.Data[_agab] & _cccg.Data[_dabb]), _efad)
			_dabb += _cccg.RowStride
			_agab += _dcec.RowStride
		}
		if _bcfg {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				for _ccgcc = 0; _ccgcc < _edcb; _ccgcc++ {
					_cccg.Data[_feea+_ccgcc] = ^(_dcec.Data[_fcfcc+_ccgcc] & _cccg.Data[_feea+_ccgcc])
				}
				_feea += _cccg.RowStride
				_fcfcc += _dcec.RowStride
			}
		}
		if _eafc {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				_cccg.Data[_dcbc] = _cbcaa(_cccg.Data[_dcbc], ^(_dcec.Data[_dgbdg] & _cccg.Data[_dcbc]), _eege)
				_dcbc += _cccg.RowStride
				_dgbdg += _dcec.RowStride
			}
		}
	case PixNotPixSrcXorDst:
		for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
			_cccg.Data[_dabb] = _cbcaa(_cccg.Data[_dabb], ^(_dcec.Data[_agab] ^ _cccg.Data[_dabb]), _efad)
			_dabb += _cccg.RowStride
			_agab += _dcec.RowStride
		}
		if _bcfg {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				for _ccgcc = 0; _ccgcc < _edcb; _ccgcc++ {
					_cccg.Data[_feea+_ccgcc] = ^(_dcec.Data[_fcfcc+_ccgcc] ^ _cccg.Data[_feea+_ccgcc])
				}
				_feea += _cccg.RowStride
				_fcfcc += _dcec.RowStride
			}
		}
		if _eafc {
			for _bgfaa = 0; _bgfaa < _dgfb; _bgfaa++ {
				_cccg.Data[_dcbc] = _cbcaa(_cccg.Data[_dcbc], ^(_dcec.Data[_dgbdg] ^ _cccg.Data[_dcbc]), _eege)
				_dcbc += _cccg.RowStride
				_dgbdg += _dcec.RowStride
			}
		}
	default:
		_db.Log.Debug("I\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070e\u0072\u0061\u0074o\u0072:\u0020\u0025\u0064", _aeb)
		return _a.Error("\u0072\u0061\u0073\u0074er\u004f\u0070\u0056\u0041\u006c\u0069\u0067\u006e\u0065\u0064\u004c\u006f\u0077", "\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func TstWriteSymbols(t *_f.T, bms *Bitmaps, src *Bitmap) {
	for _bedb := 0; _bedb < bms.Size(); _bedb++ {
		_cbbac := bms.Values[_bedb]
		_ggff := bms.Boxes[_bedb]
		_gacf := src.RasterOperation(_ggff.Min.X, _ggff.Min.Y, _cbbac.Width, _cbbac.Height, PixSrc, _cbbac, 0, 0)
		_g.NoError(t, _gacf)
	}
}
func (_fcb *Bitmap) RemoveBorderGeneral(left, right, top, bot int) (*Bitmap, error) {
	return _fcb.removeBorderGeneral(left, right, top, bot)
}
func (_adfa *Bitmaps) selectByIndicator(_gfefa *_fc.NumSlice) (_dadec *Bitmaps, _bffc error) {
	const _aecb = "\u0042i\u0074\u006d\u0061\u0070s\u002e\u0073\u0065\u006c\u0065c\u0074B\u0079I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _adfa == nil {
		return nil, _a.Error(_aecb, "\u0027\u0062\u0027 b\u0069\u0074\u006d\u0061\u0070\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _gfefa == nil {
		return nil, _a.Error(_aecb, "'\u006e\u0061\u0027\u0020\u0069\u006ed\u0069\u0063\u0061\u0074\u006f\u0072\u0073\u0020\u006eo\u0074\u0020\u0064e\u0066i\u006e\u0065\u0064")
	}
	if len(_adfa.Values) == 0 {
		return _adfa, nil
	}
	if len(*_gfefa) != len(_adfa.Values) {
		return nil, _a.Errorf(_aecb, "\u006ea\u0020\u006ce\u006e\u0067\u0074\u0068:\u0020\u0025\u0064,\u0020\u0069\u0073\u0020\u0064\u0069\u0066\u0066\u0065re\u006e\u0074\u0020t\u0068\u0061n\u0020\u0062\u0069\u0074\u006d\u0061p\u0073\u003a \u0025\u0064", len(*_gfefa), len(_adfa.Values))
	}
	var _bgbec, _dagd, _fbef int
	for _dagd = 0; _dagd < len(*_gfefa); _dagd++ {
		if _bgbec, _bffc = _gfefa.GetInt(_dagd); _bffc != nil {
			return nil, _a.Wrap(_bffc, _aecb, "f\u0069\u0072\u0073\u0074\u0020\u0063\u0068\u0065\u0063\u006b")
		}
		if _bgbec == 1 {
			_fbef++
		}
	}
	if _fbef == len(_adfa.Values) {
		return _adfa, nil
	}
	_dadec = &Bitmaps{}
	_eaddb := len(_adfa.Values) == len(_adfa.Boxes)
	for _dagd = 0; _dagd < len(*_gfefa); _dagd++ {
		if _bgbec = int((*_gfefa)[_dagd]); _bgbec == 0 {
			continue
		}
		_dadec.Values = append(_dadec.Values, _adfa.Values[_dagd])
		if _eaddb {
			_dadec.Boxes = append(_dadec.Boxes, _adfa.Boxes[_dagd])
		}
	}
	return _dadec, nil
}
func _gagf(_fag, _bca *Bitmap, _ggb int, _afc []byte, _dae int) (_bad error) {
	const _gcf = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0033"
	var (
		_fga, _ddbf, _ceb, _cbc, _eab, _eag, _dea, _fgg int
		_gfe, _bdfd, _dag, _edd                         uint32
		_daef, _gcg                                     byte
		_fef                                            uint16
	)
	_ebeae := make([]byte, 4)
	_dfa := make([]byte, 4)
	for _ceb = 0; _ceb < _fag.Height-1; _ceb, _cbc = _ceb+2, _cbc+1 {
		_fga = _ceb * _fag.RowStride
		_ddbf = _cbc * _bca.RowStride
		for _eab, _eag = 0, 0; _eab < _dae; _eab, _eag = _eab+4, _eag+1 {
			for _dea = 0; _dea < 4; _dea++ {
				_fgg = _fga + _eab + _dea
				if _fgg <= len(_fag.Data)-1 && _fgg < _fga+_fag.RowStride {
					_ebeae[_dea] = _fag.Data[_fgg]
				} else {
					_ebeae[_dea] = 0x00
				}
				_fgg = _fga + _fag.RowStride + _eab + _dea
				if _fgg <= len(_fag.Data)-1 && _fgg < _fga+(2*_fag.RowStride) {
					_dfa[_dea] = _fag.Data[_fgg]
				} else {
					_dfa[_dea] = 0x00
				}
			}
			_gfe = _gc.BigEndian.Uint32(_ebeae)
			_bdfd = _gc.BigEndian.Uint32(_dfa)
			_dag = _gfe & _bdfd
			_dag |= _dag << 1
			_edd = _gfe | _bdfd
			_edd &= _edd << 1
			_bdfd = _dag & _edd
			_bdfd &= 0xaaaaaaaa
			_gfe = _bdfd | (_bdfd << 7)
			_daef = byte(_gfe >> 24)
			_gcg = byte((_gfe >> 8) & 0xff)
			_fgg = _ddbf + _eag
			if _fgg+1 == len(_bca.Data)-1 || _fgg+1 >= _ddbf+_bca.RowStride {
				if _bad = _bca.SetByte(_fgg, _afc[_daef]); _bad != nil {
					return _a.Wrapf(_bad, _gcf, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _fgg)
				}
			} else {
				_fef = (uint16(_afc[_daef]) << 8) | uint16(_afc[_gcg])
				if _bad = _bca.setTwoBytes(_fgg, _fef); _bad != nil {
					return _a.Wrapf(_bad, _gcf, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _fgg)
				}
				_eag++
			}
		}
	}
	return nil
}
func (_eabg *Bitmap) nextOnPixelLow(_cae, _bfcd, _ebfg, _egfc, _adac int) (_fcbg _ac.Point, _adab bool, _bcgc error) {
	const _aaba = "B\u0069\u0074\u006d\u0061p.\u006ee\u0078\u0074\u004f\u006e\u0050i\u0078\u0065\u006c\u004c\u006f\u0077"
	var (
		_caf  int
		_cada byte
	)
	_bgba := _adac * _ebfg
	_baaa := _bgba + (_egfc / 8)
	if _cada, _bcgc = _eabg.GetByte(_baaa); _bcgc != nil {
		return _fcbg, false, _a.Wrap(_bcgc, _aaba, "\u0078\u0053\u0074\u0061\u0072\u0074\u0020\u0061\u006e\u0064 \u0079\u0053\u0074\u0061\u0072\u0074\u0020o\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	if _cada != 0 {
		_edag := _egfc - (_egfc % 8) + 7
		for _caf = _egfc; _caf <= _edag && _caf < _cae; _caf++ {
			if _eabg.GetPixel(_caf, _adac) {
				_fcbg.X = _caf
				_fcbg.Y = _adac
				return _fcbg, true, nil
			}
		}
	}
	_ddfg := (_egfc / 8) + 1
	_caf = 8 * _ddfg
	var _gbff int
	for _baaa = _bgba + _ddfg; _caf < _cae; _baaa, _caf = _baaa+1, _caf+8 {
		if _cada, _bcgc = _eabg.GetByte(_baaa); _bcgc != nil {
			return _fcbg, false, _a.Wrap(_bcgc, _aaba, "r\u0065\u0073\u0074\u0020of\u0020t\u0068\u0065\u0020\u006c\u0069n\u0065\u0020\u0062\u0079\u0074\u0065")
		}
		if _cada == 0 {
			continue
		}
		for _gbff = 0; _gbff < 8 && _caf < _cae; _gbff, _caf = _gbff+1, _caf+1 {
			if _eabg.GetPixel(_caf, _adac) {
				_fcbg.X = _caf
				_fcbg.Y = _adac
				return _fcbg, true, nil
			}
		}
	}
	for _egfd := _adac + 1; _egfd < _bfcd; _egfd++ {
		_bgba = _egfd * _ebfg
		for _baaa, _caf = _bgba, 0; _caf < _cae; _baaa, _caf = _baaa+1, _caf+8 {
			if _cada, _bcgc = _eabg.GetByte(_baaa); _bcgc != nil {
				return _fcbg, false, _a.Wrap(_bcgc, _aaba, "\u0066o\u006cl\u006f\u0077\u0069\u006e\u0067\u0020\u006c\u0069\u006e\u0065\u0073")
			}
			if _cada == 0 {
				continue
			}
			for _gbff = 0; _gbff < 8 && _caf < _cae; _gbff, _caf = _gbff+1, _caf+1 {
				if _eabg.GetPixel(_caf, _egfd) {
					_fcbg.X = _caf
					_fcbg.Y = _egfd
					return _fcbg, true, nil
				}
			}
		}
	}
	return _fcbg, false, nil
}
func (_bcba MorphProcess) getWidthHeight() (_feef, _agfg int) {
	return _bcba.Arguments[0], _bcba.Arguments[1]
}
func _ecc(_gagb *Bitmap, _cbf int, _ed []byte) (_dce *Bitmap, _adc error) {
	const _gb = "\u0072\u0065\u0064\u0075\u0063\u0065\u0052\u0061\u006e\u006b\u0042\u0069n\u0061\u0072\u0079\u0032"
	if _gagb == nil {
		return nil, _a.Error(_gb, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _cbf < 1 || _cbf > 4 {
		return nil, _a.Error(_gb, "\u006c\u0065\u0076\u0065\u006c\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020\u0073e\u0074\u0020\u007b\u0031\u002c\u0032\u002c\u0033\u002c\u0034\u007d")
	}
	if _gagb.Height <= 1 {
		return nil, _a.Errorf(_gb, "\u0073o\u0075\u0072c\u0065\u0020\u0068e\u0069\u0067\u0068\u0074\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0061t\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0027\u0032\u0027\u0020-\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027", _gagb.Height)
	}
	_dce = New(_gagb.Width/2, _gagb.Height/2)
	if _ed == nil {
		_ed = _faf()
	}
	_cbg := _fda(_gagb.RowStride, 2*_dce.RowStride)
	switch _cbf {
	case 1:
		_adc = _gba(_gagb, _dce, _cbf, _ed, _cbg)
	case 2:
		_adc = _cfg(_gagb, _dce, _cbf, _ed, _cbg)
	case 3:
		_adc = _gagf(_gagb, _dce, _cbf, _ed, _cbg)
	case 4:
		_adc = _cad(_gagb, _dce, _cbf, _ed, _cbg)
	}
	if _adc != nil {
		return nil, _adc
	}
	return _dce, nil
}
func _egcg(_dabed *Bitmap, _gbad *Bitmap, _cagd *Selection) (*Bitmap, error) {
	var (
		_bccg *Bitmap
		_baef error
	)
	_dabed, _baef = _gddb(_dabed, _gbad, _cagd, &_bccg)
	if _baef != nil {
		return nil, _baef
	}
	if _baef = _dabed.clearAll(); _baef != nil {
		return nil, _baef
	}
	var _fcadf SelectionValue
	for _efbg := 0; _efbg < _cagd.Height; _efbg++ {
		for _cafb := 0; _cafb < _cagd.Width; _cafb++ {
			_fcadf = _cagd.Data[_efbg][_cafb]
			if _fcadf == SelHit {
				if _baef = _dabed.RasterOperation(_cafb-_cagd.Cx, _efbg-_cagd.Cy, _gbad.Width, _gbad.Height, PixSrcOrDst, _bccg, 0, 0); _baef != nil {
					return nil, _baef
				}
			}
		}
	}
	return _dabed, nil
}
func (_cbgg *Bitmap) GetChocolateData() []byte {
	if _cbgg.Color == Vanilla {
		_cbgg.inverseData()
	}
	return _cbgg.Data
}

var _accf = [5]int{1, 2, 3, 0, 4}

func _bgeb(_ggag int) int {
	if _ggag < 0 {
		return -_ggag
	}
	return _ggag
}
func (_bbfa *Bitmap) setFourBytes(_deac int, _cffg uint32) error {
	if _deac+3 > len(_bbfa.Data)-1 {
		return _a.Errorf("\u0073\u0065\u0074F\u006f\u0075\u0072\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _deac)
	}
	_bbfa.Data[_deac] = byte((_cffg & 0xff000000) >> 24)
	_bbfa.Data[_deac+1] = byte((_cffg & 0xff0000) >> 16)
	_bbfa.Data[_deac+2] = byte((_cffg & 0xff00) >> 8)
	_bbfa.Data[_deac+3] = byte(_cffg & 0xff)
	return nil
}
func (_fbaa *BitmapsArray) GetBox(i int) (*_ac.Rectangle, error) {
	const _dcbf = "\u0042\u0069\u0074\u006dap\u0073\u0041\u0072\u0072\u0061\u0079\u002e\u0047\u0065\u0074\u0042\u006f\u0078"
	if _fbaa == nil {
		return nil, _a.Error(_dcbf, "p\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074m\u0061\u0070\u0073A\u0072r\u0061\u0079\u0027")
	}
	if i > len(_fbaa.Boxes)-1 {
		return nil, _a.Errorf(_dcbf, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _fbaa.Boxes[i], nil
}
func (_dde *Bitmap) Zero() bool {
	_cdg := _dde.Width / 8
	_ebeg := _dde.Width & 7
	var _bdb byte
	if _ebeg != 0 {
		_bdb = byte(0xff << uint(8-_ebeg))
	}
	var _dgdf, _bfa, _gcfd int
	for _bfa = 0; _bfa < _dde.Height; _bfa++ {
		_dgdf = _dde.RowStride * _bfa
		for _gcfd = 0; _gcfd < _cdg; _gcfd, _dgdf = _gcfd+1, _dgdf+1 {
			if _dde.Data[_dgdf] != 0 {
				return false
			}
		}
		if _ebeg > 0 {
			if _dde.Data[_dgdf]&_bdb != 0 {
				return false
			}
		}
	}
	return true
}
func TstASymbol(t *_f.T) *Bitmap {
	t.Helper()
	_eeebd := New(6, 6)
	_g.NoError(t, _eeebd.SetPixel(1, 0, 1))
	_g.NoError(t, _eeebd.SetPixel(2, 0, 1))
	_g.NoError(t, _eeebd.SetPixel(3, 0, 1))
	_g.NoError(t, _eeebd.SetPixel(4, 0, 1))
	_g.NoError(t, _eeebd.SetPixel(5, 1, 1))
	_g.NoError(t, _eeebd.SetPixel(1, 2, 1))
	_g.NoError(t, _eeebd.SetPixel(2, 2, 1))
	_g.NoError(t, _eeebd.SetPixel(3, 2, 1))
	_g.NoError(t, _eeebd.SetPixel(4, 2, 1))
	_g.NoError(t, _eeebd.SetPixel(5, 2, 1))
	_g.NoError(t, _eeebd.SetPixel(0, 3, 1))
	_g.NoError(t, _eeebd.SetPixel(5, 3, 1))
	_g.NoError(t, _eeebd.SetPixel(0, 4, 1))
	_g.NoError(t, _eeebd.SetPixel(5, 4, 1))
	_g.NoError(t, _eeebd.SetPixel(1, 5, 1))
	_g.NoError(t, _eeebd.SetPixel(2, 5, 1))
	_g.NoError(t, _eeebd.SetPixel(3, 5, 1))
	_g.NoError(t, _eeebd.SetPixel(4, 5, 1))
	_g.NoError(t, _eeebd.SetPixel(5, 5, 1))
	return _eeebd
}
func (_dccc *Bitmap) setTwoBytes(_fcba int, _dddc uint16) error {
	if _fcba+1 > len(_dccc.Data)-1 {
		return _a.Errorf("s\u0065\u0074\u0054\u0077\u006f\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", _fcba)
	}
	_dccc.Data[_fcba] = byte((_dddc & 0xff00) >> 8)
	_dccc.Data[_fcba+1] = byte(_dddc & 0xff)
	return nil
}
func (_fgdfe *BitmapsArray) GetBitmaps(i int) (*Bitmaps, error) {
	const _bfec = "\u0042\u0069\u0074ma\u0070\u0073\u0041\u0072\u0072\u0061\u0079\u002e\u0047\u0065\u0074\u0042\u0069\u0074\u006d\u0061\u0070\u0073"
	if _fgdfe == nil {
		return nil, _a.Error(_bfec, "p\u0072\u006f\u0076\u0069\u0064\u0065d\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074m\u0061\u0070\u0073A\u0072r\u0061\u0079\u0027")
	}
	if i > len(_fgdfe.Values)-1 {
		return nil, _a.Errorf(_bfec, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _fgdfe.Values[i], nil
}
func Dilate(d *Bitmap, s *Bitmap, sel *Selection) (*Bitmap, error) { return _egcg(d, s, sel) }

type MorphOperation int

func _cbcaa(_bdaa, _gcdc, _aeaf byte) byte { return (_bdaa &^ (_aeaf)) | (_gcdc & _aeaf) }
func (_aec *Bitmap) setPadBits(_cdgg int) {
	_gcfa := 8 - _aec.Width%8
	if _gcfa == 8 {
		return
	}
	_ecae := _aec.Width / 8
	_dcfb := _gded[_gcfa]
	if _cdgg == 0 {
		_dcfb ^= _dcfb
	}
	var _dbbd int
	for _cdba := 0; _cdba < _aec.Height; _cdba++ {
		_dbbd = _cdba*_aec.RowStride + _ecae
		if _cdgg == 0 {
			_aec.Data[_dbbd] &= _dcfb
		} else {
			_aec.Data[_dbbd] |= _dcfb
		}
	}
}
func _gbdf(_agae *_fc.Stack, _dedc, _gfec, _daaf, _dfdg, _bfeg int, _fabg *_ac.Rectangle) (_agddbd error) {
	const _ceeb = "\u0070\u0075\u0073\u0068\u0046\u0069\u006c\u006c\u0053\u0065\u0067m\u0065\u006e\u0074\u0042\u006f\u0075\u006e\u0064\u0069\u006eg\u0042\u006f\u0078"
	if _agae == nil {
		return _a.Error(_ceeb, "\u006ei\u006c \u0073\u0074\u0061\u0063\u006b \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	if _fabg == nil {
		return _a.Error(_ceeb, "\u0070\u0072\u006f\u0076i\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0069\u006da\u0067e\u002e\u0052\u0065\u0063\u0074\u0061\u006eg\u006c\u0065")
	}
	_fabg.Min.X = _fc.Min(_fabg.Min.X, _dedc)
	_fabg.Max.X = _fc.Max(_fabg.Max.X, _gfec)
	_fabg.Min.Y = _fc.Min(_fabg.Min.Y, _daaf)
	_fabg.Max.Y = _fc.Max(_fabg.Max.Y, _daaf)
	if !(_daaf+_dfdg >= 0 && _daaf+_dfdg <= _bfeg) {
		return nil
	}
	if _agae.Aux == nil {
		return _a.Error(_ceeb, "a\u0075x\u0053\u0074\u0061\u0063\u006b\u0020\u006e\u006ft\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	var _edae *fillSegment
	_egeb, _cddfa := _agae.Aux.Pop()
	if _cddfa {
		if _edae, _cddfa = _egeb.(*fillSegment); !_cddfa {
			return _a.Error(_ceeb, "a\u0075\u0078\u0053\u0074\u0061\u0063k\u0020\u0064\u0061\u0074\u0061\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0061 \u002a\u0066\u0069\u006c\u006c\u0053\u0065\u0067\u006d\u0065n\u0074")
		}
	} else {
		_edae = &fillSegment{}
	}
	_edae._eeee = _dedc
	_edae._fdddg = _gfec
	_edae._dabg = _daaf
	_edae._bege = _dfdg
	_agae.Push(_edae)
	return nil
}
func _bfcdc(_fafb *Bitmap, _bgec, _acab int, _bed, _bebg int, _dacd RasterOperator) {
	var (
		_dbceb bool
		_ffcb  bool
		_gfdb  int
		_fgga  int
		_bceg  int
		_fbcg  int
		_fcfd  bool
		_bcgb  byte
	)
	_badd := 8 - (_bgec & 7)
	_geaf := _gded[_badd]
	_bcef := _fafb.RowStride*_acab + (_bgec >> 3)
	if _bed < _badd {
		_dbceb = true
		_geaf &= _fcgba[8-_badd+_bed]
	}
	if !_dbceb {
		_gfdb = (_bed - _badd) >> 3
		if _gfdb != 0 {
			_ffcb = true
			_fgga = _bcef + 1
		}
	}
	_bceg = (_bgec + _bed) & 7
	if !(_dbceb || _bceg == 0) {
		_fcfd = true
		_bcgb = _fcgba[_bceg]
		_fbcg = _bcef + 1 + _gfdb
	}
	var _dfcb, _ccgba int
	switch _dacd {
	case PixClr:
		for _dfcb = 0; _dfcb < _bebg; _dfcb++ {
			_fafb.Data[_bcef] = _cbcaa(_fafb.Data[_bcef], 0x0, _geaf)
			_bcef += _fafb.RowStride
		}
		if _ffcb {
			for _dfcb = 0; _dfcb < _bebg; _dfcb++ {
				for _ccgba = 0; _ccgba < _gfdb; _ccgba++ {
					_fafb.Data[_fgga+_ccgba] = 0x0
				}
				_fgga += _fafb.RowStride
			}
		}
		if _fcfd {
			for _dfcb = 0; _dfcb < _bebg; _dfcb++ {
				_fafb.Data[_fbcg] = _cbcaa(_fafb.Data[_fbcg], 0x0, _bcgb)
				_fbcg += _fafb.RowStride
			}
		}
	case PixSet:
		for _dfcb = 0; _dfcb < _bebg; _dfcb++ {
			_fafb.Data[_bcef] = _cbcaa(_fafb.Data[_bcef], 0xff, _geaf)
			_bcef += _fafb.RowStride
		}
		if _ffcb {
			for _dfcb = 0; _dfcb < _bebg; _dfcb++ {
				for _ccgba = 0; _ccgba < _gfdb; _ccgba++ {
					_fafb.Data[_fgga+_ccgba] = 0xff
				}
				_fgga += _fafb.RowStride
			}
		}
		if _fcfd {
			for _dfcb = 0; _dfcb < _bebg; _dfcb++ {
				_fafb.Data[_fbcg] = _cbcaa(_fafb.Data[_fbcg], 0xff, _bcgb)
				_fbcg += _fafb.RowStride
			}
		}
	case PixNotDst:
		for _dfcb = 0; _dfcb < _bebg; _dfcb++ {
			_fafb.Data[_bcef] = _cbcaa(_fafb.Data[_bcef], ^_fafb.Data[_bcef], _geaf)
			_bcef += _fafb.RowStride
		}
		if _ffcb {
			for _dfcb = 0; _dfcb < _bebg; _dfcb++ {
				for _ccgba = 0; _ccgba < _gfdb; _ccgba++ {
					_fafb.Data[_fgga+_ccgba] = ^(_fafb.Data[_fgga+_ccgba])
				}
				_fgga += _fafb.RowStride
			}
		}
		if _fcfd {
			for _dfcb = 0; _dfcb < _bebg; _dfcb++ {
				_fafb.Data[_fbcg] = _cbcaa(_fafb.Data[_fbcg], ^_fafb.Data[_fbcg], _bcgb)
				_fbcg += _fafb.RowStride
			}
		}
	}
}
func _fge(_eeca uint, _degda byte) byte { return _degda >> _eeca << _eeca }
func _gege(_bff, _fecc, _efc *Bitmap) (*Bitmap, error) {
	const _agdb = "\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0078\u006f\u0072"
	if _fecc == nil {
		return nil, _a.Error(_agdb, "'\u0062\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _efc == nil {
		return nil, _a.Error(_agdb, "'\u0062\u0032\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _bff == _efc {
		return nil, _a.Error(_agdb, "'\u0064\u0027\u0020\u003d\u003d\u0020\u0027\u0062\u0032\u0027")
	}
	if !_fecc.SizesEqual(_efc) {
		_db.Log.Debug("\u0025s\u0020\u002d \u0042\u0069\u0074\u006da\u0070\u0020\u0027b\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074 e\u0071\u0075\u0061l\u0020\u0073i\u007a\u0065\u0020\u0077\u0069\u0074h\u0020\u0027b\u0032\u0027", _agdb)
	}
	var _abaf error
	if _bff, _abaf = _bggbg(_bff, _fecc); _abaf != nil {
		return nil, _a.Wrap(_abaf, _agdb, "\u0063\u0061n\u0027\u0074\u0020c\u0072\u0065\u0061\u0074\u0065\u0020\u0027\u0064\u0027")
	}
	if _abaf = _bff.RasterOperation(0, 0, _bff.Width, _bff.Height, PixSrcXorDst, _efc, 0, 0); _abaf != nil {
		return nil, _a.Wrap(_abaf, _agdb, "")
	}
	return _bff, nil
}

type Getter interface{ GetBitmap() *Bitmap }

func TstAddSymbol(t *_f.T, bms *Bitmaps, sym *Bitmap, x *int, y int, space int) {
	bms.AddBitmap(sym)
	_ffbd := _ac.Rect(*x, y, *x+sym.Width, y+sym.Height)
	bms.AddBox(&_ffbd)
	*x += sym.Width + space
}
func (_gbd *Bitmap) ThresholdPixelSum(thresh int, tab8 []int) (_bcac bool, _bag error) {
	const _bada = "\u0042i\u0074\u006d\u0061\u0070\u002e\u0054\u0068\u0072\u0065\u0073\u0068o\u006c\u0064\u0050\u0069\u0078\u0065\u006c\u0053\u0075\u006d"
	if tab8 == nil {
		tab8 = _cdgge()
	}
	_bdfdb := _gbd.Width >> 3
	_bdfe := _gbd.Width & 7
	_bdca := byte(0xff << uint(8-_bdfe))
	var (
		_dcdd, _gadg, _acg, _acgf int
		_adce                     byte
	)
	for _dcdd = 0; _dcdd < _gbd.Height; _dcdd++ {
		_acg = _gbd.RowStride * _dcdd
		for _gadg = 0; _gadg < _bdfdb; _gadg++ {
			_adce, _bag = _gbd.GetByte(_acg + _gadg)
			if _bag != nil {
				return false, _a.Wrap(_bag, _bada, "\u0066\u0075\u006c\u006c\u0042\u0079\u0074\u0065")
			}
			_acgf += tab8[_adce]
		}
		if _bdfe != 0 {
			_adce, _bag = _gbd.GetByte(_acg + _gadg)
			if _bag != nil {
				return false, _a.Wrap(_bag, _bada, "p\u0061\u0072\u0074\u0069\u0061\u006c\u0042\u0079\u0074\u0065")
			}
			_adce &= _bdca
			_acgf += tab8[_adce]
		}
		if _acgf > thresh {
			return true, nil
		}
	}
	return _bcac, nil
}
func (_fcef *Bitmap) clearAll() error {
	return _fcef.RasterOperation(0, 0, _fcef.Width, _fcef.Height, PixClr, nil, 0, 0)
}
func (_fec *Bitmap) SetPadBits(value int) { _fec.setPadBits(value) }
func _eecdf(_gaag *Bitmap, _aeef, _becg, _ddbd, _cabd int, _cacc RasterOperator, _cbdg *Bitmap, _faee, _edga int) error {
	var (
		_geecf       bool
		_aegf        bool
		_feebc       byte
		_gdeg        int
		_cbde        int
		_bbcc        int
		_cgdf        int
		_debg        bool
		_daff        int
		_cdd         int
		_fbc         int
		_gdab        bool
		_ecaf        byte
		_adcfa       int
		_gfbf        int
		_ecbg        int
		_deba        byte
		_ggge        int
		_bbdb        int
		_cgfae       uint
		_dfage       uint
		_acaa        byte
		_feff        shift
		_dcddd       bool
		_dagf        bool
		_agaf, _dbbb int
	)
	if _faee&7 != 0 {
		_bbdb = 8 - (_faee & 7)
	}
	if _aeef&7 != 0 {
		_cbde = 8 - (_aeef & 7)
	}
	if _bbdb == 0 && _cbde == 0 {
		_acaa = _gded[0]
	} else {
		if _cbde > _bbdb {
			_cgfae = uint(_cbde - _bbdb)
		} else {
			_cgfae = uint(8 - (_bbdb - _cbde))
		}
		_dfage = 8 - _cgfae
		_acaa = _gded[_cgfae]
	}
	if (_aeef & 7) != 0 {
		_geecf = true
		_gdeg = 8 - (_aeef & 7)
		_feebc = _gded[_gdeg]
		_bbcc = _gaag.RowStride*_becg + (_aeef >> 3)
		_cgdf = _cbdg.RowStride*_edga + (_faee >> 3)
		_ggge = 8 - (_faee & 7)
		if _gdeg > _ggge {
			_feff = _ccbce
			if _ddbd >= _bbdb {
				_dcddd = true
			}
		} else {
			_feff = _cdde
		}
	}
	if _ddbd < _gdeg {
		_aegf = true
		_feebc &= _fcgba[8-_gdeg+_ddbd]
	}
	if !_aegf {
		_daff = (_ddbd - _gdeg) >> 3
		if _daff != 0 {
			_debg = true
			_cdd = _gaag.RowStride*_becg + ((_aeef + _cbde) >> 3)
			_fbc = _cbdg.RowStride*_edga + ((_faee + _cbde) >> 3)
		}
	}
	_adcfa = (_aeef + _ddbd) & 7
	if !(_aegf || _adcfa == 0) {
		_gdab = true
		_ecaf = _fcgba[_adcfa]
		_gfbf = _gaag.RowStride*_becg + ((_aeef + _cbde) >> 3) + _daff
		_ecbg = _cbdg.RowStride*_edga + ((_faee + _cbde) >> 3) + _daff
		if _adcfa > int(_dfage) {
			_dagf = true
		}
	}
	switch _cacc {
	case PixSrc:
		if _geecf {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				if _feff == _ccbce {
					_deba = _cbdg.Data[_cgdf] << _cgfae
					if _dcddd {
						_deba = _cbcaa(_deba, _cbdg.Data[_cgdf+1]>>_dfage, _acaa)
					}
				} else {
					_deba = _cbdg.Data[_cgdf] >> _dfage
				}
				_gaag.Data[_bbcc] = _cbcaa(_gaag.Data[_bbcc], _deba, _feebc)
				_bbcc += _gaag.RowStride
				_cgdf += _cbdg.RowStride
			}
		}
		if _debg {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				for _dbbb = 0; _dbbb < _daff; _dbbb++ {
					_deba = _cbcaa(_cbdg.Data[_fbc+_dbbb]<<_cgfae, _cbdg.Data[_fbc+_dbbb+1]>>_dfage, _acaa)
					_gaag.Data[_cdd+_dbbb] = _deba
				}
				_cdd += _gaag.RowStride
				_fbc += _cbdg.RowStride
			}
		}
		if _gdab {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				_deba = _cbdg.Data[_ecbg] << _cgfae
				if _dagf {
					_deba = _cbcaa(_deba, _cbdg.Data[_ecbg+1]>>_dfage, _acaa)
				}
				_gaag.Data[_gfbf] = _cbcaa(_gaag.Data[_gfbf], _deba, _ecaf)
				_gfbf += _gaag.RowStride
				_ecbg += _cbdg.RowStride
			}
		}
	case PixNotSrc:
		if _geecf {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				if _feff == _ccbce {
					_deba = _cbdg.Data[_cgdf] << _cgfae
					if _dcddd {
						_deba = _cbcaa(_deba, _cbdg.Data[_cgdf+1]>>_dfage, _acaa)
					}
				} else {
					_deba = _cbdg.Data[_cgdf] >> _dfage
				}
				_gaag.Data[_bbcc] = _cbcaa(_gaag.Data[_bbcc], ^_deba, _feebc)
				_bbcc += _gaag.RowStride
				_cgdf += _cbdg.RowStride
			}
		}
		if _debg {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				for _dbbb = 0; _dbbb < _daff; _dbbb++ {
					_deba = _cbcaa(_cbdg.Data[_fbc+_dbbb]<<_cgfae, _cbdg.Data[_fbc+_dbbb+1]>>_dfage, _acaa)
					_gaag.Data[_cdd+_dbbb] = ^_deba
				}
				_cdd += _gaag.RowStride
				_fbc += _cbdg.RowStride
			}
		}
		if _gdab {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				_deba = _cbdg.Data[_ecbg] << _cgfae
				if _dagf {
					_deba = _cbcaa(_deba, _cbdg.Data[_ecbg+1]>>_dfage, _acaa)
				}
				_gaag.Data[_gfbf] = _cbcaa(_gaag.Data[_gfbf], ^_deba, _ecaf)
				_gfbf += _gaag.RowStride
				_ecbg += _cbdg.RowStride
			}
		}
	case PixSrcOrDst:
		if _geecf {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				if _feff == _ccbce {
					_deba = _cbdg.Data[_cgdf] << _cgfae
					if _dcddd {
						_deba = _cbcaa(_deba, _cbdg.Data[_cgdf+1]>>_dfage, _acaa)
					}
				} else {
					_deba = _cbdg.Data[_cgdf] >> _dfage
				}
				_gaag.Data[_bbcc] = _cbcaa(_gaag.Data[_bbcc], _deba|_gaag.Data[_bbcc], _feebc)
				_bbcc += _gaag.RowStride
				_cgdf += _cbdg.RowStride
			}
		}
		if _debg {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				for _dbbb = 0; _dbbb < _daff; _dbbb++ {
					_deba = _cbcaa(_cbdg.Data[_fbc+_dbbb]<<_cgfae, _cbdg.Data[_fbc+_dbbb+1]>>_dfage, _acaa)
					_gaag.Data[_cdd+_dbbb] |= _deba
				}
				_cdd += _gaag.RowStride
				_fbc += _cbdg.RowStride
			}
		}
		if _gdab {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				_deba = _cbdg.Data[_ecbg] << _cgfae
				if _dagf {
					_deba = _cbcaa(_deba, _cbdg.Data[_ecbg+1]>>_dfage, _acaa)
				}
				_gaag.Data[_gfbf] = _cbcaa(_gaag.Data[_gfbf], _deba|_gaag.Data[_gfbf], _ecaf)
				_gfbf += _gaag.RowStride
				_ecbg += _cbdg.RowStride
			}
		}
	case PixSrcAndDst:
		if _geecf {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				if _feff == _ccbce {
					_deba = _cbdg.Data[_cgdf] << _cgfae
					if _dcddd {
						_deba = _cbcaa(_deba, _cbdg.Data[_cgdf+1]>>_dfage, _acaa)
					}
				} else {
					_deba = _cbdg.Data[_cgdf] >> _dfage
				}
				_gaag.Data[_bbcc] = _cbcaa(_gaag.Data[_bbcc], _deba&_gaag.Data[_bbcc], _feebc)
				_bbcc += _gaag.RowStride
				_cgdf += _cbdg.RowStride
			}
		}
		if _debg {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				for _dbbb = 0; _dbbb < _daff; _dbbb++ {
					_deba = _cbcaa(_cbdg.Data[_fbc+_dbbb]<<_cgfae, _cbdg.Data[_fbc+_dbbb+1]>>_dfage, _acaa)
					_gaag.Data[_cdd+_dbbb] &= _deba
				}
				_cdd += _gaag.RowStride
				_fbc += _cbdg.RowStride
			}
		}
		if _gdab {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				_deba = _cbdg.Data[_ecbg] << _cgfae
				if _dagf {
					_deba = _cbcaa(_deba, _cbdg.Data[_ecbg+1]>>_dfage, _acaa)
				}
				_gaag.Data[_gfbf] = _cbcaa(_gaag.Data[_gfbf], _deba&_gaag.Data[_gfbf], _ecaf)
				_gfbf += _gaag.RowStride
				_ecbg += _cbdg.RowStride
			}
		}
	case PixSrcXorDst:
		if _geecf {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				if _feff == _ccbce {
					_deba = _cbdg.Data[_cgdf] << _cgfae
					if _dcddd {
						_deba = _cbcaa(_deba, _cbdg.Data[_cgdf+1]>>_dfage, _acaa)
					}
				} else {
					_deba = _cbdg.Data[_cgdf] >> _dfage
				}
				_gaag.Data[_bbcc] = _cbcaa(_gaag.Data[_bbcc], _deba^_gaag.Data[_bbcc], _feebc)
				_bbcc += _gaag.RowStride
				_cgdf += _cbdg.RowStride
			}
		}
		if _debg {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				for _dbbb = 0; _dbbb < _daff; _dbbb++ {
					_deba = _cbcaa(_cbdg.Data[_fbc+_dbbb]<<_cgfae, _cbdg.Data[_fbc+_dbbb+1]>>_dfage, _acaa)
					_gaag.Data[_cdd+_dbbb] ^= _deba
				}
				_cdd += _gaag.RowStride
				_fbc += _cbdg.RowStride
			}
		}
		if _gdab {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				_deba = _cbdg.Data[_ecbg] << _cgfae
				if _dagf {
					_deba = _cbcaa(_deba, _cbdg.Data[_ecbg+1]>>_dfage, _acaa)
				}
				_gaag.Data[_gfbf] = _cbcaa(_gaag.Data[_gfbf], _deba^_gaag.Data[_gfbf], _ecaf)
				_gfbf += _gaag.RowStride
				_ecbg += _cbdg.RowStride
			}
		}
	case PixNotSrcOrDst:
		if _geecf {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				if _feff == _ccbce {
					_deba = _cbdg.Data[_cgdf] << _cgfae
					if _dcddd {
						_deba = _cbcaa(_deba, _cbdg.Data[_cgdf+1]>>_dfage, _acaa)
					}
				} else {
					_deba = _cbdg.Data[_cgdf] >> _dfage
				}
				_gaag.Data[_bbcc] = _cbcaa(_gaag.Data[_bbcc], ^_deba|_gaag.Data[_bbcc], _feebc)
				_bbcc += _gaag.RowStride
				_cgdf += _cbdg.RowStride
			}
		}
		if _debg {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				for _dbbb = 0; _dbbb < _daff; _dbbb++ {
					_deba = _cbcaa(_cbdg.Data[_fbc+_dbbb]<<_cgfae, _cbdg.Data[_fbc+_dbbb+1]>>_dfage, _acaa)
					_gaag.Data[_cdd+_dbbb] |= ^_deba
				}
				_cdd += _gaag.RowStride
				_fbc += _cbdg.RowStride
			}
		}
		if _gdab {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				_deba = _cbdg.Data[_ecbg] << _cgfae
				if _dagf {
					_deba = _cbcaa(_deba, _cbdg.Data[_ecbg+1]>>_dfage, _acaa)
				}
				_gaag.Data[_gfbf] = _cbcaa(_gaag.Data[_gfbf], ^_deba|_gaag.Data[_gfbf], _ecaf)
				_gfbf += _gaag.RowStride
				_ecbg += _cbdg.RowStride
			}
		}
	case PixNotSrcAndDst:
		if _geecf {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				if _feff == _ccbce {
					_deba = _cbdg.Data[_cgdf] << _cgfae
					if _dcddd {
						_deba = _cbcaa(_deba, _cbdg.Data[_cgdf+1]>>_dfage, _acaa)
					}
				} else {
					_deba = _cbdg.Data[_cgdf] >> _dfage
				}
				_gaag.Data[_bbcc] = _cbcaa(_gaag.Data[_bbcc], ^_deba&_gaag.Data[_bbcc], _feebc)
				_bbcc += _gaag.RowStride
				_cgdf += _cbdg.RowStride
			}
		}
		if _debg {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				for _dbbb = 0; _dbbb < _daff; _dbbb++ {
					_deba = _cbcaa(_cbdg.Data[_fbc+_dbbb]<<_cgfae, _cbdg.Data[_fbc+_dbbb+1]>>_dfage, _acaa)
					_gaag.Data[_cdd+_dbbb] &= ^_deba
				}
				_cdd += _gaag.RowStride
				_fbc += _cbdg.RowStride
			}
		}
		if _gdab {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				_deba = _cbdg.Data[_ecbg] << _cgfae
				if _dagf {
					_deba = _cbcaa(_deba, _cbdg.Data[_ecbg+1]>>_dfage, _acaa)
				}
				_gaag.Data[_gfbf] = _cbcaa(_gaag.Data[_gfbf], ^_deba&_gaag.Data[_gfbf], _ecaf)
				_gfbf += _gaag.RowStride
				_ecbg += _cbdg.RowStride
			}
		}
	case PixSrcOrNotDst:
		if _geecf {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				if _feff == _ccbce {
					_deba = _cbdg.Data[_cgdf] << _cgfae
					if _dcddd {
						_deba = _cbcaa(_deba, _cbdg.Data[_cgdf+1]>>_dfage, _acaa)
					}
				} else {
					_deba = _cbdg.Data[_cgdf] >> _dfage
				}
				_gaag.Data[_bbcc] = _cbcaa(_gaag.Data[_bbcc], _deba|^_gaag.Data[_bbcc], _feebc)
				_bbcc += _gaag.RowStride
				_cgdf += _cbdg.RowStride
			}
		}
		if _debg {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				for _dbbb = 0; _dbbb < _daff; _dbbb++ {
					_deba = _cbcaa(_cbdg.Data[_fbc+_dbbb]<<_cgfae, _cbdg.Data[_fbc+_dbbb+1]>>_dfage, _acaa)
					_gaag.Data[_cdd+_dbbb] = _deba | ^_gaag.Data[_cdd+_dbbb]
				}
				_cdd += _gaag.RowStride
				_fbc += _cbdg.RowStride
			}
		}
		if _gdab {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				_deba = _cbdg.Data[_ecbg] << _cgfae
				if _dagf {
					_deba = _cbcaa(_deba, _cbdg.Data[_ecbg+1]>>_dfage, _acaa)
				}
				_gaag.Data[_gfbf] = _cbcaa(_gaag.Data[_gfbf], _deba|^_gaag.Data[_gfbf], _ecaf)
				_gfbf += _gaag.RowStride
				_ecbg += _cbdg.RowStride
			}
		}
	case PixSrcAndNotDst:
		if _geecf {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				if _feff == _ccbce {
					_deba = _cbdg.Data[_cgdf] << _cgfae
					if _dcddd {
						_deba = _cbcaa(_deba, _cbdg.Data[_cgdf+1]>>_dfage, _acaa)
					}
				} else {
					_deba = _cbdg.Data[_cgdf] >> _dfage
				}
				_gaag.Data[_bbcc] = _cbcaa(_gaag.Data[_bbcc], _deba&^_gaag.Data[_bbcc], _feebc)
				_bbcc += _gaag.RowStride
				_cgdf += _cbdg.RowStride
			}
		}
		if _debg {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				for _dbbb = 0; _dbbb < _daff; _dbbb++ {
					_deba = _cbcaa(_cbdg.Data[_fbc+_dbbb]<<_cgfae, _cbdg.Data[_fbc+_dbbb+1]>>_dfage, _acaa)
					_gaag.Data[_cdd+_dbbb] = _deba &^ _gaag.Data[_cdd+_dbbb]
				}
				_cdd += _gaag.RowStride
				_fbc += _cbdg.RowStride
			}
		}
		if _gdab {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				_deba = _cbdg.Data[_ecbg] << _cgfae
				if _dagf {
					_deba = _cbcaa(_deba, _cbdg.Data[_ecbg+1]>>_dfage, _acaa)
				}
				_gaag.Data[_gfbf] = _cbcaa(_gaag.Data[_gfbf], _deba&^_gaag.Data[_gfbf], _ecaf)
				_gfbf += _gaag.RowStride
				_ecbg += _cbdg.RowStride
			}
		}
	case PixNotPixSrcOrDst:
		if _geecf {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				if _feff == _ccbce {
					_deba = _cbdg.Data[_cgdf] << _cgfae
					if _dcddd {
						_deba = _cbcaa(_deba, _cbdg.Data[_cgdf+1]>>_dfage, _acaa)
					}
				} else {
					_deba = _cbdg.Data[_cgdf] >> _dfage
				}
				_gaag.Data[_bbcc] = _cbcaa(_gaag.Data[_bbcc], ^(_deba | _gaag.Data[_bbcc]), _feebc)
				_bbcc += _gaag.RowStride
				_cgdf += _cbdg.RowStride
			}
		}
		if _debg {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				for _dbbb = 0; _dbbb < _daff; _dbbb++ {
					_deba = _cbcaa(_cbdg.Data[_fbc+_dbbb]<<_cgfae, _cbdg.Data[_fbc+_dbbb+1]>>_dfage, _acaa)
					_gaag.Data[_cdd+_dbbb] = ^(_deba | _gaag.Data[_cdd+_dbbb])
				}
				_cdd += _gaag.RowStride
				_fbc += _cbdg.RowStride
			}
		}
		if _gdab {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				_deba = _cbdg.Data[_ecbg] << _cgfae
				if _dagf {
					_deba = _cbcaa(_deba, _cbdg.Data[_ecbg+1]>>_dfage, _acaa)
				}
				_gaag.Data[_gfbf] = _cbcaa(_gaag.Data[_gfbf], ^(_deba | _gaag.Data[_gfbf]), _ecaf)
				_gfbf += _gaag.RowStride
				_ecbg += _cbdg.RowStride
			}
		}
	case PixNotPixSrcAndDst:
		if _geecf {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				if _feff == _ccbce {
					_deba = _cbdg.Data[_cgdf] << _cgfae
					if _dcddd {
						_deba = _cbcaa(_deba, _cbdg.Data[_cgdf+1]>>_dfage, _acaa)
					}
				} else {
					_deba = _cbdg.Data[_cgdf] >> _dfage
				}
				_gaag.Data[_bbcc] = _cbcaa(_gaag.Data[_bbcc], ^(_deba & _gaag.Data[_bbcc]), _feebc)
				_bbcc += _gaag.RowStride
				_cgdf += _cbdg.RowStride
			}
		}
		if _debg {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				for _dbbb = 0; _dbbb < _daff; _dbbb++ {
					_deba = _cbcaa(_cbdg.Data[_fbc+_dbbb]<<_cgfae, _cbdg.Data[_fbc+_dbbb+1]>>_dfage, _acaa)
					_gaag.Data[_cdd+_dbbb] = ^(_deba & _gaag.Data[_cdd+_dbbb])
				}
				_cdd += _gaag.RowStride
				_fbc += _cbdg.RowStride
			}
		}
		if _gdab {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				_deba = _cbdg.Data[_ecbg] << _cgfae
				if _dagf {
					_deba = _cbcaa(_deba, _cbdg.Data[_ecbg+1]>>_dfage, _acaa)
				}
				_gaag.Data[_gfbf] = _cbcaa(_gaag.Data[_gfbf], ^(_deba & _gaag.Data[_gfbf]), _ecaf)
				_gfbf += _gaag.RowStride
				_ecbg += _cbdg.RowStride
			}
		}
	case PixNotPixSrcXorDst:
		if _geecf {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				if _feff == _ccbce {
					_deba = _cbdg.Data[_cgdf] << _cgfae
					if _dcddd {
						_deba = _cbcaa(_deba, _cbdg.Data[_cgdf+1]>>_dfage, _acaa)
					}
				} else {
					_deba = _cbdg.Data[_cgdf] >> _dfage
				}
				_gaag.Data[_bbcc] = _cbcaa(_gaag.Data[_bbcc], ^(_deba ^ _gaag.Data[_bbcc]), _feebc)
				_bbcc += _gaag.RowStride
				_cgdf += _cbdg.RowStride
			}
		}
		if _debg {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				for _dbbb = 0; _dbbb < _daff; _dbbb++ {
					_deba = _cbcaa(_cbdg.Data[_fbc+_dbbb]<<_cgfae, _cbdg.Data[_fbc+_dbbb+1]>>_dfage, _acaa)
					_gaag.Data[_cdd+_dbbb] = ^(_deba ^ _gaag.Data[_cdd+_dbbb])
				}
				_cdd += _gaag.RowStride
				_fbc += _cbdg.RowStride
			}
		}
		if _gdab {
			for _agaf = 0; _agaf < _cabd; _agaf++ {
				_deba = _cbdg.Data[_ecbg] << _cgfae
				if _dagf {
					_deba = _cbcaa(_deba, _cbdg.Data[_ecbg+1]>>_dfage, _acaa)
				}
				_gaag.Data[_gfbf] = _cbcaa(_gaag.Data[_gfbf], ^(_deba ^ _gaag.Data[_gfbf]), _ecaf)
				_gfbf += _gaag.RowStride
				_ecbg += _cbdg.RowStride
			}
		}
	default:
		_db.Log.Debug("\u004f\u0070e\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006e\u006f\u0074\u0020\u0070\u0065\u0072\u006d\u0069tt\u0065\u0064", _cacc)
		return _a.Error("\u0072a\u0073t\u0065\u0072\u004f\u0070\u0047e\u006e\u0065r\u0061\u006c\u004c\u006f\u0077", "\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065r\u0061\u0074\u0069\u006f\u006e\u0020\u006eo\u0074\u0020\u0070\u0065\u0072\u006d\u0069\u0074\u0074\u0065\u0064")
	}
	return nil
}

type MorphProcess struct {
	Operation MorphOperation
	Arguments []int
}

const (
	_ LocationFilter = iota
	LocSelectWidth
	LocSelectHeight
	LocSelectXVal
	LocSelectYVal
	LocSelectIfEither
	LocSelectIfBoth
)

func (_cgbb *Bitmap) clipRectangle(_daa, _ffde *_ac.Rectangle) (_fbf *Bitmap, _baa error) {
	const _eba = "\u0063\u006c\u0069\u0070\u0052\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065"
	if _daa == nil {
		return nil, _a.Error(_eba, "\u0070r\u006fv\u0069\u0064\u0065\u0064\u0020n\u0069\u006c \u0027\u0062\u006f\u0078\u0027")
	}
	_bgfc, _ebd := _cgbb.Width, _cgbb.Height
	_afe, _baa := ClipBoxToRectangle(_daa, _bgfc, _ebd)
	if _baa != nil {
		_db.Log.Warning("\u0027\u0062ox\u0027\u0020\u0064o\u0065\u0073\u006e\u0027t o\u0076er\u006c\u0061\u0070\u0020\u0062\u0069\u0074ma\u0070\u0020\u0027\u0062\u0027\u003a\u0020%\u0076", _baa)
		return nil, nil
	}
	_bgcb, _faae := _afe.Min.X, _afe.Min.Y
	_fgd, _fagc := _afe.Max.X-_afe.Min.X, _afe.Max.Y-_afe.Min.Y
	_fbf = New(_fgd, _fagc)
	_fbf.Text = _cgbb.Text
	if _baa = _fbf.RasterOperation(0, 0, _fgd, _fagc, PixSrc, _cgbb, _bgcb, _faae); _baa != nil {
		return nil, _a.Wrap(_baa, _eba, "")
	}
	if _ffde != nil {
		*_ffde = *_afe
	}
	return _fbf, nil
}
func _cfge(_ddfa *Bitmap, _cdbgf, _deedc, _eebdf, _cdff int, _ffcc RasterOperator, _ecbb *Bitmap, _ccgc, _eecf int) error {
	var (
		_ecgad         byte
		_afgf          int
		_fbaba         int
		_agfgd, _gggfg int
		_agdee, _fefa  int
	)
	_gdccb := _eebdf >> 3
	_fabd := _eebdf & 7
	if _fabd > 0 {
		_ecgad = _fcgba[_fabd]
	}
	_afgf = _ecbb.RowStride*_eecf + (_ccgc >> 3)
	_fbaba = _ddfa.RowStride*_deedc + (_cdbgf >> 3)
	switch _ffcc {
	case PixSrc:
		for _agdee = 0; _agdee < _cdff; _agdee++ {
			_agfgd = _afgf + _agdee*_ecbb.RowStride
			_gggfg = _fbaba + _agdee*_ddfa.RowStride
			for _fefa = 0; _fefa < _gdccb; _fefa++ {
				_ddfa.Data[_gggfg] = _ecbb.Data[_agfgd]
				_gggfg++
				_agfgd++
			}
			if _fabd > 0 {
				_ddfa.Data[_gggfg] = _cbcaa(_ddfa.Data[_gggfg], _ecbb.Data[_agfgd], _ecgad)
			}
		}
	case PixNotSrc:
		for _agdee = 0; _agdee < _cdff; _agdee++ {
			_agfgd = _afgf + _agdee*_ecbb.RowStride
			_gggfg = _fbaba + _agdee*_ddfa.RowStride
			for _fefa = 0; _fefa < _gdccb; _fefa++ {
				_ddfa.Data[_gggfg] = ^(_ecbb.Data[_agfgd])
				_gggfg++
				_agfgd++
			}
			if _fabd > 0 {
				_ddfa.Data[_gggfg] = _cbcaa(_ddfa.Data[_gggfg], ^_ecbb.Data[_agfgd], _ecgad)
			}
		}
	case PixSrcOrDst:
		for _agdee = 0; _agdee < _cdff; _agdee++ {
			_agfgd = _afgf + _agdee*_ecbb.RowStride
			_gggfg = _fbaba + _agdee*_ddfa.RowStride
			for _fefa = 0; _fefa < _gdccb; _fefa++ {
				_ddfa.Data[_gggfg] |= _ecbb.Data[_agfgd]
				_gggfg++
				_agfgd++
			}
			if _fabd > 0 {
				_ddfa.Data[_gggfg] = _cbcaa(_ddfa.Data[_gggfg], _ecbb.Data[_agfgd]|_ddfa.Data[_gggfg], _ecgad)
			}
		}
	case PixSrcAndDst:
		for _agdee = 0; _agdee < _cdff; _agdee++ {
			_agfgd = _afgf + _agdee*_ecbb.RowStride
			_gggfg = _fbaba + _agdee*_ddfa.RowStride
			for _fefa = 0; _fefa < _gdccb; _fefa++ {
				_ddfa.Data[_gggfg] &= _ecbb.Data[_agfgd]
				_gggfg++
				_agfgd++
			}
			if _fabd > 0 {
				_ddfa.Data[_gggfg] = _cbcaa(_ddfa.Data[_gggfg], _ecbb.Data[_agfgd]&_ddfa.Data[_gggfg], _ecgad)
			}
		}
	case PixSrcXorDst:
		for _agdee = 0; _agdee < _cdff; _agdee++ {
			_agfgd = _afgf + _agdee*_ecbb.RowStride
			_gggfg = _fbaba + _agdee*_ddfa.RowStride
			for _fefa = 0; _fefa < _gdccb; _fefa++ {
				_ddfa.Data[_gggfg] ^= _ecbb.Data[_agfgd]
				_gggfg++
				_agfgd++
			}
			if _fabd > 0 {
				_ddfa.Data[_gggfg] = _cbcaa(_ddfa.Data[_gggfg], _ecbb.Data[_agfgd]^_ddfa.Data[_gggfg], _ecgad)
			}
		}
	case PixNotSrcOrDst:
		for _agdee = 0; _agdee < _cdff; _agdee++ {
			_agfgd = _afgf + _agdee*_ecbb.RowStride
			_gggfg = _fbaba + _agdee*_ddfa.RowStride
			for _fefa = 0; _fefa < _gdccb; _fefa++ {
				_ddfa.Data[_gggfg] |= ^(_ecbb.Data[_agfgd])
				_gggfg++
				_agfgd++
			}
			if _fabd > 0 {
				_ddfa.Data[_gggfg] = _cbcaa(_ddfa.Data[_gggfg], ^(_ecbb.Data[_agfgd])|_ddfa.Data[_gggfg], _ecgad)
			}
		}
	case PixNotSrcAndDst:
		for _agdee = 0; _agdee < _cdff; _agdee++ {
			_agfgd = _afgf + _agdee*_ecbb.RowStride
			_gggfg = _fbaba + _agdee*_ddfa.RowStride
			for _fefa = 0; _fefa < _gdccb; _fefa++ {
				_ddfa.Data[_gggfg] &= ^(_ecbb.Data[_agfgd])
				_gggfg++
				_agfgd++
			}
			if _fabd > 0 {
				_ddfa.Data[_gggfg] = _cbcaa(_ddfa.Data[_gggfg], ^(_ecbb.Data[_agfgd])&_ddfa.Data[_gggfg], _ecgad)
			}
		}
	case PixSrcOrNotDst:
		for _agdee = 0; _agdee < _cdff; _agdee++ {
			_agfgd = _afgf + _agdee*_ecbb.RowStride
			_gggfg = _fbaba + _agdee*_ddfa.RowStride
			for _fefa = 0; _fefa < _gdccb; _fefa++ {
				_ddfa.Data[_gggfg] = _ecbb.Data[_agfgd] | ^(_ddfa.Data[_gggfg])
				_gggfg++
				_agfgd++
			}
			if _fabd > 0 {
				_ddfa.Data[_gggfg] = _cbcaa(_ddfa.Data[_gggfg], _ecbb.Data[_agfgd]|^(_ddfa.Data[_gggfg]), _ecgad)
			}
		}
	case PixSrcAndNotDst:
		for _agdee = 0; _agdee < _cdff; _agdee++ {
			_agfgd = _afgf + _agdee*_ecbb.RowStride
			_gggfg = _fbaba + _agdee*_ddfa.RowStride
			for _fefa = 0; _fefa < _gdccb; _fefa++ {
				_ddfa.Data[_gggfg] = _ecbb.Data[_agfgd] &^ (_ddfa.Data[_gggfg])
				_gggfg++
				_agfgd++
			}
			if _fabd > 0 {
				_ddfa.Data[_gggfg] = _cbcaa(_ddfa.Data[_gggfg], _ecbb.Data[_agfgd]&^(_ddfa.Data[_gggfg]), _ecgad)
			}
		}
	case PixNotPixSrcOrDst:
		for _agdee = 0; _agdee < _cdff; _agdee++ {
			_agfgd = _afgf + _agdee*_ecbb.RowStride
			_gggfg = _fbaba + _agdee*_ddfa.RowStride
			for _fefa = 0; _fefa < _gdccb; _fefa++ {
				_ddfa.Data[_gggfg] = ^(_ecbb.Data[_agfgd] | _ddfa.Data[_gggfg])
				_gggfg++
				_agfgd++
			}
			if _fabd > 0 {
				_ddfa.Data[_gggfg] = _cbcaa(_ddfa.Data[_gggfg], ^(_ecbb.Data[_agfgd] | _ddfa.Data[_gggfg]), _ecgad)
			}
		}
	case PixNotPixSrcAndDst:
		for _agdee = 0; _agdee < _cdff; _agdee++ {
			_agfgd = _afgf + _agdee*_ecbb.RowStride
			_gggfg = _fbaba + _agdee*_ddfa.RowStride
			for _fefa = 0; _fefa < _gdccb; _fefa++ {
				_ddfa.Data[_gggfg] = ^(_ecbb.Data[_agfgd] & _ddfa.Data[_gggfg])
				_gggfg++
				_agfgd++
			}
			if _fabd > 0 {
				_ddfa.Data[_gggfg] = _cbcaa(_ddfa.Data[_gggfg], ^(_ecbb.Data[_agfgd] & _ddfa.Data[_gggfg]), _ecgad)
			}
		}
	case PixNotPixSrcXorDst:
		for _agdee = 0; _agdee < _cdff; _agdee++ {
			_agfgd = _afgf + _agdee*_ecbb.RowStride
			_gggfg = _fbaba + _agdee*_ddfa.RowStride
			for _fefa = 0; _fefa < _gdccb; _fefa++ {
				_ddfa.Data[_gggfg] = ^(_ecbb.Data[_agfgd] ^ _ddfa.Data[_gggfg])
				_gggfg++
				_agfgd++
			}
			if _fabd > 0 {
				_ddfa.Data[_gggfg] = _cbcaa(_ddfa.Data[_gggfg], ^(_ecbb.Data[_agfgd] ^ _ddfa.Data[_gggfg]), _ecgad)
			}
		}
	default:
		_db.Log.Debug("\u0050\u0072ov\u0069\u0064\u0065d\u0020\u0069\u006e\u0076ali\u0064 r\u0061\u0073\u0074\u0065\u0072\u0020\u006fpe\u0072\u0061\u0074\u006f\u0072\u003a\u0020%\u0076", _ffcc)
		return _a.Error("\u0072\u0061\u0073\u0074er\u004f\u0070\u0042\u0079\u0074\u0065\u0041\u006c\u0069\u0067\u006e\u0065\u0064\u004co\u0077", "\u0069\u006e\u0076al\u0069\u0064\u0020\u0072\u0061\u0073\u0074\u0065\u0072\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u006f\u0072")
	}
	return nil
}
func (_ceff *Bitmap) nextOnPixel(_fbgc, _fcc int) (_bcaf _ac.Point, _gegf bool, _addb error) {
	const _fegg = "n\u0065\u0078\u0074\u004f\u006e\u0050\u0069\u0078\u0065\u006c"
	_bcaf, _gegf, _addb = _ceff.nextOnPixelLow(_ceff.Width, _ceff.Height, _ceff.RowStride, _fbgc, _fcc)
	if _addb != nil {
		return _bcaf, false, _a.Wrap(_addb, _fegg, "")
	}
	return _bcaf, _gegf, nil
}
func _gebb(_dgbe, _bgdf *Bitmap, _bfff, _bfad int) (*Bitmap, error) {
	const _eeafe = "d\u0069\u006c\u0061\u0074\u0065\u0042\u0072\u0069\u0063\u006b"
	if _bgdf == nil {
		_db.Log.Debug("\u0064\u0069\u006c\u0061\u0074\u0065\u0042\u0072\u0069\u0063k\u0020\u0073\u006f\u0075\u0072\u0063\u0065 \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
		return nil, _a.Error(_eeafe, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if _bfff < 1 || _bfad < 1 {
		return nil, _a.Error(_eeafe, "\u0068\u0053\u007a\u0069\u0065 \u0061\u006e\u0064\u0020\u0076\u0053\u0069\u007a\u0065\u0020\u0061\u0072\u0065 \u006e\u006f\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0031")
	}
	if _bfff == 1 && _bfad == 1 {
		_gbbg, _abbg := _bggbg(_dgbe, _bgdf)
		if _abbg != nil {
			return nil, _a.Wrap(_abbg, _eeafe, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u0026\u0026 \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _gbbg, nil
	}
	if _bfff == 1 || _bfad == 1 {
		_fgbbf := SelCreateBrick(_bfad, _bfff, _bfad/2, _bfff/2, SelHit)
		_beeb, _cabb := _egcg(_dgbe, _bgdf, _fgbbf)
		if _cabb != nil {
			return nil, _a.Wrap(_cabb, _eeafe, "\u0068s\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _beeb, nil
	}
	_edea := SelCreateBrick(1, _bfff, 0, _bfff/2, SelHit)
	_facca := SelCreateBrick(_bfad, 1, _bfad/2, 0, SelHit)
	_acag, _dbea := _egcg(nil, _bgdf, _edea)
	if _dbea != nil {
		return nil, _a.Wrap(_dbea, _eeafe, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	_dgbe, _dbea = _egcg(_dgbe, _acag, _facca)
	if _dbea != nil {
		return nil, _a.Wrap(_dbea, _eeafe, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	return _dgbe, nil
}
func (_gcdbcf *byWidth) Less(i, j int) bool { return _gcdbcf.Values[i].Width < _gcdbcf.Values[j].Width }
func (_egdf *Bitmap) connComponentsBB(_eaee int) (_agf *Boxes, _gcgf error) {
	const _agfd = "\u0042\u0069\u0074ma\u0070\u002e\u0063\u006f\u006e\u006e\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0042\u0042"
	if _eaee != 4 && _eaee != 8 {
		return nil, _a.Error(_agfd, "\u0063\u006f\u006e\u006e\u0065\u0063t\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u0027\u0034\u0027\u0020\u006fr\u0020\u0027\u0038\u0027")
	}
	if _egdf.Zero() {
		return &Boxes{}, nil
	}
	_egdf.setPadBits(0)
	_ebfc, _gcgf := _bggbg(nil, _egdf)
	if _gcgf != nil {
		return nil, _a.Wrap(_gcgf, _agfd, "\u0062\u006d\u0031")
	}
	_beaf := &_fc.Stack{}
	_beaf.Aux = &_fc.Stack{}
	_agf = &Boxes{}
	var (
		_faaf, _adad int
		_bgdd        _ac.Point
		_cac         bool
		_eeea        *_ac.Rectangle
	)
	for {
		if _bgdd, _cac, _gcgf = _ebfc.nextOnPixel(_adad, _faaf); _gcgf != nil {
			return nil, _a.Wrap(_gcgf, _agfd, "")
		}
		if !_cac {
			break
		}
		if _eeea, _gcgf = _bebaf(_ebfc, _beaf, _bgdd.X, _bgdd.Y, _eaee); _gcgf != nil {
			return nil, _a.Wrap(_gcgf, _agfd, "")
		}
		if _gcgf = _agf.Add(_eeea); _gcgf != nil {
			return nil, _a.Wrap(_gcgf, _agfd, "")
		}
		_adad = _bgdd.X
		_faaf = _bgdd.Y
	}
	return _agf, nil
}

const (
	_ccbce shift = iota
	_cdde
)

func _dcba(_afgcd *Bitmap, _adbc *Bitmap, _aabga int) (_adegc error) {
	const _effa = "\u0073\u0065\u0065\u0064\u0066\u0069\u006c\u006c\u0042\u0069\u006e\u0061r\u0079\u004c\u006f\u0077"
	_befa := _fda(_afgcd.Height, _adbc.Height)
	_bdgbf := _fda(_afgcd.RowStride, _adbc.RowStride)
	switch _aabga {
	case 4:
		_adegc = _bcda(_afgcd, _adbc, _befa, _bdgbf)
	case 8:
		_adegc = _dfdfd(_afgcd, _adbc, _befa, _bdgbf)
	default:
		return _a.Errorf(_effa, "\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0034\u0020\u006fr\u0020\u0038\u0020\u002d\u0020i\u0073\u003a \u0027\u0025\u0064\u0027", _aabga)
	}
	if _adegc != nil {
		return _a.Wrap(_adegc, _effa, "")
	}
	return nil
}

const (
	AsymmetricMorphBC BoundaryCondition = iota
	SymmetricMorphBC
)

var _ _e.Interface = &ClassedPoints{}

func _gabf(_dgbc, _fdge *Bitmap, _agddb CombinationOperator) *Bitmap {
	_facf := New(_dgbc.Width, _dgbc.Height)
	for _baeb := 0; _baeb < len(_facf.Data); _baeb++ {
		_facf.Data[_baeb] = _bgbg(_dgbc.Data[_baeb], _fdge.Data[_baeb], _agddb)
	}
	return _facf
}

const (
	_ SizeComparison = iota
	SizeSelectIfLT
	SizeSelectIfGT
	SizeSelectIfLTE
	SizeSelectIfGTE
	SizeSelectIfEQ
)

func (_effda *ClassedPoints) GetIntXByClass(i int) (int, error) {
	const _gace = "\u0043\u006c\u0061\u0073s\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047e\u0074I\u006e\u0074\u0059\u0042\u0079\u0043\u006ca\u0073\u0073"
	if i >= _effda.IntSlice.Size() {
		return 0, _a.Errorf(_gace, "\u0069\u003a\u0020\u0027\u0025\u0064\u0027 \u0069\u0073\u0020o\u0075\u0074\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065", i)
	}
	return int(_effda.XAtIndex(i)), nil
}
func (_bbee Points) GetIntY(i int) (int, error) {
	if i >= len(_bbee) {
		return 0, _a.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065t\u0049\u006e\u0074\u0059", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return int(_bbee[i].Y), nil
}
func _bcda(_bgac, _eega *Bitmap, _aeac, _gdac int) (_ddaf error) {
	const _dgbdga = "\u0073e\u0065d\u0066\u0069\u006c\u006c\u0042i\u006e\u0061r\u0079\u004c\u006f\u0077\u0034"
	var (
		_cadc, _begc, _abdf, _eadd                         int
		_eedg, _agggg, _acga, _gcfcc, _abgb, _bfgd, _cfgcb byte
	)
	for _cadc = 0; _cadc < _aeac; _cadc++ {
		_abdf = _cadc * _bgac.RowStride
		_eadd = _cadc * _eega.RowStride
		for _begc = 0; _begc < _gdac; _begc++ {
			_eedg, _ddaf = _bgac.GetByte(_abdf + _begc)
			if _ddaf != nil {
				return _a.Wrap(_ddaf, _dgbdga, "\u0066i\u0072\u0073\u0074\u0020\u0067\u0065t")
			}
			_agggg, _ddaf = _eega.GetByte(_eadd + _begc)
			if _ddaf != nil {
				return _a.Wrap(_ddaf, _dgbdga, "\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0067\u0065\u0074")
			}
			if _cadc > 0 {
				_acga, _ddaf = _bgac.GetByte(_abdf - _bgac.RowStride + _begc)
				if _ddaf != nil {
					return _a.Wrap(_ddaf, _dgbdga, "\u0069\u0020\u003e \u0030")
				}
				_eedg |= _acga
			}
			if _begc > 0 {
				_gcfcc, _ddaf = _bgac.GetByte(_abdf + _begc - 1)
				if _ddaf != nil {
					return _a.Wrap(_ddaf, _dgbdga, "\u006a\u0020\u003e \u0030")
				}
				_eedg |= _gcfcc << 7
			}
			_eedg &= _agggg
			if _eedg == 0 || (^_eedg) == 0 {
				if _ddaf = _bgac.SetByte(_abdf+_begc, _eedg); _ddaf != nil {
					return _a.Wrap(_ddaf, _dgbdga, "b\u0074\u0020\u003d\u003d 0\u0020|\u007c\u0020\u0028\u005e\u0062t\u0029\u0020\u003d\u003d\u0020\u0030")
				}
				continue
			}
			for {
				_cfgcb = _eedg
				_eedg = (_eedg | (_eedg >> 1) | (_eedg << 1)) & _agggg
				if (_eedg ^ _cfgcb) == 0 {
					if _ddaf = _bgac.SetByte(_abdf+_begc, _eedg); _ddaf != nil {
						return _a.Wrap(_ddaf, _dgbdga, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0070\u0072\u0065\u0076 \u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	for _cadc = _aeac - 1; _cadc >= 0; _cadc-- {
		_abdf = _cadc * _bgac.RowStride
		_eadd = _cadc * _eega.RowStride
		for _begc = _gdac - 1; _begc >= 0; _begc-- {
			if _eedg, _ddaf = _bgac.GetByte(_abdf + _begc); _ddaf != nil {
				return _a.Wrap(_ddaf, _dgbdga, "\u0072\u0065\u0076\u0065\u0072\u0073\u0065\u0020\u0066\u0069\u0072\u0073t\u0020\u0067\u0065\u0074")
			}
			if _agggg, _ddaf = _eega.GetByte(_eadd + _begc); _ddaf != nil {
				return _a.Wrap(_ddaf, _dgbdga, "r\u0065\u0076\u0065\u0072se\u0020g\u0065\u0074\u0020\u006d\u0061s\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _cadc < _aeac-1 {
				if _abgb, _ddaf = _bgac.GetByte(_abdf + _bgac.RowStride + _begc); _ddaf != nil {
					return _a.Wrap(_ddaf, _dgbdga, "\u0072\u0065v\u0065\u0072\u0073e\u0020\u0069\u0020\u003c\u0020\u0068\u0020\u002d\u0031")
				}
				_eedg |= _abgb
			}
			if _begc < _gdac-1 {
				if _bfgd, _ddaf = _bgac.GetByte(_abdf + _begc + 1); _ddaf != nil {
					return _a.Wrap(_ddaf, _dgbdga, "\u0072\u0065\u0076\u0065rs\u0065\u0020\u006a\u0020\u003c\u0020\u0077\u0070\u006c\u0020\u002d\u0020\u0031")
				}
				_eedg |= _bfgd >> 7
			}
			_eedg &= _agggg
			if _eedg == 0 || (^_eedg) == 0 {
				if _ddaf = _bgac.SetByte(_abdf+_begc, _eedg); _ddaf != nil {
					return _a.Wrap(_ddaf, _dgbdga, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006d\u0061\u0073k\u0065\u0064\u0020\u0062\u0079\u0074\u0065\u0020\u0066\u0061i\u006c\u0065\u0064")
				}
				continue
			}
			for {
				_cfgcb = _eedg
				_eedg = (_eedg | (_eedg >> 1) | (_eedg << 1)) & _agggg
				if (_eedg ^ _cfgcb) == 0 {
					if _ddaf = _bgac.SetByte(_abdf+_begc, _eedg); _ddaf != nil {
						return _a.Wrap(_ddaf, _dgbdga, "\u0072e\u0076\u0065\u0072\u0073e\u0020\u0073\u0065\u0074\u0074i\u006eg\u0020p\u0072\u0065\u0076\u0020\u0062\u0079\u0074e")
					}
					break
				}
			}
		}
	}
	return nil
}
func _cfc(_ebe *Bitmap, _ee *Bitmap, _ced int) (_bgf error) {
	const _aed = "e\u0078\u0070\u0061\u006edB\u0069n\u0061\u0072\u0079\u0050\u006fw\u0065\u0072\u0032\u004c\u006f\u0077"
	switch _ced {
	case 2:
		_bgf = _ga(_ebe, _ee)
	case 4:
		_bgf = _be(_ebe, _ee)
	case 8:
		_bgf = _fd(_ebe, _ee)
	default:
		return _a.Error(_aed, "\u0065\u0078p\u0061\u006e\u0073\u0069o\u006e\u0020f\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006ft\u0020\u0069\u006e\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d\u0020r\u0061\u006e\u0067\u0065")
	}
	if _bgf != nil {
		_bgf = _a.Wrap(_bgf, _aed, "")
	}
	return _bgf
}
func (_degdb *ClassedPoints) GetIntYByClass(i int) (int, error) {
	const _edaa = "\u0043\u006c\u0061\u0073s\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047e\u0074I\u006e\u0074\u0059\u0042\u0079\u0043\u006ca\u0073\u0073"
	if i >= _degdb.IntSlice.Size() {
		return 0, _a.Errorf(_edaa, "\u0069\u003a\u0020\u0027\u0025\u0064\u0027 \u0069\u0073\u0020o\u0075\u0074\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065", i)
	}
	return int(_degdb.YAtIndex(i)), nil
}
func _gcga(_bcbb, _bcgf *Bitmap, _gab, _gge, _gdd, _bgcf, _gdgd, _decbg, _ggdc, _eebd int, _egfg CombinationOperator, _aecc int) error {
	var _gbcf int
	_adb := func() {
		_gbcf++
		_gdd += _bcgf.RowStride
		_bgcf += _bcbb.RowStride
		_gdgd += _bcbb.RowStride
	}
	for _gbcf = _gab; _gbcf < _gge; _adb() {
		var _dfe uint16
		_bcdd := _gdd
		for _badf := _bgcf; _badf <= _gdgd; _badf++ {
			_gce, _beaa := _bcgf.GetByte(_bcdd)
			if _beaa != nil {
				return _beaa
			}
			_bdgd, _beaa := _bcbb.GetByte(_badf)
			if _beaa != nil {
				return _beaa
			}
			_dfe = (_dfe | (uint16(_bdgd) & 0xff)) << uint(_eebd)
			_bdgd = byte(_dfe >> 8)
			if _beaa = _bcgf.SetByte(_bcdd, _bgbg(_gce, _bdgd, _egfg)); _beaa != nil {
				return _beaa
			}
			_bcdd++
			_dfe <<= uint(_ggdc)
			if _badf == _gdgd {
				_bdgd = byte(_dfe >> (8 - uint8(_eebd)))
				if _aecc != 0 {
					_bdgd = _fge(uint(8+_decbg), _bdgd)
				}
				_gce, _beaa = _bcgf.GetByte(_bcdd)
				if _beaa != nil {
					return _beaa
				}
				if _beaa = _bcgf.SetByte(_bcdd, _bgbg(_gce, _bdgd, _egfg)); _beaa != nil {
					return _beaa
				}
			}
		}
	}
	return nil
}
func (_baac Points) Size() int   { return len(_baac) }
func (_geae *Bitmaps) Size() int { return len(_geae.Values) }
func _fcee(_edgf, _feccd *Bitmap, _cccbg, _ebdd int) (*Bitmap, error) {
	const _ebfb = "\u0065\u0072\u006f\u0064\u0065\u0042\u0072\u0069\u0063\u006b"
	if _feccd == nil {
		return nil, _a.Error(_ebfb, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _cccbg < 1 || _ebdd < 1 {
		return nil, _a.Error(_ebfb, "\u0068\u0073\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0073\u0069\u007a\u0065\u0020\u0061\u0072e\u0020\u006e\u006f\u0074\u0020\u0067\u0072e\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006fr\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0031")
	}
	if _cccbg == 1 && _ebdd == 1 {
		_geadc, _bdce := _bggbg(_edgf, _feccd)
		if _bdce != nil {
			return nil, _a.Wrap(_bdce, _ebfb, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u0026\u0026 \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _geadc, nil
	}
	if _cccbg == 1 || _ebdd == 1 {
		_gfbd := SelCreateBrick(_ebdd, _cccbg, _ebdd/2, _cccbg/2, SelHit)
		_ddge, _abab := _abbc(_edgf, _feccd, _gfbd)
		if _abab != nil {
			return nil, _a.Wrap(_abab, _ebfb, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _ddge, nil
	}
	_cgfa := SelCreateBrick(1, _cccbg, 0, _cccbg/2, SelHit)
	_bdgb := SelCreateBrick(_ebdd, 1, _ebdd/2, 0, SelHit)
	_ebbb, _fefg := _abbc(nil, _feccd, _cgfa)
	if _fefg != nil {
		return nil, _a.Wrap(_fefg, _ebfb, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	_edgf, _fefg = _abbc(_edgf, _ebbb, _bdgb)
	if _fefg != nil {
		return nil, _a.Wrap(_fefg, _ebfb, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	return _edgf, nil
}
func (_caff *Boxes) Get(i int) (*_ac.Rectangle, error) {
	const _cedd = "\u0042o\u0078\u0065\u0073\u002e\u0047\u0065t"
	if _caff == nil {
		return nil, _a.Error(_cedd, "\u0027\u0042\u006f\u0078es\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if i > len(*_caff)-1 {
		return nil, _a.Errorf(_cedd, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return (*_caff)[i], nil
}
func _bbg(_fbg *Bitmap, _dg int) (*Bitmap, error) {
	const _dff = "\u0065x\u0070a\u006e\u0064\u0042\u0069\u006ea\u0072\u0079P\u006f\u0077\u0065\u0072\u0032"
	if _fbg == nil {
		return nil, _a.Error(_dff, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _dg == 1 {
		return _bggbg(nil, _fbg)
	}
	if _dg != 2 && _dg != 4 && _dg != 8 {
		return nil, _a.Error(_dff, "\u0066\u0061\u0063t\u006f\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0069n\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_cdb := _dg * _fbg.Width
	_gfc := _dg * _fbg.Height
	_dbg := New(_cdb, _gfc)
	var _gef error
	switch _dg {
	case 2:
		_gef = _ga(_dbg, _fbg)
	case 4:
		_gef = _be(_dbg, _fbg)
	case 8:
		_gef = _fd(_dbg, _fbg)
	}
	if _gef != nil {
		return nil, _a.Wrap(_gef, _dff, "")
	}
	return _dbg, nil
}
func (_fbe *Bitmap) removeBorderGeneral(_eacc, _gbgfc, _gaefc, _eef int) (*Bitmap, error) {
	const _bee = "\u0072\u0065\u006d\u006fve\u0042\u006f\u0072\u0064\u0065\u0072\u0047\u0065\u006e\u0065\u0072\u0061\u006c"
	if _eacc < 0 || _gbgfc < 0 || _gaefc < 0 || _eef < 0 {
		return nil, _a.Error(_bee, "\u006e\u0065g\u0061\u0074\u0069\u0076\u0065\u0020\u0062\u0072\u006f\u0064\u0065\u0072\u0020\u0072\u0065\u006d\u006f\u0076\u0065\u0020\u0076\u0061lu\u0065\u0073")
	}
	_fbb, _fgdf := _fbe.Width, _fbe.Height
	_fgda := _fbb - _eacc - _gbgfc
	_gcfda := _fgdf - _gaefc - _eef
	if _fgda <= 0 {
		return nil, _a.Errorf(_bee, "w\u0069\u0064\u0074\u0068: \u0025d\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u003e\u0020\u0030", _fgda)
	}
	if _gcfda <= 0 {
		return nil, _a.Errorf(_bee, "\u0068\u0065\u0069\u0067ht\u003a\u0020\u0025\u0064\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u003e \u0030", _gcfda)
	}
	_bbe := New(_fgda, _gcfda)
	_bbe.Color = _fbe.Color
	_aee := _bbe.RasterOperation(0, 0, _fgda, _gcfda, PixSrc, _fbe, _eacc, _gaefc)
	if _aee != nil {
		return nil, _a.Wrap(_aee, _bee, "")
	}
	return _bbe, nil
}
func (_fgdb MorphProcess) verify(_dcccb int, _ccf, _bdbe *int) error {
	const _age = "\u004d\u006f\u0072\u0070hP\u0072\u006f\u0063\u0065\u0073\u0073\u002e\u0076\u0065\u0072\u0069\u0066\u0079"
	switch _fgdb.Operation {
	case MopDilation, MopErosion, MopOpening, MopClosing:
		if len(_fgdb.Arguments) != 2 {
			return _a.Error(_age, "\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0064\u0027\u002c\u0020\u0027\u0065\u0027\u002c \u0027\u006f\u0027\u002c\u0020\u0027\u0063\u0027\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0032\u0020\u0061r\u0067\u0075\u006d\u0065\u006et\u0073")
		}
		_ddff, _ggfgg := _fgdb.getWidthHeight()
		if _ddff <= 0 || _ggfgg <= 0 {
			return _a.Error(_age, "O\u0070er\u0061t\u0069o\u006e\u003a\u0020\u0027\u0064'\u002c\u0020\u0027e\u0027\u002c\u0020\u0027\u006f'\u002c\u0020\u0027c\u0027\u0020\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0062\u006f\u0074h w\u0069\u0064\u0074\u0068\u0020\u0061n\u0064\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0074\u006f\u0020b\u0065 \u003e\u003d\u0020\u0030")
		}
	case MopRankBinaryReduction:
		_dbdcg := len(_fgdb.Arguments)
		*_ccf += _dbdcg
		if _dbdcg < 1 || _dbdcg > 4 {
			return _a.Error(_age, "\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0072\u0027\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061s\u0074\u0020\u0031\u0020\u0061\u006e\u0064\u0020\u0061\u0074\u0020\u006d\u006fs\u0074\u0020\u0034\u0020\u0061\u0072g\u0075\u006d\u0065n\u0074\u0073")
		}
		for _cbba := 0; _cbba < _dbdcg; _cbba++ {
			if _fgdb.Arguments[_cbba] < 1 || _fgdb.Arguments[_cbba] > 4 {
				return _a.Error(_age, "\u0052\u0061\u006e\u006b\u0042\u0069n\u0061\u0072\u0079\u0052\u0065\u0064\u0075\u0063\u0074\u0069\u006f\u006e\u0020\u006c\u0065\u0076\u0065\u006c\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 \u00280\u002c\u0020\u0034\u003e")
			}
		}
	case MopReplicativeBinaryExpansion:
		if len(_fgdb.Arguments) == 0 {
			return _a.Error(_age, "\u0052\u0065\u0070\u006c\u0069\u0063\u0061\u0074i\u0076\u0065\u0042in\u0061\u0072\u0079\u0045\u0078\u0070a\u006e\u0073\u0069\u006f\u006e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020o\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		}
		_bdcc := _fgdb.Arguments[0]
		if _bdcc != 2 && _bdcc != 4 && _bdcc != 8 {
			return _a.Error(_age, "R\u0065\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u0076\u0065\u0042\u0069\u006e\u0061\u0072\u0079\u0045\u0078\u0070\u0061\u006e\u0073\u0069\u006f\u006e\u0020m\u0075s\u0074\u0020\u0062\u0065 \u006f\u0066 \u0066\u0061\u0063\u0074\u006f\u0072\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d")
		}
		*_ccf -= _accf[_bdcc/4]
	case MopAddBorder:
		if len(_fgdb.Arguments) == 0 {
			return _a.Error(_age, "\u0041\u0064\u0064B\u006f\u0072\u0064\u0065r\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_cgee := _fgdb.Arguments[0]
		if _dcccb > 0 {
			return _a.Error(_age, "\u0041\u0064\u0064\u0042\u006f\u0072\u0064\u0065\u0072\u0020\u006d\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020f\u0069\u0072\u0073\u0074\u0020\u006d\u006f\u0072\u0070\u0068\u0020\u0070\u0072o\u0063\u0065\u0073\u0073")
		}
		if _cgee < 1 {
			return _a.Error(_age, "\u0041\u0064\u0064\u0042o\u0072\u0064\u0065\u0072\u0020\u0076\u0061\u006c\u0075\u0065 \u006co\u0077\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u0030")
		}
		*_bdbe = _cgee
	}
	return nil
}
func (_geec *Boxes) Add(box *_ac.Rectangle) error {
	if _geec == nil {
		return _a.Error("\u0042o\u0078\u0065\u0073\u002e\u0041\u0064d", "\u0027\u0042\u006f\u0078es\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	*_geec = append(*_geec, box)
	return nil
}

var _afdc [256]uint8

func (_dbff *Points) AddPoint(x, y float32) { *_dbff = append(*_dbff, Point{x, y}) }
func TstESymbol(t *_f.T, scale ...int) *Bitmap {
	_bdfa, _bfegb := NewWithData(4, 5, []byte{0xF0, 0x80, 0xE0, 0x80, 0xF0})
	_g.NoError(t, _bfegb)
	return TstGetScaledSymbol(t, _bdfa, scale...)
}
func (_cce *Bitmap) ClipRectangle(box *_ac.Rectangle) (_ffb *Bitmap, _gbgf *_ac.Rectangle, _egb error) {
	const _gad = "\u0043\u006c\u0069\u0070\u0052\u0065\u0063\u0074\u0061\u006e\u0067\u006c\u0065"
	if box == nil {
		return nil, nil, _a.Error(_gad, "\u0062o\u0078 \u0069\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	_fca, _fgbd := _cce.Width, _cce.Height
	_dbe := _ac.Rect(0, 0, _fca, _fgbd)
	if !box.Overlaps(_dbe) {
		return nil, nil, _a.Error(_gad, "b\u006f\u0078\u0020\u0064oe\u0073n\u0027\u0074\u0020\u006f\u0076e\u0072\u006c\u0061\u0070\u0020\u0062")
	}
	_bac := box.Intersect(_dbe)
	_dcd, _ccb := _bac.Min.X, _bac.Min.Y
	_egfe, _gggb := _bac.Dx(), _bac.Dy()
	_ffb = New(_egfe, _gggb)
	_ffb.Text = _cce.Text
	if _egb = _ffb.RasterOperation(0, 0, _egfe, _gggb, PixSrc, _cce, _dcd, _ccb); _egb != nil {
		return nil, nil, _a.Wrap(_egb, _gad, "\u0050\u0069\u0078\u0053\u0072\u0063\u0020\u0074\u006f\u0020\u0063\u006ci\u0070\u0070\u0065\u0064")
	}
	_gbgf = &_bac
	return _ffb, _gbgf, nil
}
func _ffbc(_dcca *Bitmap, _babda *_fc.Stack, _aaec, _bgbc int) (_fdbg *_ac.Rectangle, _fcacf error) {
	const _bagg = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _dcca == nil {
		return nil, _a.Error(_bagg, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _babda == nil {
		return nil, _a.Error(_bagg, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	_dfda, _ffca := _dcca.Width, _dcca.Height
	_bdcf := _dfda - 1
	_ddfad := _ffca - 1
	if _aaec < 0 || _aaec > _bdcf || _bgbc < 0 || _bgbc > _ddfad || !_dcca.GetPixel(_aaec, _bgbc) {
		return nil, nil
	}
	var _agfe *_ac.Rectangle
	_agfe, _fcacf = Rect(100000, 100000, 0, 0)
	if _fcacf != nil {
		return nil, _a.Wrap(_fcacf, _bagg, "")
	}
	if _fcacf = _gbdf(_babda, _aaec, _aaec, _bgbc, 1, _ddfad, _agfe); _fcacf != nil {
		return nil, _a.Wrap(_fcacf, _bagg, "\u0069\u006e\u0069t\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	if _fcacf = _gbdf(_babda, _aaec, _aaec, _bgbc+1, -1, _ddfad, _agfe); _fcacf != nil {
		return nil, _a.Wrap(_fcacf, _bagg, "\u0032\u006ed\u0020\u0069\u006ei\u0074\u0069\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
	}
	_agfe.Min.X, _agfe.Max.X = _aaec, _aaec
	_agfe.Min.Y, _agfe.Max.Y = _bgbc, _bgbc
	var (
		_bddf *fillSegment
		_fdf  int
	)
	for _babda.Len() > 0 {
		if _bddf, _fcacf = _dabec(_babda); _fcacf != nil {
			return nil, _a.Wrap(_fcacf, _bagg, "")
		}
		_bgbc = _bddf._dabg
		for _aaec = _bddf._eeee; _aaec >= 0 && _dcca.GetPixel(_aaec, _bgbc); _aaec-- {
			if _fcacf = _dcca.SetPixel(_aaec, _bgbc, 0); _fcacf != nil {
				return nil, _a.Wrap(_fcacf, _bagg, "")
			}
		}
		if _aaec >= _bddf._eeee {
			for _aaec++; _aaec <= _bddf._fdddg && _aaec <= _bdcf && !_dcca.GetPixel(_aaec, _bgbc); _aaec++ {
			}
			_fdf = _aaec
			if !(_aaec <= _bddf._fdddg && _aaec <= _bdcf) {
				continue
			}
		} else {
			_fdf = _aaec + 1
			if _fdf < _bddf._eeee-1 {
				if _fcacf = _gbdf(_babda, _fdf, _bddf._eeee-1, _bddf._dabg, -_bddf._bege, _ddfad, _agfe); _fcacf != nil {
					return nil, _a.Wrap(_fcacf, _bagg, "\u006c\u0065\u0061\u006b\u0020\u006f\u006e\u0020\u006c\u0065\u0066\u0074 \u0073\u0069\u0064\u0065")
				}
			}
			_aaec = _bddf._eeee + 1
		}
		for {
			for ; _aaec <= _bdcf && _dcca.GetPixel(_aaec, _bgbc); _aaec++ {
				if _fcacf = _dcca.SetPixel(_aaec, _bgbc, 0); _fcacf != nil {
					return nil, _a.Wrap(_fcacf, _bagg, "\u0032n\u0064\u0020\u0073\u0065\u0074")
				}
			}
			if _fcacf = _gbdf(_babda, _fdf, _aaec-1, _bddf._dabg, _bddf._bege, _ddfad, _agfe); _fcacf != nil {
				return nil, _a.Wrap(_fcacf, _bagg, "n\u006f\u0072\u006d\u0061\u006c\u0020\u0070\u0075\u0073\u0068")
			}
			if _aaec > _bddf._fdddg+1 {
				if _fcacf = _gbdf(_babda, _bddf._fdddg+1, _aaec-1, _bddf._dabg, -_bddf._bege, _ddfad, _agfe); _fcacf != nil {
					return nil, _a.Wrap(_fcacf, _bagg, "\u006ce\u0061k\u0020\u006f\u006e\u0020\u0072i\u0067\u0068t\u0020\u0073\u0069\u0064\u0065")
				}
			}
			for _aaec++; _aaec <= _bddf._fdddg && _aaec <= _bdcf && !_dcca.GetPixel(_aaec, _bgbc); _aaec++ {
			}
			_fdf = _aaec
			if !(_aaec <= _bddf._fdddg && _aaec <= _bdcf) {
				break
			}
		}
	}
	_agfe.Max.X++
	_agfe.Max.Y++
	return _agfe, nil
}
func (_gdbg *Bitmap) ConnComponents(bms *Bitmaps, connectivity int) (_agbc *Boxes, _eded error) {
	const _eeeg = "B\u0069\u0074\u006d\u0061p.\u0043o\u006e\u006e\u0043\u006f\u006dp\u006f\u006e\u0065\u006e\u0074\u0073"
	if _gdbg == nil {
		return nil, _a.Error(_eeeg, "\u0070r\u006f\u0076\u0069\u0064e\u0064\u0020\u0065\u006d\u0070t\u0079 \u0027b\u0027\u0020\u0062\u0069\u0074\u006d\u0061p")
	}
	if connectivity != 4 && connectivity != 8 {
		return nil, _a.Error(_eeeg, "\u0063\u006f\u006ene\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u006e\u006f\u0074\u0020\u0034\u0020\u006f\u0072\u0020\u0038")
	}
	if bms == nil {
		if _agbc, _eded = _gdbg.connComponentsBB(connectivity); _eded != nil {
			return nil, _a.Wrap(_eded, _eeeg, "")
		}
	} else {
		if _agbc, _eded = _gdbg.connComponentsBitmapsBB(bms, connectivity); _eded != nil {
			return nil, _a.Wrap(_eded, _eeeg, "")
		}
	}
	return _agbc, nil
}
func _fda(_fgcc, _agbb int) int {
	if _fgcc < _agbb {
		return _fgcc
	}
	return _agbb
}
func (_abfa *Bitmaps) SortByWidth() { _gbda := (*byWidth)(_abfa); _e.Sort(_gbda) }
func TstVSymbol(t *_f.T, scale ...int) *Bitmap {
	_ccge, _dgcg := NewWithData(5, 5, []byte{0x88, 0x88, 0x88, 0x50, 0x20})
	_g.NoError(t, _dgcg)
	return TstGetScaledSymbol(t, _ccge, scale...)
}
func _fggb(_eddd, _fccf *Bitmap, _bcae, _aafb, _cba uint, _bcfb, _fcfa int, _dede bool, _afgg, _daab int) error {
	for _dccf := _bcfb; _dccf < _fcfa; _dccf++ {
		if _afgg+1 < len(_eddd.Data) {
			_gbfb := _dccf+1 == _fcfa
			_afgd, _agdc := _eddd.GetByte(_afgg)
			if _agdc != nil {
				return _agdc
			}
			_afgg++
			_afgd <<= _bcae
			_gcfcb, _agdc := _eddd.GetByte(_afgg)
			if _agdc != nil {
				return _agdc
			}
			_gcfcb >>= _aafb
			_ggfd := _afgd | _gcfcb
			if _gbfb && !_dede {
				_ggfd = _fge(_cba, _ggfd)
			}
			_agdc = _fccf.SetByte(_daab, _ggfd)
			if _agdc != nil {
				return _agdc
			}
			_daab++
			if _gbfb && _dede {
				_ecdd, _dgff := _eddd.GetByte(_afgg)
				if _dgff != nil {
					return _dgff
				}
				_ecdd <<= _bcae
				_ggfd = _fge(_cba, _ecdd)
				if _dgff = _fccf.SetByte(_daab, _ggfd); _dgff != nil {
					return _dgff
				}
			}
			continue
		}
		_bdae, _dgfa := _eddd.GetByte(_afgg)
		if _dgfa != nil {
			_db.Log.Debug("G\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0068\u0065\u0020\u0076\u0061l\u0075\u0065\u0020\u0061\u0074\u003a\u0020%\u0064\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020%\u0073", _afgg, _dgfa)
			return _dgfa
		}
		_bdae <<= _bcae
		_afgg++
		_dgfa = _fccf.SetByte(_daab, _bdae)
		if _dgfa != nil {
			return _dgfa
		}
		_daab++
	}
	return nil
}
func (_dcdae *ClassedPoints) xSortFunction() func(_fgaac int, _eddb int) bool {
	return func(_cgff, _caec int) bool { return _dcdae.XAtIndex(_cgff) < _dcdae.XAtIndex(_caec) }
}
func (_cdagf Points) XSorter() func(_eeec, _acf int) bool {
	return func(_adeb, _fcfbf int) bool { return _cdagf[_adeb].X < _cdagf[_fcfbf].X }
}
func (_ecd *Bitmap) SetPixel(x, y int, pixel byte) error {
	_dbfd := _ecd.GetByteIndex(x, y)
	if _dbfd > len(_ecd.Data)-1 {
		return _a.Errorf("\u0053\u0065\u0074\u0050\u0069\u0078\u0065\u006c", "\u0069\u006e\u0064\u0065x \u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065\u003a\u0020%\u0064", _dbfd)
	}
	_eabe := _ecd.GetBitOffset(x)
	_cfe := uint(7 - _eabe)
	_ggf := _ecd.Data[_dbfd]
	var _adfc byte
	if pixel == 1 {
		_adfc = _ggf | (pixel & 0x01 << _cfe)
	} else {
		_adfc = _ggf &^ (1 << _cfe)
	}
	_ecd.Data[_dbfd] = _adfc
	return nil
}

var (
	_fafe = _cfcb()
	_ceab = _dfb()
	_cefd = _afb()
)

func (_ffg *Bitmap) GetByteIndex(x, y int) int { return y*_ffg.RowStride + (x >> 3) }
func _cec(_ab *Bitmap, _bgb ...int) (_eed *Bitmap, _dgb error) {
	const _fg = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0043\u0061\u0073\u0063\u0061\u0064\u0065"
	if _ab == nil {
		return nil, _a.Error(_fg, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	if len(_bgb) == 0 || len(_bgb) > 4 {
		return nil, _a.Error(_fg, "t\u0068\u0065\u0072\u0065\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0061\u0074\u0020\u006cea\u0073\u0074\u0020\u006fn\u0065\u0020\u0061\u006e\u0064\u0020\u0061\u0074\u0020mo\u0073\u0074 \u0034\u0020\u006c\u0065\u0076\u0065\u006c\u0073")
	}
	if _bgb[0] <= 0 {
		_db.Log.Debug("\u006c\u0065\u0076\u0065\u006c\u0031\u0020\u003c\u003d\u0020\u0030 \u002d\u0020\u006e\u006f\u0020\u0072\u0065\u0064\u0075\u0063t\u0069\u006f\u006e")
		_eed, _dgb = _bggbg(nil, _ab)
		if _dgb != nil {
			return nil, _a.Wrap(_dgb, _fg, "l\u0065\u0076\u0065\u006c\u0031\u0020\u003c\u003d\u0020\u0030")
		}
		return _eed, nil
	}
	_ec := _faf()
	_eed = _ab
	for _fad, _fgb := range _bgb {
		if _fgb <= 0 {
			break
		}
		_eed, _dgb = _ecc(_eed, _fgb, _ec)
		if _dgb != nil {
			return nil, _a.Wrapf(_dgb, _fg, "\u006c\u0065\u0076\u0065\u006c\u0025\u0064\u0020\u0072\u0065\u0064\u0075c\u0074\u0069\u006f\u006e", _fad)
		}
	}
	return _eed, nil
}
func (_aced *Bitmap) setEightBytes(_cccc int, _bfe uint64) error {
	_ede := _aced.RowStride - (_cccc % _aced.RowStride)
	if _aced.RowStride != _aced.Width>>3 {
		_ede--
	}
	if _ede >= 8 {
		return _aced.setEightFullBytes(_cccc, _bfe)
	}
	return _aced.setEightPartlyBytes(_cccc, _ede, _bfe)
}
func _aefcd(_cee, _bdfc *Bitmap, _defe *Selection) (*Bitmap, error) {
	const _ecb = "\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u004d\u006f\u0072\u0070\u0068A\u0072\u0067\u0073\u0032"
	var _cdcg, _gffa int
	if _bdfc == nil {
		return nil, _a.Error(_ecb, "s\u006fu\u0072\u0063\u0065\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _defe == nil {
		return nil, _a.Error(_ecb, "\u0073e\u006c \u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	_cdcg = _defe.Width
	_gffa = _defe.Height
	if _cdcg == 0 || _gffa == 0 {
		return nil, _a.Error(_ecb, "\u0073\u0065\u006c\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u0020\u0030")
	}
	if _cee == nil {
		return _bdfc.createTemplate(), nil
	}
	if _gecd := _cee.resizeImageData(_bdfc); _gecd != nil {
		return nil, _gecd
	}
	return _cee, nil
}
func (_daacf *Bitmaps) AddBox(box *_ac.Rectangle) { _daacf.Boxes = append(_daacf.Boxes, box) }

type SizeComparison int

func _bebaf(_agge *Bitmap, _agga *_fc.Stack, _abde, _defc, _aadff int) (_facad *_ac.Rectangle, _dege error) {
	const _dcg = "\u0073e\u0065d\u0046\u0069\u006c\u006c\u0053\u0074\u0061\u0063\u006b\u0042\u0042"
	if _agge == nil {
		return nil, _a.Error(_dcg, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0073\u0027\u0020\u0042\u0069\u0074\u006d\u0061\u0070")
	}
	if _agga == nil {
		return nil, _a.Error(_dcg, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0027\u0073\u0074ac\u006b\u0027")
	}
	switch _aadff {
	case 4:
		if _facad, _dege = _ffbc(_agge, _agga, _abde, _defc); _dege != nil {
			return nil, _a.Wrap(_dege, _dcg, "")
		}
		return _facad, nil
	case 8:
		if _facad, _dege = _bcaa(_agge, _agga, _abde, _defc); _dege != nil {
			return nil, _a.Wrap(_dege, _dcg, "")
		}
		return _facad, nil
	default:
		return nil, _a.Errorf(_dcg, "\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076\u0069\u0074\u0079\u0020\u0069\u0073 \u006eo\u0074\u0020\u0034\u0020\u006f\u0072\u0020\u0038\u003a\u0020\u0027\u0025\u0064\u0027", _aadff)
	}
}
func _be(_ge, _gd *Bitmap) (_aa error) {
	const _dbb = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0034"
	_dcc := _gd.RowStride
	_ef := _ge.RowStride
	_cg := _gd.RowStride*4 - _ge.RowStride
	var (
		_df, _fba                            byte
		_eg                                  uint32
		_dd, _egd, _deg, _cb, _ddf, _bg, _eb int
	)
	for _deg = 0; _deg < _gd.Height; _deg++ {
		_dd = _deg * _dcc
		_egd = 4 * _deg * _ef
		for _cb = 0; _cb < _dcc; _cb++ {
			_df = _gd.Data[_dd+_cb]
			_eg = _ceab[_df]
			_bg = _egd + _cb*4
			if _cg != 0 && (_cb+1)*4 > _ge.RowStride {
				for _ddf = _cg; _ddf > 0; _ddf-- {
					_fba = byte((_eg >> uint(_ddf*8)) & 0xff)
					_eb = _bg + (_cg - _ddf)
					if _aa = _ge.SetByte(_eb, _fba); _aa != nil {
						return _a.Wrapf(_aa, _dbb, "D\u0069\u0066\u0066\u0065\u0072\u0065n\u0074\u0020\u0072\u006f\u0077\u0073\u0074\u0072\u0069d\u0065\u0073\u002e \u004b:\u0020\u0025\u0064", _ddf)
					}
				}
			} else if _aa = _ge.setFourBytes(_bg, _eg); _aa != nil {
				return _a.Wrap(_aa, _dbb, "")
			}
			if _aa = _ge.setFourBytes(_egd+_cb*4, _ceab[_gd.Data[_dd+_cb]]); _aa != nil {
				return _a.Wrap(_aa, _dbb, "")
			}
		}
		for _ddf = 1; _ddf < 4; _ddf++ {
			for _cb = 0; _cb < _ef; _cb++ {
				if _aa = _ge.SetByte(_egd+_ddf*_ef+_cb, _ge.Data[_egd+_cb]); _aa != nil {
					return _a.Wrapf(_aa, _dbb, "\u0063\u006f\u0070\u0079\u0020\u0027\u0071\u0075\u0061\u0064\u0072\u0061\u0062l\u0065\u0027\u0020\u006c\u0069\u006ee\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0079\u0074\u0065\u003a \u0027\u0025\u0064\u0027", _ddf, _cb)
				}
			}
		}
	}
	return nil
}

type Selection struct {
	Height, Width int
	Cx, Cy        int
	Name          string
	Data          [][]SelectionValue
}

func _faf() (_fbge []byte) {
	_fbge = make([]byte, 256)
	for _dda := 0; _dda < 256; _dda++ {
		_agg := byte(_dda)
		_fbge[_agg] = (_agg & 0x01) | ((_agg & 0x04) >> 1) | ((_agg & 0x10) >> 2) | ((_agg & 0x40) >> 3) | ((_agg & 0x02) << 3) | ((_agg & 0x08) << 2) | ((_agg & 0x20) << 1) | (_agg & 0x80)
	}
	return _fbge
}
func (_fdbfd *ClassedPoints) XAtIndex(i int) float32 { return (*_fdbfd.Points)[_fdbfd.IntSlice[i]].X }
func _gba(_dgf, _acb *Bitmap, _fbab int, _bc []byte, _dec int) (_add error) {
	const _dca = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0031"
	var (
		_eec, _bdc, _fab, _ebea, _cdf, _fgbf, _dga, _cdfd int
		_aadg, _cc                                        uint32
		_eff, _adf                                        byte
		_dbc                                              uint16
	)
	_gaef := make([]byte, 4)
	_aae := make([]byte, 4)
	for _fab = 0; _fab < _dgf.Height-1; _fab, _ebea = _fab+2, _ebea+1 {
		_eec = _fab * _dgf.RowStride
		_bdc = _ebea * _acb.RowStride
		for _cdf, _fgbf = 0, 0; _cdf < _dec; _cdf, _fgbf = _cdf+4, _fgbf+1 {
			for _dga = 0; _dga < 4; _dga++ {
				_cdfd = _eec + _cdf + _dga
				if _cdfd <= len(_dgf.Data)-1 && _cdfd < _eec+_dgf.RowStride {
					_gaef[_dga] = _dgf.Data[_cdfd]
				} else {
					_gaef[_dga] = 0x00
				}
				_cdfd = _eec + _dgf.RowStride + _cdf + _dga
				if _cdfd <= len(_dgf.Data)-1 && _cdfd < _eec+(2*_dgf.RowStride) {
					_aae[_dga] = _dgf.Data[_cdfd]
				} else {
					_aae[_dga] = 0x00
				}
			}
			_aadg = _gc.BigEndian.Uint32(_gaef)
			_cc = _gc.BigEndian.Uint32(_aae)
			_cc |= _aadg
			_cc |= _cc << 1
			_cc &= 0xaaaaaaaa
			_aadg = _cc | (_cc << 7)
			_eff = byte(_aadg >> 24)
			_adf = byte((_aadg >> 8) & 0xff)
			_cdfd = _bdc + _fgbf
			if _cdfd+1 == len(_acb.Data)-1 || _cdfd+1 >= _bdc+_acb.RowStride {
				_acb.Data[_cdfd] = _bc[_eff]
			} else {
				_dbc = (uint16(_bc[_eff]) << 8) | uint16(_bc[_adf])
				if _add = _acb.setTwoBytes(_cdfd, _dbc); _add != nil {
					return _a.Wrapf(_add, _dca, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _cdfd)
				}
				_fgbf++
			}
		}
	}
	return nil
}
func (_dge *Bitmap) String() string {
	var _dbgd = "\u000a"
	for _efbd := 0; _efbd < _dge.Height; _efbd++ {
		var _degd string
		for _dgac := 0; _dgac < _dge.Width; _dgac++ {
			_adcf := _dge.GetPixel(_dgac, _efbd)
			if _adcf {
				_degd += "\u0031"
			} else {
				_degd += "\u0030"
			}
		}
		_dbgd += _degd + "\u000a"
	}
	return _dbgd
}
func MorphSequence(src *Bitmap, sequence ...MorphProcess) (*Bitmap, error) {
	return _cbed(src, sequence...)
}
func _dfb() (_gag [256]uint32) {
	for _ba := 0; _ba < 256; _ba++ {
		if _ba&0x01 != 0 {
			_gag[_ba] |= 0xf
		}
		if _ba&0x02 != 0 {
			_gag[_ba] |= 0xf0
		}
		if _ba&0x04 != 0 {
			_gag[_ba] |= 0xf00
		}
		if _ba&0x08 != 0 {
			_gag[_ba] |= 0xf000
		}
		if _ba&0x10 != 0 {
			_gag[_ba] |= 0xf0000
		}
		if _ba&0x20 != 0 {
			_gag[_ba] |= 0xf00000
		}
		if _ba&0x40 != 0 {
			_gag[_ba] |= 0xf000000
		}
		if _ba&0x80 != 0 {
			_gag[_ba] |= 0xf0000000
		}
	}
	return _gag
}
func TstImageBitmapInverseData() []byte {
	_cgdfd := _faed.Copy()
	_cgdfd.InverseData()
	return _cgdfd.Data
}
func (_cgg *Bitmap) addPadBits() (_dece error) {
	const _dcf = "\u0062\u0069\u0074\u006d\u0061\u0070\u002e\u0061\u0064\u0064\u0050\u0061d\u0042\u0069\u0074\u0073"
	_bfc := _cgg.Width % 8
	if _bfc == 0 {
		return nil
	}
	_eee := _cgg.Width / 8
	_cca := _c.NewReader(_cgg.Data)
	_dac := make([]byte, _cgg.Height*_cgg.RowStride)
	_ccab := _c.NewWriterMSB(_dac)
	_bdff := make([]byte, _eee)
	var (
		_eeg  int
		_egba uint64
	)
	for _eeg = 0; _eeg < _cgg.Height; _eeg++ {
		if _, _dece = _cca.Read(_bdff); _dece != nil {
			return _a.Wrap(_dece, _dcf, "\u0066u\u006c\u006c\u0020\u0062\u0079\u0074e")
		}
		if _, _dece = _ccab.Write(_bdff); _dece != nil {
			return _a.Wrap(_dece, _dcf, "\u0066\u0075\u006c\u006c\u0020\u0062\u0079\u0074\u0065\u0073")
		}
		if _egba, _dece = _cca.ReadBits(byte(_bfc)); _dece != nil {
			return _a.Wrap(_dece, _dcf, "\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0062\u0069\u0074\u0073")
		}
		if _dece = _ccab.WriteByte(byte(_egba) << uint(8-_bfc)); _dece != nil {
			return _a.Wrap(_dece, _dcf, "\u006ca\u0073\u0074\u0020\u0062\u0079\u0074e")
		}
	}
	_cgg.Data = _ccab.Data()
	return nil
}

const (
	SelDontCare SelectionValue = iota
	SelHit
	SelMiss
)

type shift int

func NewClassedPoints(points *Points, classes _fc.IntSlice) (*ClassedPoints, error) {
	const _cabf = "\u004e\u0065w\u0043\u006c\u0061s\u0073\u0065\u0064\u0050\u006f\u0069\u006e\u0074\u0073"
	if points == nil {
		return nil, _a.Error(_cabf, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0070\u006f\u0069\u006e\u0074\u0073")
	}
	if classes == nil {
		return nil, _a.Error(_cabf, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0063\u006c\u0061ss\u0065\u0073")
	}
	_fceee := &ClassedPoints{Points: points, IntSlice: classes}
	if _afdb := _fceee.validateIntSlice(); _afdb != nil {
		return nil, _a.Wrap(_afdb, _cabf, "")
	}
	return _fceee, nil
}
func (_fdbf Points) GetIntX(i int) (int, error) {
	if i >= len(_fdbf) {
		return 0, _a.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065t\u0049\u006e\u0074\u0058", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return int(_fdbf[i].X), nil
}
func (_fegb *Boxes) SelectBySize(width, height int, tp LocationFilter, relation SizeComparison) (_edcc *Boxes, _ecea error) {
	const _aca = "\u0042o\u0078e\u0073\u002e\u0053\u0065\u006ce\u0063\u0074B\u0079\u0053\u0069\u007a\u0065"
	if _fegb == nil {
		return nil, _a.Error(_aca, "b\u006f\u0078\u0065\u0073 '\u0062'\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(*_fegb) == 0 {
		return _fegb, nil
	}
	switch tp {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _a.Errorf(_aca, "\u0069\u006e\u0076al\u0069\u0064\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", tp)
	}
	switch relation {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE:
	default:
		return nil, _a.Errorf(_aca, "i\u006e\u0076\u0061\u006c\u0069\u0064 \u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020t\u0079\u0070\u0065:\u0020'\u0025\u0064\u0027", tp)
	}
	_dagb := _fegb.makeSizeIndicator(width, height, tp, relation)
	_ggdf, _ecea := _fegb.selectWithIndicator(_dagb)
	if _ecea != nil {
		return nil, _a.Wrap(_ecea, _aca, "")
	}
	return _ggdf, nil
}
func _gefd(_gfgcb *Bitmap) (_acca *Bitmap, _dfdd int, _bead error) {
	const _bfde = "\u0042i\u0074\u006d\u0061\u0070.\u0077\u006f\u0072\u0064\u004da\u0073k\u0042y\u0044\u0069\u006c\u0061\u0074\u0069\u006fn"
	if _gfgcb == nil {
		return nil, 0, _a.Errorf(_bfde, "\u0027\u0073\u0027\u0020bi\u0074\u006d\u0061\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064")
	}
	var _bgcbb, _gfef *Bitmap
	if _bgcbb, _bead = _bggbg(nil, _gfgcb); _bead != nil {
		return nil, 0, _a.Wrap(_bead, _bfde, "\u0063\u006f\u0070\u0079\u0020\u0027\u0073\u0027")
	}
	var (
		_bcgfb       [13]int
		_fcg, _dfaba int
	)
	_eabcg := 12
	_aeda := _fc.NewNumSlice(_eabcg + 1)
	_gebe := _fc.NewNumSlice(_eabcg + 1)
	var _accb *Boxes
	for _fea := 0; _fea <= _eabcg; _fea++ {
		if _fea == 0 {
			if _gfef, _bead = _bggbg(nil, _bgcbb); _bead != nil {
				return nil, 0, _a.Wrap(_bead, _bfde, "\u0066i\u0072\u0073\u0074\u0020\u0062\u006d2")
			}
		} else {
			if _gfef, _bead = _cbed(_bgcbb, MorphProcess{Operation: MopDilation, Arguments: []int{2, 1}}); _bead != nil {
				return nil, 0, _a.Wrap(_bead, _bfde, "\u0064\u0069\u006ca\u0074\u0069\u006f\u006e\u0020\u0062\u006d\u0032")
			}
		}
		if _accb, _bead = _gfef.connComponentsBB(4); _bead != nil {
			return nil, 0, _a.Wrap(_bead, _bfde, "")
		}
		_bcgfb[_fea] = len(*_accb)
		_aeda.AddInt(_bcgfb[_fea])
		switch _fea {
		case 0:
			_fcg = _bcgfb[0]
		default:
			_dfaba = _bcgfb[_fea-1] - _bcgfb[_fea]
			_gebe.AddInt(_dfaba)
		}
		_bgcbb = _gfef
	}
	_dfcf := true
	_ebegb := 2
	var _ccg, _aegc int
	for _fcea := 1; _fcea < len(*_gebe); _fcea++ {
		if _ccg, _bead = _aeda.GetInt(_fcea); _bead != nil {
			return nil, 0, _a.Wrap(_bead, _bfde, "\u0043\u0068\u0065\u0063ki\u006e\u0067\u0020\u0062\u0065\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0069o\u006e")
		}
		if _dfcf && _ccg < int(0.3*float32(_fcg)) {
			_ebegb = _fcea + 1
			_dfcf = false
		}
		if _dfaba, _bead = _gebe.GetInt(_fcea); _bead != nil {
			return nil, 0, _a.Wrap(_bead, _bfde, "\u0067\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u006ea\u0044\u0069\u0066\u0066")
		}
		if _dfaba > _aegc {
			_aegc = _dfaba
		}
	}
	_cggb := _gfgcb.XResolution
	if _cggb == 0 {
		_cggb = 150
	}
	if _cggb > 110 {
		_ebegb++
	}
	if _ebegb < 2 {
		_db.Log.Trace("J\u0042\u0049\u0047\u0032\u0020\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u0042\u0065\u0073\u0074 \u0074\u006f\u0020\u006d\u0069\u006e\u0069\u006d\u0075\u006d a\u006c\u006c\u006fw\u0061b\u006c\u0065")
		_ebegb = 2
	}
	_dfdd = _ebegb + 1
	if _acca, _bead = _aegaf(nil, _gfgcb, _ebegb+1, 1); _bead != nil {
		return nil, 0, _a.Wrap(_bead, _bfde, "\u0067\u0065\u0074\u0074in\u0067\u0020\u006d\u0061\u0073\u006b\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
	}
	return _acca, _dfdd, nil
}
func (_afdd *Bitmaps) makeSizeIndicator(_bcad, _fdacg int, _dfdfgg LocationFilter, _edddg SizeComparison) (_faddc *_fc.NumSlice, _gabfg error) {
	const _bebb = "\u0042i\u0074\u006d\u0061\u0070s\u002e\u006d\u0061\u006b\u0065S\u0069z\u0065I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _afdd == nil {
		return nil, _a.Error(_bebb, "\u0062\u0069\u0074ma\u0070\u0073\u0020\u0027\u0062\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	switch _dfdfgg {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _a.Errorf(_bebb, "\u0070\u0072\u006f\u0076\u0069d\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u006fc\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", _dfdfgg)
	}
	switch _edddg {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE, SizeSelectIfEQ:
	default:
		return nil, _a.Errorf(_bebb, "\u0069\u006e\u0076\u0061li\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025d\u0027", _edddg)
	}
	_faddc = &_fc.NumSlice{}
	var (
		_dgffa, _gggg, _aedg int
		_ecef                *Bitmap
	)
	for _, _ecef = range _afdd.Values {
		_dgffa = 0
		_gggg, _aedg = _ecef.Width, _ecef.Height
		switch _dfdfgg {
		case LocSelectWidth:
			if (_edddg == SizeSelectIfLT && _gggg < _bcad) || (_edddg == SizeSelectIfGT && _gggg > _bcad) || (_edddg == SizeSelectIfLTE && _gggg <= _bcad) || (_edddg == SizeSelectIfGTE && _gggg >= _bcad) || (_edddg == SizeSelectIfEQ && _gggg == _bcad) {
				_dgffa = 1
			}
		case LocSelectHeight:
			if (_edddg == SizeSelectIfLT && _aedg < _fdacg) || (_edddg == SizeSelectIfGT && _aedg > _fdacg) || (_edddg == SizeSelectIfLTE && _aedg <= _fdacg) || (_edddg == SizeSelectIfGTE && _aedg >= _fdacg) || (_edddg == SizeSelectIfEQ && _aedg == _fdacg) {
				_dgffa = 1
			}
		case LocSelectIfEither:
			if (_edddg == SizeSelectIfLT && (_gggg < _bcad || _aedg < _fdacg)) || (_edddg == SizeSelectIfGT && (_gggg > _bcad || _aedg > _fdacg)) || (_edddg == SizeSelectIfLTE && (_gggg <= _bcad || _aedg <= _fdacg)) || (_edddg == SizeSelectIfGTE && (_gggg >= _bcad || _aedg >= _fdacg)) || (_edddg == SizeSelectIfEQ && (_gggg == _bcad || _aedg == _fdacg)) {
				_dgffa = 1
			}
		case LocSelectIfBoth:
			if (_edddg == SizeSelectIfLT && (_gggg < _bcad && _aedg < _fdacg)) || (_edddg == SizeSelectIfGT && (_gggg > _bcad && _aedg > _fdacg)) || (_edddg == SizeSelectIfLTE && (_gggg <= _bcad && _aedg <= _fdacg)) || (_edddg == SizeSelectIfGTE && (_gggg >= _bcad && _aedg >= _fdacg)) || (_edddg == SizeSelectIfEQ && (_gggg == _bcad && _aedg == _fdacg)) {
				_dgffa = 1
			}
		}
		_faddc.AddInt(_dgffa)
	}
	return _faddc, nil
}
func _bdde(_adeg, _bgbe *Bitmap, _bea, _addc, _dbfdc, _ffgb, _cgge int, _cfbb CombinationOperator) error {
	var _abfb int
	_bgce := func() {
		_abfb++
		_dbfdc += _bgbe.RowStride
		_ffgb += _adeg.RowStride
		_cgge += _adeg.RowStride
	}
	for _abfb = _bea; _abfb < _addc; _bgce() {
		_bdbf := _dbfdc
		for _feggd := _ffgb; _feggd <= _cgge; _feggd++ {
			_acba, _bebf := _bgbe.GetByte(_bdbf)
			if _bebf != nil {
				return _bebf
			}
			_bde, _bebf := _adeg.GetByte(_feggd)
			if _bebf != nil {
				return _bebf
			}
			if _bebf = _bgbe.SetByte(_bdbf, _bgbg(_acba, _bde, _cfbb)); _bebf != nil {
				return _bebf
			}
			_bdbf++
		}
	}
	return nil
}
func _ecaga(_eeccc, _ggac *Bitmap, _agda *Selection) (*Bitmap, error) {
	const _caffc = "c\u006c\u006f\u0073\u0065\u0042\u0069\u0074\u006d\u0061\u0070"
	var _bebfa error
	if _eeccc, _bebfa = _aefcd(_eeccc, _ggac, _agda); _bebfa != nil {
		return nil, _bebfa
	}
	_acd, _bebfa := _egcg(nil, _ggac, _agda)
	if _bebfa != nil {
		return nil, _a.Wrap(_bebfa, _caffc, "")
	}
	if _, _bebfa = _abbc(_eeccc, _acd, _agda); _bebfa != nil {
		return nil, _a.Wrap(_bebfa, _caffc, "")
	}
	return _eeccc, nil
}
func (_ffad *Bitmap) setBit(_cbgcb int) { _ffad.Data[(_cbgcb >> 3)] |= 0x80 >> uint(_cbgcb&7) }
func _cdgge() []int {
	_efaa := make([]int, 256)
	for _dcad := 0; _dcad <= 0xff; _dcad++ {
		_gbcd := byte(_dcad)
		_efaa[_gbcd] = int(_gbcd&0x1) + (int(_gbcd>>1) & 0x1) + (int(_gbcd>>2) & 0x1) + (int(_gbcd>>3) & 0x1) + (int(_gbcd>>4) & 0x1) + (int(_gbcd>>5) & 0x1) + (int(_gbcd>>6) & 0x1) + (int(_gbcd>>7) & 0x1)
	}
	return _efaa
}
func (_bdgag *ClassedPoints) Swap(i, j int) {
	_bdgag.IntSlice[i], _bdgag.IntSlice[j] = _bdgag.IntSlice[j], _bdgag.IntSlice[i]
}
func (_bdcg *byHeight) Len() int { return len(_bdcg.Values) }

const (
	PixSrc             RasterOperator = 0xc
	PixDst             RasterOperator = 0xa
	PixNotSrc          RasterOperator = 0x3
	PixNotDst          RasterOperator = 0x5
	PixClr             RasterOperator = 0x0
	PixSet             RasterOperator = 0xf
	PixSrcOrDst        RasterOperator = 0xe
	PixSrcAndDst       RasterOperator = 0x8
	PixSrcXorDst       RasterOperator = 0x6
	PixNotSrcOrDst     RasterOperator = 0xb
	PixNotSrcAndDst    RasterOperator = 0x2
	PixSrcOrNotDst     RasterOperator = 0xd
	PixSrcAndNotDst    RasterOperator = 0x4
	PixNotPixSrcOrDst  RasterOperator = 0x1
	PixNotPixSrcAndDst RasterOperator = 0x7
	PixNotPixSrcXorDst RasterOperator = 0x9
	PixPaint                          = PixSrcOrDst
	PixSubtract                       = PixNotSrcAndDst
	PixMask                           = PixSrcAndDst
)

func (_edac *Bitmaps) GroupByHeight() (*BitmapsArray, error) {
	const _eaab = "\u0047\u0072\u006f\u0075\u0070\u0042\u0079\u0048\u0065\u0069\u0067\u0068\u0074"
	if len(_edac.Values) == 0 {
		return nil, _a.Error(_eaab, "\u006eo\u0020v\u0061\u006c\u0075\u0065\u0073 \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	_fdeb := &BitmapsArray{}
	_edac.SortByHeight()
	_faab := -1
	_feefc := -1
	for _dfdgf := 0; _dfdgf < len(_edac.Values); _dfdgf++ {
		_bfdea := _edac.Values[_dfdgf].Height
		if _bfdea > _faab {
			_faab = _bfdea
			_feefc++
			_fdeb.Values = append(_fdeb.Values, &Bitmaps{})
		}
		_fdeb.Values[_feefc].AddBitmap(_edac.Values[_dfdgf])
	}
	return _fdeb, nil
}
func (_caa *Points) Add(pt *Points) error {
	const _fggd = "\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0041\u0064\u0064"
	if _caa == nil {
		return _a.Error(_fggd, "\u0070o\u0069n\u0074\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if pt == nil {
		return _a.Error(_fggd, "a\u0072\u0067\u0075\u006d\u0065\u006et\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u006eo\u0074\u0020\u0064e\u0066i\u006e\u0065\u0064")
	}
	*_caa = append(*_caa, *pt...)
	return nil
}
func (_cga *Bitmap) AddBorderGeneral(left, right, top, bot int, val int) (*Bitmap, error) {
	return _cga.addBorderGeneral(left, right, top, bot, val)
}
func MakePixelSumTab8() []int { return _cdgge() }
func _cfg(_aef, _fac *Bitmap, _bfdd int, _degc []byte, _cda int) (_gaeb error) {
	const _cgc = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0032"
	var (
		_aaf, _fae, _faea, _eca, _gec, _dbba, _fbgd, _efg int
		_bfg, _ecag, _efb, _cbe                           uint32
		_egf, _cff                                        byte
		_faa                                              uint16
	)
	_bdf := make([]byte, 4)
	_acbe := make([]byte, 4)
	for _faea = 0; _faea < _aef.Height-1; _faea, _eca = _faea+2, _eca+1 {
		_aaf = _faea * _aef.RowStride
		_fae = _eca * _fac.RowStride
		for _gec, _dbba = 0, 0; _gec < _cda; _gec, _dbba = _gec+4, _dbba+1 {
			for _fbgd = 0; _fbgd < 4; _fbgd++ {
				_efg = _aaf + _gec + _fbgd
				if _efg <= len(_aef.Data)-1 && _efg < _aaf+_aef.RowStride {
					_bdf[_fbgd] = _aef.Data[_efg]
				} else {
					_bdf[_fbgd] = 0x00
				}
				_efg = _aaf + _aef.RowStride + _gec + _fbgd
				if _efg <= len(_aef.Data)-1 && _efg < _aaf+(2*_aef.RowStride) {
					_acbe[_fbgd] = _aef.Data[_efg]
				} else {
					_acbe[_fbgd] = 0x00
				}
			}
			_bfg = _gc.BigEndian.Uint32(_bdf)
			_ecag = _gc.BigEndian.Uint32(_acbe)
			_efb = _bfg & _ecag
			_efb |= _efb << 1
			_cbe = _bfg | _ecag
			_cbe &= _cbe << 1
			_ecag = _efb | _cbe
			_ecag &= 0xaaaaaaaa
			_bfg = _ecag | (_ecag << 7)
			_egf = byte(_bfg >> 24)
			_cff = byte((_bfg >> 8) & 0xff)
			_efg = _fae + _dbba
			if _efg+1 == len(_fac.Data)-1 || _efg+1 >= _fae+_fac.RowStride {
				if _gaeb = _fac.SetByte(_efg, _degc[_egf]); _gaeb != nil {
					return _a.Wrapf(_gaeb, _cgc, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _efg)
				}
			} else {
				_faa = (uint16(_degc[_egf]) << 8) | uint16(_degc[_cff])
				if _gaeb = _fac.setTwoBytes(_efg, _faa); _gaeb != nil {
					return _a.Wrapf(_gaeb, _cgc, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _efg)
				}
				_dbba++
			}
		}
	}
	return nil
}
func (_dcadg *Bitmaps) SelectBySize(width, height int, tp LocationFilter, relation SizeComparison) (_dddcb *Bitmaps, _dgfgc error) {
	const _gcbe = "B\u0069t\u006d\u0061\u0070\u0073\u002e\u0053\u0065\u006ce\u0063\u0074\u0042\u0079Si\u007a\u0065"
	if _dcadg == nil {
		return nil, _a.Error(_gcbe, "\u0027\u0062\u0027 B\u0069\u0074\u006d\u0061\u0070\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	switch tp {
	case LocSelectWidth, LocSelectHeight, LocSelectIfEither, LocSelectIfBoth:
	default:
		return nil, _a.Errorf(_gcbe, "\u0070\u0072\u006f\u0076\u0069d\u0065\u0064\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u006fc\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0064", tp)
	}
	switch relation {
	case SizeSelectIfLT, SizeSelectIfGT, SizeSelectIfLTE, SizeSelectIfGTE, SizeSelectIfEQ:
	default:
		return nil, _a.Errorf(_gcbe, "\u0069\u006e\u0076\u0061li\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0027\u0025d\u0027", relation)
	}
	_dbgc, _dgfgc := _dcadg.makeSizeIndicator(width, height, tp, relation)
	if _dgfgc != nil {
		return nil, _a.Wrap(_dgfgc, _gcbe, "")
	}
	_dddcb, _dgfgc = _dcadg.selectByIndicator(_dbgc)
	if _dgfgc != nil {
		return nil, _a.Wrap(_dgfgc, _gcbe, "")
	}
	return _dddcb, nil
}
func (_ebac *Bitmap) centroid(_eadg, _bebc []int) (Point, error) {
	_efeg := Point{}
	_ebac.setPadBits(0)
	if len(_eadg) == 0 {
		_eadg = _caed()
	}
	if len(_bebc) == 0 {
		_bebc = _cdgge()
	}
	var _dfee, _daac, _fcfc, _gcce, _fcgb, _gaac int
	var _bbd byte
	for _fcgb = 0; _fcgb < _ebac.Height; _fcgb++ {
		_afeb := _ebac.RowStride * _fcgb
		_gcce = 0
		for _gaac = 0; _gaac < _ebac.RowStride; _gaac++ {
			_bbd = _ebac.Data[_afeb+_gaac]
			if _bbd != 0 {
				_gcce += _bebc[_bbd]
				_dfee += _eadg[_bbd] + _gaac*8*_bebc[_bbd]
			}
		}
		_fcfc += _gcce
		_daac += _gcce * _fcgb
	}
	if _fcfc != 0 {
		_efeg.X = float32(_dfee) / float32(_fcfc)
		_efeg.Y = float32(_daac) / float32(_fcfc)
	}
	return _efeg, nil
}
func RasterOperation(dest *Bitmap, dx, dy, dw, dh int, op RasterOperator, src *Bitmap, sx, sy int) error {
	return _gcaa(dest, dx, dy, dw, dh, op, src, sx, sy)
}
func (_adag Points) Get(i int) (Point, error) {
	if i > len(_adag)-1 {
		return Point{}, _a.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _adag[i], nil
}
func (_eggdf *Bitmap) RasterOperation(dx, dy, dw, dh int, op RasterOperator, src *Bitmap, sx, sy int) error {
	return _gcaa(_eggdf, dx, dy, dw, dh, op, src, sx, sy)
}
func (_cbee *ClassedPoints) validateIntSlice() error {
	const _baf = "\u0076\u0061l\u0069\u0064\u0061t\u0065\u0049\u006e\u0074\u0053\u006c\u0069\u0063\u0065"
	for _, _acec := range _cbee.IntSlice {
		if _acec >= (_cbee.Points.Size()) {
			return _a.Errorf(_baf, "c\u006c\u0061\u0073\u0073\u0020\u0069\u0064\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0061\u0020\u0076\u0061\u006ci\u0064 \u0069\u006e\u0064\u0065x\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u003a\u0020\u0025\u0064", _acec, _cbee.Points.Size())
		}
	}
	return nil
}

type Bitmaps struct {
	Values []*Bitmap
	Boxes  []*_ac.Rectangle
}

func _bccf(_dbdb, _dbga *Bitmap, _bcbd *Selection) (*Bitmap, error) {
	const _ecgac = "\u006f\u0070\u0065\u006e"
	var _bbgd error
	_dbdb, _bbgd = _aefcd(_dbdb, _dbga, _bcbd)
	if _bbgd != nil {
		return nil, _a.Wrap(_bbgd, _ecgac, "")
	}
	_ebee, _bbgd := _abbc(nil, _dbga, _bcbd)
	if _bbgd != nil {
		return nil, _a.Wrap(_bbgd, _ecgac, "")
	}
	_, _bbgd = _egcg(_dbdb, _ebee, _bcbd)
	if _bbgd != nil {
		return nil, _a.Wrap(_bbgd, _ecgac, "")
	}
	return _dbdb, nil
}
func _cad(_eecd, _cadf *Bitmap, _egc int, _eaa []byte, _cdfg int) (_dgd error) {
	const _bgea = "\u0072\u0065\u0064uc\u0065\u0052\u0061\u006e\u006b\u0042\u0069\u006e\u0061\u0072\u0079\u0032\u004c\u0065\u0076\u0065\u006c\u0034"
	var (
		_edc, _abf, _gbg, _gdg, _fada, _cef, _facc, _fee int
		_ddc, _ggg                                       uint32
		_gfed, _cgb                                      byte
		_ebc                                             uint16
	)
	_agd := make([]byte, 4)
	_dgc := make([]byte, 4)
	for _gbg = 0; _gbg < _eecd.Height-1; _gbg, _gdg = _gbg+2, _gdg+1 {
		_edc = _gbg * _eecd.RowStride
		_abf = _gdg * _cadf.RowStride
		for _fada, _cef = 0, 0; _fada < _cdfg; _fada, _cef = _fada+4, _cef+1 {
			for _facc = 0; _facc < 4; _facc++ {
				_fee = _edc + _fada + _facc
				if _fee <= len(_eecd.Data)-1 && _fee < _edc+_eecd.RowStride {
					_agd[_facc] = _eecd.Data[_fee]
				} else {
					_agd[_facc] = 0x00
				}
				_fee = _edc + _eecd.RowStride + _fada + _facc
				if _fee <= len(_eecd.Data)-1 && _fee < _edc+(2*_eecd.RowStride) {
					_dgc[_facc] = _eecd.Data[_fee]
				} else {
					_dgc[_facc] = 0x00
				}
			}
			_ddc = _gc.BigEndian.Uint32(_agd)
			_ggg = _gc.BigEndian.Uint32(_dgc)
			_ggg &= _ddc
			_ggg &= _ggg << 1
			_ggg &= 0xaaaaaaaa
			_ddc = _ggg | (_ggg << 7)
			_gfed = byte(_ddc >> 24)
			_cgb = byte((_ddc >> 8) & 0xff)
			_fee = _abf + _cef
			if _fee+1 == len(_cadf.Data)-1 || _fee+1 >= _abf+_cadf.RowStride {
				_cadf.Data[_fee] = _eaa[_gfed]
				if _dgd = _cadf.SetByte(_fee, _eaa[_gfed]); _dgd != nil {
					return _a.Wrapf(_dgd, _bgea, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0025d", _fee)
				}
			} else {
				_ebc = (uint16(_eaa[_gfed]) << 8) | uint16(_eaa[_cgb])
				if _dgd = _cadf.setTwoBytes(_fee, _ebc); _dgd != nil {
					return _a.Wrapf(_dgd, _bgea, "s\u0065\u0074\u0074\u0069\u006e\u0067 \u0074\u0077\u006f\u0020\u0062\u0079t\u0065\u0073\u0020\u0066\u0061\u0069\u006ce\u0064\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a\u0020%\u0064", _fee)
				}
				_cef++
			}
		}
	}
	return nil
}
func (_dade CombinationOperator) String() string {
	var _dgee string
	switch _dade {
	case CmbOpOr:
		_dgee = "\u004f\u0052"
	case CmbOpAnd:
		_dgee = "\u0041\u004e\u0044"
	case CmbOpXor:
		_dgee = "\u0058\u004f\u0052"
	case CmbOpXNor:
		_dgee = "\u0058\u004e\u004f\u0052"
	case CmbOpReplace:
		_dgee = "\u0052E\u0050\u004c\u0041\u0043\u0045"
	case CmbOpNot:
		_dgee = "\u004e\u004f\u0054"
	}
	return _dgee
}
func (_abafc *Bitmaps) SelectByIndexes(idx []int) (*Bitmaps, error) {
	const _caeb = "B\u0069\u0074\u006d\u0061\u0070\u0073.\u0053\u006f\u0072\u0074\u0049\u006e\u0064\u0065\u0078e\u0073\u0042\u0079H\u0065i\u0067\u0068\u0074"
	_ebdaf, _eage := _abafc.selectByIndexes(idx)
	if _eage != nil {
		return nil, _a.Wrap(_eage, _caeb, "")
	}
	return _ebdaf, nil
}
func (_dedg Points) GetGeometry(i int) (_fafd, _afea float32, _egce error) {
	if i > len(_dedg)-1 {
		return 0, 0, _a.Errorf("\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0065\u0074", "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	_cggf := _dedg[i]
	return _cggf.X, _cggf.Y, nil
}
func (_gggd *Bitmaps) String() string {
	_cbeef := _fb.Builder{}
	for _, _ffff := range _gggd.Values {
		_cbeef.WriteString(_ffff.String())
		_cbeef.WriteRune('\n')
	}
	return _cbeef.String()
}
func (_aefa *Bitmap) AddBorder(borderSize, val int) (*Bitmap, error) {
	if borderSize == 0 {
		return _aefa.Copy(), nil
	}
	_ebf, _cade := _aefa.addBorderGeneral(borderSize, borderSize, borderSize, borderSize, val)
	if _cade != nil {
		return nil, _a.Wrap(_cade, "\u0041d\u0064\u0042\u006f\u0072\u0064\u0065r", "")
	}
	return _ebf, nil
}
func _cgae(_ggbg *Bitmap, _bgbbd int) (*Bitmap, error) {
	const _aeee = "\u0065x\u0070a\u006e\u0064\u0052\u0065\u0070\u006c\u0069\u0063\u0061\u0074\u0065"
	if _ggbg == nil {
		return nil, _a.Error(_aeee, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _bgbbd <= 0 {
		return nil, _a.Error(_aeee, "i\u006e\u0076\u0061\u006cid\u0020f\u0061\u0063\u0074\u006f\u0072 \u002d\u0020\u003c\u003d\u0020\u0030")
	}
	if _bgbbd == 1 {
		_bgcef, _cagf := _bggbg(nil, _ggbg)
		if _cagf != nil {
			return nil, _a.Wrap(_cagf, _aeee, "\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u003d\u0020\u0031")
		}
		return _bgcef, nil
	}
	_defa, _aafg := _fce(_ggbg, _bgbbd, _bgbbd)
	if _aafg != nil {
		return nil, _a.Wrap(_aafg, _aeee, "")
	}
	return _defa, nil
}
func TstPSymbol(t *_f.T) *Bitmap {
	t.Helper()
	_bebbg := New(5, 8)
	_g.NoError(t, _bebbg.SetPixel(0, 0, 1))
	_g.NoError(t, _bebbg.SetPixel(1, 0, 1))
	_g.NoError(t, _bebbg.SetPixel(2, 0, 1))
	_g.NoError(t, _bebbg.SetPixel(3, 0, 1))
	_g.NoError(t, _bebbg.SetPixel(4, 1, 1))
	_g.NoError(t, _bebbg.SetPixel(0, 1, 1))
	_g.NoError(t, _bebbg.SetPixel(4, 2, 1))
	_g.NoError(t, _bebbg.SetPixel(0, 2, 1))
	_g.NoError(t, _bebbg.SetPixel(4, 3, 1))
	_g.NoError(t, _bebbg.SetPixel(0, 3, 1))
	_g.NoError(t, _bebbg.SetPixel(0, 4, 1))
	_g.NoError(t, _bebbg.SetPixel(1, 4, 1))
	_g.NoError(t, _bebbg.SetPixel(2, 4, 1))
	_g.NoError(t, _bebbg.SetPixel(3, 4, 1))
	_g.NoError(t, _bebbg.SetPixel(0, 5, 1))
	_g.NoError(t, _bebbg.SetPixel(0, 6, 1))
	_g.NoError(t, _bebbg.SetPixel(0, 7, 1))
	return _bebbg
}
func _gcaa(_gcec *Bitmap, _cfae, _gaff, _accc, _aeff int, _adba RasterOperator, _ggc *Bitmap, _eeegg, _cdbg int) error {
	const _aaff = "\u0072a\u0073t\u0065\u0072\u004f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e"
	if _gcec == nil {
		return _a.Error(_aaff, "\u006e\u0069\u006c\u0020\u0027\u0064\u0065\u0073\u0074\u0027\u0020\u0042i\u0074\u006d\u0061\u0070")
	}
	if _adba == PixDst {
		return nil
	}
	switch _adba {
	case PixClr, PixSet, PixNotDst:
		_cfdc(_gcec, _cfae, _gaff, _accc, _aeff, _adba)
		return nil
	}
	if _ggc == nil {
		_db.Log.Debug("\u0052a\u0073\u0074e\u0072\u004f\u0070\u0065r\u0061\u0074\u0069o\u006e\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020bi\u0074\u006d\u0061p\u0020\u0069s\u0020\u006e\u006f\u0074\u0020\u0064e\u0066\u0069n\u0065\u0064")
		return _a.Error(_aaff, "\u006e\u0069l\u0020\u0027\u0073r\u0063\u0027\u0020\u0062\u0069\u0074\u006d\u0061\u0070")
	}
	if _feec := _adea(_gcec, _cfae, _gaff, _accc, _aeff, _adba, _ggc, _eeegg, _cdbg); _feec != nil {
		return _a.Wrap(_feec, _aaff, "")
	}
	return nil
}
func (_eccb *Bitmaps) selectByIndexes(_baab []int) (*Bitmaps, error) {
	_facff := &Bitmaps{}
	for _, _dcdg := range _baab {
		_fbdg, _dfcbg := _eccb.GetBitmap(_dcdg)
		if _dfcbg != nil {
			return nil, _a.Wrap(_dfcbg, "\u0073e\u006ce\u0063\u0074\u0042\u0079\u0049\u006e\u0064\u0065\u0078\u0065\u0073", "")
		}
		_facff.AddBitmap(_fbdg)
	}
	return _facff, nil
}
func (_gcfcce *BitmapsArray) AddBox(box *_ac.Rectangle) { _gcfcce.Boxes = append(_gcfcce.Boxes, box) }
func (_bbcb *Selection) findMaxTranslations() (_bcbac, _cagfb, _fdad, _aage int) {
	for _cbfb := 0; _cbfb < _bbcb.Height; _cbfb++ {
		for _cbac := 0; _cbac < _bbcb.Width; _cbac++ {
			if _bbcb.Data[_cbfb][_cbac] == SelHit {
				_bcbac = _dba(_bcbac, _bbcb.Cx-_cbac)
				_cagfb = _dba(_cagfb, _bbcb.Cy-_cbfb)
				_fdad = _dba(_fdad, _cbac-_bbcb.Cx)
				_aage = _dba(_aage, _cbfb-_bbcb.Cy)
			}
		}
	}
	return _bcbac, _cagfb, _fdad, _aage
}
func _fce(_dffd *Bitmap, _fe, _bgg int) (*Bitmap, error) {
	const _gde = "e\u0078\u0070\u0061\u006edB\u0069n\u0061\u0072\u0079\u0052\u0065p\u006c\u0069\u0063\u0061\u0074\u0065"
	if _dffd == nil {
		return nil, _a.Error(_gde, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _fe <= 0 || _bgg <= 0 {
		return nil, _a.Error(_gde, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0073\u0063\u0061l\u0065\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u003a\u0020<\u003d\u0020\u0030")
	}
	if _fe == _bgg {
		if _fe == 1 {
			_dfc, _ad := _bggbg(nil, _dffd)
			if _ad != nil {
				return nil, _a.Wrap(_ad, _gde, "\u0078\u0046\u0061\u0063\u0074\u0020\u003d\u003d\u0020y\u0046\u0061\u0063\u0074")
			}
			return _dfc, nil
		}
		if _fe == 2 || _fe == 4 || _fe == 8 {
			_bda, _cfa := _bbg(_dffd, _fe)
			if _cfa != nil {
				return nil, _a.Wrap(_cfa, _gde, "\u0078\u0046a\u0063\u0074\u0020i\u006e\u0020\u007b\u0032\u002c\u0034\u002c\u0038\u007d")
			}
			return _bda, nil
		}
	}
	_ffd := _fe * _dffd.Width
	_afd := _bgg * _dffd.Height
	_bef := New(_ffd, _afd)
	_gfa := _bef.RowStride
	var (
		_aad, _dffde, _fa, _bge, _ca int
		_da                          byte
		_gga                         error
	)
	for _dffde = 0; _dffde < _dffd.Height; _dffde++ {
		_aad = _bgg * _dffde * _gfa
		for _fa = 0; _fa < _dffd.Width; _fa++ {
			if _fdd := _dffd.GetPixel(_fa, _dffde); _fdd {
				_ca = _fe * _fa
				for _bge = 0; _bge < _fe; _bge++ {
					_bef.setBit(_aad*8 + _ca + _bge)
				}
			}
		}
		for _bge = 1; _bge < _bgg; _bge++ {
			_bf := _aad + _bge*_gfa
			for _gda := 0; _gda < _gfa; _gda++ {
				if _da, _gga = _bef.GetByte(_aad + _gda); _gga != nil {
					return nil, _a.Wrapf(_gga, _gde, "\u0072\u0065\u0070\u006cic\u0061\u0074\u0069\u006e\u0067\u0020\u006c\u0069\u006e\u0065\u003a\u0020\u0027\u0025d\u0027", _bge)
				}
				if _gga = _bef.SetByte(_bf+_gda, _da); _gga != nil {
					return nil, _a.Wrap(_gga, _gde, "\u0053\u0065\u0074\u0074in\u0067\u0020\u0062\u0079\u0074\u0065\u0020\u0066\u0061\u0069\u006c\u0065\u0064")
				}
			}
		}
	}
	return _bef, nil
}
func (_gaga *Bitmaps) GetBox(i int) (*_ac.Rectangle, error) {
	const _adgga = "\u0047\u0065\u0074\u0042\u006f\u0078"
	if _gaga == nil {
		return nil, _a.Error(_adgga, "\u0070\u0072\u006f\u0076id\u0065\u0064\u0020\u006e\u0069\u006c\u0020\u0027\u0042\u0069\u0074\u006d\u0061\u0070s\u0027")
	}
	if i > len(_gaga.Boxes)-1 {
		return nil, _a.Errorf(_adgga, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _gaga.Boxes[i], nil
}
func NewWithUnpaddedData(width, height int, data []byte) (*Bitmap, error) {
	const _feg = "\u004e\u0065\u0077\u0057it\u0068\u0055\u006e\u0070\u0061\u0064\u0064\u0065\u0064\u0044\u0061\u0074\u0061"
	_fddc := _egda(width, height)
	_fddc.Data = data
	if _dcee := ((width * height) + 7) >> 3; len(data) < _dcee {
		return nil, _a.Errorf(_feg, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064a\u0074\u0061\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002e\u0020\u0054\u0068\u0065\u0020\u0064\u0061t\u0061\u0020s\u0068\u006fu\u006c\u0064\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074 l\u0065\u0061\u0073\u0074\u003a\u0020\u0027\u0025\u0064'\u0020\u0062\u0079\u0074\u0065\u0073", len(data), _dcee)
	}
	if _bcb := _fddc.addPadBits(); _bcb != nil {
		return nil, _a.Wrap(_bcb, _feg, "")
	}
	return _fddc, nil
}
func TstGetScaledSymbol(t *_f.T, sm *Bitmap, scale ...int) *Bitmap {
	if len(scale) == 0 {
		return sm
	}
	if scale[0] == 1 {
		return sm
	}
	_abbe, _bcce := MorphSequence(sm, MorphProcess{Operation: MopReplicativeBinaryExpansion, Arguments: scale})
	_g.NoError(t, _bcce)
	return _abbe
}
func _fd(_aceg, _ea *Bitmap) (_gf error) {
	const _afa = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0038"
	_agb := _ea.RowStride
	_fcf := _aceg.RowStride
	var _ddb, _ae, _ebb, _bd, _egde int
	for _ebb = 0; _ebb < _ea.Height; _ebb++ {
		_ddb = _ebb * _agb
		_ae = 8 * _ebb * _fcf
		for _bd = 0; _bd < _agb; _bd++ {
			if _gf = _aceg.setEightBytes(_ae+_bd*8, _cefd[_ea.Data[_ddb+_bd]]); _gf != nil {
				return _a.Wrap(_gf, _afa, "")
			}
		}
		for _egde = 1; _egde < 8; _egde++ {
			for _bd = 0; _bd < _fcf; _bd++ {
				if _gf = _aceg.SetByte(_ae+_egde*_fcf+_bd, _aceg.Data[_ae+_bd]); _gf != nil {
					return _a.Wrap(_gf, _afa, "")
				}
			}
		}
	}
	return nil
}
func (_dee *Bitmap) ToImage() _ac.Image {
	_ada, _cgac := _d.NewImage(_dee.Width, _dee.Height, 1, 1, _dee.Data, nil, nil)
	if _cgac != nil {
		_db.Log.Error("\u0043\u006f\u006e\u0076\u0065\u0072\u0074\u0069\u006e\u0067\u0020j\u0062\u0069\u0067\u0032\u002e\u0042\u0069\u0074m\u0061p\u0020\u0074\u006f\u0020\u0069\u006d\u0061\u0067\u0065\u0075\u0074\u0069\u006c\u002e\u0049\u006d\u0061\u0067e\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _cgac)
	}
	return _ada
}

const _bbgab = 5000

func (_dab *Bitmap) InverseData() { _dab.inverseData() }
func (_fddd Points) YSorter() func(_cfag, _ebfbb int) bool {
	return func(_debb, _efffa int) bool { return _fddd[_debb].Y < _fddd[_efffa].Y }
}
func (_bafb *ClassedPoints) ySortFunction() func(_abged int, _dbbe int) bool {
	return func(_gbgce, _eabd int) bool { return _bafb.YAtIndex(_gbgce) < _bafb.YAtIndex(_eabd) }
}
func (_fcbgb *Bitmap) setAll() error {
	_decb := _gcaa(_fcbgb, 0, 0, _fcbgb.Width, _fcbgb.Height, PixSet, nil, 0, 0)
	if _decb != nil {
		return _a.Wrap(_decb, "\u0073\u0065\u0074\u0041\u006c\u006c", "")
	}
	return nil
}
func (_ddd *Bitmap) countPixels() int {
	var (
		_ccba int
		_beb  uint8
		_fcfb byte
		_gfg  int
	)
	_bggb := _ddd.RowStride
	_gaf := uint(_ddd.Width & 0x07)
	if _gaf != 0 {
		_beb = uint8((0xff << (8 - _gaf)) & 0xff)
		_bggb--
	}
	for _bfb := 0; _bfb < _ddd.Height; _bfb++ {
		for _gfg = 0; _gfg < _bggb; _gfg++ {
			_fcfb = _ddd.Data[_bfb*_ddd.RowStride+_gfg]
			_ccba += int(_afdc[_fcfb])
		}
		if _gaf != 0 {
			_ccba += int(_afdc[_ddd.Data[_bfb*_ddd.RowStride+_gfg]&_beb])
		}
	}
	return _ccba
}

var MorphBC BoundaryCondition

func Copy(d, s *Bitmap) (*Bitmap, error) { return _bggbg(d, s) }

var _gcda = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3E, 0x78, 0x27, 0xC2, 0x27, 0x91, 0x00, 0x22, 0x48, 0x21, 0x03, 0x24, 0x91, 0x00, 0x22, 0x48, 0x21, 0x02, 0xA4, 0x95, 0x00, 0x22, 0x48, 0x21, 0x02, 0x64, 0x9B, 0x00, 0x3C, 0x78, 0x21, 0x02, 0x27, 0x91, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7F, 0xF8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func (_ffec *ClassedPoints) SortByX() { _ffec._fbdba = _ffec.xSortFunction(); _e.Sort(_ffec) }
func _dbdcge(_cbea, _adbg *Bitmap, _fcfcg, _bab int) (*Bitmap, error) {
	const _cfac = "\u0063\u006c\u006f\u0073\u0065\u0053\u0061\u0066\u0065B\u0072\u0069\u0063\u006b"
	if _adbg == nil {
		return nil, _a.Error(_cfac, "\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _fcfcg < 1 || _bab < 1 {
		return nil, _a.Error(_cfac, "\u0068s\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0073\u0069z\u0065\u0020\u006e\u006f\u0074\u0020\u003e\u003d\u0020\u0031")
	}
	if _fcfcg == 1 && _bab == 1 {
		return _bggbg(_cbea, _adbg)
	}
	if MorphBC == SymmetricMorphBC {
		_egdee, _bbc := _aegaf(_cbea, _adbg, _fcfcg, _bab)
		if _bbc != nil {
			return nil, _a.Wrap(_bbc, _cfac, "\u0053\u0079m\u006d\u0065\u0074r\u0069\u0063\u004d\u006f\u0072\u0070\u0068\u0042\u0043")
		}
		return _egdee, nil
	}
	_cabc := _dba(_fcfcg/2, _bab/2)
	_gcdb := 8 * ((_cabc + 7) / 8)
	_ffdc, _adde := _adbg.AddBorder(_gcdb, 0)
	if _adde != nil {
		return nil, _a.Wrapf(_adde, _cfac, "\u0042\u006f\u0072\u0064\u0065\u0072\u0053\u0069\u007ae\u003a\u0020\u0025\u0064", _gcdb)
	}
	var _ecgd, _gdcc *Bitmap
	if _fcfcg == 1 || _bab == 1 {
		_gfb := SelCreateBrick(_bab, _fcfcg, _bab/2, _fcfcg/2, SelHit)
		_ecgd, _adde = _ecaga(nil, _ffdc, _gfb)
		if _adde != nil {
			return nil, _a.Wrap(_adde, _cfac, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
	} else {
		_aegd := SelCreateBrick(1, _fcfcg, 0, _fcfcg/2, SelHit)
		_gggf, _ecdeb := _egcg(nil, _ffdc, _aegd)
		if _ecdeb != nil {
			return nil, _a.Wrap(_ecdeb, _cfac, "\u0072\u0065\u0067\u0075la\u0072\u0020\u002d\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0064\u0069\u006c\u0061t\u0065")
		}
		_abgg := SelCreateBrick(_bab, 1, _bab/2, 0, SelHit)
		_ecgd, _ecdeb = _egcg(nil, _gggf, _abgg)
		if _ecdeb != nil {
			return nil, _a.Wrap(_ecdeb, _cfac, "\u0072\u0065\u0067ul\u0061\u0072\u0020\u002d\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
		}
		if _, _ecdeb = _abbc(_gggf, _ecgd, _aegd); _ecdeb != nil {
			return nil, _a.Wrap(_ecdeb, _cfac, "r\u0065\u0067\u0075\u006car\u0020-\u0020\u0066\u0069\u0072\u0073t\u0020\u0065\u0072\u006f\u0064\u0065")
		}
		if _, _ecdeb = _abbc(_ecgd, _gggf, _abgg); _ecdeb != nil {
			return nil, _a.Wrap(_ecdeb, _cfac, "\u0072\u0065\u0067\u0075la\u0072\u0020\u002d\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0065\u0072\u006fd\u0065")
		}
	}
	if _gdcc, _adde = _ecgd.RemoveBorder(_gcdb); _adde != nil {
		return nil, _a.Wrap(_adde, _cfac, "\u0072e\u0067\u0075\u006c\u0061\u0072")
	}
	if _cbea == nil {
		return _gdcc, nil
	}
	if _, _adde = _bggbg(_cbea, _gdcc); _adde != nil {
		return nil, _adde
	}
	return _cbea, nil
}

type ClassedPoints struct {
	*Points
	_fc.IntSlice
	_fbdba func(_eabb, _eeccg int) bool
}

func TstOSymbol(t *_f.T, scale ...int) *Bitmap {
	_gaaca, _aefce := NewWithData(4, 5, []byte{0xF0, 0x90, 0x90, 0x90, 0xF0})
	_g.NoError(t, _aefce)
	return TstGetScaledSymbol(t, _gaaca, scale...)
}
func (_cacbf *ClassedPoints) Len() int { return _cacbf.IntSlice.Size() }
func (_cffga *Bitmaps) GroupByWidth() (*BitmapsArray, error) {
	const _aace = "\u0047\u0072\u006fu\u0070\u0042\u0079\u0057\u0069\u0064\u0074\u0068"
	if len(_cffga.Values) == 0 {
		return nil, _a.Error(_aace, "\u006eo\u0020v\u0061\u006c\u0075\u0065\u0073 \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	_efea := &BitmapsArray{}
	_cffga.SortByWidth()
	_fbfd := -1
	_ggcf := -1
	for _gcfea := 0; _gcfea < len(_cffga.Values); _gcfea++ {
		_agged := _cffga.Values[_gcfea].Width
		if _agged > _fbfd {
			_fbfd = _agged
			_ggcf++
			_efea.Values = append(_efea.Values, &Bitmaps{})
		}
		_efea.Values[_ggcf].AddBitmap(_cffga.Values[_gcfea])
	}
	return _efea, nil
}

const (
	_ SizeSelection = iota
	SizeSelectByWidth
	SizeSelectByHeight
	SizeSelectByMaxDimension
	SizeSelectByArea
	SizeSelectByPerimeter
)

func TstISymbol(t *_f.T, scale ...int) *Bitmap {
	_dfg, _bebd := NewWithData(1, 5, []byte{0x80, 0x80, 0x80, 0x80, 0x80})
	_g.NoError(t, _bebd)
	return TstGetScaledSymbol(t, _dfg, scale...)
}
func (_fedbd *byHeight) Less(i, j int) bool { return _fedbd.Values[i].Height < _fedbd.Values[j].Height }
func (_bcg *Bitmap) GetPixel(x, y int) bool {
	_fbd := _bcg.GetByteIndex(x, y)
	_ega := _bcg.GetBitOffset(x)
	_ccc := uint(7 - _ega)
	if _fbd > len(_bcg.Data)-1 {
		_db.Log.Debug("\u0054\u0072\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0067\u0065\u0074\u0020\u0070\u0069\u0078\u0065\u006c\u0020o\u0075\u0074\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0064\u0061\u0074\u0061\u0020\u0072\u0061\u006e\u0067\u0065\u002e \u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0027\u0025\u0064'\u002c\u0020\u0062m\u003a\u0020\u0027\u0025\u0073\u0027", x, y, _bcg)
		return false
	}
	if (_bcg.Data[_fbd]>>_ccc)&0x01 >= 1 {
		return true
	}
	return false
}

type SelectionValue int

func TstWSymbol(t *_f.T, scale ...int) *Bitmap {
	_aedc, _beafg := NewWithData(5, 5, []byte{0x88, 0x88, 0xA8, 0xD8, 0x88})
	_g.NoError(t, _beafg)
	return TstGetScaledSymbol(t, _aedc, scale...)
}
func SelCreateBrick(h, w int, cy, cx int, tp SelectionValue) *Selection {
	_gdacg := _ebca(h, w, "")
	_gdacg.setOrigin(cy, cx)
	var _ffbb, _fgbda int
	for _ffbb = 0; _ffbb < h; _ffbb++ {
		for _fgbda = 0; _fgbda < w; _fgbda++ {
			_gdacg.Data[_ffbb][_fgbda] = tp
		}
	}
	return _gdacg
}
func (_ecgf *Boxes) makeSizeIndicator(_bdbff, _cag int, _edeg LocationFilter, _agde SizeComparison) *_fc.NumSlice {
	_ccbc := &_fc.NumSlice{}
	var _eebdd, _acc, _ggee int
	for _, _bacd := range *_ecgf {
		_eebdd = 0
		_acc, _ggee = _bacd.Dx(), _bacd.Dy()
		switch _edeg {
		case LocSelectWidth:
			if (_agde == SizeSelectIfLT && _acc < _bdbff) || (_agde == SizeSelectIfGT && _acc > _bdbff) || (_agde == SizeSelectIfLTE && _acc <= _bdbff) || (_agde == SizeSelectIfGTE && _acc >= _bdbff) {
				_eebdd = 1
			}
		case LocSelectHeight:
			if (_agde == SizeSelectIfLT && _ggee < _cag) || (_agde == SizeSelectIfGT && _ggee > _cag) || (_agde == SizeSelectIfLTE && _ggee <= _cag) || (_agde == SizeSelectIfGTE && _ggee >= _cag) {
				_eebdd = 1
			}
		case LocSelectIfEither:
			if (_agde == SizeSelectIfLT && (_ggee < _cag || _acc < _bdbff)) || (_agde == SizeSelectIfGT && (_ggee > _cag || _acc > _bdbff)) || (_agde == SizeSelectIfLTE && (_ggee <= _cag || _acc <= _bdbff)) || (_agde == SizeSelectIfGTE && (_ggee >= _cag || _acc >= _bdbff)) {
				_eebdd = 1
			}
		case LocSelectIfBoth:
			if (_agde == SizeSelectIfLT && (_ggee < _cag && _acc < _bdbff)) || (_agde == SizeSelectIfGT && (_ggee > _cag && _acc > _bdbff)) || (_agde == SizeSelectIfLTE && (_ggee <= _cag && _acc <= _bdbff)) || (_agde == SizeSelectIfGTE && (_ggee >= _cag && _acc >= _bdbff)) {
				_eebdd = 1
			}
		}
		_ccbc.AddInt(_eebdd)
	}
	return _ccbc
}
func (_fgaa *Bitmap) setEightPartlyBytes(_dgbd, _bfae int, _dead uint64) (_edg error) {
	var (
		_bebe byte
		_cbga int
	)
	const _egfb = "\u0073\u0065\u0074\u0045ig\u0068\u0074\u0050\u0061\u0072\u0074\u006c\u0079\u0042\u0079\u0074\u0065\u0073"
	for _bgd := 1; _bgd <= _bfae; _bgd++ {
		_cbga = 64 - _bgd*8
		_bebe = byte(_dead >> uint(_cbga) & 0xff)
		_db.Log.Trace("\u0074\u0065\u006d\u0070\u003a\u0020\u0025\u0030\u0038\u0062\u002c\u0020\u0069\u006e\u0064\u0065\u0078\u003a %\u0064,\u0020\u0069\u0064\u0078\u003a\u0020\u0025\u0064\u002c\u0020\u0066\u0075l\u006c\u0042\u0079\u0074\u0065\u0073\u004e\u0075\u006d\u0062\u0065\u0072\u003a\u0020\u0025\u0064\u002c \u0073\u0068\u0069\u0066\u0074\u003a\u0020\u0025\u0064", _bebe, _dgbd, _dgbd+_bgd-1, _bfae, _cbga)
		if _edg = _fgaa.SetByte(_dgbd+_bgd-1, _bebe); _edg != nil {
			return _a.Wrap(_edg, _egfb, "\u0066\u0075\u006c\u006c\u0042\u0079\u0074\u0065")
		}
	}
	_ggfg := _fgaa.RowStride*8 - _fgaa.Width
	if _ggfg == 0 {
		return nil
	}
	_cbga -= 8
	_bebe = byte(_dead>>uint(_cbga)&0xff) << uint(_ggfg)
	if _edg = _fgaa.SetByte(_dgbd+_bfae, _bebe); _edg != nil {
		return _a.Wrap(_edg, _egfb, "\u0070\u0061\u0064\u0064\u0065\u0064")
	}
	return nil
}
func (_ffce *Bitmap) Equals(s *Bitmap) bool {
	if len(_ffce.Data) != len(s.Data) || _ffce.Width != s.Width || _ffce.Height != s.Height {
		return false
	}
	for _bbga := 0; _bbga < _ffce.Height; _bbga++ {
		_efgd := _bbga * _ffce.RowStride
		for _ead := 0; _ead < _ffce.RowStride; _ead++ {
			if _ffce.Data[_efgd+_ead] != s.Data[_efgd+_ead] {
				return false
			}
		}
	}
	return true
}
func NewWithData(width, height int, data []byte) (*Bitmap, error) {
	const _gaaf = "N\u0065\u0077\u0057\u0069\u0074\u0068\u0044\u0061\u0074\u0061"
	_fade := _egda(width, height)
	_fade.Data = data
	if len(data) < height*_fade.RowStride {
		return nil, _a.Errorf(_gaaf, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061\u0020l\u0065\u006e\u0067\u0074\u0068\u003a \u0025\u0064\u0020\u002d\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u003a\u0020\u0025\u0064", len(data), height*_fade.RowStride)
	}
	return _fade, nil
}
func TstCSymbol(t *_f.T) *Bitmap {
	t.Helper()
	_fdcg := New(6, 6)
	_g.NoError(t, _fdcg.SetPixel(1, 0, 1))
	_g.NoError(t, _fdcg.SetPixel(2, 0, 1))
	_g.NoError(t, _fdcg.SetPixel(3, 0, 1))
	_g.NoError(t, _fdcg.SetPixel(4, 0, 1))
	_g.NoError(t, _fdcg.SetPixel(0, 1, 1))
	_g.NoError(t, _fdcg.SetPixel(5, 1, 1))
	_g.NoError(t, _fdcg.SetPixel(0, 2, 1))
	_g.NoError(t, _fdcg.SetPixel(0, 3, 1))
	_g.NoError(t, _fdcg.SetPixel(0, 4, 1))
	_g.NoError(t, _fdcg.SetPixel(5, 4, 1))
	_g.NoError(t, _fdcg.SetPixel(1, 5, 1))
	_g.NoError(t, _fdcg.SetPixel(2, 5, 1))
	_g.NoError(t, _fdcg.SetPixel(3, 5, 1))
	_g.NoError(t, _fdcg.SetPixel(4, 5, 1))
	return _fdcg
}
func _bgbg(_adgf, _bcgd byte, _aagf CombinationOperator) byte {
	switch _aagf {
	case CmbOpOr:
		return _bcgd | _adgf
	case CmbOpAnd:
		return _bcgd & _adgf
	case CmbOpXor:
		return _bcgd ^ _adgf
	case CmbOpXNor:
		return ^(_bcgd ^ _adgf)
	case CmbOpNot:
		return ^(_bcgd)
	default:
		return _bcgd
	}
}
func (_fbbg *BitmapsArray) AddBitmaps(bm *Bitmaps) { _fbbg.Values = append(_fbbg.Values, bm) }
func _dba(_befd, _eagf int) int {
	if _befd > _eagf {
		return _befd
	}
	return _eagf
}
func _abbc(_bfeb, _ddac *Bitmap, _bdea *Selection) (*Bitmap, error) {
	const _aaeg = "\u0065\u0072\u006fd\u0065"
	var (
		_afgc error
		_dcde *Bitmap
	)
	_bfeb, _afgc = _gddb(_bfeb, _ddac, _bdea, &_dcde)
	if _afgc != nil {
		return nil, _a.Wrap(_afgc, _aaeg, "")
	}
	if _afgc = _bfeb.setAll(); _afgc != nil {
		return nil, _a.Wrap(_afgc, _aaeg, "")
	}
	var _fbda SelectionValue
	for _dgbf := 0; _dgbf < _bdea.Height; _dgbf++ {
		for _cacb := 0; _cacb < _bdea.Width; _cacb++ {
			_fbda = _bdea.Data[_dgbf][_cacb]
			if _fbda == SelHit {
				_afgc = _gcaa(_bfeb, _bdea.Cx-_cacb, _bdea.Cy-_dgbf, _ddac.Width, _ddac.Height, PixSrcAndDst, _dcde, 0, 0)
				if _afgc != nil {
					return nil, _a.Wrap(_afgc, _aaeg, "")
				}
			}
		}
	}
	if MorphBC == SymmetricMorphBC {
		return _bfeb, nil
	}
	_dbag, _gacd, _cfbe, _aeec := _bdea.findMaxTranslations()
	if _dbag > 0 {
		if _afgc = _bfeb.RasterOperation(0, 0, _dbag, _ddac.Height, PixClr, nil, 0, 0); _afgc != nil {
			return nil, _a.Wrap(_afgc, _aaeg, "\u0078\u0070\u0020\u003e\u0020\u0030")
		}
	}
	if _cfbe > 0 {
		if _afgc = _bfeb.RasterOperation(_ddac.Width-_cfbe, 0, _cfbe, _ddac.Height, PixClr, nil, 0, 0); _afgc != nil {
			return nil, _a.Wrap(_afgc, _aaeg, "\u0078\u006e\u0020\u003e\u0020\u0030")
		}
	}
	if _gacd > 0 {
		if _afgc = _bfeb.RasterOperation(0, 0, _ddac.Width, _gacd, PixClr, nil, 0, 0); _afgc != nil {
			return nil, _a.Wrap(_afgc, _aaeg, "\u0079\u0070\u0020\u003e\u0020\u0030")
		}
	}
	if _aeec > 0 {
		if _afgc = _bfeb.RasterOperation(0, _ddac.Height-_aeec, _ddac.Width, _aeec, PixClr, nil, 0, 0); _afgc != nil {
			return nil, _a.Wrap(_afgc, _aaeg, "\u0079\u006e\u0020\u003e\u0020\u0030")
		}
	}
	return _bfeb, nil
}
func TstImageBitmap() *Bitmap { return _faed.Copy() }
func TstWordBitmapWithSpaces(t *_f.T, scale ...int) *Bitmap {
	_bbbf := 1
	if len(scale) > 0 {
		_bbbf = scale[0]
	}
	_caddf := 3
	_efba := 9 + 7 + 15 + 2*_caddf + 2*_caddf
	_bbab := 5 + _caddf + 5 + 2*_caddf
	_efgdgf := New(_efba*_bbbf, _bbab*_bbbf)
	_egag := &Bitmaps{}
	var _ddbg *int
	_caddf *= _bbbf
	_dgeg := _caddf
	_ddbg = &_dgeg
	_dadeb := _caddf
	_abfc := TstDSymbol(t, scale...)
	TstAddSymbol(t, _egag, _abfc, _ddbg, _dadeb, 1*_bbbf)
	_abfc = TstOSymbol(t, scale...)
	TstAddSymbol(t, _egag, _abfc, _ddbg, _dadeb, _caddf)
	_abfc = TstISymbol(t, scale...)
	TstAddSymbol(t, _egag, _abfc, _ddbg, _dadeb, 1*_bbbf)
	_abfc = TstTSymbol(t, scale...)
	TstAddSymbol(t, _egag, _abfc, _ddbg, _dadeb, _caddf)
	_abfc = TstNSymbol(t, scale...)
	TstAddSymbol(t, _egag, _abfc, _ddbg, _dadeb, 1*_bbbf)
	_abfc = TstOSymbol(t, scale...)
	TstAddSymbol(t, _egag, _abfc, _ddbg, _dadeb, 1*_bbbf)
	_abfc = TstWSymbol(t, scale...)
	TstAddSymbol(t, _egag, _abfc, _ddbg, _dadeb, 0)
	*_ddbg = _caddf
	_dadeb = 5*_bbbf + _caddf
	_abfc = TstOSymbol(t, scale...)
	TstAddSymbol(t, _egag, _abfc, _ddbg, _dadeb, 1*_bbbf)
	_abfc = TstRSymbol(t, scale...)
	TstAddSymbol(t, _egag, _abfc, _ddbg, _dadeb, _caddf)
	_abfc = TstNSymbol(t, scale...)
	TstAddSymbol(t, _egag, _abfc, _ddbg, _dadeb, 1*_bbbf)
	_abfc = TstESymbol(t, scale...)
	TstAddSymbol(t, _egag, _abfc, _ddbg, _dadeb, 1*_bbbf)
	_abfc = TstVSymbol(t, scale...)
	TstAddSymbol(t, _egag, _abfc, _ddbg, _dadeb, 1*_bbbf)
	_abfc = TstESymbol(t, scale...)
	TstAddSymbol(t, _egag, _abfc, _ddbg, _dadeb, 1*_bbbf)
	_abfc = TstRSymbol(t, scale...)
	TstAddSymbol(t, _egag, _abfc, _ddbg, _dadeb, 0)
	TstWriteSymbols(t, _egag, _efgdgf)
	return _efgdgf
}
func _dabec(_gdde *_fc.Stack) (_fdcb *fillSegment, _agafd error) {
	const _dgca = "\u0070\u006f\u0070\u0046\u0069\u006c\u006c\u0053\u0065g\u006d\u0065\u006e\u0074"
	if _gdde == nil {
		return nil, _a.Error(_dgca, "\u006ei\u006c \u0073\u0074\u0061\u0063\u006b \u0070\u0072o\u0076\u0069\u0064\u0065\u0064")
	}
	if _gdde.Aux == nil {
		return nil, _a.Error(_dgca, "a\u0075x\u0053\u0074\u0061\u0063\u006b\u0020\u006e\u006ft\u0020\u0064\u0065\u0066in\u0065\u0064")
	}
	_eaeec, _edf := _gdde.Pop()
	if !_edf {
		return nil, nil
	}
	_efca, _edf := _eaeec.(*fillSegment)
	if !_edf {
		return nil, _a.Error(_dgca, "\u0073\u0074\u0061ck\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020c\u006fn\u0074a\u0069n\u0020\u002a\u0066\u0069\u006c\u006c\u0053\u0065\u0067\u006d\u0065\u006e\u0074")
	}
	_fdcb = &fillSegment{_efca._eeee, _efca._fdddg, _efca._dabg + _efca._bege, _efca._bege}
	_gdde.Aux.Push(_efca)
	return _fdcb, nil
}
func (_abe *Boxes) selectWithIndicator(_abb *_fc.NumSlice) (_fgae *Boxes, _bgaa error) {
	const _bacf = "\u0042o\u0078\u0065\u0073\u002es\u0065\u006c\u0065\u0063\u0074W\u0069t\u0068I\u006e\u0064\u0069\u0063\u0061\u0074\u006fr"
	if _abe == nil {
		return nil, _a.Error(_bacf, "b\u006f\u0078\u0065\u0073 '\u0062'\u0020\u006e\u006f\u0074\u0020d\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if _abb == nil {
		return nil, _a.Error(_bacf, "\u0027\u006ea\u0027\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if len(*_abb) != len(*_abe) {
		return nil, _a.Error(_bacf, "\u0062\u006f\u0078\u0065\u0073\u0020\u0027\u0062\u0027\u0020\u0068\u0061\u0073\u0020\u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020s\u0069\u007a\u0065\u0020\u0074h\u0061\u006e \u0027\u006e\u0061\u0027")
	}
	var _gcfc, _bce int
	for _ffgf := 0; _ffgf < len(*_abb); _ffgf++ {
		if _gcfc, _bgaa = _abb.GetInt(_ffgf); _bgaa != nil {
			return nil, _a.Wrap(_bgaa, _bacf, "\u0063\u0068\u0065\u0063\u006b\u0069\u006e\u0067\u0020c\u006f\u0075\u006e\u0074")
		}
		if _gcfc == 1 {
			_bce++
		}
	}
	if _bce == len(*_abe) {
		return _abe, nil
	}
	_cbca := Boxes{}
	for _deebb := 0; _deebb < len(*_abb); _deebb++ {
		_gcfc = int((*_abb)[_deebb])
		if _gcfc == 0 {
			continue
		}
		_cbca = append(_cbca, (*_abe)[_deebb])
	}
	_fgae = &_cbca
	return _fgae, nil
}
func DilateBrick(d, s *Bitmap, hSize, vSize int) (*Bitmap, error) { return _gebb(d, s, hSize, vSize) }
func (_gcgd *Bitmap) GetBitOffset(x int) int                      { return x & 0x07 }
func (_bgc *Bitmap) GetUnpaddedData() ([]byte, error) {
	_aba := uint(_bgc.Width & 0x07)
	if _aba == 0 {
		return _bgc.Data, nil
	}
	_ded := _bgc.Width * _bgc.Height
	if _ded%8 != 0 {
		_ded >>= 3
		_ded++
	} else {
		_ded >>= 3
	}
	_daf := make([]byte, _ded)
	_ecg := _c.NewWriterMSB(_daf)
	const _cbef = "\u0047e\u0074U\u006e\u0070\u0061\u0064\u0064\u0065\u0064\u0044\u0061\u0074\u0061"
	for _egdb := 0; _egdb < _bgc.Height; _egdb++ {
		for _eeb := 0; _eeb < _bgc.RowStride; _eeb++ {
			_bdab := _bgc.Data[_egdb*_bgc.RowStride+_eeb]
			if _eeb != _bgc.RowStride-1 {
				_bae := _ecg.WriteByte(_bdab)
				if _bae != nil {
					return nil, _a.Wrap(_bae, _cbef, "")
				}
				continue
			}
			for _affb := uint(0); _affb < _aba; _affb++ {
				_daee := _ecg.WriteBit(int(_bdab >> (7 - _affb) & 0x01))
				if _daee != nil {
					return nil, _a.Wrap(_daee, _cbef, "")
				}
			}
		}
	}
	return _daf, nil
}
func (_cbgc *Bitmap) SizesEqual(s *Bitmap) bool {
	if _cbgc == s {
		return true
	}
	if _cbgc.Width != s.Width || _cbgc.Height != s.Height {
		return false
	}
	return true
}
func CorrelationScore(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab []int) (_bacb float64, _gbgc error) {
	const _fcadg = "\u0063\u006fr\u0072\u0065\u006ca\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065"
	if bm1 == nil || bm2 == nil {
		return 0, _a.Error(_fcadg, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0062\u0069\u0074ma\u0070\u0073")
	}
	if tab == nil {
		return 0, _a.Error(_fcadg, "\u0027\u0074\u0061\u0062\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064")
	}
	if area1 <= 0 || area2 <= 0 {
		return 0, _a.Error(_fcadg, "\u0061\u0072\u0065\u0061s\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0067r\u0065a\u0074\u0065\u0072\u0020\u0074\u0068\u0061n\u0020\u0030")
	}
	_cfbd, _agceg := bm1.Width, bm1.Height
	_gegg, _cagc := bm2.Width, bm2.Height
	_dedf := _bgeb(_cfbd - _gegg)
	if _dedf > maxDiffW {
		return 0, nil
	}
	_bggg := _bgeb(_agceg - _cagc)
	if _bggg > maxDiffH {
		return 0, nil
	}
	var _facb, _def int
	if delX >= 0 {
		_facb = int(delX + 0.5)
	} else {
		_facb = int(delX - 0.5)
	}
	if delY >= 0 {
		_def = int(delY + 0.5)
	} else {
		_def = int(delY - 0.5)
	}
	_fdga := _dba(_def, 0)
	_eeeb := _fda(_cagc+_def, _agceg)
	_fadd := bm1.RowStride * _fdga
	_cgfd := bm2.RowStride * (_fdga - _def)
	_eggg := _dba(_facb, 0)
	_dad := _fda(_gegg+_facb, _cfbd)
	_eaaa := bm2.RowStride
	var _efff, _ebed int
	if _facb >= 8 {
		_efff = _facb >> 3
		_fadd += _efff
		_eggg -= _efff << 3
		_dad -= _efff << 3
		_facb &= 7
	} else if _facb <= -8 {
		_ebed = -((_facb + 7) >> 3)
		_cgfd += _ebed
		_eaaa -= _ebed
		_facb += _ebed << 3
	}
	if _eggg >= _dad || _fdga >= _eeeb {
		return 0, nil
	}
	_gfgg := (_dad + 7) >> 3
	var (
		_gcgag, _deb, _dcff   byte
		_dcea, _cfbbc, _dfcff int
	)
	switch {
	case _facb == 0:
		for _dfcff = _fdga; _dfcff < _eeeb; _dfcff, _fadd, _cgfd = _dfcff+1, _fadd+bm1.RowStride, _cgfd+bm2.RowStride {
			for _cfbbc = 0; _cfbbc < _gfgg; _cfbbc++ {
				_dcff = bm1.Data[_fadd+_cfbbc] & bm2.Data[_cgfd+_cfbbc]
				_dcea += tab[_dcff]
			}
		}
	case _facb > 0:
		if _eaaa < _gfgg {
			for _dfcff = _fdga; _dfcff < _eeeb; _dfcff, _fadd, _cgfd = _dfcff+1, _fadd+bm1.RowStride, _cgfd+bm2.RowStride {
				_gcgag, _deb = bm1.Data[_fadd], bm2.Data[_cgfd]>>uint(_facb)
				_dcff = _gcgag & _deb
				_dcea += tab[_dcff]
				for _cfbbc = 1; _cfbbc < _eaaa; _cfbbc++ {
					_gcgag, _deb = bm1.Data[_fadd+_cfbbc], (bm2.Data[_cgfd+_cfbbc]>>uint(_facb))|(bm2.Data[_cgfd+_cfbbc-1]<<uint(8-_facb))
					_dcff = _gcgag & _deb
					_dcea += tab[_dcff]
				}
				_gcgag = bm1.Data[_fadd+_cfbbc]
				_deb = bm2.Data[_cgfd+_cfbbc-1] << uint(8-_facb)
				_dcff = _gcgag & _deb
				_dcea += tab[_dcff]
			}
		} else {
			for _dfcff = _fdga; _dfcff < _eeeb; _dfcff, _fadd, _cgfd = _dfcff+1, _fadd+bm1.RowStride, _cgfd+bm2.RowStride {
				_gcgag, _deb = bm1.Data[_fadd], bm2.Data[_cgfd]>>uint(_facb)
				_dcff = _gcgag & _deb
				_dcea += tab[_dcff]
				for _cfbbc = 1; _cfbbc < _gfgg; _cfbbc++ {
					_gcgag = bm1.Data[_fadd+_cfbbc]
					_deb = (bm2.Data[_cgfd+_cfbbc] >> uint(_facb)) | (bm2.Data[_cgfd+_cfbbc-1] << uint(8-_facb))
					_dcff = _gcgag & _deb
					_dcea += tab[_dcff]
				}
			}
		}
	default:
		if _gfgg < _eaaa {
			for _dfcff = _fdga; _dfcff < _eeeb; _dfcff, _fadd, _cgfd = _dfcff+1, _fadd+bm1.RowStride, _cgfd+bm2.RowStride {
				for _cfbbc = 0; _cfbbc < _gfgg; _cfbbc++ {
					_gcgag = bm1.Data[_fadd+_cfbbc]
					_deb = bm2.Data[_cgfd+_cfbbc] << uint(-_facb)
					_deb |= bm2.Data[_cgfd+_cfbbc+1] >> uint(8+_facb)
					_dcff = _gcgag & _deb
					_dcea += tab[_dcff]
				}
			}
		} else {
			for _dfcff = _fdga; _dfcff < _eeeb; _dfcff, _fadd, _cgfd = _dfcff+1, _fadd+bm1.RowStride, _cgfd+bm2.RowStride {
				for _cfbbc = 0; _cfbbc < _gfgg-1; _cfbbc++ {
					_gcgag = bm1.Data[_fadd+_cfbbc]
					_deb = bm2.Data[_cgfd+_cfbbc] << uint(-_facb)
					_deb |= bm2.Data[_cgfd+_cfbbc+1] >> uint(8+_facb)
					_dcff = _gcgag & _deb
					_dcea += tab[_dcff]
				}
				_gcgag = bm1.Data[_fadd+_cfbbc]
				_deb = bm2.Data[_cgfd+_cfbbc] << uint(-_facb)
				_dcff = _gcgag & _deb
				_dcea += tab[_dcff]
			}
		}
	}
	_bacb = float64(_dcea) * float64(_dcea) / (float64(area1) * float64(area2))
	return _bacb, nil
}
func (_bbge *ClassedPoints) GroupByY() ([]*ClassedPoints, error) {
	const _bdga = "\u0043\u006c\u0061\u0073se\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u002e\u0047\u0072\u006f\u0075\u0070\u0042y\u0059"
	if _ddcf := _bbge.validateIntSlice(); _ddcf != nil {
		return nil, _a.Wrap(_ddcf, _bdga, "")
	}
	if _bbge.IntSlice.Size() == 0 {
		return nil, _a.Error(_bdga, "\u004e\u006f\u0020\u0063la\u0073\u0073\u0065\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064")
	}
	_bbge.SortByY()
	var (
		_agbaa []*ClassedPoints
		_deed  int
	)
	_bcdb := -1
	var _defd *ClassedPoints
	for _ccda := 0; _ccda < len(_bbge.IntSlice); _ccda++ {
		_deed = int(_bbge.YAtIndex(_ccda))
		if _deed != _bcdb {
			_defd = &ClassedPoints{Points: _bbge.Points}
			_bcdb = _deed
			_agbaa = append(_agbaa, _defd)
		}
		_defd.IntSlice = append(_defd.IntSlice, _bbge.IntSlice[_ccda])
	}
	for _, _gcdbc := range _agbaa {
		_gcdbc.SortByX()
	}
	return _agbaa, nil
}
func TstNSymbol(t *_f.T, scale ...int) *Bitmap {
	_cggee, _affa := NewWithData(4, 5, []byte{0x90, 0xD0, 0xB0, 0x90, 0x90})
	_g.NoError(t, _affa)
	return TstGetScaledSymbol(t, _cggee, scale...)
}
func Rect(x, y, w, h int) (*_ac.Rectangle, error) {
	const _fadc = "b\u0069\u0074\u006d\u0061\u0070\u002e\u0052\u0065\u0063\u0074"
	if x < 0 {
		w += x
		x = 0
		if w <= 0 {
			return nil, _a.Errorf(_fadc, "x\u003a\u0027\u0025\u0064\u0027\u0020<\u0020\u0030\u0020\u0061\u006e\u0064\u0020\u0077\u003a \u0027\u0025\u0064'\u0020<\u003d\u0020\u0030", x, w)
		}
	}
	if y < 0 {
		h += y
		y = 0
		if h <= 0 {
			return nil, _a.Error(_fadc, "\u0079\u0020\u003c 0\u0020\u0061\u006e\u0064\u0020\u0062\u006f\u0078\u0020\u006f\u0066\u0066\u0020\u002b\u0071\u0075\u0061\u0064")
		}
	}
	_gff := _ac.Rect(x, y, x+w, y+h)
	return &_gff, nil
}
func TstDSymbol(t *_f.T, scale ...int) *Bitmap {
	_ggbf, _faba := NewWithData(4, 5, []byte{0xf0, 0x90, 0x90, 0x90, 0xE0})
	_g.NoError(t, _faba)
	return TstGetScaledSymbol(t, _ggbf, scale...)
}
func (_fffg *ClassedPoints) Less(i, j int) bool { return _fffg._fbdba(i, j) }

const (
	Vanilla Color = iota
	Chocolate
)

func RankHausTest(p1, p2, p3, p4 *Bitmap, delX, delY float32, maxDiffW, maxDiffH, area1, area3 int, rank float32, tab8 []int) (_fgbaf bool, _ecga error) {
	const _cbbb = "\u0052\u0061\u006ek\u0048\u0061\u0075\u0073\u0054\u0065\u0073\u0074"
	_dgg, _efdg := p1.Width, p1.Height
	_bdag, _dbdc := p3.Width, p3.Height
	if _fc.Abs(_dgg-_bdag) > maxDiffW {
		return false, nil
	}
	if _fc.Abs(_efdg-_dbdc) > maxDiffH {
		return false, nil
	}
	_fdbc := int(float32(area1)*(1.0-rank) + 0.5)
	_bdfgb := int(float32(area3)*(1.0-rank) + 0.5)
	var _ccgb, _dabf int
	if delX >= 0 {
		_ccgb = int(delX + 0.5)
	} else {
		_ccgb = int(delX - 0.5)
	}
	if delY >= 0 {
		_dabf = int(delY + 0.5)
	} else {
		_dabf = int(delY - 0.5)
	}
	_fadg := p1.CreateTemplate()
	if _ecga = _fadg.RasterOperation(0, 0, _dgg, _efdg, PixSrc, p1, 0, 0); _ecga != nil {
		return false, _a.Wrap(_ecga, _cbbb, "p\u0031\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _ecga = _fadg.RasterOperation(_ccgb, _dabf, _dgg, _efdg, PixNotSrcAndDst, p4, 0, 0); _ecga != nil {
		return false, _a.Wrap(_ecga, _cbbb, "\u0074 \u0026\u0020\u0021\u0070\u0034")
	}
	_fgbaf, _ecga = _fadg.ThresholdPixelSum(_fdbc, tab8)
	if _ecga != nil {
		return false, _a.Wrap(_ecga, _cbbb, "\u0074\u002d\u003e\u0074\u0068\u0072\u0065\u0073\u0068\u0031")
	}
	if _fgbaf {
		return false, nil
	}
	if _ecga = _fadg.RasterOperation(_ccgb, _dabf, _bdag, _dbdc, PixSrc, p3, 0, 0); _ecga != nil {
		return false, _a.Wrap(_ecga, _cbbb, "p\u0033\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _ecga = _fadg.RasterOperation(0, 0, _bdag, _dbdc, PixNotSrcAndDst, p2, 0, 0); _ecga != nil {
		return false, _a.Wrap(_ecga, _cbbb, "\u0074 \u0026\u0020\u0021\u0070\u0032")
	}
	_fgbaf, _ecga = _fadg.ThresholdPixelSum(_bdfgb, tab8)
	if _ecga != nil {
		return false, _a.Wrap(_ecga, _cbbb, "\u0074\u002d\u003e\u0074\u0068\u0072\u0065\u0073\u0068\u0033")
	}
	return !_fgbaf, nil
}
func (_dcda *Bitmap) resizeImageData(_cde *Bitmap) error {
	if _cde == nil {
		return _a.Error("\u0072e\u0073i\u007a\u0065\u0049\u006d\u0061\u0067\u0065\u0044\u0061\u0074\u0061", "\u0073r\u0063 \u0069\u0073\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _dcda.SizesEqual(_cde) {
		return nil
	}
	_dcda.Data = make([]byte, len(_cde.Data))
	_dcda.Width = _cde.Width
	_dcda.Height = _cde.Height
	_dcda.RowStride = _cde.RowStride
	return nil
}
func (_ffgd *byWidth) Swap(i, j int) {
	_ffgd.Values[i], _ffgd.Values[j] = _ffgd.Values[j], _ffgd.Values[i]
	if _ffgd.Boxes != nil {
		_ffgd.Boxes[i], _ffgd.Boxes[j] = _ffgd.Boxes[j], _ffgd.Boxes[i]
	}
}
func (_bcea *Bitmaps) WidthSorter() func(_daeed, _feac int) bool {
	return func(_acad, _agfc int) bool { return _bcea.Values[_acad].Width < _bcea.Values[_agfc].Width }
}
func _gddb(_eaeb *Bitmap, _ddaeg *Bitmap, _adda *Selection, _beg **Bitmap) (*Bitmap, error) {
	const _cdag = "\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u004d\u006f\u0072\u0070\u0068A\u0072\u0067\u0073\u0031"
	if _ddaeg == nil {
		return nil, _a.Error(_cdag, "\u004d\u006f\u0072\u0070\u0068\u0041\u0072\u0067\u0073\u0031\u0020'\u0073\u0027\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066i\u006e\u0065\u0064")
	}
	if _adda == nil {
		return nil, _a.Error(_cdag, "\u004d\u006f\u0072\u0068p\u0041\u0072\u0067\u0073\u0031\u0020\u0027\u0073\u0065\u006c'\u0020n\u006f\u0074\u0020\u0064\u0065\u0066\u0069n\u0065\u0064")
	}
	_edad, _gage := _adda.Height, _adda.Width
	if _edad == 0 || _gage == 0 {
		return nil, _a.Error(_cdag, "\u0073\u0065\u006c\u0065ct\u0069\u006f\u006e\u0020\u006f\u0066\u0020\u0073\u0069\u007a\u0065\u0020\u0030")
	}
	if _eaeb == nil {
		_eaeb = _ddaeg.createTemplate()
		*_beg = _ddaeg
		return _eaeb, nil
	}
	_eaeb.Width = _ddaeg.Width
	_eaeb.Height = _ddaeg.Height
	_eaeb.RowStride = _ddaeg.RowStride
	_eaeb.Color = _ddaeg.Color
	_eaeb.Data = make([]byte, _ddaeg.RowStride*_ddaeg.Height)
	if _eaeb == _ddaeg {
		*_beg = _ddaeg.Copy()
	} else {
		*_beg = _ddaeg
	}
	return _eaeb, nil
}
func _fed(_efbf *Bitmap, _bcde, _fedf int, _adae, _bdgc int, _adabb RasterOperator) {
	var (
		_eacg        int
		_efed        byte
		_bbgc, _ddaa int
		_gcfae       int
	)
	_acae := _adae >> 3
	_gddd := _adae & 7
	if _gddd > 0 {
		_efed = _fcgba[_gddd]
	}
	_eacg = _efbf.RowStride*_fedf + (_bcde >> 3)
	switch _adabb {
	case PixClr:
		for _bbgc = 0; _bbgc < _bdgc; _bbgc++ {
			_gcfae = _eacg + _bbgc*_efbf.RowStride
			for _ddaa = 0; _ddaa < _acae; _ddaa++ {
				_efbf.Data[_gcfae] = 0x0
				_gcfae++
			}
			if _gddd > 0 {
				_efbf.Data[_gcfae] = _cbcaa(_efbf.Data[_gcfae], 0x0, _efed)
			}
		}
	case PixSet:
		for _bbgc = 0; _bbgc < _bdgc; _bbgc++ {
			_gcfae = _eacg + _bbgc*_efbf.RowStride
			for _ddaa = 0; _ddaa < _acae; _ddaa++ {
				_efbf.Data[_gcfae] = 0xff
				_gcfae++
			}
			if _gddd > 0 {
				_efbf.Data[_gcfae] = _cbcaa(_efbf.Data[_gcfae], 0xff, _efed)
			}
		}
	case PixNotDst:
		for _bbgc = 0; _bbgc < _bdgc; _bbgc++ {
			_gcfae = _eacg + _bbgc*_efbf.RowStride
			for _ddaa = 0; _ddaa < _acae; _ddaa++ {
				_efbf.Data[_gcfae] = ^_efbf.Data[_gcfae]
				_gcfae++
			}
			if _gddd > 0 {
				_efbf.Data[_gcfae] = _cbcaa(_efbf.Data[_gcfae], ^_efbf.Data[_gcfae], _efed)
			}
		}
	}
}
func (_afag *Bitmap) RemoveBorder(borderSize int) (*Bitmap, error) {
	if borderSize == 0 {
		return _afag.Copy(), nil
	}
	_aab, _ade := _afag.removeBorderGeneral(borderSize, borderSize, borderSize, borderSize)
	if _ade != nil {
		return nil, _a.Wrap(_ade, "\u0052\u0065\u006do\u0076\u0065\u0042\u006f\u0072\u0064\u0065\u0072", "")
	}
	return _aab, nil
}
func (_efga *Bitmap) createTemplate() *Bitmap {
	return &Bitmap{Width: _efga.Width, Height: _efga.Height, RowStride: _efga.RowStride, Color: _efga.Color, Text: _efga.Text, BitmapNumber: _efga.BitmapNumber, Special: _efga.Special, Data: make([]byte, len(_efga.Data))}
}

const (
	MopDilation MorphOperation = iota
	MopErosion
	MopOpening
	MopClosing
	MopRankBinaryReduction
	MopReplicativeBinaryExpansion
	MopAddBorder
)

func TstRSymbol(t *_f.T, scale ...int) *Bitmap {
	_bgedc, _bdfgc := NewWithData(4, 5, []byte{0xF0, 0x90, 0xF0, 0xA0, 0x90})
	_g.NoError(t, _bdfgc)
	return TstGetScaledSymbol(t, _bgedc, scale...)
}
func _ebca(_dggb, _dacg int, _gacb string) *Selection {
	_bacdc := &Selection{Height: _dggb, Width: _dacg, Name: _gacb}
	_bacdc.Data = make([][]SelectionValue, _dggb)
	for _eedcb := 0; _eedcb < _dggb; _eedcb++ {
		_bacdc.Data[_eedcb] = make([]SelectionValue, _dacg)
	}
	return _bacdc
}

type LocationFilter int

func (_bgeda *Bitmaps) SortByHeight() { _gbce := (*byHeight)(_bgeda); _e.Sort(_gbce) }

type BoundaryCondition int

func (_deaa *Bitmap) setEightFullBytes(_dfd int, _abd uint64) error {
	if _dfd+7 > len(_deaa.Data)-1 {
		return _a.Error("\u0073\u0065\u0074\u0045\u0069\u0067\u0068\u0074\u0042\u0079\u0074\u0065\u0073", "\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065")
	}
	_deaa.Data[_dfd] = byte((_abd & 0xff00000000000000) >> 56)
	_deaa.Data[_dfd+1] = byte((_abd & 0xff000000000000) >> 48)
	_deaa.Data[_dfd+2] = byte((_abd & 0xff0000000000) >> 40)
	_deaa.Data[_dfd+3] = byte((_abd & 0xff00000000) >> 32)
	_deaa.Data[_dfd+4] = byte((_abd & 0xff000000) >> 24)
	_deaa.Data[_dfd+5] = byte((_abd & 0xff0000) >> 16)
	_deaa.Data[_dfd+6] = byte((_abd & 0xff00) >> 8)
	_deaa.Data[_dfd+7] = byte(_abd & 0xff)
	return nil
}

type RasterOperator int

func (_bdgad *Bitmaps) ClipToBitmap(s *Bitmap) (*Bitmaps, error) {
	const _afec = "B\u0069t\u006d\u0061\u0070\u0073\u002e\u0043\u006c\u0069p\u0054\u006f\u0042\u0069tm\u0061\u0070"
	if _bdgad == nil {
		return nil, _a.Error(_afec, "\u0042\u0069\u0074\u006dap\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if s == nil {
		return nil, _a.Error(_afec, "\u0073o\u0075\u0072\u0063\u0065 \u0062\u0069\u0074\u006d\u0061p\u0020n\u006ft\u0020\u0064\u0065\u0066\u0069\u006e\u0065d")
	}
	_dggg := len(_bdgad.Values)
	_egac := &Bitmaps{Values: make([]*Bitmap, _dggg), Boxes: make([]*_ac.Rectangle, _dggg)}
	var (
		_daefb, _cffb *Bitmap
		_cdbf         *_ac.Rectangle
		_ddee         error
	)
	for _deced := 0; _deced < _dggg; _deced++ {
		if _daefb, _ddee = _bdgad.GetBitmap(_deced); _ddee != nil {
			return nil, _a.Wrap(_ddee, _afec, "")
		}
		if _cdbf, _ddee = _bdgad.GetBox(_deced); _ddee != nil {
			return nil, _a.Wrap(_ddee, _afec, "")
		}
		if _cffb, _ddee = s.clipRectangle(_cdbf, nil); _ddee != nil {
			return nil, _a.Wrap(_ddee, _afec, "")
		}
		if _cffb, _ddee = _cffb.And(_daefb); _ddee != nil {
			return nil, _a.Wrap(_ddee, _afec, "")
		}
		_egac.Values[_deced] = _cffb
		_egac.Boxes[_deced] = _cdbf
	}
	return _egac, nil
}
func _ga(_ace, _gg *Bitmap) (_gaa error) {
	const _gae = "\u0065\u0078\u0070\u0061nd\u0042\u0069\u006e\u0061\u0072\u0079\u0046\u0061\u0063\u0074\u006f\u0072\u0032"
	_bb := _gg.RowStride
	_cf := _ace.RowStride
	var (
		_de                      byte
		_dc                      uint16
		_dbf, _af, _ff, _cd, _ag int
	)
	for _ff = 0; _ff < _gg.Height; _ff++ {
		_dbf = _ff * _bb
		_af = 2 * _ff * _cf
		for _cd = 0; _cd < _bb; _cd++ {
			_de = _gg.Data[_dbf+_cd]
			_dc = _fafe[_de]
			_ag = _af + _cd*2
			if _ace.RowStride != _gg.RowStride*2 && (_cd+1)*2 > _ace.RowStride {
				_gaa = _ace.SetByte(_ag, byte(_dc>>8))
			} else {
				_gaa = _ace.setTwoBytes(_ag, _dc)
			}
			if _gaa != nil {
				return _a.Wrap(_gaa, _gae, "")
			}
		}
		for _cd = 0; _cd < _cf; _cd++ {
			_ag = _af + _cf + _cd
			_de = _ace.Data[_af+_cd]
			if _gaa = _ace.SetByte(_ag, _de); _gaa != nil {
				return _a.Wrapf(_gaa, _gae, "c\u006f\u0070\u0079\u0020\u0064\u006fu\u0062\u006c\u0065\u0064\u0020\u006ci\u006e\u0065\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0042\u0079\u0074\u0065\u003a\u0020\u0027\u0025d\u0027", _af+_cd, _af+_cf+_cd)
			}
		}
	}
	return nil
}
func CorrelationScoreThresholded(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab, downcount []int, scoreThreshold float32) (bool, error) {
	const _bcc = "C\u006f\u0072\u0072\u0065\u006c\u0061t\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054h\u0072\u0065\u0073h\u006fl\u0064\u0065\u0064"
	if bm1 == nil {
		return false, _a.Error(_bcc, "\u0063\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u0062\u006d1\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if bm2 == nil {
		return false, _a.Error(_bcc, "\u0063\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u0062\u006d2\u0020\u0069s\u0020\u006e\u0069\u006c")
	}
	if area1 <= 0 || area2 <= 0 {
		return false, _a.Error(_bcc, "c\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0053\u0063\u006f\u0072\u0065\u0054\u0068re\u0073\u0068\u006f\u006cd\u0065\u0064\u0020\u002d\u0020\u0061\u0072\u0065\u0061s \u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u003e\u0020\u0030")
	}
	if downcount == nil {
		return false, _a.Error(_bcc, "\u0070\u0072\u006fvi\u0064\u0065\u0064\u0020\u006e\u006f\u0020\u0027\u0064\u006f\u0077\u006e\u0063\u006f\u0075\u006e\u0074\u0027")
	}
	if tab == nil {
		return false, _a.Error(_bcc, "p\u0072\u006f\u0076\u0069de\u0064 \u006e\u0069\u006c\u0020\u0027s\u0075\u006d\u0074\u0061\u0062\u0027")
	}
	_cadfa, _gfab := bm1.Width, bm1.Height
	_bcaeg, _eeef := bm2.Width, bm2.Height
	if _fc.Abs(_cadfa-_bcaeg) > maxDiffW {
		return false, nil
	}
	if _fc.Abs(_gfab-_eeef) > maxDiffH {
		return false, nil
	}
	_feb := int(delX + _fc.Sign(delX)*0.5)
	_bbeg := int(delY + _fc.Sign(delY)*0.5)
	_dbac := int(_ce.Ceil(_ce.Sqrt(float64(scoreThreshold) * float64(area1) * float64(area2))))
	_fcga := bm2.RowStride
	_ddba := _dba(_bbeg, 0)
	_cecc := _fda(_eeef+_bbeg, _gfab)
	_agbbc := bm1.RowStride * _ddba
	_bged := bm2.RowStride * (_ddba - _bbeg)
	var _ecde int
	if _cecc <= _gfab {
		_ecde = downcount[_cecc-1]
	}
	_aadf := _dba(_feb, 0)
	_ebfcd := _fda(_bcaeg+_feb, _cadfa)
	var _aefc, _geadb int
	if _feb >= 8 {
		_aefc = _feb >> 3
		_agbbc += _aefc
		_aadf -= _aefc << 3
		_ebfcd -= _aefc << 3
		_feb &= 7
	} else if _feb <= -8 {
		_geadb = -((_feb + 7) >> 3)
		_bged += _geadb
		_fcga -= _geadb
		_feb += _geadb << 3
	}
	var (
		_ebdf, _cffe, _gdeb   int
		_fcbge, _fbgee, _adbf byte
	)
	if _aadf >= _ebfcd || _ddba >= _cecc {
		return false, nil
	}
	_cfeb := (_ebfcd + 7) >> 3
	switch {
	case _feb == 0:
		for _cffe = _ddba; _cffe < _cecc; _cffe, _agbbc, _bged = _cffe+1, _agbbc+bm1.RowStride, _bged+bm2.RowStride {
			for _gdeb = 0; _gdeb < _cfeb; _gdeb++ {
				_fcbge = bm1.Data[_agbbc+_gdeb] & bm2.Data[_bged+_gdeb]
				_ebdf += tab[_fcbge]
			}
			if _ebdf >= _dbac {
				return true, nil
			}
			if _ccca := _ebdf + downcount[_cffe] - _ecde; _ccca < _dbac {
				return false, nil
			}
		}
	case _feb > 0 && _fcga < _cfeb:
		for _cffe = _ddba; _cffe < _cecc; _cffe, _agbbc, _bged = _cffe+1, _agbbc+bm1.RowStride, _bged+bm2.RowStride {
			_fbgee = bm1.Data[_agbbc]
			_adbf = bm2.Data[_bged] >> uint(_feb)
			_fcbge = _fbgee & _adbf
			_ebdf += tab[_fcbge]
			for _gdeb = 1; _gdeb < _fcga; _gdeb++ {
				_fbgee = bm1.Data[_agbbc+_gdeb]
				_adbf = bm2.Data[_bged+_gdeb]>>uint(_feb) | bm2.Data[_bged+_gdeb-1]<<uint(8-_feb)
				_fcbge = _fbgee & _adbf
				_ebdf += tab[_fcbge]
			}
			_fbgee = bm1.Data[_agbbc+_gdeb]
			_adbf = bm2.Data[_bged+_gdeb-1] << uint(8-_feb)
			_fcbge = _fbgee & _adbf
			_ebdf += tab[_fcbge]
			if _ebdf >= _dbac {
				return true, nil
			} else if _ebdf+downcount[_cffe]-_ecde < _dbac {
				return false, nil
			}
		}
	case _feb > 0 && _fcga >= _cfeb:
		for _cffe = _ddba; _cffe < _cecc; _cffe, _agbbc, _bged = _cffe+1, _agbbc+bm1.RowStride, _bged+bm2.RowStride {
			_fbgee = bm1.Data[_agbbc]
			_adbf = bm2.Data[_bged] >> uint(_feb)
			_fcbge = _fbgee & _adbf
			_ebdf += tab[_fcbge]
			for _gdeb = 1; _gdeb < _cfeb; _gdeb++ {
				_fbgee = bm1.Data[_agbbc+_gdeb]
				_adbf = bm2.Data[_bged+_gdeb] >> uint(_feb)
				_adbf |= bm2.Data[_bged+_gdeb-1] << uint(8-_feb)
				_fcbge = _fbgee & _adbf
				_ebdf += tab[_fcbge]
			}
			if _ebdf >= _dbac {
				return true, nil
			} else if _ebdf+downcount[_cffe]-_ecde < _dbac {
				return false, nil
			}
		}
	case _cfeb < _fcga:
		for _cffe = _ddba; _cffe < _cecc; _cffe, _agbbc, _bged = _cffe+1, _agbbc+bm1.RowStride, _bged+bm2.RowStride {
			for _gdeb = 0; _gdeb < _cfeb; _gdeb++ {
				_fbgee = bm1.Data[_agbbc+_gdeb]
				_adbf = bm2.Data[_bged+_gdeb] << uint(-_feb)
				_adbf |= bm2.Data[_bged+_gdeb+1] >> uint(8+_feb)
				_fcbge = _fbgee & _adbf
				_ebdf += tab[_fcbge]
			}
			if _ebdf >= _dbac {
				return true, nil
			} else if _debf := _ebdf + downcount[_cffe] - _ecde; _debf < _dbac {
				return false, nil
			}
		}
	case _fcga >= _cfeb:
		for _cffe = _ddba; _cffe < _cecc; _cffe, _agbbc, _bged = _cffe+1, _agbbc+bm1.RowStride, _bged+bm2.RowStride {
			for _gdeb = 0; _gdeb < _cfeb; _gdeb++ {
				_fbgee = bm1.Data[_agbbc+_gdeb]
				_adbf = bm2.Data[_bged+_gdeb] << uint(-_feb)
				_adbf |= bm2.Data[_bged+_gdeb+1] >> uint(8+_feb)
				_fcbge = _fbgee & _adbf
				_ebdf += tab[_fcbge]
			}
			_fbgee = bm1.Data[_agbbc+_gdeb]
			_adbf = bm2.Data[_bged+_gdeb] << uint(-_feb)
			_fcbge = _fbgee & _adbf
			_ebdf += tab[_fcbge]
			if _ebdf >= _dbac {
				return true, nil
			} else if _ebdf+downcount[_cffe]-_ecde < _dbac {
				return false, nil
			}
		}
	}
	_aega := float32(_ebdf) * float32(_ebdf) / (float32(area1) * float32(area2))
	if _aega >= scoreThreshold {
		_db.Log.Trace("\u0063\u006f\u0075\u006e\u0074\u003a\u0020\u0025\u0064\u0020\u003c\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u006f\u006cd\u0020\u0025\u0064\u0020\u0062\u0075\u0074\u0020\u0073c\u006f\u0072\u0065\u0020\u0025\u0066\u0020\u003e\u003d\u0020\u0073\u0063\u006fr\u0065\u0054\u0068\u0072\u0065\u0073h\u006f\u006c\u0064 \u0025\u0066", _ebdf, _dbac, _aega, scoreThreshold)
	}
	return false, nil
}
func _dfdf(_ddfgd, _efac, _cea *Bitmap) (*Bitmap, error) {
	const _fdb = "\u0073\u0075\u0062\u0074\u0072\u0061\u0063\u0074"
	if _efac == nil {
		return nil, _a.Error(_fdb, "'\u0073\u0031\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	if _cea == nil {
		return nil, _a.Error(_fdb, "'\u0073\u0032\u0027\u0020\u0069\u0073\u0020\u006e\u0069\u006c")
	}
	var _aefb error
	switch {
	case _ddfgd == _efac:
		if _aefb = _ddfgd.RasterOperation(0, 0, _efac.Width, _efac.Height, PixNotSrcAndDst, _cea, 0, 0); _aefb != nil {
			return nil, _a.Wrap(_aefb, _fdb, "\u0064 \u003d\u003d\u0020\u0073\u0031")
		}
	case _ddfgd == _cea:
		if _aefb = _ddfgd.RasterOperation(0, 0, _efac.Width, _efac.Height, PixNotSrcAndDst, _efac, 0, 0); _aefb != nil {
			return nil, _a.Wrap(_aefb, _fdb, "\u0064 \u003d\u003d\u0020\u0073\u0032")
		}
	default:
		_ddfgd, _aefb = _bggbg(_ddfgd, _efac)
		if _aefb != nil {
			return nil, _a.Wrap(_aefb, _fdb, "")
		}
		if _aefb = _ddfgd.RasterOperation(0, 0, _efac.Width, _efac.Height, PixNotSrcAndDst, _cea, 0, 0); _aefb != nil {
			return nil, _a.Wrap(_aefb, _fdb, "\u0064e\u0066\u0061\u0075\u006c\u0074")
		}
	}
	return _ddfgd, nil
}
func _dfdfd(_acgae, _gccd *Bitmap, _gcecc, _aaga int) (_daed error) {
	const _bdaee = "\u0073e\u0065d\u0066\u0069\u006c\u006c\u0042i\u006e\u0061r\u0079\u004c\u006f\u0077\u0038"
	var (
		_fedb, _aggb, _dcag, _dcdf                                 int
		_bbef, _ffcce, _ffag, _fagac, _geca, _aegac, _efgf, _bcege byte
	)
	for _fedb = 0; _fedb < _gcecc; _fedb++ {
		_dcag = _fedb * _acgae.RowStride
		_dcdf = _fedb * _gccd.RowStride
		for _aggb = 0; _aggb < _aaga; _aggb++ {
			if _bbef, _daed = _acgae.GetByte(_dcag + _aggb); _daed != nil {
				return _a.Wrap(_daed, _bdaee, "\u0067e\u0074 \u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
			}
			if _ffcce, _daed = _gccd.GetByte(_dcdf + _aggb); _daed != nil {
				return _a.Wrap(_daed, _bdaee, "\u0067\u0065\u0074\u0020\u006d\u0061\u0073\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _fedb > 0 {
				if _ffag, _daed = _acgae.GetByte(_dcag - _acgae.RowStride + _aggb); _daed != nil {
					return _a.Wrap(_daed, _bdaee, "\u0069\u0020\u003e\u0020\u0030\u0020\u0062\u0079\u0074\u0065")
				}
				_bbef |= _ffag | (_ffag << 1) | (_ffag >> 1)
				if _aggb > 0 {
					if _bcege, _daed = _acgae.GetByte(_dcag - _acgae.RowStride + _aggb - 1); _daed != nil {
						return _a.Wrap(_daed, _bdaee, "\u0069\u0020\u003e\u00200 \u0026\u0026\u0020\u006a\u0020\u003e\u0020\u0030\u0020\u0062\u0079\u0074\u0065")
					}
					_bbef |= _bcege << 7
				}
				if _aggb < _aaga-1 {
					if _bcege, _daed = _acgae.GetByte(_dcag - _acgae.RowStride + _aggb + 1); _daed != nil {
						return _a.Wrap(_daed, _bdaee, "\u006a\u0020<\u0020\u0077\u0070l\u0020\u002d\u0020\u0031\u0020\u0062\u0079\u0074\u0065")
					}
					_bbef |= _bcege >> 7
				}
			}
			if _aggb > 0 {
				if _fagac, _daed = _acgae.GetByte(_dcag + _aggb - 1); _daed != nil {
					return _a.Wrap(_daed, _bdaee, "\u006a\u0020\u003e \u0030")
				}
				_bbef |= _fagac << 7
			}
			_bbef &= _ffcce
			if _bbef == 0 || ^_bbef == 0 {
				if _daed = _acgae.SetByte(_dcag+_aggb, _bbef); _daed != nil {
					return _a.Wrap(_daed, _bdaee, "\u0073e\u0074t\u0069\u006e\u0067\u0020\u0065m\u0070\u0074y\u0020\u0062\u0079\u0074\u0065")
				}
			}
			for {
				_efgf = _bbef
				_bbef = (_bbef | (_bbef >> 1) | (_bbef << 1)) & _ffcce
				if (_bbef ^ _efgf) == 0 {
					if _daed = _acgae.SetByte(_dcag+_aggb, _bbef); _daed != nil {
						return _a.Wrap(_daed, _bdaee, "\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0070\u0072\u0065\u0076 \u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	for _fedb = _gcecc - 1; _fedb >= 0; _fedb-- {
		_dcag = _fedb * _acgae.RowStride
		_dcdf = _fedb * _gccd.RowStride
		for _aggb = _aaga - 1; _aggb >= 0; _aggb-- {
			if _bbef, _daed = _acgae.GetByte(_dcag + _aggb); _daed != nil {
				return _a.Wrap(_daed, _bdaee, "\u0072\u0065\u0076er\u0073\u0065\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
			}
			if _ffcce, _daed = _gccd.GetByte(_dcdf + _aggb); _daed != nil {
				return _a.Wrap(_daed, _bdaee, "r\u0065\u0076\u0065\u0072se\u0020g\u0065\u0074\u0020\u006d\u0061s\u006b\u0020\u0062\u0079\u0074\u0065")
			}
			if _fedb < _gcecc-1 {
				if _geca, _daed = _acgae.GetByte(_dcag + _acgae.RowStride + _aggb); _daed != nil {
					return _a.Wrap(_daed, _bdaee, "\u0069\u0020\u003c\u0020h\u0020\u002d\u0020\u0031\u0020\u002d\u003e\u0020\u0067\u0065t\u0020s\u006f\u0075\u0072\u0063\u0065\u0020\u0062y\u0074\u0065")
				}
				_bbef |= _geca | (_geca << 1) | _geca>>1
				if _aggb > 0 {
					if _bcege, _daed = _acgae.GetByte(_dcag + _acgae.RowStride + _aggb - 1); _daed != nil {
						return _a.Wrap(_daed, _bdaee, "\u0069\u0020\u003c h\u002d\u0031\u0020\u0026\u0020\u006a\u0020\u003e\u00200\u0020-\u003e \u0067e\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0062\u0079\u0074\u0065")
					}
					_bbef |= _bcege << 7
				}
				if _aggb < _aaga-1 {
					if _bcege, _daed = _acgae.GetByte(_dcag + _acgae.RowStride + _aggb + 1); _daed != nil {
						return _a.Wrap(_daed, _bdaee, "\u0069\u0020\u003c\u0020\u0068\u002d\u0031\u0020\u0026\u0026\u0020\u006a\u0020\u003c\u0077\u0070\u006c\u002d\u0031\u0020\u002d\u003e\u0020\u0067e\u0074\u0020\u0073\u006f\u0075r\u0063\u0065 \u0062\u0079\u0074\u0065")
					}
					_bbef |= _bcege >> 7
				}
			}
			if _aggb < _aaga-1 {
				if _aegac, _daed = _acgae.GetByte(_dcag + _aggb + 1); _daed != nil {
					return _a.Wrap(_daed, _bdaee, "\u006a\u0020<\u0020\u0077\u0070\u006c\u0020\u002d\u0031\u0020\u002d\u003e\u0020\u0067\u0065\u0074\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u0020by\u0074\u0065")
				}
				_bbef |= _aegac >> 7
			}
			_bbef &= _ffcce
			if _bbef == 0 || (^_bbef) == 0 {
				if _daed = _acgae.SetByte(_dcag+_aggb, _bbef); _daed != nil {
					return _a.Wrap(_daed, _bdaee, "\u0073e\u0074 \u006d\u0061\u0073\u006b\u0065\u0064\u0020\u0062\u0079\u0074\u0065")
				}
			}
			for {
				_efgf = _bbef
				_bbef = (_bbef | (_bbef >> 1) | (_bbef << 1)) & _ffcce
				if (_bbef ^ _efgf) == 0 {
					if _daed = _acgae.SetByte(_dcag+_aggb, _bbef); _daed != nil {
						return _a.Wrap(_daed, _bdaee, "r\u0065\u0076\u0065\u0072se\u0020s\u0065\u0074\u0020\u0070\u0072e\u0076\u0020\u0062\u0079\u0074\u0065")
					}
					break
				}
			}
		}
	}
	return nil
}
func TstFrameBitmap() *Bitmap { return _eefb.Copy() }
func _cfcb() (_bfd [256]uint16) {
	for _aade := 0; _aade < 256; _aade++ {
		if _aade&0x01 != 0 {
			_bfd[_aade] |= 0x3
		}
		if _aade&0x02 != 0 {
			_bfd[_aade] |= 0xc
		}
		if _aade&0x04 != 0 {
			_bfd[_aade] |= 0x30
		}
		if _aade&0x08 != 0 {
			_bfd[_aade] |= 0xc0
		}
		if _aade&0x10 != 0 {
			_bfd[_aade] |= 0x300
		}
		if _aade&0x20 != 0 {
			_bfd[_aade] |= 0xc00
		}
		if _aade&0x40 != 0 {
			_bfd[_aade] |= 0x3000
		}
		if _aade&0x80 != 0 {
			_bfd[_aade] |= 0xc000
		}
	}
	return _bfd
}
func HausTest(p1, p2, p3, p4 *Bitmap, delX, delY float32, maxDiffW, maxDiffH int) (bool, error) {
	const _dgeb = "\u0048\u0061\u0075\u0073\u0054\u0065\u0073\u0074"
	_cebc, _fcca := p1.Width, p1.Height
	_befe, _efag := p3.Width, p3.Height
	if _fc.Abs(_cebc-_befe) > maxDiffW {
		return false, nil
	}
	if _fc.Abs(_fcca-_efag) > maxDiffH {
		return false, nil
	}
	_faga := int(delX + _fc.Sign(delX)*0.5)
	_abda := int(delY + _fc.Sign(delY)*0.5)
	var _fde error
	_cbd := p1.CreateTemplate()
	if _fde = _cbd.RasterOperation(0, 0, _cebc, _fcca, PixSrc, p1, 0, 0); _fde != nil {
		return false, _a.Wrap(_fde, _dgeb, "p\u0031\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _fde = _cbd.RasterOperation(_faga, _abda, _cebc, _fcca, PixNotSrcAndDst, p4, 0, 0); _fde != nil {
		return false, _a.Wrap(_fde, _dgeb, "\u0021p\u0034\u0020\u0026\u0020\u0074")
	}
	if _cbd.Zero() {
		return false, nil
	}
	if _fde = _cbd.RasterOperation(_faga, _abda, _befe, _efag, PixSrc, p3, 0, 0); _fde != nil {
		return false, _a.Wrap(_fde, _dgeb, "p\u0033\u0020\u002d\u0053\u0052\u0043\u002d\u003e\u0020\u0074")
	}
	if _fde = _cbd.RasterOperation(0, 0, _befe, _efag, PixNotSrcAndDst, p2, 0, 0); _fde != nil {
		return false, _a.Wrap(_fde, _dgeb, "\u0021p\u0032\u0020\u0026\u0020\u0074")
	}
	return _cbd.Zero(), nil
}
func (_bgcfc *Bitmap) GetComponents(components Component, maxWidth, maxHeight int) (_bdabe *Bitmaps, _fff *Boxes, _eadf error) {
	const _bgda = "B\u0069t\u006d\u0061\u0070\u002e\u0047\u0065\u0074\u0043o\u006d\u0070\u006f\u006een\u0074\u0073"
	if _bgcfc == nil {
		return nil, nil, _a.Error(_bgda, "\u0073\u006f\u0075\u0072\u0063\u0065\u0020\u0042\u0069\u0074\u006da\u0070\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069n\u0065\u0064\u002e")
	}
	switch components {
	case ComponentConn, ComponentCharacters, ComponentWords:
	default:
		return nil, nil, _a.Error(_bgda, "\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074s\u0020\u0070\u0061\u0072\u0061\u006d\u0065t\u0065\u0072")
	}
	if _bgcfc.Zero() {
		_fff = &Boxes{}
		_bdabe = &Bitmaps{}
		return _bdabe, _fff, nil
	}
	switch components {
	case ComponentConn:
		_bdabe = &Bitmaps{}
		if _fff, _eadf = _bgcfc.ConnComponents(_bdabe, 8); _eadf != nil {
			return nil, nil, _a.Wrap(_eadf, _bgda, "\u006e\u006f \u0070\u0072\u0065p\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
	case ComponentCharacters:
		_bgbb, _cccf := MorphSequence(_bgcfc, MorphProcess{Operation: MopClosing, Arguments: []int{1, 6}})
		if _cccf != nil {
			return nil, nil, _a.Wrap(_cccf, _bgda, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
		if _db.Log.IsLogLevel(_db.LogLevelTrace) {
			_db.Log.Trace("\u0043o\u006d\u0070o\u006e\u0065\u006e\u0074C\u0068\u0061\u0072a\u0063\u0074\u0065\u0072\u0073\u0020\u0062\u0069\u0074ma\u0070\u0020\u0061f\u0074\u0065r\u0020\u0063\u006c\u006f\u0073\u0069n\u0067\u003a \u0025\u0073", _bgbb.String())
		}
		_bba := &Bitmaps{}
		_fff, _cccf = _bgbb.ConnComponents(_bba, 8)
		if _cccf != nil {
			return nil, nil, _a.Wrap(_cccf, _bgda, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
		if _db.Log.IsLogLevel(_db.LogLevelTrace) {
			_db.Log.Trace("\u0043\u006f\u006d\u0070\u006f\u006ee\u006e\u0074\u0043\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u0020a\u0066\u0074\u0065\u0072\u0020\u0063\u006f\u006e\u006e\u0065\u0063\u0074\u0069\u0076i\u0074y\u003a\u0020\u0025\u0073", _bba.String())
		}
		if _bdabe, _cccf = _bba.ClipToBitmap(_bgcfc); _cccf != nil {
			return nil, nil, _a.Wrap(_cccf, _bgda, "\u0063h\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069\u006e\u0067")
		}
	case ComponentWords:
		_ccef := 1
		var _dbedc *Bitmap
		switch {
		case _bgcfc.XResolution <= 200:
			_dbedc = _bgcfc
		case _bgcfc.XResolution <= 400:
			_ccef = 2
			_dbedc, _eadf = _cec(_bgcfc, 1, 0, 0, 0)
			if _eadf != nil {
				return nil, nil, _a.Wrap(_eadf, _bgda, "w\u006f\u0072\u0064\u0020\u0070\u0072e\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0020\u002d \u0078\u0072\u0065s\u003c=\u0034\u0030\u0030")
			}
		default:
			_ccef = 4
			_dbedc, _eadf = _cec(_bgcfc, 1, 1, 0, 0)
			if _eadf != nil {
				return nil, nil, _a.Wrap(_eadf, _bgda, "\u0077\u006f\u0072\u0064 \u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073 \u002d \u0078\u0072\u0065\u0073\u0020\u003e\u00204\u0030\u0030")
			}
		}
		_eaae, _, _ebbc := _gefd(_dbedc)
		if _ebbc != nil {
			return nil, nil, _a.Wrap(_ebbc, _bgda, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
		_eada, _ebbc := _cgae(_eaae, _ccef)
		if _ebbc != nil {
			return nil, nil, _a.Wrap(_ebbc, _bgda, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
		_gbagc := &Bitmaps{}
		if _fff, _ebbc = _eada.ConnComponents(_gbagc, 4); _ebbc != nil {
			return nil, nil, _a.Wrap(_ebbc, _bgda, "\u0077\u006f\u0072\u0064\u0020\u0070r\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u002c\u0020\u0063\u006f\u006en\u0065\u0063\u0074\u0020\u0065\u0078\u0070a\u006e\u0064\u0065\u0064")
		}
		if _bdabe, _ebbc = _gbagc.ClipToBitmap(_bgcfc); _ebbc != nil {
			return nil, nil, _a.Wrap(_ebbc, _bgda, "\u0077o\u0072d\u0020\u0070\u0072\u0065\u0070\u0072\u006f\u0063\u0065\u0073\u0073")
		}
	}
	_bdabe, _eadf = _bdabe.SelectBySize(maxWidth, maxHeight, LocSelectIfBoth, SizeSelectIfLTE)
	if _eadf != nil {
		return nil, nil, _a.Wrap(_eadf, _bgda, "")
	}
	_fff, _eadf = _fff.SelectBySize(maxWidth, maxHeight, LocSelectIfBoth, SizeSelectIfLTE)
	if _eadf != nil {
		return nil, nil, _a.Wrap(_eadf, _bgda, "")
	}
	return _bdabe, _fff, nil
}
func CombineBytes(oldByte, newByte byte, op CombinationOperator) byte {
	return _bgbg(oldByte, newByte, op)
}

type Points []Point

func (_ebff *Bitmap) equivalent(_ddab *Bitmap) bool {
	if _ebff == _ddab {
		return true
	}
	if !_ebff.SizesEqual(_ddab) {
		return false
	}
	_bdg := _gabf(_ebff, _ddab, CmbOpXor)
	_afbe := _ebff.countPixels()
	_dbd := int(0.25 * float32(_afbe))
	if _bdg.thresholdPixelSum(_dbd) {
		return false
	}
	var (
		_fgbb [9][9]int
		_ceba [18][9]int
		_gbf  [9][18]int
		_feca int
		_dcfe int
	)
	_gdb := 9
	_cge := _ebff.Height / _gdb
	_fadf := _ebff.Width / _gdb
	_gecb, _fage := _cge/2, _fadf/2
	if _cge < _fadf {
		_gecb = _fadf / 2
		_fage = _cge / 2
	}
	_efd := float64(_gecb) * float64(_fage) * _ce.Pi
	_eedc := int(float64(_cge*_fadf/2) * 0.9)
	_aabg := int(float64(_fadf*_cge/2) * 0.9)
	for _gbc := 0; _gbc < _gdb; _gbc++ {
		_dage := _fadf*_gbc + _feca
		var _ffe int
		if _gbc == _gdb-1 {
			_feca = 0
			_ffe = _ebff.Width
		} else {
			_ffe = _dage + _fadf
			if ((_ebff.Width - _feca) % _gdb) > 0 {
				_feca++
				_ffe++
			}
		}
		for _decf := 0; _decf < _gdb; _decf++ {
			_gee := _cge*_decf + _dcfe
			var _dbcd int
			if _decf == _gdb-1 {
				_dcfe = 0
				_dbcd = _ebff.Height
			} else {
				_dbcd = _gee + _cge
				if (_ebff.Height-_dcfe)%_gdb > 0 {
					_dcfe++
					_dbcd++
				}
			}
			var _gfeb, _dabe, _cdbd, _aga int
			_cfb := (_dage + _ffe) / 2
			_gbcg := (_gee + _dbcd) / 2
			for _cbb := _dage; _cbb < _ffe; _cbb++ {
				for _dbce := _gee; _dbce < _dbcd; _dbce++ {
					if _bdg.GetPixel(_cbb, _dbce) {
						if _cbb < _cfb {
							_gfeb++
						} else {
							_dabe++
						}
						if _dbce < _gbcg {
							_aga++
						} else {
							_cdbd++
						}
					}
				}
			}
			_fgbb[_gbc][_decf] = _gfeb + _dabe
			_ceba[_gbc*2][_decf] = _gfeb
			_ceba[_gbc*2+1][_decf] = _dabe
			_gbf[_gbc][_decf*2] = _aga
			_gbf[_gbc][_decf*2+1] = _cdbd
		}
	}
	for _eccf := 0; _eccf < _gdb*2-1; _eccf++ {
		for _agdd := 0; _agdd < (_gdb - 1); _agdd++ {
			var _badg int
			for _adgd := 0; _adgd < 2; _adgd++ {
				for _cbgd := 0; _cbgd < 2; _cbgd++ {
					_badg += _ceba[_eccf+_adgd][_agdd+_cbgd]
				}
			}
			if _badg > _aabg {
				return false
			}
		}
	}
	for _ffa := 0; _ffa < (_gdb - 1); _ffa++ {
		for _bdbg := 0; _bdbg < ((_gdb * 2) - 1); _bdbg++ {
			var _fcac int
			for _eecc := 0; _eecc < 2; _eecc++ {
				for _gead := 0; _gead < 2; _gead++ {
					_fcac += _gbf[_ffa+_eecc][_bdbg+_gead]
				}
			}
			if _fcac > _eedc {
				return false
			}
		}
	}
	for _ege := 0; _ege < (_gdb - 2); _ege++ {
		for _gdce := 0; _gdce < (_gdb - 2); _gdce++ {
			var _bga, _dfag int
			for _cgba := 0; _cgba < 3; _cgba++ {
				for _gcd := 0; _gcd < 3; _gcd++ {
					if _cgba == _gcd {
						_bga += _fgbb[_ege+_cgba][_gdce+_gcd]
					}
					if (2 - _cgba) == _gcd {
						_dfag += _fgbb[_ege+_cgba][_gdce+_gcd]
					}
				}
			}
			if _bga > _aabg || _dfag > _aabg {
				return false
			}
		}
	}
	for _eac := 0; _eac < (_gdb - 1); _eac++ {
		for _aafe := 0; _aafe < (_gdb - 1); _aafe++ {
			var _eggd int
			for _egea := 0; _egea < 2; _egea++ {
				for _eae := 0; _eae < 2; _eae++ {
					_eggd += _fgbb[_eac+_egea][_aafe+_eae]
				}
			}
			if float64(_eggd) > _efd {
				return false
			}
		}
	}
	return true
}

var (
	_eefb *Bitmap
	_faed *Bitmap
)

func (_bacc *Bitmap) thresholdPixelSum(_feeb int) bool {
	var (
		_cgbbc int
		_cgf   uint8
		_dcb   byte
		_beba  int
	)
	_geb := _bacc.RowStride
	_cdggg := uint(_bacc.Width & 0x07)
	if _cdggg != 0 {
		_cgf = uint8((0xff << (8 - _cdggg)) & 0xff)
		_geb--
	}
	for _afaf := 0; _afaf < _bacc.Height; _afaf++ {
		for _beba = 0; _beba < _geb; _beba++ {
			_dcb = _bacc.Data[_afaf*_bacc.RowStride+_beba]
			_cgbbc += int(_afdc[_dcb])
		}
		if _cdggg != 0 {
			_dcb = _bacc.Data[_afaf*_bacc.RowStride+_beba] & _cgf
			_cgbbc += int(_afdc[_dcb])
		}
		if _cgbbc > _feeb {
			return true
		}
	}
	return false
}
func New(width, height int) *Bitmap {
	_efa := _egda(width, height)
	_efa.Data = make([]byte, height*_efa.RowStride)
	return _efa
}
func (_bcd *Bitmap) GetVanillaData() []byte {
	if _bcd.Color == Chocolate {
		_bcd.inverseData()
	}
	return _bcd.Data
}

var (
	_fcgba = []byte{0x00, 0x80, 0xC0, 0xE0, 0xF0, 0xF8, 0xFC, 0xFE, 0xFF}
	_gded  = []byte{0x00, 0x01, 0x03, 0x07, 0x0F, 0x1F, 0x3F, 0x7F, 0xFF}
)

func _gaecg(_aaag ...MorphProcess) (_acbd error) {
	const _accg = "v\u0065r\u0069\u0066\u0079\u004d\u006f\u0072\u0070\u0068P\u0072\u006f\u0063\u0065ss\u0065\u0073"
	var _dedee, _cgea int
	for _cbcb, _agdg := range _aaag {
		if _acbd = _agdg.verify(_cbcb, &_dedee, &_cgea); _acbd != nil {
			return _a.Wrap(_acbd, _accg, "")
		}
	}
	if _cgea != 0 && _dedee != 0 {
		return _a.Error(_accg, "\u004d\u006f\u0072\u0070\u0068\u0020\u0073\u0065\u0071\u0075\u0065n\u0063\u0065\u0020\u002d\u0020\u0062\u006f\u0072d\u0065r\u0020\u0061\u0064\u0064\u0065\u0064\u0020\u0062\u0075\u0074\u0020\u006e\u0065\u0074\u0020\u0072\u0065\u0064u\u0063\u0074\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020\u0030")
	}
	return nil
}

const (
	CmbOpOr CombinationOperator = iota
	CmbOpAnd
	CmbOpXor
	CmbOpXNor
	CmbOpReplace
	CmbOpNot
)

func (_cfebc *Bitmaps) GetBitmap(i int) (*Bitmap, error) {
	const _ggcb = "\u0047e\u0074\u0042\u0069\u0074\u006d\u0061p"
	if _cfebc == nil {
		return nil, _a.Error(_ggcb, "p\u0072o\u0076\u0069\u0064\u0065\u0064\u0020\u006e\u0069l\u0020\u0042\u0069\u0074ma\u0070\u0073")
	}
	if i > len(_cfebc.Values)-1 {
		return nil, _a.Errorf(_ggcb, "\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065", i)
	}
	return _cfebc.Values[i], nil
}
func TstWordBitmap(t *_f.T, scale ...int) *Bitmap {
	_dcbe := 1
	if len(scale) > 0 {
		_dcbe = scale[0]
	}
	_bbdd := 3
	_bggd := 9 + 7 + 15 + 2*_bbdd
	_egbe := 5 + _bbdd + 5
	_ecbbd := New(_bggd*_dcbe, _egbe*_dcbe)
	_fbgfa := &Bitmaps{}
	var _accbe *int
	_bbdd *= _dcbe
	_fccg := 0
	_accbe = &_fccg
	_dcaa := 0
	_fffc := TstDSymbol(t, scale...)
	TstAddSymbol(t, _fbgfa, _fffc, _accbe, _dcaa, 1*_dcbe)
	_fffc = TstOSymbol(t, scale...)
	TstAddSymbol(t, _fbgfa, _fffc, _accbe, _dcaa, _bbdd)
	_fffc = TstISymbol(t, scale...)
	TstAddSymbol(t, _fbgfa, _fffc, _accbe, _dcaa, 1*_dcbe)
	_fffc = TstTSymbol(t, scale...)
	TstAddSymbol(t, _fbgfa, _fffc, _accbe, _dcaa, _bbdd)
	_fffc = TstNSymbol(t, scale...)
	TstAddSymbol(t, _fbgfa, _fffc, _accbe, _dcaa, 1*_dcbe)
	_fffc = TstOSymbol(t, scale...)
	TstAddSymbol(t, _fbgfa, _fffc, _accbe, _dcaa, 1*_dcbe)
	_fffc = TstWSymbol(t, scale...)
	TstAddSymbol(t, _fbgfa, _fffc, _accbe, _dcaa, 0)
	*_accbe = 0
	_dcaa = 5*_dcbe + _bbdd
	_fffc = TstOSymbol(t, scale...)
	TstAddSymbol(t, _fbgfa, _fffc, _accbe, _dcaa, 1*_dcbe)
	_fffc = TstRSymbol(t, scale...)
	TstAddSymbol(t, _fbgfa, _fffc, _accbe, _dcaa, _bbdd)
	_fffc = TstNSymbol(t, scale...)
	TstAddSymbol(t, _fbgfa, _fffc, _accbe, _dcaa, 1*_dcbe)
	_fffc = TstESymbol(t, scale...)
	TstAddSymbol(t, _fbgfa, _fffc, _accbe, _dcaa, 1*_dcbe)
	_fffc = TstVSymbol(t, scale...)
	TstAddSymbol(t, _fbgfa, _fffc, _accbe, _dcaa, 1*_dcbe)
	_fffc = TstESymbol(t, scale...)
	TstAddSymbol(t, _fbgfa, _fffc, _accbe, _dcaa, 1*_dcbe)
	_fffc = TstRSymbol(t, scale...)
	TstAddSymbol(t, _fbgfa, _fffc, _accbe, _dcaa, 0)
	TstWriteSymbols(t, _fbgfa, _ecbbd)
	return _ecbbd
}
func TstTSymbol(t *_f.T, scale ...int) *Bitmap {
	_eaag, _gdfe := NewWithData(5, 5, []byte{0xF8, 0x20, 0x20, 0x20, 0x20})
	_g.NoError(t, _gdfe)
	return TstGetScaledSymbol(t, _eaag, scale...)
}

type Color int

func init() {
	const _eegd = "\u0062\u0069\u0074\u006dap\u0073\u002e\u0069\u006e\u0069\u0074\u0069\u0061\u006c\u0069\u007a\u0061\u0074\u0069o\u006e"
	_eefb = New(50, 40)
	var _dbef error
	_eefb, _dbef = _eefb.AddBorder(2, 1)
	if _dbef != nil {
		panic(_a.Wrap(_dbef, _eegd, "f\u0072\u0061\u006d\u0065\u0042\u0069\u0074\u006d\u0061\u0070"))
	}
	_faed, _dbef = NewWithData(50, 22, _gcda)
	if _dbef != nil {
		panic(_a.Wrap(_dbef, _eegd, "i\u006d\u0061\u0067\u0065\u0042\u0069\u0074\u006d\u0061\u0070"))
	}
}
func _adea(_acaf *Bitmap, _aafgc, _cbbc int, _abfe, _geac int, _fdac RasterOperator, _gedd *Bitmap, _dgbff, _cadd int) error {
	var _fbbc, _gdgdd, _cbdd, _beff int
	if _aafgc < 0 {
		_dgbff -= _aafgc
		_abfe += _aafgc
		_aafgc = 0
	}
	if _dgbff < 0 {
		_aafgc -= _dgbff
		_abfe += _dgbff
		_dgbff = 0
	}
	_fbbc = _aafgc + _abfe - _acaf.Width
	if _fbbc > 0 {
		_abfe -= _fbbc
	}
	_gdgdd = _dgbff + _abfe - _gedd.Width
	if _gdgdd > 0 {
		_abfe -= _gdgdd
	}
	if _cbbc < 0 {
		_cadd -= _cbbc
		_geac += _cbbc
		_cbbc = 0
	}
	if _cadd < 0 {
		_cbbc -= _cadd
		_geac += _cadd
		_cadd = 0
	}
	_cbdd = _cbbc + _geac - _acaf.Height
	if _cbdd > 0 {
		_geac -= _cbdd
	}
	_beff = _cadd + _geac - _gedd.Height
	if _beff > 0 {
		_geac -= _beff
	}
	if _abfe <= 0 || _geac <= 0 {
		return nil
	}
	var _fegc error
	switch {
	case _aafgc&7 == 0 && _dgbff&7 == 0:
		_fegc = _cfge(_acaf, _aafgc, _cbbc, _abfe, _geac, _fdac, _gedd, _dgbff, _cadd)
	case _aafgc&7 == _dgbff&7:
		_fegc = _eggb(_acaf, _aafgc, _cbbc, _abfe, _geac, _fdac, _gedd, _dgbff, _cadd)
	default:
		_fegc = _eecdf(_acaf, _aafgc, _cbbc, _abfe, _geac, _fdac, _gedd, _dgbff, _cadd)
	}
	if _fegc != nil {
		return _a.Wrap(_fegc, "r\u0061\u0073\u0074\u0065\u0072\u004f\u0070\u004c\u006f\u0077", "")
	}
	return nil
}
func (_eaad *byWidth) Len() int { return len(_eaad.Values) }
func _caed() []int {
	_gdcb := make([]int, 256)
	_gdcb[0] = 0
	_gdcb[1] = 7
	var _gffb int
	for _gffb = 2; _gffb < 4; _gffb++ {
		_gdcb[_gffb] = _gdcb[_gffb-2] + 6
	}
	for _gffb = 4; _gffb < 8; _gffb++ {
		_gdcb[_gffb] = _gdcb[_gffb-4] + 5
	}
	for _gffb = 8; _gffb < 16; _gffb++ {
		_gdcb[_gffb] = _gdcb[_gffb-8] + 4
	}
	for _gffb = 16; _gffb < 32; _gffb++ {
		_gdcb[_gffb] = _gdcb[_gffb-16] + 3
	}
	for _gffb = 32; _gffb < 64; _gffb++ {
		_gdcb[_gffb] = _gdcb[_gffb-32] + 2
	}
	for _gffb = 64; _gffb < 128; _gffb++ {
		_gdcb[_gffb] = _gdcb[_gffb-64] + 1
	}
	for _gffb = 128; _gffb < 256; _gffb++ {
		_gdcb[_gffb] = _gdcb[_gffb-128]
	}
	return _gdcb
}
func (_dddd *Bitmaps) AddBitmap(bm *Bitmap) { _dddd.Values = append(_dddd.Values, bm) }

type byHeight Bitmaps

func (_ebedf *ClassedPoints) SortByY()               { _ebedf._fbdba = _ebedf.ySortFunction(); _e.Sort(_ebedf) }
func (_debc *Selection) setOrigin(_bgdaf, _dega int) { _debc.Cy, _debc.Cx = _bgdaf, _dega }
func Blit(src *Bitmap, dst *Bitmap, x, y int, op CombinationOperator) error {
	var _adgg, _efgdg int
	_agc := src.RowStride - 1
	if x < 0 {
		_efgdg = -x
		x = 0
	} else if x+src.Width > dst.Width {
		_agc -= src.Width + x - dst.Width
	}
	if y < 0 {
		_adgg = -y
		y = 0
		_efgdg += src.RowStride
		_agc += src.RowStride
	} else if y+src.Height > dst.Height {
		_adgg = src.Height + y - dst.Height
	}
	var (
		_gdbb int
		_gefg error
	)
	_cffa := x & 0x07
	_bgfa := 8 - _cffa
	_aggg := src.Width & 0x07
	_bdd := _bgfa - _aggg
	_cab := _bgfa&0x07 != 0
	_aac := src.Width <= ((_agc-_efgdg)<<3)+_bgfa
	_deeb := dst.GetByteIndex(x, y)
	_aag := _adgg + dst.Height
	if src.Height > _aag {
		_gdbb = _aag
	} else {
		_gdbb = src.Height
	}
	switch {
	case !_cab:
		_gefg = _bdde(src, dst, _adgg, _gdbb, _deeb, _efgdg, _agc, op)
	case _aac:
		_gefg = _agce(src, dst, _adgg, _gdbb, _deeb, _efgdg, _agc, _bdd, _cffa, _bgfa, op)
	default:
		_gefg = _gcga(src, dst, _adgg, _gdbb, _deeb, _efgdg, _agc, _bdd, _cffa, _bgfa, op, _aggg)
	}
	return _gefg
}

type Point struct{ X, Y float32 }

func CorrelationScoreSimple(bm1, bm2 *Bitmap, area1, area2 int, delX, delY float32, maxDiffW, maxDiffH int, tab []int) (_daag float64, _efacf error) {
	const _fgba = "\u0043\u006f\u0072\u0072el\u0061\u0074\u0069\u006f\u006e\u0053\u0063\u006f\u0072\u0065\u0053\u0069\u006d\u0070l\u0065"
	if bm1 == nil || bm2 == nil {
		return _daag, _a.Error(_fgba, "n\u0069l\u0020\u0062\u0069\u0074\u006d\u0061\u0070\u0073 \u0070\u0072\u006f\u0076id\u0065\u0064")
	}
	if tab == nil {
		return _daag, _a.Error(_fgba, "\u0074\u0061\u0062\u0020\u0075\u006e\u0064\u0065\u0066\u0069\u006e\u0065\u0064")
	}
	if area1 == 0 || area2 == 0 {
		return _daag, _a.Error(_fgba, "\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0061\u0072e\u0061\u0073\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u003e\u0020\u0030")
	}
	_effd, _bgbab := bm1.Width, bm1.Height
	_cfbbce, _bfddf := bm2.Width, bm2.Height
	if _bgeb(_effd-_cfbbce) > maxDiffW {
		return 0, nil
	}
	if _bgeb(_bgbab-_bfddf) > maxDiffH {
		return 0, nil
	}
	var _egfce, _ffbf int
	if delX >= 0 {
		_egfce = int(delX + 0.5)
	} else {
		_egfce = int(delX - 0.5)
	}
	if delY >= 0 {
		_ffbf = int(delY + 0.5)
	} else {
		_ffbf = int(delY - 0.5)
	}
	_gbae := bm1.createTemplate()
	if _efacf = _gbae.RasterOperation(_egfce, _ffbf, _cfbbce, _bfddf, PixSrc, bm2, 0, 0); _efacf != nil {
		return _daag, _a.Wrap(_efacf, _fgba, "\u0062m\u0032 \u0074\u006f\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065")
	}
	if _efacf = _gbae.RasterOperation(0, 0, _effd, _bgbab, PixSrcAndDst, bm1, 0, 0); _efacf != nil {
		return _daag, _a.Wrap(_efacf, _fgba, "b\u006d\u0031\u0020\u0061\u006e\u0064\u0020\u0062\u006d\u0054")
	}
	_abge := _gbae.countPixels()
	_daag = float64(_abge) * float64(_abge) / (float64(area1) * float64(area2))
	return _daag, nil
}
func (_cfbc *byHeight) Swap(i, j int) {
	_cfbc.Values[i], _cfbc.Values[j] = _cfbc.Values[j], _cfbc.Values[i]
	if _cfbc.Boxes != nil {
		_cfbc.Boxes[i], _cfbc.Boxes[j] = _cfbc.Boxes[j], _cfbc.Boxes[i]
	}
}
func (_ddae *Bitmap) GetByte(index int) (byte, error) {
	if index > len(_ddae.Data)-1 || index < 0 {
		return 0, _a.Errorf("\u0047e\u0074\u0042\u0079\u0074\u0065", "\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065", index)
	}
	return _ddae.Data[index], nil
}
func (_gdad *ClassedPoints) YAtIndex(i int) float32 { return (*_gdad.Points)[_gdad.IntSlice[i]].Y }
func _aegaf(_cccb, _fffe *Bitmap, _accfg, _fgfc int) (*Bitmap, error) {
	const _bfea = "\u0063\u006c\u006f\u0073\u0065\u0042\u0072\u0069\u0063\u006b"
	if _fffe == nil {
		return nil, _a.Error(_bfea, "\u0073o\u0075r\u0063\u0065\u0020\u006e\u006ft\u0020\u0064e\u0066\u0069\u006e\u0065\u0064")
	}
	if _accfg < 1 || _fgfc < 1 {
		return nil, _a.Error(_bfea, "\u0068S\u0069\u007a\u0065\u0020\u0061\u006e\u0064\u0020\u0076\u0053\u0069z\u0065\u0020\u006e\u006f\u0074\u0020\u003e\u003d\u0020\u0031")
	}
	if _accfg == 1 && _fgfc == 1 {
		return _fffe.Copy(), nil
	}
	if _accfg == 1 || _fgfc == 1 {
		_fbgdb := SelCreateBrick(_fgfc, _accfg, _fgfc/2, _accfg/2, SelHit)
		var _efega error
		_cccb, _efega = _ecaga(_cccb, _fffe, _fbgdb)
		if _efega != nil {
			return nil, _a.Wrap(_efega, _bfea, "\u0068S\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031\u0020\u007c\u007c \u0076\u0053\u0069\u007a\u0065\u0020\u003d\u003d\u0020\u0031")
		}
		return _cccb, nil
	}
	_bbff := SelCreateBrick(1, _accfg, 0, _accfg/2, SelHit)
	_cdc := SelCreateBrick(_fgfc, 1, _fgfc/2, 0, SelHit)
	_gca, _beaab := _egcg(nil, _fffe, _bbff)
	if _beaab != nil {
		return nil, _a.Wrap(_beaab, _bfea, "\u0031\u0073\u0074\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	if _cccb, _beaab = _egcg(_cccb, _gca, _cdc); _beaab != nil {
		return nil, _a.Wrap(_beaab, _bfea, "\u0032\u006e\u0064\u0020\u0064\u0069\u006c\u0061\u0074\u0065")
	}
	if _, _beaab = _abbc(_gca, _cccb, _bbff); _beaab != nil {
		return nil, _a.Wrap(_beaab, _bfea, "\u0031s\u0074\u0020\u0065\u0072\u006f\u0064e")
	}
	if _, _beaab = _abbc(_cccb, _gca, _cdc); _beaab != nil {
		return nil, _a.Wrap(_beaab, _bfea, "\u0032n\u0064\u0020\u0065\u0072\u006f\u0064e")
	}
	return _cccb, nil
}
func _egda(_gbag, _cfd int) *Bitmap {
	return &Bitmap{Width: _gbag, Height: _cfd, RowStride: (_gbag + 7) >> 3}
}
func TstImageBitmapData() []byte { return _faed.Data }
func Extract(roi _ac.Rectangle, src *Bitmap) (*Bitmap, error) {
	_afg := New(roi.Dx(), roi.Dy())
	_gafg := roi.Min.X & 0x07
	_acee := 8 - _gafg
	_dbbc := uint(8 - _afg.Width&0x07)
	_abgf := src.GetByteIndex(roi.Min.X, roi.Min.Y)
	_bfcg := src.GetByteIndex(roi.Max.X-1, roi.Min.Y)
	_aabe := _afg.RowStride == _bfcg+1-_abgf
	var _gcfe int
	for _fgeb := roi.Min.Y; _fgeb < roi.Max.Y; _fgeb++ {
		_eddg := _abgf
		_dgfg := _gcfe
		switch {
		case _abgf == _bfcg:
			_edef, _ggeec := src.GetByte(_eddg)
			if _ggeec != nil {
				return nil, _ggeec
			}
			_edef <<= uint(_gafg)
			_ggeec = _afg.SetByte(_dgfg, _fge(_dbbc, _edef))
			if _ggeec != nil {
				return nil, _ggeec
			}
		case _gafg == 0:
			for _dbed := _abgf; _dbed <= _bfcg; _dbed++ {
				_dfab, _faad := src.GetByte(_eddg)
				if _faad != nil {
					return nil, _faad
				}
				_eddg++
				if _dbed == _bfcg && _aabe {
					_dfab = _fge(_dbbc, _dfab)
				}
				_faad = _afg.SetByte(_dgfg, _dfab)
				if _faad != nil {
					return nil, _faad
				}
				_dgfg++
			}
		default:
			_eagff := _fggb(src, _afg, uint(_gafg), uint(_acee), _dbbc, _abgf, _bfcg, _aabe, _eddg, _dgfg)
			if _eagff != nil {
				return nil, _eagff
			}
		}
		_abgf += src.RowStride
		_bfcg += src.RowStride
		_gcfe += _afg.RowStride
	}
	return _afg, nil
}

type SizeSelection int
type Bitmap struct {
	Width, Height            int
	BitmapNumber             int
	RowStride                int
	Data                     []byte
	Color                    Color
	Special                  int
	Text                     string
	XResolution, YResolution int
}

func (_fggf *Bitmap) SetDefaultPixel() {
	for _dbgg := range _fggf.Data {
		_fggf.Data[_dbgg] = byte(0xff)
	}
}
func (_ceae *Bitmaps) CountPixels() *_fc.NumSlice {
	_fdgd := &_fc.NumSlice{}
	for _, _agbg := range _ceae.Values {
		_fdgd.AddInt(_agbg.CountPixels())
	}
	return _fdgd
}
func _gdf(_ageg, _cfgc, _ebda *Bitmap, _geee int) (*Bitmap, error) {
	const _bfef = "\u0073\u0065\u0065\u0064\u0046\u0069\u006c\u006c\u0042i\u006e\u0061\u0072\u0079"
	if _cfgc == nil {
		return nil, _a.Error(_bfef, "s\u006fu\u0072\u0063\u0065\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _ebda == nil {
		return nil, _a.Error(_bfef, "'\u006da\u0073\u006b\u0027\u0020\u0062\u0069\u0074\u006da\u0070\u0020\u0069\u0073 n\u0069\u006c")
	}
	if _geee != 4 && _geee != 8 {
		return nil, _a.Error(_bfef, "\u0063\u006f\u006en\u0065\u0063\u0074\u0069v\u0069\u0074\u0079\u0020\u006e\u006f\u0074 \u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0034\u002c\u0038\u007d")
	}
	var _egbaf error
	_ageg, _egbaf = _bggbg(_ageg, _cfgc)
	if _egbaf != nil {
		return nil, _a.Wrap(_egbaf, _bfef, "\u0063o\u0070y\u0020\u0073\u006f\u0075\u0072c\u0065\u0020t\u006f\u0020\u0027\u0064\u0027")
	}
	_bgbgg := _cfgc.createTemplate()
	_ebda.setPadBits(0)
	for _bgge := 0; _bgge < _bbgab; _bgge++ {
		_bgbgg, _egbaf = _bggbg(_bgbgg, _ageg)
		if _egbaf != nil {
			return nil, _a.Wrapf(_egbaf, _bfef, "\u0069\u0074\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064", _bgge)
		}
		if _egbaf = _dcba(_ageg, _ebda, _geee); _egbaf != nil {
			return nil, _a.Wrapf(_egbaf, _bfef, "\u0069\u0074\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0064", _bgge)
		}
		if _bgbgg.Equals(_ageg) {
			break
		}
	}
	return _ageg, nil
}
