package testutils

import (
	_a "crypto/md5"
	_fbf "encoding/hex"
	_fe "errors"
	_ce "fmt"
	_ge "image"
	_g "image/png"
	_fb "io"
	_c "os"
	_e "os/exec"
	_b "path/filepath"
	_df "strings"
	_f "testing"

	_ba "bitbucket.org/shenghui0779/gopdf/common"
	_fed "bitbucket.org/shenghui0779/gopdf/core"
)

func CompareDictionariesDeep(d1, d2 *_fed.PdfObjectDictionary) bool {
	if len(d1.Keys()) != len(d2.Keys()) {
		_ba.Log.Debug("\u0044\u0069\u0063\u0074\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029", len(d1.Keys()), len(d2.Keys()))
		_ba.Log.Debug("\u0057\u0061s\u0020\u0027\u0025s\u0027\u0020\u0076\u0073\u0020\u0027\u0025\u0073\u0027", d1.WriteString(), d2.WriteString())
		return false
	}
	for _, _ged := range d1.Keys() {
		if _ged == "\u0050\u0061\u0072\u0065\u006e\u0074" {
			continue
		}
		_bba := _fed.TraceToDirectObject(d1.Get(_ged))
		_fff := _fed.TraceToDirectObject(d2.Get(_ged))
		if _bba == nil {
			_ba.Log.Debug("\u00761\u0020\u0069\u0073\u0020\u006e\u0069l")
			return false
		}
		if _fff == nil {
			_ba.Log.Debug("\u00762\u0020\u0069\u0073\u0020\u006e\u0069l")
			return false
		}
		switch _adg := _bba.(type) {
		case *_fed.PdfObjectDictionary:
			_gfa, _bbac := _fff.(*_fed.PdfObjectDictionary)
			if !_bbac {
				_ba.Log.Debug("\u0054\u0079\u0070\u0065 m\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020\u0025\u0054\u0020\u0076\u0073\u0020%\u0054", _bba, _fff)
				return false
			}
			if !CompareDictionariesDeep(_adg, _gfa) {
				return false
			}
			continue
		case *_fed.PdfObjectArray:
			_ecf, _baec := _fff.(*_fed.PdfObjectArray)
			if !_baec {
				_ba.Log.Debug("\u00762\u0020n\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079")
				return false
			}
			if _adg.Len() != _ecf.Len() {
				_ba.Log.Debug("\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029", _adg.Len(), _ecf.Len())
				return false
			}
			for _bbf := 0; _bbf < _adg.Len(); _bbf++ {
				_abg := _fed.TraceToDirectObject(_adg.Get(_bbf))
				_dd := _fed.TraceToDirectObject(_ecf.Get(_bbf))
				if _egdd, _ebc := _abg.(*_fed.PdfObjectDictionary); _ebc {
					_dbg, _fbe := _dd.(*_fed.PdfObjectDictionary)
					if !_fbe {
						return false
					}
					if !CompareDictionariesDeep(_egdd, _dbg) {
						return false
					}
				} else {
					if _abg.WriteString() != _dd.WriteString() {
						_ba.Log.Debug("M\u0069\u0073\u006d\u0061tc\u0068 \u0027\u0025\u0073\u0027\u0020!\u003d\u0020\u0027\u0025\u0073\u0027", _abg.WriteString(), _dd.WriteString())
						return false
					}
				}
			}
			continue
		}
		if _bba.String() != _fff.String() {
			_ba.Log.Debug("\u006b\u0065y\u003d\u0025\u0073\u0020\u004d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0021\u0020\u0027\u0025\u0073\u0027\u0020\u0021\u003d\u0020'%\u0073\u0027", _ged, _bba.String(), _fff.String())
			_ba.Log.Debug("\u0046o\u0072 \u0027\u0025\u0054\u0027\u0020\u002d\u0020\u0027\u0025\u0054\u0027", _bba, _fff)
			_ba.Log.Debug("\u0046\u006f\u0072\u0020\u0027\u0025\u002b\u0076\u0027\u0020\u002d\u0020'\u0025\u002b\u0076\u0027", _bba, _fff)
			return false
		}
	}
	return true
}
func ComparePNGFiles(file1, file2 string) (bool, error) {
	_ec, _bg := HashFile(file1)
	if _bg != nil {
		return false, _bg
	}
	_adf, _bg := HashFile(file2)
	if _bg != nil {
		return false, _bg
	}
	if _ec == _adf {
		return true, nil
	}
	_fcf, _bg := ReadPNG(file1)
	if _bg != nil {
		return false, _bg
	}
	_cge, _bg := ReadPNG(file2)
	if _bg != nil {
		return false, _bg
	}
	if _fcf.Bounds() != _cge.Bounds() {
		return false, nil
	}
	return CompareImages(_fcf, _cge)
}
func CompareImages(img1, img2 _ge.Image) (bool, error) {
	_gd := img1.Bounds()
	_cc := 0
	for _ff := 0; _ff < _gd.Size().X; _ff++ {
		for _fc := 0; _fc < _gd.Size().Y; _fc++ {
			_bb, _aff, _ea, _ := img1.At(_ff, _fc).RGBA()
			_eb, _ab, _fbc, _ := img2.At(_ff, _fc).RGBA()
			if _bb != _eb || _aff != _ab || _ea != _fbc {
				_cc++
			}
		}
	}
	_ggeg := float64(_cc) / float64(_gd.Dx()*_gd.Dy())
	if _ggeg > 0.0001 {
		_ce.Printf("\u0064\u0069\u0066f \u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0064\u0029\u000a", _ggeg, _cc)
		return false, nil
	}
	return true, nil
}

