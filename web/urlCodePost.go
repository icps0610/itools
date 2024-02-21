package web

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"

    "iTools/script"
)

func EnUrlPost(c *gin.Context) {
    text := c.PostForm("text")
    text = script.EnUrlCode(text)
    text = script.EnBase64(text)

    returnUrl := fmt.Sprintf(`/urlCode?deUrlCode=%s`, text)
    c.Redirect(http.StatusMovedPermanently, returnUrl)
}

func DeUrlPost(c *gin.Context) {
    text := c.PostForm("text")
    text = script.DeUrlCode(text)
    text = script.EnBase64(text)

    returnUrl := fmt.Sprintf(`/urlCode?enUrlCode=%s`, text)
    c.Redirect(http.StatusMovedPermanently, returnUrl)
}

var _ = fmt.Println
