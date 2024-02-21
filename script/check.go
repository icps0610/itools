package script

import (
    "fmt"
)

func IsURL(str string) bool {
    return Match(str, `^(http:\/\/|https:\/\/)`)
}

func IsWebSite(str, webSite string) bool {
    return IsURL(str) && Match(str, webSite)
}

func IsYoutube(url string) bool {
    return Match(url, `youtube`) || Match(url, `youtu\.be`)
}

func IsFacebook(url string) bool {
    return Match(url, `facebook`)
}

func isJPG(link string) bool {
    return Match(link, `jpg`)
}

func isMP4(link string) bool {
    return Match(link, `mp4`)
}

func CheckType(link string) string {
    if isMP4(link) {
        return `mp4`
    } else if isJPG(link) {
        return `jpg`
    }
    return `unknown`
}

func GrapURL(link string) string {
    return Scan(link, `(https?:\/\/.*)`, 1)
}

var _ = fmt.Println
