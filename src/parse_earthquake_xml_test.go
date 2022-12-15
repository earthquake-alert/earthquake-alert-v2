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
	"32-35_10_01_220510_VXSE51.xml",
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

func TestParseEarthquakeReport1(t *testing.T) {
	target := "32-35_01_01_100806_VXSE51.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	tsunami, err := src.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := tsunami.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震度速報",
			DateTime:         "2009-08-10T20:09:11Z",
			Status:           "通常",
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := tsunami.Parsed.Head

		require.Equal(t, head.Title, "震度速報")
		require.Equal(t, head.ReportDateTime, "2009-08-11T05:09:00+09:00")
		require.Equal(t, head.TargetDateTime, "2009-08-11T05:07:00+09:00")
		require.Equal(t, head.EventID, "20090811050711")
		require.Equal(t, head.InfoType, "発表")
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
		body := tsunami.Parsed.Body

		require.Nil(t, body.Earthquake)

		t.Run("intensity", func(t *testing.T) {
			intensity := body.Intensity

			require.NotNil(t, intensity)

			require.Len(t, intensity.Observation.CodeDefine.Types, 2)
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].Value, "地震情報／都道府県等")
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].XPath, "Pref/Code")

			require.Equal(t, intensity.Observation.MaxInt, "6-")

			require.Len(t, intensity.Observation.Pref, 16)
			require.Equal(t, intensity.Observation.Pref[0].Name, "静岡県")
			require.Equal(t, intensity.Observation.Pref[0].Code, "22")
			require.Equal(t, intensity.Observation.Pref[0].MaxInt, "6-")

			require.Len(t, intensity.Observation.Pref[0].Area, 4)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Name, "静岡県中部")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Code, "442")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].MaxInt, "6-")
		})

		require.Equal(t, body.Text, "")
		require.Nil(t, body.Comments)

	})
}
