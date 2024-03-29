package jma_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src"
	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
	"github.com/stretchr/testify/require"
)

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

// 津波警報・注意報・予報 のパーステスト
func TestParseTsunamiWarning(t *testing.T) {
	target := "32-39_11_02_120615_VTSE41.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	tsunami, err := jma.ParseTsunami(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := tsunami.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "津波警報・注意報・予報a",
			DateTime:         "2011-03-11T05:49:59Z",
			Status:           jma.Common,
			EditorialOffice:  "大阪管区気象台",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := tsunami.Head

		require.Equal(t, head.Title, "大津波警報・津波警報・津波注意報・津波予報")
		require.Equal(t, head.ReportDateTime, "2011-03-11T14:49:00+09:00")
		require.Equal(t, head.TargetDateTime, "2011-03-11T14:49:00+09:00")
		require.Equal(t, head.EventID, "20110311144640")
		require.Equal(t, head.InfoType, jma.Publication)
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
		body := tsunami.Body

		t.Run("tsunami", func(t *testing.T) {
			bodyTsunami := body.Tsunami

			// Observation は津波情報なのでnil
			require.Nil(t, bodyTsunami.Observation)

			require.Nil(t, bodyTsunami.Estimation)

			// CodeDefine
			require.Len(t, bodyTsunami.Forecast.CodeDefine.Types, 3)
			require.Equal(t, bodyTsunami.Forecast.CodeDefine.Types[0].XPath, "Item/Area/Code")
			require.Equal(t, bodyTsunami.Forecast.CodeDefine.Types[0].Value, "津波予報区")

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
			require.Len(t, bodyTsunami.Forecast.Item[0].Station, 0)
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
			require.Equal(t, bodyEarthquake[0].Magnitude.Type, jma.JMAMagnitude)
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

// 津波警報・注意報・予報 のパーステスト2
func TestParseTsunamiWarning2(t *testing.T) {
	target := "32-39_13_02_191025_VTSE41.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	tsunami, err := jma.ParseTsunami(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := tsunami.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "津波警報・注意報・予報a",
			DateTime:         "2018-12-04T09:03:00Z",
			Status:           jma.Training,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := tsunami.Head

		require.Equal(t, head.Title, "津波注意報・津波予報")
		require.Equal(t, head.ReportDateTime, "2018-12-04T18:03:00+09:00")
		require.Equal(t, head.TargetDateTime, "2018-12-04T18:03:00+09:00")
		require.Equal(t, head.EventID, "20181204180000")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "津波警報・注意報・予報")
		require.Equal(t, head.InfoKindVersion, "1.0_1")

		require.Equal(t, head.Headline.Text, `＊＊＊これは訓練です＊＊＊
津波注意報を発表しました。`)

		// Information
		require.Len(t, head.Headline.Information, 1)
		require.Equal(t, head.Headline.Information[0].Type, "津波予報領域表現")

		require.Len(t, head.Headline.Information[0].Item, 1)
		require.Equal(t, head.Headline.Information[0].Item[0].Kind.Name, "津波注意報")
		require.Equal(t, head.Headline.Information[0].Item[0].Kind.Code, "62")

		require.Equal(t, head.Headline.Information[0].Item[0].Areas.CodeType, "津波予報区")
		require.Len(t, head.Headline.Information[0].Item[0].Areas.Area, 3)
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Name, "伊豆諸島")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Code, "320")
	})

	t.Run("body", func(t *testing.T) {
		body := tsunami.Body

		t.Run("tsunami", func(t *testing.T) {
			bodyTsunami := body.Tsunami

			// Observation は津波情報なのでnil
			require.Nil(t, bodyTsunami.Observation)

			require.Nil(t, bodyTsunami.Estimation)

			// CodeDefine
			require.Len(t, bodyTsunami.Forecast.CodeDefine.Types, 3)
			require.Equal(t, bodyTsunami.Forecast.CodeDefine.Types[0].XPath, "Item/Area/Code")
			require.Equal(t, bodyTsunami.Forecast.CodeDefine.Types[0].Value, "津波予報区")

			// Item
			require.Len(t, bodyTsunami.Forecast.Item, 7)
			require.Equal(t, bodyTsunami.Forecast.Item[0].Area.Name, "千葉県内房")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Area.Code, "311")

			require.Equal(t, bodyTsunami.Forecast.Item[0].Category.Kind.Name, "津波注意報")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Category.Kind.Code, "62")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Category.LastKind.Name, "津波なし")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Category.LastKind.Code, "00")

			require.Equal(t, bodyTsunami.Forecast.Item[0].FirstHeight.Condition, "津波到達中と推測")
			require.Equal(t, bodyTsunami.Forecast.Item[0].FirstHeight.ArrivalTime, "")
			require.Equal(t, bodyTsunami.Forecast.Item[0].FirstHeight.Revise, "")

			require.NotNil(t, bodyTsunami.Forecast.Item[0].MaxHeight)
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.Condition, "")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.Revise, "")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Value, "1")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Type, "津波の高さ")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Unit, "m")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Description, "１ｍ")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Condition, "")

			// Station は津波情報なのでnil
			require.Len(t, bodyTsunami.Forecast.Item[0].Station, 0)
		})

		t.Run("earthquake", func(t *testing.T) {
			bodyEarthquake := body.Earthquake

			require.Len(t, bodyEarthquake, 1)
			require.Equal(t, bodyEarthquake[0].OriginTime, "2018-12-04T18:00:00+09:00")
			require.Equal(t, bodyEarthquake[0].ArrivalTime, "2018-12-04T18:00:00+09:00")

			// Hypocenter
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Name, "東京都２３区")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Code.Value, "350")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Coordinate.Value, "+35.6+139.7-30000/")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Coordinate.Description, "北緯３５．６度　東経１３９．７度　深さ　３０ｋｍ")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Coordinate.Datum, "日本測地系")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.NameFromMark, "館山の北北西７０ｋｍ付近")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.MarkCode.Type, "震央補助")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.MarkCode.Value, "301")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Direction, "北北西")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Distance.Value, "70")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Distance.Unit, "km")

			require.Equal(t, bodyEarthquake[0].Hypocenter.Source, "")

			require.Equal(t, bodyEarthquake[0].Magnitude.Value, "7.3")
			require.Equal(t, bodyEarthquake[0].Magnitude.Type, jma.JMAMagnitude)
			require.Equal(t, bodyEarthquake[0].Magnitude.Description, "Ｍ７．３")
			require.Equal(t, bodyEarthquake[0].Magnitude.Condition, "")
		})

		t.Run("comments", func(t *testing.T) {
			bodyComments := body.Comments

			require.Equal(t, bodyComments.WarningComment.Text, `＜津波注意報＞
海の中や海岸付近は危険です。
海の中にいる人はただちに海から上がって、海岸から離れてください。
潮の流れが速い状態が続きますので、注意報が解除されるまで海に入ったり海岸に近づいたりしないようにしてください。

＜津波予報（若干の海面変動）＞
若干の海面変動が予想されますが、被害の心配はありません。

場所によっては津波の高さが「予想される津波の高さ」より高くなる可能性があります。`)
			require.Equal(t, bodyComments.WarningComment.CodeType, "固定付加文")
			require.Equal(t, bodyComments.WarningComment.Code, "0123 0124 0132")

			require.Equal(t, bodyComments.FreeFormComment, `［予想される津波の高さの解説］
予想される津波が高いほど、より甚大な被害が生じます。
１０ｍ超　　巨大な津波が襲い壊滅的な被害が生じる。木造家屋が全壊・流失し、人は津波による流れに巻き込まれる。
１０ｍ　　　巨大な津波が襲い甚大な被害が生じる。木造家屋が全壊・流失し、人は津波による流れに巻き込まれる。
　５ｍ　　　津波が襲い甚大な被害が生じる。木造家屋が全壊・流失し、人は津波による流れに巻き込まれる。
　３ｍ　　　標高の低いところでは津波が襲い被害が生じる。木造家屋で浸水被害が発生し、人は津波による流れに巻き込まれる。
　１ｍ　　　海の中では人は速い流れに巻き込まれる。養殖いかだが流失し小型船舶が転覆する。`)
		})

		t.Run("other", func(t *testing.T) {
			require.Equal(t, body.Text, "")
		})
	})
}

