package web

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"

    "iTools/script"
)

func CookiePost(c *gin.Context) {
    cookie := c.PostForm("text")

    var resultLine string
    var result string
    for _, line := range strings.Split(cookie, "\r\n") {
        split := strings.Fields(line)
        if len(split) > 1 {
            result += fmt.Sprintf(`document.cookie = "%s=%s"`, split[0], split[1])
            result += "\r\n"

            resultLine += fmt.Sprintf(`%s=%s;`, split[0], split[1])
        }
    }
    if result != "" {
        result += "location.reload();\r\n\r\n\r\n"

        result += resultLine
    }

    result = script.EnBase64(result)

    returnUrl := fmt.Sprintf(`/cookie?cookie=%s`, result)
    c.Redirect(http.StatusMovedPermanently, returnUrl)
}

var _ = fmt.Println
