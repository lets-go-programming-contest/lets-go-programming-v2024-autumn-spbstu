package currencies

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type Float64 float64
func (f *Float64) UnmarshalText(text []byte) error {
	strValue := string(text)
	strValue = strings.ReplaceAll(strValue, ",", ".")
	value, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		return err
	}
	*f = Float64(value)
	return nil
}

func UnmarshalXML(c *Currencies, f string) error {
	text, err := os.ReadFile(f)
	if err != nil {
		return fmt.Errorf("your input file is cooked: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(text))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unsupported charset: %s", charset)
	}

	err = decoder.Decode(&c)
	if err != nil {
		return fmt.Errorf("your input file contents are cooked: %w", err)
	}

	return nil
}
