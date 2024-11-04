package currency

import (
	"fmt"
	"io"
	"net/http"
	"os"
	fp "path/filepath"
)

func DownloadCurrencyData(filepath string) error {
	if _, err := os.Stat(filepath); err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("failed to check file existence: %w", err)
	}

	url := "https://www.cbr.ru/scripts/XML_daily.asp"

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error downloading data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download data, status: %s", resp.Status)
	}

	if err := os.MkdirAll(fp.Dir(filepath), os.ModePerm); err != nil {
		return fmt.Errorf("error creating directories: %w", err)
	}

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("error writing data to file: %w", err)
	}

	return nil
}
