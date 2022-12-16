package src_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
	"github.com/stretchr/testify/require"
)

// 震度速報、震源に関する情報、震源・震度に関する情報のテストデータ
var TestEarthquakeData = []string{
	// 震度速報
	"32-35_01_01_100806_VXSE51.xml",
	"32-35_04_01_100831_VXSE51.xml",
	"32-35_04_02_100831_VXSE51.xml",
	"32-35_04_05_100831_VXSE51.xml",
	"32-35_07_02_100915_VXSE51.xml",
	"32-35_07_03_100915_VXSE51.xml",
	"32-35_07_04_100915_VXSE51.xml",
	"32-35_07_05_100915_VXSE51.xml",
	"32-35_08_01_100915_VXSE51.xml",
	"32-35_08_02_100915_VXSE51.xml",
	"32-35_08_03_100915_VXSE51.xml",
	"32-35_08_04_100915_VXSE51.xml",
	"32-35_08_05_100915_VXSE51.xml",
	"32-39_11_01_120615_VXSE51.xml",
	"32-35_10_01_220510_VXSE51.xml",
	// 震源に関する情報
	"32-35_01_02_100514_VXSE52.xml",
	"32-35_04_03_100831_VXSE52.xml",
	"32-35_06_01_100915_VXSE52.xml",
	"32-35_06_02_100915_VXSE52.xml",
	"32-35_08_06_100915_VXSE52.xml",
	"33_12_01_120615_VXSE52.xml",
	// 震源・震度に関する情報
	"32-35_01_03_100514_VXSE53.xml",
	"32-35_01_03_100806_VXSE53.xml",
	"32-35_04_04_100831_VXSE53.xml",
	"32-35_04_06_100831_VXSE53.xml",
	"32-35_06_03_100915_VXSE53.xml",
	"32-35_06_04_100915_VXSE53.xml",
	"32-35_06_05_100915_VXSE53.xml",
	"32-35_06_06_100915_VXSE53.xml",
	"32-35_07_06_100915_VXSE53.xml",
	"32-35_07_07_100915_VXSE53.xml",
	"32-35_08_07_100915_VXSE53.xml",
	"32-35_08_08_100915_VXSE53.xml",
	"32-39_05_01_100831_VXSE53.xml",
	"32-39_05_02_100831_VXSE53.xml",
	"32-39_05_04_100831_VXSE53.xml",
	"32-39_05_06_100831_VXSE53.xml",
	"32-39_05_07_100831_VXSE53.xml",
	"32-39_05_08_100831_VXSE53.xml",
	"32-39_05_12_100831_VXSE53.xml",
	"32-39_05_12_100915_VXSE53.xml",
	"32-39_11_05_120615_VXSE53.xml",
}

func TestParseEarthquake(t *testing.T) {
	for _, d := range TestEarthquakeData {
		t.Run(fmt.Sprintf("Test %s", d), func(t *testing.T) {
			testPath := filepath.Join(TEST_DATA_PATH, d)

			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			_, err = src.ParseEarthquake(row)
			require.NoError(t, err)
		})
	}

	t.Run("failed", func(t *testing.T) {
		row := "aaaaaaaa"

		_, err := src.ParseEarthquake([]byte(row))
		require.Error(t, err)
	})
}

