package parser

import "encoding/xml"

// 顕著な地震の震源要素更新のお知らせ
type EarthquakeUpdate struct {
	Row    string
	Parsed *EarthquakeUpdateInfoJmaXml
}

func ParseEarthquakeUpdate(row []byte) (*EarthquakeUpdate, error) {
	earthquake := new(EarthquakeUpdateInfoJmaXml)
	err := xml.Unmarshal(row, earthquake)
	if err != nil {
		return nil, err
	}

	return &EarthquakeUpdate{
		Row:    string(row),
		Parsed: earthquake,
	}, nil
}
