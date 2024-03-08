package pdf

import (
	"bytes"
	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

const (
	defaultDPI = 300
)

func GeneratePDF(data interface{}, dirHTML string) ([]byte, error) {
	var (
		templ *template.Template
		err   error
		body  bytes.Buffer
	)

	if templ, err = template.ParseFiles(dirHTML); err != nil {
		return nil, err
	}

	if err = templ.Execute(&body, data); err != nil {
		return nil, err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	page := wkhtmltopdf.NewPageReader(&body)

	// Aktifkan ini jika file HTML mengandung referensi lokal seperti gambar, CSS, dll.
	page.EnableLocalFileAccess.Set(true)

	// Tambahkan halaman ke generator
	pdfg.AddPage(page)

	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.Dpi.Set(defaultDPI) // Gunakan konstanta untuk DPI
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}