var (
	ErrRenderNotSupported = _fe.New("\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0073 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u006e\u0020\u0074\u0068\u0069\u0073\u0020\u0073\u0079\u0073\u0074\u0065m")
)

func HashFile(file string) (string, error) {
	_ee, _cg := _c.Open(file)
	if _cg != nil {
		return "", _cg
	}
	defer _ee.Close()
	_ggb := _a.New()
	if _, _cg = _fb.Copy(_ggb, _ee); _cg != nil {
		return "", _cg
	}
	return _fbf.EncodeToString(_ggb.Sum(nil)), nil
}
func ReadPNG(file string) (_ge.Image, error) {
	_ef, _fa := _c.Open(file)
	if _fa != nil {
		return nil, _fa
	}
	defer _ef.Close()
	return _g.Decode(_ef)
}
func ParseIndirectObjects(rawpdf string) (map[int64]_fed.PdfObject, error) {
	_fca := _fed.NewParserFromString(rawpdf)
	_adfd := map[int64]_fed.PdfObject{}
	for {
		_egb, _db := _fca.ParseIndirectObject()
		if _db != nil {
			if _db == _fb.EOF {
				break
			}
			return nil, _db
		}
		switch _de := _egb.(type) {
		case *_fed.PdfIndirectObject:
			_adfd[_de.ObjectNumber] = _egb
		case *_fed.PdfObjectStream:
			_adfd[_de.ObjectNumber] = _egb
		}
	}
	for _, _ca := range _adfd {
		_ade(_ca, _adfd)
	}
	return _adfd, nil
}
func RenderPDFToPNGs(pdfPath string, dpi int, outpathTpl string) error {
	if dpi <= 0 {
		dpi = 100
	}
	if _, _fd := _e.LookPath("\u0067\u0073"); _fd != nil {
		return ErrRenderNotSupported
	}
	return _e.Command("\u0067\u0073", "\u002d\u0073\u0044\u0045\u0056\u0049\u0043\u0045\u003d\u0070\u006e\u0067a\u006c\u0070\u0068\u0061", "\u002d\u006f", outpathTpl, _ce.Sprintf("\u002d\u0072\u0025\u0064", dpi), pdfPath).Run()
}
func _ade(_dg _fed.PdfObject, _gda map[int64]_fed.PdfObject) error {
	switch _cgb := _dg.(type) {
	case *_fed.PdfIndirectObject:
		_be := _cgb
		_ade(_be.PdfObject, _gda)
	case *_fed.PdfObjectDictionary:
		_gbc := _cgb
		for _, _egd := range _gbc.Keys() {
			_eaf := _gbc.Get(_egd)
			if _eab, _gba := _eaf.(*_fed.PdfObjectReference); _gba {
				_dc, _bab := _gda[_eab.ObjectNumber]
				if !_bab {
					return _fe.New("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074")
				}
				_gbc.Set(_egd, _dc)
			} else {
				_ade(_eaf, _gda)
			}
		}
	case *_fed.PdfObjectArray:
		_ffc := _cgb
		for _gbf, _ega := range _ffc.Elements() {
			if _aae, _cgd := _ega.(*_fed.PdfObjectReference); _cgd {
				_bde, _ag := _gda[_aae.ObjectNumber]
				if !_ag {
					return _fe.New("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074")
				}
				_ffc.Set(_gbf, _bde)
			} else {
				_ade(_ega, _gda)
			}
		}
	}
	return nil
}
func RunRenderTest(t *_f.T, pdfPath, outputDir, baselineRenderPath string, saveBaseline bool) {
	_cd := _df.TrimSuffix(_b.Base(pdfPath), _b.Ext(pdfPath))
	t.Run("\u0072\u0065\u006e\u0064\u0065\u0072", func(_bc *_f.T) {
		_geg := _b.Join(outputDir, _cd)
		_aa := _geg + "\u002d%\u0064\u002e\u0070\u006e\u0067"
		if _bcg := RenderPDFToPNGs(pdfPath, 0, _aa); _bcg != nil {
			_bc.Skip(_bcg)
		}
		for _dfd := 1; true; _dfd++ {
			_gb := _ce.Sprintf("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg", _geg, _dfd)
			_cdg := _b.Join(baselineRenderPath, _ce.Sprintf("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067", _cd, _dfd))
			if _, _fcb := _c.Stat(_gb); _fcb != nil {
				break
			}
			_bc.Logf("\u0025\u0073", _cdg)
			if saveBaseline {
				_bc.Logf("\u0043\u006fp\u0079\u0069\u006eg\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073", _gb, _cdg)
				_bd := CopyFile(_gb, _cdg)
				if _bd != nil {
					_bc.Fatalf("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _cdg, _bd)
				}
				continue
			}
			_bc.Run(_ce.Sprintf("\u0070\u0061\u0067\u0065\u0025\u0064", _dfd), func(_gf *_f.T) {
				_gf.Logf("\u0043o\u006dp\u0061\u0072\u0069\u006e\u0067 \u0025\u0073 \u0076\u0073\u0020\u0025\u0073", _gb, _cdg)
				_ga, _dff := ComparePNGFiles(_gb, _cdg)
				if _c.IsNotExist(_dff) {
					_gf.Fatal("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
				} else if !_ga {
					_gf.Fatal("\u0077\u0072\u006f\u006eg \u0070\u0061\u0067\u0065\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0065\u0064")
				}
			})
		}
	})
}
func CopyFile(src, dst string) error {
	_eg, _af := _c.Open(src)
	if _af != nil {
		return _af
	}
	defer _eg.Close()
	_gg, _af := _c.Create(dst)
	if _af != nil {
		return _af
	}
	defer _gg.Close()
	_, _af = _fb.Copy(_gg, _eg)
	return _af
}
