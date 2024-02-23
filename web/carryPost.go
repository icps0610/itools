package web

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"

    "iTools/carry64"
    "iTools/script"
)

func CarryPost(c *gin.Context) {
    num1 := c.PostForm("num1")
    carry1 := c.PostForm("carry1")
    carry2 := c.PostForm("carry2")

    carry1_int := script.To_i(carry1)
    carry2_int := script.To_i(carry2)

    var num2 string
    if carry1_int > 64 || carry2_int > 64 || carry1_int < 2 || carry2_int < 2 {
        setMeg(c, "alert", `進位 2-64`)

    } else if checkNumCarry(num1, carry1_int) == true {
        value := carry64.Decimal(num1, carry1_int)
        num2 = carry64.Carry(value, carry2_int)

    } else {
        msg := fmt.Sprintf(`數字:%s　進位:%s 錯誤`, num1, carry1)
        setMeg(c, "alert", msg)
    }

    returnUrl := fmt.Sprintf(`/carry?carry1=%s&num1=%s&carry2=%v&num2=%s`, carry1, num1, carry2, num2)
    c.Redirect(http.StatusMovedPermanently, returnUrl)
}

func checkNumCarry(num1 string, carry1_int int) bool {
    for _, num := range strings.Split(num1, "") {
        if carry64.Decimal(num, carry1_int) >= carry1_int {
            return false
        }
    }
    return true
}

var _ = fmt.Println
