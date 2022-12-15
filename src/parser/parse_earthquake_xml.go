package parser

import "encoding/xml"

// 震度速報、震源に関する情報、震源・震度に関する情報
type Earthquake struct {
	Row    string
	Parsed *EarthquakeJmaXml
}

func ParseEarthquake(row []byte) (*Earthquake, error) {
	tsunami := new(EarthquakeJmaXml)
	err := xml.Unmarshal(row, tsunami)
	if err != nil {
		return nil, err
	}

	return &Earthquake{
		Row:    string(row),
		Parsed: tsunami,
	}, nil
}
