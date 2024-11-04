package currency

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func FetchCurrencyData(filePath string) ([]Currency, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file: %w", err)
	}

	content = []byte(strings.ReplaceAll(string(content), ",", "."))

	decoder := xml.NewDecoder(bytes.NewReader(content))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unsupported charset: %s", charset)
	}

	var data struct {
		Currencies []Currency `xml:"Valute"`
	}

	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode XML: %w", err)
	}

	return data.Currencies, nil
}
