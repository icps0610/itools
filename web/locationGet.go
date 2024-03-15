package web

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func LocationGet(c *gin.Context) {

    c.HTML(http.StatusOK, `location.html`, gin.H{
        "active": c.Request.URL.Path,
    })
}

var _ = fmt.Println
