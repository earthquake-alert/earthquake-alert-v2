package src

import (
	"encoding/xml"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
)

// 地震の活動状況等に関する情報
type EarthquakeActivity struct {
	Row    string
	Parsed *jma.EarthquakeActivityJmaXml
}

func ParseEarthquakeActivity(row []byte) (*EarthquakeActivity, error) {
	earthquake := new(jma.EarthquakeActivityJmaXml)
	err := xml.Unmarshal(row, earthquake)
	if err != nil {
		return nil, err
	}

	return &EarthquakeActivity{
		Row:    string(row),
		Parsed: earthquake,
	}, nil
}