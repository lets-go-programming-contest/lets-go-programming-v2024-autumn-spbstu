package logic

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	strct "github.com/Mmmakskl/task-3/internal/structures"
)

func Parser(filePath string, conf *strct.ValCurs) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Failure open file: %w", err)
	}

	file = []byte(strings.ReplaceAll(string(file), ",", "."))

	if err = xml.Unmarshal(file, conf); err != nil {
		return fmt.Errorf("Failure decoding file: %w", err)
	}

	return nil
}
