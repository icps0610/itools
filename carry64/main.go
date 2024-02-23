package carry64

func Carry(num, gate int) string {
    var ans string
    for _, c := range ToCarry(gate, num, []int{}) {
        ans += CHR[c]
    }
    return ans
}

func Decimal(str string, gate int) int {
    return toDecimal(gate, str)
}
