package gopdf

// "regexp"
// "strings"

// TextPosition represents the exact position of text in a PDF
type TextPosition struct {
	Text     string
	Page     int
	X        float64
	Y        float64
	Width    float64
	Height   float64
	FontSize float64
}

// // FindAnchorWithPosition finds an anchor and returns its exact position
// func FindAnchorWithPosition(pdfPath, anchorText string, config *Config) (*Anchor, error) {
// 	if config == nil {
// 		config = NewConfig()
// 	}

// 	ctx, err := api.ReadContextFile(pdfPath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Process each page
// 	for i := 1; i <= ctx.PageCount; i++ {
// 		positions, err := extractTextPositions(ctx, i)
// 		if err != nil {
// 			continue
// 		}

// 		for _, pos := range positions {
// 			if textMatches(pos.Text, anchorText, config) {
// 				return &Anchor{
// 					Text:     pos.Text,
// 					Page:     pos.Page,
// 					X:        pos.X,
// 					Y:        pos.Y,
// 					FontSize: pos.FontSize,
// 				}, nil
// 			}
// 		}
// 	}

// 	return nil, errors.New("anchor text not found")
// }

// extractTextPositions extracts text with positions from a page
// // This is a simplified implementation - a full implementation would parse the PDF content stream
// func extractTextPositions(ctx *api.Context, page int) ([]TextPosition, error) {
// 	// Extract text from page
// 	text, _, _, err := api.ExtractTextForPage(ctx, page)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// In a real implementation, we would parse the content stream to get exact positions
// 	// This is a placeholder that just returns a basic position
// 	positions := []TextPosition{
// 		{
// 			Text:     text,
// 			Page:     page,
// 			X:        72.0,  // 1 inch from left
// 			Y:        720.0, // 10 inches from bottom (assuming 792pt tall page)
// 			Width:    100.0,
// 			Height:   12.0,
// 			FontSize: 12.0,
// 		},
// 	}

// 	return positions, nil
// }

// // textMatches checks if two texts match based on configuration
// func textMatches(text1, text2 string, config *Config) bool {
// 	if config.CaseSensitive {
// 		return strings.Contains(text1, text2)
// 	}

// 	// Normalize whitespace
// 	re := regexp.MustCompile(`\s+`)
// 	text1 = re.ReplaceAllString(strings.ToLower(text1), " ")
// 	text2 = re.ReplaceAllString(strings.ToLower(text2), " ")

// 	return strings.Contains(text1, text2)
// }

// // FindAnchorRegex finds an anchor using a regular expression
// func FindAnchorRegex(pdfPath, pattern string, config *Config) (*Anchor, error) {
// 	if config == nil {
// 		config = NewConfig()
// 	}

// 	ctx, err := api.ReadContextFile(pdfPath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	re, err := regexp.Compile(pattern)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Process each page
// 	for i := 1; i <= ctx.PageCount; i++ {
// 		positions, err := extractTextPositions(ctx, i)
// 		if err != nil {
// 			continue
// 		}

// 		for _, pos := range positions {
// 			if re.MatchString(pos.Text) {
// 				return &Anchor{
// 					Text:     pos.Text,
// 					Page:     pos.Page,
// 					X:        pos.X,
// 					Y:        pos.Y,
// 					FontSize: pos.FontSize,
// 				}, nil
// 			}
// 		}
// 	}

// 	return nil, errors.New("anchor pattern not found")
// }