// 震度速報パーステスト
func TestParseEarthquakeReport1(t *testing.T) {
	target := "32-35_01_01_100806_VXSE51.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震度速報",
			DateTime:         "2009-08-10T20:09:11Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "震度速報")
		require.Equal(t, head.ReportDateTime, "2009-08-11T05:09:00+09:00")
		require.Equal(t, head.TargetDateTime, "2009-08-11T05:07:00+09:00")
		require.Equal(t, head.EventID, "20090811050711")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震度速報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "１１日０５時０７分ころ、地震による強い揺れを感じました。震度３以上が観測された地域をお知らせします。")

		// Information
		require.Len(t, head.Headline.Information, 1)
		require.Equal(t, head.Headline.Information[0].Type, "震度速報")

		require.Len(t, head.Headline.Information[0].Item, 4)
		require.Equal(t, head.Headline.Information[0].Item[0].Kind.Name, "震度６弱")

		require.Equal(t, head.Headline.Information[0].Item[0].Areas.CodeType, "地震情報／細分区域")
		require.Len(t, head.Headline.Information[0].Item[0].Areas.Area, 2)
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Name, "静岡県中部")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Code, "442")
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		require.Nil(t, body.Earthquake)

		t.Run("intensity", func(t *testing.T) {
			intensity := body.Intensity

			require.NotNil(t, intensity)

			require.Len(t, intensity.Observation.CodeDefine.Types, 2)
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].Value, "地震情報／都道府県等")
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].XPath, "Pref/Code")

			require.Equal(t, intensity.Observation.MaxInt, jma.Int6l)

			require.Len(t, intensity.Observation.Pref, 16)
			require.Equal(t, intensity.Observation.Pref[0].Name, "静岡県")
			require.Equal(t, intensity.Observation.Pref[0].Code, "22")
			require.Equal(t, intensity.Observation.Pref[0].MaxInt, jma.Int6l)

			require.Len(t, intensity.Observation.Pref[0].Area, 4)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Name, "静岡県中部")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Code, "442")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].MaxInt, jma.Int6l)
		})

		require.Equal(t, body.Text, "")
		require.Nil(t, body.Comments)

	})
}

// 震度速報パーステスト（東日本大震災）
func TestParseEarthquakeReport2(t *testing.T) {
	target := "32-39_11_01_120615_VXSE51.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震度速報",
			DateTime:         "2011-03-11T05:48:10Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "震度速報")
		require.Equal(t, head.ReportDateTime, "2011-03-11T14:48:00+09:00")
		require.Equal(t, head.TargetDateTime, "2011-03-11T14:46:00+09:00")
		require.Equal(t, head.EventID, "20110311144640")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震度速報")
		require.Equal(t, head.InfoKindVersion, "1.0_1")

		require.Equal(t, head.Headline.Text, "１１日１４時４６分ころ、地震による強い揺れを感じました。震度３以上が観測された地域をお知らせします。")

		// Information
		require.Len(t, head.Headline.Information, 1)
		require.Equal(t, head.Headline.Information[0].Type, "震度速報")

		require.Len(t, head.Headline.Information[0].Item, 6)
		require.Equal(t, head.Headline.Information[0].Item[0].Kind.Name, "震度６強")

		require.Equal(t, head.Headline.Information[0].Item[0].Areas.CodeType, "地震情報／細分区域")
		require.Len(t, head.Headline.Information[0].Item[0].Areas.Area, 2)
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Name, "宮城県北部")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Code, "220")
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		require.Nil(t, body.Earthquake)

		t.Run("intensity", func(t *testing.T) {
			intensity := body.Intensity

			require.NotNil(t, intensity)

			require.Len(t, intensity.Observation.CodeDefine.Types, 2)
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].Value, "地震情報／都道府県等")
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].XPath, "Pref/Code")

			require.Equal(t, intensity.Observation.MaxInt, jma.Int6u)

			require.Len(t, intensity.Observation.Pref, 6)
			require.Equal(t, intensity.Observation.Pref[0].Name, "宮城県")
			require.Equal(t, intensity.Observation.Pref[0].Code, "04")
			require.Equal(t, intensity.Observation.Pref[0].MaxInt, jma.Int6u)

			require.Len(t, intensity.Observation.Pref[0].Area, 3)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Name, "宮城県北部")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Code, "220")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].MaxInt, jma.Int6u)
		})

		require.Equal(t, body.Text, "")

		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.ForecastComment.Text, "今後の情報に注意してください。")
		require.Equal(t, body.Comments.ForecastComment.Code, "0217")
		require.Equal(t, body.Comments.ForecastComment.CodeType, "固定付加文")
		require.Equal(t, body.Comments.FreeFormComment, "")
		require.Nil(t, body.Comments.VarComment)

	})
}

// 震度速報パーステスト（取り消し）
func TestParseEarthquakeReport3(t *testing.T) {
	target := "32-35_10_01_220510_VXSE51.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震度速報",
			DateTime:         "2022-05-10T09:10:20Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "震度速報")
		require.Equal(t, head.ReportDateTime, "2022-05-10T18:10:00+09:00")
		require.Equal(t, head.TargetDateTime, "2022-05-09T15:23:00+09:00")
		require.Equal(t, head.EventID, "20220509152315")
		require.Equal(t, head.InfoType, jma.Cancel)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震度速報")
		require.Equal(t, head.InfoKindVersion, "1.0_1")

		require.Equal(t, head.Headline.Text, "震度速報を取り消します。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		require.Nil(t, body.Earthquake)
		require.Nil(t, body.Intensity)

		require.Equal(t, body.Text, "先ほどの、震度速報を取り消します。")
		require.Nil(t, body.Comments)

	})
}

