package src

import (
	"encoding/xml"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
)

// 地震回数に関する情報
type EarthquakeCount struct {
	Row    string
	Parsed *jma.EarthquakeCountInfoJmaXml
}

func ParseEarthquakeCount(row []byte) (*EarthquakeCount, error) {
	earthquake := new(jma.EarthquakeCountInfoJmaXml)
	err := xml.Unmarshal(row, earthquake)
	if err != nil {
		return nil, err
	}

	return &EarthquakeCount{
		Row:    string(row),
		Parsed: earthquake,
	}, nil
}
