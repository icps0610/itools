package web

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"

    "iTools/conf"
    "iTools/script"
)

func UrlGet(c *gin.Context) {
    // 讀取參數和解base64
    url := queryAndDeBase64(c, "url")

    downloadLink := c.Query("downloadLink")
    downloadLink = strings.Replace(downloadLink, " ", "+", -1)

    var linkDatas []conf.LinkData
    for _, link := range strings.Split(downloadLink, "_") {
        if link != "" {
            // 解開
            link = script.DeBase64(link)

            var id, typeStr string
            if script.Match(link, `licdn`) {
                id = script.GetLinkedInID(link)
                typeStr = `mp4`

            } else if script.Match(link, `cdninstagram`) {
                id = script.GetThreadsID(link)
                typeStr = script.CheckType(link)

            } else if script.Match(link, `fbcdn`) {
                id = script.GetFacebookID(link)
                typeStr = script.CheckType(link)

            }

            linkData := conf.LinkData{id, typeStr, link}
            linkDatas = append(linkDatas, linkData)
        }
    }

    c.HTML(http.StatusOK, `url.html`, gin.H{
        "active":  c.Request.URL.Path,
        "msg":     getMsg(c),
        "bgColor": getBgColor(c),

        "url":       url,
        "linkDatas": linkDatas,
    })
}

var _ = fmt.Println
