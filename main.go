package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	fmt.Println(err)

	if err != nil {
		return
	}
	htmlStr := `<!doctype html><html><head><title>WKHTMLTOPDF TEST</title></head><body>HELLO PDF</body></html>`

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(htmlStr)))

	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)
	// Create PDF document in internal buffer
	err = pdfg.Create()
	fmt.Println(err)

	if err != nil {
		log.Fatal(err)
	}

	//Your Pdf Name
	err = pdfg.WriteFile("./output.pdf")
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done")
}
