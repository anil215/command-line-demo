package writer

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type impl struct {
}

type Writer interface {
	WriteToTxt(filename, filepath, data string) error
}

func NewWriter() Writer {
	return &impl{}
}

func (w *impl) WriteToTxt(filename, filepath, data string) error {
	f, err := os.Create(fmt.Sprintf("%s/%s", filepath, filename))

	if err != nil {
		return errors.Wrap(err, "failed to create file")
	}

	defer f.Close()

	_, err = f.WriteString(data)

	if err != nil {
		return errors.Wrap(err, "failed to write data to disk")
	}

	return nil
}
