package src

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
	"github.com/earthquake-alert/erarthquake-alert-v2/src/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const EARTHQUAKE_UPDATE_TEMPLATE_FILE = "earthquake_update.gtpl"

// 顕著な地震の震源要素更新のお知らせ
type EarthquakeUpdate struct {
	Row    string
	Parsed *jma.EarthquakeUpdateInfoJmaXml

	Images []string

	// 通常のHypocenterの型と違うため独自で定義する
	// 取消報の場合は空になる
	NewName      string
	NewLat       *float64
	NewLon       *float64
	NewDepth     *int
	NewLatLonStr string

	NewMagnitude string

	// 更新前の地震情報
	LatestEarthquake *models.Earthquake
}

func ParseEarthquakeUpdate(row []byte) (*EarthquakeUpdate, error) {
	e, err := jma.ParseEarthquakeUpdate(row)
	if err != nil {
		return nil, err
	}

	return &EarthquakeUpdate{
		Row:    string(row),
		Parsed: e,
	}, nil
}

func (e *EarthquakeUpdate) Assembly(ctx context.Context, db *sql.DB) error {
	eventId, err := e.GetEventId()
	if err != nil {
		return err
	}
	targetD, err := e.GetTargetDate()
	if err != nil {
		return err
	}

	if err := e.InsertUpdateDB(ctx, db, eventId[0], targetD); err != nil {
		return err
	}

	// 取消報の時
	// FIXME: DBに入れる
	if e.Parsed.Body.Earthquake == nil {
		return nil
	}

	// 震源要素をパースする
	if err := e.ParseEpicenter(); err != nil {
		return err
	}

	// マグニチュード表記を文字列にフォーマットする
	e.NewMagnitude = FormatMagnitude(&e.Parsed.Body.Earthquake.Magnitude)

	ea, err := e.GetOldEarthquakes(ctx, db, int64(eventId[0]))
	if err != nil {
		return err
	}
	if ea != nil {
		e.LatestEarthquake = ea
	}

	update := models.EarthquakeUpdate{
		EventID:   int64(eventId[0]),
		Lat:       null.Float64FromPtr(e.NewLat),
		Lon:       null.Float64FromPtr(e.NewLon),
		Depth:     null.IntFromPtr(e.NewDepth),
		Magnitude: null.NewString(e.NewMagnitude, true),
		Date:      targetD,
		Row:       e.Row,
	}
	if err := update.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}

	// Earthquakes テーブルを更新する
	// FIXME: 重複している
	earthquake, err := models.Earthquakes(
		models.EarthquakeWhere.EventID.EQ(int64(eventId[0])),
	).One(ctx, db)
	if err != nil {
		return err
	}
	earthquake.Lat = null.Float64FromPtr(e.NewLat)
	earthquake.Lon = null.Float64FromPtr(e.NewLon)
	earthquake.Depth = null.IntFromPtr(e.NewDepth)
	earthquake.Magnitude = null.NewString(e.NewMagnitude, true)

	if _, err := earthquake.Update(ctx, db, boil.Infer()); err != nil {
		return err
	}
	return nil
}

func (e *EarthquakeUpdate) InsertUpdateDB(ctx context.Context, db *sql.DB, eventId int64, targetDate time.Time) error {

	update := models.EarthquakeUpdate{
		EventID:   eventId,
		Lat:       null.Float64FromPtr(e.NewLat),
		Lon:       null.Float64FromPtr(e.NewLon),
		Depth:     null.IntFromPtr(e.NewDepth),
		Magnitude: null.NewString(e.NewMagnitude, true),
		Date:      targetDate,
		Row:       e.Row,
	}
	return update.Insert(ctx, db, boil.Infer())
}

// 更新前の地震情報を取得する
func (e *EarthquakeUpdate) GetOldEarthquakes(ctx context.Context, db *sql.DB, eventId int64) (*models.Earthquake, error) {
	ea, err := models.Earthquakes(
		models.EarthquakeWhere.EventID.EQ(eventId),
	).One(ctx, db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return ea, nil
}

// Coordinateを取得してパースする
func (e *EarthquakeUpdate) ParseEpicenter() error {
	ea := e.Parsed.Body.Earthquake
	if ea == nil {
		return nil
	}

	e.NewName = ea.Hypocenter.Area.Name

	find := false
	for _, epi := range ea.Hypocenter.Area.Coordinate {
		// 日本測地系のものを取得する
		if epi.Datum == "日本測地系" && epi.Type == "" {
			lat, lon, depth, err := ParseCoordinate(epi.Value)
			if err != nil {
				return err
			}

			e.NewLat = &lat
			e.NewLon = &lon
			e.NewDepth = &depth

			if epi.Description != "" {
				e.NewLatLonStr = Convert(epi.Description, true)
			} else {
				e.NewLatLonStr = "不明"
			}

			find = true
		}
	}

	// 見つからなかった場合、エラー
	if !find {
		return errors.New("earthquake epicenter coordinate is not found")
	}
	return nil
}

func (e *EarthquakeUpdate) SetImages() error {
	// TODO: 画像生成できるようにする
	return nil
}

// 前回の震源情報を返す
func (e *EarthquakeUpdate) GetOldEpicenter() string {
	if e.LatestEarthquake == nil {
		return "不明"
	}

	// 緯度経度がない場合、全ての震源要素が無いに等しいため不明と返す
	if e.LatestEarthquake.Lat.IsZero() {
		return "不明"
	}
	if e.LatestEarthquake.Lon.IsZero() {
		return "不明"
	}

	lat := e.LatestEarthquake.Lat.Float64
	lon := e.LatestEarthquake.Lon.Float64

	var depth *int = nil
	if !e.LatestEarthquake.Depth.IsZero() {
		depth = &e.LatestEarthquake.Depth.Int
	}

	return FormatLatLonDepth(lat, lon, depth)
}

func (e *EarthquakeUpdate) GetOldMagnitude() string {
	if e.LatestEarthquake == nil {
		return "M不明"
	}

	return e.LatestEarthquake.Magnitude.String
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

func (e *EarthquakeUpdate) GetEventId() ([]int, error) {
	eventId, err := strconv.Atoi(e.Parsed.Head.EventID)
	if err != nil {
		return nil, err
	}
	return []int{eventId}, nil
}
