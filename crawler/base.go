package crawler

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

    "github.com/anaskhan96/soup"

    "iTools/script"
)

var userAgent = `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Mobile Safari/537.36`

func GetDoc(url string) string {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return ""
    }

    req.Header.Set("Accept", `text/html`)

    if script.Match(url, `ithome`) {
        req.Header.Set(`User-Agent`, userAgent)
    }

    // fb instagram
    cookie := os.Getenv("cookie")
    req.Header.Set("Cookie", cookie)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return ""
    }
    defer res.Body.Close()

    if !script.Match(url, `ithome`) {
        res.Header.Set(`User-Agent`, userAgent)
    }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        return ""
    }

    // writeFile(string(body), `z:\doc.html`)
    return string(body)
}

func GetDocSoup(url string) soup.Root {
    txt := GetDoc(url)

    return soup.HTMLParse(txt)
}

func writeFile(content, path string) {
    err := ioutil.WriteFile(path, []byte(content), 0777)
    printError(err)
}

func printError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}
