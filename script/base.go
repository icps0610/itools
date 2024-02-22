package script

import (
    "encoding/base64"
    "fmt"
    "math/rand"
    "net/url"
    "os/exec"
    "runtime"
    "strconv"
    "strings"
    "time"

    "iTools/carry62"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func RandNumInt(min, max int) int {
    return min + rand.Intn(max-min+1)
}

func RunCmd(cmd string) {
    var bash, args = `bash`, `-c`
    if runtime.GOOS == "windows" {
        bash, args = `cmd`, `/c`
    }
    exec.Command(bash, args, cmd).CombinedOutput()
}

func EnBase64(str string) string {
    return base64.StdEncoding.EncodeToString([]byte(str))
}

func DeBase64(str string) string {
    s, _ := base64.StdEncoding.DecodeString(str)
    return string(s)
}

func EnURLCode(str string) string {
    return url.QueryEscape(str)
}

func DeURLCode(str string) string {
    decodedString, _ := url.QueryUnescape(str)
    return decodedString
}

func zipNum(str string) string {
    i := To_i(str)
    return carry62.Carry(i, 62)
}

func SplitText(text string) []string {
    var texts []string
    for _, t := range strings.Split(text, "\n") {
        // t = strings.Replace(t, "\r", "", -1)
        texts = append(texts, t)
    }
    return texts
}

func To_i(s string) int {
    i, _ := strconv.Atoi(s)
    return i
}

func DeUnicode(text string) string {
    for _, match := range MatchAll(text, `\\u\w{4}`) {
        replacement := deUnidcode(match)
        text = strings.Replace(text, match, replacement, 1)
    }

    for _, match := range MatchAll(text, `\\u\w{4}\\u\w{4}`) {
        replacement := GetEmoji(match)

        text = strings.Replace(text, match, replacement, 1)
    }

    return text
}

func deUnidcode(str string) string {
    deStr, err := strconv.Unquote(fmt.Sprintf(`"%s"`, str))
    if err != nil {
        return str
    }
    return deStr
}

func printError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}
