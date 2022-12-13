package src

import "encoding/xml"

// 震度速報、震源に関する情報、震源・震度に関する情報
type Earthquake struct {
	Row    string
	Parsed *EarthquakeJmaXml
}

// 地震の活動状況等に関する情報
type EarthquakeActivity struct {
	Row    string
	Parsed *EarthquakeActivityJmaXml
}

// 地震回数に関する情報
type EarthquakeCount struct {
	Row    string
	Parsed *EarthquakeCountInfoJmaXml
}

// 顕著な地震の震源要素更新のお知らせ
type EarthquakeUpdate struct {
	Row    string
	Parsed *EarthquakeUpdateInfoJmaXml
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
