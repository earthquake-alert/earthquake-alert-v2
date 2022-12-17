package src

import (
	"encoding/xml"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
)

const EARTHQUAKE_ACTIVITY_TEMPLATE_FILE = "earthquake_activity.tmpl"

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

func (e *EarthquakeActivity) GetText() (string, error) {
	return Template(EARTHQUAKE_ACTIVITY_TEMPLATE_FILE, e)
}

// テンプレートに使用するやつ
func (e *EarthquakeActivity) TempFormatHeadlineText() string {
	return Convert(e.Parsed.Head.Headline.Text, false)
}

// テンプレートに使用するやつ
func (e *EarthquakeActivity) TempFormatBodyText() string {
	return Convert(e.Parsed.Body.Text, false)
}
