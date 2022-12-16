package src

import (
	"encoding/xml"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
)

// 顕著な地震の震源要素更新のお知らせ
type EarthquakeUpdate struct {
	Row    string
	Parsed *jma.EarthquakeUpdateInfoJmaXml
}

func ParseEarthquakeUpdate(row []byte) (*EarthquakeUpdate, error) {
	earthquake := new(jma.EarthquakeUpdateInfoJmaXml)
	err := xml.Unmarshal(row, earthquake)
	if err != nil {
		return nil, err
	}

	return &EarthquakeUpdate{
		Row:    string(row),
		Parsed: earthquake,
	}, nil
}
