package web

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func TimeGet(c *gin.Context) {
    utime := c.Query("utime")
    timeFormat := c.Query("timeformat")

    c.HTML(http.StatusOK, `time.html`, gin.H{
        "active":  c.Request.URL.Path,
        "msg":     getMsg(c),
        "bgColor": getBgColor(c),

        "timeFormat": timeFormat,
        "utime":      utime,
    })
}

var _ = fmt.Println
