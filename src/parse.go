package src

import (
	"strconv"
	"strings"
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
