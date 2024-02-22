package url

import (
// "iTools/simpToTrad"
)

func Run(url string) (string, string) {
    url = processURL(url)

    title := GetUrlTitle(url)
    // title = simpToTrad.Run(title)

    return url, title
}
