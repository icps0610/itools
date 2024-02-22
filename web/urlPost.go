package web

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"

    "iTools/script"
    "iTools/url"
)

func UrlPost(c *gin.Context) {
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
