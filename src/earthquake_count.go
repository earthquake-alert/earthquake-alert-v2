package src

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
	"github.com/earthquake-alert/erarthquake-alert-v2/src/logging"
)

const EARTHQUAKE_COUNT_TEMPLATE_FILE = "earthquake_count.tmpl"
const EARTHQUAKE_COUNT_CANCEL_TEMPLATE_FILE = "earthquake_count.tmpl"

// 地震回数に関する情報
type EarthquakeCount struct {
	Row    string
	Parsed *jma.EarthquakeCountInfoJmaXml
}

type EarthquakeCountDetail struct {
	FormattedStartTime string
	FormattedEndTime   string
	Number             int
	FeltNumber         int
}

func ParseEarthquakeCount(row []byte) (*EarthquakeCount, error) {
	earthquake := new(jma.EarthquakeCountInfoJmaXml)
	err := xml.Unmarshal(row, earthquake)
	if err != nil {
		return nil, err
	}

	return &EarthquakeCount{
		Row:    string(row),
		Parsed: earthquake,
	}, nil
}

func (e *EarthquakeCount) Assembly() error {
	// 関連する情報がDB内にある場合などはここで引く
	return nil
}

// 画像は生成しない
func (e *EarthquakeCount) SetImages() error {
	return nil
}

func (e *EarthquakeCount) GetTitle() string {
	title := e.Parsed.Head.Title
	if e.GetInfoType() == jma.Cancel {
		title = fmt.Sprintf("【取消】%s", title)
	}

	return title
}

func (e *EarthquakeCount) GetTargetDate() (time.Time, error) {
	targetTime := e.Parsed.Head.TargetDateTime

	return time.Parse("2006-01-02T15:04:05+09:00", targetTime)
}

func (e *EarthquakeCount) GetInfoType() jma.InfoType {
	return e.Parsed.Head.InfoType
}

func (e *EarthquakeCount) GetText() (string, error) {
	if e.GetInfoType() == jma.Cancel {
		return Template(EARTHQUAKE_COUNT_CANCEL_TEMPLATE_FILE, e)
	}

	return Template(EARTHQUAKE_COUNT_TEMPLATE_FILE, e)
}

// テンプレートで使うやつ
func (e *EarthquakeCount) TempCountText() *EarthquakeCountDetail {
	fmt.Println("OK")
	eaItems := e.Parsed.Body.EarthquakeCount
	if eaItems == nil {
		return nil
	}
	for _, e := range eaItems.Item {
		if e.Type == "累積地震回数" {
			startTime, err := time.Parse("2006-01-02T15:04:05+09:00", e.StartTime)
			if err != nil {
				logging.Sugar.Error(err)
				return nil
			}
			endTime, err := time.Parse("2006-01-02T15:04:05+09:00", e.EndTime)
			if err != nil {
				logging.Sugar.Error(err)
				return nil
			}

			return &EarthquakeCountDetail{
				FormattedStartTime: fmt.Sprintf("%d年%d月%d日%d時ごろ", startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour()),
				FormattedEndTime:   fmt.Sprintf("%d年%d月%d日%d時ごろ", endTime.Year(), endTime.Month(), endTime.Day(), endTime.Hour()),
				Number:             e.Number,
				FeltNumber:         e.FeltNumber,
			}
		}
	}

	return nil
}

func (e *EarthquakeCount) TempFreeFormComment() string {
	c := e.Parsed.Body.Comments.FreeFormComment

	if c == "" {
		return ""
	}

	return Convert(c, true)
}

// 画像は生成しない
func (e *EarthquakeCount) GetImages() []string {
	return []string{}
}

func (e *EarthquakeCount) GetEventId() string {
	return e.Parsed.Head.EventID
}
