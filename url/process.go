package url

import (
    "fmt"

    "iTools/script"
)

func processURL(url string) string {
    url = removeFBID(url)
    url = removeUtmSource(url)
    url = processMobile(url)

    if script.IsFacebook(url) {
        url = processFB(url)
    } else {
        url = processYoutube(url)
        url = processNetflix(url)
    }

    url = script.DeUrlCode(url)

    return url
}

func processMobile(url string) string {
    keyword := `https\:\/\/m\.(.*)`
    match := script.Match(url, keyword)
    if match {
        return `https://www.` + script.Scan(url, keyword, 1)
    }
    return url
}

func processFB(url string) string {

    return url
}

func processYoutube(url string) string {
    keyword := `(?:youtube\.com\/(?:[^\/]+\/.+\/|(?:v|e(?:mbed)?)\/|.*[?&]v=)|youtu\.be\/)([A-Za-z0-9_-]{11})`
    if script.Match(url, keyword) {
        id := script.Scan(url, keyword, 1)
        url = `https://youtu.be/` + id
        //cmd := fmt.Sprintf(`curl https://img.youtube.com/vi/%s/maxresdefault.jpg --output z:\u2.png`, id)
        //runCmd(cmd)
    }
    return url
}

func processNetflix(url string) string {
    kw := `(https:\/\/www.netflix.com\/watch\/\d+)\?`
    rsg := script.ScanGroup(url, kw, 1)
    if len(rsg) > 0 {
        return rsg[0]
    }
    return url
}

var _ = fmt.Println
