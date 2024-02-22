package script

import (
    "regexp"
)

func Match(str, keyword string) bool {
    match, _ := regexp.MatchString(keyword, str)
    return match
}

func Scan(str, keyword string, i int) string {
    re := regexp.MustCompile(keyword)
    match := re.FindStringSubmatch(str)

    if len(match) > i {
        return match[i]
    }
    return ""
}

func ScanGroup(str, keyword string, is ...int) []string {
    re := regexp.MustCompile(keyword)
    match := re.FindStringSubmatch(str)
    if len(match) == 0 {
        return []string{str}
    }

    var result []string
    for _, i := range is {
        result = append(result, match[i])
    }
    return result
}

func ScanGroups(str, keyword string) []string {
    re := regexp.MustCompile(keyword)

    return re.FindStringSubmatch(str)
}

func MatchAll(str, keyword string, is ...int) []string {
    re := regexp.MustCompile(keyword)

    return re.FindAllString(str, -1)
}
