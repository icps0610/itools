package script

import (
    "fmt"
    "time"
)

func GetToday() string {
    t := time.Now()
    return fmt.Sprintf("%d-%02d-%02d", t.Year(), int(t.Month()), t.Day())
}

func GetTime(y, m, d, h, mi, sec int) time.Time {
    date := fmt.Sprintf("%v-%.2v-%.2vT%.2v:%.2v:%.2v.000+08:00", y, m, d, h, mi, sec)
    t, _ := time.Parse(time.RFC3339, date)
    return t
}

func GetDate(t time.Time) string {
    return fmt.Sprintf("%d-%02d-%02d", t.Year(), int(t.Month()), t.Day())
}

func GetTimeStamp(t time.Time) string {
    return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func GetUtime(t time.Time) int {
    return int(t.Unix())
}

func GetTimeTime(unixTime int) time.Time {
    return time.Unix(int64(unixTime), 0)
}

func DayDiff(year1, month1, day1, year2, month2, day2 int) int {
    timeTime1 := GetTime(year1, month1, day1, 0, 0, 0)
    timeTime2 := GetTime(year2, month2, day2, 0, 0, 0)

    utime1 := GetUtime(timeTime1)
    utime2 := GetUtime(timeTime2)

    return (utime1 - utime2) / 86400
}

func GetWeek(year, month, day int) int {
    monthValue := []int{6, 2, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}[month-1]
    yearValue := year % 100
    centuryValue := []int{0, 5, 3, 1}[(year/100)%4]
    w := day + monthValue + yearValue + yearValue/4 + centuryValue
    if isLeapYear(year) && month <= 2 {
        w -= 1
    }
    return w % 7
}

func isLeapYear(year int) bool {
    return (year%4 == 0 && year%100 != 0) || (year%400 == 0 && year%3200 != 0)
}
