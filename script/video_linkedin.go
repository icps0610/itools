package script

import (
    "fmt"
    "strings"

    "github.com/anaskhan96/soup"
)

func GetLinkedInID(link string) string {
    keyword := fmt.Sprintf(`\/.*/(\d+)\?`)

    return Scan(link, keyword, 1)
}

func GetLinkedInLink(doc soup.Root) string {
    var result string
    sc := doc.FindAll("script")

    if len(sc) > 0 {
        scriptTexts := sc[1].Text()
        scriptTexts = DeURLCode(scriptTexts)

        for _, scriptText := range strings.Split(scriptTexts, `,`) {
            keyword := `(https:\/\/dms.licdn.com\/playlist\/vid\/.*)"`
            result = Scan(scriptText, keyword, 1)

            if result != "" {
                break
            }
        }
    }

    return result
}

func GetLinkedInRelate(doc soup.Root) []string {
    var urls []string
    for _, div := range doc.FindAll("div", "class", "link-overlay") {
        u := div.Find("a").Attrs()["href"]
        urls = append(urls, u)
    }
    return urls
}
