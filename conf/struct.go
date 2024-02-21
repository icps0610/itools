package conf

import (
    "os"
)

var (
    Mode = getMode()
)

func getMode() string {
    var mode string
    args := os.Args
    if len(args) > 1 {
        mode = args[1]
    }
    return mode
}

type LinkData struct {
    ID   string
    Type string
    Link string
}

type File struct {
    Name   string
    Base64 string
}
