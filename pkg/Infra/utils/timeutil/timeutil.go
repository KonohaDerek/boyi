package timeutil

import (
	"time"

	"github.com/jinzhu/now"
)

// 取得起始跟結束時間, +8時間 的 起始與結束 XXXX/XX/XX 00:00 +08:00 - XXXX/XX/XX 23:59 +08:00 
func GetTimeStartAndEndAt(t time.Time) (start time.Time, end time.Time) {
	zone := time.FixedZone("", 8*60*60)

	timeNow := t.In(zone)

	start = now.With(timeNow).BeginningOfDay()
	end = now.With(timeNow).EndOfDay()

	return start, end
}
