package src_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/stretchr/testify/require"
)

const TEST_DATA_PATH = "../test_data/jma_xml/"

// 津波情報のテストデータ
var TestData = []string{
	"32-39_11_02_120615_VTSE41.xml",
	"32-39_11_03_120615_VTSE51.xml",
	"32-39_11_06_120615_VTSE41.xml",
	"32-39_11_08_120615_VTSE51.xml",
	"32-39_11_09_120615_VTSE41.xml",
	"32-39_11_10_120615_VTSE51.xml",
	"32-39_11_11_120615_VTSE41.xml",
	"32-39_11_13_120615_VTSE41.xml",

	"32-39_12_02_191025_VTSE41.xml",
	"32-39_12_03_191025_VTSE51.xml",
	"32-39_12_05_191025_VTSE52.xml",
	"32-39_12_06_191025_VTSE51.xml",
	"32-39_12_07_191025_VTSE41.xml",
	"32-39_12_08_191025_VTSE51.xml",
	"32-39_12_10_191025_VTSE51.xml",
	"32-39_12_11_191025_VTSE51.xml",
	"32-39_12_12_191025_VTSE51.xml",
	"32-39_12_13_191025_VTSE41.xml",
	"32-39_12_14_191025_VTSE51.xml",
	"32-39_13_02_191025_VTSE41.xml",
	"32-39_13_03_191025_VTSE51.xml",
	"32-39_13_07_191025_VTSE41.xml",

	"38-39_02_01_191025_VTSE41.xml",
	"38-39_02_02_191025_VTSE51.xml",
	"38-39_02_03_191025_VTSE51.xml",
	"38-39_02_04_191025_VTSE41.xml",
	"38-39_02_05_191025_VTSE51.xml",
	"38-39_03_01_210805_VTSE41.xml",
	"38-39_03_02_210805_VTSE51.xml",
	"38-39_03_03_210805_VTSE51.xml",
}

func TestParseTsunami(t *testing.T) {
	for _, d := range TestData {
		t.Run(fmt.Sprintf("Test %s", d), func(t *testing.T) {
			testPath := filepath.Join(TEST_DATA_PATH, d)

			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			_, err = src.ParseTsunami(row)
			require.NoError(t, err)
		})
	}

	t.Run("failed", func(t *testing.T) {
		row := "aaaaaaaa"

		_, err := src.ParseTsunami([]byte(row))
		require.Error(t, err)
	})
}

