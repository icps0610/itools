package web

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"

    "iTools/conf"
)

func PasswdGet(c *gin.Context) {
    chrTable := c.Query("chrTable")
    if chrTable == "" {
        chrTable = conf.ChrTable
    }

    count := c.Query("count")
    if count == "" {
        count = "10"
    }

    passwd := c.Query("passwd")

    c.HTML(http.StatusOK, `passwd.html`, gin.H{
        "active":  c.Request.URL.Path,
        "msg":     getMsg(c),
        "bgColor": getBgColor(c),

        "chrTable": chrTable,
        "count":    count,
        "passwd":   passwd,
    })
}

var _ = fmt.Println
