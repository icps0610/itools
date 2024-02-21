package web

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func UrlGet(c *gin.Context) {
    // 讀取參數和解base64
    link := queryAndDeBase64(c, "link")

    c.HTML(http.StatusOK, `url.html`, gin.H{
        "active":  c.Request.URL.Path,
        "msg":     getMsg(c),
        "bgColor": getBgColor(c),

        "link": link,
    })
}

var _ = fmt.Println
