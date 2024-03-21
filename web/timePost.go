package web

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"

    "iTools/script"
)

func UtimePost(c *gin.Context) {
    timeFormatText := c.PostForm("text")

    var returnUrl = `/time?`

    // 替算今天日期
    today := script.GetToday()
    timeFormatText = strings.Replace(timeFormatText, "today", today, -1)

    //  日期加減
    //  1970-01-01+5day
    data := script.ScanGroups(timeFormatText, `(\d+).(\d+).(\d+)(\+|\-)(\d+)day`)
    if len(data) == 6 && returnUrl == "" {
        year, month, day := script.To_i(data[1]), script.To_i(data[2]), script.To_i(data[3])
        addDay := script.To_i(data[5])
        if data[4] == "-" {
            addDay *= -1
        }

        timeTime := script.GetTime(year, month, day, 0, 0, 0).AddDate(0, 0, addDay)
        timeStamp := script.GetDate(timeTime)

        returnUrl += fmt.Sprintf(`timeformat=%s`, timeStamp)
    }

    //  計算天數差
    //  1970-01-01--1995-01-01
    data = script.ScanGroups(timeFormatText, `(\d+).(\d+).(\d+)--(\d+).(\d+).(\d+)`)
    if len(data) == 7 && returnUrl == "" {
        year1, month1, day1 := script.To_i(data[1]), script.To_i(data[2]), script.To_i(data[3])
        year2, month2, day2 := script.To_i(data[4]), script.To_i(data[5]), script.To_i(data[6])

        dayDiff := script.DayDiff(year1, month1, day1, year2, month2, day2)

        returnUrl += fmt.Sprintf(`utime=%v`, dayDiff)
    }

    // 2000-01-01 08:00:00
    data = script.ScanGroups(timeFormatText, `(\d+).(\d+).(\d+)\s(\d+).(\d+).(\d+)`)
    if len(data) == 7 && returnUrl == "" {
        year, month, day := script.To_i(data[1]), script.To_i(data[2]), script.To_i(data[3])
        hour, min, sec := script.To_i(data[4]), script.To_i(data[5]), script.To_i(data[6])

        timeTime := script.GetTime(year, month, day, hour, min, sec)
        utime := script.GetUtime(timeTime)

        returnUrl += fmt.Sprintf(`utime=%v`, utime)
        c.Redirect(http.StatusMovedPermanently, returnUrl)
    }

    // 2000-01-01
    data = script.ScanGroups(timeFormatText, `(\d+).(\d+).(\d+)`)
    if len(data) == 4 && returnUrl == "" {
        year, month, day := script.To_i(data[1]), script.To_i(data[2]), script.To_i(data[3])

        timeTime := script.GetTime(year, month, day, 0, 0, 0)
        utime := script.GetUtime(timeTime)

        returnUrl += fmt.Sprintf(`utime=%v`, utime)
    }

    c.Redirect(http.StatusMovedPermanently, returnUrl)
}

func TimeFormatPost(c *gin.Context) {
    utimeText := c.PostForm("text")
    utime := script.To_i(utimeText)

    var returnUrl = `/time`

    if utime < 0 {
        setMeg(c, "alert", "UnixTime 需要大於零")
    } else {

        timeTime := script.GetTimeTime(utime)
        timestamp := script.GetTimeStamp(timeTime)

        week := script.GetWeek(timeTime.Year(), int(timeTime.Month()), timeTime.Day())
        weekStr := []string{`一`, `二`, `三`, `四`, `五`, `六`, `日`}[week]

        returnUrl += fmt.Sprintf(`?timeformat=%s %s`, timestamp, weekStr)
    }

    c.Redirect(http.StatusMovedPermanently, returnUrl)
}

var _ = fmt.Println
