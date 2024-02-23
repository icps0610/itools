package carry64

import (
    "strings"
)

var chrTable = strings.Split(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_`, ``)

func ToCarry(gate, quotient int, arr []int) []int {
    arr = append(arr, quotient%gate)
    quotient = quotient / gate
    if quotient >= gate {
        return ToCarry(gate, quotient, arr)
    }
    arr = append(arr, quotient)
    arr = reverse(arr)

    if arr[0] == 0 {
        arr = arr[1:]
    }

    return arr
}

func toDecimal(gate int, strs string) int {
    var sum int
    for idx, str := range strs {
        int := getIndex(str)
        value := power(gate, len(strs)-1-idx)
        sum += int * value
    }
    return sum
}

func reverse(arr []int) []int {
    var nar []int
    for i := len(arr) - 1; i >= 0; i-- {
        nar = append(nar, arr[i])
    }
    return nar
}

func getIndex(r rune) int {
    for i, e := range chrTable {
        if e == string(r) {
            return i
        }
    }
    return -1
}

func power(gate, idx int) int {
    if idx <= 0 {
        return 1
    }
    return gate * power(gate, idx-1)
}
