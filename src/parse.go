package src

import (
	"math"
	"strconv"
	"strings"
	"time"
)

// EventIDをパースする
//
// 津波関連の電文では、半角スペースで区切られた複数のEventIDが存在することがある
func ParseEventID(eventId string) ([]uint64, error) {
	splitEventId := strings.Split(eventId, " ")
	parsedEventIds := make([]uint64, len(splitEventId))

	for i, e := range splitEventId {
		parsedEventId, err := strconv.ParseUint(e, 10, 64)
		if err != nil {
			return nil, err
		}
		parsedEventIds[i] = parsedEventId
	}
	return parsedEventIds, nil
}

// ISO8601拡張形式をパースする
func ParseDate(d string) (time.Time, error) {
	return time.Parse(time.RFC3339, d)
}

// 度分秒の緯度経度十進度に変換する
// = 度 + (分 / 60) + (秒 / 60)
// 引数は、00.00の形式で、[度].[分]となる。
func ParsePosition(position float64) float64 {
	isNegative := false
	if position < 0 {
		isNegative = true
	}

	p := math.Abs(position)
	d := math.Floor(p)
	m := math.Round(math.Mod(p, d)/60*10000) / 100

	if isNegative {
		return -d + m
	}
	return d + m
}
