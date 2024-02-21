package web

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

func UrlCodeGet(c *gin.Context) {
    // 讀取參數
    enUrlCode := queryAndDeBase64(c, "enUrlCode")
    deUrlCode := queryAndDeBase64(c, "deUrlCode")

    // 換行
    deUrlCode = strings.Replace(deUrlCode, "%0D%0A", "\n", -1)

    c.HTML(http.StatusOK, `urlCode.html`, gin.H{
        "active":  c.Request.URL.Path,
        "msg":     getMsg(c),
        "bgColor": getBgColor(c),

        "enUrlCode": enUrlCode,
        "deUrlCode": deUrlCode,
    })
}

var _ = fmt.Println
