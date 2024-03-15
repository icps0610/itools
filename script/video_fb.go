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
        link = DeURLCode(link)

        if IsFacebookServer(link) {
            if link != "" && NotInclude(links, link) {
                link = strings.Replace(link, `/video_redirect/?src=`, "", -1)

                links = append(links, link)
                break
            }
        }

    }

    // 因為 檢查轉址 https://fb.watch/\w+/ => https://www.facebook.com/watch?v=\d+ 所以作廢
    // 又發現可以抓到 https://www.facebook.com/stories/\d+/\w+/

    for idx, meta := range doc.FindAll("script") {
        if idx > 58 {
            text := meta.Text()
            text = FixURL(text)

            for _, textSplit := range strings.Split(text, `",`) {
                for _, link := range ScanGroups(textSplit, `(https:\/\/.*mp4\?.*&oe=\w+)`) {

                    if NotInclude(links, link) {
                        links = append(links, link)
                    }
                }

                // for _, link := range ScanGroups(textSplit, `(https:\/\/.*jpg\?.*&oe=\w+)`) {

                //     if NotInclude(links, link) {

                //         links = append(links, link)

                //     }
                // }
            }

        }

    }

    return links
}

func ReplaceWWW(url string) string {
    return strings.Replace(url, `www`, `mbasic`, 1)
}

func GetFBVideoURL(url string) string {
    keywords := []string{
        `watch\/?.*v=(\d+)`,
        `reel\/(\d+)`,
        `.*\/videos\/(\d+)`,
    }

    for _, keyword := range keywords {
        keyword = `www.facebook.com\/` + keyword

        id := Scan(url, keyword, 1)

        if id != "" {
            return fmt.Sprintf(`https://www.facebook.com/watch?v=%s`, id)
        }
    }

    return url
}
