package file

import (
	"mime/multipart"

	"github.com/gabriel-vasile/mimetype"
)

func ReadMimetypeFile(
	file *multipart.FileHeader,
) (mtype, ext string, size int64, err error) {
	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()

	mt, err := mimetype.DetectReader(src)
	mtype = mt.String()
	ext = mt.Extension()
	size = file.Size

	return
}
