package src

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
)

const EARTHQUAKE_REPORT_TEMPLATE_FILE = "earthquake_report.gtpl"
const EARTHQUAKE_EPICENTER_TEMPLATE_FILE = "earthquake_epicenter.gtpl"
const EARTHQUAKE_INFO_TEMPLATE_FILE = "earthquake_info.gtpl"
const EARTHQUAKE_COMMENTS_TEMPLATE_FILE = "earthquake_comments.gtpl"

type EarthquakeType int

const (
	EarthquakeReport    EarthquakeType = 0
	EarthquakeEpicenter EarthquakeType = 1
	EarthquakeInfo      EarthquakeType = 2
	EarthquakeCancel    EarthquakeType = 3
)

// 震度速報、震源に関する情報、震源・震度に関する情報
type Earthquake struct {
	Row    string
	Parsed *jma.EarthquakeJmaXml

	Images []string
}

func ParseEarthquake(row []byte) (*Earthquake, error) {
	e, err := jma.ParseEarthquake(row)
	if err != nil {
		return nil, err
	}

	return &Earthquake{
		Row:    string(row),
		Parsed: e,
	}, nil
}

func (c *Earthquake) Assembly(ctx context.Context, db *sql.DB) error {
	// TODO: あとで
	return nil
}

func (c *Earthquake) SetImages() error {
	// TODO: あとで
	return nil
}

func (c *Earthquake) GetTitle() string {
	title := c.Parsed.Head.Title
	if c.GetInfoType() == jma.Cancel {
		title = fmt.Sprintf("【取消】%s", title)
	}

	return title
}

func (c *Earthquake) GetText() (string, error) {
	var tempFile string
	switch c.GetType() {
	case EarthquakeEpicenter:
		tempFile = EARTHQUAKE_EPICENTER_TEMPLATE_FILE
	case EarthquakeReport:
		tempFile = EARTHQUAKE_REPORT_TEMPLATE_FILE
	case EarthquakeInfo:
		tempFile = EARTHQUAKE_INFO_TEMPLATE_FILE
	case EarthquakeCancel:
		tempFile = EARTHQUAKE_COMMENTS_TEMPLATE_FILE
	}
	return Template(tempFile, c)
}

func (c *Earthquake) GetType() EarthquakeType {
	// 地の震度に関する情報
	// 出現するもの
	// - 震度速報
	// - 震源・震度情報
	intensity := c.Parsed.Body.Intensity

	// 地震の詳細（各地の震度など）
	// 出現するもの
	// - 震源に関する情報
	// - 震源・震度情報
	earthquake := c.Parsed.Body.Earthquake

	// 震源・震度情報
	if intensity != nil && earthquake != nil {
		return EarthquakeInfo
	}
	// 震度速報
	if intensity != nil && earthquake == nil {
		return EarthquakeReport
	}
	// 震源に関する情報
	if intensity == nil && earthquake != nil {
		return EarthquakeEpicenter
	}

	return EarthquakeCancel
}

func (c *Earthquake) GetTargetDate() (time.Time, error) {
	return time.Now(), nil
}

func (c *Earthquake) GetInfoType() jma.InfoType {
	return c.Parsed.Head.InfoType
}

func (c *Earthquake) GetImages() []string {
	return c.Images
}

func (c *Earthquake) GetEventId() ([]uint64, error) {
	return ParseEventID(c.Parsed.Head.EventID)
}
