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

var TestEarthquakeUpdate = []string{
	"32-35_06_09_100915_VXSE61.xml",
	"32-35_06_10_100915_VXSE61.xml",
	"32-35_07_09_100915_VXSE61.xml",
}

func TestParseEarthquakeUpdate(t *testing.T) {
	for _, d := range TestEarthquakeUpdate {
		t.Run(fmt.Sprintf("Test %s", d), func(t *testing.T) {
			testPath := filepath.Join(TEST_DATA_PATH, d)

			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			_, err = src.ParseEarthquakeUpdate(row)
			require.NoError(t, err)
		})
	}

	t.Run("failed", func(t *testing.T) {
		row := "aaaaaaaa"

		_, err := src.ParseEarthquakeUpdate([]byte(row))
		require.Error(t, err)
	})
}

// パーステスト1
func TestParseEarthquakeUpdate1(t *testing.T) {
	target := "32-35_06_09_100915_VXSE61.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquakeUpdate(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "顕著な地震の震源要素更新のお知らせ",
			DateTime:         "2008-06-14T03:30:00Z",
			Status:           "通常",
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "顕著な地震の震源要素更新のお知らせ")
		require.Equal(t, head.ReportDateTime, "2008-06-14T12:30:00+09:00")
		require.Equal(t, head.TargetDateTime, "2008-06-14T12:30:00+09:00")
		require.Equal(t, head.EventID, "20080614084350")
		require.Equal(t, head.InfoType, "発表")
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震源要素更新のお知らせ")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "平成２０年　６月１４日１２時３０分をもって、地震の発生場所と規模を更新します。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		t.Run("earthquake", func(t *testing.T) {
			earthquake := body.Earthquake

			require.Equal(t, earthquake.OriginTime, "2008-06-14T08:43:00+09:00")
			require.Equal(t, earthquake.ArrivalTime, "2008-06-14T08:43:00+09:00")

			require.Equal(t, earthquake.Hypocenter.Area.Name, "岩手県内陸南部")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Value, "213")
			require.Len(t, earthquake.Hypocenter.Area.Coordinate, 2)
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Value, "+39.0+140.9-10000/")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Description, "北緯３９．０度　東経１４０．９度　深さ　１０ｋｍ")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Datum, "日本測地系")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Type, "")

			require.Equal(t, earthquake.Magnitude.Condition, "")
			require.Equal(t, earthquake.Magnitude.Type, "Mj")
			require.Equal(t, earthquake.Magnitude.Description, "Ｍ７．２")
			require.Equal(t, earthquake.Magnitude.Value, "7.2")
		})

		require.Equal(t, body.Text, "")
		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.FreeFormComment, "度単位の震源要素は、津波情報等を引き続き発表する場合に使用されます。")
	})
}

// パーステスト2 （取り消し）
func TestParseEarthquakeUpdate2(t *testing.T) {
	target := "32-35_06_10_100915_VXSE61.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquakeUpdate(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "顕著な地震の震源要素更新のお知らせ",
			DateTime:         "2008-06-14T03:35:00Z",
			Status:           "通常",
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "顕著な地震の震源要素更新のお知らせ")
		require.Equal(t, head.ReportDateTime, "2008-06-14T12:35:00+09:00")
		require.Equal(t, head.TargetDateTime, "2008-06-14T12:35:00+09:00")
		require.Equal(t, head.EventID, "20080614084350")
		require.Equal(t, head.InfoType, "取消")
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震源要素更新のお知らせ")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "顕著な地震の震源要素更新のお知らせを取り消します。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		require.Nil(t, body.Earthquake)
		require.Equal(t, body.Text, "先ほどの、顕著な地震の震源要素更新のお知らせを取り消します。")
		require.Nil(t, body.Comments)
	})
}

// パーステスト3
func TestParseEarthquakeUpdate3(t *testing.T) {
	target := "32-35_07_09_100915_VXSE61.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquakeUpdate(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "顕著な地震の震源要素更新のお知らせ",
			DateTime:         "2009-08-10T23:55:00Z",
			Status:           "通常",
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "顕著な地震の震源要素更新のお知らせ")
		require.Equal(t, head.ReportDateTime, "2009-08-11T08:55:00+09:00")
		require.Equal(t, head.TargetDateTime, "2009-08-11T06:45:00+09:00")
		require.Equal(t, head.EventID, "20090811050711")
		require.Equal(t, head.InfoType, "発表")
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震源要素更新のお知らせ")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "平成２１年　８月１１日０６時４５分をもって、地震の発生場所と規模を更新します。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		t.Run("earthquake", func(t *testing.T) {
			earthquake := body.Earthquake

			require.Equal(t, earthquake.OriginTime, "2009-08-11T05:07:00+09:00")
			require.Equal(t, earthquake.ArrivalTime, "2009-08-11T05:07:00+09:00")

			require.Equal(t, earthquake.Hypocenter.Area.Name, "駿河湾")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Value, "485")
			require.Len(t, earthquake.Hypocenter.Area.Coordinate, 2)
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Value, "+34.8+138.5-20000/")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Description, "北緯３４．８度　東経１３８．５度　深さ　２０ｋｍ")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Datum, "日本測地系")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Type, "")

			require.Equal(t, earthquake.Magnitude.Condition, "")
			require.Equal(t, earthquake.Magnitude.Type, "Mj")
			require.Equal(t, earthquake.Magnitude.Description, "Ｍ６．５")
			require.Equal(t, earthquake.Magnitude.Value, "6.5")
		})

		require.Equal(t, body.Text, "")
		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.FreeFormComment, "度単位の震源要素は、津波情報等を引き続き発表する場合に使用されます。")
	})
}
