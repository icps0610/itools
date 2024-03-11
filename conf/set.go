package conf

import (
    "os"
    // "time"
)

var (
    Mode     = getMode()
    Cookie   = os.Getenv("cookie")
    ChrTable = `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`

    // TimeNow = time.Now()
)

func getMode() string {
    var mode string
    args := os.Args
    if len(args) > 1 {
        mode = args[1]
    }
    return mode
}
