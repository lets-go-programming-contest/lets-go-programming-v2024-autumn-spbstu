package userErrors

import "errors"

var (
	ErrConfigFileIsNotExist   = errors.New("config file is not exist")
	ErrReadingFile            = errors.New("reading file failed")
	ErrDeserializationFailure = errors.New("deserialization failed")
	ErrInputFileIsNotExist    = errors.New("input file is not exist")
	ErrMkdirFailure           = errors.New("can't create directory")
	ErrCreatingFileFailure    = errors.New("can't create file")
	ErrWriteFile              = errors.New("writing file failed")
	ErrSerializationFailure   = errors.New("serialization failed")
	ErrEncodingConversion     = errors.New("can't encode")
)
