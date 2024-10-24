package currencies

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type Currencies struct {
	Currencies []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode  string  `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Value    float64 `xml:"Value"`
}

func Parse(c *Currencies, f string) error {
	cContent, err := os.ReadFile(f)
	if err != nil {
		return fmt.Errorf("your input file is cooked: %w", err)
	}

	cContent = []byte(strings.ReplaceAll(string(cContent), ",", "."))

	err = xml.Unmarshal(cContent, &c)
	if err != nil {
		return fmt.Errorf("your input file contents are cooked: %w", err)
	}

	return nil
}
