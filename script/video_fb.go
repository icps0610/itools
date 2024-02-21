package script

import (
    "fmt"
    "strings"

    "github.com/anaskhan96/soup"
)

func GetFacebookID(url string) string {
    groups := ScanGroups(url, `(\d+)_(\d+)_(\d+)_n.\w+`)
    var id string
    if len(groups) > 3 {
        ids := []string{zipNum(groups[1]), zipNum(groups[2]), zipNum(groups[3])}
        id = strings.Join(ids, "_")
    }

    return id
}

func GetFBLink(doc soup.Root) []string {
    var links []string
    for _, meta := range doc.FindAll("a") {
        // link := meta.Attrs()[`content`]
        link := meta.Attrs()[`href`]
        link = DeUrlCode(link)

        if Match(link, `fbcdn`) {
            if link != "" && notInclude(links, link) {
                link = strings.Replace(link, `/video_redirect/?src=`, "", -1)

                links = append(links, link)
                break
            }
        }
    }

    return links
}

func ReplaceWWW(url string) string {
    return strings.Replace(url, `www`, `mbasic`, 1)
}

var _ = fmt.Println
