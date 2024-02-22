package crawler

import (
    "fmt"
    "strings"

    "github.com/anaskhan96/soup"

    "iTools/script"
)

func GetUrlTitle(url string) string {
    var title, channel string

    // 若是FB的影片
    if script.IsFacebookVideo(url) {
        // 先替換網址
        url = script.ReplaceWWW(url)

        doc := GetDocSoup(url)

        // 依照貼文分隔
        divs := doc.FindAll("div", "class", "story_body_container")

        if len(divs) > 0 {
            div := divs[0]
            channel = getFindText(div, `a`, 0)
            title = getFindText(div, `p`, 0)
        }

        divs = doc.FindAll("div", "id", "m_story_permalink_view")

        if len(divs) > 0 {
            div := divs[0]
            channel = div.Find("strong").FullText()

            title = getFindText(div, `p`, 0)
        }
        return channel + " - " + title
    }

    doc := GetDocSoup(url)

    // 若是 Facebook
    if script.IsFacebook(url) {
        for _, meta := range doc.FindAll(`script`) {
            text := meta.FullText()

            channel := script.Scan(text, `"owner":{"__typename":"User","name":"(.*)","profile_picture"`, 1)
            channel = script.DeUnicode(channel)

            title := script.Scan(text, `,"text":"(.*)"\},"message_preferred_body":`, 1)
            title = script.DeUnicode(title)
            title = strings.Replace(title, "\n", ` `, -1)

            if channel != "" {
                return channel + " - " + title
            }
        }
    }

    // 若是 youtube
    if script.IsYoutube(url) {
        channel = getDocText(doc, "link", "itemprop", "name")
        title = getDocText(doc, "meta", "name", "title")

        return channel + " - " + title
    }

    title = getTitleText(doc)

    title = strings.TrimSpace(title)
    title = strings.Replace(title, "\n", "", -1)

    return title
}

func getDocText(doc soup.Root, tag, class, value string) string {
    root := doc.Find(tag, class, value)
    return root.Attrs()["content"]
}

func getTitleText(doc soup.Root) string {
    var title string
    ts := doc.FindAll("title")
    if len(ts) > 0 {
        title = ts[0].Text()
    }
    return title
}

func getFindText(doc soup.Root, ele string, idx int) string {
    e := doc.FindAll(ele)
    if len(e) > idx {
        return e[idx].FullText()
    }
    return ""
}

var _ = fmt.Println
