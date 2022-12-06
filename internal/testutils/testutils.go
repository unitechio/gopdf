package testutils

import (
	_e "crypto/md5"
	_d "encoding/hex"
	_gfe "errors"
	_gc "fmt"
	_fa "image"
	_gf "image/png"
	_ge "io"
	_a "os"
	_ea "os/exec"
	_be "path/filepath"
	_f "strings"
	_g "testing"

	_gcf "bitbucket.org/shenghui0779/gopdf/common"
	_ef "bitbucket.org/shenghui0779/gopdf/core"
)

func ComparePNGFiles(file1, file2 string) (bool, error) {
	_ead, _bdb := HashFile(file1)
	if _bdb != nil {
		return false, _bdb
	}
	_ba, _bdb := HashFile(file2)
	if _bdb != nil {
		return false, _bdb
	}
	if _ead == _ba {
		return true, nil
	}
	_gb, _bdb := ReadPNG(file1)
	if _bdb != nil {
		return false, _bdb
	}
	_cb, _bdb := ReadPNG(file2)
	if _bdb != nil {
		return false, _bdb
	}
	if _gb.Bounds() != _cb.Bounds() {
		return false, nil
	}
	return CompareImages(_gb, _cb)
}
func ReadPNG(file string) (_fa.Image, error) {
	_bd, _fd := _a.Open(file)
	if _fd != nil {
		return nil, _fd
	}
	defer _bd.Close()
	return _gf.Decode(_bd)
}
func HashFile(file string) (string, error) {
	_ga, _af := _a.Open(file)
	if _af != nil {
		return "", _af
	}
	defer _ga.Close()
	_ed := _e.New()
	if _, _af = _ge.Copy(_ed, _ga); _af != nil {
		return "", _af
	}
	return _d.EncodeToString(_ed.Sum(nil)), nil
}
func RunRenderTest(t *_g.T, pdfPath, outputDir, baselineRenderPath string, saveBaseline bool) {
	_db := _f.TrimSuffix(_be.Base(pdfPath), _be.Ext(pdfPath))
	t.Run("\u0072\u0065\u006e\u0064\u0065\u0072", func(_agc *_g.T) {
		_de := _be.Join(outputDir, _db)
		_ggc := _de + "\u002d%\u0064\u002e\u0070\u006e\u0067"
		if _fde := RenderPDFToPNGs(pdfPath, 0, _ggc); _fde != nil {
			_agc.Skip(_fde)
		}
		for _ce := 1; true; _ce++ {
			_bc := _gc.Sprintf("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg", _de, _ce)
			_geg := _be.Join(baselineRenderPath, _gc.Sprintf("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067", _db, _ce))
			if _, _caf := _a.Stat(_bc); _caf != nil {
				break
			}
			_agc.Logf("\u0025\u0073", _geg)
			if saveBaseline {
				_agc.Logf("\u0043\u006fp\u0079\u0069\u006eg\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073", _bc, _geg)
				_bdf := CopyFile(_bc, _geg)
				if _bdf != nil {
					_agc.Fatalf("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _geg, _bdf)
				}
				continue
			}
			_agc.Run(_gc.Sprintf("\u0070\u0061\u0067\u0065\u0025\u0064", _ce), func(_bae *_g.T) {
				_bae.Logf("\u0043o\u006dp\u0061\u0072\u0069\u006e\u0067 \u0025\u0073 \u0076\u0073\u0020\u0025\u0073", _bc, _geg)
				_bfg, _fb := ComparePNGFiles(_bc, _geg)
				if _a.IsNotExist(_fb) {
					_bae.Fatal("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
				} else if !_bfg {
					_bae.Fatal("\u0077\u0072\u006f\u006eg \u0070\u0061\u0067\u0065\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0065\u0064")
				}
			})
		}
	})
}
func RenderPDFToPNGs(pdfPath string, dpi int, outpathTpl string) error {
	if dpi <= 0 {
		dpi = 100
	}
	if _, _bb := _ea.LookPath("\u0067\u0073"); _bb != nil {
		return ErrRenderNotSupported
	}
	return _ea.Command("\u0067\u0073", "\u002d\u0073\u0044\u0045\u0056\u0049\u0043\u0045\u003d\u0070\u006e\u0067a\u006c\u0070\u0068\u0061", "\u002d\u006f", outpathTpl, _gc.Sprintf("\u002d\u0072\u0025\u0064", dpi), pdfPath).Run()
}
func CompareImages(img1, img2 _fa.Image) (bool, error) {
	_bf := img1.Bounds()
	_edf := 0
	for _gfef := 0; _gfef < _bf.Size().X; _gfef++ {
		for _ec := 0; _ec < _bf.Size().Y; _ec++ {
			_bg, _gea, _ag, _ := img1.At(_gfef, _ec).RGBA()
			_aa, _ad, _ae, _ := img2.At(_gfef, _ec).RGBA()
			if _bg != _aa || _gea != _ad || _ag != _ae {
				_edf++
			}
		}
	}
	_efe := float64(_edf) / float64(_bf.Dx()*_bf.Dy())
	if _efe > 0.0001 {
		_gc.Printf("\u0064\u0069\u0066f \u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0064\u0029\u000a", _efe, _edf)
		return false, nil
	}
	return true, nil
}
func CopyFile(src, dst string) error {
	_c, _dg := _a.Open(src)
	if _dg != nil {
		return _dg
	}
	defer _c.Close()
	_ca, _dg := _a.Create(dst)
	if _dg != nil {
		return _dg
	}
	defer _ca.Close()
	_, _dg = _ge.Copy(_ca, _c)
	return _dg
}

var (
	ErrRenderNotSupported = _gfe.New("\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0073 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u006e\u0020\u0074\u0068\u0069\u0073\u0020\u0073\u0079\u0073\u0074\u0065m")
)

func CompareDictionariesDeep(d1, d2 *_ef.PdfObjectDictionary) bool {
	if len(d1.Keys()) != len(d2.Keys()) {
		_gcf.Log.Debug("\u0044\u0069\u0063\u0074\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029", len(d1.Keys()), len(d2.Keys()))
		_gcf.Log.Debug("\u0057\u0061s\u0020\u0027\u0025s\u0027\u0020\u0076\u0073\u0020\u0027\u0025\u0073\u0027", d1.WriteString(), d2.WriteString())
		return false
	}
	for _, _dfdb := range d1.Keys() {
		if _dfdb == "\u0050\u0061\u0072\u0065\u006e\u0074" {
			continue
		}
		_eg := _ef.TraceToDirectObject(d1.Get(_dfdb))
		_eff := _ef.TraceToDirectObject(d2.Get(_dfdb))
		if _eg == nil {
			_gcf.Log.Debug("\u00761\u0020\u0069\u0073\u0020\u006e\u0069l")
			return false
		}
		if _eff == nil {
			_gcf.Log.Debug("\u00762\u0020\u0069\u0073\u0020\u006e\u0069l")
			return false
		}
		switch _fac := _eg.(type) {
		case *_ef.PdfObjectDictionary:
			_ff, _deg := _eff.(*_ef.PdfObjectDictionary)
			if !_deg {
				_gcf.Log.Debug("\u0054\u0079\u0070\u0065 m\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020\u0025\u0054\u0020\u0076\u0073\u0020%\u0054", _eg, _eff)
				return false
			}
			if !CompareDictionariesDeep(_fac, _ff) {
				return false
			}
			continue
		case *_ef.PdfObjectArray:
			_ee, _bgd := _eff.(*_ef.PdfObjectArray)
			if !_bgd {
				_gcf.Log.Debug("\u00762\u0020n\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
				return false
			}
			if _fac.Len() != _ee.Len() {
				_gcf.Log.Debug("\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029", _fac.Len(), _ee.Len())
				return false
			}
			for _bea := 0; _bea < _fac.Len(); _bea++ {
				_fga := _ef.TraceToDirectObject(_fac.Get(_bea))
				_ggg := _ef.TraceToDirectObject(_ee.Get(_bea))
				if _cafa, _dgd := _fga.(*_ef.PdfObjectDictionary); _dgd {
					_dde, _ffb := _ggg.(*_ef.PdfObjectDictionary)
					if !_ffb {
						return false
					}
					if !CompareDictionariesDeep(_cafa, _dde) {
						return false
					}
				} else {
					if _fga.WriteString() != _ggg.WriteString() {
						_gcf.Log.Debug("M\u0069\u0073\u006d\u0061tc\u0068 \u0027\u0025\u0073\u0027\u0020!\u003d\u0020\u0027\u0025\u0073\u0027", _fga.WriteString(), _ggg.WriteString())
						return false
					}
				}
			}
			continue
		}
		if _eg.String() != _eff.String() {
			_gcf.Log.Debug("\u006b\u0065y\u003d\u0025\u0073\u0020\u004d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0021\u0020\u0027\u0025\u0073\u0027\u0020\u0021\u003d\u0020'%\u0073\u0027", _dfdb, _eg.String(), _eff.String())
			_gcf.Log.Debug("\u0046o\u0072 \u0027\u0025\u0054\u0027\u0020\u002d\u0020\u0027\u0025\u0054\u0027", _eg, _eff)
			_gcf.Log.Debug("\u0046\u006f\u0072\u0020\u0027\u0025\u002b\u0076\u0027\u0020\u002d\u0020'\u0025\u002b\u0076\u0027", _eg, _eff)
			return false
		}
	}
	return true
}
func ParseIndirectObjects(rawpdf string) (map[int64]_ef.PdfObject, error) {
	_eac := _ef.NewParserFromString(rawpdf)
	_bdd := map[int64]_ef.PdfObject{}
	for {
		_fe, _eadb := _eac.ParseIndirectObject()
		if _eadb != nil {
			if _eadb == _ge.EOF {
				break
			}
			return nil, _eadb
		}
		switch _df := _fe.(type) {
		case *_ef.PdfIndirectObject:
			_bdd[_df.ObjectNumber] = _fe
		case *_ef.PdfObjectStream:
			_bdd[_df.ObjectNumber] = _fe
		}
	}
	for _, _gbf := range _bdd {
		_eadd(_gbf, _bdd)
	}
	return _bdd, nil
}
func _eadd(_bge _ef.PdfObject, _fdf map[int64]_ef.PdfObject) error {
	switch _cab := _bge.(type) {
	case *_ef.PdfIndirectObject:
		_dd := _cab
		_eadd(_dd.PdfObject, _fdf)
	case *_ef.PdfObjectDictionary:
		_cc := _cab
		for _, _dba := range _cc.Keys() {
			_cea := _cc.Get(_dba)
			if _aad, _aga := _cea.(*_ef.PdfObjectReference); _aga {
				_cf, _fc := _fdf[_aad.ObjectNumber]
				if !_fc {
					return _gfe.New("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074")
				}
				_cc.Set(_dba, _cf)
			} else {
				_eadd(_cea, _fdf)
			}
		}
	case *_ef.PdfObjectArray:
		_edg := _cab
		for _cce, _fbe := range _edg.Elements() {
			if _gae, _gbe := _fbe.(*_ef.PdfObjectReference); _gbe {
				_gbea, _dfd := _fdf[_gae.ObjectNumber]
				if !_dfd {
					return _gfe.New("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074")
				}
				_edg.Set(_cce, _gbea)
			} else {
				_eadd(_fbe, _fdf)
			}
		}
	}
	return nil
}
