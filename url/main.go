package url

import (
    "fmt"

    "iTools/crawler"
    // "iTools/simpToTrad"
)

func Run(url string) (string, string) {
    url = processURL(url)

    title := crawler.GetUrlTitle(url)
    // title = simpToTrad.Run(title)

    return url, title
}

var _ = fmt.Println
