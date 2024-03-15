package urlProcess

import (
    "fmt"
    "strings"

    "github.com/anaskhan96/soup"

    "iTools/crawler"
    "iTools/script"
)

func GetTitle(url string) (string, string) {

    var title, channel string

    // 若是FB的影片
    if script.IsFacebookVideo(url) {

        // 先替換網址
        url = script.ReplaceWWW(url)

        docText := crawler.GetDoc(url)
        doc := crawler.GetTextSoup(docText)
        // FB - https://www.facebook.com/reel/\d+
        // FB - https://www.facebook.com/watch?v=\d+

        divs := doc.FindAll("div", "class", "story_body_container")
        if len(divs) > 0 {
            div := divs[0]
            channel = getFindText(div, `a`, 0)
            title = getFindText(div, `p`, 0)

        }

        // https://www.facebook.com/\w+/videos/\d+

        if title == "" {
            divs = doc.FindAll("div", "id", "m_story_permalink_view")
            if len(divs) > 0 {
                div := divs[0]
                sg := div.FindAll("strong")
                if len(sg) > 0 {
                    channel = sg[0].FullText()
                }

                title = getFindText(div, `p`, 0)

            }
        }

        if title == "" {
            divs := doc.FindAll("title")
            if len(divs) > 0 {
                title = divs[0].Text()
            }
        }

        return tileAndChannel(title, channel), docText
    }

    docText := crawler.GetDoc(url)
    doc := crawler.GetTextSoup(docText)

    // 有些有
    // FB - https://www.facebook.com/photo/?fbid=\d+&set=pob.\d+

    if script.IsFacebook(url) {
        for _, meta := range doc.FindAll(`script`) {
            text := meta.FullText()

            channel := script.Scan(text, `"owner":\{"__typename":"User","name":"(.*)","profile_picture"`, 1)
            // channel := script.Scan(text, `"User","owner_as_page":\{"name":"(.*)","id"`, 1)

            title := script.Scan(text, `,"text":"(.*)"\},"message_preferred_body":`, 1)
            // title := script.Scan(text, `ranges":\[\],"color_ranges":\[\],"text":"(.*)"\},"feedback":\{`, 1)

            if channel != "" && title != "" {
                channel = script.DeUnicode(channel)

                title = strings.Replace(title, `\n`, ` `, -1)
                title = script.DeUnicode(title)

                title = tileAndChannel(title, channel)
            }
        }

    } else if script.IsYoutube(url) {

        // 若是 youtube

        channel = getDocText(doc, "link", "itemprop", "name")
        title = getDocText(doc, "meta", "name", "title")

        title = tileAndChannel(title, channel)

    } else if script.IsInstagram(url) {

        // 若是 instagram

        channel = getDocText(doc, "meta", "name", "twitter:title")
        channel = strings.Replace(channel, ` • Instagram reel`, ``, -1)

        title = getDocText(doc, "meta", "name", "description")
        title = script.Scan(title, `"(.*)"`, 1)

        title = tileAndChannel(title, channel)
    } else {

        // FB - https://www.facebook.com/\w+/posts/\d+?ref=embed_post
        // FB - https://www.facebook.com/permalink.php?story_fbid=\d+&id=\d+&ref=embed_post
        // FB - https://www.facebook.com/groups/\d+/permalink/\d+/?app=fbl
        // FB - https://www.facebook.com/share/v/\w+/?mibextid=oFDknk
        // FB - https://www.facebook.com/photo/?fbid=\d+&set=a.\d+

        title = getTitleText(doc)

        title = strings.TrimSpace(title)
        title = strings.Replace(title, "\n", "", -1)
    }

    return title, docText
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

func tileAndChannel(title, channel string) string {
    if channel != "" {
        channel = fmt.Sprintf(`[ %s ]`, channel)
    }
    return fmt.Sprintf(`%s %s`, title, channel)
}

var _ = fmt.Println