// 津波情報 のパーステスト
func TestParseTsunamiInfo(t *testing.T) {
	target := "32-39_11_03_120615_VTSE51.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	tsunami, err := jma.ParseTsunami(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := tsunami.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "津波情報a",
			DateTime:         "2011-03-11T05:50:46Z",
			Status:           jma.Common,
			EditorialOffice:  "大阪管区気象台",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := tsunami.Head

		require.Equal(t, head.Title, "各地の満潮時刻・津波到達予想時刻に関する情報")
		require.Equal(t, head.ReportDateTime, "2011-03-11T14:50:00+09:00")
		require.Equal(t, head.TargetDateTime, "2011-03-11T14:50:00+09:00")
		require.Equal(t, head.EventID, "20110311144640")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "1")
		require.Equal(t, head.InfoKind, "津波情報")
		require.Equal(t, head.InfoKindVersion, "1.0_1")

		require.Equal(t, head.Headline.Text, "各地の満潮時刻と津波到達予想時刻をお知らせします。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := tsunami.Body

		t.Run("tsunami", func(t *testing.T) {
			bodyTsunami := body.Tsunami

			// Observation は津波情報なのでnil
			require.Nil(t, bodyTsunami.Observation)

			require.Nil(t, bodyTsunami.Estimation)

			// CodeDefine
			require.Len(t, bodyTsunami.Forecast.CodeDefine.Types, 4)
			require.Equal(t, bodyTsunami.Forecast.CodeDefine.Types[0].XPath, "Item/Area/Code")
			require.Equal(t, bodyTsunami.Forecast.CodeDefine.Types[0].Value, "津波予報区")

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

			require.Len(t, bodyTsunami.Forecast.Item[0].Station, 4)
			require.Equal(t, bodyTsunami.Forecast.Item[0].Station[0].Name, "宮古")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Station[0].Code, "21001")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Station[0].HighTideDateTime, "2011-03-11T19:43:00+09:00")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Station[0].FirstHeight.ArrivalTime, "2011-03-11T15:20:00+09:00")
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
			require.Equal(t, bodyEarthquake[0].Magnitude.Type, jma.JMAMagnitude)
			require.Equal(t, bodyEarthquake[0].Magnitude.Description, "Ｍ８を超える巨大地震")
			require.Equal(t, bodyEarthquake[0].Magnitude.Condition, "不明")
		})

		t.Run("comments", func(t *testing.T) {
			bodyComments := body.Comments

			require.Equal(t, bodyComments.WarningComment.Text, "津波と満潮が重なると、津波はより高くなりますので一層厳重な警戒が必要です。")
			require.Equal(t, bodyComments.WarningComment.CodeType, "固定付加文")
			require.Equal(t, bodyComments.WarningComment.Code, "0109")

			require.Equal(t, bodyComments.FreeFormComment, "")
		})

		t.Run("other", func(t *testing.T) {
			require.Equal(t, body.Text, "")
		})
	})
}

