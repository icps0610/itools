package web

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"

    "iTools/script"
)

func RemPost(c *gin.Context) {
    var returnURL = `/rem`
    text := c.PostForm("text")
    text = script.DeURLCode(text)

    id := getRemText(c)
    if id != "" {
        remPath := fmt.Sprintf(`%srem_%s`, TempDirPath, id)
        script.WriteFile(text, remPath)
    }

    c.Redirect(http.StatusMovedPermanently, returnURL)
}

func RemIDPost(c *gin.Context) {
    var returnURL = `/rem`

    id := c.PostForm("id")
    setRemText(c, id)

    c.Redirect(http.StatusMovedPermanently, returnURL)
}

func setRemText(c *gin.Context, msg string) {
    setSession(c, "rem", msg)
}

func getRemText(c *gin.Context) string {
    return getSession(c, "rem")
}

func setRemClear(c *gin.Context) {
    setRemText(c, "")
}

var _ = fmt.Println
