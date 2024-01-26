package pdf

import (
	"bytes"
	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
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

	// Terapkan data template HTML yang telah dianalisis dan simpan hasilnya dalam sebuah buffer
	if err = templ.Execute(&body, data); err != nil {
		return nil, err
	}

	// Inisialisasi generator wkhtmltopdf
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	// Baca halaman HTML sebagai halaman PDF
	page := wkhtmltopdf.NewPageReader(&body)

	// Aktifkan ini jika file HTML mengandung referensi lokal seperti gambar, CSS, dll.
	page.EnableLocalFileAccess.Set(true)

	// Tambahkan halaman ke generator
	pdfg.AddPage(page)

	// Manipulasi atribut halaman sesuai kebutuhan
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	// Buat PDF
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}
