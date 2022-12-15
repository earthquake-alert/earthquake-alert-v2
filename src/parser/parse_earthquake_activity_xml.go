package parser

import "encoding/xml"

// 地震の活動状況等に関する情報
type EarthquakeActivity struct {
	Row    string
	Parsed *EarthquakeActivityJmaXml
}

func ParseEarthquakeActivity(row []byte) (*EarthquakeActivity, error) {
	earthquake := new(EarthquakeActivityJmaXml)
	err := xml.Unmarshal(row, earthquake)
	if err != nil {
		return nil, err
	}

	return &EarthquakeActivity{
		Row:    string(row),
		Parsed: earthquake,
	}, nil
}
