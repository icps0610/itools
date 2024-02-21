package web

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func Base64Get(c *gin.Context) {
    // 讀取參數
    enBase64 := queryAndDeBase64(c, "enBase64")

    deBase64 := c.Query("deBase64")

    c.HTML(http.StatusOK, `base64.html`, gin.H{
        "active":  c.Request.URL.Path,
        "msg":     getMsg(c),
        "bgColor": getBgColor(c),

        "enBase64": enBase64,
        "deBase64": deBase64,
    })
}

var _ = fmt.Println
