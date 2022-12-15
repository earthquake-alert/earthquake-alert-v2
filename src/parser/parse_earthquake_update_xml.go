package parser

// 顕著な地震の震源要素更新のお知らせ
type EarthquakeUpdate struct {
	Row    string
	Parsed *EarthquakeUpdateInfoJmaXml
}
