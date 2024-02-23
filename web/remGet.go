package web

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"

    "iTools/conf"
    "iTools/script"
)

func RemGet(c *gin.Context) {
    var text string
    id := getRemText(c)
    if id == "" {
        id = script.RandList(conf.ChrTable, 10)
    }

    remPath := fmt.Sprintf(`%srem_%s`, TempDirPath, id)
    text = script.ReadFile(remPath)

    c.HTML(http.StatusOK, `rem.html`, gin.H{
        "active":  c.Request.URL.Path,
        "msg":     getMsg(c),
        "bgColor": getBgColor(c),

        "id":   id,
        "text": text,
    })
}

var _ = fmt.Println
