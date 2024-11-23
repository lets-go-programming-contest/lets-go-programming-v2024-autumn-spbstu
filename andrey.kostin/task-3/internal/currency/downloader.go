package currency

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	errProc "github.com/IDevFrye/task-3/internal/errors"
)

func DownloadCurrencyData(filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		return nil
	}

	url := "https://www.cbr.ru/scripts/XML_daily.asp?date_req=02/03/2002"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("%w: %s", errProc.ErrDataDownload, err.Error())
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("%w: %s", errProc.ErrDataDownload, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w, status: %s", errProc.ErrDataDownload, resp.Status)
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return fmt.Errorf("%w: %s", errProc.ErrDirectoryCreation, err.Error())
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("%w: %s", errProc.ErrFileCreation, err.Error())
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("%w: %s", errProc.ErrDataWriting, err.Error())
	}

	fmt.Println("Currency data downloaded successfully")
	return nil
}
