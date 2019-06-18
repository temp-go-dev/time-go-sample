package common

import (
	"strings"
	"time"

	pro "github.com/temp-go-dev/time-go-sample/property"
)

var (
	now  time.Time
	prop pro.Propertys
)

func GetNow() time.Time {

	timelayout := Timeformat("yyyy-MM-dd HH:mm:ss MST")

	if now.IsZero() {
		propertySetNow := prop.App_common_time_set_now
		if propertySetNow == "" {
			// 現在時刻
			now = time.Now()
			strNow := now.Format(timelayout)
			now, _ = time.Parse(timelayout, strNow)

		} else {
			// システムプロパティで現在時刻が設定されていれば取得
			now, _ = time.Parse(timelayout, propertySetNow)
		}
	}

	return now

}

func SetNowForDebug(nowDebug time.Time) time.Time {

	now = nowDebug

	return now

}

func GetNowString(layout string) string {

	var timeNow time.Time
	var fotmat string

	timeNow = time.Now()

	fotmat = Timeformat(layout)

	return timeNow.Format(fotmat)
}

func Timeformat(layout string) string {

	// 文字列中で指定したパターンにマッチする部分を置換する
	// 最後の引数は回数(<0で無制限)
	layout = strings.Replace(layout, "yyyy", "2006", -1)
	layout = strings.Replace(layout, "MM", "01", -1)
	layout = strings.Replace(layout, "dd", "02", -1)
	layout = strings.Replace(layout, "HH", "15", -1)
	layout = strings.Replace(layout, "mm", "04", -1)
	layout = strings.Replace(layout, "ss", "05", -1)

	return layout
}
