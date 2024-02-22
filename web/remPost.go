package web

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"

    "iTools/script"
)

func RemPost(c *gin.Context) {
    text := c.PostForm("text")
    text = script.DeURLCode(text)

    setRemText(c, text)

    c.Redirect(http.StatusMovedPermanently, `/rem`)
}

func RemClear(c *gin.Context) {
    setRemClear(c)

    c.Redirect(http.StatusMovedPermanently, `/rem`)
}

func RemRmLast(c *gin.Context) {
    text := getRemText(c)

    texts := script.SplitText(text)
    texts = removeLastEmpty(texts)

    text = strings.Join(texts, "\n")

    setRemText(c, text)

    c.Redirect(http.StatusMovedPermanently, `/rem`)
}

func removeLastEmpty(texts []string) []string {
    var lastIdx = len(texts) - 1

    if lastIdx > 1 && strings.TrimSpace(texts[lastIdx]) == "" {
        return texts[:lastIdx-1]
    }

    return texts[:lastIdx]
}

func setRemText(c *gin.Context, msg string) {
    setSession(c, "rem", msg)
}

func setRemClear(c *gin.Context) {
    setRemText(c, "")
}

func getRemText(c *gin.Context) string {
    return getSession(c, "rem")
}

var _ = fmt.Println
