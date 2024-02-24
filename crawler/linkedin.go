package crawler

import (
    "iTools/script"
)

func GetDocLink(url string) string {
    doc := GetURLSoup(url)
    link := script.GetLinkedInLink(doc)

    return link
}
