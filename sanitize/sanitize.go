package sanitize

import (
	_e "bitbucket.org/shenghui0779/gopdf/common"
	_g "bitbucket.org/shenghui0779/gopdf/core"
)

// SanitizationOpts specifies the objects to be removed during sanitization.
type SanitizationOpts struct {
	// JavaScript specifies wether JavaScript action should be removed. JavaScript Actions, section 12.6.4.16 of PDF32000_2008
	JavaScript bool

	// URI specifies if URI actions should be removed. 12.6.4.7 URI Actions, PDF32000_2008.
	URI bool

	// GoToR removes remote GoTo actions. 12.6.4.3 Remote Go-To Actions, PDF32000_2008.
	GoToR bool

	// GoTo specifies wether GoTo actions should be removed. 12.6.4.2 Go-To Actions, PDF32000_2008.
	GoTo bool

	// RenditionJS enables removing of `JS` entry from a Rendition Action.
	// The `JS` entry has a value of text string or stream containing a JavaScript script that shall be executed when the action is triggered.
	// 12.6.4.13 Rendition Actions Table 214, PDF32000_2008.
	RenditionJS bool

	// OpenAction removes OpenAction entry from the document catalog.
	OpenAction bool

	// Launch specifies wether Launch Action should be removed.
	// A launch action launches an application or opens or prints a document.
	// 12.6.4.5 Launch Actions, PDF32000_2008.
	Launch bool
}