// 津波警報・注意報・予報 のパーステスト
func TestParseTsunamiWarning(t *testing.T) {
	target := "32-39_11_02_120615_VTSE41.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	tsunami, err := src.ParseTsunami(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := tsunami.Parsed.Control

		require.Equal(t, control, src.JmaXmlControl{
			Title:            "津波警報・注意報・予報a",
			DateTime:         "2011-03-11T05:49:59Z",
			Status:           "通常",
			EditorialOffice:  "大阪管区気象台",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := tsunami.Parsed.Head

		require.Equal(t, head.Title, "大津波警報・津波警報・津波注意報・津波予報")
		require.Equal(t, head.ReportDateTime, "2011-03-11T14:49:00+09:00")
		require.Equal(t, head.TargetDateTime, "2011-03-11T14:49:00+09:00")
		require.Equal(t, head.EventID, "20110311144640")
		require.Equal(t, head.InfoType, "発表")
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "津波警報・注意報・予報")
		require.Equal(t, head.InfoKindVersion, "1.0_1")

		require.Equal(t, head.Headline.Text, `東日本大震災クラスの津波が来襲します。
大津波警報・津波警報を発表しました。
ただちに避難してください。`)

		// Information
		require.Len(t, head.Headline.Information, 1)
		require.Equal(t, head.Headline.Information[0].Type, "津波予報領域表現")

		require.Len(t, head.Headline.Information[0].Item, 2)
		require.Equal(t, head.Headline.Information[0].Item[0].Kind.Name, "大津波警報")
		require.Equal(t, head.Headline.Information[0].Item[0].Kind.Code, "52")

		require.Equal(t, head.Headline.Information[0].Item[0].Areas.CodeType, "津波予報区")
		require.Len(t, head.Headline.Information[0].Item[0].Areas.Area, 1)
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Name, "東北地方太平洋沿岸")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Code, "291")
	})

	t.Run("body", func(t *testing.T) {
		body := tsunami.Parsed.Body

		t.Run("tsunami", func(t *testing.T) {
			bodyTsunami := body.Tsunami

			// Observation は津波情報なのでnil
			require.Nil(t, bodyTsunami.Observation)

			// CodeDefine
			require.Len(t, bodyTsunami.Forecast.CodeDefine.Type, 3)
			require.Equal(t, bodyTsunami.Forecast.CodeDefine.Type[0].XPath, "Item/Area/Code")
			require.Equal(t, bodyTsunami.Forecast.CodeDefine.Type[0].Value, "津波予報区")

			// Item
			require.Len(t, bodyTsunami.Forecast.Item, 43)
			require.Equal(t, bodyTsunami.Forecast.Item[0].Area.Name, "岩手県")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Area.Code, "210")

			require.Equal(t, bodyTsunami.Forecast.Item[0].Category.Kind.Name, "大津波警報：発表")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Category.Kind.Code, "53")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Category.LastKind.Name, "津波なし")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Category.LastKind.Code, "00")

			require.Equal(t, bodyTsunami.Forecast.Item[0].FirstHeight.Condition, "津波到達中と推測")
			require.Equal(t, bodyTsunami.Forecast.Item[0].FirstHeight.ArrivalTime, "")
			require.Equal(t, bodyTsunami.Forecast.Item[0].FirstHeight.Revise, "")

			require.NotNil(t, bodyTsunami.Forecast.Item[0].MaxHeight)
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.Condition, "")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.Revise, "")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Value, "NaN")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Type, "津波の高さ")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Unit, "m")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Description, "巨大")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Condition, "不明")

			// Station は津波情報なのでnil
			require.Nil(t, bodyTsunami.Forecast.Item[0].Station)
		})

		t.Run("earthquake", func(t *testing.T) {
			bodyEarthquake := body.Earthquake

			require.Len(t, bodyEarthquake, 1)
			require.Equal(t, bodyEarthquake[0].OriginTime, "2011-03-11T14:46:00+09:00")
			require.Equal(t, bodyEarthquake[0].ArrivalTime, "2011-03-11T14:46:00+09:00")

			// Hypocenter
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Name, "三陸沖")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Code.Value, "288")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Coordinate.Value, "+38.0+142.9-10000/")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Coordinate.Description, "北緯３８．０度　東経１４２．９度　深さ　１０ｋｍ")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Coordinate.Datum, "日本測地系")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.NameFromMark, "牡鹿半島の東南東１３０ｋｍ付近")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.MarkCode.Type, "震央補助")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.MarkCode.Value, "202")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Direction, "東南東")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Distance.Value, "130")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Distance.Unit, "km")

			require.Equal(t, bodyEarthquake[0].Hypocenter.Source, "")

			require.Equal(t, bodyEarthquake[0].Magnitude.Value, "NaN")
			require.Equal(t, bodyEarthquake[0].Magnitude.Type, "Mj")
			require.Equal(t, bodyEarthquake[0].Magnitude.Description, "Ｍ８を超える巨大地震")
			require.Equal(t, bodyEarthquake[0].Magnitude.Condition, "不明")
		})

		t.Run("comments", func(t *testing.T) {
			bodyComments := body.Comments

			require.Equal(t, bodyComments.WarningComment.Text, `東日本大震災クラスの津波が来襲します。

ただちに避難してください。

＜大津波警報＞
大きな津波が襲い甚大な被害が発生します。
沿岸部や川沿いにいる人はただちに高台や避難ビルなど安全な場所へ避難してください。
津波は繰り返し襲ってきます。警報が解除されるまで安全な場所から離れないでください。

＜津波警報＞
津波による被害が発生します。
沿岸部や川沿いにいる人はただちに高台や避難ビルなど安全な場所へ避難してください。
津波は繰り返し襲ってきます。警報が解除されるまで安全な場所から離れないでください。

＜津波注意報＞
海の中や海岸付近は危険です。
海の中にいる人はただちに海から上がって、海岸から離れてください。
潮の流れが速い状態が続きますので、注意報が解除されるまで海に入ったり海岸に近づいたりしないようにしてください。

＜津波予報（若干の海面変動）＞
若干の海面変動が予想されますが、被害の心配はありません。

警報が発表された沿岸部や川沿いにいる人はただちに高台や避難ビルなど安全な場所へ避難してください。
到達予想時刻は、予報区のなかで最も早く津波が到達する時刻です。場所によっては、この時刻よりもかなり遅れて津波が襲ってくることがあります。
到達予想時刻から津波が最も高くなるまでに数時間以上かかることがありますので、観測された津波の高さにかかわらず、警報が解除されるまで安全な場所から離れないでください。`)
			require.Equal(t, bodyComments.WarningComment.CodeType, "固定付加文")
			require.Equal(t, bodyComments.WarningComment.Code, "0141 0149 0121 0122 0123 0124 0131")

			require.Equal(t, bodyComments.FreeFormComment, "")
		})

		t.Run("other", func(t *testing.T) {
			require.Equal(t, body.Text, "")
		})
	})
}

