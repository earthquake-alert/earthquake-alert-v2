package src

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
)

const EARTHQUAKE_UPDATE_TEMPLATE_FILE = ""

// 顕著な地震の震源要素更新のお知らせ
type EarthquakeUpdate struct {
	Row    string
	Parsed *jma.EarthquakeUpdateInfoJmaXml

	Images []string
}

func ParseEarthquakeUpdate(row []byte) (*EarthquakeUpdate, error) {
	earthquake := new(jma.EarthquakeUpdateInfoJmaXml)
	err := xml.Unmarshal(row, earthquake)
	if err != nil {
		return nil, err
	}

	return &EarthquakeUpdate{
		Row:    string(row),
		Parsed: earthquake,
	}, nil
}

func (e *EarthquakeUpdate) Assembly() error {
	// TODO: 同じEventIDの震源を取得する
	return nil
}

func (e *EarthquakeUpdate) SetImages() error {
	// TODO: 画像生成できるようにする
	return nil
}

func (e *EarthquakeUpdate) GetTitle() string {
	title := "震源要素更新のお知らせ"
	if e.GetInfoType() == jma.Cancel {
		title = fmt.Sprintf("【取消】%s", title)
	}

	return title
}

func (e *EarthquakeUpdate) GetTargetDate() (time.Time, error) {
	targetTime := e.Parsed.Head.TargetDateTime

	return time.Parse("2006-01-02T15:04:05+09:00", targetTime)
}

func (e *EarthquakeUpdate) GetInfoType() jma.InfoType {
	return e.Parsed.Head.InfoType
}

func (e *EarthquakeUpdate) GetText() (string, error) {
	return Template(EARTHQUAKE_UPDATE_TEMPLATE_FILE, e)
}

func (e *EarthquakeUpdate) GetImages() []string {
	return e.Images
}

func (e *EarthquakeUpdate) GetEventId() string {
	return e.Parsed.Head.EventID
}
