package script

import (
    "fmt"
    "strings"
)

func GetThreadsID(url string) string {
    groups := ScanGroups(url, `https:\/\/.*\/(.+)_?n?\.[jpg|mp4]`)

    var id string
    if len(groups) > 1 {
        id = strings.Join(groups[1:], "_")
    }

    return id
}

func GetThreadsLink(texts string) []string {
    // 修正 /
    texts = strings.Replace(texts, `\/`, `/`, -1)
    texts = strings.Replace(texts, `\u0025`, `%`, -1)

    var splitText string
    // threads
    if Match(texts, `thread_items`) {
        splitText = splitKeyWordText(texts, `thread_items`, 1)

    } else if Match(texts, `Instagram`) {
        // Instagram
        splitText = splitKeyWordText(texts, `ScheduledServerJS`, 5)
    }

    var links []string
    // 用 , 分段
    for _, text := range strings.Split(splitText, ",") {
        // 去除有解析度(通常是profile_pic_url) stp=dst-jpg 沒被截圖
        if (isJPG(text) && !Match(text, `_\w{1}\d+x\d+`) && Match(text, `stp=dst-jpg`)) || isMP4(text) {
            keyWord := `url":"(https:\/\/scontent.cdninstagram.com.*10d13b)`
            link := Scan(text, keyWord, 1)

            if link != "" && notInclude(links, link) {

                links = append(links, link)
            }
        }
    }

    return links
}

func splitKeyWordText(texts, keyWord string, idxK int) string {
    for idx, text := range strings.Split(texts, keyWord) {
        if idx == idxK {
            return text
        }
    }
    return ""
}

func notInclude(arr []string, e1 string) bool {
    for _, e := range arr {
        if e == e1 {
            return false
        }
    }
    return true
}

var _ = fmt.Println