// // 津波情報 のパーステスト
// func TestParseTsunamiInfo(t *testing.T) {
// 	target := "32-39_11_03_120615_VTSE51.xml"
// }

// // 沖合の津波観測に関する情報 のパーステスト
// func TestParseTsunamiOffshoreInfo(t *testing.T) {
// 	target := "32-39_12_05_191025_VTSE52.xml"
// }

func TestStatus(t *testing.T) {
	statusData := map[string]src.Status{
		"32-39_11_02_120615_VTSE41.xml": src.Common,
		"32-39_12_05_191025_VTSE52.xml": src.Training,
	}

	for d, s := range statusData {
		t.Run(fmt.Sprintf("Test IsStatus %s", d), func(t *testing.T) {
			testPath := filepath.Join(TEST_DATA_PATH, d)

			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			tsunami, err := src.ParseTsunami(row)
			require.NoError(t, err)

			require.Equal(t, tsunami.Status(), s)
		})
	}
}

func TestIsCommon(t *testing.T) {
	statusData := map[string]bool{
		"32-39_11_02_120615_VTSE41.xml": true,
		"32-39_12_05_191025_VTSE52.xml": false,
	}

	for d, s := range statusData {
		t.Run(fmt.Sprintf("Test IsStatus %s", d), func(t *testing.T) {
			testPath := filepath.Join(TEST_DATA_PATH, d)

			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			tsunami, err := src.ParseTsunami(row)
			require.NoError(t, err)

			require.Equal(t, tsunami.IsCommon(), s)
		})
	}
}

func TestInfoType(t *testing.T) {
	statusData := map[string]src.InfoType{
		"32-35_01_01_100806_VXSE51.xml": src.Publication,
		"38-39_03_01_210805_VTSE41.xml": src.Cancel,
	}

	for d, s := range statusData {
		t.Run(fmt.Sprintf("Test IsStatus %s", d), func(t *testing.T) {
			testPath := filepath.Join(TEST_DATA_PATH, d)

			row, err := os.ReadFile(testPath)
			require.NoError(t, err)

			tsunami, err := src.ParseTsunami(row)
			require.NoError(t, err)

			require.Equal(t, tsunami.InfoType(), s)
		})
	}
}
