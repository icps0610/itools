package conf

import (
    "os"
)

var (
    Mode     = getMode()
    Cookie   = os.Getenv("cookie")
    ChrTable = `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`
)

func getMode() string {
    var mode string
    args := os.Args
    if len(args) > 1 {
        mode = args[1]
    }
    return mode
}
