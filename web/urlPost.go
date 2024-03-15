package web

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "strings"

    // "iTools/conf"
    "iTools/crawler"
    "iTools/linkedin"
    "iTools/script"
    // "iTools/simpToTrad"
    "iTools/urlProcess"
)

func UrlPost(c *gin.Context) {

    url := c.PostForm("url")
    // 只擷取網址部分
    url = script.GrapURL(url)

    var returnUrl = `/url?`
    var docText string

    if url == "" {
        setMeg(c, "alert", "請輸入網址")
    } else {
        url = urlProcess.ProcessURL(url)

        // 開始爬
        var title string
        title, docText = urlProcess.GetTitle(url)
        // title = simpToTrad.Run(title)

        msg := script.EnBase64(title + "\n" + url)
        returnUrl += fmt.Sprintf(`url=%s`, msg)
    }

    // download

    var links []string

    // 判斷網址
    if script.IsWebSite(url, `threads`) || script.IsWebSite(url, `instagram`) {

        // 開始爬
        // docText = crawler.GetDoc(url)

        // 獲取影片網址
        lks := script.GetThreadsLink(docText)
        links = append(links, lks...)

    } else if script.IsWebSite(url, `linkedin`) {

        // 轉換 doc
        doc := crawler.GetTextSoup(docText)

        link := script.GetLinkedInLink(doc)
        links = append(links, link)

        // 其他相關貼文
        urls := script.GetLinkedInRelate(doc)
        lks := linkedin.JobQueene(urls)

        links = append(links, lks...)

    } else if script.IsFacebook(url) {

        // 轉換 doc
        doc := crawler.GetTextSoup(docText)

        lks := script.GetFBLink(doc)

        links = append(links, lks...)

    }

    // if conf.Mode == `pi` && script.IsYoutube(url) {
    //     cmd := fmt.Sprintf(`echo 'chdir /d z:\' > /tmp/u2.bat`)
    //     script.RunCmd(cmd)

    //     cmd = fmt.Sprintf(`echo yt-dlp -f mp4 %s >> /tmp/u2.bat`, url)
    //     script.RunCmd(cmd)

    //     cmd = fmt.Sprintf(`echo yt-dlp -f m4a %s >> /tmp/u2.bat`, url)
    //     script.RunCmd(cmd)

    //     cmd = fmt.Sprintf(`chmod 777 /tmp/u2.bat`)
    //     script.RunCmd(cmd)

    // setMeg(c, "success", fmt.Sprintf(`已將 %s 列入排程`, url))
    // }

    // if len(links) == 0 {
    //     setMeg(c, "alert", "請輸入 Facebook / threads / LinkedIn 網址")
    // }

    if len(links) > 0 {
        // 把網址改base64合併回傳
        returnUrl += fmt.Sprintf(`&downloadLink=%s`, enBase64AndCombLink(links))
    }

    c.Redirect(http.StatusMovedPermanently, returnUrl)
}

func parseUnicodeEscape(input string) []uint16 {
    // 移除開頭的 "\u"
    input = input[2:]

    // 分割字串為單個 Unicode 轉義序列
    escapeSequences := splitString(input, "\\u")

    // 解析每個 Unicode 轉義序列並轉換成 Unicode 碼點
    var codePoints []uint16
    for _, esc := range escapeSequences {
        // 將十六進制轉換為整數
        value, err := strconv.ParseUint(esc, 16, 32)
        if err == nil {
            codePoints = append(codePoints, uint16(value))
        }
    }
    return codePoints
}

func splitString(s, sep string) []string {
    var parts []string
    for len(s) > 0 {
        i := 0
        // 找到分隔符號的位置
        if i = len(s); i == 0 {
            break
        }
        // 添加子字串到切片中
        parts = append(parts, s[:i])
        // 更新剩餘的字串
        s = s[i:]
    }
    return parts
}

func enBase64AndCombLink(links []string) string {
    for idx := range links {
        links[idx] = script.EnBase64(links[idx])
    }
    return strings.Join(links, "_")
}
