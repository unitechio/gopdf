# Go PDF

Golang PDF library for creating and processing PDF files (pure go)

[Examples](https://unidoc.io/unipdf-examples)

> 1. No `init` Required
> 2. No `LICENSE_API_KEY` Required

### usage

```go
package main

import (
	"fmt"
	"time"

	"bitbucket.org/shenghui0779/gopdf/common"
	"bitbucket.org/shenghui0779/gopdf/creator"
	"bitbucket.org/shenghui0779/gopdf/model"
)

func main() {
	err := RunPdfReport("sample.pdf")
	if err != nil {
		panic(err)
	}
}

func RunPdfReport(outputPath string) error {
	helvetica, _ := model.NewStandard14Font("Helvetica")
	helveticaBold, _ := model.NewStandard14Font("Helvetica-Bold")

	c := creator.New()
	c.SetPageMargins(50, 50, 100, 70)

	// Generate the table of contents.
	c.AddTOC = true
	toc := c.TOC()
	hstyle := c.NewTextStyle()
	hstyle.Color = creator.ColorRGBFromArithmetic(0.2, 0.2, 0.2)
	hstyle.FontSize = 28
	toc.SetHeading("Table of Contents", hstyle)
	lstyle := c.NewTextStyle()
	lstyle.FontSize = 14
	toc.SetLineStyle(lstyle)

	logoImg, err := c.NewImageFromURL("https://golang.google.cn/doc/gopher/pkg.png")
	if err != nil {
		return err
	}

	logoImg.ScaleToHeight(25)
	logoImg.SetPos(58, 20)

	DoDocumentControl(c, helvetica, helveticaBold)
	DoFeatureOverview(c, helvetica, helveticaBold)

	// Setup a front page (always placed first).
	c.CreateFrontPage(func(args creator.FrontpageFunctionArgs) {
		DoFirstPage(c, helvetica, helveticaBold)
	})

	// Draw a header on each page.
	c.DrawHeader(func(block *creator.Block, args creator.HeaderFunctionArgs) {
		// Draw the header on a block. The block size is the size of the page's top margins.
		block.Draw(logoImg)
	})

	// Draw footer on each page.
	c.DrawFooter(func(block *creator.Block, args creator.FooterFunctionArgs) {
		// Draw the on a block for each page.
		p := c.NewParagraph("gopdf")
		p.SetFont(helvetica)
		p.SetFontSize(8)
		p.SetPos(50, 20)
		p.SetColor(creator.ColorRGBFrom8bit(63, 68, 76))
		block.Draw(p)

		strPage := fmt.Sprintf("Page %d of %d", args.PageNum, args.TotalPages)
		p = c.NewParagraph(strPage)
		p.SetFont(helvetica)
		p.SetFontSize(8)
		p.SetPos(300, 20)
		p.SetColor(creator.ColorRGBFrom8bit(63, 68, 76))
		block.Draw(p)
	})

	return c.WriteToFile(outputPath)
}

// Generates the front page.
func DoFirstPage(c *creator.Creator, fontRegular *model.PdfFont, fontBold *model.PdfFont) {
	helvetica, _ := model.NewStandard14Font("Helvetica")
	helveticaBold, _ := model.NewStandard14Font("Helvetica-Bold")

	p := c.NewParagraph("GoPDF")
	p.SetFont(helvetica)
	p.SetFontSize(48)
	p.SetMargins(85, 0, 150, 0)
	p.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	c.Draw(p)

	p = c.NewParagraph("Example Report")
	p.SetFont(helveticaBold)
	p.SetFontSize(30)
	p.SetMargins(85, 0, 0, 0)
	p.SetColor(creator.ColorRGBFrom8bit(45, 148, 215))
	c.Draw(p)

	t := time.Now().UTC()
	dateStr := t.Format("2006-01-02 15:04:05")

	p = c.NewParagraph(dateStr)
	p.SetFont(helveticaBold)
	p.SetFontSize(12)
	p.SetMargins(90, 0, 5, 0)
	p.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	c.Draw(p)
}

// Document control page.
func DoDocumentControl(c *creator.Creator, fontRegular *model.PdfFont, fontBold *model.PdfFont) {
	ch := c.NewChapter("Document control")
	ch.SetMargins(0, 0, 40, 0)
	ch.GetHeading().SetFont(fontRegular)
	ch.GetHeading().SetFontSize(18)
	ch.GetHeading().SetColor(creator.ColorRGBFrom8bit(72, 86, 95))

	sc := ch.NewSubchapter("Issuer details")
	sc.GetHeading().SetFont(fontRegular)
	sc.GetHeading().SetFontSize(18)
	sc.GetHeading().SetColor(creator.ColorRGBFrom8bit(72, 86, 95))

	issuerTable := c.NewTable(2)
	issuerTable.SetMargins(0, 0, 30, 0)

	pColor := creator.ColorRGBFrom8bit(72, 86, 95)
	bgColor := creator.ColorRGBFrom8bit(56, 68, 67)

	p := c.NewParagraph("Issuer")
	p.SetFont(fontBold)
	p.SetFontSize(10)
	p.SetColor(creator.ColorWhite)
	cell := issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetBackgroundColor(bgColor)
	cell.SetContent(p)

	p = c.NewParagraph("GoPDF")
	p.SetFont(fontRegular)
	p.SetFontSize(10)
	p.SetColor(pColor)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetContent(p)

	p = c.NewParagraph("Address")
	p.SetFont(fontBold)
	p.SetFontSize(10)
	p.SetColor(creator.ColorWhite)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetBackgroundColor(bgColor)
	cell.SetContent(p)

	p = c.NewParagraph("Nanjing, China")
	p.SetFont(fontRegular)
	p.SetFontSize(10)
	p.SetColor(pColor)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetContent(p)

	p = c.NewParagraph("Email")
	p.SetFont(fontBold)
	p.SetFontSize(10)
	p.SetColor(creator.ColorWhite)
	cell = issuerTable.NewCell()
	cell.SetBackgroundColor(bgColor)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetContent(p)

	p = c.NewParagraph("example@xxx.io")
	p.SetFont(fontRegular)
	p.SetFontSize(10)
	p.SetColor(pColor)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetContent(p)

	p = c.NewParagraph("Web")
	p.SetFont(fontBold)
	p.SetFontSize(10)
	p.SetColor(creator.ColorWhite)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetBackgroundColor(bgColor)
	cell.SetContent(p)

	p = c.NewParagraph("example.io")
	p.SetFont(fontRegular)
	p.SetFontSize(10)
	p.SetColor(pColor)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetContent(p)

	p = c.NewParagraph("Author")
	p.SetFont(fontBold)
	p.SetFontSize(10)
	p.SetColor(creator.ColorWhite)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetBackgroundColor(bgColor)
	cell.SetContent(p)

	p = c.NewParagraph("GoPDF report generator")
	p.SetFont(fontRegular)
	p.SetFontSize(10)
	p.SetColor(pColor)
	cell = issuerTable.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetContent(p)

	sc.Add(issuerTable)

	// 1.2 - Document history
	sc = ch.NewSubchapter("Document History")
	sc.SetMargins(0, 0, 5, 0)
	sc.GetHeading().SetFont(fontRegular)
	sc.GetHeading().SetFontSize(18)
	sc.GetHeading().SetColor(pColor)

	histTable := c.NewTable(3)
	histTable.SetMargins(0, 0, 30, 50)

	histCols := []string{"Date Issued", "GoPDF Version", "Type/Change"}
	for _, histCol := range histCols {
		p = c.NewParagraph(histCol)
		p.SetFont(fontBold)
		p.SetFontSize(10)
		p.SetColor(creator.ColorWhite)
		cell = histTable.NewCell()
		cell.SetBackgroundColor(bgColor)
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
		cell.SetVerticalAlignment(creator.CellVerticalAlignmentMiddle)
		cell.SetContent(p)
	}

	dateStr := common.ReleasedAt.Format("2006-01-02 15:04:05")

	histVals := []string{dateStr, common.Version, "First issue"}
	for _, histVal := range histVals {
		p = c.NewParagraph(histVal)
		p.SetFont(fontRegular)
		p.SetFontSize(10)
		p.SetColor(pColor)
		cell = histTable.NewCell()
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
		cell.SetVerticalAlignment(creator.CellVerticalAlignmentMiddle)
		cell.SetContent(p)
	}

	sc.Add(histTable)

	if err := c.Draw(ch); err != nil {
		panic(err)
	}
}

// Chapter giving an overview of features.
// TODO: Add code snippets and show more styles and options.
func DoFeatureOverview(c *creator.Creator, fontRegular *model.PdfFont, fontBold *model.PdfFont) {
	// Ensure that the chapter starts on a new page.
	c.NewPage()

	ch := c.NewChapter("Feature overview")

	chapterFont := fontRegular
	chapterFontColor := creator.ColorRGBFrom8bit(72, 86, 95)
	chapterFontSize := 18.0

	normalFont := fontRegular
	normalFontColor := creator.ColorRGBFrom8bit(72, 86, 95)
	normalFontSize := 10.0

	bgColor := creator.ColorRGBFrom8bit(56, 68, 67)

	ch.GetHeading().SetFont(chapterFont)
	ch.GetHeading().SetFontSize(chapterFontSize)
	ch.GetHeading().SetColor(chapterFontColor)

	p := c.NewParagraph("This chapter demonstrates a few of the features of GoPDF that can be used for report generation.")
	p.SetFont(normalFont)
	p.SetFontSize(normalFontSize)
	p.SetColor(normalFontColor)
	p.SetMargins(0, 0, 5, 0)
	ch.Add(p)

	// Paragraphs.
	sc := ch.NewSubchapter("Paragraphs")
	sc.GetHeading().SetMargins(0, 0, 20, 0)
	sc.GetHeading().SetFont(chapterFont)
	sc.GetHeading().SetFontSize(chapterFontSize)
	sc.GetHeading().SetColor(chapterFontColor)

	p = c.NewParagraph("Paragraphs are used to represent text, as little as a single character, a word or multiple words forming multiple sentences GoPDF handles automatically wrapping those across lines and pages, making it relatively easy to work with. They can also be left, center, right aligned or justified as illustrated below:")
	p.SetFont(normalFont)
	p.SetFontSize(normalFontSize)
	p.SetColor(normalFontColor)
	p.SetMargins(0, 0, 5, 0)
	sc.Add(p)

	// Example paragraphs:
	loremTxt := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	alignments := []creator.TextAlignment{creator.TextAlignmentLeft, creator.TextAlignmentCenter, creator.TextAlignmentRight, creator.TextAlignmentJustify}
	for j := 0; j < 4; j++ {
		p = c.NewParagraph(loremTxt)
		p.SetFont(normalFont)
		p.SetFontSize(normalFontSize)
		p.SetColor(normalFontColor)
		p.SetMargins(20, 0, 10, 10)
		p.SetTextAlignment(alignments[j%4])

		sc.Add(p)
	}

	sc = ch.NewSubchapter("Tables")
	// Mock table: Priority table.
	priTable := c.NewTable(2)
	priTable.SetMargins(40, 40, 10, 0)
	// Column headers:
	tableCols := []string{"Priority", "Items fulfilled / available"}
	for _, tableCol := range tableCols {
		p = c.NewParagraph(tableCol)
		p.SetFont(fontBold)
		p.SetFontSize(10)
		p.SetColor(creator.ColorWhite)
		cell := priTable.NewCell()
		cell.SetBackgroundColor(bgColor)
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetContent(p)
	}
	items := [][]string{
		{"High", "52/80"},
		{"Medium", "32/100"},
		{"Low", "10/90"},
	}
	for _, lineItems := range items {
		for _, item := range lineItems {
			p = c.NewParagraph(item)
			p.SetFont(fontBold)
			p.SetFontSize(10)
			p.SetColor(creator.ColorWhite)
			cell := priTable.NewCell()
			cell.SetBackgroundColor(bgColor)
			cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
			cell.SetContent(p)
		}
	}
	sc.Add(priTable)

	sc = ch.NewSubchapter("Images")
	sc.GetHeading().SetMargins(0, 0, 20, 0)
	sc.GetHeading().SetFont(chapterFont)
	sc.GetHeading().SetFontSize(chapterFontSize)
	sc.GetHeading().SetColor(chapterFontColor)

	p = c.NewParagraph("Images can be loaded from multiple file formats, example from a PNG image:")
	p.SetFont(normalFont)
	p.SetFontSize(normalFontSize)
	p.SetColor(normalFontColor)
	p.SetMargins(0, 0, 5, 5)
	sc.Add(p)

	img, err := c.NewImageFromURL("https://golang.google.cn/blog/gopher/header.jpg")
	if err != nil {
		panic(err)
	}
	img.ScaleToHeight(50)
	sc.Add(img)

	sc = ch.NewSubchapter("Headers and footers")
	sc.GetHeading().SetMargins(0, 0, 20, 0)
	sc.GetHeading().SetFont(chapterFont)
	sc.GetHeading().SetFontSize(chapterFontSize)
	sc.GetHeading().SetColor(chapterFontColor)

	p = c.NewParagraph("Convenience functions are provided to generate headers and footers, see: Creator.DrawHeader and Creator.DrawFooter They both set a function that accepts a block which the header/footer is drawn on for each page. More information is provided in the arguments, allowing to skip header/footer on specific pages and showing page number and count.")
	p.SetFont(normalFont)
	p.SetFontSize(normalFontSize)
	p.SetColor(normalFontColor)
	p.SetMargins(0, 0, 5, 0)
	sc.Add(p)

	sc = ch.NewSubchapter("Table of contents generation")
	sc.GetHeading().SetMargins(0, 0, 20, 0)
	sc.GetHeading().SetFont(chapterFont)
	sc.GetHeading().SetFontSize(chapterFontSize)
	sc.GetHeading().SetColor(chapterFontColor)

	p = c.NewParagraph("A convenience function is provided to generate table of contents as can be seen on Creator.CreateTableOfContents and in our example code on unidoc.io.")
	p.SetFont(normalFont)
	p.SetFontSize(normalFontSize)
	p.SetColor(normalFontColor)
	p.SetMargins(0, 0, 5, 0)
	sc.Add(p)

	c.Draw(ch)
}
```