// 津波情報 のパーステスト2
func TestParseTsunamiInfo2(t *testing.T) {
	target := "32-39_11_08_120615_VTSE51.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	tsunami, err := jma.ParseTsunami(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := tsunami.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "津波情報a",
			DateTime:         "2011-03-11T05:59:27Z",
			Status:           jma.Common,
			EditorialOffice:  "大阪管区気象台",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := tsunami.Head

		require.Equal(t, head.Title, "津波観測に関する情報")
		require.Equal(t, head.ReportDateTime, "2011-03-11T14:59:00+09:00")
		require.Equal(t, head.TargetDateTime, "2011-03-11T14:58:00+09:00")
		require.Equal(t, head.EventID, "20110311144640")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "2")
		require.Equal(t, head.InfoKind, "津波情報")
		require.Equal(t, head.InfoKindVersion, "1.0_1")

		require.Equal(t, head.Headline.Text, "１１日１４時５８分現在の、津波の観測値をお知らせします。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := tsunami.Body

		t.Run("tsunami", func(t *testing.T) {
			bodyTsunami := body.Tsunami

			require.Nil(t, bodyTsunami.Estimation)

			// Observation
			require.Len(t, bodyTsunami.Observation.CodeDefine.Types, 2)
			require.Equal(t, bodyTsunami.Observation.CodeDefine.Types[0].Value, "津波予報区")
			require.Equal(t, bodyTsunami.Observation.CodeDefine.Types[0].XPath, "Item/Area/Code")

			require.Len(t, bodyTsunami.Observation.Item, 1)
			require.Equal(t, bodyTsunami.Observation.Item[0].Area.Name, "岩手県")
			require.Equal(t, bodyTsunami.Observation.Item[0].Area.Code, "210")

			require.Len(t, bodyTsunami.Observation.Item[0].Station, 1)
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].Name, "大船渡")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].Code, "21002")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].FirstHeight.ArrivalTime, "2011-03-11T14:46:00+09:00")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].FirstHeight.Initial, "引き")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].FirstHeight.Condition, "")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].FirstHeight.Revise, "")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].MaxHeight.ArrivalTime, "")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].MaxHeight.Condition, "観測中")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].MaxHeight.Revise, "")

			require.Nil(t, bodyTsunami.Observation.Item[0].Station[0].MaxHeight.TsunamiHeight)

			// CodeDefine
			require.Len(t, bodyTsunami.Forecast.CodeDefine.Types, 4)
			require.Equal(t, bodyTsunami.Forecast.CodeDefine.Types[0].XPath, "Item/Area/Code")
			require.Equal(t, bodyTsunami.Forecast.CodeDefine.Types[0].Value, "津波予報区")

			// Item
			require.Len(t, bodyTsunami.Forecast.Item, 43)
			require.Equal(t, bodyTsunami.Forecast.Item[0].Area.Name, "青森県太平洋沿岸")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Area.Code, "201")

			require.Equal(t, bodyTsunami.Forecast.Item[0].Category.Kind.Name, "大津波警報：発表")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Category.Kind.Code, "53")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Category.LastKind.Name, "津波警報")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Category.LastKind.Code, "51")

			require.Equal(t, bodyTsunami.Forecast.Item[0].FirstHeight.Condition, "")
			require.Equal(t, bodyTsunami.Forecast.Item[0].FirstHeight.ArrivalTime, "2011-03-11T15:30:00+09:00")
			require.Equal(t, bodyTsunami.Forecast.Item[0].FirstHeight.Revise, "")

			require.NotNil(t, bodyTsunami.Forecast.Item[0].MaxHeight)
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.Condition, "")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.Revise, "")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Value, "NaN")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Type, "津波の高さ")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Unit, "m")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Description, "巨大")
			require.Equal(t, bodyTsunami.Forecast.Item[0].MaxHeight.TsunamiHeight.Condition, "不明")

			require.Len(t, bodyTsunami.Forecast.Item[0].Station, 3)
			require.Equal(t, bodyTsunami.Forecast.Item[0].Station[0].Name, "八戸")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Station[0].Code, "20101")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Station[0].HighTideDateTime, "2011-03-11T19:30:00+09:00")
			require.Equal(t, bodyTsunami.Forecast.Item[0].Station[0].FirstHeight.ArrivalTime, "2011-03-11T15:50:00+09:00")
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
			require.Equal(t, bodyEarthquake[0].Magnitude.Type, jma.JMAMagnitude)
			require.Equal(t, bodyEarthquake[0].Magnitude.Description, "Ｍ８を超える巨大地震")
			require.Equal(t, bodyEarthquake[0].Magnitude.Condition, "不明")
		})

		t.Run("comments", func(t *testing.T) {
			bodyComments := body.Comments

			require.Equal(t, bodyComments.WarningComment.Text, `津波による潮位変化が観測されてから最大波が観測されるまでに数時間以上かかることがあります。

場所によっては、観測した津波の高さよりさらに大きな津波が到達しているおそれがあります。

今後、津波の高さは更に高くなることも考えられます。`)
			require.Equal(t, bodyComments.WarningComment.CodeType, "固定付加文")
			require.Equal(t, bodyComments.WarningComment.Code, "0114 0111 0112")

			require.Equal(t, bodyComments.FreeFormComment, "")
		})

		t.Run("other", func(t *testing.T) {
			require.Equal(t, body.Text, "")
		})
	})
}

