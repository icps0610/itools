package web

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"

    "iTools/script"
    "iTools/url"
)

func UrlPost(c *gin.Context) {
    // 輸入的 Unicode 轉義序列

    urlstr := c.PostForm("url")
    urlstr = script.DeURLCode(urlstr)

    // 只擷取網址部分
    urlstr = script.GrapURL(urlstr)

    var msg string
    if urlstr == "" || !script.IsURL(urlstr) {
        setMeg(c, "alert", "請輸入網址")
    } else {
        u, title := url.Run(urlstr)

        msg = script.EnBase64(title + "\n" + u)
    }

    returnUrl := fmt.Sprintf(`/url?link=%s`, msg)
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
