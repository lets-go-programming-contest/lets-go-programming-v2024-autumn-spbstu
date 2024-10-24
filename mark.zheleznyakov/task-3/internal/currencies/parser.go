package currencies

import (
	"encoding/xml"
	"fmt"
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

	err = xml.Unmarshal(cContent, &c)
	if err != nil {
		return fmt.Errorf("your input file contents are cooked: %w", err)
	}

	return nil
}
