package src

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
	"github.com/earthquake-alert/erarthquake-alert-v2/src/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const EARTHQUAKE_ACTIVITY_TEMPLATE_FILE = "earthquake_activity.gtpl"

// 地震の活動状況等に関する情報
type EarthquakeActivity struct {
	Row    string
	Parsed *jma.EarthquakeActivityJmaXml
}

func ParseEarthquakeActivity(row []byte) (*EarthquakeActivity, error) {
	e, err := jma.ParseEarthquakeActivity(row)
	if err != nil {
		return nil, err
	}

	return &EarthquakeActivity{
		Row:    string(row),
		Parsed: e,
	}, nil
}

func (e *EarthquakeActivity) Assembly(ctx context.Context, db *sql.DB) error {
	eventId, err := e.GetEventId()
	if err != nil {
		return err
	}
	d, err := e.GetTargetDate()
	if err != nil {
		return err
	}

	activity := &models.EarthquakeActivity{
		EventID: int64(eventId[0]),
		Date:    d,
		Row:     e.Row,
	}

	return activity.Insert(ctx, db, boil.Infer())
}

// 画像は生成しないのでなにもしない
func (e *EarthquakeActivity) SetImages() error {
	return nil
}

// タイトル
// 取り消し報の場合`【取消】`を付与する
func (e *EarthquakeActivity) GetTitle() string {
	title := e.Parsed.Control.Title

	if e.GetInfoType() == jma.Cancel {
		title = fmt.Sprintf("【取消】%s", title)
	}

	return title
}

func (e *EarthquakeActivity) GetTargetDate() (time.Time, error) {
	targetTime := e.Parsed.Head.TargetDateTime

	return time.Parse("2006-01-02T15:04:05+09:00", targetTime)
}

func (e *EarthquakeActivity) GetInfoType() jma.InfoType {
	return e.Parsed.Head.InfoType
}

func (e *EarthquakeActivity) GetText() (string, error) {
	return Template(EARTHQUAKE_ACTIVITY_TEMPLATE_FILE, e)
}

// 画像は生成しないので空
func (e *EarthquakeActivity) GetImages() []string {
	return []string{}
}

func (e *EarthquakeActivity) GetEventId() ([]int, error) {
	eventId, err := strconv.Atoi(e.Parsed.Head.EventID)
	if err != nil {
		return nil, err
	}
	return []int{eventId}, nil
}
