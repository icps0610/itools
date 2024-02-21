package web

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"

    "iTools/crawler"
    "iTools/linkedin"
    "iTools/script"
)

func DownloadPost(c *gin.Context) {
    // 讀取 url 值
    url := c.PostForm("url")

    // 只擷取網址部分
    url = script.GrapURL(url)

    var links []string

    // 判斷網址
    if script.IsWebSite(url, `threads`) || script.IsWebSite(url, `instagram`) {

        // 開始爬
        text := crawler.GetDoc(url)

        // 獲取影片網址
        lks := script.GetThreadsLink(text)
        links = append(links, lks...)

    } else if script.IsWebSite(url, `linkedin`) {

        // 開始爬
        doc := crawler.GetDocSoup(url)

        link := script.GetLinkedInLink(doc)
        links = append(links, link)

        // 其他相關貼文
        urls := script.GetLinkedInRelate(doc)
        lks := linkedin.JobQueene(urls)

        links = append(links, lks...)

    } else if script.IsWebSite(url, `facebook`) {

        // reel
        id := script.Scan(url, `https:\/\/www.facebook.com\/reel\/(\d+)`, 1)
        if id != "" {
            url = fmt.Sprintf(`https://www.facebook.com/watch?v=%s`, id)
        }

        // 先替換網址
        url = script.ReplaceWWW(url)

        // 開始爬
        doc := crawler.GetDocSoup(url)

        lks := script.GetFBLink(doc)

        links = append(links, lks...)
    }

    if len(links) == 0 {
        setMeg(c, "alert", "請輸入 Facebook / threads / LinkedIn 網址")
    }

    var returnUrl = `/download`

    // 把網址改base64合併回傳
    returnUrl = enBase64AndCombLink(links)
    c.Redirect(http.StatusMovedPermanently, returnUrl)
}

func enBase64AndCombLink(links []string) string {
    for idx := range links {
        links[idx] = script.EnBase64(links[idx])
    }
    combLinks := strings.Join(links, "_")
    return fmt.Sprintf(`/download?link=%s`, combLinks)
}