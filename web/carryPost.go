package web

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"

    "iTools/carry62"
    "iTools/script"
)

func CarryPost(c *gin.Context) {
    num1 := c.PostForm("num1")
    carry1 := c.PostForm("carry1")
    carry2 := c.PostForm("carry2")

    carry1_int := script.To_i(carry1)
    carry2_int := script.To_i(carry2)

    var num2 string
    if carry1_int > 62 || carry2_int > 62 || carry1_int < 2 || carry2_int < 2 {
        setMeg(c, "alert", `進位 2-62`)

    } else if checkNumCarry(num1, carry1_int) == true {
        value := carry62.Decimal(num1, carry1_int)
        num2 = carry62.Carry(value, carry2_int)

    } else {
        msg := fmt.Sprintf(`數字:%s　進位:%s 錯誤`, num1, carry1)
        setMeg(c, "alert", msg)
    }

    returnUrl := fmt.Sprintf(`/carry?num1=%s&carry1=%s&num2=%v&carry2=%s`, num1, carry1, num2, carry2)
    c.Redirect(http.StatusMovedPermanently, returnUrl)
}

func checkNumCarry(num1 string, carry1_int int) bool {
    for _, num := range strings.Split(num1, "") {
        if carry62.Decimal(num, carry1_int) >= carry1_int {
            return false
        }
    }
    return true
}

var _ = fmt.Println