func (_gg *Sanitizer) processObjects(_f []_g.PdfObject) ([]_g.PdfObject, error) {
	_ee := []_g.PdfObject{}
	_fa := _gg._ef
	for _, _bg := range _f {
		switch _cb := _bg.(type) {
		case *_g.PdfIndirectObject:
			_bb, _eb := _g.GetDict(_cb)
			if _eb {
				if _ge, _a := _g.GetName(_bb.Get("\u0054\u0079\u0070\u0065")); _a && *_ge == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
					if _, _ea := _g.GetIndirect(_bb.Get("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e")); _ea && _fa.OpenAction {
						_bb.Remove("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e")
					}
				} else if _ba, _eg := _g.GetName(_bb.Get("\u0053")); _eg {
					switch *_ba {
					case "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074":
						if _fa.JavaScript {
							if _ad, _gf := _g.GetStream(_bb.Get("\u004a\u0053")); _gf {
								_ac := []byte{}
								_eae, _fe := _g.MakeStream(_ac, nil)
								if _fe == nil {
									*_ad = *_eae
								}
							}
							_e.Log.Debug("\u004a\u0061\u0076\u0061\u0073\u0063\u0072\u0069\u0070\u0074\u0020a\u0063\u0074\u0069\u006f\u006e\u0020\u0073\u006b\u0069\u0070p\u0065\u0064\u002e")
							continue
						}
					case "\u0055\u0052\u0049":
						if _fa.URI {
							_e.Log.Debug("\u0055\u0052\u0049\u0020ac\u0074\u0069\u006f\u006e\u0020\u0073\u006b\u0069\u0070\u0070\u0065\u0064\u002e")
							continue
						}
					case "\u0047\u006f\u0054\u006f":
						if _fa.GoTo {
							_e.Log.Debug("G\u004fT\u004f\u0020\u0061\u0063\u0074\u0069\u006f\u006e \u0073\u006b\u0069\u0070pe\u0064\u002e")
							continue
						}
					case "\u0047\u006f\u0054o\u0052":
						if _fa.GoToR {
							_e.Log.Debug("R\u0065\u006d\u006f\u0074\u0065\u0020G\u006f\u0054\u004f\u0020\u0061\u0063\u0074\u0069\u006fn\u0020\u0073\u006bi\u0070p\u0065\u0064\u002e")
							continue
						}
					case "\u004c\u0061\u0075\u006e\u0063\u0068":
						if _fa.Launch {
							_e.Log.Debug("\u004a\u0061\u0076\u0061\u0073\u0063\u0072\u0069\u0070\u0074\u0020a\u0063\u0074\u0069\u006f\u006e\u0020\u0073\u006b\u0069\u0070p\u0065\u0064\u002e")
							continue
						}
					case "\u0052e\u006e\u0064\u0069\u0074\u0069\u006fn":
						if _bbe, _bge := _g.GetStream(_bb.Get("\u004a\u0053")); _bge {
							_ab := []byte{}
							_cc, _bd := _g.MakeStream(_ab, nil)
							if _bd == nil {
								*_bbe = *_cc
							}
						}
					}
				} else if _eff := _bb.Get("\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"); _eff != nil && _fa.JavaScript {
					continue
				} else if _bga, _ebg := _g.GetName(_bb.Get("\u0054\u0079\u0070\u0065")); _ebg && *_bga == "\u0041\u006e\u006eo\u0074" && _fa.JavaScript {
					if _fee, _db := _g.GetIndirect(_bb.Get("\u0050\u0061\u0072\u0065\u006e\u0074")); _db {
						if _ega, _egf := _g.GetDict(_fee.PdfObject); _egf {
							if _fd, _bag := _g.GetDict(_ega.Get("\u0041\u0041")); _bag {
								_abf, _fg := _g.GetIndirect(_fd.Get("\u004b"))
								if _fg {
									if _gfd, _feg := _g.GetDict(_abf.PdfObject); _feg {
										if _bf, _ca := _g.GetName(_gfd.Get("\u0053")); _ca && *_bf == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" {
											_gfd.Clear()
										} else if _fgf := _fd.Get("\u0046"); _fgf != nil {
											if _df, _fgb := _g.GetIndirect(_fgf); _fgb {
												if _eea, _bc := _g.GetDict(_df.PdfObject); _bc {
													if _da, _bba := _g.GetName(_eea.Get("\u0053")); _bba && *_da == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" {
														_eea.Clear()
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		case *_g.PdfObjectStream:
			_e.Log.Debug("\u0070d\u0066\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0073t\u0072e\u0061m\u0020\u0074\u0079\u0070\u0065\u0020\u0025T", _cb)
		case *_g.PdfObjectStreams:
			_e.Log.Debug("\u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0074\u0072\u0065\u0061\u006d\u0073\u0020\u0074\u0079\u0070e\u0020\u0025\u0054", _cb)
		default:
			_e.Log.Debug("u\u006e\u006b\u006e\u006fwn\u0020p\u0064\u0066\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0025\u0054", _cb)
		}
		_ee = append(_ee, _bg)
	}
	_gg.analyze(_ee)
	return _ee, nil
}

// New returns a new sanitizer object.
func New(opts SanitizationOpts) *Sanitizer { return &Sanitizer{_ef: opts} }

func (_gee *Sanitizer) analyze(_aba []_g.PdfObject) {
	_eec := map[string]int{}
	for _, _ga := range _aba {
		switch _gc := _ga.(type) {
		case *_g.PdfIndirectObject:
			_eaed, _ada := _g.GetDict(_gc.PdfObject)
			if _ada {
				if _cg, _abd := _g.GetName(_eaed.Get("\u0054\u0079\u0070\u0065")); _abd && *_cg == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
					if _, _aa := _g.GetIndirect(_eaed.Get("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e")); _aa {
						_eec["\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e"]++
					}
				} else if _eee, _bde := _g.GetName(_eaed.Get("\u0053")); _bde {
					_baf := _eee.String()
					if _baf == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" || _baf == "\u0055\u0052\u0049" || _baf == "\u0047\u006f\u0054\u006f" || _baf == "\u0047\u006f\u0054o\u0052" || _baf == "\u004c\u0061\u0075\u006e\u0063\u0068" {
						_eec[_baf]++
					} else if _baf == "\u0052e\u006e\u0064\u0069\u0074\u0069\u006fn" {
						if _, _gce := _g.GetStream(_eaed.Get("\u004a\u0053")); _gce {
							_eec[_baf]++
						}
					}
				} else if _cce := _eaed.Get("\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"); _cce != nil {
					_eec["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++
				} else if _gfde, _cd := _g.GetIndirect(_eaed.Get("\u0050\u0061\u0072\u0065\u006e\u0074")); _cd {
					if _ed, _baa := _g.GetDict(_gfde.PdfObject); _baa {
						if _abb, _fc := _g.GetDict(_ed.Get("\u0041\u0041")); _fc {
							_cgc := _abb.Get("\u004b")
							_ebc, _ggb := _g.GetIndirect(_cgc)
							if _ggb {
								if _aed, _dfb := _g.GetDict(_ebc.PdfObject); _dfb {
									if _bfbb, _efg := _g.GetName(_aed.Get("\u0053")); _efg && *_bfbb == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" {
										_eec["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++
									} else if _, _ged := _g.GetString(_aed.Get("\u004a\u0053")); _ged {
										_eec["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++
									} else {
										_ebe := _abb.Get("\u0046")
										if _ebe != nil {
											_cfe, _aec := _g.GetIndirect(_ebe)
											if _aec {
												if _fda, _fdd := _g.GetDict(_cfe.PdfObject); _fdd {
													if _ff, _geb := _g.GetName(_fda.Get("\u0053")); _geb {
														_dg := _ff.String()
														_eec[_dg]++
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	_gee._c = _eec
}

// Sanitizer represents a sanitizer object.
// It implements the Optimizer interface to access the objects field from the writer.
type Sanitizer struct {
	_ef SanitizationOpts
	_c  map[string]int
}

// Optimize optimizes `objects` and returns updated list of objects.
func (_gd *Sanitizer) Optimize(objects []_g.PdfObject) ([]_g.PdfObject, error) {
	return _gd.processObjects(objects)
}

// GetSuspiciousObjects returns a count of each detected suspicious object.
func (_dgc *Sanitizer) GetSuspiciousObjects() map[string]int { return _dgc._c }
