package web

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func RemGet(c *gin.Context) {
    text := getRemText(c)

    c.HTML(http.StatusOK, `rem.html`, gin.H{
        "active":  c.Request.URL.Path,
        "msg":     getMsg(c),
        "bgColor": getBgColor(c),

        "text": text,
    })
}

var _ = fmt.Println
