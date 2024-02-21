package web

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func CarryGet(c *gin.Context) {

    num1 := initQuery(c, "num1", "0")
    num2 := initQuery(c, "num2", "0")

    carry1 := initQuery(c, "carry1", "10")
    carry2 := initQuery(c, "carry2", "10")

    c.HTML(http.StatusOK, `carry.html`, gin.H{
        "active":  c.Request.URL.Path,
        "msg":     getMsg(c),
        "bgColor": getBgColor(c),

        "num1":   num1,
        "num2":   num2,
        "carry1": carry1,
        "carry2": carry2,
    })
}

func initQuery(c *gin.Context, str, init string) string {
    query := c.Query(str)
    if query == "" {
        query = init
    }
    return query
}

var _ = fmt.Println
