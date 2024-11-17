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

    const (
        DirPerm = 0755
        FilePerm = 0644
    )
    directory := filepath.Dir(FilePath)
    err = os.MkdirAll(directory, DirPerm)
    if err != nil {
        return 
    }
    file, err := os.OpenFile(FilePath, os.O_CREATE|os.O_RDWR, FilePerm)
    if err != nil {
        return
    }
    defer file.Close()
    os.WriteFile(file.Name(), Result, FilePerm)
    if err != nil {
        return
    }
    return 
}
