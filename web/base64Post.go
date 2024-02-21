package web

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"

    "iTools/script"
)

func EnBase64Post(c *gin.Context) {
    text := c.PostForm("text")
    text = script.EnBase64(text)

    returnUrl := fmt.Sprintf(`/base64?deBase64=%s`, text)
    c.Redirect(http.StatusMovedPermanently, returnUrl)
}

func DeBase64Post(c *gin.Context) {
    text := c.PostForm("text")

    returnUrl := fmt.Sprintf(`/base64?enBase64=%s`, text)
    c.Redirect(http.StatusMovedPermanently, returnUrl)
}

var _ = fmt.Println
