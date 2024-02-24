package web

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func IndexGet(c *gin.Context) {
    c.Redirect(http.StatusMovedPermanently, `/url`)
}
