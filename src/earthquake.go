package src

import (
	"encoding/xml"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
)

// 震度速報、震源に関する情報、震源・震度に関する情報
type Earthquake struct {
	Row    string
	Parsed *jma.EarthquakeJmaXml
}

func ParseEarthquake(row []byte) (*Earthquake, error) {
	earthquake := new(jma.EarthquakeJmaXml)
	err := xml.Unmarshal(row, earthquake)
	if err != nil {
		return nil, err
	}

	return &Earthquake{
		Row:    string(row),
		Parsed: earthquake,
	}, nil
}
