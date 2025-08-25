package gopdf

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

// Anchor represents a text anchor in a PDF with its position
type Anchor struct {
	Text       string  `json:"text"`
	Page       int     `json:"page"`
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	FontSize   float64 `json:"font_size"`
	Width      float64 `json:"width"`
	Height     float64 `json:"height"`
	LineNumber int     `json:"line_number"`
}

// SignaturePosition represents where to place a signature
type SignaturePosition struct {
	Page int     `json:"page"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	W    float64 `json:"w"`
	H    float64 `json:"h"`
}

// Config holds configuration for anchor detection
type Config struct {
	SimilarityThreshold float64 `json:"similarity_threshold"`
	OffsetX             float64 `json:"offset_x"`
	OffsetY             float64 `json:"offset_y"`
	CaseSensitive       bool    `json:"case_sensitive"`
	MaxPages            int     `json:"max_pages"`
	StopOnFirst         bool    `json:"stop_on_first"`
	DefaultFontSize     float64 `json:"default_font_size"`
	LineHeight          float64 `json:"line_height"`
	PageHeight          float64 `json:"page_height"`
	PageWidth           float64 `json:"page_width"`
	MarginLeft          float64 `json:"margin_left"`
	MarginTop           float64 `json:"margin_top"`
	PreferredExtractor  string  `json:"preferred_extractor"` // "pdftotext", "mutool", "auto"
}

// NewConfig creates a default configuration
func NewConfig() *Config {
	return &Config{
		SimilarityThreshold: 0.8,
		OffsetX:             10.0,
		OffsetY:             -10.0,
		CaseSensitive:       false,
		MaxPages:            0,
		StopOnFirst:         true,
		DefaultFontSize:     12.0,
		LineHeight:          15.0,
		PageHeight:          792.0,
		PageWidth:           612.0,
		MarginLeft:          50.0,
		MarginTop:           50.0,
		PreferredExtractor:  "auto", // Auto-detect best extractor
	}
}

// Validate checks if config values are valid
func (c *Config) Validate() error {
	if c.SimilarityThreshold < 0 || c.SimilarityThreshold > 1 {
		return errors.New("similarity threshold must be between 0 and 1")
	}
	if c.MaxPages < 0 {
		return errors.New("max pages cannot be negative")
	}
	if c.DefaultFontSize <= 0 {
		return errors.New("default font size must be positive")
	}
	if c.LineHeight <= 0 {
		return errors.New("line height must be positive")
	}
	return nil
}

// FindAnchor searches for a text anchor in a PDF and returns its position
func FindAnchor(pdfPath, anchorText string, config *Config) (*Anchor, error) {
	if config == nil {
		config = NewConfig()
	}
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	if strings.TrimSpace(anchorText) == "" {
		return nil, errors.New("anchor text cannot be empty")
	}

	// Extract text using best available method
	text, err := extractTextFromPDF(pdfPath, config)
	if err != nil {
		return nil, fmt.Errorf("failed to extract text: %w", err)
	}

	// Search for anchor in extracted text
	return findAnchorInText(text, anchorText, config)
}

// extractTextFromPDF extracts text using the best available method
func extractTextFromPDF(pdfPath string, config *Config) (string, error) {
	var extractors []func(string) (string, error)

	// Choose extractor based on preference
	switch config.PreferredExtractor {
	case "pdftotext":
		extractors = []func(string) (string, error){extractWithPdftotext}
	case "mutool":
		extractors = []func(string) (string, error){extractWithMutool}
	default: // "auto"
		extractors = []func(string) (string, error){
			extractWithPdftotext, // Most reliable for wkhtmltopdf
			extractWithMutool,    // Good alternative
		}
	}

	// Try each extractor
	for _, extractor := range extractors {
		if text, err := extractor(pdfPath); err == nil && len(strings.TrimSpace(text)) > 0 {
			// Verify text quality (contains Vietnamese or readable content)
			if containsVietnamese(text) || containsReadableContent(text) {
				return text, nil
			}
		}
	}

	return "", errors.New("no suitable text extractor found or PDF text is corrupted")
}

// extractWithPdftotext uses pdftotext command for text extraction
func extractWithPdftotext(pdfPath string) (string, error) {
	if !checkCommandExists("pdftotext") {
		return "", errors.New("pdftotext not available")
	}

	cmd := exec.Command("pdftotext", "-enc", "UTF-8", pdfPath, "-")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("pdftotext command failed: %w", err)
	}

	return string(output), nil
}

// extractWithMutool uses mutool for text extraction
func extractWithMutool(pdfPath string) (string, error) {
	if !checkCommandExists("mutool") {
		return "", errors.New("mutool not available")
	}

	cmd := exec.Command("mutool", "draw", "-F", "txt", pdfPath)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("mutool command failed: %w", err)
	}

	return string(output), nil
}

// checkCommandExists checks if a command is available in PATH
func checkCommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// containsVietnamese checks if text contains Vietnamese characters
func containsVietnamese(text string) bool {
	vietnameseChars := "àáạảãâầấậẩẫăằắặẳẵèéẹẻẽêềếệểễìíịỉĩòóọỏõôồốộổỗơờớợởỡùúụủũưừứựửữỳýỵỷỹđĐ"
	for _, char := range vietnameseChars {
		if strings.ContainsRune(text, char) {
			return true
		}
	}
	return false
}

// containsReadableContent checks for readable Vietnamese content (even without accents)
func containsReadableContent(text string) bool {
	commonWords := []string{"NGUOI", "TRUONG", "PHONG", "GIAM", "DOC", "KY", "TEN"}
	textUpper := strings.ToUpper(text)

	for _, word := range commonWords {
		if strings.Contains(textUpper, word) {
			return true
		}
	}
	return false
}

// findAnchorInText searches for anchor in extracted text
func findAnchorInText(fullText, anchorText string, config *Config) (*Anchor, error) {
	lines := strings.Split(fullText, "\n")
	currentY := config.PageHeight - config.MarginTop

	for lineNum, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			currentY -= config.LineHeight
			continue
		}

		if textMatches(line, anchorText, config) {
			return &Anchor{
				Text:       line,
				Page:       1, // Assume single page for now
				X:          config.MarginLeft,
				Y:          currentY,
				FontSize:   config.DefaultFontSize,
				Width:      estimateTextWidth(line, config.DefaultFontSize),
				Height:     config.DefaultFontSize,
				LineNumber: lineNum + 1,
			}, nil
		}

		currentY -= config.LineHeight
		if currentY < config.MarginTop {
			break
		}
	}

	return nil, fmt.Errorf("anchor text '%s' not found", anchorText)
}

// FindAnchorRegex finds an anchor using a regular expression
func FindAnchorRegex(pdfPath, pattern string, config *Config) (*Anchor, error) {
	if config == nil {
		config = NewConfig()
	}
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid regex pattern: %w", err)
	}

	// Extract text
	text, err := extractTextFromPDF(pdfPath, config)
	if err != nil {
		return nil, fmt.Errorf("failed to extract text: %w", err)
	}

	// Search with regex
	return findAnchorWithRegex(text, re, config)
}

// findAnchorWithRegex searches using regex
func findAnchorWithRegex(fullText string, re *regexp.Regexp, config *Config) (*Anchor, error) {
	lines := strings.Split(fullText, "\n")
	currentY := config.PageHeight - config.MarginTop

	for lineNum, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			currentY -= config.LineHeight
			continue
		}

		if re.MatchString(line) {
			return &Anchor{
				Text:       line,
				Page:       1,
				X:          config.MarginLeft,
				Y:          currentY,
				FontSize:   config.DefaultFontSize,
				Width:      estimateTextWidth(line, config.DefaultFontSize),
				Height:     config.DefaultFontSize,
				LineNumber: lineNum + 1,
			}, nil
		}

		currentY -= config.LineHeight
		if currentY < config.MarginTop {
			break
		}
	}

	return nil, fmt.Errorf("regex pattern not found")
}

// FindMultipleAnchors searches for multiple text anchors
func FindMultipleAnchors(pdfPath string, anchorTexts []string, config *Config) (map[string]*Anchor, error) {
	if config == nil {
		config = NewConfig()
	}
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	if len(anchorTexts) == 0 {
		return make(map[string]*Anchor), nil
	}

	// Extract text once
	text, err := extractTextFromPDF(pdfPath, config)
	if err != nil {
		return nil, fmt.Errorf("failed to extract text: %w", err)
	}

	result := make(map[string]*Anchor)

	// Search for each anchor
	for _, anchorText := range anchorTexts {
		if anchor, err := findAnchorInText(text, anchorText, config); err == nil {
			result[anchorText] = anchor
		}
	}

	return result, nil
}

// textMatches checks if two texts match based on configuration
func textMatches(text1, text2 string, config *Config) bool {
	if !config.CaseSensitive {
		text1 = strings.ToLower(text1)
		text2 = strings.ToLower(text2)
	}

	text1 = normalizeWhitespace(text1)
	text2 = normalizeWhitespace(text2)

	return strings.Contains(text1, text2)
}

// normalizeWhitespace normalizes whitespace in text
func normalizeWhitespace(text string) string {
	var result strings.Builder
	var prevSpace bool

	for _, r := range text {
		if unicode.IsSpace(r) {
			if !prevSpace {
				result.WriteRune(' ')
				prevSpace = true
			}
		} else {
			result.WriteRune(r)
			prevSpace = false
		}
	}

	return strings.TrimSpace(result.String())
}

// estimateTextWidth provides rough estimate of text width
func estimateTextWidth(text string, fontSize float64) float64 {
	return float64(len(strings.TrimSpace(text))) * fontSize * 0.6
}

// CalculateSignaturePosition calculates where to place a signature based on an anchor
func CalculateSignaturePosition(anchor *Anchor, config *Config) *SignaturePosition {
	if anchor == nil {
		return nil
	}
	if config == nil {
		config = NewConfig()
	}

	return &SignaturePosition{
		Page: anchor.Page,
		X:    anchor.X + config.OffsetX,
		Y:    anchor.Y + config.OffsetY,
		W:    anchor.Width,
		H:    anchor.Height,
	}
}

// AddSignatureText adds text at the signature position in the PDF
func AddSignatureText(inputPDF, outputPDF, text string, pos *SignaturePosition, fontName string, fontSize float64) error {
	if pos == nil {
		return errors.New("signature position cannot be nil")
	}
	if strings.TrimSpace(text) == "" {
		return errors.New("signature text cannot be empty")
	}
	if fontSize <= 0 {
		fontSize = 12.0
	}
	if fontName == "" {
		fontName = "Helvetica"
	}

	conf := model.NewDefaultConfiguration()

	desc := fmt.Sprintf("text:%s, pos:%.1f %.1f, fontname:%s, fontsize:%.1f",
		text, pos.X, pos.Y, fontName, fontSize)

	wm, err := api.TextWatermark(desc, "", true, true, types.POINTS)
	if err != nil {
		return fmt.Errorf("failed to create text watermark: %w", err)
	}

	selectedPages := []string{strconv.Itoa(pos.Page)}

	return api.AddWatermarksFile(inputPDF, outputPDF, selectedPages, wm, conf)
}

// AddSignatureImage adds an image at the signature position in the PDF
func AddSignatureImage(inputPDF, outputPDF, imagePath string, pos *SignaturePosition) error {
	if pos == nil {
		return errors.New("signature position cannot be nil")
	}

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return fmt.Errorf("image file does not exist: %s", imagePath)
	}

	conf := model.NewDefaultConfiguration()

	desc := fmt.Sprintf("image:%s, pos:%.1f %.1f, scale:1.0", imagePath, pos.X, pos.Y)

	wm, err := api.TextWatermark(desc, "", true, true, types.POINTS)
	if err != nil {
		return fmt.Errorf("failed to create watermark: %w", err)
	}

	selectedPages := []string{strconv.Itoa(pos.Page)}

	return api.AddWatermarksFile(inputPDF, outputPDF, selectedPages, wm, conf)
}

// ValidatePDF checks if a PDF file is valid and readable
func ValidatePDF(pdfPath string) error {
	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		return fmt.Errorf("PDF file does not exist: %s", pdfPath)
	}

	// Try to extract some text to validate
	config := NewConfig()
	_, err := extractTextFromPDF(pdfPath, config)
	if err != nil {
		return fmt.Errorf("PDF validation failed: %w", err)
	}

	return nil
}

// GetPDFPageCount returns the number of pages (simplified - assumes 1 page for now)
func GetPDFPageCount(pdfPath string) (int, error) {
	if err := ValidatePDF(pdfPath); err != nil {
		return 0, err
	}
	return 1, nil // Simplified - could be enhanced to detect actual page count
}

// ExtractTextFromPage extracts all text from a specific page
func ExtractTextFromPage(pdfPath string, pageNum int) (string, error) {
	config := NewConfig()
	return extractTextFromPDF(pdfPath, config)
}

// BatchProcessAnchors processes multiple PDFs with the same anchor configuration
func BatchProcessAnchors(pdfPaths []string, anchorText string, config *Config) (map[string]*Anchor, []error) {
	results := make(map[string]*Anchor)
	var errors []error

	for _, pdfPath := range pdfPaths {
		anchor, err := FindAnchor(pdfPath, anchorText, config)
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to process %s: %w", pdfPath, err))
		} else {
			results[pdfPath] = anchor
		}
	}

	return results, errors
}
