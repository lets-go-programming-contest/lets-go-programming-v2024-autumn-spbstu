package currency

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	errProc "github.com/IDevFrye/task-3/internal/errors"
	"golang.org/x/text/encoding/charmap"
)

func FetchCurrencyData(filePath string) ([]Currency, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errProc.ErrInputFileNotFound, err.Error())
	}

	if len(content) == 0 {
		return nil, errProc.ErrEmptyInputFile
	}

	content = []byte(strings.ReplaceAll(string(content), ",", "."))

	decoder := xml.NewDecoder(bytes.NewReader(content))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("%w: %s", errProc.ErrUnsupportedCharset, charset)
	}

	var data struct {
		Currencies []Currency `xml:"Valute"`
	}

	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("%w: %s", errProc.ErrInvalidXMLFormat, err.Error())
	}

	return data.Currencies, nil
}
