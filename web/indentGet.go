package web

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func IndentGet(c *gin.Context) {
    // 讀取參數和解base64
    indent := queryAndDeBase64(c, "indent")

    c.HTML(http.StatusOK, `indent.html`, gin.H{
        "active":  c.Request.URL.Path,
        "msg":     getMsg(c),
        "bgColor": getBgColor(c),

        "indent": indent,
    })
}

var _ = fmt.Println
