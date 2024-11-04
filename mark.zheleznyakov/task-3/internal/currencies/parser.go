package currencies

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"os"
	"sort"
	"strings"
)

type Currencies struct {
	Currencies []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode  string  `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"Value" json:"value"`
}

func Parse(c *Currencies, f string) error {
	sort.Slice(c.Currencies, func(i, j int) bool {
		return c.Currencies[i].Value > c.Currencies[j].Value
	})

	cContent, err := os.ReadFile(f)
	if err != nil {
		return fmt.Errorf("your input file is cooked: %w", err)
	}

	cContent = []byte(strings.ReplaceAll(string(cContent), ",", "."))

	decoder := xml.NewDecoder(bytes.NewReader(cContent))
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
