package web

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"

    "iTools/conf"
    "iTools/script"
)

func PasswdPost(c *gin.Context) {
    chrTable := c.PostForm("chrTable")
    if chrTable == "" {
        chrTable = conf.ChrTable
    }

    countStr := c.PostForm("count")
    count := script.To_i(countStr)
    if count <= 0 {
        count = 10
    }

    passwd := script.RandList(chrTable, count)

    returnUrl := fmt.Sprintf(`/passwd?passwd=%s&chrTable=%s&count=%v`, passwd, chrTable, count)
    c.Redirect(http.StatusMovedPermanently, returnUrl)
}