// 震源に関する情報パーステスト
func TestParseEarthquakeEpicenter1(t *testing.T) {
	target := "32-35_01_02_100514_VXSE52.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震源に関する情報",
			DateTime:         "2009-10-01T04:48:03Z",
			Status:           jma.Training,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "震源に関する情報")
		require.Equal(t, head.ReportDateTime, "2009-10-01T13:48:00+09:00")
		require.Equal(t, head.TargetDateTime, "2009-10-01T13:48:00+09:00")
		require.Equal(t, head.EventID, "20091001134500")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震源速報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "　１日１３時４５分ころ、地震がありました。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		t.Run("earthquake", func(t *testing.T) {
			earthquake := body.Earthquake
			require.NotNil(t, earthquake)

			require.Equal(t, earthquake.OriginTime, "2009-10-01T13:45:00+09:00")
			require.Equal(t, earthquake.ArrivalTime, "2009-10-01T13:45:00+09:00")

			require.Equal(t, earthquake.Hypocenter.Area.Name, "駿河湾")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Value, "485")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Datum, "日本測地系")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Description, "北緯３４．８度　東経１３８．５度　深さ　１０ｋｍ")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Value, "+34.8+138.5-10000/")
			require.Equal(t, earthquake.Hypocenter.Area.NameFromMark, "")
			require.Nil(t, earthquake.Hypocenter.Area.MarkCode)
			require.Equal(t, earthquake.Hypocenter.Area.Direction, "")
			require.Nil(t, earthquake.Hypocenter.Area.Distance)

			require.Equal(t, earthquake.Hypocenter.Source, "")

			require.Equal(t, earthquake.Magnitude.Condition, "")
			require.Equal(t, earthquake.Magnitude.Type, jma.JMAMagnitude)
			require.Equal(t, earthquake.Magnitude.Description, "Ｍ６．６")
			require.Equal(t, earthquake.Magnitude.Value, "6.6")

		})

		require.Nil(t, body.Intensity)

		require.Equal(t, body.Text, "")

		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.ForecastComment.Text, "この地震による津波の心配はありません。")
		require.Equal(t, body.Comments.ForecastComment.Code, "0203")
		require.Equal(t, body.Comments.ForecastComment.CodeType, "固定付加文")
		require.Equal(t, body.Comments.FreeFormComment, "")
		require.Nil(t, body.Comments.VarComment)
	})
}

// 震源に関する情報パーステスト
func TestParseEarthquakeEpicenter2(t *testing.T) {
	target := "33_12_01_120615_VXSE52.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震源に関する情報",
			DateTime:         "2012-02-14T12:42:53Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "震源に関する情報")
		require.Equal(t, head.ReportDateTime, "2012-02-14T21:42:00+09:00")
		require.Equal(t, head.TargetDateTime, "2012-02-14T21:42:00+09:00")
		require.Equal(t, head.EventID, "20120214214013")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震源速報")
		require.Equal(t, head.InfoKindVersion, "1.0_1")

		require.Equal(t, head.Headline.Text, "１４日２１時４０分ころ、地震がありました。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		t.Run("earthquake", func(t *testing.T) {
			earthquake := body.Earthquake
			require.NotNil(t, earthquake)

			require.Equal(t, earthquake.OriginTime, "2012-02-14T21:40:00+09:00")
			require.Equal(t, earthquake.ArrivalTime, "2012-02-14T21:40:00+09:00")

			require.Equal(t, earthquake.Hypocenter.Area.Name, "岐阜県美濃中西部")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Value, "432")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Datum, "日本測地系")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Description, "北緯３５．６度　東経１３６．６度　深さ　１０ｋｍ")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Value, "+35.6+136.6-10000/")
			require.Equal(t, earthquake.Hypocenter.Area.NameFromMark, "")
			require.Nil(t, earthquake.Hypocenter.Area.MarkCode)
			require.Equal(t, earthquake.Hypocenter.Area.Direction, "")
			require.Nil(t, earthquake.Hypocenter.Area.Distance)

			require.Equal(t, earthquake.Hypocenter.Source, "")

			require.Equal(t, earthquake.Magnitude.Condition, "不明")
			require.Equal(t, earthquake.Magnitude.Type, jma.JMAMagnitude)
			require.Equal(t, earthquake.Magnitude.Description, "Ｍ８を超える巨大地震")
			require.Equal(t, earthquake.Magnitude.Value, "NaN")

		})

		require.Nil(t, body.Intensity)

		require.Equal(t, body.Text, "")

		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.ForecastComment.Text, `この地震による津波の心配はありません。
この地震について、緊急地震速報を発表しています。`)
		require.Equal(t, body.Comments.ForecastComment.Code, "0215 0241")
		require.Equal(t, body.Comments.ForecastComment.CodeType, "固定付加文")
		require.Equal(t, body.Comments.FreeFormComment, "")
		require.Nil(t, body.Comments.VarComment)
	})
}

