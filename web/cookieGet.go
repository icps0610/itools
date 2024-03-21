package web

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func CookieGet(c *gin.Context) {
    // 讀取參數和解base64
    cookie := queryAndDeBase64(c, "cookie")

    c.HTML(http.StatusOK, `cookie.html`, gin.H{
        "active":  c.Request.URL.Path,
        "msg":     getMsg(c),
        "bgColor": getBgColor(c),

        "cookie": cookie,
    })
}

var _ = fmt.Println
