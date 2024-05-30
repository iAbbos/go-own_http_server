package writer

import (
	"github.com/codecrafters-io/http-server-starter-go/internal/entity"
	"io"
)

type Writer struct {
	writer io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{writer: w}
}

func (w *Writer) Write(resp *entity.Response) error {
	res := resp.Marshal()

	_, err := w.writer.Write(res)
	if err != nil {
		return err
	}

	return nil
}
