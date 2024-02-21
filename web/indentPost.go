package web

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"

    "iTools/script"
)

func IndentPost(c *gin.Context) {
    indent := c.PostForm("indent")

    if indent != "" {
        indent = strings.Replace(indent, "\r\n", "", -1)
        indent = strings.Replace(indent, "\n", "", -1)
        indent = strings.TrimSpace(indent)

        indent = fmt.Sprintf(`?indent=%s`, script.EnBase64(indent))
    }

    returnUrl := fmt.Sprintf(`/indent%s`, indent)
    c.Redirect(http.StatusMovedPermanently, returnUrl)
}

var _ = fmt.Println
