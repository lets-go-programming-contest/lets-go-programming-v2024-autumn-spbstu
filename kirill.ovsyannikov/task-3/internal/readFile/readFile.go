package readFile

import (
	"bytes"
	"encoding/xml"
	"os"
	"strconv"
	"strings"
	errors "task-3/internal/errorsExt"

	"golang.org/x/net/html/charset"
)

type Float32 float32

func (f *Float32) UnmarshalText(text []byte) error {
	strText := string(text)
	strText = strings.ReplaceAll(strText, ",", ".")

	value, err := strconv.ParseFloat(strText, 32)
	if err != nil {
		return errors.ErrorLocate(err)
	}

	*f = Float32(value)

	return nil
}

func ParseXML(inputFile string) (ValCurs, error) {
	var currencies ValCurs

	data, err := os.ReadFile(inputFile)
	if err != nil {
		return currencies, errors.ErrorLocate(err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&currencies)
	if err != nil {
		return currencies, errors.ErrorLocate(err)
	}

	return currencies, nil
}