// 震源・震度に関する情報のパーステスト
func TestParseEarthquakeDetails1(t *testing.T) {
	target := "32-35_01_03_100806_VXSE53.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震源・震度に関する情報",
			DateTime:         "2009-10-01T04:50:01Z",
			Status:           jma.Training,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "震源・震度情報")
		require.Equal(t, head.ReportDateTime, "2009-10-01T13:50:00+09:00")
		require.Equal(t, head.TargetDateTime, "2009-10-01T13:50:00+09:00")
		require.Equal(t, head.EventID, "20091001134500")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "1")
		require.Equal(t, head.InfoKind, "地震情報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "　１日１３時４５分ころ、地震がありました。各地の震度をお知らせします。")

		// Information
		require.Len(t, head.Headline.Information, 2)
		require.Equal(t, head.Headline.Information[0].Type, "震源・震度に関する情報（細分区域）")
		require.Len(t, head.Headline.Information[0].Item, 3)
		require.Equal(t, head.Headline.Information[0].Item[0].Kind.Name, "震度５弱")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.CodeType, "地震情報／細分区域")
		require.Len(t, head.Headline.Information[0].Item[0].Areas.Area, 2)
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Name, "静岡県伊豆")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Code, "440")
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		t.Run("earthquake", func(t *testing.T) {
			earthquake := body.Earthquake
			require.NotNil(t, earthquake)

			require.Equal(t, earthquake.OriginTime, "2009-10-01T13:45:00+09:00")
			require.Equal(t, earthquake.ArrivalTime, "2009-10-01T13:45:00+09:00")

			require.Equal(t, earthquake.Hypocenter.Area.Name, "駿河湾")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Value, "485")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Datum, "日本測地系")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Description, "北緯３４．８度　東経１３８．５度　深さ　１０ｋｍ")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Value, "+34.8+138.5-10000/")
			require.Equal(t, earthquake.Hypocenter.Area.NameFromMark, "")
			require.Nil(t, earthquake.Hypocenter.Area.MarkCode)
			require.Equal(t, earthquake.Hypocenter.Area.Direction, "")
			require.Nil(t, earthquake.Hypocenter.Area.Distance)

			require.Equal(t, earthquake.Hypocenter.Source, "")

			require.Equal(t, earthquake.Magnitude.Condition, "")
			require.Equal(t, earthquake.Magnitude.Type, jma.JMAMagnitude)
			require.Equal(t, earthquake.Magnitude.Description, "Ｍ５．９")
			require.Equal(t, earthquake.Magnitude.Value, "5.9")

		})

		t.Run("intensity", func(t *testing.T) {
			intensity := body.Intensity

			require.NotNil(t, intensity)

			require.Len(t, intensity.Observation.CodeDefine.Types, 4)
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].Value, "地震情報／都道府県等")
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].XPath, "Pref/Code")

			require.Equal(t, intensity.Observation.MaxInt, jma.Int5l)

			require.Len(t, intensity.Observation.Pref, 8)
			require.Equal(t, intensity.Observation.Pref[0].Name, "静岡県")
			require.Equal(t, intensity.Observation.Pref[0].Code, "22")
			require.Equal(t, intensity.Observation.Pref[0].MaxInt, jma.Int5l)

			require.Len(t, intensity.Observation.Pref[0].Area, 4)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Name, "静岡県伊豆")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Code, "440")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].MaxInt, jma.Int5l)

			require.Len(t, intensity.Observation.Pref[0].Area[0].City, 10)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].Name, "西伊豆町")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].Code, "2230600")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].MaxInt, jma.Int5l)
			require.Len(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation, 1)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation[0].Name, "西伊豆町仁科＊")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation[0].Code, "2230630")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation[0].Int, jma.Int5l)

		})

		require.Equal(t, body.Text, "")

		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.ForecastComment.Text, "この地震による津波の心配はありません。")
		require.Equal(t, body.Comments.ForecastComment.Code, "0203")
		require.Equal(t, body.Comments.ForecastComment.CodeType, "固定付加文")
		require.Equal(t, body.Comments.FreeFormComment, "")
		require.Equal(t, body.Comments.VarComment.Text, "＊印は気象庁以外の震度観測点についての情報です。")
		require.Equal(t, body.Comments.VarComment.Code, "0210")
		require.Equal(t, body.Comments.VarComment.CodeType, "固定付加文")
	})
}

