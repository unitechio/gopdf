package sanitize

import (
	_e "bitbucket.org/shenghui0779/gopdf/common"
	_f "bitbucket.org/shenghui0779/gopdf/core"
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

// GetSuspiciousObjects returns a count of each detected suspicious object.
func (_dfg *Sanitizer) GetSuspiciousObjects() map[string]int { return _dfg._ee }

// New returns a new sanitizer object.
func New(opts SanitizationOpts) *Sanitizer { return &Sanitizer{_a: opts} }

// Optimize optimizes `objects` and returns updated list of objects.
func (_d *Sanitizer) Optimize(objects []_f.PdfObject) ([]_f.PdfObject, error) {
	return _d.processObjects(objects)
}
func (_gf *Sanitizer) processObjects(_c []_f.PdfObject) ([]_f.PdfObject, error) {
	_eea := []_f.PdfObject{}
	_ec := _gf._a
	for _, _gb := range _c {
		switch _df := _gb.(type) {
		case *_f.PdfIndirectObject:
			_eb, _ef := _f.GetDict(_df)
			if _ef {
				if _ed, _da := _f.GetName(_eb.Get("\u0054\u0079\u0070\u0065")); _da && *_ed == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
					if _, _gd := _f.GetIndirect(_eb.Get("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e")); _gd && _ec.OpenAction {
						_eb.Remove("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e")
					}
				} else if _edf, _cd := _f.GetName(_eb.Get("\u0053")); _cd {
					switch *_edf {
					case "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074":
						if _ec.JavaScript {
							if _ab, _gfc := _f.GetStream(_eb.Get("\u004a\u0053")); _gfc {
								_ca := []byte{}
								_ga, _aa := _f.MakeStream(_ca, nil)
								if _aa == nil {
									*_ab = *_ga
								}
							}
							_e.Log.Debug("\u004a\u0061\u0076\u0061\u0073\u0063\u0072\u0069\u0070\u0074\u0020a\u0063\u0074\u0069\u006f\u006e\u0020\u0073\u006b\u0069\u0070p\u0065\u0064\u002e")
							continue
						}
					case "\u0055\u0052\u0049":
						if _ec.URI {
							_e.Log.Debug("\u0055\u0052\u0049\u0020ac\u0074\u0069\u006f\u006e\u0020\u0073\u006b\u0069\u0070\u0070\u0065\u0064\u002e")
							continue
						}
					case "\u0047\u006f\u0054\u006f":
						if _ec.GoTo {
							_e.Log.Debug("G\u004fT\u004f\u0020\u0061\u0063\u0074\u0069\u006f\u006e \u0073\u006b\u0069\u0070pe\u0064\u002e")
							continue
						}
					case "\u0047\u006f\u0054o\u0052":
						if _ec.GoToR {
							_e.Log.Debug("R\u0065\u006d\u006f\u0074\u0065\u0020G\u006f\u0054\u004f\u0020\u0061\u0063\u0074\u0069\u006fn\u0020\u0073\u006bi\u0070p\u0065\u0064\u002e")
							continue
						}
					case "\u004c\u0061\u0075\u006e\u0063\u0068":
						if _ec.Launch {
							_e.Log.Debug("\u004a\u0061\u0076\u0061\u0073\u0063\u0072\u0069\u0070\u0074\u0020a\u0063\u0074\u0069\u006f\u006e\u0020\u0073\u006b\u0069\u0070p\u0065\u0064\u002e")
							continue
						}
					case "\u0052e\u006e\u0064\u0069\u0074\u0069\u006fn":
						if _ge, _ac := _f.GetStream(_eb.Get("\u004a\u0053")); _ac {
							_dff := []byte{}
							_cb, _cda := _f.MakeStream(_dff, nil)
							if _cda == nil {
								*_ge = *_cb
							}
						}
					}
				} else if _gg := _eb.Get("\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"); _gg != nil && _ec.JavaScript {
					continue
				} else if _cae, _ebb := _f.GetName(_eb.Get("\u0054\u0079\u0070\u0065")); _ebb && *_cae == "\u0041\u006e\u006eo\u0074" && _ec.JavaScript {
					if _cg, _ebf := _f.GetIndirect(_eb.Get("\u0050\u0061\u0072\u0065\u006e\u0074")); _ebf {
						if _b, _ebba := _f.GetDict(_cg.PdfObject); _ebba {
							if _dd, _eec := _f.GetDict(_b.Get("\u0041\u0041")); _eec {
								_abg, _fe := _f.GetIndirect(_dd.Get("\u004b"))
								if _fe {
									if _eac, _ba := _f.GetDict(_abg.PdfObject); _ba {
										if _fb, _ce := _f.GetName(_eac.Get("\u0053")); _ce && *_fb == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" {
											_eac.Clear()
										} else if _dfb := _dd.Get("\u0046"); _dfb != nil {
											if _dc, _bc := _f.GetIndirect(_dfb); _bc {
												if _be, _feg := _f.GetDict(_dc.PdfObject); _feg {
													if _fa, _fc := _f.GetName(_be.Get("\u0053")); _fc && *_fa == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" {
														_be.Clear()
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
		case *_f.PdfObjectStream:
			_e.Log.Debug("\u0070d\u0066\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0073t\u0072e\u0061m\u0020\u0074\u0079\u0070\u0065\u0020\u0025T", _df)
		case *_f.PdfObjectStreams:
			_e.Log.Debug("\u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0074\u0072\u0065\u0061\u006d\u0073\u0020\u0074\u0079\u0070e\u0020\u0025\u0054", _df)
		default:
			_e.Log.Debug("u\u006e\u006b\u006e\u006fwn\u0020p\u0064\u0066\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0025\u0054", _df)
		}
		_eea = append(_eea, _gb)
	}
	_gf.analyze(_eea)
	return _eea, nil
}

// Sanitizer represents a sanitizer object.
// It implements the Optimizer interface to access the objects field from the writer.
type Sanitizer struct {
	_a  SanitizationOpts
	_ee map[string]int
}

func (_fd *Sanitizer) analyze(_bd []_f.PdfObject) {
	_cf := map[string]int{}
	for _, _beb := range _bd {
		switch _gbg := _beb.(type) {
		case *_f.PdfIndirectObject:
			_dda, _cea := _f.GetDict(_gbg.PdfObject)
			if _cea {
				if _ag, _fca := _f.GetName(_dda.Get("\u0054\u0079\u0070\u0065")); _fca && *_ag == "\u0043a\u0074\u0061\u006c\u006f\u0067" {
					if _, _bda := _f.GetIndirect(_dda.Get("\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e")); _bda {
						_cf["\u004f\u0070\u0065\u006e\u0041\u0063\u0074\u0069\u006f\u006e"]++
					}
				} else if _caf, _bce := _f.GetName(_dda.Get("\u0053")); _bce {
					_abe := _caf.String()
					if _abe == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" || _abe == "\u0055\u0052\u0049" || _abe == "\u0047\u006f\u0054\u006f" || _abe == "\u0047\u006f\u0054o\u0052" || _abe == "\u004c\u0061\u0075\u006e\u0063\u0068" {
						_cf[_abe]++
					} else if _abe == "\u0052e\u006e\u0064\u0069\u0074\u0069\u006fn" {
						if _, _cbc := _f.GetStream(_dda.Get("\u004a\u0053")); _cbc {
							_cf[_abe]++
						}
					}
				} else if _fcb := _dda.Get("\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"); _fcb != nil {
					_cf["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++
				} else if _eda, _ad := _f.GetIndirect(_dda.Get("\u0050\u0061\u0072\u0065\u006e\u0074")); _ad {
					if _fde, _aaa := _f.GetDict(_eda.PdfObject); _aaa {
						if _fg, _fbc := _f.GetDict(_fde.Get("\u0041\u0041")); _fbc {
							_bac := _fg.Get("\u004b")
							_fcg, _gc := _f.GetIndirect(_bac)
							if _gc {
								if _gbc, _gcg := _f.GetDict(_fcg.PdfObject); _gcg {
									if _agc, _ceag := _f.GetName(_gbc.Get("\u0053")); _ceag && *_agc == "\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074" {
										_cf["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++
									} else if _, _ae := _f.GetString(_gbc.Get("\u004a\u0053")); _ae {
										_cf["\u004a\u0061\u0076\u0061\u0053\u0063\u0072\u0069\u0070\u0074"]++
									} else {
										_abeg := _fg.Get("\u0046")
										if _abeg != nil {
											_cac, _caef := _f.GetIndirect(_abeg)
											if _caef {
												if _eaa, _faf := _f.GetDict(_cac.PdfObject); _faf {
													if _de, _dffg := _f.GetName(_eaa.Get("\u0053")); _dffg {
														_fac := _de.String()
														_cf[_fac]++
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
	_fd._ee = _cf
}
