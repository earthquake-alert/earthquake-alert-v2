package parser

import "encoding/xml"

// 地震回数に関する情報
type EarthquakeCount struct {
	Row    string
	Parsed *EarthquakeCountInfoJmaXml
}

func ParseEarthquakeCount(row []byte) (*EarthquakeCount, error) {
	earthquake := new(EarthquakeCountInfoJmaXml)
	err := xml.Unmarshal(row, earthquake)
	if err != nil {
		return nil, err
	}

	return &EarthquakeCount{
		Row:    string(row),
		Parsed: earthquake,
	}, nil
}
