package conf

import (
    "os"
)

var (
    Mode   = getMode()
    Cookie = os.Getenv("cookie")
)

func getMode() string {
    var mode string
    args := os.Args
    if len(args) > 1 {
        mode = args[1]
    }
    return mode
}
