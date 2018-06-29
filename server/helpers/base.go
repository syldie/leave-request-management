package helpers

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

// CheckErr ...
func CheckErr(errMsg string, err error) {
	if err != nil {
		beego.Warning(errMsg, err)
	}
}

// BytesToString ...
func BytesToString(data []byte) string {
	return string(data[:])
}

// ArrayToString ...
func ArrayToString(arr []int64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(arr), " ", delim, -1), "[]")
}

// GetTotalDay ...
func GetTotalDay(from string, to string) int64 {
	dateFrom, _ := time.Parse("2006-01-02", from)
	dateTo, _ := time.Parse("2006-01-02", to)
	diff := dateTo.Sub(dateFrom)

	return int64(diff.Hours() / 24)
}
