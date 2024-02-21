package script

import (
    "strings"
)

func FixURL(url string) string {
    // 修正 /
    url = strings.Replace(url, `\/`, `/`, -1)
    url = strings.Replace(url, `\u0025`, `%`, -1)

    return url
}
