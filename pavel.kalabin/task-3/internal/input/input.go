package input

import (
    "os"
    "path/filepath"
)

func ReadXML(FilePath string) (contents []byte, err error) {
    contents, err = os.ReadFile(FilePath)
    return
}

func WriteResult(FilePath string, Result []byte) (err error) {

    directory := filepath.Dir(FilePath)
    err = os.MkdirAll(directory, os.ModePerm)
    if err != nil {
        return 
    }
    file, err := os.OpenFile(FilePath, os.O_CREATE|os.O_RDWR, 0644)
    if err != nil {
        return
    }
    defer file.Close()
    os.WriteFile(file.Name(), Result, os.ModePerm)
    if err != nil {
        return
    }
    return 
}
