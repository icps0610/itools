package crawler

import (
    "iTools/script"
)

func GetDocLink(url string) string {
    doc := GetDocSoup(url)
    link := script.GetLinkedInLink(doc)

    return link
}
