package valuteProcessing

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/hahapathetic/task-3/internal/valuteStructs"
	"golang.org/x/text/encoding/charmap"
	"io"
	"os"
	"sort"
	"strings"
)

func ParseFromXML(valuteRates *valuteStructs.ValuteRate, filePath string) error {
	sort.Slice(valuteRates.ValuteRate, func(i, j int) bool {
		return valuteRates.ValuteRate[i].Value > valuteRates.ValuteRate[j].Value
	})

	configFile, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	configFile = []byte(strings.ReplaceAll(string(configFile), ",", "."))

	decoder := xml.NewDecoder(bytes.NewReader(configFile))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unsupported charset: %s", charset)
	}

	err = decoder.Decode(&valuteRates)
	if err != nil {
		return fmt.Errorf("failed to decode: %w", err)
	}

	return nil
}
