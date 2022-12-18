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

var TestEarthquakeCount = []string{
	"32-35_03_01_100514_VXSE60.xml",
	"32-35_10_02_220510_VXSE60.xml",
}

func TestParseEarthquakeCount(t *testing.T) {
	for _, d := range TestEarthquakeCount {
		t.Run(fmt.Sprintf("Test %s", d), func(t *testing.T) {
			testPath := filepath.Join(TEST_DATA_PATH, d)

			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			_, err = src.ParseEarthquakeCount(row)
			require.NoError(t, err)
		})
	}

	t.Run("failed", func(t *testing.T) {
		row := "aaaaaaaa"

		_, err := src.ParseEarthquakeCount([]byte(row))
		require.Error(t, err)
	})
}

// パーステスト1
func TestParseEarthquakeCount1(t *testing.T) {
	target := "32-35_03_01_100514_VXSE60.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquakeCount(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "地震回数に関する情報",
			DateTime:         "2008-08-26T03:00:15Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "地震回数に関する情報")
		require.Equal(t, head.ReportDateTime, "2008-08-26T12:00:00+09:00")
		require.Equal(t, head.TargetDateTime, "2008-08-26T12:00:00+09:00")
		require.Equal(t, head.EventID, "20080824150500")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "1")
		require.Equal(t, head.InfoKind, "地震回数情報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "地震回数に関する情報をお知らせします。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		t.Run("earthquake count", func(t *testing.T) {
			earthquake := body.EarthquakeCount

			require.Len(t, earthquake.Item, 5)
			require.Equal(t, earthquake.Item[0].Type, "地震回数")
			require.Equal(t, earthquake.Item[0].StartTime, "2008-08-24T15:00:00+09:00")
			require.Equal(t, earthquake.Item[0].EndTime, "2008-08-25T09:00:00+09:00")
			require.Equal(t, earthquake.Item[0].Number, 1587)
			require.Equal(t, earthquake.Item[0].FeltNumber, 1)
		})

		require.Equal(t, body.Text, "")
		require.Equal(t, body.NextAdvisory, "次の「地震回数に関する情報」は、２６日１８時００分頃に発表します。")

		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.FreeFormComment, `　８月２４日１５時過ぎから伊豆半島東方沖で地震が発生しています。この
付近で発生した地震については、震度３以上の場合は「震源・震度情報」で
発表しますが、震度２以下の場合は、「地震回数に関する情報」（本情報）
で地震回数をまとめて発表します。`)
	})
}

// パーステスト2 （取り消し）
func TestParseEarthquakeCount2(t *testing.T) {
	target := "32-35_10_02_220510_VXSE60.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := src.ParseEarthquakeCount(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "地震回数に関する情報",
			DateTime:         "2022-05-10T09:49:07Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Parsed.Head

		require.Equal(t, head.Title, "地震回数に関する情報")
		require.Equal(t, head.ReportDateTime, "2022-05-10T18:49:00+09:00")
		require.Equal(t, head.TargetDateTime, "2022-05-10T18:49:00+09:00")
		require.Equal(t, head.EventID, "20220510173400")
		require.Equal(t, head.InfoType, jma.Cancel)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "地震回数情報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "地震回数に関する情報を取り消します。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Parsed.Body

		require.Nil(t, body.EarthquakeCount)

		require.Equal(t, body.Text, "先ほどの、地震回数に関する情報を取り消します。")
		require.Equal(t, body.NextAdvisory, "")

		require.Nil(t, body.Comments)
	})
}

func TestEarthquakeCountGetText(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		target := "32-35_03_01_100514_VXSE60.xml"

		testPath := filepath.Join(TEST_DATA_PATH, target)
		row, err := os.ReadFile(testPath)
		require.NoError(t, err)

		ea, err := src.ParseEarthquakeCount(row)
		require.NoError(t, err)

		text, err := ea.GetText()
		require.NoError(t, err)

		require.Equal(t, text, "地震回数に関する情報をお知らせします。\n  2008年8月24日15時ごろから2008年8月25日12時ごろまでの間に、1704回（うち有感地震1回）の地震が発生しています。\n次の「地震回数に関する情報」は、26日18時00分頃に発表します。\n 8月24日15時過ぎから伊豆半島東方沖で地震が発生しています。この付近で発生した地震については、震度3以上の場合は「震源・震度情報」で発表しますが、震度2以下の場合は、「地震回数に関する情報」(本情報)で地震回数をまとめて発表します。")
	})

	t.Run("2", func(t *testing.T) {
		target := "32-35_10_02_220510_VXSE60.xml"

		testPath := filepath.Join(TEST_DATA_PATH, target)
		row, err := os.ReadFile(testPath)
		require.NoError(t, err)

		ea, err := src.ParseEarthquakeCount(row)
		require.NoError(t, err)

		text, err := ea.GetText()
		require.NoError(t, err)

		require.Equal(t, text, "地震回数に関する情報を取り消します。\n先ほどの、地震回数に関する情報を取り消します。")
	})
}
