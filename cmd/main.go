package main

import (
	"fmt"
	"log"

	"github.com/unitechio/gopdf"
)

func main() {
	pdfPath := "contract.pdf"

	fmt.Println("🚀 Testing Final gopdf Solution")
	fmt.Println("===============================")

	// Create config
	config := gopdf.NewConfig()
	config.PreferredExtractor = "auto" // Auto-detect best method

	// Test 1: Find single anchor
	fmt.Println("\n1️⃣ Testing single anchor search...")
	anchor, err := gopdf.FindAnchor(pdfPath, "NGƯỜI ĐỀ NGHỊ", config)
	if err != nil {
		log.Printf("❌ FindAnchor failed: %v", err)
	} else {
		fmt.Printf("✅ Found anchor: %s\n", anchor.Text)
		fmt.Printf("   Page: %d, Position: (%.1f, %.1f), Line: %d\n",
			anchor.Page, anchor.X, anchor.Y, anchor.LineNumber)
	}

	// Test 2: Find multiple anchors
	fmt.Println("\n2️⃣ Testing multiple anchor search...")
	anchors := []string{
		"NGƯỜI ĐỀ NGHỊ",
		"TRƯỞNG PHÒNG",
		"KẾ TOÁN TRƯỞNG",
		"GIÁM ĐỐC",
		"Nguyễn Văn A",
	}

	results, err := gopdf.FindMultipleAnchors(pdfPath, anchors, config)
	if err != nil {
		log.Printf("❌ FindMultipleAnchors failed: %v", err)
	} else {
		fmt.Printf("✅ Found %d anchors:\n", len(results))
		for text, anchor := range results {
			fmt.Printf("   '%s' -> Line %d, Position (%.1f, %.1f)\n",
				text, anchor.LineNumber, anchor.X, anchor.Y)
		}
	}

	// Test 3: Calculate signature positions
	fmt.Println("\n3️⃣ Testing signature position calculation...")
	if anchor != nil {
		sigPos := gopdf.CalculateSignaturePosition(anchor, config)
		if sigPos != nil {
			fmt.Printf("✅ Signature position calculated:\n")
			fmt.Printf("   Page: %d, Position: (%.1f, %.1f), Size: (%.1f, %.1f)\n",
				sigPos.Page, sigPos.X, sigPos.Y, sigPos.W, sigPos.H)

			// Test 4: Add signature text
			fmt.Println("\n4️⃣ Testing signature text addition...")
			err = gopdf.AddSignatureText(
				pdfPath,
				"signed_contract.pdf",
				"John Doe - Signed",
				sigPos,
				"Arial",
				12.0,
			)
			if err != nil {
				log.Printf("❌ AddSignatureText failed: %v", err)
			} else {
				fmt.Println("✅ Signature text added successfully!")
				fmt.Println("   Output: signed_contract.pdf")
			}
		}
	}

	// Test 5: Regex search
	fmt.Println("\n5️⃣ Testing regex anchor search...")
	regexAnchor, err := gopdf.FindAnchorRegex(pdfPath, `(?i)nguyễn.*văn.*a`, config)
	if err != nil {
		log.Printf("❌ FindAnchorRegex failed: %v", err)
	} else {
		fmt.Printf("✅ Regex found: '%s' at line %d\n", regexAnchor.Text, regexAnchor.LineNumber)
	}

	// Test 6: Batch processing
	fmt.Println("\n6️⃣ Testing batch processing...")
	pdfFiles := []string{pdfPath} // Only one file for demo
	batchResults, batchErrors := gopdf.BatchProcessAnchors(pdfFiles, "TRƯỞNG PHÒNG", config)

	if len(batchErrors) > 0 {
		fmt.Printf("❌ Batch processing errors: %d\n", len(batchErrors))
		for _, err := range batchErrors {
			fmt.Printf("   %v\n", err)
		}
	}

	if len(batchResults) > 0 {
		fmt.Printf("✅ Batch processing successful: %d files\n", len(batchResults))
		for file, anchor := range batchResults {
			fmt.Printf("   %s -> '%s' at line %d\n", file, anchor.Text, anchor.LineNumber)
		}
	}

	fmt.Println("\n🎉 All tests completed!")
}
