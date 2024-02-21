package simpToTrad

import (
    "net/url"

    "iTools/crawler"
    "iTools/script"
)

func Run(str string) string {
    // https://docs.zhconvert.org/api/convert/
    // Taiwan Traditional Simplified

    str = `https://api.zhconvert.org/convert?converter=Taiwan&text=` + url.QueryEscape(str)
    bodyText := crawler.GetDoc(str)
    return script.Scan(bodyText, `"text":"(.*)","diff"`, 1)
}