// 津波情報 のパーステスト2 （取り消し）
func TestParseTsunamiInfo3(t *testing.T) {
	target := "38-39_03_03_210805_VTSE51.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	tsunami, err := jma.ParseTsunami(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := tsunami.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "津波情報a",
			DateTime:         "2021-08-05T03:56:58Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := tsunami.Head

		require.Equal(t, head.Title, "津波観測に関する情報")
		require.Equal(t, head.ReportDateTime, "2021-08-05T12:56:00+09:00")
		require.Equal(t, head.TargetDateTime, "2021-08-05T13:05:00+09:00")
		require.Equal(t, head.EventID, "20210805103531")
		require.Equal(t, head.InfoType, jma.Cancel)
		require.Equal(t, head.Serial, "1")
		require.Equal(t, head.InfoKind, "津波情報")
		require.Equal(t, head.InfoKindVersion, "1.0_1")

		require.Equal(t, head.Headline.Text, "津波観測に関する情報を取り消します。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := tsunami.Body

		t.Run("other", func(t *testing.T) {
			require.Equal(t, body.Text, "先ほどの、津波観測に関する情報を取り消します。")
		})
	})
}

// 沖合の津波観測に関する情報 のパーステスト
func TestParseTsunamiOffshoreInfo(t *testing.T) {
	target := "32-39_12_05_191025_VTSE52.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	tsunami, err := src.ParseTsunami(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := tsunami.Parsed.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "沖合の津波観測に関する情報",
			DateTime:         "2016-08-31T22:15:30Z",
			Status:           jma.Training,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := tsunami.Parsed.Head

		require.Equal(t, head.Title, "沖合の津波観測に関する情報")
		require.Equal(t, head.ReportDateTime, "2016-09-01T07:15:00+09:00")
		require.Equal(t, head.TargetDateTime, "2016-09-01T07:15:00+09:00")
		require.Equal(t, head.EventID, "20160901071000")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "1")
		require.Equal(t, head.InfoKind, "津波情報")
		require.Equal(t, head.InfoKindVersion, "1.0_1")

		require.Equal(t, head.Headline.Text, `＊＊＊これは訓練です＊＊＊
高い津波を沖合で観測しました。`)

		// Information
		require.Len(t, head.Headline.Information, 1)
		require.Equal(t, head.Headline.Information[0].Type, "沖合の津波観測に関する情報")
		require.Len(t, head.Headline.Information[0].Item, 1)
		require.Equal(t, head.Headline.Information[0].Item[0].Kind.Code, "")
		require.Equal(t, head.Headline.Information[0].Item[0].Kind.Name, "沖合の津波観測に関する情報")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.CodeType, "潮位観測点")
		require.Len(t, head.Headline.Information[0].Item[0].Areas.Area, 8)
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Name, "静岡御前崎沖")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Code, "38090")
	})

	t.Run("body", func(t *testing.T) {
		body := tsunami.Parsed.Body

		t.Run("tsunami", func(t *testing.T) {
			bodyTsunami := body.Tsunami

			// Observation
			require.Len(t, bodyTsunami.Observation.CodeDefine.Types, 2)
			require.Equal(t, bodyTsunami.Observation.CodeDefine.Types[0].Value, "津波予報区")
			require.Equal(t, bodyTsunami.Observation.CodeDefine.Types[0].XPath, "Item/Area/Code")

			require.Len(t, bodyTsunami.Observation.Item, 1)
			require.Equal(t, bodyTsunami.Observation.Item[0].Area.Name, "")
			require.Equal(t, bodyTsunami.Observation.Item[0].Area.Code, "")

			require.Len(t, bodyTsunami.Observation.Item[0].Station, 8)
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].Name, "静岡御前崎沖")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].Code, "38090")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].Sensor, "ＧＰＳ波浪計")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].FirstHeight.ArrivalTime, "2016-09-01T07:10:00+09:00")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].FirstHeight.Initial, "押し")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].FirstHeight.Condition, "")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].FirstHeight.Revise, "")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].MaxHeight.DateTime, "2016-09-01T07:15:00+09:00")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].MaxHeight.Condition, "重要")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].MaxHeight.Revise, "")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].MaxHeight.TsunamiHeight.Condition, "")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].MaxHeight.TsunamiHeight.Type, "これまでの最大波の高さ")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].MaxHeight.TsunamiHeight.Unit, "m")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].MaxHeight.TsunamiHeight.Description, "１．８ｍ")
			require.Equal(t, bodyTsunami.Observation.Item[0].Station[0].MaxHeight.TsunamiHeight.Value, "1.8")

			require.Nil(t, bodyTsunami.Forecast)

			require.NotNil(t, bodyTsunami.Estimation)
			require.Len(t, bodyTsunami.Estimation.CodeDefine.Types, 1)
			require.Equal(t, bodyTsunami.Estimation.CodeDefine.Types[0].XPath, "Item/Area/Code")
			require.Equal(t, bodyTsunami.Estimation.CodeDefine.Types[0].Value, "沿岸地域")

			require.Len(t, bodyTsunami.Estimation.Item, 5)
			require.Equal(t, bodyTsunami.Estimation.Item[0].Area.Name, "静岡県")
			require.Equal(t, bodyTsunami.Estimation.Item[0].Area.Code, "380")
			require.Equal(t, bodyTsunami.Estimation.Item[0].FirstHeight.ArrivalTime, "2016-09-01T07:17:00+09:00")
			require.Equal(t, bodyTsunami.Estimation.Item[0].FirstHeight.Condition, "早いところでは既に津波到達と推定")
			require.Equal(t, bodyTsunami.Estimation.Item[0].MaxHeight.Condition, "重要")
			require.Equal(t, bodyTsunami.Estimation.Item[0].MaxHeight.DateTime, "2016-09-01T07:22:00+09:00")
			require.Equal(t, bodyTsunami.Estimation.Item[0].MaxHeight.Revise, "")
			require.Equal(t, bodyTsunami.Estimation.Item[0].MaxHeight.TsunamiHeight.Type, "津波の高さ")
			require.Equal(t, bodyTsunami.Estimation.Item[0].MaxHeight.TsunamiHeight.Unit, "m")
			require.Equal(t, bodyTsunami.Estimation.Item[0].MaxHeight.TsunamiHeight.Condition, "不明")
			require.Equal(t, bodyTsunami.Estimation.Item[0].MaxHeight.TsunamiHeight.Description, "巨大")
			require.Equal(t, bodyTsunami.Estimation.Item[0].MaxHeight.TsunamiHeight.Value, "NaN")
		})

		t.Run("earthquake", func(t *testing.T) {
			bodyEarthquake := body.Earthquake

			require.Len(t, bodyEarthquake, 1)
			require.Equal(t, bodyEarthquake[0].OriginTime, "2016-09-01T07:10:00+09:00")
			require.Equal(t, bodyEarthquake[0].ArrivalTime, "2016-09-01T07:10:00+09:00")

			// Hypocenter
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Name, "和歌山県南方沖")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Code.Value, "689")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Coordinate.Value, "+33.2+136.0-10000/")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Coordinate.Description, "北緯３３．２度　東経１３６．０度　深さ　１０ｋｍ")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Coordinate.Datum, "日本測地系")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.NameFromMark, "潮岬の南東３０ｋｍ付近")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.MarkCode.Type, "震央補助")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.MarkCode.Value, "502")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Direction, "南東")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Distance.Value, "30")
			require.Equal(t, bodyEarthquake[0].Hypocenter.Area.Distance.Unit, "km")

			require.Equal(t, bodyEarthquake[0].Hypocenter.Source, "")

			require.Equal(t, bodyEarthquake[0].Magnitude.Value, "NaN")
			require.Equal(t, bodyEarthquake[0].Magnitude.Type, jma.JMAMagnitude)
			require.Equal(t, bodyEarthquake[0].Magnitude.Description, "Ｍ８を超える巨大地震")
			require.Equal(t, bodyEarthquake[0].Magnitude.Condition, "不明")
		})

		t.Run("comments", func(t *testing.T) {
			bodyComments := body.Comments

			require.Equal(t, bodyComments.WarningComment.Text, "沖合での観測値であり、沿岸では津波はさらに高くなります。")
			require.Equal(t, bodyComments.WarningComment.CodeType, "固定付加文")
			require.Equal(t, bodyComments.WarningComment.Code, "0115")

			require.Equal(t, bodyComments.FreeFormComment, "")
		})

		t.Run("other", func(t *testing.T) {
			require.Equal(t, body.Text, "")
		})
	})
}