// 震源・震度に関する情報のパーステスト（東日本大震災）
func TestParseEarthquakeDetails2(t *testing.T) {
	target := "32-39_11_05_120615_VXSE53.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震源・震度に関する情報",
			DateTime:         "2011-03-11T05:54:58Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "震源・震度情報")
		require.Equal(t, head.ReportDateTime, "2011-03-11T14:54:00+09:00")
		require.Equal(t, head.TargetDateTime, "2011-03-11T14:54:00+09:00")
		require.Equal(t, head.EventID, "20110311144640")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "1")
		require.Equal(t, head.InfoKind, "地震情報")
		require.Equal(t, head.InfoKindVersion, "1.0_1")

		require.Equal(t, head.Headline.Text, "１１日１４時４６分ころ、地震がありました。")

		// Information
		require.Len(t, head.Headline.Information, 2)
		require.Equal(t, head.Headline.Information[0].Type, "震源・震度に関する情報（細分区域）")
		require.Len(t, head.Headline.Information[0].Item, 7)
		require.Equal(t, head.Headline.Information[0].Item[0].Kind.Name, "震度７")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.CodeType, "地震情報／細分区域")
		require.Len(t, head.Headline.Information[0].Item[0].Areas.Area, 1)
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Name, "宮城県北部")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Code, "220")
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		t.Run("earthquake", func(t *testing.T) {
			earthquake := body.Earthquake
			require.NotNil(t, earthquake)

			require.Equal(t, earthquake.OriginTime, "2011-03-11T14:46:00+09:00")
			require.Equal(t, earthquake.ArrivalTime, "2011-03-11T14:46:00+09:00")

			require.Equal(t, earthquake.Hypocenter.Area.Name, "三陸沖")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Value, "288")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Datum, "日本測地系")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Description, "北緯３８．０度　東経１４２．９度　深さ　１０ｋｍ")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Value, "+38.0+142.9-10000/")
			require.Equal(t, earthquake.Hypocenter.Area.NameFromMark, "牡鹿半島の東南東１３０ｋｍ付近")
			require.Equal(t, earthquake.Hypocenter.Area.MarkCode.Type, "震央補助")
			require.Equal(t, earthquake.Hypocenter.Area.MarkCode.Value, "202")
			require.Equal(t, earthquake.Hypocenter.Area.Direction, "東南東")
			require.Equal(t, earthquake.Hypocenter.Area.Distance.Unit, "km")
			require.Equal(t, earthquake.Hypocenter.Area.Distance.Value, "130")

			require.Equal(t, earthquake.Hypocenter.Source, "")

			require.Equal(t, earthquake.Magnitude.Condition, "不明")
			require.Equal(t, earthquake.Magnitude.Type, jma.JMAMagnitude)
			require.Equal(t, earthquake.Magnitude.Description, "Ｍ８を超える巨大地震")
			require.Equal(t, earthquake.Magnitude.Value, "NaN")

		})

		t.Run("intensity", func(t *testing.T) {
			intensity := body.Intensity

			require.NotNil(t, intensity)

			require.Len(t, intensity.Observation.CodeDefine.Types, 4)
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].Value, "地震情報／都道府県等")
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].XPath, "Pref/Code")

			require.Equal(t, intensity.Observation.MaxInt, jma.Int7)

			require.Len(t, intensity.Observation.Pref, 41)
			require.Equal(t, intensity.Observation.Pref[0].Name, "宮城県")
			require.Equal(t, intensity.Observation.Pref[0].Code, "04")
			require.Equal(t, intensity.Observation.Pref[0].MaxInt, jma.Int7)

			require.Len(t, intensity.Observation.Pref[0].Area, 3)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Name, "宮城県北部")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Code, "220")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].MaxInt, jma.Int7)

			require.Len(t, intensity.Observation.Pref[0].Area[0].City, 8)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].Name, "栗原市")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].Code, "0421300")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].MaxInt, jma.Int7)
			require.Len(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation, 2)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation[0].Name, "栗原市築館＊")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation[0].Code, "0421320")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation[0].Int, jma.Int7)

		})

		require.Equal(t, body.Text, "")

		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.ForecastComment.Text, `津波警報等（大津波警報・津波警報あるいは津波注意報）を発表中です。
この地震について、緊急地震速報を発表しています。`)
		require.Equal(t, body.Comments.ForecastComment.Code, "0211 0241")
		require.Equal(t, body.Comments.ForecastComment.CodeType, "固定付加文")
		require.Equal(t, body.Comments.FreeFormComment, "")
		require.Equal(t, body.Comments.VarComment.Text, "＊印は気象庁以外の震度観測点についての情報です。")
		require.Equal(t, body.Comments.VarComment.Code, "0262")
		require.Equal(t, body.Comments.VarComment.CodeType, "固定付加文")
	})
}

