package currency

import "errors"

var (
	ErrConfigNotFound     = errors.New("configuration file not found or unreadable")
	ErrInputFileNotFound  = errors.New("input file not found")
	ErrInvalidXMLFormat   = errors.New("failed to decode XML: incorrect format")
	ErrDirectoryCreation  = errors.New("error creating directories")
	ErrFileCreation       = errors.New("error creating file")
	ErrDataWriting        = errors.New("error writing data to file")
	ErrEmptyInputFile     = errors.New("input file is empty")
	ErrUnsupportedCharset = errors.New("unsupported charset in XML data")
	ErrDataDownload       = errors.New("failed to download data")
	ErrJSONMarshal        = errors.New("failed to marshal data to JSON")
)
