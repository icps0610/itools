package urlProcess

import (
    "iTools/script"
)

func removeRegex(str, keyword string) string {
    if !script.Match(str, keyword) {
        return str
    }
    return removeUtmSource(script.Scan(str, keyword, 1))
}

func removeFBID(str string) string {
    keyword := `(.*)(\?|\&)fbclid=.*`
    return removeRegex(str, keyword)
}

func removeUtmSource(str string) string {
    keyword := `(.*)(\?|\&)utm_\w+\=.*`
    return removeRegex(str, keyword)
}

// func removeTail(str string) string {
//     keyword := `(.*)\?.*`
//     return removeRegex(str, keyword)
// }
