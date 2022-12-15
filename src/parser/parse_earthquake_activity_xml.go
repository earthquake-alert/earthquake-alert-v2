package parser

// 地震の活動状況等に関する情報
type EarthquakeActivity struct {
	Row    string
	Parsed *EarthquakeActivityJmaXml
}
