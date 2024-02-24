package web

// import (
//     "fmt"
//     "net/http"
//     "strings"

//     "github.com/gin-gonic/gin"

//     "iTools/crawler"
//     "iTools/linkedin"
//     "iTools/script"
// )

// func DownloadPost(c *gin.Context) {
//     // 讀取 url 值
//     url := c.PostForm("url")

//     // 只擷取網址部分
//     url = script.GrapURL(url)

//     var links []string

//     // 判斷網址
//     if script.IsWebSite(url, `threads`) || script.IsWebSite(url, `instagram`) {

//         // 開始爬
//         text := crawler.GetDoc(url)

//         // 獲取影片網址
//         lks := script.GetThreadsLink(text)
//         links = append(links, lks...)

//     } else if script.IsWebSite(url, `linkedin`) {

//         // 開始爬
//         doc := crawler.GetURLSoup(url)

//         link := script.GetLinkedInLink(doc)
//         links = append(links, link)

//         // 其他相關貼文
//         urls := script.GetLinkedInRelate(doc)
//         lks := linkedin.JobQueene(urls)

//         links = append(links, lks...)

//     } else if script.IsFacebook(url) {

//         url := crawler.GetRedirectURL(url)
//         url = crawler.GetRedirectURL(url)

//         url = script.GetFBVideoURL(url)

//         // 先替換網址
//         url = script.ReplaceWWW(url)

//         // 開始爬
//         doc := crawler.GetURLSoup(url)

//         lks := script.GetFBLink(doc)

//         links = append(links, lks...)

//     }

//     if len(links) == 0 {
//         setMeg(c, "alert", "請輸入 Facebook / threads / LinkedIn 網址")
//     }

//     var returnUrl = `/download`

//     // 把網址改base64合併回傳
//     returnUrl += fmt.Sprintf(`?url=%s`, url)
//     returnUrl += fmt.Sprintf(`&downloadLink=%s`, enBase64AndCombLink(links))

//     c.Redirect(http.StatusMovedPermanently, returnUrl)
// }

// func enBase64AndCombLink(links []string) string {
//     for idx := range links {
//         links[idx] = script.EnBase64(links[idx])
//     }
//     return strings.Join(links, "_")
// }
