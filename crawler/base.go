package crawler

import (
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/anaskhan96/soup"

    "iTools/conf"
    "iTools/script"
)

func getHttpRequest(url string) *http.Request {

    req, _ := http.NewRequest("GET", url, nil)

    req.Header.Set("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7`)

    // req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")

    if script.Match(url, `ithome`) {
        req.Header.Set(`User-Agent`, userAgent)
    }

    var cookie = `over18=1;`

    // fb instagram

    if script.IsFacebook(url) || script.IsInstagram(url) {
        cookie += conf.Cookie
    }

    // if script.IsWebSite(url, `twitter`) {
    //     cookie += `auth_token=5e465e71bccc07061ab6e9fa19b61a572b8f9260;`
    // }

    // fmt.Println(cookie)

    req.Header.Set("Cookie", cookie)

    return req
}

func GetDoc(url string) string {

    req := getHttpRequest(url)

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

    if conf.Mode == `debug` {
        script.WriteFile(string(body), `z:\doc.html`)
    }

    return string(body)
}

func GetRedirectURL(url string) string {
    req := getHttpRequest(url)

    // 檢查轉址 https://fb.watch/\w+/ => https://www.facebook.com/watch?v=\d+
    client := &http.Client{
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        }}

    res, err := client.Do(req)
    if err != nil {
        return ""
    }

    if res != nil && res.StatusCode == http.StatusFound {
        redirectURL, _ := res.Location()
        return redirectURL.String()
    }

    return url
}

func GetURLSoup(url string) soup.Root {
    txt := GetDoc(url)

    return soup.HTMLParse(txt)
}

func GetTextSoup(docText string) soup.Root {

    return soup.HTMLParse(docText)
}

var userAgent = `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Mobile Safari/537.36`

var _ = fmt.Println