// 震源・震度に関する情報のパーステスト（遠地地震）
func TestParseEarthquakeDetails3(t *testing.T) {
	target := "32-35_01_03_100514_VXSE53.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震源・震度に関する情報",
			DateTime:         "2009-11-09T11:15:00Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "遠地地震に関する情報")
		require.Equal(t, head.ReportDateTime, "2009-11-09T20:15:00+09:00")
		require.Equal(t, head.TargetDateTime, "2009-11-09T20:15:00+09:00")
		require.Equal(t, head.EventID, "20091109194822")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "1")
		require.Equal(t, head.InfoKind, "地震情報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "　９日１９時４５分ころ、海外で規模の大きな地震がありました。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		t.Run("earthquake", func(t *testing.T) {
			earthquake := body.Earthquake
			require.NotNil(t, earthquake)

			require.Equal(t, earthquake.OriginTime, "2009-11-09T19:45:00+09:00")
			require.Equal(t, earthquake.ArrivalTime, "2009-11-09T19:45:00+09:00")

			require.Equal(t, earthquake.Hypocenter.Area.Name, "南太平洋")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Value, "950")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Datum, "")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Description, "南緯１７．２度　東経１７８．６度　深さ５７０ｋｍ")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Value, "-17.2+178.6-570000/")
			require.Equal(t, earthquake.Hypocenter.Area.NameFromMark, "")
			require.Nil(t, earthquake.Hypocenter.Area.MarkCode)
			require.Equal(t, earthquake.Hypocenter.Area.Direction, "")
			require.Nil(t, earthquake.Hypocenter.Area.Distance)
			require.Equal(t, earthquake.Hypocenter.Area.DetailedName, "フィジー諸島")
			require.Equal(t, earthquake.Hypocenter.Area.DetailedCode.Type, "詳細震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.DetailedCode.Value, "182")

			require.Equal(t, earthquake.Hypocenter.Source, "ＰＴＷＣ")

			require.Equal(t, earthquake.Magnitude.Condition, "")
			require.Equal(t, earthquake.Magnitude.Type, jma.MomentMagnitude)
			require.Equal(t, earthquake.Magnitude.Description, "Ｍ７．１")
			require.Equal(t, earthquake.Magnitude.Value, "7.1")

		})

		require.Nil(t, body.Intensity)

		require.Equal(t, body.Text, "")

		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.ForecastComment.Text, "この地震による津波の心配はありません。")
		require.Equal(t, body.Comments.ForecastComment.Code, "0203")
		require.Equal(t, body.Comments.ForecastComment.CodeType, "固定付加文")
		require.Equal(t, body.Comments.FreeFormComment, "")
		require.Nil(t, body.Comments.VarComment)
	})
}
